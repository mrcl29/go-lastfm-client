package golastfmclient

import (
	"context"
	"net/url"
	"strconv"

	"github.com/mrcl29/go-lastfm-client/internal/jsonutil"
)

// TrackService handles API calls related to tracks.
type TrackService struct {
	client APIClient
}

// NewTrackService creates a new track service.
func NewTrackService(client APIClient) *TrackService {
	return &TrackService{client: client}
}

// GetInfo gets the metadata for a track.
// See: http://www.last.fm/api/show/track.getInfo
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//   - options: Additional options (e.g. mbid, username, autocorrect).
//
// Returns:
//   - *Track: The track details.
//   - error: Error if the request fails.
func (s *TrackService) GetInfo(ctx context.Context, artist, track string, options url.Values) (*Track, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp trackGetInfoResponse
	err := s.client.Call(ctx, "GET", "track.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Track, nil
}

// Search searches for a track by name.
// See: http://www.last.fm/api/show/track.search
//
// Parameters:
//   - ctx: Context for the request.
//   - track: The track name to search for.
//   - options: Additional options (e.g. artist, page, limit).
//
// Returns:
//   - TrackList: A slice of tracks matching the search.
//   - error: Error if the request fails.
func (s *TrackService) Search(ctx context.Context, track string, options url.Values) (TrackList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("track", track)

	var resp trackSearchResponse
	err := s.client.Call(ctx, "GET", "track.search", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Results.TrackMatches.Track, nil
}

// Scrobble adds a track-play to a user's profile.
// See: http://www.last.fm/api/show/track.scrobble
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//   - timestamp: The time the track started playing (Unix timestamp).
//   - options: Additional options (e.g. album, trackNumber, context, streamId, albumArtist).
//
// Returns:
//   - ScrobbleResults: Results of the scrobble request.
//   - error: Error if the request fails.
func (s *TrackService) Scrobble(ctx context.Context, artist, track string, timestamp int64, options url.Values) (ScrobbleResults, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)
	params.Set("timestamp", strconv.FormatInt(timestamp, 10))

	var resp trackScrobbleResponse
	err := s.client.Call(ctx, "POST", "track.scrobble", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Scrobbles.Scrobble, nil
}

// UpdateNowPlaying notifies Last.fm that a user has started listening to a track.
// See: http://www.last.fm/api/show/track.updateNowPlaying
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//   - options: Additional options (e.g. album, trackNumber, context, albumArtist).
//
// Returns:
//   - *ScrobbleResult: Result of the now playing update.
//   - error: Error if the request fails.
func (s *TrackService) UpdateNowPlaying(ctx context.Context, artist, track string, options url.Values) (*ScrobbleResult, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp trackNowPlayingResponse
	err := s.client.Call(ctx, "POST", "track.updateNowPlaying", params, &resp)
	if err != nil {
		return nil, err
	}
	return (*ScrobbleResult)(&resp.NowPlaying), nil
}

// AddTags tags a track using a list of user supplied tags.
// See: http://www.last.fm/api/show/track.addTags
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//   - tags: A comma-separated list of tags to apply.
//
// Returns:
//   - error: Error if the request fails.
func (s *TrackService) AddTags(ctx context.Context, artist, track, tags string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("track", track)
	params.Set("tags", tags)

	return s.client.Call(ctx, "POST", "track.addTags", params, nil)
}

// RemoveTag removes a user supplied tag from a track.
// See: http://www.last.fm/api/show/track.removeTag
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//   - tag: A single tag to remove.
//
// Returns:
//   - error: Error if the request fails.
func (s *TrackService) RemoveTag(ctx context.Context, artist, track, tag string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("track", track)
	params.Set("tag", tag)

	return s.client.Call(ctx, "POST", "track.removeTag", params, nil)
}

// GetTags gets the tags applied by an individual user to a track.
// See: http://www.last.fm/api/show/track.getTags
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//   - options: Additional options (e.g. user, mbid, autocorrect).
//
// Returns:
//   - TagList: A slice of tags applied by the user.
//   - error: Error if the request fails.
func (s *TrackService) GetTags(ctx context.Context, artist, track string, options url.Values) (TagList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp trackGetTagsResponse
	err := s.client.Call(ctx, "GET", "track.getTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Tags.Tag, nil
}

// GetTopTags gets the top tags for a track.
// See: http://www.last.fm/api/show/track.getTopTags
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//   - options: Additional options (e.g. mbid, autocorrect).
//
// Returns:
//   - TagList: A slice of top tags for the track.
//   - error: Error if the request fails.
func (s *TrackService) GetTopTags(ctx context.Context, artist, track string, options url.Values) (TagList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp trackGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "track.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.TopTags.Tag, nil
}

// GetCorrection gets the corrected artist/track names.
// See: http://www.last.fm/api/show/track.getCorrection
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//
// Returns:
//   - *Track: The corrected track details.
//   - error: Error if the request fails.
func (s *TrackService) GetCorrection(ctx context.Context, artist, track string) (*Track, error) {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp trackGetCorrectionResponse
	err := s.client.Call(ctx, "GET", "track.getCorrection", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Corrections.Correction.Track, nil
}

// GetSimilar gets similar tracks.
// See: http://www.last.fm/api/show/track.getSimilar
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//   - options: Additional options (e.g. limit, mbid, autocorrect).
//
// Returns:
//   - TrackList: A slice of similar tracks.
//   - error: Error if the request fails.
func (s *TrackService) GetSimilar(ctx context.Context, artist, track string, options url.Values) (TrackList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp trackGetSimilarResponse
	err := s.client.Call(ctx, "GET", "track.getSimilar", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.SimilarTracks.Track, nil
}

// Love loves a track.
// See: http://www.last.fm/api/show/track.love
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//
// Returns:
//   - error: Error if the request fails.
func (s *TrackService) Love(ctx context.Context, artist, track string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("track", track)

	return s.client.Call(ctx, "POST", "track.love", params, nil)
}

// Unlove unloves a track.
// See: http://www.last.fm/api/show/track.unlove
//
// Parameters:
//   - ctx: Context for the request.
//   - artist: The artist name.
//   - track: The track name.
//
// Returns:
//   - error: Error if the request fails.
func (s *TrackService) Unlove(ctx context.Context, artist, track string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("track", track)

	return s.client.Call(ctx, "POST", "track.unlove", params, nil)
}

type trackGetInfoResponse struct {
	Track Track `json:"track"`
}

type trackSearchResponse struct {
	Results struct {
		TrackMatches struct {
			Track TrackList `json:"track"`
		} `json:"trackmatches"`
	} `json:"results"`
}

type trackScrobbleResponse struct {
	Scrobbles struct {
		Scrobble Scrobbles `json:"scrobble"`
	} `json:"scrobbles"`
}

// Scrobbles is a slice of ScrobbleResult that handles both a single object and an array in JSON.
type Scrobbles []ScrobbleResult

// UnmarshalJSON implements json.Unmarshaler.
func (s *Scrobbles) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]ScrobbleResult)(s))
}

// ScrobbleResults is an alias for Scrobbles for backward compatibility or clarity.
type ScrobbleResults = Scrobbles

// ScrobbleResult represents the result of a single scrobble.
type ScrobbleResult struct {
	Artist    trackCorrectionEntity `json:"artist"`
	Album     trackCorrectionEntity `json:"album"`
	Track     trackCorrectionEntity `json:"track"`
	Timestamp string                `json:"timestamp"`
}

type trackCorrectionEntity struct {
	Text      string `json:"#text"`
	Corrected string `json:"corrected"`
}

type trackNowPlayingResponse struct {
	NowPlaying trackNowPlayingResult `json:"nowplaying"`
}

type trackNowPlayingResult struct {
	Artist    trackCorrectionEntity `json:"artist"`
	Album     trackCorrectionEntity `json:"album"`
	Track     trackCorrectionEntity `json:"track"`
	Timestamp string                `json:"timestamp"`
}

type trackGetTagsResponse struct {
	Tags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
		Track  string  `json:"track"`
	} `json:"tags"`
}

type trackGetTopTagsResponse struct {
	TopTags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
		Track  string  `json:"track"`
	} `json:"toptags"`
}

type trackGetCorrectionResponse struct {
	Corrections struct {
		Correction trackCorrection `json:"correction"`
	} `json:"corrections"`
}

type trackCorrection struct {
	Track Track `json:"track"`
}

type trackGetSimilarResponse struct {
	SimilarTracks struct {
		Track TrackList `json:"track"`
		Attr  struct {
			Artist string `json:"artist"`
		} `json:"@attr"`
	} `json:"similartracks"`
}
