package model

// FileInfo contains information when listing files or fetching files from DBFS api
type FileInfo struct {
	Path     string `json:"path,omitempty"`
	IsDir    bool   `json:"is_dir,omitempty"`
	FileSize int64  `json:"file_size,omitempty"`
}
