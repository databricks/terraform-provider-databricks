package workspace

import (
	"bytes"
	"context"
	"log"
	"path/filepath"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func workspaceFileUploadOptionFunc(i *workspace.Import) {
	i.Format = "RAW"
	i.Overwrite = true
	i.ForceSendFields = []string{"Content"}
}

// ResourceWorkspaceFile manages files in workspace
func ResourceWorkspaceFile() common.Resource {
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
		"workspace_path": {
			Type:     schema.TypeString,
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
			err = client.Workspace.Upload(ctx, path, bytes.NewReader(content), workspaceFileUploadOptionFunc)
			if err != nil {
				if isParentDoesntExistError(err) {
					parent := filepath.ToSlash(filepath.Dir(path))
					log.Printf("[DEBUG] Parent folder '%s' doesn't exist, creating...", parent)
					err = client.Workspace.MkdirsByPath(ctx, parent)
					if err != nil {
						return err
					}
					err = client.Workspace.Upload(ctx, path, bytes.NewReader(content), workspaceFileUploadOptionFunc)
				}
				if err != nil {
					return err
				}
			}
			d.SetId(path)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			client, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			objectStatus, err := common.RetryOnTimeout(ctx, func(ctx context.Context) (*workspace.ObjectInfo, error) {
				return client.Workspace.GetStatusByPath(ctx, d.Id())
			})
			if err != nil {
				return err
			}
			SetWorkspaceObjectComputedProperties(d, c)
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
			return client.Workspace.Upload(ctx, d.Id(), bytes.NewReader(content), workspaceFileUploadOptionFunc)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			client, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return client.Workspace.Delete(ctx, workspace.Delete{Path: d.Id(), Recursive: false})
		},
	}
}
