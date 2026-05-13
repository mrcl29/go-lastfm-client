package golastfmclient

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	defaultBaseURL   = "https://ws.audioscrobbler.com/2.0/"
	defaultUserAgent = "go-lastfm-client/0.1.0 (github.com/mrcl29/go-lastfm-client)"
)

// APIClient defines the interface required by the services.
type APIClient interface {
	Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error
}

// Client is a Last.fm API client.
type Client struct {
	apiKey     string
	apiSecret  string
	sessionKey string
	baseURL    string
	userAgent  string
	httpClient *http.Client

	Auth    *AuthService
	Track   *TrackService
	Artist  *ArtistService
	Album   *AlbumService
	User    *UserService
	Chart   *ChartService
	Geo     *GeoService
	Tag     *TagService
	Library *LibraryService
}

// New creates a new Last.fm API client.
func New(apiKey string, apiSecret string, opts ...Option) *Client {
	c := &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		baseURL:   defaultBaseURL,
		userAgent: defaultUserAgent,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Auth = &AuthService{client: c}
	c.Track = NewTrackService(c)
	c.Artist = NewArtistService(c)
	c.Album = NewAlbumService(c)
	c.User = NewUserService(c)
	c.Chart = NewChartService(c)
	c.Geo = NewGeoService(c)
	c.Tag = NewTagService(c)
	c.Library = NewLibraryService(c)

	return c
}

// get performs a GET request to the Last.fm API.
func (c *Client) get(ctx context.Context, method string, params url.Values, target interface{}) error {
	return c.call(ctx, http.MethodGet, method, params, target)
}

// post performs a POST request to the Last.fm API.
func (c *Client) post(ctx context.Context, method string, params url.Values, target interface{}) error {
	return c.call(ctx, http.MethodPost, method, params, target)
}

// Call performs a request to the Last.fm API.
func (c *Client) Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error {
	return c.call(ctx, httpMethod, apiMethod, params, target)
}

func (c *Client) call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error {
	if params == nil {
		params = make(url.Values)
	}

	params.Set("method", apiMethod)
	params.Set("api_key", c.apiKey)
	params.Set("format", "json")

	if c.sessionKey != "" {
		params.Set("sk", c.sessionKey)
	}

	if c.apiSecret != "" {
		params.Set("api_sig", c.sign(params))
	}

	var req *http.Request
	var err error

	if httpMethod == http.MethodPost {
		req, err = http.NewRequestWithContext(ctx, httpMethod, c.baseURL, strings.NewReader(params.Encode()))
		if err != nil {
			return fmt.Errorf("create request: %w", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		u, err := url.Parse(c.baseURL)
		if err != nil {
			return fmt.Errorf("parse base URL: %w", err)
		}
		u.RawQuery = params.Encode()
		req, err = http.NewRequestWithContext(ctx, httpMethod, u.String(), nil)
		if err != nil {
			return fmt.Errorf("create request: %w", err)
		}
	}

	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return fmt.Errorf("unexpected status code %d (failed to decode error: %w)", resp.StatusCode, err)
		}
		return &apiErr
	}

	if target != nil {
		if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}

func (c *Client) sign(params url.Values) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "format" || k == "callback" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString(params.Get(k))
	}
	sb.WriteString(c.apiSecret)

	h := md5.New()
	_, _ = io.WriteString(h, sb.String())
	return fmt.Sprintf("%x", h.Sum(nil))
}
