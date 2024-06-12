// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AddBlock struct {
	// The base64-encoded data to append to the stream. This has a limit of 1
	// MB.
	Data string `tfsdk:"data"`
	// The handle on an open stream.
	Handle int64 `tfsdk:"handle"`
}

type AddBlockResponse struct {
}

type Close struct {
	// The handle on an open stream.
	Handle int64 `tfsdk:"handle"`
}

type CloseResponse struct {
}

type Create struct {
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite bool `tfsdk:"overwrite"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path string `tfsdk:"path"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Create) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Create) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Create a directory
type CreateDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath string `tfsdk:"-" url:"-"`
}

type CreateDirectoryResponse struct {
}

type CreateResponse struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	Handle int64 `tfsdk:"handle"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Delete struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path.
	Path string `tfsdk:"path"`
	// Whether or not to recursively delete the directory's contents. Deleting
	// empty directories can be done without providing the recursive flag.
	Recursive bool `tfsdk:"recursive"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Delete) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Delete) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a directory
type DeleteDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath string `tfsdk:"-" url:"-"`
}

type DeleteDirectoryResponse struct {
}

// Delete a file
type DeleteFileRequest struct {
	// The absolute path of the file.
	FilePath string `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

type DirectoryEntry struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize int64 `tfsdk:"file_size"`
	// True if the path is a directory.
	IsDirectory bool `tfsdk:"is_directory"`
	// Last modification time of given file in milliseconds since unix epoch.
	LastModified int64 `tfsdk:"last_modified"`
	// The name of the file or directory. This is the last component of the
	// path.
	Name string `tfsdk:"name"`
	// The absolute path of the file or directory.
	Path string `tfsdk:"path"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DirectoryEntry) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DirectoryEntry) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Download a file
type DownloadRequest struct {
	// The absolute path of the file.
	FilePath string `tfsdk:"-" url:"-"`
}

type DownloadResponse struct {
	ContentLength int64 `tfsdk:"-" url:"-" header:"content-length,omitempty"`

	ContentType string `tfsdk:"-" url:"-" header:"content-type,omitempty"`

	Contents io.ReadCloser `tfsdk:"-"`

	LastModified string `tfsdk:"-" url:"-" header:"last-modified,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DownloadResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DownloadResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileInfo struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize int64 `tfsdk:"file_size"`
	// True if the path is a directory.
	IsDir bool `tfsdk:"is_dir"`
	// Last modification time of given file in milliseconds since epoch.
	ModificationTime int64 `tfsdk:"modification_time"`
	// The absolute path of the file or directory.
	Path string `tfsdk:"path"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *FileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get directory metadata
type GetDirectoryMetadataRequest struct {
	// The absolute path of a directory.
	DirectoryPath string `tfsdk:"-" url:"-"`
}

type GetDirectoryMetadataResponse struct {
}

// Get file metadata
type GetMetadataRequest struct {
	// The absolute path of the file.
	FilePath string `tfsdk:"-" url:"-"`
}

type GetMetadataResponse struct {
	ContentLength int64 `tfsdk:"-" url:"-" header:"content-length,omitempty"`

	ContentType string `tfsdk:"-" url:"-" header:"content-type,omitempty"`

	LastModified string `tfsdk:"-" url:"-" header:"last-modified,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the information of a file or directory
type GetStatusRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path string `tfsdk:"-" url:"path"`
}

// List directory contents or file details
type ListDbfsRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path string `tfsdk:"-" url:"path"`
}

// List directory contents
type ListDirectoryContentsRequest struct {
	// The absolute path of a directory.
	DirectoryPath string `tfsdk:"-" url:"-"`
	// The maximum number of directory entries to return. The response may
	// contain fewer entries. If the response contains a `next_page_token`,
	// there may be more entries, even if fewer than `page_size` entries are in
	// the response.
	//
	// We recommend not to set this value unless you are intentionally listing
	// less than the complete directory contents.
	//
	// If unspecified, at most 1000 directory entries will be returned. The
	// maximum value is 1000. Values above 1000 will be coerced to 1000.
	PageSize int64 `tfsdk:"-" url:"page_size,omitempty"`
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the contents of this directory. Provide this
	// token to retrieve the next page of directory entries. When providing a
	// `page_token`, all other parameters provided to the request must match the
	// previous request. To list all of the entries in a directory, it is
	// necessary to continue requesting pages of entries until the response
	// contains no `next_page_token`. Note that the number of entries returned
	// must not be used to determine when the listing is complete.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListDirectoryContentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDirectoryContentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDirectoryResponse struct {
	// Array of DirectoryEntry.
	Contents []DirectoryEntry `tfsdk:"contents"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListDirectoryResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDirectoryResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListStatusResponse struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	Files []FileInfo `tfsdk:"files"`
}

type MkDirs struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	Path string `tfsdk:"path"`
}

type MkDirsResponse struct {
}

type Move struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path.
	DestinationPath string `tfsdk:"destination_path"`
	// The source path of the file or directory. The path should be the absolute
	// DBFS path.
	SourcePath string `tfsdk:"source_path"`
}

type MoveResponse struct {
}

type Put struct {
	// This parameter might be absent, and instead a posted file will be used.
	Contents string `tfsdk:"contents"`
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite bool `tfsdk:"overwrite"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path string `tfsdk:"path"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Put) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Put) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PutResponse struct {
}

// Get the contents of a file
type ReadDbfsRequest struct {
	// The number of bytes to read starting from the offset. This has a limit of
	// 1 MB, and a default value of 0.5 MB.
	Length int64 `tfsdk:"-" url:"length,omitempty"`
	// The offset to read from in bytes.
	Offset int64 `tfsdk:"-" url:"offset,omitempty"`
	// The path of the file to read. The path should be the absolute DBFS path.
	Path string `tfsdk:"-" url:"path"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ReadDbfsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReadDbfsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ReadResponse struct {
	// The number of bytes read (could be less than ``length`` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	BytesRead int64 `tfsdk:"bytes_read"`
	// The base64-encoded contents of the file read.
	Data string `tfsdk:"data"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ReadResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReadResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Upload a file
type UploadRequest struct {
	Contents io.ReadCloser `tfsdk:"-"`
	// The absolute path of the file.
	FilePath string `tfsdk:"-" url:"-"`
	// If true, an existing file will be overwritten.
	Overwrite bool `tfsdk:"-" url:"overwrite,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UploadRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UploadRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UploadResponse struct {
}
