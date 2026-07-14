package storage

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceDbfsFile() common.Resource {
	s := map[string]*schema.Schema{
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
	}
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)
	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			newClient, err := m.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			limitFileSize := d.Get("limit_file_size").(bool)
			dbfsAPI := NewDbfsAPI(ctx, newClient)
			fileInfo, err := dbfsAPI.Status(d.Get("path").(string))
			if err != nil {
				return err
			}
			if limitFileSize && fileInfo.FileSize > 4e6 {
				return fmt.Errorf("size of %s is too large: %d bytes",
					fileInfo.Path, fileInfo.FileSize)
			}
			d.SetId(fileInfo.Path)
			d.Set("path", fileInfo.Path)
			d.Set("file_size", fileInfo.FileSize)
			content, err := dbfsAPI.Read(fileInfo.Path)
			if err != nil {
				return err
			}
			d.Set("content", base64.StdEncoding.EncodeToString(content))
			return nil
		},
	}
}
