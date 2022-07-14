package clusters

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCluster() *schema.Resource {

	return common.DataResource(ClusterInfo{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*ClusterInfo)
		var err error
		clusterAPI := NewClustersAPI(ctx, c)
		*data, err = clusterAPI.Get(data.ClusterID)
		if err != nil {
			return err
		}

		return nil
	})
}
