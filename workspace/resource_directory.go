package workspace

import (
	"context"
	"errors"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceDirectory manages directories
func ResourceDirectory() *schema.Resource {
	s := map[string]*schema.Schema{
		"path": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"delete_recursive": {
			Type:     schema.TypeBool,
			Default:  false,
			Optional: true,
		},
	}

	return common.Resource{
		Schema:        s,
		SchemaVersion: 1,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			notebooksAPI := NewNotebooksAPI(ctx, c)
			path := d.Get("path").(string)
			if err := notebooksAPI.Mkdirs(path); err != nil {
				// TODO: handle RESOURCE_ALREADY_EXISTS
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
			if objectStatus.ObjectType != Directory {
				// TODO: better error message
				return errors.New("different object type on this path other than a directory")
			}
			return common.StructToData(objectStatus, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewNotebooksAPI(ctx, c).Delete(d.Id(), d.Get("delete_recursive").(bool))
		},
	}.ToResource()
}
