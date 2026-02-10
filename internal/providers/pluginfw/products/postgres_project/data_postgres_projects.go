// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_project

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/postgres"
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

const dataSourcesName = "postgres_projects"

var _ datasource.DataSourceWithConfigure = &ProjectsDataSource{}

func DataSourceProjects() datasource.DataSource {
	return &ProjectsDataSource{}
}

// ProjectsData extends the main model with additional fields.
type ProjectsData struct {
	Postgres types.List `tfsdk:"projects"`
	// Upper bound for items returned. Cannot be negative.
	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (ProjectsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"projects":        reflect.TypeOf(ProjectData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m ProjectsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["projects"] = attrs["projects"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type ProjectsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *ProjectsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *ProjectsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ProjectsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Project",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ProjectsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ProjectsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config ProjectsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest postgres.ListProjectsRequest
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

	response, err := client.Postgres.ListProjectsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list postgres_projects", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var project ProjectData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &project)...)
		if resp.Diagnostics.HasError() {
			return
		}
		project.ProviderConfigData = config.ProviderConfigData

		results = append(results, project.ToObjectValue(ctx))
	}

	config.Postgres = types.ListValueMust(ProjectData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
