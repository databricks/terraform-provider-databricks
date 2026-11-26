package u2m

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/credentials/u2m/cache"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"golang.org/x/oauth2"
)

type tokenCacheMock struct {
	store  func(key string, t *oauth2.Token) error
	lookup func(key string) (*oauth2.Token, error)
}

func (m *tokenCacheMock) Store(key string, t *oauth2.Token) error {
	return m.store(key, t)
}

func (m *tokenCacheMock) Lookup(key string) (*oauth2.Token, error) {
	return m.lookup(key)
}

func TestToken(t *testing.T) {
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != "https://abc/oidc/accounts/xyz" {
				t.Fatalf("lookup(): want key 'https://abc/oidc/accounts/xyz', got %s", key)
			}
			return &oauth2.Token{
				AccessToken: "bcd",
				Expiry:      time.Now().Add(1 * time.Hour),
			}, nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://abc", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(context.Background(), WithTokenCache(cache), WithOAuthArgument(arg))
	if err != nil {
		t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	tok, err := p.Token()

	if err != nil {
		t.Fatalf("p.Token(): want no error, got %v", err)
	}
	if tok.AccessToken != "bcd" {
		t.Errorf("p.Token(): want access token 'bcd', got %s", tok.AccessToken)
	}
	if tok.RefreshToken != "" {
		t.Errorf("p.Token(): want refresh token '', got %s", tok.RefreshToken)
	}
}

func TestToken_WithProfile(t *testing.T) {
	profileKey := "my-profile"
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != profileKey {
				t.Fatalf("lookup(): want key %q, got %q", profileKey, key)
			}
			return &oauth2.Token{
				AccessToken: "profile-token",
				Expiry:      time.Now().Add(1 * time.Hour),
			}, nil
		},
	}
	arg, err := NewProfileAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz", profileKey)
	if err != nil {
		t.Fatalf("NewProfileAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(context.Background(), WithTokenCache(cache), WithOAuthArgument(arg))
	if err != nil {
		t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	tok, err := p.Token()
	if err != nil {
		t.Fatalf("p.Token(): want no error, got %v", err)
	}
	if tok.AccessToken != "profile-token" {
		t.Errorf("p.Token(): want access token 'profile-token', got %s", tok.AccessToken)
	}
}

type MockOAuthEndpointSupplier struct{}

func (m MockOAuthEndpointSupplier) GetAccountOAuthEndpoints(ctx context.Context, accountHost string, accountId string) (*OAuthAuthorizationServer, error) {
	return &OAuthAuthorizationServer{
		AuthorizationEndpoint: fmt.Sprintf("%s/oidc/accounts/%s/v1/authorize", accountHost, accountId),
		TokenEndpoint:         fmt.Sprintf("%s/oidc/accounts/%s/v1/token", accountHost, accountId),
	}, nil
}

func (m MockOAuthEndpointSupplier) GetWorkspaceOAuthEndpoints(ctx context.Context, workspaceHost string) (*OAuthAuthorizationServer, error) {
	return &OAuthAuthorizationServer{
		AuthorizationEndpoint: fmt.Sprintf("%s/oidc/v1/authorize", workspaceHost),
		TokenEndpoint:         fmt.Sprintf("%s/oidc/v1/token", workspaceHost),
	}, nil
}

func (m MockOAuthEndpointSupplier) GetUnifiedOAuthEndpoints(ctx context.Context, host string, accountId string) (*OAuthAuthorizationServer, error) {
	return &OAuthAuthorizationServer{
		AuthorizationEndpoint: fmt.Sprintf("%s/oidc/accounts/%s/v1/authorize", host, accountId),
		TokenEndpoint:         fmt.Sprintf("%s/oidc/accounts/%s/v1/token", host, accountId),
	}, nil
}

func (m MockOAuthEndpointSupplier) GetEndpointsFromURL(_ context.Context, _ string) (*OAuthAuthorizationServer, error) {
	return nil, ErrOAuthNotSupported
}

func TestToken_RefreshesExpiredAccessToken(t *testing.T) {
	ctx := context.Background()
	expectedKey := "https://accounts.cloud.databricks.com/oidc/accounts/xyz"
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != expectedKey {
				t.Fatalf("lookup(): want key %s, got %s", expectedKey, key)
			}
			return &oauth2.Token{
				AccessToken:  "expired",
				RefreshToken: "cde",
				Expiry:       time.Now().Add(-1 * time.Minute),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			if key != expectedKey {
				t.Fatalf("store(): want key %s, got %s", expectedKey, key)
			}
			if tok.RefreshToken != "def" {
				t.Fatalf("store(): want refresh token 'def', got %s", tok.RefreshToken)
			}
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(cache),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=refreshed&refresh_token=def`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/x-www-form-urlencoded"},
					},
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Errorf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	tok, err := p.Token()
	if err != nil {
		t.Fatalf("p.Token(): want no error, got %v", err)
	}
	if tok.AccessToken != "refreshed" {
		t.Errorf("p.Token(): want access token 'refreshed', got %s", tok.AccessToken)
	}
	if tok.RefreshToken != "" {
		t.Errorf("p.Token(): want refresh token '', got %s", tok.RefreshToken)
	}
}

func TestToken_RefreshesTokenExpiringSoon(t *testing.T) {
	ctx := context.Background()
	expectedKey := "https://accounts.cloud.databricks.com/oidc/accounts/xyz"
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != expectedKey {
				t.Fatalf("lookup(): want key %s, got %s", expectedKey, key)
			}
			return &oauth2.Token{
				AccessToken:  "expiring-soon",
				RefreshToken: "cde",
				Expiry:       time.Now().Add(4 * time.Minute),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(c),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=refreshed&refresh_token=def`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/x-www-form-urlencoded"},
					},
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	tok, err := p.Token()
	if err != nil {
		t.Fatalf("p.Token(): want no error, got %v", err)
	}
	if tok.AccessToken != "refreshed" {
		t.Errorf("p.Token(): want access token 'refreshed', got %s", tok.AccessToken)
	}
}

func TestToken_ReturnsStillValidTokenWhenProactiveRefreshFails(t *testing.T) {
	ctx := context.Background()
	expectedKey := "https://accounts.cloud.databricks.com/oidc/accounts/xyz"
	transport := fixtures.SliceTransport{
		{
			Method:   "POST",
			Resource: "/oidc/accounts/xyz/v1/token",
			Response: `{"error": "temporarily_unavailable", "error_description": "temporarily unavailable"}`,
			Status:   503,
		},
	}
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != expectedKey {
				t.Fatalf("lookup(): want key %s, got %s", expectedKey, key)
			}
			return &oauth2.Token{
				AccessToken:  "still-valid",
				RefreshToken: "cde",
				Expiry:       time.Now().Add(4 * time.Minute),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			t.Fatalf("store(): unexpected call - refresh should not succeed")
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(c),
		WithHttpClient(&http.Client{Transport: transport}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	tok, err := p.Token()
	if err != nil {
		t.Fatalf("p.Token(): want no error, got %v", err)
	}
	if tok.AccessToken != "still-valid" {
		t.Errorf("p.Token(): want access token 'still-valid', got %s", tok.AccessToken)
	}
	if tok.RefreshToken != "" {
		t.Errorf("p.Token(): want refresh token '', got %s", tok.RefreshToken)
	}
	if transport[0].Method != "" {
		t.Errorf("refresh(): want proactive refresh attempt, but request was not sent")
	}
}

func TestToken_DoesNotRefreshTokenNotExpiringSoon(t *testing.T) {
	transport := fixtures.SliceTransport{}
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			return &oauth2.Token{
				AccessToken: "still-valid",
				Expiry:      time.Now().Add(6 * time.Minute),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			t.Fatalf("store(): unexpected call — token should not be refreshed")
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(c),
		WithHttpClient(&http.Client{Transport: transport}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	tok, err := p.Token()
	if err != nil {
		t.Fatalf("p.Token(): want no error, got %v", err)
	}
	if tok.AccessToken != "still-valid" {
		t.Errorf("p.Token(): want access token 'still-valid', got %s", tok.AccessToken)
	}
}

func TestToken_ZeroExpiryDoesNotTriggerRefresh(t *testing.T) {
	transport := fixtures.SliceTransport{}
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			return &oauth2.Token{
				AccessToken: "no-expiry",
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			t.Fatalf("store(): unexpected call — zero-expiry token should not be refreshed")
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(c),
		WithHttpClient(&http.Client{Transport: transport}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	tok, err := p.Token()
	if err != nil {
		t.Fatalf("p.Token(): want no error, got %v", err)
	}
	if tok.AccessToken != "no-expiry" {
		t.Errorf("p.Token(): want access token 'no-expiry', got %s", tok.AccessToken)
	}
}

func TestToken_ReturnsError(t *testing.T) {
	ctx := context.Background()
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != "https://accounts.cloud.databricks.com/oidc/accounts/xyz" {
				t.Fatalf("lookup(): want key 'https://accounts.cloud.databricks.com/oidc/accounts/xyz', got %s", key)
			}
			return &oauth2.Token{
				AccessToken:  "expired",
				RefreshToken: "cde",
				Expiry:       time.Now().Add(-1 * time.Minute),
			}, nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(cache),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `{"error": "invalid_grant", "error_description": "Invalid Client"}`,
					Status:   401,
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Errorf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()
	tok, err := p.Token()

	if tok != nil {
		t.Errorf("p.Token(): want nil, got %v", tok)
	}
	if !strings.Contains(err.Error(), "Invalid Client (error code: invalid_grant)") {
		t.Errorf("p.Token(): want error containing 'Invalid Client (error code: invalid_grant)', got %v", err)
	}
}

func TestToken_ReturnsInvalidRefreshTokenError(t *testing.T) {
	ctx := context.Background()
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != "https://accounts.cloud.databricks.com/oidc/accounts/xyz" {
				t.Fatalf("lookup(): want key 'https://accounts.cloud.databricks.com/oidc/accounts/xyz', got %s", key)
			}
			return &oauth2.Token{
				AccessToken:  "expired",
				RefreshToken: "cde",
				Expiry:       time.Now().Add(-1 * time.Minute),
			}, nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(cache),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `{"error": "invalid_grant", "error_description": "Refresh token is invalid"}`,
					Status:   401,
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Errorf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()
	tok, err := p.Token()
	if tok != nil {
		t.Fatalf("p.Token(): want nil, got %v", tok)
	}
	target := &InvalidRefreshTokenError{}
	if !errors.As(err, &target) {
		t.Fatalf("p.Token(): want error of type InvalidRefreshTokenError, got %v", err)
	}
}

func TestToken_ReturnsMissingRefreshTokenErrorWhenExpired(t *testing.T) {
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			return &oauth2.Token{
				AccessToken: "expired",
				Expiry:      time.Now().Add(-1 * time.Minute),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			t.Fatalf("store(): unexpected call — Token() should fail before writing when refresh token is missing")
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}
	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(c),
		WithHttpClient(&http.Client{Transport: fixtures.SliceTransport{}}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	tok, err := p.Token()
	if err == nil {
		t.Fatal("Token(): want error when refresh token is missing, got nil")
	}
	if tok != nil {
		t.Errorf("Token(): want nil token on error, got %v", tok)
	}
	if !errors.Is(err, ErrMissingRefreshToken) {
		t.Fatalf("Token(): want ErrMissingRefreshToken, got %v", err)
	}
	want := "token refresh: cached token has no refresh token"
	if got := err.Error(); got != want {
		t.Errorf("Token(): want error %q, got %q", want, got)
	}
}

func TestToken_ReturnsExistingTokenWhenNearExpiryAndNoRefreshToken(t *testing.T) {
	// Token is still valid but within the proactive refresh window (5 min).
	// Token() attempts refresh, refresh() returns ErrMissingRefreshToken.
	// Token() falls back to the existing token since it is still valid.
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			return &oauth2.Token{
				AccessToken: "near-expiry",
				Expiry:      time.Now().Add(3 * time.Minute), // within tokenRefreshBuffer
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			t.Fatalf("store(): unexpected call — Token() should return existing token without migrating")
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}
	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(c),
		WithHttpClient(&http.Client{Transport: fixtures.SliceTransport{}}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	tok, err := p.Token()
	if err != nil {
		t.Fatalf("Token(): want no error when falling back to valid token, got %v", err)
	}
	if tok == nil {
		t.Fatal("Token(): want token, got nil")
	}
	if tok.AccessToken != "near-expiry" {
		t.Errorf("Token(): want access token 'near-expiry', got %s", tok.AccessToken)
	}
}

func TestForceRefreshToken_RefreshesValidToken(t *testing.T) {
	refreshCalled := false
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			return &oauth2.Token{
				AccessToken:  "still-valid",
				RefreshToken: "refresh-me",
				Expiry:       time.Now().Add(1 * time.Hour),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			refreshCalled = true
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}
	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(c),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=force-refreshed&refresh_token=new-refresh`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/x-www-form-urlencoded"},
					},
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	tok, err := p.ForceRefreshToken()
	if err != nil {
		t.Fatalf("ForceRefreshToken(): want no error, got %v", err)
	}
	if tok.AccessToken != "force-refreshed" {
		t.Errorf("ForceRefreshToken(): want access token 'force-refreshed', got %s", tok.AccessToken)
	}
	if tok.RefreshToken != "" {
		t.Errorf("ForceRefreshToken(): want refresh token redacted, got %s", tok.RefreshToken)
	}
	if !refreshCalled {
		t.Error("ForceRefreshToken(): want refresh to be called for valid token, but it was not")
	}
}

func TestForceRefreshToken_RecoversConcurrentCacheUpdate(t *testing.T) {
	now := time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC)
	old := &oauth2.Token{
		AccessToken:  "old-access",
		RefreshToken: "old-refresh",
		Expiry:       now.Add(time.Hour),
	}
	winner := &oauth2.Token{
		AccessToken:  "winner-access",
		RefreshToken: "winner-refresh",
		Expiry:       now.Add(time.Hour - 30*time.Second),
	}
	lookupCalls := 0
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			lookupCalls++
			if lookupCalls <= 2 {
				return old, nil
			}
			return winner, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			if tok.AccessToken != "candidate-access" {
				t.Fatalf("store(): want candidate access token, got %q", tok.AccessToken)
			}
			return errors.New("concurrent cache update")
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}
	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(c),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=candidate-access&refresh_token=candidate-refresh&expires_in=3600`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/x-www-form-urlencoded"},
					},
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	tok, err := p.ForceRefreshToken()
	if err != nil {
		t.Fatalf("ForceRefreshToken(): want no error, got %v", err)
	}
	if tok.AccessToken != winner.AccessToken {
		t.Errorf("ForceRefreshToken(): want winner access token %q, got %q", winner.AccessToken, tok.AccessToken)
	}
	if tok.RefreshToken != "" {
		t.Errorf("ForceRefreshToken(): want refresh token redacted, got %q", tok.RefreshToken)
	}
	if lookupCalls != 3 {
		t.Errorf("Lookup(): want 3 calls, got %d", lookupCalls)
	}
}

func TestIsFreshReplacement(t *testing.T) {
	now := time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC)
	old := &oauth2.Token{AccessToken: "old", Expiry: now.Add(time.Hour)}
	candidate := &oauth2.Token{AccessToken: "candidate", Expiry: now.Add(time.Hour)}

	tests := []struct {
		name      string
		candidate *oauth2.Token
		cached    *oauth2.Token
		want      bool
	}{
		{
			name:      "same token",
			candidate: candidate,
			cached:    &oauth2.Token{AccessToken: "old", Expiry: now.Add(time.Hour)},
			want:      false,
		},
		{
			name:      "expired replacement",
			candidate: candidate,
			cached:    &oauth2.Token{AccessToken: "winner", Expiry: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)},
			want:      false,
		},
		{
			name:      "replacement expires too soon",
			candidate: candidate,
			cached:    &oauth2.Token{AccessToken: "winner", Expiry: candidate.Expiry.Add(-time.Minute - time.Second)},
			want:      false,
		},
		{
			name:      "replacement expiry within tolerance",
			candidate: candidate,
			cached:    &oauth2.Token{AccessToken: "winner", Expiry: candidate.Expiry.Add(-time.Minute)},
			want:      true,
		},
		{
			name:      "replacement expires later",
			candidate: candidate,
			cached:    &oauth2.Token{AccessToken: "winner", Expiry: candidate.Expiry.Add(time.Minute)},
			want:      true,
		},
		{
			name:      "expiring candidate and non-expiring replacement",
			candidate: candidate,
			cached:    &oauth2.Token{AccessToken: "winner"},
			want:      true,
		},
		{
			name:      "non-expiring candidate and replacement inside refresh buffer",
			candidate: &oauth2.Token{AccessToken: "candidate"},
			cached:    &oauth2.Token{AccessToken: "winner", Expiry: now.Add(time.Minute)},
			want:      false,
		},
		{
			name:      "non-expiring candidate and replacement",
			candidate: &oauth2.Token{AccessToken: "candidate"},
			cached:    &oauth2.Token{AccessToken: "winner"},
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFreshReplacement(old, tt.candidate, tt.cached); got != tt.want {
				t.Errorf("isFreshReplacement(): want %t, got %t", tt.want, got)
			}
		})
	}
}

func TestForceRefreshToken_WithInMemoryCachePreservesCachedRefreshToken(t *testing.T) {
	tokenCache := cache.NewInMemoryTokenCache()
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}
	if err := tokenCache.Store(arg.GetCacheKey(), &oauth2.Token{
		AccessToken:  "still-valid",
		RefreshToken: "refresh-me",
		Expiry:       time.Now().Add(1 * time.Hour),
	}); err != nil {
		t.Fatalf("Store(): %v", err)
	}

	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(tokenCache),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=force-refreshed&refresh_token=new-refresh`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/x-www-form-urlencoded"},
					},
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	tok, err := p.ForceRefreshToken()
	if err != nil {
		t.Fatalf("ForceRefreshToken(): want no error, got %v", err)
	}
	if tok.RefreshToken != "" {
		t.Fatalf("ForceRefreshToken(): want refresh token redacted, got %q", tok.RefreshToken)
	}

	cached, err := tokenCache.Lookup(arg.GetCacheKey())
	if err != nil {
		t.Fatalf("Lookup(): want cached token, got %v", err)
	}
	if cached.RefreshToken != "new-refresh" {
		t.Fatalf("Lookup(): want cached refresh token %q, got %q", "new-refresh", cached.RefreshToken)
	}
}

func TestForceRefreshToken_FailsWithoutRefreshToken(t *testing.T) {
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			return &oauth2.Token{
				AccessToken: "still-valid",
				Expiry:      time.Now().Add(1 * time.Hour),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			t.Fatalf("store(): unexpected call — ForceRefreshToken() should fail before writing when refresh token is missing")
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}
	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(c),
		WithHttpClient(&http.Client{Transport: fixtures.SliceTransport{}}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	tok, err := p.ForceRefreshToken()
	if err == nil {
		t.Fatal("ForceRefreshToken(): want error when refresh token is missing, got nil")
	}
	if tok != nil {
		t.Errorf("ForceRefreshToken(): want nil token on error, got %v", tok)
	}
	if !errors.Is(err, ErrMissingRefreshToken) {
		t.Fatalf("ForceRefreshToken(): want ErrMissingRefreshToken, got %v", err)
	}
	want := "forced token refresh: cached token has no refresh token"
	if got := err.Error(); got != want {
		t.Errorf("ForceRefreshToken(): want error %q, got %q", want, got)
	}
}

func TestForceRefreshToken_FailsOnRefreshError(t *testing.T) {
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			return &oauth2.Token{
				AccessToken:  "still-valid",
				RefreshToken: "bad-refresh",
				Expiry:       time.Now().Add(1 * time.Hour),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			t.Fatalf("store(): unexpected call — refresh should not succeed")
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}
	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(c),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `{"error": "temporarily_unavailable", "error_description": "temporarily unavailable"}`,
					Status:   503,
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	tok, err := p.ForceRefreshToken()
	if err == nil {
		t.Fatal("ForceRefreshToken(): want error when refresh fails, got nil")
	}
	if tok != nil {
		t.Errorf("ForceRefreshToken(): want nil token on error, got %v", tok)
	}
	want := "forced token refresh: temporarily unavailable (error code: temporarily_unavailable)"
	if got := err.Error(); got != want {
		t.Errorf("ForceRefreshToken(): want error %q, got %q", want, got)
	}
}

func TestForceRefreshToken_InvalidRefreshTokenError(t *testing.T) {
	c := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			return &oauth2.Token{
				AccessToken:  "expired",
				RefreshToken: "bad-refresh",
				Expiry:       time.Now().Add(-1 * time.Minute),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			t.Fatalf("store(): unexpected call — refresh should not succeed")
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}
	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(c),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `{"error": "invalid_grant", "error_description": "Refresh token is invalid"}`,
					Status:   401,
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	tok, err := p.ForceRefreshToken()
	if tok != nil {
		t.Fatalf("ForceRefreshToken(): want nil token, got %v", tok)
	}
	target := &InvalidRefreshTokenError{}
	if !errors.As(err, &target) {
		t.Fatalf("ForceRefreshToken(): want InvalidRefreshTokenError, got %v", err)
	}
}

func TestChallenge(t *testing.T) {
	ctx := context.Background()

	browserOpened := make(chan string)
	browser := func(redirect string) error {
		u, err := url.ParseRequestURI(redirect)
		if err != nil {
			return err
		}
		if u.Path != "/oidc/accounts/xyz/v1/authorize" {
			t.Fatalf("browser(): want path '/oidc/accounts/xyz/v1/authorize', got %s", u.Path)
		}
		// for now we're ignoring asserting the fields of the redirect
		query := u.Query()
		browserOpened <- query.Get("state")
		return nil
	}
	cache := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			if key != "https://accounts.cloud.databricks.com/oidc/accounts/xyz" {
				t.Fatalf("store(): want key 'https://accounts.cloud.databricks.com/oidc/accounts/xyz', got %s", key)
			}
			if tok.AccessToken != "__THAT__" {
				t.Fatalf("store(): want access token '__THAT__', got %s", tok.AccessToken)
			}
			if tok.RefreshToken != "__SOMETHING__" {
				t.Fatalf("store(): want refresh token '__SOMETHING__', got %s", tok.RefreshToken)
			}
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}

	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(cache),
		WithBrowser(browser),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=__THAT__&refresh_token=__SOMETHING__`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/x-www-form-urlencoded"},
					},
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Errorf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	errc := make(chan error)
	go func() {
		err := p.Challenge()
		errc <- err
		close(errc)
	}()

	state := <-browserOpened
	resp, err := http.Get(fmt.Sprintf("http://localhost:8020?code=__THIS__&state=%s", state))
	if err != nil {
		t.Fatalf("http.Get(): want no error, got %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("http.Get(): want status code 200, got %d", resp.StatusCode)
	}

	err = <-errc
	if err != nil {
		t.Fatalf("p.Challenge(): want no error, got %v", err)
	}
}

func TestAuthorizationConfig_GroupIDOnlyChangesAuthorizationEndpoint(t *testing.T) {
	workspaceArgument, err := NewBasicWorkspaceOAuthArgument("https://workspace.cloud.databricks.com")
	if err != nil {
		t.Fatalf("NewBasicWorkspaceOAuthArgument(): %v", err)
	}

	accountArgument, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "account-123")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}

	unifiedArgument, err := NewBasicUnifiedOAuthArgument("https://unified.cloud.databricks.com", "account-123")
	if err != nil {
		t.Fatalf("NewBasicUnifiedOAuthArgument(): %v", err)
	}

	testCases := []struct {
		name     string
		argument OAuthArgument
	}{
		{
			name:     "workspace",
			argument: workspaceArgument,
		},
		{
			name:     "account",
			argument: accountArgument,
		},
		{
			name:     "unified",
			argument: unifiedArgument,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			persistentAuth, err := NewPersistentAuth(
				context.Background(),
				WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
				WithOAuthArgument(testCase.argument),
				WithGroupID("group-A"),
			)
			if err != nil {
				t.Fatalf("NewPersistentAuth(): %v", err)
			}

			baseConfig, err := persistentAuth.oauth2Config()
			if err != nil {
				t.Fatalf("oauth2Config(): %v", err)
			}
			if strings.Contains(baseConfig.Endpoint.AuthURL, "assume_group") {
				t.Errorf("oauth2Config() AuthURL = %q, want no assume_group", baseConfig.Endpoint.AuthURL)
			}

			config, err := persistentAuth.authorizationConfig()
			if err != nil {
				t.Fatalf("authorizationConfig(): %v", err)
			}

			authorizationURL, err := url.Parse(config.Endpoint.AuthURL)
			if err != nil {
				t.Fatalf("url.Parse(AuthURL): %v", err)
			}
			if got := authorizationURL.Query()["assume_group"]; len(got) != 1 || got[0] != "group-A" {
				t.Errorf("authorization assume_group = %v, want [group-A]", got)
			}
			if strings.Contains(config.Endpoint.TokenURL, "assume_group") {
				t.Errorf("TokenURL = %q, want no assume_group", config.Endpoint.TokenURL)
			}
		})
	}
}

func TestChallenge_GroupIDStoresTokenUnderProfileKey(t *testing.T) {
	const profile = "prod"

	var storedKey string
	tokenCache := &tokenCacheMock{
		store: func(key string, _ *oauth2.Token) error {
			storedKey = key
			return nil
		},
	}

	browserOpened := make(chan string, 1)
	argument, err := NewProfileWorkspaceOAuthArgument("https://workspace.cloud.databricks.com", profile)
	if err != nil {
		t.Fatalf("NewProfileWorkspaceOAuthArgument(): %v", err)
	}

	persistentAuth, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(tokenCache),
		WithBrowser(func(authorizationURL string) error {
			browserOpened <- authorizationURL
			return nil
		}),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/v1/token",
					Response: `{"access_token":"access","refresh_token":"refresh","token_type":"Bearer","expires_in":3600}`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/json"},
					},
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(argument),
		WithGroupID("group-A"),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer persistentAuth.Close()

	completeChallenge(t, persistentAuth, browserOpened)

	if storedKey != profile {
		t.Errorf("Store() key = %q, want %q", storedKey, profile)
	}
}

func TestChallenge_ReturnsErrorOnFailure(t *testing.T) {
	ctx := context.Background()
	browserOpened := make(chan string)
	browser := func(redirect string) error {
		u, err := url.ParseRequestURI(redirect)
		if err != nil {
			return err
		}
		if u.Path != "/oidc/accounts/xyz/v1/authorize" {
			t.Fatalf("browser(): want path '/oidc/accounts/xyz/v1/authorize', got %s", u.Path)
		}
		// for now we're ignoring asserting the fields of the redirect
		query := u.Query()
		browserOpened <- query.Get("state")
		return nil
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}

	p, err := NewPersistentAuth(ctx, WithBrowser(browser), WithOAuthArgument(arg))
	if err != nil {
		t.Errorf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	errc := make(chan error)
	go func() {
		err := p.Challenge()
		errc <- err
		close(errc)
	}()

	<-browserOpened
	resp, err := http.Get("http://localhost:8020?error=access_denied&error_description=Policy%20evaluation%20failed%20for%20this%20request")
	if err != nil {
		t.Fatalf("http.Get(): want no error, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 400 {
		t.Fatalf("http.Get(): want status code 400, got %d", resp.StatusCode)
	}

	err = <-errc
	if err == nil {
		t.Fatalf("p.Challenge(): want error, got nil")
	}
	if !strings.Contains(err.Error(), "authorize: access_denied: Policy evaluation failed for this request") {
		t.Fatalf("p.Challenge(): want error containing 'authorize: access_denied: Policy evaluation failed for this request', got %v", err)
	}
}

func TestPersistentAuth_startListener_startFrom8020(t *testing.T) {
	pa := &PersistentAuth{}
	pa.netListen = func(_, address string) (net.Listener, error) {
		return nil, nil
	}

	gotErr := pa.startListener(context.Background())

	if gotErr != nil {
		t.Fatalf("pa.startListener(): want no error, got %v", gotErr)
	}
	if pa.redirectAddr != "localhost:8020" {
		t.Errorf("pa.redirectAddr should be localhost:8020, got %s", pa.redirectAddr)
	}
}

func TestPersistentAuth_startListener_incrementalFallBack(t *testing.T) {
	pa := &PersistentAuth{}
	pa.netListen = func(_, address string) (net.Listener, error) {
		if address == "localhost:8020" {
			return nil, fmt.Errorf("address already in use")
		}
		if address == "localhost:8021" {
			return nil, fmt.Errorf("address already in use")
		}
		return nil, nil
	}

	gotErr := pa.startListener(context.Background())

	if gotErr != nil {
		t.Fatalf("pa.startListener(): want no error, got %v", gotErr)
	}
	if pa.redirectAddr != "localhost:8022" {
		t.Errorf("pa.redirectAddr should be localhost:8022, got %s", pa.redirectAddr)
	}
}

func TestPersistentAuth_startListener_noAvailablePort(t *testing.T) {
	pa := &PersistentAuth{}
	pa.netListen = func(_, address string) (net.Listener, error) {
		return nil, fmt.Errorf("address already in use")
	}

	gotErr := pa.startListener(context.Background())

	if !errors.Is(gotErr, errNoPortAvailable) {
		t.Fatalf("pa.startListener(): want error %v, got %v", errNoPortAvailable, gotErr)
	}
}

func TestPersistentAuth_startListener_maxPortFallbackIncluded(t *testing.T) {
	maxAddress := fmt.Sprintf("localhost:%d", maxPortFallback)
	pa := &PersistentAuth{}
	pa.netListen = func(_, address string) (net.Listener, error) {
		if address == maxAddress {
			return nil, nil
		}
		return nil, fmt.Errorf("address already in use")
	}

	gotErr := pa.startListener(context.Background())

	if gotErr != nil {
		t.Fatalf("pa.startListener(): want no error, got %v", gotErr)
	}
	if pa.redirectAddr != maxAddress {
		t.Errorf("pa.redirectAddr should be %s, got %s", maxAddress, pa.redirectAddr)
	}
}

func TestPersistentAuth_startListener_explicitPort(t *testing.T) {
	explicitPort := 1337
	pa := &PersistentAuth{port: explicitPort}
	pa.netListen = func(_, address string) (net.Listener, error) {
		return nil, nil
	}

	gotErr := pa.startListener(context.Background())

	if gotErr != nil {
		t.Fatalf("pa.startListener(): want no error, got %v", gotErr)
	}
	if pa.redirectAddr != "localhost:1337" {
		t.Errorf("pa.redirectAddr should be localhost:1337, got %s", pa.redirectAddr)
	}
}

func TestPersistentAuth_startListener_explicitPortNoFallBack(t *testing.T) {
	testError := errors.New("test error")
	explicitPort := 1337
	pa := &PersistentAuth{port: explicitPort}
	pa.netListen = func(_, address string) (net.Listener, error) {
		if address == "localhost:1337" {
			return nil, testError
		}
		return nil, nil
	}

	gotErr := pa.startListener(context.Background())

	if !errors.Is(gotErr, testError) {
		t.Fatalf("pa.startListener(): want error %v, got %v", testError, gotErr)
	}
}

// TestU2M_ScopesAndOfflineAccess verifies that OAuth scopes are correctly configured
// and sent during the authorization flow, and that the disableOfflineAccess flag
// correctly controls whether offline_access is added to the scope.
func TestU2M_ScopesAndOfflineAccess(t *testing.T) {
	const (
		testWorkspaceHost = "https://workspace.cloud.databricks.com"
		testTokenEndpoint = "/oidc/v1/token"
		testCallbackURL   = "http://localhost:8020"
	)

	tests := []struct {
		name           string
		scopes         []string
		disableOffline bool
		want           string
	}{
		{
			name:           "nil scopes uses default with offline_access",
			scopes:         nil,
			disableOffline: false,
			want:           "offline_access all-apis",
		},
		{
			name:           "empty scopes uses default with offline_access",
			scopes:         []string{},
			disableOffline: false,
			want:           "offline_access all-apis",
		},
		{
			name:           "single scope with offline_access",
			scopes:         []string{"dashboards"},
			disableOffline: false,
			want:           "offline_access dashboards",
		},
		{
			name:           "multiple scopes with offline_access",
			scopes:         []string{"files", "jobs", "mlflow:read"},
			disableOffline: false,
			want:           "offline_access files jobs mlflow:read",
		},
		{
			name:           "disable offline_access",
			scopes:         []string{"files", "jobs"},
			disableOffline: true,
			want:           "files jobs",
		},
		{
			name:           "nil scopes with disable offline_access",
			scopes:         nil,
			disableOffline: true,
			want:           "all-apis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			var scopeReceived, stateReceived string
			browserCalled := make(chan struct{})
			defer close(browserCalled)
			browser := func(redirect string) error {
				u, err := url.ParseRequestURI(redirect)
				if err != nil {
					return err
				}
				query := u.Query()
				scopeReceived = query.Get("scope")
				stateReceived = query.Get("state")
				browserCalled <- struct{}{}
				return nil
			}

			cache := &tokenCacheMock{
				store: func(key string, tok *oauth2.Token) error {
					return nil
				},
			}

			arg, err := NewBasicWorkspaceOAuthArgument(testWorkspaceHost)
			if err != nil {
				t.Fatalf("NewBasicWorkspaceOAuthArgument(): want no error, got %v", err)
			}

			var tokenResponse string
			if tt.disableOffline {
				tokenResponse = `access_token=token`
			} else {
				tokenResponse = `access_token=token&refresh_token=refresh`
			}

			opts := []PersistentAuthOption{
				WithTokenCache(cache),
				WithBrowser(browser),
				WithHttpClient(&http.Client{
					Transport: fixtures.SliceTransport{
						{
							Method:   "POST",
							Resource: testTokenEndpoint,
							Response: tokenResponse,
							ResponseHeaders: map[string][]string{
								"Content-Type": {"application/x-www-form-urlencoded"},
							},
						},
					},
				}),
				WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
				WithOAuthArgument(arg),
				WithDisableOfflineAccess(tt.disableOffline),
				WithScopes(tt.scopes),
			}

			p, err := NewPersistentAuth(ctx, opts...)
			if err != nil {
				t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
			}
			defer p.Close()

			errc := make(chan error)
			defer close(errc)
			go func() {
				err := p.Challenge()
				errc <- err
			}()

			select {
			case <-browserCalled:
			case <-time.After(5 * time.Second):
				t.Fatal("timed out waiting for browser to be called")
			}

			if scopeReceived != tt.want {
				t.Errorf("scope: want %q, got %q", tt.want, scopeReceived)
			}

			resp, err := http.Get(fmt.Sprintf("%s?code=__CODE__&state=%s", testCallbackURL, stateReceived))
			if err != nil {
				t.Fatalf("http.Get(): want no error, got %v", err)
			}
			defer resp.Body.Close()

			select {
			case err = <-errc:
			case <-time.After(5 * time.Second):
				t.Fatal("timed out waiting for Challenge() to complete")
			}
			if err != nil {
				t.Fatalf("p.Challenge(): want no error, got %v", err)
			}
		})
	}
}

func TestChallenge_Discovery(t *testing.T) {
	// Mock token server that responds to POST /oidc/v1/token.
	tokenServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("token server: want POST, got %s", r.Method)
		}
		if r.URL.Path != "/oidc/v1/token" {
			t.Errorf("token server: want path /oidc/v1/token, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"access_token":"discovery-access-token","refresh_token":"discovery-refresh-token","token_type":"Bearer","expires_in":3600}`)
	}))
	defer tokenServer.Close()

	// The issuer is the mock server URL + /oidc.
	issuer := tokenServer.URL + "/oidc"

	browserOpened := make(chan string, 1)
	browserMock := func(u string) error {
		browserOpened <- u
		return nil
	}

	storedTokens := map[string]*oauth2.Token{}
	cacheMock := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			storedTokens[key] = tok
			return nil
		},
	}

	arg, err := NewBasicDiscoveryOAuthArgument("discovery-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): %v", err)
	}

	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(cacheMock),
		WithBrowser(browserMock),
		WithHttpClient(tokenServer.Client()),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
		WithDiscoveryLogin(),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	errc := make(chan error, 1)
	go func() {
		errc <- p.Challenge()
	}()

	// Wait for browser to be called and extract state from the authorize URL.
	var state string
	select {
	case authURL := <-browserOpened:
		u, err := url.Parse(authURL)
		if err != nil {
			t.Fatalf("parsing auth URL: %v", err)
		}
		destURL := u.Query().Get("destination_url")
		dest, err := url.Parse(destURL)
		if err != nil {
			t.Fatalf("parsing destination_url: %v", err)
		}
		state = dest.Query().Get("state")
		if state == "" {
			t.Fatal("state is empty in authorize URL")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for browser to be called")
	}

	// Fire the callback with code, state, and iss.
	callbackURL := fmt.Sprintf("http://%s?code=test-code&state=%s&iss=%s",
		p.redirectAddr, url.QueryEscape(state), url.QueryEscape(issuer))
	resp, err := http.Get(callbackURL)
	if err != nil {
		t.Fatalf("callback GET: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("callback: want status 200, got %d", resp.StatusCode)
	}

	// Wait for Challenge to complete.
	select {
	case err := <-errc:
		if err != nil {
			t.Fatalf("Challenge(): %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for Challenge to complete")
	}

	// Verify discovered host was set on the argument.
	expectedHost, err := DeriveHostFromIssuer(issuer)
	if err != nil {
		t.Fatalf("DeriveHostFromIssuer(%q): %v", issuer, err)
	}
	if arg.GetDiscoveredHost() != expectedHost {
		t.Errorf("discovered host = %q, want %q", arg.GetDiscoveredHost(), expectedHost)
	}
	if len(storedTokens) != 1 {
		t.Fatalf("store count: want 1 key (profile), got %d", len(storedTokens))
	}
	storedToken := storedTokens["discovery-profile"]
	if storedToken == nil {
		t.Fatalf("stored token for profile key is nil")
	}
	if storedToken.AccessToken != "discovery-access-token" {
		t.Errorf("access token = %q, want %q", storedToken.AccessToken, "discovery-access-token")
	}
	if storedToken.RefreshToken != "discovery-refresh-token" {
		t.Errorf("refresh token = %q, want %q", storedToken.RefreshToken, "discovery-refresh-token")
	}
}

// ----------------------------------------------------------------------------
// Test helper functions
// ----------------------------------------------------------------------------

// completeChallenge simulates the callback for a browser authorization URL and
// waits for the OAuth challenge to finish.
func completeChallenge(t *testing.T, persistentAuth *PersistentAuth, browserOpened <-chan string) {
	t.Helper()

	errorChannel := make(chan error, 1)
	go func() {
		errorChannel <- persistentAuth.Challenge()
	}()

	var authorizationURL string
	select {
	case authorizationURL = <-browserOpened:
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for browser authorization")
	}

	parsedAuthorizationURL, err := url.ParseRequestURI(authorizationURL)
	if err != nil {
		t.Fatalf("url.ParseRequestURI(): %v", err)
	}

	state := parsedAuthorizationURL.Query().Get("state")
	callbackClient := &http.Client{Timeout: 5 * time.Second}
	callbackURL := fmt.Sprintf("http://%s?code=code&state=%s", persistentAuth.redirectAddr, url.QueryEscape(state))
	response, err := callbackClient.Get(callbackURL)
	if err != nil {
		t.Fatalf("callback: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("callback status = %d, want %d", response.StatusCode, http.StatusOK)
	}

	select {
	case err := <-errorChannel:
		if err != nil {
			t.Fatalf("Challenge(): %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for Challenge() to complete")
	}
}
