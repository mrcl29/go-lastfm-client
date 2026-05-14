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
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - options: Additional options (e.g. lang, mbid, username, autocorrect).
//
// Returns:
//   - *Artist: The artist details.
//   - error: Error if the request fails.
func (s *ArtistService) GetInfo(ctx context.Context, artist string, options url.Values) (*Artist, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp artistGetInfoResponse
	err := s.client.Call(ctx, "GET", "artist.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Artist, nil
}

// Search searches for an artist by name.
// See: http://www.last.fm/api/show/artist.search
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name to search for.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - ArtistList: A slice of artists matching the search.
//   - error: Error if the request fails.
func (s *ArtistService) Search(ctx context.Context, artist string, options url.Values) (ArtistList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp artistSearchResponse
	err := s.client.Call(ctx, "GET", "artist.search", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Results.ArtistMatches.Artist, nil
}

// GetSimilar gets all the artists similar to this artist.
// See: http://www.last.fm/api/show/artist.getSimilar
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - options: Additional options (e.g. limit, mbid, autocorrect).
//
// Returns:
//   - ArtistList: A slice of similar artists.
//   - error: Error if the request fails.
func (s *ArtistService) GetSimilar(ctx context.Context, artist string, options url.Values) (ArtistList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp artistGetSimilarResponse
	err := s.client.Call(ctx, "GET", "artist.getSimilar", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.SimilarArtists.Artist, nil
}

// GetCorrection checks whether the supplied artist has a correction to a canonical artist.
// See: http://www.last.fm/api/show/artist.getCorrection
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name to correct.
//
// Returns:
//   - *Artist: The corrected artist details.
//   - error: Error if the request fails.
func (s *ArtistService) GetCorrection(ctx context.Context, artist string) (*Artist, error) {
	params := url.Values{}
	params.Set("artist", artist)

	var resp artistGetCorrectionResponse
	err := s.client.Call(ctx, "GET", "artist.getCorrection", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Corrections.Correction.Artist, nil
}

// GetTopAlbums gets the top albums for an artist on Last.fm, ordered by popularity.
// See: http://www.last.fm/api/show/artist.getTopAlbums
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - options: Additional options (e.g. page, limit, mbid, autocorrect).
//
// Returns:
//   - AlbumList: A slice of top albums for the artist.
//   - *Attr: Pagination and metadata.
//   - error: Error if the request fails.
func (s *ArtistService) GetTopAlbums(ctx context.Context, artist string, options url.Values) (AlbumList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp artistGetTopAlbumsResponse
	err := s.client.Call(ctx, "GET", "artist.getTopAlbums", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.TopAlbums.Album, &resp.TopAlbums.Attr, nil
}

// GetTopTracks gets the top tracks by an artist on Last.fm, ordered by popularity.
// See: http://www.last.fm/api/show/artist.getTopTracks
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - options: Additional options (e.g. page, limit, mbid, autocorrect).
//
// Returns:
//   - TrackList: A slice of top tracks for the artist.
//   - *Attr: Pagination and metadata.
//   - error: Error if the request fails.
func (s *ArtistService) GetTopTracks(ctx context.Context, artist string, options url.Values) (TrackList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp artistGetTopTracksResponse
	err := s.client.Call(ctx, "GET", "artist.getTopTracks", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.TopTracks.Track, &resp.TopTracks.Attr, nil
}

// AddTags tags an artist using a list of user supplied tags.
// See: http://www.last.fm/api/show/artist.addTags
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - tags: A comma-separated list of tags to apply.
//
// Returns:
//   - error: Error if the request fails.
func (s *ArtistService) AddTags(ctx context.Context, artist, tags string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("tags", tags)

	return s.client.Call(ctx, "POST", "artist.addTags", params, nil)
}

// RemoveTag removes a user supplied tag from an artist.
// See: http://www.last.fm/api/show/artist.removeTag
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - tag: A single tag to remove.
//
// Returns:
//   - error: Error if the request fails.
func (s *ArtistService) RemoveTag(ctx context.Context, artist, tag string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("tag", tag)

	return s.client.Call(ctx, "POST", "artist.removeTag", params, nil)
}

// GetTags gets the tags applied by an individual user to an artist.
// See: http://www.last.fm/api/show/artist.getTags
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - options: Additional options (e.g. user, mbid, autocorrect).
//
// Returns:
//   - TagList: A slice of tags applied by the user.
//   - error: Error if the request fails.
func (s *ArtistService) GetTags(ctx context.Context, artist string, options url.Values) (TagList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp artistGetTagsResponse
	err := s.client.Call(ctx, "GET", "artist.getTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Tags.Tag, nil
}

// GetTopTags gets the top tags for an artist.
// See: http://www.last.fm/api/show/artist.getTopTags
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - options: Additional options (e.g. mbid, autocorrect).
//
// Returns:
//   - TagList: A slice of top tags for the artist.
//   - error: Error if the request fails.
func (s *ArtistService) GetTopTags(ctx context.Context, artist string, options url.Values) (TagList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)

	var resp artistGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "artist.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.TopTags.Tag, nil
}

type artistGetInfoResponse struct {
	Artist Artist `json:"artist"`
}

type artistSearchResponse struct {
	Results struct {
		ArtistMatches struct {
			Artist ArtistList `json:"artist"`
		} `json:"artistmatches"`
	} `json:"results"`
}

type artistGetSimilarResponse struct {
	SimilarArtists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"similarartists"`
}

type artistGetCorrectionResponse struct {
	Corrections struct {
		Correction artistCorrection `json:"correction"`
	} `json:"corrections"`
}

type artistCorrection struct {
	Artist Artist `json:"artist"`
	Index  string `json:"index"`
}

type artistGetTopAlbumsResponse struct {
	TopAlbums struct {
		Album AlbumList `json:"album"`
		Attr  Attr      `json:"@attr"`
	} `json:"topalbums"`
}

type artistGetTopTracksResponse struct {
	TopTracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"toptracks"`
}

type artistGetTagsResponse struct {
	Tags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
	} `json:"tags"`
}

type artistGetTopTagsResponse struct {
	TopTags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
	} `json:"toptags"`
}
