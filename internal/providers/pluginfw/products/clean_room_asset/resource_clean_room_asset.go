// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clean_room_asset

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/cleanrooms"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/cleanrooms_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "clean_room_asset"

var _ resource.ResourceWithConfigure = &CleanRoomAssetResource{}

func ResourceCleanRoomAsset() resource.Resource {
	return &CleanRoomAssetResource{}
}

type CleanRoomAssetResource struct {
	Client *autogen.DatabricksClient
}

// CleanRoomAssetExtended extends the main model with additional fields.
type CleanRoomAssetExtended struct {
	cleanrooms_tf.CleanRoomAsset
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// CleanRoomAssetExtended struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m CleanRoomAssetExtended) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.CleanRoomAsset.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetExtended
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetExtended) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.CleanRoomAsset.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m CleanRoomAssetExtended) Type(ctx context.Context) attr.Type {
	embeddedType := m.CleanRoomAsset.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *CleanRoomAssetExtended) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan CleanRoomAssetExtended) {
	m.CleanRoomAsset.SyncFieldsDuringCreateOrUpdate(ctx, plan.CleanRoomAsset)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *CleanRoomAssetExtended) SyncFieldsDuringRead(ctx context.Context, existingState CleanRoomAssetExtended) {
	m.CleanRoomAsset.SyncFieldsDuringRead(ctx, existingState.CleanRoomAsset)
}

func (r *CleanRoomAssetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *CleanRoomAssetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, CleanRoomAssetExtended{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
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

func (r *CleanRoomAssetResource) update(ctx context.Context, plan CleanRoomAssetExtended, diags *diag.Diagnostics, state *tfsdk.State) {
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

	var newState CleanRoomAssetExtended
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *CleanRoomAssetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan CleanRoomAssetExtended
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

	var newState CleanRoomAssetExtended

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

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

	var existingState CleanRoomAssetExtended
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

	var newState CleanRoomAssetExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *CleanRoomAssetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan CleanRoomAssetExtended
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

	var state CleanRoomAssetExtended
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
