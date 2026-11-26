package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/exp/slices"
	"golang.org/x/oauth2"
)

// CredentialsStrategy responsible for configuring static or refreshable
// authentication credentials for Databricks REST APIs
type CredentialsStrategy interface {
	// Name returns human-addressable name of this credentials provider strategy
	Name() string

	// Configure creates CredentialsProvider or returns nil if a given credentials
	// strategy are not configured. It returns an error if credentials are misconfigured.
	// Takes a context and a pointer to a Config instance, that holds auth mutex.
	Configure(context.Context, *Config) (credentials.CredentialsProvider, error)
}

type Loader interface {
	// Name is human-addressable representation of this config resolver
	Name() string
	Configure(*Config) error
}

// HostType represents the type of API the configured host supports.
type HostType string

const (
	// WorkspaceHost supports only workspace-level APIs.
	WorkspaceHost HostType = "WORKSPACE_HOST"
	// AccountHost supports only account-level APIs.
	AccountHost HostType = "ACCOUNT_HOST"
	// UnifiedHost supports both workspace-level and account-level APIs.
	UnifiedHost HostType = "UNIFIED_HOST"
	// HostTypeUnknown is the zero value for an unset/empty host type.
	HostTypeUnknown HostType = ""
)

// UnmarshalJSON automatically normalizes the host type string when parsing JSON.
func (h *HostType) UnmarshalJSON(data []byte) error {
	var rawString string
	if err := json.Unmarshal(data, &rawString); err != nil {
		return err
	}
	*h = normalizeHostType(rawString)
	return nil
}

func normalizeHostType(hostType string) HostType {
	switch strings.ToLower(hostType) {
	case "workspace":
		return WorkspaceHost
	case "account":
		return AccountHost
	case "unified":
		return UnifiedHost
	default:
		return HostTypeUnknown
	}
}

// ConfigType represents the type of API this config is valid for.
type ConfigType string

const (
	// WorkspaceConfig is valid for workspace-level API requests.
	WorkspaceConfig ConfigType = "WORKSPACE_CONFIG"
	// AccountConfig is valid for account-level API requests.
	AccountConfig ConfigType = "ACCOUNT_CONFIG"
	// InvalidConfig is returned when the config is not valid for either workspace-level or account-level APIs.
	InvalidConfig ConfigType = "INVALID_CONFIG"
)

var defaultScopes = []string{"all-apis"}

// DefaultCredentialStrategy is the default factory for creating a
// [CredentialsStrategy] when [Config.Credentials] is not set. It can be
// overridden to change the default authentication behavior for all Config
// instances.
//
// Note: This variable is not thread safe and is expected to be set during
// program initialization or before any concurrent use. Changing it at runtime
// in a concurrent context may lead to race conditions and undefined behavior.
var DefaultCredentialStrategyProvider func() CredentialsStrategy = func() CredentialsStrategy {
	return &DefaultCredentials{}
}

// Config represents configuration for Databricks Connectivity
type Config struct {
	// Credentials holds an instance of Credentials Strategy to authenticate with Databricks REST APIs.
	// If no credentials strategy is specified, `DefaultCredentials` are implicitly used.
	Credentials CredentialsStrategy

	// Databricks host (either of workspace endpoint or Accounts API endpoint)
	Host string `name:"host" env:"DATABRICKS_HOST"`

	ClusterID           string `name:"cluster_id" env:"DATABRICKS_CLUSTER_ID"`
	WarehouseID         string `name:"warehouse_id" env:"DATABRICKS_WAREHOUSE_ID"`
	ServerlessComputeID string `name:"serverless_compute_id" env:"DATABRICKS_SERVERLESS_COMPUTE_ID"`

	// URL of the metadata service that provides authentication credentials.
	MetadataServiceURL string `name:"metadata_service_url" env:"DATABRICKS_METADATA_SERVICE_URL" auth:"metadata-service,sensitive"`

	// Databricks Account ID for Accounts API. This field is used in dependencies.
	AccountID string `name:"account_id" env:"DATABRICKS_ACCOUNT_ID"`

	// Databricks Workspace ID for Workspace clients when working with unified hosts.
	// Accepts either a classic numeric workspace ID or a CPDR connection ID; the
	// server disambiguates.
	WorkspaceID string `name:"workspace_id" env:"DATABRICKS_WORKSPACE_ID"`

	// GroupID is the ID of the Databricks group whose role is assumed through
	// role-based access control (RBAC). A non-empty value requires a supported
	// Databricks OAuth credential strategy and fixes the role for this client.
	// Databricks CLI authentication does not support group role assumption.
	//
	// Warning: As of August 2026, group role assumption works only for workspace
	// authorization. Account and unified authentication are not supported. M2M
	// and WIF still pass the group ID to those OAuth servers so they can support
	// group role assumption in the future without requiring an SDK change.
	GroupID string `name:"group_id" env:"DATABRICKS_GROUP_ID"`

	Token    string `name:"token" env:"DATABRICKS_TOKEN" auth:"pat,sensitive"`
	Username string `name:"username" env:"DATABRICKS_USERNAME" auth:"basic"`
	Password string `name:"password" env:"DATABRICKS_PASSWORD" auth:"basic,sensitive"`

	// Connection profile specified within ~/.databrickscfg.
	Profile string `name:"profile" env:"DATABRICKS_CONFIG_PROFILE"`

	// Location of the Databricks CLI credentials file, that is created
	// by `databricks configure --token` command. By default, it is located
	// in ~/.databrickscfg.
	ConfigFile string `name:"config_file" env:"DATABRICKS_CONFIG_FILE"`

	// OAuthPort is the port to use for the OAuth2 callback server. If not set,
	// the default port with fallback is used. This means that setting a port
	// will disable the fallback mechanism.
	OAuthCallbackPort int `name:"oauth_callback_port" env:"DATABRICKS_OAUTH_CALLBACK_PORT" auth:"-"`

	GoogleServiceAccount string `name:"google_service_account" env:"DATABRICKS_GOOGLE_SERVICE_ACCOUNT" auth:"google" auth_types:"google-id"`
	GoogleCredentials    string `name:"google_credentials" env:"GOOGLE_CREDENTIALS" auth:"google,sensitive" auth_types:"google-credentials"`

	// Azure Resource Manager ID for Azure Databricks workspace, which is exhanged for a Host
	AzureResourceID string `name:"azure_workspace_resource_id" env:"DATABRICKS_AZURE_RESOURCE_ID" auth:"azure" auth_types:"azure-cli,azure-msi"`

	AzureUseMSI       bool   `name:"azure_use_msi" env:"ARM_USE_MSI" auth:"azure" auth_types:"azure-msi"`
	AzureClientSecret string `name:"azure_client_secret" env:"ARM_CLIENT_SECRET" auth:"azure,sensitive" auth_types:"azure-client-secret"`
	AzureClientID     string `name:"azure_client_id" env:"ARM_CLIENT_ID" auth:"azure" auth_types:"azure-client-secret,azure-msi"`
	AzureTenantID     string `name:"azure_tenant_id" env:"ARM_TENANT_ID" auth:"azure" auth_types:"azure-cli,azure-client-secret"`

	// Parameters to request Azure OIDC token on behalf of Github Actions.
	// Ref: https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-cloud-providers
	//
	// TODO: Users are not expected to set these environment variables. They
	// should be removed from the config in the future.
	ActionsIDTokenRequestURL   string `name:"actions_id_token_request_url" env:"ACTIONS_ID_TOKEN_REQUEST_URL"`
	ActionsIDTokenRequestToken string `name:"actions_id_token_request_token" env:"ACTIONS_ID_TOKEN_REQUEST_TOKEN"`

	// AzureEnvironment (PUBLIC, USGOVERNMENT, CHINA) has specific set of API endpoints. Starting from v0.26.0,
	// the environment is determined based on the workspace hostname, if it's specified.
	AzureEnvironment string `name:"azure_environment" env:"ARM_ENVIRONMENT"`

	// Azure Login Application ID. Must be set if authenticating for non-production workspaces. Starting from v0.26.0,
	// the correct Azure Login App ID is determined based on the Azure Databricks Workspace hostname.
	//
	// Deprecated: this configuration property no longer has any effect and will be removed in the future
	// versions of Go SDK.
	AzureLoginAppID string `name:"azure_login_app_id" env:"DATABRICKS_AZURE_LOGIN_APP_ID" auth:"azure"`

	ClientID     string `name:"client_id" env:"DATABRICKS_CLIENT_ID" auth:"oauth" auth_types:"oauth-m2m"`
	ClientSecret string `name:"client_secret" env:"DATABRICKS_CLIENT_SECRET" auth:"oauth,sensitive" auth_types:"oauth-m2m"`

	// Scopes is a list of OAuth scopes to request when authenticating.
	//
	// WARNING:
	//   - This feature is still in development and may not work as expected
	//   - This feature is EXPERIMENTAL and may change or be removed without notice.
	//   - Do NOT use this feature in production environments.
	//
	// Notes:
	//   - If Scopes is nil or empty, the default ["all-apis"] scope will be used for backward compatibility.
	//   - For U2M authentication, the "offline_access" scope will automatically be added to obtain a refresh token
	//     unless you set DisableOAuthRefreshToken to true.
	//   - You cannot set Scopes via environment variables.
	//   - The scopes list will be sorted in-place during configuration resolution.
	//   - The U2M token cache currently does NOT support differentiated caching for scopes.
	Scopes []string `name:"scopes" auth:"-"`

	// DisableOAuthRefreshToken controls whether a refresh token should be requested
	// during the U2M authentication flow (default to false).
	DisableOAuthRefreshToken bool `name:"disable_oauth_refresh_token" env:"DATABRICKS_DISABLE_OAUTH_REFRESH_TOKEN" auth:"-"`

	// DisableAsyncTokenRefresh controls whether asynchronous token refresh
	// should be disabled for OAuth tokens.
	DisableAsyncTokenRefresh bool

	// Path to the Databricks CLI (version >= 0.100.0).
	DatabricksCliPath string `name:"databricks_cli_path" env:"DATABRICKS_CLI_PATH" auth_types:"databricks-cli"`

	// When multiple auth attributes are available in the environment, use the auth type
	// specified by this argument. This argument also holds currently selected auth.
	AuthType string `name:"auth_type" env:"DATABRICKS_AUTH_TYPE" auth:"-"`

	// Path to the file containing an OIDC ID token.
	OIDCTokenFilepath string `name:"databricks_id_token_filepath" env:"DATABRICKS_OIDC_TOKEN_FILEPATH" auth:"file-oidc"`

	// Environment variable name that contains an OIDC ID token.
	OIDCTokenEnv string `name:"oidc_token_env" env:"DATABRICKS_OIDC_TOKEN_ENV" auth:"env-oidc"`

	// Skip SSL certificate verification for HTTP calls.
	// Use at your own risk or for unit testing purposes.
	InsecureSkipVerify bool `name:"skip_verify" auth:"-"`

	// Number of seconds for HTTP timeout. Default is 60 (1 minute).
	HTTPTimeoutSeconds int `name:"http_timeout_seconds" auth:"-"`

	// Truncate JSON fields in JSON above this limit. Default is 96.
	DebugTruncateBytes int `name:"debug_truncate_bytes" env:"DATABRICKS_DEBUG_TRUNCATE_BYTES" auth:"-"`

	// Debug HTTP headers of requests made by the provider. Default is false.
	DebugHeaders bool `name:"debug_headers" env:"DATABRICKS_DEBUG_HEADERS" auth:"-"`

	// Headers, if set, is called on every request made by the client to set
	// additional HTTP headers, on top of normal authentication. It runs after
	// authentication, so it must not overwrite headers such as Authorization.
	Headers func(*http.Request) error

	// Maximum number of requests per second made to Databricks REST API. Default is 15 RPS.
	RateLimitPerSecond int `name:"rate_limit" env:"DATABRICKS_RATE_LIMIT" auth:"-"`

	// Number of seconds to keep retrying HTTP requests. Default is 300 (5 minutes).
	// If negative, the client will retry on retriable errors indefinitely.
	RetryTimeoutSeconds int `name:"retry_timeout_seconds" auth:"-"`

	// HTTPTransport can be overriden for unit testing and together with tooling like https://github.com/google/go-replayers
	HTTPTransport http.RoundTripper

	// Cloud is the cloud provider for this Databricks deployment (AWS, Azure, GCP, or CloudUnknown).
	//
	// Experimental: subject to change.
	Cloud environment.Cloud `name:"cloud" env:"DATABRICKS_CLOUD" auth:"-"`

	DatabricksEnvironment *environment.DatabricksEnvironment

	// When using Workload Identity Federation, the audience to specify when fetching an ID token from the ID token supplier.
	TokenAudience string `name:"audience" env:"DATABRICKS_TOKEN_AUDIENCE" auth:"-"`

	Loaders []Loader

	// resolvedHostType is the host type resolved from the /.well-known/databricks-config
	// discovery endpoint.
	resolvedHostType HostType

	// marker for configuration resolving
	resolved bool

	// internal client used in for authentication purposes:
	//  - Databricks Metadata Service: request/refresh tokens from parent processes, like Databricks VSCode extension
	//	- Azure Active Directory (AAD): request/refresh OAuth tokens
	//  - Azure Instance Metadata Service (IMDS): request/refresh OAuth tokens for Azure Managed Identity
	//  - Azure Resource Manager (ARM): resolve host if only Azure Databricks Resource ID provided
	refreshClient *httpclient.ApiClient

	// internal background context used for authentication purposes together with refreshClient
	refreshCtx context.Context

	// internal client used to fetch Azure Tenant ID from Databricks Login endpoint
	azureTenantIdFetchClient *http.Client

	// marker for testing fixture
	isTesting bool

	// Mutex used by Authenticate method to guard `auth`, which
	// has to be lazily created on the first request to Databricks API.
	// It is done because databricks host and token may often be available
	// only in the middle of Terraform DAG execution.
	// This mutex is also used for config resolution.
	mu sync.Mutex

	// HTTP request interceptor, that assigns Authorization header
	credentialsProvider credentials.CredentialsProvider

	// Keep track of the source of each attribute
	attrSource map[string]Source

	// OpenID Connect discovery URL. When set, OIDC endpoints are fetched directly
	// from this URL instead of the default host-type-based well-known endpoint logic.
	// Mirrors discoveryUrl in the Java SDK.
	//
	// Experimental: subject to change.
	DiscoveryURL string `name:"discovery_url" env:"DATABRICKS_DISCOVERY_URL" auth:"-"`

	// HostMetadataResolver, if set, overrides the default HTTP fetch of
	// /.well-known/databricks-config during config resolution. This allows
	// callers (e.g., the CLI) to provide cached metadata directly.
	// Use [Config.DefaultHostMetadataResolver] to delegate to the SDK's
	// built-in fetch on cache miss.
	//
	// Experimental: subject to change.
	HostMetadataResolver HostMetadataResolver `auth:"-"`
}

// NewWithWorkspaceHost returns a new instance of the Config with the host set to
// the workspace host. Fields that are not relevant to workspace-level config,
// like account ID, are omitted. Workspace-level attributes that cannot be
// computed from the host alone, like Azure Resource ID, are also omitted.
func (c *Config) NewWithWorkspaceHost(host string) (*Config, error) {
	err := c.EnsureResolved()
	if err != nil {
		return nil, err
	}

	fieldsToSkip := map[string]struct{}{
		"Host":            {},
		"AzureResourceID": {},
		"AccountID":       {},
	}
	res := &Config{}
	cv := reflect.ValueOf(c).Elem()
	resv := reflect.ValueOf(res).Elem()
	for i := 0; i < resv.NumField(); i++ {
		field := resv.Field(i)
		if !field.CanSet() {
			continue
		}
		if _, ok := fieldsToSkip[resv.Type().Field(i).Name]; ok {
			continue
		}
		field.Set(cv.Field(i))
	}

	res.Host = host
	res.isTesting = c.isTesting
	// We need to reresolve the config with the updated host in general. For
	// example, the audience for OAuth tokens provided by GCP is derived from
	// the host, so account-level tokens cannot be used at workspace-level or
	// vice-versa.
	//
	// In the future, when unified login is widely available, we may be able to
	// reuse the authentication visitor specifically for in-house OAuth.
	return res, nil
}

// Authenticate adds special headers to HTTP request to authorize it to work with Databricks REST API
func (c *Config) Authenticate(r *http.Request) error {
	err := c.EnsureResolved()
	if err != nil {
		return err
	}
	err = c.authenticateIfNeeded()
	if err != nil {
		return err
	}
	return c.credentialsProvider.SetHeaders(r)
}

// Authenticate returns an OAuth token for the current configuration. It will
// return an error if the CredentialsStrategy does not support OAuth tokens.
//
// Deprecated: Use GetTokenSource instead.
func (c *Config) GetToken() (*oauth2.Token, error) {
	ts := c.GetTokenSource()
	return ts.Token(context.Background())
}

// GetTokenSource returns an OAuth token source for the current configuration.
// It will return an error if the CredentialsStrategy does not support OAuth
// tokens.
func (c *Config) GetTokenSource() auth.TokenSource {
	if err := c.EnsureResolved(); err != nil {
		return errorTokenSource(err)
	}
	if err := c.authenticateIfNeeded(); err != nil {
		return errorTokenSource(err)
	}
	if h, ok := c.credentialsProvider.(credentials.OAuthCredentialsProvider); ok {
		return h
	} else {
		return errorTokenSource(fmt.Errorf("OAuth Token not supported for current auth type %s", c.AuthType))
	}
}

func errorTokenSource(err error) auth.TokenSource {
	return auth.TokenSourceFn(func(context.Context) (*oauth2.Token, error) {
		return nil, err
	})
}

// IsAzure returns if the client is configured for Azure Databricks.
func (c *Config) IsAzure() bool {
	if c.AzureResourceID != "" {
		return true
	}
	if c.Cloud != environment.CloudUnknown {
		return c.Cloud == environment.CloudAzure
	}
	return c.Environment().Cloud == environment.CloudAzure
}

// IsGcp returns if the client is configured for Databricks on Google Cloud.
func (c *Config) IsGcp() bool {
	if c.Cloud != environment.CloudUnknown {
		return c.Cloud == environment.CloudGCP
	}
	return c.Environment().Cloud == environment.CloudGCP
}

// IsAws returns if the client is configured for Databricks on AWS.
func (c *Config) IsAws() bool {
	if c.Cloud != environment.CloudUnknown {
		return c.Cloud == environment.CloudAWS
	}
	return c.Host != "" && !c.IsAzure() && !c.IsGcp()
}

// IsAccountClient returns true if client is configured for Accounts API.
//
// Deprecated: Use HostType() if possible, or ConfigType() if necessary.
func (c *Config) IsAccountClient() bool {
	if c.AccountID != "" && c.isTesting {
		return true
	}

	accountsPrefixes := []string{
		"https://accounts.",
		"https://accounts-dod.",
	}
	for _, prefix := range accountsPrefixes {
		if strings.HasPrefix(c.Host, prefix) {
			return true
		}
	}
	return false
}

// normalizedHost returns the normalized host for the client.
// Small utility function to avoid duplicating the logic in HostType().
func normalizedHost(host string) string {
	if host != "" && !strings.Contains(host, "://") {
		host = "https://" + host
	}
	return host
}

// HostType returns the type of host that the client is configured for.
// When host metadata has been resolved (via /.well-known/databricks-config),
// the resolved host type is returned directly. Otherwise, the host type is
// inferred from URL patterns as a fallback.
//
// Deprecated: The SDK is moving towards host-agnostic behavior where host URL
// patterns are no longer used to determine client capabilities. This method is
// retained for backwards compatibility but should not be used in new code.
func (c *Config) HostType() HostType {
	// If host metadata resolved a known host type, use it.
	if c.resolvedHostType != HostTypeUnknown {
		return c.resolvedHostType
	}

	// Normalize the host to ensure the scheme is present before checking
	// prefixes. Profiles saved without "https://" (e.g. from user input)
	// would otherwise fail the prefix check and be misclassified as
	// workspace hosts.
	host := normalizedHost(c.Host)
	accountsPrefixes := []string{
		"https://accounts.",
		"https://accounts-dod.",
	}
	for _, prefix := range accountsPrefixes {
		if strings.HasPrefix(host, prefix) {
			return AccountHost
		}
	}

	return WorkspaceHost
}

// ConfigType returns the type of config that the client is configured for.
// Returns InvalidConfig if the config is invalid.
// Use of this function should be avoided where possible, because we plan
// to remove WorkspaceClient and AccountClient in favor of a single unified
// client in the future.
func (c *Config) ConfigType() ConfigType {
	switch c.HostType() {
	case AccountHost:
		return AccountConfig
	case WorkspaceHost:
		return WorkspaceConfig
	default:
		return InvalidConfig
	}
}

func (c *Config) EnsureResolved() error {
	if c.resolved {
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.resolved {
		return nil
	}
	if len(c.Loaders) == 0 {
		c.Loaders = []Loader{
			ConfigAttributes,
			ConfigFile,
		}
	}
	ctx := context.Background()
	for _, loader := range c.Loaders {
		logger.Tracef(ctx, "Loading config via %s", loader.Name())
		err := loader.Configure(c)
		if err != nil {
			return c.wrapDebug(fmt.Errorf("resolve: %w", err))
		}
	}
	err := ConfigAttributes.Validate(c)
	if err != nil {
		return c.wrapDebug(fmt.Errorf("validate: %w", err))
	}
	c.refreshCtx = ctx
	c.refreshClient = httpclient.NewApiClient(httpclient.ClientConfig{
		DebugHeaders:       c.DebugHeaders,
		DebugTruncateBytes: c.DebugTruncateBytes,
		InsecureSkipVerify: c.InsecureSkipVerify,
		RetryTimeout:       time.Duration(c.RetryTimeoutSeconds) * time.Second,
		HTTPTimeout:        time.Duration(c.HTTPTimeoutSeconds) * time.Second,
		Transport:          c.HTTPTransport,
		ErrorMapper:        c.refreshTokenErrorMapper,
		TransientErrors: []string{
			"throttled",
			"too many requests",
			"429",
			"request limit exceeded",
			"rate limit",
		},
	})
	if c.azureTenantIdFetchClient == nil {
		c.azureTenantIdFetchClient = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				// Do not follow redirects
				return http.ErrUseLastResponse
			},
		}
	}
	slices.Sort(c.Scopes)
	c.Scopes = slices.Compact(c.Scopes)
	c.resolveHostMetadata(ctx)
	c.resolved = true
	return nil
}

func (c *Config) WithTesting() *Config {
	c.isTesting = true
	return c
}

func (c *Config) CanonicalHostName() string {
	// Missing host is tolerated here.
	_ = c.fixHostIfNeeded()
	return c.Host
}

func (c *Config) GetScopes() []string {
	if len(c.Scopes) == 0 {
		return defaultScopes
	}
	return c.Scopes
}

func (c *Config) wrapDebug(err error) error {
	debug := ConfigAttributes.DebugString(c)
	if debug == "" {
		return err
	}
	return fmt.Errorf("%w. %s", err, debug)
}

// authenticateIfNeeded lazily authenticates across authorizers or returns error
func (c *Config) authenticateIfNeeded() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.credentialsProvider != nil {
		return nil
	}
	if c.Credentials == nil {
		c.Credentials = DefaultCredentialStrategyProvider()
	}
	if err := c.fixHostIfNeeded(); err != nil && !errors.Is(err, ErrNoHostConfigured) {
		return err
	}
	ctx := c.refreshClient.InContextForOAuth2(c.refreshCtx)
	credentialsProvider, err := c.Credentials.Configure(ctx, c)
	if err != nil {
		return c.wrapDebug(fmt.Errorf("%s auth: %w", c.Credentials.Name(), err))
	}
	if credentialsProvider == nil {
		return c.wrapDebug(fmt.Errorf("%s auth: not configured", c.Credentials.Name()))
	}
	c.credentialsProvider = credentialsProvider
	c.AuthType = c.Credentials.Name()
	if err := c.fixHostIfNeeded(); err != nil {
		return err
	}
	// TODO: error customization
	return nil
}

func (c *Config) fixHostIfNeeded() error {
	// Nothing to fix if the host isn't set.
	if c.Host == "" {
		return nil
	}
	parsedHost, err := url.Parse(c.Host)
	if err != nil {
		return err
	}
	// If the host is empty, assume the scheme wasn't included.
	if parsedHost.Host == "" {
		parsedHost, err = url.Parse("https://" + c.Host)
		if err != nil {
			return err
		}
	}
	if parsedHost.Hostname() == "" {
		return ErrNoHostConfigured
	}
	// SPOG URLs pasted from the Databricks UI carry the workspace ID as
	// ?o= (or ?workspace_id=) and the account ID as ?a= (or ?account_id=).
	// Promote those into WorkspaceID/AccountID before we strip the query,
	// so requests get the X-Databricks-Org-Id header instead of hitting the
	// SPOG without routing and getting back the login HTML page.
	if parsedHost.RawQuery != "" {
		q := parsedHost.Query()
		if c.WorkspaceID == "" {
			c.WorkspaceID = workspaceIDFromQuery(q)
		}
		if c.AccountID == "" {
			c.AccountID = accountIDFromQuery(q)
		}
	}
	// Create new instance to ensure other fields are initialized as empty.
	parsedHost = &url.URL{
		Scheme: parsedHost.Scheme,
		Host:   parsedHost.Host,
	}
	// Store sanitized version of c.Host.
	c.Host = parsedHost.String()
	return nil
}

func workspaceIDFromQuery(q url.Values) string {
	// ?w= is the unified workspace addressing query parameter. It supersedes
	// the older ?o=/?workspace_id= forms and accepts a broader range of
	// workspace identifier formats — both classic numeric workspace IDs and
	// other identifier formats the server understands — so we don't apply the
	// numeric-only validation that we use for ?o=/?workspace_id=.
	if v := q.Get("w"); v != "" {
		return v
	}
	for _, key := range []string{"o", "workspace_id"} {
		v := q.Get(key)
		if v == "" {
			continue
		}
		if _, err := strconv.ParseInt(v, 10, 64); err == nil {
			return v
		}
	}
	return ""
}

func accountIDFromQuery(q url.Values) string {
	for _, key := range []string{"a", "account_id"} {
		if v := q.Get(key); v != "" {
			return v
		}
	}
	return ""
}

// ErrNoHostConfigured is the error returned when a user tries to authenticate
// without a host configured. Applications can check for this error to provide
// more user-friendly error messages.
var ErrNoHostConfigured = fmt.Errorf("no host configured")

func (c *Config) refreshTokenErrorMapper(ctx context.Context, resp common.ResponseWrapper) error {
	defaultErr := httpclient.DefaultErrorMapper(ctx, resp)
	if defaultErr == nil {
		return nil
	}
	err, ok := defaultErr.(*httpclient.HttpError)
	if !ok {
		return err
	}
	if c.IsAzure() {
		return c.mapAzureError(err)
	}
	return &tokenError{
		message: err.Message,
		err:     err,
	}
}

// getOidcEndpoints returns the OAuth endpoints for the current configuration.
// If DiscoveryURL is set, endpoints are fetched directly from it. Otherwise
// falls back to host-type-based well-known endpoint logic.
func (c *Config) getOidcEndpoints(ctx context.Context) (*u2m.OAuthAuthorizationServer, error) {
	c.EnsureResolved()
	oauthClient := &u2m.BasicOAuthEndpointSupplier{
		Client: c.refreshClient,
	}
	if c.DiscoveryURL != "" {
		return oauthClient.GetEndpointsFromURL(ctx, c.DiscoveryURL)
	}
	host := c.CanonicalHostName()
	switch c.HostType() {
	case AccountHost:
		return oauthClient.GetAccountOAuthEndpoints(ctx, host, c.AccountID)
	case WorkspaceHost:
		return oauthClient.GetWorkspaceOAuthEndpoints(ctx, host)
	default:
		return nil, fmt.Errorf("unknown host type: %v", c.HostType())
	}
}

// resolveHostMetadata populates missing config fields from the host's
// /.well-known/databricks-config discovery endpoint. It back-fills AccountID,
// WorkspaceID, Cloud, and DiscoveryURL (with any {account_id} placeholder substituted)
// if those fields are not already set.
//
// Errors from the metadata endpoint are non-fatal: a warning is logged and the
// method returns without modifying the config. This mirrors the Python SDK
// behavior where metadata resolution is best-effort during config init.
func (c *Config) resolveHostMetadata(ctx context.Context) {
	if c.Host == "" {
		return
	}
	if err := c.fixHostIfNeeded(); err != nil {
		logger.Warnf(ctx, "Failed to fix host for metadata resolution: %v", err)
		return
	}

	resolver := c.HostMetadataResolver
	if resolver == nil {
		if factory := DefaultHostMetadataResolverFactory; factory != nil {
			resolver = factory(c)
		}
	}

	var meta *HostMetadata
	var err error
	if resolver != nil {
		meta, err = resolver(ctx, c.CanonicalHostName())
	} else {
		meta, err = getHostMetadata(ctx, c.CanonicalHostName(), c.refreshClient)
	}
	if err != nil {
		logger.Warnf(ctx, "Failed to resolve host metadata: %v. Falling back to user config.", err)
		return
	}
	if meta == nil {
		return
	}
	if c.AccountID == "" && meta.AccountID != "" {
		logger.Debugf(ctx, "Resolved account_id from host metadata: %q", meta.AccountID)
		c.AccountID = meta.AccountID
	}
	if c.WorkspaceID == "" && meta.WorkspaceID != "" {
		logger.Debugf(ctx, "Resolved workspace_id from host metadata: %q", meta.WorkspaceID)
		c.WorkspaceID = meta.WorkspaceID
	}
	if c.Cloud == "" && meta.Cloud != environment.CloudUnknown {
		logger.Debugf(ctx, "Resolved cloud from host metadata: %q", meta.Cloud)
		c.Cloud = meta.Cloud
	}
	if c.Cloud == "" {
		c.Cloud = c.Environment().Cloud
		logger.Debugf(ctx, "Resolved cloud from hostname: %q", c.Cloud)
	}
	if c.resolvedHostType == HostTypeUnknown && meta.HostType != HostTypeUnknown {
		logger.Debugf(ctx, "Resolved host_type from host metadata: %q", meta.HostType)
		c.resolvedHostType = meta.HostType
	}
	if c.TokenAudience == "" && len(meta.TokenFederationDefaultOIDCAudiences) > 0 {
		logger.Debugf(ctx, "Resolved token_audience from host metadata token_federation_default_oidc_audiences: %q", meta.TokenFederationDefaultOIDCAudiences[0])
		c.TokenAudience = meta.TokenFederationDefaultOIDCAudiences[0]
	}
	if c.TokenAudience == "" && meta.WorkspaceID == "" && c.AccountID != "" {
		logger.Debugf(ctx, "Setting token_audience to account_id for account host: %q", c.AccountID)
		c.TokenAudience = c.AccountID
	}
	if c.DiscoveryURL == "" {
		if meta.OIDCEndpoint == "" {
			logger.Warnf(ctx, "Host metadata missing oidc_endpoint; skipping discovery URL resolution")
			return
		}
		oidcRoot := meta.OIDCEndpoint
		if strings.Contains(oidcRoot, "{account_id}") {
			if c.AccountID == "" {
				logger.Warnf(ctx, "Host metadata oidc_endpoint contains {account_id} placeholder but account_id is not set; skipping discovery URL resolution")
				return
			}
			oidcRoot = strings.ReplaceAll(oidcRoot, "{account_id}", c.AccountID)
		}
		oidcRoot = strings.TrimRight(oidcRoot, "/")
		discoveryURL := oidcRoot + "/.well-known/oauth-authorization-server"
		logger.Debugf(ctx, "Resolved discovery_url from host metadata: %q", discoveryURL)
		c.DiscoveryURL = discoveryURL
	}
}

// DefaultHostMetadataResolver returns a [HostMetadataResolver] that fetches
// host metadata from the /.well-known/databricks-config endpoint using the
// SDK's internal HTTP client (with retries, proxy support, and debug logging).
//
// This is intended for use inside a custom [Config.HostMetadataResolver] to
// delegate to the SDK's default fetch on cache miss. It must only be called
// from within a resolver during [Config.EnsureResolved], when the internal
// HTTP client is guaranteed to be initialized.
//
// Experimental: subject to change.
func (c *Config) DefaultHostMetadataResolver() HostMetadataResolver {
	return func(ctx context.Context, host string) (*HostMetadata, error) {
		return getHostMetadata(ctx, host, c.refreshClient)
	}
}

func (c *Config) getOAuthArgument() (u2m.OAuthArgument, error) {
	err := c.EnsureResolved()
	if err != nil {
		return nil, err
	}
	host := c.CanonicalHostName()
	profile := c.Profile
	switch c.HostType() {
	case AccountHost:
		return u2m.NewProfileAccountOAuthArgument(host, c.AccountID, profile)
	case WorkspaceHost:
		return u2m.NewProfileWorkspaceOAuthArgument(host, profile)
	default:
		return nil, fmt.Errorf("unknown host type: %v", c.HostType())
	}
}
