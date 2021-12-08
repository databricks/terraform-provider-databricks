package scim

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroupMember bind group with member
func ResourceGroupMember() *schema.Resource {
	return common.NewPairID("group_id", "member_id").BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Patch(groupID, PatchRequest("add", "members", memberID))
		},
		ReadContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			group, err := NewGroupsAPI(ctx, c).Read(groupID)
			hasMember := ComplexValues(group.Members).HasValue(memberID)
			if err == nil && !hasMember {
				return common.NotFound("Group has no member")
			}
			return err
		},
		DeleteContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Patch(groupID, PatchRequest(
				"remove", fmt.Sprintf(`members[value eq "%s"]`, memberID), ""))
		},
	})
}
