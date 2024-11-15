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

	resource "databricks_app" "this" {
		name = "{var.RANDOM}"
		description = "%s"
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
