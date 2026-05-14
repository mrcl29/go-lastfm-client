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
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - ArtistList: A slice of artists in the user's library.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *LibraryService) GetArtists(ctx context.Context, user string, options url.Values) (ArtistList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp libraryGetArtistsResponse
	err := s.client.Call(ctx, "GET", "library.getArtists", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Artists.Artist, &resp.Artists.Attr, nil
}

type libraryGetArtistsResponse struct {
	Artists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"artists"`
}
