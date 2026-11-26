// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aisearch

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// Column information (name and data type) for an index column. Surfaced on
// `Index.column_info`.
type ColumnInfo struct {
	// Name of the column.
	Name string `json:"name,omitempty"`
	// Data type of the column (e.g., "string", "int", "array<float>").
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateEndpointRequest struct {
	// The Endpoint resource to create. Fields other than `endpoint.name` carry
	// the desired configuration; `endpoint.name` is server-assigned from
	// `parent` and `endpoint_id`.
	Endpoint Endpoint `json:"endpoint"`
	// Short name of the endpoint, which becomes the last segment of `name`. The
	// server rejects a request without it. The name can have at most 50
	// characters, can contain only lowercase letters, digits, hyphens, periods,
	// and underscores, and must start and end with a letter or digit.
	EndpointId string `json:"-" url:"endpoint_id,omitempty"`
	// The Workspace where this Endpoint will be created. Format:
	// `workspaces/{workspace_id}`
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateIndexRequest struct {
	// The Index resource to create. Fields other than `index.name` carry the
	// desired configuration; `index.name` is server-assigned from `parent` and
	// `index_id`.
	Index Index `json:"index"`
	// Three-level Unity Catalog table name for the new index, in the form
	// `catalog.schema.name`. The server composes the full `Index.name` as
	// `{parent}/indexes/{index_id}` and rejects an empty value.
	IndexId string `json:"-" url:"index_id,omitempty"`
	// The Endpoint where this Index will be created. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// User-defined key/value tag attached to an AI Search endpoint for cost
// attribution and access control.
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

// Per-row outcome of a data-plane upsert or delete operation.
type DataModificationResult struct {
	// Primary keys of rows that failed to process.
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of rows processed successfully.
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataModificationResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataModificationResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Overall outcome of a data-plane upsert or delete. Mirrors the legacy
// `databricks.brickindexscheduler.UpsertDeleteDataStatus` value-for-value.
type DataModificationStatus string

const DataModificationStatusFailure DataModificationStatus = `FAILURE`

const DataModificationStatusPartialSuccess DataModificationStatus = `PARTIAL_SUCCESS`

const DataModificationStatusSuccess DataModificationStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *DataModificationStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataModificationStatus) Set(v string) error {
	switch v {
	case `FAILURE`, `PARTIAL_SUCCESS`, `SUCCESS`:
		*f = DataModificationStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILURE", "PARTIAL_SUCCESS", "SUCCESS"`, v)
	}
}

// Values returns all possible values for DataModificationStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataModificationStatus) Values() []DataModificationStatus {
	return []DataModificationStatus{
		DataModificationStatusFailure,
		DataModificationStatusPartialSuccess,
		DataModificationStatusSuccess,
	}
}

// Type always returns DataModificationStatus to satisfy [pflag.Value] interface
func (f *DataModificationStatus) Type() string {
	return "DataModificationStatus"
}

type DeleteEndpointRequest struct {
	// Full resource name of the endpoint to delete. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`
	Name string `json:"-" url:"-"`
}

func (s *DeleteEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteIndexRequest struct {
	// Full resource name of the index to delete. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name string `json:"-" url:"-"`
}

func (s *DeleteIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Specification for a Delta Sync index — the index is kept in sync with a
// source Delta table.
type DeltaSyncIndexSpec struct {
	// Columns of the source table to include in the index. When empty, all
	// columns are synced. The primary key and the embedding source or vector
	// column are always synced. Only synced columns can be returned in query
	// results or used in filters.
	ColumnsToSync []string `json:"columns_to_sync,omitempty"`
	// Columns whose text Databricks embeds during sync, using the model
	// endpoint named on each column. Set either this field or
	// `embedding_vector_columns`, not both.
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// Columns of the source table that already hold precomputed embedding
	// vectors. Set either this field or `embedding_source_columns`, not both.
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// Three-level name of a Delta table to write the synced rows and computed
	// embeddings to. It applies only when embeddings are computed from
	// `embedding_source_columns`, and the table must not already exist.
	EmbeddingWritebackTable string `json:"embedding_writeback_table,omitempty"`
	// The ID of the pipeline that is used to sync the index.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Sync mode of the pipeline that keeps the index up to date. `TRIGGERED`
	// syncs at creation and then each time SyncIndex is called; `CONTINUOUS`
	// processes source changes as they arrive. Storage Optimized endpoints
	// accept only `TRIGGERED`, and an unset value is rejected.
	PipelineType PipelineType `json:"pipeline_type"`
	// Three-level Unity Catalog name of the source Delta table, in the form
	// `catalog.schema.table`.
	SourceTable string `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeltaSyncIndexSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncIndexSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Specification for a Direct Access index — the customer manages vectors and
// metadata directly.
type DirectAccessIndexSpec struct {
	// Columns that contain the text to embed, each with the model endpoint that
	// computes its embeddings.
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// Columns that hold precomputed embedding vectors. A single vector column
	// is supported, and it cannot be the primary key.
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// The schema of the index in JSON format, mapping each column name to its
	// type. Supported types are `integer`, `long`, `float`, `double`,
	// `boolean`, `string`, `date`, `timestamp`. Supported types for vector
	// columns: `array<float>`, `array<double>`. The schema must include the
	// primary key column and the embedding vector column.
	SchemaJson string `json:"schema_json,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DirectAccessIndexSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DirectAccessIndexSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Name of an embedding source column and its associated embedding model
// endpoint.
type EmbeddingSourceColumn struct {
	// Name of the embedding model endpoint, used by default for both ingestion
	// and querying.
	EmbeddingModelEndpoint string `json:"embedding_model_endpoint,omitempty"`
	// Name of the embedding model endpoint which, if specified, is used for
	// querying (not ingestion).
	ModelEndpointNameForQuery string `json:"model_endpoint_name_for_query,omitempty"`
	// Name of the column whose text is embedded.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EmbeddingSourceColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingSourceColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Name and dimension of an embedding vector column.
type EmbeddingVectorColumn struct {
	// Number of elements in each embedding vector. It must be positive.
	EmbeddingDimension int `json:"embedding_dimension,omitempty"`
	// Name of the column that holds the precomputed embedding vectors.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EmbeddingVectorColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingVectorColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An AI Search endpoint — compute infrastructure that hosts AI Search indexes
// and serves queries against them. Customers create, query, and delete
// endpoints; the system manages provisioning, scaling, and health status.
type Endpoint struct {
	// ID of the serverless budget policy to apply to the endpoint. The caller
	// must have permission to use the policy. Set either this field or
	// `usage_policy_id`, not both. Responses do not echo this field; they
	// report the applied policy in `effective_budget_policy_id`.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Time the endpoint was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// User name of the principal that created the endpoint. For a user, it is
	// the user's email address; for any other principal, it is the principal's
	// unique name.
	Creator string `json:"creator,omitempty"`
	// Custom tags assigned to the endpoint. On update, the list replaces all of
	// the endpoint's existing tags.
	CustomTags []CustomTag `json:"custom_tags,omitempty"`
	// ID of the budget policy that is currently applied to the endpoint.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// Current lifecycle state of the endpoint, with an optional message that
	// explains it.
	EndpointStatus *EndpointStatus `json:"endpoint_status,omitempty"`
	// Type of endpoint. Required on create and immutable thereafter.
	EndpointType EndpointType `json:"endpoint_type"`
	// System-generated UUID that uniquely identifies the endpoint.
	Id string `json:"id,omitempty"`
	// Number of indexes on the endpoint
	IndexCount int `json:"index_count,omitempty"`
	// User name of the principal that last updated the endpoint, in the same
	// format as `creator`.
	LastUpdatedUser string `json:"last_updated_user,omitempty"`
	// Full resource name of the endpoint, in the form
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`. On create, the
	// server composes it from `parent` and `endpoint_id` and ignores any value
	// supplied in the request body.
	Name string `json:"name,omitempty"`
	// The client-supplied desired number of replicas for the endpoint, applied
	// at create/update time. Mutually exclusive with `target_qps`.
	ReplicaCount int `json:"replica_count,omitempty"`
	// Scaling information for the endpoint
	ScalingInfo *EndpointScalingInfo `json:"scaling_info,omitempty"`
	// Target QPS for the endpoint. Mutually exclusive with `replica_count`.
	// Best-effort; the system does not guarantee this QPS will be achieved.
	// Responses report the requested value in
	// `scaling_info.requested_target_qps`.
	TargetQps int `json:"target_qps,omitempty"`
	// Throughput information for the endpoint
	ThroughputInfo *EndpointThroughputInfo `json:"throughput_info,omitempty"`
	// Time the endpoint was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// ID of the usage policy to apply to the endpoint. The server treats it as
	// the same setting as `budget_policy_id`, so set only one of the two.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Endpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Endpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Scaling information for a Standard endpoint: the current scaling state, the
// requested QPS target, and the progress of an in-progress scaling change.
type EndpointScalingInfo struct {
	// Reason the in-progress scaling change is delayed. When set,
	// `estimated_completion_time` is unset. Unset when the change is not
	// delayed or no scaling change is in progress.
	DelayedReason EndpointScalingInfoDelayReason `json:"delayed_reason,omitempty"`
	// Estimated time at which the in-progress scaling change will complete. The
	// estimate may change as the scaling change progresses. Unset when no
	// scaling change is in progress, when no estimate is available, or when the
	// change is delayed (see `delayed_reason`).
	EstimatedCompletionTime *time.Time `json:"estimated_completion_time,omitempty"`
	// Percentage of index initialization completed, from 0 to 100. Set only
	// during the `INITIALIZING_INDEX` phase; unset in other phases or when
	// progress is not available.
	IndexInitializationPercentage int `json:"index_initialization_percentage,omitempty"`
	// Current phase of the in-progress scaling change. Unset when no scaling
	// change is in progress or the phase is not available.
	Phase EndpointScalingInfoPhase `json:"phase,omitempty"`
	// The requested QPS target for the endpoint. Best-effort; the system does
	// not guarantee this QPS will be achieved.
	RequestedTargetQps int64 `json:"requested_target_qps,omitempty"`
	// The current state of the scaling change request.
	State ScalingChangeState `json:"state,omitempty"`
	// Earliest time a new Target QPS change is accepted. Set only while a
	// scaling change is in progress, the change lock is on, and the endpoint
	// has a recorded last throughput change; may be in the past once the lock
	// has expired.
	TargetQpsChangeAllowedAfterTime *time.Time `json:"target_qps_change_allowed_after_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointScalingInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointScalingInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Reason an in-progress scaling change is delayed.
type EndpointScalingInfoDelayReason string

const EndpointScalingInfoDelayReasonCapacityUnavailable EndpointScalingInfoDelayReason = `CAPACITY_UNAVAILABLE`

const EndpointScalingInfoDelayReasonScaleDownBlocked EndpointScalingInfoDelayReason = `SCALE_DOWN_BLOCKED`

const EndpointScalingInfoDelayReasonTakingLongerThanExpected EndpointScalingInfoDelayReason = `TAKING_LONGER_THAN_EXPECTED`

// String representation for [fmt.Print]
func (f *EndpointScalingInfoDelayReason) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointScalingInfoDelayReason) Set(v string) error {
	switch v {
	case `CAPACITY_UNAVAILABLE`, `SCALE_DOWN_BLOCKED`, `TAKING_LONGER_THAN_EXPECTED`:
		*f = EndpointScalingInfoDelayReason(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAPACITY_UNAVAILABLE", "SCALE_DOWN_BLOCKED", "TAKING_LONGER_THAN_EXPECTED"`, v)
	}
}

// Values returns all possible values for EndpointScalingInfoDelayReason.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointScalingInfoDelayReason) Values() []EndpointScalingInfoDelayReason {
	return []EndpointScalingInfoDelayReason{
		EndpointScalingInfoDelayReasonCapacityUnavailable,
		EndpointScalingInfoDelayReasonScaleDownBlocked,
		EndpointScalingInfoDelayReasonTakingLongerThanExpected,
	}
}

// Type always returns EndpointScalingInfoDelayReason to satisfy [pflag.Value] interface
func (f *EndpointScalingInfoDelayReason) Type() string {
	return "EndpointScalingInfoDelayReason"
}

// Phase of an in-progress scaling change.
type EndpointScalingInfoPhase string

const EndpointScalingInfoPhaseAdjustingCapacity EndpointScalingInfoPhase = `ADJUSTING_CAPACITY`

const EndpointScalingInfoPhaseEstimating EndpointScalingInfoPhase = `ESTIMATING`

const EndpointScalingInfoPhaseFinalizing EndpointScalingInfoPhase = `FINALIZING`

const EndpointScalingInfoPhaseInitializingIndex EndpointScalingInfoPhase = `INITIALIZING_INDEX`

// String representation for [fmt.Print]
func (f *EndpointScalingInfoPhase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointScalingInfoPhase) Set(v string) error {
	switch v {
	case `ADJUSTING_CAPACITY`, `ESTIMATING`, `FINALIZING`, `INITIALIZING_INDEX`:
		*f = EndpointScalingInfoPhase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADJUSTING_CAPACITY", "ESTIMATING", "FINALIZING", "INITIALIZING_INDEX"`, v)
	}
}

// Values returns all possible values for EndpointScalingInfoPhase.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointScalingInfoPhase) Values() []EndpointScalingInfoPhase {
	return []EndpointScalingInfoPhase{
		EndpointScalingInfoPhaseAdjustingCapacity,
		EndpointScalingInfoPhaseEstimating,
		EndpointScalingInfoPhaseFinalizing,
		EndpointScalingInfoPhaseInitializingIndex,
	}
}

// Type always returns EndpointScalingInfoPhase to satisfy [pflag.Value] interface
func (f *EndpointScalingInfoPhase) Type() string {
	return "EndpointScalingInfoPhase"
}

// Lifecycle and health state of an AI Search endpoint, along with any
// human-readable detail about that state.
type EndpointStatus struct {
	// Human-readable detail about the endpoint's current state or the reason
	// for a state transition.
	Message string `json:"message,omitempty"`
	// Current lifecycle state of the endpoint. See `State` for the meaning of
	// each value.
	State EndpointStatusState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Lifecycle state of an AI Search endpoint, used by both Standard and Storage
// Optimized SKUs.
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

// Throughput information for an AI Search endpoint, including requested and
// current concurrency settings.
type EndpointThroughputInfo struct {
	// Additional information about the throughput change request
	ChangeRequestMessage string `json:"change_request_message,omitempty"`
	// The state of the most recent throughput change request
	ChangeRequestState ThroughputChangeRequestState `json:"change_request_state,omitempty"`
	// The current concurrency (total CPU) allocated to the endpoint
	CurrentConcurrency float64 `json:"current_concurrency,omitempty"`
	// The current utilization of concurrency as a percentage (0-100)
	CurrentConcurrencyUtilizationPercentage float64 `json:"current_concurrency_utilization_percentage,omitempty"`
	// The current number of replicas allocated to the endpoint
	CurrentNumReplicas int `json:"current_num_replicas,omitempty"`
	// The maximum concurrency allowed for this endpoint
	MaximumConcurrencyAllowed float64 `json:"maximum_concurrency_allowed,omitempty"`
	// The minimum concurrency allowed for this endpoint
	MinimalConcurrencyAllowed float64 `json:"minimal_concurrency_allowed,omitempty"`
	// The requested concurrency (total CPU) for the endpoint
	RequestedConcurrency float64 `json:"requested_concurrency,omitempty"`
	// The requested number of replicas for the endpoint
	RequestedNumReplicas int `json:"requested_num_replicas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointThroughputInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointThroughputInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	// Facet rows; each row is `[facet_column_name, value_or_range, count]`.
	FacetArray [][]json.RawMessage `json:"facet_array,omitempty"`
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

type GetEndpointRequest struct {
	// Full resource name of the endpoint. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`
	Name string `json:"-" url:"-"`
}

func (s *GetEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetIndexRequest struct {
	// Full resource name of the index. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name string `json:"-" url:"-"`
}

func (s *GetIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An AI Search index — a searchable collection of vectors and metadata hosted
// on an AI Search endpoint. Indexes are children of endpoints; customers
// create, get, list, and delete them. The `{index}` segment of the resource
// name is the index's Unity Catalog table name.
type Index struct {
	// Creator of the index.
	Creator string `json:"creator,omitempty"`
	// Specification for a Delta Sync index. Set when `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec *DeltaSyncIndexSpec `json:"delta_sync_index_spec,omitempty"`
	// Specification for a Direct Access index. Set when `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec *DirectAccessIndexSpec `json:"direct_access_index_spec,omitempty"`
	// Name of the endpoint associated with the index. Ignored on create — the
	// endpoint is taken from `CreateIndexRequest.parent`; populated only on
	// output.
	Endpoint string `json:"endpoint,omitempty"`
	// Subtype of the index. Use `HYBRID` for vector and hybrid search, or
	// `FULL_TEXT` for keyword-only search without embeddings; `VECTOR` is not a
	// supported create value. Set on create and immutable thereafter.
	IndexSubtype IndexSubtype `json:"index_subtype,omitempty"`
	// Type of index. A `DELTA_SYNC` index is kept in sync with a source Delta
	// table; a `DIRECT_ACCESS` index holds rows that you write through
	// UpsertData and RemoveData. Set on create and immutable thereafter.
	IndexType IndexType `json:"index_type"`
	// Name of the AI Search index. Server-assigned full resource path
	// (`workspaces/{workspace}/endpoints/{endpoint}/indexes/{index}`) on
	// output, where `{index}` is the index's Unity Catalog table name. On
	// create, the user-supplied UC table name is conveyed via
	// `CreateIndexRequest.index_id`; the server composes the full `name` and
	// returns it on the response.
	Name string `json:"name,omitempty"`
	// Name of the column that uniquely identifies each row. It must exist in
	// the source table (Delta Sync) or in `schema_json` (Direct Access) and
	// cannot be the embedding source or vector column. Set on create and
	// immutable thereafter.
	PrimaryKey string `json:"primary_key"`
	// Current status of the index.
	Status *IndexStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Index) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Index) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Lifecycle and health state of an AI Search index, along with human-readable
// detail about that state and basic indexing progress.
type IndexStatus struct {
	// Lifecycle and ingestion state. An online index can still have a failed
	// update.
	DetailedState IndexStatusState `json:"detailed_state,omitempty"`
	// Index API URL used to perform operations on the index.
	IndexUrl string `json:"index_url,omitempty"`
	// Number of rows indexed.
	IndexedRowCount int64 `json:"indexed_row_count,omitempty"`
	// Human-readable detail about the index's current state.
	Message string `json:"message,omitempty"`
	// Whether the index can serve queries. After CreateIndex, poll GetIndex
	// until true; message describes progress. This does not guarantee freshness
	// or update success; check detailed_state for update failures.
	Ready bool `json:"ready,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IndexStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IndexStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Detailed lifecycle state of the index.
type IndexStatusState string

const IndexStatusStateDownloading IndexStatusState = `DOWNLOADING`

const IndexStatusStateOffline IndexStatusState = `OFFLINE`

const IndexStatusStateOfflineFailed IndexStatusState = `OFFLINE_FAILED`

const IndexStatusStateOnline IndexStatusState = `ONLINE`

const IndexStatusStateOnlineContinuousUpdate IndexStatusState = `ONLINE_CONTINUOUS_UPDATE`

const IndexStatusStateOnlineDirectAccess IndexStatusState = `ONLINE_DIRECT_ACCESS`

const IndexStatusStateOnlineNoPendingUpdate IndexStatusState = `ONLINE_NO_PENDING_UPDATE`

const IndexStatusStateOnlinePipelineFailed IndexStatusState = `ONLINE_PIPELINE_FAILED`

const IndexStatusStateOnlineTriggeredUpdate IndexStatusState = `ONLINE_TRIGGERED_UPDATE`

const IndexStatusStateOnlineUpdatingPipelineResources IndexStatusState = `ONLINE_UPDATING_PIPELINE_RESOURCES`

const IndexStatusStateProvisioning IndexStatusState = `PROVISIONING`

const IndexStatusStateProvisioningEndpoint IndexStatusState = `PROVISIONING_ENDPOINT`

const IndexStatusStateProvisioningIndex IndexStatusState = `PROVISIONING_INDEX`

const IndexStatusStateProvisioningInitialSnapshot IndexStatusState = `PROVISIONING_INITIAL_SNAPSHOT`

const IndexStatusStateProvisioningPipelineResources IndexStatusState = `PROVISIONING_PIPELINE_RESOURCES`

// String representation for [fmt.Print]
func (f *IndexStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IndexStatusState) Set(v string) error {
	switch v {
	case `DOWNLOADING`, `OFFLINE`, `OFFLINE_FAILED`, `ONLINE`, `ONLINE_CONTINUOUS_UPDATE`, `ONLINE_DIRECT_ACCESS`, `ONLINE_NO_PENDING_UPDATE`, `ONLINE_PIPELINE_FAILED`, `ONLINE_TRIGGERED_UPDATE`, `ONLINE_UPDATING_PIPELINE_RESOURCES`, `PROVISIONING`, `PROVISIONING_ENDPOINT`, `PROVISIONING_INDEX`, `PROVISIONING_INITIAL_SNAPSHOT`, `PROVISIONING_PIPELINE_RESOURCES`:
		*f = IndexStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DOWNLOADING", "OFFLINE", "OFFLINE_FAILED", "ONLINE", "ONLINE_CONTINUOUS_UPDATE", "ONLINE_DIRECT_ACCESS", "ONLINE_NO_PENDING_UPDATE", "ONLINE_PIPELINE_FAILED", "ONLINE_TRIGGERED_UPDATE", "ONLINE_UPDATING_PIPELINE_RESOURCES", "PROVISIONING", "PROVISIONING_ENDPOINT", "PROVISIONING_INDEX", "PROVISIONING_INITIAL_SNAPSHOT", "PROVISIONING_PIPELINE_RESOURCES"`, v)
	}
}

// Values returns all possible values for IndexStatusState.
//
// There is no guarantee on the order of the values in the slice.
func (f *IndexStatusState) Values() []IndexStatusState {
	return []IndexStatusState{
		IndexStatusStateDownloading,
		IndexStatusStateOffline,
		IndexStatusStateOfflineFailed,
		IndexStatusStateOnline,
		IndexStatusStateOnlineContinuousUpdate,
		IndexStatusStateOnlineDirectAccess,
		IndexStatusStateOnlineNoPendingUpdate,
		IndexStatusStateOnlinePipelineFailed,
		IndexStatusStateOnlineTriggeredUpdate,
		IndexStatusStateOnlineUpdatingPipelineResources,
		IndexStatusStateProvisioning,
		IndexStatusStateProvisioningEndpoint,
		IndexStatusStateProvisioningIndex,
		IndexStatusStateProvisioningInitialSnapshot,
		IndexStatusStateProvisioningPipelineResources,
	}
}

// Type always returns IndexStatusState to satisfy [pflag.Value] interface
func (f *IndexStatusState) Type() string {
	return "IndexStatusState"
}

// The subtype of the AI Search index, determining the indexing and retrieval
// strategy. - `VECTOR`: Not a supported create value — do not select it. Use
// `HYBRID` (vector + hybrid search) or `FULL_TEXT` (full-text only). It is the
// proto2 default (`= 0`) solely to mirror the legacy `index_v2.proto` enum
// value-for-value; it is not an offered index subtype. - `FULL_TEXT`: An index
// that uses full-text search without vector embeddings. - `HYBRID`: An index
// that uses vector embeddings for similarity search and hybrid search.
type IndexSubtype string

// An index that uses full-text search without vector embeddings.
const IndexSubtypeFullText IndexSubtype = `FULL_TEXT`

// An index that uses vector embeddings for similarity search and hybrid search.
const IndexSubtypeHybrid IndexSubtype = `HYBRID`

// Not a supported create value — do not select it. Use `HYBRID` (vector +
// hybrid
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

// There are 2 types of AI Search indexes: - `DELTA_SYNC`: An index that
// automatically syncs with a source Delta Table, automatically and
// incrementally updating the index as the underlying data in the Delta Table
// changes. - `DIRECT_ACCESS`: An index that supports direct read and write of
// vectors and metadata through our REST and SDK APIs. With this model, the user
// manages index updates.
type IndexType string

// An index that automatically syncs with a source Delta Table,
const IndexTypeDeltaSync IndexType = `DELTA_SYNC`

// An index that supports direct read and write of vectors and
const IndexTypeDirectAccess IndexType = `DIRECT_ACCESS`

// String representation for [fmt.Print]
func (f *IndexType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IndexType) Set(v string) error {
	switch v {
	case `DELTA_SYNC`, `DIRECT_ACCESS`:
		*f = IndexType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTA_SYNC", "DIRECT_ACCESS"`, v)
	}
}

// Values returns all possible values for IndexType.
//
// There is no guarantee on the order of the values in the slice.
func (f *IndexType) Values() []IndexType {
	return []IndexType{
		IndexTypeDeltaSync,
		IndexTypeDirectAccess,
	}
}

// Type always returns IndexType to satisfy [pflag.Value] interface
func (f *IndexType) Type() string {
	return "IndexType"
}

type ListEndpointsRequest struct {
	// Maximum number of endpoints to return in one page. The server applies its
	// own page size limit, so a page can hold fewer endpoints than `page_size`
	// even when more remain.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Workspace that owns this collection of endpoints. Format:
	// `workspaces/{workspace_id}`
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for ListEndpoints carrying the page of endpoints and an optional
// continuation token.
type ListEndpointsResponse struct {
	// The endpoints in the workspace.
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	// A token that can be used to get the next page of results. Empty when
	// there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListIndexesRequest struct {
	// Maximum number of indexes to return. When unset, every index on the
	// endpoint is returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The Endpoint that owns this collection of indexes. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListIndexesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListIndexesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for ListIndexes carrying the page of indexes and an optional
// continuation token.
type ListIndexesResponse struct {
	// The indexes on the endpoint. The field is named `indexes` (not the
	// irregular plural `indices`) to satisfy core::0132, which derives the
	// response field name from the ListIndexes method.
	// core::0158::response-plural-first-field independently computes the
	// resource plural as `indices` and is satisfied via a scoped field
	// exception below.
	Indexes []Index `json:"indexes,omitempty"`
	// A token that can be used to get the next page of results. Empty when
	// there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListIndexesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListIndexesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Pipeline execution mode for a Delta Sync index. Required on create for Delta
// Sync indexes; the legacy backend rejects an unset value with
// INVALID_PARAMETER_VALUE. - `TRIGGERED`: the pipeline stops after refreshing
// the source table once, using the data available when the update started. -
// `CONTINUOUS`: the pipeline processes new data as it arrives in the source
// table to keep the index fresh.
type PipelineType string

// the pipeline processes new data as it arrives in the source table to
const PipelineTypeContinuous PipelineType = `CONTINUOUS`

// the pipeline stops after refreshing the source table once, using the
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

// Request to query (search) an AI Search index. The legacy `num_results` count
// is exposed as `max_results`; v1 returns up to `max_results` rows in a single
// response (no cursor pagination — see the note on `max_results` below).
type QueryIndexRequest struct {
	// Columns to include in each result row. Only columns that are part of the
	// index can be returned.
	Columns []string `json:"columns"`
	// Columns whose values are sent to the reranker.
	ColumnsToRerank []string `json:"columns_to_rerank,omitempty"`
	// Facets to compute over the matched results (e.g. `"category TOP 5"`).
	Facets []string `json:"facets,omitempty"`
	// JSON object of filters to apply, serialized as a string. Each key is a
	// column name with an optional operator suffix such as `<`, `>=`, `NOT`, or
	// `OR`, and each value is the value to compare against. Only columns that
	// are part of the index can be filtered. See
	// https://docs.databricks.com/aws/en/ai-search/query-ai-search#filters for
	// the supported operators.
	FiltersJson string `json:"filters_json,omitempty"`
	// Maximum number of results to return. Defaults to 10. It corresponds to
	// `num_results` in the Vector Search API.
	MaxResults int `json:"max_results,omitempty"`
	// Full resource name of the index to query. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name string `json:"-" url:"-"`
	// Text columns to search for `query_text`. When empty, all text columns are
	// searched.
	QueryColumns []string `json:"query_columns,omitempty"`
	// Query text. Required for Delta Sync indexes that compute embeddings from
	// a model endpoint, and for `HYBRID` and `FULL_TEXT` queries.
	QueryText string `json:"query_text,omitempty"`
	// Query type: `ANN` for vector similarity search, `HYBRID` for combined
	// vector and keyword search, or `FULL_TEXT` for keyword search without
	// embeddings. Defaults to `ANN`.
	QueryType string `json:"query_type,omitempty"`
	// Query vector. Required for Direct Access indexes and Delta Sync indexes
	// with self-managed vectors.
	QueryVector []float64 `json:"query_vector,omitempty"`
	// If set, results are reranked before being returned. See
	// https://docs.databricks.com/aws/en/ai-search/query-ai-search#rerank for
	// details.
	Reranker *RerankerConfig `json:"reranker,omitempty"`
	// Score threshold for the approximate nearest-neighbor search; results that
	// score below it are not returned. Defaults to 0.0.
	ScoreThreshold float64 `json:"score_threshold,omitempty"`
	// Sort clauses, e.g. `["rating DESC", "price ASC"]`. Overrides relevance
	// ordering.
	SortColumns []string `json:"sort_columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for QueryIndex carrying the matched rows and their column metadata.
type QueryIndexResponse struct {
	// Facet aggregation rows, when facets were requested.
	FacetResult *FacetResultData `json:"facet_result,omitempty"`
	// Metadata describing the result columns.
	Manifest *ResultManifest `json:"manifest,omitempty"`
	// The matched result rows.
	Result *ResultData `json:"result,omitempty"`
}

func (s *QueryIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to remove rows by primary key from a Direct Access AI Search index.
// Named RemoveData (not DeleteData) so the linter does not classify it as a
// standard AIP-135 Delete method — it deletes rows within an index, not the
// index resource.
type RemoveDataRequest struct {
	// Full resource name of the index. Must be a Direct Access index. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name string `json:"-" url:"-"`
	// Primary keys of the rows to remove.
	PrimaryKeys []string `json:"primary_keys"`
}

func (s *RemoveDataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response for RemoveData.
type RemoveDataResponse struct {
	// Per-row outcome of the delete.
	Result *DataModificationResult `json:"result,omitempty"`
	// Overall status of the delete: `SUCCESS` when every row was removed,
	// `PARTIAL_SUCCESS` when some rows failed (see
	// `result.failed_primary_keys`), or `FAILURE` when the operation failed.
	Status DataModificationStatus `json:"status,omitempty"`
}

func (s *RemoveDataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Configuration for reranking query results with a reranker model.
type RerankerConfig struct {
	// Reranker identifier: "databricks_reranker" for the base model, or a Model
	// Serving endpoint name when `model_type` is MODEL_TYPE_FINETUNED.
	Model string `json:"model,omitempty"`
	// Discriminator for how `model` is interpreted.
	ModelType RerankerConfigModelType `json:"model_type,omitempty"`
	// Parameters controlling reranking.
	Parameters *RerankerConfigRerankerParameters `json:"parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RerankerConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RerankerConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// How the `model` field is interpreted.
type RerankerConfigModelType string

const RerankerConfigModelTypeModelTypeBase RerankerConfigModelType = `MODEL_TYPE_BASE`

const RerankerConfigModelTypeModelTypeFinetuned RerankerConfigModelType = `MODEL_TYPE_FINETUNED`

// String representation for [fmt.Print]
func (f *RerankerConfigModelType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RerankerConfigModelType) Set(v string) error {
	switch v {
	case `MODEL_TYPE_BASE`, `MODEL_TYPE_FINETUNED`:
		*f = RerankerConfigModelType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MODEL_TYPE_BASE", "MODEL_TYPE_FINETUNED"`, v)
	}
}

// Values returns all possible values for RerankerConfigModelType.
//
// There is no guarantee on the order of the values in the slice.
func (f *RerankerConfigModelType) Values() []RerankerConfigModelType {
	return []RerankerConfigModelType{
		RerankerConfigModelTypeModelTypeBase,
		RerankerConfigModelTypeModelTypeFinetuned,
	}
}

// Type always returns RerankerConfigModelType to satisfy [pflag.Value] interface
func (f *RerankerConfigModelType) Type() string {
	return "RerankerConfigModelType"
}

// Parameters controlling how the reranker processes results.
type RerankerConfigRerankerParameters struct {
	// Columns whose values are concatenated and sent to the reranker.
	ColumnsToRerank []string `json:"columns_to_rerank,omitempty"`
}

func (s *RerankerConfigRerankerParameters) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The rows of a query result set.
type ResultData struct {
	// Result rows; each row is a list of column values aligned with the
	// manifest columns.
	DataArray [][]json.RawMessage `json:"data_array,omitempty"`
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

// Metadata describing the columns of a query result set.
type ResultManifest struct {
	// Number of columns in the result set.
	ColumnCount int `json:"column_count,omitempty"`
	// Information about each column in the result set.
	Columns []ColumnInfo `json:"columns,omitempty"`
	// Number of columns in the facet result.
	FacetColumnCount int `json:"facet_column_count,omitempty"`
	// Information about each facet column.
	FacetColumns []ColumnInfo `json:"facet_columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResultManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// State of the most recent scaling change request for a Standard endpoint.
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

// Request to scan (paginate over) the rows of an AI Search index. Models the
// legacy `num_results` / `last_primary_key` cursor as AIP-158 `page_size` /
// `page_token`.
type ScanIndexRequest struct {
	// Full resource name of the index to scan. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name string `json:"-" url:"-"`
	// Maximum number of rows to return in this page. Defaults to 10.
	PageSize int `json:"page_size,omitempty"`
	// Page token from a previous response; if unset, scanning starts from the
	// beginning.
	PageToken string `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ScanIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ScanIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for ScanIndex carrying a page of rows and an optional continuation
// token.
type ScanIndexResponse struct {
	// Rows in this page, ordered by primary key. Each row is a JSON object that
	// maps column names to values.
	Data []map[string]json.RawMessage `json:"data,omitempty"`
	// Token for the next page; empty when the scan is exhausted.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ScanIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ScanIndexResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to synchronize a Delta Sync AI Search index with its source Delta
// table.
type SyncIndexRequest struct {
	// Full resource name of the index to synchronize. Must be a Delta Sync
	// index. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name string `json:"-" url:"-"`
}

func (s *SyncIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response for SyncIndex. Empty today; reserved so future sync metadata (e.g.
// an operation handle) can be added without breaking the wire contract.
type SyncIndexResponse struct {
}

func (s *SyncIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// State of the most recent throughput change request issued against a Storage
// Optimized endpoint. Surfaced on
// `EndpointThroughputInfo.change_request_state`.
type ThroughputChangeRequestState string

const ThroughputChangeRequestStateChangeAdjusted ThroughputChangeRequestState = `CHANGE_ADJUSTED`

const ThroughputChangeRequestStateChangeFailed ThroughputChangeRequestState = `CHANGE_FAILED`

const ThroughputChangeRequestStateChangeInProgress ThroughputChangeRequestState = `CHANGE_IN_PROGRESS`

const ThroughputChangeRequestStateChangeReachedMaximum ThroughputChangeRequestState = `CHANGE_REACHED_MAXIMUM`

const ThroughputChangeRequestStateChangeReachedMinimum ThroughputChangeRequestState = `CHANGE_REACHED_MINIMUM`

const ThroughputChangeRequestStateChangeSuccess ThroughputChangeRequestState = `CHANGE_SUCCESS`

// String representation for [fmt.Print]
func (f *ThroughputChangeRequestState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ThroughputChangeRequestState) Set(v string) error {
	switch v {
	case `CHANGE_ADJUSTED`, `CHANGE_FAILED`, `CHANGE_IN_PROGRESS`, `CHANGE_REACHED_MAXIMUM`, `CHANGE_REACHED_MINIMUM`, `CHANGE_SUCCESS`:
		*f = ThroughputChangeRequestState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CHANGE_ADJUSTED", "CHANGE_FAILED", "CHANGE_IN_PROGRESS", "CHANGE_REACHED_MAXIMUM", "CHANGE_REACHED_MINIMUM", "CHANGE_SUCCESS"`, v)
	}
}

// Values returns all possible values for ThroughputChangeRequestState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ThroughputChangeRequestState) Values() []ThroughputChangeRequestState {
	return []ThroughputChangeRequestState{
		ThroughputChangeRequestStateChangeAdjusted,
		ThroughputChangeRequestStateChangeFailed,
		ThroughputChangeRequestStateChangeInProgress,
		ThroughputChangeRequestStateChangeReachedMaximum,
		ThroughputChangeRequestStateChangeReachedMinimum,
		ThroughputChangeRequestStateChangeSuccess,
	}
}

// Type always returns ThroughputChangeRequestState to satisfy [pflag.Value] interface
func (f *ThroughputChangeRequestState) Type() string {
	return "ThroughputChangeRequestState"
}

type UpdateEndpointRequest struct {
	// The Endpoint resource to update. `endpoint.name` carries the full
	// resource path.
	Endpoint Endpoint `json:"endpoint"`
	// Full resource name of the endpoint, in the form
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}`. On create, the
	// server composes it from `parent` and `endpoint_id` and ignores any value
	// supplied in the request body.
	Name string `json:"-" url:"-"`
	// Comma-separated list of `Endpoint` fields to update. The updatable fields
	// are `budget_policy_id`, `usage_policy_id`, `custom_tags`,
	// `replica_count`, and `target_qps`. The server rejects nested paths and
	// unknown fields.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to upsert rows into a Direct Access AI Search index.
type UpsertDataRequest struct {
	// JSON array of the rows to insert or update, serialized as a string. Each
	// element is an object that maps column names to values and must include
	// the primary key column; vector columns are given as arrays of numbers. A
	// row whose primary key already exists is overwritten. The array must not
	// be empty.
	InputsJson string `json:"inputs_json"`
	// Full resource name of the index. Must be a Direct Access index. Format:
	// `workspaces/{workspace_id}/endpoints/{endpoint_id}/indexes/{index_id}`
	Name string `json:"-" url:"-"`
}

func (s *UpsertDataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response for UpsertData.
type UpsertDataResponse struct {
	// Per-row outcome of the upsert.
	Result *DataModificationResult `json:"result,omitempty"`
	// Overall status of the upsert: `SUCCESS` when every row was written,
	// `PARTIAL_SUCCESS` when some rows failed (see
	// `result.failed_primary_keys`), or `FAILURE` when the operation failed.
	Status DataModificationStatus `json:"status,omitempty"`
}

func (s *UpsertDataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
