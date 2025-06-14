package serviceprincipalfederationpolicy

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/oauth2_tf"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const resourceName = "service_principal_federation_policy"

type resourceServicePrincipalFederationPolicy struct {
	client *common.DatabricksClient
}

type servicePrincipalFederationPolicyResource struct {
	oauth2_tf.FederationPolicy
	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id"`
}

func (rspfp *resourceServicePrincipalFederationPolicy) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	ac, diags := rspfp.client.GetAccountClient()
	response.Diagnostics.Append(diags...)

	var spfpr servicePrincipalFederationPolicyResource
	response.Diagnostics.Append(request.Plan.Get(ctx, &spfpr)...)
	if response.Diagnostics.HasError() {
		return
	}
	var federationPolicyGoSdk oauth2.FederationPolicy
	response.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, spfpr, &federationPolicyGoSdk)...)
	if response.Diagnostics.HasError() {
		return
	}

	finalSpfpr, err := ac.ServicePrincipalFederationPolicy.Create(ctx, oauth2.CreateServicePrincipalFederationPolicyRequest{
		Policy:             &federationPolicyGoSdk,
		ServicePrincipalId: spfpr.ServicePrincipalId.ValueInt64(),
	})
	if err != nil {
		response.Diagnostics.AddError("failed to create ServicePrincipalFederationPolicy", err.Error())
		return
	}

	var newSpfpr servicePrincipalFederationPolicyResource
	response.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, finalSpfpr, &newSpfpr)...)
	if response.Diagnostics.HasError() {
		return
	}
	newSpfpr.ServicePrincipalId = spfpr.ServicePrincipalId
	response.Diagnostics.Append(response.State.Set(ctx, newSpfpr)...)
	if response.Diagnostics.HasError() {
		return
	}

}

func (rspfp *resourceServicePrincipalFederationPolicy) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {

	//panic("read")

	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	ac, diags := rspfp.client.GetAccountClient()
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	var spfpr servicePrincipalFederationPolicyResource
	response.Diagnostics.Append(request.State.Get(ctx, &spfpr)...)
	if response.Diagnostics.HasError() {
		return
	}

	spfpGoSdk, err := ac.ServicePrincipalFederationPolicy.GetByServicePrincipalIdAndPolicyId(ctx,
		spfpr.ServicePrincipalId.ValueInt64(), getLastPartOfName(&spfpr))
	if err != nil {
		response.Diagnostics.AddError("failed to read servicePrincipalFederationPolicy", err.Error())
		return
	}

	var newSpfpr servicePrincipalFederationPolicyResource
	response.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, spfpGoSdk, &newSpfpr)...)
	if response.Diagnostics.HasError() {
		return
	}
	newSpfpr.ServicePrincipalId = spfpr.ServicePrincipalId
	response.Diagnostics.Append(response.State.Set(ctx, newSpfpr)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func getLastPartOfName(d *servicePrincipalFederationPolicyResource) string {
	name := d.Name.ValueString()
	return name[strings.LastIndex(name, "/")+1:]
}

func (rspfp *resourceServicePrincipalFederationPolicy) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	ac, diags := rspfp.client.GetAccountClient()
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	var spfpr servicePrincipalFederationPolicyResource
	response.Diagnostics.Append(request.Plan.Get(ctx, &spfpr)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Update the app
	var federationPolicyGoSdk oauth2.FederationPolicy
	response.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, spfpr, &federationPolicyGoSdk)...)
	if response.Diagnostics.HasError() {
		return
	}
	updatedFederationPolicyGoSdk, err := ac.ServicePrincipalFederationPolicy.Update(ctx,
		oauth2.UpdateServicePrincipalFederationPolicyRequest{
			Policy:             &federationPolicyGoSdk,
			PolicyId:           getLastPartOfName(&spfpr),
			ServicePrincipalId: spfpr.ServicePrincipalId.ValueInt64(),
		},
	)
	if err != nil {
		response.Diagnostics.AddError("failed to update app", err.Error())
		return
	}

	var newSpfpr servicePrincipalFederationPolicyResource
	response.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, updatedFederationPolicyGoSdk, &newSpfpr)...)
	if response.Diagnostics.HasError() {
		return
	}

	newSpfpr.ServicePrincipalId = spfpr.ServicePrincipalId
	response.Diagnostics.Append(response.State.Set(ctx, newSpfpr)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (rspfp *resourceServicePrincipalFederationPolicy) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	ac, diags := rspfp.client.GetAccountClient()
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	var spfpr servicePrincipalFederationPolicyResource
	response.Diagnostics.Append(request.State.Get(ctx, &spfpr)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := ac.ServicePrincipalFederationPolicy.DeleteByServicePrincipalIdAndPolicyId(
		ctx,
		spfpr.ServicePrincipalId.ValueInt64(),
		getLastPartOfName(&spfpr),
	)
	if err != nil && !apierr.IsMissing(err) {
		response.Diagnostics.AddError("failed to delete ServicePrincipalFederationPolicy", err.Error())
		return
	}
}

func ResourceServicePrincipalFederationPolicyResource() resource.Resource {
	return &resourceServicePrincipalFederationPolicy{}
}

func (rspfp *resourceServicePrincipalFederationPolicy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, servicePrincipalFederationPolicyResource{}, func(cs tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		cs.AddPlanModifier(int64planmodifier.RequiresReplace(), "service_principal_id")
		cs.AddPlanModifier(stringplanmodifier.RequiresReplace(), "name")
		cs.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "name")
		cs.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "create_time")
		cs.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "uid")
		cs.SetComputed("name")
		cs.SetRequired("oidc_policy")
		cs.SetRequired("service_principal_id")
		return cs
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Service Principal Federation Policy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (rspfp *resourceServicePrincipalFederationPolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}

func (rspfp *resourceServicePrincipalFederationPolicy) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

func (rspfp *resourceServicePrincipalFederationPolicy) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if rspfp.client == nil && req.ProviderData != nil {
		rspfp.client = pluginfwcommon.ConfigureResource(req, resp)
	}
}
