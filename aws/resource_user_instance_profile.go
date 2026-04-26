package aws

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/scim"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceUserInstanceProfile binds user and instance profile
func ResourceUserInstanceProfile() common.Resource {
	r := common.NewPairID("user_id", "instance_profile_id").Schema(func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		m["instance_profile_id"].ValidateDiagFunc = ValidArn
		common.AddApiField(m)
		common.AddNamespaceInSchema(m)
		common.NamespaceCustomizeSchemaMap(m)
		return m
	}).BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return scim.NewUsersAPI(ctx, c, common.GetApiLevel(d)).Patch(userID, scim.PatchRequestWithValue("add", "roles", roleARN))
		},
		ReadContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			user, err := scim.NewUsersAPI(ctx, c, common.GetApiLevel(d)).Read(userID, "roles")
			hasRole := scim.ComplexValues(user.Roles).HasValue(roleARN)
			if err == nil && !hasRole {
				return apierr.ErrNotFound
			}
			return err
		},
		DeleteContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return scim.NewUsersAPI(ctx, c, common.GetApiLevel(d)).Patch(userID, scim.PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN)))
		},
	})
	r.DeprecationMessage = "Please migrate to `databricks_user_role`. This resource will be removed in v0.5.x"
	r.CustomizeDiff = func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
		return common.CustomizeDiffDualResources(ctx, d, c)
	}
	return r
}
