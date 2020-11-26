package identity

import (
	"context"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroupMember bind group with member
func ResourceGroupMember() *schema.Resource {
	return util.NewPairID("group_id", "member_id").BindResource(util.BindResource{
		CreateContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).PatchR(groupID, scimPatchRequest("add", "members", memberID))
		},
		ReadContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			group, err := NewGroupsAPI(ctx, c).Read(groupID)
			if err == nil && !group.HasMember(memberID) {
				return common.NotFound("Group has no member")
			}
			return err
		},
		DeleteContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).PatchR(groupID, scimPatchRequest(
				"remove", fmt.Sprintf(`members[value eq "%s"]`, memberID), ""))
		},
	})
}
