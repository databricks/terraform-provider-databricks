package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/stretchr/testify/assert"
)

func TestAccGroupMemberResource(t *testing.T) {
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_group" "root" {
				display_name = "tf-{var.RANDOM}"
			}
			resource "databricks_group" "first" {
				display_name = "tf-{var.RANDOM}-first"
			}
			resource "databricks_group" "second" {
				display_name = "tf-{var.RANDOM}-second"
			}
			resource "databricks_group_member" "rf" {
				group_id = databricks_group.root.id
				member_id = databricks_group.first.id
			}
			resource "databricks_group_member" "rs" {
				group_id = databricks_group.root.id
				member_id = databricks_group.second.id
			}`,
			Callback: func(ctx context.Context, client *common.DatabricksClient, id string) error {
				g, err := scim.NewGroupsAPI(ctx, client).Read(id)
				if err != nil {
					return err
				}
				assert.Len(t, g.Members, 2)
				return nil
			},
		},
	})
}
