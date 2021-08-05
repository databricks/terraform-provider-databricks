package acceptance

import (
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUserResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	config := qa.EnvironmentTemplate(t, `
	resource "databricks_user" "first" {
		user_name = "eerste+{var.RANDOM}@example.com"
		display_name = "Eerste {var.RANDOM}"
	}

	resource "databricks_user" "second" {
		user_name = "tweede+{var.RANDOM}@example.com"
		display_name = "Tweede {var.RANDOM}"
		allow_cluster_create = true
	}

	resource "databricks_user" "third" {
		user_name = "derde+{var.RANDOM}@example.com"
		display_name = "Derde {var.RANDOM}"
		allow_instance_pool_create = true
	}`)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_user.first", "allow_cluster_create", "false"),
					resource.TestCheckResourceAttr("databricks_user.first", "allow_instance_pool_create", "false"),
					resource.TestCheckResourceAttr("databricks_user.second", "allow_cluster_create", "true"),
					resource.TestCheckResourceAttr("databricks_user.second", "allow_instance_pool_create", "false"),
					resource.TestCheckResourceAttr("databricks_user.third", "allow_cluster_create", "false"),
					resource.TestCheckResourceAttr("databricks_user.third", "allow_instance_pool_create", "true"),
				),
			},
			{
				Config: config,
			},
		},
	})
}
