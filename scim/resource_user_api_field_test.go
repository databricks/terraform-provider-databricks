package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceUserCreate_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Users",
				Response: User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Users/abc?attributes=userName,displayName,active,externalId,entitlements",
				Response: User{
					UserName:    "me@example.com",
					DisplayName: "me",
					Active:      true,
					ID:          "abc",
				},
			},
		},
		Resource:  ResourceUser(),
		AccountID: "acc-123",
		HCL: `
			user_name = "me@example.com"
			api = "account"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceUserCreate_ApiFieldWorkspace(t *testing.T) {
	// Even with AccountID set, api = "workspace" routes to workspace SCIM
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/preview/scim/v2/Users",
				Response: User{
					ID: "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/def?attributes=userName,displayName,active,externalId,entitlements",
				Response: User{
					UserName:    "ws@example.com",
					DisplayName: "ws",
					Active:      true,
					ID:          "def",
				},
			},
		},
		Resource:            ResourceUser(),
		AccountID:           "acc-123",
		ProviderWorkspaceID: "12345",
		HCL: `
			user_name = "ws@example.com"
			api = "workspace"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceUserDelete_ApiFieldAccount(t *testing.T) {
	// api = "account" should disable (PATCH) instead of delete
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Users/abc",
			},
		},
		Resource:  ResourceUser(),
		AccountID: "acc-123",
		HCL: `
			user_name = "me@example.com"
			api = "account"
		`,
		Delete: true,
		ID:     "abc",
	}.ApplyNoError(t)
}

func TestResourceUserDelete_ApiFieldWorkspace(t *testing.T) {
	// api = "workspace" should actually delete even with AccountID set
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
			},
		},
		Resource:            ResourceUser(),
		AccountID:           "acc-123",
		ProviderWorkspaceID: "12345",
		HCL: `
			user_name = "me@example.com"
			api = "workspace"
		`,
		Delete: true,
		ID:     "abc",
	}.ApplyNoError(t)
}
