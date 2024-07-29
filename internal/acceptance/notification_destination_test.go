package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccNDEmail(t *testing.T) {
	display_name := "Email Notification Destination - " + qa.RandomName()
	workspaceLevel(t, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "` + display_name + `"
			config {
				email {
					addresses = ["` + qa.RandomEmail() + `"]
				}
			}
		}
		`,
	}, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "` + display_name + `"
			config {
				email {
					addresses = ["` + qa.RandomEmail() + `", "` + qa.RandomEmail() + `"]
				}
			}
		}
		`,
		Check: resourceCheck("databricks_notification_destination.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			ndResource, err := w.NotificationDestinations.Get(ctx, settings.GetNotificationDestinationRequest{
				Id: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, settings.DestinationType("EMAIL"), ndResource.DestinationType)
			assert.Equal(t, display_name, ndResource.DisplayName)
			require.NoError(t, err)
			return nil
		}),
	})
}

func TestAccNDSlack(t *testing.T) {
	display_name := "Notification Destination - " + qa.RandomName()
	workspaceLevel(t, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "` + display_name + `"
			config {
				slack {
					url = "https://hooks.slack.com/services/{var.RANDOM}"
				}
			}
		}
		`,
		Check: resourceCheck("databricks_notification_destination.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			ndResource, err := w.NotificationDestinations.Get(ctx, settings.GetNotificationDestinationRequest{
				Id: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, settings.DestinationType("SLACK"), ndResource.DestinationType)
			assert.Equal(t, display_name, ndResource.DisplayName)
			require.NoError(t, err)
			return nil
		}),
	}, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "` + display_name + `"
			config {
				slack {
					url = "https://hooks.slack.com/services/{var.RANDOM}"
				}
			}
		}
		`,
		Check: resourceCheck("databricks_notification_destination.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			ndResource, err := w.NotificationDestinations.Get(ctx, settings.GetNotificationDestinationRequest{
				Id: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, settings.DestinationType("SLACK"), ndResource.DestinationType)
			assert.Equal(t, display_name, ndResource.DisplayName)
			require.NoError(t, err)
			return nil
		}),
	})
}

func TestAccNDMicrosoftTeams(t *testing.T) {
	display_name := "Notification Destination - " + qa.RandomName()
	workspaceLevel(t, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "` + display_name + `"
			config {
				microsoft_teams {
					url = "https://outlook.office.com/webhook/{var.RANDOM}"
				}
			}
		}
		`,
	}, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "` + display_name + `"
			config {
				microsoft_teams {
					url = "https://outlook.office.com/webhook/{var.RANDOM}"
				}
			}
		}
		`,
		Check: resourceCheck("databricks_notification_destination.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			ndResource, err := w.NotificationDestinations.Get(ctx, settings.GetNotificationDestinationRequest{
				Id: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, settings.DestinationType("MICROSOFT_TEAMS"), ndResource.DestinationType)
			assert.Equal(t, display_name, ndResource.DisplayName)
			require.NoError(t, err)
			return nil
		}),
	})
}

func TestAccNDPagerduty(t *testing.T) {
	display_name := "Notification Destination - " + qa.RandomName()
	workspaceLevel(t, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "` + display_name + `"
			config {
				pagerduty {
					integration_key = "{var.RANDOM}"
				}
			}
		}
		`,
	}, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "` + display_name + `"
			config {
				pagerduty {
					integration_key = "{var.RANDOM}"
				}
			}
		}
		`,
		Check: resourceCheck("databricks_notification_destination.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			ndResource, err := w.NotificationDestinations.Get(ctx, settings.GetNotificationDestinationRequest{
				Id: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, settings.DestinationType("PAGERDUTY"), ndResource.DestinationType)
			assert.Equal(t, display_name, ndResource.DisplayName)
			require.NoError(t, err)
			return nil
		}),
	})
}

func TestAccNDGenericWebhook(t *testing.T) {
	display_name := "Notification Destination - " + qa.RandomName()
	workspaceLevel(t, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "` + display_name + `"
			config {
				generic_webhook {
					url = "https://webhook.site/{var.RANDOM}"
					password = "password"
				}
			}
		}
		`,
	}, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "` + display_name + `"
			config {
				generic_webhook {
					url = "https://webhook.site/{var.RANDOM}"
					username = "username2"
				}
			}
		}
		`,
		Check: resourceCheck("databricks_notification_destination.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			ndResource, err := w.NotificationDestinations.Get(ctx, settings.GetNotificationDestinationRequest{
				Id: id,
			})
			if err != nil {
				return err
			}
			assert.Equal(t, settings.DestinationType("WEBHOOK"), ndResource.DestinationType)
			assert.Equal(t, display_name, ndResource.DisplayName)
			require.NoError(t, err)
			return nil
		}),
	})
}
