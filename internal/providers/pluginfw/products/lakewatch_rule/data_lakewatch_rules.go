// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_rule

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourcesName = "lakewatch_rules"

var _ datasource.DataSourceWithConfigure = &RulesDataSource{}

func DataSourceRules() datasource.DataSource {
	return &RulesDataSource{}
}

// RulesData extends the main model with additional fields.
type RulesData struct {
	Lakewatch types.List `tfsdk:"rules"`
	// Name of the Unity Catalog catalog to list rules from. Required.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Filter expression that restricts the results, such as `rule_name =
	// "my_rule"` to return only the rule with that name. The name is matched
	// against the current UC name, which reads report in `effective_rule_name`.
	Filter types.String `tfsdk:"filter"`
	// Maximum number of rules to return. The server may return fewer. If unset
	// or 0, the server picks a default page size.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Name of a Unity Catalog schema in `catalog_name`. If set, only rules in
	// this schema are returned. If unset, rules in all schemas of the catalog
	// are returned.
	SchemaName types.String `tfsdk:"schema_name"`
	// How much of each rule to return. Defaults to `RULE_VIEW_FULL`.
	// `RULE_VIEW_BASIC` returns only each rule's identity and Unity Catalog
	// location.
	View               types.String `tfsdk:"view"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (RulesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rules":           reflect.TypeOf(RuleData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m RulesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["view"] = attrs["view"].SetOptional()

	attrs["rules"] = attrs["rules"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type RulesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *RulesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *RulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, RulesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Rule",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *RulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *RulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config RulesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest lakewatch.ListRulesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfigData
	resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Lakewatch.ListRulesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list lakewatch_rules", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var rule RuleData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &rule)...)
		if resp.Diagnostics.HasError() {
			return
		}
		rule.ProviderConfigData = config.ProviderConfigData

		results = append(results, rule.ToObjectValue(ctx))
	}

	config.Lakewatch = types.ListValueMust(RuleData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
