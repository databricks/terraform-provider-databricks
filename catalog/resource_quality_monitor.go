package catalog

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const qualityMonitorDefaultProvisionTimeout = 15 * time.Minute

func WaitForMonitor(w *databricks.WorkspaceClient, ctx context.Context, monitorName string) error {
	return retry.RetryContext(ctx, qualityMonitorDefaultProvisionTimeout, func() *retry.RetryError {
		endpoint, err := w.QualityMonitors.GetByTableName(ctx, monitorName)
		if err != nil {
			return retry.NonRetryableError(err)
		}

		switch endpoint.Status {
		case catalog.MonitorInfoStatusMonitorStatusActive:
			return nil
		case catalog.MonitorInfoStatusMonitorStatusError, catalog.MonitorInfoStatusMonitorStatusFailed:
			return retry.NonRetryableError(fmt.Errorf("monitor status retrund %s for monitor: %s", endpoint.Status, monitorName))
		}
		return retry.RetryableError(fmt.Errorf("monitor %s is still pending", monitorName))
	})
}

func ResourceQualityMonitor() common.Resource {
	monitorSchema := common.StructToSchema(
		catalog.MonitorInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			common.CustomizeSchemaPath(m, "assets_dir").SetRequired()
			common.CustomizeSchemaPath(m, "output_schema_name").SetRequired()
			common.CustomizeSchemaPath(m, "table_name").SetRequired()
			common.CustomizeSchemaPath(m).AddNewField("skip_builtin_dashboard", &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Required: false,
			})
			common.CustomizeSchemaPath(m).AddNewField("warehouse_id", &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Required: false,
			})
			common.CustomizeSchemaPath(m, "monitor_version").SetReadOnly()
			common.CustomizeSchemaPath(m, "drift_metrics_table_name").SetReadOnly()
			common.CustomizeSchemaPath(m, "profile_metrics_table_name").SetReadOnly()
			common.CustomizeSchemaPath(m, "status").SetReadOnly()
			common.CustomizeSchemaPath(m, "dashboard_id").SetReadOnly()
			
			return m
		},
	)

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			var create catalog.CreateMonitor
			common.DataToStructPointer(d, monitorSchema, &create)
			create.TableName = d.Get("table_name").(string)

			endpoint, err := w.QualityMonitors.Create(ctx, create)
			if err != nil {
				return err
			}
			err = WaitForMonitor(w, ctx, create.TableName)
			if err != nil {
				return err
			}
			d.SetId(endpoint.TableName)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			endpoint, err := w.QualityMonitors.GetByTableName(ctx, d.Id())
			if err != nil {
				return err

			}
			return common.StructToData(endpoint, monitorSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var update catalog.UpdateMonitor
			common.DataToStructPointer(d, monitorSchema, &update)
			update.TableName = d.Get("table_name").(string)
			_, err = w.QualityMonitors.Update(ctx, update)
			if err != nil {
				return err
			}
			return WaitForMonitor(w, ctx, update.TableName)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.QualityMonitors.DeleteByTableName(ctx, d.Id())
		},
		Schema: monitorSchema,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(qualityMonitorDefaultProvisionTimeout),
		},
	}
}
