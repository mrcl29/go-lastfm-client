package geo

import (
	"context"
	"net/url"
)

// APIClient defines the interface required by the geo service.
type APIClient interface {
	Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error
}

// Service handles API calls related to geography.
type Service struct {
	client APIClient
}

// New creates a new geo service.
func New(client APIClient) *Service {
	return &Service{client: client}
}

// GetTopArtists gets the most popular artists on Last.fm by country.
// See: http://www.last.fm/api/show/geo.getTopArtists
func (s *Service) GetTopArtists(ctx context.Context, country string, options url.Values) (*GetTopArtistsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("country", country)

	var resp GetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "geo.getTopArtists", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTracks gets the most popular tracks on Last.fm last week by country.
// See: http://www.last.fm/api/show/geo.getTopTracks
func (s *Service) GetTopTracks(ctx context.Context, country string, options url.Values) (*GetTopTracksResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("country", country)

	var resp GetTopTracksResponse
	err := s.client.Call(ctx, "GET", "geo.getTopTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
