package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

const resourceName = "function"

var _ resource.ResourceWithConfigure = &FunctionResource{}

func ResourceFunction() resource.Resource {
	return &FunctionResource{}
}

type FunctionResource struct {
	Client *common.DatabricksClient
}

func (r *FunctionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

func (r *FunctionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, catalog_tf.FunctionInfo{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetRequired("name")
		c.SetRequired("catalog_name")
		c.SetRequired("schema_name")
		c.SetRequired("input_params")
		c.SetRequired("data_type")
		c.SetRequired("full_data_type")
		c.SetRequired("routine_defintion")
		c.SetRequired("is_deterministic")
		c.SetRequired("is_null_call")
		c.SetRequired("specific_name")
		c.SetRequired("routine_body")
		c.SetRequired("security_type")
		c.SetRequired("sql_data_access")
		c.SetRequired("parameter_style")

		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "name")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "catalog_name")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "schema_name")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "input_params")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "data_type")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "full_data_type")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "routine_definition")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "is_deterministic")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "is_null_call")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "specific_name")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "routine_body")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "security_type")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "sql_data_access")
		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "parameter_style")

		c.SetReadOnly("full_name")
		c.SetReadOnly("created_at")
		c.SetReadOnly("created_by")
		c.SetReadOnly("updated_at")
		c.SetReadOnly("updated_by")

		return c
	})

	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Function",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FunctionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if r.Client == nil && req.ProviderData != nil {
		r.Client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (r *FunctionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("full_name"), req, resp)
}

func (r *FunctionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var planFunc catalog_tf.FunctionInfo
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planFunc)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var createReq catalog.CreateFunctionRequest

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, planFunc, &createReq)...)
	if resp.Diagnostics.HasError() {
		return
	}

	funcInfo, err := w.Functions.Create(ctx, createReq)
	if err != nil {
		resp.Diagnostics.AddError("failed to create function", err.Error())
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, funcInfo, &planFunc)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, planFunc)...)
}

func (r *FunctionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var planFunc catalog_tf.FunctionInfo
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planFunc)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var updateReq catalog.UpdateFunction

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, planFunc, &updateReq)...)
	if resp.Diagnostics.HasError() {
		return
	}

	funcInfo, err := w.Functions.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("failed to update function", err.Error())
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, funcInfo, &planFunc)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, planFunc)...)
}

func (r *FunctionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var stateFunc catalog_tf.FunctionInfo

	resp.Diagnostics.Append(req.State.Get(ctx, &stateFunc)...)
	if resp.Diagnostics.HasError() {
		return
	}

	funcName := stateFunc.FullName.ValueString()

	funcInfo, err := w.Functions.GetByName(ctx, funcName)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get function", err.Error())
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, funcInfo, &stateFunc)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, stateFunc)...)
}

func (r *FunctionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteReq catalog_tf.DeleteFunctionRequest
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("full_name"), &deleteReq.Name)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := w.Functions.DeleteByName(ctx, deleteReq.Name.ValueString())
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete function", err.Error())
	}
}
