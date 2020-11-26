package acceptance

import (
	"context"
	"fmt"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
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
					content = base64encode("# Databricks notebook source\nprint(1)")
					path = "/Beginning/%[1]s/Init"
					overwrite = true
					mkdirs = true
					language = "PYTHON"
					format = "SOURCE"
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
						func(client *common.DatabricksClient, id string) error {
							ctx := context.Background()
							permissions, err := NewPermissionsAPI(ctx, client).Read(id)
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
					content = base64encode("# Databricks notebook source\nprint(1)")
					path = "/Beginning/%[1]s/Init"
					overwrite = true
					mkdirs = true
					language = "PYTHON"
					format = "SOURCE"
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
					func(client *common.DatabricksClient, id string) error {
						ctx := context.Background()
						permissions, err := NewPermissionsAPI(ctx, client).Read(id)
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
