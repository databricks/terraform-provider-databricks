package storage

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"io"
	"os"

	"github.com/databricks/databricks-sdk-go/service/files"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type hashReadCloser struct {
	io.Reader
	io.Closer
	hash.Hash
}

// Note: we don't use workspace.ReadContent here because that one reads all the content into memory,
// and `resource_file` supports files of up to 5GB.
func getContentReader(data *schema.ResourceData) (*hashReadCloser, error) {
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
	hash := md5.New()
	tee := io.TeeReader(reader, hash)
	teeCloser := hashReadCloser{tee, reader, hash}
	return &teeCloser, err
}

func upload(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient, path string) error {
	w, err := common.WorkspaceClientUnifiedProvider(ctx, data, c)
	if err != nil {
		return err
	}
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
	data.Set("remote_file_modified", false)
	data.Set("md5", hex.EncodeToString((*reader).Sum(nil)))
	data.SetId(path)
	return nil
}

func ResourceFile() common.Resource {
	s := workspace.FileContentSchema(map[string]*schema.Schema{
		"modification_time": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"file_size": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"remote_file_modified": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			path := data.Get("path").(string)
			err := upload(ctx, data, c, path)
			if err != nil {
				return err
			}
			data.SetId(path)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := common.WorkspaceClientUnifiedProvider(ctx, data, c)
			if err != nil {
				return err
			}

			path := data.Id()
			metadata, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: path})
			if err != nil {
				return err
			}
			storedModificationTime := data.Get("modification_time").(string)

			data.Set("remote_file_modified", storedModificationTime != metadata.LastModified)

			// Do not store here the modification time. If the update fails, we will keep the wrong one in the state.

			return common.StructToData(metadata, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			path := data.Id()
			return upload(ctx, data, c, path)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := common.WorkspaceClientUnifiedProvider(ctx, data, c)
			if err != nil {
				return err
			}
			path := data.Id()
			err = w.Files.Delete(ctx, files.DeleteFileRequest{FilePath: path})
			return err
		},
	}
}
