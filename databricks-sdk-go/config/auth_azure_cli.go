package config

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/oauth2"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/authconv"
	"github.com/databricks/databricks-sdk-go/logger"
)

// The header used to pass the service management token to the Databricks backend.
const xDatabricksAzureSpManagementToken = "X-Databricks-Azure-SP-Management-Token"

// The header used to pass the workspace resource ID to the Databricks backend.
const xDatabricksAzureWorkspaceResourceId = "X-Databricks-Azure-Workspace-Resource-Id"

type AzureCliCredentials struct {
}

func (c AzureCliCredentials) Name() string {
	return "azure-cli"
}

// implementing azureHostResolver for ensureWorkspaceUrl to work
func (c AzureCliCredentials) tokenSourceFor(_ context.Context, cfg *Config, _, resource string) auth.TokenSource {
	return &azureCliTokenSource{
		resource:      resource,
		azureTenantId: cfg.AzureTenantID,
	}
}

// There are three scenarios:
//
//  1. The user has logged in with the Azure CLI as a user and has access to the service management endpoint.
//  2. The user has logged in with the Azure CLI as a user and does not have access to the service management endpoint.
//  3. The user has logged in with the Azure CLI as a service principal, and must have access to the service management
//     endpoint to authenticate.
//
// If the user can't access the service management endpoint, we assume they are in case 2 and do not pass the service
// management token. Otherwise, we always pass the service management token.
func (c AzureCliCredentials) getVisitor(ctx context.Context, cfg *Config, inner auth.TokenSource) (func(*http.Request) error, error) {
	opts := cacheOptions(cfg)
	ts := &azureCliTokenSource{
		resource:            cfg.Environment().AzureServiceManagementEndpoint(),
		workspaceResourceId: cfg.AzureResourceID,
		azureTenantId:       cfg.AzureTenantID,
	}
	t, err := ts.Token(ctx)
	if err != nil {
		logger.Debugf(ctx, "Not including service management token in headers: %v", err)
		return azureVisitor(cfg, refreshableAuthVisitor(inner, opts...)), nil
	}
	management := azureReuseTokenSource(t, ts, opts...)
	return azureVisitor(cfg, serviceToServiceVisitor(inner, management, xDatabricksAzureSpManagementToken, false, opts...)), nil
}

func (c AzureCliCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if err := unsupportedGroupRoleAssumption(cfg, c.Name()); err != nil {
		return nil, err
	}
	if !cfg.IsAzure() {
		return nil, nil
	}
	// Set the azure tenant ID from host if available
	err := cfg.loadAzureTenantId(ctx)
	if err != nil {
		return nil, fmt.Errorf("load tenant id: %w", err)
	}
	// Eagerly get a token to fail fast in case the user is not logged in with the Azure CLI.
	ts := &azureCliTokenSource{
		resource:            cfg.Environment().AzureApplicationID,
		workspaceResourceId: cfg.AzureResourceID,
		azureTenantId:       cfg.AzureTenantID,
	}
	t, err := ts.Token(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "No subscription found") {
			// auth is not configured
			return nil, nil
		}
		if strings.Contains(err.Error(), "executable file not found") {
			logger.Debugf(ctx, "Most likely Azure CLI is not installed. "+
				"See https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest for details")
			return nil, nil
		}
		return nil, err
	}
	err = cfg.azureEnsureWorkspaceUrl(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("resolve host: %w", err)
	}
	visitor, err := c.getVisitor(ctx, cfg, azureReuseTokenSource(t, ts, cacheOptions(cfg)...))
	if err != nil {
		return nil, err
	}
	logger.Infof(ctx, "Using Azure CLI authentication with AAD tokens")
	return newVisitorOAuthCredentials(visitor, ts), nil
}

// NewAzureCliTokenSource returns [oauth2.TokenSource] for a passwordless authentication via Azure CLI (`az login`)
func NewAzureCliTokenSource(ctx context.Context, resource, azureTenantId string) oauth2.TokenSource {
	return authconv.OAuth2TokenSource(&azureCliTokenSource{
		resource:      resource,
		azureTenantId: azureTenantId,
	})
}

type azureCliTokenSource struct {
	resource            string
	workspaceResourceId string
	azureTenantId       string
}

type internalCliToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
	ExpiresOn    string `json:"expiresOn"`
}

func (ts *azureCliTokenSource) getSubscription(ctx context.Context) string {
	if ts.workspaceResourceId == "" {
		return ""
	}
	components := strings.Split(ts.workspaceResourceId, "/")
	if len(components) < 3 {
		logger.Warnf(ctx, "Invalid azure workspace resource ID")
		return ""
	}
	return components[2]
}

func (ts *azureCliTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	tokenBytes, err := ts.getTokenBytes(ctx)
	if err != nil {
		return nil, err
	}
	var it internalCliToken
	err = json.Unmarshal(tokenBytes, &it)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal CLI result: %w", err)
	}
	expiresOn, err := time.ParseInLocation("2006-01-02 15:04:05.999999", it.ExpiresOn, time.Local)
	if err != nil {
		return nil, fmt.Errorf("cannot parse expiry: %w", err)
	}
	tenantIdDebug := ""
	if ts.azureTenantId != "" {
		tenantIdDebug = fmt.Sprintf(" for tenant %s", ts.azureTenantId)
	}
	logger.Infof(ctx, "Refreshed OAuth token for %s%s from Azure CLI, which expires on %s",
		ts.resource, tenantIdDebug, it.ExpiresOn)

	var extra map[string]interface{}
	err = json.Unmarshal(tokenBytes, &extra)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal extra: %w", err)
	}
	return (&oauth2.Token{
		AccessToken:  it.AccessToken,
		RefreshToken: it.RefreshToken,
		TokenType:    it.TokenType,
		Expiry:       expiresOn,
	}).WithExtra(extra), nil
}

func (ts *azureCliTokenSource) getTokenBytes(ctx context.Context) ([]byte, error) {
	// When fetching an access token from the CLI with a managed identity, the tenant ID should not be specified.
	// https://github.com/hashicorp/go-azure-sdk/pull/910/files demonstrates how to detect whether the CLI is authenticated
	// using a managed identity or federated token.
	accountRaw, err := runCommand(ctx, "az", []string{"account", "show", "--output", "json"})
	if err != nil {
		return nil, fmt.Errorf("cannot get account info: %w", err)
	}
	var account struct {
		User struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"user"`
	}
	if err := json.Unmarshal(accountRaw, &account); err != nil {
		return nil, fmt.Errorf("cannot unmarshal account info: %w", err)
	}

	isMsi := account.User.Type == "servicePrincipal" && (account.User.Name == "systemAssignedIdentity" || account.User.Name == "userAssignedIdentity")
	if !isMsi {
		// For regular service principals, try with tenant ID first
		result, err := ts.getTokenBytesWithTenantId(ctx, ts.azureTenantId)
		if err != nil && account.User.Type == "servicePrincipal" {
			// If it fails for service principals, it might be a federated token scenario
			// where tenant ID should not be specified. Fall back to no tenant ID.
			logger.Infof(ctx, "Failed to get token with tenant ID for service principal, trying without tenant ID: %v", err)
			return ts.getTokenBytesWithTenantId(ctx, "")
		}
		return result, err
	}
	return ts.getTokenBytesWithTenantId(ctx, "")
}

func (ts *azureCliTokenSource) getTokenBytesWithTenantId(ctx context.Context, tenantId string) ([]byte, error) {
	args := []string{"account", "get-access-token", "--resource",
		ts.resource, "--output", "json"}
	if tenantId != "" {
		args = append(args, "--tenant", tenantId)
	}
	subscription := ts.getSubscription(ctx)
	if subscription != "" && ts.azureTenantId == "" {
		// This will fail if the user has access to the workspace, but not to the subscription
		// itself.
		// In such case, we fall back to not using the subscription.
		// This should only be attempted when the tenant ID is not known.
		result, err := runCommand(ctx, "az", append(args, "--subscription", subscription))
		if err == nil {
			return result, nil
		}
		logger.Infof(ctx, "Failed to get token for subscription. Using resource only token.")
	}
	result, err := runCommand(ctx, "az", args)
	if ee, ok := err.(*exec.ExitError); ok {
		return nil, fmt.Errorf("cannot get access token: %s", string(ee.Stderr))
	}
	if err != nil {
		return nil, fmt.Errorf("cannot get access token: %w", err)
	}
	return result, nil
}
