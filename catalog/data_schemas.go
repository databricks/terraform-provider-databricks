package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceSchemas() common.Resource {
	return common.WorkspaceDataWithUnifiedProvider(func(ctx context.Context, data *struct {
		CatalogName string   `json:"catalog_name"`
		Ids         []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}, w *databricks.WorkspaceClient) error {
		schemas, err := w.Schemas.ListAll(ctx, catalog.ListSchemasRequest{CatalogName: data.CatalogName})
		if err != nil {
			return err
		}
		for _, v := range schemas {
			data.Ids = append(data.Ids, v.FullName)
		}
		return nil
	})
}
