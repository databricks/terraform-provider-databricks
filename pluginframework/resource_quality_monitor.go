package pluginframework

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/service/catalog_tf"
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

type QualityMonitorResource struct {
	Client *common.DatabricksClient
}

func (r *QualityMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "databricks_lakehouse_monitor_pluginframework"
}

func (r *QualityMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Lakehouse Monitor. MonitorInfo struct is used to create the schema",
		// TODO: Add CustomizeSchemaPath once it is supported in the plugin framework
		Attributes: common.PluginFrameworkResourceStructToSchemaMap(catalog_tf.MonitorInfo{}),
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
	var createMonitorTfSDK catalog_tf.CreateMonitor
	diags := req.Plan.Get(ctx, &createMonitorTfSDK)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var createMonitorGoSDK catalog.CreateMonitor
	err = common.TfSdkToGoSdkStruct(createMonitorTfSDK, &createMonitorGoSDK, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to convert Tf SDK struct to Go SDK struct", err.Error())
		return
	}
	endpoint, err := w.QualityMonitors.Create(ctx, createMonitorGoSDK)
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
	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var getMonitor catalog_tf.GetQualityMonitorRequest
	diags := req.State.Get(ctx, &getMonitor)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpoint, err := w.QualityMonitors.GetByTableName(ctx, getMonitor.TableName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to get monitor", err.Error())
		return
	}
	var monitorInfoTfSDK catalog_tf.MonitorInfo
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
	var updateMonitorTfSDK catalog_tf.UpdateMonitor
	diags := req.Plan.Get(ctx, &updateMonitorTfSDK)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var updateMonitorGoSDK catalog.UpdateMonitor
	err = common.TfSdkToGoSdkStruct(updateMonitorTfSDK, &updateMonitorGoSDK, ctx)
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
		resp.Diagnostics.AddError("Failed to get updated monitor", err.Error())
		return
	}
}

func (r *QualityMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var deleteRequest catalog_tf.DeleteQualityMonitorRequest
	diags := req.State.Get(ctx, &deleteRequest)
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
