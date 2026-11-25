package config

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestAzureReuseTokenSource(t *testing.T) {
	mockSource := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return &oauth2.Token{
			Expiry: time.Now().Add(35 * time.Second),
		}, nil
	})

	// Assert the token is not valid if it expires in 35 seconds.
	adjustedSource := azureReuseTokenSource(nil, mockSource)
	token, err := adjustedSource.Token(context.Background())
	assert.NoError(t, err)
	assert.False(t, token.Valid())
}

func TestServiceToServiceVisitorWithFallback_BothSucceed(t *testing.T) {
	primary := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return &oauth2.Token{AccessToken: "primary-token"}, nil
	})
	secondary := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return &oauth2.Token{AccessToken: "secondary-token"}, nil
	})
	visitor := serviceToServiceVisitor(primary, secondary, "X-Secondary", true)

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)
	err = visitor(req)
	require.NoError(t, err)
	assert.Equal(t, "Bearer primary-token", req.Header.Get("Authorization"))
	assert.Equal(t, "secondary-token", req.Header.Get("X-Secondary"))
}

func TestServiceToServiceVisitorWithFallback_SecondaryFails_SkipsHeader(t *testing.T) {
	primary := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return &oauth2.Token{AccessToken: "primary-token"}, nil
	})
	secondary := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return nil, fmt.Errorf("secondary failed")
	})
	visitor := serviceToServiceVisitor(primary, secondary, "X-Secondary", true)

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)
	err = visitor(req)
	require.NoError(t, err)
	assert.Equal(t, "Bearer primary-token", req.Header.Get("Authorization"))
	assert.Empty(t, req.Header.Get("X-Secondary"))
}

func TestServiceToServiceVisitor_SecondaryFails_NotOptional_ReturnsError(t *testing.T) {
	primary := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return &oauth2.Token{AccessToken: "primary-token"}, nil
	})
	secondary := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return nil, fmt.Errorf("secondary failed")
	})
	visitor := serviceToServiceVisitor(primary, secondary, "X-Secondary", false)

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)
	err = visitor(req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cloud token")
}

func TestServiceToServiceVisitorWithFallback_PrimaryFails_ReturnsError(t *testing.T) {
	primary := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return nil, fmt.Errorf("primary failed")
	})
	secondary := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return &oauth2.Token{AccessToken: "secondary-token"}, nil
	})
	visitor := serviceToServiceVisitor(primary, secondary, "X-Secondary", true)

	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)
	err = visitor(req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "inner token")
}
