package model

// DBFSFileInfo contains information when listing files or fetching files from DBFS api
type DBFSFileInfo struct {
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
