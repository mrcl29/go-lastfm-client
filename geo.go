package golastfmclient

import (
	"context"
	"net/url"
)

// GeoService handles API calls related to geography.
type GeoService struct {
	client APIClient
}

// NewGeoService creates a new geo service.
func NewGeoService(client APIClient) *GeoService {
	return &GeoService{client: client}
}

// GetTopArtists gets the most popular artists on Last.fm by country.
// See: http://www.last.fm/api/show/geo.getTopArtists
func (s *GeoService) GetTopArtists(ctx context.Context, country string, options url.Values) (*GeoGetTopArtistsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("country", country)

	var resp GeoGetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "geo.getTopArtists", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTracks gets the most popular tracks on Last.fm last week by country.
// See: http://www.last.fm/api/show/geo.getTopTracks
func (s *GeoService) GetTopTracks(ctx context.Context, country string, options url.Values) (*GeoGetTopTracksResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("country", country)

	var resp GeoGetTopTracksResponse
	err := s.client.Call(ctx, "GET", "geo.getTopTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GeoGetTopArtistsResponse is the response from geo.getTopArtists.
type GeoGetTopArtistsResponse struct {
	TopArtists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"topartists"`
}

// GeoGetTopTracksResponse is the response from geo.getTopTracks.
type GeoGetTopTracksResponse struct {
	Tracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"tracks"`
}
