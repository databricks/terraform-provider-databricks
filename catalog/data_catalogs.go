package catalog

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCatalogs() *schema.Resource {
	var data struct {
		Ids []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.DataResource(&data, func(ctx context.Context, c *common.DatabricksClient) error {
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
