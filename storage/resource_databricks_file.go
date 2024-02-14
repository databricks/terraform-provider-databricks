package storage

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"os"

	"github.com/databricks/databricks-sdk-go/service/files"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getContentReader(data *schema.ResourceData) (io.ReadCloser, error) {
	source := data.Get("source").(string)
	var reader io.ReadCloser
	var err error
	if source != "" {
		reader, err = os.Open(source)
		if err != nil {
			return nil, err
		}
	}
	contentBase64 := data.Get("content_base64").(string)
	if contentBase64 != "" {
		decodedString, err := base64.StdEncoding.DecodeString(contentBase64)
		if err != nil {
			return nil, err
		}
		reader = io.NopCloser(bytes.NewReader(decodedString))
		if err != nil {
			return nil, err
		}
	}
	return reader, err
}

func ResourceFile() common.Resource {
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
			reader, err := getContentReader(data)
			if err != nil {
				return err
			}
			err = w.Files.Upload(ctx, files.UploadRequest{Contents: reader, FilePath: path})
			if err != nil {
				return err
			}

			metadata, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: path})
			if err != nil {
				return err
			}
			data.Set("modification_time", metadata.LastModified)
			data.Set("file_size", metadata.ContentLength)

			data.SetId(path)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			path := data.Id()
			metadata, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: path})
			if err != nil {
				return err
			}
			data.Set("modification_time", metadata.LastModified)
			data.Set("file_size", metadata.ContentLength)
			return common.StructToData(metadata, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			path := data.Id()
			reader, err := getContentReader(data)
			if err != nil {
				return err
			}
			err = w.Files.Upload(ctx, files.UploadRequest{Contents: reader, FilePath: path})
			if err != nil {
				return err
			}
			metadata, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: path})
			if err != nil {
				return err
			}
			data.Set("modification_time", metadata.LastModified)
			data.Set("file_size", metadata.ContentLength)

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
	}
}
