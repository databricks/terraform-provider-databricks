// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package knowledge_assistant

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

const dataSourcesName = "knowledge_assistants"

var _ datasource.DataSourceWithConfigure = &KnowledgeAssistantsDataSource{}

func DataSourceKnowledgeAssistants() datasource.DataSource {
	return &KnowledgeAssistantsDataSource{}
}

// KnowledgeAssistantsData extends the main model with additional fields.
type KnowledgeAssistantsData struct {
	KnowledgeAssistants types.List `tfsdk:"knowledge_assistants"`
	// The maximum number of knowledge assistants to return. If unspecified, at
	// most 100 knowledge assistants will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (KnowledgeAssistantsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_assistants": reflect.TypeOf(KnowledgeAssistantData{}),
		"provider_config":      reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m KnowledgeAssistantsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["knowledge_assistants"] = attrs["knowledge_assistants"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type KnowledgeAssistantsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *KnowledgeAssistantsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *KnowledgeAssistantsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, KnowledgeAssistantsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks KnowledgeAssistant",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *KnowledgeAssistantsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *KnowledgeAssistantsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config KnowledgeAssistantsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest knowledgeassistants.ListKnowledgeAssistantsRequest
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

	response, err := client.KnowledgeAssistants.ListKnowledgeAssistantsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list knowledge_assistants", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var knowledge_assistant KnowledgeAssistantData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &knowledge_assistant)...)
		if resp.Diagnostics.HasError() {
			return
		}
		knowledge_assistant.ProviderConfigData = config.ProviderConfigData

		results = append(results, knowledge_assistant.ToObjectValue(ctx))
	}

	config.KnowledgeAssistants = types.ListValueMust(KnowledgeAssistantData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
