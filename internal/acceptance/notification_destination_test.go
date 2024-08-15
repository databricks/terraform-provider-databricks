package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkND(t *testing.T, display_name string, config_type settings.DestinationType) func(*terraform.State) error {
	return resourceCheck("databricks_notification_destination.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
		assert.Equal(t, config_type, ndResource.DestinationType)
		assert.Equal(t, display_name, ndResource.DisplayName)
		require.NoError(t, err)
		return nil
	})
}

func TestAccNDEmail(t *testing.T) {
	display_name := "Email Notification Destination - " + qa.RandomName()
	workspaceLevel(t, LegacyStep{
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
	}, LegacyStep{
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
		Check: checkND(t, display_name, settings.DestinationTypeEmail),
	})
}

func TestAccNDSlack(t *testing.T) {
	display_name := "Notification Destination - " + qa.RandomName()
	workspaceLevel(t, LegacyStep{
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
		Check: checkND(t, display_name, settings.DestinationTypeSlack),
	}, LegacyStep{
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
		Check: checkND(t, display_name, settings.DestinationTypeSlack),
	})
}

func TestAccNDMicrosoftTeams(t *testing.T) {
	display_name := "Notification Destination - " + qa.RandomName()
	workspaceLevel(t, LegacyStep{
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
	}, LegacyStep{
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
		Check: checkND(t, display_name, settings.DestinationTypeMicrosoftTeams),
	})
}

func TestAccNDPagerduty(t *testing.T) {
	display_name := "Notification Destination - " + qa.RandomName()
	workspaceLevel(t, LegacyStep{
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
	}, LegacyStep{
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
		Check: checkND(t, display_name, settings.DestinationTypePagerduty),
	})
}

func TestAccNDGenericWebhook(t *testing.T) {
	display_name := "Notification Destination - " + qa.RandomName()
	workspaceLevel(t, LegacyStep{
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
	}, LegacyStep{
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
		Check: checkND(t, display_name, settings.DestinationTypeWebhook),
	})
}

func TestAccConfigTypeChange(t *testing.T) {
	display_name := "Notification Destination - " + qa.RandomName()
	workspaceLevel(t, LegacyStep{
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
		Check: checkND(t, display_name, settings.DestinationTypeSlack),
	}, LegacyStep{
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
		Check: checkND(t, display_name, settings.DestinationTypeMicrosoftTeams),
	})
}
