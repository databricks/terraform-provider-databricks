package aws

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/scim"
)

func ResourceUserRole() common.Resource {
	return common.NewPairID("user_id", "role").BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, userID, role string, c *common.DatabricksClient) error {
			return scim.NewUsersAPI(ctx, c).Patch(userID, scim.PatchRequestWithValue("add", "roles", role))
		},
		ReadContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			user, err := scim.NewUsersAPI(ctx, c).Read(userID, "roles")
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
		DeleteContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			return scim.NewUsersAPI(ctx, c).Patch(userID, scim.PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN)))
		},
	})
}
