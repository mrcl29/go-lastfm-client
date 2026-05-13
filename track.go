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
func (s *TrackService) GetInfo(ctx context.Context, artist, track string, options url.Values) (*TrackGetInfoResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp TrackGetInfoResponse
	err := s.client.Call(ctx, "GET", "track.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Search searches for a track by name.
// See: http://www.last.fm/api/show/track.search
func (s *TrackService) Search(ctx context.Context, track string, options url.Values) (*TrackSearchResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("track", track)

	var resp TrackSearchResponse
	err := s.client.Call(ctx, "GET", "track.search", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Scrobble adds a track-play to a user's profile.
// See: http://www.last.fm/api/show/track.scrobble
func (s *TrackService) Scrobble(ctx context.Context, artist, track string, timestamp int64, options url.Values) (*TrackScrobbleResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)
	params.Set("timestamp", strconv.FormatInt(timestamp, 10))

	var resp TrackScrobbleResponse
	err := s.client.Call(ctx, "POST", "track.scrobble", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateNowPlaying notifies Last.fm that a user has started listening to a track.
// See: http://www.last.fm/api/show/track.updateNowPlaying
func (s *TrackService) UpdateNowPlaying(ctx context.Context, artist, track string, options url.Values) (*TrackNowPlayingResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp TrackNowPlayingResponse
	err := s.client.Call(ctx, "POST", "track.updateNowPlaying", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddTags tags a track using a list of user supplied tags.
// See: http://www.last.fm/api/show/track.addTags
func (s *TrackService) AddTags(ctx context.Context, artist, track, tags string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("track", track)
	params.Set("tags", tags)

	return s.client.Call(ctx, "POST", "track.addTags", params, nil)
}

// RemoveTag removes a user supplied tag from a track.
// See: http://www.last.fm/api/show/track.removeTag
func (s *TrackService) RemoveTag(ctx context.Context, artist, track, tag string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("track", track)
	params.Set("tag", tag)

	return s.client.Call(ctx, "POST", "track.removeTag", params, nil)
}

// GetTags gets the tags applied by an individual user to a track.
// See: http://www.last.fm/api/show/track.getTags
func (s *TrackService) GetTags(ctx context.Context, artist, track string, options url.Values) (*TrackGetTagsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp TrackGetTagsResponse
	err := s.client.Call(ctx, "GET", "track.getTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTags gets the top tags for a track.
// See: http://www.last.fm/api/show/track.getTopTags
func (s *TrackService) GetTopTags(ctx context.Context, artist, track string, options url.Values) (*TrackGetTopTagsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp TrackGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "track.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCorrection gets the corrected artist/track names.
// See: http://www.last.fm/api/show/track.getCorrection
func (s *TrackService) GetCorrection(ctx context.Context, artist, track string) (*TrackGetCorrectionResponse, error) {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp TrackGetCorrectionResponse
	err := s.client.Call(ctx, "GET", "track.getCorrection", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSimilar gets similar tracks.
// See: http://www.last.fm/api/show/track.getSimilar
func (s *TrackService) GetSimilar(ctx context.Context, artist, track string, options url.Values) (*TrackGetSimilarResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("artist", artist)
	params.Set("track", track)

	var resp TrackGetSimilarResponse
	err := s.client.Call(ctx, "GET", "track.getSimilar", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Love loves a track.
// See: http://www.last.fm/api/show/track.love
func (s *TrackService) Love(ctx context.Context, artist, track string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("track", track)

	return s.client.Call(ctx, "POST", "track.love", params, nil)
}

// Unlove unloves a track.
// See: http://www.last.fm/api/show/track.unlove
func (s *TrackService) Unlove(ctx context.Context, artist, track string) error {
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("track", track)

	return s.client.Call(ctx, "POST", "track.unlove", params, nil)
}

// TrackGetInfoResponse is the response from track.getInfo.
type TrackGetInfoResponse struct {
	Track Track `json:"track"`
}

// TrackSearchResponse is the response from track.search.
type TrackSearchResponse struct {
	Results struct {
		TrackMatches struct {
			Track TrackList `json:"track"`
		} `json:"trackmatches"`
	} `json:"results"`
}

// TrackScrobbleResponse is the response from track.scrobble.
type TrackScrobbleResponse struct {
	Scrobbles struct {
		Scrobble ScrobbleResults `json:"scrobble"`
		Attr     struct {
			Accepted int `json:"accepted"`
			Ignored  int `json:"ignored"`
		} `json:"@attr"`
	} `json:"scrobbles"`
}

// ScrobbleResults is a slice of ScrobbleResult that handles both a single object and an array in JSON.
type ScrobbleResults []ScrobbleResult

// UnmarshalJSON implements json.Unmarshaler.
func (s *ScrobbleResults) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]ScrobbleResult)(s))
}

// ScrobbleResult represents the result of a single scrobble.
type ScrobbleResult struct {
	Artist    TrackCorrectionEntity `json:"artist"`
	Album     TrackCorrectionEntity `json:"album"`
	Track     TrackCorrectionEntity `json:"track"`
	Timestamp string                `json:"timestamp"`
}

// TrackCorrectionEntity represents corrected metadata in a scrobble response.
type TrackCorrectionEntity struct {
	Text      string `json:"#text"`
	Corrected string `json:"corrected"`
}

// TrackNowPlayingResponse is the response from track.updateNowPlaying.
type TrackNowPlayingResponse struct {
	NowPlaying struct {
		Artist    TrackCorrectionEntity `json:"artist"`
		Album     TrackCorrectionEntity `json:"album"`
		Track     TrackCorrectionEntity `json:"track"`
		Timestamp string                `json:"timestamp"`
	} `json:"nowplaying"`
}

// TrackGetTagsResponse is the response from track.getTags.
type TrackGetTagsResponse struct {
	Tags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
		Track  string  `json:"track"`
	} `json:"tags"`
}

// TrackGetTopTagsResponse is the response from track.getTopTags.
type TrackGetTopTagsResponse struct {
	TopTags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
		Track  string  `json:"track"`
	} `json:"toptags"`
}

// TrackGetCorrectionResponse is the response from track.getCorrection.
type TrackGetCorrectionResponse struct {
	Corrections struct {
		Correction TrackCorrection `json:"correction"`
	} `json:"corrections"`
}

// TrackCorrection represents a track correction.
type TrackCorrection struct {
	Track Track `json:"track"`
}

// TrackGetSimilarResponse is the response from track.getSimilar.
type TrackGetSimilarResponse struct {
	SimilarTracks struct {
		Track TrackList `json:"track"`
		Attr  struct {
			Artist string `json:"artist"`
		} `json:"@attr"`
	} `json:"similartracks"`
}
