// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_network_policy

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
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

const resourceName = "account_network_policy"

var _ resource.ResourceWithConfigure = &AccountNetworkPolicyResource{}

func ResourceAccountNetworkPolicy() resource.Resource {
	return &AccountNetworkPolicyResource{}
}

type AccountNetworkPolicyResource struct {
	Client *autogen.DatabricksClient
}

// AccountNetworkPolicyExtended extends the main model with additional fields.
type AccountNetworkPolicyExtended struct {
	settings_tf.AccountNetworkPolicy
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// AccountNetworkPolicyExtended struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AccountNetworkPolicyExtended) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.AccountNetworkPolicy.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountNetworkPolicyExtended
// only implements ToObjectValue() and Type().
func (m AccountNetworkPolicyExtended) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.AccountNetworkPolicy.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AccountNetworkPolicyExtended) Type(ctx context.Context) attr.Type {
	embeddedType := m.AccountNetworkPolicy.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *AccountNetworkPolicyExtended) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan AccountNetworkPolicyExtended) {
	m.AccountNetworkPolicy.SyncFieldsDuringCreateOrUpdate(ctx, plan.AccountNetworkPolicy)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *AccountNetworkPolicyExtended) SyncFieldsDuringRead(ctx context.Context, existingState AccountNetworkPolicyExtended) {
	m.AccountNetworkPolicy.SyncFieldsDuringRead(ctx, existingState.AccountNetworkPolicy)
}

func (r *AccountNetworkPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *AccountNetworkPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, AccountNetworkPolicyExtended{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "network_policy_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks account_network_policy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AccountNetworkPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *AccountNetworkPolicyResource) update(ctx context.Context, plan AccountNetworkPolicyExtended, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetAccountClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var account_network_policy settings.AccountNetworkPolicy

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &account_network_policy)...)
	if diags.HasError() {
		return
	}

	updateRequest := settings.UpdateNetworkPolicyRequest{
		NetworkPolicy:   account_network_policy,
		NetworkPolicyId: plan.NetworkPolicyId.ValueString(),
	}

	response, err := client.NetworkPolicies.UpdateNetworkPolicyRpc(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update account_network_policy", err.Error())
		return
	}

	var newState AccountNetworkPolicyExtended
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *AccountNetworkPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan AccountNetworkPolicyExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var account_network_policy settings.AccountNetworkPolicy

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &account_network_policy)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := settings.CreateNetworkPolicyRequest{
		NetworkPolicy: account_network_policy,
	}

	response, err := client.NetworkPolicies.CreateNetworkPolicyRpc(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create account_network_policy", err.Error())
		return
	}

	var newState AccountNetworkPolicyExtended

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

func (r *AccountNetworkPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState AccountNetworkPolicyExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest settings.GetNetworkPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.NetworkPolicies.GetNetworkPolicyRpc(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get account_network_policy", err.Error())
		return
	}

	var newState AccountNetworkPolicyExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *AccountNetworkPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan AccountNetworkPolicyExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *AccountNetworkPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state AccountNetworkPolicyExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest settings.DeleteNetworkPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.NetworkPolicies.DeleteNetworkPolicyRpc(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete account_network_policy", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &AccountNetworkPolicyResource{}

func (r *AccountNetworkPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: network_policy_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	networkPolicyId := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_policy_id"), networkPolicyId)...)
}
