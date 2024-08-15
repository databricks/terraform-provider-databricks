package acceptance

import (
	"testing"
)

func TestAccIPACLListsResourceFullLifecycle(t *testing.T) {
	workspaceLevel(t, LegacyStep{
		Template: `
		resource "databricks_ip_access_list" "this" {
			label = "tf-{var.RANDOM}"
			list_type = "BLOCK"
			ip_addresses = [
				"10.0.10.25",
				"10.0.10.0/24"
			]
		}`,
	}, LegacyStep{
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
