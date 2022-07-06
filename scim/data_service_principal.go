package scim

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceServicePrincipal returns information about the spn specified by the application_id
func DataSourceServicePrincipal() *schema.Resource {
	type spnData struct {
		ApplicationID string `json:"application_id,omitempty" tf:"computed"`
		DisplayName   string `json:"display_name,omitempty" tf:"computed"`
		SpID          string `json:"sp_id,omitempty" tf:"computed"`
		Home          string `json:"home,omitempty" tf:"computed"`
		Repos         string `json:"repos,omitempty" tf:"computed"`
		Active        bool   `json:"active,omitempty" tf:"computed"`
		ExternalID    string `json:"external_id,omitempty" tf:"computed"`
	}
	return common.DataResource(spnData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		response := e.(*spnData)
		spnAPI := NewServicePrincipalsAPI(ctx, c)
		spList, err := spnAPI.filter(fmt.Sprintf("applicationId eq '%s'", response.ApplicationID))
		if err != nil {
			return err
		}
		if len(spList) == 0 {
			return fmt.Errorf("cannot find SP with ID %s", response.ApplicationID)
		}
		sp := spList[0]
		response.DisplayName = sp.DisplayName
		response.Home = fmt.Sprintf("/Users/%s", sp.ApplicationID)
		response.Repos = fmt.Sprintf("/Repos/%s", sp.ApplicationID)
		response.ExternalID = sp.ExternalID
		response.Active = sp.Active
		response.SpID = sp.ID
		return nil
	})
}
