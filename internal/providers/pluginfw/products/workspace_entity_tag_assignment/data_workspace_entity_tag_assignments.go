// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace_entity_tag_assignment

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/tags"
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

const dataSourcesName = "workspace_entity_tag_assignments"

var _ datasource.DataSourceWithConfigure = &TagAssignmentsDataSource{}

func DataSourceTagAssignments() datasource.DataSource {
	return &TagAssignmentsDataSource{}
}

// TagAssignmentsData extends the main model with additional fields.
type TagAssignmentsData struct {
	WorkspaceEntityTagAssignments types.List `tfsdk:"tag_assignments"`
	// Optional. Maximum number of tag assignments to return in a single page
	PageSize           types.Int64  `tfsdk:"page_size"`
	EntityType         types.String `tfsdk:"entity_type"`
	EntityId           types.String `tfsdk:"entity_id"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (TagAssignmentsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignments": reflect.TypeOf(TagAssignmentData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m TagAssignmentsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["tag_assignments"] = attrs["tag_assignments"].SetComputed()
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["entity_id"] = attrs["entity_id"].SetRequired()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type TagAssignmentsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *TagAssignmentsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *TagAssignmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, TagAssignmentsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks TagAssignment",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *TagAssignmentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *TagAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config TagAssignmentsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest tags.ListTagAssignmentsRequest
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

	response, err := client.WorkspaceEntityTagAssignments.ListTagAssignmentsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list workspace_entity_tag_assignments", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var tag_assignment TagAssignmentData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &tag_assignment)...)
		if resp.Diagnostics.HasError() {
			return
		}
		tag_assignment.ProviderConfigData = config.ProviderConfigData

		results = append(results, tag_assignment.ToObjectValue(ctx))
	}

	config.WorkspaceEntityTagAssignments = types.ListValueMust(TagAssignmentData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
