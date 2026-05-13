package album

import "encoding/json"

// Album represents a Last.fm album.
type Album struct {
	Name        string  `json:"name"`
	Artist      string  `json:"artist"`
	ID          string  `json:"id"`
	MBID        string  `json:"mbid"`
	URL         string  `json:"url"`
	ReleaseDate string  `json:"releasedate"`
	Image       []Image `json:"image"`
	Listeners   string  `json:"listeners"`
	Playcount   string  `json:"playcount"`
	TopTags     TopTags `json:"toptags"`
	Tracks      Tracks  `json:"tracks"`
	Wiki        Wiki    `json:"wiki"`
	Streamable  string  `json:"streamable,omitempty"`
}

// Image represents a Last.fm image.
type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

// TopTags represents top tags for an album.
type TopTags struct {
	Tag TagList `json:"tag"`
}

// TagList is a slice of Tag that handles both a single object and an array in JSON.
type TagList []Tag

// UnmarshalJSON implements json.Unmarshaler.
func (t *TagList) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '{' {
		var r Tag
		if err := json.Unmarshal(data, &r); err != nil {
			return err
		}
		*t = []Tag{r}
		return nil
	}
	var r []Tag
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*t = r
	return nil
}

// Tag represents a single tag.
type Tag struct {
	Name  string      `json:"name"`
	URL   string      `json:"url"`
	Count json.Number `json:"count,omitempty"`
}

// Tracks represents the tracklist of an album.
type Tracks struct {
	Track TrackList `json:"track"`
}

// TrackList is a slice of Track that handles both a single object and an array in JSON.
type TrackList []Track

// UnmarshalJSON implements json.Unmarshaler.
func (t *TrackList) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '{' {
		var r Track
		if err := json.Unmarshal(data, &r); err != nil {
			return err
		}
		*t = []Track{r}
		return nil
	}
	var r []Track
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*t = r
	return nil
}

// Track represents a track on an album.
type Track struct {
	Name       string      `json:"name"`
	Duration   json.Number `json:"duration"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Streamable Streamable  `json:"streamable"`
	Artist     Artist      `json:"artist"`
	Rank       json.Number `json:"rank"`
}

// Streamable represents streamable status.
type Streamable struct {
	Text      string `json:"#text"`
	FullTrack string `json:"fulltrack"`
}

// Artist represents a track artist.
type Artist struct {
	Name string `json:"name"`
	MBID string `json:"mbid"`
	URL  string `json:"url"`
}

// Wiki represents the wiki information for an album.
type Wiki struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

// GetInfoResponse is the response from album.getInfo.
type GetInfoResponse struct {
	Album Album `json:"album"`
}

// SearchResponse is the response from album.search.
type SearchResponse struct {
	Results struct {
		AlbumMatches struct {
			Album AlbumList `json:"album"`
		} `json:"albummatches"`
	} `json:"results"`
}

// AlbumList is a slice of Album that handles both a single object and an array in JSON.
type AlbumList []Album

// UnmarshalJSON implements json.Unmarshaler.
func (a *AlbumList) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '{' {
		var r Album
		if err := json.Unmarshal(data, &r); err != nil {
			return err
		}
		*a = []Album{r}
		return nil
	}
	var r []Album
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*a = r
	return nil
}

// GetTagsResponse is the response from album.getTags.
type GetTagsResponse struct {
	Tags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
		Album  string  `json:"album"`
	} `json:"tags"`
}

// GetTopTagsResponse is the response from album.getTopTags.
type GetTopTagsResponse struct {
	TopTags struct {
		Tag    TagList `json:"tag"`
		Artist string  `json:"artist"`
		Album  string  `json:"album"`
	} `json:"toptags"`
}

