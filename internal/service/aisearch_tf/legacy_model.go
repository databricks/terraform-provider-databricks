// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package aisearch_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Column information (name and data type) for an index column. Surfaced on
// `Index.column_info`.
type ColumnInfo_SdkV2 struct {
	// Name of the column.
	Name types.String `tfsdk:"name"`
	// Data type of the column (e.g., "string", "int", "array<float>").
	TypeText types.String `tfsdk:"type_text"`
}

func (to *ColumnInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ColumnInfo_SdkV2) {
}

func (to *ColumnInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ColumnInfo_SdkV2) {
}

func (m ColumnInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetComputed()
	attrs["type_text"] = attrs["type_text"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ColumnInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ColumnInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m ColumnInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":      m.Name,
			"type_text": m.TypeText,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ColumnInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":      types.StringType,
			"type_text": types.StringType,
		},
	}
}

type CreateEndpointRequest_SdkV2 struct {
	// The Endpoint resource to create. Fields other than `endpoint.name` carry
	// the desired configuration; `endpoint.name` is server-assigned from
	// `parent` and `endpoint_id`.
	Endpoint types.List `tfsdk:"endpoint"`
	// The user-supplied short name for the Endpoint, per AIP-133. The server
	// composes the full `Endpoint.name` as `{parent}/endpoints/{endpoint_id}`.
	// AIP-133 does not list `endpoint_id` as a fields-may-be-required entry, so
	// we annotate it OPTIONAL on the wire; the server still rejects empty
	// values with INVALID_PARAMETER_VALUE.
	EndpointId types.String `tfsdk:"-"`
	// The Workspace where this Endpoint will be created. Format:
	// `workspaces/{workspace_id}`
	Parent types.String `tfsdk:"-"`
}

func (to *CreateEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				// Recursively sync the fields of Endpoint
				toEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (to *CreateEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m CreateEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
	attrs["endpoint"] = attrs["endpoint"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint":    m.Endpoint,
			"endpoint_id": m.EndpointId,
			"parent":      m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint": basetypes.ListType{
				ElemType: Endpoint_SdkV2{}.Type(ctx),
			},
			"endpoint_id": types.StringType,
			"parent":      types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in CreateEndpointRequest_SdkV2 as
// a Endpoint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateEndpointRequest_SdkV2) GetEndpoint(ctx context.Context) (Endpoint_SdkV2, bool) {
	var e Endpoint_SdkV2
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v []Endpoint_SdkV2
	d := m.Endpoint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEndpoint sets the value of the Endpoint field in CreateEndpointRequest_SdkV2.
func (m *CreateEndpointRequest_SdkV2) SetEndpoint(ctx context.Context, v Endpoint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoint"]
	m.Endpoint = types.ListValueMust(t, vs)
}

type CreateIndexRequest_SdkV2 struct {
	// The Index resource to create. Fields other than `index.name` carry the
	// desired configuration; `index.name` is server-assigned from `parent` and
	// `index_id`.
	Index types.List `tfsdk:"index"`
	// The user-supplied Unity Catalog table name for the Index, per AIP-133.
	// The server composes the full `Index.name` as
	// `{parent}/indexes/{index_id}`. AIP-133 does not list `index_id` as a
	// fields-may-be-required entry, so we annotate it OPTIONAL on the wire; the
	// server still rejects empty values with INVALID_PARAMETER_VALUE.
	IndexId types.String `tfsdk:"-"`
	// The Endpoint where this Index will be created. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`
	Parent types.String `tfsdk:"-"`
}

func (to *CreateIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateIndexRequest_SdkV2) {
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

func (to *CreateIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateIndexRequest_SdkV2) {
	if !from.Index.IsNull() && !from.Index.IsUnknown() {
		if toIndex, ok := to.GetIndex(ctx); ok {
			if fromIndex, ok := from.GetIndex(ctx); ok {
				toIndex.SyncFieldsDuringRead(ctx, fromIndex)
				to.SetIndex(ctx, toIndex)
			}
		}
	}
}

func (m CreateIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["index"] = attrs["index"].SetRequired()
	attrs["index"] = attrs["index"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["index_id"] = attrs["index_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"index": reflect.TypeOf(Index_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index":    m.Index,
			"index_id": m.IndexId,
			"parent":   m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index": basetypes.ListType{
				ElemType: Index_SdkV2{}.Type(ctx),
			},
			"index_id": types.StringType,
			"parent":   types.StringType,
		},
	}
}

// GetIndex returns the value of the Index field in CreateIndexRequest_SdkV2 as
// a Index_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateIndexRequest_SdkV2) GetIndex(ctx context.Context) (Index_SdkV2, bool) {
	var e Index_SdkV2
	if m.Index.IsNull() || m.Index.IsUnknown() {
		return e, false
	}
	var v []Index_SdkV2
	d := m.Index.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIndex sets the value of the Index field in CreateIndexRequest_SdkV2.
func (m *CreateIndexRequest_SdkV2) SetIndex(ctx context.Context, v Index_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["index"]
	m.Index = types.ListValueMust(t, vs)
}

// User-defined key/value tag attached to an AI Search endpoint for cost
// attribution and access control.
type CustomTag_SdkV2 struct {
	// Key field for an AI Search endpoint tag.
	Key types.String `tfsdk:"key"`
	// [Optional] Value field for an AI Search endpoint tag.
	Value types.String `tfsdk:"value"`
}

func (to *CustomTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomTag_SdkV2) {
}

func (to *CustomTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CustomTag_SdkV2) {
}

func (m CustomTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CustomTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomTag_SdkV2
// only implements ToObjectValue() and Type().
func (m CustomTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Per-row outcome of a data-plane upsert or delete operation.
type DataModificationResult_SdkV2 struct {
	// Primary keys of rows that failed to process.
	FailedPrimaryKeys types.List `tfsdk:"failed_primary_keys"`
	// Count of rows processed successfully.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count"`
}

func (to *DataModificationResult_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataModificationResult_SdkV2) {
	if !from.FailedPrimaryKeys.IsNull() && !from.FailedPrimaryKeys.IsUnknown() && to.FailedPrimaryKeys.IsNull() && len(from.FailedPrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailedPrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailedPrimaryKeys = from.FailedPrimaryKeys
	}
}

func (to *DataModificationResult_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DataModificationResult_SdkV2) {
	if !from.FailedPrimaryKeys.IsNull() && !from.FailedPrimaryKeys.IsUnknown() && to.FailedPrimaryKeys.IsNull() && len(from.FailedPrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailedPrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailedPrimaryKeys = from.FailedPrimaryKeys
	}
}

func (m DataModificationResult_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["failed_primary_keys"] = attrs["failed_primary_keys"].SetComputed()
	attrs["success_row_count"] = attrs["success_row_count"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataModificationResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DataModificationResult_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failed_primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataModificationResult_SdkV2
// only implements ToObjectValue() and Type().
func (m DataModificationResult_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failed_primary_keys": m.FailedPrimaryKeys,
			"success_row_count":   m.SuccessRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataModificationResult_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failed_primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"success_row_count": types.Int64Type,
		},
	}
}

// GetFailedPrimaryKeys returns the value of the FailedPrimaryKeys field in DataModificationResult_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DataModificationResult_SdkV2) GetFailedPrimaryKeys(ctx context.Context) ([]types.String, bool) {
	if m.FailedPrimaryKeys.IsNull() || m.FailedPrimaryKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.FailedPrimaryKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailedPrimaryKeys sets the value of the FailedPrimaryKeys field in DataModificationResult_SdkV2.
func (m *DataModificationResult_SdkV2) SetFailedPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["failed_primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FailedPrimaryKeys = types.ListValueMust(t, vs)
}

type DeleteEndpointRequest_SdkV2 struct {
	// Full resource name of the endpoint to delete. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`
	Name types.String `tfsdk:"-"`
}

func (to *DeleteEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEndpointRequest_SdkV2) {
}

func (to *DeleteEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteEndpointRequest_SdkV2) {
}

func (m DeleteEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteIndexRequest_SdkV2 struct {
	// Full resource name of the index to delete. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name types.String `tfsdk:"-"`
}

func (to *DeleteIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteIndexRequest_SdkV2) {
}

func (to *DeleteIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteIndexRequest_SdkV2) {
}

func (m DeleteIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Specification for a Delta Sync index — the index is kept in sync with a
// source Delta table.
type DeltaSyncIndexSpec_SdkV2 struct {
	// [Optional] Select the columns to sync with the index. If left blank, all
	// columns from the source table are synced. The primary key column and
	// embedding source or vector column are always synced.
	ColumnsToSync types.List `tfsdk:"columns_to_sync"`
	// The columns that contain the embedding source.
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns"`
	// [Optional] Name of the Delta table to sync the index contents and
	// computed embeddings to.
	EmbeddingWritebackTable types.String `tfsdk:"embedding_writeback_table"`
	// The ID of the pipeline that is used to sync the index.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// Pipeline execution mode. Required on create — the backend rejects an
	// unset value. Storage Optimized endpoints accept only `TRIGGERED`;
	// Standard endpoints accept both. No explicit `stage` — a REQUIRED field
	// staged below its service would be dropped from combined specs while
	// remaining in `required`, tripping the OpenAPI required-vs-properties
	// consistency check. The field inherits the service's launch stage.
	PipelineType types.String `tfsdk:"pipeline_type"`
	// The full name of the source Delta table.
	SourceTable types.String `tfsdk:"source_table"`
}

func (to *DeltaSyncIndexSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaSyncIndexSpec_SdkV2) {
	if !from.ColumnsToSync.IsNull() && !from.ColumnsToSync.IsUnknown() && to.ColumnsToSync.IsNull() && len(from.ColumnsToSync.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToSync, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToSync = from.ColumnsToSync
	}
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() {
		if toEmbeddingSourceColumns, ok := to.GetEmbeddingSourceColumns(ctx); ok {
			if fromEmbeddingSourceColumns, ok := from.GetEmbeddingSourceColumns(ctx); ok {
				// Recursively sync the fields of each EmbeddingSourceColumns element by position.
				for i := range toEmbeddingSourceColumns {
					if i < len(fromEmbeddingSourceColumns) {
						toEmbeddingSourceColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEmbeddingSourceColumns[i])
					}
				}
				to.SetEmbeddingSourceColumns(ctx, toEmbeddingSourceColumns)
			}
		}
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() {
		if toEmbeddingVectorColumns, ok := to.GetEmbeddingVectorColumns(ctx); ok {
			if fromEmbeddingVectorColumns, ok := from.GetEmbeddingVectorColumns(ctx); ok {
				// Recursively sync the fields of each EmbeddingVectorColumns element by position.
				for i := range toEmbeddingVectorColumns {
					if i < len(fromEmbeddingVectorColumns) {
						toEmbeddingVectorColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEmbeddingVectorColumns[i])
					}
				}
				to.SetEmbeddingVectorColumns(ctx, toEmbeddingVectorColumns)
			}
		}
	}
}

func (to *DeltaSyncIndexSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeltaSyncIndexSpec_SdkV2) {
	if !from.ColumnsToSync.IsNull() && !from.ColumnsToSync.IsUnknown() && to.ColumnsToSync.IsNull() && len(from.ColumnsToSync.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToSync, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToSync = from.ColumnsToSync
	}
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() {
		if toEmbeddingSourceColumns, ok := to.GetEmbeddingSourceColumns(ctx); ok {
			if fromEmbeddingSourceColumns, ok := from.GetEmbeddingSourceColumns(ctx); ok {
				for i := range toEmbeddingSourceColumns {
					if i < len(fromEmbeddingSourceColumns) {
						toEmbeddingSourceColumns[i].SyncFieldsDuringRead(ctx, fromEmbeddingSourceColumns[i])
					}
				}
				to.SetEmbeddingSourceColumns(ctx, toEmbeddingSourceColumns)
			}
		}
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() {
		if toEmbeddingVectorColumns, ok := to.GetEmbeddingVectorColumns(ctx); ok {
			if fromEmbeddingVectorColumns, ok := from.GetEmbeddingVectorColumns(ctx); ok {
				for i := range toEmbeddingVectorColumns {
					if i < len(fromEmbeddingVectorColumns) {
						toEmbeddingVectorColumns[i].SyncFieldsDuringRead(ctx, fromEmbeddingVectorColumns[i])
					}
				}
				to.SetEmbeddingVectorColumns(ctx, toEmbeddingVectorColumns)
			}
		}
	}
}

func (m DeltaSyncIndexSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns_to_sync"] = attrs["columns_to_sync"].SetOptional()
	attrs["embedding_source_columns"] = attrs["embedding_source_columns"].SetOptional()
	attrs["embedding_vector_columns"] = attrs["embedding_vector_columns"].SetOptional()
	attrs["embedding_writeback_table"] = attrs["embedding_writeback_table"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetComputed()
	attrs["pipeline_type"] = attrs["pipeline_type"].SetRequired()
	attrs["source_table"] = attrs["source_table"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSyncIndexSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeltaSyncIndexSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns_to_sync":          reflect.TypeOf(types.String{}),
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn_SdkV2{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSyncIndexSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m DeltaSyncIndexSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns_to_sync":           m.ColumnsToSync,
			"embedding_source_columns":  m.EmbeddingSourceColumns,
			"embedding_vector_columns":  m.EmbeddingVectorColumns,
			"embedding_writeback_table": m.EmbeddingWritebackTable,
			"pipeline_id":               m.PipelineId,
			"pipeline_type":             m.PipelineType,
			"source_table":              m.SourceTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaSyncIndexSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns_to_sync": basetypes.ListType{
				ElemType: types.StringType,
			},
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn_SdkV2{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn_SdkV2{}.Type(ctx),
			},
			"embedding_writeback_table": types.StringType,
			"pipeline_id":               types.StringType,
			"pipeline_type":             types.StringType,
			"source_table":              types.StringType,
		},
	}
}

// GetColumnsToSync returns the value of the ColumnsToSync field in DeltaSyncIndexSpec_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncIndexSpec_SdkV2) GetColumnsToSync(ctx context.Context) ([]types.String, bool) {
	if m.ColumnsToSync.IsNull() || m.ColumnsToSync.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ColumnsToSync.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumnsToSync sets the value of the ColumnsToSync field in DeltaSyncIndexSpec_SdkV2.
func (m *DeltaSyncIndexSpec_SdkV2) SetColumnsToSync(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_sync"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToSync = types.ListValueMust(t, vs)
}

// GetEmbeddingSourceColumns returns the value of the EmbeddingSourceColumns field in DeltaSyncIndexSpec_SdkV2 as
// a slice of EmbeddingSourceColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncIndexSpec_SdkV2) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn_SdkV2, bool) {
	if m.EmbeddingSourceColumns.IsNull() || m.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn_SdkV2
	d := m.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DeltaSyncIndexSpec_SdkV2.
func (m *DeltaSyncIndexSpec_SdkV2) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DeltaSyncIndexSpec_SdkV2 as
// a slice of EmbeddingVectorColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncIndexSpec_SdkV2) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn_SdkV2, bool) {
	if m.EmbeddingVectorColumns.IsNull() || m.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn_SdkV2
	d := m.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DeltaSyncIndexSpec_SdkV2.
func (m *DeltaSyncIndexSpec_SdkV2) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

// Specification for a Direct Access index — the customer manages vectors and
// metadata directly.
type DirectAccessIndexSpec_SdkV2 struct {
	// The columns that contain the embedding source.
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns"`
	// The schema of the index in JSON format. Supported types are `integer`,
	// `long`, `float`, `double`, `boolean`, `string`, `date`, `timestamp`.
	// Supported types for vector columns: `array<float>`, `array<double>`.
	SchemaJson types.String `tfsdk:"schema_json"`
}

func (to *DirectAccessIndexSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DirectAccessIndexSpec_SdkV2) {
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() {
		if toEmbeddingSourceColumns, ok := to.GetEmbeddingSourceColumns(ctx); ok {
			if fromEmbeddingSourceColumns, ok := from.GetEmbeddingSourceColumns(ctx); ok {
				// Recursively sync the fields of each EmbeddingSourceColumns element by position.
				for i := range toEmbeddingSourceColumns {
					if i < len(fromEmbeddingSourceColumns) {
						toEmbeddingSourceColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEmbeddingSourceColumns[i])
					}
				}
				to.SetEmbeddingSourceColumns(ctx, toEmbeddingSourceColumns)
			}
		}
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() {
		if toEmbeddingVectorColumns, ok := to.GetEmbeddingVectorColumns(ctx); ok {
			if fromEmbeddingVectorColumns, ok := from.GetEmbeddingVectorColumns(ctx); ok {
				// Recursively sync the fields of each EmbeddingVectorColumns element by position.
				for i := range toEmbeddingVectorColumns {
					if i < len(fromEmbeddingVectorColumns) {
						toEmbeddingVectorColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEmbeddingVectorColumns[i])
					}
				}
				to.SetEmbeddingVectorColumns(ctx, toEmbeddingVectorColumns)
			}
		}
	}
}

func (to *DirectAccessIndexSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DirectAccessIndexSpec_SdkV2) {
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() {
		if toEmbeddingSourceColumns, ok := to.GetEmbeddingSourceColumns(ctx); ok {
			if fromEmbeddingSourceColumns, ok := from.GetEmbeddingSourceColumns(ctx); ok {
				for i := range toEmbeddingSourceColumns {
					if i < len(fromEmbeddingSourceColumns) {
						toEmbeddingSourceColumns[i].SyncFieldsDuringRead(ctx, fromEmbeddingSourceColumns[i])
					}
				}
				to.SetEmbeddingSourceColumns(ctx, toEmbeddingSourceColumns)
			}
		}
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() {
		if toEmbeddingVectorColumns, ok := to.GetEmbeddingVectorColumns(ctx); ok {
			if fromEmbeddingVectorColumns, ok := from.GetEmbeddingVectorColumns(ctx); ok {
				for i := range toEmbeddingVectorColumns {
					if i < len(fromEmbeddingVectorColumns) {
						toEmbeddingVectorColumns[i].SyncFieldsDuringRead(ctx, fromEmbeddingVectorColumns[i])
					}
				}
				to.SetEmbeddingVectorColumns(ctx, toEmbeddingVectorColumns)
			}
		}
	}
}

func (m DirectAccessIndexSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["embedding_source_columns"] = attrs["embedding_source_columns"].SetOptional()
	attrs["embedding_vector_columns"] = attrs["embedding_vector_columns"].SetOptional()
	attrs["schema_json"] = attrs["schema_json"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DirectAccessIndexSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DirectAccessIndexSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn_SdkV2{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectAccessIndexSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m DirectAccessIndexSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_source_columns": m.EmbeddingSourceColumns,
			"embedding_vector_columns": m.EmbeddingVectorColumns,
			"schema_json":              m.SchemaJson,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DirectAccessIndexSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn_SdkV2{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn_SdkV2{}.Type(ctx),
			},
			"schema_json": types.StringType,
		},
	}
}

// GetEmbeddingSourceColumns returns the value of the EmbeddingSourceColumns field in DirectAccessIndexSpec_SdkV2 as
// a slice of EmbeddingSourceColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DirectAccessIndexSpec_SdkV2) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn_SdkV2, bool) {
	if m.EmbeddingSourceColumns.IsNull() || m.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn_SdkV2
	d := m.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DirectAccessIndexSpec_SdkV2.
func (m *DirectAccessIndexSpec_SdkV2) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DirectAccessIndexSpec_SdkV2 as
// a slice of EmbeddingVectorColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DirectAccessIndexSpec_SdkV2) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn_SdkV2, bool) {
	if m.EmbeddingVectorColumns.IsNull() || m.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn_SdkV2
	d := m.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DirectAccessIndexSpec_SdkV2.
func (m *DirectAccessIndexSpec_SdkV2) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

// Name of an embedding source column and its associated embedding model
// endpoint.
type EmbeddingSourceColumn_SdkV2 struct {
	// Name of the embedding model endpoint, used by default for both ingestion
	// and querying.
	EmbeddingModelEndpoint types.String `tfsdk:"embedding_model_endpoint"`
	// Name of the embedding model endpoint which, if specified, is used for
	// querying (not ingestion).
	ModelEndpointNameForQuery types.String `tfsdk:"model_endpoint_name_for_query"`
	// Name of the source column.
	Name types.String `tfsdk:"name"`
}

func (to *EmbeddingSourceColumn_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmbeddingSourceColumn_SdkV2) {
}

func (to *EmbeddingSourceColumn_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EmbeddingSourceColumn_SdkV2) {
}

func (m EmbeddingSourceColumn_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["embedding_model_endpoint"] = attrs["embedding_model_endpoint"].SetOptional()
	attrs["model_endpoint_name_for_query"] = attrs["model_endpoint_name_for_query"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmbeddingSourceColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EmbeddingSourceColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingSourceColumn_SdkV2
// only implements ToObjectValue() and Type().
func (m EmbeddingSourceColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_model_endpoint":      m.EmbeddingModelEndpoint,
			"model_endpoint_name_for_query": m.ModelEndpointNameForQuery,
			"name":                          m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EmbeddingSourceColumn_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_model_endpoint":      types.StringType,
			"model_endpoint_name_for_query": types.StringType,
			"name":                          types.StringType,
		},
	}
}

// Name and dimension of an embedding vector column.
type EmbeddingVectorColumn_SdkV2 struct {
	// Dimension of the embedding vector.
	EmbeddingDimension types.Int64 `tfsdk:"embedding_dimension"`
	// Name of the column.
	Name types.String `tfsdk:"name"`
}

func (to *EmbeddingVectorColumn_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmbeddingVectorColumn_SdkV2) {
}

func (to *EmbeddingVectorColumn_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EmbeddingVectorColumn_SdkV2) {
}

func (m EmbeddingVectorColumn_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["embedding_dimension"] = attrs["embedding_dimension"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmbeddingVectorColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EmbeddingVectorColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingVectorColumn_SdkV2
// only implements ToObjectValue() and Type().
func (m EmbeddingVectorColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_dimension": m.EmbeddingDimension,
			"name":                m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EmbeddingVectorColumn_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_dimension": types.Int64Type,
			"name":                types.StringType,
		},
	}
}

// An AI Search endpoint — compute infrastructure that hosts AI Search indexes
// and serves queries against them. Customers create, query, and delete
// endpoints; the system manages provisioning, scaling, and health status.
type Endpoint_SdkV2 struct {
	// The user-selected budget policy id for the endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// Time the endpoint was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Creator of the endpoint
	Creator types.String `tfsdk:"creator"`
	// The custom tags assigned to the endpoint
	CustomTags types.List `tfsdk:"custom_tags"`
	// The budget policy id applied to the endpoint
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`
	// Current status of the endpoint
	EndpointStatus types.List `tfsdk:"endpoint_status"`
	// Type of endpoint. Required on create and immutable thereafter.
	EndpointType types.String `tfsdk:"endpoint_type"`
	// Unique identifier of the endpoint
	Id types.String `tfsdk:"id"`
	// Number of indexes on the endpoint
	IndexCount types.Int64 `tfsdk:"index_count"`
	// User who last updated the endpoint
	LastUpdatedUser types.String `tfsdk:"last_updated_user"`
	// Name of the AI Search endpoint. Server-assigned full resource path
	// (`workspaces/{workspace}/endpoints/{endpoint}`) on output. On create, the
	// user-supplied short name is conveyed via
	// `CreateEndpointRequest.endpoint_id`; the server composes the full `name`
	// and returns it on the response.
	Name types.String `tfsdk:"name"`
	// The client-supplied desired number of replicas for the endpoint, applied
	// at create/update time. Mutually exclusive with `target_qps`.
	ReplicaCount types.Int64 `tfsdk:"replica_count"`
	// Scaling information for the endpoint
	ScalingInfo types.List `tfsdk:"scaling_info"`
	// Target QPS for the endpoint. Mutually exclusive with `replica_count`.
	// Best-effort; the system does not guarantee this QPS will be achieved.
	TargetQps types.Int64 `tfsdk:"target_qps"`
	// Throughput information for the endpoint
	ThroughputInfo types.List `tfsdk:"throughput_info"`
	// Time the endpoint was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// The usage policy id applied to the endpoint.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

func (to *Endpoint_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Endpoint_SdkV2) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() {
		if toCustomTags, ok := to.GetCustomTags(ctx); ok {
			if fromCustomTags, ok := from.GetCustomTags(ctx); ok {
				// Recursively sync the fields of each CustomTags element by position.
				for i := range toCustomTags {
					if i < len(fromCustomTags) {
						toCustomTags[i].SyncFieldsDuringCreateOrUpdate(ctx, fromCustomTags[i])
					}
				}
				to.SetCustomTags(ctx, toCustomTags)
			}
		}
	}
	if !from.EndpointStatus.IsNull() && !from.EndpointStatus.IsUnknown() {
		if toEndpointStatus, ok := to.GetEndpointStatus(ctx); ok {
			if fromEndpointStatus, ok := from.GetEndpointStatus(ctx); ok {
				// Recursively sync the fields of EndpointStatus
				toEndpointStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromEndpointStatus)
				to.SetEndpointStatus(ctx, toEndpointStatus)
			}
		}
	}
	if !from.ScalingInfo.IsNull() && !from.ScalingInfo.IsUnknown() {
		if toScalingInfo, ok := to.GetScalingInfo(ctx); ok {
			if fromScalingInfo, ok := from.GetScalingInfo(ctx); ok {
				// Recursively sync the fields of ScalingInfo
				toScalingInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromScalingInfo)
				to.SetScalingInfo(ctx, toScalingInfo)
			}
		}
	}
	if !from.ThroughputInfo.IsNull() && !from.ThroughputInfo.IsUnknown() {
		if toThroughputInfo, ok := to.GetThroughputInfo(ctx); ok {
			if fromThroughputInfo, ok := from.GetThroughputInfo(ctx); ok {
				// Recursively sync the fields of ThroughputInfo
				toThroughputInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromThroughputInfo)
				to.SetThroughputInfo(ctx, toThroughputInfo)
			}
		}
	}
}

func (to *Endpoint_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Endpoint_SdkV2) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() {
		if toCustomTags, ok := to.GetCustomTags(ctx); ok {
			if fromCustomTags, ok := from.GetCustomTags(ctx); ok {
				for i := range toCustomTags {
					if i < len(fromCustomTags) {
						toCustomTags[i].SyncFieldsDuringRead(ctx, fromCustomTags[i])
					}
				}
				to.SetCustomTags(ctx, toCustomTags)
			}
		}
	}
	if !from.EndpointStatus.IsNull() && !from.EndpointStatus.IsUnknown() {
		if toEndpointStatus, ok := to.GetEndpointStatus(ctx); ok {
			if fromEndpointStatus, ok := from.GetEndpointStatus(ctx); ok {
				toEndpointStatus.SyncFieldsDuringRead(ctx, fromEndpointStatus)
				to.SetEndpointStatus(ctx, toEndpointStatus)
			}
		}
	}
	if !from.ScalingInfo.IsNull() && !from.ScalingInfo.IsUnknown() {
		if toScalingInfo, ok := to.GetScalingInfo(ctx); ok {
			if fromScalingInfo, ok := from.GetScalingInfo(ctx); ok {
				toScalingInfo.SyncFieldsDuringRead(ctx, fromScalingInfo)
				to.SetScalingInfo(ctx, toScalingInfo)
			}
		}
	}
	if !from.ThroughputInfo.IsNull() && !from.ThroughputInfo.IsUnknown() {
		if toThroughputInfo, ok := to.GetThroughputInfo(ctx); ok {
			if fromThroughputInfo, ok := from.GetThroughputInfo(ctx); ok {
				toThroughputInfo.SyncFieldsDuringRead(ctx, fromThroughputInfo)
				to.SetThroughputInfo(ctx, toThroughputInfo)
			}
		}
	}
}

func (m Endpoint_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetComputed()
	attrs["endpoint_status"] = attrs["endpoint_status"].SetComputed()
	attrs["endpoint_status"] = attrs["endpoint_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["endpoint_type"] = attrs["endpoint_type"].SetRequired()
	attrs["endpoint_type"] = attrs["endpoint_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetComputed()
	attrs["index_count"] = attrs["index_count"].SetComputed()
	attrs["last_updated_user"] = attrs["last_updated_user"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["replica_count"] = attrs["replica_count"].SetOptional()
	attrs["scaling_info"] = attrs["scaling_info"].SetComputed()
	attrs["scaling_info"] = attrs["scaling_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["target_qps"] = attrs["target_qps"].SetOptional()
	attrs["throughput_info"] = attrs["throughput_info"].SetComputed()
	attrs["throughput_info"] = attrs["throughput_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Endpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Endpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags":     reflect.TypeOf(CustomTag_SdkV2{}),
		"endpoint_status": reflect.TypeOf(EndpointStatus_SdkV2{}),
		"scaling_info":    reflect.TypeOf(EndpointScalingInfo_SdkV2{}),
		"throughput_info": reflect.TypeOf(EndpointThroughputInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Endpoint_SdkV2
// only implements ToObjectValue() and Type().
func (m Endpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id":           m.BudgetPolicyId,
			"create_time":                m.CreateTime,
			"creator":                    m.Creator,
			"custom_tags":                m.CustomTags,
			"effective_budget_policy_id": m.EffectiveBudgetPolicyId,
			"endpoint_status":            m.EndpointStatus,
			"endpoint_type":              m.EndpointType,
			"id":                         m.Id,
			"index_count":                m.IndexCount,
			"last_updated_user":          m.LastUpdatedUser,
			"name":                       m.Name,
			"replica_count":              m.ReplicaCount,
			"scaling_info":               m.ScalingInfo,
			"target_qps":                 m.TargetQps,
			"throughput_info":            m.ThroughputInfo,
			"update_time":                m.UpdateTime,
			"usage_policy_id":            m.UsagePolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Endpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"create_time":      timetypes.RFC3339{}.Type(ctx),
			"creator":          types.StringType,
			"custom_tags": basetypes.ListType{
				ElemType: CustomTag_SdkV2{}.Type(ctx),
			},
			"effective_budget_policy_id": types.StringType,
			"endpoint_status": basetypes.ListType{
				ElemType: EndpointStatus_SdkV2{}.Type(ctx),
			},
			"endpoint_type":     types.StringType,
			"id":                types.StringType,
			"index_count":       types.Int64Type,
			"last_updated_user": types.StringType,
			"name":              types.StringType,
			"replica_count":     types.Int64Type,
			"scaling_info": basetypes.ListType{
				ElemType: EndpointScalingInfo_SdkV2{}.Type(ctx),
			},
			"target_qps": types.Int64Type,
			"throughput_info": basetypes.ListType{
				ElemType: EndpointThroughputInfo_SdkV2{}.Type(ctx),
			},
			"update_time":     timetypes.RFC3339{}.Type(ctx),
			"usage_policy_id": types.StringType,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in Endpoint_SdkV2 as
// a slice of CustomTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint_SdkV2) GetCustomTags(ctx context.Context) ([]CustomTag_SdkV2, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag_SdkV2
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in Endpoint_SdkV2.
func (m *Endpoint_SdkV2) SetCustomTags(ctx context.Context, v []CustomTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

// GetEndpointStatus returns the value of the EndpointStatus field in Endpoint_SdkV2 as
// a EndpointStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint_SdkV2) GetEndpointStatus(ctx context.Context) (EndpointStatus_SdkV2, bool) {
	var e EndpointStatus_SdkV2
	if m.EndpointStatus.IsNull() || m.EndpointStatus.IsUnknown() {
		return e, false
	}
	var v []EndpointStatus_SdkV2
	d := m.EndpointStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEndpointStatus sets the value of the EndpointStatus field in Endpoint_SdkV2.
func (m *Endpoint_SdkV2) SetEndpointStatus(ctx context.Context, v EndpointStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoint_status"]
	m.EndpointStatus = types.ListValueMust(t, vs)
}

// GetScalingInfo returns the value of the ScalingInfo field in Endpoint_SdkV2 as
// a EndpointScalingInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint_SdkV2) GetScalingInfo(ctx context.Context) (EndpointScalingInfo_SdkV2, bool) {
	var e EndpointScalingInfo_SdkV2
	if m.ScalingInfo.IsNull() || m.ScalingInfo.IsUnknown() {
		return e, false
	}
	var v []EndpointScalingInfo_SdkV2
	d := m.ScalingInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetScalingInfo sets the value of the ScalingInfo field in Endpoint_SdkV2.
func (m *Endpoint_SdkV2) SetScalingInfo(ctx context.Context, v EndpointScalingInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scaling_info"]
	m.ScalingInfo = types.ListValueMust(t, vs)
}

// GetThroughputInfo returns the value of the ThroughputInfo field in Endpoint_SdkV2 as
// a EndpointThroughputInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint_SdkV2) GetThroughputInfo(ctx context.Context) (EndpointThroughputInfo_SdkV2, bool) {
	var e EndpointThroughputInfo_SdkV2
	if m.ThroughputInfo.IsNull() || m.ThroughputInfo.IsUnknown() {
		return e, false
	}
	var v []EndpointThroughputInfo_SdkV2
	d := m.ThroughputInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetThroughputInfo sets the value of the ThroughputInfo field in Endpoint_SdkV2.
func (m *Endpoint_SdkV2) SetThroughputInfo(ctx context.Context, v EndpointThroughputInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["throughput_info"]
	m.ThroughputInfo = types.ListValueMust(t, vs)
}

// Scaling information for a Storage Optimized endpoint — current scaling
// state and the requested QPS target the system is scaling toward.
type EndpointScalingInfo_SdkV2 struct {
	// The requested QPS target for the endpoint. Best-effort; the system does
	// not guarantee this QPS will be achieved.
	RequestedTargetQps types.Int64 `tfsdk:"requested_target_qps"`
	// The current state of the scaling change request.
	State types.String `tfsdk:"state"`
}

func (to *EndpointScalingInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointScalingInfo_SdkV2) {
}

func (to *EndpointScalingInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointScalingInfo_SdkV2) {
}

func (m EndpointScalingInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["requested_target_qps"] = attrs["requested_target_qps"].SetOptional()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointScalingInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointScalingInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointScalingInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointScalingInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"requested_target_qps": m.RequestedTargetQps,
			"state":                m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointScalingInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"requested_target_qps": types.Int64Type,
			"state":                types.StringType,
		},
	}
}

// Lifecycle and health state of an AI Search endpoint, along with any
// human-readable detail about that state.
type EndpointStatus_SdkV2 struct {
	// Human-readable detail about the endpoint's current state or the reason
	// for a state transition.
	Message types.String `tfsdk:"message"`
	// Current lifecycle state of the endpoint. See `State` for the meaning of
	// each value.
	State types.String `tfsdk:"state"`
}

func (to *EndpointStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointStatus_SdkV2) {
}

func (to *EndpointStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointStatus_SdkV2) {
}

func (m EndpointStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": m.Message,
			"state":   m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

// Throughput information for an AI Search endpoint, including requested and
// current concurrency settings.
type EndpointThroughputInfo_SdkV2 struct {
	// Additional information about the throughput change request
	ChangeRequestMessage types.String `tfsdk:"change_request_message"`
	// The state of the most recent throughput change request
	ChangeRequestState types.String `tfsdk:"change_request_state"`
	// The current concurrency (total CPU) allocated to the endpoint
	CurrentConcurrency types.Float64 `tfsdk:"current_concurrency"`
	// The current utilization of concurrency as a percentage (0-100)
	CurrentConcurrencyUtilizationPercentage types.Float64 `tfsdk:"current_concurrency_utilization_percentage"`
	// The current number of replicas allocated to the endpoint
	CurrentNumReplicas types.Int64 `tfsdk:"current_num_replicas"`
	// The maximum concurrency allowed for this endpoint
	MaximumConcurrencyAllowed types.Float64 `tfsdk:"maximum_concurrency_allowed"`
	// The minimum concurrency allowed for this endpoint
	MinimalConcurrencyAllowed types.Float64 `tfsdk:"minimal_concurrency_allowed"`
	// The requested concurrency (total CPU) for the endpoint
	RequestedConcurrency types.Float64 `tfsdk:"requested_concurrency"`
	// The requested number of replicas for the endpoint
	RequestedNumReplicas types.Int64 `tfsdk:"requested_num_replicas"`
}

func (to *EndpointThroughputInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointThroughputInfo_SdkV2) {
}

func (to *EndpointThroughputInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointThroughputInfo_SdkV2) {
}

func (m EndpointThroughputInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["change_request_message"] = attrs["change_request_message"].SetComputed()
	attrs["change_request_state"] = attrs["change_request_state"].SetComputed()
	attrs["current_concurrency"] = attrs["current_concurrency"].SetComputed()
	attrs["current_concurrency_utilization_percentage"] = attrs["current_concurrency_utilization_percentage"].SetComputed()
	attrs["current_num_replicas"] = attrs["current_num_replicas"].SetComputed()
	attrs["maximum_concurrency_allowed"] = attrs["maximum_concurrency_allowed"].SetOptional()
	attrs["minimal_concurrency_allowed"] = attrs["minimal_concurrency_allowed"].SetOptional()
	attrs["requested_concurrency"] = attrs["requested_concurrency"].SetOptional()
	attrs["requested_num_replicas"] = attrs["requested_num_replicas"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointThroughputInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointThroughputInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointThroughputInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointThroughputInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"change_request_message":                     m.ChangeRequestMessage,
			"change_request_state":                       m.ChangeRequestState,
			"current_concurrency":                        m.CurrentConcurrency,
			"current_concurrency_utilization_percentage": m.CurrentConcurrencyUtilizationPercentage,
			"current_num_replicas":                       m.CurrentNumReplicas,
			"maximum_concurrency_allowed":                m.MaximumConcurrencyAllowed,
			"minimal_concurrency_allowed":                m.MinimalConcurrencyAllowed,
			"requested_concurrency":                      m.RequestedConcurrency,
			"requested_num_replicas":                     m.RequestedNumReplicas,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointThroughputInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"change_request_message":                     types.StringType,
			"change_request_state":                       types.StringType,
			"current_concurrency":                        types.Float64Type,
			"current_concurrency_utilization_percentage": types.Float64Type,
			"current_num_replicas":                       types.Int64Type,
			"maximum_concurrency_allowed":                types.Float64Type,
			"minimal_concurrency_allowed":                types.Float64Type,
			"requested_concurrency":                      types.Float64Type,
			"requested_num_replicas":                     types.Int64Type,
		},
	}
}

// Facet aggregation rows returned by a query.
type FacetResultData_SdkV2 struct {
	// Facet rows; each row is `[facet_column_name, value_or_range, count]`.
	FacetArray types.List `tfsdk:"facet_array"`
	// Number of facet rows returned.
	FacetRowCount types.Int64 `tfsdk:"facet_row_count"`
}

func (to *FacetResultData_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FacetResultData_SdkV2) {
	if !from.FacetArray.IsNull() && !from.FacetArray.IsUnknown() && to.FacetArray.IsNull() && len(from.FacetArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FacetArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FacetArray = from.FacetArray
	}
}

func (to *FacetResultData_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FacetResultData_SdkV2) {
	if !from.FacetArray.IsNull() && !from.FacetArray.IsUnknown() && to.FacetArray.IsNull() && len(from.FacetArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FacetArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FacetArray = from.FacetArray
	}
}

func (m FacetResultData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["facet_array"] = attrs["facet_array"].SetComputed()
	attrs["facet_row_count"] = attrs["facet_row_count"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FacetResultData.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FacetResultData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"facet_array": reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FacetResultData_SdkV2
// only implements ToObjectValue() and Type().
func (m FacetResultData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"facet_array":     m.FacetArray,
			"facet_row_count": m.FacetRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FacetResultData_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"facet_array": basetypes.ListType{
				ElemType: basetypes.ListType{
					ElemType: jsontypes.NormalizedType{},
				},
			},
			"facet_row_count": types.Int64Type,
		},
	}
}

// GetFacetArray returns the value of the FacetArray field in FacetResultData_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *FacetResultData_SdkV2) GetFacetArray(ctx context.Context) ([]types.Object, bool) {
	if m.FacetArray.IsNull() || m.FacetArray.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.FacetArray.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFacetArray sets the value of the FacetArray field in FacetResultData_SdkV2.
func (m *FacetResultData_SdkV2) SetFacetArray(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["facet_array"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FacetArray = types.ListValueMust(t, vs)
}

type GetEndpointRequest_SdkV2 struct {
	// Full resource name of the endpoint. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`
	Name types.String `tfsdk:"-"`
}

func (to *GetEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEndpointRequest_SdkV2) {
}

func (to *GetEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetEndpointRequest_SdkV2) {
}

func (m GetEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetIndexRequest_SdkV2 struct {
	// Full resource name of the index. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name types.String `tfsdk:"-"`
}

func (to *GetIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetIndexRequest_SdkV2) {
}

func (to *GetIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetIndexRequest_SdkV2) {
}

func (m GetIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// An AI Search index — a searchable collection of vectors and metadata hosted
// on an AI Search endpoint. Indexes are children of endpoints; customers
// create, get, list, and delete them. The `{index}` segment of the resource
// name is the index's Unity Catalog table name.
type Index_SdkV2 struct {
	// Creator of the index.
	Creator types.String `tfsdk:"creator"`
	// Specification for a Delta Sync index. Set when `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec types.List `tfsdk:"delta_sync_index_spec"`
	// Specification for a Direct Access index. Set when `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec types.List `tfsdk:"direct_access_index_spec"`
	// Name of the endpoint associated with the index. Ignored on create — the
	// endpoint is taken from `CreateIndexRequest.parent`; populated only on
	// output.
	Endpoint types.String `tfsdk:"endpoint"`
	// The subtype of the index. Set on create and immutable thereafter.
	IndexSubtype types.String `tfsdk:"index_subtype"`
	// Type of index. Required on create and immutable thereafter.
	IndexType types.String `tfsdk:"index_type"`
	// Name of the AI Search index. Server-assigned full resource path
	// (`workspaces/{workspace}/endpoints/{endpoint}/indexes/{index}`) on
	// output, where `{index}` is the index's Unity Catalog table name. On
	// create, the user-supplied UC table name is conveyed via
	// `CreateIndexRequest.index_id`; the server composes the full `name` and
	// returns it on the response.
	Name types.String `tfsdk:"name"`
	// Primary key of the index. Set on create and immutable thereafter.
	PrimaryKey types.String `tfsdk:"primary_key"`
	// Current status of the index.
	Status types.List `tfsdk:"status"`
}

func (to *Index_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Index_SdkV2) {
	if !from.DeltaSyncIndexSpec.IsNull() && !from.DeltaSyncIndexSpec.IsUnknown() {
		if toDeltaSyncIndexSpec, ok := to.GetDeltaSyncIndexSpec(ctx); ok {
			if fromDeltaSyncIndexSpec, ok := from.GetDeltaSyncIndexSpec(ctx); ok {
				// Recursively sync the fields of DeltaSyncIndexSpec
				toDeltaSyncIndexSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDeltaSyncIndexSpec)
				to.SetDeltaSyncIndexSpec(ctx, toDeltaSyncIndexSpec)
			}
		}
	}
	if !from.DirectAccessIndexSpec.IsNull() && !from.DirectAccessIndexSpec.IsUnknown() {
		if toDirectAccessIndexSpec, ok := to.GetDirectAccessIndexSpec(ctx); ok {
			if fromDirectAccessIndexSpec, ok := from.GetDirectAccessIndexSpec(ctx); ok {
				// Recursively sync the fields of DirectAccessIndexSpec
				toDirectAccessIndexSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDirectAccessIndexSpec)
				to.SetDirectAccessIndexSpec(ctx, toDirectAccessIndexSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (to *Index_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Index_SdkV2) {
	if !from.DeltaSyncIndexSpec.IsNull() && !from.DeltaSyncIndexSpec.IsUnknown() {
		if toDeltaSyncIndexSpec, ok := to.GetDeltaSyncIndexSpec(ctx); ok {
			if fromDeltaSyncIndexSpec, ok := from.GetDeltaSyncIndexSpec(ctx); ok {
				toDeltaSyncIndexSpec.SyncFieldsDuringRead(ctx, fromDeltaSyncIndexSpec)
				to.SetDeltaSyncIndexSpec(ctx, toDeltaSyncIndexSpec)
			}
		}
	}
	if !from.DirectAccessIndexSpec.IsNull() && !from.DirectAccessIndexSpec.IsUnknown() {
		if toDirectAccessIndexSpec, ok := to.GetDirectAccessIndexSpec(ctx); ok {
			if fromDirectAccessIndexSpec, ok := from.GetDirectAccessIndexSpec(ctx); ok {
				toDirectAccessIndexSpec.SyncFieldsDuringRead(ctx, fromDirectAccessIndexSpec)
				to.SetDirectAccessIndexSpec(ctx, toDirectAccessIndexSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (m Index_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].SetOptional()
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].SetOptional()
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["endpoint"] = attrs["endpoint"].SetComputed()
	attrs["index_subtype"] = attrs["index_subtype"].SetOptional()
	attrs["index_subtype"] = attrs["index_subtype"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["index_type"] = attrs["index_type"].SetRequired()
	attrs["index_type"] = attrs["index_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["primary_key"] = attrs["primary_key"].SetRequired()
	attrs["primary_key"] = attrs["primary_key"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Index.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Index_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_sync_index_spec":    reflect.TypeOf(DeltaSyncIndexSpec_SdkV2{}),
		"direct_access_index_spec": reflect.TypeOf(DirectAccessIndexSpec_SdkV2{}),
		"status":                   reflect.TypeOf(IndexStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Index_SdkV2
// only implements ToObjectValue() and Type().
func (m Index_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator":                  m.Creator,
			"delta_sync_index_spec":    m.DeltaSyncIndexSpec,
			"direct_access_index_spec": m.DirectAccessIndexSpec,
			"endpoint":                 m.Endpoint,
			"index_subtype":            m.IndexSubtype,
			"index_type":               m.IndexType,
			"name":                     m.Name,
			"primary_key":              m.PrimaryKey,
			"status":                   m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Index_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creator": types.StringType,
			"delta_sync_index_spec": basetypes.ListType{
				ElemType: DeltaSyncIndexSpec_SdkV2{}.Type(ctx),
			},
			"direct_access_index_spec": basetypes.ListType{
				ElemType: DirectAccessIndexSpec_SdkV2{}.Type(ctx),
			},
			"endpoint":      types.StringType,
			"index_subtype": types.StringType,
			"index_type":    types.StringType,
			"name":          types.StringType,
			"primary_key":   types.StringType,
			"status": basetypes.ListType{
				ElemType: IndexStatus_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDeltaSyncIndexSpec returns the value of the DeltaSyncIndexSpec field in Index_SdkV2 as
// a DeltaSyncIndexSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Index_SdkV2) GetDeltaSyncIndexSpec(ctx context.Context) (DeltaSyncIndexSpec_SdkV2, bool) {
	var e DeltaSyncIndexSpec_SdkV2
	if m.DeltaSyncIndexSpec.IsNull() || m.DeltaSyncIndexSpec.IsUnknown() {
		return e, false
	}
	var v []DeltaSyncIndexSpec_SdkV2
	d := m.DeltaSyncIndexSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeltaSyncIndexSpec sets the value of the DeltaSyncIndexSpec field in Index_SdkV2.
func (m *Index_SdkV2) SetDeltaSyncIndexSpec(ctx context.Context, v DeltaSyncIndexSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["delta_sync_index_spec"]
	m.DeltaSyncIndexSpec = types.ListValueMust(t, vs)
}

// GetDirectAccessIndexSpec returns the value of the DirectAccessIndexSpec field in Index_SdkV2 as
// a DirectAccessIndexSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Index_SdkV2) GetDirectAccessIndexSpec(ctx context.Context) (DirectAccessIndexSpec_SdkV2, bool) {
	var e DirectAccessIndexSpec_SdkV2
	if m.DirectAccessIndexSpec.IsNull() || m.DirectAccessIndexSpec.IsUnknown() {
		return e, false
	}
	var v []DirectAccessIndexSpec_SdkV2
	d := m.DirectAccessIndexSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDirectAccessIndexSpec sets the value of the DirectAccessIndexSpec field in Index_SdkV2.
func (m *Index_SdkV2) SetDirectAccessIndexSpec(ctx context.Context, v DirectAccessIndexSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["direct_access_index_spec"]
	m.DirectAccessIndexSpec = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Index_SdkV2 as
// a IndexStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Index_SdkV2) GetStatus(ctx context.Context) (IndexStatus_SdkV2, bool) {
	var e IndexStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []IndexStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in Index_SdkV2.
func (m *Index_SdkV2) SetStatus(ctx context.Context, v IndexStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

// Lifecycle and health state of an AI Search index, along with human-readable
// detail about that state and basic indexing progress.
type IndexStatus_SdkV2 struct {
	// Index API URL used to perform operations on the index.
	IndexUrl types.String `tfsdk:"index_url"`
	// Number of rows indexed.
	IndexedRowCount types.Int64 `tfsdk:"indexed_row_count"`
	// Human-readable detail about the index's current state.
	Message types.String `tfsdk:"message"`
	// Whether the index is ready for search.
	Ready types.Bool `tfsdk:"ready"`
}

func (to *IndexStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IndexStatus_SdkV2) {
}

func (to *IndexStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IndexStatus_SdkV2) {
}

func (m IndexStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["index_url"] = attrs["index_url"].SetComputed()
	attrs["indexed_row_count"] = attrs["indexed_row_count"].SetComputed()
	attrs["message"] = attrs["message"].SetComputed()
	attrs["ready"] = attrs["ready"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IndexStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m IndexStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IndexStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m IndexStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_url":         m.IndexUrl,
			"indexed_row_count": m.IndexedRowCount,
			"message":           m.Message,
			"ready":             m.Ready,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IndexStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_url":         types.StringType,
			"indexed_row_count": types.Int64Type,
			"message":           types.StringType,
			"ready":             types.BoolType,
		},
	}
}

type ListEndpointsRequest_SdkV2 struct {
	// Best-effort upper bound on the number of results to return. Honored as an
	// upper bound by the shim: `page_size` only narrows the legacy backend's
	// response, never widens it, so the practical cap is `min(page_size,
	// legacy_fixed_page_size)`.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken types.String `tfsdk:"-"`
	// The Workspace that owns this collection of endpoints. Format:
	// `workspaces/{workspace_id}`
	Parent types.String `tfsdk:"-"`
}

func (to *ListEndpointsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsRequest_SdkV2) {
}

func (to *ListEndpointsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsRequest_SdkV2) {
}

func (m ListEndpointsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEndpointsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListEndpointsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response for ListEndpoints carrying the page of endpoints and an optional
// continuation token.
type ListEndpointsResponse_SdkV2 struct {
	// The endpoints in the workspace.
	Endpoints types.List `tfsdk:"endpoints"`
	// A token that can be used to get the next page of results. Empty when
	// there are no more results.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListEndpointsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsResponse_SdkV2) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() {
		if toEndpoints, ok := to.GetEndpoints(ctx); ok {
			if fromEndpoints, ok := from.GetEndpoints(ctx); ok {
				// Recursively sync the fields of each Endpoints element by position.
				for i := range toEndpoints {
					if i < len(fromEndpoints) {
						toEndpoints[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEndpoints[i])
					}
				}
				to.SetEndpoints(ctx, toEndpoints)
			}
		}
	}
}

func (to *ListEndpointsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsResponse_SdkV2) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() {
		if toEndpoints, ok := to.GetEndpoints(ctx); ok {
			if fromEndpoints, ok := from.GetEndpoints(ctx); ok {
				for i := range toEndpoints {
					if i < len(fromEndpoints) {
						toEndpoints[i].SyncFieldsDuringRead(ctx, fromEndpoints[i])
					}
				}
				to.SetEndpoints(ctx, toEndpoints)
			}
		}
	}
}

func (m ListEndpointsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoints"] = attrs["endpoints"].SetComputed()
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEndpointsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(Endpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListEndpointsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints":       m.Endpoints,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: Endpoint_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetEndpoints returns the value of the Endpoints field in ListEndpointsResponse_SdkV2 as
// a slice of Endpoint_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListEndpointsResponse_SdkV2) GetEndpoints(ctx context.Context) ([]Endpoint_SdkV2, bool) {
	if m.Endpoints.IsNull() || m.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []Endpoint_SdkV2
	d := m.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointsResponse_SdkV2.
func (m *ListEndpointsResponse_SdkV2) SetEndpoints(ctx context.Context, v []Endpoint_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Endpoints = types.ListValueMust(t, vs)
}

type ListIndexesRequest_SdkV2 struct {
	// Best-effort upper bound on the number of results to return. Honored as an
	// upper bound by the shim: `page_size` only narrows the legacy backend's
	// response, never widens it, so the practical cap is `min(page_size,
	// legacy_fixed_page_size)`.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken types.String `tfsdk:"-"`
	// The Endpoint that owns this collection of indexes. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`
	Parent types.String `tfsdk:"-"`
}

func (to *ListIndexesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListIndexesRequest_SdkV2) {
}

func (to *ListIndexesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListIndexesRequest_SdkV2) {
}

func (m ListIndexesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIndexesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListIndexesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIndexesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListIndexesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListIndexesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response for ListIndexes carrying the page of indexes and an optional
// continuation token.
type ListIndexesResponse_SdkV2 struct {
	// The indexes on the endpoint. The field is named `indexes` (not the
	// irregular plural `indices`) to satisfy core::0132, which derives the
	// response field name from the ListIndexes method.
	// core::0158::response-plural-first-field independently computes the
	// resource plural as `indices` and is satisfied via a scoped field
	// exception below.
	Indexes types.List `tfsdk:"indexes"`
	// A token that can be used to get the next page of results. Empty when
	// there are no more results.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListIndexesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListIndexesResponse_SdkV2) {
	if !from.Indexes.IsNull() && !from.Indexes.IsUnknown() && to.Indexes.IsNull() && len(from.Indexes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Indexes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Indexes = from.Indexes
	}
	if !from.Indexes.IsNull() && !from.Indexes.IsUnknown() {
		if toIndexes, ok := to.GetIndexes(ctx); ok {
			if fromIndexes, ok := from.GetIndexes(ctx); ok {
				// Recursively sync the fields of each Indexes element by position.
				for i := range toIndexes {
					if i < len(fromIndexes) {
						toIndexes[i].SyncFieldsDuringCreateOrUpdate(ctx, fromIndexes[i])
					}
				}
				to.SetIndexes(ctx, toIndexes)
			}
		}
	}
}

func (to *ListIndexesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListIndexesResponse_SdkV2) {
	if !from.Indexes.IsNull() && !from.Indexes.IsUnknown() && to.Indexes.IsNull() && len(from.Indexes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Indexes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Indexes = from.Indexes
	}
	if !from.Indexes.IsNull() && !from.Indexes.IsUnknown() {
		if toIndexes, ok := to.GetIndexes(ctx); ok {
			if fromIndexes, ok := from.GetIndexes(ctx); ok {
				for i := range toIndexes {
					if i < len(fromIndexes) {
						toIndexes[i].SyncFieldsDuringRead(ctx, fromIndexes[i])
					}
				}
				to.SetIndexes(ctx, toIndexes)
			}
		}
	}
}

func (m ListIndexesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["indexes"] = attrs["indexes"].SetComputed()
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIndexesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListIndexesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"indexes": reflect.TypeOf(Index_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIndexesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListIndexesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"indexes":         m.Indexes,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListIndexesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"indexes": basetypes.ListType{
				ElemType: Index_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetIndexes returns the value of the Indexes field in ListIndexesResponse_SdkV2 as
// a slice of Index_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListIndexesResponse_SdkV2) GetIndexes(ctx context.Context) ([]Index_SdkV2, bool) {
	if m.Indexes.IsNull() || m.Indexes.IsUnknown() {
		return nil, false
	}
	var v []Index_SdkV2
	d := m.Indexes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIndexes sets the value of the Indexes field in ListIndexesResponse_SdkV2.
func (m *ListIndexesResponse_SdkV2) SetIndexes(ctx context.Context, v []Index_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["indexes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Indexes = types.ListValueMust(t, vs)
}

// Request to query (search) an AI Search index. The legacy `num_results` count
// is exposed as `max_results`; v1 returns up to `max_results` rows in a single
// response (no cursor pagination — see the note on `max_results` below).
type QueryIndexRequest_SdkV2 struct {
	// Column names to include in each result row.
	Columns types.List `tfsdk:"columns"`
	// Columns whose values are sent to the reranker.
	ColumnsToRerank types.List `tfsdk:"columns_to_rerank"`
	// Facets to compute over the matched results (e.g. `"category TOP 5"`).
	Facets types.List `tfsdk:"facets"`
	// JSON string describing query filters (e.g. `{"id >": 5}`).
	FiltersJson types.String `tfsdk:"filters_json"`
	// Maximum number of results to return (the legacy `num_results`). Defaults
	// to 10.
	MaxResults types.Int64 `tfsdk:"max_results"`
	// Full resource name of the index to query. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name types.String `tfsdk:"-"`
	// Text columns to search for `query_text`. When empty, all text columns are
	// searched.
	QueryColumns types.List `tfsdk:"query_columns"`
	// Query text. Required for Delta Sync indexes that compute embeddings from
	// a model endpoint.
	QueryText types.String `tfsdk:"query_text"`
	// Query type: `ANN`, `HYBRID`, or `FULL_TEXT`. Defaults to `ANN`.
	QueryType types.String `tfsdk:"query_type"`
	// Query vector. Required for Direct Access indexes and Delta Sync indexes
	// with self-managed vectors.
	QueryVector types.List `tfsdk:"query_vector"`
	// If set, results are reranked before being returned.
	Reranker types.List `tfsdk:"reranker"`
	// Score threshold for the approximate nearest-neighbor search. Defaults to
	// 0.0.
	ScoreThreshold types.Float64 `tfsdk:"score_threshold"`
	// Sort clauses, e.g. `["rating DESC", "price ASC"]`. Overrides relevance
	// ordering.
	SortColumns types.List `tfsdk:"sort_columns"`
}

func (to *QueryIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryIndexRequest_SdkV2) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
	if !from.Facets.IsNull() && !from.Facets.IsUnknown() && to.Facets.IsNull() && len(from.Facets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Facets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Facets = from.Facets
	}
	if !from.QueryColumns.IsNull() && !from.QueryColumns.IsUnknown() && to.QueryColumns.IsNull() && len(from.QueryColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QueryColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QueryColumns = from.QueryColumns
	}
	if !from.QueryVector.IsNull() && !from.QueryVector.IsUnknown() && to.QueryVector.IsNull() && len(from.QueryVector.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QueryVector, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QueryVector = from.QueryVector
	}
	if !from.Reranker.IsNull() && !from.Reranker.IsUnknown() {
		if toReranker, ok := to.GetReranker(ctx); ok {
			if fromReranker, ok := from.GetReranker(ctx); ok {
				// Recursively sync the fields of Reranker
				toReranker.SyncFieldsDuringCreateOrUpdate(ctx, fromReranker)
				to.SetReranker(ctx, toReranker)
			}
		}
	}
	if !from.SortColumns.IsNull() && !from.SortColumns.IsUnknown() && to.SortColumns.IsNull() && len(from.SortColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SortColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SortColumns = from.SortColumns
	}
}

func (to *QueryIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryIndexRequest_SdkV2) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
	if !from.Facets.IsNull() && !from.Facets.IsUnknown() && to.Facets.IsNull() && len(from.Facets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Facets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Facets = from.Facets
	}
	if !from.QueryColumns.IsNull() && !from.QueryColumns.IsUnknown() && to.QueryColumns.IsNull() && len(from.QueryColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QueryColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QueryColumns = from.QueryColumns
	}
	if !from.QueryVector.IsNull() && !from.QueryVector.IsUnknown() && to.QueryVector.IsNull() && len(from.QueryVector.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QueryVector, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QueryVector = from.QueryVector
	}
	if !from.Reranker.IsNull() && !from.Reranker.IsUnknown() {
		if toReranker, ok := to.GetReranker(ctx); ok {
			if fromReranker, ok := from.GetReranker(ctx); ok {
				toReranker.SyncFieldsDuringRead(ctx, fromReranker)
				to.SetReranker(ctx, toReranker)
			}
		}
	}
	if !from.SortColumns.IsNull() && !from.SortColumns.IsUnknown() && to.SortColumns.IsNull() && len(from.SortColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SortColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SortColumns = from.SortColumns
	}
}

func (m QueryIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetRequired()
	attrs["columns_to_rerank"] = attrs["columns_to_rerank"].SetOptional()
	attrs["facets"] = attrs["facets"].SetOptional()
	attrs["filters_json"] = attrs["filters_json"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["query_columns"] = attrs["query_columns"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["query_type"] = attrs["query_type"].SetOptional()
	attrs["query_vector"] = attrs["query_vector"].SetOptional()
	attrs["reranker"] = attrs["reranker"].SetOptional()
	attrs["reranker"] = attrs["reranker"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["score_threshold"] = attrs["score_threshold"].SetOptional()
	attrs["sort_columns"] = attrs["sort_columns"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m QueryIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns":           reflect.TypeOf(types.String{}),
		"columns_to_rerank": reflect.TypeOf(types.String{}),
		"facets":            reflect.TypeOf(types.String{}),
		"query_columns":     reflect.TypeOf(types.String{}),
		"query_vector":      reflect.TypeOf(types.Float64{}),
		"reranker":          reflect.TypeOf(RerankerConfig_SdkV2{}),
		"sort_columns":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns":           m.Columns,
			"columns_to_rerank": m.ColumnsToRerank,
			"facets":            m.Facets,
			"filters_json":      m.FiltersJson,
			"max_results":       m.MaxResults,
			"name":              m.Name,
			"query_columns":     m.QueryColumns,
			"query_text":        m.QueryText,
			"query_type":        m.QueryType,
			"query_vector":      m.QueryVector,
			"reranker":          m.Reranker,
			"score_threshold":   m.ScoreThreshold,
			"sort_columns":      m.SortColumns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"columns_to_rerank": basetypes.ListType{
				ElemType: types.StringType,
			},
			"facets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"filters_json": types.StringType,
			"max_results":  types.Int64Type,
			"name":         types.StringType,
			"query_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"query_text": types.StringType,
			"query_type": types.StringType,
			"query_vector": basetypes.ListType{
				ElemType: types.Float64Type,
			},
			"reranker": basetypes.ListType{
				ElemType: RerankerConfig_SdkV2{}.Type(ctx),
			},
			"score_threshold": types.Float64Type,
			"sort_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetColumns returns the value of the Columns field in QueryIndexRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryIndexRequest_SdkV2) GetColumns(ctx context.Context) ([]types.String, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in QueryIndexRequest_SdkV2.
func (m *QueryIndexRequest_SdkV2) SetColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

// GetColumnsToRerank returns the value of the ColumnsToRerank field in QueryIndexRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryIndexRequest_SdkV2) GetColumnsToRerank(ctx context.Context) ([]types.String, bool) {
	if m.ColumnsToRerank.IsNull() || m.ColumnsToRerank.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ColumnsToRerank.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumnsToRerank sets the value of the ColumnsToRerank field in QueryIndexRequest_SdkV2.
func (m *QueryIndexRequest_SdkV2) SetColumnsToRerank(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_rerank"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToRerank = types.ListValueMust(t, vs)
}

// GetFacets returns the value of the Facets field in QueryIndexRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryIndexRequest_SdkV2) GetFacets(ctx context.Context) ([]types.String, bool) {
	if m.Facets.IsNull() || m.Facets.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Facets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFacets sets the value of the Facets field in QueryIndexRequest_SdkV2.
func (m *QueryIndexRequest_SdkV2) SetFacets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["facets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Facets = types.ListValueMust(t, vs)
}

// GetQueryColumns returns the value of the QueryColumns field in QueryIndexRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryIndexRequest_SdkV2) GetQueryColumns(ctx context.Context) ([]types.String, bool) {
	if m.QueryColumns.IsNull() || m.QueryColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.QueryColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryColumns sets the value of the QueryColumns field in QueryIndexRequest_SdkV2.
func (m *QueryIndexRequest_SdkV2) SetQueryColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.QueryColumns = types.ListValueMust(t, vs)
}

// GetQueryVector returns the value of the QueryVector field in QueryIndexRequest_SdkV2 as
// a slice of types.Float64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryIndexRequest_SdkV2) GetQueryVector(ctx context.Context) ([]types.Float64, bool) {
	if m.QueryVector.IsNull() || m.QueryVector.IsUnknown() {
		return nil, false
	}
	var v []types.Float64
	d := m.QueryVector.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryVector sets the value of the QueryVector field in QueryIndexRequest_SdkV2.
func (m *QueryIndexRequest_SdkV2) SetQueryVector(ctx context.Context, v []types.Float64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_vector"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.QueryVector = types.ListValueMust(t, vs)
}

// GetReranker returns the value of the Reranker field in QueryIndexRequest_SdkV2 as
// a RerankerConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryIndexRequest_SdkV2) GetReranker(ctx context.Context) (RerankerConfig_SdkV2, bool) {
	var e RerankerConfig_SdkV2
	if m.Reranker.IsNull() || m.Reranker.IsUnknown() {
		return e, false
	}
	var v []RerankerConfig_SdkV2
	d := m.Reranker.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetReranker sets the value of the Reranker field in QueryIndexRequest_SdkV2.
func (m *QueryIndexRequest_SdkV2) SetReranker(ctx context.Context, v RerankerConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["reranker"]
	m.Reranker = types.ListValueMust(t, vs)
}

// GetSortColumns returns the value of the SortColumns field in QueryIndexRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryIndexRequest_SdkV2) GetSortColumns(ctx context.Context) ([]types.String, bool) {
	if m.SortColumns.IsNull() || m.SortColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SortColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSortColumns sets the value of the SortColumns field in QueryIndexRequest_SdkV2.
func (m *QueryIndexRequest_SdkV2) SetSortColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sort_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SortColumns = types.ListValueMust(t, vs)
}

// Response for QueryIndex carrying the matched rows and their column metadata.
type QueryIndexResponse_SdkV2 struct {
	// Facet aggregation rows, when facets were requested.
	FacetResult types.List `tfsdk:"facet_result"`
	// Metadata describing the result columns.
	Manifest types.List `tfsdk:"manifest"`
	// The matched result rows.
	Result types.List `tfsdk:"result"`
}

func (to *QueryIndexResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryIndexResponse_SdkV2) {
	if !from.FacetResult.IsNull() && !from.FacetResult.IsUnknown() {
		if toFacetResult, ok := to.GetFacetResult(ctx); ok {
			if fromFacetResult, ok := from.GetFacetResult(ctx); ok {
				// Recursively sync the fields of FacetResult
				toFacetResult.SyncFieldsDuringCreateOrUpdate(ctx, fromFacetResult)
				to.SetFacetResult(ctx, toFacetResult)
			}
		}
	}
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				// Recursively sync the fields of Manifest
				toManifest.SyncFieldsDuringCreateOrUpdate(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				// Recursively sync the fields of Result
				toResult.SyncFieldsDuringCreateOrUpdate(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (to *QueryIndexResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryIndexResponse_SdkV2) {
	if !from.FacetResult.IsNull() && !from.FacetResult.IsUnknown() {
		if toFacetResult, ok := to.GetFacetResult(ctx); ok {
			if fromFacetResult, ok := from.GetFacetResult(ctx); ok {
				toFacetResult.SyncFieldsDuringRead(ctx, fromFacetResult)
				to.SetFacetResult(ctx, toFacetResult)
			}
		}
	}
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				toManifest.SyncFieldsDuringRead(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				toResult.SyncFieldsDuringRead(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (m QueryIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["facet_result"] = attrs["facet_result"].SetComputed()
	attrs["facet_result"] = attrs["facet_result"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["manifest"] = attrs["manifest"].SetComputed()
	attrs["manifest"] = attrs["manifest"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["result"] = attrs["result"].SetComputed()
	attrs["result"] = attrs["result"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m QueryIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"facet_result": reflect.TypeOf(FacetResultData_SdkV2{}),
		"manifest":     reflect.TypeOf(ResultManifest_SdkV2{}),
		"result":       reflect.TypeOf(ResultData_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"facet_result": m.FacetResult,
			"manifest":     m.Manifest,
			"result":       m.Result,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"facet_result": basetypes.ListType{
				ElemType: FacetResultData_SdkV2{}.Type(ctx),
			},
			"manifest": basetypes.ListType{
				ElemType: ResultManifest_SdkV2{}.Type(ctx),
			},
			"result": basetypes.ListType{
				ElemType: ResultData_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFacetResult returns the value of the FacetResult field in QueryIndexResponse_SdkV2 as
// a FacetResultData_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryIndexResponse_SdkV2) GetFacetResult(ctx context.Context) (FacetResultData_SdkV2, bool) {
	var e FacetResultData_SdkV2
	if m.FacetResult.IsNull() || m.FacetResult.IsUnknown() {
		return e, false
	}
	var v []FacetResultData_SdkV2
	d := m.FacetResult.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFacetResult sets the value of the FacetResult field in QueryIndexResponse_SdkV2.
func (m *QueryIndexResponse_SdkV2) SetFacetResult(ctx context.Context, v FacetResultData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["facet_result"]
	m.FacetResult = types.ListValueMust(t, vs)
}

// GetManifest returns the value of the Manifest field in QueryIndexResponse_SdkV2 as
// a ResultManifest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryIndexResponse_SdkV2) GetManifest(ctx context.Context) (ResultManifest_SdkV2, bool) {
	var e ResultManifest_SdkV2
	if m.Manifest.IsNull() || m.Manifest.IsUnknown() {
		return e, false
	}
	var v []ResultManifest_SdkV2
	d := m.Manifest.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManifest sets the value of the Manifest field in QueryIndexResponse_SdkV2.
func (m *QueryIndexResponse_SdkV2) SetManifest(ctx context.Context, v ResultManifest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["manifest"]
	m.Manifest = types.ListValueMust(t, vs)
}

// GetResult returns the value of the Result field in QueryIndexResponse_SdkV2 as
// a ResultData_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryIndexResponse_SdkV2) GetResult(ctx context.Context) (ResultData_SdkV2, bool) {
	var e ResultData_SdkV2
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v []ResultData_SdkV2
	d := m.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in QueryIndexResponse_SdkV2.
func (m *QueryIndexResponse_SdkV2) SetResult(ctx context.Context, v ResultData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	m.Result = types.ListValueMust(t, vs)
}

// Request to remove rows by primary key from a Direct Access AI Search index.
// Named RemoveData (not DeleteData) so the linter does not classify it as a
// standard AIP-135 Delete method — it deletes rows within an index, not the
// index resource.
type RemoveDataRequest_SdkV2 struct {
	// Full resource name of the index. Must be a Direct Access index. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name types.String `tfsdk:"-"`
	// Primary keys of the rows to remove.
	PrimaryKeys types.List `tfsdk:"primary_keys"`
}

func (to *RemoveDataRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RemoveDataRequest_SdkV2) {
}

func (to *RemoveDataRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RemoveDataRequest_SdkV2) {
}

func (m RemoveDataRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["primary_keys"] = attrs["primary_keys"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveDataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RemoveDataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveDataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RemoveDataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":         m.Name,
			"primary_keys": m.PrimaryKeys,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RemoveDataRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPrimaryKeys returns the value of the PrimaryKeys field in RemoveDataRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RemoveDataRequest_SdkV2) GetPrimaryKeys(ctx context.Context) ([]types.String, bool) {
	if m.PrimaryKeys.IsNull() || m.PrimaryKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.PrimaryKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrimaryKeys sets the value of the PrimaryKeys field in RemoveDataRequest_SdkV2.
func (m *RemoveDataRequest_SdkV2) SetPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrimaryKeys = types.ListValueMust(t, vs)
}

// Response for RemoveData.
type RemoveDataResponse_SdkV2 struct {
	// Per-row outcome of the delete.
	Result types.List `tfsdk:"result"`
	// Overall status of the delete.
	Status types.String `tfsdk:"status"`
}

func (to *RemoveDataResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RemoveDataResponse_SdkV2) {
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				// Recursively sync the fields of Result
				toResult.SyncFieldsDuringCreateOrUpdate(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (to *RemoveDataResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RemoveDataResponse_SdkV2) {
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				toResult.SyncFieldsDuringRead(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (m RemoveDataResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["result"] = attrs["result"].SetComputed()
	attrs["result"] = attrs["result"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RemoveDataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RemoveDataResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"result": reflect.TypeOf(DataModificationResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RemoveDataResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RemoveDataResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result": m.Result,
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RemoveDataResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"result": basetypes.ListType{
				ElemType: DataModificationResult_SdkV2{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

// GetResult returns the value of the Result field in RemoveDataResponse_SdkV2 as
// a DataModificationResult_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RemoveDataResponse_SdkV2) GetResult(ctx context.Context) (DataModificationResult_SdkV2, bool) {
	var e DataModificationResult_SdkV2
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v []DataModificationResult_SdkV2
	d := m.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in RemoveDataResponse_SdkV2.
func (m *RemoveDataResponse_SdkV2) SetResult(ctx context.Context, v DataModificationResult_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	m.Result = types.ListValueMust(t, vs)
}

// Configuration for reranking query results with a reranker model.
type RerankerConfig_SdkV2 struct {
	// Reranker identifier: "databricks_reranker" for the base model, or a Model
	// Serving endpoint name when `model_type` is MODEL_TYPE_FINETUNED.
	Model types.String `tfsdk:"model"`
	// Discriminator for how `model` is interpreted.
	ModelType types.String `tfsdk:"model_type"`
	// Parameters controlling reranking.
	Parameters types.List `tfsdk:"parameters"`
}

func (to *RerankerConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RerankerConfig_SdkV2) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() {
		if toParameters, ok := to.GetParameters(ctx); ok {
			if fromParameters, ok := from.GetParameters(ctx); ok {
				// Recursively sync the fields of Parameters
				toParameters.SyncFieldsDuringCreateOrUpdate(ctx, fromParameters)
				to.SetParameters(ctx, toParameters)
			}
		}
	}
}

func (to *RerankerConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RerankerConfig_SdkV2) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() {
		if toParameters, ok := to.GetParameters(ctx); ok {
			if fromParameters, ok := from.GetParameters(ctx); ok {
				toParameters.SyncFieldsDuringRead(ctx, fromParameters)
				to.SetParameters(ctx, toParameters)
			}
		}
	}
}

func (m RerankerConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model"] = attrs["model"].SetOptional()
	attrs["model_type"] = attrs["model_type"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["parameters"] = attrs["parameters"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RerankerConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RerankerConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(RerankerConfigRerankerParameters_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RerankerConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m RerankerConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model":      m.Model,
			"model_type": m.ModelType,
			"parameters": m.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RerankerConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model":      types.StringType,
			"model_type": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: RerankerConfigRerankerParameters_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetParameters returns the value of the Parameters field in RerankerConfig_SdkV2 as
// a RerankerConfigRerankerParameters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RerankerConfig_SdkV2) GetParameters(ctx context.Context) (RerankerConfigRerankerParameters_SdkV2, bool) {
	var e RerankerConfigRerankerParameters_SdkV2
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return e, false
	}
	var v []RerankerConfigRerankerParameters_SdkV2
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetParameters sets the value of the Parameters field in RerankerConfig_SdkV2.
func (m *RerankerConfig_SdkV2) SetParameters(ctx context.Context, v RerankerConfigRerankerParameters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	m.Parameters = types.ListValueMust(t, vs)
}

// Parameters controlling how the reranker processes results.
type RerankerConfigRerankerParameters_SdkV2 struct {
	// Columns whose values are concatenated and sent to the reranker.
	ColumnsToRerank types.List `tfsdk:"columns_to_rerank"`
}

func (to *RerankerConfigRerankerParameters_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RerankerConfigRerankerParameters_SdkV2) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
}

func (to *RerankerConfigRerankerParameters_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RerankerConfigRerankerParameters_SdkV2) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
}

func (m RerankerConfigRerankerParameters_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns_to_rerank"] = attrs["columns_to_rerank"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RerankerConfigRerankerParameters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RerankerConfigRerankerParameters_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns_to_rerank": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RerankerConfigRerankerParameters_SdkV2
// only implements ToObjectValue() and Type().
func (m RerankerConfigRerankerParameters_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns_to_rerank": m.ColumnsToRerank,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RerankerConfigRerankerParameters_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns_to_rerank": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetColumnsToRerank returns the value of the ColumnsToRerank field in RerankerConfigRerankerParameters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RerankerConfigRerankerParameters_SdkV2) GetColumnsToRerank(ctx context.Context) ([]types.String, bool) {
	if m.ColumnsToRerank.IsNull() || m.ColumnsToRerank.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ColumnsToRerank.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumnsToRerank sets the value of the ColumnsToRerank field in RerankerConfigRerankerParameters_SdkV2.
func (m *RerankerConfigRerankerParameters_SdkV2) SetColumnsToRerank(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_rerank"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToRerank = types.ListValueMust(t, vs)
}

// The rows of a query result set.
type ResultData_SdkV2 struct {
	// Result rows; each row is a list of column values aligned with the
	// manifest columns.
	DataArray types.List `tfsdk:"data_array"`
	// Number of rows in the result set.
	RowCount types.Int64 `tfsdk:"row_count"`
}

func (to *ResultData_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultData_SdkV2) {
	if !from.DataArray.IsNull() && !from.DataArray.IsUnknown() && to.DataArray.IsNull() && len(from.DataArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataArray = from.DataArray
	}
}

func (to *ResultData_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResultData_SdkV2) {
	if !from.DataArray.IsNull() && !from.DataArray.IsUnknown() && to.DataArray.IsNull() && len(from.DataArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataArray = from.DataArray
	}
}

func (m ResultData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_array"] = attrs["data_array"].SetComputed()
	attrs["row_count"] = attrs["row_count"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultData.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResultData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_array": reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultData_SdkV2
// only implements ToObjectValue() and Type().
func (m ResultData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_array": m.DataArray,
			"row_count":  m.RowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultData_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_array": basetypes.ListType{
				ElemType: basetypes.ListType{
					ElemType: jsontypes.NormalizedType{},
				},
			},
			"row_count": types.Int64Type,
		},
	}
}

// GetDataArray returns the value of the DataArray field in ResultData_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultData_SdkV2) GetDataArray(ctx context.Context) ([]types.Object, bool) {
	if m.DataArray.IsNull() || m.DataArray.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.DataArray.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataArray sets the value of the DataArray field in ResultData_SdkV2.
func (m *ResultData_SdkV2) SetDataArray(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_array"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DataArray = types.ListValueMust(t, vs)
}

// Metadata describing the columns of a query result set.
type ResultManifest_SdkV2 struct {
	// Number of columns in the result set.
	ColumnCount types.Int64 `tfsdk:"column_count"`
	// Information about each column in the result set.
	Columns types.List `tfsdk:"columns"`
	// Number of columns in the facet result.
	FacetColumnCount types.Int64 `tfsdk:"facet_column_count"`
	// Information about each facet column.
	FacetColumns types.List `tfsdk:"facet_columns"`
}

func (to *ResultManifest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultManifest_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() {
		if toColumns, ok := to.GetColumns(ctx); ok {
			if fromColumns, ok := from.GetColumns(ctx); ok {
				// Recursively sync the fields of each Columns element by position.
				for i := range toColumns {
					if i < len(fromColumns) {
						toColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromColumns[i])
					}
				}
				to.SetColumns(ctx, toColumns)
			}
		}
	}
	if !from.FacetColumns.IsNull() && !from.FacetColumns.IsUnknown() && to.FacetColumns.IsNull() && len(from.FacetColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FacetColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FacetColumns = from.FacetColumns
	}
	if !from.FacetColumns.IsNull() && !from.FacetColumns.IsUnknown() {
		if toFacetColumns, ok := to.GetFacetColumns(ctx); ok {
			if fromFacetColumns, ok := from.GetFacetColumns(ctx); ok {
				// Recursively sync the fields of each FacetColumns element by position.
				for i := range toFacetColumns {
					if i < len(fromFacetColumns) {
						toFacetColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromFacetColumns[i])
					}
				}
				to.SetFacetColumns(ctx, toFacetColumns)
			}
		}
	}
}

func (to *ResultManifest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResultManifest_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() {
		if toColumns, ok := to.GetColumns(ctx); ok {
			if fromColumns, ok := from.GetColumns(ctx); ok {
				for i := range toColumns {
					if i < len(fromColumns) {
						toColumns[i].SyncFieldsDuringRead(ctx, fromColumns[i])
					}
				}
				to.SetColumns(ctx, toColumns)
			}
		}
	}
	if !from.FacetColumns.IsNull() && !from.FacetColumns.IsUnknown() && to.FacetColumns.IsNull() && len(from.FacetColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FacetColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FacetColumns = from.FacetColumns
	}
	if !from.FacetColumns.IsNull() && !from.FacetColumns.IsUnknown() {
		if toFacetColumns, ok := to.GetFacetColumns(ctx); ok {
			if fromFacetColumns, ok := from.GetFacetColumns(ctx); ok {
				for i := range toFacetColumns {
					if i < len(fromFacetColumns) {
						toFacetColumns[i].SyncFieldsDuringRead(ctx, fromFacetColumns[i])
					}
				}
				to.SetFacetColumns(ctx, toFacetColumns)
			}
		}
	}
}

func (m ResultManifest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column_count"] = attrs["column_count"].SetComputed()
	attrs["columns"] = attrs["columns"].SetComputed()
	attrs["facet_column_count"] = attrs["facet_column_count"].SetComputed()
	attrs["facet_columns"] = attrs["facet_columns"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultManifest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResultManifest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns":       reflect.TypeOf(ColumnInfo_SdkV2{}),
		"facet_columns": reflect.TypeOf(ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResultManifest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count":       m.ColumnCount,
			"columns":            m.Columns,
			"facet_column_count": m.FacetColumnCount,
			"facet_columns":      m.FacetColumns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultManifest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column_count": types.Int64Type,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo_SdkV2{}.Type(ctx),
			},
			"facet_column_count": types.Int64Type,
			"facet_columns": basetypes.ListType{
				ElemType: ColumnInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in ResultManifest_SdkV2 as
// a slice of ColumnInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultManifest_SdkV2) GetColumns(ctx context.Context) ([]ColumnInfo_SdkV2, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo_SdkV2
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in ResultManifest_SdkV2.
func (m *ResultManifest_SdkV2) SetColumns(ctx context.Context, v []ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

// GetFacetColumns returns the value of the FacetColumns field in ResultManifest_SdkV2 as
// a slice of ColumnInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultManifest_SdkV2) GetFacetColumns(ctx context.Context) ([]ColumnInfo_SdkV2, bool) {
	if m.FacetColumns.IsNull() || m.FacetColumns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo_SdkV2
	d := m.FacetColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFacetColumns sets the value of the FacetColumns field in ResultManifest_SdkV2.
func (m *ResultManifest_SdkV2) SetFacetColumns(ctx context.Context, v []ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["facet_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FacetColumns = types.ListValueMust(t, vs)
}

// Request to scan (paginate over) the rows of an AI Search index. Models the
// legacy `num_results` / `last_primary_key` cursor as AIP-158 `page_size` /
// `page_token`.
type ScanIndexRequest_SdkV2 struct {
	// Full resource name of the index to scan. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name types.String `tfsdk:"-"`
	// Maximum number of rows to return in this page.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Page token from a previous response; if unset, scanning starts from the
	// beginning.
	PageToken types.String `tfsdk:"page_token"`
}

func (to *ScanIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ScanIndexRequest_SdkV2) {
}

func (to *ScanIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ScanIndexRequest_SdkV2) {
}

func (m ScanIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ScanIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ScanIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ScanIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ScanIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":       m.Name,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ScanIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":       types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Response for ScanIndex carrying a page of rows and an optional continuation
// token.
type ScanIndexResponse_SdkV2 struct {
	// The rows in this page, each a struct of column name to value.
	Data types.List `tfsdk:"data"`
	// Token for the next page; empty when the scan is exhausted.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ScanIndexResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ScanIndexResponse_SdkV2) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
}

func (to *ScanIndexResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ScanIndexResponse_SdkV2) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
}

func (m ScanIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data"] = attrs["data"].SetComputed()
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ScanIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ScanIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data": reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ScanIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ScanIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":            m.Data,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ScanIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data": basetypes.ListType{
				ElemType: basetypes.MapType{
					ElemType: jsontypes.NormalizedType{},
				},
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetData returns the value of the Data field in ScanIndexResponse_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *ScanIndexResponse_SdkV2) GetData(ctx context.Context) ([]types.Object, bool) {
	if m.Data.IsNull() || m.Data.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in ScanIndexResponse_SdkV2.
func (m *ScanIndexResponse_SdkV2) SetData(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Data = types.ListValueMust(t, vs)
}

// Request to synchronize a Delta Sync AI Search index with its source Delta
// table.
type SyncIndexRequest_SdkV2 struct {
	// Full resource name of the index to synchronize. Must be a Delta Sync
	// index. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name types.String `tfsdk:"-"`
}

func (to *SyncIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncIndexRequest_SdkV2) {
}

func (to *SyncIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncIndexRequest_SdkV2) {
}

func (m SyncIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Response for SyncIndex. Empty today; reserved so future sync metadata (e.g.
// an operation handle) can be added without breaking the wire contract.
type SyncIndexResponse_SdkV2 struct {
}

func (to *SyncIndexResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncIndexResponse_SdkV2) {
}

func (to *SyncIndexResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncIndexResponse_SdkV2) {
}

func (m SyncIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SyncIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateEndpointRequest_SdkV2 struct {
	// The Endpoint resource to update. `endpoint.name` carries the full
	// resource path.
	Endpoint types.List `tfsdk:"endpoint"`
	// Name of the AI Search endpoint. Server-assigned full resource path
	// (`workspaces/{workspace}/endpoints/{endpoint}`) on output. On create, the
	// user-supplied short name is conveyed via
	// `CreateEndpointRequest.endpoint_id`; the server composes the full `name`
	// and returns it on the response.
	Name types.String `tfsdk:"-"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				// Recursively sync the fields of Endpoint
				toEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (to *UpdateEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m UpdateEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
	attrs["endpoint"] = attrs["endpoint"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint":    m.Endpoint,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint": basetypes.ListType{
				ElemType: Endpoint_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in UpdateEndpointRequest_SdkV2 as
// a Endpoint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEndpointRequest_SdkV2) GetEndpoint(ctx context.Context) (Endpoint_SdkV2, bool) {
	var e Endpoint_SdkV2
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v []Endpoint_SdkV2
	d := m.Endpoint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEndpoint sets the value of the Endpoint field in UpdateEndpointRequest_SdkV2.
func (m *UpdateEndpointRequest_SdkV2) SetEndpoint(ctx context.Context, v Endpoint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoint"]
	m.Endpoint = types.ListValueMust(t, vs)
}

// Request to upsert rows into a Direct Access AI Search index.
type UpsertDataRequest_SdkV2 struct {
	// JSON document describing the rows to upsert.
	InputsJson types.String `tfsdk:"inputs_json"`
	// Full resource name of the index. Must be a Direct Access index. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name types.String `tfsdk:"-"`
}

func (to *UpsertDataRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpsertDataRequest_SdkV2) {
}

func (to *UpsertDataRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpsertDataRequest_SdkV2) {
}

func (m UpsertDataRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inputs_json"] = attrs["inputs_json"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertDataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpsertDataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpsertDataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inputs_json": m.InputsJson,
			"name":        m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpsertDataRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inputs_json": types.StringType,
			"name":        types.StringType,
		},
	}
}

// Response for UpsertData.
type UpsertDataResponse_SdkV2 struct {
	// Per-row outcome of the upsert.
	Result types.List `tfsdk:"result"`
	// Overall status of the upsert.
	Status types.String `tfsdk:"status"`
}

func (to *UpsertDataResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpsertDataResponse_SdkV2) {
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				// Recursively sync the fields of Result
				toResult.SyncFieldsDuringCreateOrUpdate(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (to *UpsertDataResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpsertDataResponse_SdkV2) {
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				toResult.SyncFieldsDuringRead(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (m UpsertDataResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["result"] = attrs["result"].SetComputed()
	attrs["result"] = attrs["result"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertDataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpsertDataResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"result": reflect.TypeOf(DataModificationResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpsertDataResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result": m.Result,
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpsertDataResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"result": basetypes.ListType{
				ElemType: DataModificationResult_SdkV2{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

// GetResult returns the value of the Result field in UpsertDataResponse_SdkV2 as
// a DataModificationResult_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpsertDataResponse_SdkV2) GetResult(ctx context.Context) (DataModificationResult_SdkV2, bool) {
	var e DataModificationResult_SdkV2
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v []DataModificationResult_SdkV2
	d := m.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in UpsertDataResponse_SdkV2.
func (m *UpsertDataResponse_SdkV2) SetResult(ctx context.Context, v DataModificationResult_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	m.Result = types.ListValueMust(t, vs)
}
