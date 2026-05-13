package user

import (
	"context"
	"net/url"
)

// APIClient defines the interface required by the user service.
type APIClient interface {
	Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error
}

// Service handles API calls related to users.
type Service struct {
	client APIClient
}

// New creates a new user service.
func New(client APIClient) *Service {
	return &Service{client: client}
}

// GetInfo gets the metadata for a user.
// See: http://www.last.fm/api/show/user.getInfo
func (s *Service) GetInfo(ctx context.Context, user string, options url.Values) (*GetInfoResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	if user != "" {
		params.Set("user", user)
	}

	var resp GetInfoResponse
	err := s.client.Call(ctx, "GET", "user.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRecentTracks gets a list of the recent tracks listened to by this user.
// See: http://www.last.fm/api/show/user.getRecentTracks
func (s *Service) GetRecentTracks(ctx context.Context, user string, options url.Values) (*GetRecentTracksResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("user", user)

	var resp GetRecentTracksResponse
	err := s.client.Call(ctx, "GET", "user.getRecentTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopAlbums gets the top albums listened to by a user.
// See: http://www.last.fm/api/show/user.getTopAlbums
func (s *Service) GetTopAlbums(ctx context.Context, user string, options url.Values) (*GetTopAlbumsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("user", user)

	var resp GetTopAlbumsResponse
	err := s.client.Call(ctx, "GET", "user.getTopAlbums", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopArtists gets the top artists listened to by a user.
// See: http://www.last.fm/api/show/user.getTopArtists
func (s *Service) GetTopArtists(ctx context.Context, user string, options url.Values) (*GetTopArtistsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("user", user)

	var resp GetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "user.getTopArtists", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTracks gets the top tracks listened to by a user.
// See: http://www.last.fm/api/show/user.getTopTracks
func (s *Service) GetTopTracks(ctx context.Context, user string, options url.Values) (*GetTopTracksResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("user", user)

	var resp GetTopTracksResponse
	err := s.client.Call(ctx, "GET", "user.getTopTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLovedTracks gets the loved tracks for a user.
// See: http://www.last.fm/api/show/user.getLovedTracks
func (s *Service) GetLovedTracks(ctx context.Context, user string, options url.Values) (*GetLovedTracksResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("user", user)

	var resp GetLovedTracksResponse
	err := s.client.Call(ctx, "GET", "user.getLovedTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFriends gets the friends for a user.
// See: http://www.last.fm/api/show/user.getFriends
func (s *Service) GetFriends(ctx context.Context, user string, options url.Values) (*GetFriendsResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("user", user)

	var resp GetFriendsResponse
	err := s.client.Call(ctx, "GET", "user.getFriends", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
