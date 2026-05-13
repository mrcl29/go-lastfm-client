package chart

import (
	"encoding/json"

	"github.com/mrcl29/go-lastfm-client/internal/jsonutil"
)

// Artist represents an artist in the chart.
type Artist struct {
	Name       string      `json:"name"`
	Playcount  json.Number `json:"playcount"`
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

// Tag represents a tag in the chart.
type Tag struct {
	Name       string      `json:"name"`
	URL        string      `json:"url"`
	Reach      json.Number `json:"reach"`
	Taggings   json.Number `json:"taggings"`
	Streamable string      `json:"streamable"`
	Wiki       Wiki        `json:"wiki"`
}

// TagList is a slice of Tag that handles both a single object and an array in JSON.
type TagList []Tag

// UnmarshalJSON implements json.Unmarshaler.
func (t *TagList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Tag)(t))
}

// Track represents a track in the chart.
type Track struct {
	Name       string      `json:"name"`
	Duration   json.Number `json:"duration"`
	Playcount  json.Number `json:"playcount"`
	Listeners  json.Number `json:"listeners"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Streamable Streamable  `json:"streamable"`
	Artist     Artist      `json:"artist"`
	Image      []Image     `json:"image"`
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

// Wiki represents the wiki information for a tag.
type Wiki struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

// Streamable represents streamable status.
type Streamable struct {
	Text      string `json:"#text"`
	FullTrack string `json:"fulltrack"`
}

// GetTopArtistsResponse is the response from chart.getTopArtists.
type GetTopArtistsResponse struct {
	Artists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"artists"`
}

// GetTopTagsResponse is the response from chart.getTopTags.
type GetTopTagsResponse struct {
	Tags struct {
		Tag  TagList `json:"tag"`
		Attr Attr    `json:"@attr"`
	} `json:"tags"`
}

// GetTopTracksResponse is the response from chart.getTopTracks.
type GetTopTracksResponse struct {
	Tracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"tracks"`
}

// Attr represents metadata for chart responses.
type Attr struct {
	Page       json.Number `json:"page"`
	PerPage    json.Number `json:"perPage"`
	TotalPages json.Number `json:"totalPages"`
	Total      json.Number `json:"total"`
}
