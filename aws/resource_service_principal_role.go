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
func ResourceServicePrincipalRole() *schema.Resource {
	r := common.NewPairID("service_principal_id", "role").BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, servicePrincipalID, role string, c *common.DatabricksClient) error {
			return scim.NewServicePrincipalsAPI(ctx, c).Patch(servicePrincipalID, scim.PatchRequest("add", "roles", role))
		},
		ReadContext: func(ctx context.Context, servicePrincipalID, roleARN string, c *common.DatabricksClient) error {
			servicePrincipal, err := scim.NewServicePrincipalsAPI(ctx, c).Read(servicePrincipalID)
			hasRole := scim.ComplexValues(servicePrincipal.Roles).HasValue(roleARN)
			if err == nil && !hasRole {
				return apierr.NotFound("Service Principal has no role")
			}
			return err
		},
		DeleteContext: func(ctx context.Context, servicePrincipalID, roleARN string, c *common.DatabricksClient) error {
			return scim.NewServicePrincipalsAPI(ctx, c).Patch(servicePrincipalID, scim.PatchRequest(
				"remove", fmt.Sprintf(`roles[value eq "%s"]`, roleARN), ""))
		},
	})
	return r
}
