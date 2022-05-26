package catalog

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func updateFunctionFactory(securable string, updatable []string) func(context.Context, *schema.ResourceData, *common.DatabricksClient) error {
	return func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		patch := map[string]interface{}{}
		for _, field := range updatable {
			old, new := d.GetChange(field)
			if !d.HasChange(field) {
				continue
			}
			if field == "name" && old == "" {
				continue
			}
			patch[field] = new
		}
		if len(patch) == 0 {
			return nil
		}
		switch securable {
		case "metastore":
			return NewMetastoresAPI(ctx, c).updateMetastore(d.Id(), patch)
		case "catalog":
			return NewCatalogsAPI(ctx, c).updateCatalog(d.Id(), patch)
		case "schema":
			return NewSchemasAPI(ctx, c).updateSchema(d.Id(), patch)
		case "table":
			return NewTablesAPI(ctx, c).updateTable(d.Id(), patch)
		default:
			return nil
		}
	}
}
