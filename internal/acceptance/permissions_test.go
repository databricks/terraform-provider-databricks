package acceptance

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/permissions"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getId(id string) string {
	idParts := strings.Split(id, "/")
	return idParts[len(idParts)-1]
}

func TestAccPermissions_Notebook(t *testing.T) {
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
					w := databricks.Must(databricks.NewWorkspaceClient())
					permissions, err := w.Workspace.GetPermissions(ctx, workspace.GetWorkspaceObjectPermissionsRequest{
						WorkspaceObjectId:   getId(id),
						WorkspaceObjectType: "notebooks",
					})
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
				w := databricks.Must(databricks.NewWorkspaceClient())
				permissions, err := w.Workspace.GetPermissions(ctx, workspace.GetWorkspaceObjectPermissionsRequest{
					WorkspaceObjectId:   getId(id),
					WorkspaceObjectType: "notebook",
				})
				if err != nil {
					return err
				}
				assert.GreaterOrEqual(t, len(permissions.AccessControlList), 2)
				return nil
			}),
	})
}

func TestAccPermissions_Repo(t *testing.T) {
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
					w := databricks.Must(databricks.NewWorkspaceClient())
					permissions, err := w.Repos.GetPermissions(ctx, workspace.GetRepoPermissionsRequest{
						RepoId: getId(id),
					})
					if err != nil {
						return err
					}
					assert.GreaterOrEqual(t, len(permissions.AccessControlList), 2)
					return nil
				}),
		),
	})
}

func TestAccPermissions_SqlWarehouse(t *testing.T) {
	// Random string to annotate newly created groups
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	// Validates export attribute "object_type" for the permissions resource
	// is set to warehouses
	checkObjectType := resource.TestCheckResourceAttr("databricks_permissions.this",
		"object_type", "warehouses")

	// Asserts value of a permission level for a group
	assertPermissionLevel := func(t *testing.T, permissionId, groupName, permissionLevel string) {
		// Query permissions on warehouse
		w := databricks.Must(databricks.NewWorkspaceClient())
		warehousePermissions, err := w.Warehouses.GetPermissionsByWarehouseId(context.Background(), permissionId)
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

func TestAccPermissions_Jobs(t *testing.T) {
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

const noPermissions = ""

type permissionSettings struct {
	// Name of the SP or group
	name string
	// If true, the resource will not be created
	skipCreation    bool
	permissionLevel string
}

func simpleSettings(permissionLevel ...string) []permissionSettings {
	var settings []permissionSettings
	for _, level := range permissionLevel {
		settings = append(settings, permissionSettings{permissionLevel: level})
	}
	return settings
}

func makePermissionsTestStage(idAttribute, idValue string, currentUserPermission string, servicePrincipalPermissions []permissionSettings, groupPermissions []permissionSettings) string {
	var resources string
	var accessControlBlocks string
	for i, servicePrincipal := range servicePrincipalPermissions {
		if !servicePrincipal.skipCreation {
			resources += fmt.Sprintf(`
			resource "databricks_service_principal" "_%d" {
				display_name = "{var.STICKY_RANDOM}-%d"
			}`, i, i)
		}
		var spName string
		if servicePrincipal.name == "" {
			spName = fmt.Sprintf("databricks_service_principal._%d.application_id", i)
		} else {
			spName = fmt.Sprintf(`"%s"`, servicePrincipal.name)
		}
		accessControlBlocks += fmt.Sprintf(`
		access_control {
			permission_level = "%s"
			service_principal_name = %s
		}`, servicePrincipal.permissionLevel, spName)
	}
	for i, group := range groupPermissions {
		if !group.skipCreation {
			resources += fmt.Sprintf(`
			resource "databricks_group" "_%d" {
				display_name = "{var.STICKY_RANDOM}-%d"
			}`, i, i)
		}
		var groupName string
		if group.name == "" {
			groupName = fmt.Sprintf("databricks_group._%d.display_name", i)
		} else {
			groupName = fmt.Sprintf(`"%s"`, group.name)
		}
		accessControlBlocks += fmt.Sprintf(`
		access_control {
			group_name = %s
			permission_level = "%s"
		}`, groupName, group.permissionLevel)
	}
	if currentUserPermission != "" {
		accessControlBlocks += fmt.Sprintf(`
		access_control {
			permission_level = "%s"
			service_principal_name = data.databricks_current_user.me.user_name
		}`, currentUserPermission)
	}
	return fmt.Sprintf(`
	data databricks_current_user me {}
	%s
	resource "databricks_permissions" "this" {
		%s = %s
		%s
	}
	`, resources, idAttribute, idValue, accessControlBlocks)
}

func TestAccPermissions_ClusterPolicy(t *testing.T) {
	policyTemplate := `
resource "databricks_cluster_policy" "this" {
	name = "{var.STICKY_RANDOM}"
	definition = jsonencode({
		"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
			"type": "fixed",
			"value": "jdbc:sqlserver://<jdbc-url>"
		}
	})
}`
	WorkspaceLevel(t, Step{
		Template: policyTemplate + makePermissionsTestStage("cluster_policy_id", "databricks_cluster_policy.this.id", noPermissions, nil, simpleSettings("CAN_USE")),
	}, Step{
		Template: policyTemplate + makePermissionsTestStage("cluster_policy_id", "databricks_cluster_policy.this.id", "CAN_USE", nil, simpleSettings("CAN_USE", "CAN_USE")),
	})
}

func TestAccPermissions_InstancePool(t *testing.T) {
	policyTemplate := `
data "databricks_node_type" "smallest" {
	local_disk = true
}
resource "databricks_instance_pool" "this" {
	instance_pool_name = "{var.STICKY_RANDOM}"
	min_idle_instances = 0
	max_capacity = 1
	node_type_id = data.databricks_node_type.smallest.id
	idle_instance_autotermination_minutes = 10
}`
	WorkspaceLevel(t, Step{
		Template: policyTemplate + makePermissionsTestStage("instance_pool_id", "databricks_instance_pool.this.id", noPermissions, nil, simpleSettings("CAN_ATTACH_TO")),
	}, Step{
		Template: policyTemplate + makePermissionsTestStage("instance_pool_id", "databricks_instance_pool.this.id", "CAN_MANAGE", nil, simpleSettings("CAN_ATTACH_TO", "CAN_MANAGE")),
	}, Step{
		Template:    policyTemplate + makePermissionsTestStage("instance_pool_id", "databricks_instance_pool.this.id", "CAN_ATTACH_TO", nil, nil),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for instance-pool, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Cluster(t *testing.T) {
	policyTemplate := `
data "databricks_spark_version" "latest" {
}
resource "databricks_cluster" "this" {
	cluster_name = "singlenode-{var.RANDOM}"
	spark_version = data.databricks_spark_version.latest.id
	instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
	num_workers = 0
	autotermination_minutes = 10
	spark_conf = {
		"spark.databricks.cluster.profile" = "singleNode"
		"spark.master" = "local[*]"
	}
	custom_tags = {
		"ResourceClass" = "SingleNode"
	}
}`
	WorkspaceLevel(t, Step{
		Template: policyTemplate + makePermissionsTestStage("cluster_id", "databricks_cluster.this.id", noPermissions, nil, simpleSettings("CAN_ATTACH_TO")),
	}, Step{
		Template: policyTemplate + makePermissionsTestStage("cluster_id", "databricks_cluster.this.id", "CAN_MANAGE", nil, simpleSettings("CAN_ATTACH_TO", "CAN_RESTART", "CAN_MANAGE")),
	}, Step{
		Template:    policyTemplate + makePermissionsTestStage("cluster_id", "databricks_cluster.this.id", "CAN_ATTACH_TO", nil, simpleSettings("CAN_ATTACH_TO", "CAN_RESTART", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for cluster, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Pipeline(t *testing.T) {
	policyTemplate := `
locals {
	name = "{var.STICKY_RANDOM}"
}
resource "databricks_pipeline" "this" {
	name = "${local.name}"
	storage = "/test/${local.name}"

	library {
		notebook {
			path = databricks_notebook.this.path
		}
	}
	continuous = false
}` + dltNotebookResource
	WorkspaceLevel(t, Step{
		Template: policyTemplate + makePermissionsTestStage("pipeline_id", "databricks_pipeline.this.id", noPermissions, nil, simpleSettings("CAN_VIEW")),
	}, Step{
		Template: policyTemplate + makePermissionsTestStage("pipeline_id", "databricks_pipeline.this.id", "IS_OWNER", nil, simpleSettings("CAN_VIEW", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template: policyTemplate + makePermissionsTestStage("pipeline_id", "databricks_pipeline.this.id", "CAN_MANAGE", simpleSettings("IS_OWNER"), simpleSettings("CAN_VIEW", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template:    policyTemplate + makePermissionsTestStage("pipeline_id", "databricks_pipeline.this.id", "CAN_RUN", nil, simpleSettings("CAN_VIEW", "CAN_RUN", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for pipelines, allowed levels: CAN_MANAGE, IS_OWNER"),
	}, Step{
		Template: policyTemplate,
		Check: resourceCheck("databricks_pipeline.this", func(ctx context.Context, c *common.DatabricksClient, id string) error {
			w, err := c.WorkspaceClient()
			assert.NoError(t, err)
			pipeline, err := w.Pipelines.GetByPipelineId(context.Background(), id)
			assert.NoError(t, err)
			permissions, err := w.Pipelines.GetPermissionsByPipelineId(context.Background(), id)
			assert.NoError(t, err)
			found := false
			for _, acl := range permissions.AccessControlList {
				if acl.ServicePrincipalName == pipeline.CreatorUserName {
					assert.Equal(t, pipelines.PipelinePermissionLevel("IS_OWNER"), acl.AllPermissions[0].PermissionLevel)
					found = true
					break
				}
			}
			if !found {
				assert.Fail(t, "pipeline creator not found in permissions")
			}
			return nil
		}),
	})
}

func TestAccPermissions_Notebook_Path(t *testing.T) {
	notebookTemplate := `
	resource "databricks_directory" "this" {
		path = "/permissions_test/{var.STICKY_RANDOM}"
	}
	resource "databricks_notebook" "this" {
		source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
		path = "${databricks_directory.this.path}/test_notebook"
	}`
	WorkspaceLevel(t, Step{
		Template: notebookTemplate + makePermissionsTestStage("notebook_path", "databricks_notebook.this.id", noPermissions, nil, simpleSettings("CAN_RUN")),
	}, Step{
		Template: notebookTemplate + makePermissionsTestStage("notebook_path", "databricks_notebook.this.id", "CAN_MANAGE", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: notebookTemplate + makePermissionsTestStage("notebook_path", "databricks_notebook.this.id", noPermissions, simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    notebookTemplate + makePermissionsTestStage("notebook_path", "databricks_notebook.this.id", "CAN_READ", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for notebook, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Notebook_Id(t *testing.T) {
	notebookTemplate := `
	resource "databricks_directory" "this" {
		path = "/permissions_test/{var.STICKY_RANDOM}"
	}
	resource "databricks_notebook" "this" {
		source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
		path = "${databricks_directory.this.path}/test_notebook"
	}`
	WorkspaceLevel(t, Step{
		Template: notebookTemplate + makePermissionsTestStage("notebook_id", "databricks_notebook.this.object_id", noPermissions, nil, simpleSettings("CAN_RUN")),
	}, Step{
		Template: notebookTemplate + makePermissionsTestStage("notebook_id", "databricks_notebook.this.object_id", "CAN_MANAGE", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: notebookTemplate + makePermissionsTestStage("notebook_id", "databricks_notebook.this.object_id", noPermissions, simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    notebookTemplate + makePermissionsTestStage("notebook_id", "databricks_notebook.this.object_id", "CAN_READ", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for notebook, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Directory_Path(t *testing.T) {
	directoryTemplate := `
	resource "databricks_directory" "this" {
		path = "/permissions_test/{var.STICKY_RANDOM}"
	}`
	WorkspaceLevel(t, Step{
		Template: directoryTemplate + makePermissionsTestStage("directory_path", "databricks_directory.this.id", noPermissions, nil, simpleSettings("CAN_RUN")),
	}, Step{
		Template: directoryTemplate + makePermissionsTestStage("directory_path", "databricks_directory.this.id", "CAN_MANAGE", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: directoryTemplate + makePermissionsTestStage("directory_path", "databricks_directory.this.id", noPermissions, simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    directoryTemplate + makePermissionsTestStage("directory_path", "databricks_directory.this.id", "CAN_READ", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for notebook, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Directory_Id(t *testing.T) {
	directoryTemplate := `
	resource "databricks_directory" "this" {
		path = "/permissions_test/{var.STICKY_RANDOM}"
	}`
	WorkspaceLevel(t, Step{
		Template: directoryTemplate + makePermissionsTestStage("directory_id", "databricks_directory.this.object_id", noPermissions, nil, simpleSettings("CAN_RUN")),
	}, Step{
		Template: directoryTemplate + makePermissionsTestStage("directory_id", "databricks_directory.this.object_id", "CAN_MANAGE", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: directoryTemplate + makePermissionsTestStage("directory_id", "databricks_directory.this.object_id", noPermissions, simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    directoryTemplate + makePermissionsTestStage("directory_id", "databricks_directory.this.object_id", "CAN_READ", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for directory, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_WorkspaceFile_Path(t *testing.T) {
	workspaceFile := `
	resource "databricks_directory" "this" {
		path = "/permissions_test/{var.STICKY_RANDOM}"
	}
	resource "databricks_workspace_file" "this" {
		source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
		path = "${databricks_directory.this.path}/test_notebook"
	}`
	WorkspaceLevel(t, Step{
		Template: workspaceFile + makePermissionsTestStage("workspace_file_path", "databricks_workspace_file.this.id", noPermissions, nil, simpleSettings("CAN_RUN")),
	}, Step{
		Template: workspaceFile + makePermissionsTestStage("workspace_file_path", "databricks_workspace_file.this.id", "CAN_MANAGE", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: workspaceFile + makePermissionsTestStage("workspace_file_path", "databricks_workspace_file.this.id", noPermissions, simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    workspaceFile + makePermissionsTestStage("workspace_file_path", "databricks_workspace_file.this.id", "CAN_READ", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for file, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_WorkspaceFile_Id(t *testing.T) {
	workspaceFile := `
	resource "databricks_directory" "this" {
		path = "/permissions_test/{var.STICKY_RANDOM}"
	}
	resource "databricks_workspace_file" "this" {
		source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
		path = "${databricks_directory.this.path}/test_notebook"
	}`
	WorkspaceLevel(t, Step{
		Template: workspaceFile + makePermissionsTestStage("workspace_file_id", "databricks_workspace_file.this.object_id", noPermissions, nil, simpleSettings("CAN_RUN")),
	}, Step{
		Template: workspaceFile + makePermissionsTestStage("workspace_file_id", "databricks_workspace_file.this.object_id", "CAN_MANAGE", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: workspaceFile + makePermissionsTestStage("workspace_file_id", "databricks_workspace_file.this.object_id", noPermissions, simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    workspaceFile + makePermissionsTestStage("workspace_file_id", "databricks_workspace_file.this.object_id", "CAN_READ", simpleSettings("CAN_MANAGE"), simpleSettings("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for file, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Authorization_Passwords(t *testing.T) {
	skipf(t)("ACLs for passwords are disabled on testing infrastructure")
	WorkspaceLevel(t, Step{
		Template: makePermissionsTestStage("authorization", "\"passwords\"", noPermissions, nil, simpleSettings("CAN_USE")),
	}, Step{
		Template: makePermissionsTestStage("authorization", "\"passwords\"", noPermissions, nil, []permissionSettings{{name: "admins", skipCreation: true, permissionLevel: "CAN_USE"}}),
	})
}

func TestAccPermissions_Authorization_Tokens(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: makePermissionsTestStage("authorization", "\"tokens\"", noPermissions, nil, simpleSettings("CAN_USE")),
	}, Step{
		Template: makePermissionsTestStage("authorization", "\"tokens\"", noPermissions, nil, []permissionSettings{{name: "users", skipCreation: true, permissionLevel: "CAN_USE"}}),
	})
}

func TestAccPermissions_SqlWarehouses(t *testing.T) {
	sqlWarehouseTemplate := `
	resource "databricks_sql_global_config" "this" {
		enable_serverless_compute = true
	}
	resource "databricks_sql_endpoint" "this" {
		depends_on = [databricks_sql_global_config.this]
		name = "{var.STICKY_RANDOM}"
		cluster_size = "2X-Small"
		enable_serverless_compute = true
	}`
	// Note: ideally we could test making a new user the owner of the warehouse, but the new user
	// needs cluster creation permissions, and the SCIM API doesn't provide get-after-put consistency,
	// so this would introduce flakiness.
	WorkspaceLevel(t, Step{
		Template: sqlWarehouseTemplate + makePermissionsTestStage("sql_endpoint_id", "databricks_sql_endpoint.this.id", noPermissions, nil, simpleSettings("CAN_USE")),
	}, Step{
		Template:    sqlWarehouseTemplate + makePermissionsTestStage("sql_endpoint_id", "databricks_sql_endpoint.this.id", "CAN_USE", nil, simpleSettings("CAN_USE", "CAN_MANAGE", "CAN_MONITOR")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for warehouses, allowed levels: CAN_MANAGE, IS_OWNER"),
	}, Step{
		Template: sqlWarehouseTemplate,
		Check: resourceCheck("databricks_sql_endpoint.this", func(ctx context.Context, c *common.DatabricksClient, id string) error {
			w, err := c.WorkspaceClient()
			assert.NoError(t, err)
			warehouse, err := w.Warehouses.GetById(context.Background(), id)
			assert.NoError(t, err)
			permissions, err := w.Warehouses.GetPermissionsByWarehouseId(context.Background(), id)
			assert.NoError(t, err)
			found := false
			for _, acl := range permissions.AccessControlList {
				if acl.ServicePrincipalName == warehouse.CreatorName {
					assert.Equal(t, sql.WarehousePermissionLevel("IS_OWNER"), acl.AllPermissions[0].PermissionLevel)
					found = true
					break
				}
			}
			if !found {
				assert.Fail(t, "pipeline creator not found in permissions")
			}
			return nil
		}),
	})
}

func TestAccPermissions_SqlDashboard(t *testing.T) {
	dashboardTemplate := `
	resource "databricks_sql_dashboard" "this" {
		name = "{var.STICKY_RANDOM}"
	}`
	WorkspaceLevel(t, Step{
		Template: dashboardTemplate + makePermissionsTestStage("sql_dashboard_id", "databricks_sql_dashboard.this.id", noPermissions, nil, simpleSettings("CAN_VIEW")),
	}, Step{
		Template: dashboardTemplate + makePermissionsTestStage("sql_dashboard_id", "databricks_sql_dashboard.this.id", "CAN_MANAGE", nil, simpleSettings("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template:    dashboardTemplate + makePermissionsTestStage("sql_dashboard_id", "databricks_sql_dashboard.this.id", "CAN_VIEW", nil, simpleSettings("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for dashboard, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_SqlAlert(t *testing.T) {
	alertTemplate := `
	resource "databricks_sql_query" "this" {
		name = "{var.STICKY_RANDOM}-query"
		query = "SELECT 1 AS p1, 2 as p2"
		data_source_id = "{env.TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID}"
	}
	resource "databricks_sql_alert" "this" {
		name = "{var.STICKY_RANDOM}-alert"
		query_id = databricks_sql_query.this.id
		options {
			column = "p1"
			op = ">="
			value = "3"
			muted = false
		}
	}`
	WorkspaceLevel(t, Step{
		Template: alertTemplate + makePermissionsTestStage("sql_alert_id", "databricks_sql_alert.this.id", noPermissions, nil, simpleSettings("CAN_VIEW")),
	}, Step{
		Template: alertTemplate + makePermissionsTestStage("sql_alert_id", "databricks_sql_alert.this.id", "CAN_MANAGE", nil, simpleSettings("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template:    alertTemplate + makePermissionsTestStage("sql_alert_id", "databricks_sql_alert.this.id", "CAN_VIEW", nil, simpleSettings("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for alert, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_SqlQuery(t *testing.T) {
	queryTemplate := `
	resource "databricks_sql_query" "this" {
		name = "{var.STICKY_RANDOM}-query"
		query = "SELECT 1 AS p1, 2 as p2"
		data_source_id = "{env.TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID}"
	}`
	WorkspaceLevel(t, Step{
		Template: queryTemplate + makePermissionsTestStage("sql_query_id", "databricks_sql_query.this.id", noPermissions, nil, simpleSettings("CAN_VIEW")),
	}, Step{
		Template: queryTemplate + makePermissionsTestStage("sql_query_id", "databricks_sql_query.this.id", "CAN_MANAGE", nil, simpleSettings("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template:    queryTemplate + makePermissionsTestStage("sql_query_id", "databricks_sql_query.this.id", "CAN_VIEW", nil, simpleSettings("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for query, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Dashboard(t *testing.T) {
	dashboardTemplate := `
	resource "databricks_directory" "this" {
		path = "/permissions_test/{var.STICKY_RANDOM}"
	}
	resource "databricks_dashboard" "dashboard" {
		display_name = "TF New Dashboard"
		warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
		parent_path = databricks_directory.this.path
	}
	`
	WorkspaceLevel(t, Step{
		Template: dashboardTemplate + makePermissionsTestStage("dashboard_id", "databricks_dashboard.dashboard.id", noPermissions, nil, simpleSettings("CAN_READ")),
	}, Step{
		Template: dashboardTemplate + makePermissionsTestStage("dashboard_id", "databricks_dashboard.dashboard.id", "CAN_MANAGE", nil, simpleSettings("CAN_READ", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template:    dashboardTemplate + makePermissionsTestStage("dashboard_id", "databricks_dashboard.dashboard.id", "CAN_READ", nil, simpleSettings("CAN_READ", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for dashboard, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Experiment(t *testing.T) {
	experimentTemplate := `
	resource "databricks_directory" "this" {
		path = "/permissions_test/{var.STICKY_RANDOM}"
	}
	resource "databricks_mlflow_experiment" "this" {
		name = "${databricks_directory.this.path}/experiment"
	}`
	WorkspaceLevel(t, Step{
		Template: experimentTemplate + makePermissionsTestStage("experiment_id", "databricks_mlflow_experiment.this.id", noPermissions, nil, simpleSettings("CAN_READ")),
	}, Step{
		Template: experimentTemplate + makePermissionsTestStage("experiment_id", "databricks_mlflow_experiment.this.id", "CAN_MANAGE", nil, simpleSettings("CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    experimentTemplate + makePermissionsTestStage("experiment_id", "databricks_mlflow_experiment.this.id", "CAN_READ", nil, simpleSettings("CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for mlflowExperiment, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_RegisteredModel(t *testing.T) {
	modelTemplate := `
	resource "databricks_mlflow_model" "m1" {
		name = "tf-{var.STICKY_RANDOM}"
		description = "tf-{var.STICKY_RANDOM} description"
	}
	`
	WorkspaceLevel(t, Step{
		Template: modelTemplate + makePermissionsTestStage("registered_model_id", "databricks_mlflow_model.m1.registered_model_id", noPermissions, nil, simpleSettings("CAN_READ")),
	}, Step{
		Template: modelTemplate + makePermissionsTestStage("registered_model_id", "databricks_mlflow_model.m1.registered_model_id", "CAN_MANAGE", nil, simpleSettings("CAN_READ", "CAN_EDIT", "CAN_MANAGE_STAGING_VERSIONS", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE")),
	}, Step{
		Template:    modelTemplate + makePermissionsTestStage("registered_model_id", "databricks_mlflow_model.m1.registered_model_id", "CAN_READ", nil, simpleSettings("CAN_READ", "CAN_EDIT", "CAN_MANAGE_STAGING_VERSIONS", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for registered-model, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_ServingEndpoint(t *testing.T) {
	skipf(t)("Serving endpoint update is flaky")
	endpointTemplate := `
	resource "databricks_model_serving" "endpoint" {
		name = "{var.STICKY_RANDOM}"
		config {
			served_models {
				name = "prod_model"
				model_name = "experiment-fixture-model"
				model_version = "1"
				workload_size = "Small"
				scale_to_zero_enabled = true
			}
			traffic_config {
				routes {
					served_model_name = "prod_model"
					traffic_percentage = 100
				}
			}
		}
	}`
	WorkspaceLevel(t, Step{
		Template: endpointTemplate + makePermissionsTestStage("serving_endpoint_id", "databricks_model_serving.endpoint.id", noPermissions, nil, simpleSettings("CAN_VIEW")),
	}, Step{
		Template: endpointTemplate + makePermissionsTestStage("serving_endpoint_id", "databricks_model_serving.endpoint.id", "CAN_MANAGE", nil, simpleSettings("CAN_VIEW", "CAN_QUERY", "CAN_MANAGE")),
	}, Step{
		Template:    endpointTemplate + makePermissionsTestStage("serving_endpoint_id", "databricks_model_serving.endpoint.id", "CAN_VIEW", nil, simpleSettings("CAN_VIEW", "CAN_QUERY", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for servingEndpoint, allowed levels: CAN_MANAGE"),
	})
}
