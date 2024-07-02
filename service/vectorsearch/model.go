// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type ColumnInfo struct {
	// Name of the column.
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateEndpoint struct {
	// Type of endpoint.
	EndpointType EndpointType `tfsdk:"endpoint_type"`
	// Name of endpoint
	Name string `tfsdk:"name"`
}

type CreateVectorIndexRequest struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecRequest `tfsdk:"delta_sync_index_spec"`
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `tfsdk:"direct_access_index_spec"`
	// Name of the endpoint to be used for serving the index
	EndpointName string `tfsdk:"endpoint_name"`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType VectorIndexType `tfsdk:"index_type"`
	// Name of the index
	Name string `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey string `tfsdk:"primary_key"`
}

type CreateVectorIndexResponse struct {
	VectorIndex *VectorIndex `tfsdk:"vector_index"`
}

// Result of the upsert or delete operation.
type DeleteDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []string `tfsdk:"failed_primary_keys"`
	// Count of successfully processed rows.
	SuccessRowCount int64 `tfsdk:"success_row_count"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeleteDataResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDataResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status of the delete operation.
type DeleteDataStatus string

const DeleteDataStatusFailure DeleteDataStatus = `FAILURE`

const DeleteDataStatusPartialSuccess DeleteDataStatus = `PARTIAL_SUCCESS`

const DeleteDataStatusSuccess DeleteDataStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *DeleteDataStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeleteDataStatus) Set(v string) error {
	switch v {
	case `FAILURE`, `PARTIAL_SUCCESS`, `SUCCESS`:
		*f = DeleteDataStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILURE", "PARTIAL_SUCCESS", "SUCCESS"`, v)
	}
}

// Type always returns DeleteDataStatus to satisfy [pflag.Value] interface
func (f *DeleteDataStatus) Type() string {
	return "DeleteDataStatus"
}

// Request payload for deleting data from a vector index.
type DeleteDataVectorIndexRequest struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	IndexName string `tfsdk:"-" url:"-"`
	// List of primary keys for the data to be deleted.
	PrimaryKeys []string `tfsdk:"primary_keys"`
}

// Response to a delete data vector index request.
type DeleteDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *DeleteDataResult `tfsdk:"result"`
	// Status of the delete operation.
	Status DeleteDataStatus `tfsdk:"status"`
}

// Delete an endpoint
type DeleteEndpointRequest struct {
	// Name of the endpoint
	EndpointName string `tfsdk:"-" url:"-"`
}

type DeleteEndpointResponse struct {
}

// Delete an index
type DeleteIndexRequest struct {
	// Name of the index
	IndexName string `tfsdk:"-" url:"-"`
}

type DeleteIndexResponse struct {
}

type DeltaSyncVectorIndexSpecRequest struct {
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []EmbeddingSourceColumn `tfsdk:"embedding_source_columns"`
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	EmbeddingVectorColumns []EmbeddingVectorColumn `tfsdk:"embedding_vector_columns"`
	// [Optional] Automatically sync the vector index contents and computed
	// embeddings to the specified Delta table. The only supported table name is
	// the index name with the suffix `_writeback_table`.
	EmbeddingWritebackTable string `tfsdk:"embedding_writeback_table"`
	// Pipeline execution mode.
	//
	// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the
	// system stops processing after successfully refreshing the source table in
	// the pipeline once, ensuring the table is updated based on the data
	// available when the update started. - `CONTINUOUS`: If the pipeline uses
	// continuous execution, the pipeline processes new data as it arrives in
	// the source table to keep vector index fresh.
	PipelineType PipelineType `tfsdk:"pipeline_type"`
	// The name of the source table.
	SourceTable string `tfsdk:"source_table"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeltaSyncVectorIndexSpecRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncVectorIndexSpecRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeltaSyncVectorIndexSpecResponse struct {
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []EmbeddingSourceColumn `tfsdk:"embedding_source_columns"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns []EmbeddingVectorColumn `tfsdk:"embedding_vector_columns"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable string `tfsdk:"embedding_writeback_table"`
	// The ID of the pipeline that is used to sync the index.
	PipelineId string `tfsdk:"pipeline_id"`
	// Pipeline execution mode.
	//
	// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the
	// system stops processing after successfully refreshing the source table in
	// the pipeline once, ensuring the table is updated based on the data
	// available when the update started. - `CONTINUOUS`: If the pipeline uses
	// continuous execution, the pipeline processes new data as it arrives in
	// the source table to keep vector index fresh.
	PipelineType PipelineType `tfsdk:"pipeline_type"`
	// The name of the source table.
	SourceTable string `tfsdk:"source_table"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DeltaSyncVectorIndexSpecResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncVectorIndexSpecResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DirectAccessVectorIndexSpec struct {
	// Contains the optional model endpoint to use during query time.
	EmbeddingSourceColumns []EmbeddingSourceColumn `tfsdk:"embedding_source_columns"`

	EmbeddingVectorColumns []EmbeddingVectorColumn `tfsdk:"embedding_vector_columns"`
	// The schema of the index in JSON format.
	//
	// Supported types are `integer`, `long`, `float`, `double`, `boolean`,
	// `string`, `date`, `timestamp`.
	//
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	SchemaJson string `tfsdk:"schema_json"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *DirectAccessVectorIndexSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DirectAccessVectorIndexSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EmbeddingSourceColumn struct {
	// Name of the embedding model endpoint
	EmbeddingModelEndpointName string `tfsdk:"embedding_model_endpoint_name"`
	// Name of the column
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EmbeddingSourceColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingSourceColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EmbeddingVectorColumn struct {
	// Dimension of the embedding vector
	EmbeddingDimension int `tfsdk:"embedding_dimension"`
	// Name of the column
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EmbeddingVectorColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingVectorColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointInfo struct {
	// Timestamp of endpoint creation
	CreationTimestamp int64 `tfsdk:"creation_timestamp"`
	// Creator of the endpoint
	Creator string `tfsdk:"creator"`
	// Current status of the endpoint
	EndpointStatus *EndpointStatus `tfsdk:"endpoint_status"`
	// Type of endpoint.
	EndpointType EndpointType `tfsdk:"endpoint_type"`
	// Unique identifier of the endpoint
	Id string `tfsdk:"id"`
	// Timestamp of last update to the endpoint
	LastUpdatedTimestamp int64 `tfsdk:"last_updated_timestamp"`
	// User who last updated the endpoint
	LastUpdatedUser string `tfsdk:"last_updated_user"`
	// Name of endpoint
	Name string `tfsdk:"name"`
	// Number of indexes on the endpoint
	NumIndexes int `tfsdk:"num_indexes"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status information of an endpoint
type EndpointStatus struct {
	// Additional status message
	Message string `tfsdk:"message"`
	// Current state of the endpoint
	State EndpointStatusState `tfsdk:"state"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EndpointStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Current state of the endpoint
type EndpointStatusState string

const EndpointStatusStateOffline EndpointStatusState = `OFFLINE`

const EndpointStatusStateOnline EndpointStatusState = `ONLINE`

const EndpointStatusStateProvisioning EndpointStatusState = `PROVISIONING`

// String representation for [fmt.Print]
func (f *EndpointStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStatusState) Set(v string) error {
	switch v {
	case `OFFLINE`, `ONLINE`, `PROVISIONING`:
		*f = EndpointStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OFFLINE", "ONLINE", "PROVISIONING"`, v)
	}
}

// Type always returns EndpointStatusState to satisfy [pflag.Value] interface
func (f *EndpointStatusState) Type() string {
	return "EndpointStatusState"
}

// Type of endpoint.
type EndpointType string

const EndpointTypeStandard EndpointType = `STANDARD`

// String representation for [fmt.Print]
func (f *EndpointType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointType) Set(v string) error {
	switch v {
	case `STANDARD`:
		*f = EndpointType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "STANDARD"`, v)
	}
}

// Type always returns EndpointType to satisfy [pflag.Value] interface
func (f *EndpointType) Type() string {
	return "EndpointType"
}

// Get an endpoint
type GetEndpointRequest struct {
	// Name of the endpoint
	EndpointName string `tfsdk:"-" url:"-"`
}

// Get an index
type GetIndexRequest struct {
	// Name of the index
	IndexName string `tfsdk:"-" url:"-"`
}

type ListEndpointResponse struct {
	// An array of Endpoint objects
	Endpoints []EndpointInfo `tfsdk:"endpoints"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListEndpointResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List all endpoints
type ListEndpointsRequest struct {
	// Token for pagination
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListEndpointsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List indexes
type ListIndexesRequest struct {
	// Name of the endpoint
	EndpointName string `tfsdk:"-" url:"endpoint_name"`
	// Token for pagination
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListIndexesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListIndexesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListValue struct {
	Values []Value `tfsdk:"values"`
}

type ListVectorIndexesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `tfsdk:"next_page_token"`

	VectorIndexes []MiniVectorIndex `tfsdk:"vector_indexes"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListVectorIndexesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVectorIndexesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Key-value pair.
type MapStringValueEntry struct {
	// Column name.
	Key string `tfsdk:"key"`
	// Column value, nullable.
	Value *Value `tfsdk:"value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *MapStringValueEntry) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MapStringValueEntry) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MiniVectorIndex struct {
	// The user who created the index.
	Creator string `tfsdk:"creator"`
	// Name of the endpoint associated with the index
	EndpointName string `tfsdk:"endpoint_name"`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType VectorIndexType `tfsdk:"index_type"`
	// Name of the index
	Name string `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey string `tfsdk:"primary_key"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *MiniVectorIndex) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MiniVectorIndex) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Pipeline execution mode.
//
// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the system
// stops processing after successfully refreshing the source table in the
// pipeline once, ensuring the table is updated based on the data available when
// the update started. - `CONTINUOUS`: If the pipeline uses continuous
// execution, the pipeline processes new data as it arrives in the source table
// to keep vector index fresh.
type PipelineType string

// If the pipeline uses continuous execution, the pipeline processes new data as
// it arrives in the source table to keep vector index fresh.
const PipelineTypeContinuous PipelineType = `CONTINUOUS`

// If the pipeline uses the triggered execution mode, the system stops
// processing after successfully refreshing the source table in the pipeline
// once, ensuring the table is updated based on the data available when the
// update started.
const PipelineTypeTriggered PipelineType = `TRIGGERED`

// String representation for [fmt.Print]
func (f *PipelineType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PipelineType) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `TRIGGERED`:
		*f = PipelineType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "TRIGGERED"`, v)
	}
}

// Type always returns PipelineType to satisfy [pflag.Value] interface
func (f *PipelineType) Type() string {
	return "PipelineType"
}

type QueryVectorIndexRequest struct {
	// List of column names to include in the response.
	Columns []string `tfsdk:"columns"`
	// JSON string representing query filters.
	//
	// Example filters: - `{"id <": 5}`: Filter for id less than 5. - `{"id >":
	// 5}`: Filter for id greater than 5. - `{"id <=": 5}`: Filter for id less
	// than equal to 5. - `{"id >=": 5}`: Filter for id greater than equal to 5.
	// - `{"id": 5}`: Filter for id equal to 5.
	FiltersJson string `tfsdk:"filters_json"`
	// Name of the vector index to query.
	IndexName string `tfsdk:"-" url:"-"`
	// Number of results to return. Defaults to 10.
	NumResults int `tfsdk:"num_results"`
	// Query text. Required for Delta Sync Index using model endpoint.
	QueryText string `tfsdk:"query_text"`
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	QueryVector []float64 `tfsdk:"query_vector"`
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold float64 `tfsdk:"score_threshold"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *QueryVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryVectorIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryVectorIndexResponse struct {
	// Metadata about the result set.
	Manifest *ResultManifest `tfsdk:"manifest"`
	// Data returned in the query result.
	Result *ResultData `tfsdk:"result"`
}

// Data returned in the query result.
type ResultData struct {
	// Data rows returned in the query.
	DataArray [][]string `tfsdk:"data_array"`
	// Number of rows in the result set.
	RowCount int `tfsdk:"row_count"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ResultData) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultData) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata about the result set.
type ResultManifest struct {
	// Number of columns in the result set.
	ColumnCount int `tfsdk:"column_count"`
	// Information about each column in the result set.
	Columns []ColumnInfo `tfsdk:"columns"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ResultManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request payload for scanning data from a vector index.
type ScanVectorIndexRequest struct {
	// Name of the vector index to scan.
	IndexName string `tfsdk:"-" url:"-"`
	// Primary key of the last entry returned in the previous scan.
	LastPrimaryKey string `tfsdk:"last_primary_key"`
	// Number of results to return. Defaults to 10.
	NumResults int `tfsdk:"num_results"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ScanVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ScanVectorIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response to a scan vector index request.
type ScanVectorIndexResponse struct {
	// List of data entries
	Data []Struct `tfsdk:"data"`
	// Primary key of the last entry.
	LastPrimaryKey string `tfsdk:"last_primary_key"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ScanVectorIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ScanVectorIndexResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Struct struct {
	// Data entry, corresponding to a row in a vector index.
	Fields []MapStringValueEntry `tfsdk:"fields"`
}

// Synchronize an index
type SyncIndexRequest struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName string `tfsdk:"-" url:"-"`
}

type SyncIndexResponse struct {
}

// Result of the upsert or delete operation.
type UpsertDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []string `tfsdk:"failed_primary_keys"`
	// Count of successfully processed rows.
	SuccessRowCount int64 `tfsdk:"success_row_count"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpsertDataResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpsertDataResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status of the upsert operation.
type UpsertDataStatus string

const UpsertDataStatusFailure UpsertDataStatus = `FAILURE`

const UpsertDataStatusPartialSuccess UpsertDataStatus = `PARTIAL_SUCCESS`

const UpsertDataStatusSuccess UpsertDataStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *UpsertDataStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpsertDataStatus) Set(v string) error {
	switch v {
	case `FAILURE`, `PARTIAL_SUCCESS`, `SUCCESS`:
		*f = UpsertDataStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILURE", "PARTIAL_SUCCESS", "SUCCESS"`, v)
	}
}

// Type always returns UpsertDataStatus to satisfy [pflag.Value] interface
func (f *UpsertDataStatus) Type() string {
	return "UpsertDataStatus"
}

// Request payload for upserting data into a vector index.
type UpsertDataVectorIndexRequest struct {
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	IndexName string `tfsdk:"-" url:"-"`
	// JSON string representing the data to be upserted.
	InputsJson string `tfsdk:"inputs_json"`
}

// Response to an upsert data vector index request.
type UpsertDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *UpsertDataResult `tfsdk:"result"`
	// Status of the upsert operation.
	Status UpsertDataStatus `tfsdk:"status"`
}

type Value struct {
	BoolValue bool `tfsdk:"bool_value"`

	ListValue *ListValue `tfsdk:"list_value"`

	NullValue string `tfsdk:"null_value"`

	NumberValue float64 `tfsdk:"number_value"`

	StringValue string `tfsdk:"string_value"`

	StructValue *Struct `tfsdk:"struct_value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Value) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Value) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VectorIndex struct {
	// The user who created the index.
	Creator string `tfsdk:"creator"`

	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecResponse `tfsdk:"delta_sync_index_spec"`

	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `tfsdk:"direct_access_index_spec"`
	// Name of the endpoint associated with the index
	EndpointName string `tfsdk:"endpoint_name"`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType VectorIndexType `tfsdk:"index_type"`
	// Name of the index
	Name string `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey string `tfsdk:"primary_key"`

	Status *VectorIndexStatus `tfsdk:"status"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *VectorIndex) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorIndex) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VectorIndexStatus struct {
	// Index API Url to be used to perform operations on the index
	IndexUrl string `tfsdk:"index_url"`
	// Number of rows indexed
	IndexedRowCount int64 `tfsdk:"indexed_row_count"`
	// Message associated with the index status
	Message string `tfsdk:"message"`
	// Whether the index is ready for search
	Ready bool `tfsdk:"ready"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *VectorIndexStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorIndexStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// There are 2 types of Vector Search indexes:
//
// - `DELTA_SYNC`: An index that automatically syncs with a source Delta Table,
// automatically and incrementally updating the index as the underlying data in
// the Delta Table changes. - `DIRECT_ACCESS`: An index that supports direct
// read and write of vectors and metadata through our REST and SDK APIs. With
// this model, the user manages index updates.
type VectorIndexType string

// An index that automatically syncs with a source Delta Table, automatically
// and incrementally updating the index as the underlying data in the Delta
// Table changes.
const VectorIndexTypeDeltaSync VectorIndexType = `DELTA_SYNC`

// An index that supports direct read and write of vectors and metadata through
// our REST and SDK APIs. With this model, the user manages index updates.
const VectorIndexTypeDirectAccess VectorIndexType = `DIRECT_ACCESS`

// String representation for [fmt.Print]
func (f *VectorIndexType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VectorIndexType) Set(v string) error {
	switch v {
	case `DELTA_SYNC`, `DIRECT_ACCESS`:
		*f = VectorIndexType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTA_SYNC", "DIRECT_ACCESS"`, v)
	}
}

// Type always returns VectorIndexType to satisfy [pflag.Value] interface
func (f *VectorIndexType) Type() string {
	return "VectorIndexType"
}
