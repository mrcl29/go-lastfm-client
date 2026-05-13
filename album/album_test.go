package album

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

func TestService_GetInfo(t *testing.T) {
	resp := `{
		"album": {
			"name": "Believe",
			"artist": "Cher",
			"playcount": "1000"
		}
	}`
	
	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "album.getInfo",
		params: url.Values{
			"artist": {"Cher"},
			"album":  {"Believe"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetInfo(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Believe", res.Album.Name)
	assert.Equal(t, "1000", res.Album.Playcount)
}

func TestService_Search(t *testing.T) {
	resp := `{
		"results": {
			"albummatches": {
				"album": [
					{"name": "Believe", "artist": "Cher"}
				]
			}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "album.search",
		params: url.Values{
			"album": {"Believe"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.Search(context.Background(), "Believe", nil)

	assert.NoError(t, err)
	assert.Len(t, res.Results.AlbumMatches.Album, 1)
	assert.Equal(t, "Believe", res.Results.AlbumMatches.Album[0].Name)
}

func TestService_GetTags(t *testing.T) {
	resp := `{
		"tags": {
			"tag": [{"name": "pop", "url": "..."}],
			"artist": "Cher",
			"album": "Believe"
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "album.getTags",
		params: url.Values{
			"artist": {"Cher"},
			"album":  {"Believe"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetTags(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Equal(t, "pop", res.Tags.Tag[0].Name)
}

func TestService_GetTopTags(t *testing.T) {
	resp := `{
		"toptags": {
			"tag": [{"name": "pop", "count": "100"}],
			"artist": "Cher",
			"album": "Believe"
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "album.getTopTags",
		params: url.Values{
			"artist": {"Cher"},
			"album":  {"Believe"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetTopTags(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Equal(t, "pop", res.TopTags.Tag[0].Name)
	assert.Equal(t, json.Number("100"), res.TopTags.Tag[0].Count)
}
