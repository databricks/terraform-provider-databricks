package clusters

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceCluster() common.Resource {
	type clusterData struct {
		Id          string       `json:"id,omitempty" tf:"computed"`
		ClusterId   string       `json:"cluster_id,omitempty" tf:"computed"`
		Name        string       `json:"cluster_name,omitempty" tf:"computed"`
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
			namedClusters := []ClusterInfo{}
			for _, clst := range clusters {
				cluster := clst
				if cluster.ClusterName == data.Name {
					namedClusters = append(namedClusters, cluster)
				}
			}
			if len(namedClusters) == 0 {
				return fmt.Errorf("there is no cluster with name '%s'", data.Name)
			}
			if len(namedClusters) > 1 {
				return fmt.Errorf("there is more than one cluster with name '%s'", data.Name)
			}
			data.ClusterInfo = &namedClusters[0]
		} else if data.ClusterId != "" {
			cls, err := clusterAPI.Get(data.ClusterId)
			if err != nil {
				return err
			}
			data.ClusterInfo = &cls
		} else {
			return fmt.Errorf("you need to specify either `cluster_name` or `cluster_id`")
		}
		data.Id = data.ClusterInfo.ClusterID
		data.ClusterId = data.ClusterInfo.ClusterID

		return nil
	})
}
