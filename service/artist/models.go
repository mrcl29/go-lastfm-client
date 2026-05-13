package artist

import (
	"encoding/json"

	"github.com/mrcl29/go-lastfm-client/internal/jsonutil"
)

// Artist represents a Last.fm artist.
type Artist struct {
	Name       string      `json:"name"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Image      []Image     `json:"image"`
	Streamable string      `json:"streamable"`
	Stats      Stats       `json:"stats"`
	Similar    Similar     `json:"similar"`
	Tags       Tags        `json:"tags"`
	Bio        Bio         `json:"bio"`
	Match      float64     `json:"match,string,omitempty"`
}

// Image represents a Last.fm image.
type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

// Stats represents statistics for an artist.
type Stats struct {
	Listeners json.Number `json:"listeners"`
	Playcount json.Number `json:"playcount"`
}

// Similar represents similar artists.
type Similar struct {
	Artist ArtistList `json:"artist"`
}

// ArtistList is a slice of Artist that handles both a single object and an array in JSON.
type ArtistList []Artist

// UnmarshalJSON implements json.Unmarshaler.
func (a *ArtistList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Artist)(a))
}

// Tags represents tags for an artist.
type Tags struct {
	Tag TagList `json:"tag"`
}

// TagList is a slice of Tag that handles both a single object and an array in JSON.
type TagList []Tag

// UnmarshalJSON implements json.Unmarshaler.
func (t *TagList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Tag)(t))
}

// Tag represents a single tag.
type Tag struct {
	Name  string      `json:"name"`
	URL   string      `json:"url"`
	Count json.Number `json:"count,omitempty"`
}

// Bio represents biography information for an artist.
type Bio struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

// GetInfoResponse is the response from artist.getInfo.
type GetInfoResponse struct {
	Artist Artist `json:"artist"`
}

// SearchResponse is the response from artist.search.
type SearchResponse struct {
	Results struct {
		ArtistMatches struct {
			Artist ArtistList `json:"artist"`
		} `json:"artistmatches"`
	} `json:"results"`
}

// GetSimilarResponse is the response from artist.getSimilar.
type GetSimilarResponse struct {
	SimilarArtists struct {
		Artist ArtistList `json:"artist"`
		Attr   struct {
			Artist string `json:"artist"`
		} `json:"@attr"`
	} `json:"similarartists"`
}

// GetCorrectionResponse is the response from artist.getCorrection.
type GetCorrectionResponse struct {
	Corrections struct {
		Correction Correction `json:"correction"`
	} `json:"corrections"`
}

// Correction represents an artist correction.
type Correction struct {
	Artist Artist `json:"artist"`
	Index  string `json:"index"`
}

// GetTopAlbumsResponse is the response from artist.getTopAlbums.
type GetTopAlbumsResponse struct {
	TopAlbums struct {
		Album AlbumList `json:"album"`
		Attr  struct {
			Artist string `json:"artist"`
			Page   string `json:"page"`
			Total  string `json:"total"`
		} `json:"@attr"`
	} `json:"topalbums"`
}

// AlbumList is a slice of Album that handles both a single object and an array in JSON.
type AlbumList []Album

// UnmarshalJSON implements json.Unmarshaler.
func (a *AlbumList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Album)(a))
}

// Album represents an album in an artist's top albums list.
type Album struct {
	Name      string      `json:"name"`
	Playcount json.Number `json:"playcount"`
	MBID      string      `json:"mbid"`
	URL       string      `json:"url"`
	Artist    Artist      `json:"artist"`
	Image     []Image     `json:"image"`
}

// GetTopTracksResponse is the response from artist.getTopTracks.
type GetTopTracksResponse struct {
	TopTracks struct {
		Track TrackList `json:"track"`
		Attr  struct {
			Artist string `json:"artist"`
			Page   string `json:"page"`
			Total  string `json:"total"`
		} `json:"@attr"`
	} `json:"toptracks"`
}

// TrackList is a slice of Track that handles both a single object and an array in JSON.
type TrackList []Track

// UnmarshalJSON implements json.Unmarshaler.
func (t *TrackList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Track)(t))
}

// Track represents a track in an artist's top tracks list.
type Track struct {
	Name      string      `json:"name"`
	Playcount json.Number `json:"playcount"`
	Listeners json.Number `json:"listeners"`
	MBID      string      `json:"mbid"`
	URL       string      `json:"url"`
	Artist    Artist      `json:"artist"`
	Image     []Image     `json:"image"`
}

// GetTagsResponse is the response from artist.getTags.
type GetTagsResponse struct {
	Tags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
	} `json:"tags"`
}

// GetTopTagsResponse is the response from artist.getTopTags.
type GetTopTagsResponse struct {
	TopTags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
	} `json:"toptags"`
}
