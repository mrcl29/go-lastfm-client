package artist

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
		"artist": {
			"name": "Cher",
			"stats": {
				"listeners": "100",
				"playcount": "1000"
			}
		}
	}`
	
	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "artist.getInfo",
		params: url.Values{
			"artist": {"Cher"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetInfo(context.Background(), "Cher", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Cher", res.Artist.Name)
	assert.Equal(t, json.Number("100"), res.Artist.Stats.Listeners)
}

func TestService_Search(t *testing.T) {
	resp := `{
		"results": {
			"artistmatches": {
				"artist": [
					{"name": "Cher"}
				]
			}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "artist.search",
		params: url.Values{
			"artist": {"Cher"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.Search(context.Background(), "Cher", nil)

	assert.NoError(t, err)
	assert.Len(t, res.Results.ArtistMatches.Artist, 1)
	assert.Equal(t, "Cher", res.Results.ArtistMatches.Artist[0].Name)
}

func TestService_GetSimilar(t *testing.T) {
	resp := `{
		"similarartists": {
			"artist": [
				{"name": "Madonna", "match": "1"}
			],
			"@attr": {"artist": "Cher"}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "artist.getSimilar",
		params: url.Values{
			"artist": {"Cher"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetSimilar(context.Background(), "Cher", nil)

	assert.NoError(t, err)
	assert.Len(t, res.SimilarArtists.Artist, 1)
	assert.Equal(t, "Madonna", res.SimilarArtists.Artist[0].Name)
}

func TestService_GetCorrection(t *testing.T) {
	resp := `{
		"corrections": {
			"correction": {
				"artist": {"name": "Guns N' Roses"}
			}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "artist.getCorrection",
		params: url.Values{
			"artist": {"Guns and Roses"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetCorrection(context.Background(), "Guns and Roses")

	assert.NoError(t, err)
	assert.Equal(t, "Guns N' Roses", res.Corrections.Correction.Artist.Name)
}
