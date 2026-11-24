// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Jobs, Policy Compliance For Jobs, etc.
package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type JobsInterface interface {

	// WaitGetRunJobTerminatedOrSkipped repeatedly calls [JobsAPI.GetRun] and waits to reach TERMINATED or SKIPPED state
	WaitGetRunJobTerminatedOrSkipped(ctx context.Context, runId int64,
		timeout time.Duration, callback func(*Run)) (*Run, error)

	// Cancels all active runs of a job. The runs are canceled asynchronously, so it
	// doesn't prevent new runs from being started.
	CancelAllRuns(ctx context.Context, request CancelAllRuns) error

	// Cancels a job run or a task run. The run is canceled asynchronously, so it
	// may still be running when this request completes.
	CancelRun(ctx context.Context, cancelRun CancelRun) (*WaitGetRunJobTerminatedOrSkipped[struct{}], error)

	// Calls [JobsAPIInterface.CancelRun] and waits to reach TERMINATED or SKIPPED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[Run](60*time.Minute) functional option.
	//
	// Deprecated: use [JobsAPIInterface.CancelRun].Get() or [JobsAPIInterface.WaitGetRunJobTerminatedOrSkipped]
	CancelRunAndWait(ctx context.Context, cancelRun CancelRun, options ...retries.Option[Run]) (*Run, error)

	// Cancels a job run or a task run. The run is canceled asynchronously, so it
	// may still be running when this request completes.
	CancelRunByRunId(ctx context.Context, runId int64) error

	CancelRunByRunIdAndWait(ctx context.Context, runId int64, options ...retries.Option[Run]) (*Run, error)

	// Create a new job.
	Create(ctx context.Context, request CreateJob) (*CreateResponse, error)

	// Deletes a job.
	Delete(ctx context.Context, request DeleteJob) error

	// Deletes a job.
	DeleteByJobId(ctx context.Context, jobId int64) error

	// Deletes a non-active run. Returns an error if the run is active.
	DeleteRun(ctx context.Context, request DeleteRun) error

	// Deletes a non-active run. Returns an error if the run is active.
	DeleteRunByRunId(ctx context.Context, runId int64) error

	// Export and retrieve the job run task.
	ExportRun(ctx context.Context, request ExportRunRequest) (*ExportRunOutput, error)

	// Retrieves the details for a single job.
	//
	// Large arrays in the results will be paginated when they exceed 100 elements.
	// A request for a single job will return all properties for that job, and the
	// first 100 elements of array properties (`tasks`, `job_clusters`,
	// `environments` and `parameters`). Use the `next_page_token` field to check
	// for more results and pass its value as the `page_token` in subsequent
	// requests. If any array properties have more than 100 elements, additional
	// results will be returned on subsequent requests. Arrays without additional
	// results will be empty on later pages.
	Get(ctx context.Context, request GetJobRequest) (*Job, error)

	// Retrieves the details for a single job.
	//
	// Large arrays in the results will be paginated when they exceed 100 elements.
	// A request for a single job will return all properties for that job, and the
	// first 100 elements of array properties (`tasks`, `job_clusters`,
	// `environments` and `parameters`). Use the `next_page_token` field to check
	// for more results and pass its value as the `page_token` in subsequent
	// requests. If any array properties have more than 100 elements, additional
	// results will be returned on subsequent requests. Arrays without additional
	// results will be empty on later pages.
	GetByJobId(ctx context.Context, jobId int64) (*Job, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetJobPermissionLevelsRequest) (*GetJobPermissionLevelsResponse, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevelsByJobId(ctx context.Context, jobId string) (*GetJobPermissionLevelsResponse, error)

	// Gets the permissions of a job. Jobs can inherit permissions from their root
	// object.
	GetPermissions(ctx context.Context, request GetJobPermissionsRequest) (*JobPermissions, error)

	// Gets the permissions of a job. Jobs can inherit permissions from their root
	// object.
	GetPermissionsByJobId(ctx context.Context, jobId string) (*JobPermissions, error)

	// Retrieves the metadata of a run.
	//
	// Large arrays in the results will be paginated when they exceed 100 elements.
	// A request for a single run will return all properties for that run, and the
	// first 100 elements of array properties (`tasks`, `job_clusters`,
	// `job_parameters` and `repair_history`). Use the next_page_token field to
	// check for more results and pass its value as the page_token in subsequent
	// requests. If any array properties have more than 100 elements, additional
	// results will be returned on subsequent requests. Arrays without additional
	// results will be empty on later pages.
	GetRun(ctx context.Context, request GetRunRequest) (*Run, error)

	// Retrieve the output and metadata of a single task run. When a notebook task
	// returns a value through the `dbutils.notebook.exit()` call, you can use this
	// endpoint to retrieve that value. Databricks restricts this API to returning
	// the first 5 MB of the output. To return a larger result, you can store job
	// results in a cloud storage service.
	//
	// This endpoint validates that the __run_id__ parameter is valid and returns an
	// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
	// automatically removed after 60 days. If you to want to reference them beyond
	// 60 days, you must save old run results before they expire.
	GetRunOutput(ctx context.Context, request GetRunOutputRequest) (*RunOutput, error)

	// Retrieve the output and metadata of a single task run. When a notebook task
	// returns a value through the `dbutils.notebook.exit()` call, you can use this
	// endpoint to retrieve that value. Databricks restricts this API to returning
	// the first 5 MB of the output. To return a larger result, you can store job
	// results in a cloud storage service.
	//
	// This endpoint validates that the __run_id__ parameter is valid and returns an
	// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
	// automatically removed after 60 days. If you to want to reference them beyond
	// 60 days, you must save old run results before they expire.
	GetRunOutputByRunId(ctx context.Context, runId int64) (*RunOutput, error)

	// Retrieves a list of jobs.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListJobsRequest) listing.Iterator[BaseJob]

	// Retrieves a list of jobs.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListJobsRequest) ([]BaseJob, error)

	// BaseJobSettingsNameToJobIdMap calls [JobsAPI.ListAll] and creates a map of results with [BaseJob].Settings.Name as key and [BaseJob].JobId as value.
	//
	// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
	//
	// Note: All [BaseJob] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	BaseJobSettingsNameToJobIdMap(ctx context.Context, request ListJobsRequest) (map[string]int64, error)

	// GetBySettingsName calls [JobsAPI.BaseJobSettingsNameToJobIdMap] and returns a single [BaseJob].
	//
	// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
	//
	// Note: All [BaseJob] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetBySettingsName(ctx context.Context, name string) (*BaseJob, error)

	// List runs in descending order by start time.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListRuns(ctx context.Context, request ListRunsRequest) listing.Iterator[BaseRun]

	// List runs in descending order by start time.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListRunsAll(ctx context.Context, request ListRunsRequest) ([]BaseRun, error)

	// Re-run one or more tasks. Tasks are re-run as part of the original job run.
	// They use the current job and task settings, and can be viewed in the history
	// for the original job run.
	RepairRun(ctx context.Context, repairRun RepairRun) (*WaitGetRunJobTerminatedOrSkipped[RepairRunResponse], error)

	// Calls [JobsAPIInterface.RepairRun] and waits to reach TERMINATED or SKIPPED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[Run](60*time.Minute) functional option.
	//
	// Deprecated: use [JobsAPIInterface.RepairRun].Get() or [JobsAPIInterface.WaitGetRunJobTerminatedOrSkipped]
	RepairRunAndWait(ctx context.Context, repairRun RepairRun, options ...retries.Option[Run]) (*Run, error)

	// Overwrite all settings for the given job. Use the [_Update_
	// endpoint](:method:jobs/update) to update job settings partially.
	Reset(ctx context.Context, request ResetJob) error

	// Run a job and return the `run_id` of the triggered run.
	RunNow(ctx context.Context, runNow RunNow) (*WaitGetRunJobTerminatedOrSkipped[RunNowResponse], error)

	// Calls [JobsAPIInterface.RunNow] and waits to reach TERMINATED or SKIPPED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[Run](60*time.Minute) functional option.
	//
	// Deprecated: use [JobsAPIInterface.RunNow].Get() or [JobsAPIInterface.WaitGetRunJobTerminatedOrSkipped]
	RunNowAndWait(ctx context.Context, runNow RunNow, options ...retries.Option[Run]) (*Run, error)

	// Sets permissions on an object, replacing existing permissions if they exist.
	// Deletes all direct permissions if none are specified. Objects can inherit
	// permissions from their root object.
	SetPermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error)

	// Submit a one-time run. This endpoint allows you to submit a workload directly
	// without creating a job. Runs submitted using this endpoint don’t display in
	// the UI. Use the `jobs/runs/get` API to check the run state after the job is
	// submitted.
	//
	// **Important:** Jobs submitted using this endpoint are not saved as a job.
	// They do not show up in the Jobs UI, and do not retry when they fail. Because
	// they are not saved, Databricks cannot auto-optimize serverless compute in
	// case of failure. If your job fails, you may want to use classic compute to
	// specify the compute needs for the job. Alternatively, use the `POST
	// /jobs/create` and `POST /jobs/run-now` endpoints to create and run a saved
	// job.
	Submit(ctx context.Context, submitRun SubmitRun) (*WaitGetRunJobTerminatedOrSkipped[SubmitRunResponse], error)

	// Calls [JobsAPIInterface.Submit] and waits to reach TERMINATED or SKIPPED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[Run](60*time.Minute) functional option.
	//
	// Deprecated: use [JobsAPIInterface.Submit].Get() or [JobsAPIInterface.WaitGetRunJobTerminatedOrSkipped]
	SubmitAndWait(ctx context.Context, submitRun SubmitRun, options ...retries.Option[Run]) (*Run, error)

	// Add, update, or remove specific settings of an existing job. Use the [_Reset_
	// endpoint](:method:jobs/reset) to overwrite all job settings.
	Update(ctx context.Context, request UpdateJob) error

	// Updates the permissions on a job. Jobs can inherit permissions from their
	// root object.
	UpdatePermissions(ctx context.Context, request JobPermissionsRequest) (*JobPermissions, error)
}

func NewJobs(client *client.DatabricksClient) *JobsAPI {
	return &JobsAPI{
		jobsImpl: jobsImpl{
			client: client,
		},
	}
}

// The Jobs API allows you to create, edit, and delete jobs.
//
// You can use a Databricks job to run a data processing or data analysis task
// in a Databricks cluster with scalable resources. Your job can consist of a
// single task or can be a large, multi-task workflow with complex dependencies.
// Databricks manages the task orchestration, cluster management, monitoring,
// and error reporting for all of your jobs. You can run your jobs immediately
// or periodically through an easy-to-use scheduling system. You can implement
// job tasks using notebooks, JARS, Spark Declarative Pipelines, or Python,
// Scala, Spark submit, and Java applications.
//
// You should never hard code secrets or store them in plain text. Use the
// [Secrets CLI] to manage secrets in the [Databricks CLI]. Use the [Secrets
// utility] to reference secrets in notebooks and jobs.
//
// [Databricks CLI]: https://docs.databricks.com/dev-tools/cli/index.html
// [Secrets CLI]: https://docs.databricks.com/dev-tools/cli/secrets-cli.html
// [Secrets utility]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-secrets
type JobsAPI struct {
	jobsImpl
}

// WaitGetRunJobTerminatedOrSkipped repeatedly calls [JobsAPI.GetRun] and waits to reach TERMINATED or SKIPPED state
func (a *JobsAPI) WaitGetRunJobTerminatedOrSkipped(ctx context.Context, runId int64,
	timeout time.Duration, callback func(*Run)) (*Run, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[Run](ctx, timeout, func() (*Run, *retries.Err) {
		run, err := a.GetRun(ctx, GetRunRequest{
			RunId: runId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(run)
		}
		status := run.State.LifeCycleState
		statusMessage := fmt.Sprintf("current status: %s", status)
		if run.State != nil {
			statusMessage = run.State.StateMessage
		}
		switch status {
		case RunLifeCycleStateTerminated, RunLifeCycleStateSkipped: // target state
			return run, nil
		case RunLifeCycleStateInternalError:
			err := fmt.Errorf("failed to reach %s or %s, got %s: %s",
				RunLifeCycleStateTerminated, RunLifeCycleStateSkipped, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetRunJobTerminatedOrSkipped is a wrapper that calls [JobsAPI.WaitGetRunJobTerminatedOrSkipped] and waits to reach TERMINATED or SKIPPED state.
type WaitGetRunJobTerminatedOrSkipped[R any] struct {
	Response *R
	RunId    int64 `json:"run_id"`
	Poll     func(time.Duration, func(*Run)) (*Run, error)
	callback func(*Run)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetRunJobTerminatedOrSkipped[R]) OnProgress(callback func(*Run)) *WaitGetRunJobTerminatedOrSkipped[R] {
	w.callback = callback
	return w
}

// Get the Run with the default timeout of 20 minutes.
func (w *WaitGetRunJobTerminatedOrSkipped[R]) Get() (*Run, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the Run with custom timeout.
func (w *WaitGetRunJobTerminatedOrSkipped[R]) GetWithTimeout(timeout time.Duration) (*Run, error) {
	return w.Poll(timeout, w.callback)
}

// Cancels a job run or a task run. The run is canceled asynchronously, so it
// may still be running when this request completes.
func (a *JobsAPI) CancelRun(ctx context.Context, cancelRun CancelRun) (*WaitGetRunJobTerminatedOrSkipped[struct{}], error) {
	err := a.jobsImpl.CancelRun(ctx, cancelRun)
	if err != nil {
		return nil, err
	}
	return &WaitGetRunJobTerminatedOrSkipped[struct{}]{

		RunId: cancelRun.RunId,
		Poll: func(timeout time.Duration, callback func(*Run)) (*Run, error) {
			return a.WaitGetRunJobTerminatedOrSkipped(ctx, cancelRun.RunId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [JobsAPI.CancelRun] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
//
// Deprecated: use [JobsAPI.CancelRun].Get() or [JobsAPI.WaitGetRunJobTerminatedOrSkipped]
func (a *JobsAPI) CancelRunAndWait(ctx context.Context, cancelRun CancelRun, options ...retries.Option[Run]) (*Run, error) {
	wait, err := a.CancelRun(ctx, cancelRun)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *Run) {
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Cancels a job run or a task run. The run is canceled asynchronously, so it
// may still be running when this request completes.
func (a *JobsAPI) CancelRunByRunId(ctx context.Context, runId int64) error {
	return a.jobsImpl.CancelRun(ctx, CancelRun{
		RunId: runId,
	})
}

func (a *JobsAPI) CancelRunByRunIdAndWait(ctx context.Context, runId int64, options ...retries.Option[Run]) (*Run, error) {
	return a.CancelRunAndWait(ctx, CancelRun{
		RunId: runId,
	}, options...)
}

// Deletes a job.
func (a *JobsAPI) DeleteByJobId(ctx context.Context, jobId int64) error {
	return a.jobsImpl.Delete(ctx, DeleteJob{
		JobId: jobId,
	})
}

// Deletes a non-active run. Returns an error if the run is active.
func (a *JobsAPI) DeleteRunByRunId(ctx context.Context, runId int64) error {
	return a.jobsImpl.DeleteRun(ctx, DeleteRun{
		RunId: runId,
	})
}

// Retrieves the details for a single job.
//
// Large arrays in the results will be paginated when they exceed 100 elements.
// A request for a single job will return all properties for that job, and the
// first 100 elements of array properties (`tasks`, `job_clusters`,
// `environments` and `parameters`). Use the `next_page_token` field to check
// for more results and pass its value as the `page_token` in subsequent
// requests. If any array properties have more than 100 elements, additional
// results will be returned on subsequent requests. Arrays without additional
// results will be empty on later pages.
func (a *JobsAPI) GetByJobId(ctx context.Context, jobId int64) (*Job, error) {
	return a.jobsImpl.Get(ctx, GetJobRequest{
		JobId: jobId,
	})
}

// Gets the permission levels that a user can have on an object.
func (a *JobsAPI) GetPermissionLevelsByJobId(ctx context.Context, jobId string) (*GetJobPermissionLevelsResponse, error) {
	return a.jobsImpl.GetPermissionLevels(ctx, GetJobPermissionLevelsRequest{
		JobId: jobId,
	})
}

// Gets the permissions of a job. Jobs can inherit permissions from their root
// object.
func (a *JobsAPI) GetPermissionsByJobId(ctx context.Context, jobId string) (*JobPermissions, error) {
	return a.jobsImpl.GetPermissions(ctx, GetJobPermissionsRequest{
		JobId: jobId,
	})
}

// Retrieve the output and metadata of a single task run. When a notebook task
// returns a value through the `dbutils.notebook.exit()` call, you can use this
// endpoint to retrieve that value. Databricks restricts this API to returning
// the first 5 MB of the output. To return a larger result, you can store job
// results in a cloud storage service.
//
// This endpoint validates that the __run_id__ parameter is valid and returns an
// HTTP status code 400 if the __run_id__ parameter is invalid. Runs are
// automatically removed after 60 days. If you to want to reference them beyond
// 60 days, you must save old run results before they expire.
func (a *JobsAPI) GetRunOutputByRunId(ctx context.Context, runId int64) (*RunOutput, error) {
	return a.jobsImpl.GetRunOutput(ctx, GetRunOutputRequest{
		RunId: runId,
	})
}

// BaseJobSettingsNameToJobIdMap calls [JobsAPI.ListAll] and creates a map of results with [BaseJob].Settings.Name as key and [BaseJob].JobId as value.
//
// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
//
// Note: All [BaseJob] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) BaseJobSettingsNameToJobIdMap(ctx context.Context, request ListJobsRequest) (map[string]int64, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]int64{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Settings.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Settings.Name: %s", key)
		}
		mapping[key] = v.JobId
	}
	return mapping, nil
}

// GetBySettingsName calls [JobsAPI.BaseJobSettingsNameToJobIdMap] and returns a single [BaseJob].
//
// Returns an error if there's more than one [BaseJob] with the same .Settings.Name.
//
// Note: All [BaseJob] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *JobsAPI) GetBySettingsName(ctx context.Context, name string) (*BaseJob, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListJobsRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]BaseJob{}
	for _, v := range result {
		key := v.Settings.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("BaseJob named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of BaseJob named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Re-run one or more tasks. Tasks are re-run as part of the original job run.
// They use the current job and task settings, and can be viewed in the history
// for the original job run.
func (a *JobsAPI) RepairRun(ctx context.Context, repairRun RepairRun) (*WaitGetRunJobTerminatedOrSkipped[RepairRunResponse], error) {
	repairRunResponse, err := a.jobsImpl.RepairRun(ctx, repairRun)
	if err != nil {
		return nil, err
	}
	return &WaitGetRunJobTerminatedOrSkipped[RepairRunResponse]{
		Response: repairRunResponse,
		RunId:    repairRun.RunId,
		Poll: func(timeout time.Duration, callback func(*Run)) (*Run, error) {
			return a.WaitGetRunJobTerminatedOrSkipped(ctx, repairRun.RunId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [JobsAPI.RepairRun] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
//
// Deprecated: use [JobsAPI.RepairRun].Get() or [JobsAPI.WaitGetRunJobTerminatedOrSkipped]
func (a *JobsAPI) RepairRunAndWait(ctx context.Context, repairRun RepairRun, options ...retries.Option[Run]) (*Run, error) {
	wait, err := a.RepairRun(ctx, repairRun)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *Run) {
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Run a job and return the `run_id` of the triggered run.
func (a *JobsAPI) RunNow(ctx context.Context, runNow RunNow) (*WaitGetRunJobTerminatedOrSkipped[RunNowResponse], error) {
	runNowResponse, err := a.jobsImpl.RunNow(ctx, runNow)
	if err != nil {
		return nil, err
	}
	return &WaitGetRunJobTerminatedOrSkipped[RunNowResponse]{
		Response: runNowResponse,
		RunId:    runNowResponse.RunId,
		Poll: func(timeout time.Duration, callback func(*Run)) (*Run, error) {
			return a.WaitGetRunJobTerminatedOrSkipped(ctx, runNowResponse.RunId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [JobsAPI.RunNow] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
//
// Deprecated: use [JobsAPI.RunNow].Get() or [JobsAPI.WaitGetRunJobTerminatedOrSkipped]
func (a *JobsAPI) RunNowAndWait(ctx context.Context, runNow RunNow, options ...retries.Option[Run]) (*Run, error) {
	wait, err := a.RunNow(ctx, runNow)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *Run) {
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Submit a one-time run. This endpoint allows you to submit a workload directly
// without creating a job. Runs submitted using this endpoint don’t display in
// the UI. Use the `jobs/runs/get` API to check the run state after the job is
// submitted.
//
// **Important:** Jobs submitted using this endpoint are not saved as a job.
// They do not show up in the Jobs UI, and do not retry when they fail. Because
// they are not saved, Databricks cannot auto-optimize serverless compute in
// case of failure. If your job fails, you may want to use classic compute to
// specify the compute needs for the job. Alternatively, use the `POST
// /jobs/create` and `POST /jobs/run-now` endpoints to create and run a saved
// job.
func (a *JobsAPI) Submit(ctx context.Context, submitRun SubmitRun) (*WaitGetRunJobTerminatedOrSkipped[SubmitRunResponse], error) {
	submitRunResponse, err := a.jobsImpl.Submit(ctx, submitRun)
	if err != nil {
		return nil, err
	}
	return &WaitGetRunJobTerminatedOrSkipped[SubmitRunResponse]{
		Response: submitRunResponse,
		RunId:    submitRunResponse.RunId,
		Poll: func(timeout time.Duration, callback func(*Run)) (*Run, error) {
			return a.WaitGetRunJobTerminatedOrSkipped(ctx, submitRunResponse.RunId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [JobsAPI.Submit] and waits to reach TERMINATED or SKIPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Run](60*time.Minute) functional option.
//
// Deprecated: use [JobsAPI.Submit].Get() or [JobsAPI.WaitGetRunJobTerminatedOrSkipped]
func (a *JobsAPI) SubmitAndWait(ctx context.Context, submitRun SubmitRun, options ...retries.Option[Run]) (*Run, error) {
	wait, err := a.Submit(ctx, submitRun)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[Run]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *Run) {
		for _, o := range options {
			o(&retries.Info[Run]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

type PolicyComplianceForJobsInterface interface {

	// Updates a job so the job clusters that are created when running the job
	// (specified in `new_cluster`) are compliant with the current versions of their
	// respective cluster policies. All-purpose clusters used in the job will not be
	// updated.
	EnforceCompliance(ctx context.Context, request EnforcePolicyComplianceRequest) (*EnforcePolicyComplianceResponse, error)

	// Returns the policy compliance status of a job. Jobs could be out of
	// compliance if a cluster policy they use was updated after the job was last
	// edited and some of its job clusters no longer comply with their updated
	// policies.
	GetCompliance(ctx context.Context, request GetPolicyComplianceRequest) (*GetPolicyComplianceResponse, error)

	// Returns the policy compliance status of a job. Jobs could be out of
	// compliance if a cluster policy they use was updated after the job was last
	// edited and some of its job clusters no longer comply with their updated
	// policies.
	GetComplianceByJobId(ctx context.Context, jobId int64) (*GetPolicyComplianceResponse, error)

	// Returns the policy compliance status of all jobs that use a given policy.
	// Jobs could be out of compliance if a cluster policy they use was updated
	// after the job was last edited and its job clusters no longer comply with the
	// updated policy.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListCompliance(ctx context.Context, request ListJobComplianceRequest) listing.Iterator[JobCompliance]

	// Returns the policy compliance status of all jobs that use a given policy.
	// Jobs could be out of compliance if a cluster policy they use was updated
	// after the job was last edited and its job clusters no longer comply with the
	// updated policy.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListComplianceAll(ctx context.Context, request ListJobComplianceRequest) ([]JobCompliance, error)
}

func NewPolicyComplianceForJobs(client *client.DatabricksClient) *PolicyComplianceForJobsAPI {
	return &PolicyComplianceForJobsAPI{
		policyComplianceForJobsImpl: policyComplianceForJobsImpl{
			client: client,
		},
	}
}

// The compliance APIs allow you to view and manage the policy compliance status
// of jobs in your workspace. This API currently only supports compliance
// controls for cluster policies.
//
// A job is in compliance if its cluster configurations satisfy the rules of all
// their respective cluster policies. A job could be out of compliance if a
// cluster policy it uses was updated after the job was last edited. The job is
// considered out of compliance if any of its clusters no longer comply with
// their updated policies.
//
// The get and list compliance APIs allow you to view the policy compliance
// status of a job. The enforce compliance API allows you to update a job so
// that it becomes compliant with all of its policies.
type PolicyComplianceForJobsAPI struct {
	policyComplianceForJobsImpl
}

// Returns the policy compliance status of a job. Jobs could be out of
// compliance if a cluster policy they use was updated after the job was last
// edited and some of its job clusters no longer comply with their updated
// policies.
func (a *PolicyComplianceForJobsAPI) GetComplianceByJobId(ctx context.Context, jobId int64) (*GetPolicyComplianceResponse, error) {
	return a.policyComplianceForJobsImpl.GetCompliance(ctx, GetPolicyComplianceRequest{
		JobId: jobId,
	})
}
