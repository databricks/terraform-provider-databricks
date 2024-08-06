package pluginframework

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/path"
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

var _ resource.Resource = &QualityMonitorResource{}

func ResourceQualityMonitor() func() resource.Resource {
	return func() resource.Resource {
		return &QualityMonitorResource{}
	}
}

type MonitorInfoExtended struct {
	catalog_tf.MonitorInfo
	// Adding new fields.
	WarehouseId          types.String `tfsdk:"warehouse_id" tf:"optional"`
	SkipBuiltinDashboard types.Bool   `tfsdk:"skip_builtin_dashboard" tf:"optional"`
}

type QualityMonitorResource struct {
	Client *common.DatabricksClient
}

func (r *QualityMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "databricks_lakehouse_monitor_pluginframework"
}

func (r *QualityMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Lakehouse Monitor. MonitorInfo struct is used to create the schema",
		Attributes: common.PluginFrameworkResourceStructToSchemaMap(MonitorInfoExtended{}, func(c common.CustomizableSchemaPluginFramework) common.CustomizableSchemaPluginFramework {
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
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*common.DatabricksClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *common.DatabricksClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.Client = client
}

func (r *QualityMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var monitorInfoTfSDK MonitorInfoExtended
	diags := req.Plan.Get(ctx, &monitorInfoTfSDK)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var createMonitorGoSDK catalog.CreateMonitor
	err = common.TfSdkToGoSdkStruct(monitorInfoTfSDK, &createMonitorGoSDK, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to convert Tf SDK struct to Go SDK struct", err.Error())
		return
	}
	endpoint, err := w.QualityMonitors.Create(ctx, createMonitorGoSDK)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get created monitor", err.Error())
		return
	}
	err = WaitForMonitor(w, ctx, endpoint.TableName)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for newly created monitor", err.Error())
		return
	}

	// Get the created monitor.
	new_endpoint, err := w.QualityMonitors.GetByTableName(ctx, createMonitorGoSDK.TableName)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get newly created monitor", err.Error())
		return
	}

	var newMonitorInfoTfSDK MonitorInfoExtended
	err = common.GoSdkToTfSdkStruct(new_endpoint, &newMonitorInfoTfSDK, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to convert Go SDK struct to TF SDK struct", err.Error())
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newMonitorInfoTfSDK)...)
}

func (r *QualityMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var getMonitor catalog_tf.GetQualityMonitorRequest
	diags := req.State.GetAttribute(ctx, path.Root("table_name"), &getMonitor.TableName)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpoint, err := w.QualityMonitors.GetByTableName(ctx, getMonitor.TableName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to get monitor", err.Error())
		return
	}
	var monitorInfoTfSDK MonitorInfoExtended
	err = common.GoSdkToTfSdkStruct(endpoint, &monitorInfoTfSDK, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to convert Go SDK struct to TF SDK struct", err.Error())
	}
	diags = resp.State.Set(ctx, monitorInfoTfSDK)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *QualityMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var monitorInfoTfSDK MonitorInfoExtended
	diags := req.Plan.Get(ctx, &monitorInfoTfSDK)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Plan is not adding `dashboard_id`, but it is in the state.
	diags = req.State.GetAttribute(ctx, path.Root("dashboard_id"), &monitorInfoTfSDK.DashboardId)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var updateMonitorGoSDK catalog.UpdateMonitor
	err = common.TfSdkToGoSdkStruct(monitorInfoTfSDK, &updateMonitorGoSDK, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to convert Tf SDK struct to Go SDK struct", err.Error())
		return
	}
	_, err = w.QualityMonitors.Update(ctx, updateMonitorGoSDK)
	if err != nil {
		resp.Diagnostics.AddError("Failed to update monitor", err.Error())
		return
	}
	err = WaitForMonitor(w, ctx, updateMonitorGoSDK.TableName)
	if err != nil {
		resp.Diagnostics.AddError("Failed to wait for updated monitor", err.Error())
		return
	}

	// Get the created monitor.
	new_endpoint, err := w.QualityMonitors.GetByTableName(ctx, updateMonitorGoSDK.TableName)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get newly created monitor", err.Error())
		return
	}

	var newMonitorInfoTfSDK MonitorInfoExtended
	err = common.GoSdkToTfSdkStruct(new_endpoint, &newMonitorInfoTfSDK, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to convert Go SDK struct to TF SDK struct", err.Error())
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newMonitorInfoTfSDK)...)
}

func (r *QualityMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var deleteRequest catalog_tf.DeleteQualityMonitorRequest
	diags := req.State.GetAttribute(ctx, path.Root("table_name"), &deleteRequest.TableName)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	err = w.QualityMonitors.DeleteByTableName(ctx, deleteRequest.TableName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete monitor", err.Error())
		return
	}
}
