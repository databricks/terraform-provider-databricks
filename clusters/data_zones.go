package clusters

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
)

// DataSourceClusterZones ...
func DataSourceClusterZones() common.Resource {
	return common.WorkspaceDataWithUnifiedProvider(func(ctx context.Context, data *struct {
		Id          string   `json:"id,omitempty" tf:"computed"`
		DefaultZone string   `json:"default_zone,omitempty" tf:"computed"`
		Zones       []string `json:"zones,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		zonesInfo, err := w.Clusters.ListZones(ctx)
		if err != nil {
			return err
		}
		data.Id = zonesInfo.DefaultZone
		data.DefaultZone = zonesInfo.DefaultZone
		data.Zones = zonesInfo.Zones
		return nil
	})
}
