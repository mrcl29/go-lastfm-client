package golastfmclient

import (
	"context"
	"net/url"
)

// LibraryService handles API calls related to a user's library.
type LibraryService struct {
	client APIClient
}

// NewLibraryService creates a new library service.
func NewLibraryService(client APIClient) *LibraryService {
	return &LibraryService{client: client}
}

// GetArtists gets a paginated list of all the artists in a user's library.
// See: http://www.last.fm/api/show/library.getArtists
func (s *LibraryService) GetArtists(ctx context.Context, user string, options url.Values) (*LibraryGetArtistsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp LibraryGetArtistsResponse
	err := s.client.Call(ctx, "GET", "library.getArtists", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LibraryGetArtistsResponse is the response from library.getArtists.
type LibraryGetArtistsResponse struct {
	Artists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"artists"`
}
