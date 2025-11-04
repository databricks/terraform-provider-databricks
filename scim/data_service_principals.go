package scim

import (
	"context"
	"fmt"
	"sort"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceServicePrincipals searches for service principals based on display_name
func DataSourceServicePrincipals() common.Resource {
	type spnsData struct {
		DisplayNameContains string   `json:"display_name_contains,omitempty" tf:"computed"`
		ApplicationIDs      []string `json:"application_ids,omitempty" tf:"computed,slice_set"`
	}
	s := common.StructToSchema(spnsData{}, nil)
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)
	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var response spnsData
			common.DataToStructPointer(d, s, &response)
			spnAPI := NewServicePrincipalsAPI(ctx, newClient)

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
			err = common.StructToData(response, s, d)
			if err != nil {
				return err
			}
			d.SetId("_")
			return nil
		},
	}
}
