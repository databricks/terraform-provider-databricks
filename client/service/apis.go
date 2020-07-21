package service

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
)

// DatabricksClient is the client struct that contains clients for all the services available on Databricks
type DatabricksClient struct {
	Host       string
	Token      string
	Profile    string
	ConfigFile string
	BasicAuth  struct {
		Username string
		Password string
	}
	AzureAuth          AzureAuth
	InsecureSkipVerify bool
	TimeoutSeconds     int
	tokenCreateTime    int64
	tokenExpiryTime    int64
	authType           string
	userAgent          string
	httpClient         *retryablehttp.Client
	uriPrefix          string
	auth               func(r *http.Request) (*http.Request, error)
}

// Configure validates and configures the client
func (c *DatabricksClient) ConfigureOld(version string) error {
	c.userAgent = fmt.Sprintf("databricks-tf-provider/%s", version)

	c.configureHTTPCLient()
	c.AzureAuth.databricksClient = c

	var hasCredentials bool
	for _, authProvider := range []func() (bool, error){
		c.configureAuthWithDirectParams,
		c.AzureAuth.configureWithClientSecret,
		c.AzureAuth.configureWithAzureCLI,
		c.configureFromDatabricksCfg} {
		success, err := authProvider()
		if success {
			hasCredentials = true
			break
		}
		if err != nil {
			return err
		}
	}
	if !hasCredentials {
		return fmt.Errorf("Authentication is not configured for provider. Please configure it\n" +
			"through one of the following options:\n" +
			"1. DATABRICKS_HOST + DATABRICKS_TOKEN environment variables.\n" +
			"2. host + token provider arguments.\n" +
			"3. azure_auth configuration block.\n" +
			"4. Run `databricks configure --token` that will create ~/.databrickscfg file.\n\n" +
			"Please check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for details")
	}
	if c.authType == "" {
		c.authType = "Bearer"
	}
	parsedURI, err := url.Parse(c.Host)
	if err != nil {
		return err
	}
	c.uriPrefix = fmt.Sprintf("%s://%s/api", parsedURI.Scheme, parsedURI.Host)
	return nil
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
	// parsedURI, err := url.Parse(c.Host)
	// if err != nil {
	// 	return err
	// }
	// c.uriPrefix = fmt.Sprintf("%s://%s/api", parsedURI.Scheme, parsedURI.Host)
	return nil
}

func (c *DatabricksClient) findAndApplyAuthorizer() error {
	for _, authProvider := range []func() (func(r *http.Request) (*http.Request, error), error){
		c.configureAuthWithDirectParams2,
		c.AzureAuth.configureWithClientSecret2,
		c.configureFromDatabricksCfg2,
		// c.AzureAuth.configureWithAzureCLI,
	} {
		authorizer, err := authProvider()
		if err != nil {
			return err
		}
		if authorizer == nil {
			continue
		}
		c.auth = authorizer
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

func (c *DatabricksClient) configureAuthWithDirectParams2() (func(r *http.Request) (*http.Request, error), error) {
	authType := "Bearer"
	var needsHostBecause string
	username := c.BasicAuth.Username
	if username != "" && c.BasicAuth.Password != "" {
		authType = "Basic"
		c.BasicAuth.Password = ""
		needsHostBecause = "basic_auth"
		c.Token = c.encodeBasicAuth(username, c.BasicAuth.Password)
		log.Printf("[INFO] Using basic auth for user '%s'", username)
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

func (c *DatabricksClient) configureFromDatabricksCfg2() (func(r *http.Request) (*http.Request, error), error) {
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

func (c *DatabricksClient) authorizer(authType, token string) func(r *http.Request) (*http.Request, error) {
	return func(r *http.Request) (*http.Request, error) {
		r.Header.Set("Authorization", fmt.Sprintf("%s %s", authType, token))
		return r, nil
	}
}

func (c *DatabricksClient) configureAuthWithDirectParams() (bool, error) {
	var needsHostBecause string
	if c.BasicAuth.Username != "" && c.BasicAuth.Password != "" {
		c.Token = c.encodeBasicAuth(c.BasicAuth.Username, c.BasicAuth.Password)
		c.BasicAuth.Password = ""
		needsHostBecause = "basic_auth"
		c.authType = "Basic"
	} else if c.Token != "" {
		needsHostBecause = "token"
	}
	if needsHostBecause != "" && c.Host == "" {
		return false, fmt.Errorf("Host is empty, but is required by %s", needsHostBecause)
	}
	if c.Token == "" || c.Host == "" {
		// direct params are not configured
		return false, nil
	}
	return true, nil
}

func (c *DatabricksClient) encodeBasicAuth(username, password string) string {
	tokenUnB64 := fmt.Sprintf("%s:%s", username, password)
	return base64.StdEncoding.EncodeToString([]byte(tokenUnB64))
}

// configureFromDatabricksCfg sets Host and Token from ~/.databrickscfg file if it exists
func (c *DatabricksClient) configureFromDatabricksCfg() (bool, error) {
	_, err := os.Stat(c.ConfigFile)
	if os.IsNotExist(err) {
		// early return for non-configured machines
		return false, nil
	}
	configFile, err := homedir.Expand(c.ConfigFile)
	if err != nil {
		return false, err
	}
	cfg, err := ini.Load(configFile)
	if err != nil {
		return false, err
	}
	if c.Profile == "" {
		c.Profile = "DEFAULT"
	}
	dbcli := cfg.Section(c.Profile)
	c.Host = dbcli.Key("host").String()
	if c.Host == "" {
		return false, fmt.Errorf("config file %s is corrupt: cannot find host in %s profile",
			configFile, c.Profile)
	}
	if dbcli.HasKey("username") && dbcli.HasKey("password") {
		username := dbcli.Key("username").String()
		password := dbcli.Key("password").String()
		c.Token = c.encodeBasicAuth(username, password)
		c.authType = "Basic"
	} else {
		c.Token = dbcli.Key("token").String()
	}
	if c.Token == "" {
		return false, fmt.Errorf("config file %s is corrupt: cannot find token in %s profile",
			configFile, c.Profile)
	}
	return true, nil
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
				// TODO: This probably should be a configuration at the provider level and optional and not a fixed val
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
