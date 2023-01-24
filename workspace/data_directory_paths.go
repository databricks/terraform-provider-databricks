package workspace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceDirectoryPaths ...
func DataSourceDirectoryPaths() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
			path := d.Get("path").(string)
			recursive := d.Get("recursive").(bool)
			directoryList, err := NewNotebooksAPI(ctx, m).ListDirectories(path, recursive)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(path)
			var directoryPathList []map[string]string
			for _, v := range directoryList {
				directoryPathMap := map[string]string{}
				directoryPathMap["path"] = v.Path
				directoryPathList = append(directoryPathList, directoryPathMap)
			}
			// nolint
			d.Set("directory_path_list", directoryPathList)
			return nil
		},
		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"recursive": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"directory_path_list": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				Set: PathListHash,
			},
		},
	}
}
