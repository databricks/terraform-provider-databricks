package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceTables() *schema.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		CatalogName string   `json:"catalog_name"`
		SchemaName  string   `json:"schema_name"`
		Ids         []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}, w *databricks.WorkspaceClient) error {
		tables, err := w.Tables.ListAll(ctx, catalog.ListTablesRequest{CatalogName: data.CatalogName, SchemaName: data.SchemaName})
		if err != nil {
			return err
		}
		for _, v := range tables {
			if v.TableType != "VIEW" {
				data.Ids = append(data.Ids, v.FullName)
			}
		}
		return nil
	})
}
