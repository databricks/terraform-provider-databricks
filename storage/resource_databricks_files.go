package storage

import (
	"bufio"
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/databricks/databricks-sdk-go/service/files"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
		Schema: s,
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			path := data.Get("path").(string)
			source := data.Get("source").(string)
			reader, err := os.Open(source)
			if err != nil {
				return err
			}

			err = w.Files.Upload(ctx, files.UploadRequest{Contents: reader, FilePath: path})
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

			path := data.Id()
			fileInfo, err := w.Files.GetStatus(ctx, files.GetStatusRequest{Path: path})
			if err != nil {
				return err
			}
			data.Set("modification_time", fileInfo.ModificationTime)
			data.Set("file_size", fileInfo.FileSize)

			downloadResponse, err := w.Files.Download(ctx, files.DownloadRequest{FilePath: path})
			if err != nil {
				return err
			}

			dataReadCloser := downloadResponse.Contents
			defer dataReadCloser.Close()
			reader := bufio.NewReader(dataReadCloser)
			dataByte, err := ioutil.ReadAll(reader)
			if err != nil {
				return err
			}
			data.Set("md5", fmt.Sprintf("%x", md5.Sum(dataByte)))
			return common.StructToData(fileInfo, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			path := data.Id()
			source := data.Get("source").(string)
			reader, err := os.Open(source)
			if err != nil {
				return err
			}
			err = w.Files.Upload(ctx, files.UploadRequest{Contents: reader, FilePath: path})
			return err
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			path := data.Id()
			err = w.Files.Delete(ctx, files.DeleteFileRequest{FilePath: path})
			return err
		},
	}.ToResource()
}
