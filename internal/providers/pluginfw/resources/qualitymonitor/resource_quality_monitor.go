package qualitymonitor

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

const qualityMonitorDefaultProvisionTimeout = 15 * time.Minute

var _ resource.ResourceWithConfigure = &QualityMonitorResource{}

func ResourceQualityMonitor() resource.Resource {
	return &QualityMonitorResource{}
}

func waitForMonitor(ctx context.Context, w *databricks.WorkspaceClient, monitor *catalog.MonitorInfo) diag.Diagnostics {
	err := retry.RetryContext(ctx, qualityMonitorDefaultProvisionTimeout, func() *retry.RetryError {
		newMonitor, err := w.QualityMonitors.GetByTableName(ctx, monitor.TableName)
		*monitor = *newMonitor
		if err != nil {
			return retry.NonRetryableError(err)
		}

		switch newMonitor.Status {
		case catalog.MonitorInfoStatusMonitorStatusActive:
			return nil
		case catalog.MonitorInfoStatusMonitorStatusError, catalog.MonitorInfoStatusMonitorStatusFailed:
			return retry.NonRetryableError(fmt.Errorf("monitor status returned %s for monitor: %s", newMonitor.Status, newMonitor.TableName))
		}
		return retry.RetryableError(fmt.Errorf("monitor %s is still pending", newMonitor.TableName))
	})
	if err != nil {
		return diag.Diagnostics{diag.NewErrorDiagnostic("failed to get monitor", err.Error())}
	}
	return nil
}

type MonitorInfoExtended struct {
	catalog_tf.MonitorInfo
	WarehouseId          types.String `tfsdk:"warehouse_id" tf:"optional"`
	SkipBuiltinDashboard types.Bool   `tfsdk:"skip_builtin_dashboard" tf:"optional"`
}

type QualityMonitorResource struct {
	Client *common.DatabricksClient
}

func (r *QualityMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "databricks_quality_monitor_pluginframework"
}

func (r *QualityMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Quality Monitor",
		Attributes: tfschema.ResourceStructToSchemaMap(MonitorInfoExtended{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
			c.SetRequired("assets_dir")
			c.SetRequired("output_schema_name")
			c.SetRequired("table_name")
			c.SetReadOnly("monitor_version")
			c.SetReadOnly("drift_metrics_table_name")
			c.SetReadOnly("profile_metrics_table_name")
			c.SetReadOnly("status")
			c.SetReadOnly("dashboard_id")
			c.SetReadOnly("schedule", "pause_status")
			return c
		}),
	}
}

func (d *QualityMonitorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (r *QualityMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var monitorInfoTfSDK MonitorInfoExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &monitorInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var createMonitorGoSDK catalog.CreateMonitor
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, monitorInfoTfSDK, &createMonitorGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}
	monitor, err := w.QualityMonitors.Create(ctx, createMonitorGoSDK)
	if err != nil {
		resp.Diagnostics.AddError("failed to get created monitor", err.Error())
		return
	}
	resp.Diagnostics.Append(waitForMonitor(ctx, w, monitor)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var newMonitorInfoTfSDK MonitorInfoExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, monitor, &newMonitorInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newMonitorInfoTfSDK)...)
}

func (r *QualityMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var getMonitor catalog_tf.GetQualityMonitorRequest
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("table_name"), &getMonitor.TableName)...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpoint, err := w.QualityMonitors.GetByTableName(ctx, getMonitor.TableName.ValueString())
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
		}
		resp.Diagnostics.AddError("failed to get monitor", err.Error())
		return
	}
	var monitorInfoTfSDK MonitorInfoExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, endpoint, &monitorInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, monitorInfoTfSDK)...)
}

func (r *QualityMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var monitorInfoTfSDK MonitorInfoExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &monitorInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Plan is not adding `dashboard_id`, but it is in the state.
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("dashboard_id"), &monitorInfoTfSDK.DashboardId)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var updateMonitorGoSDK catalog.UpdateMonitor
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, monitorInfoTfSDK, &updateMonitorGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}
	monitor, err := w.QualityMonitors.Update(ctx, updateMonitorGoSDK)
	if err != nil {
		resp.Diagnostics.AddError("failed to update monitor", err.Error())
		return
	}
	resp.Diagnostics.Append(waitForMonitor(ctx, w, monitor)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var newMonitorInfoTfSDK MonitorInfoExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, monitor, &newMonitorInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newMonitorInfoTfSDK)...)
}

func (r *QualityMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest catalog_tf.DeleteQualityMonitorRequest
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("table_name"), &deleteRequest.TableName)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := w.QualityMonitors.DeleteByTableName(ctx, deleteRequest.TableName.ValueString())
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete monitor", err.Error())
	}
}
