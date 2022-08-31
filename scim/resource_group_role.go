package scim

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroupRole bind group with role
func ResourceGroupRole() *schema.Resource {
	return common.NewPairID("group_id", "role_arn").BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Patch(groupID, PatchRequest("add", "roles", roleARN))
		},
		ReadContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			group, err := NewGroupsAPI(ctx, c).Read(groupID)
			hasRole := ComplexValues(group.Roles).HasValue(roleARN)
			if err == nil && !hasRole {
				return common.NotFound("Group has no roleARN")
			}
			return err
		},
		DeleteContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Patch(groupID, PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN), ""))
		},
	})
}
