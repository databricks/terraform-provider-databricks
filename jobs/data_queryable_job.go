package jobs

import (
	"context"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceQueryableJob() *schema.Resource {
	type queryableJobData struct {
		Ids  []string `json:"ids,omitempty" tf:"computed,slice_set"`
		Id   string   `json:"id,omitempty"`
		Name string   `json:"job_name,omitempty"`
	}
	return common.DataResource(queryableJobData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*queryableJobData)
		jobsAPI := NewJobsAPI(ctx, c)
		list, err := jobsAPI.List()
		if err != nil {
			return err
		}
		for _, currentJob := range list.Jobs {
			name := currentJob.Settings.Name
			id := currentJob.ID()
			if id == data.Id || name == data.Name {
				data.Ids = append(data.Ids, currentJob.ID())
			}
		}
		return nil
	})
}
