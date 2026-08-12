package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
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
		Resource:            ResourceGroup(),
		AccountID:           "acc-123",
		ProviderWorkspaceID: "12345",
		HCL: `
			display_name = "ws-group"
			api = "workspace"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceGroupCreate_ApiFieldNotSet_FallsBackToHostInference(t *testing.T) {
	// When api is NOT set, account host routes to account SCIM (backwards compatible).
	// Host inference is driven by the host metadata discovery — a 200 on
	// /.well-known/databricks-config with host_type=ACCOUNT_HOST resolves the
	// config to AccountHost, which is what a real accounts.* host would do.
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				Response:     map[string]string{"host_type": "account"},
				ReuseRequest: true,
			},
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
		`,
		Create: true,
	}.ApplyNoError(t)
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
