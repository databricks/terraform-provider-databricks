package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceCatalogs() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		Ids []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}, w *databricks.WorkspaceClient) error {
		var req catalog.ListCatalogsRequest = catalog.ListCatalogsRequest{}
		catalogs, err := w.Catalogs.ListAll(ctx, req)
		if err != nil {
			return err
		}
		for _, v := range catalogs {
			data.Ids = append(data.Ids, v.Name)
		}
		return nil
	})
}
