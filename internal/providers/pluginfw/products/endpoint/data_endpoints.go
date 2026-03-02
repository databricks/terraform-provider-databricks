// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package endpoint

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/networking"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "endpoints"

var _ datasource.DataSourceWithConfigure = &EndpointsDataSource{}

func DataSourceEndpoints() datasource.DataSource {
	return &EndpointsDataSource{}
}

// EndpointsData extends the main model with additional fields.
type EndpointsData struct {
	Endpoints types.List `tfsdk:"items"`

	PageSize types.Int64 `tfsdk:"page_size"`

	Parent types.String `tfsdk:"parent"`
}

func (EndpointsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(EndpointData{}),
	}
}

func (m EndpointsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["items"] = attrs["items"].SetComputed()
	return attrs
}

type EndpointsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *EndpointsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *EndpointsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, EndpointsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Endpoint",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *EndpointsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *EndpointsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config EndpointsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest networking.ListEndpointsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Endpoints.ListEndpointsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list endpoints", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var endpoint EndpointData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &endpoint)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, endpoint.ToObjectValue(ctx))
	}

	config.Endpoints = types.ListValueMust(EndpointData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
