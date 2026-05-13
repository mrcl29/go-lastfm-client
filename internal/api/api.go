package api

import (
	"context"
	"fmt"
	"net/url"
)

// Client defines the interface required by the services.
type Client interface {
	Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error
}

// Error represents a Last.fm API error.
type Error struct {
	Code    int    `json:"error"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("lastfm error %d: %s", e.Code, e.Message)
}
