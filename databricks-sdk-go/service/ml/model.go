// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/duration"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type Activity struct {
	ActivityType ActivityType `json:"activity_type,omitempty"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Source stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	FromStage string `json:"from_stage,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	SystemComment string `json:"system_comment,omitempty"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	ToStage string `json:"to_stage,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Activity) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Activity) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An action that a user (with sufficient permissions) could take on an activity
// or comment.
//
// For activities, valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request
//
// * `REJECT_TRANSITION_REQUEST`: Reject a transition request
//
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
//
// For comments, valid values are: * `EDIT_COMMENT`: Edit the comment
//
// * `DELETE_COMMENT`: Delete the comment
type ActivityAction string

// Approve a transition request
const ActivityActionApproveTransitionRequest ActivityAction = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const ActivityActionCancelTransitionRequest ActivityAction = `CANCEL_TRANSITION_REQUEST`

// Delete the comment
const ActivityActionDeleteComment ActivityAction = `DELETE_COMMENT`

// Edit the comment
const ActivityActionEditComment ActivityAction = `EDIT_COMMENT`

// Reject a transition request
const ActivityActionRejectTransitionRequest ActivityAction = `REJECT_TRANSITION_REQUEST`

// String representation for [fmt.Print]
func (f *ActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActivityAction) Set(v string) error {
	switch v {
	case `APPROVE_TRANSITION_REQUEST`, `CANCEL_TRANSITION_REQUEST`, `DELETE_COMMENT`, `EDIT_COMMENT`, `REJECT_TRANSITION_REQUEST`:
		*f = ActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVE_TRANSITION_REQUEST", "CANCEL_TRANSITION_REQUEST", "DELETE_COMMENT", "EDIT_COMMENT", "REJECT_TRANSITION_REQUEST"`, v)
	}
}

// Values returns all possible values for ActivityAction.
//
// There is no guarantee on the order of the values in the slice.
func (f *ActivityAction) Values() []ActivityAction {
	return []ActivityAction{
		ActivityActionApproveTransitionRequest,
		ActivityActionCancelTransitionRequest,
		ActivityActionDeleteComment,
		ActivityActionEditComment,
		ActivityActionRejectTransitionRequest,
	}
}

// Type always returns ActivityAction to satisfy [pflag.Value] interface
func (f *ActivityAction) Type() string {
	return "ActivityAction"
}

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition.
//
// * `REQUESTED_TRANSITION`: User requested the corresponding stage transition.
//
// * `CANCELLED_REQUEST`: User cancelled an existing transition request.
//
// * `APPROVED_REQUEST`: User approved the corresponding stage transition.
//
// * `REJECTED_REQUEST`: User rejected the coressponding stage transition.
//
// * `SYSTEM_TRANSITION`: For events performed as a side effect, such as
// archiving existing model versions in a stage.
type ActivityType string

// User applied the corresponding stage transition.
const ActivityTypeAppliedTransition ActivityType = `APPLIED_TRANSITION`

// User approved the corresponding stage transition.
const ActivityTypeApprovedRequest ActivityType = `APPROVED_REQUEST`

// User cancelled an existing transition request.
const ActivityTypeCancelledRequest ActivityType = `CANCELLED_REQUEST`

const ActivityTypeNewComment ActivityType = `NEW_COMMENT`

// User rejected the coressponding stage transition.
const ActivityTypeRejectedRequest ActivityType = `REJECTED_REQUEST`

// User requested the corresponding stage transition.
const ActivityTypeRequestedTransition ActivityType = `REQUESTED_TRANSITION`

// For events performed as a side effect, such as archiving existing model
// versions in a stage.
const ActivityTypeSystemTransition ActivityType = `SYSTEM_TRANSITION`

// String representation for [fmt.Print]
func (f *ActivityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActivityType) Set(v string) error {
	switch v {
	case `APPLIED_TRANSITION`, `APPROVED_REQUEST`, `CANCELLED_REQUEST`, `NEW_COMMENT`, `REJECTED_REQUEST`, `REQUESTED_TRANSITION`, `SYSTEM_TRANSITION`:
		*f = ActivityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPLIED_TRANSITION", "APPROVED_REQUEST", "CANCELLED_REQUEST", "NEW_COMMENT", "REJECTED_REQUEST", "REQUESTED_TRANSITION", "SYSTEM_TRANSITION"`, v)
	}
}

// Values returns all possible values for ActivityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ActivityType) Values() []ActivityType {
	return []ActivityType{
		ActivityTypeAppliedTransition,
		ActivityTypeApprovedRequest,
		ActivityTypeCancelledRequest,
		ActivityTypeNewComment,
		ActivityTypeRejectedRequest,
		ActivityTypeRequestedTransition,
		ActivityTypeSystemTransition,
	}
}

// Type always returns ActivityType to satisfy [pflag.Value] interface
func (f *ActivityType) Type() string {
	return "ActivityType"
}

// An aggregation function applied over a time window.
type AggregationFunction struct {
	ApproxCountDistinct *ApproxCountDistinctFunction `json:"approx_count_distinct,omitempty"`

	ApproxPercentile *ApproxPercentileFunction `json:"approx_percentile,omitempty"`

	Avg *AvgFunction `json:"avg,omitempty"`

	CountFunction *CountFunction `json:"count_function,omitempty"`

	First *FirstFunction `json:"first,omitempty"`

	FirstDistinct *FirstDistinctFunction `json:"first_distinct,omitempty"`

	FirstN *FirstNFunction `json:"first_n,omitempty"`

	Last *LastFunction `json:"last,omitempty"`

	LastDistinct *LastDistinctFunction `json:"last_distinct,omitempty"`

	LastN *LastNFunction `json:"last_n,omitempty"`

	Max *MaxFunction `json:"max,omitempty"`

	Min *MinFunction `json:"min,omitempty"`

	StddevPop *StddevPopFunction `json:"stddev_pop,omitempty"`

	StddevSamp *StddevSampFunction `json:"stddev_samp,omitempty"`

	Sum *SumFunction `json:"sum,omitempty"`
	// The time window over which the aggregation is computed.
	TimeWindow *TimeWindow `json:"time_window,omitempty"`

	VarPop *VarPopFunction `json:"var_pop,omitempty"`

	VarSamp *VarSampFunction `json:"var_samp,omitempty"`
}

func (s *AggregationFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to identify and approve a model version stage transition
// request.
type ApproveTransitionRequest struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions bool `json:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ApproveTransitionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ApproveTransitionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ApproveTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	Activity *Activity `json:"activity,omitempty"`
}

func (s *ApproveTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Computes the approximate count of distinct values.
type ApproxCountDistinctFunction struct {
	// The input column from which the approximate count of distinct values is
	// computed.
	Input string `json:"input"`
	// The error the user accepts, as a fraction of the count, between 0.02 (by
	// default) and 0.3. This is one standard error of the estimate, the same
	// meaning as Spark's relativeSD: 0.02 means the result is within ±2% of
	// the true count about 68% of the time (one standard deviation covers 68%
	// of the region). Feature Store picks the sketch size that gives this
	// relative error. If neither relative_error nor relative_sd is set, 0.02 is
	// default chosen at feature creation.
	RelativeError float64 `json:"relative_error,omitempty"`
	// The maximum relative standard deviation allowed (default defined by
	// Spark). Deprecated: use relative_error. When relative_error is set, reads
	// return the same value here for clients that don't have relative_error,
	// and a request may only send that value.
	RelativeSd float64 `json:"relative_sd,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ApproxCountDistinctFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ApproxCountDistinctFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Computes the approximate percentile of values.
type ApproxPercentileFunction struct {
	// The accuracy parameter (higher is more accurate but slower). Deprecated:
	// use relative_error. When relative_error is set, reads return ceil(1 /
	// relative_error) here for clients that don't have relative_error, and a
	// request may only send that value.
	Accuracy int64 `json:"accuracy,omitempty"`
	// The input column from which the approximate percentile is computed.
	Input string `json:"input"`
	// The percentile value to compute (between 0 and 1).
	Percentile float64 `json:"percentile"`
	// The error the user accepts, as a fraction of rank, between 0.01 (by
	// default) and 0.3. The result is within this rank error with 99%
	// confidence: at 0.01, the value returned for p95 is between the true p94
	// and p96 at least 99% of the time. Feature Store picks the sketch size
	// that gives this error. If neither relative_error nor accuracy is set,
	// 0.01 is default chosen at feature creation.
	RelativeError float64 `json:"relative_error,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ApproxPercentileFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ApproxPercentileFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AuthConfig struct {
	// Mutual-TLS authentication. See MtlsConfig.
	MtlsConfig *MtlsConfig `json:"mtls_config,omitempty"`
	// Name of the Unity Catalog service credential. This value will be set
	// under the option databricks.serviceCredential
	UcServiceCredentialName string `json:"uc_service_credential_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AuthConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AuthConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Computes the average of values.
type AvgFunction struct {
	// The input column from which the average is computed. For Kafka sources,
	// use dot-prefixed path notation (e.g., "value.amount"). For nested fields,
	// the leaf node name is used. Colon-prefixed notation (e.g.,
	// "value:amount") is supported for backwards compatibility but is
	// deprecated; migrate to dot notation.
	Input string `json:"input"`
}

func (s *AvgFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BackfillFeaturesRequest struct {
	// Time ranges of source data to backfill. Ranges must not overlap.
	BackfillRanges []BackfillRange `json:"backfill_ranges"`
	// The budget policy ID, in UUID format, used to attribute the serverless
	// compute cost of this backfill. If not specified, a default budget policy
	// may be applied.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Full names of the features to backfill. All features must share the same
	// interval between values and the same offset. The interval is the
	// slide_duration of a sliding window or the window_duration of a tumbling
	// window.
	FeatureFullNames []string `json:"feature_full_names"`
	// Idempotency token for the request.
	RequestId string `json:"request_id,omitempty"`
	// Custom tags to associate with this backfill. They are applied to the
	// backfill job and forwarded to the underlying compute as Databricks
	// resource tags, so backfill cost can be attributed in the billing system
	// tables. These tags apply only to the backfill compute; they are not
	// applied to the Unity Catalog Feature resources themselves, whose tags are
	// managed separately through the Unity Catalog tagging API. A maximum of 25
	// tags is supported; keys and values are subject to the same limitations as
	// Databricks resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BackfillFeaturesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BackfillFeaturesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Result of a completed backfill.
type BackfillFeaturesResponse struct {
}

func (s *BackfillFeaturesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Progress and configuration for a backfill.
type BackfillOperationMetadata struct {
	// The requested time ranges, each aligned outward to the boundaries at
	// which the features emit values. Ranges that overlap after alignment are
	// combined into one. The backfill stops at the most recent materialized
	// values, so a range may extend past the values that it recomputed.
	BackfillRanges []BackfillRange `json:"backfill_ranges,omitempty"`
	// Full names of the features targeted by the backfill.
	FeatureFullNames []string `json:"feature_full_names,omitempty"`
	// Current state of the backfill.
	State BackfillOperationMetadataState `json:"state,omitempty"`
}

func (s *BackfillOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Lifecycle state of a backfill.
type BackfillOperationMetadataState string

const BackfillOperationMetadataStateCancelled BackfillOperationMetadataState = `CANCELLED`

const BackfillOperationMetadataStateFailed BackfillOperationMetadataState = `FAILED`

const BackfillOperationMetadataStatePending BackfillOperationMetadataState = `PENDING`

const BackfillOperationMetadataStateRunning BackfillOperationMetadataState = `RUNNING`

const BackfillOperationMetadataStateSucceeded BackfillOperationMetadataState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *BackfillOperationMetadataState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *BackfillOperationMetadataState) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCEEDED`:
		*f = BackfillOperationMetadataState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "PENDING", "RUNNING", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for BackfillOperationMetadataState.
//
// There is no guarantee on the order of the values in the slice.
func (f *BackfillOperationMetadataState) Values() []BackfillOperationMetadataState {
	return []BackfillOperationMetadataState{
		BackfillOperationMetadataStateCancelled,
		BackfillOperationMetadataStateFailed,
		BackfillOperationMetadataStatePending,
		BackfillOperationMetadataStateRunning,
		BackfillOperationMetadataStateSucceeded,
	}
}

// Type always returns BackfillOperationMetadataState to satisfy [pflag.Value] interface
func (f *BackfillOperationMetadataState) Type() string {
	return "BackfillOperationMetadataState"
}

// A time range of source data for a backfill. The range is aligned outward to
// the boundaries at which the feature emits values: every slide_duration for a
// sliding window or every window_duration for a tumbling window, shifted by
// offset. Feature values up to one window_duration after the aligned end may
// change. The backfill stops at the most recent materialized values and leaves
// later values to materialization.
type BackfillRange struct {
	// End of the backfill range, exclusive. If unset, defaults to the most
	// recent materialized values.
	EndTime *time.Time `json:"end_time,omitempty"`
	// Start of the backfill range, inclusive. If unset, defaults to the
	// earliest source timestamp of the feature.
	StartTime *time.Time `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BackfillRange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BackfillRange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BackfillSource struct {
	// The full three-part name (catalog, schema, name) of the Delta table
	// containing the historical data to backfill.
	DeltaTableName string `json:"delta_table_name,omitempty"`
	// Deprecated: Use delta_table_name instead. Kept for backwards
	// compatibility. The Delta table source containing the historical data to
	// backfill. Only the delta table name is used for backfill, other fields
	// are ignored.
	DeltaTableSource *DeltaTableSource `json:"delta_table_source,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BackfillSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BackfillSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BatchCreateMaterializedFeaturesRequest struct {
	// The requests to create materialized features.
	Requests []CreateMaterializedFeatureRequest `json:"requests"`
}

func (s *BatchCreateMaterializedFeaturesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BatchCreateMaterializedFeaturesResponse struct {
	// The created materialized features with assigned IDs.
	MaterializedFeatures []MaterializedFeature `json:"materialized_features,omitempty"`
}

func (s *BatchCreateMaterializedFeaturesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The request message for `CancelOperation` method.
type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name string `json:"-" url:"-"`
}

func (s *CancelOperationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ColumnIdentifier struct {
	// String representation of the column name using dot-prefixed path
	// notation.
	VariantExprPath string `json:"variant_expr_path"`
}

func (s *ColumnIdentifier) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A ColumnSelection function, equivalent to the LAST() record of an entity over
// a lifetime window
type ColumnSelection struct {
	// Column name from source to select as the feature value.
	Column string `json:"column"`
}

func (s *ColumnSelection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An action that a user (with sufficient permissions) could take on an activity
// or comment.
//
// For activities, valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request
//
// * `REJECT_TRANSITION_REQUEST`: Reject a transition request
//
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
//
// For comments, valid values are: * `EDIT_COMMENT`: Edit the comment
//
// * `DELETE_COMMENT`: Delete the comment
type CommentActivityAction string

// Approve a transition request
const CommentActivityActionApproveTransitionRequest CommentActivityAction = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const CommentActivityActionCancelTransitionRequest CommentActivityAction = `CANCEL_TRANSITION_REQUEST`

// Delete the comment
const CommentActivityActionDeleteComment CommentActivityAction = `DELETE_COMMENT`

// Edit the comment
const CommentActivityActionEditComment CommentActivityAction = `EDIT_COMMENT`

// Reject a transition request
const CommentActivityActionRejectTransitionRequest CommentActivityAction = `REJECT_TRANSITION_REQUEST`

// String representation for [fmt.Print]
func (f *CommentActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CommentActivityAction) Set(v string) error {
	switch v {
	case `APPROVE_TRANSITION_REQUEST`, `CANCEL_TRANSITION_REQUEST`, `DELETE_COMMENT`, `EDIT_COMMENT`, `REJECT_TRANSITION_REQUEST`:
		*f = CommentActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVE_TRANSITION_REQUEST", "CANCEL_TRANSITION_REQUEST", "DELETE_COMMENT", "EDIT_COMMENT", "REJECT_TRANSITION_REQUEST"`, v)
	}
}

// Values returns all possible values for CommentActivityAction.
//
// There is no guarantee on the order of the values in the slice.
func (f *CommentActivityAction) Values() []CommentActivityAction {
	return []CommentActivityAction{
		CommentActivityActionApproveTransitionRequest,
		CommentActivityActionCancelTransitionRequest,
		CommentActivityActionDeleteComment,
		CommentActivityActionEditComment,
		CommentActivityActionRejectTransitionRequest,
	}
}

// Type always returns CommentActivityAction to satisfy [pflag.Value] interface
func (f *CommentActivityAction) Type() string {
	return "CommentActivityAction"
}

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type CommentObject struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []CommentActivityAction `json:"available_actions,omitempty"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CommentObject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CommentObject) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ContinuousWindow struct {
	// The offset of the continuous window (must be non-positive).
	Offset string `json:"offset,omitempty"`
	// The duration of the continuous window (must be positive).
	WindowDuration string `json:"window_duration"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ContinuousWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContinuousWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Computes the count of values.
type CountFunction struct {
	// The input column from which the count is computed. For Kafka sources, use
	// dot-prefixed path notation (e.g., "value.amount"). For nested fields, the
	// leaf node name is used. Colon-prefixed notation (e.g., "value:amount") is
	// supported for backwards compatibility but is deprecated; migrate to dot
	// notation.
	Input string `json:"input"`
}

func (s *CountFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to create a comment on a model version.
type CreateComment struct {
	// User-provided comment on the action.
	Comment string `json:"comment"`
	// Name of the model.
	Name string `json:"name"`
	// Version of the model.
	Version string `json:"version"`
}

func (s *CreateComment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateCommentResponse struct {
	// New comment object
	Comment *CommentObject `json:"comment,omitempty"`
}

func (s *CreateCommentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateExperiment struct {
	// URI of the location where the artifacts of the experiment's runs are
	// stored, such as a Unity Catalog volume path of the form
	// `dbfs:/Volumes/<catalog>/<schema>/<volume>/<path>`. The URI must have a
	// scheme other than `file` and must not point under the MLflow-managed
	// prefix `dbfs:/databricks/mlflow-tracking`. If omitted, artifacts are
	// stored in MLflow-managed storage. See
	// https://docs.databricks.com/aws/en/mlflow/experiments for the supported
	// locations.
	ArtifactLocation string `json:"artifact_location,omitempty"`
	// Name of the experiment, as an absolute path in the workspace. Experiments
	// cannot be created under `/Repos`.
	Name string `json:"name"`
	// Tags to set on the new experiment, as key-value pairs. A request can
	// include at most 100 tags.
	Tags []ExperimentTag `json:"tags,omitempty"`
	// The location where the experiment's traces are stored. When set, the
	// underlying storage is provisioned and the experiment's traces are routed
	// to it. When unset, traces are stored in the default MLflow backend. This
	// field cannot be updated after the experiment is created.
	TraceLocation *ExperimentTraceLocation `json:"trace_location,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateExperiment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExperiment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateExperimentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExperimentResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateFeatureRequest struct {
	// Feature to create.
	Feature Feature `json:"feature"`
}

func (s *CreateFeatureRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateFeatureSpecRequest struct {
	// The feature spec to create.
	FeatureSpec FeatureSpec `json:"feature_spec"`
}

func (s *CreateFeatureSpecRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateFeatureTagRequest struct {
	FeatureName string `json:"-" url:"-"`

	FeatureTag FeatureTag `json:"feature_tag"`

	TableName string `json:"-" url:"-"`
}

func (s *CreateFeatureTagRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateForecastingExperimentRequest struct {
	// The column in the training table used to customize weights for each time
	// series.
	CustomWeightsColumn string `json:"custom_weights_column,omitempty"`
	// The path in the workspace to store the created experiment.
	ExperimentPath string `json:"experiment_path,omitempty"`
	// The time interval between consecutive rows in the time series data.
	// Possible values include: '1 second', '1 minute', '5 minutes', '10
	// minutes', '15 minutes', '30 minutes', 'Hourly', 'Daily', 'Weekly',
	// 'Monthly', 'Quarterly', 'Yearly'.
	ForecastGranularity string `json:"forecast_granularity"`
	// The number of time steps into the future to make predictions, calculated
	// as a multiple of forecast_granularity. This value represents how far
	// ahead the model should forecast.
	ForecastHorizon int64 `json:"forecast_horizon"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store future feature data
	// for predictions.
	FutureFeatureDataPath string `json:"future_feature_data_path,omitempty"`
	// The region code(s) to automatically add holiday features. Currently
	// supports only one region.
	HolidayRegions []string `json:"holiday_regions,omitempty"`
	// Specifies the list of feature columns to include in model training. These
	// columns must exist in the training data and be of type string, numerical,
	// or boolean. If not specified, no additional features will be included.
	// Note: Certain columns are automatically handled: - Automatically
	// excluded: split_column, target_column, custom_weights_column. -
	// Automatically included: time_column.
	IncludeFeatures []string `json:"include_features,omitempty"`
	// The maximum duration for the experiment in minutes. The experiment stops
	// automatically if it exceeds this limit.
	MaxRuntime int64 `json:"max_runtime,omitempty"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store predictions.
	PredictionDataPath string `json:"prediction_data_path,omitempty"`
	// The evaluation metric used to optimize the forecasting model.
	PrimaryMetric string `json:"primary_metric,omitempty"`
	// The fully qualified path of a Unity Catalog model, formatted as
	// catalog_name.schema_name.model_name, used to store the best model.
	RegisterTo string `json:"register_to,omitempty"`
	// // The column in the training table used for custom data splits. Values
	// must be 'train', 'validate', or 'test'.
	SplitColumn string `json:"split_column,omitempty"`
	// The column in the input training table used as the prediction target for
	// model training. The values in this column are used as the ground truth
	// for model training.
	TargetColumn string `json:"target_column"`
	// The column in the input training table that represents each row's
	// timestamp.
	TimeColumn string `json:"time_column"`
	// The column in the training table used to group the dataset for predicting
	// individual time series.
	TimeseriesIdentifierColumns []string `json:"timeseries_identifier_columns,omitempty"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used as training data for the
	// forecasting model.
	TrainDataPath string `json:"train_data_path"`
	// List of frameworks to include for model tuning. Possible values are
	// 'Prophet', 'ARIMA', 'DeepAR'. An empty list includes all supported
	// frameworks.
	TrainingFrameworks []string `json:"training_frameworks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateForecastingExperimentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateForecastingExperimentRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateForecastingExperimentResponse struct {
	// The unique ID of the created forecasting experiment
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateForecastingExperimentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateForecastingExperimentResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateIssueRequest struct {
	// Output-only fields are ignored.
	Issue Issue `json:"issue"`
	// Optional client-chosen ID: iss- followed by 32 lowercase hex digits. A
	// retry with the same ID fails with ALREADY_EXISTS rather than making a
	// duplicate. The server generates one if unset.
	IssueId string `json:"-" url:"issue_id,omitempty"`
	// The experiment to create the Issue in, as `experiments/{experiment}`,
	// where `{experiment}` is the experiment ID. For example,
	// `experiments/1234567890`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateIssueRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateIssueRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateKafkaConfigRequest struct {
	KafkaConfig KafkaConfig `json:"kafka_config"`
}

func (s *CreateKafkaConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateLoggedModelRequest struct {
	// The ID of the experiment that owns the model.
	ExperimentId string `json:"experiment_id"`
	// The type of the model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	ModelType string `json:"model_type,omitempty"`
	// The name of the model (optional). If not specified one will be generated.
	Name string `json:"name,omitempty"`
	// Parameters attached to the model.
	Params []LoggedModelParameter `json:"params,omitempty"`
	// The ID of the run that created the model.
	SourceRunId string `json:"source_run_id,omitempty"`
	// Tags attached to the model.
	Tags []LoggedModelTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateLoggedModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateLoggedModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateLoggedModelResponse struct {
	// The newly created logged model.
	Model *LoggedModel `json:"model,omitempty"`
}

func (s *CreateLoggedModelResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateMaterializationSetRequest struct {
	// The materialization set to create. Set `materialization_set_id` to choose
	// the identifier, or omit it to have one assigned.
	//
	// Members with a cron schedule or table trigger must request the same
	// `paused` value. An omitted `paused` defaults to `false`. A set can be
	// created paused only if all members are batch materializations.
	MaterializationSet MaterializationSet `json:"materialization_set"`
	// Idempotency token for the request. A second request carrying a token that
	// has already been used returns the operation the first one created,
	// whatever state it reached, and creates no second set. The token is not
	// checked against the rest of the request, so reusing one for a different
	// set returns the original operation rather than an error. Omit it and each
	// request creates its own set.
	RequestId string `json:"-" url:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateMaterializationSetRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateMaterializationSetRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateMaterializedFeatureRequest struct {
	// The materialized feature to create.
	MaterializedFeature MaterializedFeature `json:"materialized_feature"`
}

func (s *CreateMaterializedFeatureRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateModelRequest struct {
	// Optional description for registered model.
	Description string `json:"description,omitempty"`
	// Register models under this name
	Name string `json:"name"`
	// Additional metadata for registered model.
	Tags []ModelTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateModelResponse struct {
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

func (s *CreateModelResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateModelVersionRequest struct {
	// Description of the model version.
	Description string `json:"description,omitempty"`
	// Name of the existing registered model to create the version under.
	Name string `json:"name"`
	// ID of the MLflow run that produced the model artifacts, if `source` comes
	// from a run in the MLflow tracking server. The ID is at most 32 characters
	// long.
	RunId string `json:"run_id,omitempty"`
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	RunLink string `json:"run_link,omitempty"`
	// URI of the model artifacts to register, such as the artifact path of an
	// MLflow run or a `models:/<model_name>/<version>` URI of an existing
	// `READY` model version. The registry copies the artifacts asynchronously,
	// so the new version starts in `PENDING_REGISTRATION` status.
	Source string `json:"source"`
	// Tags to set on the new model version.
	Tags []ModelVersionTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateModelVersionResponse struct {
	// The new model version, including the version number that the registry
	// assigned.
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

func (s *CreateModelVersionResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateOnlineStoreRequest struct {
	// Online store to create.
	OnlineStore OnlineStore `json:"online_store"`
}

func (s *CreateOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to create a registry webhook.
type CreateRegistryWebhook struct {
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events []RegistryWebhookEvent `json:"events"`
	// External HTTPS URL called on event trigger (by using a POST request).
	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"`
	// ID of the job that the webhook runs.
	JobSpec *JobSpec `json:"job_spec,omitempty"`
	// If model name is not specified, a registry-wide webhook is created that
	// listens for the specified events across all versions of all registered
	// models.
	ModelName string `json:"model_name,omitempty"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRegistryWebhook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRegistryWebhook) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRun struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// The name of the run.
	RunName string `json:"run_name,omitempty"`
	// Unix timestamp in milliseconds of when the run started.
	StartTime int64 `json:"start_time,omitempty"`
	// Additional metadata for run.
	Tags []RunTag `json:"tags,omitempty"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRunResponse struct {
	// The newly created run.
	Run *Run `json:"run,omitempty"`
}

func (s *CreateRunResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateStreamRequest struct {
	// The Stream to create.
	Stream Stream `json:"stream"`
}

func (s *CreateStreamRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to create a model version stage transition request.
type CreateTransitionRequest struct {
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateTransitionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTransitionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTransitionRequestResponse struct {
	// New activity generated for stage transition request.
	Request *TransitionRequest `json:"request,omitempty"`
}

func (s *CreateTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateWebhookResponse struct {
	Webhook *RegistryWebhook `json:"webhook,omitempty"`
}

func (s *CreateWebhookResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A cron-based schedule trigger for batch materialization.
type CronSchedule struct {
	// The cron expression defining the schedule (e.g., "0 0 * * *" for daily at
	// midnight). The schedule is interpreted in timezone_id (defaults to UTC).
	// Required when mode is MANUAL (or unset). Left empty when mode is DERIVED,
	// where the service computes it (aligned to UTC) from the features' window
	// timing and fills it in on the response.
	CronExpression string `json:"cron_expression,omitempty"`
	// How the schedule is determined. Defaults to MANUAL when unset.
	Mode CronScheduleMode `json:"mode,omitempty"`
	// Whether automatic materialization is paused. Defaults to `false`.
	// Supported only for materializations in a materialization set. Use the
	// `UpdateMaterializationSet` API to pause or resume the set.
	//
	// Pausing does not cancel running or queued runs. Runs can still be
	// triggered through the `TriggerMaterializationSet` API. Resuming does not
	// submit a run. Creating a paused materialization skips its initial run.
	Paused bool `json:"paused,omitempty"`
	// A Java timezone ID. The schedule is resolved with respect to this
	// timezone. Defaults to UTC when omitted. Can only be configured for MANUAL
	// schedules; DERIVED schedules are always aligned to UTC.
	TimezoneId string `json:"timezone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CronSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CronSchedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The way a materialization schedule is arrived at.
type CronScheduleMode string

const CronScheduleModeDerived CronScheduleMode = `DERIVED`

const CronScheduleModeManual CronScheduleMode = `MANUAL`

// String representation for [fmt.Print]
func (f *CronScheduleMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CronScheduleMode) Set(v string) error {
	switch v {
	case `DERIVED`, `MANUAL`:
		*f = CronScheduleMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DERIVED", "MANUAL"`, v)
	}
}

// Values returns all possible values for CronScheduleMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CronScheduleMode) Values() []CronScheduleMode {
	return []CronScheduleMode{
		CronScheduleModeDerived,
		CronScheduleModeManual,
	}
}

// Type always returns CronScheduleMode to satisfy [pflag.Value] interface
func (f *CronScheduleMode) Type() string {
	return "CronScheduleMode"
}

// A CustomUdf function applies a registered Unity Catalog function row-wise to
// source columns, producing a single output column per row.
type CustomUdf struct {
	// Fully qualified 3-part Unity Catalog path of the function to apply.
	FunctionPath string `json:"function_path"`
	// Binds each UC function parameter to a source column. May be empty for
	// zero-argument functions (e.g. a timestamp generator).
	InputBindings []InputBinding `json:"input_bindings,omitempty"`
}

func (s *CustomUdf) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Specifies the data source backing a feature. Exactly one source type must be
// set.
type DataSource struct {
	// A Delta table data source.
	DeltaTableSource *DeltaTableSource `json:"delta_table_source,omitempty"`
	// A data source composed from registered upstream Features.
	FeatureViewSource *FeatureViewSource `json:"feature_view_source,omitempty"`
	// A Unity Catalog managed or foreign Iceberg table source.
	IcebergTableSource *IcebergTableSource `json:"iceberg_table_source,omitempty"`
	// A Kafka stream data source.
	KafkaSource *KafkaSource `json:"kafka_source,omitempty"`
	// Completeness timing for this Feature's use of the source. This
	// configuration is part of the Feature definition; it does not modify the
	// underlying table or stream.
	Lateness *SourceLateness `json:"lateness,omitempty"`
	// A request-time data source.
	RequestSource *RequestSource `json:"request_source,omitempty"`
	// A Spark Declarative Pipeline data source.
	SdpSource *SdpSource `json:"sdp_source,omitempty"`
	// A Stream data source.
	StreamSource *StreamSource `json:"stream_source,omitempty"`
}

func (s *DataSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	Details []json.RawMessage `json:"details,omitempty"`

	ErrorCode ErrorCode `json:"error_code,omitempty"`

	Message string `json:"message,omitempty"`

	StackTrace string `json:"stack_trace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksServiceExceptionWithDetailsProto) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksServiceExceptionWithDetailsProto) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Dataset. Represents a reference to data used for training, testing, or
// evaluation during the model development process.
type Dataset struct {
	// Dataset digest, e.g. an md5 hash of the dataset that uniquely identifies
	// it within datasets of the same name.
	Digest string `json:"digest"`
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	Name string `json:"name"`
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	Profile string `json:"profile,omitempty"`
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	Schema string `json:"schema,omitempty"`
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	Source string `json:"source"`
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	SourceType string `json:"source_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Dataset) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Dataset) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// DatasetInput. Represents a dataset and input tags.
type DatasetInput struct {
	// The dataset being used as a Run input.
	Dataset Dataset `json:"dataset"`
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	Tags []InputTag `json:"tags,omitempty"`
}

func (s *DatasetInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A reference to a Feature.
type DeclarativeFeature struct {
	// Three-part Unity Catalog name (catalog.schema.feature) of the Feature.
	Feature string `json:"feature"`
}

func (s *DeclarativeFeature) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A default value for a looked-up feature.
type DefaultValue struct {
	// Default value for the feature.
	DefaultValue string `json:"default_value,omitempty"`
	// Name of the feature: its output name if it is renamed, otherwise its name
	// in the feature table.
	Feature string `json:"feature"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DefaultValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DefaultValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteCommentRequest struct {
	// Unique identifier of an activity
	Id string `json:"-" url:"id"`
}

func (s *DeleteCommentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
}

func (s *DeleteExperiment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteFeatureRequest struct {
	// Name of the feature to delete.
	FullName string `json:"-" url:"-"`
}

func (s *DeleteFeatureRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteFeatureSpecRequest struct {
	// Three-part name (catalog.schema.feature_spec) of the feature spec to
	// delete.
	Name string `json:"-" url:"-"`
}

func (s *DeleteFeatureSpecRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteFeatureTagRequest struct {
	// The name of the feature within the feature table.
	FeatureName string `json:"-" url:"-"`
	// The key of the tag to delete.
	Key string `json:"-" url:"-"`
	// The name of the feature table.
	TableName string `json:"-" url:"-"`
}

func (s *DeleteFeatureTagRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteKafkaConfigRequest struct {
	// Name of the Kafka config to delete.
	Name string `json:"-" url:"-"`
}

func (s *DeleteKafkaConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteLoggedModelRequest struct {
	// The ID of the logged model to delete.
	ModelId string `json:"-" url:"-"`
}

func (s *DeleteLoggedModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteLoggedModelTagRequest struct {
	// The ID of the logged model to delete the tag from.
	ModelId string `json:"-" url:"-"`
	// The tag key.
	TagKey string `json:"-" url:"-"`
}

func (s *DeleteLoggedModelTagRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteMaterializationSetRequest struct {
	// Return success instead of NOT_FOUND when the set does not exist, so a
	// repeated destroy is idempotent.
	AllowMissing bool `json:"-" url:"allow_missing,omitempty"`
	// The ID of the materialization set to delete.
	MaterializationSetId string `json:"-" url:"-"`
	// Idempotency token for the request. A second request carrying a token that
	// has already been used returns the operation the first one created,
	// whatever state it reached, and deletes nothing further. The token is not
	// checked against the rest of the request. Omit it and each request deletes
	// independently.
	RequestId string `json:"-" url:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteMaterializationSetRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteMaterializationSetRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteMaterializedFeatureRequest struct {
	// The ID of the materialized feature to delete.
	MaterializedFeatureId string `json:"-" url:"-"`
}

func (s *DeleteMaterializedFeatureRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"-" url:"name"`
}

func (s *DeleteModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key string `json:"-" url:"key"`
	// Name of the registered model that the tag was logged under.
	Name string `json:"-" url:"name"`
}

func (s *DeleteModelTagRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteModelVersionRequest struct {
	// Name of the registered model
	Name string `json:"-" url:"name"`
	// Model version number
	Version string `json:"-" url:"version"`
}

func (s *DeleteModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteModelVersionTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key string `json:"-" url:"key"`
	// Name of the registered model that the tag was logged under.
	Name string `json:"-" url:"name"`
	// Model version number that the tag was logged under.
	Version string `json:"-" url:"version"`
}

func (s *DeleteModelVersionTagRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteOnlineStoreRequest struct {
	// Name of the online store to delete.
	Name string `json:"-" url:"-"`
}

func (s *DeleteOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteOnlineTableRequest struct {
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName string `json:"-" url:"-"`
}

func (s *DeleteOnlineTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteRun struct {
	// ID of the run to delete.
	RunId string `json:"run_id"`
}

func (s *DeleteRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteRuns struct {
	// The ID of the experiment containing the runs to delete.
	ExperimentId string `json:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	MaxRuns int `json:"max_runs,omitempty"`
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	MaxTimestampMillis int64 `json:"max_timestamp_millis"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteRuns) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRuns) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteRunsResponse struct {
	// The number of runs deleted.
	RunsDeleted int `json:"runs_deleted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteStreamRequest struct {
	// Full three-part name (catalog.schema.stream) of the Stream to delete.
	Name string `json:"-" url:"-"`
}

func (s *DeleteStreamRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteTag struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	Key string `json:"key"`
	// ID of the run that the tag was logged under. Must be provided.
	RunId string `json:"run_id"`
}

func (s *DeleteTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteTransitionRequestRequest struct {
	// User-provided comment on the action.
	Comment string `json:"-" url:"comment,omitempty"`
	// Username of the user who created this request. Of the transition requests
	// matching the specified details, only the one transition created by this
	// user will be deleted.
	Creator string `json:"-" url:"creator"`
	// Name of the model.
	Name string `json:"-" url:"name"`
	// Target stage of the transition request. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"-" url:"stage"`
	// Version of the model.
	Version string `json:"-" url:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteTransitionRequestRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteTransitionRequestRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	Activity *Activity `json:"activity,omitempty"`
}

func (s *DeleteTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	Id string `json:"-" url:"id"`
}

func (s *DeleteWebhookRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeltaTableSource struct {
	// Schema of the resulting dataframe after transformations, in Spark
	// StructType JSON format (from df.schema.json()). Required if
	// transformation_sql is specified. Example:
	// {"type":"struct","fields":[{"name":"col_a","type":"integer","nullable":true,"metadata":{}},{"name":"col_c","type":"integer","nullable":true,"metadata":{}}]}
	DataframeSchema string `json:"dataframe_schema,omitempty"`

	EntityColumns []string `json:"entity_columns,omitempty"`
	// A single WHERE clause to filter the Delta table before applying
	// transformations. The predicate is evaluated for each row.
	FilterCondition string `json:"filter_condition,omitempty"`
	// The full three-part (catalog, schema, table) name of the Delta table.
	FullName string `json:"full_name"`

	TimeseriesColumn string `json:"timeseries_column,omitempty"`
	// A single SQL SELECT expression applied after filter_condition. The
	// expression must contain all required columns (e.g., "*, col_a + col_b AS
	// col_c"). If transformation_sql is not provided, all columns of the delta
	// table are present in the DataSource dataframe.
	TransformationSql string `json:"transformation_sql,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeltaTableSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaTableSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Direct connection configs for mTLS, as Kafka Connections do not support mTLS
// yet.
type DirectMtlsConfig struct {
	// A comma-separated list of host:port pairs for the Kafka bootstrap
	// servers.
	BootstrapServers string `json:"bootstrap_servers"`
	// Mutual-TLS authentication configuration.
	MtlsConfig MtlsConfig `json:"mtls_config"`
}

func (s *DirectMtlsConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Schema definitions provided directly on the Stream, as opposed to referencing
// a schema registry. To resolve schemas from a registry instead, use
// SchemaRegistryConfig.
type DirectSchemas struct {
	// Schema for the message key. This is only used for Kafka streams. For
	// Kafka, at least one of payload_schema or key_schema must be specified.
	KeySchema *SchemaConfig `json:"key_schema,omitempty"`
	// Schema for the message payload. For Kafka, this is the value schema.
	// Unless the platform supports another schema (e.g. keys for Kafka), this
	// must be specified.
	PayloadSchema *SchemaConfig `json:"payload_schema,omitempty"`
}

func (s *DirectSchemas) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Binds one entity column of a Lookup's upstream Feature to the input that
// supplies its values.
type EntityBinding struct {
	// Specifies the entity column of the upstream Feature.
	Entity string `json:"entity"`
	// Specifies where the column's values come from.
	Input InputReference `json:"input"`
}

func (s *EntityBinding) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EntityColumn struct {
	// The name of the entity column. For Kafka sources, use dot-prefixed path
	// notation to reference fields within the key or value schema (e.g.,
	// "value.user_id", "key.partition_key"). For nested fields, the leaf node
	// name (e.g., "user_id" from "value.trip_details.user_id") is what will be
	// present in materialized tables and expected to match at query time.
	// Colon-prefixed notation (e.g., "value:user_id") is supported for
	// backwards compatibility but is deprecated; migrate to dot notation.
	Name string `json:"name"`
}

func (s *EntityColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Error codes returned by Databricks APIs to indicate specific failure
// conditions.
type ErrorCode string

const ErrorCodeAborted ErrorCode = `ABORTED`

const ErrorCodeAlreadyExists ErrorCode = `ALREADY_EXISTS`

const ErrorCodeBadRequest ErrorCode = `BAD_REQUEST`

const ErrorCodeCancelled ErrorCode = `CANCELLED`

const ErrorCodeCatalogAlreadyExists ErrorCode = `CATALOG_ALREADY_EXISTS`

const ErrorCodeCatalogDoesNotExist ErrorCode = `CATALOG_DOES_NOT_EXIST`

const ErrorCodeCatalogNotEmpty ErrorCode = `CATALOG_NOT_EMPTY`

const ErrorCodeCouldNotAcquireLock ErrorCode = `COULD_NOT_ACQUIRE_LOCK`

const ErrorCodeCustomerUnauthorized ErrorCode = `CUSTOMER_UNAUTHORIZED`

const ErrorCodeDacAlreadyExists ErrorCode = `DAC_ALREADY_EXISTS`

const ErrorCodeDacDoesNotExist ErrorCode = `DAC_DOES_NOT_EXIST`

const ErrorCodeDataLoss ErrorCode = `DATA_LOSS`

const ErrorCodeDeadlineExceeded ErrorCode = `DEADLINE_EXCEEDED`

const ErrorCodeDeploymentTimeout ErrorCode = `DEPLOYMENT_TIMEOUT`

const ErrorCodeDirectoryNotEmpty ErrorCode = `DIRECTORY_NOT_EMPTY`

const ErrorCodeDirectoryProtected ErrorCode = `DIRECTORY_PROTECTED`

const ErrorCodeDryRunFailed ErrorCode = `DRY_RUN_FAILED`

const ErrorCodeEndpointNotFound ErrorCode = `ENDPOINT_NOT_FOUND`

const ErrorCodeExternalLocationAlreadyExists ErrorCode = `EXTERNAL_LOCATION_ALREADY_EXISTS`

const ErrorCodeExternalLocationDoesNotExist ErrorCode = `EXTERNAL_LOCATION_DOES_NOT_EXIST`

const ErrorCodeFeatureDisabled ErrorCode = `FEATURE_DISABLED`

const ErrorCodeGitConflict ErrorCode = `GIT_CONFLICT`

const ErrorCodeGitRemoteError ErrorCode = `GIT_REMOTE_ERROR`

const ErrorCodeGitSensitiveTokenDetected ErrorCode = `GIT_SENSITIVE_TOKEN_DETECTED`

const ErrorCodeGitUnknownRef ErrorCode = `GIT_UNKNOWN_REF`

const ErrorCodeGitUrlNotOnAllowList ErrorCode = `GIT_URL_NOT_ON_ALLOW_LIST`

const ErrorCodeInsecurePartnerResponse ErrorCode = `INSECURE_PARTNER_RESPONSE`

const ErrorCodeInternalError ErrorCode = `INTERNAL_ERROR`

const ErrorCodeInvalidParameterValue ErrorCode = `INVALID_PARAMETER_VALUE`

const ErrorCodeInvalidState ErrorCode = `INVALID_STATE`

const ErrorCodeInvalidStateTransition ErrorCode = `INVALID_STATE_TRANSITION`

const ErrorCodeIoError ErrorCode = `IO_ERROR`

const ErrorCodeIpynbFileInRepo ErrorCode = `IPYNB_FILE_IN_REPO`

const ErrorCodeMalformedPartnerResponse ErrorCode = `MALFORMED_PARTNER_RESPONSE`

const ErrorCodeMalformedRequest ErrorCode = `MALFORMED_REQUEST`

const ErrorCodeManagedResourceGroupDoesNotExist ErrorCode = `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`

const ErrorCodeMaxBlockSizeExceeded ErrorCode = `MAX_BLOCK_SIZE_EXCEEDED`

const ErrorCodeMaxChildNodeSizeExceeded ErrorCode = `MAX_CHILD_NODE_SIZE_EXCEEDED`

const ErrorCodeMaxListSizeExceeded ErrorCode = `MAX_LIST_SIZE_EXCEEDED`

const ErrorCodeMaxNotebookSizeExceeded ErrorCode = `MAX_NOTEBOOK_SIZE_EXCEEDED`

const ErrorCodeMaxReadSizeExceeded ErrorCode = `MAX_READ_SIZE_EXCEEDED`

const ErrorCodeMetastoreAlreadyExists ErrorCode = `METASTORE_ALREADY_EXISTS`

const ErrorCodeMetastoreDoesNotExist ErrorCode = `METASTORE_DOES_NOT_EXIST`

const ErrorCodeMetastoreNotEmpty ErrorCode = `METASTORE_NOT_EMPTY`

const ErrorCodeNotFound ErrorCode = `NOT_FOUND`

const ErrorCodeNotImplemented ErrorCode = `NOT_IMPLEMENTED`

const ErrorCodePartialDelete ErrorCode = `PARTIAL_DELETE`

const ErrorCodePermissionDenied ErrorCode = `PERMISSION_DENIED`

const ErrorCodePermissionNotPropagated ErrorCode = `PERMISSION_NOT_PROPAGATED`

const ErrorCodePrincipalDoesNotExist ErrorCode = `PRINCIPAL_DOES_NOT_EXIST`

const ErrorCodeProjectsOperationTimeout ErrorCode = `PROJECTS_OPERATION_TIMEOUT`

const ErrorCodeProviderAlreadyExists ErrorCode = `PROVIDER_ALREADY_EXISTS`

const ErrorCodeProviderDoesNotExist ErrorCode = `PROVIDER_DOES_NOT_EXIST`

const ErrorCodeProviderShareNotAccessible ErrorCode = `PROVIDER_SHARE_NOT_ACCESSIBLE`

const ErrorCodeQuotaExceeded ErrorCode = `QUOTA_EXCEEDED`

const ErrorCodeRecipientAlreadyExists ErrorCode = `RECIPIENT_ALREADY_EXISTS`

const ErrorCodeRecipientDoesNotExist ErrorCode = `RECIPIENT_DOES_NOT_EXIST`

const ErrorCodeRequestLimitExceeded ErrorCode = `REQUEST_LIMIT_EXCEEDED`

const ErrorCodeResourceAlreadyExists ErrorCode = `RESOURCE_ALREADY_EXISTS`

const ErrorCodeResourceConflict ErrorCode = `RESOURCE_CONFLICT`

const ErrorCodeResourceDoesNotExist ErrorCode = `RESOURCE_DOES_NOT_EXIST`

const ErrorCodeResourceExhausted ErrorCode = `RESOURCE_EXHAUSTED`

const ErrorCodeResourceLimitExceeded ErrorCode = `RESOURCE_LIMIT_EXCEEDED`

const ErrorCodeSchemaAlreadyExists ErrorCode = `SCHEMA_ALREADY_EXISTS`

const ErrorCodeSchemaDoesNotExist ErrorCode = `SCHEMA_DOES_NOT_EXIST`

const ErrorCodeSchemaNotEmpty ErrorCode = `SCHEMA_NOT_EMPTY`

const ErrorCodeSearchQueryTooLong ErrorCode = `SEARCH_QUERY_TOO_LONG`

const ErrorCodeSearchQueryTooShort ErrorCode = `SEARCH_QUERY_TOO_SHORT`

const ErrorCodeServiceUnderMaintenance ErrorCode = `SERVICE_UNDER_MAINTENANCE`

const ErrorCodeShareAlreadyExists ErrorCode = `SHARE_ALREADY_EXISTS`

const ErrorCodeShareDoesNotExist ErrorCode = `SHARE_DOES_NOT_EXIST`

const ErrorCodeStorageCredentialAlreadyExists ErrorCode = `STORAGE_CREDENTIAL_ALREADY_EXISTS`

const ErrorCodeStorageCredentialDoesNotExist ErrorCode = `STORAGE_CREDENTIAL_DOES_NOT_EXIST`

const ErrorCodeTableAlreadyExists ErrorCode = `TABLE_ALREADY_EXISTS`

const ErrorCodeTableDoesNotExist ErrorCode = `TABLE_DOES_NOT_EXIST`

const ErrorCodeTemporarilyUnavailable ErrorCode = `TEMPORARILY_UNAVAILABLE`

const ErrorCodeUnauthenticated ErrorCode = `UNAUTHENTICATED`

const ErrorCodeUnavailable ErrorCode = `UNAVAILABLE`

const ErrorCodeUnknown ErrorCode = `UNKNOWN`

const ErrorCodeUnparseableHttpError ErrorCode = `UNPARSEABLE_HTTP_ERROR`

const ErrorCodeWorkspaceTemporarilyUnavailable ErrorCode = `WORKSPACE_TEMPORARILY_UNAVAILABLE`

// String representation for [fmt.Print]
func (f *ErrorCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ErrorCode) Set(v string) error {
	switch v {
	case `ABORTED`, `ALREADY_EXISTS`, `BAD_REQUEST`, `CANCELLED`, `CATALOG_ALREADY_EXISTS`, `CATALOG_DOES_NOT_EXIST`, `CATALOG_NOT_EMPTY`, `COULD_NOT_ACQUIRE_LOCK`, `CUSTOMER_UNAUTHORIZED`, `DAC_ALREADY_EXISTS`, `DAC_DOES_NOT_EXIST`, `DATA_LOSS`, `DEADLINE_EXCEEDED`, `DEPLOYMENT_TIMEOUT`, `DIRECTORY_NOT_EMPTY`, `DIRECTORY_PROTECTED`, `DRY_RUN_FAILED`, `ENDPOINT_NOT_FOUND`, `EXTERNAL_LOCATION_ALREADY_EXISTS`, `EXTERNAL_LOCATION_DOES_NOT_EXIST`, `FEATURE_DISABLED`, `GIT_CONFLICT`, `GIT_REMOTE_ERROR`, `GIT_SENSITIVE_TOKEN_DETECTED`, `GIT_UNKNOWN_REF`, `GIT_URL_NOT_ON_ALLOW_LIST`, `INSECURE_PARTNER_RESPONSE`, `INTERNAL_ERROR`, `INVALID_PARAMETER_VALUE`, `INVALID_STATE`, `INVALID_STATE_TRANSITION`, `IO_ERROR`, `IPYNB_FILE_IN_REPO`, `MALFORMED_PARTNER_RESPONSE`, `MALFORMED_REQUEST`, `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`, `MAX_BLOCK_SIZE_EXCEEDED`, `MAX_CHILD_NODE_SIZE_EXCEEDED`, `MAX_LIST_SIZE_EXCEEDED`, `MAX_NOTEBOOK_SIZE_EXCEEDED`, `MAX_READ_SIZE_EXCEEDED`, `METASTORE_ALREADY_EXISTS`, `METASTORE_DOES_NOT_EXIST`, `METASTORE_NOT_EMPTY`, `NOT_FOUND`, `NOT_IMPLEMENTED`, `PARTIAL_DELETE`, `PERMISSION_DENIED`, `PERMISSION_NOT_PROPAGATED`, `PRINCIPAL_DOES_NOT_EXIST`, `PROJECTS_OPERATION_TIMEOUT`, `PROVIDER_ALREADY_EXISTS`, `PROVIDER_DOES_NOT_EXIST`, `PROVIDER_SHARE_NOT_ACCESSIBLE`, `QUOTA_EXCEEDED`, `RECIPIENT_ALREADY_EXISTS`, `RECIPIENT_DOES_NOT_EXIST`, `REQUEST_LIMIT_EXCEEDED`, `RESOURCE_ALREADY_EXISTS`, `RESOURCE_CONFLICT`, `RESOURCE_DOES_NOT_EXIST`, `RESOURCE_EXHAUSTED`, `RESOURCE_LIMIT_EXCEEDED`, `SCHEMA_ALREADY_EXISTS`, `SCHEMA_DOES_NOT_EXIST`, `SCHEMA_NOT_EMPTY`, `SEARCH_QUERY_TOO_LONG`, `SEARCH_QUERY_TOO_SHORT`, `SERVICE_UNDER_MAINTENANCE`, `SHARE_ALREADY_EXISTS`, `SHARE_DOES_NOT_EXIST`, `STORAGE_CREDENTIAL_ALREADY_EXISTS`, `STORAGE_CREDENTIAL_DOES_NOT_EXIST`, `TABLE_ALREADY_EXISTS`, `TABLE_DOES_NOT_EXIST`, `TEMPORARILY_UNAVAILABLE`, `UNAUTHENTICATED`, `UNAVAILABLE`, `UNKNOWN`, `UNPARSEABLE_HTTP_ERROR`, `WORKSPACE_TEMPORARILY_UNAVAILABLE`:
		*f = ErrorCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABORTED", "ALREADY_EXISTS", "BAD_REQUEST", "CANCELLED", "CATALOG_ALREADY_EXISTS", "CATALOG_DOES_NOT_EXIST", "CATALOG_NOT_EMPTY", "COULD_NOT_ACQUIRE_LOCK", "CUSTOMER_UNAUTHORIZED", "DAC_ALREADY_EXISTS", "DAC_DOES_NOT_EXIST", "DATA_LOSS", "DEADLINE_EXCEEDED", "DEPLOYMENT_TIMEOUT", "DIRECTORY_NOT_EMPTY", "DIRECTORY_PROTECTED", "DRY_RUN_FAILED", "ENDPOINT_NOT_FOUND", "EXTERNAL_LOCATION_ALREADY_EXISTS", "EXTERNAL_LOCATION_DOES_NOT_EXIST", "FEATURE_DISABLED", "GIT_CONFLICT", "GIT_REMOTE_ERROR", "GIT_SENSITIVE_TOKEN_DETECTED", "GIT_UNKNOWN_REF", "GIT_URL_NOT_ON_ALLOW_LIST", "INSECURE_PARTNER_RESPONSE", "INTERNAL_ERROR", "INVALID_PARAMETER_VALUE", "INVALID_STATE", "INVALID_STATE_TRANSITION", "IO_ERROR", "IPYNB_FILE_IN_REPO", "MALFORMED_PARTNER_RESPONSE", "MALFORMED_REQUEST", "MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST", "MAX_BLOCK_SIZE_EXCEEDED", "MAX_CHILD_NODE_SIZE_EXCEEDED", "MAX_LIST_SIZE_EXCEEDED", "MAX_NOTEBOOK_SIZE_EXCEEDED", "MAX_READ_SIZE_EXCEEDED", "METASTORE_ALREADY_EXISTS", "METASTORE_DOES_NOT_EXIST", "METASTORE_NOT_EMPTY", "NOT_FOUND", "NOT_IMPLEMENTED", "PARTIAL_DELETE", "PERMISSION_DENIED", "PERMISSION_NOT_PROPAGATED", "PRINCIPAL_DOES_NOT_EXIST", "PROJECTS_OPERATION_TIMEOUT", "PROVIDER_ALREADY_EXISTS", "PROVIDER_DOES_NOT_EXIST", "PROVIDER_SHARE_NOT_ACCESSIBLE", "QUOTA_EXCEEDED", "RECIPIENT_ALREADY_EXISTS", "RECIPIENT_DOES_NOT_EXIST", "REQUEST_LIMIT_EXCEEDED", "RESOURCE_ALREADY_EXISTS", "RESOURCE_CONFLICT", "RESOURCE_DOES_NOT_EXIST", "RESOURCE_EXHAUSTED", "RESOURCE_LIMIT_EXCEEDED", "SCHEMA_ALREADY_EXISTS", "SCHEMA_DOES_NOT_EXIST", "SCHEMA_NOT_EMPTY", "SEARCH_QUERY_TOO_LONG", "SEARCH_QUERY_TOO_SHORT", "SERVICE_UNDER_MAINTENANCE", "SHARE_ALREADY_EXISTS", "SHARE_DOES_NOT_EXIST", "STORAGE_CREDENTIAL_ALREADY_EXISTS", "STORAGE_CREDENTIAL_DOES_NOT_EXIST", "TABLE_ALREADY_EXISTS", "TABLE_DOES_NOT_EXIST", "TEMPORARILY_UNAVAILABLE", "UNAUTHENTICATED", "UNAVAILABLE", "UNKNOWN", "UNPARSEABLE_HTTP_ERROR", "WORKSPACE_TEMPORARILY_UNAVAILABLE"`, v)
	}
}

// Values returns all possible values for ErrorCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		ErrorCodeAborted,
		ErrorCodeAlreadyExists,
		ErrorCodeBadRequest,
		ErrorCodeCancelled,
		ErrorCodeCatalogAlreadyExists,
		ErrorCodeCatalogDoesNotExist,
		ErrorCodeCatalogNotEmpty,
		ErrorCodeCouldNotAcquireLock,
		ErrorCodeCustomerUnauthorized,
		ErrorCodeDacAlreadyExists,
		ErrorCodeDacDoesNotExist,
		ErrorCodeDataLoss,
		ErrorCodeDeadlineExceeded,
		ErrorCodeDeploymentTimeout,
		ErrorCodeDirectoryNotEmpty,
		ErrorCodeDirectoryProtected,
		ErrorCodeDryRunFailed,
		ErrorCodeEndpointNotFound,
		ErrorCodeExternalLocationAlreadyExists,
		ErrorCodeExternalLocationDoesNotExist,
		ErrorCodeFeatureDisabled,
		ErrorCodeGitConflict,
		ErrorCodeGitRemoteError,
		ErrorCodeGitSensitiveTokenDetected,
		ErrorCodeGitUnknownRef,
		ErrorCodeGitUrlNotOnAllowList,
		ErrorCodeInsecurePartnerResponse,
		ErrorCodeInternalError,
		ErrorCodeInvalidParameterValue,
		ErrorCodeInvalidState,
		ErrorCodeInvalidStateTransition,
		ErrorCodeIoError,
		ErrorCodeIpynbFileInRepo,
		ErrorCodeMalformedPartnerResponse,
		ErrorCodeMalformedRequest,
		ErrorCodeManagedResourceGroupDoesNotExist,
		ErrorCodeMaxBlockSizeExceeded,
		ErrorCodeMaxChildNodeSizeExceeded,
		ErrorCodeMaxListSizeExceeded,
		ErrorCodeMaxNotebookSizeExceeded,
		ErrorCodeMaxReadSizeExceeded,
		ErrorCodeMetastoreAlreadyExists,
		ErrorCodeMetastoreDoesNotExist,
		ErrorCodeMetastoreNotEmpty,
		ErrorCodeNotFound,
		ErrorCodeNotImplemented,
		ErrorCodePartialDelete,
		ErrorCodePermissionDenied,
		ErrorCodePermissionNotPropagated,
		ErrorCodePrincipalDoesNotExist,
		ErrorCodeProjectsOperationTimeout,
		ErrorCodeProviderAlreadyExists,
		ErrorCodeProviderDoesNotExist,
		ErrorCodeProviderShareNotAccessible,
		ErrorCodeQuotaExceeded,
		ErrorCodeRecipientAlreadyExists,
		ErrorCodeRecipientDoesNotExist,
		ErrorCodeRequestLimitExceeded,
		ErrorCodeResourceAlreadyExists,
		ErrorCodeResourceConflict,
		ErrorCodeResourceDoesNotExist,
		ErrorCodeResourceExhausted,
		ErrorCodeResourceLimitExceeded,
		ErrorCodeSchemaAlreadyExists,
		ErrorCodeSchemaDoesNotExist,
		ErrorCodeSchemaNotEmpty,
		ErrorCodeSearchQueryTooLong,
		ErrorCodeSearchQueryTooShort,
		ErrorCodeServiceUnderMaintenance,
		ErrorCodeShareAlreadyExists,
		ErrorCodeShareDoesNotExist,
		ErrorCodeStorageCredentialAlreadyExists,
		ErrorCodeStorageCredentialDoesNotExist,
		ErrorCodeTableAlreadyExists,
		ErrorCodeTableDoesNotExist,
		ErrorCodeTemporarilyUnavailable,
		ErrorCodeUnauthenticated,
		ErrorCodeUnavailable,
		ErrorCodeUnknown,
		ErrorCodeUnparseableHttpError,
		ErrorCodeWorkspaceTemporarilyUnavailable,
	}
}

// Type always returns ErrorCode to satisfy [pflag.Value] interface
func (f *ErrorCode) Type() string {
	return "ErrorCode"
}

// An experiment and its metadata.
type Experiment struct {
	// URI of the location where the artifacts of the experiment's runs are
	// stored. Experiments created without an explicit location use
	// MLflow-managed storage under
	// `dbfs:/databricks/mlflow-tracking/<experiment_id>`.
	ArtifactLocation string `json:"artifact_location,omitempty"`
	// Time when the experiment was created, as a Unix timestamp in
	// milliseconds.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Unique identifier for the experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Time when the experiment metadata was last updated, as a Unix timestamp
	// in milliseconds.
	LastUpdateTime int64 `json:"last_update_time,omitempty"`
	// Current life cycle stage of the experiment: `active` or `deleted`.
	// Deleted experiments are excluded from list and search results unless
	// `view_type` requests them, and can be restored until they are permanently
	// deleted.
	LifecycleStage string `json:"lifecycle_stage,omitempty"`
	// Name of the experiment. A workspace experiment is named by its absolute
	// path in the workspace; a notebook experiment has the path of its
	// notebook.
	Name string `json:"name,omitempty"`
	// Tags set on the experiment, as key-value pairs, including system tags
	// such as `mlflow.ownerId` that MLflow maintains.
	Tags []ExperimentTag `json:"tags,omitempty"`
	// The location where the experiment's traces are stored. Unset when traces
	// are stored in the default MLflow backend. This field cannot be updated
	// after the experiment is created.
	TraceLocation *ExperimentTraceLocation `json:"trace_location,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Experiment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Experiment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExperimentAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExperimentAccessControlResponse struct {
	// All permissions.
	AllPermissions []ExperimentPermission `json:"all_permissions,omitempty"`
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

func (s *ExperimentAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExperimentPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type ExperimentPermissionLevel string

const ExperimentPermissionLevelCanEdit ExperimentPermissionLevel = `CAN_EDIT`

const ExperimentPermissionLevelCanManage ExperimentPermissionLevel = `CAN_MANAGE`

const ExperimentPermissionLevelCanRead ExperimentPermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *ExperimentPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExperimentPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`:
		*f = ExperimentPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ"`, v)
	}
}

// Values returns all possible values for ExperimentPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExperimentPermissionLevel) Values() []ExperimentPermissionLevel {
	return []ExperimentPermissionLevel{
		ExperimentPermissionLevelCanEdit,
		ExperimentPermissionLevelCanManage,
		ExperimentPermissionLevelCanRead,
	}
}

// Type always returns ExperimentPermissionLevel to satisfy [pflag.Value] interface
func (f *ExperimentPermissionLevel) Type() string {
	return "ExperimentPermissionLevel"
}

type ExperimentPermissions struct {
	AccessControlList []ExperimentAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExperimentPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExperimentPermissionsRequest struct {
	AccessControlList []ExperimentAccessControlRequest `json:"access_control_list,omitempty"`
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" url:"-"`
}

func (s *ExperimentPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A tag for an experiment.
type ExperimentTag struct {
	// The tag key. Keys are at most 255 characters long.
	Key string `json:"key,omitempty"`
	// The tag value. Values are at most 65535 characters long.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExperimentTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExperimentTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The storage location for an experiment's traces.
type ExperimentTraceLocation struct {
	// A Unity Catalog schema where the experiment's traces are stored as Delta
	// tables.
	UcTraceLocation *UcTraceLocation `json:"uc_trace_location,omitempty"`
}

func (s *ExperimentTraceLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A trigger that disables automatic materialization, including the initial run
// at creation. Runs are instead triggered through the
// `TriggerMaterializationSet` API. Supported only for batch materializations in
// a materialization set.
type ExternalTrigger struct {
}

func (s *ExternalTrigger) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Feature struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// Time at which this feature was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Username of the feature creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The description of the feature.
	Description string `json:"description,omitempty"`
	// The entity columns for the feature, used as aggregation keys and for
	// query-time lookup. Optional since entities are not set for RequestSource
	// features or on-demand calculated features.
	Entities []EntityColumn `json:"entities,omitempty"`

	FilterCondition string `json:"filter_condition,omitempty"`
	// The full three-part name (catalog, schema, name) of the feature.
	FullName string `json:"full_name"`
	// The function by which the feature is computed.
	Function Function `json:"function"`

	Inputs []string `json:"inputs,omitempty"`
	// Lineage context information for this feature. WARNING: This field is
	// primarily intended for internal use by Databricks systems and is
	// automatically populated when features are created through Databricks
	// notebooks or jobs. Users should not manually set this field as incorrect
	// values may lead to inaccurate lineage tracking or unexpected behavior.
	// This field will be set by feature-engineering client and should be left
	// unset by SDK and terraform users.
	LineageContext *LineageContext `json:"lineage_context,omitempty"`
	// Name of the feature, extracted from the full three-part name
	// (catalog.schema.name).
	Name string `json:"name,omitempty"`
	// Output type in Databricks SQL DDL (e.g., "BIGINT", "DECIMAL(10,2)",
	// "ARRAY<STRING>", "STRUCT<name:STRING>").
	OutputTypeText string `json:"output_type_text,omitempty"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// The data source of the feature.
	Source DataSource `json:"source"`

	TimeWindow *TimeWindow `json:"time_window,omitempty"`
	// Column recording time, used for point-in-time joins, backfills, and
	// aggregations. Optional since a timeseries column is not set for
	// RequestSource features or on-demand calculated features.
	TimeseriesColumn *TimeseriesColumn `json:"timeseries_column,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Feature) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Feature) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Python UDF that computes a feature.
type FeatureFunction struct {
	// Bindings of the UDF's parameters to features. Order is not significant,
	// and GetFeatureSpec can return them in a different order.
	InputBindings []FeatureFunctionInputBinding `json:"input_bindings,omitempty"`
	// Output name of the feature that the UDF computes.
	Output string `json:"output,omitempty"`
	// Three-part name of the Python UDF.
	Udf string `json:"udf"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FeatureFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FeatureFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Binds one parameter of the UDF to a feature.
type FeatureFunctionInputBinding struct {
	// Name of the feature bound to the parameter.
	BoundFeature string `json:"bound_feature"`
	// Name of the UDF parameter.
	Parameter string `json:"parameter"`
}

func (s *FeatureFunctionInputBinding) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type FeatureLineage struct {
	// List of feature specs that contain this feature.
	FeatureSpecs []FeatureLineageFeatureSpec `json:"feature_specs,omitempty"`
	// List of Unity Catalog models that were trained on this feature.
	Models []FeatureLineageModel `json:"models,omitempty"`
	// List of online features that use this feature as source.
	OnlineFeatures []FeatureLineageOnlineFeature `json:"online_features,omitempty"`
}

func (s *FeatureLineage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type FeatureLineageFeatureSpec struct {
	// The full name of the feature spec in Unity Catalog.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FeatureLineageFeatureSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FeatureLineageFeatureSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FeatureLineageModel struct {
	// The full name of the model in Unity Catalog.
	Name string `json:"name,omitempty"`
	// The version of the model.
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FeatureLineageModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FeatureLineageModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FeatureLineageOnlineFeature struct {
	// The name of the online feature (column name).
	FeatureName string `json:"feature_name,omitempty"`
	// The full name of the online table in Unity Catalog.
	TableName string `json:"table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FeatureLineageOnlineFeature) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FeatureLineageOnlineFeature) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Feature list wrap all the features for a model version
type FeatureList struct {
	Features []LinkedFeature `json:"features,omitempty"`
}

func (s *FeatureList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Features to look up from a feature table.
type FeatureLookup struct {
	// Default values for features. Order is not significant, and GetFeatureSpec
	// can return them in a different order.
	DefaultValues []DefaultValue `json:"default_values,omitempty"`
	// Names of the features to look up, or empty to look up every feature in
	// the table.
	FeatureNames []string `json:"feature_names,omitempty"`
	// Lookback window to use when performing point-in-time lookup.
	LookbackWindow *duration.Duration `json:"lookback_window,omitempty"`
	// Keys to use when joining with the feature table.
	LookupKeys []string `json:"lookup_keys"`
	// Output names for features, in place of their names in the feature table.
	// Order is not significant, and GetFeatureSpec can return them in a
	// different order.
	RenameOutputs []RenameOutput `json:"rename_outputs,omitempty"`
	// Three-part name of the feature table.
	Table string `json:"table"`
	// Keys to use when performing point-in-time lookup.
	TimestampLookupKeys []string `json:"timestamp_lookup_keys,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FeatureLookup) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FeatureLookup) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A reference to one registered upstream Feature. A message rather than a bare
// name so an upstream can later be pinned more precisely (e.g. by version)
// without a breaking type change.
type FeatureReference struct {
	// The three-part full name of the upstream Feature.
	Feature string `json:"feature"`
}

func (s *FeatureReference) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A named set of features to serve, stored as a Unity Catalog function.
type FeatureSpec struct {
	// Output names of columns to leave out of the feature spec. Order is not
	// significant, and GetFeatureSpec can return them in a different order.
	ExcludeColumns []string `json:"exclude_columns,omitempty"`
	// Extra pip requirements to install in the serving container, in addition
	// to the defaults and those resolved from the features.
	ExtraPipRequirements *PipRequirements `json:"extra_pip_requirements,omitempty"`
	// The Features (either a declarative one, or a lookup from a feature table,
	// or a feature function) to serve. A declarative feature
	// (`DeclarativeFeature`) cannot be combined with lookups (`FeatureLookup`)
	// and functions (`FeatureFunction`) in one feature spec.
	Features []Features `json:"features,omitempty"`
	// Three-part name (catalog.schema.feature_spec) of the feature spec.
	Name string `json:"name"`
}

func (s *FeatureSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Represents a tag on a feature in a feature table.
type FeatureTag struct {
	Key string `json:"key"`

	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FeatureTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FeatureTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A data source composed from registered upstream Features.
type FeatureViewSource struct {
	// The upstream Features this source reads. Must include at least one
	// feature.
	FeatureReferences []FeatureReference `json:"feature_references,omitempty"`
}

func (s *FeatureViewSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// One entry of a feature spec: a Feature, a feature table lookup, or a feature
// function.
type Features struct {
	// A Feature, referenced by its three-part Unity Catalog name.
	DeclarativeFeature *DeclarativeFeature `json:"declarative_feature,omitempty"`
	// A feature function.
	FeatureFunction *FeatureFunction `json:"feature_function,omitempty"`
	// A feature table lookup.
	FeatureLookup *FeatureLookup `json:"feature_lookup,omitempty"`
}

func (s *Features) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A single field definition within a FlatSchema, specifying the field name and
// its scalar data type. Does not support nested or complex types (arrays, maps,
// structs).
type FieldDefinition struct {
	// The scalar data type of the field.
	DataType ScalarDataType `json:"data_type"`
	// The name of the field.
	Name string `json:"name"`
}

func (s *FieldDefinition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Metadata of a single artifact file or directory.
type FileInfo struct {
	// The size in bytes of the file. Unset for directories.
	FileSize int64 `json:"file_size,omitempty"`
	// Whether the path is a directory.
	IsDir bool `json:"is_dir,omitempty"`
	// The path relative to the root artifact directory run.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FinalizeLoggedModelRequest struct {
	// The ID of the logged model to finalize.
	ModelId string `json:"-" url:"-"`
	// Whether or not the model is ready for use.
	// ``"LOGGED_MODEL_UPLOAD_FAILED"`` indicates that something went wrong when
	// logging the model weights / agent code.
	Status LoggedModelStatus `json:"status"`
}

func (s *FinalizeLoggedModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type FinalizeLoggedModelResponse struct {
	// The updated logged model.
	Model *LoggedModel `json:"model,omitempty"`
}

func (s *FinalizeLoggedModelResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Returns the first N distinct values, ordered by the feature's timeseries
// column.
type FirstDistinctFunction struct {
	// The input column from which the first N distinct values are returned.
	Input string `json:"input"`
	// The number of distinct values to return.
	N int64 `json:"n"`
}

func (s *FirstDistinctFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Returns the first value.
type FirstFunction struct {
	// The input column from which the first value is returned.
	Input string `json:"input"`
}

func (s *FirstFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Returns the first N values, ordered by the feature's timeseries column.
type FirstNFunction struct {
	// The input column from which the first N values are returned.
	Input string `json:"input"`
	// The number of values to return.
	N int64 `json:"n"`
}

func (s *FirstNFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A flat (non-nested) schema for request-time fields, defined as an ordered
// list of field definitions. This schema only supports scalar types.
type FlatSchema struct {
	// The list of fields in this schema.
	Fields []FieldDefinition `json:"fields"`
}

func (s *FlatSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Represents a forecasting experiment with its unique identifier, URL, and
// state.
type ForecastingExperiment struct {
	// The unique ID for the forecasting experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// The URL to the forecasting experiment page.
	ExperimentPageUrl string `json:"experiment_page_url,omitempty"`
	// The current state of the forecasting experiment.
	State ForecastingExperimentState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ForecastingExperiment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForecastingExperiment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ForecastingExperimentState string

const ForecastingExperimentStateCancelled ForecastingExperimentState = `CANCELLED`

const ForecastingExperimentStateFailed ForecastingExperimentState = `FAILED`

const ForecastingExperimentStatePending ForecastingExperimentState = `PENDING`

const ForecastingExperimentStateRunning ForecastingExperimentState = `RUNNING`

const ForecastingExperimentStateSucceeded ForecastingExperimentState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *ForecastingExperimentState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ForecastingExperimentState) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCEEDED`:
		*f = ForecastingExperimentState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "PENDING", "RUNNING", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for ForecastingExperimentState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ForecastingExperimentState) Values() []ForecastingExperimentState {
	return []ForecastingExperimentState{
		ForecastingExperimentStateCancelled,
		ForecastingExperimentStateFailed,
		ForecastingExperimentStatePending,
		ForecastingExperimentStateRunning,
		ForecastingExperimentStateSucceeded,
	}
}

// Type always returns ForecastingExperimentState to satisfy [pflag.Value] interface
func (f *ForecastingExperimentState) Type() string {
	return "ForecastingExperimentState"
}

type Function struct {
	// An aggregation function applied over a time window.
	AggregationFunction *AggregationFunction `json:"aggregation_function,omitempty"`
	// Selects the latest value of a single column in a data source
	ColumnSelection *ColumnSelection `json:"column_selection,omitempty"`
	// Applies a registered Unity Catalog function row-wise to source columns.
	CustomUdf *CustomUdf `json:"custom_udf,omitempty"`

	ExtraParameters []FunctionExtraParameter `json:"extra_parameters,omitempty"`

	FunctionType FunctionFunctionType `json:"function_type,omitempty"`
	// Looks up an upstream Feature using entity values from other Features or
	// the request.
	Lookup *Lookup `json:"lookup,omitempty"`
	// Selects the secondary key values an upstream Feature has stored for the
	// request's entity.
	SecondaryKeys *SecondaryKeys `json:"secondary_keys,omitempty"`
}

func (s *Function) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type FunctionExtraParameter struct {
	// The name of the parameter.
	Key string `json:"key"`
	// The value of the parameter.
	Value string `json:"value"`
}

func (s *FunctionExtraParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type FunctionFunctionType string

const FunctionFunctionTypeApproxCountDistinct FunctionFunctionType = `APPROX_COUNT_DISTINCT`

const FunctionFunctionTypeApproxPercentile FunctionFunctionType = `APPROX_PERCENTILE`

const FunctionFunctionTypeAvg FunctionFunctionType = `AVG`

const FunctionFunctionTypeCount FunctionFunctionType = `COUNT`

const FunctionFunctionTypeFirst FunctionFunctionType = `FIRST`

const FunctionFunctionTypeFunctionTypeUnspecified FunctionFunctionType = `FUNCTION_TYPE_UNSPECIFIED`

const FunctionFunctionTypeLast FunctionFunctionType = `LAST`

const FunctionFunctionTypeMax FunctionFunctionType = `MAX`

const FunctionFunctionTypeMin FunctionFunctionType = `MIN`

const FunctionFunctionTypeStddevPop FunctionFunctionType = `STDDEV_POP`

const FunctionFunctionTypeStddevSamp FunctionFunctionType = `STDDEV_SAMP`

const FunctionFunctionTypeSum FunctionFunctionType = `SUM`

const FunctionFunctionTypeVarPop FunctionFunctionType = `VAR_POP`

const FunctionFunctionTypeVarSamp FunctionFunctionType = `VAR_SAMP`

// String representation for [fmt.Print]
func (f *FunctionFunctionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionFunctionType) Set(v string) error {
	switch v {
	case `APPROX_COUNT_DISTINCT`, `APPROX_PERCENTILE`, `AVG`, `COUNT`, `FIRST`, `FUNCTION_TYPE_UNSPECIFIED`, `LAST`, `MAX`, `MIN`, `STDDEV_POP`, `STDDEV_SAMP`, `SUM`, `VAR_POP`, `VAR_SAMP`:
		*f = FunctionFunctionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROX_COUNT_DISTINCT", "APPROX_PERCENTILE", "AVG", "COUNT", "FIRST", "FUNCTION_TYPE_UNSPECIFIED", "LAST", "MAX", "MIN", "STDDEV_POP", "STDDEV_SAMP", "SUM", "VAR_POP", "VAR_SAMP"`, v)
	}
}

// Values returns all possible values for FunctionFunctionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionFunctionType) Values() []FunctionFunctionType {
	return []FunctionFunctionType{
		FunctionFunctionTypeApproxCountDistinct,
		FunctionFunctionTypeApproxPercentile,
		FunctionFunctionTypeAvg,
		FunctionFunctionTypeCount,
		FunctionFunctionTypeFirst,
		FunctionFunctionTypeFunctionTypeUnspecified,
		FunctionFunctionTypeLast,
		FunctionFunctionTypeMax,
		FunctionFunctionTypeMin,
		FunctionFunctionTypeStddevPop,
		FunctionFunctionTypeStddevSamp,
		FunctionFunctionTypeSum,
		FunctionFunctionTypeVarPop,
		FunctionFunctionTypeVarSamp,
	}
}

// Type always returns FunctionFunctionType to satisfy [pflag.Value] interface
func (f *FunctionFunctionType) Type() string {
	return "FunctionFunctionType"
}

type GetByNameRequest struct {
	// Name of the associated experiment.
	ExperimentName string `json:"-" url:"experiment_name"`
}

func (s *GetByNameRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExperimentByNameResponse struct {
	// Experiment details.
	Experiment *Experiment `json:"experiment,omitempty"`
}

func (s *GetExperimentByNameResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" url:"-"`
}

func (s *GetExperimentPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExperimentPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ExperimentPermissionsDescription `json:"permission_levels,omitempty"`
}

func (s *GetExperimentPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" url:"-"`
}

func (s *GetExperimentPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId string `json:"-" url:"experiment_id"`
}

func (s *GetExperimentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExperimentResponse struct {
	// Experiment details.
	Experiment *Experiment `json:"experiment,omitempty"`
}

func (s *GetExperimentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetFeatureLineageRequest struct {
	// The name of the feature.
	FeatureName string `json:"-" url:"-"`
	// The full name of the feature table in Unity Catalog.
	TableName string `json:"-" url:"-"`
}

func (s *GetFeatureLineageRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetFeatureRequest struct {
	// Name of the feature to get.
	FullName string `json:"-" url:"-"`
}

func (s *GetFeatureRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetFeatureSpecRequest struct {
	// Three-part name (catalog.schema.feature_spec) of the feature spec to get.
	Name string `json:"-" url:"-"`
}

func (s *GetFeatureSpecRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetFeatureTagRequest struct {
	FeatureName string `json:"-" url:"-"`

	Key string `json:"-" url:"-"`

	TableName string `json:"-" url:"-"`
}

func (s *GetFeatureTagRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetForecastingExperimentRequest struct {
	// The unique ID of a forecasting experiment
	ExperimentId string `json:"-" url:"-"`
}

func (s *GetForecastingExperimentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetHistoryRequest struct {
	// Maximum number of Metric records to return per paginated request. Default
	// is set to 25,000. If set higher than 25,000, a request Exception will be
	// raised.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Name of the metric.
	MetricKey string `json:"-" url:"metric_key"`
	// Token indicating the page of metric histories to fetch.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// ID of the run from which to fetch metric values. Must be provided.
	RunId string `json:"-" url:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run from which to fetch
	// metric values. This field will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetHistoryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetHistoryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetIssueRequest struct {
	// The Issue to get, as `experiments/{experiment}/issues/{issue}`, where
	// `{experiment}` is the experiment ID and `{issue}` is the Issue ID: `iss-`
	// followed by 32 lowercase hex digits. For example,
	// `experiments/1234567890/issues/iss-0123456789abcdef0123456789abcdef`.
	Name string `json:"-" url:"-"`
}

func (s *GetIssueRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetKafkaConfigRequest struct {
	// Name of the Kafka config to get.
	Name string `json:"-" url:"-"`
}

func (s *GetKafkaConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetLatestVersionsRequest struct {
	// Registered model unique name identifier.
	Name string `json:"name"`
	// List of stages.
	Stages []string `json:"stages,omitempty"`
}

func (s *GetLatestVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions []ModelVersion `json:"model_versions,omitempty"`
}

func (s *GetLatestVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetLoggedModelRequest struct {
	// The ID of the logged model to retrieve.
	ModelId string `json:"-" url:"-"`
}

func (s *GetLoggedModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetLoggedModelResponse struct {
	// The retrieved logged model.
	Model *LoggedModel `json:"model,omitempty"`
}

func (s *GetLoggedModelResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetMaterializationSetRequest struct {
	// The ID of the materialization set.
	MaterializationSetId string `json:"-" url:"-"`
}

func (s *GetMaterializationSetRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetMaterializedFeatureRequest struct {
	// The ID of the materialized feature.
	MaterializedFeatureId string `json:"-" url:"-"`
}

func (s *GetMaterializedFeatureRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetMetricHistoryResponse struct {
	// All logged values for this metric if `max_results` is not specified in
	// the request or if the total count of metrics returned is less than the
	// service level pagination threshold. Otherwise, this is one page of
	// results.
	Metrics []Metric `json:"metrics,omitempty"`
	// A token that can be used to issue a query for the next page of metric
	// history values. A missing token indicates that no additional metrics are
	// available to fetch.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetMetricHistoryResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetMetricHistoryResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"-" url:"name"`
}

func (s *GetModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetModelResponse struct {
	RegisteredModelDatabricks *ModelDatabricks `json:"registered_model_databricks,omitempty"`
}

func (s *GetModelResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	Name string `json:"-" url:"name"`
	// Model version number
	Version string `json:"-" url:"version"`
}

func (s *GetModelVersionDownloadUriRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetModelVersionDownloadUriResponse struct {
	// URI corresponding to where artifacts for this model version are stored.
	ArtifactUri string `json:"artifact_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetModelVersionDownloadUriResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetModelVersionDownloadUriResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetModelVersionRequest struct {
	// Name of the registered model
	Name string `json:"-" url:"name"`
	// Model version number
	Version string `json:"-" url:"version"`
}

func (s *GetModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetModelVersionResponse struct {
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

func (s *GetModelVersionResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetOnlineStoreRequest struct {
	// Name of the online store to get.
	Name string `json:"-" url:"-"`
}

func (s *GetOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `json:"-" url:"-"`
}

func (s *GetOperationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" url:"-"`
}

func (s *GetRegisteredModelPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRegisteredModelPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []RegisteredModelPermissionsDescription `json:"permission_levels,omitempty"`
}

func (s *GetRegisteredModelPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" url:"-"`
}

func (s *GetRegisteredModelPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	RunId string `json:"-" url:"run_id"`
	// [Deprecated, use `run_id` instead] ID of the run to fetch. This field
	// will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetRunRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRunRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetRunResponse struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	Run *Run `json:"run,omitempty"`
}

func (s *GetRunResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetStreamRequest struct {
	// Full three-part name (catalog.schema.stream) of the Stream to get.
	Name string `json:"-" url:"-"`
}

func (s *GetStreamRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type HttpUrlSpec struct {
	// Value of the authorization header that should be sent in the request sent
	// by the wehbook. It should be of the form `"<auth type> <credentials>"`.
	// If set to an empty string, no authorization header will be included in
	// the request.
	Authorization string `json:"authorization,omitempty"`
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification bool `json:"enable_ssl_verification,omitempty"`
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	Secret string `json:"secret,omitempty"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url string `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *HttpUrlSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s HttpUrlSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type HttpUrlSpecWithoutSecret struct {
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification bool `json:"enable_ssl_verification,omitempty"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *HttpUrlSpecWithoutSecret) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s HttpUrlSpecWithoutSecret) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Unity Catalog managed or foreign Iceberg table, in Iceberg v2 or v3 format.
// Table identity, format, schema, and permissions are resolved through Unity
// Catalog.
type IcebergTableSource struct {
	// Schema of the resulting dataframe after transformations, in Spark
	// StructType JSON format (from df.schema.json()). Required if
	// transformation_sql is specified. Example:
	// {"type":"struct","fields":[{"name":"col_a","type":"integer","nullable":true,"metadata":{}},{"name":"col_c","type":"integer","nullable":true,"metadata":{}}]}
	DataframeSchema string `json:"dataframe_schema,omitempty"`
	// A single WHERE clause to filter the Iceberg table before applying
	// transformations. The predicate is evaluated for each row.
	FilterCondition string `json:"filter_condition,omitempty"`
	// Full three-part name of the source table. May be omitted on reads when
	// inaccessible.
	FullName string `json:"full_name"`
	// A single SQL SELECT expression applied after filter_condition. The
	// expression must contain all required columns (e.g., "*, col_a + col_b AS
	// col_c"). If transformation_sql is not provided, all columns of the
	// Iceberg table are present in the DataSource dataframe.
	TransformationSql string `json:"transformation_sql,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IcebergTableSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IcebergTableSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration for the Databricks-managed ingestion pipeline. Groups the
// ingestion destination (required) and optional backfill source.
type IngestionConfig struct {
	// The ID of the Databricks Job that performs the historical backfill of the
	// ingestion Delta table.
	BackfillJobId int64 `json:"backfill_job_id,omitempty"`
	// A user-provided source for backfilling data. Historical data is used when
	// creating a training set from streaming features linked to this Stream.
	// The backfill data stored in this location will be copied into the
	// ingestion table for offline querying and training. The schema for this
	// source must match exactly that of the key and payload schemas specified
	// for this Stream, except that it may omit any columns listed in
	// excluded_columns.
	BackfillSource *BackfillSource `json:"backfill_source,omitempty"`
	// The ID of the budget policy used to attribute the serverless compute cost
	// of this stream's managed ingestion. If not specified, a default budget
	// policy may be applied.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Column paths used to identify duplicate rows during ingestion; only one
	// row per distinct combination of these values is kept. Use dot notation
	// for nested fields (e.g. `value.user_id`). Empty list means every column
	// is compared.
	DeduplicationColumns []string `json:"deduplication_columns,omitempty"`
	// Destination for the Databricks-managed Delta table that holds an offline
	// copy of the streaming data for querying and training. This table contains
	// both 1) forward-filled data from the Stream and 2) backfilled data from
	// the BackfillSource (if provided). This table is created and managed by
	// Databricks and is deleted when the Stream is deleted.
	IngestionDestination IngestionDestination `json:"ingestion_destination"`
	// The ID of the Databricks Job that performs the forward-fill ingestion.
	IngestionJobId int64 `json:"ingestion_job_id,omitempty"`
	// The ID of the SDP pipeline that continuously copies new events from the
	// streaming source into the ingestion Delta table.
	IngestionPipelineId string `json:"ingestion_pipeline_id,omitempty"`
	// Custom tags to associate with this stream's managed ingestion. They are
	// applied to the ingestion pipeline and its forward-fill and backfill jobs,
	// and forwarded to the underlying compute as cluster tags, so ingestion
	// cost can be attributed in the billing system tables. These tags apply
	// only to the managed ingestion compute; they are not applied to the Stream
	// entity itself, and are distinct from any Unity Catalog tags on the
	// Stream. A maximum of 25 tags is supported; keys and values are subject to
	// the same limitations as cluster tags.
	Tags map[string]string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IngestionConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Destination for the Databricks-managed Delta table that holds an offline copy
// of the streaming data for querying and training.
type IngestionDestination struct {
	// The full three-part name (catalog, schema, name) of the Delta table to be
	// created for ingestion.
	DeltaTableName string `json:"delta_table_name,omitempty"`
	// The full three-part name (catalog, schema, name) of an existing view to
	// use instead of creating a Delta table. The view is dropped when the
	// Stream is deleted.
	ViewName string `json:"view_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IngestionDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Binds a single UC function parameter to a source column.
type InputBinding struct {
	// Source column whose value is passed for this parameter at execution time.
	Column string `json:"column"`
	// Name of the UC function parameter.
	Parameter string `json:"parameter"`
}

func (s *InputBinding) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The input that supplies a value. A value from the request comes from a
// Feature whose source is a RequestSource.
type InputReference struct {
	// Specifies another Feature whose output supplies the value.
	Feature *FeatureReference `json:"feature,omitempty"`
}

func (s *InputReference) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Tag for a dataset input.
type InputTag struct {
	// The tag key.
	Key string `json:"key"`
	// The tag value.
	Value string `json:"value"`
}

func (s *InputTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A quality or operational problem found in an experiment's traces.
type Issue struct {
	// At most 10 values of at most 64 UTF-8 bytes each, unique
	// case-insensitively. Leading and trailing whitespace is trimmed.
	Categories []string `json:"categories,omitempty"`
	// When the Issue was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The email of the user whose request created the Issue. source_type tells
	// whether issue detection created it.
	CreatedBy string `json:"created_by,omitempty"`
	// Markdown, at most 32 KiB.
	Description string `json:"description,omitempty"`
	// The last segment of `name`. Issue assessments on traces reference the
	// Issue by this ID.
	IssueId string `json:"issue_id,omitempty"`
	// The request time of the latest trace the Issue affects. Required, and
	// can't be cleared, unless source_type is MANUAL. A MANUAL Issue without
	// one returns its create_time.
	LastSeenTime *time.Time `json:"last_seen_time,omitempty"`
	// The email of the user whose request last updated the Issue.
	LastUpdatedBy string `json:"last_updated_by,omitempty"`
	// Additional details about the Issue, such as how issue detection observed
	// it. Not a stable contract: keys and values can change without notice, so
	// clients must tolerate unknown and missing keys. At most 20 entries, with
	// keys of at most 64 UTF-8 bytes. Values are at most 256 UTF-8 bytes,
	// except recent_affected_trace_ids (4 KiB) and suggested_fix (16 KiB).
	Metadata map[string]string `json:"metadata,omitempty"`
	// The Issue's resource name, as `experiments/{experiment}/issues/{issue}`,
	// where `{experiment}` is the experiment ID and `{issue}` is the Issue ID:
	// `iss-` followed by 32 lowercase hex digits. For example,
	// `experiments/1234567890/issues/iss-0123456789abcdef0123456789abcdef`.
	Name string `json:"name,omitempty"`
	// Cannot be cleared: an Issue without a severity would match no severity
	// filter.
	Severity IssueSeverity `json:"severity"`
	// Set once at create. For ADHOC_RUN and SCHEDULED_RUN, the ID of the issue
	// detection run in the same experiment that created the Issue. Absent for
	// MANUAL.
	SourceId string `json:"source_id,omitempty"`
	// Set once at create. Defaults to MANUAL.
	SourceType IssueSourceType `json:"source_type,omitempty"`
	// Defaults to PENDING on create, and can't be cleared.
	Status IssueStatus `json:"status,omitempty"`
	// Caller-owned; the server never writes system values here. At most 10
	// entries. Keys match ^[A-Za-z0-9_.:/-]+$ and are at most 128 UTF-8 bytes;
	// values are at most 256 UTF-8 bytes.
	Tags map[string]string `json:"tags,omitempty"`
	// Non-empty, at most 250 UTF-8 bytes.
	Title string `json:"title"`
	// When the Issue was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Issue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Issue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// How serious an Issue is, ordered LOW < MEDIUM < HIGH.
type IssueSeverity string

const IssueSeverityHigh IssueSeverity = `HIGH`

const IssueSeverityLow IssueSeverity = `LOW`

const IssueSeverityMedium IssueSeverity = `MEDIUM`

// String representation for [fmt.Print]
func (f *IssueSeverity) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IssueSeverity) Set(v string) error {
	switch v {
	case `HIGH`, `LOW`, `MEDIUM`:
		*f = IssueSeverity(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "HIGH", "LOW", "MEDIUM"`, v)
	}
}

// Values returns all possible values for IssueSeverity.
//
// There is no guarantee on the order of the values in the slice.
func (f *IssueSeverity) Values() []IssueSeverity {
	return []IssueSeverity{
		IssueSeverityHigh,
		IssueSeverityLow,
		IssueSeverityMedium,
	}
}

// Type always returns IssueSeverity to satisfy [pflag.Value] interface
func (f *IssueSeverity) Type() string {
	return "IssueSeverity"
}

// The field that search results are ordered by. SEVERITY breaks ties by
// last_seen_time, in the same direction. Results are grouped by status first
// (PENDING, RESOLVED, then REJECTED, whatever the sort_order), and remaining
// ties are broken by issue ID, so paging is stable.
type IssueSortKey string

const IssueSortKeyCreateTime IssueSortKey = `CREATE_TIME`

const IssueSortKeyLastSeenTime IssueSortKey = `LAST_SEEN_TIME`

const IssueSortKeySeverity IssueSortKey = `SEVERITY`

const IssueSortKeyUpdateTime IssueSortKey = `UPDATE_TIME`

// String representation for [fmt.Print]
func (f *IssueSortKey) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IssueSortKey) Set(v string) error {
	switch v {
	case `CREATE_TIME`, `LAST_SEEN_TIME`, `SEVERITY`, `UPDATE_TIME`:
		*f = IssueSortKey(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATE_TIME", "LAST_SEEN_TIME", "SEVERITY", "UPDATE_TIME"`, v)
	}
}

// Values returns all possible values for IssueSortKey.
//
// There is no guarantee on the order of the values in the slice.
func (f *IssueSortKey) Values() []IssueSortKey {
	return []IssueSortKey{
		IssueSortKeyCreateTime,
		IssueSortKeyLastSeenTime,
		IssueSortKeySeverity,
		IssueSortKeyUpdateTime,
	}
}

// Type always returns IssueSortKey to satisfy [pflag.Value] interface
func (f *IssueSortKey) Type() string {
	return "IssueSortKey"
}

// What created an Issue: ADHOC_RUN, an issue detection run started on demand;
// SCHEDULED_RUN, one started by a schedule; SINGLE_TRACE_ANALYSIS, an analysis
// of a single trace; or MANUAL, a person.
type IssueSourceType string

const IssueSourceTypeAdhocRun IssueSourceType = `ADHOC_RUN`

const IssueSourceTypeManual IssueSourceType = `MANUAL`

const IssueSourceTypeScheduledRun IssueSourceType = `SCHEDULED_RUN`

const IssueSourceTypeSingleTraceAnalysis IssueSourceType = `SINGLE_TRACE_ANALYSIS`

// String representation for [fmt.Print]
func (f *IssueSourceType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IssueSourceType) Set(v string) error {
	switch v {
	case `ADHOC_RUN`, `MANUAL`, `SCHEDULED_RUN`, `SINGLE_TRACE_ANALYSIS`:
		*f = IssueSourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADHOC_RUN", "MANUAL", "SCHEDULED_RUN", "SINGLE_TRACE_ANALYSIS"`, v)
	}
}

// Values returns all possible values for IssueSourceType.
//
// There is no guarantee on the order of the values in the slice.
func (f *IssueSourceType) Values() []IssueSourceType {
	return []IssueSourceType{
		IssueSourceTypeAdhocRun,
		IssueSourceTypeManual,
		IssueSourceTypeScheduledRun,
		IssueSourceTypeSingleTraceAnalysis,
	}
}

// Type always returns IssueSourceType to satisfy [pflag.Value] interface
func (f *IssueSourceType) Type() string {
	return "IssueSourceType"
}

// The triage status of an Issue: PENDING means open and awaiting triage,
// RESOLVED means judged real and addressed, and REJECTED means judged not to be
// an issue.
type IssueStatus string

const IssueStatusPending IssueStatus = `PENDING`

const IssueStatusRejected IssueStatus = `REJECTED`

const IssueStatusResolved IssueStatus = `RESOLVED`

// String representation for [fmt.Print]
func (f *IssueStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IssueStatus) Set(v string) error {
	switch v {
	case `PENDING`, `REJECTED`, `RESOLVED`:
		*f = IssueStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PENDING", "REJECTED", "RESOLVED"`, v)
	}
}

// Values returns all possible values for IssueStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *IssueStatus) Values() []IssueStatus {
	return []IssueStatus{
		IssueStatusPending,
		IssueStatusRejected,
		IssueStatusResolved,
	}
}

// Type always returns IssueStatus to satisfy [pflag.Value] interface
func (f *IssueStatus) Type() string {
	return "IssueStatus"
}

// One predicate over Issue.tags.
type IssueTagFilter struct {
	// Matched exactly.
	Key string `json:"key"`
	// Matched exactly, and ORed together. Empty matches any Issue carrying the
	// key.
	Values []string `json:"values,omitempty"`
}

func (s *IssueTagFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Which fields each search result includes: BASIC has every field except
// `description`, and may omit large `metadata` entries; FULL has every field.
type IssueView string

const IssueViewBasic IssueView = `BASIC`

const IssueViewFull IssueView = `FULL`

// String representation for [fmt.Print]
func (f *IssueView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IssueView) Set(v string) error {
	switch v {
	case `BASIC`, `FULL`:
		*f = IssueView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BASIC", "FULL"`, v)
	}
}

// Values returns all possible values for IssueView.
//
// There is no guarantee on the order of the values in the slice.
func (f *IssueView) Values() []IssueView {
	return []IssueView{
		IssueViewBasic,
		IssueViewFull,
	}
}

// Type always returns IssueView to satisfy [pflag.Value] interface
func (f *IssueView) Type() string {
	return "IssueView"
}

type JobContext struct {
	// The job ID where this API invoked.
	JobId int64 `json:"job_id,omitempty"`
	// The job run ID where this API was invoked.
	JobRunId int64 `json:"job_run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobContext) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobContext) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobSpec struct {
	// The personal access token used to authorize webhook's job runs.
	AccessToken string `json:"access_token"`
	// ID of the job that the webhook runs.
	JobId string `json:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobSpecWithoutSecret struct {
	// ID of the job that the webhook runs.
	JobId string `json:"job_id,omitempty"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobSpecWithoutSecret) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobSpecWithoutSecret) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KafkaConfig struct {
	// Authentication configuration for connection to topics.
	AuthConfig AuthConfig `json:"auth_config"`
	// A user-provided and managed source for backfilling data. Historical data
	// is used when creating a training set from streaming features linked to
	// this Kafka config. In the future, a separate table will be maintained by
	// Databricks for forward filling data. The schema for this source must
	// match exactly that of the key and value schemas specified for this Kafka
	// config.
	BackfillSource *BackfillSource `json:"backfill_source,omitempty"`
	// A comma-separated list of host/port pairs pointing to Kafka cluster.
	BootstrapServers string `json:"bootstrap_servers"`
	// Catch-all for miscellaneous options. Keys should be source options or
	// Kafka consumer options (kafka.*)
	ExtraOptions map[string]string `json:"extra_options,omitempty"`
	// Configuration for ingesting Kafka data into a Databricks-managed Delta
	// table.
	IngestionConfig *IngestionConfig `json:"ingestion_config,omitempty"`
	// Schema configuration for extracting message keys from topics. At least
	// one of key_schema and value_schema must be provided.
	KeySchema *SchemaConfig `json:"key_schema,omitempty"`
	// Name that uniquely identifies this Kafka config within the metastore.
	// This will be the identifier used from the Feature object to reference
	// these configs for a feature. Can be distinct from topic name.
	Name string `json:"name"`
	// Options to configure which Kafka topics to pull data from.
	SubscriptionMode SubscriptionMode `json:"subscription_mode"`
	// Schema configuration for extracting message values from topics. At least
	// one of key_schema and value_schema must be provided.
	ValueSchema *SchemaConfig `json:"value_schema,omitempty"`
}

func (s *KafkaConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type KafkaSource struct {
	EntityColumnIdentifiers []ColumnIdentifier `json:"entity_column_identifiers,omitempty"`
	// The filter condition applied to the source data before aggregation.
	FilterCondition string `json:"filter_condition,omitempty"`
	// Name of the Kafka source, used to identify it. This is used to look up
	// the corresponding KafkaConfig object. Can be distinct from topic name.
	Name string `json:"name"`

	TimeseriesColumnIdentifier *ColumnIdentifier `json:"timeseries_column_identifier,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KafkaSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KafkaSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Kafka-specific configuration for a Stream.
type KafkaStreamConfig struct {
	// Optional Kafka source or consumer options, validated against a
	// server-side allowlist at request time. Allowed keys: -
	// `maxOffsetsPerTrigger` - `startingOffsets` - `includeHeaders` -
	// `kafka.request.timeout.ms` - `kafka.session.timeout.ms` -
	// `kafka.max.partition.fetch.bytes` The following keys are ingestion-only
	// and are stripped before being forwarded to the materialization pipeline:
	// - `maxOffsetsPerTrigger` - `startingOffsets` Auth and connection details
	// belong on the parent Stream's `connection_config`, not here.
	ExtraOptions map[string]string `json:"extra_options,omitempty"`
	// Options to configure which Kafka topics to pull data from.
	SubscriptionMode KafkaSubscriptionMode `json:"subscription_mode"`
}

func (s *KafkaStreamConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Subscription mode for Kafka topic selection, matching standard Spark
// Structured Streaming options.
type KafkaSubscriptionMode struct {
	// A JSON string that contains the specific topic-partitions to consume
	// from. For example, for '{"topicA":[0,1],"topicB":[2,4]}', topicA's 0'th
	// and 1st partitions will be consumed from.
	Assign string `json:"assign,omitempty"`
	// A comma-separated list of Kafka topics to read from. For example,
	// 'topicA,topicB,topicC'.
	Subscribe string `json:"subscribe,omitempty"`
	// A regular expression matching topics to subscribe to. For example,
	// 'topic.*' will subscribe to all topics starting with 'topic'.
	SubscribePattern string `json:"subscribe_pattern,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KafkaSubscriptionMode) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KafkaSubscriptionMode) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Kinesis-specific configuration for a Stream. For the underlying connector and
// its source options, see the Databricks documentation on connecting to Amazon
// Kinesis (https://docs.databricks.com/aws/en/connect/streaming/kinesis).
type KinesisStreamConfig struct {
	// Optional Kinesis source options, validated against a server-side
	// allowlist at request time. Allowed keys: - `consumerMode` -
	// `consumerNamePrefix` - `maxFetchRate` - `minFetchPeriod` -
	// `maxFetchDuration` - `maxRecordsPerFetch` - `shardsPerTask` -
	// `fetchBufferSize` - `shardFetchInterval` `consumerMode` must be `efo` or
	// `polling` (case-insensitive). `maxRecordsPerFetch` applies only during
	// ingestion and does not affect the materialization pipeline. Auth and
	// connection details belong on the parent Stream's `connection_config`, not
	// here.
	ExtraOptions map[string]string `json:"extra_options,omitempty"`
	// Kinesis stream ARNs to read from.
	StreamArns *StreamArnList `json:"stream_arns,omitempty"`
	// Kinesis stream names to read from.
	StreamNames *StreamNameList `json:"stream_names,omitempty"`
}

func (s *KinesisStreamConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Returns the last N distinct values, ordered by the feature's timeseries
// column.
type LastDistinctFunction struct {
	// The input column from which the last N distinct values are returned.
	Input string `json:"input"`
	// The number of distinct values to return.
	N int64 `json:"n"`
}

func (s *LastDistinctFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Returns the last value.
type LastFunction struct {
	// The input column from which the last value is returned.
	Input string `json:"input"`
}

func (s *LastFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Returns the last N values, ordered by the feature's timeseries column.
type LastNFunction struct {
	// The input column from which the last N values are returned.
	Input string `json:"input"`
	// The number of values to return.
	N int64 `json:"n"`
}

func (s *LastNFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Lineage context information for tracking where an API was invoked. This will
// allow us to track lineage, which currently uses caller entity information for
// use across the Lineage Client and Observability in Lumberjack.
type LineageContext struct {
	// Job context information including job ID and run ID.
	JobContext *JobContext `json:"job_context,omitempty"`
	// The notebook ID where this API was invoked.
	NotebookId int64 `json:"notebook_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LineageContext) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LineageContext) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Feature for model version.
type LinkedFeature struct {
	// Feature name
	FeatureName string `json:"feature_name,omitempty"`
	// Feature table id
	FeatureTableId string `json:"feature_table_id,omitempty"`
	// Feature table name
	FeatureTableName string `json:"feature_table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LinkedFeature) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LinkedFeature) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListArtifactsRequest struct {
	// The token indicating the page of artifact results to fetch. `page_token`
	// is not supported when listing artifacts in UC Volumes. A maximum of 1000
	// artifacts will be retrieved for UC Volumes. Please call
	// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
	// Volumes, which supports pagination. See [List directory contents | Files
	// API](/api/workspace/files/listdirectorycontents).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	Path string `json:"-" url:"path,omitempty"`
	// ID of the run whose artifacts to list. Must be provided.
	RunId string `json:"-" url:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListArtifactsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListArtifactsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListArtifactsResponse struct {
	// The file location and metadata for artifacts.
	Files []FileInfo `json:"files,omitempty"`
	// The token that can be used to retrieve the next page of artifact results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The root artifact directory for the run.
	RootUri string `json:"root_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListArtifactsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListArtifactsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExperimentsRequest struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it'll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// Page token returned as `next_page_token` by a previous call. Omit it to
	// fetch the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType ViewType `json:"-" url:"view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExperimentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExperimentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExperimentsResponse struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	Experiments []Experiment `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExperimentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExperimentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFeatureTagsRequest struct {
	FeatureName string `json:"-" url:"-"`
	// The maximum number of results to return.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	TableName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFeatureTagsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFeatureTagsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message for ListFeatureTag.
type ListFeatureTagsResponse struct {
	FeatureTags []FeatureTag `json:"feature_tags,omitempty"`
	// Pagination token to request the next page of results for this query.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFeatureTagsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFeatureTagsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFeaturesRequest struct {
	// Name of parent catalog for features of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// The maximum number of results to return.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFeaturesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFeaturesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFeaturesResponse struct {
	// List of features.
	Features []Feature `json:"features,omitempty"`
	// Pagination token to request the next page of results for this query.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFeaturesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFeaturesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListKafkaConfigsRequest struct {
	// The maximum number of results to return.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListKafkaConfigsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListKafkaConfigsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListKafkaConfigsResponse struct {
	// List of Kafka configs. Schemas are not included in the response.
	KafkaConfigs []KafkaConfig `json:"kafka_configs"`
	// Pagination token to request the next page of results for this query.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListKafkaConfigsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListKafkaConfigsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListMaterializedFeaturesRequest struct {
	// Filter by feature name. If specified, only materialized features
	// materialized from this feature will be returned.
	FeatureName string `json:"-" url:"feature_name,omitempty"`
	// The maximum number of results to return. Defaults to 100 if not
	// specified. Cannot be greater than 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMaterializedFeaturesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMaterializedFeaturesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListMaterializedFeaturesResponse struct {
	// List of materialized features.
	MaterializedFeatures []MaterializedFeature `json:"materialized_features,omitempty"`
	// Pagination token to request the next page of results for this query.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMaterializedFeaturesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMaterializedFeaturesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListModelsResponse struct {
	// Pagination token to request next page of models for the same query.
	NextPageToken string `json:"next_page_token,omitempty"`

	RegisteredModels []Model `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListModelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListOnlineStoresRequest struct {
	// The maximum number of results to return. Defaults to 100 if not
	// specified.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListOnlineStoresRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListOnlineStoresRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListOnlineStoresResponse struct {
	// Pagination token to request the next page of results for this query.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of online stores.
	OnlineStores []OnlineStore `json:"online_stores,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListOnlineStoresResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListOnlineStoresResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken string `json:"next_page_token,omitempty"`
	// Array of registry webhooks.
	Webhooks []RegistryWebhook `json:"webhooks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRegistryWebhooks) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRegistryWebhooks) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListStreamsRequest struct {
	// The maximum number of results to return.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Two-part name (catalog.schema) of the parent under which to list Streams.
	Parent string `json:"-" url:"parent,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListStreamsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStreamsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response to a ListStreamsRequest.
//
// NOTE: Results are post-filtered by access permission on each stream's
// ingestion table. This means: - Returned results may be fewer than page_size
// (including zero) - Page token points to next unfiltered batch, not next
// filtered batch, and may point to an item that will be filtered out Callers
// should paginate until next_page_token is empty to retrieve all accessible
// streams.
type ListStreamsResponse struct {
	// Pagination token to request the next page of results for this query.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of Streams.
	Streams []Stream `json:"streams,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListStreamsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStreamsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTransitionRequestsRequest struct {
	// Name of the registered model.
	Name string `json:"-" url:"name"`
	// Version of the model.
	Version string `json:"-" url:"version"`
}

func (s *ListTransitionRequestsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListTransitionRequestsResponse struct {
	// Array of open transition requests.
	Requests []Activity `json:"requests,omitempty"`
}

func (s *ListTransitionRequestsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListWebhooksRequest struct {
	// Events that trigger the webhook. * `MODEL_VERSION_CREATED`: A new model
	// version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	//
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	Events []RegistryWebhookEvent `json:"-" url:"events,omitempty"`

	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// Registered model name If not specified, all webhooks associated with the
	// specified events are listed, regardless of their associated model.
	ModelName string `json:"-" url:"model_name,omitempty"`
	// Token indicating the page of artifact results to fetch
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWebhooksRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWebhooksRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LogBatch struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	Metrics []Metric `json:"metrics,omitempty"`
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	Params []Param `json:"params,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id,omitempty"`
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	Tags []RunTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogBatch) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogBatch) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LogInputs struct {
	// Dataset inputs
	Datasets []DatasetInput `json:"datasets,omitempty"`
	// Model inputs
	Models []ModelInput `json:"models,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id"`
}

func (s *LogInputs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type LogLoggedModelParamsRequest struct {
	// The ID of the logged model to log params for.
	ModelId string `json:"-" url:"-"`
	// Parameters to attach to the model.
	Params []LoggedModelParameter `json:"params,omitempty"`
}

func (s *LogLoggedModelParamsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type LogMetric struct {
	// Dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	DatasetName string `json:"dataset_name,omitempty"`
	// Name of the metric.
	Key string `json:"key"`
	// ID of the logged model associated with the metric, if applicable
	ModelId string `json:"model_id,omitempty"`
	// ID of the run under which to log the metric. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Step at which to log the metric
	Step int64 `json:"step,omitempty"`
	// Unix timestamp in milliseconds at the time metric was logged.
	Timestamp int64 `json:"timestamp"`
	// Double value of the metric being logged.
	Value float64 `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogMetric) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogMetric) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LogModel struct {
	// MLmodel file in json format.
	ModelJson string `json:"model_json,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LogOutputsRequest struct {
	// The model outputs from the Run.
	Models []ModelOutput `json:"models,omitempty"`
	// The ID of the Run from which to log outputs.
	RunId string `json:"run_id"`
}

func (s *LogOutputsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type LogParam struct {
	// Name of the param. Maximum size is 255 bytes.
	Key string `json:"key"`
	// ID of the run under which to log the param. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// String value of the param being logged. Maximum size is 500 bytes.
	Value string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogParam) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogParam) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A logged model message includes logged model attributes, tags, registration
// info, params, and linked run metrics.
type LoggedModel struct {
	// The params and metrics attached to the logged model.
	Data *LoggedModelData `json:"data,omitempty"`
	// The logged model attributes such as model ID, status, tags, etc.
	Info *LoggedModelInfo `json:"info,omitempty"`
}

func (s *LoggedModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A LoggedModelData message includes logged model params and linked metrics.
type LoggedModelData struct {
	// Performance metrics linked to the model.
	Metrics []Metric `json:"metrics,omitempty"`
	// Immutable string key-value pairs of the model.
	Params []LoggedModelParameter `json:"params,omitempty"`
}

func (s *LoggedModelData) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A LoggedModelInfo includes logged model attributes, tags, and registration
// info.
type LoggedModelInfo struct {
	// The URI of the directory where model artifacts are stored.
	ArtifactUri string `json:"artifact_uri,omitempty"`
	// The timestamp when the model was created in milliseconds since the UNIX
	// epoch.
	CreationTimestampMs int64 `json:"creation_timestamp_ms,omitempty"`
	// The ID of the user or principal that created the model.
	CreatorId int64 `json:"creator_id,omitempty"`
	// The ID of the experiment that owns the model.
	ExperimentId string `json:"experiment_id,omitempty"`
	// The timestamp when the model was last updated in milliseconds since the
	// UNIX epoch.
	LastUpdatedTimestampMs int64 `json:"last_updated_timestamp_ms,omitempty"`
	// The unique identifier for the logged model.
	ModelId string `json:"model_id,omitempty"`
	// The type of model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	ModelType string `json:"model_type,omitempty"`
	// The name of the model.
	Name string `json:"name,omitempty"`
	// The ID of the run that created the model.
	SourceRunId string `json:"source_run_id,omitempty"`
	// The status of whether or not the model is ready for use.
	Status LoggedModelStatus `json:"status,omitempty"`
	// Details on the current model status.
	StatusMessage string `json:"status_message,omitempty"`
	// Mutable string key-value pairs set on the model.
	Tags []LoggedModelTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LoggedModelInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LoggedModelInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Parameter associated with a LoggedModel.
type LoggedModelParameter struct {
	// The key identifying this param.
	Key string `json:"key,omitempty"`
	// The value of this param.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LoggedModelParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LoggedModelParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A LoggedModelStatus enum value represents the status of a logged model.
type LoggedModelStatus string

const LoggedModelStatusLoggedModelPending LoggedModelStatus = `LOGGED_MODEL_PENDING`

const LoggedModelStatusLoggedModelReady LoggedModelStatus = `LOGGED_MODEL_READY`

const LoggedModelStatusLoggedModelUploadFailed LoggedModelStatus = `LOGGED_MODEL_UPLOAD_FAILED`

// String representation for [fmt.Print]
func (f *LoggedModelStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LoggedModelStatus) Set(v string) error {
	switch v {
	case `LOGGED_MODEL_PENDING`, `LOGGED_MODEL_READY`, `LOGGED_MODEL_UPLOAD_FAILED`:
		*f = LoggedModelStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LOGGED_MODEL_PENDING", "LOGGED_MODEL_READY", "LOGGED_MODEL_UPLOAD_FAILED"`, v)
	}
}

// Values returns all possible values for LoggedModelStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *LoggedModelStatus) Values() []LoggedModelStatus {
	return []LoggedModelStatus{
		LoggedModelStatusLoggedModelPending,
		LoggedModelStatusLoggedModelReady,
		LoggedModelStatusLoggedModelUploadFailed,
	}
}

// Type always returns LoggedModelStatus to satisfy [pflag.Value] interface
func (f *LoggedModelStatus) Type() string {
	return "LoggedModelStatus"
}

// Tag for a LoggedModel.
type LoggedModelTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LoggedModelTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LoggedModelTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Looks up an upstream Feature using entity values bound from other Features or
// from the request. When a binding is an array, such as the output of a
// SecondaryKeys Feature, returns one value per position as an array aligned
// with that binding. The Feature's source must be a FeatureViewSource that
// includes `feature` and every Feature referenced in `entities`. Currently,
// exactly one entity must be bound to a Feature: a SecondaryKeys Feature whose
// key has the same name as that entity. Its target can be `feature` or another
// Feature. Every other entity must be bound to a request column of the same
// name. Names are compared without struct prefixes, so `user_id` matches
// `value.user_id`. A FeatureSpec can't also look up `feature` with all of its
// entities supplied.
type Lookup struct {
	// One binding per entity column of the upstream Feature. Every entity
	// column must be bound exactly once.
	Entities []EntityBinding `json:"entities"`
	// The upstream Feature to look up.
	Feature FeatureReference `json:"feature"`
}

func (s *Lookup) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Recipients to notify when a materialization run fails.
type MaterializationFailureNotification struct {
	// IDs of the notification destinations (for example Slack, Microsoft Teams,
	// PagerDuty, or a generic webhook) to notify. Not supported for streaming
	// materialized features.
	DestinationIds []string `json:"destination_ids,omitempty"`
	// Email addresses to notify.
	EmailAddresses []string `json:"email_addresses,omitempty"`
	// If true, notify only when the final attempt of a run fails. If false or
	// unset, notify on every failed attempt, including attempts that will be
	// retried.
	//
	// Batch materialization does not retry failures that need a fix on your
	// side, such as missing permissions or invalid configuration, so the first
	// attempt is the final one. Other batch failures are retried up to twice.
	//
	// Streaming materialization restarts a failed update indefinitely unless
	// the error cannot be retried. With this set, a streaming materialized
	// feature notifies only on errors that cannot be retried, and never on
	// failures that are restarted.
	FinalAttemptOnly bool `json:"final_attempt_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MaterializationFailureNotification) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MaterializationFailureNotification) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Notifications for the jobs and pipelines that materialize features, one field
// per trigger.
type MaterializationNotifications struct {
	// Who to notify when a run fails, and on which attempts.
	OnFailure *MaterializationFailureNotification `json:"on_failure,omitempty"`
	// Who to notify when a run succeeds.
	OnSuccess *MaterializationSuccessNotification `json:"on_success,omitempty"`
}

func (s *MaterializationNotifications) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Progress of a requested run, which may materialize multiple members of the
// set.
type MaterializationRunProgress struct {
	// The failure, including a failure to submit the run. Present when `state`
	// is `FAILED`.
	Error *DatabricksServiceExceptionWithDetailsProto `json:"error,omitempty"`
	// IDs of the materialized features assigned to this run when the operation
	// was created.
	MaterializedFeatureIds []string `json:"materialized_feature_ids,omitempty"`
	// Current state of the requested run.
	State MaterializationRunProgressState `json:"state,omitempty"`
}

func (s *MaterializationRunProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Lifecycle state of a requested materialization run.
type MaterializationRunProgressState string

const MaterializationRunProgressStateCancelled MaterializationRunProgressState = `CANCELLED`

const MaterializationRunProgressStateFailed MaterializationRunProgressState = `FAILED`

const MaterializationRunProgressStatePending MaterializationRunProgressState = `PENDING`

const MaterializationRunProgressStateRunning MaterializationRunProgressState = `RUNNING`

const MaterializationRunProgressStateSucceeded MaterializationRunProgressState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *MaterializationRunProgressState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MaterializationRunProgressState) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCEEDED`:
		*f = MaterializationRunProgressState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "PENDING", "RUNNING", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for MaterializationRunProgressState.
//
// There is no guarantee on the order of the values in the slice.
func (f *MaterializationRunProgressState) Values() []MaterializationRunProgressState {
	return []MaterializationRunProgressState{
		MaterializationRunProgressStateCancelled,
		MaterializationRunProgressStateFailed,
		MaterializationRunProgressStatePending,
		MaterializationRunProgressStateRunning,
		MaterializationRunProgressStateSucceeded,
	}
}

// Type always returns MaterializationRunProgressState to satisfy [pflag.Value] interface
func (f *MaterializationRunProgressState) Type() string {
	return "MaterializationRunProgressState"
}

// A group of materialized features created and managed as a single unit.
type MaterializationSet struct {
	// Server-assigned unique identifier for the materialization set.
	MaterializationSetId string `json:"materialization_set_id,omitempty"`
	// The materialized features in this set. Order is not significant. An
	// `UpdateMaterializationSet` request with `materialized_features` or `*` in
	// its update mask deletes omitted materialized features.
	MaterializedFeatures []MaterializedFeature `json:"materialized_features,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MaterializationSet) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MaterializationSet) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// What a partly failed DeleteMaterializationSet got through, attached to the
// error so a caller can tell which materialized feature members are gone and
// which ones still need attention. The set itself is left in place whenever any
// member failed, so a retry can still find the members.
type MaterializationSetDeletionProgress struct {
	// Members that are no longer materialized, including any that were already
	// gone.
	DeletedMaterializedFeatureIds []string `json:"deleted_materialized_feature_ids,omitempty"`
	// Members that could not be deleted, one entry each.
	FailedMaterializedFeatures []MaterializedFeatureChangeFailure `json:"failed_materialized_features,omitempty"`
}

func (s *MaterializationSetDeletionProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Progress of a materialization set operation.
type MaterializationSetOperationMetadata struct {
	// Time at which the operation was requested.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Set by a `DeleteMaterializationSet` operation.
	DeletionProgress *MaterializationSetDeletionProgress `json:"deletion_progress,omitempty"`
	// Full names of the features the set materializes, one entry per feature.
	FeatureFullNames []string `json:"feature_full_names,omitempty"`
	// ID of the materialization set the operation targets.
	MaterializationSetId string `json:"materialization_set_id,omitempty"`
	// Current state of the operation.
	State MaterializationSetOperationMetadataState `json:"state,omitempty"`
	// Set by a `TriggerMaterializationSet` operation.
	TriggerProgress *MaterializationSetTriggerProgress `json:"trigger_progress,omitempty"`
	// Set by an `UpdateMaterializationSet` operation.
	UpdateProgress *MaterializationSetUpdateProgress `json:"update_progress,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MaterializationSetOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MaterializationSetOperationMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Lifecycle state of a materialization set operation.
type MaterializationSetOperationMetadataState string

const MaterializationSetOperationMetadataStateCancelled MaterializationSetOperationMetadataState = `CANCELLED`

const MaterializationSetOperationMetadataStateFailed MaterializationSetOperationMetadataState = `FAILED`

const MaterializationSetOperationMetadataStatePending MaterializationSetOperationMetadataState = `PENDING`

const MaterializationSetOperationMetadataStateRunning MaterializationSetOperationMetadataState = `RUNNING`

const MaterializationSetOperationMetadataStateSucceeded MaterializationSetOperationMetadataState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *MaterializationSetOperationMetadataState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MaterializationSetOperationMetadataState) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCEEDED`:
		*f = MaterializationSetOperationMetadataState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "PENDING", "RUNNING", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for MaterializationSetOperationMetadataState.
//
// There is no guarantee on the order of the values in the slice.
func (f *MaterializationSetOperationMetadataState) Values() []MaterializationSetOperationMetadataState {
	return []MaterializationSetOperationMetadataState{
		MaterializationSetOperationMetadataStateCancelled,
		MaterializationSetOperationMetadataStateFailed,
		MaterializationSetOperationMetadataStatePending,
		MaterializationSetOperationMetadataStateRunning,
		MaterializationSetOperationMetadataStateSucceeded,
	}
}

// Type always returns MaterializationSetOperationMetadataState to satisfy [pflag.Value] interface
func (f *MaterializationSetOperationMetadataState) Type() string {
	return "MaterializationSetOperationMetadataState"
}

// Progress and outcomes of the runs requested by `TriggerMaterializationSet`.
// Entries remain available after the operation finishes.
//
// Once all requested runs have finished, the operation fails if any run failed.
// Otherwise, it is cancelled if any run was cancelled, and succeeds if all runs
// succeeded.
type MaterializationSetTriggerProgress struct {
	// One entry per requested run, including runs that fail or are cancelled
	// before submission.
	Runs []MaterializationRunProgress `json:"runs,omitempty"`
}

func (s *MaterializationSetTriggerProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Progress and outcomes of an `UpdateMaterializationSet` operation.
type MaterializationSetUpdateProgress struct {
	// IDs of materialized features that were deleted.
	DeletedMaterializedFeatureIds []string `json:"deleted_materialized_feature_ids,omitempty"`
	// Materialized features that could not be updated or deleted, one entry
	// each.
	FailedMaterializedFeatures []MaterializedFeatureChangeFailure `json:"failed_materialized_features,omitempty"`
	// IDs of materialized features whose mutable state was updated.
	UpdatedMaterializedFeatureIds []string `json:"updated_materialized_feature_ids,omitempty"`
}

func (s *MaterializationSetUpdateProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Recipients to notify when a materialization run succeeds.
type MaterializationSuccessNotification struct {
	// IDs of the notification destinations (for example Slack, Microsoft Teams,
	// PagerDuty, or a generic webhook) to notify. Not supported for streaming
	// materialized features.
	DestinationIds []string `json:"destination_ids,omitempty"`
	// Email addresses to notify.
	EmailAddresses []string `json:"email_addresses,omitempty"`
}

func (s *MaterializationSuccessNotification) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A materialized feature represents a feature that is continuously computed and
// stored.
type MaterializedFeature struct {
	// The ID of the budget policy used to attribute the serverless compute cost
	// of this materialization. If not specified, a default budget policy may be
	// applied.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`

	CronSchedule string `json:"cron_schedule,omitempty"`
	// A cron-based schedule trigger for batch materialization.
	CronScheduleTrigger *CronSchedule `json:"cron_schedule_trigger,omitempty"`
	// A trigger that disables automatic materialization, including the initial
	// run at creation. Runs are instead triggered through the
	// `TriggerMaterializationSet` API. Supported only for batch
	// materializations in a materialization set.
	ExternalTrigger *ExternalTrigger `json:"external_trigger,omitempty"`
	// The full name of the feature in Unity Catalog.
	FeatureName string `json:"feature_name"`
	// True if this is an online materialized feature. False if it is an offline
	// materialized feature.
	IsOnline bool `json:"is_online,omitempty"`
	// The ID of the job that materializes the feature. This is present for both
	// batch and streaming features.
	JobId int64 `json:"job_id,omitempty"`
	// The timestamp when the pipeline last ran and updated the materialized
	// feature values. If the pipeline has not run yet, this field will be null.
	LastMaterializationTime string `json:"last_materialization_time,omitempty"`
	// Name of the latest backfill operation on this materialized feature.
	// Format: operations/{operation_id}.
	LatestBackfillOperation string `json:"latest_backfill_operation,omitempty"`
	// The ID of the materialization set containing this materialized feature.
	// Unset for a materialized feature that is not part of a set.
	MaterializationSetId string `json:"materialization_set_id,omitempty"`
	// Server-assigned unique identifier for the materialized feature.
	MaterializedFeatureId string `json:"materialized_feature_id,omitempty"`
	// Notification configuration around the lifecycle of the job or pipeline
	// that materializes this feature. Features which are materialized in the
	// same pipeline will share the same notification.
	Notifications *MaterializationNotifications `json:"notifications,omitempty"`
	// Destination for writing feature values to an offline Delta table. The
	// resulting table is returned as `table_name`.
	OfflineStoreConfig *OfflineStoreConfig `json:"offline_store_config,omitempty"`
	// Destination for writing feature values to an online Lakebase table. The
	// resulting table is returned as `table_name`.
	OnlineStoreConfig *OnlineStoreConfig `json:"online_store_config,omitempty"`
	// Operations associated with this materialized feature.
	Operations *MaterializedFeatureOperations `json:"operations,omitempty"`
	// The ID of the pipeline that materializes this feature. This is only
	// present for streaming features.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Deprecated: This legacy pipeline setting does not control batch job
	// scheduling.
	PipelineScheduleState MaterializedFeaturePipelineScheduleState `json:"pipeline_schedule_state,omitempty"`
	// The Structured Streaming trigger mode used for materialization. Real-time
	// mode (RTM) targets sub-second latency for operational workloads;
	// micro-batch mode (MBM) favors cost efficiency for ETL and analytics
	// workloads.
	StreamingMode *StreamingMode `json:"streaming_mode,omitempty"`
	// The fully qualified Unity Catalog path to the table containing the
	// materialized feature (Delta table or Lakebase table). Output only.
	TableName string `json:"table_name,omitempty"`
	// A trigger that starts batch materialization when the upstream source
	// table changes.
	TableTrigger *TableTrigger `json:"table_trigger,omitempty"`
	// Custom tags to associate with this materialization. They are applied to
	// the materialization job (for batch features) or pipeline (for streaming
	// features) and forwarded to the underlying compute as cluster tags, so
	// materialization cost can be attributed in the billing system tables.
	// These tags apply only to the materialization compute; they are not
	// applied to the Unity Catalog Feature resource itself, whose tags are
	// managed separately through the Unity Catalog tagging API. A maximum of 25
	// tags is supported; keys and values are subject to the same limitations as
	// cluster tags.
	Tags map[string]string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MaterializedFeature) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MaterializedFeature) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Why one materialization set member could not be changed.
type MaterializedFeatureChangeFailure struct {
	// Databricks error code name of the failure, for example
	// "RESOURCE_CONFLICT".
	ErrorCode string `json:"error_code,omitempty"`
	// The member that could not be changed.
	MaterializedFeatureId string `json:"materialized_feature_id,omitempty"`
	// The failure's message.
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MaterializedFeatureChangeFailure) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MaterializedFeatureChangeFailure) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// References to operations on a materialized feature.
type MaterializedFeatureOperations struct {
	// Name of the latest entity purge operation. Format:
	// operations/{operation_id}.
	LatestEntityPurgeOperation string `json:"latest_entity_purge_operation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MaterializedFeatureOperations) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MaterializedFeatureOperations) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MaterializedFeaturePipelineScheduleState string

const MaterializedFeaturePipelineScheduleStateActive MaterializedFeaturePipelineScheduleState = `ACTIVE`

const MaterializedFeaturePipelineScheduleStatePaused MaterializedFeaturePipelineScheduleState = `PAUSED`

const MaterializedFeaturePipelineScheduleStateSnapshot MaterializedFeaturePipelineScheduleState = `SNAPSHOT`

// String representation for [fmt.Print]
func (f *MaterializedFeaturePipelineScheduleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MaterializedFeaturePipelineScheduleState) Set(v string) error {
	switch v {
	case `ACTIVE`, `PAUSED`, `SNAPSHOT`:
		*f = MaterializedFeaturePipelineScheduleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "PAUSED", "SNAPSHOT"`, v)
	}
}

// Values returns all possible values for MaterializedFeaturePipelineScheduleState.
//
// There is no guarantee on the order of the values in the slice.
func (f *MaterializedFeaturePipelineScheduleState) Values() []MaterializedFeaturePipelineScheduleState {
	return []MaterializedFeaturePipelineScheduleState{
		MaterializedFeaturePipelineScheduleStateActive,
		MaterializedFeaturePipelineScheduleStatePaused,
		MaterializedFeaturePipelineScheduleStateSnapshot,
	}
}

// Type always returns MaterializedFeaturePipelineScheduleState to satisfy [pflag.Value] interface
func (f *MaterializedFeaturePipelineScheduleState) Type() string {
	return "MaterializedFeaturePipelineScheduleState"
}

// Computes the maximum value.
type MaxFunction struct {
	// The input column from which the maximum is computed.
	Input string `json:"input"`
}

func (s *MaxFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Metric associated with a run, represented as a key-value pair.
type Metric struct {
	// The dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	DatasetName string `json:"dataset_name,omitempty"`
	// The key identifying the metric.
	Key string `json:"key,omitempty"`
	// The ID of the logged model or registered model version associated with
	// the metric, if applicable.
	ModelId string `json:"model_id,omitempty"`
	// The ID of the run containing the metric.
	RunId string `json:"run_id,omitempty"`
	// The step at which the metric was logged.
	Step int64 `json:"step,omitempty"`
	// Time at which the metric was recorded, as a Unix timestamp in
	// milliseconds.
	Timestamp int64 `json:"timestamp,omitempty"`
	// The value of the metric.
	Value float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Metric) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Metric) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Computes the minimum value.
type MinFunction struct {
	// The input column from which the minimum is computed.
	Input string `json:"input"`
}

func (s *MinFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Model struct {
	// Timestamp recorded when this `registered_model` was created.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Description of this `registered_model`.
	Description string `json:"description,omitempty"`
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	LatestVersions []ModelVersion `json:"latest_versions,omitempty"`
	// Unique name for the model.
	Name string `json:"name,omitempty"`
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	Tags []ModelTag `json:"tags,omitempty"`
	// User that created this `registered_model`
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Model) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Model) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ModelDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// User-specified description for the object.
	Description string `json:"description,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// Last update time of the object, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Array of model versions, each the latest version for its stage.
	LatestVersions []ModelVersion `json:"latest_versions,omitempty"`
	// Name of the model.
	Name string `json:"name,omitempty"`
	// Permission level granted for the requesting user on this registered model
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// Array of tags associated with the model.
	Tags []ModelTag `json:"tags,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelDatabricks) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelDatabricks) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Represents a LoggedModel or Registered Model Version input to a Run.
type ModelInput struct {
	// The unique identifier of the model.
	ModelId string `json:"model_id"`
}

func (s *ModelInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Represents a LoggedModel output of a Run.
type ModelOutput struct {
	// The unique identifier of the model.
	ModelId string `json:"model_id"`
	// The step at which the model was produced.
	Step int64 `json:"step"`
}

func (s *ModelOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Tag for a registered model
type ModelTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A version of a registered model in the Workspace Model Registry. Each version
// points to the model artifacts it was registered from and moves through the
// stages `None`, `Staging`, `Production` and `Archived`.
type ModelVersion struct {
	// Time when the model version was created, as a Unix timestamp in
	// milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Current stage of the model version. Valid values are `None`, `Staging`,
	// `Production` and `Archived`. A new model version starts in `None`.
	CurrentStage string `json:"current_stage,omitempty"`
	// User-provided description of the model version.
	Description string `json:"description,omitempty"`
	// Time when the metadata of the model version was last updated, as a Unix
	// timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Name of the registered model that the version belongs to.
	Name string `json:"name,omitempty"`
	// ID of the MLflow run that produced the model artifacts, if `source` comes
	// from a run in the MLflow tracking server.
	RunId string `json:"run_id,omitempty"`
	// URL of the MLflow run that generated the model version, as given in the
	// create request. It is typically set when the run is hosted on a different
	// tracking server than the registry.
	RunLink string `json:"run_link,omitempty"`
	// URI of the model artifacts that the model version was registered from, as
	// given in the create request.
	Source string `json:"source,omitempty"`
	// Registration status of the model version. The version can be used once
	// the status is `READY`.
	Status ModelVersionStatus `json:"status,omitempty"`
	// Details on the current `status`, such as the reason why registration
	// failed.
	StatusMessage string `json:"status_message,omitempty"`
	// Key-value tags attached to the model version.
	Tags []ModelVersionTag `json:"tags,omitempty"`
	// User that created this `model_version`.
	UserId string `json:"user_id,omitempty"`
	// Version number of the model version, as a string. Version numbers start
	// at 1 for each registered model.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelVersion) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelVersion) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ModelVersionDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Current stage of the model version. Valid values are `None`, `Staging`,
	// `Production` and `Archived`.
	CurrentStage string `json:"current_stage,omitempty"`
	// User-specified description for the object.
	Description string `json:"description,omitempty"`
	// Email Subscription Status: This is the subscription status of the user to
	// the model version Users get subscribed by interacting with the model
	// version.
	EmailSubscriptionStatus RegistryEmailSubscriptionType `json:"email_subscription_status,omitempty"`
	// Feature lineage of `model_version`.
	FeatureList *FeatureList `json:"feature_list,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Name of the model.
	Name string `json:"name,omitempty"`
	// Open stage transition requests for the model version.
	OpenRequests []Activity `json:"open_requests,omitempty"`
	// Permission level of the requesting user on the registered model that the
	// version belongs to.
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	RunId string `json:"run_id,omitempty"`
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	RunLink string `json:"run_link,omitempty"`
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	Source string `json:"source,omitempty"`
	// Registration status of the model version. The version can be used once
	// the status is `READY`.
	Status Status `json:"status,omitempty"`
	// Details on the current status, for example why registration failed.
	StatusMessage string `json:"status_message,omitempty"`
	// Array of tags that are associated with the model version.
	Tags []ModelVersionTag `json:"tags,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`
	// Version of the model.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelVersionDatabricks) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelVersionDatabricks) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks.
//
// * `FAILED_REGISTRATION`: Request to register a new model version has failed.
//
// * `READY`: Model version is ready for use.
type ModelVersionStatus string

// Request to register a new model version has failed.
const ModelVersionStatusFailedRegistration ModelVersionStatus = `FAILED_REGISTRATION`

// Request to register a new model version is pending as server performs
// background tasks.
const ModelVersionStatusPendingRegistration ModelVersionStatus = `PENDING_REGISTRATION`

// Model version is ready for use.
const ModelVersionStatusReady ModelVersionStatus = `READY`

// String representation for [fmt.Print]
func (f *ModelVersionStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelVersionStatus) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = ModelVersionStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Values returns all possible values for ModelVersionStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ModelVersionStatus) Values() []ModelVersionStatus {
	return []ModelVersionStatus{
		ModelVersionStatusFailedRegistration,
		ModelVersionStatusPendingRegistration,
		ModelVersionStatusReady,
	}
}

// Type always returns ModelVersionStatus to satisfy [pflag.Value] interface
func (f *ModelVersionStatus) Type() string {
	return "ModelVersionStatus"
}

type ModelVersionTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelVersionTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelVersionTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Mutual-TLS (mTLS) authentication configuration. The keystore (client
// certificate + private key) and truststore (CAs trusted to verify the broker)
// live as JKS files on Unity Catalog volumes, with their passwords stored in
// Databricks secret scopes. This matches the SSL setup pattern documented at
// https://docs.databricks.com/en/connect/streaming/kafka/authentication#use-ssl-to-connect-databricks-to-kafka.
type MtlsConfig struct {
	// Set to true only when the broker certificate's SAN intentionally does not
	// match the connection endpoint — for example when reaching the cluster
	// through a PrivateLink endpoint whose DNS name is not in the broker
	// certificate. Skipping the hostname check removes a defense against
	// man-in-the-middle attacks; do not enable casually. mTLS client
	// authentication is unaffected by this option.
	//
	// See the Apache Kafka SSL security guide for background on this check:
	// https://kafka.apache.org/42/security/encryption-and-authentication-using-ssl/#host-name-verification
	DisableHostnameVerification bool `json:"disable_hostname_verification,omitempty"`
	// Secret-scope reference for the private key password. Often the same value
	// as the keystore password (keytool's default), but provided as a separate
	// field because Apache Kafka requires it as a distinct option
	// (kafka.ssl.key.password).
	KeyPasswordRef SecretScopeReference `json:"key_password_ref"`
	// Unity Catalog volume path to the JKS keystore file containing the client
	// certificate and private key. e.g.
	// "/Volumes/<catalog>/<schema>/<volume>/client.jks". The materialization
	// compute must have read permission on this volume.
	KeystoreLocation string `json:"keystore_location"`
	// Secret-scope reference for the JKS keystore password.
	KeystorePasswordRef SecretScopeReference `json:"keystore_password_ref"`
	// Unity Catalog volume path to the JKS truststore file containing the CA
	// certificate(s) trusted to verify the Kafka broker's server certificate.
	// e.g. "/Volumes/<catalog>/<schema>/<volume>/truststore.jks".
	TruststoreLocation string `json:"truststore_location"`
	// Secret-scope reference for the JKS truststore password.
	TruststorePasswordRef SecretScopeReference `json:"truststore_password_ref"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MtlsConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MtlsConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration for offline store destination.
type OfflineStoreConfig struct {
	// The Unity Catalog catalog name.
	CatalogName string `json:"catalog_name"`
	// The Unity Catalog schema name.
	SchemaName string `json:"schema_name"`
	// Prefix for Unity Catalog table name. The materialized feature will be
	// stored in a table with this prefix and a generated postfix.
	TableNamePrefix string `json:"table_name_prefix"`
}

func (s *OfflineStoreConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An OnlineStore is a logical database instance that stores and serves features
// online.
type OnlineStore struct {
	// The capacity of the online store. Valid values are "CU_1", "CU_2",
	// "CU_4", "CU_8".
	Capacity string `json:"capacity"`
	// The timestamp when the online store was created.
	CreationTime string `json:"creation_time,omitempty"`
	// The email of the creator of the online store.
	Creator string `json:"creator,omitempty"`
	// The name of the online store. This is the unique identifier for the
	// online store.
	Name string `json:"name"`
	// The number of read replicas for the online store. Defaults to 0.
	ReadReplicaCount int `json:"read_replica_count,omitempty"`
	// The current state of the online store.
	State OnlineStoreState `json:"state,omitempty"`
	// The usage policy applied to the online store to track billing.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OnlineStore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineStore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration for online store destination.
type OnlineStoreConfig struct {
	// The Unity Catalog catalog name. This name is also used as the Lakebase
	// logical database name. Quoting is handled by the backend where needed, do
	// not pre-quote it.
	CatalogName string `json:"catalog_name"`
	// The name of the target online store.
	OnlineStoreName string `json:"online_store_name"`
	// The Unity Catalog schema name. This name is also used as the Lakebase
	// schema name under the database. Quoting is handled by the backend where
	// needed, do not pre-quote it.
	SchemaName string `json:"schema_name"`
	// Prefix for Unity Catalog table name. The materialized feature will be
	// stored in a Lakebase table with this prefix and a generated postfix.
	TableNamePrefix string `json:"table_name_prefix"`
}

func (s *OnlineStoreConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type OnlineStoreState string

const OnlineStoreStateAvailable OnlineStoreState = `AVAILABLE`

const OnlineStoreStateDeleting OnlineStoreState = `DELETING`

const OnlineStoreStateFailingOver OnlineStoreState = `FAILING_OVER`

const OnlineStoreStateStarting OnlineStoreState = `STARTING`

const OnlineStoreStateStopped OnlineStoreState = `STOPPED`

const OnlineStoreStateUpdating OnlineStoreState = `UPDATING`

// String representation for [fmt.Print]
func (f *OnlineStoreState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OnlineStoreState) Set(v string) error {
	switch v {
	case `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`:
		*f = OnlineStoreState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVAILABLE", "DELETING", "FAILING_OVER", "STARTING", "STOPPED", "UPDATING"`, v)
	}
}

// Values returns all possible values for OnlineStoreState.
//
// There is no guarantee on the order of the values in the slice.
func (f *OnlineStoreState) Values() []OnlineStoreState {
	return []OnlineStoreState{
		OnlineStoreStateAvailable,
		OnlineStoreStateDeleting,
		OnlineStoreStateFailingOver,
		OnlineStoreStateStarting,
		OnlineStoreStateStopped,
		OnlineStoreStateUpdating,
	}
}

// Type always returns OnlineStoreState to satisfy [pflag.Value] interface
func (f *OnlineStoreState) Type() string {
	return "OnlineStoreState"
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `json:"done,omitempty"`
	// The error result of the operation in case of failure or cancellation.
	Error *DatabricksServiceExceptionWithDetailsProto `json:"error,omitempty"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata json.RawMessage `json:"metadata,omitempty"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name string `json:"name,omitempty"`
	// The normal, successful response of the operation.
	Response json.RawMessage `json:"response,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Operation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Operation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Param associated with a run.
type Param struct {
	// Key identifying this param.
	Key string `json:"key,omitempty"`
	// Value associated with this param.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Param) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Param) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions].
//
// [MLflow Model permissions]: https://docs.databricks.com/aws/en/machine-learning/manage-model-lifecycle/workspace-model-registry#permissions
type PermissionLevel string

const PermissionLevelCanCreateRegisteredModel PermissionLevel = `CAN_CREATE_REGISTERED_MODEL`

const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageStagingVersions PermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanRead PermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *PermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PermissionLevel) Set(v string) error {
	switch v {
	case `CAN_CREATE_REGISTERED_MODEL`, `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_CREATE_REGISTERED_MODEL", "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Values returns all possible values for PermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PermissionLevel) Values() []PermissionLevel {
	return []PermissionLevel{
		PermissionLevelCanCreateRegisteredModel,
		PermissionLevelCanEdit,
		PermissionLevelCanManage,
		PermissionLevelCanManageProductionVersions,
		PermissionLevelCanManageStagingVersions,
		PermissionLevelCanRead,
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

// Pip requirements to install in a serving container.
type PipRequirements struct {
	// Pip requirement strings (PEP 508) or UC Volume whl paths (/Volumes/...).
	Requirements []string `json:"requirements"`
}

func (s *PipRequirements) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A Protocol Buffer schema paired with the name of the message within it that
// describes the Kafka payload. A .proto file may declare multiple messages;
// message_name disambiguates.
type ProtoSchemaSpec struct {
	// The fully-qualified name of the message within schema_text that describes
	// the Kafka payload (e.g. "Event" or "com.example.Event" if schema_text
	// declares a package). Identifies which message is used to decode each
	// Kafka record — a .proto file may declare multiple messages but only one
	// represents the payload. Must not be empty.
	MessageName string `json:"message_name"`
	// The raw .proto file text (proto2 and proto3 syntax supported, see
	// https://protobuf.dev/programming-guides/proto3/ and
	// https://protobuf.dev/programming-guides/proto2/).
	SchemaText string `json:"schema_text"`
}

func (s *ProtoSchemaSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PublishSpec struct {
	// Budget policy id used to attribute the serverless compute cost of the
	// synced online-table sync pipeline. Applied only when the sync pipeline is
	// first created (the initial publish of a new online table); republishing
	// to an existing online table does not update it.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Full Unity Catalog name of one of the features materialized in the source
	// table, used to derive the synced online table's entity and timeseries
	// columns. Required for view sources without a UC PrimaryKeyConstraint;
	// ignored when the source already has one.
	FullFeatureName string `json:"full_feature_name,omitempty"`
	// The name of the target online store.
	OnlineStore string `json:"online_store"`
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName string `json:"online_table_name"`
	// The publish mode of the pipeline that syncs the online table with the
	// source table.
	PublishMode PublishSpecPublishMode `json:"publish_mode"`
	// Custom tags to apply to the synced online-table sync pipeline created for
	// this publish. They are forwarded to the pipeline's compute as cluster
	// tags so its cost can be attributed in the billing system tables. Applied
	// only when the sync pipeline is first created (the initial publish of a
	// new online table); republishing to an existing online table does not
	// update them.
	Tags map[string]string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PublishSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublishSpecPublishMode string

const PublishSpecPublishModeContinuous PublishSpecPublishMode = `CONTINUOUS`

const PublishSpecPublishModeSnapshot PublishSpecPublishMode = `SNAPSHOT`

const PublishSpecPublishModeTriggered PublishSpecPublishMode = `TRIGGERED`

// String representation for [fmt.Print]
func (f *PublishSpecPublishMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PublishSpecPublishMode) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `SNAPSHOT`, `TRIGGERED`:
		*f = PublishSpecPublishMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "SNAPSHOT", "TRIGGERED"`, v)
	}
}

// Values returns all possible values for PublishSpecPublishMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *PublishSpecPublishMode) Values() []PublishSpecPublishMode {
	return []PublishSpecPublishMode{
		PublishSpecPublishModeContinuous,
		PublishSpecPublishModeSnapshot,
		PublishSpecPublishModeTriggered,
	}
}

// Type always returns PublishSpecPublishMode to satisfy [pflag.Value] interface
func (f *PublishSpecPublishMode) Type() string {
	return "PublishSpecPublishMode"
}

type PublishTableRequest struct {
	// The specification for publishing the online table from the source table.
	PublishSpec PublishSpec `json:"publish_spec"`
	// The full three-part (catalog, schema, table) name of the source table.
	SourceTableName string `json:"-" url:"-"`
}

func (s *PublishTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PublishTableResponse struct {
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName string `json:"online_table_name,omitempty"`
	// The ID of the pipeline that syncs the online table with the source table.
	PipelineId string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PublishTableResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishTableResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Progress and configuration for a feature entity purge.
type PurgeFeatureEntitiesMetadata struct {
	// Time at which the purge operation was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Whether this purge deletes matching rows instead of setting feature
	// values to null.
	DeleteRows bool `json:"delete_rows,omitempty"`
	// Fully qualified name of the Unity Catalog Delta table containing the
	// entity keys to purge.
	EntitiesTable string `json:"entities_table,omitempty"`
	// Version of the entities table used by the purge.
	EntitiesTableVersion string `json:"entities_table_version,omitempty"`
	// Fully qualified names of the features targeted by the purge.
	Features []string `json:"features,omitempty"`
	// ID of the job that executes this purge.
	JobId int64 `json:"job_id,omitempty"`
	// Current state of the purge operation.
	State PurgeFeatureEntitiesMetadataState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PurgeFeatureEntitiesMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PurgeFeatureEntitiesMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Lifecycle state of a feature entity purge.
type PurgeFeatureEntitiesMetadataState string

const PurgeFeatureEntitiesMetadataStateCancelled PurgeFeatureEntitiesMetadataState = `CANCELLED`

const PurgeFeatureEntitiesMetadataStateFailed PurgeFeatureEntitiesMetadataState = `FAILED`

const PurgeFeatureEntitiesMetadataStatePending PurgeFeatureEntitiesMetadataState = `PENDING`

const PurgeFeatureEntitiesMetadataStatePendingCleanup PurgeFeatureEntitiesMetadataState = `PENDING_CLEANUP`

const PurgeFeatureEntitiesMetadataStatePendingPurge PurgeFeatureEntitiesMetadataState = `PENDING_PURGE`

const PurgeFeatureEntitiesMetadataStateRunning PurgeFeatureEntitiesMetadataState = `RUNNING`

const PurgeFeatureEntitiesMetadataStateRunningCleanup PurgeFeatureEntitiesMetadataState = `RUNNING_CLEANUP`

const PurgeFeatureEntitiesMetadataStateRunningPurge PurgeFeatureEntitiesMetadataState = `RUNNING_PURGE`

const PurgeFeatureEntitiesMetadataStateSucceeded PurgeFeatureEntitiesMetadataState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *PurgeFeatureEntitiesMetadataState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PurgeFeatureEntitiesMetadataState) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `PENDING`, `PENDING_CLEANUP`, `PENDING_PURGE`, `RUNNING`, `RUNNING_CLEANUP`, `RUNNING_PURGE`, `SUCCEEDED`:
		*f = PurgeFeatureEntitiesMetadataState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "PENDING", "PENDING_CLEANUP", "PENDING_PURGE", "RUNNING", "RUNNING_CLEANUP", "RUNNING_PURGE", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for PurgeFeatureEntitiesMetadataState.
//
// There is no guarantee on the order of the values in the slice.
func (f *PurgeFeatureEntitiesMetadataState) Values() []PurgeFeatureEntitiesMetadataState {
	return []PurgeFeatureEntitiesMetadataState{
		PurgeFeatureEntitiesMetadataStateCancelled,
		PurgeFeatureEntitiesMetadataStateFailed,
		PurgeFeatureEntitiesMetadataStatePending,
		PurgeFeatureEntitiesMetadataStatePendingCleanup,
		PurgeFeatureEntitiesMetadataStatePendingPurge,
		PurgeFeatureEntitiesMetadataStateRunning,
		PurgeFeatureEntitiesMetadataStateRunningCleanup,
		PurgeFeatureEntitiesMetadataStateRunningPurge,
		PurgeFeatureEntitiesMetadataStateSucceeded,
	}
}

// Type always returns PurgeFeatureEntitiesMetadataState to satisfy [pflag.Value] interface
func (f *PurgeFeatureEntitiesMetadataState) Type() string {
	return "PurgeFeatureEntitiesMetadataState"
}

// Request to purge materialized feature values for entities listed in a Unity
// Catalog Delta table.
type PurgeFeatureEntitiesRequest struct {
	// The budget policy ID, in UUID format, used to attribute the serverless
	// compute cost of this purge. If not specified, a default budget policy may
	// be applied.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// If false feature values are set to null on purge but rows are not
	// deleted. If true then the full rows containing entities are deleted.
	// Defaults to false. When true, the request must include every feature
	// which shares the same `table_name` property in their MaterializedFeature.
	// Features from multiple tables may be included if all features
	// materialized in each table are included.
	DeleteRows bool `json:"delete_rows,omitempty"`
	// Fully qualified name of the Unity Catalog Delta table containing the
	// entity keys to purge. The table may contain a subset of each feature's
	// entity-key columns. A partial key match deletes all feature rows matching
	// the provided key values. Non-key columns are rejected; null key values
	// are allowed.
	EntitiesTable string `json:"entities_table"`
	// Fully qualified names of the features to purge. At least one nonempty
	// feature name is required. A request may contain at most 500 features;
	// submit additional features in separate requests. Duplicate features are
	// rejected.
	Features []string `json:"features"`
	// Optional UUID4 idempotency token for the request.
	RequestId string `json:"request_id,omitempty"`
	// Custom tags to associate with this purge. They are applied to the purge
	// job and forwarded to the underlying compute as Databricks resource tags,
	// so purge cost can be attributed in the billing system tables. These tags
	// apply only to the purge compute; they are not applied to the Unity
	// Catalog Feature resources themselves, whose tags are managed separately
	// through the Unity Catalog tagging API. A maximum of 25 tags is supported;
	// keys and values are subject to the same limitations as Databricks
	// resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PurgeFeatureEntitiesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PurgeFeatureEntitiesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Result of a completed feature entity purge.
type PurgeFeatureEntitiesResponse struct {
	// Operation-level error, if the purge failed outside an individual feature
	// target.
	Error *DatabricksServiceExceptionWithDetailsProto `json:"error,omitempty"`
	// Metadata about the purge operation.
	Metadata *PurgeFeatureEntitiesMetadata `json:"metadata,omitempty"`
	// Per-feature purge results.
	Results []PurgeFeatureEntitiesResult `json:"results,omitempty"`
	// State of the purge operation.
	State PurgeFeatureEntitiesMetadataState `json:"state,omitempty"`
}

func (s *PurgeFeatureEntitiesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Result of purging one feature.
type PurgeFeatureEntitiesResult struct {
	// Error encountered while purging this feature, if any.
	Error *DatabricksServiceExceptionWithDetailsProto `json:"error,omitempty"`
	// Fully qualified name of the feature that was purged.
	Feature string `json:"feature,omitempty"`
	// State of the offline purge for this feature.
	OfflineState PurgeFeatureEntitiesResultState `json:"offline_state,omitempty"`
	// State of the online purge for this feature.
	OnlineState PurgeFeatureEntitiesResultState `json:"online_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PurgeFeatureEntitiesResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PurgeFeatureEntitiesResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Terminal state of a purge for one store type.
type PurgeFeatureEntitiesResultState string

const PurgeFeatureEntitiesResultStateFailed PurgeFeatureEntitiesResultState = `FAILED`

const PurgeFeatureEntitiesResultStateNotApplicable PurgeFeatureEntitiesResultState = `NOT_APPLICABLE`

const PurgeFeatureEntitiesResultStateSucceeded PurgeFeatureEntitiesResultState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *PurgeFeatureEntitiesResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PurgeFeatureEntitiesResultState) Set(v string) error {
	switch v {
	case `FAILED`, `NOT_APPLICABLE`, `SUCCEEDED`:
		*f = PurgeFeatureEntitiesResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "NOT_APPLICABLE", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for PurgeFeatureEntitiesResultState.
//
// There is no guarantee on the order of the values in the slice.
func (f *PurgeFeatureEntitiesResultState) Values() []PurgeFeatureEntitiesResultState {
	return []PurgeFeatureEntitiesResultState{
		PurgeFeatureEntitiesResultStateFailed,
		PurgeFeatureEntitiesResultStateNotApplicable,
		PurgeFeatureEntitiesResultStateSucceeded,
	}
}

// Type always returns PurgeFeatureEntitiesResultState to satisfy [pflag.Value] interface
func (f *PurgeFeatureEntitiesResultState) Type() string {
	return "PurgeFeatureEntitiesResultState"
}

type RegisteredModelAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelAccessControlResponse struct {
	// All permissions.
	AllPermissions []RegisteredModelPermission `json:"all_permissions,omitempty"`
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

func (s *RegisteredModelAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type RegisteredModelPermissionLevel string

const RegisteredModelPermissionLevelCanEdit RegisteredModelPermissionLevel = `CAN_EDIT`

const RegisteredModelPermissionLevelCanManage RegisteredModelPermissionLevel = `CAN_MANAGE`

const RegisteredModelPermissionLevelCanManageProductionVersions RegisteredModelPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const RegisteredModelPermissionLevelCanManageStagingVersions RegisteredModelPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const RegisteredModelPermissionLevelCanRead RegisteredModelPermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *RegisteredModelPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegisteredModelPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*f = RegisteredModelPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Values returns all possible values for RegisteredModelPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegisteredModelPermissionLevel) Values() []RegisteredModelPermissionLevel {
	return []RegisteredModelPermissionLevel{
		RegisteredModelPermissionLevelCanEdit,
		RegisteredModelPermissionLevelCanManage,
		RegisteredModelPermissionLevelCanManageProductionVersions,
		RegisteredModelPermissionLevelCanManageStagingVersions,
		RegisteredModelPermissionLevelCanRead,
	}
}

// Type always returns RegisteredModelPermissionLevel to satisfy [pflag.Value] interface
func (f *RegisteredModelPermissionLevel) Type() string {
	return "RegisteredModelPermissionLevel"
}

type RegisteredModelPermissions struct {
	AccessControlList []RegisteredModelAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegisteredModelPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegisteredModelPermissionsRequest struct {
	AccessControlList []RegisteredModelAccessControlRequest `json:"access_control_list,omitempty"`
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" url:"-"`
}

func (s *RegisteredModelPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// .. note:: Experimental: This entity may change or be removed in a future
// release without warning. Email subscription types for registry notifications:
// - `ALL_EVENTS`: Subscribed to all events. - `DEFAULT`: Default subscription
// type. - `SUBSCRIBED`: Subscribed to notifications. - `UNSUBSCRIBED`: Not
// subscribed to notifications.
type RegistryEmailSubscriptionType string

// Subscribed to all events.
const RegistryEmailSubscriptionTypeAllEvents RegistryEmailSubscriptionType = `ALL_EVENTS`

// Default subscription type.
const RegistryEmailSubscriptionTypeDefault RegistryEmailSubscriptionType = `DEFAULT`

// Subscribed to notifications.
const RegistryEmailSubscriptionTypeSubscribed RegistryEmailSubscriptionType = `SUBSCRIBED`

// Not subscribed to notifications.
const RegistryEmailSubscriptionTypeUnsubscribed RegistryEmailSubscriptionType = `UNSUBSCRIBED`

// String representation for [fmt.Print]
func (f *RegistryEmailSubscriptionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryEmailSubscriptionType) Set(v string) error {
	switch v {
	case `ALL_EVENTS`, `DEFAULT`, `SUBSCRIBED`, `UNSUBSCRIBED`:
		*f = RegistryEmailSubscriptionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_EVENTS", "DEFAULT", "SUBSCRIBED", "UNSUBSCRIBED"`, v)
	}
}

// Values returns all possible values for RegistryEmailSubscriptionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegistryEmailSubscriptionType) Values() []RegistryEmailSubscriptionType {
	return []RegistryEmailSubscriptionType{
		RegistryEmailSubscriptionTypeAllEvents,
		RegistryEmailSubscriptionTypeDefault,
		RegistryEmailSubscriptionTypeSubscribed,
		RegistryEmailSubscriptionTypeUnsubscribed,
	}
}

// Type always returns RegistryEmailSubscriptionType to satisfy [pflag.Value] interface
func (f *RegistryEmailSubscriptionType) Type() string {
	return "RegistryEmailSubscriptionType"
}

type RegistryWebhook struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events []RegistryWebhookEvent `json:"events,omitempty"`

	HttpUrlSpec *HttpUrlSpecWithoutSecret `json:"http_url_spec,omitempty"`
	// Webhook ID
	Id string `json:"id,omitempty"`

	JobSpec *JobSpecWithoutSecret `json:"job_spec,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Name of the model whose events would trigger this webhook.
	ModelName string `json:"model_name,omitempty"`

	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RegistryWebhook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegistryWebhook) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RegistryWebhookEvent string

const RegistryWebhookEventCommentCreated RegistryWebhookEvent = `COMMENT_CREATED`

const RegistryWebhookEventModelVersionCreated RegistryWebhookEvent = `MODEL_VERSION_CREATED`

const RegistryWebhookEventModelVersionTagSet RegistryWebhookEvent = `MODEL_VERSION_TAG_SET`

const RegistryWebhookEventModelVersionTransitionedStage RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_STAGE`

const RegistryWebhookEventModelVersionTransitionedToArchived RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`

const RegistryWebhookEventModelVersionTransitionedToProduction RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`

const RegistryWebhookEventModelVersionTransitionedToStaging RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_STAGING`

const RegistryWebhookEventRegisteredModelCreated RegistryWebhookEvent = `REGISTERED_MODEL_CREATED`

const RegistryWebhookEventTransitionRequestCreated RegistryWebhookEvent = `TRANSITION_REQUEST_CREATED`

const RegistryWebhookEventTransitionRequestToArchivedCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`

const RegistryWebhookEventTransitionRequestToProductionCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`

const RegistryWebhookEventTransitionRequestToStagingCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_STAGING_CREATED`

// String representation for [fmt.Print]
func (f *RegistryWebhookEvent) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryWebhookEvent) Set(v string) error {
	switch v {
	case `COMMENT_CREATED`, `MODEL_VERSION_CREATED`, `MODEL_VERSION_TAG_SET`, `MODEL_VERSION_TRANSITIONED_STAGE`, `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`, `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`, `MODEL_VERSION_TRANSITIONED_TO_STAGING`, `REGISTERED_MODEL_CREATED`, `TRANSITION_REQUEST_CREATED`, `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`, `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`, `TRANSITION_REQUEST_TO_STAGING_CREATED`:
		*f = RegistryWebhookEvent(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COMMENT_CREATED", "MODEL_VERSION_CREATED", "MODEL_VERSION_TAG_SET", "MODEL_VERSION_TRANSITIONED_STAGE", "MODEL_VERSION_TRANSITIONED_TO_ARCHIVED", "MODEL_VERSION_TRANSITIONED_TO_PRODUCTION", "MODEL_VERSION_TRANSITIONED_TO_STAGING", "REGISTERED_MODEL_CREATED", "TRANSITION_REQUEST_CREATED", "TRANSITION_REQUEST_TO_ARCHIVED_CREATED", "TRANSITION_REQUEST_TO_PRODUCTION_CREATED", "TRANSITION_REQUEST_TO_STAGING_CREATED"`, v)
	}
}

// Values returns all possible values for RegistryWebhookEvent.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegistryWebhookEvent) Values() []RegistryWebhookEvent {
	return []RegistryWebhookEvent{
		RegistryWebhookEventCommentCreated,
		RegistryWebhookEventModelVersionCreated,
		RegistryWebhookEventModelVersionTagSet,
		RegistryWebhookEventModelVersionTransitionedStage,
		RegistryWebhookEventModelVersionTransitionedToArchived,
		RegistryWebhookEventModelVersionTransitionedToProduction,
		RegistryWebhookEventModelVersionTransitionedToStaging,
		RegistryWebhookEventRegisteredModelCreated,
		RegistryWebhookEventTransitionRequestCreated,
		RegistryWebhookEventTransitionRequestToArchivedCreated,
		RegistryWebhookEventTransitionRequestToProductionCreated,
		RegistryWebhookEventTransitionRequestToStagingCreated,
	}
}

// Type always returns RegistryWebhookEvent to satisfy [pflag.Value] interface
func (f *RegistryWebhookEvent) Type() string {
	return "RegistryWebhookEvent"
}

// Enable or disable triggering the webhook, or put the webhook into test mode.
// The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an associated
// event happens.
//
// * `DISABLED`: Webhook is not triggered.
//
// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is not
// triggered on a real event.
type RegistryWebhookStatus string

// Webhook is triggered when an associated event happens.
const RegistryWebhookStatusActive RegistryWebhookStatus = `ACTIVE`

// Webhook is not triggered.
const RegistryWebhookStatusDisabled RegistryWebhookStatus = `DISABLED`

// Webhook can be triggered through the test endpoint, but is not triggered on a
// real event.
const RegistryWebhookStatusTestMode RegistryWebhookStatus = `TEST_MODE`

// String representation for [fmt.Print]
func (f *RegistryWebhookStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryWebhookStatus) Set(v string) error {
	switch v {
	case `ACTIVE`, `DISABLED`, `TEST_MODE`:
		*f = RegistryWebhookStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DISABLED", "TEST_MODE"`, v)
	}
}

// Values returns all possible values for RegistryWebhookStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegistryWebhookStatus) Values() []RegistryWebhookStatus {
	return []RegistryWebhookStatus{
		RegistryWebhookStatusActive,
		RegistryWebhookStatusDisabled,
		RegistryWebhookStatusTestMode,
	}
}

// Type always returns RegistryWebhookStatus to satisfy [pflag.Value] interface
func (f *RegistryWebhookStatus) Type() string {
	return "RegistryWebhookStatus"
}

// Details required to identify and reject a model version stage transition
// request.
type RejectTransitionRequest struct {
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RejectTransitionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RejectTransitionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RejectTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	Activity *Activity `json:"activity,omitempty"`
}

func (s *RejectTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RenameModelRequest struct {
	// Registered model unique name identifier.
	Name string `json:"name"`
	// If provided, updates the name for this `registered_model`.
	NewName string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RenameModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RenameModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RenameModelResponse struct {
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

func (s *RenameModelResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An output name for a looked-up feature.
type RenameOutput struct {
	// Name of the feature in the feature table.
	Feature string `json:"feature"`
	// Output name for the feature.
	Output string `json:"output"`
}

func (s *RenameOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A request-time data source whose value is provided at inference time: offline
// batch scoring or online serving endpoint
type RequestSource struct {
	// A schema containing scalar or nested fields, in Spark StructType JSON
	// format (from df.schema.json()). This preserves field, array-element, and
	// map-value nullability.
	DataframeSchema string `json:"dataframe_schema,omitempty"`
	// A flat schema with scalar-typed fields only.
	FlatSchema *FlatSchema `json:"flat_schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RequestSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RequestSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
}

func (s *RestoreExperiment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RestoreRun struct {
	// ID of the run to restore.
	RunId string `json:"run_id"`
}

func (s *RestoreRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RestoreRuns struct {
	// The ID of the experiment containing the runs to restore.
	ExperimentId string `json:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	MaxRuns int `json:"max_runs,omitempty"`
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	MinTimestampMillis int64 `json:"min_timestamp_millis"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestoreRuns) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestoreRuns) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RestoreRunsResponse struct {
	// The number of runs restored.
	RunsRestored int `json:"runs_restored,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestoreRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestoreRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A rolling time window with an optional non-negative delay.
type RollingWindow struct {
	// Non-negative analytic lag that evaluates the window this far in the past.
	// Use this for timing variations unrelated to source lateness, such as a
	// 30-day count as of one week ago. If unset, the analytic lag is zero. It
	// composes with source.lateness when both are set.
	Delay *duration.Duration `json:"delay,omitempty"`
	// The duration of the rolling window. Must be positive when set; absent
	// means lifetime (aggregate over the entity's entire history).
	WindowDuration *duration.Duration `json:"window_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RollingWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RollingWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A single run.
type Run struct {
	// Run data.
	Data *RunData `json:"data,omitempty"`
	// Run metadata.
	Info *RunInfo `json:"info,omitempty"`
	// Run inputs.
	Inputs *RunInputs `json:"inputs,omitempty"`
}

func (s *Run) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Run data (metrics, params, and tags).
type RunData struct {
	// Run metrics.
	Metrics []Metric `json:"metrics,omitempty"`
	// Run parameters.
	Params []Param `json:"params,omitempty"`
	// Additional metadata key-value pairs.
	Tags []RunTag `json:"tags,omitempty"`
}

func (s *RunData) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Metadata of a single run.
type RunInfo struct {
	// URI of the directory where the run's artifacts are stored. MLflow sets it
	// when the run is created, as the `artifact_location` of the run's
	// experiment followed by `/<run_id>/artifacts`.
	ArtifactUri string `json:"artifact_uri,omitempty"`
	// Unix timestamp of when the run ended in milliseconds.
	EndTime int64 `json:"end_time,omitempty"`
	// ID of the experiment that the run belongs to.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Current life cycle stage of the run: `active` or `deleted`.
	LifecycleStage string `json:"lifecycle_stage,omitempty"`
	// Unique identifier for the run.
	RunId string `json:"run_id,omitempty"`
	// The name of the run.
	RunName string `json:"run_name,omitempty"`
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Unix timestamp of when the run started in milliseconds.
	StartTime int64 `json:"start_time,omitempty"`
	// Current status of the run.
	Status RunInfoStatus `json:"status,omitempty"`
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status of a run.
type RunInfoStatus string

const RunInfoStatusFailed RunInfoStatus = `FAILED`

const RunInfoStatusFinished RunInfoStatus = `FINISHED`

const RunInfoStatusKilled RunInfoStatus = `KILLED`

const RunInfoStatusRunning RunInfoStatus = `RUNNING`

const RunInfoStatusScheduled RunInfoStatus = `SCHEDULED`

// String representation for [fmt.Print]
func (f *RunInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunInfoStatus) Set(v string) error {
	switch v {
	case `FAILED`, `FINISHED`, `KILLED`, `RUNNING`, `SCHEDULED`:
		*f = RunInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "FINISHED", "KILLED", "RUNNING", "SCHEDULED"`, v)
	}
}

// Values returns all possible values for RunInfoStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunInfoStatus) Values() []RunInfoStatus {
	return []RunInfoStatus{
		RunInfoStatusFailed,
		RunInfoStatusFinished,
		RunInfoStatusKilled,
		RunInfoStatusRunning,
		RunInfoStatusScheduled,
	}
}

// Type always returns RunInfoStatus to satisfy [pflag.Value] interface
func (f *RunInfoStatus) Type() string {
	return "RunInfoStatus"
}

// Run inputs.
type RunInputs struct {
	// Datasets used as inputs to the run.
	DatasetInputs []DatasetInput `json:"dataset_inputs,omitempty"`
	// Model inputs to the Run.
	ModelInputs []ModelInput `json:"model_inputs,omitempty"`
}

func (s *RunInputs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Tag for a run.
type RunTag struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A sawtooth window served via the hybrid batch + streaming path. The batch
// pipeline maintains daily partial aggregates for the bulk of the window while
// the streaming pipeline maintains the most recent day(s), and serving merges
// them on read. Same field shape as RollingWindow, but a distinct type so the
// control plane can explicitly identify hybrid (sawtooth) features rather than
// inferring hybrid behavior from window_duration.
type SawtoothWindow struct {
	// Delay is not currently supported for Sawtooth windows.
	Delay *duration.Duration `json:"delay,omitempty"`
	// The duration of the window. Must be positive and span more than two days
	// when set, so that both the batch (N-1 day) and stale-path (N-2 day)
	// partial aggregates are well defined. The duration need not be a whole
	// number of days (e.g. 3 days 15 minutes is allowed). Absent means lifetime
	// (aggregate over the entity's entire history).
	WindowDuration *duration.Duration `json:"window_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SawtoothWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SawtoothWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Scalar data types for request-time field definitions. Only flat (non-nested)
// types are supported.
type ScalarDataType string

const ScalarDataTypeBinary ScalarDataType = `BINARY`

const ScalarDataTypeBoolean ScalarDataType = `BOOLEAN`

const ScalarDataTypeDate ScalarDataType = `DATE`

const ScalarDataTypeDecimal ScalarDataType = `DECIMAL`

const ScalarDataTypeDouble ScalarDataType = `DOUBLE`

const ScalarDataTypeFloat ScalarDataType = `FLOAT`

const ScalarDataTypeInteger ScalarDataType = `INTEGER`

const ScalarDataTypeLong ScalarDataType = `LONG`

const ScalarDataTypeShort ScalarDataType = `SHORT`

const ScalarDataTypeString ScalarDataType = `STRING`

const ScalarDataTypeTimestamp ScalarDataType = `TIMESTAMP`

// String representation for [fmt.Print]
func (f *ScalarDataType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ScalarDataType) Set(v string) error {
	switch v {
	case `BINARY`, `BOOLEAN`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INTEGER`, `LONG`, `SHORT`, `STRING`, `TIMESTAMP`:
		*f = ScalarDataType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BINARY", "BOOLEAN", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "INTEGER", "LONG", "SHORT", "STRING", "TIMESTAMP"`, v)
	}
}

// Values returns all possible values for ScalarDataType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ScalarDataType) Values() []ScalarDataType {
	return []ScalarDataType{
		ScalarDataTypeBinary,
		ScalarDataTypeBoolean,
		ScalarDataTypeDate,
		ScalarDataTypeDecimal,
		ScalarDataTypeDouble,
		ScalarDataTypeFloat,
		ScalarDataTypeInteger,
		ScalarDataTypeLong,
		ScalarDataTypeShort,
		ScalarDataTypeString,
		ScalarDataTypeTimestamp,
	}
}

// Type always returns ScalarDataType to satisfy [pflag.Value] interface
func (f *ScalarDataType) Type() string {
	return "ScalarDataType"
}

type SchemaConfig struct {
	// Avro schema in JSON format
	// (https://avro.apache.org/docs/current/specification/).
	AvroSchema string `json:"avro_schema,omitempty"`
	// Schema of the JSON object in standard IETF JSON schema format
	// (https://json-schema.org/).
	JsonSchema string `json:"json_schema,omitempty"`
	// Protocol Buffer schema with its payload message name.
	ProtoSchema *ProtoSchemaSpec `json:"proto_schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SchemaConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SchemaConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Schema locator for one side (payload or key) of a message. Identifies which
// schema to use in the schema registry and the serialization format.
type SchemaLocator struct {
	// Confluent Schema Registry schema locator.
	ConfluentSchema *SchemaLocatorConfluentSchema `json:"confluent_schema,omitempty"`
	// Serialization format for this schema.
	Format SchemaLocatorFormat `json:"format"`
}

func (s *SchemaLocator) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Confluent Schema Registry schema locator. The value to provide for `subject`
// depends on the naming strategy configured in your registry: -
// TopicNameStrategy (default): "{topic}-key" or "{topic}-value" e.g. for topic
// "transactions" use "transactions-value" for the payload and
// "transactions-key" for the key. - RecordNameStrategy: the fully-qualified
// record name e.g. "com.example.Payment" for Avro, the bare message name
// (without package) for Protobuf, or the `title` field value for JSON. -
// TopicRecordNameStrategy: "{topic}-{fully-qualified-record-name}" e.g.
// "transactions-com.example.Payment".
type SchemaLocatorConfluentSchema struct {
	// The Confluent schema registry subject name.
	Subject string `json:"subject"`
}

func (s *SchemaLocatorConfluentSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Supported serialization formats for a schema registry schema.
type SchemaLocatorFormat string

const SchemaLocatorFormatFormatAvro SchemaLocatorFormat = `FORMAT_AVRO`

const SchemaLocatorFormatFormatJson SchemaLocatorFormat = `FORMAT_JSON`

const SchemaLocatorFormatFormatProtobuf SchemaLocatorFormat = `FORMAT_PROTOBUF`

// String representation for [fmt.Print]
func (f *SchemaLocatorFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SchemaLocatorFormat) Set(v string) error {
	switch v {
	case `FORMAT_AVRO`, `FORMAT_JSON`, `FORMAT_PROTOBUF`:
		*f = SchemaLocatorFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FORMAT_AVRO", "FORMAT_JSON", "FORMAT_PROTOBUF"`, v)
	}
}

// Values returns all possible values for SchemaLocatorFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *SchemaLocatorFormat) Values() []SchemaLocatorFormat {
	return []SchemaLocatorFormat{
		SchemaLocatorFormatFormatAvro,
		SchemaLocatorFormatFormatJson,
		SchemaLocatorFormatFormatProtobuf,
	}
}

// Type always returns SchemaLocatorFormat to satisfy [pflag.Value] interface
func (f *SchemaLocatorFormat) Type() string {
	return "SchemaLocatorFormat"
}

// Configuration for resolving a Stream's schema from an external schema
// registry (e.g. Confluent).
type SchemaRegistryConfig struct {
	// Reference to the schema registry API secret in a Databricks secret scope.
	// Set this only if required for authentication for the schema registry.
	ApiSecretRef *SecretScopeReference `json:"api_secret_ref,omitempty"`
	// Schema locator for the message key. Only used for Kafka streams. At least
	// one of payload_schema_locator or key_schema_locator must be set.
	KeySchemaLocator *SchemaLocator `json:"key_schema_locator,omitempty"`
	// Schema locator for the message payload. For Kafka this is the value. At
	// least one of payload_schema_locator or key_schema_locator must be set.
	PayloadSchemaLocator *SchemaLocator `json:"payload_schema_locator,omitempty"`
	// A Schema Registry UC Connection object.
	UcConnection string `json:"uc_connection,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SchemaRegistryConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SchemaRegistryConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Spark Declarative Pipeline (SDP) data source.
type SdpSource struct {
	// Schema of the resulting dataframe after transformations, in Spark
	// StructType JSON format (from df.schema.json()). Required if
	// transformation_sql is specified. Example:
	// {"type":"struct","fields":[{"name":"col_a","type":"integer","nullable":true,"metadata":{}},{"name":"col_c","type":"integer","nullable":true,"metadata":{}}]}
	DataframeSchema string `json:"dataframe_schema,omitempty"`
	// Single WHERE clause to filter the pipeline's output table before applying
	// transformations. Will be row-wise evaluated, so should only include
	// conditionals and projections.
	FilterCondition string `json:"filter_condition,omitempty"`
	// The full three-part names of the tables the pipeline reads from. Required
	// to materialize this feature with a TableTrigger, which starts a run when
	// these tables update. Each table must exist when the feature is created.
	// Feature Store doesn't confirm that the pipeline reads from these tables.
	// At most 25 tables.
	InputTables []string `json:"input_tables,omitempty"`
	// The full three-part (catalog, schema, table) name of the pipeline's
	// output table to read from. The table doesn't need to exist until the
	// pipeline first runs. Required.
	OutputTable string `json:"output_table,omitempty"`
	// The ID of the Spark Declarative Pipeline that Feature Store runs to
	// produce this source, before each materialization or backfill. The
	// pipeline must already exist when the feature is created, and you need
	// CAN_RUN, CAN_MANAGE, or owner permission on it. It doesn't need to have
	// run yet. Deleting the feature doesn't delete the pipeline. If the
	// pipeline is deleted, this field reads back empty. Required.
	PipelineId string `json:"pipeline_id,omitempty"`
	// State tracking pipeline execution and output table validation. PENDING
	// until the pipeline has run at least once and created the output table;
	// then RESOLVED on success, or ERROR on failure.
	PipelineState SdpSourceState `json:"pipeline_state,omitempty"`
	// A single SQL SELECT expression applied after filter_condition. Should
	// contain all the columns needed (eg. "SELECT *, col_a + col_b AS col_c
	// FROM x.y.z WHERE col_a > 0" would have `transformation_sql` "*, col_a +
	// col_b AS col_c") If transformation_sql is not provided, all columns of
	// the pipeline's output table are present in the DataSource dataframe.
	TransformationSql string `json:"transformation_sql,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SdpSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SdpSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// State of the SDP source relative to pipeline execution and output table
// validation.
type SdpSourceState string

const SdpSourceStateError SdpSourceState = `ERROR`

const SdpSourceStatePending SdpSourceState = `PENDING`

const SdpSourceStateResolved SdpSourceState = `RESOLVED`

// String representation for [fmt.Print]
func (f *SdpSourceState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SdpSourceState) Set(v string) error {
	switch v {
	case `ERROR`, `PENDING`, `RESOLVED`:
		*f = SdpSourceState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ERROR", "PENDING", "RESOLVED"`, v)
	}
}

// Values returns all possible values for SdpSourceState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SdpSourceState) Values() []SdpSourceState {
	return []SdpSourceState{
		SdpSourceStateError,
		SdpSourceStatePending,
		SdpSourceStateResolved,
	}
}

// Type always returns SdpSourceState to satisfy [pflag.Value] interface
func (f *SdpSourceState) Type() string {
	return "SdpSourceState"
}

type SearchExperiments struct {
	// Filter condition on the experiment name and tags. The `name` attribute
	// and `tags.<key>` support the operators `=`, `!=`, `LIKE` and `ILIKE`
	// (case-insensitive `LIKE`), string values are wrapped in single quotes,
	// and several conditions can be combined with `AND`. No wildcard is added
	// automatically, so include `%` in a `LIKE` pattern where you need one.
	Filter string `json:"filter,omitempty"`
	// Maximum number of experiments to return per page. Larger values are
	// capped by the server, so use `page_token` to retrieve further pages.
	MaxResults int64 `json:"max_results,omitempty"`
	// Columns to order the results by, each given as `name`, `basename` (the
	// last segment of the name), `experiment_id`, `creation_time` or
	// `last_update_time` with an optional `ASC` or `DESC` suffix, where `ASC`
	// is the default. If omitted, results are ordered by last update time with
	// the most recently updated experiment first. Tiebreaks are done by
	// experiment id DESC.
	OrderBy []string `json:"order_by,omitempty"`
	// Page token returned as `next_page_token` by a previous call. Omit it to
	// fetch the first page.
	PageToken string `json:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType ViewType `json:"view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchExperiments) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchExperiments) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	Experiments []Experiment `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchExperimentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchExperimentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request for SearchIssues.
type SearchIssuesRequest struct {
	// Matches Issues with any of these categories, case-insensitively.
	Categories []string `json:"categories,omitempty"`
	// Defaults to 20; values above 100 are clamped.
	PageSize int `json:"page_size,omitempty"`
	// Expires after 15 minutes. Every other argument except `view` must match
	// the request that returned the token.
	PageToken string `json:"page_token,omitempty"`
	// The experiment to search, as `experiments/{experiment}`, where
	// `{experiment}` is the experiment ID. For example,
	// `experiments/1234567890`.
	Parent string `json:"-" url:"-"`
	// Case-insensitive substring match on title.
	Query string `json:"query,omitempty"`
	// Matches Issues with any of these severities.
	Severities []IssueSeverity `json:"severities,omitempty"`
	// Defaults to SEVERITY.
	SortKey IssueSortKey `json:"sort_key,omitempty"`
	// Defaults to DESCENDING.
	SortOrder SortOrder `json:"sort_order,omitempty"`
	// Matches Issues with any of these statuses. Different filters AND
	// together, and an empty filter matches every Issue.
	Statuses []IssueStatus `json:"statuses,omitempty"`
	// Matches Issues that satisfy every filter. At most 10 filters of 20 values
	// each.
	TagFilters []IssueTagFilter `json:"tag_filters,omitempty"`
	// Defaults to BASIC.
	View IssueView `json:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchIssuesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchIssuesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// One page of SearchIssues results.
type SearchIssuesResponse struct {
	// The page's Issues, in the requested order.
	Issues []Issue `json:"issues,omitempty"`
	// Absent on the last page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Exact number of matching Issues.
	TotalSize int `json:"total_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchIssuesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchIssuesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchLoggedModelsDataset struct {
	// The digest of the dataset.
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset.
	DatasetName string `json:"dataset_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchLoggedModelsDataset) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchLoggedModelsDataset) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchLoggedModelsOrderBy struct {
	// Whether the search results order is ascending or not.
	Ascending bool `json:"ascending,omitempty"`
	// If ``field_name`` refers to a metric, this field specifies the digest of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name and digest will be considered for ordering. This
	// field may only be set if ``dataset_name`` is also set.
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// If ``field_name`` refers to a metric, this field specifies the name of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name will be considered for ordering. This field may
	// only be set if ``field_name`` refers to a metric.
	DatasetName string `json:"dataset_name,omitempty"`
	// The name of the field to order by, e.g. "metrics.accuracy".
	FieldName string `json:"field_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchLoggedModelsOrderBy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchLoggedModelsOrderBy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchLoggedModelsRequest struct {
	// List of datasets on which to apply the metrics filter clauses. For
	// example, a filter with `metrics.accuracy > 0.9` and dataset info with
	// name "test_dataset" means we will return all logged models with accuracy
	// > 0.9 on the test_dataset. Metric values from ANY dataset matching the
	// criteria are considered. If no datasets are specified, then metrics
	// across all datasets are considered in the filter.
	Datasets []SearchLoggedModelsDataset `json:"datasets,omitempty"`
	// The IDs of the experiments in which to search for logged models.
	ExperimentIds []string `json:"experiment_ids,omitempty"`
	// A filter expression over logged model info and data that allows returning
	// a subset of logged models. The syntax is a subset of SQL that supports
	// AND'ing together binary operations.
	//
	// Example: ``params.alpha < 0.3 AND metrics.accuracy > 0.9``.
	Filter string `json:"filter,omitempty"`
	// The maximum number of Logged Models to return. The maximum limit is 50.
	MaxResults int `json:"max_results,omitempty"`
	// The list of columns for ordering the results, with additional fields for
	// sorting criteria.
	OrderBy []SearchLoggedModelsOrderBy `json:"order_by,omitempty"`
	// The token indicating the page of logged models to fetch.
	PageToken string `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchLoggedModelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchLoggedModelsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchLoggedModelsResponse struct {
	// Logged models that match the search criteria.
	Models []LoggedModel `json:"models,omitempty"`
	// The token that can be used to retrieve the next page of logged models.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchLoggedModelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchLoggedModelsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchModelVersionsRequest struct {
	// Filter condition on `name`, `run_id`, `source` or `tags.<key>`. The
	// filter must constrain at least one of `name`, `run_id` or `source`.
	// Conditions use `=`, `!=`, `LIKE` or `ILIKE`, and `run_id` also supports
	// `IN`. String values are wrapped in single quotes, and several conditions
	// can be combined with `AND`.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of models desired. Max threshold is 10K.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// Columns to order the results by, each given as `name`, `version`,
	// `current_stage` or `last_updated_timestamp` with an optional `ASC` or
	// `DESC` suffix, where `ASC` is the default. Tiebreaks are done by latest
	// stage transition timestamp, followed by name ASC, followed by version
	// DESC.
	OrderBy []string `json:"-" url:"order_by,omitempty"`
	// Pagination token to go to next page based on previous search query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchModelVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchModelVersionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchModelVersionsResponse struct {
	// Model versions that match the search filter.
	ModelVersions []ModelVersion `json:"model_versions,omitempty"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchModelVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchModelVersionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchModelsRequest struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	OrderBy []string `json:"-" url:"order_by,omitempty"`
	// Pagination token to go to the next page based on a previous search query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchModelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchModelsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Registered Models that match the search criteria.
	RegisteredModels []Model `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchModelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchModelsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchRuns struct {
	// List of experiment IDs to search over.
	ExperimentIds []string `json:"experiment_ids,omitempty"`
	// A filter expression over params, metrics, and tags, that allows returning
	// a subset of runs. The syntax is a subset of SQL that supports ANDing
	// together binary operations between a param, metric, or tag and a
	// constant.
	//
	// Example: `metrics.rmse < 1 and params.model_class = 'LogisticRegression'`
	//
	// You can select columns with special characters (hyphen, space, period,
	// etc.) by using double quotes: `metrics."model class" = 'LinearRegression'
	// and tags."user-name" = 'Tomas'`
	//
	// Supported operators are `=`, `!=`, `>`, `>=`, `<`, and `<=`.
	Filter string `json:"filter,omitempty"`
	// Maximum number of runs desired. Max threshold is 50000
	MaxResults int `json:"max_results,omitempty"`
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional `"DESC"` or `"ASC"` annotation, where `"ASC"`
	// is the default. Example: `["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"]`. Tiebreaks are done by start_time `DESC` followed by
	// `run_id` for runs with the same start time (and this is the default
	// ordering criterion if order_by is not provided).
	OrderBy []string `json:"order_by,omitempty"`
	// Token for the current page of runs.
	PageToken string `json:"page_token,omitempty"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	RunViewType ViewType `json:"run_view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchRuns) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchRuns) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SearchRunsResponse struct {
	// Token for the next page of runs.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Runs that match the search criteria.
	Runs []Run `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Returns the values of one entity column that an upstream Feature has stored
// for the request's other entity columns, as an array. For example, given a
// Feature with entities (user_id, ad_id), returns the ad_ids stored for the
// requested user_id. The Feature's source must be a FeatureViewSource that
// includes `feature`.
type SecondaryKeys struct {
	// The upstream Feature whose stored keys are selected.
	Feature FeatureReference `json:"feature"`
	// The entity column of the upstream Feature to return values for. It is not
	// supplied by the request.
	Key string `json:"key"`
	// Maximum number of keys to return per entity, in primary key order. Must
	// not exceed the workspace limit.
	MaxKeys int64 `json:"max_keys"`
}

func (s *SecondaryKeys) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Reference to an entry in a Databricks secret scope.
type SecretScopeReference struct {
	// The key within the scope.
	Key string `json:"key"`
	// The Databricks secret scope name.
	Scope string `json:"scope"`
}

func (s *SecretScopeReference) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SetExperimentTag struct {
	// ID of the experiment under which to log the tag. Must be provided.
	ExperimentId string `json:"experiment_id"`
	// Key of the tag. Keys are at most 255 characters long. Setting a key that
	// already exists replaces its value. System tags such as `mlflow.ownerId`,
	// `mlflow.ownerEmail` and `mlflow.experimentType` cannot be set.
	Key string `json:"key"`
	// Value of the tag. Values are at most 65535 characters long.
	Value string `json:"value"`
}

func (s *SetExperimentTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SetLoggedModelTagsRequest struct {
	// The ID of the logged model to set the tags on.
	ModelId string `json:"-" url:"-"`
	// The tags to set on the logged model.
	Tags []LoggedModelTag `json:"tags,omitempty"`
}

func (s *SetLoggedModelTagsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SetModelTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key string `json:"key"`
	// Unique name of the model.
	Name string `json:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value string `json:"value"`
}

func (s *SetModelTagRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SetModelVersionTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key string `json:"key"`
	// Unique name of the model.
	Name string `json:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value string `json:"value"`
	// Model version number.
	Version string `json:"version"`
}

func (s *SetModelVersionTagRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SetTag struct {
	// Name of the tag. Keys up to 250 bytes in size are supported.
	Key string `json:"key"`
	// ID of the run under which to log the tag. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	Value string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SetTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SetTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SlidingWindow struct {
	// Non-negative analytic lag that evaluates the window this far in the past.
	// Use this for timing variations unrelated to source lateness, such as a
	// 30-day count as of one week ago. If unset, the analytic lag is zero. It
	// composes with source.lateness when both are set.
	Delay *duration.Duration `json:"delay,omitempty"`
	// Non-negative phase shift from the default midnight UTC alignment. For
	// example, offset=22h on a 24h slide produces boundaries at 22:00 UTC
	// (17:00 New York in standard time) instead of midnight UTC. If unset, the
	// offset is zero. Must be shorter than slide_duration (and therefore
	// window_duration).
	Offset *duration.Duration `json:"offset,omitempty"`
	// The slide duration (interval by which windows advance, must be positive
	// and less than duration).
	SlideDuration string `json:"slide_duration"`
	// The duration of the sliding window. Must be positive when set; absent
	// means lifetime (aggregate over the entity's entire history).
	WindowDuration string `json:"window_duration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SlidingWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SlidingWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The direction of the ordering.
type SortOrder string

const SortOrderAscending SortOrder = `ASCENDING`

const SortOrderDescending SortOrder = `DESCENDING`

// String representation for [fmt.Print]
func (f *SortOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SortOrder) Set(v string) error {
	switch v {
	case `ASCENDING`, `DESCENDING`:
		*f = SortOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASCENDING", "DESCENDING"`, v)
	}
}

// Values returns all possible values for SortOrder.
//
// There is no guarantee on the order of the values in the slice.
func (f *SortOrder) Values() []SortOrder {
	return []SortOrder{
		SortOrderAscending,
		SortOrderDescending,
	}
}

// Type always returns SortOrder to satisfy [pflag.Value] interface
func (f *SortOrder) Type() string {
	return "SortOrder"
}

// Configures when event-time data from this source is considered complete for a
// Feature.
type SourceLateness struct {
	// Non-negative time to wait after a window ends before treating its source
	// data as complete. Training shifts the eligible evaluation time backwards
	// by this duration so it does not join data that would still have been
	// settling online. Materialization waits for the duration to elapse before
	// publishing the window. If unset, source data is considered settled
	// immediately.
	SettlingDelay *duration.Duration `json:"settling_delay,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SourceLateness) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SourceLateness) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks.
//
// * `FAILED_REGISTRATION`: Request to register a new model version has failed.
//
// * `READY`: Model version is ready for use.
type Status string

// Request to register a new model version has failed.
const StatusFailedRegistration Status = `FAILED_REGISTRATION`

// Request to register a new model version is pending as server performs
// background tasks.
const StatusPendingRegistration Status = `PENDING_REGISTRATION`

// Model version is ready for use.
const StatusReady Status = `READY`

// String representation for [fmt.Print]
func (f *Status) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Status) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = Status(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Values returns all possible values for Status.
//
// There is no guarantee on the order of the values in the slice.
func (f *Status) Values() []Status {
	return []Status{
		StatusFailedRegistration,
		StatusPendingRegistration,
		StatusReady,
	}
}

// Type always returns Status to satisfy [pflag.Value] interface
func (f *Status) Type() string {
	return "Status"
}

// Computes the population standard deviation.
type StddevPopFunction struct {
	// The input column from which the population standard deviation is
	// computed. For Kafka sources, use dot-prefixed path notation (e.g.,
	// "value.amount"). For nested fields, the leaf node name is used.
	// Colon-prefixed notation (e.g., "value:amount") is supported for backwards
	// compatibility but is deprecated; migrate to dot notation.
	Input string `json:"input"`
}

func (s *StddevPopFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Computes the sample standard deviation.
type StddevSampFunction struct {
	// The input column from which the sample standard deviation is computed.
	Input string `json:"input"`
}

func (s *StddevSampFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A Stream is a governed UC entity representing an external streaming data
// source. The source_config field determines the streaming platform source
// (e.g. Kafka, Kinesis).
type Stream struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Specifies how to connect and authenticate to the stream platform.
	ConnectionConfig StreamConnectionConfig `json:"connection_config"`
	// Time at which this Stream was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Username of the Stream creator.
	CreatedBy string `json:"created_by,omitempty"`
	// User-provided description.
	Description string `json:"description,omitempty"`
	// Column paths (dot notation, e.g. "value.email" for Kafka) to drop. A path
	// may reference a struct, in which case all of its nested fields are
	// dropped (e.g. "value.address" drops "value.address.city" and
	// "value.address.zip"). These columns are not written to the ingestion
	// table and cannot be referenced by any feature. They are dropped from
	// ingestion, backfill, and materialization. For direct schemas, each column
	// must exist in the relevant key or payload schema. With a schema registry,
	// a column can be excluded before it exists. A column cannot also be a
	// deduplication column in the ingestion_config.
	ExcludedColumns []string `json:"excluded_columns,omitempty"`
	// Configuration for streaming data ingestion: the managed table storing an
	// offline copy of forward fill data and optional historical backfill.
	IngestionConfig IngestionConfig `json:"ingestion_config"`
	// Full three-part (catalog.schema.stream) name of the stream.
	Name string `json:"name"`
	// Optional SQL predicate to filter which record types from a streaming
	// channel (e.g. a topic for Kafka) belong to this Stream. Events that do
	// not match are not written to the ingestion table and are not used in
	// materialization. Example: "value.event_type = 'transaction'".
	RecordTypeFilter string `json:"record_type_filter,omitempty"`
	// Schema definitions for the stream, provided either directly on the Stream
	// or resolved from an external schema registry through a UC Connection.
	SchemaConfig StreamSchemaConfig `json:"schema_config"`
	// Source-specific configuration. Determines the streaming platform source.
	SourceConfig StreamSourceConfig `json:"source_config"`
	// Time at which this Stream was last modified.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Username of user who last modified the Stream.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Stream) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Stream) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of Kinesis stream ARNs to read from.
type StreamArnList struct {
	// Kinesis stream ARNs to read from. For example,
	// 'arn:aws:kinesis:us-west-2:111122223333:stream/stream-a'.
	Arns []string `json:"arns,omitempty"`
}

func (s *StreamArnList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Specifies how to connect and authenticate to the stream platform.
type StreamConnectionConfig struct {
	// Direct mTLS configuration for stream platform access.
	DirectMtlsConfig *DirectMtlsConfig `json:"direct_mtls_config,omitempty"`
	// Name of an existing UC Connection for stream platform access. Must be the
	// correct type for the streaming platform (e.g. a Kafka Connection for a
	// Kafka Stream, or a Kinesis Connection for a Kinesis Stream).
	UcConnectionName string `json:"uc_connection_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StreamConnectionConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StreamConnectionConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of Kinesis stream names to read from.
type StreamNameList struct {
	// Kinesis stream names to read from.
	Names []string `json:"names,omitempty"`
}

func (s *StreamNameList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Schema definitions for the stream. Feature store supports both direct schemas
// and schema registries.
type StreamSchemaConfig struct {
	// Schema definitions provided directly on the Stream.
	DirectSchemas *DirectSchemas `json:"direct_schemas,omitempty"`
	// Resolve schemas from an external schema registry.
	SchemaRegistryConfig *SchemaRegistryConfig `json:"schema_registry_config,omitempty"`
}

func (s *StreamSchemaConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A Stream entity used as a data source for a feature.
type StreamSource struct {
	// Schema of the resulting dataframe after transformations, in Spark
	// StructType JSON format (from df.schema.json()). Any subsequent functions
	// operate against this dataframe.
	DataframeSchema string `json:"dataframe_schema,omitempty"`
	// The filter condition applied to the source data before aggregation.
	FilterCondition string `json:"filter_condition,omitempty"`
	// Three-part full name of the Stream (catalog.schema.stream).
	FullName string `json:"full_name"`
	// The pipeline runs these SQL statements immediately after conversion into
	// the schema specified on the Stream object.
	TransformationSql string `json:"transformation_sql,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StreamSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StreamSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Source-specific configuration. Determines the streaming platform source.
type StreamSourceConfig struct {
	// Configuration for Apache Kafka streams.
	KafkaStreamConfig *KafkaStreamConfig `json:"kafka_stream_config,omitempty"`
	// Configuration for AWS Kinesis Data Streams.
	KinesisStreamConfig *KinesisStreamConfig `json:"kinesis_stream_config,omitempty"`
}

func (s *StreamSourceConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The streaming mode configuration for a streaming materialization pipeline.
type StreamingMode struct {
	// The desired data freshness for feature materialization, expressed as a
	// duration string (e.g. "1 minute").
	FreshnessTarget string `json:"freshness_target,omitempty"`
	// The type of streaming mode used by the materialization pipeline.
	Mode StreamingModeStreamingModeType `json:"mode,omitempty"`
	// Number of shuffle partitions for streaming materialization of this
	// feature. Higher values process high-throughput features with more
	// parallelism at higher compute cost. Materialized features which are
	// computed together will use the largest value set among them.
	ShufflePartitions int `json:"shuffle_partitions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StreamingMode) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StreamingMode) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StreamingModeStreamingModeType string

const StreamingModeStreamingModeTypeStreamingModeTypeMbm StreamingModeStreamingModeType = `STREAMING_MODE_TYPE_MBM`

const StreamingModeStreamingModeTypeStreamingModeTypeRtm StreamingModeStreamingModeType = `STREAMING_MODE_TYPE_RTM`

// String representation for [fmt.Print]
func (f *StreamingModeStreamingModeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *StreamingModeStreamingModeType) Set(v string) error {
	switch v {
	case `STREAMING_MODE_TYPE_MBM`, `STREAMING_MODE_TYPE_RTM`:
		*f = StreamingModeStreamingModeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "STREAMING_MODE_TYPE_MBM", "STREAMING_MODE_TYPE_RTM"`, v)
	}
}

// Values returns all possible values for StreamingModeStreamingModeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *StreamingModeStreamingModeType) Values() []StreamingModeStreamingModeType {
	return []StreamingModeStreamingModeType{
		StreamingModeStreamingModeTypeStreamingModeTypeMbm,
		StreamingModeStreamingModeTypeStreamingModeTypeRtm,
	}
}

// Type always returns StreamingModeStreamingModeType to satisfy [pflag.Value] interface
func (f *StreamingModeStreamingModeType) Type() string {
	return "StreamingModeStreamingModeType"
}

// Deprecated: Use KafkaSubscriptionMode instead.
type SubscriptionMode struct {
	// A JSON string that contains the specific topic-partitions to consume
	// from. For example, for '{"topicA":[0,1],"topicB":[2,4]}', topicA's 0'th
	// and 1st partitions will be consumed from.
	Assign string `json:"assign,omitempty"`
	// A comma-separated list of Kafka topics to read from. For example,
	// 'topicA,topicB,topicC'.
	Subscribe string `json:"subscribe,omitempty"`
	// A regular expression matching topics to subscribe to. For example,
	// 'topic.*' will subscribe to all topics starting with 'topic'.
	SubscribePattern string `json:"subscribe_pattern,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SubscriptionMode) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SubscriptionMode) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Computes the sum of values.
type SumFunction struct {
	// The input column from which the sum is computed. For Kafka sources, use
	// dot-prefixed path notation (e.g., "value.amount"). For nested fields, the
	// leaf node name is used. Colon-prefixed notation (e.g., "value:amount") is
	// supported for backwards compatibility but is deprecated; migrate to dot
	// notation.
	Input string `json:"input"`
}

func (s *SumFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A trigger that starts batch materialization when the feature's source table
// changes: a DeltaTableSource's table, or an SdpSource's input_tables.
type TableTrigger struct {
	// Which source table updates trigger a materialization run. If not set, any
	// source table update triggers a run.
	Condition TableTriggerTableTriggerCondition `json:"condition,omitempty"`
	// The minimum time between two triggered runs, so a frequently updated
	// source table does not trigger runs too often. If not set, a feature with
	// a sliding window uses half its slide_duration and a feature with a
	// tumbling window uses half its window_duration, clamped to between 5
	// minutes and 1 hour. Other features have no minimum.
	MinTimeBetweenTriggers *duration.Duration `json:"min_time_between_triggers,omitempty"`
	// Whether automatic materialization is paused. Defaults to `false`.
	// Supported only for materializations in a materialization set. Use the
	// `UpdateMaterializationSet` API to pause or resume the set.
	//
	// Pausing does not cancel running or queued runs. Runs can still be
	// triggered through the `TriggerMaterializationSet` API. Resuming does not
	// submit a run. Creating a paused materialization skips its initial run.
	Paused bool `json:"paused,omitempty"`
	// How long the source table must go without updates before a run starts, so
	// a series of updates settles before the run. If not set, a run starts as
	// soon as the source table updates.
	WaitAfterLastChange *duration.Duration `json:"wait_after_last_change,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableTrigger) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableTrigger) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Which source table updates trigger a materialization run.
type TableTriggerTableTriggerCondition string

const TableTriggerTableTriggerConditionTableTriggerConditionAllUpdated TableTriggerTableTriggerCondition = `TABLE_TRIGGER_CONDITION_ALL_UPDATED`

const TableTriggerTableTriggerConditionTableTriggerConditionAnyUpdated TableTriggerTableTriggerCondition = `TABLE_TRIGGER_CONDITION_ANY_UPDATED`

// String representation for [fmt.Print]
func (f *TableTriggerTableTriggerCondition) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableTriggerTableTriggerCondition) Set(v string) error {
	switch v {
	case `TABLE_TRIGGER_CONDITION_ALL_UPDATED`, `TABLE_TRIGGER_CONDITION_ANY_UPDATED`:
		*f = TableTriggerTableTriggerCondition(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TABLE_TRIGGER_CONDITION_ALL_UPDATED", "TABLE_TRIGGER_CONDITION_ANY_UPDATED"`, v)
	}
}

// Values returns all possible values for TableTriggerTableTriggerCondition.
//
// There is no guarantee on the order of the values in the slice.
func (f *TableTriggerTableTriggerCondition) Values() []TableTriggerTableTriggerCondition {
	return []TableTriggerTableTriggerCondition{
		TableTriggerTableTriggerConditionTableTriggerConditionAllUpdated,
		TableTriggerTableTriggerConditionTableTriggerConditionAnyUpdated,
	}
}

// Type always returns TableTriggerTableTriggerCondition to satisfy [pflag.Value] interface
func (f *TableTriggerTableTriggerCondition) Type() string {
	return "TableTriggerTableTriggerCondition"
}

// Details required to test a registry webhook.
type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event RegistryWebhookEvent `json:"event,omitempty"`
	// Webhook ID
	Id string `json:"id"`
}

func (s *TestRegistryWebhookRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TestRegistryWebhookResponse struct {
	// Body of the response from the webhook URL
	Body string `json:"body,omitempty"`
	// Status code returned by the webhook URL
	StatusCode int `json:"status_code,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TestRegistryWebhookResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TestRegistryWebhookResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TimeWindow struct {
	Continuous *ContinuousWindow `json:"continuous,omitempty"`

	Rolling *RollingWindow `json:"rolling,omitempty"`
	// A sawtooth window served via the hybrid batch + streaming path.
	Sawtooth *SawtoothWindow `json:"sawtooth,omitempty"`

	Sliding *SlidingWindow `json:"sliding,omitempty"`
	// Earliest event-time boundary at which the Feature may emit an output.
	// This gates outputs, not the historical inputs read by a window. For
	// example, a 365-day window with start_time=2026-01-01 begins emitting
	// partial-window values on that date instead of waiting for 365 days of
	// data; a lifetime window produces no output before start_time. If unset,
	// tumbling and fixed-duration sliding windows first emit at an
	// offset-aligned boundary after a full window can be formed. If unset,
	// lifetime sliding windows and rolling windows emit as soon as eligible
	// source data exists. For sawtooth windows, start_time must not be later
	// than the time the Feature is created. Not supported for rolling windows
	// on Features with a stream source.
	StartTime *time.Time `json:"start_time,omitempty"`

	Tumbling *TumblingWindow `json:"tumbling,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TimeWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TimeWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TimeseriesColumn struct {
	// The name of the timeseries column. For Kafka sources, use dot-prefixed
	// path notation to reference fields within the key or value schema (e.g.,
	// "value.event_timestamp"). For nested fields, the leaf node name (e.g.,
	// "event_timestamp" from "value.event_details.event_timestamp") is what
	// will be present in materialized tables and expected to match at query
	// time. Colon-prefixed notation (e.g., "value:event_timestamp") is
	// supported for backwards compatibility but is deprecated; migrate to dot
	// notation.
	Name string `json:"name"`
}

func (s *TimeseriesColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to transition a model version's stage.
type TransitionModelVersionStageDatabricks struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions bool `json:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage string `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TransitionModelVersionStageDatabricks) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TransitionModelVersionStageDatabricks) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []ActivityAction `json:"available_actions,omitempty"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	ToStage string `json:"to_stage,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TransitionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TransitionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TransitionStageResponse struct {
	// Updated model version
	ModelVersionDatabricks *ModelVersionDatabricks `json:"model_version_databricks,omitempty"`
}

func (s *TransitionStageResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to trigger materialization for a materialization set.
type TriggerMaterializationSetRequest struct {
	// The ID of the materialization set to trigger.
	MaterializationSetId string `json:"-" url:"-"`
	// Idempotency token in UUID format. Reusing it returns the original
	// operation without starting additional runs, even if the request differs
	// or the operation has finished. If omitted, each request triggers
	// independently.
	RequestId string `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TriggerMaterializationSetRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TriggerMaterializationSetRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Result of a successful `TriggerMaterializationSet` operation.
type TriggerMaterializationSetResponse struct {
}

func (s *TriggerMaterializationSetResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TumblingWindow struct {
	// Non-negative analytic lag that evaluates the window this far in the past.
	// Use this for timing variations unrelated to source lateness, such as a
	// 30-day count as of one week ago. If unset, the analytic lag is zero. It
	// composes with source.lateness when both are set.
	Delay *duration.Duration `json:"delay,omitempty"`
	// Non-negative phase shift from the default midnight UTC alignment. For
	// example, offset=22h on a 24h window produces boundaries at 22:00 UTC
	// (17:00 New York in standard time) instead of midnight UTC. If unset, the
	// offset is zero. Must be shorter than window_duration.
	Offset *duration.Duration `json:"offset,omitempty"`
	// The duration of each tumbling window (non-overlapping, fixed-duration
	// windows).
	WindowDuration string `json:"window_duration"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TumblingWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TumblingWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Unity Catalog trace storage location. Traces are stored as Delta tables in
// the specified catalog and schema.
type UcTraceLocation struct {
	// The name of the Unity Catalog catalog.
	Catalog string `json:"catalog"`
	// The trace-table prefix actually in effect: `table_prefix` if it was set
	// on creation, otherwise the server-generated default.
	EffectiveTablePrefix string `json:"effective_table_prefix,omitempty"`
	// The name of the Unity Catalog schema within `catalog`.
	Schema string `json:"schema"`
	// The prefix for the trace tables, which are named
	// `{catalog}.{schema}.{table_prefix}_otel_*`. May only contain letters,
	// digits, and underscores, and may be at most 238 characters. When unset, a
	// server-generated prefix derived from the experiment ID is used and this
	// field stays empty on read; the resolved value is always available in
	// `effective_table_prefix`.
	TablePrefix string `json:"table_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UcTraceLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UcTraceLocation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to edit a comment on a model version.
type UpdateComment struct {
	// User-provided comment on the action.
	Comment string `json:"comment"`
	// Unique identifier of an activity
	Id string `json:"id"`
}

func (s *UpdateComment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateCommentResponse struct {
	// Updated comment object
	Comment *CommentObject `json:"comment,omitempty"`
}

func (s *UpdateCommentResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
	// New name for the experiment, as an absolute path in the workspace. The
	// new name must be unique.
	NewName string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateExperiment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateExperiment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateFeatureRequest struct {
	// Feature whose full_name identifies the target. Only description is
	// mutable.
	Feature Feature `json:"feature"`
	// The full three-part name (catalog, schema, name) of the feature.
	FullName string `json:"-" url:"-"`
	// Fields to update. The only supported path is description.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateFeatureRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateFeatureTagRequest struct {
	FeatureName string `json:"-" url:"-"`

	FeatureTag FeatureTag `json:"feature_tag"`

	Key string `json:"-" url:"-"`

	TableName string `json:"-" url:"-"`
	// The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateFeatureTagRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateFeatureTagRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateIssueRequest struct {
	// `issue.name` identifies the target.
	Issue Issue `json:"issue"`
	// The Issue's resource name, as `experiments/{experiment}/issues/{issue}`,
	// where `{experiment}` is the experiment ID and `{issue}` is the Issue ID:
	// `iss-` followed by 32 lowercase hex digits. For example,
	// `experiments/1234567890/issues/iss-0123456789abcdef0123456789abcdef`.
	Name string `json:"-" url:"-"`
	// Accepts title, description, status, severity, categories, tags,
	// last_seen_time, and metadata.<key>. A path with an empty value clears the
	// field, except that title, status, and severity can't be cleared, and
	// last_seen_time can be cleared only when source_type is MANUAL. categories
	// and tags are replaced whole. metadata is updated per key: a
	// metadata.<key> path sets that key from issue.metadata, or removes it when
	// issue.metadata lacks it, and leaves the other keys alone. A bare metadata
	// path is rejected.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateIssueRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateKafkaConfigRequest struct {
	// The Kafka config to update.
	KafkaConfig KafkaConfig `json:"kafka_config"`
	// Name that uniquely identifies this Kafka config within the metastore.
	// This will be the identifier used from the Feature object to reference
	// these configs for a feature. Can be distinct from topic name.
	Name string `json:"-" url:"-"`
	// The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateKafkaConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateMaterializationSetRequest struct {
	// The desired state of the materialization set. `materialization_set_id`
	// identifies the set.
	//
	// A request that changes pause settings must request the same `paused`
	// value for all retained members with a cron schedule or table trigger.
	MaterializationSet MaterializationSet `json:"materialization_set"`
	// Server-assigned unique identifier for the materialization set.
	MaterializationSetId string `json:"-" url:"-"`
	// Idempotency token for the request. A second request carrying a token that
	// has already been used returns the operation the first one created,
	// whatever state it reached, and changes nothing further. The token is not
	// checked against the rest of the request. Omit it and each request updates
	// independently.
	RequestId string `json:"-" url:"request_id,omitempty"`
	// The fields to update, relative to the materialization set. Supported
	// values are `materialized_features` and `*`.
	//
	// The `materialized_features` mask updates retained members and deletes
	// omitted materialized features. The `*` mask selects all fields of the set
	// and follows the same membership rules. Adding members is not supported.
	//
	// To change pause settings without changing membership, include every
	// member of the set. An omitted `paused` is reset to `false`.
	//
	// Fields outside the mask are unchanged. Trigger types and settings other
	// than `paused` must remain unchanged.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateMaterializationSetRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMaterializationSetRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateMaterializedFeatureRequest struct {
	// The materialized feature to update.
	MaterializedFeature MaterializedFeature `json:"materialized_feature"`
	// Server-assigned unique identifier for the materialized feature.
	MaterializedFeatureId string `json:"-" url:"-"`
	// The fields to update. Only `pipeline_schedule_state` is supported.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateMaterializedFeatureRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description string `json:"description,omitempty"`
	// Registered model unique name identifier.
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateModelResponse struct {
	RegisteredModel *Model `json:"registered_model,omitempty"`
}

func (s *UpdateModelResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateModelVersionRequest struct {
	// New description for the model version.
	Description string `json:"description,omitempty"`
	// Name of the registered model
	Name string `json:"name"`
	// Model version number
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateModelVersionResponse struct {
	// The updated model version.
	ModelVersion *ModelVersion `json:"model_version,omitempty"`
}

func (s *UpdateModelVersionResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateOnlineStoreRequest struct {
	// The name of the online store. This is the unique identifier for the
	// online store.
	Name string `json:"-" url:"-"`
	// Online store to update.
	OnlineStore OnlineStore `json:"online_store"`
	// The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details required to update a registry webhook. Only the fields that need to
// be updated should be specified, and both `http_url_spec` and `job_spec`
// should not be specified in the same request.
type UpdateRegistryWebhook struct {
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events []RegistryWebhookEvent `json:"events,omitempty"`

	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"`
	// Webhook ID
	Id string `json:"id"`

	JobSpec *JobSpec `json:"job_spec,omitempty"`

	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRegistryWebhook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRegistryWebhook) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRun struct {
	// Unix timestamp in milliseconds of when the run ended.
	EndTime int64 `json:"end_time,omitempty"`
	// ID of the run to update. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// Updated name of the run.
	RunName string `json:"run_name,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run to update. This field
	// will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Updated status of the run.
	Status UpdateRunStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRunResponse struct {
	// Updated metadata of the run.
	RunInfo *RunInfo `json:"run_info,omitempty"`
}

func (s *UpdateRunResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Status of a run.
type UpdateRunStatus string

const UpdateRunStatusFailed UpdateRunStatus = `FAILED`

const UpdateRunStatusFinished UpdateRunStatus = `FINISHED`

const UpdateRunStatusKilled UpdateRunStatus = `KILLED`

const UpdateRunStatusRunning UpdateRunStatus = `RUNNING`

const UpdateRunStatusScheduled UpdateRunStatus = `SCHEDULED`

// String representation for [fmt.Print]
func (f *UpdateRunStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateRunStatus) Set(v string) error {
	switch v {
	case `FAILED`, `FINISHED`, `KILLED`, `RUNNING`, `SCHEDULED`:
		*f = UpdateRunStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "FINISHED", "KILLED", "RUNNING", "SCHEDULED"`, v)
	}
}

// Values returns all possible values for UpdateRunStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpdateRunStatus) Values() []UpdateRunStatus {
	return []UpdateRunStatus{
		UpdateRunStatusFailed,
		UpdateRunStatusFinished,
		UpdateRunStatusKilled,
		UpdateRunStatusRunning,
		UpdateRunStatusScheduled,
	}
}

// Type always returns UpdateRunStatus to satisfy [pflag.Value] interface
func (f *UpdateRunStatus) Type() string {
	return "UpdateRunStatus"
}

type UpdateStreamRequest struct {
	// Full three-part (catalog.schema.stream) name of the stream.
	Name string `json:"-" url:"-"`
	// The Stream to update.
	Stream Stream `json:"stream"`
	// The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateStreamRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWebhookResponse struct {
	Webhook *RegistryWebhook `json:"webhook,omitempty"`
}

func (s *UpdateWebhookResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Computes the population variance.
type VarPopFunction struct {
	// The input column from which the population variance is computed.
	Input string `json:"input"`
}

func (s *VarPopFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Computes the sample variance.
type VarSampFunction struct {
	// The input column from which the sample variance is computed.
	Input string `json:"input"`
}

func (s *VarSampFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Qualifier for the view type.
type ViewType string

const ViewTypeActiveOnly ViewType = `ACTIVE_ONLY`

const ViewTypeAll ViewType = `ALL`

const ViewTypeDeletedOnly ViewType = `DELETED_ONLY`

// String representation for [fmt.Print]
func (f *ViewType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ViewType) Set(v string) error {
	switch v {
	case `ACTIVE_ONLY`, `ALL`, `DELETED_ONLY`:
		*f = ViewType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE_ONLY", "ALL", "DELETED_ONLY"`, v)
	}
}

// Values returns all possible values for ViewType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ViewType) Values() []ViewType {
	return []ViewType{
		ViewTypeActiveOnly,
		ViewTypeAll,
		ViewTypeDeletedOnly,
	}
}

// Type always returns ViewType to satisfy [pflag.Value] interface
func (f *ViewType) Type() string {
	return "ViewType"
}
