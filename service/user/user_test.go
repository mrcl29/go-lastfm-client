package user

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

	if target == nil {
		return nil
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
		"user": {
			"name": "rj",
			"playcount": "12345"
		}
	}`
	
	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "user.getInfo",
		params: url.Values{
			"user": {"rj"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetInfo(context.Background(), "rj", nil)

	assert.NoError(t, err)
	assert.Equal(t, "rj", res.User.Name)
	assert.Equal(t, json.Number("12345"), res.User.Playcount)
}

func TestService_GetRecentTracks(t *testing.T) {
	resp := `{
		"recenttracks": {
			"track": [
				{"name": "Believe", "artist": {"name": "Cher"}}
			],
			"@attr": {"user": "rj"}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "user.getRecentTracks",
		params: url.Values{
			"user": {"rj"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetRecentTracks(context.Background(), "rj", nil)

	assert.NoError(t, err)
	assert.Len(t, res.RecentTracks.Track, 1)
	assert.Equal(t, "Believe", res.RecentTracks.Track[0].Name)
}

func TestService_GetTopArtists(t *testing.T) {
	resp := `{
		"topartists": {
			"artist": [
				{"name": "Cher", "playcount": "100"}
			],
			"@attr": {"user": "rj"}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "user.getTopArtists",
		params: url.Values{
			"user": {"rj"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetTopArtists(context.Background(), "rj", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Cher", res.TopArtists.Artist[0].Name)
	assert.Equal(t, json.Number("100"), res.TopArtists.Artist[0].Playcount)
}

func TestService_GetPersonalTags(t *testing.T) {
	resp := `{
		"taggings": {
			"artists": {
				"artist": [{"name": "Cher"}]
			},
			"@attr": {"user": "rj", "tag": "diva"}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "user.getPersonalTags",
		params: url.Values{
			"user":        {"rj"},
			"tag":         {"diva"},
			"taggingtype": {"artist"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetPersonalTags(context.Background(), "rj", "diva", "artist", nil)

	assert.NoError(t, err)
	assert.NotNil(t, res.Taggings.Artists)
	assert.Equal(t, "Cher", res.Taggings.Artists.Artist[0].Name)
}

func TestService_GetTopTags(t *testing.T) {
	resp := `{
		"toptags": {
			"tag": [{"name": "pop", "count": "100"}],
			"@attr": {"user": "rj"}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "user.getTopTags",
		params: url.Values{
			"user": {"rj"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetTopTags(context.Background(), "rj", nil)

	assert.NoError(t, err)
	assert.Len(t, res.TopTags.Tag, 1)
	assert.Equal(t, "pop", res.TopTags.Tag[0].Name)
}

func TestService_GetWeeklyChartList(t *testing.T) {
	resp := `{
		"weeklychartlist": {
			"chart": [{"from": "1108296002", "to": "1108900802"}],
			"@attr": {"user": "rj"}
		}
	}`

	mock := &MockClient{
		t:          t,
		httpMethod: "GET",
		apiMethod:  "user.getWeeklyChartList",
		params: url.Values{
			"user": {"rj"},
		},
		respJSON: resp,
	}

	service := New(mock)
	res, err := service.GetWeeklyChartList(context.Background(), "rj")

	assert.NoError(t, err)
	assert.Len(t, res.WeeklyChartList.Chart, 1)
	assert.Equal(t, "1108296002", res.WeeklyChartList.Chart[0].From)
}
