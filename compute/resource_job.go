package compute

import (
	"context"
	"fmt"
	"log"
	"sort"
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
	client := m.(*common.DatabricksClient)
	if client.UseMutiltaskJobs {
		ctx = context.WithValue(ctx, common.Api, common.API_2_1)
	}
	return JobsAPI{client, ctx}
}

// JobsAPI exposes the Jobs API
type JobsAPI struct {
	client  *common.DatabricksClient
	context context.Context
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
func (a JobsAPI) RunsCancel(runID int64, timeout time.Duration) error {
	var response interface{}
	err := a.client.Post(a.context, "/jobs/runs/cancel", map[string]interface{}{
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
	err := a.client.Get(a.context, "/jobs/runs/get", map[string]interface{}{
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

func (a JobsAPI) Restart(id string, timeout time.Duration) error {
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
		return a.Start(jobID, timeout)
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
	return a.Start(jobID, timeout)
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

var jobSchema = common.StructToSchema(JobSettings{},
	func(s map[string]*schema.Schema) map[string]*schema.Schema {
		jobSettingsSchema(&s, "")
		jobSettingsSchema(&s["task"].Elem.(*schema.Resource).Schema, "task.0.")
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
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(DefaultProvisionTimeout),
			Update: schema.DefaultTimeout(DefaultProvisionTimeout),
		},
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			var js JobSettings
			err := common.DiffToStructPointer(d, jobSchema, &js)
			if err != nil {
				return err
			}
			alwaysRunning := d.Get("always_running").(bool)
			if alwaysRunning && js.MaxConcurrentRuns > 1 {
				return fmt.Errorf("`always_running` must be specified only with `max_concurrent_runs = 1`")
			}
			c := m.(*common.DatabricksClient)
			if c.UseMutiltaskJobs {
				for _, task := range js.Tasks {
					err = validateClusterDefinition(*task.NewCluster)
					if err != nil {
						return fmt.Errorf("task %s invalid: %w", task.TaskKey, err)
					}
				}
			}
			if js.NewCluster != nil {
				err = validateClusterDefinition(*js.NewCluster)
				if err != nil {
					return fmt.Errorf("invalid job cluster: %w", err)
				}
			}
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var js JobSettings
			err := common.DataToStructPointer(d, jobSchema, &js)
			if err != nil {
				return err
			}
			sort.Slice(js.Tasks, func(i, j int) bool {
				return js.Tasks[i].TaskKey < js.Tasks[j].TaskKey
			})
			jobsAPI := NewJobsAPI(ctx, c)
			job, err := jobsAPI.Create(js)
			if err != nil {
				return err
			}
			d.SetId(job.ID())
			if d.Get("always_running").(bool) {
				// TODO: test this with c.UseMutiltaskJobs
				return jobsAPI.Start(job.JobID, d.Timeout(schema.TimeoutCreate))
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			job, err := NewJobsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			sort.Slice(job.Settings.Tasks, func(i, j int) bool {
				return job.Settings.Tasks[i].TaskKey < job.Settings.Tasks[j].TaskKey
			})
			d.Set("url", c.FormatURL("#job/", d.Id()))
			return common.StructToData(*job.Settings, jobSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var js JobSettings
			err := common.DataToStructPointer(d, jobSchema, &js)
			if err != nil {
				return err
			}
			jobsAPI := NewJobsAPI(ctx, c)
			err = jobsAPI.Update(d.Id(), js)
			if err != nil {
				return err
			}
			if d.Get("always_running").(bool) {
				// TODO: test this with c.UseMutiltaskJobs
				return jobsAPI.Restart(d.Id(), d.Timeout(schema.TimeoutUpdate))
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewJobsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
