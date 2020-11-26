package identity

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroupInstanceProfile defines group role resource
func ResourceGroupInstanceProfile() *schema.Resource {
	return util.NewPairID("group_id", "instance_profile_id").BindResource(util.BindResource{
		ReadContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			group, err := NewGroupsAPI(ctx, c).Read(groupID)
			if err == nil && !group.HasRole(roleARN) {
				return common.NotFound("Group has no instance profile")
			}
			return err
		},
		CreateContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			err := validateInstanceProfileARN(roleARN)
			if err != nil {
				return err
			}
			return NewGroupsAPI(ctx, c).PatchR(groupID, scimPatchRequest("add", "roles", roleARN))
		},
		DeleteContext: func(ctx context.Context, groupID, roleARN string, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).PatchR(groupID, scimPatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN), ""))
		},
	})
}

func validateInstanceProfileARN(v string) error {
	instanceProfileArn, err := arn.Parse(v)
	if err != nil {
		return fmt.Errorf("Illegal instance profile %s: %s", v, err)
	}
	if !strings.HasPrefix(instanceProfileArn.Resource, "instance-profile") {
		return fmt.Errorf("Not an instance profile ARN: %s", v)
	}
	return nil
}
