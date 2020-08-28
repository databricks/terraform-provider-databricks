package workspace

import (
	"hash/fnv"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceNotebookPaths() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNotebookPathsRead,
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
				Set: NotebookPathListHash,
			},
		},
	}
}

func dataSourceNotebookPathsRead(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	recursive := d.Get("recursive").(bool)

	client := m.(*common.DatabricksClient)

	notebookList, err := NewNotebooksAPI(client).List(path, recursive)
	if err != nil {
		return err
	}

	d.SetId(path)
	err = d.Set("recursive", recursive)
	if err != nil {
		return err
	}

	err = d.Set("path", path)
	if err != nil {
		return err
	}

	var notebookPathList []map[string]string
	for _, v := range notebookList {
		notebookPathMap := map[string]string{}
		notebookPathMap["path"] = v.Path
		notebookPathMap["language"] = string(v.Language)
		notebookPathList = append(notebookPathList, notebookPathMap)
	}

	err = d.Set("notebook_path_list", notebookPathList)

	return err
}

// NotebookPathListHash a hash
func NotebookPathListHash(v interface{}) int {
	h := fnv.New32a()
	m := v.(map[string]interface{})
	var err error
	if v, ok := m["path"]; ok {
		_, err = h.Write([]byte(v.(string)))
		if err != nil {
			return 0
		}
	}
	c := int(h.Sum32())
	if -c >= 0 {
		return -c
	}
	return c
}
