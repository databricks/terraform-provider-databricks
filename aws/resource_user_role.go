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
	return common.NewPairID("user_id", "role").Schema(func(m map[string]*schema.Schema) map[string]*schema.Schema {
		return common.AddApiField(m)
	}).BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, userID, role string, c *common.DatabricksClient, d *schema.ResourceData) error {
			return scim.NewUsersAPI(ctx, c, common.GetApiLevel(d)).Patch(userID, scim.PatchRequestWithValue("add", "roles", role))
		},
		ReadContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient, d *schema.ResourceData) error {
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
			return scim.NewUsersAPI(ctx, c, common.GetApiLevel(d)).Patch(userID, scim.PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN)))
		},
	})
}
