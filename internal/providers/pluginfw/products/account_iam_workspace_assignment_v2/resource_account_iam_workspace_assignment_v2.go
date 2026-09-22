// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_iam_workspace_assignment_v2

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/iamv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
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

const resourceName = "account_iam_workspace_assignment_v2"

var _ resource.ResourceWithConfigure = &WorkspaceAssignmentResource{}

func ResourceWorkspaceAssignment() resource.Resource {
	return &WorkspaceAssignmentResource{}
}

type WorkspaceAssignmentResource struct {
	Client *autogen.DatabricksClient
}

// WorkspaceAssignment extends the main model with additional fields.
type WorkspaceAssignment struct {
	// The account ID parent of the workspace where the principal is assigned
	AccountId types.String `tfsdk:"account_id"`
	// Every entitlement the principal holds in this workspace, whether granted
	// directly or through group membership. Get responses populate this field.
	// List responses leave it empty.
	EffectiveEntitlements types.Set `tfsdk:"effective_entitlements"`
	// Entitlements granted directly to the principal on this workspace. This is
	// the only client-settable field. Create and update manage exactly this
	// set, including entitlements the principal also holds through a group.
	// List responses leave this field empty. Get a single principal to read its
	// entitlements.
	Entitlements types.Set `tfsdk:"entitlements"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// The type of the principal (user/service principal/group) that is
	// assigned.
	PrincipalType types.String `tfsdk:"principal_type"`
	// The workspace ID where the principal is assigned.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// WorkspaceAssignment struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m WorkspaceAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_entitlements": reflect.TypeOf(types.String{}),
		"entitlements":           reflect.TypeOf(types.String{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAssignment
// only implements ToObjectValue() and Type().
func (m WorkspaceAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"account_id": m.AccountId,
			"effective_entitlements": m.EffectiveEntitlements,
			"entitlements":           m.Entitlements,
			"principal_id":           m.PrincipalId,
			"principal_type":         m.PrincipalType,
			"workspace_id":           m.WorkspaceId,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m WorkspaceAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"account_id": types.StringType,
			"effective_entitlements": basetypes.SetType{
				ElemType: types.StringType,
			},
			"entitlements": basetypes.SetType{
				ElemType: types.StringType,
			},
			"principal_id":   types.Int64Type,
			"principal_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *WorkspaceAssignment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceAssignment) {
	if !from.EffectiveEntitlements.IsNull() && !from.EffectiveEntitlements.IsUnknown() && to.EffectiveEntitlements.IsNull() && len(from.EffectiveEntitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveEntitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveEntitlements = from.EffectiveEntitlements
	}
	if !from.Entitlements.IsNull() && !from.Entitlements.IsUnknown() && to.Entitlements.IsNull() && len(from.Entitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entitlements = from.Entitlements
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *WorkspaceAssignment) SyncFieldsDuringRead(ctx context.Context, from WorkspaceAssignment) {
	if !from.EffectiveEntitlements.IsNull() && !from.EffectiveEntitlements.IsUnknown() && to.EffectiveEntitlements.IsNull() && len(from.EffectiveEntitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveEntitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveEntitlements = from.EffectiveEntitlements
	}
	if !from.Entitlements.IsNull() && !from.Entitlements.IsUnknown() && to.Entitlements.IsNull() && len(from.Entitlements.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entitlements, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entitlements = from.Entitlements
	}
}

func (m WorkspaceAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["effective_entitlements"] = attrs["effective_entitlements"].SetComputed()
	attrs["entitlements"] = attrs["entitlements"].SetOptional()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["principal_type"] = attrs["principal_type"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["principal_id"] = attrs["principal_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetEffectiveEntitlements returns the value of the EffectiveEntitlements field in WorkspaceAssignment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceAssignment) GetEffectiveEntitlements(ctx context.Context) ([]types.String, bool) {
	if m.EffectiveEntitlements.IsNull() || m.EffectiveEntitlements.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.EffectiveEntitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveEntitlements sets the value of the EffectiveEntitlements field in WorkspaceAssignment.
func (m *WorkspaceAssignment) SetEffectiveEntitlements(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveEntitlements = types.SetValueMust(t, vs)
}

// GetEntitlements returns the value of the Entitlements field in WorkspaceAssignment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceAssignment) GetEntitlements(ctx context.Context) ([]types.String, bool) {
	if m.Entitlements.IsNull() || m.Entitlements.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Entitlements.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntitlements sets the value of the Entitlements field in WorkspaceAssignment.
func (m *WorkspaceAssignment) SetEntitlements(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["entitlements"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Entitlements = types.SetValueMust(t, vs)
}

func (r *WorkspaceAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *WorkspaceAssignmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, WorkspaceAssignment{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks account_iam_workspace_assignment_v2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *WorkspaceAssignmentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *WorkspaceAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan WorkspaceAssignment
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var workspace_assignment iamv2.WorkspaceAssignment

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &workspace_assignment)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := iamv2.CreateWorkspaceAssignmentRequest{
		WorkspaceAssignment: workspace_assignment,
		WorkspaceId:         plan.WorkspaceId.ValueInt64(),
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.IamV2.CreateWorkspaceAssignment(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create account_iam_workspace_assignment_v2", err.Error())
		return
	}

	var newState WorkspaceAssignment

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

func (r *WorkspaceAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState WorkspaceAssignment
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest iamv2.GetWorkspaceAssignmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.IamV2.GetWorkspaceAssignment(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get account_iam_workspace_assignment_v2", err.Error())
		return
	}

	var newState WorkspaceAssignment
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *WorkspaceAssignmentResource) update(ctx context.Context, plan WorkspaceAssignment, diags *diag.Diagnostics, state *tfsdk.State) {
	var workspace_assignment iamv2.WorkspaceAssignment

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &workspace_assignment)...)
	if diags.HasError() {
		return
	}

	updateRequest := iamv2.UpdateWorkspaceAssignmentRequest{
		WorkspaceAssignment: workspace_assignment,
		PrincipalId:         plan.PrincipalId.ValueInt64(),
		WorkspaceId:         plan.WorkspaceId.ValueInt64(),
		UpdateMask:          *fieldmask.New(strings.Split("entitlements", ",")),
	}

	client, clientDiags := r.Client.GetAccountClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.IamV2.UpdateWorkspaceAssignment(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update account_iam_workspace_assignment_v2", err.Error())
		return
	}

	var newState WorkspaceAssignment

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *WorkspaceAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan WorkspaceAssignment
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *WorkspaceAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state WorkspaceAssignment
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest iamv2.DeleteWorkspaceAssignmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.IamV2.DeleteWorkspaceAssignment(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete account_iam_workspace_assignment_v2", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &WorkspaceAssignmentResource{}

func (r *WorkspaceAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: workspace_id,principal_id. Got: %q",
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
	principalId, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse import identifier", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("principal_id"), principalId)...)
}
