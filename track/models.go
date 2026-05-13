package track

import "encoding/json"

// Track represents a Last.fm track.
type Track struct {
	Name       string      `json:"name"`
	Duration   json.Number `json:"duration"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Streamable string      `json:"streamable"`
	Artist     Artist      `json:"artist"`
	Album      Album       `json:"album"`
	TopTags    TopTags     `json:"toptags"`
	Wiki       Wiki        `json:"wiki"`
}

// Artist represents a Last.fm artist within a track.
type Artist struct {
	Name string `json:"name"`
	MBID string `json:"mbid"`
	URL  string `json:"url"`
}

// Album represents a Last.fm album within a track.
type Album struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	MBID   string `json:"mbid"`
	URL    string `json:"url"`
}

// TopTags represents the top tags for a track.
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
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Wiki represents the wiki information for a track.
type Wiki struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

// GetInfoResponse is the response from track.getInfo.
type GetInfoResponse struct {
	Track Track `json:"track"`
}

// SearchResponse is the response from track.search.
type SearchResponse struct {
	Results struct {
		TrackMatches struct {
			Track TrackList `json:"track"`
		} `json:"trackmatches"`
	} `json:"results"`
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

// ScrobbleResponse is the response from track.scrobble.
type ScrobbleResponse struct {
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
	if len(data) > 0 && data[0] == '{' {
		var r ScrobbleResult
		if err := json.Unmarshal(data, &r); err != nil {
			return err
		}
		*s = []ScrobbleResult{r}
		return nil
	}
	var r []ScrobbleResult
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*s = r
	return nil
}

// ScrobbleResult represents the result of a single scrobble.
type ScrobbleResult struct {
	Artist    Correction `json:"artist"`
	Album     Correction `json:"album"`
	Track     Correction `json:"track"`
	Timestamp string     `json:"timestamp"`
}

// Correction represents corrected metadata in a scrobble response.
type Correction struct {
	Text      string `json:"#text"`
	Corrected string `json:"corrected"`
}

// NowPlayingResponse is the response from track.updateNowPlaying.
type NowPlayingResponse struct {
	NowPlaying struct {
		Artist    Correction `json:"artist"`
		Album     Correction `json:"album"`
		Track     Correction `json:"track"`
		Timestamp string     `json:"timestamp"`
	} `json:"nowplaying"`
}
