package tokens

import (
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/databricks/databricks-sdk-go/service/oauth2"
)

func TestServicePrincipalSecretCreate(t *testing.T) {
	expireTime := time.Now().Add(20 * time.Second).Format(time.RFC3339)
	currentTime := time.Now().Format(time.RFC3339)
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/xyz/servicePrincipals/123/credentials/secrets",
				ExpectedRequest: oauth2.CreateServicePrincipalSecretRequest{
					ServicePrincipalId: 123,
					Lifetime:           "20s",
				},
				Response: oauth2.CreateServicePrincipalSecretResponse{
					Secret:     "qwe",
					Id:         "003",
					Status:     "ACTIVE",
					ExpireTime: expireTime,
					CreateTime: currentTime,
					UpdateTime: currentTime,
					SecretHash: "abc",
				},
			},
		},
		Resource:  ResourceServicePrincipalSecret(),
		Create:    true,
		AccountID: "xyz",
		HCL: `
		service_principal_id = "123"
		lifetime  	  		 = "20s"
		`,
		New: true,
	}.ApplyAndExpectData(t, map[string]any{
		"secret":               "qwe",
		"status":               "ACTIVE",
		"service_principal_id": "123",
		"id":                   "003",
		"expire_time":          expireTime,
		"create_time":          currentTime,
		"update_time":          currentTime,
		"secret_hash":          "abc",
		"lifetime":             "20s",
	})
}

func TestServicePrincipalSecretDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/xyz/servicePrincipals/123/credentials/secrets/003?",
			},
		},
		Resource:  ResourceServicePrincipalSecret(),
		ID:        "003",
		Delete:    true,
		AccountID: "xyz",
		HCL: `
		service_principal_id = "123"
		`,
	}.ApplyNoError(t)
}

func TestServicePrincipalSecretRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/xyz/servicePrincipals/123/credentials/secrets?",
				Response: oauth2.ListServicePrincipalSecretsResponse{
					Secrets: []oauth2.SecretInfo{
						{
							Id:         "003",
							SecretHash: "abc",
							Status:     "ACTIVE",
						},
					},
				},
			},
		},
		Resource:  ResourceServicePrincipalSecret(),
		ID:        "003",
		Read:      true,
		AccountID: "xyz",
		HCL: `
		service_principal_id = "123"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"secret_hash": "abc",
		"status":      "ACTIVE",
	})
}

func TestServicePrincipalSecretReadRemoved(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/xyz/servicePrincipals/123/credentials/secrets?",
				Response: oauth2.ListServicePrincipalSecretsResponse{
					Secrets: []oauth2.SecretInfo{
						{
							Id:         "004",
							SecretHash: "abc",
							Status:     "ACTIVE",
						},
					},
				},
			},
		},
		Resource:  ResourceServicePrincipalSecret(),
		ID:        "003",
		Read:      true,
		Removed:   true,
		AccountID: "xyz",
		HCL: `
		service_principal_id = "123"
		`,
	}.ApplyNoError(t)
}
