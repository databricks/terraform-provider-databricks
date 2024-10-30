package acceptance

import (
	"fmt"
	"testing"
)

var (
	appTemplate = `
	resource "databricks_sql_endpoint" "this" {
		name = "tf-{var.RANDOM}"
		cluster_size = "2X-Small"
		max_num_clusters = 1

		tags {
			custom_tags {
				key   = "Owner"
				value = "eng-dev-ecosystem-team_at_databricks.com"
			}
		}
	}

	resource "databricks_directory" "my_custom_directory" {
 		path = "/Shared/provider-test/xx_{var.RANDOM}"
	}
	
	resource "databricks_workspace_file" "this" {
		source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
		path = "${databricks_directory.my_custom_directory.path}/tf-test-python.py"
	}
	resource "databricks_app" "this" {
		name = "{var.RANDOM}"
		description = "%s"
		source_code_path = databricks_directory.my_custom_directory.path
		mode = "SNAPSHOT"
		resource {
			name = "warehouse"
			description = "warehouse for app"
			sql_warehouse {
				id = databricks_sql_endpoint.this.id
				permission = "CAN_MANAGE"
			}		
		}
	}`
)

func TestAccAppCreate(t *testing.T) {
	loadWorkspaceEnv(t)
	if isGcp(t) {
		skipf(t)("not available on GCP")
	}
	WorkspaceLevel(t, Step{
		Template: fmt.Sprintf(appTemplate, "My app"),
	})
}

func TestAccAppUpdate(t *testing.T) {
	loadWorkspaceEnv(t)
	if isGcp(t) {
		skipf(t)("not available on GCP")
	}
	WorkspaceLevel(t, Step{
		Template: fmt.Sprintf(appTemplate, "My app"),
	}, Step{
		Template: fmt.Sprintf(appTemplate, "My new app"),
	})
}
