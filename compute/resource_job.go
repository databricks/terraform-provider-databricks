package compute

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/databrickslabs/terraform-provider-databricks/common"
)

// NewJobsAPI creates JobsAPI instance from provider meta
func NewJobsAPI(ctx context.Context, m interface{}) JobsAPI {
	return JobsAPI{m.(*common.DatabricksClient), ctx, 5 * time.Minute}
}

// JobsAPI exposes the Jobs API
type JobsAPI struct {
	client  *common.DatabricksClient
	context context.Context
	timeout time.Duration
}

// List all jobs
func (a JobsAPI) List() (l JobList, err error) {
	err = a.client.Get(a.context, "/jobs/list", nil, &l)
	return
}

// RunsList ...
func (a JobsAPI) RunsList(r JobRunsListRequest) (jrl JobRunsList, err error) {
	err = a.client.Get(a.context, "/jobs/runs/list", r, &jrl)
	return
}

// RunsCancel ...
func (a JobsAPI) RunsCancel(runID int64) error {
	var response interface{}
	err := a.client.Post(a.context, "/jobs/runs/cancel", map[string]interface{}{
		"run_id": runID,
	}, &response)
	if err != nil {
		return err
	}
	return a.waitForRunState(runID, "TERMINATED")
}

func (a JobsAPI) waitForRunState(runID int64, desiredState string) error {
	return resource.RetryContext(a.context, a.timeout, func() *resource.RetryError {
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
	err := a.client.Get(a.context, "/jobs/runs/get", map[string]interface{}{
		"run_id": runID,
	}, &jr)
	return jr, err
}

func (a JobsAPI) Run(jobID int64) error {
	runID, err := a.RunNow(jobID)
	if err != nil {
		return fmt.Errorf("cannot start job run: %v", err)
	}
	return a.waitForRunState(runID, "RUNNING")
}

func (a JobsAPI) Restart(id string) error {
	jobID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}
	runs, err := a.RunsList(JobRunsListRequest{JobID: jobID, ActiveOnly: true})
	if err != nil {
		return err
	}
	if len(runs.Runs) == 0 {
		// nothing to cancel
		return a.Run(jobID)
	}
	if len(runs.Runs) > 1 {
		return fmt.Errorf("`always_running` must be specified only with "+
			"`max_concurrent_runs = 1`. There are %d active runs", len(runs.Runs))
	}
	if len(runs.Runs) == 1 {
		activeRun := runs.Runs[0]
		err = a.RunsCancel(activeRun.RunID)
		if err != nil {
			return fmt.Errorf("cannot cancel run %d: %v", activeRun.RunID, err)
		}
	}
	return a.Run(jobID)
}

// Create creates a job on the workspace given the job settings
func (a JobsAPI) Create(jobSettings JobSettings) (Job, error) {
	var job Job
	err := a.client.Post(a.context, "/jobs/create", jobSettings, &job)
	return job, err
}

// Update updates a job given the id and a new set of job settings
func (a JobsAPI) Update(id string, jobSettings JobSettings) error {
	jobID, err := strconv.ParseInt(id, 10, 32)
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
	jobID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return
	}
	err = wrapMissingJobError(a.client.Get(a.context, "/jobs/get", map[string]int64{
		"job_id": jobID,
	}, &job), id)
	return
}

// Delete deletes the job given a job id
func (a JobsAPI) Delete(id string) error {
	jobID, err := strconv.ParseInt(id, 10, 32)
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
	apiErr, ok := err.(common.APIError)
	if !ok {
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

var jobSchema = common.StructToSchema(JobSettings{},
	func(s map[string]*schema.Schema) map[string]*schema.Schema {
		if p, err := common.SchemaPath(s, "new_cluster", "num_workers"); err == nil {
			p.Optional = true
			p.Default = 0
			p.Type = schema.TypeInt
			p.ValidateDiagFunc = validation.ToDiagFunc(validation.IntAtLeast(0))
			p.Required = false
		}
		if p, err := common.SchemaPath(s, "schedule", "pause_status"); err == nil {
			p.ValidateFunc = validation.StringInSlice([]string{"PAUSED", "UNPAUSED"}, false)
		}
		if v, err := common.SchemaPath(s, "new_cluster", "spark_conf"); err == nil {
			v.DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				isPossiblyLegacyConfig := k == "new_cluster.0.spark_conf.%" && old == "1" && new == "0"
				isLegacyConfig := k == "new_cluster.0.spark_conf.spark.databricks.delta.preview.enabled"
				if isPossiblyLegacyConfig || isLegacyConfig {
					log.Printf("[DEBUG] Suppressing diff for k=%#v old=%#v new=%#v", k, old, new)
					return true
				}
				return false
			}
		}
		if v, err := common.SchemaPath(s, "new_cluster", "aws_attributes"); err == nil {
			v.DiffSuppressFunc = common.MakeEmptyBlockSuppressFunc("new_cluster.0.aws_attributes.#")
		}
		if v, err := common.SchemaPath(s, "new_cluster", "azure_attributes"); err == nil {
			v.DiffSuppressFunc = common.MakeEmptyBlockSuppressFunc("new_cluster.0.azure_attributes.#")
		}
		if v, err := common.SchemaPath(s, "new_cluster", "gcp_attributes"); err == nil {
			v.DiffSuppressFunc = common.MakeEmptyBlockSuppressFunc("new_cluster.0.gcp_attributes.#")
		}
		s["email_notifications"].DiffSuppressFunc = common.MakeEmptyBlockSuppressFunc("email_notifications.#")
		s["url"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		s["always_running"] = &schema.Schema{
			Optional: true,
			Default:  false,
			Type:     schema.TypeBool,
		}
		return s
	})

// ResourceJob ...
func ResourceJob() *schema.Resource {
	return common.Resource{
		Schema:        jobSchema,
		SchemaVersion: 2,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c interface{}) error {
			alwaysRunning := d.Get("always_running").(bool)
			maxConcurrentRuns := d.Get("max_concurrent_runs").(int)
			if alwaysRunning && maxConcurrentRuns > 1 {
				return fmt.Errorf("`always_running` must be specified only with `max_concurrent_runs = 1`")
			}
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var js JobSettings
			err := common.DataToStructPointer(d, jobSchema, &js)
			if err != nil {
				return err
			}
			if js.NewCluster != nil {
				if err = validateClusterDefinition(*js.NewCluster); err != nil {
					return err
				}
			}
			jobsAPI := NewJobsAPI(ctx, c)
			job, err := jobsAPI.Create(js)
			if err != nil {
				return err
			}
			d.SetId(job.ID())
			if d.Get("always_running").(bool) {
				return jobsAPI.Run(job.JobID)
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			job, err := NewJobsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			d.Set("url", c.FormatURL("#job/", d.Id()))
			return common.StructToData(*job.Settings, jobSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var js JobSettings
			err := common.DataToStructPointer(d, jobSchema, &js)
			if err != nil {
				return err
			}
			if js.NewCluster != nil {
				err = validateClusterDefinition(*js.NewCluster)
				if err != nil {
					return err
				}
			}
			jobsAPI := NewJobsAPI(ctx, c)
			err = jobsAPI.Update(d.Id(), js)
			if err != nil {
				return err
			}
			if d.Get("always_running").(bool) {
				return jobsAPI.Restart(d.Id())
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewJobsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
