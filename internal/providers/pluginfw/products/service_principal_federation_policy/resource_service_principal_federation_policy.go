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
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
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
	// Creation time of the federation policy.
	CreateTime types.String `tfsdk:"create_time"`
	// Description of the federation policy.
	Description types.String `tfsdk:"description"`
	// Resource name for the federation policy. Example values include
	// `accounts/<account-id>/federationPolicies/my-federation-policy` for
	// Account Federation Policies, and
	// `accounts/<account-id>/servicePrincipals/<service-principal-id>/federationPolicies/my-federation-policy`
	// for Service Principal Federation Policies. Typically an output parameter,
	// which does not need to be specified in create or update requests. If
	// specified in a request, must match the value in the request URL.
	Name types.String `tfsdk:"name"`

	OidcPolicy types.Object `tfsdk:"oidc_policy"`
	// The ID of the federation policy. Output only.
	PolicyId types.String `tfsdk:"policy_id"`
	// The service principal ID that this federation policy applies to. Output
	// only. Only set for service principal federation policies.
	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id"`
	// Unique, immutable id of the federation policy.
	Uid types.String `tfsdk:"uid"`
	// Last update time of the federation policy.
	UpdateTime types.String `tfsdk:"update_time"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// FederationPolicy struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m FederationPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"oidc_policy": reflect.TypeOf(oauth2_tf.OidcFederationPolicy{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FederationPolicy
// only implements ToObjectValue() and Type().
func (m FederationPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"create_time": m.CreateTime,
			"description":          m.Description,
			"name":                 m.Name,
			"oidc_policy":          m.OidcPolicy,
			"policy_id":            m.PolicyId,
			"service_principal_id": m.ServicePrincipalId,
			"uid":                  m.Uid,
			"update_time":          m.UpdateTime,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m FederationPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": types.StringType,
			"description":          types.StringType,
			"name":                 types.StringType,
			"oidc_policy":          oauth2_tf.OidcFederationPolicy{}.Type(ctx),
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
			"uid":                  types.StringType,
			"update_time":          types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *FederationPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FederationPolicy) {
	if !from.OidcPolicy.IsNull() && !from.OidcPolicy.IsUnknown() {
		if toOidcPolicy, ok := to.GetOidcPolicy(ctx); ok {
			if fromOidcPolicy, ok := from.GetOidcPolicy(ctx); ok {
				// Recursively sync the fields of OidcPolicy
				toOidcPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromOidcPolicy)
				to.SetOidcPolicy(ctx, toOidcPolicy)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *FederationPolicy) SyncFieldsDuringRead(ctx context.Context, from FederationPolicy) {
	if !from.OidcPolicy.IsNull() && !from.OidcPolicy.IsUnknown() {
		if toOidcPolicy, ok := to.GetOidcPolicy(ctx); ok {
			if fromOidcPolicy, ok := from.GetOidcPolicy(ctx); ok {
				toOidcPolicy.SyncFieldsDuringRead(ctx, fromOidcPolicy)
				to.SetOidcPolicy(ctx, toOidcPolicy)
			}
		}
	}
}

func (m FederationPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["oidc_policy"] = attrs["oidc_policy"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetComputed()
	attrs["policy_id"] = attrs["policy_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["policy_id"] = attrs["policy_id"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()
	attrs["service_principal_id"] = attrs["service_principal_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["service_principal_id"] = attrs["service_principal_id"].SetOptional()
	attrs["service_principal_id"] = attrs["service_principal_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["uid"] = attrs["uid"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()

	attrs["service_principal_id"] = attrs["service_principal_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["policy_id"] = attrs["policy_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetOidcPolicy returns the value of the OidcPolicy field in FederationPolicy as
// a oauth2_tf.OidcFederationPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *FederationPolicy) GetOidcPolicy(ctx context.Context) (oauth2_tf.OidcFederationPolicy, bool) {
	var e oauth2_tf.OidcFederationPolicy
	if m.OidcPolicy.IsNull() || m.OidcPolicy.IsUnknown() {
		return e, false
	}
	var v oauth2_tf.OidcFederationPolicy
	d := m.OidcPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOidcPolicy sets the value of the OidcPolicy field in FederationPolicy.
func (m *FederationPolicy) SetOidcPolicy(ctx context.Context, v oauth2_tf.OidcFederationPolicy) {
	vs := v.ToObjectValue(ctx)
	m.OidcPolicy = vs
}

func (r *FederationPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *FederationPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, FederationPolicy{}, nil)
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
