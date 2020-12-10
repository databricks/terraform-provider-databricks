package acceptance

import (
	"context"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccGroupMemberResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	config := qa.EnvironmentTemplate(t, `
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
	}`)
	var rootID, subID *string
	assertMembers := func(num int) func(
		context.Context, *common.DatabricksClient, string) error {
		return func(ctx context.Context, client *common.DatabricksClient, id string) error {
			g, err := NewGroupsAPI(ctx, client).Read(id)
			if err != nil {
				return err
			}
			assert.Len(t, g.Members, num)
			return nil
		}
	}
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config:  config,
				Destroy: false,
				Check: acceptance.ResourceCheck("databricks_group",
					func(ctx context.Context, client *common.DatabricksClient, id string) error {
						*rootID = id
						return nil
					}),
			},
			{
				PreConfig: func() {
					//	manually create another subgroup
					ctx := context.Background()
					client := common.CommonEnvironmentClient()
					subGroupC, _ := NewGroupsAPI(ctx, client).Create(qa.RandomName("manual-tf-"), nil, nil, nil)
					*subID = subGroupC.ID
					//  Add new subgroup to root group
					err := NewGroupsAPI(ctx, client).Patch(*rootID, []string{*subID}, nil, GroupMembersPath)
					assert.NoError(t, err, err)
				},
				Config: config,
				Check:  acceptance.ResourceCheck("databricks_group", assertMembers(3)),
			},
			{
				Config: `resource "databricks_group" "root" {
					display_name = "group-membership-step3"
				}`,
				Check: acceptance.ResourceCheck("databricks_group", assertMembers(1)),
			},
		},
	})
}
