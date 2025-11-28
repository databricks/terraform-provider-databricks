package exporter

import (
	"fmt"
	"log"

	"github.com/databricks/databricks-sdk-go/service/dataquality"
	"github.com/databricks/databricks-sdk-go/service/qualitymonitorv2"
	data_quality_monitor "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/data_quality_monitor"
	quality_monitor_v2 "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/quality_monitor_v2"
)

func listQualityMonitorsV2(ic *importContext) error {
	it := ic.workspaceClient.QualityMonitorV2.ListQualityMonitor(ic.Context, qualitymonitorv2.ListQualityMonitorRequest{})
	for it.HasNext(ic.Context) {
		monitor, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		ic.Emit(&resource{
			Resource: "databricks_quality_monitor_v2",
			ID:       fmt.Sprintf("%s,%s", monitor.ObjectType, monitor.ObjectId),
		})
	}
	return nil
}

func importQualityMonitorV2(ic *importContext, r *resource) error {
	log.Printf("[DEBUG] Importing quality monitor v2: %s", r.ID)
	// Convert Plugin Framework state to Go SDK struct
	var monitor qualitymonitorv2.QualityMonitor
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper,
		quality_monitor_v2.QualityMonitor{}, &monitor); err != nil {
		return err
	}

	// Emit dependencies based on object_type
	// For "schema" object type, emit the schema resource
	if monitor.ObjectType == "schema" && monitor.ObjectId != "" {
		// object_id is the schema UUID, but we need to emit by schema full name
		// Try to get the schema full name if it's available in any config
		// For now, we rely on the Depends references in importables.go
		// which will match object_id to schema_id
		ic.Emit(&resource{
			Resource: "databricks_schema",
			ID:       monitor.ObjectId,
		})
	}

	return nil
}

func importDataQualityMonitor(ic *importContext, r *resource) error {
	log.Printf("[DEBUG] Importing data quality monitor: %s", r.ID)
	// Convert Plugin Framework state to Go SDK struct
	var monitor dataquality.Monitor
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper,
		data_quality_monitor.Monitor{}, &monitor); err != nil {
		return err
	}

	// TODO: Figure out how to emit the monitored schema/table directly by object_id (UUID)
	// Currently relying on Depends references (object_id -> table_id/schema_id) and
	// full name fields (monitored_table_name). Could be more robust if we could lookup
	// table/schema by UUID and emit them directly here.
	// Challenges: SDK's Tables.Get() and Schemas.Get() only accept full names, not UUIDs

	// Emit dependencies based on monitor type
	if monitor.ObjectType == "table" && monitor.DataProfilingConfig != nil {
		config := monitor.DataProfilingConfig

		// Emit the monitored table - table's import will emit schema and catalog
		if config.MonitoredTableName != "" {
			ic.Emit(&resource{
				Resource: "databricks_sql_table",
				ID:       config.MonitoredTableName,
			})
		}

		// Warehouse for running monitor queries
		if config.WarehouseId != "" {
			ic.Emit(&resource{
				Resource: "databricks_sql_endpoint",
				ID:       config.WarehouseId,
			})
		}

		// Baseline table for drift analysis - table's import will emit its schema/catalog
		if config.BaselineTableName != "" {
			ic.Emit(&resource{
				Resource: "databricks_sql_table",
				ID:       config.BaselineTableName,
			})
		}

		// Notification emails
		if config.NotificationSettings != nil &&
			config.NotificationSettings.OnFailure != nil {
			for _, email := range config.NotificationSettings.OnFailure.EmailAddresses {
				ic.emitUserOrServicePrincipal(email)
			}
		}
	}
	// Note: For schema monitors (object_type == "schema"), the schema dependency is already
	// handled via the Depends reference: object_id -> schema_id. No need to emit it here.

	return nil
}
