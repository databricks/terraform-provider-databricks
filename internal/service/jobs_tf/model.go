// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package jobs_tf

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type BaseJob struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime types.Int64 `tfsdk:"created_time" tf:"optional"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id" tf:"computed,optional"`
	// The canonical identifier for this job.
	JobId types.Int64 `tfsdk:"job_id" tf:"optional"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings types.List `tfsdk:"settings" tf:"optional,object"`
}

func (newState *BaseJob) SyncEffectiveFieldsDuringCreateOrUpdate(plan BaseJob) {
}

func (newState *BaseJob) SyncEffectiveFieldsDuringRead(existingState BaseJob) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BaseJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BaseJob) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings": reflect.TypeOf(JobSettings{}),
	}
}

// ToObjectType returns the representation of BaseJob in the Terraform plugin framework type
// system.
func (a BaseJob) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_time":               types.Int64Type,
			"creator_user_name":          types.StringType,
			"effective_budget_policy_id": types.StringType,
			"job_id":                     types.Int64Type,
			"settings": basetypes.ListType{
				ElemType: JobSettings{}.ToObjectType(ctx),
			},
		},
	}
}

type BaseRun struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber types.Int64 `tfsdk:"attempt_number" tf:"optional"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration types.Int64 `tfsdk:"cleanup_duration" tf:"optional"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance types.List `tfsdk:"cluster_instance" tf:"optional,object"`
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec types.List `tfsdk:"cluster_spec" tf:"optional,object"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// Description of the run
	Description types.String `tfsdk:"description" tf:"optional"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration types.Int64 `tfsdk:"execution_duration" tf:"optional"`
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
	GitSource types.List `tfsdk:"git_source" tf:"optional,object"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters types.List `tfsdk:"job_clusters" tf:"optional"`
	// The canonical identifier of the job that contains this run.
	JobId types.Int64 `tfsdk:"job_id" tf:"optional"`
	// Job-level parameters used in the run
	JobParameters types.List `tfsdk:"job_parameters" tf:"optional"`
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	JobRunId types.Int64 `tfsdk:"job_run_id" tf:"optional"`
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob types.Int64 `tfsdk:"number_in_job" tf:"optional"`
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId types.Int64 `tfsdk:"original_attempt_run_id" tf:"optional"`
	// The parameters used for this run.
	OverridingParameters types.List `tfsdk:"overriding_parameters" tf:"optional,object"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration types.Int64 `tfsdk:"queue_duration" tf:"optional"`
	// The repair history of the run.
	RepairHistory types.List `tfsdk:"repair_history" tf:"optional"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration types.Int64 `tfsdk:"run_duration" tf:"optional"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName types.String `tfsdk:"run_name" tf:"optional"`
	// The URL to the detail page of the run.
	RunPageUrl types.String `tfsdk:"run_page_url" tf:"optional"`
	// The type of a run. * `JOB_RUN`: Normal job run. A run created with
	// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
	// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
	// :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	RunType types.String `tfsdk:"run_type" tf:"optional"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration types.Int64 `tfsdk:"setup_duration" tf:"optional"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Deprecated. Please use the `status` field instead.
	State types.List `tfsdk:"state" tf:"optional,object"`
	// The current status of the run
	Status types.List `tfsdk:"status" tf:"optional,object"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls.
	Tasks types.List `tfsdk:"tasks" tf:"optional"`
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
	Trigger types.String `tfsdk:"trigger" tf:"optional"`
	// Additional details about what triggered the run
	TriggerInfo types.List `tfsdk:"trigger_info" tf:"optional,object"`
}

func (newState *BaseRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan BaseRun) {
}

func (newState *BaseRun) SyncEffectiveFieldsDuringRead(existingState BaseRun) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BaseRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BaseRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster_instance":      reflect.TypeOf(ClusterInstance{}),
		"cluster_spec":          reflect.TypeOf(ClusterSpec{}),
		"git_source":            reflect.TypeOf(GitSource{}),
		"job_clusters":          reflect.TypeOf(JobCluster{}),
		"job_parameters":        reflect.TypeOf(JobParameter{}),
		"overriding_parameters": reflect.TypeOf(RunParameters{}),
		"repair_history":        reflect.TypeOf(RepairHistoryItem{}),
		"schedule":              reflect.TypeOf(CronSchedule{}),
		"state":                 reflect.TypeOf(RunState{}),
		"status":                reflect.TypeOf(RunStatus{}),
		"tasks":                 reflect.TypeOf(RunTask{}),
		"trigger_info":          reflect.TypeOf(TriggerInfo{}),
	}
}

// ToObjectType returns the representation of BaseRun in the Terraform plugin framework type
// system.
func (a BaseRun) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attempt_number":   types.Int64Type,
			"cleanup_duration": types.Int64Type,
			"cluster_instance": basetypes.ListType{
				ElemType: ClusterInstance{}.ToObjectType(ctx),
			},
			"cluster_spec": basetypes.ListType{
				ElemType: ClusterSpec{}.ToObjectType(ctx),
			},
			"creator_user_name":  types.StringType,
			"description":        types.StringType,
			"end_time":           types.Int64Type,
			"execution_duration": types.Int64Type,
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.ToObjectType(ctx),
			},
			"job_clusters": basetypes.ListType{
				ElemType: JobCluster{}.ToObjectType(ctx),
			},
			"job_id": types.Int64Type,
			"job_parameters": basetypes.ListType{
				ElemType: JobParameter{}.ToObjectType(ctx),
			},
			"job_run_id":              types.Int64Type,
			"number_in_job":           types.Int64Type,
			"original_attempt_run_id": types.Int64Type,
			"overriding_parameters": basetypes.ListType{
				ElemType: RunParameters{}.ToObjectType(ctx),
			},
			"queue_duration": types.Int64Type,
			"repair_history": basetypes.ListType{
				ElemType: RepairHistoryItem{}.ToObjectType(ctx),
			},
			"run_duration": types.Int64Type,
			"run_id":       types.Int64Type,
			"run_name":     types.StringType,
			"run_page_url": types.StringType,
			"run_type":     types.StringType,
			"schedule": basetypes.ListType{
				ElemType: CronSchedule{}.ToObjectType(ctx),
			},
			"setup_duration": types.Int64Type,
			"start_time":     types.Int64Type,
			"state": basetypes.ListType{
				ElemType: RunState{}.ToObjectType(ctx),
			},
			"status": basetypes.ListType{
				ElemType: RunStatus{}.ToObjectType(ctx),
			},
			"tasks": basetypes.ListType{
				ElemType: RunTask{}.ToObjectType(ctx),
			},
			"trigger": types.StringType,
			"trigger_info": basetypes.ListType{
				ElemType: TriggerInfo{}.ToObjectType(ctx),
			},
		},
	}
}

type CancelAllRuns struct {
	// Optional boolean parameter to cancel all queued runs. If no job_id is
	// provided, all queued runs in the workspace are canceled.
	AllQueuedRuns types.Bool `tfsdk:"all_queued_runs" tf:"optional"`
	// The canonical identifier of the job to cancel all runs of.
	JobId types.Int64 `tfsdk:"job_id" tf:"optional"`
}

func (newState *CancelAllRuns) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelAllRuns) {
}

func (newState *CancelAllRuns) SyncEffectiveFieldsDuringRead(existingState CancelAllRuns) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelAllRuns.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelAllRuns) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CancelAllRuns in the Terraform plugin framework type
// system.
func (a CancelAllRuns) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_queued_runs": types.BoolType,
			"job_id":          types.Int64Type,
		},
	}
}

type CancelAllRunsResponse struct {
}

func (newState *CancelAllRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelAllRunsResponse) {
}

func (newState *CancelAllRunsResponse) SyncEffectiveFieldsDuringRead(existingState CancelAllRunsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelAllRunsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelAllRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CancelAllRunsResponse in the Terraform plugin framework type
// system.
func (a CancelAllRunsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CancelRun struct {
	// This field is required.
	RunId types.Int64 `tfsdk:"run_id" tf:""`
}

func (newState *CancelRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelRun) {
}

func (newState *CancelRun) SyncEffectiveFieldsDuringRead(existingState CancelRun) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CancelRun in the Terraform plugin framework type
// system.
func (a CancelRun) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

type CancelRunResponse struct {
}

func (newState *CancelRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelRunResponse) {
}

func (newState *CancelRunResponse) SyncEffectiveFieldsDuringRead(existingState CancelRunResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CancelRunResponse in the Terraform plugin framework type
// system.
func (a CancelRunResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The canonical identifier for the Spark context used by a run. This field
	// is filled in once the run begins execution. This value can be used to
	// view the Spark UI by browsing to
	// `/#setting/sparkui/$cluster_id/$spark_context_id`. The Spark UI continues
	// to be available after the run has completed.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	SparkContextId types.String `tfsdk:"spark_context_id" tf:"optional"`
}

func (newState *ClusterInstance) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterInstance) {
}

func (newState *ClusterInstance) SyncEffectiveFieldsDuringRead(existingState ClusterInstance) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterInstance.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterInstance) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ClusterInstance in the Terraform plugin framework type
// system.
func (a ClusterInstance) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":       types.StringType,
			"spark_context_id": types.StringType,
		},
	}
}

type ClusterSpec struct {
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId types.String `tfsdk:"existing_cluster_id" tf:"optional"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey types.String `tfsdk:"job_cluster_key" tf:"optional"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries types.List `tfsdk:"library" tf:"optional"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster compute.ClusterSpec `tfsdk:"new_cluster" tf:"optional,object"`
}

func (newState *ClusterSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterSpec) {
}

func (newState *ClusterSpec) SyncEffectiveFieldsDuringRead(existingState ClusterSpec) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"library":     reflect.TypeOf(compute.Library{}),
		"new_cluster": reflect.TypeOf(compute.ClusterSpec{}),
	}
}

// ToObjectType returns the representation of ClusterSpec in the Terraform plugin framework type
// system.
func (a ClusterSpec) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"existing_cluster_id": types.StringType,
			"job_cluster_key":     types.StringType,
			"library": basetypes.ListType{
				ElemType: compute_tf.Library{}.ToObjectType(ctx),
			},
			"new_cluster": basetypes.ListType{
				ElemType: compute_tf.ClusterSpec{}.ToObjectType(ctx),
			},
		},
	}
}

type ConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	Left types.String `tfsdk:"left" tf:""`
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
	Op types.String `tfsdk:"op" tf:""`
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right types.String `tfsdk:"right" tf:""`
}

func (newState *ConditionTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan ConditionTask) {
}

func (newState *ConditionTask) SyncEffectiveFieldsDuringRead(existingState ConditionTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConditionTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ConditionTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ConditionTask in the Terraform plugin framework type
// system.
func (a ConditionTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"left":  types.StringType,
			"op":    types.StringType,
			"right": types.StringType,
		},
	}
}

type Continuous struct {
	// Indicate whether the continuous execution of the job is paused or not.
	// Defaults to UNPAUSED.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
}

func (newState *Continuous) SyncEffectiveFieldsDuringCreateOrUpdate(plan Continuous) {
}

func (newState *Continuous) SyncEffectiveFieldsDuringRead(existingState Continuous) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Continuous.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Continuous) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of Continuous in the Terraform plugin framework type
// system.
func (a Continuous) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pause_status": types.StringType,
		},
	}
}

type CreateJob struct {
	// List of permissions to set on the job.
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id" tf:"optional"`
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous types.List `tfsdk:"continuous" tf:"optional,object"`
	// Deployment information for jobs managed by external sources.
	Deployment types.List `tfsdk:"deployment" tf:"optional,object"`
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode types.String `tfsdk:"edit_mode" tf:"optional"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications types.List `tfsdk:"email_notifications" tf:"optional,object"`
	// A list of task execution environment specifications that can be
	// referenced by serverless tasks of this job. An environment is required to
	// be present for serverless tasks. For serverless notebook tasks, the
	// environment is accessible in the notebook environment panel. For other
	// serverless tasks, the task environment is required to be specified using
	// environment_key in the task settings.
	Environments types.List `tfsdk:"environment" tf:"optional"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format types.String `tfsdk:"format" tf:"optional"`
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
	GitSource types.List `tfsdk:"git_source" tf:"optional,object"`
	// An optional set of health rules that can be defined for this job.
	Health types.List `tfsdk:"health" tf:"optional,object"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters types.List `tfsdk:"job_cluster" tf:"optional"`
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
	MaxConcurrentRuns types.Int64 `tfsdk:"max_concurrent_runs" tf:"optional"`
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings types.List `tfsdk:"notification_settings" tf:"optional,object"`
	// Job-level parameter definitions
	Parameters types.List `tfsdk:"parameter" tf:"optional"`
	// The queue settings of the job.
	Queue types.List `tfsdk:"queue" tf:"optional,object"`
	// Write-only setting. Specifies the user, service principal or group that
	// the job/pipeline runs as. If not specified, the job/pipeline runs as the
	// user who created the job/pipeline.
	//
	// Either `user_name` or `service_principal_name` should be specified. If
	// not, an error is thrown.
	RunAs types.List `tfsdk:"run_as" tf:"optional,object"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags types.Map `tfsdk:"tags" tf:"optional"`
	// A list of task specifications to be executed by this job.
	Tasks types.List `tfsdk:"task" tf:"optional"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds" tf:"optional"`
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger types.List `tfsdk:"trigger" tf:"optional,object"`
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications types.List `tfsdk:"webhook_notifications" tf:"optional,object"`
}

func (newState *CreateJob) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateJob) {
}

func (newState *CreateJob) SyncEffectiveFieldsDuringRead(existingState CreateJob) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateJob) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list":   reflect.TypeOf(JobAccessControlRequest{}),
		"continuous":            reflect.TypeOf(Continuous{}),
		"deployment":            reflect.TypeOf(JobDeployment{}),
		"email_notifications":   reflect.TypeOf(JobEmailNotifications{}),
		"environment":           reflect.TypeOf(JobEnvironment{}),
		"git_source":            reflect.TypeOf(GitSource{}),
		"health":                reflect.TypeOf(JobsHealthRules{}),
		"job_cluster":           reflect.TypeOf(JobCluster{}),
		"notification_settings": reflect.TypeOf(JobNotificationSettings{}),
		"parameter":             reflect.TypeOf(JobParameterDefinition{}),
		"queue":                 reflect.TypeOf(QueueSettings{}),
		"run_as":                reflect.TypeOf(JobRunAs{}),
		"schedule":              reflect.TypeOf(CronSchedule{}),
		"tags":                  reflect.TypeOf(types.String{}),
		"task":                  reflect.TypeOf(Task{}),
		"trigger":               reflect.TypeOf(TriggerSettings{}),
		"webhook_notifications": reflect.TypeOf(WebhookNotifications{}),
	}
}

// ToObjectType returns the representation of CreateJob in the Terraform plugin framework type
// system.
func (a CreateJob) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: JobAccessControlRequest{}.ToObjectType(ctx),
			},
			"budget_policy_id": types.StringType,
			"continuous": basetypes.ListType{
				ElemType: Continuous{}.ToObjectType(ctx),
			},
			"deployment": basetypes.ListType{
				ElemType: JobDeployment{}.ToObjectType(ctx),
			},
			"description": types.StringType,
			"edit_mode":   types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: JobEmailNotifications{}.ToObjectType(ctx),
			},
			"environment": basetypes.ListType{
				ElemType: JobEnvironment{}.ToObjectType(ctx),
			},
			"format": types.StringType,
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.ToObjectType(ctx),
			},
			"health": basetypes.ListType{
				ElemType: JobsHealthRules{}.ToObjectType(ctx),
			},
			"job_cluster": basetypes.ListType{
				ElemType: JobCluster{}.ToObjectType(ctx),
			},
			"max_concurrent_runs": types.Int64Type,
			"name":                types.StringType,
			"notification_settings": basetypes.ListType{
				ElemType: JobNotificationSettings{}.ToObjectType(ctx),
			},
			"parameter": basetypes.ListType{
				ElemType: JobParameterDefinition{}.ToObjectType(ctx),
			},
			"queue": basetypes.ListType{
				ElemType: QueueSettings{}.ToObjectType(ctx),
			},
			"run_as": basetypes.ListType{
				ElemType: JobRunAs{}.ToObjectType(ctx),
			},
			"schedule": basetypes.ListType{
				ElemType: CronSchedule{}.ToObjectType(ctx),
			},
			"tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"task": basetypes.ListType{
				ElemType: Task{}.ToObjectType(ctx),
			},
			"timeout_seconds": types.Int64Type,
			"trigger": basetypes.ListType{
				ElemType: TriggerSettings{}.ToObjectType(ctx),
			},
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.ToObjectType(ctx),
			},
		},
	}
}

// Job was created successfully
type CreateResponse struct {
	// The canonical identifier for the newly created job.
	JobId types.Int64 `tfsdk:"job_id" tf:"optional"`
}

func (newState *CreateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateResponse) {
}

func (newState *CreateResponse) SyncEffectiveFieldsDuringRead(existingState CreateResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CreateResponse in the Terraform plugin framework type
// system.
func (a CreateResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.Int64Type,
		},
	}
}

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// A Cron expression using Quartz syntax that describes the schedule for a
	// job. See [Cron Trigger] for details. This field is required.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression types.String `tfsdk:"quartz_cron_expression" tf:""`
	// A Java timezone ID. The schedule for a job is resolved with respect to
	// this timezone. See [Java TimeZone] for details. This field is required.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId types.String `tfsdk:"timezone_id" tf:""`
}

func (newState *CronSchedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan CronSchedule) {
}

func (newState *CronSchedule) SyncEffectiveFieldsDuringRead(existingState CronSchedule) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CronSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CronSchedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CronSchedule in the Terraform plugin framework type
// system.
func (a CronSchedule) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pause_status":           types.StringType,
			"quartz_cron_expression": types.StringType,
			"timezone_id":            types.StringType,
		},
	}
}

type DbtOutput struct {
	// An optional map of headers to send when retrieving the artifact from the
	// `artifacts_link`.
	ArtifactsHeaders types.Map `tfsdk:"artifacts_headers" tf:"optional"`
	// A pre-signed URL to download the (compressed) dbt artifacts. This link is
	// valid for a limited time (30 minutes). This information is only available
	// after the run has finished.
	ArtifactsLink types.String `tfsdk:"artifacts_link" tf:"optional"`
}

func (newState *DbtOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DbtOutput) {
}

func (newState *DbtOutput) SyncEffectiveFieldsDuringRead(existingState DbtOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DbtOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DbtOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"artifacts_headers": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of DbtOutput in the Terraform plugin framework type
// system.
func (a DbtOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifacts_headers": basetypes.MapType{
				ElemType: types.StringType,
			},
			"artifacts_link": types.StringType,
		},
	}
}

type DbtTask struct {
	// Optional name of the catalog to use. The value is the top level in the
	// 3-level namespace of Unity Catalog (catalog / schema / relation). The
	// catalog value can only be specified if a warehouse_id is specified.
	// Requires dbt-databricks >= 1.1.1.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// A list of dbt commands to execute. All commands must start with `dbt`.
	// This parameter must not be empty. A maximum of up to 10 commands can be
	// provided.
	Commands types.List `tfsdk:"commands" tf:""`
	// Optional (relative) path to the profiles directory. Can only be specified
	// if no warehouse_id is specified. If no warehouse_id is specified and this
	// folder is unset, the root directory is used.
	ProfilesDirectory types.String `tfsdk:"profiles_directory" tf:"optional"`
	// Path to the project directory. Optional for Git sourced tasks, in which
	// case if no value is provided, the root of the Git repository is used.
	ProjectDirectory types.String `tfsdk:"project_directory" tf:"optional"`
	// Optional schema to write to. This parameter is only used when a
	// warehouse_id is also provided. If not provided, the `default` schema is
	// used.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// Optional location type of the project directory. When set to `WORKSPACE`,
	// the project will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the project will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: Project is located in Databricks workspace. * `GIT`:
	// Project is located in cloud Git provider.
	Source types.String `tfsdk:"source" tf:"optional"`
	// ID of the SQL warehouse to connect to. If provided, we automatically
	// generate and provide the profile and connection details to dbt. It can be
	// overridden on a per-command basis by using the `--profiles-dir` command
	// line argument.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *DbtTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan DbtTask) {
}

func (newState *DbtTask) SyncEffectiveFieldsDuringRead(existingState DbtTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DbtTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DbtTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"commands": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of DbtTask in the Terraform plugin framework type
// system.
func (a DbtTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog": types.StringType,
			"commands": basetypes.ListType{
				ElemType: types.StringType,
			},
			"profiles_directory": types.StringType,
			"project_directory":  types.StringType,
			"schema":             types.StringType,
			"source":             types.StringType,
			"warehouse_id":       types.StringType,
		},
	}
}

type DeleteJob struct {
	// The canonical identifier of the job to delete. This field is required.
	JobId types.Int64 `tfsdk:"job_id" tf:""`
}

func (newState *DeleteJob) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteJob) {
}

func (newState *DeleteJob) SyncEffectiveFieldsDuringRead(existingState DeleteJob) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteJob) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteJob in the Terraform plugin framework type
// system.
func (a DeleteJob) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.Int64Type,
		},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteResponse in the Terraform plugin framework type
// system.
func (a DeleteResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRun struct {
	// ID of the run to delete.
	RunId types.Int64 `tfsdk:"run_id" tf:""`
}

func (newState *DeleteRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRun) {
}

func (newState *DeleteRun) SyncEffectiveFieldsDuringRead(existingState DeleteRun) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteRun in the Terraform plugin framework type
// system.
func (a DeleteRun) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

type DeleteRunResponse struct {
}

func (newState *DeleteRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRunResponse) {
}

func (newState *DeleteRunResponse) SyncEffectiveFieldsDuringRead(existingState DeleteRunResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteRunResponse in the Terraform plugin framework type
// system.
func (a DeleteRunResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Represents a change to the job cluster's settings that would be required for
// the job clusters to become compliant with their policies.
type EnforcePolicyComplianceForJobResponseJobClusterSettingsChange struct {
	// The field where this change would be made, prepended with the job cluster
	// key.
	Field types.String `tfsdk:"field" tf:"optional"`
	// The new value of this field after enforcing policy compliance (either a
	// number, a boolean, or a string) converted to a string. This is intended
	// to be read by a human. The typed new value of this field can be retrieved
	// by reading the settings field in the API response.
	NewValue types.String `tfsdk:"new_value" tf:"optional"`
	// The previous value of this field before enforcing policy compliance
	// (either a number, a boolean, or a string) converted to a string. This is
	// intended to be read by a human. The type of the field can be retrieved by
	// reading the settings field in the API response.
	PreviousValue types.String `tfsdk:"previous_value" tf:"optional"`
}

func (newState *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) {
}

func (newState *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) SyncEffectiveFieldsDuringRead(existingState EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnforcePolicyComplianceForJobResponseJobClusterSettingsChange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of EnforcePolicyComplianceForJobResponseJobClusterSettingsChange in the Terraform plugin framework type
// system.
func (a EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"field":          types.StringType,
			"new_value":      types.StringType,
			"previous_value": types.StringType,
		},
	}
}

type EnforcePolicyComplianceRequest struct {
	// The ID of the job you want to enforce policy compliance on.
	JobId types.Int64 `tfsdk:"job_id" tf:""`
	// If set, previews changes made to the job to comply with its policy, but
	// does not update the job.
	ValidateOnly types.Bool `tfsdk:"validate_only" tf:"optional"`
}

func (newState *EnforcePolicyComplianceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnforcePolicyComplianceRequest) {
}

func (newState *EnforcePolicyComplianceRequest) SyncEffectiveFieldsDuringRead(existingState EnforcePolicyComplianceRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnforcePolicyComplianceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnforcePolicyComplianceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of EnforcePolicyComplianceRequest in the Terraform plugin framework type
// system.
func (a EnforcePolicyComplianceRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id":        types.Int64Type,
			"validate_only": types.BoolType,
		},
	}
}

type EnforcePolicyComplianceResponse struct {
	// Whether any changes have been made to the job cluster settings for the
	// job to become compliant with its policies.
	HasChanges types.Bool `tfsdk:"has_changes" tf:"optional"`
	// A list of job cluster changes that have been made to the job’s cluster
	// settings in order for all job clusters to become compliant with their
	// policies.
	JobClusterChanges types.List `tfsdk:"job_cluster_changes" tf:"optional"`
	// Updated job settings after policy enforcement. Policy enforcement only
	// applies to job clusters that are created when running the job (which are
	// specified in new_cluster) and does not apply to existing all-purpose
	// clusters. Updated job settings are derived by applying policy default
	// values to the existing job clusters in order to satisfy policy
	// requirements.
	Settings types.List `tfsdk:"settings" tf:"optional,object"`
}

func (newState *EnforcePolicyComplianceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnforcePolicyComplianceResponse) {
}

func (newState *EnforcePolicyComplianceResponse) SyncEffectiveFieldsDuringRead(existingState EnforcePolicyComplianceResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnforcePolicyComplianceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnforcePolicyComplianceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"job_cluster_changes": reflect.TypeOf(EnforcePolicyComplianceForJobResponseJobClusterSettingsChange{}),
		"settings":            reflect.TypeOf(JobSettings{}),
	}
}

// ToObjectType returns the representation of EnforcePolicyComplianceResponse in the Terraform plugin framework type
// system.
func (a EnforcePolicyComplianceResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"has_changes": types.BoolType,
			"job_cluster_changes": basetypes.ListType{
				ElemType: EnforcePolicyComplianceForJobResponseJobClusterSettingsChange{}.ToObjectType(ctx),
			},
			"settings": basetypes.ListType{
				ElemType: JobSettings{}.ToObjectType(ctx),
			},
		},
	}
}

// Run was exported successfully.
type ExportRunOutput struct {
	// The exported content in HTML format (one for every view item). To extract
	// the HTML notebook from the JSON response, download and run this [Python
	// script].
	//
	// [Python script]: https://docs.databricks.com/en/_static/examples/extract.py
	Views types.List `tfsdk:"views" tf:"optional"`
}

func (newState *ExportRunOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportRunOutput) {
}

func (newState *ExportRunOutput) SyncEffectiveFieldsDuringRead(existingState ExportRunOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportRunOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExportRunOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"views": reflect.TypeOf(ViewItem{}),
	}
}

// ToObjectType returns the representation of ExportRunOutput in the Terraform plugin framework type
// system.
func (a ExportRunOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"views": basetypes.ListType{
				ElemType: ViewItem{}.ToObjectType(ctx),
			},
		},
	}
}

// Export and retrieve a job run
type ExportRunRequest struct {
	// The canonical identifier for the run. This field is required.
	RunId types.Int64 `tfsdk:"-"`
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	ViewsToExport types.String `tfsdk:"-"`
}

func (newState *ExportRunRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportRunRequest) {
}

func (newState *ExportRunRequest) SyncEffectiveFieldsDuringRead(existingState ExportRunRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportRunRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExportRunRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ExportRunRequest in the Terraform plugin framework type
// system.
func (a ExportRunRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id":          types.Int64Type,
			"views_to_export": types.StringType,
		},
	}
}

type FileArrivalTriggerConfiguration struct {
	// If set, the trigger starts a run only after the specified amount of time
	// passed since the last time the trigger fired. The minimum allowed value
	// is 60 seconds
	MinTimeBetweenTriggersSeconds types.Int64 `tfsdk:"min_time_between_triggers_seconds" tf:"optional"`
	// URL to be monitored for file arrivals. The path must point to the root or
	// a subpath of the external location.
	Url types.String `tfsdk:"url" tf:""`
	// If set, the trigger starts a run only after no file activity has occurred
	// for the specified amount of time. This makes it possible to wait for a
	// batch of incoming files to arrive before triggering a run. The minimum
	// allowed value is 60 seconds.
	WaitAfterLastChangeSeconds types.Int64 `tfsdk:"wait_after_last_change_seconds" tf:"optional"`
}

func (newState *FileArrivalTriggerConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileArrivalTriggerConfiguration) {
}

func (newState *FileArrivalTriggerConfiguration) SyncEffectiveFieldsDuringRead(existingState FileArrivalTriggerConfiguration) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileArrivalTriggerConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FileArrivalTriggerConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of FileArrivalTriggerConfiguration in the Terraform plugin framework type
// system.
func (a FileArrivalTriggerConfiguration) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"min_time_between_triggers_seconds": types.Int64Type,
			"url":                               types.StringType,
			"wait_after_last_change_seconds":    types.Int64Type,
		},
	}
}

type ForEachStats struct {
	// Sample of 3 most common error messages occurred during the iteration.
	ErrorMessageStats types.List `tfsdk:"error_message_stats" tf:"optional"`
	// Describes stats of the iteration. Only latest retries are considered.
	TaskRunStats types.List `tfsdk:"task_run_stats" tf:"optional,object"`
}

func (newState *ForEachStats) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForEachStats) {
}

func (newState *ForEachStats) SyncEffectiveFieldsDuringRead(existingState ForEachStats) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForEachStats.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ForEachStats) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error_message_stats": reflect.TypeOf(ForEachTaskErrorMessageStats{}),
		"task_run_stats":      reflect.TypeOf(ForEachTaskTaskRunStats{}),
	}
}

// ToObjectType returns the representation of ForEachStats in the Terraform plugin framework type
// system.
func (a ForEachStats) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_message_stats": basetypes.ListType{
				ElemType: ForEachTaskErrorMessageStats{}.ToObjectType(ctx),
			},
			"task_run_stats": basetypes.ListType{
				ElemType: ForEachTaskTaskRunStats{}.ToObjectType(ctx),
			},
		},
	}
}

type ForEachTask struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	Concurrency types.Int64 `tfsdk:"concurrency" tf:"optional"`
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs types.String `tfsdk:"inputs" tf:""`
	// Configuration for the task that will be run for each element in the array
	Task types.List `tfsdk:"task" tf:"object"`
}

func (newState *ForEachTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForEachTask) {
}

func (newState *ForEachTask) SyncEffectiveFieldsDuringRead(existingState ForEachTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForEachTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ForEachTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"task": reflect.TypeOf(Task{}),
	}
}

// ToObjectType returns the representation of ForEachTask in the Terraform plugin framework type
// system.
func (a ForEachTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"concurrency": types.Int64Type,
			"inputs":      types.StringType,
			"task": basetypes.ListType{
				ElemType: Task{}.ToObjectType(ctx),
			},
		},
	}
}

type ForEachTaskErrorMessageStats struct {
	// Describes the count of such error message encountered during the
	// iterations.
	Count types.Int64 `tfsdk:"count" tf:"optional"`
	// Describes the error message occured during the iterations.
	ErrorMessage types.String `tfsdk:"error_message" tf:"optional"`
	// Describes the termination reason for the error message.
	TerminationCategory types.String `tfsdk:"termination_category" tf:"optional"`
}

func (newState *ForEachTaskErrorMessageStats) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForEachTaskErrorMessageStats) {
}

func (newState *ForEachTaskErrorMessageStats) SyncEffectiveFieldsDuringRead(existingState ForEachTaskErrorMessageStats) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForEachTaskErrorMessageStats.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ForEachTaskErrorMessageStats) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ForEachTaskErrorMessageStats in the Terraform plugin framework type
// system.
func (a ForEachTaskErrorMessageStats) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"count":                types.Int64Type,
			"error_message":        types.StringType,
			"termination_category": types.StringType,
		},
	}
}

type ForEachTaskTaskRunStats struct {
	// Describes the iteration runs having an active lifecycle state or an
	// active run sub state.
	ActiveIterations types.Int64 `tfsdk:"active_iterations" tf:"optional"`
	// Describes the number of failed and succeeded iteration runs.
	CompletedIterations types.Int64 `tfsdk:"completed_iterations" tf:"optional"`
	// Describes the number of failed iteration runs.
	FailedIterations types.Int64 `tfsdk:"failed_iterations" tf:"optional"`
	// Describes the number of iteration runs that have been scheduled.
	ScheduledIterations types.Int64 `tfsdk:"scheduled_iterations" tf:"optional"`
	// Describes the number of succeeded iteration runs.
	SucceededIterations types.Int64 `tfsdk:"succeeded_iterations" tf:"optional"`
	// Describes the length of the list of items to iterate over.
	TotalIterations types.Int64 `tfsdk:"total_iterations" tf:"optional"`
}

func (newState *ForEachTaskTaskRunStats) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForEachTaskTaskRunStats) {
}

func (newState *ForEachTaskTaskRunStats) SyncEffectiveFieldsDuringRead(existingState ForEachTaskTaskRunStats) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForEachTaskTaskRunStats.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ForEachTaskTaskRunStats) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ForEachTaskTaskRunStats in the Terraform plugin framework type
// system.
func (a ForEachTaskTaskRunStats) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active_iterations":    types.Int64Type,
			"completed_iterations": types.Int64Type,
			"failed_iterations":    types.Int64Type,
			"scheduled_iterations": types.Int64Type,
			"succeeded_iterations": types.Int64Type,
			"total_iterations":     types.Int64Type,
		},
	}
}

// Get job permission levels
type GetJobPermissionLevelsRequest struct {
	// The job for which to get or manage permissions.
	JobId types.String `tfsdk:"-"`
}

func (newState *GetJobPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetJobPermissionLevelsRequest) {
}

func (newState *GetJobPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetJobPermissionLevelsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetJobPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetJobPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetJobPermissionLevelsRequest in the Terraform plugin framework type
// system.
func (a GetJobPermissionLevelsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.StringType,
		},
	}
}

type GetJobPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetJobPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetJobPermissionLevelsResponse) {
}

func (newState *GetJobPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetJobPermissionLevelsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetJobPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetJobPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(JobPermissionsDescription{}),
	}
}

// ToObjectType returns the representation of GetJobPermissionLevelsResponse in the Terraform plugin framework type
// system.
func (a GetJobPermissionLevelsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: JobPermissionsDescription{}.ToObjectType(ctx),
			},
		},
	}
}

// Get job permissions
type GetJobPermissionsRequest struct {
	// The job for which to get or manage permissions.
	JobId types.String `tfsdk:"-"`
}

func (newState *GetJobPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetJobPermissionsRequest) {
}

func (newState *GetJobPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetJobPermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetJobPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetJobPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetJobPermissionsRequest in the Terraform plugin framework type
// system.
func (a GetJobPermissionsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.StringType,
		},
	}
}

// Get a single job
type GetJobRequest struct {
	// The canonical identifier of the job to retrieve information about. This
	// field is required.
	JobId types.Int64 `tfsdk:"-"`
}

func (newState *GetJobRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetJobRequest) {
}

func (newState *GetJobRequest) SyncEffectiveFieldsDuringRead(existingState GetJobRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetJobRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetJobRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetJobRequest in the Terraform plugin framework type
// system.
func (a GetJobRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.Int64Type,
		},
	}
}

// Get job policy compliance
type GetPolicyComplianceRequest struct {
	// The ID of the job whose compliance status you are requesting.
	JobId types.Int64 `tfsdk:"-"`
}

func (newState *GetPolicyComplianceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPolicyComplianceRequest) {
}

func (newState *GetPolicyComplianceRequest) SyncEffectiveFieldsDuringRead(existingState GetPolicyComplianceRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPolicyComplianceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPolicyComplianceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetPolicyComplianceRequest in the Terraform plugin framework type
// system.
func (a GetPolicyComplianceRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.Int64Type,
		},
	}
}

type GetPolicyComplianceResponse struct {
	// Whether the job is compliant with its policies or not. Jobs could be out
	// of compliance if a policy they are using was updated after the job was
	// last edited and some of its job clusters no longer comply with their
	// updated policies.
	IsCompliant types.Bool `tfsdk:"is_compliant" tf:"optional"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	Violations types.Map `tfsdk:"violations" tf:"optional"`
}

func (newState *GetPolicyComplianceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPolicyComplianceResponse) {
}

func (newState *GetPolicyComplianceResponse) SyncEffectiveFieldsDuringRead(existingState GetPolicyComplianceResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPolicyComplianceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPolicyComplianceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"violations": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of GetPolicyComplianceResponse in the Terraform plugin framework type
// system.
func (a GetPolicyComplianceResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_compliant": types.BoolType,
			"violations": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// Get the output for a single run
type GetRunOutputRequest struct {
	// The canonical identifier for the run.
	RunId types.Int64 `tfsdk:"-"`
}

func (newState *GetRunOutputRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRunOutputRequest) {
}

func (newState *GetRunOutputRequest) SyncEffectiveFieldsDuringRead(existingState GetRunOutputRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRunOutputRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRunOutputRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetRunOutputRequest in the Terraform plugin framework type
// system.
func (a GetRunOutputRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

// Get a single job run
type GetRunRequest struct {
	// Whether to include the repair history in the response.
	IncludeHistory types.Bool `tfsdk:"-"`
	// Whether to include resolved parameter values in the response.
	IncludeResolvedValues types.Bool `tfsdk:"-"`
	// To list the next page of job tasks, set this field to the value of the
	// `next_page_token` returned in the GetJob response.
	PageToken types.String `tfsdk:"-"`
	// The canonical identifier of the run for which to retrieve the metadata.
	// This field is required.
	RunId types.Int64 `tfsdk:"-"`
}

func (newState *GetRunRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRunRequest) {
}

func (newState *GetRunRequest) SyncEffectiveFieldsDuringRead(existingState GetRunRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRunRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRunRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetRunRequest in the Terraform plugin framework type
// system.
func (a GetRunRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_history":         types.BoolType,
			"include_resolved_values": types.BoolType,
			"page_token":              types.StringType,
			"run_id":                  types.Int64Type,
		},
	}
}

// Read-only state of the remote repository at the time the job was run. This
// field is only included on job runs.
type GitSnapshot struct {
	// Commit that was used to execute the run. If git_branch was specified,
	// this points to the HEAD of the branch at the time of the run; if git_tag
	// was specified, this points to the commit the tag points to.
	UsedCommit types.String `tfsdk:"used_commit" tf:"optional"`
}

func (newState *GitSnapshot) SyncEffectiveFieldsDuringCreateOrUpdate(plan GitSnapshot) {
}

func (newState *GitSnapshot) SyncEffectiveFieldsDuringRead(existingState GitSnapshot) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GitSnapshot.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GitSnapshot) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GitSnapshot in the Terraform plugin framework type
// system.
func (a GitSnapshot) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"used_commit": types.StringType,
		},
	}
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
	GitBranch types.String `tfsdk:"branch" tf:"optional"`
	// Commit to be checked out and used by this job. This field cannot be
	// specified in conjunction with git_branch or git_tag.
	GitCommit types.String `tfsdk:"commit" tf:"optional"`
	// Unique identifier of the service used to host the Git repository. The
	// value is case insensitive.
	GitProvider types.String `tfsdk:"git_provider" tf:""`
	// Read-only state of the remote repository at the time the job was run.
	// This field is only included on job runs.
	GitSnapshot types.List `tfsdk:"git_snapshot" tf:"optional,object"`
	// Name of the tag to be checked out and used by this job. This field cannot
	// be specified in conjunction with git_branch or git_commit.
	GitTag types.String `tfsdk:"tag" tf:"optional"`
	// URL of the repository to be cloned by this job.
	GitUrl types.String `tfsdk:"url" tf:""`
	// The source of the job specification in the remote repository when the job
	// is source controlled.
	JobSource types.List `tfsdk:"job_source" tf:"optional,object"`
}

func (newState *GitSource) SyncEffectiveFieldsDuringCreateOrUpdate(plan GitSource) {
}

func (newState *GitSource) SyncEffectiveFieldsDuringRead(existingState GitSource) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GitSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GitSource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"git_snapshot": reflect.TypeOf(GitSnapshot{}),
		"job_source":   reflect.TypeOf(JobSource{}),
	}
}

// ToObjectType returns the representation of GitSource in the Terraform plugin framework type
// system.
func (a GitSource) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":       types.StringType,
			"commit":       types.StringType,
			"git_provider": types.StringType,
			"git_snapshot": basetypes.ListType{
				ElemType: GitSnapshot{}.ToObjectType(ctx),
			},
			"tag": types.StringType,
			"url": types.StringType,
			"job_source": basetypes.ListType{
				ElemType: JobSource{}.ToObjectType(ctx),
			},
		},
	}
}

// Job was retrieved successfully.
type Job struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime types.Int64 `tfsdk:"created_time" tf:"optional"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id" tf:"computed,optional"`
	// The canonical identifier for this job.
	JobId types.Int64 `tfsdk:"job_id" tf:"optional"`
	// The email of an active workspace user or the application ID of a service
	// principal that the job runs as. This value can be changed by setting the
	// `run_as` field when creating or updating a job.
	//
	// By default, `run_as_user_name` is based on the current job settings and
	// is set to the creator of the job if job access control is disabled or to
	// the user with the `is_owner` permission if job access control is enabled.
	RunAsUserName types.String `tfsdk:"run_as_user_name" tf:"optional"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings types.List `tfsdk:"settings" tf:"optional,object"`
}

func (newState *Job) SyncEffectiveFieldsDuringCreateOrUpdate(plan Job) {
}

func (newState *Job) SyncEffectiveFieldsDuringRead(existingState Job) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Job.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Job) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings": reflect.TypeOf(JobSettings{}),
	}
}

// ToObjectType returns the representation of Job in the Terraform plugin framework type
// system.
func (a Job) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_time":               types.Int64Type,
			"creator_user_name":          types.StringType,
			"effective_budget_policy_id": types.StringType,
			"job_id":                     types.Int64Type,
			"run_as_user_name":           types.StringType,
			"settings": basetypes.ListType{
				ElemType: JobSettings{}.ToObjectType(ctx),
			},
		},
	}
}

type JobAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *JobAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobAccessControlRequest) {
}

func (newState *JobAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState JobAccessControlRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of JobAccessControlRequest in the Terraform plugin framework type
// system.
func (a JobAccessControlRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type JobAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *JobAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobAccessControlResponse) {
}

func (newState *JobAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState JobAccessControlResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(JobPermission{}),
	}
}

// ToObjectType returns the representation of JobAccessControlResponse in the Terraform plugin framework type
// system.
func (a JobAccessControlResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: JobPermission{}.ToObjectType(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type JobCluster struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	JobClusterKey types.String `tfsdk:"job_cluster_key" tf:""`
	// If new_cluster, a description of a cluster that is created for each task.
	NewCluster compute.ClusterSpec `tfsdk:"new_cluster" tf:"object"`
}

func (newState *JobCluster) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobCluster) {
}

func (newState *JobCluster) SyncEffectiveFieldsDuringRead(existingState JobCluster) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobCluster) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"new_cluster": reflect.TypeOf(compute.ClusterSpec{}),
	}
}

// ToObjectType returns the representation of JobCluster in the Terraform plugin framework type
// system.
func (a JobCluster) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_cluster_key": types.StringType,
			"new_cluster": basetypes.ListType{
				ElemType: compute_tf.ClusterSpec{}.ToObjectType(ctx),
			},
		},
	}
}

type JobCompliance struct {
	// Whether this job is in compliance with the latest version of its policy.
	IsCompliant types.Bool `tfsdk:"is_compliant" tf:"optional"`
	// Canonical unique identifier for a job.
	JobId types.Int64 `tfsdk:"job_id" tf:""`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	Violations types.Map `tfsdk:"violations" tf:"optional"`
}

func (newState *JobCompliance) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobCompliance) {
}

func (newState *JobCompliance) SyncEffectiveFieldsDuringRead(existingState JobCompliance) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobCompliance.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobCompliance) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"violations": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of JobCompliance in the Terraform plugin framework type
// system.
func (a JobCompliance) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_compliant": types.BoolType,
			"job_id":       types.Int64Type,
			"violations": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

type JobDeployment struct {
	// The kind of deployment that manages the job.
	//
	// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
	Kind types.String `tfsdk:"kind" tf:""`
	// Path of the file that contains deployment metadata.
	MetadataFilePath types.String `tfsdk:"metadata_file_path" tf:"optional"`
}

func (newState *JobDeployment) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobDeployment) {
}

func (newState *JobDeployment) SyncEffectiveFieldsDuringRead(existingState JobDeployment) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobDeployment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobDeployment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of JobDeployment in the Terraform plugin framework type
// system.
func (a JobDeployment) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"kind":               types.StringType,
			"metadata_file_path": types.StringType,
		},
	}
}

type JobEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	NoAlertForSkippedRuns types.Bool `tfsdk:"no_alert_for_skipped_runs" tf:"optional"`
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded types.List `tfsdk:"on_duration_warning_threshold_exceeded" tf:"optional"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure types.List `tfsdk:"on_failure" tf:"optional"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart types.List `tfsdk:"on_start" tf:"optional"`
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	OnStreamingBacklogExceeded types.List `tfsdk:"on_streaming_backlog_exceeded" tf:"optional"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess types.List `tfsdk:"on_success" tf:"optional"`
}

func (newState *JobEmailNotifications) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobEmailNotifications) {
}

func (newState *JobEmailNotifications) SyncEffectiveFieldsDuringRead(existingState JobEmailNotifications) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobEmailNotifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobEmailNotifications) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_duration_warning_threshold_exceeded": reflect.TypeOf(types.String{}),
		"on_failure":                             reflect.TypeOf(types.String{}),
		"on_start":                               reflect.TypeOf(types.String{}),
		"on_streaming_backlog_exceeded":          reflect.TypeOf(types.String{}),
		"on_success":                             reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of JobEmailNotifications in the Terraform plugin framework type
// system.
func (a JobEmailNotifications) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"no_alert_for_skipped_runs": types.BoolType,
			"on_duration_warning_threshold_exceeded": basetypes.ListType{
				ElemType: types.StringType,
			},
			"on_failure": basetypes.ListType{
				ElemType: types.StringType,
			},
			"on_start": basetypes.ListType{
				ElemType: types.StringType,
			},
			"on_streaming_backlog_exceeded": basetypes.ListType{
				ElemType: types.StringType,
			},
			"on_success": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type JobEnvironment struct {
	// The key of an environment. It has to be unique within a job.
	EnvironmentKey types.String `tfsdk:"environment_key" tf:""`
	// The environment entity used to preserve serverless environment side panel
	// and jobs' environment for non-notebook task. In this minimal environment
	// spec, only pip dependencies are supported.
	Spec compute.Environment `tfsdk:"spec" tf:"optional,object"`
}

func (newState *JobEnvironment) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobEnvironment) {
}

func (newState *JobEnvironment) SyncEffectiveFieldsDuringRead(existingState JobEnvironment) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobEnvironment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobEnvironment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec": reflect.TypeOf(compute.Environment{}),
	}
}

// ToObjectType returns the representation of JobEnvironment in the Terraform plugin framework type
// system.
func (a JobEnvironment) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"environment_key": types.StringType,
			"spec": basetypes.ListType{
				ElemType: compute_tf.Environment{}.ToObjectType(ctx),
			},
		},
	}
}

type JobNotificationSettings struct {
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns types.Bool `tfsdk:"no_alert_for_canceled_runs" tf:"optional"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns types.Bool `tfsdk:"no_alert_for_skipped_runs" tf:"optional"`
}

func (newState *JobNotificationSettings) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobNotificationSettings) {
}

func (newState *JobNotificationSettings) SyncEffectiveFieldsDuringRead(existingState JobNotificationSettings) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobNotificationSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobNotificationSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of JobNotificationSettings in the Terraform plugin framework type
// system.
func (a JobNotificationSettings) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"no_alert_for_canceled_runs": types.BoolType,
			"no_alert_for_skipped_runs":  types.BoolType,
		},
	}
}

type JobParameter struct {
	// The optional default value of the parameter
	Default types.String `tfsdk:"default" tf:"optional"`
	// The name of the parameter
	Name types.String `tfsdk:"name" tf:"optional"`
	// The value used in the run
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *JobParameter) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobParameter) {
}

func (newState *JobParameter) SyncEffectiveFieldsDuringRead(existingState JobParameter) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobParameter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobParameter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of JobParameter in the Terraform plugin framework type
// system.
func (a JobParameter) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default": types.StringType,
			"name":    types.StringType,
			"value":   types.StringType,
		},
	}
}

type JobParameterDefinition struct {
	// Default value of the parameter.
	Default types.String `tfsdk:"default" tf:""`
	// The name of the defined parameter. May only contain alphanumeric
	// characters, `_`, `-`, and `.`
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *JobParameterDefinition) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobParameterDefinition) {
}

func (newState *JobParameterDefinition) SyncEffectiveFieldsDuringRead(existingState JobParameterDefinition) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobParameterDefinition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobParameterDefinition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of JobParameterDefinition in the Terraform plugin framework type
// system.
func (a JobParameterDefinition) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default": types.StringType,
			"name":    types.StringType,
		},
	}
}

type JobPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *JobPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobPermission) {
}

func (newState *JobPermission) SyncEffectiveFieldsDuringRead(existingState JobPermission) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of JobPermission in the Terraform plugin framework type
// system.
func (a JobPermission) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

type JobPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *JobPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobPermissions) {
}

func (newState *JobPermissions) SyncEffectiveFieldsDuringRead(existingState JobPermissions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(JobAccessControlResponse{}),
	}
}

// ToObjectType returns the representation of JobPermissions in the Terraform plugin framework type
// system.
func (a JobPermissions) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: JobAccessControlResponse{}.ToObjectType(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type JobPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *JobPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobPermissionsDescription) {
}

func (newState *JobPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState JobPermissionsDescription) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of JobPermissionsDescription in the Terraform plugin framework type
// system.
func (a JobPermissionsDescription) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type JobPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The job for which to get or manage permissions.
	JobId types.String `tfsdk:"-"`
}

func (newState *JobPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobPermissionsRequest) {
}

func (newState *JobPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState JobPermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(JobAccessControlRequest{}),
	}
}

// ToObjectType returns the representation of JobPermissionsRequest in the Terraform plugin framework type
// system.
func (a JobPermissionsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: JobAccessControlRequest{}.ToObjectType(ctx),
			},
			"job_id": types.StringType,
		},
	}
}

// Write-only setting. Specifies the user, service principal or group that the
// job/pipeline runs as. If not specified, the job/pipeline runs as the user who
// created the job/pipeline.
//
// Either `user_name` or `service_principal_name` should be specified. If not,
// an error is thrown.
type JobRunAs struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// The email of an active workspace user. Non-admin users can only set this
	// field to their own email.
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *JobRunAs) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobRunAs) {
}

func (newState *JobRunAs) SyncEffectiveFieldsDuringRead(existingState JobRunAs) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobRunAs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobRunAs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of JobRunAs in the Terraform plugin framework type
// system.
func (a JobRunAs) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type JobSettings struct {
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id" tf:"optional"`
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous types.List `tfsdk:"continuous" tf:"optional,object"`
	// Deployment information for jobs managed by external sources.
	Deployment types.List `tfsdk:"deployment" tf:"optional,object"`
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode types.String `tfsdk:"edit_mode" tf:"optional"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications types.List `tfsdk:"email_notifications" tf:"optional,object"`
	// A list of task execution environment specifications that can be
	// referenced by serverless tasks of this job. An environment is required to
	// be present for serverless tasks. For serverless notebook tasks, the
	// environment is accessible in the notebook environment panel. For other
	// serverless tasks, the task environment is required to be specified using
	// environment_key in the task settings.
	Environments types.List `tfsdk:"environment" tf:"optional"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format types.String `tfsdk:"format" tf:"optional"`
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
	GitSource types.List `tfsdk:"git_source" tf:"optional,object"`
	// An optional set of health rules that can be defined for this job.
	Health types.List `tfsdk:"health" tf:"optional,object"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters types.List `tfsdk:"job_cluster" tf:"optional"`
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
	MaxConcurrentRuns types.Int64 `tfsdk:"max_concurrent_runs" tf:"optional"`
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings types.List `tfsdk:"notification_settings" tf:"optional,object"`
	// Job-level parameter definitions
	Parameters types.List `tfsdk:"parameter" tf:"optional"`
	// The queue settings of the job.
	Queue types.List `tfsdk:"queue" tf:"optional,object"`
	// Write-only setting. Specifies the user, service principal or group that
	// the job/pipeline runs as. If not specified, the job/pipeline runs as the
	// user who created the job/pipeline.
	//
	// Either `user_name` or `service_principal_name` should be specified. If
	// not, an error is thrown.
	RunAs types.List `tfsdk:"run_as" tf:"optional,object"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags types.Map `tfsdk:"tags" tf:"optional"`
	// A list of task specifications to be executed by this job.
	Tasks types.List `tfsdk:"task" tf:"optional"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds" tf:"optional"`
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger types.List `tfsdk:"trigger" tf:"optional,object"`
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications types.List `tfsdk:"webhook_notifications" tf:"optional,object"`
}

func (newState *JobSettings) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobSettings) {
}

func (newState *JobSettings) SyncEffectiveFieldsDuringRead(existingState JobSettings) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"continuous":            reflect.TypeOf(Continuous{}),
		"deployment":            reflect.TypeOf(JobDeployment{}),
		"email_notifications":   reflect.TypeOf(JobEmailNotifications{}),
		"environment":           reflect.TypeOf(JobEnvironment{}),
		"git_source":            reflect.TypeOf(GitSource{}),
		"health":                reflect.TypeOf(JobsHealthRules{}),
		"job_cluster":           reflect.TypeOf(JobCluster{}),
		"notification_settings": reflect.TypeOf(JobNotificationSettings{}),
		"parameter":             reflect.TypeOf(JobParameterDefinition{}),
		"queue":                 reflect.TypeOf(QueueSettings{}),
		"run_as":                reflect.TypeOf(JobRunAs{}),
		"schedule":              reflect.TypeOf(CronSchedule{}),
		"tags":                  reflect.TypeOf(types.String{}),
		"task":                  reflect.TypeOf(Task{}),
		"trigger":               reflect.TypeOf(TriggerSettings{}),
		"webhook_notifications": reflect.TypeOf(WebhookNotifications{}),
	}
}

// ToObjectType returns the representation of JobSettings in the Terraform plugin framework type
// system.
func (a JobSettings) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"continuous": basetypes.ListType{
				ElemType: Continuous{}.ToObjectType(ctx),
			},
			"deployment": basetypes.ListType{
				ElemType: JobDeployment{}.ToObjectType(ctx),
			},
			"description": types.StringType,
			"edit_mode":   types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: JobEmailNotifications{}.ToObjectType(ctx),
			},
			"environment": basetypes.ListType{
				ElemType: JobEnvironment{}.ToObjectType(ctx),
			},
			"format": types.StringType,
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.ToObjectType(ctx),
			},
			"health": basetypes.ListType{
				ElemType: JobsHealthRules{}.ToObjectType(ctx),
			},
			"job_cluster": basetypes.ListType{
				ElemType: JobCluster{}.ToObjectType(ctx),
			},
			"max_concurrent_runs": types.Int64Type,
			"name":                types.StringType,
			"notification_settings": basetypes.ListType{
				ElemType: JobNotificationSettings{}.ToObjectType(ctx),
			},
			"parameter": basetypes.ListType{
				ElemType: JobParameterDefinition{}.ToObjectType(ctx),
			},
			"queue": basetypes.ListType{
				ElemType: QueueSettings{}.ToObjectType(ctx),
			},
			"run_as": basetypes.ListType{
				ElemType: JobRunAs{}.ToObjectType(ctx),
			},
			"schedule": basetypes.ListType{
				ElemType: CronSchedule{}.ToObjectType(ctx),
			},
			"tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"task": basetypes.ListType{
				ElemType: Task{}.ToObjectType(ctx),
			},
			"timeout_seconds": types.Int64Type,
			"trigger": basetypes.ListType{
				ElemType: TriggerSettings{}.ToObjectType(ctx),
			},
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.ToObjectType(ctx),
			},
		},
	}
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
	DirtyState types.String `tfsdk:"dirty_state" tf:"optional"`
	// Name of the branch which the job is imported from.
	ImportFromGitBranch types.String `tfsdk:"import_from_git_branch" tf:""`
	// Path of the job YAML file that contains the job specification.
	JobConfigPath types.String `tfsdk:"job_config_path" tf:""`
}

func (newState *JobSource) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobSource) {
}

func (newState *JobSource) SyncEffectiveFieldsDuringRead(existingState JobSource) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobSource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of JobSource in the Terraform plugin framework type
// system.
func (a JobSource) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dirty_state":            types.StringType,
			"import_from_git_branch": types.StringType,
			"job_config_path":        types.StringType,
		},
	}
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
	Metric types.String `tfsdk:"metric" tf:""`
	// Specifies the operator used to compare the health metric value with the
	// specified threshold.
	Op types.String `tfsdk:"op" tf:""`
	// Specifies the threshold value that the health metric should obey to
	// satisfy the health rule.
	Value types.Int64 `tfsdk:"value" tf:""`
}

func (newState *JobsHealthRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobsHealthRule) {
}

func (newState *JobsHealthRule) SyncEffectiveFieldsDuringRead(existingState JobsHealthRule) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobsHealthRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobsHealthRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of JobsHealthRule in the Terraform plugin framework type
// system.
func (a JobsHealthRule) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metric": types.StringType,
			"op":     types.StringType,
			"value":  types.Int64Type,
		},
	}
}

// An optional set of health rules that can be defined for this job.
type JobsHealthRules struct {
	Rules types.List `tfsdk:"rules" tf:"optional"`
}

func (newState *JobsHealthRules) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobsHealthRules) {
}

func (newState *JobsHealthRules) SyncEffectiveFieldsDuringRead(existingState JobsHealthRules) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobsHealthRules.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobsHealthRules) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rules": reflect.TypeOf(JobsHealthRule{}),
	}
}

// ToObjectType returns the representation of JobsHealthRules in the Terraform plugin framework type
// system.
func (a JobsHealthRules) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"rules": basetypes.ListType{
				ElemType: JobsHealthRule{}.ToObjectType(ctx),
			},
		},
	}
}

type ListJobComplianceForPolicyResponse struct {
	// A list of jobs and their policy compliance statuses.
	Jobs types.List `tfsdk:"jobs" tf:"optional"`
	// This field represents the pagination token to retrieve the next page of
	// results. If this field is not in the response, it means no further
	// results for the request.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If this field is not in the response, it means no further
	// results for the request.
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`
}

func (newState *ListJobComplianceForPolicyResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListJobComplianceForPolicyResponse) {
}

func (newState *ListJobComplianceForPolicyResponse) SyncEffectiveFieldsDuringRead(existingState ListJobComplianceForPolicyResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListJobComplianceForPolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListJobComplianceForPolicyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"jobs": reflect.TypeOf(JobCompliance{}),
	}
}

// ToObjectType returns the representation of ListJobComplianceForPolicyResponse in the Terraform plugin framework type
// system.
func (a ListJobComplianceForPolicyResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"jobs": basetypes.ListType{
				ElemType: JobCompliance{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
		},
	}
}

// List job policy compliance
type ListJobComplianceRequest struct {
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token that can be used to navigate to the next page or previous
	// page as returned by `next_page_token` or `prev_page_token`.
	PageToken types.String `tfsdk:"-"`
	// Canonical unique identifier for the cluster policy.
	PolicyId types.String `tfsdk:"-"`
}

func (newState *ListJobComplianceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListJobComplianceRequest) {
}

func (newState *ListJobComplianceRequest) SyncEffectiveFieldsDuringRead(existingState ListJobComplianceRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListJobComplianceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListJobComplianceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListJobComplianceRequest in the Terraform plugin framework type
// system.
func (a ListJobComplianceRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"policy_id":  types.StringType,
		},
	}
}

// List jobs
type ListJobsRequest struct {
	// Whether to include task and cluster details in the response.
	ExpandTasks types.Bool `tfsdk:"-"`
	// The number of jobs to return. This value must be greater than 0 and less
	// or equal to 100. The default value is 20.
	Limit types.Int64 `tfsdk:"-"`
	// A filter on the list based on the exact (case insensitive) job name.
	Name types.String `tfsdk:"-"`
	// The offset of the first job to return, relative to the most recently
	// created job. Deprecated since June 2023. Use `page_token` to iterate
	// through the pages instead.
	Offset types.Int64 `tfsdk:"-"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of jobs respectively.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListJobsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListJobsRequest) {
}

func (newState *ListJobsRequest) SyncEffectiveFieldsDuringRead(existingState ListJobsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListJobsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListJobsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListJobsRequest in the Terraform plugin framework type
// system.
func (a ListJobsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"expand_tasks": types.BoolType,
			"limit":        types.Int64Type,
			"name":         types.StringType,
			"offset":       types.Int64Type,
			"page_token":   types.StringType,
		},
	}
}

// List of jobs was retrieved successfully.
type ListJobsResponse struct {
	// If true, additional jobs matching the provided filter are available for
	// listing.
	HasMore types.Bool `tfsdk:"has_more" tf:"optional"`
	// The list of jobs. Only included in the response if there are jobs to
	// list.
	Jobs types.List `tfsdk:"jobs" tf:"optional"`
	// A token that can be used to list the next page of jobs (if applicable).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// A token that can be used to list the previous page of jobs (if
	// applicable).
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`
}

func (newState *ListJobsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListJobsResponse) {
}

func (newState *ListJobsResponse) SyncEffectiveFieldsDuringRead(existingState ListJobsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListJobsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListJobsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"jobs": reflect.TypeOf(BaseJob{}),
	}
}

// ToObjectType returns the representation of ListJobsResponse in the Terraform plugin framework type
// system.
func (a ListJobsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"has_more": types.BoolType,
			"jobs": basetypes.ListType{
				ElemType: BaseJob{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
		},
	}
}

// List job runs
type ListRunsRequest struct {
	// If active_only is `true`, only active runs are included in the results;
	// otherwise, lists both active and completed runs. An active run is a run
	// in the `QUEUED`, `PENDING`, `RUNNING`, or `TERMINATING`. This field
	// cannot be `true` when completed_only is `true`.
	ActiveOnly types.Bool `tfsdk:"-"`
	// If completed_only is `true`, only completed runs are included in the
	// results; otherwise, lists both active and completed runs. This field
	// cannot be `true` when active_only is `true`.
	CompletedOnly types.Bool `tfsdk:"-"`
	// Whether to include task and cluster details in the response.
	ExpandTasks types.Bool `tfsdk:"-"`
	// The job for which to list runs. If omitted, the Jobs service lists runs
	// from all jobs.
	JobId types.Int64 `tfsdk:"-"`
	// The number of runs to return. This value must be greater than 0 and less
	// than 25. The default value is 20. If a request specifies a limit of 0,
	// the service instead uses the maximum limit.
	Limit types.Int64 `tfsdk:"-"`
	// The offset of the first run to return, relative to the most recent run.
	// Deprecated since June 2023. Use `page_token` to iterate through the pages
	// instead.
	Offset types.Int64 `tfsdk:"-"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of runs respectively.
	PageToken types.String `tfsdk:"-"`
	// The type of runs to return. For a description of run types, see
	// :method:jobs/getRun.
	RunType types.String `tfsdk:"-"`
	// Show runs that started _at or after_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_to_ to filter
	// by a time range.
	StartTimeFrom types.Int64 `tfsdk:"-"`
	// Show runs that started _at or before_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_from_ to
	// filter by a time range.
	StartTimeTo types.Int64 `tfsdk:"-"`
}

func (newState *ListRunsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRunsRequest) {
}

func (newState *ListRunsRequest) SyncEffectiveFieldsDuringRead(existingState ListRunsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRunsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRunsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListRunsRequest in the Terraform plugin framework type
// system.
func (a ListRunsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active_only":     types.BoolType,
			"completed_only":  types.BoolType,
			"expand_tasks":    types.BoolType,
			"job_id":          types.Int64Type,
			"limit":           types.Int64Type,
			"offset":          types.Int64Type,
			"page_token":      types.StringType,
			"run_type":        types.StringType,
			"start_time_from": types.Int64Type,
			"start_time_to":   types.Int64Type,
		},
	}
}

// List of runs was retrieved successfully.
type ListRunsResponse struct {
	// If true, additional runs matching the provided filter are available for
	// listing.
	HasMore types.Bool `tfsdk:"has_more" tf:"optional"`
	// A token that can be used to list the next page of runs (if applicable).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// A token that can be used to list the previous page of runs (if
	// applicable).
	PrevPageToken types.String `tfsdk:"prev_page_token" tf:"optional"`
	// A list of runs, from most recently started to least. Only included in the
	// response if there are runs to list.
	Runs types.List `tfsdk:"runs" tf:"optional"`
}

func (newState *ListRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRunsResponse) {
}

func (newState *ListRunsResponse) SyncEffectiveFieldsDuringRead(existingState ListRunsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRunsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(BaseRun{}),
	}
}

// ToObjectType returns the representation of ListRunsResponse in the Terraform plugin framework type
// system.
func (a ListRunsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"has_more":        types.BoolType,
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
			"runs": basetypes.ListType{
				ElemType: BaseRun{}.ToObjectType(ctx),
			},
		},
	}
}

type NotebookOutput struct {
	// The value passed to
	// [dbutils.notebook.exit()](/notebooks/notebook-workflows.html#notebook-workflows-exit).
	// Databricks restricts this API to return the first 5 MB of the value. For
	// a larger result, your job can store the results in a cloud storage
	// service. This field is absent if `dbutils.notebook.exit()` was never
	// called.
	Result types.String `tfsdk:"result" tf:"optional"`
	// Whether or not the result was truncated.
	Truncated types.Bool `tfsdk:"truncated" tf:"optional"`
}

func (newState *NotebookOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan NotebookOutput) {
}

func (newState *NotebookOutput) SyncEffectiveFieldsDuringRead(existingState NotebookOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotebookOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NotebookOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of NotebookOutput in the Terraform plugin framework type
// system.
func (a NotebookOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"result":    types.StringType,
			"truncated": types.BoolType,
		},
	}
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
	BaseParameters types.Map `tfsdk:"base_parameters" tf:"optional"`
	// The path of the notebook to be run in the Databricks workspace or remote
	// repository. For notebooks stored in the Databricks workspace, the path
	// must be absolute and begin with a slash. For notebooks stored in a remote
	// repository, the path must be relative. This field is required.
	NotebookPath types.String `tfsdk:"notebook_path" tf:""`
	// Optional location type of the notebook. When set to `WORKSPACE`, the
	// notebook will be retrieved from the local Databricks workspace. When set
	// to `GIT`, the notebook will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`:
	// Notebook is located in Databricks workspace. * `GIT`: Notebook is located
	// in cloud Git provider.
	Source types.String `tfsdk:"source" tf:"optional"`
	// Optional `warehouse_id` to run the notebook on a SQL warehouse. Classic
	// SQL warehouses are NOT supported, please use serverless or pro SQL
	// warehouses.
	//
	// Note that SQL warehouses only support SQL cells; if the notebook contains
	// non-SQL cells, the run will fail.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *NotebookTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan NotebookTask) {
}

func (newState *NotebookTask) SyncEffectiveFieldsDuringRead(existingState NotebookTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotebookTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NotebookTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"base_parameters": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of NotebookTask in the Terraform plugin framework type
// system.
func (a NotebookTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"base_parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"notebook_path": types.StringType,
			"source":        types.StringType,
			"warehouse_id":  types.StringType,
		},
	}
}

type PeriodicTriggerConfiguration struct {
	// The interval at which the trigger should run.
	Interval types.Int64 `tfsdk:"interval" tf:""`
	// The unit of time for the interval.
	Unit types.String `tfsdk:"unit" tf:""`
}

func (newState *PeriodicTriggerConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan PeriodicTriggerConfiguration) {
}

func (newState *PeriodicTriggerConfiguration) SyncEffectiveFieldsDuringRead(existingState PeriodicTriggerConfiguration) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PeriodicTriggerConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PeriodicTriggerConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PeriodicTriggerConfiguration in the Terraform plugin framework type
// system.
func (a PeriodicTriggerConfiguration) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"interval": types.Int64Type,
			"unit":     types.StringType,
		},
	}
}

type PipelineParams struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh types.Bool `tfsdk:"full_refresh" tf:"optional"`
}

func (newState *PipelineParams) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineParams) {
}

func (newState *PipelineParams) SyncEffectiveFieldsDuringRead(existingState PipelineParams) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineParams.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineParams) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PipelineParams in the Terraform plugin framework type
// system.
func (a PipelineParams) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_refresh": types.BoolType,
		},
	}
}

type PipelineTask struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh types.Bool `tfsdk:"full_refresh" tf:"optional"`
	// The full name of the pipeline task to execute.
	PipelineId types.String `tfsdk:"pipeline_id" tf:""`
}

func (newState *PipelineTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineTask) {
}

func (newState *PipelineTask) SyncEffectiveFieldsDuringRead(existingState PipelineTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PipelineTask in the Terraform plugin framework type
// system.
func (a PipelineTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_refresh": types.BoolType,
			"pipeline_id":  types.StringType,
		},
	}
}

type PythonWheelTask struct {
	// Named entry point to use, if it does not exist in the metadata of the
	// package it executes the function from the package directly using
	// `$packageName.$entryPoint()`
	EntryPoint types.String `tfsdk:"entry_point" tf:""`
	// Command-line parameters passed to Python wheel task in the form of
	// `["--name=task", "--data=dbfs:/path/to/data.json"]`. Leave it empty if
	// `parameters` is not null.
	NamedParameters types.Map `tfsdk:"named_parameters" tf:"optional"`
	// Name of the package to execute
	PackageName types.String `tfsdk:"package_name" tf:""`
	// Command-line parameters passed to Python wheel task. Leave it empty if
	// `named_parameters` is not null.
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
}

func (newState *PythonWheelTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan PythonWheelTask) {
}

func (newState *PythonWheelTask) SyncEffectiveFieldsDuringRead(existingState PythonWheelTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PythonWheelTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PythonWheelTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"named_parameters": reflect.TypeOf(types.String{}),
		"parameters":       reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of PythonWheelTask in the Terraform plugin framework type
// system.
func (a PythonWheelTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entry_point": types.StringType,
			"named_parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"package_name": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type QueueDetails struct {
	// The reason for queuing the run. * `ACTIVE_RUNS_LIMIT_REACHED`: The run
	// was queued due to reaching the workspace limit of active task runs. *
	// `MAX_CONCURRENT_RUNS_REACHED`: The run was queued due to reaching the
	// per-job limit of concurrent job runs. *
	// `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`: The run was queued due to reaching
	// the workspace limit of active run job tasks.
	Code types.String `tfsdk:"code" tf:"optional"`
	// A descriptive message with the queuing details. This field is
	// unstructured, and its exact format is subject to change.
	Message types.String `tfsdk:"message" tf:"optional"`
}

func (newState *QueueDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueueDetails) {
}

func (newState *QueueDetails) SyncEffectiveFieldsDuringRead(existingState QueueDetails) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueueDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueueDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of QueueDetails in the Terraform plugin framework type
// system.
func (a QueueDetails) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"code":    types.StringType,
			"message": types.StringType,
		},
	}
}

type QueueSettings struct {
	// If true, enable queueing for the job. This is a required field.
	Enabled types.Bool `tfsdk:"enabled" tf:""`
}

func (newState *QueueSettings) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueueSettings) {
}

func (newState *QueueSettings) SyncEffectiveFieldsDuringRead(existingState QueueSettings) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueueSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueueSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of QueueSettings in the Terraform plugin framework type
// system.
func (a QueueSettings) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

type RepairHistoryItem struct {
	// The end time of the (repaired) run.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// The ID of the repair. Only returned for the items that represent a repair
	// in `repair_history`.
	Id types.Int64 `tfsdk:"id" tf:"optional"`
	// The start time of the (repaired) run.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Deprecated. Please use the `status` field instead.
	State types.List `tfsdk:"state" tf:"optional,object"`
	// The current status of the run
	Status types.List `tfsdk:"status" tf:"optional,object"`
	// The run IDs of the task runs that ran as part of this repair history
	// item.
	TaskRunIds types.List `tfsdk:"task_run_ids" tf:"optional"`
	// The repair history item type. Indicates whether a run is the original run
	// or a repair run.
	Type types.String `tfsdk:"type" tf:"optional"`
}

func (newState *RepairHistoryItem) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepairHistoryItem) {
}

func (newState *RepairHistoryItem) SyncEffectiveFieldsDuringRead(existingState RepairHistoryItem) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepairHistoryItem.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepairHistoryItem) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"state":        reflect.TypeOf(RunState{}),
		"status":       reflect.TypeOf(RunStatus{}),
		"task_run_ids": reflect.TypeOf(types.Int64{}),
	}
}

// ToObjectType returns the representation of RepairHistoryItem in the Terraform plugin framework type
// system.
func (a RepairHistoryItem) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time":   types.Int64Type,
			"id":         types.Int64Type,
			"start_time": types.Int64Type,
			"state": basetypes.ListType{
				ElemType: RunState{}.ToObjectType(ctx),
			},
			"status": basetypes.ListType{
				ElemType: RunStatus{}.ToObjectType(ctx),
			},
			"task_run_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"type": types.StringType,
		},
	}
}

type RepairRun struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands types.List `tfsdk:"dbt_commands" tf:"optional"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams types.List `tfsdk:"jar_params" tf:"optional"`
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters types.Map `tfsdk:"job_parameters" tf:"optional"`
	// The ID of the latest repair. This parameter is not required when
	// repairing a run for the first time, but must be provided on subsequent
	// requests to repair the same run.
	LatestRepairId types.Int64 `tfsdk:"latest_repair_id" tf:"optional"`
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
	NotebookParams types.Map `tfsdk:"notebook_params" tf:"optional"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams types.List `tfsdk:"pipeline_params" tf:"optional,object"`

	PythonNamedParams types.Map `tfsdk:"python_named_params" tf:"optional"`
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
	PythonParams types.List `tfsdk:"python_params" tf:"optional"`
	// If true, repair all failed tasks. Only one of `rerun_tasks` or
	// `rerun_all_failed_tasks` can be used.
	RerunAllFailedTasks types.Bool `tfsdk:"rerun_all_failed_tasks" tf:"optional"`
	// If true, repair all tasks that depend on the tasks in `rerun_tasks`, even
	// if they were previously successful. Can be also used in combination with
	// `rerun_all_failed_tasks`.
	RerunDependentTasks types.Bool `tfsdk:"rerun_dependent_tasks" tf:"optional"`
	// The task keys of the task runs to repair.
	RerunTasks types.List `tfsdk:"rerun_tasks" tf:"optional"`
	// The job run ID of the run to repair. The run must not be in progress.
	RunId types.Int64 `tfsdk:"run_id" tf:""`
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
	SparkSubmitParams types.List `tfsdk:"spark_submit_params" tf:"optional"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams types.Map `tfsdk:"sql_params" tf:"optional"`
}

func (newState *RepairRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepairRun) {
}

func (newState *RepairRun) SyncEffectiveFieldsDuringRead(existingState RepairRun) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepairRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepairRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbt_commands":        reflect.TypeOf(types.String{}),
		"jar_params":          reflect.TypeOf(types.String{}),
		"job_parameters":      reflect.TypeOf(types.String{}),
		"notebook_params":     reflect.TypeOf(types.String{}),
		"pipeline_params":     reflect.TypeOf(PipelineParams{}),
		"python_named_params": reflect.TypeOf(types.String{}),
		"python_params":       reflect.TypeOf(types.String{}),
		"rerun_tasks":         reflect.TypeOf(types.String{}),
		"spark_submit_params": reflect.TypeOf(types.String{}),
		"sql_params":          reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of RepairRun in the Terraform plugin framework type
// system.
func (a RepairRun) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbt_commands": basetypes.ListType{
				ElemType: types.StringType,
			},
			"jar_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"job_parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"latest_repair_id": types.Int64Type,
			"notebook_params": basetypes.MapType{
				ElemType: types.StringType,
			},
			"pipeline_params": basetypes.ListType{
				ElemType: PipelineParams{}.ToObjectType(ctx),
			},
			"python_named_params": basetypes.MapType{
				ElemType: types.StringType,
			},
			"python_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"rerun_all_failed_tasks": types.BoolType,
			"rerun_dependent_tasks":  types.BoolType,
			"rerun_tasks": basetypes.ListType{
				ElemType: types.StringType,
			},
			"run_id": types.Int64Type,
			"spark_submit_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"sql_params": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// Run repair was initiated.
type RepairRunResponse struct {
	// The ID of the repair. Must be provided in subsequent repairs using the
	// `latest_repair_id` field to ensure sequential repairs.
	RepairId types.Int64 `tfsdk:"repair_id" tf:"optional"`
}

func (newState *RepairRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepairRunResponse) {
}

func (newState *RepairRunResponse) SyncEffectiveFieldsDuringRead(existingState RepairRunResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepairRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepairRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of RepairRunResponse in the Terraform plugin framework type
// system.
func (a RepairRunResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repair_id": types.Int64Type,
		},
	}
}

type ResetJob struct {
	// The canonical identifier of the job to reset. This field is required.
	JobId types.Int64 `tfsdk:"job_id" tf:""`
	// The new settings of the job. These settings completely replace the old
	// settings.
	//
	// Changes to the field `JobBaseSettings.timeout_seconds` are applied to
	// active runs. Changes to other fields are applied to future runs only.
	NewSettings types.List `tfsdk:"new_settings" tf:"object"`
}

func (newState *ResetJob) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResetJob) {
}

func (newState *ResetJob) SyncEffectiveFieldsDuringRead(existingState ResetJob) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResetJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResetJob) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"new_settings": reflect.TypeOf(JobSettings{}),
	}
}

// ToObjectType returns the representation of ResetJob in the Terraform plugin framework type
// system.
func (a ResetJob) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.Int64Type,
			"new_settings": basetypes.ListType{
				ElemType: JobSettings{}.ToObjectType(ctx),
			},
		},
	}
}

type ResetResponse struct {
}

func (newState *ResetResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResetResponse) {
}

func (newState *ResetResponse) SyncEffectiveFieldsDuringRead(existingState ResetResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResetResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ResetResponse in the Terraform plugin framework type
// system.
func (a ResetResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ResolvedConditionTaskValues struct {
	Left types.String `tfsdk:"left" tf:"optional"`

	Right types.String `tfsdk:"right" tf:"optional"`
}

func (newState *ResolvedConditionTaskValues) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedConditionTaskValues) {
}

func (newState *ResolvedConditionTaskValues) SyncEffectiveFieldsDuringRead(existingState ResolvedConditionTaskValues) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedConditionTaskValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedConditionTaskValues) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ResolvedConditionTaskValues in the Terraform plugin framework type
// system.
func (a ResolvedConditionTaskValues) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"left":  types.StringType,
			"right": types.StringType,
		},
	}
}

type ResolvedDbtTaskValues struct {
	Commands types.List `tfsdk:"commands" tf:"optional"`
}

func (newState *ResolvedDbtTaskValues) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedDbtTaskValues) {
}

func (newState *ResolvedDbtTaskValues) SyncEffectiveFieldsDuringRead(existingState ResolvedDbtTaskValues) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedDbtTaskValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedDbtTaskValues) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"commands": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ResolvedDbtTaskValues in the Terraform plugin framework type
// system.
func (a ResolvedDbtTaskValues) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"commands": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type ResolvedNotebookTaskValues struct {
	BaseParameters types.Map `tfsdk:"base_parameters" tf:"optional"`
}

func (newState *ResolvedNotebookTaskValues) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedNotebookTaskValues) {
}

func (newState *ResolvedNotebookTaskValues) SyncEffectiveFieldsDuringRead(existingState ResolvedNotebookTaskValues) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedNotebookTaskValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedNotebookTaskValues) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"base_parameters": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ResolvedNotebookTaskValues in the Terraform plugin framework type
// system.
func (a ResolvedNotebookTaskValues) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"base_parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

type ResolvedParamPairValues struct {
	Parameters types.Map `tfsdk:"parameters" tf:"optional"`
}

func (newState *ResolvedParamPairValues) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedParamPairValues) {
}

func (newState *ResolvedParamPairValues) SyncEffectiveFieldsDuringRead(existingState ResolvedParamPairValues) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedParamPairValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedParamPairValues) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ResolvedParamPairValues in the Terraform plugin framework type
// system.
func (a ResolvedParamPairValues) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

type ResolvedPythonWheelTaskValues struct {
	NamedParameters types.Map `tfsdk:"named_parameters" tf:"optional"`

	Parameters types.List `tfsdk:"parameters" tf:"optional"`
}

func (newState *ResolvedPythonWheelTaskValues) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedPythonWheelTaskValues) {
}

func (newState *ResolvedPythonWheelTaskValues) SyncEffectiveFieldsDuringRead(existingState ResolvedPythonWheelTaskValues) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedPythonWheelTaskValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedPythonWheelTaskValues) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"named_parameters": reflect.TypeOf(types.String{}),
		"parameters":       reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ResolvedPythonWheelTaskValues in the Terraform plugin framework type
// system.
func (a ResolvedPythonWheelTaskValues) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"named_parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"parameters": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type ResolvedRunJobTaskValues struct {
	JobParameters types.Map `tfsdk:"job_parameters" tf:"optional"`

	Parameters types.Map `tfsdk:"parameters" tf:"optional"`
}

func (newState *ResolvedRunJobTaskValues) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedRunJobTaskValues) {
}

func (newState *ResolvedRunJobTaskValues) SyncEffectiveFieldsDuringRead(existingState ResolvedRunJobTaskValues) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedRunJobTaskValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedRunJobTaskValues) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"job_parameters": reflect.TypeOf(types.String{}),
		"parameters":     reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ResolvedRunJobTaskValues in the Terraform plugin framework type
// system.
func (a ResolvedRunJobTaskValues) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

type ResolvedStringParamsValues struct {
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
}

func (newState *ResolvedStringParamsValues) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedStringParamsValues) {
}

func (newState *ResolvedStringParamsValues) SyncEffectiveFieldsDuringRead(existingState ResolvedStringParamsValues) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedStringParamsValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedStringParamsValues) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ResolvedStringParamsValues in the Terraform plugin framework type
// system.
func (a ResolvedStringParamsValues) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parameters": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type ResolvedValues struct {
	ConditionTask types.List `tfsdk:"condition_task" tf:"optional,object"`

	DbtTask types.List `tfsdk:"dbt_task" tf:"optional,object"`

	NotebookTask types.List `tfsdk:"notebook_task" tf:"optional,object"`

	PythonWheelTask types.List `tfsdk:"python_wheel_task" tf:"optional,object"`

	RunJobTask types.List `tfsdk:"run_job_task" tf:"optional,object"`

	SimulationTask types.List `tfsdk:"simulation_task" tf:"optional,object"`

	SparkJarTask types.List `tfsdk:"spark_jar_task" tf:"optional,object"`

	SparkPythonTask types.List `tfsdk:"spark_python_task" tf:"optional,object"`

	SparkSubmitTask types.List `tfsdk:"spark_submit_task" tf:"optional,object"`

	SqlTask types.List `tfsdk:"sql_task" tf:"optional,object"`
}

func (newState *ResolvedValues) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedValues) {
}

func (newState *ResolvedValues) SyncEffectiveFieldsDuringRead(existingState ResolvedValues) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedValues) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition_task":    reflect.TypeOf(ResolvedConditionTaskValues{}),
		"dbt_task":          reflect.TypeOf(ResolvedDbtTaskValues{}),
		"notebook_task":     reflect.TypeOf(ResolvedNotebookTaskValues{}),
		"python_wheel_task": reflect.TypeOf(ResolvedPythonWheelTaskValues{}),
		"run_job_task":      reflect.TypeOf(ResolvedRunJobTaskValues{}),
		"simulation_task":   reflect.TypeOf(ResolvedParamPairValues{}),
		"spark_jar_task":    reflect.TypeOf(ResolvedStringParamsValues{}),
		"spark_python_task": reflect.TypeOf(ResolvedStringParamsValues{}),
		"spark_submit_task": reflect.TypeOf(ResolvedStringParamsValues{}),
		"sql_task":          reflect.TypeOf(ResolvedParamPairValues{}),
	}
}

// ToObjectType returns the representation of ResolvedValues in the Terraform plugin framework type
// system.
func (a ResolvedValues) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition_task": basetypes.ListType{
				ElemType: ResolvedConditionTaskValues{}.ToObjectType(ctx),
			},
			"dbt_task": basetypes.ListType{
				ElemType: ResolvedDbtTaskValues{}.ToObjectType(ctx),
			},
			"notebook_task": basetypes.ListType{
				ElemType: ResolvedNotebookTaskValues{}.ToObjectType(ctx),
			},
			"python_wheel_task": basetypes.ListType{
				ElemType: ResolvedPythonWheelTaskValues{}.ToObjectType(ctx),
			},
			"run_job_task": basetypes.ListType{
				ElemType: ResolvedRunJobTaskValues{}.ToObjectType(ctx),
			},
			"simulation_task": basetypes.ListType{
				ElemType: ResolvedParamPairValues{}.ToObjectType(ctx),
			},
			"spark_jar_task": basetypes.ListType{
				ElemType: ResolvedStringParamsValues{}.ToObjectType(ctx),
			},
			"spark_python_task": basetypes.ListType{
				ElemType: ResolvedStringParamsValues{}.ToObjectType(ctx),
			},
			"spark_submit_task": basetypes.ListType{
				ElemType: ResolvedStringParamsValues{}.ToObjectType(ctx),
			},
			"sql_task": basetypes.ListType{
				ElemType: ResolvedParamPairValues{}.ToObjectType(ctx),
			},
		},
	}
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
	AttemptNumber types.Int64 `tfsdk:"attempt_number" tf:"optional"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration types.Int64 `tfsdk:"cleanup_duration" tf:"optional"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance types.List `tfsdk:"cluster_instance" tf:"optional,object"`
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec types.List `tfsdk:"cluster_spec" tf:"optional,object"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// Description of the run
	Description types.String `tfsdk:"description" tf:"optional"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration types.Int64 `tfsdk:"execution_duration" tf:"optional"`
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
	GitSource types.List `tfsdk:"git_source" tf:"optional,object"`
	// Only populated by for-each iterations. The parent for-each task is
	// located in tasks array.
	Iterations types.List `tfsdk:"iterations" tf:"optional"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters types.List `tfsdk:"job_clusters" tf:"optional"`
	// The canonical identifier of the job that contains this run.
	JobId types.Int64 `tfsdk:"job_id" tf:"optional"`
	// Job-level parameters used in the run
	JobParameters types.List `tfsdk:"job_parameters" tf:"optional"`
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	JobRunId types.Int64 `tfsdk:"job_run_id" tf:"optional"`
	// A token that can be used to list the next page of sub-resources.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob types.Int64 `tfsdk:"number_in_job" tf:"optional"`
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId types.Int64 `tfsdk:"original_attempt_run_id" tf:"optional"`
	// The parameters used for this run.
	OverridingParameters types.List `tfsdk:"overriding_parameters" tf:"optional,object"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration types.Int64 `tfsdk:"queue_duration" tf:"optional"`
	// The repair history of the run.
	RepairHistory types.List `tfsdk:"repair_history" tf:"optional"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration types.Int64 `tfsdk:"run_duration" tf:"optional"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName types.String `tfsdk:"run_name" tf:"optional"`
	// The URL to the detail page of the run.
	RunPageUrl types.String `tfsdk:"run_page_url" tf:"optional"`
	// The type of a run. * `JOB_RUN`: Normal job run. A run created with
	// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
	// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
	// :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	RunType types.String `tfsdk:"run_type" tf:"optional"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration types.Int64 `tfsdk:"setup_duration" tf:"optional"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Deprecated. Please use the `status` field instead.
	State types.List `tfsdk:"state" tf:"optional,object"`
	// The current status of the run
	Status types.List `tfsdk:"status" tf:"optional,object"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls.
	Tasks types.List `tfsdk:"tasks" tf:"optional"`
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
	Trigger types.String `tfsdk:"trigger" tf:"optional"`
	// Additional details about what triggered the run
	TriggerInfo types.List `tfsdk:"trigger_info" tf:"optional,object"`
}

func (newState *Run) SyncEffectiveFieldsDuringCreateOrUpdate(plan Run) {
}

func (newState *Run) SyncEffectiveFieldsDuringRead(existingState Run) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Run.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Run) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster_instance":      reflect.TypeOf(ClusterInstance{}),
		"cluster_spec":          reflect.TypeOf(ClusterSpec{}),
		"git_source":            reflect.TypeOf(GitSource{}),
		"iterations":            reflect.TypeOf(RunTask{}),
		"job_clusters":          reflect.TypeOf(JobCluster{}),
		"job_parameters":        reflect.TypeOf(JobParameter{}),
		"overriding_parameters": reflect.TypeOf(RunParameters{}),
		"repair_history":        reflect.TypeOf(RepairHistoryItem{}),
		"schedule":              reflect.TypeOf(CronSchedule{}),
		"state":                 reflect.TypeOf(RunState{}),
		"status":                reflect.TypeOf(RunStatus{}),
		"tasks":                 reflect.TypeOf(RunTask{}),
		"trigger_info":          reflect.TypeOf(TriggerInfo{}),
	}
}

// ToObjectType returns the representation of Run in the Terraform plugin framework type
// system.
func (a Run) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attempt_number":   types.Int64Type,
			"cleanup_duration": types.Int64Type,
			"cluster_instance": basetypes.ListType{
				ElemType: ClusterInstance{}.ToObjectType(ctx),
			},
			"cluster_spec": basetypes.ListType{
				ElemType: ClusterSpec{}.ToObjectType(ctx),
			},
			"creator_user_name":  types.StringType,
			"description":        types.StringType,
			"end_time":           types.Int64Type,
			"execution_duration": types.Int64Type,
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.ToObjectType(ctx),
			},
			"iterations": basetypes.ListType{
				ElemType: RunTask{}.ToObjectType(ctx),
			},
			"job_clusters": basetypes.ListType{
				ElemType: JobCluster{}.ToObjectType(ctx),
			},
			"job_id": types.Int64Type,
			"job_parameters": basetypes.ListType{
				ElemType: JobParameter{}.ToObjectType(ctx),
			},
			"job_run_id":              types.Int64Type,
			"next_page_token":         types.StringType,
			"number_in_job":           types.Int64Type,
			"original_attempt_run_id": types.Int64Type,
			"overriding_parameters": basetypes.ListType{
				ElemType: RunParameters{}.ToObjectType(ctx),
			},
			"queue_duration": types.Int64Type,
			"repair_history": basetypes.ListType{
				ElemType: RepairHistoryItem{}.ToObjectType(ctx),
			},
			"run_duration": types.Int64Type,
			"run_id":       types.Int64Type,
			"run_name":     types.StringType,
			"run_page_url": types.StringType,
			"run_type":     types.StringType,
			"schedule": basetypes.ListType{
				ElemType: CronSchedule{}.ToObjectType(ctx),
			},
			"setup_duration": types.Int64Type,
			"start_time":     types.Int64Type,
			"state": basetypes.ListType{
				ElemType: RunState{}.ToObjectType(ctx),
			},
			"status": basetypes.ListType{
				ElemType: RunStatus{}.ToObjectType(ctx),
			},
			"tasks": basetypes.ListType{
				ElemType: RunTask{}.ToObjectType(ctx),
			},
			"trigger": types.StringType,
			"trigger_info": basetypes.ListType{
				ElemType: TriggerInfo{}.ToObjectType(ctx),
			},
		},
	}
}

type RunConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	Left types.String `tfsdk:"left" tf:""`
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
	Op types.String `tfsdk:"op" tf:""`
	// The condition expression evaluation result. Filled in if the task was
	// successfully completed. Can be `"true"` or `"false"`
	Outcome types.String `tfsdk:"outcome" tf:"optional"`
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right types.String `tfsdk:"right" tf:""`
}

func (newState *RunConditionTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunConditionTask) {
}

func (newState *RunConditionTask) SyncEffectiveFieldsDuringRead(existingState RunConditionTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunConditionTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunConditionTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of RunConditionTask in the Terraform plugin framework type
// system.
func (a RunConditionTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"left":    types.StringType,
			"op":      types.StringType,
			"outcome": types.StringType,
			"right":   types.StringType,
		},
	}
}

type RunForEachTask struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	Concurrency types.Int64 `tfsdk:"concurrency" tf:"optional"`
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs types.String `tfsdk:"inputs" tf:""`
	// Read only field. Populated for GetRun and ListRuns RPC calls and stores
	// the execution stats of an For each task
	Stats types.List `tfsdk:"stats" tf:"optional,object"`
	// Configuration for the task that will be run for each element in the array
	Task types.List `tfsdk:"task" tf:"object"`
}

func (newState *RunForEachTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunForEachTask) {
}

func (newState *RunForEachTask) SyncEffectiveFieldsDuringRead(existingState RunForEachTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunForEachTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunForEachTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stats": reflect.TypeOf(ForEachStats{}),
		"task":  reflect.TypeOf(Task{}),
	}
}

// ToObjectType returns the representation of RunForEachTask in the Terraform plugin framework type
// system.
func (a RunForEachTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"concurrency": types.Int64Type,
			"inputs":      types.StringType,
			"stats": basetypes.ListType{
				ElemType: ForEachStats{}.ToObjectType(ctx),
			},
			"task": basetypes.ListType{
				ElemType: Task{}.ToObjectType(ctx),
			},
		},
	}
}

type RunJobOutput struct {
	// The run id of the triggered job run
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
}

func (newState *RunJobOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunJobOutput) {
}

func (newState *RunJobOutput) SyncEffectiveFieldsDuringRead(existingState RunJobOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunJobOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunJobOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of RunJobOutput in the Terraform plugin framework type
// system.
func (a RunJobOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

type RunJobTask struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands types.List `tfsdk:"dbt_commands" tf:"optional"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams types.List `tfsdk:"jar_params" tf:"optional"`
	// ID of the job to trigger.
	JobId types.Int64 `tfsdk:"job_id" tf:""`
	// Job-level parameters used to trigger the job.
	JobParameters types.Map `tfsdk:"job_parameters" tf:"optional"`
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
	NotebookParams types.Map `tfsdk:"notebook_params" tf:"optional"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams types.List `tfsdk:"pipeline_params" tf:"optional,object"`

	PythonNamedParams types.Map `tfsdk:"python_named_params" tf:"optional"`
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
	PythonParams types.List `tfsdk:"python_params" tf:"optional"`
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
	SparkSubmitParams types.List `tfsdk:"spark_submit_params" tf:"optional"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams types.Map `tfsdk:"sql_params" tf:"optional"`
}

func (newState *RunJobTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunJobTask) {
}

func (newState *RunJobTask) SyncEffectiveFieldsDuringRead(existingState RunJobTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunJobTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunJobTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbt_commands":        reflect.TypeOf(types.String{}),
		"jar_params":          reflect.TypeOf(types.String{}),
		"job_parameters":      reflect.TypeOf(types.String{}),
		"notebook_params":     reflect.TypeOf(types.String{}),
		"pipeline_params":     reflect.TypeOf(PipelineParams{}),
		"python_named_params": reflect.TypeOf(types.String{}),
		"python_params":       reflect.TypeOf(types.String{}),
		"spark_submit_params": reflect.TypeOf(types.String{}),
		"sql_params":          reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of RunJobTask in the Terraform plugin framework type
// system.
func (a RunJobTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbt_commands": basetypes.ListType{
				ElemType: types.StringType,
			},
			"jar_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"job_id": types.Int64Type,
			"job_parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"notebook_params": basetypes.MapType{
				ElemType: types.StringType,
			},
			"pipeline_params": basetypes.ListType{
				ElemType: PipelineParams{}.ToObjectType(ctx),
			},
			"python_named_params": basetypes.MapType{
				ElemType: types.StringType,
			},
			"python_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"spark_submit_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"sql_params": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

type RunNow struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands types.List `tfsdk:"dbt_commands" tf:"optional"`
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
	IdempotencyToken types.String `tfsdk:"idempotency_token" tf:"optional"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams types.List `tfsdk:"jar_params" tf:"optional"`
	// The ID of the job to be executed
	JobId types.Int64 `tfsdk:"job_id" tf:""`
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters types.Map `tfsdk:"job_parameters" tf:"optional"`
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
	NotebookParams types.Map `tfsdk:"notebook_params" tf:"optional"`
	// A list of task keys to run inside of the job. If this field is not
	// provided, all tasks in the job will be run.
	Only types.List `tfsdk:"only" tf:"optional"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams types.List `tfsdk:"pipeline_params" tf:"optional,object"`

	PythonNamedParams types.Map `tfsdk:"python_named_params" tf:"optional"`
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
	PythonParams types.List `tfsdk:"python_params" tf:"optional"`
	// The queue settings of the run.
	Queue types.List `tfsdk:"queue" tf:"optional,object"`
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
	SparkSubmitParams types.List `tfsdk:"spark_submit_params" tf:"optional"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams types.Map `tfsdk:"sql_params" tf:"optional"`
}

func (newState *RunNow) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunNow) {
}

func (newState *RunNow) SyncEffectiveFieldsDuringRead(existingState RunNow) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunNow.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunNow) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbt_commands":        reflect.TypeOf(types.String{}),
		"jar_params":          reflect.TypeOf(types.String{}),
		"job_parameters":      reflect.TypeOf(types.String{}),
		"notebook_params":     reflect.TypeOf(types.String{}),
		"only":                reflect.TypeOf(types.String{}),
		"pipeline_params":     reflect.TypeOf(PipelineParams{}),
		"python_named_params": reflect.TypeOf(types.String{}),
		"python_params":       reflect.TypeOf(types.String{}),
		"queue":               reflect.TypeOf(QueueSettings{}),
		"spark_submit_params": reflect.TypeOf(types.String{}),
		"sql_params":          reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of RunNow in the Terraform plugin framework type
// system.
func (a RunNow) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbt_commands": basetypes.ListType{
				ElemType: types.StringType,
			},
			"idempotency_token": types.StringType,
			"jar_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"job_id": types.Int64Type,
			"job_parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"notebook_params": basetypes.MapType{
				ElemType: types.StringType,
			},
			"only": basetypes.ListType{
				ElemType: types.StringType,
			},
			"pipeline_params": basetypes.ListType{
				ElemType: PipelineParams{}.ToObjectType(ctx),
			},
			"python_named_params": basetypes.MapType{
				ElemType: types.StringType,
			},
			"python_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"queue": basetypes.ListType{
				ElemType: QueueSettings{}.ToObjectType(ctx),
			},
			"spark_submit_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"sql_params": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// Run was started successfully.
type RunNowResponse struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob types.Int64 `tfsdk:"number_in_job" tf:"optional"`
	// The globally unique ID of the newly triggered run.
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
}

func (newState *RunNowResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunNowResponse) {
}

func (newState *RunNowResponse) SyncEffectiveFieldsDuringRead(existingState RunNowResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunNowResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunNowResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of RunNowResponse in the Terraform plugin framework type
// system.
func (a RunNowResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"number_in_job": types.Int64Type,
			"run_id":        types.Int64Type,
		},
	}
}

// Run output was retrieved successfully.
type RunOutput struct {
	// The output of a dbt task, if available.
	DbtOutput types.List `tfsdk:"dbt_output" tf:"optional,object"`
	// An error message indicating why a task failed or why output is not
	// available. The message is unstructured, and its exact format is subject
	// to change.
	Error types.String `tfsdk:"error" tf:"optional"`
	// If there was an error executing the run, this field contains any
	// available stack traces.
	ErrorTrace types.String `tfsdk:"error_trace" tf:"optional"`

	Info types.String `tfsdk:"info" tf:"optional"`
	// The output from tasks that write to standard streams (stdout/stderr) such
	// as spark_jar_task, spark_python_task, python_wheel_task.
	//
	// It's not supported for the notebook_task, pipeline_task or
	// spark_submit_task.
	//
	// Databricks restricts this API to return the last 5 MB of these logs.
	Logs types.String `tfsdk:"logs" tf:"optional"`
	// Whether the logs are truncated.
	LogsTruncated types.Bool `tfsdk:"logs_truncated" tf:"optional"`
	// All details of the run except for its output.
	Metadata types.List `tfsdk:"metadata" tf:"optional,object"`
	// The output of a notebook task, if available. A notebook task that
	// terminates (either successfully or with a failure) without calling
	// `dbutils.notebook.exit()` is considered to have an empty output. This
	// field is set but its result value is empty. Databricks restricts this API
	// to return the first 5 MB of the output. To return a larger result, use
	// the [ClusterLogConf] field to configure log storage for the job cluster.
	//
	// [ClusterLogConf]: https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterlogconf
	NotebookOutput types.List `tfsdk:"notebook_output" tf:"optional,object"`
	// The output of a run job task, if available
	RunJobOutput types.List `tfsdk:"run_job_output" tf:"optional,object"`
	// The output of a SQL task, if available.
	SqlOutput types.List `tfsdk:"sql_output" tf:"optional,object"`
}

func (newState *RunOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunOutput) {
}

func (newState *RunOutput) SyncEffectiveFieldsDuringRead(existingState RunOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbt_output":      reflect.TypeOf(DbtOutput{}),
		"metadata":        reflect.TypeOf(Run{}),
		"notebook_output": reflect.TypeOf(NotebookOutput{}),
		"run_job_output":  reflect.TypeOf(RunJobOutput{}),
		"sql_output":      reflect.TypeOf(SqlOutput{}),
	}
}

// ToObjectType returns the representation of RunOutput in the Terraform plugin framework type
// system.
func (a RunOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbt_output": basetypes.ListType{
				ElemType: DbtOutput{}.ToObjectType(ctx),
			},
			"error":          types.StringType,
			"error_trace":    types.StringType,
			"info":           types.StringType,
			"logs":           types.StringType,
			"logs_truncated": types.BoolType,
			"metadata": basetypes.ListType{
				ElemType: Run{}.ToObjectType(ctx),
			},
			"notebook_output": basetypes.ListType{
				ElemType: NotebookOutput{}.ToObjectType(ctx),
			},
			"run_job_output": basetypes.ListType{
				ElemType: RunJobOutput{}.ToObjectType(ctx),
			},
			"sql_output": basetypes.ListType{
				ElemType: SqlOutput{}.ToObjectType(ctx),
			},
		},
	}
}

type RunParameters struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands types.List `tfsdk:"dbt_commands" tf:"optional"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams types.List `tfsdk:"jar_params" tf:"optional"`
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
	NotebookParams types.Map `tfsdk:"notebook_params" tf:"optional"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams types.List `tfsdk:"pipeline_params" tf:"optional,object"`

	PythonNamedParams types.Map `tfsdk:"python_named_params" tf:"optional"`
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
	PythonParams types.List `tfsdk:"python_params" tf:"optional"`
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
	SparkSubmitParams types.List `tfsdk:"spark_submit_params" tf:"optional"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams types.Map `tfsdk:"sql_params" tf:"optional"`
}

func (newState *RunParameters) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunParameters) {
}

func (newState *RunParameters) SyncEffectiveFieldsDuringRead(existingState RunParameters) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunParameters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunParameters) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbt_commands":        reflect.TypeOf(types.String{}),
		"jar_params":          reflect.TypeOf(types.String{}),
		"notebook_params":     reflect.TypeOf(types.String{}),
		"pipeline_params":     reflect.TypeOf(PipelineParams{}),
		"python_named_params": reflect.TypeOf(types.String{}),
		"python_params":       reflect.TypeOf(types.String{}),
		"spark_submit_params": reflect.TypeOf(types.String{}),
		"sql_params":          reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of RunParameters in the Terraform plugin framework type
// system.
func (a RunParameters) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbt_commands": basetypes.ListType{
				ElemType: types.StringType,
			},
			"jar_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"notebook_params": basetypes.MapType{
				ElemType: types.StringType,
			},
			"pipeline_params": basetypes.ListType{
				ElemType: PipelineParams{}.ToObjectType(ctx),
			},
			"python_named_params": basetypes.MapType{
				ElemType: types.StringType,
			},
			"python_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"spark_submit_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"sql_params": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// The current state of the run.
type RunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response.
	LifeCycleState types.String `tfsdk:"life_cycle_state" tf:"optional"`
	// The reason indicating why the run was queued.
	QueueReason types.String `tfsdk:"queue_reason" tf:"optional"`
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states.
	ResultState types.String `tfsdk:"result_state" tf:"optional"`
	// A descriptive message for the current state. This field is unstructured,
	// and its exact format is subject to change.
	StateMessage types.String `tfsdk:"state_message" tf:"optional"`
	// A value indicating whether a run was canceled manually by a user or by
	// the scheduler because the run timed out.
	UserCancelledOrTimedout types.Bool `tfsdk:"user_cancelled_or_timedout" tf:"optional"`
}

func (newState *RunState) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunState) {
}

func (newState *RunState) SyncEffectiveFieldsDuringRead(existingState RunState) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunState.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunState) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of RunState in the Terraform plugin framework type
// system.
func (a RunState) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"life_cycle_state":           types.StringType,
			"queue_reason":               types.StringType,
			"result_state":               types.StringType,
			"state_message":              types.StringType,
			"user_cancelled_or_timedout": types.BoolType,
		},
	}
}

// The current status of the run
type RunStatus struct {
	// If the run was queued, details about the reason for queuing the run.
	QueueDetails types.List `tfsdk:"queue_details" tf:"optional,object"`
	// The current state of the run.
	State types.String `tfsdk:"state" tf:"optional"`
	// If the run is in a TERMINATING or TERMINATED state, details about the
	// reason for terminating the run.
	TerminationDetails types.List `tfsdk:"termination_details" tf:"optional,object"`
}

func (newState *RunStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunStatus) {
}

func (newState *RunStatus) SyncEffectiveFieldsDuringRead(existingState RunStatus) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"queue_details":       reflect.TypeOf(QueueDetails{}),
		"termination_details": reflect.TypeOf(TerminationDetails{}),
	}
}

// ToObjectType returns the representation of RunStatus in the Terraform plugin framework type
// system.
func (a RunStatus) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"queue_details": basetypes.ListType{
				ElemType: QueueDetails{}.ToObjectType(ctx),
			},
			"state": types.StringType,
			"termination_details": basetypes.ListType{
				ElemType: TerminationDetails{}.ToObjectType(ctx),
			},
		},
	}
}

// Used when outputting a child run, in GetRun or ListRuns.
type RunTask struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber types.Int64 `tfsdk:"attempt_number" tf:"optional"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration types.Int64 `tfsdk:"cleanup_duration" tf:"optional"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance types.List `tfsdk:"cluster_instance" tf:"optional,object"`
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask types.List `tfsdk:"condition_task" tf:"optional,object"`
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask types.List `tfsdk:"dbt_task" tf:"optional,object"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn types.List `tfsdk:"depends_on" tf:"optional"`
	// An optional description for this task.
	Description types.String `tfsdk:"description" tf:"optional"`
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications types.List `tfsdk:"email_notifications" tf:"optional,object"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey types.String `tfsdk:"environment_key" tf:"optional"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration types.Int64 `tfsdk:"execution_duration" tf:"optional"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId types.String `tfsdk:"existing_cluster_id" tf:"optional"`
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask types.List `tfsdk:"for_each_task" tf:"optional,object"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks. If `git_source` is set,
	// these tasks retrieve the file from the remote repository by default.
	// However, this behavior can be overridden by setting `source` to
	// `WORKSPACE` on the task. Note: dbt and SQL File tasks support only
	// version-controlled sources. If dbt or SQL File tasks are used,
	// `git_source` must be defined on the job.
	GitSource types.List `tfsdk:"git_source" tf:"optional,object"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey types.String `tfsdk:"job_cluster_key" tf:"optional"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries types.List `tfsdk:"library" tf:"optional"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster compute.ClusterSpec `tfsdk:"new_cluster" tf:"optional,object"`
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask types.List `tfsdk:"notebook_task" tf:"optional,object"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings types.List `tfsdk:"notification_settings" tf:"optional,object"`
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask types.List `tfsdk:"pipeline_task" tf:"optional,object"`
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask types.List `tfsdk:"python_wheel_task" tf:"optional,object"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration types.Int64 `tfsdk:"queue_duration" tf:"optional"`
	// Parameter values including resolved references
	ResolvedValues types.List `tfsdk:"resolved_values" tf:"optional,object"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration types.Int64 `tfsdk:"run_duration" tf:"optional"`
	// The ID of the task run.
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf types.String `tfsdk:"run_if" tf:"optional"`
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask types.List `tfsdk:"run_job_task" tf:"optional,object"`

	RunPageUrl types.String `tfsdk:"run_page_url" tf:"optional"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration types.Int64 `tfsdk:"setup_duration" tf:"optional"`
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask types.List `tfsdk:"spark_jar_task" tf:"optional,object"`
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask types.List `tfsdk:"spark_python_task" tf:"optional,object"`
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
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
	SparkSubmitTask types.List `tfsdk:"spark_submit_task" tf:"optional,object"`
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask types.List `tfsdk:"sql_task" tf:"optional,object"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Deprecated. Please use the `status` field instead.
	State types.List `tfsdk:"state" tf:"optional,object"`
	// The current status of the run
	Status types.List `tfsdk:"status" tf:"optional,object"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey types.String `tfsdk:"task_key" tf:""`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds" tf:"optional"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications types.List `tfsdk:"webhook_notifications" tf:"optional,object"`
}

func (newState *RunTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunTask) {
}

func (newState *RunTask) SyncEffectiveFieldsDuringRead(existingState RunTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster_instance":      reflect.TypeOf(ClusterInstance{}),
		"condition_task":        reflect.TypeOf(RunConditionTask{}),
		"dbt_task":              reflect.TypeOf(DbtTask{}),
		"depends_on":            reflect.TypeOf(TaskDependency{}),
		"email_notifications":   reflect.TypeOf(JobEmailNotifications{}),
		"for_each_task":         reflect.TypeOf(RunForEachTask{}),
		"git_source":            reflect.TypeOf(GitSource{}),
		"library":               reflect.TypeOf(compute.Library{}),
		"new_cluster":           reflect.TypeOf(compute.ClusterSpec{}),
		"notebook_task":         reflect.TypeOf(NotebookTask{}),
		"notification_settings": reflect.TypeOf(TaskNotificationSettings{}),
		"pipeline_task":         reflect.TypeOf(PipelineTask{}),
		"python_wheel_task":     reflect.TypeOf(PythonWheelTask{}),
		"resolved_values":       reflect.TypeOf(ResolvedValues{}),
		"run_job_task":          reflect.TypeOf(RunJobTask{}),
		"spark_jar_task":        reflect.TypeOf(SparkJarTask{}),
		"spark_python_task":     reflect.TypeOf(SparkPythonTask{}),
		"spark_submit_task":     reflect.TypeOf(SparkSubmitTask{}),
		"sql_task":              reflect.TypeOf(SqlTask{}),
		"state":                 reflect.TypeOf(RunState{}),
		"status":                reflect.TypeOf(RunStatus{}),
		"webhook_notifications": reflect.TypeOf(WebhookNotifications{}),
	}
}

// ToObjectType returns the representation of RunTask in the Terraform plugin framework type
// system.
func (a RunTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attempt_number":   types.Int64Type,
			"cleanup_duration": types.Int64Type,
			"cluster_instance": basetypes.ListType{
				ElemType: ClusterInstance{}.ToObjectType(ctx),
			},
			"condition_task": basetypes.ListType{
				ElemType: RunConditionTask{}.ToObjectType(ctx),
			},
			"dbt_task": basetypes.ListType{
				ElemType: DbtTask{}.ToObjectType(ctx),
			},
			"depends_on": basetypes.ListType{
				ElemType: TaskDependency{}.ToObjectType(ctx),
			},
			"description": types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: JobEmailNotifications{}.ToObjectType(ctx),
			},
			"end_time":            types.Int64Type,
			"environment_key":     types.StringType,
			"execution_duration":  types.Int64Type,
			"existing_cluster_id": types.StringType,
			"for_each_task": basetypes.ListType{
				ElemType: RunForEachTask{}.ToObjectType(ctx),
			},
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.ToObjectType(ctx),
			},
			"job_cluster_key": types.StringType,
			"library": basetypes.ListType{
				ElemType: compute_tf.Library{}.ToObjectType(ctx),
			},
			"new_cluster": basetypes.ListType{
				ElemType: compute_tf.ClusterSpec{}.ToObjectType(ctx),
			},
			"notebook_task": basetypes.ListType{
				ElemType: NotebookTask{}.ToObjectType(ctx),
			},
			"notification_settings": basetypes.ListType{
				ElemType: TaskNotificationSettings{}.ToObjectType(ctx),
			},
			"pipeline_task": basetypes.ListType{
				ElemType: PipelineTask{}.ToObjectType(ctx),
			},
			"python_wheel_task": basetypes.ListType{
				ElemType: PythonWheelTask{}.ToObjectType(ctx),
			},
			"queue_duration": types.Int64Type,
			"resolved_values": basetypes.ListType{
				ElemType: ResolvedValues{}.ToObjectType(ctx),
			},
			"run_duration": types.Int64Type,
			"run_id":       types.Int64Type,
			"run_if":       types.StringType,
			"run_job_task": basetypes.ListType{
				ElemType: RunJobTask{}.ToObjectType(ctx),
			},
			"run_page_url":   types.StringType,
			"setup_duration": types.Int64Type,
			"spark_jar_task": basetypes.ListType{
				ElemType: SparkJarTask{}.ToObjectType(ctx),
			},
			"spark_python_task": basetypes.ListType{
				ElemType: SparkPythonTask{}.ToObjectType(ctx),
			},
			"spark_submit_task": basetypes.ListType{
				ElemType: SparkSubmitTask{}.ToObjectType(ctx),
			},
			"sql_task": basetypes.ListType{
				ElemType: SqlTask{}.ToObjectType(ctx),
			},
			"start_time": types.Int64Type,
			"state": basetypes.ListType{
				ElemType: RunState{}.ToObjectType(ctx),
			},
			"status": basetypes.ListType{
				ElemType: RunStatus{}.ToObjectType(ctx),
			},
			"task_key":        types.StringType,
			"timeout_seconds": types.Int64Type,
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.ToObjectType(ctx),
			},
		},
	}
}

type SparkJarTask struct {
	// Deprecated since 04/2016. Provide a `jar` through the `libraries` field
	// instead. For an example, see :method:jobs/create.
	JarUri types.String `tfsdk:"jar_uri" tf:"optional"`
	// The full name of the class containing the main method to be executed.
	// This class must be contained in a JAR provided as a library.
	//
	// The code must use `SparkContext.getOrCreate` to obtain a Spark context;
	// otherwise, runs of the job fail.
	MainClassName types.String `tfsdk:"main_class_name" tf:"optional"`
	// Parameters passed to the main method.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
}

func (newState *SparkJarTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkJarTask) {
}

func (newState *SparkJarTask) SyncEffectiveFieldsDuringRead(existingState SparkJarTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkJarTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SparkJarTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of SparkJarTask in the Terraform plugin framework type
// system.
func (a SparkJarTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"jar_uri":         types.StringType,
			"main_class_name": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type SparkPythonTask struct {
	// Command line parameters passed to the Python file.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// The Python file to be executed. Cloud file URIs (such as dbfs:/, s3:/,
	// adls:/, gcs:/) and workspace paths are supported. For python files stored
	// in the Databricks workspace, the path must be absolute and begin with
	// `/`. For files stored in a remote repository, the path must be relative.
	// This field is required.
	PythonFile types.String `tfsdk:"python_file" tf:""`
	// Optional location type of the Python file. When set to `WORKSPACE` or not
	// specified, the file will be retrieved from the local Databricks workspace
	// or cloud location (if the `python_file` has a URI format). When set to
	// `GIT`, the Python file will be retrieved from a Git repository defined in
	// `git_source`.
	//
	// * `WORKSPACE`: The Python file is located in a Databricks workspace or at
	// a cloud filesystem URI. * `GIT`: The Python file is located in a remote
	// Git repository.
	Source types.String `tfsdk:"source" tf:"optional"`
}

func (newState *SparkPythonTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkPythonTask) {
}

func (newState *SparkPythonTask) SyncEffectiveFieldsDuringRead(existingState SparkPythonTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkPythonTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SparkPythonTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of SparkPythonTask in the Terraform plugin framework type
// system.
func (a SparkPythonTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parameters": basetypes.ListType{
				ElemType: types.StringType,
			},
			"python_file": types.StringType,
			"source":      types.StringType,
		},
	}
}

type SparkSubmitTask struct {
	// Command-line parameters passed to spark submit.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
}

func (newState *SparkSubmitTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkSubmitTask) {
}

func (newState *SparkSubmitTask) SyncEffectiveFieldsDuringRead(existingState SparkSubmitTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkSubmitTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SparkSubmitTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of SparkSubmitTask in the Terraform plugin framework type
// system.
func (a SparkSubmitTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parameters": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type SqlAlertOutput struct {
	// The state of the SQL alert.
	//
	// * UNKNOWN: alert yet to be evaluated * OK: alert evaluated and did not
	// fulfill trigger conditions * TRIGGERED: alert evaluated and fulfilled
	// trigger conditions
	AlertState types.String `tfsdk:"alert_state" tf:"optional"`
	// The link to find the output results.
	OutputLink types.String `tfsdk:"output_link" tf:"optional"`
	// The text of the SQL query. Can Run permission of the SQL query associated
	// with the SQL alert is required to view this field.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// Information about SQL statements executed in the run.
	SqlStatements types.List `tfsdk:"sql_statements" tf:"optional"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *SqlAlertOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlAlertOutput) {
}

func (newState *SqlAlertOutput) SyncEffectiveFieldsDuringRead(existingState SqlAlertOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlAlertOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlAlertOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sql_statements": reflect.TypeOf(SqlStatementOutput{}),
	}
}

// ToObjectType returns the representation of SqlAlertOutput in the Terraform plugin framework type
// system.
func (a SqlAlertOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_state": types.StringType,
			"output_link": types.StringType,
			"query_text":  types.StringType,
			"sql_statements": basetypes.ListType{
				ElemType: SqlStatementOutput{}.ToObjectType(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

type SqlDashboardOutput struct {
	// The canonical identifier of the SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
	// Widgets executed in the run. Only SQL query based widgets are listed.
	Widgets types.List `tfsdk:"widgets" tf:"optional"`
}

func (newState *SqlDashboardOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlDashboardOutput) {
}

func (newState *SqlDashboardOutput) SyncEffectiveFieldsDuringRead(existingState SqlDashboardOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlDashboardOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlDashboardOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"widgets": reflect.TypeOf(SqlDashboardWidgetOutput{}),
	}
}

// ToObjectType returns the representation of SqlDashboardOutput in the Terraform plugin framework type
// system.
func (a SqlDashboardOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warehouse_id": types.StringType,
			"widgets": basetypes.ListType{
				ElemType: SqlDashboardWidgetOutput{}.ToObjectType(ctx),
			},
		},
	}
}

type SqlDashboardWidgetOutput struct {
	// Time (in epoch milliseconds) when execution of the SQL widget ends.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// The information about the error when execution fails.
	Error types.List `tfsdk:"error" tf:"optional,object"`
	// The link to find the output results.
	OutputLink types.String `tfsdk:"output_link" tf:"optional"`
	// Time (in epoch milliseconds) when execution of the SQL widget starts.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// The execution status of the SQL widget.
	Status types.String `tfsdk:"status" tf:"optional"`
	// The canonical identifier of the SQL widget.
	WidgetId types.String `tfsdk:"widget_id" tf:"optional"`
	// The title of the SQL widget.
	WidgetTitle types.String `tfsdk:"widget_title" tf:"optional"`
}

func (newState *SqlDashboardWidgetOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlDashboardWidgetOutput) {
}

func (newState *SqlDashboardWidgetOutput) SyncEffectiveFieldsDuringRead(existingState SqlDashboardWidgetOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlDashboardWidgetOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlDashboardWidgetOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(SqlOutputError{}),
	}
}

// ToObjectType returns the representation of SqlDashboardWidgetOutput in the Terraform plugin framework type
// system.
func (a SqlDashboardWidgetOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time": types.Int64Type,
			"error": basetypes.ListType{
				ElemType: SqlOutputError{}.ToObjectType(ctx),
			},
			"output_link":  types.StringType,
			"start_time":   types.Int64Type,
			"status":       types.StringType,
			"widget_id":    types.StringType,
			"widget_title": types.StringType,
		},
	}
}

type SqlOutput struct {
	// The output of a SQL alert task, if available.
	AlertOutput types.List `tfsdk:"alert_output" tf:"optional,object"`
	// The output of a SQL dashboard task, if available.
	DashboardOutput types.List `tfsdk:"dashboard_output" tf:"optional,object"`
	// The output of a SQL query task, if available.
	QueryOutput types.List `tfsdk:"query_output" tf:"optional,object"`
}

func (newState *SqlOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlOutput) {
}

func (newState *SqlOutput) SyncEffectiveFieldsDuringRead(existingState SqlOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_output":     reflect.TypeOf(SqlAlertOutput{}),
		"dashboard_output": reflect.TypeOf(SqlDashboardOutput{}),
		"query_output":     reflect.TypeOf(SqlQueryOutput{}),
	}
}

// ToObjectType returns the representation of SqlOutput in the Terraform plugin framework type
// system.
func (a SqlOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_output": basetypes.ListType{
				ElemType: SqlAlertOutput{}.ToObjectType(ctx),
			},
			"dashboard_output": basetypes.ListType{
				ElemType: SqlDashboardOutput{}.ToObjectType(ctx),
			},
			"query_output": basetypes.ListType{
				ElemType: SqlQueryOutput{}.ToObjectType(ctx),
			},
		},
	}
}

type SqlOutputError struct {
	// The error message when execution fails.
	Message types.String `tfsdk:"message" tf:"optional"`
}

func (newState *SqlOutputError) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlOutputError) {
}

func (newState *SqlOutputError) SyncEffectiveFieldsDuringRead(existingState SqlOutputError) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlOutputError.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlOutputError) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of SqlOutputError in the Terraform plugin framework type
// system.
func (a SqlOutputError) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
		},
	}
}

type SqlQueryOutput struct {
	EndpointId types.String `tfsdk:"endpoint_id" tf:"optional"`
	// The link to find the output results.
	OutputLink types.String `tfsdk:"output_link" tf:"optional"`
	// The text of the SQL query. Can Run permission of the SQL query is
	// required to view this field.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// Information about SQL statements executed in the run.
	SqlStatements types.List `tfsdk:"sql_statements" tf:"optional"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *SqlQueryOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlQueryOutput) {
}

func (newState *SqlQueryOutput) SyncEffectiveFieldsDuringRead(existingState SqlQueryOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlQueryOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlQueryOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sql_statements": reflect.TypeOf(SqlStatementOutput{}),
	}
}

// ToObjectType returns the representation of SqlQueryOutput in the Terraform plugin framework type
// system.
func (a SqlQueryOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_id": types.StringType,
			"output_link": types.StringType,
			"query_text":  types.StringType,
			"sql_statements": basetypes.ListType{
				ElemType: SqlStatementOutput{}.ToObjectType(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

type SqlStatementOutput struct {
	// A key that can be used to look up query details.
	LookupKey types.String `tfsdk:"lookup_key" tf:"optional"`
}

func (newState *SqlStatementOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlStatementOutput) {
}

func (newState *SqlStatementOutput) SyncEffectiveFieldsDuringRead(existingState SqlStatementOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlStatementOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlStatementOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of SqlStatementOutput in the Terraform plugin framework type
// system.
func (a SqlStatementOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"lookup_key": types.StringType,
		},
	}
}

type SqlTask struct {
	// If alert, indicates that this job must refresh a SQL alert.
	Alert types.List `tfsdk:"alert" tf:"optional,object"`
	// If dashboard, indicates that this job must refresh a SQL dashboard.
	Dashboard types.List `tfsdk:"dashboard" tf:"optional,object"`
	// If file, indicates that this job runs a SQL file in a remote Git
	// repository.
	File types.List `tfsdk:"file" tf:"optional,object"`
	// Parameters to be used for each run of this job. The SQL alert task does
	// not support custom parameters.
	Parameters types.Map `tfsdk:"parameters" tf:"optional"`
	// If query, indicates that this job must execute a SQL query.
	Query types.List `tfsdk:"query" tf:"optional,object"`
	// The canonical identifier of the SQL warehouse. Recommended to use with
	// serverless or pro SQL warehouses. Classic SQL warehouses are only
	// supported for SQL alert, dashboard and query tasks and are limited to
	// scheduled single-task jobs.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:""`
}

func (newState *SqlTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTask) {
}

func (newState *SqlTask) SyncEffectiveFieldsDuringRead(existingState SqlTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert":      reflect.TypeOf(SqlTaskAlert{}),
		"dashboard":  reflect.TypeOf(SqlTaskDashboard{}),
		"file":       reflect.TypeOf(SqlTaskFile{}),
		"parameters": reflect.TypeOf(types.String{}),
		"query":      reflect.TypeOf(SqlTaskQuery{}),
	}
}

// ToObjectType returns the representation of SqlTask in the Terraform plugin framework type
// system.
func (a SqlTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert": basetypes.ListType{
				ElemType: SqlTaskAlert{}.ToObjectType(ctx),
			},
			"dashboard": basetypes.ListType{
				ElemType: SqlTaskDashboard{}.ToObjectType(ctx),
			},
			"file": basetypes.ListType{
				ElemType: SqlTaskFile{}.ToObjectType(ctx),
			},
			"parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"query": basetypes.ListType{
				ElemType: SqlTaskQuery{}.ToObjectType(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

type SqlTaskAlert struct {
	// The canonical identifier of the SQL alert.
	AlertId types.String `tfsdk:"alert_id" tf:""`
	// If true, the alert notifications are not sent to subscribers.
	PauseSubscriptions types.Bool `tfsdk:"pause_subscriptions" tf:"optional"`
	// If specified, alert notifications are sent to subscribers.
	Subscriptions types.List `tfsdk:"subscriptions" tf:"optional"`
}

func (newState *SqlTaskAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTaskAlert) {
}

func (newState *SqlTaskAlert) SyncEffectiveFieldsDuringRead(existingState SqlTaskAlert) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTaskAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTaskAlert) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriptions": reflect.TypeOf(SqlTaskSubscription{}),
	}
}

// ToObjectType returns the representation of SqlTaskAlert in the Terraform plugin framework type
// system.
func (a SqlTaskAlert) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id":            types.StringType,
			"pause_subscriptions": types.BoolType,
			"subscriptions": basetypes.ListType{
				ElemType: SqlTaskSubscription{}.ToObjectType(ctx),
			},
		},
	}
}

type SqlTaskDashboard struct {
	// Subject of the email sent to subscribers of this task.
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// The canonical identifier of the SQL dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:""`
	// If true, the dashboard snapshot is not taken, and emails are not sent to
	// subscribers.
	PauseSubscriptions types.Bool `tfsdk:"pause_subscriptions" tf:"optional"`
	// If specified, dashboard snapshots are sent to subscriptions.
	Subscriptions types.List `tfsdk:"subscriptions" tf:"optional"`
}

func (newState *SqlTaskDashboard) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTaskDashboard) {
}

func (newState *SqlTaskDashboard) SyncEffectiveFieldsDuringRead(existingState SqlTaskDashboard) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTaskDashboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTaskDashboard) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriptions": reflect.TypeOf(SqlTaskSubscription{}),
	}
}

// ToObjectType returns the representation of SqlTaskDashboard in the Terraform plugin framework type
// system.
func (a SqlTaskDashboard) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_subject":      types.StringType,
			"dashboard_id":        types.StringType,
			"pause_subscriptions": types.BoolType,
			"subscriptions": basetypes.ListType{
				ElemType: SqlTaskSubscription{}.ToObjectType(ctx),
			},
		},
	}
}

type SqlTaskFile struct {
	// Path of the SQL file. Must be relative if the source is a remote Git
	// repository and absolute for workspace paths.
	Path types.String `tfsdk:"path" tf:""`
	// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL
	// file will be retrieved from the local Databricks workspace. When set to
	// `GIT`, the SQL file will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: SQL file is located in Databricks workspace. * `GIT`: SQL
	// file is located in cloud Git provider.
	Source types.String `tfsdk:"source" tf:"optional"`
}

func (newState *SqlTaskFile) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTaskFile) {
}

func (newState *SqlTaskFile) SyncEffectiveFieldsDuringRead(existingState SqlTaskFile) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTaskFile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTaskFile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of SqlTaskFile in the Terraform plugin framework type
// system.
func (a SqlTaskFile) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":   types.StringType,
			"source": types.StringType,
		},
	}
}

type SqlTaskQuery struct {
	// The canonical identifier of the SQL query.
	QueryId types.String `tfsdk:"query_id" tf:""`
}

func (newState *SqlTaskQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTaskQuery) {
}

func (newState *SqlTaskQuery) SyncEffectiveFieldsDuringRead(existingState SqlTaskQuery) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTaskQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTaskQuery) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of SqlTaskQuery in the Terraform plugin framework type
// system.
func (a SqlTaskQuery) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_id": types.StringType,
		},
	}
}

type SqlTaskSubscription struct {
	// The canonical identifier of the destination to receive email
	// notification. This parameter is mutually exclusive with user_name. You
	// cannot set both destination_id and user_name for subscription
	// notifications.
	DestinationId types.String `tfsdk:"destination_id" tf:"optional"`
	// The user name to receive the subscription email. This parameter is
	// mutually exclusive with destination_id. You cannot set both
	// destination_id and user_name for subscription notifications.
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *SqlTaskSubscription) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTaskSubscription) {
}

func (newState *SqlTaskSubscription) SyncEffectiveFieldsDuringRead(existingState SqlTaskSubscription) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTaskSubscription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTaskSubscription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of SqlTaskSubscription in the Terraform plugin framework type
// system.
func (a SqlTaskSubscription) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_id": types.StringType,
			"user_name":      types.StringType,
		},
	}
}

type SubmitRun struct {
	// List of permissions to set on the job.
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The user specified id of the budget policy to use for this one-time run.
	// If not specified, the run will be not be attributed to any budget policy.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id" tf:"optional"`
	// An optional set of email addresses notified when the run begins or
	// completes.
	EmailNotifications types.List `tfsdk:"email_notifications" tf:"optional,object"`
	// A list of task execution environment specifications that can be
	// referenced by tasks of this run.
	Environments types.List `tfsdk:"environments" tf:"optional"`
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
	GitSource types.List `tfsdk:"git_source" tf:"optional,object"`
	// An optional set of health rules that can be defined for this job.
	Health types.List `tfsdk:"health" tf:"optional,object"`
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
	IdempotencyToken types.String `tfsdk:"idempotency_token" tf:"optional"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// run.
	NotificationSettings types.List `tfsdk:"notification_settings" tf:"optional,object"`
	// The queue settings of the one-time run.
	Queue types.List `tfsdk:"queue" tf:"optional,object"`
	// Specifies the user or service principal that the job runs as. If not
	// specified, the job runs as the user who submits the request.
	RunAs types.List `tfsdk:"run_as" tf:"optional,object"`
	// An optional name for the run. The default value is `Untitled`.
	RunName types.String `tfsdk:"run_name" tf:"optional"`

	Tasks types.List `tfsdk:"tasks" tf:"optional"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds" tf:"optional"`
	// A collection of system notification IDs to notify when the run begins or
	// completes.
	WebhookNotifications types.List `tfsdk:"webhook_notifications" tf:"optional,object"`
}

func (newState *SubmitRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubmitRun) {
}

func (newState *SubmitRun) SyncEffectiveFieldsDuringRead(existingState SubmitRun) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubmitRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubmitRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list":   reflect.TypeOf(JobAccessControlRequest{}),
		"email_notifications":   reflect.TypeOf(JobEmailNotifications{}),
		"environments":          reflect.TypeOf(JobEnvironment{}),
		"git_source":            reflect.TypeOf(GitSource{}),
		"health":                reflect.TypeOf(JobsHealthRules{}),
		"notification_settings": reflect.TypeOf(JobNotificationSettings{}),
		"queue":                 reflect.TypeOf(QueueSettings{}),
		"run_as":                reflect.TypeOf(JobRunAs{}),
		"tasks":                 reflect.TypeOf(SubmitTask{}),
		"webhook_notifications": reflect.TypeOf(WebhookNotifications{}),
	}
}

// ToObjectType returns the representation of SubmitRun in the Terraform plugin framework type
// system.
func (a SubmitRun) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: JobAccessControlRequest{}.ToObjectType(ctx),
			},
			"budget_policy_id": types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: JobEmailNotifications{}.ToObjectType(ctx),
			},
			"environments": basetypes.ListType{
				ElemType: JobEnvironment{}.ToObjectType(ctx),
			},
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.ToObjectType(ctx),
			},
			"health": basetypes.ListType{
				ElemType: JobsHealthRules{}.ToObjectType(ctx),
			},
			"idempotency_token": types.StringType,
			"notification_settings": basetypes.ListType{
				ElemType: JobNotificationSettings{}.ToObjectType(ctx),
			},
			"queue": basetypes.ListType{
				ElemType: QueueSettings{}.ToObjectType(ctx),
			},
			"run_as": basetypes.ListType{
				ElemType: JobRunAs{}.ToObjectType(ctx),
			},
			"run_name": types.StringType,
			"tasks": basetypes.ListType{
				ElemType: SubmitTask{}.ToObjectType(ctx),
			},
			"timeout_seconds": types.Int64Type,
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.ToObjectType(ctx),
			},
		},
	}
}

// Run was created and started successfully.
type SubmitRunResponse struct {
	// The canonical identifier for the newly submitted run.
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
}

func (newState *SubmitRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubmitRunResponse) {
}

func (newState *SubmitRunResponse) SyncEffectiveFieldsDuringRead(existingState SubmitRunResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubmitRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubmitRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of SubmitRunResponse in the Terraform plugin framework type
// system.
func (a SubmitRunResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

type SubmitTask struct {
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask types.List `tfsdk:"condition_task" tf:"optional,object"`
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask types.List `tfsdk:"dbt_task" tf:"optional,object"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn types.List `tfsdk:"depends_on" tf:"optional"`
	// An optional description for this task.
	Description types.String `tfsdk:"description" tf:"optional"`
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications types.List `tfsdk:"email_notifications" tf:"optional,object"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey types.String `tfsdk:"environment_key" tf:"optional"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId types.String `tfsdk:"existing_cluster_id" tf:"optional"`
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask types.List `tfsdk:"for_each_task" tf:"optional,object"`
	// An optional set of health rules that can be defined for this job.
	Health types.List `tfsdk:"health" tf:"optional,object"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries types.List `tfsdk:"library" tf:"optional"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster compute.ClusterSpec `tfsdk:"new_cluster" tf:"optional,object"`
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask types.List `tfsdk:"notebook_task" tf:"optional,object"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings types.List `tfsdk:"notification_settings" tf:"optional,object"`
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask types.List `tfsdk:"pipeline_task" tf:"optional,object"`
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask types.List `tfsdk:"python_wheel_task" tf:"optional,object"`
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf types.String `tfsdk:"run_if" tf:"optional"`
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask types.List `tfsdk:"run_job_task" tf:"optional,object"`
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask types.List `tfsdk:"spark_jar_task" tf:"optional,object"`
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask types.List `tfsdk:"spark_python_task" tf:"optional,object"`
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
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
	SparkSubmitTask types.List `tfsdk:"spark_submit_task" tf:"optional,object"`
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask types.List `tfsdk:"sql_task" tf:"optional,object"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey types.String `tfsdk:"task_key" tf:""`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds" tf:"optional"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications types.List `tfsdk:"webhook_notifications" tf:"optional,object"`
}

func (newState *SubmitTask) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubmitTask) {
}

func (newState *SubmitTask) SyncEffectiveFieldsDuringRead(existingState SubmitTask) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubmitTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubmitTask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition_task":        reflect.TypeOf(ConditionTask{}),
		"dbt_task":              reflect.TypeOf(DbtTask{}),
		"depends_on":            reflect.TypeOf(TaskDependency{}),
		"email_notifications":   reflect.TypeOf(JobEmailNotifications{}),
		"for_each_task":         reflect.TypeOf(ForEachTask{}),
		"health":                reflect.TypeOf(JobsHealthRules{}),
		"library":               reflect.TypeOf(compute.Library{}),
		"new_cluster":           reflect.TypeOf(compute.ClusterSpec{}),
		"notebook_task":         reflect.TypeOf(NotebookTask{}),
		"notification_settings": reflect.TypeOf(TaskNotificationSettings{}),
		"pipeline_task":         reflect.TypeOf(PipelineTask{}),
		"python_wheel_task":     reflect.TypeOf(PythonWheelTask{}),
		"run_job_task":          reflect.TypeOf(RunJobTask{}),
		"spark_jar_task":        reflect.TypeOf(SparkJarTask{}),
		"spark_python_task":     reflect.TypeOf(SparkPythonTask{}),
		"spark_submit_task":     reflect.TypeOf(SparkSubmitTask{}),
		"sql_task":              reflect.TypeOf(SqlTask{}),
		"webhook_notifications": reflect.TypeOf(WebhookNotifications{}),
	}
}

// ToObjectType returns the representation of SubmitTask in the Terraform plugin framework type
// system.
func (a SubmitTask) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition_task": basetypes.ListType{
				ElemType: ConditionTask{}.ToObjectType(ctx),
			},
			"dbt_task": basetypes.ListType{
				ElemType: DbtTask{}.ToObjectType(ctx),
			},
			"depends_on": basetypes.ListType{
				ElemType: TaskDependency{}.ToObjectType(ctx),
			},
			"description": types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: JobEmailNotifications{}.ToObjectType(ctx),
			},
			"environment_key":     types.StringType,
			"existing_cluster_id": types.StringType,
			"for_each_task": basetypes.ListType{
				ElemType: ForEachTask{}.ToObjectType(ctx),
			},
			"health": basetypes.ListType{
				ElemType: JobsHealthRules{}.ToObjectType(ctx),
			},
			"library": basetypes.ListType{
				ElemType: compute_tf.Library{}.ToObjectType(ctx),
			},
			"new_cluster": basetypes.ListType{
				ElemType: compute_tf.ClusterSpec{}.ToObjectType(ctx),
			},
			"notebook_task": basetypes.ListType{
				ElemType: NotebookTask{}.ToObjectType(ctx),
			},
			"notification_settings": basetypes.ListType{
				ElemType: TaskNotificationSettings{}.ToObjectType(ctx),
			},
			"pipeline_task": basetypes.ListType{
				ElemType: PipelineTask{}.ToObjectType(ctx),
			},
			"python_wheel_task": basetypes.ListType{
				ElemType: PythonWheelTask{}.ToObjectType(ctx),
			},
			"run_if": types.StringType,
			"run_job_task": basetypes.ListType{
				ElemType: RunJobTask{}.ToObjectType(ctx),
			},
			"spark_jar_task": basetypes.ListType{
				ElemType: SparkJarTask{}.ToObjectType(ctx),
			},
			"spark_python_task": basetypes.ListType{
				ElemType: SparkPythonTask{}.ToObjectType(ctx),
			},
			"spark_submit_task": basetypes.ListType{
				ElemType: SparkSubmitTask{}.ToObjectType(ctx),
			},
			"sql_task": basetypes.ListType{
				ElemType: SqlTask{}.ToObjectType(ctx),
			},
			"task_key":        types.StringType,
			"timeout_seconds": types.Int64Type,
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.ToObjectType(ctx),
			},
		},
	}
}

type TableUpdateTriggerConfiguration struct {
	// The table(s) condition based on which to trigger a job run.
	Condition types.String `tfsdk:"condition" tf:"optional"`
	// If set, the trigger starts a run only after the specified amount of time
	// has passed since the last time the trigger fired. The minimum allowed
	// value is 60 seconds.
	MinTimeBetweenTriggersSeconds types.Int64 `tfsdk:"min_time_between_triggers_seconds" tf:"optional"`
	// A list of Delta tables to monitor for changes. The table name must be in
	// the format `catalog_name.schema_name.table_name`.
	TableNames types.List `tfsdk:"table_names" tf:"optional"`
	// If set, the trigger starts a run only after no table updates have
	// occurred for the specified time and can be used to wait for a series of
	// table updates before triggering a run. The minimum allowed value is 60
	// seconds.
	WaitAfterLastChangeSeconds types.Int64 `tfsdk:"wait_after_last_change_seconds" tf:"optional"`
}

func (newState *TableUpdateTriggerConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableUpdateTriggerConfiguration) {
}

func (newState *TableUpdateTriggerConfiguration) SyncEffectiveFieldsDuringRead(existingState TableUpdateTriggerConfiguration) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableUpdateTriggerConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableUpdateTriggerConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_names": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of TableUpdateTriggerConfiguration in the Terraform plugin framework type
// system.
func (a TableUpdateTriggerConfiguration) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition":                         types.StringType,
			"min_time_between_triggers_seconds": types.Int64Type,
			"table_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"wait_after_last_change_seconds": types.Int64Type,
		},
	}
}

type Task struct {
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask types.List `tfsdk:"condition_task" tf:"optional,object"`
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask types.List `tfsdk:"dbt_task" tf:"optional,object"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete before executing this
	// task. The task will run only if the `run_if` condition is true. The key
	// is `task_key`, and the value is the name assigned to the dependent task.
	DependsOn types.List `tfsdk:"depends_on" tf:"optional"`
	// An optional description for this task.
	Description types.String `tfsdk:"description" tf:"optional"`
	// An option to disable auto optimization in serverless
	DisableAutoOptimization types.Bool `tfsdk:"disable_auto_optimization" tf:"optional"`
	// An optional set of email addresses that is notified when runs of this
	// task begin or complete as well as when this task is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications types.List `tfsdk:"email_notifications" tf:"optional,object"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey types.String `tfsdk:"environment_key" tf:"optional"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId types.String `tfsdk:"existing_cluster_id" tf:"optional"`
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask types.List `tfsdk:"for_each_task" tf:"optional,object"`
	// An optional set of health rules that can be defined for this job.
	Health types.List `tfsdk:"health" tf:"optional,object"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey types.String `tfsdk:"job_cluster_key" tf:"optional"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries types.List `tfsdk:"library" tf:"optional"`
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value `-1` means
	// to retry indefinitely and the value `0` means to never retry.
	MaxRetries types.Int64 `tfsdk:"max_retries" tf:"optional"`
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	MinRetryIntervalMillis types.Int64 `tfsdk:"min_retry_interval_millis" tf:"optional"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster compute.ClusterSpec `tfsdk:"new_cluster" tf:"optional,object"`
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask types.List `tfsdk:"notebook_task" tf:"optional,object"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task.
	NotificationSettings types.List `tfsdk:"notification_settings" tf:"optional,object"`
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask types.List `tfsdk:"pipeline_task" tf:"optional,object"`
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask types.List `tfsdk:"python_wheel_task" tf:"optional,object"`
	// An optional policy to specify whether to retry a job when it times out.
	// The default behavior is to not retry on timeout.
	RetryOnTimeout types.Bool `tfsdk:"retry_on_timeout" tf:"optional"`
	// An optional value specifying the condition determining whether the task
	// is run once its dependencies have been completed.
	//
	// * `ALL_SUCCESS`: All dependencies have executed and succeeded *
	// `AT_LEAST_ONE_SUCCESS`: At least one dependency has succeeded *
	// `NONE_FAILED`: None of the dependencies have failed and at least one was
	// executed * `ALL_DONE`: All dependencies have been completed *
	// `AT_LEAST_ONE_FAILED`: At least one dependency failed * `ALL_FAILED`: ALl
	// dependencies have failed
	RunIf types.String `tfsdk:"run_if" tf:"optional"`
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask types.List `tfsdk:"run_job_task" tf:"optional,object"`
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask types.List `tfsdk:"spark_jar_task" tf:"optional,object"`
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask types.List `tfsdk:"spark_python_task" tf:"optional,object"`
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
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
	SparkSubmitTask types.List `tfsdk:"spark_submit_task" tf:"optional,object"`
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask types.List `tfsdk:"sql_task" tf:"optional,object"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey types.String `tfsdk:"task_key" tf:""`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds types.Int64 `tfsdk:"timeout_seconds" tf:"optional"`
	// A collection of system notification IDs to notify when runs of this task
	// begin or complete. The default behavior is to not send any system
	// notifications.
	WebhookNotifications types.List `tfsdk:"webhook_notifications" tf:"optional,object"`
}

func (newState *Task) SyncEffectiveFieldsDuringCreateOrUpdate(plan Task) {
}

func (newState *Task) SyncEffectiveFieldsDuringRead(existingState Task) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Task.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Task) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition_task":        reflect.TypeOf(ConditionTask{}),
		"dbt_task":              reflect.TypeOf(DbtTask{}),
		"depends_on":            reflect.TypeOf(TaskDependency{}),
		"email_notifications":   reflect.TypeOf(TaskEmailNotifications{}),
		"for_each_task":         reflect.TypeOf(ForEachTask{}),
		"health":                reflect.TypeOf(JobsHealthRules{}),
		"library":               reflect.TypeOf(compute.Library{}),
		"new_cluster":           reflect.TypeOf(compute.ClusterSpec{}),
		"notebook_task":         reflect.TypeOf(NotebookTask{}),
		"notification_settings": reflect.TypeOf(TaskNotificationSettings{}),
		"pipeline_task":         reflect.TypeOf(PipelineTask{}),
		"python_wheel_task":     reflect.TypeOf(PythonWheelTask{}),
		"run_job_task":          reflect.TypeOf(RunJobTask{}),
		"spark_jar_task":        reflect.TypeOf(SparkJarTask{}),
		"spark_python_task":     reflect.TypeOf(SparkPythonTask{}),
		"spark_submit_task":     reflect.TypeOf(SparkSubmitTask{}),
		"sql_task":              reflect.TypeOf(SqlTask{}),
		"webhook_notifications": reflect.TypeOf(WebhookNotifications{}),
	}
}

// ToObjectType returns the representation of Task in the Terraform plugin framework type
// system.
func (a Task) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition_task": basetypes.ListType{
				ElemType: ConditionTask{}.ToObjectType(ctx),
			},
			"dbt_task": basetypes.ListType{
				ElemType: DbtTask{}.ToObjectType(ctx),
			},
			"depends_on": basetypes.ListType{
				ElemType: TaskDependency{}.ToObjectType(ctx),
			},
			"description":               types.StringType,
			"disable_auto_optimization": types.BoolType,
			"email_notifications": basetypes.ListType{
				ElemType: TaskEmailNotifications{}.ToObjectType(ctx),
			},
			"environment_key":     types.StringType,
			"existing_cluster_id": types.StringType,
			"for_each_task": basetypes.ListType{
				ElemType: ForEachTask{}.ToObjectType(ctx),
			},
			"health": basetypes.ListType{
				ElemType: JobsHealthRules{}.ToObjectType(ctx),
			},
			"job_cluster_key": types.StringType,
			"library": basetypes.ListType{
				ElemType: compute_tf.Library{}.ToObjectType(ctx),
			},
			"max_retries":               types.Int64Type,
			"min_retry_interval_millis": types.Int64Type,
			"new_cluster": basetypes.ListType{
				ElemType: compute_tf.ClusterSpec{}.ToObjectType(ctx),
			},
			"notebook_task": basetypes.ListType{
				ElemType: NotebookTask{}.ToObjectType(ctx),
			},
			"notification_settings": basetypes.ListType{
				ElemType: TaskNotificationSettings{}.ToObjectType(ctx),
			},
			"pipeline_task": basetypes.ListType{
				ElemType: PipelineTask{}.ToObjectType(ctx),
			},
			"python_wheel_task": basetypes.ListType{
				ElemType: PythonWheelTask{}.ToObjectType(ctx),
			},
			"retry_on_timeout": types.BoolType,
			"run_if":           types.StringType,
			"run_job_task": basetypes.ListType{
				ElemType: RunJobTask{}.ToObjectType(ctx),
			},
			"spark_jar_task": basetypes.ListType{
				ElemType: SparkJarTask{}.ToObjectType(ctx),
			},
			"spark_python_task": basetypes.ListType{
				ElemType: SparkPythonTask{}.ToObjectType(ctx),
			},
			"spark_submit_task": basetypes.ListType{
				ElemType: SparkSubmitTask{}.ToObjectType(ctx),
			},
			"sql_task": basetypes.ListType{
				ElemType: SqlTask{}.ToObjectType(ctx),
			},
			"task_key":        types.StringType,
			"timeout_seconds": types.Int64Type,
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.ToObjectType(ctx),
			},
		},
	}
}

type TaskDependency struct {
	// Can only be specified on condition task dependencies. The outcome of the
	// dependent task that must be met for this task to run.
	Outcome types.String `tfsdk:"outcome" tf:"optional"`
	// The name of the task this task depends on.
	TaskKey types.String `tfsdk:"task_key" tf:""`
}

func (newState *TaskDependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan TaskDependency) {
}

func (newState *TaskDependency) SyncEffectiveFieldsDuringRead(existingState TaskDependency) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TaskDependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TaskDependency) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of TaskDependency in the Terraform plugin framework type
// system.
func (a TaskDependency) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"outcome":  types.StringType,
			"task_key": types.StringType,
		},
	}
}

type TaskEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	NoAlertForSkippedRuns types.Bool `tfsdk:"no_alert_for_skipped_runs" tf:"optional"`
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded types.List `tfsdk:"on_duration_warning_threshold_exceeded" tf:"optional"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure types.List `tfsdk:"on_failure" tf:"optional"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart types.List `tfsdk:"on_start" tf:"optional"`
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	OnStreamingBacklogExceeded types.List `tfsdk:"on_streaming_backlog_exceeded" tf:"optional"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess types.List `tfsdk:"on_success" tf:"optional"`
}

func (newState *TaskEmailNotifications) SyncEffectiveFieldsDuringCreateOrUpdate(plan TaskEmailNotifications) {
}

func (newState *TaskEmailNotifications) SyncEffectiveFieldsDuringRead(existingState TaskEmailNotifications) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TaskEmailNotifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TaskEmailNotifications) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_duration_warning_threshold_exceeded": reflect.TypeOf(types.String{}),
		"on_failure":                             reflect.TypeOf(types.String{}),
		"on_start":                               reflect.TypeOf(types.String{}),
		"on_streaming_backlog_exceeded":          reflect.TypeOf(types.String{}),
		"on_success":                             reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of TaskEmailNotifications in the Terraform plugin framework type
// system.
func (a TaskEmailNotifications) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"no_alert_for_skipped_runs": types.BoolType,
			"on_duration_warning_threshold_exceeded": basetypes.ListType{
				ElemType: types.StringType,
			},
			"on_failure": basetypes.ListType{
				ElemType: types.StringType,
			},
			"on_start": basetypes.ListType{
				ElemType: types.StringType,
			},
			"on_streaming_backlog_exceeded": basetypes.ListType{
				ElemType: types.StringType,
			},
			"on_success": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type TaskNotificationSettings struct {
	// If true, do not send notifications to recipients specified in `on_start`
	// for the retried runs and do not send notifications to recipients
	// specified in `on_failure` until the last retry of the run.
	AlertOnLastAttempt types.Bool `tfsdk:"alert_on_last_attempt" tf:"optional"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns types.Bool `tfsdk:"no_alert_for_canceled_runs" tf:"optional"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns types.Bool `tfsdk:"no_alert_for_skipped_runs" tf:"optional"`
}

func (newState *TaskNotificationSettings) SyncEffectiveFieldsDuringCreateOrUpdate(plan TaskNotificationSettings) {
}

func (newState *TaskNotificationSettings) SyncEffectiveFieldsDuringRead(existingState TaskNotificationSettings) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TaskNotificationSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TaskNotificationSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of TaskNotificationSettings in the Terraform plugin framework type
// system.
func (a TaskNotificationSettings) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_on_last_attempt":      types.BoolType,
			"no_alert_for_canceled_runs": types.BoolType,
			"no_alert_for_skipped_runs":  types.BoolType,
		},
	}
}

type TerminationDetails struct {
	// The code indicates why the run was terminated. Additional codes might be
	// introduced in future releases. * `SUCCESS`: The run was completed
	// successfully. * `USER_CANCELED`: The run was successfully canceled during
	// execution by a user. * `CANCELED`: The run was canceled during execution
	// by the Databricks platform; for example, if the maximum run duration was
	// exceeded. * `SKIPPED`: Run was never executed, for example, if the
	// upstream task run failed, the dependency type condition was not met, or
	// there were no material tasks to execute. * `INTERNAL_ERROR`: The run
	// encountered an unexpected error. Refer to the state message for further
	// details. * `DRIVER_ERROR`: The run encountered an error while
	// communicating with the Spark Driver. * `CLUSTER_ERROR`: The run failed
	// due to a cluster error. Refer to the state message for further details. *
	// `REPOSITORY_CHECKOUT_FAILED`: Failed to complete the checkout due to an
	// error when communicating with the third party service. *
	// `INVALID_CLUSTER_REQUEST`: The run failed because it issued an invalid
	// request to start the cluster. * `WORKSPACE_RUN_LIMIT_EXCEEDED`: The
	// workspace has reached the quota for the maximum number of concurrent
	// active runs. Consider scheduling the runs over a larger time frame. *
	// `FEATURE_DISABLED`: The run failed because it tried to access a feature
	// unavailable for the workspace. * `CLUSTER_REQUEST_LIMIT_EXCEEDED`: The
	// number of cluster creation, start, and upsize requests have exceeded the
	// allotted rate limit. Consider spreading the run execution over a larger
	// time frame. * `STORAGE_ACCESS_ERROR`: The run failed due to an error when
	// accessing the customer blob storage. Refer to the state message for
	// further details. * `RUN_EXECUTION_ERROR`: The run was completed with task
	// failures. For more details, refer to the state message or run output. *
	// `UNAUTHORIZED_ERROR`: The run failed due to a permission issue while
	// accessing a resource. Refer to the state message for further details. *
	// `LIBRARY_INSTALLATION_ERROR`: The run failed while installing the
	// user-requested library. Refer to the state message for further details.
	// The causes might include, but are not limited to: The provided library is
	// invalid, there are insufficient permissions to install the library, and
	// so forth. * `MAX_CONCURRENT_RUNS_EXCEEDED`: The scheduled run exceeds the
	// limit of maximum concurrent runs set for the job. *
	// `MAX_SPARK_CONTEXTS_EXCEEDED`: The run is scheduled on a cluster that has
	// already reached the maximum number of contexts it is configured to
	// create. See: [Link]. * `RESOURCE_NOT_FOUND`: A resource necessary for run
	// execution does not exist. Refer to the state message for further details.
	// * `INVALID_RUN_CONFIGURATION`: The run failed due to an invalid
	// configuration. Refer to the state message for further details. *
	// `CLOUD_FAILURE`: The run failed due to a cloud provider issue. Refer to
	// the state message for further details. * `MAX_JOB_QUEUE_SIZE_EXCEEDED`:
	// The run was skipped due to reaching the job level queue size limit.
	//
	// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
	Code types.String `tfsdk:"code" tf:"optional"`
	// A descriptive message with the termination details. This field is
	// unstructured and the format might change.
	Message types.String `tfsdk:"message" tf:"optional"`
	// * `SUCCESS`: The run terminated without any issues * `INTERNAL_ERROR`: An
	// error occurred in the Databricks platform. Please look at the [status
	// page] or contact support if the issue persists. * `CLIENT_ERROR`: The run
	// was terminated because of an error caused by user input or the job
	// configuration. * `CLOUD_FAILURE`: The run was terminated because of an
	// issue with your cloud provider.
	//
	// [status page]: https://status.databricks.com/
	Type types.String `tfsdk:"type" tf:"optional"`
}

func (newState *TerminationDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan TerminationDetails) {
}

func (newState *TerminationDetails) SyncEffectiveFieldsDuringRead(existingState TerminationDetails) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TerminationDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TerminationDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of TerminationDetails in the Terraform plugin framework type
// system.
func (a TerminationDetails) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"code":    types.StringType,
			"message": types.StringType,
			"type":    types.StringType,
		},
	}
}

// Additional details about what triggered the run
type TriggerInfo struct {
	// The run id of the Run Job task run
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
}

func (newState *TriggerInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan TriggerInfo) {
}

func (newState *TriggerInfo) SyncEffectiveFieldsDuringRead(existingState TriggerInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TriggerInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TriggerInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of TriggerInfo in the Terraform plugin framework type
// system.
func (a TriggerInfo) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

type TriggerSettings struct {
	// File arrival trigger settings.
	FileArrival types.List `tfsdk:"file_arrival" tf:"optional,object"`
	// Whether this trigger is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// Periodic trigger settings.
	Periodic types.List `tfsdk:"periodic" tf:"optional,object"`
	// Old table trigger settings name. Deprecated in favor of `table_update`.
	Table types.List `tfsdk:"table" tf:"optional,object"`

	TableUpdate types.List `tfsdk:"table_update" tf:"optional,object"`
}

func (newState *TriggerSettings) SyncEffectiveFieldsDuringCreateOrUpdate(plan TriggerSettings) {
}

func (newState *TriggerSettings) SyncEffectiveFieldsDuringRead(existingState TriggerSettings) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TriggerSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TriggerSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_arrival": reflect.TypeOf(FileArrivalTriggerConfiguration{}),
		"periodic":     reflect.TypeOf(PeriodicTriggerConfiguration{}),
		"table":        reflect.TypeOf(TableUpdateTriggerConfiguration{}),
		"table_update": reflect.TypeOf(TableUpdateTriggerConfiguration{}),
	}
}

// ToObjectType returns the representation of TriggerSettings in the Terraform plugin framework type
// system.
func (a TriggerSettings) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_arrival": basetypes.ListType{
				ElemType: FileArrivalTriggerConfiguration{}.ToObjectType(ctx),
			},
			"pause_status": types.StringType,
			"periodic": basetypes.ListType{
				ElemType: PeriodicTriggerConfiguration{}.ToObjectType(ctx),
			},
			"table": basetypes.ListType{
				ElemType: TableUpdateTriggerConfiguration{}.ToObjectType(ctx),
			},
			"table_update": basetypes.ListType{
				ElemType: TableUpdateTriggerConfiguration{}.ToObjectType(ctx),
			},
		},
	}
}

type UpdateJob struct {
	// Remove top-level fields in the job settings. Removing nested fields is
	// not supported, except for tasks and job clusters (`tasks/task_1`). This
	// field is optional.
	FieldsToRemove types.List `tfsdk:"fields_to_remove" tf:"optional"`
	// The canonical identifier of the job to update. This field is required.
	JobId types.Int64 `tfsdk:"job_id" tf:""`
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
	NewSettings types.List `tfsdk:"new_settings" tf:"optional,object"`
}

func (newState *UpdateJob) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateJob) {
}

func (newState *UpdateJob) SyncEffectiveFieldsDuringRead(existingState UpdateJob) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateJob) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fields_to_remove": reflect.TypeOf(types.String{}),
		"new_settings":     reflect.TypeOf(JobSettings{}),
	}
}

// ToObjectType returns the representation of UpdateJob in the Terraform plugin framework type
// system.
func (a UpdateJob) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fields_to_remove": basetypes.ListType{
				ElemType: types.StringType,
			},
			"job_id": types.Int64Type,
			"new_settings": basetypes.ListType{
				ElemType: JobSettings{}.ToObjectType(ctx),
			},
		},
	}
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of UpdateResponse in the Terraform plugin framework type
// system.
func (a UpdateResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ViewItem struct {
	// Content of the view.
	Content types.String `tfsdk:"content" tf:"optional"`
	// Name of the view item. In the case of code view, it would be the
	// notebook’s name. In the case of dashboard view, it would be the
	// dashboard’s name.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Type of the view item.
	Type types.String `tfsdk:"type" tf:"optional"`
}

func (newState *ViewItem) SyncEffectiveFieldsDuringCreateOrUpdate(plan ViewItem) {
}

func (newState *ViewItem) SyncEffectiveFieldsDuringRead(existingState ViewItem) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ViewItem.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ViewItem) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ViewItem in the Terraform plugin framework type
// system.
func (a ViewItem) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"name":    types.StringType,
			"type":    types.StringType,
		},
	}
}

type Webhook struct {
	Id types.String `tfsdk:"id" tf:""`
}

func (newState *Webhook) SyncEffectiveFieldsDuringCreateOrUpdate(plan Webhook) {
}

func (newState *Webhook) SyncEffectiveFieldsDuringRead(existingState Webhook) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Webhook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Webhook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of Webhook in the Terraform plugin framework type
// system.
func (a Webhook) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type WebhookNotifications struct {
	// An optional list of system notification IDs to call when the duration of
	// a run exceeds the threshold specified for the `RUN_DURATION_SECONDS`
	// metric in the `health` field. A maximum of 3 destinations can be
	// specified for the `on_duration_warning_threshold_exceeded` property.
	OnDurationWarningThresholdExceeded types.List `tfsdk:"on_duration_warning_threshold_exceeded" tf:"optional"`
	// An optional list of system notification IDs to call when the run fails. A
	// maximum of 3 destinations can be specified for the `on_failure` property.
	OnFailure types.List `tfsdk:"on_failure" tf:"optional"`
	// An optional list of system notification IDs to call when the run starts.
	// A maximum of 3 destinations can be specified for the `on_start` property.
	OnStart types.List `tfsdk:"on_start" tf:"optional"`
	// An optional list of system notification IDs to call when any streaming
	// backlog thresholds are exceeded for any stream. Streaming backlog
	// thresholds can be set in the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes. A maximum of 3 destinations
	// can be specified for the `on_streaming_backlog_exceeded` property.
	OnStreamingBacklogExceeded types.List `tfsdk:"on_streaming_backlog_exceeded" tf:"optional"`
	// An optional list of system notification IDs to call when the run
	// completes successfully. A maximum of 3 destinations can be specified for
	// the `on_success` property.
	OnSuccess types.List `tfsdk:"on_success" tf:"optional"`
}

func (newState *WebhookNotifications) SyncEffectiveFieldsDuringCreateOrUpdate(plan WebhookNotifications) {
}

func (newState *WebhookNotifications) SyncEffectiveFieldsDuringRead(existingState WebhookNotifications) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WebhookNotifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WebhookNotifications) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_duration_warning_threshold_exceeded": reflect.TypeOf(Webhook{}),
		"on_failure":                             reflect.TypeOf(Webhook{}),
		"on_start":                               reflect.TypeOf(Webhook{}),
		"on_streaming_backlog_exceeded":          reflect.TypeOf(Webhook{}),
		"on_success":                             reflect.TypeOf(Webhook{}),
	}
}

// ToObjectType returns the representation of WebhookNotifications in the Terraform plugin framework type
// system.
func (a WebhookNotifications) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"on_duration_warning_threshold_exceeded": basetypes.ListType{
				ElemType: Webhook{}.ToObjectType(ctx),
			},
			"on_failure": basetypes.ListType{
				ElemType: Webhook{}.ToObjectType(ctx),
			},
			"on_start": basetypes.ListType{
				ElemType: Webhook{}.ToObjectType(ctx),
			},
			"on_streaming_backlog_exceeded": basetypes.ListType{
				ElemType: Webhook{}.ToObjectType(ctx),
			},
			"on_success": basetypes.ListType{
				ElemType: Webhook{}.ToObjectType(ctx),
			},
		},
	}
}
