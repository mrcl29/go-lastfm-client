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
	// 1. Setup credentials from Environment
	apiKey := os.Getenv("LASTFM_API_KEY")
	apiSecret := os.Getenv("LASTFM_API_SECRET")
	username := os.Getenv("LASTFM_USERNAME")
	password := os.Getenv("LASTFM_PASSWORD")

	if apiKey == "" || apiSecret == "" || username == "" || password == "" {
		log.Fatal("LASTFM_API_KEY, LASTFM_API_SECRET, LASTFM_USERNAME, and LASTFM_PASSWORD must be set")
	}

	ctx := context.Background()
	client := lastfm.New(apiKey, apiSecret)

	// 2. Authenticate (Mobile flow for script convenience)
	fmt.Printf("--- Authenticating user: %s ---\n", username)
	session, err := client.Auth.GetMobileSession(ctx, username, password)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}
	fmt.Printf("✓ Successfully authenticated! Session Key: %s...\n\n", session.Key[:10])

	// Create an authenticated client for the rest of the steps
	authClient := lastfm.New(apiKey, apiSecret, lastfm.WithSessionKey(session.Key))

	// 3. User Info
	fmt.Println("--- Fetching User Profile ---")
	userRes, err := authClient.User.GetInfo(ctx, "", nil)
	if err != nil {
		log.Fatalf("Failed to get user info: %v", err)
	}
	fmt.Printf("User: %s (Playcount: %s, Registered: %s)\n\n", 
		userRes.User.Name, userRes.User.Playcount, userRes.User.Registered.Text)

	// 4. Realistic Discovery Flow
	artistName := "Daft Punk"
	fmt.Printf("--- Discovering Artist: %s ---\n", artistName)
	
	// Get Artist Info
	artistInfo, err := authClient.Artist.GetInfo(ctx, artistName, nil)
	if err != nil {
		log.Fatalf("Failed to get artist info: %v", err)
	}
	fmt.Printf("Artist: %s\nSummary: %s\n", artistInfo.Artist.Name, artistInfo.Artist.Bio.Summary)

	// Get Top Albums
	topAlbums, err := authClient.Artist.GetTopAlbums(ctx, artistName, nil)
	if err != nil {
		log.Fatalf("Failed to get top albums: %v", err)
	}
	if len(topAlbums.TopAlbums.Album) > 0 {
		bestAlbum := topAlbums.TopAlbums.Album[0]
		fmt.Printf("\nMost popular album: %s (Plays: %s)\n", bestAlbum.Name, bestAlbum.Playcount)

		// Get Album Info (with tracks)
		albumDetails, err := authClient.Album.GetInfo(ctx, artistName, bestAlbum.Name, nil)
		if err != nil {
			log.Fatalf("Failed to get album details: %v", err)
		}
		fmt.Printf("Tracks in '%s':\n", bestAlbum.Name)
		for i, track := range albumDetails.Album.Tracks.Track {
			fmt.Printf("  %d. %s\n", i+1, track.Name)
			if i >= 4 { // Just show first 5
				break
			}
		}
	}

	// 5. Scrobbling Flow
	scrobbleArtist := "Daft Punk"
	scrobbleTrack := "Around the World"
	fmt.Printf("\n--- Simulation: Listening to '%s' by '%s' ---\n", scrobbleTrack, scrobbleArtist)

	// Update Now Playing
	fmt.Println("Updating 'Now Playing' status...")
	npRes, err := authClient.Track.UpdateNowPlaying(ctx, scrobbleArtist, scrobbleTrack, nil)
	if err != nil {
		log.Fatalf("Failed to update now playing: %v", err)
	}
	fmt.Printf("✓ Last.fm confirmed: Now playing %s\n", npRes.NowPlaying.Track.Text)

	// Wait a moment to simulate listening (optional)
	fmt.Println("Waiting 2 seconds before scrobbling...")
	time.Sleep(2 * time.Second)

	// Final Scrobble
	fmt.Println("Sending final scrobble...")
	timestamp := time.Now().Unix()
	scrobbleRes, err := authClient.Track.Scrobble(ctx, scrobbleArtist, scrobbleTrack, timestamp, nil)
	if err != nil {
		log.Fatalf("Scrobble failed: %v", err)
	}
	
	if scrobbleRes.Scrobbles.Attr.Accepted > 0 {
		fmt.Printf("✓ Scrobble Accepted! (Artist: %s, Track: %s, Timestamp: %s)\n", 
			scrobbleRes.Scrobbles.Scrobble[0].Artist.Text,
			scrobbleRes.Scrobbles.Scrobble[0].Track.Text,
			scrobbleRes.Scrobbles.Scrobble[0].Timestamp)
	}

	fmt.Println("\n--- Flow completed successfully! ---")
}
