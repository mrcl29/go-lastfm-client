package golastfmclient

import (
	"context"
	"net/url"
)

// AlbumService handles API calls related to albums.
type AlbumService struct {
	client APIClient
}

// NewAlbumService creates a new album service.
func NewAlbumService(client APIClient) *AlbumService {
	return &AlbumService{client: client}
}

// GetInfo gets the metadata and tracklist for an album.
// See: http://www.last.fm/api/show/album.getInfo
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - album: The album name.
//   - options: Additional options (e.g. lang, mbid, username).
//
// Returns:
//   - *Album: The album details.
//   - error: Error if the request fails.
func (s *AlbumService) GetInfo(ctx context.Context, artist, album string, options url.Values) (*Album, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	if artist != "" {
		params.Set("artist", artist)
	}
	if album != "" {
		params.Set("album", album)
	}

	var resp albumGetInfoResponse
	err := s.client.Call(ctx, "GET", "album.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Album, nil
}

// Search searches for an album by name.
// See: http://www.last.fm/api/show/album.search
//
// Parameters:
//   - ctx: Context for the request.
//   - album: The album name to search for.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - AlbumList: A slice of albums matching the search.
//   - error: Error if the request fails.
func (s *AlbumService) Search(ctx context.Context, album string, options url.Values) (AlbumList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("album", album)

	var resp albumSearchResponse
	err := s.client.Call(ctx, "GET", "album.search", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Results.AlbumMatches.Album, nil
}

// GetTags gets the tags applied by an individual user to an album.
// See: http://www.last.fm/api/show/album.getTags
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - album: The album name.
//   - options: Additional options (e.g. user, mbid).
//
// Returns:
//   - TagList: A slice of tags applied by the user.
//   - error: Error if the request fails.
func (s *AlbumService) GetTags(ctx context.Context, artist, album string, options url.Values) (TagList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	if artist != "" {
		params.Set("artist", artist)
	}
	if album != "" {
		params.Set("album", album)
	}

	var resp albumGetTagsResponse
	err := s.client.Call(ctx, "GET", "album.getTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Tags.Tag, nil
}

// GetTopTags gets the top tags for an album.
// See: http://www.last.fm/api/show/album.getTopTags
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - album: The album name.
//   - options: Additional options (e.g. mbid).
//
// Returns:
//   - TagList: A slice of top tags for the album.
//   - error: Error if the request fails.
func (s *AlbumService) GetTopTags(ctx context.Context, artist, album string, options url.Values) (TagList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	if artist != "" {
		params.Set("artist", artist)
	}
	if album != "" {
		params.Set("album", album)
	}

	var resp albumGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "album.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.TopTags.Tag, nil
}

// AddTags tags an album using a list of user supplied tags.
// See: http://www.last.fm/api/show/album.addTags
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - album: The album name.
//   - tags: A comma-separated list of tags to apply.
//
// Returns:
//   - error: Error if the request fails.
func (s *AlbumService) AddTags(ctx context.Context, artist, album, tags string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("album", album)
	params.Set("tags", tags)

	return s.client.Call(ctx, "POST", "album.addTags", params, nil)
}

// RemoveTag removes a user supplied tag from an album.
// See: http://www.last.fm/api/show/album.removeTag
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - album: The album name.
//   - tag: A single tag to remove.
//
// Returns:
//   - error: Error if the request fails.
func (s *AlbumService) RemoveTag(ctx context.Context, artist, album, tag string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("album", album)
	params.Set("tag", tag)

	return s.client.Call(ctx, "POST", "album.removeTag", params, nil)
}

type albumGetInfoResponse struct {
	Album Album `json:"album"`
}

type albumSearchResponse struct {
	Results struct {
		AlbumMatches struct {
			Album AlbumList `json:"album"`
		} `json:"albummatches"`
	} `json:"results"`
}

type albumGetTagsResponse struct {
	Tags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
		Album  string  `json:"album"`
	} `json:"tags"`
}

type albumGetTopTagsResponse struct {
	TopTags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
		Album  string  `json:"album"`
	} `json:"toptags"`
}

