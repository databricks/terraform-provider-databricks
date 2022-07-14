package sql

import (
	"context"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceWarehouses() *schema.Resource {
	type warehousesData struct {
		WarehouseNameContains string   `json:"warehouse_name_contains,omitempty"`
		Ids                   []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.DataResource(warehousesData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*warehousesData)
		a := NewSQLEndpointsAPI(ctx, c)
		list, err := a.List()
		if err != nil {
			return err
		}
		name_contains := data.WarehouseNameContains
		for _, v := range list.Endpoints {
			match_name := strings.Contains(strings.ToLower(v.Name), name_contains)
			if name_contains != "" && !match_name {
				continue
			}
			data.Ids = append(data.Ids, v.ID)
		}
		return nil
	})
}
