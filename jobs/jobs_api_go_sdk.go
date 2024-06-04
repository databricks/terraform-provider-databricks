package jobs

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/repos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (js *JobSettingsResource) adjustTasks() {
	js.sortTasksByKey()
	for _, task := range js.Tasks {
		sort.Slice(task.DependsOn, func(i, j int) bool {
			return task.DependsOn[i].TaskKey < task.DependsOn[j].TaskKey
		})
		sortWebhookNotifications(task.WebhookNotifications)
	}
}

func (js *JobSettingsResource) sortTasksByKey() {
	sort.Slice(js.Tasks, func(i, j int) bool {
		return js.Tasks[i].TaskKey < js.Tasks[j].TaskKey
	})
}

func adjustTasks(cj *jobs.CreateJob) {
	sortTasksByKey(cj)
	for _, task := range cj.Tasks {
		sort.Slice(task.DependsOn, func(i, j int) bool {
			return task.DependsOn[i].TaskKey < task.DependsOn[j].TaskKey
		})
		sortWebhookNotifications(task.WebhookNotifications)
	}
}

func sortTasksByKey(cj *jobs.CreateJob) {
	sort.Slice(cj.Tasks, func(i, j int) bool {
		return cj.Tasks[i].TaskKey < cj.Tasks[j].TaskKey
	})
}

func (js *JobSettingsResource) sortWebhooksByID() {
	sortWebhookNotifications(js.WebhookNotifications)
}

func sortWebhooksByID(cj *jobs.CreateJob) {
	sortWebhookNotifications(cj.WebhookNotifications)
}

func (js *JobSettingsResource) isMultiTask() bool {
	return js.Format == "MULTI_TASK" || len(js.Tasks) > 0
}

func getJobLifecycleManagerGoSdk(d *schema.ResourceData, m *common.DatabricksClient) jobLifecycleManager {
	if d.Get("always_running").(bool) {
		return alwaysRunningLifecycleManagerGoSdk{d: d, m: m}
	}
	if d.Get("control_run_state").(bool) {
		return controlRunStateLifecycleManagerGoSdk{d: d, m: m}
	}
	return noopLifecycleManager{}
}

type alwaysRunningLifecycleManagerGoSdk struct {
	d *schema.ResourceData
	m *common.DatabricksClient
}

func (a alwaysRunningLifecycleManagerGoSdk) OnCreate(ctx context.Context) error {
	w, err := a.m.WorkspaceClient()
	if err != nil {
		return err
	}
	jobID, err := parseJobId(a.d.Id())
	if err != nil {
		return err
	}

	return Start(jobID, a.d.Timeout(schema.TimeoutCreate), w, ctx)
}

func (a alwaysRunningLifecycleManagerGoSdk) OnUpdate(ctx context.Context) error {
	w, err := a.m.WorkspaceClient()
	if err != nil {
		return err
	}
	jobID, err := parseJobId(a.d.Id())
	if err != nil {
		return err
	}

	err = StopActiveRun(jobID, a.d.Timeout(schema.TimeoutUpdate), w, ctx)

	if err != nil {
		return err
	}
	return Start(jobID, a.d.Timeout(schema.TimeoutUpdate), w, ctx)
}

type controlRunStateLifecycleManagerGoSdk struct {
	d *schema.ResourceData
	m *common.DatabricksClient
}

func (c controlRunStateLifecycleManagerGoSdk) OnCreate(ctx context.Context) error {
	return nil
}

func (c controlRunStateLifecycleManagerGoSdk) OnUpdate(ctx context.Context) error {
	if c.d.Get("continuous") == nil {
		return nil
	}

	jobID, err := parseJobId(c.d.Id())
	if err != nil {
		return err
	}

	w, err := c.m.WorkspaceClient()
	if err != nil {
		return err
	}

	// Only use RunNow to stop the active run if the job is unpaused.
	pauseStatus := c.d.Get("continuous.0.pause_status").(string)
	if pauseStatus == "UNPAUSED" {
		// Previously, RunNow() was not supported for continuous jobs. Now, calling RunNow()
		// on a continuous job works, cancelling the active run if there is one, and resetting
		// the exponential backoff timer. So, we try to call RunNow() first, and if it fails,
		// we call StopActiveRun() instead.
		_, err := w.Jobs.RunNow(ctx, jobs.RunNow{
			JobId: jobID,
		})

		if err == nil {
			return nil
		}

		// RunNow() returns 404 when the feature is disabled.
		var apiErr *apierr.APIError
		if errors.As(err, &apiErr) && apiErr.StatusCode != 404 {
			return err
		}
	}

	return StopActiveRun(jobID, c.d.Timeout(schema.TimeoutUpdate), w, ctx)
}

func updateJobClusterSpec(clusterSpec *compute.ClusterSpec, d *schema.ResourceData) error {
	err := clusters.ModifyRequestOnInstancePool(clusterSpec)
	if err != nil {
		return err
	}
	err = clusters.FixInstancePoolChangeIfAny(d, clusterSpec)
	if err != nil {
		return err
	}
	err = clusters.SetForceSendFieldsForCluster(clusterSpec, d)
	if err != nil {
		return err
	}
	return nil
}

func prepareJobSettingsForUpdateGoSdk(d *schema.ResourceData, js *JobSettingsResource) error {
	if js.NewCluster != nil {
		err := updateJobClusterSpec(js.NewCluster, d)
		if err != nil {
			return err
		}
	}
	for _, task := range js.Tasks {
		if task.NewCluster != nil {
			err := updateJobClusterSpec(task.NewCluster, d)
			if err != nil {
				return err
			}
		}
	}
	for i := range js.JobClusters {
		err := updateJobClusterSpec(&js.JobClusters[i].NewCluster, d)
		if err != nil {
			return err
		}
	}
	return nil
}

func prepareJobSettingsForCreateGoSdk(d *schema.ResourceData, jc *JobCreateStruct) error {
	// We always need to add NumWorkers into ForceSendField for the go-sdk client.
	// Before the go-sdk migration, the field `num_workers` was required, so we always sent it.
	for _, task := range jc.Tasks {
		if task.NewCluster != nil {
			err := clusters.SetForceSendFieldsForCluster(task.NewCluster, d)
			if err != nil {
				return err
			}
		}
	}
	for i := range jc.JobClusters {
		err := clusters.SetForceSendFieldsForCluster(&jc.JobClusters[i].NewCluster, d)
		if err != nil {
			return err
		}
	}
	return nil
}

func Create(createJob jobs.CreateJob, w *databricks.WorkspaceClient, ctx context.Context) (int64, error) {
	adjustTasks(&createJob)
	sortWebhooksByID(&createJob)
	var gitSource *jobs.GitSource = createJob.GitSource
	if gitSource != nil && gitSource.GitProvider == "" {
		var provider jobs.GitProvider = jobs.GitProvider(repos.GetGitProviderFromUrl(gitSource.GitUrl))
		gitSource.GitProvider = provider
		if gitSource.GitProvider == "" {
			return 0, fmt.Errorf("git source is not empty but Git Provider is not specified and cannot be guessed by url %+v", gitSource)
		}
		if gitSource.GitBranch == "" && gitSource.GitTag == "" && gitSource.GitCommit == "" {
			return 0, fmt.Errorf("git source is not empty but none of branch, commit and tag is specified")
		}
	}
	res, err := w.Jobs.Create(ctx, createJob)
	return res.JobId, err
}

func Update(jobID int64, js JobSettingsResource, w *databricks.WorkspaceClient, ctx context.Context) error {
	err := w.Jobs.Reset(ctx, jobs.ResetJob{
		JobId:       jobID,
		NewSettings: js.JobSettings,
	})
	return wrapMissingJobError(err, fmt.Sprintf("%d", jobID))
}

func Read(jobID int64, w *databricks.WorkspaceClient, ctx context.Context) (job *jobs.Job, err error) {
	job, err = w.Jobs.GetByJobId(ctx, jobID)
	err = wrapMissingJobError(err, fmt.Sprintf("%d", jobID))
	if job.Settings != nil {
		js := JobSettingsResource{JobSettings: *job.Settings}
		js.adjustTasks()
		js.sortWebhooksByID()
	}

	// Populate the `run_as` field. In the settings struct it can only be set on write and is not
	// returned on read. Therefore, we populate it from the top-level `run_as_user_name` field so
	// that Terraform can still diff it with the intended state.
	if job.Settings != nil && job.RunAsUserName != "" {
		if common.StringIsUUID(job.RunAsUserName) {
			job.Settings.RunAs = &jobs.JobRunAs{
				ServicePrincipalName: job.RunAsUserName,
			}
		} else {
			job.Settings.RunAs = &jobs.JobRunAs{
				UserName: job.RunAsUserName,
			}
		}
	}

	return
}

func Start(jobID int64, timeout time.Duration, w *databricks.WorkspaceClient, ctx context.Context) error {
	res, err := w.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: jobID,
	})
	if err != nil {
		return err
	}

	_, err = res.GetWithTimeout(timeout)
	if err != nil {
		return err
	}
	return nil
}

func StopActiveRun(jobID int64, timeout time.Duration, w *databricks.WorkspaceClient, ctx context.Context) error {
	runs, err := w.Jobs.ListRunsAll(ctx, jobs.ListRunsRequest{
		JobId:      jobID,
		ActiveOnly: true,
	})
	if err != nil {
		return err
	}
	if len(runs) > 1 {
		return fmt.Errorf("`always_running` must be specified only with "+
			"`max_concurrent_runs = 1`. There are %d active runs", len(runs))
	}
	if len(runs) == 1 {
		activeRun := runs[0]
		res, err := w.Jobs.CancelRun(ctx, jobs.CancelRun{
			RunId: activeRun.RunId,
		})
		if err != nil {
			return fmt.Errorf("cannot cancel run %d: %v", activeRun.RunId, err)
		}
		_, err = res.GetWithTimeout(timeout)
		if err != nil {
			return fmt.Errorf("cannot cancel run, error waiting for run to be terminated %d: %v", activeRun.RunId, err)
		}
	}
	return nil
}
