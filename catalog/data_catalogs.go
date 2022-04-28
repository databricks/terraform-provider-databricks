package catalog

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCatalogs() *schema.Resource {
	type catalogsData struct {
		Ids []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.DataResource(catalogsData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*catalogsData)
		catalogsAPI := NewCatalogsAPI(ctx, c)
		catalogs, err := catalogsAPI.list()
		if err != nil {
			return err
		}
		for _, v := range catalogs.Catalogs {
			data.Ids = append(data.Ids, v.Name)
		}
		return nil
	})
}
