package acceptance

import (
	"context"
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
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
	// need to wait for internal state to stabilise, otherwise will get 409 Resource Conflict
	waitForState := func(ctx context.Context, client *common.DatabricksClient, id string) error {
		// sleep for 3 seconds
		time.Sleep(3 * time.Second)
		return nil
	}
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
		Check: resourceCheck("databricks_ip_access_list.this", waitForState),
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
		Check: resourceCheck("databricks_ip_access_list.this", waitForState),
	})
}
