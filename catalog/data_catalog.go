package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceCatalog() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		Id      string               `json:"id,omitempty" tf:"computed"`
		Name    string               `json:"name"`
		Catalog *catalog.CatalogInfo `json:"catalog_info,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient,
	) error {
		catalog, err := w.Catalogs.GetByName(ctx, data.Name)
		if err != nil {
			return err
		}
		data.Catalog = catalog
		data.Id = catalog.Name
		return nil
	})
}
