package clusters

import (
	"context"
	"fmt"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/libraries"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceLibrary() *schema.Resource {
	s := common.StructToSchema(libraries.Library{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["cluster_id"] = &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		}
		return m
	})
	parseId := func(id string) (string, string) {
		split := strings.SplitN(id, "/", 2)
		if len(split) != 2 {
			return "unknown", "unknown"
		}
		return split[0], split[1]
	}
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			clusterID := d.Get("cluster_id").(string)
			err := NewClustersAPI(ctx, c).Start(clusterID)
			if err != nil {
				return err
			}
			var lib libraries.Library
			common.DataToStructPointer(d, s, &lib)
			librariesAPI := libraries.NewLibrariesAPI(ctx, c)
			err = librariesAPI.Install(libraries.ClusterLibraryList{
				ClusterID: clusterID,
				Libraries: []libraries.Library{lib},
			})
			if err != nil {
				return err
			}
			_, err = librariesAPI.WaitForLibrariesInstalled(libraries.Wait{
				ClusterID: clusterID,
				Timeout:   d.Timeout(schema.TimeoutCreate),
				IsRunning: true,
			})
			if err != nil {
				return err
			}
			d.SetId(fmt.Sprintf("%s/%s", clusterID, lib.String()))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			clusterID, libraryRep := parseId(d.Id())
			cll, err := libraries.NewLibrariesAPI(ctx, c).WaitForLibrariesInstalled(libraries.Wait{
				ClusterID: clusterID,
				Timeout:   d.Timeout(schema.TimeoutRead),
				IsRefresh: true,
			})
			if err != nil {
				return err
			}
			for _, v := range cll.LibraryStatuses {
				thisRep := v.Library.String()
				if thisRep == libraryRep {
					// library is found
					return nil
				}
			}
			return common.NotFound(fmt.Sprintf("cannot find %s on %s", libraryRep, clusterID))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			clusterID, libraryRep := parseId(d.Id())
			err := NewClustersAPI(ctx, c).Start(clusterID)
			if err != nil {
				return err
			}
			librariesAPI := libraries.NewLibrariesAPI(ctx, c)
			cll, err := librariesAPI.ClusterStatus(clusterID)
			if err != nil {
				return err
			}
			for _, v := range cll.LibraryStatuses {
				if v.Library.String() != libraryRep {
					continue
				}
				return librariesAPI.Uninstall(libraries.ClusterLibraryList{
					ClusterID: clusterID,
					Libraries: []libraries.Library{*v.Library},
				})
			}
			return common.NotFound(fmt.Sprintf("cannot find %s on %s", libraryRep, clusterID))
		},
	}.ToResource()
}
