package clusters

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceClusters() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
			clusters, err := NewClustersAPI(ctx, i).List()
			if err != nil {
				return diag.FromErr(err)
			}
			ids := schema.NewSet(schema.HashString, []interface{}{})
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
