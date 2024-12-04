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

func (a ColumnInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ColumnInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Name": types.StringType,
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

func (a CreateEndpoint) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateEndpoint) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EndpointType": types.StringType,
			"Name":         types.StringType,
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

func (a CreateVectorIndexRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DeltaSyncIndexSpec":    reflect.TypeOf(DeltaSyncVectorIndexSpecRequest{}),
		"DirectAccessIndexSpec": reflect.TypeOf(DirectAccessVectorIndexSpec{}),
	}
}

func (a CreateVectorIndexRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DeltaSyncIndexSpec":    DeltaSyncVectorIndexSpecRequest{}.ToAttrType(ctx),
			"DirectAccessIndexSpec": DirectAccessVectorIndexSpec{}.ToAttrType(ctx),
			"EndpointName":          types.StringType,
			"IndexType":             types.StringType,
			"Name":                  types.StringType,
			"PrimaryKey":            types.StringType,
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

func (a CreateVectorIndexResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"VectorIndex": reflect.TypeOf(VectorIndex{}),
	}
}

func (a CreateVectorIndexResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"VectorIndex": VectorIndex{}.ToAttrType(ctx),
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

func (a DeleteDataResult) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"FailedPrimaryKeys": reflect.TypeOf(types.StringType),
	}
}

func (a DeleteDataResult) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FailedPrimaryKeys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"SuccessRowCount": types.Int64Type,
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

func (a DeleteDataVectorIndexRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"PrimaryKeys": reflect.TypeOf(types.StringType),
	}
}

func (a DeleteDataVectorIndexRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IndexName": types.StringType,
			"PrimaryKeys": basetypes.ListType{
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

func (a DeleteDataVectorIndexResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Result": reflect.TypeOf(DeleteDataResult{}),
	}
}

func (a DeleteDataVectorIndexResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Result": DeleteDataResult{}.ToAttrType(ctx),
			"Status": types.StringType,
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

func (a DeleteEndpointRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteEndpointRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EndpointName": types.StringType,
		},
	}
}

type DeleteEndpointResponse struct {
}

func (newState *DeleteEndpointResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteEndpointResponse) {
}

func (newState *DeleteEndpointResponse) SyncEffectiveFieldsDuringRead(existingState DeleteEndpointResponse) {
}

func (a DeleteEndpointResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteEndpointResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeleteIndexRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteIndexRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IndexName": types.StringType,
		},
	}
}

type DeleteIndexResponse struct {
}

func (newState *DeleteIndexResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteIndexResponse) {
}

func (newState *DeleteIndexResponse) SyncEffectiveFieldsDuringRead(existingState DeleteIndexResponse) {
}

func (a DeleteIndexResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteIndexResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeltaSyncVectorIndexSpecRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ColumnsToSync":          reflect.TypeOf(types.StringType),
		"EmbeddingSourceColumns": reflect.TypeOf(EmbeddingSourceColumn{}),
		"EmbeddingVectorColumns": reflect.TypeOf(EmbeddingVectorColumn{}),
	}
}

func (a DeltaSyncVectorIndexSpecRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ColumnsToSync": basetypes.ListType{
				ElemType: types.StringType,
			},
			"EmbeddingSourceColumns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn{}.ToAttrType(ctx),
			},
			"EmbeddingVectorColumns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn{}.ToAttrType(ctx),
			},
			"EmbeddingWritebackTable": types.StringType,
			"PipelineType":            types.StringType,
			"SourceTable":             types.StringType,
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

func (a DeltaSyncVectorIndexSpecResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"EmbeddingSourceColumns": reflect.TypeOf(EmbeddingSourceColumn{}),
		"EmbeddingVectorColumns": reflect.TypeOf(EmbeddingVectorColumn{}),
	}
}

func (a DeltaSyncVectorIndexSpecResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EmbeddingSourceColumns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn{}.ToAttrType(ctx),
			},
			"EmbeddingVectorColumns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn{}.ToAttrType(ctx),
			},
			"EmbeddingWritebackTable": types.StringType,
			"PipelineId":              types.StringType,
			"PipelineType":            types.StringType,
			"SourceTable":             types.StringType,
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

func (a DirectAccessVectorIndexSpec) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"EmbeddingSourceColumns": reflect.TypeOf(EmbeddingSourceColumn{}),
		"EmbeddingVectorColumns": reflect.TypeOf(EmbeddingVectorColumn{}),
	}
}

func (a DirectAccessVectorIndexSpec) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EmbeddingSourceColumns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn{}.ToAttrType(ctx),
			},
			"EmbeddingVectorColumns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn{}.ToAttrType(ctx),
			},
			"SchemaJson": types.StringType,
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

func (a EmbeddingSourceColumn) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a EmbeddingSourceColumn) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EmbeddingModelEndpointName": types.StringType,
			"Name":                       types.StringType,
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

func (a EmbeddingVectorColumn) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a EmbeddingVectorColumn) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EmbeddingDimension": types.Int64Type,
			"Name":               types.StringType,
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

func (a EndpointInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"EndpointStatus": reflect.TypeOf(EndpointStatus{}),
	}
}

func (a EndpointInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"CreationTimestamp":    types.Int64Type,
			"Creator":              types.StringType,
			"EndpointStatus":       EndpointStatus{}.ToAttrType(ctx),
			"EndpointType":         types.StringType,
			"Id":                   types.StringType,
			"LastUpdatedTimestamp": types.Int64Type,
			"LastUpdatedUser":      types.StringType,
			"Name":                 types.StringType,
			"NumIndexes":           types.Int64Type,
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

func (a EndpointStatus) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a EndpointStatus) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Message": types.StringType,
			"State":   types.StringType,
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

func (a GetEndpointRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetEndpointRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EndpointName": types.StringType,
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

func (a GetIndexRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetIndexRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IndexName": types.StringType,
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

func (a ListEndpointResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Endpoints": reflect.TypeOf(EndpointInfo{}),
	}
}

func (a ListEndpointResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Endpoints": basetypes.ListType{
				ElemType: EndpointInfo{}.ToAttrType(ctx),
			},
			"NextPageToken": types.StringType,
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

func (a ListEndpointsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListEndpointsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"PageToken": types.StringType,
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

func (a ListIndexesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListIndexesRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EndpointName": types.StringType,
			"PageToken":    types.StringType,
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

func (a ListValue) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Values": reflect.TypeOf(Value{}),
	}
}

func (a ListValue) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Values": basetypes.ListType{
				ElemType: Value{}.ToAttrType(ctx),
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

func (a ListVectorIndexesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"VectorIndexes": reflect.TypeOf(MiniVectorIndex{}),
	}
}

func (a ListVectorIndexesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"NextPageToken": types.StringType,
			"VectorIndexes": basetypes.ListType{
				ElemType: MiniVectorIndex{}.ToAttrType(ctx),
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

func (a MapStringValueEntry) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Value": reflect.TypeOf(Value{}),
	}
}

func (a MapStringValueEntry) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Key":   types.StringType,
			"Value": Value{}.ToAttrType(ctx),
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

func (a MiniVectorIndex) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a MiniVectorIndex) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Creator":      types.StringType,
			"EndpointName": types.StringType,
			"IndexType":    types.StringType,
			"Name":         types.StringType,
			"PrimaryKey":   types.StringType,
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

func (a QueryVectorIndexNextPageRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a QueryVectorIndexNextPageRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"EndpointName": types.StringType,
			"IndexName":    types.StringType,
			"PageToken":    types.StringType,
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

func (a QueryVectorIndexRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Columns":     reflect.TypeOf(types.StringType),
		"QueryVector": reflect.TypeOf(types.Float64Type),
	}
}

func (a QueryVectorIndexRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"FiltersJson": types.StringType,
			"IndexName":   types.StringType,
			"NumResults":  types.Int64Type,
			"QueryText":   types.StringType,
			"QueryType":   types.StringType,
			"QueryVector": basetypes.ListType{
				ElemType: types.Float64Type,
			},
			"ScoreThreshold": types.Float64Type,
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

func (a QueryVectorIndexResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Manifest": reflect.TypeOf(ResultManifest{}),
		"Result":   reflect.TypeOf(ResultData{}),
	}
}

func (a QueryVectorIndexResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Manifest":      ResultManifest{}.ToAttrType(ctx),
			"NextPageToken": types.StringType,
			"Result":        ResultData{}.ToAttrType(ctx),
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

func (a ResultData) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DataArray": reflect.TypeOf(types.StringType),
	}
}

func (a ResultData) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"DataArray": basetypes.ListType{
				ElemType: basetypes.ListType{
					ElemType: types.StringType,
				},
			},
			"RowCount": types.Int64Type,
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

func (a ResultManifest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Columns": reflect.TypeOf(ColumnInfo{}),
	}
}

func (a ResultManifest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ColumnCount": types.Int64Type,
			"Columns": basetypes.ListType{
				ElemType: ColumnInfo{}.ToAttrType(ctx),
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

func (a ScanVectorIndexRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ScanVectorIndexRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IndexName":      types.StringType,
			"LastPrimaryKey": types.StringType,
			"NumResults":     types.Int64Type,
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

func (a ScanVectorIndexResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Data": reflect.TypeOf(Struct{}),
	}
}

func (a ScanVectorIndexResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Data": basetypes.ListType{
				ElemType: Struct{}.ToAttrType(ctx),
			},
			"LastPrimaryKey": types.StringType,
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

func (a Struct) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Fields": reflect.TypeOf(MapStringValueEntry{}),
	}
}

func (a Struct) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Fields": basetypes.ListType{
				ElemType: MapStringValueEntry{}.ToAttrType(ctx),
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

func (a SyncIndexRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SyncIndexRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IndexName": types.StringType,
		},
	}
}

type SyncIndexResponse struct {
}

func (newState *SyncIndexResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncIndexResponse) {
}

func (newState *SyncIndexResponse) SyncEffectiveFieldsDuringRead(existingState SyncIndexResponse) {
}

func (a SyncIndexResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SyncIndexResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a UpsertDataResult) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"FailedPrimaryKeys": reflect.TypeOf(types.StringType),
	}
}

func (a UpsertDataResult) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"FailedPrimaryKeys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"SuccessRowCount": types.Int64Type,
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

func (a UpsertDataVectorIndexRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpsertDataVectorIndexRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IndexName":  types.StringType,
			"InputsJson": types.StringType,
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

func (a UpsertDataVectorIndexResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Result": reflect.TypeOf(UpsertDataResult{}),
	}
}

func (a UpsertDataVectorIndexResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Result": UpsertDataResult{}.ToAttrType(ctx),
			"Status": types.StringType,
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

func (a Value) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ListValue":   reflect.TypeOf(ListValue{}),
		"StructValue": reflect.TypeOf(Struct{}),
	}
}

func (a Value) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"BoolValue":   types.BoolType,
			"ListValue":   ListValue{}.ToAttrType(ctx),
			"NullValue":   types.StringType,
			"NumberValue": types.Float64Type,
			"StringValue": types.StringType,
			"StructValue": Struct{}.ToAttrType(ctx),
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

func (a VectorIndex) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DeltaSyncIndexSpec":    reflect.TypeOf(DeltaSyncVectorIndexSpecResponse{}),
		"DirectAccessIndexSpec": reflect.TypeOf(DirectAccessVectorIndexSpec{}),
		"Status":                reflect.TypeOf(VectorIndexStatus{}),
	}
}

func (a VectorIndex) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"Creator":               types.StringType,
			"DeltaSyncIndexSpec":    DeltaSyncVectorIndexSpecResponse{}.ToAttrType(ctx),
			"DirectAccessIndexSpec": DirectAccessVectorIndexSpec{}.ToAttrType(ctx),
			"EndpointName":          types.StringType,
			"IndexType":             types.StringType,
			"Name":                  types.StringType,
			"PrimaryKey":            types.StringType,
			"Status":                VectorIndexStatus{}.ToAttrType(ctx),
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

func (a VectorIndexStatus) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a VectorIndexStatus) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"IndexUrl":        types.StringType,
			"IndexedRowCount": types.Int64Type,
			"Message":         types.StringType,
			"Ready":           types.BoolType,
		},
	}
}
