package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestAccCreateRecipientDb2Open(t *testing.T) {
	qa.RequireCloudEnv(t, "aws-uc-prod")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_recipient" "db2open" {
			  name = "{var.RANDOM}-terraform-db2open-recipient"
			  comment = "made by terraform"
			  authentication_type = "TOKEN"
			  sharing_code = "{var.RANDOM}"
			  ip_access_list {
				allowed_ip_addresses = ["10.0.0.0/16"] // using private ip for acc testing
			  }
			}`,
		},
	})
}

func TestAccCreateRecipientDb2DbAws(t *testing.T) {
	qa.RequireCloudEnv(t, "aws-uc-prod")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_metastore" "recipient_metastore" {
			  name = "{var.RANDOM}-terraform-recipient-metastore"
			  storage_root = format("s3a://%s/%s", "{var.RANDOM}", "{var.RANDOM}")
			  delta_sharing_scope = "INTERNAL"
			  delta_sharing_recipient_token_lifetime_in_seconds = "60000"
			  force_destroy = true
			  lifecycle { ignore_changes = [storage_root] } // fake storage root is causing issues
			}
			
			resource "databricks_recipient" "db2db" {
			  name = "{var.RANDOM}-terraform-db2db-recipient"
			  comment = "made by terraform"
			  authentication_type = "DATABRICKS"
			  data_recipient_global_metastore_id = databricks_metastore.recipient_metastore.global_metastore_id
			}`,
		},
	})
}
