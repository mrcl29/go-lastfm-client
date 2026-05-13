package artist

import "encoding/json"

// Artist represents a Last.fm artist.
type Artist struct {
	Name       string      `json:"name"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Image      []Image     `json:"image"`
	Streamable string      `json:"streamable"`
	Stats      Stats       `json:"stats"`
	Similar    Similar     `json:"similar"`
	Tags       Tags        `json:"tags"`
	Bio        Bio         `json:"bio"`
	Match      float64     `json:"match,string,omitempty"`
}

// Image represents a Last.fm image.
type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

// Stats represents statistics for an artist.
type Stats struct {
	Listeners json.Number `json:"listeners"`
	Playcount json.Number `json:"playcount"`
}

// Similar represents similar artists.
type Similar struct {
	Artist []Artist `json:"artist"`
}

// Tags represents tags for an artist.
type Tags struct {
	Tag []Tag `json:"tag"`
}

// Tag represents a single tag.
type Tag struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Bio represents biography information for an artist.
type Bio struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

// GetInfoResponse is the response from artist.getInfo.
type GetInfoResponse struct {
	Artist Artist `json:"artist"`
}

// SearchResponse is the response from artist.search.
type SearchResponse struct {
	Results struct {
		ArtistMatches struct {
			Artist []Artist `json:"artist"`
		} `json:"artistmatches"`
	} `json:"results"`
}

// GetSimilarResponse is the response from artist.getSimilar.
type GetSimilarResponse struct {
	SimilarArtists struct {
		Artist []Artist `json:"artist"`
		Attr   struct {
			Artist string `json:"artist"`
		} `json:"@attr"`
	} `json:"similarartists"`
}

// GetCorrectionResponse is the response from artist.getCorrection.
type GetCorrectionResponse struct {
	Corrections struct {
		Correction Correction `json:"correction"`
	} `json:"corrections"`
}

// Correction represents an artist correction.
type Correction struct {
	Artist Artist `json:"artist"`
	Index  string `json:"index"`
}

// GetTopAlbumsResponse is the response from artist.getTopAlbums.
type GetTopAlbumsResponse struct {
	TopAlbums struct {
		Album []Album `json:"album"`
		Attr  struct {
			Artist string `json:"artist"`
			Page   string `json:"page"`
			Total  string `json:"total"`
		} `json:"@attr"`
	} `json:"topalbums"`
}

// Album represents an album in an artist's top albums list.
type Album struct {
	Name      string      `json:"name"`
	Playcount json.Number `json:"playcount"`
	MBID      string      `json:"mbid"`
	URL       string      `json:"url"`
	Artist    Artist      `json:"artist"`
	Image     []Image     `json:"image"`
}

// GetTopTracksResponse is the response from artist.getTopTracks.
type GetTopTracksResponse struct {
	TopTracks struct {
		Track []Track `json:"track"`
		Attr  struct {
			Artist string `json:"artist"`
			Page   string `json:"page"`
			Total  string `json:"total"`
		} `json:"@attr"`
	} `json:"toptracks"`
}

// Track represents a track in an artist's top tracks list.
type Track struct {
	Name      string      `json:"name"`
	Playcount json.Number `json:"playcount"`
	Listeners json.Number `json:"listeners"`
	MBID      string      `json:"mbid"`
	URL       string      `json:"url"`
	Artist    Artist      `json:"artist"`
	Image     []Image     `json:"image"`
}
