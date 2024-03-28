package clusters

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type LibraryResource struct {
	compute.Library
}

func (LibraryResource) CustomizeSchema(s map[string]*schema.Schema, path []string) map[string]*schema.Schema {
	common.CustomizeSchemaPath(s).AddNewField("cluster_id", &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	})
	return s
}

func ResourceLibrary() common.Resource {
	libraySdkSchema := common.StructToSchema(LibraryResource{}, nil)
	parseId := func(id string) (string, string) {
		split := strings.SplitN(id, "/", 2)
		if len(split) != 2 {
			return "unknown", "unknown"
		}
		return split[0], split[1]
	}
	return common.Resource{
		Schema: libraySdkSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			clusterID := d.Get("cluster_id").(string)
			_, err = StartClusterAndGetInfo(ctx, w, clusterID)
			if err != nil {
				return err
			}
			var lib compute.Library
			common.DataToStructPointer(d, libraySdkSchema, &lib)
			err = w.Libraries.Install(ctx, compute.InstallLibraries{
				ClusterId: clusterID,
				Libraries: []compute.Library{lib},
			})
			if err != nil {
				return err
			}
			_, err = libraries.WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
				ClusterID: clusterID,
				IsRunning: true,
			}, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				return err
			}
			d.SetId(fmt.Sprintf("%s/%s", clusterID, lib.String()))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			clusterID, libraryRep := parseId(d.Id())
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			cll, err := libraries.WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
				ClusterID: clusterID,
				IsRefresh: true,
			}, d.Timeout(schema.TimeoutRead))
			if err != nil {
				return err
			}
			for _, v := range cll.LibraryStatuses {
				thisRep := v.Library.String()
				if thisRep == libraryRep {
					common.StructToData(v.Library, libraySdkSchema, d)
					d.Set("cluster_id", clusterID)
					return nil
				}
			}
			return apierr.NotFound(fmt.Sprintf("cannot find %s on %s", libraryRep, clusterID))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			clusterID, libraryRep := parseId(d.Id())
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			_, err = StartClusterAndGetInfo(ctx, w, clusterID)
			if err != nil {
				return err
			}
			cll, err := w.Libraries.ClusterStatusByClusterId(ctx, clusterID)
			if err != nil {
				return err
			}
			for _, v := range cll.LibraryStatuses {
				if v.Library.String() != libraryRep {
					continue
				}
				return w.Libraries.Uninstall(ctx, compute.UninstallLibraries{
					ClusterId: clusterID,
					Libraries: []compute.Library{*v.Library},
				})
			}
			return apierr.NotFound(fmt.Sprintf("cannot find %s on %s", libraryRep, clusterID))
		},
	}
}
