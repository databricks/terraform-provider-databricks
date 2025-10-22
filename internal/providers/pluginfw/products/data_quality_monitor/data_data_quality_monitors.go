// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package data_quality_monitor

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/dataquality"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "data_quality_monitors"

var _ datasource.DataSourceWithConfigure = &MonitorsDataSource{}

func DataSourceMonitors() datasource.DataSource {
	return &MonitorsDataSource{}
}

// MonitorsData extends the main model with additional fields.
type MonitorsData struct {
	DataQuality types.List `tfsdk:"monitors"`

	PageSize types.Int64 `tfsdk:"page_size"`
}

func (MonitorsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"monitors": reflect.TypeOf(MonitorData{}),
	}
}

func (m MonitorsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["monitors"] = attrs["monitors"].SetComputed()
	return attrs
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *MonitorsData) SyncFieldsDuringRead(ctx context.Context, from MonitorsData) {
}

type MonitorsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *MonitorsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *MonitorsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, MonitorsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Monitor",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *MonitorsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *MonitorsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config MonitorsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest dataquality.ListMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DataQuality.ListMonitorAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list data_quality_monitors", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var monitor MonitorData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &monitor)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, monitor.ToObjectValue(ctx))
	}

	var newState MonitorsData
	newState.DataQuality = types.ListValueMust(MonitorData{}.Type(ctx), results)
	newState.SyncFieldsDuringRead(ctx, config)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
