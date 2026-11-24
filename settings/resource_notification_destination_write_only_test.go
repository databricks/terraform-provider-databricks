package settings

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNDCreateWriteOnly(t *testing.T) {
	for _, tc := range []struct {
		name     string
		hcl      string
		config   settings.Config
		readBack settings.Config
		cleared  []string
	}{
		{
			name: "generic_webhook",
			hcl: `generic_webhook {
				url_wo      = "https://webhook.site/abc?token=secret"
				username    = "svc"
				password_wo = "password"
			}`,
			config: settings.Config{GenericWebhook: &settings.GenericWebhookConfig{
				Url: "https://webhook.site/abc?token=secret", Username: "svc", Password: "password"}},
			readBack: settings.Config{GenericWebhook: &settings.GenericWebhookConfig{
				UrlSet: true, UsernameSet: true, PasswordSet: true}},
			cleared: []string{"config.0.generic_webhook.0.url", "config.0.generic_webhook.0.password"},
		},
		{
			name: "generic_webhook_username",
			hcl: `generic_webhook {
				url_wo      = "https://webhook.site/abc"
				username_wo = "svc"
			}`,
			config: settings.Config{GenericWebhook: &settings.GenericWebhookConfig{
				Url: "https://webhook.site/abc", Username: "svc"}},
			readBack: settings.Config{GenericWebhook: &settings.GenericWebhookConfig{
				UrlSet: true, UsernameSet: true}},
			cleared: []string{"config.0.generic_webhook.0.url", "config.0.generic_webhook.0.username"},
		},
		{
			name: "slack",
			hcl: `slack {
				url_wo         = "https://hooks.slack.com/services/T/B/secret"
				oauth_token_wo = "xoxb-secret"
				channel_id     = "C123"
			}`,
			config: settings.Config{Slack: &settings.SlackConfig{
				Url: "https://hooks.slack.com/services/T/B/secret", OauthToken: "xoxb-secret", ChannelId: "C123"}},
			readBack: settings.Config{Slack: &settings.SlackConfig{
				UrlSet: true, OauthTokenSet: true, ChannelIdSet: true}},
			cleared: []string{"config.0.slack.0.url", "config.0.slack.0.oauth_token"},
		},
		{
			name: "microsoft_teams",
			hcl: `microsoft_teams {
				url_wo         = "https://outlook.office.com/webhook/secret"
				auth_secret_wo = "teams-secret"
			}`,
			config: settings.Config{MicrosoftTeams: &settings.MicrosoftTeamsConfig{
				Url: "https://outlook.office.com/webhook/secret", AuthSecret: "teams-secret"}},
			readBack: settings.Config{MicrosoftTeams: &settings.MicrosoftTeamsConfig{
				UrlSet: true, AuthSecretSet: true}},
			cleared: []string{"config.0.microsoft_teams.0.url", "config.0.microsoft_teams.0.auth_secret"},
		},
		{
			name: "pagerduty",
			hcl: `pagerduty {
				integration_key_wo = "pd-secret"
			}`,
			config: settings.Config{Pagerduty: &settings.PagerdutyConfig{IntegrationKey: "pd-secret"}},
			readBack: settings.Config{Pagerduty: &settings.PagerdutyConfig{
				IntegrationKeySet: true}},
			cleared: []string{"config.0.pagerduty.0.integration_key"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config, readBack := tc.config, tc.readBack
			d, err := qa.ResourceFixture{
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					e := w.GetMockNotificationDestinationsAPI().EXPECT()
					e.Create(mock.Anything, settings.CreateNotificationDestinationRequest{
						DisplayName: "Notification Destination",
						Config:      &config,
					}).Return(&settings.NotificationDestination{Id: "xyz"}, nil)
					e.Get(mock.Anything, settings.GetNotificationDestinationRequest{Id: "xyz"}).Return(
						&settings.NotificationDestination{
							Id:          "xyz",
							DisplayName: "Notification Destination",
							Config:      &readBack,
						}, nil)
				},
				Resource: ResourceNotificationDestination(),
				Create:   true,
				HCL: `
					display_name      = "Notification Destination"
					config_wo_version = 1
					config {
						` + tc.hcl + `
					}`,
			}.Apply(t)
			assert.NoError(t, err)
			assert.Equal(t, 1, d.Get("config_wo_version"))
			for _, k := range tc.cleared {
				assert.Equal(t, "", d.Get(k), k)
			}
		})
	}
}

func TestNDUpdateResendsWriteOnly(t *testing.T) {
	// The API replaces the whole config block on update, so the write-only values must be
	// sent with every update, also when only another field changes.
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateNotificationDestinationRequest{
				Id:          "xyz",
				DisplayName: "Notification Destination - 2",
				Config: &settings.Config{
					GenericWebhook: &settings.GenericWebhookConfig{
						Url:         "https://webhook.site/abc",
						Password:    "password",
						UrlSet:      true,
						PasswordSet: true,
					},
				},
			}).Return(&settings.NotificationDestination{Id: "xyz"}, nil)
			e.Get(mock.Anything, settings.GetNotificationDestinationRequest{Id: "xyz"}).Return(
				&settings.NotificationDestination{
					Id:          "xyz",
					DisplayName: "Notification Destination - 2",
					Config: &settings.Config{
						GenericWebhook: &settings.GenericWebhookConfig{UrlSet: true, PasswordSet: true},
					},
				}, nil)
		},
		Resource: ResourceNotificationDestination(),
		Update:   true,
		ID:       "xyz",
		HCL: `
			display_name      = "Notification Destination - 2"
			config_wo_version = 1
			config {
				generic_webhook {
					url_wo      = "https://webhook.site/abc"
					password_wo = "password"
				}
			}`,
		InstanceState: map[string]string{
			"id":                                 "xyz",
			"display_name":                       "Notification Destination",
			"config_wo_version":                  "1",
			"config.#":                           "1",
			"config.0.generic_webhook.#":         "1",
			"config.0.generic_webhook.0.url_set": "true",
			"config.0.generic_webhook.0.password_set": "true",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"display_name":                        "Notification Destination - 2",
		"config.0.generic_webhook.0.url":      "",
		"config.0.generic_webhook.0.password": "",
	})
}

func TestNDSwitchToWriteOnlyClearsState(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateNotificationDestinationRequest{
				Id:          "xyz",
				DisplayName: "Notification Destination",
				Config: &settings.Config{
					GenericWebhook: &settings.GenericWebhookConfig{
						Url:         "https://webhook.site/new",
						Password:    "new-password",
						UrlSet:      true,
						PasswordSet: true,
					},
				},
			}).Return(&settings.NotificationDestination{Id: "xyz"}, nil)
			e.Get(mock.Anything, settings.GetNotificationDestinationRequest{Id: "xyz"}).Return(
				&settings.NotificationDestination{
					Id:          "xyz",
					DisplayName: "Notification Destination",
					Config: &settings.Config{
						GenericWebhook: &settings.GenericWebhookConfig{UrlSet: true, PasswordSet: true},
					},
				}, nil)
		},
		Resource: ResourceNotificationDestination(),
		Update:   true,
		ID:       "xyz",
		HCL: `
			display_name      = "Notification Destination"
			config_wo_version = 1
			config {
				generic_webhook {
					url_wo      = "https://webhook.site/new"
					password_wo = "new-password"
				}
			}`,
		InstanceState: map[string]string{
			"id":                                      "xyz",
			"display_name":                            "Notification Destination",
			"config.#":                                "1",
			"config.0.generic_webhook.#":              "1",
			"config.0.generic_webhook.0.url":          "https://webhook.site/old",
			"config.0.generic_webhook.0.url_set":      "true",
			"config.0.generic_webhook.0.password":     "old-password",
			"config.0.generic_webhook.0.password_set": "true",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"config_wo_version":                   1,
		"config.0.generic_webhook.0.url":      "",
		"config.0.generic_webhook.0.password": "",
	})
}

func TestNDWriteOnlyValidation(t *testing.T) {
	for _, tc := range []struct {
		name, hcl, err string
	}{
		{
			name: "url_missing",
			hcl:  `config_wo_version = 1` + "\n" + `config { generic_webhook { username = "svc" } }`,
			err:  "[config.#.generic_webhook.#.url] Invalid combination of arguments",
		},
		{
			name: "url_and_url_wo",
			hcl:  `config_wo_version = 1` + "\n" + `config { generic_webhook { url = "a" url_wo = "b" } }`,
			err:  "[config.#.generic_webhook.#.url] Invalid combination of arguments. [config.#.generic_webhook.#.url_wo] Conflicting configuration arguments",
		},
		{
			name: "password_and_password_wo",
			hcl:  `config_wo_version = 1` + "\n" + `config { generic_webhook { url_wo = "a" password = "b" password_wo = "c" } }`,
			err:  "[config.#.generic_webhook.#.password] Conflicting configuration arguments. [config.#.generic_webhook.#.password_wo] Conflicting configuration arguments",
		},
		{
			name: "write_only_without_version",
			hcl:  `config { slack { url_wo = "a" } }`,
			err:  "[config.#.slack.#.url_wo] Missing required argument",
		},
		{
			name: "integration_key_missing",
			hcl:  `config_wo_version = 1` + "\n" + `config { pagerduty { } }`,
			err:  "[config.#.pagerduty.#.integration_key] Invalid combination of arguments",
		},
		{
			name: "version_below_one",
			hcl:  `config_wo_version = 0` + "\n" + `config { slack { url_wo = "a" } }`,
			err:  "[config_wo_version] expected config_wo_version to be at least (1), got 0",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			qa.ResourceFixture{
				Resource: ResourceNotificationDestination(),
				Create:   true,
				HCL:      `display_name = "Notification Destination"` + "\n" + tc.hcl,
			}.ExpectError(t, "invalid config supplied. "+tc.err)
		})
	}
}

func TestApplyWriteOnlyConfigNilConfig(t *testing.T) {
	d := ResourceNotificationDestination().ToResource().TestResourceData()
	assert.NotPanics(t, func() { applyWriteOnlyConfig(d, nil) })
}
