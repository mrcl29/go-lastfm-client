package golastfmclient_test

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mrcl29/go-lastfm-client"
	"github.com/stretchr/testify/assert"
)

func TestArtistService_GetInfo(t *testing.T) {
	resp := `{
		"artist": {
			"name": "Cher",
			"stats": {
				"listeners": "100",
				"playcount": "1000"
			}
		}
	}`
	
	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "artist.getInfo",
		Params: url.Values{
			"artist": {"Cher"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewArtistService(mock)
	res, err := service.GetInfo(context.Background(), "Cher", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Cher", res.Artist.Name)
	assert.Equal(t, json.Number("100"), res.Artist.Stats.Listeners)
}

func TestArtistService_Search(t *testing.T) {
	resp := `{
		"results": {
			"artistmatches": {
				"artist": [
					{"name": "Cher"}
				]
			}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "artist.search",
		Params: url.Values{
			"artist": {"Cher"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewArtistService(mock)
	res, err := service.Search(context.Background(), "Cher", nil)

	assert.NoError(t, err)
	assert.Len(t, res.Results.ArtistMatches.Artist, 1)
	assert.Equal(t, "Cher", res.Results.ArtistMatches.Artist[0].Name)
}

func TestArtistService_GetSimilar(t *testing.T) {
	resp := `{
		"similarartists": {
			"artist": [
				{"name": "Madonna", "match": "1"}
			],
			"@attr": {"artist": "Cher"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "artist.getSimilar",
		Params: url.Values{
			"artist": {"Cher"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewArtistService(mock)
	res, err := service.GetSimilar(context.Background(), "Cher", nil)

	assert.NoError(t, err)
	assert.Len(t, res.SimilarArtists.Artist, 1)
	assert.Equal(t, "Madonna", res.SimilarArtists.Artist[0].Name)
}

func TestArtistService_GetCorrection(t *testing.T) {
	resp := `{
		"corrections": {
			"correction": {
				"artist": {"name": "Guns N' Roses"}
			}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "artist.getCorrection",
		Params: url.Values{
			"artist": {"Guns and Roses"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewArtistService(mock)
	res, err := service.GetCorrection(context.Background(), "Guns and Roses")

	assert.NoError(t, err)
	assert.Equal(t, "Guns N' Roses", res.Corrections.Correction.Artist.Name)
}

func TestArtistService_AddTags(t *testing.T) {
	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "POST",
		APIMethod:  "artist.addTags",
		Params: url.Values{
			"artist": {"Cher"},
			"tags":   {"pop,diva"},
		},
		RespJSON: `{}`,
	}

	service := golastfmclient.NewArtistService(mock)
	err := service.AddTags(context.Background(), "Cher", "pop,diva")
	assert.NoError(t, err)
}

func TestArtistService_RemoveTag(t *testing.T) {
	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "POST",
		APIMethod:  "artist.removeTag",
		Params: url.Values{
			"artist": {"Cher"},
			"tag":    {"pop"},
		},
		RespJSON: `{}`,
	}

	service := golastfmclient.NewArtistService(mock)
	err := service.RemoveTag(context.Background(), "Cher", "pop")
	assert.NoError(t, err)
}

func TestArtistService_GetTags(t *testing.T) {
	resp := `{
		"tags": {
			"tag": [{"name": "pop", "url": "..."}],
			"artist": "Cher"
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "artist.getTags",
		Params: url.Values{
			"artist": {"Cher"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewArtistService(mock)
	res, err := service.GetTags(context.Background(), "Cher", nil)

	assert.NoError(t, err)
	assert.Len(t, res.Tags.Tag, 1)
	assert.Equal(t, "pop", res.Tags.Tag[0].Name)
}

func TestArtistService_GetTopTags(t *testing.T) {
	resp := `{
		"toptags": {
			"tag": [{"name": "pop", "count": "100"}],
			"artist": "Cher"
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "artist.getTopTags",
		Params: url.Values{
			"artist": {"Cher"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewArtistService(mock)
	res, err := service.GetTopTags(context.Background(), "Cher", nil)

	assert.NoError(t, err)
	assert.Len(t, res.TopTags.Tag, 1)
	assert.Equal(t, "pop", res.TopTags.Tag[0].Name)
}
