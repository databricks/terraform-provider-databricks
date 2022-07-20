package sql

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceWarehouse() *schema.Resource {
	type SQLWarehouseInfo struct {
		ID                      string          `json:"id"`
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
		err := c.Get(ctx, fmt.Sprintf("/sql/warehouses/%s", data.ID), nil, data)
		if err != nil {
			return err
		}
		endpointsAPI := NewSQLEndpointsAPI(ctx, c)
		data.DataSourceID, err = endpointsAPI.ResolveDataSourceID(data.ID)
		if err != nil {
			return err
		}
		return nil
	})
}
