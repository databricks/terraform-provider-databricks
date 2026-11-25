// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"context"
)

// DBFS API makes it simple to interact with various data sources without having
// to include a users credentials every time to read a file.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DbfsService interface {

	// Appends a block of data to the stream specified by the input handle. If
	// the handle does not exist, this call will throw an exception with
	// ``RESOURCE_DOES_NOT_EXIST``.
	//
	// If the block of data exceeds 1 MB, this call will throw an exception with
	// ``MAX_BLOCK_SIZE_EXCEEDED``.
	AddBlock(ctx context.Context, request AddBlock) error

	// Closes the stream specified by the input handle. If the handle does not
	// exist, this call throws an exception with ``RESOURCE_DOES_NOT_EXIST``.
	Close(ctx context.Context, request Close) error

	// Opens a stream to write to a file and returns a handle to this stream.
	// There is a 10 minute idle timeout on this handle. If a file or directory
	// already exists on the given path and __overwrite__ is set to false, this
	// call will throw an exception with ``RESOURCE_ALREADY_EXISTS``.
	//
	// A typical workflow for file upload would be:
	//
	// 1. Issue a ``create`` call and get a handle. 2. Issue one or more
	// ``add-block`` calls with the handle you have. 3. Issue a ``close`` call
	// with the handle you have.
	Create(ctx context.Context, request Create) (*CreateResponse, error)

	// Delete the file or directory (optionally recursively delete all files in
	// the directory). This call throws an exception with `IO_ERROR` if the path
	// is a non-empty directory and `recursive` is set to `false` or on other
	// similar errors.
	//
	// When you delete a large number of files, the delete operation is done in
	// increments. The call returns a response after approximately 45 seconds
	// with an error message (503 Service Unavailable) asking you to re-invoke
	// the delete operation until the directory structure is fully deleted.
	//
	// For operations that delete more than 10K files, we discourage using the
	// DBFS REST API, but advise you to perform such operations in the context
	// of a cluster, using the [File system utility
	// (dbutils.fs)](/dev-tools/databricks-utils.html#dbutils-fs). `dbutils.fs`
	// covers the functional scope of the DBFS REST API, but from notebooks.
	// Running such operations using notebooks provides better control and
	// manageability, such as selective deletes, and the possibility to automate
	// periodic delete jobs.
	Delete(ctx context.Context, request Delete) error

	// Gets the file information for a file or directory. If the file or
	// directory does not exist, this call throws an exception with
	// `RESOURCE_DOES_NOT_EXIST`.
	GetStatus(ctx context.Context, request GetStatusRequest) (*FileInfo, error)

	// List the contents of a directory, or details of the file. If the file or
	// directory does not exist, this call throws an exception with
	// `RESOURCE_DOES_NOT_EXIST`.
	//
	// When calling list on a large directory, the list operation will time out
	// after approximately 60 seconds. We strongly recommend using list only on
	// directories containing less than 10K files and discourage using the DBFS
	// REST API for operations that list more than 10K files. Instead, we
	// recommend that you perform such operations in the context of a cluster,
	// using the [File system utility
	// (dbutils.fs)](/dev-tools/databricks-utils.html#dbutils-fs), which
	// provides the same functionality without timing out.
	List(ctx context.Context, request ListDbfsRequest) (*ListStatusResponse, error)

	// Creates the given directory and necessary parent directories if they do
	// not exist. If a file (not a directory) exists at any prefix of the input
	// path, this call throws an exception with `RESOURCE_ALREADY_EXISTS`.
	// **Note**: If this operation fails, it might have succeeded in creating
	// some of the necessary parent directories.
	Mkdirs(ctx context.Context, request MkDirs) error

	// Moves a file from one location to another location within DBFS. If the
	// source file does not exist, this call throws an exception with
	// `RESOURCE_DOES_NOT_EXIST`. If a file already exists in the destination
	// path, this call throws an exception with `RESOURCE_ALREADY_EXISTS`. If
	// the given source path is a directory, this call always recursively moves
	// all files.
	Move(ctx context.Context, request Move) error

	// Uploads a file through the use of multipart form post. It is mainly used
	// for streaming uploads, but can also be used as a convenient single call
	// for data upload.
	//
	// Alternatively you can pass contents as base64 string.
	//
	// The amount of data that can be passed (when not streaming) using the
	// __contents__ parameter is limited to 1 MB. `MAX_BLOCK_SIZE_EXCEEDED` will
	// be thrown if this limit is exceeded.
	//
	// If you want to upload large files, use the streaming upload. For details,
	// see :method:dbfs/create, :method:dbfs/addBlock, :method:dbfs/close.
	Put(ctx context.Context, request Put) error

	// Returns the contents of a file. If the file does not exist, this call
	// throws an exception with `RESOURCE_DOES_NOT_EXIST`. If the path is a
	// directory, the read length is negative, or if the offset is negative,
	// this call throws an exception with `INVALID_PARAMETER_VALUE`. If the read
	// length exceeds 1 MB, this call throws an exception with
	// `MAX_READ_SIZE_EXCEEDED`.
	//
	// If `offset + length` exceeds the number of bytes in a file, it reads the
	// contents until the end of file.
	Read(ctx context.Context, request ReadDbfsRequest) (*ReadResponse, error)
}

// The Files API is a standard HTTP API that allows you to read, write, list,
// and delete files and directories by referring to their URI. The API makes
// working with file content as raw bytes easier and more efficient.
//
// The API supports [Unity Catalog volumes], where files and directories to
// operate on are specified using their volume URI path, which follows the
// format
// /Volumes/&lt;catalog_name&gt;/&lt;schema_name&gt;/&lt;volume_name&gt;/&lt;path_to_file&gt;.
//
// The Files API has two distinct endpoints, one for working with files
// (`/fs/files`) and another one for working with directories
// (`/fs/directories`). Both endpoints use the standard HTTP methods GET, HEAD,
// PUT, and DELETE to manage files and directories specified using their URI
// path. The path is always absolute.
//
// Use of Files API may incur Databricks data transfer charges.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Unity Catalog volumes]: https://docs.databricks.com/en/connect/unity-catalog/volumes.html
type FilesService interface {

	// Creates an empty directory. If necessary, also creates any parent
	// directories of the new, empty directory (like the shell command `mkdir
	// -p`). If called on an existing directory, returns a success response;
	// this method is idempotent (it will succeed if the directory already
	// exists).
	CreateDirectory(ctx context.Context, request CreateDirectoryRequest) error

	// Deletes a file. If the request is successful, there is no response body.
	Delete(ctx context.Context, request DeleteFileRequest) error

	// Deletes an empty directory.
	//
	// To delete a non-empty directory, first delete all of its contents. This
	// can be done by listing the directory contents and deleting each file and
	// subdirectory recursively.
	DeleteDirectory(ctx context.Context, request DeleteDirectoryRequest) error

	// Downloads a file. The file contents are the response body. This is a
	// standard HTTP file download, not a JSON RPC. It supports the Range and
	// If-Unmodified-Since HTTP headers.
	Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error)

	// Get the metadata of a directory. The response HTTP headers contain the
	// metadata. There is no response body.
	//
	// This method is useful to check if a directory exists and the caller has
	// access to it.
	//
	// If you wish to ensure the directory exists, you can instead use `PUT`,
	// which will create the directory if it does not exist, and is idempotent
	// (it will succeed if the directory already exists).
	GetDirectoryMetadata(ctx context.Context, request GetDirectoryMetadataRequest) error

	// Get the metadata of a file. The response HTTP headers contain the
	// metadata. There is no response body.
	GetMetadata(ctx context.Context, request GetMetadataRequest) (*GetMetadataResponse, error)

	// Returns the contents of a directory. If there is no directory at the
	// specified path, the API returns an HTTP 404 error.
	ListDirectoryContents(ctx context.Context, request ListDirectoryContentsRequest) (*ListDirectoryResponse, error)

	// Uploads a file of up to 5 GiB. The file contents should be sent as the
	// request body as raw bytes (an octet stream); do not encode or otherwise
	// modify the bytes before sending. The contents of the resulting file will
	// be exactly the bytes sent in the request body. If the request is
	// successful, there is no response body.
	Upload(ctx context.Context, request UploadRequest) error
}
