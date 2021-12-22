package storage

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
)

// FileList contains list of file metadata entries
type FileList struct {
	Files []FileInfo `json:"files,omitempty"`
}

// FileInfo contains information when listing files or fetching files from DBFS api
type FileInfo struct {
	Path     string `json:"path,omitempty"`
	IsDir    bool   `json:"is_dir,omitempty"`
	FileSize int64  `json:"file_size,omitempty"`
}

// createHandle contains the payload to create a handle which is a connection for uploading blocks of file data
type createHandle struct {
	Path      string `json:"path,omitempty"`
	Overwrite bool   `json:"overwrite,omitempty"`
}

// handleResponse contains the response from making an handle request
type handleResponse struct {
	Handle int64 `json:"handle,omitempty"`
}

// addBlock contains the payload to upload a block of base64 data to a handle
type addBlock struct {
	Data   string `json:"data,omitempty"`
	Handle int64  `json:"handle,omitempty"`
}

// ReadResponse contains the response from reading a portion of a file in DBFS
type ReadResponse struct {
	BytesRead int64  `json:"bytes_read"`
	Data      string `json:"data"`
}

// NewDbfsAPI creates DBFSAPI instance from provider meta
func NewDbfsAPI(ctx context.Context, m interface{}) DbfsAPI {
	return DbfsAPI{m.(*common.DatabricksClient), ctx}
}

// DbfsAPI exposes the DBFS API
type DbfsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a file on DBFS
func (a DbfsAPI) Create(path string, contents []byte, overwrite bool) (err error) {
	handle, err := a.createHandle(path, overwrite)
	if err != nil {
		err = fmt.Errorf("cannot create handle: %w", err)
	}
	defer func() {
		cerr := a.closeHandle(handle)
		if cerr != nil {
			err = fmt.Errorf("cannot close handle: %w", cerr)
		}
	}()
	buffer := bytes.NewBuffer(contents)
	for {
		byteChunk := buffer.Next(1e6)
		if len(byteChunk) == 0 {
			break
		}
		b64Data := base64.StdEncoding.EncodeToString(byteChunk)
		err = a.addBlock(b64Data, handle)
		if err != nil {
			err = fmt.Errorf("cannot add block: %w", err)
		}
	}
	return
}

func (a DbfsAPI) createHandle(path string, overwrite bool) (int64, error) {
	var h handleResponse
	err := a.client.Post(a.context, "/dbfs/create", createHandle{path, overwrite}, &h)
	return h.Handle, err
}

func (a DbfsAPI) addBlock(data string, handle int64) error {
	return a.client.Post(a.context, "/dbfs/add-block", addBlock{data, handle}, nil)
}

func (a DbfsAPI) closeHandle(handle int64) error {
	return a.client.Post(a.context, "/dbfs/close", handleResponse{handle}, nil)
}

// List returns a list of files in DBFS and the recursive flag lets you recursively list files
func (a DbfsAPI) List(path string, recursive bool) ([]FileInfo, error) {
	if recursive {
		var paths []FileInfo
		err := a.recursiveAddPaths(path, &paths)
		if err != nil {
			return nil, err
		}
		return paths, err
	}
	return a.list(path)
}

func (a DbfsAPI) recursiveAddPaths(path string, pathList *[]FileInfo) error {
	fileInfoList, err := a.list(path)
	if err != nil {
		return err
	}
	for _, v := range fileInfoList {
		if !v.IsDir {
			*pathList = append(*pathList, v)
		} else if v.IsDir {
			err := a.recursiveAddPaths(v.Path, pathList)
			if err != nil {
				return fmt.Errorf("cannot list subfolder: %w", err)
			}
		}
	}
	return nil
}

func (a DbfsAPI) list(path string) ([]FileInfo, error) {
	var dbfsList FileList
	err := a.client.Get(a.context, "/dbfs/list", map[string]interface{}{
		"path": path,
	}, &dbfsList)
	if err != nil {
		err = fmt.Errorf("cannot list %s: %w", path, err)
	}
	return dbfsList.Files, err
}

// Delete deletes a file in DBFS via API
func (a DbfsAPI) Delete(path string, recursive bool) error {
	return a.client.Post(a.context, "/dbfs/delete", dbfsRequest{
		Path:      path,
		Recursive: recursive,
	}, nil)
}

type dbfsRequest struct {
	Path      string `json:"path,omitempty" url:"path,omitempty"`
	Offset    int64  `json:"offset,omitempty" url:"offset,omitempty"`
	Length    int64  `json:"length,omitempty" url:"length,omitempty"`
	Recursive bool   `json:"recursive,omitempty" url:"recursive,omitempty"`
}

// Read returns the contents of a file
func (a DbfsAPI) Read(path string) (content []byte, err error) {
	fetchLoop := true
	offSet := int64(0)
	length := int64(1e6)
	for fetchLoop {
		bytesRead, bytes, err := a.read(path, offSet, length)
		if err != nil {
			return content, fmt.Errorf("cannot read %s: %w", path, err)
		}
		if bytesRead == 0 || bytesRead < length {
			fetchLoop = false
		}
		content = append(content, bytes...)
		offSet += length
	}
	return content, err
}

func (a DbfsAPI) read(path string, offset, length int64) (int64, []byte, error) {
	bytesRead, data, err := a.readString(path, offset, length)
	if err != nil {
		return bytesRead, nil, err
	}
	dataBytes, err := base64.StdEncoding.DecodeString(data)
	return bytesRead, dataBytes, err
}

// readString reads a "block" of data in DBFS given a offset and length as a base64 encoded string
func (a DbfsAPI) readString(path string, offset, length int64) (int64, string, error) {
	var readBytes ReadResponse
	err := a.client.Get(a.context, "/dbfs/read", dbfsRequest{
		Path:   path,
		Offset: offset,
		Length: length,
	}, &readBytes)
	return readBytes.BytesRead, readBytes.Data, err
}

// Status returns the status of a file in DBFS
func (a DbfsAPI) Status(path string) (f FileInfo, err error) {
	err = a.client.Get(a.context, "/dbfs/get-status", map[string]interface{}{
		"path": path,
	}, &f)
	return
}
