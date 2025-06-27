package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceVolumes() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		CatalogName string   `json:"catalog_name"`
		SchemaName  string   `json:"schema_name"`
		Ids         []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}, w *databricks.WorkspaceClient) error {
		volumes, err := w.Volumes.ListAll(ctx, catalog.ListVolumesRequest{CatalogName: data.CatalogName, SchemaName: data.SchemaName})
		if err != nil {
			return err
		}
		for _, v := range volumes {
			data.Ids = append(data.Ids, v.FullName)
		}
		return nil
	})
}
