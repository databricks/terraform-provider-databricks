package acceptance

import (
	"context"
	"fmt"
	"os"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/scim"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

// https://github.com/databrickslabs/terraform-provider-databricks/issues/1099
func TestAccGroupsExternalIdAndScimProvisioning(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	name := qa.RandomName("tfgroup")
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: `resource "databricks_group" "this" {
					display_name = "` + name + `"
					allow_cluster_create = true
				}`,
				Check: acceptance.ResourceCheck("databricks_group.this",
					func(ctx context.Context, client *common.DatabricksClient, id string) error {
						groupsAPI := scim.NewGroupsAPI(ctx, client)
						group, err := groupsAPI.Read(id)
						if err != nil {
							return err
						}
						// external SCIM change
						return groupsAPI.UpdateNameAndEntitlements(
							id, group.DisplayName, qa.RandomName("ext-id"), group.Entitlements)
					}),
			},
			{
				Config: `resource "databricks_group" "this" {
					display_name = "` + name + `"
					allow_cluster_create = true
				}`,
			},
		},
	})
}

func TestAccGroupResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	displayName := fmt.Sprintf("tf group test %s", randomStr)
	newDisplayName := fmt.Sprintf("new tf group test %s", randomStr)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: testDatabricksGroup(displayName),
				Check: resource.TestCheckResourceAttr(
					"databricks_group.my_group", "display_name", displayName),
				Destroy: false,
			},
			{
				Config:             testDatabricksGroup(newDisplayName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
		},
	})
}

func TestAccGroupResource_verify_entitlements(t *testing.T) {
	// TODO: CHECK THESE RESOURCES FOR GENERIC DESTROY
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	displayName := fmt.Sprintf("tf group test %s", randomStr)
	newDisplayName := fmt.Sprintf("new tf group test %s", randomStr)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config: testDatabricksGroupEntitlements(displayName, "true", "true"),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_group.my_group", "allow_cluster_create", "true"),
					resource.TestCheckResourceAttr("databricks_group.my_group", "allow_instance_pool_create", "true"),
				),
				Destroy: false,
			},
			{
				Config:             testDatabricksGroup(newDisplayName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
		},
	})
}

func testDatabricksGroup(groupName string) string {
	return fmt.Sprintf(`
		resource "databricks_group" "my_group" {
			display_name = "%s"
		}`, groupName)
}

func testDatabricksGroupEntitlements(groupName, allowClusterCreate, allowPoolCreate string) string {
	return fmt.Sprintf(`
		resource "databricks_group" "my_group" {
			display_name = "%s"
			allow_cluster_create = %s
			allow_instance_pool_create = %s
		}`, groupName, allowClusterCreate, allowPoolCreate)
}
