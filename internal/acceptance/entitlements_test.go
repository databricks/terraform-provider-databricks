package acceptance

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

type entitlement struct {
	name  string
	value bool
}

func (e entitlement) String() string {
	return fmt.Sprintf("%s = %t", e.name, e.value)
}

func entitlementsStepBuilder(t *testing.T, c **httpclient.ApiClient, groupName string) func(entitlements []entitlement) step {
	return func(entitlements []entitlement) step {
		entitlementsBuf := strings.Builder{}
		for _, entitlement := range entitlements {
			entitlementsBuf.WriteString(fmt.Sprintf("%s\n", entitlement.String()))
		}
		return step{
			Template: fmt.Sprintf(`
			data "databricks_group" "example" {
				display_name = "%s"
			}
			
			resource "databricks_entitlements" "entitlements_users" {
				group_id              = data.databricks_group.example.id
				%s
			}
		`, groupName, entitlementsBuf.String()),
			Check: func(s *terraform.State) error {
				groupId := s.RootModule().Resources["data.databricks_group.example"].Primary.ID
				var res iam.Group
				ctx := context.Background()
				err := (*c).Do(ctx, "GET", fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%s?attributes=entitlements", groupId),
					httpclient.WithResponseUnmarshal(&res))
				assert.NoError(t, err)
				receivedEntitlements := make([]string, 0, len(res.Entitlements))
				for _, entitlement := range res.Entitlements {
					receivedEntitlements = append(receivedEntitlements, entitlement.Value)
				}
				expectedEntitlements := make([]string, 0, len(entitlements))
				for _, entitlement := range entitlements {
					if entitlement.value {
						expectedEntitlements = append(expectedEntitlements, strings.ReplaceAll(entitlement.name, "_", "-"))
					}
				}
				assert.ElementsMatch(t, expectedEntitlements, receivedEntitlements)
				return nil
			},
		}

	}
}

func makeEntitlementsSteps(t *testing.T, entitlementsSteps [][]entitlement) []step {
	groupName := RandomName("entitlements-")
	var c *httpclient.ApiClient
	makeEntitlementsStep := entitlementsStepBuilder(t, &c, groupName)
	steps := make([]step, len(entitlementsSteps))
	for i, entitlements := range entitlementsSteps {
		steps[i] = makeEntitlementsStep(entitlements)
	}
	steps[0].PreConfig = makePreconfig(t, &c, groupName)
	return steps
}

func makePreconfig(t *testing.T, c **httpclient.ApiClient, groupName string) func() {
	return func() {
		w := databricks.Must(databricks.NewWorkspaceClient())
		var err error
		*c, err = w.Config.NewApiClient()
		assert.NoError(t, err)
		ctx := context.Background()
		group, err := w.Groups.Create(ctx, iam.Group{
			DisplayName: groupName,
		})
		assert.NoError(t, err)
		t.Cleanup(func() {
			err := w.Groups.DeleteById(ctx, group.Id)
			assert.NoError(t, err)
		})
	}
}

func TestAccEntitlementsAddToEmpty(t *testing.T) {
	steps := makeEntitlementsSteps(t, [][]entitlement{
		{},
		{
			{"allow_cluster_create", true},
			{"allow_instance_pool_create", true},
			{"workspace_access", true},
			{"databricks_sql_access", true},
		},
	})
	workspaceLevel(t, steps...)
}

func TestAccEntitlementsSetExplicitlyToFalse(t *testing.T) {
	steps := makeEntitlementsSteps(t, [][]entitlement{
		{
			{"allow_cluster_create", false},
			{"allow_instance_pool_create", false},
			{"workspace_access", false},
			{"databricks_sql_access", false},
		},
		{},
		{
			{"allow_cluster_create", false},
			{"allow_instance_pool_create", false},
			{"workspace_access", false},
			{"databricks_sql_access", false},
		},
	})
	workspaceLevel(t, steps...)
}

func TestAccEntitlementsRemoveExisting(t *testing.T) {
	steps := makeEntitlementsSteps(t, [][]entitlement{
		{
			{"allow_cluster_create", true},
			{"allow_instance_pool_create", true},
			{"workspace_access", true},
			{"databricks_sql_access", true},
		},
		{},
	})
	workspaceLevel(t, steps...)
}

func TestAccEntitlementsSomeTrueSomeFalse(t *testing.T) {
	steps := makeEntitlementsSteps(t, [][]entitlement{
		{
			{"allow_cluster_create", false},
			{"allow_instance_pool_create", false},
			{"workspace_access", true},
			{"databricks_sql_access", true},
		},
	})
	workspaceLevel(t, steps...)
}
