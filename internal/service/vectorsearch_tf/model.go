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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ColumnInfo struct {
	// Name of the column.
	Name types.String `tfsdk:"name" tf:"optional"`
}

type CreateEndpoint struct {
	// Type of endpoint.
	EndpointType types.String `tfsdk:"endpoint_type" tf:""`
	// Name of endpoint
	Name types.String `tfsdk:"name" tf:""`
}

type CreateVectorIndexRequest struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecRequest `tfsdk:"delta_sync_index_spec" tf:"optional"`
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `tfsdk:"direct_access_index_spec" tf:"optional"`
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

type CreateVectorIndexResponse struct {
	VectorIndex *VectorIndex `tfsdk:"vector_index" tf:"optional"`
}

// Result of the upsert or delete operation.
type DeleteDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []types.String `tfsdk:"failed_primary_keys" tf:"optional"`
	// Count of successfully processed rows.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count" tf:"optional"`
}

// Request payload for deleting data from a vector index.
type DeleteDataVectorIndexRequest struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	IndexName types.String `tfsdk:"-"`
	// List of primary keys for the data to be deleted.
	PrimaryKeys []types.String `tfsdk:"primary_keys" tf:""`
}

// Response to a delete data vector index request.
type DeleteDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *DeleteDataResult `tfsdk:"result" tf:"optional"`
	// Status of the delete operation.
	Status types.String `tfsdk:"status" tf:"optional"`
}

// Delete an endpoint
type DeleteEndpointRequest struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
}

type DeleteEndpointResponse struct {
}

// Delete an index
type DeleteIndexRequest struct {
	// Name of the index
	IndexName types.String `tfsdk:"-"`
}

type DeleteIndexResponse struct {
}

type DeltaSyncVectorIndexSpecRequest struct {
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	ColumnsToSync []types.String `tfsdk:"columns_to_sync" tf:"optional"`
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []EmbeddingSourceColumn `tfsdk:"embedding_source_columns" tf:"optional"`
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	EmbeddingVectorColumns []EmbeddingVectorColumn `tfsdk:"embedding_vector_columns" tf:"optional"`
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

type DeltaSyncVectorIndexSpecResponse struct {
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []EmbeddingSourceColumn `tfsdk:"embedding_source_columns" tf:"optional"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns []EmbeddingVectorColumn `tfsdk:"embedding_vector_columns" tf:"optional"`
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

type DirectAccessVectorIndexSpec struct {
	// Contains the optional model endpoint to use during query time.
	EmbeddingSourceColumns []EmbeddingSourceColumn `tfsdk:"embedding_source_columns" tf:"optional"`

	EmbeddingVectorColumns []EmbeddingVectorColumn `tfsdk:"embedding_vector_columns" tf:"optional"`
	// The schema of the index in JSON format.
	//
	// Supported types are `integer`, `long`, `float`, `double`, `boolean`,
	// `string`, `date`, `timestamp`.
	//
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	SchemaJson types.String `tfsdk:"schema_json" tf:"optional"`
}

type EmbeddingSourceColumn struct {
	// Name of the embedding model endpoint
	EmbeddingModelEndpointName types.String `tfsdk:"embedding_model_endpoint_name" tf:"optional"`
	// Name of the column
	Name types.String `tfsdk:"name" tf:"optional"`
}

type EmbeddingVectorColumn struct {
	// Dimension of the embedding vector
	EmbeddingDimension types.Int64 `tfsdk:"embedding_dimension" tf:"optional"`
	// Name of the column
	Name types.String `tfsdk:"name" tf:"optional"`
}

type EndpointInfo struct {
	// Timestamp of endpoint creation
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Creator of the endpoint
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// Current status of the endpoint
	EndpointStatus *EndpointStatus `tfsdk:"endpoint_status" tf:"optional"`
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

// Status information of an endpoint
type EndpointStatus struct {
	// Additional status message
	Message types.String `tfsdk:"message" tf:"optional"`
	// Current state of the endpoint
	State types.String `tfsdk:"state" tf:"optional"`
}

// Get an endpoint
type GetEndpointRequest struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
}

// Get an index
type GetIndexRequest struct {
	// Name of the index
	IndexName types.String `tfsdk:"-"`
}

type ListEndpointResponse struct {
	// An array of Endpoint objects
	Endpoints []EndpointInfo `tfsdk:"endpoints" tf:"optional"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List all endpoints
type ListEndpointsRequest struct {
	// Token for pagination
	PageToken types.String `tfsdk:"-"`
}

// List indexes
type ListIndexesRequest struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
	// Token for pagination
	PageToken types.String `tfsdk:"-"`
}

type ListValue struct {
	Values []Value `tfsdk:"values" tf:"optional"`
}

type ListVectorIndexesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	VectorIndexes []MiniVectorIndex `tfsdk:"vector_indexes" tf:"optional"`
}

// Key-value pair.
type MapStringValueEntry struct {
	// Column name.
	Key types.String `tfsdk:"key" tf:"optional"`
	// Column value, nullable.
	Value *Value `tfsdk:"value" tf:"optional"`
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

type QueryVectorIndexRequest struct {
	// List of column names to include in the response.
	Columns []types.String `tfsdk:"columns" tf:""`
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
	QueryVector []types.Float64 `tfsdk:"query_vector" tf:"optional"`
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold types.Float64 `tfsdk:"score_threshold" tf:"optional"`
}

type QueryVectorIndexResponse struct {
	// Metadata about the result set.
	Manifest *ResultManifest `tfsdk:"manifest" tf:"optional"`
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Data returned in the query result.
	Result *ResultData `tfsdk:"result" tf:"optional"`
}

// Data returned in the query result.
type ResultData struct {
	// Data rows returned in the query.
	DataArray [][]types.String `tfsdk:"data_array" tf:"optional"`
	// Number of rows in the result set.
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
}

// Metadata about the result set.
type ResultManifest struct {
	// Number of columns in the result set.
	ColumnCount types.Int64 `tfsdk:"column_count" tf:"optional"`
	// Information about each column in the result set.
	Columns []ColumnInfo `tfsdk:"columns" tf:"optional"`
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

// Response to a scan vector index request.
type ScanVectorIndexResponse struct {
	// List of data entries
	Data []Struct `tfsdk:"data" tf:"optional"`
	// Primary key of the last entry.
	LastPrimaryKey types.String `tfsdk:"last_primary_key" tf:"optional"`
}

type Struct struct {
	// Data entry, corresponding to a row in a vector index.
	Fields []MapStringValueEntry `tfsdk:"fields" tf:"optional"`
}

// Synchronize an index
type SyncIndexRequest struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName types.String `tfsdk:"-"`
}

type SyncIndexResponse struct {
}

// Result of the upsert or delete operation.
type UpsertDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []types.String `tfsdk:"failed_primary_keys" tf:"optional"`
	// Count of successfully processed rows.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count" tf:"optional"`
}

// Request payload for upserting data into a vector index.
type UpsertDataVectorIndexRequest struct {
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	IndexName types.String `tfsdk:"-"`
	// JSON string representing the data to be upserted.
	InputsJson types.String `tfsdk:"inputs_json" tf:""`
}

// Response to an upsert data vector index request.
type UpsertDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *UpsertDataResult `tfsdk:"result" tf:"optional"`
	// Status of the upsert operation.
	Status types.String `tfsdk:"status" tf:"optional"`
}

type Value struct {
	BoolValue types.Bool `tfsdk:"bool_value" tf:"optional"`

	ListValue *ListValue `tfsdk:"list_value" tf:"optional"`

	NullValue types.String `tfsdk:"null_value" tf:"optional"`

	NumberValue types.Float64 `tfsdk:"number_value" tf:"optional"`

	StringValue types.String `tfsdk:"string_value" tf:"optional"`

	StructValue *Struct `tfsdk:"struct_value" tf:"optional"`
}

type VectorIndex struct {
	// The user who created the index.
	Creator types.String `tfsdk:"creator" tf:"optional"`

	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecResponse `tfsdk:"delta_sync_index_spec" tf:"optional"`

	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `tfsdk:"direct_access_index_spec" tf:"optional"`
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

	Status *VectorIndexStatus `tfsdk:"status" tf:"optional"`
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
