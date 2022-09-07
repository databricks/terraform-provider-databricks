package acceptance

import (
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccEntitlementResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	config := acceptance.EnvironmentTemplate(t, `
	resource "databricks_user" "first" {
		user_name = "tf-eerste+{var.RANDOM}@example.com"
		display_name = "Eerste {var.RANDOM}"
	}

	resource "databricks_group" "second" {
		display_name = "{var.RANDOM} group"
	}
	
	resource "databricks_entitlements" "first_entitlements" {
		user_id                    = databricks_user.first.id
		allow_cluster_create       = true
		allow_instance_pool_create = true
	}	

	resource "databricks_entitlements" "second_entitlements" {
		group_id                   = databricks_group.second.id
		allow_cluster_create       = true
		allow_instance_pool_create = true
	}		
	`)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_entitlements.first_entitlements", "allow_cluster_create", "true"),
					resource.TestCheckResourceAttr("databricks_entitlements.first_entitlements", "allow_instance_pool_create", "true"),
					resource.TestCheckResourceAttr("databricks_entitlements.second_entitlements", "allow_cluster_create", "true"),
					resource.TestCheckResourceAttr("databricks_entitlements.second_entitlements", "allow_instance_pool_create", "true"),
				),
			},
			{
				Config: config,
			},
		},
	})
}

func TestAccServicePrincipalEntitlementsResourceOnAzure(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloud != "azure" {
		t.Skip("Test is only for CLOUD_ENV=azure")
	}
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_service_principal" "this" {
				application_id = "00000000-1234-5678-0000-000000000001"
				display_name = "SPN {var.RANDOM}"
			}

			resource "databricks_entitlements" "service_principal" {
				service_principal_id       = databricks_service_principal.this.id
				allow_cluster_create       = true
				allow_instance_pool_create = true
			}`,
		},
	})
}

func TestAccServicePrincipalEntitlementsResourceOnAws(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloud != "aws" {
		t.Skip("Test is only for CLOUD_ENV=aws")
	}
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_service_principal" "this" {
				display_name = "SPN {var.RANDOM}"
			}

			resource "databricks_entitlements" "service_principal" {
				service_principal_id       = databricks_service_principal.this.id
				allow_cluster_create       = true
				allow_instance_pool_create = true
			}`,
		},
	})
}
