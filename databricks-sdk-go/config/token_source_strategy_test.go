package config

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"golang.org/x/oauth2"
)

// authHeader returns the Authorization header set by the given credentials
// provider. It returns an empty string if the provider is nil.
func authHeader(cp credentials.CredentialsProvider) string {
	if cp == nil {
		return ""
	}
	req := &http.Request{Header: http.Header{}}
	cp.SetHeaders(req)
	return req.Header.Get("Authorization")
}

func TestTokenSourceStrategy_Configure(t *testing.T) {
	testCases := []struct {
		desc       string
		ts         auth.TokenSource
		wantHeader string
		wantError  bool
	}{
		{
			desc: "token source return error",
			ts: auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
				return nil, errors.New("test error")
			}),
			wantError: true,
		},
		{
			desc: "token source return token",
			ts: auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
				return &oauth2.Token{
					AccessToken: "token-123",
				}, nil
			}),
			wantHeader: "Bearer token-123",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			start := &tokenSourceStrategy{
				name: "test-strategy",
				ts:   tc.ts,
			}

			cp, err := start.Configure(context.Background(), &Config{})
			gotHeader := authHeader(cp)

			if tc.wantError && err == nil {
				t.Errorf("Expected error, but got none")
			}
			if !tc.wantError && err != nil {
				t.Errorf("Expected no error, but got %q", err)
			}
			if gotHeader != tc.wantHeader {
				t.Errorf("Expected header %q, but got %q", tc.wantHeader, gotHeader)
			}
		})
	}
}
