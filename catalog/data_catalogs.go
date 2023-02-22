package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCatalogs() *schema.Resource {
	type catalogsData struct {
		Ids []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.WorkspaceData(func(ctx context.Context, data *catalogsData, w *databricks.WorkspaceClient) error {
		catalogs, err := w.Catalogs.ListAll(ctx)
		if err != nil {
			return err
		}
		for _, v := range catalogs {
			data.Ids = append(data.Ids, v.Name)
		}
		return nil
	})
}
