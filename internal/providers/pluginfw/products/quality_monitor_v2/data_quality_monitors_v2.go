// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package quality_monitor_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/qualitymonitorv2"
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

const dataSourcesName = "quality_monitors_v2"

var _ datasource.DataSourceWithConfigure = &QualityMonitorsDataSource{}

func DataSourceQualityMonitors() datasource.DataSource {
	return &QualityMonitorsDataSource{}
}

// QualityMonitorsData extends the main model with additional fields.
type QualityMonitorsData struct {
	QualityMonitorV2   types.List   `tfsdk:"quality_monitors"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (QualityMonitorsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitors": reflect.TypeOf(QualityMonitorData{}),
		"provider_config":  reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m QualityMonitorsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["quality_monitors"] = attrs["quality_monitors"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type QualityMonitorsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *QualityMonitorsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *QualityMonitorsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, QualityMonitorsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks QualityMonitor",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *QualityMonitorsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *QualityMonitorsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config QualityMonitorsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest qualitymonitorv2.ListQualityMonitorRequest
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

	response, err := client.QualityMonitorV2.ListQualityMonitorAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list quality_monitors_v2", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var quality_monitor QualityMonitorData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &quality_monitor)...)
		if resp.Diagnostics.HasError() {
			return
		}
		quality_monitor.ProviderConfigData = config.ProviderConfigData

		results = append(results, quality_monitor.ToObjectValue(ctx))
	}

	var newState QualityMonitorsData
	newState.QualityMonitorV2 = types.ListValueMust(QualityMonitorData{}.Type(ctx), results)
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
