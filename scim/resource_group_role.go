package scim

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroupRole bind group with role
func ResourceGroupRole() *schema.Resource {
	return common.NewPairID("group_id", "role").BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, groupID, role string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Patch(groupID, PatchRequest("add", "roles", role))
		},
		ReadContext: func(ctx context.Context, groupID, role string, c *common.DatabricksClient) error {
			group, err := NewGroupsAPI(ctx, c).Read(groupID)
			hasRole := ComplexValues(group.Roles).HasValue(role)
			if err == nil && !hasRole {
				return apierr.NotFound("Group has no role")
			}
			return err
		},
		DeleteContext: func(ctx context.Context, groupID, role string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Patch(groupID, PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, role), ""))
		},
	})
}
