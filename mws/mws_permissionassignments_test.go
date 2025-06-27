package mws_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestUcAccAssignGroupToWorkspace(t *testing.T) {
	acceptance.UnityAccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_group" "this" {
			display_name = "TF {var.RANDOM}"
		}
		resource "databricks_mws_permission_assignment" "this" {
			workspace_id = {env.DUMMY_WORKSPACE_ID}
			principal_id = databricks_group.this.id
			permissions  = ["USER"]
		}`,
	}, acceptance.Step{
		Template: `
		resource "databricks_group" "this" {
			display_name = "TF {var.RANDOM}"
		}
		resource "databricks_mws_permission_assignment" "this" {
			workspace_id = {env.DUMMY_WORKSPACE_ID}
			principal_id = databricks_group.this.id
			permissions  = ["ADMIN"]
		}`,
	}, acceptance.Step{
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
	acceptance.UnityAccountLevel(t, acceptance.Step{
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
