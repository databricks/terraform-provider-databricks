package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

var expectedUserDisablePatchRequest = patchRequest{
	Operations: []patchOperation{
		patchOperation{
			Op:   "replace",
			Path: "active",
			Value: []ComplexValue{
				{
					Value: "false",
				},
			},
		},
	},
	Schemas: []URN{PatchOp},
}

func TestResourceUserDeleteAsDisable_NoError(t *testing.T) {
	qa.ResourceFixture{
		AccountID: "00000000-0000-0000-0000-000000000001",
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/accounts/00000000-0000-0000-0000-000000000001/scim/v2/Users/abc",
				ExpectedRequest: expectedUserDisablePatchRequest,
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name    = "abc"
			disable_as_user_deletion = true
		`,
	}.ApplyNoError(t)
}

func TestResourceUserDeleteAsDisable_NoErrorEmptyParams(t *testing.T) {
	qa.ResourceFixture{
		AccountID: "00000000-0000-0000-0000-000000000001",
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/accounts/00000000-0000-0000-0000-000000000001/scim/v2/Users/abc",
				ExpectedRequest: expectedUserDisablePatchRequest,
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name    = "abc"
		`,
	}.ApplyNoError(t)
}

func TestResourceUserDeleteAsDisable_IgnoreIfNotAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/Users/abc",
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name = "abc"
			disable_as_user_deletion = true
		`,
	}.ApplyNoError(t)
}
