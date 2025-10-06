package dashboards

import (
	"context"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/dashboards_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "dashboards"

func DataSourceDashboards() datasource.DataSource {
	return &DashboardsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &DashboardsDataSource{}

type DashboardsDataSource struct {
	Client *common.DatabricksClient
}

type DashboardsInfo struct {
	DashboardNameContains types.String `tfsdk:"dashboard_name_contains"`
	Dashboards            types.List   `tfsdk:"dashboards"`
	ProviderConfig        types.Object `tfsdk:"provider_config"`
}

func (DashboardsInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_name_contains"] = attrs["dashboard_name_contains"].SetOptional()
	attrs["dashboards"] = attrs["dashboards"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (DashboardsInfo) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dashboards":      reflect.TypeOf(dashboards_tf.Dashboard{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (d *DashboardsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *DashboardsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DashboardsInfo{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *DashboardsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func AppendDiagAndCheckErrors(resp *datasource.ReadResponse, diags diag.Diagnostics) bool {
	resp.Diagnostics.Append(diags...)
	return resp.Diagnostics.HasError()
}

func (d *DashboardsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var dashboardInfo DashboardsInfo
	if AppendDiagAndCheckErrors(resp, req.Config.Get(ctx, &dashboardInfo)) {
		return
	}
	var workspaceID string
	if !dashboardInfo.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfigData
		resp.Diagnostics.Append(dashboardInfo.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}
	w, diags := d.Client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	dashboardName := strings.ToLower(dashboardInfo.DashboardNameContains.ValueString())

	dashboardsGoSdk, err := w.Lakeview.ListAll(ctx, dashboards.ListDashboardsRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to fetch Dashboards", err.Error())
		return
	}

	var dashboardsTfSdk []attr.Value
	for _, dashboard := range dashboardsGoSdk {
		if dashboardName != "" && !strings.Contains(strings.ToLower(dashboard.DisplayName), dashboardName) {
			continue
		}

		var dashboardResponse dashboards_tf.Dashboard
		if AppendDiagAndCheckErrors(resp, converters.GoSdkToTfSdkStruct(ctx, dashboard, &dashboardResponse)) {
			return
		}
		dashboardsTfSdk = append(dashboardsTfSdk, dashboardResponse.ToObjectValue(ctx))
	}

	dashboardInfo.Dashboards = types.ListValueMust(dashboards_tf.Dashboard{}.Type(ctx), dashboardsTfSdk)
	resp.Diagnostics.Append(resp.State.Set(ctx, dashboardInfo)...)

}
