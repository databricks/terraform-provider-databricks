package databricks

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDBFSFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceDBFSFileCreate,
		Read:   resourceDBFSFileRead,
		Delete: resourceDBFSFileDelete,
		Update: resourceDBFSFileUpdate,

		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"mkdirs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"validate_remote_file": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"file_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceDBFSFileCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	path := d.Get("path").(string)
	content := d.Get("content").(string)
	overwrite := d.Get("overwrite").(bool)
	mkdirs := d.Get("mkdirs").(bool)

	if mkdirs {
		parentDir := filepath.Dir(path)
		err := client.DBFS().Mkdirs(parentDir)
		if err != nil {
			return err
		}
	}

	err := client.DBFS().Create(path, overwrite, content)
	if err != nil {
		return err
	}

	d.SetId(path)
	err = d.Set("content", content)
	if err != nil {
		return err
	}
	err = d.Set("overwrite", overwrite)
	if err != nil {
		return err
	}
	err = d.Set("mkdirs", mkdirs)
	if err != nil {
		return err
	}

	return resourceDBFSFileRead(d, m)
}

func resourceDBFSFileRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)

	fileInfo, err := client.DBFS().Status(id)
	if err != nil {
		if isDBFSFileMissing(err.Error(), id) {
			log.Printf("Missing dbfs file with id: %s.", id)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("path", fileInfo.Path)
	if err != nil {
		return err
	}
	err = d.Set("file_size", fileInfo.FileSize)

	if validateRemoteFile, ok := d.GetOk("validate_remote_file"); ok {
		validateFile := validateRemoteFile.(bool)
		if validateFile {
			log.Println("Validating remote file!")
			data, err := client.DBFS().Read(id)
			if err != nil {
				return err
			}
			err = d.Set("content", data)
			if err != nil {
				return err
			}
		}
	}

	return err
}

func resourceDBFSFileUpdate(d *schema.ResourceData, m interface{}) error {
	overwrite := d.Get("overwrite").(bool)
	mkdirs := d.Get("mkdirs").(bool)

	err := d.Set("overwrite", overwrite)
	if err != nil {
		return err
	}
	err = d.Set("mkdirs", mkdirs)
	if err != nil {
		return err
	}
	return resourceDBFSFileRead(d, m)
}
func resourceDBFSFileDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	err := client.DBFS().Delete(id, false)
	return err
}

func isDBFSFileMissing(errorMsg, resourceID string) bool {
	return strings.Contains(errorMsg, "RESOURCE_DOES_NOT_EXIST") &&
		strings.Contains(errorMsg, fmt.Sprintf("No file or directory exists on path %s.", resourceID))
}
