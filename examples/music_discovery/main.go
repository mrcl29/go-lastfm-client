package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/mrcl29/go-lastfm-client"
)

func main() {
	apiKey := os.Getenv("LASTFM_API_KEY")
	if apiKey == "" {
		log.Fatal("LASTFM_API_KEY must be set")
	}

	// For read-only calls, we only need the API Key
	client := golastfmclient.New(apiKey, "")
	ctx := context.Background()

	// 1. Get Top Artists from Chart
	fmt.Println("--- Top Artists Chart ---")
	chart, err := client.Chart.GetTopArtists(ctx, nil)
	if err != nil {
		log.Fatalf("Error getting top artists: %v", err)
	}

	for i, artist := range chart.Artists.Artist {
		if i >= 5 {
			break
		}
		fmt.Printf("%d. %s (%s listeners)\n", i+1, artist.Name, artist.Listeners.String())
	}

	// 2. Get Similar Artists
	artistName := "Radiohead"
	fmt.Printf("\n--- Artists similar to %s ---\n", artistName)
	similar, err := client.Artist.GetSimilar(ctx, artistName, nil)
	if err != nil {
		log.Fatalf("Error getting similar artists: %v", err)
	}

	for i, artist := range similar.SimilarArtists.Artist {
		if i >= 5 {
			break
		}
		fmt.Printf("%d. %s\n", i+1, artist.Name)
	}

	// 3. Get Top Tracks for a Tag
	tagName := "electronic"
	fmt.Printf("\n--- Top tracks tagged as '%s' ---\n", tagName)
	options := url.Values{}
	options.Set("limit", "5")
	tagTracks, err := client.Tag.GetTopTracks(ctx, tagName, options)
	if err != nil {
		log.Fatalf("Error getting top tracks for tag: %v", err)
	}

	for i, track := range tagTracks.Tracks.Track {
		name := ""
		switch a := track.Artist.(type) {
		case string:
			name = a
		case map[string]interface{}:
			if n, ok := a["name"].(string); ok {
				name = n
			}
		}
		fmt.Printf("%d. %s - %s\n", i+1, name, track.Name)
	}
}
