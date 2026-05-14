package golastfmclient_test

import (
	"context"
	"net/url"
	"testing"

	golastfmclient "github.com/mrcl29/go-lastfm-client"
	"github.com/stretchr/testify/assert"
)

func TestGeoService_GetTopArtists(t *testing.T) {
	resp := `{
		"topartists": {
			"artist": [{"name": "Cher", "listeners": "100"}],
			"@attr": {"country": "Spain"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "geo.getTopArtists",
		Params: url.Values{
			"country": {"Spain"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewGeoService(mock)
	res, attr, err := service.GetTopArtists(context.Background(), "Spain", nil)

	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, "Cher", res[0].Name)
	assert.Equal(t, "Spain", attr.Country)
}

func TestGeoService_GetTopTracks(t *testing.T) {
	resp := `{
		"tracks": {
			"track": [{"name": "Believe", "artist": {"name": "Cher"}}],
			"@attr": {"country": "Spain"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "geo.getTopTracks",
		Params: url.Values{
			"country": {"Spain"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewGeoService(mock)
	res, attr, err := service.GetTopTracks(context.Background(), "Spain", nil)

	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, "Believe", res[0].Name)
	assert.Equal(t, "Spain", attr.Country)
}
