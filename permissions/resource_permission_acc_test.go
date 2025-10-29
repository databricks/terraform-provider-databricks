package permissions_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccPermission_Cluster(t *testing.T) {
	acceptance.LoadDebugEnvIfRunsFromIDE(t, "workspace")
	clusterTemplate := `
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
}

resource "databricks_group" "group1" {
	display_name = "permission-group1-{var.RANDOM}"
}

resource "databricks_group" "group2" {
	display_name = "permission-group2-{var.RANDOM}"
}

resource "databricks_permission" "cluster_group1" {
	cluster_id       = databricks_cluster.this.id
	group_name       = databricks_group.group1.display_name
	permission_level = "CAN_ATTACH_TO"
}

resource "databricks_permission" "cluster_group2" {
	cluster_id       = databricks_cluster.this.id
	group_name       = databricks_group.group2.display_name
	permission_level = "CAN_RESTART"
}
`

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: clusterTemplate,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permission.cluster_group1", "permission_level", "CAN_ATTACH_TO"),
			resource.TestCheckResourceAttr("databricks_permission.cluster_group2", "permission_level", "CAN_RESTART"),
			func(s *terraform.State) error {
				w := databricks.Must(databricks.NewWorkspaceClient())
				clusterId := s.RootModule().Resources["databricks_cluster.this"].Primary.ID
				permissions, err := w.Permissions.GetByRequestObjectTypeAndRequestObjectId(context.Background(), "clusters", clusterId)
				assert.NoError(t, err)

				group1Name := s.RootModule().Resources["databricks_group.group1"].Primary.Attributes["display_name"]
				group2Name := s.RootModule().Resources["databricks_group.group2"].Primary.Attributes["display_name"]

				// Verify both permissions exist
				foundGroup1 := false
				foundGroup2 := false
				for _, acl := range permissions.AccessControlList {
					if acl.GroupName == group1Name {
						assert.Equal(t, iam.PermissionLevelCanAttachTo, acl.AllPermissions[0].PermissionLevel)
						foundGroup1 = true
					}
					if acl.GroupName == group2Name {
						assert.Equal(t, iam.PermissionLevelCanRestart, acl.AllPermissions[0].PermissionLevel)
						foundGroup2 = true
					}
				}
				assert.True(t, foundGroup1, "Group1 permission not found")
				assert.True(t, foundGroup2, "Group2 permission not found")
				return nil
			},
		),
	})
}

func TestAccPermission_Job(t *testing.T) {
	acceptance.LoadDebugEnvIfRunsFromIDE(t, "workspace")
	template := `
resource "databricks_job" "this" {
	name = "permission-job-{var.RANDOM}"
}

resource "databricks_group" "viewers" {
	display_name = "permission-viewers-{var.RANDOM}"
}

resource "databricks_user" "test_user" {
	user_name = "permission-test-{var.RANDOM}@example.com"
	force = true
}

resource "databricks_permission" "job_group" {
	job_id           = databricks_job.this.id
	group_name       = databricks_group.viewers.display_name
	permission_level = "CAN_VIEW"
}

resource "databricks_permission" "job_user" {
	job_id           = databricks_job.this.id
	user_name        = databricks_user.test_user.user_name
	permission_level = "CAN_MANAGE_RUN"
}
`

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permission.job_group", "permission_level", "CAN_VIEW"),
			resource.TestCheckResourceAttr("databricks_permission.job_user", "permission_level", "CAN_MANAGE_RUN"),
			resource.TestCheckResourceAttr("databricks_permission.job_group", "object_type", "jobs"),
		),
	})
}

func TestAccPermission_Notebook(t *testing.T) {
	acceptance.LoadDebugEnvIfRunsFromIDE(t, "workspace")
	template := `
resource "databricks_directory" "this" {
	path = "/permission-test-{var.RANDOM}"
}

resource "databricks_notebook" "this" {
	source = "{var.CWD}/../storage/testdata/tf-test-python.py"
	path = "${databricks_directory.this.path}/test_notebook"
}

resource "databricks_group" "editors" {
	display_name = "permission-editors-{var.RANDOM}"
}

resource "databricks_group" "runners" {
	display_name = "permission-runners-{var.RANDOM}"
}

resource "databricks_permission" "notebook_editors" {
	notebook_path    = databricks_notebook.this.path
	group_name       = databricks_group.editors.display_name
	permission_level = "CAN_EDIT"
}

resource "databricks_permission" "notebook_runners" {
	notebook_path    = databricks_notebook.this.path
	group_name       = databricks_group.runners.display_name
	permission_level = "CAN_RUN"
}
`

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permission.notebook_editors", "permission_level", "CAN_EDIT"),
			resource.TestCheckResourceAttr("databricks_permission.notebook_runners", "permission_level", "CAN_RUN"),
		),
	})
}

func TestAccPermission_Authorization_Tokens(t *testing.T) {
	acceptance.LoadDebugEnvIfRunsFromIDE(t, "workspace")
	template := `
resource "databricks_group" "team_a" {
	display_name = "permission-team-a-{var.RANDOM}"
}

resource "databricks_group" "team_b" {
	display_name = "permission-team-b-{var.RANDOM}"
}

resource "databricks_group" "team_c" {
	display_name = "permission-team-c-{var.RANDOM}"
}

# This demonstrates the key benefit: each team's token permissions
# can be managed independently, unlike databricks_permissions
resource "databricks_permission" "tokens_team_a" {
	authorization    = "tokens"
	group_name       = databricks_group.team_a.display_name
	permission_level = "CAN_USE"
}

resource "databricks_permission" "tokens_team_b" {
	authorization    = "tokens"
	group_name       = databricks_group.team_b.display_name
	permission_level = "CAN_USE"
}

resource "databricks_permission" "tokens_team_c" {
	authorization    = "tokens"
	group_name       = databricks_group.team_c.display_name
	permission_level = "CAN_USE"
}
`

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permission.tokens_team_a", "authorization", "tokens"),
			resource.TestCheckResourceAttr("databricks_permission.tokens_team_a", "permission_level", "CAN_USE"),
			resource.TestCheckResourceAttr("databricks_permission.tokens_team_b", "authorization", "tokens"),
			resource.TestCheckResourceAttr("databricks_permission.tokens_team_c", "authorization", "tokens"),
			func(s *terraform.State) error {
				w := databricks.Must(databricks.NewWorkspaceClient())
				permissions, err := w.Permissions.GetByRequestObjectTypeAndRequestObjectId(context.Background(), "authorization", "tokens")
				assert.NoError(t, err)

				teamA := s.RootModule().Resources["databricks_group.team_a"].Primary.Attributes["display_name"]
				teamB := s.RootModule().Resources["databricks_group.team_b"].Primary.Attributes["display_name"]
				teamC := s.RootModule().Resources["databricks_group.team_c"].Primary.Attributes["display_name"]

				foundA, foundB, foundC := false, false, false
				for _, acl := range permissions.AccessControlList {
					if acl.GroupName == teamA {
						foundA = true
					}
					if acl.GroupName == teamB {
						foundB = true
					}
					if acl.GroupName == teamC {
						foundC = true
					}
				}
				assert.True(t, foundA, "Team A permission not found")
				assert.True(t, foundB, "Team B permission not found")
				assert.True(t, foundC, "Team C permission not found")
				return nil
			},
		),
	}, acceptance.Step{
		// Update one permission independently - remove team_b
		Template: `
resource "databricks_group" "team_a" {
	display_name = "permission-team-a-{var.RANDOM}"
}

resource "databricks_group" "team_b" {
	display_name = "permission-team-b-{var.RANDOM}"
}

resource "databricks_group" "team_c" {
	display_name = "permission-team-c-{var.RANDOM}"
}

resource "databricks_permission" "tokens_team_a" {
	authorization    = "tokens"
	group_name       = databricks_group.team_a.display_name
	permission_level = "CAN_USE"
}

resource "databricks_permission" "tokens_team_c" {
	authorization    = "tokens"
	group_name       = databricks_group.team_c.display_name
	permission_level = "CAN_USE"
}
`,
		Check: func(s *terraform.State) error {
			w := databricks.Must(databricks.NewWorkspaceClient())
			permissions, err := w.Permissions.GetByRequestObjectTypeAndRequestObjectId(context.Background(), "authorization", "tokens")
			assert.NoError(t, err)

			teamA := s.RootModule().Resources["databricks_group.team_a"].Primary.Attributes["display_name"]
			teamB := s.RootModule().Resources["databricks_group.team_b"].Primary.Attributes["display_name"]
			teamC := s.RootModule().Resources["databricks_group.team_c"].Primary.Attributes["display_name"]

			foundA, foundB, foundC := false, false, false
			for _, acl := range permissions.AccessControlList {
				if acl.GroupName == teamA {
					foundA = true
				}
				if acl.GroupName == teamB {
					foundB = true
				}
				if acl.GroupName == teamC {
					foundC = true
				}
			}
			assert.True(t, foundA, "Team A permission should still exist")
			assert.False(t, foundB, "Team B permission should be removed")
			assert.True(t, foundC, "Team C permission should still exist")
			return nil
		},
	})
}

func TestAccPermission_Update(t *testing.T) {
	acceptance.LoadDebugEnvIfRunsFromIDE(t, "workspace")
	template1 := `
resource "databricks_job" "this" {
	name = "permission-update-{var.RANDOM}"
}

resource "databricks_group" "test" {
	display_name = "permission-update-{var.RANDOM}"
}

resource "databricks_permission" "job_group" {
	job_id           = databricks_job.this.id
	group_name       = databricks_group.test.display_name
	permission_level = "CAN_VIEW"
}
`

	template2 := `
resource "databricks_job" "this" {
	name = "permission-update-{var.RANDOM}"
}

resource "databricks_group" "test" {
	display_name = "permission-update-{var.RANDOM}"
}

resource "databricks_permission" "job_group" {
	job_id           = databricks_job.this.id
	group_name       = databricks_group.test.display_name
	permission_level = "CAN_MANAGE"
}
`

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template1,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permission.job_group", "permission_level", "CAN_VIEW"),
		),
	}, acceptance.Step{
		Template: template2,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permission.job_group", "permission_level", "CAN_MANAGE"),
			func(s *terraform.State) error {
				w := databricks.Must(databricks.NewWorkspaceClient())
				jobId := s.RootModule().Resources["databricks_job.this"].Primary.ID
				permissions, err := w.Permissions.GetByRequestObjectTypeAndRequestObjectId(context.Background(), "jobs", jobId)
				assert.NoError(t, err)

				groupName := s.RootModule().Resources["databricks_group.test"].Primary.Attributes["display_name"]
				for _, acl := range permissions.AccessControlList {
					if acl.GroupName == groupName {
						assert.Equal(t, iam.PermissionLevelCanManage, acl.AllPermissions[0].PermissionLevel)
						return nil
					}
				}
				return fmt.Errorf("permission not found for group %s", groupName)
			},
		),
	})
}

func TestAccPermission_ServicePrincipal(t *testing.T) {
	acceptance.LoadDebugEnvIfRunsFromIDE(t, "workspace")
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("Service principals have different behavior on GCP")
	}

	template := `
resource "databricks_job" "this" {
	name = "permission-sp-{var.RANDOM}"
}

resource "databricks_service_principal" "sp" {
	display_name = "permission-sp-{var.RANDOM}"
}

resource "databricks_permission" "job_sp" {
	job_id                   = databricks_job.this.id
	service_principal_name   = databricks_service_principal.sp.application_id
	permission_level         = "CAN_MANAGE"
}
`

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permission.job_sp", "permission_level", "CAN_MANAGE"),
			resource.TestCheckResourceAttrSet("databricks_permission.job_sp", "service_principal_name"),
		),
	})
}

func TestAccPermission_Import(t *testing.T) {
	acceptance.LoadDebugEnvIfRunsFromIDE(t, "workspace")
	template := `
resource "databricks_job" "this" {
	name = "permission-import-{var.RANDOM}"
}

resource "databricks_group" "test" {
	display_name = "permission-import-{var.RANDOM}"
}

resource "databricks_permission" "job_group" {
	job_id           = databricks_job.this.id
	group_name       = databricks_group.test.display_name
	permission_level = "CAN_VIEW"
}
`

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
	}, acceptance.Step{
		Template:          template,
		ResourceName:      "databricks_permission.job_group",
		ImportState:       true,
		ImportStateVerify: true,
	})
}

func TestAccPermission_SqlEndpoint(t *testing.T) {
	acceptance.LoadDebugEnvIfRunsFromIDE(t, "workspace")
	template := `
resource "databricks_sql_endpoint" "this" {
	name = "permission-sql-{var.RANDOM}"
	cluster_size = "2X-Small"
	tags {
		custom_tags {
			key   = "Owner"
			value = "eng-dev-ecosystem-team_at_databricks.com"
		}
	}
}

resource "databricks_group" "sql_users" {
	display_name = "permission-sql-users-{var.RANDOM}"
}

resource "databricks_permission" "warehouse_users" {
	sql_endpoint_id  = databricks_sql_endpoint.this.id
	group_name       = databricks_group.sql_users.display_name
	permission_level = "CAN_USE"
}
`

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permission.warehouse_users", "permission_level", "CAN_USE"),
		),
	})
}

func TestAccPermission_InstancePool(t *testing.T) {
	acceptance.LoadDebugEnvIfRunsFromIDE(t, "workspace")
	template := `
data "databricks_node_type" "smallest" {
	local_disk = true
}

resource "databricks_instance_pool" "this" {
	instance_pool_name = "permission-pool-{var.RANDOM}"
	min_idle_instances = 0
	max_capacity = 1
	node_type_id = data.databricks_node_type.smallest.id
	idle_instance_autotermination_minutes = 10
}

resource "databricks_group" "pool_users" {
	display_name = "permission-pool-users-{var.RANDOM}"
}

resource "databricks_permission" "pool_access" {
	instance_pool_id = databricks_instance_pool.this.id
	group_name       = databricks_group.pool_users.display_name
	permission_level = "CAN_ATTACH_TO"
}
`

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permission.pool_access", "permission_level", "CAN_ATTACH_TO"),
		),
	})
}

func TestAccPermission_ClusterPolicy(t *testing.T) {
	acceptance.LoadDebugEnvIfRunsFromIDE(t, "workspace")
	template := `
resource "databricks_cluster_policy" "this" {
	name = "permission-policy-{var.RANDOM}"
	definition = jsonencode({
		"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
			"type": "fixed",
			"value": "jdbc:sqlserver://<jdbc-url>"
		}
	})
}

resource "databricks_group" "policy_users" {
	display_name = "permission-policy-users-{var.RANDOM}"
}

resource "databricks_permission" "policy_access" {
	cluster_policy_id = databricks_cluster_policy.this.id
	group_name        = databricks_group.policy_users.display_name
	permission_level  = "CAN_USE"
}
`

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_permission.policy_access", "permission_level", "CAN_USE"),
		),
	})
}
