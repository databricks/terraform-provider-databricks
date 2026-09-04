// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace_iam_workspace_assignment_v2

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

const dataSourcesName = "workspace_iam_workspace_assignments_v2"

var _ datasource.DataSourceWithConfigure = &WorkspaceAssignmentsDataSource{}

func DataSourceWorkspaceAssignments() datasource.DataSource {
	return &WorkspaceAssignmentsDataSource{}
}

// WorkspaceAssignmentsData extends the main model with additional fields.
type WorkspaceAssignmentsData struct {
	WorkspaceIamV2 types.List `tfsdk:"workspace_assignments"`
	// The maximum number of workspace assignments to return. The service may
	// return fewer than this value. If not provided, defaults to 1000, which is
	// also the maximum allowed. Requests for more than the maximum are clamped
	// to 1000.
	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (WorkspaceAssignmentsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_assignments": reflect.TypeOf(WorkspaceAssignmentData{}),
		"provider_config":       reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m WorkspaceAssignmentsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["workspace_assignments"] = attrs["workspace_assignments"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type WorkspaceAssignmentsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *WorkspaceAssignmentsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *WorkspaceAssignmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, WorkspaceAssignmentsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks WorkspaceAssignment",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *WorkspaceAssignmentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *WorkspaceAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config WorkspaceAssignmentsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest iamv2.ListWorkspaceAssignmentsProxyRequest
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

	response, err := client.WorkspaceIamV2.ListWorkspaceAssignmentsProxyAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list workspace_iam_workspace_assignments_v2", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var workspace_assignment WorkspaceAssignmentData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &workspace_assignment)...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspace_assignment.ProviderConfigData = config.ProviderConfigData

		results = append(results, workspace_assignment.ToObjectValue(ctx))
	}

	config.WorkspaceIamV2 = types.ListValueMust(WorkspaceAssignmentData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
