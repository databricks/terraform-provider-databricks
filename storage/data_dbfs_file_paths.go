package storage

import (
	"context"

	"github.com/databrickslabs/databricks-terraform/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceDBFSFilePaths ...
func DataSourceDBFSFilePaths() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			path := d.Get("path").(string)
			recursive := d.Get("recursive").(bool)
			paths, err := NewDbfsAPI(ctx, m).List(path, recursive)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(path)
			pathList := []map[string]interface{}{}
			for _, pathInfo := range paths {
				pathData := map[string]interface{}{}
				pathData["path"] = pathInfo.Path
				pathData["file_size"] = pathInfo.FileSize
				pathList = append(pathList, pathData)
			}
			// nolint
			d.Set("path_list", pathList)
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
			"path_list": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"file_size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
				Set: workspace.PathListHash,
			},
		},
	}
}
