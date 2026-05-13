package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mrcl29/go-lastfm-client"
)

func main() {
	apiKey := os.Getenv("LASTFM_API_KEY")
	apiSecret := os.Getenv("LASTFM_API_SECRET")
	username := os.Getenv("LASTFM_USERNAME")
	password := os.Getenv("LASTFM_PASSWORD")

	if apiKey == "" || apiSecret == "" || username == "" || password == "" {
		log.Fatal("LASTFM_API_KEY, LASTFM_API_SECRET, LASTFM_USERNAME and LASTFM_PASSWORD must be set")
	}

	client := golastfmclient.New(apiKey, apiSecret)

	ctx := context.Background()
	session, err := client.Auth.GetMobileSession(ctx, username, password)
	if err != nil {
		log.Fatalf("Error getting mobile session: %v", err)
	}

	fmt.Printf("Successfully authenticated as %s\n", session.Name)
	fmt.Printf("Session Key: %s\n", session.Key)
}
