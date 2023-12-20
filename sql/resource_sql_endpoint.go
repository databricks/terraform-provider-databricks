package sql

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ClusterSizes for SQL endpoints
var (
	ClusterSizes   = []string{"2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large"}
	MaxNumClusters = 30
)

type SqlWarehouse struct {
	sql.GetWarehouseResponse

	// The data source ID is not part of the endpoint API response.
	// We manually resolve it by retrieving the list of data sources
	// and matching this entity's endpoint ID.
	DataSourceId string `json:"data_source_id,omitempty" tf:"computed"`
}

func getSqlWarehouse(ctx context.Context, w *databricks.WorkspaceClient, id string) (SqlWarehouse, error) {
	se, err := w.Warehouses.GetById(ctx, id)
	if err != nil {
		return SqlWarehouse{}, err
	}
	warehouse := SqlWarehouse{
		GetWarehouseResponse: *se,
	}
	return warehouse, nil
}

func resolveDataSourceID(ctx context.Context, w *databricks.WorkspaceClient, warehouseId string) (string, error) {
	list, err := w.DataSources.List(ctx)
	if err != nil {
		return "", err
	}
	for _, ds := range list {
		if ds.WarehouseId == warehouseId {
			return ds.Id, nil
		}
	}
	return "", fmt.Errorf("no data source found for endpoint %s", warehouseId)
}

func setRequired(s *schema.Schema) {
	s.Optional = false
	s.Required = true
}

func ResourceSqlEndpoint() *schema.Resource {
	s := common.StructToSchema(SqlWarehouse{}, func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		m["id"].Computed = true
		common.SetDefault(m["auto_stop_mins"], 120)
		common.SetSuppressDiff(m["channel"])
		common.MustSchemaPath(m, "channel", "name").Default = "CHANNEL_NAME_CURRENT"
		common.SetReadOnly(m["creator_name"])
		m["cluster_size"].ValidateDiagFunc = validation.ToDiagFunc(
			validation.StringInSlice(ClusterSizes, false))
		common.SetDefault(m["enable_photon"], true)
		common.SetSuppressDiff(m["enable_serverless_compute"])
		common.SetReadOnly(m["health"])
		common.SetReadOnly(m["jdbc_url"])
		common.SetDefault(m["max_num_clusters"], 1)
		m["max_num_clusters"].ValidateDiagFunc = validation.ToDiagFunc(
			validation.IntBetween(1, MaxNumClusters))
		common.SetSuppressDiff(m["min_num_clusters"])
		common.SetReadOnly(m["num_active_sessions"])
		common.SetSuppressDiff(m["num_clusters"])
		common.SetReadOnly(m["odbc_params"])
		common.SetDefault(m["spot_instance_policy"], "COST_OPTIMIZED")
		common.SetReadOnly(m["state"])
		common.SetSuppressDiff(m["tags"])
		common.SetSuppressDiff(m["warehouse_type"])
		m["warehouse_type"].ValidateDiagFunc = validation.ToDiagFunc(
			validation.StringInSlice([]string{"PRO", "CLASSIC"}, false))
		return m
	})
	f, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.SortKeys = true
	spew.Fdump(f, s)
	return common.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var se sql.CreateWarehouseRequest
			common.DataToStructPointer(d, s, &se)
			wait, err := w.Warehouses.Create(ctx, se)
			if err != nil {
				return fmt.Errorf("failed creating warehouse: %w", err)
			}
			resp, err := wait.Get()
			if err != nil {
				return fmt.Errorf("failed waiting for warehouse to start: %w", err)
			}
			d.SetId(resp.Id)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			warehouse, err := getSqlWarehouse(ctx, w, d.Id())
			if err != nil {
				return err
			}
			warehouse.DataSourceId, err = resolveDataSourceID(ctx, w, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(warehouse, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var se sql.EditWarehouseRequest
			common.DataToStructPointer(d, s, &se)
			se.Id = d.Id()
			_, err = w.Warehouses.Edit(ctx, se)
			if err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Warehouses.DeleteById(ctx, d.Id())
		},
		Schema: s,
	}.ToResource()
}
