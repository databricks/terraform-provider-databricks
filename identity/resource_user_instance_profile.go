package identity

import (
	"context"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceUserInstanceProfile binds user and instance profile
func ResourceUserInstanceProfile() *schema.Resource {
	return util.NewPairID("user_id", "instance_profile_id").Schema(func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		m["instance_profile_id"].ValidateDiagFunc = ValidInstanceProfile
		return m
	}).BindResource(util.BindResource{
		CreateContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			return NewUsersAPI(ctx, c).PatchR(userID, scimPatchRequest("add", "roles", roleARN))
		},
		ReadContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			user, err := NewUsersAPI(ctx, c).Read(userID)
			if err == nil && !user.HasRole(roleARN) {
				return common.NotFound("User has no role")
			}
			return err
		},
		DeleteContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			return NewUsersAPI(ctx, c).PatchR(userID, scimPatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN), ""))
		},
	})
}
