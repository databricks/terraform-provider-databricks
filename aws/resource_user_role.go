package aws

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceUserRole() common.Resource {
	r := common.NewPairID("user_id", "role").Schema(func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.AddApiField(m)
		common.AddNamespaceInSchema(m)
		common.NamespaceCustomizeSchemaMap(m)
		return m
	}).BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, userID, role string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			return scim.NewUsersAPI(ctx, c, common.GetApiLevel(d)).Patch(userID, scim.PatchRequestWithValue("add", "roles", role))
		},
		ReadContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			user, err := scim.NewUsersAPI(ctx, c, common.GetApiLevel(d)).Read(userID, "roles")
			hasRole := scim.ComplexValues(user.Roles).HasValue(roleARN)
			if err == nil && !hasRole {
				return &apierr.APIError{
					ErrorCode:  "NOT_FOUND",
					StatusCode: 404,
					Message:    "User has no role",
				}
			}
			return err
		},
		DeleteContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			return scim.NewUsersAPI(ctx, c, common.GetApiLevel(d)).Patch(userID, scim.PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN)))
		},
	})
	r.CustomizeDiff = func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
		return common.CustomizeDiffDualResources(ctx, d, c)
	}
	r.IsDual = true
	return r
}
