package lastfm

import "fmt"

// APIError represents a Last.fm API error.
type APIError struct {
	Code    int    `json:"error"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("lastfm error %d: %s", e.Code, e.Message)
}
