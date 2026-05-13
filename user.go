package golastfmclient

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/mrcl29/go-lastfm-client/internal/jsonutil"
)

// UserService handles API calls related to users.
type UserService struct {
	client APIClient
}

// NewUserService creates a new user service.
func NewUserService(client APIClient) *UserService {
	return &UserService{client: client}
}

// GetInfo gets the metadata for a user.
// See: http://www.last.fm/api/show/user.getInfo
func (s *UserService) GetInfo(ctx context.Context, user string, options url.Values) (*UserGetInfoResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	if user != "" {
		params.Set("user", user)
	}

	var resp UserGetInfoResponse
	err := s.client.Call(ctx, "GET", "user.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRecentTracks gets a list of the recent tracks listened to by this user.
// See: http://www.last.fm/api/show/user.getRecentTracks
func (s *UserService) GetRecentTracks(ctx context.Context, user string, options url.Values) (*UserGetRecentTracksResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp UserGetRecentTracksResponse
	err := s.client.Call(ctx, "GET", "user.getRecentTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopAlbums gets the top albums listened to by a user.
// See: http://www.last.fm/api/show/user.getTopAlbums
func (s *UserService) GetTopAlbums(ctx context.Context, user string, options url.Values) (*UserGetTopAlbumsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp UserGetTopAlbumsResponse
	err := s.client.Call(ctx, "GET", "user.getTopAlbums", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopArtists gets the top artists listened to by a user.
// See: http://www.last.fm/api/show/user.getTopArtists
func (s *UserService) GetTopArtists(ctx context.Context, user string, options url.Values) (*UserGetTopArtistsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp UserGetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "user.getTopArtists", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTracks gets the top tracks listened to by a user.
// See: http://www.last.fm/api/show/user.getTopTracks
func (s *UserService) GetTopTracks(ctx context.Context, user string, options url.Values) (*UserGetTopTracksResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp UserGetTopTracksResponse
	err := s.client.Call(ctx, "GET", "user.getTopTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLovedTracks gets the loved tracks for a user.
// See: http://www.last.fm/api/show/user.getLovedTracks
func (s *UserService) GetLovedTracks(ctx context.Context, user string, options url.Values) (*UserGetLovedTracksResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp UserGetLovedTracksResponse
	err := s.client.Call(ctx, "GET", "user.getLovedTracks", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFriends gets the friends for a user.
// See: http://www.last.fm/api/show/user.getFriends
func (s *UserService) GetFriends(ctx context.Context, user string, options url.Values) (*UserGetFriendsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp UserGetFriendsResponse
	err := s.client.Call(ctx, "GET", "user.getFriends", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPersonalTags gets the user's personal tags.
// See: http://www.last.fm/api/show/user.getPersonalTags
func (s *UserService) GetPersonalTags(ctx context.Context, user, tag, taggingType string, options url.Values) (*UserGetPersonalTagsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)
	params.Set("tag", tag)
	params.Set("taggingtype", taggingType)

	var resp UserGetPersonalTagsResponse
	err := s.client.Call(ctx, "GET", "user.getPersonalTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopTags gets the top tags for a user.
// See: http://www.last.fm/api/show/user.getTopTags
func (s *UserService) GetTopTags(ctx context.Context, user string, options url.Values) (*UserGetTopTagsResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp UserGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "user.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWeeklyAlbumChart gets the weekly album chart for a user.
// See: http://www.last.fm/api/show/user.getWeeklyAlbumChart
func (s *UserService) GetWeeklyAlbumChart(ctx context.Context, user string, options url.Values) (*UserGetWeeklyAlbumChartResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp UserGetWeeklyAlbumChartResponse
	err := s.client.Call(ctx, "GET", "user.getWeeklyAlbumChart", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWeeklyArtistChart gets the weekly artist chart for a user.
// See: http://www.last.fm/api/show/user.getWeeklyArtistChart
func (s *UserService) GetWeeklyArtistChart(ctx context.Context, user string, options url.Values) (*UserGetWeeklyArtistChartResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp UserGetWeeklyArtistChartResponse
	err := s.client.Call(ctx, "GET", "user.getWeeklyArtistChart", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWeeklyChartList gets a list of available weekly charts for a user.
// See: http://www.last.fm/api/show/user.getWeeklyChartList
func (s *UserService) GetWeeklyChartList(ctx context.Context, user string) (*UserGetWeeklyChartListResponse, error) {
	params := url.Values{}
	params.Set("user", user)

	var resp UserGetWeeklyChartListResponse
	err := s.client.Call(ctx, "GET", "user.getWeeklyChartList", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWeeklyTrackChart gets the weekly track chart for a user.
// See: http://www.last.fm/api/show/user.getWeeklyTrackChart
func (s *UserService) GetWeeklyTrackChart(ctx context.Context, user string, options url.Values) (*UserGetWeeklyTrackChartResponse, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp UserGetWeeklyTrackChartResponse
	err := s.client.Call(ctx, "GET", "user.getWeeklyTrackChart", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UserGetInfoResponse is the response from user.getInfo.
type UserGetInfoResponse struct {
	User User `json:"user"`
}

// UserGetRecentTracksResponse is the response from user.getRecentTracks.
type UserGetRecentTracksResponse struct {
	RecentTracks struct {
		Track TrackList `json:"track"`
		Attr  struct {
			User       string      `json:"user"`
			Page       json.Number `json:"page"`
			PerPage    json.Number `json:"perPage"`
			TotalPages json.Number `json:"totalPages"`
			Total      json.Number `json:"total"`
		} `json:"@attr"`
	} `json:"recenttracks"`
}

// UserGetTopAlbumsResponse is the response from user.getTopAlbums.
type UserGetTopAlbumsResponse struct {
	TopAlbums struct {
		Album AlbumList `json:"album"`
		Attr  struct {
			User       string      `json:"user"`
			Page       json.Number `json:"page"`
			PerPage    json.Number `json:"perPage"`
			TotalPages json.Number `json:"totalPages"`
			Total      json.Number `json:"total"`
		} `json:"@attr"`
	} `json:"topalbums"`
}

// UserGetTopArtistsResponse is the response from user.getTopArtists.
type UserGetTopArtistsResponse struct {
	TopArtists struct {
		Artist ArtistList `json:"artist"`
		Attr   struct {
			User       string      `json:"user"`
			Page       json.Number `json:"page"`
			PerPage    json.Number `json:"perPage"`
			TotalPages json.Number `json:"totalPages"`
			Total      json.Number `json:"total"`
		} `json:"@attr"`
	} `json:"topartists"`
}

// UserGetTopTracksResponse is the response from user.getTopTracks.
type UserGetTopTracksResponse struct {
	TopTracks struct {
		Track TrackList `json:"track"`
		Attr  struct {
			User       string      `json:"user"`
			Page       json.Number `json:"page"`
			PerPage    json.Number `json:"perPage"`
			TotalPages json.Number `json:"totalPages"`
			Total      json.Number `json:"total"`
		} `json:"@attr"`
	} `json:"toptracks"`
}

// UserGetLovedTracksResponse is the response from user.getLovedTracks.
type UserGetLovedTracksResponse struct {
	LovedTracks struct {
		Track TrackList `json:"track"`
		Attr  struct {
			User       string      `json:"user"`
			Page       json.Number `json:"page"`
			PerPage    json.Number `json:"perPage"`
			TotalPages json.Number `json:"totalPages"`
			Total      json.Number `json:"total"`
		} `json:"@attr"`
	} `json:"lovedtracks"`
}

// UserGetFriendsResponse is the response from user.getFriends.
type UserGetFriendsResponse struct {
	Friends struct {
		User FriendList `json:"user"`
		Attr struct {
			User       string      `json:"user"`
			Page       json.Number `json:"page"`
			PerPage    json.Number `json:"perPage"`
			TotalPages json.Number `json:"totalPages"`
			Total      json.Number `json:"total"`
		} `json:"@attr"`
	} `json:"friends"`
}

// Friend represents a Last.fm user's friend.
type Friend struct {
	Name       string  `json:"name"`
	RealName   string  `json:"realname"`
	Image      []Image `json:"image"`
	URL        string  `json:"url"`
	Country    string  `json:"country"`
	Registered Date    `json:"registered"`
}

// FriendList is a slice of Friend that handles both a single object and an array in JSON.
type FriendList []Friend

// UnmarshalJSON implements json.Unmarshaler.
func (f *FriendList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Friend)(f))
}

// UserGetPersonalTagsResponse is the response from user.getPersonalTags.
type UserGetPersonalTagsResponse struct {
	Taggings struct {
		Artists *struct {
			Artist ArtistList `json:"artist"`
		} `json:"artists,omitempty"`
		Albums *struct {
			Album AlbumList `json:"album"`
		} `json:"albums,omitempty"`
		Tracks *struct {
			Track TrackList `json:"track"`
		} `json:"tracks,omitempty"`
		Attr struct {
			User       string      `json:"user"`
			Tag        string      `json:"tag"`
			Page       json.Number `json:"page"`
			PerPage    json.Number `json:"perPage"`
			TotalPages json.Number `json:"totalPages"`
			Total      json.Number `json:"total"`
		} `json:"@attr"`
	} `json:"taggings"`
}

// UserGetTopTagsResponse is the response from user.getTopTags.
type UserGetTopTagsResponse struct {
	TopTags struct {
		Tag  TagList `json:"tag"`
		Attr struct {
			User string `json:"user"`
		} `json:"@attr"`
	} `json:"toptags"`
}

// UserGetWeeklyAlbumChartResponse is the response from user.getWeeklyAlbumChart.
type UserGetWeeklyAlbumChartResponse struct {
	WeeklyAlbumChart struct {
		Album AlbumList `json:"album"`
		Attr  struct {
			User string `json:"user"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"@attr"`
	} `json:"weeklyalbumchart"`
}

// UserGetWeeklyArtistChartResponse is the response from user.getWeeklyArtistChart.
type UserGetWeeklyArtistChartResponse struct {
	WeeklyArtistChart struct {
		Artist ArtistList `json:"artist"`
		Attr   struct {
			User string `json:"user"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"@attr"`
	} `json:"weeklyartistchart"`
}

// UserGetWeeklyChartListResponse is the response from user.getWeeklyChartList.
type UserGetWeeklyChartListResponse struct {
	WeeklyChartList struct {
		Chart ChartList `json:"chart"`
		Attr  struct {
			User string `json:"user"`
		} `json:"@attr"`
	} `json:"weeklychartlist"`
}

// UserGetWeeklyTrackChartResponse is the response from user.getWeeklyTrackChart.
type UserGetWeeklyTrackChartResponse struct {
	WeeklyTrackChart struct {
		Track TrackList `json:"track"`
		Attr  struct {
			User string `json:"user"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"@attr"`
	} `json:"weeklytrackchart"`
}
