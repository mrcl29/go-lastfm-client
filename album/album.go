package album

import (
	"context"
	"net/url"
)

// APIClient defines the interface required by the album service.
type APIClient interface {
	Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error
}

// Service handles API calls related to albums.
type Service struct {
	client APIClient
}

// New creates a new album service.
func New(client APIClient) *Service {
	return &Service{client: client}
}

// GetInfo gets the metadata and tracklist for an album.
// See: http://www.last.fm/api/show/album.getInfo
func (s *Service) GetInfo(ctx context.Context, artist, album string, options url.Values) (*GetInfoResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	if artist != "" {
		params.Set("artist", artist)
	}
	if album != "" {
		params.Set("album", album)
	}

	var resp GetInfoResponse
	err := s.client.Call(ctx, "GET", "album.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Search searches for an album by name.
// See: http://www.last.fm/api/show/album.search
func (s *Service) Search(ctx context.Context, album string, options url.Values) (*SearchResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("album", album)

	var resp SearchResponse
	err := s.client.Call(ctx, "GET", "album.search", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTags gets the tags applied by an individual user to an album.
// See: http://www.last.fm/api/show/album.getTags
func (s *Service) GetTags(ctx context.Context, artist, album string, options url.Values) (*GetTagsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	if artist != "" {
		params.Set("artist", artist)
	}
	if album != "" {
		params.Set("album", album)
	}

	var resp GetTagsResponse
	err := s.client.Call(ctx, "GET", "album.getTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTags gets the top tags for an album.
// See: http://www.last.fm/api/show/album.getTopTags
func (s *Service) GetTopTags(ctx context.Context, artist, album string, options url.Values) (*GetTopTagsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	if artist != "" {
		params.Set("artist", artist)
	}
	if album != "" {
		params.Set("album", album)
	}

	var resp GetTopTagsResponse
	err := s.client.Call(ctx, "GET", "album.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
