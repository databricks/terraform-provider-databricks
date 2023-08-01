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
		if data.ID == "" && data.Name == "" {
			return fmt.Errorf("either 'id' or 'name' should be provided")
		}
		endpointsAPI := NewSQLEndpointsAPI(ctx, c)
		selected := []DataSource{}
		dataSources, err := endpointsAPI.listDataSources()
		if err != nil {
			return err
		}
		for _, source := range dataSources {
			if data.Name != "" && source.Name == data.Name {
				selected = append(selected, source)
			} else if data.ID != "" && source.EndpointID == data.ID {
				selected = append(selected, source)
				break
			}
		}
		if len(selected) == 0 {
			if data.Name != "" {
				return fmt.Errorf("can't find SQL warehouse with the name '%s'", data.Name)
			} else {
				return fmt.Errorf("can't find SQL warehouse with the ID '%s'", data.ID)
			}
		}
		if len(selected) > 1 {
			if data.Name != "" {
				return fmt.Errorf("there are multiple SQL warehouses with the name '%s'", data.Name)
			} else {
				return fmt.Errorf("there are multiple SQL warehouses with the ID '%s'", data.ID)
			}
		}
		id = selected[0].EndpointID
		err = c.Get(ctx, fmt.Sprintf("/sql/warehouses/%s", id), nil, data)
		if err != nil {
			return err
		}
		data.DataSourceID = selected[0].ID
		return nil
	})
}
