// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_iam_service_principal_v2

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iamv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
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

const resourceName = "account_iam_service_principal_v2"

var _ resource.ResourceWithConfigure = &ServicePrincipalResource{}

func ResourceServicePrincipal() resource.Resource {
	return &ServicePrincipalResource{}
}

type ServicePrincipalResource struct {
	Client *autogen.DatabricksClient
}

// ServicePrincipal extends the main model with additional fields.
type ServicePrincipal struct {
	// The ID of the account that contains the service principal.
	AccountId types.String `tfsdk:"account_id"`
	// Whether the service principal is active in the account.
	AccountSpStatus types.String `tfsdk:"account_sp_status"`
	// The service principal's application ID, a UUID. If you don't set it when
	// you create the service principal, Databricks generates one. It must be
	// unique in the account and can't be changed.
	ApplicationId types.String `tfsdk:"application_id"`
	// The service principal's display name.
	DisplayName types.String `tfsdk:"display_name"`
	// The ID of the service principal in your identity provider (IdP).
	ExternalId types.String `tfsdk:"external_id"`
	// The resource name of the service principal, in one of these formats: -
	// Account APIs:
	// `accounts/{account_id}/service-principals/{service_principal_id}` -
	// Workspace APIs: `service-principals/{service_principal_id}`
	Name types.String `tfsdk:"name"`
	// The resource name of the account that contains the service principal, in
	// the format `accounts/{account_id}`. Workspace APIs return the same
	// format.
	Parent types.String `tfsdk:"parent"`
	// The Databricks ID of the service principal.
	ServicePrincipalId types.String `tfsdk:"service_principal_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ServicePrincipal struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ServicePrincipal) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServicePrincipal
// only implements ToObjectValue() and Type().
func (m ServicePrincipal) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"account_id": m.AccountId,
			"account_sp_status":    m.AccountSpStatus,
			"application_id":       m.ApplicationId,
			"display_name":         m.DisplayName,
			"external_id":          m.ExternalId,
			"name":                 m.Name,
			"parent":               m.Parent,
			"service_principal_id": m.ServicePrincipalId,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ServicePrincipal) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"account_id": types.StringType,
			"account_sp_status":    types.StringType,
			"application_id":       types.StringType,
			"display_name":         types.StringType,
			"external_id":          types.StringType,
			"name":                 types.StringType,
			"parent":               types.StringType,
			"service_principal_id": types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *ServicePrincipal) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServicePrincipal) {
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *ServicePrincipal) SyncFieldsDuringRead(ctx context.Context, from ServicePrincipal) {
}

func (m ServicePrincipal) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_id"] = attrs["account_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["account_sp_status"] = attrs["account_sp_status"].SetRequired()
	attrs["application_id"] = attrs["application_id"].SetComputed()
	attrs["application_id"] = attrs["application_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["application_id"] = attrs["application_id"].SetOptional()
	attrs["application_id"] = attrs["application_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()
	attrs["service_principal_id"] = attrs["service_principal_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

func (r *ServicePrincipalResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *ServicePrincipalResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, ServicePrincipal{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks account_iam_service_principal_v2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ServicePrincipalResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *ServicePrincipalResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ServicePrincipal
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var service_principal iamv2.ServicePrincipal

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &service_principal)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := iamv2.CreateServicePrincipalRequest{
		ServicePrincipal: service_principal,
		Parent:           plan.Parent.ValueString(),
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.IamV2.CreateServicePrincipal(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create account_iam_service_principal_v2", err.Error())
		return
	}

	var newState ServicePrincipal

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

func (r *ServicePrincipalResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState ServicePrincipal
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest iamv2.GetServicePrincipalRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.IamV2.GetServicePrincipal(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get account_iam_service_principal_v2", err.Error())
		return
	}

	var newState ServicePrincipal
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *ServicePrincipalResource) update(ctx context.Context, plan ServicePrincipal, diags *diag.Diagnostics, state *tfsdk.State) {
	var service_principal iamv2.ServicePrincipal

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &service_principal)...)
	if diags.HasError() {
		return
	}

	updateRequest := iamv2.UpdateServicePrincipalRequest{
		ServicePrincipal: service_principal,
		Name:             plan.Name.ValueString(),
		UpdateMask:       "account_sp_status,display_name,external_id",
	}

	client, clientDiags := r.Client.GetAccountClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.IamV2.UpdateServicePrincipal(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update account_iam_service_principal_v2", err.Error())
		return
	}

	var newState ServicePrincipal

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *ServicePrincipalResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ServicePrincipal
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *ServicePrincipalResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state ServicePrincipal
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest iamv2.DeleteServicePrincipalRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.IamV2.DeleteServicePrincipal(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete account_iam_service_principal_v2", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &ServicePrincipalResource{}

func (r *ServicePrincipalResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: name. Got: %q",
				req.ID,
			),
		)
		return
	}

	name := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
}
