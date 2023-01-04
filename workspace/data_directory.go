package workspace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceDirectory ...
func DataSourceDirectory() *schema.Resource {
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
	}
	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
			notebooksAPI := NewNotebooksAPI(ctx, m)
			path := d.Get("path").(string)
			data, err := notebooksAPI.Read(path)
			if err != nil {
				return diag.FromErr(err)
			}
			if data.ObjectType != Directory { // should we support Repos as well?
				return diag.Errorf("'%s' isn't a directory", path)
			}
			d.SetId(data.Path)
			d.Set("object_id", data.ObjectID)
			d.Set("path", path)
			return nil
		},
	}
}
