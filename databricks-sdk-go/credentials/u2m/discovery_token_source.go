package u2m

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

// defaultLoginDatabricksHost is the production host used for discovery login
// when no override is configured via WithDiscoveryHost.
const defaultLoginDatabricksHost = "https://login.databricks.com"

// DeriveHostFromIssuer extracts the workspace host (scheme + host) from an
// issuer URL returned in the OAuth callback iss parameter.
//
// Examples:
//
//	"https://adb-xxx.azuredatabricks.net/oidc"           -> "https://adb-xxx.azuredatabricks.net"
//	"https://nike.databricks.com/oidc/accounts/xxx"      -> "https://nike.databricks.com"
func DeriveHostFromIssuer(issuer string) (string, error) {
	if issuer == "" {
		return "", fmt.Errorf("issuer must not be empty")
	}
	u, err := url.Parse(issuer)
	if err != nil {
		return "", fmt.Errorf("parsing issuer URL %q: %w", issuer, err)
	}
	// Allow http for localhost (consistent with validateHost in workspace_oauth_argument.go).
	local := u.Scheme == "http" && u.Hostname() == "127.0.0.1"
	if u.Scheme != "https" && !local {
		return "", fmt.Errorf("issuer must use https scheme: %q", issuer)
	}
	if u.Host == "" {
		return "", fmt.Errorf("issuer must have a non-empty host: %q", issuer)
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host), nil
}

// DeriveTokenEndpoint derives the token endpoint from an issuer URL by
// appending /v1/token to the issuer path.
//
// Example:
//
//	"https://adb-xxx.net/oidc" -> "https://adb-xxx.net/oidc/v1/token"
func DeriveTokenEndpoint(issuer string) string {
	return strings.TrimRight(issuer, "/") + "/v1/token"
}

// discoveryTargetAccount is the value of the `target` query parameter that
// tells login.databricks.com to land the user on the account selector instead
// of the workspace selector. Used when the caller has signalled (e.g. via
// WithDiscoveryAccountTarget) that they only want account-level access.
const discoveryTargetAccount = "ACCOUNT"

// BuildDiscoveryAuthorizeURL builds the login.databricks.com URL that initiates
// the discovery OAuth flow. The OIDC authorize path with all OAuth query params
// is URL-encoded as the destination_url parameter.
func BuildDiscoveryAuthorizeURL(redirectAddr, state string, pkce PKCEParams, scopes []string) string {
	return buildDiscoveryAuthorizeURL(defaultLoginDatabricksHost, redirectAddr, state, pkce, scopes, "", "")
}

// buildDiscoveryAuthorizeURL builds the discovery authorize URL against the
// given host. Trailing slashes on host are trimmed so the result is
// well-formed regardless of how an override is written. When target is
// non-empty it is set as the top-level `target` query parameter, which
// login.databricks.com uses to route the user to a specific selector page
// (e.g. "ACCOUNT" for the account selector).
func buildDiscoveryAuthorizeURL(host, redirectAddr, state string, pkce PKCEParams, scopes []string, target, groupID string) string {
	// Build the nested OIDC authorize path with query parameters.
	authParams := url.Values{}
	authParams.Set("client_id", appClientID)
	authParams.Set("redirect_uri", "http://"+redirectAddr)
	authParams.Set("response_type", "code")
	authParams.Set("scope", strings.Join(scopes, " "))
	authParams.Set("state", state)
	authParams.Set("code_challenge", pkce.Challenge)
	authParams.Set("code_challenge_method", pkce.ChallengeMethod)

	if groupID != "" {
		authParams.Set("assume_group", groupID)
	}

	destinationURL := "/oidc/v1/authorize?" + authParams.Encode()

	// Wrap the authorize path as the destination_url query parameter on the
	// discovery host.
	topParams := url.Values{}
	if target != "" {
		topParams.Set("target", target)
	}
	topParams.Set("destination_url", destinationURL)
	return strings.TrimRight(host, "/") + "/?" + topParams.Encode()
}

// PKCEParams holds the PKCE challenge parameters used to build the discovery
// authorize URL. This mirrors authhandler.PKCEParams but is used directly so
// the caller does not need to import authhandler.
type PKCEParams struct {
	Challenge       string
	ChallengeMethod string
	Verifier        string
}

// discoveryTokenSource handles the OAuth PKCE flow for login.databricks.com
// discovery login. Unlike the standard flow, the token endpoint is not known
// until the callback provides the iss parameter identifying the workspace.
type discoveryTokenSource struct {
	pa *PersistentAuth
	// host overrides defaultLoginDatabricksHost when non-empty.
	host string
	// target is the value of the top-level `target` query parameter on the
	// authorize URL. When non-empty (e.g. "ACCOUNT"), login.databricks.com
	// routes the user directly to the corresponding selector.
	target string
}

// challenge initiates the discovery OAuth flow through login.databricks.com.
// It builds a custom authorize URL, opens the browser, waits for the callback,
// derives the workspace host and token endpoint from the iss parameter, and
// exchanges the authorization code for tokens.
func (d *discoveryTokenSource) challenge() error {
	cb, err := d.pa.newCallbackServer()
	if err != nil {
		return fmt.Errorf("callback server: %w", err)
	}
	defer cb.Close()

	state, authPKCE, err := d.pa.stateAndPKCE()
	if err != nil {
		return fmt.Errorf("state and pkce: %w", err)
	}

	scopes := d.pa.resolveScopes()

	pkce := PKCEParams{
		Challenge:       authPKCE.Challenge,
		ChallengeMethod: authPKCE.ChallengeMethod,
		Verifier:        authPKCE.Verifier,
	}
	host := d.host
	if host == "" {
		host = defaultLoginDatabricksHost
	}
	authorizeURL := buildDiscoveryAuthorizeURL(host, d.pa.redirectAddr, state, pkce, scopes, d.target, d.pa.groupID)

	code, returnedState, issuer, err := cb.handlerWithIssuer(authorizeURL)
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}

	// Validate state matches what we generated before consuming callback data.
	if returnedState != state {
		return fmt.Errorf("state mismatch: expected %q, got %q", state, returnedState)
	}

	if issuer == "" {
		return fmt.Errorf("discovery login failed: callback did not include an issuer (iss) parameter")
	}

	// Derive host and token endpoint from the issuer.
	discoveredHost, err := DeriveHostFromIssuer(issuer)
	if err != nil {
		return fmt.Errorf("deriving host from issuer: %w", err)
	}
	tokenEndpoint := DeriveTokenEndpoint(issuer)

	// Exchange authorization code for tokens.
	cfg := &oauth2.Config{
		ClientID: appClientID,
		Endpoint: oauth2.Endpoint{
			TokenURL:  tokenEndpoint,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: "http://" + d.pa.redirectAddr,
	}
	ctx := d.pa.setOAuthContext(d.pa.ctx)
	token, err := cfg.Exchange(ctx, code,
		oauth2.SetAuthURLParam("code_verifier", pkce.Verifier))
	if err != nil {
		return fmt.Errorf("token exchange: %w", err)
	}

	discoveryArg, ok := d.pa.oAuthArgument.(DiscoveryOAuthArgument)
	if !ok {
		return fmt.Errorf("discovery login requires DiscoveryOAuthArgument, got %T", d.pa.oAuthArgument)
	}
	discoveryArg.SetDiscoveredHost(discoveredHost)

	if err := d.pa.cache.Store(d.pa.oAuthArgument.GetCacheKey(), token); err != nil {
		return fmt.Errorf("storing token: %w", err)
	}
	return nil
}
