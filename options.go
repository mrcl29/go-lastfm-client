package golastfmclient

import (
	"net/http"
	"strings"
)

// Option is a functional option for configuring the client.
type Option func(*Client)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL sets a custom base URL.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		if !strings.HasSuffix(baseURL, "/") {
			baseURL += "/"
		}
		c.baseURL = baseURL
	}
}

// WithUserAgent sets a custom User-Agent header.
func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

// WithSessionKey sets the session key for authenticated requests.
func WithSessionKey(sessionKey string) Option {
	return func(c *Client) {
		c.sessionKey = sessionKey
	}
}
