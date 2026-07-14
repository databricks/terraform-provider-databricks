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
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AddBlock_SdkV2 struct {
	// The base64-encoded data to append to the stream. This has a limit of 1
	// MB.
	Data types.String `tfsdk:"data"`
	// The handle on an open stream.
	Handle types.Int64 `tfsdk:"handle"`
}

func (to *AddBlock_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AddBlock_SdkV2) {
}

func (to *AddBlock_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AddBlock_SdkV2) {
}

func (m AddBlock_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data"] = attrs["data"].SetRequired()
	attrs["handle"] = attrs["handle"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddBlock.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AddBlock_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddBlock_SdkV2
// only implements ToObjectValue() and Type().
func (m AddBlock_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":   m.Data,
			"handle": m.Handle,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AddBlock_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data":   types.StringType,
			"handle": types.Int64Type,
		},
	}
}

type AddBlockResponse_SdkV2 struct {
}

func (to *AddBlockResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AddBlockResponse_SdkV2) {
}

func (to *AddBlockResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AddBlockResponse_SdkV2) {
}

func (m AddBlockResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddBlockResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AddBlockResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddBlockResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m AddBlockResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m AddBlockResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Close_SdkV2 struct {
	// The handle on an open stream.
	Handle types.Int64 `tfsdk:"handle"`
}

func (to *Close_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Close_SdkV2) {
}

func (to *Close_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Close_SdkV2) {
}

func (m Close_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["handle"] = attrs["handle"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Close.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Close_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Close_SdkV2
// only implements ToObjectValue() and Type().
func (m Close_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"handle": m.Handle,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Close_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"handle": types.Int64Type,
		},
	}
}

type CloseResponse_SdkV2 struct {
}

func (to *CloseResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CloseResponse_SdkV2) {
}

func (to *CloseResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CloseResponse_SdkV2) {
}

func (m CloseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CloseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CloseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m CloseResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Create_SdkV2 struct {
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite types.Bool `tfsdk:"overwrite"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path"`
}

func (to *Create_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Create_SdkV2) {
}

func (to *Create_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Create_SdkV2) {
}

func (m Create_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["overwrite"] = attrs["overwrite"].SetOptional()
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Create.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Create_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Create_SdkV2
// only implements ToObjectValue() and Type().
func (m Create_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"overwrite": m.Overwrite,
			"path":      m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Create_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"overwrite": types.BoolType,
			"path":      types.StringType,
		},
	}
}

type CreateDirectoryRequest_SdkV2 struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-"`
}

func (to *CreateDirectoryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDirectoryRequest_SdkV2) {
}

func (to *CreateDirectoryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateDirectoryRequest_SdkV2) {
}

func (m CreateDirectoryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["directory_path"] = attrs["directory_path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDirectoryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDirectoryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDirectoryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateDirectoryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"directory_path": m.DirectoryPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDirectoryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"directory_path": types.StringType,
		},
	}
}

type CreateResponse_SdkV2 struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	Handle types.Int64 `tfsdk:"handle"`
}

func (to *CreateResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateResponse_SdkV2) {
}

func (to *CreateResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateResponse_SdkV2) {
}

func (m CreateResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["handle"] = attrs["handle"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"handle": m.Handle,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"handle": types.Int64Type,
		},
	}
}

type Delete_SdkV2 struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path.
	Path types.String `tfsdk:"path"`
	// Whether or not to recursively delete the directory's contents. Deleting
	// empty directories can be done without providing the recursive flag.
	Recursive types.Bool `tfsdk:"recursive"`
}

func (to *Delete_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Delete_SdkV2) {
}

func (to *Delete_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Delete_SdkV2) {
}

func (m Delete_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()
	attrs["recursive"] = attrs["recursive"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Delete.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Delete_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Delete_SdkV2
// only implements ToObjectValue() and Type().
func (m Delete_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path":      m.Path,
			"recursive": m.Recursive,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Delete_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":      types.StringType,
			"recursive": types.BoolType,
		},
	}
}

type DeleteDirectoryRequest_SdkV2 struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-"`
}

func (to *DeleteDirectoryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDirectoryRequest_SdkV2) {
}

func (to *DeleteDirectoryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDirectoryRequest_SdkV2) {
}

func (m DeleteDirectoryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["directory_path"] = attrs["directory_path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDirectoryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDirectoryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDirectoryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDirectoryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"directory_path": m.DirectoryPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDirectoryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"directory_path": types.StringType,
		},
	}
}

type DeleteFileRequest_SdkV2 struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
}

func (to *DeleteFileRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFileRequest_SdkV2) {
}

func (to *DeleteFileRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteFileRequest_SdkV2) {
}

func (m DeleteFileRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_path"] = attrs["file_path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFileRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteFileRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFileRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteFileRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_path": m.FilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteFileRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_path": types.StringType,
		},
	}
}

type DeleteResponse_SdkV2 struct {
}

func (to *DeleteResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteResponse_SdkV2) {
}

func (to *DeleteResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteResponse_SdkV2) {
}

func (m DeleteResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DirectoryEntry_SdkV2 struct {
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

func (to *DirectoryEntry_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DirectoryEntry_SdkV2) {
}

func (to *DirectoryEntry_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DirectoryEntry_SdkV2) {
}

func (m DirectoryEntry_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_size"] = attrs["file_size"].SetOptional()
	attrs["is_directory"] = attrs["is_directory"].SetOptional()
	attrs["last_modified"] = attrs["last_modified"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DirectoryEntry.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DirectoryEntry_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectoryEntry_SdkV2
// only implements ToObjectValue() and Type().
func (m DirectoryEntry_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_size":     m.FileSize,
			"is_directory":  m.IsDirectory,
			"last_modified": m.LastModified,
			"name":          m.Name,
			"path":          m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DirectoryEntry_SdkV2) Type(ctx context.Context) attr.Type {
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

type DownloadRequest_SdkV2 struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
}

func (to *DownloadRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DownloadRequest_SdkV2) {
}

func (to *DownloadRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DownloadRequest_SdkV2) {
}

func (m DownloadRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_path"] = attrs["file_path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DownloadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DownloadRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DownloadRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DownloadRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_path": m.FilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DownloadRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_path": types.StringType,
		},
	}
}

type DownloadResponse_SdkV2 struct {
	// The length of the HTTP response body in bytes.
	ContentLength types.Int64 `tfsdk:"-"`

	ContentType types.String `tfsdk:"-"`

	Contents types.Object `tfsdk:"-"`
	// The last modified time of the file in HTTP-date (RFC 7231) format.
	LastModified types.String `tfsdk:"-"`
}

func (to *DownloadResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DownloadResponse_SdkV2) {
}

func (to *DownloadResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DownloadResponse_SdkV2) {
}

func (m DownloadResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["contents"] = attrs["contents"].SetOptional()
	attrs["content_length"] = attrs["content_length"].SetOptional()
	attrs["content_type"] = attrs["content_type"].SetOptional()
	attrs["last_modified"] = attrs["last_modified"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DownloadResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DownloadResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DownloadResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DownloadResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content_length": m.ContentLength,
			"content_type":   m.ContentType,
			"contents":       m.Contents,
			"last_modified":  m.LastModified,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DownloadResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content_length": types.Int64Type,
			"content_type":   types.StringType,
			"contents":       types.ObjectType{},
			"last_modified":  types.StringType,
		},
	}
}

type FileInfo_SdkV2 struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize types.Int64 `tfsdk:"file_size"`
	// True if the path is a directory.
	IsDir types.Bool `tfsdk:"is_dir"`
	// Last modification time of given file in milliseconds since epoch.
	ModificationTime types.Int64 `tfsdk:"modification_time"`
	// The absolute path of the file or directory.
	Path types.String `tfsdk:"path"`
}

func (to *FileInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileInfo_SdkV2) {
}

func (to *FileInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FileInfo_SdkV2) {
}

func (m FileInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_size"] = attrs["file_size"].SetOptional()
	attrs["is_dir"] = attrs["is_dir"].SetOptional()
	attrs["modification_time"] = attrs["modification_time"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FileInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m FileInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_size":         m.FileSize,
			"is_dir":            m.IsDir,
			"modification_time": m.ModificationTime,
			"path":              m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FileInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_size":         types.Int64Type,
			"is_dir":            types.BoolType,
			"modification_time": types.Int64Type,
			"path":              types.StringType,
		},
	}
}

type GetDirectoryMetadataRequest_SdkV2 struct {
	// The absolute path of a directory.
	DirectoryPath types.String `tfsdk:"-"`
}

func (to *GetDirectoryMetadataRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDirectoryMetadataRequest_SdkV2) {
}

func (to *GetDirectoryMetadataRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDirectoryMetadataRequest_SdkV2) {
}

func (m GetDirectoryMetadataRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["directory_path"] = attrs["directory_path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDirectoryMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDirectoryMetadataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDirectoryMetadataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDirectoryMetadataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"directory_path": m.DirectoryPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDirectoryMetadataRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"directory_path": types.StringType,
		},
	}
}

type GetMetadataRequest_SdkV2 struct {
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
}

func (to *GetMetadataRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetMetadataRequest_SdkV2) {
}

func (to *GetMetadataRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetMetadataRequest_SdkV2) {
}

func (m GetMetadataRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_path"] = attrs["file_path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetMetadataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetadataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetMetadataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_path": m.FilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetMetadataRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_path": types.StringType,
		},
	}
}

type GetMetadataResponse_SdkV2 struct {
	// The length of the HTTP response body in bytes.
	ContentLength types.Int64 `tfsdk:"-"`

	ContentType types.String `tfsdk:"-"`
	// The last modified time of the file in HTTP-date (RFC 7231) format.
	LastModified types.String `tfsdk:"-"`
}

func (to *GetMetadataResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetMetadataResponse_SdkV2) {
}

func (to *GetMetadataResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetMetadataResponse_SdkV2) {
}

func (m GetMetadataResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content_length"] = attrs["content_length"].SetOptional()
	attrs["content_type"] = attrs["content_type"].SetOptional()
	attrs["last_modified"] = attrs["last_modified"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMetadataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetMetadataResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetadataResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetMetadataResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content_length": m.ContentLength,
			"content_type":   m.ContentType,
			"last_modified":  m.LastModified,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetMetadataResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content_length": types.Int64Type,
			"content_type":   types.StringType,
			"last_modified":  types.StringType,
		},
	}
}

type GetStatusRequest_SdkV2 struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path types.String `tfsdk:"-"`
}

func (to *GetStatusRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatusRequest_SdkV2) {
}

func (to *GetStatusRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetStatusRequest_SdkV2) {
}

func (m GetStatusRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetStatusRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatusRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetStatusRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStatusRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type ListDbfsRequest_SdkV2 struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path types.String `tfsdk:"-"`
}

func (to *ListDbfsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDbfsRequest_SdkV2) {
}

func (to *ListDbfsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDbfsRequest_SdkV2) {
}

func (m ListDbfsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDbfsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDbfsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDbfsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDbfsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type ListDirectoryContentsRequest_SdkV2 struct {
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

func (to *ListDirectoryContentsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDirectoryContentsRequest_SdkV2) {
}

func (to *ListDirectoryContentsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDirectoryContentsRequest_SdkV2) {
}

func (m ListDirectoryContentsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["directory_path"] = attrs["directory_path"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDirectoryContentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDirectoryContentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDirectoryContentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDirectoryContentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"directory_path": m.DirectoryPath,
			"page_size":      m.PageSize,
			"page_token":     m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDirectoryContentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"directory_path": types.StringType,
			"page_size":      types.Int64Type,
			"page_token":     types.StringType,
		},
	}
}

type ListDirectoryResponse_SdkV2 struct {
	// Array of DirectoryEntry.
	Contents types.List `tfsdk:"contents"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDirectoryResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDirectoryResponse_SdkV2) {
	if !from.Contents.IsNull() && !from.Contents.IsUnknown() && to.Contents.IsNull() && len(from.Contents.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Contents, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Contents = from.Contents
	}
}

func (to *ListDirectoryResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDirectoryResponse_SdkV2) {
	if !from.Contents.IsNull() && !from.Contents.IsUnknown() && to.Contents.IsNull() && len(from.Contents.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Contents, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Contents = from.Contents
	}
}

func (m ListDirectoryResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["contents"] = attrs["contents"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDirectoryResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDirectoryResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"contents": reflect.TypeOf(DirectoryEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDirectoryResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDirectoryResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents":        m.Contents,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDirectoryResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": basetypes.ListType{
				ElemType: DirectoryEntry_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetContents returns the value of the Contents field in ListDirectoryResponse_SdkV2 as
// a slice of DirectoryEntry_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDirectoryResponse_SdkV2) GetContents(ctx context.Context) ([]DirectoryEntry_SdkV2, bool) {
	if m.Contents.IsNull() || m.Contents.IsUnknown() {
		return nil, false
	}
	var v []DirectoryEntry_SdkV2
	d := m.Contents.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetContents sets the value of the Contents field in ListDirectoryResponse_SdkV2.
func (m *ListDirectoryResponse_SdkV2) SetContents(ctx context.Context, v []DirectoryEntry_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["contents"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Contents = types.ListValueMust(t, vs)
}

type ListStatusResponse_SdkV2 struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	Files types.List `tfsdk:"files"`
}

func (to *ListStatusResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListStatusResponse_SdkV2) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
}

func (to *ListStatusResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListStatusResponse_SdkV2) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
}

func (m ListStatusResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["files"] = attrs["files"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStatusResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListStatusResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"files": reflect.TypeOf(FileInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStatusResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListStatusResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"files": m.Files,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListStatusResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"files": basetypes.ListType{
				ElemType: FileInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFiles returns the value of the Files field in ListStatusResponse_SdkV2 as
// a slice of FileInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListStatusResponse_SdkV2) GetFiles(ctx context.Context) ([]FileInfo_SdkV2, bool) {
	if m.Files.IsNull() || m.Files.IsUnknown() {
		return nil, false
	}
	var v []FileInfo_SdkV2
	d := m.Files.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in ListStatusResponse_SdkV2.
func (m *ListStatusResponse_SdkV2) SetFiles(ctx context.Context, v []FileInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["files"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Files = types.ListValueMust(t, vs)
}

type MkDirs_SdkV2 struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path"`
}

func (to *MkDirs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MkDirs_SdkV2) {
}

func (to *MkDirs_SdkV2) SyncFieldsDuringRead(ctx context.Context, from MkDirs_SdkV2) {
}

func (m MkDirs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MkDirs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MkDirs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MkDirs_SdkV2
// only implements ToObjectValue() and Type().
func (m MkDirs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MkDirs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type MkDirsResponse_SdkV2 struct {
}

func (to *MkDirsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MkDirsResponse_SdkV2) {
}

func (to *MkDirsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from MkDirsResponse_SdkV2) {
}

func (m MkDirsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MkDirsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MkDirsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MkDirsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m MkDirsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m MkDirsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Move_SdkV2 struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path.
	DestinationPath types.String `tfsdk:"destination_path"`
	// The source path of the file or directory. The path should be the absolute
	// DBFS path.
	SourcePath types.String `tfsdk:"source_path"`
}

func (to *Move_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Move_SdkV2) {
}

func (to *Move_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Move_SdkV2) {
}

func (m Move_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_path"] = attrs["destination_path"].SetRequired()
	attrs["source_path"] = attrs["source_path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Move.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Move_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Move_SdkV2
// only implements ToObjectValue() and Type().
func (m Move_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_path": m.DestinationPath,
			"source_path":      m.SourcePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Move_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_path": types.StringType,
			"source_path":      types.StringType,
		},
	}
}

type MoveResponse_SdkV2 struct {
}

func (to *MoveResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MoveResponse_SdkV2) {
}

func (to *MoveResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from MoveResponse_SdkV2) {
}

func (m MoveResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MoveResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MoveResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MoveResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m MoveResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m MoveResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Put_SdkV2 struct {
	// This parameter might be absent, and instead a posted file will be used.
	Contents types.String `tfsdk:"contents"`
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite types.Bool `tfsdk:"overwrite"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path"`
}

func (to *Put_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Put_SdkV2) {
}

func (to *Put_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Put_SdkV2) {
}

func (m Put_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["contents"] = attrs["contents"].SetOptional()
	attrs["overwrite"] = attrs["overwrite"].SetOptional()
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Put.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Put_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Put_SdkV2
// only implements ToObjectValue() and Type().
func (m Put_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents":  m.Contents,
			"overwrite": m.Overwrite,
			"path":      m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Put_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents":  types.StringType,
			"overwrite": types.BoolType,
			"path":      types.StringType,
		},
	}
}

type PutResponse_SdkV2 struct {
}

func (to *PutResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutResponse_SdkV2) {
}

func (to *PutResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PutResponse_SdkV2) {
}

func (m PutResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PutResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m PutResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m PutResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ReadDbfsRequest_SdkV2 struct {
	// The number of bytes to read starting from the offset. This has a limit of
	// 1 MB, and a default value of 0.5 MB.
	Length types.Int64 `tfsdk:"-"`
	// The offset to read from in bytes.
	Offset types.Int64 `tfsdk:"-"`
	// The path of the file to read. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"-"`
}

func (to *ReadDbfsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReadDbfsRequest_SdkV2) {
}

func (to *ReadDbfsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ReadDbfsRequest_SdkV2) {
}

func (m ReadDbfsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()
	attrs["offset"] = attrs["offset"].SetOptional()
	attrs["length"] = attrs["length"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReadDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ReadDbfsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReadDbfsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ReadDbfsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"length": m.Length,
			"offset": m.Offset,
			"path":   m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ReadDbfsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"length": types.Int64Type,
			"offset": types.Int64Type,
			"path":   types.StringType,
		},
	}
}

type ReadResponse_SdkV2 struct {
	// The number of bytes read (could be less than ``length`` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	BytesRead types.Int64 `tfsdk:"bytes_read"`
	// The base64-encoded contents of the file read.
	Data types.String `tfsdk:"data"`
}

func (to *ReadResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReadResponse_SdkV2) {
}

func (to *ReadResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ReadResponse_SdkV2) {
}

func (m ReadResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bytes_read"] = attrs["bytes_read"].SetOptional()
	attrs["data"] = attrs["data"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReadResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ReadResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReadResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ReadResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bytes_read": m.BytesRead,
			"data":       m.Data,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ReadResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bytes_read": types.Int64Type,
			"data":       types.StringType,
		},
	}
}

type UploadRequest_SdkV2 struct {
	Contents types.Object `tfsdk:"-"`
	// The absolute path of the file.
	FilePath types.String `tfsdk:"-"`
	// If true or unspecified, an existing file will be overwritten. If false,
	// an error will be returned if the path points to an existing file.
	Overwrite types.Bool `tfsdk:"-"`
}

func (to *UploadRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UploadRequest_SdkV2) {
}

func (to *UploadRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UploadRequest_SdkV2) {
}

func (m UploadRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["contents"] = attrs["contents"].SetRequired()
	attrs["file_path"] = attrs["file_path"].SetRequired()
	attrs["overwrite"] = attrs["overwrite"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UploadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UploadRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UploadRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UploadRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents":  m.Contents,
			"file_path": m.FilePath,
			"overwrite": m.Overwrite,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UploadRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents":  types.ObjectType{},
			"file_path": types.StringType,
			"overwrite": types.BoolType,
		},
	}
}
