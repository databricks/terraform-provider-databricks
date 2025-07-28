// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clean_room_auto_approval_rule

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

const resourceName = "clean_room_auto_approval_rule"

var _ resource.ResourceWithConfigure = &CleanRoomAutoApprovalRuleResource{}

func ResourceCleanRoomAutoApprovalRule() resource.Resource {
	return &CleanRoomAutoApprovalRuleResource{}
}

type CleanRoomAutoApprovalRuleResource struct {
	Client *autogen.DatabricksClient
}

func (r *CleanRoomAutoApprovalRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *CleanRoomAutoApprovalRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, cleanrooms_tf.CleanRoomAutoApprovalRule{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "rule_id")
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

func (r *CleanRoomAutoApprovalRuleResource) update(ctx context.Context, plan cleanrooms_tf.CleanRoomAutoApprovalRule, diags *diag.Diagnostics, state *tfsdk.State) {
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

	var newState cleanrooms_tf.CleanRoomAutoApprovalRule
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *CleanRoomAutoApprovalRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan cleanrooms_tf.CleanRoomAutoApprovalRule
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

	var newState cleanrooms_tf.CleanRoomAutoApprovalRule

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

func (r *CleanRoomAutoApprovalRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState cleanrooms_tf.CleanRoomAutoApprovalRule
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

	var newState cleanrooms_tf.CleanRoomAutoApprovalRule

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *CleanRoomAutoApprovalRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan cleanrooms_tf.CleanRoomAutoApprovalRule
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

	var state cleanrooms_tf.CleanRoomAutoApprovalRule
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
