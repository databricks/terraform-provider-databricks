package aws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/scim"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceServicePrincipalRoleCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: scim.PatchRequestWithValue(
					"add",
					"roles",
					"arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile"),
				Response: scim.User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=roles",
				Response: scim.User{
					Schemas:     []scim.URN{scim.ServicePrincipalSchema},
					DisplayName: "ABC SP",
					Roles: []scim.ComplexValue{
						{
							Value: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceServicePrincipalRole(),
		State: map[string]any{
			"service_principal_id": "abc",
			"role":                 "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{"id": "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile"})
}

func TestResourceServicePrincipalRoleCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceServicePrincipalRole(),
		State: map[string]any{
			"service_principal_id": "abc",
			"role":                 "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.ExpectError(t, "Internal error happened")
}

func TestResourceServicePrincipalRoleRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=roles",
				Response: scim.User{
					Schemas:     []scim.URN{scim.ServicePrincipalSchema},
					DisplayName: "ABC SP",
					Roles: []scim.ComplexValue{
						{
							Value: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceServicePrincipalRole(),
		Read:     true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.ApplyAndExpectData(t, map[string]any{"id": "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile"})
}

func TestResourceServicePrincipalRoleRead_NoRole(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=roles",
				Response: scim.User{
					Schemas:     []scim.URN{scim.ServicePrincipalSchema},
					DisplayName: "ABC SP",
					ID:          "abc",
				},
			},
		},
		Resource: ResourceServicePrincipalRole(),
		Read:     true,
		Removed:  true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalRoleRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc?attributes=roles",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceServicePrincipalRole(),
		Read:     true,
		Removed:  true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.ApplyNoError(t)
}
