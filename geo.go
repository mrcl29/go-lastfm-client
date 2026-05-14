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
//
// Parameters:
//   - ctx: Context for the request.
//   - country: A country name, as defined by the ISO 3166-1 country names standard.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - ArtistList: A slice of top artists.
//   - *Attr: Pagination and metadata.
//   - error: Error if the request fails.
func (s *GeoService) GetTopArtists(ctx context.Context, country string, options url.Values) (ArtistList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("country", country)

	var resp geoGetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "geo.getTopArtists", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.TopArtists.Artist, &resp.TopArtists.Attr, nil
}

// GetTopTracks gets the most popular tracks on Last.fm last week by country.
// See: http://www.last.fm/api/show/geo.getTopTracks
//
// Parameters:
//   - ctx: Context for the request.
//   - country: A country name, as defined by the ISO 3166-1 country names standard.
//   - options: Additional options (e.g. page, limit, location).
//
// Returns:
//   - TrackList: A slice of top tracks.
//   - *Attr: Pagination and metadata.
//   - error: Error if the request fails.
func (s *GeoService) GetTopTracks(ctx context.Context, country string, options url.Values) (TrackList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("country", country)

	var resp geoGetTopTracksResponse
	err := s.client.Call(ctx, "GET", "geo.getTopTracks", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Tracks.Track, &resp.Tracks.Attr, nil
}

type geoGetTopArtistsResponse struct {
	TopArtists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"topartists"`
}

type geoGetTopTracksResponse struct {
	Tracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"tracks"`
}
