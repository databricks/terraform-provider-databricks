package acceptance

import (
	"fmt"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
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
							permissions, err := NewPermissionsAPI(client).Read(id)
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
						permissions, err := NewPermissionsAPI(client).Read(id)
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

func TestAccDatabricksJobPermissionsResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	client := new(common.DatabricksClient)
	nodeType := qa.GetCloudInstanceType(client)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "databricks_job" "this" {
					name = "First %[1]s"
					timeout_seconds = 3600
					max_retries = 1
					max_concurrent_runs = 1

					new_cluster  {
						num_workers   = 1
						spark_version = "6.6.x-scala2.11"
						node_type_id  = "%[2]s"
					}
				}

				resource "databricks_group" "first" {
					display_name = "First %[1]s"
				}

				resource "databricks_permissions" "dummy" {
					job_id = databricks_job.this.id
					access_control {
						group_name = databricks_group.first.display_name
						permission_level = "CAN_MANAGE"
					}
				}`, randomName, nodeType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_permissions.dummy",
						"object_type", "job"),
					acceptance.ResourceCheck("databricks_permissions.dummy",
						func(client *common.DatabricksClient, id string) error {
							job, err := compute.NewJobsAPI(client).Read(id)

							assert.NoError(t, err)

							userName := job.CreatorUserName

							userACL := AccessControlChange{
								UserName:        &userName,
								PermissionLevel: "IS_OWNER",
							}

							accessControlChange := []*AccessControlChange{&userACL}

							jobACL := AccessControlChangeList{
								AccessControlList: accessControlChange,
							}

							param := &jobACL
							permissionsAPIUpdateErr := NewPermissionsAPI(client).SetOrDelete(fmt.Sprintf("/jobs/%s/", id), param)

							assert.NoError(t, permissionsAPIUpdateErr)

							permissions, permissionsAPIReadErr := NewPermissionsAPI(client).Read(id)

							assert.NoError(t, permissionsAPIReadErr)

							fmt.Printf("%s", permissions)
							assert.Len(t, permissions.AccessControlList, 2)
							return nil
						}),
				),
			},
		},
	})
}
