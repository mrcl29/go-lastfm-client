package artist

import (
	"context"
	"net/url"
)

// APIClient defines the interface required by the artist service.
type APIClient interface {
	Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error
}

// Service handles API calls related to artists.
type Service struct {
	client APIClient
}

// New creates a new artist service.
func New(client APIClient) *Service {
	return &Service{client: client}
}

// GetInfo gets the metadata for an artist.
// See: http://www.last.fm/api/show/artist.getInfo
func (s *Service) GetInfo(ctx context.Context, artist string, options url.Values) (*GetInfoResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("artist", artist)

	var resp GetInfoResponse
	err := s.client.Call(ctx, "GET", "artist.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Search searches for an artist by name.
// See: http://www.last.fm/api/show/artist.search
func (s *Service) Search(ctx context.Context, artist string, options url.Values) (*SearchResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("artist", artist)

	var resp SearchResponse
	err := s.client.Call(ctx, "GET", "artist.search", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSimilar gets all the artists similar to this artist.
// See: http://www.last.fm/api/show/artist.getSimilar
func (s *Service) GetSimilar(ctx context.Context, artist string, options url.Values) (*GetSimilarResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("artist", artist)

	var resp GetSimilarResponse
	err := s.client.Call(ctx, "GET", "artist.getSimilar", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCorrection checks whether the supplied artist has a correction to a canonical artist.
// See: http://www.last.fm/api/show/artist.getCorrection
func (s *Service) GetCorrection(ctx context.Context, artist string) (*GetCorrectionResponse, error) {
	params := url.Values{}
	params.Set("artist", artist)

	var resp GetCorrectionResponse
	err := s.client.Call(ctx, "GET", "artist.getCorrection", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopAlbums gets the top albums for an artist on Last.fm, ordered by popularity.
// See: http://www.last.fm/api/show/artist.getTopAlbums
func (s *Service) GetTopAlbums(ctx context.Context, artist string, options url.Values) (*GetTopAlbumsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("artist", artist)

	var resp GetTopAlbumsResponse
	err := s.client.Call(ctx, "GET", "artist.getTopAlbums", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTracks gets the top tracks by an artist on Last.fm, ordered by popularity.
// See: http://www.last.fm/api/show/artist.getTopTracks
func (s *Service) GetTopTracks(ctx context.Context, artist string, options url.Values) (*GetTopTracksResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("artist", artist)

	var resp GetTopTracksResponse
	err := s.client.Call(ctx, "GET", "artist.getTopTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
