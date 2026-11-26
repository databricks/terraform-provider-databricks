package dataplane

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"golang.org/x/oauth2"
)

type mockClient func(context.Context, string, *oauth2.Token) (*oauth2.Token, error)

func (m mockClient) GetOAuthToken(ctx context.Context, authDetails string, t *oauth2.Token) (*oauth2.Token, error) {
	return m(ctx, authDetails, t)
}

func TestDataPlaneTokenSource_Token(t *testing.T) {
	testErr := fmt.Errorf("test error")
	testToken := &oauth2.Token{AccessToken: "access token"}

	testCases := []struct {
		desc      string
		apiClient OAuthClient
		cpts      auth.TokenSource
		wantToken *oauth2.Token
		wantErr   error
	}{
		{
			desc: "Failing control plane token source",
			cpts: auth.TokenSourceFn(func(context.Context) (*oauth2.Token, error) {
				return testToken, testErr
			}),
			wantErr: testErr,
		},
		{
			desc: "Failing oauth endpoint",
			cpts: auth.TokenSourceFn(func(context.Context) (*oauth2.Token, error) {
				return testToken, nil
			}),
			apiClient: mockClient(func(context.Context, string, *oauth2.Token) (*oauth2.Token, error) {
				return nil, testErr
			}),
			wantErr: testErr,
		},
		{
			desc: "Successful token retrieval",
			cpts: auth.TokenSourceFn(func(context.Context) (*oauth2.Token, error) {
				return &oauth2.Token{AccessToken: "control plane test token"}, nil
			}),
			apiClient: mockClient(func(context.Context, string, *oauth2.Token) (*oauth2.Token, error) {
				return testToken, nil
			}),
			wantToken: testToken,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			dpts := NewEndpointTokenSource(tc.apiClient, tc.cpts)

			gotToken, gotErr := dpts.Token(context.Background(), "endpoint", "authDetails")

			if gotErr != tc.wantErr {
				t.Errorf("Token(): got error %v, want %v", gotErr, tc.wantErr)
			}
			if gotToken != tc.wantToken {
				t.Errorf("Token(): got token %v, want %v", gotToken, tc.wantToken)
			}
		})
	}
}
