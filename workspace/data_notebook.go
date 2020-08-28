package workspace

import (
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceNotebook() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNotebookRead,
		Schema: map[string]*schema.Schema{

			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"format": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(DBC),
					string(Source),
					string(HTML),
				}, false),
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"language": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceNotebookRead(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	format := d.Get("format").(string)
	client := m.(*common.DatabricksClient)

	notebookInfo, err := NewNotebooksAPI(client).Read(path)
	if err != nil {
		return err
	}
	notebookContent, err := NewNotebooksAPI(client).Export(path, ExportFormat(format))
	if err != nil {
		return err
	}

	d.SetId(path)
	err = d.Set("path", path)
	if err != nil {
		return err
	}
	err = d.Set("format", format)
	if err != nil {
		return err
	}

	err = d.Set("language", string(notebookInfo.Language))
	if err != nil {
		return err
	}
	err = d.Set("object_id", int(notebookInfo.ObjectID))
	if err != nil {
		return err
	}
	err = d.Set("object_type", string(notebookInfo.ObjectType))
	if err != nil {
		return err
	}

	err = d.Set("content", notebookContent)

	return err
}
