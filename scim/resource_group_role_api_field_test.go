package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceGroupRoleCreate_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Groups/abc",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Groups/abc?attributes=roles",
				Response: Group{
					Roles: []ComplexValue{{Value: "arn:aws:iam::999999999999:role/foo"}},
				},
			},
		},
		Resource:  ResourceGroupRole(),
		AccountID: "acc-123",
		HCL: `
			group_id = "abc"
			role = "arn:aws:iam::999999999999:role/foo"
			api = "account"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceGroupRoleCreate_ApiFieldWorkspace(t *testing.T) {
	// Even with AccountID set, api = "workspace" routes to workspace SCIM
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/abc?attributes=roles",
				Response: Group{
					Roles: []ComplexValue{{Value: "arn:aws:iam::999999999999:role/foo"}},
				},
			},
		},
		Resource:            ResourceGroupRole(),
		AccountID:           "acc-123",
		ProviderWorkspaceID: "12345",
		HCL: `
			group_id = "abc"
			role = "arn:aws:iam::999999999999:role/foo"
			api = "workspace"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceGroupRoleRead_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Groups/abc?attributes=roles",
				Response: Group{
					Roles: []ComplexValue{{Value: "arn:aws:iam::999999999999:role/foo"}},
				},
			},
		},
		Resource:  ResourceGroupRole(),
		AccountID: "acc-123",
		HCL: `
			group_id = "abc"
			role = "arn:aws:iam::999999999999:role/foo"
			api = "account"
		`,
		New:  true,
		Read: true,
		ID:   "abc|arn:aws:iam::999999999999:role/foo",
	}.ApplyNoError(t)
}
