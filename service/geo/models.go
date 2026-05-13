package geo

import (
	"encoding/json"

	"github.com/mrcl29/go-lastfm-client/internal/jsonutil"
)

// Artist represents an artist in a geo response.
type Artist struct {
	Name       string      `json:"name"`
	Listeners  json.Number `json:"listeners"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Streamable string      `json:"streamable"`
	Image      []Image     `json:"image"`
}

// ArtistList is a slice of Artist that handles both a single object and an array in JSON.
type ArtistList []Artist

// UnmarshalJSON implements json.Unmarshaler.
func (a *ArtistList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Artist)(a))
}

// Track represents a track in a geo response.
type Track struct {
	Name       string      `json:"name"`
	Duration   json.Number `json:"duration"`
	Listeners  json.Number `json:"listeners"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Streamable Streamable  `json:"streamable"`
	Artist     Artist      `json:"artist"`
	Image      []Image     `json:"image"`
	Attr       struct {
		Rank json.Number `json:"rank"`
	} `json:"@attr"`
}

// TrackList is a slice of Track that handles both a single object and an array in JSON.
type TrackList []Track

// UnmarshalJSON implements json.Unmarshaler.
func (t *TrackList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Track)(t))
}

// Image represents a Last.fm image.
type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

// Streamable represents streamable status.
type Streamable struct {
	Text      string `json:"#text"`
	FullTrack string `json:"fulltrack"`
}

// GetTopArtistsResponse is the response from geo.getTopArtists.
type GetTopArtistsResponse struct {
	TopArtists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"topartists"`
}

// GetTopTracksResponse is the response from geo.getTopTracks.
type GetTopTracksResponse struct {
	Tracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"tracks"`
}

// Attr represents metadata for geo responses.
type Attr struct {
	Country    string      `json:"country"`
	Page       json.Number `json:"page"`
	PerPage    json.Number `json:"perPage"`
	TotalPages json.Number `json:"totalPages"`
	Total      json.Number `json:"total"`
}
