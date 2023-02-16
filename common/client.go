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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DatabricksClient holds properties needed for authentication and HTTP client setup
// fields with `name` struct tags become Terraform provider attributes. `env` struct tag
// can hold one or more coma-separated env variable names to find value, if not specified
// directly. `auth` struct tag describes the type of conflicting authentication used.
type DatabricksClient struct {
	*client.DatabricksClient

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
		if attr.Internal && !c.Config.DebugHeaders {
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
	return c.Config.IsAzure()
}

// IsAws returns true if client is configured for AWS
func (c *DatabricksClient) IsAws() bool {
	return c.Config.IsAws()
}

// IsGcp returns true if client is configured for GCP
func (c *DatabricksClient) IsGcp() bool {
	return c.Config.IsGcp()
}

// FormatURL creates URL from the client Host and additional strings
func (c *DatabricksClient) FormatURL(strs ...string) string {
	host := c.Config.Host
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
