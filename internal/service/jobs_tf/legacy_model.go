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

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"

	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type BaseJob_SdkV2 struct {
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
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id" tf:"computed"`
	// The canonical identifier for this job.
	JobId types.Int64 `tfsdk:"job_id" tf:"optional"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings types.List `tfsdk:"settings" tf:"optional,object"`
}

func (newState *BaseJob_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BaseJob_SdkV2) {
}

func (newState *BaseJob_SdkV2) SyncEffectiveFieldsDuringRead(existingState BaseJob_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BaseJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BaseJob_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings": reflect.TypeOf(JobSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BaseJob_SdkV2
// only implements ToObjectValue() and Type().
func (o BaseJob_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_time":               o.CreatedTime,
			"creator_user_name":          o.CreatorUserName,
			"effective_budget_policy_id": o.EffectiveBudgetPolicyId,
			"job_id":                     o.JobId,
			"settings":                   o.Settings,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BaseJob_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_time":               types.Int64Type,
			"creator_user_name":          types.StringType,
			"effective_budget_policy_id": types.StringType,
			"job_id":                     types.Int64Type,
			"settings": basetypes.ListType{
				ElemType: JobSettings{}.Type(ctx),
			},
		},
	}
}

// GetSettings returns the value of the Settings field in BaseJob_SdkV2 as
// a JobSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseJob_SdkV2) GetSettings(ctx context.Context) (JobSettings_SdkV2, bool) {
	var e JobSettings_SdkV2
	if o.Settings.IsNull() || o.Settings.IsUnknown() {
		return e, false
	}
	var v []JobSettings_SdkV2
	d := o.Settings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSettings sets the value of the Settings field in BaseJob_SdkV2.
func (o *BaseJob_SdkV2) SetSettings(ctx context.Context, v JobSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings"]
	o.Settings = types.ListValueMust(t, vs)
}

type BaseRun_SdkV2 struct {
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
	// * `CONTINUOUS_RESTART`: Indicates a run created by user to manually
	// restart a continuous job run.
	Trigger types.String `tfsdk:"trigger" tf:"optional"`
	// Additional details about what triggered the run
	TriggerInfo types.List `tfsdk:"trigger_info" tf:"optional,object"`
}

func (newState *BaseRun_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BaseRun_SdkV2) {
}

func (newState *BaseRun_SdkV2) SyncEffectiveFieldsDuringRead(existingState BaseRun_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BaseRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BaseRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster_instance":      reflect.TypeOf(ClusterInstance_SdkV2{}),
		"cluster_spec":          reflect.TypeOf(ClusterSpec_SdkV2{}),
		"git_source":            reflect.TypeOf(GitSource_SdkV2{}),
		"job_clusters":          reflect.TypeOf(JobCluster_SdkV2{}),
		"job_parameters":        reflect.TypeOf(JobParameter_SdkV2{}),
		"overriding_parameters": reflect.TypeOf(RunParameters_SdkV2{}),
		"repair_history":        reflect.TypeOf(RepairHistoryItem_SdkV2{}),
		"schedule":              reflect.TypeOf(CronSchedule_SdkV2{}),
		"state":                 reflect.TypeOf(RunState_SdkV2{}),
		"status":                reflect.TypeOf(RunStatus_SdkV2{}),
		"tasks":                 reflect.TypeOf(RunTask_SdkV2{}),
		"trigger_info":          reflect.TypeOf(TriggerInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BaseRun_SdkV2
// only implements ToObjectValue() and Type().
func (o BaseRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attempt_number":          o.AttemptNumber,
			"cleanup_duration":        o.CleanupDuration,
			"cluster_instance":        o.ClusterInstance,
			"cluster_spec":            o.ClusterSpec,
			"creator_user_name":       o.CreatorUserName,
			"description":             o.Description,
			"end_time":                o.EndTime,
			"execution_duration":      o.ExecutionDuration,
			"git_source":              o.GitSource,
			"job_clusters":            o.JobClusters,
			"job_id":                  o.JobId,
			"job_parameters":          o.JobParameters,
			"job_run_id":              o.JobRunId,
			"number_in_job":           o.NumberInJob,
			"original_attempt_run_id": o.OriginalAttemptRunId,
			"overriding_parameters":   o.OverridingParameters,
			"queue_duration":          o.QueueDuration,
			"repair_history":          o.RepairHistory,
			"run_duration":            o.RunDuration,
			"run_id":                  o.RunId,
			"run_name":                o.RunName,
			"run_page_url":            o.RunPageUrl,
			"run_type":                o.RunType,
			"schedule":                o.Schedule,
			"setup_duration":          o.SetupDuration,
			"start_time":              o.StartTime,
			"state":                   o.State,
			"status":                  o.Status,
			"tasks":                   o.Tasks,
			"trigger":                 o.Trigger,
			"trigger_info":            o.TriggerInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BaseRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attempt_number":   types.Int64Type,
			"cleanup_duration": types.Int64Type,
			"cluster_instance": basetypes.ListType{
				ElemType: ClusterInstance{}.Type(ctx),
			},
			"cluster_spec": basetypes.ListType{
				ElemType: ClusterSpec{}.Type(ctx),
			},
			"creator_user_name":  types.StringType,
			"description":        types.StringType,
			"end_time":           types.Int64Type,
			"execution_duration": types.Int64Type,
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.Type(ctx),
			},
			"job_clusters": basetypes.ListType{
				ElemType: JobCluster{}.Type(ctx),
			},
			"job_id": types.Int64Type,
			"job_parameters": basetypes.ListType{
				ElemType: JobParameter{}.Type(ctx),
			},
			"job_run_id":              types.Int64Type,
			"number_in_job":           types.Int64Type,
			"original_attempt_run_id": types.Int64Type,
			"overriding_parameters": basetypes.ListType{
				ElemType: RunParameters{}.Type(ctx),
			},
			"queue_duration": types.Int64Type,
			"repair_history": basetypes.ListType{
				ElemType: RepairHistoryItem{}.Type(ctx),
			},
			"run_duration": types.Int64Type,
			"run_id":       types.Int64Type,
			"run_name":     types.StringType,
			"run_page_url": types.StringType,
			"run_type":     types.StringType,
			"schedule": basetypes.ListType{
				ElemType: CronSchedule{}.Type(ctx),
			},
			"setup_duration": types.Int64Type,
			"start_time":     types.Int64Type,
			"state": basetypes.ListType{
				ElemType: RunState{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: RunStatus{}.Type(ctx),
			},
			"tasks": basetypes.ListType{
				ElemType: RunTask{}.Type(ctx),
			},
			"trigger": types.StringType,
			"trigger_info": basetypes.ListType{
				ElemType: TriggerInfo{}.Type(ctx),
			},
		},
	}
}

// GetClusterInstance returns the value of the ClusterInstance field in BaseRun_SdkV2 as
// a ClusterInstance_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetClusterInstance(ctx context.Context) (ClusterInstance_SdkV2, bool) {
	var e ClusterInstance_SdkV2
	if o.ClusterInstance.IsNull() || o.ClusterInstance.IsUnknown() {
		return e, false
	}
	var v []ClusterInstance_SdkV2
	d := o.ClusterInstance.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterInstance sets the value of the ClusterInstance field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetClusterInstance(ctx context.Context, v ClusterInstance_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_instance"]
	o.ClusterInstance = types.ListValueMust(t, vs)
}

// GetClusterSpec returns the value of the ClusterSpec field in BaseRun_SdkV2 as
// a ClusterSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetClusterSpec(ctx context.Context) (ClusterSpec_SdkV2, bool) {
	var e ClusterSpec_SdkV2
	if o.ClusterSpec.IsNull() || o.ClusterSpec.IsUnknown() {
		return e, false
	}
	var v []ClusterSpec_SdkV2
	d := o.ClusterSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterSpec sets the value of the ClusterSpec field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetClusterSpec(ctx context.Context, v ClusterSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_spec"]
	o.ClusterSpec = types.ListValueMust(t, vs)
}

// GetGitSource returns the value of the GitSource field in BaseRun_SdkV2 as
// a GitSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetGitSource(ctx context.Context) (GitSource_SdkV2, bool) {
	var e GitSource_SdkV2
	if o.GitSource.IsNull() || o.GitSource.IsUnknown() {
		return e, false
	}
	var v []GitSource_SdkV2
	d := o.GitSource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitSource sets the value of the GitSource field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetGitSource(ctx context.Context, v GitSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["git_source"]
	o.GitSource = types.ListValueMust(t, vs)
}

// GetJobClusters returns the value of the JobClusters field in BaseRun_SdkV2 as
// a slice of JobCluster_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetJobClusters(ctx context.Context) ([]JobCluster_SdkV2, bool) {
	if o.JobClusters.IsNull() || o.JobClusters.IsUnknown() {
		return nil, false
	}
	var v []JobCluster_SdkV2
	d := o.JobClusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobClusters sets the value of the JobClusters field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetJobClusters(ctx context.Context, v []JobCluster_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobClusters = types.ListValueMust(t, vs)
}

// GetJobParameters returns the value of the JobParameters field in BaseRun_SdkV2 as
// a slice of JobParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetJobParameters(ctx context.Context) ([]JobParameter_SdkV2, bool) {
	if o.JobParameters.IsNull() || o.JobParameters.IsUnknown() {
		return nil, false
	}
	var v []JobParameter_SdkV2
	d := o.JobParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobParameters sets the value of the JobParameters field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetJobParameters(ctx context.Context, v []JobParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobParameters = types.ListValueMust(t, vs)
}

// GetOverridingParameters returns the value of the OverridingParameters field in BaseRun_SdkV2 as
// a RunParameters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetOverridingParameters(ctx context.Context) (RunParameters_SdkV2, bool) {
	var e RunParameters_SdkV2
	if o.OverridingParameters.IsNull() || o.OverridingParameters.IsUnknown() {
		return e, false
	}
	var v []RunParameters_SdkV2
	d := o.OverridingParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOverridingParameters sets the value of the OverridingParameters field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetOverridingParameters(ctx context.Context, v RunParameters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["overriding_parameters"]
	o.OverridingParameters = types.ListValueMust(t, vs)
}

// GetRepairHistory returns the value of the RepairHistory field in BaseRun_SdkV2 as
// a slice of RepairHistoryItem_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetRepairHistory(ctx context.Context) ([]RepairHistoryItem_SdkV2, bool) {
	if o.RepairHistory.IsNull() || o.RepairHistory.IsUnknown() {
		return nil, false
	}
	var v []RepairHistoryItem_SdkV2
	d := o.RepairHistory.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRepairHistory sets the value of the RepairHistory field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetRepairHistory(ctx context.Context, v []RepairHistoryItem_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["repair_history"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RepairHistory = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in BaseRun_SdkV2 as
// a CronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetSchedule(ctx context.Context) (CronSchedule_SdkV2, bool) {
	var e CronSchedule_SdkV2
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []CronSchedule_SdkV2
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetSchedule(ctx context.Context, v CronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}

// GetState returns the value of the State field in BaseRun_SdkV2 as
// a RunState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetState(ctx context.Context) (RunState_SdkV2, bool) {
	var e RunState_SdkV2
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v []RunState_SdkV2
	d := o.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetState(ctx context.Context, v RunState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	o.State = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in BaseRun_SdkV2 as
// a RunStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetStatus(ctx context.Context) (RunStatus_SdkV2, bool) {
	var e RunStatus_SdkV2
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []RunStatus_SdkV2
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetStatus(ctx context.Context, v RunStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
}

// GetTasks returns the value of the Tasks field in BaseRun_SdkV2 as
// a slice of RunTask_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetTasks(ctx context.Context) ([]RunTask_SdkV2, bool) {
	if o.Tasks.IsNull() || o.Tasks.IsUnknown() {
		return nil, false
	}
	var v []RunTask_SdkV2
	d := o.Tasks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTasks sets the value of the Tasks field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetTasks(ctx context.Context, v []RunTask_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tasks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tasks = types.ListValueMust(t, vs)
}

// GetTriggerInfo returns the value of the TriggerInfo field in BaseRun_SdkV2 as
// a TriggerInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *BaseRun_SdkV2) GetTriggerInfo(ctx context.Context) (TriggerInfo_SdkV2, bool) {
	var e TriggerInfo_SdkV2
	if o.TriggerInfo.IsNull() || o.TriggerInfo.IsUnknown() {
		return e, false
	}
	var v []TriggerInfo_SdkV2
	d := o.TriggerInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTriggerInfo sets the value of the TriggerInfo field in BaseRun_SdkV2.
func (o *BaseRun_SdkV2) SetTriggerInfo(ctx context.Context, v TriggerInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["trigger_info"]
	o.TriggerInfo = types.ListValueMust(t, vs)
}

type CancelAllRuns_SdkV2 struct {
	// Optional boolean parameter to cancel all queued runs. If no job_id is
	// provided, all queued runs in the workspace are canceled.
	AllQueuedRuns types.Bool `tfsdk:"all_queued_runs" tf:"optional"`
	// The canonical identifier of the job to cancel all runs of.
	JobId types.Int64 `tfsdk:"job_id" tf:"optional"`
}

func (newState *CancelAllRuns_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelAllRuns_SdkV2) {
}

func (newState *CancelAllRuns_SdkV2) SyncEffectiveFieldsDuringRead(existingState CancelAllRuns_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelAllRuns.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelAllRuns_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelAllRuns_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelAllRuns_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_queued_runs": o.AllQueuedRuns,
			"job_id":          o.JobId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CancelAllRuns_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_queued_runs": types.BoolType,
			"job_id":          types.Int64Type,
		},
	}
}

type CancelAllRunsResponse_SdkV2 struct {
}

func (newState *CancelAllRunsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelAllRunsResponse_SdkV2) {
}

func (newState *CancelAllRunsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CancelAllRunsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelAllRunsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelAllRunsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelAllRunsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelAllRunsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CancelAllRunsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CancelRun_SdkV2 struct {
	// This field is required.
	RunId types.Int64 `tfsdk:"run_id" tf:""`
}

func (newState *CancelRun_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelRun_SdkV2) {
}

func (newState *CancelRun_SdkV2) SyncEffectiveFieldsDuringRead(existingState CancelRun_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelRun_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CancelRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

type CancelRunResponse_SdkV2 struct {
}

func (newState *CancelRunResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelRunResponse_SdkV2) {
}

func (newState *CancelRunResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CancelRunResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CancelRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Stores the run state of the clean rooms notebook task.
type CleanRoomTaskRunState_SdkV2 struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response.
	LifeCycleState types.String `tfsdk:"life_cycle_state" tf:"optional"`
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states.
	ResultState types.String `tfsdk:"result_state" tf:"optional"`
}

func (newState *CleanRoomTaskRunState_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomTaskRunState_SdkV2) {
}

func (newState *CleanRoomTaskRunState_SdkV2) SyncEffectiveFieldsDuringRead(existingState CleanRoomTaskRunState_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomTaskRunState.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomTaskRunState_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomTaskRunState_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomTaskRunState_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"life_cycle_state": o.LifeCycleState,
			"result_state":     o.ResultState,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomTaskRunState_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"life_cycle_state": types.StringType,
			"result_state":     types.StringType,
		},
	}
}

type CleanRoomsNotebookTask_SdkV2 struct {
	// The clean room that the notebook belongs to.
	CleanRoomName types.String `tfsdk:"clean_room_name" tf:""`
	// Checksum to validate the freshness of the notebook resource (i.e. the
	// notebook being run is the latest version). It can be fetched by calling
	// the :method:cleanroomassets/get API.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// Base parameters to be used for the clean room notebook job.
	NotebookBaseParameters types.Map `tfsdk:"notebook_base_parameters" tf:"optional"`
	// Name of the notebook being run.
	NotebookName types.String `tfsdk:"notebook_name" tf:""`
}

func (newState *CleanRoomsNotebookTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomsNotebookTask_SdkV2) {
}

func (newState *CleanRoomsNotebookTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState CleanRoomsNotebookTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomsNotebookTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomsNotebookTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notebook_base_parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomsNotebookTask_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomsNotebookTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name":          o.CleanRoomName,
			"etag":                     o.Etag,
			"notebook_base_parameters": o.NotebookBaseParameters,
			"notebook_name":            o.NotebookName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomsNotebookTask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"etag":            types.StringType,
			"notebook_base_parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"notebook_name": types.StringType,
		},
	}
}

// GetNotebookBaseParameters returns the value of the NotebookBaseParameters field in CleanRoomsNotebookTask_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomsNotebookTask_SdkV2) GetNotebookBaseParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.NotebookBaseParameters.IsNull() || o.NotebookBaseParameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.NotebookBaseParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookBaseParameters sets the value of the NotebookBaseParameters field in CleanRoomsNotebookTask_SdkV2.
func (o *CleanRoomsNotebookTask_SdkV2) SetNotebookBaseParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_base_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NotebookBaseParameters = types.MapValueMust(t, vs)
}

type ClusterInstance_SdkV2 struct {
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

func (newState *ClusterInstance_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterInstance_SdkV2) {
}

func (newState *ClusterInstance_SdkV2) SyncEffectiveFieldsDuringRead(existingState ClusterInstance_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterInstance.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterInstance_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterInstance_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterInstance_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cluster_id":       o.ClusterId,
			"spark_context_id": o.SparkContextId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterInstance_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cluster_id":       types.StringType,
			"spark_context_id": types.StringType,
		},
	}
}

type ClusterSpec_SdkV2 struct {
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
	NewCluster types.List `tfsdk:"new_cluster" tf:"optional,object"`
}

func (newState *ClusterSpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterSpec_SdkV2) {
}

func (newState *ClusterSpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState ClusterSpec_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ClusterSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"library":     reflect.TypeOf(compute_tf.Library_SdkV2{}),
		"new_cluster": reflect.TypeOf(compute_tf.ClusterSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"existing_cluster_id": o.ExistingClusterId,
			"job_cluster_key":     o.JobClusterKey,
			"library":             o.Libraries,
			"new_cluster":         o.NewCluster,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"existing_cluster_id": types.StringType,
			"job_cluster_key":     types.StringType,
			"library": basetypes.ListType{
				ElemType: compute_tf.Library{}.Type(ctx),
			},
			"new_cluster": basetypes.ListType{
				ElemType: compute_tf.ClusterSpec{}.Type(ctx),
			},
		},
	}
}

// GetLibraries returns the value of the Libraries field in ClusterSpec_SdkV2 as
// a slice of compute_tf.Library_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetLibraries(ctx context.Context) ([]compute_tf.Library_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []compute_tf.Library_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetLibraries(ctx context.Context, v []compute_tf.Library_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["library"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

// GetNewCluster returns the value of the NewCluster field in ClusterSpec_SdkV2 as
// a compute_tf.ClusterSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterSpec_SdkV2) GetNewCluster(ctx context.Context) (compute_tf.ClusterSpec_SdkV2, bool) {
	var e compute_tf.ClusterSpec_SdkV2
	if o.NewCluster.IsNull() || o.NewCluster.IsUnknown() {
		return e, false
	}
	var v []compute_tf.ClusterSpec_SdkV2
	d := o.NewCluster.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNewCluster sets the value of the NewCluster field in ClusterSpec_SdkV2.
func (o *ClusterSpec_SdkV2) SetNewCluster(ctx context.Context, v compute_tf.ClusterSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["new_cluster"]
	o.NewCluster = types.ListValueMust(t, vs)
}

type ConditionTask_SdkV2 struct {
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

func (newState *ConditionTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ConditionTask_SdkV2) {
}

func (newState *ConditionTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState ConditionTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConditionTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ConditionTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConditionTask_SdkV2
// only implements ToObjectValue() and Type().
func (o ConditionTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"left":  o.Left,
			"op":    o.Op,
			"right": o.Right,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ConditionTask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"left":  types.StringType,
			"op":    types.StringType,
			"right": types.StringType,
		},
	}
}

type Continuous_SdkV2 struct {
	// Indicate whether the continuous execution of the job is paused or not.
	// Defaults to UNPAUSED.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
}

func (newState *Continuous_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Continuous_SdkV2) {
}

func (newState *Continuous_SdkV2) SyncEffectiveFieldsDuringRead(existingState Continuous_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Continuous.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Continuous_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Continuous_SdkV2
// only implements ToObjectValue() and Type().
func (o Continuous_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pause_status": o.PauseStatus,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Continuous_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pause_status": types.StringType,
		},
	}
}

type CreateJob_SdkV2 struct {
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
	// Write-only setting. Specifies the user or service principal that the job
	// runs as. If not specified, the job runs as the user who created the job.
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

func (newState *CreateJob_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateJob_SdkV2) {
}

func (newState *CreateJob_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateJob_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateJob_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list":   reflect.TypeOf(JobAccessControlRequest_SdkV2{}),
		"continuous":            reflect.TypeOf(Continuous_SdkV2{}),
		"deployment":            reflect.TypeOf(JobDeployment_SdkV2{}),
		"email_notifications":   reflect.TypeOf(JobEmailNotifications_SdkV2{}),
		"environment":           reflect.TypeOf(JobEnvironment_SdkV2{}),
		"git_source":            reflect.TypeOf(GitSource_SdkV2{}),
		"health":                reflect.TypeOf(JobsHealthRules_SdkV2{}),
		"job_cluster":           reflect.TypeOf(JobCluster_SdkV2{}),
		"notification_settings": reflect.TypeOf(JobNotificationSettings_SdkV2{}),
		"parameter":             reflect.TypeOf(JobParameterDefinition_SdkV2{}),
		"queue":                 reflect.TypeOf(QueueSettings_SdkV2{}),
		"run_as":                reflect.TypeOf(JobRunAs_SdkV2{}),
		"schedule":              reflect.TypeOf(CronSchedule_SdkV2{}),
		"tags":                  reflect.TypeOf(types.String{}),
		"task":                  reflect.TypeOf(Task_SdkV2{}),
		"trigger":               reflect.TypeOf(TriggerSettings_SdkV2{}),
		"webhook_notifications": reflect.TypeOf(WebhookNotifications_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateJob_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateJob_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list":   o.AccessControlList,
			"budget_policy_id":      o.BudgetPolicyId,
			"continuous":            o.Continuous,
			"deployment":            o.Deployment,
			"description":           o.Description,
			"edit_mode":             o.EditMode,
			"email_notifications":   o.EmailNotifications,
			"environment":           o.Environments,
			"format":                o.Format,
			"git_source":            o.GitSource,
			"health":                o.Health,
			"job_cluster":           o.JobClusters,
			"max_concurrent_runs":   o.MaxConcurrentRuns,
			"name":                  o.Name,
			"notification_settings": o.NotificationSettings,
			"parameter":             o.Parameters,
			"queue":                 o.Queue,
			"run_as":                o.RunAs,
			"schedule":              o.Schedule,
			"tags":                  o.Tags,
			"task":                  o.Tasks,
			"timeout_seconds":       o.TimeoutSeconds,
			"trigger":               o.Trigger,
			"webhook_notifications": o.WebhookNotifications,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateJob_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: JobAccessControlRequest{}.Type(ctx),
			},
			"budget_policy_id": types.StringType,
			"continuous": basetypes.ListType{
				ElemType: Continuous{}.Type(ctx),
			},
			"deployment": basetypes.ListType{
				ElemType: JobDeployment{}.Type(ctx),
			},
			"description": types.StringType,
			"edit_mode":   types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: JobEmailNotifications{}.Type(ctx),
			},
			"environment": basetypes.ListType{
				ElemType: JobEnvironment{}.Type(ctx),
			},
			"format": types.StringType,
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.Type(ctx),
			},
			"health": basetypes.ListType{
				ElemType: JobsHealthRules{}.Type(ctx),
			},
			"job_cluster": basetypes.ListType{
				ElemType: JobCluster{}.Type(ctx),
			},
			"max_concurrent_runs": types.Int64Type,
			"name":                types.StringType,
			"notification_settings": basetypes.ListType{
				ElemType: JobNotificationSettings{}.Type(ctx),
			},
			"parameter": basetypes.ListType{
				ElemType: JobParameterDefinition{}.Type(ctx),
			},
			"queue": basetypes.ListType{
				ElemType: QueueSettings{}.Type(ctx),
			},
			"run_as": basetypes.ListType{
				ElemType: JobRunAs{}.Type(ctx),
			},
			"schedule": basetypes.ListType{
				ElemType: CronSchedule{}.Type(ctx),
			},
			"tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"task": basetypes.ListType{
				ElemType: Task{}.Type(ctx),
			},
			"timeout_seconds": types.Int64Type,
			"trigger": basetypes.ListType{
				ElemType: TriggerSettings{}.Type(ctx),
			},
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.Type(ctx),
			},
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in CreateJob_SdkV2 as
// a slice of JobAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetAccessControlList(ctx context.Context) ([]JobAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []JobAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetAccessControlList(ctx context.Context, v []JobAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

// GetContinuous returns the value of the Continuous field in CreateJob_SdkV2 as
// a Continuous_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetContinuous(ctx context.Context) (Continuous_SdkV2, bool) {
	var e Continuous_SdkV2
	if o.Continuous.IsNull() || o.Continuous.IsUnknown() {
		return e, false
	}
	var v []Continuous_SdkV2
	d := o.Continuous.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetContinuous sets the value of the Continuous field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetContinuous(ctx context.Context, v Continuous_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["continuous"]
	o.Continuous = types.ListValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in CreateJob_SdkV2 as
// a JobDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetDeployment(ctx context.Context) (JobDeployment_SdkV2, bool) {
	var e JobDeployment_SdkV2
	if o.Deployment.IsNull() || o.Deployment.IsUnknown() {
		return e, false
	}
	var v []JobDeployment_SdkV2
	d := o.Deployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeployment sets the value of the Deployment field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetDeployment(ctx context.Context, v JobDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["deployment"]
	o.Deployment = types.ListValueMust(t, vs)
}

// GetEmailNotifications returns the value of the EmailNotifications field in CreateJob_SdkV2 as
// a JobEmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetEmailNotifications(ctx context.Context) (JobEmailNotifications_SdkV2, bool) {
	var e JobEmailNotifications_SdkV2
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []JobEmailNotifications_SdkV2
	d := o.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetEmailNotifications(ctx context.Context, v JobEmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	o.EmailNotifications = types.ListValueMust(t, vs)
}

// GetEnvironments returns the value of the Environments field in CreateJob_SdkV2 as
// a slice of JobEnvironment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetEnvironments(ctx context.Context) ([]JobEnvironment_SdkV2, bool) {
	if o.Environments.IsNull() || o.Environments.IsUnknown() {
		return nil, false
	}
	var v []JobEnvironment_SdkV2
	d := o.Environments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironments sets the value of the Environments field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetEnvironments(ctx context.Context, v []JobEnvironment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Environments = types.ListValueMust(t, vs)
}

// GetGitSource returns the value of the GitSource field in CreateJob_SdkV2 as
// a GitSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetGitSource(ctx context.Context) (GitSource_SdkV2, bool) {
	var e GitSource_SdkV2
	if o.GitSource.IsNull() || o.GitSource.IsUnknown() {
		return e, false
	}
	var v []GitSource_SdkV2
	d := o.GitSource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitSource sets the value of the GitSource field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetGitSource(ctx context.Context, v GitSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["git_source"]
	o.GitSource = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in CreateJob_SdkV2 as
// a JobsHealthRules_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetHealth(ctx context.Context) (JobsHealthRules_SdkV2, bool) {
	var e JobsHealthRules_SdkV2
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v []JobsHealthRules_SdkV2
	d := o.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetHealth(ctx context.Context, v JobsHealthRules_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	o.Health = types.ListValueMust(t, vs)
}

// GetJobClusters returns the value of the JobClusters field in CreateJob_SdkV2 as
// a slice of JobCluster_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetJobClusters(ctx context.Context) ([]JobCluster_SdkV2, bool) {
	if o.JobClusters.IsNull() || o.JobClusters.IsUnknown() {
		return nil, false
	}
	var v []JobCluster_SdkV2
	d := o.JobClusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobClusters sets the value of the JobClusters field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetJobClusters(ctx context.Context, v []JobCluster_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_cluster"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobClusters = types.ListValueMust(t, vs)
}

// GetNotificationSettings returns the value of the NotificationSettings field in CreateJob_SdkV2 as
// a JobNotificationSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetNotificationSettings(ctx context.Context) (JobNotificationSettings_SdkV2, bool) {
	var e JobNotificationSettings_SdkV2
	if o.NotificationSettings.IsNull() || o.NotificationSettings.IsUnknown() {
		return e, false
	}
	var v []JobNotificationSettings_SdkV2
	d := o.NotificationSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotificationSettings sets the value of the NotificationSettings field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetNotificationSettings(ctx context.Context, v JobNotificationSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notification_settings"]
	o.NotificationSettings = types.ListValueMust(t, vs)
}

// GetParameters returns the value of the Parameters field in CreateJob_SdkV2 as
// a slice of JobParameterDefinition_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetParameters(ctx context.Context) ([]JobParameterDefinition_SdkV2, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []JobParameterDefinition_SdkV2
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetParameters(ctx context.Context, v []JobParameterDefinition_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameter"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

// GetQueue returns the value of the Queue field in CreateJob_SdkV2 as
// a QueueSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetQueue(ctx context.Context) (QueueSettings_SdkV2, bool) {
	var e QueueSettings_SdkV2
	if o.Queue.IsNull() || o.Queue.IsUnknown() {
		return e, false
	}
	var v []QueueSettings_SdkV2
	d := o.Queue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueue sets the value of the Queue field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetQueue(ctx context.Context, v QueueSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["queue"]
	o.Queue = types.ListValueMust(t, vs)
}

// GetRunAs returns the value of the RunAs field in CreateJob_SdkV2 as
// a JobRunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetRunAs(ctx context.Context) (JobRunAs_SdkV2, bool) {
	var e JobRunAs_SdkV2
	if o.RunAs.IsNull() || o.RunAs.IsUnknown() {
		return e, false
	}
	var v []JobRunAs_SdkV2
	d := o.RunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunAs sets the value of the RunAs field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetRunAs(ctx context.Context, v JobRunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_as"]
	o.RunAs = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in CreateJob_SdkV2 as
// a CronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetSchedule(ctx context.Context) (CronSchedule_SdkV2, bool) {
	var e CronSchedule_SdkV2
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []CronSchedule_SdkV2
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetSchedule(ctx context.Context, v CronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateJob_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetTags(ctx context.Context) (map[string]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.MapValueMust(t, vs)
}

// GetTasks returns the value of the Tasks field in CreateJob_SdkV2 as
// a slice of Task_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetTasks(ctx context.Context) ([]Task_SdkV2, bool) {
	if o.Tasks.IsNull() || o.Tasks.IsUnknown() {
		return nil, false
	}
	var v []Task_SdkV2
	d := o.Tasks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTasks sets the value of the Tasks field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetTasks(ctx context.Context, v []Task_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["task"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tasks = types.ListValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in CreateJob_SdkV2 as
// a TriggerSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetTrigger(ctx context.Context) (TriggerSettings_SdkV2, bool) {
	var e TriggerSettings_SdkV2
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v []TriggerSettings_SdkV2
	d := o.Trigger.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrigger sets the value of the Trigger field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetTrigger(ctx context.Context, v TriggerSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["trigger"]
	o.Trigger = types.ListValueMust(t, vs)
}

// GetWebhookNotifications returns the value of the WebhookNotifications field in CreateJob_SdkV2 as
// a WebhookNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateJob_SdkV2) GetWebhookNotifications(ctx context.Context) (WebhookNotifications_SdkV2, bool) {
	var e WebhookNotifications_SdkV2
	if o.WebhookNotifications.IsNull() || o.WebhookNotifications.IsUnknown() {
		return e, false
	}
	var v []WebhookNotifications_SdkV2
	d := o.WebhookNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhookNotifications sets the value of the WebhookNotifications field in CreateJob_SdkV2.
func (o *CreateJob_SdkV2) SetWebhookNotifications(ctx context.Context, v WebhookNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook_notifications"]
	o.WebhookNotifications = types.ListValueMust(t, vs)
}

// Job was created successfully
type CreateResponse_SdkV2 struct {
	// The canonical identifier for the newly created job.
	JobId types.Int64 `tfsdk:"job_id" tf:"optional"`
}

func (newState *CreateResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateResponse_SdkV2) {
}

func (newState *CreateResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id": o.JobId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.Int64Type,
		},
	}
}

type CronSchedule_SdkV2 struct {
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

func (newState *CronSchedule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CronSchedule_SdkV2) {
}

func (newState *CronSchedule_SdkV2) SyncEffectiveFieldsDuringRead(existingState CronSchedule_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CronSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CronSchedule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CronSchedule_SdkV2
// only implements ToObjectValue() and Type().
func (o CronSchedule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pause_status":           o.PauseStatus,
			"quartz_cron_expression": o.QuartzCronExpression,
			"timezone_id":            o.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CronSchedule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pause_status":           types.StringType,
			"quartz_cron_expression": types.StringType,
			"timezone_id":            types.StringType,
		},
	}
}

type DbtOutput_SdkV2 struct {
	// An optional map of headers to send when retrieving the artifact from the
	// `artifacts_link`.
	ArtifactsHeaders types.Map `tfsdk:"artifacts_headers" tf:"optional"`
	// A pre-signed URL to download the (compressed) dbt artifacts. This link is
	// valid for a limited time (30 minutes). This information is only available
	// after the run has finished.
	ArtifactsLink types.String `tfsdk:"artifacts_link" tf:"optional"`
}

func (newState *DbtOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DbtOutput_SdkV2) {
}

func (newState *DbtOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState DbtOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DbtOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DbtOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"artifacts_headers": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DbtOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o DbtOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifacts_headers": o.ArtifactsHeaders,
			"artifacts_link":    o.ArtifactsLink,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DbtOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifacts_headers": basetypes.MapType{
				ElemType: types.StringType,
			},
			"artifacts_link": types.StringType,
		},
	}
}

// GetArtifactsHeaders returns the value of the ArtifactsHeaders field in DbtOutput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *DbtOutput_SdkV2) GetArtifactsHeaders(ctx context.Context) (map[string]types.String, bool) {
	if o.ArtifactsHeaders.IsNull() || o.ArtifactsHeaders.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.ArtifactsHeaders.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetArtifactsHeaders sets the value of the ArtifactsHeaders field in DbtOutput_SdkV2.
func (o *DbtOutput_SdkV2) SetArtifactsHeaders(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["artifacts_headers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ArtifactsHeaders = types.MapValueMust(t, vs)
}

type DbtTask_SdkV2 struct {
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

func (newState *DbtTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DbtTask_SdkV2) {
}

func (newState *DbtTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState DbtTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DbtTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DbtTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"commands": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DbtTask_SdkV2
// only implements ToObjectValue() and Type().
func (o DbtTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog":            o.Catalog,
			"commands":           o.Commands,
			"profiles_directory": o.ProfilesDirectory,
			"project_directory":  o.ProjectDirectory,
			"schema":             o.Schema,
			"source":             o.Source,
			"warehouse_id":       o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DbtTask_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetCommands returns the value of the Commands field in DbtTask_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *DbtTask_SdkV2) GetCommands(ctx context.Context) ([]types.String, bool) {
	if o.Commands.IsNull() || o.Commands.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Commands.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCommands sets the value of the Commands field in DbtTask_SdkV2.
func (o *DbtTask_SdkV2) SetCommands(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["commands"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Commands = types.ListValueMust(t, vs)
}

type DeleteJob_SdkV2 struct {
	// The canonical identifier of the job to delete. This field is required.
	JobId types.Int64 `tfsdk:"job_id" tf:""`
}

func (newState *DeleteJob_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteJob_SdkV2) {
}

func (newState *DeleteJob_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteJob_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteJob_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteJob_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteJob_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id": o.JobId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteJob_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.Int64Type,
		},
	}
}

type DeleteResponse_SdkV2 struct {
}

func (newState *DeleteResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse_SdkV2) {
}

func (newState *DeleteResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRun_SdkV2 struct {
	// ID of the run to delete.
	RunId types.Int64 `tfsdk:"run_id" tf:""`
}

func (newState *DeleteRun_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRun_SdkV2) {
}

func (newState *DeleteRun_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteRun_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRun_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

type DeleteRunResponse_SdkV2 struct {
}

func (newState *DeleteRunResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRunResponse_SdkV2) {
}

func (newState *DeleteRunResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteRunResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Represents a change to the job cluster's settings that would be required for
// the job clusters to become compliant with their policies.
type EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2 struct {
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

func (newState *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2) {
}

func (newState *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2) SyncEffectiveFieldsDuringRead(existingState EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnforcePolicyComplianceForJobResponseJobClusterSettingsChange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2
// only implements ToObjectValue() and Type().
func (o EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"field":          o.Field,
			"new_value":      o.NewValue,
			"previous_value": o.PreviousValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"field":          types.StringType,
			"new_value":      types.StringType,
			"previous_value": types.StringType,
		},
	}
}

type EnforcePolicyComplianceRequest_SdkV2 struct {
	// The ID of the job you want to enforce policy compliance on.
	JobId types.Int64 `tfsdk:"job_id" tf:""`
	// If set, previews changes made to the job to comply with its policy, but
	// does not update the job.
	ValidateOnly types.Bool `tfsdk:"validate_only" tf:"optional"`
}

func (newState *EnforcePolicyComplianceRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnforcePolicyComplianceRequest_SdkV2) {
}

func (newState *EnforcePolicyComplianceRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState EnforcePolicyComplianceRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnforcePolicyComplianceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnforcePolicyComplianceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnforcePolicyComplianceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o EnforcePolicyComplianceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id":        o.JobId,
			"validate_only": o.ValidateOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnforcePolicyComplianceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id":        types.Int64Type,
			"validate_only": types.BoolType,
		},
	}
}

type EnforcePolicyComplianceResponse_SdkV2 struct {
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

func (newState *EnforcePolicyComplianceResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnforcePolicyComplianceResponse_SdkV2) {
}

func (newState *EnforcePolicyComplianceResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState EnforcePolicyComplianceResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnforcePolicyComplianceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnforcePolicyComplianceResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"job_cluster_changes": reflect.TypeOf(EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2{}),
		"settings":            reflect.TypeOf(JobSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnforcePolicyComplianceResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o EnforcePolicyComplianceResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"has_changes":         o.HasChanges,
			"job_cluster_changes": o.JobClusterChanges,
			"settings":            o.Settings,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnforcePolicyComplianceResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"has_changes": types.BoolType,
			"job_cluster_changes": basetypes.ListType{
				ElemType: EnforcePolicyComplianceForJobResponseJobClusterSettingsChange{}.Type(ctx),
			},
			"settings": basetypes.ListType{
				ElemType: JobSettings{}.Type(ctx),
			},
		},
	}
}

// GetJobClusterChanges returns the value of the JobClusterChanges field in EnforcePolicyComplianceResponse_SdkV2 as
// a slice of EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EnforcePolicyComplianceResponse_SdkV2) GetJobClusterChanges(ctx context.Context) ([]EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2, bool) {
	if o.JobClusterChanges.IsNull() || o.JobClusterChanges.IsUnknown() {
		return nil, false
	}
	var v []EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2
	d := o.JobClusterChanges.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobClusterChanges sets the value of the JobClusterChanges field in EnforcePolicyComplianceResponse_SdkV2.
func (o *EnforcePolicyComplianceResponse_SdkV2) SetJobClusterChanges(ctx context.Context, v []EnforcePolicyComplianceForJobResponseJobClusterSettingsChange_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_cluster_changes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobClusterChanges = types.ListValueMust(t, vs)
}

// GetSettings returns the value of the Settings field in EnforcePolicyComplianceResponse_SdkV2 as
// a JobSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EnforcePolicyComplianceResponse_SdkV2) GetSettings(ctx context.Context) (JobSettings_SdkV2, bool) {
	var e JobSettings_SdkV2
	if o.Settings.IsNull() || o.Settings.IsUnknown() {
		return e, false
	}
	var v []JobSettings_SdkV2
	d := o.Settings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSettings sets the value of the Settings field in EnforcePolicyComplianceResponse_SdkV2.
func (o *EnforcePolicyComplianceResponse_SdkV2) SetSettings(ctx context.Context, v JobSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings"]
	o.Settings = types.ListValueMust(t, vs)
}

// Run was exported successfully.
type ExportRunOutput_SdkV2 struct {
	// The exported content in HTML format (one for every view item). To extract
	// the HTML notebook from the JSON response, download and run this [Python
	// script].
	//
	// [Python script]: https://docs.databricks.com/en/_static/examples/extract.py
	Views types.List `tfsdk:"views" tf:"optional"`
}

func (newState *ExportRunOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportRunOutput_SdkV2) {
}

func (newState *ExportRunOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExportRunOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportRunOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExportRunOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"views": reflect.TypeOf(ViewItem_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportRunOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o ExportRunOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"views": o.Views,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExportRunOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"views": basetypes.ListType{
				ElemType: ViewItem{}.Type(ctx),
			},
		},
	}
}

// GetViews returns the value of the Views field in ExportRunOutput_SdkV2 as
// a slice of ViewItem_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExportRunOutput_SdkV2) GetViews(ctx context.Context) ([]ViewItem_SdkV2, bool) {
	if o.Views.IsNull() || o.Views.IsUnknown() {
		return nil, false
	}
	var v []ViewItem_SdkV2
	d := o.Views.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetViews sets the value of the Views field in ExportRunOutput_SdkV2.
func (o *ExportRunOutput_SdkV2) SetViews(ctx context.Context, v []ViewItem_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["views"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Views = types.ListValueMust(t, vs)
}

// Export and retrieve a job run
type ExportRunRequest_SdkV2 struct {
	// The canonical identifier for the run. This field is required.
	RunId types.Int64 `tfsdk:"-"`
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	ViewsToExport types.String `tfsdk:"-"`
}

func (newState *ExportRunRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportRunRequest_SdkV2) {
}

func (newState *ExportRunRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExportRunRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportRunRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExportRunRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportRunRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ExportRunRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id":          o.RunId,
			"views_to_export": o.ViewsToExport,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExportRunRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id":          types.Int64Type,
			"views_to_export": types.StringType,
		},
	}
}

type FileArrivalTriggerConfiguration_SdkV2 struct {
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

func (newState *FileArrivalTriggerConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileArrivalTriggerConfiguration_SdkV2) {
}

func (newState *FileArrivalTriggerConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState FileArrivalTriggerConfiguration_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileArrivalTriggerConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FileArrivalTriggerConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileArrivalTriggerConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o FileArrivalTriggerConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"min_time_between_triggers_seconds": o.MinTimeBetweenTriggersSeconds,
			"url":                               o.Url,
			"wait_after_last_change_seconds":    o.WaitAfterLastChangeSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FileArrivalTriggerConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"min_time_between_triggers_seconds": types.Int64Type,
			"url":                               types.StringType,
			"wait_after_last_change_seconds":    types.Int64Type,
		},
	}
}

type ForEachStats_SdkV2 struct {
	// Sample of 3 most common error messages occurred during the iteration.
	ErrorMessageStats types.List `tfsdk:"error_message_stats" tf:"optional"`
	// Describes stats of the iteration. Only latest retries are considered.
	TaskRunStats types.List `tfsdk:"task_run_stats" tf:"optional,object"`
}

func (newState *ForEachStats_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForEachStats_SdkV2) {
}

func (newState *ForEachStats_SdkV2) SyncEffectiveFieldsDuringRead(existingState ForEachStats_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForEachStats.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ForEachStats_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error_message_stats": reflect.TypeOf(ForEachTaskErrorMessageStats_SdkV2{}),
		"task_run_stats":      reflect.TypeOf(ForEachTaskTaskRunStats_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForEachStats_SdkV2
// only implements ToObjectValue() and Type().
func (o ForEachStats_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_message_stats": o.ErrorMessageStats,
			"task_run_stats":      o.TaskRunStats,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ForEachStats_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_message_stats": basetypes.ListType{
				ElemType: ForEachTaskErrorMessageStats{}.Type(ctx),
			},
			"task_run_stats": basetypes.ListType{
				ElemType: ForEachTaskTaskRunStats{}.Type(ctx),
			},
		},
	}
}

// GetErrorMessageStats returns the value of the ErrorMessageStats field in ForEachStats_SdkV2 as
// a slice of ForEachTaskErrorMessageStats_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ForEachStats_SdkV2) GetErrorMessageStats(ctx context.Context) ([]ForEachTaskErrorMessageStats_SdkV2, bool) {
	if o.ErrorMessageStats.IsNull() || o.ErrorMessageStats.IsUnknown() {
		return nil, false
	}
	var v []ForEachTaskErrorMessageStats_SdkV2
	d := o.ErrorMessageStats.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetErrorMessageStats sets the value of the ErrorMessageStats field in ForEachStats_SdkV2.
func (o *ForEachStats_SdkV2) SetErrorMessageStats(ctx context.Context, v []ForEachTaskErrorMessageStats_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error_message_stats"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ErrorMessageStats = types.ListValueMust(t, vs)
}

// GetTaskRunStats returns the value of the TaskRunStats field in ForEachStats_SdkV2 as
// a ForEachTaskTaskRunStats_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ForEachStats_SdkV2) GetTaskRunStats(ctx context.Context) (ForEachTaskTaskRunStats_SdkV2, bool) {
	var e ForEachTaskTaskRunStats_SdkV2
	if o.TaskRunStats.IsNull() || o.TaskRunStats.IsUnknown() {
		return e, false
	}
	var v []ForEachTaskTaskRunStats_SdkV2
	d := o.TaskRunStats.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTaskRunStats sets the value of the TaskRunStats field in ForEachStats_SdkV2.
func (o *ForEachStats_SdkV2) SetTaskRunStats(ctx context.Context, v ForEachTaskTaskRunStats_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["task_run_stats"]
	o.TaskRunStats = types.ListValueMust(t, vs)
}

type ForEachTask_SdkV2 struct {
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

func (newState *ForEachTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForEachTask_SdkV2) {
}

func (newState *ForEachTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState ForEachTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForEachTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ForEachTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"task": reflect.TypeOf(Task_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForEachTask_SdkV2
// only implements ToObjectValue() and Type().
func (o ForEachTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"concurrency": o.Concurrency,
			"inputs":      o.Inputs,
			"task":        o.Task,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ForEachTask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"concurrency": types.Int64Type,
			"inputs":      types.StringType,
			"task": basetypes.ListType{
				ElemType: Task{}.Type(ctx),
			},
		},
	}
}

// GetTask returns the value of the Task field in ForEachTask_SdkV2 as
// a Task_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ForEachTask_SdkV2) GetTask(ctx context.Context) (Task_SdkV2, bool) {
	var e Task_SdkV2
	if o.Task.IsNull() || o.Task.IsUnknown() {
		return e, false
	}
	var v []Task_SdkV2
	d := o.Task.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTask sets the value of the Task field in ForEachTask_SdkV2.
func (o *ForEachTask_SdkV2) SetTask(ctx context.Context, v Task_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["task"]
	o.Task = types.ListValueMust(t, vs)
}

type ForEachTaskErrorMessageStats_SdkV2 struct {
	// Describes the count of such error message encountered during the
	// iterations.
	Count types.Int64 `tfsdk:"count" tf:"optional"`
	// Describes the error message occured during the iterations.
	ErrorMessage types.String `tfsdk:"error_message" tf:"optional"`
	// Describes the termination reason for the error message.
	TerminationCategory types.String `tfsdk:"termination_category" tf:"optional"`
}

func (newState *ForEachTaskErrorMessageStats_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForEachTaskErrorMessageStats_SdkV2) {
}

func (newState *ForEachTaskErrorMessageStats_SdkV2) SyncEffectiveFieldsDuringRead(existingState ForEachTaskErrorMessageStats_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForEachTaskErrorMessageStats.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ForEachTaskErrorMessageStats_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForEachTaskErrorMessageStats_SdkV2
// only implements ToObjectValue() and Type().
func (o ForEachTaskErrorMessageStats_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"count":                o.Count,
			"error_message":        o.ErrorMessage,
			"termination_category": o.TerminationCategory,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ForEachTaskErrorMessageStats_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"count":                types.Int64Type,
			"error_message":        types.StringType,
			"termination_category": types.StringType,
		},
	}
}

type ForEachTaskTaskRunStats_SdkV2 struct {
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

func (newState *ForEachTaskTaskRunStats_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForEachTaskTaskRunStats_SdkV2) {
}

func (newState *ForEachTaskTaskRunStats_SdkV2) SyncEffectiveFieldsDuringRead(existingState ForEachTaskTaskRunStats_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForEachTaskTaskRunStats.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ForEachTaskTaskRunStats_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForEachTaskTaskRunStats_SdkV2
// only implements ToObjectValue() and Type().
func (o ForEachTaskTaskRunStats_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"active_iterations":    o.ActiveIterations,
			"completed_iterations": o.CompletedIterations,
			"failed_iterations":    o.FailedIterations,
			"scheduled_iterations": o.ScheduledIterations,
			"succeeded_iterations": o.SucceededIterations,
			"total_iterations":     o.TotalIterations,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ForEachTaskTaskRunStats_SdkV2) Type(ctx context.Context) attr.Type {
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
type GetJobPermissionLevelsRequest_SdkV2 struct {
	// The job for which to get or manage permissions.
	JobId types.String `tfsdk:"-"`
}

func (newState *GetJobPermissionLevelsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetJobPermissionLevelsRequest_SdkV2) {
}

func (newState *GetJobPermissionLevelsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetJobPermissionLevelsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetJobPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetJobPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetJobPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetJobPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id": o.JobId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetJobPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.StringType,
		},
	}
}

type GetJobPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetJobPermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetJobPermissionLevelsResponse_SdkV2) {
}

func (newState *GetJobPermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetJobPermissionLevelsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetJobPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetJobPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(JobPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetJobPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetJobPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetJobPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: JobPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetJobPermissionLevelsResponse_SdkV2 as
// a slice of JobPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetJobPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]JobPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []JobPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetJobPermissionLevelsResponse_SdkV2.
func (o *GetJobPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []JobPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

// Get job permissions
type GetJobPermissionsRequest_SdkV2 struct {
	// The job for which to get or manage permissions.
	JobId types.String `tfsdk:"-"`
}

func (newState *GetJobPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetJobPermissionsRequest_SdkV2) {
}

func (newState *GetJobPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetJobPermissionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetJobPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetJobPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetJobPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetJobPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id": o.JobId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetJobPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.StringType,
		},
	}
}

// Get a single job
type GetJobRequest_SdkV2 struct {
	// The canonical identifier of the job to retrieve information about. This
	// field is required.
	JobId types.Int64 `tfsdk:"-"`
}

func (newState *GetJobRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetJobRequest_SdkV2) {
}

func (newState *GetJobRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetJobRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetJobRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetJobRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetJobRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetJobRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id": o.JobId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetJobRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.Int64Type,
		},
	}
}

// Get job policy compliance
type GetPolicyComplianceRequest_SdkV2 struct {
	// The ID of the job whose compliance status you are requesting.
	JobId types.Int64 `tfsdk:"-"`
}

func (newState *GetPolicyComplianceRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPolicyComplianceRequest_SdkV2) {
}

func (newState *GetPolicyComplianceRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPolicyComplianceRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPolicyComplianceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPolicyComplianceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPolicyComplianceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPolicyComplianceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id": o.JobId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPolicyComplianceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.Int64Type,
		},
	}
}

type GetPolicyComplianceResponse_SdkV2 struct {
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

func (newState *GetPolicyComplianceResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPolicyComplianceResponse_SdkV2) {
}

func (newState *GetPolicyComplianceResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetPolicyComplianceResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPolicyComplianceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPolicyComplianceResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"violations": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPolicyComplianceResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPolicyComplianceResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_compliant": o.IsCompliant,
			"violations":   o.Violations,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPolicyComplianceResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_compliant": types.BoolType,
			"violations": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetViolations returns the value of the Violations field in GetPolicyComplianceResponse_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPolicyComplianceResponse_SdkV2) GetViolations(ctx context.Context) (map[string]types.String, bool) {
	if o.Violations.IsNull() || o.Violations.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Violations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetViolations sets the value of the Violations field in GetPolicyComplianceResponse_SdkV2.
func (o *GetPolicyComplianceResponse_SdkV2) SetViolations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["violations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Violations = types.MapValueMust(t, vs)
}

// Get the output for a single run
type GetRunOutputRequest_SdkV2 struct {
	// The canonical identifier for the run.
	RunId types.Int64 `tfsdk:"-"`
}

func (newState *GetRunOutputRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRunOutputRequest_SdkV2) {
}

func (newState *GetRunOutputRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetRunOutputRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRunOutputRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRunOutputRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRunOutputRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRunOutputRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRunOutputRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

// Get a single job run
type GetRunRequest_SdkV2 struct {
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

func (newState *GetRunRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRunRequest_SdkV2) {
}

func (newState *GetRunRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetRunRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRunRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRunRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRunRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRunRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_history":         o.IncludeHistory,
			"include_resolved_values": o.IncludeResolvedValues,
			"page_token":              o.PageToken,
			"run_id":                  o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRunRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
type GitSnapshot_SdkV2 struct {
	// Commit that was used to execute the run. If git_branch was specified,
	// this points to the HEAD of the branch at the time of the run; if git_tag
	// was specified, this points to the commit the tag points to.
	UsedCommit types.String `tfsdk:"used_commit" tf:"optional"`
}

func (newState *GitSnapshot_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GitSnapshot_SdkV2) {
}

func (newState *GitSnapshot_SdkV2) SyncEffectiveFieldsDuringRead(existingState GitSnapshot_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GitSnapshot.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GitSnapshot_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GitSnapshot_SdkV2
// only implements ToObjectValue() and Type().
func (o GitSnapshot_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"used_commit": o.UsedCommit,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GitSnapshot_SdkV2) Type(ctx context.Context) attr.Type {
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
type GitSource_SdkV2 struct {
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

func (newState *GitSource_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GitSource_SdkV2) {
}

func (newState *GitSource_SdkV2) SyncEffectiveFieldsDuringRead(existingState GitSource_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GitSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GitSource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"git_snapshot": reflect.TypeOf(GitSnapshot_SdkV2{}),
		"job_source":   reflect.TypeOf(JobSource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GitSource_SdkV2
// only implements ToObjectValue() and Type().
func (o GitSource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":       o.GitBranch,
			"commit":       o.GitCommit,
			"git_provider": o.GitProvider,
			"git_snapshot": o.GitSnapshot,
			"tag":          o.GitTag,
			"url":          o.GitUrl,
			"job_source":   o.JobSource,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GitSource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":       types.StringType,
			"commit":       types.StringType,
			"git_provider": types.StringType,
			"git_snapshot": basetypes.ListType{
				ElemType: GitSnapshot{}.Type(ctx),
			},
			"tag": types.StringType,
			"url": types.StringType,
			"job_source": basetypes.ListType{
				ElemType: JobSource{}.Type(ctx),
			},
		},
	}
}

// GetGitSnapshot returns the value of the GitSnapshot field in GitSource_SdkV2 as
// a GitSnapshot_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GitSource_SdkV2) GetGitSnapshot(ctx context.Context) (GitSnapshot_SdkV2, bool) {
	var e GitSnapshot_SdkV2
	if o.GitSnapshot.IsNull() || o.GitSnapshot.IsUnknown() {
		return e, false
	}
	var v []GitSnapshot_SdkV2
	d := o.GitSnapshot.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitSnapshot sets the value of the GitSnapshot field in GitSource_SdkV2.
func (o *GitSource_SdkV2) SetGitSnapshot(ctx context.Context, v GitSnapshot_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["git_snapshot"]
	o.GitSnapshot = types.ListValueMust(t, vs)
}

// GetJobSource returns the value of the JobSource field in GitSource_SdkV2 as
// a JobSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GitSource_SdkV2) GetJobSource(ctx context.Context) (JobSource_SdkV2, bool) {
	var e JobSource_SdkV2
	if o.JobSource.IsNull() || o.JobSource.IsUnknown() {
		return e, false
	}
	var v []JobSource_SdkV2
	d := o.JobSource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSource sets the value of the JobSource field in GitSource_SdkV2.
func (o *GitSource_SdkV2) SetJobSource(ctx context.Context, v JobSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_source"]
	o.JobSource = types.ListValueMust(t, vs)
}

// Job was retrieved successfully.
type Job_SdkV2 struct {
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
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id" tf:"computed"`
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

func (newState *Job_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Job_SdkV2) {
}

func (newState *Job_SdkV2) SyncEffectiveFieldsDuringRead(existingState Job_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Job.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Job_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings": reflect.TypeOf(JobSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Job_SdkV2
// only implements ToObjectValue() and Type().
func (o Job_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_time":               o.CreatedTime,
			"creator_user_name":          o.CreatorUserName,
			"effective_budget_policy_id": o.EffectiveBudgetPolicyId,
			"job_id":                     o.JobId,
			"run_as_user_name":           o.RunAsUserName,
			"settings":                   o.Settings,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Job_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_time":               types.Int64Type,
			"creator_user_name":          types.StringType,
			"effective_budget_policy_id": types.StringType,
			"job_id":                     types.Int64Type,
			"run_as_user_name":           types.StringType,
			"settings": basetypes.ListType{
				ElemType: JobSettings{}.Type(ctx),
			},
		},
	}
}

// GetSettings returns the value of the Settings field in Job_SdkV2 as
// a JobSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Job_SdkV2) GetSettings(ctx context.Context) (JobSettings_SdkV2, bool) {
	var e JobSettings_SdkV2
	if o.Settings.IsNull() || o.Settings.IsUnknown() {
		return e, false
	}
	var v []JobSettings_SdkV2
	d := o.Settings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSettings sets the value of the Settings field in Job_SdkV2.
func (o *Job_SdkV2) SetSettings(ctx context.Context, v JobSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings"]
	o.Settings = types.ListValueMust(t, vs)
}

type JobAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *JobAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobAccessControlRequest_SdkV2) {
}

func (newState *JobAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobAccessControlRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o JobAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type JobAccessControlResponse_SdkV2 struct {
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

func (newState *JobAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobAccessControlResponse_SdkV2) {
}

func (newState *JobAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobAccessControlResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(JobPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o JobAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: JobPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in JobAccessControlResponse_SdkV2 as
// a slice of JobPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]JobPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []JobPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in JobAccessControlResponse_SdkV2.
func (o *JobAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []JobPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type JobCluster_SdkV2 struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	JobClusterKey types.String `tfsdk:"job_cluster_key" tf:""`
	// If new_cluster, a description of a cluster that is created for each task.
	NewCluster types.List `tfsdk:"new_cluster" tf:"object"`
}

func (newState *JobCluster_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobCluster_SdkV2) {
}

func (newState *JobCluster_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobCluster_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobCluster.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobCluster_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"new_cluster": reflect.TypeOf(compute_tf.ClusterSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobCluster_SdkV2
// only implements ToObjectValue() and Type().
func (o JobCluster_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_cluster_key": o.JobClusterKey,
			"new_cluster":     o.NewCluster,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobCluster_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_cluster_key": types.StringType,
			"new_cluster": basetypes.ListType{
				ElemType: compute_tf.ClusterSpec{}.Type(ctx),
			},
		},
	}
}

// GetNewCluster returns the value of the NewCluster field in JobCluster_SdkV2 as
// a compute_tf.ClusterSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobCluster_SdkV2) GetNewCluster(ctx context.Context) (compute_tf.ClusterSpec_SdkV2, bool) {
	var e compute_tf.ClusterSpec_SdkV2
	if o.NewCluster.IsNull() || o.NewCluster.IsUnknown() {
		return e, false
	}
	var v []compute_tf.ClusterSpec_SdkV2
	d := o.NewCluster.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNewCluster sets the value of the NewCluster field in JobCluster_SdkV2.
func (o *JobCluster_SdkV2) SetNewCluster(ctx context.Context, v compute_tf.ClusterSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["new_cluster"]
	o.NewCluster = types.ListValueMust(t, vs)
}

type JobCompliance_SdkV2 struct {
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

func (newState *JobCompliance_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobCompliance_SdkV2) {
}

func (newState *JobCompliance_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobCompliance_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobCompliance.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobCompliance_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"violations": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobCompliance_SdkV2
// only implements ToObjectValue() and Type().
func (o JobCompliance_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_compliant": o.IsCompliant,
			"job_id":       o.JobId,
			"violations":   o.Violations,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobCompliance_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetViolations returns the value of the Violations field in JobCompliance_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobCompliance_SdkV2) GetViolations(ctx context.Context) (map[string]types.String, bool) {
	if o.Violations.IsNull() || o.Violations.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Violations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetViolations sets the value of the Violations field in JobCompliance_SdkV2.
func (o *JobCompliance_SdkV2) SetViolations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["violations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Violations = types.MapValueMust(t, vs)
}

type JobDeployment_SdkV2 struct {
	// The kind of deployment that manages the job.
	//
	// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
	Kind types.String `tfsdk:"kind" tf:""`
	// Path of the file that contains deployment metadata.
	MetadataFilePath types.String `tfsdk:"metadata_file_path" tf:"optional"`
}

func (newState *JobDeployment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobDeployment_SdkV2) {
}

func (newState *JobDeployment_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobDeployment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobDeployment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobDeployment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobDeployment_SdkV2
// only implements ToObjectValue() and Type().
func (o JobDeployment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kind":               o.Kind,
			"metadata_file_path": o.MetadataFilePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobDeployment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"kind":               types.StringType,
			"metadata_file_path": types.StringType,
		},
	}
}

type JobEmailNotifications_SdkV2 struct {
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

func (newState *JobEmailNotifications_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobEmailNotifications_SdkV2) {
}

func (newState *JobEmailNotifications_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobEmailNotifications_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobEmailNotifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobEmailNotifications_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_duration_warning_threshold_exceeded": reflect.TypeOf(types.String{}),
		"on_failure":                             reflect.TypeOf(types.String{}),
		"on_start":                               reflect.TypeOf(types.String{}),
		"on_streaming_backlog_exceeded":          reflect.TypeOf(types.String{}),
		"on_success":                             reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobEmailNotifications_SdkV2
// only implements ToObjectValue() and Type().
func (o JobEmailNotifications_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"no_alert_for_skipped_runs":              o.NoAlertForSkippedRuns,
			"on_duration_warning_threshold_exceeded": o.OnDurationWarningThresholdExceeded,
			"on_failure":                             o.OnFailure,
			"on_start":                               o.OnStart,
			"on_streaming_backlog_exceeded":          o.OnStreamingBacklogExceeded,
			"on_success":                             o.OnSuccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobEmailNotifications_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetOnDurationWarningThresholdExceeded returns the value of the OnDurationWarningThresholdExceeded field in JobEmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobEmailNotifications_SdkV2) GetOnDurationWarningThresholdExceeded(ctx context.Context) ([]types.String, bool) {
	if o.OnDurationWarningThresholdExceeded.IsNull() || o.OnDurationWarningThresholdExceeded.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnDurationWarningThresholdExceeded.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnDurationWarningThresholdExceeded sets the value of the OnDurationWarningThresholdExceeded field in JobEmailNotifications_SdkV2.
func (o *JobEmailNotifications_SdkV2) SetOnDurationWarningThresholdExceeded(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_duration_warning_threshold_exceeded"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnDurationWarningThresholdExceeded = types.ListValueMust(t, vs)
}

// GetOnFailure returns the value of the OnFailure field in JobEmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobEmailNotifications_SdkV2) GetOnFailure(ctx context.Context) ([]types.String, bool) {
	if o.OnFailure.IsNull() || o.OnFailure.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnFailure.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnFailure sets the value of the OnFailure field in JobEmailNotifications_SdkV2.
func (o *JobEmailNotifications_SdkV2) SetOnFailure(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_failure"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnFailure = types.ListValueMust(t, vs)
}

// GetOnStart returns the value of the OnStart field in JobEmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobEmailNotifications_SdkV2) GetOnStart(ctx context.Context) ([]types.String, bool) {
	if o.OnStart.IsNull() || o.OnStart.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnStart.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnStart sets the value of the OnStart field in JobEmailNotifications_SdkV2.
func (o *JobEmailNotifications_SdkV2) SetOnStart(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_start"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnStart = types.ListValueMust(t, vs)
}

// GetOnStreamingBacklogExceeded returns the value of the OnStreamingBacklogExceeded field in JobEmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobEmailNotifications_SdkV2) GetOnStreamingBacklogExceeded(ctx context.Context) ([]types.String, bool) {
	if o.OnStreamingBacklogExceeded.IsNull() || o.OnStreamingBacklogExceeded.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnStreamingBacklogExceeded.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnStreamingBacklogExceeded sets the value of the OnStreamingBacklogExceeded field in JobEmailNotifications_SdkV2.
func (o *JobEmailNotifications_SdkV2) SetOnStreamingBacklogExceeded(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_streaming_backlog_exceeded"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnStreamingBacklogExceeded = types.ListValueMust(t, vs)
}

// GetOnSuccess returns the value of the OnSuccess field in JobEmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobEmailNotifications_SdkV2) GetOnSuccess(ctx context.Context) ([]types.String, bool) {
	if o.OnSuccess.IsNull() || o.OnSuccess.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnSuccess.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnSuccess sets the value of the OnSuccess field in JobEmailNotifications_SdkV2.
func (o *JobEmailNotifications_SdkV2) SetOnSuccess(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_success"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnSuccess = types.ListValueMust(t, vs)
}

type JobEnvironment_SdkV2 struct {
	// The key of an environment. It has to be unique within a job.
	EnvironmentKey types.String `tfsdk:"environment_key" tf:""`
	// The environment entity used to preserve serverless environment side panel
	// and jobs' environment for non-notebook task. In this minimal environment
	// spec, only pip dependencies are supported.
	Spec types.List `tfsdk:"spec" tf:"optional,object"`
}

func (newState *JobEnvironment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobEnvironment_SdkV2) {
}

func (newState *JobEnvironment_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobEnvironment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobEnvironment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobEnvironment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec": reflect.TypeOf(compute_tf.Environment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobEnvironment_SdkV2
// only implements ToObjectValue() and Type().
func (o JobEnvironment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"environment_key": o.EnvironmentKey,
			"spec":            o.Spec,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobEnvironment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"environment_key": types.StringType,
			"spec": basetypes.ListType{
				ElemType: compute_tf.Environment{}.Type(ctx),
			},
		},
	}
}

// GetSpec returns the value of the Spec field in JobEnvironment_SdkV2 as
// a compute_tf.Environment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobEnvironment_SdkV2) GetSpec(ctx context.Context) (compute_tf.Environment_SdkV2, bool) {
	var e compute_tf.Environment_SdkV2
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v []compute_tf.Environment_SdkV2
	d := o.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in JobEnvironment_SdkV2.
func (o *JobEnvironment_SdkV2) SetSpec(ctx context.Context, v compute_tf.Environment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	o.Spec = types.ListValueMust(t, vs)
}

type JobNotificationSettings_SdkV2 struct {
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns types.Bool `tfsdk:"no_alert_for_canceled_runs" tf:"optional"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns types.Bool `tfsdk:"no_alert_for_skipped_runs" tf:"optional"`
}

func (newState *JobNotificationSettings_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobNotificationSettings_SdkV2) {
}

func (newState *JobNotificationSettings_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobNotificationSettings_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobNotificationSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobNotificationSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobNotificationSettings_SdkV2
// only implements ToObjectValue() and Type().
func (o JobNotificationSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"no_alert_for_canceled_runs": o.NoAlertForCanceledRuns,
			"no_alert_for_skipped_runs":  o.NoAlertForSkippedRuns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobNotificationSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"no_alert_for_canceled_runs": types.BoolType,
			"no_alert_for_skipped_runs":  types.BoolType,
		},
	}
}

type JobParameter_SdkV2 struct {
	// The optional default value of the parameter
	Default types.String `tfsdk:"default" tf:"optional"`
	// The name of the parameter
	Name types.String `tfsdk:"name" tf:"optional"`
	// The value used in the run
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *JobParameter_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobParameter_SdkV2) {
}

func (newState *JobParameter_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobParameter_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobParameter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobParameter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobParameter_SdkV2
// only implements ToObjectValue() and Type().
func (o JobParameter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default": o.Default,
			"name":    o.Name,
			"value":   o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobParameter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default": types.StringType,
			"name":    types.StringType,
			"value":   types.StringType,
		},
	}
}

type JobParameterDefinition_SdkV2 struct {
	// Default value of the parameter.
	Default types.String `tfsdk:"default" tf:""`
	// The name of the defined parameter. May only contain alphanumeric
	// characters, `_`, `-`, and `.`
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *JobParameterDefinition_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobParameterDefinition_SdkV2) {
}

func (newState *JobParameterDefinition_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobParameterDefinition_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobParameterDefinition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobParameterDefinition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobParameterDefinition_SdkV2
// only implements ToObjectValue() and Type().
func (o JobParameterDefinition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default": o.Default,
			"name":    o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobParameterDefinition_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default": types.StringType,
			"name":    types.StringType,
		},
	}
}

type JobPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *JobPermission_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobPermission_SdkV2) {
}

func (newState *JobPermission_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobPermission_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o JobPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in JobPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in JobPermission_SdkV2.
func (o *JobPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type JobPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *JobPermissions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobPermissions_SdkV2) {
}

func (newState *JobPermissions_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobPermissions_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(JobAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o JobPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: JobAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in JobPermissions_SdkV2 as
// a slice of JobAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]JobAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []JobAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in JobPermissions_SdkV2.
func (o *JobPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []JobAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type JobPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *JobPermissionsDescription_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobPermissionsDescription_SdkV2) {
}

func (newState *JobPermissionsDescription_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobPermissionsDescription_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o JobPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type JobPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The job for which to get or manage permissions.
	JobId types.String `tfsdk:"-"`
}

func (newState *JobPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobPermissionsRequest_SdkV2) {
}

func (newState *JobPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobPermissionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(JobAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o JobPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"job_id":              o.JobId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: JobAccessControlRequest{}.Type(ctx),
			},
			"job_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in JobPermissionsRequest_SdkV2 as
// a slice of JobAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]JobAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []JobAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in JobPermissionsRequest_SdkV2.
func (o *JobPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []JobAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

// Write-only setting. Specifies the user or service principal that the job runs
// as. If not specified, the job runs as the user who created the job.
//
// Either `user_name` or `service_principal_name` should be specified. If not,
// an error is thrown.
type JobRunAs_SdkV2 struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// The email of an active workspace user. Non-admin users can only set this
	// field to their own email.
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *JobRunAs_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobRunAs_SdkV2) {
}

func (newState *JobRunAs_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobRunAs_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobRunAs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobRunAs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobRunAs_SdkV2
// only implements ToObjectValue() and Type().
func (o JobRunAs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobRunAs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type JobSettings_SdkV2 struct {
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
	// Write-only setting. Specifies the user or service principal that the job
	// runs as. If not specified, the job runs as the user who created the job.
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

func (newState *JobSettings_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobSettings_SdkV2) {
}

func (newState *JobSettings_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobSettings_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"continuous":            reflect.TypeOf(Continuous_SdkV2{}),
		"deployment":            reflect.TypeOf(JobDeployment_SdkV2{}),
		"email_notifications":   reflect.TypeOf(JobEmailNotifications_SdkV2{}),
		"environment":           reflect.TypeOf(JobEnvironment_SdkV2{}),
		"git_source":            reflect.TypeOf(GitSource_SdkV2{}),
		"health":                reflect.TypeOf(JobsHealthRules_SdkV2{}),
		"job_cluster":           reflect.TypeOf(JobCluster_SdkV2{}),
		"notification_settings": reflect.TypeOf(JobNotificationSettings_SdkV2{}),
		"parameter":             reflect.TypeOf(JobParameterDefinition_SdkV2{}),
		"queue":                 reflect.TypeOf(QueueSettings_SdkV2{}),
		"run_as":                reflect.TypeOf(JobRunAs_SdkV2{}),
		"schedule":              reflect.TypeOf(CronSchedule_SdkV2{}),
		"tags":                  reflect.TypeOf(types.String{}),
		"task":                  reflect.TypeOf(Task_SdkV2{}),
		"trigger":               reflect.TypeOf(TriggerSettings_SdkV2{}),
		"webhook_notifications": reflect.TypeOf(WebhookNotifications_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobSettings_SdkV2
// only implements ToObjectValue() and Type().
func (o JobSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id":      o.BudgetPolicyId,
			"continuous":            o.Continuous,
			"deployment":            o.Deployment,
			"description":           o.Description,
			"edit_mode":             o.EditMode,
			"email_notifications":   o.EmailNotifications,
			"environment":           o.Environments,
			"format":                o.Format,
			"git_source":            o.GitSource,
			"health":                o.Health,
			"job_cluster":           o.JobClusters,
			"max_concurrent_runs":   o.MaxConcurrentRuns,
			"name":                  o.Name,
			"notification_settings": o.NotificationSettings,
			"parameter":             o.Parameters,
			"queue":                 o.Queue,
			"run_as":                o.RunAs,
			"schedule":              o.Schedule,
			"tags":                  o.Tags,
			"task":                  o.Tasks,
			"timeout_seconds":       o.TimeoutSeconds,
			"trigger":               o.Trigger,
			"webhook_notifications": o.WebhookNotifications,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"continuous": basetypes.ListType{
				ElemType: Continuous{}.Type(ctx),
			},
			"deployment": basetypes.ListType{
				ElemType: JobDeployment{}.Type(ctx),
			},
			"description": types.StringType,
			"edit_mode":   types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: JobEmailNotifications{}.Type(ctx),
			},
			"environment": basetypes.ListType{
				ElemType: JobEnvironment{}.Type(ctx),
			},
			"format": types.StringType,
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.Type(ctx),
			},
			"health": basetypes.ListType{
				ElemType: JobsHealthRules{}.Type(ctx),
			},
			"job_cluster": basetypes.ListType{
				ElemType: JobCluster{}.Type(ctx),
			},
			"max_concurrent_runs": types.Int64Type,
			"name":                types.StringType,
			"notification_settings": basetypes.ListType{
				ElemType: JobNotificationSettings{}.Type(ctx),
			},
			"parameter": basetypes.ListType{
				ElemType: JobParameterDefinition{}.Type(ctx),
			},
			"queue": basetypes.ListType{
				ElemType: QueueSettings{}.Type(ctx),
			},
			"run_as": basetypes.ListType{
				ElemType: JobRunAs{}.Type(ctx),
			},
			"schedule": basetypes.ListType{
				ElemType: CronSchedule{}.Type(ctx),
			},
			"tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"task": basetypes.ListType{
				ElemType: Task{}.Type(ctx),
			},
			"timeout_seconds": types.Int64Type,
			"trigger": basetypes.ListType{
				ElemType: TriggerSettings{}.Type(ctx),
			},
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.Type(ctx),
			},
		},
	}
}

// GetContinuous returns the value of the Continuous field in JobSettings_SdkV2 as
// a Continuous_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetContinuous(ctx context.Context) (Continuous_SdkV2, bool) {
	var e Continuous_SdkV2
	if o.Continuous.IsNull() || o.Continuous.IsUnknown() {
		return e, false
	}
	var v []Continuous_SdkV2
	d := o.Continuous.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetContinuous sets the value of the Continuous field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetContinuous(ctx context.Context, v Continuous_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["continuous"]
	o.Continuous = types.ListValueMust(t, vs)
}

// GetDeployment returns the value of the Deployment field in JobSettings_SdkV2 as
// a JobDeployment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetDeployment(ctx context.Context) (JobDeployment_SdkV2, bool) {
	var e JobDeployment_SdkV2
	if o.Deployment.IsNull() || o.Deployment.IsUnknown() {
		return e, false
	}
	var v []JobDeployment_SdkV2
	d := o.Deployment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeployment sets the value of the Deployment field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetDeployment(ctx context.Context, v JobDeployment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["deployment"]
	o.Deployment = types.ListValueMust(t, vs)
}

// GetEmailNotifications returns the value of the EmailNotifications field in JobSettings_SdkV2 as
// a JobEmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetEmailNotifications(ctx context.Context) (JobEmailNotifications_SdkV2, bool) {
	var e JobEmailNotifications_SdkV2
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []JobEmailNotifications_SdkV2
	d := o.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetEmailNotifications(ctx context.Context, v JobEmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	o.EmailNotifications = types.ListValueMust(t, vs)
}

// GetEnvironments returns the value of the Environments field in JobSettings_SdkV2 as
// a slice of JobEnvironment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetEnvironments(ctx context.Context) ([]JobEnvironment_SdkV2, bool) {
	if o.Environments.IsNull() || o.Environments.IsUnknown() {
		return nil, false
	}
	var v []JobEnvironment_SdkV2
	d := o.Environments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironments sets the value of the Environments field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetEnvironments(ctx context.Context, v []JobEnvironment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Environments = types.ListValueMust(t, vs)
}

// GetGitSource returns the value of the GitSource field in JobSettings_SdkV2 as
// a GitSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetGitSource(ctx context.Context) (GitSource_SdkV2, bool) {
	var e GitSource_SdkV2
	if o.GitSource.IsNull() || o.GitSource.IsUnknown() {
		return e, false
	}
	var v []GitSource_SdkV2
	d := o.GitSource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitSource sets the value of the GitSource field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetGitSource(ctx context.Context, v GitSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["git_source"]
	o.GitSource = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in JobSettings_SdkV2 as
// a JobsHealthRules_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetHealth(ctx context.Context) (JobsHealthRules_SdkV2, bool) {
	var e JobsHealthRules_SdkV2
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v []JobsHealthRules_SdkV2
	d := o.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetHealth(ctx context.Context, v JobsHealthRules_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	o.Health = types.ListValueMust(t, vs)
}

// GetJobClusters returns the value of the JobClusters field in JobSettings_SdkV2 as
// a slice of JobCluster_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetJobClusters(ctx context.Context) ([]JobCluster_SdkV2, bool) {
	if o.JobClusters.IsNull() || o.JobClusters.IsUnknown() {
		return nil, false
	}
	var v []JobCluster_SdkV2
	d := o.JobClusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobClusters sets the value of the JobClusters field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetJobClusters(ctx context.Context, v []JobCluster_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_cluster"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobClusters = types.ListValueMust(t, vs)
}

// GetNotificationSettings returns the value of the NotificationSettings field in JobSettings_SdkV2 as
// a JobNotificationSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetNotificationSettings(ctx context.Context) (JobNotificationSettings_SdkV2, bool) {
	var e JobNotificationSettings_SdkV2
	if o.NotificationSettings.IsNull() || o.NotificationSettings.IsUnknown() {
		return e, false
	}
	var v []JobNotificationSettings_SdkV2
	d := o.NotificationSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotificationSettings sets the value of the NotificationSettings field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetNotificationSettings(ctx context.Context, v JobNotificationSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notification_settings"]
	o.NotificationSettings = types.ListValueMust(t, vs)
}

// GetParameters returns the value of the Parameters field in JobSettings_SdkV2 as
// a slice of JobParameterDefinition_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetParameters(ctx context.Context) ([]JobParameterDefinition_SdkV2, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []JobParameterDefinition_SdkV2
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetParameters(ctx context.Context, v []JobParameterDefinition_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameter"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

// GetQueue returns the value of the Queue field in JobSettings_SdkV2 as
// a QueueSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetQueue(ctx context.Context) (QueueSettings_SdkV2, bool) {
	var e QueueSettings_SdkV2
	if o.Queue.IsNull() || o.Queue.IsUnknown() {
		return e, false
	}
	var v []QueueSettings_SdkV2
	d := o.Queue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueue sets the value of the Queue field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetQueue(ctx context.Context, v QueueSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["queue"]
	o.Queue = types.ListValueMust(t, vs)
}

// GetRunAs returns the value of the RunAs field in JobSettings_SdkV2 as
// a JobRunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetRunAs(ctx context.Context) (JobRunAs_SdkV2, bool) {
	var e JobRunAs_SdkV2
	if o.RunAs.IsNull() || o.RunAs.IsUnknown() {
		return e, false
	}
	var v []JobRunAs_SdkV2
	d := o.RunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunAs sets the value of the RunAs field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetRunAs(ctx context.Context, v JobRunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_as"]
	o.RunAs = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in JobSettings_SdkV2 as
// a CronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetSchedule(ctx context.Context) (CronSchedule_SdkV2, bool) {
	var e CronSchedule_SdkV2
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []CronSchedule_SdkV2
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetSchedule(ctx context.Context, v CronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in JobSettings_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetTags(ctx context.Context) (map[string]types.String, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.MapValueMust(t, vs)
}

// GetTasks returns the value of the Tasks field in JobSettings_SdkV2 as
// a slice of Task_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetTasks(ctx context.Context) ([]Task_SdkV2, bool) {
	if o.Tasks.IsNull() || o.Tasks.IsUnknown() {
		return nil, false
	}
	var v []Task_SdkV2
	d := o.Tasks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTasks sets the value of the Tasks field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetTasks(ctx context.Context, v []Task_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["task"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tasks = types.ListValueMust(t, vs)
}

// GetTrigger returns the value of the Trigger field in JobSettings_SdkV2 as
// a TriggerSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetTrigger(ctx context.Context) (TriggerSettings_SdkV2, bool) {
	var e TriggerSettings_SdkV2
	if o.Trigger.IsNull() || o.Trigger.IsUnknown() {
		return e, false
	}
	var v []TriggerSettings_SdkV2
	d := o.Trigger.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrigger sets the value of the Trigger field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetTrigger(ctx context.Context, v TriggerSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["trigger"]
	o.Trigger = types.ListValueMust(t, vs)
}

// GetWebhookNotifications returns the value of the WebhookNotifications field in JobSettings_SdkV2 as
// a WebhookNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *JobSettings_SdkV2) GetWebhookNotifications(ctx context.Context) (WebhookNotifications_SdkV2, bool) {
	var e WebhookNotifications_SdkV2
	if o.WebhookNotifications.IsNull() || o.WebhookNotifications.IsUnknown() {
		return e, false
	}
	var v []WebhookNotifications_SdkV2
	d := o.WebhookNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhookNotifications sets the value of the WebhookNotifications field in JobSettings_SdkV2.
func (o *JobSettings_SdkV2) SetWebhookNotifications(ctx context.Context, v WebhookNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook_notifications"]
	o.WebhookNotifications = types.ListValueMust(t, vs)
}

// The source of the job specification in the remote repository when the job is
// source controlled.
type JobSource_SdkV2 struct {
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

func (newState *JobSource_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobSource_SdkV2) {
}

func (newState *JobSource_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobSource_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobSource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobSource_SdkV2
// only implements ToObjectValue() and Type().
func (o JobSource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dirty_state":            o.DirtyState,
			"import_from_git_branch": o.ImportFromGitBranch,
			"job_config_path":        o.JobConfigPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobSource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dirty_state":            types.StringType,
			"import_from_git_branch": types.StringType,
			"job_config_path":        types.StringType,
		},
	}
}

type JobsHealthRule_SdkV2 struct {
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
	Metric types.String `tfsdk:"metric" tf:""`
	// Specifies the operator used to compare the health metric value with the
	// specified threshold.
	Op types.String `tfsdk:"op" tf:""`
	// Specifies the threshold value that the health metric should obey to
	// satisfy the health rule.
	Value types.Int64 `tfsdk:"value" tf:""`
}

func (newState *JobsHealthRule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobsHealthRule_SdkV2) {
}

func (newState *JobsHealthRule_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobsHealthRule_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobsHealthRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobsHealthRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobsHealthRule_SdkV2
// only implements ToObjectValue() and Type().
func (o JobsHealthRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metric": o.Metric,
			"op":     o.Op,
			"value":  o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobsHealthRule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metric": types.StringType,
			"op":     types.StringType,
			"value":  types.Int64Type,
		},
	}
}

// An optional set of health rules that can be defined for this job.
type JobsHealthRules_SdkV2 struct {
	Rules types.List `tfsdk:"rules" tf:"optional"`
}

func (newState *JobsHealthRules_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobsHealthRules_SdkV2) {
}

func (newState *JobsHealthRules_SdkV2) SyncEffectiveFieldsDuringRead(existingState JobsHealthRules_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobsHealthRules.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a JobsHealthRules_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rules": reflect.TypeOf(JobsHealthRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobsHealthRules_SdkV2
// only implements ToObjectValue() and Type().
func (o JobsHealthRules_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"rules": o.Rules,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobsHealthRules_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"rules": basetypes.ListType{
				ElemType: JobsHealthRule{}.Type(ctx),
			},
		},
	}
}

// GetRules returns the value of the Rules field in JobsHealthRules_SdkV2 as
// a slice of JobsHealthRule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *JobsHealthRules_SdkV2) GetRules(ctx context.Context) ([]JobsHealthRule_SdkV2, bool) {
	if o.Rules.IsNull() || o.Rules.IsUnknown() {
		return nil, false
	}
	var v []JobsHealthRule_SdkV2
	d := o.Rules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRules sets the value of the Rules field in JobsHealthRules_SdkV2.
func (o *JobsHealthRules_SdkV2) SetRules(ctx context.Context, v []JobsHealthRule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Rules = types.ListValueMust(t, vs)
}

type ListJobComplianceForPolicyResponse_SdkV2 struct {
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

func (newState *ListJobComplianceForPolicyResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListJobComplianceForPolicyResponse_SdkV2) {
}

func (newState *ListJobComplianceForPolicyResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListJobComplianceForPolicyResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListJobComplianceForPolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListJobComplianceForPolicyResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"jobs": reflect.TypeOf(JobCompliance_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListJobComplianceForPolicyResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListJobComplianceForPolicyResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"jobs":            o.Jobs,
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListJobComplianceForPolicyResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"jobs": basetypes.ListType{
				ElemType: JobCompliance{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
		},
	}
}

// GetJobs returns the value of the Jobs field in ListJobComplianceForPolicyResponse_SdkV2 as
// a slice of JobCompliance_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListJobComplianceForPolicyResponse_SdkV2) GetJobs(ctx context.Context) ([]JobCompliance_SdkV2, bool) {
	if o.Jobs.IsNull() || o.Jobs.IsUnknown() {
		return nil, false
	}
	var v []JobCompliance_SdkV2
	d := o.Jobs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobs sets the value of the Jobs field in ListJobComplianceForPolicyResponse_SdkV2.
func (o *ListJobComplianceForPolicyResponse_SdkV2) SetJobs(ctx context.Context, v []JobCompliance_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["jobs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Jobs = types.ListValueMust(t, vs)
}

// List job policy compliance
type ListJobComplianceRequest_SdkV2 struct {
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

func (newState *ListJobComplianceRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListJobComplianceRequest_SdkV2) {
}

func (newState *ListJobComplianceRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListJobComplianceRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListJobComplianceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListJobComplianceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListJobComplianceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListJobComplianceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
			"policy_id":  o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListJobComplianceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"policy_id":  types.StringType,
		},
	}
}

// List jobs
type ListJobsRequest_SdkV2 struct {
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

func (newState *ListJobsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListJobsRequest_SdkV2) {
}

func (newState *ListJobsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListJobsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListJobsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListJobsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListJobsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListJobsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"expand_tasks": o.ExpandTasks,
			"limit":        o.Limit,
			"name":         o.Name,
			"offset":       o.Offset,
			"page_token":   o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListJobsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
type ListJobsResponse_SdkV2 struct {
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

func (newState *ListJobsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListJobsResponse_SdkV2) {
}

func (newState *ListJobsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListJobsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListJobsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListJobsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"jobs": reflect.TypeOf(BaseJob_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListJobsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListJobsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"has_more":        o.HasMore,
			"jobs":            o.Jobs,
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListJobsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"has_more": types.BoolType,
			"jobs": basetypes.ListType{
				ElemType: BaseJob{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
		},
	}
}

// GetJobs returns the value of the Jobs field in ListJobsResponse_SdkV2 as
// a slice of BaseJob_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListJobsResponse_SdkV2) GetJobs(ctx context.Context) ([]BaseJob_SdkV2, bool) {
	if o.Jobs.IsNull() || o.Jobs.IsUnknown() {
		return nil, false
	}
	var v []BaseJob_SdkV2
	d := o.Jobs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobs sets the value of the Jobs field in ListJobsResponse_SdkV2.
func (o *ListJobsResponse_SdkV2) SetJobs(ctx context.Context, v []BaseJob_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["jobs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Jobs = types.ListValueMust(t, vs)
}

// List job runs
type ListRunsRequest_SdkV2 struct {
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

func (newState *ListRunsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRunsRequest_SdkV2) {
}

func (newState *ListRunsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListRunsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRunsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRunsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRunsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListRunsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"active_only":     o.ActiveOnly,
			"completed_only":  o.CompletedOnly,
			"expand_tasks":    o.ExpandTasks,
			"job_id":          o.JobId,
			"limit":           o.Limit,
			"offset":          o.Offset,
			"page_token":      o.PageToken,
			"run_type":        o.RunType,
			"start_time_from": o.StartTimeFrom,
			"start_time_to":   o.StartTimeTo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRunsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
type ListRunsResponse_SdkV2 struct {
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

func (newState *ListRunsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRunsResponse_SdkV2) {
}

func (newState *ListRunsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListRunsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRunsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRunsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(BaseRun_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRunsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListRunsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"has_more":        o.HasMore,
			"next_page_token": o.NextPageToken,
			"prev_page_token": o.PrevPageToken,
			"runs":            o.Runs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRunsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"has_more":        types.BoolType,
			"next_page_token": types.StringType,
			"prev_page_token": types.StringType,
			"runs": basetypes.ListType{
				ElemType: BaseRun{}.Type(ctx),
			},
		},
	}
}

// GetRuns returns the value of the Runs field in ListRunsResponse_SdkV2 as
// a slice of BaseRun_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListRunsResponse_SdkV2) GetRuns(ctx context.Context) ([]BaseRun_SdkV2, bool) {
	if o.Runs.IsNull() || o.Runs.IsUnknown() {
		return nil, false
	}
	var v []BaseRun_SdkV2
	d := o.Runs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRuns sets the value of the Runs field in ListRunsResponse_SdkV2.
func (o *ListRunsResponse_SdkV2) SetRuns(ctx context.Context, v []BaseRun_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["runs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Runs = types.ListValueMust(t, vs)
}

type NotebookOutput_SdkV2 struct {
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

func (newState *NotebookOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NotebookOutput_SdkV2) {
}

func (newState *NotebookOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState NotebookOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotebookOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NotebookOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o NotebookOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result":    o.Result,
			"truncated": o.Truncated,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NotebookOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"result":    types.StringType,
			"truncated": types.BoolType,
		},
	}
}

type NotebookTask_SdkV2 struct {
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

func (newState *NotebookTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NotebookTask_SdkV2) {
}

func (newState *NotebookTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState NotebookTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotebookTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NotebookTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"base_parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookTask_SdkV2
// only implements ToObjectValue() and Type().
func (o NotebookTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"base_parameters": o.BaseParameters,
			"notebook_path":   o.NotebookPath,
			"source":          o.Source,
			"warehouse_id":    o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NotebookTask_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetBaseParameters returns the value of the BaseParameters field in NotebookTask_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NotebookTask_SdkV2) GetBaseParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.BaseParameters.IsNull() || o.BaseParameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.BaseParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBaseParameters sets the value of the BaseParameters field in NotebookTask_SdkV2.
func (o *NotebookTask_SdkV2) SetBaseParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["base_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.BaseParameters = types.MapValueMust(t, vs)
}

type PeriodicTriggerConfiguration_SdkV2 struct {
	// The interval at which the trigger should run.
	Interval types.Int64 `tfsdk:"interval" tf:""`
	// The unit of time for the interval.
	Unit types.String `tfsdk:"unit" tf:""`
}

func (newState *PeriodicTriggerConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PeriodicTriggerConfiguration_SdkV2) {
}

func (newState *PeriodicTriggerConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState PeriodicTriggerConfiguration_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PeriodicTriggerConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PeriodicTriggerConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PeriodicTriggerConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o PeriodicTriggerConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"interval": o.Interval,
			"unit":     o.Unit,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PeriodicTriggerConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"interval": types.Int64Type,
			"unit":     types.StringType,
		},
	}
}

type PipelineParams_SdkV2 struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh types.Bool `tfsdk:"full_refresh" tf:"optional"`
}

func (newState *PipelineParams_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineParams_SdkV2) {
}

func (newState *PipelineParams_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineParams_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineParams.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineParams_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineParams_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineParams_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_refresh": o.FullRefresh,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineParams_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_refresh": types.BoolType,
		},
	}
}

type PipelineTask_SdkV2 struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh types.Bool `tfsdk:"full_refresh" tf:"optional"`
	// The full name of the pipeline task to execute.
	PipelineId types.String `tfsdk:"pipeline_id" tf:""`
}

func (newState *PipelineTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineTask_SdkV2) {
}

func (newState *PipelineTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineTask_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_refresh": o.FullRefresh,
			"pipeline_id":  o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineTask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_refresh": types.BoolType,
			"pipeline_id":  types.StringType,
		},
	}
}

type PythonWheelTask_SdkV2 struct {
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

func (newState *PythonWheelTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PythonWheelTask_SdkV2) {
}

func (newState *PythonWheelTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState PythonWheelTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PythonWheelTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PythonWheelTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"named_parameters": reflect.TypeOf(types.String{}),
		"parameters":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PythonWheelTask_SdkV2
// only implements ToObjectValue() and Type().
func (o PythonWheelTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entry_point":      o.EntryPoint,
			"named_parameters": o.NamedParameters,
			"package_name":     o.PackageName,
			"parameters":       o.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PythonWheelTask_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetNamedParameters returns the value of the NamedParameters field in PythonWheelTask_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PythonWheelTask_SdkV2) GetNamedParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.NamedParameters.IsNull() || o.NamedParameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.NamedParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNamedParameters sets the value of the NamedParameters field in PythonWheelTask_SdkV2.
func (o *PythonWheelTask_SdkV2) SetNamedParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["named_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NamedParameters = types.MapValueMust(t, vs)
}

// GetParameters returns the value of the Parameters field in PythonWheelTask_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PythonWheelTask_SdkV2) GetParameters(ctx context.Context) ([]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in PythonWheelTask_SdkV2.
func (o *PythonWheelTask_SdkV2) SetParameters(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

type QueueDetails_SdkV2 struct {
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

func (newState *QueueDetails_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueueDetails_SdkV2) {
}

func (newState *QueueDetails_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueueDetails_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueueDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueueDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueueDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o QueueDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"code":    o.Code,
			"message": o.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueueDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"code":    types.StringType,
			"message": types.StringType,
		},
	}
}

type QueueSettings_SdkV2 struct {
	// If true, enable queueing for the job. This is a required field.
	Enabled types.Bool `tfsdk:"enabled" tf:""`
}

func (newState *QueueSettings_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueueSettings_SdkV2) {
}

func (newState *QueueSettings_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueueSettings_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueueSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueueSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueueSettings_SdkV2
// only implements ToObjectValue() and Type().
func (o QueueSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled": o.Enabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueueSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

type RepairHistoryItem_SdkV2 struct {
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
	Type_ types.String `tfsdk:"type" tf:"optional"`
}

func (newState *RepairHistoryItem_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepairHistoryItem_SdkV2) {
}

func (newState *RepairHistoryItem_SdkV2) SyncEffectiveFieldsDuringRead(existingState RepairHistoryItem_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepairHistoryItem.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepairHistoryItem_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"state":        reflect.TypeOf(RunState_SdkV2{}),
		"status":       reflect.TypeOf(RunStatus_SdkV2{}),
		"task_run_ids": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepairHistoryItem_SdkV2
// only implements ToObjectValue() and Type().
func (o RepairHistoryItem_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time":     o.EndTime,
			"id":           o.Id,
			"start_time":   o.StartTime,
			"state":        o.State,
			"status":       o.Status,
			"task_run_ids": o.TaskRunIds,
			"type":         o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepairHistoryItem_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time":   types.Int64Type,
			"id":         types.Int64Type,
			"start_time": types.Int64Type,
			"state": basetypes.ListType{
				ElemType: RunState{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: RunStatus{}.Type(ctx),
			},
			"task_run_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"type": types.StringType,
		},
	}
}

// GetState returns the value of the State field in RepairHistoryItem_SdkV2 as
// a RunState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairHistoryItem_SdkV2) GetState(ctx context.Context) (RunState_SdkV2, bool) {
	var e RunState_SdkV2
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v []RunState_SdkV2
	d := o.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in RepairHistoryItem_SdkV2.
func (o *RepairHistoryItem_SdkV2) SetState(ctx context.Context, v RunState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	o.State = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in RepairHistoryItem_SdkV2 as
// a RunStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairHistoryItem_SdkV2) GetStatus(ctx context.Context) (RunStatus_SdkV2, bool) {
	var e RunStatus_SdkV2
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []RunStatus_SdkV2
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in RepairHistoryItem_SdkV2.
func (o *RepairHistoryItem_SdkV2) SetStatus(ctx context.Context, v RunStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
}

// GetTaskRunIds returns the value of the TaskRunIds field in RepairHistoryItem_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairHistoryItem_SdkV2) GetTaskRunIds(ctx context.Context) ([]types.Int64, bool) {
	if o.TaskRunIds.IsNull() || o.TaskRunIds.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.TaskRunIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTaskRunIds sets the value of the TaskRunIds field in RepairHistoryItem_SdkV2.
func (o *RepairHistoryItem_SdkV2) SetTaskRunIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["task_run_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TaskRunIds = types.ListValueMust(t, vs)
}

type RepairRun_SdkV2 struct {
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

func (newState *RepairRun_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepairRun_SdkV2) {
}

func (newState *RepairRun_SdkV2) SyncEffectiveFieldsDuringRead(existingState RepairRun_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepairRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepairRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbt_commands":        reflect.TypeOf(types.String{}),
		"jar_params":          reflect.TypeOf(types.String{}),
		"job_parameters":      reflect.TypeOf(types.String{}),
		"notebook_params":     reflect.TypeOf(types.String{}),
		"pipeline_params":     reflect.TypeOf(PipelineParams_SdkV2{}),
		"python_named_params": reflect.TypeOf(types.String{}),
		"python_params":       reflect.TypeOf(types.String{}),
		"rerun_tasks":         reflect.TypeOf(types.String{}),
		"spark_submit_params": reflect.TypeOf(types.String{}),
		"sql_params":          reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepairRun_SdkV2
// only implements ToObjectValue() and Type().
func (o RepairRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbt_commands":           o.DbtCommands,
			"jar_params":             o.JarParams,
			"job_parameters":         o.JobParameters,
			"latest_repair_id":       o.LatestRepairId,
			"notebook_params":        o.NotebookParams,
			"pipeline_params":        o.PipelineParams,
			"python_named_params":    o.PythonNamedParams,
			"python_params":          o.PythonParams,
			"rerun_all_failed_tasks": o.RerunAllFailedTasks,
			"rerun_dependent_tasks":  o.RerunDependentTasks,
			"rerun_tasks":            o.RerunTasks,
			"run_id":                 o.RunId,
			"spark_submit_params":    o.SparkSubmitParams,
			"sql_params":             o.SqlParams,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepairRun_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: PipelineParams{}.Type(ctx),
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

// GetDbtCommands returns the value of the DbtCommands field in RepairRun_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairRun_SdkV2) GetDbtCommands(ctx context.Context) ([]types.String, bool) {
	if o.DbtCommands.IsNull() || o.DbtCommands.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DbtCommands.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDbtCommands sets the value of the DbtCommands field in RepairRun_SdkV2.
func (o *RepairRun_SdkV2) SetDbtCommands(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbt_commands"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DbtCommands = types.ListValueMust(t, vs)
}

// GetJarParams returns the value of the JarParams field in RepairRun_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairRun_SdkV2) GetJarParams(ctx context.Context) ([]types.String, bool) {
	if o.JarParams.IsNull() || o.JarParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.JarParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJarParams sets the value of the JarParams field in RepairRun_SdkV2.
func (o *RepairRun_SdkV2) SetJarParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["jar_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JarParams = types.ListValueMust(t, vs)
}

// GetJobParameters returns the value of the JobParameters field in RepairRun_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairRun_SdkV2) GetJobParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.JobParameters.IsNull() || o.JobParameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.JobParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobParameters sets the value of the JobParameters field in RepairRun_SdkV2.
func (o *RepairRun_SdkV2) SetJobParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobParameters = types.MapValueMust(t, vs)
}

// GetNotebookParams returns the value of the NotebookParams field in RepairRun_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairRun_SdkV2) GetNotebookParams(ctx context.Context) (map[string]types.String, bool) {
	if o.NotebookParams.IsNull() || o.NotebookParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.NotebookParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookParams sets the value of the NotebookParams field in RepairRun_SdkV2.
func (o *RepairRun_SdkV2) SetNotebookParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NotebookParams = types.MapValueMust(t, vs)
}

// GetPipelineParams returns the value of the PipelineParams field in RepairRun_SdkV2 as
// a PipelineParams_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairRun_SdkV2) GetPipelineParams(ctx context.Context) (PipelineParams_SdkV2, bool) {
	var e PipelineParams_SdkV2
	if o.PipelineParams.IsNull() || o.PipelineParams.IsUnknown() {
		return e, false
	}
	var v []PipelineParams_SdkV2
	d := o.PipelineParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPipelineParams sets the value of the PipelineParams field in RepairRun_SdkV2.
func (o *RepairRun_SdkV2) SetPipelineParams(ctx context.Context, v PipelineParams_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pipeline_params"]
	o.PipelineParams = types.ListValueMust(t, vs)
}

// GetPythonNamedParams returns the value of the PythonNamedParams field in RepairRun_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairRun_SdkV2) GetPythonNamedParams(ctx context.Context) (map[string]types.String, bool) {
	if o.PythonNamedParams.IsNull() || o.PythonNamedParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.PythonNamedParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPythonNamedParams sets the value of the PythonNamedParams field in RepairRun_SdkV2.
func (o *RepairRun_SdkV2) SetPythonNamedParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_named_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PythonNamedParams = types.MapValueMust(t, vs)
}

// GetPythonParams returns the value of the PythonParams field in RepairRun_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairRun_SdkV2) GetPythonParams(ctx context.Context) ([]types.String, bool) {
	if o.PythonParams.IsNull() || o.PythonParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PythonParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPythonParams sets the value of the PythonParams field in RepairRun_SdkV2.
func (o *RepairRun_SdkV2) SetPythonParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PythonParams = types.ListValueMust(t, vs)
}

// GetRerunTasks returns the value of the RerunTasks field in RepairRun_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairRun_SdkV2) GetRerunTasks(ctx context.Context) ([]types.String, bool) {
	if o.RerunTasks.IsNull() || o.RerunTasks.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.RerunTasks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRerunTasks sets the value of the RerunTasks field in RepairRun_SdkV2.
func (o *RepairRun_SdkV2) SetRerunTasks(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rerun_tasks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RerunTasks = types.ListValueMust(t, vs)
}

// GetSparkSubmitParams returns the value of the SparkSubmitParams field in RepairRun_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairRun_SdkV2) GetSparkSubmitParams(ctx context.Context) ([]types.String, bool) {
	if o.SparkSubmitParams.IsNull() || o.SparkSubmitParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SparkSubmitParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkSubmitParams sets the value of the SparkSubmitParams field in RepairRun_SdkV2.
func (o *RepairRun_SdkV2) SetSparkSubmitParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_submit_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkSubmitParams = types.ListValueMust(t, vs)
}

// GetSqlParams returns the value of the SqlParams field in RepairRun_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepairRun_SdkV2) GetSqlParams(ctx context.Context) (map[string]types.String, bool) {
	if o.SqlParams.IsNull() || o.SqlParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SqlParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlParams sets the value of the SqlParams field in RepairRun_SdkV2.
func (o *RepairRun_SdkV2) SetSqlParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SqlParams = types.MapValueMust(t, vs)
}

// Run repair was initiated.
type RepairRunResponse_SdkV2 struct {
	// The ID of the repair. Must be provided in subsequent repairs using the
	// `latest_repair_id` field to ensure sequential repairs.
	RepairId types.Int64 `tfsdk:"repair_id" tf:"optional"`
}

func (newState *RepairRunResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepairRunResponse_SdkV2) {
}

func (newState *RepairRunResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState RepairRunResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepairRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepairRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepairRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RepairRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repair_id": o.RepairId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepairRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repair_id": types.Int64Type,
		},
	}
}

type ResetJob_SdkV2 struct {
	// The canonical identifier of the job to reset. This field is required.
	JobId types.Int64 `tfsdk:"job_id" tf:""`
	// The new settings of the job. These settings completely replace the old
	// settings.
	//
	// Changes to the field `JobBaseSettings.timeout_seconds` are applied to
	// active runs. Changes to other fields are applied to future runs only.
	NewSettings types.List `tfsdk:"new_settings" tf:"object"`
}

func (newState *ResetJob_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResetJob_SdkV2) {
}

func (newState *ResetJob_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResetJob_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResetJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResetJob_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"new_settings": reflect.TypeOf(JobSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResetJob_SdkV2
// only implements ToObjectValue() and Type().
func (o ResetJob_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id":       o.JobId,
			"new_settings": o.NewSettings,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResetJob_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id": types.Int64Type,
			"new_settings": basetypes.ListType{
				ElemType: JobSettings{}.Type(ctx),
			},
		},
	}
}

// GetNewSettings returns the value of the NewSettings field in ResetJob_SdkV2 as
// a JobSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResetJob_SdkV2) GetNewSettings(ctx context.Context) (JobSettings_SdkV2, bool) {
	var e JobSettings_SdkV2
	if o.NewSettings.IsNull() || o.NewSettings.IsUnknown() {
		return e, false
	}
	var v []JobSettings_SdkV2
	d := o.NewSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNewSettings sets the value of the NewSettings field in ResetJob_SdkV2.
func (o *ResetJob_SdkV2) SetNewSettings(ctx context.Context, v JobSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["new_settings"]
	o.NewSettings = types.ListValueMust(t, vs)
}

type ResetResponse_SdkV2 struct {
}

func (newState *ResetResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResetResponse_SdkV2) {
}

func (newState *ResetResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResetResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResetResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResetResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ResetResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ResetResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ResolvedConditionTaskValues_SdkV2 struct {
	Left types.String `tfsdk:"left" tf:"optional"`

	Right types.String `tfsdk:"right" tf:"optional"`
}

func (newState *ResolvedConditionTaskValues_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedConditionTaskValues_SdkV2) {
}

func (newState *ResolvedConditionTaskValues_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResolvedConditionTaskValues_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedConditionTaskValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedConditionTaskValues_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolvedConditionTaskValues_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolvedConditionTaskValues_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"left":  o.Left,
			"right": o.Right,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolvedConditionTaskValues_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"left":  types.StringType,
			"right": types.StringType,
		},
	}
}

type ResolvedDbtTaskValues_SdkV2 struct {
	Commands types.List `tfsdk:"commands" tf:"optional"`
}

func (newState *ResolvedDbtTaskValues_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedDbtTaskValues_SdkV2) {
}

func (newState *ResolvedDbtTaskValues_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResolvedDbtTaskValues_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedDbtTaskValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedDbtTaskValues_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"commands": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolvedDbtTaskValues_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolvedDbtTaskValues_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"commands": o.Commands,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolvedDbtTaskValues_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"commands": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetCommands returns the value of the Commands field in ResolvedDbtTaskValues_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedDbtTaskValues_SdkV2) GetCommands(ctx context.Context) ([]types.String, bool) {
	if o.Commands.IsNull() || o.Commands.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Commands.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCommands sets the value of the Commands field in ResolvedDbtTaskValues_SdkV2.
func (o *ResolvedDbtTaskValues_SdkV2) SetCommands(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["commands"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Commands = types.ListValueMust(t, vs)
}

type ResolvedNotebookTaskValues_SdkV2 struct {
	BaseParameters types.Map `tfsdk:"base_parameters" tf:"optional"`
}

func (newState *ResolvedNotebookTaskValues_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedNotebookTaskValues_SdkV2) {
}

func (newState *ResolvedNotebookTaskValues_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResolvedNotebookTaskValues_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedNotebookTaskValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedNotebookTaskValues_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"base_parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolvedNotebookTaskValues_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolvedNotebookTaskValues_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"base_parameters": o.BaseParameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolvedNotebookTaskValues_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"base_parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetBaseParameters returns the value of the BaseParameters field in ResolvedNotebookTaskValues_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedNotebookTaskValues_SdkV2) GetBaseParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.BaseParameters.IsNull() || o.BaseParameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.BaseParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBaseParameters sets the value of the BaseParameters field in ResolvedNotebookTaskValues_SdkV2.
func (o *ResolvedNotebookTaskValues_SdkV2) SetBaseParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["base_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.BaseParameters = types.MapValueMust(t, vs)
}

type ResolvedParamPairValues_SdkV2 struct {
	Parameters types.Map `tfsdk:"parameters" tf:"optional"`
}

func (newState *ResolvedParamPairValues_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedParamPairValues_SdkV2) {
}

func (newState *ResolvedParamPairValues_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResolvedParamPairValues_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedParamPairValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedParamPairValues_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolvedParamPairValues_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolvedParamPairValues_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parameters": o.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolvedParamPairValues_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetParameters returns the value of the Parameters field in ResolvedParamPairValues_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedParamPairValues_SdkV2) GetParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ResolvedParamPairValues_SdkV2.
func (o *ResolvedParamPairValues_SdkV2) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.MapValueMust(t, vs)
}

type ResolvedPythonWheelTaskValues_SdkV2 struct {
	NamedParameters types.Map `tfsdk:"named_parameters" tf:"optional"`

	Parameters types.List `tfsdk:"parameters" tf:"optional"`
}

func (newState *ResolvedPythonWheelTaskValues_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedPythonWheelTaskValues_SdkV2) {
}

func (newState *ResolvedPythonWheelTaskValues_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResolvedPythonWheelTaskValues_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedPythonWheelTaskValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedPythonWheelTaskValues_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"named_parameters": reflect.TypeOf(types.String{}),
		"parameters":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolvedPythonWheelTaskValues_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolvedPythonWheelTaskValues_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"named_parameters": o.NamedParameters,
			"parameters":       o.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolvedPythonWheelTaskValues_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetNamedParameters returns the value of the NamedParameters field in ResolvedPythonWheelTaskValues_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedPythonWheelTaskValues_SdkV2) GetNamedParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.NamedParameters.IsNull() || o.NamedParameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.NamedParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNamedParameters sets the value of the NamedParameters field in ResolvedPythonWheelTaskValues_SdkV2.
func (o *ResolvedPythonWheelTaskValues_SdkV2) SetNamedParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["named_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NamedParameters = types.MapValueMust(t, vs)
}

// GetParameters returns the value of the Parameters field in ResolvedPythonWheelTaskValues_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedPythonWheelTaskValues_SdkV2) GetParameters(ctx context.Context) ([]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ResolvedPythonWheelTaskValues_SdkV2.
func (o *ResolvedPythonWheelTaskValues_SdkV2) SetParameters(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

type ResolvedRunJobTaskValues_SdkV2 struct {
	JobParameters types.Map `tfsdk:"job_parameters" tf:"optional"`

	Parameters types.Map `tfsdk:"parameters" tf:"optional"`
}

func (newState *ResolvedRunJobTaskValues_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedRunJobTaskValues_SdkV2) {
}

func (newState *ResolvedRunJobTaskValues_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResolvedRunJobTaskValues_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedRunJobTaskValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedRunJobTaskValues_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"job_parameters": reflect.TypeOf(types.String{}),
		"parameters":     reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolvedRunJobTaskValues_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolvedRunJobTaskValues_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_parameters": o.JobParameters,
			"parameters":     o.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolvedRunJobTaskValues_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetJobParameters returns the value of the JobParameters field in ResolvedRunJobTaskValues_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedRunJobTaskValues_SdkV2) GetJobParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.JobParameters.IsNull() || o.JobParameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.JobParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobParameters sets the value of the JobParameters field in ResolvedRunJobTaskValues_SdkV2.
func (o *ResolvedRunJobTaskValues_SdkV2) SetJobParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobParameters = types.MapValueMust(t, vs)
}

// GetParameters returns the value of the Parameters field in ResolvedRunJobTaskValues_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedRunJobTaskValues_SdkV2) GetParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ResolvedRunJobTaskValues_SdkV2.
func (o *ResolvedRunJobTaskValues_SdkV2) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.MapValueMust(t, vs)
}

type ResolvedStringParamsValues_SdkV2 struct {
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
}

func (newState *ResolvedStringParamsValues_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedStringParamsValues_SdkV2) {
}

func (newState *ResolvedStringParamsValues_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResolvedStringParamsValues_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedStringParamsValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedStringParamsValues_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolvedStringParamsValues_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolvedStringParamsValues_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parameters": o.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolvedStringParamsValues_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parameters": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetParameters returns the value of the Parameters field in ResolvedStringParamsValues_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedStringParamsValues_SdkV2) GetParameters(ctx context.Context) ([]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in ResolvedStringParamsValues_SdkV2.
func (o *ResolvedStringParamsValues_SdkV2) SetParameters(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

type ResolvedValues_SdkV2 struct {
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

func (newState *ResolvedValues_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResolvedValues_SdkV2) {
}

func (newState *ResolvedValues_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResolvedValues_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResolvedValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResolvedValues_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"condition_task":    reflect.TypeOf(ResolvedConditionTaskValues_SdkV2{}),
		"dbt_task":          reflect.TypeOf(ResolvedDbtTaskValues_SdkV2{}),
		"notebook_task":     reflect.TypeOf(ResolvedNotebookTaskValues_SdkV2{}),
		"python_wheel_task": reflect.TypeOf(ResolvedPythonWheelTaskValues_SdkV2{}),
		"run_job_task":      reflect.TypeOf(ResolvedRunJobTaskValues_SdkV2{}),
		"simulation_task":   reflect.TypeOf(ResolvedParamPairValues_SdkV2{}),
		"spark_jar_task":    reflect.TypeOf(ResolvedStringParamsValues_SdkV2{}),
		"spark_python_task": reflect.TypeOf(ResolvedStringParamsValues_SdkV2{}),
		"spark_submit_task": reflect.TypeOf(ResolvedStringParamsValues_SdkV2{}),
		"sql_task":          reflect.TypeOf(ResolvedParamPairValues_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResolvedValues_SdkV2
// only implements ToObjectValue() and Type().
func (o ResolvedValues_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"condition_task":    o.ConditionTask,
			"dbt_task":          o.DbtTask,
			"notebook_task":     o.NotebookTask,
			"python_wheel_task": o.PythonWheelTask,
			"run_job_task":      o.RunJobTask,
			"simulation_task":   o.SimulationTask,
			"spark_jar_task":    o.SparkJarTask,
			"spark_python_task": o.SparkPythonTask,
			"spark_submit_task": o.SparkSubmitTask,
			"sql_task":          o.SqlTask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResolvedValues_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"condition_task": basetypes.ListType{
				ElemType: ResolvedConditionTaskValues{}.Type(ctx),
			},
			"dbt_task": basetypes.ListType{
				ElemType: ResolvedDbtTaskValues{}.Type(ctx),
			},
			"notebook_task": basetypes.ListType{
				ElemType: ResolvedNotebookTaskValues{}.Type(ctx),
			},
			"python_wheel_task": basetypes.ListType{
				ElemType: ResolvedPythonWheelTaskValues{}.Type(ctx),
			},
			"run_job_task": basetypes.ListType{
				ElemType: ResolvedRunJobTaskValues{}.Type(ctx),
			},
			"simulation_task": basetypes.ListType{
				ElemType: ResolvedParamPairValues{}.Type(ctx),
			},
			"spark_jar_task": basetypes.ListType{
				ElemType: ResolvedStringParamsValues{}.Type(ctx),
			},
			"spark_python_task": basetypes.ListType{
				ElemType: ResolvedStringParamsValues{}.Type(ctx),
			},
			"spark_submit_task": basetypes.ListType{
				ElemType: ResolvedStringParamsValues{}.Type(ctx),
			},
			"sql_task": basetypes.ListType{
				ElemType: ResolvedParamPairValues{}.Type(ctx),
			},
		},
	}
}

// GetConditionTask returns the value of the ConditionTask field in ResolvedValues_SdkV2 as
// a ResolvedConditionTaskValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedValues_SdkV2) GetConditionTask(ctx context.Context) (ResolvedConditionTaskValues_SdkV2, bool) {
	var e ResolvedConditionTaskValues_SdkV2
	if o.ConditionTask.IsNull() || o.ConditionTask.IsUnknown() {
		return e, false
	}
	var v []ResolvedConditionTaskValues_SdkV2
	d := o.ConditionTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConditionTask sets the value of the ConditionTask field in ResolvedValues_SdkV2.
func (o *ResolvedValues_SdkV2) SetConditionTask(ctx context.Context, v ResolvedConditionTaskValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition_task"]
	o.ConditionTask = types.ListValueMust(t, vs)
}

// GetDbtTask returns the value of the DbtTask field in ResolvedValues_SdkV2 as
// a ResolvedDbtTaskValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedValues_SdkV2) GetDbtTask(ctx context.Context) (ResolvedDbtTaskValues_SdkV2, bool) {
	var e ResolvedDbtTaskValues_SdkV2
	if o.DbtTask.IsNull() || o.DbtTask.IsUnknown() {
		return e, false
	}
	var v []ResolvedDbtTaskValues_SdkV2
	d := o.DbtTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbtTask sets the value of the DbtTask field in ResolvedValues_SdkV2.
func (o *ResolvedValues_SdkV2) SetDbtTask(ctx context.Context, v ResolvedDbtTaskValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbt_task"]
	o.DbtTask = types.ListValueMust(t, vs)
}

// GetNotebookTask returns the value of the NotebookTask field in ResolvedValues_SdkV2 as
// a ResolvedNotebookTaskValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedValues_SdkV2) GetNotebookTask(ctx context.Context) (ResolvedNotebookTaskValues_SdkV2, bool) {
	var e ResolvedNotebookTaskValues_SdkV2
	if o.NotebookTask.IsNull() || o.NotebookTask.IsUnknown() {
		return e, false
	}
	var v []ResolvedNotebookTaskValues_SdkV2
	d := o.NotebookTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebookTask sets the value of the NotebookTask field in ResolvedValues_SdkV2.
func (o *ResolvedValues_SdkV2) SetNotebookTask(ctx context.Context, v ResolvedNotebookTaskValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_task"]
	o.NotebookTask = types.ListValueMust(t, vs)
}

// GetPythonWheelTask returns the value of the PythonWheelTask field in ResolvedValues_SdkV2 as
// a ResolvedPythonWheelTaskValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedValues_SdkV2) GetPythonWheelTask(ctx context.Context) (ResolvedPythonWheelTaskValues_SdkV2, bool) {
	var e ResolvedPythonWheelTaskValues_SdkV2
	if o.PythonWheelTask.IsNull() || o.PythonWheelTask.IsUnknown() {
		return e, false
	}
	var v []ResolvedPythonWheelTaskValues_SdkV2
	d := o.PythonWheelTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPythonWheelTask sets the value of the PythonWheelTask field in ResolvedValues_SdkV2.
func (o *ResolvedValues_SdkV2) SetPythonWheelTask(ctx context.Context, v ResolvedPythonWheelTaskValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_wheel_task"]
	o.PythonWheelTask = types.ListValueMust(t, vs)
}

// GetRunJobTask returns the value of the RunJobTask field in ResolvedValues_SdkV2 as
// a ResolvedRunJobTaskValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedValues_SdkV2) GetRunJobTask(ctx context.Context) (ResolvedRunJobTaskValues_SdkV2, bool) {
	var e ResolvedRunJobTaskValues_SdkV2
	if o.RunJobTask.IsNull() || o.RunJobTask.IsUnknown() {
		return e, false
	}
	var v []ResolvedRunJobTaskValues_SdkV2
	d := o.RunJobTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunJobTask sets the value of the RunJobTask field in ResolvedValues_SdkV2.
func (o *ResolvedValues_SdkV2) SetRunJobTask(ctx context.Context, v ResolvedRunJobTaskValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_job_task"]
	o.RunJobTask = types.ListValueMust(t, vs)
}

// GetSimulationTask returns the value of the SimulationTask field in ResolvedValues_SdkV2 as
// a ResolvedParamPairValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedValues_SdkV2) GetSimulationTask(ctx context.Context) (ResolvedParamPairValues_SdkV2, bool) {
	var e ResolvedParamPairValues_SdkV2
	if o.SimulationTask.IsNull() || o.SimulationTask.IsUnknown() {
		return e, false
	}
	var v []ResolvedParamPairValues_SdkV2
	d := o.SimulationTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSimulationTask sets the value of the SimulationTask field in ResolvedValues_SdkV2.
func (o *ResolvedValues_SdkV2) SetSimulationTask(ctx context.Context, v ResolvedParamPairValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["simulation_task"]
	o.SimulationTask = types.ListValueMust(t, vs)
}

// GetSparkJarTask returns the value of the SparkJarTask field in ResolvedValues_SdkV2 as
// a ResolvedStringParamsValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedValues_SdkV2) GetSparkJarTask(ctx context.Context) (ResolvedStringParamsValues_SdkV2, bool) {
	var e ResolvedStringParamsValues_SdkV2
	if o.SparkJarTask.IsNull() || o.SparkJarTask.IsUnknown() {
		return e, false
	}
	var v []ResolvedStringParamsValues_SdkV2
	d := o.SparkJarTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkJarTask sets the value of the SparkJarTask field in ResolvedValues_SdkV2.
func (o *ResolvedValues_SdkV2) SetSparkJarTask(ctx context.Context, v ResolvedStringParamsValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_jar_task"]
	o.SparkJarTask = types.ListValueMust(t, vs)
}

// GetSparkPythonTask returns the value of the SparkPythonTask field in ResolvedValues_SdkV2 as
// a ResolvedStringParamsValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedValues_SdkV2) GetSparkPythonTask(ctx context.Context) (ResolvedStringParamsValues_SdkV2, bool) {
	var e ResolvedStringParamsValues_SdkV2
	if o.SparkPythonTask.IsNull() || o.SparkPythonTask.IsUnknown() {
		return e, false
	}
	var v []ResolvedStringParamsValues_SdkV2
	d := o.SparkPythonTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkPythonTask sets the value of the SparkPythonTask field in ResolvedValues_SdkV2.
func (o *ResolvedValues_SdkV2) SetSparkPythonTask(ctx context.Context, v ResolvedStringParamsValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_python_task"]
	o.SparkPythonTask = types.ListValueMust(t, vs)
}

// GetSparkSubmitTask returns the value of the SparkSubmitTask field in ResolvedValues_SdkV2 as
// a ResolvedStringParamsValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedValues_SdkV2) GetSparkSubmitTask(ctx context.Context) (ResolvedStringParamsValues_SdkV2, bool) {
	var e ResolvedStringParamsValues_SdkV2
	if o.SparkSubmitTask.IsNull() || o.SparkSubmitTask.IsUnknown() {
		return e, false
	}
	var v []ResolvedStringParamsValues_SdkV2
	d := o.SparkSubmitTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkSubmitTask sets the value of the SparkSubmitTask field in ResolvedValues_SdkV2.
func (o *ResolvedValues_SdkV2) SetSparkSubmitTask(ctx context.Context, v ResolvedStringParamsValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_submit_task"]
	o.SparkSubmitTask = types.ListValueMust(t, vs)
}

// GetSqlTask returns the value of the SqlTask field in ResolvedValues_SdkV2 as
// a ResolvedParamPairValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ResolvedValues_SdkV2) GetSqlTask(ctx context.Context) (ResolvedParamPairValues_SdkV2, bool) {
	var e ResolvedParamPairValues_SdkV2
	if o.SqlTask.IsNull() || o.SqlTask.IsUnknown() {
		return e, false
	}
	var v []ResolvedParamPairValues_SdkV2
	d := o.SqlTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlTask sets the value of the SqlTask field in ResolvedValues_SdkV2.
func (o *ResolvedValues_SdkV2) SetSqlTask(ctx context.Context, v ResolvedParamPairValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_task"]
	o.SqlTask = types.ListValueMust(t, vs)
}

// Run was retrieved successfully
type Run_SdkV2 struct {
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
	// * `CONTINUOUS_RESTART`: Indicates a run created by user to manually
	// restart a continuous job run.
	Trigger types.String `tfsdk:"trigger" tf:"optional"`
	// Additional details about what triggered the run
	TriggerInfo types.List `tfsdk:"trigger_info" tf:"optional,object"`
}

func (newState *Run_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Run_SdkV2) {
}

func (newState *Run_SdkV2) SyncEffectiveFieldsDuringRead(existingState Run_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Run.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Run_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cluster_instance":      reflect.TypeOf(ClusterInstance_SdkV2{}),
		"cluster_spec":          reflect.TypeOf(ClusterSpec_SdkV2{}),
		"git_source":            reflect.TypeOf(GitSource_SdkV2{}),
		"iterations":            reflect.TypeOf(RunTask_SdkV2{}),
		"job_clusters":          reflect.TypeOf(JobCluster_SdkV2{}),
		"job_parameters":        reflect.TypeOf(JobParameter_SdkV2{}),
		"overriding_parameters": reflect.TypeOf(RunParameters_SdkV2{}),
		"repair_history":        reflect.TypeOf(RepairHistoryItem_SdkV2{}),
		"schedule":              reflect.TypeOf(CronSchedule_SdkV2{}),
		"state":                 reflect.TypeOf(RunState_SdkV2{}),
		"status":                reflect.TypeOf(RunStatus_SdkV2{}),
		"tasks":                 reflect.TypeOf(RunTask_SdkV2{}),
		"trigger_info":          reflect.TypeOf(TriggerInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Run_SdkV2
// only implements ToObjectValue() and Type().
func (o Run_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attempt_number":          o.AttemptNumber,
			"cleanup_duration":        o.CleanupDuration,
			"cluster_instance":        o.ClusterInstance,
			"cluster_spec":            o.ClusterSpec,
			"creator_user_name":       o.CreatorUserName,
			"description":             o.Description,
			"end_time":                o.EndTime,
			"execution_duration":      o.ExecutionDuration,
			"git_source":              o.GitSource,
			"iterations":              o.Iterations,
			"job_clusters":            o.JobClusters,
			"job_id":                  o.JobId,
			"job_parameters":          o.JobParameters,
			"job_run_id":              o.JobRunId,
			"next_page_token":         o.NextPageToken,
			"number_in_job":           o.NumberInJob,
			"original_attempt_run_id": o.OriginalAttemptRunId,
			"overriding_parameters":   o.OverridingParameters,
			"queue_duration":          o.QueueDuration,
			"repair_history":          o.RepairHistory,
			"run_duration":            o.RunDuration,
			"run_id":                  o.RunId,
			"run_name":                o.RunName,
			"run_page_url":            o.RunPageUrl,
			"run_type":                o.RunType,
			"schedule":                o.Schedule,
			"setup_duration":          o.SetupDuration,
			"start_time":              o.StartTime,
			"state":                   o.State,
			"status":                  o.Status,
			"tasks":                   o.Tasks,
			"trigger":                 o.Trigger,
			"trigger_info":            o.TriggerInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Run_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attempt_number":   types.Int64Type,
			"cleanup_duration": types.Int64Type,
			"cluster_instance": basetypes.ListType{
				ElemType: ClusterInstance{}.Type(ctx),
			},
			"cluster_spec": basetypes.ListType{
				ElemType: ClusterSpec{}.Type(ctx),
			},
			"creator_user_name":  types.StringType,
			"description":        types.StringType,
			"end_time":           types.Int64Type,
			"execution_duration": types.Int64Type,
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.Type(ctx),
			},
			"iterations": basetypes.ListType{
				ElemType: RunTask{}.Type(ctx),
			},
			"job_clusters": basetypes.ListType{
				ElemType: JobCluster{}.Type(ctx),
			},
			"job_id": types.Int64Type,
			"job_parameters": basetypes.ListType{
				ElemType: JobParameter{}.Type(ctx),
			},
			"job_run_id":              types.Int64Type,
			"next_page_token":         types.StringType,
			"number_in_job":           types.Int64Type,
			"original_attempt_run_id": types.Int64Type,
			"overriding_parameters": basetypes.ListType{
				ElemType: RunParameters{}.Type(ctx),
			},
			"queue_duration": types.Int64Type,
			"repair_history": basetypes.ListType{
				ElemType: RepairHistoryItem{}.Type(ctx),
			},
			"run_duration": types.Int64Type,
			"run_id":       types.Int64Type,
			"run_name":     types.StringType,
			"run_page_url": types.StringType,
			"run_type":     types.StringType,
			"schedule": basetypes.ListType{
				ElemType: CronSchedule{}.Type(ctx),
			},
			"setup_duration": types.Int64Type,
			"start_time":     types.Int64Type,
			"state": basetypes.ListType{
				ElemType: RunState{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: RunStatus{}.Type(ctx),
			},
			"tasks": basetypes.ListType{
				ElemType: RunTask{}.Type(ctx),
			},
			"trigger": types.StringType,
			"trigger_info": basetypes.ListType{
				ElemType: TriggerInfo{}.Type(ctx),
			},
		},
	}
}

// GetClusterInstance returns the value of the ClusterInstance field in Run_SdkV2 as
// a ClusterInstance_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetClusterInstance(ctx context.Context) (ClusterInstance_SdkV2, bool) {
	var e ClusterInstance_SdkV2
	if o.ClusterInstance.IsNull() || o.ClusterInstance.IsUnknown() {
		return e, false
	}
	var v []ClusterInstance_SdkV2
	d := o.ClusterInstance.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterInstance sets the value of the ClusterInstance field in Run_SdkV2.
func (o *Run_SdkV2) SetClusterInstance(ctx context.Context, v ClusterInstance_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_instance"]
	o.ClusterInstance = types.ListValueMust(t, vs)
}

// GetClusterSpec returns the value of the ClusterSpec field in Run_SdkV2 as
// a ClusterSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetClusterSpec(ctx context.Context) (ClusterSpec_SdkV2, bool) {
	var e ClusterSpec_SdkV2
	if o.ClusterSpec.IsNull() || o.ClusterSpec.IsUnknown() {
		return e, false
	}
	var v []ClusterSpec_SdkV2
	d := o.ClusterSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterSpec sets the value of the ClusterSpec field in Run_SdkV2.
func (o *Run_SdkV2) SetClusterSpec(ctx context.Context, v ClusterSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_spec"]
	o.ClusterSpec = types.ListValueMust(t, vs)
}

// GetGitSource returns the value of the GitSource field in Run_SdkV2 as
// a GitSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetGitSource(ctx context.Context) (GitSource_SdkV2, bool) {
	var e GitSource_SdkV2
	if o.GitSource.IsNull() || o.GitSource.IsUnknown() {
		return e, false
	}
	var v []GitSource_SdkV2
	d := o.GitSource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitSource sets the value of the GitSource field in Run_SdkV2.
func (o *Run_SdkV2) SetGitSource(ctx context.Context, v GitSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["git_source"]
	o.GitSource = types.ListValueMust(t, vs)
}

// GetIterations returns the value of the Iterations field in Run_SdkV2 as
// a slice of RunTask_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetIterations(ctx context.Context) ([]RunTask_SdkV2, bool) {
	if o.Iterations.IsNull() || o.Iterations.IsUnknown() {
		return nil, false
	}
	var v []RunTask_SdkV2
	d := o.Iterations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIterations sets the value of the Iterations field in Run_SdkV2.
func (o *Run_SdkV2) SetIterations(ctx context.Context, v []RunTask_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["iterations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Iterations = types.ListValueMust(t, vs)
}

// GetJobClusters returns the value of the JobClusters field in Run_SdkV2 as
// a slice of JobCluster_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetJobClusters(ctx context.Context) ([]JobCluster_SdkV2, bool) {
	if o.JobClusters.IsNull() || o.JobClusters.IsUnknown() {
		return nil, false
	}
	var v []JobCluster_SdkV2
	d := o.JobClusters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobClusters sets the value of the JobClusters field in Run_SdkV2.
func (o *Run_SdkV2) SetJobClusters(ctx context.Context, v []JobCluster_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_clusters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobClusters = types.ListValueMust(t, vs)
}

// GetJobParameters returns the value of the JobParameters field in Run_SdkV2 as
// a slice of JobParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetJobParameters(ctx context.Context) ([]JobParameter_SdkV2, bool) {
	if o.JobParameters.IsNull() || o.JobParameters.IsUnknown() {
		return nil, false
	}
	var v []JobParameter_SdkV2
	d := o.JobParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobParameters sets the value of the JobParameters field in Run_SdkV2.
func (o *Run_SdkV2) SetJobParameters(ctx context.Context, v []JobParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobParameters = types.ListValueMust(t, vs)
}

// GetOverridingParameters returns the value of the OverridingParameters field in Run_SdkV2 as
// a RunParameters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetOverridingParameters(ctx context.Context) (RunParameters_SdkV2, bool) {
	var e RunParameters_SdkV2
	if o.OverridingParameters.IsNull() || o.OverridingParameters.IsUnknown() {
		return e, false
	}
	var v []RunParameters_SdkV2
	d := o.OverridingParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOverridingParameters sets the value of the OverridingParameters field in Run_SdkV2.
func (o *Run_SdkV2) SetOverridingParameters(ctx context.Context, v RunParameters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["overriding_parameters"]
	o.OverridingParameters = types.ListValueMust(t, vs)
}

// GetRepairHistory returns the value of the RepairHistory field in Run_SdkV2 as
// a slice of RepairHistoryItem_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetRepairHistory(ctx context.Context) ([]RepairHistoryItem_SdkV2, bool) {
	if o.RepairHistory.IsNull() || o.RepairHistory.IsUnknown() {
		return nil, false
	}
	var v []RepairHistoryItem_SdkV2
	d := o.RepairHistory.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRepairHistory sets the value of the RepairHistory field in Run_SdkV2.
func (o *Run_SdkV2) SetRepairHistory(ctx context.Context, v []RepairHistoryItem_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["repair_history"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RepairHistory = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in Run_SdkV2 as
// a CronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetSchedule(ctx context.Context) (CronSchedule_SdkV2, bool) {
	var e CronSchedule_SdkV2
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []CronSchedule_SdkV2
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in Run_SdkV2.
func (o *Run_SdkV2) SetSchedule(ctx context.Context, v CronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}

// GetState returns the value of the State field in Run_SdkV2 as
// a RunState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetState(ctx context.Context) (RunState_SdkV2, bool) {
	var e RunState_SdkV2
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v []RunState_SdkV2
	d := o.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in Run_SdkV2.
func (o *Run_SdkV2) SetState(ctx context.Context, v RunState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	o.State = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Run_SdkV2 as
// a RunStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetStatus(ctx context.Context) (RunStatus_SdkV2, bool) {
	var e RunStatus_SdkV2
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []RunStatus_SdkV2
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in Run_SdkV2.
func (o *Run_SdkV2) SetStatus(ctx context.Context, v RunStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
}

// GetTasks returns the value of the Tasks field in Run_SdkV2 as
// a slice of RunTask_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetTasks(ctx context.Context) ([]RunTask_SdkV2, bool) {
	if o.Tasks.IsNull() || o.Tasks.IsUnknown() {
		return nil, false
	}
	var v []RunTask_SdkV2
	d := o.Tasks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTasks sets the value of the Tasks field in Run_SdkV2.
func (o *Run_SdkV2) SetTasks(ctx context.Context, v []RunTask_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tasks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tasks = types.ListValueMust(t, vs)
}

// GetTriggerInfo returns the value of the TriggerInfo field in Run_SdkV2 as
// a TriggerInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetTriggerInfo(ctx context.Context) (TriggerInfo_SdkV2, bool) {
	var e TriggerInfo_SdkV2
	if o.TriggerInfo.IsNull() || o.TriggerInfo.IsUnknown() {
		return e, false
	}
	var v []TriggerInfo_SdkV2
	d := o.TriggerInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTriggerInfo sets the value of the TriggerInfo field in Run_SdkV2.
func (o *Run_SdkV2) SetTriggerInfo(ctx context.Context, v TriggerInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["trigger_info"]
	o.TriggerInfo = types.ListValueMust(t, vs)
}

type RunConditionTask_SdkV2 struct {
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

func (newState *RunConditionTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunConditionTask_SdkV2) {
}

func (newState *RunConditionTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunConditionTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunConditionTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunConditionTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunConditionTask_SdkV2
// only implements ToObjectValue() and Type().
func (o RunConditionTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"left":    o.Left,
			"op":      o.Op,
			"outcome": o.Outcome,
			"right":   o.Right,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunConditionTask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"left":    types.StringType,
			"op":      types.StringType,
			"outcome": types.StringType,
			"right":   types.StringType,
		},
	}
}

type RunForEachTask_SdkV2 struct {
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

func (newState *RunForEachTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunForEachTask_SdkV2) {
}

func (newState *RunForEachTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunForEachTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunForEachTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunForEachTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stats": reflect.TypeOf(ForEachStats_SdkV2{}),
		"task":  reflect.TypeOf(Task_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunForEachTask_SdkV2
// only implements ToObjectValue() and Type().
func (o RunForEachTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"concurrency": o.Concurrency,
			"inputs":      o.Inputs,
			"stats":       o.Stats,
			"task":        o.Task,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunForEachTask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"concurrency": types.Int64Type,
			"inputs":      types.StringType,
			"stats": basetypes.ListType{
				ElemType: ForEachStats{}.Type(ctx),
			},
			"task": basetypes.ListType{
				ElemType: Task{}.Type(ctx),
			},
		},
	}
}

// GetStats returns the value of the Stats field in RunForEachTask_SdkV2 as
// a ForEachStats_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunForEachTask_SdkV2) GetStats(ctx context.Context) (ForEachStats_SdkV2, bool) {
	var e ForEachStats_SdkV2
	if o.Stats.IsNull() || o.Stats.IsUnknown() {
		return e, false
	}
	var v []ForEachStats_SdkV2
	d := o.Stats.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStats sets the value of the Stats field in RunForEachTask_SdkV2.
func (o *RunForEachTask_SdkV2) SetStats(ctx context.Context, v ForEachStats_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["stats"]
	o.Stats = types.ListValueMust(t, vs)
}

// GetTask returns the value of the Task field in RunForEachTask_SdkV2 as
// a Task_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunForEachTask_SdkV2) GetTask(ctx context.Context) (Task_SdkV2, bool) {
	var e Task_SdkV2
	if o.Task.IsNull() || o.Task.IsUnknown() {
		return e, false
	}
	var v []Task_SdkV2
	d := o.Task.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTask sets the value of the Task field in RunForEachTask_SdkV2.
func (o *RunForEachTask_SdkV2) SetTask(ctx context.Context, v Task_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["task"]
	o.Task = types.ListValueMust(t, vs)
}

type RunJobOutput_SdkV2 struct {
	// The run id of the triggered job run
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
}

func (newState *RunJobOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunJobOutput_SdkV2) {
}

func (newState *RunJobOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunJobOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunJobOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunJobOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunJobOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o RunJobOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunJobOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

type RunJobTask_SdkV2 struct {
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

func (newState *RunJobTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunJobTask_SdkV2) {
}

func (newState *RunJobTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunJobTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunJobTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunJobTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbt_commands":        reflect.TypeOf(types.String{}),
		"jar_params":          reflect.TypeOf(types.String{}),
		"job_parameters":      reflect.TypeOf(types.String{}),
		"notebook_params":     reflect.TypeOf(types.String{}),
		"pipeline_params":     reflect.TypeOf(PipelineParams_SdkV2{}),
		"python_named_params": reflect.TypeOf(types.String{}),
		"python_params":       reflect.TypeOf(types.String{}),
		"spark_submit_params": reflect.TypeOf(types.String{}),
		"sql_params":          reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunJobTask_SdkV2
// only implements ToObjectValue() and Type().
func (o RunJobTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbt_commands":        o.DbtCommands,
			"jar_params":          o.JarParams,
			"job_id":              o.JobId,
			"job_parameters":      o.JobParameters,
			"notebook_params":     o.NotebookParams,
			"pipeline_params":     o.PipelineParams,
			"python_named_params": o.PythonNamedParams,
			"python_params":       o.PythonParams,
			"spark_submit_params": o.SparkSubmitParams,
			"sql_params":          o.SqlParams,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunJobTask_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: PipelineParams{}.Type(ctx),
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

// GetDbtCommands returns the value of the DbtCommands field in RunJobTask_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunJobTask_SdkV2) GetDbtCommands(ctx context.Context) ([]types.String, bool) {
	if o.DbtCommands.IsNull() || o.DbtCommands.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DbtCommands.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDbtCommands sets the value of the DbtCommands field in RunJobTask_SdkV2.
func (o *RunJobTask_SdkV2) SetDbtCommands(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbt_commands"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DbtCommands = types.ListValueMust(t, vs)
}

// GetJarParams returns the value of the JarParams field in RunJobTask_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunJobTask_SdkV2) GetJarParams(ctx context.Context) ([]types.String, bool) {
	if o.JarParams.IsNull() || o.JarParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.JarParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJarParams sets the value of the JarParams field in RunJobTask_SdkV2.
func (o *RunJobTask_SdkV2) SetJarParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["jar_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JarParams = types.ListValueMust(t, vs)
}

// GetJobParameters returns the value of the JobParameters field in RunJobTask_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunJobTask_SdkV2) GetJobParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.JobParameters.IsNull() || o.JobParameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.JobParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobParameters sets the value of the JobParameters field in RunJobTask_SdkV2.
func (o *RunJobTask_SdkV2) SetJobParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobParameters = types.MapValueMust(t, vs)
}

// GetNotebookParams returns the value of the NotebookParams field in RunJobTask_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunJobTask_SdkV2) GetNotebookParams(ctx context.Context) (map[string]types.String, bool) {
	if o.NotebookParams.IsNull() || o.NotebookParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.NotebookParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookParams sets the value of the NotebookParams field in RunJobTask_SdkV2.
func (o *RunJobTask_SdkV2) SetNotebookParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NotebookParams = types.MapValueMust(t, vs)
}

// GetPipelineParams returns the value of the PipelineParams field in RunJobTask_SdkV2 as
// a PipelineParams_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunJobTask_SdkV2) GetPipelineParams(ctx context.Context) (PipelineParams_SdkV2, bool) {
	var e PipelineParams_SdkV2
	if o.PipelineParams.IsNull() || o.PipelineParams.IsUnknown() {
		return e, false
	}
	var v []PipelineParams_SdkV2
	d := o.PipelineParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPipelineParams sets the value of the PipelineParams field in RunJobTask_SdkV2.
func (o *RunJobTask_SdkV2) SetPipelineParams(ctx context.Context, v PipelineParams_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pipeline_params"]
	o.PipelineParams = types.ListValueMust(t, vs)
}

// GetPythonNamedParams returns the value of the PythonNamedParams field in RunJobTask_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunJobTask_SdkV2) GetPythonNamedParams(ctx context.Context) (map[string]types.String, bool) {
	if o.PythonNamedParams.IsNull() || o.PythonNamedParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.PythonNamedParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPythonNamedParams sets the value of the PythonNamedParams field in RunJobTask_SdkV2.
func (o *RunJobTask_SdkV2) SetPythonNamedParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_named_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PythonNamedParams = types.MapValueMust(t, vs)
}

// GetPythonParams returns the value of the PythonParams field in RunJobTask_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunJobTask_SdkV2) GetPythonParams(ctx context.Context) ([]types.String, bool) {
	if o.PythonParams.IsNull() || o.PythonParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PythonParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPythonParams sets the value of the PythonParams field in RunJobTask_SdkV2.
func (o *RunJobTask_SdkV2) SetPythonParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PythonParams = types.ListValueMust(t, vs)
}

// GetSparkSubmitParams returns the value of the SparkSubmitParams field in RunJobTask_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunJobTask_SdkV2) GetSparkSubmitParams(ctx context.Context) ([]types.String, bool) {
	if o.SparkSubmitParams.IsNull() || o.SparkSubmitParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SparkSubmitParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkSubmitParams sets the value of the SparkSubmitParams field in RunJobTask_SdkV2.
func (o *RunJobTask_SdkV2) SetSparkSubmitParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_submit_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkSubmitParams = types.ListValueMust(t, vs)
}

// GetSqlParams returns the value of the SqlParams field in RunJobTask_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunJobTask_SdkV2) GetSqlParams(ctx context.Context) (map[string]types.String, bool) {
	if o.SqlParams.IsNull() || o.SqlParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SqlParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlParams sets the value of the SqlParams field in RunJobTask_SdkV2.
func (o *RunJobTask_SdkV2) SetSqlParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SqlParams = types.MapValueMust(t, vs)
}

type RunNow_SdkV2 struct {
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

func (newState *RunNow_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunNow_SdkV2) {
}

func (newState *RunNow_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunNow_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunNow.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunNow_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbt_commands":        reflect.TypeOf(types.String{}),
		"jar_params":          reflect.TypeOf(types.String{}),
		"job_parameters":      reflect.TypeOf(types.String{}),
		"notebook_params":     reflect.TypeOf(types.String{}),
		"only":                reflect.TypeOf(types.String{}),
		"pipeline_params":     reflect.TypeOf(PipelineParams_SdkV2{}),
		"python_named_params": reflect.TypeOf(types.String{}),
		"python_params":       reflect.TypeOf(types.String{}),
		"queue":               reflect.TypeOf(QueueSettings_SdkV2{}),
		"spark_submit_params": reflect.TypeOf(types.String{}),
		"sql_params":          reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunNow_SdkV2
// only implements ToObjectValue() and Type().
func (o RunNow_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbt_commands":        o.DbtCommands,
			"idempotency_token":   o.IdempotencyToken,
			"jar_params":          o.JarParams,
			"job_id":              o.JobId,
			"job_parameters":      o.JobParameters,
			"notebook_params":     o.NotebookParams,
			"only":                o.Only,
			"pipeline_params":     o.PipelineParams,
			"python_named_params": o.PythonNamedParams,
			"python_params":       o.PythonParams,
			"queue":               o.Queue,
			"spark_submit_params": o.SparkSubmitParams,
			"sql_params":          o.SqlParams,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunNow_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: PipelineParams{}.Type(ctx),
			},
			"python_named_params": basetypes.MapType{
				ElemType: types.StringType,
			},
			"python_params": basetypes.ListType{
				ElemType: types.StringType,
			},
			"queue": basetypes.ListType{
				ElemType: QueueSettings{}.Type(ctx),
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

// GetDbtCommands returns the value of the DbtCommands field in RunNow_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetDbtCommands(ctx context.Context) ([]types.String, bool) {
	if o.DbtCommands.IsNull() || o.DbtCommands.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DbtCommands.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDbtCommands sets the value of the DbtCommands field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetDbtCommands(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbt_commands"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DbtCommands = types.ListValueMust(t, vs)
}

// GetJarParams returns the value of the JarParams field in RunNow_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetJarParams(ctx context.Context) ([]types.String, bool) {
	if o.JarParams.IsNull() || o.JarParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.JarParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJarParams sets the value of the JarParams field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetJarParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["jar_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JarParams = types.ListValueMust(t, vs)
}

// GetJobParameters returns the value of the JobParameters field in RunNow_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetJobParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.JobParameters.IsNull() || o.JobParameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.JobParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobParameters sets the value of the JobParameters field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetJobParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JobParameters = types.MapValueMust(t, vs)
}

// GetNotebookParams returns the value of the NotebookParams field in RunNow_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetNotebookParams(ctx context.Context) (map[string]types.String, bool) {
	if o.NotebookParams.IsNull() || o.NotebookParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.NotebookParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookParams sets the value of the NotebookParams field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetNotebookParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NotebookParams = types.MapValueMust(t, vs)
}

// GetOnly returns the value of the Only field in RunNow_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetOnly(ctx context.Context) ([]types.String, bool) {
	if o.Only.IsNull() || o.Only.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Only.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnly sets the value of the Only field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetOnly(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["only"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Only = types.ListValueMust(t, vs)
}

// GetPipelineParams returns the value of the PipelineParams field in RunNow_SdkV2 as
// a PipelineParams_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetPipelineParams(ctx context.Context) (PipelineParams_SdkV2, bool) {
	var e PipelineParams_SdkV2
	if o.PipelineParams.IsNull() || o.PipelineParams.IsUnknown() {
		return e, false
	}
	var v []PipelineParams_SdkV2
	d := o.PipelineParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPipelineParams sets the value of the PipelineParams field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetPipelineParams(ctx context.Context, v PipelineParams_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pipeline_params"]
	o.PipelineParams = types.ListValueMust(t, vs)
}

// GetPythonNamedParams returns the value of the PythonNamedParams field in RunNow_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetPythonNamedParams(ctx context.Context) (map[string]types.String, bool) {
	if o.PythonNamedParams.IsNull() || o.PythonNamedParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.PythonNamedParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPythonNamedParams sets the value of the PythonNamedParams field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetPythonNamedParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_named_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PythonNamedParams = types.MapValueMust(t, vs)
}

// GetPythonParams returns the value of the PythonParams field in RunNow_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetPythonParams(ctx context.Context) ([]types.String, bool) {
	if o.PythonParams.IsNull() || o.PythonParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PythonParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPythonParams sets the value of the PythonParams field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetPythonParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PythonParams = types.ListValueMust(t, vs)
}

// GetQueue returns the value of the Queue field in RunNow_SdkV2 as
// a QueueSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetQueue(ctx context.Context) (QueueSettings_SdkV2, bool) {
	var e QueueSettings_SdkV2
	if o.Queue.IsNull() || o.Queue.IsUnknown() {
		return e, false
	}
	var v []QueueSettings_SdkV2
	d := o.Queue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueue sets the value of the Queue field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetQueue(ctx context.Context, v QueueSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["queue"]
	o.Queue = types.ListValueMust(t, vs)
}

// GetSparkSubmitParams returns the value of the SparkSubmitParams field in RunNow_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetSparkSubmitParams(ctx context.Context) ([]types.String, bool) {
	if o.SparkSubmitParams.IsNull() || o.SparkSubmitParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SparkSubmitParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkSubmitParams sets the value of the SparkSubmitParams field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetSparkSubmitParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_submit_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkSubmitParams = types.ListValueMust(t, vs)
}

// GetSqlParams returns the value of the SqlParams field in RunNow_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunNow_SdkV2) GetSqlParams(ctx context.Context) (map[string]types.String, bool) {
	if o.SqlParams.IsNull() || o.SqlParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SqlParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlParams sets the value of the SqlParams field in RunNow_SdkV2.
func (o *RunNow_SdkV2) SetSqlParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SqlParams = types.MapValueMust(t, vs)
}

// Run was started successfully.
type RunNowResponse_SdkV2 struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob types.Int64 `tfsdk:"number_in_job" tf:"optional"`
	// The globally unique ID of the newly triggered run.
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
}

func (newState *RunNowResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunNowResponse_SdkV2) {
}

func (newState *RunNowResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunNowResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunNowResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunNowResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunNowResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RunNowResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"number_in_job": o.NumberInJob,
			"run_id":        o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunNowResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"number_in_job": types.Int64Type,
			"run_id":        types.Int64Type,
		},
	}
}

// Run output was retrieved successfully.
type RunOutput_SdkV2 struct {
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

func (newState *RunOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunOutput_SdkV2) {
}

func (newState *RunOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbt_output":      reflect.TypeOf(DbtOutput_SdkV2{}),
		"metadata":        reflect.TypeOf(Run_SdkV2{}),
		"notebook_output": reflect.TypeOf(NotebookOutput_SdkV2{}),
		"run_job_output":  reflect.TypeOf(RunJobOutput_SdkV2{}),
		"sql_output":      reflect.TypeOf(SqlOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o RunOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbt_output":      o.DbtOutput,
			"error":           o.Error,
			"error_trace":     o.ErrorTrace,
			"info":            o.Info,
			"logs":            o.Logs,
			"logs_truncated":  o.LogsTruncated,
			"metadata":        o.Metadata,
			"notebook_output": o.NotebookOutput,
			"run_job_output":  o.RunJobOutput,
			"sql_output":      o.SqlOutput,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dbt_output": basetypes.ListType{
				ElemType: DbtOutput{}.Type(ctx),
			},
			"error":          types.StringType,
			"error_trace":    types.StringType,
			"info":           types.StringType,
			"logs":           types.StringType,
			"logs_truncated": types.BoolType,
			"metadata": basetypes.ListType{
				ElemType: Run{}.Type(ctx),
			},
			"notebook_output": basetypes.ListType{
				ElemType: NotebookOutput{}.Type(ctx),
			},
			"run_job_output": basetypes.ListType{
				ElemType: RunJobOutput{}.Type(ctx),
			},
			"sql_output": basetypes.ListType{
				ElemType: SqlOutput{}.Type(ctx),
			},
		},
	}
}

// GetDbtOutput returns the value of the DbtOutput field in RunOutput_SdkV2 as
// a DbtOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunOutput_SdkV2) GetDbtOutput(ctx context.Context) (DbtOutput_SdkV2, bool) {
	var e DbtOutput_SdkV2
	if o.DbtOutput.IsNull() || o.DbtOutput.IsUnknown() {
		return e, false
	}
	var v []DbtOutput_SdkV2
	d := o.DbtOutput.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbtOutput sets the value of the DbtOutput field in RunOutput_SdkV2.
func (o *RunOutput_SdkV2) SetDbtOutput(ctx context.Context, v DbtOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbt_output"]
	o.DbtOutput = types.ListValueMust(t, vs)
}

// GetMetadata returns the value of the Metadata field in RunOutput_SdkV2 as
// a Run_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunOutput_SdkV2) GetMetadata(ctx context.Context) (Run_SdkV2, bool) {
	var e Run_SdkV2
	if o.Metadata.IsNull() || o.Metadata.IsUnknown() {
		return e, false
	}
	var v []Run_SdkV2
	d := o.Metadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetadata sets the value of the Metadata field in RunOutput_SdkV2.
func (o *RunOutput_SdkV2) SetMetadata(ctx context.Context, v Run_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metadata"]
	o.Metadata = types.ListValueMust(t, vs)
}

// GetNotebookOutput returns the value of the NotebookOutput field in RunOutput_SdkV2 as
// a NotebookOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunOutput_SdkV2) GetNotebookOutput(ctx context.Context) (NotebookOutput_SdkV2, bool) {
	var e NotebookOutput_SdkV2
	if o.NotebookOutput.IsNull() || o.NotebookOutput.IsUnknown() {
		return e, false
	}
	var v []NotebookOutput_SdkV2
	d := o.NotebookOutput.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebookOutput sets the value of the NotebookOutput field in RunOutput_SdkV2.
func (o *RunOutput_SdkV2) SetNotebookOutput(ctx context.Context, v NotebookOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_output"]
	o.NotebookOutput = types.ListValueMust(t, vs)
}

// GetRunJobOutput returns the value of the RunJobOutput field in RunOutput_SdkV2 as
// a RunJobOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunOutput_SdkV2) GetRunJobOutput(ctx context.Context) (RunJobOutput_SdkV2, bool) {
	var e RunJobOutput_SdkV2
	if o.RunJobOutput.IsNull() || o.RunJobOutput.IsUnknown() {
		return e, false
	}
	var v []RunJobOutput_SdkV2
	d := o.RunJobOutput.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunJobOutput sets the value of the RunJobOutput field in RunOutput_SdkV2.
func (o *RunOutput_SdkV2) SetRunJobOutput(ctx context.Context, v RunJobOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_job_output"]
	o.RunJobOutput = types.ListValueMust(t, vs)
}

// GetSqlOutput returns the value of the SqlOutput field in RunOutput_SdkV2 as
// a SqlOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunOutput_SdkV2) GetSqlOutput(ctx context.Context) (SqlOutput_SdkV2, bool) {
	var e SqlOutput_SdkV2
	if o.SqlOutput.IsNull() || o.SqlOutput.IsUnknown() {
		return e, false
	}
	var v []SqlOutput_SdkV2
	d := o.SqlOutput.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlOutput sets the value of the SqlOutput field in RunOutput_SdkV2.
func (o *RunOutput_SdkV2) SetSqlOutput(ctx context.Context, v SqlOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_output"]
	o.SqlOutput = types.ListValueMust(t, vs)
}

type RunParameters_SdkV2 struct {
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

func (newState *RunParameters_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunParameters_SdkV2) {
}

func (newState *RunParameters_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunParameters_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunParameters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunParameters_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dbt_commands":        reflect.TypeOf(types.String{}),
		"jar_params":          reflect.TypeOf(types.String{}),
		"notebook_params":     reflect.TypeOf(types.String{}),
		"pipeline_params":     reflect.TypeOf(PipelineParams_SdkV2{}),
		"python_named_params": reflect.TypeOf(types.String{}),
		"python_params":       reflect.TypeOf(types.String{}),
		"spark_submit_params": reflect.TypeOf(types.String{}),
		"sql_params":          reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunParameters_SdkV2
// only implements ToObjectValue() and Type().
func (o RunParameters_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dbt_commands":        o.DbtCommands,
			"jar_params":          o.JarParams,
			"notebook_params":     o.NotebookParams,
			"pipeline_params":     o.PipelineParams,
			"python_named_params": o.PythonNamedParams,
			"python_params":       o.PythonParams,
			"spark_submit_params": o.SparkSubmitParams,
			"sql_params":          o.SqlParams,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunParameters_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: PipelineParams{}.Type(ctx),
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

// GetDbtCommands returns the value of the DbtCommands field in RunParameters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunParameters_SdkV2) GetDbtCommands(ctx context.Context) ([]types.String, bool) {
	if o.DbtCommands.IsNull() || o.DbtCommands.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DbtCommands.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDbtCommands sets the value of the DbtCommands field in RunParameters_SdkV2.
func (o *RunParameters_SdkV2) SetDbtCommands(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbt_commands"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DbtCommands = types.ListValueMust(t, vs)
}

// GetJarParams returns the value of the JarParams field in RunParameters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunParameters_SdkV2) GetJarParams(ctx context.Context) ([]types.String, bool) {
	if o.JarParams.IsNull() || o.JarParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.JarParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJarParams sets the value of the JarParams field in RunParameters_SdkV2.
func (o *RunParameters_SdkV2) SetJarParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["jar_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.JarParams = types.ListValueMust(t, vs)
}

// GetNotebookParams returns the value of the NotebookParams field in RunParameters_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunParameters_SdkV2) GetNotebookParams(ctx context.Context) (map[string]types.String, bool) {
	if o.NotebookParams.IsNull() || o.NotebookParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.NotebookParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookParams sets the value of the NotebookParams field in RunParameters_SdkV2.
func (o *RunParameters_SdkV2) SetNotebookParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NotebookParams = types.MapValueMust(t, vs)
}

// GetPipelineParams returns the value of the PipelineParams field in RunParameters_SdkV2 as
// a PipelineParams_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunParameters_SdkV2) GetPipelineParams(ctx context.Context) (PipelineParams_SdkV2, bool) {
	var e PipelineParams_SdkV2
	if o.PipelineParams.IsNull() || o.PipelineParams.IsUnknown() {
		return e, false
	}
	var v []PipelineParams_SdkV2
	d := o.PipelineParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPipelineParams sets the value of the PipelineParams field in RunParameters_SdkV2.
func (o *RunParameters_SdkV2) SetPipelineParams(ctx context.Context, v PipelineParams_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pipeline_params"]
	o.PipelineParams = types.ListValueMust(t, vs)
}

// GetPythonNamedParams returns the value of the PythonNamedParams field in RunParameters_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunParameters_SdkV2) GetPythonNamedParams(ctx context.Context) (map[string]types.String, bool) {
	if o.PythonNamedParams.IsNull() || o.PythonNamedParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.PythonNamedParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPythonNamedParams sets the value of the PythonNamedParams field in RunParameters_SdkV2.
func (o *RunParameters_SdkV2) SetPythonNamedParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_named_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PythonNamedParams = types.MapValueMust(t, vs)
}

// GetPythonParams returns the value of the PythonParams field in RunParameters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunParameters_SdkV2) GetPythonParams(ctx context.Context) ([]types.String, bool) {
	if o.PythonParams.IsNull() || o.PythonParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PythonParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPythonParams sets the value of the PythonParams field in RunParameters_SdkV2.
func (o *RunParameters_SdkV2) SetPythonParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PythonParams = types.ListValueMust(t, vs)
}

// GetSparkSubmitParams returns the value of the SparkSubmitParams field in RunParameters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunParameters_SdkV2) GetSparkSubmitParams(ctx context.Context) ([]types.String, bool) {
	if o.SparkSubmitParams.IsNull() || o.SparkSubmitParams.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SparkSubmitParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparkSubmitParams sets the value of the SparkSubmitParams field in RunParameters_SdkV2.
func (o *RunParameters_SdkV2) SetSparkSubmitParams(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_submit_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SparkSubmitParams = types.ListValueMust(t, vs)
}

// GetSqlParams returns the value of the SqlParams field in RunParameters_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunParameters_SdkV2) GetSqlParams(ctx context.Context) (map[string]types.String, bool) {
	if o.SqlParams.IsNull() || o.SqlParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.SqlParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlParams sets the value of the SqlParams field in RunParameters_SdkV2.
func (o *RunParameters_SdkV2) SetSqlParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SqlParams = types.MapValueMust(t, vs)
}

// The current state of the run.
type RunState_SdkV2 struct {
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

func (newState *RunState_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunState_SdkV2) {
}

func (newState *RunState_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunState_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunState.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunState_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunState_SdkV2
// only implements ToObjectValue() and Type().
func (o RunState_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"life_cycle_state":           o.LifeCycleState,
			"queue_reason":               o.QueueReason,
			"result_state":               o.ResultState,
			"state_message":              o.StateMessage,
			"user_cancelled_or_timedout": o.UserCancelledOrTimedout,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunState_SdkV2) Type(ctx context.Context) attr.Type {
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
type RunStatus_SdkV2 struct {
	// If the run was queued, details about the reason for queuing the run.
	QueueDetails types.List `tfsdk:"queue_details" tf:"optional,object"`
	// The current state of the run.
	State types.String `tfsdk:"state" tf:"optional"`
	// If the run is in a TERMINATING or TERMINATED state, details about the
	// reason for terminating the run.
	TerminationDetails types.List `tfsdk:"termination_details" tf:"optional,object"`
}

func (newState *RunStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunStatus_SdkV2) {
}

func (newState *RunStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunStatus_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"queue_details":       reflect.TypeOf(QueueDetails_SdkV2{}),
		"termination_details": reflect.TypeOf(TerminationDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o RunStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"queue_details":       o.QueueDetails,
			"state":               o.State,
			"termination_details": o.TerminationDetails,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"queue_details": basetypes.ListType{
				ElemType: QueueDetails{}.Type(ctx),
			},
			"state": types.StringType,
			"termination_details": basetypes.ListType{
				ElemType: TerminationDetails{}.Type(ctx),
			},
		},
	}
}

// GetQueueDetails returns the value of the QueueDetails field in RunStatus_SdkV2 as
// a QueueDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunStatus_SdkV2) GetQueueDetails(ctx context.Context) (QueueDetails_SdkV2, bool) {
	var e QueueDetails_SdkV2
	if o.QueueDetails.IsNull() || o.QueueDetails.IsUnknown() {
		return e, false
	}
	var v []QueueDetails_SdkV2
	d := o.QueueDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueueDetails sets the value of the QueueDetails field in RunStatus_SdkV2.
func (o *RunStatus_SdkV2) SetQueueDetails(ctx context.Context, v QueueDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["queue_details"]
	o.QueueDetails = types.ListValueMust(t, vs)
}

// GetTerminationDetails returns the value of the TerminationDetails field in RunStatus_SdkV2 as
// a TerminationDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunStatus_SdkV2) GetTerminationDetails(ctx context.Context) (TerminationDetails_SdkV2, bool) {
	var e TerminationDetails_SdkV2
	if o.TerminationDetails.IsNull() || o.TerminationDetails.IsUnknown() {
		return e, false
	}
	var v []TerminationDetails_SdkV2
	d := o.TerminationDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTerminationDetails sets the value of the TerminationDetails field in RunStatus_SdkV2.
func (o *RunStatus_SdkV2) SetTerminationDetails(ctx context.Context, v TerminationDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["termination_details"]
	o.TerminationDetails = types.ListValueMust(t, vs)
}

// Used when outputting a child run, in GetRun or ListRuns.
type RunTask_SdkV2 struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber types.Int64 `tfsdk:"attempt_number" tf:"optional"`
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	CleanRoomsNotebookTask types.List `tfsdk:"clean_rooms_notebook_task" tf:"optional,object"`
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
	NewCluster types.List `tfsdk:"new_cluster" tf:"optional,object"`
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

func (newState *RunTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunTask_SdkV2) {
}

func (newState *RunTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_rooms_notebook_task": reflect.TypeOf(CleanRoomsNotebookTask_SdkV2{}),
		"cluster_instance":          reflect.TypeOf(ClusterInstance_SdkV2{}),
		"condition_task":            reflect.TypeOf(RunConditionTask_SdkV2{}),
		"dbt_task":                  reflect.TypeOf(DbtTask_SdkV2{}),
		"depends_on":                reflect.TypeOf(TaskDependency_SdkV2{}),
		"email_notifications":       reflect.TypeOf(JobEmailNotifications_SdkV2{}),
		"for_each_task":             reflect.TypeOf(RunForEachTask_SdkV2{}),
		"git_source":                reflect.TypeOf(GitSource_SdkV2{}),
		"library":                   reflect.TypeOf(compute_tf.Library_SdkV2{}),
		"new_cluster":               reflect.TypeOf(compute_tf.ClusterSpec_SdkV2{}),
		"notebook_task":             reflect.TypeOf(NotebookTask_SdkV2{}),
		"notification_settings":     reflect.TypeOf(TaskNotificationSettings_SdkV2{}),
		"pipeline_task":             reflect.TypeOf(PipelineTask_SdkV2{}),
		"python_wheel_task":         reflect.TypeOf(PythonWheelTask_SdkV2{}),
		"resolved_values":           reflect.TypeOf(ResolvedValues_SdkV2{}),
		"run_job_task":              reflect.TypeOf(RunJobTask_SdkV2{}),
		"spark_jar_task":            reflect.TypeOf(SparkJarTask_SdkV2{}),
		"spark_python_task":         reflect.TypeOf(SparkPythonTask_SdkV2{}),
		"spark_submit_task":         reflect.TypeOf(SparkSubmitTask_SdkV2{}),
		"sql_task":                  reflect.TypeOf(SqlTask_SdkV2{}),
		"state":                     reflect.TypeOf(RunState_SdkV2{}),
		"status":                    reflect.TypeOf(RunStatus_SdkV2{}),
		"webhook_notifications":     reflect.TypeOf(WebhookNotifications_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunTask_SdkV2
// only implements ToObjectValue() and Type().
func (o RunTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attempt_number":            o.AttemptNumber,
			"clean_rooms_notebook_task": o.CleanRoomsNotebookTask,
			"cleanup_duration":          o.CleanupDuration,
			"cluster_instance":          o.ClusterInstance,
			"condition_task":            o.ConditionTask,
			"dbt_task":                  o.DbtTask,
			"depends_on":                o.DependsOn,
			"description":               o.Description,
			"email_notifications":       o.EmailNotifications,
			"end_time":                  o.EndTime,
			"environment_key":           o.EnvironmentKey,
			"execution_duration":        o.ExecutionDuration,
			"existing_cluster_id":       o.ExistingClusterId,
			"for_each_task":             o.ForEachTask,
			"git_source":                o.GitSource,
			"job_cluster_key":           o.JobClusterKey,
			"library":                   o.Libraries,
			"new_cluster":               o.NewCluster,
			"notebook_task":             o.NotebookTask,
			"notification_settings":     o.NotificationSettings,
			"pipeline_task":             o.PipelineTask,
			"python_wheel_task":         o.PythonWheelTask,
			"queue_duration":            o.QueueDuration,
			"resolved_values":           o.ResolvedValues,
			"run_duration":              o.RunDuration,
			"run_id":                    o.RunId,
			"run_if":                    o.RunIf,
			"run_job_task":              o.RunJobTask,
			"run_page_url":              o.RunPageUrl,
			"setup_duration":            o.SetupDuration,
			"spark_jar_task":            o.SparkJarTask,
			"spark_python_task":         o.SparkPythonTask,
			"spark_submit_task":         o.SparkSubmitTask,
			"sql_task":                  o.SqlTask,
			"start_time":                o.StartTime,
			"state":                     o.State,
			"status":                    o.Status,
			"task_key":                  o.TaskKey,
			"timeout_seconds":           o.TimeoutSeconds,
			"webhook_notifications":     o.WebhookNotifications,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunTask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attempt_number": types.Int64Type,
			"clean_rooms_notebook_task": basetypes.ListType{
				ElemType: CleanRoomsNotebookTask{}.Type(ctx),
			},
			"cleanup_duration": types.Int64Type,
			"cluster_instance": basetypes.ListType{
				ElemType: ClusterInstance{}.Type(ctx),
			},
			"condition_task": basetypes.ListType{
				ElemType: RunConditionTask{}.Type(ctx),
			},
			"dbt_task": basetypes.ListType{
				ElemType: DbtTask{}.Type(ctx),
			},
			"depends_on": basetypes.ListType{
				ElemType: TaskDependency{}.Type(ctx),
			},
			"description": types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: JobEmailNotifications{}.Type(ctx),
			},
			"end_time":            types.Int64Type,
			"environment_key":     types.StringType,
			"execution_duration":  types.Int64Type,
			"existing_cluster_id": types.StringType,
			"for_each_task": basetypes.ListType{
				ElemType: RunForEachTask{}.Type(ctx),
			},
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.Type(ctx),
			},
			"job_cluster_key": types.StringType,
			"library": basetypes.ListType{
				ElemType: compute_tf.Library{}.Type(ctx),
			},
			"new_cluster": basetypes.ListType{
				ElemType: compute_tf.ClusterSpec{}.Type(ctx),
			},
			"notebook_task": basetypes.ListType{
				ElemType: NotebookTask{}.Type(ctx),
			},
			"notification_settings": basetypes.ListType{
				ElemType: TaskNotificationSettings{}.Type(ctx),
			},
			"pipeline_task": basetypes.ListType{
				ElemType: PipelineTask{}.Type(ctx),
			},
			"python_wheel_task": basetypes.ListType{
				ElemType: PythonWheelTask{}.Type(ctx),
			},
			"queue_duration": types.Int64Type,
			"resolved_values": basetypes.ListType{
				ElemType: ResolvedValues{}.Type(ctx),
			},
			"run_duration": types.Int64Type,
			"run_id":       types.Int64Type,
			"run_if":       types.StringType,
			"run_job_task": basetypes.ListType{
				ElemType: RunJobTask{}.Type(ctx),
			},
			"run_page_url":   types.StringType,
			"setup_duration": types.Int64Type,
			"spark_jar_task": basetypes.ListType{
				ElemType: SparkJarTask{}.Type(ctx),
			},
			"spark_python_task": basetypes.ListType{
				ElemType: SparkPythonTask{}.Type(ctx),
			},
			"spark_submit_task": basetypes.ListType{
				ElemType: SparkSubmitTask{}.Type(ctx),
			},
			"sql_task": basetypes.ListType{
				ElemType: SqlTask{}.Type(ctx),
			},
			"start_time": types.Int64Type,
			"state": basetypes.ListType{
				ElemType: RunState{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: RunStatus{}.Type(ctx),
			},
			"task_key":        types.StringType,
			"timeout_seconds": types.Int64Type,
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.Type(ctx),
			},
		},
	}
}

// GetCleanRoomsNotebookTask returns the value of the CleanRoomsNotebookTask field in RunTask_SdkV2 as
// a CleanRoomsNotebookTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetCleanRoomsNotebookTask(ctx context.Context) (CleanRoomsNotebookTask_SdkV2, bool) {
	var e CleanRoomsNotebookTask_SdkV2
	if o.CleanRoomsNotebookTask.IsNull() || o.CleanRoomsNotebookTask.IsUnknown() {
		return e, false
	}
	var v []CleanRoomsNotebookTask_SdkV2
	d := o.CleanRoomsNotebookTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCleanRoomsNotebookTask sets the value of the CleanRoomsNotebookTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetCleanRoomsNotebookTask(ctx context.Context, v CleanRoomsNotebookTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_rooms_notebook_task"]
	o.CleanRoomsNotebookTask = types.ListValueMust(t, vs)
}

// GetClusterInstance returns the value of the ClusterInstance field in RunTask_SdkV2 as
// a ClusterInstance_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetClusterInstance(ctx context.Context) (ClusterInstance_SdkV2, bool) {
	var e ClusterInstance_SdkV2
	if o.ClusterInstance.IsNull() || o.ClusterInstance.IsUnknown() {
		return e, false
	}
	var v []ClusterInstance_SdkV2
	d := o.ClusterInstance.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetClusterInstance sets the value of the ClusterInstance field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetClusterInstance(ctx context.Context, v ClusterInstance_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cluster_instance"]
	o.ClusterInstance = types.ListValueMust(t, vs)
}

// GetConditionTask returns the value of the ConditionTask field in RunTask_SdkV2 as
// a RunConditionTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetConditionTask(ctx context.Context) (RunConditionTask_SdkV2, bool) {
	var e RunConditionTask_SdkV2
	if o.ConditionTask.IsNull() || o.ConditionTask.IsUnknown() {
		return e, false
	}
	var v []RunConditionTask_SdkV2
	d := o.ConditionTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConditionTask sets the value of the ConditionTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetConditionTask(ctx context.Context, v RunConditionTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition_task"]
	o.ConditionTask = types.ListValueMust(t, vs)
}

// GetDbtTask returns the value of the DbtTask field in RunTask_SdkV2 as
// a DbtTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetDbtTask(ctx context.Context) (DbtTask_SdkV2, bool) {
	var e DbtTask_SdkV2
	if o.DbtTask.IsNull() || o.DbtTask.IsUnknown() {
		return e, false
	}
	var v []DbtTask_SdkV2
	d := o.DbtTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbtTask sets the value of the DbtTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetDbtTask(ctx context.Context, v DbtTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbt_task"]
	o.DbtTask = types.ListValueMust(t, vs)
}

// GetDependsOn returns the value of the DependsOn field in RunTask_SdkV2 as
// a slice of TaskDependency_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetDependsOn(ctx context.Context) ([]TaskDependency_SdkV2, bool) {
	if o.DependsOn.IsNull() || o.DependsOn.IsUnknown() {
		return nil, false
	}
	var v []TaskDependency_SdkV2
	d := o.DependsOn.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependsOn sets the value of the DependsOn field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetDependsOn(ctx context.Context, v []TaskDependency_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["depends_on"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DependsOn = types.ListValueMust(t, vs)
}

// GetEmailNotifications returns the value of the EmailNotifications field in RunTask_SdkV2 as
// a JobEmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetEmailNotifications(ctx context.Context) (JobEmailNotifications_SdkV2, bool) {
	var e JobEmailNotifications_SdkV2
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []JobEmailNotifications_SdkV2
	d := o.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetEmailNotifications(ctx context.Context, v JobEmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	o.EmailNotifications = types.ListValueMust(t, vs)
}

// GetForEachTask returns the value of the ForEachTask field in RunTask_SdkV2 as
// a RunForEachTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetForEachTask(ctx context.Context) (RunForEachTask_SdkV2, bool) {
	var e RunForEachTask_SdkV2
	if o.ForEachTask.IsNull() || o.ForEachTask.IsUnknown() {
		return e, false
	}
	var v []RunForEachTask_SdkV2
	d := o.ForEachTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetForEachTask sets the value of the ForEachTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetForEachTask(ctx context.Context, v RunForEachTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["for_each_task"]
	o.ForEachTask = types.ListValueMust(t, vs)
}

// GetGitSource returns the value of the GitSource field in RunTask_SdkV2 as
// a GitSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetGitSource(ctx context.Context) (GitSource_SdkV2, bool) {
	var e GitSource_SdkV2
	if o.GitSource.IsNull() || o.GitSource.IsUnknown() {
		return e, false
	}
	var v []GitSource_SdkV2
	d := o.GitSource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitSource sets the value of the GitSource field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetGitSource(ctx context.Context, v GitSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["git_source"]
	o.GitSource = types.ListValueMust(t, vs)
}

// GetLibraries returns the value of the Libraries field in RunTask_SdkV2 as
// a slice of compute_tf.Library_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetLibraries(ctx context.Context) ([]compute_tf.Library_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []compute_tf.Library_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetLibraries(ctx context.Context, v []compute_tf.Library_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["library"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

// GetNewCluster returns the value of the NewCluster field in RunTask_SdkV2 as
// a compute_tf.ClusterSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetNewCluster(ctx context.Context) (compute_tf.ClusterSpec_SdkV2, bool) {
	var e compute_tf.ClusterSpec_SdkV2
	if o.NewCluster.IsNull() || o.NewCluster.IsUnknown() {
		return e, false
	}
	var v []compute_tf.ClusterSpec_SdkV2
	d := o.NewCluster.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNewCluster sets the value of the NewCluster field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetNewCluster(ctx context.Context, v compute_tf.ClusterSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["new_cluster"]
	o.NewCluster = types.ListValueMust(t, vs)
}

// GetNotebookTask returns the value of the NotebookTask field in RunTask_SdkV2 as
// a NotebookTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetNotebookTask(ctx context.Context) (NotebookTask_SdkV2, bool) {
	var e NotebookTask_SdkV2
	if o.NotebookTask.IsNull() || o.NotebookTask.IsUnknown() {
		return e, false
	}
	var v []NotebookTask_SdkV2
	d := o.NotebookTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebookTask sets the value of the NotebookTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetNotebookTask(ctx context.Context, v NotebookTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_task"]
	o.NotebookTask = types.ListValueMust(t, vs)
}

// GetNotificationSettings returns the value of the NotificationSettings field in RunTask_SdkV2 as
// a TaskNotificationSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetNotificationSettings(ctx context.Context) (TaskNotificationSettings_SdkV2, bool) {
	var e TaskNotificationSettings_SdkV2
	if o.NotificationSettings.IsNull() || o.NotificationSettings.IsUnknown() {
		return e, false
	}
	var v []TaskNotificationSettings_SdkV2
	d := o.NotificationSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotificationSettings sets the value of the NotificationSettings field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetNotificationSettings(ctx context.Context, v TaskNotificationSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notification_settings"]
	o.NotificationSettings = types.ListValueMust(t, vs)
}

// GetPipelineTask returns the value of the PipelineTask field in RunTask_SdkV2 as
// a PipelineTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetPipelineTask(ctx context.Context) (PipelineTask_SdkV2, bool) {
	var e PipelineTask_SdkV2
	if o.PipelineTask.IsNull() || o.PipelineTask.IsUnknown() {
		return e, false
	}
	var v []PipelineTask_SdkV2
	d := o.PipelineTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPipelineTask sets the value of the PipelineTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetPipelineTask(ctx context.Context, v PipelineTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pipeline_task"]
	o.PipelineTask = types.ListValueMust(t, vs)
}

// GetPythonWheelTask returns the value of the PythonWheelTask field in RunTask_SdkV2 as
// a PythonWheelTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetPythonWheelTask(ctx context.Context) (PythonWheelTask_SdkV2, bool) {
	var e PythonWheelTask_SdkV2
	if o.PythonWheelTask.IsNull() || o.PythonWheelTask.IsUnknown() {
		return e, false
	}
	var v []PythonWheelTask_SdkV2
	d := o.PythonWheelTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPythonWheelTask sets the value of the PythonWheelTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetPythonWheelTask(ctx context.Context, v PythonWheelTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_wheel_task"]
	o.PythonWheelTask = types.ListValueMust(t, vs)
}

// GetResolvedValues returns the value of the ResolvedValues field in RunTask_SdkV2 as
// a ResolvedValues_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetResolvedValues(ctx context.Context) (ResolvedValues_SdkV2, bool) {
	var e ResolvedValues_SdkV2
	if o.ResolvedValues.IsNull() || o.ResolvedValues.IsUnknown() {
		return e, false
	}
	var v []ResolvedValues_SdkV2
	d := o.ResolvedValues.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResolvedValues sets the value of the ResolvedValues field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetResolvedValues(ctx context.Context, v ResolvedValues_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resolved_values"]
	o.ResolvedValues = types.ListValueMust(t, vs)
}

// GetRunJobTask returns the value of the RunJobTask field in RunTask_SdkV2 as
// a RunJobTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetRunJobTask(ctx context.Context) (RunJobTask_SdkV2, bool) {
	var e RunJobTask_SdkV2
	if o.RunJobTask.IsNull() || o.RunJobTask.IsUnknown() {
		return e, false
	}
	var v []RunJobTask_SdkV2
	d := o.RunJobTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunJobTask sets the value of the RunJobTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetRunJobTask(ctx context.Context, v RunJobTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_job_task"]
	o.RunJobTask = types.ListValueMust(t, vs)
}

// GetSparkJarTask returns the value of the SparkJarTask field in RunTask_SdkV2 as
// a SparkJarTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetSparkJarTask(ctx context.Context) (SparkJarTask_SdkV2, bool) {
	var e SparkJarTask_SdkV2
	if o.SparkJarTask.IsNull() || o.SparkJarTask.IsUnknown() {
		return e, false
	}
	var v []SparkJarTask_SdkV2
	d := o.SparkJarTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkJarTask sets the value of the SparkJarTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetSparkJarTask(ctx context.Context, v SparkJarTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_jar_task"]
	o.SparkJarTask = types.ListValueMust(t, vs)
}

// GetSparkPythonTask returns the value of the SparkPythonTask field in RunTask_SdkV2 as
// a SparkPythonTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetSparkPythonTask(ctx context.Context) (SparkPythonTask_SdkV2, bool) {
	var e SparkPythonTask_SdkV2
	if o.SparkPythonTask.IsNull() || o.SparkPythonTask.IsUnknown() {
		return e, false
	}
	var v []SparkPythonTask_SdkV2
	d := o.SparkPythonTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkPythonTask sets the value of the SparkPythonTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetSparkPythonTask(ctx context.Context, v SparkPythonTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_python_task"]
	o.SparkPythonTask = types.ListValueMust(t, vs)
}

// GetSparkSubmitTask returns the value of the SparkSubmitTask field in RunTask_SdkV2 as
// a SparkSubmitTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetSparkSubmitTask(ctx context.Context) (SparkSubmitTask_SdkV2, bool) {
	var e SparkSubmitTask_SdkV2
	if o.SparkSubmitTask.IsNull() || o.SparkSubmitTask.IsUnknown() {
		return e, false
	}
	var v []SparkSubmitTask_SdkV2
	d := o.SparkSubmitTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkSubmitTask sets the value of the SparkSubmitTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetSparkSubmitTask(ctx context.Context, v SparkSubmitTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_submit_task"]
	o.SparkSubmitTask = types.ListValueMust(t, vs)
}

// GetSqlTask returns the value of the SqlTask field in RunTask_SdkV2 as
// a SqlTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetSqlTask(ctx context.Context) (SqlTask_SdkV2, bool) {
	var e SqlTask_SdkV2
	if o.SqlTask.IsNull() || o.SqlTask.IsUnknown() {
		return e, false
	}
	var v []SqlTask_SdkV2
	d := o.SqlTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlTask sets the value of the SqlTask field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetSqlTask(ctx context.Context, v SqlTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_task"]
	o.SqlTask = types.ListValueMust(t, vs)
}

// GetState returns the value of the State field in RunTask_SdkV2 as
// a RunState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetState(ctx context.Context) (RunState_SdkV2, bool) {
	var e RunState_SdkV2
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v []RunState_SdkV2
	d := o.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetState(ctx context.Context, v RunState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	o.State = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in RunTask_SdkV2 as
// a RunStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetStatus(ctx context.Context) (RunStatus_SdkV2, bool) {
	var e RunStatus_SdkV2
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []RunStatus_SdkV2
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetStatus(ctx context.Context, v RunStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
}

// GetWebhookNotifications returns the value of the WebhookNotifications field in RunTask_SdkV2 as
// a WebhookNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RunTask_SdkV2) GetWebhookNotifications(ctx context.Context) (WebhookNotifications_SdkV2, bool) {
	var e WebhookNotifications_SdkV2
	if o.WebhookNotifications.IsNull() || o.WebhookNotifications.IsUnknown() {
		return e, false
	}
	var v []WebhookNotifications_SdkV2
	d := o.WebhookNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhookNotifications sets the value of the WebhookNotifications field in RunTask_SdkV2.
func (o *RunTask_SdkV2) SetWebhookNotifications(ctx context.Context, v WebhookNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook_notifications"]
	o.WebhookNotifications = types.ListValueMust(t, vs)
}

type SparkJarTask_SdkV2 struct {
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

func (newState *SparkJarTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkJarTask_SdkV2) {
}

func (newState *SparkJarTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState SparkJarTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkJarTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SparkJarTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkJarTask_SdkV2
// only implements ToObjectValue() and Type().
func (o SparkJarTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"jar_uri":         o.JarUri,
			"main_class_name": o.MainClassName,
			"parameters":      o.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparkJarTask_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetParameters returns the value of the Parameters field in SparkJarTask_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SparkJarTask_SdkV2) GetParameters(ctx context.Context) ([]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in SparkJarTask_SdkV2.
func (o *SparkJarTask_SdkV2) SetParameters(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

type SparkPythonTask_SdkV2 struct {
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

func (newState *SparkPythonTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkPythonTask_SdkV2) {
}

func (newState *SparkPythonTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState SparkPythonTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkPythonTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SparkPythonTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkPythonTask_SdkV2
// only implements ToObjectValue() and Type().
func (o SparkPythonTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parameters":  o.Parameters,
			"python_file": o.PythonFile,
			"source":      o.Source,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparkPythonTask_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetParameters returns the value of the Parameters field in SparkPythonTask_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SparkPythonTask_SdkV2) GetParameters(ctx context.Context) ([]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in SparkPythonTask_SdkV2.
func (o *SparkPythonTask_SdkV2) SetParameters(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

type SparkSubmitTask_SdkV2 struct {
	// Command-line parameters passed to spark submit.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
}

func (newState *SparkSubmitTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparkSubmitTask_SdkV2) {
}

func (newState *SparkSubmitTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState SparkSubmitTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparkSubmitTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SparkSubmitTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparkSubmitTask_SdkV2
// only implements ToObjectValue() and Type().
func (o SparkSubmitTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parameters": o.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparkSubmitTask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parameters": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetParameters returns the value of the Parameters field in SparkSubmitTask_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SparkSubmitTask_SdkV2) GetParameters(ctx context.Context) ([]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in SparkSubmitTask_SdkV2.
func (o *SparkSubmitTask_SdkV2) SetParameters(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
}

type SqlAlertOutput_SdkV2 struct {
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

func (newState *SqlAlertOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlAlertOutput_SdkV2) {
}

func (newState *SqlAlertOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlAlertOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlAlertOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlAlertOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sql_statements": reflect.TypeOf(SqlStatementOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlAlertOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlAlertOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_state":    o.AlertState,
			"output_link":    o.OutputLink,
			"query_text":     o.QueryText,
			"sql_statements": o.SqlStatements,
			"warehouse_id":   o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlAlertOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_state": types.StringType,
			"output_link": types.StringType,
			"query_text":  types.StringType,
			"sql_statements": basetypes.ListType{
				ElemType: SqlStatementOutput{}.Type(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

// GetSqlStatements returns the value of the SqlStatements field in SqlAlertOutput_SdkV2 as
// a slice of SqlStatementOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlAlertOutput_SdkV2) GetSqlStatements(ctx context.Context) ([]SqlStatementOutput_SdkV2, bool) {
	if o.SqlStatements.IsNull() || o.SqlStatements.IsUnknown() {
		return nil, false
	}
	var v []SqlStatementOutput_SdkV2
	d := o.SqlStatements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlStatements sets the value of the SqlStatements field in SqlAlertOutput_SdkV2.
func (o *SqlAlertOutput_SdkV2) SetSqlStatements(ctx context.Context, v []SqlStatementOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_statements"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SqlStatements = types.ListValueMust(t, vs)
}

type SqlDashboardOutput_SdkV2 struct {
	// The canonical identifier of the SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
	// Widgets executed in the run. Only SQL query based widgets are listed.
	Widgets types.List `tfsdk:"widgets" tf:"optional"`
}

func (newState *SqlDashboardOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlDashboardOutput_SdkV2) {
}

func (newState *SqlDashboardOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlDashboardOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlDashboardOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlDashboardOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"widgets": reflect.TypeOf(SqlDashboardWidgetOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlDashboardOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlDashboardOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warehouse_id": o.WarehouseId,
			"widgets":      o.Widgets,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlDashboardOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warehouse_id": types.StringType,
			"widgets": basetypes.ListType{
				ElemType: SqlDashboardWidgetOutput{}.Type(ctx),
			},
		},
	}
}

// GetWidgets returns the value of the Widgets field in SqlDashboardOutput_SdkV2 as
// a slice of SqlDashboardWidgetOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlDashboardOutput_SdkV2) GetWidgets(ctx context.Context) ([]SqlDashboardWidgetOutput_SdkV2, bool) {
	if o.Widgets.IsNull() || o.Widgets.IsUnknown() {
		return nil, false
	}
	var v []SqlDashboardWidgetOutput_SdkV2
	d := o.Widgets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWidgets sets the value of the Widgets field in SqlDashboardOutput_SdkV2.
func (o *SqlDashboardOutput_SdkV2) SetWidgets(ctx context.Context, v []SqlDashboardWidgetOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["widgets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Widgets = types.ListValueMust(t, vs)
}

type SqlDashboardWidgetOutput_SdkV2 struct {
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

func (newState *SqlDashboardWidgetOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlDashboardWidgetOutput_SdkV2) {
}

func (newState *SqlDashboardWidgetOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlDashboardWidgetOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlDashboardWidgetOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlDashboardWidgetOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(SqlOutputError_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlDashboardWidgetOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlDashboardWidgetOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time":     o.EndTime,
			"error":        o.Error,
			"output_link":  o.OutputLink,
			"start_time":   o.StartTime,
			"status":       o.Status,
			"widget_id":    o.WidgetId,
			"widget_title": o.WidgetTitle,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlDashboardWidgetOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time": types.Int64Type,
			"error": basetypes.ListType{
				ElemType: SqlOutputError{}.Type(ctx),
			},
			"output_link":  types.StringType,
			"start_time":   types.Int64Type,
			"status":       types.StringType,
			"widget_id":    types.StringType,
			"widget_title": types.StringType,
		},
	}
}

// GetError returns the value of the Error field in SqlDashboardWidgetOutput_SdkV2 as
// a SqlOutputError_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlDashboardWidgetOutput_SdkV2) GetError(ctx context.Context) (SqlOutputError_SdkV2, bool) {
	var e SqlOutputError_SdkV2
	if o.Error.IsNull() || o.Error.IsUnknown() {
		return e, false
	}
	var v []SqlOutputError_SdkV2
	d := o.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in SqlDashboardWidgetOutput_SdkV2.
func (o *SqlDashboardWidgetOutput_SdkV2) SetError(ctx context.Context, v SqlOutputError_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	o.Error = types.ListValueMust(t, vs)
}

type SqlOutput_SdkV2 struct {
	// The output of a SQL alert task, if available.
	AlertOutput types.List `tfsdk:"alert_output" tf:"optional,object"`
	// The output of a SQL dashboard task, if available.
	DashboardOutput types.List `tfsdk:"dashboard_output" tf:"optional,object"`
	// The output of a SQL query task, if available.
	QueryOutput types.List `tfsdk:"query_output" tf:"optional,object"`
}

func (newState *SqlOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlOutput_SdkV2) {
}

func (newState *SqlOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_output":     reflect.TypeOf(SqlAlertOutput_SdkV2{}),
		"dashboard_output": reflect.TypeOf(SqlDashboardOutput_SdkV2{}),
		"query_output":     reflect.TypeOf(SqlQueryOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_output":     o.AlertOutput,
			"dashboard_output": o.DashboardOutput,
			"query_output":     o.QueryOutput,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_output": basetypes.ListType{
				ElemType: SqlAlertOutput{}.Type(ctx),
			},
			"dashboard_output": basetypes.ListType{
				ElemType: SqlDashboardOutput{}.Type(ctx),
			},
			"query_output": basetypes.ListType{
				ElemType: SqlQueryOutput{}.Type(ctx),
			},
		},
	}
}

// GetAlertOutput returns the value of the AlertOutput field in SqlOutput_SdkV2 as
// a SqlAlertOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlOutput_SdkV2) GetAlertOutput(ctx context.Context) (SqlAlertOutput_SdkV2, bool) {
	var e SqlAlertOutput_SdkV2
	if o.AlertOutput.IsNull() || o.AlertOutput.IsUnknown() {
		return e, false
	}
	var v []SqlAlertOutput_SdkV2
	d := o.AlertOutput.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlertOutput sets the value of the AlertOutput field in SqlOutput_SdkV2.
func (o *SqlOutput_SdkV2) SetAlertOutput(ctx context.Context, v SqlAlertOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert_output"]
	o.AlertOutput = types.ListValueMust(t, vs)
}

// GetDashboardOutput returns the value of the DashboardOutput field in SqlOutput_SdkV2 as
// a SqlDashboardOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlOutput_SdkV2) GetDashboardOutput(ctx context.Context) (SqlDashboardOutput_SdkV2, bool) {
	var e SqlDashboardOutput_SdkV2
	if o.DashboardOutput.IsNull() || o.DashboardOutput.IsUnknown() {
		return e, false
	}
	var v []SqlDashboardOutput_SdkV2
	d := o.DashboardOutput.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDashboardOutput sets the value of the DashboardOutput field in SqlOutput_SdkV2.
func (o *SqlOutput_SdkV2) SetDashboardOutput(ctx context.Context, v SqlDashboardOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dashboard_output"]
	o.DashboardOutput = types.ListValueMust(t, vs)
}

// GetQueryOutput returns the value of the QueryOutput field in SqlOutput_SdkV2 as
// a SqlQueryOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlOutput_SdkV2) GetQueryOutput(ctx context.Context) (SqlQueryOutput_SdkV2, bool) {
	var e SqlQueryOutput_SdkV2
	if o.QueryOutput.IsNull() || o.QueryOutput.IsUnknown() {
		return e, false
	}
	var v []SqlQueryOutput_SdkV2
	d := o.QueryOutput.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryOutput sets the value of the QueryOutput field in SqlOutput_SdkV2.
func (o *SqlOutput_SdkV2) SetQueryOutput(ctx context.Context, v SqlQueryOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_output"]
	o.QueryOutput = types.ListValueMust(t, vs)
}

type SqlOutputError_SdkV2 struct {
	// The error message when execution fails.
	Message types.String `tfsdk:"message" tf:"optional"`
}

func (newState *SqlOutputError_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlOutputError_SdkV2) {
}

func (newState *SqlOutputError_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlOutputError_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlOutputError.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlOutputError_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlOutputError_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlOutputError_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlOutputError_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
		},
	}
}

type SqlQueryOutput_SdkV2 struct {
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

func (newState *SqlQueryOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlQueryOutput_SdkV2) {
}

func (newState *SqlQueryOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlQueryOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlQueryOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlQueryOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sql_statements": reflect.TypeOf(SqlStatementOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlQueryOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlQueryOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_id":    o.EndpointId,
			"output_link":    o.OutputLink,
			"query_text":     o.QueryText,
			"sql_statements": o.SqlStatements,
			"warehouse_id":   o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlQueryOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_id": types.StringType,
			"output_link": types.StringType,
			"query_text":  types.StringType,
			"sql_statements": basetypes.ListType{
				ElemType: SqlStatementOutput{}.Type(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

// GetSqlStatements returns the value of the SqlStatements field in SqlQueryOutput_SdkV2 as
// a slice of SqlStatementOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlQueryOutput_SdkV2) GetSqlStatements(ctx context.Context) ([]SqlStatementOutput_SdkV2, bool) {
	if o.SqlStatements.IsNull() || o.SqlStatements.IsUnknown() {
		return nil, false
	}
	var v []SqlStatementOutput_SdkV2
	d := o.SqlStatements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlStatements sets the value of the SqlStatements field in SqlQueryOutput_SdkV2.
func (o *SqlQueryOutput_SdkV2) SetSqlStatements(ctx context.Context, v []SqlStatementOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_statements"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SqlStatements = types.ListValueMust(t, vs)
}

type SqlStatementOutput_SdkV2 struct {
	// A key that can be used to look up query details.
	LookupKey types.String `tfsdk:"lookup_key" tf:"optional"`
}

func (newState *SqlStatementOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlStatementOutput_SdkV2) {
}

func (newState *SqlStatementOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlStatementOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlStatementOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlStatementOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlStatementOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlStatementOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"lookup_key": o.LookupKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlStatementOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"lookup_key": types.StringType,
		},
	}
}

type SqlTask_SdkV2 struct {
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

func (newState *SqlTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTask_SdkV2) {
}

func (newState *SqlTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert":      reflect.TypeOf(SqlTaskAlert_SdkV2{}),
		"dashboard":  reflect.TypeOf(SqlTaskDashboard_SdkV2{}),
		"file":       reflect.TypeOf(SqlTaskFile_SdkV2{}),
		"parameters": reflect.TypeOf(types.String{}),
		"query":      reflect.TypeOf(SqlTaskQuery_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlTask_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert":        o.Alert,
			"dashboard":    o.Dashboard,
			"file":         o.File,
			"parameters":   o.Parameters,
			"query":        o.Query,
			"warehouse_id": o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlTask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert": basetypes.ListType{
				ElemType: SqlTaskAlert{}.Type(ctx),
			},
			"dashboard": basetypes.ListType{
				ElemType: SqlTaskDashboard{}.Type(ctx),
			},
			"file": basetypes.ListType{
				ElemType: SqlTaskFile{}.Type(ctx),
			},
			"parameters": basetypes.MapType{
				ElemType: types.StringType,
			},
			"query": basetypes.ListType{
				ElemType: SqlTaskQuery{}.Type(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

// GetAlert returns the value of the Alert field in SqlTask_SdkV2 as
// a SqlTaskAlert_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlTask_SdkV2) GetAlert(ctx context.Context) (SqlTaskAlert_SdkV2, bool) {
	var e SqlTaskAlert_SdkV2
	if o.Alert.IsNull() || o.Alert.IsUnknown() {
		return e, false
	}
	var v []SqlTaskAlert_SdkV2
	d := o.Alert.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAlert sets the value of the Alert field in SqlTask_SdkV2.
func (o *SqlTask_SdkV2) SetAlert(ctx context.Context, v SqlTaskAlert_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["alert"]
	o.Alert = types.ListValueMust(t, vs)
}

// GetDashboard returns the value of the Dashboard field in SqlTask_SdkV2 as
// a SqlTaskDashboard_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlTask_SdkV2) GetDashboard(ctx context.Context) (SqlTaskDashboard_SdkV2, bool) {
	var e SqlTaskDashboard_SdkV2
	if o.Dashboard.IsNull() || o.Dashboard.IsUnknown() {
		return e, false
	}
	var v []SqlTaskDashboard_SdkV2
	d := o.Dashboard.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDashboard sets the value of the Dashboard field in SqlTask_SdkV2.
func (o *SqlTask_SdkV2) SetDashboard(ctx context.Context, v SqlTaskDashboard_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dashboard"]
	o.Dashboard = types.ListValueMust(t, vs)
}

// GetFile returns the value of the File field in SqlTask_SdkV2 as
// a SqlTaskFile_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlTask_SdkV2) GetFile(ctx context.Context) (SqlTaskFile_SdkV2, bool) {
	var e SqlTaskFile_SdkV2
	if o.File.IsNull() || o.File.IsUnknown() {
		return e, false
	}
	var v []SqlTaskFile_SdkV2
	d := o.File.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFile sets the value of the File field in SqlTask_SdkV2.
func (o *SqlTask_SdkV2) SetFile(ctx context.Context, v SqlTaskFile_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file"]
	o.File = types.ListValueMust(t, vs)
}

// GetParameters returns the value of the Parameters field in SqlTask_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlTask_SdkV2) GetParameters(ctx context.Context) (map[string]types.String, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in SqlTask_SdkV2.
func (o *SqlTask_SdkV2) SetParameters(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.MapValueMust(t, vs)
}

// GetQuery returns the value of the Query field in SqlTask_SdkV2 as
// a SqlTaskQuery_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlTask_SdkV2) GetQuery(ctx context.Context) (SqlTaskQuery_SdkV2, bool) {
	var e SqlTaskQuery_SdkV2
	if o.Query.IsNull() || o.Query.IsUnknown() {
		return e, false
	}
	var v []SqlTaskQuery_SdkV2
	d := o.Query.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuery sets the value of the Query field in SqlTask_SdkV2.
func (o *SqlTask_SdkV2) SetQuery(ctx context.Context, v SqlTaskQuery_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query"]
	o.Query = types.ListValueMust(t, vs)
}

type SqlTaskAlert_SdkV2 struct {
	// The canonical identifier of the SQL alert.
	AlertId types.String `tfsdk:"alert_id" tf:""`
	// If true, the alert notifications are not sent to subscribers.
	PauseSubscriptions types.Bool `tfsdk:"pause_subscriptions" tf:"optional"`
	// If specified, alert notifications are sent to subscribers.
	Subscriptions types.List `tfsdk:"subscriptions" tf:"optional"`
}

func (newState *SqlTaskAlert_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTaskAlert_SdkV2) {
}

func (newState *SqlTaskAlert_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlTaskAlert_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTaskAlert.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTaskAlert_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriptions": reflect.TypeOf(SqlTaskSubscription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlTaskAlert_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlTaskAlert_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_id":            o.AlertId,
			"pause_subscriptions": o.PauseSubscriptions,
			"subscriptions":       o.Subscriptions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlTaskAlert_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_id":            types.StringType,
			"pause_subscriptions": types.BoolType,
			"subscriptions": basetypes.ListType{
				ElemType: SqlTaskSubscription{}.Type(ctx),
			},
		},
	}
}

// GetSubscriptions returns the value of the Subscriptions field in SqlTaskAlert_SdkV2 as
// a slice of SqlTaskSubscription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlTaskAlert_SdkV2) GetSubscriptions(ctx context.Context) ([]SqlTaskSubscription_SdkV2, bool) {
	if o.Subscriptions.IsNull() || o.Subscriptions.IsUnknown() {
		return nil, false
	}
	var v []SqlTaskSubscription_SdkV2
	d := o.Subscriptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubscriptions sets the value of the Subscriptions field in SqlTaskAlert_SdkV2.
func (o *SqlTaskAlert_SdkV2) SetSubscriptions(ctx context.Context, v []SqlTaskSubscription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subscriptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Subscriptions = types.ListValueMust(t, vs)
}

type SqlTaskDashboard_SdkV2 struct {
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

func (newState *SqlTaskDashboard_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTaskDashboard_SdkV2) {
}

func (newState *SqlTaskDashboard_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlTaskDashboard_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTaskDashboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTaskDashboard_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriptions": reflect.TypeOf(SqlTaskSubscription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlTaskDashboard_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlTaskDashboard_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_subject":      o.CustomSubject,
			"dashboard_id":        o.DashboardId,
			"pause_subscriptions": o.PauseSubscriptions,
			"subscriptions":       o.Subscriptions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlTaskDashboard_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_subject":      types.StringType,
			"dashboard_id":        types.StringType,
			"pause_subscriptions": types.BoolType,
			"subscriptions": basetypes.ListType{
				ElemType: SqlTaskSubscription{}.Type(ctx),
			},
		},
	}
}

// GetSubscriptions returns the value of the Subscriptions field in SqlTaskDashboard_SdkV2 as
// a slice of SqlTaskSubscription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SqlTaskDashboard_SdkV2) GetSubscriptions(ctx context.Context) ([]SqlTaskSubscription_SdkV2, bool) {
	if o.Subscriptions.IsNull() || o.Subscriptions.IsUnknown() {
		return nil, false
	}
	var v []SqlTaskSubscription_SdkV2
	d := o.Subscriptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubscriptions sets the value of the Subscriptions field in SqlTaskDashboard_SdkV2.
func (o *SqlTaskDashboard_SdkV2) SetSubscriptions(ctx context.Context, v []SqlTaskSubscription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subscriptions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Subscriptions = types.ListValueMust(t, vs)
}

type SqlTaskFile_SdkV2 struct {
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

func (newState *SqlTaskFile_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTaskFile_SdkV2) {
}

func (newState *SqlTaskFile_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlTaskFile_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTaskFile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTaskFile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlTaskFile_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlTaskFile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path":   o.Path,
			"source": o.Source,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlTaskFile_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":   types.StringType,
			"source": types.StringType,
		},
	}
}

type SqlTaskQuery_SdkV2 struct {
	// The canonical identifier of the SQL query.
	QueryId types.String `tfsdk:"query_id" tf:""`
}

func (newState *SqlTaskQuery_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTaskQuery_SdkV2) {
}

func (newState *SqlTaskQuery_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlTaskQuery_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTaskQuery.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTaskQuery_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlTaskQuery_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlTaskQuery_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_id": o.QueryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlTaskQuery_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_id": types.StringType,
		},
	}
}

type SqlTaskSubscription_SdkV2 struct {
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

func (newState *SqlTaskSubscription_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlTaskSubscription_SdkV2) {
}

func (newState *SqlTaskSubscription_SdkV2) SyncEffectiveFieldsDuringRead(existingState SqlTaskSubscription_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlTaskSubscription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SqlTaskSubscription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlTaskSubscription_SdkV2
// only implements ToObjectValue() and Type().
func (o SqlTaskSubscription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_id": o.DestinationId,
			"user_name":      o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlTaskSubscription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_id": types.StringType,
			"user_name":      types.StringType,
		},
	}
}

type SubmitRun_SdkV2 struct {
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

func (newState *SubmitRun_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubmitRun_SdkV2) {
}

func (newState *SubmitRun_SdkV2) SyncEffectiveFieldsDuringRead(existingState SubmitRun_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubmitRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubmitRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list":   reflect.TypeOf(JobAccessControlRequest_SdkV2{}),
		"email_notifications":   reflect.TypeOf(JobEmailNotifications_SdkV2{}),
		"environments":          reflect.TypeOf(JobEnvironment_SdkV2{}),
		"git_source":            reflect.TypeOf(GitSource_SdkV2{}),
		"health":                reflect.TypeOf(JobsHealthRules_SdkV2{}),
		"notification_settings": reflect.TypeOf(JobNotificationSettings_SdkV2{}),
		"queue":                 reflect.TypeOf(QueueSettings_SdkV2{}),
		"run_as":                reflect.TypeOf(JobRunAs_SdkV2{}),
		"tasks":                 reflect.TypeOf(SubmitTask_SdkV2{}),
		"webhook_notifications": reflect.TypeOf(WebhookNotifications_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SubmitRun_SdkV2
// only implements ToObjectValue() and Type().
func (o SubmitRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list":   o.AccessControlList,
			"budget_policy_id":      o.BudgetPolicyId,
			"email_notifications":   o.EmailNotifications,
			"environments":          o.Environments,
			"git_source":            o.GitSource,
			"health":                o.Health,
			"idempotency_token":     o.IdempotencyToken,
			"notification_settings": o.NotificationSettings,
			"queue":                 o.Queue,
			"run_as":                o.RunAs,
			"run_name":              o.RunName,
			"tasks":                 o.Tasks,
			"timeout_seconds":       o.TimeoutSeconds,
			"webhook_notifications": o.WebhookNotifications,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SubmitRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: JobAccessControlRequest{}.Type(ctx),
			},
			"budget_policy_id": types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: JobEmailNotifications{}.Type(ctx),
			},
			"environments": basetypes.ListType{
				ElemType: JobEnvironment{}.Type(ctx),
			},
			"git_source": basetypes.ListType{
				ElemType: GitSource{}.Type(ctx),
			},
			"health": basetypes.ListType{
				ElemType: JobsHealthRules{}.Type(ctx),
			},
			"idempotency_token": types.StringType,
			"notification_settings": basetypes.ListType{
				ElemType: JobNotificationSettings{}.Type(ctx),
			},
			"queue": basetypes.ListType{
				ElemType: QueueSettings{}.Type(ctx),
			},
			"run_as": basetypes.ListType{
				ElemType: JobRunAs{}.Type(ctx),
			},
			"run_name": types.StringType,
			"tasks": basetypes.ListType{
				ElemType: SubmitTask{}.Type(ctx),
			},
			"timeout_seconds": types.Int64Type,
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.Type(ctx),
			},
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in SubmitRun_SdkV2 as
// a slice of JobAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitRun_SdkV2) GetAccessControlList(ctx context.Context) ([]JobAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []JobAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SubmitRun_SdkV2.
func (o *SubmitRun_SdkV2) SetAccessControlList(ctx context.Context, v []JobAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

// GetEmailNotifications returns the value of the EmailNotifications field in SubmitRun_SdkV2 as
// a JobEmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitRun_SdkV2) GetEmailNotifications(ctx context.Context) (JobEmailNotifications_SdkV2, bool) {
	var e JobEmailNotifications_SdkV2
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []JobEmailNotifications_SdkV2
	d := o.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in SubmitRun_SdkV2.
func (o *SubmitRun_SdkV2) SetEmailNotifications(ctx context.Context, v JobEmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	o.EmailNotifications = types.ListValueMust(t, vs)
}

// GetEnvironments returns the value of the Environments field in SubmitRun_SdkV2 as
// a slice of JobEnvironment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitRun_SdkV2) GetEnvironments(ctx context.Context) ([]JobEnvironment_SdkV2, bool) {
	if o.Environments.IsNull() || o.Environments.IsUnknown() {
		return nil, false
	}
	var v []JobEnvironment_SdkV2
	d := o.Environments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironments sets the value of the Environments field in SubmitRun_SdkV2.
func (o *SubmitRun_SdkV2) SetEnvironments(ctx context.Context, v []JobEnvironment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Environments = types.ListValueMust(t, vs)
}

// GetGitSource returns the value of the GitSource field in SubmitRun_SdkV2 as
// a GitSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitRun_SdkV2) GetGitSource(ctx context.Context) (GitSource_SdkV2, bool) {
	var e GitSource_SdkV2
	if o.GitSource.IsNull() || o.GitSource.IsUnknown() {
		return e, false
	}
	var v []GitSource_SdkV2
	d := o.GitSource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGitSource sets the value of the GitSource field in SubmitRun_SdkV2.
func (o *SubmitRun_SdkV2) SetGitSource(ctx context.Context, v GitSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["git_source"]
	o.GitSource = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in SubmitRun_SdkV2 as
// a JobsHealthRules_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitRun_SdkV2) GetHealth(ctx context.Context) (JobsHealthRules_SdkV2, bool) {
	var e JobsHealthRules_SdkV2
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v []JobsHealthRules_SdkV2
	d := o.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in SubmitRun_SdkV2.
func (o *SubmitRun_SdkV2) SetHealth(ctx context.Context, v JobsHealthRules_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	o.Health = types.ListValueMust(t, vs)
}

// GetNotificationSettings returns the value of the NotificationSettings field in SubmitRun_SdkV2 as
// a JobNotificationSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitRun_SdkV2) GetNotificationSettings(ctx context.Context) (JobNotificationSettings_SdkV2, bool) {
	var e JobNotificationSettings_SdkV2
	if o.NotificationSettings.IsNull() || o.NotificationSettings.IsUnknown() {
		return e, false
	}
	var v []JobNotificationSettings_SdkV2
	d := o.NotificationSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotificationSettings sets the value of the NotificationSettings field in SubmitRun_SdkV2.
func (o *SubmitRun_SdkV2) SetNotificationSettings(ctx context.Context, v JobNotificationSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notification_settings"]
	o.NotificationSettings = types.ListValueMust(t, vs)
}

// GetQueue returns the value of the Queue field in SubmitRun_SdkV2 as
// a QueueSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitRun_SdkV2) GetQueue(ctx context.Context) (QueueSettings_SdkV2, bool) {
	var e QueueSettings_SdkV2
	if o.Queue.IsNull() || o.Queue.IsUnknown() {
		return e, false
	}
	var v []QueueSettings_SdkV2
	d := o.Queue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueue sets the value of the Queue field in SubmitRun_SdkV2.
func (o *SubmitRun_SdkV2) SetQueue(ctx context.Context, v QueueSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["queue"]
	o.Queue = types.ListValueMust(t, vs)
}

// GetRunAs returns the value of the RunAs field in SubmitRun_SdkV2 as
// a JobRunAs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitRun_SdkV2) GetRunAs(ctx context.Context) (JobRunAs_SdkV2, bool) {
	var e JobRunAs_SdkV2
	if o.RunAs.IsNull() || o.RunAs.IsUnknown() {
		return e, false
	}
	var v []JobRunAs_SdkV2
	d := o.RunAs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunAs sets the value of the RunAs field in SubmitRun_SdkV2.
func (o *SubmitRun_SdkV2) SetRunAs(ctx context.Context, v JobRunAs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_as"]
	o.RunAs = types.ListValueMust(t, vs)
}

// GetTasks returns the value of the Tasks field in SubmitRun_SdkV2 as
// a slice of SubmitTask_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitRun_SdkV2) GetTasks(ctx context.Context) ([]SubmitTask_SdkV2, bool) {
	if o.Tasks.IsNull() || o.Tasks.IsUnknown() {
		return nil, false
	}
	var v []SubmitTask_SdkV2
	d := o.Tasks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTasks sets the value of the Tasks field in SubmitRun_SdkV2.
func (o *SubmitRun_SdkV2) SetTasks(ctx context.Context, v []SubmitTask_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tasks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tasks = types.ListValueMust(t, vs)
}

// GetWebhookNotifications returns the value of the WebhookNotifications field in SubmitRun_SdkV2 as
// a WebhookNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitRun_SdkV2) GetWebhookNotifications(ctx context.Context) (WebhookNotifications_SdkV2, bool) {
	var e WebhookNotifications_SdkV2
	if o.WebhookNotifications.IsNull() || o.WebhookNotifications.IsUnknown() {
		return e, false
	}
	var v []WebhookNotifications_SdkV2
	d := o.WebhookNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhookNotifications sets the value of the WebhookNotifications field in SubmitRun_SdkV2.
func (o *SubmitRun_SdkV2) SetWebhookNotifications(ctx context.Context, v WebhookNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook_notifications"]
	o.WebhookNotifications = types.ListValueMust(t, vs)
}

// Run was created and started successfully.
type SubmitRunResponse_SdkV2 struct {
	// The canonical identifier for the newly submitted run.
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
}

func (newState *SubmitRunResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubmitRunResponse_SdkV2) {
}

func (newState *SubmitRunResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState SubmitRunResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubmitRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubmitRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SubmitRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SubmitRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SubmitRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

type SubmitTask_SdkV2 struct {
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	CleanRoomsNotebookTask types.List `tfsdk:"clean_rooms_notebook_task" tf:"optional,object"`
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
	NewCluster types.List `tfsdk:"new_cluster" tf:"optional,object"`
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

func (newState *SubmitTask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubmitTask_SdkV2) {
}

func (newState *SubmitTask_SdkV2) SyncEffectiveFieldsDuringRead(existingState SubmitTask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubmitTask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubmitTask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_rooms_notebook_task": reflect.TypeOf(CleanRoomsNotebookTask_SdkV2{}),
		"condition_task":            reflect.TypeOf(ConditionTask_SdkV2{}),
		"dbt_task":                  reflect.TypeOf(DbtTask_SdkV2{}),
		"depends_on":                reflect.TypeOf(TaskDependency_SdkV2{}),
		"email_notifications":       reflect.TypeOf(JobEmailNotifications_SdkV2{}),
		"for_each_task":             reflect.TypeOf(ForEachTask_SdkV2{}),
		"health":                    reflect.TypeOf(JobsHealthRules_SdkV2{}),
		"library":                   reflect.TypeOf(compute_tf.Library_SdkV2{}),
		"new_cluster":               reflect.TypeOf(compute_tf.ClusterSpec_SdkV2{}),
		"notebook_task":             reflect.TypeOf(NotebookTask_SdkV2{}),
		"notification_settings":     reflect.TypeOf(TaskNotificationSettings_SdkV2{}),
		"pipeline_task":             reflect.TypeOf(PipelineTask_SdkV2{}),
		"python_wheel_task":         reflect.TypeOf(PythonWheelTask_SdkV2{}),
		"run_job_task":              reflect.TypeOf(RunJobTask_SdkV2{}),
		"spark_jar_task":            reflect.TypeOf(SparkJarTask_SdkV2{}),
		"spark_python_task":         reflect.TypeOf(SparkPythonTask_SdkV2{}),
		"spark_submit_task":         reflect.TypeOf(SparkSubmitTask_SdkV2{}),
		"sql_task":                  reflect.TypeOf(SqlTask_SdkV2{}),
		"webhook_notifications":     reflect.TypeOf(WebhookNotifications_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SubmitTask_SdkV2
// only implements ToObjectValue() and Type().
func (o SubmitTask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_rooms_notebook_task": o.CleanRoomsNotebookTask,
			"condition_task":            o.ConditionTask,
			"dbt_task":                  o.DbtTask,
			"depends_on":                o.DependsOn,
			"description":               o.Description,
			"email_notifications":       o.EmailNotifications,
			"environment_key":           o.EnvironmentKey,
			"existing_cluster_id":       o.ExistingClusterId,
			"for_each_task":             o.ForEachTask,
			"health":                    o.Health,
			"library":                   o.Libraries,
			"new_cluster":               o.NewCluster,
			"notebook_task":             o.NotebookTask,
			"notification_settings":     o.NotificationSettings,
			"pipeline_task":             o.PipelineTask,
			"python_wheel_task":         o.PythonWheelTask,
			"run_if":                    o.RunIf,
			"run_job_task":              o.RunJobTask,
			"spark_jar_task":            o.SparkJarTask,
			"spark_python_task":         o.SparkPythonTask,
			"spark_submit_task":         o.SparkSubmitTask,
			"sql_task":                  o.SqlTask,
			"task_key":                  o.TaskKey,
			"timeout_seconds":           o.TimeoutSeconds,
			"webhook_notifications":     o.WebhookNotifications,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SubmitTask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_rooms_notebook_task": basetypes.ListType{
				ElemType: CleanRoomsNotebookTask{}.Type(ctx),
			},
			"condition_task": basetypes.ListType{
				ElemType: ConditionTask{}.Type(ctx),
			},
			"dbt_task": basetypes.ListType{
				ElemType: DbtTask{}.Type(ctx),
			},
			"depends_on": basetypes.ListType{
				ElemType: TaskDependency{}.Type(ctx),
			},
			"description": types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: JobEmailNotifications{}.Type(ctx),
			},
			"environment_key":     types.StringType,
			"existing_cluster_id": types.StringType,
			"for_each_task": basetypes.ListType{
				ElemType: ForEachTask{}.Type(ctx),
			},
			"health": basetypes.ListType{
				ElemType: JobsHealthRules{}.Type(ctx),
			},
			"library": basetypes.ListType{
				ElemType: compute_tf.Library{}.Type(ctx),
			},
			"new_cluster": basetypes.ListType{
				ElemType: compute_tf.ClusterSpec{}.Type(ctx),
			},
			"notebook_task": basetypes.ListType{
				ElemType: NotebookTask{}.Type(ctx),
			},
			"notification_settings": basetypes.ListType{
				ElemType: TaskNotificationSettings{}.Type(ctx),
			},
			"pipeline_task": basetypes.ListType{
				ElemType: PipelineTask{}.Type(ctx),
			},
			"python_wheel_task": basetypes.ListType{
				ElemType: PythonWheelTask{}.Type(ctx),
			},
			"run_if": types.StringType,
			"run_job_task": basetypes.ListType{
				ElemType: RunJobTask{}.Type(ctx),
			},
			"spark_jar_task": basetypes.ListType{
				ElemType: SparkJarTask{}.Type(ctx),
			},
			"spark_python_task": basetypes.ListType{
				ElemType: SparkPythonTask{}.Type(ctx),
			},
			"spark_submit_task": basetypes.ListType{
				ElemType: SparkSubmitTask{}.Type(ctx),
			},
			"sql_task": basetypes.ListType{
				ElemType: SqlTask{}.Type(ctx),
			},
			"task_key":        types.StringType,
			"timeout_seconds": types.Int64Type,
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.Type(ctx),
			},
		},
	}
}

// GetCleanRoomsNotebookTask returns the value of the CleanRoomsNotebookTask field in SubmitTask_SdkV2 as
// a CleanRoomsNotebookTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetCleanRoomsNotebookTask(ctx context.Context) (CleanRoomsNotebookTask_SdkV2, bool) {
	var e CleanRoomsNotebookTask_SdkV2
	if o.CleanRoomsNotebookTask.IsNull() || o.CleanRoomsNotebookTask.IsUnknown() {
		return e, false
	}
	var v []CleanRoomsNotebookTask_SdkV2
	d := o.CleanRoomsNotebookTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCleanRoomsNotebookTask sets the value of the CleanRoomsNotebookTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetCleanRoomsNotebookTask(ctx context.Context, v CleanRoomsNotebookTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_rooms_notebook_task"]
	o.CleanRoomsNotebookTask = types.ListValueMust(t, vs)
}

// GetConditionTask returns the value of the ConditionTask field in SubmitTask_SdkV2 as
// a ConditionTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetConditionTask(ctx context.Context) (ConditionTask_SdkV2, bool) {
	var e ConditionTask_SdkV2
	if o.ConditionTask.IsNull() || o.ConditionTask.IsUnknown() {
		return e, false
	}
	var v []ConditionTask_SdkV2
	d := o.ConditionTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConditionTask sets the value of the ConditionTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetConditionTask(ctx context.Context, v ConditionTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition_task"]
	o.ConditionTask = types.ListValueMust(t, vs)
}

// GetDbtTask returns the value of the DbtTask field in SubmitTask_SdkV2 as
// a DbtTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetDbtTask(ctx context.Context) (DbtTask_SdkV2, bool) {
	var e DbtTask_SdkV2
	if o.DbtTask.IsNull() || o.DbtTask.IsUnknown() {
		return e, false
	}
	var v []DbtTask_SdkV2
	d := o.DbtTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbtTask sets the value of the DbtTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetDbtTask(ctx context.Context, v DbtTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbt_task"]
	o.DbtTask = types.ListValueMust(t, vs)
}

// GetDependsOn returns the value of the DependsOn field in SubmitTask_SdkV2 as
// a slice of TaskDependency_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetDependsOn(ctx context.Context) ([]TaskDependency_SdkV2, bool) {
	if o.DependsOn.IsNull() || o.DependsOn.IsUnknown() {
		return nil, false
	}
	var v []TaskDependency_SdkV2
	d := o.DependsOn.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependsOn sets the value of the DependsOn field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetDependsOn(ctx context.Context, v []TaskDependency_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["depends_on"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DependsOn = types.ListValueMust(t, vs)
}

// GetEmailNotifications returns the value of the EmailNotifications field in SubmitTask_SdkV2 as
// a JobEmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetEmailNotifications(ctx context.Context) (JobEmailNotifications_SdkV2, bool) {
	var e JobEmailNotifications_SdkV2
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []JobEmailNotifications_SdkV2
	d := o.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetEmailNotifications(ctx context.Context, v JobEmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	o.EmailNotifications = types.ListValueMust(t, vs)
}

// GetForEachTask returns the value of the ForEachTask field in SubmitTask_SdkV2 as
// a ForEachTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetForEachTask(ctx context.Context) (ForEachTask_SdkV2, bool) {
	var e ForEachTask_SdkV2
	if o.ForEachTask.IsNull() || o.ForEachTask.IsUnknown() {
		return e, false
	}
	var v []ForEachTask_SdkV2
	d := o.ForEachTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetForEachTask sets the value of the ForEachTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetForEachTask(ctx context.Context, v ForEachTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["for_each_task"]
	o.ForEachTask = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in SubmitTask_SdkV2 as
// a JobsHealthRules_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetHealth(ctx context.Context) (JobsHealthRules_SdkV2, bool) {
	var e JobsHealthRules_SdkV2
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v []JobsHealthRules_SdkV2
	d := o.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetHealth(ctx context.Context, v JobsHealthRules_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	o.Health = types.ListValueMust(t, vs)
}

// GetLibraries returns the value of the Libraries field in SubmitTask_SdkV2 as
// a slice of compute_tf.Library_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetLibraries(ctx context.Context) ([]compute_tf.Library_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []compute_tf.Library_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetLibraries(ctx context.Context, v []compute_tf.Library_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["library"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

// GetNewCluster returns the value of the NewCluster field in SubmitTask_SdkV2 as
// a compute_tf.ClusterSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetNewCluster(ctx context.Context) (compute_tf.ClusterSpec_SdkV2, bool) {
	var e compute_tf.ClusterSpec_SdkV2
	if o.NewCluster.IsNull() || o.NewCluster.IsUnknown() {
		return e, false
	}
	var v []compute_tf.ClusterSpec_SdkV2
	d := o.NewCluster.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNewCluster sets the value of the NewCluster field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetNewCluster(ctx context.Context, v compute_tf.ClusterSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["new_cluster"]
	o.NewCluster = types.ListValueMust(t, vs)
}

// GetNotebookTask returns the value of the NotebookTask field in SubmitTask_SdkV2 as
// a NotebookTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetNotebookTask(ctx context.Context) (NotebookTask_SdkV2, bool) {
	var e NotebookTask_SdkV2
	if o.NotebookTask.IsNull() || o.NotebookTask.IsUnknown() {
		return e, false
	}
	var v []NotebookTask_SdkV2
	d := o.NotebookTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebookTask sets the value of the NotebookTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetNotebookTask(ctx context.Context, v NotebookTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_task"]
	o.NotebookTask = types.ListValueMust(t, vs)
}

// GetNotificationSettings returns the value of the NotificationSettings field in SubmitTask_SdkV2 as
// a TaskNotificationSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetNotificationSettings(ctx context.Context) (TaskNotificationSettings_SdkV2, bool) {
	var e TaskNotificationSettings_SdkV2
	if o.NotificationSettings.IsNull() || o.NotificationSettings.IsUnknown() {
		return e, false
	}
	var v []TaskNotificationSettings_SdkV2
	d := o.NotificationSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotificationSettings sets the value of the NotificationSettings field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetNotificationSettings(ctx context.Context, v TaskNotificationSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notification_settings"]
	o.NotificationSettings = types.ListValueMust(t, vs)
}

// GetPipelineTask returns the value of the PipelineTask field in SubmitTask_SdkV2 as
// a PipelineTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetPipelineTask(ctx context.Context) (PipelineTask_SdkV2, bool) {
	var e PipelineTask_SdkV2
	if o.PipelineTask.IsNull() || o.PipelineTask.IsUnknown() {
		return e, false
	}
	var v []PipelineTask_SdkV2
	d := o.PipelineTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPipelineTask sets the value of the PipelineTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetPipelineTask(ctx context.Context, v PipelineTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pipeline_task"]
	o.PipelineTask = types.ListValueMust(t, vs)
}

// GetPythonWheelTask returns the value of the PythonWheelTask field in SubmitTask_SdkV2 as
// a PythonWheelTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetPythonWheelTask(ctx context.Context) (PythonWheelTask_SdkV2, bool) {
	var e PythonWheelTask_SdkV2
	if o.PythonWheelTask.IsNull() || o.PythonWheelTask.IsUnknown() {
		return e, false
	}
	var v []PythonWheelTask_SdkV2
	d := o.PythonWheelTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPythonWheelTask sets the value of the PythonWheelTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetPythonWheelTask(ctx context.Context, v PythonWheelTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_wheel_task"]
	o.PythonWheelTask = types.ListValueMust(t, vs)
}

// GetRunJobTask returns the value of the RunJobTask field in SubmitTask_SdkV2 as
// a RunJobTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetRunJobTask(ctx context.Context) (RunJobTask_SdkV2, bool) {
	var e RunJobTask_SdkV2
	if o.RunJobTask.IsNull() || o.RunJobTask.IsUnknown() {
		return e, false
	}
	var v []RunJobTask_SdkV2
	d := o.RunJobTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunJobTask sets the value of the RunJobTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetRunJobTask(ctx context.Context, v RunJobTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_job_task"]
	o.RunJobTask = types.ListValueMust(t, vs)
}

// GetSparkJarTask returns the value of the SparkJarTask field in SubmitTask_SdkV2 as
// a SparkJarTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetSparkJarTask(ctx context.Context) (SparkJarTask_SdkV2, bool) {
	var e SparkJarTask_SdkV2
	if o.SparkJarTask.IsNull() || o.SparkJarTask.IsUnknown() {
		return e, false
	}
	var v []SparkJarTask_SdkV2
	d := o.SparkJarTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkJarTask sets the value of the SparkJarTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetSparkJarTask(ctx context.Context, v SparkJarTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_jar_task"]
	o.SparkJarTask = types.ListValueMust(t, vs)
}

// GetSparkPythonTask returns the value of the SparkPythonTask field in SubmitTask_SdkV2 as
// a SparkPythonTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetSparkPythonTask(ctx context.Context) (SparkPythonTask_SdkV2, bool) {
	var e SparkPythonTask_SdkV2
	if o.SparkPythonTask.IsNull() || o.SparkPythonTask.IsUnknown() {
		return e, false
	}
	var v []SparkPythonTask_SdkV2
	d := o.SparkPythonTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkPythonTask sets the value of the SparkPythonTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetSparkPythonTask(ctx context.Context, v SparkPythonTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_python_task"]
	o.SparkPythonTask = types.ListValueMust(t, vs)
}

// GetSparkSubmitTask returns the value of the SparkSubmitTask field in SubmitTask_SdkV2 as
// a SparkSubmitTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetSparkSubmitTask(ctx context.Context) (SparkSubmitTask_SdkV2, bool) {
	var e SparkSubmitTask_SdkV2
	if o.SparkSubmitTask.IsNull() || o.SparkSubmitTask.IsUnknown() {
		return e, false
	}
	var v []SparkSubmitTask_SdkV2
	d := o.SparkSubmitTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkSubmitTask sets the value of the SparkSubmitTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetSparkSubmitTask(ctx context.Context, v SparkSubmitTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_submit_task"]
	o.SparkSubmitTask = types.ListValueMust(t, vs)
}

// GetSqlTask returns the value of the SqlTask field in SubmitTask_SdkV2 as
// a SqlTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetSqlTask(ctx context.Context) (SqlTask_SdkV2, bool) {
	var e SqlTask_SdkV2
	if o.SqlTask.IsNull() || o.SqlTask.IsUnknown() {
		return e, false
	}
	var v []SqlTask_SdkV2
	d := o.SqlTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlTask sets the value of the SqlTask field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetSqlTask(ctx context.Context, v SqlTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_task"]
	o.SqlTask = types.ListValueMust(t, vs)
}

// GetWebhookNotifications returns the value of the WebhookNotifications field in SubmitTask_SdkV2 as
// a WebhookNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SubmitTask_SdkV2) GetWebhookNotifications(ctx context.Context) (WebhookNotifications_SdkV2, bool) {
	var e WebhookNotifications_SdkV2
	if o.WebhookNotifications.IsNull() || o.WebhookNotifications.IsUnknown() {
		return e, false
	}
	var v []WebhookNotifications_SdkV2
	d := o.WebhookNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhookNotifications sets the value of the WebhookNotifications field in SubmitTask_SdkV2.
func (o *SubmitTask_SdkV2) SetWebhookNotifications(ctx context.Context, v WebhookNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook_notifications"]
	o.WebhookNotifications = types.ListValueMust(t, vs)
}

type TableUpdateTriggerConfiguration_SdkV2 struct {
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

func (newState *TableUpdateTriggerConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableUpdateTriggerConfiguration_SdkV2) {
}

func (newState *TableUpdateTriggerConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState TableUpdateTriggerConfiguration_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableUpdateTriggerConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableUpdateTriggerConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableUpdateTriggerConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o TableUpdateTriggerConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"condition":                         o.Condition,
			"min_time_between_triggers_seconds": o.MinTimeBetweenTriggersSeconds,
			"table_names":                       o.TableNames,
			"wait_after_last_change_seconds":    o.WaitAfterLastChangeSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableUpdateTriggerConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetTableNames returns the value of the TableNames field in TableUpdateTriggerConfiguration_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableUpdateTriggerConfiguration_SdkV2) GetTableNames(ctx context.Context) ([]types.String, bool) {
	if o.TableNames.IsNull() || o.TableNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.TableNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableNames sets the value of the TableNames field in TableUpdateTriggerConfiguration_SdkV2.
func (o *TableUpdateTriggerConfiguration_SdkV2) SetTableNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TableNames = types.ListValueMust(t, vs)
}

type Task_SdkV2 struct {
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	CleanRoomsNotebookTask types.List `tfsdk:"clean_rooms_notebook_task" tf:"optional,object"`
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
	NewCluster types.List `tfsdk:"new_cluster" tf:"optional,object"`
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

func (newState *Task_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Task_SdkV2) {
}

func (newState *Task_SdkV2) SyncEffectiveFieldsDuringRead(existingState Task_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Task.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Task_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_rooms_notebook_task": reflect.TypeOf(CleanRoomsNotebookTask_SdkV2{}),
		"condition_task":            reflect.TypeOf(ConditionTask_SdkV2{}),
		"dbt_task":                  reflect.TypeOf(DbtTask_SdkV2{}),
		"depends_on":                reflect.TypeOf(TaskDependency_SdkV2{}),
		"email_notifications":       reflect.TypeOf(TaskEmailNotifications_SdkV2{}),
		"for_each_task":             reflect.TypeOf(ForEachTask_SdkV2{}),
		"health":                    reflect.TypeOf(JobsHealthRules_SdkV2{}),
		"library":                   reflect.TypeOf(compute_tf.Library_SdkV2{}),
		"new_cluster":               reflect.TypeOf(compute_tf.ClusterSpec_SdkV2{}),
		"notebook_task":             reflect.TypeOf(NotebookTask_SdkV2{}),
		"notification_settings":     reflect.TypeOf(TaskNotificationSettings_SdkV2{}),
		"pipeline_task":             reflect.TypeOf(PipelineTask_SdkV2{}),
		"python_wheel_task":         reflect.TypeOf(PythonWheelTask_SdkV2{}),
		"run_job_task":              reflect.TypeOf(RunJobTask_SdkV2{}),
		"spark_jar_task":            reflect.TypeOf(SparkJarTask_SdkV2{}),
		"spark_python_task":         reflect.TypeOf(SparkPythonTask_SdkV2{}),
		"spark_submit_task":         reflect.TypeOf(SparkSubmitTask_SdkV2{}),
		"sql_task":                  reflect.TypeOf(SqlTask_SdkV2{}),
		"webhook_notifications":     reflect.TypeOf(WebhookNotifications_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Task_SdkV2
// only implements ToObjectValue() and Type().
func (o Task_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_rooms_notebook_task": o.CleanRoomsNotebookTask,
			"condition_task":            o.ConditionTask,
			"dbt_task":                  o.DbtTask,
			"depends_on":                o.DependsOn,
			"description":               o.Description,
			"disable_auto_optimization": o.DisableAutoOptimization,
			"email_notifications":       o.EmailNotifications,
			"environment_key":           o.EnvironmentKey,
			"existing_cluster_id":       o.ExistingClusterId,
			"for_each_task":             o.ForEachTask,
			"health":                    o.Health,
			"job_cluster_key":           o.JobClusterKey,
			"library":                   o.Libraries,
			"max_retries":               o.MaxRetries,
			"min_retry_interval_millis": o.MinRetryIntervalMillis,
			"new_cluster":               o.NewCluster,
			"notebook_task":             o.NotebookTask,
			"notification_settings":     o.NotificationSettings,
			"pipeline_task":             o.PipelineTask,
			"python_wheel_task":         o.PythonWheelTask,
			"retry_on_timeout":          o.RetryOnTimeout,
			"run_if":                    o.RunIf,
			"run_job_task":              o.RunJobTask,
			"spark_jar_task":            o.SparkJarTask,
			"spark_python_task":         o.SparkPythonTask,
			"spark_submit_task":         o.SparkSubmitTask,
			"sql_task":                  o.SqlTask,
			"task_key":                  o.TaskKey,
			"timeout_seconds":           o.TimeoutSeconds,
			"webhook_notifications":     o.WebhookNotifications,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Task_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_rooms_notebook_task": basetypes.ListType{
				ElemType: CleanRoomsNotebookTask{}.Type(ctx),
			},
			"condition_task": basetypes.ListType{
				ElemType: ConditionTask{}.Type(ctx),
			},
			"dbt_task": basetypes.ListType{
				ElemType: DbtTask{}.Type(ctx),
			},
			"depends_on": basetypes.ListType{
				ElemType: TaskDependency{}.Type(ctx),
			},
			"description":               types.StringType,
			"disable_auto_optimization": types.BoolType,
			"email_notifications": basetypes.ListType{
				ElemType: TaskEmailNotifications{}.Type(ctx),
			},
			"environment_key":     types.StringType,
			"existing_cluster_id": types.StringType,
			"for_each_task": basetypes.ListType{
				ElemType: ForEachTask{}.Type(ctx),
			},
			"health": basetypes.ListType{
				ElemType: JobsHealthRules{}.Type(ctx),
			},
			"job_cluster_key": types.StringType,
			"library": basetypes.ListType{
				ElemType: compute_tf.Library{}.Type(ctx),
			},
			"max_retries":               types.Int64Type,
			"min_retry_interval_millis": types.Int64Type,
			"new_cluster": basetypes.ListType{
				ElemType: compute_tf.ClusterSpec{}.Type(ctx),
			},
			"notebook_task": basetypes.ListType{
				ElemType: NotebookTask{}.Type(ctx),
			},
			"notification_settings": basetypes.ListType{
				ElemType: TaskNotificationSettings{}.Type(ctx),
			},
			"pipeline_task": basetypes.ListType{
				ElemType: PipelineTask{}.Type(ctx),
			},
			"python_wheel_task": basetypes.ListType{
				ElemType: PythonWheelTask{}.Type(ctx),
			},
			"retry_on_timeout": types.BoolType,
			"run_if":           types.StringType,
			"run_job_task": basetypes.ListType{
				ElemType: RunJobTask{}.Type(ctx),
			},
			"spark_jar_task": basetypes.ListType{
				ElemType: SparkJarTask{}.Type(ctx),
			},
			"spark_python_task": basetypes.ListType{
				ElemType: SparkPythonTask{}.Type(ctx),
			},
			"spark_submit_task": basetypes.ListType{
				ElemType: SparkSubmitTask{}.Type(ctx),
			},
			"sql_task": basetypes.ListType{
				ElemType: SqlTask{}.Type(ctx),
			},
			"task_key":        types.StringType,
			"timeout_seconds": types.Int64Type,
			"webhook_notifications": basetypes.ListType{
				ElemType: WebhookNotifications{}.Type(ctx),
			},
		},
	}
}

// GetCleanRoomsNotebookTask returns the value of the CleanRoomsNotebookTask field in Task_SdkV2 as
// a CleanRoomsNotebookTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetCleanRoomsNotebookTask(ctx context.Context) (CleanRoomsNotebookTask_SdkV2, bool) {
	var e CleanRoomsNotebookTask_SdkV2
	if o.CleanRoomsNotebookTask.IsNull() || o.CleanRoomsNotebookTask.IsUnknown() {
		return e, false
	}
	var v []CleanRoomsNotebookTask_SdkV2
	d := o.CleanRoomsNotebookTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCleanRoomsNotebookTask sets the value of the CleanRoomsNotebookTask field in Task_SdkV2.
func (o *Task_SdkV2) SetCleanRoomsNotebookTask(ctx context.Context, v CleanRoomsNotebookTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_rooms_notebook_task"]
	o.CleanRoomsNotebookTask = types.ListValueMust(t, vs)
}

// GetConditionTask returns the value of the ConditionTask field in Task_SdkV2 as
// a ConditionTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetConditionTask(ctx context.Context) (ConditionTask_SdkV2, bool) {
	var e ConditionTask_SdkV2
	if o.ConditionTask.IsNull() || o.ConditionTask.IsUnknown() {
		return e, false
	}
	var v []ConditionTask_SdkV2
	d := o.ConditionTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConditionTask sets the value of the ConditionTask field in Task_SdkV2.
func (o *Task_SdkV2) SetConditionTask(ctx context.Context, v ConditionTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["condition_task"]
	o.ConditionTask = types.ListValueMust(t, vs)
}

// GetDbtTask returns the value of the DbtTask field in Task_SdkV2 as
// a DbtTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetDbtTask(ctx context.Context) (DbtTask_SdkV2, bool) {
	var e DbtTask_SdkV2
	if o.DbtTask.IsNull() || o.DbtTask.IsUnknown() {
		return e, false
	}
	var v []DbtTask_SdkV2
	d := o.DbtTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDbtTask sets the value of the DbtTask field in Task_SdkV2.
func (o *Task_SdkV2) SetDbtTask(ctx context.Context, v DbtTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dbt_task"]
	o.DbtTask = types.ListValueMust(t, vs)
}

// GetDependsOn returns the value of the DependsOn field in Task_SdkV2 as
// a slice of TaskDependency_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetDependsOn(ctx context.Context) ([]TaskDependency_SdkV2, bool) {
	if o.DependsOn.IsNull() || o.DependsOn.IsUnknown() {
		return nil, false
	}
	var v []TaskDependency_SdkV2
	d := o.DependsOn.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependsOn sets the value of the DependsOn field in Task_SdkV2.
func (o *Task_SdkV2) SetDependsOn(ctx context.Context, v []TaskDependency_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["depends_on"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DependsOn = types.ListValueMust(t, vs)
}

// GetEmailNotifications returns the value of the EmailNotifications field in Task_SdkV2 as
// a TaskEmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetEmailNotifications(ctx context.Context) (TaskEmailNotifications_SdkV2, bool) {
	var e TaskEmailNotifications_SdkV2
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []TaskEmailNotifications_SdkV2
	d := o.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in Task_SdkV2.
func (o *Task_SdkV2) SetEmailNotifications(ctx context.Context, v TaskEmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	o.EmailNotifications = types.ListValueMust(t, vs)
}

// GetForEachTask returns the value of the ForEachTask field in Task_SdkV2 as
// a ForEachTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetForEachTask(ctx context.Context) (ForEachTask_SdkV2, bool) {
	var e ForEachTask_SdkV2
	if o.ForEachTask.IsNull() || o.ForEachTask.IsUnknown() {
		return e, false
	}
	var v []ForEachTask_SdkV2
	d := o.ForEachTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetForEachTask sets the value of the ForEachTask field in Task_SdkV2.
func (o *Task_SdkV2) SetForEachTask(ctx context.Context, v ForEachTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["for_each_task"]
	o.ForEachTask = types.ListValueMust(t, vs)
}

// GetHealth returns the value of the Health field in Task_SdkV2 as
// a JobsHealthRules_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetHealth(ctx context.Context) (JobsHealthRules_SdkV2, bool) {
	var e JobsHealthRules_SdkV2
	if o.Health.IsNull() || o.Health.IsUnknown() {
		return e, false
	}
	var v []JobsHealthRules_SdkV2
	d := o.Health.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHealth sets the value of the Health field in Task_SdkV2.
func (o *Task_SdkV2) SetHealth(ctx context.Context, v JobsHealthRules_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["health"]
	o.Health = types.ListValueMust(t, vs)
}

// GetLibraries returns the value of the Libraries field in Task_SdkV2 as
// a slice of compute_tf.Library_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetLibraries(ctx context.Context) ([]compute_tf.Library_SdkV2, bool) {
	if o.Libraries.IsNull() || o.Libraries.IsUnknown() {
		return nil, false
	}
	var v []compute_tf.Library_SdkV2
	d := o.Libraries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLibraries sets the value of the Libraries field in Task_SdkV2.
func (o *Task_SdkV2) SetLibraries(ctx context.Context, v []compute_tf.Library_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["library"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Libraries = types.ListValueMust(t, vs)
}

// GetNewCluster returns the value of the NewCluster field in Task_SdkV2 as
// a compute_tf.ClusterSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetNewCluster(ctx context.Context) (compute_tf.ClusterSpec_SdkV2, bool) {
	var e compute_tf.ClusterSpec_SdkV2
	if o.NewCluster.IsNull() || o.NewCluster.IsUnknown() {
		return e, false
	}
	var v []compute_tf.ClusterSpec_SdkV2
	d := o.NewCluster.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNewCluster sets the value of the NewCluster field in Task_SdkV2.
func (o *Task_SdkV2) SetNewCluster(ctx context.Context, v compute_tf.ClusterSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["new_cluster"]
	o.NewCluster = types.ListValueMust(t, vs)
}

// GetNotebookTask returns the value of the NotebookTask field in Task_SdkV2 as
// a NotebookTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetNotebookTask(ctx context.Context) (NotebookTask_SdkV2, bool) {
	var e NotebookTask_SdkV2
	if o.NotebookTask.IsNull() || o.NotebookTask.IsUnknown() {
		return e, false
	}
	var v []NotebookTask_SdkV2
	d := o.NotebookTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebookTask sets the value of the NotebookTask field in Task_SdkV2.
func (o *Task_SdkV2) SetNotebookTask(ctx context.Context, v NotebookTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_task"]
	o.NotebookTask = types.ListValueMust(t, vs)
}

// GetNotificationSettings returns the value of the NotificationSettings field in Task_SdkV2 as
// a TaskNotificationSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetNotificationSettings(ctx context.Context) (TaskNotificationSettings_SdkV2, bool) {
	var e TaskNotificationSettings_SdkV2
	if o.NotificationSettings.IsNull() || o.NotificationSettings.IsUnknown() {
		return e, false
	}
	var v []TaskNotificationSettings_SdkV2
	d := o.NotificationSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotificationSettings sets the value of the NotificationSettings field in Task_SdkV2.
func (o *Task_SdkV2) SetNotificationSettings(ctx context.Context, v TaskNotificationSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notification_settings"]
	o.NotificationSettings = types.ListValueMust(t, vs)
}

// GetPipelineTask returns the value of the PipelineTask field in Task_SdkV2 as
// a PipelineTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetPipelineTask(ctx context.Context) (PipelineTask_SdkV2, bool) {
	var e PipelineTask_SdkV2
	if o.PipelineTask.IsNull() || o.PipelineTask.IsUnknown() {
		return e, false
	}
	var v []PipelineTask_SdkV2
	d := o.PipelineTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPipelineTask sets the value of the PipelineTask field in Task_SdkV2.
func (o *Task_SdkV2) SetPipelineTask(ctx context.Context, v PipelineTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pipeline_task"]
	o.PipelineTask = types.ListValueMust(t, vs)
}

// GetPythonWheelTask returns the value of the PythonWheelTask field in Task_SdkV2 as
// a PythonWheelTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetPythonWheelTask(ctx context.Context) (PythonWheelTask_SdkV2, bool) {
	var e PythonWheelTask_SdkV2
	if o.PythonWheelTask.IsNull() || o.PythonWheelTask.IsUnknown() {
		return e, false
	}
	var v []PythonWheelTask_SdkV2
	d := o.PythonWheelTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPythonWheelTask sets the value of the PythonWheelTask field in Task_SdkV2.
func (o *Task_SdkV2) SetPythonWheelTask(ctx context.Context, v PythonWheelTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["python_wheel_task"]
	o.PythonWheelTask = types.ListValueMust(t, vs)
}

// GetRunJobTask returns the value of the RunJobTask field in Task_SdkV2 as
// a RunJobTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetRunJobTask(ctx context.Context) (RunJobTask_SdkV2, bool) {
	var e RunJobTask_SdkV2
	if o.RunJobTask.IsNull() || o.RunJobTask.IsUnknown() {
		return e, false
	}
	var v []RunJobTask_SdkV2
	d := o.RunJobTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunJobTask sets the value of the RunJobTask field in Task_SdkV2.
func (o *Task_SdkV2) SetRunJobTask(ctx context.Context, v RunJobTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_job_task"]
	o.RunJobTask = types.ListValueMust(t, vs)
}

// GetSparkJarTask returns the value of the SparkJarTask field in Task_SdkV2 as
// a SparkJarTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetSparkJarTask(ctx context.Context) (SparkJarTask_SdkV2, bool) {
	var e SparkJarTask_SdkV2
	if o.SparkJarTask.IsNull() || o.SparkJarTask.IsUnknown() {
		return e, false
	}
	var v []SparkJarTask_SdkV2
	d := o.SparkJarTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkJarTask sets the value of the SparkJarTask field in Task_SdkV2.
func (o *Task_SdkV2) SetSparkJarTask(ctx context.Context, v SparkJarTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_jar_task"]
	o.SparkJarTask = types.ListValueMust(t, vs)
}

// GetSparkPythonTask returns the value of the SparkPythonTask field in Task_SdkV2 as
// a SparkPythonTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetSparkPythonTask(ctx context.Context) (SparkPythonTask_SdkV2, bool) {
	var e SparkPythonTask_SdkV2
	if o.SparkPythonTask.IsNull() || o.SparkPythonTask.IsUnknown() {
		return e, false
	}
	var v []SparkPythonTask_SdkV2
	d := o.SparkPythonTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkPythonTask sets the value of the SparkPythonTask field in Task_SdkV2.
func (o *Task_SdkV2) SetSparkPythonTask(ctx context.Context, v SparkPythonTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_python_task"]
	o.SparkPythonTask = types.ListValueMust(t, vs)
}

// GetSparkSubmitTask returns the value of the SparkSubmitTask field in Task_SdkV2 as
// a SparkSubmitTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetSparkSubmitTask(ctx context.Context) (SparkSubmitTask_SdkV2, bool) {
	var e SparkSubmitTask_SdkV2
	if o.SparkSubmitTask.IsNull() || o.SparkSubmitTask.IsUnknown() {
		return e, false
	}
	var v []SparkSubmitTask_SdkV2
	d := o.SparkSubmitTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparkSubmitTask sets the value of the SparkSubmitTask field in Task_SdkV2.
func (o *Task_SdkV2) SetSparkSubmitTask(ctx context.Context, v SparkSubmitTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spark_submit_task"]
	o.SparkSubmitTask = types.ListValueMust(t, vs)
}

// GetSqlTask returns the value of the SqlTask field in Task_SdkV2 as
// a SqlTask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetSqlTask(ctx context.Context) (SqlTask_SdkV2, bool) {
	var e SqlTask_SdkV2
	if o.SqlTask.IsNull() || o.SqlTask.IsUnknown() {
		return e, false
	}
	var v []SqlTask_SdkV2
	d := o.SqlTask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSqlTask sets the value of the SqlTask field in Task_SdkV2.
func (o *Task_SdkV2) SetSqlTask(ctx context.Context, v SqlTask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sql_task"]
	o.SqlTask = types.ListValueMust(t, vs)
}

// GetWebhookNotifications returns the value of the WebhookNotifications field in Task_SdkV2 as
// a WebhookNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Task_SdkV2) GetWebhookNotifications(ctx context.Context) (WebhookNotifications_SdkV2, bool) {
	var e WebhookNotifications_SdkV2
	if o.WebhookNotifications.IsNull() || o.WebhookNotifications.IsUnknown() {
		return e, false
	}
	var v []WebhookNotifications_SdkV2
	d := o.WebhookNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhookNotifications sets the value of the WebhookNotifications field in Task_SdkV2.
func (o *Task_SdkV2) SetWebhookNotifications(ctx context.Context, v WebhookNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook_notifications"]
	o.WebhookNotifications = types.ListValueMust(t, vs)
}

type TaskDependency_SdkV2 struct {
	// Can only be specified on condition task dependencies. The outcome of the
	// dependent task that must be met for this task to run.
	Outcome types.String `tfsdk:"outcome" tf:"optional"`
	// The name of the task this task depends on.
	TaskKey types.String `tfsdk:"task_key" tf:""`
}

func (newState *TaskDependency_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TaskDependency_SdkV2) {
}

func (newState *TaskDependency_SdkV2) SyncEffectiveFieldsDuringRead(existingState TaskDependency_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TaskDependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TaskDependency_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TaskDependency_SdkV2
// only implements ToObjectValue() and Type().
func (o TaskDependency_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"outcome":  o.Outcome,
			"task_key": o.TaskKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TaskDependency_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"outcome":  types.StringType,
			"task_key": types.StringType,
		},
	}
}

type TaskEmailNotifications_SdkV2 struct {
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

func (newState *TaskEmailNotifications_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TaskEmailNotifications_SdkV2) {
}

func (newState *TaskEmailNotifications_SdkV2) SyncEffectiveFieldsDuringRead(existingState TaskEmailNotifications_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TaskEmailNotifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TaskEmailNotifications_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_duration_warning_threshold_exceeded": reflect.TypeOf(types.String{}),
		"on_failure":                             reflect.TypeOf(types.String{}),
		"on_start":                               reflect.TypeOf(types.String{}),
		"on_streaming_backlog_exceeded":          reflect.TypeOf(types.String{}),
		"on_success":                             reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TaskEmailNotifications_SdkV2
// only implements ToObjectValue() and Type().
func (o TaskEmailNotifications_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"no_alert_for_skipped_runs":              o.NoAlertForSkippedRuns,
			"on_duration_warning_threshold_exceeded": o.OnDurationWarningThresholdExceeded,
			"on_failure":                             o.OnFailure,
			"on_start":                               o.OnStart,
			"on_streaming_backlog_exceeded":          o.OnStreamingBacklogExceeded,
			"on_success":                             o.OnSuccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TaskEmailNotifications_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetOnDurationWarningThresholdExceeded returns the value of the OnDurationWarningThresholdExceeded field in TaskEmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TaskEmailNotifications_SdkV2) GetOnDurationWarningThresholdExceeded(ctx context.Context) ([]types.String, bool) {
	if o.OnDurationWarningThresholdExceeded.IsNull() || o.OnDurationWarningThresholdExceeded.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnDurationWarningThresholdExceeded.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnDurationWarningThresholdExceeded sets the value of the OnDurationWarningThresholdExceeded field in TaskEmailNotifications_SdkV2.
func (o *TaskEmailNotifications_SdkV2) SetOnDurationWarningThresholdExceeded(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_duration_warning_threshold_exceeded"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnDurationWarningThresholdExceeded = types.ListValueMust(t, vs)
}

// GetOnFailure returns the value of the OnFailure field in TaskEmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TaskEmailNotifications_SdkV2) GetOnFailure(ctx context.Context) ([]types.String, bool) {
	if o.OnFailure.IsNull() || o.OnFailure.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnFailure.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnFailure sets the value of the OnFailure field in TaskEmailNotifications_SdkV2.
func (o *TaskEmailNotifications_SdkV2) SetOnFailure(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_failure"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnFailure = types.ListValueMust(t, vs)
}

// GetOnStart returns the value of the OnStart field in TaskEmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TaskEmailNotifications_SdkV2) GetOnStart(ctx context.Context) ([]types.String, bool) {
	if o.OnStart.IsNull() || o.OnStart.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnStart.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnStart sets the value of the OnStart field in TaskEmailNotifications_SdkV2.
func (o *TaskEmailNotifications_SdkV2) SetOnStart(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_start"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnStart = types.ListValueMust(t, vs)
}

// GetOnStreamingBacklogExceeded returns the value of the OnStreamingBacklogExceeded field in TaskEmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TaskEmailNotifications_SdkV2) GetOnStreamingBacklogExceeded(ctx context.Context) ([]types.String, bool) {
	if o.OnStreamingBacklogExceeded.IsNull() || o.OnStreamingBacklogExceeded.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnStreamingBacklogExceeded.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnStreamingBacklogExceeded sets the value of the OnStreamingBacklogExceeded field in TaskEmailNotifications_SdkV2.
func (o *TaskEmailNotifications_SdkV2) SetOnStreamingBacklogExceeded(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_streaming_backlog_exceeded"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnStreamingBacklogExceeded = types.ListValueMust(t, vs)
}

// GetOnSuccess returns the value of the OnSuccess field in TaskEmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TaskEmailNotifications_SdkV2) GetOnSuccess(ctx context.Context) ([]types.String, bool) {
	if o.OnSuccess.IsNull() || o.OnSuccess.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnSuccess.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnSuccess sets the value of the OnSuccess field in TaskEmailNotifications_SdkV2.
func (o *TaskEmailNotifications_SdkV2) SetOnSuccess(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_success"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnSuccess = types.ListValueMust(t, vs)
}

type TaskNotificationSettings_SdkV2 struct {
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

func (newState *TaskNotificationSettings_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TaskNotificationSettings_SdkV2) {
}

func (newState *TaskNotificationSettings_SdkV2) SyncEffectiveFieldsDuringRead(existingState TaskNotificationSettings_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TaskNotificationSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TaskNotificationSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TaskNotificationSettings_SdkV2
// only implements ToObjectValue() and Type().
func (o TaskNotificationSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alert_on_last_attempt":      o.AlertOnLastAttempt,
			"no_alert_for_canceled_runs": o.NoAlertForCanceledRuns,
			"no_alert_for_skipped_runs":  o.NoAlertForSkippedRuns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TaskNotificationSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alert_on_last_attempt":      types.BoolType,
			"no_alert_for_canceled_runs": types.BoolType,
			"no_alert_for_skipped_runs":  types.BoolType,
		},
	}
}

type TerminationDetails_SdkV2 struct {
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
	Type_ types.String `tfsdk:"type" tf:"optional"`
}

func (newState *TerminationDetails_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TerminationDetails_SdkV2) {
}

func (newState *TerminationDetails_SdkV2) SyncEffectiveFieldsDuringRead(existingState TerminationDetails_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TerminationDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TerminationDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TerminationDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o TerminationDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"code":    o.Code,
			"message": o.Message,
			"type":    o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TerminationDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"code":    types.StringType,
			"message": types.StringType,
			"type":    types.StringType,
		},
	}
}

// Additional details about what triggered the run
type TriggerInfo_SdkV2 struct {
	// The run id of the Run Job task run
	RunId types.Int64 `tfsdk:"run_id" tf:"optional"`
}

func (newState *TriggerInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TriggerInfo_SdkV2) {
}

func (newState *TriggerInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState TriggerInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TriggerInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TriggerInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TriggerInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o TriggerInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TriggerInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.Int64Type,
		},
	}
}

type TriggerSettings_SdkV2 struct {
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

func (newState *TriggerSettings_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TriggerSettings_SdkV2) {
}

func (newState *TriggerSettings_SdkV2) SyncEffectiveFieldsDuringRead(existingState TriggerSettings_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TriggerSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TriggerSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_arrival": reflect.TypeOf(FileArrivalTriggerConfiguration_SdkV2{}),
		"periodic":     reflect.TypeOf(PeriodicTriggerConfiguration_SdkV2{}),
		"table":        reflect.TypeOf(TableUpdateTriggerConfiguration_SdkV2{}),
		"table_update": reflect.TypeOf(TableUpdateTriggerConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TriggerSettings_SdkV2
// only implements ToObjectValue() and Type().
func (o TriggerSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_arrival": o.FileArrival,
			"pause_status": o.PauseStatus,
			"periodic":     o.Periodic,
			"table":        o.Table,
			"table_update": o.TableUpdate,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TriggerSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_arrival": basetypes.ListType{
				ElemType: FileArrivalTriggerConfiguration{}.Type(ctx),
			},
			"pause_status": types.StringType,
			"periodic": basetypes.ListType{
				ElemType: PeriodicTriggerConfiguration{}.Type(ctx),
			},
			"table": basetypes.ListType{
				ElemType: TableUpdateTriggerConfiguration{}.Type(ctx),
			},
			"table_update": basetypes.ListType{
				ElemType: TableUpdateTriggerConfiguration{}.Type(ctx),
			},
		},
	}
}

// GetFileArrival returns the value of the FileArrival field in TriggerSettings_SdkV2 as
// a FileArrivalTriggerConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TriggerSettings_SdkV2) GetFileArrival(ctx context.Context) (FileArrivalTriggerConfiguration_SdkV2, bool) {
	var e FileArrivalTriggerConfiguration_SdkV2
	if o.FileArrival.IsNull() || o.FileArrival.IsUnknown() {
		return e, false
	}
	var v []FileArrivalTriggerConfiguration_SdkV2
	d := o.FileArrival.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileArrival sets the value of the FileArrival field in TriggerSettings_SdkV2.
func (o *TriggerSettings_SdkV2) SetFileArrival(ctx context.Context, v FileArrivalTriggerConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["file_arrival"]
	o.FileArrival = types.ListValueMust(t, vs)
}

// GetPeriodic returns the value of the Periodic field in TriggerSettings_SdkV2 as
// a PeriodicTriggerConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TriggerSettings_SdkV2) GetPeriodic(ctx context.Context) (PeriodicTriggerConfiguration_SdkV2, bool) {
	var e PeriodicTriggerConfiguration_SdkV2
	if o.Periodic.IsNull() || o.Periodic.IsUnknown() {
		return e, false
	}
	var v []PeriodicTriggerConfiguration_SdkV2
	d := o.Periodic.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPeriodic sets the value of the Periodic field in TriggerSettings_SdkV2.
func (o *TriggerSettings_SdkV2) SetPeriodic(ctx context.Context, v PeriodicTriggerConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["periodic"]
	o.Periodic = types.ListValueMust(t, vs)
}

// GetTable returns the value of the Table field in TriggerSettings_SdkV2 as
// a TableUpdateTriggerConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TriggerSettings_SdkV2) GetTable(ctx context.Context) (TableUpdateTriggerConfiguration_SdkV2, bool) {
	var e TableUpdateTriggerConfiguration_SdkV2
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []TableUpdateTriggerConfiguration_SdkV2
	d := o.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in TriggerSettings_SdkV2.
func (o *TriggerSettings_SdkV2) SetTable(ctx context.Context, v TableUpdateTriggerConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	o.Table = types.ListValueMust(t, vs)
}

// GetTableUpdate returns the value of the TableUpdate field in TriggerSettings_SdkV2 as
// a TableUpdateTriggerConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TriggerSettings_SdkV2) GetTableUpdate(ctx context.Context) (TableUpdateTriggerConfiguration_SdkV2, bool) {
	var e TableUpdateTriggerConfiguration_SdkV2
	if o.TableUpdate.IsNull() || o.TableUpdate.IsUnknown() {
		return e, false
	}
	var v []TableUpdateTriggerConfiguration_SdkV2
	d := o.TableUpdate.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableUpdate sets the value of the TableUpdate field in TriggerSettings_SdkV2.
func (o *TriggerSettings_SdkV2) SetTableUpdate(ctx context.Context, v TableUpdateTriggerConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table_update"]
	o.TableUpdate = types.ListValueMust(t, vs)
}

type UpdateJob_SdkV2 struct {
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

func (newState *UpdateJob_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateJob_SdkV2) {
}

func (newState *UpdateJob_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateJob_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateJob_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fields_to_remove": reflect.TypeOf(types.String{}),
		"new_settings":     reflect.TypeOf(JobSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateJob_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateJob_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fields_to_remove": o.FieldsToRemove,
			"job_id":           o.JobId,
			"new_settings":     o.NewSettings,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateJob_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fields_to_remove": basetypes.ListType{
				ElemType: types.StringType,
			},
			"job_id": types.Int64Type,
			"new_settings": basetypes.ListType{
				ElemType: JobSettings{}.Type(ctx),
			},
		},
	}
}

// GetFieldsToRemove returns the value of the FieldsToRemove field in UpdateJob_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateJob_SdkV2) GetFieldsToRemove(ctx context.Context) ([]types.String, bool) {
	if o.FieldsToRemove.IsNull() || o.FieldsToRemove.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.FieldsToRemove.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFieldsToRemove sets the value of the FieldsToRemove field in UpdateJob_SdkV2.
func (o *UpdateJob_SdkV2) SetFieldsToRemove(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["fields_to_remove"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FieldsToRemove = types.ListValueMust(t, vs)
}

// GetNewSettings returns the value of the NewSettings field in UpdateJob_SdkV2 as
// a JobSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateJob_SdkV2) GetNewSettings(ctx context.Context) (JobSettings_SdkV2, bool) {
	var e JobSettings_SdkV2
	if o.NewSettings.IsNull() || o.NewSettings.IsUnknown() {
		return e, false
	}
	var v []JobSettings_SdkV2
	d := o.NewSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNewSettings sets the value of the NewSettings field in UpdateJob_SdkV2.
func (o *UpdateJob_SdkV2) SetNewSettings(ctx context.Context, v JobSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["new_settings"]
	o.NewSettings = types.ListValueMust(t, vs)
}

type UpdateResponse_SdkV2 struct {
}

func (newState *UpdateResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse_SdkV2) {
}

func (newState *UpdateResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ViewItem_SdkV2 struct {
	// Content of the view.
	Content types.String `tfsdk:"content" tf:"optional"`
	// Name of the view item. In the case of code view, it would be the
	// notebook’s name. In the case of dashboard view, it would be the
	// dashboard’s name.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Type of the view item.
	Type_ types.String `tfsdk:"type" tf:"optional"`
}

func (newState *ViewItem_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ViewItem_SdkV2) {
}

func (newState *ViewItem_SdkV2) SyncEffectiveFieldsDuringRead(existingState ViewItem_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ViewItem.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ViewItem_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ViewItem_SdkV2
// only implements ToObjectValue() and Type().
func (o ViewItem_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": o.Content,
			"name":    o.Name,
			"type":    o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ViewItem_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"name":    types.StringType,
			"type":    types.StringType,
		},
	}
}

type Webhook_SdkV2 struct {
	Id types.String `tfsdk:"id" tf:""`
}

func (newState *Webhook_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Webhook_SdkV2) {
}

func (newState *Webhook_SdkV2) SyncEffectiveFieldsDuringRead(existingState Webhook_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Webhook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Webhook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Webhook_SdkV2
// only implements ToObjectValue() and Type().
func (o Webhook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Webhook_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type WebhookNotifications_SdkV2 struct {
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

func (newState *WebhookNotifications_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WebhookNotifications_SdkV2) {
}

func (newState *WebhookNotifications_SdkV2) SyncEffectiveFieldsDuringRead(existingState WebhookNotifications_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WebhookNotifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WebhookNotifications_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_duration_warning_threshold_exceeded": reflect.TypeOf(Webhook_SdkV2{}),
		"on_failure":                             reflect.TypeOf(Webhook_SdkV2{}),
		"on_start":                               reflect.TypeOf(Webhook_SdkV2{}),
		"on_streaming_backlog_exceeded":          reflect.TypeOf(Webhook_SdkV2{}),
		"on_success":                             reflect.TypeOf(Webhook_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WebhookNotifications_SdkV2
// only implements ToObjectValue() and Type().
func (o WebhookNotifications_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"on_duration_warning_threshold_exceeded": o.OnDurationWarningThresholdExceeded,
			"on_failure":                             o.OnFailure,
			"on_start":                               o.OnStart,
			"on_streaming_backlog_exceeded":          o.OnStreamingBacklogExceeded,
			"on_success":                             o.OnSuccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WebhookNotifications_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"on_duration_warning_threshold_exceeded": basetypes.ListType{
				ElemType: Webhook{}.Type(ctx),
			},
			"on_failure": basetypes.ListType{
				ElemType: Webhook{}.Type(ctx),
			},
			"on_start": basetypes.ListType{
				ElemType: Webhook{}.Type(ctx),
			},
			"on_streaming_backlog_exceeded": basetypes.ListType{
				ElemType: Webhook{}.Type(ctx),
			},
			"on_success": basetypes.ListType{
				ElemType: Webhook{}.Type(ctx),
			},
		},
	}
}

// GetOnDurationWarningThresholdExceeded returns the value of the OnDurationWarningThresholdExceeded field in WebhookNotifications_SdkV2 as
// a slice of Webhook_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WebhookNotifications_SdkV2) GetOnDurationWarningThresholdExceeded(ctx context.Context) ([]Webhook_SdkV2, bool) {
	if o.OnDurationWarningThresholdExceeded.IsNull() || o.OnDurationWarningThresholdExceeded.IsUnknown() {
		return nil, false
	}
	var v []Webhook_SdkV2
	d := o.OnDurationWarningThresholdExceeded.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnDurationWarningThresholdExceeded sets the value of the OnDurationWarningThresholdExceeded field in WebhookNotifications_SdkV2.
func (o *WebhookNotifications_SdkV2) SetOnDurationWarningThresholdExceeded(ctx context.Context, v []Webhook_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_duration_warning_threshold_exceeded"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnDurationWarningThresholdExceeded = types.ListValueMust(t, vs)
}

// GetOnFailure returns the value of the OnFailure field in WebhookNotifications_SdkV2 as
// a slice of Webhook_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WebhookNotifications_SdkV2) GetOnFailure(ctx context.Context) ([]Webhook_SdkV2, bool) {
	if o.OnFailure.IsNull() || o.OnFailure.IsUnknown() {
		return nil, false
	}
	var v []Webhook_SdkV2
	d := o.OnFailure.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnFailure sets the value of the OnFailure field in WebhookNotifications_SdkV2.
func (o *WebhookNotifications_SdkV2) SetOnFailure(ctx context.Context, v []Webhook_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_failure"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnFailure = types.ListValueMust(t, vs)
}

// GetOnStart returns the value of the OnStart field in WebhookNotifications_SdkV2 as
// a slice of Webhook_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WebhookNotifications_SdkV2) GetOnStart(ctx context.Context) ([]Webhook_SdkV2, bool) {
	if o.OnStart.IsNull() || o.OnStart.IsUnknown() {
		return nil, false
	}
	var v []Webhook_SdkV2
	d := o.OnStart.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnStart sets the value of the OnStart field in WebhookNotifications_SdkV2.
func (o *WebhookNotifications_SdkV2) SetOnStart(ctx context.Context, v []Webhook_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_start"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnStart = types.ListValueMust(t, vs)
}

// GetOnStreamingBacklogExceeded returns the value of the OnStreamingBacklogExceeded field in WebhookNotifications_SdkV2 as
// a slice of Webhook_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WebhookNotifications_SdkV2) GetOnStreamingBacklogExceeded(ctx context.Context) ([]Webhook_SdkV2, bool) {
	if o.OnStreamingBacklogExceeded.IsNull() || o.OnStreamingBacklogExceeded.IsUnknown() {
		return nil, false
	}
	var v []Webhook_SdkV2
	d := o.OnStreamingBacklogExceeded.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnStreamingBacklogExceeded sets the value of the OnStreamingBacklogExceeded field in WebhookNotifications_SdkV2.
func (o *WebhookNotifications_SdkV2) SetOnStreamingBacklogExceeded(ctx context.Context, v []Webhook_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_streaming_backlog_exceeded"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnStreamingBacklogExceeded = types.ListValueMust(t, vs)
}

// GetOnSuccess returns the value of the OnSuccess field in WebhookNotifications_SdkV2 as
// a slice of Webhook_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WebhookNotifications_SdkV2) GetOnSuccess(ctx context.Context) ([]Webhook_SdkV2, bool) {
	if o.OnSuccess.IsNull() || o.OnSuccess.IsUnknown() {
		return nil, false
	}
	var v []Webhook_SdkV2
	d := o.OnSuccess.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnSuccess sets the value of the OnSuccess field in WebhookNotifications_SdkV2.
func (o *WebhookNotifications_SdkV2) SetOnSuccess(ctx context.Context, v []Webhook_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_success"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnSuccess = types.ListValueMust(t, vs)
}
