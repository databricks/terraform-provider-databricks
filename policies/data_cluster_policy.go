package policies

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getPolicy(ClusterPoliciesAPI ClusterPoliciesAPI, name string) (*ClusterPolicy, error) {
	policyList, err := ClusterPoliciesAPI.List()
	if err != nil {
		return nil, err
	}
	for _, policy := range policyList {
		if policy.Name == name {
			return &policy, nil
		}
	}

	return nil, fmt.Errorf("cluster policy '%s' wasn't found", name)
}

// DataSourceClusterPolicy returns information about cluster policy specified by name
func DataSourceClusterPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"definition": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_clusters_per_user": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
			ClusterPoliciesAPI := NewClusterPoliciesAPI(ctx, m)
			policy, err := getPolicy(ClusterPoliciesAPI, d.Get("name").(string))
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(policy.PolicyID)
			d.Set("definition", policy.Definition)
			d.Set("max_clusters_per_user", policy.MaxClustersPerUser)
			return nil
		},
	}
}
