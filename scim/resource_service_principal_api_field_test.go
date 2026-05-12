package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceServicePrincipalCreate_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/ServicePrincipals",
				Response: User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/ServicePrincipals/abc?attributes=userName,displayName,active,externalId,entitlements",
				Response: User{
					DisplayName:   "Example SP",
					Active:        true,
					ID:            "abc",
					ApplicationID: "bcd",
				},
			},
		},
		Resource:  ResourceServicePrincipal(),
		AccountID: "acc-123",
		HCL: `
			display_name = "Example SP"
			api = "account"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalCreate_ApiFieldWorkspace(t *testing.T) {
	// Even when AccountID is set on the provider, api = "workspace" should
	// route to workspace SCIM (preview path, not account path).
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals",
				Response: User{
					ID: "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/def?attributes=userName,displayName,active,externalId,entitlements",
				Response: User{
					DisplayName:   "WS SP",
					Active:        true,
					ID:            "def",
					ApplicationID: "efg",
				},
			},
		},
		Resource:            ResourceServicePrincipal(),
		AccountID:           "acc-123",
		ProviderWorkspaceID: "12345",
		HCL: `
			display_name = "WS SP"
			api = "workspace"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalRead_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/ServicePrincipals/abc?attributes=userName,displayName,active,externalId,entitlements",
				Response: User{
					ID:            "abc",
					ApplicationID: "bcd",
					DisplayName:   "SP Account",
					Active:        true,
				},
			},
		},
		Resource:  ResourceServicePrincipal(),
		AccountID: "acc-123",
		HCL: `
			display_name = "SP Account"
			api = "account"
		`,
		New:  true,
		Read: true,
		ID:   "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"display_name":   "SP Account",
		"application_id": "bcd",
	})
}

func TestResourceServicePrincipalDelete_ApiFieldAccount(t *testing.T) {
	// When api = "account", delete should disable (PATCH) rather than delete.
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/ServicePrincipals/abc",
			},
		},
		Resource:  ResourceServicePrincipal(),
		AccountID: "acc-123",
		HCL: `
			display_name = "SP Account"
			api = "account"
		`,
		Delete: true,
		ID:     "abc",
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalDelete_ApiFieldWorkspace(t *testing.T) {
	// When api = "workspace", delete should actually DELETE (not PATCH/disable),
	// even when AccountID is set on the provider.
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
			},
		},
		Resource:            ResourceServicePrincipal(),
		AccountID:           "acc-123",
		ProviderWorkspaceID: "12345",
		HCL: `
			display_name = "SP WS"
			api = "workspace"
		`,
		Delete: true,
		ID:     "abc",
	}.ApplyNoError(t)
}
