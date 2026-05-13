package track

import (
	"context"
	"net/url"
	"strconv"
)

// APIClient defines the interface required by the track service.
// This allows the track package to remain independent of the main lastfm package.
type APIClient interface {
	Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error
}

// Service handles API calls related to tracks.
type Service struct {
	client APIClient
}

// New creates a new track service.
func New(client APIClient) *Service {
	return &Service{client: client}
}

// GetInfo gets the metadata for a track.
// See: http://www.last.fm/api/show/track.getInfo
func (s *Service) GetInfo(ctx context.Context, artist, track string, options url.Values) (*GetInfoResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp GetInfoResponse
	err := s.client.Call(ctx, "GET", "track.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Search searches for a track by name.
// See: http://www.last.fm/api/show/track.search
func (s *Service) Search(ctx context.Context, track string, options url.Values) (*SearchResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("track", track)

	var resp SearchResponse
	err := s.client.Call(ctx, "GET", "track.search", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Scrobble adds a track-play to a user's profile.
// See: http://www.last.fm/api/show/track.scrobble
func (s *Service) Scrobble(ctx context.Context, artist, track string, timestamp int64, options url.Values) (*ScrobbleResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("artist", artist)
	params.Set("track", track)
	params.Set("timestamp", strconv.FormatInt(timestamp, 10))

	var resp ScrobbleResponse
	err := s.client.Call(ctx, "POST", "track.scrobble", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateNowPlaying notifies Last.fm that a user has started listening to a track.
// See: http://www.last.fm/api/show/track.updateNowPlaying
func (s *Service) UpdateNowPlaying(ctx context.Context, artist, track string, options url.Values) (*NowPlayingResponse, error) {
	params := url.Values{}
	if options != nil {
		for k, v := range options {
			params[k] = v
		}
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp NowPlayingResponse
	err := s.client.Call(ctx, "POST", "track.updateNowPlaying", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
