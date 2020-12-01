package workspace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceNotebookPaths ...
func DataSourceNotebookPaths() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			path := d.Get("path").(string)
			recursive := d.Get("recursive").(bool)
			notebookList, err := NewNotebooksAPI(ctx, m).List(path, recursive)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(path)
			if err = d.Set("recursive", recursive); err != nil {
				return diag.FromErr(err)
			}
			if err = d.Set("path", path); err != nil {
				return diag.FromErr(err)
			}
			var notebookPathList []map[string]string
			for _, v := range notebookList {
				notebookPathMap := map[string]string{}
				notebookPathMap["path"] = v.Path
				notebookPathMap["language"] = string(v.Language)
				notebookPathList = append(notebookPathList, notebookPathMap)
			}
			// nolint
			d.Set("notebook_path_list", notebookPathList)
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
			"notebook_path_list": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"language": {
							Type:     schema.TypeString,
							Optional: true,
						},
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
