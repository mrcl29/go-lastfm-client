package golastfmclient

import (
	"context"
	"net/url"
)

// ArtistService handles API calls related to artists.
type ArtistService struct {
	client APIClient
}

// NewArtistService creates a new artist service.
func NewArtistService(client APIClient) *ArtistService {
	return &ArtistService{client: client}
}

// GetInfo gets the metadata for an artist.
// See: http://www.last.fm/api/show/artist.getInfo
func (s *ArtistService) GetInfo(ctx context.Context, artist string, options url.Values) (*ArtistGetInfoResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp ArtistGetInfoResponse
	err := s.client.Call(ctx, "GET", "artist.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Search searches for an artist by name.
// See: http://www.last.fm/api/show/artist.search
func (s *ArtistService) Search(ctx context.Context, artist string, options url.Values) (*ArtistSearchResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp ArtistSearchResponse
	err := s.client.Call(ctx, "GET", "artist.search", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSimilar gets all the artists similar to this artist.
// See: http://www.last.fm/api/show/artist.getSimilar
func (s *ArtistService) GetSimilar(ctx context.Context, artist string, options url.Values) (*ArtistGetSimilarResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp ArtistGetSimilarResponse
	err := s.client.Call(ctx, "GET", "artist.getSimilar", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCorrection checks whether the supplied artist has a correction to a canonical artist.
// See: http://www.last.fm/api/show/artist.getCorrection
func (s *ArtistService) GetCorrection(ctx context.Context, artist string) (*ArtistGetCorrectionResponse, error) {
	params := url.Values{}
	params.Set("artist", artist)

	var resp ArtistGetCorrectionResponse
	err := s.client.Call(ctx, "GET", "artist.getCorrection", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopAlbums gets the top albums for an artist on Last.fm, ordered by popularity.
// See: http://www.last.fm/api/show/artist.getTopAlbums
func (s *ArtistService) GetTopAlbums(ctx context.Context, artist string, options url.Values) (*ArtistGetTopAlbumsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp ArtistGetTopAlbumsResponse
	err := s.client.Call(ctx, "GET", "artist.getTopAlbums", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTracks gets the top tracks by an artist on Last.fm, ordered by popularity.
// See: http://www.last.fm/api/show/artist.getTopTracks
func (s *ArtistService) GetTopTracks(ctx context.Context, artist string, options url.Values) (*ArtistGetTopTracksResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp ArtistGetTopTracksResponse
	err := s.client.Call(ctx, "GET", "artist.getTopTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddTags tags an artist using a list of user supplied tags.
// See: http://www.last.fm/api/show/artist.addTags
func (s *ArtistService) AddTags(ctx context.Context, artist, tags string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("tags", tags)

	return s.client.Call(ctx, "POST", "artist.addTags", params, nil)
}

// RemoveTag removes a user supplied tag from an artist.
// See: http://www.last.fm/api/show/artist.removeTag
func (s *ArtistService) RemoveTag(ctx context.Context, artist, tag string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("tag", tag)

	return s.client.Call(ctx, "POST", "artist.removeTag", params, nil)
}

// GetTags gets the tags applied by an individual user to an artist.
// See: http://www.last.fm/api/show/artist.getTags
func (s *ArtistService) GetTags(ctx context.Context, artist string, options url.Values) (*ArtistGetTagsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp ArtistGetTagsResponse
	err := s.client.Call(ctx, "GET", "artist.getTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTags gets the top tags for an artist.
// See: http://www.last.fm/api/show/artist.getTopTags
func (s *ArtistService) GetTopTags(ctx context.Context, artist string, options url.Values) (*ArtistGetTopTagsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp ArtistGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "artist.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ArtistGetInfoResponse is the response from artist.getInfo.
type ArtistGetInfoResponse struct {
	Artist Artist `json:"artist"`
}

// ArtistSearchResponse is the response from artist.search.
type ArtistSearchResponse struct {
	Results struct {
		ArtistMatches struct {
			Artist ArtistList `json:"artist"`
		} `json:"artistmatches"`
	} `json:"results"`
}

// ArtistGetSimilarResponse is the response from artist.getSimilar.
type ArtistGetSimilarResponse struct {
	SimilarArtists struct {
		Artist ArtistList `json:"artist"`
		Attr   struct {
			Artist string `json:"artist"`
		} `json:"@attr"`
	} `json:"similarartists"`
}

// ArtistGetCorrectionResponse is the response from artist.getCorrection.
type ArtistGetCorrectionResponse struct {
	Corrections struct {
		Correction ArtistCorrection `json:"correction"`
	} `json:"corrections"`
}

// ArtistCorrection represents an artist correction.
type ArtistCorrection struct {
	Artist Artist `json:"artist"`
	Index  string `json:"index"`
}

// ArtistGetTopAlbumsResponse is the response from artist.getTopAlbums.
type ArtistGetTopAlbumsResponse struct {
	TopAlbums struct {
		Album AlbumList `json:"album"`
		Attr  struct {
			Artist string `json:"artist"`
			Page   string `json:"page"`
			Total  string `json:"total"`
		} `json:"@attr"`
	} `json:"topalbums"`
}

// ArtistGetTopTracksResponse is the response from artist.getTopTracks.
type ArtistGetTopTracksResponse struct {
	TopTracks struct {
		Track TrackList `json:"track"`
		Attr  struct {
			Artist string `json:"artist"`
			Page   string `json:"page"`
			Total  string `json:"total"`
		} `json:"@attr"`
	} `json:"toptracks"`
}

// ArtistGetTagsResponse is the response from artist.getTags.
type ArtistGetTagsResponse struct {
	Tags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
	} `json:"tags"`
}

// ArtistGetTopTagsResponse is the response from artist.getTopTags.
type ArtistGetTopTagsResponse struct {
	TopTags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
	} `json:"toptags"`
}
