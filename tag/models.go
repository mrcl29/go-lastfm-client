package tag

import "encoding/json"

// Tag represents a Last.fm tag.
type Tag struct {
	Name       string      `json:"name"`
	Count      json.Number `json:"count,omitempty"`
	Reach      json.Number `json:"reach,omitempty"`
	Taggings   json.Number `json:"taggings,omitempty"`
	Streamable string      `json:"streamable,omitempty"`
	URL        string      `json:"url"`
	Wiki       Wiki        `json:"wiki,omitempty"`
}

// Wiki represents the wiki information for a tag.
type Wiki struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

// Album represents an album in a tag response.
type Album struct {
	Name       string  `json:"name"`
	MBID       string  `json:"mbid"`
	URL        string  `json:"url"`
	Artist     Artist  `json:"artist"`
	Image      []Image `json:"image"`
	Attr       Attr    `json:"@attr"`
}

// Artist represents an artist in a tag response.
type Artist struct {
	Name       string  `json:"name"`
	MBID       string  `json:"mbid"`
	URL        string  `json:"url"`
	Streamable string  `json:"streamable"`
	Image      []Image `json:"image"`
	Attr       Attr    `json:"@attr,omitempty"`
}

// Track represents a track in a tag response.
type Track struct {
	Name       string     `json:"name"`
	Duration   string     `json:"duration"`
	MBID       string     `json:"mbid"`
	URL        string     `json:"url"`
	Streamable Streamable `json:"streamable"`
	Artist     Artist     `json:"artist"`
	Image      []Image    `json:"image"`
	Attr       Attr       `json:"@attr"`
}

// Image represents a Last.fm image.
type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

// Streamable represents streamable status.
type Streamable struct {
	Text      string `json:"#text"`
	FullTrack string `json:"fulltrack"`
}

// Attr represents metadata for tag entity responses.
type Attr struct {
	Rank string `json:"rank,omitempty"`
}

// ResponseAttr represents pagination metadata.
type ResponseAttr struct {
	Tag        string `json:"tag,omitempty"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

// GetInfoResponse is the response from tag.getInfo.
type GetInfoResponse struct {
	Tag Tag `json:"tag"`
}

// GetSimilarResponse is the response from tag.getSimilar.
type GetSimilarResponse struct {
	SimilarTags struct {
		Tag  []Tag `json:"tag"`
		Attr struct {
			Tag string `json:"tag"`
		} `json:"@attr"`
	} `json:"similartags"`
}

// GetTopAlbumsResponse is the response from tag.getTopAlbums.
type GetTopAlbumsResponse struct {
	Albums struct {
		Album []Album      `json:"album"`
		Attr  ResponseAttr `json:"@attr"`
	} `json:"albums"`
}

// GetTopArtistsResponse is the response from tag.getTopArtists.
type GetTopArtistsResponse struct {
	TopArtists struct {
		Artist []Artist     `json:"artist"`
		Attr   ResponseAttr `json:"@attr"`
	} `json:"topartists"`
}

// GetTopTagsResponse is the response from tag.getTopTags.
type GetTopTagsResponse struct {
	TopTags struct {
		Tag  []Tag `json:"tag"`
		Attr struct {
			NumRes     int `json:"num_res"`
			Offset     int `json:"offset"`
			Total      int `json:"total"`
		} `json:"@attr"`
	} `json:"toptags"`
}

// GetTopTracksResponse is the response from tag.getTopTracks.
type GetTopTracksResponse struct {
	Tracks struct {
		Track []Track      `json:"track"`
		Attr  ResponseAttr `json:"@attr"`
	} `json:"tracks"`
}

// GetWeeklyChartListResponse is the response from tag.getWeeklyChartList.
type GetWeeklyChartListResponse struct {
	WeeklyChartList struct {
		Chart []struct {
			Text string `json:"#text"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"chart"`
		Attr struct {
			Tag string `json:"tag"`
		} `json:"@attr"`
	} `json:"weeklychartlist"`
}
