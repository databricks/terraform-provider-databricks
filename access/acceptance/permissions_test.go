package acceptance

import (
	"fmt"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccDatabricksPermissionsResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	acceptance.AccTest(t, resource.TestCase{

		Steps: []resource.TestStep{
			{
				Config: testClusterPolicyPermissions(randomName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_permissions.dummy_can_use",
						"object_type", "cluster-policy"),
					acceptance.ResourceCheck("databricks_permissions.dummy_can_use",
						func(client *common.DatabricksClient, id string) error {
							permissions, err := NewPermissionsAPI(client).Read(id)
							if err != nil {
								return err
							}
							assert.Len(t, permissions.AccessControlList, 3)
							return nil
						}),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testClusterPolicyPermissionsSecondGroupAdded(randomName),
				Check: acceptance.ResourceCheck("databricks_permissions.dummy_can_use",
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

func testClusterPolicyPermissions(name string) string {
	return fmt.Sprintf(`
	resource "databricks_cluster_policy" "something_simple" {
		name = "Terraform Policy %[1]s"
		definition = jsonencode({
			"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
				"type": "forbidden"
			}
		})
	}
	resource "databricks_scim_group" "dummy_group" {
		display_name = "Terraform Group %[1]s"
	}
	resource "databricks_permissions" "dummy_can_use" {
		cluster_policy_id = databricks_cluster_policy.something_simple.id
		access_control {
			group_name = databricks_scim_group.dummy_group.display_name
			permission_level = "CAN_USE"
		}
	}
	`, name)
}

func testClusterPolicyPermissionsSecondGroupAdded(name string) string {
	return fmt.Sprintf(`
	resource "databricks_cluster_policy" "something_simple" {
		name = "Terraform Policy %[1]s"
		definition = jsonencode({
			"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
				"type": "forbidden"
			},
			"spark_conf.spark.secondkey": {
				"type": "forbidden"
			}
		})
	}
	resource "databricks_scim_group" "dummy_group" {
		display_name = "Terraform Group %[1]s"
	}
	resource "databricks_scim_group" "second_group" {
		display_name = "Terraform Second Group %[1]s"
	}
	resource "databricks_permissions" "dummy_can_use" {
		cluster_policy_id = databricks_cluster_policy.something_simple.id
		access_control {
			group_name = databricks_scim_group.dummy_group.display_name
			permission_level = "CAN_USE"
		}
		access_control {
			group_name = databricks_scim_group.second_group.display_name
			permission_level = "CAN_USE"
		}
	}
	`, name)
}

func TestAccNotebookPermissions(t *testing.T) {
	// TODO: fails with big run...
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
				resource "databricks_scim_group" "group" {
					display_name = "Terraform Group %[1]s"
				}
				resource "databricks_permissions" "manage" {
					directory_path = "/Beginning/%[1]s"
					access_control {
						group_name = databricks_scim_group.group.display_name
						permission_level = "CAN_MANAGE"
					}
				}
				`, randomName),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_permissions.manage",
						"object_type", "directory"),
					resource.TestCheckResourceAttr("databricks_permissions.manage",
						"access_control.0.group_name", "Terraform Group "+randomName),
					acceptance.ResourceCheck("databricks_permissions.manage",
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
		},
	})
}
