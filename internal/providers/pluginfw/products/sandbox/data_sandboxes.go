// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sandbox

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/sandbox"
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

const dataSourcesName = "sandboxes"

var _ datasource.DataSourceWithConfigure = &SandboxesDataSource{}

func DataSourceSandboxes() datasource.DataSource {
	return &SandboxesDataSource{}
}

// SandboxesData extends the main model with additional fields.
type SandboxesData struct {
	Sandbox types.List `tfsdk:"sandboxes"`

	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (SandboxesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sandboxes":       reflect.TypeOf(SandboxData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m SandboxesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["sandboxes"] = attrs["sandboxes"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type SandboxesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *SandboxesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *SandboxesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SandboxesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Sandbox",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SandboxesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *SandboxesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config SandboxesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest sandbox.ListSandboxesRequest
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

	response, err := client.Sandbox.ListSandboxesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list sandboxes", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var sandbox SandboxData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &sandbox)...)
		if resp.Diagnostics.HasError() {
			return
		}
		sandbox.ProviderConfigData = config.ProviderConfigData

		results = append(results, sandbox.ToObjectValue(ctx))
	}

	config.Sandbox = types.ListValueMust(SandboxData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
