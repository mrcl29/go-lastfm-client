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
func (s *AlbumService) GetInfo(ctx context.Context, artist, album string, options url.Values) (*AlbumGetInfoResponse, error) {
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

	var resp AlbumGetInfoResponse
	err := s.client.Call(ctx, "GET", "album.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Search searches for an album by name.
// See: http://www.last.fm/api/show/album.search
func (s *AlbumService) Search(ctx context.Context, album string, options url.Values) (*AlbumSearchResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("album", album)

	var resp AlbumSearchResponse
	err := s.client.Call(ctx, "GET", "album.search", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTags gets the tags applied by an individual user to an album.
// See: http://www.last.fm/api/show/album.getTags
func (s *AlbumService) GetTags(ctx context.Context, artist, album string, options url.Values) (*AlbumGetTagsResponse, error) {
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

	var resp AlbumGetTagsResponse
	err := s.client.Call(ctx, "GET", "album.getTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTags gets the top tags for an album.
// See: http://www.last.fm/api/show/album.getTopTags
func (s *AlbumService) GetTopTags(ctx context.Context, artist, album string, options url.Values) (*AlbumGetTopTagsResponse, error) {
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

	var resp AlbumGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "album.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddTags tags an album using a list of user supplied tags.
// See: http://www.last.fm/api/show/album.addTags
func (s *AlbumService) AddTags(ctx context.Context, artist, album, tags string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("album", album)
	params.Set("tags", tags)

	return s.client.Call(ctx, "POST", "album.addTags", params, nil)
}

// RemoveTag removes a user supplied tag from an album.
// See: http://www.last.fm/api/show/album.removeTag
func (s *AlbumService) RemoveTag(ctx context.Context, artist, album, tag string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("album", album)
	params.Set("tag", tag)

	return s.client.Call(ctx, "POST", "album.removeTag", params, nil)
}

// AlbumGetInfoResponse is the response from album.getInfo.
type AlbumGetInfoResponse struct {
	Album Album `json:"album"`
}

// AlbumSearchResponse is the response from album.search.
type AlbumSearchResponse struct {
	Results struct {
		AlbumMatches struct {
			Album AlbumList `json:"album"`
		} `json:"albummatches"`
	} `json:"results"`
}

// AlbumGetTagsResponse is the response from album.getTags.
type AlbumGetTagsResponse struct {
	Tags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
		Album  string  `json:"album"`
	} `json:"tags"`
}

// AlbumGetTopTagsResponse is the response from album.getTopTags.
type AlbumGetTopTagsResponse struct {
	TopTags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
		Album  string  `json:"album"`
	} `json:"toptags"`
}
