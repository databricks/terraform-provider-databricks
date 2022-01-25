package aws

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/scim"
)

func TestUserRoleCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceUserRole(),
		qa.CornerCaseID("a|b"),
		qa.CornerCaseSkipCRUD("create"))
}

func TestUserRoleCreate_AndGetResourceDrift(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Users/a",
				ExpectedRequest: scim.PatchRequest("add", "roles", "b"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/a",
				Response: scim.User{},
			},
		},
		Create:   true,
		Resource: ResourceUserRole(),
		HCL: `
		user_id = "a"
		role = "b"
		`,
	}.ExpectError(t, "User has no role")
}
