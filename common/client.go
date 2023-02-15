package common

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	*client.DatabricksClient

	Config *config.Config

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

	// Terraform provider instance to include Terraform binary version in
	// User-Agent header
	Provider *schema.Provider

	// callback used to create API1.2 call wrapper, which simplifies unit tessting
	commandFactory func(context.Context, *DatabricksClient) CommandExecutor
}

func (c *DatabricksClient) WorkspaceClient() (*databricks.WorkspaceClient, error) {
	return databricks.NewWorkspaceClient((*databricks.Config)(c.DatabricksClient.Config))
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
	// create dummy http request
	req, _ := http.NewRequestWithContext(ctx, "GET", "/", nil)
	// Ensure that client is authenticated
	err := c.DatabricksClient.Config.Authenticate(req)
	if err != nil {
		return nil, fmt.Errorf("cannot authenticate parent client: %w", err)
	}
	cfg := &config.Config{
		Host:                 url,
		Username:             c.Config.Username,
		Password:             c.Config.Password,
		Token:                c.Config.Token,
		ClientID:             c.Config.ClientID,
		ClientSecret:         c.Config.ClientSecret,
		GoogleServiceAccount: c.Config.GoogleServiceAccount,
		GoogleCredentials:    c.Config.GoogleCredentials,
		InsecureSkipVerify:   c.Config.InsecureSkipVerify,
		HTTPTimeoutSeconds:   c.Config.HTTPTimeoutSeconds,
		DebugTruncateBytes:   c.Config.DebugTruncateBytes,
		DebugHeaders:         c.Config.DebugHeaders,
		RateLimitPerSecond:   c.Config.RateLimitPerSecond,
	}
	client, err := client.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot configure new client: %w", err)
	}
	// copy all client configuration options except Databricks CLI profile
	return &DatabricksClient{
		DatabricksClient: client,
		Provider:         c.Provider,
		commandFactory:   c.commandFactory,
	}, nil
}
