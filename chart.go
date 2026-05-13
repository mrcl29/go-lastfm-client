package golastfmclient

import (
	"context"
	"net/url"
)

// ChartService handles API calls related to charts.
type ChartService struct {
	client APIClient
}

// NewChartService creates a new chart service.
func NewChartService(client APIClient) *ChartService {
	return &ChartService{client: client}
}

// GetTopArtists gets the top artists chart.
// See: http://www.last.fm/api/show/chart.getTopArtists
func (s *ChartService) GetTopArtists(ctx context.Context, options url.Values) (*ChartGetTopArtistsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}

	var resp ChartGetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "chart.getTopArtists", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTags gets the top tags chart.
// See: http://www.last.fm/api/show/chart.getTopTags
func (s *ChartService) GetTopTags(ctx context.Context, options url.Values) (*ChartGetTopTagsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}

	var resp ChartGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "chart.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTracks gets the top tracks chart.
// See: http://www.last.fm/api/show/chart.getTopTracks
func (s *ChartService) GetTopTracks(ctx context.Context, options url.Values) (*ChartGetTopTracksResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}

	var resp ChartGetTopTracksResponse
	err := s.client.Call(ctx, "GET", "chart.getTopTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChartGetTopArtistsResponse is the response from chart.getTopArtists.
type ChartGetTopArtistsResponse struct {
	Artists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"artists"`
}

// ChartGetTopTagsResponse is the response from chart.getTopTags.
type ChartGetTopTagsResponse struct {
	Tags struct {
		Tag  TagList `json:"tag"`
		Attr Attr    `json:"@attr"`
	} `json:"tags"`
}

// ChartGetTopTracksResponse is the response from chart.getTopTracks.
type ChartGetTopTracksResponse struct {
	Tracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"tracks"`
}
