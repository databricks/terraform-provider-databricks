package qa

import (
	"context"
	"errors"

	databricks "github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/terraform-provider-databricks/common"
)

type testingClient struct {
	a      *mocks.MockAccountClient
	w      *mocks.MockWorkspaceClient
	config *config.Config
}

// AccountClient implements common.DatabricksAPI.
func (t testingClient) AccountClient() (*databricks.AccountClient, error) {
	if t.a == nil {
		return nil, errors.New("AccountClient not set")
	}
	return t.a.AccountClient, nil
}

// AccountOrWorkspaceRequest implements common.DatabricksAPI.
func (t testingClient) AccountOrWorkspaceRequest(accCallback func(*databricks.AccountClient) error, wsCallback func(*databricks.WorkspaceClient) error) error {
	if t.Config().IsAccountClient() {
		return accCallback(t.a.AccountClient)
	}
	return wsCallback(t.w.WorkspaceClient)
}

// ClientForHost implements common.DatabricksAPI.
func (t testingClient) ClientForHost(ctx context.Context, url string) (*common.DatabricksClient, error) {
	panic("unimplemented")
}

// CommandExecutor implements common.DatabricksAPI.
func (testingClient) CommandExecutor(ctx context.Context) common.CommandExecutor {
	panic("unimplemented")
}

// Config implements common.DatabricksAPI.
func (t testingClient) Config() *config.Config {
	return t.config
}

// Delete implements common.DatabricksAPI.
func (testingClient) Delete(ctx context.Context, path string, request any) error {
	panic("unimplemented")
}

// DeleteWithResponse implements common.DatabricksAPI.
func (testingClient) DeleteWithResponse(ctx context.Context, path string, request any, response any) error {
	panic("unimplemented")
}

// FormatURL implements common.DatabricksAPI.
func (testingClient) FormatURL(strs ...string) string {
	panic("unimplemented")
}

// Get implements common.DatabricksAPI.
func (testingClient) Get(ctx context.Context, path string, request any, response any) error {
	panic("unimplemented")
}

// GetAzureJwtProperty implements common.DatabricksAPI.
func (testingClient) GetAzureJwtProperty(key string) (any, error) {
	panic("unimplemented")
}

// IsAws implements common.DatabricksAPI.
func (t testingClient) IsAws() bool {
	return t.Config().IsAws()
}

// IsAzure implements common.DatabricksAPI.
func (t testingClient) IsAzure() bool {
	return t.Config().IsAzure()
}

// IsGcp implements common.DatabricksAPI.
func (t testingClient) IsGcp() bool {
	return t.Config().IsGcp()
}

// Patch implements common.DatabricksAPI.
func (testingClient) Patch(ctx context.Context, path string, request any) error {
	panic("unimplemented")
}

// PatchWithResponse implements common.DatabricksAPI.
func (testingClient) PatchWithResponse(ctx context.Context, path string, request any, response any) error {
	panic("unimplemented")
}

// Post implements common.DatabricksAPI.
func (testingClient) Post(ctx context.Context, path string, request any, response any) error {
	panic("unimplemented")
}

// Put implements common.DatabricksAPI.
func (testingClient) Put(ctx context.Context, path string, request any) error {
	panic("unimplemented")
}

// Scim implements common.DatabricksAPI.
func (testingClient) Scim(ctx context.Context, method string, path string, request any, response any) error {
	panic("unimplemented")
}

// SetAccountId implements common.DatabricksAPI.
func (testingClient) SetAccountId(accountId string) error {
	panic("unimplemented")
}

// WithCommandExecutor implements common.DatabricksAPI.
func (testingClient) WithCommandExecutor(cef func(context.Context, *common.DatabricksClient) common.CommandExecutor) {
	panic("unimplemented")
}

// WithCommandMock implements common.DatabricksAPI.
func (testingClient) WithCommandMock(mock common.CommandMock) {
	panic("unimplemented")
}

// WorkspaceClient implements common.DatabricksAPI.
func (t testingClient) WorkspaceClient() (*databricks.WorkspaceClient, error) {
	if t.w == nil {
		return nil, errors.New("WorkspaceClient not set")
	}
	return t.w.WorkspaceClient, nil
}

var _ common.DatabricksAPI = testingClient{}

var ErrImATeapot = errors.New("I'm a teapot")
