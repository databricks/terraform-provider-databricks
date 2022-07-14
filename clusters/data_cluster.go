package clusters

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCluster() *schema.Resource {
	type clusterData struct {
		ClusterId   string       `json:"cluster_id"`
		ClusterInfo *ClusterInfo `json:"cluster_info,omitempty" tf:"computed"`
	}
	return common.DataResource(clusterData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*clusterData)
		clusterAPI := NewClustersAPI(ctx, c)
		clusterInfo, err := clusterAPI.Get(data.ClusterId)

		data.ClusterInfo = &clusterInfo
		if err != nil {
			return err
		}

		return nil
	})
}
