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

func ResourceQualityMonitor() resource.Resource {
	return &QualityMonitorResource{}
}

type QualityMonitorResource struct{}

func (r *QualityMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lakehouse_monitor_plugin_framework"
}

func (r *QualityMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Lakehouse Monitor. MonitorInfo struct is used to create the schema",
		// Attributes:  common.PluginFrameworkResourceStructToSchemaMap(catalog.MonitorInfo{}),
		Attributes: map[string]schema.Attribute{},
	}
}

func (r *QualityMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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
	endpoint, err := w.QualityMonitors.Create(ctx, create)
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

func (r *QualityMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	c := common.DatabricksClient{}
	w, err := c.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var getMonitor catalog.GetQualityMonitorRequest
	diags := req.State.Get(ctx, &getMonitor)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpoint, err := w.QualityMonitors.GetByTableName(ctx, getMonitor.TableName)
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

func (r *QualityMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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
	_, err = w.QualityMonitors.Update(ctx, updateRequest)
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

func (r *QualityMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	c := common.DatabricksClient{}
	w, err := c.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var deleteRequest catalog.DeleteQualityMonitorRequest
	diags := req.State.Get(ctx, &deleteRequest)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	err = w.QualityMonitors.DeleteByTableName(ctx, deleteRequest.TableName)
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete monitor", err.Error())
		return
	}
}
