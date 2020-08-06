package storage

import (
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceDBFSFile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDBFSFileRead,
		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"limit_file_size": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"file_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceDBFSFileRead(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	limitFileSize := d.Get("limit_file_size").(bool)
	client := m.(*common.DatabricksClient)

	fileInfo, err := NewDBFSAPI(client).Status(path)
	if err != nil {
		return err
	}

	d.SetId(path)
	err = d.Set("path", fileInfo.Path)
	if err != nil {
		return err
	}
	err = d.Set("file_size", fileInfo.FileSize)
	if err != nil {
		return err
	}

	if limitFileSize {
		if fileInfo.FileSize < 4e6 {
			content, err := NewDBFSAPI(client).Read(path)
			if err != nil {
				return err
			}
			err = d.Set("content", content)
			if err != nil {
				return err
			}
		} else if fileInfo.FileSize > 4e6 {
			err = d.Set("content", "File Size is too Large!")
		}
	} else {
		content, err := NewDBFSAPI(client).Read(path)
		if err != nil {
			return err
		}
		err = d.Set("content", content)
		if err != nil {
			return err
		}
	}

	return err
}
