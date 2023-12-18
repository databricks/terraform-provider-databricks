package acceptance

import (
	"testing"
	"time"
)

func TestAccIPACLListsResourceFullLifecycle(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_ip_access_list" "this" {
			label = "tf-{var.RANDOM}"
			list_type = "BLOCK"
			ip_addresses = [
				"10.0.10.25",
				"10.0.10.0/24"
			]
		}`,
	}, step{
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

func TestMwsAccIPACLListsResourceFullLifecycle(t *testing.T) {
	accountLevel(t, step{
		Template: `
		resource "databricks_ip_access_list" "this" {
			label = "tf-{var.RANDOM}"
			list_type = "BLOCK"
			ip_addresses = [
				"10.0.10.25",
				"10.0.10.0/24"
			]
		}`,
	}, step{
		PreConfig: func() { time.Sleep(1 * time.Second) },
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
