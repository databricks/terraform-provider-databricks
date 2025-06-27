package aws

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/scim"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroupInstanceProfile defines group role resource
func ResourceGroupInstanceProfile() common.Resource {
	r := common.NewPairID("group_id", "instance_profile_id").Schema(func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		m["instance_profile_id"].ValidateDiagFunc = ValidArn
		return m
	}).BindResource(common.BindResource{
		ReadContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			group, err := scim.NewGroupsAPI(ctx, c).Read(groupID, "roles")
			hasRole := scim.ComplexValues(group.Roles).HasValue(roleARN)
			if err == nil && !hasRole {
				return &apierr.APIError{
					ErrorCode:  "NOT_FOUND",
					StatusCode: 404,
					Message:    "Group has no instance profile",
				}
			}
			return err
		},
		CreateContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			return scim.NewGroupsAPI(ctx, c).Patch(groupID, scim.PatchRequestWithValue("add", "roles", roleARN))
		},
		DeleteContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			return scim.NewGroupsAPI(ctx, c).Patch(groupID, scim.PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN)))
		},
	})
	r.DeprecationMessage = "Please migrate to `databricks_group_role`"
	return r
}
