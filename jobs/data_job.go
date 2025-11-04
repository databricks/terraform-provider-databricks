package jobs

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceJob() common.Resource {
	type queryableJobData struct {
		Id      string `json:"id,omitempty" tf:"computed"`
		JobId   string `json:"job_id,omitempty" tf:"computed"`
		Name    string `json:"name,omitempty" tf:"computed"`
		JobName string `json:"job_name,omitempty" tf:"computed"`
		Job     *Job   `json:"job_settings,omitempty" tf:"computed"`
	}
	s := common.StructToSchema(queryableJobData{}, nil)
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)
	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var data queryableJobData
			common.DataToStructPointer(d, s, &data)
			jobsAPI := NewJobsAPI(ctx, newClient)
			var list []Job
			if data.Id == "" {
				data.Id = data.JobId
			}
			if data.Name == "" {
				data.Name = data.JobName
			}
			if data.Name != "" {
				// if name is provided, need to list all jobs ny name
				list, err = jobsAPI.ListByName(data.Name, true)
			} else {
				// otherwise, just read the job
				var job Job
				job, err = jobsAPI.Read(data.Id)
				if err != nil {
					return err
				}
				data.Job = &job
				data.Name = job.Settings.Name
			}
			if err != nil {
				return err
			}
			for _, job := range list {
				currentJob := job // De-referencing the temp variable used by the loop
				currentJobId := currentJob.ID()
				currentJobName := currentJob.Settings.Name
				if currentJobName == data.Name || currentJobId == data.Id {
					data.Job = &currentJob
					data.Name = currentJobName
					data.Id = currentJobId
					data.JobId = currentJobId
					break
				}
			}
			if data.Job == nil {
				return fmt.Errorf("no job found with specified name")
			}
			err = common.StructToData(data, s, d)
			if err != nil {
				return err
			}
			if data.Id != "" {
				d.SetId(data.Id)
			} else {
				d.SetId("_")
			}
			return nil
		},
	}
}
