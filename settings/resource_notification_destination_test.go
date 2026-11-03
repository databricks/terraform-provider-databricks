package settings

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/mock"
)

func TestNDCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI().EXPECT()
			e.Create(mock.Anything, settings.CreateNotificationDestinationRequest{
				DisplayName: "Notification Destination",
				Config: &settings.Config{
					GenericWebhook: &settings.GenericWebhookConfig{
						Url:      "https://webhook.site/abc",
						Password: "password",
					},
				},
			}).Return(&settings.NotificationDestination{
				Id:              "xyz",
				DisplayName:     "Notification Destination",
				DestinationType: "WEBHOOK",
				Config: &settings.Config{
					GenericWebhook: &settings.GenericWebhookConfig{
						UrlSet:      true,
						PasswordSet: true,
					},
				},
			}, nil)
			e.Get(mock.Anything, settings.GetNotificationDestinationRequest{
				Id: "xyz",
			}).Return(&settings.NotificationDestination{
				Id:              "xyz",
				DisplayName:     "Notification Destination",
				DestinationType: "WEBHOOK",
				Config: &settings.Config{
					GenericWebhook: &settings.GenericWebhookConfig{
						UrlSet:      true,
						PasswordSet: true,
					},
				},
			}, nil)
		},
		Resource: ResourceNotificationDestination(),
		Create:   true,
		HCL: `
			display_name = "Notification Destination"
			config {
				generic_webhook {
					url = "https://webhook.site/abc"
					password = "password"
				}
			}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           "xyz",
		"display_name": "Notification Destination",
	})
}

func TestNDRead(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockNotificationDestinationsAPI().EXPECT().Get(mock.Anything, settings.GetNotificationDestinationRequest{
				Id: "xyz",
			}).Return(&settings.NotificationDestination{
				Id:              "xyz",
				DisplayName:     "Notification Destination",
				DestinationType: "EMAIL",
				Config: &settings.Config{
					Email: &settings.EmailConfig{
						Addresses: []string{"abc@email.com"},
					},
				},
			}, nil)
		},
		Resource: ResourceNotificationDestination(),
		Read:     true,
		ID:       "xyz",
		HCL: `
			display_name = "Notification Destination"
			config {
				email {
					addresses = ["abc@email.com"]
				}
			}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           "xyz",
		"display_name": "Notification Destination",
	})
}

func TestNDUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateNotificationDestinationRequest{
				Id:          "xyz",
				DisplayName: "Notification Destination - 2",
				Config: &settings.Config{
					Email: &settings.EmailConfig{
						Addresses: []string{"pqr@email.com"},
					},
				},
			}).Return(&settings.NotificationDestination{
				Id:              "xyz",
				DisplayName:     "Notification Destination - 2",
				DestinationType: "EMAIL",
				Config: &settings.Config{
					Email: &settings.EmailConfig{
						Addresses: []string{"pqr@email.com"},
					},
				},
			}, nil)
			e.Get(mock.Anything, settings.GetNotificationDestinationRequest{
				Id: "xyz",
			}).Return(&settings.NotificationDestination{
				Id:              "xyz",
				DisplayName:     "Notification Destination - 2",
				DestinationType: "EMAIL",
				Config: &settings.Config{
					Email: &settings.EmailConfig{
						Addresses: []string{"pqr@email.com"},
					},
				},
			}, nil)
		},
		Resource: ResourceNotificationDestination(),
		Update:   true,
		ID:       "xyz",
		HCL: `
			display_name = "Notification Destination - 2"
			config {
				email {
					addresses = ["pqr@email.com"]
				}
			}
		`,
		InstanceState: map[string]string{
			"id":                           "xyz",
			"display_name":                 "Notification Destination",
			"config.#":                     "1",
			"config.0.email.#":             "1",
			"config.0.email.0.addresses.#": "1",
			"config.0.email.0.addresses.0": "abc@email.com",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id":           "xyz",
		"display_name": "Notification Destination - 2",
	})
}

func TestNDDelete(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockNotificationDestinationsAPI().EXPECT().Delete(mock.Anything, settings.DeleteNotificationDestinationRequest{
				Id: "xyz",
			}).Return(nil)
		},
		Resource: ResourceNotificationDestination(),
		Delete:   true,
		ID:       "xyz",
		HCL: `
			display_name = "Notification Destination"
			config {
				generic_webhook {
					url = "https://webhook.site/abc"
					password = "password"
				}
			}
		`,
	}.ApplyNoError(t)
}

func TestNDCreateDatabricksAppSlack(t *testing.T) {
	nd := &settings.NotificationDestination{
		Id:              "xyz",
		DisplayName:     "Notification Destination",
		DestinationType: settings.DestinationTypeDatabricksAppSlack,
		Config: &settings.Config{
			DatabricksAppSlack: &settings.DatabricksAppSlackConfig{
				ChannelIdSet: true,
				TeamIdSet:    true,
			},
		},
	}
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI().EXPECT()
			e.Create(mock.Anything, settings.CreateNotificationDestinationRequest{
				DisplayName: "Notification Destination",
				Config: &settings.Config{
					DatabricksAppSlack: &settings.DatabricksAppSlackConfig{
						ChannelId: "C0123456789",
						TeamId:    "T0123456789",
					},
				},
			}).Return(nd, nil)
			e.Get(mock.Anything, settings.GetNotificationDestinationRequest{Id: "xyz"}).Return(nd, nil)
		},
		Resource: ResourceNotificationDestination(),
		Create:   true,
		HCL: `
			display_name = "Notification Destination"
			config {
				databricks_app_slack {
					channel_id = "C0123456789"
					team_id    = "T0123456789"
				}
			}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":               "xyz",
		"destination_type": "DATABRICKS_APP_SLACK",
		"config.0.databricks_app_slack.0.channel_id":     "C0123456789",
		"config.0.databricks_app_slack.0.team_id":        "T0123456789",
		"config.0.databricks_app_slack.0.channel_id_set": true,
		"config.0.databricks_app_slack.0.team_id_set":    true,
	})
}

func TestNDCreateDatabricksAppTeams(t *testing.T) {
	nd := &settings.NotificationDestination{
		Id:              "xyz",
		DisplayName:     "Notification Destination",
		DestinationType: settings.DestinationTypeDatabricksAppTeams,
		Config: &settings.Config{
			DatabricksAppTeams: &settings.DatabricksAppTeamsConfig{
				ChannelIdSet: true,
				TeamIdSet:    true,
				TenantIdSet:  true,
			},
		},
	}
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockNotificationDestinationsAPI().EXPECT()
			e.Create(mock.Anything, settings.CreateNotificationDestinationRequest{
				DisplayName: "Notification Destination",
				Config: &settings.Config{
					DatabricksAppTeams: &settings.DatabricksAppTeamsConfig{
						ChannelId: "19:abc@thread.tacv2",
						TeamId:    "team-guid",
						TenantId:  "tenant-guid",
					},
				},
			}).Return(nd, nil)
			e.Get(mock.Anything, settings.GetNotificationDestinationRequest{Id: "xyz"}).Return(nd, nil)
		},
		Resource: ResourceNotificationDestination(),
		Create:   true,
		HCL: `
			display_name = "Notification Destination"
			config {
				databricks_app_teams {
					channel_id = "19:abc@thread.tacv2"
					team_id    = "team-guid"
					tenant_id  = "tenant-guid"
				}
			}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":               "xyz",
		"destination_type": "DATABRICKS_APP_TEAMS",
		"config.0.databricks_app_teams.0.channel_id":    "19:abc@thread.tacv2",
		"config.0.databricks_app_teams.0.team_id":       "team-guid",
		"config.0.databricks_app_teams.0.tenant_id":     "tenant-guid",
		"config.0.databricks_app_teams.0.tenant_id_set": true,
	})
}

func TestNDReadKeepsInputOnlyDatabricksAppFields(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockNotificationDestinationsAPI().EXPECT().Get(mock.Anything, settings.GetNotificationDestinationRequest{
				Id: "xyz",
			}).Return(&settings.NotificationDestination{
				Id:              "xyz",
				DisplayName:     "Notification Destination",
				DestinationType: settings.DestinationTypeDatabricksAppSlack,
				Config: &settings.Config{
					DatabricksAppSlack: &settings.DatabricksAppSlackConfig{
						ChannelIdSet: true,
						TeamIdSet:    true,
					},
				},
			}, nil)
		},
		Resource: ResourceNotificationDestination(),
		Read:     true,
		ID:       "xyz",
		HCL: `
			display_name = "Notification Destination"
			config {
				databricks_app_slack {
					channel_id = "C0123456789"
					team_id    = "T0123456789"
				}
			}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"config.0.databricks_app_slack.0.channel_id": "C0123456789",
		"config.0.databricks_app_slack.0.team_id":    "T0123456789",
	})
}

func TestNDDatabricksAppRequiresExactlyOneConfig(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceNotificationDestination(),
		Create:   true,
		HCL: `
			display_name = "Notification Destination"
			config {
				databricks_app_slack {
					channel_id = "C0123456789"
					team_id    = "T0123456789"
				}
				email {
					addresses = ["abc@email.com"]
				}
			}
		`,
	}.ExpectError(t, "invalid config supplied. [config.#.slack] Invalid combination of arguments")
}
