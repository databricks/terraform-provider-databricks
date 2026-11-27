package authconv

import (
	"fmt"
	"testing"

	"golang.org/x/oauth2"
)

type mockOauth2TokenSource func() (*oauth2.Token, error)

func (t mockOauth2TokenSource) Token() (*oauth2.Token, error) {
	return t()
}

func TestIndepotency(t *testing.T) {
	wantErr := fmt.Errorf("test error")
	wantToken := &oauth2.Token{AccessToken: "test token"}
	ts := mockOauth2TokenSource(func() (*oauth2.Token, error) {
		return wantToken, wantErr
	})

	gotToken, gotErr := OAuth2TokenSource(AuthTokenSource(ts)).Token()

	if gotErr != wantErr {
		t.Errorf("Token() = %v, want %v", gotErr, wantErr)
	}
	if gotToken != wantToken {
		t.Errorf("Token() = %v, want %v", gotToken, wantToken)
	}
}
