package service

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
)

// DatabricksClient is the client struct that contains clients for all the services available on Databricks
type DatabricksClient struct {
	Host               string
	Token              string
	Username           string
	Password           string
	Profile            string
	ConfigFile         string
	AzureAuth          AzureAuth
	InsecureSkipVerify bool
	TimeoutSeconds     int
	userAgent          string
	httpClient         *retryablehttp.Client
	authVisitor        func(r *http.Request) error
}

// Configure ...
func (c *DatabricksClient) Configure(version string) error {
	c.configureHTTPCLient()
	c.AzureAuth.databricksClient = c
	c.userAgent = fmt.Sprintf("databricks-tf-provider/%s", version)
	err := c.findAndApplyAuthorizer()
	if err != nil {
		return err
	}
	if c.Host != "" && !(strings.HasPrefix(c.Host, "https://") || strings.HasPrefix(c.Host, "http://")) {
		// azurerm_databricks_workspace.*.workspace_url is giving URL without scheme
		// so that is why this line is here
		c.Host = "https://" + c.Host
	}
	return nil
}

func (c *DatabricksClient) findAndApplyAuthorizer() error {
	for _, authProvider := range []func() (func(r *http.Request) error, error){
		c.configureAuthWithDirectParams,
		c.AzureAuth.configureWithClientSecret,
		c.AzureAuth.configureWithAzureCLI,
		c.configureFromDatabricksCfg,
	} {
		authorizer, err := authProvider()
		if err != nil {
			return err
		}
		if authorizer == nil {
			continue
		}
		c.authVisitor = authorizer
		return nil
	}
	return fmt.Errorf("Authentication is not configured for provider. Please configure it\n" +
		"through one of the following options:\n" +
		"1. DATABRICKS_HOST + DATABRICKS_TOKEN environment variables.\n" +
		"2. host + token provider arguments.\n" +
		"3. azure_databricks_workspace_id + AZ CLI authentication.\n" +
		"4. azure_databricks_workspace_id + azure_client_id + azure_client_secret + azure_tenant_id " +
		"for Azure Service Principal authentication.\n" +
		"5. Run `databricks configure --token` that will create ~/.databrickscfg file.\n\n" +
		"Please check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for details")
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
		return nil, fmt.Errorf("Host is empty, but is required by %s", needsHostBecause)
	}
	if c.Token == "" || c.Host == "" {
		log.Printf("[INFO] No direct authentication params configured")
		return nil, nil
	}
	log.Printf("[INFO] Successfully configured Bearer authentication")
	return c.authorizer(authType, c.Token), nil
}

func (c *DatabricksClient) configureFromDatabricksCfg() (func(r *http.Request) error, error) {
	_, err := os.Stat(c.ConfigFile)
	if os.IsNotExist(err) {
		log.Printf("[INFO] ~/.databrickscfg not found on current host")
		// early return for non-configured machines
		return nil, nil
	}
	configFile, err := homedir.Expand(c.ConfigFile)
	if err != nil {
		return nil, err
	}
	cfg, err := ini.Load(configFile)
	if err != nil {
		return nil, err
	}
	if c.Profile == "" {
		log.Printf("[INFO] Using DEFAULT profile from ~/.databrickscfg")
		c.Profile = "DEFAULT"
	}
	dbcli := cfg.Section(c.Profile)
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
	log.Printf("[INFO] Successfully configured %s authentication from ~/.databrickscfg", authType)
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
	if c.TimeoutSeconds == 0 {
		c.TimeoutSeconds = 60
	}
	// Set up a retryable HTTP Client to handle cases where the service returns
	// a transient error on initial creation
	retryDelayDuration := 10 * time.Second
	retryMaximumDuration := 5 * time.Minute
	defaultTransport := http.DefaultTransport.(*http.Transport)
	c.httpClient = &retryablehttp.Client{
		HTTPClient: &http.Client{
			Timeout: time.Duration(c.TimeoutSeconds) * time.Second,
			Transport: &http.Transport{
				Proxy:                 defaultTransport.Proxy,
				DialContext:           defaultTransport.DialContext,
				MaxIdleConns:          defaultTransport.MaxIdleConns,
				IdleConnTimeout:       defaultTransport.IdleConnTimeout,
				TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
				ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: c.InsecureSkipVerify,
				},
			},
		},
		CheckRetry: checkHTTPRetry,
		// Using a linear retry rather than the default exponential retry
		// as the creation condition is normally passed after 30-40 seconds
		// Setting the retry interval to 10 seconds. Setting RetryWaitMin and RetryWaitMax
		// to the same value removes jitter (which would be useful in a high-volume traffic scenario
		// but wouldn't add much here)
		Backoff:      retryablehttp.LinearJitterBackoff,
		RetryWaitMin: retryDelayDuration,
		RetryWaitMax: retryDelayDuration,
		RetryMax:     int(retryMaximumDuration / retryDelayDuration), // + request & response log hooks
	}
}

// NewClientFromEnvironment makes very good client for testing purposes
func NewClientFromEnvironment() *DatabricksClient {
	client := DatabricksClient{
		Host:       os.Getenv("DATABRICKS_HOST"),
		Token:      os.Getenv("DATABRICKS_TOKEN"),
		Username:   os.Getenv("DATABRICKS_USERNAME"),
		Password:   os.Getenv("DATABRICKS_PASSWORD"),
		ConfigFile: os.Getenv("DATABRICKS_CONFIG_FILE"),
		Profile:    os.Getenv("DATABRICKS_CONFIG_PROFILE"),
		AzureAuth: AzureAuth{
			ResourceID:     os.Getenv("DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID"),
			WorkspaceName:  os.Getenv("DATABRICKS_AZURE_WORKSPACE_NAME"),
			ResourceGroup:  os.Getenv("DATABRICKS_AZURE_RESOURCE_GROUP"),
			SubscriptionID: os.Getenv("DATABRICKS_AZURE_SUBSCRIPTION_ID"),
			ClientID:       os.Getenv("DATABRICKS_AZURE_CLIENT_ID"),
			ClientSecret:   os.Getenv("DATABRICKS_AZURE_CLIENT_SECRET"),
			TenantID:       os.Getenv("DATABRICKS_AZURE_TENANT_ID"),
		},
	}
	err := client.Configure("dev")
	if err != nil {
		panic(err)
	}
	return &client
}

// IsAzure returns true if Azure is configured
func (c *DatabricksClient) IsAzure() bool {
	return c.AzureAuth.resourceID() != ""
}

// Clusters returns an instance of ClustersAPI
func (c *DatabricksClient) Clusters() ClustersAPI {
	return ClustersAPI{client: c}
}

// ClusterPolicies returns an instance of ClusterPoliciesAPI
func (c *DatabricksClient) ClusterPolicies() ClusterPoliciesAPI {
	return ClusterPoliciesAPI{client: c}
}

// Secrets returns an instance of SecretsAPI
func (c *DatabricksClient) Secrets() SecretsAPI {
	return SecretsAPI{client: c}
}

// SecretScopes returns an instance of SecretScopesAPI
func (c *DatabricksClient) SecretScopes() SecretScopesAPI {
	return SecretScopesAPI{client: c}
}

// SecretAcls returns an instance of SecretAclsAPI
func (c *DatabricksClient) SecretAcls() SecretAclsAPI {
	return SecretAclsAPI{client: c}
}

// Tokens returns an instance of TokensAPI
func (c *DatabricksClient) Tokens() TokensAPI {
	return TokensAPI{client: c}
}

// Users returns an instance of UsersAPI
func (c *DatabricksClient) Users() UsersAPI {
	return UsersAPI{client: c}
}

// Groups returns an instance of GroupsAPI
func (c *DatabricksClient) Groups() GroupsAPI {
	return GroupsAPI{client: c}
}

// Notebooks returns an instance of NotebooksAPI
func (c *DatabricksClient) Notebooks() NotebooksAPI {
	return NotebooksAPI{client: c}
}

// Jobs returns an instance of JobsAPI
func (c *DatabricksClient) Jobs() JobsAPI {
	return JobsAPI{client: c}
}

// DBFS returns an instance of DBFSAPI
func (c *DatabricksClient) DBFS() DBFSAPI {
	return DBFSAPI{client: c}
}

// Libraries returns an instance of LibrariesAPI
func (c *DatabricksClient) Libraries() LibrariesAPI {
	return LibrariesAPI{client: c}
}

// InstancePools returns an instance of InstancePoolsAPI
func (c *DatabricksClient) InstancePools() InstancePoolsAPI {
	return InstancePoolsAPI{client: c}
}

// InstanceProfiles returns an instance of InstanceProfilesAPI
func (c *DatabricksClient) InstanceProfiles() InstanceProfilesAPI {
	return InstanceProfilesAPI{client: c}
}

// Commands returns an instance of CommandsAPI
func (c *DatabricksClient) Commands() CommandsAPI {
	return CommandsAPI{client: c}
}

// MWSCredentials returns an instance of MWSCredentialsAPI
func (c *DatabricksClient) MWSCredentials() MWSCredentialsAPI {
	return MWSCredentialsAPI{client: c}
}

// MWSStorageConfigurations returns an instance of MWSStorageConfigurationsAPI
func (c *DatabricksClient) MWSStorageConfigurations() MWSStorageConfigurationsAPI {
	return MWSStorageConfigurationsAPI{client: c}
}

// MWSWorkspaces returns an instance of MWSWorkspacesAPI
func (c *DatabricksClient) MWSWorkspaces() MWSWorkspacesAPI {
	return MWSWorkspacesAPI{client: c}
}

// MWSNetworks returns an instance of MWSNetworksAPI
func (c *DatabricksClient) MWSNetworks() MWSNetworksAPI {
	return MWSNetworksAPI{client: c}
}

// MWSCustomerManagedKeys returns an instance of MWSCustomerManagedKeysAPI
func (c *DatabricksClient) MWSCustomerManagedKeys() MWSCustomerManagedKeysAPI {
	return MWSCustomerManagedKeysAPI{client: c}
}

// Permissions returns an instance of CommandsAPI
func (c *DatabricksClient) Permissions() PermissionsAPI {
	return PermissionsAPI{client: c}
}
