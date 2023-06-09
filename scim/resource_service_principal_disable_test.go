package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/require"
)

var expectedServicePrincipalDisablePatchRequest = patchRequest{
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

func TestResourceServicePrincipalDeleteAsDisable_NoError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: expectedServicePrincipalDisablePatchRequest,
			},
		},
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			application_id    = "abc"
			disable_as_user_deletion = true
		`,
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalDeleteAsDisable_ErrorForceDeleteRepos(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			application_id    = "abc"
			disable_as_user_deletion = true
			force_delete_repos = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceServicePrincipalDeleteAsDisable_ErrorForceDeleteHomeDir(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			application_id    = "abc"
			disable_as_user_deletion = true
			force_delete_home_dir = true
		`,
	}.Apply(t)
	require.Error(t, err, err)
}

func TestResourceServicePrincipalDeleteAsDisable_NoErrorEmptyParams(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: expectedServicePrincipalDisablePatchRequest,
			},
		},
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			application_id    = "abc"
		`,
	}.ApplyNoError(t)
}
