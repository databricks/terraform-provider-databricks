package jobs

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/jobs"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/databricks/terraform-provider-databricks/repos"
)

// NotebookTask contains the information for notebook jobs
type NotebookTask struct {
	NotebookPath   string            `json:"notebook_path"`
	Source         string            `json:"source,omitempty" tf:"suppress_diff"`
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
}

// SparkPythonTask contains the information for python jobs
type SparkPythonTask struct {
	PythonFile string   `json:"python_file"`
	Source     string   `json:"source,omitempty" tf:"suppress_diff"`
	Parameters []string `json:"parameters,omitempty"`
}

// SparkJarTask contains the information for jar jobs
type SparkJarTask struct {
	JarURI        string   `json:"jar_uri,omitempty"`
	MainClassName string   `json:"main_class_name,omitempty"`
	Parameters    []string `json:"parameters,omitempty"`
}

// SparkSubmitTask contains the information for spark submit jobs
type SparkSubmitTask struct {
	Parameters []string `json:"parameters,omitempty"`
}

// PythonWheelTask contains the information for python wheel jobs
type PythonWheelTask struct {
	EntryPoint      string            `json:"entry_point,omitempty"`
	PackageName     string            `json:"package_name,omitempty"`
	Parameters      []string          `json:"parameters,omitempty"`
	NamedParameters map[string]string `json:"named_parameters,omitempty"`
}

// PipelineTask contains the information for pipeline jobs
type PipelineTask struct {
	PipelineID  string `json:"pipeline_id"`
	FullRefresh bool   `json:"full_refresh,omitempty"`
}

type SqlQueryTask struct {
	QueryID string `json:"query_id"`
}

type SqlSubscription struct {
	UserName      string `json:"user_name,omitempty"`
	DestinationID string `json:"destination_id,omitempty"`
}

type SqlDashboardTask struct {
	DashboardID        string            `json:"dashboard_id"`
	Subscriptions      []SqlSubscription `json:"subscriptions,omitempty"`
	CustomSubject      string            `json:"custom_subject,omitempty"`
	PauseSubscriptions bool              `json:"pause_subscriptions,omitempty"`
}

type SqlAlertTask struct {
	AlertID            string            `json:"alert_id"`
	Subscriptions      []SqlSubscription `json:"subscriptions"`
	PauseSubscriptions bool              `json:"pause_subscriptions,omitempty"`
}

type SqlFileTask struct {
	Path string `json:"path"`
}

// SqlTask contains information about DBSQL task
// TODO: add validation & conflictsWith
type SqlTask struct {
	Query       *SqlQueryTask     `json:"query,omitempty"`
	Dashboard   *SqlDashboardTask `json:"dashboard,omitempty"`
	Alert       *SqlAlertTask     `json:"alert,omitempty"`
	File        *SqlFileTask      `json:"file,omitempty"`
	WarehouseID string            `json:"warehouse_id,omitempty"`
	Parameters  map[string]string `json:"parameters,omitempty"`
}

// DbtTask contains information about DBT task
// TODO: add validation for non-empty commands
type DbtTask struct {
	Commands          []string `json:"commands"`
	ProfilesDirectory string   `json:"profiles_directory,omitempty"`
	ProjectDirectory  string   `json:"project_directory,omitempty"`
	Schema            string   `json:"schema,omitempty" tf:"default:default"`
	Catalog           string   `json:"catalog,omitempty"`
	WarehouseId       string   `json:"warehouse_id,omitempty"`
}

// EmailNotifications contains the information for email notifications after job or task run start or completion
type EmailNotifications struct {
	OnStart               []string `json:"on_start,omitempty"`
	OnSuccess             []string `json:"on_success,omitempty"`
	OnFailure             []string `json:"on_failure,omitempty"`
	NoAlertForSkippedRuns bool     `json:"no_alert_for_skipped_runs,omitempty"`
	AlertOnLastAttempt    bool     `json:"alert_on_last_attempt,omitempty"`
}

// WebhookNotifications contains the information for webhook notifications sent after job start or completion.
type WebhookNotifications struct {
	OnStart   []Webhook `json:"on_start,omitempty"`
	OnSuccess []Webhook `json:"on_success,omitempty"`
	OnFailure []Webhook `json:"on_failure,omitempty"`
}

// NotificationSettings control the notification settings for a job
type NotificationSettings struct {
	NoAlertForSkippedRuns  bool `json:"no_alert_for_skipped_runs,omitempty"`
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
}

func (wn *WebhookNotifications) Sort() {
	if wn == nil {
		return
	}

	notifs := [][]Webhook{wn.OnStart, wn.OnFailure, wn.OnSuccess}
	for _, ns := range notifs {
		sort.Slice(ns, func(i, j int) bool {
			return ns[i].ID < ns[j].ID
		})
	}
}

// Webhook contains a reference by id to one of the centrally configured webhooks.
type Webhook struct {
	ID string `json:"id"`
}

// CronSchedule contains the information for the quartz cron expression
type CronSchedule struct {
	QuartzCronExpression string `json:"quartz_cron_expression"`
	TimezoneID           string `json:"timezone_id"`
	PauseStatus          string `json:"pause_status,omitempty" tf:"computed"`
}

// BEGIN Jobs + Repo integration preview
type GitSource struct {
	Url      string `json:"git_url" tf:"alias:url"`
	Provider string `json:"git_provider,omitempty" tf:"alias:provider"`
	Branch   string `json:"git_branch,omitempty" tf:"alias:branch"`
	Tag      string `json:"git_tag,omitempty" tf:"alias:tag"`
	Commit   string `json:"git_commit,omitempty" tf:"alias:commit"`
}

// End Jobs + Repo integration preview

type JobTaskSettings struct {
	TaskKey     string                `json:"task_key,omitempty"`
	Description string                `json:"description,omitempty"`
	DependsOn   []jobs.TaskDependency `json:"depends_on,omitempty"`

	// BEGIN Jobs + RunIf preview
	RunIf string `json:"run_if,omitempty" tf:"suppress_diff"`
	// END Jobs + RunIf preview

	ExistingClusterID string              `json:"existing_cluster_id,omitempty" tf:"group:cluster_type"`
	NewCluster        *clusters.Cluster   `json:"new_cluster,omitempty" tf:"group:cluster_type"`
	JobClusterKey     string              `json:"job_cluster_key,omitempty" tf:"group:cluster_type"`
	ComputeKey        string              `json:"compute_key,omitempty" tf:"group:cluster_type"`
	Libraries         []libraries.Library `json:"libraries,omitempty" tf:"slice_set,alias:library"`

	NotebookTask    *NotebookTask    `json:"notebook_task,omitempty" tf:"group:task_type"`
	SparkJarTask    *SparkJarTask    `json:"spark_jar_task,omitempty" tf:"group:task_type"`
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty" tf:"group:task_type"`
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty" tf:"group:task_type"`
	PipelineTask    *PipelineTask    `json:"pipeline_task,omitempty" tf:"group:task_type"`
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty" tf:"group:task_type"`
	SqlTask         *SqlTask         `json:"sql_task,omitempty" tf:"group:task_type"`
	DbtTask         *DbtTask         `json:"dbt_task,omitempty" tf:"group:task_type"`

	// ConditionTask is in private preview
	ConditionTask *jobs.ConditionTask `json:"condition_task,omitempty" tf:"group:task_type"`

	EmailNotifications     *EmailNotifications `json:"email_notifications,omitempty" tf:"suppress_diff"`
	TimeoutSeconds         int32               `json:"timeout_seconds,omitempty"`
	MaxRetries             int32               `json:"max_retries,omitempty"`
	MinRetryIntervalMillis int32               `json:"min_retry_interval_millis,omitempty"`
	RetryOnTimeout         bool                `json:"retry_on_timeout,omitempty" tf:"computed"`
}

type JobCluster struct {
	JobClusterKey string            `json:"job_cluster_key,omitempty" tf:"group:cluster_type"`
	NewCluster    *clusters.Cluster `json:"new_cluster,omitempty" tf:"group:cluster_type"`
}

type JobCompute struct {
	ComputeKey  string               `json:"compute_key,omitempty" tf:"group:cluster_type"`
	ComputeSpec *compute.ComputeSpec `json:"spec,omitempty" tf:"group:cluster_type"`
}

type ContinuousConf struct {
	PauseStatus string `json:"pause_status,omitempty" tf:"computed"`
}

type Queue struct {
}

type JobRunAs struct {
	UserName             string `json:"user_name,omitempty"`
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
}

type FileArrival struct {
	URL                           string `json:"url"`
	MinTimeBetweenTriggersSeconds int32  `json:"min_time_between_triggers_seconds,omitempty"`
	WaitAfterLastChangeSeconds    int32  `json:"wait_after_last_change_seconds,omitempty"`
}

type Trigger struct {
	FileArrival *FileArrival `json:"file_arrival"`
	PauseStatus string       `json:"pause_status,omitempty" tf:"computed"`
}

// JobSettings contains the information for configuring a job on databricks
type JobSettings struct {
	Name string `json:"name,omitempty" tf:"default:Untitled"`

	// BEGIN Jobs API 2.0
	ExistingClusterID      string              `json:"existing_cluster_id,omitempty" tf:"group:cluster_type"`
	NewCluster             *clusters.Cluster   `json:"new_cluster,omitempty" tf:"group:cluster_type"`
	NotebookTask           *NotebookTask       `json:"notebook_task,omitempty" tf:"group:task_type"`
	SparkJarTask           *SparkJarTask       `json:"spark_jar_task,omitempty" tf:"group:task_type"`
	SparkPythonTask        *SparkPythonTask    `json:"spark_python_task,omitempty" tf:"group:task_type"`
	SparkSubmitTask        *SparkSubmitTask    `json:"spark_submit_task,omitempty" tf:"group:task_type"`
	PipelineTask           *PipelineTask       `json:"pipeline_task,omitempty" tf:"group:task_type"`
	PythonWheelTask        *PythonWheelTask    `json:"python_wheel_task,omitempty" tf:"group:task_type"`
	DbtTask                *DbtTask            `json:"dbt_task,omitempty" tf:"group:task_type"`
	Libraries              []libraries.Library `json:"libraries,omitempty" tf:"slice_set,alias:library"`
	TimeoutSeconds         int32               `json:"timeout_seconds,omitempty"`
	MaxRetries             int32               `json:"max_retries,omitempty"`
	MinRetryIntervalMillis int32               `json:"min_retry_interval_millis,omitempty"`
	RetryOnTimeout         bool                `json:"retry_on_timeout,omitempty"`
	// END Jobs API 2.0

	// BEGIN Jobs API 2.1
	Tasks       []JobTaskSettings `json:"tasks,omitempty" tf:"alias:task"`
	Format      string            `json:"format,omitempty" tf:"computed"`
	JobClusters []JobCluster      `json:"job_clusters,omitempty" tf:"alias:job_cluster"`
	Compute     []JobCompute      `json:"compute,omitempty" tf:"alias:compute"`
	// END Jobs API 2.1

	// BEGIN Jobs + Repo integration preview
	GitSource *GitSource `json:"git_source,omitempty"`
	// END Jobs + Repo integration preview

	Schedule             *CronSchedule         `json:"schedule,omitempty"`
	Continuous           *ContinuousConf       `json:"continuous,omitempty"`
	Trigger              *Trigger              `json:"trigger,omitempty"`
	MaxConcurrentRuns    int32                 `json:"max_concurrent_runs,omitempty"`
	EmailNotifications   *EmailNotifications   `json:"email_notifications,omitempty" tf:"suppress_diff"`
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty" tf:"suppress_diff"`
	NotificationSettings *NotificationSettings `json:"notification_settings,omitempty"`
	Tags                 map[string]string     `json:"tags,omitempty"`
	Queue                *Queue                `json:"queue,omitempty"`
	RunAs                *JobRunAs             `json:"run_as,omitempty"`
}

func (js *JobSettings) isMultiTask() bool {
	return js.Format == "MULTI_TASK" || len(js.Tasks) > 0
}

func (js *JobSettings) sortTasksByKey() {
	sort.Slice(js.Tasks, func(i, j int) bool {
		return js.Tasks[i].TaskKey < js.Tasks[j].TaskKey
	})
}

func (js *JobSettings) sortWebhooksByID() {
	js.WebhookNotifications.Sort()
}

// JobListResponse returns a list of all jobs
type JobListResponse struct {
	Jobs    []Job `json:"jobs"`
	HasMore bool  `json:"has_more,omitempty"`
}

// Job contains the information when using a GET request from the Databricks Jobs api
type Job struct {
	JobID           int64        `json:"job_id,omitempty"`
	CreatorUserName string       `json:"creator_user_name,omitempty"`
	RunAsUserName   string       `json:"run_as_user_name,omitempty" tf:"computed"`
	Settings        *JobSettings `json:"settings,omitempty"`
	CreatedTime     int64        `json:"created_time,omitempty"`
}

// ID returns job id as string
func (j Job) ID() string {
	return fmt.Sprintf("%d", j.JobID)
}

// RunParameters used to pass params to tasks
type RunParameters struct {
	// a shortcut field to reuse this type for RunNow
	JobID int64 `json:"job_id,omitempty"`

	NotebookParams    map[string]string `json:"notebook_params,omitempty"`
	JarParams         []string          `json:"jar_params,omitempty"`
	PythonParams      []string          `json:"python_params,omitempty"`
	SparkSubmitParams []string          `json:"spark_submit_params,omitempty"`
}

// RunState of the job
type RunState struct {
	ResultState    string `json:"result_state,omitempty"`
	LifeCycleState string `json:"life_cycle_state,omitempty"`
	StateMessage   string `json:"state_message,omitempty"`
}

// JobRun is a simplified representation of corresponding entity
type JobRun struct {
	JobID       int64    `json:"job_id,omitempty"`
	RunID       int64    `json:"run_id,omitempty"`
	NumberInJob int64    `json:"number_in_job,omitempty"`
	StartTime   int64    `json:"start_time,omitempty"`
	State       RunState `json:"state,omitempty"`
	Trigger     string   `json:"trigger,omitempty"`
	RuntType    string   `json:"run_type,omitempty"`

	OverridingParameters RunParameters `json:"overriding_parameters,omitempty"`
}

// JobRunsListRequest used to do what it sounds like
type JobRunsListRequest struct {
	JobID         int64 `url:"job_id,omitempty"`
	ActiveOnly    bool  `url:"active_only,omitempty"`
	CompletedOnly bool  `url:"completed_only,omitempty"`
	Offset        int32 `url:"offset,omitempty"`
	Limit         int32 `url:"limit,omitempty"`
}

// JobRunsList returns a page of job runs
type JobRunsList struct {
	Runs    []JobRun `json:"runs"`
	HasMore bool     `json:"has_more"`
}

// UpdateJobRequest used to do what it sounds like
type UpdateJobRequest struct {
	JobID       int64        `json:"job_id,omitempty" url:"job_id,omitempty"`
	NewSettings *JobSettings `json:"new_settings,omitempty" url:"new_settings,omitempty"`
}

// NewJobsAPI creates JobsAPI instance from provider meta
func NewJobsAPI(ctx context.Context, m any) JobsAPI {
	client := m.(*common.DatabricksClient)
	return JobsAPI{client, ctx}
}

// JobsAPI exposes the Jobs API
type JobsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// List all jobs matching the name. If name is empty, returns all jobs
func (a JobsAPI) ListByName(name string, expandTasks bool) ([]Job, error) {
	jobs := []Job{}
	params := map[string]interface{}{
		"limit":        25,
		"expand_tasks": expandTasks,
	}
	if name != "" {
		params["name"] = name
	}
	offset := 0

	ctx := context.WithValue(a.context, common.Api, common.API_2_1)
	for {
		var resp JobListResponse
		params["offset"] = offset
		err := a.client.Get(ctx, "/jobs/list", params, &resp)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, resp.Jobs...)
		if !resp.HasMore {
			break
		}
		offset += len(resp.Jobs)
	}
	return jobs, nil
}

// List all jobs
func (a JobsAPI) List() (l []Job, err error) {
	l, err = a.ListByName("", false)
	return
}

// RunsList returns a job runs list
func (a JobsAPI) RunsList(r JobRunsListRequest) (jrl JobRunsList, err error) {
	err = a.client.Get(a.context, "/jobs/runs/list", r, &jrl)
	return
}

// RunsCancel cancels job run and waits till it's finished
func (a JobsAPI) RunsCancel(runID int64, timeout time.Duration) error {
	var response any
	err := a.client.Post(a.context, "/jobs/runs/cancel", map[string]any{
		"run_id": runID,
	}, &response)
	if err != nil {
		return err
	}
	return a.waitForRunState(runID, "TERMINATED", timeout)
}

func (a JobsAPI) waitForRunState(runID int64, desiredState string, timeout time.Duration) error {
	return resource.RetryContext(a.context, timeout, func() *resource.RetryError {
		jobRun, err := a.RunsGet(runID)
		if err != nil {
			return resource.NonRetryableError(
				fmt.Errorf("cannot get job %s: %v", desiredState, err))
		}
		state := jobRun.State
		if state.LifeCycleState == desiredState {
			return nil
		}
		if state.LifeCycleState == "INTERNAL_ERROR" {
			return resource.NonRetryableError(
				fmt.Errorf("cannot get job %s: %s",
					desiredState, state.StateMessage))
		}
		return resource.RetryableError(
			fmt.Errorf("run is %s: %s",
				state.LifeCycleState,
				state.StateMessage))
	})
}

// RunNow triggers the job and returns a run ID
func (a JobsAPI) RunNow(jobID int64) (int64, error) {
	var jr JobRun
	err := a.client.Post(a.context, "/jobs/run-now", RunParameters{
		JobID: jobID,
	}, &jr)
	return jr.RunID, err
}

// RunsGet to retrieve information about the run
func (a JobsAPI) RunsGet(runID int64) (JobRun, error) {
	var jr JobRun
	err := a.client.Get(a.context, "/jobs/runs/get", map[string]any{
		"run_id": runID,
	}, &jr)
	return jr, err
}

func (a JobsAPI) Start(jobID int64, timeout time.Duration) error {
	runID, err := a.RunNow(jobID)
	if err != nil {
		return fmt.Errorf("cannot start job run: %v", err)
	}
	return a.waitForRunState(runID, "RUNNING", timeout)
}

func (a JobsAPI) StopActiveRun(jobID int64, timeout time.Duration) error {
	runs, err := a.RunsList(JobRunsListRequest{JobID: jobID, ActiveOnly: true})
	if err != nil {
		return err
	}
	if len(runs.Runs) > 1 {
		return fmt.Errorf("`always_running` must be specified only with "+
			"`max_concurrent_runs = 1`. There are %d active runs", len(runs.Runs))
	}
	if len(runs.Runs) == 1 {
		activeRun := runs.Runs[0]
		err = a.RunsCancel(activeRun.RunID, timeout)
		if err != nil {
			return fmt.Errorf("cannot cancel run %d: %v", activeRun.RunID, err)
		}
	}
	return nil
}

// Create creates a job on the workspace given the job settings
func (a JobsAPI) Create(jobSettings JobSettings) (Job, error) {
	var job Job
	jobSettings.sortTasksByKey()
	jobSettings.sortWebhooksByID()
	var gitSource *GitSource = jobSettings.GitSource
	if gitSource != nil && gitSource.Provider == "" {
		gitSource.Provider = repos.GetGitProviderFromUrl(gitSource.Url)
		if gitSource.Provider == "" {
			return job, fmt.Errorf("git source is not empty but Git Provider is not specified and cannot be guessed by url %+v", gitSource)
		}
		if gitSource.Branch == "" && gitSource.Tag == "" && gitSource.Commit == "" {
			return job, fmt.Errorf("git source is not empty but none of branch, commit and tag is specified")
		}
	}
	err := a.client.Post(a.context, "/jobs/create", jobSettings, &job)
	return job, err
}

// Update updates a job given the id and a new set of job settings
func (a JobsAPI) Update(id string, jobSettings JobSettings) error {
	jobID, err := parseJobId(id)
	if err != nil {
		return err
	}
	return wrapMissingJobError(a.client.Post(a.context, "/jobs/reset", UpdateJobRequest{
		JobID:       jobID,
		NewSettings: &jobSettings,
	}, nil), id)
}

// Read returns the job object with all the attributes
func (a JobsAPI) Read(id string) (job Job, err error) {
	jobID, err := parseJobId(id)
	if err != nil {
		return
	}
	err = wrapMissingJobError(a.client.Get(a.context, "/jobs/get", map[string]int64{
		"job_id": jobID,
	}, &job), id)
	if job.Settings != nil {
		job.Settings.sortTasksByKey()
		job.Settings.sortWebhooksByID()
	}
	return
}

// Delete deletes the job given a job id
func (a JobsAPI) Delete(id string) error {
	jobID, err := parseJobId(id)
	if err != nil {
		return err
	}
	return wrapMissingJobError(a.client.Post(a.context, "/jobs/delete", map[string]int64{
		"job_id": jobID,
	}, nil), id)
}

func wrapMissingJobError(err error, id string) error {
	if err == nil {
		return nil
	}
	var apiErr *apierr.APIError
	if !errors.As(err, &apiErr) {
		return err
	}
	if apiErr.IsMissing() {
		return err
	}
	// fix non-compliant error code
	if strings.Contains(apiErr.Message,
		fmt.Sprintf("Job %s does not exist.", id)) {
		apiErr.StatusCode = 404
		return apiErr
	}
	return err
}

func jobSettingsSchema(s *map[string]*schema.Schema, prefix string) {
	if p, err := common.SchemaPath(*s, "new_cluster", "num_workers"); err == nil {
		p.Optional = true
		p.Default = 0
		p.Type = schema.TypeInt
		p.ValidateDiagFunc = validation.ToDiagFunc(validation.IntAtLeast(0))
		p.Required = false
	}
	if v, err := common.SchemaPath(*s, "new_cluster", "spark_conf"); err == nil {
		reSize := common.MustCompileKeyRE(prefix + "new_cluster.0.spark_conf.%")
		reConf := common.MustCompileKeyRE(prefix + "new_cluster.0.spark_conf.spark.databricks.delta.preview.enabled")
		v.DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
			isPossiblyLegacyConfig := reSize.Match([]byte(k)) && old == "1" && new == "0"
			isLegacyConfig := reConf.Match([]byte(k))
			if isPossiblyLegacyConfig || isLegacyConfig {
				log.Printf("[DEBUG] Suppressing diff for k=%#v old=%#v new=%#v", k, old, new)
				return true
			}
			return false
		}
	}
}

func gitSourceSchema(r *schema.Resource, prefix string) {
	r.Schema["url"].ValidateFunc = validation.IsURLWithHTTPS
	(*r.Schema["tag"]).ConflictsWith = []string{"git_source.0.branch", "git_source.0.commit"}
	(*r.Schema["branch"]).ConflictsWith = []string{"git_source.0.commit", "git_source.0.tag"}
	(*r.Schema["commit"]).ConflictsWith = []string{"git_source.0.branch", "git_source.0.tag"}
}

var jobSchema = common.StructToSchema(JobSettings{},
	func(s map[string]*schema.Schema) map[string]*schema.Schema {
		jobSettingsSchema(&s, "")
		jobSettingsSchema(&s["task"].Elem.(*schema.Resource).Schema, "task.0.")
		jobSettingsSchema(&s["job_cluster"].Elem.(*schema.Resource).Schema, "job_cluster.0.")
		gitSourceSchema(s["git_source"].Elem.(*schema.Resource), "")
		if p, err := common.SchemaPath(s, "schedule", "pause_status"); err == nil {
			p.ValidateFunc = validation.StringInSlice([]string{"PAUSED", "UNPAUSED"}, false)
		}
		s["max_concurrent_runs"].ValidateDiagFunc = validation.ToDiagFunc(validation.IntAtLeast(1))
		s["max_concurrent_runs"].Default = 1
		s["url"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		s["always_running"] = &schema.Schema{
			Optional:      true,
			Default:       false,
			Type:          schema.TypeBool,
			Deprecated:    "always_running will be replaced by control_run_state in the next major release.",
			ConflictsWith: []string{"control_run_state", "continuous"},
		}
		s["control_run_state"] = &schema.Schema{
			Optional:      true,
			Default:       false,
			Type:          schema.TypeBool,
			ConflictsWith: []string{"always_running"},
		}
		s["schedule"].ConflictsWith = []string{"continuous", "trigger"}
		s["continuous"].ConflictsWith = []string{"schedule", "trigger"}
		s["trigger"].ConflictsWith = []string{"schedule", "continuous"}
		return s
	})

func parseJobId(id string) (int64, error) {
	return strconv.ParseInt(id, 10, 64)
}

// Callbacks to manage runs for jobs after creation and update.
//
// There are three types of lifecycle management for jobs:
//  1. always_running: When enabled, a new run will be started after the job configuration is updated.
//     An existing active run will be cancelled if one exists.
//  2. control_run_state: When enabled, stops the active run of continuous jobs after the job configuration is updated.
//  3. Noop: No lifecycle management.
//
// always_running is deprecated but still supported for backwards compatibility.
type jobLifecycleManager interface {
	OnCreate(ctx context.Context) error
	OnUpdate(ctx context.Context) error
}

func getJobLifecycleManager(d *schema.ResourceData, m any) jobLifecycleManager {
	if d.Get("always_running").(bool) {
		return alwaysRunningLifecycleManager{d: d, m: m}
	}
	if d.Get("control_run_state").(bool) {
		return controlRunStateLifecycleManager{d: d, m: m}
	}
	return noopLifecycleManager{}
}

type noopLifecycleManager struct{}

func (n noopLifecycleManager) OnCreate(ctx context.Context) error {
	return nil
}
func (n noopLifecycleManager) OnUpdate(ctx context.Context) error {
	return nil
}

type alwaysRunningLifecycleManager struct {
	d *schema.ResourceData
	m any
}

func (a alwaysRunningLifecycleManager) OnCreate(ctx context.Context) error {
	jobID, err := parseJobId(a.d.Id())
	if err != nil {
		return err
	}
	return NewJobsAPI(ctx, a.m).Start(jobID, a.d.Timeout(schema.TimeoutCreate))
}
func (a alwaysRunningLifecycleManager) OnUpdate(ctx context.Context) error {
	api := NewJobsAPI(ctx, a.m)
	jobID, err := parseJobId(a.d.Id())
	if err != nil {
		return err
	}
	err = api.StopActiveRun(jobID, a.d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	return api.Start(jobID, a.d.Timeout(schema.TimeoutUpdate))
}

type controlRunStateLifecycleManager struct {
	d *schema.ResourceData
	m any
}

func (c controlRunStateLifecycleManager) OnCreate(ctx context.Context) error {
	return nil
}

func (c controlRunStateLifecycleManager) OnUpdate(ctx context.Context) error {
	if c.d.Get("continuous") == nil {
		return nil
	}

	jobID, err := parseJobId(c.d.Id())
	if err != nil {
		return err
	}

	api := NewJobsAPI(ctx, c.m)

	// Only use RunNow to stop the active run if the job is unpaused.
	pauseStatus := c.d.Get("continuous.0.pause_status").(string)
	if pauseStatus == "UNPAUSED" {
		// Previously, RunNow() was not supported for continuous jobs. Now, calling RunNow()
		// on a continuous job works, cancelling the active run if there is one, and resetting
		// the exponential backoff timer. So, we try to call RunNow() first, and if it fails,
		// we call StopActiveRun() instead.
		_, err = api.RunNow(jobID)

		if err == nil {
			return nil
		}

		// RunNow() returns 404 when the feature is disabled.
		var apiErr *apierr.APIError
		if errors.As(err, &apiErr) && apiErr.StatusCode != 404 {
			return err
		}
	}

	return api.StopActiveRun(jobID, c.d.Timeout(schema.TimeoutUpdate))
}

func prepareJobSettingsForUpdate(js JobSettings) {
	if js.NewCluster != nil {
		js.NewCluster.ModifyRequestOnInstancePool()
	}
	for _, task := range js.Tasks {
		if task.NewCluster != nil {
			task.NewCluster.ModifyRequestOnInstancePool()
		}
	}
	for _, jc := range js.JobClusters {
		if jc.NewCluster != nil {
			jc.NewCluster.ModifyRequestOnInstancePool()
		}
	}
}

func ResourceJob() *schema.Resource {
	getReadCtx := func(ctx context.Context, d *schema.ResourceData) context.Context {
		var js JobSettings
		common.DataToStructPointer(d, jobSchema, &js)
		if js.isMultiTask() {
			return context.WithValue(ctx, common.Api, common.API_2_1)
		}
		return ctx
	}
	return common.Resource{
		Schema:        jobSchema,
		SchemaVersion: 2,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(clusters.DefaultProvisionTimeout),
			Update: schema.DefaultTimeout(clusters.DefaultProvisionTimeout),
		},
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			var js JobSettings
			common.DiffToStructPointer(d, jobSchema, &js)
			alwaysRunning := d.Get("always_running").(bool)
			if alwaysRunning && js.MaxConcurrentRuns > 1 {
				return fmt.Errorf("`always_running` must be specified only with `max_concurrent_runs = 1`")
			}
			controlRunState := d.Get("control_run_state").(bool)
			if controlRunState {
				if js.Continuous == nil {
					return fmt.Errorf("`control_run_state` must be specified only with `continuous`")
				}
				if js.MaxConcurrentRuns > 1 {
					return fmt.Errorf("`control_run_state` must be specified only with `max_concurrent_runs = 1`")
				}
			}
			for _, task := range js.Tasks {
				if task.NewCluster == nil {
					continue
				}
				if err := task.NewCluster.Validate(); err != nil {
					return fmt.Errorf("task %s invalid: %w", task.TaskKey, err)
				}
			}
			if js.NewCluster != nil {
				if err := js.NewCluster.Validate(); err != nil {
					return fmt.Errorf("invalid job cluster: %w", err)
				}
			}
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var js JobSettings
			common.DataToStructPointer(d, jobSchema, &js)
			if js.isMultiTask() {
				ctx = context.WithValue(ctx, common.Api, common.API_2_1)
			}
			jobsAPI := NewJobsAPI(ctx, c)
			job, err := jobsAPI.Create(js)
			if err != nil {
				return err
			}
			d.SetId(job.ID())
			return getJobLifecycleManager(d, c).OnCreate(ctx)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ctx = getReadCtx(ctx, d)
			job, err := NewJobsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			d.Set("url", c.FormatURL("#job/", d.Id()))
			return common.StructToData(*job.Settings, jobSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var js JobSettings
			common.DataToStructPointer(d, jobSchema, &js)
			if js.isMultiTask() {
				ctx = context.WithValue(ctx, common.Api, common.API_2_1)
			}

			prepareJobSettingsForUpdate(js)

			jobsAPI := NewJobsAPI(ctx, c)
			err := jobsAPI.Update(d.Id(), js)
			if err != nil {
				return err
			}
			return getJobLifecycleManager(d, c).OnUpdate(ctx)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ctx = getReadCtx(ctx, d)
			return NewJobsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
