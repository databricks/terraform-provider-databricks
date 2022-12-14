package common

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/api/option"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
)

// Default settings
const (
	DefaultTruncateBytes      = 96
	DefaultRateLimitPerSecond = 15
	DefaultHTTPTimeoutSeconds = 60
)

// DatabricksClient holds properties needed for authentication and HTTP client setup
// fields with `name` struct tags become Terraform provider attributes. `env` struct tag
// can hold one or more coma-separated env variable names to find value, if not specified
// directly. `auth` struct tag describes the type of conflicting authentication used.
type DatabricksClient struct {
	Config *config.Config
	Client *client.DatabricksClient

	Host     string `name:"host" env:"DATABRICKS_HOST"`
	Token    string `name:"token" env:"DATABRICKS_TOKEN" auth:"token,sensitive"`
	Username string `name:"username" env:"DATABRICKS_USERNAME" auth:"password"`
	Password string `name:"password" env:"DATABRICKS_PASSWORD" auth:"password,sensitive"`

	ClientID      string `name:"client_id" env:"DATABRICKS_CLIENT_ID" auth:"oauth"`
	ClientSecret  string `name:"client_secret" env:"DATABRICKS_CLIENT_SECRET" auth:"oauth,sensitive"`
	TokenEndpoint string `name:"token_endpoint" env:"DATABRICKS_TOKEN_ENDPOINT" auth:"oauth"`

	// Databricks Account ID for Accounts API. This field is used in dependencies.
	AccountID string `name:"account_id" env:"DATABRICKS_ACCOUNT_ID"`

	// Connection profile specified within ~/.databrickscfg.
	Profile string `name:"profile" env:"DATABRICKS_CONFIG_PROFILE" auth:"config profile"`

	// Location of the Databricks CLI credentials file, that is created
	// by `databricks configure --token` command. By default, it is located
	// in ~/.databrickscfg.
	ConfigFile string `name:"config_file" env:"DATABRICKS_CONFIG_FILE"`

	GoogleServiceAccount string `name:"google_service_account" env:"DATABRICKS_GOOGLE_SERVICE_ACCOUNT" auth:"google"`
	GoogleCredentials    string `name:"google_credentials" env:"GOOGLE_CREDENTIALS" auth:"google,sensitive"`

	AzureResourceID           string `name:"azure_workspace_resource_id" env:"DATABRICKS_AZURE_RESOURCE_ID" auth:"azure"`
	AzureUseMSI               bool   `name:"azure_use_msi" env:"ARM_USE_MSI" auth:"azure"`
	AzureClientSecret         string `name:"azure_client_secret" env:"ARM_CLIENT_SECRET" auth:"azure,sensitive"`
	AzureClientID             string `name:"azure_client_id" env:"ARM_CLIENT_ID" auth:"azure"`
	AzureTenantID             string `name:"azure_tenant_id" env:"ARM_TENANT_ID" auth:"azure"`
	AzurermEnvironment        string `name:"azure_environment" env:"ARM_ENVIRONMENT"`
	AzureDatabricksLoginAppId string `name:"azure_login_app_id" env:"DATABRICKS_AZURE_LOGIN_APP_ID" auth:"azure"`

	// When multiple auth attributes are available in the environment, use the auth type
	// specified by this argument. This argument also holds currently selected auth.
	AuthType string `name:"auth_type" auth:"-"`

	// Azure Environment endpoints
	AzureEnvironment *azure.Environment

	// Skip SSL certificate verification for HTTP calls.
	// Use at your own risk or for unit testing purposes.
	InsecureSkipVerify bool `name:"skip_verify" auth:"-"`
	HTTPTimeoutSeconds int  `name:"http_timeout_seconds" auth:"-"`

	// Truncate JSON fields in JSON above this limit. Default is 96.
	DebugTruncateBytes int `name:"debug_truncate_bytes" env:"DATABRICKS_DEBUG_TRUNCATE_BYTES" auth:"-"`

	// Debug HTTP headers of requests made by the provider. Default is false.
	DebugHeaders bool `name:"debug_headers" env:"DATABRICKS_DEBUG_HEADERS" auth:"-"`

	// Maximum number of requests per second made to Databricks REST API.
	RateLimitPerSecond int `name:"rate_limit" env:"DATABRICKS_RATE_LIMIT" auth:"-"`

	// OAuth token refreshers for Azure to be used within `authVisitor`
	azureAuthorizer autorest.Authorizer

	// options used to enable unit testing mode for OIDC
	googleAuthOptions []option.ClientOption

	// Mutex used by Authenticate method to guard `authVisitor`, which
	// has to be lazily created on the first request to Databricks API.
	// It is done because databricks host and token may often be available
	// only in the middle of Terraform DAG execution.
	authMutex sync.Mutex

	// HTTP request interceptor, that assigns Authorization header
	authVisitor func(r *http.Request) error

	// Databricks REST API rate limiter
	rateLimiter *rate.Limiter

	// Terraform provider instance to include Terraform binary version in
	// User-Agent header
	Provider *schema.Provider

	// retryalble HTTP client
	httpClient *retryablehttp.Client

	// configuration attributes that were used to initialise client.
	configAttributesUsed []string

	// callback used to create API1.2 call wrapper, which simplifies unit tessting
	commandFactory func(context.Context, *DatabricksClient) CommandExecutor
}

type ConfigAttribute struct {
	Name      string
	Kind      reflect.Kind
	EnvVars   []string
	Auth      string
	Sensitive bool
	Internal  bool
	num       int
}

func (ca *ConfigAttribute) Set(client *DatabricksClient, i any) error {
	rv := reflect.ValueOf(client)
	field := rv.Elem().Field(ca.num)
	switch ca.Kind {
	case reflect.String:
		field.SetString(i.(string))
	case reflect.Bool:
		field.SetBool(i.(bool))
	case reflect.Int:
		field.SetInt(int64(i.(int)))
	default:
		// must extensively test with providerFixture to avoid this one
		return fmt.Errorf("cannot set %s of unknown type %s", ca.Name, reflectKind(ca.Kind))
	}
	return nil
}

func (ca *ConfigAttribute) GetString(client *DatabricksClient) string {
	rv := reflect.ValueOf(client)
	field := rv.Elem().Field(ca.num)
	return fmt.Sprintf("%v", field.Interface())
}

// ClientAttributes returns meta-representation of DatabricksClient configuration options
func ClientAttributes() (attrs []ConfigAttribute) {
	t := reflect.TypeOf(DatabricksClient{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		nameTag := field.Tag.Get("name")
		if nameTag == "" {
			continue
		}
		sensitive := false
		auth := field.Tag.Get("auth")
		authSplit := strings.Split(auth, ",")
		if len(authSplit) == 2 {
			auth = authSplit[0]
			sensitive = authSplit[1] == "sensitive"
		}
		// internal config fields are skipped in debugging
		internal := false
		if auth == "-" {
			auth = ""
			internal = true
		}
		attr := ConfigAttribute{
			Name:      nameTag,
			Auth:      auth,
			Kind:      field.Type.Kind(),
			Sensitive: sensitive,
			Internal:  internal,
			num:       i,
		}
		envTag := field.Tag.Get("env")
		if envTag != "" {
			attr.EnvVars = strings.Split(envTag, ",")
		}
		attrs = append(attrs, attr)
	}
	return
}

// Configure client to work, optionally specifying configuration attributes used
func (c *DatabricksClient) Configure(attrsUsed ...string) error {
	c.configAttributesUsed = attrsUsed
	c.configureHTTPCLient()
	if c.DebugTruncateBytes == 0 {
		c.DebugTruncateBytes = DefaultTruncateBytes
	}
	// AzureEnvironment could be used in the different contexts, not only for Auzre Authentication
	// lack of this lead to crash (see issue #831)
	azureEnvironment, err := c.getAzureEnvironment()
	if err != nil {
		return fmt.Errorf("cannot get azure environment: %w", err)
	}
	c.AzureEnvironment = &azureEnvironment

	return nil
}

func (c *DatabricksClient) configDebugString() string {
	debug := []string{}
	for _, attr := range ClientAttributes() {
		if attr.Internal && !c.DebugHeaders {
			continue
		}
		value := attr.GetString(c)
		if value == "" {
			continue
		}
		if attr.Name == "azure_use_msi" && value == "false" {
			// include Azure MSI info only when it's relevant
			continue
		}
		if attr.Sensitive {
			value = "***REDACTED***"
		}
		debug = append(debug, fmt.Sprintf("%s=%v", attr.Name, value))
	}
	return strings.Join(debug, ", ") // lgtm[go/clear-text-logging]
}

// Authenticate lazily authenticates across authorizers or returns error
func (c *DatabricksClient) Authenticate(ctx context.Context) error {
	if c.authVisitor != nil {
		return nil
	}
	c.authMutex.Lock()
	defer c.authMutex.Unlock()
	if c.authVisitor != nil {
		return nil
	}
	// Fix host prior to auth, because it may be used in the OIDC flow as "audience" field.
	// If necessary, this function adds a scheme and strips a trailing slash.
	if err := c.fixHost(); err != nil {
		return err
	}
	type auth struct {
		configure func(context.Context) (func(*http.Request) error, error)
		name      string
	}
	providers := []auth{
		{c.configureWithPat, "pat"},
		{c.configureWithBasicAuth, "basic"},
		{c.configureWithOAuthM2M, "oauth-m2m"},
		{c.configureWithAzureClientSecret, "azure-client-secret"},
		{c.configureWithAzureManagedIdentity, "azure-msi"},
		{c.configureWithAzureCLI, "azure-cli"},
		{c.configureWithGoogleCrendentials, "google-creds"},
		{c.configureWithGoogleForAccountsAPI, "google-accounts"},
		{c.configureWithGoogleForWorkspace, "google-workspace"},
		{c.configureWithDatabricksCfg, "databricks-cli"},
	}
	// try configuring authentication with different methods
	for _, auth := range providers {
		if c.AuthType != "" && auth.name != c.AuthType {
			// ignore other auth types if one is explicitly enforced
			log.Printf("[INFO] Ignoring %s auth, because %s is preferred", auth.name, c.AuthType)
			continue
		}
		authorizer, err := auth.configure(ctx)
		if err != nil {
			return c.niceAuthError(fmt.Sprintf("cannot configure %s auth: %s", auth.name, err))
		}
		if authorizer == nil {
			// try the next method.
			continue
		}
		// even though this may complain about clear text logging, passwords are replaced with `***`
		log.Printf("[INFO] Configured %s auth: %s", auth.name, c.configDebugString()) // lgtm[go/clear-text-logging]
		c.authVisitor = authorizer
		c.AuthType = auth.name
		c.fixHost()
		return nil
	}
	if c.AuthType != "" {
		return c.niceAuthError(fmt.Sprintf("cannot configure %s auth", c.AuthType))
	}
	if c.Host == "" && IsData.GetOrUnknown(ctx) == "yes" {
		return c.niceAuthError("workspace is most likely not created yet, because the `host` " +
			"is empty. Please add `depends_on = [databricks_mws_workspaces.this]` or " +
			"`depends_on = [azurerm_databricks_workspace.this]` to every data resource. See " +
			"https://www.terraform.io/docs/language/resources/behavior.html more info")
	}
	return c.niceAuthError("authentication is not configured for provider.")
}

func (c *DatabricksClient) niceAuthError(message string) error {
	info := ""
	if len(c.configAttributesUsed) > 0 {
		envs := []string{}
		attrs := []string{}
		usedAsEnv := map[string]bool{}
		for _, attr := range ClientAttributes() {
			if len(attr.EnvVars) == 0 {
				continue
			}
			for _, envVar := range attr.EnvVars {
				value := os.Getenv(envVar)
				if value == "" {
					continue
				}
				usedAsEnv[attr.Name] = true
				envs = append(envs, envVar)
			}
		}
		for _, attr := range c.configAttributesUsed {
			if usedAsEnv[attr] {
				continue
			}
			attrs = append(attrs, attr)
		}
		infos := []string{}
		if len(attrs) > 0 {
			infos = append(infos, fmt.Sprintf("Attributes used: %s", strings.Join(attrs, ", ")))
		}
		if len(envs) > 0 {
			infos = append(infos, fmt.Sprintf("Environment variables used: %s", strings.Join(envs, ", ")))
		}
		info = ". " + strings.Join(infos, ". ")
	}
	info = strings.TrimSuffix(info, ".")
	message = strings.TrimSuffix(message, ".")
	docUrl := "https://registry.terraform.io/providers/databricks/databricks/latest/docs#authentication"
	return fmt.Errorf("%s%s. Please check %s for details", message, info, docUrl)
}

func (c *DatabricksClient) fixHost() error {
	// Nothing to fix if the host isn't set.
	if c.Host == "" {
		return nil
	}

	u, err := url.Parse(c.Host)
	if err != nil {
		return err
	}

	// If the host is empty, assume the scheme wasn't included.
	if u.Host == "" {
		u, err = url.Parse("https://" + c.Host)
		if err != nil {
			return err
		}
	}

	// Create new instance to ensure other fields are initialized as empty.
	u = &url.URL{
		Scheme: u.Scheme,
		Host:   u.Host,
	}

	// Store sanitized version of c.Host.
	c.Host = u.String()
	return nil
}

func (c *DatabricksClient) configureWithPat(ctx context.Context) (func(*http.Request) error, error) {
	if !(c.Token != "" && c.Host != "") {
		return nil, nil
	}
	log.Printf("[INFO] Using directly configured PAT authentication")
	return c.authorizer("Bearer", c.Token), nil
}

func (c *DatabricksClient) configureWithBasicAuth(ctx context.Context) (func(*http.Request) error, error) {
	if !(c.Username != "" && c.Password != "" && c.Host != "") {
		return nil, nil
	}
	b64 := c.encodeBasicAuth(c.Username, c.Password)
	log.Printf("[INFO] Using directly configured basic authentication")
	return c.authorizer("Basic", b64), nil
}

func (c *DatabricksClient) configureWithDatabricksCfg(ctx context.Context) (func(r *http.Request) error, error) {
	configFile := c.ConfigFile
	if configFile == "" {
		configFile = "~/.databrickscfg"
	}
	configFile, err := homedir.Expand(configFile)
	if err != nil {
		return nil, fmt.Errorf("cannot find homedir: %w", err)
	}
	_, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		// early return for non-configured machines
		log.Printf("[DEBUG] %s not found on current host", configFile)
		return nil, nil
	}
	cfg, err := ini.Load(configFile)
	if err != nil {
		return nil, fmt.Errorf("cannot parse config file: %w", err)
	}
	if c.Profile == "" {
		log.Printf("[INFO] Using DEFAULT profile from %s", configFile)
		c.Profile = "DEFAULT"
	}
	dbcli := cfg.Section(c.Profile)
	if len(dbcli.Keys()) == 0 {
		// here we meet a heavy user of Databricks CLI
		return nil, fmt.Errorf("%s has no %s profile configured", configFile, c.Profile)
	}
	c.Host = dbcli.Key("host").String()
	if c.Host == "" {
		return nil, fmt.Errorf("config file %s is corrupt: cannot find host in %s profile",
			configFile, c.Profile)
	}
	token := ""
	authType := "Bearer"
	if dbcli.HasKey("username") && dbcli.HasKey("password") {
		c.Username = dbcli.Key("username").String()
		c.Password = dbcli.Key("password").String()
		token = c.encodeBasicAuth(c.Username, c.Password)
		authType = "Basic"
	} else {
		c.Token = dbcli.Key("token").String()
		token = c.Token
	}
	if token == "" {
		return nil, fmt.Errorf("config file %s is corrupt: cannot find token in %s profile",
			configFile, c.Profile)
	}
	log.Printf("[INFO] Using %s authentication from ~/.databrickscfg", authType)
	return c.authorizer(authType, token), nil
}

func (c *DatabricksClient) authorizer(authType, token string) func(r *http.Request) error {
	return func(r *http.Request) error {
		r.Header.Set("Authorization", fmt.Sprintf("%s %s", authType, token))
		return nil
	}
}

func (c *DatabricksClient) encodeBasicAuth(username, password string) string {
	tokenUnB64 := fmt.Sprintf("%s:%s", username, password)
	return base64.StdEncoding.EncodeToString([]byte(tokenUnB64))
}

func (c *DatabricksClient) configureHTTPCLient() {
	if c.HTTPTimeoutSeconds == 0 {
		c.HTTPTimeoutSeconds = DefaultHTTPTimeoutSeconds
	}
	if c.RateLimitPerSecond == 0 {
		c.RateLimitPerSecond = DefaultRateLimitPerSecond
	}
	c.rateLimiter = rate.NewLimiter(rate.Limit(c.RateLimitPerSecond), 1)
	// Set up a retryable HTTP Client to handle cases where the service returns
	// a transient error on initial creation
	retryDelayDuration := 10 * time.Second
	retryMaximumDuration := 5 * time.Minute
	defaultTransport := http.DefaultTransport.(*http.Transport)
	c.httpClient = &retryablehttp.Client{
		HTTPClient: &http.Client{
			Timeout: time.Duration(c.HTTPTimeoutSeconds) * time.Second,
			Transport: &http.Transport{
				Proxy:                 defaultTransport.Proxy,
				DialContext:           defaultTransport.DialContext,
				MaxIdleConns:          defaultTransport.MaxIdleConns,
				IdleConnTimeout:       defaultTransport.IdleConnTimeout * 3,
				TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout * 3,
				ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: c.InsecureSkipVerify,
				},
			},
		},
		CheckRetry: c.checkHTTPRetry,
		// Using a linear retry rather than the default exponential retry
		// as the creation condition is normally passed after 30-40 seconds
		// Setting the retry interval to 10 seconds. Setting RetryWaitMin and RetryWaitMax
		// to the same value removes jitter (which would be useful in a high-volume traffic scenario
		// but wouldn't add much here)
		Backoff:      retryablehttp.LinearJitterBackoff,
		RetryWaitMin: retryDelayDuration,
		RetryWaitMax: retryDelayDuration,
		RetryMax:     int(retryMaximumDuration / retryDelayDuration),
	}
}

// IsAzure returns true if client is configured for Azure Databricks - either by using AAD auth or with host+token combination
func (c *DatabricksClient) IsAzure() bool {
	return c.AzureResourceID != "" || c.AzureClientID != "" || c.AzureUseMSI || strings.Contains(c.Host, ".azuredatabricks.net")
}

// IsAws returns true if client is configured for AWS
func (c *DatabricksClient) IsAws() bool {
	return !c.IsAzure() && !c.IsGcp()
}

// IsGcp returns true if client is configured for GCP
func (c *DatabricksClient) IsGcp() bool {
	return c.GoogleServiceAccount != "" || strings.Contains(c.Host, ".gcp.databricks.com")
}

// FormatURL creates URL from the client Host and additional strings
func (c *DatabricksClient) FormatURL(strs ...string) string {
	host := c.Host
	if !strings.HasSuffix(host, "/") {
		host += "/"
	}
	data := append([]string{host}, strs...)
	return strings.Join(data, "")
}

// ClientForHost creates a new DatabricksClient instance with the same auth parameters,
// but for the given host. Authentication has to be reinitialized, as Google OIDC has
// different authorizers, depending if it's workspace or Accounts API we're talking to.
func (c *DatabricksClient) ClientForHost(ctx context.Context, url string) (*DatabricksClient, error) {
	log.Printf("[INFO] Creating client for host %s based on %s", url, c.configDebugString()) // lgtm[go/clear-text-logging]
	// Ensure that client is authenticated
	err := c.Authenticate(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot authenticate parent client: %w", err)
	}
	// copy all client configuration options except Databricks CLI profile
	return &DatabricksClient{
		Host:                 url,
		Username:             c.Username,
		Password:             c.Password,
		Token:                c.Token,
		ClientID:             c.ClientID,
		ClientSecret:         c.ClientSecret,
		GoogleServiceAccount: c.GoogleServiceAccount,
		GoogleCredentials:    c.GoogleCredentials,
		AzurermEnvironment:   c.AzurermEnvironment,
		InsecureSkipVerify:   c.InsecureSkipVerify,
		HTTPTimeoutSeconds:   c.HTTPTimeoutSeconds,
		DebugTruncateBytes:   c.DebugTruncateBytes,
		DebugHeaders:         c.DebugHeaders,
		RateLimitPerSecond:   c.RateLimitPerSecond,
		Provider:             c.Provider,
		rateLimiter:          c.rateLimiter,
		httpClient:           c.httpClient,
		configAttributesUsed: c.configAttributesUsed,
		commandFactory:       c.commandFactory,
	}, nil
}
