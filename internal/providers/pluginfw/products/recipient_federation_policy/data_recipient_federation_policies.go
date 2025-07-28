// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package recipient_federation_policy

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/sharing_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "recipient_federation_policies"

var _ datasource.DataSourceWithConfigure = &FederationPoliciesDataSource{}

func DataSourceFederationPolicies() datasource.DataSource {
	return &FederationPoliciesDataSource{}
}

type FederationPoliciesList struct {
	sharing_tf.ListFederationPoliciesRequest
	RecipientFederationPolicies types.List `tfsdk:"policies"`
}

func (c FederationPoliciesList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policies"] = attrs["policies"].SetComputed()
	return attrs
}

func (FederationPoliciesList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(sharing_tf.FederationPolicy{}),
	}
}

type FederationPoliciesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *FederationPoliciesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *FederationPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, FederationPoliciesList{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks FederationPolicy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FederationPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *FederationPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config FederationPoliciesList
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest sharing.ListFederationPoliciesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.RecipientFederationPolicies.ListAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list recipient_federation_policies", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var federation_policy sharing_tf.FederationPolicy
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &federation_policy)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, federation_policy.ToObjectValue(ctx))
	}

	var newState FederationPoliciesList
	newState.RecipientFederationPolicies = types.ListValueMust(sharing_tf.FederationPolicy{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
