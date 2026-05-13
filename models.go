package golastfmclient

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

// Artist represents a Last.fm artist.
type Artist struct {
	Name       string      `json:"name"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url,omitempty"`
	Image      []Image     `json:"image,omitempty"`
	Streamable string      `json:"streamable,omitempty"`
	Stats      *Stats      `json:"stats,omitempty"`
	Similar    *Similar    `json:"similar,omitempty"`
	Tags       *Tags       `json:"tags,omitempty"`
	Bio        *Bio        `json:"bio,omitempty"`
	Match      float64     `json:"match,string,omitempty"`
	Listeners  json.Number `json:"listeners,omitempty"`
	Playcount  json.Number `json:"playcount,omitempty"`
	Tagcount   json.Number `json:"tagcount,omitempty"`
	Rank       json.Number `json:"rank,omitempty"`
}

// ArtistList is a slice of Artist that handles both a single object and an array in JSON.
type ArtistList []Artist

// UnmarshalJSON implements json.Unmarshaler.
func (a *ArtistList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Artist)(a))
}

// Album represents a Last.fm album.
type Album struct {
	Name        string      `json:"name,omitempty"`
	Title       string      `json:"title,omitempty"`  // Some endpoints use "title"
	Text        string      `json:"#text,omitempty"`  // Some endpoints use "#text" for name
	Artist      interface{} `json:"artist,omitempty"` // Can be string or Artist struct
	ID          string      `json:"id,omitempty"`
	MBID        string      `json:"mbid,omitempty"`
	URL         string      `json:"url,omitempty"`
	ReleaseDate string      `json:"releasedate,omitempty"`
	Image       []Image     `json:"image,omitempty"`
	Listeners   string      `json:"listeners,omitempty"`
	Playcount   json.Number `json:"playcount,omitempty"`
	TopTags     *TopTags    `json:"toptags,omitempty"`
	Tracks      *Tracks     `json:"tracks,omitempty"`
	Wiki        *Wiki       `json:"wiki,omitempty"`
	Streamable  string      `json:"streamable,omitempty"`
	Rank        json.Number `json:"rank,omitempty"`
}

// AlbumList is a slice of Album that handles both a single object and an array in JSON.
type AlbumList []Album

// UnmarshalJSON implements json.Unmarshaler.
func (a *AlbumList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Album)(a))
}

// Track represents a Last.fm track.
type Track struct {
	Name       string      `json:"name"`
	Duration   json.Number `json:"duration,omitempty"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Streamable interface{} `json:"streamable,omitempty"` // Can be string or Streamable struct
	Artist     interface{} `json:"artist,omitempty"`     // Can be string or Artist struct
	Album      interface{} `json:"album,omitempty"`      // Can be string or Album struct
	Rank       json.Number `json:"rank,omitempty"`
	Playcount  json.Number `json:"playcount,omitempty"`
	Listeners  json.Number `json:"listeners,omitempty"`
	Image      []Image     `json:"image,omitempty"`
	TopTags    *TopTags    `json:"toptags,omitempty"`
	Wiki       *Wiki       `json:"wiki,omitempty"`
	Date       *Date       `json:"date,omitempty"`
	Attr       *TrackAttr  `json:"@attr,omitempty"`
}

// TrackList is a slice of Track that handles both a single object and an array in JSON.
type TrackList []Track

// UnmarshalJSON implements json.Unmarshaler.
func (t *TrackList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Track)(t))
}

// Tag represents a Last.fm tag.
type Tag struct {
	Name       string      `json:"name"`
	URL        string      `json:"url"`
	Count      json.Number `json:"count,omitempty"`
	Reach      json.Number `json:"reach,omitempty"`
	Taggings   json.Number `json:"taggings,omitempty"`
	Streamable string      `json:"streamable,omitempty"`
	Wiki       *Wiki       `json:"wiki,omitempty"`
}

// TagList is a slice of Tag that handles both a single object and an array in JSON.
type TagList []Tag

// UnmarshalJSON implements json.Unmarshaler.
func (t *TagList) UnmarshalJSON(data []byte) error {
	return jsonutil.UnmarshalList(data, (*[]Tag)(t))
}

// Image represents a Last.fm image.
type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

// Wiki represents the wiki information for an entity.
type Wiki struct {
	Published string `json:"published,omitempty"`
	Summary   string `json:"summary,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Streamable represents streamable status.
type Streamable struct {
	Text      string `json:"#text"`
	FullTrack string `json:"fulltrack"`
}

// Stats represents statistics for an artist.
type Stats struct {
	Listeners json.Number `json:"listeners"`
	Playcount json.Number `json:"playcount"`
}

// Similar represents similar artists or tags.
type Similar struct {
	Artist ArtistList `json:"artist,omitempty"`
	Tag    TagList    `json:"tag,omitempty"`
}

// Tags represents tags for an artist.
type Tags struct {
	Tag TagList `json:"tag"`
}

// TopTags represents top tags for an entity.
type TopTags struct {
	Tag TagList `json:"tag"`
}

// Tracks represents the tracklist of an album.
type Tracks struct {
	Track TrackList `json:"track"`
}

// Bio represents biography information for an artist.
type Bio struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

// Date represents a Last.fm date response.
type Date struct {
	UnixTime json.Number `json:"unixtime"`
	Text     json.Number `json:"#text"`
}

// TrackAttr represents metadata for a track (e.g. nowplaying).
type TrackAttr struct {
	NowPlaying string      `json:"nowplaying"`
	Rank       json.Number `json:"rank,omitempty"`
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

// Attr represents metadata for responses.
type Attr struct {
	Page       json.Number `json:"page,omitempty"`
	PerPage    json.Number `json:"perPage,omitempty"`
	TotalPages json.Number `json:"totalPages,omitempty"`
	Total      json.Number `json:"total,omitempty"`
	User       string      `json:"user,omitempty"`
	Country    string      `json:"country,omitempty"`
	Artist     string      `json:"artist,omitempty"`
	Tag        string      `json:"tag,omitempty"`
	From       string      `json:"from,omitempty"`
	To         string      `json:"to,omitempty"`
}
