package notificationdestinations_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Notification destinations are created and read back by ID, which is
// read-after-write consistent, so creation can be asserted directly against the
// resource state.
func TestAccNotificationDestinationCreate(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
			resource "databricks_notification_destination" "email" {
				display_name = "tf-test-{var.STICKY_RANDOM}-email"
				config {
					email {
						addresses = ["abc@gmail.com"]
					}
				}
			}

			resource "databricks_notification_destination" "slack" {
				display_name = "tf-test-{var.STICKY_RANDOM}-slack"
				config {
					slack {
						url = "https://hooks.slack.com/services/..."
					}
				}
			}
		`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("databricks_notification_destination.email", "id"),
			resource.TestCheckResourceAttr("databricks_notification_destination.email", "destination_type", "EMAIL"),
			resource.TestCheckResourceAttrSet("databricks_notification_destination.slack", "id"),
			resource.TestCheckResourceAttr("databricks_notification_destination.slack", "destination_type", "SLACK"),
		),
	})
}

// The notification destinations list API is eventually consistent, so the data
// source is validated against a destination that already exists on the workspace
// (TEST_NOTIFICATION_DESTINATION_ID) rather than one created by the test: a
// freshly created destination may not yet appear in the list response.
func TestAccNotificationDestinationsDataSource(t *testing.T) {
	destinationID := acceptance.GetEnvOrSkipTest(t, "TEST_NOTIFICATION_DESTINATION_ID")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `data "databricks_notification_destinations" "this" {}`,
		Check: func(s *terraform.State) error {
			ds, ok := s.Modules[0].Resources["data.databricks_notification_destinations.this"]
			if !ok {
				return fmt.Errorf("data.databricks_notification_destinations.this is missing from the state")
			}
			count, err := strconv.Atoi(ds.Primary.Attributes["notification_destinations.#"])
			if err != nil {
				return fmt.Errorf("notification_destinations count is not a number: %w", err)
			}
			for i := 0; i < count; i++ {
				if ds.Primary.Attributes[fmt.Sprintf("notification_destinations.%d.id", i)] == destinationID {
					return nil
				}
			}
			return fmt.Errorf("notification destination %s not returned by the data source (%d listed)", destinationID, count)
		},
	})
}
