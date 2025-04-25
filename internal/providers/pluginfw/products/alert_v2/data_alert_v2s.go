// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package alert_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/sql_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "alert_v2s"

var _ datasource.DataSourceWithConfigure = &AlertV2sDataSource{}

func DataSourceAlertV2s() datasource.DataSource {
	return &AlertV2sDataSource{}
}

type AlertV2sList struct {
	sql_tf.ListAlertsV2Request
	AlertsV2 types.List `tfsdk:"alert_v2s"`
}

func (c AlertV2sList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alert_v2s"] = attrs["alert_v2s"].SetComputed()
	return attrs
}

func (AlertV2sList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"alert_v2s": reflect.TypeOf(sql_tf.AlertV2{}),
	}
}

type AlertV2sDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *AlertV2sDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *AlertV2sDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, AlertV2sList{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks AlertV2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AlertV2sDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *AlertV2sDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config AlertV2sList
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest sql.ListAlertsV2Request
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.AlertsV2.ListAlertsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list alert_v2s", err.Error())
		return
	}

	var alert_v2s = []attr.Value{}
	for _, item := range response {
		var alert_v2 sql_tf.AlertV2
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &alert_v2)...)
		if resp.Diagnostics.HasError() {
			return
		}
		alert_v2s = append(alert_v2s, alert_v2.ToObjectValue(ctx))
	}

	var newState AlertV2sList
	newState.AlertsV2 = types.ListValueMust(sql_tf.AlertV2{}.Type(ctx), alert_v2s)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
