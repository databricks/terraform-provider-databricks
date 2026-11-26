// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type ColumnInfo struct {
	// Name of the column.
	Name string `json:"name,omitempty"`
	// Data type of the column (e.g., "string", "int", "array<float>")
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateEndpoint struct {
	// Budget policy to attribute the endpoint's usage to. If omitted, the
	// default budget policy applies.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Type of endpoint to create. STANDARD endpoints are optimized for
	// low-latency queries and support Target QPS scaling. STORAGE_OPTIMIZED
	// endpoints hold much larger indexes and index faster, with higher query
	// latency. See
	// https://docs.databricks.com/aws/en/ai-search/ai-search#endpoint-options.
	EndpointType EndpointType `json:"endpoint_type"`
	// Name of the AI Search endpoint. The name must be unique within the
	// workspace, have at most 50 characters, contain only lowercase letters,
	// digits, hyphens, underscores and periods, and start and end with a letter
	// or digit.
	Name string `json:"name"`
	// Target queries per second (QPS) for the endpoint. Supported for STANDARD
	// endpoints only. The actual replica count is calculated at index
	// creation/sync time based on this value. Best-effort target; the system
	// does not guarantee this QPS will be achieved.
	TargetQps int64 `json:"target_qps,omitempty"`
	// The usage policy id to be applied once we've migrated to usage policies
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateVectorIndexRequest struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecRequest `json:"delta_sync_index_spec,omitempty"`
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `json:"direct_access_index_spec,omitempty"`
	// Name of the endpoint to be used for serving the index
	EndpointName string `json:"endpoint_name"`
	// The subtype of the index. Use `HYBRID` or `FULL_TEXT`. `VECTOR` is not
	// supported.
	IndexSubtype IndexSubtype `json:"index_subtype,omitempty"`
	// Type of index to create, which determines which spec field is required.
	IndexType VectorIndexType `json:"index_type"`
	// Three-level Unity Catalog name of the index to create.
	Name string `json:"name"`
	// Name of the column that uniquely identifies each row. For a Delta Sync
	// index it is a column of the source table; for a Direct Vector Access
	// index it must be declared in `schema_json`. It cannot be the embedding
	// source or embedding vector column.
	PrimaryKey string `json:"primary_key"`
}

func (s *CreateVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CustomTag struct {
	// Key field for an AI Search endpoint tag.
	Key string `json:"key"`
	// [Optional] Value field for an AI Search endpoint tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of successfully processed rows.
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDataResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDataResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for DeleteDataStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeleteDataStatus) Values() []DeleteDataStatus {
	return []DeleteDataStatus{
		DeleteDataStatusFailure,
		DeleteDataStatusPartialSuccess,
		DeleteDataStatusSuccess,
	}
}

// Type always returns DeleteDataStatus to satisfy [pflag.Value] interface
func (f *DeleteDataStatus) Type() string {
	return "DeleteDataStatus"
}

type DeleteDataVectorIndexRequest struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	IndexName string `json:"-" url:"-"`
	// Primary keys of the rows to delete. At least one key is required.
	PrimaryKeys []string `json:"-" url:"primary_keys"`
}

func (s *DeleteDataVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *DeleteDataResult `json:"result,omitempty"`
	// Status of the delete operation.
	Status DeleteDataStatus `json:"status,omitempty"`
}

func (s *DeleteDataVectorIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteEndpointRequest struct {
	// Name of the AI Search endpoint
	EndpointName string `json:"-" url:"-"`
}

func (s *DeleteEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteIndexRequest struct {
	// Three-level Unity Catalog name of the index.
	IndexName string `json:"-" url:"-"`
}

func (s *DeleteIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeltaSyncVectorIndexSpecRequest struct {
	// [Optional] Alias for columns_to_sync. Select the columns to include in
	// the vector index. If you leave this field blank, all columns from the
	// source table are included. The primary key column and embedding source
	// column or embedding vector column are always included. Only one of
	// columns_to_sync or columns_to_index may be specified.
	ColumnsToIndex []string `json:"columns_to_index,omitempty"`
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	ColumnsToSync []string `json:"columns_to_sync,omitempty"`
	// The source table column whose text Databricks embeds, together with the
	// embedding model endpoint to use. Set this for an index with
	// Databricks-managed embeddings. Exactly one of `embedding_source_columns`
	// or `embedding_vector_columns` must be specified unless the index subtype
	// is `FULL_TEXT`, and only one column is supported.
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The source table column that contains precomputed embedding vectors. Set
	// this instead of `embedding_source_columns` for an index with self-managed
	// embeddings.
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable string `json:"embedding_writeback_table,omitempty"`
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	PipelineType PipelineType `json:"pipeline_type,omitempty"`
	// Three-level Unity Catalog name of the source table that the index is
	// synced from.
	SourceTable string `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeltaSyncVectorIndexSpecRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncVectorIndexSpecRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeltaSyncVectorIndexSpecResponse struct {
	// [Optional] Alias for columns_to_sync. Select the columns to include in
	// the vector index. If you leave this field blank, all columns from the
	// source table are included. The primary key column and embedding source
	// column or embedding vector column are always included. Only one of
	// columns_to_sync or columns_to_index may be specified.
	ColumnsToIndex []string `json:"columns_to_index,omitempty"`
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	ColumnsToSync []string `json:"columns_to_sync,omitempty"`
	// The source table column whose text Databricks embeds, together with the
	// embedding model endpoint to use. Set this for an index with
	// Databricks-managed embeddings. Exactly one of `embedding_source_columns`
	// or `embedding_vector_columns` must be specified unless the index subtype
	// is `FULL_TEXT`, and only one column is supported.
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The source table column that contains precomputed embedding vectors. Set
	// this instead of `embedding_source_columns` for an index with self-managed
	// embeddings.
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable string `json:"embedding_writeback_table,omitempty"`
	// The ID of the pipeline that is used to sync the index.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	PipelineType PipelineType `json:"pipeline_type,omitempty"`
	// Three-level Unity Catalog name of the source table that the index is
	// synced from.
	SourceTable string `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeltaSyncVectorIndexSpecResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncVectorIndexSpecResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DirectAccessVectorIndexSpec struct {
	// Embedding source configuration used to embed `query_text` when the index
	// is queried with text, naming the embedding model endpoint to use. The
	// index stores the vectors you supply through upsert; this setting never
	// computes embeddings for them.
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The column that holds the embedding vectors you write to the index.
	// Exactly one column must be specified, and `schema_json` must declare it
	// as `array<float>` or `array<double>`.
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// JSON object that maps each column name of the index to its type,
	// serialized to a string. The primary key and the embedding vector column
	// must be present. Supported types are `integer`, `long`, `float`,
	// `double`, `boolean`, `string`, `date`, `timestamp`. Supported types for
	// vector column: `array<float>`, `array<double>`.
	SchemaJson string `json:"schema_json,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DirectAccessVectorIndexSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DirectAccessVectorIndexSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EmbeddingSourceColumn struct {
	// Name of the embedding model endpoint, used by default for both ingestion
	// and querying.
	EmbeddingModelEndpointName string `json:"embedding_model_endpoint_name,omitempty"`
	// Name of the embedding model endpoint which, if specified, is used for
	// querying (not ingestion).
	ModelEndpointNameForQuery string `json:"model_endpoint_name_for_query,omitempty"`
	// Name of the column that contains the text to embed.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EmbeddingSourceColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingSourceColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EmbeddingVectorColumn struct {
	// Dimension of the embedding vector
	EmbeddingDimension int `json:"embedding_dimension,omitempty"`
	// Name of the column that contains the embedding vectors.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EmbeddingVectorColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingVectorColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointInfo struct {
	// The user-selected budget policy id for the endpoint.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Time at which the endpoint was created, in milliseconds since the Unix
	// epoch.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// User or service principal that created the endpoint. Users are identified
	// by username and service principals by application ID.
	Creator string `json:"creator,omitempty"`
	// The custom tags assigned to the endpoint
	CustomTags []CustomTag `json:"custom_tags,omitempty"`
	// Budget policy ID in effect for the endpoint. This is the policy set on
	// the endpoint when one was set, and otherwise the default budget policy.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// Current status of the endpoint
	EndpointStatus *EndpointStatus `json:"endpoint_status,omitempty"`
	// Type of the endpoint. See
	// https://docs.databricks.com/aws/en/ai-search/ai-search#endpoint-options
	// for how the types differ.
	EndpointType EndpointType `json:"endpoint_type,omitempty"`
	// Unique identifier of the endpoint
	Id string `json:"id,omitempty"`
	// Time at which the endpoint was last updated, in milliseconds since the
	// Unix epoch.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// User or service principal that last updated the endpoint, in the same
	// form as creator.
	LastUpdatedUser string `json:"last_updated_user,omitempty"`
	// Name of the AI Search endpoint
	Name string `json:"name,omitempty"`
	// Number of indexes on the endpoint
	NumIndexes int `json:"num_indexes,omitempty"`
	// Scaling information for the endpoint. Present only for STANDARD endpoints
	// that have a Target QPS set.
	ScalingInfo *EndpointScalingInfo `json:"scaling_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Scaling information for a Standard endpoint: the current scaling state, the
// requested QPS target, and the progress of an in-progress scaling change.
type EndpointScalingInfo struct {
	// Reason the in-progress scaling change is delayed. When set,
	// `estimated_completion_timestamp` is unset. Unset when the change is not
	// delayed or no scaling change is in progress.
	DelayedReason EndpointScalingInfoScalingDelayReason `json:"delayed_reason,omitempty"`
	// Estimated time at which the in-progress scaling change will complete, in
	// milliseconds since the Unix epoch. The estimate may change as the scaling
	// change progresses. Unset when no scaling change is in progress, when no
	// estimate is available, or when the change is delayed (see
	// `delayed_reason`).
	EstimatedCompletionTimestamp int64 `json:"estimated_completion_timestamp,omitempty"`
	// Percentage of index initialization completed, from 0 to 100. Set only
	// during the `INITIALIZING_INDEX` phase; unset in other phases or when
	// progress is not available.
	IndexInitializationPercentage int `json:"index_initialization_percentage,omitempty"`
	// Current phase of the in-progress scaling change. Unset when no scaling
	// change is in progress or the phase is not available.
	Phase EndpointScalingInfoScalingPhase `json:"phase,omitempty"`
	// The requested QPS target for the endpoint. Best-effort; the system does
	// not guarantee this QPS will be achieved.
	RequestedTargetQps int64 `json:"requested_target_qps,omitempty"`
	// The current state of the scaling change request.
	State ScalingChangeState `json:"state,omitempty"`
	// Earliest time a new Target QPS change is accepted, in milliseconds since
	// the Unix epoch. Set only while a scaling change is in progress, the
	// change lock is on, and the endpoint has a recorded last throughput
	// change; may be in the past once the lock has expired.
	TargetQpsChangeAllowedAfterTimestamp int64 `json:"target_qps_change_allowed_after_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointScalingInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointScalingInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Reason an in-progress scaling change is delayed.
type EndpointScalingInfoScalingDelayReason string

const EndpointScalingInfoScalingDelayReasonCapacityUnavailable EndpointScalingInfoScalingDelayReason = `CAPACITY_UNAVAILABLE`

const EndpointScalingInfoScalingDelayReasonScaleDownBlocked EndpointScalingInfoScalingDelayReason = `SCALE_DOWN_BLOCKED`

const EndpointScalingInfoScalingDelayReasonTakingLongerThanExpected EndpointScalingInfoScalingDelayReason = `TAKING_LONGER_THAN_EXPECTED`

// String representation for [fmt.Print]
func (f *EndpointScalingInfoScalingDelayReason) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointScalingInfoScalingDelayReason) Set(v string) error {
	switch v {
	case `CAPACITY_UNAVAILABLE`, `SCALE_DOWN_BLOCKED`, `TAKING_LONGER_THAN_EXPECTED`:
		*f = EndpointScalingInfoScalingDelayReason(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAPACITY_UNAVAILABLE", "SCALE_DOWN_BLOCKED", "TAKING_LONGER_THAN_EXPECTED"`, v)
	}
}

// Values returns all possible values for EndpointScalingInfoScalingDelayReason.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointScalingInfoScalingDelayReason) Values() []EndpointScalingInfoScalingDelayReason {
	return []EndpointScalingInfoScalingDelayReason{
		EndpointScalingInfoScalingDelayReasonCapacityUnavailable,
		EndpointScalingInfoScalingDelayReasonScaleDownBlocked,
		EndpointScalingInfoScalingDelayReasonTakingLongerThanExpected,
	}
}

// Type always returns EndpointScalingInfoScalingDelayReason to satisfy [pflag.Value] interface
func (f *EndpointScalingInfoScalingDelayReason) Type() string {
	return "EndpointScalingInfoScalingDelayReason"
}

// Phase of an in-progress scaling change.
type EndpointScalingInfoScalingPhase string

const EndpointScalingInfoScalingPhaseAdjustingCapacity EndpointScalingInfoScalingPhase = `ADJUSTING_CAPACITY`

const EndpointScalingInfoScalingPhaseEstimating EndpointScalingInfoScalingPhase = `ESTIMATING`

const EndpointScalingInfoScalingPhaseFinalizing EndpointScalingInfoScalingPhase = `FINALIZING`

const EndpointScalingInfoScalingPhaseInitializingIndex EndpointScalingInfoScalingPhase = `INITIALIZING_INDEX`

// String representation for [fmt.Print]
func (f *EndpointScalingInfoScalingPhase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointScalingInfoScalingPhase) Set(v string) error {
	switch v {
	case `ADJUSTING_CAPACITY`, `ESTIMATING`, `FINALIZING`, `INITIALIZING_INDEX`:
		*f = EndpointScalingInfoScalingPhase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADJUSTING_CAPACITY", "ESTIMATING", "FINALIZING", "INITIALIZING_INDEX"`, v)
	}
}

// Values returns all possible values for EndpointScalingInfoScalingPhase.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointScalingInfoScalingPhase) Values() []EndpointScalingInfoScalingPhase {
	return []EndpointScalingInfoScalingPhase{
		EndpointScalingInfoScalingPhaseAdjustingCapacity,
		EndpointScalingInfoScalingPhaseEstimating,
		EndpointScalingInfoScalingPhaseFinalizing,
		EndpointScalingInfoScalingPhaseInitializingIndex,
	}
}

// Type always returns EndpointScalingInfoScalingPhase to satisfy [pflag.Value] interface
func (f *EndpointScalingInfoScalingPhase) Type() string {
	return "EndpointScalingInfoScalingPhase"
}

// Status information of an endpoint
type EndpointStatus struct {
	// Additional detail about the state, such as the reason the endpoint is
	// OFFLINE. Absent when there is nothing to report.
	Message string `json:"message,omitempty"`
	// Current state of the endpoint
	State EndpointStatusState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Current state of the endpoint
type EndpointStatusState string

const EndpointStatusStateDeleted EndpointStatusState = `DELETED`

const EndpointStatusStateOffline EndpointStatusState = `OFFLINE`

const EndpointStatusStateOnline EndpointStatusState = `ONLINE`

const EndpointStatusStateProvisioning EndpointStatusState = `PROVISIONING`

const EndpointStatusStateRedState EndpointStatusState = `RED_STATE`

const EndpointStatusStateStarting EndpointStatusState = `STARTING`

const EndpointStatusStateStopped EndpointStatusState = `STOPPED`

const EndpointStatusStateStopping EndpointStatusState = `STOPPING`

const EndpointStatusStateYellowState EndpointStatusState = `YELLOW_STATE`

// String representation for [fmt.Print]
func (f *EndpointStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStatusState) Set(v string) error {
	switch v {
	case `DELETED`, `OFFLINE`, `ONLINE`, `PROVISIONING`, `RED_STATE`, `STARTING`, `STOPPED`, `STOPPING`, `YELLOW_STATE`:
		*f = EndpointStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETED", "OFFLINE", "ONLINE", "PROVISIONING", "RED_STATE", "STARTING", "STOPPED", "STOPPING", "YELLOW_STATE"`, v)
	}
}

// Values returns all possible values for EndpointStatusState.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointStatusState) Values() []EndpointStatusState {
	return []EndpointStatusState{
		EndpointStatusStateDeleted,
		EndpointStatusStateOffline,
		EndpointStatusStateOnline,
		EndpointStatusStateProvisioning,
		EndpointStatusStateRedState,
		EndpointStatusStateStarting,
		EndpointStatusStateStopped,
		EndpointStatusStateStopping,
		EndpointStatusStateYellowState,
	}
}

// Type always returns EndpointStatusState to satisfy [pflag.Value] interface
func (f *EndpointStatusState) Type() string {
	return "EndpointStatusState"
}

// Type of endpoint.
type EndpointType string

const EndpointTypeStandard EndpointType = `STANDARD`

const EndpointTypeStorageOptimized EndpointType = `STORAGE_OPTIMIZED`

// String representation for [fmt.Print]
func (f *EndpointType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointType) Set(v string) error {
	switch v {
	case `STANDARD`, `STORAGE_OPTIMIZED`:
		*f = EndpointType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "STANDARD", "STORAGE_OPTIMIZED"`, v)
	}
}

// Values returns all possible values for EndpointType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointType) Values() []EndpointType {
	return []EndpointType{
		EndpointTypeStandard,
		EndpointTypeStorageOptimized,
	}
}

// Type always returns EndpointType to satisfy [pflag.Value] interface
func (f *EndpointType) Type() string {
	return "EndpointType"
}

// Facet aggregation rows returned by a query.
type FacetResultData struct {
	// Facet rows. Each row is `[facet_column_name, value_or_range, count]`.
	FacetArray [][]string `json:"facet_array,omitempty"`
	// Number of facet rows returned.
	FacetRowCount int `json:"facet_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FacetResultData) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FacetResultData) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Failure details for an index or its latest ingestion update.
type FailedStatus struct {
	// Error code associated with the error message, when available.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message, when available.
	ErrorMessage string `json:"error_message,omitempty"`
	// [Standard endpoint] timestamp when the last processed commit was
	// committed to the source table. [Storage Optimized endpoint] timestamp
	// when the last processed commit was processed by the pipeline.
	LastProcessedCommitTimestamp *time.Time `json:"last_processed_commit_timestamp,omitempty"`
	// Last source Delta table version fully processed while the index is still
	// queryable.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FailedStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FailedStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetEndpointRequest struct {
	// Name of the endpoint
	EndpointName string `json:"-" url:"-"`
}

func (s *GetEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetIndexRequest struct {
	// If true, the URL returned for the index is guaranteed to be compatible
	// with the reranker. Currently this means we return the CP URL regardless
	// of how the index is being accessed. If not set or set to false, the URL
	// may still be compatible with the reranker depending on what URL we
	// return.
	EnsureRerankerCompatible bool `json:"-" url:"ensure_reranker_compatible,omitempty"`
	// Three-level Unity Catalog name of the index.
	IndexName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetVectorSearchEndpointPermissionLevelsRequest struct {
	// The vector search endpoint for which to get or manage permissions.
	EndpointId string `json:"-" url:"-"`
}

func (s *GetVectorSearchEndpointPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetVectorSearchEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []VectorSearchEndpointPermissionsDescription `json:"permission_levels,omitempty"`
}

func (s *GetVectorSearchEndpointPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetVectorSearchEndpointPermissionsRequest struct {
	// The vector search endpoint for which to get or manage permissions.
	EndpointId string `json:"-" url:"-"`
}

func (s *GetVectorSearchEndpointPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Lifecycle and ingestion state of an index.
type IndexStatus string

const IndexStatusDownloading IndexStatus = `DOWNLOADING`

const IndexStatusOffline IndexStatus = `OFFLINE`

const IndexStatusOfflineFailed IndexStatus = `OFFLINE_FAILED`

const IndexStatusOnline IndexStatus = `ONLINE`

const IndexStatusOnlineContinuousUpdate IndexStatus = `ONLINE_CONTINUOUS_UPDATE`

const IndexStatusOnlineDirectAccess IndexStatus = `ONLINE_DIRECT_ACCESS`

const IndexStatusOnlineNoPendingUpdate IndexStatus = `ONLINE_NO_PENDING_UPDATE`

const IndexStatusOnlinePipelineFailed IndexStatus = `ONLINE_PIPELINE_FAILED`

const IndexStatusOnlineTriggeredUpdate IndexStatus = `ONLINE_TRIGGERED_UPDATE`

const IndexStatusOnlineUpdatingPipelineResources IndexStatus = `ONLINE_UPDATING_PIPELINE_RESOURCES`

const IndexStatusProvisioning IndexStatus = `PROVISIONING`

const IndexStatusProvisioningEndpoint IndexStatus = `PROVISIONING_ENDPOINT`

const IndexStatusProvisioningIndex IndexStatus = `PROVISIONING_INDEX`

const IndexStatusProvisioningInitialSnapshot IndexStatus = `PROVISIONING_INITIAL_SNAPSHOT`

const IndexStatusProvisioningPipelineResources IndexStatus = `PROVISIONING_PIPELINE_RESOURCES`

// String representation for [fmt.Print]
func (f *IndexStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IndexStatus) Set(v string) error {
	switch v {
	case `DOWNLOADING`, `OFFLINE`, `OFFLINE_FAILED`, `ONLINE`, `ONLINE_CONTINUOUS_UPDATE`, `ONLINE_DIRECT_ACCESS`, `ONLINE_NO_PENDING_UPDATE`, `ONLINE_PIPELINE_FAILED`, `ONLINE_TRIGGERED_UPDATE`, `ONLINE_UPDATING_PIPELINE_RESOURCES`, `PROVISIONING`, `PROVISIONING_ENDPOINT`, `PROVISIONING_INDEX`, `PROVISIONING_INITIAL_SNAPSHOT`, `PROVISIONING_PIPELINE_RESOURCES`:
		*f = IndexStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DOWNLOADING", "OFFLINE", "OFFLINE_FAILED", "ONLINE", "ONLINE_CONTINUOUS_UPDATE", "ONLINE_DIRECT_ACCESS", "ONLINE_NO_PENDING_UPDATE", "ONLINE_PIPELINE_FAILED", "ONLINE_TRIGGERED_UPDATE", "ONLINE_UPDATING_PIPELINE_RESOURCES", "PROVISIONING", "PROVISIONING_ENDPOINT", "PROVISIONING_INDEX", "PROVISIONING_INITIAL_SNAPSHOT", "PROVISIONING_PIPELINE_RESOURCES"`, v)
	}
}

// Values returns all possible values for IndexStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *IndexStatus) Values() []IndexStatus {
	return []IndexStatus{
		IndexStatusDownloading,
		IndexStatusOffline,
		IndexStatusOfflineFailed,
		IndexStatusOnline,
		IndexStatusOnlineContinuousUpdate,
		IndexStatusOnlineDirectAccess,
		IndexStatusOnlineNoPendingUpdate,
		IndexStatusOnlinePipelineFailed,
		IndexStatusOnlineTriggeredUpdate,
		IndexStatusOnlineUpdatingPipelineResources,
		IndexStatusProvisioning,
		IndexStatusProvisioningEndpoint,
		IndexStatusProvisioningIndex,
		IndexStatusProvisioningInitialSnapshot,
		IndexStatusProvisioningPipelineResources,
	}
}

// Type always returns IndexStatus to satisfy [pflag.Value] interface
func (f *IndexStatus) Type() string {
	return "IndexStatus"
}

// The subtype of the AI Search index, determining the indexing and retrieval
// strategy. - `VECTOR`: Not supported. Use `HYBRID` instead. - `FULL_TEXT`: An
// index that uses full-text search without vector embeddings. - `HYBRID`: An
// index that uses vector embeddings for similarity search and hybrid search.
type IndexSubtype string

// An index that uses full-text search without vector embeddings.
const IndexSubtypeFullText IndexSubtype = `FULL_TEXT`

// An index that uses vector embeddings for similarity search and hybrid search.
const IndexSubtypeHybrid IndexSubtype = `HYBRID`

// Not supported. Use `HYBRID` instead.
const IndexSubtypeVector IndexSubtype = `VECTOR`

// String representation for [fmt.Print]
func (f *IndexSubtype) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IndexSubtype) Set(v string) error {
	switch v {
	case `FULL_TEXT`, `HYBRID`, `VECTOR`:
		*f = IndexSubtype(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FULL_TEXT", "HYBRID", "VECTOR"`, v)
	}
}

// Values returns all possible values for IndexSubtype.
//
// There is no guarantee on the order of the values in the slice.
func (f *IndexSubtype) Values() []IndexSubtype {
	return []IndexSubtype{
		IndexSubtypeFullText,
		IndexSubtypeHybrid,
		IndexSubtypeVector,
	}
}

// Type always returns IndexSubtype to satisfy [pflag.Value] interface
func (f *IndexSubtype) Type() string {
	return "IndexSubtype"
}

type ListEndpointResponse struct {
	// An array of Endpoint objects
	Endpoints []EndpointInfo `json:"endpoints,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsRequest struct {
	// Token for pagination
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListIndexesRequest struct {
	// Name of the endpoint whose indexes to list.
	EndpointName string `json:"-" url:"endpoint_name"`
	// Page token returned by a previous call, used to get the next page of
	// results.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListIndexesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListIndexesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListValue struct {
	// Repeated field of dynamically typed values.
	Values []Value `json:"values,omitempty"`
}

func (s *ListValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListVectorIndexesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Indexes served by the endpoint. Each entry is a summary without the index
	// specification and status; use Get an index to retrieve them.
	VectorIndexes []MiniVectorIndex `json:"vector_indexes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	Key string `json:"key,omitempty"`
	// Column value, nullable.
	Value *Value `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MapStringValueEntry) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MapStringValueEntry) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metric specification
type Metric struct {
	// Labels that restrict the metric to matching series, for example to a
	// single index. Endpoint-level metrics do not accept labels.
	Labels []MetricLabel `json:"labels,omitempty"`
	// Metric name. Supported metrics include: `EndpointHealthStatus`,
	// `EndpointConcurrency`, `IndexRequestCount`, `IndexRequestLatency`,
	// `IndexRequestErrorSourceCount`, `IndexRequestGuidanceCount`,
	// `IndexRequestGuidancePeakRate`, `IndexIngestedRows`, and
	// `IndexEmbeddedRows`. An unknown name is rejected with an error that lists
	// the available metric names.
	Name string `json:"name,omitempty"`
	// Percentile for latency metrics, such as `IndexRequestLatency`, which
	// require it. One of 0.5, 0.75, 0.9, 0.95, or 0.99.
	Percentile float64 `json:"percentile,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Metric) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Metric) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Label for a metric
type MetricLabel struct {
	// Label name, such as `error_source`, `failure_cause`, `operation`,
	// `request_path`, `query_type`, `response_code`, or `index_name`. Supported
	// labels depend on the metric.
	Name string `json:"name,omitempty"`
	// Label value, as an RE2 regular expression. The label also splits the
	// result by its values, so `.+` returns one series per value. Some labels,
	// such as `operation`, accept only `.+`, `.*`, or listed values joined by
	// `|`.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MetricLabel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetricLabel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Single metric value at a specific timestamp
type MetricValue struct {
	// Timestamp of the metric value, in seconds since the Unix epoch.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Value of the metric at `timestamp`.
	Value float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MetricValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetricValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Collection of metric values for a specific metric
type MetricValues struct {
	// The metric this time series was computed for, including the labels that
	// identify the series.
	Metric *Metric `json:"metric,omitempty"`
	// Each entry is one data point of the time series.
	Values []MetricValue `json:"values,omitempty"`
}

func (s *MetricValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An AI Search index stores vectors and metadata on an endpoint and serves
// vector similarity, full-text and hybrid search queries over them. A Delta
// Sync index is kept in sync with a source table, with embeddings either
// computed by Databricks from a text column or supplied by you in a vector
// column; a Direct Vector Access index holds rows that you write through the
// upsert and delete data APIs. See
// https://docs.databricks.com/aws/en/ai-search/create-ai-search for how to
// choose between them.
type MiniVectorIndex struct {
	// The user who created the index.
	Creator string `json:"creator,omitempty"`
	// ID of the endpoint associated with the index.
	EndpointId string `json:"endpoint_id,omitempty"`
	// Name of the endpoint that serves the index.
	EndpointName string `json:"endpoint_name,omitempty"`
	// Subtype of the index. A `HYBRID` index supports vector similarity,
	// full-text and hybrid queries; a `FULL_TEXT` index has no embeddings and
	// supports only full-text queries.
	IndexSubtype IndexSubtype `json:"index_subtype,omitempty"`
	// Type of the index, which determines whether `delta_sync_index_spec` or
	// `direct_access_index_spec` describes it.
	IndexType VectorIndexType `json:"index_type,omitempty"`
	// Three-level Unity Catalog name of the index.
	Name string `json:"name,omitempty"`
	// Name of the column that uniquely identifies each row of the index.
	PrimaryKey string `json:"primary_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MiniVectorIndex) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MiniVectorIndex) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PatchEndpointBudgetPolicyRequest struct {
	// The budget policy id to be applied
	BudgetPolicyId string `json:"budget_policy_id"`
	// Name of the AI Search endpoint
	EndpointName string `json:"-" url:"-"`
}

func (s *PatchEndpointBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PatchEndpointBudgetPolicyResponse struct {
	// The requested budget policy ID, returned when the endpoint was already
	// using it.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The budget policy applied to the AI Search endpoint.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PatchEndpointBudgetPolicyResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PatchEndpointBudgetPolicyResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PatchEndpointRequest struct {
	// Name of the AI Search endpoint
	EndpointName string `json:"-" url:"-"`
	// Target QPS for the endpoint. Best-effort; the system does not guarantee
	// this QPS will be achieved. Set to -1 to clear the Target QPS, as if it
	// had never been set.
	TargetQps int64 `json:"target_qps,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PatchEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PatchEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the triggered
// execution mode, the system stops processing after successfully refreshing the
// source table in the pipeline once, ensuring the table is updated based on the
// data available when the update started. - `CONTINUOUS`: If the pipeline uses
// continuous execution, the pipeline processes new data as it arrives in the
// source table to keep vector index fresh.
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

// Values returns all possible values for PipelineType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PipelineType) Values() []PipelineType {
	return []PipelineType{
		PipelineTypeContinuous,
		PipelineTypeTriggered,
	}
}

// Type always returns PipelineType to satisfy [pflag.Value] interface
func (f *PipelineType) Type() string {
	return "PipelineType"
}

// Request payload for getting next page of results.
type QueryVectorIndexNextPageRequest struct {
	// Name of the endpoint that serves the index.
	EndpointName string `json:"endpoint_name,omitempty"`
	// Name of the vector index to query.
	IndexName string `json:"-" url:"-"`
	// Page token returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` API.
	PageToken string `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryVectorIndexNextPageRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryVectorIndexNextPageRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryVectorIndexRequest struct {
	// Names of the columns to return for each result. At least one column is
	// required, and only columns that are part of the index can be returned.
	Columns []string `json:"columns"`
	// Column names used to retrieve data to send to the reranker.
	ColumnsToRerank []string `json:"columns_to_rerank,omitempty"`
	// Facets to compute over the matched results. Each entry has one of these
	// forms: `"<column>"` - top 10 distinct values by count `"<column> TOP
	// <n>"` - top n distinct values, where n > 0 `"<column> BUCKETS
	// [[from,to],...]"` - inclusive numeric ranges `TOP` and `BUCKETS` are
	// case-insensitive. A column may appear at most once.
	Facets []string `json:"facets,omitempty"`
	// Filters to apply to the results, as a JSON object serialized to a string.
	// Each key is a column name, optionally followed by a space and an
	// operator, and each value is the value or list of values to compare the
	// column against. A key without an operator matches the value exactly, or
	// any of the listed values. Supported operators are `NOT`, `<`, `<=`, `>`,
	// `>=`, `OR` and `LIKE`; see
	// https://docs.databricks.com/aws/en/ai-search/query-ai-search for their
	// semantics. This field is not supported on storage-optimized endpoints.
	FiltersJson string `json:"filters_json,omitempty"`
	// Name of the vector index to query.
	IndexName string `json:"-" url:"-"`
	// Maximum number of results to return. It must be greater than 0 and
	// defaults to 10. By default, a `HYBRID` query can return at most 200
	// results.
	NumResults int `json:"num_results,omitempty"`
	// Text columns to search for `query_text`. When empty, all text columns are
	// searched.
	QueryColumns []string `json:"query_columns,omitempty"`
	// Text to search for. Use it for an index with Databricks-managed
	// embeddings, which embeds the text with the index's embedding model
	// endpoint, and for `FULL_TEXT` queries on any index. A `HYBRID` query on
	// an index with managed embeddings requires only `query_text`.
	QueryText string `json:"query_text,omitempty"`
	// The query type to use. Choices are `ANN` and `HYBRID` and `FULL_TEXT`.
	// Defaults to `ANN`.
	QueryType string `json:"query_type,omitempty"`
	// Embedding vector to search for. Use it for an index whose embeddings you
	// compute yourself: a Direct Vector Access index or a Delta Sync index with
	// `embedding_vector_columns`. A `HYBRID` query on a Delta Sync index with
	// `embedding_vector_columns`, or on a Direct Vector Access index without
	// `embedding_source_columns`, requires both `query_vector` and
	// `query_text`. A Direct Vector Access index with
	// `embedding_source_columns` can be queried with `query_text` alone.
	QueryVector []float64 `json:"query_vector,omitempty"`
	// If set, the top 50 results are reranked with the Databricks Reranker
	// model before returning the `num_results` results to the user. The setting
	// `columns_to_rerank` selects which columns are used for reranking. For
	// each datapoint, the columns selected are concatenated before being sent
	// to the reranking model. See
	// https://docs.databricks.com/aws/en/vector-search/query-vector-search#rerank
	// for more information.
	Reranker *RerankerConfig `json:"reranker,omitempty"`
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold float64 `json:"score_threshold,omitempty"`
	// Sort results by column values instead of the default relevance ordering.
	// Each clause has the form `"<column> ASC"` or `"<column> DESC"`, for
	// example `["rating DESC", "price ASC"]`.
	SortColumns []string `json:"sort_columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryVectorIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryVectorIndexResponse struct {
	// Facet aggregation rows returned by a query.
	FacetResult *FacetResultData `json:"facet_result,omitempty"`
	// Metadata about the result set.
	Manifest *ResultManifest `json:"manifest,omitempty"`
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	// The maximum number of results that can be returned is 10,000.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Data returned in the query result.
	Result *ResultData `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryVectorIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryVectorIndexResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RerankerConfig struct {
	// Reranker identifier: - When model_type=BASE/UNSPECIFIED: must be
	// "databricks_reranker". - When model_type=FINETUNED: the Model Serving
	// endpoint name hosting a finetuned reranker.
	Model string `json:"model,omitempty"`
	// Parameters that control how the reranker processes the query results.
	Parameters *RerankerConfigRerankerParameters `json:"parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RerankerConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RerankerConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RerankerConfigRerankerParameters struct {
	// Names of the columns whose values are sent to the reranker. List the most
	// relevant column first, since the order affects the ranking.
	ColumnsToRerank []string `json:"columns_to_rerank,omitempty"`
}

func (s *RerankerConfigRerankerParameters) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Data returned in the query result.
type ResultData struct {
	// Data rows returned in the query.
	DataArray [][]string `json:"data_array,omitempty"`
	// Number of rows in the result set.
	RowCount int `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	ColumnCount int `json:"column_count,omitempty"`
	// Information about each column in the result set.
	Columns []ColumnInfo `json:"columns,omitempty"`
	// Number of columns in `facet_result`.
	FacetColumnCount int `json:"facet_column_count,omitempty"`
	// Information about each column in `facet_result`.
	FacetColumns []ColumnInfo `json:"facet_columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResultManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to retrieve user-visible metrics
type RetrieveUserVisibleMetricsRequest struct {
	// End of the time range. Defaults to the current time. Must be later than
	// `start_time`.
	EndTime string `json:"end_time,omitempty"`
	// Requested interval between data points, in seconds. Defaults to 60, which
	// is also the minimum. The server may return a finer or coarser interval:
	// 60 seconds only for time ranges of 1 hour or less, otherwise 300 seconds
	// or more. Read the interval from the returned timestamps.
	GranularityInSeconds int `json:"granularity_in_seconds,omitempty"`
	// Metrics to retrieve. Each entry names a metric and optionally narrows it
	// with labels or a percentile. An unknown metric name fails the whole
	// request, and an empty list returns no series.
	Metrics []Metric `json:"metrics,omitempty"`
	// Name of the endpoint whose metrics to retrieve.
	Name string `json:"-" url:"-"`
	// Not used. The response returns all data points in one page.
	PageToken string `json:"page_token,omitempty"`
	// Start of the time range. Required. Must be within the last 169 hours.
	StartTime string `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RetrieveUserVisibleMetricsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RetrieveUserVisibleMetricsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing user-visible metrics
type RetrieveUserVisibleMetricsResponse struct {
	// One time series for each requested metric and each combination of label
	// values it was broken down by.
	MetricValues []MetricValues `json:"metric_values,omitempty"`
	// Not set. The response returns all data points in one page.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RetrieveUserVisibleMetricsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RetrieveUserVisibleMetricsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ScalingChangeState string

const ScalingChangeStateScalingChangeApplied ScalingChangeState = `SCALING_CHANGE_APPLIED`

const ScalingChangeStateScalingChangeInProgress ScalingChangeState = `SCALING_CHANGE_IN_PROGRESS`

const ScalingChangeStateScalingChangeUnspecified ScalingChangeState = `SCALING_CHANGE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *ScalingChangeState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ScalingChangeState) Set(v string) error {
	switch v {
	case `SCALING_CHANGE_APPLIED`, `SCALING_CHANGE_IN_PROGRESS`, `SCALING_CHANGE_UNSPECIFIED`:
		*f = ScalingChangeState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SCALING_CHANGE_APPLIED", "SCALING_CHANGE_IN_PROGRESS", "SCALING_CHANGE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for ScalingChangeState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ScalingChangeState) Values() []ScalingChangeState {
	return []ScalingChangeState{
		ScalingChangeStateScalingChangeApplied,
		ScalingChangeStateScalingChangeInProgress,
		ScalingChangeStateScalingChangeUnspecified,
	}
}

// Type always returns ScalingChangeState to satisfy [pflag.Value] interface
func (f *ScalingChangeState) Type() string {
	return "ScalingChangeState"
}

type ScanVectorIndexRequest struct {
	// Name of the vector index to scan.
	IndexName string `json:"-" url:"-"`
	// Primary key of the last entry returned in the previous scan.
	LastPrimaryKey string `json:"last_primary_key,omitempty"`
	// Number of results to return. Defaults to 10.
	NumResults int `json:"num_results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	Data []Struct `json:"data,omitempty"`
	// Primary key of the last entry.
	LastPrimaryKey string `json:"last_primary_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ScanVectorIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ScanVectorIndexResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StartEndpointRequest struct {
	// Name of the AI Search endpoint
	EndpointName string `json:"-" url:"-"`
}

func (s *StartEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type StopEndpointRequest struct {
	// Name of the AI Search endpoint
	EndpointName string `json:"-" url:"-"`
}

func (s *StopEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Struct struct {
	// Data entry, corresponding to a row in a vector index.
	Fields []MapStringValueEntry `json:"fields,omitempty"`
}

func (s *Struct) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SyncIndexRequest struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName string `json:"-" url:"-"`
}

func (s *SyncIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateEndpointCustomTagsRequest struct {
	// The new custom tags for the AI Search endpoint
	CustomTags []CustomTag `json:"custom_tags"`
	// Name of the AI Search endpoint
	EndpointName string `json:"-" url:"-"`
}

func (s *UpdateEndpointCustomTagsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateEndpointCustomTagsResponse struct {
	// All the custom tags that are applied to the AI Search endpoint.
	CustomTags []CustomTag `json:"custom_tags,omitempty"`
	// The name of the AI Search endpoint whose custom tags were updated.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateEndpointCustomTagsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateEndpointCustomTagsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpsertDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of successfully processed rows.
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpsertDataResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpsertDataResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for UpsertDataStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpsertDataStatus) Values() []UpsertDataStatus {
	return []UpsertDataStatus{
		UpsertDataStatusFailure,
		UpsertDataStatusPartialSuccess,
		UpsertDataStatusSuccess,
	}
}

// Type always returns UpsertDataStatus to satisfy [pflag.Value] interface
func (f *UpsertDataStatus) Type() string {
	return "UpsertDataStatus"
}

type UpsertDataVectorIndexRequest struct {
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	IndexName string `json:"-" url:"-"`
	// Rows to upsert, as a JSON array of objects serialized to a string. Each
	// object maps column names from `schema_json` to values and must include
	// the primary key. A row whose primary key already exists in the index is
	// replaced.
	InputsJson string `json:"inputs_json"`
}

func (s *UpsertDataVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpsertDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *UpsertDataResult `json:"result,omitempty"`
	// Status of the upsert operation.
	Status UpsertDataStatus `json:"status,omitempty"`
}

func (s *UpsertDataVectorIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Value struct {
	BoolValue bool `json:"bool_value,omitempty"`

	ListValue *ListValue `json:"list_value,omitempty"`

	NumberValue float64 `json:"number_value,omitempty"`

	StringValue string `json:"string_value,omitempty"`

	StructValue *Struct `json:"struct_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Value) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Value) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An AI Search index stores vectors and metadata on an endpoint and serves
// vector similarity, full-text and hybrid search queries over them. A Delta
// Sync index is kept in sync with a source table, with embeddings either
// computed by Databricks from a text column or supplied by you in a vector
// column; a Direct Vector Access index holds rows that you write through the
// upsert and delete data APIs. See
// https://docs.databricks.com/aws/en/ai-search/create-ai-search for how to
// choose between them.
type VectorIndex struct {
	// The user who created the index.
	Creator string `json:"creator,omitempty"`
	// Specification of a Delta Sync index. Present only when `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecResponse `json:"delta_sync_index_spec,omitempty"`
	// Specification of a Direct Vector Access index. Present only when
	// `index_type` is `DIRECT_ACCESS`.
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `json:"direct_access_index_spec,omitempty"`
	// ID of the endpoint associated with the index.
	EndpointId string `json:"endpoint_id,omitempty"`
	// Name of the endpoint that serves the index.
	EndpointName string `json:"endpoint_name,omitempty"`
	// Subtype of the index. A `HYBRID` index supports vector similarity,
	// full-text and hybrid queries; a `FULL_TEXT` index has no embeddings and
	// supports only full-text queries.
	IndexSubtype IndexSubtype `json:"index_subtype,omitempty"`
	// Type of the index, which determines whether `delta_sync_index_spec` or
	// `direct_access_index_spec` describes it.
	IndexType VectorIndexType `json:"index_type,omitempty"`
	// Three-level Unity Catalog name of the index.
	Name string `json:"name,omitempty"`
	// Name of the column that uniquely identifies each row of the index.
	PrimaryKey string `json:"primary_key,omitempty"`
	// Current status of the index, including whether it is ready to serve
	// queries and how many rows it holds.
	Status *VectorIndexStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VectorIndex) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorIndex) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VectorIndexStatus struct {
	// Lifecycle and ingestion state. An online index can still have a failed
	// update.
	DetailedState IndexStatus `json:"detailed_state,omitempty"`
	// Failure details, when available, for OFFLINE_FAILED or
	// ONLINE_PIPELINE_FAILED.
	FailedStatus *FailedStatus `json:"failed_status,omitempty"`
	// Index API Url to be used to perform operations on the index
	IndexUrl string `json:"index_url,omitempty"`
	// Number of rows indexed
	IndexedRowCount int64 `json:"indexed_row_count,omitempty"`
	// Human-readable description of the current status.
	Message string `json:"message,omitempty"`
	// Whether the index can serve queries. This does not mean the latest update
	// succeeded or that all source-table changes have been indexed; check
	// detailed_state for update failures.
	Ready bool `json:"ready,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VectorIndexStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorIndexStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// There are 2 types of AI Search indexes: - `DELTA_SYNC`: An index that
// automatically syncs with a source Delta Table, automatically and
// incrementally updating the index as the underlying data in the Delta Table
// changes. - `DIRECT_ACCESS`: An index that supports direct read and write of
// vectors and metadata through our REST and SDK APIs. With this model, the user
// manages index updates.
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

// Values returns all possible values for VectorIndexType.
//
// There is no guarantee on the order of the values in the slice.
func (f *VectorIndexType) Values() []VectorIndexType {
	return []VectorIndexType{
		VectorIndexTypeDeltaSync,
		VectorIndexTypeDirectAccess,
	}
}

// Type always returns VectorIndexType to satisfy [pflag.Value] interface
func (f *VectorIndexType) Type() string {
	return "VectorIndexType"
}

type VectorSearchEndpointAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel VectorSearchEndpointPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VectorSearchEndpointAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorSearchEndpointAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VectorSearchEndpointAccessControlResponse struct {
	// All permissions.
	AllPermissions []VectorSearchEndpointPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VectorSearchEndpointAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorSearchEndpointAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VectorSearchEndpointPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel VectorSearchEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VectorSearchEndpointPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorSearchEndpointPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type VectorSearchEndpointPermissionLevel string

const VectorSearchEndpointPermissionLevelCanCreate VectorSearchEndpointPermissionLevel = `CAN_CREATE`

const VectorSearchEndpointPermissionLevelCanManage VectorSearchEndpointPermissionLevel = `CAN_MANAGE`

const VectorSearchEndpointPermissionLevelCanUse VectorSearchEndpointPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (f *VectorSearchEndpointPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VectorSearchEndpointPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_CREATE`, `CAN_MANAGE`, `CAN_USE`:
		*f = VectorSearchEndpointPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_CREATE", "CAN_MANAGE", "CAN_USE"`, v)
	}
}

// Values returns all possible values for VectorSearchEndpointPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *VectorSearchEndpointPermissionLevel) Values() []VectorSearchEndpointPermissionLevel {
	return []VectorSearchEndpointPermissionLevel{
		VectorSearchEndpointPermissionLevelCanCreate,
		VectorSearchEndpointPermissionLevelCanManage,
		VectorSearchEndpointPermissionLevelCanUse,
	}
}

// Type always returns VectorSearchEndpointPermissionLevel to satisfy [pflag.Value] interface
func (f *VectorSearchEndpointPermissionLevel) Type() string {
	return "VectorSearchEndpointPermissionLevel"
}

type VectorSearchEndpointPermissions struct {
	AccessControlList []VectorSearchEndpointAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VectorSearchEndpointPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorSearchEndpointPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VectorSearchEndpointPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel VectorSearchEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VectorSearchEndpointPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorSearchEndpointPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VectorSearchEndpointPermissionsRequest struct {
	AccessControlList []VectorSearchEndpointAccessControlRequest `json:"access_control_list,omitempty"`
	// The vector search endpoint for which to get or manage permissions.
	EndpointId string `json:"-" url:"-"`
}

func (s *VectorSearchEndpointPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
