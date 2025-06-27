package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

var expectedUserDisablePatchRequest = patchRequest{
	Operations: []patchOperation{
		{
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

func TestResourceUserDeleteAsDisableInAccount_NoError(t *testing.T) {
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

func TestResourceUserDeleteAsDisableInAccount_NoErrorEmptyParams(t *testing.T) {
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

func TestResourceUserDeleteAsDisableInAccount_HardDelete(t *testing.T) {
	qa.ResourceFixture{
		AccountID: "00000000-0000-0000-0000-000000000001",
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/00000000-0000-0000-0000-000000000001/scim/v2/Users/abc",
			},
		},
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name    = "abc"
			disable_as_user_deletion = false
		`,
	}.ApplyNoError(t)
}

func TestResourceUserDeleteAsDisableInWorkspace_NoError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/Users/abc",
				ExpectedRequest: expectedUserDisablePatchRequest,
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

func TestResourceUserDeleteAsDisableInWorkspace_ErrorForceDeleteRepos(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name = "abc"
			disable_as_user_deletion = true
			force_delete_repos = true
		`,
	}.ExpectError(t, "force_delete_repos: cannot force delete if disable_as_user_deletion is set")
}

func TestResourceUserDeleteAsDisableInWorkspace_ErrorForceDeleteHomeDir(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceUser(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			user_name = "abc"
			disable_as_user_deletion = true
			force_delete_home_dir = true
		`,
	}.ExpectError(t, "force_delete_home_dir: cannot force delete if disable_as_user_deletion is set")
}
