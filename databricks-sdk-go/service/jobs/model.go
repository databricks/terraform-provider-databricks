// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/duration"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

// Configuration for a task that runs the `ai_classify` function.
//
// The label set comes from exactly one of two places: `inline_labels`, which
// holds up to 10 labels, or a Unity Catalog table named by `label_table`. Use
// `label_table` for a larger label set.
type AiClassifyFunction struct {
	// When true, includes a per-label confidence score in the response.
	EnableConfidenceScores bool `json:"enable_confidence_scores,omitempty"`
	// When true, includes a rationale explaining each classification in the
	// response.
	EnableRationales bool `json:"enable_rationales,omitempty"`
	// Labels listed inline, for a small label set.
	InlineLabels *AiClassifyFunctionInlineLabels `json:"inline_labels,omitempty"`
	// Global description of the task and domain, passed to the model to improve
	// classification quality.
	Instructions string `json:"instructions,omitempty"`
	// When true, more than one label can be returned for each input.
	Multilabel bool `json:"multilabel,omitempty"`
	// The ai_classify function version to invoke (for example, `2.1`).
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiClassifyFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiClassifyFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A label set listed inline.
type AiClassifyFunctionInlineLabels struct {
	// The labels an input can be classified as. At least two labels are
	// required unless `multilabel` is true, which also accepts a single label.
	Labels []AiClassifyFunctionLabel `json:"labels"`
}

func (s *AiClassifyFunctionInlineLabels) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// One label the input can be classified as.
type AiClassifyFunctionLabel struct {
	// Guidance describing what belongs under this label.
	Description string `json:"description,omitempty"`
	// The name returned when an input is classified as this label.
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiClassifyFunctionLabel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiClassifyFunctionLabel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration for a task that runs the `ai_extract` function.
//
// The fields to extract come from exactly one of two places: an existing AI
// Extract agent, named by `agent_id`, or the inline `schema` and
// `instructions`. Set `agent_id` or the inline fields, not both.
type AiExtractFunction struct {
	// When true, includes citation metadata locating each extracted value in
	// the source.
	EnableCitations bool `json:"enable_citations,omitempty"`
	// When true, includes a per-field confidence score in the response.
	EnableConfidenceScores bool `json:"enable_confidence_scores,omitempty"`
	// Global description of the task and domain, passed to the model to improve
	// extraction quality.
	Instructions string `json:"instructions,omitempty"`
	// The extraction mode. Defaults to standard extraction when unspecified.
	Mode AiExtractFunctionMode `json:"mode,omitempty"`
	// The extraction schema defining the fields to extract, as a JSON string.
	// Either an array of field names (e.g. ["company", "valuation"]) or an
	// object mapping each field name to a type and an optional description.
	// Supported types are string, integer, number, boolean, enum, object, and
	// array; all fields are optional.
	Schema string `json:"schema,omitempty"`
	// The ai_extract function version to invoke (for example, `2.1`).
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiExtractFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiExtractFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The extraction mode.
type AiExtractFunctionMode string

const AiExtractFunctionModeModePrecision AiExtractFunctionMode = `MODE_PRECISION`

// String representation for [fmt.Print]
func (f *AiExtractFunctionMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AiExtractFunctionMode) Set(v string) error {
	switch v {
	case `MODE_PRECISION`:
		*f = AiExtractFunctionMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MODE_PRECISION"`, v)
	}
}

// Values returns all possible values for AiExtractFunctionMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *AiExtractFunctionMode) Values() []AiExtractFunctionMode {
	return []AiExtractFunctionMode{
		AiExtractFunctionModeModePrecision,
	}
}

// Type always returns AiExtractFunctionMode to satisfy [pflag.Value] interface
func (f *AiExtractFunctionMode) Type() string {
	return "AiExtractFunctionMode"
}

// A task that applies an AI function to a set of input documents or rows.
type AiFunctionTask struct {
	// Classifies each input document or row into one or more labels.
	AiClassify *AiClassifyFunction `json:"ai_classify,omitempty"`
	// Extracts structured fields from the input documents or rows.
	AiExtract *AiExtractFunction `json:"ai_extract,omitempty"`
	// Parses the input documents into structured text.
	AiParseDocument *AiParseDocumentFunction `json:"ai_parse_document,omitempty"`
	// A Unity Catalog Delta table. The function is applied to each row.
	DeltaTable *DeltaTableSource `json:"delta_table,omitempty"`
	// Output destination table. When `table` is unset the task generates a
	// stable name from the function and task identity so a downstream task can
	// reference it.
	Output QualifiedName `json:"output"`
	// A Unity Catalog volume. The function is applied to each file.
	Volume *VolumeSource `json:"volume,omitempty"`
}

func (s *AiFunctionTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The result of an AI function task run. Also reported mid-run, so the counts
// can be partial.
type AiFunctionTaskOutput struct {
	// Coarse estimate of the time left to run the function over the inputs,
	// from the throughput so far. Covers that work only, not the merge into the
	// output table that follows it.
	EstimatedTimeRemaining *duration.Duration `json:"estimated_time_remaining,omitempty"`
	// Inputs permanently failed so far. Each lands in the output table as an
	// error row, and does not on its own fail the run.
	FailedRowCount int64 `json:"failed_row_count,omitempty"`
	// The table the run writes its results to. Reported even when
	// `output.table` is unset and the name is generated, so a client can read
	// the table.
	OutputTable *QualifiedName `json:"output_table,omitempty"`
	// Inputs with a result so far. Reported per batch, so it can lag mid-run.
	SucceededRowCount int64 `json:"succeeded_row_count,omitempty"`
	// Inputs to process: rows for a Delta table source, files for a volume
	// source. Stays well inside 2^53, so a JavaScript client can hold it
	// exactly.
	TotalRowCount int64 `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiFunctionTaskOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiFunctionTaskOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration for a task that runs the `ai_parse_document` function.
type AiParseDocumentFunction struct {
	// Element types to generate AI descriptions for. Leave empty to generate no
	// descriptions; use ALL for every supported type, or list specific types to
	// restrict generation to only those.
	DescriptionElementTypes []AiParseDocumentFunctionDescriptionElementType `json:"description_element_types,omitempty"`
	// Unity Catalog volume path where rendered page and element images are
	// written.
	ImageOutputPath string `json:"image_output_path,omitempty"`
	// Pages to parse (1-indexed), as a comma-separated list of page numbers or
	// ranges (e.g. "1,3,5-10").
	PageRange string `json:"page_range,omitempty"`
	// The ai_parse_document output schema version to invoke (for example,
	// `2.0`).
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiParseDocumentFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiParseDocumentFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An element type that AI descriptions can be generated for.
type AiParseDocumentFunctionDescriptionElementType string

const AiParseDocumentFunctionDescriptionElementTypeDescriptionElementTypeAll AiParseDocumentFunctionDescriptionElementType = `DESCRIPTION_ELEMENT_TYPE_ALL`

const AiParseDocumentFunctionDescriptionElementTypeDescriptionElementTypeFigure AiParseDocumentFunctionDescriptionElementType = `DESCRIPTION_ELEMENT_TYPE_FIGURE`

// String representation for [fmt.Print]
func (f *AiParseDocumentFunctionDescriptionElementType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AiParseDocumentFunctionDescriptionElementType) Set(v string) error {
	switch v {
	case `DESCRIPTION_ELEMENT_TYPE_ALL`, `DESCRIPTION_ELEMENT_TYPE_FIGURE`:
		*f = AiParseDocumentFunctionDescriptionElementType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DESCRIPTION_ELEMENT_TYPE_ALL", "DESCRIPTION_ELEMENT_TYPE_FIGURE"`, v)
	}
}

// Values returns all possible values for AiParseDocumentFunctionDescriptionElementType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AiParseDocumentFunctionDescriptionElementType) Values() []AiParseDocumentFunctionDescriptionElementType {
	return []AiParseDocumentFunctionDescriptionElementType{
		AiParseDocumentFunctionDescriptionElementTypeDescriptionElementTypeAll,
		AiParseDocumentFunctionDescriptionElementTypeDescriptionElementTypeFigure,
	}
}

// Type always returns AiParseDocumentFunctionDescriptionElementType to satisfy [pflag.Value] interface
func (f *AiParseDocumentFunctionDescriptionElementType) Type() string {
	return "AiParseDocumentFunctionDescriptionElementType"
}

// AiRuntimeTask: multi-node GPU compute task definition for Databricks AI
// Runtime workloads.
//
// Jobs-framework-level concepts (retries, per-task timeout, idempotency token,
// usage/budget policy, permissions) live on the surrounding TaskSettings /
// run-submit request and are intentionally NOT duplicated here. Users compose
// `ai_runtime_task` with the standard Jobs/DABs task wrapper to get those.
type AiRuntimeTask struct {
	// Workspace or UC volume path of the code-source archive, unpacked on each
	// node and exposed through `$CODE_SOURCE`. Set by first-party tooling; not
	// for direct callers.
	CodeSourcePath string `json:"code_source_path,omitempty"`
	// Deployment specs for this task. Exactly one deployment is currently
	// supported (a single entry where every node runs the same command); this
	// is a current-Preview constraint. Role-split workloads (driver + worker,
	// parameter server, separate eval node, etc.) with multiple entries are the
	// eventual intent but not yet supported.
	Deployments []DeploymentSpec `json:"deployments"`
	// Optional Docker image URL for a custom container image. When set, the
	// task runs on the specified container image instead of the default
	// Databricks client image. Format: `{organization}/{repository}:{tag}`
	DockerImageUrl string `json:"docker_image_url,omitempty"`
	// MLflow experiment name for this run. If an experiment with this name
	// already exists under the calling user, the run is appended to it;
	// otherwise a new experiment is created. To target a specific MLflow
	// storage location (for example, when running as a service principal), set
	// `mlflow_experiment_directory`.
	Experiment string `json:"experiment"`
	// Optional root location for MLflow artifacts logged by the run. If this
	// field isn't specified the default artifact location will be in dbfs i.e.
	// `dbfs:/databricks/mlflow-tracking/<experiment_id>/...` If dbfs access is
	// restricted or UC is preferred this can be a custom location in UC:
	// `dbfs:/Volumes/<catalog>/<schema>/<volume>/...` The location should be
	// unique for each experiment.
	MlflowArtifactLocation string `json:"mlflow_artifact_location,omitempty"`
	// Optional workspace directory under which the MLflow experiment named in
	// `experiment` is created. Must start with `/Workspace`. Set this when
	// running as a service principal that has no default user directory; for
	// regular users the experiment defaults to the user's home directory.
	MlflowExperimentDirectory string `json:"mlflow_experiment_directory,omitempty"`
	// Optional display name for the MLflow run created under `experiment`. If
	// omitted, MLflow generates a default name.
	MlflowRun string `json:"mlflow_run,omitempty"`
	// Scheduling priority class for the workload. May only be set together with
	// a GPU Pool (a deployment's `compute.accelerator_pool_id` or legacy
	// `compute.provisioned_capacity_id`); it is rejected on a workload that
	// runs on on-demand capacity.
	PriorityClass AiRuntimeTaskPriorityClass `json:"priority_class,omitempty"`
	// Optional Unity Catalog path for a custom container image. When set, the
	// task runs on the specified container image instead of the default
	// Databricks client image. Format: `{catalog}.{schema}.{image_name}:{tag}`
	UnityCatalogImagePath string `json:"unity_catalog_image_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiRuntimeTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiRuntimeTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// AiRuntimeTaskOutput: output identifiers for an AiRuntimeTask run — the
// MLflow experiment and run IDs the task wrote to.
//
// Run lifecycle and termination status are not on this message; they live on
// the surrounding `RunTask.status` field (see `runs.proto:RunTask.status`).
type AiRuntimeTaskOutput struct {
	// MLflow experiment ID the run was logged to. Use it to look up the
	// experiment in MLflow APIs or the workspace MLflow UI.
	MlflowExperimentId string `json:"mlflow_experiment_id,omitempty"`
	// MLflow run ID for this task execution. Use it to look up the run in
	// MLflow APIs or the workspace MLflow UI.
	MlflowRunId string `json:"mlflow_run_id,omitempty"`
	// Human-readable status message for this run, suitable for display to the
	// user (for example, that the run is still waiting for GPU compute). Set by
	// the server only when there is something to surface; empty otherwise.
	StatusMessage string `json:"status_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiRuntimeTaskOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiRuntimeTaskOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Scheduling priority class for a workload — its priority and preemptability
// when the scheduler ranks pending work.
type AiRuntimeTaskPriorityClass string

const AiRuntimeTaskPriorityClassHigh AiRuntimeTaskPriorityClass = `HIGH`

const AiRuntimeTaskPriorityClassNormal AiRuntimeTaskPriorityClass = `NORMAL`

const AiRuntimeTaskPriorityClassPreemptable AiRuntimeTaskPriorityClass = `PREEMPTABLE`

// String representation for [fmt.Print]
func (f *AiRuntimeTaskPriorityClass) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AiRuntimeTaskPriorityClass) Set(v string) error {
	switch v {
	case `HIGH`, `NORMAL`, `PREEMPTABLE`:
		*f = AiRuntimeTaskPriorityClass(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "HIGH", "NORMAL", "PREEMPTABLE"`, v)
	}
}

// Values returns all possible values for AiRuntimeTaskPriorityClass.
//
// There is no guarantee on the order of the values in the slice.
func (f *AiRuntimeTaskPriorityClass) Values() []AiRuntimeTaskPriorityClass {
	return []AiRuntimeTaskPriorityClass{
		AiRuntimeTaskPriorityClassHigh,
		AiRuntimeTaskPriorityClassNormal,
		AiRuntimeTaskPriorityClassPreemptable,
	}
}

// Type always returns AiRuntimeTaskPriorityClass to satisfy [pflag.Value] interface
func (f *AiRuntimeTaskPriorityClass) Type() string {
	return "AiRuntimeTaskPriorityClass"
}

// Same alert evaluation state as in redash-v2/api/proto/alertsv2/alerts.proto
type AlertEvaluationState string

const AlertEvaluationStateError AlertEvaluationState = `ERROR`

const AlertEvaluationStateOk AlertEvaluationState = `OK`

const AlertEvaluationStateTriggered AlertEvaluationState = `TRIGGERED`

const AlertEvaluationStateUnknown AlertEvaluationState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *AlertEvaluationState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertEvaluationState) Set(v string) error {
	switch v {
	case `ERROR`, `OK`, `TRIGGERED`, `UNKNOWN`:
		*f = AlertEvaluationState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ERROR", "OK", "TRIGGERED", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for AlertEvaluationState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertEvaluationState) Values() []AlertEvaluationState {
	return []AlertEvaluationState{
		AlertEvaluationStateError,
		AlertEvaluationStateOk,
		AlertEvaluationStateTriggered,
		AlertEvaluationStateUnknown,
	}
}

// Type always returns AlertEvaluationState to satisfy [pflag.Value] interface
func (f *AlertEvaluationState) Type() string {
	return "AlertEvaluationState"
}

type AlertTask struct {
	// The alert_id is the canonical identifier of the alert.
	AlertId string `json:"alert_id,omitempty"`
	// Per-run parameter overrides, keyed by parameter name, applied onto the
	// alert's stored query parameters before the query is executed. Only scalar
	// values are supported. Values may reference job parameters with
	// `{{job.parameters.*}}`, which are resolved before the task runs. An
	// override whose key does not match a stored parameter fails the task run.
	// Limited to 10000 characters when serialized as JSON; keys must be 1-100
	// characters and contain only letters, digits, underscores, dashes, and
	// periods.
	Parameters map[string]string `json:"parameters,omitempty"`
	// The subscribers receive alert evaluation result notifications after the
	// alert task is completed. The number of subscriptions is limited to 100.
	Subscribers []AlertTaskSubscriber `json:"subscribers,omitempty"`
	// The warehouse_id identifies the warehouse settings used by the alert
	// task.
	WarehouseId string `json:"warehouse_id,omitempty"`
	// The workspace_path is the path to the alert file in the workspace. The
	// path: * must start with "/Workspace" * must be a normalized path. User
	// has to select only one of alert_id or workspace_path to identify the
	// alert.
	WorkspacePath string `json:"workspace_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertTaskOutput struct {
	AlertState AlertEvaluationState `json:"alert_state,omitempty"`
}

func (s *AlertTaskOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Represents a subscriber that will receive alert notifications. A subscriber
// can be either a user (via email) or a notification destination (via
// destination_id).
type AlertTaskSubscriber struct {
	DestinationId string `json:"destination_id,omitempty"`
	// A valid workspace email address.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertTaskSubscriber) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertTaskSubscriber) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AuthenticationMethod string

const AuthenticationMethodOauth AuthenticationMethod = `OAUTH`

const AuthenticationMethodPat AuthenticationMethod = `PAT`

// String representation for [fmt.Print]
func (f *AuthenticationMethod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AuthenticationMethod) Set(v string) error {
	switch v {
	case `OAUTH`, `PAT`:
		*f = AuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OAUTH", "PAT"`, v)
	}
}

// Values returns all possible values for AuthenticationMethod.
//
// There is no guarantee on the order of the values in the slice.
func (f *AuthenticationMethod) Values() []AuthenticationMethod {
	return []AuthenticationMethod{
		AuthenticationMethodOauth,
		AuthenticationMethodPat,
	}
}

// Type always returns AuthenticationMethod to satisfy [pflag.Value] interface
func (f *AuthenticationMethod) Type() string {
	return "AuthenticationMethod"
}

type BaseJob struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime int64 `json:"created_time,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// The id of the usage policy used by this job for cost attribution
	// purposes.
	EffectiveUsagePolicyId string `json:"effective_usage_policy_id,omitempty"`
	// Indicates if the job has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/get endpoint.
	// It is only relevant for API 2.2 :method:jobs/list requests with
	// `expand_tasks=true`.
	HasMore bool `json:"has_more,omitempty"`
	// The canonical identifier for this job.
	JobId int64 `json:"job_id,omitempty"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *JobSettings `json:"settings,omitempty"`
	// Per-trigger runtime information for the multi-trigger surface. Same
	// length and order as `JobSettings.triggers`; `trigger_details[i]`
	// corresponds to `triggers[i]`. Sub-fields (`state`, `history`) are
	// populated independently based on the `GetJob.include_trigger_state` /
	// `include_trigger_history` flags.
	TriggerDetails []TriggerDetails `json:"trigger_details,omitempty"`
	// State of the trigger associated with the job.
	TriggerState *TriggerStateProto `json:"trigger_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BaseJob) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BaseJob) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BaseRun struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `json:"cluster_instance,omitempty"`
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec *ClusterSpec `json:"cluster_spec,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// ID of the deployment that produced the job when this run was created.
	// Used to look up deployment metadata from the Deployment Metadata service.
	// Only set for job runs of jobs with a `BUNDLE` deployment.
	DeploymentId string `json:"deployment_id,omitempty"`
	// Description of the run
	Description string `json:"description,omitempty"`
	// The window the run aims to finish within, in hours, for a Cost-optimized
	// serverless run. Required only when `performance_target` is
	// `COST_OPTIMIZED`. Must be one of the supported values: `NO_FLEXIBILITY`,
	// `4h`, `12h`, `24h`, `72h`. `NO_FLEXIBILITY` means the run uses the
	// default cost-optimized performance level without additional cost-saving
	// optimizations.
	DurationTarget string `json:"duration_target,omitempty"`
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance. * `STANDARD`:
	// Enables cost-efficient execution of serverless workloads. *
	// `COST_OPTIMIZED`: Enables lower job costs by optimizing compute for your
	// selected target duration time. Must provide a duration target.
	EffectivePerformanceTarget PerformanceTarget `json:"effective_performance_target,omitempty"`
	// The id of the usage policy used by this run for cost attribution
	// purposes.
	EffectiveUsagePolicyId string `json:"effective_usage_policy_id,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// Snapshot of `JobSettings.environment_variables` as it was at run launch
	// — the full list of named environment-variable entries the job defined.
	// To find which entry a given task ran with, look at
	// `RunTaskSettings.environment_variables_key`.
	EnvironmentVariables []JobEnvironmentVariables `json:"environment_variables,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *GitSource `json:"git_source,omitempty"`
	// Indicates if the run has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/getrun
	// endpoint. It is only relevant for API 2.2 :method:jobs/listruns requests
	// with `expand_tasks=true`.
	HasMore bool `json:"has_more,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings. If more than 100
	// job clusters are available, you can paginate through them using
	// :method:jobs/getrun.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// The canonical identifier of the job that contains this run.
	JobId int64 `json:"job_id,omitempty"`
	// Job-level parameters used in the run
	JobParameters []JobParameter `json:"job_parameters,omitempty"`
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	JobRunId int64 `json:"job_run_id,omitempty"`
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId int64 `json:"original_attempt_run_id,omitempty"`
	// The parameters used for this run.
	OverridingParameters *RunParameters `json:"overriding_parameters,omitempty"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration int64 `json:"queue_duration,omitempty"`
	// The repair history of the run.
	RepairHistory []RepairHistoryItem `json:"repair_history,omitempty"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int64 `json:"run_duration,omitempty"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId int64 `json:"run_id,omitempty"`
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName string `json:"run_name,omitempty"`
	// The URL to the detail page of the run.
	RunPageUrl string `json:"run_page_url,omitempty"`

	RunType RunType `json:"run_type,omitempty"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64 `json:"setup_duration,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// Deprecated. Please use the `status` field instead.
	State *RunState `json:"state,omitempty"`

	Status *RunStatus `json:"status,omitempty"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run results. If
	// more than 100 tasks are available, you can paginate through them using
	// :method:jobs/getrun. Use the `next_page_token` field at the object root
	// to determine if more results are available.
	Tasks []RunTask `json:"tasks,omitempty"`

	Trigger TriggerType `json:"trigger,omitempty"`

	TriggerInfo *TriggerInfo `json:"trigger_info,omitempty"`
	// ID of the deployment version that produced the job when this run was
	// created. Identifies a specific snapshot of the deployment in the
	// Deployment Metadata service. Only set for job runs of jobs with a
	// `BUNDLE` deployment.
	VersionId string `json:"version_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BaseRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BaseRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CancelAllRuns struct {
	// Optional boolean parameter to cancel all queued runs. If no job_id is
	// provided, all queued runs in the workspace are canceled.
	AllQueuedRuns bool `json:"all_queued_runs,omitempty"`
	// The canonical identifier of the job to cancel all runs of.
	JobId int64 `json:"job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CancelAllRuns) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CancelAllRuns) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CancelRun struct {
	// This field is required.
	RunId int64 `json:"run_id"`
}

func (s *CancelRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Copied from elastic-spark-common/api/messages/runs.proto. Using the original
// definition to remove coupling with jobs API definition
type CleanRoomTaskRunLifeCycleState string

const CleanRoomTaskRunLifeCycleStateBlocked CleanRoomTaskRunLifeCycleState = `BLOCKED`

const CleanRoomTaskRunLifeCycleStateInternalError CleanRoomTaskRunLifeCycleState = `INTERNAL_ERROR`

const CleanRoomTaskRunLifeCycleStatePending CleanRoomTaskRunLifeCycleState = `PENDING`

const CleanRoomTaskRunLifeCycleStateQueued CleanRoomTaskRunLifeCycleState = `QUEUED`

const CleanRoomTaskRunLifeCycleStateRunning CleanRoomTaskRunLifeCycleState = `RUNNING`

const CleanRoomTaskRunLifeCycleStateRunLifeCycleStateUnspecified CleanRoomTaskRunLifeCycleState = `RUN_LIFE_CYCLE_STATE_UNSPECIFIED`

const CleanRoomTaskRunLifeCycleStateSkipped CleanRoomTaskRunLifeCycleState = `SKIPPED`

const CleanRoomTaskRunLifeCycleStateTerminated CleanRoomTaskRunLifeCycleState = `TERMINATED`

const CleanRoomTaskRunLifeCycleStateTerminating CleanRoomTaskRunLifeCycleState = `TERMINATING`

const CleanRoomTaskRunLifeCycleStateWaitingForRetry CleanRoomTaskRunLifeCycleState = `WAITING_FOR_RETRY`

// String representation for [fmt.Print]
func (f *CleanRoomTaskRunLifeCycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomTaskRunLifeCycleState) Set(v string) error {
	switch v {
	case `BLOCKED`, `INTERNAL_ERROR`, `PENDING`, `QUEUED`, `RUNNING`, `RUN_LIFE_CYCLE_STATE_UNSPECIFIED`, `SKIPPED`, `TERMINATED`, `TERMINATING`, `WAITING_FOR_RETRY`:
		*f = CleanRoomTaskRunLifeCycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "INTERNAL_ERROR", "PENDING", "QUEUED", "RUNNING", "RUN_LIFE_CYCLE_STATE_UNSPECIFIED", "SKIPPED", "TERMINATED", "TERMINATING", "WAITING_FOR_RETRY"`, v)
	}
}

// Values returns all possible values for CleanRoomTaskRunLifeCycleState.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomTaskRunLifeCycleState) Values() []CleanRoomTaskRunLifeCycleState {
	return []CleanRoomTaskRunLifeCycleState{
		CleanRoomTaskRunLifeCycleStateBlocked,
		CleanRoomTaskRunLifeCycleStateInternalError,
		CleanRoomTaskRunLifeCycleStatePending,
		CleanRoomTaskRunLifeCycleStateQueued,
		CleanRoomTaskRunLifeCycleStateRunning,
		CleanRoomTaskRunLifeCycleStateRunLifeCycleStateUnspecified,
		CleanRoomTaskRunLifeCycleStateSkipped,
		CleanRoomTaskRunLifeCycleStateTerminated,
		CleanRoomTaskRunLifeCycleStateTerminating,
		CleanRoomTaskRunLifeCycleStateWaitingForRetry,
	}
}

// Type always returns CleanRoomTaskRunLifeCycleState to satisfy [pflag.Value] interface
func (f *CleanRoomTaskRunLifeCycleState) Type() string {
	return "CleanRoomTaskRunLifeCycleState"
}

// Copied from elastic-spark-common/api/messages/runs.proto. Using the original
// definition to avoid cyclic dependency.
type CleanRoomTaskRunResultState string

const CleanRoomTaskRunResultStateCanceled CleanRoomTaskRunResultState = `CANCELED`

const CleanRoomTaskRunResultStateDisabled CleanRoomTaskRunResultState = `DISABLED`

const CleanRoomTaskRunResultStateEvicted CleanRoomTaskRunResultState = `EVICTED`

const CleanRoomTaskRunResultStateExcluded CleanRoomTaskRunResultState = `EXCLUDED`

const CleanRoomTaskRunResultStateFailed CleanRoomTaskRunResultState = `FAILED`

const CleanRoomTaskRunResultStateMaximumConcurrentRunsReached CleanRoomTaskRunResultState = `MAXIMUM_CONCURRENT_RUNS_REACHED`

const CleanRoomTaskRunResultStateRunResultStateUnspecified CleanRoomTaskRunResultState = `RUN_RESULT_STATE_UNSPECIFIED`

const CleanRoomTaskRunResultStateSuccess CleanRoomTaskRunResultState = `SUCCESS`

const CleanRoomTaskRunResultStateSuccessWithFailures CleanRoomTaskRunResultState = `SUCCESS_WITH_FAILURES`

const CleanRoomTaskRunResultStateTimedout CleanRoomTaskRunResultState = `TIMEDOUT`

const CleanRoomTaskRunResultStateUpstreamCanceled CleanRoomTaskRunResultState = `UPSTREAM_CANCELED`

const CleanRoomTaskRunResultStateUpstreamEvicted CleanRoomTaskRunResultState = `UPSTREAM_EVICTED`

const CleanRoomTaskRunResultStateUpstreamFailed CleanRoomTaskRunResultState = `UPSTREAM_FAILED`

// String representation for [fmt.Print]
func (f *CleanRoomTaskRunResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomTaskRunResultState) Set(v string) error {
	switch v {
	case `CANCELED`, `DISABLED`, `EVICTED`, `EXCLUDED`, `FAILED`, `MAXIMUM_CONCURRENT_RUNS_REACHED`, `RUN_RESULT_STATE_UNSPECIFIED`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `TIMEDOUT`, `UPSTREAM_CANCELED`, `UPSTREAM_EVICTED`, `UPSTREAM_FAILED`:
		*f = CleanRoomTaskRunResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "DISABLED", "EVICTED", "EXCLUDED", "FAILED", "MAXIMUM_CONCURRENT_RUNS_REACHED", "RUN_RESULT_STATE_UNSPECIFIED", "SUCCESS", "SUCCESS_WITH_FAILURES", "TIMEDOUT", "UPSTREAM_CANCELED", "UPSTREAM_EVICTED", "UPSTREAM_FAILED"`, v)
	}
}

// Values returns all possible values for CleanRoomTaskRunResultState.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomTaskRunResultState) Values() []CleanRoomTaskRunResultState {
	return []CleanRoomTaskRunResultState{
		CleanRoomTaskRunResultStateCanceled,
		CleanRoomTaskRunResultStateDisabled,
		CleanRoomTaskRunResultStateEvicted,
		CleanRoomTaskRunResultStateExcluded,
		CleanRoomTaskRunResultStateFailed,
		CleanRoomTaskRunResultStateMaximumConcurrentRunsReached,
		CleanRoomTaskRunResultStateRunResultStateUnspecified,
		CleanRoomTaskRunResultStateSuccess,
		CleanRoomTaskRunResultStateSuccessWithFailures,
		CleanRoomTaskRunResultStateTimedout,
		CleanRoomTaskRunResultStateUpstreamCanceled,
		CleanRoomTaskRunResultStateUpstreamEvicted,
		CleanRoomTaskRunResultStateUpstreamFailed,
	}
}

// Type always returns CleanRoomTaskRunResultState to satisfy [pflag.Value] interface
func (f *CleanRoomTaskRunResultState) Type() string {
	return "CleanRoomTaskRunResultState"
}

// Stores the run state of the clean rooms notebook task.
type CleanRoomTaskRunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response. Note: Additional states might be
	// introduced in future releases.
	LifeCycleState CleanRoomTaskRunLifeCycleState `json:"life_cycle_state,omitempty"`
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states. Note: Additional states might be introduced in
	// future releases.
	ResultState CleanRoomTaskRunResultState `json:"result_state,omitempty"`
}

func (s *CleanRoomTaskRunState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Clean Rooms notebook task for V1 Clean Room service (GA). Replaces the
// deprecated CleanRoomNotebookTask (defined above) which was for V0 service.
type CleanRoomsNotebookTask struct {
	// The clean room that the notebook belongs to.
	CleanRoomName string `json:"clean_room_name"`
	// Checksum to validate the freshness of the notebook resource (i.e. the
	// notebook being run is the latest version). It can be fetched by calling
	// the :method:cleanroomassets/get API.
	Etag string `json:"etag,omitempty"`
	// Base parameters to be used for the clean room notebook job.
	NotebookBaseParameters map[string]string `json:"notebook_base_parameters,omitempty"`
	// Name of the notebook being run.
	NotebookName string `json:"notebook_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CleanRoomsNotebookTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CleanRoomsNotebookTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput struct {
	// The run state of the clean rooms notebook task.
	CleanRoomJobRunState *CleanRoomTaskRunState `json:"clean_room_job_run_state,omitempty"`
	// The notebook output for the clean room run
	NotebookOutput *NotebookOutput `json:"notebook_output,omitempty"`
	// Information on how to access the output schema for the clean room run
	OutputSchemaInfo *OutputSchemaInfo `json:"output_schema_info,omitempty"`
}

func (s *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ClusterInstance struct {
	// The canonical identifier for the cluster used by a run. This field is
	// always available for runs on existing clusters. For runs on new clusters,
	// it becomes available once the cluster is created. This value can be used
	// to view logs by browsing to `/#setting/sparkui/$cluster_id/driver-logs`.
	// The logs continue to be available after the run completes.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	ClusterId string `json:"cluster_id,omitempty"`
	// The canonical identifier for the Spark context used by a run. This field
	// is filled in once the run begins execution. This value can be used to
	// view the Spark UI by browsing to
	// `/#setting/sparkui/$cluster_id/$spark_context_id`. The Spark UI continues
	// to be available after the run has completed.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	SparkContextId string `json:"spark_context_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterInstance) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterInstance) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterSpec struct {
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string `json:"job_cluster_key,omitempty"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec `json:"new_cluster,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A single upstream completion that contributed to triggering the run, captured
// at trigger time
type CompletedJob struct {
	// The completion time of the upstream run
	CompletionTime *time.Time `json:"completion_time,omitempty"`
	// The ID of the upstream job
	JobId string `json:"job_id,omitempty"`
	// The ID of the upstream run whose completion contributed to the trigger
	RunId string `json:"run_id,omitempty"`
	// The termination code of the upstream run
	TerminationCode TerminationCodeCode `json:"termination_code,omitempty"`
	// The ID of the workspace containing the upstream job
	WorkspaceId string `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CompletedJob) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CompletedJob) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Compute struct {
	// Hardware accelerator configuration for Serverless GPU workloads.
	HardwareAccelerator compute.HardwareAcceleratorType `json:"hardware_accelerator,omitempty"`
}

func (s *Compute) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ComputeConfig struct {
	// IDof the GPU pool to use.
	GpuNodePoolId string `json:"gpu_node_pool_id,omitempty"`
	// GPU type.
	GpuType string `json:"gpu_type,omitempty"`
	// Number of GPUs.
	NumGpus int `json:"num_gpus"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComputeConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComputeConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// ComputeSpec: compute configuration — accelerator type and total accelerator
// count across all nodes.
type ComputeSpec struct {
	// Total number of accelerators across all nodes. Must be a positive
	// multiple of the per-node accelerator count encoded in `accelerator_type`.
	// For example, `GPU_8xH100` with `accelerator_count: 16` allocates 2 nodes
	// (8 GPUs per node).
	AcceleratorCount int `json:"accelerator_count"`
	// Hardware accelerator type (for example, `GPU_1xA10` or `GPU_8xH100`). The
	// number of accelerators per node is encoded in the enum value —
	// `GPU_8xH100` means 8 H100 GPUs per node.
	AcceleratorType ComputeSpecAcceleratorType `json:"accelerator_type"`
	// Legacy alias for `accelerator_pool_id`. Use `accelerator_pool_id` for new
	// configurations. If both fields are set, they must contain the same ID.
	ProvisionedCapacityId string `json:"provisioned_capacity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComputeSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComputeSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Hardware accelerator type for the AiRuntime workload. Per-node accelerator
// count is encoded in the value name (e.g. `GPU_8xH100` means 8 H100s per
// node).
type ComputeSpecAcceleratorType string

const ComputeSpecAcceleratorTypeGpu1xA10 ComputeSpecAcceleratorType = `GPU_1xA10`

const ComputeSpecAcceleratorTypeGpu1xH100 ComputeSpecAcceleratorType = `GPU_1xH100`

const ComputeSpecAcceleratorTypeGpu8xB300 ComputeSpecAcceleratorType = `GPU_8xB300`

const ComputeSpecAcceleratorTypeGpu8xH100 ComputeSpecAcceleratorType = `GPU_8xH100`

// String representation for [fmt.Print]
func (f *ComputeSpecAcceleratorType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ComputeSpecAcceleratorType) Set(v string) error {
	switch v {
	case `GPU_1xA10`, `GPU_1xH100`, `GPU_8xB300`, `GPU_8xH100`:
		*f = ComputeSpecAcceleratorType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GPU_1xA10", "GPU_1xH100", "GPU_8xB300", "GPU_8xH100"`, v)
	}
}

// Values returns all possible values for ComputeSpecAcceleratorType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ComputeSpecAcceleratorType) Values() []ComputeSpecAcceleratorType {
	return []ComputeSpecAcceleratorType{
		ComputeSpecAcceleratorTypeGpu1xA10,
		ComputeSpecAcceleratorTypeGpu1xH100,
		ComputeSpecAcceleratorTypeGpu8xB300,
		ComputeSpecAcceleratorTypeGpu8xH100,
	}
}

// Type always returns ComputeSpecAcceleratorType to satisfy [pflag.Value] interface
func (f *ComputeSpecAcceleratorType) Type() string {
	return "ComputeSpecAcceleratorType"
}

type Condition string

const ConditionAllUpdated Condition = `ALL_UPDATED`

const ConditionAnyUpdated Condition = `ANY_UPDATED`

// String representation for [fmt.Print]
func (f *Condition) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Condition) Set(v string) error {
	switch v {
	case `ALL_UPDATED`, `ANY_UPDATED`:
		*f = Condition(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_UPDATED", "ANY_UPDATED"`, v)
	}
}

// Values returns all possible values for Condition.
//
// There is no guarantee on the order of the values in the slice.
func (f *Condition) Values() []Condition {
	return []Condition{
		ConditionAllUpdated,
		ConditionAnyUpdated,
	}
}

// Type always returns Condition to satisfy [pflag.Value] interface
func (f *Condition) Type() string {
	return "Condition"
}

type ConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	Left string `json:"left"`
	// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
	// operands. This means that `“12.0” == “12”` will evaluate to
	// `false`. * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`,
	// `LESS_THAN_OR_EQUAL` operators perform numeric comparison of their
	// operands. `“12.0” >= “12”` will evaluate to `true`, `“10.0”
	// >= “12”` will evaluate to `false`.
	//
	// The boolean comparison to task values can be implemented with operators
	// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it
	// will be serialized to `“true”` or `“false”` for the comparison.
	Op ConditionTaskOp `json:"op"`
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right string `json:"right"`
}

func (s *ConditionTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
// operands. This means that `“12.0” == “12”` will evaluate to `false`.
// * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`
// operators perform numeric comparison of their operands. `“12.0” >=
// “12”` will evaluate to `true`, `“10.0” >= “12”` will evaluate to
// `false`.
//
// The boolean comparison to task values can be implemented with operators
// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it will
// be serialized to `“true”` or `“false”` for the comparison.
type ConditionTaskOp string

const ConditionTaskOpEqualTo ConditionTaskOp = `EQUAL_TO`

const ConditionTaskOpGreaterThan ConditionTaskOp = `GREATER_THAN`

const ConditionTaskOpGreaterThanOrEqual ConditionTaskOp = `GREATER_THAN_OR_EQUAL`

const ConditionTaskOpLessThan ConditionTaskOp = `LESS_THAN`

const ConditionTaskOpLessThanOrEqual ConditionTaskOp = `LESS_THAN_OR_EQUAL`

const ConditionTaskOpNotEqual ConditionTaskOp = `NOT_EQUAL`

// String representation for [fmt.Print]
func (f *ConditionTaskOp) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConditionTaskOp) Set(v string) error {
	switch v {
	case `EQUAL_TO`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`:
		*f = ConditionTaskOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL_TO", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "NOT_EQUAL"`, v)
	}
}

// Values returns all possible values for ConditionTaskOp.
//
// There is no guarantee on the order of the values in the slice.
func (f *ConditionTaskOp) Values() []ConditionTaskOp {
	return []ConditionTaskOp{
		ConditionTaskOpEqualTo,
		ConditionTaskOpGreaterThan,
		ConditionTaskOpGreaterThanOrEqual,
		ConditionTaskOpLessThan,
		ConditionTaskOpLessThanOrEqual,
		ConditionTaskOpNotEqual,
	}
}

// Type always returns ConditionTaskOp to satisfy [pflag.Value] interface
func (f *ConditionTaskOp) Type() string {
	return "ConditionTaskOp"
}

type Continuous struct {
	// Defines when platform-initiated maintenance may run for this job. If
	// unspecified, maintenance may run at any time.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty"`
	// Indicate whether the continuous execution of the job is paused or not.
	// Defaults to UNPAUSED.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
	// Indicate whether the continuous job is applying task level retries or
	// not. Defaults to NEVER.
	TaskRetryMode TaskRetryMode `json:"task_retry_mode,omitempty"`
}

func (s *Continuous) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Continuous trigger. Stripped-down counterpart to `ContinuousSettings`:
// `pause_status` is owned by the enclosing `TriggerConfiguration` and
// intentionally omitted here.
type ContinuousTriggerConfiguration struct {
	// Defines when platform-initiated maintenance may run for this trigger. If
	// unspecified, maintenance may run at any time.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty"`
	// Whether the continuous job applies task-level retries. Defaults to NEVER.
	TaskRetryMode TaskRetryMode `json:"task_retry_mode,omitempty"`
}

func (s *ContinuousTriggerConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ContinuousTriggerState struct {
	// The number of consecutive unsuccessful runs. Resets to `0` after a
	// healthy run and determines the exponential backoff before the next run.
	ConsecutiveFailures int `json:"consecutive_failures,omitempty"`
	// Whether the trigger is delaying the next run with exponential backoff
	// after repeated failures.
	IsBackingOff bool `json:"is_backing_off,omitempty"`
	// The time at which the trigger next attempts to start a run, in epoch
	// milliseconds.
	NextAttemptMs int64 `json:"next_attempt_ms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ContinuousTriggerState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContinuousTriggerState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateJob struct {
	// List of permissions to set on the job.
	AccessControlList []JobAccessControlRequest `json:"access_control_list,omitempty"`
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	//
	// Pipelines started by a continuous job also run continuously, regardless
	// of their own pipeline mode setting.
	Continuous *Continuous `json:"continuous,omitempty"`
	// Deployment information for jobs managed by external sources.
	Deployment *JobDeployment `json:"deployment,omitempty"`
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	Description string `json:"description,omitempty"`
	// The window the run aims to finish within, in hours, for a Cost-optimized
	// serverless run. Required only when `performance_target` is
	// `COST_OPTIMIZED`. Must be one of the supported values: `NO_FLEXIBILITY`,
	// `4h`, `12h`, `24h`, `72h`. `NO_FLEXIBILITY` means the run uses the
	// default cost-optimized performance level without additional cost-saving
	// optimizations.
	DurationTarget string `json:"duration_target,omitempty"`
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode JobEditMode `json:"edit_mode,omitempty"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// Named environment-variable entries that tasks can reference by key from
	// `TaskSettings.environment_variables_key`. Each entry's `spec` holds
	// inline `variables` and optional `.env` `files`. Maximum 10 entries per
	// job. A task can reference at most one entry from this list.
	EnvironmentVariables []JobEnvironmentVariables `json:"environment_variables,omitempty"`
	// A list of task execution environment specifications that can be
	// referenced by tasks that use serverless compute or a compute resource
	// that uses Environments mode.
	//
	// For notebook tasks that use serverless compute or a compute resource that
	// uses Environments mode, if the `environment_key` is not specified, the
	// notebook environment will be used if present. If a jobs environment is
	// specified, it will override the notebook environment.
	//
	// For other tasks that use serverless compute or a compute resource that
	// uses Environments mode, the task environment is required to be specified
	// using `environment_key` in the task settings.
	Environments []JobEnvironment `json:"environments,omitempty"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format Format `json:"format,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *GitSource `json:"git_source,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `json:"health,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// An optional maximum allowed number of concurrent runs of the job. Set
	// this value if you want to be able to execute multiple runs of the same
	// job concurrently. This is useful for example if you trigger your job on a
	// frequent schedule and want to allow consecutive runs to overlap with each
	// other, or if you want to trigger multiple runs which differ by their
	// input parameters. This setting affects only new runs. For example,
	// suppose the job’s concurrency is 4 and there are 4 concurrent active
	// runs. Then setting the concurrency to 3 won’t kill any of the active
	// runs. However, from then on, new runs are skipped unless there are fewer
	// than 3 active runs. This value cannot exceed 1000. Setting this value to
	// `0` causes all new runs to be skipped.
	MaxConcurrentRuns int `json:"max_concurrent_runs,omitempty"`
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name string `json:"name,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *JobNotificationSettings `json:"notification_settings,omitempty"`
	// Job-level parameter definitions
	Parameters []JobParameterDefinition `json:"parameters,omitempty"`
	// Path of the job parent folder in workspace file tree. If absent, the job
	// doesn't have a parent folder.
	//
	// Example: /Workspace/Projects/my_project
	ParentPath string `json:"parent_path,omitempty"`
	// The performance mode on a serverless job. This field determines the level
	// of compute performance or cost-efficiency for the run. The performance
	// target does not apply to tasks that run on Serverless GPU compute.
	//
	// * `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance. * `STANDARD`:
	// Enables cost-efficient execution of serverless workloads. *
	// `COST_OPTIMIZED`: Enables lower job costs by optimizing compute for your
	// selected target duration time. Must provide a duration target.
	PerformanceTarget PerformanceTarget `json:"performance_target,omitempty"`
	// The queue settings of the job.
	Queue *QueueSettings `json:"queue,omitempty"`
	// The user or service principal that the job runs as, if specified in the
	// request. This field indicates the explicit configuration of `run_as` for
	// the job. To find the value in all cases, explicit or implicit, use
	// `run_as_user_name`.
	RunAs *JobRunAs `json:"run_as,omitempty"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags map[string]string `json:"tags,omitempty"`
	// A list of task specifications to be executed by this job. It supports up
	// to 1000 elements in write endpoints (:method:jobs/create,
	// :method:jobs/reset, :method:jobs/update, :method:jobs/submit). Read
	// endpoints return only 100 tasks. If more than 100 tasks are available,
	// you can paginate through them using :method:jobs/get. Use the
	// `next_page_token` field at the object root to determine if more results
	// are available.
	Tasks []Task `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger *TriggerSettings `json:"trigger,omitempty"`
	// List of triggers attached to this job. A run starts when any active
	// trigger evaluates to true. Cannot be set in the same request as the
	// legacy `schedule`, `trigger`, or `continuous` fields. Gated behind the
	// "Multiple Triggers" feature preview.
	Triggers []TriggerConfiguration `json:"triggers,omitempty"`
	// The id of the user specified usage policy to use for this job. If not
	// specified, a default usage policy may be applied when creating or
	// modifying the job. See `effective_usage_policy_id` for the usage policy
	// used by this workload.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateJob) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateJob) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Job was created successfully
type CreateResponse struct {
	// The canonical identifier for the newly created job.
	JobId int64 `json:"job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
	// A Cron expression using Quartz syntax that describes the schedule for a
	// job. See [Cron Trigger] for details. This field is required.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// SQL condition that must be satisfied before a scheduled run is triggered.
	// The condition is evaluated after the cron expression fires and must
	// return a truthy result for the run to proceed.
	SqlCondition *SqlConditionConfiguration `json:"sql_condition,omitempty"`
	// A Java timezone ID. The schedule for a job is resolved with respect to
	// this timezone. See [Java TimeZone] for details. This field is required.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId string `json:"timezone_id"`
}

func (s *CronSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Cron schedule trigger. Stripped-down counterpart to `CronSchedule`:
// `pause_status` and `sql_condition` are owned by the enclosing
// `TriggerConfiguration` and intentionally omitted here.
type CronTriggerConfiguration struct {
	// A Cron expression using Quartz syntax that describes the schedule for
	// this trigger. See [Cron Trigger] for details.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// A Java timezone ID. The schedule is resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId string `json:"timezone_id"`
}

func (s *CronTriggerConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DashboardPageSnapshot struct {
	PageDisplayName string `json:"page_display_name,omitempty"`

	WidgetErrorDetails []WidgetErrorDetail `json:"widget_error_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DashboardPageSnapshot) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardPageSnapshot) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configures the Lakeview Dashboard job task type.
type DashboardTask struct {
	// The identifier of the dashboard to refresh.
	DashboardId string `json:"dashboard_id,omitempty"`
	// Dashboard task parameters. Used to apply dashboard filter values during
	// dashboard task execution. Parameter values get applied to any dashboard
	// filters that have a matching URL identifier as the parameter key. The
	// parameter value format is dependent on the filter type: - For text and
	// single-select filters, provide a single value (e.g. `"value"`) - For date
	// and datetime filters, provide the value in ISO 8601 format (e.g.
	// `"2000-01-01T00:00:00"`) - For multi-select filters, provide a JSON array
	// of values (e.g. `"[\"value1\",\"value2\"]"`) - For range and date range
	// filters, provide a JSON object with `start` and `end` (e.g.
	// `"{\"start\":\"1\",\"end\":\"10\"}"`)
	Filters map[string]string `json:"filters,omitempty"`
	// Optional: subscription configuration for sending the dashboard snapshot.
	Subscription *Subscription `json:"subscription,omitempty"`
	// Optional: The warehouse id to execute the dashboard with for the
	// schedule. If not specified, the default warehouse of the dashboard will
	// be used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DashboardTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DashboardTaskOutput struct {
	// Should only be populated for manual PDF download jobs.
	PageSnapshots []DashboardPageSnapshot `json:"page_snapshots,omitempty"`
}

func (s *DashboardTaskOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Days of week that can be referenced by Jobs scheduling settings.
type DayOfWeek string

const DayOfWeekFriday DayOfWeek = `FRIDAY`

const DayOfWeekMonday DayOfWeek = `MONDAY`

const DayOfWeekSaturday DayOfWeek = `SATURDAY`

const DayOfWeekSunday DayOfWeek = `SUNDAY`

const DayOfWeekThursday DayOfWeek = `THURSDAY`

const DayOfWeekTuesday DayOfWeek = `TUESDAY`

const DayOfWeekWednesday DayOfWeek = `WEDNESDAY`

// String representation for [fmt.Print]
func (f *DayOfWeek) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DayOfWeek) Set(v string) error {
	switch v {
	case `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`:
		*f = DayOfWeek(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY", "WEDNESDAY"`, v)
	}
}

// Values returns all possible values for DayOfWeek.
//
// There is no guarantee on the order of the values in the slice.
func (f *DayOfWeek) Values() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
	}
}

// Type always returns DayOfWeek to satisfy [pflag.Value] interface
func (f *DayOfWeek) Type() string {
	return "DayOfWeek"
}

// Format of response retrieved from dbt Cloud, for inclusion in output
// Deprecated in favor of DbtPlatformJobRunStep
type DbtCloudJobRunStep struct {
	// Orders the steps in the job
	Index int `json:"index,omitempty"`
	// Output of the step
	Logs string `json:"logs,omitempty"`
	// Name of the step in the job
	Name string `json:"name,omitempty"`
	// State of the step
	Status DbtPlatformRunStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DbtCloudJobRunStep) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DbtCloudJobRunStep) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Deprecated in favor of DbtPlatformTask
type DbtCloudTask struct {
	// The resource name of the UC connection that authenticates the dbt Cloud
	// for this task
	ConnectionResourceName string `json:"connection_resource_name,omitempty"`
	// Id of the dbt Cloud job to be triggered
	DbtCloudJobId int64 `json:"dbt_cloud_job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DbtCloudTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DbtCloudTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Deprecated in favor of DbtPlatformTaskOutput
type DbtCloudTaskOutput struct {
	// Id of the job run in dbt Cloud
	DbtCloudJobRunId int64 `json:"dbt_cloud_job_run_id,omitempty"`
	// Steps of the job run as received from dbt Cloud
	DbtCloudJobRunOutput []DbtCloudJobRunStep `json:"dbt_cloud_job_run_output,omitempty"`
	// Url where full run details can be viewed
	DbtCloudJobRunUrl string `json:"dbt_cloud_job_run_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DbtCloudTaskOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DbtCloudTaskOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DbtOutput struct {
	// An optional map of headers to send when retrieving the artifact from the
	// `artifacts_link`.
	ArtifactsHeaders map[string]string `json:"artifacts_headers,omitempty"`
	// A pre-signed URL to download the (compressed) dbt artifacts. This link is
	// valid for a limited time (30 minutes). This information is only available
	// after the run has finished.
	ArtifactsLink string `json:"artifacts_link,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DbtOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DbtOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Format of response retrieved from dbt platform, for inclusion in output
type DbtPlatformJobRunStep struct {
	// Orders the steps in the job
	Index int `json:"index,omitempty"`
	// Output of the step
	Logs string `json:"logs,omitempty"`
	// Whether the logs of this step have been truncated. If true, the logs has
	// been truncated to 10000 characters.
	LogsTruncated bool `json:"logs_truncated,omitempty"`
	// Name of the step in the job
	Name string `json:"name,omitempty"`
	// Whether the name of the job has been truncated. If true, the name has
	// been truncated to 100 characters.
	NameTruncated bool `json:"name_truncated,omitempty"`
	// State of the step
	Status DbtPlatformRunStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DbtPlatformJobRunStep) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DbtPlatformJobRunStep) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response enumeration from calling the dbt platform API, for inclusion in
// output
type DbtPlatformRunStatus string

const DbtPlatformRunStatusCancelled DbtPlatformRunStatus = `CANCELLED`

const DbtPlatformRunStatusError DbtPlatformRunStatus = `ERROR`

const DbtPlatformRunStatusQueued DbtPlatformRunStatus = `QUEUED`

const DbtPlatformRunStatusRunning DbtPlatformRunStatus = `RUNNING`

const DbtPlatformRunStatusStarting DbtPlatformRunStatus = `STARTING`

const DbtPlatformRunStatusSuccess DbtPlatformRunStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *DbtPlatformRunStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DbtPlatformRunStatus) Set(v string) error {
	switch v {
	case `CANCELLED`, `ERROR`, `QUEUED`, `RUNNING`, `STARTING`, `SUCCESS`:
		*f = DbtPlatformRunStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "ERROR", "QUEUED", "RUNNING", "STARTING", "SUCCESS"`, v)
	}
}

// Values returns all possible values for DbtPlatformRunStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *DbtPlatformRunStatus) Values() []DbtPlatformRunStatus {
	return []DbtPlatformRunStatus{
		DbtPlatformRunStatusCancelled,
		DbtPlatformRunStatusError,
		DbtPlatformRunStatusQueued,
		DbtPlatformRunStatusRunning,
		DbtPlatformRunStatusStarting,
		DbtPlatformRunStatusSuccess,
	}
}

// Type always returns DbtPlatformRunStatus to satisfy [pflag.Value] interface
func (f *DbtPlatformRunStatus) Type() string {
	return "DbtPlatformRunStatus"
}

type DbtPlatformTask struct {
	// The resource name of the UC connection that authenticates the dbt
	// platform for this task
	ConnectionResourceName string `json:"connection_resource_name,omitempty"`
	// Id of the dbt platform job to be triggered. Specified as a string for
	// maximum compatibility with clients.
	DbtPlatformJobId string `json:"dbt_platform_job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DbtPlatformTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DbtPlatformTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DbtPlatformTaskOutput struct {
	// Id of the job run in dbt platform. Specified as a string for maximum
	// compatibility with clients.
	DbtPlatformJobRunId string `json:"dbt_platform_job_run_id,omitempty"`
	// Steps of the job run as received from dbt platform
	DbtPlatformJobRunOutput []DbtPlatformJobRunStep `json:"dbt_platform_job_run_output,omitempty"`
	// Url where full run details can be viewed
	DbtPlatformJobRunUrl string `json:"dbt_platform_job_run_url,omitempty"`
	// Whether the number of steps in the output has been truncated. If true,
	// the output will contain the first 20 steps of the output.
	StepsTruncated bool `json:"steps_truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DbtPlatformTaskOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DbtPlatformTaskOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DbtTask struct {
	// Optional name of the catalog to use. The value is the top level in the
	// 3-level namespace of Unity Catalog (catalog / schema / relation). The
	// catalog value can only be specified if a warehouse_id is specified.
	// Requires dbt-databricks >= 1.1.1.
	Catalog string `json:"catalog,omitempty"`
	// A list of dbt commands to execute. All commands must start with `dbt`.
	// This parameter must not be empty. A maximum of up to 10 commands can be
	// provided.
	Commands []string `json:"commands"`
	// Optional (relative) path to the profiles directory. Can only be specified
	// if no warehouse_id is specified. If no warehouse_id is specified and this
	// folder is unset, the root directory is used.
	ProfilesDirectory string `json:"profiles_directory,omitempty"`
	// Path to the project directory. Optional for Git sourced tasks, in which
	// case if no value is provided, the root of the Git repository is used.
	ProjectDirectory string `json:"project_directory,omitempty"`
	// Optional schema to write to. This parameter is only used when a
	// warehouse_id is also provided. If not provided, the `default` schema is
	// used.
	Schema string `json:"schema,omitempty"`
	// Optional location type of the project directory. When set to `WORKSPACE`,
	// the project will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the project will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: Project is located in Databricks workspace. * `GIT`:
	// Project is located in cloud Git provider.
	Source Source `json:"source,omitempty"`
	// ID of the SQL warehouse to connect to. If provided, we automatically
	// generate and provide the profile and connection details to dbt. It can be
	// overridden on a per-command basis by using the `--profiles-dir` command
	// line argument.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DbtTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DbtTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteJob struct {
	// The canonical identifier of the job to delete. This field is required.
	JobId int64 `json:"job_id"`
}

func (s *DeleteJob) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteRun struct {
	// ID of the run to delete.
	RunId int64 `json:"run_id"`
}

func (s *DeleteRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A Unity Catalog Delta table input source. The function is applied to each
// row.
type DeltaTableSource struct {
	// Name of the column holding the content the function is applied to.
	InputColumn string `json:"input_column"`
	// Columns that together identify a row across incremental runs, so results
	// can be matched back to their source rows. When unset, the task derives an
	// implicit key from Delta row tracking. If row tracking is also disabled
	// there is no stable key, so each run reprocesses the entire source (a full
	// refresh) rather than only the changed rows; set a key column or enable
	// row tracking to keep runs incremental.
	MergeKeyColumns []string `json:"merge_key_columns,omitempty"`
	// Fully qualified three-part name of the source table
	// (catalog.schema.table).
	TableName string `json:"table_name"`
}

func (s *DeltaTableSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// DeploymentSpec: configuration for one deployment within an AiRuntimeTask.
// Each entry in `AiRuntimeTask.deployments` describes a group of nodes that
// share the same command and compute. Many single-program training algorithms
// use a single entry where every node runs the same command; role-split
// workloads (driver + worker, parameter server, separate eval node, etc.) use
// multiple entries.
type DeploymentSpec struct {
	// Workspace path of the script to run on each node in this deployment.
	// Upload the script to this path and supply the path here. When the task
	// runs, the file at this path is run on each node; if it fails, the task
	// fails with its exit code.
	//
	// Example script contents:
	//
	// ```bash # Plain Python: python train.py --epochs 10
	//
	// # Multi-GPU via accelerate: accelerate launch train.py --config
	// config.yaml
	//
	// # Distributed via torchrun: torchrun --nproc_per_node=8 train.py ```
	CommandPath string `json:"command_path,omitempty"`
	// Compute resources allocated to each node in this deployment.
	Compute ComputeSpec `json:"compute"`
	// Optional human-readable name for this deployment (for example, `driver`,
	// `worker`, `param_server`). Used for log and UI display. Distinct names
	// are recommended so deployments can be told apart, but uniqueness is not
	// enforced.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeploymentSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeploymentSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Represents a change to the job cluster's settings that would be required for
// the job clusters to become compliant with their policies.
type EnforcePolicyComplianceForJobResponseJobClusterSettingsChange struct {
	// The field where this change would be made, prepended with the job cluster
	// key.
	Field string `json:"field,omitempty"`
	// The new value of this field after enforcing policy compliance (either a
	// number, a boolean, or a string) converted to a string. This is intended
	// to be read by a human. The typed new value of this field can be retrieved
	// by reading the settings field in the API response.
	NewValue string `json:"new_value,omitempty"`
	// The previous value of this field before enforcing policy compliance
	// (either a number, a boolean, or a string) converted to a string. This is
	// intended to be read by a human. The type of the field can be retrieved by
	// reading the settings field in the API response.
	PreviousValue string `json:"previous_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnforcePolicyComplianceRequest struct {
	// The ID of the job you want to enforce policy compliance on.
	JobId int64 `json:"job_id"`
	// If set, previews changes made to the job to comply with its policy, but
	// does not update the job.
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnforcePolicyComplianceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnforcePolicyComplianceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnforcePolicyComplianceResponse struct {
	// Whether any changes have been made to the job cluster settings for the
	// job to become compliant with its policies.
	HasChanges bool `json:"has_changes,omitempty"`
	// A list of job cluster changes that have been made to the job’s cluster
	// settings in order for all job clusters to become compliant with their
	// policies.
	JobClusterChanges []EnforcePolicyComplianceForJobResponseJobClusterSettingsChange `json:"job_cluster_changes,omitempty"`
	// Updated job settings after policy enforcement. Policy enforcement only
	// applies to job clusters that are created when running the job (which are
	// specified in new_cluster) and does not apply to existing all-purpose
	// clusters. Updated job settings are derived by applying policy default
	// values to the existing job clusters in order to satisfy policy
	// requirements.
	Settings *JobSettings `json:"settings,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnforcePolicyComplianceResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnforcePolicyComplianceResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Run was exported successfully.
type ExportRunOutput struct {
	// The exported content in HTML format (one for every view item). To extract
	// the HTML notebook from the JSON response, download and run this [Python
	// script](/_static/examples/extract.py).
	Views []ViewItem `json:"views,omitempty"`
}

func (s *ExportRunOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ExportRunRequest struct {
	// The canonical identifier for the run. This field is required.
	RunId int64 `json:"-" url:"run_id"`
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	ViewsToExport ViewsToExport `json:"-" url:"views_to_export,omitempty"`
}

func (s *ExportRunRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type FileArrivalTriggerConfiguration struct {
	// If set, the trigger starts a run only after the specified amount of time
	// passed since the last time the trigger fired. The minimum allowed value
	// is 60 seconds
	MinTimeBetweenTriggersSeconds int `json:"min_time_between_triggers_seconds,omitempty"`
	// URL to be monitored for file arrivals. The path must point to the root or
	// a subpath of the external location.
	Url string `json:"url"`
	// If set, the trigger starts a run only after no file activity has occurred
	// for the specified amount of time. This makes it possible to wait for a
	// batch of incoming files to arrive before triggering a run. The minimum
	// allowed value is 60 seconds.
	WaitAfterLastChangeSeconds int `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileArrivalTriggerConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileArrivalTriggerConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FileArrivalTriggerState struct {
	// Indicates whether the trigger leverages file events to detect file
	// arrivals.
	UsingFileEvents bool `json:"using_file_events,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileArrivalTriggerState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileArrivalTriggerState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ForEachStats struct {
	// Sample of 3 most common error messages occurred during the iteration.
	ErrorMessageStats []ForEachTaskErrorMessageStats `json:"error_message_stats,omitempty"`
	// Describes stats of the iteration. Only latest retries are considered.
	TaskRunStats *ForEachTaskTaskRunStats `json:"task_run_stats,omitempty"`
}

func (s *ForEachStats) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ForEachTask struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	Concurrency int `json:"concurrency,omitempty"`
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs string `json:"inputs"`
	// Configuration for the task that will be run for each element in the array
	Task Task `json:"task"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ForEachTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForEachTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ForEachTaskErrorMessageStats struct {
	// Describes the count of such error message encountered during the
	// iterations.
	Count int `json:"count,omitempty"`
	// Describes the error message occurred during the iterations.
	ErrorMessage string `json:"error_message,omitempty"`
	// Describes the termination reason for the error message.
	TerminationCategory string `json:"termination_category,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ForEachTaskErrorMessageStats) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForEachTaskErrorMessageStats) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ForEachTaskTaskRunStats struct {
	// Describes the iteration runs having an active lifecycle state or an
	// active run sub state.
	ActiveIterations int `json:"active_iterations,omitempty"`
	// Describes the number of failed and succeeded iteration runs.
	CompletedIterations int `json:"completed_iterations,omitempty"`
	// Describes the number of failed iteration runs.
	FailedIterations int `json:"failed_iterations,omitempty"`
	// Describes the number of iteration runs that have been scheduled.
	ScheduledIterations int `json:"scheduled_iterations,omitempty"`
	// Describes the number of succeeded iteration runs.
	SucceededIterations int `json:"succeeded_iterations,omitempty"`
	// Describes the length of the list of items to iterate over.
	TotalIterations int `json:"total_iterations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ForEachTaskTaskRunStats) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForEachTaskTaskRunStats) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Format string

const FormatMultiTask Format = `MULTI_TASK`

const FormatSingleTask Format = `SINGLE_TASK`

// String representation for [fmt.Print]
func (f *Format) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Format) Set(v string) error {
	switch v {
	case `MULTI_TASK`, `SINGLE_TASK`:
		*f = Format(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MULTI_TASK", "SINGLE_TASK"`, v)
	}
}

// Values returns all possible values for Format.
//
// There is no guarantee on the order of the values in the slice.
func (f *Format) Values() []Format {
	return []Format{
		FormatMultiTask,
		FormatSingleTask,
	}
}

// Type always returns Format to satisfy [pflag.Value] interface
func (f *Format) Type() string {
	return "Format"
}

// DEPRECATED — use `AiRuntimeTask` for all new BYOT multi-node GPU workloads
// (see ai_runtime_task.proto). `AiRuntimeTask` is the only supported BYOT task
// type for new workloads; this proto is retained only for AIR CLI (fka SGCLI)
// pywheel backwards compatibility and will be removed once the pywheel →
// databricks-cli migration completes (post- PuPr).
type GenAiComputeTask struct {
	// Command launcher to run the actual script, e.g. bash, python etc.
	Command string `json:"command,omitempty"`

	Compute *ComputeConfig `json:"compute,omitempty"`
	// Runtime image
	DlRuntimeImage string `json:"dl_runtime_image"`
	// Optional string containing the name of the MLflow experiment to log the
	// run to. If name is not found, backend will create the mlflow experiment
	// using the name.
	MlflowExperimentName string `json:"mlflow_experiment_name,omitempty"`
	// Optional location type of the training script. When set to `WORKSPACE`,
	// the script will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the script will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`: Script
	// is located in Databricks workspace. * `GIT`: Script is located in cloud
	// Git provider.
	Source Source `json:"source,omitempty"`
	// The training script file path to be executed. Cloud file URIs (such as
	// dbfs:/, s3:/, adls:/, gcs:/) and workspace paths are supported. For
	// python files stored in the Databricks workspace, the path must be
	// absolute and begin with `/`. For files stored in a remote repository, the
	// path must be relative. This field is required.
	TrainingScriptPath string `json:"training_script_path,omitempty"`
	// Optional string containing model parameters passed to the training script
	// in yaml format. If present, then the content in yaml_parameters_file_path
	// will be ignored.
	YamlParameters string `json:"yaml_parameters,omitempty"`
	// Optional path to a YAML file containing model parameters passed to the
	// training script.
	YamlParametersFilePath string `json:"yaml_parameters_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenAiComputeTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenAiComputeTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetJobPermissionLevelsRequest struct {
	// The job for which to get or manage permissions.
	JobId string `json:"-" url:"-"`
}

func (s *GetJobPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetJobPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []JobPermissionsDescription `json:"permission_levels,omitempty"`
}

func (s *GetJobPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetJobPermissionsRequest struct {
	// The job for which to get or manage permissions.
	JobId string `json:"-" url:"-"`
}

func (s *GetJobPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetJobRequest struct {
	// Flag that indicates that trigger state should be included in the
	// response.
	IncludeTriggerState bool `json:"-" url:"include_trigger_state,omitempty"`
	// The canonical identifier of the job to retrieve information about. This
	// field is required.
	JobId int64 `json:"-" url:"job_id"`
	// Use `next_page_token` returned from the previous GetJob response to
	// request the next page of the job's array properties.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetJobRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetJobRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPolicyComplianceRequest struct {
	// The ID of the job whose compliance status you are requesting.
	JobId int64 `json:"-" url:"job_id"`
}

func (s *GetPolicyComplianceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPolicyComplianceResponse struct {
	// Whether the job is compliant with its policies or not. Jobs could be out
	// of compliance if a policy they are using was updated after the job was
	// last edited and some of its job clusters no longer comply with their
	// updated policies.
	IsCompliant bool `json:"is_compliant,omitempty"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	Violations map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPolicyComplianceResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPolicyComplianceResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetRunOutputRequest struct {
	// The canonical identifier for the run.
	RunId int64 `json:"-" url:"run_id"`
}

func (s *GetRunOutputRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetRunRequest struct {
	// Whether to include the repair history in the response.
	IncludeHistory bool `json:"-" url:"include_history,omitempty"`
	// Whether to include resolved parameter values in the response.
	IncludeResolvedValues bool `json:"-" url:"include_resolved_values,omitempty"`
	// Use `next_page_token` returned from the previous GetRun response to
	// request the next page of the run's array properties.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The canonical identifier of the run for which to retrieve the metadata.
	// This field is required.
	RunId int64 `json:"-" url:"run_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetRunRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRunRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GitProvider string

const GitProviderAwsCodeCommit GitProvider = `awsCodeCommit`

const GitProviderAzureDevOpsServices GitProvider = `azureDevOpsServices`

const GitProviderBitbucketCloud GitProvider = `bitbucketCloud`

const GitProviderBitbucketServer GitProvider = `bitbucketServer`

const GitProviderGitHub GitProvider = `gitHub`

const GitProviderGitHubEnterprise GitProvider = `gitHubEnterprise`

const GitProviderGitLab GitProvider = `gitLab`

const GitProviderGitLabEnterpriseEdition GitProvider = `gitLabEnterpriseEdition`

// String representation for [fmt.Print]
func (f *GitProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GitProvider) Set(v string) error {
	switch v {
	case `awsCodeCommit`, `azureDevOpsServices`, `bitbucketCloud`, `bitbucketServer`, `gitHub`, `gitHubEnterprise`, `gitLab`, `gitLabEnterpriseEdition`:
		*f = GitProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "awsCodeCommit", "azureDevOpsServices", "bitbucketCloud", "bitbucketServer", "gitHub", "gitHubEnterprise", "gitLab", "gitLabEnterpriseEdition"`, v)
	}
}

// Values returns all possible values for GitProvider.
//
// There is no guarantee on the order of the values in the slice.
func (f *GitProvider) Values() []GitProvider {
	return []GitProvider{
		GitProviderAwsCodeCommit,
		GitProviderAzureDevOpsServices,
		GitProviderBitbucketCloud,
		GitProviderBitbucketServer,
		GitProviderGitHub,
		GitProviderGitHubEnterprise,
		GitProviderGitLab,
		GitProviderGitLabEnterpriseEdition,
	}
}

// Type always returns GitProvider to satisfy [pflag.Value] interface
func (f *GitProvider) Type() string {
	return "GitProvider"
}

// Read-only state of the remote repository at the time the job was run. This
// field is only included on job runs.
type GitSnapshot struct {
	// Commit that was used to execute the run. If git_branch was specified,
	// this points to the HEAD of the branch at the time of the run; if git_tag
	// was specified, this points to the commit the tag points to.
	UsedCommit string `json:"used_commit,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GitSnapshot) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GitSnapshot) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An optional specification for a remote Git repository containing the source
// code used by tasks. Version-controlled source code is supported by notebook,
// dbt, Python script, and SQL File tasks.
//
// If `git_source` is set, these tasks retrieve the file from the remote
// repository by default. However, this behavior can be overridden by setting
// `source` to `WORKSPACE` on the task.
//
// Note: dbt and SQL File tasks support only version-controlled sources. If dbt
// or SQL File tasks are used, `git_source` must be defined on the job.
type GitSource struct {
	// Name of the branch to be checked out and used by this job. This field
	// cannot be specified in conjunction with git_tag or git_commit.
	GitBranch string `json:"git_branch,omitempty"`
	// Commit to be checked out and used by this job. This field cannot be
	// specified in conjunction with git_branch or git_tag.
	GitCommit string `json:"git_commit,omitempty"`
	// Unique identifier of the service used to host the Git repository. The
	// value is case insensitive.
	GitProvider GitProvider `json:"git_provider"`
	// Read-only state of the remote repository at the time the job was run.
	// This field is only included on job runs.
	GitSnapshot *GitSnapshot `json:"git_snapshot,omitempty"`
	// Name of the tag to be checked out and used by this job. This field cannot
	// be specified in conjunction with git_branch or git_commit.
	GitTag string `json:"git_tag,omitempty"`
	// URL of the repository to be cloned by this job.
	GitUrl string `json:"git_url"`
	// The source of the job specification in the remote repository when the job
	// is source controlled.
	JobSource *JobSource `json:"job_source,omitempty"`
	// Sparse checkout settings for the repository. When set, only the
	// directories listed in `patterns` are checked out.
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GitSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GitSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Job was retrieved successfully.
type Job struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime int64 `json:"created_time,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// The id of the usage policy used by this job for cost attribution
	// purposes.
	EffectiveUsagePolicyId string `json:"effective_usage_policy_id,omitempty"`
	// Indicates if the job has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/get endpoint.
	// It is only relevant for API 2.2 :method:jobs/list requests with
	// `expand_tasks=true`.
	HasMore bool `json:"has_more,omitempty"`
	// The canonical identifier for this job.
	JobId int64 `json:"job_id,omitempty"`
	// A token that can be used to list the next page of array properties.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The email of an active workspace user or the application ID of a service
	// principal that the job runs as. This value can be changed by setting the
	// `run_as` field when creating or updating a job.
	//
	// By default, `run_as_user_name` is based on the current job settings and
	// is set to the creator of the job if job access control is disabled or to
	// the user with the `is_owner` permission if job access control is enabled.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *JobSettings `json:"settings,omitempty"`
	// Per-trigger runtime information for the multi-trigger surface. Same
	// length and order as `JobSettings.triggers`; `trigger_details[i]`
	// corresponds to `triggers[i]`. Sub-fields (`state`, `history`) are
	// populated independently based on the `GetJob.include_trigger_state` /
	// `include_trigger_history` flags.
	TriggerDetails []TriggerDetails `json:"trigger_details,omitempty"`
	// State of the trigger associated with the job.
	TriggerState *TriggerStateProto `json:"trigger_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Job) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Job) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobAccessControlResponse struct {
	// All permissions.
	AllPermissions []JobPermission `json:"all_permissions,omitempty"`
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

func (s *JobAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobCluster struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	JobClusterKey string `json:"job_cluster_key"`
	// If new_cluster, a description of a cluster that is created for each task.
	NewCluster *compute.ClusterSpec `json:"new_cluster,omitempty"`
	// Serverless compute settings for this cluster. At most one job cluster per
	// job may set this field; the rate limit defined on the referenced policy
	// applies across all tasks bound to this cluster.
	ServerlessCompute *JobClusterServerlessCompute `json:"serverless_compute,omitempty"`
	// The ID of the serverless compute object to bind this cluster to. At most
	// one JobCluster per job may set this field; the rate limit defined on the
	// referenced serverless compute applies across all tasks bound to this
	// cluster.
	ServerlessComputeId string `json:"serverless_compute_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobCluster) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobCluster) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Serverless compute settings for a job cluster.
type JobClusterServerlessCompute struct {
	// The ID of the serverless compute policy to apply to this cluster.
	PolicyId string `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobClusterServerlessCompute) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobClusterServerlessCompute) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A trigger that starts a run after every upstream job completes with a
// configured status.
type JobCompletionTriggerConfiguration struct {
	// The status that all upstream jobs must have before this trigger fires.
	// Defaults to SUCCESS.
	IfStatus JobCompletionTriggerConfigurationIfStatus `json:"if_status,omitempty"`
	// The upstream jobs to monitor. Each job must be in a workspace attached to
	// the same Unity Catalog metastore as the downstream job. A maximum of 10
	// jobs can be monitored.
	Jobs []JobCompletionTriggerConfigurationMonitoredJob `json:"jobs"`
}

func (s *JobCompletionTriggerConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The status that all upstream jobs must have before a job completion trigger
// fires.
type JobCompletionTriggerConfigurationIfStatus string

const JobCompletionTriggerConfigurationIfStatusDone JobCompletionTriggerConfigurationIfStatus = `DONE`

const JobCompletionTriggerConfigurationIfStatusFailed JobCompletionTriggerConfigurationIfStatus = `FAILED`

const JobCompletionTriggerConfigurationIfStatusSuccess JobCompletionTriggerConfigurationIfStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *JobCompletionTriggerConfigurationIfStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobCompletionTriggerConfigurationIfStatus) Set(v string) error {
	switch v {
	case `DONE`, `FAILED`, `SUCCESS`:
		*f = JobCompletionTriggerConfigurationIfStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DONE", "FAILED", "SUCCESS"`, v)
	}
}

// Values returns all possible values for JobCompletionTriggerConfigurationIfStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobCompletionTriggerConfigurationIfStatus) Values() []JobCompletionTriggerConfigurationIfStatus {
	return []JobCompletionTriggerConfigurationIfStatus{
		JobCompletionTriggerConfigurationIfStatusDone,
		JobCompletionTriggerConfigurationIfStatusFailed,
		JobCompletionTriggerConfigurationIfStatusSuccess,
	}
}

// Type always returns JobCompletionTriggerConfigurationIfStatus to satisfy [pflag.Value] interface
func (f *JobCompletionTriggerConfigurationIfStatus) Type() string {
	return "JobCompletionTriggerConfigurationIfStatus"
}

// An upstream job to monitor.
type JobCompletionTriggerConfigurationMonitoredJob struct {
	// The ID of the upstream job.
	JobId string `json:"job_id"`
	// The ID of the workspace containing the upstream job. This can differ from
	// the caller's workspace and downstream services enforce correct
	// entitlements and CAN_VIEW permissions for the user + job combination.
	WorkspaceId string `json:"workspace_id"`
}

func (s *JobCompletionTriggerConfigurationMonitoredJob) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Contains the immutable set of upstream completions that triggered the run, in
// configured jobs order
type JobCompletionTriggerRunInfoDetails struct {
	CompletedJobs []CompletedJob `json:"completed_jobs,omitempty"`
}

func (s *JobCompletionTriggerRunInfoDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Runtime state for a job completion trigger.
type JobCompletionTriggerState struct {
	// State for each configured upstream job, in the same order as `jobs`.
	MonitoredJobStates []JobCompletionTriggerStateMonitoredJobState `json:"monitored_job_states,omitempty"`
}

func (s *JobCompletionTriggerState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type JobCompletionTriggerStateMonitoredJobState struct {
	// Whether this job has contributed a matching completion toward the next
	// trigger fire.
	HasSeenCompletion bool `json:"has_seen_completion,omitempty"`
	// Whether the upstream job was missing or deleted when last observed by
	// this trigger. Omitted when deletion status is unknown. This observation
	// may become stale while the trigger is paused or evaluation cannot access
	// the upstream job.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// The ID of the upstream job. For a job monitored by name, the ID of the
	// job that the name last resolved to. Omitted until the name resolves to a
	// job.
	JobId string `json:"job_id,omitempty"`
	// The completion time of the latest matching upstream run observed by the
	// trigger.
	LastSeenCompletionTime *time.Time `json:"last_seen_completion_time,omitempty"`
	// The termination code of the latest matching upstream run observed by the
	// trigger.
	LastSeenResult TerminationCodeCode `json:"last_seen_result,omitempty"`
	// The ID of the latest matching upstream run observed by the trigger.
	LastSeenRunId string `json:"last_seen_run_id,omitempty"`
	// The termination code of the latest upstream run if it does not satisfy
	// `if_status`. Unset when the latest run matches or no run has been
	// observed.
	LatestNonMatchingResult TerminationCodeCode `json:"latest_non_matching_result,omitempty"`
	// The ID of the workspace containing the upstream job.
	WorkspaceId string `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobCompletionTriggerStateMonitoredJobState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobCompletionTriggerStateMonitoredJobState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobCompliance struct {
	// Whether this job is in compliance with the latest version of its policy.
	IsCompliant bool `json:"is_compliant,omitempty"`
	// Canonical unique identifier for a job.
	JobId int64 `json:"job_id"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	Violations map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobCompliance) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobCompliance) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobDeployment struct {
	// ID of the deployment that manages this job. Only set when `kind` is
	// `BUNDLE`. Used to look up deployment metadata from the Deployment
	// Metadata service.
	DeploymentId string `json:"deployment_id,omitempty"`
	// The kind of deployment that manages the job.
	//
	// * `BUNDLE`: The job is managed by Databricks Asset Bundle. *
	// `SYSTEM_MANAGED`: The job is managed by Databricks and is read-only.
	Kind JobDeploymentKind `json:"kind"`
	// Path of the file that contains deployment metadata.
	MetadataFilePath string `json:"metadata_file_path,omitempty"`
	// ID of the version of the deployment that produced this job. Only set when
	// `kind` is `BUNDLE`. Identifies a specific snapshot of the deployment in
	// the Deployment Metadata service.
	VersionId string `json:"version_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobDeployment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobDeployment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// * `BUNDLE`: The job is managed by Databricks Asset Bundle. *
// `SYSTEM_MANAGED`: The job is managed by Databricks and is read-only.
type JobDeploymentKind string

// The job is managed by Databricks Asset Bundle.
const JobDeploymentKindBundle JobDeploymentKind = `BUNDLE`

// The job is managed by <Databricks> and is read-only.
const JobDeploymentKindSystemManaged JobDeploymentKind = `SYSTEM_MANAGED`

// String representation for [fmt.Print]
func (f *JobDeploymentKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobDeploymentKind) Set(v string) error {
	switch v {
	case `BUNDLE`, `SYSTEM_MANAGED`:
		*f = JobDeploymentKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BUNDLE", "SYSTEM_MANAGED"`, v)
	}
}

// Values returns all possible values for JobDeploymentKind.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobDeploymentKind) Values() []JobDeploymentKind {
	return []JobDeploymentKind{
		JobDeploymentKindBundle,
		JobDeploymentKindSystemManaged,
	}
}

// Type always returns JobDeploymentKind to satisfy [pflag.Value] interface
func (f *JobDeploymentKind) Type() string {
	return "JobDeploymentKind"
}

// Edit mode of the job.
//
// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
// `EDITABLE`: The job is in an editable state and can be modified.
type JobEditMode string

// The job is in an editable state and can be modified.
const JobEditModeEditable JobEditMode = `EDITABLE`

// The job is in a locked UI state and cannot be modified.
const JobEditModeUiLocked JobEditMode = `UI_LOCKED`

// String representation for [fmt.Print]
func (f *JobEditMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobEditMode) Set(v string) error {
	switch v {
	case `EDITABLE`, `UI_LOCKED`:
		*f = JobEditMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EDITABLE", "UI_LOCKED"`, v)
	}
}

// Values returns all possible values for JobEditMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobEditMode) Values() []JobEditMode {
	return []JobEditMode{
		JobEditModeEditable,
		JobEditModeUiLocked,
	}
}

// Type always returns JobEditMode to satisfy [pflag.Value] interface
func (f *JobEditMode) Type() string {
	return "JobEditMode"
}

type JobEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded []string `json:"on_duration_warning_threshold_exceeded,omitempty"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure []string `json:"on_failure,omitempty"`
	// A list of email addresses to notify when platform-initiated maintenance
	// completes for a continuous job.
	OnMaintenanceComplete []string `json:"on_maintenance_complete,omitempty"`
	// A list of email addresses to notify when platform-initiated maintenance
	// starts for a continuous job.
	OnMaintenanceStart []string `json:"on_maintenance_start,omitempty"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []string `json:"on_start,omitempty"`
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	OnStreamingBacklogExceeded []string `json:"on_streaming_backlog_exceeded,omitempty"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess []string `json:"on_success,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobEmailNotifications) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobEmailNotifications) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobEnvironment struct {
	// The key of an environment. It has to be unique within a job.
	EnvironmentKey string `json:"environment_key"`
	// The environment specification, including the environment version and the
	// dependencies to install before a task runs.
	Spec *compute.Environment `json:"spec,omitempty"`
}

func (s *JobEnvironment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A named environment-variable entry, defined once at the job level and
// referenced by key from one or more tasks. Entries live on
// `JobSettings.environment_variables`, and tasks select one via
// `TaskSettings.environment_variables_key`.
type JobEnvironmentVariables struct {
	// Identifier for this entry. Must be unique within
	// `JobSettings.environment_variables`. Tasks reference it from
	// `TaskSettings.environment_variables_key`.
	EnvironmentVariablesKey string `json:"environment_variables_key"`
	// The environment variable specification.
	Spec *JobEnvironmentVariablesSpec `json:"spec,omitempty"`
}

func (s *JobEnvironmentVariables) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The environment variables and files associated with a job environment
// variable entry. Runtime environment variables override inline `variables`,
// which override values from `files`, on duplicate keys.
type JobEnvironmentVariablesSpec struct {
	// Workspace (`/Workspace/...`) or UC Volumes (`/Volumes/...`) paths to
	// `.env` files. Maximum 5 files. Files are read, parsed, and merged at task
	// execution time, not at job creation or update API call time.
	//
	// File format: each line containing a variable must be exactly `KEY=VALUE`.
	// Empty and whitespace-only lines, and lines beginning with `#`, are
	// ignored. Keys must match the same regex as inlined variable names
	// (`^[A-Za-z_][A-Za-z0-9_]*$`); the value continues to the end of the line.
	// No other syntax is supported — no inline comments, no quoted values, no
	// escape sequences, no variable interpolation. Any other line that does not
	// match the `KEY=VALUE` shape fails the run.
	//
	// Size limits: maximum 32,768 bytes (32 KiB) per file on disk; maximum
	// 1,024 bytes (1 KiB) per `KEY=VALUE` line combined. Files or lines
	// exceeding these limits fail the run.
	//
	// On a duplicate key, the later file wins; `variables` override values from
	// any file.
	Files []string `json:"files,omitempty"`
	// Environment variables specified directly as key/value pairs. Maximum 20
	// entries.
	//
	// Each key must be 1 to 256 characters and match
	// `^[A-Za-z_][A-Za-z0-9_]*$`: it must start with an ASCII letter or
	// underscore and contain only ASCII letters, digits, and underscores. Each
	// value can be any Unicode string of up to 512 characters, including an
	// empty string.
	Variables map[string]string `json:"variables,omitempty"`
}

func (s *JobEnvironmentVariablesSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type JobNotificationSettings struct {
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobNotificationSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobNotificationSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobParameter struct {
	// The optional default value of the parameter
	Default string `json:"default,omitempty"`
	// The name of the parameter
	Name string `json:"name,omitempty"`
	// The value used in the run
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobParameterDefinition struct {
	// Default value of the parameter.
	Default string `json:"default"`
	// The name of the defined parameter. May only contain alphanumeric
	// characters, `_`, `-`, and `.`
	Name string `json:"name"`
}

func (s *JobParameterDefinition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type JobPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type JobPermissionLevel string

const JobPermissionLevelCanManage JobPermissionLevel = `CAN_MANAGE`

const JobPermissionLevelCanManageRun JobPermissionLevel = `CAN_MANAGE_RUN`

const JobPermissionLevelCanView JobPermissionLevel = `CAN_VIEW`

const JobPermissionLevelIsOwner JobPermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (f *JobPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`:
		*f = JobPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_MANAGE_RUN", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Values returns all possible values for JobPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobPermissionLevel) Values() []JobPermissionLevel {
	return []JobPermissionLevel{
		JobPermissionLevelCanManage,
		JobPermissionLevelCanManageRun,
		JobPermissionLevelCanView,
		JobPermissionLevelIsOwner,
	}
}

// Type always returns JobPermissionLevel to satisfy [pflag.Value] interface
func (f *JobPermissionLevel) Type() string {
	return "JobPermissionLevel"
}

type JobPermissions struct {
	AccessControlList []JobAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobPermissionsRequest struct {
	AccessControlList []JobAccessControlRequest `json:"access_control_list,omitempty"`
	// The job for which to get or manage permissions.
	JobId string `json:"-" url:"-"`
}

func (s *JobPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Write-only setting. Specifies the user, service principal, or group that the
// job runs as. If not specified, the job runs as the user who created the job.
//
// One of `user_name`, `service_principal_name`, or `group_name` should be
// specified. If not, an error is thrown.
type JobRunAs struct {
	// Group name of an account group assigned to the workspace. When set, all
	// tasks run as the group and the group's permissions are used for data
	// access. Setting this field requires being a member of the group, or
	// having the `Assume` permission on the group.
	GroupName string `json:"group_name,omitempty"`
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The email of an active workspace user. Non-admin users can only set this
	// field to their own email.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobRunAs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobRunAs) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobSettings struct {
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	//
	// Pipelines started by a continuous job also run continuously, regardless
	// of their own pipeline mode setting.
	Continuous *Continuous `json:"continuous,omitempty"`
	// Deployment information for jobs managed by external sources.
	Deployment *JobDeployment `json:"deployment,omitempty"`
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	Description string `json:"description,omitempty"`
	// The window the run aims to finish within, in hours, for a Cost-optimized
	// serverless run. Required only when `performance_target` is
	// `COST_OPTIMIZED`. Must be one of the supported values: `NO_FLEXIBILITY`,
	// `4h`, `12h`, `24h`, `72h`. `NO_FLEXIBILITY` means the run uses the
	// default cost-optimized performance level without additional cost-saving
	// optimizations.
	DurationTarget string `json:"duration_target,omitempty"`
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode JobEditMode `json:"edit_mode,omitempty"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// Named environment-variable entries that tasks can reference by key from
	// `TaskSettings.environment_variables_key`. Each entry's `spec` holds
	// inline `variables` and optional `.env` `files`. Maximum 10 entries per
	// job. A task can reference at most one entry from this list.
	EnvironmentVariables []JobEnvironmentVariables `json:"environment_variables,omitempty"`
	// A list of task execution environment specifications that can be
	// referenced by tasks that use serverless compute or a compute resource
	// that uses Environments mode.
	//
	// For notebook tasks that use serverless compute or a compute resource that
	// uses Environments mode, if the `environment_key` is not specified, the
	// notebook environment will be used if present. If a jobs environment is
	// specified, it will override the notebook environment.
	//
	// For other tasks that use serverless compute or a compute resource that
	// uses Environments mode, the task environment is required to be specified
	// using `environment_key` in the task settings.
	Environments []JobEnvironment `json:"environments,omitempty"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format Format `json:"format,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *GitSource `json:"git_source,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `json:"health,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// An optional maximum allowed number of concurrent runs of the job. Set
	// this value if you want to be able to execute multiple runs of the same
	// job concurrently. This is useful for example if you trigger your job on a
	// frequent schedule and want to allow consecutive runs to overlap with each
	// other, or if you want to trigger multiple runs which differ by their
	// input parameters. This setting affects only new runs. For example,
	// suppose the job’s concurrency is 4 and there are 4 concurrent active
	// runs. Then setting the concurrency to 3 won’t kill any of the active
	// runs. However, from then on, new runs are skipped unless there are fewer
	// than 3 active runs. This value cannot exceed 1000. Setting this value to
	// `0` causes all new runs to be skipped.
	MaxConcurrentRuns int `json:"max_concurrent_runs,omitempty"`
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name string `json:"name,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *JobNotificationSettings `json:"notification_settings,omitempty"`
	// Job-level parameter definitions
	Parameters []JobParameterDefinition `json:"parameters,omitempty"`
	// Path of the job parent folder in workspace file tree. If absent, the job
	// doesn't have a parent folder.
	//
	// Example: /Workspace/Projects/my_project
	ParentPath string `json:"parent_path,omitempty"`
	// The performance mode on a serverless job. This field determines the level
	// of compute performance or cost-efficiency for the run. The performance
	// target does not apply to tasks that run on Serverless GPU compute.
	//
	// * `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance. * `STANDARD`:
	// Enables cost-efficient execution of serverless workloads. *
	// `COST_OPTIMIZED`: Enables lower job costs by optimizing compute for your
	// selected target duration time. Must provide a duration target.
	PerformanceTarget PerformanceTarget `json:"performance_target,omitempty"`
	// The queue settings of the job.
	Queue *QueueSettings `json:"queue,omitempty"`
	// The user or service principal that the job runs as, if specified in the
	// request. This field indicates the explicit configuration of `run_as` for
	// the job. To find the value in all cases, explicit or implicit, use
	// `run_as_user_name`.
	RunAs *JobRunAs `json:"run_as,omitempty"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags map[string]string `json:"tags,omitempty"`
	// A list of task specifications to be executed by this job. It supports up
	// to 1000 elements in write endpoints (:method:jobs/create,
	// :method:jobs/reset, :method:jobs/update, :method:jobs/submit). Read
	// endpoints return only 100 tasks. If more than 100 tasks are available,
	// you can paginate through them using :method:jobs/get. Use the
	// `next_page_token` field at the object root to determine if more results
	// are available.
	Tasks []Task `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger *TriggerSettings `json:"trigger,omitempty"`
	// List of triggers attached to this job. A run starts when any active
	// trigger evaluates to true. Cannot be set in the same request as the
	// legacy `schedule`, `trigger`, or `continuous` fields. Gated behind the
	// "Multiple Triggers" feature preview.
	Triggers []TriggerConfiguration `json:"triggers,omitempty"`
	// The id of the user specified usage policy to use for this job. If not
	// specified, a default usage policy may be applied when creating or
	// modifying the job. See `effective_usage_policy_id` for the usage policy
	// used by this workload.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *JobSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The source of the job specification in the remote repository when the job is
// source controlled.
type JobSource struct {
	// Dirty state indicates the job is not fully synced with the job
	// specification in the remote repository.
	//
	// Possible values are: * `NOT_SYNCED`: The job is not yet synced with the
	// remote job specification. Import the remote job specification from UI to
	// make the job fully synced. * `DISCONNECTED`: The job is temporary
	// disconnected from the remote job specification and is allowed for live
	// edit. Import the remote job specification again from UI to make the job
	// fully synced.
	DirtyState JobSourceDirtyState `json:"dirty_state,omitempty"`
	// Name of the branch which the job is imported from.
	ImportFromGitBranch string `json:"import_from_git_branch"`
	// Path of the job YAML file that contains the job specification.
	JobConfigPath string `json:"job_config_path"`
}

func (s *JobSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Dirty state indicates the job is not fully synced with the job specification
// in the remote repository.
//
// Possible values are: * `NOT_SYNCED`: The job is not yet synced with the
// remote job specification. Import the remote job specification from UI to make
// the job fully synced. * `DISCONNECTED`: The job is temporary disconnected
// from the remote job specification and is allowed for live edit. Import the
// remote job specification again from UI to make the job fully synced.
type JobSourceDirtyState string

// The job is temporary disconnected from the remote job specification and is
// allowed for live edit. Import the remote job specification again from UI to
// make the job fully synced.
const JobSourceDirtyStateDisconnected JobSourceDirtyState = `DISCONNECTED`

// The job is not yet synced with the remote job specification. Import the
// remote job specification from UI to make the job fully synced.
const JobSourceDirtyStateNotSynced JobSourceDirtyState = `NOT_SYNCED`

// String representation for [fmt.Print]
func (f *JobSourceDirtyState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobSourceDirtyState) Set(v string) error {
	switch v {
	case `DISCONNECTED`, `NOT_SYNCED`:
		*f = JobSourceDirtyState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISCONNECTED", "NOT_SYNCED"`, v)
	}
}

// Values returns all possible values for JobSourceDirtyState.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobSourceDirtyState) Values() []JobSourceDirtyState {
	return []JobSourceDirtyState{
		JobSourceDirtyStateDisconnected,
		JobSourceDirtyStateNotSynced,
	}
}

// Type always returns JobSourceDirtyState to satisfy [pflag.Value] interface
func (f *JobSourceDirtyState) Type() string {
	return "JobSourceDirtyState"
}

// Specifies the health metric that is being evaluated for a particular health
// rule.
//
// * `RUN_DURATION_SECONDS`: Expected total time for a run in seconds. *
// `STREAMING_BACKLOG_BYTES`: An estimate of the maximum bytes of data waiting
// to be consumed across all streams. This metric is in Public Preview. *
// `STREAMING_BACKLOG_RECORDS`: An estimate of the maximum offset lag across all
// streams. This metric is in Public Preview. * `STREAMING_BACKLOG_SECONDS`: An
// estimate of the maximum consumer delay across all streams. This metric is in
// Public Preview. * `STREAMING_BACKLOG_FILES`: An estimate of the maximum
// number of outstanding files across all streams. This metric is in Public
// Preview.
type JobsHealthMetric string

// Expected total time for a run in seconds.
const JobsHealthMetricRunDurationSeconds JobsHealthMetric = `RUN_DURATION_SECONDS`

// An estimate of the maximum bytes of data waiting to be consumed across all
// streams. This metric is in Public Preview.
const JobsHealthMetricStreamingBacklogBytes JobsHealthMetric = `STREAMING_BACKLOG_BYTES`

// An estimate of the maximum number of outstanding files across all streams.
// This metric is in Public Preview.
const JobsHealthMetricStreamingBacklogFiles JobsHealthMetric = `STREAMING_BACKLOG_FILES`

// An estimate of the maximum offset lag across all streams. This metric is in
// Public Preview.
const JobsHealthMetricStreamingBacklogRecords JobsHealthMetric = `STREAMING_BACKLOG_RECORDS`

// An estimate of the maximum consumer delay across all streams. This metric is
// in Public Preview.
const JobsHealthMetricStreamingBacklogSeconds JobsHealthMetric = `STREAMING_BACKLOG_SECONDS`

// String representation for [fmt.Print]
func (f *JobsHealthMetric) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobsHealthMetric) Set(v string) error {
	switch v {
	case `RUN_DURATION_SECONDS`, `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_FILES`, `STREAMING_BACKLOG_RECORDS`, `STREAMING_BACKLOG_SECONDS`:
		*f = JobsHealthMetric(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "RUN_DURATION_SECONDS", "STREAMING_BACKLOG_BYTES", "STREAMING_BACKLOG_FILES", "STREAMING_BACKLOG_RECORDS", "STREAMING_BACKLOG_SECONDS"`, v)
	}
}

// Values returns all possible values for JobsHealthMetric.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobsHealthMetric) Values() []JobsHealthMetric {
	return []JobsHealthMetric{
		JobsHealthMetricRunDurationSeconds,
		JobsHealthMetricStreamingBacklogBytes,
		JobsHealthMetricStreamingBacklogFiles,
		JobsHealthMetricStreamingBacklogRecords,
		JobsHealthMetricStreamingBacklogSeconds,
	}
}

// Type always returns JobsHealthMetric to satisfy [pflag.Value] interface
func (f *JobsHealthMetric) Type() string {
	return "JobsHealthMetric"
}

// Specifies the operator used to compare the health metric value with the
// specified threshold.
type JobsHealthOperator string

const JobsHealthOperatorGreaterThan JobsHealthOperator = `GREATER_THAN`

// String representation for [fmt.Print]
func (f *JobsHealthOperator) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobsHealthOperator) Set(v string) error {
	switch v {
	case `GREATER_THAN`:
		*f = JobsHealthOperator(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GREATER_THAN"`, v)
	}
}

// Values returns all possible values for JobsHealthOperator.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobsHealthOperator) Values() []JobsHealthOperator {
	return []JobsHealthOperator{
		JobsHealthOperatorGreaterThan,
	}
}

// Type always returns JobsHealthOperator to satisfy [pflag.Value] interface
func (f *JobsHealthOperator) Type() string {
	return "JobsHealthOperator"
}

type JobsHealthRule struct {
	// Specifies the health metric that is being evaluated for a particular
	// health rule.
	//
	// * `RUN_DURATION_SECONDS`: Expected total time for a run in seconds. *
	// `STREAMING_BACKLOG_BYTES`: An estimate of the maximum bytes of data
	// waiting to be consumed across all streams. This metric is in Public
	// Preview. * `STREAMING_BACKLOG_RECORDS`: An estimate of the maximum offset
	// lag across all streams. This metric is in Public Preview. *
	// `STREAMING_BACKLOG_SECONDS`: An estimate of the maximum consumer delay
	// across all streams. This metric is in Public Preview. *
	// `STREAMING_BACKLOG_FILES`: An estimate of the maximum number of
	// outstanding files across all streams. This metric is in Public Preview.
	Metric JobsHealthMetric `json:"metric"`
	// Specifies the operator used to compare the health metric value with the
	// specified threshold.
	Op JobsHealthOperator `json:"op"`
	// Specifies the threshold value that the health metric should obey to
	// satisfy the health rule.
	Value int64 `json:"value"`
}

func (s *JobsHealthRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An optional set of health rules that can be defined for this job.
type JobsHealthRules struct {
	// The health rules to evaluate for each run.
	Rules []JobsHealthRule `json:"rules,omitempty"`
}

func (s *JobsHealthRules) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListJobComplianceForPolicyResponse struct {
	// A list of jobs and their policy compliance statuses.
	Jobs []JobCompliance `json:"jobs,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If this field is not in the response, it means no further
	// results for the request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If this field is not in the response, it means no further
	// results for the request.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListJobComplianceForPolicyResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListJobComplianceForPolicyResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListJobComplianceRequest struct {
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token that can be used to navigate to the next page or previous
	// page as returned by `next_page_token` or `prev_page_token`.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Canonical unique identifier for the cluster policy.
	PolicyId string `json:"-" url:"policy_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListJobComplianceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListJobComplianceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListJobsRequest struct {
	// Whether to include task and cluster details in the response. Note that
	// only the first 100 elements will be shown. Use :method:jobs/get to
	// paginate through all tasks and clusters.
	ExpandTasks bool `json:"-" url:"expand_tasks,omitempty"`
	// The number of jobs to return. This value must be greater than 0 and less
	// or equal to 100. The default value is 20.
	Limit int `json:"-" url:"limit,omitempty"`
	// A filter on the list based on the exact (case insensitive) job name.
	Name string `json:"-" url:"name,omitempty"`
	// The offset of the first job to return, relative to the most recently
	// created job. Deprecated since June 2023. Use `page_token` to iterate
	// through the pages instead.
	Offset int `json:"-" url:"offset,omitempty"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of jobs respectively.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListJobsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListJobsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List of jobs was retrieved successfully.
type ListJobsResponse struct {
	// If true, additional jobs matching the provided filter are available for
	// listing.
	HasMore bool `json:"has_more,omitempty"`
	// The list of jobs. Only included in the response if there are jobs to
	// list.
	Jobs []BaseJob `json:"jobs,omitempty"`
	// A token that can be used to list the next page of jobs (if applicable).
	NextPageToken string `json:"next_page_token,omitempty"`
	// A token that can be used to list the previous page of jobs (if
	// applicable).
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListJobsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListJobsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRunsRequest struct {
	// If active_only is `true`, only active runs are included in the results;
	// otherwise, lists both active and completed runs. An active run is a run
	// in the `QUEUED`, `PENDING`, `RUNNING`, or `TERMINATING`. This field
	// cannot be `true` when completed_only is `true`.
	ActiveOnly bool `json:"-" url:"active_only,omitempty"`
	// If completed_only is `true`, only completed runs are included in the
	// results; otherwise, lists both active and completed runs. This field
	// cannot be `true` when active_only is `true`.
	CompletedOnly bool `json:"-" url:"completed_only,omitempty"`
	// Whether to include task and cluster details in the response. Note that
	// only the first 100 elements will be shown. Use :method:jobs/getrun to
	// paginate through all tasks and clusters.
	ExpandTasks bool `json:"-" url:"expand_tasks,omitempty"`
	// The job for which to list runs. If omitted, the Jobs service lists runs
	// from all jobs.
	JobId int64 `json:"-" url:"job_id,omitempty"`
	// The number of runs to return. This value must be greater than 0 and less
	// than 25. The default value is 20. If a request specifies a limit of 0,
	// the service instead uses the maximum limit.
	Limit int `json:"-" url:"limit,omitempty"`
	// The offset of the first run to return, relative to the most recent run.
	// Deprecated since June 2023. Use `page_token` to iterate through the pages
	// instead.
	Offset int `json:"-" url:"offset,omitempty"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of runs respectively.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The type of runs to return. For a description of run types, see
	// :method:jobs/getRun.
	RunType RunType `json:"-" url:"run_type,omitempty"`
	// Show runs that started _at or after_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_to_ to filter
	// by a time range.
	StartTimeFrom int64 `json:"-" url:"start_time_from,omitempty"`
	// Show runs that started _at or before_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_from_ to
	// filter by a time range.
	StartTimeTo int64 `json:"-" url:"start_time_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRunsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRunsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List of runs was retrieved successfully.
type ListRunsResponse struct {
	// If true, additional runs matching the provided filter are available for
	// listing.
	HasMore bool `json:"has_more,omitempty"`
	// A token that can be used to list the next page of runs (if applicable).
	NextPageToken string `json:"next_page_token,omitempty"`
	// A token that can be used to list the previous page of runs (if
	// applicable).
	PrevPageToken string `json:"prev_page_token,omitempty"`
	// A list of runs, from most recently started to least. Only included in the
	// response if there are runs to list.
	Runs []BaseRun `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A recurring weekly time window during which platform-initiated maintenance is
// allowed to run for a continuous job.
type MaintenanceWindow struct {
	// The day of week on which maintenance is allowed to happen. This field is
	// required.
	DayOfWeek DayOfWeek `json:"day_of_week"`
	// An integer between 0 and 23 denoting the start hour for the maintenance
	// window in the 24-hour day. Platform-initiated maintenance is triggered
	// only within a one-hour window starting at this hour. This field is
	// required.
	StartHour int `json:"start_hour"`
	// A Java timezone ID. The maintenance window is resolved with respect to
	// this timezone. See [Java TimeZone] for details. This field is required.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId string `json:"timezone_id"`
}

func (s *MaintenanceWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ModelTriggerConfiguration struct {
	// Aliases of the model versions to monitor. Can only be used in conjunction
	// with condition MODEL_ALIAS_SET.
	Aliases []string `json:"aliases,omitempty"`
	// The condition based on which to trigger a job run.
	Condition ModelTriggerConfigurationCondition `json:"condition"`
	// If set, the trigger starts a run only after the specified amount of time
	// has passed since the last time the trigger fired. The minimum allowed
	// value is 60 seconds.
	MinTimeBetweenTriggersSeconds int `json:"min_time_between_triggers_seconds,omitempty"`
	// Name of the securable to monitor ("mycatalog.myschema.mymodel" in the
	// case of model-level triggers, "mycatalog.myschema" in the case of
	// schema-level triggers) or empty in the case of metastore-level triggers.
	SecurableName string `json:"securable_name,omitempty"`
	// If set, the trigger starts a run only after no model updates have
	// occurred for the specified time and can be used to wait for a series of
	// model updates before triggering a run. The minimum allowed value is 60
	// seconds.
	WaitAfterLastChangeSeconds int `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ModelTriggerConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelTriggerConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ModelTriggerConfigurationCondition string

const ModelTriggerConfigurationConditionModelAliasSet ModelTriggerConfigurationCondition = `MODEL_ALIAS_SET`

const ModelTriggerConfigurationConditionModelCreated ModelTriggerConfigurationCondition = `MODEL_CREATED`

const ModelTriggerConfigurationConditionModelVersionReady ModelTriggerConfigurationCondition = `MODEL_VERSION_READY`

// String representation for [fmt.Print]
func (f *ModelTriggerConfigurationCondition) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelTriggerConfigurationCondition) Set(v string) error {
	switch v {
	case `MODEL_ALIAS_SET`, `MODEL_CREATED`, `MODEL_VERSION_READY`:
		*f = ModelTriggerConfigurationCondition(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MODEL_ALIAS_SET", "MODEL_CREATED", "MODEL_VERSION_READY"`, v)
	}
}

// Values returns all possible values for ModelTriggerConfigurationCondition.
//
// There is no guarantee on the order of the values in the slice.
func (f *ModelTriggerConfigurationCondition) Values() []ModelTriggerConfigurationCondition {
	return []ModelTriggerConfigurationCondition{
		ModelTriggerConfigurationConditionModelAliasSet,
		ModelTriggerConfigurationConditionModelCreated,
		ModelTriggerConfigurationConditionModelVersionReady,
	}
}

// Type always returns ModelTriggerConfigurationCondition to satisfy [pflag.Value] interface
func (f *ModelTriggerConfigurationCondition) Type() string {
	return "ModelTriggerConfigurationCondition"
}

// Runtime state for a model trigger. Currently empty because model triggers do
// not expose any trigger-specific runtime state.
type ModelTriggerState struct {
}

func (s *ModelTriggerState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type NotebookOutput struct {
	// The value passed to
	// [dbutils.notebook.exit()](/notebooks/notebook-workflows.html#notebook-workflows-exit).
	// Databricks restricts this API to return the first 5 MB of the value. For
	// a larger result, your job can store the results in a cloud storage
	// service. This field is absent if `dbutils.notebook.exit()` was never
	// called.
	Result string `json:"result,omitempty"`
	// Whether or not the result was truncated.
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NotebookOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotebookOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NotebookTask struct {
	// Base parameters to be used for each run of this job. If the run is
	// initiated by a call to :method:jobs/run Now with parameters specified,
	// the two parameters maps are merged. If the same key is specified in
	// `base_parameters` and in `run-now`, the value from `run-now` is used. Use
	// [Task parameter variables] to set parameters containing information about
	// job runs.
	//
	// If the notebook takes a parameter that is not specified in the job’s
	// `base_parameters` or the `run-now` override parameters, the default value
	// from the notebook is used.
	//
	// Retrieve these parameters in a notebook using [dbutils.widgets.get].
	//
	// The JSON representation of this field cannot exceed 1MB.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-widgets
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
	// The path of the notebook to be run in the Databricks workspace or remote
	// repository. For notebooks stored in the Databricks workspace, the path
	// must be absolute and begin with a slash. For notebooks stored in a remote
	// repository, the path must be relative. This field is required.
	NotebookPath string `json:"notebook_path"`
	// Optional location type of the notebook. When set to `WORKSPACE`, the
	// notebook will be retrieved from the local Databricks workspace. When set
	// to `GIT`, the notebook will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`:
	// Notebook is located in Databricks workspace. * `GIT`: Notebook is located
	// in cloud Git provider.
	Source Source `json:"source,omitempty"`
	// Optional `warehouse_id` to run the notebook on a SQL warehouse. Classic
	// SQL warehouses are NOT supported, please use serverless or pro SQL
	// warehouses.
	//
	// Note that SQL warehouses only support SQL cells; if the notebook contains
	// non-SQL cells, the run will fail.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NotebookTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotebookTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Stores the catalog name, schema name, and the output schema expiration time
// for the clean room run.
type OutputSchemaInfo struct {
	CatalogName string `json:"catalog_name,omitempty"`
	// The expiration time for the output schema as a Unix timestamp in
	// milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`

	SchemaName string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OutputSchemaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OutputSchemaInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PauseStatus string

const PauseStatusPaused PauseStatus = `PAUSED`

const PauseStatusUnpaused PauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *PauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = PauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Values returns all possible values for PauseStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *PauseStatus) Values() []PauseStatus {
	return []PauseStatus{
		PauseStatusPaused,
		PauseStatusUnpaused,
	}
}

// Type always returns PauseStatus to satisfy [pflag.Value] interface
func (f *PauseStatus) Type() string {
	return "PauseStatus"
}

// Per-trigger runtime state for the multi-trigger surface. Mirrors
// `TriggerConfiguration`'s trigger-type variants 1:1; each entry sets exactly
// one variant matching the corresponding trigger's type. Variants with no
// runtime state today (`schedule`, `model`) are emitted as empty messages.
type PerTriggerState struct {
	// State of the continuous trigger.
	Continuous *ContinuousTriggerState `json:"continuous,omitempty"`
	// State of the file arrival trigger.
	FileArrival *FileArrivalTriggerState `json:"file_arrival,omitempty"`
	// State of the job completion trigger.
	JobCompletion *JobCompletionTriggerState `json:"job_completion,omitempty"`
	// State of the model trigger.
	Model *ModelTriggerState `json:"model,omitempty"`
	// Whether this trigger is paused or not. Mirrors the configured
	// pause_status.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
	// State of the periodic trigger.
	Periodic *PeriodicTriggerState `json:"periodic,omitempty"`
	// State of the cron schedule trigger.
	Schedule *ScheduleTriggerState `json:"schedule,omitempty"`
	// State for SQL condition evaluation, can coexist with other trigger
	// states.
	SqlCondition *SqlConditionState `json:"sql_condition,omitempty"`
	// State of the table update trigger.
	TableUpdate *TableTriggerState `json:"table_update,omitempty"`
}

func (s *PerTriggerState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// PerformanceTarget defines how performant (lower latency) or cost efficient
// the execution of run on serverless compute should be. The performance mode on
// the job or pipeline should map to a performance setting that is passed to
// Cluster Manager (see cluster-common PerformanceTarget).
type PerformanceTarget string

const PerformanceTargetPerformanceOptimized PerformanceTarget = `PERFORMANCE_OPTIMIZED`

const PerformanceTargetStandard PerformanceTarget = `STANDARD`

// String representation for [fmt.Print]
func (f *PerformanceTarget) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PerformanceTarget) Set(v string) error {
	switch v {
	case `PERFORMANCE_OPTIMIZED`, `STANDARD`:
		*f = PerformanceTarget(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PERFORMANCE_OPTIMIZED", "STANDARD"`, v)
	}
}

// Values returns all possible values for PerformanceTarget.
//
// There is no guarantee on the order of the values in the slice.
func (f *PerformanceTarget) Values() []PerformanceTarget {
	return []PerformanceTarget{
		PerformanceTargetPerformanceOptimized,
		PerformanceTargetStandard,
	}
}

// Type always returns PerformanceTarget to satisfy [pflag.Value] interface
func (f *PerformanceTarget) Type() string {
	return "PerformanceTarget"
}

type PeriodicTriggerConfiguration struct {
	// The interval at which the trigger should run.
	Interval int `json:"interval"`
	// The unit of time for the interval.
	Unit PeriodicTriggerConfigurationTimeUnit `json:"unit"`
}

func (s *PeriodicTriggerConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PeriodicTriggerConfigurationTimeUnit string

const PeriodicTriggerConfigurationTimeUnitDays PeriodicTriggerConfigurationTimeUnit = `DAYS`

const PeriodicTriggerConfigurationTimeUnitHours PeriodicTriggerConfigurationTimeUnit = `HOURS`

const PeriodicTriggerConfigurationTimeUnitMinutes PeriodicTriggerConfigurationTimeUnit = `MINUTES`

const PeriodicTriggerConfigurationTimeUnitWeeks PeriodicTriggerConfigurationTimeUnit = `WEEKS`

// String representation for [fmt.Print]
func (f *PeriodicTriggerConfigurationTimeUnit) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PeriodicTriggerConfigurationTimeUnit) Set(v string) error {
	switch v {
	case `DAYS`, `HOURS`, `MINUTES`, `WEEKS`:
		*f = PeriodicTriggerConfigurationTimeUnit(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DAYS", "HOURS", "MINUTES", "WEEKS"`, v)
	}
}

// Values returns all possible values for PeriodicTriggerConfigurationTimeUnit.
//
// There is no guarantee on the order of the values in the slice.
func (f *PeriodicTriggerConfigurationTimeUnit) Values() []PeriodicTriggerConfigurationTimeUnit {
	return []PeriodicTriggerConfigurationTimeUnit{
		PeriodicTriggerConfigurationTimeUnitDays,
		PeriodicTriggerConfigurationTimeUnitHours,
		PeriodicTriggerConfigurationTimeUnitMinutes,
		PeriodicTriggerConfigurationTimeUnitWeeks,
	}
}

// Type always returns PeriodicTriggerConfigurationTimeUnit to satisfy [pflag.Value] interface
func (f *PeriodicTriggerConfigurationTimeUnit) Type() string {
	return "PeriodicTriggerConfigurationTimeUnit"
}

type PeriodicTriggerState struct {
	// The time of the next scheduled run, in epoch milliseconds. Only set while
	// the trigger is not paused.
	NextRunTime int64 `json:"next_run_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PeriodicTriggerState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PeriodicTriggerState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineParams struct {
	// If true, triggers a full refresh on the spark declarative pipeline.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// A list of tables to update with fullRefresh.
	FullRefreshSelection []string `json:"full_refresh_selection,omitempty"`
	// Flow names to selectively refresh. These are unioned with other selective
	// refresh options (refresh_selection, full_refresh_selection) to determine
	// the final set of flows to refresh.
	RefreshFlowSelection []string `json:"refresh_flow_selection,omitempty"`
	// A list of tables to update without fullRefresh.
	RefreshSelection []string `json:"refresh_selection,omitempty"`
	// A list of streaming flows to reset checkpoints without clearing data.
	ResetCheckpointSelection []string `json:"reset_checkpoint_selection,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineParams) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineParams) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineTask struct {
	// If true, triggers a full refresh on the spark declarative pipeline.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// A list of tables to update with fullRefresh.
	FullRefreshSelection []string `json:"full_refresh_selection,omitempty"`
	// Key/value-map of parameters passed to the pipeline execution. Limited to
	// 10k characters in total.
	Parameters map[string]string `json:"parameters,omitempty"`
	// The full name of the pipeline task to execute.
	PipelineId string `json:"pipeline_id"`
	// Flow names to selectively refresh. These are unioned with other selective
	// refresh options (refresh_selection, full_refresh_selection) to determine
	// the final set of flows to refresh.
	RefreshFlowSelection []string `json:"refresh_flow_selection,omitempty"`
	// A list of tables to update without fullRefresh.
	RefreshSelection []string `json:"refresh_selection,omitempty"`
	// A list of streaming flows to reset checkpoints without clearing data.
	ResetCheckpointSelection []string `json:"reset_checkpoint_selection,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PowerBiModel struct {
	// How the published Power BI model authenticates to Databricks
	AuthenticationMethod AuthenticationMethod `json:"authentication_method,omitempty"`
	// The name of the Power BI model
	ModelName string `json:"model_name,omitempty"`
	// Whether to overwrite existing Power BI models
	OverwriteExisting bool `json:"overwrite_existing,omitempty"`
	// The default storage mode of the Power BI model
	StorageMode StorageMode `json:"storage_mode,omitempty"`
	// The name of the Power BI workspace of the model
	WorkspaceName string `json:"workspace_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PowerBiModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PowerBiModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PowerBiTable struct {
	// The catalog name in Databricks
	Catalog string `json:"catalog,omitempty"`
	// The table name in Databricks
	Name string `json:"name,omitempty"`
	// The schema name in Databricks
	Schema string `json:"schema,omitempty"`
	// The Power BI storage mode of the table
	StorageMode StorageMode `json:"storage_mode,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PowerBiTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PowerBiTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PowerBiTask struct {
	// The resource name of the UC connection to authenticate from Databricks to
	// Power BI
	ConnectionResourceName string `json:"connection_resource_name,omitempty"`
	// The semantic model to update
	PowerBiModel *PowerBiModel `json:"power_bi_model,omitempty"`
	// Whether the model should be refreshed after the update
	RefreshAfterUpdate bool `json:"refresh_after_update,omitempty"`
	// The tables to be exported to Power BI
	Tables []PowerBiTable `json:"tables,omitempty"`
	// The SQL warehouse ID to use as the Power BI data source
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PowerBiTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PowerBiTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PythonOperatorTask struct {
	// Fully qualified name of the main class or function. For example,
	// `my_project.my_function` or `my_project.MyOperator`.
	Main string `json:"main,omitempty"`
	// An ordered list of task parameters.
	Parameters []PythonOperatorTaskParameter `json:"parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PythonOperatorTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PythonOperatorTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PythonOperatorTaskParameter struct {
	Name string `json:"name,omitempty"`

	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PythonOperatorTaskParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PythonOperatorTaskParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PythonWheelTask struct {
	// Named entry point to use, if it does not exist in the metadata of the
	// package it executes the function from the package directly using
	// `$packageName.$entryPoint()`
	EntryPoint string `json:"entry_point"`
	// Command-line parameters passed to Python wheel task in the form of
	// `["--name=task", "--data=dbfs:/path/to/data.json"]`. Leave it empty if
	// `parameters` is not null.
	NamedParameters map[string]string `json:"named_parameters,omitempty"`
	// Name of the package to execute
	PackageName string `json:"package_name"`
	// Command-line parameters passed to Python wheel task. Leave it empty if
	// `named_parameters` is not null.
	Parameters []string `json:"parameters,omitempty"`
}

func (s *PythonWheelTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A three-level Unity Catalog name (catalog.schema.table); any level may be
// unset.
type QualifiedName struct {
	// Unity Catalog catalog.
	Catalog string `json:"catalog,omitempty"`
	// Unity Catalog schema, within the catalog.
	Schema string `json:"schema,omitempty"`
	// Table, within the catalog and schema.
	Table string `json:"table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QualifiedName) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QualifiedName) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueueDetails struct {
	Code QueueDetailsCodeCode `json:"code,omitempty"`
	// A descriptive message with the queuing details. This field is
	// unstructured, and its exact format is subject to change.
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueueDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueueDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The reason for queuing the run. * `ACTIVE_RUNS_LIMIT_REACHED`: The run was
// queued due to reaching the workspace limit of active task runs. *
// `MAX_CONCURRENT_RUNS_REACHED`: The run was queued due to reaching the per-job
// limit of concurrent job runs. * `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`: The run
// was queued due to reaching the workspace limit of active run job tasks.
type QueueDetailsCodeCode string

// The run was queued due to reaching the workspace limit of active task runs.
const QueueDetailsCodeCodeActiveRunsLimitReached QueueDetailsCodeCode = `ACTIVE_RUNS_LIMIT_REACHED`

// The run was queued due to reaching the workspace limit of active run job
// tasks.
const QueueDetailsCodeCodeActiveRunJobTasksLimitReached QueueDetailsCodeCode = `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`

// The run was queued due to reaching the per-job limit of concurrent job runs.
const QueueDetailsCodeCodeMaxConcurrentRunsReached QueueDetailsCodeCode = `MAX_CONCURRENT_RUNS_REACHED`

// String representation for [fmt.Print]
func (f *QueueDetailsCodeCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *QueueDetailsCodeCode) Set(v string) error {
	switch v {
	case `ACTIVE_RUNS_LIMIT_REACHED`, `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`, `MAX_CONCURRENT_RUNS_REACHED`:
		*f = QueueDetailsCodeCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE_RUNS_LIMIT_REACHED", "ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED", "MAX_CONCURRENT_RUNS_REACHED"`, v)
	}
}

// Values returns all possible values for QueueDetailsCodeCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *QueueDetailsCodeCode) Values() []QueueDetailsCodeCode {
	return []QueueDetailsCodeCode{
		QueueDetailsCodeCodeActiveRunsLimitReached,
		QueueDetailsCodeCodeActiveRunJobTasksLimitReached,
		QueueDetailsCodeCodeMaxConcurrentRunsReached,
	}
}

// Type always returns QueueDetailsCodeCode to satisfy [pflag.Value] interface
func (f *QueueDetailsCodeCode) Type() string {
	return "QueueDetailsCodeCode"
}

type QueueSettings struct {
	// If true, enable queueing for the job. This is a required field.
	Enabled bool `json:"enabled"`
}

func (s *QueueSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RepairHistoryItem struct {
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance. * `STANDARD`:
	// Enables cost-efficient execution of serverless workloads. *
	// `COST_OPTIMIZED`: Enables lower job costs by optimizing compute for your
	// selected target duration time. Must provide a duration target.
	EffectivePerformanceTarget PerformanceTarget `json:"effective_performance_target,omitempty"`
	// The end time of the (repaired) run.
	EndTime int64 `json:"end_time,omitempty"`
	// The ID of the repair. Only returned for the items that represent a repair
	// in `repair_history`.
	Id int64 `json:"id,omitempty"`
	// The start time of the (repaired) run.
	StartTime int64 `json:"start_time,omitempty"`
	// Deprecated. Please use the `status` field instead.
	State *RunState `json:"state,omitempty"`

	Status *RunStatus `json:"status,omitempty"`
	// The run IDs of the task runs that ran as part of this repair history
	// item.
	TaskRunIds []int64 `json:"task_run_ids,omitempty"`
	// The repair history item type. Indicates whether a run is the original run
	// or a repair run.
	Type RepairHistoryItemType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RepairHistoryItem) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepairHistoryItem) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The repair history item type. Indicates whether a run is the original run or
// a repair run.
type RepairHistoryItemType string

const RepairHistoryItemTypeOriginal RepairHistoryItemType = `ORIGINAL`

const RepairHistoryItemTypeRepair RepairHistoryItemType = `REPAIR`

// String representation for [fmt.Print]
func (f *RepairHistoryItemType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RepairHistoryItemType) Set(v string) error {
	switch v {
	case `ORIGINAL`, `REPAIR`:
		*f = RepairHistoryItemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ORIGINAL", "REPAIR"`, v)
	}
}

// Values returns all possible values for RepairHistoryItemType.
//
// There is no guarantee on the order of the values in the slice.
func (f *RepairHistoryItemType) Values() []RepairHistoryItemType {
	return []RepairHistoryItemType{
		RepairHistoryItemTypeOriginal,
		RepairHistoryItemTypeRepair,
	}
}

// Type always returns RepairHistoryItemType to satisfy [pflag.Value] interface
func (f *RepairHistoryItemType) Type() string {
	return "RepairHistoryItemType"
}

type RepairRun struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// The window the run aims to finish within, in hours, for a Cost-optimized
	// serverless run. Required only when `performance_target` is
	// `COST_OPTIMIZED`. Must be one of the supported values: `NO_FLEXIBILITY`,
	// `4h`, `12h`, `24h`, `72h`. `NO_FLEXIBILITY` means the run uses the
	// default cost-optimized performance level without additional cost-saving
	// optimizations.
	DurationTarget string `json:"duration_target,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	JarParams []string `json:"jar_params,omitempty"`
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters map[string]string `json:"job_parameters,omitempty"`
	// The ID of the latest repair. This parameter is not required when
	// repairing a run for the first time, but must be provided on subsequent
	// requests to repair the same run.
	LatestRepairId int64 `json:"latest_repair_id,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	NotebookParams map[string]string `json:"notebook_params,omitempty"`
	// The performance mode on a serverless job. The performance target
	// determines the level of compute performance or cost-efficiency for the
	// run. This field overrides the performance target defined on the job
	// level.
	//
	// * `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance. * `STANDARD`:
	// Enables cost-efficient execution of serverless workloads. *
	// `COST_OPTIMIZED`: Enables lower job costs by optimizing compute for your
	// selected target duration time. Must provide a duration target.
	PerformanceTarget PerformanceTarget `json:"performance_target,omitempty"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *PipelineParams `json:"pipeline_params,omitempty"`

	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	PythonParams []string `json:"python_params,omitempty"`
	// If true, repair all failed tasks. Only one of `rerun_tasks` or
	// `rerun_all_failed_tasks` can be used.
	RerunAllFailedTasks bool `json:"rerun_all_failed_tasks,omitempty"`
	// If true, repair all tasks that depend on the tasks in `rerun_tasks`, even
	// if they were previously successful. Can be also used in combination with
	// `rerun_all_failed_tasks`.
	RerunDependentTasks bool `json:"rerun_dependent_tasks,omitempty"`
	// The task keys of the task runs to repair.
	RerunTasks []string `json:"rerun_tasks,omitempty"`
	// The job run ID of the run to repair. The run must not be in progress.
	RunId int64 `json:"run_id"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	SqlParams map[string]string `json:"sql_params,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RepairRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepairRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Run repair was initiated.
type RepairRunResponse struct {
	// The ID of the repair. Must be provided in subsequent repairs using the
	// `latest_repair_id` field to ensure sequential repairs.
	RepairId int64 `json:"repair_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RepairRunResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepairRunResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ResetJob struct {
	// The canonical identifier of the job to reset. This field is required.
	JobId int64 `json:"job_id"`
	// The new settings of the job. These settings completely replace the old
	// settings.
	//
	// Changes to the field `JobBaseSettings.timeout_seconds` are applied to
	// active runs. Changes to other fields are applied to future runs only.
	NewSettings JobSettings `json:"new_settings"`
}

func (s *ResetJob) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolvedConditionTaskValues struct {
	Left string `json:"left,omitempty"`

	Right string `json:"right,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResolvedConditionTaskValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResolvedConditionTaskValues) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ResolvedDbtTaskValues struct {
	Commands []string `json:"commands,omitempty"`
}

func (s *ResolvedDbtTaskValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolvedNotebookTaskValues struct {
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
}

func (s *ResolvedNotebookTaskValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolvedParamPairValues struct {
	Parameters map[string]string `json:"parameters,omitempty"`
}

func (s *ResolvedParamPairValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolvedPipelineTaskValues struct {
	// Key/value-map of parameters passed to the pipeline execution. Limited to
	// 10k characters in total.
	Parameters map[string]string `json:"parameters,omitempty"`
}

func (s *ResolvedPipelineTaskValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolvedPythonWheelTaskValues struct {
	NamedParameters map[string]string `json:"named_parameters,omitempty"`

	Parameters []string `json:"parameters,omitempty"`
}

func (s *ResolvedPythonWheelTaskValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolvedRunJobTaskValues struct {
	JobParameters map[string]string `json:"job_parameters,omitempty"`

	Parameters map[string]string `json:"parameters,omitempty"`
}

func (s *ResolvedRunJobTaskValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolvedStringParamsValues struct {
	Parameters []string `json:"parameters,omitempty"`
}

func (s *ResolvedStringParamsValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolvedValues struct {
	// Resolved values for an AI Function task: the Delta table input with
	// upstream-task references replaced by concrete values before the run.
	AiFunctionTask *ResolvedValuesAiFunctionTaskResolvedValues `json:"ai_function_task,omitempty"`
	// Resolved values for an AI Runtime task — env_vars with
	// `{{tasks.<key>.values.<name>}}` references substituted to concrete values
	// before submission to the training service.
	AiRuntimeTask *ResolvedValuesAiRuntimeTaskResolvedValues `json:"ai_runtime_task,omitempty"`

	ConditionTask *ResolvedConditionTaskValues `json:"condition_task,omitempty"`

	DbtTask *ResolvedDbtTaskValues `json:"dbt_task,omitempty"`

	NotebookTask *ResolvedNotebookTaskValues `json:"notebook_task,omitempty"`

	PipelineTask *ResolvedPipelineTaskValues `json:"pipeline_task,omitempty"`

	PythonWheelTask *ResolvedPythonWheelTaskValues `json:"python_wheel_task,omitempty"`

	RunJobTask *ResolvedRunJobTaskValues `json:"run_job_task,omitempty"`

	SimulationTask *ResolvedParamPairValues `json:"simulation_task,omitempty"`

	SparkJarTask *ResolvedStringParamsValues `json:"spark_jar_task,omitempty"`

	SparkPythonTask *ResolvedStringParamsValues `json:"spark_python_task,omitempty"`

	SparkSubmitTask *ResolvedStringParamsValues `json:"spark_submit_task,omitempty"`

	SqlTask *ResolvedParamPairValues `json:"sql_task,omitempty"`
}

func (s *ResolvedValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Resolved values for an AI Function task's Delta table input.
type ResolvedValuesAiFunctionTaskResolvedValues struct {
	// The Delta table source with any upstream-task references (such as
	// `{{tasks.<task_key>.values.output_table_name}}`) in the table name
	// replaced by concrete values before the run.
	DeltaTable *DeltaTableSource `json:"delta_table,omitempty"`
}

func (s *ResolvedValuesAiFunctionTaskResolvedValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Resolved values for an AiRuntimeTask after dynamic-value substitution, so
// Jobs can expand `{{tasks.<key>.values.<name>}}` references before submission.
type ResolvedValuesAiRuntimeTaskResolvedValues struct {
}

func (s *ResolvedValuesAiRuntimeTaskResolvedValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Run was retrieved successfully
type Run struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `json:"cluster_instance,omitempty"`
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec *ClusterSpec `json:"cluster_spec,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// ID of the deployment that produced the job when this run was created.
	// Used to look up deployment metadata from the Deployment Metadata service.
	// Only set for job runs of jobs with a `BUNDLE` deployment.
	DeploymentId string `json:"deployment_id,omitempty"`
	// Description of the run
	Description string `json:"description,omitempty"`
	// The window the run aims to finish within, in hours, for a Cost-optimized
	// serverless run. Required only when `performance_target` is
	// `COST_OPTIMIZED`. Must be one of the supported values: `NO_FLEXIBILITY`,
	// `4h`, `12h`, `24h`, `72h`. `NO_FLEXIBILITY` means the run uses the
	// default cost-optimized performance level without additional cost-saving
	// optimizations.
	DurationTarget string `json:"duration_target,omitempty"`
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance. * `STANDARD`:
	// Enables cost-efficient execution of serverless workloads. *
	// `COST_OPTIMIZED`: Enables lower job costs by optimizing compute for your
	// selected target duration time. Must provide a duration target.
	EffectivePerformanceTarget PerformanceTarget `json:"effective_performance_target,omitempty"`
	// The id of the usage policy used by this run for cost attribution
	// purposes.
	EffectiveUsagePolicyId string `json:"effective_usage_policy_id,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// Snapshot of `JobSettings.environment_variables` as it was at run launch
	// — the full list of named environment-variable entries the job defined.
	// To find which entry a given task ran with, look at
	// `RunTaskSettings.environment_variables_key`.
	EnvironmentVariables []JobEnvironmentVariables `json:"environment_variables,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *GitSource `json:"git_source,omitempty"`
	// Indicates if the run has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/getrun
	// endpoint. It is only relevant for API 2.2 :method:jobs/listruns requests
	// with `expand_tasks=true`.
	HasMore bool `json:"has_more,omitempty"`
	// Only populated by for-each iterations. The parent for-each task is
	// located in tasks array.
	Iterations []RunTask `json:"iterations,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings. If more than 100
	// job clusters are available, you can paginate through them using
	// :method:jobs/getrun.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// The canonical identifier of the job that contains this run.
	JobId int64 `json:"job_id,omitempty"`
	// Job-level parameters used in the run
	JobParameters []JobParameter `json:"job_parameters,omitempty"`
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	JobRunId int64 `json:"job_run_id,omitempty"`
	// A token that can be used to list the next page of array properties.
	NextPageToken string `json:"next_page_token,omitempty"`
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId int64 `json:"original_attempt_run_id,omitempty"`
	// The parameters used for this run.
	OverridingParameters *RunParameters `json:"overriding_parameters,omitempty"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration int64 `json:"queue_duration,omitempty"`
	// The repair history of the run.
	RepairHistory []RepairHistoryItem `json:"repair_history,omitempty"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int64 `json:"run_duration,omitempty"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId int64 `json:"run_id,omitempty"`
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName string `json:"run_name,omitempty"`
	// The URL to the detail page of the run.
	RunPageUrl string `json:"run_page_url,omitempty"`

	RunType RunType `json:"run_type,omitempty"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64 `json:"setup_duration,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// Deprecated. Please use the `status` field instead.
	State *RunState `json:"state,omitempty"`

	Status *RunStatus `json:"status,omitempty"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run results. If
	// more than 100 tasks are available, you can paginate through them using
	// :method:jobs/getrun. Use the `next_page_token` field at the object root
	// to determine if more results are available.
	Tasks []RunTask `json:"tasks,omitempty"`

	Trigger TriggerType `json:"trigger,omitempty"`

	TriggerInfo *TriggerInfo `json:"trigger_info,omitempty"`
	// ID of the deployment version that produced the job when this run was
	// created. Identifies a specific snapshot of the deployment in the
	// Deployment Metadata service. Only set for job runs of jobs with a
	// `BUNDLE` deployment.
	VersionId string `json:"version_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Run) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Run) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	Left string `json:"left"`
	// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
	// operands. This means that `“12.0” == “12”` will evaluate to
	// `false`. * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`,
	// `LESS_THAN_OR_EQUAL` operators perform numeric comparison of their
	// operands. `“12.0” >= “12”` will evaluate to `true`, `“10.0”
	// >= “12”` will evaluate to `false`.
	//
	// The boolean comparison to task values can be implemented with operators
	// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it
	// will be serialized to `“true”` or `“false”` for the comparison.
	Op ConditionTaskOp `json:"op"`
	// The condition expression evaluation result. Filled in if the task was
	// successfully completed. Can be `"true"` or `"false"`
	Outcome string `json:"outcome,omitempty"`
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right string `json:"right"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunConditionTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunConditionTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunForEachTask struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	Concurrency int `json:"concurrency,omitempty"`
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs string `json:"inputs"`
	// Read only field. Populated for GetRun and ListRuns RPC calls and stores
	// the execution stats of a `For each` task.
	Stats *ForEachStats `json:"stats,omitempty"`
	// Configuration for the task that will be run for each element in the array
	Task Task `json:"task"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunForEachTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunForEachTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An optional value indicating the condition that determines whether the task
// should be run once its dependencies have been completed. When omitted,
// defaults to `ALL_SUCCESS`.
//
// Possible values are: * `ALL_SUCCESS`: All dependencies have executed and
// succeeded * `AT_LEAST_ONE_SUCCESS`: At least one dependency has succeeded *
// `NONE_FAILED`: None of the dependencies have failed and at least one was
// executed * `ALL_DONE`: All dependencies have been completed *
// `AT_LEAST_ONE_FAILED`: At least one dependency failed * `ALL_FAILED`: ALl
// dependencies have failed
type RunIf string

// All dependencies have been completed
const RunIfAllDone RunIf = `ALL_DONE`

// ALl dependencies have failed
const RunIfAllFailed RunIf = `ALL_FAILED`

// All dependencies have executed and succeeded
const RunIfAllSuccess RunIf = `ALL_SUCCESS`

// At least one dependency failed
const RunIfAtLeastOneFailed RunIf = `AT_LEAST_ONE_FAILED`

// At least one dependency has succeeded
const RunIfAtLeastOneSuccess RunIf = `AT_LEAST_ONE_SUCCESS`

// None of the dependencies have failed and at least one was executed
const RunIfNoneFailed RunIf = `NONE_FAILED`

// String representation for [fmt.Print]
func (f *RunIf) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunIf) Set(v string) error {
	switch v {
	case `ALL_DONE`, `ALL_FAILED`, `ALL_SUCCESS`, `AT_LEAST_ONE_FAILED`, `AT_LEAST_ONE_SUCCESS`, `NONE_FAILED`:
		*f = RunIf(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_DONE", "ALL_FAILED", "ALL_SUCCESS", "AT_LEAST_ONE_FAILED", "AT_LEAST_ONE_SUCCESS", "NONE_FAILED"`, v)
	}
}

// Values returns all possible values for RunIf.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunIf) Values() []RunIf {
	return []RunIf{
		RunIfAllDone,
		RunIfAllFailed,
		RunIfAllSuccess,
		RunIfAtLeastOneFailed,
		RunIfAtLeastOneSuccess,
		RunIfNoneFailed,
	}
}

// Type always returns RunIf to satisfy [pflag.Value] interface
func (f *RunIf) Type() string {
	return "RunIf"
}

type RunJobOutput struct {
	// The run id of the triggered job run
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunJobOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunJobOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunJobTask struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	JarParams []string `json:"jar_params,omitempty"`
	// ID of the job to trigger.
	JobId int64 `json:"job_id"`
	// Job-level parameters used to trigger the job.
	JobParameters map[string]string `json:"job_parameters,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	NotebookParams map[string]string `json:"notebook_params,omitempty"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *PipelineParams `json:"pipeline_params,omitempty"`

	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	PythonParams []string `json:"python_params,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	SqlParams map[string]string `json:"sql_params,omitempty"`
}

func (s *RunJobTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A value indicating the run's lifecycle state. The possible values are: *
// `QUEUED`: The run is queued. * `PENDING`: The run is waiting to be executed
// while the cluster and execution context are being prepared. * `RUNNING`: The
// task of this run is being executed. * `TERMINATING`: The task of this run has
// completed, and the cluster and execution context are being cleaned up. *
// `TERMINATED`: The task of this run has completed, and the cluster and
// execution context have been cleaned up. This state is terminal. * `SKIPPED`:
// This run was aborted because a previous run of the same job was already
// active. This state is terminal. * `INTERNAL_ERROR`: An exceptional state that
// indicates a failure in the Jobs service, such as network failure over a long
// period. If a run on a new cluster ends in the `INTERNAL_ERROR` state, the
// Jobs service terminates the cluster as soon as possible. This state is
// terminal. * `BLOCKED`: The run is blocked on an upstream dependency. *
// `WAITING_FOR_RETRY`: The run is waiting for a retry.
type RunLifeCycleState string

// The run is blocked on an upstream dependency.
const RunLifeCycleStateBlocked RunLifeCycleState = `BLOCKED`

// An exceptional state that indicates a failure in the Jobs service, such as
// network failure over a long period. If a run on a new cluster ends in the
// `INTERNAL_ERROR` state, the Jobs service terminates the cluster as soon as
// possible. This state is terminal.
const RunLifeCycleStateInternalError RunLifeCycleState = `INTERNAL_ERROR`

// The run is waiting to be executed while the cluster and execution context are
// being prepared.
const RunLifeCycleStatePending RunLifeCycleState = `PENDING`

// The run is queued.
const RunLifeCycleStateQueued RunLifeCycleState = `QUEUED`

// The task of this run is being executed.
const RunLifeCycleStateRunning RunLifeCycleState = `RUNNING`

// This run was aborted because a previous run of the same job was already
// active. This state is terminal.
const RunLifeCycleStateSkipped RunLifeCycleState = `SKIPPED`

// The task of this run has completed, and the cluster and execution context
// have been cleaned up. This state is terminal.
const RunLifeCycleStateTerminated RunLifeCycleState = `TERMINATED`

// The task of this run has completed, and the cluster and execution context are
// being cleaned up.
const RunLifeCycleStateTerminating RunLifeCycleState = `TERMINATING`

// The run is waiting for a retry.
const RunLifeCycleStateWaitingForRetry RunLifeCycleState = `WAITING_FOR_RETRY`

// String representation for [fmt.Print]
func (f *RunLifeCycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunLifeCycleState) Set(v string) error {
	switch v {
	case `BLOCKED`, `INTERNAL_ERROR`, `PENDING`, `QUEUED`, `RUNNING`, `SKIPPED`, `TERMINATED`, `TERMINATING`, `WAITING_FOR_RETRY`:
		*f = RunLifeCycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "INTERNAL_ERROR", "PENDING", "QUEUED", "RUNNING", "SKIPPED", "TERMINATED", "TERMINATING", "WAITING_FOR_RETRY"`, v)
	}
}

// Values returns all possible values for RunLifeCycleState.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunLifeCycleState) Values() []RunLifeCycleState {
	return []RunLifeCycleState{
		RunLifeCycleStateBlocked,
		RunLifeCycleStateInternalError,
		RunLifeCycleStatePending,
		RunLifeCycleStateQueued,
		RunLifeCycleStateRunning,
		RunLifeCycleStateSkipped,
		RunLifeCycleStateTerminated,
		RunLifeCycleStateTerminating,
		RunLifeCycleStateWaitingForRetry,
	}
}

// Type always returns RunLifeCycleState to satisfy [pflag.Value] interface
func (f *RunLifeCycleState) Type() string {
	return "RunLifeCycleState"
}

// The current state of the run.
type RunLifecycleStateV2State string

const RunLifecycleStateV2StateBlocked RunLifecycleStateV2State = `BLOCKED`

const RunLifecycleStateV2StatePending RunLifecycleStateV2State = `PENDING`

const RunLifecycleStateV2StateQueued RunLifecycleStateV2State = `QUEUED`

const RunLifecycleStateV2StateRunning RunLifecycleStateV2State = `RUNNING`

const RunLifecycleStateV2StateTerminated RunLifecycleStateV2State = `TERMINATED`

const RunLifecycleStateV2StateTerminating RunLifecycleStateV2State = `TERMINATING`

const RunLifecycleStateV2StateWaiting RunLifecycleStateV2State = `WAITING`

// String representation for [fmt.Print]
func (f *RunLifecycleStateV2State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunLifecycleStateV2State) Set(v string) error {
	switch v {
	case `BLOCKED`, `PENDING`, `QUEUED`, `RUNNING`, `TERMINATED`, `TERMINATING`, `WAITING`:
		*f = RunLifecycleStateV2State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "PENDING", "QUEUED", "RUNNING", "TERMINATED", "TERMINATING", "WAITING"`, v)
	}
}

// Values returns all possible values for RunLifecycleStateV2State.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunLifecycleStateV2State) Values() []RunLifecycleStateV2State {
	return []RunLifecycleStateV2State{
		RunLifecycleStateV2StateBlocked,
		RunLifecycleStateV2StatePending,
		RunLifecycleStateV2StateQueued,
		RunLifecycleStateV2StateRunning,
		RunLifecycleStateV2StateTerminated,
		RunLifecycleStateV2StateTerminating,
		RunLifecycleStateV2StateWaiting,
	}
}

// Type always returns RunLifecycleStateV2State to satisfy [pflag.Value] interface
func (f *RunLifecycleStateV2State) Type() string {
	return "RunLifecycleStateV2State"
}

type RunNow struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// The window the run aims to finish within, in hours, for a Cost-optimized
	// serverless run. Required only when `performance_target` is
	// `COST_OPTIMIZED`. Must be one of the supported values: `NO_FLEXIBILITY`,
	// `4h`, `12h`, `24h`, `72h`. `NO_FLEXIBILITY` means the run uses the
	// default cost-optimized performance level without additional cost-saving
	// optimizations.
	DurationTarget string `json:"duration_target,omitempty"`
	// An optional token to guarantee the idempotency of job run requests. If a
	// run with the provided token already exists, the request does not create a
	// new run but returns the ID of the existing run instead. If a run with the
	// provided token is deleted, an error is returned.
	//
	// If you specify the idempotency token, upon failure you can retry until
	// the request succeeds. Databricks guarantees that exactly one run is
	// launched with that idempotency token.
	//
	// This token must have at most 64 characters.
	IdempotencyToken string `json:"idempotency_token,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	JarParams []string `json:"jar_params,omitempty"`
	// The ID of the job to be executed
	JobId int64 `json:"job_id"`
	// Job-level parameters used in the run, for example `"param":
	// "overriding_val"`.
	JobParameters map[string]string `json:"job_parameters,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	NotebookParams map[string]string `json:"notebook_params,omitempty"`
	// A list of task keys to run inside of the job. If this field is not
	// provided, all tasks in the job will be run.
	//
	// Prefix a task key with `+` to also run its upstream tasks, or suffix it
	// with `+` to also run its downstream tasks. For example, `+my_task` runs
	// `my_task` and everything upstream of it, `my_task+` runs `my_task` and
	// everything downstream of it, and `+my_task+` runs both. A task key with
	// no `+` runs only that task.
	Only []string `json:"only,omitempty"`
	// The performance mode on a serverless job. The performance target
	// determines the level of compute performance or cost-efficiency for the
	// run. This field overrides the performance target defined on the job
	// level.
	//
	// * `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance. * `STANDARD`:
	// Enables cost-efficient execution of serverless workloads. *
	// `COST_OPTIMIZED`: Enables lower job costs by optimizing compute for your
	// selected target duration time. Must provide a duration target.
	PerformanceTarget PerformanceTarget `json:"performance_target,omitempty"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *PipelineParams `json:"pipeline_params,omitempty"`

	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	PythonParams []string `json:"python_params,omitempty"`
	// The queue settings of the run.
	Queue *QueueSettings `json:"queue,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	SqlParams map[string]string `json:"sql_params,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunNow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunNow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Run was started successfully.
type RunNowResponse struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// The globally unique ID of the newly triggered run.
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunNowResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunNowResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Run output was retrieved successfully.
type RunOutput struct {
	// The output of an AI function task, if available
	AiFunctionTaskOutput *AiFunctionTaskOutput `json:"ai_function_task_output,omitempty"`
	// The output of an AiRuntimeTask, if available — MLflow identifiers,
	// artifact paths, and per-replica allocated compute. Run lifecycle /
	// termination status lives on the surrounding framework `RunTask.status`
	// (`runs.proto:RunTask.status` of type `RunStatus`), not on this output.
	// See `tasks/genai/ai_runtime_task.proto:AiRuntimeTaskOutput`.
	AiRuntimeTaskOutput *AiRuntimeTaskOutput `json:"ai_runtime_task_output,omitempty"`
	// The output of an alert task, if available
	AlertOutput *AlertTaskOutput `json:"alert_output,omitempty"`
	// The output of a clean rooms notebook task, if available
	CleanRoomsNotebookOutput *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput `json:"clean_rooms_notebook_output,omitempty"`
	// The output of a dashboard task, if available
	DashboardOutput *DashboardTaskOutput `json:"dashboard_output,omitempty"`
	// Deprecated in favor of the new dbt_platform_output
	DbtCloudOutput *DbtCloudTaskOutput `json:"dbt_cloud_output,omitempty"`
	// The output of a dbt task, if available.
	DbtOutput *DbtOutput `json:"dbt_output,omitempty"`

	DbtPlatformOutput *DbtPlatformTaskOutput `json:"dbt_platform_output,omitempty"`
	// An error message indicating why a task failed or why output is not
	// available. The message is unstructured, and its exact format is subject
	// to change.
	Error string `json:"error,omitempty"`
	// If there was an error executing the run, this field contains any
	// available stack traces.
	ErrorTrace string `json:"error_trace,omitempty"`

	Info string `json:"info,omitempty"`
	// The output from tasks that write to standard streams (stdout/stderr) such
	// as spark_jar_task, spark_python_task, python_wheel_task.
	//
	// It's not supported for the notebook_task, pipeline_task or
	// spark_submit_task.
	//
	// Databricks restricts this API to return the last 5 MB of these logs.
	Logs string `json:"logs,omitempty"`
	// Whether the logs are truncated.
	LogsTruncated bool `json:"logs_truncated,omitempty"`
	// All details of the run except for its output.
	Metadata *Run `json:"metadata,omitempty"`
	// The output of a notebook task, if available. A notebook task that
	// terminates (either successfully or with a failure) without calling
	// `dbutils.notebook.exit()` is considered to have an empty output. This
	// field is set but its result value is empty. Databricks restricts this API
	// to return the first 5 MB of the output. To return a larger result, use
	// the [ClusterLogConf] field to configure log storage for the job cluster.
	//
	// [ClusterLogConf]: https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterlogconf
	NotebookOutput *NotebookOutput `json:"notebook_output,omitempty"`
	// The output of a run job task, if available
	RunJobOutput *RunJobOutput `json:"run_job_output,omitempty"`
	// The output of a SQL task, if available.
	SqlOutput *SqlOutput `json:"sql_output,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunParameters struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	JarParams []string `json:"jar_params,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	NotebookParams map[string]string `json:"notebook_params,omitempty"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *PipelineParams `json:"pipeline_params,omitempty"`

	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	PythonParams []string `json:"python_params,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	//
	// ⚠ **Deprecation note** Use [job parameters] to pass information down to
	// tasks.
	//
	// [job parameters]: https://docs.databricks.com/jobs/job-parameters.html#job-parameter-pushdown
	SqlParams map[string]string `json:"sql_params,omitempty"`
}

func (s *RunParameters) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A value indicating the run's result. The possible values are: * `SUCCESS`:
// The task completed successfully. * `FAILED`: The task completed with an
// error. * `TIMEDOUT`: The run was stopped after reaching the timeout. *
// `CANCELED`: The run was canceled at user request. *
// `MAXIMUM_CONCURRENT_RUNS_REACHED`: The run was skipped because the maximum
// concurrent runs were reached. * `EXCLUDED`: The run was skipped because the
// necessary conditions were not met. * `SUCCESS_WITH_FAILURES`: The job run
// completed successfully with some failures; leaf tasks were successful. *
// `UPSTREAM_FAILED`: The run was skipped because of an upstream failure. *
// `UPSTREAM_CANCELED`: The run was skipped because an upstream task was
// canceled. * `DISABLED`: The run was skipped because it was disabled
// explicitly by the user.
type RunResultState string

// The run was canceled at user request.
const RunResultStateCanceled RunResultState = `CANCELED`

// The run was skipped because it was disabled explicitly by the user.
const RunResultStateDisabled RunResultState = `DISABLED`

// The run was skipped because the necessary conditions were not met.
const RunResultStateExcluded RunResultState = `EXCLUDED`

// The task completed with an error.
const RunResultStateFailed RunResultState = `FAILED`

// The run was skipped because the maximum concurrent runs were reached.
const RunResultStateMaximumConcurrentRunsReached RunResultState = `MAXIMUM_CONCURRENT_RUNS_REACHED`

// The task completed successfully.
const RunResultStateSuccess RunResultState = `SUCCESS`

// The job run completed successfully with some failures; leaf tasks were
// successful.
const RunResultStateSuccessWithFailures RunResultState = `SUCCESS_WITH_FAILURES`

// The run was stopped after reaching the timeout.
const RunResultStateTimedout RunResultState = `TIMEDOUT`

// The run was skipped because an upstream task was canceled.
const RunResultStateUpstreamCanceled RunResultState = `UPSTREAM_CANCELED`

// The run was skipped because of an upstream failure.
const RunResultStateUpstreamFailed RunResultState = `UPSTREAM_FAILED`

// String representation for [fmt.Print]
func (f *RunResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunResultState) Set(v string) error {
	switch v {
	case `CANCELED`, `DISABLED`, `EXCLUDED`, `FAILED`, `MAXIMUM_CONCURRENT_RUNS_REACHED`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `TIMEDOUT`, `UPSTREAM_CANCELED`, `UPSTREAM_FAILED`:
		*f = RunResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "DISABLED", "EXCLUDED", "FAILED", "MAXIMUM_CONCURRENT_RUNS_REACHED", "SUCCESS", "SUCCESS_WITH_FAILURES", "TIMEDOUT", "UPSTREAM_CANCELED", "UPSTREAM_FAILED"`, v)
	}
}

// Values returns all possible values for RunResultState.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunResultState) Values() []RunResultState {
	return []RunResultState{
		RunResultStateCanceled,
		RunResultStateDisabled,
		RunResultStateExcluded,
		RunResultStateFailed,
		RunResultStateMaximumConcurrentRunsReached,
		RunResultStateSuccess,
		RunResultStateSuccessWithFailures,
		RunResultStateTimedout,
		RunResultStateUpstreamCanceled,
		RunResultStateUpstreamFailed,
	}
}

// Type always returns RunResultState to satisfy [pflag.Value] interface
func (f *RunResultState) Type() string {
	return "RunResultState"
}

// The current state of the run.
type RunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response. Note: Additional states might be
	// introduced in future releases.
	LifeCycleState RunLifeCycleState `json:"life_cycle_state,omitempty"`
	// The reason indicating why the run was queued.
	QueueReason string `json:"queue_reason,omitempty"`
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states. Note: Additional states might be introduced in
	// future releases.
	ResultState RunResultState `json:"result_state,omitempty"`
	// A descriptive message for the current state. This field is unstructured,
	// and its exact format is subject to change.
	StateMessage string `json:"state_message,omitempty"`
	// A value indicating whether a run was canceled manually by a user or by
	// the scheduler because the run timed out.
	UserCancelledOrTimedout bool `json:"user_cancelled_or_timedout,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The current status of the run
type RunStatus struct {
	// If the run was queued, details about the reason for queuing the run.
	QueueDetails *QueueDetails `json:"queue_details,omitempty"`

	State RunLifecycleStateV2State `json:"state,omitempty"`
	// If the run is in a TERMINATING or TERMINATED state, details about the
	// reason for terminating the run.
	TerminationDetails *TerminationDetails `json:"termination_details,omitempty"`
}

func (s *RunStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Used when outputting a child run, in GetRun or ListRuns.
type RunTask struct {
	// Applies an AI function to a set of input documents or rows.
	AiFunctionTask *AiFunctionTask `json:"ai_function_task,omitempty"`
	// The task runs a multi-gpu compute workload on Databricks AI Runtime.
	// Specify the accelerator type and count, the command to run, and where the
	// workload's code and MLflow output are stored.
	AiRuntimeTask *AiRuntimeTask `json:"ai_runtime_task,omitempty"`
	// The task evaluates a Databricks alert and sends notifications to
	// subscribers when the `alert_task` field is present.
	AlertTask *AlertTask `json:"alert_task,omitempty"`
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/clean-rooms/index.html
	CleanRoomsNotebookTask *CleanRoomsNotebookTask `json:"clean_rooms_notebook_task,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `json:"cluster_instance,omitempty"`
	// Task level compute configuration.
	Compute *Compute `json:"compute,omitempty"`
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask *RunConditionTask `json:"condition_task,omitempty"`
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	DashboardTask *DashboardTask `json:"dashboard_task,omitempty"`
	// Task type for dbt cloud, deprecated in favor of the new name
	// dbt_platform_task
	DbtCloudTask *DbtCloudTask `json:"dbt_cloud_task,omitempty"`

	DbtPlatformTask *DbtPlatformTask `json:"dbt_platform_task,omitempty"`
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask *DbtTask `json:"dbt_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn []TaskDependency `json:"depends_on,omitempty"`
	// An optional description for this task.
	Description string `json:"description,omitempty"`
	// An option to disable auto optimization in serverless
	DisableAutoOptimization bool `json:"disable_auto_optimization,omitempty"`
	// An optional flag to disable the task. If set to true, the task will not
	// run even if it is part of a job.
	Disabled bool `json:"disabled,omitempty"`
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance. * `STANDARD`:
	// Enables cost-efficient execution of serverless workloads. *
	// `COST_OPTIMIZED`: Enables lower job costs by optimizing compute for your
	// selected target duration time. Must provide a duration target.
	EffectivePerformanceTarget PerformanceTarget `json:"effective_performance_target,omitempty"`
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute or a compute resource that uses Environments mode.
	EnvironmentKey string `json:"environment_key,omitempty"`
	// Reference to a `JobEnvironmentVariables` entry defined in
	// `RunSettings.environment_variables` for one-time runs or preserved in
	// `Run.environment_variables` for run snapshots. The selected entry's
	// variables are applied to this task at execution time. This field supports
	// serverless tasks using environment version 5 or later.
	EnvironmentVariablesKey string `json:"environment_variables_key,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask *RunForEachTask `json:"for_each_task,omitempty"`

	GenAiComputeTask *GenAiComputeTask `json:"gen_ai_compute_task,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks. If `git_source` is set,
	// these tasks retrieve the file from the remote repository by default.
	// However, this behavior can be overridden by setting `source` to
	// `WORKSPACE` on the task. Note: dbt and SQL File tasks support only
	// version-controlled sources. If dbt or SQL File tasks are used,
	// `git_source` must be defined on the job.
	GitSource *GitSource `json:"git_source,omitempty"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string `json:"job_cluster_key,omitempty"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value `-1` means
	// to retry indefinitely and the value `0` means to never retry.
	MaxRetries int `json:"max_retries,omitempty"`
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	MinRetryIntervalMillis int `json:"min_retry_interval_millis,omitempty"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec `json:"new_cluster,omitempty"`
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings *TaskNotificationSettings `json:"notification_settings,omitempty"`
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask *PipelineTask `json:"pipeline_task,omitempty"`
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	PowerBiTask *PowerBiTask `json:"power_bi_task,omitempty"`
	// The task runs a Python operator task.
	PythonOperatorTask *PythonOperatorTask `json:"python_operator_task,omitempty"`
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration int64 `json:"queue_duration,omitempty"`
	// Parameter values including resolved references
	ResolvedValues *ResolvedValues `json:"resolved_values,omitempty"`
	// An optional policy to specify whether to retry a job when it times out.
	// The default behavior is to not retry on timeout.
	RetryOnTimeout bool `json:"retry_on_timeout,omitempty"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int64 `json:"run_duration,omitempty"`
	// The ID of the task run.
	RunId int64 `json:"run_id,omitempty"`
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf RunIf `json:"run_if,omitempty"`
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask *RunJobTask `json:"run_job_task,omitempty"`

	RunPageUrl string `json:"run_page_url,omitempty"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64 `json:"setup_duration,omitempty"`
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
	// (Legacy) The task runs the spark-submit script when the spark_submit_task
	// field is present. Databricks recommends using the spark_jar_task instead;
	// see [Spark Submit task for jobs](/jobs/spark-submit).
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask *SqlTask `json:"sql_task,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// Deprecated. Please use the `status` field instead.
	State *RunState `json:"state,omitempty"`

	Status *RunStatus `json:"status,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of a run. * `JOB_RUN`: Normal job run. A run created with
// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
// :method:jobs/submit.
//
// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
type RunType string

// Normal job run. A run created with :method:jobs/runNow.
const RunTypeJobRun RunType = `JOB_RUN`

// Submit run. A run created with :method:jobs/submit.
const RunTypeSubmitRun RunType = `SUBMIT_RUN`

// Workflow run. A run created with
// [dbutils.notebook.run](/dev-tools/databricks-utils.html#dbutils-workflow).
const RunTypeWorkflowRun RunType = `WORKFLOW_RUN`

// String representation for [fmt.Print]
func (f *RunType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunType) Set(v string) error {
	switch v {
	case `JOB_RUN`, `SUBMIT_RUN`, `WORKFLOW_RUN`:
		*f = RunType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "JOB_RUN", "SUBMIT_RUN", "WORKFLOW_RUN"`, v)
	}
}

// Values returns all possible values for RunType.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunType) Values() []RunType {
	return []RunType{
		RunTypeJobRun,
		RunTypeSubmitRun,
		RunTypeWorkflowRun,
	}
}

// Type always returns RunType to satisfy [pflag.Value] interface
func (f *RunType) Type() string {
	return "RunType"
}

// Runtime state for a schedule trigger. Currently empty because schedule
// triggers do not expose any trigger-specific runtime state.
type ScheduleTriggerState struct {
}

func (s *ScheduleTriggerState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL file
// will be retrieved\ from the local Databricks workspace. When set to `GIT`,
// the SQL file will be retrieved from a Git repository defined in `git_source`.
// If the value is empty, the task will use `GIT` if `git_source` is defined and
// `WORKSPACE` otherwise.
//
// * `WORKSPACE`: SQL file is located in Databricks workspace. * `GIT`: SQL file
// is located in cloud Git provider.
type Source string

// SQL file is located in cloud Git provider.
const SourceGit Source = `GIT`

// SQL file is located in <Databricks> workspace.
const SourceWorkspace Source = `WORKSPACE`

// String representation for [fmt.Print]
func (f *Source) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Source) Set(v string) error {
	switch v {
	case `GIT`, `WORKSPACE`:
		*f = Source(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GIT", "WORKSPACE"`, v)
	}
}

// Values returns all possible values for Source.
//
// There is no guarantee on the order of the values in the slice.
func (f *Source) Values() []Source {
	return []Source{
		SourceGit,
		SourceWorkspace,
	}
}

// Type always returns Source to satisfy [pflag.Value] interface
func (f *Source) Type() string {
	return "Source"
}

type SparkJarTask struct {
	// Deprecated since 04/2016. For classic compute, provide a `jar` through
	// the `libraries` field instead. For serverless compute, provide a `jar`
	// though the `java_dependencies` field inside the `environments` list.
	//
	// See the examples of classic and serverless compute usage at the top of
	// the page.
	JarUri string `json:"jar_uri,omitempty"`
	// The full name of the class containing the main method to be executed.
	// This class must be contained in a JAR provided as a library.
	//
	// The code must use `SparkContext.getOrCreate` to obtain a Spark context;
	// otherwise, runs of the job fail.
	MainClassName string `json:"main_class_name,omitempty"`
	// Parameters passed to the main method.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string `json:"parameters,omitempty"`
	// Deprecated. A value of `false` is no longer supported.
	RunAsRepl bool `json:"run_as_repl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SparkJarTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SparkJarTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SparkPythonTask struct {
	// Command line parameters passed to the Python file.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string `json:"parameters,omitempty"`
	// The Python file to be executed. Cloud file URIs (such as dbfs:/, s3:/,
	// adls:/, gcs:/) and workspace paths are supported. For python files stored
	// in the Databricks workspace, the path must be absolute and begin with
	// `/`. For files stored in a remote repository, the path must be relative.
	// This field is required.
	PythonFile string `json:"python_file"`
	// Optional location type of the Python file. When set to `WORKSPACE` or not
	// specified, the file will be retrieved from the local Databricks workspace
	// or cloud location (if the `python_file` has a URI format). When set to
	// `GIT`, the Python file will be retrieved from a Git repository defined in
	// `git_source`.
	//
	// * `WORKSPACE`: The Python file is located in a Databricks workspace or at
	// a cloud filesystem URI. * `GIT`: The Python file is located in a remote
	// Git repository.
	Source Source `json:"source,omitempty"`
}

func (s *SparkPythonTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SparkSubmitTask struct {
	// Command-line parameters passed to spark submit.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string `json:"parameters,omitempty"`
}

func (s *SparkSubmitTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SparseCheckout struct {
	// List of patterns to include for sparse checkout. Each pattern is the path
	// of a directory relative to the repository root.
	Patterns []string `json:"patterns,omitempty"`
}

func (s *SparseCheckout) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SqlAlertOutput struct {
	AlertState SqlAlertState `json:"alert_state,omitempty"`
	// The link to find the output results.
	OutputLink string `json:"output_link,omitempty"`
	// The text of the SQL query. Can Run permission of the SQL query associated
	// with the SQL alert is required to view this field.
	QueryText string `json:"query_text,omitempty"`
	// Information about SQL statements executed in the run.
	SqlStatements []SqlStatementOutput `json:"sql_statements,omitempty"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlAlertOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlAlertOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of the SQL alert.
//
// * UNKNOWN: alert yet to be evaluated * OK: alert evaluated and did not
// fulfill trigger conditions * TRIGGERED: alert evaluated and fulfilled trigger
// conditions
type SqlAlertState string

const SqlAlertStateOk SqlAlertState = `OK`

const SqlAlertStateTriggered SqlAlertState = `TRIGGERED`

const SqlAlertStateUnknown SqlAlertState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *SqlAlertState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SqlAlertState) Set(v string) error {
	switch v {
	case `OK`, `TRIGGERED`, `UNKNOWN`:
		*f = SqlAlertState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OK", "TRIGGERED", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for SqlAlertState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SqlAlertState) Values() []SqlAlertState {
	return []SqlAlertState{
		SqlAlertStateOk,
		SqlAlertStateTriggered,
		SqlAlertStateUnknown,
	}
}

// Type always returns SqlAlertState to satisfy [pflag.Value] interface
func (f *SqlAlertState) Type() string {
	return "SqlAlertState"
}

type SqlConditionConfiguration struct {
	// The ID of the SQL query to evaluate as the trigger condition. Exactly one
	// of `sql_query_id` or `file` must be set.
	SqlQueryId string `json:"sql_query_id,omitempty"`
	// Determines how the SQL query result is interpreted to decide whether the
	// condition fires. Must be set to a recognized value when provided. When
	// unset on an existing serialized configuration, the server preserves the
	// original semantics by interpreting it as `QUERY_RETURNS_ROWS`. New
	// configurations should set this explicitly — explicit
	// `SQL_CONDITION_TRIGGER_MODE_UNSPECIFIED` is rejected at validation.
	TriggerMode SqlConditionTriggerMode `json:"trigger_mode,omitempty"`
	// The canonical identifier of the SQL warehouse to run the condition query
	// against.
	WarehouseId string `json:"warehouse_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlConditionConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlConditionConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// SQL condition evaluation details captured at the time the run was triggered
type SqlConditionRunInfoDetails struct {
	// Whether the last condition evaluation was satisfied. For
	// QUERY_RETURNS_ROWS, the query returned at least one row. For
	// RESULT_VALUE_CHANGES, the result value changed.
	ConditionEvaluationSatisfied bool `json:"condition_evaluation_satisfied,omitempty"`
	// The ID of the SQL session of the condition evaluation. Not set by current
	// evaluations. Use condition_evaluation_sql_statement_id instead.
	ConditionEvaluationSqlSessionId string `json:"condition_evaluation_sql_session_id,omitempty"`
	// The SQL statement ID of the condition evaluation. Set for both the
	// QUERY_RETURNS_ROWS and RESULT_VALUE_CHANGES trigger modes.
	ConditionEvaluationSqlStatementId string `json:"condition_evaluation_sql_statement_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlConditionRunInfoDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlConditionRunInfoDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlConditionState struct {
	// Whether the last condition evaluation was satisfied. For
	// QUERY_RETURNS_ROWS, the query returned at least one row. For
	// RESULT_VALUE_CHANGES, the result value changed.
	LatestConditionEvaluationSatisfied bool `json:"latest_condition_evaluation_satisfied,omitempty"`
	// The ID of the SQL session of the latest condition evaluation. Not set by
	// current evaluations. Use latest_condition_evaluation_sql_statement_id
	// instead.
	LatestConditionEvaluationSqlSessionId string `json:"latest_condition_evaluation_sql_session_id,omitempty"`
	// The ID of the SQL statement executed for the latest condition evaluation.
	// Populated for both QUERY_RETURNS_ROWS and RESULT_VALUE_CHANGES.
	LatestConditionEvaluationSqlStatementId string `json:"latest_condition_evaluation_sql_statement_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlConditionState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlConditionState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The strategy used to evaluate a SQL condition trigger against a query result
// set.
//
// * `SQL_CONDITION_TRIGGER_MODE_UNSPECIFIED`: Sentinel zero-value. Not a valid
// input — the validator rejects this when sent explicitly. Internally treated
// as `QUERY_RETURNS_ROWS` when reading legacy data that predates this field. *
// `QUERY_RETURNS_ROWS`: Fires whenever the result set has at least one row.
// Zero rows means the condition is not met. This is the original SQL condition
// behavior. * `RESULT_VALUE_CHANGES`: Fires whenever the query's single result
// value differs from the previous evaluation. The first evaluation stores the
// result as a baseline and does not fire. Queries must return exactly one cell
// (one row, one column).
type SqlConditionTriggerMode string

// Fires whenever the result set has at least one row. Zero rows means
const SqlConditionTriggerModeQueryReturnsRows SqlConditionTriggerMode = `QUERY_RETURNS_ROWS`

// Fires whenever the query's single result value differs from the
const SqlConditionTriggerModeResultValueChanges SqlConditionTriggerMode = `RESULT_VALUE_CHANGES`

// String representation for [fmt.Print]
func (f *SqlConditionTriggerMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SqlConditionTriggerMode) Set(v string) error {
	switch v {
	case `QUERY_RETURNS_ROWS`, `RESULT_VALUE_CHANGES`:
		*f = SqlConditionTriggerMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "QUERY_RETURNS_ROWS", "RESULT_VALUE_CHANGES"`, v)
	}
}

// Values returns all possible values for SqlConditionTriggerMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *SqlConditionTriggerMode) Values() []SqlConditionTriggerMode {
	return []SqlConditionTriggerMode{
		SqlConditionTriggerModeQueryReturnsRows,
		SqlConditionTriggerModeResultValueChanges,
	}
}

// Type always returns SqlConditionTriggerMode to satisfy [pflag.Value] interface
func (f *SqlConditionTriggerMode) Type() string {
	return "SqlConditionTriggerMode"
}

type SqlDashboardOutput struct {
	// The canonical identifier of the SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`
	// Widgets executed in the run. Only SQL query based widgets are listed.
	Widgets []SqlDashboardWidgetOutput `json:"widgets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlDashboardOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlDashboardOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlDashboardWidgetOutput struct {
	// Time (in epoch milliseconds) when execution of the SQL widget ends.
	EndTime int64 `json:"end_time,omitempty"`
	// The information about the error when execution fails.
	Error *SqlOutputError `json:"error,omitempty"`
	// The link to find the output results.
	OutputLink string `json:"output_link,omitempty"`
	// Time (in epoch milliseconds) when execution of the SQL widget starts.
	StartTime int64 `json:"start_time,omitempty"`
	// The execution status of the SQL widget.
	Status SqlDashboardWidgetOutputStatus `json:"status,omitempty"`
	// The canonical identifier of the SQL widget.
	WidgetId string `json:"widget_id,omitempty"`
	// The title of the SQL widget.
	WidgetTitle string `json:"widget_title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlDashboardWidgetOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlDashboardWidgetOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlDashboardWidgetOutputStatus string

const SqlDashboardWidgetOutputStatusCancelled SqlDashboardWidgetOutputStatus = `CANCELLED`

const SqlDashboardWidgetOutputStatusFailed SqlDashboardWidgetOutputStatus = `FAILED`

const SqlDashboardWidgetOutputStatusPending SqlDashboardWidgetOutputStatus = `PENDING`

const SqlDashboardWidgetOutputStatusRunning SqlDashboardWidgetOutputStatus = `RUNNING`

const SqlDashboardWidgetOutputStatusSuccess SqlDashboardWidgetOutputStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *SqlDashboardWidgetOutputStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SqlDashboardWidgetOutputStatus) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCESS`:
		*f = SqlDashboardWidgetOutputStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "PENDING", "RUNNING", "SUCCESS"`, v)
	}
}

// Values returns all possible values for SqlDashboardWidgetOutputStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *SqlDashboardWidgetOutputStatus) Values() []SqlDashboardWidgetOutputStatus {
	return []SqlDashboardWidgetOutputStatus{
		SqlDashboardWidgetOutputStatusCancelled,
		SqlDashboardWidgetOutputStatusFailed,
		SqlDashboardWidgetOutputStatusPending,
		SqlDashboardWidgetOutputStatusRunning,
		SqlDashboardWidgetOutputStatusSuccess,
	}
}

// Type always returns SqlDashboardWidgetOutputStatus to satisfy [pflag.Value] interface
func (f *SqlDashboardWidgetOutputStatus) Type() string {
	return "SqlDashboardWidgetOutputStatus"
}

type SqlOutput struct {
	// The output of a SQL alert task, if available.
	AlertOutput *SqlAlertOutput `json:"alert_output,omitempty"`
	// The output of a SQL dashboard task, if available.
	DashboardOutput *SqlDashboardOutput `json:"dashboard_output,omitempty"`
	// The output of a SQL file task, if available.
	FileOutput *SqlTaskSqlFileOutput `json:"file_output,omitempty"`
	// The output of a SQL query task, if available.
	QueryOutput *SqlQueryOutput `json:"query_output,omitempty"`
}

func (s *SqlOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SqlOutputError struct {
	// The error message when execution fails.
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlOutputError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlOutputError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlQueryOutput struct {
	EndpointId string `json:"endpoint_id,omitempty"`
	// Deprecated. The link to find the output results, which are only
	// accessible from the UI.
	OutputLink string `json:"output_link,omitempty"`
	// The text of the SQL query. Can Run permission of the SQL query is
	// required to view this field.
	QueryText string `json:"query_text,omitempty"`
	// Information about SQL statements executed in the run.
	SqlStatements []SqlStatementOutput `json:"sql_statements,omitempty"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlQueryOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlQueryOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlStatementOutput struct {
	// A key that can be used to look up query details.
	LookupKey string `json:"lookup_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlStatementOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlStatementOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlStatementQueryResult struct {
	// The metadata for each column in the result set.
	Columns []SqlStatementQueryResultColumn `json:"columns,omitempty"`
	// The rows returned by the query.
	Rows []SqlStatementQueryResultRow `json:"rows,omitempty"`
	// Whether the result set was truncated.
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlStatementQueryResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlStatementQueryResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlStatementQueryResultColumn struct {
	// The name of the column.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlStatementQueryResultColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlStatementQueryResultColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlStatementQueryResultRow struct {
	// The values in the row, ordered to match the columns.
	Values []SqlStatementQueryResultRowRowValue `json:"values,omitempty"`
}

func (s *SqlStatementQueryResultRow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SqlStatementQueryResultRowRowValue struct {
	// The value rendered as a string.
	StringValue string `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlStatementQueryResultRowRowValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlStatementQueryResultRowRowValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlTask struct {
	// If alert, indicates that this job must refresh a SQL alert.
	Alert *SqlTaskAlert `json:"alert,omitempty"`
	// If dashboard, indicates that this job must refresh a SQL dashboard.
	Dashboard *SqlTaskDashboard `json:"dashboard,omitempty"`
	// If file, indicates that this job runs a SQL file in a remote Git
	// repository.
	File *SqlTaskFile `json:"file,omitempty"`
	// Parameters to be used for each run of this job. The SQL alert task does
	// not support custom parameters.
	Parameters map[string]string `json:"parameters,omitempty"`
	// If query, indicates that this job must execute a SQL query.
	Query *SqlTaskQuery `json:"query,omitempty"`
	// The canonical identifier of the SQL warehouse. Recommended to use with
	// serverless or pro SQL warehouses. Classic SQL warehouses are only
	// supported for SQL alert, dashboard and query tasks and are limited to
	// scheduled single-task jobs.
	WarehouseId string `json:"warehouse_id"`
}

func (s *SqlTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SqlTaskAlert struct {
	// The canonical identifier of the SQL alert.
	AlertId string `json:"alert_id"`
	// If true, the alert notifications are not sent to subscribers.
	PauseSubscriptions bool `json:"pause_subscriptions,omitempty"`
	// If specified, alert notifications are sent to subscribers.
	Subscriptions []SqlTaskSubscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlTaskAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlTaskAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlTaskDashboard struct {
	// Subject of the email sent to subscribers of this task.
	CustomSubject string `json:"custom_subject,omitempty"`
	// The canonical identifier of the SQL dashboard.
	DashboardId string `json:"dashboard_id"`
	// If true, the dashboard snapshot is not taken, and emails are not sent to
	// subscribers.
	PauseSubscriptions bool `json:"pause_subscriptions,omitempty"`
	// If specified, dashboard snapshots are sent to subscriptions.
	Subscriptions []SqlTaskSubscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlTaskDashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlTaskDashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlTaskFile struct {
	// Path of the SQL file. Must be relative if the source is a remote Git
	// repository and absolute for workspace paths.
	Path string `json:"path"`
	// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL
	// file will be retrieved from the local Databricks workspace. When set to
	// `GIT`, the SQL file will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: SQL file is located in Databricks workspace. * `GIT`: SQL
	// file is located in cloud Git provider.
	Source Source `json:"source,omitempty"`
}

func (s *SqlTaskFile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SqlTaskQuery struct {
	// The canonical identifier of the SQL query.
	QueryId string `json:"query_id"`
}

func (s *SqlTaskQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SqlTaskSqlFileOutput struct {
	// The result of the executed SQL statement, if available.
	QueryResult *SqlStatementQueryResult `json:"query_result,omitempty"`
	// The text of the SQL query.
	QueryText string `json:"query_text,omitempty"`
	// Information about SQL statements executed in the run.
	SqlStatements []SqlStatementOutput `json:"sql_statements,omitempty"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlTaskSqlFileOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlTaskSqlFileOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlTaskSubscription struct {
	// The canonical identifier of the destination to receive email
	// notification. This parameter is mutually exclusive with user_name. You
	// cannot set both destination_id and user_name for subscription
	// notifications.
	DestinationId string `json:"destination_id,omitempty"`
	// The user name to receive the subscription email. This parameter is
	// mutually exclusive with destination_id. You cannot set both
	// destination_id and user_name for subscription notifications.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SqlTaskSubscription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlTaskSubscription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StorageMode string

const StorageModeDirectQuery StorageMode = `DIRECT_QUERY`

const StorageModeDual StorageMode = `DUAL`

const StorageModeImport StorageMode = `IMPORT`

// String representation for [fmt.Print]
func (f *StorageMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *StorageMode) Set(v string) error {
	switch v {
	case `DIRECT_QUERY`, `DUAL`, `IMPORT`:
		*f = StorageMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECT_QUERY", "DUAL", "IMPORT"`, v)
	}
}

// Values returns all possible values for StorageMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *StorageMode) Values() []StorageMode {
	return []StorageMode{
		StorageModeDirectQuery,
		StorageModeDual,
		StorageModeImport,
	}
}

// Type always returns StorageMode to satisfy [pflag.Value] interface
func (f *StorageMode) Type() string {
	return "StorageMode"
}

type SubmitRun struct {
	// List of permissions to set on the job.
	AccessControlList []JobAccessControlRequest `json:"access_control_list,omitempty"`
	// The user specified id of the budget policy to use for this one-time run.
	// If not specified, the run will be not be attributed to any budget policy.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The window the run aims to finish within, in hours, for a Cost-optimized
	// serverless run. Required only when `performance_target` is
	// `COST_OPTIMIZED`. Must be one of the supported values: `NO_FLEXIBILITY`,
	// `4h`, `12h`, `24h`, `72h`. `NO_FLEXIBILITY` means the run uses the
	// default cost-optimized performance level without additional cost-saving
	// optimizations.
	DurationTarget string `json:"duration_target,omitempty"`
	// An optional set of email addresses notified when the run begins or
	// completes.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// Named environment-variable entries that tasks of this one-time run can
	// reference by key from `RunTaskSettings.environment_variables_key`. Each
	// entry's `spec` holds inline `variables` and optional `.env` `files`.
	// Handled identically to `JobSettings.environment_variables`. Maximum 10
	// entries. Entries are independent of one another — there is no
	// cross-entry merging.
	EnvironmentVariables []JobEnvironmentVariables `json:"environment_variables,omitempty"`
	// A list of task execution environment specifications that can be
	// referenced by tasks of this run.
	Environments []JobEnvironment `json:"environments,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *GitSource `json:"git_source,omitempty"`

	Health *JobsHealthRules `json:"health,omitempty"`
	// An optional token that can be used to guarantee the idempotency of job
	// run requests. If a run with the provided token already exists, the
	// request does not create a new run but returns the ID of the existing run
	// instead. If a run with the provided token is deleted, an error is
	// returned.
	//
	// If you specify the idempotency token, upon failure you can retry until
	// the request succeeds. Databricks guarantees that exactly one run is
	// launched with that idempotency token.
	//
	// This token must have at most 64 characters.
	IdempotencyToken string `json:"idempotency_token,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// run.
	NotificationSettings *JobNotificationSettings `json:"notification_settings,omitempty"`
	// The performance mode on a serverless one-time run. This field determines
	// the level of compute performance or cost-efficiency for the run. The
	// performance target does not apply to tasks that run on Serverless GPU
	// compute.
	//
	// * `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance. * `STANDARD`:
	// Enables cost-efficient execution of serverless workloads. *
	// `COST_OPTIMIZED`: Enables lower job costs by optimizing compute for your
	// selected target duration time. Must provide a duration target.
	PerformanceTarget PerformanceTarget `json:"performance_target,omitempty"`
	// The queue settings of the one-time run.
	Queue *QueueSettings `json:"queue,omitempty"`
	// Specifies the user or service principal that the job runs as. If not
	// specified, the job runs as the user who submits the request.
	RunAs *JobRunAs `json:"run_as,omitempty"`
	// An optional name for the run. The default value is `Untitled`.
	RunName string `json:"run_name,omitempty"`

	Tasks []SubmitTask `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// The user specified id of the usage policy to use for this one-time run.
	// If not specified, a default usage policy may be applied when creating or
	// modifying the job.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SubmitRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SubmitRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Run was created and started successfully.
type SubmitRunResponse struct {
	// The canonical identifier for the newly submitted run.
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SubmitRunResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SubmitRunResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SubmitTask struct {
	// Applies an AI function to a set of input documents or rows.
	AiFunctionTask *AiFunctionTask `json:"ai_function_task,omitempty"`
	// The task runs a multi-gpu compute workload on Databricks AI Runtime.
	// Specify the accelerator type and count, the command to run, and where the
	// workload's code and MLflow output are stored.
	AiRuntimeTask *AiRuntimeTask `json:"ai_runtime_task,omitempty"`
	// The task evaluates a Databricks alert and sends notifications to
	// subscribers when the `alert_task` field is present.
	AlertTask *AlertTask `json:"alert_task,omitempty"`
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/clean-rooms/index.html
	CleanRoomsNotebookTask *CleanRoomsNotebookTask `json:"clean_rooms_notebook_task,omitempty"`
	// Task level compute configuration.
	Compute *Compute `json:"compute,omitempty"`
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask *ConditionTask `json:"condition_task,omitempty"`
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	DashboardTask *DashboardTask `json:"dashboard_task,omitempty"`
	// Task type for dbt cloud, deprecated in favor of the new name
	// dbt_platform_task
	DbtCloudTask *DbtCloudTask `json:"dbt_cloud_task,omitempty"`

	DbtPlatformTask *DbtPlatformTask `json:"dbt_platform_task,omitempty"`
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask *DbtTask `json:"dbt_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn []TaskDependency `json:"depends_on,omitempty"`
	// An optional description for this task.
	Description string `json:"description,omitempty"`
	// An option to disable auto optimization in serverless
	DisableAutoOptimization bool `json:"disable_auto_optimization,omitempty"`
	// An optional flag to disable the task. If set to true, the task will not
	// run even if it is part of a job.
	Disabled bool `json:"disabled,omitempty"`
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute or a compute resource that uses Environments mode.
	EnvironmentKey string `json:"environment_key,omitempty"`
	// Reference to a `JobEnvironmentVariables` entry defined in
	// `RunSettings.environment_variables` for one-time runs or preserved in
	// `Run.environment_variables` for run snapshots. The selected entry's
	// variables are applied to this task at execution time. This field supports
	// serverless tasks using environment version 5 or later.
	EnvironmentVariablesKey string `json:"environment_variables_key,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask *ForEachTask `json:"for_each_task,omitempty"`

	GenAiComputeTask *GenAiComputeTask `json:"gen_ai_compute_task,omitempty"`

	Health *JobsHealthRules `json:"health,omitempty"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value `-1` means
	// to retry indefinitely and the value `0` means to never retry.
	MaxRetries int `json:"max_retries,omitempty"`
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	MinRetryIntervalMillis int `json:"min_retry_interval_millis,omitempty"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec `json:"new_cluster,omitempty"`
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings *TaskNotificationSettings `json:"notification_settings,omitempty"`
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask *PipelineTask `json:"pipeline_task,omitempty"`
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	PowerBiTask *PowerBiTask `json:"power_bi_task,omitempty"`
	// The task runs a Python operator task.
	PythonOperatorTask *PythonOperatorTask `json:"python_operator_task,omitempty"`
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty"`
	// An optional policy to specify whether to retry a job when it times out.
	// The default behavior is to not retry on timeout.
	RetryOnTimeout bool `json:"retry_on_timeout,omitempty"`
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf RunIf `json:"run_if,omitempty"`
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask *RunJobTask `json:"run_job_task,omitempty"`
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
	// (Legacy) The task runs the spark-submit script when the spark_submit_task
	// field is present. Databricks recommends using the spark_jar_task instead;
	// see [Spark Submit task for jobs](/jobs/spark-submit).
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask *SqlTask `json:"sql_task,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SubmitTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SubmitTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Subscription struct {
	// Optional: The configuration for the files attached to the dashboard
	// snapshot notification. It controls the underlying data of dashboard
	// widgets that is attached as data files and which pages are included in
	// the attached PDF. If not specified, the notification is sent with only
	// the default snapshot.
	Attachment *SubscriptionAttachment `json:"attachment,omitempty"`
	// Optional: Allows users to specify a custom subject line on the email sent
	// to subscribers.
	CustomSubject string `json:"custom_subject,omitempty"`
	// When true, the subscription will not send emails.
	Paused bool `json:"paused,omitempty"`
	// The list of subscribers to send the snapshot of the dashboard to.
	Subscribers []SubscriptionSubscriber `json:"subscribers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Subscription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Subscription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The configuration for the files attached to the dashboard snapshot
// notification sent to subscribers. `data_attachment` describes the underlying
// data of dashboard widgets attached as data files, and `page_attachment`
// describes which dashboard pages are included in the attached PDF.
type SubscriptionAttachment struct {
	// Optional: The configuration for the underlying data of dashboard widgets
	// attached as tabular data files. If not specified, no tabular data is
	// attached.
	DataAttachment *SubscriptionAttachmentDataAttachment `json:"data_attachment,omitempty"`
	// Optional: The selection of dashboard pages included in the attached PDF.
	// If not specified, only the first page is included.
	PageAttachment *SubscriptionAttachmentPageAttachment `json:"page_attachment,omitempty"`
}

func (s *SubscriptionAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The configuration for the tabular data files attached to the dashboard
// snapshot. `widgets` lists the widgets whose data is exported and
// `file_format` is the output format; each referenced widget's data is exported
// as a separate file.
type SubscriptionAttachmentDataAttachment struct {
	// The file format of the exported tabular data files.
	FileFormat SubscriptionAttachmentDataAttachmentFileFormat `json:"file_format"`
	// When true, the dashboard's applied filter values are included alongside
	// the tabular data as an additional attached file.
	IncludeAppliedFilters bool `json:"include_applied_filters,omitempty"`
	// The widgets whose underlying data is exported. At least one widget must
	// be referenced.
	Widgets []SubscriptionAttachmentDataAttachmentWidget `json:"widgets"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SubscriptionAttachmentDataAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SubscriptionAttachmentDataAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The output format of an exported tabular data file. Must be set to a value
// other than `FILE_FORMAT_UNSPECIFIED`.
type SubscriptionAttachmentDataAttachmentFileFormat string

const SubscriptionAttachmentDataAttachmentFileFormatFileFormatCsv SubscriptionAttachmentDataAttachmentFileFormat = `FILE_FORMAT_CSV`

const SubscriptionAttachmentDataAttachmentFileFormatFileFormatExcel SubscriptionAttachmentDataAttachmentFileFormat = `FILE_FORMAT_EXCEL`

const SubscriptionAttachmentDataAttachmentFileFormatFileFormatTsv SubscriptionAttachmentDataAttachmentFileFormat = `FILE_FORMAT_TSV`

// String representation for [fmt.Print]
func (f *SubscriptionAttachmentDataAttachmentFileFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SubscriptionAttachmentDataAttachmentFileFormat) Set(v string) error {
	switch v {
	case `FILE_FORMAT_CSV`, `FILE_FORMAT_EXCEL`, `FILE_FORMAT_TSV`:
		*f = SubscriptionAttachmentDataAttachmentFileFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE_FORMAT_CSV", "FILE_FORMAT_EXCEL", "FILE_FORMAT_TSV"`, v)
	}
}

// Values returns all possible values for SubscriptionAttachmentDataAttachmentFileFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *SubscriptionAttachmentDataAttachmentFileFormat) Values() []SubscriptionAttachmentDataAttachmentFileFormat {
	return []SubscriptionAttachmentDataAttachmentFileFormat{
		SubscriptionAttachmentDataAttachmentFileFormatFileFormatCsv,
		SubscriptionAttachmentDataAttachmentFileFormatFileFormatExcel,
		SubscriptionAttachmentDataAttachmentFileFormatFileFormatTsv,
	}
}

// Type always returns SubscriptionAttachmentDataAttachmentFileFormat to satisfy [pflag.Value] interface
func (f *SubscriptionAttachmentDataAttachmentFileFormat) Type() string {
	return "SubscriptionAttachmentDataAttachmentFileFormat"
}

// Identifies a dashboard widget whose data to attach. `page_id` and `widget_id`
// are the URL identifiers from the dashboard's serialized definition returned
// by the Get dashboard API (`serialized_dashboard`), and together they uniquely
// identify the widget.
type SubscriptionAttachmentDataAttachmentWidget struct {
	// The URL identifier of the page containing the widget. This is the value
	// of the `name` field of the corresponding entry under `pages` in the
	// serialized dashboard definition. Format: alphanumeric, hyphen, and
	// underscore characters allowed, max length 63.
	PageId string `json:"page_id"`
	// The URL identifier of the widget. This is the value of the `name` field
	// of the corresponding widget under that page's `layout` in the serialized
	// dashboard definition. Format: alphanumeric, hyphen, and underscore
	// characters allowed, max length 63.
	WidgetId string `json:"widget_id"`
}

func (s *SubscriptionAttachmentDataAttachmentWidget) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The selection of dashboard pages included in the attached PDF.
// `page_selection` chooses between all pages, the first page only, or an
// explicit subset listed in `page_ids`.
type SubscriptionAttachmentPageAttachment struct {
	// The identifiers of the pages to include in the PDF. Only specified when
	// `page_selection` is `PAGE_SELECTION_SUBSET`. Each value is the identifier
	// of a page in the dashboard's serialized definition returned by the Get
	// dashboard API (`serialized_dashboard`): the value of the `name` field of
	// an entry under `pages`. Format: alphanumeric, hyphen, and underscore
	// characters allowed, max length 63.
	PageIds []string `json:"page_ids,omitempty"`
	// How the attached pages are selected. When set to `PAGE_SELECTION_SUBSET`,
	// the pages to include are listed in `page_ids`.
	PageSelection SubscriptionAttachmentPageAttachmentPageSelection `json:"page_selection,omitempty"`
}

func (s *SubscriptionAttachmentPageAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Determines which dashboard pages are attached to the notification.
type SubscriptionAttachmentPageAttachmentPageSelection string

const SubscriptionAttachmentPageAttachmentPageSelectionPageSelectionAll SubscriptionAttachmentPageAttachmentPageSelection = `PAGE_SELECTION_ALL`

const SubscriptionAttachmentPageAttachmentPageSelectionPageSelectionSubset SubscriptionAttachmentPageAttachmentPageSelection = `PAGE_SELECTION_SUBSET`

// String representation for [fmt.Print]
func (f *SubscriptionAttachmentPageAttachmentPageSelection) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SubscriptionAttachmentPageAttachmentPageSelection) Set(v string) error {
	switch v {
	case `PAGE_SELECTION_ALL`, `PAGE_SELECTION_SUBSET`:
		*f = SubscriptionAttachmentPageAttachmentPageSelection(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAGE_SELECTION_ALL", "PAGE_SELECTION_SUBSET"`, v)
	}
}

// Values returns all possible values for SubscriptionAttachmentPageAttachmentPageSelection.
//
// There is no guarantee on the order of the values in the slice.
func (f *SubscriptionAttachmentPageAttachmentPageSelection) Values() []SubscriptionAttachmentPageAttachmentPageSelection {
	return []SubscriptionAttachmentPageAttachmentPageSelection{
		SubscriptionAttachmentPageAttachmentPageSelectionPageSelectionAll,
		SubscriptionAttachmentPageAttachmentPageSelectionPageSelectionSubset,
	}
}

// Type always returns SubscriptionAttachmentPageAttachmentPageSelection to satisfy [pflag.Value] interface
func (f *SubscriptionAttachmentPageAttachmentPageSelection) Type() string {
	return "SubscriptionAttachmentPageAttachmentPageSelection"
}

type SubscriptionSubscriber struct {
	// A snapshot of the dashboard will be sent to the destination when the
	// `destination_id` field is present.
	DestinationId string `json:"destination_id,omitempty"`
	// A snapshot of the dashboard will be sent to the user's email when the
	// `user_name` field is present.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SubscriptionSubscriber) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SubscriptionSubscriber) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableState struct {
	// Whether or not the table has seen updates since either the creation of
	// the trigger or the last successful evaluation of the trigger
	HasSeenUpdates bool `json:"has_seen_updates,omitempty"`
	// Full table name of the table to monitor, e.g.
	// `mycatalog.myschema.mytable`
	TableName string `json:"table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableTriggerState struct {
	// The last observed state of each monitored table.
	LastSeenTableStates []TableState `json:"last_seen_table_states,omitempty"`
	// Indicates whether the trigger is using scalable monitoring.
	UsingScalableMonitoring bool `json:"using_scalable_monitoring,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableTriggerState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableTriggerState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableUpdateTriggerConfiguration struct {
	// The table(s) condition based on which to trigger a job run.
	Condition Condition `json:"condition,omitempty"`
	// If set, the trigger starts a run only after the specified amount of time
	// has passed since the last time the trigger fired. The minimum allowed
	// value is 60 seconds.
	MinTimeBetweenTriggersSeconds int `json:"min_time_between_triggers_seconds,omitempty"`
	// A list of tables to monitor for changes. The table name must be in the
	// format `catalog_name.schema_name.table_name`.
	TableNames []string `json:"table_names"`
	// If set, the trigger starts a run only after no table updates have
	// occurred for the specified time and can be used to wait for a series of
	// table updates before triggering a run. The minimum allowed value is 60
	// seconds.
	WaitAfterLastChangeSeconds int `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableUpdateTriggerConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableUpdateTriggerConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Task struct {
	// Applies an AI function to a set of input documents or rows.
	AiFunctionTask *AiFunctionTask `json:"ai_function_task,omitempty"`
	// The task runs a multi-gpu compute workload on Databricks AI Runtime.
	// Specify the accelerator type and count, the command to run, and where the
	// workload's code and MLflow output are stored.
	AiRuntimeTask *AiRuntimeTask `json:"ai_runtime_task,omitempty"`
	// The task evaluates a Databricks alert and sends notifications to
	// subscribers when the `alert_task` field is present.
	AlertTask *AlertTask `json:"alert_task,omitempty"`
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/clean-rooms/index.html
	CleanRoomsNotebookTask *CleanRoomsNotebookTask `json:"clean_rooms_notebook_task,omitempty"`
	// Task level compute configuration.
	Compute *Compute `json:"compute,omitempty"`
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask *ConditionTask `json:"condition_task,omitempty"`
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	DashboardTask *DashboardTask `json:"dashboard_task,omitempty"`
	// Task type for dbt cloud, deprecated in favor of the new name
	// dbt_platform_task
	DbtCloudTask *DbtCloudTask `json:"dbt_cloud_task,omitempty"`

	DbtPlatformTask *DbtPlatformTask `json:"dbt_platform_task,omitempty"`
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask *DbtTask `json:"dbt_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete before executing this
	// task. The task will run only if the `run_if` condition is true. The key
	// is `task_key`, and the value is the name assigned to the dependent task.
	DependsOn []TaskDependency `json:"depends_on,omitempty"`
	// An optional description for this task.
	Description string `json:"description,omitempty"`
	// An option to disable auto optimization in serverless
	DisableAutoOptimization bool `json:"disable_auto_optimization,omitempty"`
	// An optional flag to disable the task. If set to true, the task will not
	// run even if it is part of a job.
	Disabled bool `json:"disabled,omitempty"`
	// An optional set of email addresses that is notified when runs of this
	// task begin or complete as well as when this task is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *TaskEmailNotifications `json:"email_notifications,omitempty"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute or a compute resource that uses Environments mode.
	EnvironmentKey string `json:"environment_key,omitempty"`
	// Reference to a `JobEnvironmentVariables` entry defined in
	// `JobSettings.environment_variables`. The selected entry's variables are
	// applied to this task at execution time. This field supports serverless
	// tasks using environment version 5 or later.
	EnvironmentVariablesKey string `json:"environment_variables_key,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask *ForEachTask `json:"for_each_task,omitempty"`

	GenAiComputeTask *GenAiComputeTask `json:"gen_ai_compute_task,omitempty"`
	// An optional set of health rules that can be defined for this task.
	Health *JobsHealthRules `json:"health,omitempty"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string `json:"job_cluster_key,omitempty"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value `-1` means
	// to retry indefinitely and the value `0` means to never retry.
	MaxRetries int `json:"max_retries,omitempty"`
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	MinRetryIntervalMillis int `json:"min_retry_interval_millis,omitempty"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec `json:"new_cluster,omitempty"`
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task.
	NotificationSettings *TaskNotificationSettings `json:"notification_settings,omitempty"`
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask *PipelineTask `json:"pipeline_task,omitempty"`
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	PowerBiTask *PowerBiTask `json:"power_bi_task,omitempty"`
	// The task runs a Python operator task.
	PythonOperatorTask *PythonOperatorTask `json:"python_operator_task,omitempty"`
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty"`
	// An optional policy to specify whether to retry a job when it times out.
	// The default behavior is to not retry on timeout.
	RetryOnTimeout bool `json:"retry_on_timeout,omitempty"`
	// An optional value specifying the condition determining whether the task
	// is run once its dependencies have been completed.
	//
	// * `ALL_SUCCESS`: All dependencies have executed and succeeded *
	// `AT_LEAST_ONE_SUCCESS`: At least one dependency has succeeded *
	// `NONE_FAILED`: None of the dependencies have failed and at least one was
	// executed * `ALL_DONE`: All dependencies have been completed *
	// `AT_LEAST_ONE_FAILED`: At least one dependency failed * `ALL_FAILED`: ALl
	// dependencies have failed
	RunIf RunIf `json:"run_if,omitempty"`
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask *RunJobTask `json:"run_job_task,omitempty"`
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
	// (Legacy) The task runs the spark-submit script when the spark_submit_task
	// field is present. Databricks recommends using the spark_jar_task instead;
	// see [Spark Submit task for jobs](/jobs/spark-submit).
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask *SqlTask `json:"sql_task,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when runs of this task
	// begin or complete. The default behavior is to not send any system
	// notifications.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Task) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Task) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TaskDependency struct {
	// Can only be specified on condition task dependencies. The outcome of the
	// dependent task that must be met for this task to run.
	Outcome string `json:"outcome,omitempty"`
	// The name of the task this task depends on.
	TaskKey string `json:"task_key"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TaskDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TaskDependency) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TaskEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded []string `json:"on_duration_warning_threshold_exceeded,omitempty"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure []string `json:"on_failure,omitempty"`
	// A list of email addresses to notify when platform-initiated maintenance
	// completes for a continuous job.
	OnMaintenanceComplete []string `json:"on_maintenance_complete,omitempty"`
	// A list of email addresses to notify when platform-initiated maintenance
	// starts for a continuous job.
	OnMaintenanceStart []string `json:"on_maintenance_start,omitempty"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []string `json:"on_start,omitempty"`
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	OnStreamingBacklogExceeded []string `json:"on_streaming_backlog_exceeded,omitempty"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess []string `json:"on_success,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TaskEmailNotifications) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TaskEmailNotifications) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TaskNotificationSettings struct {
	// If true, do not send notifications to recipients specified in `on_start`
	// for the retried runs and do not send notifications to recipients
	// specified in `on_failure` until the last retry of the run.
	AlertOnLastAttempt bool `json:"alert_on_last_attempt,omitempty"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TaskNotificationSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TaskNotificationSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// task retry mode of the continuous job * NEVER: The failed task will not be
// retried. * ON_FAILURE: Retry a failed task if at least one other task in the
// job is still running its first attempt. When this condition is no longer met
// or the retry limit is reached, the job run is cancelled and a new run is
// started.
type TaskRetryMode string

const TaskRetryModeNever TaskRetryMode = `NEVER`

const TaskRetryModeOnFailure TaskRetryMode = `ON_FAILURE`

// String representation for [fmt.Print]
func (f *TaskRetryMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TaskRetryMode) Set(v string) error {
	switch v {
	case `NEVER`, `ON_FAILURE`:
		*f = TaskRetryMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NEVER", "ON_FAILURE"`, v)
	}
}

// Values returns all possible values for TaskRetryMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *TaskRetryMode) Values() []TaskRetryMode {
	return []TaskRetryMode{
		TaskRetryModeNever,
		TaskRetryModeOnFailure,
	}
}

// Type always returns TaskRetryMode to satisfy [pflag.Value] interface
func (f *TaskRetryMode) Type() string {
	return "TaskRetryMode"
}

// The code indicates why the run was terminated. Additional codes might be
// introduced in future releases. * `SUCCESS`: The run was completed
// successfully. * `SUCCESS_WITH_FAILURES`: The run was completed successfully
// but some child runs failed. * `USER_CANCELED`: The run was successfully
// canceled during execution by a user. * `CANCELED`: The run was canceled
// during execution by the Databricks platform; for example, if the maximum run
// duration was exceeded. * `SKIPPED`: Run was never executed, for example, if
// the upstream task run failed, the dependency type condition was not met, or
// there were no material tasks to execute. * `INTERNAL_ERROR`: The run
// encountered an unexpected error. Refer to the state message for further
// details. * `DRIVER_ERROR`: The run encountered an error while communicating
// with the Spark Driver. * `CLUSTER_ERROR`: The run failed due to a cluster
// error. Refer to the state message for further details. *
// `REPOSITORY_CHECKOUT_FAILED`: Failed to complete the checkout due to an error
// when communicating with the third party service. * `INVALID_CLUSTER_REQUEST`:
// The run failed because it issued an invalid request to start the cluster. *
// `WORKSPACE_RUN_LIMIT_EXCEEDED`: The workspace has reached the quota for the
// maximum number of concurrent active runs. Consider scheduling the runs over a
// larger time frame. * `FEATURE_DISABLED`: The run failed because it tried to
// access a feature unavailable for the workspace. *
// `CLUSTER_REQUEST_LIMIT_EXCEEDED`: The number of cluster creation, start, and
// upsize requests have exceeded the allotted rate limit. Consider spreading the
// run execution over a larger time frame. * `STORAGE_ACCESS_ERROR`: The run
// failed due to an error when accessing the customer blob storage. Refer to the
// state message for further details. * `RUN_EXECUTION_ERROR`: The run was
// completed with task failures. For more details, refer to the state message or
// run output. * `UNAUTHORIZED_ERROR`: The run failed due to a permission issue
// while accessing a resource. Refer to the state message for further details. *
// `LIBRARY_INSTALLATION_ERROR`: The run failed while installing the
// user-requested library. Refer to the state message for further details. The
// causes might include, but are not limited to: The provided library is
// invalid, there are insufficient permissions to install the library, and so
// forth. * `MAX_CONCURRENT_RUNS_EXCEEDED`: The scheduled run exceeds the limit
// of maximum concurrent runs set for the job. * `MAX_SPARK_CONTEXTS_EXCEEDED`:
// The run is scheduled on a cluster that has already reached the maximum number
// of contexts it is configured to create. See: [Link]. * `RESOURCE_NOT_FOUND`:
// A resource necessary for run execution does not exist. Refer to the state
// message for further details. * `INVALID_RUN_CONFIGURATION`: The run failed
// due to an invalid configuration. Refer to the state message for further
// details. * `CLOUD_FAILURE`: The run failed due to a cloud provider issue.
// Refer to the state message for further details. *
// `MAX_JOB_QUEUE_SIZE_EXCEEDED`: The run was skipped due to reaching the job
// level queue size limit. * `DISABLED`: The run was never executed because it
// was disabled explicitly by the user. * `BREAKING_CHANGE`: Run failed because
// of an intentional breaking change in Spark, but it will be retried with a
// mitigation config. * `CLUSTER_TERMINATED_BY_USER`: The run failed because the
// externally managed cluster entered an unusable state, likely due to the user
// terminating or restarting it outside the jobs service.
//
// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
type TerminationCodeCode string

// Run failed because of an intentional breaking change in Spark, but it will be
// retried with a mitigation config.
const TerminationCodeCodeBreakingChange TerminationCodeCode = `BREAKING_CHANGE`

const TerminationCodeCodeBudgetPolicyLimitExceeded TerminationCodeCode = `BUDGET_POLICY_LIMIT_EXCEEDED`

// The run was canceled during execution by the <Databricks> platform; for
// example, if the maximum run duration was exceeded.
const TerminationCodeCodeCanceled TerminationCodeCode = `CANCELED`

// The run failed due to a cloud provider issue. Refer to the state message for
// further details.
const TerminationCodeCodeCloudFailure TerminationCodeCode = `CLOUD_FAILURE`

// The run failed due to a cluster error. Refer to the state message for further
// details.
const TerminationCodeCodeClusterError TerminationCodeCode = `CLUSTER_ERROR`

// The number of cluster creation, start, and upsize requests have exceeded the
// allotted rate limit. Consider spreading the run execution over a larger time
// frame.
const TerminationCodeCodeClusterRequestLimitExceeded TerminationCodeCode = `CLUSTER_REQUEST_LIMIT_EXCEEDED`

// The run was never executed because it was disabled explicitly by the user.
const TerminationCodeCodeDisabled TerminationCodeCode = `DISABLED`

// The run encountered an error while communicating with the Spark Driver.
const TerminationCodeCodeDriverError TerminationCodeCode = `DRIVER_ERROR`

// The run failed because it tried to access a feature unavailable for the
// workspace.
const TerminationCodeCodeFeatureDisabled TerminationCodeCode = `FEATURE_DISABLED`

// The run encountered an unexpected error. Refer to the state message for
// further details.
const TerminationCodeCodeInternalError TerminationCodeCode = `INTERNAL_ERROR`

// The run failed because it issued an invalid request to start the cluster.
const TerminationCodeCodeInvalidClusterRequest TerminationCodeCode = `INVALID_CLUSTER_REQUEST`

// The run failed due to an invalid configuration. Refer to the state message
// for further details.
const TerminationCodeCodeInvalidRunConfiguration TerminationCodeCode = `INVALID_RUN_CONFIGURATION`

// The run failed while installing the user-requested library. Refer to the
// state message for further details. The causes might include, but are not
// limited to: The provided library is invalid, there are insufficient
// permissions to install the library, and so forth.
const TerminationCodeCodeLibraryInstallationError TerminationCodeCode = `LIBRARY_INSTALLATION_ERROR`

// The scheduled run exceeds the limit of maximum concurrent runs set for the
// job.
const TerminationCodeCodeMaxConcurrentRunsExceeded TerminationCodeCode = `MAX_CONCURRENT_RUNS_EXCEEDED`

// The run was skipped due to reaching the job level queue size limit.
const TerminationCodeCodeMaxJobQueueSizeExceeded TerminationCodeCode = `MAX_JOB_QUEUE_SIZE_EXCEEDED`

// The run is scheduled on a cluster that has already reached the maximum number
// of contexts it is configured to create. See: [Link].
//
// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
const TerminationCodeCodeMaxSparkContextsExceeded TerminationCodeCode = `MAX_SPARK_CONTEXTS_EXCEEDED`

// Failed to complete the checkout due to an error when communicating with the
// third party service.
const TerminationCodeCodeRepositoryCheckoutFailed TerminationCodeCode = `REPOSITORY_CHECKOUT_FAILED`

// A resource necessary for run execution does not exist. Refer to the state
// message for further details.
const TerminationCodeCodeResourceNotFound TerminationCodeCode = `RESOURCE_NOT_FOUND`

// The run was completed with task failures. For more details, refer to the
// state message or run output.
const TerminationCodeCodeRunExecutionError TerminationCodeCode = `RUN_EXECUTION_ERROR`

// Run was never executed, for example, if the upstream task run failed, the
// dependency type condition was not met, or there were no material tasks to
// execute.
const TerminationCodeCodeSkipped TerminationCodeCode = `SKIPPED`

// The run failed due to an error when accessing the customer blob storage.
// Refer to the state message for further details.
const TerminationCodeCodeStorageAccessError TerminationCodeCode = `STORAGE_ACCESS_ERROR`

// The run was completed successfully.
const TerminationCodeCodeSuccess TerminationCodeCode = `SUCCESS`

// The run was completed successfully but some child runs failed.
const TerminationCodeCodeSuccessWithFailures TerminationCodeCode = `SUCCESS_WITH_FAILURES`

// The run failed due to a permission issue while accessing a resource. Refer to
// the state message for further details.
const TerminationCodeCodeUnauthorizedError TerminationCodeCode = `UNAUTHORIZED_ERROR`

// The run was successfully canceled during execution by a user.
const TerminationCodeCodeUserCanceled TerminationCodeCode = `USER_CANCELED`

// The workspace has reached the quota for the maximum number of concurrent
// active runs. Consider scheduling the runs over a larger time frame.
const TerminationCodeCodeWorkspaceRunLimitExceeded TerminationCodeCode = `WORKSPACE_RUN_LIMIT_EXCEEDED`

// String representation for [fmt.Print]
func (f *TerminationCodeCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationCodeCode) Set(v string) error {
	switch v {
	case `BREAKING_CHANGE`, `BUDGET_POLICY_LIMIT_EXCEEDED`, `CANCELED`, `CLOUD_FAILURE`, `CLUSTER_ERROR`, `CLUSTER_REQUEST_LIMIT_EXCEEDED`, `DISABLED`, `DRIVER_ERROR`, `FEATURE_DISABLED`, `INTERNAL_ERROR`, `INVALID_CLUSTER_REQUEST`, `INVALID_RUN_CONFIGURATION`, `LIBRARY_INSTALLATION_ERROR`, `MAX_CONCURRENT_RUNS_EXCEEDED`, `MAX_JOB_QUEUE_SIZE_EXCEEDED`, `MAX_SPARK_CONTEXTS_EXCEEDED`, `REPOSITORY_CHECKOUT_FAILED`, `RESOURCE_NOT_FOUND`, `RUN_EXECUTION_ERROR`, `SKIPPED`, `STORAGE_ACCESS_ERROR`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `UNAUTHORIZED_ERROR`, `USER_CANCELED`, `WORKSPACE_RUN_LIMIT_EXCEEDED`:
		*f = TerminationCodeCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BREAKING_CHANGE", "BUDGET_POLICY_LIMIT_EXCEEDED", "CANCELED", "CLOUD_FAILURE", "CLUSTER_ERROR", "CLUSTER_REQUEST_LIMIT_EXCEEDED", "DISABLED", "DRIVER_ERROR", "FEATURE_DISABLED", "INTERNAL_ERROR", "INVALID_CLUSTER_REQUEST", "INVALID_RUN_CONFIGURATION", "LIBRARY_INSTALLATION_ERROR", "MAX_CONCURRENT_RUNS_EXCEEDED", "MAX_JOB_QUEUE_SIZE_EXCEEDED", "MAX_SPARK_CONTEXTS_EXCEEDED", "REPOSITORY_CHECKOUT_FAILED", "RESOURCE_NOT_FOUND", "RUN_EXECUTION_ERROR", "SKIPPED", "STORAGE_ACCESS_ERROR", "SUCCESS", "SUCCESS_WITH_FAILURES", "UNAUTHORIZED_ERROR", "USER_CANCELED", "WORKSPACE_RUN_LIMIT_EXCEEDED"`, v)
	}
}

// Values returns all possible values for TerminationCodeCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *TerminationCodeCode) Values() []TerminationCodeCode {
	return []TerminationCodeCode{
		TerminationCodeCodeBreakingChange,
		TerminationCodeCodeBudgetPolicyLimitExceeded,
		TerminationCodeCodeCanceled,
		TerminationCodeCodeCloudFailure,
		TerminationCodeCodeClusterError,
		TerminationCodeCodeClusterRequestLimitExceeded,
		TerminationCodeCodeDisabled,
		TerminationCodeCodeDriverError,
		TerminationCodeCodeFeatureDisabled,
		TerminationCodeCodeInternalError,
		TerminationCodeCodeInvalidClusterRequest,
		TerminationCodeCodeInvalidRunConfiguration,
		TerminationCodeCodeLibraryInstallationError,
		TerminationCodeCodeMaxConcurrentRunsExceeded,
		TerminationCodeCodeMaxJobQueueSizeExceeded,
		TerminationCodeCodeMaxSparkContextsExceeded,
		TerminationCodeCodeRepositoryCheckoutFailed,
		TerminationCodeCodeResourceNotFound,
		TerminationCodeCodeRunExecutionError,
		TerminationCodeCodeSkipped,
		TerminationCodeCodeStorageAccessError,
		TerminationCodeCodeSuccess,
		TerminationCodeCodeSuccessWithFailures,
		TerminationCodeCodeUnauthorizedError,
		TerminationCodeCodeUserCanceled,
		TerminationCodeCodeWorkspaceRunLimitExceeded,
	}
}

// Type always returns TerminationCodeCode to satisfy [pflag.Value] interface
func (f *TerminationCodeCode) Type() string {
	return "TerminationCodeCode"
}

type TerminationDetails struct {
	Code TerminationCodeCode `json:"code,omitempty"`
	// A descriptive message with the termination details. This field is
	// unstructured and the format might change.
	Message string `json:"message,omitempty"`

	Type TerminationTypeType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TerminationDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TerminationDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// * `SUCCESS`: The run terminated without any issues * `INTERNAL_ERROR`: An
// error occurred in the Databricks platform. Please look at the [status page]
// or contact support if the issue persists. * `CLIENT_ERROR`: The run was
// terminated because of an error caused by user input or the job configuration.
// * `CLOUD_FAILURE`: The run was terminated because of an issue with your cloud
// provider.
//
// [status page]: https://status.databricks.com/
type TerminationTypeType string

// The run was terminated because of an error caused by user input or the job
// configuration.
const TerminationTypeTypeClientError TerminationTypeType = `CLIENT_ERROR`

// The run was terminated because of an issue with your cloud provider.
const TerminationTypeTypeCloudFailure TerminationTypeType = `CLOUD_FAILURE`

// An error occurred in the <Databricks> platform. Please look at the [status
// page] or contact support if the issue persists.
//
// [status page]: https://status.databricks.com/
const TerminationTypeTypeInternalError TerminationTypeType = `INTERNAL_ERROR`

// The run terminated without any issues
const TerminationTypeTypeSuccess TerminationTypeType = `SUCCESS`

// String representation for [fmt.Print]
func (f *TerminationTypeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationTypeType) Set(v string) error {
	switch v {
	case `CLIENT_ERROR`, `CLOUD_FAILURE`, `INTERNAL_ERROR`, `SUCCESS`:
		*f = TerminationTypeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLIENT_ERROR", "CLOUD_FAILURE", "INTERNAL_ERROR", "SUCCESS"`, v)
	}
}

// Values returns all possible values for TerminationTypeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TerminationTypeType) Values() []TerminationTypeType {
	return []TerminationTypeType{
		TerminationTypeTypeClientError,
		TerminationTypeTypeCloudFailure,
		TerminationTypeTypeInternalError,
		TerminationTypeTypeSuccess,
	}
}

// Type always returns TerminationTypeType to satisfy [pflag.Value] interface
func (f *TerminationTypeType) Type() string {
	return "TerminationTypeType"
}

// A single trigger attached to a job via `JobSettings.triggers`. Exactly one of
// the trigger-type fields (`periodic`, `schedule`, `continuous`,
// `file_arrival`, `table_update`, `model`, `job_completion`) must be set;
// mutual exclusivity is enforced in the API handler rather than via `oneof` so
// that codegen, validation, and JSON serialization across SDKs and Terraform
// behave consistently.
type TriggerConfiguration struct {
	// Continuous trigger configuration.
	Continuous *ContinuousTriggerConfiguration `json:"continuous,omitempty"`
	// File arrival trigger configuration.
	FileArrival *FileArrivalTriggerConfiguration `json:"file_arrival,omitempty"`
	// Job completion trigger configuration.
	JobCompletion *JobCompletionTriggerConfiguration `json:"job_completion,omitempty"`
	// Model trigger configuration.
	Model *ModelTriggerConfiguration `json:"model,omitempty"`
	// Whether this trigger is paused. Defaults to UNPAUSED when unset; the
	// server always returns an explicit value on read.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
	// Trigger type: exactly one must be set; mutual exclusivity is enforced in
	// the API handler Periodic trigger configuration.
	Periodic *PeriodicTriggerConfiguration `json:"periodic,omitempty"`
	// Cron schedule trigger configuration.
	Schedule *CronTriggerConfiguration `json:"schedule,omitempty"`
	// Optional SQL condition that gates whether this trigger fires.
	SqlCondition *SqlConditionConfiguration `json:"sql_condition,omitempty"`
	// Table update trigger configuration.
	TableUpdate *TableUpdateTriggerConfiguration `json:"table_update,omitempty"`
}

func (s *TriggerConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Per-trigger runtime details returned by `GetJob`. Same length and order as
// `JobSettings.triggers`; sub-fields are populated independently based on the
// corresponding `GetJob.include_trigger_state` / `include_trigger_history`
// flags.
type TriggerDetails struct {
	// Recent evaluation history. Populated when
	// `GetJob.include_trigger_history` is set.
	History *TriggerHistory `json:"history,omitempty"`
	// Current runtime state. Populated when `GetJob.include_trigger_state` is
	// set.
	State *PerTriggerState `json:"state,omitempty"`
}

func (s *TriggerDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TriggerEvaluation struct {
	// Human-readable description of the trigger evaluation result. Explains why
	// the trigger evaluation triggered or did not trigger a run, or failed.
	Description string `json:"description,omitempty"`
	// The ID of the run that was triggered by the trigger evaluation. Only
	// returned if a run was triggered.
	RunId int64 `json:"run_id,omitempty"`
	// Timestamp at which the trigger was evaluated.
	Timestamp int64 `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TriggerEvaluation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TriggerEvaluation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TriggerHistory struct {
	// The last time the trigger failed to evaluate.
	LastFailed *TriggerEvaluation `json:"last_failed,omitempty"`
	// The last time the trigger was evaluated but did not trigger a run.
	LastNotTriggered *TriggerEvaluation `json:"last_not_triggered,omitempty"`
	// The last time the run was triggered due to a file arrival.
	LastTriggered *TriggerEvaluation `json:"last_triggered,omitempty"`
}

func (s *TriggerHistory) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Additional details about what triggered the run
type TriggerInfo struct {
	// Additional details for job runs triggered by a job completion trigger
	// (trigger: "JOB_COMPLETION")
	JobCompletion *JobCompletionTriggerRunInfoDetails `json:"job_completion,omitempty"`
	// The run id of the Run Job task run
	RunId int64 `json:"run_id,omitempty"`
	// SQL condition evaluation details for this run
	SqlCondition *SqlConditionRunInfoDetails `json:"sql_condition,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TriggerInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TriggerInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TriggerSettings struct {
	// File arrival trigger settings.
	FileArrival *FileArrivalTriggerConfiguration `json:"file_arrival,omitempty"`
	// Job completion trigger settings.
	JobCompletion *JobCompletionTriggerConfiguration `json:"job_completion,omitempty"`

	Model *ModelTriggerConfiguration `json:"model,omitempty"`
	// Whether this trigger is paused or not.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
	// Periodic trigger settings.
	Periodic *PeriodicTriggerConfiguration `json:"periodic,omitempty"`
	// SQL condition that must be satisfied for the trigger to fire. Can be used
	// in combination with other trigger types and runs *after* other trigger
	// types conditions are evaluated.
	SqlCondition *SqlConditionConfiguration `json:"sql_condition,omitempty"`
	// Table update trigger settings.
	TableUpdate *TableUpdateTriggerConfiguration `json:"table_update,omitempty"`
}

func (s *TriggerSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TriggerStateProto struct {
	// State of the file arrival trigger.
	FileArrival *FileArrivalTriggerState `json:"file_arrival,omitempty"`
	// State of the job completion trigger.
	JobCompletion *JobCompletionTriggerState `json:"job_completion,omitempty"`
	// Whether this trigger is paused or not. For continuous schedules, it can
	// differ from the configured pause_status whenever a paused continuous job
	// is kickstarted by an operation other than an update, such as a run-now.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
	// State for SQL condition evaluation, can coexist with other trigger
	// states.
	SqlCondition *SqlConditionState `json:"sql_condition,omitempty"`
	// State of the table update trigger.
	Table *TableTriggerState `json:"table,omitempty"`
}

func (s *TriggerStateProto) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The type of trigger that fired this run.
//
// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
// occurs you triggered a single run on demand through the UI or the API. *
// `RETRY`: Indicates a run that is triggered as a retry of a previously failed
// run. This occurs when you request to re-run the job in case of failures. *
// `RUN_JOB_TASK`: Indicates a run that is triggered using a Run Job task. *
// `FILE_ARRIVAL`: Indicates a run that is triggered by a file arrival. *
// `CONTINUOUS`: Indicates a run that is triggered by a continuous job. *
// `TABLE`: Indicates a run that is triggered by a table update. *
// `CONTINUOUS_RESTART`: Indicates a run created by user to manually restart a
// continuous job run. * `MODEL`: Indicates a run that is triggered by a model
// update. * `JOB_COMPLETION`: Indicates a run that is triggered by an upstream
// job completion.
type TriggerType string

// Indicates a run that is triggered by a continuous job.
const TriggerTypeContinuous TriggerType = `CONTINUOUS`

// Indicates a run created by user to manually restart a continuous job run.
const TriggerTypeContinuousRestart TriggerType = `CONTINUOUS_RESTART`

// Indicates a run that is triggered by a file arrival.
const TriggerTypeFileArrival TriggerType = `FILE_ARRIVAL`

// Indicates a run that is triggered by an upstream job completion.
const TriggerTypeJobCompletion TriggerType = `JOB_COMPLETION`

// One time triggers that fire a single run. This occurs you triggered a single
// run on demand through the UI or the API.
const TriggerTypeOneTime TriggerType = `ONE_TIME`

// Schedules that periodically trigger runs, such as a cron scheduler.
const TriggerTypePeriodic TriggerType = `PERIODIC`

// Indicates a run that is triggered as a retry of a previously failed run. This
// occurs when you request to re-run the job in case of failures.
const TriggerTypeRetry TriggerType = `RETRY`

// Indicates a run that is triggered using a Run Job task.
const TriggerTypeRunJobTask TriggerType = `RUN_JOB_TASK`

// Indicates a run that is triggered by a table update.
const TriggerTypeTable TriggerType = `TABLE`

// String representation for [fmt.Print]
func (f *TriggerType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TriggerType) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `CONTINUOUS_RESTART`, `FILE_ARRIVAL`, `JOB_COMPLETION`, `ONE_TIME`, `PERIODIC`, `RETRY`, `RUN_JOB_TASK`, `TABLE`:
		*f = TriggerType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "CONTINUOUS_RESTART", "FILE_ARRIVAL", "JOB_COMPLETION", "ONE_TIME", "PERIODIC", "RETRY", "RUN_JOB_TASK", "TABLE"`, v)
	}
}

// Values returns all possible values for TriggerType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TriggerType) Values() []TriggerType {
	return []TriggerType{
		TriggerTypeContinuous,
		TriggerTypeContinuousRestart,
		TriggerTypeFileArrival,
		TriggerTypeJobCompletion,
		TriggerTypeOneTime,
		TriggerTypePeriodic,
		TriggerTypeRetry,
		TriggerTypeRunJobTask,
		TriggerTypeTable,
	}
}

// Type always returns TriggerType to satisfy [pflag.Value] interface
func (f *TriggerType) Type() string {
	return "TriggerType"
}

type UpdateJob struct {
	// Remove top-level fields in the job settings. Removing nested fields is
	// not supported, except for tasks and job clusters (`tasks/task_1`). This
	// field is optional.
	FieldsToRemove []string `json:"fields_to_remove,omitempty"`
	// The canonical identifier of the job to update. This field is required.
	JobId int64 `json:"job_id"`
	// The new settings for the job.
	//
	// Top-level fields specified in `new_settings` are completely replaced,
	// except for arrays which are merged. That is, new and existing entries are
	// completely replaced based on the respective key fields, i.e. `task_key`
	// or `job_cluster_key`, while previous entries are kept.
	//
	// Partially updating nested fields is not supported.
	//
	// Changes to the field `JobSettings.timeout_seconds` are applied to active
	// runs. Changes to other fields are applied to future runs only.
	NewSettings *JobSettings `json:"new_settings,omitempty"`
}

func (s *UpdateJob) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ViewItem struct {
	// Content of the view.
	Content string `json:"content,omitempty"`
	// Name of the view item. In the case of code view, it would be the
	// notebook’s name. In the case of dashboard view, it would be the
	// dashboard’s name.
	Name string `json:"name,omitempty"`
	// Type of the view item.
	Type ViewType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ViewItem) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ViewItem) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// * `NOTEBOOK`: Notebook view item. * `DASHBOARD`: Dashboard view item.
type ViewType string

// Dashboard view item.
const ViewTypeDashboard ViewType = `DASHBOARD`

// Notebook view item.
const ViewTypeNotebook ViewType = `NOTEBOOK`

// String representation for [fmt.Print]
func (f *ViewType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ViewType) Set(v string) error {
	switch v {
	case `DASHBOARD`, `NOTEBOOK`:
		*f = ViewType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DASHBOARD", "NOTEBOOK"`, v)
	}
}

// Values returns all possible values for ViewType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ViewType) Values() []ViewType {
	return []ViewType{
		ViewTypeDashboard,
		ViewTypeNotebook,
	}
}

// Type always returns ViewType to satisfy [pflag.Value] interface
func (f *ViewType) Type() string {
	return "ViewType"
}

// * `CODE`: Code view of the notebook. * `DASHBOARDS`: All dashboard views of
// the notebook. * `ALL`: All views of the notebook.
type ViewsToExport string

// All views of the notebook.
const ViewsToExportAll ViewsToExport = `ALL`

// Code view of the notebook.
const ViewsToExportCode ViewsToExport = `CODE`

// All dashboard views of the notebook.
const ViewsToExportDashboards ViewsToExport = `DASHBOARDS`

// String representation for [fmt.Print]
func (f *ViewsToExport) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ViewsToExport) Set(v string) error {
	switch v {
	case `ALL`, `CODE`, `DASHBOARDS`:
		*f = ViewsToExport(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL", "CODE", "DASHBOARDS"`, v)
	}
}

// Values returns all possible values for ViewsToExport.
//
// There is no guarantee on the order of the values in the slice.
func (f *ViewsToExport) Values() []ViewsToExport {
	return []ViewsToExport{
		ViewsToExportAll,
		ViewsToExportCode,
		ViewsToExportDashboards,
	}
}

// Type always returns ViewsToExport to satisfy [pflag.Value] interface
func (f *ViewsToExport) Type() string {
	return "ViewsToExport"
}

// A Unity Catalog volume input source. The function is applied to each file.
type VolumeSource struct {
	// Unity Catalog volume path to the input files (e.g.
	// `/Volumes/<catalog>/<schema>/<volume>/...`).
	VolumePath string `json:"volume_path"`
}

func (s *VolumeSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Webhook struct {
	// The ID of the notification destination to call. Must be a UUID.
	Id string `json:"id"`
}

func (s *Webhook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type WebhookNotifications struct {
	// An optional list of system notification IDs to call when the duration of
	// a run exceeds the threshold specified for the `RUN_DURATION_SECONDS`
	// metric in the `health` field. A maximum of 3 destinations can be
	// specified for the `on_duration_warning_threshold_exceeded` property.
	OnDurationWarningThresholdExceeded []Webhook `json:"on_duration_warning_threshold_exceeded,omitempty"`
	// An optional list of system notification IDs to call when the run fails. A
	// maximum of 3 destinations can be specified for the `on_failure` property.
	OnFailure []Webhook `json:"on_failure,omitempty"`
	// An optional list of system notification IDs to call when
	// platform-initiated maintenance completes for a continuous job. A maximum
	// of 3 destinations can be specified for the `on_maintenance_complete`
	// property.
	OnMaintenanceComplete []Webhook `json:"on_maintenance_complete,omitempty"`
	// An optional list of system notification IDs to call when
	// platform-initiated maintenance starts for a continuous job. A maximum of
	// 3 destinations can be specified for the `on_maintenance_start` property.
	OnMaintenanceStart []Webhook `json:"on_maintenance_start,omitempty"`
	// An optional list of system notification IDs to call when the run starts.
	// A maximum of 3 destinations can be specified for the `on_start` property.
	OnStart []Webhook `json:"on_start,omitempty"`
	// An optional list of system notification IDs to call when any streaming
	// backlog thresholds are exceeded for any stream. Streaming backlog
	// thresholds can be set in the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes. A maximum of 3 destinations
	// can be specified for the `on_streaming_backlog_exceeded` property.
	OnStreamingBacklogExceeded []Webhook `json:"on_streaming_backlog_exceeded,omitempty"`
	// An optional list of system notification IDs to call when the run
	// completes successfully. A maximum of 3 destinations can be specified for
	// the `on_success` property.
	OnSuccess []Webhook `json:"on_success,omitempty"`
}

func (s *WebhookNotifications) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type WidgetErrorDetail struct {
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WidgetErrorDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WidgetErrorDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
