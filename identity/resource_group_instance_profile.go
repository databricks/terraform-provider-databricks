package identity

import (
	"context"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroupInstanceProfile defines group role resource
func ResourceGroupInstanceProfile() *schema.Resource {
	return util.NewPairID("group_id", "instance_profile_id").Schema(func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		m["instance_profile_id"].ValidateDiagFunc = ValidInstanceProfile
		return m
	}).BindResource(util.BindResource{
		ReadContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			group, err := NewGroupsAPI(ctx, c).Read(groupID)
			if err == nil && !group.HasRole(roleARN) {
				return common.NotFound("Group has no instance profile")
			}
			return err
		},
		CreateContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).PatchR(groupID, scimPatchRequest("add", "roles", roleARN))
		},
		DeleteContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).PatchR(groupID, scimPatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN), ""))
		},
	})
}
