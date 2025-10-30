package scim

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/databricks/terraform-provider-databricks/common"
)

// DataSourceServicePrincipals searches for service principals based on display_name
func DataSourceServicePrincipals() common.Resource {
	type spnsData struct {
		DisplayNameContains string                 `json:"display_name_contains,omitempty" tf:"computed"`
		ApplicationIDs      []string               `json:"application_ids,omitempty" tf:"computed,slice_set"`
		ServicePrincipals   []servicePrincipalData `json:"service_principals,omitempty" tf:"computed"`
	}
	return common.DataResource(spnsData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		response := e.(*spnsData)
		spnAPI := NewServicePrincipalsAPI(ctx, c)

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
		return nil
	})
}
