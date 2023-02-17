package common

import (
	"context"
	"fmt"
	"net/http"
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

// IsAzure returns true if client is configured for Azure Databricks - either by using AAD auth or with host+token combination
func (c *DatabricksClient) IsAzure() bool {
	return c.Config.IsAzure()
}

// IsAws returns true if client is configured for AWS
func (c *DatabricksClient) IsAws() bool {
	return !c.IsGcp() && !c.IsAzure()
}

// IsGcp returns true if client is configured for GCP
func (c *DatabricksClient) IsGcp() bool {
	return c.Config.GoogleServiceAccount != "" || c.Config.IsGcp()
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
