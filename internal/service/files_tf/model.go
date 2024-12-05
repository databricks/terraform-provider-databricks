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
	"context"
	"io"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddBlock.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AddBlock) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of AddBlock in the Terraform plugin framework type
// system.
func (a AddBlock) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data":   types.StringType,
			"handle": types.Int64Type,
		},
	}
}

type AddBlockResponse struct {
}

func (newState *AddBlockResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AddBlockResponse) {
}

func (newState *AddBlockResponse) SyncEffectiveFieldsDuringRead(existingState AddBlockResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddBlockResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AddBlockResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of AddBlockResponse in the Terraform plugin framework type
// system.
func (a AddBlockResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Close struct {
	// The handle on an open stream.
	Handle types.Int64 `tfsdk:"handle" tf:""`
}

func (newState *Close) SyncEffectiveFieldsDuringCreateOrUpdate(plan Close) {
}

func (newState *Close) SyncEffectiveFieldsDuringRead(existingState Close) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Close.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Close) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of Close in the Terraform plugin framework type
// system.
func (a Close) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"handle": types.Int64Type,
		},
	}
}

type CloseResponse struct {
}

func (newState *CloseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloseResponse) {
}

func (newState *CloseResponse) SyncEffectiveFieldsDuringRead(existingState CloseResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CloseResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of CloseResponse in the Terraform plugin framework type
// system.
func (a CloseResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Create.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Create) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of Create in the Terraform plugin framework type
// system.
func (a Create) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"overwrite": types.BoolType,
			"path":      types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDirectoryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDirectoryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of CreateDirectoryRequest in the Terraform plugin framework type
// system.
func (a CreateDirectoryRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"directory_path": types.StringType,
		},
	}
}

type CreateDirectoryResponse struct {
}

func (newState *CreateDirectoryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateDirectoryResponse) {
}

func (newState *CreateDirectoryResponse) SyncEffectiveFieldsDuringRead(existingState CreateDirectoryResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDirectoryResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDirectoryResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of CreateDirectoryResponse in the Terraform plugin framework type
// system.
func (a CreateDirectoryResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of CreateResponse in the Terraform plugin framework type
// system.
func (a CreateResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"handle": types.Int64Type,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Delete.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Delete) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of Delete in the Terraform plugin framework type
// system.
func (a Delete) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":      types.StringType,
			"recursive": types.BoolType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDirectoryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDirectoryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteDirectoryRequest in the Terraform plugin framework type
// system.
func (a DeleteDirectoryRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"directory_path": types.StringType,
		},
	}
}

type DeleteDirectoryResponse struct {
}

func (newState *DeleteDirectoryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDirectoryResponse) {
}

func (newState *DeleteDirectoryResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDirectoryResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDirectoryResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDirectoryResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteDirectoryResponse in the Terraform plugin framework type
// system.
func (a DeleteDirectoryResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFileRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteFileRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteFileRequest in the Terraform plugin framework type
// system.
func (a DeleteFileRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_path": types.StringType,
		},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DeleteResponse in the Terraform plugin framework type
// system.
func (a DeleteResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DirectoryEntry.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DirectoryEntry) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DirectoryEntry in the Terraform plugin framework type
// system.
func (a DirectoryEntry) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_size":     types.Int64Type,
			"is_directory":  types.BoolType,
			"last_modified": types.Int64Type,
			"name":          types.StringType,
			"path":          types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DownloadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DownloadRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DownloadRequest in the Terraform plugin framework type
// system.
func (a DownloadRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_path": types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DownloadResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DownloadResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of DownloadResponse in the Terraform plugin framework type
// system.
func (a DownloadResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content-length": types.Int64Type,
			"content-type":   types.StringType,
			"contents":       types.ObjectType{},
			"last-modified":  types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FileInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of FileInfo in the Terraform plugin framework type
// system.
func (a FileInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_size":         types.Int64Type,
			"is_dir":            types.BoolType,
			"modification_time": types.Int64Type,
			"path":              types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDirectoryMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDirectoryMetadataRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetDirectoryMetadataRequest in the Terraform plugin framework type
// system.
func (a GetDirectoryMetadataRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"directory_path": types.StringType,
		},
	}
}

type GetDirectoryMetadataResponse struct {
}

func (newState *GetDirectoryMetadataResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDirectoryMetadataResponse) {
}

func (newState *GetDirectoryMetadataResponse) SyncEffectiveFieldsDuringRead(existingState GetDirectoryMetadataResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDirectoryMetadataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDirectoryMetadataResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetDirectoryMetadataResponse in the Terraform plugin framework type
// system.
func (a GetDirectoryMetadataResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetMetadataRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetMetadataRequest in the Terraform plugin framework type
// system.
func (a GetMetadataRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_path": types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMetadataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetMetadataResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetMetadataResponse in the Terraform plugin framework type
// system.
func (a GetMetadataResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content-length": types.Int64Type,
			"content-type":   types.StringType,
			"last-modified":  types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GetStatusRequest in the Terraform plugin framework type
// system.
func (a GetStatusRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ListDbfsRequest in the Terraform plugin framework type
// system.
func (a ListDbfsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDirectoryContentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDirectoryContentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ListDirectoryContentsRequest in the Terraform plugin framework type
// system.
func (a ListDirectoryContentsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"directory_path": types.StringType,
			"page_size":      types.Int64Type,
			"page_token":     types.StringType,
		},
	}
}

type ListDirectoryResponse struct {
	// Array of DirectoryEntry.
	Contents types.List `tfsdk:"contents" tf:"optional"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListDirectoryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDirectoryResponse) {
}

func (newState *ListDirectoryResponse) SyncEffectiveFieldsDuringRead(existingState ListDirectoryResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDirectoryResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDirectoryResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"contents": reflect.TypeOf(DirectoryEntry{}),
	}
}

// ToAttrType returns the representation of ListDirectoryResponse in the Terraform plugin framework type
// system.
func (a ListDirectoryResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": basetypes.ListType{
				ElemType: DirectoryEntry{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

type ListStatusResponse struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	Files types.List `tfsdk:"files" tf:"optional"`
}

func (newState *ListStatusResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListStatusResponse) {
}

func (newState *ListStatusResponse) SyncEffectiveFieldsDuringRead(existingState ListStatusResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStatusResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListStatusResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"files": reflect.TypeOf(FileInfo{}),
	}
}

// ToAttrType returns the representation of ListStatusResponse in the Terraform plugin framework type
// system.
func (a ListStatusResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"files": basetypes.ListType{
				ElemType: FileInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type MkDirs struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path" tf:""`
}

func (newState *MkDirs) SyncEffectiveFieldsDuringCreateOrUpdate(plan MkDirs) {
}

func (newState *MkDirs) SyncEffectiveFieldsDuringRead(existingState MkDirs) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MkDirs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MkDirs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of MkDirs in the Terraform plugin framework type
// system.
func (a MkDirs) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type MkDirsResponse struct {
}

func (newState *MkDirsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MkDirsResponse) {
}

func (newState *MkDirsResponse) SyncEffectiveFieldsDuringRead(existingState MkDirsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MkDirsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MkDirsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of MkDirsResponse in the Terraform plugin framework type
// system.
func (a MkDirsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Move.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Move) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of Move in the Terraform plugin framework type
// system.
func (a Move) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_path": types.StringType,
			"source_path":      types.StringType,
		},
	}
}

type MoveResponse struct {
}

func (newState *MoveResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MoveResponse) {
}

func (newState *MoveResponse) SyncEffectiveFieldsDuringRead(existingState MoveResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MoveResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MoveResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of MoveResponse in the Terraform plugin framework type
// system.
func (a MoveResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Put.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Put) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of Put in the Terraform plugin framework type
// system.
func (a Put) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents":  types.StringType,
			"overwrite": types.BoolType,
			"path":      types.StringType,
		},
	}
}

type PutResponse struct {
}

func (newState *PutResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutResponse) {
}

func (newState *PutResponse) SyncEffectiveFieldsDuringRead(existingState PutResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of PutResponse in the Terraform plugin framework type
// system.
func (a PutResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReadDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReadDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ReadDbfsRequest in the Terraform plugin framework type
// system.
func (a ReadDbfsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"length": types.Int64Type,
			"offset": types.Int64Type,
			"path":   types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReadResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReadResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of ReadResponse in the Terraform plugin framework type
// system.
func (a ReadResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bytes_read": types.Int64Type,
			"data":       types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UploadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UploadRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of UploadRequest in the Terraform plugin framework type
// system.
func (a UploadRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents":  types.ObjectType{},
			"file_path": types.StringType,
			"overwrite": types.BoolType,
		},
	}
}

type UploadResponse struct {
}

func (newState *UploadResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UploadResponse) {
}

func (newState *UploadResponse) SyncEffectiveFieldsDuringRead(existingState UploadResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UploadResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UploadResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of UploadResponse in the Terraform plugin framework type
// system.
func (a UploadResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}
