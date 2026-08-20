package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceGroupMemberCreate_ApiFieldAccount(t *testing.T) {
	globalGroupsCache = newGroupCache()
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/Groups/abc",
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/accounts/acc-123/scim/v2/Groups?attributes=id%2Cmembers&count=10000&startIndex=1",
				ReuseRequest: true,
				Response: GroupList{
					TotalResults: 1,
					Resources: []Group{
						{ID: "abc", Members: []ComplexValue{{Value: "def"}}},
					},
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
	globalGroupsCache = newGroupCache()
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.0/accounts/acc-123/scim/v2/Groups?attributes=id%2Cmembers&count=10000&startIndex=1",
				ReuseRequest: true,
				Response: GroupList{
					TotalResults: 1,
					Resources: []Group{
						{ID: "abc", Members: []ComplexValue{{Value: "def"}}},
					},
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
