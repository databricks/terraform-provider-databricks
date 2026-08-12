// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package environments_workspace_base_environment

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/environments"
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

const dataSourcesName = "environments_workspace_base_environments"

var _ datasource.DataSourceWithConfigure = &WorkspaceBaseEnvironmentsDataSource{}

func DataSourceWorkspaceBaseEnvironments() datasource.DataSource {
	return &WorkspaceBaseEnvironmentsDataSource{}
}

// WorkspaceBaseEnvironmentsData extends the main model with additional fields.
type WorkspaceBaseEnvironmentsData struct {
	Environments types.List `tfsdk:"workspace_base_environments"`
	// The maximum number of environments to return per page. Default is 1000.
	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (WorkspaceBaseEnvironmentsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_base_environments": reflect.TypeOf(WorkspaceBaseEnvironmentData{}),
		"provider_config":             reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m WorkspaceBaseEnvironmentsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["workspace_base_environments"] = attrs["workspace_base_environments"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type WorkspaceBaseEnvironmentsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *WorkspaceBaseEnvironmentsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *WorkspaceBaseEnvironmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, WorkspaceBaseEnvironmentsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks WorkspaceBaseEnvironment",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *WorkspaceBaseEnvironmentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *WorkspaceBaseEnvironmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config WorkspaceBaseEnvironmentsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest environments.ListWorkspaceBaseEnvironmentsRequest
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

	response, err := client.Environments.ListWorkspaceBaseEnvironmentsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list environments_workspace_base_environments", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var workspace_base_environment WorkspaceBaseEnvironmentData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &workspace_base_environment)...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspace_base_environment.ProviderConfigData = config.ProviderConfigData

		results = append(results, workspace_base_environment.ToObjectValue(ctx))
	}

	config.Environments = types.ListValueMust(WorkspaceBaseEnvironmentData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
