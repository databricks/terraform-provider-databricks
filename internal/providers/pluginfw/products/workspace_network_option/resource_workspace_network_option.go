// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace_network_option

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "workspace_network_option"

var _ resource.ResourceWithConfigure = &WorkspaceNetworkOptionResource{}

func ResourceWorkspaceNetworkOption() resource.Resource {
	return &WorkspaceNetworkOptionResource{}
}

type WorkspaceNetworkOptionResource struct {
	Client *autogen.DatabricksClient
}

// WorkspaceNetworkOption extends the main model with additional fields.
type WorkspaceNetworkOption struct {
	// The network policy ID to apply to the workspace. This controls the
	// network access rules for all serverless compute resources in the
	// workspace. Each workspace can only be linked to one policy at a time. If
	// no policy is explicitly assigned, the workspace will use
	// 'default-policy'.
	NetworkPolicyId types.String `tfsdk:"network_policy_id"`
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// WorkspaceNetworkOption struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m WorkspaceNetworkOption) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceNetworkOption
// only implements ToObjectValue() and Type().
func (m WorkspaceNetworkOption) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"network_policy_id": m.NetworkPolicyId,
			"workspace_id": m.WorkspaceId,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m WorkspaceNetworkOption) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"network_policy_id": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *WorkspaceNetworkOption) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceNetworkOption) {
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *WorkspaceNetworkOption) SyncFieldsDuringRead(ctx context.Context, from WorkspaceNetworkOption) {
}

func (m WorkspaceNetworkOption) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_policy_id"] = attrs["network_policy_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()

	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

func (r *WorkspaceNetworkOptionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *WorkspaceNetworkOptionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, WorkspaceNetworkOption{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks workspace_network_option",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *WorkspaceNetworkOptionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *WorkspaceNetworkOptionResource) update(ctx context.Context, plan WorkspaceNetworkOption, diags *diag.Diagnostics, state *tfsdk.State) {
	var workspace_network_option settings.WorkspaceNetworkOption

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &workspace_network_option)...)
	if diags.HasError() {
		return
	}

	updateRequest := settings.UpdateWorkspaceNetworkOptionRequest{
		WorkspaceNetworkOption: workspace_network_option,
		WorkspaceId:            plan.WorkspaceId.ValueInt64(),
	}

	client, clientDiags := r.Client.GetAccountClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.WorkspaceNetworkConfiguration.UpdateWorkspaceNetworkOptionRpc(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update workspace_network_option", err.Error())
		return
	}

	var newState WorkspaceNetworkOption
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *WorkspaceNetworkOptionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan WorkspaceNetworkOption
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *WorkspaceNetworkOptionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState WorkspaceNetworkOption
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest settings.GetWorkspaceNetworkOptionRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
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

	var newState WorkspaceNetworkOption
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *WorkspaceNetworkOptionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan WorkspaceNetworkOption
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
