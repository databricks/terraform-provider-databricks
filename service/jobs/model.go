// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BaseJob struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime types.Int64 `tfsdk:"created_time"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// The canonical identifier for this job.
	JobId types.Int64 `tfsdk:"job_id"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *JobSettings `tfsdk:"settings"`
}

type BaseRun struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0\. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` \> 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber types.Int64 `tfsdk:"attempt_number"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration types.Int64 `tfsdk:"cleanup_duration"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `tfsdk:"cluster_instance"`
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec *ClusterSpec `tfsdk:"cluster_spec"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// Description of the run
	Description types.String `tfsdk:"description"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime types.Int64 `tfsdk:"end_time"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration types.Int64 `tfsdk:"execution_duration"`
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
	GitSource *GitSource `tfsdk:"git_source"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `tfsdk:"job_clusters"`
	// The canonical identifier of the job that contains this run.
	JobId types.Int64 `tfsdk:"job_id"`
	// Job-level parameters used in the run
	JobParameters []JobParameter `tfsdk:"job_parameters"`
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob types.Int64 `tfsdk:"number_in_job"`
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId types.Int64 `tfsdk:"original_attempt_run_id"`
	// The parameters used for this run.
	OverridingParameters *RunParameters `tfsdk:"overriding_parameters"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration types.Int64 `tfsdk:"queue_duration"`
	// The repair history of the run.
	RepairHistory []RepairHistoryItem `tfsdk:"repair_history"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration types.Int64 `tfsdk:"run_duration"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId types.Int64 `tfsdk:"run_id"`
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName types.String `tfsdk:"run_name"`
	// The URL to the detail page of the run.
	RunPageUrl types.String `tfsdk:"run_page_url"`
	// The type of a run. * `JOB_RUN`: Normal job run. A run created with
	// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
	// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
	// :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	RunType RunType `tfsdk:"run_type"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *CronSchedule `tfsdk:"schedule"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration types.Int64 `tfsdk:"setup_duration"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime types.Int64 `tfsdk:"start_time"`
	// The current state of the run.
	State *RunState `tfsdk:"state"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls.
	Tasks []RunTask `tfsdk:"tasks"`
	// The type of trigger that fired this run.
	//
	// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
	// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
	// occurs you triggered a single run on demand through the UI or the API. *
	// `RETRY`: Indicates a run that is triggered as a retry of a previously
	// failed run. This occurs when you request to re-run the job in case of
	// failures. * `RUN_JOB_TASK`: Indicates a run that is triggered using a Run
	// Job task. * `FILE_ARRIVAL`: Indicates a run that is triggered by a file
	// arrival. * `TABLE`: Indicates a run that is triggered by a table update.
	Trigger TriggerType `tfsdk:"trigger"`
	// Additional details about what triggered the run
	TriggerInfo *TriggerInfo `tfsdk:"trigger_info"`
}

type CancelAllRuns struct {
	// Optional boolean parameter to cancel all queued runs. If no job_id is
	// provided, all queued runs in the workspace are canceled.
	AllQueuedRuns types.Bool `tfsdk:"all_queued_runs"`
	// The canonical identifier of the job to cancel all runs of.
	JobId types.Int64 `tfsdk:"job_id"`
}

type CancelAllRunsResponse struct {
}

type CancelRun struct {
	// This field is required.
	RunId types.Int64 `tfsdk:"run_id"`
}

type CancelRunResponse struct {
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
	ClusterId types.String `tfsdk:"cluster_id"`
	// The canonical identifier for the Spark context used by a run. This field
	// is filled in once the run begins execution. This value can be used to
	// view the Spark UI by browsing to
	// `/#setting/sparkui/$cluster_id/$spark_context_id`. The Spark UI continues
	// to be available after the run has completed.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	SparkContextId types.String `tfsdk:"spark_context_id"`
}

type ClusterSpec struct {
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId types.String `tfsdk:"existing_cluster_id"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey types.String `tfsdk:"job_cluster_key"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library `tfsdk:"libraries"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec `tfsdk:"new_cluster"`
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

// Type always returns Condition to satisfy [pflag.Value] interface
func (f *Condition) Type() string {
	return "Condition"
}

type ConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	Left types.String `tfsdk:"left"`
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
	Op ConditionTaskOp `tfsdk:"op"`
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right types.String `tfsdk:"right"`
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

// Type always returns ConditionTaskOp to satisfy [pflag.Value] interface
func (f *ConditionTaskOp) Type() string {
	return "ConditionTaskOp"
}

type Continuous struct {
	// Indicate whether the continuous execution of the job is paused or not.
	// Defaults to UNPAUSED.
	PauseStatus PauseStatus `tfsdk:"pause_status"`
}

type CreateJob struct {
	// List of permissions to set on the job.
	AccessControlList []iam.AccessControlRequest `tfsdk:"access_control_list"`
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous *Continuous `tfsdk:"continuous"`
	// Deployment information for jobs managed by external sources.
	Deployment *JobDeployment `tfsdk:"deployment"`
	// An optional description for the job. The maximum length is 1024
	// characters in UTF-8 encoding.
	Description types.String `tfsdk:"description"`
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode JobEditMode `tfsdk:"edit_mode"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications *JobEmailNotifications `tfsdk:"email_notifications"`
	// A list of task execution environment specifications that can be
	// referenced by tasks of this job.
	Environments []JobEnvironment `tfsdk:"environments"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format Format `tfsdk:"format"`
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
	GitSource *GitSource `tfsdk:"git_source"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `tfsdk:"health"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `tfsdk:"job_clusters"`
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
	MaxConcurrentRuns types.Int64 `tfsdk:"max_concurrent_runs"`
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name types.String `tfsdk:"name"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *JobNotificationSettings `tfsdk:"notification_settings"`
	// Job-level parameter definitions
	Parameters []JobParameterDefinition `tfsdk:"parameters"`
	// The queue settings of the job.
	Queue *QueueSettings `tfsdk:"queue"`
	// Write-only setting, available only in Create/Update/Reset and Submit
	// calls. Specifies the user or service principal that the job runs as. If
	// not specified, the job runs as the user who created the job.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	RunAs *JobRunAs `tfsdk:"run_as"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule *CronSchedule `tfsdk:"schedule"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags map[types.String]types.String `tfsdk:"tags"`
	// A list of task specifications to be executed by this job.
	Tasks []Task `tfsdk:"tasks"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds"`
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger *TriggerSettings `tfsdk:"trigger"`
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications *WebhookNotifications `tfsdk:"webhook_notifications"`
}

// Job was created successfully
type CreateResponse struct {
	// The canonical identifier for the newly created job.
	JobId types.Int64 `tfsdk:"job_id"`
}

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	PauseStatus PauseStatus `tfsdk:"pause_status"`
	// A Cron expression using Quartz syntax that describes the schedule for a
	// job. See [Cron Trigger] for details. This field is required.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression types.String `tfsdk:"quartz_cron_expression"`
	// A Java timezone ID. The schedule for a job is resolved with respect to
	// this timezone. See [Java TimeZone] for details. This field is required.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId types.String `tfsdk:"timezone_id"`
}

type DbtOutput struct {
	// An optional map of headers to send when retrieving the artifact from the
	// `artifacts_link`.
	ArtifactsHeaders map[types.String]types.String `tfsdk:"artifacts_headers"`
	// A pre-signed URL to download the (compressed) dbt artifacts. This link is
	// valid for a limited time (30 minutes). This information is only available
	// after the run has finished.
	ArtifactsLink types.String `tfsdk:"artifacts_link"`
}

type DbtTask struct {
	// Optional name of the catalog to use. The value is the top level in the
	// 3-level namespace of Unity Catalog (catalog / schema / relation). The
	// catalog value can only be specified if a warehouse_id is specified.
	// Requires dbt-databricks >= 1.1.1.
	Catalog types.String `tfsdk:"catalog"`
	// A list of dbt commands to execute. All commands must start with `dbt`.
	// This parameter must not be empty. A maximum of up to 10 commands can be
	// provided.
	Commands []types.String `tfsdk:"commands"`
	// Optional (relative) path to the profiles directory. Can only be specified
	// if no warehouse_id is specified. If no warehouse_id is specified and this
	// folder is unset, the root directory is used.
	ProfilesDirectory types.String `tfsdk:"profiles_directory"`
	// Path to the project directory. Optional for Git sourced tasks, in which
	// case if no value is provided, the root of the Git repository is used.
	ProjectDirectory types.String `tfsdk:"project_directory"`
	// Optional schema to write to. This parameter is only used when a
	// warehouse_id is also provided. If not provided, the `default` schema is
	// used.
	Schema types.String `tfsdk:"schema"`
	// Optional location type of the project directory. When set to `WORKSPACE`,
	// the project will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the project will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: Project is located in Databricks workspace. * `GIT`:
	// Project is located in cloud Git provider.
	Source Source `tfsdk:"source"`
	// ID of the SQL warehouse to connect to. If provided, we automatically
	// generate and provide the profile and connection details to dbt. It can be
	// overridden on a per-command basis by using the `--profiles-dir` command
	// line argument.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

type DeleteJob struct {
	// The canonical identifier of the job to delete. This field is required.
	JobId types.Int64 `tfsdk:"job_id"`
}

type DeleteResponse struct {
}

type DeleteRun struct {
	// ID of the run to delete.
	RunId types.Int64 `tfsdk:"run_id"`
}

type DeleteRunResponse struct {
}

// Run was exported successfully.
type ExportRunOutput struct {
	// The exported content in HTML format (one for every view item). To extract
	// the HTML notebook from the JSON response, download and run this [Python
	// script].
	//
	// [Python script]: https://docs.databricks.com/en/_static/examples/extract.py
	Views []ViewItem `tfsdk:"views"`
}

// Export and retrieve a job run
type ExportRunRequest struct {
	// The canonical identifier for the run. This field is required.
	RunId types.Int64 `tfsdk:"-" url:"run_id"`
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	ViewsToExport ViewsToExport `tfsdk:"-" url:"views_to_export,omitempty"`
}

type FileArrivalTriggerConfiguration struct {
	// If set, the trigger starts a run only after the specified amount of time
	// passed since the last time the trigger fired. The minimum allowed value
	// is 60 seconds
	MinTimeBetweenTriggersSeconds types.Int64 `tfsdk:"min_time_between_triggers_seconds"`
	// URL to be monitored for file arrivals. The path must point to the root or
	// a subpath of the external location.
	Url types.String `tfsdk:"url"`
	// If set, the trigger starts a run only after no file activity has occurred
	// for the specified amount of time. This makes it possible to wait for a
	// batch of incoming files to arrive before triggering a run. The minimum
	// allowed value is 60 seconds.
	WaitAfterLastChangeSeconds types.Int64 `tfsdk:"wait_after_last_change_seconds"`
}

type ForEachStats struct {
	// Sample of 3 most common error messages occurred during the iteration.
	ErrorMessageStats []ForEachTaskErrorMessageStats `tfsdk:"error_message_stats"`
	// Describes stats of the iteration. Only latest retries are considered.
	TaskRunStats *ForEachTaskTaskRunStats `tfsdk:"task_run_stats"`
}

type ForEachTask struct {
	// Controls the number of active iterations task runs. Default is 20,
	// maximum allowed is 100.
	Concurrency types.Int64 `tfsdk:"concurrency"`
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs types.String `tfsdk:"inputs"`
	// Configuration for the task that will be run for each element in the array
	Task Task `tfsdk:"task"`
}

type ForEachTaskErrorMessageStats struct {
	// Describes the count of such error message encountered during the
	// iterations.
	Count types.Int64 `tfsdk:"count"`
	// Describes the error message occured during the iterations.
	ErrorMessage types.String `tfsdk:"error_message"`
	// Describes the termination reason for the error message.
	TerminationCategory types.String `tfsdk:"termination_category"`
}

type ForEachTaskTaskRunStats struct {
	// Describes the iteration runs having an active lifecycle state or an
	// active run sub state.
	ActiveIterations types.Int64 `tfsdk:"active_iterations"`
	// Describes the number of failed and succeeded iteration runs.
	CompletedIterations types.Int64 `tfsdk:"completed_iterations"`
	// Describes the number of failed iteration runs.
	FailedIterations types.Int64 `tfsdk:"failed_iterations"`
	// Describes the number of iteration runs that have been scheduled.
	ScheduledIterations types.Int64 `tfsdk:"scheduled_iterations"`
	// Describes the number of succeeded iteration runs.
	SucceededIterations types.Int64 `tfsdk:"succeeded_iterations"`
	// Describes the length of the list of items to iterate over.
	TotalIterations types.Int64 `tfsdk:"total_iterations"`
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

// Type always returns Format to satisfy [pflag.Value] interface
func (f *Format) Type() string {
	return "Format"
}

// Get job permission levels
type GetJobPermissionLevelsRequest struct {
	// The job for which to get or manage permissions.
	JobId types.String `tfsdk:"-" url:"-"`
}

type GetJobPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []JobPermissionsDescription `tfsdk:"permission_levels"`
}

// Get job permissions
type GetJobPermissionsRequest struct {
	// The job for which to get or manage permissions.
	JobId types.String `tfsdk:"-" url:"-"`
}

// Get a single job
type GetJobRequest struct {
	// The canonical identifier of the job to retrieve information about. This
	// field is required.
	JobId types.Int64 `tfsdk:"-" url:"job_id"`
}

// Get the output for a single run
type GetRunOutputRequest struct {
	// The canonical identifier for the run.
	RunId types.Int64 `tfsdk:"-" url:"run_id"`
}

// Get a single job run
type GetRunRequest struct {
	// Whether to include the repair history in the response.
	IncludeHistory types.Bool `tfsdk:"-" url:"include_history,omitempty"`
	// Whether to include resolved parameter values in the response.
	IncludeResolvedValues types.Bool `tfsdk:"-" url:"include_resolved_values,omitempty"`
	// The canonical identifier of the run for which to retrieve the metadata.
	// This field is required.
	RunId types.Int64 `tfsdk:"-" url:"run_id"`
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
	UsedCommit types.String `tfsdk:"used_commit"`
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
	GitBranch types.String `tfsdk:"git_branch"`
	// Commit to be checked out and used by this job. This field cannot be
	// specified in conjunction with git_branch or git_tag.
	GitCommit types.String `tfsdk:"git_commit"`
	// Unique identifier of the service used to host the Git repository. The
	// value is case insensitive.
	GitProvider GitProvider `tfsdk:"git_provider"`
	// Read-only state of the remote repository at the time the job was run.
	// This field is only included on job runs.
	GitSnapshot *GitSnapshot `tfsdk:"git_snapshot"`
	// Name of the tag to be checked out and used by this job. This field cannot
	// be specified in conjunction with git_branch or git_commit.
	GitTag types.String `tfsdk:"git_tag"`
	// URL of the repository to be cloned by this job.
	GitUrl types.String `tfsdk:"git_url"`
	// The source of the job specification in the remote repository when the job
	// is source controlled.
	JobSource *JobSource `tfsdk:"job_source"`
}

// Job was retrieved successfully.
type Job struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime types.Int64 `tfsdk:"created_time"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// The canonical identifier for this job.
	JobId types.Int64 `tfsdk:"job_id"`
	// The email of an active workspace user or the application ID of a service
	// principal that the job runs as. This value can be changed by setting the
	// `run_as` field when creating or updating a job.
	//
	// By default, `run_as_user_name` is based on the current job settings and
	// is set to the creator of the job if job access control is disabled or to
	// the user with the `is_owner` permission if job access control is enabled.
	RunAsUserName types.String `tfsdk:"run_as_user_name"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *JobSettings `tfsdk:"settings"`
}

type JobAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel JobPermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

type JobAccessControlResponse struct {
	// All permissions.
	AllPermissions []JobPermission `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

type JobCluster struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	JobClusterKey types.String `tfsdk:"job_cluster_key"`
	// If new_cluster, a description of a cluster that is created for each task.
	NewCluster compute.ClusterSpec `tfsdk:"new_cluster"`
}

type JobDeployment struct {
	// The kind of deployment that manages the job.
	//
	// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
	Kind JobDeploymentKind `tfsdk:"kind"`
	// Path of the file that contains deployment metadata.
	MetadataFilePath types.String `tfsdk:"metadata_file_path"`
}

// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
type JobDeploymentKind string

// The job is managed by Databricks Asset Bundle.
const JobDeploymentKindBundle JobDeploymentKind = `BUNDLE`

// String representation for [fmt.Print]
func (f *JobDeploymentKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobDeploymentKind) Set(v string) error {
	switch v {
	case `BUNDLE`:
		*f = JobDeploymentKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BUNDLE"`, v)
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

// Type always returns JobEditMode to satisfy [pflag.Value] interface
func (f *JobEditMode) Type() string {
	return "JobEditMode"
}

type JobEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped.
	NoAlertForSkippedRuns types.Bool `tfsdk:"no_alert_for_skipped_runs"`
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded []types.String `tfsdk:"on_duration_warning_threshold_exceeded"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure []types.String `tfsdk:"on_failure"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []types.String `tfsdk:"on_start"`
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	OnStreamingBacklogExceeded []types.String `tfsdk:"on_streaming_backlog_exceeded"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess []types.String `tfsdk:"on_success"`
}

type JobEnvironment struct {
	// The key of an environment. It has to be unique within a job.
	EnvironmentKey types.String `tfsdk:"environment_key"`
	// The environment entity used to preserve serverless environment side panel
	// and jobs' environment for non-notebook task. In this minimal environment
	// spec, only pip dependencies are supported.
	Spec *compute.Environment `tfsdk:"spec"`
}

type JobNotificationSettings struct {
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns types.Bool `tfsdk:"no_alert_for_canceled_runs"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns types.Bool `tfsdk:"no_alert_for_skipped_runs"`
}

type JobParameter struct {
	// The optional default value of the parameter
	Default types.String `tfsdk:"default"`
	// The name of the parameter
	Name types.String `tfsdk:"name"`
	// The value used in the run
	Value types.String `tfsdk:"value"`
}

type JobParameterDefinition struct {
	// Default value of the parameter.
	Default types.String `tfsdk:"default"`
	// The name of the defined parameter. May only contain alphanumeric
	// characters, `_`, `-`, and `.`
	Name types.String `tfsdk:"name"`
}

type JobPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel JobPermissionLevel `tfsdk:"permission_level"`
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

// Type always returns JobPermissionLevel to satisfy [pflag.Value] interface
func (f *JobPermissionLevel) Type() string {
	return "JobPermissionLevel"
}

type JobPermissions struct {
	AccessControlList []JobAccessControlResponse `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

type JobPermissionsDescription struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel JobPermissionLevel `tfsdk:"permission_level"`
}

type JobPermissionsRequest struct {
	AccessControlList []JobAccessControlRequest `tfsdk:"access_control_list"`
	// The job for which to get or manage permissions.
	JobId types.String `tfsdk:"-" url:"-"`
}

// Write-only setting, available only in Create/Update/Reset and Submit calls.
// Specifies the user or service principal that the job runs as. If not
// specified, the job runs as the user who created the job.
//
// Only `user_name` or `service_principal_name` can be specified. If both are
// specified, an error is thrown.
type JobRunAs struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// The email of an active workspace user. Non-admin users can only set this
	// field to their own email.
	UserName types.String `tfsdk:"user_name"`
}

type JobSettings struct {
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous *Continuous `tfsdk:"continuous"`
	// Deployment information for jobs managed by external sources.
	Deployment *JobDeployment `tfsdk:"deployment"`
	// An optional description for the job. The maximum length is 1024
	// characters in UTF-8 encoding.
	Description types.String `tfsdk:"description"`
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode JobEditMode `tfsdk:"edit_mode"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications *JobEmailNotifications `tfsdk:"email_notifications"`
	// A list of task execution environment specifications that can be
	// referenced by tasks of this job.
	Environments []JobEnvironment `tfsdk:"environments"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format Format `tfsdk:"format"`
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
	GitSource *GitSource `tfsdk:"git_source"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `tfsdk:"health"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `tfsdk:"job_clusters"`
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
	MaxConcurrentRuns types.Int64 `tfsdk:"max_concurrent_runs"`
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name types.String `tfsdk:"name"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *JobNotificationSettings `tfsdk:"notification_settings"`
	// Job-level parameter definitions
	Parameters []JobParameterDefinition `tfsdk:"parameters"`
	// The queue settings of the job.
	Queue *QueueSettings `tfsdk:"queue"`
	// Write-only setting, available only in Create/Update/Reset and Submit
	// calls. Specifies the user or service principal that the job runs as. If
	// not specified, the job runs as the user who created the job.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	RunAs *JobRunAs `tfsdk:"run_as"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule *CronSchedule `tfsdk:"schedule"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags map[types.String]types.String `tfsdk:"tags"`
	// A list of task specifications to be executed by this job.
	Tasks []Task `tfsdk:"tasks"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds"`
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger *TriggerSettings `tfsdk:"trigger"`
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications *WebhookNotifications `tfsdk:"webhook_notifications"`
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
	DirtyState JobSourceDirtyState `tfsdk:"dirty_state"`
	// Name of the branch which the job is imported from.
	ImportFromGitBranch types.String `tfsdk:"import_from_git_branch"`
	// Path of the job YAML file that contains the job specification.
	JobConfigPath types.String `tfsdk:"job_config_path"`
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

// Type always returns JobSourceDirtyState to satisfy [pflag.Value] interface
func (f *JobSourceDirtyState) Type() string {
	return "JobSourceDirtyState"
}

// Specifies the health metric that is being evaluated for a particular health
// rule.
//
// * `RUN_DURATION_SECONDS`: Expected total time for a run in seconds. *
// `STREAMING_BACKLOG_BYTES`: An estimate of the maximum bytes of data waiting
// to be consumed across all streams. This metric is in Private Preview. *
// `STREAMING_BACKLOG_RECORDS`: An estimate of the maximum offset lag across all
// streams. This metric is in Private Preview. * `STREAMING_BACKLOG_SECONDS`: An
// estimate of the maximum consumer delay across all streams. This metric is in
// Private Preview. * `STREAMING_BACKLOG_FILES`: An estimate of the maximum
// number of outstanding files across all streams. This metric is in Private
// Preview.
type JobsHealthMetric string

// Expected total time for a run in seconds.
const JobsHealthMetricRunDurationSeconds JobsHealthMetric = `RUN_DURATION_SECONDS`

// An estimate of the maximum bytes of data waiting to be consumed across all
// streams. This metric is in Private Preview.
const JobsHealthMetricStreamingBacklogBytes JobsHealthMetric = `STREAMING_BACKLOG_BYTES`

// An estimate of the maximum number of outstanding files across all streams.
// This metric is in Private Preview.
const JobsHealthMetricStreamingBacklogFiles JobsHealthMetric = `STREAMING_BACKLOG_FILES`

// An estimate of the maximum offset lag across all streams. This metric is in
// Private Preview.
const JobsHealthMetricStreamingBacklogRecords JobsHealthMetric = `STREAMING_BACKLOG_RECORDS`

// An estimate of the maximum consumer delay across all streams. This metric is
// in Private Preview.
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
	// waiting to be consumed across all streams. This metric is in Private
	// Preview. * `STREAMING_BACKLOG_RECORDS`: An estimate of the maximum offset
	// lag across all streams. This metric is in Private Preview. *
	// `STREAMING_BACKLOG_SECONDS`: An estimate of the maximum consumer delay
	// across all streams. This metric is in Private Preview. *
	// `STREAMING_BACKLOG_FILES`: An estimate of the maximum number of
	// outstanding files across all streams. This metric is in Private Preview.
	Metric JobsHealthMetric `tfsdk:"metric"`
	// Specifies the operator used to compare the health metric value with the
	// specified threshold.
	Op JobsHealthOperator `tfsdk:"op"`
	// Specifies the threshold value that the health metric should obey to
	// satisfy the health rule.
	Value types.Int64 `tfsdk:"value"`
}

// An optional set of health rules that can be defined for this job.
type JobsHealthRules struct {
	Rules []JobsHealthRule `tfsdk:"rules"`
}

// List jobs
type ListJobsRequest struct {
	// Whether to include task and cluster details in the response.
	ExpandTasks types.Bool `tfsdk:"-" url:"expand_tasks,omitempty"`
	// The number of jobs to return. This value must be greater than 0 and less
	// or equal to 100. The default value is 20.
	Limit types.Int64 `tfsdk:"-" url:"limit,omitempty"`
	// A filter on the list based on the exact (case insensitive) job name.
	Name types.String `tfsdk:"-" url:"name,omitempty"`
	// The offset of the first job to return, relative to the most recently
	// created job. Deprecated since June 2023. Use `page_token` to iterate
	// through the pages instead.
	Offset types.Int64 `tfsdk:"-" url:"offset,omitempty"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of jobs respectively.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

// List of jobs was retrieved successfully.
type ListJobsResponse struct {
	// If true, additional jobs matching the provided filter are available for
	// listing.
	HasMore types.Bool `tfsdk:"has_more"`
	// The list of jobs. Only included in the response if there are jobs to
	// list.
	Jobs []BaseJob `tfsdk:"jobs"`
	// A token that can be used to list the next page of jobs (if applicable).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// A token that can be used to list the previous page of jobs (if
	// applicable).
	PrevPageToken types.String `tfsdk:"prev_page_token"`
}

// List job runs
type ListRunsRequest struct {
	// If active_only is `true`, only active runs are included in the results;
	// otherwise, lists both active and completed runs. An active run is a run
	// in the `QUEUED`, `PENDING`, `RUNNING`, or `TERMINATING`. This field
	// cannot be `true` when completed_only is `true`.
	ActiveOnly types.Bool `tfsdk:"-" url:"active_only,omitempty"`
	// If completed_only is `true`, only completed runs are included in the
	// results; otherwise, lists both active and completed runs. This field
	// cannot be `true` when active_only is `true`.
	CompletedOnly types.Bool `tfsdk:"-" url:"completed_only,omitempty"`
	// Whether to include task and cluster details in the response.
	ExpandTasks types.Bool `tfsdk:"-" url:"expand_tasks,omitempty"`
	// The job for which to list runs. If omitted, the Jobs service lists runs
	// from all jobs.
	JobId types.Int64 `tfsdk:"-" url:"job_id,omitempty"`
	// The number of runs to return. This value must be greater than 0 and less
	// than 25. The default value is 20. If a request specifies a limit of 0,
	// the service instead uses the maximum limit.
	Limit types.Int64 `tfsdk:"-" url:"limit,omitempty"`
	// The offset of the first run to return, relative to the most recent run.
	// Deprecated since June 2023. Use `page_token` to iterate through the pages
	// instead.
	Offset types.Int64 `tfsdk:"-" url:"offset,omitempty"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of runs respectively.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
	// The type of runs to return. For a description of run types, see
	// :method:jobs/getRun.
	RunType RunType `tfsdk:"-" url:"run_type,omitempty"`
	// Show runs that started _at or after_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_to_ to filter
	// by a time range.
	StartTimeFrom types.Int64 `tfsdk:"-" url:"start_time_from,omitempty"`
	// Show runs that started _at or before_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_from_ to
	// filter by a time range.
	StartTimeTo types.Int64 `tfsdk:"-" url:"start_time_to,omitempty"`
}

// List of runs was retrieved successfully.
type ListRunsResponse struct {
	// If true, additional runs matching the provided filter are available for
	// listing.
	HasMore types.Bool `tfsdk:"has_more"`
	// A token that can be used to list the next page of runs (if applicable).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// A token that can be used to list the previous page of runs (if
	// applicable).
	PrevPageToken types.String `tfsdk:"prev_page_token"`
	// A list of runs, from most recently started to least. Only included in the
	// response if there are runs to list.
	Runs []BaseRun `tfsdk:"runs"`
}

type NotebookOutput struct {
	// The value passed to
	// [dbutils.notebook.exit()](/notebooks/notebook-workflows.html#notebook-workflows-exit).
	// Databricks restricts this API to return the first 5 MB of the value. For
	// a larger result, your job can store the results in a cloud storage
	// service. This field is absent if `dbutils.notebook.exit()` was never
	// called.
	Result types.String `tfsdk:"result"`
	// Whether or not the result was truncated.
	Truncated types.Bool `tfsdk:"truncated"`
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
	BaseParameters map[types.String]types.String `tfsdk:"base_parameters"`
	// The path of the notebook to be run in the Databricks workspace or remote
	// repository. For notebooks stored in the Databricks workspace, the path
	// must be absolute and begin with a slash. For notebooks stored in a remote
	// repository, the path must be relative. This field is required.
	NotebookPath types.String `tfsdk:"notebook_path"`
	// Optional location type of the notebook. When set to `WORKSPACE`, the
	// notebook will be retrieved from the local Databricks workspace. When set
	// to `GIT`, the notebook will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`:
	// Notebook is located in Databricks workspace. * `GIT`: Notebook is located
	// in cloud Git provider.
	Source Source `tfsdk:"source"`
	// Optional `warehouse_id` to run the notebook on a SQL warehouse. Classic
	// SQL warehouses are NOT supported, please use serverless or pro SQL
	// warehouses.
	//
	// Note that SQL warehouses only support SQL cells; if the notebook contains
	// non-SQL cells, the run will fail.
	WarehouseId types.String `tfsdk:"warehouse_id"`
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

// Type always returns PauseStatus to satisfy [pflag.Value] interface
func (f *PauseStatus) Type() string {
	return "PauseStatus"
}

type PeriodicTriggerConfiguration struct {
	// The interval at which the trigger should run.
	Interval types.Int64 `tfsdk:"interval"`
	// The unit of time for the interval.
	Unit PeriodicTriggerConfigurationTimeUnit `tfsdk:"unit"`
}

type PeriodicTriggerConfigurationTimeUnit string

const PeriodicTriggerConfigurationTimeUnitDays PeriodicTriggerConfigurationTimeUnit = `DAYS`

const PeriodicTriggerConfigurationTimeUnitHours PeriodicTriggerConfigurationTimeUnit = `HOURS`

const PeriodicTriggerConfigurationTimeUnitTimeUnitUnspecified PeriodicTriggerConfigurationTimeUnit = `TIME_UNIT_UNSPECIFIED`

const PeriodicTriggerConfigurationTimeUnitWeeks PeriodicTriggerConfigurationTimeUnit = `WEEKS`

// String representation for [fmt.Print]
func (f *PeriodicTriggerConfigurationTimeUnit) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PeriodicTriggerConfigurationTimeUnit) Set(v string) error {
	switch v {
	case `DAYS`, `HOURS`, `TIME_UNIT_UNSPECIFIED`, `WEEKS`:
		*f = PeriodicTriggerConfigurationTimeUnit(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DAYS", "HOURS", "TIME_UNIT_UNSPECIFIED", "WEEKS"`, v)
	}
}

// Type always returns PeriodicTriggerConfigurationTimeUnit to satisfy [pflag.Value] interface
func (f *PeriodicTriggerConfigurationTimeUnit) Type() string {
	return "PeriodicTriggerConfigurationTimeUnit"
}

type PipelineParams struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh types.Bool `tfsdk:"full_refresh"`
}

type PipelineTask struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh types.Bool `tfsdk:"full_refresh"`
	// The full name of the pipeline task to execute.
	PipelineId types.String `tfsdk:"pipeline_id"`
}

type PythonWheelTask struct {
	// Named entry point to use, if it does not exist in the metadata of the
	// package it executes the function from the package directly using
	// `$packageName.$entryPoint()`
	EntryPoint types.String `tfsdk:"entry_point"`
	// Command-line parameters passed to Python wheel task in the form of
	// `["--name=task", "--data=dbfs:/path/to/data.json"]`. Leave it empty if
	// `parameters` is not null.
	NamedParameters map[types.String]types.String `tfsdk:"named_parameters"`
	// Name of the package to execute
	PackageName types.String `tfsdk:"package_name"`
	// Command-line parameters passed to Python wheel task. Leave it empty if
	// `named_parameters` is not null.
	Parameters []types.String `tfsdk:"parameters"`
}

type QueueSettings struct {
	// If true, enable queueing for the job. This is a required field.
	Enabled types.Bool `tfsdk:"enabled"`
}

type RepairHistoryItem struct {
	// The end time of the (repaired) run.
	EndTime types.Int64 `tfsdk:"end_time"`
	// The ID of the repair. Only returned for the items that represent a repair
	// in `repair_history`.
	Id types.Int64 `tfsdk:"id"`
	// The start time of the (repaired) run.
	StartTime types.Int64 `tfsdk:"start_time"`
	// The current state of the run.
	State *RunState `tfsdk:"state"`
	// The run IDs of the task runs that ran as part of this repair history
	// item.
	TaskRunIds []types.Int64 `tfsdk:"task_run_ids"`
	// The repair history item type. Indicates whether a run is the original run
	// or a repair run.
	Type RepairHistoryItemType `tfsdk:"type"`
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

// Type always returns RepairHistoryItemType to satisfy [pflag.Value] interface
func (f *RepairHistoryItemType) Type() string {
	return "RepairHistoryItemType"
}

type RepairRun struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []types.String `tfsdk:"dbt_commands"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables](/jobs.html\"#parameter-variables\") to set
	// parameters containing information about job runs.
	JarParams []types.String `tfsdk:"jar_params"`
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters map[types.String]types.String `tfsdk:"job_parameters"`
	// The ID of the latest repair. This parameter is not required when
	// repairing a run for the first time, but must be provided on subsequent
	// requests to repair the same run.
	LatestRepairId types.Int64 `tfsdk:"latest_repair_id"`
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
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[types.String]types.String `tfsdk:"notebook_params"`

	PipelineParams *PipelineParams `tfsdk:"pipeline_params"`

	PythonNamedParams map[types.String]types.String `tfsdk:"python_named_params"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []types.String `tfsdk:"python_params"`
	// If true, repair all failed tasks. Only one of `rerun_tasks` or
	// `rerun_all_failed_tasks` can be used.
	RerunAllFailedTasks types.Bool `tfsdk:"rerun_all_failed_tasks"`
	// If true, repair all tasks that depend on the tasks in `rerun_tasks`, even
	// if they were previously successful. Can be also used in combination with
	// `rerun_all_failed_tasks`.
	RerunDependentTasks types.Bool `tfsdk:"rerun_dependent_tasks"`
	// The task keys of the task runs to repair.
	RerunTasks []types.String `tfsdk:"rerun_tasks"`
	// The job run ID of the run to repair. The run must not be in progress.
	RunId types.Int64 `tfsdk:"run_id"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []types.String `tfsdk:"spark_submit_params"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[types.String]types.String `tfsdk:"sql_params"`
}

// Run repair was initiated.
type RepairRunResponse struct {
	// The ID of the repair. Must be provided in subsequent repairs using the
	// `latest_repair_id` field to ensure sequential repairs.
	RepairId types.Int64 `tfsdk:"repair_id"`
}

type ResetJob struct {
	// The canonical identifier of the job to reset. This field is required.
	JobId types.Int64 `tfsdk:"job_id"`
	// The new settings of the job. These settings completely replace the old
	// settings.
	//
	// Changes to the field `JobBaseSettings.timeout_seconds` are applied to
	// active runs. Changes to other fields are applied to future runs only.
	NewSettings JobSettings `tfsdk:"new_settings"`
}

type ResetResponse struct {
}

type ResolvedConditionTaskValues struct {
	Left types.String `tfsdk:"left"`

	Right types.String `tfsdk:"right"`
}

type ResolvedDbtTaskValues struct {
	Commands []types.String `tfsdk:"commands"`
}

type ResolvedNotebookTaskValues struct {
	BaseParameters map[types.String]types.String `tfsdk:"base_parameters"`
}

type ResolvedParamPairValues struct {
	Parameters map[types.String]types.String `tfsdk:"parameters"`
}

type ResolvedPythonWheelTaskValues struct {
	NamedParameters map[types.String]types.String `tfsdk:"named_parameters"`

	Parameters []types.String `tfsdk:"parameters"`
}

type ResolvedRunJobTaskValues struct {
	JobParameters map[types.String]types.String `tfsdk:"job_parameters"`

	Parameters map[types.String]types.String `tfsdk:"parameters"`
}

type ResolvedStringParamsValues struct {
	Parameters []types.String `tfsdk:"parameters"`
}

type ResolvedValues struct {
	ConditionTask *ResolvedConditionTaskValues `tfsdk:"condition_task"`

	DbtTask *ResolvedDbtTaskValues `tfsdk:"dbt_task"`

	NotebookTask *ResolvedNotebookTaskValues `tfsdk:"notebook_task"`

	PythonWheelTask *ResolvedPythonWheelTaskValues `tfsdk:"python_wheel_task"`

	RunJobTask *ResolvedRunJobTaskValues `tfsdk:"run_job_task"`

	SimulationTask *ResolvedParamPairValues `tfsdk:"simulation_task"`

	SparkJarTask *ResolvedStringParamsValues `tfsdk:"spark_jar_task"`

	SparkPythonTask *ResolvedStringParamsValues `tfsdk:"spark_python_task"`

	SparkSubmitTask *ResolvedStringParamsValues `tfsdk:"spark_submit_task"`

	SqlTask *ResolvedParamPairValues `tfsdk:"sql_task"`
}

// Run was retrieved successfully
type Run struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0\. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` \> 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber types.Int64 `tfsdk:"attempt_number"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration types.Int64 `tfsdk:"cleanup_duration"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `tfsdk:"cluster_instance"`
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec *ClusterSpec `tfsdk:"cluster_spec"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name"`
	// Description of the run
	Description types.String `tfsdk:"description"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime types.Int64 `tfsdk:"end_time"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration types.Int64 `tfsdk:"execution_duration"`
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
	GitSource *GitSource `tfsdk:"git_source"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `tfsdk:"job_clusters"`
	// The canonical identifier of the job that contains this run.
	JobId types.Int64 `tfsdk:"job_id"`
	// Job-level parameters used in the run
	JobParameters []JobParameter `tfsdk:"job_parameters"`
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob types.Int64 `tfsdk:"number_in_job"`
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId types.Int64 `tfsdk:"original_attempt_run_id"`
	// The parameters used for this run.
	OverridingParameters *RunParameters `tfsdk:"overriding_parameters"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration types.Int64 `tfsdk:"queue_duration"`
	// The repair history of the run.
	RepairHistory []RepairHistoryItem `tfsdk:"repair_history"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration types.Int64 `tfsdk:"run_duration"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId types.Int64 `tfsdk:"run_id"`
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName types.String `tfsdk:"run_name"`
	// The URL to the detail page of the run.
	RunPageUrl types.String `tfsdk:"run_page_url"`
	// The type of a run. * `JOB_RUN`: Normal job run. A run created with
	// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
	// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
	// :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	RunType RunType `tfsdk:"run_type"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *CronSchedule `tfsdk:"schedule"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration types.Int64 `tfsdk:"setup_duration"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime types.Int64 `tfsdk:"start_time"`
	// The current state of the run.
	State *RunState `tfsdk:"state"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls.
	Tasks []RunTask `tfsdk:"tasks"`
	// The type of trigger that fired this run.
	//
	// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
	// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
	// occurs you triggered a single run on demand through the UI or the API. *
	// `RETRY`: Indicates a run that is triggered as a retry of a previously
	// failed run. This occurs when you request to re-run the job in case of
	// failures. * `RUN_JOB_TASK`: Indicates a run that is triggered using a Run
	// Job task. * `FILE_ARRIVAL`: Indicates a run that is triggered by a file
	// arrival. * `TABLE`: Indicates a run that is triggered by a table update.
	Trigger TriggerType `tfsdk:"trigger"`
	// Additional details about what triggered the run
	TriggerInfo *TriggerInfo `tfsdk:"trigger_info"`
}

type RunConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	Left types.String `tfsdk:"left"`
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
	Op ConditionTaskOp `tfsdk:"op"`
	// The condition expression evaluation result. Filled in if the task was
	// successfully completed. Can be `"true"` or `"false"`
	Outcome types.String `tfsdk:"outcome"`
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right types.String `tfsdk:"right"`
}

type RunForEachTask struct {
	// Controls the number of active iterations task runs. Default is 20,
	// maximum allowed is 100.
	Concurrency types.Int64 `tfsdk:"concurrency"`
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs types.String `tfsdk:"inputs"`
	// Read only field. Populated for GetRun and ListRuns RPC calls and stores
	// the execution stats of an For each task
	Stats *ForEachStats `tfsdk:"stats"`
	// Configuration for the task that will be run for each element in the array
	Task Task `tfsdk:"task"`
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

// Type always returns RunIf to satisfy [pflag.Value] interface
func (f *RunIf) Type() string {
	return "RunIf"
}

type RunJobOutput struct {
	// The run id of the triggered job run
	RunId types.Int64 `tfsdk:"run_id"`
}

type RunJobTask struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []types.String `tfsdk:"dbt_commands"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables](/jobs.html\"#parameter-variables\") to set
	// parameters containing information about job runs.
	JarParams []types.String `tfsdk:"jar_params"`
	// ID of the job to trigger.
	JobId types.Int64 `tfsdk:"job_id"`
	// Job-level parameters used to trigger the job.
	JobParameters map[types.String]types.String `tfsdk:"job_parameters"`
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
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[types.String]types.String `tfsdk:"notebook_params"`

	PipelineParams *PipelineParams `tfsdk:"pipeline_params"`

	PythonNamedParams map[types.String]types.String `tfsdk:"python_named_params"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []types.String `tfsdk:"python_params"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []types.String `tfsdk:"spark_submit_params"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[types.String]types.String `tfsdk:"sql_params"`
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

// Type always returns RunLifeCycleState to satisfy [pflag.Value] interface
func (f *RunLifeCycleState) Type() string {
	return "RunLifeCycleState"
}

type RunNow struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []types.String `tfsdk:"dbt_commands"`
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
	//
	// For more information, see [How to ensure idempotency for jobs].
	//
	// [How to ensure idempotency for jobs]: https://kb.databricks.com/jobs/jobs-idempotency.html
	IdempotencyToken types.String `tfsdk:"idempotency_token"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables](/jobs.html\"#parameter-variables\") to set
	// parameters containing information about job runs.
	JarParams []types.String `tfsdk:"jar_params"`
	// The ID of the job to be executed
	JobId types.Int64 `tfsdk:"job_id"`
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters map[types.String]types.String `tfsdk:"job_parameters"`
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
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[types.String]types.String `tfsdk:"notebook_params"`

	PipelineParams *PipelineParams `tfsdk:"pipeline_params"`

	PythonNamedParams map[types.String]types.String `tfsdk:"python_named_params"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []types.String `tfsdk:"python_params"`
	// The queue settings of the run.
	Queue *QueueSettings `tfsdk:"queue"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []types.String `tfsdk:"spark_submit_params"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[types.String]types.String `tfsdk:"sql_params"`
}

// Run was started successfully.
type RunNowResponse struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob types.Int64 `tfsdk:"number_in_job"`
	// The globally unique ID of the newly triggered run.
	RunId types.Int64 `tfsdk:"run_id"`
}

// Run output was retrieved successfully.
type RunOutput struct {
	// The output of a dbt task, if available.
	DbtOutput *DbtOutput `tfsdk:"dbt_output"`
	// An error message indicating why a task failed or why output is not
	// available. The message is unstructured, and its exact format is subject
	// to change.
	Error types.String `tfsdk:"error"`
	// If there was an error executing the run, this field contains any
	// available stack traces.
	ErrorTrace types.String `tfsdk:"error_trace"`

	Info types.String `tfsdk:"info"`
	// The output from tasks that write to standard streams (stdout/stderr) such
	// as spark_jar_task, spark_python_task, python_wheel_task.
	//
	// It's not supported for the notebook_task, pipeline_task or
	// spark_submit_task.
	//
	// Databricks restricts this API to return the last 5 MB of these logs.
	Logs types.String `tfsdk:"logs"`
	// Whether the logs are truncated.
	LogsTruncated types.Bool `tfsdk:"logs_truncated"`
	// All details of the run except for its output.
	Metadata *Run `tfsdk:"metadata"`
	// The output of a notebook task, if available. A notebook task that
	// terminates (either successfully or with a failure) without calling
	// `dbutils.notebook.exit()` is considered to have an empty output. This
	// field is set but its result value is empty. Databricks restricts this API
	// to return the first 5 MB of the output. To return a larger result, use
	// the [ClusterLogConf] field to configure log storage for the job cluster.
	//
	// [ClusterLogConf]: https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterlogconf
	NotebookOutput *NotebookOutput `tfsdk:"notebook_output"`
	// The output of a run job task, if available
	RunJobOutput *RunJobOutput `tfsdk:"run_job_output"`
	// The output of a SQL task, if available.
	SqlOutput *SqlOutput `tfsdk:"sql_output"`
}

type RunParameters struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []types.String `tfsdk:"dbt_commands"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables](/jobs.html\"#parameter-variables\") to set
	// parameters containing information about job runs.
	JarParams []types.String `tfsdk:"jar_params"`
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
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[types.String]types.String `tfsdk:"notebook_params"`

	PipelineParams *PipelineParams `tfsdk:"pipeline_params"`

	PythonNamedParams map[types.String]types.String `tfsdk:"python_named_params"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []types.String `tfsdk:"python_params"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []types.String `tfsdk:"spark_submit_params"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[types.String]types.String `tfsdk:"sql_params"`
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
// canceled.
type RunResultState string

// The run was canceled at user request.
const RunResultStateCanceled RunResultState = `CANCELED`

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
	case `CANCELED`, `EXCLUDED`, `FAILED`, `MAXIMUM_CONCURRENT_RUNS_REACHED`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `TIMEDOUT`, `UPSTREAM_CANCELED`, `UPSTREAM_FAILED`:
		*f = RunResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "EXCLUDED", "FAILED", "MAXIMUM_CONCURRENT_RUNS_REACHED", "SUCCESS", "SUCCESS_WITH_FAILURES", "TIMEDOUT", "UPSTREAM_CANCELED", "UPSTREAM_FAILED"`, v)
	}
}

// Type always returns RunResultState to satisfy [pflag.Value] interface
func (f *RunResultState) Type() string {
	return "RunResultState"
}

// The current state of the run.
type RunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response.
	LifeCycleState RunLifeCycleState `tfsdk:"life_cycle_state"`
	// The reason indicating why the run was queued.
	QueueReason types.String `tfsdk:"queue_reason"`
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states.
	ResultState RunResultState `tfsdk:"result_state"`
	// A descriptive message for the current state. This field is unstructured,
	// and its exact format is subject to change.
	StateMessage types.String `tfsdk:"state_message"`
	// A value indicating whether a run was canceled manually by a user or by
	// the scheduler because the run timed out.
	UserCancelledOrTimedout types.Bool `tfsdk:"user_cancelled_or_timedout"`
}

// Used when outputting a child run, in GetRun or ListRuns.
type RunTask struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0\. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` \> 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber types.Int64 `tfsdk:"attempt_number"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration types.Int64 `tfsdk:"cleanup_duration"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `tfsdk:"cluster_instance"`
	// If condition_task, specifies a condition with an outcome that can be used
	// to control the execution of other tasks. Does not require a cluster to
	// execute and does not support retries or notifications.
	ConditionTask *RunConditionTask `tfsdk:"condition_task"`
	// If dbt_task, indicates that this must execute a dbt task. It requires
	// both Databricks SQL and the ability to use a serverless or a pro SQL
	// warehouse.
	DbtTask *DbtTask `tfsdk:"dbt_task"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn []TaskDependency `tfsdk:"depends_on"`
	// An optional description for this task.
	Description types.String `tfsdk:"description"`
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications `tfsdk:"email_notifications"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime types.Int64 `tfsdk:"end_time"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey types.String `tfsdk:"environment_key"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration types.Int64 `tfsdk:"execution_duration"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId types.String `tfsdk:"existing_cluster_id"`
	// If for_each_task, indicates that this task must execute the nested task
	// within it.
	ForEachTask *RunForEachTask `tfsdk:"for_each_task"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks. If `git_source` is set,
	// these tasks retrieve the file from the remote repository by default.
	// However, this behavior can be overridden by setting `source` to
	// `WORKSPACE` on the task. Note: dbt and SQL File tasks support only
	// version-controlled sources. If dbt or SQL File tasks are used,
	// `git_source` must be defined on the job.
	GitSource *GitSource `tfsdk:"git_source"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey types.String `tfsdk:"job_cluster_key"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library `tfsdk:"libraries"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec `tfsdk:"new_cluster"`
	// If notebook_task, indicates that this task must run a notebook. This
	// field may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `tfsdk:"notebook_task"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings *TaskNotificationSettings `tfsdk:"notification_settings"`
	// If pipeline_task, indicates that this task must execute a Pipeline.
	PipelineTask *PipelineTask `tfsdk:"pipeline_task"`
	// If python_wheel_task, indicates that this job must execute a PythonWheel.
	PythonWheelTask *PythonWheelTask `tfsdk:"python_wheel_task"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration types.Int64 `tfsdk:"queue_duration"`
	// Parameter values including resolved references
	ResolvedValues *ResolvedValues `tfsdk:"resolved_values"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration types.Int64 `tfsdk:"run_duration"`
	// The ID of the task run.
	RunId types.Int64 `tfsdk:"run_id"`
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf RunIf `tfsdk:"run_if"`
	// If run_job_task, indicates that this task must execute another job.
	RunJobTask *RunJobTask `tfsdk:"run_job_task"`

	RunPageUrl types.String `tfsdk:"run_page_url"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration types.Int64 `tfsdk:"setup_duration"`
	// If spark_jar_task, indicates that this task must run a JAR.
	SparkJarTask *SparkJarTask `tfsdk:"spark_jar_task"`
	// If spark_python_task, indicates that this task must run a Python file.
	SparkPythonTask *SparkPythonTask `tfsdk:"spark_python_task"`
	// If `spark_submit_task`, indicates that this task must be launched by the
	// spark submit script. This task can run only on new clusters.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	SparkSubmitTask *SparkSubmitTask `tfsdk:"spark_submit_task"`
	// If sql_task, indicates that this job must execute a SQL task.
	SqlTask *SqlTask `tfsdk:"sql_task"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime types.Int64 `tfsdk:"start_time"`
	// The current state of the run.
	State *RunState `tfsdk:"state"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey types.String `tfsdk:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications *WebhookNotifications `tfsdk:"webhook_notifications"`
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

// Workflow run. A run created with [dbutils.notebook.run].
//
// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
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

// Type always returns RunType to satisfy [pflag.Value] interface
func (f *RunType) Type() string {
	return "RunType"
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

// Type always returns Source to satisfy [pflag.Value] interface
func (f *Source) Type() string {
	return "Source"
}

type SparkJarTask struct {
	// Deprecated since 04/2016. Provide a `jar` through the `libraries` field
	// instead. For an example, see :method:jobs/create.
	JarUri types.String `tfsdk:"jar_uri"`
	// The full name of the class containing the main method to be executed.
	// This class must be contained in a JAR provided as a library.
	//
	// The code must use `SparkContext.getOrCreate` to obtain a Spark context;
	// otherwise, runs of the job fail.
	MainClassName types.String `tfsdk:"main_class_name"`
	// Parameters passed to the main method.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []types.String `tfsdk:"parameters"`
}

type SparkPythonTask struct {
	// Command line parameters passed to the Python file.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []types.String `tfsdk:"parameters"`
	// The Python file to be executed. Cloud file URIs (such as dbfs:/, s3:/,
	// adls:/, gcs:/) and workspace paths are supported. For python files stored
	// in the Databricks workspace, the path must be absolute and begin with
	// `/`. For files stored in a remote repository, the path must be relative.
	// This field is required.
	PythonFile types.String `tfsdk:"python_file"`
	// Optional location type of the Python file. When set to `WORKSPACE` or not
	// specified, the file will be retrieved from the local Databricks workspace
	// or cloud location (if the `python_file` has a URI format). When set to
	// `GIT`, the Python file will be retrieved from a Git repository defined in
	// `git_source`.
	//
	// * `WORKSPACE`: The Python file is located in a Databricks workspace or at
	// a cloud filesystem URI. * `GIT`: The Python file is located in a remote
	// Git repository.
	Source Source `tfsdk:"source"`
}

type SparkSubmitTask struct {
	// Command-line parameters passed to spark submit.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []types.String `tfsdk:"parameters"`
}

type SqlAlertOutput struct {
	// The state of the SQL alert.
	//
	// * UNKNOWN: alert yet to be evaluated * OK: alert evaluated and did not
	// fulfill trigger conditions * TRIGGERED: alert evaluated and fulfilled
	// trigger conditions
	AlertState SqlAlertState `tfsdk:"alert_state"`
	// The link to find the output results.
	OutputLink types.String `tfsdk:"output_link"`
	// The text of the SQL query. Can Run permission of the SQL query associated
	// with the SQL alert is required to view this field.
	QueryText types.String `tfsdk:"query_text"`
	// Information about SQL statements executed in the run.
	SqlStatements []SqlStatementOutput `tfsdk:"sql_statements"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id"`
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

// Type always returns SqlAlertState to satisfy [pflag.Value] interface
func (f *SqlAlertState) Type() string {
	return "SqlAlertState"
}

type SqlDashboardOutput struct {
	// The canonical identifier of the SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id"`
	// Widgets executed in the run. Only SQL query based widgets are listed.
	Widgets []SqlDashboardWidgetOutput `tfsdk:"widgets"`
}

type SqlDashboardWidgetOutput struct {
	// Time (in epoch milliseconds) when execution of the SQL widget ends.
	EndTime types.Int64 `tfsdk:"end_time"`
	// The information about the error when execution fails.
	Error *SqlOutputError `tfsdk:"error"`
	// The link to find the output results.
	OutputLink types.String `tfsdk:"output_link"`
	// Time (in epoch milliseconds) when execution of the SQL widget starts.
	StartTime types.Int64 `tfsdk:"start_time"`
	// The execution status of the SQL widget.
	Status SqlDashboardWidgetOutputStatus `tfsdk:"status"`
	// The canonical identifier of the SQL widget.
	WidgetId types.String `tfsdk:"widget_id"`
	// The title of the SQL widget.
	WidgetTitle types.String `tfsdk:"widget_title"`
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

// Type always returns SqlDashboardWidgetOutputStatus to satisfy [pflag.Value] interface
func (f *SqlDashboardWidgetOutputStatus) Type() string {
	return "SqlDashboardWidgetOutputStatus"
}

type SqlOutput struct {
	// The output of a SQL alert task, if available.
	AlertOutput *SqlAlertOutput `tfsdk:"alert_output"`
	// The output of a SQL dashboard task, if available.
	DashboardOutput *SqlDashboardOutput `tfsdk:"dashboard_output"`
	// The output of a SQL query task, if available.
	QueryOutput *SqlQueryOutput `tfsdk:"query_output"`
}

type SqlOutputError struct {
	// The error message when execution fails.
	Message types.String `tfsdk:"message"`
}

type SqlQueryOutput struct {
	EndpointId types.String `tfsdk:"endpoint_id"`
	// The link to find the output results.
	OutputLink types.String `tfsdk:"output_link"`
	// The text of the SQL query. Can Run permission of the SQL query is
	// required to view this field.
	QueryText types.String `tfsdk:"query_text"`
	// Information about SQL statements executed in the run.
	SqlStatements []SqlStatementOutput `tfsdk:"sql_statements"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

type SqlStatementOutput struct {
	// A key that can be used to look up query details.
	LookupKey types.String `tfsdk:"lookup_key"`
}

type SqlTask struct {
	// If alert, indicates that this job must refresh a SQL alert.
	Alert *SqlTaskAlert `tfsdk:"alert"`
	// If dashboard, indicates that this job must refresh a SQL dashboard.
	Dashboard *SqlTaskDashboard `tfsdk:"dashboard"`
	// If file, indicates that this job runs a SQL file in a remote Git
	// repository.
	File *SqlTaskFile `tfsdk:"file"`
	// Parameters to be used for each run of this job. The SQL alert task does
	// not support custom parameters.
	Parameters map[types.String]types.String `tfsdk:"parameters"`
	// If query, indicates that this job must execute a SQL query.
	Query *SqlTaskQuery `tfsdk:"query"`
	// The canonical identifier of the SQL warehouse. Recommended to use with
	// serverless or pro SQL warehouses. Classic SQL warehouses are only
	// supported for SQL alert, dashboard and query tasks and are limited to
	// scheduled single-task jobs.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

type SqlTaskAlert struct {
	// The canonical identifier of the SQL alert.
	AlertId types.String `tfsdk:"alert_id"`
	// If true, the alert notifications are not sent to subscribers.
	PauseSubscriptions types.Bool `tfsdk:"pause_subscriptions"`
	// If specified, alert notifications are sent to subscribers.
	Subscriptions []SqlTaskSubscription `tfsdk:"subscriptions"`
}

type SqlTaskDashboard struct {
	// Subject of the email sent to subscribers of this task.
	CustomSubject types.String `tfsdk:"custom_subject"`
	// The canonical identifier of the SQL dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// If true, the dashboard snapshot is not taken, and emails are not sent to
	// subscribers.
	PauseSubscriptions types.Bool `tfsdk:"pause_subscriptions"`
	// If specified, dashboard snapshots are sent to subscriptions.
	Subscriptions []SqlTaskSubscription `tfsdk:"subscriptions"`
}

type SqlTaskFile struct {
	// Path of the SQL file. Must be relative if the source is a remote Git
	// repository and absolute for workspace paths.
	Path types.String `tfsdk:"path"`
	// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL
	// file will be retrieved from the local Databricks workspace. When set to
	// `GIT`, the SQL file will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: SQL file is located in Databricks workspace. * `GIT`: SQL
	// file is located in cloud Git provider.
	Source Source `tfsdk:"source"`
}

type SqlTaskQuery struct {
	// The canonical identifier of the SQL query.
	QueryId types.String `tfsdk:"query_id"`
}

type SqlTaskSubscription struct {
	// The canonical identifier of the destination to receive email
	// notification. This parameter is mutually exclusive with user_name. You
	// cannot set both destination_id and user_name for subscription
	// notifications.
	DestinationId types.String `tfsdk:"destination_id"`
	// The user name to receive the subscription email. This parameter is
	// mutually exclusive with destination_id. You cannot set both
	// destination_id and user_name for subscription notifications.
	UserName types.String `tfsdk:"user_name"`
}

type SubmitRun struct {
	// List of permissions to set on the job.
	AccessControlList []iam.AccessControlRequest `tfsdk:"access_control_list"`
	// An optional set of email addresses notified when the run begins or
	// completes.
	EmailNotifications *JobEmailNotifications `tfsdk:"email_notifications"`
	// A list of task execution environment specifications that can be
	// referenced by tasks of this run.
	Environments []JobEnvironment `tfsdk:"environments"`
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
	GitSource *GitSource `tfsdk:"git_source"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `tfsdk:"health"`
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
	//
	// For more information, see [How to ensure idempotency for jobs].
	//
	// [How to ensure idempotency for jobs]: https://kb.databricks.com/jobs/jobs-idempotency.html
	IdempotencyToken types.String `tfsdk:"idempotency_token"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// run.
	NotificationSettings *JobNotificationSettings `tfsdk:"notification_settings"`
	// The queue settings of the one-time run.
	Queue *QueueSettings `tfsdk:"queue"`
	// Specifies the user or service principal that the job runs as. If not
	// specified, the job runs as the user who submits the request.
	RunAs *JobRunAs `tfsdk:"run_as"`
	// An optional name for the run. The default value is `Untitled`.
	RunName types.String `tfsdk:"run_name"`

	Tasks []SubmitTask `tfsdk:"tasks"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds"`
	// A collection of system notification IDs to notify when the run begins or
	// completes.
	WebhookNotifications *WebhookNotifications `tfsdk:"webhook_notifications"`
}

// Run was created and started successfully.
type SubmitRunResponse struct {
	// The canonical identifier for the newly submitted run.
	RunId types.Int64 `tfsdk:"run_id"`
}

type SubmitTask struct {
	// If condition_task, specifies a condition with an outcome that can be used
	// to control the execution of other tasks. Does not require a cluster to
	// execute and does not support retries or notifications.
	ConditionTask *ConditionTask `tfsdk:"condition_task"`
	// If dbt_task, indicates that this must execute a dbt task. It requires
	// both Databricks SQL and the ability to use a serverless or a pro SQL
	// warehouse.
	DbtTask *DbtTask `tfsdk:"dbt_task"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn []TaskDependency `tfsdk:"depends_on"`
	// An optional description for this task.
	Description types.String `tfsdk:"description"`
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications `tfsdk:"email_notifications"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey types.String `tfsdk:"environment_key"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId types.String `tfsdk:"existing_cluster_id"`
	// If for_each_task, indicates that this task must execute the nested task
	// within it.
	ForEachTask *ForEachTask `tfsdk:"for_each_task"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `tfsdk:"health"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library `tfsdk:"libraries"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec `tfsdk:"new_cluster"`
	// If notebook_task, indicates that this task must run a notebook. This
	// field may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `tfsdk:"notebook_task"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings *TaskNotificationSettings `tfsdk:"notification_settings"`
	// If pipeline_task, indicates that this task must execute a Pipeline.
	PipelineTask *PipelineTask `tfsdk:"pipeline_task"`
	// If python_wheel_task, indicates that this job must execute a PythonWheel.
	PythonWheelTask *PythonWheelTask `tfsdk:"python_wheel_task"`
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf RunIf `tfsdk:"run_if"`
	// If run_job_task, indicates that this task must execute another job.
	RunJobTask *RunJobTask `tfsdk:"run_job_task"`
	// If spark_jar_task, indicates that this task must run a JAR.
	SparkJarTask *SparkJarTask `tfsdk:"spark_jar_task"`
	// If spark_python_task, indicates that this task must run a Python file.
	SparkPythonTask *SparkPythonTask `tfsdk:"spark_python_task"`
	// If `spark_submit_task`, indicates that this task must be launched by the
	// spark submit script. This task can run only on new clusters.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	SparkSubmitTask *SparkSubmitTask `tfsdk:"spark_submit_task"`
	// If sql_task, indicates that this job must execute a SQL task.
	SqlTask *SqlTask `tfsdk:"sql_task"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey types.String `tfsdk:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications *WebhookNotifications `tfsdk:"webhook_notifications"`
}

type TableUpdateTriggerConfiguration struct {
	// The table(s) condition based on which to trigger a job run.
	Condition Condition `tfsdk:"condition"`
	// If set, the trigger starts a run only after the specified amount of time
	// has passed since the last time the trigger fired. The minimum allowed
	// value is 60 seconds.
	MinTimeBetweenTriggersSeconds types.Int64 `tfsdk:"min_time_between_triggers_seconds"`
	// A list of Delta tables to monitor for changes. The table name must be in
	// the format `catalog_name.schema_name.table_name`.
	TableNames []types.String `tfsdk:"table_names"`
	// If set, the trigger starts a run only after no table updates have
	// occurred for the specified time and can be used to wait for a series of
	// table updates before triggering a run. The minimum allowed value is 60
	// seconds.
	WaitAfterLastChangeSeconds types.Int64 `tfsdk:"wait_after_last_change_seconds"`
}

type Task struct {
	// If condition_task, specifies a condition with an outcome that can be used
	// to control the execution of other tasks. Does not require a cluster to
	// execute and does not support retries or notifications.
	ConditionTask *ConditionTask `tfsdk:"condition_task"`
	// If dbt_task, indicates that this must execute a dbt task. It requires
	// both Databricks SQL and the ability to use a serverless or a pro SQL
	// warehouse.
	DbtTask *DbtTask `tfsdk:"dbt_task"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete before executing this
	// task. The task will run only if the `run_if` condition is true. The key
	// is `task_key`, and the value is the name assigned to the dependent task.
	DependsOn []TaskDependency `tfsdk:"depends_on"`
	// An optional description for this task.
	Description types.String `tfsdk:"description"`
	// An option to disable auto optimization in serverless
	DisableAutoOptimization types.Bool `tfsdk:"disable_auto_optimization"`
	// An optional set of email addresses that is notified when runs of this
	// task begin or complete as well as when this task is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *TaskEmailNotifications `tfsdk:"email_notifications"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey types.String `tfsdk:"environment_key"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId types.String `tfsdk:"existing_cluster_id"`
	// If for_each_task, indicates that this task must execute the nested task
	// within it.
	ForEachTask *ForEachTask `tfsdk:"for_each_task"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `tfsdk:"health"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey types.String `tfsdk:"job_cluster_key"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library `tfsdk:"libraries"`
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value `-1` means
	// to retry indefinitely and the value `0` means to never retry.
	MaxRetries types.Int64 `tfsdk:"max_retries"`
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	MinRetryIntervalMillis types.Int64 `tfsdk:"min_retry_interval_millis"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec `tfsdk:"new_cluster"`
	// If notebook_task, indicates that this task must run a notebook. This
	// field may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `tfsdk:"notebook_task"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task.
	NotificationSettings *TaskNotificationSettings `tfsdk:"notification_settings"`
	// If pipeline_task, indicates that this task must execute a Pipeline.
	PipelineTask *PipelineTask `tfsdk:"pipeline_task"`
	// If python_wheel_task, indicates that this job must execute a PythonWheel.
	PythonWheelTask *PythonWheelTask `tfsdk:"python_wheel_task"`
	// An optional policy to specify whether to retry a job when it times out.
	// The default behavior is to not retry on timeout.
	RetryOnTimeout types.Bool `tfsdk:"retry_on_timeout"`
	// An optional value specifying the condition determining whether the task
	// is run once its dependencies have been completed.
	//
	// * `ALL_SUCCESS`: All dependencies have executed and succeeded *
	// `AT_LEAST_ONE_SUCCESS`: At least one dependency has succeeded *
	// `NONE_FAILED`: None of the dependencies have failed and at least one was
	// executed * `ALL_DONE`: All dependencies have been completed *
	// `AT_LEAST_ONE_FAILED`: At least one dependency failed * `ALL_FAILED`: ALl
	// dependencies have failed
	RunIf RunIf `tfsdk:"run_if"`
	// If run_job_task, indicates that this task must execute another job.
	RunJobTask *RunJobTask `tfsdk:"run_job_task"`
	// If spark_jar_task, indicates that this task must run a JAR.
	SparkJarTask *SparkJarTask `tfsdk:"spark_jar_task"`
	// If spark_python_task, indicates that this task must run a Python file.
	SparkPythonTask *SparkPythonTask `tfsdk:"spark_python_task"`
	// If `spark_submit_task`, indicates that this task must be launched by the
	// spark submit script. This task can run only on new clusters.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	SparkSubmitTask *SparkSubmitTask `tfsdk:"spark_submit_task"`
	// If sql_task, indicates that this job must execute a SQL task.
	SqlTask *SqlTask `tfsdk:"sql_task"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey types.String `tfsdk:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds"`
	// A collection of system notification IDs to notify when runs of this task
	// begin or complete. The default behavior is to not send any system
	// notifications.
	WebhookNotifications *WebhookNotifications `tfsdk:"webhook_notifications"`
}

type TaskDependency struct {
	// Can only be specified on condition task dependencies. The outcome of the
	// dependent task that must be met for this task to run.
	Outcome types.String `tfsdk:"outcome"`
	// The name of the task this task depends on.
	TaskKey types.String `tfsdk:"task_key"`
}

type TaskEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped.
	NoAlertForSkippedRuns types.Bool `tfsdk:"no_alert_for_skipped_runs"`
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded []types.String `tfsdk:"on_duration_warning_threshold_exceeded"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure []types.String `tfsdk:"on_failure"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []types.String `tfsdk:"on_start"`
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	OnStreamingBacklogExceeded []types.String `tfsdk:"on_streaming_backlog_exceeded"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess []types.String `tfsdk:"on_success"`
}

type TaskNotificationSettings struct {
	// If true, do not send notifications to recipients specified in `on_start`
	// for the retried runs and do not send notifications to recipients
	// specified in `on_failure` until the last retry of the run.
	AlertOnLastAttempt types.Bool `tfsdk:"alert_on_last_attempt"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns types.Bool `tfsdk:"no_alert_for_canceled_runs"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns types.Bool `tfsdk:"no_alert_for_skipped_runs"`
}

// Additional details about what triggered the run
type TriggerInfo struct {
	// The run id of the Run Job task run
	RunId types.Int64 `tfsdk:"run_id"`
}

type TriggerSettings struct {
	// File arrival trigger settings.
	FileArrival *FileArrivalTriggerConfiguration `tfsdk:"file_arrival"`
	// Whether this trigger is paused or not.
	PauseStatus PauseStatus `tfsdk:"pause_status"`
	// Periodic trigger settings.
	Periodic *PeriodicTriggerConfiguration `tfsdk:"periodic"`
	// Old table trigger settings name. Deprecated in favor of `table_update`.
	Table *TableUpdateTriggerConfiguration `tfsdk:"table"`

	TableUpdate *TableUpdateTriggerConfiguration `tfsdk:"table_update"`
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
// `TABLE`: Indicates a run that is triggered by a table update.
type TriggerType string

// Indicates a run that is triggered by a file arrival.
const TriggerTypeFileArrival TriggerType = `FILE_ARRIVAL`

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
	case `FILE_ARRIVAL`, `ONE_TIME`, `PERIODIC`, `RETRY`, `RUN_JOB_TASK`, `TABLE`:
		*f = TriggerType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE_ARRIVAL", "ONE_TIME", "PERIODIC", "RETRY", "RUN_JOB_TASK", "TABLE"`, v)
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
	FieldsToRemove []types.String `tfsdk:"fields_to_remove"`
	// The canonical identifier of the job to update. This field is required.
	JobId types.Int64 `tfsdk:"job_id"`
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
	NewSettings *JobSettings `tfsdk:"new_settings"`
}

type UpdateResponse struct {
}

type ViewItem struct {
	// Content of the view.
	Content types.String `tfsdk:"content"`
	// Name of the view item. In the case of code view, it would be the
	// notebook’s name. In the case of dashboard view, it would be the
	// dashboard’s name.
	Name types.String `tfsdk:"name"`
	// Type of the view item.
	Type ViewType `tfsdk:"type"`
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

// Type always returns ViewsToExport to satisfy [pflag.Value] interface
func (f *ViewsToExport) Type() string {
	return "ViewsToExport"
}

type Webhook struct {
	Id types.String `tfsdk:"id"`
}

type WebhookNotifications struct {
	// An optional list of system notification IDs to call when the duration of
	// a run exceeds the threshold specified for the `RUN_DURATION_SECONDS`
	// metric in the `health` field. A maximum of 3 destinations can be
	// specified for the `on_duration_warning_threshold_exceeded` property.
	OnDurationWarningThresholdExceeded []Webhook `tfsdk:"on_duration_warning_threshold_exceeded"`
	// An optional list of system notification IDs to call when the run fails. A
	// maximum of 3 destinations can be specified for the `on_failure` property.
	OnFailure []Webhook `tfsdk:"on_failure"`
	// An optional list of system notification IDs to call when the run starts.
	// A maximum of 3 destinations can be specified for the `on_start` property.
	OnStart []Webhook `tfsdk:"on_start"`
	// An optional list of system notification IDs to call when any streaming
	// backlog thresholds are exceeded for any stream. Streaming backlog
	// thresholds can be set in the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes. A maximum of 3 destinations
	// can be specified for the `on_streaming_backlog_exceeded` property.
	OnStreamingBacklogExceeded []Webhook `tfsdk:"on_streaming_backlog_exceeded"`
	// An optional list of system notification IDs to call when the run
	// completes successfully. A maximum of 3 destinations can be specified for
	// the `on_success` property.
	OnSuccess []Webhook `tfsdk:"on_success"`
}
