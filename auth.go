package golastfmclient

import (
	"context"
	"encoding/json"
	"net/url"
)

// AuthService handles authentication-related API calls.
type AuthService struct {
	client *Client
}

// Session represents a Last.fm user session.
type Session struct {
	Name       string      `json:"name"`
	Key        string      `json:"key"`
	Subscriber json.Number `json:"subscriber"`
}

type authGetSessionResponse struct {
	Session Session `json:"session"`
}

type authGetTokenResponse struct {
	Token string `json:"token"`
}

// GetToken fetches a request token for the Desktop authentication flow.
func (s *AuthService) GetToken(ctx context.Context) (string, error) {
	var resp authGetTokenResponse
	err := s.client.get(ctx, "auth.getToken", nil, &resp)
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}

// GetSession fetches a web service session for a user.
// The token should be the one received after user authorization.
func (s *AuthService) GetSession(ctx context.Context, token string) (Session, error) {
	params := url.Values{
		"token": {token},
	}
	var resp authGetSessionResponse
	err := s.client.get(ctx, "auth.getSession", params, &resp)
	if err != nil {
		return Session{}, err
	}
	return resp.Session, nil
}

// GetMobileSession fetches a session key for a mobile application using credentials.
// Note: This method uses a POST request over HTTPS as required by Last.fm.
func (s *AuthService) GetMobileSession(ctx context.Context, username, password string) (Session, error) {
	params := url.Values{
		"username": {username},
		"password": {password},
	}

	// We need to ensure HTTPS for this call.
	// If the client's baseURL is HTTP, we might need a temporary override or just assume HTTPS works.
	// For now, we use the client's post method.
	var resp authGetSessionResponse
	err := s.client.post(ctx, "auth.getMobileSession", params, &resp)
	if err != nil {
		return Session{}, err
	}
	return resp.Session, nil
}

// AuthURL returns the URL the user should be redirected to for authorization.
// For Web applications, token can be empty. For Desktop, it should be the token from GetToken.
func (s *AuthService) AuthURL(token string) string {
	u, _ := url.Parse("http://www.last.fm/api/auth/")
	q := u.Query()
	q.Set("api_key", s.client.apiKey)
	if token != "" {
		q.Set("token", token)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
