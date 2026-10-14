package sql

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ClusterSizes for SQL endpoints
var (
	ClusterSizes    = []string{"2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large", "5X-Large"}
	MaxNumClusters  = 40
	ForceSendFields = []string{"enable_serverless_compute", "enable_photon", "auto_stop_mins"}
)

type SqlWarehouse struct {
	sql.GetWarehouseResponse
	realtimeWarehouseFields

	// The data source ID is not part of the endpoint API response.
	// We manually resolve it by retrieving the list of data sources
	// and matching this entity's endpoint ID.
	DataSourceId string `json:"data_source_id,omitempty" tf:"computed"`
	common.Namespace
}

func getSqlWarehouse(ctx context.Context, w *databricks.WorkspaceClient, id string) (*SqlWarehouse, error) {
	se, err := w.Warehouses.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	warehouse := SqlWarehouse{
		GetWarehouseResponse: *se,
	}
	switch warehouse.WarehouseType {
	case "REALTIME", "REYDEN":
		warehouse.WarehouseType = "REALTIME"
		warehouse.StatementTimeout = 3600
		if err := doRealtimeWarehouseRequest(ctx, w, http.MethodGet, "/api/2.0/sql/warehouses/"+id, nil, &warehouse.realtimeWarehouseFields); err != nil {
			return nil, err
		}
	}
	return &warehouse, nil
}

func resolveDataSourceID(ctx context.Context, w *databricks.WorkspaceClient, warehouseId string) (string, error) {
	list, err := w.DataSources.List(ctx)
	if err != nil {
		// Don't fail if DataSources.List errors
		return "", nil
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
		m["auto_stop_mins"].Optional = true
		m["auto_stop_mins"].Computed = true
		m["statement_timeout"].ValidateDiagFunc = validation.ToDiagFunc(validation.IntAtLeast(0))
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
			SetValidateDiagFunc(validation.ToDiagFunc(validation.StringInSlice([]string{"PRO", "CLASSIC", "REALTIME"}, false)))

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

		common.NamespaceCustomizeSchemaMap(m)
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
			if se.WarehouseType == "REALTIME" {
				var response sql.CreateWarehouseResponse
				err := writeRealtimeWarehouse(ctx, w, "/api/2.0/sql/warehouses", se, d.Get("statement_timeout").(int), &response)
				if err != nil {
					return fmt.Errorf("failed creating warehouse: %w", err)
				}
				d.SetId(response.Id)
				// REALTIME warehouses are created stopped.
				return nil
			}
			wait, err := w.Warehouses.Create(ctx, se)
			if err != nil {
				return fmt.Errorf("failed creating warehouse: %w", err)
			}

			d.SetId(wait.Id)

			if d.Get("no_wait").(bool) {
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
			if warehouse.WarehouseType != "REALTIME" {
				warehouse.DataSourceId, err = resolveDataSourceID(ctx, w, d.Id())
				if err != nil {
					return err
				}
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
			if se.WarehouseType == "REALTIME" {
				return writeRealtimeWarehouse(ctx, w, "/api/2.0/sql/warehouses/"+d.Id()+"/edit", se, d.Get("statement_timeout").(int), nil)
			}
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
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			if err := common.NamespaceCustomizeDiff(ctx, d, c); err != nil {
				return err
			}
			if err := customizeSqlWarehouseDiff(d); err != nil {
				return err
			}
			return d.Clear("health")
		},
	}
}

func customizeSqlWarehouseDiff(d *schema.ResourceDiff) error {
	config := d.GetRawConfig()
	if config.IsNull() {
		return nil
	}
	isRealtime := d.Get("warehouse_type") == "REALTIME"
	if d.Id() != "" && d.NewValueKnown("warehouse_type") && d.HasChange("warehouse_type") {
		oldType, _ := d.GetChange("warehouse_type")
		if isRealtime || oldType == "REALTIME" {
			return fmt.Errorf("changing warehouse_type to or from REALTIME is not supported")
		}
	}
	if config.GetAttr("auto_stop_mins").IsNull() && !d.NewValueKnown("warehouse_type") {
		return d.SetNewComputed("auto_stop_mins")
	}
	defaults := map[string]any{"auto_stop_mins": 120}
	if isRealtime {
		defaults["auto_stop_mins"] = 10
		defaults["statement_timeout"] = 3600
		defaults["enable_serverless_compute"] = true
	} else if d.NewValueKnown("warehouse_type") && !config.GetAttr("statement_timeout").IsNull() {
		return fmt.Errorf("statement_timeout is only supported for REALTIME warehouses")
	}
	for field, value := range defaults {
		if config.GetAttr(field).IsNull() {
			if err := d.SetNew(field, value); err != nil {
				return err
			}
		}
	}
	if isRealtime {
		return validateRealtimeWarehouse(d)
	}
	return nil
}

func validateRealtimeWarehouse(d *schema.ResourceDiff) error {
	if mins := d.Get("auto_stop_mins").(int); d.NewValueKnown("auto_stop_mins") && mins != 0 && mins < 5 {
		return fmt.Errorf("auto_stop_mins must be 0 or at least 5 for REALTIME warehouses")
	}
	for _, field := range []string{"min_num_clusters", "max_num_clusters"} {
		if !d.GetRawConfig().GetAttr(field).IsNull() && d.NewValueKnown(field) && d.Get(field).(int) != 1 {
			return fmt.Errorf("%s must be 1 for REALTIME warehouses", field)
		}
	}
	for _, field := range []string{"enable_photon", "enable_serverless_compute"} {
		if d.NewValueKnown(field) && !d.Get(field).(bool) {
			return fmt.Errorf("%s must be true for REALTIME warehouses", field)
		}
	}
	if d.NewValueKnown("spot_instance_policy") && d.Get("spot_instance_policy") != "COST_OPTIMIZED" {
		return fmt.Errorf("spot_instance_policy cannot be configured for REALTIME warehouses")
	}
	return nil
}
