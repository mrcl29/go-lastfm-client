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
	assert.Equal(t, "rj", res.Name)
	assert.Equal(t, json.Number("12345"), res.Playcount)
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
	res, attr, err := service.GetRecentTracks(context.Background(), "rj", nil)

	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, "Believe", res[0].Name)
	assert.Equal(t, "rj", attr.User)
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
	res, attr, err := service.GetTopArtists(context.Background(), "rj", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Cher", res[0].Name)
	assert.Equal(t, json.Number("100"), res[0].Playcount)
	assert.Equal(t, "rj", attr.User)
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
	res, attr, err := service.GetPersonalTags(context.Background(), "rj", "diva", "artist", nil)

	assert.NoError(t, err)
	assert.NotNil(t, res.Artists)
	assert.Equal(t, "Cher", res.Artists[0].Name)
	assert.Equal(t, "diva", attr.Tag)
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
	assert.Len(t, res, 1)
	assert.Equal(t, "pop", res[0].Name)
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
	assert.Len(t, res, 1)
	assert.Equal(t, "1108296002", res[0].From)
}
