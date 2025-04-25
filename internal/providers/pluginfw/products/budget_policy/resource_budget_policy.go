// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package budget_policy

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/billing_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

const resourceName = "budget_policy"

var _ resource.ResourceWithConfigure = &BudgetPolicyResource{}

func ResourceBudgetPolicy() resource.Resource {
	return &BudgetPolicyResource{}
}

type BudgetPolicyResource struct {
	Client *autogen.DatabricksClient
}

func (r *BudgetPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *BudgetPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, billing_tf.BudgetPolicy{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "policy_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks budget_policy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *BudgetPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *BudgetPolicyResource) update(ctx context.Context, plan billing_tf.BudgetPolicy, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetAccountClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var budget_policy billing.BudgetPolicy
	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &budget_policy)...)
	if diags.HasError() {
		return
	}

	var updateRequest = billing.UpdateBudgetPolicyRequest{Policy: &budget_policy}
	updateRequest.PolicyId = plan.PolicyId.ValueString()

	response, err := client.BudgetPolicy.Update(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update budget_policy", err.Error())
		return
	}

	var newState billing_tf.BudgetPolicy
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncEffectiveFieldsDuringCreateOrUpdate(plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *BudgetPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan billing_tf.BudgetPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var budget_policy billing.BudgetPolicy
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &budget_policy)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.BudgetPolicy.Create(ctx, billing.CreateBudgetPolicyRequest{Policy: &budget_policy})
	if err != nil {
		resp.Diagnostics.AddError("failed to create budget_policy", err.Error())
		return
	}

	var newState billing_tf.BudgetPolicy

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncEffectiveFieldsDuringCreateOrUpdate(plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *BudgetPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState billing_tf.BudgetPolicy
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest billing.GetBudgetPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.BudgetPolicy.Get(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get budget_policy", err.Error())
		return
	}

	var newState billing_tf.BudgetPolicy
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncEffectiveFieldsDuringRead(existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *BudgetPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan billing_tf.BudgetPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *BudgetPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state billing_tf.BudgetPolicy
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest billing.DeleteBudgetPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.BudgetPolicy.Delete(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete budget_policy", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &BudgetPolicyResource{}

func (r *BudgetPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: policy_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	policyId := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("policy_id"), policyId)...)
}
