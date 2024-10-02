package acceptance

import (
	"testing"
)

func TestUcAccAssignGroupToWorkspace(t *testing.T) {
	UnityAccountLevel(t, Step{
		Template: `
		resource "databricks_group" "this" {
			display_name = "TF {var.RANDOM}"
		}
		resource "databricks_mws_permission_assignment" "this" {
			workspace_id = {env.DUMMY_WORKSPACE_ID}
			principal_id = databricks_group.this.id
			permissions  = ["USER"]
		}`,
	}, Step{
		Template: `
		resource "databricks_group" "this" {
			display_name = "TF {var.RANDOM}"
		}
		resource "databricks_mws_permission_assignment" "this" {
			workspace_id = {env.DUMMY_WORKSPACE_ID}
			principal_id = databricks_group.this.id
			permissions  = ["ADMIN"]
		}`,
	}, Step{
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
	UnityAccountLevel(t, Step{
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
