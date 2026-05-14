package golastfmclient

import (
	"context"
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
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options.
//
// Returns:
//   - *User: The user details.
//   - error: Error if the request fails.
func (s *UserService) GetInfo(ctx context.Context, user string, options url.Values) (*User, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	if user != "" {
		params.Set("user", user)
	}

	var resp userGetInfoResponse
	err := s.client.Call(ctx, "GET", "user.getInfo", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.User, nil
}

// GetRecentTracks gets a list of the recent tracks listened to by this user.
// See: http://www.last.fm/api/show/user.getRecentTracks
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options (e.g. page, limit, from, to).
//
// Returns:
//   - TrackList: A slice of recent tracks.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *UserService) GetRecentTracks(ctx context.Context, user string, options url.Values) (TrackList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp userGetRecentTracksResponse
	err := s.client.Call(ctx, "GET", "user.getRecentTracks", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.RecentTracks.Track, (*Attr)(&resp.RecentTracks.Attr), nil
}

// GetTopAlbums gets the top albums listened to by a user.
// See: http://www.last.fm/api/show/user.getTopAlbums
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options (e.g. period, page, limit).
//
// Returns:
//   - AlbumList: A slice of top albums.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *UserService) GetTopAlbums(ctx context.Context, user string, options url.Values) (AlbumList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp userGetTopAlbumsResponse
	err := s.client.Call(ctx, "GET", "user.getTopAlbums", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.TopAlbums.Album, (*Attr)(&resp.TopAlbums.Attr), nil
}

// GetTopArtists gets the top artists listened to by a user.
// See: http://www.last.fm/api/show/user.getTopArtists
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options (e.g. period, page, limit).
//
// Returns:
//   - ArtistList: A slice of top artists.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *UserService) GetTopArtists(ctx context.Context, user string, options url.Values) (ArtistList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp userGetTopArtistsResponse
	err := s.client.Call(ctx, "GET", "user.getTopArtists", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.TopArtists.Artist, (*Attr)(&resp.TopArtists.Attr), nil
}

// GetTopTracks gets the top tracks listened to by a user.
// See: http://www.last.fm/api/show/user.getTopTracks
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options (e.g. period, page, limit).
//
// Returns:
//   - TrackList: A slice of top tracks.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *UserService) GetTopTracks(ctx context.Context, user string, options url.Values) (TrackList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp userGetTopTracksResponse
	err := s.client.Call(ctx, "GET", "user.getTopTracks", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.TopTracks.Track, (*Attr)(&resp.TopTracks.Attr), nil
}

// GetLovedTracks gets the loved tracks for a user.
// See: http://www.last.fm/api/show/user.getLovedTracks
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - TrackList: A slice of loved tracks.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *UserService) GetLovedTracks(ctx context.Context, user string, options url.Values) (TrackList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp userGetLovedTracksResponse
	err := s.client.Call(ctx, "GET", "user.getLovedTracks", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.LovedTracks.Track, (*Attr)(&resp.LovedTracks.Attr), nil
}

// GetFriends gets the friends for a user.
// See: http://www.last.fm/api/show/user.getFriends
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options (e.g. recenttracks, page, limit).
//
// Returns:
//   - FriendList: A slice of friends.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *UserService) GetFriends(ctx context.Context, user string, options url.Values) (FriendList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp userGetFriendsResponse
	err := s.client.Call(ctx, "GET", "user.getFriends", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Friends.User, (*Attr)(&resp.Friends.Attr), nil
}

// UserTaggings represents the results of a personal tags request.
type UserTaggings struct {
	Artists ArtistList
	Albums  AlbumList
	Tracks  TrackList
}

// GetPersonalTags gets the user's personal tags.
// See: http://www.last.fm/api/show/user.getPersonalTags
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - tag: The tag name.
//   - taggingType: The type of items to return (artist, album, or track).
//   - options: Additional options (e.g. page, limit).
//
// Returns:
//   - *UserTaggings: The tagged items.
//   - *Attr: Pagination metadata.
//   - error: Error if the request fails.
func (s *UserService) GetPersonalTags(ctx context.Context, user, tag, taggingType string, options url.Values) (*UserTaggings, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)
	params.Set("tag", tag)
	params.Set("taggingtype", taggingType)

	var resp userGetPersonalTagsResponse
	err := s.client.Call(ctx, "GET", "user.getPersonalTags", params, &resp)
	if err != nil {
		return nil, nil, err
	}

	taggings := &UserTaggings{}
	if resp.Taggings.Artists != nil {
		taggings.Artists = resp.Taggings.Artists.Artist
	}
	if resp.Taggings.Albums != nil {
		taggings.Albums = resp.Taggings.Albums.Album
	}
	if resp.Taggings.Tracks != nil {
		taggings.Tracks = resp.Taggings.Tracks.Track
	}

	return taggings, (*Attr)(&resp.Taggings.Attr), nil
}

// GetTopTags gets the top tags for a user.
// See: http://www.last.fm/api/show/user.getTopTags
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options.
//
// Returns:
//   - TagList: A slice of top tags.
//   - error: Error if the request fails.
func (s *UserService) GetTopTags(ctx context.Context, user string, options url.Values) (TagList, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp userGetTopTagsResponse
	err := s.client.Call(ctx, "GET", "user.getTopTags", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.TopTags.Tag, nil
}

// GetWeeklyAlbumChart gets the weekly album chart for a user.
// See: http://www.last.fm/api/show/user.getWeeklyAlbumChart
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options (e.g. from, to).
//
// Returns:
//   - AlbumList: A slice of albums in the chart.
//   - *Attr: Metadata (user, from, to).
//   - error: Error if the request fails.
func (s *UserService) GetWeeklyAlbumChart(ctx context.Context, user string, options url.Values) (AlbumList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp userGetWeeklyAlbumChartResponse
	err := s.client.Call(ctx, "GET", "user.getWeeklyAlbumChart", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.WeeklyAlbumChart.Album, (*Attr)(&resp.WeeklyAlbumChart.Attr), nil
}

// GetWeeklyArtistChart gets the weekly artist chart for a user.
// See: http://www.last.fm/api/show/user.getWeeklyArtistChart
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options (e.g. from, to).
//
// Returns:
//   - ArtistList: A slice of artists in the chart.
//   - *Attr: Metadata (user, from, to).
//   - error: Error if the request fails.
func (s *UserService) GetWeeklyArtistChart(ctx context.Context, user string, options url.Values) (ArtistList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp userGetWeeklyArtistChartResponse
	err := s.client.Call(ctx, "GET", "user.getWeeklyArtistChart", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.WeeklyArtistChart.Artist, (*Attr)(&resp.WeeklyArtistChart.Attr), nil
}

// GetWeeklyChartList gets a list of available weekly charts for a user.
// See: http://www.last.fm/api/show/user.getWeeklyChartList
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//
// Returns:
//   - ChartList: A slice of available charts.
//   - error: Error if the request fails.
func (s *UserService) GetWeeklyChartList(ctx context.Context, user string) (ChartList, error) {
	params := url.Values{}
	params.Set("user", user)

	var resp userGetWeeklyChartListResponse
	err := s.client.Call(ctx, "GET", "user.getWeeklyChartList", params, &resp)
	if err != nil {
		return nil, err
	}
	return resp.WeeklyChartList.Chart, nil
}

// GetWeeklyTrackChart gets the weekly track chart for a user.
// See: http://www.last.fm/api/show/user.getWeeklyTrackChart
//
// Parameters:
//   - ctx: Context for the request.
//   - user: The user name.
//   - options: Additional options (e.g. from, to).
//
// Returns:
//   - TrackList: A slice of tracks in the chart.
//   - *Attr: Metadata (user, from, to).
//   - error: Error if the request fails.
func (s *UserService) GetWeeklyTrackChart(ctx context.Context, user string, options url.Values) (TrackList, *Attr, error) {
	params := url.Values{}
	for k, v := range options {
		params[k] = v
	}
	params.Set("user", user)

	var resp userGetWeeklyTrackChartResponse
	err := s.client.Call(ctx, "GET", "user.getWeeklyTrackChart", params, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.WeeklyTrackChart.Track, (*Attr)(&resp.WeeklyTrackChart.Attr), nil
}

// userGetInfoResponse is the response from user.getInfo.
type userGetInfoResponse struct {
	User User `json:"user"`
}

// userGetRecentTracksResponse is the response from user.getRecentTracks.
type userGetRecentTracksResponse struct {
	RecentTracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"recenttracks"`
}

// userGetTopAlbumsResponse is the response from user.getTopAlbums.
type userGetTopAlbumsResponse struct {
	TopAlbums struct {
		Album AlbumList `json:"album"`
		Attr  Attr      `json:"@attr"`
	} `json:"topalbums"`
}

// userGetTopArtistsResponse is the response from user.getTopArtists.
type userGetTopArtistsResponse struct {
	TopArtists struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"topartists"`
}

// userGetTopTracksResponse is the response from user.getTopTracks.
type userGetTopTracksResponse struct {
	TopTracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"toptracks"`
}

// userGetLovedTracksResponse is the response from user.getLovedTracks.
type userGetLovedTracksResponse struct {
	LovedTracks struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"lovedtracks"`
}

// userGetFriendsResponse is the response from user.getFriends.
type userGetFriendsResponse struct {
	Friends struct {
		User FriendList `json:"user"`
		Attr Attr       `json:"@attr"`
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

// userGetPersonalTagsResponse is the response from user.getPersonalTags.
type userGetPersonalTagsResponse struct {
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
		Attr Attr `json:"@attr"`
	} `json:"taggings"`
}

// userGetTopTagsResponse is the response from user.getTopTags.
type userGetTopTagsResponse struct {
	TopTags struct {
		Tag  TagList `json:"tag"`
		Attr Attr    `json:"@attr"`
	} `json:"toptags"`
}

// userGetWeeklyAlbumChartResponse is the response from user.getWeeklyAlbumChart.
type userGetWeeklyAlbumChartResponse struct {
	WeeklyAlbumChart struct {
		Album AlbumList `json:"album"`
		Attr  Attr      `json:"@attr"`
	} `json:"weeklyalbumchart"`
}

// userGetWeeklyArtistChartResponse is the response from user.getWeeklyArtistChart.
type userGetWeeklyArtistChartResponse struct {
	WeeklyArtistChart struct {
		Artist ArtistList `json:"artist"`
		Attr   Attr       `json:"@attr"`
	} `json:"weeklyartistchart"`
}

// userGetWeeklyChartListResponse is the response from user.getWeeklyChartList.
type userGetWeeklyChartListResponse struct {
	WeeklyChartList struct {
		Chart ChartList `json:"chart"`
		Attr  Attr      `json:"@attr"`
	} `json:"weeklychartlist"`
}

// userGetWeeklyTrackChartResponse is the response from user.getWeeklyTrackChart.
type userGetWeeklyTrackChartResponse struct {
	WeeklyTrackChart struct {
		Track TrackList `json:"track"`
		Attr  Attr      `json:"@attr"`
	} `json:"weeklytrackchart"`
}
