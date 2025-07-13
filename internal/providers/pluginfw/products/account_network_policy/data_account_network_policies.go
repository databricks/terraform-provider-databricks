// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_network_policy

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "account_network_policies"

var _ datasource.DataSourceWithConfigure = &AccountNetworkPoliciesDataSource{}

func DataSourceAccountNetworkPolicies() datasource.DataSource {
	return &AccountNetworkPoliciesDataSource{}
}

type AccountNetworkPoliciesList struct {
	settings_tf.ListNetworkPoliciesRequest
	NetworkPolicies types.List `tfsdk:"items"`
}

func (c AccountNetworkPoliciesList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["items"] = attrs["items"].SetComputed()
	return attrs
}

func (AccountNetworkPoliciesList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(settings_tf.AccountNetworkPolicy{}),
	}
}

type AccountNetworkPoliciesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *AccountNetworkPoliciesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *AccountNetworkPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, AccountNetworkPoliciesList{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks AccountNetworkPolicy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AccountNetworkPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *AccountNetworkPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config AccountNetworkPoliciesList
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest settings.ListNetworkPoliciesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.NetworkPolicies.ListNetworkPoliciesRpcAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list account_network_policies", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var account_network_policy settings_tf.AccountNetworkPolicy
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &account_network_policy)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, account_network_policy.ToObjectValue(ctx))
	}

	var newState AccountNetworkPoliciesList
	newState.NetworkPolicies = types.ListValueMust(settings_tf.AccountNetworkPolicy{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
