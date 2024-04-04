package acceptance

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccEntitlementResource(t *testing.T) {
	var conf = `
	resource "databricks_user" "first" {
		user_name = "tf-eerste+{var.RANDOM}@example.com"
		display_name = "Eerste {var.RANDOM}"
		allow_cluster_create       = true
		allow_instance_pool_create = true		
	}

	resource "databricks_group" "second" {
		display_name = "{var.RANDOM} group"
		allow_cluster_create       = true
		allow_instance_pool_create = true		
	}

	resource "databricks_group" "third" {
		display_name = "{var.RANDOM} group 2"
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
	
	resource "databricks_entitlements" "third_entitlements" {
		group_id                   = databricks_group.third.id
		allow_cluster_create       = false
		allow_instance_pool_create = false
		databricks_sql_access      = false
		workspace_access           = false
	}`
	workspaceLevel(t, step{
		Template: conf,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_entitlements.first_entitlements", "allow_cluster_create", "true"),
			resource.TestCheckResourceAttr("databricks_entitlements.first_entitlements", "allow_instance_pool_create", "true"),
			resource.TestCheckResourceAttr("databricks_entitlements.second_entitlements", "allow_cluster_create", "true"),
			resource.TestCheckResourceAttr("databricks_entitlements.second_entitlements", "allow_instance_pool_create", "true"),
		),
	}, step{
		Template: conf,
	})
}

func TestAccServicePrincipalEntitlementsResourceOnAzure(t *testing.T) {
	// this test should run only on Azure, so just expect SPN config to be there or fail
	// TODO: change to SDK so that we can be explicit if we want to fetch entitlements
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	workspaceLevel(t, step{
		Template: `resource "databricks_service_principal" "this" {
			application_id = "{var.RANDOM_UUID}"
			allow_cluster_create       = true
			allow_instance_pool_create = true
			display_name = "SPN {var.RANDOM}"
			force = true			
		}

		resource "databricks_entitlements" "service_principal" {
			service_principal_id       = databricks_service_principal.this.id
			allow_cluster_create       = true
			allow_instance_pool_create = true
		}`,
	})
}

func TestAccServicePrincipalEntitlementsResourceOnAws(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	workspaceLevel(t, step{
		Template: `
		resource "databricks_service_principal" "this" {
			display_name = "SPN {var.RANDOM}"
			allow_cluster_create       = true
			allow_instance_pool_create = true				
		}

		resource "databricks_entitlements" "service_principal" {
			service_principal_id       = databricks_service_principal.this.id
			allow_cluster_create       = true
			allow_instance_pool_create = true
		}`,
	})
}

func TestAccIssue1860(t *testing.T) {
	groupName := RandomName("entitlements-")
	workspaceLevel(t, step{
		PreConfig: func() {
			w := databricks.Must(databricks.NewWorkspaceClient())
			ctx := context.Background()
			_, err := w.Groups.Create(ctx, iam.Group{
				DisplayName: groupName,
			})
			assert.NoError(t, err)
		},
		Template: fmt.Sprintf(`
		  data "databricks_group" "example" {
			display_name = "%s"
		  }
		  
		  resource "databricks_entitlements" "entitlements_users" {
			group_id              = data.databricks_group.example.id
		  }
		`, groupName),
	}, step{
		Template: fmt.Sprintf(`
		  data "databricks_group" "example" {
			display_name = "%s"
		  }
		  resource "databricks_entitlements" "entitlements_users" {
			group_id                   = data.databricks_group.example.id
			allow_cluster_create       = true
			allow_instance_pool_create = true
			workspace_access           = true
			databricks_sql_access      = true
		  }
		  `, groupName),
	}, step{
		Template: fmt.Sprintf(`
		  data "databricks_group" "example" {
			display_name = "%s"
		  }
		  resource "databricks_entitlements" "entitlements_users" {
			group_id                   = data.databricks_group.example.id
			allow_cluster_create       = false
			allow_instance_pool_create = false
			workspace_access           = true
			databricks_sql_access      = true
		  }
		  `, groupName),
	})
}
