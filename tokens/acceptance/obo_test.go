package acceptance

import (
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"

	"testing"
)

func TestAwsAccOboTokenResource(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
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
		},
	})
}
