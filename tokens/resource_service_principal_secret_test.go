package tokens

import (
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestServicePrincipalSecretCreate(t *testing.T) {
	expireTime := time.Now().Add(20 * time.Second).Format(time.RFC3339)
	currentTime := time.Now().Format(time.RFC3339)
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockServicePrincipalSecretsAPI().EXPECT()
			e.Create(mock.Anything, oauth2.CreateServicePrincipalSecretRequest{
				ServicePrincipalId: 123,
				Lifetime:           "20s",
			}).Return(&oauth2.CreateServicePrincipalSecretResponse{
				Secret:     "qwe",
				Id:         "003",
				Status:     "ACTIVE",
				ExpireTime: expireTime,
				CreateTime: currentTime,
				UpdateTime: currentTime,
				SecretHash: "abc",
			}, nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockServicePrincipalSecretsAPI().EXPECT()
			e.Delete(mock.Anything, oauth2.DeleteServicePrincipalSecretRequest{
				ServicePrincipalId: 123,
				SecretId:           "003",
			}).Return(nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockServicePrincipalSecretsAPI().EXPECT()
			e.ListByServicePrincipalId(mock.Anything, int64(123)).Return(&oauth2.ListServicePrincipalSecretsResponse{
				Secrets: []oauth2.SecretInfo{
					{
						Id:         "003",
						SecretHash: "abc",
						Status:     "ACTIVE",
					},
				},
			}, nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockServicePrincipalSecretsAPI().EXPECT()
			e.ListByServicePrincipalId(mock.Anything, int64(123)).Return(&oauth2.ListServicePrincipalSecretsResponse{
				Secrets: []oauth2.SecretInfo{
					{
						Id:         "004",
						SecretHash: "abc",
						Status:     "ACTIVE",
					},
				},
			}, nil)
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
