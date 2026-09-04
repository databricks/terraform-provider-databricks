// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace_iam_direct_group_member_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/iamv2"
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

const dataSourcesName = "workspace_iam_direct_group_members_v2"

var _ datasource.DataSourceWithConfigure = &DirectGroupMembersDataSource{}

func DataSourceDirectGroupMembers() datasource.DataSource {
	return &DirectGroupMembersDataSource{}
}

// DirectGroupMembersData extends the main model with additional fields.
type DirectGroupMembersData struct {
	WorkspaceIamV2 types.List `tfsdk:"direct_group_members"`
	// The maximum number of members to return. The service may return fewer
	// than this value. If not provided, defaults to 1000, which is also the
	// maximum allowed. Requests for more than the maximum are clamped to 1000.
	PageSize           types.Int64  `tfsdk:"page_size"`
	GroupId            types.Int64  `tfsdk:"group_id"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (DirectGroupMembersData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"direct_group_members": reflect.TypeOf(DirectGroupMemberData{}),
		"provider_config":      reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m DirectGroupMembersData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["direct_group_members"] = attrs["direct_group_members"].SetComputed()
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type DirectGroupMembersDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *DirectGroupMembersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *DirectGroupMembersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DirectGroupMembersData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks DirectGroupMember",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DirectGroupMembersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DirectGroupMembersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config DirectGroupMembersData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest iamv2.ListDirectGroupMembersProxyRequest
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

	response, err := client.WorkspaceIamV2.ListDirectGroupMembersProxyAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list workspace_iam_direct_group_members_v2", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var direct_group_member DirectGroupMemberData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &direct_group_member)...)
		if resp.Diagnostics.HasError() {
			return
		}
		direct_group_member.ProviderConfigData = config.ProviderConfigData

		results = append(results, direct_group_member.ToObjectValue(ctx))
	}

	config.WorkspaceIamV2 = types.ListValueMust(DirectGroupMemberData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
