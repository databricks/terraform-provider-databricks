package scim

import (
	"context"
	"fmt"
	"sort"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceServicePrincipals searches for service principals based on display_name
func DataSourceServicePrincipals() *schema.Resource {
	type spnsData struct {
		DisplayNameContains string   `json:"display_name_contains,omitempty" tf:"computed"`
		ApplicationIDs      []string `json:"application_ids,omitempty" tf:"computed,slice_set"`
	}
	return common.DataResource(spnsData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		response := e.(*spnsData)
		spnAPI := NewServicePrincipalsAPI(ctx, c)

		spList, err := spnAPI.filter(fmt.Sprintf("displayName co '%s'", response.DisplayNameContains))
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
