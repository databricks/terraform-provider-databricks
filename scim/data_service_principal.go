package scim

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// onlyOneNonEmpty checks if only one string in the slice is non-empty
func onlyOneNonEmpty(strings []string) bool {
	maxLen := 0
	sumLen := 0

	for _, s := range strings {
		l := len(s)
		if l > maxLen {
			maxLen = l
		}
		sumLen += l
	}
	if sumLen != maxLen {
		return false
	} else {
		return true
	}
}

// DataSourceServicePrincipal returns information about the spn specified by the application_id, id, or display_name
func DataSourceServicePrincipal() *schema.Resource {
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

		if !onlyOneNonEmpty([]string{response.ApplicationID, response.DisplayName, response.SpID}) {
			return fmt.Errorf("please specify only one of: application_id, sp_id, or display_name")
		}
		if response.ApplicationID != "" {
			spList, err = spnAPI.Filter(fmt.Sprintf(`applicationId eq "%s"`, response.ApplicationID), true)
		} else if response.SpID != "" {
			spList, err = spnAPI.Filter(fmt.Sprintf(`id eq "%s"`, response.SpID), true)
		} else if response.DisplayName != "" {
			spList, err = spnAPI.Filter(fmt.Sprintf(`displayName eq "%s"`, response.DisplayName), true)
		} else {
			return fmt.Errorf("please specify either application_id, display_name, or sp_id")
		}
		if err != nil {
			return err
		}
		if len(spList) == 0 {
			if response.ApplicationID != "" {
				return fmt.Errorf("cannot find SP with an application ID %s", response.ApplicationID)
			} else if response.SpID != "" {
				return fmt.Errorf("cannot find SP with an ID %s", response.SpID)
			} else {
				return fmt.Errorf("cannot find SP with name %s", response.DisplayName)
			}
		} else if len(spList) > 1 {
			return fmt.Errorf("there are more than %d Service Principals with name %s", len(spList), response.DisplayName)
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
