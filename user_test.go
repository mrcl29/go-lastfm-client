package golastfmclient_test

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mrcl29/go-lastfm-client"
	"github.com/stretchr/testify/assert"
)

func TestUserService_GetInfo(t *testing.T) {
	resp := `{
		"user": {
			"name": "rj",
			"playcount": "12345"
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "user.getInfo",
		Params: url.Values{
			"user": {"rj"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewUserService(mock)
	res, err := service.GetInfo(context.Background(), "rj", nil)

	assert.NoError(t, err)
	assert.Equal(t, "rj", res.User.Name)
	assert.Equal(t, json.Number("12345"), res.User.Playcount)
}

func TestUserService_GetRecentTracks(t *testing.T) {
	resp := `{
		"recenttracks": {
			"track": [
				{"name": "Believe", "artist": {"name": "Cher"}}
			],
			"@attr": {"user": "rj"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "user.getRecentTracks",
		Params: url.Values{
			"user": {"rj"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewUserService(mock)
	res, err := service.GetRecentTracks(context.Background(), "rj", nil)

	assert.NoError(t, err)
	assert.Len(t, res.RecentTracks.Track, 1)
	assert.Equal(t, "Believe", res.RecentTracks.Track[0].Name)
}

func TestUserService_GetTopArtists(t *testing.T) {
	resp := `{
		"topartists": {
			"artist": [
				{"name": "Cher", "playcount": "100"}
			],
			"@attr": {"user": "rj"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "user.getTopArtists",
		Params: url.Values{
			"user": {"rj"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewUserService(mock)
	res, err := service.GetTopArtists(context.Background(), "rj", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Cher", res.TopArtists.Artist[0].Name)
	assert.Equal(t, json.Number("100"), res.TopArtists.Artist[0].Playcount)
}

func TestUserService_GetPersonalTags(t *testing.T) {
	resp := `{
		"taggings": {
			"artists": {
				"artist": [{"name": "Cher"}]
			},
			"@attr": {"user": "rj", "tag": "diva"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "user.getPersonalTags",
		Params: url.Values{
			"user":        {"rj"},
			"tag":         {"diva"},
			"taggingtype": {"artist"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewUserService(mock)
	res, err := service.GetPersonalTags(context.Background(), "rj", "diva", "artist", nil)

	assert.NoError(t, err)
	assert.NotNil(t, res.Taggings.Artists)
	assert.Equal(t, "Cher", res.Taggings.Artists.Artist[0].Name)
}

func TestUserService_GetTopTags(t *testing.T) {
	resp := `{
		"toptags": {
			"tag": [{"name": "pop", "count": "100"}],
			"@attr": {"user": "rj"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "user.getTopTags",
		Params: url.Values{
			"user": {"rj"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewUserService(mock)
	res, err := service.GetTopTags(context.Background(), "rj", nil)

	assert.NoError(t, err)
	assert.Len(t, res.TopTags.Tag, 1)
	assert.Equal(t, "pop", res.TopTags.Tag[0].Name)
}

func TestUserService_GetWeeklyChartList(t *testing.T) {
	resp := `{
		"weeklychartlist": {
			"chart": [{"from": "1108296002", "to": "1108900802"}],
			"@attr": {"user": "rj"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "user.getWeeklyChartList",
		Params: url.Values{
			"user": {"rj"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewUserService(mock)
	res, err := service.GetWeeklyChartList(context.Background(), "rj")

	assert.NoError(t, err)
	assert.Len(t, res.WeeklyChartList.Chart, 1)
	assert.Equal(t, "1108296002", res.WeeklyChartList.Chart[0].From)
}
