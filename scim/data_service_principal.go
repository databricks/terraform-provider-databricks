package scim

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceServicePrincipal returns information about the spn specified by the application_id, id, display_name, or scim_id
func DataSourceServicePrincipal() common.Resource {
	type spnData struct {
		ApplicationID  string `json:"application_id,omitempty" tf:"computed"`
		DisplayName    string `json:"display_name,omitempty" tf:"computed"`
		SpID           string `json:"sp_id,omitempty" tf:"computed"`
		ScimID         string `json:"scim_id,omitempty" tf:"computed"`
		ID             string `json:"id,omitempty" tf:"computed"`
		Home           string `json:"home,omitempty" tf:"computed"`
		Repos          string `json:"repos,omitempty" tf:"computed"`
		Active         bool   `json:"active,omitempty" tf:"computed"`
		ExternalID     string `json:"external_id,omitempty" tf:"computed"`
		AclPrincipalID string `json:"acl_principal_id,omitempty" tf:"computed"`
	}

	s := common.StructToSchema(spnData{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		s["application_id"].ExactlyOneOf = []string{"application_id", "display_name", "scim_id"}
		s["display_name"].ExactlyOneOf = []string{"application_id", "display_name", "scim_id"}
		s["scim_id"].ExactlyOneOf = []string{"application_id", "display_name", "scim_id"}
		return s
	})

	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			var response spnData
			var spList []User
			var err error

			common.DataToStructPointer(d, s, &response)
			spnAPI := NewServicePrincipalsAPI(ctx, m)

			if response.ApplicationID != "" {
				spList, err = spnAPI.Filter(fmt.Sprintf(`applicationId eq "%s"`, response.ApplicationID), true)
			} else if response.ScimID != "" {
				spList, err = spnAPI.Filter(fmt.Sprintf(`id eq "%s"`, response.ScimID), true)
			} else if response.DisplayName != "" {
				spList, err = spnAPI.Filter(fmt.Sprintf(`displayName eq "%s"`, response.DisplayName), true)
			} else {
				return fmt.Errorf("please specify either application_id, display_name, or scim_id")
			}
			if err != nil {
				return err
			}

			if len(spList) == 0 {
				if response.ApplicationID != "" {
					return fmt.Errorf("cannot find SP with an application ID %s", response.ApplicationID)
				} else if response.ScimID != "" {
					return fmt.Errorf("cannot find SP with an ID %s", response.ScimID)
				} else {
					return fmt.Errorf("cannot find SP with name %s", response.DisplayName)
				}
			} else if len(spList) > 1 {
				return fmt.Errorf("there are %d Service Principals with name %s", len(spList), response.DisplayName)
			}

			sp := spList[0]
			response.DisplayName = sp.DisplayName
			response.ApplicationID = sp.ApplicationID
			response.Home = fmt.Sprintf("/Users/%s", sp.ApplicationID)
			response.Repos = fmt.Sprintf("/Repos/%s", sp.ApplicationID)
			response.AclPrincipalID = fmt.Sprintf("servicePrincipals/%s", sp.ApplicationID)
			response.ExternalID = sp.ExternalID
			response.Active = sp.Active
			response.ScimID = sp.ID
			response.SpID = sp.ID

			err = common.StructToData(response, s, d)
			d.SetId(sp.ID)
			return err
		},
	}
}
