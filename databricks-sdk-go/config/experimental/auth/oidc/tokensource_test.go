package oidc

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
)

func errPrefix(s string) *string {
	return &s
}

func hasPrefix(err error, prefix string) bool {
	return strings.HasPrefix(err.Error(), prefix)
}

func TestDatabricksOidcTokenSource(t *testing.T) {
	testCases := []struct {
		desc                 string
		clientID             string
		accountID            string
		host                 string
		tokenAudience        string
		httpTransport        http.RoundTripper
		oidcEndpointProvider func(context.Context) (*u2m.OAuthAuthorizationServer, error)
		idToken              string
		wantAudience         string
		tokenProviderError   error
		wantToken            string
		wantErrPrefix        *string
	}{
		{
			desc:          "missing host",
			clientID:      "client-id",
			tokenAudience: "token-audience",
			wantErrPrefix: errPrefix("missing Host"),
		},
		{
			desc: "token provider error",

			clientID:      "client-id",
			host:          "http://host.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			wantAudience:       "token-audience",
			tokenProviderError: errors.New("error getting id token"),
			wantErrPrefix:      errPrefix("error getting id token"),
		},
		{
			desc:          "databricks workspace server error",
			clientID:      "client-id",
			host:          "http://host.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
				"POST /oidc/v1/token": {
					Status: http.StatusInternalServerError,
					ExpectedHeaders: map[string]string{
						"Content-Type": "application/x-www-form-urlencoded",
					},
				},
			},
			wantAudience:  "token-audience",
			idToken:       "id-token-42",
			wantErrPrefix: errPrefix("oauth2: cannot fetch token: Internal Server Error"),
		},
		{
			desc:          "invalid auth token",
			clientID:      "client-id",
			host:          "http://host.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
				"POST /oidc/v1/token": {
					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Content-Type": "application/x-www-form-urlencoded",
					},
					Response: map[string]string{
						"foo": "bar",
					},
				},
			},
			wantAudience:  "token-audience",
			idToken:       "id-token-42",
			wantErrPrefix: errPrefix("oauth2: server response missing access_token"),
		},
		{
			desc:          "success WIF workspace",
			clientID:      "client-id",
			host:          "http://host.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
				"POST /oidc/v1/token": {

					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Content-Type": "application/x-www-form-urlencoded",
					},
					ExpectedRequest: url.Values{
						"client_id":          {"client-id"},
						"scope":              {"all-apis"},
						"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
						"subject_token":      {"id-token-42"},
						"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
					},
					Response: map[string]string{
						"token_type":    "access-token",
						"access_token":  "test-auth-token",
						"refresh_token": "refresh",
						"expires_on":    "0",
					},
				},
			},
			wantAudience: "token-audience",
			idToken:      "id-token-42",
			wantToken:    "test-auth-token",
		},
		{
			desc:          "success WIF account",
			clientID:      "client-id",
			accountID:     "ac123",
			host:          "https://accounts.databricks.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
				"POST /oidc/v1/token": {
					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Content-Type": "application/x-www-form-urlencoded",
					},
					ExpectedRequest: url.Values{
						"client_id":          {"client-id"},
						"scope":              {"all-apis"},
						"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
						"subject_token":      {"id-token-42"},
						"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
					},
					Response: map[string]string{
						"token_type":    "access-token",
						"access_token":  "test-auth-token",
						"refresh_token": "refresh",
						"expires_on":    "0",
					},
				},
			},
			wantAudience: "token-audience",
			idToken:      "id-token-42",
			wantToken:    "test-auth-token",
		},
		{
			desc:      "default token audience account",
			clientID:  "client-id",
			accountID: "ac123",
			host:      "https://accounts.databricks.com",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
				"POST /oidc/v1/token": {
					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Content-Type": "application/x-www-form-urlencoded",
					},
					Response: map[string]string{
						"token_type":    "access-token",
						"access_token":  "test-auth-token",
						"refresh_token": "refresh",
						"expires_on":    "0",
					},
				},
			},
			wantAudience: "ac123",
			idToken:      "id-token-42",
			wantToken:    "test-auth-token",
		},
		{
			desc:     "default token audience workspace",
			clientID: "client-id",
			host:     "https://host.com",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
				"POST /oidc/v1/token": {
					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Content-Type": "application/x-www-form-urlencoded",
					},
					Response: map[string]string{
						"token_type":    "access-token",
						"access_token":  "test-auth-token",
						"refresh_token": "refresh",
						"expires_on":    "0",
					},
				},
			},
			wantAudience: "https://host.com/oidc/v1/token",
			idToken:      "id-token-42",
			wantToken:    "test-auth-token",
		},
		{
			desc:          "success account-wide",
			host:          "http://host.com",
			tokenAudience: "token-audience",
			oidcEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
				return &u2m.OAuthAuthorizationServer{
					TokenEndpoint: "https://host.com/oidc/v1/token",
				}, nil
			},
			httpTransport: fixtures.MappingTransport{
				"POST /oidc/v1/token": {

					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Content-Type": "application/x-www-form-urlencoded",
					},
					ExpectedRequest: url.Values{
						"scope":              {"all-apis"},
						"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
						"subject_token":      {"id-token-42"},
						"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
					},
					Response: map[string]string{
						"token_type":    "access-token",
						"access_token":  "test-auth-token",
						"refresh_token": "refresh",
						"expires_on":    "0",
					},
				},
			},
			wantAudience: "token-audience",
			idToken:      "id-token-42",
			wantToken:    "test-auth-token",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			var gotAudience string // set when IDTokenSource is called
			cfg := DatabricksOIDCTokenSourceConfig{
				ClientID:              tc.clientID,
				AccountID:             tc.accountID,
				Host:                  tc.host,
				TokenEndpointProvider: tc.oidcEndpointProvider,
				Audience:              tc.tokenAudience,
				IDTokenSource: IDTokenSourceFn(func(ctx context.Context, aud string) (*IDToken, error) {
					gotAudience = aud
					return &IDToken{Value: tc.idToken}, tc.tokenProviderError
				}),
			}

			ts := NewDatabricksOIDCTokenSource(cfg)
			if tc.httpTransport != nil {
				ts.(*databricksOIDCTokenSource).cfg.TokenEndpointProvider = func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
					return &u2m.OAuthAuthorizationServer{
						TokenEndpoint: "https://host.com/oidc/v1/token",
					}, nil
				}
			}

			ctx := context.Background()
			if tc.httpTransport != nil {
				ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
					Transport: tc.httpTransport,
				})
			}

			token, err := ts.Token(ctx)
			if tc.wantErrPrefix == nil && err != nil {
				t.Errorf("Token(ctx): got error %q, want none", err)
			}
			if tc.wantErrPrefix != nil && !hasPrefix(err, *tc.wantErrPrefix) {
				t.Errorf("Token(ctx): got error %q, want error with prefix %q", err, *tc.wantErrPrefix)
			}
			if tc.wantAudience != gotAudience {
				t.Errorf("mockTokenProvider: got audience %s, want %s", gotAudience, tc.wantAudience)
			}
			tokenValue := ""
			if token != nil {
				tokenValue = token.AccessToken
			}
			if diff := cmp.Diff(tc.wantToken, tokenValue); diff != "" {
				t.Errorf("Authenticate(): mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWIF_Scopes(t *testing.T) {
	const (
		testClientID    = "test-client-id"
		testIDToken     = "test-id-token"
		testAccessToken = "test-access-token"
		testTokenPath   = "/oidc/v1/token"
		testHost        = "https://host.com"
	)

	tests := []struct {
		name   string
		scopes []string
		want   string
	}{
		{
			name:   "nil scopes uses default",
			scopes: nil,
			want:   "all-apis",
		},
		{
			name:   "empty scopes uses default",
			scopes: []string{},
			want:   "all-apis",
		},
		{
			name:   "single scope",
			scopes: []string{"dashboards"},
			want:   "dashboards",
		},
		{
			name:   "multiple scopes",
			scopes: []string{"jobs", "files:read", "mlflow"},
			want:   "jobs files:read mlflow",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := DatabricksOIDCTokenSourceConfig{
				ClientID: testClientID,
				Host:     testHost,
				TokenEndpointProvider: func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
					return &u2m.OAuthAuthorizationServer{
						TokenEndpoint: testHost + testTokenPath,
					}, nil
				},
				Audience: "token-audience",
				IDTokenSource: IDTokenSourceFn(func(ctx context.Context, aud string) (*IDToken, error) {
					return &IDToken{Value: testIDToken}, nil
				}),
				scopes: tt.scopes,
			}

			ts := NewDatabricksOIDCTokenSource(cfg)

			// The scope assertion: verifies the token source sends the correct scope parameter.
			expectedRequest := url.Values{
				"client_id":          {testClientID},
				"scope":              {tt.want},
				"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
				"subject_token":      {testIDToken},
				"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
			}

			ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
				Transport: fixtures.MappingTransport{
					"POST " + testTokenPath: {
						Status: http.StatusOK,
						ExpectedHeaders: map[string]string{
							"Content-Type": "application/x-www-form-urlencoded",
						},
						ExpectedRequest: expectedRequest,
						Response: map[string]string{
							"token_type":   "Bearer",
							"access_token": testAccessToken,
						},
					},
				},
			})

			token, err := ts.Token(ctx)
			if err != nil {
				t.Fatalf("Token(ctx): got error %q, want none", err)
			}
			if token.AccessToken != testAccessToken {
				t.Errorf("Token(ctx): got access token %q, want %q", token.AccessToken, testAccessToken)
			}
		})
	}
}

// TestDatabricksOIDCTokenSource_GroupRoleForm verifies that workspace, account,
// and unified token exchanges send the requested group to the selected token
// endpoint, including subsequent exchanges.
func TestDatabricksOIDCTokenSource_GroupRoleForm(t *testing.T) {
	testCases := []struct {
		name          string
		host          string
		accountID     string
		tokenEndpoint string
	}{
		{
			name:          "workspace",
			host:          "https://workspace.cloud.databricks.com",
			tokenEndpoint: "https://workspace.cloud.databricks.com/oidc/v1/token",
		},
		{
			name:          "account",
			host:          "https://accounts.cloud.databricks.com",
			accountID:     "account-123",
			tokenEndpoint: "https://accounts.cloud.databricks.com/oidc/accounts/account-123/v1/token",
		},
		{
			name:          "unified",
			host:          "https://unified.cloud.databricks.com",
			accountID:     "account-123",
			tokenEndpoint: "https://unified.cloud.databricks.com/oidc/accounts/account-123/v1/token",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			requestCount := 0
			transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
				requestCount++

				if got := req.URL.String(); got != testCase.tokenEndpoint {
					t.Errorf("Token request endpoint = %q, want %q", got, testCase.tokenEndpoint)
				}
				if err := req.ParseForm(); err != nil {
					t.Fatalf("Token request body could not be parsed: %v", err)
				}
				if got := req.Form.Get("assume_group"); got != "group-123" {
					t.Errorf("assume_group = %q, want %q", got, "group-123")
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     "200 OK",
					Header:     http.Header{"Content-Type": {"application/json"}},
					Body:       io.NopCloser(strings.NewReader(`{"access_token":"role-token","token_type":"Bearer"}`)),
					Request:    req,
				}, nil
			})

			ts := NewDatabricksOIDCTokenSource(DatabricksOIDCTokenSourceConfig{
				ClientID:  "client-id",
				AccountID: testCase.accountID,
				Host:      testCase.host,
				GroupID:   "group-123",
				TokenEndpointProvider: func(context.Context) (*u2m.OAuthAuthorizationServer, error) {
					return &u2m.OAuthAuthorizationServer{TokenEndpoint: testCase.tokenEndpoint}, nil
				},
				IDTokenSource: IDTokenSourceFn(func(context.Context, string) (*IDToken, error) {
					return &IDToken{Value: "id-token"}, nil
				}),
			})

			ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Transport: transport})

			for exchange := range 2 {
				if _, err := ts.Token(ctx); err != nil {
					t.Fatalf("Token exchange %d failed: %v", exchange+1, err)
				}
			}

			if requestCount != 2 {
				t.Errorf("Token endpoint was called %d times, want 2 calls", requestCount)
			}
		})
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

// RoundTrip calls the function used as an HTTP transport by the test.
func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
