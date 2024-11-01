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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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

// TODO: Update as needed to fit the requirements of the resource.
func (r *FunctionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(catalog_tf.FunctionInfo{}, nil)

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

func AppendDiagAndCheckErrors(resp *diag.Diagnostics, diags diag.Diagnostics) bool {
	resp.Append(diags...)
	return resp.HasError()
}

func (r *FunctionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

}

func (r *FunctionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := r.Client.GetWorkspaceClient()
	if AppendDiagAndCheckErrors(&resp.Diagnostics, diags) {
		return
	}

	var planFunc catalog_tf.FunctionInfo
	if AppendDiagAndCheckErrors(&resp.Diagnostics, req.Plan.Get(ctx, &planFunc)) {
		return
	}

	var updateReq catalog.UpdateFunction

	if AppendDiagAndCheckErrors(&resp.Diagnostics, converters.TfSdkToGoSdkStruct(ctx, planFunc, &updateReq)) {
		return
	}

	funcInfo, err := w.Functions.Update(ctx, updateReq)
	if err != nil {
		resp.Diagnostics.AddError("failed to update function", err.Error())
	}

	if AppendDiagAndCheckErrors(&resp.Diagnostics, converters.GoSdkToTfSdkStruct(ctx, funcInfo, &planFunc)) {
		return
	}

	if AppendDiagAndCheckErrors(&resp.Diagnostics, resp.State.Set(ctx, funcInfo)) {
		return
	}
}

func (r *FunctionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	w, diags := r.Client.GetWorkspaceClient()
	if AppendDiagAndCheckErrors(&resp.Diagnostics, diags) {
		return
	}

	var stateFunc catalog_tf.FunctionInfo
	if AppendDiagAndCheckErrors(&resp.Diagnostics, req.State.Get(ctx, &stateFunc)) {
		return
	}

	funcName := stateFunc.Name.ValueString()

	funcInfo, err := w.Functions.GetByName(ctx, funcName)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get function", err.Error())
		return
	}

	if AppendDiagAndCheckErrors(&resp.Diagnostics, converters.GoSdkToTfSdkStruct(ctx, funcInfo, &stateFunc)) {
		return
	}

	if AppendDiagAndCheckErrors(&resp.Diagnostics, resp.State.Set(ctx, stateFunc)) {
		return
	}
}

func (r *FunctionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
