// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mason_session_store

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/mason"
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

const dataSourcesName = "mason_session_stores"

var _ datasource.DataSourceWithConfigure = &SessionStoresDataSource{}

func DataSourceSessionStores() datasource.DataSource {
	return &SessionStoresDataSource{}
}

// SessionStoresData extends the main model with additional fields.
type SessionStoresData struct {
	Mason types.List `tfsdk:"session_stores"`
	// Maximum number of session stores to return. Defaults to 10; must be
	// between 1 and 100.
	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (SessionStoresData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"session_stores":  reflect.TypeOf(SessionStoreData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m SessionStoresData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["session_stores"] = attrs["session_stores"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type SessionStoresDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *SessionStoresDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *SessionStoresDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SessionStoresData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks SessionStore",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SessionStoresDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *SessionStoresDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config SessionStoresData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest mason.ListSessionStoresRequest
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

	response, err := client.Mason.ListSessionStoresAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list mason_session_stores", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var session_store SessionStoreData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &session_store)...)
		if resp.Diagnostics.HasError() {
			return
		}
		session_store.ProviderConfigData = config.ProviderConfigData

		results = append(results, session_store.ToObjectValue(ctx))
	}

	config.Mason = types.ListValueMust(SessionStoreData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
