package serving

import (
	"context"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/serving_tf"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DataSourceServingEndpoints() datasource.DataSource {
	return &ServingEndpointsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &ServingEndpointsDataSource{}

type ServingEndpointsDataSource struct {
	Client *common.DatabricksClient
}

type ServingEndpointsData struct {
	Endpoints []serving_tf.ServingEndpoint `tfsdk:"endpoints" tf:"optional,computed"`
}

func (d *ServingEndpointsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "databricks_serving_endpoints"
}

func (d *ServingEndpointsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ServingEndpointsData{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *ServingEndpointsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *ServingEndpointsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var endpoints ServingEndpointsData
	diags = req.Config.Get(ctx, &endpoints)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpointsInfoSdk, err := w.ServingEndpoints.ListAll(ctx)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
		}
		resp.Diagnostics.AddError("failed to list endpoints", err.Error())
		return
	}
	for _, endpoint := range endpointsInfoSdk {
		var endpointsInfo serving_tf.ServingEndpoint
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, endpoint, &endpointsInfo)...)
		if resp.Diagnostics.HasError() {
			return
		}
		endpoints.Endpoints = append(endpoints.Endpoints, endpointsInfo)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, endpoints)...)
}
