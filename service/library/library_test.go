package library

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

func TestService_GetArtists(t *testing.T) {
	resp := `{
		"artists": {
			"artist": [
				{"name": "Dream Theater", "playcount": "1346"}
			],
			"@attr": {
				"user": "RJ",
				"page": "1",
				"perPage": "50",
				"totalPages": "20",
				"total": "1000"
			}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "library.getArtists",
		params: url.Values{
			"user": {"RJ"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetArtists(context.Background(), "RJ", nil)

	assert.NoError(t, err)
	assert.Len(t, res.Artists.Artist, 1)
	assert.Equal(t, "Dream Theater", res.Artists.Artist[0].Name)
	assert.Equal(t, "1346", res.Artists.Artist[0].Playcount.String())
}
