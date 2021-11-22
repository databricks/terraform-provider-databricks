package clusters

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceClusterZones ...
func DataSourceClusterZones() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			zonesInfo, err := NewClustersAPI(ctx, m).ListZones()
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(zonesInfo.DefaultZone)
			d.Set("default_zone", zonesInfo.DefaultZone)
			d.Set("zones", zonesInfo.Zones)
			return nil
		},
		Schema: map[string]*schema.Schema{
			"default_zone": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"zones": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
		},
	}
}
