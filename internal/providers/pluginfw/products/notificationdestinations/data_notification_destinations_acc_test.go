package notificationdestinations_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
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
		require.Equal(t, 2, notificationsCountInt, "expected two notification destinations")

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
		require.True(t, foundSlackId, "slack notification id not found in destinations")
		return nil
	}
}

func waitForNotificationsInList(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		emailNotif := s.Modules[0].Resources["databricks_notification_destination.email_notification"]
		slackNotif := s.Modules[0].Resources["databricks_notification_destination.slack_notification"]

		return acceptance.ResourceCheck("databricks_notification_destination.email_notification", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, "")
			if err != nil {
				return err
			}
			expectedIDs := map[string]bool{
				emailNotif.Primary.ID: false,
				slackNotif.Primary.ID: false,
			}
			return resource.RetryContext(ctx, 2*time.Minute, func() *resource.RetryError {
				destinations, err := w.NotificationDestinations.ListAll(ctx, settings.ListNotificationDestinationsRequest{})
				if err != nil {
					return resource.NonRetryableError(err)
				}
				for _, destination := range destinations {
					if _, ok := expectedIDs[destination.Id]; ok {
						expectedIDs[destination.Id] = true
					}
				}
				for expectedID, found := range expectedIDs {
					if !found {
						return resource.RetryableError(fmt.Errorf("notification destination %s not found in list response", expectedID))
					}
				}
				return nil
			})
		})(s)
	}
}

func TestAccNotificationsCreation(t *testing.T) {
	displayNamePrefix := "Notification Destination - " + qa.RandomName()
	slackWebhookURL := "https://hooks.slack.com/services/" + qa.RandomName()
	notificationDestinationsTemplate := `
		resource "databricks_notification_destination" "email_notification" {
			display_name = "` + displayNamePrefix + ` email"
			config {
				email {
					addresses = ["` + qa.RandomEmail() + `"]
				}
			}
		}

		resource "databricks_notification_destination" "slack_notification" {
			display_name = "` + displayNamePrefix + ` slack"
			config {
				slack {
					url = "` + slackWebhookURL + `"
				}
			}
		}
	`
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notificationDestinationsTemplate,
		Check:    waitForNotificationsInList(t),
	}, acceptance.Step{
		Template: `
			` + notificationDestinationsTemplate + `
			data "databricks_notification_destinations" "this" {
				display_name_contains = "` + displayNamePrefix + `"
				depends_on = [databricks_notification_destination.email_notification, databricks_notification_destination.slack_notification]
			}
		`,
		Check: CheckDataSourceNotificationsPopulated(t),
	})
}
