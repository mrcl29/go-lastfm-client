package library

import (
	"context"
	"net/url"
)

// APIClient defines the interface required by the library service.
type APIClient interface {
	Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error
}

// Service handles API calls related to a user's library.
type Service struct {
	client APIClient
}

// New creates a new library service.
func New(client APIClient) *Service {
	return &Service{client: client}
}

// GetArtists gets a paginated list of all the artists in a user's library.
// See: http://www.last.fm/api/show/library.getArtists
func (s *Service) GetArtists(ctx context.Context, user string, options url.Values) (*GetArtistsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("user", user)

	var resp GetArtistsResponse
	err := s.client.Call(ctx, "GET", "library.getArtists", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
