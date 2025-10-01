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
	Id   string `json:"id" tf:"computed,optional"`
	Name string `json:"name" tf:"computed,optional"`
	common.ProviderConfig
}

func DataSourceWarehouse() common.Resource {
	return common.WorkspaceDataWithParams(func(ctx context.Context, data sqlWarehouseDataParams, w *databricks.WorkspaceClient) (*SqlWarehouse, error) {
		if data.Id == "" && data.Name == "" {
			return nil, fmt.Errorf("either 'id' or 'name' should be provided")
		}
		selected := []sql.DataSource{}
		dataSources, err := w.DataSources.List(ctx)
		if err != nil {
			return nil, err
		}
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
		warehouse, err := getSqlWarehouse(ctx, w, selected[0].WarehouseId)
		if err != nil {
			return nil, err
		}
		warehouse.DataSourceId = selected[0].Id
		return warehouse, nil
	})
}
