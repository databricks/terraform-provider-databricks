package acceptance

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/permissions"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccDatabricksPermissionsResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	workspaceLevel(t, step{
		Template: fmt.Sprintf(`
		resource "databricks_notebook" "this" {
			content_base64 = base64encode("# Databricks notebook source\nprint(1)")
			path = "/Beginning/%[1]s/Init"
			language = "PYTHON"
		}
		resource "databricks_group" "first" {
			display_name = "First %[1]s"
		}
		resource "databricks_permissions" "dummy" {
			notebook_path = databricks_notebook.this.id
			access_control {
				group_name = databricks_group.first.display_name
				permission_level = "CAN_MANAGE"
			}
		}`, randomName),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permissions.dummy",
				"object_type", "notebook"),
			resourceCheck("databricks_permissions.dummy",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
					permissions, err := permissions.NewPermissionsAPI(ctx, client).Read(id)
					if err != nil {
						return err
					}
					assert.GreaterOrEqual(t, len(permissions.AccessControlList), 1)
					return nil
				}),
		),
	}, step{
		Template: fmt.Sprintf(`
		resource "databricks_notebook" "this" {
			content_base64 = base64encode("# Databricks notebook source\nprint(1)")
			path = "/Beginning/%[1]s/Init"
			language = "PYTHON"
		}
		resource "databricks_group" "first" {
			display_name = "First %[1]s"
		}
		resource "databricks_group" "second" {
			display_name = "Second %[1]s"
		}
		resource "databricks_permissions" "dummy" {
			notebook_path = databricks_notebook.this.id
			access_control {
				group_name = databricks_group.first.display_name
				permission_level = "CAN_MANAGE"
			}
			access_control {
				group_name = databricks_group.second.display_name
				permission_level = "CAN_RUN"
			}
		}`, randomName),
		Check: resourceCheck("databricks_permissions.dummy",
			func(ctx context.Context, client *common.DatabricksClient, id string) error {
				permissions, err := permissions.NewPermissionsAPI(ctx, client).Read(id)
				if err != nil {
					return err
				}
				assert.GreaterOrEqual(t, len(permissions.AccessControlList), 2)
				return nil
			}),
	})
}

func TestAccDatabricksReposPermissionsResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	workspaceLevel(t, step{
		Template: fmt.Sprintf(`
		resource "databricks_repo" "this" {
			url = "https://github.com/databrickslabs/tempo.git"
			path = "/Repos/terraform-tests/tempo-%[1]s"
		}
		resource "databricks_group" "first" {
			display_name = "First %[1]s"
		}
		resource "databricks_group" "second" {
			display_name = "Second %[1]s"
		}
		resource "databricks_permissions" "dummy" {
			repo_path = databricks_repo.this.path
			access_control {
				group_name = databricks_group.first.display_name
				permission_level = "CAN_MANAGE"
			}
			access_control {
				group_name = databricks_group.second.display_name
				permission_level = "CAN_RUN"
			}
		}`, randomName),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permissions.dummy",
				"object_type", "repo"),
			resourceCheck("databricks_permissions.dummy",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
					permissions, err := permissions.NewPermissionsAPI(ctx, client).Read(id)
					if err != nil {
						return err
					}
					assert.GreaterOrEqual(t, len(permissions.AccessControlList), 2)
					return nil
				}),
		),
	})
}
