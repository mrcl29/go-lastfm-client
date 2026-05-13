# go-lastfm-client

[![Go Version](https://img.shields.io/github/go-mod/go-version/mrcl29/go-lastfm-client)](https://golang.org/)
[![License](https://img.shields.io/github/license/mrcl29/go-lastfm-client)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrcl29/go-lastfm-client)](https://goreportcard.com/report/github.com/mrcl29/go-lastfm-client)
[![Go Reference](https://pkg.go.dev/badge/github.com/mrcl29/go-lastfm-client.svg)](https://pkg.go.dev/github.com/mrcl29/go-lastfm-client)
[![CI](https://github.com/mrcl29/lastfm/actions/workflows/ci.yml/badge.svg)](https://github.com/mrcl29/lastfm/actions/workflows/ci.yml)

A professional, feature-rich, and idiomatic Go client for the [Last.fm API](https://www.last.fm/api).

Built with scalability and resilience in mind, this library handles the complexities of the Last.fm API—like its polymorphic JSON responses—so you can focus on building your application.

## ✨ Features

- **Full Service Coverage**: Implementation for `Album`, `Artist`, `Track`, `User`, `Tag`, `Geo`, `Chart`, and `Library` services.
- **Polymorphic JSON Resilience**: Automatically handles Last.fm's "object or array" response inconsistency using custom unmarshaling logic.
- **Context Support**: Every API call supports `context.Context` for proper timeout and cancellation management.
- **Authentication**: Easy-to-use support for both Desktop (Token-based) and Mobile (Credentials-based) authentication flows.
- **Functional Options**: Clean client initialization using the functional options pattern.
- **Request Signing**: Automatic generation of `api_sig` for all authenticated methods.
- **Modern Go**: Leverages Go 1.26 features and follows standard project layout conventions.

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
    // API Secret is optional for unauthenticated calls
    client := golastfmclient.New("YOUR_API_KEY", "")
    ctx := context.Background()

    // Search for an artist
    res, err := client.Artist.Search(ctx, "Cher", nil)
    if err != nil {
        panic(err)
    }
    
    for _, artist := range res.Results.ArtistMatches.Artist {
        fmt.Printf("Artist: %s (Listeners: %s)\n", artist.Name, artist.Listeners.String())
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
    client := golastfmclient.New("API_KEY", "API_SECRET", golastfmclient.WithSessionKey("USER_SESSION_KEY"))
    ctx := context.Background()

    // Update Now Playing
    _, err := client.Track.UpdateNowPlaying(ctx, "Daft Punk", "Around the World", nil)
    if err != nil {
        panic(err)
    }

    // Send Scrobble
    timestamp := time.Now().Unix()
    _, err = client.Track.Scrobble(ctx, "Daft Punk", "Around the World", timestamp, nil)
    if err != nil {
        panic(err)
    }
}
```

## 🏗️ Architecture & Design

### Polymorphic JSON Handling

Last.fm's API is known for returning a single object when there is one result and an array when there are many. This library solves this globally using a custom `internal/jsonutil` package that ensures your code always receives a slice, regardless of the raw JSON structure.

### Service-Oriented Design

The client is organized into domain-specific services, accessible directly from the main `Client` struct:

- `client.Album`: Metadata, tags, search.
- `client.Artist`: Bio, similar artists, top albums/tracks.
- `client.Track`: Scrobbling, "Now Playing", corrections, similarity.
- `client.User`: Profile, recent tracks, loved tracks, personal tags, weekly charts.
- `client.Tag`: Tag metadata, top artists/tracks by tag.
- `client.Library`: User library discovery.
- `client.Chart`: Global charts for artists, tracks, and tags.
- `client.Geo`: Top artists and tracks by country.
- `client.Auth`: Token and Session management.

## 📚 Examples

Check the [examples/](examples/) directory for more detailed use cases:

- [Music Discovery](examples/music_discovery/): Charts, similarity, and tag-based search.
- [Mobile Auth](examples/auth_mobile/): Authenticating with username/password.
- [Desktop Auth](examples/auth_desktop/): Web-based token authorization.
- [Scrobble & Love](examples/scrobble_and_love/): Practical scrobbling flow.

## 🧪 Testing

The library is designed with testability as a core principle. All services use an `APIClient` interface, allowing for easy mocking.

Run tests:

```bash
go test ./...
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
