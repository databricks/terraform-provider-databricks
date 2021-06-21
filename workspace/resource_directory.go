package workspace

import (
	"context"
	"fmt"

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
		"object_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"delete_recursive": {
			Type:     schema.TypeBool,
			Default:  false,
			Optional: true,
		},
	}

	directoryRead := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		notebooksAPI := NewNotebooksAPI(ctx, c)
		objectStatus, err := notebooksAPI.Read(d.Id())
		if err != nil {
			return err
		}
		if objectStatus.ObjectType != Directory {
			d.SetId("")
			return fmt.Errorf("different object type, %s, on this path other than a directory", objectStatus.ObjectType)
		}
		return common.StructToData(objectStatus, s, d)
	}

	return common.Resource{
		Schema: s,
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
		Read:   directoryRead,
		Update: directoryRead,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewNotebooksAPI(ctx, c).Delete(d.Id(), d.Get("delete_recursive").(bool))
		},
	}.ToResource()
}
