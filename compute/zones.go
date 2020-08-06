package compute

import (
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceClusterZones() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, m interface{}) error {
			client := m.(*common.DatabricksClient)

			zonesInfo, err := NewClustersAPI(client).ListZones()
			if err != nil {
				return err
			}

			d.SetId(zonesInfo.DefaultZone)
			err = d.Set("default_zone", zonesInfo.DefaultZone)
			if err != nil {
				return err
			}
			err = d.Set("zones", zonesInfo.Zones)
			return err
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
