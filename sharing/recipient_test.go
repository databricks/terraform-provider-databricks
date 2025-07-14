package sharing_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestUcAccCreateRecipientDb2Open(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_recipient" "db2open" {
			name = "{var.RANDOM}-terraform-db2open-recipient"
			comment = "made by terraform"
			authentication_type = "TOKEN"
			sharing_code = "{var.RANDOM}"
			ip_access_list {
			// using private ip for acc testing
			allowed_ip_addresses = ["10.0.0.0/16"]
			}
		}`,
	})
}

func TestUcAccCreateRecipientDb2DbAws(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_metastore" "recipient_metastore" {
			name = "{var.RANDOM}-terraform-recipient-metastore"
			storage_root = "s3://{env.TEST_BUCKET}/test{var.RANDOM}"
			delta_sharing_scope = "INTERNAL"
			delta_sharing_recipient_token_lifetime_in_seconds = "60000"
			force_destroy = true
			lifecycle {
			// fake storage root is causing issues
			ignore_changes = [storage_root]
			}
		}

		resource "databricks_recipient" "db2db" {
			name = "{var.RANDOM}-terraform-db2db-recipient"
			comment = "made by terraform"
			authentication_type = "DATABRICKS"
			data_recipient_global_metastore_id = databricks_metastore.recipient_metastore.global_metastore_id
		}`,
	})
}

func TestUcAccUpdateRecipientDb2Open(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: recipientTemplateWithOwner("made by terraform", "account users"),
	}, acceptance.Step{
		Template: recipientTemplateWithOwner("made by terraform -- updated comment", "account users"),
	}, acceptance.Step{
		Template: recipientTemplateWithOwner("made by terraform -- updated comment", "{env.TEST_DATA_ENG_GROUP}"),
	}, acceptance.Step{
		Template: recipientTemplateWithOwner("made by terraform -- updated comment 2", "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"),
	})
}

func recipientTemplateWithOwner(comment string, owner string) string {
	return fmt.Sprintf(`
		resource "databricks_recipient" "db2open" {
			name = "{var.STICKY_RANDOM}-terraform-db2open-recipient"
			comment = "%s"
			owner = "%s"
			authentication_type = "TOKEN"
			sharing_code = "{var.STICKY_RANDOM}"
			ip_access_list {
			// using private ip for acc testing
			allowed_ip_addresses = ["10.0.0.0/16"]
			}
		}
	`, comment, owner)
}
