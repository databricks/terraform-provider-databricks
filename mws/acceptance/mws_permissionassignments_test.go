package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestUcAccAssignGroupToWorkspace(t *testing.T) {
	qa.RequireCloudEnv(t, "ucacct")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_group" "this" {
				display_name = "TF {var.RANDOM}"
			}
			resource "databricks_mws_permission_assignment" "this" {
				workspace_id = {env.TEST_WORKSPACE_ID}
				principal_id = databricks_group.this.id
				permissions  = ["USER"]
			}`,
		},
		{
			Template: `
			resource "databricks_group" "this" {
				display_name = "TF {var.RANDOM}"
			}
			resource "databricks_mws_permission_assignment" "this" {
				workspace_id = {env.TEST_WORKSPACE_ID}
				principal_id = databricks_group.this.id
				permissions  = ["ADMIN"]
			}`,
		},
		{
			Template: `
			resource "databricks_group" "this" {
				display_name = "TF {var.RANDOM}"
			}
			resource "databricks_mws_permission_assignment" "this" {
				workspace_id = {env.TEST_WORKSPACE_ID}
				principal_id = databricks_group.this.id
				permissions  = ["USER"]
			}`,
		},
	})
}

func TestAccAssignSpnToWorkspace(t *testing.T) {
	qa.RequireCloudEnv(t, "unity-catalog-account")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_service_principal" "this" {
				display_name = "TF {var.RANDOM}"
			}
			resource "databricks_mws_permission_assignment" "this" {
				workspace_id = {env.TEST_WORKSPACE_ID}
				principal_id = databricks_service_principal.this.id
				permissions  = ["USER"]
			}`,
		},
	})
}
