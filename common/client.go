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

// DatabricksClient holds properties needed for authentication and HTTP client setup
// fields with `name` struct tags become Terraform provider attributes. `env` struct tag
// can hold one or more coma-separated env variable names to find value, if not specified
// directly. `auth` struct tag describes the type of conflicting authentication used.
type DatabricksClient struct {
	*client.DatabricksClient

	// callback used to create API1.2 call wrapper, which simplifies unit tessting
	commandFactory        func(context.Context, *DatabricksClient) CommandExecutor
	cachedAccountClient   *databricks.AccountClient
	cachedWorkspaceClient *databricks.WorkspaceClient
	mu                    sync.Mutex
}

func (c *DatabricksClient) AccountClient() (*databricks.AccountClient, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cachedAccountClient != nil {
		return c.cachedAccountClient, nil
	}
	a, err := databricks.NewAccountClient((*databricks.Config)(c.DatabricksClient.Config))
	if err != nil {
		return nil, err
	}
	c.cachedAccountClient = a
	return a, nil
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
	internalImpl := w.CurrentUser.Impl()
	w.CurrentUser.WithImpl(&cachedMe{
		internalImpl: internalImpl,
	})
	c.cachedWorkspaceClient = w
	return w, nil
}

// Get on path
func (c *DatabricksClient) Get(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodGet, path, request, response, c.addApiPrefix)
}

// Post on path
func (c *DatabricksClient) Post(ctx context.Context, path string, request any, response any) error {
	return c.Do(ctx, http.MethodPost, path, request, response, c.addApiPrefix)
}

// Delete on path
func (c *DatabricksClient) Delete(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodDelete, path, request, nil, c.addApiPrefix)
}

// Patch on path
func (c *DatabricksClient) Patch(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodPatch, path, request, nil, c.addApiPrefix)
}

// Put on path
func (c *DatabricksClient) Put(ctx context.Context, path string, request any) error {
	return c.Do(ctx, http.MethodPut, path, request, nil, c.addApiPrefix)
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
	r.Header.Set("Content-Type", "application/scim+json; charset=utf-8")
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
	return c.Do(ctx, method, path, request, response, c.addApiPrefix, c.scimVisitor)
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
