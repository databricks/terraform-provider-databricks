package acceptance

import (
	"testing"
)

func TestUcAccAssignGroupToWorkspace(t *testing.T) {
	unityAccountLevel(t, LegacyStep{
		Template: `
		resource "databricks_group" "this" {
			display_name = "TF {var.RANDOM}"
		}
		resource "databricks_mws_permission_assignment" "this" {
			workspace_id = {env.DUMMY_WORKSPACE_ID}
			principal_id = databricks_group.this.id
			permissions  = ["USER"]
		}`,
	}, LegacyStep{
		Template: `
		resource "databricks_group" "this" {
			display_name = "TF {var.RANDOM}"
		}
		resource "databricks_mws_permission_assignment" "this" {
			workspace_id = {env.DUMMY_WORKSPACE_ID}
			principal_id = databricks_group.this.id
			permissions  = ["ADMIN"]
		}`,
	}, LegacyStep{
		Template: `
		resource "databricks_group" "this" {
			display_name = "TF {var.RANDOM}"
		}
		resource "databricks_mws_permission_assignment" "this" {
			workspace_id = {env.DUMMY_WORKSPACE_ID}
			principal_id = databricks_group.this.id
			permissions  = ["USER"]
		}`,
	})
}

func TestAccAssignSpnToWorkspace(t *testing.T) {
	unityAccountLevel(t, LegacyStep{
		Template: `
		resource "databricks_service_principal" "this" {
			display_name = "TF {var.RANDOM}"
		}
		resource "databricks_mws_permission_assignment" "this" {
			workspace_id = {env.DUMMY_WORKSPACE_ID}
			principal_id = databricks_service_principal.this.id
			permissions  = ["USER"]
		}`,
	})
}
