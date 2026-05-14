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
//
// Parameters:
//   - ctx: Context for the request.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - ArtistList: A slice of top artists.
//   - *Attr: Pagination and metadata.
//   - error: Error if the request fails.
func (s *ChartService) GetTopArtists(ctx context.Context, options url.Values) (ArtistList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}

	var resp chartGetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "chart.getTopArtists", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Artists.Artist, &resp.Artists.Attr, nil
}

// GetTopTags gets the top tags chart.
// See: http://www.last.fm/api/show/chart.getTopTags
//
// Parameters:
//   - ctx: Context for the request.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - TagList: A slice of top tags.
//   - *Attr: Pagination and metadata.
//   - error: Error if the request fails.
func (s *ChartService) GetTopTags(ctx context.Context, options url.Values) (TagList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}

	var resp chartGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "chart.getTopTags", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Tags.Tag, &resp.Tags.Attr, nil
}

// GetTopTracks gets the top tracks chart.
// See: http://www.last.fm/api/show/chart.getTopTracks
//
// Parameters:
//   - ctx: Context for the request.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - TrackList: A slice of top tracks.
//   - *Attr: Pagination and metadata.
//   - error: Error if the request fails.
func (s *ChartService) GetTopTracks(ctx context.Context, options url.Values) (TrackList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}

	var resp chartGetTopTracksResponse
	err := s.client.Call(ctx, "GET", "chart.getTopTracks", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Tracks.Track, &resp.Tracks.Attr, nil
}

type chartGetTopArtistsResponse struct {
	Artists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"artists"`
}

type chartGetTopTagsResponse struct {
	Tags struct {
		Tag  TagList `json:"tag"`
		Attr Attr    `json:"@attr"`
	} `json:"tags"`
}

type chartGetTopTracksResponse struct {
	Tracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"tracks"`
}
