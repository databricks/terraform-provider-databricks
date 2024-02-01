package workspace

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceDirectory ...
func DataSourceDirectory() common.Resource {
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
		"workspace_path": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			notebooksAPI := NewNotebooksAPI(ctx, m)
			path := d.Get("path").(string)
			data, err := notebooksAPI.Read(path)
			if err != nil {
				return err
			}
			if data.ObjectType != Directory { // should we support Repos as well?
				return fmt.Errorf("'%s' isn't a directory", path)
			}
			d.SetId(data.Path)
			d.Set("object_id", data.ObjectID)
			d.Set("path", path)
			d.Set("workspace_path", "/Workspace"+data.Path)
			return nil
		},
	}
}
