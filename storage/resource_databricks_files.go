package storage

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// look at resource_workspace_file.go for ref, we aren't yet going to use modification time to be consistent with other file resources
func ResourceFiles() *schema.Resource {
	s := workspace.FileContentSchema(map[string]*schema.Schema{
		"modification_time": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"file_size": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	})
	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			path := data.Get("path").(string)
			source := data.Get("source").(string)
			reader := io.Reader(strings.NewReader(source))

			err = w.Files.Upload(ctx, source, reader)
			if err != nil {
				return err
			}

			data.SetId(path)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			path := data.Get("path").(string)
			reader, err := w.Files.Download(ctx, path)
			if err != nil {
				return err
			}

			fileInfo, err := w.Files.GetStatus(ctx, path)
			if err != nil {
				return err
			}

			data.Set("modification_time", fileInfo.ModificationTime)

			source := data.Get("source").(string)
			content, err := workspace.readFileContent(source)
			if err != nil {
				return err
			}
			data.Set("md5", fmt.Sprintf("%x", md5.Sum(content)))

			reader.Read([]byte(source))
			return nil
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			path := data.Get("path").(string)
			err = w.Files.Upload(ctx, path)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			path := data.Get("path").(string)
			err = w.Files.Delete(ctx, path)
			return err
		},
		Schema: s,
	}.ToResource()
}
