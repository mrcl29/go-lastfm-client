package user

import "encoding/json"

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
}

// Album represents a Last.fm album within user responses.
type Album struct {
	Name string `json:"#text"`
	MBID string `json:"mbid"`
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

// GetInfoResponse is the response from user.getInfo.
type GetInfoResponse struct {
	User User `json:"user"`
}

// GetRecentTracksResponse is the response from user.getRecentTracks.
type GetRecentTracksResponse struct {
	RecentTracks struct {
		Track []Track `json:"track"`
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
		Album []struct {
			Name      string      `json:"name"`
			Playcount json.Number `json:"playcount"`
			MBID      string      `json:"mbid"`
			URL       string      `json:"url"`
			Artist    Artist      `json:"artist"`
			Image     []Image     `json:"image"`
		} `json:"album"`
		Attr struct {
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
		Artist []Artist `json:"artist"`
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
		Track []Track `json:"track"`
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
		Track []Track `json:"track"`
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
		User []Friend `json:"user"`
		Attr struct {
			User       string      `json:"user"`
			Page       json.Number `json:"page"`
			PerPage    json.Number `json:"perPage"`
			TotalPages json.Number `json:"totalPages"`
			Total      json.Number `json:"total"`
		} `json:"@attr"`
	} `json:"friends"`
}
