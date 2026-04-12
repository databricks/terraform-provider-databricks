package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceGroupMemberCreate_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Groups/abc",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Groups/abc?attributes=members",
				Response: Group{
					Members: []ComplexValue{{Value: "def"}},
				},
			},
		},
		Resource:  ResourceGroupMember(),
		AccountID: "acc-123",
		HCL: `
			group_id = "abc"
			member_id = "def"
			api = "account"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceGroupMemberRead_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Groups/abc?attributes=members",
				Response: Group{
					Members: []ComplexValue{{Value: "def"}},
				},
			},
		},
		Resource:  ResourceGroupMember(),
		AccountID: "acc-123",
		HCL: `
			group_id = "abc"
			member_id = "def"
			api = "account"
		`,
		New:  true,
		Read: true,
		ID:   "abc|def",
	}.ApplyNoError(t)
}
