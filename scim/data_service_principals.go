package scim

import (
	"context"
	"fmt"
	"sort"

	"github.com/databricks/terraform-provider-databricks/common"
)

// DataSourceServicePrincipals searches for service principals based on display_name
func DataSourceServicePrincipals() common.Resource {
	type spnsData struct {
		DisplayNameContains string   `json:"display_name_contains,omitempty" tf:"computed"`
		ApplicationIDs      []string `json:"application_ids,omitempty" tf:"computed,slice_set"`
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
			return fmt.Errorf("cannot find SPs with display name containing %s", response.DisplayNameContains)
		}
		for _, sp := range spList {
			response.ApplicationIDs = append(response.ApplicationIDs, sp.ApplicationID)
		}
		sort.Strings(response.ApplicationIDs)
		return nil
	})
}
