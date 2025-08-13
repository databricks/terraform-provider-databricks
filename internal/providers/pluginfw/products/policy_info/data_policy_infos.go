// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package policy_info

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "policy_infos"

var _ datasource.DataSourceWithConfigure = &PolicyInfosDataSource{}

func DataSourcePolicyInfos() datasource.DataSource {
	return &PolicyInfosDataSource{}
}

type PolicyInfosList struct {
	catalog_tf.ListPoliciesRequest
	Policies types.List `tfsdk:"policies"`
}

func (c PolicyInfosList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policies"] = attrs["policies"].SetComputed()
	return attrs
}

func (PolicyInfosList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(catalog_tf.PolicyInfo{}),
	}
}

type PolicyInfosDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *PolicyInfosDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *PolicyInfosDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, PolicyInfosList{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks PolicyInfo",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *PolicyInfosDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *PolicyInfosDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PolicyInfosList
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest catalog.ListPoliciesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Policies.ListPoliciesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list policy_infos", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var policy_info catalog_tf.PolicyInfo
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &policy_info)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, policy_info.ToObjectValue(ctx))
	}

	var newState PolicyInfosList
	newState.Policies = types.ListValueMust(catalog_tf.PolicyInfo{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
