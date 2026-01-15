package sql

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
)

// Note that these fields are both marked as computed/optional because users can specify either the name or the ID
// of the warehouse to retrieve.
type sqlWarehouseDataParams struct {
	common.Namespace
	Id   string `json:"id" tf:"computed,optional"`
	Name string `json:"name" tf:"computed,optional"`
}

func DataSourceWarehouse() common.Resource {
	return common.WorkspaceDataWithParams(func(ctx context.Context, data sqlWarehouseDataParams, w *databricks.WorkspaceClient) (*SqlWarehouse, error) {
		if data.Id == "" && data.Name == "" {
			return nil, fmt.Errorf("either 'id' or 'name' should be provided")
		}

		var warehouseId string
		var dataSourceId string

		dataSources, err := w.DataSources.List(ctx)
		if err == nil {
			selected := []sql.DataSource{}
			for _, source := range dataSources {
				if data.Name != "" && source.Name == data.Name {
					selected = append(selected, source)
				} else if data.Id != "" && source.WarehouseId == data.Id {
					selected = append(selected, source)
					break
				}
			}
			if len(selected) == 0 {
				if data.Name != "" {
					return nil, fmt.Errorf("can't find SQL warehouse with the name '%s'", data.Name)
				} else {
					return nil, fmt.Errorf("can't find SQL warehouse with the ID '%s'", data.Id)
				}
			}
			if len(selected) > 1 {
				if data.Name != "" {
					return nil, fmt.Errorf("there are multiple SQL warehouses with the name '%s'", data.Name)
				} else {
					return nil, fmt.Errorf("there are multiple SQL warehouses with the ID '%s'", data.Id)
				}
			}
			warehouseId = selected[0].WarehouseId
			dataSourceId = selected[0].Id
		} else {
			// DataSources.List failed; fall back to finding warehouse without DataSourceId
			if data.Name != "" {
				// Find warehouse by name using Warehouses.ListAll
				list, err := w.Warehouses.ListAll(ctx, sql.ListWarehousesRequest{})
				if err != nil {
					return nil, err
				}
				var matched []sql.EndpointInfo
				for _, wh := range list {
					if wh.Name == data.Name {
						matched = append(matched, wh)
					}
				}
				if len(matched) == 0 {
					return nil, fmt.Errorf("can't find SQL warehouse with the name '%s'", data.Name)
				}
				if len(matched) > 1 {
					return nil, fmt.Errorf("there are multiple SQL warehouses with the name '%s'", data.Name)
				}
				warehouseId = matched[0].Id
			} else {
				warehouseId = data.Id
			}
		}

		warehouse, err := getSqlWarehouse(ctx, w, warehouseId)
		if err != nil {
			return nil, err
		}
		warehouse.DataSourceId = dataSourceId
		return warehouse, nil
	})
}
