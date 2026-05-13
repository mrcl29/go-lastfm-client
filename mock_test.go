package golastfmclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockClient implements the APIClient interface for testing.
type MockClient struct {
	T          *testing.T
	HTTPMethod string
	APIMethod  string
	Params     url.Values
	RespJSON   string
	StatusCode int
}

func (m *MockClient) Call(ctx context.Context, httpMethod string, apiMethod string, params url.Values, target interface{}) error {
	assert.Equal(m.T, m.HTTPMethod, httpMethod)
	assert.Equal(m.T, m.APIMethod, apiMethod)

	for k, v := range m.Params {
		assert.Equal(m.T, v, params[k])
	}

	if m.StatusCode != 0 && m.StatusCode != http.StatusOK {
		return &mockError{Message: "Error"}
	}

	if target == nil {
		return nil
	}

	return json.Unmarshal([]byte(m.RespJSON), target)
}

type mockError struct {
	Code    int
	Message string
}

func (e *mockError) Error() string {
	return e.Message
}
