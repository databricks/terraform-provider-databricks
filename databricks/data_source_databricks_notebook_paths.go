package databricks

import (
	"bytes"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNotebookPaths() *schema.Resource {
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
				Set: notebookPathListHash,
			},
		},
	}
}

func dataSourceNotebookPathsRead(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	recursive := d.Get("recursive").(bool)

	client := m.(*service.DBApiClient)

	notebookList, err := client.Notebooks().List(path, recursive)
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

func notebookPathListHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["path"]; ok {
		buf.WriteString(v.(string))
	}
	return hashcode.String(buf.String())
}
