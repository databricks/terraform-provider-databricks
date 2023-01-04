package clusters

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCluster() *schema.Resource {
	type clusterData struct {
		Id          string       `json:"id,omitempty" tf:"computed"`
		ClusterId   string       `json:"cluster_id,omitempty" tf:"computed"`
		Name        string       `json:"name,omitempty" tf:"computed"`
		ClusterInfo *ClusterInfo `json:"cluster_info,omitempty" tf:"computed"`
	}
	return common.DataResource(clusterData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*clusterData)
		clusterAPI := NewClustersAPI(ctx, c)
		if data.Name != "" {
			clusters, err := clusterAPI.List()
			if err != nil {
				return err
			}
			for _, clst := range clusters {
				cluster := clst
				if cluster.ClusterName == data.Name {
					data.ClusterInfo = &cluster
					break
				}
			}
			if data.ClusterInfo == nil {
				return fmt.Errorf("there is no cluster with name '%s'", data.Name)
			}
		} else if data.ClusterId != "" {
			cls, err := clusterAPI.Get(data.ClusterId)
			if err != nil {
				return err
			}
			data.ClusterInfo = &cls
		} else {
			return fmt.Errorf("you need to specify either `name` or `cluster_id`")
		}
		data.Id = data.ClusterInfo.ClusterID
		data.ClusterId = data.ClusterInfo.ClusterID

		return nil
	})
}
