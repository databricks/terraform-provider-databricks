package acceptance

import (
	"testing"
)

func TestAccAwsOboTokenResource(t *testing.T) {
	// running this test temporarily on the UC WS level
	// until infrastructure gets AWS specific markers
	unityWorkspaceLevel(t, LegacyStep{
		Template: `
		// dummy: {env.TEST_GLOBAL_METASTORE_ID}
		resource "databricks_service_principal" "this" {
			display_name = "tf-{var.RANDOM}"
		}

		data "databricks_group" "admins" {
			display_name = "admins"
		}
		
		resource "databricks_group_member" "this" {
			group_id = data.databricks_group.admins.id
			member_id = databricks_service_principal.this.id
		}

		resource "databricks_obo_token" "this" {
			depends_on = [databricks_group_member.this]
			application_id = databricks_service_principal.this.application_id
			comment = "PAT on behalf of ${databricks_service_principal.this.display_name}"
			lifetime_seconds = 3600
		}`,
	})
}
