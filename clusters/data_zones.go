package clusters

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceClusterZones ...
func DataSourceClusterZones() common.Resource {
	return common.Resource{
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			zonesInfo, err := NewClustersAPI(ctx, m).ListZones()
			if err != nil {
				return err
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
