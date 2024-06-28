// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"io"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddBlock struct {
	// The base64-encoded data to append to the stream. This has a limit of 1
	// MB.
	Data types.String `tfsdk:"data"`
	// The handle on an open stream.
	Handle types.Int64 `tfsdk:"handle"`
}

type AddBlockResponse struct {
}

type Close struct {
	// The handle on an open stream.
	Handle types.Int64 `tfsdk:"handle"`
}

type CloseResponse struct {
}

type Create struct {
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite types.Bool `tfsdk:"overwrite"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path"`
}

// Create a directory
type CreateDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-" url:"-"`
}

type CreateDirectoryResponse struct {
}

type CreateResponse struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	Handle types.Int64 `tfsdk:"handle"`
}

type Delete struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path.
	Path types.String `tfsdk:"path"`
	// Whether or not to recursively delete the directory's contents. Deleting
	// empty directories can be done without providing the recursive flag.
	Recursive types.Bool `tfsdk:"recursive"`
}

// Delete a directory
type DeleteDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-" url:"-"`
}

type DeleteDirectoryResponse struct {
}

// Delete a file
type DeleteFileRequest struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

type DirectoryEntry struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize types.Int64 `tfsdk:"file_size"`
	// True if the path is a directory.
	IsDirectory types.Bool `tfsdk:"is_directory"`
	// Last modification time of given file in milliseconds since unix epoch.
	LastModified types.Int64 `tfsdk:"last_modified"`
	// The name of the file or directory. This is the last component of the
	// path.
	Name types.String `tfsdk:"name"`
	// The absolute path of the file or directory.
	Path types.String `tfsdk:"path"`
}

// Download a file
type DownloadRequest struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-" url:"-"`
}

type DownloadResponse struct {
	ContentLength types.Int64 `tfsdk:"-" url:"-" header:"content-length,omitempty"`

	ContentType types.String `tfsdk:"-" url:"-" header:"content-type,omitempty"`

	Contents io.ReadCloser `tfsdk:"-"`

	LastModified types.String `tfsdk:"-" url:"-" header:"last-modified,omitempty"`
}

type FileInfo struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize types.Int64 `tfsdk:"file_size"`
	// True if the path is a directory.
	IsDir types.Bool `tfsdk:"is_dir"`
	// Last modification time of given file in milliseconds since epoch.
	ModificationTime types.Int64 `tfsdk:"modification_time"`
	// The absolute path of the file or directory.
	Path types.String `tfsdk:"path"`
}

// Get directory metadata
type GetDirectoryMetadataRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-" url:"-"`
}

type GetDirectoryMetadataResponse struct {
}

// Get file metadata
type GetMetadataRequest struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-" url:"-"`
}

type GetMetadataResponse struct {
	ContentLength types.Int64 `tfsdk:"-" url:"-" header:"content-length,omitempty"`

	ContentType types.String `tfsdk:"-" url:"-" header:"content-type,omitempty"`

	LastModified types.String `tfsdk:"-" url:"-" header:"last-modified,omitempty"`
}

// Get the information of a file or directory
type GetStatusRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path types.String `tfsdk:"-" url:"path"`
}

// List directory contents or file details
type ListDbfsRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path types.String `tfsdk:"-" url:"path"`
}

// List directory contents
type ListDirectoryContentsRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-" url:"-"`
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
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the contents of this directory. Provide this
	// token to retrieve the next page of directory entries. When providing a
	// `page_token`, all other parameters provided to the request must match the
	// previous request. To list all of the entries in a directory, it is
	// necessary to continue requesting pages of entries until the response
	// contains no `next_page_token`. Note that the number of entries returned
	// must not be used to determine when the listing is complete.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListDirectoryResponse struct {
	// Array of DirectoryEntry.
	Contents types.List `tfsdk:"contents"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

type ListStatusResponse struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	Files types.List `tfsdk:"files"`
}

type MkDirs struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path"`
}

type MkDirsResponse struct {
}

type Move struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path.
	DestinationPath types.String `tfsdk:"destination_path"`
	// The source path of the file or directory. The path should be the absolute
	// DBFS path.
	SourcePath types.String `tfsdk:"source_path"`
}

type MoveResponse struct {
}

type Put struct {
	// This parameter might be absent, and instead a posted file will be used.
	Contents types.String `tfsdk:"contents"`
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite types.Bool `tfsdk:"overwrite"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path"`
}

type PutResponse struct {
}

// Get the contents of a file
type ReadDbfsRequest struct {
	// The number of bytes to read starting from the offset. This has a limit of
	// 1 MB, and a default value of 0.5 MB.
	Length types.Int64 `tfsdk:"-" url:"length,omitempty"`
	// The offset to read from in bytes.
	Offset types.Int64 `tfsdk:"-" url:"offset,omitempty"`
	// The path of the file to read. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"-" url:"path"`
}

type ReadResponse struct {
	// The number of bytes read (could be less than ``length`` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	BytesRead types.Int64 `tfsdk:"bytes_read"`
	// The base64-encoded contents of the file read.
	Data types.String `tfsdk:"data"`
}

// Upload a file
type UploadRequest struct {
	Contents io.ReadCloser `tfsdk:"-"`
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-" url:"-"`
	// If true, an existing file will be overwritten.
	Overwrite types.Bool `tfsdk:"-" url:"overwrite,omitempty"`
}

type UploadResponse struct {
}
