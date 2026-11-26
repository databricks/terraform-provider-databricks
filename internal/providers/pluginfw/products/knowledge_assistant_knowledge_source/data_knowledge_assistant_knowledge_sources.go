// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package knowledge_assistant_knowledge_source

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/knowledgeassistants"
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

const dataSourcesName = "knowledge_assistant_knowledge_sources"

var _ datasource.DataSourceWithConfigure = &KnowledgeSourcesDataSource{}

func DataSourceKnowledgeSources() datasource.DataSource {
	return &KnowledgeSourcesDataSource{}
}

// KnowledgeSourcesData extends the main model with additional fields.
type KnowledgeSourcesData struct {
	KnowledgeAssistants types.List `tfsdk:"knowledge_sources"`

	PageSize types.Int64 `tfsdk:"page_size"`
	// Parent resource to list from. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent             types.String `tfsdk:"parent"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (KnowledgeSourcesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_sources": reflect.TypeOf(KnowledgeSourceData{}),
		"provider_config":   reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m KnowledgeSourcesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["knowledge_sources"] = attrs["knowledge_sources"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type KnowledgeSourcesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *KnowledgeSourcesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *KnowledgeSourcesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, KnowledgeSourcesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks KnowledgeSource",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *KnowledgeSourcesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *KnowledgeSourcesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config KnowledgeSourcesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest knowledgeassistants.ListKnowledgeSourcesRequest
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

	response, err := client.KnowledgeAssistants.ListKnowledgeSourcesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list knowledge_assistant_knowledge_sources", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var knowledge_source KnowledgeSourceData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &knowledge_source)...)
		if resp.Diagnostics.HasError() {
			return
		}
		knowledge_source.ProviderConfigData = config.ProviderConfigData

		results = append(results, knowledge_source.ToObjectValue(ctx))
	}

	config.KnowledgeAssistants = types.ListValueMust(KnowledgeSourceData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
