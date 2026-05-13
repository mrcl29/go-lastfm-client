package chart

import (
	"context"
	"net/url"
)

// APIClient defines the interface required by the chart service.
type APIClient interface {
	Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error
}

// Service handles API calls related to charts.
type Service struct {
	client APIClient
}

// New creates a new chart service.
func New(client APIClient) *Service {
	return &Service{client: client}
}

// GetTopArtists gets the top artists chart.
// See: http://www.last.fm/api/show/chart.getTopArtists
func (s *Service) GetTopArtists(ctx context.Context, options url.Values) (*GetTopArtistsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}

	var resp GetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "chart.getTopArtists", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTags gets the top tags chart.
// See: http://www.last.fm/api/show/chart.getTopTags
func (s *Service) GetTopTags(ctx context.Context, options url.Values) (*GetTopTagsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}

	var resp GetTopTagsResponse
	err := s.client.Call(ctx, "GET", "chart.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTracks gets the top tracks chart.
// See: http://www.last.fm/api/show/chart.getTopTracks
func (s *Service) GetTopTracks(ctx context.Context, options url.Values) (*GetTopTracksResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}

	var resp GetTopTracksResponse
	err := s.client.Call(ctx, "GET", "chart.getTopTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
