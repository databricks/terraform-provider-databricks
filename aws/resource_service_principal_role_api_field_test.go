package aws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/scim"
)

func TestResourceServicePrincipalRoleCreate_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/ServicePrincipals/abc",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/ServicePrincipals/abc?attributes=roles",
				Response: scim.User{
					Roles: []scim.ComplexValue{{Value: "arn:aws:iam::999999999999:role/foo"}},
				},
			},
		},
		Resource:  ResourceServicePrincipalRole(),
		AccountID: "acc-123",
		HCL: `
			service_principal_id = "abc"
			role = "arn:aws:iam::999999999999:role/foo"
			api = "account"
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalRoleRead_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/acc-123/scim/v2/ServicePrincipals/abc?attributes=roles",
				Response: scim.User{
					Roles: []scim.ComplexValue{{Value: "arn:aws:iam::999999999999:role/foo"}},
				},
			},
		},
		Resource:  ResourceServicePrincipalRole(),
		AccountID: "acc-123",
		HCL: `
			service_principal_id = "abc"
			role = "arn:aws:iam::999999999999:role/foo"
			api = "account"
		`,
		New:  true,
		Read: true,
		ID:   "abc|arn:aws:iam::999999999999:role/foo",
	}.ApplyNoError(t)
}
