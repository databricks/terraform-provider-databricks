package scim

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroupRole bind group with role
func ResourceGroupRole() common.Resource {
	r := common.NewPairID("group_id", "role").Schema(func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.AddApiField(m)
		common.AddNamespaceInSchema(m)
		common.NamespaceCustomizeSchemaMap(m)
		return m
	}).BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, groupID, role string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			return NewGroupsAPI(ctx, c, common.GetApiLevel(d)).Patch(groupID, PatchRequestWithValue("add", "roles", role))
		},
		ReadContext: func(ctx context.Context, groupID, role string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			group, err := NewGroupsAPI(ctx, c, common.GetApiLevel(d)).Read(groupID, "roles")
			hasRole := ComplexValues(group.Roles).HasValue(role)
			if err == nil && !hasRole {
				return &apierr.APIError{
					ErrorCode:  "NOT_FOUND",
					StatusCode: 404,
					Message:    "Group has no role",
				}
			}
			return err
		},
		DeleteContext: func(ctx context.Context, groupID, role string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			return NewGroupsAPI(ctx, c, common.GetApiLevel(d)).Patch(groupID, PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, role)))
		},
	})
	r.CustomizeDiff = func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
		return common.CustomizeDiffDualResources(ctx, d, c)
	}
	r.IsDual = true
	return r
}
