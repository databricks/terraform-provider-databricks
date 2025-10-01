package storage

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DbfsFileData struct {
	Path          string `json:"path"`
	LimitFileSize bool   `json:"limit_file_size"`
	Content       string `json:"content"`
	FileSize      int64  `json:"file_size,omitempty" tf:"computed"`
	DbfsPath      string `json:"dbfs_path,omitempty" tf:"computed"`
	common.ProviderConfig
}

func DataSourceDbfsFile() common.Resource {
	s := common.StructToSchema(DbfsFileData{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["path"].Required = true
		m["path"].ForceNew = true
		m["limit_file_size"].Required = true
		m["limit_file_size"].ForceNew = true
		m["content"].Computed = true
		m["content"].ForceNew = true
		m["file_size"].Computed = true
		m["file_size"].ForceNew = true
		m["dbfs_path"].Computed = true
		m["dbfs_path"].ForceNew = true

		// Add provider_config customizations
		common.CustomizeSchemaPath(m, "provider_config").SetOptional()
		common.CustomizeSchemaPath(m, "provider_config", "workspace_id").SetRequired()

		return m
	})
	return common.Resource{
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			var dfd DbfsFileData
			common.DataToStructPointer(d, s, &dfd)

			dbfsAPI := NewDbfsAPI(ctx, m)
			fileInfo, err := dbfsAPI.Status(dfd.Path)
			if err != nil {
				return err
			}
			if dfd.LimitFileSize && fileInfo.FileSize > 4e6 {
				return fmt.Errorf("size of %s is too large: %d bytes",
					fileInfo.Path, fileInfo.FileSize)
			}

			d.SetId(fileInfo.Path)
			dfd.Path = fileInfo.Path
			dfd.FileSize = fileInfo.FileSize
			dfd.DbfsPath = fmt.Sprintf("dbfs:%s", fileInfo.Path)

			content, err := dbfsAPI.Read(fileInfo.Path)
			if err != nil {
				return err
			}
			dfd.Content = base64.StdEncoding.EncodeToString(content)

			return common.StructToData(dfd, s, d)
		},
		Schema: s,
	}
}
