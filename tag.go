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
func (s *TagService) GetInfo(ctx context.Context, tag string, options url.Values) (*TagGetInfoResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp TagGetInfoResponse
	err := s.client.Call(ctx, "GET", "tag.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSimilar gets all the tags similar to this tag.
// See: http://www.last.fm/api/show/tag.getSimilar
func (s *TagService) GetSimilar(ctx context.Context, tag string, options url.Values) (*TagGetSimilarResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp TagGetSimilarResponse
	err := s.client.Call(ctx, "GET", "tag.getSimilar", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopAlbums gets the top albums for a tag.
// See: http://www.last.fm/api/show/tag.getTopAlbums
func (s *TagService) GetTopAlbums(ctx context.Context, tag string, options url.Values) (*TagGetTopAlbumsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp TagGetTopAlbumsResponse
	err := s.client.Call(ctx, "GET", "tag.getTopAlbums", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopArtists gets the top artists for a tag.
// See: http://www.last.fm/api/show/tag.getTopArtists
func (s *TagService) GetTopArtists(ctx context.Context, tag string, options url.Values) (*TagGetTopArtistsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp TagGetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "tag.getTopArtists", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTags gets the top tags on Last.fm.
// See: http://www.last.fm/api/show/tag.getTopTags
func (s *TagService) GetTopTags(ctx context.Context, options url.Values) (*TagGetTopTagsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}

	var resp TagGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "tag.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTracks gets the top tracks for a tag.
// See: http://www.last.fm/api/show/tag.getTopTracks
func (s *TagService) GetTopTracks(ctx context.Context, tag string, options url.Values) (*TagGetTopTracksResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp TagGetTopTracksResponse
	err := s.client.Call(ctx, "GET", "tag.getTopTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWeeklyChartList gets a list of available charts for this tag.
// See: http://www.last.fm/api/show/tag.getWeeklyChartList
func (s *TagService) GetWeeklyChartList(ctx context.Context, tag string, options url.Values) (*TagGetWeeklyChartListResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("tag", tag)

	var resp TagGetWeeklyChartListResponse
	err := s.client.Call(ctx, "GET", "tag.getWeeklyChartList", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TagGetInfoResponse is the response from tag.getInfo.
type TagGetInfoResponse struct {
	Tag Tag `json:"tag"`
}

// TagGetSimilarResponse is the response from tag.getSimilar.
type TagGetSimilarResponse struct {
	SimilarTags struct {
		Tag  TagList `json:"tag"`
		Attr struct {
			Tag string `json:"tag"`
		} `json:"@attr"`
	} `json:"similartags"`
}

// TagGetTopAlbumsResponse is the response from tag.getTopAlbums.
type TagGetTopAlbumsResponse struct {
	Albums struct {
		Album AlbumList `json:"album"`
		Attr  Attr      `json:"@attr"`
	} `json:"albums"`
}

// TagGetTopArtistsResponse is the response from tag.getTopArtists.
type TagGetTopArtistsResponse struct {
	TopArtists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"topartists"`
}

// TagGetTopTagsResponse is the response from tag.getTopTags.
type TagGetTopTagsResponse struct {
	TopTags struct {
		Tag  TagList `json:"tag"`
		Attr struct {
			NumRes int `json:"num_res"`
			Offset int `json:"offset"`
			Total  int `json:"total"`
		} `json:"@attr"`
	} `json:"toptags"`
}

// TagGetTopTracksResponse is the response from tag.getTopTracks.
type TagGetTopTracksResponse struct {
	Tracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"tracks"`
}

// TagGetWeeklyChartListResponse is the response from tag.getWeeklyChartList.
type TagGetWeeklyChartListResponse struct {
	WeeklyChartList struct {
		Chart ChartList `json:"chart"`
		Attr  struct {
			Tag string `json:"tag"`
		} `json:"@attr"`
	} `json:"weeklychartlist"`
}
