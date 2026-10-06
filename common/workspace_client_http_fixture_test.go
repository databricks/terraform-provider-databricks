package common_test

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkspaceClientsFromUnifiedAccountAreIsolatedAtHTTPBoundary(t *testing.T) {
	client, server, err := qa.StrictHTTPFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/account-id/workspaces/111?",
			Response: provisioning.Workspace{WorkspaceId: 111, DeploymentName: "workspace-one"},
		},
		{
			Method:          "GET",
			Resource:        "/api/2.0/preview/scim/v2/Me?",
			ExpectedHeaders: map[string]string{"X-Databricks-Workspace-Id": "111"},
			Response:        iam.User{Id: "user-one", UserName: "one@databricks.com"},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/account-id/workspaces/222?",
			Response: provisioning.Workspace{WorkspaceId: 222, DeploymentName: "workspace-two"},
		},
		{
			Method:          "GET",
			Resource:        "/api/2.0/preview/scim/v2/Me?",
			ExpectedHeaders: map[string]string{"X-Databricks-Workspace-Id": "222"},
			Response:        iam.User{Id: "user-two", UserName: "two@databricks.com"},
		},
	})
	require.NoError(t, err)

	client.Config.AccountID = "account-id"
	client.Config.HostMetadataResolver = func(context.Context, string) (*config.HostMetadata, error) {
		return &config.HostMetadata{HostType: config.UnifiedHost}, nil
	}

	first, err := client.WorkspaceClientForWorkspace(context.Background(), "111")
	require.NoError(t, err)
	firstUser, err := first.CurrentUser.Me(context.Background(), iam.MeRequest{})
	require.NoError(t, err)

	second, err := client.WorkspaceClientForWorkspace(context.Background(), "222")
	require.NoError(t, err)
	secondUser, err := second.CurrentUser.Me(context.Background(), iam.MeRequest{})
	require.NoError(t, err)

	assert.Equal(t, "one@databricks.com", firstUser.UserName)
	assert.Equal(t, "two@databricks.com", secondUser.UserName)
	assert.Equal(t, "111", first.Config.WorkspaceID)
	assert.Equal(t, "222", second.Config.WorkspaceID)
	assert.Empty(t, client.Config.WorkspaceID)
	assert.Equal(t, server.URL, first.Config.Host)
	assert.Equal(t, server.URL, second.Config.Host)
}
