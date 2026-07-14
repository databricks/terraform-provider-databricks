package storage

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceDbfsFile manages files on DBFS
func ResourceDbfsFile() common.Resource {
	s := workspace.FileContentSchema(map[string]*schema.Schema{
		"file_size": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"dbfs_path": {
			Type:     schema.TypeString,
			Computed: true,
		},
	})
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)
	return common.Resource{
		SchemaVersion: 1,
		Schema:        s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			path := d.Get("path").(string)
			content, err := workspace.ReadContent(d)
			if err != nil {
				return err
			}
			if err = NewDbfsAPI(ctx, newClient).Create(path, content, true); err != nil {
				return err
			}
			d.SetId(path)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			dbfsAPI := NewDbfsAPI(ctx, newClient)
			fileInfo, err := dbfsAPI.Status(d.Id())
			if err != nil {
				return err
			}
			d.Set("path", fileInfo.Path)
			d.Set("dbfs_path", fmt.Sprint("dbfs:", fileInfo.Path))
			d.Set("file_size", fileInfo.FileSize)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return NewDbfsAPI(ctx, newClient).Delete(d.Id(), false)
		},
		StateUpgraders: []schema.StateUpgrader{
			{
				Version: 0,
				Type:    DbfsFileV0(),
				Upgrade: workspace.MigrateV0,
			},
		},
	}
}
