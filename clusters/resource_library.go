package clusters

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type LibraryResource struct {
	compute.Library
	common.Namespace
}

func (LibraryResource) CustomizeSchemaResourceSpecific(s *common.CustomizableSchema) *common.CustomizableSchema {
	s.AddNewField("cluster_id", &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	})
	return s
}

const EggDeprecationWarning = "The `egg` library type is deprecated. Please use `whl` or `pypi` instead."

func (LibraryResource) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	common.NamespaceCustomizeSchema(s)
	s.SchemaPath("egg").SetDeprecated(EggDeprecationWarning)
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
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			cll, err := libraries.WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
				ClusterID: clusterID,
				IsRefresh: true,
			}, d.Timeout(schema.TimeoutRead))
			if err != nil {
				err = common.IgnoreNotFoundError(err)
				if err == nil {
					log.Printf("[WARN] %s is not found, ignoring it", clusterID)
					d.SetId("")
				}
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
			return &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    fmt.Sprintf("cannot find %s on %s", libraryRep, clusterID),
			}
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			clusterID, libraryRep := parseId(d.Id())
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
			return &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    fmt.Sprintf("cannot find %s on %s", libraryRep, clusterID),
			}
		},
	}
}
