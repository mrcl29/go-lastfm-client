package geo

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
		return &mockError{Message: "Error"}
	}

	return json.Unmarshal([]byte(m.respJSON), target)
}

type mockError struct {
	Message string
}

func (e *mockError) Error() string {
	return e.Message
}

func TestService_GetTopArtists(t *testing.T) {
	resp := `{
		"topartists": {
			"artist": [{"name": "Cher", "listeners": "100"}],
			"@attr": {"country": "Spain"}
		}
	}`
	
	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "geo.getTopArtists",
		params: url.Values{
			"country": {"Spain"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetTopArtists(context.Background(), "Spain", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Cher", res.TopArtists.Artist[0].Name)
	assert.Equal(t, "Spain", res.TopArtists.Attr.Country)
}

func TestService_GetTopTracks(t *testing.T) {
	resp := `{
		"tracks": {
			"track": [{"name": "Believe", "artist": {"name": "Cher"}}],
			"@attr": {"country": "Spain"}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "geo.getTopTracks",
		params: url.Values{
			"country": {"Spain"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetTopTracks(context.Background(), "Spain", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Believe", res.Tracks.Track[0].Name)
	assert.Equal(t, "Spain", res.Tracks.Attr.Country)
}
