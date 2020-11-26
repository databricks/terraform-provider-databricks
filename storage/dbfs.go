package storage

import (
	"context"
	"encoding/base64"

	"github.com/databrickslabs/databricks-terraform/common"
)

// FileInfo contains information when listing files or fetching files from DBFS api
type FileInfo struct {
	Path     string `json:"path,omitempty"`
	IsDir    bool   `json:"is_dir,omitempty"`
	FileSize int64  `json:"file_size,omitempty"`
}

// DBFSHandleRequest contains the payload to create a handle which is a connection for uploading blocks of file data
type DBFSHandleRequest struct {
	Path      string `json:"path,omitempty" url:"path,omitempty"`
	Overwrite bool   `json:"overwrite,omitempty" url:"overwrite,omitempty"`
}

// DBFSHandleResponse contains the response from making an handle request
type DBFSHandleResponse struct {
	Handle int64 `json:"handle,omitempty" url:"handle,omitempty"`
}

// DBFSBlockRequest contains the payload to upload a block of base64 data to a handle
type DBFSBlockRequest struct {
	Data   string `json:"data,omitempty" url:"data,omitempty"`
	Handle int64  `json:"handle,omitempty" url:"handle,omitempty"`
}

// DBFSCloseRequest contains the payload close an opened connection (handle) to a dbfs path
type DBFSCloseRequest struct {
	Handle int64 `json:"handle,omitempty" url:"handle,omitempty"`
}

// DBFSReadResponse contains the response from reading a portion of a file in DBFS
type DBFSReadResponse struct {
	BytesRead int64  `json:"bytes_read,omitempty" url:"bytes_read,omitempty"`
	Data      string `json:"data,omitempty" url:"data,omitempty"`
}

// DBFSMkdirRequest contains the payload to make a directory in dbfs
type DBFSMkdirRequest struct {
	Path string `json:"path,omitempty" url:"path,omitempty"`
}

// DBFSDeleteRequest contains the payload to delete a file/directory in dbfs
type DBFSDeleteRequest struct {
	Path      string `json:"path,omitempty" url:"path,omitempty"`
	Recursive bool   `json:"recursive,omitempty" url:"recursive,omitempty"`
}

// NewDBFSAPI creates DBFSAPI instance from provider meta
func NewDBFSAPI(m interface{}) DBFSAPI {
	return DBFSAPI{m.(*common.DatabricksClient), context.TODO()}
}

// DBFSAPI exposes the DBFS API
type DBFSAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a file in DBFS given data string in base64
func (a DBFSAPI) Create(path string, overwrite bool, data string) (err error) {
	byteArr, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}
	byteChunks := split(byteArr, 1e6)
	handle, err := a.createHandle(path, overwrite)
	if err != nil {
		return
	}
	defer func() {
		cerr := a.closeHandle(handle)
		if cerr != nil {
			err = cerr
		}
	}()
	for _, byteChunk := range byteChunks {
		b64Data := base64.StdEncoding.EncodeToString(byteChunk)
		err = a.addBlock(b64Data, handle)
		if err != nil {
			return
		}
	}
	return
}

// Read returns the contents of a file in DBFS as a base64 encoded string
func (a DBFSAPI) Read(path string) (string, error) {
	var bytesFetched []byte
	fetchLoop := true
	offSet := int64(0)
	length := int64(1e6)
	for fetchLoop {
		bytesRead, bytes, err := a.read(path, offSet, length)
		if err != nil {
			return "", err
		}
		if bytesRead == 0 || bytesRead < length {
			fetchLoop = false
		}

		bytesFetched = append(bytesFetched, bytes...)
		offSet += length
	}
	resp := base64.StdEncoding.EncodeToString(bytesFetched)
	return resp, nil
}

func (a DBFSAPI) read(path string, offset, length int64) (int64, []byte, error) {
	bytesRead, data, err := a.ReadString(path, offset, length)
	if err != nil {
		return bytesRead, nil, err
	}
	dataBytes, err := base64.StdEncoding.DecodeString(data)
	return bytesRead, dataBytes, err
}

// List returns a list of files in DBFS and the recursive flag lets you recursively list files
func (a DBFSAPI) List(path string, recursive bool) ([]FileInfo, error) {
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

func (a DBFSAPI) recursiveAddPaths(path string, pathList *[]FileInfo) error {
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
				return err
			}
		}
	}
	return nil
}

// Move moves the file between DBFS locations via DBFS api
func (a DBFSAPI) Move(src string, tgt string) error {
	return a.client.Post(a.context, "/dbfs/move", map[string]string{
		"source_path":      src,
		"destination_path": tgt,
	}, nil)
}

// Delete deletes a file in DBFS via API
func (a DBFSAPI) Delete(path string, recursive bool) error {
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

// ReadString reads a "block" of data in DBFS given a offset and length as a base64 encoded string
func (a DBFSAPI) ReadString(path string, offset, length int64) (int64, string, error) {
	var readBytes struct {
		BytesRead int64  `json:"bytes_read,omitempty"`
		Data      string `json:"data,omitempty"`
	}
	err := a.client.Get(a.context, "/dbfs/read", dbfsRequest{
		Path:   path,
		Offset: offset,
		Length: length,
	}, &readBytes)
	return readBytes.BytesRead, readBytes.Data, err
}

// Status returns the status of a file in DBFS
func (a DBFSAPI) Status(path string) (f FileInfo, err error) {
	err = a.client.Get(a.context, "/dbfs/get-status", map[string]interface{}{
		"path": path,
	}, &f)
	return
}

type FileList struct {
	Files []FileInfo `json:"files,omitempty" url:"files,omitempty"`
}

func (a DBFSAPI) list(path string) ([]FileInfo, error) {
	var dbfsList FileList
	err := a.client.Get(a.context, "/dbfs/list", map[string]interface{}{
		"path": path,
	}, &dbfsList)
	return dbfsList.Files, err
}

// Mkdirs makes the directories in DBFS include the parent paths
func (a DBFSAPI) Mkdirs(path string) error {
	return a.client.Post(a.context, "/dbfs/mkdirs", map[string]interface{}{
		"path": path,
	}, nil)
}

func (a DBFSAPI) createHandle(path string, overwrite bool) (int64, error) {
	var handle struct {
		Handle int64 `json:"handle,omitempty" url:"handle,omitempty"`
	}
	err := a.client.Post(a.context, "/dbfs/create", map[string]interface{}{
		"path":      path,
		"overwrite": overwrite,
	}, &handle)
	return handle.Handle, err
}

func (a DBFSAPI) addBlock(data string, handle int64) error {
	return a.client.Post(a.context, "/dbfs/add-block", map[string]interface{}{
		"data":   data,
		"handle": handle,
	}, nil)
}

func (a DBFSAPI) closeHandle(handle int64) error {
	return a.client.Post(a.context, "/dbfs/close", map[string]interface{}{
		"handle": handle,
	}, nil)
}

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf)
	}
	return chunks
}
