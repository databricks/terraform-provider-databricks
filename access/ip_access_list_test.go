package access_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccIPACLListsResourceFullLifecycle(t *testing.T) {
	acceptance.WorkspaceLevel(t, Step{
		Template: `
		resource "databricks_ip_access_list" "this" {
			label = "tf-{var.RANDOM}"
			list_type = "BLOCK"
			ip_addresses = [
				"10.0.10.25",
				"10.0.10.0/24"
			]
		}`,
	}, Step{
		Template: `
		resource "databricks_ip_access_list" "this" {
			label = "tf-{var.RANDOM}"
			list_type = "BLOCK"
			ip_addresses = [
				"10.0.11.25",
				"10.0.11.0/24"
			]
		}`,
	})
}
