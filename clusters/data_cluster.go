package clusters

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceCluster() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		common.Namespace
		Id          string                  `json:"id,omitempty" tf:"computed"`
		ClusterId   string                  `json:"cluster_id,omitempty" tf:"computed"`
		Name        string                  `json:"cluster_name,omitempty" tf:"computed"`
		ClusterInfo *compute.ClusterDetails `json:"cluster_info,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		if data.Name != "" {
			clusters, err := w.Clusters.ListAll(ctx, compute.ListClustersRequest{})
			if err != nil {
				return err
			}
			namedClusters := []compute.ClusterDetails{}
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
			cls, err := w.Clusters.GetByClusterId(ctx, data.ClusterId)
			if err != nil {
				return err
			}
			data.ClusterInfo = cls
		} else {
			return fmt.Errorf("you need to specify either `cluster_name` or `cluster_id`")
		}
		data.Id = data.ClusterInfo.ClusterId
		data.ClusterId = data.ClusterInfo.ClusterId
		data.Name = data.ClusterInfo.ClusterName

		return nil
	})
}
