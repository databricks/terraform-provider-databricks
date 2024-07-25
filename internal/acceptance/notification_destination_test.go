package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestAccNDEmail(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_notification_destination" "this" {
			display_name = "Email Notification Destination"
			config {
				email {
					addresses = ["` + qa.RandomEmail() + `"]
				}
			}
		}
		`,
	})
}
