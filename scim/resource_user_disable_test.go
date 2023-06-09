package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/require"
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
			user_name    = "abc"
			disable_as_user_deletion = true
		`,
	}.ApplyNoError(t)
}

func TestResourceUserDeleteAsDisable_ForceDeleteRepos(t *testing.T) {
	_, err := qa.ResourceFixture{
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
			user_name    = "abc"
			disable_as_user_deletion = true
			force_delete_repos = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceUserDeleteAsDisable_ForceDeleteHomeDir(t *testing.T) {
	_, err := qa.ResourceFixture{
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
			user_name    = "abc"
			disable_as_user_deletion = true
			force_delete_home_dir = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceUserDeleteAsDisable_NoErrorEmptyParams(t *testing.T) {
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
			user_name    = "abc"
		`,
	}.ApplyNoError(t)
}
