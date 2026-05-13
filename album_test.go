package golastfmclient_test

import (
	"context"
	"net/url"
	"testing"

	golastfmclient "github.com/mrcl29/go-lastfm-client"
	"github.com/stretchr/testify/assert"
)

func TestAlbumService_GetInfo(t *testing.T) {
	resp := `{
		"album": {
			"name": "Believe",
			"artist": "Cher",
			"playcount": "1000"
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "album.getInfo",
		Params: url.Values{
			"artist": {"Cher"},
			"album":  {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewAlbumService(mock)
	res, err := service.GetInfo(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Believe", res.Album.Name)
	assert.Equal(t, "1000", res.Album.Playcount.String())
}

func TestAlbumService_Search(t *testing.T) {
	resp := `{
		"results": {
			"albummatches": {
				"album": [
					{"name": "Believe", "artist": "Cher"}
				]
			}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "album.search",
		Params: url.Values{
			"album": {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewAlbumService(mock)
	res, err := service.Search(context.Background(), "Believe", nil)

	assert.NoError(t, err)
	assert.Len(t, res.Results.AlbumMatches.Album, 1)
	assert.Equal(t, "Believe", res.Results.AlbumMatches.Album[0].Name)
}

func TestAlbumService_GetTags(t *testing.T) {
	resp := `{
		"tags": {
			"tag": [{"name": "pop", "url": "..."}],
			"artist": "Cher",
			"album": "Believe"
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "album.getTags",
		Params: url.Values{
			"artist": {"Cher"},
			"album":  {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewAlbumService(mock)
	res, err := service.GetTags(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Equal(t, "pop", res.Tags.Tag[0].Name)
}

func TestAlbumService_GetTopTags(t *testing.T) {
	resp := `{
		"toptags": {
			"tag": [{"name": "pop", "count": "100"}],
			"artist": "Cher",
			"album": "Believe"
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "album.getTopTags",
		Params: url.Values{
			"artist": {"Cher"},
			"album":  {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewAlbumService(mock)
	res, err := service.GetTopTags(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Equal(t, "pop", res.TopTags.Tag[0].Name)
	assert.Equal(t, "100", res.TopTags.Tag[0].Count.String())
}

func TestAlbumService_AddTags(t *testing.T) {
	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "POST",
		APIMethod:  "album.addTags",
		Params: url.Values{
			"artist": {"Cher"},
			"album":  {"Believe"},
			"tags":   {"pop,rock"},
		},
		RespJSON: `{}`,
	}

	service := golastfmclient.NewAlbumService(mock)
	err := service.AddTags(context.Background(), "Cher", "Believe", "pop,rock")
	assert.NoError(t, err)
}

func TestAlbumService_RemoveTag(t *testing.T) {
	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "POST",
		APIMethod:  "album.removeTag",
		Params: url.Values{
			"artist": {"Cher"},
			"album":  {"Believe"},
			"tag":    {"rock"},
		},
		RespJSON: `{}`,
	}

	service := golastfmclient.NewAlbumService(mock)
	err := service.RemoveTag(context.Background(), "Cher", "Believe", "rock")
	assert.NoError(t, err)
}
