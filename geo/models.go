package geo

import "encoding/json"

// Artist represents an artist in a geo response.
type Artist struct {
	Name       string      `json:"name"`
	Listeners  json.Number `json:"listeners"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Streamable string      `json:"streamable"`
	Image      []Image     `json:"image"`
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
		Artist []Artist `json:"artist"`
		Attr   Attr     `json:"@attr"`
	} `json:"topartists"`
}

// GetTopTracksResponse is the response from geo.getTopTracks.
type GetTopTracksResponse struct {
	Tracks struct {
		Track []Track `json:"track"`
		Attr  Attr    `json:"@attr"`
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
