// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clean_room_auto_approval_rule

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

const resourceName = "clean_room_auto_approval_rule"

var _ resource.ResourceWithConfigure = &CleanRoomAutoApprovalRuleResource{}

func ResourceCleanRoomAutoApprovalRule() resource.Resource {
	return &CleanRoomAutoApprovalRuleResource{}
}

type CleanRoomAutoApprovalRuleResource struct {
	Client *autogen.DatabricksClient
}

// CleanRoomAutoApprovalRule extends the main model with additional fields.
type CleanRoomAutoApprovalRule struct {
	cleanrooms_tf.CleanRoomAutoApprovalRule
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// CleanRoomAutoApprovalRule struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m CleanRoomAutoApprovalRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.CleanRoomAutoApprovalRule.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAutoApprovalRule
// only implements ToObjectValue() and Type().
func (m CleanRoomAutoApprovalRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.CleanRoomAutoApprovalRule.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()
	embeddedAttrs["workspace_id"] = m.WorkspaceID

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m CleanRoomAutoApprovalRule) Type(ctx context.Context) attr.Type {
	embeddedType := m.CleanRoomAutoApprovalRule.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()
	attrTypes["workspace_id"] = types.StringType

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *CleanRoomAutoApprovalRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan CleanRoomAutoApprovalRule) {
	m.CleanRoomAutoApprovalRule.SyncFieldsDuringCreateOrUpdate(ctx, plan.CleanRoomAutoApprovalRule)
	m.WorkspaceID = plan.WorkspaceID
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *CleanRoomAutoApprovalRule) SyncFieldsDuringRead(ctx context.Context, existingState CleanRoomAutoApprovalRule) {
	m.CleanRoomAutoApprovalRule.SyncFieldsDuringRead(ctx, existingState.CleanRoomAutoApprovalRule)
	m.WorkspaceID = existingState.WorkspaceID
}

func (r *CleanRoomAutoApprovalRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *CleanRoomAutoApprovalRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, CleanRoomAutoApprovalRule{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "rule_id")
		c.SetOptional("workspace_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks clean_room_auto_approval_rule",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CleanRoomAutoApprovalRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *CleanRoomAutoApprovalRuleResource) update(ctx context.Context, plan CleanRoomAutoApprovalRule, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var clean_room_auto_approval_rule cleanrooms.CleanRoomAutoApprovalRule

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &clean_room_auto_approval_rule)...)
	if diags.HasError() {
		return
	}

	updateRequest := cleanrooms.UpdateCleanRoomAutoApprovalRuleRequest{
		AutoApprovalRule: clean_room_auto_approval_rule,
		RuleId:           plan.RuleId.ValueString(),
	}

	response, err := client.CleanRoomAutoApprovalRules.Update(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update clean_room_auto_approval_rule", err.Error())
		return
	}

	var newState CleanRoomAutoApprovalRule
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *CleanRoomAutoApprovalRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan CleanRoomAutoApprovalRule
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var clean_room_auto_approval_rule cleanrooms.CleanRoomAutoApprovalRule

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &clean_room_auto_approval_rule)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := cleanrooms.CreateCleanRoomAutoApprovalRuleRequest{
		AutoApprovalRule: clean_room_auto_approval_rule,
	}

	response, err := client.CleanRoomAutoApprovalRules.Create(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create clean_room_auto_approval_rule", err.Error())
		return
	}

	var newState CleanRoomAutoApprovalRule

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

func (r *CleanRoomAutoApprovalRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState CleanRoomAutoApprovalRule
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest cleanrooms.GetCleanRoomAutoApprovalRuleRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.CleanRoomAutoApprovalRules.Get(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get clean_room_auto_approval_rule", err.Error())
		return
	}

	var newState CleanRoomAutoApprovalRule
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *CleanRoomAutoApprovalRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan CleanRoomAutoApprovalRule
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *CleanRoomAutoApprovalRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CleanRoomAutoApprovalRule
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest cleanrooms.DeleteCleanRoomAutoApprovalRuleRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.CleanRoomAutoApprovalRules.Delete(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete clean_room_auto_approval_rule", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &CleanRoomAutoApprovalRuleResource{}

func (r *CleanRoomAutoApprovalRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: rule_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	ruleId := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("rule_id"), ruleId)...)
}
