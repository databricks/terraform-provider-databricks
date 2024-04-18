package clusters

import (
	"context"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceClusters() common.Resource {
	return common.Resource{
		Read: func(ctx context.Context, d *schema.ResourceData, i *common.DatabricksClient) error {
			clusters, err := NewClustersAPI(ctx, i).List()
			if err != nil {
				return err
			}
			ids := schema.NewSet(schema.HashString, []any{})
			name_contains := strings.ToLower(d.Get("cluster_name_contains").(string))
			for _, v := range clusters {
				match_name := strings.Contains(strings.ToLower(v.ClusterName), name_contains)
				if name_contains != "" && !match_name {
					continue
				}
				ids.Add(v.ClusterID)
			}
			d.Set("ids", ids)
			d.SetId("_")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"ids": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cluster_name_contains": {
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}
