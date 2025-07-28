// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clean_room_asset

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/cleanrooms"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/cleanrooms_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

const resourceName = "clean_room_asset"

var _ resource.ResourceWithConfigure = &CleanRoomAssetResource{}

func ResourceCleanRoomAsset() resource.Resource {
	return &CleanRoomAssetResource{}
}

type CleanRoomAssetResource struct {
	Client *autogen.DatabricksClient
}

func (r *CleanRoomAssetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *CleanRoomAssetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, cleanrooms_tf.CleanRoomAsset{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "clean_room_name")
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "name")
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "asset_type")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks clean_room_asset",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CleanRoomAssetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *CleanRoomAssetResource) update(ctx context.Context, plan cleanrooms_tf.CleanRoomAsset, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var clean_room_asset cleanrooms.CleanRoomAsset

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &clean_room_asset)...)
	if diags.HasError() {
		return
	}

	updateRequest := cleanrooms.UpdateCleanRoomAssetRequest{
		Asset:         clean_room_asset,
		AssetType:     cleanrooms.CleanRoomAssetAssetType(plan.AssetType.ValueString()),
		CleanRoomName: plan.CleanRoomName.ValueString(),
		Name:          plan.Name.ValueString(),
	}

	response, err := client.CleanRoomAssets.Update(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update clean_room_asset", err.Error())
		return
	}

	var newState cleanrooms_tf.CleanRoomAsset
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *CleanRoomAssetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan cleanrooms_tf.CleanRoomAsset
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var clean_room_asset cleanrooms.CleanRoomAsset

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &clean_room_asset)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := cleanrooms.CreateCleanRoomAssetRequest{
		Asset:         clean_room_asset,
		CleanRoomName: plan.CleanRoomName.ValueString(),
	}

	response, err := client.CleanRoomAssets.Create(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create clean_room_asset", err.Error())
		return
	}

	var newState cleanrooms_tf.CleanRoomAsset

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CleanRoomAssetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState cleanrooms_tf.CleanRoomAsset
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest cleanrooms.GetCleanRoomAssetRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.CleanRoomAssets.Get(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get clean_room_asset", err.Error())
		return
	}

	var newState cleanrooms_tf.CleanRoomAsset

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *CleanRoomAssetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan cleanrooms_tf.CleanRoomAsset
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *CleanRoomAssetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state cleanrooms_tf.CleanRoomAsset
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest cleanrooms.DeleteCleanRoomAssetRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.CleanRoomAssets.Delete(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete clean_room_asset", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &CleanRoomAssetResource{}

func (r *CleanRoomAssetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: clean_room_name,name,asset_type. Got: %q",
				req.ID,
			),
		)
		return
	}

	cleanRoomName := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("clean_room_name"), cleanRoomName)...)
	name := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
	assetType := parts[2]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("asset_type"), assetType)...)
}
