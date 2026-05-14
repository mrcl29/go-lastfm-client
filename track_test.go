package golastfmclient_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/mrcl29/go-lastfm-client"
	"github.com/stretchr/testify/assert"
)

func TestTrackService_GetInfo(t *testing.T) {
	resp := `{
		"track": {
			"name": "Believe",
			"artist": {
				"name": "Cher"
			}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "track.getInfo",
		Params: url.Values{
			"artist": {"Cher"},
			"track":  {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewTrackService(mock)
	track, err := service.GetInfo(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Believe", track.Name)
	// track.Artist is interface{} in unified models, so we need to check carefully or cast
	// In the JSON above, it's a map/struct.
	artistMap := track.Artist.(map[string]interface{})
	assert.Equal(t, "Cher", artistMap["name"])
}

func TestTrackService_Search(t *testing.T) {
	resp := `{
		"results": {
			"trackmatches": {
				"track": [
					{"name": "Believe", "artist": {"name": "Cher"}}
				]
			}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "track.search",
		Params: url.Values{
			"track": {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewTrackService(mock)
	tracks, err := service.Search(context.Background(), "Believe", nil)

	assert.NoError(t, err)
	assert.Len(t, tracks, 1)
	assert.Equal(t, "Believe", tracks[0].Name)
}

func TestTrackService_Scrobble(t *testing.T) {
	resp := `{
		"scrobbles": {
			"scrobble": [
				{
					"track": {"#text": "Believe", "corrected": "0"},
					"artist": {"#text": "Cher", "corrected": "0"}
				}
			],
			"@attr": {"accepted": 1, "ignored": 0}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "POST",
		APIMethod:  "track.scrobble",
		Params: url.Values{
			"artist":    {"Cher"},
			"track":     {"Believe"},
			"timestamp": {"123456789"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewTrackService(mock)
	scrobbles, err := service.Scrobble(context.Background(), "Cher", "Believe", 123456789, nil)

	assert.NoError(t, err)
	assert.Len(t, scrobbles, 1)
	assert.Equal(t, "Believe", scrobbles[0].Track.Text)
}

func TestTrackService_UpdateNowPlaying(t *testing.T) {
	resp := `{
		"nowplaying": {
			"track": {"#text": "Believe", "corrected": "0"},
			"artist": {"#text": "Cher", "corrected": "0"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "POST",
		APIMethod:  "track.updateNowPlaying",
		Params: url.Values{
			"artist": {"Cher"},
			"track":  {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewTrackService(mock)
	res, err := service.UpdateNowPlaying(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Equal(t, "Believe", res.Track.Text)
}

func TestTrackService_Love(t *testing.T) {
	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "POST",
		APIMethod:  "track.love",
		Params: url.Values{
			"artist": {"Cher"},
			"track":  {"Believe"},
		},
		RespJSON: `{}`,
	}

	service := golastfmclient.NewTrackService(mock)
	err := service.Love(context.Background(), "Cher", "Believe")
	assert.NoError(t, err)
}

func TestTrackService_Unlove(t *testing.T) {
	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "POST",
		APIMethod:  "track.unlove",
		Params: url.Values{
			"artist": {"Cher"},
			"track":  {"Believe"},
		},
		RespJSON: `{}`,
	}

	service := golastfmclient.NewTrackService(mock)
	err := service.Unlove(context.Background(), "Cher", "Believe")
	assert.NoError(t, err)
}

func TestTrackService_GetTags(t *testing.T) {
	resp := `{
		"tags": {
			"tag": [{"name": "pop", "url": "..."}],
			"artist": "Cher",
			"track": "Believe"
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "track.getTags",
		Params: url.Values{
			"artist": {"Cher"},
			"track":  {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewTrackService(mock)
	tags, err := service.GetTags(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, "pop", tags[0].Name)
}

func TestTrackService_GetTopTags(t *testing.T) {
	resp := `{
		"toptags": {
			"tag": [{"name": "pop", "url": "..."}],
			"artist": "Cher",
			"track": "Believe"
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "track.getTopTags",
		Params: url.Values{
			"artist": {"Cher"},
			"track":  {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewTrackService(mock)
	tags, err := service.GetTopTags(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, "pop", tags[0].Name)
}

func TestTrackService_GetCorrection(t *testing.T) {
	resp := `{
		"corrections": {
			"correction": {
				"track": {"name": "Believe", "artist": {"name": "Cher"}}
			}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "track.getCorrection",
		Params: url.Values{
			"artist": {"Cher"},
			"track":  {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewTrackService(mock)
	track, err := service.GetCorrection(context.Background(), "Cher", "Believe")

	assert.NoError(t, err)
	assert.Equal(t, "Believe", track.Name)
}

func TestTrackService_GetSimilar(t *testing.T) {
	resp := `{
		"similartracks": {
			"track": [
				{"name": "Strong Enough", "artist": {"name": "Cher"}}
			],
			"@attr": {"artist": "Cher"}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "track.getSimilar",
		Params: url.Values{
			"artist": {"Cher"},
			"track":  {"Believe"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewTrackService(mock)
	tracks, err := service.GetSimilar(context.Background(), "Cher", "Believe", nil)

	assert.NoError(t, err)
	assert.Len(t, tracks, 1)
	assert.Equal(t, "Strong Enough", tracks[0].Name)
}
