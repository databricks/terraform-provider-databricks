package catalog

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSchemas() *schema.Resource {
	type schemasData struct {
		CatalogName string   `json:"catalog_name"`
		Ids         []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.DataResource(schemasData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*schemasData)
		schemasAPI := NewSchemasAPI(ctx, c)
		schemas, err := schemasAPI.listByCatalog(data.CatalogName)
		if err != nil {
			return err
		}
		for _, v := range schemas.Schemas {
			data.Ids = append(data.Ids, v.FullName)
		}
		return nil
	})
}
