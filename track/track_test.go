package track

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockClient implements the APIClient interface for testing.
type MockClient struct {
	t          *testing.T
	httpMethod string
	apiMethod  string
	params     url.Values
	respJSON   string
	statusCode int
}

func (m *MockClient) Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error {
	assert.Equal(m.t, m.httpMethod, httpMethod)
	assert.Equal(m.t, m.apiMethod, apiMethod)
	
	for k, v := range m.params {
		assert.Equal(m.t, v, params[k])
	}

	if m.statusCode != 0 && m.statusCode != http.StatusOK {
		return &mockError{Code: 1, Message: "Error"}
	}

	return json.Unmarshal([]byte(m.respJSON), target)
}

type mockError struct {
	Code    int
	Message string
}

func (e *mockError) Error() string {
	return e.Message
}

func TestService_GetInfo(t *testing.T) {
	resp := `{
		"track": {
			"name": "Believe",
			"artist": {
				"name": "Cher"
			}
		}
	}`
	
	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "track.getInfo",
		params: url.Values{
			"artist": {"Cher"},
			"track":  {"Believe"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetInfo(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Believe", res.Track.Name)
	assert.Equal(t, "Cher", res.Track.Artist.Name)
}

func TestService_Search(t *testing.T) {
	resp := `{
		"results": {
			"trackmatches": {
				"track": [
					{"name": "Believe", "artist": {"name": "Cher"}}
				]
			}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "track.search",
		params: url.Values{
			"track": {"Believe"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.Search(context.Background(), "Believe", nil)

	assert.NoError(t, err)
	assert.Len(t, res.Results.TrackMatches.Track, 1)
	assert.Equal(t, "Believe", res.Results.TrackMatches.Track[0].Name)
}

func TestService_Scrobble(t *testing.T) {
	resp := `{
		"scrobbles": {
			"scrobble": [
				{
					"track": {"#text": "Believe", "corrected": "0"},
					"artist": {"#text": "Cher", "corrected": "0"}
				}
			],
			"@attr": {"accepted": 1, "ignored": 0}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "POST",
		apiMethod:  "track.scrobble",
		params: url.Values{
			"artist":    {"Cher"},
			"track":     {"Believe"},
			"timestamp": {"123456789"},
		},
		respJSON: resp,
	}

	service := New(mock)
	// Note: in track.go I used url.QueryEscape(string(timestamp)) which is WRONG for int64.
	// I should fix that in track.go
	res, err := service.Scrobble(context.Background(), "Cher", "Believe", 123456789, nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, res.Scrobbles.Attr.Accepted)
}

func TestService_UpdateNowPlaying(t *testing.T) {
	resp := `{
		"nowplaying": {
			"track": {"#text": "Believe", "corrected": "0"},
			"artist": {"#text": "Cher", "corrected": "0"}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "POST",
		apiMethod:  "track.updateNowPlaying",
		params: url.Values{
			"artist": {"Cher"},
			"track":  {"Believe"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.UpdateNowPlaying(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Believe", res.NowPlaying.Track.Text)
}
