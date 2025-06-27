package scim

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
)

// ResourceGroupMember bind group with member
func ResourceGroupMember() common.Resource {
	return common.NewPairID("group_id", "member_id").BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Patch(groupID, PatchRequestWithValue("add", "members", memberID))
		},
		ReadContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			group, err := NewGroupsAPI(ctx, c).Read(groupID, "members")
			hasMember := ComplexValues(group.Members).HasValue(memberID)
			if err == nil && !hasMember {
				return &apierr.APIError{
					ErrorCode:  "NOT_FOUND",
					StatusCode: 404,
					Message:    "Group has no member",
				}
			}
			return err
		},
		DeleteContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Patch(groupID, PatchRequest(
				"remove", fmt.Sprintf(`members[value eq "%s"]`, memberID)))
		},
	})
}
