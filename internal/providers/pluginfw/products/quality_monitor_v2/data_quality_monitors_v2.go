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
	"github.com/databricks/terraform-provider-databricks/internal/service/qualitymonitorv2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "quality_monitors_v2"

var _ datasource.DataSourceWithConfigure = &QualityMonitorsDataSource{}

func DataSourceQualityMonitors() datasource.DataSource {
	return &QualityMonitorsDataSource{}
}

type QualityMonitorsList struct {
	qualitymonitorv2_tf.ListQualityMonitorRequest
	QualityMonitorV2 types.List `tfsdk:"quality_monitors"`
}

func (c QualityMonitorsList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["quality_monitors"] = attrs["quality_monitors"].SetComputed()
	return attrs
}

func (QualityMonitorsList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quality_monitors": reflect.TypeOf(qualitymonitorv2_tf.QualityMonitor{}),
	}
}

type QualityMonitorsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *QualityMonitorsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *QualityMonitorsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, QualityMonitorsList{}, nil)
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

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config QualityMonitorsList
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest qualitymonitorv2.ListQualityMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
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
		var quality_monitor qualitymonitorv2_tf.QualityMonitor
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &quality_monitor)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, quality_monitor.ToObjectValue(ctx))
	}

	var newState QualityMonitorsList
	newState.QualityMonitorV2 = types.ListValueMust(qualitymonitorv2_tf.QualityMonitor{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
