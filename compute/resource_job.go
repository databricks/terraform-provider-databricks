package compute

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
)

// NewJobsAPI creates JobsAPI instance from provider meta
func NewJobsAPI(ctx context.Context, m interface{}) JobsAPI {
	return JobsAPI{m.(*common.DatabricksClient), context.TODO()}
}

// JobsAPI exposes the Jobs API
type JobsAPI struct {
	client  *common.DatabricksClient
	context context.Context
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

var jobSchema = internal.StructToSchema(JobSettings{},
	func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["existing_cluster_id"].Description = "If existing_cluster_id, the ID " +
			"of an existing cluster that will be used for all runs of this job. " +
			"When running jobs on an existing cluster, you may need to manually " +
			"restart the cluster if it stops responding. We strongly suggest to use " +
			"`new_cluster` for greater reliability."
		s["new_cluster"].Description = "Same set of parameters as for " +
			"[databricks_cluster](cluster.md) resource."
		if v, err := internal.SchemaPath(s, "new_cluster", "spark_conf"); err == nil {
			v.DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				isPossiblyLegacyConfig := "new_cluster.0.spark_conf.%" == k && "1" == old && "0" == new
				isLegacyConfig := "new_cluster.0.spark_conf.spark.databricks.delta.preview.enabled" == k
				if isPossiblyLegacyConfig || isLegacyConfig {
					log.Printf("[DEBUG] Suppressing diff for k=%#v old=%#v new=%#v", k, old, new)
					return true
				}
				return false
			}
		}
		s["name"].Description = "An optional name for the job. The default value is Untitled."
		s["library"].Description = "An optional list of libraries to be installed on " +
			"the cluster that will execute the job. The default value is an empty list."
		s["email_notifications"].Description = "An optional set of email addresses " +
			"notified when runs of this job begin and complete and when this job is " +
			"deleted. The default behavior is to not send any emails."
		s["timeout_seconds"].Description = "An optional timeout applied to each run " +
			"of this job. The default behavior is to have no timeout."
		s["max_retries"].Description = "An optional maximum number of times to retry " +
			"an unsuccessful run. A run is considered to be unsuccessful if it " +
			"completes with a FAILED result_state or INTERNAL_ERROR life_cycle_state. " +
			"The value -1 means to retry indefinitely and the value 0 means to never " +
			"retry. The default behavior is to never retry."
		s["min_retry_interval_millis"].Description = "An optional minimal interval in " +
			"milliseconds between the start of the failed run and the subsequent retry run. " +
			"The default behavior is that unsuccessful runs are immediately retried."
		s["retry_on_timeout"].Description = "An optional policy to specify whether to " +
			"retry a job when it times out. The default behavior is to not retry on timeout."
		s["schedule"].Description = "An optional periodic schedule for this job. " +
			"The default behavior is that the job runs when triggered by clicking " +
			"Run Now in the Jobs UI or sending an API request to runNow."
		s["max_concurrent_runs"].Description = "An optional maximum allowed number of " +
			"concurrent runs of the job."
		return s
	})

// ResourceJob ...
func ResourceJob() *schema.Resource {
	return util.CommonResource{
		Schema: jobSchema,
		SchemaVersion: 2,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var js JobSettings
			err := internal.DataToStructPointer(d, jobSchema, &js)
			if err != nil {
				return err
			}
			job, err := NewJobsAPI(ctx, c).Create(js)
			if err != nil {
				return err
			}
			d.SetId(job.ID())
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			job, err := NewJobsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			return internal.StructToData(*job.Settings, jobSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var js JobSettings
			err := internal.DataToStructPointer(d, jobSchema, &js)
			if err != nil {
				return err
			}
			return NewJobsAPI(ctx, c).Update(d.Id(), js)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewJobsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
