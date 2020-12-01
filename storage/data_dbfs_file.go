package storage

import (
	"context"
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceDBFSFile ...
func DataSourceDBFSFile() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			limitFileSize := d.Get("limit_file_size").(bool)
			dbfsAPI := NewDbfsAPI(ctx, m)
			fileInfo, err := dbfsAPI.Status(d.Get("path").(string))
			if err != nil {
				return diag.FromErr(err)
			}
			// TODO: DEPRECATE/ make default
			if limitFileSize && fileInfo.FileSize > 4e6 {
				return diag.Errorf("Size of %s is too large: %d bytes",
					fileInfo.Path, fileInfo.FileSize)
			}
			d.SetId(fileInfo.Path)
			d.Set("path", fileInfo.Path)
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
				Deprecated: "Would become client property",
				Type:       schema.TypeBool,
				Required:   true,
				ForceNew:   true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"file_size": {
				Deprecated: "Rename to size?...",
				Type:       schema.TypeInt,
				Computed:   true,
			},
		},
	}
}
