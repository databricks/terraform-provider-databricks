package notificationdestinations_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

func CheckDataSourceNotificationsPopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.Modules[0].Resources["data.databricks_notification_destinations.this"]
		require.True(t, ok, "data.databricks_notification_destinations.this has to be there")

		emailNotif := s.Modules[0].Resources["databricks_notification_destination.email_notification"]
		slackNotif := s.Modules[0].Resources["databricks_notification_destination.slack_notification"]

		emailId := emailNotif.Primary.ID
		slackId := slackNotif.Primary.ID

		notificationsCount := ds.Primary.Attributes["notification_destinations.#"]
		notificationsCountInt, err := strconv.Atoi(notificationsCount)
		require.NoError(t, err, "notification destinations count is not a number")
		require.NotEqual(t, 0, notificationsCountInt, "notification destinations list is empty")

		foundEmailId := false
		foundSlackId := false
		for i := 0; i < notificationsCountInt; i++ {
			id := ds.Primary.Attributes[fmt.Sprintf("notification_destinations.%d.id", i)]
			if id == emailId {
				foundEmailId = true
			}
			if id == slackId {
				foundSlackId = true
			}
		}
		require.True(t, foundEmailId, "email notification id not found in destinations")
		require.True(t, foundSlackId, "email notification id not found in destinations")
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

			resource "databricks_notification_destination" "slack_notification" {
				display_name = "slack notification destination"
				config {
					slack {
					url = "https://hooks.slack.com/services/..."
					}
				}
			}

			data "databricks_notification_destinations" "this" {
				depends_on = [databricks_notification_destination.email_notification, databricks_notification_destination.slack_notification]
			}
		`,
		Check: CheckDataSourceNotificationsPopulated(t),
	})
}
