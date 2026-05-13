package golastfmclient_test

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mrcl29/go-lastfm-client"
	"github.com/stretchr/testify/assert"
)

func TestTagService_GetInfo(t *testing.T) {
	resp := `{
		"tag": {
			"name": "disco",
			"reach": "12345"
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "tag.getInfo",
		Params: url.Values{
			"tag": {"disco"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewTagService(mock)
	res, err := service.GetInfo(context.Background(), "disco", nil)

	assert.NoError(t, err)
	assert.Equal(t, "disco", res.Tag.Name)
	assert.Equal(t, json.Number("12345"), res.Tag.Reach)
}

func TestTagService_GetTopAlbums(t *testing.T) {
	resp := `{
		"albums": {
			"album": [
				{"name": "Believe", "artist": {"name": "Cher"}}
			],
			"@attr": {"tag": "disco"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "tag.getTopAlbums",
		Params: url.Values{
			"tag": {"disco"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewTagService(mock)
	res, err := service.GetTopAlbums(context.Background(), "disco", nil)

	assert.NoError(t, err)
	assert.Len(t, res.Albums.Album, 1)
	assert.Equal(t, "Believe", res.Albums.Album[0].Name)
}

func TestTagService_GetTopTags(t *testing.T) {
	resp := `{
		"toptags": {
			"tag": [
				{"name": "pop", "count": "100"}
			],
			"@attr": {"total": 1}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "tag.getTopTags",
		Params:     url.Values{},
		RespJSON:   resp,
	}

	service := golastfmclient.NewTagService(mock)
	res, err := service.GetTopTags(context.Background(), nil)

	assert.NoError(t, err)
	assert.Len(t, res.TopTags.Tag, 1)
	assert.Equal(t, "pop", res.TopTags.Tag[0].Name)
}
