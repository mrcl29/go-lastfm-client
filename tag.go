package golastfmclient

import (
	"context"
	"net/url"
)

// TagService handles API calls related to tags.
type TagService struct {
	client APIClient
}

// NewTagService creates a new tag service.
func NewTagService(client APIClient) *TagService {
	return &TagService{client: client}
}

// GetInfo gets the metadata for a tag.
// See: http://www.last.fm/api/show/tag.getInfo
//
// Parameters:
//   - ctx: Context for the request.
//   - tag: The tag name.
//   - options: Additional options (e.g. lang).
//
// Returns:
//   - *Tag: The tag details.
//   - error: Error if the request fails.
func (s *TagService) GetInfo(ctx context.Context, tag string, options url.Values) (*Tag, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp tagGetInfoResponse
	err := s.client.Call(ctx, "GET", "tag.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Tag, nil
}

// GetSimilar gets all the tags similar to this tag.
// See: http://www.last.fm/api/show/tag.getSimilar
//
// Parameters:
//   - ctx: Context for the request.
//   - tag: The tag name.
//   - options: Additional options.
//
// Returns:
//   - TagList: A slice of similar tags.
//   - error: Error if the request fails.
func (s *TagService) GetSimilar(ctx context.Context, tag string, options url.Values) (TagList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp tagGetSimilarResponse
	err := s.client.Call(ctx, "GET", "tag.getSimilar", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.SimilarTags.Tag, nil
}

// GetTopAlbums gets the top albums for a tag.
// See: http://www.last.fm/api/show/tag.getTopAlbums
//
// Parameters:
//   - ctx: Context for the request.
//   - tag: The tag name.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - AlbumList: A slice of top albums for the tag.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *TagService) GetTopAlbums(ctx context.Context, tag string, options url.Values) (AlbumList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp tagGetTopAlbumsResponse
	err := s.client.Call(ctx, "GET", "tag.getTopAlbums", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Albums.Album, &resp.Albums.Attr, nil
}

// GetTopArtists gets the top artists for a tag.
// See: http://www.last.fm/api/show/tag.getTopArtists
//
// Parameters:
//   - ctx: Context for the request.
//   - tag: The tag name.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - ArtistList: A slice of top artists for the tag.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *TagService) GetTopArtists(ctx context.Context, tag string, options url.Values) (ArtistList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp tagGetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "tag.getTopArtists", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.TopArtists.Artist, &resp.TopArtists.Attr, nil
}

// GetTopTags gets the top tags on Last.fm.
// See: http://www.last.fm/api/show/tag.getTopTags
//
// Parameters:
//   - ctx: Context for the request.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - TagList: A slice of top tags.
//   - error: Error if the request fails.
func (s *TagService) GetTopTags(ctx context.Context, options url.Values) (TagList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}

	var resp tagGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "tag.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.TopTags.Tag, nil
}

// GetTopTracks gets the top tracks for a tag.
// See: http://www.last.fm/api/show/tag.getTopTracks
//
// Parameters:
//   - ctx: Context for the request.
//   - tag: The tag name.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - TrackList: A slice of top tracks for the tag.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *TagService) GetTopTracks(ctx context.Context, tag string, options url.Values) (TrackList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp tagGetTopTracksResponse
	err := s.client.Call(ctx, "GET", "tag.getTopTracks", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Tracks.Track, &resp.Tracks.Attr, nil
}

// GetWeeklyChartList gets a list of available charts for this tag.
// See: http://www.last.fm/api/show/tag.getWeeklyChartList
//
// Parameters:
//   - ctx: Context for the request.
//   - tag: The tag name.
//   - options: Additional options.
//
// Returns:
//   - ChartList: A slice of available charts.
//   - error: Error if the request fails.
func (s *TagService) GetWeeklyChartList(ctx context.Context, tag string, options url.Values) (ChartList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp tagGetWeeklyChartListResponse
	err := s.client.Call(ctx, "GET", "tag.getWeeklyChartList", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.WeeklyChartList.Chart, nil
}

type tagGetInfoResponse struct {
	Tag Tag `json:"tag"`
}

type tagGetSimilarResponse struct {
	SimilarTags struct {
		Tag TagList `json:"tag"`
	} `json:"similartags"`
}

type tagGetTopAlbumsResponse struct {
	Albums struct {
		Album AlbumList `json:"album"`
		Attr  Attr      `json:"@attr"`
	} `json:"albums"`
}

type tagGetTopArtistsResponse struct {
	TopArtists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"topartists"`
}

type tagGetTopTagsResponse struct {
	TopTags struct {
		Tag TagList `json:"tag"`
	} `json:"toptags"`
}

type tagGetTopTracksResponse struct {
	Tracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"tracks"`
}

type tagGetWeeklyChartListResponse struct {
	WeeklyChartList struct {
		Chart ChartList `json:"chart"`
	} `json:"weeklychartlist"`
}
