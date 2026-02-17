// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps_space

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/apps"
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

const dataSourcesName = "apps_spaces"

var _ datasource.DataSourceWithConfigure = &SpacesDataSource{}

func DataSourceSpaces() datasource.DataSource {
	return &SpacesDataSource{}
}

// SpacesData extends the main model with additional fields.
type SpacesData struct {
	Apps types.List `tfsdk:"spaces"`
	// Upper bound for items returned.
	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (SpacesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spaces":          reflect.TypeOf(SpaceData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m SpacesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["spaces"] = attrs["spaces"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type SpacesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *SpacesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *SpacesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SpacesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Space",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SpacesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *SpacesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config SpacesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest apps.ListSpacesRequest
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

	response, err := client.Apps.ListSpacesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list apps_spaces", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var space SpaceData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &space)...)
		if resp.Diagnostics.HasError() {
			return
		}
		space.ProviderConfigData = config.ProviderConfigData

		results = append(results, space.ToObjectValue(ctx))
	}

	config.Apps = types.ListValueMust(SpaceData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
