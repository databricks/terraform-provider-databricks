package jobs

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceQueryableJob() *schema.Resource {
	type queryableJobData struct {
		// data resources output fields annotated with tf:"computed"
		Ids []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.DataResource(queryableJobData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*queryableJobData)
		// TODO: implement me
		data.Ids = append(data.Ids, "..") // replace
		return nil
	})
}
