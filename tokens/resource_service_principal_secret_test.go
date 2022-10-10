package tokens

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestServicePrincipalSecretCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/xyz/servicePrincipals/abc/credentials/secrets",
				Response: ListServicePrincipalSecrets{
					Secrets: []ServicePrincipalSecret{
						{
							ID: "001",
						},
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/xyz/servicePrincipals/abc/credentials/secrets",
				Response: ServicePrincipalSecret{
					Secret: "qwe",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/xyz/servicePrincipals/abc/credentials/secrets",
				Response: ListServicePrincipalSecrets{
					Secrets: []ServicePrincipalSecret{
						{
							ID: "001",
						},
						{
							ID: "002",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/xyz/servicePrincipals/abc/credentials/secrets",
				Response: ListServicePrincipalSecrets{
					Secrets: []ServicePrincipalSecret{
						{
							ID:     "001",
							Status: "foo",
						},
						{
							ID:     "002",
							Status: "bar",
						},
					},
				},
			},
		},
		Resource:  ResourceServicePrincipalSecret(),
		Create:    true,
		AccountID: "xyz",
		HCL: `
		service_principal_id = "abc"
		`,
		New: true,
	}.ApplyAndExpectData(t, map[string]any{
		"secret": "qwe",
		"status": "bar",
	})
}

func TestServicePrincipalSecretDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/xyz/servicePrincipals/abc/credentials/secrets/003",
			},
		},
		Resource:  ResourceServicePrincipalSecret(),
		ID:        "003",
		Delete:    true,
		AccountID: "xyz",
		HCL: `
		service_principal_id = "abc"
		`,
	}.ApplyNoError(t)
}

func TestServicePrincipalSecretFuzz(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceServicePrincipalSecret(),
		qa.CornerCaseExpectError("must have `account_id` on provider"))
}
