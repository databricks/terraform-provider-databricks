package acceptance

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/scim"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestMwsAccGroupsExternalIdAndScimProvisioning(t *testing.T) {
	name := qa.RandomName("tfgroup")
	accountLevel(t, step{
		Template: `resource "databricks_group" "this" {
			display_name = "` + name + `"
		}`,
		Check: resourceCheck("databricks_group.this",
			func(ctx context.Context, client *common.DatabricksClient, id string) error {
				// duplicate code between workspace level and account level, because clients
				// might get different
				groupsAPI := scim.NewGroupsAPI(ctx, client)
				group, err := groupsAPI.Read(id)
				if err != nil {
					return err
				}
				// external SCIM change
				return groupsAPI.UpdateNameAndEntitlements(
					id, group.DisplayName, qa.RandomName("ext-id"), group.Entitlements)
			}),
	}, step{
		Template: `resource "databricks_group" "this" {
			display_name = "` + name + `"
		}`,
	})
}

// https://github.com/databricks/terraform-provider-databricks/issues/1099
func TestAccGroupsExternalIdAndScimProvisioning(t *testing.T) {
	name := qa.RandomName("tfgroup")
	workspaceLevel(t, step{
		Template: `resource "databricks_group" "this" {
			display_name = "` + name + `"
			allow_cluster_create = true
		}`,
		Check: resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_group.this", "allow_cluster_create", "true"),
			resource.TestCheckResourceAttr("databricks_group.this", "allow_instance_pool_create", "false"),
			resourceCheck("databricks_group.this",
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
		),
	}, step{
		Template: `resource "databricks_group" "this" {
			display_name = "` + name + `"
			allow_cluster_create = true
		}`,
	})
}
