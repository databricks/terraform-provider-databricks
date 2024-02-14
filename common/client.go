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
	"github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type cachedMe struct {
	internalImpl iam.CurrentUserService
	cachedUser   *iam.User
	mu           sync.Mutex
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

type cachedWorkspaceClient struct {
	apiClient *client.DatabricksClient
	client    *databricks.WorkspaceClient
}

// DatabricksClient holds properties needed for authentication and HTTP client setup
// fields with `name` struct tags become Terraform provider attributes. `env` struct tag
// can hold one or more coma-separated env variable names to find value, if not specified
// directly. `auth` struct tag describes the type of conflicting authentication used.
type DatabricksClient struct {
	*client.DatabricksClient

	// callback used to create API1.2 call wrapper, which simplifies unit testing
	commandFactory      func(context.Context, *DatabricksClient) CommandExecutor
	cachedAccountClient *databricks.AccountClient
	mu                  sync.Mutex

	// cachedWorkspaceClients is a cache of WorkspaceClients by workspace ID. This is used
	// when the provider is configured at the account-level and used to manage workspace-
	// level resources.
	cachedWorkspaceClients   map[int64]*cachedWorkspaceClient
	cachedWorkspaceClientsMu sync.Mutex

	// currentWorkspaceID is a cache of the workspace ID for the current cached workspace
	// client. It should only be set either when the provider is configured at the workspace-
	// level (in which case it is lazily loaded) or when getting a workspace-level client
	// from a provider configured at the account-level (in which case it is eagerly
	// initialized).
	currentWorkspaceID   int64
	currentWorkspaceIDMu sync.Mutex
}

var defaultWorkspaceId int64 = -1

// WorkspaceClient() returns a WorkspaceClient for the current workspace. Fails if the provider is
// configured at the account level.
func (c *DatabricksClient) WorkspaceClient() (*databricks.WorkspaceClient, error) {
	clients, err := c.getConfiguredWorkspaceClient(context.Background(), defaultWorkspaceId)
	if err != nil {
		return nil, err
	}
	return clients.client, nil
}

// CurrentWorkspaceID() returns the workspace ID of the workspace that this client is connected to.
func (c *DatabricksClient) CurrentWorkspaceID(ctx context.Context) (int64, error) {
	if c.currentWorkspaceID != 0 {
		return c.currentWorkspaceID, nil
	}
	c.currentWorkspaceIDMu.Lock()
	defer c.currentWorkspaceIDMu.Unlock()
	if c.currentWorkspaceID != 0 {
		return c.currentWorkspaceID, nil
	}
	w, err := c.WorkspaceClient()
	if err != nil {
		return 0, err
	}
	c.currentWorkspaceID, err = w.CurrentWorkspaceID(ctx)
	if err != nil {
		return 0, err
	}
	return c.currentWorkspaceID, nil
}

// InWorkspace() returns a new DatabricksClient instance configured to interact with the specified
// workspace. The underlying WorkspaceClient and DatabricksClient instances are shared by all callers
// of this method.
func (c *DatabricksClient) InWorkspace(ctx context.Context, workspaceId int64) (*DatabricksClient, error) {
	clients, err := c.getConfiguredWorkspaceClient(ctx, workspaceId)
	if err != nil {
		return nil, err
	}
	return &DatabricksClient{
		DatabricksClient:       clients.apiClient,
		commandFactory:         c.commandFactory,
		cachedWorkspaceClients: c.cachedWorkspaceClients,
		currentWorkspaceID:     workspaceId,
	}, nil
}

// getConfiguredWorkspaceClient gets the workspace client for the given workspace ID. This will be
// retrieved from the cache if it has already been created. Otherwise, it will be constructed from
// the account client if the provider is configured at the account level, or from the existing
// workspace client if the provider is configured at the workspace level.
func (c *DatabricksClient) getConfiguredWorkspaceClient(ctx context.Context, workspaceId int64) (*cachedWorkspaceClient, error) {
	if c.cachedWorkspaceClients == nil {
		c.cachedWorkspaceClientsMu.Lock()
		if c.cachedWorkspaceClients == nil {
			c.cachedWorkspaceClients = make(map[int64]*cachedWorkspaceClient)
		}
		c.cachedWorkspaceClientsMu.Unlock()
	}
	if clients, ok := c.cachedWorkspaceClients[workspaceId]; ok {
		return clients, nil
	}
	c.cachedWorkspaceClientsMu.Lock()
	defer c.cachedWorkspaceClientsMu.Unlock()
	if w, ok := c.cachedWorkspaceClients[workspaceId]; ok {
		return w, nil
	}
	w, err := c.makeWorkspaceClient(ctx, workspaceId)
	if err != nil {
		return nil, err
	}
	internalImpl := w.CurrentUser.Impl()
	w.CurrentUser.WithImpl(&cachedMe{
		internalImpl: internalImpl,
	})
	apiClient, err := client.New(w.Config)
	if err != nil {
		return nil, err
	}
	clients := cachedWorkspaceClient{
		apiClient: apiClient,
		client:    w,
	}
	c.cachedWorkspaceClients[workspaceId] = &clients
	return &clients, nil
}

func (c *DatabricksClient) makeWorkspaceClient(ctx context.Context, workspaceId int64) (*databricks.WorkspaceClient, error) {
	if c.Config.IsAccountClient() {
		// If provider is configured at the account-level, construct the workspace client from the account
		// client and cache it.
		a, err := c.AccountClient()
		if err != nil {
			return nil, err
		}
		ws, err := a.Workspaces.GetByWorkspaceId(ctx, workspaceId)
		if err != nil {
			return nil, err
		}
		w, err := a.GetWorkspaceClient(*ws)
		if err != nil {
			return nil, err
		}
		return w, nil
	}
	// If configured at the workspace level, verify that the specified workspace ID matches the current workspace ID.
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(c.DatabricksClient.Config))
	if err != nil {
		return nil, err
	}
	// Only check current workspace ID if it isn't defaultWorkspaceId
	if workspaceId != defaultWorkspaceId {
		return w, nil
	}
	id, err := w.CurrentWorkspaceID(ctx)
	if err != nil {
		return nil, err
	}
	if workspaceId != id {
		return nil, fmt.Errorf("workspace ID %d does not match current workspace ID %d", workspaceId, id)
	}
	return w, nil
}

// Set the cached workspace client.
func (c *DatabricksClient) SetWorkspaceClient(w *databricks.WorkspaceClient) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cachedWorkspaceClients = map[int64]*cachedWorkspaceClient{
		defaultWorkspaceId: {
			apiClient: c.DatabricksClient,
			client:    w,
		},
	}
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

func (c *DatabricksClient) AccountClient() (*databricks.AccountClient, error) {
	if c.cachedAccountClient != nil {
		return c.cachedAccountClient, nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
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
	return c.Do(ctx, http.MethodGet, path, nil, request, response, c.addApiPrefix)
}

// Post on path
func (c *DatabricksClient) Post(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodPost, path, nil, request, response, c.addApiPrefix)
}

// Delete on path. Ignores succesfull responses from the server.
func (c *DatabricksClient) Delete(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodDelete, path, nil, request, nil, c.addApiPrefix)
}

// Delete on path. Deserializes the response into the response parameter.
func (c *DatabricksClient) DeleteWithResponse(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodDelete, path, nil, request, response, c.addApiPrefix)
}

// Patch on path. Ignores succesfull responses from the server.
func (c *DatabricksClient) Patch(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodPatch, path, nil, request, nil, c.addApiPrefix)
}

// Patch on path. Deserializes the response into the response parameter.
func (c *DatabricksClient) PatchWithResponse(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodPatch, path, nil, request, response, c.addApiPrefix)
}

// Put on path
func (c *DatabricksClient) Put(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodPut, path, nil, request, nil, c.addApiPrefix)
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
	}, request, response, c.addApiPrefix, c.scimVisitor)
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
