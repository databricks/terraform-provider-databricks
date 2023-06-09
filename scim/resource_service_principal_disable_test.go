package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
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
		AccountID: "00000000-0000-0000-0000-000000000001",
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/accounts/00000000-0000-0000-0000-000000000001/scim/v2/ServicePrincipals/abc",
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

func TestResourceServicePrincipalDeleteAsDisable_NoErrorEmptyParams(t *testing.T) {
	qa.ResourceFixture{
		AccountID: "00000000-0000-0000-0000-000000000001",
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PATCH",
				Resource:        "/api/2.0/accounts/00000000-0000-0000-0000-000000000001/scim/v2/ServicePrincipals/abc",
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

func TestResourceServicePrincipalDeleteAsDisable_IgnoreIfNotAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
			},
		},
		Resource: ResourceServicePrincipal(),
		Delete:   true,
		ID:       "abc",
		HCL: `
			application_id = "abc"
			disable_as_user_deletion = true
		`,
	}.ApplyNoError(t)
}
