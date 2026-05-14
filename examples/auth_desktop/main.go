package main

import (
	"context"
	"fmt"
	"log"
	"os"

	golastfmclient "github.com/mrcl29/go-lastfm-client"
)

func main() {
	apiKey := os.Getenv("LASTFM_API_KEY")
	apiSecret := os.Getenv("LASTFM_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		log.Fatal("LASTFM_API_KEY and LASTFM_API_SECRET must be set")
	}

	client := golastfmclient.New(apiKey, apiSecret)
	ctx := context.Background()

	// 1. Get a token
	token, err := client.Auth.GetToken(ctx)
	if err != nil {
		log.Fatalf("Error getting token: %v", err)
	}

	// 2. Authorize the token
	authURL := client.Auth.AuthURL(token)
	fmt.Printf("Please authorize the application by visiting this URL:\n%s\n\n", authURL)
	fmt.Println("Press Enter after you have authorized the application...")
	if _, err := fmt.Scanln(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// 3. Get the session
	session, err := client.Auth.GetSession(ctx, token)
	if err != nil {
		log.Fatalf("Error getting session: %v", err)
	}

	fmt.Printf("Successfully authenticated as %s\n", session.Name)
	fmt.Printf("Session Key: %s\n", session.Key)
}
