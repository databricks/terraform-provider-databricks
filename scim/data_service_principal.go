package scim

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
)

// DataSourceServicePrincipal returns information about the spn specified by the application_id
func DataSourceServicePrincipal() common.Resource {
	type spnData struct {
		ApplicationID  string `json:"application_id,omitempty" tf:"computed"`
		DisplayName    string `json:"display_name,omitempty" tf:"computed"`
		SpID           string `json:"sp_id,omitempty" tf:"computed"`
		ID             string `json:"id,omitempty" tf:"computed"`
		Home           string `json:"home,omitempty" tf:"computed"`
		Repos          string `json:"repos,omitempty" tf:"computed"`
		Active         bool   `json:"active,omitempty" tf:"computed"`
		ExternalID     string `json:"external_id,omitempty" tf:"computed"`
		AclPrincipalID string `json:"acl_principal_id,omitempty" tf:"computed"`
	}
	return common.DataResource(spnData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		response := e.(*spnData)
		spnAPI := NewServicePrincipalsAPI(ctx, c)
		var spList []User
		var err error
		if response.ApplicationID != "" && response.DisplayName != "" {
			return fmt.Errorf("please specify only one of application_id or display_name")
		}
		if response.ApplicationID != "" {
			spList, err = spnAPI.Filter(fmt.Sprintf(`applicationId eq "%s"`, response.ApplicationID), true)
		} else if response.DisplayName != "" {
			spList, err = spnAPI.Filter(fmt.Sprintf(`displayName eq "%s"`, response.DisplayName), true)
		} else {
			return fmt.Errorf("please specify either application_id or display_name")
		}
		if err != nil {
			return err
		}
		if len(spList) == 0 {
			if response.ApplicationID != "" {
				return fmt.Errorf("cannot find SP with ID %s", response.ApplicationID)
			} else {
				return fmt.Errorf("cannot find SP with name %s", response.DisplayName)
			}
		} else if len(spList) > 1 {
			return fmt.Errorf("there are more than 1 service principal with name %s", response.DisplayName)
		}

		sp := spList[0]
		response.DisplayName = sp.DisplayName
		response.ApplicationID = sp.ApplicationID
		response.Home = fmt.Sprintf("/Users/%s", sp.ApplicationID)
		response.Repos = fmt.Sprintf("/Repos/%s", sp.ApplicationID)
		response.AclPrincipalID = fmt.Sprintf("servicePrincipals/%s", sp.ApplicationID)
		response.ExternalID = sp.ExternalID
		response.Active = sp.Active
		response.SpID = sp.ID
		response.ID = sp.ID
		return nil
	})
}
