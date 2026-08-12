// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_role

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

const dataSourcesName = "postgres_roles"

var _ datasource.DataSourceWithConfigure = &RolesDataSource{}

func DataSourceRoles() datasource.DataSource {
	return &RolesDataSource{}
}

// RolesData extends the main model with additional fields.
type RolesData struct {
	Postgres types.List `tfsdk:"roles"`
	// Upper bound for items returned. Cannot be negative.
	PageSize types.Int64 `tfsdk:"page_size"`
	// The Branch that owns this collection of roles. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent             types.String `tfsdk:"parent"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (RolesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"roles":           reflect.TypeOf(RoleData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m RolesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["roles"] = attrs["roles"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type RolesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *RolesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *RolesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, RolesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Role",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *RolesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *RolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config RolesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest postgres.ListRolesRequest
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

	response, err := client.Postgres.ListRolesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list postgres_roles", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var role RoleData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &role)...)
		if resp.Diagnostics.HasError() {
			return
		}
		role.ProviderConfigData = config.ProviderConfigData

		results = append(results, role.ToObjectValue(ctx))
	}

	config.Postgres = types.ListValueMust(RoleData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
