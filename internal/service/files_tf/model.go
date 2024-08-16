// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package files_tf

import (
	"io"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddBlock struct {
	// The base64-encoded data to append to the stream. This has a limit of 1
	// MB.
	Data types.String `tfsdk:"data" tf:""`
	// The handle on an open stream.
	Handle types.Int64 `tfsdk:"handle" tf:""`
}

type AddBlockResponse struct {
}

type Close struct {
	// The handle on an open stream.
	Handle types.Int64 `tfsdk:"handle" tf:""`
}

type CloseResponse struct {
}

type Create struct {
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite types.Bool `tfsdk:"overwrite" tf:"optional"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path" tf:""`
}

// Create a directory
type CreateDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-"`
}

type CreateDirectoryResponse struct {
}

type CreateResponse struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	Handle types.Int64 `tfsdk:"handle" tf:"optional"`
}

type Delete struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path.
	Path types.String `tfsdk:"path" tf:""`
	// Whether or not to recursively delete the directory's contents. Deleting
	// empty directories can be done without providing the recursive flag.
	Recursive types.Bool `tfsdk:"recursive" tf:"optional"`
}

// Delete a directory
type DeleteDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-"`
}

type DeleteDirectoryResponse struct {
}

// Delete a file
type DeleteFileRequest struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
}

type DeleteResponse struct {
}

type DirectoryEntry struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize types.Int64 `tfsdk:"file_size" tf:"optional"`
	// True if the path is a directory.
	IsDirectory types.Bool `tfsdk:"is_directory" tf:"optional"`
	// Last modification time of given file in milliseconds since unix epoch.
	LastModified types.Int64 `tfsdk:"last_modified" tf:"optional"`
	// The name of the file or directory. This is the last component of the
	// path.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The absolute path of the file or directory.
	Path types.String `tfsdk:"path" tf:"optional"`
}

// Download a file
type DownloadRequest struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
}

type DownloadResponse struct {
	ContentLength types.Int64 `tfsdk:"-"`

	ContentType types.String `tfsdk:"-"`

	Contents io.ReadCloser `tfsdk:"-"`

	LastModified types.String `tfsdk:"-"`
}

type FileInfo struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize types.Int64 `tfsdk:"file_size" tf:"optional"`
	// True if the path is a directory.
	IsDir types.Bool `tfsdk:"is_dir" tf:"optional"`
	// Last modification time of given file in milliseconds since epoch.
	ModificationTime types.Int64 `tfsdk:"modification_time" tf:"optional"`
	// The absolute path of the file or directory.
	Path types.String `tfsdk:"path" tf:"optional"`
}

// Get directory metadata
type GetDirectoryMetadataRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-"`
}

type GetDirectoryMetadataResponse struct {
}

// Get file metadata
type GetMetadataRequest struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
}

type GetMetadataResponse struct {
	ContentLength types.Int64 `tfsdk:"-"`

	ContentType types.String `tfsdk:"-"`

	LastModified types.String `tfsdk:"-"`
}

// Get the information of a file or directory
type GetStatusRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path types.String `tfsdk:"-"`
}

// List directory contents or file details
type ListDbfsRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path types.String `tfsdk:"-"`
}

// List directory contents
type ListDirectoryContentsRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-"`
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
	PageSize types.Int64 `tfsdk:"-"`
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the contents of this directory. Provide this
	// token to retrieve the next page of directory entries. When providing a
	// `page_token`, all other parameters provided to the request must match the
	// previous request. To list all of the entries in a directory, it is
	// necessary to continue requesting pages of entries until the response
	// contains no `next_page_token`. Note that the number of entries returned
	// must not be used to determine when the listing is complete.
	PageToken types.String `tfsdk:"-"`
}

type ListDirectoryResponse struct {
	// Array of DirectoryEntry.
	Contents []DirectoryEntry `tfsdk:"contents" tf:"optional"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

type ListStatusResponse struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	Files []FileInfo `tfsdk:"files" tf:"optional"`
}

type MkDirs struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path" tf:""`
}

type MkDirsResponse struct {
}

type Move struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path.
	DestinationPath types.String `tfsdk:"destination_path" tf:""`
	// The source path of the file or directory. The path should be the absolute
	// DBFS path.
	SourcePath types.String `tfsdk:"source_path" tf:""`
}

type MoveResponse struct {
}

type Put struct {
	// This parameter might be absent, and instead a posted file will be used.
	Contents types.String `tfsdk:"contents" tf:"optional"`
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite types.Bool `tfsdk:"overwrite" tf:"optional"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path" tf:""`
}

type PutResponse struct {
}

// Get the contents of a file
type ReadDbfsRequest struct {
	// The number of bytes to read starting from the offset. This has a limit of
	// 1 MB, and a default value of 0.5 MB.
	Length types.Int64 `tfsdk:"-"`
	// The offset to read from in bytes.
	Offset types.Int64 `tfsdk:"-"`
	// The path of the file to read. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"-"`
}

type ReadResponse struct {
	// The number of bytes read (could be less than ``length`` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	BytesRead types.Int64 `tfsdk:"bytes_read" tf:"optional"`
	// The base64-encoded contents of the file read.
	Data types.String `tfsdk:"data" tf:"optional"`
}

// Upload a file
type UploadRequest struct {
	Contents io.ReadCloser `tfsdk:"-"`
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
	// If true, an existing file will be overwritten.
	Overwrite types.Bool `tfsdk:"-"`
}

type UploadResponse struct {
}
