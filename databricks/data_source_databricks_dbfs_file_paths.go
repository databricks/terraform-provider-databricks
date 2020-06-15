package databricks

import (
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceDBFSFilePaths() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDBFSFilePathsRead,
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
				Set: notebookPathListHash,
			},
		},
	}
}

func dataSourceDBFSFilePathsRead(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	recursive := d.Get("recursive").(bool)
	client := m.(*service.DBApiClient)

	paths, err := client.DBFS().List(path, recursive)
	if err != nil {
		return err
	}

	d.SetId(path)
	pathList := []map[string]interface{}{}
	for _, pathInfo := range paths {
		pathData := map[string]interface{}{}
		pathData["path"] = pathInfo.Path
		pathData["file_size"] = pathInfo.FileSize
		pathList = append(pathList, pathData)
	}

	err = d.Set("path_list", pathList)

	return err
}
