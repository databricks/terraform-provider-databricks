package sql

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceWarehouse() *schema.Resource {
	type SQLWarehouseInfo struct {
		ID                      string          `json:"id,omitempty" tf:"computed"`
		Name                    string          `json:"name,omitempty" tf:"computed"`
		ClusterSize             string          `json:"cluster_size,omitempty" tf:"computed"`
		AutoStopMinutes         int             `json:"auto_stop_mins,omitempty" tf:"computed"`
		MinNumClusters          int             `json:"min_num_clusters,omitempty" tf:"computed"`
		MaxNumClusters          int             `json:"max_num_clusters,omitempty" tf:"computed"`
		NumClusters             int             `json:"num_clusters,omitempty" tf:"computed"`
		EnablePhoton            bool            `json:"enable_photon,omitempty" tf:"computed"`
		EnableServerlessCompute bool            `json:"enable_serverless_compute,omitempty" tf:"computed"`
		InstanceProfileARN      string          `json:"instance_profile_arn,omitempty" tf:"computed"`
		State                   string          `json:"state,omitempty" tf:"computed"`
		JdbcURL                 string          `json:"jdbc_url,omitempty" tf:"computed"`
		OdbcParams              *OdbcParams     `json:"odbc_params,omitempty" tf:"computed"`
		Tags                    *Tags           `json:"tags,omitempty" tf:"computed"`
		SpotInstancePolicy      string          `json:"spot_instance_policy,omitempty" tf:"computed"`
		Channel                 *ReleaseChannel `json:"channel,omitempty" tf:"computed"`
		DataSourceID            string          `json:"data_source_id,omitempty" tf:"computed"`
	}

	return common.DataResource(SQLWarehouseInfo{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*SQLWarehouseInfo)
		var id string
		endpointsAPI := NewSQLEndpointsAPI(ctx, c)
		if data.Name != "" {
			endpoints, err := endpointsAPI.List()
			if err != nil {
				return err
			}
			selected := []SQLEndpoint{}
			for _, endpoint := range endpoints.Endpoints {
				if endpoint.Name == data.Name {
					selected = append(selected, endpoint)
				}
			}
			if len(selected) == 0 {
				return fmt.Errorf("can't find SQL warehouse with the name '%s'", data.Name)
			}
			if len(selected) > 1 {
				return fmt.Errorf("there are multiple SQL warehouses with the name '%s'", data.Name)
			}
			id = selected[0].ID
		} else if data.ID != "" {
			id = data.ID
		} else {
			return fmt.Errorf("either 'id' or 'name' should be provided")
		}
		err := c.Get(ctx, fmt.Sprintf("/sql/warehouses/%s", id), nil, data)
		if err != nil {
			return err
		}
		data.DataSourceID, err = endpointsAPI.ResolveDataSourceID(data.ID)
		if err != nil {
			return err
		}
		return nil
	})
}
