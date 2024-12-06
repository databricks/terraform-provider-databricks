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

func (newState *AddBlock) SyncEffectiveFieldsDuringCreateOrUpdate(plan AddBlock) {
}

func (newState *AddBlock) SyncEffectiveFieldsDuringRead(existingState AddBlock) {
}

type AddBlockResponse struct {
}

func (newState *AddBlockResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AddBlockResponse) {
}

func (newState *AddBlockResponse) SyncEffectiveFieldsDuringRead(existingState AddBlockResponse) {
}

type Close struct {
	// The handle on an open stream.
	Handle types.Int64 `tfsdk:"handle" tf:""`
}

func (newState *Close) SyncEffectiveFieldsDuringCreateOrUpdate(plan Close) {
}

func (newState *Close) SyncEffectiveFieldsDuringRead(existingState Close) {
}

type CloseResponse struct {
}

func (newState *CloseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloseResponse) {
}

func (newState *CloseResponse) SyncEffectiveFieldsDuringRead(existingState CloseResponse) {
}

type Create struct {
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite types.Bool `tfsdk:"overwrite" tf:"optional"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path" tf:""`
}

func (newState *Create) SyncEffectiveFieldsDuringCreateOrUpdate(plan Create) {
}

func (newState *Create) SyncEffectiveFieldsDuringRead(existingState Create) {
}

// Create a directory
type CreateDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-"`
}

func (newState *CreateDirectoryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateDirectoryRequest) {
}

func (newState *CreateDirectoryRequest) SyncEffectiveFieldsDuringRead(existingState CreateDirectoryRequest) {
}

type CreateDirectoryResponse struct {
}

func (newState *CreateDirectoryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateDirectoryResponse) {
}

func (newState *CreateDirectoryResponse) SyncEffectiveFieldsDuringRead(existingState CreateDirectoryResponse) {
}

type CreateResponse struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	Handle types.Int64 `tfsdk:"handle" tf:"optional"`
}

func (newState *CreateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateResponse) {
}

func (newState *CreateResponse) SyncEffectiveFieldsDuringRead(existingState CreateResponse) {
}

type Delete struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path.
	Path types.String `tfsdk:"path" tf:""`
	// Whether or not to recursively delete the directory's contents. Deleting
	// empty directories can be done without providing the recursive flag.
	Recursive types.Bool `tfsdk:"recursive" tf:"optional"`
}

func (newState *Delete) SyncEffectiveFieldsDuringCreateOrUpdate(plan Delete) {
}

func (newState *Delete) SyncEffectiveFieldsDuringRead(existingState Delete) {
}

// Delete a directory
type DeleteDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-"`
}

func (newState *DeleteDirectoryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDirectoryRequest) {
}

func (newState *DeleteDirectoryRequest) SyncEffectiveFieldsDuringRead(existingState DeleteDirectoryRequest) {
}

type DeleteDirectoryResponse struct {
}

func (newState *DeleteDirectoryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDirectoryResponse) {
}

func (newState *DeleteDirectoryResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDirectoryResponse) {
}

// Delete a file
type DeleteFileRequest struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
}

func (newState *DeleteFileRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteFileRequest) {
}

func (newState *DeleteFileRequest) SyncEffectiveFieldsDuringRead(existingState DeleteFileRequest) {
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
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

func (newState *DirectoryEntry) SyncEffectiveFieldsDuringCreateOrUpdate(plan DirectoryEntry) {
}

func (newState *DirectoryEntry) SyncEffectiveFieldsDuringRead(existingState DirectoryEntry) {
}

// Download a file
type DownloadRequest struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
}

func (newState *DownloadRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DownloadRequest) {
}

func (newState *DownloadRequest) SyncEffectiveFieldsDuringRead(existingState DownloadRequest) {
}

type DownloadResponse struct {
	ContentLength types.Int64 `tfsdk:"-"`

	ContentType types.String `tfsdk:"-"`

	Contents io.ReadCloser `tfsdk:"-"`

	LastModified types.String `tfsdk:"-"`
}

func (newState *DownloadResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DownloadResponse) {
}

func (newState *DownloadResponse) SyncEffectiveFieldsDuringRead(existingState DownloadResponse) {
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

func (newState *FileInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileInfo) {
}

func (newState *FileInfo) SyncEffectiveFieldsDuringRead(existingState FileInfo) {
}

// Get directory metadata
type GetDirectoryMetadataRequest struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-"`
}

func (newState *GetDirectoryMetadataRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDirectoryMetadataRequest) {
}

func (newState *GetDirectoryMetadataRequest) SyncEffectiveFieldsDuringRead(existingState GetDirectoryMetadataRequest) {
}

type GetDirectoryMetadataResponse struct {
}

func (newState *GetDirectoryMetadataResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDirectoryMetadataResponse) {
}

func (newState *GetDirectoryMetadataResponse) SyncEffectiveFieldsDuringRead(existingState GetDirectoryMetadataResponse) {
}

// Get file metadata
type GetMetadataRequest struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
}

func (newState *GetMetadataRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetadataRequest) {
}

func (newState *GetMetadataRequest) SyncEffectiveFieldsDuringRead(existingState GetMetadataRequest) {
}

type GetMetadataResponse struct {
	ContentLength types.Int64 `tfsdk:"-"`

	ContentType types.String `tfsdk:"-"`

	LastModified types.String `tfsdk:"-"`
}

func (newState *GetMetadataResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetadataResponse) {
}

func (newState *GetMetadataResponse) SyncEffectiveFieldsDuringRead(existingState GetMetadataResponse) {
}

// Get the information of a file or directory
type GetStatusRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path types.String `tfsdk:"-"`
}

func (newState *GetStatusRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStatusRequest) {
}

func (newState *GetStatusRequest) SyncEffectiveFieldsDuringRead(existingState GetStatusRequest) {
}

// List directory contents or file details
type ListDbfsRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path types.String `tfsdk:"-"`
}

func (newState *ListDbfsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDbfsRequest) {
}

func (newState *ListDbfsRequest) SyncEffectiveFieldsDuringRead(existingState ListDbfsRequest) {
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

func (newState *ListDirectoryContentsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDirectoryContentsRequest) {
}

func (newState *ListDirectoryContentsRequest) SyncEffectiveFieldsDuringRead(existingState ListDirectoryContentsRequest) {
}

type ListDirectoryResponse struct {
	// Array of DirectoryEntry.
	Contents []DirectoryEntry `tfsdk:"contents" tf:"optional"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListDirectoryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDirectoryResponse) {
}

func (newState *ListDirectoryResponse) SyncEffectiveFieldsDuringRead(existingState ListDirectoryResponse) {
}

type ListStatusResponse struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	Files []FileInfo `tfsdk:"files" tf:"optional"`
}

func (newState *ListStatusResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListStatusResponse) {
}

func (newState *ListStatusResponse) SyncEffectiveFieldsDuringRead(existingState ListStatusResponse) {
}

type MkDirs struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path" tf:""`
}

func (newState *MkDirs) SyncEffectiveFieldsDuringCreateOrUpdate(plan MkDirs) {
}

func (newState *MkDirs) SyncEffectiveFieldsDuringRead(existingState MkDirs) {
}

type MkDirsResponse struct {
}

func (newState *MkDirsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MkDirsResponse) {
}

func (newState *MkDirsResponse) SyncEffectiveFieldsDuringRead(existingState MkDirsResponse) {
}

type Move struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path.
	DestinationPath types.String `tfsdk:"destination_path" tf:""`
	// The source path of the file or directory. The path should be the absolute
	// DBFS path.
	SourcePath types.String `tfsdk:"source_path" tf:""`
}

func (newState *Move) SyncEffectiveFieldsDuringCreateOrUpdate(plan Move) {
}

func (newState *Move) SyncEffectiveFieldsDuringRead(existingState Move) {
}

type MoveResponse struct {
}

func (newState *MoveResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MoveResponse) {
}

func (newState *MoveResponse) SyncEffectiveFieldsDuringRead(existingState MoveResponse) {
}

type Put struct {
	// This parameter might be absent, and instead a posted file will be used.
	Contents types.String `tfsdk:"contents" tf:"optional"`
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite types.Bool `tfsdk:"overwrite" tf:"optional"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path" tf:""`
}

func (newState *Put) SyncEffectiveFieldsDuringCreateOrUpdate(plan Put) {
}

func (newState *Put) SyncEffectiveFieldsDuringRead(existingState Put) {
}

type PutResponse struct {
}

func (newState *PutResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutResponse) {
}

func (newState *PutResponse) SyncEffectiveFieldsDuringRead(existingState PutResponse) {
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

func (newState *ReadDbfsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReadDbfsRequest) {
}

func (newState *ReadDbfsRequest) SyncEffectiveFieldsDuringRead(existingState ReadDbfsRequest) {
}

type ReadResponse struct {
	// The number of bytes read (could be less than ``length`` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	BytesRead types.Int64 `tfsdk:"bytes_read" tf:"optional"`
	// The base64-encoded contents of the file read.
	Data types.String `tfsdk:"data" tf:"optional"`
}

func (newState *ReadResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReadResponse) {
}

func (newState *ReadResponse) SyncEffectiveFieldsDuringRead(existingState ReadResponse) {
}

// Upload a file
type UploadRequest struct {
	Contents io.ReadCloser `tfsdk:"-"`
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
	// If true, an existing file will be overwritten.
	Overwrite types.Bool `tfsdk:"-"`
}

func (newState *UploadRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UploadRequest) {
}

func (newState *UploadRequest) SyncEffectiveFieldsDuringRead(existingState UploadRequest) {
}

type UploadResponse struct {
}

func (newState *UploadResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UploadResponse) {
}

func (newState *UploadResponse) SyncEffectiveFieldsDuringRead(existingState UploadResponse) {
}
