package u2m

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"golang.org/x/oauth2"
)

func TestCallbackServer_HandlerWithIssuerBindsIssuerToResult(t *testing.T) {
	browserOpened := make(chan string, 1)
	cb := &callbackServer{
		ctx: context.Background(),
		browser: func(redirect string) error {
			browserOpened <- redirect
			return nil
		},
		renderErrCh: make(chan error),
		feedbackCh:  make(chan oauthResult, 2),
		tmpl:        template.Must(template.New("page").Parse("")),
	}

	const authCodeURL = "https://login.databricks.com/?destination_url=%2Foidc%2Fv1%2Fauthorize"
	legitimate := struct {
		code   string
		state  string
		issuer string
	}{
		code:   "legit-code",
		state:  "legit-state",
		issuer: "https://adb-123.azuredatabricks.net/oidc",
	}

	serveCallback := func(code, state, issuer string) {
		callbackURL := fmt.Sprintf("/?code=%s&state=%s&iss=%s",
			url.QueryEscape(code), url.QueryEscape(state), url.QueryEscape(issuer))
		req := httptest.NewRequest(http.MethodGet, callbackURL, nil)
		rec := httptest.NewRecorder()
		cb.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("ServeHTTP status = %d, want %d", rec.Code, http.StatusOK)
		}
	}

	// Queue two callbacks with different code/state/iss; the handler must
	// return the values from the first result and not interleave.
	serveCallback(legitimate.code, legitimate.state, legitimate.issuer)
	serveCallback("second-code", "second-state", "https://other.example/oidc")

	code, state, issuer, err := cb.handlerWithIssuer(authCodeURL)
	if err != nil {
		t.Fatalf("handlerWithIssuer(): %v", err)
	}
	if code != legitimate.code {
		t.Errorf("code = %q, want %q", code, legitimate.code)
	}
	if state != legitimate.state {
		t.Errorf("state = %q, want %q", state, legitimate.state)
	}
	if issuer != legitimate.issuer {
		t.Errorf("issuer = %q, want %q", issuer, legitimate.issuer)
	}

	select {
	case got := <-browserOpened:
		if got != authCodeURL {
			t.Errorf("browser opened %q, want %q", got, authCodeURL)
		}
	default:
		t.Fatal("browser was not opened")
	}
}

func TestCallbackServer_ExtractsIssuer(t *testing.T) {
	ctx := context.Background()

	browserOpened := make(chan string)
	browser := func(redirect string) error {
		browserOpened <- redirect
		return nil
	}
	tokenCache := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}

	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(tokenCache),
		WithBrowser(browser),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=__ACCESS__&refresh_token=__REFRESH__`,
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

	// Start a listener so we can create the callback server.
	err = p.startListener(ctx)
	if err != nil {
		t.Fatalf("startListener(): %v", err)
	}

	cb, err := p.newCallbackServer()
	if err != nil {
		t.Fatalf("newCallbackServer(): %v", err)
	}
	defer cb.Close()

	// Fire a callback with iss parameter.
	issuerURL := "https://adb-123.azuredatabricks.net/oidc"
	resp, err := http.Get(fmt.Sprintf("http://%s?code=xxx&state=yyy&iss=%s", p.redirectAddr, issuerURL))
	if err != nil {
		t.Fatalf("http.Get(): %v", err)
	}
	defer resp.Body.Close()

	res := <-cb.feedbackCh
	if got := res.Issuer; got != issuerURL {
		t.Fatalf("result Issuer: want %q, got %q", issuerURL, got)
	}
}

func TestCallbackServer_NoIssuer(t *testing.T) {
	ctx := context.Background()

	tokenCache := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}

	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(tokenCache),
		WithBrowser(func(string) error { return nil }),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	err = p.startListener(ctx)
	if err != nil {
		t.Fatalf("startListener(): %v", err)
	}

	cb, err := p.newCallbackServer()
	if err != nil {
		t.Fatalf("newCallbackServer(): %v", err)
	}
	defer cb.Close()

	// Fire a callback without iss parameter.
	resp, err := http.Get(fmt.Sprintf("http://%s?code=xxx&state=yyy", p.redirectAddr))
	if err != nil {
		t.Fatalf("http.Get(): %v", err)
	}
	defer resp.Body.Close()

	res := <-cb.feedbackCh
	if got := res.Issuer; got != "" {
		t.Fatalf("result Issuer: want %q, got %q", "", got)
	}
}

func TestCallbackServer_IssuerWithAccountPath(t *testing.T) {
	ctx := context.Background()

	tokenCache := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}

	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(tokenCache),
		WithBrowser(func(string) error { return nil }),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	err = p.startListener(ctx)
	if err != nil {
		t.Fatalf("startListener(): %v", err)
	}

	cb, err := p.newCallbackServer()
	if err != nil {
		t.Fatalf("newCallbackServer(): %v", err)
	}
	defer cb.Close()

	// Fire a callback with iss containing an account path.
	issuerURL := "https://nike.databricks.com/oidc/accounts/abc-123-def"
	resp, err := http.Get(fmt.Sprintf("http://%s?code=xxx&state=yyy&iss=%s", p.redirectAddr, issuerURL))
	if err != nil {
		t.Fatalf("http.Get(): %v", err)
	}
	defer resp.Body.Close()

	res := <-cb.feedbackCh
	if got := res.Issuer; got != issuerURL {
		t.Fatalf("result Issuer: want %q, got %q", issuerURL, got)
	}
}
