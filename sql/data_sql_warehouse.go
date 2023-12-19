package sql

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SqlWarehouseDataParams struct {
	Id   string `json:"id" tf:"computed,optional"`
	Name string `json:"name" tf:"computed,optional"`
}

func DataSourceWarehouse() *schema.Resource {
	return common.WorkspaceData2[SqlWarehouse, SqlWarehouseDataParams](func(ctx context.Context, data SqlWarehouseDataParams, warehouse *SqlWarehouse, w *databricks.WorkspaceClient) error {
		if data.Id == "" && data.Name == "" {
			return fmt.Errorf("either 'id' or 'name' should be provided")
		}
		selected := []sql.DataSource{}
		dataSources, err := w.DataSources.List(ctx)
		if err != nil {
			return err
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
				return fmt.Errorf("can't find SQL warehouse with the name '%s'", data.Name)
			} else {
				return fmt.Errorf("can't find SQL warehouse with the ID '%s'", data.Id)
			}
		}
		if len(selected) > 1 {
			if data.Name != "" {
				return fmt.Errorf("there are multiple SQL warehouses with the name '%s'", data.Name)
			} else {
				return fmt.Errorf("there are multiple SQL warehouses with the ID '%s'", data.Id)
			}
		}
		*warehouse, err = getSqlWarehouse(ctx, w, selected[0].WarehouseId)
		warehouse.DataSourceId = selected[0].Id
		if err != nil {
			return err
		}
		return nil
	})
}
