package clusters

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceClusters() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		Id                  string                        `json:"id,omitempty" tf:"computed"`
		Ids                 []string                      `json:"ids,omitempty" tf:"computed,slice_set"`
		ClusterNameContains string                        `json:"cluster_name_contains,omitempty"`
		FilterBy            *compute.ListClustersFilterBy `json:"filter_by,omitempty"`
	}, w *databricks.WorkspaceClient,
	) error {
		clusters, err := w.Clusters.ListAll(ctx, compute.ListClustersRequest{
			FilterBy: data.FilterBy,
		})
		if err != nil {
			return err
		}
		ids := make([]string, 0, len(clusters))
		name_contains := strings.ToLower(data.ClusterNameContains)
		for _, v := range clusters {
			match_name := strings.Contains(strings.ToLower(v.ClusterName), name_contains)
			if name_contains != "" && !match_name {
				continue
			}
			ids = append(ids, v.ClusterId)
		}
		data.Ids = ids
		data.Id = "_"
		return nil
	})
}
