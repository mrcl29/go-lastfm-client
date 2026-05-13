package lastfm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthService_GetToken(t *testing.T) {
	is := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		is.Equal("auth.getToken", r.URL.Query().Get("method"))
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"token": "test_token",
		})
	}))
	defer server.Close()

	client := New("key", "secret", WithBaseURL(server.URL))
	token, err := client.Auth.GetToken(context.Background())

	is.NoError(err)
	is.Equal("test_token", token)
}

func TestAuthService_GetSession(t *testing.T) {
	is := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		is.Equal("auth.getSession", query.Get("method"))
		is.Equal("token_val", query.Get("token"))
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"session": map[string]string{
				"name": "user",
				"key":  "session_key",
			},
		})
	}))
	defer server.Close()

	client := New("key", "secret", WithBaseURL(server.URL))
	session, err := client.Auth.GetSession(context.Background(), "token_val")

	is.NoError(err)
	is.Equal("user", session.Name)
	is.Equal("session_key", session.Key)
}

func TestAuthService_GetMobileSession(t *testing.T) {
	is := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		is.Equal(http.MethodPost, r.Method)
		
		err := r.ParseForm()
		is.NoError(err)
		
		is.Equal("auth.getMobileSession", r.Form.Get("method"))
		is.Equal("user", r.Form.Get("username"))
		is.Equal("pass", r.Form.Get("password"))

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"session": map[string]string{
				"name": "user",
				"key":  "mobile_session_key",
			},
		})
	}))
	defer server.Close()

	client := New("key", "secret", WithBaseURL(server.URL))
	session, err := client.Auth.GetMobileSession(context.Background(), "user", "pass")

	is.NoError(err)
	is.Equal("user", session.Name)
	is.Equal("mobile_session_key", session.Key)
}

func TestAuthService_AuthURL(t *testing.T) {
	is := assert.New(t)

	client := New("test_key", "secret")
	
	is.Equal("http://www.last.fm/api/auth/?api_key=test_key", client.Auth.AuthURL(""))
	is.Equal("http://www.last.fm/api/auth/?api_key=test_key&token=t1", client.Auth.AuthURL("t1"))
}
