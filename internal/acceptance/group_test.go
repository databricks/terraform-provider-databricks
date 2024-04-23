package acceptance

import (
	"context"

	"github.com/stretchr/testify/assert"

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
				group, err := groupsAPI.Read(id, "displayName,entitlements")
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
					group, err := groupsAPI.Read(id, "displayName,entitlements")
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

func TestMwsAccGroupsUpdateDisplayName(t *testing.T) {
	nameInit := qa.RandomName("tfgroup")
	nameUpdate := qa.RandomName("tfgroup")
	accountLevel(t, step{
		Template: `resource "databricks_group" "this" {
			display_name = "` + nameInit + `"
		}`,
		Check: resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_group.this", "display_name", nameInit),
			resourceCheck("databricks_group.this",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
					groupsAPI := scim.NewGroupsAPI(ctx, client)
					group, err := groupsAPI.Read(id, "displayName,entitlements")
					if err != nil {
						return err
					}
					assert.Equal(t, group.DisplayName, nameInit)
					return nil
				}),
		),
	}, step{
		Template: `resource "databricks_group" "this" {
			display_name = "` + nameUpdate + `"
		}`,
		Check: resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_group.this", "display_name", nameUpdate),
			resourceCheck("databricks_group.this",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
					groupsAPI := scim.NewGroupsAPI(ctx, client)
					group, err := groupsAPI.Read(id, "displayName,entitlements")
					if err != nil {
						return err
					}
					assert.Equal(t, group.DisplayName, nameUpdate)
					return nil
				}),
		),
	})
}
func TestAccGroupsUpdateDisplayName(t *testing.T) {
	nameInit := qa.RandomName("tfgroup")
	nameUpdate := qa.RandomName("tfgroup")
	workspaceLevel(t, step{
		Template: `resource "databricks_group" "this" {
			display_name = "` + nameInit + `"
		}`,
		Check: resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_group.this", "display_name", nameInit),
			resourceCheck("databricks_group.this",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
					groupsAPI := scim.NewGroupsAPI(ctx, client)
					group, err := groupsAPI.Read(id, "displayName,entitlements")
					if err != nil {
						return err
					}
					assert.Equal(t, group.DisplayName, nameInit)
					return nil
				}),
		),
	}, step{
		Template: `resource "databricks_group" "this" {
			display_name = "` + nameUpdate + `"
		}`,
		Check: resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_group.this", "display_name", nameUpdate),
			resourceCheck("databricks_group.this",
				func(ctx context.Context, client *common.DatabricksClient, id string) error {
					groupsAPI := scim.NewGroupsAPI(ctx, client)
					group, err := groupsAPI.Read(id, "displayName,entitlements")
					if err != nil {
						return err
					}
					assert.Equal(t, group.DisplayName, nameUpdate)
					return nil
				}),
		),
	})
}
