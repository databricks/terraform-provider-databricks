package apps

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestResourceCustomAppIntegrationCreate(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockCustomAppIntegrationAPI().EXPECT()
			api.Create(mock.Anything, oauth2.CreateCustomAppIntegration{
				Name: "custom_integration_name",
				RedirectUrls: []string{
					"https://example.com",
				},
				Scopes: []string{
					"all",
				},
				TokenAccessPolicy: &oauth2.TokenAccessPolicy{
					AccessTokenTtlInMinutes:  60,
					RefreshTokenTtlInMinutes: 30,
				},
			}).Return(&oauth2.CreateCustomAppIntegrationOutput{
				ClientId:      "client_id",
				ClientSecret:  "client_secret",
				IntegrationId: "integration_id",
			}, nil)
			api.GetByIntegrationId(mock.Anything, "integration_id").Return(
				&oauth2.GetCustomAppIntegrationOutput{
					Name: "custom_integration_name",
					RedirectUrls: []string{
						"https://example.com",
					},
					Scopes: []string{
						"all",
					},
					TokenAccessPolicy: &oauth2.TokenAccessPolicy{
						AccessTokenTtlInMinutes:  60,
						RefreshTokenTtlInMinutes: 30,
					},
					ClientId:      "client_id",
					IntegrationId: "integration_id",
				}, nil,
			)
		},
		Create:    true,
		AccountID: "account_id",
		HCL: `
		name = "custom_integration_name"
		redirect_urls = ["https://example.com"]
		scopes = ["all"]
		token_access_policy {
			access_token_ttl_in_minutes = 60
			refresh_token_ttl_in_minutes = 30
		}`,
		Resource: ResourceCustomAppIntegration(),
	}.ApplyAndExpectData(t, map[string]any{
		"name":           "custom_integration_name",
		"integration_id": "integration_id",
		"client_id":      "client_id",
		"client_secret":  "client_secret",
	})
}

func TestResourceCustomAppIntegrationRead(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockCustomAppIntegrationAPI().EXPECT().GetByIntegrationId(mock.Anything, "integration_id").Return(
				&oauth2.GetCustomAppIntegrationOutput{
					Name: "custom_integration_name",
					RedirectUrls: []string{
						"https://example.com",
					},
					Scopes: []string{
						"all",
					},
					TokenAccessPolicy: &oauth2.TokenAccessPolicy{
						AccessTokenTtlInMinutes:  60,
						RefreshTokenTtlInMinutes: 30,
					},
					ClientId:      "client_id",
					IntegrationId: "integration_id",
				}, nil,
			)
		},
		Resource:  ResourceCustomAppIntegration(),
		Read:      true,
		New:       true,
		AccountID: "account_id",
		ID:        "integration_id",
	}.ApplyAndExpectData(t, map[string]any{
		"name":           "custom_integration_name",
		"integration_id": "integration_id",
		"client_id":      "client_id",
	})
}

func TestResourceCustomAppIntegrationUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockCustomAppIntegrationAPI().EXPECT()
			api.Update(mock.Anything, oauth2.UpdateCustomAppIntegration{
				RedirectUrls: []string{
					"https://example.com",
				},
				TokenAccessPolicy: &oauth2.TokenAccessPolicy{
					AccessTokenTtlInMinutes:  30,
					RefreshTokenTtlInMinutes: 30,
				},
			}).Return(nil)
			api.GetByIntegrationId(mock.Anything, "integration_id").Return(
				&oauth2.GetCustomAppIntegrationOutput{
					Name: "custom_integration_name",
					RedirectUrls: []string{
						"https://example.com",
					},
					Scopes: []string{
						"all",
					},
					TokenAccessPolicy: &oauth2.TokenAccessPolicy{
						AccessTokenTtlInMinutes:  30,
						RefreshTokenTtlInMinutes: 30,
					},
					ClientId:      "client_id",
					IntegrationId: "integration_id",
				}, nil,
			)
		},
		Resource: ResourceCustomAppIntegration(),
		Update:   true,
		HCL: `
		name = "custom_integration_name"
		redirect_urls = ["https://example.com"]
		scopes = ["all"]
		token_access_policy {
			access_token_ttl_in_minutes = 30
			refresh_token_ttl_in_minutes = 30
		}`,
		InstanceState: map[string]string{
			"name":            "custom_integration_name",
			"integration_id":  "integration_id",
			"client_id":       "client_id",
			"scopes.#":        "1",
			"scopes.0":        "all",
			"redirect_urls.#": "1",
			"redirect_urls.0": "https://example.com",
			"token_access_policy.access_token_ttl_in_minutes":  "30",
			"token_access_policy.refresh_token_ttl_in_minutes": "30",
		},
		AccountID: "account_id",
		ID:        "integration_id",
	}.ApplyAndExpectData(t, map[string]any{
		"name": "custom_integration_name",
		"token_access_policy.0.access_token_ttl_in_minutes": 30,
	})
}

func TestResourceCustomAppIntegrationDelete(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockCustomAppIntegrationAPI().EXPECT().DeleteByIntegrationId(mock.Anything, "integration_id").Return(nil)
		},
		Resource:  ResourceCustomAppIntegration(),
		AccountID: "account_id",
		Delete:    true,
		ID:        "integration_id",
	}.ApplyAndExpectData(t, nil)
}
