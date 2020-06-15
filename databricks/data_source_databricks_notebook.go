package databricks

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceNotebook() *schema.Resource {
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
					string(model.DBC),
					string(model.Source),
					string(model.HTML),
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
	client := m.(*service.DBApiClient)

	notebookInfo, err := client.Notebooks().Read(path)
	if err != nil {
		return err
	}
	notebookContent, err := client.Notebooks().Export(path, model.ExportFormat(format))
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
