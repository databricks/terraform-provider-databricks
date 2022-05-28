package aws

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/scim"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceServicePrincipalInstanceProfile binds service principal and instance profile
func ResourceServicePrincipalRole() *schema.Resource {
	r := common.NewPairID("service_principal_id", "role").Schema(func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		m["instance_profile_id"].ValidateDiagFunc = ValidInstanceProfile
		return m
	}).BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, servicePrincipalID, role string, c *common.DatabricksClient) error {
			return scim.NewServicePrincipalsAPI(ctx, c).Patch(servicePrincipalID, scim.PatchRequest("add", "roles", role))
		},
		ReadContext: func(ctx context.Context, servicePrincipalID, roleARN string, c *common.DatabricksClient) error {
			user, err := scim.NewServicePrincipalsAPI(ctx, c).Read(servicePrincipalID)
			hasRole := scim.ComplexValues(user.Roles).HasValue(roleARN)
			if err == nil && !hasRole {
				return common.NotFound("Service Principal has no role")
			}
			return err
		},
		DeleteContext: func(ctx context.Context, userID, roleARN string, c *common.DatabricksClient) error {
			return scim.NewServicePrincipalsAPI(ctx, c).Patch(userID, scim.PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN), ""))
		},
	})
	return r
}
