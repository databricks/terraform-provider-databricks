package storage

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceDbfsFilePaths() common.Resource {
	return common.Resource{
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			path := d.Get("path").(string)
			recursive := d.Get("recursive").(bool)
			paths, err := NewDbfsAPI(ctx, m).List(path, recursive)
			if err != nil {
				return err
			}
			d.SetId(path)
			pathList := []map[string]any{}
			for _, pathInfo := range paths {
				pathData := map[string]any{}
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
