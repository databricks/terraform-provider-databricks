// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package vectorsearch_tf

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ColumnInfo struct {
	// Name of the column.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnInfo) {
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringRead(existingState ColumnInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ColumnInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ColumnInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo
// only implements ToObjectValue() and Type().
func (o ColumnInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ColumnInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type CreateEndpoint struct {
	// Type of endpoint.
	EndpointType types.String `tfsdk:"endpoint_type" tf:""`
	// Name of endpoint
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *CreateEndpoint) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateEndpoint) {
}

func (newState *CreateEndpoint) SyncEffectiveFieldsDuringRead(existingState CreateEndpoint) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateEndpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateEndpoint
// only implements ToObjectValue() and Type().
func (o CreateEndpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_type": o.EndpointType,
			"name":          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_type": types.StringType,
			"name":          types.StringType,
		},
	}
}

type CreateVectorIndexRequest struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec types.List `tfsdk:"delta_sync_index_spec" tf:"optional,object"`
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec types.List `tfsdk:"direct_access_index_spec" tf:"optional,object"`
	// Name of the endpoint to be used for serving the index
	EndpointName types.String `tfsdk:"endpoint_name" tf:""`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType types.String `tfsdk:"index_type" tf:""`
	// Name of the index
	Name types.String `tfsdk:"name" tf:""`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key" tf:""`
}

func (newState *CreateVectorIndexRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVectorIndexRequest) {
}

func (newState *CreateVectorIndexRequest) SyncEffectiveFieldsDuringRead(existingState CreateVectorIndexRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVectorIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_sync_index_spec":    reflect.TypeOf(DeltaSyncVectorIndexSpecRequest{}),
		"direct_access_index_spec": reflect.TypeOf(DirectAccessVectorIndexSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVectorIndexRequest
// only implements ToObjectValue() and Type().
func (o CreateVectorIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_sync_index_spec":    o.DeltaSyncIndexSpec,
			"direct_access_index_spec": o.DirectAccessIndexSpec,
			"endpoint_name":            o.EndpointName,
			"index_type":               o.IndexType,
			"name":                     o.Name,
			"primary_key":              o.PrimaryKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateVectorIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_sync_index_spec": basetypes.ListType{
				ElemType: DeltaSyncVectorIndexSpecRequest{}.Type(ctx),
			},
			"direct_access_index_spec": basetypes.ListType{
				ElemType: DirectAccessVectorIndexSpec{}.Type(ctx),
			},
			"endpoint_name": types.StringType,
			"index_type":    types.StringType,
			"name":          types.StringType,
			"primary_key":   types.StringType,
		},
	}
}

type CreateVectorIndexResponse struct {
	VectorIndex types.List `tfsdk:"vector_index" tf:"optional,object"`
}

func (newState *CreateVectorIndexResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVectorIndexResponse) {
}

func (newState *CreateVectorIndexResponse) SyncEffectiveFieldsDuringRead(existingState CreateVectorIndexResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVectorIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"vector_index": reflect.TypeOf(VectorIndex{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVectorIndexResponse
// only implements ToObjectValue() and Type().
func (o CreateVectorIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vector_index": o.VectorIndex,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateVectorIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vector_index": basetypes.ListType{
				ElemType: VectorIndex{}.Type(ctx),
			},
		},
	}
}

// Result of the upsert or delete operation.
type DeleteDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys types.List `tfsdk:"failed_primary_keys" tf:"optional"`
	// Count of successfully processed rows.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count" tf:"optional"`
}

func (newState *DeleteDataResult) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDataResult) {
}

func (newState *DeleteDataResult) SyncEffectiveFieldsDuringRead(existingState DeleteDataResult) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDataResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDataResult) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failed_primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataResult
// only implements ToObjectValue() and Type().
func (o DeleteDataResult) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failed_primary_keys": o.FailedPrimaryKeys,
			"success_row_count":   o.SuccessRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDataResult) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failed_primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"success_row_count": types.Int64Type,
		},
	}
}

// Request payload for deleting data from a vector index.
type DeleteDataVectorIndexRequest struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	IndexName types.String `tfsdk:"-"`
	// List of primary keys for the data to be deleted.
	PrimaryKeys types.List `tfsdk:"primary_keys" tf:""`
}

func (newState *DeleteDataVectorIndexRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDataVectorIndexRequest) {
}

func (newState *DeleteDataVectorIndexRequest) SyncEffectiveFieldsDuringRead(existingState DeleteDataVectorIndexRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDataVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDataVectorIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataVectorIndexRequest
// only implements ToObjectValue() and Type().
func (o DeleteDataVectorIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":   o.IndexName,
			"primary_keys": o.PrimaryKeys,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDataVectorIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
			"primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// Response to a delete data vector index request.
type DeleteDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result types.List `tfsdk:"result" tf:"optional,object"`
	// Status of the delete operation.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *DeleteDataVectorIndexResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDataVectorIndexResponse) {
}

func (newState *DeleteDataVectorIndexResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDataVectorIndexResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDataVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDataVectorIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"result": reflect.TypeOf(DeleteDataResult{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataVectorIndexResponse
// only implements ToObjectValue() and Type().
func (o DeleteDataVectorIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result": o.Result,
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDataVectorIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"result": basetypes.ListType{
				ElemType: DeleteDataResult{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

// Delete an endpoint
type DeleteEndpointRequest struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (newState *DeleteEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteEndpointRequest) {
}

func (newState *DeleteEndpointRequest) SyncEffectiveFieldsDuringRead(existingState DeleteEndpointRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointRequest
// only implements ToObjectValue() and Type().
func (o DeleteEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": o.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
		},
	}
}

type DeleteEndpointResponse struct {
}

func (newState *DeleteEndpointResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteEndpointResponse) {
}

func (newState *DeleteEndpointResponse) SyncEffectiveFieldsDuringRead(existingState DeleteEndpointResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEndpointResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteEndpointResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointResponse
// only implements ToObjectValue() and Type().
func (o DeleteEndpointResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteEndpointResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete an index
type DeleteIndexRequest struct {
	// Name of the index
	IndexName types.String `tfsdk:"-"`
}

func (newState *DeleteIndexRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteIndexRequest) {
}

func (newState *DeleteIndexRequest) SyncEffectiveFieldsDuringRead(existingState DeleteIndexRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIndexRequest
// only implements ToObjectValue() and Type().
func (o DeleteIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": o.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
		},
	}
}

type DeleteIndexResponse struct {
}

func (newState *DeleteIndexResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteIndexResponse) {
}

func (newState *DeleteIndexResponse) SyncEffectiveFieldsDuringRead(existingState DeleteIndexResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIndexResponse
// only implements ToObjectValue() and Type().
func (o DeleteIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeltaSyncVectorIndexSpecRequest struct {
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	ColumnsToSync types.List `tfsdk:"columns_to_sync" tf:"optional"`
	// The columns that contain the embedding source.
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns" tf:"optional"`
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns" tf:"optional"`
	// [Optional] Automatically sync the vector index contents and computed
	// embeddings to the specified Delta table. The only supported table name is
	// the index name with the suffix `_writeback_table`.
	EmbeddingWritebackTable types.String `tfsdk:"embedding_writeback_table" tf:"optional"`
	// Pipeline execution mode.
	//
	// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the
	// system stops processing after successfully refreshing the source table in
	// the pipeline once, ensuring the table is updated based on the data
	// available when the update started. - `CONTINUOUS`: If the pipeline uses
	// continuous execution, the pipeline processes new data as it arrives in
	// the source table to keep vector index fresh.
	PipelineType types.String `tfsdk:"pipeline_type" tf:"optional"`
	// The name of the source table.
	SourceTable types.String `tfsdk:"source_table" tf:"optional"`
}

func (newState *DeltaSyncVectorIndexSpecRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaSyncVectorIndexSpecRequest) {
}

func (newState *DeltaSyncVectorIndexSpecRequest) SyncEffectiveFieldsDuringRead(existingState DeltaSyncVectorIndexSpecRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSyncVectorIndexSpecRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeltaSyncVectorIndexSpecRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns_to_sync":          reflect.TypeOf(types.String{}),
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSyncVectorIndexSpecRequest
// only implements ToObjectValue() and Type().
func (o DeltaSyncVectorIndexSpecRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns_to_sync":           o.ColumnsToSync,
			"embedding_source_columns":  o.EmbeddingSourceColumns,
			"embedding_vector_columns":  o.EmbeddingVectorColumns,
			"embedding_writeback_table": o.EmbeddingWritebackTable,
			"pipeline_type":             o.PipelineType,
			"source_table":              o.SourceTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaSyncVectorIndexSpecRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns_to_sync": basetypes.ListType{
				ElemType: types.StringType,
			},
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn{}.Type(ctx),
			},
			"embedding_writeback_table": types.StringType,
			"pipeline_type":             types.StringType,
			"source_table":              types.StringType,
		},
	}
}

type DeltaSyncVectorIndexSpecResponse struct {
	// The columns that contain the embedding source.
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns" tf:"optional"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns" tf:"optional"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable types.String `tfsdk:"embedding_writeback_table" tf:"optional"`
	// The ID of the pipeline that is used to sync the index.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// Pipeline execution mode.
	//
	// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the
	// system stops processing after successfully refreshing the source table in
	// the pipeline once, ensuring the table is updated based on the data
	// available when the update started. - `CONTINUOUS`: If the pipeline uses
	// continuous execution, the pipeline processes new data as it arrives in
	// the source table to keep vector index fresh.
	PipelineType types.String `tfsdk:"pipeline_type" tf:"optional"`
	// The name of the source table.
	SourceTable types.String `tfsdk:"source_table" tf:"optional"`
}

func (newState *DeltaSyncVectorIndexSpecResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaSyncVectorIndexSpecResponse) {
}

func (newState *DeltaSyncVectorIndexSpecResponse) SyncEffectiveFieldsDuringRead(existingState DeltaSyncVectorIndexSpecResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSyncVectorIndexSpecResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeltaSyncVectorIndexSpecResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSyncVectorIndexSpecResponse
// only implements ToObjectValue() and Type().
func (o DeltaSyncVectorIndexSpecResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_source_columns":  o.EmbeddingSourceColumns,
			"embedding_vector_columns":  o.EmbeddingVectorColumns,
			"embedding_writeback_table": o.EmbeddingWritebackTable,
			"pipeline_id":               o.PipelineId,
			"pipeline_type":             o.PipelineType,
			"source_table":              o.SourceTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaSyncVectorIndexSpecResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn{}.Type(ctx),
			},
			"embedding_writeback_table": types.StringType,
			"pipeline_id":               types.StringType,
			"pipeline_type":             types.StringType,
			"source_table":              types.StringType,
		},
	}
}

type DirectAccessVectorIndexSpec struct {
	// Contains the optional model endpoint to use during query time.
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns" tf:"optional"`

	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns" tf:"optional"`
	// The schema of the index in JSON format.
	//
	// Supported types are `integer`, `long`, `float`, `double`, `boolean`,
	// `string`, `date`, `timestamp`.
	//
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	SchemaJson types.String `tfsdk:"schema_json" tf:"optional"`
}

func (newState *DirectAccessVectorIndexSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan DirectAccessVectorIndexSpec) {
}

func (newState *DirectAccessVectorIndexSpec) SyncEffectiveFieldsDuringRead(existingState DirectAccessVectorIndexSpec) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DirectAccessVectorIndexSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DirectAccessVectorIndexSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectAccessVectorIndexSpec
// only implements ToObjectValue() and Type().
func (o DirectAccessVectorIndexSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_source_columns": o.EmbeddingSourceColumns,
			"embedding_vector_columns": o.EmbeddingVectorColumns,
			"schema_json":              o.SchemaJson,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DirectAccessVectorIndexSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn{}.Type(ctx),
			},
			"schema_json": types.StringType,
		},
	}
}

type EmbeddingSourceColumn struct {
	// Name of the embedding model endpoint
	EmbeddingModelEndpointName types.String `tfsdk:"embedding_model_endpoint_name" tf:"optional"`
	// Name of the column
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *EmbeddingSourceColumn) SyncEffectiveFieldsDuringCreateOrUpdate(plan EmbeddingSourceColumn) {
}

func (newState *EmbeddingSourceColumn) SyncEffectiveFieldsDuringRead(existingState EmbeddingSourceColumn) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmbeddingSourceColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EmbeddingSourceColumn) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingSourceColumn
// only implements ToObjectValue() and Type().
func (o EmbeddingSourceColumn) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_model_endpoint_name": o.EmbeddingModelEndpointName,
			"name":                          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EmbeddingSourceColumn) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_model_endpoint_name": types.StringType,
			"name":                          types.StringType,
		},
	}
}

type EmbeddingVectorColumn struct {
	// Dimension of the embedding vector
	EmbeddingDimension types.Int64 `tfsdk:"embedding_dimension" tf:"optional"`
	// Name of the column
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *EmbeddingVectorColumn) SyncEffectiveFieldsDuringCreateOrUpdate(plan EmbeddingVectorColumn) {
}

func (newState *EmbeddingVectorColumn) SyncEffectiveFieldsDuringRead(existingState EmbeddingVectorColumn) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmbeddingVectorColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EmbeddingVectorColumn) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingVectorColumn
// only implements ToObjectValue() and Type().
func (o EmbeddingVectorColumn) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_dimension": o.EmbeddingDimension,
			"name":                o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EmbeddingVectorColumn) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_dimension": types.Int64Type,
			"name":                types.StringType,
		},
	}
}

type EndpointInfo struct {
	// Timestamp of endpoint creation
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Creator of the endpoint
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// Current status of the endpoint
	EndpointStatus types.List `tfsdk:"endpoint_status" tf:"optional,object"`
	// Type of endpoint.
	EndpointType types.String `tfsdk:"endpoint_type" tf:"optional"`
	// Unique identifier of the endpoint
	Id types.String `tfsdk:"id" tf:"optional"`
	// Timestamp of last update to the endpoint
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// User who last updated the endpoint
	LastUpdatedUser types.String `tfsdk:"last_updated_user" tf:"optional"`
	// Name of endpoint
	Name types.String `tfsdk:"name" tf:"optional"`
	// Number of indexes on the endpoint
	NumIndexes types.Int64 `tfsdk:"num_indexes" tf:"optional"`
}

func (newState *EndpointInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointInfo) {
}

func (newState *EndpointInfo) SyncEffectiveFieldsDuringRead(existingState EndpointInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint_status": reflect.TypeOf(EndpointStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointInfo
// only implements ToObjectValue() and Type().
func (o EndpointInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":     o.CreationTimestamp,
			"creator":                o.Creator,
			"endpoint_status":        o.EndpointStatus,
			"endpoint_type":          o.EndpointType,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"last_updated_user":      o.LastUpdatedUser,
			"name":                   o.Name,
			"num_indexes":            o.NumIndexes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
			"endpoint_status": basetypes.ListType{
				ElemType: EndpointStatus{}.Type(ctx),
			},
			"endpoint_type":          types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"last_updated_user":      types.StringType,
			"name":                   types.StringType,
			"num_indexes":            types.Int64Type,
		},
	}
}

// Status information of an endpoint
type EndpointStatus struct {
	// Additional status message
	Message types.String `tfsdk:"message" tf:"optional"`
	// Current state of the endpoint
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *EndpointStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointStatus) {
}

func (newState *EndpointStatus) SyncEffectiveFieldsDuringRead(existingState EndpointStatus) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointStatus
// only implements ToObjectValue() and Type().
func (o EndpointStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
			"state":   o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

// Get an endpoint
type GetEndpointRequest struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (newState *GetEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEndpointRequest) {
}

func (newState *GetEndpointRequest) SyncEffectiveFieldsDuringRead(existingState GetEndpointRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEndpointRequest
// only implements ToObjectValue() and Type().
func (o GetEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": o.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
		},
	}
}

// Get an index
type GetIndexRequest struct {
	// Name of the index
	IndexName types.String `tfsdk:"-"`
}

func (newState *GetIndexRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetIndexRequest) {
}

func (newState *GetIndexRequest) SyncEffectiveFieldsDuringRead(existingState GetIndexRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIndexRequest
// only implements ToObjectValue() and Type().
func (o GetIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": o.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
		},
	}
}

type ListEndpointResponse struct {
	// An array of Endpoint objects
	Endpoints types.List `tfsdk:"endpoints" tf:"optional"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListEndpointResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListEndpointResponse) {
}

func (newState *ListEndpointResponse) SyncEffectiveFieldsDuringRead(existingState ListEndpointResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListEndpointResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(EndpointInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointResponse
// only implements ToObjectValue() and Type().
func (o ListEndpointResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints":       o.Endpoints,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListEndpointResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: EndpointInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// List all endpoints
type ListEndpointsRequest struct {
	// Token for pagination
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListEndpointsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListEndpointsRequest) {
}

func (newState *ListEndpointsRequest) SyncEffectiveFieldsDuringRead(existingState ListEndpointsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListEndpointsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsRequest
// only implements ToObjectValue() and Type().
func (o ListEndpointsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListEndpointsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

// List indexes
type ListIndexesRequest struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
	// Token for pagination
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListIndexesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListIndexesRequest) {
}

func (newState *ListIndexesRequest) SyncEffectiveFieldsDuringRead(existingState ListIndexesRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIndexesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListIndexesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIndexesRequest
// only implements ToObjectValue() and Type().
func (o ListIndexesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": o.EndpointName,
			"page_token":    o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListIndexesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
			"page_token":    types.StringType,
		},
	}
}

type ListValue struct {
	Values types.List `tfsdk:"values" tf:"optional"`
}

func (newState *ListValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListValue) {
}

func (newState *ListValue) SyncEffectiveFieldsDuringRead(existingState ListValue) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(Value{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListValue
// only implements ToObjectValue() and Type().
func (o ListValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"values": o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"values": basetypes.ListType{
				ElemType: Value{}.Type(ctx),
			},
		},
	}
}

type ListVectorIndexesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	VectorIndexes types.List `tfsdk:"vector_indexes" tf:"optional"`
}

func (newState *ListVectorIndexesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVectorIndexesResponse) {
}

func (newState *ListVectorIndexesResponse) SyncEffectiveFieldsDuringRead(existingState ListVectorIndexesResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVectorIndexesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVectorIndexesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"vector_indexes": reflect.TypeOf(MiniVectorIndex{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVectorIndexesResponse
// only implements ToObjectValue() and Type().
func (o ListVectorIndexesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"vector_indexes":  o.VectorIndexes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListVectorIndexesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"vector_indexes": basetypes.ListType{
				ElemType: MiniVectorIndex{}.Type(ctx),
			},
		},
	}
}

// Key-value pair.
type MapStringValueEntry struct {
	// Column name.
	Key types.String `tfsdk:"key" tf:"optional"`
	// Column value, nullable.
	Value types.List `tfsdk:"value" tf:"optional,object"`
}

func (newState *MapStringValueEntry) SyncEffectiveFieldsDuringCreateOrUpdate(plan MapStringValueEntry) {
}

func (newState *MapStringValueEntry) SyncEffectiveFieldsDuringRead(existingState MapStringValueEntry) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MapStringValueEntry.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MapStringValueEntry) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(Value{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MapStringValueEntry
// only implements ToObjectValue() and Type().
func (o MapStringValueEntry) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MapStringValueEntry) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key": types.StringType,
			"value": basetypes.ListType{
				ElemType: Value{}.Type(ctx),
			},
		},
	}
}

type MiniVectorIndex struct {
	// The user who created the index.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// Name of the endpoint associated with the index
	EndpointName types.String `tfsdk:"endpoint_name" tf:"optional"`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType types.String `tfsdk:"index_type" tf:"optional"`
	// Name of the index
	Name types.String `tfsdk:"name" tf:"optional"`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key" tf:"optional"`
}

func (newState *MiniVectorIndex) SyncEffectiveFieldsDuringCreateOrUpdate(plan MiniVectorIndex) {
}

func (newState *MiniVectorIndex) SyncEffectiveFieldsDuringRead(existingState MiniVectorIndex) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MiniVectorIndex.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MiniVectorIndex) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MiniVectorIndex
// only implements ToObjectValue() and Type().
func (o MiniVectorIndex) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator":       o.Creator,
			"endpoint_name": o.EndpointName,
			"index_type":    o.IndexType,
			"name":          o.Name,
			"primary_key":   o.PrimaryKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MiniVectorIndex) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creator":       types.StringType,
			"endpoint_name": types.StringType,
			"index_type":    types.StringType,
			"name":          types.StringType,
			"primary_key":   types.StringType,
		},
	}
}

// Request payload for getting next page of results.
type QueryVectorIndexNextPageRequest struct {
	// Name of the endpoint.
	EndpointName types.String `tfsdk:"endpoint_name" tf:"optional"`
	// Name of the vector index to query.
	IndexName types.String `tfsdk:"-"`
	// Page token returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` API.
	PageToken types.String `tfsdk:"page_token" tf:"optional"`
}

func (newState *QueryVectorIndexNextPageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryVectorIndexNextPageRequest) {
}

func (newState *QueryVectorIndexNextPageRequest) SyncEffectiveFieldsDuringRead(existingState QueryVectorIndexNextPageRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryVectorIndexNextPageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryVectorIndexNextPageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexNextPageRequest
// only implements ToObjectValue() and Type().
func (o QueryVectorIndexNextPageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": o.EndpointName,
			"index_name":    o.IndexName,
			"page_token":    o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryVectorIndexNextPageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
			"index_name":    types.StringType,
			"page_token":    types.StringType,
		},
	}
}

type QueryVectorIndexRequest struct {
	// List of column names to include in the response.
	Columns types.List `tfsdk:"columns" tf:""`
	// JSON string representing query filters.
	//
	// Example filters: - `{"id <": 5}`: Filter for id less than 5. - `{"id >":
	// 5}`: Filter for id greater than 5. - `{"id <=": 5}`: Filter for id less
	// than equal to 5. - `{"id >=": 5}`: Filter for id greater than equal to 5.
	// - `{"id": 5}`: Filter for id equal to 5.
	FiltersJson types.String `tfsdk:"filters_json" tf:"optional"`
	// Name of the vector index to query.
	IndexName types.String `tfsdk:"-"`
	// Number of results to return. Defaults to 10.
	NumResults types.Int64 `tfsdk:"num_results" tf:"optional"`
	// Query text. Required for Delta Sync Index using model endpoint.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// The query type to use. Choices are `ANN` and `HYBRID`. Defaults to `ANN`.
	QueryType types.String `tfsdk:"query_type" tf:"optional"`
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	QueryVector types.List `tfsdk:"query_vector" tf:"optional"`
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold types.Float64 `tfsdk:"score_threshold" tf:"optional"`
}

func (newState *QueryVectorIndexRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryVectorIndexRequest) {
}

func (newState *QueryVectorIndexRequest) SyncEffectiveFieldsDuringRead(existingState QueryVectorIndexRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryVectorIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns":      reflect.TypeOf(types.String{}),
		"query_vector": reflect.TypeOf(types.Float64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexRequest
// only implements ToObjectValue() and Type().
func (o QueryVectorIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns":         o.Columns,
			"filters_json":    o.FiltersJson,
			"index_name":      o.IndexName,
			"num_results":     o.NumResults,
			"query_text":      o.QueryText,
			"query_type":      o.QueryType,
			"query_vector":    o.QueryVector,
			"score_threshold": o.ScoreThreshold,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryVectorIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"filters_json": types.StringType,
			"index_name":   types.StringType,
			"num_results":  types.Int64Type,
			"query_text":   types.StringType,
			"query_type":   types.StringType,
			"query_vector": basetypes.ListType{
				ElemType: types.Float64Type,
			},
			"score_threshold": types.Float64Type,
		},
	}
}

type QueryVectorIndexResponse struct {
	// Metadata about the result set.
	Manifest types.List `tfsdk:"manifest" tf:"optional,object"`
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Data returned in the query result.
	Result types.List `tfsdk:"result" tf:"optional,object"`
}

func (newState *QueryVectorIndexResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryVectorIndexResponse) {
}

func (newState *QueryVectorIndexResponse) SyncEffectiveFieldsDuringRead(existingState QueryVectorIndexResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryVectorIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(ResultManifest{}),
		"result":   reflect.TypeOf(ResultData{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexResponse
// only implements ToObjectValue() and Type().
func (o QueryVectorIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"manifest":        o.Manifest,
			"next_page_token": o.NextPageToken,
			"result":          o.Result,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryVectorIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"manifest": basetypes.ListType{
				ElemType: ResultManifest{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"result": basetypes.ListType{
				ElemType: ResultData{}.Type(ctx),
			},
		},
	}
}

// Data returned in the query result.
type ResultData struct {
	// Data rows returned in the query.
	DataArray types.List `tfsdk:"data_array" tf:"optional"`
	// Number of rows in the result set.
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
}

func (newState *ResultData) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultData) {
}

func (newState *ResultData) SyncEffectiveFieldsDuringRead(existingState ResultData) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultData.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResultData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_array": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultData
// only implements ToObjectValue() and Type().
func (o ResultData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_array": o.DataArray,
			"row_count":  o.RowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResultData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_array": basetypes.ListType{
				ElemType: basetypes.ListType{
					ElemType: types.StringType,
				},
			},
			"row_count": types.Int64Type,
		},
	}
}

// Metadata about the result set.
type ResultManifest struct {
	// Number of columns in the result set.
	ColumnCount types.Int64 `tfsdk:"column_count" tf:"optional"`
	// Information about each column in the result set.
	Columns types.List `tfsdk:"columns" tf:"optional"`
}

func (newState *ResultManifest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultManifest) {
}

func (newState *ResultManifest) SyncEffectiveFieldsDuringRead(existingState ResultManifest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultManifest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResultManifest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest
// only implements ToObjectValue() and Type().
func (o ResultManifest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count": o.ColumnCount,
			"columns":      o.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResultManifest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column_count": types.Int64Type,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo{}.Type(ctx),
			},
		},
	}
}

// Request payload for scanning data from a vector index.
type ScanVectorIndexRequest struct {
	// Name of the vector index to scan.
	IndexName types.String `tfsdk:"-"`
	// Primary key of the last entry returned in the previous scan.
	LastPrimaryKey types.String `tfsdk:"last_primary_key" tf:"optional"`
	// Number of results to return. Defaults to 10.
	NumResults types.Int64 `tfsdk:"num_results" tf:"optional"`
}

func (newState *ScanVectorIndexRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ScanVectorIndexRequest) {
}

func (newState *ScanVectorIndexRequest) SyncEffectiveFieldsDuringRead(existingState ScanVectorIndexRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ScanVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ScanVectorIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ScanVectorIndexRequest
// only implements ToObjectValue() and Type().
func (o ScanVectorIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":       o.IndexName,
			"last_primary_key": o.LastPrimaryKey,
			"num_results":      o.NumResults,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ScanVectorIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name":       types.StringType,
			"last_primary_key": types.StringType,
			"num_results":      types.Int64Type,
		},
	}
}

// Response to a scan vector index request.
type ScanVectorIndexResponse struct {
	// List of data entries
	Data types.List `tfsdk:"data" tf:"optional"`
	// Primary key of the last entry.
	LastPrimaryKey types.String `tfsdk:"last_primary_key" tf:"optional"`
}

func (newState *ScanVectorIndexResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ScanVectorIndexResponse) {
}

func (newState *ScanVectorIndexResponse) SyncEffectiveFieldsDuringRead(existingState ScanVectorIndexResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ScanVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ScanVectorIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data": reflect.TypeOf(Struct{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ScanVectorIndexResponse
// only implements ToObjectValue() and Type().
func (o ScanVectorIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":             o.Data,
			"last_primary_key": o.LastPrimaryKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ScanVectorIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data": basetypes.ListType{
				ElemType: Struct{}.Type(ctx),
			},
			"last_primary_key": types.StringType,
		},
	}
}

type Struct struct {
	// Data entry, corresponding to a row in a vector index.
	Fields types.List `tfsdk:"fields" tf:"optional"`
}

func (newState *Struct) SyncEffectiveFieldsDuringCreateOrUpdate(plan Struct) {
}

func (newState *Struct) SyncEffectiveFieldsDuringRead(existingState Struct) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Struct.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Struct) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fields": reflect.TypeOf(MapStringValueEntry{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Struct
// only implements ToObjectValue() and Type().
func (o Struct) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fields": o.Fields,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Struct) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fields": basetypes.ListType{
				ElemType: MapStringValueEntry{}.Type(ctx),
			},
		},
	}
}

// Synchronize an index
type SyncIndexRequest struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName types.String `tfsdk:"-"`
}

func (newState *SyncIndexRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncIndexRequest) {
}

func (newState *SyncIndexRequest) SyncEffectiveFieldsDuringRead(existingState SyncIndexRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncIndexRequest
// only implements ToObjectValue() and Type().
func (o SyncIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": o.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
		},
	}
}

type SyncIndexResponse struct {
}

func (newState *SyncIndexResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncIndexResponse) {
}

func (newState *SyncIndexResponse) SyncEffectiveFieldsDuringRead(existingState SyncIndexResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncIndexResponse
// only implements ToObjectValue() and Type().
func (o SyncIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SyncIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Result of the upsert or delete operation.
type UpsertDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys types.List `tfsdk:"failed_primary_keys" tf:"optional"`
	// Count of successfully processed rows.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count" tf:"optional"`
}

func (newState *UpsertDataResult) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpsertDataResult) {
}

func (newState *UpsertDataResult) SyncEffectiveFieldsDuringRead(existingState UpsertDataResult) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertDataResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpsertDataResult) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failed_primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataResult
// only implements ToObjectValue() and Type().
func (o UpsertDataResult) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failed_primary_keys": o.FailedPrimaryKeys,
			"success_row_count":   o.SuccessRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpsertDataResult) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failed_primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"success_row_count": types.Int64Type,
		},
	}
}

// Request payload for upserting data into a vector index.
type UpsertDataVectorIndexRequest struct {
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	IndexName types.String `tfsdk:"-"`
	// JSON string representing the data to be upserted.
	InputsJson types.String `tfsdk:"inputs_json" tf:""`
}

func (newState *UpsertDataVectorIndexRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpsertDataVectorIndexRequest) {
}

func (newState *UpsertDataVectorIndexRequest) SyncEffectiveFieldsDuringRead(existingState UpsertDataVectorIndexRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertDataVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpsertDataVectorIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataVectorIndexRequest
// only implements ToObjectValue() and Type().
func (o UpsertDataVectorIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":  o.IndexName,
			"inputs_json": o.InputsJson,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpsertDataVectorIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name":  types.StringType,
			"inputs_json": types.StringType,
		},
	}
}

// Response to an upsert data vector index request.
type UpsertDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result types.List `tfsdk:"result" tf:"optional,object"`
	// Status of the upsert operation.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *UpsertDataVectorIndexResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpsertDataVectorIndexResponse) {
}

func (newState *UpsertDataVectorIndexResponse) SyncEffectiveFieldsDuringRead(existingState UpsertDataVectorIndexResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertDataVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpsertDataVectorIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"result": reflect.TypeOf(UpsertDataResult{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataVectorIndexResponse
// only implements ToObjectValue() and Type().
func (o UpsertDataVectorIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result": o.Result,
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpsertDataVectorIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"result": basetypes.ListType{
				ElemType: UpsertDataResult{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

type Value struct {
	BoolValue types.Bool `tfsdk:"bool_value" tf:"optional"`

	ListValue types.List `tfsdk:"list_value" tf:"optional,object"`

	NullValue types.String `tfsdk:"null_value" tf:"optional"`

	NumberValue types.Float64 `tfsdk:"number_value" tf:"optional"`

	StringValue types.String `tfsdk:"string_value" tf:"optional"`

	StructValue types.List `tfsdk:"struct_value" tf:"optional,object"`
}

func (newState *Value) SyncEffectiveFieldsDuringCreateOrUpdate(plan Value) {
}

func (newState *Value) SyncEffectiveFieldsDuringRead(existingState Value) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Value.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Value) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"list_value":   reflect.TypeOf(ListValue{}),
		"struct_value": reflect.TypeOf(Struct{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Value
// only implements ToObjectValue() and Type().
func (o Value) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bool_value":   o.BoolValue,
			"list_value":   o.ListValue,
			"null_value":   o.NullValue,
			"number_value": o.NumberValue,
			"string_value": o.StringValue,
			"struct_value": o.StructValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Value) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bool_value": types.BoolType,
			"list_value": basetypes.ListType{
				ElemType: ListValue{}.Type(ctx),
			},
			"null_value":   types.StringType,
			"number_value": types.Float64Type,
			"string_value": types.StringType,
			"struct_value": basetypes.ListType{
				ElemType: Struct{}.Type(ctx),
			},
		},
	}
}

type VectorIndex struct {
	// The user who created the index.
	Creator types.String `tfsdk:"creator" tf:"optional"`

	DeltaSyncIndexSpec types.List `tfsdk:"delta_sync_index_spec" tf:"optional,object"`

	DirectAccessIndexSpec types.List `tfsdk:"direct_access_index_spec" tf:"optional,object"`
	// Name of the endpoint associated with the index
	EndpointName types.String `tfsdk:"endpoint_name" tf:"optional"`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType types.String `tfsdk:"index_type" tf:"optional"`
	// Name of the index
	Name types.String `tfsdk:"name" tf:"optional"`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key" tf:"optional"`

	Status types.List `tfsdk:"status" tf:"optional,object"`
}

func (newState *VectorIndex) SyncEffectiveFieldsDuringCreateOrUpdate(plan VectorIndex) {
}

func (newState *VectorIndex) SyncEffectiveFieldsDuringRead(existingState VectorIndex) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorIndex.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a VectorIndex) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_sync_index_spec":    reflect.TypeOf(DeltaSyncVectorIndexSpecResponse{}),
		"direct_access_index_spec": reflect.TypeOf(DirectAccessVectorIndexSpec{}),
		"status":                   reflect.TypeOf(VectorIndexStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorIndex
// only implements ToObjectValue() and Type().
func (o VectorIndex) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator":                  o.Creator,
			"delta_sync_index_spec":    o.DeltaSyncIndexSpec,
			"direct_access_index_spec": o.DirectAccessIndexSpec,
			"endpoint_name":            o.EndpointName,
			"index_type":               o.IndexType,
			"name":                     o.Name,
			"primary_key":              o.PrimaryKey,
			"status":                   o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o VectorIndex) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creator": types.StringType,
			"delta_sync_index_spec": basetypes.ListType{
				ElemType: DeltaSyncVectorIndexSpecResponse{}.Type(ctx),
			},
			"direct_access_index_spec": basetypes.ListType{
				ElemType: DirectAccessVectorIndexSpec{}.Type(ctx),
			},
			"endpoint_name": types.StringType,
			"index_type":    types.StringType,
			"name":          types.StringType,
			"primary_key":   types.StringType,
			"status": basetypes.ListType{
				ElemType: VectorIndexStatus{}.Type(ctx),
			},
		},
	}
}

type VectorIndexStatus struct {
	// Index API Url to be used to perform operations on the index
	IndexUrl types.String `tfsdk:"index_url" tf:"optional"`
	// Number of rows indexed
	IndexedRowCount types.Int64 `tfsdk:"indexed_row_count" tf:"optional"`
	// Message associated with the index status
	Message types.String `tfsdk:"message" tf:"optional"`
	// Whether the index is ready for search
	Ready types.Bool `tfsdk:"ready" tf:"optional"`
}

func (newState *VectorIndexStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan VectorIndexStatus) {
}

func (newState *VectorIndexStatus) SyncEffectiveFieldsDuringRead(existingState VectorIndexStatus) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorIndexStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a VectorIndexStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorIndexStatus
// only implements ToObjectValue() and Type().
func (o VectorIndexStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_url":         o.IndexUrl,
			"indexed_row_count": o.IndexedRowCount,
			"message":           o.Message,
			"ready":             o.Ready,
		})
}

// Type implements basetypes.ObjectValuable.
func (o VectorIndexStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_url":         types.StringType,
			"indexed_row_count": types.Int64Type,
			"message":           types.StringType,
			"ready":             types.BoolType,
		},
	}
}

// Status of the delete operation.

// Current state of the endpoint

// Type of endpoint.

// Pipeline execution mode.
//
// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the system
// stops processing after successfully refreshing the source table in the
// pipeline once, ensuring the table is updated based on the data available when
// the update started. - `CONTINUOUS`: If the pipeline uses continuous
// execution, the pipeline processes new data as it arrives in the source table
// to keep vector index fresh.

// Status of the upsert operation.

// There are 2 types of Vector Search indexes:
//
// - `DELTA_SYNC`: An index that automatically syncs with a source Delta Table,
// automatically and incrementally updating the index as the underlying data in
// the Delta Table changes. - `DIRECT_ACCESS`: An index that supports direct
// read and write of vectors and metadata through our REST and SDK APIs. With
// this model, the user manages index updates.
