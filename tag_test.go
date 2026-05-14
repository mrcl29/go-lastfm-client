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
	tag, err := service.GetInfo(context.Background(), "disco", nil)

	assert.NoError(t, err)
	assert.Equal(t, "disco", tag.Name)
	assert.Equal(t, json.Number("12345"), tag.Reach)
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
	albums, attr, err := service.GetTopAlbums(context.Background(), "disco", nil)

	assert.NoError(t, err)
	assert.Len(t, albums, 1)
	assert.Equal(t, "Believe", albums[0].Name)
	assert.Equal(t, "disco", attr.Tag)
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
	tags, err := service.GetTopTags(context.Background(), nil)

	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, "pop", tags[0].Name)
}
