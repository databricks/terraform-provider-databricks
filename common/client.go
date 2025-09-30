// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package common

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

	// cachedWorkspaceClients is a map of workspace clients for each workspace ID
	// populated when fetching a WorkspaceClient for a specific workspace ID using
	// a provider configured at the account level
	cachedWorkspaceClients map[int64]*databricks.WorkspaceClient

	// cachedAccountClient is a cached account client authenticated to the account
	// configured for the provider
	cachedAccountClient *databricks.AccountClient

	// mu synchronizes access to all cached clients.
	mu sync.Mutex
}

// GetWorkspaceClient returns the Databricks WorkspaceClient or a diagnostics if that fails.
// This is used by resources and data sources that are developed over plugin framework.
func (c *DatabricksClient) GetWorkspaceClient() (*databricks.WorkspaceClient, diag.Diagnostics) {
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
	a, err := c.accountClient()
	if err != nil {
		return nil, err
	}
	workspace, err := a.Workspaces.Get(ctx, provisioning.GetWorkspaceRequest{WorkspaceId: workspaceId})
	if err != nil {
		return nil, err
	}
	w, err := a.GetWorkspaceClient(*workspace)
	if err != nil {
		return nil, err
	}
	w.CurrentUser = newCachedMe(w.CurrentUser)
	c.cachedWorkspaceClients[workspace.WorkspaceId] = w
	return w, nil
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

func (c *DatabricksClient) AccountOrWorkspaceRequest(accCallback func(*databricks.AccountClient) error, wsCallback func(*databricks.WorkspaceClient) error) error {
	if c.Config.IsAccountClient() {
		a, err := c.AccountClient()
		if err != nil {
			return err
		}
		return accCallback(a)
	} else {
		ws, err := c.WorkspaceClient()
		if err != nil {
			return err
		}
		return wsCallback(ws)
	}
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

// Put on path
func (c *DatabricksClient) PutWithResponse(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodPut, path, nil, nil, request, response, c.addApiPrefix)
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

// scimVisitor is a separate method for the sake of unit tests
func (c *DatabricksClient) scimVisitor(r *http.Request) error {
	if c.Config.IsAccountClient() && c.Config.AccountID != "" {
		// until `/preview` is there for workspace scim,
		// `/api/2.0` is added by completeUrl visitor
		r.URL.Path = strings.ReplaceAll(r.URL.Path, "/api/2.0/preview",
			fmt.Sprintf("/api/2.0/accounts/%s", c.Config.AccountID))
	}
	return nil
}

// Scim sets SCIM headers
func (c *DatabricksClient) Scim(ctx context.Context, method, path string, request any, response any) error {
	return c.Do(ctx, method, path, map[string]string{
		"Content-Type": "application/scim+json; charset=utf-8",
	}, nil, request, response, c.addApiPrefix, c.scimVisitor)
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
