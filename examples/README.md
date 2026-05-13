# go-lastfm-client Examples

This directory contains examples of how to use the `go-lastfm-client` library.

## Prerequisites

You will need a Last.fm API Key and Secret. You can get them by creating an API account at [https://www.last.fm/api/account/create](https://www.last.fm/api/account/create).

## Running the Examples

Set the following environment variables before running the examples:

```bash
export LASTFM_API_KEY="your_api_key"
export LASTFM_API_SECRET="your_api_secret"
export LASTFM_USERNAME="your_username"
export LASTFM_PASSWORD="your_password"
export LASTFM_SESSION_KEY="your_session_key" # Optional, can be obtained via auth examples
```

### 1. Music Discovery (Read-only)

Demonstrates how to fetch charts, similar artists, and top tracks for a tag. Only requires `LASTFM_API_KEY`.

```bash
go run examples/music_discovery/main.go
```

### 2. Mobile Authentication

Demonstrates how to obtain a session key using a username and password. Requires `LASTFM_API_KEY`, `LASTFM_API_SECRET`, `LASTFM_USERNAME`, and `LASTFM_PASSWORD`.

```bash
go run examples/auth_mobile/main.go
```

### 3. Desktop Authentication

Demonstrates the web-based authentication flow where the user authorizes a token in their browser. Requires `LASTFM_API_KEY` and `LASTFM_API_SECRET`.

```bash
go run examples/auth_desktop/main.go
```

### 4. Scrobble and Love

Demonstrates how to update the "Now Playing" status, scrobble a track, and mark it as loved. Requires `LASTFM_API_KEY`, `LASTFM_API_SECRET`, and `LASTFM_SESSION_KEY`.

```bash
go run examples/scrobble_and_love/main.go
```
