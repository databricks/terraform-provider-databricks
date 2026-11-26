package jobs

import (
	"context"

	"github.com/databricks/databricks-sdk-go/listing"
)

// List fetches a list of jobs.
// If expand_tasks is true, the response will include the full list of tasks and job_clusters for each job.
// This function handles pagination two ways: paginates all the jobs in the list and paginates all the tasks and job_clusters for each job.
func (a *JobsAPI) List(ctx context.Context, request ListJobsRequest) listing.Iterator[BaseJob] {
	// Fetch jobs with limited elements in top level arrays
	jobsList := a.jobsImpl.List(ctx, request)

	if !request.ExpandTasks {
		return jobsList
	}

	return &expandedJobsIterator{
		originalIterator: jobsList,
		service:          a,
	}
}

// expandedJobsIterator is a custom iterator that for each job calls job/get in order to fetch full list of tasks and job_clusters.
type expandedJobsIterator struct {
	originalIterator listing.Iterator[BaseJob]
	service          *JobsAPI
}

func (e *expandedJobsIterator) HasNext(ctx context.Context) bool {
	return e.originalIterator.HasNext(ctx)
}

func (e *expandedJobsIterator) Next(ctx context.Context) (BaseJob, error) {
	job, err := e.originalIterator.Next(ctx)
	if err != nil {
		return BaseJob{}, err
	}
	if !job.HasMore {
		return job, nil
	}

	// Fully fetch all top level arrays for the job
	getJobRequest := GetJobRequest{JobId: job.JobId}
	fullJob, err := e.service.Get(ctx, getJobRequest)
	if err != nil {
		return BaseJob{}, err
	}

	job.Settings.Tasks = fullJob.Settings.Tasks
	job.Settings.JobClusters = fullJob.Settings.JobClusters
	job.Settings.Parameters = fullJob.Settings.Parameters
	job.Settings.Environments = fullJob.Settings.Environments
	job.HasMore = false

	return job, nil
}

// ListRuns fetches a list of job runs.
// If expand_tasks is true, the response will include the full list of tasks and job_clusters for each run.
// This function handles pagination two ways: paginates all the runs in the list and paginates all the tasks and job_clusters for each run.
func (a *JobsAPI) ListRuns(ctx context.Context, request ListRunsRequest) listing.Iterator[BaseRun] {
	runsList := a.jobsImpl.ListRuns(ctx, request)

	if !request.ExpandTasks {
		return runsList
	}

	return &expandedRunsIterator{
		originalIterator: runsList,
		service:          a,
	}
}

// expandedRunsIterator is a custom iterator that for each run calls runs/get in order to fetch full list of tasks and job_clusters.
type expandedRunsIterator struct {
	originalIterator listing.Iterator[BaseRun]
	service          *JobsAPI
}

func (e *expandedRunsIterator) HasNext(ctx context.Context) bool {
	return e.originalIterator.HasNext(ctx)
}

func (e *expandedRunsIterator) Next(ctx context.Context) (BaseRun, error) {
	run, err := e.originalIterator.Next(ctx)
	if err != nil {
		return BaseRun{}, err
	}
	if !run.HasMore {
		return run, nil
	}

	// Fully fetch all top level arrays for the job run.
	getRunRequest := GetRunRequest{RunId: run.RunId}
	fullRun, err := e.service.GetRun(ctx, getRunRequest)
	if err != nil {
		return BaseRun{}, err
	}

	run.Tasks = fullRun.Tasks
	run.JobClusters = fullRun.JobClusters
	run.JobParameters = fullRun.JobParameters
	run.RepairHistory = fullRun.RepairHistory
	run.HasMore = false

	return run, nil
}

// GetRun retrieves a run based on the provided request.
// It handles pagination if the run contains multiple iterations or tasks.
func (a *JobsAPI) GetRun(ctx context.Context, request GetRunRequest) (*Run, error) {
	run, err := a.jobsImpl.GetRun(ctx, request)
	if err != nil {
		return nil, err
	}

	// When querying a Job run, a page token is returned when there are more than 100 tasks. No iterations are defined for a Job run. Therefore, the next page in the response only includes the next page of tasks.
	// When querying a ForEach task run, a page token is returned when there are more than 100 iterations. Only a single task is returned, corresponding to the ForEach task itself. Therefore, the client only reads the iterations from the next page and not the tasks.
	isPaginatingIterations := len(run.Iterations) > 0

	// runs/get response includes next_page_token as long as there are more pages to fetch.
	for run.NextPageToken != "" {
		request.PageToken = run.NextPageToken
		nextRun, err := a.jobsImpl.GetRun(ctx, request)
		if err != nil {
			return nil, err
		}

		if isPaginatingIterations {
			run.Iterations = append(run.Iterations, nextRun.Iterations...)
		} else {
			run.Tasks = append(run.Tasks, nextRun.Tasks...)
		}
		// Each new page of runs/get response includes the next page of the job_clusters, job_parameters, and repair history.
		run.JobClusters = append(run.JobClusters, nextRun.JobClusters...)
		run.JobParameters = append(run.JobParameters, nextRun.JobParameters...)
		run.RepairHistory = append(run.RepairHistory, nextRun.RepairHistory...)
		run.NextPageToken = nextRun.NextPageToken
	}

	return run, nil
}

// Get retrieves a job based on the provided request.
// It handles pagination if the job contains multiple tasks, job_clusters, job_parameters or environments.
func (a *JobsAPI) Get(ctx context.Context, request GetJobRequest) (*Job, error) {
	job, err := a.jobsImpl.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	// jobs/get response includes next_page_token as long as there are more pages to fetch.
	for job.NextPageToken != "" {
		request.PageToken = job.NextPageToken
		nextJob, err := a.jobsImpl.Get(ctx, request)
		if err != nil {
			return nil, err
		}

		// Each new page of jobs/get response includes the next page of the tasks, job_clusters, job_parameters, and environments.
		job.Settings.Tasks = append(job.Settings.Tasks, nextJob.Settings.Tasks...)
		job.Settings.JobClusters = append(job.Settings.JobClusters, nextJob.Settings.JobClusters...)
		job.Settings.Parameters = append(job.Settings.Parameters, nextJob.Settings.Parameters...)
		job.Settings.Environments = append(job.Settings.Environments, nextJob.Settings.Environments...)
		job.NextPageToken = nextJob.NextPageToken
	}

	return job, nil
}
