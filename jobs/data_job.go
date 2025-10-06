package jobs

import (
	"context"
	"fmt"
	"strconv"

	"github.com/databricks/databricks-sdk-go"
	sdk_jobs "github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceJob() common.Resource {
	type queryableJobData struct {
		Id      string        `json:"id,omitempty" tf:"computed"`
		JobId   string        `json:"job_id,omitempty" tf:"computed"`
		Name    string        `json:"name,omitempty" tf:"computed"`
		JobName string        `json:"job_name,omitempty" tf:"computed"`
		Job     *sdk_jobs.Job `json:"job_settings,omitempty" tf:"computed"`
	}
	return common.WorkspaceDataWithCustomizeFunc(func(ctx context.Context, data *queryableJobData, w *databricks.WorkspaceClient) error {
		var err error
		if data.Id == "" {
			data.Id = data.JobId
		}
		if data.Name == "" {
			data.Name = data.JobName
		}
		if data.Name != "" {
			// if name is provided, need to list all jobs ny name
			// find only job ID, then use `GetByJobId` to get the job with all fields
			list, err := w.Jobs.ListAll(ctx, sdk_jobs.ListJobsRequest{ExpandTasks: false, Name: data.Name, Limit: 100})
			if err != nil {
				return err
			}
			found := false
			for _, job := range list {
				currentJob := job // De-referencing the temp variable used by the loop
				currentJobName := currentJob.Settings.Name
				jobIdString := strconv.FormatInt(currentJob.JobId, 10)
				if currentJobName == data.Name || jobIdString == data.Id {
					data.Name = currentJobName
					data.Id = jobIdString
					found = true
					break // break the loop after we found the job
				}
			}
			if !found {
				return fmt.Errorf("no job found with specified name")
			}
		}
		// read the job by ID
		jobId, err := strconv.ParseInt(data.Id, 10, 64)
		if err != nil {
			return err
		}
		job, err := w.Jobs.Get(ctx, sdk_jobs.GetJobRequest{JobId: jobId})
		if err != nil {
			return err
		}
		data.Job = job
		data.Name = job.Settings.Name
		data.JobName = job.Settings.Name
		data.JobId = data.Id

		return nil
	}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// Customize the git_source schema to fix ConflictsWith paths
		if gitSource, err := common.SchemaPath(s, "job_settings", "settings", "git_source"); err == nil {
			if gitSourceResource, ok := gitSource.Elem.(*schema.Resource); ok {
				gitSourceSchema(gitSourceResource.Schema, "job_settings.0.settings.0.")
			}
		}
		return s
	})
}
