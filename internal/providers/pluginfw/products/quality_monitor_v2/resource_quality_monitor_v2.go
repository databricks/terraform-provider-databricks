// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package quality_monitor_v2

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/qualitymonitorv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/qualitymonitorv2_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

const resourceName = "quality_monitor_v2"

var _ resource.ResourceWithConfigure = &QualityMonitorV2Resource{}

func ResourceQualityMonitorV2() resource.Resource {
	return &QualityMonitorV2Resource{}
}

type QualityMonitorV2Resource struct {
	Client *autogen.DatabricksClient
}

func (r *QualityMonitorV2Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *QualityMonitorV2Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, qualitymonitorv2_tf.QualityMonitor{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "object_type")
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "object_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks quality_monitor_v2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *QualityMonitorV2Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *QualityMonitorV2Resource) update(ctx context.Context, plan qualitymonitorv2_tf.QualityMonitor, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var quality_monitor qualitymonitorv2.QualityMonitor
	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &quality_monitor)...)
	if diags.HasError() {
		return
	}

	updateRequest := qualitymonitorv2.UpdateQualityMonitorRequest{
		QualityMonitor: quality_monitor,
		ObjectType:     plan.ObjectType.ValueString(),
		ObjectId:       plan.ObjectId.ValueString(),
	}

	response, err := client.QualityMonitorV2.UpdateQualityMonitor(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update quality_monitor_v2", err.Error())
		return
	}

	var newState qualitymonitorv2_tf.QualityMonitor
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncEffectiveFieldsDuringCreateOrUpdate(plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *QualityMonitorV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan qualitymonitorv2_tf.QualityMonitor
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var quality_monitor qualitymonitorv2.QualityMonitor
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &quality_monitor)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := qualitymonitorv2.CreateQualityMonitorRequest{
		QualityMonitor: quality_monitor,
	}

	response, err := client.QualityMonitorV2.CreateQualityMonitor(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create quality_monitor_v2", err.Error())
		return
	}

	var newState qualitymonitorv2_tf.QualityMonitor

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncEffectiveFieldsDuringCreateOrUpdate(plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *QualityMonitorV2Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState qualitymonitorv2_tf.QualityMonitor
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest qualitymonitorv2.GetQualityMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.QualityMonitorV2.GetQualityMonitor(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get quality_monitor_v2", err.Error())
		return
	}

	var newState qualitymonitorv2_tf.QualityMonitor
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncEffectiveFieldsDuringRead(existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *QualityMonitorV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan qualitymonitorv2_tf.QualityMonitor
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *QualityMonitorV2Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state qualitymonitorv2_tf.QualityMonitor
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest qualitymonitorv2.DeleteQualityMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.QualityMonitorV2.DeleteQualityMonitor(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete quality_monitor_v2", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &QualityMonitorV2Resource{}

func (r *QualityMonitorV2Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: object_type,object_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	objectType := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("object_type"), objectType)...)
	objectId := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("object_id"), objectId)...)
}
