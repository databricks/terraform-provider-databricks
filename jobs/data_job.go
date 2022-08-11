package jobs

import (
	"context"
	"fmt"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceQueryableJob() *schema.Resource {
	type queryableJobData struct {
		Id      string `json:"job_id,omitempty"`
		Name    string `json:"job_name,omitempty"`
		JobInfo *Job   `json:"job_info,omitempty" tf:"computed"`
	}
	return common.DataResource(queryableJobData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*queryableJobData)
		jobsAPI := NewJobsAPI(ctx, c)
		list, err := jobsAPI.List()
		if err != nil {
			return err
		}
		for _, currentJob := range list.Jobs {
			if currentJob.Settings.Name == data.Name || currentJob.ID() == data.Id {
				data.JobInfo = &currentJob
			}
		}
		if data.JobInfo == nil {
			return fmt.Errorf("no job found with specified name or id")
		}
		return nil
	})
}
