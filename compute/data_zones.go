package compute

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceClusterZones() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, m interface{}) error {
			zonesInfo, err := NewClustersAPI(m).ListZones()
			if err != nil {
				return err
			}
			d.SetId(zonesInfo.DefaultZone)
			err = d.Set("default_zone", zonesInfo.DefaultZone)
			if err != nil {
				return err
			}
			return d.Set("zones", zonesInfo.Zones)
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
