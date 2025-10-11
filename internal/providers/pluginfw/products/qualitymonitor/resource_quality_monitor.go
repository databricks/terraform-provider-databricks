package qualitymonitor

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "quality_monitor"

const qualityMonitorDefaultProvisionTimeout = 15 * time.Minute

var _ resource.ResourceWithConfigure = &QualityMonitorResource{}

func ResourceQualityMonitor() resource.Resource {
	return &QualityMonitorResource{}
}

func waitForMonitor(ctx context.Context, w *databricks.WorkspaceClient, monitor *catalog.MonitorInfo) diag.Diagnostics {
	updatedMonitor, err := retries.Poll[catalog.MonitorInfo](ctx, qualityMonitorDefaultProvisionTimeout, func() (*catalog.MonitorInfo, *retries.Err) {
		newMonitor, err := w.QualityMonitors.Get(ctx, catalog.GetQualityMonitorRequest{
			TableName: monitor.TableName,
		})
		if err != nil {
			return nil, retries.Halt(fmt.Errorf("failed to get monitor: %s", err))
		}

		switch newMonitor.Status {
		case catalog.MonitorInfoStatusMonitorStatusActive:
			return newMonitor, nil
		case catalog.MonitorInfoStatusMonitorStatusError, catalog.MonitorInfoStatusMonitorStatusFailed:
			return nil, retries.Halt(fmt.Errorf("monitor status returned %s for monitor: %s", newMonitor.Status, newMonitor.TableName))
		}
		return nil, retries.Continue(fmt.Errorf("monitor %s is still pending", newMonitor.TableName))
	})
	if err != nil {
		return diag.Diagnostics{diag.NewErrorDiagnostic("failed to get monitor", err.Error())}
	}
	*monitor = *updatedMonitor
	return nil
}

type MonitorInfoExtended struct {
	catalog_tf.MonitorInfo_SdkV2
	WarehouseId          types.String `tfsdk:"warehouse_id"`
	SkipBuiltinDashboard types.Bool   `tfsdk:"skip_builtin_dashboard"`
	ID                   types.String `tfsdk:"id"` // Adding ID field to stay compatible with SDKv2
	ProviderConfig       types.Object `tfsdk:"provider_config"`
}

var _ pluginfwcommon.ComplexFieldTypeProvider = MonitorInfoExtended{}

func (m MonitorInfoExtended) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	attrs := m.MonitorInfo_SdkV2.GetComplexFieldTypes(ctx)
	attrs["provider_config"] = reflect.TypeOf(tfschema.ProviderConfig{})
	return attrs
}

type QualityMonitorResource struct {
	Client *common.DatabricksClient
}

func (r *QualityMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

func (r *QualityMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, MonitorInfoExtended{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.ConfigureAsSdkV2Compatible()
		c.SetRequired("assets_dir")
		c.SetReadOnly("monitor_version")
		c.SetReadOnly("drift_metrics_table_name")
		c.SetReadOnly("profile_metrics_table_name")
		c.SetReadOnly("status")
		c.SetReadOnly("dashboard_id")
		c.SetReadOnly("schedule", "pause_status")
		c.SetOptional("warehouse_id")
		c.SetOptional("skip_builtin_dashboard")
		c.SetComputed("id")
		c.SetOptional("id")
		c.SetOptional("provider_config")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Quality Monitor",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (d *QualityMonitorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if d.Client == nil && req.ProviderData != nil {
		d.Client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (d *QualityMonitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("table_name"), req, resp)
}

func (r *QualityMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	var monitorInfoTfSDK MonitorInfoExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &monitorInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var workspaceID string
	if !monitorInfoTfSDK.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfigData
		resp.Diagnostics.Append(monitorInfoTfSDK.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}

	w, diags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
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
		resp.Diagnostics.AddError("failed to create monitor", err.Error())
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

	// Set the ID to the table name
	newMonitorInfoTfSDK.ID = newMonitorInfoTfSDK.TableName
	// We need it to fill additional fields as they are not returned by the API
	newMonitorInfoTfSDK.WarehouseId = monitorInfoTfSDK.WarehouseId
	newMonitorInfoTfSDK.SkipBuiltinDashboard = monitorInfoTfSDK.SkipBuiltinDashboard
	newMonitorInfoTfSDK.ProviderConfig = monitorInfoTfSDK.ProviderConfig

	resp.Diagnostics.Append(resp.State.Set(ctx, newMonitorInfoTfSDK)...)
}

func (r *QualityMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var monitorInfoTfSDK MonitorInfoExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &monitorInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var workspaceID string
	if !monitorInfoTfSDK.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfigData
		resp.Diagnostics.Append(monitorInfoTfSDK.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}

	w, diags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint, err := w.QualityMonitors.Get(ctx, catalog.GetQualityMonitorRequest{
		TableName: monitorInfoTfSDK.TableName.ValueString(),
	})
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get monitor", err.Error())
		return
	}
	var newMonitorInfoTfSDK MonitorInfoExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, endpoint, &newMonitorInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newMonitorInfoTfSDK.ID = monitorInfoTfSDK.TableName
	if monitorInfoTfSDK.WarehouseId.ValueString() != "" {
		newMonitorInfoTfSDK.WarehouseId = monitorInfoTfSDK.WarehouseId
	}
	if monitorInfoTfSDK.SkipBuiltinDashboard.ValueBool() {
		newMonitorInfoTfSDK.SkipBuiltinDashboard = monitorInfoTfSDK.SkipBuiltinDashboard
	}

	newMonitorInfoTfSDK.ProviderConfig = monitorInfoTfSDK.ProviderConfig
	resp.Diagnostics.Append(resp.State.Set(ctx, newMonitorInfoTfSDK)...)
}

func (r *QualityMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

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
	if updateMonitorGoSDK.Schedule != nil {
		updateMonitorGoSDK.Schedule.PauseStatus = ""
	}

	var workspaceID string
	if !monitorInfoTfSDK.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfigData
		resp.Diagnostics.Append(monitorInfoTfSDK.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}

	w, diags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
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
	// We need it to fill additional fields as they are not returned by the API
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("warehouse_id"), &newMonitorInfoTfSDK.WarehouseId)...)
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("skip_builtin_dashboard"), &newMonitorInfoTfSDK.SkipBuiltinDashboard)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newMonitorInfoTfSDK.ProviderConfig = monitorInfoTfSDK.ProviderConfig
	resp.Diagnostics.Append(resp.State.Set(ctx, newMonitorInfoTfSDK)...)
}

func (r *QualityMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var monitorInfoTfSDK MonitorInfoExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &monitorInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var workspaceID string
	if !monitorInfoTfSDK.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfigData
		resp.Diagnostics.Append(monitorInfoTfSDK.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}

	w, diags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := w.QualityMonitors.Delete(ctx, catalog.DeleteQualityMonitorRequest{
		TableName: monitorInfoTfSDK.TableName.ValueString(),
	})
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete monitor", err.Error())
	}
}
