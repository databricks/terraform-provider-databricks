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
	return common.WorkspaceData(func(ctx context.Context, data *warehousesData, w *databricks.WorkspaceClient) error {
		list, err := w.Warehouses.ListAll(ctx, sql.ListWarehousesRequest{})
		if err != nil {
			return err
		}
		name_contains := (*data).WarehouseNameContains
		for _, e := range list {
			match_name := strings.Contains(strings.ToLower(e.Name), name_contains)
			if name_contains != "" && !match_name {
				continue
			}
			data.Ids = append(data.Ids, e.Id)
		}

		sort.Strings(data.Ids)
		return nil
	})
}
