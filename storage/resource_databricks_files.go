package storage

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFiles() *schema.Resource {
	s := common.StructToSchema(any, func(m map[string]*schema.Schema) map[string]*schema.Schema {})
	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			w.Files.Upload(ctx, "")
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			w.Files.Download(ctx, "")
			return nil
		},
		Schema: s,
	}.ToResource()
}
