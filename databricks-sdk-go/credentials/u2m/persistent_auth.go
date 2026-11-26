package u2m

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	cache "github.com/databricks/databricks-sdk-go/credentials/u2m/cache"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/authhandler"
)

const (
	// appClientId is the default client ID used by the SDK for U2M OAuth.
	appClientID = "databricks-cli"

	// defaultPort is the default port for the OAuth2 callback server. If the
	// port is already in use, the next port is tried (8021, 8022, etc.).
	defaultPort = 8020

	// maxPortFallback is the maximum port to try when using the fallback
	// mechanism.
	maxPortFallback = 8040

	// listenerTimeout is the maximum duration spent trying to acquire a
	// listener (including port selection).
	listenerTimeout = 45 * time.Second

	// tokenRefreshBuffer is the duration before token expiry at which the
	// token is proactively refreshed. This prevents callers from receiving
	// near-expired tokens that may expire before the next request.
	tokenRefreshBuffer = 5 * time.Minute

	// Cache update recovery checks immediately and then waits for 25, 50, 100,
	// and 200 milliseconds. This gives a concurrent cache writer time to finish
	// while bounding the delay for persistent storage failures to 375 milliseconds.
	cacheUpdateRecoveryAttempts = 5

	cacheUpdateRecoveryInitialDelay = 25 * time.Millisecond
	cacheUpdateRecoveryDelayFactor  = 2

	// Concurrent refreshes can finish at slightly different times, so their
	// expiration times need not be identical.
	cacheUpdateRecoveryExpiryDelta = time.Minute
)

var (
	// Internal errors used for testing.
	errListenerTimeout = errors.New("failed to listen on any port: timeout")
	errNoPortAvailable = errors.New("no port available to listen on")
)

// PersistentAuth is an OAuth manager that handles the U2M OAuth flow. Tokens
// are stored in and looked up from the provided cache. Tokens include the
// refresh token. On load, if the access token is expired or close to expiry,
// it is refreshed using the refresh token.
//
// The PersistentAuth is safe for concurrent use. The token cache is locked
// during token retrieval, refresh and storage.
type PersistentAuth struct {
	// cache is the token cache to store and lookup tokens.
	cache cache.TokenCache

	// client is the HTTP client to use for OAuth2 requests.
	client *http.Client

	// endpointSupplier is the HTTP endpointSupplier to use for OAuth2 requests.
	endpointSupplier OAuthEndpointSupplier

	// oAuthArgument defines the workspace or account to authenticate to and the
	// cache key for the token.
	oAuthArgument OAuthArgument

	// browser is the function to open a URL in the default browser.
	browser func(url string) error

	// ln is the listener for the OAuth2 callback server.
	ln net.Listener

	// ctx is the context to use for underlying operations. This is needed in
	// order to implement the oauth2.TokenSource interface.
	ctx context.Context

	// redirectAddr is the redirect address for OAuth2 callbacks. The value is
	// set to localhost:PORT by startListener which will dynamically assign a
	// random port. If a value is already provided, it will be used instead
	// (e.g. for testing).
	redirectAddr string

	// Optional port to use for the OAuth2 callback server. If set to 0, the
	// default port with fallback is used. This means that setting a port will
	// disable the fallback mechanism.
	port int

	// netListen is an optional function to listen on a TCP address. If not set,
	// it will use net.Listen by default. This is useful for testing.
	netListen func(network, address string) (net.Listener, error)

	// scopes is the list of OAuth scopes to request.
	scopes []string

	// disableOfflineAccess controls whether offline_access scope is requested.
	// When true, offline_access will NOT be automatically added to scopes,
	// meaning the token will not include a refresh token.
	disableOfflineAccess bool

	// discoveryMode enables the login.databricks.com discovery flow.
	// When true, Challenge() uses the discovery token source instead of
	// the standard authhandler flow.
	discoveryMode bool

	// discoveryHost overrides the default login.databricks.com host used by
	// the discovery flow. Empty means the production host.
	discoveryHost string

	// discoveryAccountTarget, when true, instructs the discovery flow to set
	// the top-level `target=ACCOUNT` query parameter on the authorize URL so
	// login.databricks.com lands the user on the account selector instead of
	// the workspace selector. Use for account-only logins.
	discoveryAccountTarget bool

	// groupID is the ID of the Databricks group whose role is assumed through
	// RBAC. The authorization code and refresh token retain this selection.
	groupID string
}

type PersistentAuthOption func(*PersistentAuth)

// WithTokenCache sets the token cache for the PersistentAuth.
func WithTokenCache(c cache.TokenCache) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.cache = c
	}
}

// WithHttpClient sets the HTTP client for the PersistentAuth.
func WithHttpClient(c *http.Client) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.client = c
	}
}

// WithOAuthEndpointSupplier sets the OAuth endpoint supplier for the
// PersistentAuth.
func WithOAuthEndpointSupplier(c OAuthEndpointSupplier) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.endpointSupplier = c
	}
}

// WithOAuthArgument sets the OAuthArgument for the PersistentAuth.
func WithOAuthArgument(arg OAuthArgument) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.oAuthArgument = arg
	}
}

// WithBrowser sets the browser function for the PersistentAuth.
func WithBrowser(b func(url string) error) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.browser = b
	}
}

// WithPort sets the port for the PersistentAuth.
func WithPort(port int) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.port = port
	}
}

// WithScopes sets the OAuth scopes for the PersistentAuth.
func WithScopes(scopes []string) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.scopes = scopes
	}
}

// WithDisableOfflineAccess controls whether offline_access scope is requested.
func WithDisableOfflineAccess(disable bool) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.disableOfflineAccess = disable
	}
}

// WithDiscoveryLogin enables the login.databricks.com discovery flow.
// When enabled, Challenge() routes through login.databricks.com instead
// of directly to a workspace OIDC endpoint.
//
// This option is only valid with [DiscoveryOAuthArgument], which is a
// bootstrap-only argument type for discovery login. Once the workspace host
// has been discovered, callers should construct the usual host-based
// OAuthArgument for future PersistentAuth instances.
func WithDiscoveryLogin() PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.discoveryMode = true
	}
}

// WithDiscoveryHost overrides the default https://login.databricks.com host
// used by the discovery login flow. Intended for testing and development
// against non-production environments; has no effect unless WithDiscoveryLogin
// is also set. If host has no scheme, https:// is prepended. Trailing slashes
// are trimmed.
func WithDiscoveryHost(host string) PersistentAuthOption {
	return func(a *PersistentAuth) {
		if host != "" && !strings.Contains(host, "://") {
			host = "https://" + host
		}
		a.discoveryHost = host
	}
}

// WithDiscoveryAccountTarget sets the top-level `target=ACCOUNT` query
// parameter on the discovery authorize URL so login.databricks.com lands the
// user on the account selector instead of the workspace selector. Use for
// account-only logins where workspace selection would be a wasted step.
//
// Has no effect unless WithDiscoveryLogin is also set.
func WithDiscoveryAccountTarget() PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.discoveryAccountTarget = true
	}
}

// WithGroupID configures the ID of the Databricks group whose role is assumed
// during a new U2M browser authorization through role-based access control
// (RBAC). The group ID is sent only on the authorization request; code
// exchange, refresh, and API requests do not resend it.
//
// Warning: Account authentication does not support group role assumption.
// As of August 2026, unified authentication rejects the group ID. The SDK still
// passes it to the OAuth server so unified role assumption may work in the future
// if server support is added.
func WithGroupID(groupID string) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.groupID = groupID
	}
}

// NewPersistentAuth creates a new PersistentAuth with the provided options.
func NewPersistentAuth(ctx context.Context, opts ...PersistentAuthOption) (*PersistentAuth, error) {
	p := &PersistentAuth{}
	for _, opt := range opts {
		opt(p)
	}
	// By default, PersistentAuth uses the default ApiClient to make HTTP
	// requests. Furthermore, if the endpointSupplier is not provided, it uses
	// this same client to fetch the OAuth endpoints. If the HTTP client is
	// provided but the endpointSupplier is not, we construct a default
	// ApiClient for use with BasicOAuthClient.
	apiClient := httpclient.NewApiClient(httpclient.ClientConfig{})
	if p.client == nil {
		p.client = &http.Client{
			Transport: apiClient,
			// 30 seconds matches the default timeout of the ApiClient
			Timeout: 30 * time.Second,
		}
	}
	if p.endpointSupplier == nil {
		p.endpointSupplier = &BasicOAuthEndpointSupplier{
			Client: apiClient,
		}
	}
	if p.cache == nil {
		p.cache = cache.NewInMemoryTokenCache()
	}
	if err := p.validateArg(); err != nil {
		return nil, err
	}
	if p.browser == nil {
		p.browser = browser.OpenURL
	}
	p.ctx = ctx
	return p, nil
}

// loadToken loads the cached OAuth2 token for the configured OAuthArgument
// using GetCacheKey(). The returned token may be expired; callers are
// responsible for deciding whether and how to refresh it.
func (a *PersistentAuth) loadToken() (*oauth2.Token, error) {
	t, err := a.cache.Lookup(a.oAuthArgument.GetCacheKey())
	if err != nil {
		return nil, fmt.Errorf("cache: %w", err)
	}
	return t, nil
}

// Token loads the OAuth2 token for the given OAuthArgument from the cache. If
// the token is expired or close to expiry, it is refreshed using the refresh
// token. When a proactive refresh (token still valid but near expiry) fails,
// the existing token is returned so the caller is not blocked.
func (a *PersistentAuth) Token() (*oauth2.Token, error) {
	t, err := a.loadToken()
	if err != nil {
		return nil, err
	}
	if needsRefresh(t) {
		if refreshedToken, err := a.refresh(t); err == nil {
			t = refreshedToken
		} else if !t.Valid() {
			return nil, fmt.Errorf("token refresh: %w", err)
		} else {
			logger.Debugf(a.ctx, "proactive token refresh failed, returning existing token: %v", err)
		}
	}
	t.RefreshToken = ""
	return t, nil
}

// ForceRefreshToken loads the OAuth2 token from the cache by GetCacheKey(),
// refreshes it unconditionally, and stores the refreshed token back under the
// same key. Unlike Token(), if the refresh fails the error is always returned
// -- the caller explicitly asked for a fresh token, so silently falling back
// to a stale one would be incorrect.
func (a *PersistentAuth) ForceRefreshToken() (*oauth2.Token, error) {
	t, err := a.loadToken()
	if err != nil {
		return nil, err
	}
	t, err = a.refresh(t)
	if err != nil {
		return nil, fmt.Errorf("forced token refresh: %w", err)
	}
	t.RefreshToken = ""
	return t, nil
}

// needsRefresh returns true when the token should be refreshed, either because
// it is no longer valid or because it will expire within the
// tokenRefreshBuffer window.
func needsRefresh(t *oauth2.Token) bool {
	if !t.Valid() {
		return true
	}
	return !t.Expiry.IsZero() && time.Until(t.Expiry) < tokenRefreshBuffer
}

// isFreshReplacement reports whether cached changed from old, is valid, and
// does not expire significantly sooner than candidate.
func isFreshReplacement(old, candidate, cached *oauth2.Token) bool {
	if cached.AccessToken == old.AccessToken || !cached.Valid() {
		return false
	}
	if cached.Expiry.IsZero() {
		return true
	}
	if candidate.Expiry.IsZero() {
		return false
	}
	return !cached.Expiry.Before(candidate.Expiry.Add(-cacheUpdateRecoveryExpiryDelta))
}

// recoverCacheUpdate checks whether a concurrent cache update completed.
// Retrying reads instead of writes avoids recreating the write race.
func (a *PersistentAuth) recoverCacheUpdate(old, candidate *oauth2.Token) *oauth2.Token {
	delay := cacheUpdateRecoveryInitialDelay
	for attempt := 0; attempt < cacheUpdateRecoveryAttempts; attempt++ {
		if attempt > 0 {
			timer := time.NewTimer(delay)
			select {
			case <-a.ctx.Done():
				timer.Stop()
				return nil
			case <-timer.C:
			}
			delay *= cacheUpdateRecoveryDelayFactor
		}

		cached, err := a.cache.Lookup(a.oAuthArgument.GetCacheKey())
		if err == nil && isFreshReplacement(old, candidate, cached) {
			return cached
		}
	}
	return nil
}

// refresh refreshes the token for the given OAuthArgument, storing the new
// token in the cache.
//
// TODO: This read-refresh-write sequence is not coordinated across processes.
// Because the CLI is stateless, two separate CLI invocations can load the same
// cached refresh token, both attempt a refresh, and race to update the cache.
// This should be fixed in a follow-up by adding cross-process coordination
// around refresh and cache writes.
func (a *PersistentAuth) refresh(oldToken *oauth2.Token) (*oauth2.Token, error) {
	// Fail fast with ErrMissingRefreshToken instead of letting the oauth2
	// library attempt to refresh and return a misleading error (e.g. "token
	// expired" when the real problem is that the cached token is not
	// refresh-capable).
	if oldToken.RefreshToken == "" {
		return nil, ErrMissingRefreshToken
	}
	cfg, err := a.oauth2Config()
	if err != nil {
		return nil, err
	}
	ctx := a.setOAuthContext(a.ctx)
	// Force the oauth2 library to refresh by ensuring the token appears
	// expired. PersistentAuth owns the refresh decision (including the
	// proactive buffer), so the oauth2 library should always perform the
	// refresh when asked.
	expired := *oldToken
	expired.Expiry = time.Now().Add(-time.Minute)
	t, err := cfg.TokenSource(ctx, &expired).Token()
	if err != nil {
		// The default RoundTripper of our httpclient.ApiClient returns an error
		// if the response status code is not 2xx. This isn't compliant with the
		// RoundTripper interface, so this error isn't handled by the oauth2
		// library. We need to handle it here.
		var internalHttpError *httpclient.HttpError
		if errors.As(err, &internalHttpError) {
			// error fields
			// https://datatracker.ietf.org/doc/html/rfc6749#section-5.2
			var errResponse struct {
				Error            string `json:"error"`
				ErrorDescription string `json:"error_description"`
			}
			if unmarshalErr := json.Unmarshal([]byte(internalHttpError.Message), &errResponse); unmarshalErr != nil {
				return nil, fmt.Errorf("unmarshal: %w", unmarshalErr)
			}
			// Invalid refresh tokens get their own error type so they can be
			// better presented to users.
			if errResponse.ErrorDescription == "Refresh token is invalid" {
				return nil, &InvalidRefreshTokenError{err}
			}
			return nil, fmt.Errorf("%s (error code: %s)", errResponse.ErrorDescription, errResponse.Error)
		}

		// Handle responses from well-behaved *http.Client implementations.
		var httpErr *oauth2.RetrieveError
		if errors.As(err, &httpErr) {
			// Invalid refresh tokens get their own error type so they can be
			// better presented to users.
			if httpErr.ErrorDescription == "Refresh token is invalid" {
				return nil, &InvalidRefreshTokenError{err}
			}
			return nil, fmt.Errorf("%s (error code: %s)", httpErr.ErrorDescription, httpErr.ErrorCode)
		}
		return nil, err
	}
	err = a.cache.Store(a.oAuthArgument.GetCacheKey(), t)
	if err != nil {
		if cached := a.recoverCacheUpdate(oldToken, t); cached != nil {
			return cached, nil
		}
		return nil, fmt.Errorf("cache update: %w", err)
	}
	return t, nil
}

// Challenge initiates the OAuth2 login flow for the given OAuthArgument. The
// OAuth2 flow is started by opening the browser to the OAuth2 authorization
// URL. The user is redirected to the callback server on appRedirectAddr. The
// callback server listens for the redirect from the identity provider and
// exchanges the authorization code for an access token.
func (a *PersistentAuth) Challenge() error {
	if a.discoveryMode {
		return a.discoveryChallenge()
	}
	err := a.startListener(a.ctx)
	if err != nil {
		return fmt.Errorf("starting listener: %w", err)
	}
	// The listener will be closed by the callback server automatically, but if
	// the callback server is not created, we need to close the listener manually.
	defer a.Close()

	cfg, err := a.authorizationConfig()
	if err != nil {
		return fmt.Errorf("building authorization config: %w", err)
	}
	cb, err := a.newCallbackServer()
	if err != nil {
		return fmt.Errorf("callback server: %w", err)
	}
	defer cb.Close()

	state, pkce, err := a.stateAndPKCE()
	if err != nil {
		return fmt.Errorf("state and pkce: %w", err)
	}
	// make OAuth2 library use our client
	ctx := a.setOAuthContext(a.ctx)
	ts := authhandler.TokenSourceWithPKCE(ctx, cfg, state, cb.Handler, pkce)
	t, err := ts.Token()
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	err = a.cache.Store(a.oAuthArgument.GetCacheKey(), t)
	if err != nil {
		return fmt.Errorf("store: %w", err)
	}
	return nil
}

// discoveryChallenge handles the login.databricks.com discovery flow.
// The listener must be started before the discovery token source is invoked
// because the challenge needs the redirect address to build the authorize URL.
func (a *PersistentAuth) discoveryChallenge() error {
	err := a.startListener(a.ctx)
	if err != nil {
		return fmt.Errorf("starting listener: %w", err)
	}
	defer a.Close()
	ds := &discoveryTokenSource{pa: a, host: a.discoveryHost}
	if a.discoveryAccountTarget {
		ds.target = discoveryTargetAccount
	}
	return ds.challenge()
}

// startListener starts a listener on appRedirectAddr, retrying if the address
// is already in use.
func (pa *PersistentAuth) startListener(ctx context.Context) error {
	if pa.port != 0 { // if port is set, use it
		return pa.startListenerWithPort(pa.port)
	}
	return pa.startListenerWithFallback(ctx)
}

// startListenerWithFallback starts a listener that will try to find a free
// port to listen on starting from the default port and incrementing by 1 until
// a free port is found.
func (pa *PersistentAuth) startListenerWithFallback(ctx context.Context) error {
	startTime := time.Now()
	for port := defaultPort; port <= maxPortFallback; port++ {
		if time.Since(startTime) > listenerTimeout {
			return errListenerTimeout
		}
		if err := pa.startListenerWithPort(port); err != nil {
			logger.Debugf(ctx, "failed to listen on %d: %v, retrying", port, err)
			continue
		}
		logger.Debugf(ctx, "OAuth callback server listening on %s", pa.redirectAddr)
		return nil
	}
	return errNoPortAvailable
}

func (pa *PersistentAuth) startListenerWithPort(port int) error {
	addr := fmt.Sprintf("localhost:%d", port)
	listener, err := pa.listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}
	pa.ln = listener
	pa.redirectAddr = addr
	return nil
}

func (pa *PersistentAuth) listen(network, addr string) (net.Listener, error) {
	if pa.netListen != nil {
		return pa.netListen(network, addr)
	} else {
		return net.Listen(network, addr)
	}
}

func (a *PersistentAuth) Close() error {
	if a.ln == nil {
		return nil
	}
	return a.ln.Close()
}

// validateArg ensures that the OAuthArgument is either a WorkspaceOAuthArgument,
// AccountOAuthArgument, UnifiedOAuthArgument, or a bootstrap-only
// DiscoveryOAuthArgument paired with WithDiscoveryLogin.
func (a *PersistentAuth) validateArg() error {
	if a.oAuthArgument == nil {
		return errors.New("missing OAuthArgument")
	}
	_, isWorkspaceArg := a.oAuthArgument.(WorkspaceOAuthArgument)
	_, isAccountArg := a.oAuthArgument.(AccountOAuthArgument)
	_, isUnifiedArg := a.oAuthArgument.(UnifiedOAuthArgument)
	_, isDiscoveryArg := a.oAuthArgument.(DiscoveryOAuthArgument)
	if !isWorkspaceArg && !isAccountArg && !isUnifiedArg && !isDiscoveryArg {
		return fmt.Errorf("unsupported OAuthArgument type: %T, must implement WorkspaceOAuthArgument, AccountOAuthArgument, UnifiedOAuthArgument, or DiscoveryOAuthArgument", a.oAuthArgument)
	}
	if isDiscoveryArg && !a.discoveryMode {
		return fmt.Errorf("discovery OAuthArgument %T requires WithDiscoveryLogin; after discovery, construct a WorkspaceOAuthArgument with the discovered host", a.oAuthArgument)
	}
	if a.discoveryMode && !isDiscoveryArg {
		return fmt.Errorf("discovery login requires DiscoveryOAuthArgument, got %T", a.oAuthArgument)
	}
	return nil
}

// resolveScopes returns the effective OAuth scopes for this PersistentAuth.
// It defaults to "all-apis" for backwards compatibility, and prepends
// "offline_access" unless disabled.
func (a *PersistentAuth) resolveScopes() []string {
	scopes := a.scopes
	if len(scopes) == 0 {
		scopes = []string{"all-apis"}
	}
	if !a.disableOfflineAccess {
		scopes = append([]string{"offline_access"}, scopes...)
	}
	return scopes
}

// oauth2Config returns the common OAuth2 configuration used for authorization
// and token refresh.
func (a *PersistentAuth) oauth2Config() (*oauth2.Config, error) {
	scopes := a.resolveScopes()

	var endpoints *OAuthAuthorizationServer
	var err error
	switch argg := a.oAuthArgument.(type) {
	case WorkspaceOAuthArgument:
		endpoints, err = a.endpointSupplier.GetWorkspaceOAuthEndpoints(a.ctx, argg.GetWorkspaceHost())
	case AccountOAuthArgument:
		endpoints, err = a.endpointSupplier.GetAccountOAuthEndpoints(
			a.ctx, argg.GetAccountHost(), argg.GetAccountId())
	case UnifiedOAuthArgument:
		endpoints, err = a.endpointSupplier.GetUnifiedOAuthEndpoints(a.ctx, argg.GetHost(), argg.GetAccountId())
	case DiscoveryOAuthArgument:
		return nil, fmt.Errorf("discovery OAuthArgument %T is only supported with WithDiscoveryLogin during Challenge; after discovery, construct a WorkspaceOAuthArgument with the discovered host", a.oAuthArgument)
	default:
		return nil, fmt.Errorf("unsupported OAuthArgument type: %T, must implement either WorkspaceOAuthArgument, AccountOAuthArgument or UnifiedOAuthArgument interface", a.oAuthArgument)
	}
	if err != nil {
		return nil, fmt.Errorf("fetching OAuth endpoints: %w", err)
	}
	return &oauth2.Config{
		ClientID: appClientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:   endpoints.AuthorizationEndpoint,
			TokenURL:  endpoints.TokenEndpoint,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: fmt.Sprintf("http://%s", a.redirectAddr),
		Scopes:      scopes,
	}, nil
}

// authorizationConfig extends the common OAuth2 configuration with parameters
// used only for the initial authorization request, such as group role assumption.
func (a *PersistentAuth) authorizationConfig() (*oauth2.Config, error) {
	config, err := a.oauth2Config()
	if err != nil {
		return nil, err
	}

	if a.groupID == "" {
		return config, nil
	}

	config.Endpoint.AuthURL, err = authorizationURLWithGroupID(config.Endpoint.AuthURL, a.groupID)
	if err != nil {
		return nil, fmt.Errorf("adding group ID to authorization endpoint: %w", err)
	}

	return config, nil
}

func authorizationURLWithGroupID(rawURL, groupID string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("parsing authorization endpoint %q: %w", rawURL, err)
	}

	query := parsedURL.Query()
	query.Set("assume_group", groupID)
	parsedURL.RawQuery = query.Encode()

	return parsedURL.String(), nil
}

func (a *PersistentAuth) stateAndPKCE() (string, *authhandler.PKCEParams, error) {
	verifier, err := a.randomString(64)
	if err != nil {
		return "", nil, fmt.Errorf("verifier: %w", err)
	}
	verifierSha256 := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(verifierSha256[:])
	state, err := a.randomString(16)
	if err != nil {
		return "", nil, fmt.Errorf("state: %w", err)
	}
	return state, &authhandler.PKCEParams{
		Challenge:       challenge,
		ChallengeMethod: "S256",
		Verifier:        verifier,
	}, nil
}

func (a *PersistentAuth) randomString(size int) (string, error) {
	raw := make([]byte, size)
	// ignore error as rand.Reader never returns an error
	_, err := rand.Read(raw)
	if err != nil {
		return "", fmt.Errorf("rand.Read: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(raw), nil
}

func (a *PersistentAuth) setOAuthContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, oauth2.HTTPClient, a.client)
}

var _ oauth2.TokenSource = (*PersistentAuth)(nil)
