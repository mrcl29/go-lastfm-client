# go-lastfm-client

[![Go Version](https://img.shields.io/github/go-mod/go-version/mrcl29/go-lastfm-client)](https://golang.org/)
[![License](https://img.shields.io/github/license/mrcl29/go-lastfm-client)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrcl29/go-lastfm-client)](https://goreportcard.com/report/github.com/mrcl29/go-lastfm-client)

A professional, feature-rich, and idiomatic Go client for the [Last.fm API](https://www.last.fm/api). 

Built with scalability and resilience in mind, this library handles the complexities of the Last.fm API—like its polymorphic JSON responses—so you can focus on building your application.

## ✨ Features

- **Full Service Coverage**: Implementation for `Album`, `Artist`, `Track`, `User`, `Tag`, `Geo`, `Chart`, and `Library` services.
- **Polymorphic JSON Resilience**: Automatically handles Last.fm's "object or array" response inconsistency using custom unmarshaling logic.
- **Context Support**: Every API call supports `context.Context` for proper timeout and cancellation management.
- **Authentication**: Easy-to-use support for both Desktop (Token-based) and Mobile (Credentials-based) authentication flows.
- **Functional Options**: Clean client initialization using the functional options pattern.
- **Request Signing**: Automatic generation of `api_sig` for all authenticated methods.
- **Modern Go**: Leverages Go generics for internal utilities and follows standard project layout conventions.

## 🚀 Installation

```bash
go get github.com/mrcl29/go-lastfm-client
```

## 🛠️ Quick Start

### 1. Unauthenticated Metadata Retrieval

```go
package main

import (
	"context"
	"fmt"
	"github.com/mrcl29/go-lastfm-client"
)

func main() {
	client := lastfm.New("YOUR_API_KEY", "YOUR_API_SECRET")
	ctx := context.Background()

	// Search for an artist
	res, _ := client.Artist.Search(ctx, "Cher", nil)
	
	for _, artist := range res.Results.ArtistMatches.Artist {
		fmt.Printf("Artist: %s (Listeners: %s)\n", artist.Name, artist.Stats.Listeners)
	}
}
```

### 2. Authenticated Operations (Scrobbling)

```go
package main

import (
	"context"
	"time"
	"github.com/mrcl29/go-lastfm-client"
)

func main() {
	// Initialize with a session key
	client := lastfm.New("API_KEY", "API_SECRET", lastfm.WithSessionKey("USER_SESSION_KEY"))
	ctx := context.Background()

	// Update Now Playing
	_, _ = client.Track.UpdateNowPlaying(ctx, "Daft Punk", "Around the World", nil)

	// Send Scrobble
	timestamp := time.Now().Unix()
	_, _ = client.Track.Scrobble(ctx, "Daft Punk", "Around the World", timestamp, nil)
}
```

## 🏗️ Architecture & Design

### Polymorphic JSON Handling
Last.fm's API is known for returning a single object when there is one result and an array when there are many. This library solves this globally using a generic `internal/jsonutil` package that ensures your code always receives a slice, regardless of the raw JSON structure.

### Service-Oriented Design
The client is organized into domain-specific services, accessible directly from the main `Client` struct:

- `client.Album`: Metadata, tags, search.
- `client.Artist`: Bio, similar artists, top albums/tracks.
- `client.Track`: Scrobbling, "Now Playing", corrections, similarity.
- `client.User`: Profile, recent tracks, loved tracks, personal tags, weekly charts.
- `client.Tag`: Tag metadata, top artists/tracks by tag.
- `client.Library`: User library discovery.
- `client.Auth`: Token and Session management.

## 🧪 Testing

The library is designed with testability as a core principle. All services use an `APIClient` interface, allowing for easy mocking. We use `github.com/stretchr/testify` for all assertions.

Run tests:
```bash
go test ./...
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
