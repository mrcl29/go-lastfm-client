package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mrcl29/go-lastfm-client"
)

func main() {
	apiKey := os.Getenv("LASTFM_API_KEY")
	apiSecret := os.Getenv("LASTFM_API_SECRET")
	sessionKey := os.Getenv("LASTFM_SESSION_KEY")

	if apiKey == "" || apiSecret == "" || sessionKey == "" {
		log.Fatal("LASTFM_API_KEY, LASTFM_API_SECRET and LASTFM_SESSION_KEY must be set")
	}

	// Initialize client with a session key
	client := golastfmclient.New(apiKey, apiSecret, golastfmclient.WithSessionKey(sessionKey))
	ctx := context.Background()

	artist := "Cher"
	track := "Believe"

	// 1. Update Now Playing
	fmt.Printf("Updating now playing: %s - %s...\n", artist, track)
	np, err := client.Track.UpdateNowPlaying(ctx, artist, track, nil)
	if err != nil {
		log.Fatalf("Error updating now playing: %v", err)
	}
	fmt.Printf("Now playing updated: %s\n", np.NowPlaying.Track.Text)

	// 2. Scrobble the track
	timestamp := time.Now().Unix()
	fmt.Printf("Scrobbling: %s - %s at %d...\n", artist, track, timestamp)
	scr, err := client.Track.Scrobble(ctx, artist, track, timestamp, nil)
	if err != nil {
		log.Fatalf("Error scrobbling: %v", err)
	}
	fmt.Printf("Scrobble accepted: %d\n", scr.Scrobbles.Attr.Accepted)

	// 3. Love the track
	fmt.Printf("Loving the track: %s - %s...\n", artist, track)
	err = client.Track.Love(ctx, artist, track)
	if err != nil {
		log.Fatalf("Error loving track: %v", err)
	}
	fmt.Println("Successfully loved the track!")
}
