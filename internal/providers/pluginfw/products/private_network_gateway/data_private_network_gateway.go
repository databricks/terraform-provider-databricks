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
	"github.com/databricks/terraform-provider-databricks/internal/service/networking_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "private_network_gateway"

var _ datasource.DataSourceWithConfigure = &PrivateNetworkGatewayDataSource{}

func DataSourcePrivateNetworkGateway() datasource.DataSource {
	return &PrivateNetworkGatewayDataSource{}
}

type PrivateNetworkGatewayDataSource struct {
	Client *autogen.DatabricksClient
}

// PrivateNetworkGatewayData extends the main model with additional fields.
type PrivateNetworkGatewayData struct {
	// The AWS connection used by the gateway.
	AwsCloudConnection types.Object `tfsdk:"aws_cloud_connection"`
	// The Azure connection used by the gateway.
	AzureCloudConnection types.Object `tfsdk:"azure_cloud_connection"`
	// The provisioned bandwidth tier for an Azure gateway, in gigabits per
	// second. Required when creating an Azure gateway.
	BandwidthTierGigabitsPerSecond types.Int64 `tfsdk:"bandwidth_tier_gigabits_per_second"`
	// The time when the gateway was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The destinations routed through this gateway.
	Destinations types.List `tfsdk:"destinations"`
	// The human-readable name of the gateway.
	DisplayName types.String `tfsdk:"display_name"`
	// The failure reason when the gateway is in the FAILED state.
	ErrorMessage types.String `tfsdk:"error_message"`
	// The canonical resource name of the gateway, in the form
	// `accounts/{account_id}/network-connectivity-configs/{ncc_id}/private-network-gateways/{gateway_id}`.
	Name types.String `tfsdk:"name"`
	// The DNS resolvers used for private name resolution.
	PrivateDnsResolvers types.List `tfsdk:"private_dns_resolvers"`
	// The current lifecycle state of the gateway.
	State types.String `tfsdk:"state"`
	// The traffic routed through this gateway.
	TrafficMode types.String `tfsdk:"traffic_mode"`
	// The time when the gateway was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// PrivateNetworkGatewayData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m PrivateNetworkGatewayData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_cloud_connection":   reflect.TypeOf(networking_tf.PrivateNetworkGatewayAwsCloudConnection{}),
		"azure_cloud_connection": reflect.TypeOf(networking_tf.PrivateNetworkGatewayAzureCloudConnection{}),
		"destinations":           reflect.TypeOf(networking_tf.PrivateNetworkGatewayDestination{}),
		"private_dns_resolvers":  reflect.TypeOf(networking_tf.PrivateNetworkGatewayPrivateDnsResolver{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGatewayData
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGatewayData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_cloud_connection":               m.AwsCloudConnection,
			"azure_cloud_connection":             m.AzureCloudConnection,
			"bandwidth_tier_gigabits_per_second": m.BandwidthTierGigabitsPerSecond,
			"create_time":                        m.CreateTime,
			"destinations":                       m.Destinations,
			"display_name":                       m.DisplayName,
			"error_message":                      m.ErrorMessage,
			"name":                               m.Name,
			"private_dns_resolvers":              m.PrivateDnsResolvers,
			"state":                              m.State,
			"traffic_mode":                       m.TrafficMode,
			"update_time":                        m.UpdateTime,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m PrivateNetworkGatewayData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_cloud_connection":               networking_tf.PrivateNetworkGatewayAwsCloudConnection{}.Type(ctx),
			"azure_cloud_connection":             networking_tf.PrivateNetworkGatewayAzureCloudConnection{}.Type(ctx),
			"bandwidth_tier_gigabits_per_second": types.Int64Type,
			"create_time":                        timetypes.RFC3339{}.Type(ctx),
			"destinations": basetypes.ListType{
				ElemType: networking_tf.PrivateNetworkGatewayDestination{}.Type(ctx),
			},
			"display_name":  types.StringType,
			"error_message": types.StringType,
			"name":          types.StringType,
			"private_dns_resolvers": basetypes.ListType{
				ElemType: networking_tf.PrivateNetworkGatewayPrivateDnsResolver{}.Type(ctx),
			},
			"state":        types.StringType,
			"traffic_mode": types.StringType,
			"update_time":  timetypes.RFC3339{}.Type(ctx),
		},
	}
}

func (m PrivateNetworkGatewayData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_cloud_connection"] = attrs["aws_cloud_connection"].SetComputed()
	attrs["azure_cloud_connection"] = attrs["azure_cloud_connection"].SetComputed()
	attrs["bandwidth_tier_gigabits_per_second"] = attrs["bandwidth_tier_gigabits_per_second"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["destinations"] = attrs["destinations"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["error_message"] = attrs["error_message"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["private_dns_resolvers"] = attrs["private_dns_resolvers"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["traffic_mode"] = attrs["traffic_mode"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

func (r *PrivateNetworkGatewayDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *PrivateNetworkGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, PrivateNetworkGatewayData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks PrivateNetworkGateway",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *PrivateNetworkGatewayDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *PrivateNetworkGatewayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config PrivateNetworkGatewayData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest networking.GetPrivateNetworkGatewayRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.PrivateNetworkGateways.GetPrivateNetworkGateway(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get private_network_gateway", err.Error())
		return
	}

	var newState PrivateNetworkGatewayData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
