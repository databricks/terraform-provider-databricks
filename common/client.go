package common

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/api/option"

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

type DatabricksClient struct {
	Host       string
	Token      string
	Username   string
	Password   string
	Profile    string
	ConfigFile string
	AccountID  string

	GoogleServiceAccount string

	AzureWorkspaceName        string
	AzureResourceGroup        string
	AzureSubscriptionID       string
	AzureDatabricksResourceID string
	AzureClientSecret         string
	AzureClientID             string
	AzureTenantID             string
	AzurermEnvironment        string

	// temporary workaround to use PAT tokens instead of AAD tokens
	AzurePATTokenDurationSeconds string
	AzureUsePATForCLI            bool
	AzureUsePATForSPN            bool

	AzureEnvironment *azure.Environment

	InsecureSkipVerify bool
	HTTPTimeoutSeconds int
	DebugTruncateBytes int
	DebugHeaders       bool
	RateLimitPerSecond int

	// OAuth token refreshers for Azure to be used within `authVisitor`
	azureAuthorizer autorest.Authorizer

	// session temporary PAT token if `UsePATForSPN` or `UsePATForCLI` are true
	temporaryPat *tokenResponse

	// options used to enable unit testing mode for OIDC
	googleAuthOptions []option.ClientOption

	// Context used during provider initialisation,
	// mostly for OAuth-based validation.
	InitContext context.Context

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

type ConfigOption struct {
	Kind      string
	Name      string
	Sensitive bool
	EnvVars   []string
}

var configOptions = []ConfigOption{
	{"direct", "host", false, []string{"DATABRICKS_HOST"}}, // aws, azure, gcp
	{"host", "token", true, []string{"DATABRICKS_TOKEN"}},
	{"host", "username", false, []string{"DATABRICKS_USERNAME"}},
	{"host", "password", true, []string{"DATABRICKS_PASSWORD"}},
	{"direct", "config_file", false, []string{"DATABRICKS_CONFIG_FILE"}},
	{"config_file", "profile", false, []string{"DATABRICKS_CONFIG_PROFILE"}},
	{"direct", "azure_workspace_resource_id", false, []string{"DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID", "AZURE_DATABRICKS_WORKSPACE_RESOURCE_ID"}},
	{"direct", "azure_subscription_id", false, []string{"DATABRICKS_AZURE_SUBSCRIPTION_ID", "ARM_SUBSCRIPTION_ID"}},
	{"direct", "azure_resource_group", false, []string{"DATABRICKS_AZURE_RESOURCE_GROUP"}},
	{"direct", "azure_workspace_name", false, []string{"DATABRICKS_AZURE_WORKSPACE_NAME"}},
	{"direct", "azure_client_id", false, []string{"DATABRICKS_AZURE_CLIENT_ID", "ARM_CLIENT_ID"}},
	{"direct", "azure_client_secret", true, []string{"DATABRICKS_AZURE_CLIENT_SECRET", "ARM_CLIENT_SECRET"}},
	{"direct", "azure_tenant_id", false, []string{"DATABRICKS_AZURE_TENANT_ID", "ARM_TENANT_ID"}},
	{"direct", "azure_environment", false, []string{"ARM_ENVIRONMENT"}},
	{"direct", "google_service_account", false, []string{"DATABRICKS_GOOGLE_SERVICE_ACCOUNT"}},
}

// Configure client to work, optionally specifying configuration attributes used
func (c *DatabricksClient) Configure(attrsUsed ...string) error {
	c.configAttributesUsed = attrsUsed
	c.configureHTTPCLient()
	if c.DebugTruncateBytes == 0 {
		c.DebugTruncateBytes = DefaultTruncateBytes
	}
	return nil
}

// Authenticate authenticates across providers or returns error
func (c *DatabricksClient) Authenticate() error {
	if c.authVisitor != nil {
		return nil
	}
	c.authMutex.Lock()
	defer c.authMutex.Unlock()
	if c.authVisitor != nil {
		return nil
	}
	authorizers := []func() (func(r *http.Request) error, error){
		c.configureAuthWithDirectParams,
		c.configureWithClientSecret,
		c.configureWithAzureCLI,
		c.configureWithGoogleForAccountsAPI,
		c.configureWithGoogleForWorkspace,
		c.configureFromDatabricksCfg,
	}
	for _, authProvider := range authorizers {
		authorizer, err := authProvider()
		if err != nil {
			return fmt.Errorf("cannot configure auth: %w", err)
		}
		if authorizer == nil {
			continue
		}
		c.authVisitor = authorizer
		c.fixHost()
		return nil
	}
	info := ""
	if len(c.configAttributesUsed) > 0 {
		// TODO: add env vars
		info = fmt.Sprintf("Attributes used: %s - %d", strings.Join(c.configAttributesUsed, ", "), len(configOptions))
	}
	docUrl := "https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs#authentication"
	return fmt.Errorf("authentication is not configured for provider. %sPlease check %s for details", info, docUrl)
}

func (c *DatabricksClient) fixHost() {
	if c.Host != "" && !(strings.HasPrefix(c.Host, "https://") || strings.HasPrefix(c.Host, "http://")) {
		// azurerm_databricks_workspace.*.workspace_url is giving URL without scheme
		// so that is why this line is here
		c.Host = "https://" + c.Host
	}
}

func (c *DatabricksClient) configureAuthWithDirectParams() (func(r *http.Request) error, error) {
	authType := "Bearer"
	var needsHostBecause string
	if c.Username != "" && c.Password != "" {
		authType = "Basic"
		needsHostBecause = "basic_auth"
		c.Token = c.encodeBasicAuth(c.Username, c.Password)
		c.Password = ""
		log.Printf("[INFO] Using basic auth for user '%s'", c.Username)
	} else if c.Token != "" {
		needsHostBecause = "token"
	}
	if needsHostBecause != "" && c.Host == "" {
		return nil, fmt.Errorf("host is empty, but is required by %s", needsHostBecause)
	}
	if c.Token == "" || c.Host == "" {
		return nil, nil
	}
	log.Printf("[INFO] Using directly configured host+%s authentication", needsHostBecause)
	return c.authorizer(authType, c.Token), nil
}

func (c *DatabricksClient) configureFromDatabricksCfg() (func(r *http.Request) error, error) {
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
		log.Printf("[INFO] ~/.databrickscfg not found on current host")
		// early return for non-configured machines
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
	authType := "Bearer"
	if dbcli.HasKey("username") && dbcli.HasKey("password") {
		username := dbcli.Key("username").String()
		password := dbcli.Key("password").String()
		c.Token = c.encodeBasicAuth(username, password)
		authType = "Basic"
	} else {
		c.Token = dbcli.Key("token").String()
	}
	if c.Token == "" {
		return nil, fmt.Errorf("config file %s is corrupt: cannot find token in %s profile",
			configFile, c.Profile)
	}
	log.Printf("[INFO] Using %s authentication from ~/.databrickscfg", authType)
	return c.authorizer(authType, c.Token), nil
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
	if c.InitContext == nil {
		c.InitContext = context.Background()
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
	return c.resourceID() != "" || strings.Contains(c.Host, ".azuredatabricks.net")
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
