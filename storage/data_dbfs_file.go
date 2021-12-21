package storage

import (
	"context"
	"fmt"	
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceDBFSFile ...
func DataSourceDBFSFile() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			limitFileSize := d.Get("limit_file_size").(int)
			dbfsAPI := NewDbfsAPI(ctx, m)
			fileInfo, err := dbfsAPI.Status(d.Get("path").(string))
			if err != nil {
				return diag.FromErr(err)
			}
			if fileInfo.FileSize > 4e6 {
				return diag.Errorf("Size of %s is too large: %d bytes",
					fileInfo.Path, fileInfo.FileSize)
			}
			if fileInfo.FileSize <= limitFileSize {
				return diag.Errorf("Size of %s (%d bytes) is under limit_file_size limit (%d bytes)",
					fileInfo.Path, fileInfo.FileSize, limitFileSize)
			}
			d.SetId(fileInfo.Path)
			d.Set("path", fileInfo.Path)
			d.Set("dbfs_path", fmt.Sprint("dbfs:", fileInfo.Path))			
			d.Set("file_size", fileInfo.FileSize)
			content, err := dbfsAPI.Read(fileInfo.Path)
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("content", base64.StdEncoding.EncodeToString(content))
			return nil
		},
		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"limit_file_size": {
				Type:     schema.TypeInt,
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
			"dbfs_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
