package user

import (
	"encoding/json"

	"github.com/mrcl29/go-lastfm-client/internal/jsonutil"
)

// User represents a Last.fm user.
type User struct {
	Name       string      `json:"name"`
	RealName   string      `json:"realname"`
	Image      []Image     `json:"image"`
	URL        string      `json:"url"`
	Country    string      `json:"country"`
	Age        json.Number `json:"age"`
	Gender     string      `json:"gender"`
	Subscriber json.Number `json:"subscriber"`
	Playcount  json.Number `json:"playcount"`
	Playlists  json.Number `json:"playlists"`
	Bootstrap  json.Number `json:"bootstrap"`
	Registered Date        `json:"registered"`
}

// Image represents a Last.fm image.
type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

// Date represents a Last.fm date response.
type Date struct {
	UnixTime json.Number `json:"unixtime"`
	Text     json.Number `json:"#text"`
}

// Track represents a Last.fm track in the context of user recent/top tracks.
type Track struct {
	Name       string      `json:"name"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Artist     Artist      `json:"artist"`
	Album      Album       `json:"album,omitempty"`
	Image      []Image     `json:"image,omitempty"`
	Date       Date        `json:"date,omitempty"`
	Attr       *TrackAttr  `json:"@attr,omitempty"`
	Playcount  json.Number `json:"playcount,omitempty"`
	Duration   json.Number `json:"duration,omitempty"`
	Streamable string      `json:"streamable,omitempty"`
}

// TrackList is a slice of Track that handles both a single object and an array in JSON.
type TrackList []Track

// UnmarshalJSON implements json.Unmarshaler.
func (t *TrackList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Track)(t))
}

// TrackAttr represents metadata for a track (e.g. nowplaying).
type TrackAttr struct {
	NowPlaying string      `json:"nowplaying"`
	Rank       json.Number `json:"rank,omitempty"`
}

// Artist represents a Last.fm artist within user responses.
type Artist struct {
	Name      string      `json:"name"`
	MBID      string      `json:"mbid"`
	URL       string      `json:"url,omitempty"`
	Playcount json.Number `json:"playcount,omitempty"`
	Rank      json.Number `json:"rank,omitempty"`
}

// ArtistList is a slice of Artist that handles both a single object and an array in JSON.
type ArtistList []Artist

// UnmarshalJSON implements json.Unmarshaler.
func (a *ArtistList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Artist)(a))
}

// Album represents a Last.fm album within user responses.
type Album struct {
	Name      string      `json:"#text"`
	Title     string      `json:"name,omitempty"` // Some endpoints use "name" instead of "#text"
	MBID      string      `json:"mbid"`
	Playcount json.Number `json:"playcount,omitempty"`
	Artist    *Artist     `json:"artist,omitempty"`
	Image     []Image     `json:"image,omitempty"`
	URL       string      `json:"url,omitempty"`
	Rank      json.Number `json:"rank,omitempty"`
}

// AlbumList is a slice of Album that handles both a single object and an array in JSON.
type AlbumList []Album

// UnmarshalJSON implements json.Unmarshaler.
func (a *AlbumList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Album)(a))
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

// Tag represents a Last.fm tag within user responses.
type Tag struct {
	Name  string      `json:"name"`
	Count json.Number `json:"count"`
	URL   string      `json:"url"`
}

// TagList is a slice of Tag that handles both a single object and an array in JSON.
type TagList []Tag

// UnmarshalJSON implements json.Unmarshaler.
func (t *TagList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Tag)(t))
}

// Chart represents a weekly chart entry.
type Chart struct {
	Text string `json:"#text"`
	From string `json:"from"`
	To   string `json:"to"`
}

// ChartList is a slice of Chart that handles both a single object and an array in JSON.
type ChartList []Chart

// UnmarshalJSON implements json.Unmarshaler.
func (c *ChartList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Chart)(c))
}

// GetInfoResponse is the response from user.getInfo.
type GetInfoResponse struct {
	User User `json:"user"`
}

// GetRecentTracksResponse is the response from user.getRecentTracks.
type GetRecentTracksResponse struct {
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

// GetTopAlbumsResponse is the response from user.getTopAlbums.
type GetTopAlbumsResponse struct {
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

// GetTopArtistsResponse is the response from user.getTopArtists.
type GetTopArtistsResponse struct {
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

// GetTopTracksResponse is the response from user.getTopTracks.
type GetTopTracksResponse struct {
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

// GetLovedTracksResponse is the response from user.getLovedTracks.
type GetLovedTracksResponse struct {
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

// GetFriendsResponse is the response from user.getFriends.
type GetFriendsResponse struct {
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

// GetPersonalTagsResponse is the response from user.getPersonalTags.
type GetPersonalTagsResponse struct {
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

// GetTopTagsResponse is the response from user.getTopTags.
type GetTopTagsResponse struct {
	TopTags struct {
		Tag  TagList `json:"tag"`
		Attr struct {
			User string `json:"user"`
		} `json:"@attr"`
	} `json:"toptags"`
}

// GetWeeklyAlbumChartResponse is the response from user.getWeeklyAlbumChart.
type GetWeeklyAlbumChartResponse struct {
	WeeklyAlbumChart struct {
		Album AlbumList `json:"album"`
		Attr  struct {
			User string `json:"user"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"@attr"`
	} `json:"weeklyalbumchart"`
}

// GetWeeklyArtistChartResponse is the response from user.getWeeklyArtistChart.
type GetWeeklyArtistChartResponse struct {
	WeeklyArtistChart struct {
		Artist ArtistList `json:"artist"`
		Attr   struct {
			User string `json:"user"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"@attr"`
	} `json:"weeklyartistchart"`
}

// GetWeeklyChartListResponse is the response from user.getWeeklyChartList.
type GetWeeklyChartListResponse struct {
	WeeklyChartList struct {
		Chart ChartList `json:"chart"`
		Attr  struct {
			User string `json:"user"`
		} `json:"@attr"`
	} `json:"weeklychartlist"`
}

// GetWeeklyTrackChartResponse is the response from user.getWeeklyTrackChart.
type GetWeeklyTrackChartResponse struct {
	WeeklyTrackChart struct {
		Track TrackList `json:"track"`
		Attr  struct {
			User string `json:"user"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"@attr"`
	} `json:"weeklytrackchart"`
}
