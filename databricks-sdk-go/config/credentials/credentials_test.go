package credentials

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"golang.org/x/oauth2"
)

type ctxKey struct{}

func TestOAuthCredentialsProvider_SetHeaders_passesRequestContext(t *testing.T) {
	var gotCtx context.Context
	ts := auth.TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
		gotCtx = ctx
		return &oauth2.Token{AccessToken: "token"}, nil
	})
	cp := NewOAuthCredentialsProviderFromTokenSource(ts)

	ctx := context.WithValue(context.Background(), ctxKey{}, "request")
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.test/api", nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := cp.SetHeaders(r); err != nil {
		t.Fatalf("SetHeaders() error = %v", err)
	}

	if got := gotCtx.Value(ctxKey{}); got != "request" {
		t.Errorf("token source context value = %v, want %q", got, "request")
	}
	if got := r.Header.Get("Authorization"); got != "Bearer token" {
		t.Errorf("Authorization header = %q, want %q", got, "Bearer token")
	}
}

func TestOAuthCredentialsProvider_SetHeaders_canceledRequestReturnsPromptly(t *testing.T) {
	// The refresh is held until the test ends. SetHeaders must return when the
	// request is canceled without waiting for it.
	release := make(chan struct{})
	defer close(release)
	ts := auth.NewCachedTokenSource(auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		<-release
		return &oauth2.Token{AccessToken: "token"}, nil
	}))
	cp := NewOAuthCredentialsProviderFromTokenSource(ts)

	ctx, cancel := context.WithCancel(context.Background())
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.test/api", nil)
	if err != nil {
		t.Fatal(err)
	}
	cancel()

	errc := make(chan error, 1)
	go func() {
		errc <- cp.SetHeaders(r)
	}()

	select {
	case err := <-errc:
		if !errors.Is(err, context.Canceled) {
			t.Errorf("SetHeaders() error = %v, want context.Canceled", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("SetHeaders() did not return after the request context was canceled")
	}
}

func TestFromCredentials_passesRequestContext(t *testing.T) {
	var gotCtx context.Context
	c := auth.CredentialsFn(func(ctx context.Context) (map[string]string, error) {
		gotCtx = ctx
		return map[string]string{"Authorization": "Bearer token"}, nil
	})
	cp := FromCredentials(c)

	ctx := context.WithValue(context.Background(), ctxKey{}, "request")
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.test/api", nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := cp.SetHeaders(r); err != nil {
		t.Fatalf("SetHeaders() error = %v", err)
	}

	if got := gotCtx.Value(ctxKey{}); got != "request" {
		t.Errorf("AuthHeaders context value = %v, want %q", got, "request")
	}
}
