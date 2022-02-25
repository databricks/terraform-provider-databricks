package jobs

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceJobs() *schema.Resource {
	var response struct {
		Ids map[string]string `json:"ids,omitempty" tf:"computed"`
	}
	return common.DataResource(&response, func(ctx context.Context, c *common.DatabricksClient) error {
		jobsAPI := NewJobsAPI(ctx, c)
		list, err := jobsAPI.List()
		if err != nil {
			return err
		}
		response.Ids = map[string]string{}
		for _, v := range list.Jobs {
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
