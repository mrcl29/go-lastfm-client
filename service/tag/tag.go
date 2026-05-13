package tag

import (
	"context"
	"net/url"
)

// APIClient defines the interface required by the tag service.
type APIClient interface {
	Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error
}

// Service handles API calls related to tags.
type Service struct {
	client APIClient
}

// New creates a new tag service.
func New(client APIClient) *Service {
	return &Service{client: client}
}

// GetInfo gets the metadata for a tag.
// See: http://www.last.fm/api/show/tag.getInfo
func (s *Service) GetInfo(ctx context.Context, tag string, options url.Values) (*GetInfoResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("tag", tag)

	var resp GetInfoResponse
	err := s.client.Call(ctx, "GET", "tag.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSimilar gets all the tags similar to this tag.
// See: http://www.last.fm/api/show/tag.getSimilar
func (s *Service) GetSimilar(ctx context.Context, tag string, options url.Values) (*GetSimilarResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("tag", tag)

	var resp GetSimilarResponse
	err := s.client.Call(ctx, "GET", "tag.getSimilar", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopAlbums gets the top albums for a tag.
// See: http://www.last.fm/api/show/tag.getTopAlbums
func (s *Service) GetTopAlbums(ctx context.Context, tag string, options url.Values) (*GetTopAlbumsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("tag", tag)

	var resp GetTopAlbumsResponse
	err := s.client.Call(ctx, "GET", "tag.getTopAlbums", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopArtists gets the top artists for a tag.
// See: http://www.last.fm/api/show/tag.getTopArtists
func (s *Service) GetTopArtists(ctx context.Context, tag string, options url.Values) (*GetTopArtistsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("tag", tag)

	var resp GetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "tag.getTopArtists", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTags gets the top tags on Last.fm.
// See: http://www.last.fm/api/show/tag.getTopTags
func (s *Service) GetTopTags(ctx context.Context, options url.Values) (*GetTopTagsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}

	var resp GetTopTagsResponse
	err := s.client.Call(ctx, "GET", "tag.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTracks gets the top tracks for a tag.
// See: http://www.last.fm/api/show/tag.getTopTracks
func (s *Service) GetTopTracks(ctx context.Context, tag string, options url.Values) (*GetTopTracksResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("tag", tag)

	var resp GetTopTracksResponse
	err := s.client.Call(ctx, "GET", "tag.getTopTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWeeklyChartList gets a list of available charts for this tag.
// See: http://www.last.fm/api/show/tag.getWeeklyChartList
func (s *Service) GetWeeklyChartList(ctx context.Context, tag string, options url.Values) (*GetWeeklyChartListResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("tag", tag)

	var resp GetWeeklyChartListResponse
	err := s.client.Call(ctx, "GET", "tag.getWeeklyChartList", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
