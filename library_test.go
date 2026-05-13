package golastfmclient_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/mrcl29/go-lastfm-client"
	"github.com/stretchr/testify/assert"
)

func TestLibraryService_GetArtists(t *testing.T) {
	resp := `{
		"artists": {
			"artist": [
				{"name": "Dream Theater", "playcount": "1346"}
			],
			"@attr": {
				"user": "RJ",
				"page": "1",
				"perPage": "50",
				"totalPages": "20",
				"total": "1000"
			}
		}
	}`

	mock := &golastfmclient.MockClient{
		T:          t,
		HTTPMethod: "GET",
		APIMethod:  "library.getArtists",
		Params: url.Values{
			"user": {"RJ"},
		},
		RespJSON: resp,
	}

	service := golastfmclient.NewLibraryService(mock)
	res, err := service.GetArtists(context.Background(), "RJ", nil)

	assert.NoError(t, err)
	assert.Len(t, res.Artists.Artist, 1)
	assert.Equal(t, "Dream Theater", res.Artists.Artist[0].Name)
	assert.Equal(t, "1346", res.Artists.Artist[0].Playcount.String())
}
