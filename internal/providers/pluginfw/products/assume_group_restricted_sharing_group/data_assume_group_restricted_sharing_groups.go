// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package assume_group_restricted_sharing_group

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/iam"
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

const dataSourcesName = "assume_group_restricted_sharing_groups"

var _ datasource.DataSourceWithConfigure = &RestrictedSharingGroupsDataSource{}

func DataSourceRestrictedSharingGroups() datasource.DataSource {
	return &RestrictedSharingGroupsDataSource{}
}

// RestrictedSharingGroupsData extends the main model with additional fields.
type RestrictedSharingGroupsData struct {
	AssumeGroup types.List `tfsdk:"restricted_sharing_groups"`
	// The maximum number of groups to return. The service may return fewer. If
	// unspecified or 0, at most 100 groups are returned. Values above 100 are
	// coerced to 100. Negative values are rejected.
	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (RestrictedSharingGroupsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"restricted_sharing_groups": reflect.TypeOf(RestrictedSharingGroupData{}),
		"provider_config":           reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m RestrictedSharingGroupsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["restricted_sharing_groups"] = attrs["restricted_sharing_groups"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type RestrictedSharingGroupsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *RestrictedSharingGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *RestrictedSharingGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, RestrictedSharingGroupsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks RestrictedSharingGroup",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *RestrictedSharingGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *RestrictedSharingGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config RestrictedSharingGroupsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest iam.ListRestrictedSharingGroupsRequest
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

	response, err := client.AssumeGroup.ListRestrictedSharingGroupsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list assume_group_restricted_sharing_groups", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var restricted_sharing_group RestrictedSharingGroupData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &restricted_sharing_group)...)
		if resp.Diagnostics.HasError() {
			return
		}
		restricted_sharing_group.ProviderConfigData = config.ProviderConfigData

		results = append(results, restricted_sharing_group.ToObjectValue(ctx))
	}

	config.AssumeGroup = types.ListValueMust(RestrictedSharingGroupData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
