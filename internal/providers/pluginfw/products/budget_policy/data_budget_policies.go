// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package budget_policy

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/billing_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "budget_policies"

var _ datasource.DataSourceWithConfigure = &BudgetPoliciesDataSource{}

func DataSourceBudgetPolicies() datasource.DataSource {
	return &BudgetPoliciesDataSource{}
}

type BudgetPoliciesList struct {
	billing_tf.ListBudgetPoliciesRequest
	BudgetPolicy types.List `tfsdk:"budget_policies"`
}

func (c BudgetPoliciesList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policies"] = attrs["budget_policies"].SetComputed()
	return attrs
}

func (BudgetPoliciesList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"budget_policies": reflect.TypeOf(billing_tf.BudgetPolicy{}),
	}
}

type BudgetPoliciesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *BudgetPoliciesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *BudgetPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, BudgetPoliciesList{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks BudgetPolicy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *BudgetPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *BudgetPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config BudgetPoliciesList
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest billing.ListBudgetPoliciesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.BudgetPolicy.ListAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list budget_policies", err.Error())
		return
	}

	var budget_policies = []attr.Value{}
	for _, item := range response {
		var budget_policy billing_tf.BudgetPolicy
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &budget_policy)...)
		if resp.Diagnostics.HasError() {
			return
		}
		budget_policies = append(budget_policies, budget_policy.ToObjectValue(ctx))
	}

	var newState BudgetPoliciesList
	newState.BudgetPolicy = types.ListValueMust(billing_tf.BudgetPolicy{}.Type(ctx), budget_policies)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
