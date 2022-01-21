package aws

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/scim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceUserRole() *schema.Resource {
	return common.NewPairID("user_id", "role").BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, userID, role string, c *common.DatabricksClient) error {
			return scim.NewUsersAPI(ctx, c).Patch(userID, scim.PatchRequest("add", "roles", role))
		},
		ReadContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			user, err := scim.NewUsersAPI(ctx, c).Read(userID)
			hasRole := scim.ComplexValues(user.Roles).HasValue(roleARN)
			if err == nil && !hasRole {
				return common.NotFound("User has no role")
			}
			return err
		},
		DeleteContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			return scim.NewUsersAPI(ctx, c).Patch(userID, scim.PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN), ""))
		},
	})
}