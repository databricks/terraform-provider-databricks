package compute

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceClusterSparkVersions() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceClusterSparkVersionsRead,

		Schema: map[string]*schema.Schema{
			"version": {
				Type:     schema.TypeString,
				Computed: true},
			"latest": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"long_term_support": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ml": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gpu": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"scala": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2.12",
			},
			"genomics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

// DataSourceClusterSparkVersionsRead performs the Databricks runtime lookup.
func DataSourceClusterSparkVersionsRead(d *schema.ResourceData, m interface{}) error {

	sparkVersions, err := NewClustersAPI(m).ListSparkVersions()
	if err != nil {
		return err
	}
	fmt.Printf("Received %d", len(sparkVersions.SparkVerions))
	latest := d.Get("latest").(bool)
	lts := d.Get("long_term_support").(bool)
	ml := d.Get("ml").(bool)
	gpu := d.Get("gpu").(bool)
	genomics := d.Get("genomics").(bool)

	var clusterSparkVersions []string
	if scala, ok := d.GetOk("scala"); ok {
		for _, version := range sparkVersions.SparkVerions {
			if strings.Contains(version.Key, scala.(string)) {
				runtimeFilter := (strings.Contains(version.Key, "ml") == ml &&
					strings.Contains(version.Key, "gpu") == gpu &&
					strings.Contains(version.Name, "Genomics") == genomics &&
					strings.Contains(version.Name, "LTS") == lts)

				if runtimeFilter {
					clusterSparkVersions = append(clusterSparkVersions, version.Key)
				}
			}
		}
	}

	if len(clusterSparkVersions) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again")
	}

	if len(clusterSparkVersions) > 1 {
		if latest {
			sort.Sort(sort.Reverse(sort.StringSlice(clusterSparkVersions)))
		} else {
			return fmt.Errorf("Your query returned multiple results. Please change your search criteria and try again")
		}
	}

	return d.Set("version", clusterSparkVersions[0])
}
