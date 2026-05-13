package golastfmclient

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Get(t *testing.T) {
	is := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		is.Equal(http.MethodGet, r.Method)
		is.Equal("/2.0/", r.URL.Path)

		query := r.URL.Query()
		is.Equal("artist.getInfo", query.Get("method"))
		is.Equal("test_api_key", query.Get("api_key"))
		is.Equal("json", query.Get("format"))
		is.Equal("Cher", query.Get("artist"))

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"artist": map[string]string{
				"name": "Cher",
			},
		})
	}))
	defer server.Close()

	client := New("test_api_key", "test_api_secret", WithBaseURL(server.URL+"/2.0/"))

	var target map[string]interface{}
	err := client.get(context.Background(), "artist.getInfo", url.Values{"artist": {"Cher"}}, &target)

	is.NoError(err)
	is.NotNil(target)
	artist := target["artist"].(map[string]interface{})
	is.Equal("Cher", artist["name"])
}

func TestClient_Post(t *testing.T) {
	is := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		is.Equal(http.MethodPost, r.Method)
		
		err := r.ParseForm()
		is.NoError(err)

		form := r.PostForm
		is.Equal("track.scrobble", form.Get("method"))
		is.Equal("test_api_key", form.Get("api_key"))
		is.Equal("json", form.Get("format"))
		is.NotEmpty(form.Get("api_sig"))

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "ok",
		})
	}))
	defer server.Close()

	client := New("test_api_key", "test_api_secret", WithBaseURL(server.URL+"/2.0/"))

	var target map[string]interface{}
	err := client.post(context.Background(), "track.scrobble", url.Values{"track": {"Believe"}}, &target)

	is.NoError(err)
	is.Equal("ok", target["status"])
}

func TestClient_Error(t *testing.T) {
	is := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   10,
			"message": "Invalid API Key",
		})
	}))
	defer server.Close()

	client := New("invalid_key", "test_api_secret", WithBaseURL(server.URL+"/2.0/"))

	err := client.get(context.Background(), "artist.getInfo", nil, nil)

	is.Error(err)
	apiErr, ok := err.(*APIError)
	is.True(ok)
	is.Equal(10, apiErr.Code)
	is.Equal("Invalid API Key", apiErr.Message)
}

func TestClient_Sign(t *testing.T) {
	is := assert.New(t)

	client := New("test_api_key", "test_api_secret")
	params := url.Values{
		"method":  {"auth.getSession"},
		"api_key": {"test_api_key"},
		"token":   {"test_token"},
		"format":  {"json"},
	}

	sig := client.sign(params)
	
	// Alphabetical order of keys (excluding format): api_key, method, token
	// Concatenation: api_key + value + method + value + token + value + secret
	expectedInput := "api_keytest_api_keymethodauth.getSessiontokentest_tokentest_api_secret"
	h := md5.New()
	h.Write([]byte(expectedInput))
	expectedSig := fmt.Sprintf("%x", h.Sum(nil))

	is.Equal(expectedSig, sig)
}
