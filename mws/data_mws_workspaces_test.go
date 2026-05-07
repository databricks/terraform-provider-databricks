package mws

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceMwsWorkspaces(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/workspaces",

				Response: []Workspace{
					{
						WorkspaceName: "bcd",
						WorkspaceID:   123,
					},
					{
						WorkspaceName: "def",
						WorkspaceID:   456,
					},
				},
			},
		},
		AccountID:   "abc",
		Resource:    DataSourceMwsWorkspaces(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]any{
			"bcd": 123,
			"def": 456,
		},
	})
}

func TestCatalogsData_Error(t *testing.T) {
	qa.ResourceFixture{
		AccountID:   "abc",
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceMwsWorkspaces(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}

// TestDataSourceMwsWorkspaces_AccountLevelNoHookFailure is a regression test
// for https://github.com/databricks/terraform-provider-databricks/issues/5672.
// At account level, the post-Read provider_config hook used to fail with
// "cannot populate provider_config for mws workspaces: failed to resolve
// workspace_id" because the data source has no workspace context. The fix
// short-circuits the hook for resources that opt in via
// common.Resource.SkipProviderConfigStatePopulation.
func TestDataSourceMwsWorkspaces_AccountLevelNoHookFailure(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/workspaces",
			Response: []Workspace{{WorkspaceName: "n", WorkspaceID: 1}},
		},
	})
	assert.NoError(t, err)
	defer server.Close()

	// Simulate a real account-level provider: no cached workspace ID and
	// AccountID set on the config. Without the fix, this would crash inside
	// populateProviderConfigInState trying to resolve a workspace ID.
	client.SetCachedWorkspaceID(0)
	client.Config.AccountID = "abc"

	r := DataSourceMwsWorkspaces().ToResource()
	d := r.TestResourceData()
	d.SetId("_")
	ctx := context.WithValue(context.Background(), common.ResourceName, "mws_workspaces")
	diags := r.ReadContext(ctx, d, client)

	assert.False(t, diags.HasError(), "post-Read hook must be skipped for account-only data source: %v", diags)
}

func TestDataSourceMwsWorkspaces_Empty(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/workspaces",

				Response: []Workspace{},
			},
		},
		AccountID:   "abc",
		Resource:    DataSourceMwsWorkspaces(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]any{},
	})
}
