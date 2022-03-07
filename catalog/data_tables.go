package catalog

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceTables() *schema.Resource {
	var data struct {
		CatalogName string   `json:"catalog_name"`
		SchemaName  string   `json:"schema_name"`
		Ids         []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.DataResource(&data, func(ctx context.Context, c *common.DatabricksClient) error {
		tablesAPI := NewTablesAPI(ctx, c)
		tables, err := tablesAPI.listTables(data.CatalogName, data.SchemaName)
		if err != nil {
			return err
		}
		for _, v := range tables.Tables {
			data.Ids = append(data.Ids, v.FullName())
		}
		return nil
	})
}
