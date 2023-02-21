package acceptance

import (
	"testing"
)

func TestIPACLListsResourceFullLifecycle(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_workspace_conf" "features" {
			enable_ip_access_lists = "true"
		}
		  
		resource "databricks_ip_access_list" "this" {
			label = "tf-{var.RANDOM}"
			list_type = "BLOCK"
			ip_addresses = [
				"10.0.10.25",
				"10.0.10.0/24"
			]
			depends_on = [databricks_workspace_conf.features]
		}`,
	})
}
