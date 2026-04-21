// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package common

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type cachedMe struct {
	internalImpl iam.CurrentUserService
	cachedUser   *iam.User
	mu           sync.Mutex
}

func newCachedMe(inner iam.CurrentUserService) *cachedMe {
	return &cachedMe{
		internalImpl: inner,
	}
}

func (a *cachedMe) Me(ctx context.Context) (*iam.User, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.cachedUser != nil {
		return a.cachedUser, nil
	}
	user, err := a.internalImpl.Me(ctx)
	if err != nil {
		return user, err
	}
	a.cachedUser = user
	return user, err
}

// DatabricksClient holds properties needed for authentication and HTTP client setup
// fields with `name` struct tags become Terraform provider attributes. `env` struct tag
// can hold one or more coma-separated env variable names to find value, if not specified
// directly. `auth` struct tag describes the type of conflicting authentication used.
type DatabricksClient struct {
	*client.DatabricksClient

	// callback used to create API1.2 call wrapper, which simplifies unit testing
	commandFactory func(context.Context, *DatabricksClient) CommandExecutor

	// cachedWorkspaceClient is a cached workspace client authenticated to the workspace
	// configured for the provider
	cachedWorkspaceClient *databricks.WorkspaceClient

	// cachedWorkspaceID is the cached workspace ID of the workspace client
	// authenticated to the workspace configured for the provider
	cachedWorkspaceID int64

	// cachedDatabricksClients is a map of Databricks Clients authenticated to the workspaces
	// configured for the provider. The key is the workspace ID. This is used by legacy SDKv2
	// resources and data sources not using Go SDK.
	cachedDatabricksClients map[int64]*client.DatabricksClient

	// cachedWorkspaceClients is a map of workspace clients for each workspace ID
	// populated when fetching a WorkspaceClient for a specific workspace ID using
	// a provider configured at the account level
	cachedWorkspaceClients map[int64]*databricks.WorkspaceClient

	// cachedAccountClient is a cached account client authenticated to the account
	// configured for the provider
	cachedAccountClient *databricks.AccountClient

	// mu synchronizes access to all cached clients.
	mu sync.Mutex

	// muLegacy synchronizes access to all cached clients.
	// This is used by legacy SDKv2 resources and data sources not using Go SDK where
	// a new client is created.
	muLegacy sync.Mutex
}

// GetWorkspaceClientForUnifiedProviderWithDiagnostics returns the Databricks
// WorkspaceClient for workspace level resources or diagnostics if that fails
// for terraform provider, the provider can be configured at account level or workspace level.
// This implementation will be used by resources and data sources that are developed
// over plugin framework.
func (c *DatabricksClient) GetWorkspaceClientForUnifiedProviderWithDiagnostics(
	ctx context.Context, workspaceID string,
) (*databricks.WorkspaceClient, diag.Diagnostics) {
	w, err := c.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
	if err != nil {
		return nil, diag.Diagnostics{diag.NewErrorDiagnostic("failed to get workspace client", err.Error())}
	}
	return w, nil
}

// GetWorkspaceClientForUnifiedProvider returns the Databricks
// WorkspaceClient for workspace level resources or diagnostics if that fails
// for terraform provider, the provider can be configured at account level or workspace level.
// This implementation will be used by resources and data sources that are developed
// over SDKv2.
func (c *DatabricksClient) GetWorkspaceClientForUnifiedProvider(
	ctx context.Context, workspaceID string,
) (*databricks.WorkspaceClient, error) {
	// The provider can be configured at account level or workspace level.
	if c.HostTypeForTerraform() != config.WorkspaceHost {
		return c.getWorkspaceClientForAccountUnifiedHost(ctx, workspaceID)
	}
	return c.getWorkspaceClientForWorkspaceConfiguredProvider(ctx, workspaceID)
}

// getWorkspaceClientForAccountUnifiedHost gets the workspace client for
// the workspace ID specified in the resource when the provider is configured
// at account level.
func (c *DatabricksClient) getWorkspaceClientForAccountUnifiedHost(
	ctx context.Context, workspaceID string,
) (*databricks.WorkspaceClient, error) {
	// If workspace_id is not provided in provider_config, use the provider-level
	// workspace_id from SDK config as fallback
	if workspaceID == "" {
		workspaceID = c.Config.WorkspaceID
	}
	if workspaceID == "" {
		return nil, fmt.Errorf("managing workspace-level resources requires a workspace_id, " +
			"but none was found in the resource's provider_config block or the provider's workspace_id attribute")
	}

	// Parse the workspace ID to int.
	workspaceIDInt, err := parseWorkspaceID(workspaceID)
	if err != nil {
		return nil, err
	}

	// Get the workspace client for the workspace ID.
	w, err := c.WorkspaceClientForWorkspace(ctx, workspaceIDInt)
	if err != nil {
		return nil, fmt.Errorf("failed to get workspace client with workspace_id %d: %w", workspaceIDInt, err)
	}
	return w, nil
}

// getWorkspaceClientForWorkspaceConfiguredProvider gets the workspace client for
// the workspace ID specified in the resource when the provider is configured at workspace level.
func (c *DatabricksClient) getWorkspaceClientForWorkspaceConfiguredProvider(
	ctx context.Context, workspaceID string,
) (*databricks.WorkspaceClient, error) {
	// Provider is configured at workspace level and we get the
	// workspace client from the provider.
	if workspaceID == "" {
		return c.WorkspaceClient()
	}

	workspaceIDInt, err := parseWorkspaceID(workspaceID)
	if err != nil {
		return nil, err
	}

	// Check if the workspace ID specified in the resource matches
	// the workspace ID of the provider configured workspace client.
	w, err := c.WorkspaceClient()
	if err != nil {
		return nil, err
	}

	err = c.validateWorkspaceIDFromProvider(ctx, workspaceIDInt, w)
	if err != nil {
		return nil, fmt.Errorf("failed to validate workspace_id: %w", err)
	}
	// The provider is configured at the workspace level and the
	// workspace ID matches
	return w, nil
}

// parseWorkspaceID parses the workspace ID from string to int64.
func parseWorkspaceID(workspaceID string) (int64, error) {
	workspaceIDInt, err := strconv.ParseInt(workspaceID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse workspace_id, please check if the workspace_id in provider_config is a valid integer: %w", err)

	}
	return workspaceIDInt, nil
}

// CurrentWorkspaceID returns the workspace ID for a workspace-level provider.
// It uses the cached value if available, otherwise makes an API call to resolve it.
func (c *DatabricksClient) CurrentWorkspaceID(ctx context.Context) (int64, error) {
	if c.cachedWorkspaceID != 0 {
		return c.cachedWorkspaceID, nil
	}
	w, err := c.WorkspaceClient()
	if err != nil {
		return 0, err
	}
	err = c.setCachedWorkspaceID(ctx, w)
	if err != nil {
		return 0, err
	}
	return c.cachedWorkspaceID, nil
}

// validateWorkspaceIDFromProvider validates the workspace ID specified in the
// resource or data soruce matches the workspace ID of the provider configured workspace client.
func (c *DatabricksClient) validateWorkspaceIDFromProvider(ctx context.Context, workspaceID int64,
	w *databricks.WorkspaceClient) error {
	// If the workspace ID is not cached, we make an API call to the workspace to get
	// the current workspace ID and cache it.
	if c.cachedWorkspaceID == 0 {
		err := c.setCachedWorkspaceID(ctx, w)
		if err != nil {
			return err
		}
	}

	if c.cachedWorkspaceID != workspaceID {
		return fmt.Errorf("workspace_id mismatch: provider is configured for workspace %d but got %d in provider_config. "+
			"please check the workspace_id provided in provider_config",
			c.cachedWorkspaceID, workspaceID)
	}
	return nil
}

// setCachedWorkspaceID sets the cached workspace ID.
func (c *DatabricksClient) setCachedWorkspaceID(ctx context.Context, w *databricks.WorkspaceClient) error {
	// Acquire the lock to avoid race conditions.
	c.mu.Lock()
	defer c.mu.Unlock()
	// Double checked locking
	if c.cachedWorkspaceID == 0 {
		id, err := w.CurrentWorkspaceID(ctx)
		if err != nil {
			return fmt.Errorf("failed to get the workspace_id: %w",
				err)
		}
		c.cachedWorkspaceID = id
	}
	return nil
}

// GetWorkspaceClient returns the Databricks WorkspaceClient or a diagnostics if
// that fails. This is used by resources and data sources that are developed
// over plugin framework.
func (c *DatabricksClient) GetWorkspaceClient() (*databricks.WorkspaceClient,
	diag.Diagnostics) {
	w, err := c.WorkspaceClient()
	if err != nil {
		return nil, diag.Diagnostics{diag.NewErrorDiagnostic("Failed to get workspace client", err.Error())}
	}
	return w, nil
}

func (c *DatabricksClient) WorkspaceClient() (*databricks.WorkspaceClient, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cachedWorkspaceClient != nil {
		return c.cachedWorkspaceClient, nil
	}
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(c.DatabricksClient.Config))
	if err != nil {
		return nil, err
	}
	w.CurrentUser = newCachedMe(w.CurrentUser)
	c.cachedWorkspaceClient = w
	return w, nil
}

// Set the cached workspace client.
func (c *DatabricksClient) SetWorkspaceClient(w *databricks.WorkspaceClient) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cachedWorkspaceClient = w
}

func (c *DatabricksClient) WorkspaceClientForWorkspace(ctx context.Context, workspaceId int64) (*databricks.WorkspaceClient, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cachedWorkspaceClients == nil {
		c.cachedWorkspaceClients = make(map[int64]*databricks.WorkspaceClient)
	}
	if client, ok := c.cachedWorkspaceClients[workspaceId]; ok {
		return client, nil
	}
	// Get workspace client via account API.
	w, err := c.workspaceClientViaAccountAPI(ctx, workspaceId)
	if err != nil {
		// Fallback: create workspace client on the same host with workspace_id set.
		// This works for unified hosts that can route by workspace_id through X-Databricks-Org-Id header.
		// Note: GetWorkspaceClient(*workspace) supports all host, this is the case when users don't have access to the account
		w, err = c.tryWorkspaceClientDirect(workspaceId)
		if err != nil {
			return nil, err
		}
	}
	w.CurrentUser = newCachedMe(w.CurrentUser)
	c.cachedWorkspaceClients[workspaceId] = w
	return w, nil
}

// workspaceClientViaAccountAPI resolves the workspace deployment URL via the
// account API and creates a workspace client pointing to it.
func (c *DatabricksClient) workspaceClientViaAccountAPI(ctx context.Context, workspaceId int64) (*databricks.WorkspaceClient, error) {
	a, err := c.accountClient()
	if err != nil {
		return nil, err
	}
	workspace, err := a.Workspaces.Get(ctx, provisioning.GetWorkspaceRequest{WorkspaceId: workspaceId})
	if err != nil {
		return nil, err
	}
	return a.GetWorkspaceClient(*workspace)
}

// tryWorkspaceClientDirect creates a workspace client on the same host with
// workspace_id set, then validates it with a lightweight API call. This works
// for unified hosts that can route requests by workspace_id.
func (c *DatabricksClient) tryWorkspaceClientDirect(workspaceId int64) (*databricks.WorkspaceClient, error) {
	cfg, err := c.Config.NewWithWorkspaceHost(c.Config.Host)
	if err != nil {
		return nil, err
	}
	cfg.AccountID = c.Config.AccountID
	cfg.WorkspaceID = fmt.Sprintf("%d", workspaceId)
	return databricks.NewWorkspaceClient((*databricks.Config)(cfg))
}

// SetWorkspaceClientForWorkspace sets the cached workspace client for a specific workspace ID.
func (c *DatabricksClient) SetWorkspaceClientForWorkspace(workspaceId int64, w *databricks.WorkspaceClient) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cachedWorkspaceClients == nil {
		c.cachedWorkspaceClients = make(map[int64]*databricks.WorkspaceClient)
	}
	c.cachedWorkspaceClients[workspaceId] = w
}

// Set the cached account client.
func (c *DatabricksClient) SetAccountClient(a *databricks.AccountClient) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cachedAccountClient = a
}

// GetProviderWorkspaceID returns the provider-level workspace_id from Config.
// Satisfies the tfschema.UnifiedProviderClient interface.
func (c *DatabricksClient) GetProviderWorkspaceID() string {
	return c.Config.WorkspaceID
}

// ValidateWorkspaceAccess validates that the workspace client for the given
// workspace_id is reachable. Satisfies the tfschema.UnifiedProviderClient interface.
func (c *DatabricksClient) ValidateWorkspaceAccess(ctx context.Context, workspaceID string) diag.Diagnostics {
	_, diags := c.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	return diags
}

// SetCachedWorkspaceID sets the cached workspace ID directly.
// This is used by test infrastructure to pre-populate the cache and prevent
// lazy CurrentWorkspaceID API calls during unit tests.
func (c *DatabricksClient) SetCachedWorkspaceID(id int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cachedWorkspaceID = id
}

func (c *DatabricksClient) setAccountId(accountId string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if accountId == "" {
		return nil
	}
	oldAccountID := c.DatabricksClient.Config.AccountID
	if oldAccountID != "" && oldAccountID != accountId {
		return fmt.Errorf("account ID is already set to %s", oldAccountID)
	}
	c.DatabricksClient.Config.AccountID = accountId
	return nil
}

// GetAccountClient returns the Databricks Account client or a diagnostics if that fails.
// This is used by resources and data sources that are developed over plugin framework.
func (c *DatabricksClient) GetAccountClient() (*databricks.AccountClient, diag.Diagnostics) {
	a, err := c.AccountClient()
	if err != nil {
		return nil, diag.Diagnostics{diag.NewErrorDiagnostic("Failed to get account client", err.Error())}
	}
	return a, nil
}

func (c *DatabricksClient) AccountClient() (*databricks.AccountClient, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.accountClient()
}

// accountClient returns the Databricks Account client or an error if that fails.
// The `mu` mutex must be held by the current goroutine when calling this method.
func (c *DatabricksClient) accountClient() (*databricks.AccountClient, error) {
	if c.cachedAccountClient != nil {
		return c.cachedAccountClient, nil
	}
	acc, err := databricks.NewAccountClient((*databricks.Config)(c.DatabricksClient.Config))
	if err != nil {
		return nil, err
	}
	c.cachedAccountClient = acc
	return acc, nil
}

func (c *DatabricksClient) AccountClientWithAccountIdFromConfig(d *schema.ResourceData) (*databricks.AccountClient, error) {
	accountID, ok := d.GetOk("account_id")
	if ok {
		err := c.setAccountId(accountID.(string))
		if err != nil {
			return nil, err
		}
	}
	return c.AccountClient()
}

func (c *DatabricksClient) AccountClientWithAccountIdFromPair(d *schema.ResourceData, p *Pair) (*databricks.AccountClient, string, error) {
	accountID, resourceId, err := p.Unpack(d)
	if err != nil {
		return nil, "", err
	}
	err = c.setAccountId(accountID)
	if err != nil {
		return nil, "", err
	}
	a, err := c.AccountClient()
	if err != nil {
		return nil, "", err
	}
	return a, resourceId, nil
}

// AccountOrWorkspaceRequest routes the request to account or workspace callbacks.
// It checks the `api` field in the resource data first. If set, it takes precedence
// over host-based inference. When `api` is not set, it falls back to the provider's
// host type (the original behavior).
func (c *DatabricksClient) AccountOrWorkspaceRequest(d *schema.ResourceData, accCallback func(*databricks.AccountClient) error, wsCallback func(*databricks.WorkspaceClient) error) error {
	if IsAccountLevel(d, c) {
		a, err := c.AccountClient()
		if err != nil {
			return err
		}
		return accCallback(a)
	}
	ws, err := c.WorkspaceClient()
	if err != nil {
		return err
	}
	return wsCallback(ws)
}

// Get on path
func (c *DatabricksClient) Get(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodGet, path, nil, nil, request, response, c.addApiPrefix)
}

// Post on path
func (c *DatabricksClient) Post(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodPost, path, nil, nil, request, response, c.addApiPrefix)
}

// Delete on path. Ignores succesfull responses from the server.
func (c *DatabricksClient) Delete(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodDelete, path, nil, nil, request, nil, c.addApiPrefix)
}

// Delete on path. Deserializes the response into the response parameter.
func (c *DatabricksClient) DeleteWithResponse(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodDelete, path, nil, nil, request, response, c.addApiPrefix)
}

// Patch on path. Ignores succesfull responses from the server.
func (c *DatabricksClient) Patch(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodPatch, path, nil, nil, request, nil, c.addApiPrefix)
}

// Patch on path. Deserializes the response into the response parameter.
func (c *DatabricksClient) PatchWithResponse(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodPatch, path, nil, nil, request, response, c.addApiPrefix)
}

// Put on path
func (c *DatabricksClient) Put(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodPut, path, nil, nil, request, nil, c.addApiPrefix)
}

const (
	// ApiLevelAccount indicates the resource should use account-level APIs.
	ApiLevelAccount = "account"
	// ApiLevelWorkspace indicates the resource should use workspace-level APIs.
	ApiLevelWorkspace = "workspace"
)

// AddApiField adds the `api` field to a resource schema. This field allows users
// to explicitly specify whether the resource should use account-level or
// workspace-level APIs. When set, it takes precedence over host-based inference.
func AddApiField(s map[string]*schema.Schema) map[string]*schema.Schema {
	s["api"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		ValidateFunc: validation.StringInSlice(
			[]string{ApiLevelAccount, ApiLevelWorkspace}, false,
		),
		Description: "Specifies whether to use account-level or workspace-level API. " +
			"Valid values are `account` and `workspace`. When not set, the API level " +
			"is inferred from the provider host.",
	}
	return s
}

// GetApiLevel returns the value of the `api` field from resource data,
// or empty string if not set.
func GetApiLevel(d *schema.ResourceData) string {
	if v, ok := d.GetOk("api"); ok {
		level := v.(string)
		if level == ApiLevelAccount || level == ApiLevelWorkspace {
			return level
		}
	}
	return ""
}

// GetApiLevelFromDiff returns the planned (new) value of the `api` field from
// a resource diff, or empty string if not set. This mirrors GetApiLevel but
// works with ResourceDiff (used in CustomizeDiff hooks).
func GetApiLevelFromDiff(d *schema.ResourceDiff) string {
	if v, ok := d.GetOk("api"); ok {
		level := v.(string)
		if level == ApiLevelAccount || level == ApiLevelWorkspace {
			return level
		}
	}
	return ""
}

// IsAccountLevelFromDiff determines whether a resource should use account-level APIs.
// This mirrors IsAccountLevel but works with ResourceDiff (used in CustomizeDiff).
func IsAccountLevelFromDiff(d *schema.ResourceDiff, c *DatabricksClient) bool {
	return isAccountLevelFromApiLevel(GetApiLevelFromDiff(d), c)
}

// isAccountLevelFromApiLevel determines whether a resource should use account-level APIs
// based on the api level string and the client's host type.
func isAccountLevelFromApiLevel(apiLevel string, c *DatabricksClient) bool {
	switch apiLevel {
	case ApiLevelAccount:
		return true
	case ApiLevelWorkspace:
		return false
	default:
		return c.HostTypeForTerraform() == config.AccountHost
	}
}

// IsAccountLevel determines whether a resource should use account-level APIs.
// It checks the `api` field first. If set, it takes precedence. Otherwise, it
// falls back to the provider's host type.
func IsAccountLevel(d *schema.ResourceData, c *DatabricksClient) bool {
	return isAccountLevelFromApiLevel(GetApiLevel(d), c)
}

type ApiVersion string

const (
	API_1_2 ApiVersion = "1.2"
	API_2_0 ApiVersion = "2.0"
	API_2_1 ApiVersion = "2.1"
)

func (c *DatabricksClient) addApiPrefix(r *http.Request) error {
	if r.URL == nil {
		return fmt.Errorf("no URL found in request")
	}
	ctx := r.Context()
	av, ok := ctx.Value(Api).(ApiVersion)
	if !ok {
		av = API_2_0
	}
	r.URL.Path = fmt.Sprintf("/api/%s%s", av, r.URL.Path)
	return nil
}

// scimVisitorForLevel returns a request visitor that rewrites SCIM URL paths
// for account-level requests. The apiLevel parameter takes precedence over
// host-based inference when non-empty.
func (c *DatabricksClient) scimVisitorForLevel(apiLevel string) func(*http.Request) error {
	return func(r *http.Request) error {
		var isAccount bool
		if apiLevel != "" {
			// Explicit api field takes precedence over host-based inference
			isAccount = apiLevel == ApiLevelAccount
		} else {
			isAccount = c.HostTypeForTerraform() == config.AccountHost
		}
		if isAccount {
			// until `/preview` is there for workspace scim,
			// `/api/2.0` is added by completeUrl visitor
			r.URL.Path = strings.ReplaceAll(r.URL.Path, "/api/2.0/preview",
				fmt.Sprintf("/api/2.0/accounts/%s", c.Config.AccountID))
		}
		return nil
	}
}

// Scim sets SCIM headers. The apiLevel parameter controls whether account-level
// or workspace-level SCIM endpoints are used. Pass "" to infer from the provider host.
func (c *DatabricksClient) Scim(ctx context.Context, method, path string, request any, response any, apiLevel string) error {
	return c.Do(ctx, method, path, map[string]string{
		"Content-Type": "application/scim+json; charset=utf-8",
	}, nil, request, response, c.addApiPrefix, c.scimVisitorForLevel(apiLevel))
}

// IsAzure returns true if client is configured for Azure Databricks - either by using AAD auth or with host+token combination
func (c *DatabricksClient) IsAzure() bool {
	return c.Config.IsAzure()
}

// acceptance.IsAws returns true if client is configured for AWS
func (c *DatabricksClient) IsAws() bool {
	return !c.IsGcp() && !c.IsAzure()
}

// acceptance.IsGcp returns true if client is configured for GCP
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
	cfg, err := c.DatabricksClient.Config.NewWithWorkspaceHost(url)
	if err != nil {
		return nil, fmt.Errorf("cannot configure new client: %w", err)
	}
	client, err := client.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot configure new client: %w", err)
	}
	// copy all client configuration options except Databricks CLI profile
	return &DatabricksClient{
		DatabricksClient: client,
		commandFactory:   c.commandFactory,
	}, nil
}

func (aa *DatabricksClient) GetAzureJwtProperty(key string) (any, error) {
	if !aa.IsAzure() {
		return "", fmt.Errorf("can't get Azure JWT token in non-Azure environment")
	}
	if key == "tid" && aa.Config.AzureTenantID != "" {
		return aa.Config.AzureTenantID, nil
	}
	request, err := http.NewRequest("GET", aa.Config.Host, nil)
	if err != nil {
		return nil, err
	}
	err = aa.Config.Authenticate(request)
	if err != nil {
		return nil, err
	}
	header := request.Header.Get("Authorization")
	var stoken string
	if len(header) > 0 && strings.HasPrefix(string(header), "Bearer ") {
		log.Printf("[DEBUG] Got Bearer token")
		stoken = strings.TrimSpace(strings.TrimPrefix(string(header), "Bearer "))
	}
	if stoken == "" {
		return nil, fmt.Errorf("can't obtain Azure JWT token")
	}
	if strings.HasPrefix(stoken, "dapi") {
		return nil, fmt.Errorf("can't use Databricks PAT")
	}
	parser := jwt.Parser{SkipClaimsValidation: true}
	token, _, err := parser.ParseUnverified(stoken, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	v, ok := claims[key]
	if !ok {
		return nil, fmt.Errorf("can't find field '%s' in parsed JWT", key)
	}
	return v, nil
}

func CommonEnvironmentClient() *DatabricksClient {
	c, err := client.New(&config.Config{})
	if err != nil {
		panic(err)
	}
	return &DatabricksClient{
		DatabricksClient: c,
		commandFactory: func(ctx context.Context, dc *DatabricksClient) CommandExecutor {
			panic("command executor not initalized")
		},
	}
}
