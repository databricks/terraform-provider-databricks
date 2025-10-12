package sql

import (
	"context"
	"sort"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceWarehouses() common.Resource {
	type warehousesData struct {
		WarehouseNameContains string   `json:"warehouse_name_contains,omitempty"`
		Ids                   []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.WorkspaceDataWithUnifiedProvider(func(ctx context.Context, data *warehousesData, w *databricks.WorkspaceClient) error {
		list, err := w.Warehouses.ListAll(ctx, sql.ListWarehousesRequest{})
		if err != nil {
			return err
		}
		nameContains := strings.ToLower(data.WarehouseNameContains)
		for _, warehouse := range list {
			if nameContains != "" && !strings.Contains(strings.ToLower(warehouse.Name), nameContains) {
				continue
			}
			data.Ids = append(data.Ids, warehouse.Id)
		}

		sort.Strings(data.Ids)
		return nil
	})
}
