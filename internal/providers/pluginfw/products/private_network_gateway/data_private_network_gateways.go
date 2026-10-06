// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package private_network_gateway

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

const dataSourcesName = "private_network_gateways"

var _ datasource.DataSourceWithConfigure = &PrivateNetworkGatewaysDataSource{}

func DataSourcePrivateNetworkGateways() datasource.DataSource {
	return &PrivateNetworkGatewaysDataSource{}
}

// PrivateNetworkGatewaysData extends the main model with additional fields.
type PrivateNetworkGatewaysData struct {
	PrivateNetworkGateways types.List `tfsdk:"private_network_gateways"`
	// The network connectivity configuration containing the gateways.
	Parent types.String `tfsdk:"parent"`
}

func (PrivateNetworkGatewaysData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"private_network_gateways": reflect.TypeOf(PrivateNetworkGatewayData{}),
	}
}

func (m PrivateNetworkGatewaysData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()

	attrs["private_network_gateways"] = attrs["private_network_gateways"].SetComputed()
	return attrs
}

type PrivateNetworkGatewaysDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *PrivateNetworkGatewaysDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *PrivateNetworkGatewaysDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, PrivateNetworkGatewaysData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks PrivateNetworkGateway",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *PrivateNetworkGatewaysDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *PrivateNetworkGatewaysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config PrivateNetworkGatewaysData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest networking.ListPrivateNetworkGatewaysRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.PrivateNetworkGateways.ListPrivateNetworkGatewaysAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list private_network_gateways", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var private_network_gateway PrivateNetworkGatewayData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &private_network_gateway)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, private_network_gateway.ToObjectValue(ctx))
	}

	config.PrivateNetworkGateways = types.ListValueMust(PrivateNetworkGatewayData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
