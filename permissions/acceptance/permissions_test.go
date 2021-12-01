package acceptance

import (
	"context"
	"fmt"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/permissions"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccDatabricksPermissionsResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
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
					acceptance.ResourceCheck("databricks_permissions.dummy",
						func(ctx context.Context, client *common.DatabricksClient, id string) error {
							permissions, err := permissions.NewPermissionsAPI(ctx, client).Read(id)
							if err != nil {
								return err
							}
							assert.Len(t, permissions.AccessControlList, 2)
							return nil
						}),
				),
			},
			{
				Config: fmt.Sprintf(`
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
				Check: acceptance.ResourceCheck("databricks_permissions.dummy",
					func(ctx context.Context, client *common.DatabricksClient, id string) error {
						permissions, err := permissions.NewPermissionsAPI(ctx, client).Read(id)
						if err != nil {
							return err
						}
						assert.Len(t, permissions.AccessControlList, 3)
						return nil
					}),
			},
		},
	})
}

func TestAccDatabricksReposPermissionsResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "databricks_repo" "this" {
					url = "https://github.com/databrickslabs/tempo.git"
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
					acceptance.ResourceCheck("databricks_permissions.dummy",
						func(ctx context.Context, client *common.DatabricksClient, id string) error {
							permissions, err := permissions.NewPermissionsAPI(ctx, client).Read(id)
							if err != nil {
								return err
							}
							assert.Len(t, permissions.AccessControlList, 4)
							return nil
						}),
				),
			},
		},
	})
}
