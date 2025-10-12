package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ClusterSizes for SQL endpoints
var (
	ClusterSizes    = []string{"2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large"}
	MaxNumClusters  = 30
	ForceSendFields = []string{"enable_serverless_compute", "enable_photon", "auto_stop_mins"}
)

type SqlWarehouse struct {
	sql.GetWarehouseResponse

	// The data source ID is not part of the endpoint API response.
	// We manually resolve it by retrieving the list of data sources
	// and matching this entity's endpoint ID.
	DataSourceId string `json:"data_source_id,omitempty" tf:"computed"`
}

func getSqlWarehouse(ctx context.Context, w *databricks.WorkspaceClient, id string) (*SqlWarehouse, error) {
	se, err := w.Warehouses.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	warehouse := SqlWarehouse{
		GetWarehouseResponse: *se,
	}
	return &warehouse, nil
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

func ResourceSqlEndpoint() common.Resource {
	s := common.StructToSchema(SqlWarehouse{}, func(
		m map[string]*schema.Schema) map[string]*schema.Schema {
		m["id"].Computed = true
		common.SetDefault(m["auto_stop_mins"], 120)
		common.CustomizeSchemaPath(m, "channel").SetSuppressDiff()
		common.MustSchemaPath(m, "channel", "name").Default = "CHANNEL_NAME_CURRENT"
		common.SetRequired(m["cluster_size"])
		common.SetReadOnly(m["creator_name"])
		m["cluster_size"].ValidateDiagFunc = validation.ToDiagFunc(
			validation.StringInSlice(ClusterSizes, false))
		common.SetDefault(m["enable_photon"], true)
		m["enable_serverless_compute"].Computed = true
		common.SetReadOnly(m["health"])
		common.SetReadOnly(m["jdbc_url"])
		common.SetDefault(m["max_num_clusters"], 1)
		m["max_num_clusters"].ValidateDiagFunc = validation.ToDiagFunc(
			validation.IntBetween(1, MaxNumClusters))
		common.CustomizeSchemaPath(m, "min_num_clusters").SetSuppressDiff()
		common.SetRequired(m["name"])
		common.SetReadOnly(m["num_active_sessions"])
		common.SetReadOnly(m["num_clusters"])
		common.SetReadOnly(m["odbc_params"])
		common.SetDefault(m["spot_instance_policy"], "COST_OPTIMIZED")
		common.SetReadOnly(m["state"])
		common.CustomizeSchemaPath(m, "tags").SetSuppressDiff()
		common.SetRequired(common.MustSchemaPath(m, "tags", "custom_tags", "key"))
		common.SetRequired(common.MustSchemaPath(m, "tags", "custom_tags", "value"))
		common.CustomizeSchemaPath(m, "warehouse_type").
			SetSuppressDiff().
			SetValidateDiagFunc(validation.ToDiagFunc(validation.StringInSlice([]string{"PRO", "CLASSIC"}, false)))

		// Add no_wait field to schema
		m["no_wait"] = &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "If true, skip waiting for the warehouse to start after creation.",
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == "" && new == "false" {
					return true
				}
				return old == new
			},
		}

		return m
	})
	return common.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var se sql.CreateWarehouseRequest
			common.DataToStructPointer(d, s, &se)
			common.SetForceSendFields(&se, d, ForceSendFields)
			wait, err := w.Warehouses.Create(ctx, se)
			if err != nil {
				return fmt.Errorf("failed creating warehouse: %w", err)
			}

			d.SetId(wait.Id)

			// Check if no_wait flag is set to true
			noWait, ok := d.GetOk("no_wait")
			if ok && noWait.(bool) {
				return nil
			}

			// Wait for warehouse to start if no_wait is false or not set
			_, err = wait.Get()
			if err != nil {
				// Rollback by deleting the warehouse
				rollbackErr := w.Warehouses.DeleteById(ctx, wait.Id)
				if rollbackErr != nil {
					return fmt.Errorf("failed waiting for warehouse to start: %w. when rolling back, also failed: %w", err, rollbackErr)
				}
				return fmt.Errorf("failed waiting for warehouse to start: %w", err)
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var se sql.EditWarehouseRequest
			common.DataToStructPointer(d, s, &se)
			common.SetForceSendFields(&se, d, ForceSendFields)
			se.Id = d.Id()
			_, err = w.Warehouses.Edit(ctx, se)
			if err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.Warehouses.DeleteById(ctx, d.Id())
		},
		Schema: s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return d.Clear("health")
		},
	}
}
