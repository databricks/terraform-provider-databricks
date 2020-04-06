package db

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"path/filepath"
)

func resourceNotebook() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotebookCreate,
		Read:   resourceNotebookRead,
		Delete: resourceNotebookDelete,

		Schema: map[string]*schema.Schema{
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(model.Scala),
					string(model.Python),
					string(model.R),
					string(model.SQL),
				}, false),
			},
			"overwrite": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"mkdirs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  string(model.Source),
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(model.DBC),
					string(model.Source),
					string(model.Html),
				}, false),
			},
			"object_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceNotebookCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	path := d.Get("path").(string)
	content := d.Get("content").(string)
	language := d.Get("language").(string)
	format := d.Get("format").(string)
	overwrite := d.Get("overwrite").(bool)
	mkdirs := d.Get("mkdirs").(bool)

	if mkdirs == true {
		parentDir := filepath.Dir(path)
		err := client.Notebooks().Mkdirs(parentDir)
		if err != nil {
			return err
		}
	}

	err := client.Notebooks().Create(path, content, model.Language(language), model.ExportFormat(format), overwrite)
	if err != nil {
		return err
	}
	d.SetId(path)
	err = d.Set("content", content)
	if err != nil {
		return err
	}
	err = d.Set("format", format)
	if err != nil {
		return err
	}
	err = d.Set("overwrite", overwrite)
	if err != nil {
		return err
	}
	err = d.Set("mkdirs", overwrite)
	if err != nil {
		return err
	}
	return resourceNotebookRead(d, m)
}

func resourceNotebookRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(service.DBApiClient)
	notebookInfo, err := client.Notebooks().Read(id)
	if err != nil {
		return err
	}
	err = d.Set("path", id)
	if err != nil {
		return err
	}

	err = d.Set("language", string(notebookInfo.Language))
	if err != nil {
		return err
	}
	err = d.Set("object_id", int(notebookInfo.ObjectId))
	if err != nil {
		return err
	}
	err = d.Set("object_type", string(notebookInfo.ObjectType))

	return err
}

func resourceNotebookDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(service.DBApiClient)
	err := client.Notebooks().Delete(id, true)
	return err
}
