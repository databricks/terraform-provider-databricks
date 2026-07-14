package aws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/scim"
)

func TestResourceUserRoleCreate_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Users/abc",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Users/abc?attributes=roles",
				Response: scim.User{
					Roles: []scim.ComplexValue{{Value: "arn:aws:iam::999999999999:role/foo"}},
				},
			},
		},
		Resource:  ResourceUserRole(),
		AccountID: "acc-123",
		HCL: `
			user_id = "abc"
			role = "arn:aws:iam::999999999999:role/foo"
			api = "account"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceUserRoleRead_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Users/abc?attributes=roles",
				Response: scim.User{
					Roles: []scim.ComplexValue{{Value: "arn:aws:iam::999999999999:role/foo"}},
				},
			},
		},
		Resource:  ResourceUserRole(),
		AccountID: "acc-123",
		HCL: `
			user_id = "abc"
			role = "arn:aws:iam::999999999999:role/foo"
			api = "account"
		`,
		New:  true,
		Read: true,
		ID:   "abc|arn:aws:iam::999999999999:role/foo",
	}.ApplyNoError(t)
}
