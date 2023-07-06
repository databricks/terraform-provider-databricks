package workspace

import (
	"context"
	"encoding/base64"
	"path/filepath"

	ws_api "github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceWorkspaceFile manages files in workspace
func ResourceWorkspaceFile() *schema.Resource {
	s := FileContentSchema(map[string]*schema.Schema{
		"url": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"object_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
	})
	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			client, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			path := d.Get("path").(string)
			parent := filepath.ToSlash(filepath.Dir(path))
			if parent != "/" {
				err = client.Workspace.MkdirsByPath(ctx, parent)
				if err != nil {
					return err
				}
			}
			err = client.Workspace.Import(ctx, ws_api.Import{
				Content:   base64.StdEncoding.EncodeToString(content),
				Format:    ws_api.ImportFormatAuto,
				Path:      path,
				Overwrite: true,
			})
			if err != nil {
				return err
			}
			d.SetId(path)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			client, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			objectStatus, err := client.Workspace.GetStatusByPath(ctx, d.Id())
			if err != nil {
				return err
			}
			d.Set("url", c.FormatURL("#workspace", d.Id()))
			return common.StructToData(objectStatus, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			client, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			return client.Workspace.Import(ctx, ws_api.Import{
				Content:   base64.StdEncoding.EncodeToString(content),
				Format:    ws_api.ImportFormatAuto,
				Overwrite: true,
				Path:      d.Id(),
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			client, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return client.Workspace.Delete(ctx, ws_api.Delete{Path: d.Id(), Recursive: true})
		},
	}.ToResource()
}
