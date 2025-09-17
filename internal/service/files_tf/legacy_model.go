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

func (toState *AddBlock_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AddBlock_SdkV2) {
}

func (toState *AddBlock_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AddBlock_SdkV2) {
}

func (c AddBlock_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AddBlock_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddBlock_SdkV2
// only implements ToObjectValue() and Type().
func (o AddBlock_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":   o.Data,
			"handle": o.Handle,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AddBlock_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data":   types.StringType,
			"handle": types.Int64Type,
		},
	}
}

type AddBlockResponse_SdkV2 struct {
}

func (toState *AddBlockResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AddBlockResponse_SdkV2) {
}

func (toState *AddBlockResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState AddBlockResponse_SdkV2) {
}

func (c AddBlockResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AddBlockResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AddBlockResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AddBlockResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o AddBlockResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o AddBlockResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Close_SdkV2 struct {
	// The handle on an open stream.
	Handle types.Int64 `tfsdk:"handle"`
}

func (toState *Close_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Close_SdkV2) {
}

func (toState *Close_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Close_SdkV2) {
}

func (c Close_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Close_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Close_SdkV2
// only implements ToObjectValue() and Type().
func (o Close_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"handle": o.Handle,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Close_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"handle": types.Int64Type,
		},
	}
}

type CloseResponse_SdkV2 struct {
}

func (toState *CloseResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CloseResponse_SdkV2) {
}

func (toState *CloseResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CloseResponse_SdkV2) {
}

func (c CloseResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloseResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CloseResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloseResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CloseResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CloseResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *Create_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Create_SdkV2) {
}

func (toState *Create_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Create_SdkV2) {
}

func (c Create_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Create_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Create_SdkV2
// only implements ToObjectValue() and Type().
func (o Create_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"overwrite": o.Overwrite,
			"path":      o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Create_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *CreateDirectoryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateDirectoryRequest_SdkV2) {
}

func (toState *CreateDirectoryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateDirectoryRequest_SdkV2) {
}

func (c CreateDirectoryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateDirectoryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDirectoryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateDirectoryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"directory_path": o.DirectoryPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDirectoryRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *CreateResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateResponse_SdkV2) {
}

func (toState *CreateResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateResponse_SdkV2) {
}

func (c CreateResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"handle": o.Handle,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *Delete_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Delete_SdkV2) {
}

func (toState *Delete_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Delete_SdkV2) {
}

func (c Delete_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Delete_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Delete_SdkV2
// only implements ToObjectValue() and Type().
func (o Delete_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path":      o.Path,
			"recursive": o.Recursive,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Delete_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteDirectoryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteDirectoryRequest_SdkV2) {
}

func (toState *DeleteDirectoryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteDirectoryRequest_SdkV2) {
}

func (c DeleteDirectoryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDirectoryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDirectoryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDirectoryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"directory_path": o.DirectoryPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDirectoryRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteFileRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteFileRequest_SdkV2) {
}

func (toState *DeleteFileRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteFileRequest_SdkV2) {
}

func (c DeleteFileRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteFileRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFileRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteFileRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_path": o.FilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteFileRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_path": types.StringType,
		},
	}
}

type DeleteResponse_SdkV2 struct {
}

func (toState *DeleteResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteResponse_SdkV2) {
}

func (toState *DeleteResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteResponse_SdkV2) {
}

func (c DeleteResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DirectoryEntry_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DirectoryEntry_SdkV2) {
}

func (toState *DirectoryEntry_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DirectoryEntry_SdkV2) {
}

func (c DirectoryEntry_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DirectoryEntry_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectoryEntry_SdkV2
// only implements ToObjectValue() and Type().
func (o DirectoryEntry_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_size":     o.FileSize,
			"is_directory":  o.IsDirectory,
			"last_modified": o.LastModified,
			"name":          o.Name,
			"path":          o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DirectoryEntry_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DownloadRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DownloadRequest_SdkV2) {
}

func (toState *DownloadRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DownloadRequest_SdkV2) {
}

func (c DownloadRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DownloadRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DownloadRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DownloadRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_path": o.FilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DownloadRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DownloadResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DownloadResponse_SdkV2) {
}

func (toState *DownloadResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DownloadResponse_SdkV2) {
}

func (c DownloadResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DownloadResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DownloadResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DownloadResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content_length": o.ContentLength,
			"content_type":   o.ContentType,
			"contents":       o.Contents,
			"last_modified":  o.LastModified,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DownloadResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *FileInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FileInfo_SdkV2) {
}

func (toState *FileInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FileInfo_SdkV2) {
}

func (c FileInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FileInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o FileInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_size":         o.FileSize,
			"is_dir":            o.IsDir,
			"modification_time": o.ModificationTime,
			"path":              o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FileInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetDirectoryMetadataRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetDirectoryMetadataRequest_SdkV2) {
}

func (toState *GetDirectoryMetadataRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetDirectoryMetadataRequest_SdkV2) {
}

func (c GetDirectoryMetadataRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetDirectoryMetadataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDirectoryMetadataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetDirectoryMetadataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"directory_path": o.DirectoryPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDirectoryMetadataRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetMetadataRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetMetadataRequest_SdkV2) {
}

func (toState *GetMetadataRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetMetadataRequest_SdkV2) {
}

func (c GetMetadataRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetMetadataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetadataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetMetadataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_path": o.FilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetMetadataRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetMetadataResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetMetadataResponse_SdkV2) {
}

func (toState *GetMetadataResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetMetadataResponse_SdkV2) {
}

func (c GetMetadataResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetMetadataResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetadataResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetMetadataResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content_length": o.ContentLength,
			"content_type":   o.ContentType,
			"last_modified":  o.LastModified,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetMetadataResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetStatusRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetStatusRequest_SdkV2) {
}

func (toState *GetStatusRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetStatusRequest_SdkV2) {
}

func (c GetStatusRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetStatusRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatusRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetStatusRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStatusRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ListDbfsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListDbfsRequest_SdkV2) {
}

func (toState *ListDbfsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListDbfsRequest_SdkV2) {
}

func (c ListDbfsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListDbfsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDbfsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListDbfsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDbfsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ListDirectoryContentsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListDirectoryContentsRequest_SdkV2) {
}

func (toState *ListDirectoryContentsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListDirectoryContentsRequest_SdkV2) {
}

func (c ListDirectoryContentsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListDirectoryContentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDirectoryContentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListDirectoryContentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"directory_path": o.DirectoryPath,
			"page_size":      o.PageSize,
			"page_token":     o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDirectoryContentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ListDirectoryResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListDirectoryResponse_SdkV2) {
}

func (toState *ListDirectoryResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListDirectoryResponse_SdkV2) {
}

func (c ListDirectoryResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListDirectoryResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"contents": reflect.TypeOf(DirectoryEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDirectoryResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListDirectoryResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents":        o.Contents,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDirectoryResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListDirectoryResponse_SdkV2) GetContents(ctx context.Context) ([]DirectoryEntry_SdkV2, bool) {
	if o.Contents.IsNull() || o.Contents.IsUnknown() {
		return nil, false
	}
	var v []DirectoryEntry_SdkV2
	d := o.Contents.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetContents sets the value of the Contents field in ListDirectoryResponse_SdkV2.
func (o *ListDirectoryResponse_SdkV2) SetContents(ctx context.Context, v []DirectoryEntry_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["contents"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Contents = types.ListValueMust(t, vs)
}

type ListStatusResponse_SdkV2 struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	Files types.List `tfsdk:"files"`
}

func (toState *ListStatusResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListStatusResponse_SdkV2) {
}

func (toState *ListStatusResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListStatusResponse_SdkV2) {
}

func (c ListStatusResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListStatusResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"files": reflect.TypeOf(FileInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStatusResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListStatusResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"files": o.Files,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListStatusResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListStatusResponse_SdkV2) GetFiles(ctx context.Context) ([]FileInfo_SdkV2, bool) {
	if o.Files.IsNull() || o.Files.IsUnknown() {
		return nil, false
	}
	var v []FileInfo_SdkV2
	d := o.Files.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in ListStatusResponse_SdkV2.
func (o *ListStatusResponse_SdkV2) SetFiles(ctx context.Context, v []FileInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["files"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Files = types.ListValueMust(t, vs)
}

type MkDirs_SdkV2 struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	Path types.String `tfsdk:"path"`
}

func (toState *MkDirs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan MkDirs_SdkV2) {
}

func (toState *MkDirs_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState MkDirs_SdkV2) {
}

func (c MkDirs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a MkDirs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MkDirs_SdkV2
// only implements ToObjectValue() and Type().
func (o MkDirs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MkDirs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type MkDirsResponse_SdkV2 struct {
}

func (toState *MkDirsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan MkDirsResponse_SdkV2) {
}

func (toState *MkDirsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState MkDirsResponse_SdkV2) {
}

func (c MkDirsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MkDirsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MkDirsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MkDirsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o MkDirsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o MkDirsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *Move_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Move_SdkV2) {
}

func (toState *Move_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Move_SdkV2) {
}

func (c Move_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Move_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Move_SdkV2
// only implements ToObjectValue() and Type().
func (o Move_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_path": o.DestinationPath,
			"source_path":      o.SourcePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Move_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_path": types.StringType,
			"source_path":      types.StringType,
		},
	}
}

type MoveResponse_SdkV2 struct {
}

func (toState *MoveResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan MoveResponse_SdkV2) {
}

func (toState *MoveResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState MoveResponse_SdkV2) {
}

func (c MoveResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MoveResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MoveResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MoveResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o MoveResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o MoveResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *Put_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Put_SdkV2) {
}

func (toState *Put_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Put_SdkV2) {
}

func (c Put_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Put_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Put_SdkV2
// only implements ToObjectValue() and Type().
func (o Put_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents":  o.Contents,
			"overwrite": o.Overwrite,
			"path":      o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Put_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *PutResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PutResponse_SdkV2) {
}

func (toState *PutResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PutResponse_SdkV2) {
}

func (c PutResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o PutResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PutResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ReadDbfsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ReadDbfsRequest_SdkV2) {
}

func (toState *ReadDbfsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ReadDbfsRequest_SdkV2) {
}

func (c ReadDbfsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ReadDbfsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReadDbfsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ReadDbfsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"length": o.Length,
			"offset": o.Offset,
			"path":   o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ReadDbfsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ReadResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ReadResponse_SdkV2) {
}

func (toState *ReadResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ReadResponse_SdkV2) {
}

func (c ReadResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ReadResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReadResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ReadResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bytes_read": o.BytesRead,
			"data":       o.Data,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ReadResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *UploadRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UploadRequest_SdkV2) {
}

func (toState *UploadRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UploadRequest_SdkV2) {
}

func (c UploadRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UploadRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UploadRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UploadRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents":  o.Contents,
			"file_path": o.FilePath,
			"overwrite": o.Overwrite,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UploadRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents":  types.ObjectType{},
			"file_path": types.StringType,
			"overwrite": types.BoolType,
		},
	}
}
