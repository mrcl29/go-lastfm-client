package golastfmclient_test

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"

	golastfmclient "github.com/mrcl29/go-lastfm-client"
	"github.com/stretchr/testify/assert"
)

func TestChartService_GetTopArtists(t *testing.T) {
	resp := `{
		"artists": {
			"artist": [{"name": "Cher", "listeners": "100"}],
			"@attr": {"total": "1"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "chart.getTopArtists",
		Params:     url.Values{},
		RespJSON:   resp,
	}

	service := golastfmclient.NewChartService(mock)
	res, attr, err := service.GetTopArtists(context.Background(), nil)

	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, "Cher", res[0].Name)
	assert.Equal(t, json.Number("100"), res[0].Listeners)
	assert.Equal(t, "1", attr.Total.String())
}

func TestChartService_GetTopTags(t *testing.T) {
	resp := `{
		"tags": {
			"tag": [{"name": "pop", "taggings": "1000"}],
			"@attr": {"total": "1"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "chart.getTopTags",
		Params:     url.Values{},
		RespJSON:   resp,
	}

	service := golastfmclient.NewChartService(mock)
	res, attr, err := service.GetTopTags(context.Background(), nil)

	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, "pop", res[0].Name)
	assert.Equal(t, "1", attr.Total.String())
}

func TestChartService_GetTopTracks(t *testing.T) {
	resp := `{
		"tracks": {
			"track": [{"name": "Believe", "artist": {"name": "Cher"}}],
			"@attr": {"total": "1"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "chart.getTopTracks",
		Params:     url.Values{},
		RespJSON:   resp,
	}

	service := golastfmclient.NewChartService(mock)
	res, attr, err := service.GetTopTracks(context.Background(), nil)

	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, "Believe", res[0].Name)
	artistMap := res[0].Artist.(map[string]interface{})
	assert.Equal(t, "Cher", artistMap["name"])
	assert.Equal(t, "1", attr.Total.String())
}
