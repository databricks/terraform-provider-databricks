package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/unitycatalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSchemas() *schema.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		CatalogName string   `json:"catalog_name"`
		Ids         []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}, w *databricks.WorkspaceClient) error {
		schemas, err := w.Schemas.ListAll(ctx, unitycatalog.ListSchemasRequest{CatalogName: data.CatalogName})
		if err != nil {
			return err
		}
		for _, v := range schemas {
			data.Ids = append(data.Ids, v.FullName)
		}
		return nil
	})
}
