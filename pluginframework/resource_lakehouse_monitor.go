package pluginframework

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

const lakehouseMonitorDefaultProvisionTimeout = 15 * time.Minute

func WaitForMonitor(w *databricks.WorkspaceClient, ctx context.Context, monitorName string) error {
	return retry.RetryContext(ctx, lakehouseMonitorDefaultProvisionTimeout, func() *retry.RetryError {
		endpoint, err := w.LakehouseMonitors.GetByTableName(ctx, monitorName)
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

func ResourceLakehouseMonitor() resource.Resource {
	return &LakehouseMonitorResource{}
}

type LakehouseMonitorResource struct{}

func (r *LakehouseMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lakehouse_monitor_plugin_framework"
}

func (r *LakehouseMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Schema for lakehouse monitor",
		// We would need similar method to common.StructToSchema but since this is PoC, we are not implementing it here
		// Also this isn't complete, using incomplete schema for PoCs just to check if it's working
		Attributes: map[string]schema.Attribute{
			"assets_dir": schema.StringAttribute{
				Description: "The directory to store monitoring assets (e.g. dashboard, metric tables)",
				Optional:    true,
			},
			"baseline_table_name": schema.StringAttribute{
				Description: "Name of the baseline table from which drift metrics are computed from Columns in the monitored table should also be present in the baseline table.",
				Optional:    true,
			},
			"dashboard_id": schema.StringAttribute{
				Description: "Id of dashboard that visualizes the computed metrics. This can be empty if the monitor is in PENDING state.",
				Optional:    true,
			},
			"table_name": schema.StringAttribute{
				Description: "The full name of the table to monitor. Format:__catalog_name__.__schema_name__.__table_name__.",
				Required:    true,
			},
		},
	}
}

func (r *LakehouseMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	c := common.DatabricksClient{}
	w, err := c.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var create catalog.CreateMonitor
	diags := req.Plan.Get(ctx, &create)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpoint, err := w.LakehouseMonitors.Create(ctx, create)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get create monitor", err.Error())
		return
	}
	err = WaitForMonitor(w, ctx, endpoint.TableName)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get newly created monitor", err.Error())
		return
	}
}

func (r *LakehouseMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	c := common.DatabricksClient{}
	w, err := c.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var getMonitor catalog.GetLakehouseMonitorRequest
	diags := req.State.Get(ctx, &getMonitor)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpoint, err := w.LakehouseMonitors.GetByTableName(ctx, getMonitor.TableName)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get monitor", err.Error())
		return
	}
	diags = resp.State.Set(ctx, endpoint)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *LakehouseMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	c := common.DatabricksClient{}
	w, err := c.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var updateRequest catalog.UpdateMonitor
	diags := req.Plan.Get(ctx, &updateRequest)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err = w.LakehouseMonitors.Update(ctx, updateRequest)
	if err != nil {
		resp.Diagnostics.AddError("Failed to update monitor", err.Error())
		return
	}
	err = WaitForMonitor(w, ctx, updateRequest.TableName)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get updated monitor", err.Error())
		return
	}
}

func (r *LakehouseMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	c := common.DatabricksClient{}
	w, err := c.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var deleteRequest catalog.DeleteLakehouseMonitorRequest
	diags := req.State.Get(ctx, &deleteRequest)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	err = w.LakehouseMonitors.DeleteByTableName(ctx, deleteRequest.TableName)
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete monitor", err.Error())
		return
	}
}
