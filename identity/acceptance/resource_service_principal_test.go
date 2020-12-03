package acceptance

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccServicePrincipalResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	config := qa.EnvironmentTemplate(t, `
        data "databricks_group" "admins" {
        	display_name = "admins"
        }
	resource "databricks_service_principal" "sp_first" {
		application_id = "00000000-1234-5678-0000-000000000001"
		display_name = "Eerste {var.RANDOM}"
	}
	resource "databricks_service_principal" "sp_second" {
		application_id = "00000000-1234-5678-0000-000000000002"
		display_name = "Tweede {var.RANDOM}"
        	allow_cluster_create = true
	}
	resource "databricks_service_principal" "sp_third" {
		application_id = "00000000-1234-5678-0000-000000000003"
        	allow_instance_pool_create = true
	}`)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_service_principal.sp_first", "allow_cluster_create", "false"),
					resource.TestCheckResourceAttr("databricks_service_principal.sp_first", "allow_instance_pool_create", "false"),
					resource.TestCheckResourceAttr("databricks_service_principal.sp_second", "allow_cluster_create", "true"),
					resource.TestCheckResourceAttr("databricks_service_principal.sp_second", "allow_instance_pool_create", "false"),
					resource.TestCheckResourceAttr("databricks_service_principal.sp_third", "allow_cluster_create", "false"),
					resource.TestCheckResourceAttr("databricks_service_principal.sp_third", "allow_instance_pool_create", "true"),
				),
				Destroy: false,
			},
		},
	})
}
