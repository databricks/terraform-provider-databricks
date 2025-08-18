// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace_network_option

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

const resourceName = "workspace_network_option"

var _ resource.ResourceWithConfigure = &WorkspaceNetworkOptionResource{}

func ResourceWorkspaceNetworkOption() resource.Resource {
	return &WorkspaceNetworkOptionResource{}
}

type WorkspaceNetworkOptionResource struct {
	Client *autogen.DatabricksClient
}

func (r *WorkspaceNetworkOptionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *WorkspaceNetworkOptionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, settings_tf.WorkspaceNetworkOption{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(int64planmodifier.UseStateForUnknown(), "workspace_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks workspace_network_option",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *WorkspaceNetworkOptionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *WorkspaceNetworkOptionResource) update(ctx context.Context, plan settings_tf.WorkspaceNetworkOption, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetAccountClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var workspace_network_option settings.WorkspaceNetworkOption

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &workspace_network_option)...)
	if diags.HasError() {
		return
	}

	updateRequest := settings.UpdateWorkspaceNetworkOptionRequest{
		WorkspaceNetworkOption: workspace_network_option,
		WorkspaceId:            plan.WorkspaceId.ValueInt64(),
	}

	response, err := client.WorkspaceNetworkConfiguration.UpdateWorkspaceNetworkOptionRpc(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update workspace_network_option", err.Error())
		return
	}

	var newState settings_tf.WorkspaceNetworkOption
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *WorkspaceNetworkOptionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan settings_tf.WorkspaceNetworkOption
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *WorkspaceNetworkOptionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState settings_tf.WorkspaceNetworkOption
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest settings.GetWorkspaceNetworkOptionRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.WorkspaceNetworkConfiguration.GetWorkspaceNetworkOptionRpc(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get workspace_network_option", err.Error())
		return
	}

	var newState settings_tf.WorkspaceNetworkOption
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *WorkspaceNetworkOptionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan settings_tf.WorkspaceNetworkOption
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *WorkspaceNetworkOptionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

var _ resource.ResourceWithImportState = &WorkspaceNetworkOptionResource{}

func (r *WorkspaceNetworkOptionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: workspace_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	workspaceId, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse import identifier", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("workspace_id"), workspaceId)...)
}
