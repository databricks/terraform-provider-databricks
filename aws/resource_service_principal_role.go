package aws

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceServicePrincipalRole binds service principal and instance profile
func ResourceServicePrincipalRole() common.Resource {
	r := common.NewPairID("service_principal_id", "role").Schema(func(m map[string]*schema.Schema) map[string]*schema.Schema {
		return common.AddApiField(m)
	}).BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, servicePrincipalID, role string, c *common.DatabricksClient, d *schema.ResourceData) error {
			return scim.NewServicePrincipalsAPI(ctx, c, common.GetApiLevel(d)).Patch(servicePrincipalID, scim.PatchRequestWithValue("add", "roles", role))
		},
		ReadContext: func(ctx context.Context, servicePrincipalID, roleARN string, c *common.DatabricksClient, d *schema.ResourceData) error {
			servicePrincipal, err := scim.NewServicePrincipalsAPI(ctx, c, common.GetApiLevel(d)).Read(servicePrincipalID, "roles")
			hasRole := scim.ComplexValues(servicePrincipal.Roles).HasValue(roleARN)
			if err == nil && !hasRole {
				return &apierr.APIError{
					ErrorCode:  "NOT_FOUND",
					StatusCode: 404,
					Message:    "Service Principal has no role",
				}
			}
			return err
		},
		DeleteContext: func(ctx context.Context, servicePrincipalID, roleARN string, c *common.DatabricksClient, d *schema.ResourceData) error {
			return scim.NewServicePrincipalsAPI(ctx, c, common.GetApiLevel(d)).Patch(servicePrincipalID, scim.PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN)))
		},
	})
	return r
}
