package oidc

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
)

func TestGithubIDTokenSource(t *testing.T) {
	testCases := []struct {
		desc              string
		tokenRequestUrl   string
		tokenRequestToken string
		audience          string
		httpTransport     http.RoundTripper
		wantToken         *IDToken
		wantErrPrefix     *string
	}{
		{
			desc:              "missing request token url",
			tokenRequestToken: "token-1337",
			wantErrPrefix:     errPrefix("missing ActionsIDTokenRequestURL"),
		},
		{
			desc:            "missing request token token",
			tokenRequestUrl: "http://endpoint.com/test?version=1",
			wantErrPrefix:   errPrefix("missing ActionsIDTokenRequestToken"),
		},
		{
			desc:              "error getting token",
			tokenRequestToken: "token-1337",
			tokenRequestUrl:   "http://endpoint.com/test?version=1",
			httpTransport: fixtures.MappingTransport{
				"GET /test?version=1": {
					Status: http.StatusInternalServerError,
					ExpectedHeaders: map[string]string{
						"Authorization": "Bearer token-1337",
						"Accept":        "application/json",
					},
				},
			},
			wantErrPrefix: errPrefix("failed to request ID token from"),
		},
		{
			desc:              "success",
			tokenRequestToken: "token-1337",
			tokenRequestUrl:   "http://endpoint.com/test?version=1",
			httpTransport: fixtures.MappingTransport{
				"GET /test?version=1": {
					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Authorization": "Bearer token-1337",
						"Accept":        "application/json",
					},
					Response: `{"value": "id-token-42"}`,
				},
			},
			wantToken: &IDToken{
				Value: "id-token-42",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			cli := httpclient.NewApiClient(httpclient.ClientConfig{
				Transport: tc.httpTransport,
			})
			p := &githubIDTokenSource{
				actionsIDTokenRequestURL:   tc.tokenRequestUrl,
				actionsIDTokenRequestToken: tc.tokenRequestToken,
				refreshClient:              cli,
			}
			token, gotErr := p.IDToken(context.Background(), tc.audience)

			if tc.wantErrPrefix == nil && gotErr != nil {
				t.Errorf("Authenticate(): got error %q, want none", gotErr)
			}
			if tc.wantErrPrefix != nil && !hasPrefix(gotErr, *tc.wantErrPrefix) {
				t.Errorf("Authenticate(): got error %q, want error with prefix %q", gotErr, *tc.wantErrPrefix)
			}
			if diff := cmp.Diff(tc.wantToken, token); diff != "" {
				t.Errorf("Authenticate(): mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
