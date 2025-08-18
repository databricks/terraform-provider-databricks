// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package policy_info

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

const resourceName = "policy_info"

var _ resource.ResourceWithConfigure = &PolicyInfoResource{}

func ResourcePolicyInfo() resource.Resource {
	return &PolicyInfoResource{}
}

type PolicyInfoResource struct {
	Client *autogen.DatabricksClient
}

func (r *PolicyInfoResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *PolicyInfoResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, catalog_tf.PolicyInfo{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "on_securable_type")
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "on_securable_fullname")
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "name")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks policy_info",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *PolicyInfoResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *PolicyInfoResource) update(ctx context.Context, plan catalog_tf.PolicyInfo, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var policy_info catalog.PolicyInfo

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &policy_info)...)
	if diags.HasError() {
		return
	}

	updateRequest := catalog.UpdatePolicyRequest{
		PolicyInfo:          policy_info,
		Name:                plan.Name.ValueString(),
		OnSecurableFullname: plan.OnSecurableFullname.ValueString(),
		OnSecurableType:     plan.OnSecurableType.ValueString(),
		UpdateMask:          "column_mask,comment,except_principals,for_securable_type,match_columns,policy_type,row_filter,to_principals,when_condition",
	}

	response, err := client.Policies.UpdatePolicy(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update policy_info", err.Error())
		return
	}

	var newState catalog_tf.PolicyInfo
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *PolicyInfoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan catalog_tf.PolicyInfo
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var policy_info catalog.PolicyInfo

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &policy_info)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := catalog.CreatePolicyRequest{
		PolicyInfo: policy_info,
	}

	response, err := client.Policies.CreatePolicy(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create policy_info", err.Error())
		return
	}

	var newState catalog_tf.PolicyInfo

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

func (r *PolicyInfoResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState catalog_tf.PolicyInfo
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Policies.GetPolicy(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get policy_info", err.Error())
		return
	}

	var newState catalog_tf.PolicyInfo
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *PolicyInfoResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan catalog_tf.PolicyInfo
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *PolicyInfoResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state catalog_tf.PolicyInfo
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest catalog.DeletePolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.Policies.DeletePolicy(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete policy_info", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &PolicyInfoResource{}

func (r *PolicyInfoResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: on_securable_type,on_securable_fullname,name. Got: %q",
				req.ID,
			),
		)
		return
	}

	onSecurableType := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("on_securable_type"), onSecurableType)...)
	onSecurableFullname := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("on_securable_fullname"), onSecurableFullname)...)
	name := parts[2]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
}
