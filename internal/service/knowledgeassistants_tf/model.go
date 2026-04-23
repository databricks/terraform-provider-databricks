// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package knowledgeassistants_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateKnowledgeAssistantRequest struct {
	// The Knowledge Assistant to create.
	KnowledgeAssistant types.Object `tfsdk:"knowledge_assistant"`
}

func (to *CreateKnowledgeAssistantRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateKnowledgeAssistantRequest) {
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				// Recursively sync the fields of KnowledgeAssistant
				toKnowledgeAssistant.SyncFieldsDuringCreateOrUpdate(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
}

func (to *CreateKnowledgeAssistantRequest) SyncFieldsDuringRead(ctx context.Context, from CreateKnowledgeAssistantRequest) {
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				toKnowledgeAssistant.SyncFieldsDuringRead(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
}

func (m CreateKnowledgeAssistantRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateKnowledgeAssistantRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateKnowledgeAssistantRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_assistant": reflect.TypeOf(KnowledgeAssistant{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateKnowledgeAssistantRequest
// only implements ToObjectValue() and Type().
func (m CreateKnowledgeAssistantRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistant": m.KnowledgeAssistant,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateKnowledgeAssistantRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistant": KnowledgeAssistant{}.Type(ctx),
		},
	}
}

// GetKnowledgeAssistant returns the value of the KnowledgeAssistant field in CreateKnowledgeAssistantRequest as
// a KnowledgeAssistant value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateKnowledgeAssistantRequest) GetKnowledgeAssistant(ctx context.Context) (KnowledgeAssistant, bool) {
	var e KnowledgeAssistant
	if m.KnowledgeAssistant.IsNull() || m.KnowledgeAssistant.IsUnknown() {
		return e, false
	}
	var v KnowledgeAssistant
	d := m.KnowledgeAssistant.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKnowledgeAssistant sets the value of the KnowledgeAssistant field in CreateKnowledgeAssistantRequest.
func (m *CreateKnowledgeAssistantRequest) SetKnowledgeAssistant(ctx context.Context, v KnowledgeAssistant) {
	vs := v.ToObjectValue(ctx)
	m.KnowledgeAssistant = vs
}

type CreateKnowledgeSourceRequest struct {
	KnowledgeSource types.Object `tfsdk:"knowledge_source"`
	// Parent resource where this source will be created. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent types.String `tfsdk:"-"`
}

func (to *CreateKnowledgeSourceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateKnowledgeSourceRequest) {
	if !from.KnowledgeSource.IsNull() && !from.KnowledgeSource.IsUnknown() {
		if toKnowledgeSource, ok := to.GetKnowledgeSource(ctx); ok {
			if fromKnowledgeSource, ok := from.GetKnowledgeSource(ctx); ok {
				// Recursively sync the fields of KnowledgeSource
				toKnowledgeSource.SyncFieldsDuringCreateOrUpdate(ctx, fromKnowledgeSource)
				to.SetKnowledgeSource(ctx, toKnowledgeSource)
			}
		}
	}
}

func (to *CreateKnowledgeSourceRequest) SyncFieldsDuringRead(ctx context.Context, from CreateKnowledgeSourceRequest) {
	if !from.KnowledgeSource.IsNull() && !from.KnowledgeSource.IsUnknown() {
		if toKnowledgeSource, ok := to.GetKnowledgeSource(ctx); ok {
			if fromKnowledgeSource, ok := from.GetKnowledgeSource(ctx); ok {
				toKnowledgeSource.SyncFieldsDuringRead(ctx, fromKnowledgeSource)
				to.SetKnowledgeSource(ctx, toKnowledgeSource)
			}
		}
	}
}

func (m CreateKnowledgeSourceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_source"] = attrs["knowledge_source"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateKnowledgeSourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateKnowledgeSourceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_source": reflect.TypeOf(KnowledgeSource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateKnowledgeSourceRequest
// only implements ToObjectValue() and Type().
func (m CreateKnowledgeSourceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_source": m.KnowledgeSource,
			"parent":           m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateKnowledgeSourceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_source": KnowledgeSource{}.Type(ctx),
			"parent":           types.StringType,
		},
	}
}

// GetKnowledgeSource returns the value of the KnowledgeSource field in CreateKnowledgeSourceRequest as
// a KnowledgeSource value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateKnowledgeSourceRequest) GetKnowledgeSource(ctx context.Context) (KnowledgeSource, bool) {
	var e KnowledgeSource
	if m.KnowledgeSource.IsNull() || m.KnowledgeSource.IsUnknown() {
		return e, false
	}
	var v KnowledgeSource
	d := m.KnowledgeSource.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKnowledgeSource sets the value of the KnowledgeSource field in CreateKnowledgeSourceRequest.
func (m *CreateKnowledgeSourceRequest) SetKnowledgeSource(ctx context.Context, v KnowledgeSource) {
	vs := v.ToObjectValue(ctx)
	m.KnowledgeSource = vs
}

type DeleteKnowledgeAssistantRequest struct {
	// The resource name of the knowledge assistant to be deleted. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteKnowledgeAssistantRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteKnowledgeAssistantRequest) {
}

func (to *DeleteKnowledgeAssistantRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteKnowledgeAssistantRequest) {
}

func (m DeleteKnowledgeAssistantRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteKnowledgeAssistantRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteKnowledgeAssistantRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteKnowledgeAssistantRequest
// only implements ToObjectValue() and Type().
func (m DeleteKnowledgeAssistantRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteKnowledgeAssistantRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteKnowledgeSourceRequest struct {
	// The resource name of the Knowledge Source to delete. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteKnowledgeSourceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteKnowledgeSourceRequest) {
}

func (to *DeleteKnowledgeSourceRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteKnowledgeSourceRequest) {
}

func (m DeleteKnowledgeSourceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteKnowledgeSourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteKnowledgeSourceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteKnowledgeSourceRequest
// only implements ToObjectValue() and Type().
func (m DeleteKnowledgeSourceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteKnowledgeSourceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// FileTableSpec specifies a file table source configuration.
type FileTableSpec struct {
	// The name of the column containing BINARY file content to be indexed.
	FileCol types.String `tfsdk:"file_col"`
	// Full UC name of the table, in the format of
	// {CATALOG}.{SCHEMA}.{TABLE_NAME}.
	TableName types.String `tfsdk:"table_name"`
}

func (to *FileTableSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileTableSpec) {
}

func (to *FileTableSpec) SyncFieldsDuringRead(ctx context.Context, from FileTableSpec) {
}

func (m FileTableSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_col"] = attrs["file_col"].SetRequired()
	attrs["table_name"] = attrs["table_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileTableSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FileTableSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileTableSpec
// only implements ToObjectValue() and Type().
func (m FileTableSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_col":   m.FileCol,
			"table_name": m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FileTableSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_col":   types.StringType,
			"table_name": types.StringType,
		},
	}
}

// FilesSpec specifies a files source configuration.
type FilesSpec struct {
	// A UC volume path that includes a list of files.
	Path types.String `tfsdk:"path"`
}

func (to *FilesSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FilesSpec) {
}

func (to *FilesSpec) SyncFieldsDuringRead(ctx context.Context, from FilesSpec) {
}

func (m FilesSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FilesSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FilesSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FilesSpec
// only implements ToObjectValue() and Type().
func (m FilesSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FilesSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type GetKnowledgeAssistantPermissionLevelsRequest struct {
	// The knowledge assistant for which to get or manage permissions.
	KnowledgeAssistantId types.String `tfsdk:"-"`
}

func (to *GetKnowledgeAssistantPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetKnowledgeAssistantPermissionLevelsRequest) {
}

func (to *GetKnowledgeAssistantPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetKnowledgeAssistantPermissionLevelsRequest) {
}

func (m GetKnowledgeAssistantPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistant_id"] = attrs["knowledge_assistant_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetKnowledgeAssistantPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetKnowledgeAssistantPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetKnowledgeAssistantPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetKnowledgeAssistantPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistant_id": m.KnowledgeAssistantId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetKnowledgeAssistantPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistant_id": types.StringType,
		},
	}
}

type GetKnowledgeAssistantPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetKnowledgeAssistantPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetKnowledgeAssistantPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetKnowledgeAssistantPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetKnowledgeAssistantPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetKnowledgeAssistantPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetKnowledgeAssistantPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetKnowledgeAssistantPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(KnowledgeAssistantPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetKnowledgeAssistantPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetKnowledgeAssistantPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetKnowledgeAssistantPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: KnowledgeAssistantPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetKnowledgeAssistantPermissionLevelsResponse as
// a slice of KnowledgeAssistantPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetKnowledgeAssistantPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]KnowledgeAssistantPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeAssistantPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetKnowledgeAssistantPermissionLevelsResponse.
func (m *GetKnowledgeAssistantPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []KnowledgeAssistantPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetKnowledgeAssistantPermissionsRequest struct {
	// The knowledge assistant for which to get or manage permissions.
	KnowledgeAssistantId types.String `tfsdk:"-"`
}

func (to *GetKnowledgeAssistantPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetKnowledgeAssistantPermissionsRequest) {
}

func (to *GetKnowledgeAssistantPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetKnowledgeAssistantPermissionsRequest) {
}

func (m GetKnowledgeAssistantPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistant_id"] = attrs["knowledge_assistant_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetKnowledgeAssistantPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetKnowledgeAssistantPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetKnowledgeAssistantPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetKnowledgeAssistantPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistant_id": m.KnowledgeAssistantId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetKnowledgeAssistantPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistant_id": types.StringType,
		},
	}
}

type GetKnowledgeAssistantRequest struct {
	// The resource name of the knowledge assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetKnowledgeAssistantRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetKnowledgeAssistantRequest) {
}

func (to *GetKnowledgeAssistantRequest) SyncFieldsDuringRead(ctx context.Context, from GetKnowledgeAssistantRequest) {
}

func (m GetKnowledgeAssistantRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetKnowledgeAssistantRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetKnowledgeAssistantRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetKnowledgeAssistantRequest
// only implements ToObjectValue() and Type().
func (m GetKnowledgeAssistantRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetKnowledgeAssistantRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetKnowledgeSourceRequest struct {
	// The resource name of the Knowledge Source. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetKnowledgeSourceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetKnowledgeSourceRequest) {
}

func (to *GetKnowledgeSourceRequest) SyncFieldsDuringRead(ctx context.Context, from GetKnowledgeSourceRequest) {
}

func (m GetKnowledgeSourceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetKnowledgeSourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetKnowledgeSourceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetKnowledgeSourceRequest
// only implements ToObjectValue() and Type().
func (m GetKnowledgeSourceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetKnowledgeSourceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// IndexSpec specifies a vector search index source configuration.
type IndexSpec struct {
	// The column that specifies a link or reference to where the information
	// came from.
	DocUriCol types.String `tfsdk:"doc_uri_col"`
	// Full UC name of the vector search index, in the format of
	// {CATALOG}.{SCHEMA}.{INDEX_NAME}.
	IndexName types.String `tfsdk:"index_name"`
	// The column that includes the document text for retrieval.
	TextCol types.String `tfsdk:"text_col"`
}

func (to *IndexSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IndexSpec) {
}

func (to *IndexSpec) SyncFieldsDuringRead(ctx context.Context, from IndexSpec) {
}

func (m IndexSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["doc_uri_col"] = attrs["doc_uri_col"].SetRequired()
	attrs["index_name"] = attrs["index_name"].SetRequired()
	attrs["text_col"] = attrs["text_col"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IndexSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m IndexSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IndexSpec
// only implements ToObjectValue() and Type().
func (m IndexSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"doc_uri_col": m.DocUriCol,
			"index_name":  m.IndexName,
			"text_col":    m.TextCol,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IndexSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"doc_uri_col": types.StringType,
			"index_name":  types.StringType,
			"text_col":    types.StringType,
		},
	}
}

// Entity message that represents a knowledge assistant. Note: REQUIRED
// annotations below represent create-time requirements. For updates, required
// fields are determined by the update mask.
type KnowledgeAssistant struct {
	// Creation timestamp.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The creator of the Knowledge Assistant.
	Creator types.String `tfsdk:"creator"`
	// Description of what this agent can do (user-facing). Required when
	// creating a Knowledge Assistant. When updating a Knowledge Assistant,
	// optional unless included in update_mask.
	Description types.String `tfsdk:"description"`
	// The display name of the Knowledge Assistant, unique at workspace level.
	// Required when creating a Knowledge Assistant. When updating a Knowledge
	// Assistant, optional unless included in update_mask.
	DisplayName types.String `tfsdk:"display_name"`
	// The name of the knowledge assistant agent endpoint.
	EndpointName types.String `tfsdk:"endpoint_name"`
	// Error details when the Knowledge Assistant is in FAILED state.
	ErrorInfo types.String `tfsdk:"error_info"`
	// The MLflow experiment ID.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// The universally unique identifier (UUID) of the Knowledge Assistant.
	Id types.String `tfsdk:"id"`
	// Additional global instructions on how the agent should generate answers.
	// Optional on create and update. When updating a Knowledge Assistant,
	// include this field in update_mask to modify it.
	Instructions types.String `tfsdk:"instructions"`
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"name"`
	// State of the Knowledge Assistant. Not returned in List responses.
	State types.String `tfsdk:"state"`
}

func (to *KnowledgeAssistant) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistant) {
}

func (to *KnowledgeAssistant) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistant) {
}

func (m KnowledgeAssistant) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetRequired()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetComputed()
	attrs["error_info"] = attrs["error_info"].SetComputed()
	attrs["experiment_id"] = attrs["experiment_id"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["instructions"] = attrs["instructions"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistant.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistant) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistant
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistant) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":   m.CreateTime,
			"creator":       m.Creator,
			"description":   m.Description,
			"display_name":  m.DisplayName,
			"endpoint_name": m.EndpointName,
			"error_info":    m.ErrorInfo,
			"experiment_id": m.ExperimentId,
			"id":            m.Id,
			"instructions":  m.Instructions,
			"name":          m.Name,
			"state":         m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistant) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":   timetypes.RFC3339{}.Type(ctx),
			"creator":       types.StringType,
			"description":   types.StringType,
			"display_name":  types.StringType,
			"endpoint_name": types.StringType,
			"error_info":    types.StringType,
			"experiment_id": types.StringType,
			"id":            types.StringType,
			"instructions":  types.StringType,
			"name":          types.StringType,
			"state":         types.StringType,
		},
	}
}

type KnowledgeAssistantAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *KnowledgeAssistantAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantAccessControlRequest) {
}

func (to *KnowledgeAssistantAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantAccessControlRequest) {
}

func (m KnowledgeAssistantAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantAccessControlRequest
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistantAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type KnowledgeAssistantAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *KnowledgeAssistantAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *KnowledgeAssistantAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m KnowledgeAssistantAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(KnowledgeAssistantPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantAccessControlResponse
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistantAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: KnowledgeAssistantPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in KnowledgeAssistantAccessControlResponse as
// a slice of KnowledgeAssistantPermission values.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeAssistantAccessControlResponse) GetAllPermissions(ctx context.Context) ([]KnowledgeAssistantPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeAssistantPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in KnowledgeAssistantAccessControlResponse.
func (m *KnowledgeAssistantAccessControlResponse) SetAllPermissions(ctx context.Context, v []KnowledgeAssistantPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type KnowledgeAssistantPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *KnowledgeAssistantPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *KnowledgeAssistantPermission) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m KnowledgeAssistantPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantPermission
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistantPermission) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

// GetInheritedFromObject returns the value of the InheritedFromObject field in KnowledgeAssistantPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeAssistantPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in KnowledgeAssistantPermission.
func (m *KnowledgeAssistantPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type KnowledgeAssistantPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *KnowledgeAssistantPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *KnowledgeAssistantPermissions) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m KnowledgeAssistantPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(KnowledgeAssistantAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantPermissions
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistantPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: KnowledgeAssistantAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in KnowledgeAssistantPermissions as
// a slice of KnowledgeAssistantAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeAssistantPermissions) GetAccessControlList(ctx context.Context) ([]KnowledgeAssistantAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeAssistantAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in KnowledgeAssistantPermissions.
func (m *KnowledgeAssistantPermissions) SetAccessControlList(ctx context.Context, v []KnowledgeAssistantAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type KnowledgeAssistantPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *KnowledgeAssistantPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantPermissionsDescription) {
}

func (to *KnowledgeAssistantPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantPermissionsDescription) {
}

func (m KnowledgeAssistantPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantPermissionsDescription
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistantPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type KnowledgeAssistantPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The knowledge assistant for which to get or manage permissions.
	KnowledgeAssistantId types.String `tfsdk:"-"`
}

func (to *KnowledgeAssistantPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *KnowledgeAssistantPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m KnowledgeAssistantPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["knowledge_assistant_id"] = attrs["knowledge_assistant_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(KnowledgeAssistantAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantPermissionsRequest
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list":    m.AccessControlList,
			"knowledge_assistant_id": m.KnowledgeAssistantId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistantPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: KnowledgeAssistantAccessControlRequest{}.Type(ctx),
			},
			"knowledge_assistant_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in KnowledgeAssistantPermissionsRequest as
// a slice of KnowledgeAssistantAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeAssistantPermissionsRequest) GetAccessControlList(ctx context.Context) ([]KnowledgeAssistantAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeAssistantAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in KnowledgeAssistantPermissionsRequest.
func (m *KnowledgeAssistantPermissionsRequest) SetAccessControlList(ctx context.Context, v []KnowledgeAssistantAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

// KnowledgeSource represents a source of knowledge for the KnowledgeAssistant.
// Used in create/update requests and returned in Get/List responses. Note:
// REQUIRED annotations below represent create-time requirements. For updates,
// required fields are determined by the update mask.
type KnowledgeSource struct {
	// Timestamp when this knowledge source was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Description of the knowledge source. Required when creating a Knowledge
	// Source. When updating a Knowledge Source, optional unless included in
	// update_mask.
	Description types.String `tfsdk:"description"`
	// Human-readable display name of the knowledge source. Required when
	// creating a Knowledge Source. When updating a Knowledge Source, optional
	// unless included in update_mask.
	DisplayName types.String `tfsdk:"display_name"`

	FileTable types.Object `tfsdk:"file_table"`

	Files types.Object `tfsdk:"files"`

	Id types.String `tfsdk:"id"`

	Index types.Object `tfsdk:"index"`
	// Timestamp representing the cutoff before which content in this knowledge
	// source is being ingested.
	KnowledgeCutoffTime timetypes.RFC3339 `tfsdk:"knowledge_cutoff_time"`
	// Full resource name:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name types.String `tfsdk:"name"`
	// The type of the source: "index", "files", or "file_table". Required when
	// creating a Knowledge Source. When updating a Knowledge Source, this field
	// is ignored.
	SourceType types.String `tfsdk:"source_type"`

	State types.String `tfsdk:"state"`
}

func (to *KnowledgeSource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeSource) {
	if !from.FileTable.IsNull() && !from.FileTable.IsUnknown() {
		if toFileTable, ok := to.GetFileTable(ctx); ok {
			if fromFileTable, ok := from.GetFileTable(ctx); ok {
				// Recursively sync the fields of FileTable
				toFileTable.SyncFieldsDuringCreateOrUpdate(ctx, fromFileTable)
				to.SetFileTable(ctx, toFileTable)
			}
		}
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				// Recursively sync the fields of Files
				toFiles.SyncFieldsDuringCreateOrUpdate(ctx, fromFiles)
				to.SetFiles(ctx, toFiles)
			}
		}
	}
	if !from.Index.IsNull() && !from.Index.IsUnknown() {
		if toIndex, ok := to.GetIndex(ctx); ok {
			if fromIndex, ok := from.GetIndex(ctx); ok {
				// Recursively sync the fields of Index
				toIndex.SyncFieldsDuringCreateOrUpdate(ctx, fromIndex)
				to.SetIndex(ctx, toIndex)
			}
		}
	}
}

func (to *KnowledgeSource) SyncFieldsDuringRead(ctx context.Context, from KnowledgeSource) {
	if !from.FileTable.IsNull() && !from.FileTable.IsUnknown() {
		if toFileTable, ok := to.GetFileTable(ctx); ok {
			if fromFileTable, ok := from.GetFileTable(ctx); ok {
				toFileTable.SyncFieldsDuringRead(ctx, fromFileTable)
				to.SetFileTable(ctx, toFileTable)
			}
		}
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				toFiles.SyncFieldsDuringRead(ctx, fromFiles)
				to.SetFiles(ctx, toFiles)
			}
		}
	}
	if !from.Index.IsNull() && !from.Index.IsUnknown() {
		if toIndex, ok := to.GetIndex(ctx); ok {
			if fromIndex, ok := from.GetIndex(ctx); ok {
				toIndex.SyncFieldsDuringRead(ctx, fromIndex)
				to.SetIndex(ctx, toIndex)
			}
		}
	}
}

func (m KnowledgeSource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetRequired()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["file_table"] = attrs["file_table"].SetOptional()
	attrs["files"] = attrs["files"].SetOptional()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["index"] = attrs["index"].SetOptional()
	attrs["knowledge_cutoff_time"] = attrs["knowledge_cutoff_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["source_type"] = attrs["source_type"].SetRequired()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeSource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_table": reflect.TypeOf(FileTableSpec{}),
		"files":      reflect.TypeOf(FilesSpec{}),
		"index":      reflect.TypeOf(IndexSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeSource
// only implements ToObjectValue() and Type().
func (m KnowledgeSource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":           m.CreateTime,
			"description":           m.Description,
			"display_name":          m.DisplayName,
			"file_table":            m.FileTable,
			"files":                 m.Files,
			"id":                    m.Id,
			"index":                 m.Index,
			"knowledge_cutoff_time": m.KnowledgeCutoffTime,
			"name":                  m.Name,
			"source_type":           m.SourceType,
			"state":                 m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeSource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":           timetypes.RFC3339{}.Type(ctx),
			"description":           types.StringType,
			"display_name":          types.StringType,
			"file_table":            FileTableSpec{}.Type(ctx),
			"files":                 FilesSpec{}.Type(ctx),
			"id":                    types.StringType,
			"index":                 IndexSpec{}.Type(ctx),
			"knowledge_cutoff_time": timetypes.RFC3339{}.Type(ctx),
			"name":                  types.StringType,
			"source_type":           types.StringType,
			"state":                 types.StringType,
		},
	}
}

// GetFileTable returns the value of the FileTable field in KnowledgeSource as
// a FileTableSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeSource) GetFileTable(ctx context.Context) (FileTableSpec, bool) {
	var e FileTableSpec
	if m.FileTable.IsNull() || m.FileTable.IsUnknown() {
		return e, false
	}
	var v FileTableSpec
	d := m.FileTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFileTable sets the value of the FileTable field in KnowledgeSource.
func (m *KnowledgeSource) SetFileTable(ctx context.Context, v FileTableSpec) {
	vs := v.ToObjectValue(ctx)
	m.FileTable = vs
}

// GetFiles returns the value of the Files field in KnowledgeSource as
// a FilesSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeSource) GetFiles(ctx context.Context) (FilesSpec, bool) {
	var e FilesSpec
	if m.Files.IsNull() || m.Files.IsUnknown() {
		return e, false
	}
	var v FilesSpec
	d := m.Files.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in KnowledgeSource.
func (m *KnowledgeSource) SetFiles(ctx context.Context, v FilesSpec) {
	vs := v.ToObjectValue(ctx)
	m.Files = vs
}

// GetIndex returns the value of the Index field in KnowledgeSource as
// a IndexSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeSource) GetIndex(ctx context.Context) (IndexSpec, bool) {
	var e IndexSpec
	if m.Index.IsNull() || m.Index.IsUnknown() {
		return e, false
	}
	var v IndexSpec
	d := m.Index.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIndex sets the value of the Index field in KnowledgeSource.
func (m *KnowledgeSource) SetIndex(ctx context.Context, v IndexSpec) {
	vs := v.ToObjectValue(ctx)
	m.Index = vs
}

type ListKnowledgeAssistantsRequest struct {
	// The maximum number of knowledge assistants to return. If unspecified, at
	// most 100 knowledge assistants will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListKnowledgeAssistants` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListKnowledgeAssistantsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListKnowledgeAssistantsRequest) {
}

func (to *ListKnowledgeAssistantsRequest) SyncFieldsDuringRead(ctx context.Context, from ListKnowledgeAssistantsRequest) {
}

func (m ListKnowledgeAssistantsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListKnowledgeAssistantsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListKnowledgeAssistantsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListKnowledgeAssistantsRequest
// only implements ToObjectValue() and Type().
func (m ListKnowledgeAssistantsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListKnowledgeAssistantsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// A list of Knowledge Assistants.
type ListKnowledgeAssistantsResponse struct {
	KnowledgeAssistants types.List `tfsdk:"knowledge_assistants"`
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListKnowledgeAssistantsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListKnowledgeAssistantsResponse) {
	if !from.KnowledgeAssistants.IsNull() && !from.KnowledgeAssistants.IsUnknown() && to.KnowledgeAssistants.IsNull() && len(from.KnowledgeAssistants.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for KnowledgeAssistants, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.KnowledgeAssistants = from.KnowledgeAssistants
	}
}

func (to *ListKnowledgeAssistantsResponse) SyncFieldsDuringRead(ctx context.Context, from ListKnowledgeAssistantsResponse) {
	if !from.KnowledgeAssistants.IsNull() && !from.KnowledgeAssistants.IsUnknown() && to.KnowledgeAssistants.IsNull() && len(from.KnowledgeAssistants.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for KnowledgeAssistants, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.KnowledgeAssistants = from.KnowledgeAssistants
	}
}

func (m ListKnowledgeAssistantsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistants"] = attrs["knowledge_assistants"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListKnowledgeAssistantsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListKnowledgeAssistantsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_assistants": reflect.TypeOf(KnowledgeAssistant{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListKnowledgeAssistantsResponse
// only implements ToObjectValue() and Type().
func (m ListKnowledgeAssistantsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistants": m.KnowledgeAssistants,
			"next_page_token":      m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListKnowledgeAssistantsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistants": basetypes.ListType{
				ElemType: KnowledgeAssistant{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetKnowledgeAssistants returns the value of the KnowledgeAssistants field in ListKnowledgeAssistantsResponse as
// a slice of KnowledgeAssistant values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListKnowledgeAssistantsResponse) GetKnowledgeAssistants(ctx context.Context) ([]KnowledgeAssistant, bool) {
	if m.KnowledgeAssistants.IsNull() || m.KnowledgeAssistants.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeAssistant
	d := m.KnowledgeAssistants.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKnowledgeAssistants sets the value of the KnowledgeAssistants field in ListKnowledgeAssistantsResponse.
func (m *ListKnowledgeAssistantsResponse) SetKnowledgeAssistants(ctx context.Context, v []KnowledgeAssistant) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["knowledge_assistants"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.KnowledgeAssistants = types.ListValueMust(t, vs)
}

type ListKnowledgeSourcesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Parent resource to list from. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListKnowledgeSourcesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListKnowledgeSourcesRequest) {
}

func (to *ListKnowledgeSourcesRequest) SyncFieldsDuringRead(ctx context.Context, from ListKnowledgeSourcesRequest) {
}

func (m ListKnowledgeSourcesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListKnowledgeSourcesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListKnowledgeSourcesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListKnowledgeSourcesRequest
// only implements ToObjectValue() and Type().
func (m ListKnowledgeSourcesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListKnowledgeSourcesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListKnowledgeSourcesResponse struct {
	KnowledgeSources types.List `tfsdk:"knowledge_sources"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListKnowledgeSourcesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListKnowledgeSourcesResponse) {
	if !from.KnowledgeSources.IsNull() && !from.KnowledgeSources.IsUnknown() && to.KnowledgeSources.IsNull() && len(from.KnowledgeSources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for KnowledgeSources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.KnowledgeSources = from.KnowledgeSources
	}
}

func (to *ListKnowledgeSourcesResponse) SyncFieldsDuringRead(ctx context.Context, from ListKnowledgeSourcesResponse) {
	if !from.KnowledgeSources.IsNull() && !from.KnowledgeSources.IsUnknown() && to.KnowledgeSources.IsNull() && len(from.KnowledgeSources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for KnowledgeSources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.KnowledgeSources = from.KnowledgeSources
	}
}

func (m ListKnowledgeSourcesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_sources"] = attrs["knowledge_sources"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListKnowledgeSourcesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListKnowledgeSourcesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_sources": reflect.TypeOf(KnowledgeSource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListKnowledgeSourcesResponse
// only implements ToObjectValue() and Type().
func (m ListKnowledgeSourcesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_sources": m.KnowledgeSources,
			"next_page_token":   m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListKnowledgeSourcesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_sources": basetypes.ListType{
				ElemType: KnowledgeSource{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetKnowledgeSources returns the value of the KnowledgeSources field in ListKnowledgeSourcesResponse as
// a slice of KnowledgeSource values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListKnowledgeSourcesResponse) GetKnowledgeSources(ctx context.Context) ([]KnowledgeSource, bool) {
	if m.KnowledgeSources.IsNull() || m.KnowledgeSources.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeSource
	d := m.KnowledgeSources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKnowledgeSources sets the value of the KnowledgeSources field in ListKnowledgeSourcesResponse.
func (m *ListKnowledgeSourcesResponse) SetKnowledgeSources(ctx context.Context, v []KnowledgeSource) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["knowledge_sources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.KnowledgeSources = types.ListValueMust(t, vs)
}

type SyncKnowledgeSourcesRequest struct {
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"-"`
}

func (to *SyncKnowledgeSourcesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncKnowledgeSourcesRequest) {
}

func (to *SyncKnowledgeSourcesRequest) SyncFieldsDuringRead(ctx context.Context, from SyncKnowledgeSourcesRequest) {
}

func (m SyncKnowledgeSourcesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncKnowledgeSourcesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncKnowledgeSourcesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncKnowledgeSourcesRequest
// only implements ToObjectValue() and Type().
func (m SyncKnowledgeSourcesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncKnowledgeSourcesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type UpdateKnowledgeAssistantRequest struct {
	// The Knowledge Assistant update payload. Only fields listed in update_mask
	// are updated. REQUIRED annotations on Knowledge Assistant fields describe
	// create-time requirements and do not mean all those fields are required
	// for update.
	KnowledgeAssistant types.Object `tfsdk:"knowledge_assistant"`
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"-"`
	// Comma-delimited list of fields to update on the Knowledge Assistant.
	// Allowed values: `display_name`, `description`, `instructions`. Examples:
	// - `display_name` - `description,instructions`
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateKnowledgeAssistantRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateKnowledgeAssistantRequest) {
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				// Recursively sync the fields of KnowledgeAssistant
				toKnowledgeAssistant.SyncFieldsDuringCreateOrUpdate(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
}

func (to *UpdateKnowledgeAssistantRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateKnowledgeAssistantRequest) {
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				toKnowledgeAssistant.SyncFieldsDuringRead(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
}

func (m UpdateKnowledgeAssistantRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateKnowledgeAssistantRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateKnowledgeAssistantRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_assistant": reflect.TypeOf(KnowledgeAssistant{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateKnowledgeAssistantRequest
// only implements ToObjectValue() and Type().
func (m UpdateKnowledgeAssistantRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistant": m.KnowledgeAssistant,
			"name":                m.Name,
			"update_mask":         m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateKnowledgeAssistantRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistant": KnowledgeAssistant{}.Type(ctx),
			"name":                types.StringType,
			"update_mask":         types.StringType,
		},
	}
}

// GetKnowledgeAssistant returns the value of the KnowledgeAssistant field in UpdateKnowledgeAssistantRequest as
// a KnowledgeAssistant value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateKnowledgeAssistantRequest) GetKnowledgeAssistant(ctx context.Context) (KnowledgeAssistant, bool) {
	var e KnowledgeAssistant
	if m.KnowledgeAssistant.IsNull() || m.KnowledgeAssistant.IsUnknown() {
		return e, false
	}
	var v KnowledgeAssistant
	d := m.KnowledgeAssistant.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKnowledgeAssistant sets the value of the KnowledgeAssistant field in UpdateKnowledgeAssistantRequest.
func (m *UpdateKnowledgeAssistantRequest) SetKnowledgeAssistant(ctx context.Context, v KnowledgeAssistant) {
	vs := v.ToObjectValue(ctx)
	m.KnowledgeAssistant = vs
}

type UpdateKnowledgeSourceRequest struct {
	// The Knowledge Source update payload. Only fields listed in update_mask
	// are updated. REQUIRED annotations on Knowledge Source fields describe
	// create-time requirements and do not mean all those fields are required
	// for update.
	KnowledgeSource types.Object `tfsdk:"knowledge_source"`
	// The resource name of the Knowledge Source to update. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name types.String `tfsdk:"-"`
	// Comma-delimited list of fields to update on the Knowledge Source. Allowed
	// values: `display_name`, `description`. Examples: - `display_name` -
	// `display_name,description`
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateKnowledgeSourceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateKnowledgeSourceRequest) {
	if !from.KnowledgeSource.IsNull() && !from.KnowledgeSource.IsUnknown() {
		if toKnowledgeSource, ok := to.GetKnowledgeSource(ctx); ok {
			if fromKnowledgeSource, ok := from.GetKnowledgeSource(ctx); ok {
				// Recursively sync the fields of KnowledgeSource
				toKnowledgeSource.SyncFieldsDuringCreateOrUpdate(ctx, fromKnowledgeSource)
				to.SetKnowledgeSource(ctx, toKnowledgeSource)
			}
		}
	}
}

func (to *UpdateKnowledgeSourceRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateKnowledgeSourceRequest) {
	if !from.KnowledgeSource.IsNull() && !from.KnowledgeSource.IsUnknown() {
		if toKnowledgeSource, ok := to.GetKnowledgeSource(ctx); ok {
			if fromKnowledgeSource, ok := from.GetKnowledgeSource(ctx); ok {
				toKnowledgeSource.SyncFieldsDuringRead(ctx, fromKnowledgeSource)
				to.SetKnowledgeSource(ctx, toKnowledgeSource)
			}
		}
	}
}

func (m UpdateKnowledgeSourceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_source"] = attrs["knowledge_source"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateKnowledgeSourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateKnowledgeSourceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_source": reflect.TypeOf(KnowledgeSource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateKnowledgeSourceRequest
// only implements ToObjectValue() and Type().
func (m UpdateKnowledgeSourceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_source": m.KnowledgeSource,
			"name":             m.Name,
			"update_mask":      m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateKnowledgeSourceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_source": KnowledgeSource{}.Type(ctx),
			"name":             types.StringType,
			"update_mask":      types.StringType,
		},
	}
}

// GetKnowledgeSource returns the value of the KnowledgeSource field in UpdateKnowledgeSourceRequest as
// a KnowledgeSource value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateKnowledgeSourceRequest) GetKnowledgeSource(ctx context.Context) (KnowledgeSource, bool) {
	var e KnowledgeSource
	if m.KnowledgeSource.IsNull() || m.KnowledgeSource.IsUnknown() {
		return e, false
	}
	var v KnowledgeSource
	d := m.KnowledgeSource.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKnowledgeSource sets the value of the KnowledgeSource field in UpdateKnowledgeSourceRequest.
func (m *UpdateKnowledgeSourceRequest) SetKnowledgeSource(ctx context.Context, v KnowledgeSource) {
	vs := v.ToObjectValue(ctx)
	m.KnowledgeSource = vs
}
