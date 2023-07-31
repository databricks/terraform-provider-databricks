package jobs

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceJobs() *schema.Resource {
	type jobsData struct {
		Ids map[string]string `json:"ids,omitempty" tf:"computed"`
	}
	return common.DataResource(jobsData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		response := e.(*jobsData)
		jobsAPI := NewJobsAPI(ctx, c)
		list, err := jobsAPI.List()
		if err != nil {
			return err
		}
		response.Ids = map[string]string{}
		for _, v := range list {
			name := v.Settings.Name
			_, duplicateName := response.Ids[name]
			if duplicateName {
				return fmt.Errorf("duplicate job name detected: %s", name)
			}
			response.Ids[name] = v.ID()
		}
		return nil
	})
}
