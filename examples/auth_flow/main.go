package main

import (
	"bufio"
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
	username := os.Getenv("LASTFM_USERNAME")
	password := os.Getenv("LASTFM_PASSWORD")
	sessionKey := os.Getenv("LASTFM_SESSION_KEY")

	if apiKey == "" || apiSecret == "" {
		log.Fatal("LASTFM_API_KEY and LASTFM_API_SECRET must be set")
	}

	ctx := context.Background()
	client := lastfm.New(apiKey, apiSecret)

	fmt.Println("--- 1. Testing Unauthenticated Flux (Artist Search) ---")
	searchRes, err := client.Artist.Search(ctx, "Cher", nil)
	if err != nil {
		log.Fatalf("Search failed: %v", err)
	}
	if len(searchRes.Results.ArtistMatches.Artist) > 0 {
		fmt.Printf("Found artist: %s\n", searchRes.Results.ArtistMatches.Artist[0].Name)
	}

	fmt.Println("\n--- 2. Testing Authenticated Flux (Scrobble) ---")
	sk := sessionKey
	if sk == "" {
		if username == "" || password == "" {
			fmt.Println("Skipping authenticated flux: LASTFM_USERNAME/PASSWORD not set")
		} else {
			session, err := client.Auth.GetMobileSession(ctx, username, password)
			if err != nil {
				log.Fatalf("Mobile authentication failed: %v", err)
			}
			sk = session.Key
			fmt.Printf("Authenticated via Mobile as: %s\n", session.Name)
		}
	}

	if sk != "" {
		authClient := lastfm.New(apiKey, apiSecret, lastfm.WithSessionKey(sk))
		
		// Scrobble
		timestamp := time.Now().Unix()
		scrobbleRes, err := authClient.Track.Scrobble(ctx, "Cher", "Believe", timestamp, nil)
		if err != nil {
			log.Fatalf("Scrobble failed: %v", err)
		}
		fmt.Printf("Scrobble successful! Accepted: %d\n", scrobbleRes.Scrobbles.Attr.Accepted)
	}

	fmt.Println("\n--- 3. Testing Desktop Auth Flow ---")
	token, err := client.Auth.GetToken(ctx)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}

	authURL := client.Auth.AuthURL(token)
	fmt.Printf("Fetched request token: %s\n", token)
	
	fmt.Printf("\n--- ACTION REQUIRED ---\n")
	fmt.Printf("1. Go to: %s\n", authURL)
	fmt.Printf("2. Authorize the application in your browser.\n")
	fmt.Printf("3. Once authorized, come back here and press [Enter] to continue...\n")
	
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	session, err := client.Auth.GetSession(ctx, token)
	if err != nil {
		log.Fatalf("Failed to get session after authorization: %v", err)
	}
	
	fmt.Printf("Successfully authenticated as %s\n", session.Name)
	fmt.Printf("Your Session Key: %s\n", session.Key)
	fmt.Println("\nSuccess!")
}
