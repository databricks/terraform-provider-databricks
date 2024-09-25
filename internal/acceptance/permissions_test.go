package acceptance

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

//
// databricks_permissions testing support
//

type permissionSettings struct {
	// Name of the SP or group
	name string
	// If true, the resource will not be created
	skipCreation    bool
	permissionLevel string
}

type makePermissionsConfig struct {
	currentUserPermission string
	servicePrincipal      []permissionSettings
	group                 []permissionSettings
	user                  []permissionSettings
}

func servicePrincipalPermissions(permissionLevel ...string) func(*makePermissionsConfig) {
	return func(config *makePermissionsConfig) {
		if config.servicePrincipal != nil {
			panic("service_principal permissions already set")
		}
		config.servicePrincipal = simpleSettings(permissionLevel...)
	}
}

func groupPermissions(permissionLevel ...string) func(*makePermissionsConfig) {
	return func(config *makePermissionsConfig) {
		if config.group != nil {
			panic("group permissions already set")
		}
		config.group = simpleSettings(permissionLevel...)
	}
}

func userPermissions(permissionLevel ...string) func(*makePermissionsConfig) {
	return func(config *makePermissionsConfig) {
		if config.user != nil {
			panic("user permissions already set")
		}
		config.user = simpleSettings(permissionLevel...)
	}
}

func allPrincipalPermissions(permissionLevel ...string) func(*makePermissionsConfig) {
	return func(config *makePermissionsConfig) {
		if config.servicePrincipal != nil || config.group != nil || config.user != nil {
			panic("permissions already set")
		}
		config.servicePrincipal = simpleSettings(permissionLevel...)
		config.group = simpleSettings(permissionLevel...)
		config.user = simpleSettings(permissionLevel...)
	}
}

func currentUserPermission(permissionLevel string) func(*makePermissionsConfig) {
	return func(config *makePermissionsConfig) {
		if config.currentUserPermission != "" {
			panic("current user permission already set")
		}
		config.currentUserPermission = permissionLevel
	}
}

func customPermission(name string, permissionSettings permissionSettings) func(*makePermissionsConfig) {
	return func(config *makePermissionsConfig) {
		switch name {
		case "service_principal":
			config.servicePrincipal = append(config.servicePrincipal, permissionSettings)
		case "group":
			config.group = append(config.group, permissionSettings)
		case "user":
			config.user = append(config.user, permissionSettings)
		default:
			panic(fmt.Sprintf("unknown permission type: %s", name))
		}
	}
}

func simpleSettings(permissionLevel ...string) []permissionSettings {
	var settings []permissionSettings
	for _, level := range permissionLevel {
		settings = append(settings, permissionSettings{permissionLevel: level})
	}
	return settings
}

func makePermissionsTestStage(idAttribute, idValue string, permissionOptions ...func(*makePermissionsConfig)) string {
	config := makePermissionsConfig{}
	for _, option := range permissionOptions {
		option(&config)
	}
	var resources string
	var accessControlBlocks string
	addPermissions := func(permissionSettings []permissionSettings, resourceType, resourceNameAttribute, idAttribute, accessControlAttribute string, getName func(int) string) {
		for i, permission := range permissionSettings {
			if !permission.skipCreation {
				resources += fmt.Sprintf(`
				resource "%s" "_%d" {
					%s = "%s"
				}`, resourceType, i, resourceNameAttribute, getName(i))
			}
			var name string
			if permission.name == "" {
				name = fmt.Sprintf("%s._%d.%s", resourceType, i, idAttribute)
			} else {
				name = fmt.Sprintf(`"%s"`, permission.name)
			}
			accessControlBlocks += fmt.Sprintf(`
			access_control {
				%s = %s
				permission_level = "%s"
			}`, accessControlAttribute, name, permission.permissionLevel)
		}
	}
	addPermissions(config.servicePrincipal, "databricks_service_principal", "display_name", "application_id", "service_principal_name", func(i int) string {
		return fmt.Sprintf("{var.STICKY_RANDOM}-%d", i)
	})
	addPermissions(config.group, "databricks_group", "display_name", "display_name", "group_name", func(i int) string {
		return fmt.Sprintf("{var.STICKY_RANDOM}-%d", i)
	})
	addPermissions(config.user, "databricks_user", "user_name", "user_name", "user_name", func(i int) string {
		return fmt.Sprintf("{var.STICKY_RANDOM}-%d@databricks.com", i)
	})
	if config.currentUserPermission != "" {
		accessControlBlocks += fmt.Sprintf(`
		access_control {
			permission_level = "%s"
			service_principal_name = data.databricks_current_user.me.user_name
		}`, config.currentUserPermission)
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

func assertContainsPermission(t *testing.T, permissions *iam.ObjectPermissions, principalType, name string, permissionLevel iam.PermissionLevel) {
	for _, acl := range permissions.AccessControlList {
		switch principalType {
		case "user":
			if acl.UserName == name {
				assert.Equal(t, permissionLevel, acl.AllPermissions[0].PermissionLevel)
				return
			}
		case "service_principal":
			if acl.ServicePrincipalName == name {
				assert.Equal(t, permissionLevel, acl.AllPermissions[0].PermissionLevel)
				return
			}
		case "group":
			if acl.GroupName == name {
				assert.Equal(t, permissionLevel, acl.AllPermissions[0].PermissionLevel)
				return
			}
		}
	}
	assert.Fail(t, "permission not found for %s %s", principalType, name)
}

//
// databricks_permissions acceptance tests
//

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
		Template: policyTemplate + makePermissionsTestStage("cluster_policy_id", "databricks_cluster_policy.this.id", groupPermissions("CAN_USE")),
	}, Step{
		Template: policyTemplate + makePermissionsTestStage("cluster_policy_id", "databricks_cluster_policy.this.id", currentUserPermission("CAN_USE"), allPrincipalPermissions("CAN_USE")),
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
		Template: policyTemplate + makePermissionsTestStage("instance_pool_id", "databricks_instance_pool.this.id", groupPermissions("CAN_ATTACH_TO")),
	}, Step{
		Template: policyTemplate + makePermissionsTestStage("instance_pool_id", "databricks_instance_pool.this.id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_ATTACH_TO", "CAN_MANAGE")),
	}, Step{
		Template:    policyTemplate + makePermissionsTestStage("instance_pool_id", "databricks_instance_pool.this.id", currentUserPermission("CAN_ATTACH_TO")),
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
		Template: policyTemplate + makePermissionsTestStage("cluster_id", "databricks_cluster.this.id", groupPermissions("CAN_ATTACH_TO")),
	}, Step{
		Template: policyTemplate + makePermissionsTestStage("cluster_id", "databricks_cluster.this.id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_ATTACH_TO", "CAN_RESTART", "CAN_MANAGE")),
	}, Step{
		Template:    policyTemplate + makePermissionsTestStage("cluster_id", "databricks_cluster.this.id", currentUserPermission("CAN_ATTACH_TO")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for cluster, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Job(t *testing.T) {
	template := `
	resource "databricks_job" "this" {
		name = "{var.STICKY_RANDOM}"
	}`
	WorkspaceLevel(t, Step{
		Template: template + makePermissionsTestStage("job_id", "databricks_job.this.id", groupPermissions("CAN_VIEW")),
	}, Step{
		Template: template + makePermissionsTestStage("job_id", "databricks_job.this.id", currentUserPermission("IS_OWNER"), allPrincipalPermissions("CAN_VIEW", "CAN_MANAGE_RUN", "CAN_MANAGE")),
	}, Step{
		Template: template + makePermissionsTestStage("job_id", "databricks_job.this.id", currentUserPermission("CAN_MANAGE"), userPermissions("IS_OWNER")),
	}, Step{
		Template:    template + makePermissionsTestStage("job_id", "databricks_job.this.id", currentUserPermission("CAN_MANAGE_RUN")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for job, allowed levels: CAN_MANAGE, IS_OWNER"),
	}, Step{
		Template: template,
		Check: func(s *terraform.State) error {
			w := databricks.Must(databricks.NewWorkspaceClient())
			jobId := s.RootModule().Resources["databricks_job.this"].Primary.ID
			permissions, err := w.Permissions.GetByRequestObjectTypeAndRequestObjectId(context.Background(), "jobs", jobId)
			assert.NoError(t, err)
			idInt, err := strconv.Atoi(jobId)
			assert.NoError(t, err)
			job, err := w.Jobs.GetByJobId(context.Background(), int64(idInt))
			assert.NoError(t, err)
			assertContainsPermission(t, permissions, "service_principal", job.CreatorUserName, iam.PermissionLevelIsOwner)
			return nil
		},
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
		Template: policyTemplate + makePermissionsTestStage("pipeline_id", "databricks_pipeline.this.id", groupPermissions("CAN_VIEW")),
	}, Step{
		Template: policyTemplate + makePermissionsTestStage("pipeline_id", "databricks_pipeline.this.id", currentUserPermission("IS_OWNER"), allPrincipalPermissions("CAN_VIEW", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template: policyTemplate + makePermissionsTestStage("pipeline_id", "databricks_pipeline.this.id", currentUserPermission("CAN_MANAGE"), userPermissions("IS_OWNER"), groupPermissions("CAN_VIEW", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template:    policyTemplate + makePermissionsTestStage("pipeline_id", "databricks_pipeline.this.id", currentUserPermission("CAN_RUN")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for pipelines, allowed levels: CAN_MANAGE, IS_OWNER"),
	}, Step{
		Template: policyTemplate,
		Check: resourceCheck("databricks_pipeline.this", func(ctx context.Context, c *common.DatabricksClient, id string) error {
			w, err := c.WorkspaceClient()
			assert.NoError(t, err)
			pipeline, err := w.Pipelines.GetByPipelineId(context.Background(), id)
			assert.NoError(t, err)
			permissions, err := w.Permissions.GetByRequestObjectTypeAndRequestObjectId(context.Background(), "pipelines", id)
			assert.NoError(t, err)
			assertContainsPermission(t, permissions, "service_principal", pipeline.CreatorUserName, iam.PermissionLevelIsOwner)
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
		Template: notebookTemplate + makePermissionsTestStage("notebook_path", "databricks_notebook.this.id", groupPermissions("CAN_RUN")),
	}, Step{
		Template: notebookTemplate + makePermissionsTestStage("notebook_path", "databricks_notebook.this.id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: notebookTemplate + makePermissionsTestStage("notebook_path", "databricks_notebook.this.id", allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    notebookTemplate + makePermissionsTestStage("notebook_path", "databricks_notebook.this.id", currentUserPermission("CAN_READ")),
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
		Template: notebookTemplate + makePermissionsTestStage("notebook_id", "databricks_notebook.this.object_id", groupPermissions("CAN_RUN")),
	}, Step{
		Template: notebookTemplate + makePermissionsTestStage("notebook_id", "databricks_notebook.this.object_id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: notebookTemplate + makePermissionsTestStage("notebook_id", "databricks_notebook.this.object_id", allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    notebookTemplate + makePermissionsTestStage("notebook_id", "databricks_notebook.this.object_id", currentUserPermission("CAN_READ")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for notebook, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Directory_Path(t *testing.T) {
	directoryTemplate := `
	resource "databricks_directory" "this" {
		path = "/permissions_test/{var.STICKY_RANDOM}"
	}`
	WorkspaceLevel(t, Step{
		Template: directoryTemplate + makePermissionsTestStage("directory_path", "databricks_directory.this.id", groupPermissions("CAN_RUN")),
	}, Step{
		Template: directoryTemplate + makePermissionsTestStage("directory_path", "databricks_directory.this.id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: directoryTemplate + makePermissionsTestStage("directory_path", "databricks_directory.this.id", allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    directoryTemplate + makePermissionsTestStage("directory_path", "databricks_directory.this.id", currentUserPermission("CAN_READ")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for directory, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Directory_Id(t *testing.T) {
	directoryTemplate := `
	resource "databricks_directory" "this" {
		path = "/permissions_test/{var.STICKY_RANDOM}"
	}`
	WorkspaceLevel(t, Step{
		Template: directoryTemplate + makePermissionsTestStage("directory_id", "databricks_directory.this.object_id", groupPermissions("CAN_RUN")),
	}, Step{
		Template: directoryTemplate + makePermissionsTestStage("directory_id", "databricks_directory.this.object_id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: directoryTemplate + makePermissionsTestStage("directory_id", "databricks_directory.this.object_id", allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    directoryTemplate + makePermissionsTestStage("directory_id", "databricks_directory.this.object_id", currentUserPermission("CAN_READ")),
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
		Template: workspaceFile + makePermissionsTestStage("workspace_file_path", "databricks_workspace_file.this.id", groupPermissions("CAN_RUN")),
	}, Step{
		Template: workspaceFile + makePermissionsTestStage("workspace_file_path", "databricks_workspace_file.this.id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: workspaceFile + makePermissionsTestStage("workspace_file_path", "databricks_workspace_file.this.id", allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    workspaceFile + makePermissionsTestStage("workspace_file_path", "databricks_workspace_file.this.id", currentUserPermission("CAN_READ")),
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
		Template: workspaceFile + makePermissionsTestStage("workspace_file_id", "databricks_workspace_file.this.object_id", groupPermissions("CAN_RUN")),
	}, Step{
		Template: workspaceFile + makePermissionsTestStage("workspace_file_id", "databricks_workspace_file.this.object_id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		// The current user can be removed from permissions since they inherit permissions from the directory they created.
		Template: workspaceFile + makePermissionsTestStage("workspace_file_id", "databricks_workspace_file.this.object_id", allPrincipalPermissions("CAN_RUN", "CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    workspaceFile + makePermissionsTestStage("workspace_file_id", "databricks_workspace_file.this.object_id", currentUserPermission("CAN_READ")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for file, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Repo_Id(t *testing.T) {
	template := `
	resource "databricks_repo" "this" {
		url = "https://github.com/databrickslabs/tempo.git"
		path = "/Repos/terraform-tests/tempo-{var.STICKY_RANDOM}"
	}
	`
	WorkspaceLevel(t, Step{
		Template: template + makePermissionsTestStage("repo_id", "databricks_repo.this.id", groupPermissions("CAN_MANAGE", "CAN_READ")),
		// Check: resource.ComposeTestCheckFunc(
		// 	resource.TestCheckResourceAttr("databricks_permissions.this", "object_type", "repo"),
		// 	func(s *terraform.State) error {
		// 		w := databricks.Must(databricks.NewWorkspaceClient())
		// 		repoId := s.RootModule().Resources["databricks_repo.this"].Primary.ID
		// 		permissions, err := w.Permissions.GetByRequestObjectTypeAndRequestObjectId(context.Background(), "repos", repoId)
		// 		assert.NoError(t, err)
		// 		group1Name := s.RootModule().Resources["databricks_group._0"].Primary.Attributes["display_name"]
		// 		assertContainsPermission(t, permissions, "group", group1Name, iam.PermissionLevelCanManage)
		// 		group2Name := s.RootModule().Resources["databricks_group._1"].Primary.Attributes["display_name"]
		// 		assertContainsPermission(t, permissions, "group", group2Name, iam.PermissionLevelCanRun)
		// 		return nil
		// 	},
		// ),
	}, Step{
		Template: template + makePermissionsTestStage("repo_id", "databricks_repo.this.id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_READ", "CAN_MANAGE", "CAN_RUN", "CAN_EDIT")),
	}, Step{
		Template: template + makePermissionsTestStage("repo_id", "databricks_repo.this.id", allPrincipalPermissions("CAN_READ", "CAN_MANAGE", "CAN_RUN", "CAN_EDIT")),
	}, Step{
		Template:    template + makePermissionsTestStage("repo_id", "databricks_repo.this.id", currentUserPermission("CAN_READ")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for repo, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Repo_Path(t *testing.T) {
	template := `
	resource "databricks_repo" "this" {
		url = "https://github.com/databrickslabs/tempo.git"
		path = "/Repos/terraform-tests/tempo-{var.STICKY_RANDOM}"
	}
	`
	WorkspaceLevel(t, Step{
		Template: template + makePermissionsTestStage("repo_path", "databricks_repo.this.path", groupPermissions("CAN_MANAGE", "CAN_RUN")),
		// Check: resource.ComposeTestCheckFunc(
		// 	resource.TestCheckResourceAttr("databricks_permissions.this", "object_type", "repo"),
		// 	func(s *terraform.State) error {
		// 		w := databricks.Must(databricks.NewWorkspaceClient())
		// 		repoId := s.RootModule().Resources["databricks_repo.this"].Primary.ID
		// 		permissions, err := w.Permissions.GetByRequestObjectTypeAndRequestObjectId(context.Background(), "repos", repoId)
		// 		assert.NoError(t, err)
		// 		group1Name := s.RootModule().Resources["databricks_group._0"].Primary.Attributes["display_name"]
		// 		assertContainsPermission(t, permissions, "group", group1Name, iam.PermissionLevelCanManage)
		// 		group2Name := s.RootModule().Resources["databricks_group._1"].Primary.Attributes["display_name"]
		// 		assertContainsPermission(t, permissions, "group", group2Name, iam.PermissionLevelCanRun)
		// 		return nil
		// 	},
		// ),
	}, Step{
		Template: template + makePermissionsTestStage("repo_id", "databricks_repo.this.id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_READ", "CAN_MANAGE", "CAN_RUN", "CAN_EDIT")),
	}, Step{
		Template: template + makePermissionsTestStage("repo_id", "databricks_repo.this.id", allPrincipalPermissions("CAN_READ", "CAN_MANAGE", "CAN_RUN", "CAN_EDIT")),
	}, Step{
		Template:    template + makePermissionsTestStage("repo_id", "databricks_repo.this.id", currentUserPermission("CAN_READ")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for repo, allowed levels: CAN_MANAGE"),
	})
}

func TestAccPermissions_Authorization_Passwords(t *testing.T) {
	skipf(t)("ACLs for passwords are disabled on testing workspaces")
	WorkspaceLevel(t, Step{
		Template: makePermissionsTestStage("authorization", "\"passwords\"", groupPermissions("CAN_USE")),
	}, Step{
		Template: makePermissionsTestStage("authorization", "\"passwords\"", customPermission("group", permissionSettings{name: "admins", skipCreation: true, permissionLevel: "CAN_USE"})),
	})
}

func TestAccPermissions_Authorization_Tokens(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: makePermissionsTestStage("authorization", "\"tokens\"", groupPermissions("CAN_USE")),
	}, Step{
		Template: makePermissionsTestStage("authorization", "\"tokens\"", customPermission("group", permissionSettings{name: "users", skipCreation: true, permissionLevel: "CAN_USE"})),
	})
}

func TestAccPermissions_SqlWarehouses(t *testing.T) {
	sqlWarehouseTemplate := `
	resource "databricks_sql_endpoint" "this" {
		depends_on = [databricks_sql_global_config.this]
		name = "{var.STICKY_RANDOM}"
		cluster_size = "2X-Small"
	}`
	// Note: ideally we could test making a new user the owner of the warehouse, but the new user
	// needs cluster creation permissions, and the SCIM API doesn't provide get-after-put consistency,
	// so this would introduce flakiness.
	WorkspaceLevel(t, Step{
		Template: sqlWarehouseTemplate + makePermissionsTestStage("sql_endpoint_id", "databricks_sql_endpoint.this.id", groupPermissions("CAN_USE")),
	}, Step{
		Template: sqlWarehouseTemplate + makePermissionsTestStage("sql_endpoint_id", "databricks_sql_endpoint.this.id", currentUserPermission("IS_OWNER"), allPrincipalPermissions("CAN_USE", "CAN_MANAGE", "CAN_MONITOR")),
	}, Step{
		Template: sqlWarehouseTemplate + makePermissionsTestStage("sql_endpoint_id", "databricks_sql_endpoint.this.id", currentUserPermission("CAN_MANAGE"), servicePrincipalPermissions("IS_OWNER")),
	}, Step{
		Template:    sqlWarehouseTemplate + makePermissionsTestStage("sql_endpoint_id", "databricks_sql_endpoint.this.id", currentUserPermission("CAN_USE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for warehouses, allowed levels: CAN_MANAGE, IS_OWNER"),
	}, Step{
		Template: sqlWarehouseTemplate,
		Check: func(s *terraform.State) error {
			w := databricks.Must(databricks.NewWorkspaceClient())
			id := s.RootModule().Resources["databricks_sql_endpoint.this"].Primary.ID
			warehouse, err := w.Warehouses.GetById(context.Background(), id)
			assert.NoError(t, err)
			permissions, err := w.Permissions.GetByRequestObjectTypeAndRequestObjectId(context.Background(), "warehouses", id)
			assert.NoError(t, err)
			assertContainsPermission(t, permissions, "service_principal", warehouse.CreatorName, iam.PermissionLevelIsOwner)
			return nil
		},
	})
}

func TestAccPermissions_SqlDashboard(t *testing.T) {
	dashboardTemplate := `
	resource "databricks_sql_dashboard" "this" {
		name = "{var.STICKY_RANDOM}"
	}`
	WorkspaceLevel(t, Step{
		Template: dashboardTemplate + makePermissionsTestStage("sql_dashboard_id", "databricks_sql_dashboard.this.id", groupPermissions("CAN_VIEW")),
	}, Step{
		Template: dashboardTemplate + makePermissionsTestStage("sql_dashboard_id", "databricks_sql_dashboard.this.id", currentUserPermission("CAN_MANAGE"), allPrincipalPermissions("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template:    dashboardTemplate + makePermissionsTestStage("sql_dashboard_id", "databricks_sql_dashboard.this.id", currentUserPermission("CAN_VIEW")),
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
		Template: alertTemplate + makePermissionsTestStage("sql_alert_id", "databricks_sql_alert.this.id", groupPermissions("CAN_VIEW")),
	}, Step{
		Template: alertTemplate + makePermissionsTestStage("sql_alert_id", "databricks_sql_alert.this.id", currentUserPermission("CAN_MANAGE"), groupPermissions("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template:    alertTemplate + makePermissionsTestStage("sql_alert_id", "databricks_sql_alert.this.id", currentUserPermission("CAN_VIEW"), groupPermissions("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
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
		Template: queryTemplate + makePermissionsTestStage("sql_query_id", "databricks_sql_query.this.id", groupPermissions("CAN_VIEW")),
	}, Step{
		Template: queryTemplate + makePermissionsTestStage("sql_query_id", "databricks_sql_query.this.id", currentUserPermission("CAN_MANAGE"), groupPermissions("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template:    queryTemplate + makePermissionsTestStage("sql_query_id", "databricks_sql_query.this.id", currentUserPermission("CAN_VIEW"), groupPermissions("CAN_VIEW", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
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
		Template: dashboardTemplate + makePermissionsTestStage("dashboard_id", "databricks_dashboard.dashboard.id", groupPermissions("CAN_READ")),
	}, Step{
		Template: dashboardTemplate + makePermissionsTestStage("dashboard_id", "databricks_dashboard.dashboard.id", currentUserPermission("CAN_MANAGE"), groupPermissions("CAN_READ", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
	}, Step{
		Template:    dashboardTemplate + makePermissionsTestStage("dashboard_id", "databricks_dashboard.dashboard.id", currentUserPermission("CAN_READ"), groupPermissions("CAN_READ", "CAN_EDIT", "CAN_RUN", "CAN_MANAGE")),
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
		Template: experimentTemplate + makePermissionsTestStage("experiment_id", "databricks_mlflow_experiment.this.id", groupPermissions("CAN_READ")),
	}, Step{
		Template: experimentTemplate + makePermissionsTestStage("experiment_id", "databricks_mlflow_experiment.this.id", currentUserPermission("CAN_MANAGE"), groupPermissions("CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
	}, Step{
		Template:    experimentTemplate + makePermissionsTestStage("experiment_id", "databricks_mlflow_experiment.this.id", currentUserPermission("CAN_READ"), groupPermissions("CAN_READ", "CAN_EDIT", "CAN_MANAGE")),
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
		Template: modelTemplate + makePermissionsTestStage("registered_model_id", "databricks_mlflow_model.m1.registered_model_id", groupPermissions("CAN_READ")),
	}, Step{
		Template: modelTemplate + makePermissionsTestStage("registered_model_id", "databricks_mlflow_model.m1.registered_model_id", currentUserPermission("CAN_MANAGE"), groupPermissions("CAN_READ", "CAN_EDIT", "CAN_MANAGE_STAGING_VERSIONS", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE")),
	}, Step{
		Template:    modelTemplate + makePermissionsTestStage("registered_model_id", "databricks_mlflow_model.m1.registered_model_id", currentUserPermission("CAN_READ"), groupPermissions("CAN_READ", "CAN_EDIT", "CAN_MANAGE_STAGING_VERSIONS", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE")),
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
		Template: endpointTemplate + makePermissionsTestStage("serving_endpoint_id", "databricks_model_serving.endpoint.id", groupPermissions("CAN_VIEW")),
	}, Step{
		Template: endpointTemplate + makePermissionsTestStage("serving_endpoint_id", "databricks_model_serving.endpoint.id", currentUserPermission("CAN_MANAGE"), groupPermissions("CAN_VIEW", "CAN_QUERY", "CAN_MANAGE")),
	}, Step{
		Template:    endpointTemplate + makePermissionsTestStage("serving_endpoint_id", "databricks_model_serving.endpoint.id", currentUserPermission("CAN_VIEW"), groupPermissions("CAN_VIEW", "CAN_QUERY", "CAN_MANAGE")),
		ExpectError: regexp.MustCompile("cannot remove management permissions for the current user for servingEndpoint, allowed levels: CAN_MANAGE"),
	})
}
