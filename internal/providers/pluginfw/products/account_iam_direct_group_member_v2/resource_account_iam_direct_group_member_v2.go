// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_iam_direct_group_member_v2

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iamv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "account_iam_direct_group_member_v2"

var _ resource.ResourceWithConfigure = &DirectGroupMemberResource{}

func ResourceDirectGroupMember() resource.Resource {
	return &DirectGroupMemberResource{}
}

type DirectGroupMemberResource struct {
	Client *autogen.DatabricksClient
}

// DirectGroupMember extends the main model with additional fields.
type DirectGroupMember struct {
	// Display name of the principal.
	DisplayName types.String `tfsdk:"display_name"`
	// The external ID of the principal in Databricks.
	ExternalId types.String `tfsdk:"external_id"`
	// The internal ID of the group this member belongs to.
	GroupId types.Int64 `tfsdk:"group_id"`
	// The source of group membership (internal or from identity provider).
	MembershipSource types.String `tfsdk:"membership_source"`
	// Internal ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// The type of the principal (user/service principal/group).
	PrincipalType types.String `tfsdk:"principal_type"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// DirectGroupMember struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m DirectGroupMember) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectGroupMember
// only implements ToObjectValue() and Type().
func (m DirectGroupMember) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"display_name": m.DisplayName,
			"external_id":       m.ExternalId,
			"group_id":          m.GroupId,
			"membership_source": m.MembershipSource,
			"principal_id":      m.PrincipalId,
			"principal_type":    m.PrincipalType,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m DirectGroupMember) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"display_name": types.StringType,
			"external_id":       types.StringType,
			"group_id":          types.Int64Type,
			"membership_source": types.StringType,
			"principal_id":      types.Int64Type,
			"principal_type":    types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *DirectGroupMember) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DirectGroupMember) {
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *DirectGroupMember) SyncFieldsDuringRead(ctx context.Context, from DirectGroupMember) {
}

func (m DirectGroupMember) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["display_name"] = attrs["display_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["external_id"] = attrs["external_id"].SetComputed()
	attrs["external_id"] = attrs["external_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["group_id"] = attrs["group_id"].SetComputed()
	attrs["group_id"] = attrs["group_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["group_id"] = attrs["group_id"].SetOptional()
	attrs["group_id"] = attrs["group_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["membership_source"] = attrs["membership_source"].SetComputed()
	attrs["membership_source"] = attrs["membership_source"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["principal_id"] = attrs["principal_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["principal_type"] = attrs["principal_type"].SetComputed()
	attrs["principal_type"] = attrs["principal_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	attrs["group_id"] = attrs["group_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["principal_id"] = attrs["principal_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

func (r *DirectGroupMemberResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *DirectGroupMemberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, DirectGroupMember{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks account_iam_direct_group_member_v2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DirectGroupMemberResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *DirectGroupMemberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan DirectGroupMember
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var direct_group_member iamv2.DirectGroupMember

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &direct_group_member)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := iamv2.CreateDirectGroupMemberRequest{
		DirectGroupMember: direct_group_member,
		GroupId:           plan.GroupId.ValueInt64(),
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.IamV2.CreateDirectGroupMember(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create account_iam_direct_group_member_v2", err.Error())
		return
	}

	var newState DirectGroupMember

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

func (r *DirectGroupMemberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState DirectGroupMember
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest iamv2.GetDirectGroupMemberRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.IamV2.GetDirectGroupMember(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get account_iam_direct_group_member_v2", err.Error())
		return
	}

	var newState DirectGroupMember
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *DirectGroupMemberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// VARIANT_IMMUTABLE resources do not support updates - all changes require replacement.
	resp.Diagnostics.AddError("Update not supported", "This resource does not support updates. All changes require replacement.")
}

func (r *DirectGroupMemberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state DirectGroupMember
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest iamv2.DeleteDirectGroupMemberRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.IamV2.DeleteDirectGroupMember(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete account_iam_direct_group_member_v2", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &DirectGroupMemberResource{}

func (r *DirectGroupMemberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: group_id,principal_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	groupId, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse import identifier", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("group_id"), groupId)...)
	principalId, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse import identifier", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("principal_id"), principalId)...)
}
