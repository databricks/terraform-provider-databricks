package notificationdestinations_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CheckDataSourceNotificationsPopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.Modules[0].Resources["data.databricks_notification_destinations.this"]
		require.True(t, ok, "data.databricks_notification_destinations.this has to be there")

		notificationCount := ds.Primary.Attributes["notification_destinations.#"]
		require.Equal(t, "2", notificationCount, "expected two notifications")

		notificationIds := []string{
			ds.Primary.Attributes["notification_destinations.0.id"],
			ds.Primary.Attributes["notification_destinations.1.id"],
		}

		expectedNotificationIds := []string{
			s.Modules[0].Resources["databricks_notification_destination.email_notification"].Primary.ID,
			s.Modules[0].Resources["databricks_notification_destination.slack_notification"].Primary.ID,
		}

		assert.ElementsMatch(t, expectedNotificationIds, notificationIds, "expected notification_destination ids to match")

		return nil
	}
}

func TestAccNotificationsCreation(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
			resource "databricks_notification_destination" "email_notification" {
				display_name = "email notification destination"
				config {
					email {
					addresses = ["abc@gmail.com"]
					}
				}
			}

			resource "databricks_notification_destination" "notification" {
				display_name = "slack notification destination"
				config {
					slack {
					url = "https://hooks.slack.com/services/..."
					}
				}
			}

			data "databricks_notification_destinations" "this" {}
		`,
		Check: CheckDataSourceNotificationsPopulated(t),
	})
}
