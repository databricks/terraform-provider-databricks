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
