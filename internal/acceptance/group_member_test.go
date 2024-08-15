package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/stretchr/testify/assert"
)

const groupMemberTest = `
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
}`

func TestMwsAccGroupMemberResource(t *testing.T) {
	accountLevel(t, LegacyStep{
		Template: groupMemberTest,
		Check: resourceCheck("databricks_group.root",
			func(ctx context.Context, client *common.DatabricksClient, id string) error {
				g, err := scim.NewGroupsAPI(ctx, client).Read(id, "members")
				if err != nil {
					return err
				}
				assert.Len(t, g.Members, 2)
				return nil
			}),
	})
}

func TestAccGroupMemberResource(t *testing.T) {
	workspaceLevel(t, LegacyStep{
		Template: groupMemberTest,
		Check: resourceCheck("databricks_group.root",
			func(ctx context.Context, client *common.DatabricksClient, id string) error {
				g, err := scim.NewGroupsAPI(ctx, client).Read(id, "members")
				if err != nil {
					return err
				}
				assert.Len(t, g.Members, 2)
				return nil
			}),
	})
}
