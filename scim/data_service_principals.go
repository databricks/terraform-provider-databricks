package scim

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceServicePrincipals searches for service principals based on display_name
func DataSourceServicePrincipals() common.Resource {
	type spnsData struct {
		DisplayNameContains string                 `json:"display_name_contains,omitempty" tf:"computed"`
		ApplicationIDs      []string               `json:"application_ids,omitempty" tf:"computed,slice_set"`
		ServicePrincipals   []servicePrincipalData `json:"service_principals,omitempty" tf:"computed"`
	}
	s := common.StructToSchema(spnsData{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.AddApiField(m)
		return m
	})
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)

	return common.Resource{
		IsDual: true,
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			if err := common.ValidateApiLevelForUnifiedHostFromData(d, m); err != nil {
				return err
			}
			newClient, err := m.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			var response spnsData
			common.DataToStructPointer(d, s, &response)
			spnAPI := NewServicePrincipalsAPI(ctx, newClient, common.GetApiLevel(d))

			var filter string
			if response.DisplayNameContains != "" {
				filter = fmt.Sprintf(`displayName co "%s"`, response.DisplayNameContains)
			}
			spList, err := spnAPI.Filter(filter, true)
			if err != nil {
				return err
			}
			if len(spList) == 0 {
				log.Printf("[INFO] cannot find SPs with display name containing %s", response.DisplayNameContains)
			}
			for _, sp := range spList {
				response.ApplicationIDs = append(response.ApplicationIDs, sp.ApplicationID)
				response.ServicePrincipals = append(response.ServicePrincipals, servicePrincipalData{
					ApplicationID:  sp.ApplicationID,
					DisplayName:    sp.DisplayName,
					Active:         sp.Active,
					ID:             sp.ID,
					SpID:           sp.ID,
					ScimID:         sp.ID,
					ExternalID:     sp.ExternalID,
					AclPrincipalID: fmt.Sprintf("servicePrincipals/%s", sp.ApplicationID),
					Home:           fmt.Sprintf("/Users/%s", sp.ApplicationID),
					Repos:          fmt.Sprintf("/Repos/%s", sp.ApplicationID),
				})
			}
			sort.Strings(response.ApplicationIDs)
			sort.Slice(response.ServicePrincipals, func(i, j int) bool {
				return response.ServicePrincipals[i].ApplicationID < response.ServicePrincipals[j].ApplicationID
			})
			if err := common.StructToData(&response, s, d); err != nil {
				return err
			}
			d.SetId("_")
			return nil
		},
	}
}
