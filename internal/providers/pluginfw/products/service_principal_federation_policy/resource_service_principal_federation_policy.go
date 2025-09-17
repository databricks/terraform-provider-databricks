// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package service_principal_federation_policy

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/oauth2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "service_principal_federation_policy"

var _ resource.ResourceWithConfigure = &FederationPolicyResource{}

func ResourceFederationPolicy() resource.Resource {
	return &FederationPolicyResource{}
}

type FederationPolicyResource struct {
	Client *autogen.DatabricksClient
}

// FederationPolicy extends the main model with additional fields.
type FederationPolicy struct {
	oauth2_tf.FederationPolicy
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// FederationPolicy struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m FederationPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.FederationPolicy.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FederationPolicy
// only implements ToObjectValue() and Type().
func (m FederationPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.FederationPolicy.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m FederationPolicy) Type(ctx context.Context) attr.Type {
	embeddedType := m.FederationPolicy.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *FederationPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan FederationPolicy) {
	m.FederationPolicy.SyncFieldsDuringCreateOrUpdate(ctx, plan.FederationPolicy)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *FederationPolicy) SyncFieldsDuringRead(ctx context.Context, existingState FederationPolicy) {
	m.FederationPolicy.SyncFieldsDuringRead(ctx, existingState.FederationPolicy)
}

func (r *FederationPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *FederationPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, FederationPolicy{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(int64planmodifier.UseStateForUnknown(), "service_principal_id")
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "policy_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks service_principal_federation_policy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FederationPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *FederationPolicyResource) update(ctx context.Context, plan FederationPolicy, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetAccountClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var federation_policy oauth2.FederationPolicy

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &federation_policy)...)
	if diags.HasError() {
		return
	}

	updateRequest := oauth2.UpdateServicePrincipalFederationPolicyRequest{
		Policy:             federation_policy,
		PolicyId:           plan.PolicyId.ValueString(),
		ServicePrincipalId: plan.ServicePrincipalId.ValueInt64(),
		UpdateMask:         "description,oidc_policy",
	}

	response, err := client.ServicePrincipalFederationPolicy.Update(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update service_principal_federation_policy", err.Error())
		return
	}

	var newState FederationPolicy
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *FederationPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan FederationPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var federation_policy oauth2.FederationPolicy

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &federation_policy)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := oauth2.CreateServicePrincipalFederationPolicyRequest{
		Policy:             federation_policy,
		PolicyId:           plan.PolicyId.ValueString(),
		ServicePrincipalId: plan.ServicePrincipalId.ValueInt64(),
	}

	response, err := client.ServicePrincipalFederationPolicy.Create(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create service_principal_federation_policy", err.Error())
		return
	}

	var newState FederationPolicy

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

func (r *FederationPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState FederationPolicy
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest oauth2.GetServicePrincipalFederationPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.ServicePrincipalFederationPolicy.Get(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get service_principal_federation_policy", err.Error())
		return
	}

	var newState FederationPolicy
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *FederationPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan FederationPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *FederationPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state FederationPolicy
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest oauth2.DeleteServicePrincipalFederationPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.ServicePrincipalFederationPolicy.Delete(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete service_principal_federation_policy", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &FederationPolicyResource{}

func (r *FederationPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: service_principal_id,policy_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	servicePrincipalId, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse import identifier", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("service_principal_id"), servicePrincipalId)...)
	policyId := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("policy_id"), policyId)...)
}
