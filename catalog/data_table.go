package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceTable() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		Id    string             `json:"id,omitempty" tf:"computed"`
		Name  string             `json:"name"`
		Table *catalog.TableInfo `json:"table_info,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		table, err := w.Tables.GetByFullName(ctx, data.Name)
		if err != nil {
			return err
		}
		data.Table = table
		data.Id = table.TableId
		return nil
	})
}
