package tag

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
		"tag": {
			"name": "disco",
			"reach": "12345"
		}
	}`
	
	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "tag.getInfo",
		params: url.Values{
			"tag": {"disco"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetInfo(context.Background(), "disco", nil)

	assert.NoError(t, err)
	assert.Equal(t, "disco", res.Tag.Name)
	assert.Equal(t, json.Number("12345"), res.Tag.Reach)
}

func TestService_GetTopAlbums(t *testing.T) {
	resp := `{
		"albums": {
			"album": [
				{"name": "Believe", "artist": {"name": "Cher"}}
			],
			"@attr": {"tag": "disco"}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "tag.getTopAlbums",
		params: url.Values{
			"tag": {"disco"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetTopAlbums(context.Background(), "disco", nil)

	assert.NoError(t, err)
	assert.Len(t, res.Albums.Album, 1)
	assert.Equal(t, "Believe", res.Albums.Album[0].Name)
}

func TestService_GetTopTags(t *testing.T) {
	resp := `{
		"toptags": {
			"tag": [
				{"name": "pop", "count": "100"}
			],
			"@attr": {"total": 1}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "tag.getTopTags",
		params:     url.Values{},
		respJSON:   resp,
	}

	service := New(mock)
	res, err := service.GetTopTags(context.Background(), nil)

	assert.NoError(t, err)
	assert.Len(t, res.TopTags.Tag, 1)
	assert.Equal(t, "pop", res.TopTags.Tag[0].Name)
}
