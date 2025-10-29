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

// BudgetPoliciesData extends the main model with additional fields.
type BudgetPoliciesData struct {
	BudgetPolicy types.List `tfsdk:"policies"`
	// A filter to apply to the list of policies.
	FilterBy types.Object `tfsdk:"filter_by"`
	// The maximum number of budget policies to return. If unspecified, at most
	// 100 budget policies will be returned. The maximum value is 1000; values
	// above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"page_size"`
	// The sort specification.
	SortSpec types.Object `tfsdk:"sort_spec"`
}

func (BudgetPoliciesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies":  reflect.TypeOf(BudgetPolicyData{}),
		"filter_by": reflect.TypeOf(billing_tf.Filter{}),
		"sort_spec": reflect.TypeOf(billing_tf.SortSpec{}),
	}
}

func (m BudgetPoliciesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["filter_by"] = attrs["filter_by"].SetOptional()
	attrs["sort_spec"] = attrs["sort_spec"].SetOptional()

	attrs["policies"] = attrs["policies"].SetComputed()
	return attrs
}

type BudgetPoliciesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *BudgetPoliciesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *BudgetPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, BudgetPoliciesData{}, nil)
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

	var config BudgetPoliciesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest billing.ListBudgetPoliciesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.BudgetPolicy.ListAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list budget_policies", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var budget_policy BudgetPolicyData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &budget_policy)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, budget_policy.ToObjectValue(ctx))
	}

	config.BudgetPolicy = types.ListValueMust(BudgetPolicyData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
