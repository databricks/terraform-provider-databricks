package scim

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceGroupCreate_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Groups",
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Groups/abc?attributes=displayName,externalId,entitlements",
				Response: Group{
					ID:          "abc",
					DisplayName: "test-group",
				},
			},
		},
		Resource:  ResourceGroup(),
		AccountID: "acc-123",
		HCL: `
			display_name = "test-group"
			api = "account"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceGroupCreate_ApiFieldWorkspace(t *testing.T) {
	// Even with AccountID set, api = "workspace" routes to workspace SCIM
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Groups",
				Response: Group{
					ID: "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/def?attributes=displayName,externalId,entitlements",
				Response: Group{
					ID:          "def",
					DisplayName: "ws-group",
				},
			},
		},
		Resource:  ResourceGroup(),
		AccountID: "acc-123",
		HCL: `
			display_name = "ws-group"
			api = "workspace"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceGroupCreate_ApiFieldNotSet_FallsBackToHostInference(t *testing.T) {
	// When api is NOT set, the host URL determines the SCIM endpoint used.
	// With a non-accounts host (e.g. test server on localhost), routes to workspace SCIM.
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Groups",
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,externalId,entitlements",
				Response: Group{
					ID:          "abc",
					DisplayName: "test-group",
				},
			},
		},
		Resource:  ResourceGroup(),
		AccountID: "acc-123",
		HCL: `
			display_name = "test-group"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceGroupCreate_ApiFieldNotSet_AccountsHostInference(t *testing.T) {
	// When api is NOT set but the host IS an accounts URL,
	// HostTypeForTerraform() infers account level from the URL prefix.
	// This is the production path: users configure host = "https://accounts.cloud.databricks.com"
	// and IsAccountLevel correctly returns true without needing api = "account".
	c := &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "acc-123",
			},
		},
	}
	d := ResourceGroup().ToResource().TestResourceData()
	// api field is NOT set — IsAccountLevel should still return true
	// because HostTypeForTerraform() detects the accounts URL prefix
	assert.Equal(t, config.AccountHost, c.HostTypeForTerraform())
	assert.True(t, common.IsAccountLevel(d, c))
}

func TestResourceGroupCreate_ApiFieldNotSet_WorkspaceHost(t *testing.T) {
	// When api is NOT set and provider is workspace-level, routes to workspace SCIM (backwards compatible)
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Groups",
				Response: Group{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=displayName,externalId,entitlements",
				Response: Group{
					ID:          "abc",
					DisplayName: "test-group",
				},
			},
		},
		Resource: ResourceGroup(),
		HCL: `
			display_name = "test-group"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceGroupRead_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Groups/abc?attributes=displayName,externalId,entitlements",
				Response: Group{
					ID:          "abc",
					DisplayName: "test-group",
				},
			},
		},
		Resource:  ResourceGroup(),
		AccountID: "acc-123",
		HCL: `
			display_name = "test-group"
			api = "account"
		`,
		New:  true,
		Read: true,
		ID:   "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"display_name": "test-group",
	})
}
