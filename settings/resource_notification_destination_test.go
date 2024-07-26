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
						Addresses: []string{"abc@gmail.com"},
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
					addresses = ["abc@gmail.com"]
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
				DisplayName: "Notification Destination",
				Config: &settings.Config{
					Email: &settings.EmailConfig{
						Addresses: []string{"pqr@gmail.com"},
					},
				},
			}).Return(&settings.NotificationDestination{
				Id:              "xyz",
				DisplayName:     "Notification Destination",
				DestinationType: "EMAIL",
				Config: &settings.Config{
					Email: &settings.EmailConfig{
						Addresses: []string{"pqr@gmail.com"},
					},
				},
			}, nil)
			e.Get(mock.Anything, settings.GetNotificationDestinationRequest{
				Id: "xyz",
			}).Return(&settings.NotificationDestination{
				Id:              "xyz",
				DisplayName:     "Notification Destination",
				DestinationType: "EMAIL",
				Config: &settings.Config{
					Email: &settings.EmailConfig{
						Addresses: []string{"pqr@gmail.com"},
					},
				},
			}, nil)
		},
		Resource: ResourceNotificationDestination(),
		Update:   true,
		ID:       "xyz",
		HCL: `
			display_name = "Notification Destination"
			config {
				email {
					addresses = ["pqr@gmail.com"]
				}
			}
		`,
		InstanceState: map[string]string{
			"id":           "xyz",
			"display_name": "Notification Destination",
			"config":       `{"email":{"addresses":["abc@gmail.com"]}}`,
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id":           "xyz",
		"display_name": "Notification Destination",
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

// func TestNDConflictingFields(t *testing.T) {
// 	qa.ResourceFixture{
// 		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
// 			e := w.GetMockNotificationDestinationsAPI().EXPECT()
// 			e.Create(mock.Anything, settings.CreateNotificationDestinationRequest{
// 				DisplayName: "Notification Destination",
// 				Config: &settings.Config{
// 					GenericWebhook: &settings.GenericWebhookConfig{
// 						Url:      "https://webhook.site/abc",
// 						Password: "password",
// 					},
// 					Email: &settings.EmailConfig{
// 						Addresses: []string{"abc@gmail.com"},
// 					},
// 				},
// 			})
// 		},
// 		Resource: ResourceNotificationDestination(),
// 		Create:   true,
// 		HCL: `
// 			display_name = "Notification Destination"
// 			config {
// 				generic_webhook {
// 					url = "https://webhook.site/abc"
// 					password = "password"
// 				}
// 				email {
// 					addresses = ["abc@gmail.com"]
// 				}
// 			}
// 		`,
// 	}.ExpectError(t, "invalid config supplied. [config.#.email] Invalid combination of arguments. [config.#.generic_webhook] Invalid combination of arguments. [config.#.microsoft_teams] Invalid combination of arguments. [config.#.pagerduty] Invalid combination of arguments. [config.#.slack] Invalid combination of ars")
// }
