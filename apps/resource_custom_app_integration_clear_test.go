package apps

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	mockoauth2 "github.com/databricks/databricks-sdk-go/experimental/mocks/service/oauth2"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func customAppIntegrationState(extra map[string]string) map[string]string {
	state := map[string]string{
		"name":                  "custom_integration_name",
		"integration_id":        "integration_id",
		"client_id":             "client_id",
		"scopes.#":              "2",
		"scopes.0":              "all-apis",
		"scopes.1":              "sql",
		"redirect_urls.#":       "1",
		"redirect_urls.0":       "https://example.com",
		"token_access_policy.#": "1",
	}
	for k, v := range extra {
		state[k] = v
	}
	return state
}

var rotationOnState = map[string]string{
	"token_access_policy.0.access_token_ttl_in_minutes":          "30",
	"token_access_policy.0.refresh_token_ttl_in_minutes":         "720",
	"token_access_policy.0.enable_single_use_refresh_tokens":     "true",
	"token_access_policy.0.absolute_session_lifetime_in_minutes": "1440",
}

func expectCustomAppIntegrationGet(api *mockoauth2.MockCustomAppIntegrationInterface_Expecter, policy *oauth2.TokenAccessPolicy) {
	api.GetByIntegrationId(mock.Anything, "integration_id").Return(
		&oauth2.GetCustomAppIntegrationOutput{
			Name:              "custom_integration_name",
			RedirectUrls:      []string{"https://example.com"},
			Scopes:            []string{"all-apis", "sql"},
			TokenAccessPolicy: policy,
			ClientId:          "client_id",
			IntegrationId:     "integration_id",
		}, nil)
}

func TestResourceCustomAppIntegrationUpdateDisableRotation(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockCustomAppIntegrationAPI().EXPECT()
			api.Update(mock.Anything, oauth2.UpdateCustomAppIntegration{
				IntegrationId: "integration_id",
				RedirectUrls:  []string{"https://example.com"},
				Scopes:        []string{"all-apis", "sql"},
				TokenAccessPolicy: &oauth2.TokenAccessPolicy{
					AccessTokenTtlInMinutes:      30,
					RefreshTokenTtlInMinutes:     720,
					EnableSingleUseRefreshTokens: false,
					ForceSendFields:              []string{"EnableSingleUseRefreshTokens"},
				},
			}).Return(nil)
			expectCustomAppIntegrationGet(api, &oauth2.TokenAccessPolicy{
				AccessTokenTtlInMinutes:          30,
				RefreshTokenTtlInMinutes:         720,
				AbsoluteSessionLifetimeInMinutes: 1440,
			})
		},
		Resource: ResourceCustomAppIntegration(),
		Update:   true,
		HCL: `
		name = "custom_integration_name"
		redirect_urls = ["https://example.com"]
		scopes = ["all-apis", "sql"]
		token_access_policy {
			access_token_ttl_in_minutes = 30
			refresh_token_ttl_in_minutes = 720
			enable_single_use_refresh_tokens = false
		}`,
		InstanceState: customAppIntegrationState(rotationOnState),
		AccountID:     "account_id",
		Host:          "https://accounts.cloud.databricks.com",
		ID:            "integration_id",
	}.ApplyAndExpectData(t, map[string]any{
		"token_access_policy.0.enable_single_use_refresh_tokens": false,
	})
}

func TestResourceCustomAppIntegrationUpdateKeepsRotationFlagWhenUnchanged(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockCustomAppIntegrationAPI().EXPECT()
			api.Update(mock.Anything, oauth2.UpdateCustomAppIntegration{
				IntegrationId: "integration_id",
				RedirectUrls:  []string{"https://example.com"},
				Scopes:        []string{"all-apis", "sql"},
				TokenAccessPolicy: &oauth2.TokenAccessPolicy{
					AccessTokenTtlInMinutes:  15,
					RefreshTokenTtlInMinutes: 720,
				},
			}).Return(nil)
			expectCustomAppIntegrationGet(api, &oauth2.TokenAccessPolicy{
				AccessTokenTtlInMinutes:  15,
				RefreshTokenTtlInMinutes: 720,
			})
		},
		Resource: ResourceCustomAppIntegration(),
		Update:   true,
		HCL: `
		name = "custom_integration_name"
		redirect_urls = ["https://example.com"]
		scopes = ["all-apis", "sql"]
		token_access_policy {
			access_token_ttl_in_minutes = 15
			refresh_token_ttl_in_minutes = 720
		}`,
		InstanceState: customAppIntegrationState(map[string]string{
			"token_access_policy.0.access_token_ttl_in_minutes":      "30",
			"token_access_policy.0.refresh_token_ttl_in_minutes":     "720",
			"token_access_policy.0.enable_single_use_refresh_tokens": "false",
		}),
		AccountID: "account_id",
		Host:      "https://accounts.cloud.databricks.com",
		ID:        "integration_id",
	}.ApplyNoError(t)
}

func TestResourceCustomAppIntegrationClearUserAuthorizedScopesFails(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceCustomAppIntegration(),
		Update:   true,
		HCL: `
		name = "custom_integration_name"
		redirect_urls = ["https://example.com"]
		scopes = ["all-apis", "sql"]
		token_access_policy {
			access_token_ttl_in_minutes = 30
			refresh_token_ttl_in_minutes = 720
			enable_single_use_refresh_tokens = true
		}`,
		InstanceState: customAppIntegrationState(map[string]string{
			"user_authorized_scopes.#":                                   "1",
			"user_authorized_scopes.0":                                   "sql",
			"token_access_policy.0.access_token_ttl_in_minutes":          "30",
			"token_access_policy.0.refresh_token_ttl_in_minutes":         "720",
			"token_access_policy.0.enable_single_use_refresh_tokens":     "true",
			"token_access_policy.0.absolute_session_lifetime_in_minutes": "1440",
		}),
		AccountID: "account_id",
		Host:      "https://accounts.cloud.databricks.com",
		ID:        "integration_id",
	}.ExpectError(t, "user_authorized_scopes can't be removed from an existing custom app integration, "+
		"because the API keeps the current scopes [sql] when an empty list is sent. "+
		"Keep at least one scope, or recreate the integration (this issues a new client ID and secret)")
}

func TestResourceCustomAppIntegrationRemovePolicyResetsToDefaults(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockCustomAppIntegrationAPI().EXPECT()
			api.Update(mock.Anything, oauth2.UpdateCustomAppIntegration{
				IntegrationId: "integration_id",
				RedirectUrls:  []string{"https://example.com"},
				Scopes:        []string{"all-apis", "sql"},
				TokenAccessPolicy: &oauth2.TokenAccessPolicy{
					AccessTokenTtlInMinutes:  60,
					RefreshTokenTtlInMinutes: 10080,
					ForceSendFields:          []string{"EnableSingleUseRefreshTokens"},
				},
			}).Return(nil)
			expectCustomAppIntegrationGet(api, &oauth2.TokenAccessPolicy{
				AccessTokenTtlInMinutes:          60,
				RefreshTokenTtlInMinutes:         10080,
				AbsoluteSessionLifetimeInMinutes: 1440,
			})
		},
		Resource: ResourceCustomAppIntegration(),
		Update:   true,
		HCL: `
		name = "custom_integration_name"
		redirect_urls = ["https://example.com"]
		scopes = ["all-apis", "sql"]`,
		InstanceState: customAppIntegrationState(rotationOnState),
		AccountID:     "account_id",
		Host:          "https://accounts.cloud.databricks.com",
		ID:            "integration_id",
	}.ApplyAndExpectData(t, map[string]any{
		"token_access_policy.0.access_token_ttl_in_minutes":      60,
		"token_access_policy.0.refresh_token_ttl_in_minutes":     10080,
		"token_access_policy.0.enable_single_use_refresh_tokens": false,
	})
}

func TestResourceCustomAppIntegrationNoPolicyNoDiffWithServerDefaults(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceCustomAppIntegration(),
		HCL: `
		name = "custom_integration_name"
		redirect_urls = ["https://example.com"]
		scopes = ["all-apis", "sql"]`,
		InstanceState: customAppIntegrationState(map[string]string{
			"token_access_policy.0.access_token_ttl_in_minutes":          "60",
			"token_access_policy.0.refresh_token_ttl_in_minutes":         "10080",
			"token_access_policy.0.enable_single_use_refresh_tokens":     "false",
			"token_access_policy.0.absolute_session_lifetime_in_minutes": "129600",
			"client_secret":    "secret",
			"create_time":      "2026-11-19T14:37:43.700Z",
			"created_by":       "100",
			"creator_username": "someone",
		}),
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{},
		AccountID:    "account_id",
		Host:         "https://accounts.cloud.databricks.com",
		ID:           "integration_id",
	}.ApplyNoError(t)
}
