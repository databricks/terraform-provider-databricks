package workspace

import (
	"context"
	"encoding/base64"
	"path/filepath"

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
			notebooksAPI := NewNotebooksAPI(ctx, c)
			path := d.Get("path").(string)
			parent := filepath.ToSlash(filepath.Dir(path))
			if parent != "/" {
				err = notebooksAPI.Mkdirs(parent)
				if err != nil {
					return err
				}
			}
			createNotebook := ImportPath{
				Content:   base64.StdEncoding.EncodeToString(content),
				Format:    Auto,
				Path:      path,
				Overwrite: true,
			}
			err = notebooksAPI.Create(createNotebook)
			if err != nil {
				return err
			}
			d.SetId(path)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			notebooksAPI := NewNotebooksAPI(ctx, c)
			objectStatus, err := notebooksAPI.Read(d.Id())
			if err != nil {
				return err
			}
			d.Set("url", c.FormatURL("#workspace", d.Id()))
			return common.StructToData(objectStatus, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			notebooksAPI := NewNotebooksAPI(ctx, c)
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			return notebooksAPI.Create(ImportPath{
				Content:   base64.StdEncoding.EncodeToString(content),
				Format:    Auto,
				Overwrite: true,
				Path:      d.Id(),
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewNotebooksAPI(ctx, c).Delete(d.Id(), true)
		},
	}.ToResource()
}
