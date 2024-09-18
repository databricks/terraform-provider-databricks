package acceptance

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/permissions"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccDatabricksPermissionsResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	WorkspaceLevel(t, Step{
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
	}, Step{
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
	WorkspaceLevel(t, Step{
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

func TestAccDatabricksPermissionsForSqlWarehouses(t *testing.T) {
	// Random string to annotate newly created groups
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	// Create a client to query the permissions API
	c, err := client.New(&config.Config{})
	require.NoError(t, err)
	permissionsClient := permissions.NewPermissionsAPI(context.Background(), &common.DatabricksClient{DatabricksClient: c})

	// Validates export attribute "object_type" for the permissions resource
	// is set to warehouses
	checkObjectType := resource.TestCheckResourceAttr("databricks_permissions.this",
		"object_type", "warehouses")

	// Asserts value of a permission level for a group
	assertPermissionLevel := func(t *testing.T, permissionId, groupName, permissionLevel string) {
		// Query permissions on warehouse
		warehousePermissions, err := permissionsClient.Read(permissionId)
		require.NoError(t, err)

		// Assert expected permission level is present
		assert.Contains(t, warehousePermissions.AccessControlList, permissions.AccessControlApiResponse{
			GroupName: groupName,
			AllPermissions: []permissions.PermissionApiResponse{
				{
					PermissionLevel: permissionLevel,
				},
			},
		})
	}

	// Get permission ID from the terraform state
	getPermissionId := func(s *terraform.State) string {
		resourcePermission, ok := s.RootModule().Resources["databricks_permissions.this"]
		require.True(t, ok, "could not find permissions resource: databricks_permissions.this")
		return resourcePermission.Primary.ID
	}

	// Configuration for step 1 of the test. Create a databricks_permissions
	// resources, assigning a group CAN_MANAGE  permission to the warehouse.
	config1 := fmt.Sprintf(`
	resource "databricks_group" "one" {
		display_name = "test-warehouse-permission-one-%s"
	}
	resource "databricks_permissions" "this" {
		sql_endpoint_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
		access_control {
			group_name = databricks_group.one.display_name
			permission_level = "CAN_MANAGE"
		}
	}`, randomName)

	// Configuration for step 2 of the test. Create another group and update
	// permissions to CAN_USE for the second group
	config2 := fmt.Sprintf(`			
	resource "databricks_group" "one" {
		display_name = "test-warehouse-permission-one-%[1]s"
	}
	resource "databricks_group" "two" {
		display_name = "test-warehouse-permission-two-%[1]s"
	}
	resource "databricks_permissions" "this" {
		sql_endpoint_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
		access_control {
			group_name = databricks_group.one.display_name
			permission_level = "CAN_MANAGE"
		}
		access_control {
			group_name = databricks_group.two.display_name
			permission_level = "CAN_USE"
		}
	}`, randomName)

	WorkspaceLevel(t,
		Step{
			Template: config1,
			Check: resource.ComposeTestCheckFunc(
				checkObjectType,
				func(s *terraform.State) error {
					id := getPermissionId(s)
					assertPermissionLevel(t, id, "test-warehouse-permission-one-"+randomName, "CAN_MANAGE")
					return nil
				},
			),
		},
		Step{
			Template: config2,
			Check: func(s *terraform.State) error {
				id := getPermissionId(s)
				assertPermissionLevel(t, id, "test-warehouse-permission-one-"+randomName, "CAN_MANAGE")
				assertPermissionLevel(t, id, "test-warehouse-permission-two-"+randomName, "CAN_USE")
				return nil
			},
		},
	)
}

func TestAccDatabricksPermissionsForJobs(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		data databricks_current_user me {}

		resource "databricks_job" "this" {
			name = "{var.RANDOM}"
		}

		resource "databricks_permissions" "this" {
			job_id = databricks_job.this.id
			access_control {
			    permission_level = "IS_OWNER"
				service_principal_name = data.databricks_current_user.me.user_name
			}
		}
		`,
	}, Step{
		Template: `
		data databricks_current_user me {}

		resource "databricks_job" "this" {
			name = "{var.RANDOM}"
		}

		resource "databricks_service_principal" "this" {
			display_name = "{var.RANDOM}"
		}

		resource "databricks_permissions" "this" {
			job_id = databricks_job.this.id
			# Lower the current users permissions to CAN_MANAGE and set a new owner
			access_control {
			    permission_level = "CAN_MANAGE"
				service_principal_name = data.databricks_current_user.me.user_name
			}
			access_control {
			    permission_level = "IS_OWNER"
				service_principal_name = databricks_service_principal.this.application_id
			}
		}
		`,
	}, Step{
		Template: `
		resource "databricks_job" "this" {
			name = "{var.RANDOM}"
		}
		`,
		// The current user should be the owner after permissions are removed.
		Check: func(s *terraform.State) error {
			job, ok := s.RootModule().Resources["databricks_job.this"]
			require.True(t, ok, "could not find job resource: databricks_job.this")
			w := databricks.Must(databricks.NewWorkspaceClient())
			permissions, err := getCurrentUserPermissions(context.Background(), t, w, job.Primary.ID)
			assert.NoError(t, err)
			assert.Len(t, permissions, 1)
			assert.Equal(t, jobs.JobPermissionLevelIsOwner, permissions[0].PermissionLevel)
			return nil
		},
	})
}

// getCurrentUserPermissions gets the permissions for the current user on a job with a given ID. If the user
// does not have any permissions on the job, an error is returned. This does not check whether the user belongs
// to any groups that have permissions on the job; it only checks the user's direct permissions.
func getCurrentUserPermissions(ctx context.Context, t *testing.T, w *databricks.WorkspaceClient, jobId string) ([]jobs.JobPermission, error) {
	permissions, err := w.Jobs.GetPermissions(ctx, jobs.GetJobPermissionsRequest{
		JobId: jobId,
	})
	require.NoError(t, err)
	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)
	for _, acl := range permissions.AccessControlList {
		if acl.ServicePrincipalName != me.UserName && acl.UserName != me.UserName {
			continue
		}
		return acl.AllPermissions, nil
	}
	return nil, fmt.Errorf("could not find current user %s in permissions for job %s", me.UserName, jobId)
}
