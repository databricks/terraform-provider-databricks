package common

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigureWithOAuthM2M(t *testing.T) {
	defer CleanupEnvironment()()
	cnt := []int{0}
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			if req.RequestURI ==
				"/oidc/.well-known/oauth-authorization-server" {
				_, err := rw.Write([]byte(
					`{"token_endpoint": "http://localhost/oauth/token"}`))
				assert.NoError(t, err)
				cnt[0]++
				return
			}
			assert.Fail(t, fmt.Sprintf("Received unexpected call: %s %s",
				req.Method, req.RequestURI))
		}))
	defer server.Close()

	c := &DatabricksClient{
		Host:         server.URL,
		ClientID:     "abc",
		ClientSecret: "bcd",
	}
	_, err := c.configureWithOAuthM2M(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, cnt[0])
}

func TestConfigureWithOAuthOIDCUnavailableSkips(t *testing.T) {
	defer CleanupEnvironment()()
	c := &DatabricksClient{
		Host:         "http://localhost:22",
		ClientID:     "abc",
		ClientSecret: "bcd",
	}
	_, err := c.configureWithOAuthM2M(context.Background())
	assert.NoError(t, err)
}
