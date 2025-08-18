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
		return m
	}).BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			return scim.NewUsersAPI(ctx, c).Patch(userID, scim.PatchRequestWithValue("add", "roles", roleARN))
		},
		ReadContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			user, err := scim.NewUsersAPI(ctx, c).Read(userID, "roles")
			hasRole := scim.ComplexValues(user.Roles).HasValue(roleARN)
			if err == nil && !hasRole {
				return apierr.ErrNotFound
			}
			return err
		},
		DeleteContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			return scim.NewUsersAPI(ctx, c).Patch(userID, scim.PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN)))
		},
	})
	r.DeprecationMessage = "Please migrate to `databricks_user_role`. This resource will be removed in v0.5.x"
	return r
}
