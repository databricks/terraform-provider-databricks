// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package private_network_gateway

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/networking"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/networking_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "private_network_gateway"

var _ resource.ResourceWithConfigure = &PrivateNetworkGatewayResource{}

func ResourcePrivateNetworkGateway() resource.Resource {
	return &PrivateNetworkGatewayResource{}
}

type PrivateNetworkGatewayResource struct {
	Client *autogen.DatabricksClient
}

// PrivateNetworkGateway extends the main model with additional fields.
type PrivateNetworkGateway struct {
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
	// The network connectivity configuration that will contain the gateway.
	Parent types.String `tfsdk:"parent"`
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
// PrivateNetworkGateway struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m PrivateNetworkGateway) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGateway
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGateway) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"aws_cloud_connection": m.AwsCloudConnection,
			"azure_cloud_connection":             m.AzureCloudConnection,
			"bandwidth_tier_gigabits_per_second": m.BandwidthTierGigabitsPerSecond,
			"create_time":                        m.CreateTime,
			"destinations":                       m.Destinations,
			"display_name":                       m.DisplayName,
			"error_message":                      m.ErrorMessage,
			"name":                               m.Name,
			"parent":                             m.Parent,
			"private_dns_resolvers":              m.PrivateDnsResolvers,
			"state":                              m.State,
			"traffic_mode":                       m.TrafficMode,
			"update_time":                        m.UpdateTime,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m PrivateNetworkGateway) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"aws_cloud_connection": networking_tf.PrivateNetworkGatewayAwsCloudConnection{}.Type(ctx),
			"azure_cloud_connection":             networking_tf.PrivateNetworkGatewayAzureCloudConnection{}.Type(ctx),
			"bandwidth_tier_gigabits_per_second": types.Int64Type,
			"create_time":                        timetypes.RFC3339{}.Type(ctx),
			"destinations": basetypes.ListType{
				ElemType: networking_tf.PrivateNetworkGatewayDestination{}.Type(ctx),
			},
			"display_name":  types.StringType,
			"error_message": types.StringType,
			"name":          types.StringType,
			"parent":        types.StringType,
			"private_dns_resolvers": basetypes.ListType{
				ElemType: networking_tf.PrivateNetworkGatewayPrivateDnsResolver{}.Type(ctx),
			},
			"state":        types.StringType,
			"traffic_mode": types.StringType,
			"update_time":  timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *PrivateNetworkGateway) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateNetworkGateway) {
	if !from.AwsCloudConnection.IsNull() && !from.AwsCloudConnection.IsUnknown() {
		if toAwsCloudConnection, ok := to.GetAwsCloudConnection(ctx); ok {
			if fromAwsCloudConnection, ok := from.GetAwsCloudConnection(ctx); ok {
				// Recursively sync the fields of AwsCloudConnection
				toAwsCloudConnection.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsCloudConnection)
				to.SetAwsCloudConnection(ctx, toAwsCloudConnection)
			}
		}
	}
	if !from.AzureCloudConnection.IsNull() && !from.AzureCloudConnection.IsUnknown() {
		if toAzureCloudConnection, ok := to.GetAzureCloudConnection(ctx); ok {
			if fromAzureCloudConnection, ok := from.GetAzureCloudConnection(ctx); ok {
				// Recursively sync the fields of AzureCloudConnection
				toAzureCloudConnection.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureCloudConnection)
				to.SetAzureCloudConnection(ctx, toAzureCloudConnection)
			}
		}
	}
	if !from.Destinations.IsNull() && !from.Destinations.IsUnknown() && to.Destinations.IsNull() && len(from.Destinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Destinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Destinations = from.Destinations
	}
	if !from.Destinations.IsNull() && !from.Destinations.IsUnknown() {
		if toDestinations, ok := to.GetDestinations(ctx); ok {
			if fromDestinations, ok := from.GetDestinations(ctx); ok {
				// Recursively sync the fields of each Destinations element by position.
				for i := range toDestinations {
					if i < len(fromDestinations) {
						toDestinations[i].SyncFieldsDuringCreateOrUpdate(ctx, fromDestinations[i])
					}
				}
				to.SetDestinations(ctx, toDestinations)
			}
		}
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.PrivateDnsResolvers.IsNull() && !from.PrivateDnsResolvers.IsUnknown() && to.PrivateDnsResolvers.IsNull() && len(from.PrivateDnsResolvers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrivateDnsResolvers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrivateDnsResolvers = from.PrivateDnsResolvers
	}
	if !from.PrivateDnsResolvers.IsNull() && !from.PrivateDnsResolvers.IsUnknown() {
		if toPrivateDnsResolvers, ok := to.GetPrivateDnsResolvers(ctx); ok {
			if fromPrivateDnsResolvers, ok := from.GetPrivateDnsResolvers(ctx); ok {
				// Recursively sync the fields of each PrivateDnsResolvers element by position.
				for i := range toPrivateDnsResolvers {
					if i < len(fromPrivateDnsResolvers) {
						toPrivateDnsResolvers[i].SyncFieldsDuringCreateOrUpdate(ctx, fromPrivateDnsResolvers[i])
					}
				}
				to.SetPrivateDnsResolvers(ctx, toPrivateDnsResolvers)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *PrivateNetworkGateway) SyncFieldsDuringRead(ctx context.Context, from PrivateNetworkGateway) {
	if !from.AwsCloudConnection.IsNull() && !from.AwsCloudConnection.IsUnknown() {
		if toAwsCloudConnection, ok := to.GetAwsCloudConnection(ctx); ok {
			if fromAwsCloudConnection, ok := from.GetAwsCloudConnection(ctx); ok {
				toAwsCloudConnection.SyncFieldsDuringRead(ctx, fromAwsCloudConnection)
				to.SetAwsCloudConnection(ctx, toAwsCloudConnection)
			}
		}
	}
	if !from.AzureCloudConnection.IsNull() && !from.AzureCloudConnection.IsUnknown() {
		if toAzureCloudConnection, ok := to.GetAzureCloudConnection(ctx); ok {
			if fromAzureCloudConnection, ok := from.GetAzureCloudConnection(ctx); ok {
				toAzureCloudConnection.SyncFieldsDuringRead(ctx, fromAzureCloudConnection)
				to.SetAzureCloudConnection(ctx, toAzureCloudConnection)
			}
		}
	}
	if !from.Destinations.IsNull() && !from.Destinations.IsUnknown() && to.Destinations.IsNull() && len(from.Destinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Destinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Destinations = from.Destinations
	}
	if !from.Destinations.IsNull() && !from.Destinations.IsUnknown() {
		if toDestinations, ok := to.GetDestinations(ctx); ok {
			if fromDestinations, ok := from.GetDestinations(ctx); ok {
				for i := range toDestinations {
					if i < len(fromDestinations) {
						toDestinations[i].SyncFieldsDuringRead(ctx, fromDestinations[i])
					}
				}
				to.SetDestinations(ctx, toDestinations)
			}
		}
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.PrivateDnsResolvers.IsNull() && !from.PrivateDnsResolvers.IsUnknown() && to.PrivateDnsResolvers.IsNull() && len(from.PrivateDnsResolvers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrivateDnsResolvers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrivateDnsResolvers = from.PrivateDnsResolvers
	}
	if !from.PrivateDnsResolvers.IsNull() && !from.PrivateDnsResolvers.IsUnknown() {
		if toPrivateDnsResolvers, ok := to.GetPrivateDnsResolvers(ctx); ok {
			if fromPrivateDnsResolvers, ok := from.GetPrivateDnsResolvers(ctx); ok {
				for i := range toPrivateDnsResolvers {
					if i < len(fromPrivateDnsResolvers) {
						toPrivateDnsResolvers[i].SyncFieldsDuringRead(ctx, fromPrivateDnsResolvers[i])
					}
				}
				to.SetPrivateDnsResolvers(ctx, toPrivateDnsResolvers)
			}
		}
	}
}

func (m PrivateNetworkGateway) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_cloud_connection"] = attrs["aws_cloud_connection"].SetOptional()
	attrs["aws_cloud_connection"] = attrs["aws_cloud_connection"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["azure_cloud_connection"] = attrs["azure_cloud_connection"].SetOptional()
	attrs["azure_cloud_connection"] = attrs["azure_cloud_connection"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["bandwidth_tier_gigabits_per_second"] = attrs["bandwidth_tier_gigabits_per_second"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["destinations"] = attrs["destinations"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["error_message"] = attrs["error_message"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["private_dns_resolvers"] = attrs["private_dns_resolvers"].SetOptional()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["traffic_mode"] = attrs["traffic_mode"].SetRequired()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetAwsCloudConnection returns the value of the AwsCloudConnection field in PrivateNetworkGateway as
// a networking_tf.PrivateNetworkGatewayAwsCloudConnection value.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGateway) GetAwsCloudConnection(ctx context.Context) (networking_tf.PrivateNetworkGatewayAwsCloudConnection, bool) {
	var e networking_tf.PrivateNetworkGatewayAwsCloudConnection
	if m.AwsCloudConnection.IsNull() || m.AwsCloudConnection.IsUnknown() {
		return e, false
	}
	var v networking_tf.PrivateNetworkGatewayAwsCloudConnection
	d := m.AwsCloudConnection.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsCloudConnection sets the value of the AwsCloudConnection field in PrivateNetworkGateway.
func (m *PrivateNetworkGateway) SetAwsCloudConnection(ctx context.Context, v networking_tf.PrivateNetworkGatewayAwsCloudConnection) {
	vs := v.ToObjectValue(ctx)
	m.AwsCloudConnection = vs
}

// GetAzureCloudConnection returns the value of the AzureCloudConnection field in PrivateNetworkGateway as
// a networking_tf.PrivateNetworkGatewayAzureCloudConnection value.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGateway) GetAzureCloudConnection(ctx context.Context) (networking_tf.PrivateNetworkGatewayAzureCloudConnection, bool) {
	var e networking_tf.PrivateNetworkGatewayAzureCloudConnection
	if m.AzureCloudConnection.IsNull() || m.AzureCloudConnection.IsUnknown() {
		return e, false
	}
	var v networking_tf.PrivateNetworkGatewayAzureCloudConnection
	d := m.AzureCloudConnection.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureCloudConnection sets the value of the AzureCloudConnection field in PrivateNetworkGateway.
func (m *PrivateNetworkGateway) SetAzureCloudConnection(ctx context.Context, v networking_tf.PrivateNetworkGatewayAzureCloudConnection) {
	vs := v.ToObjectValue(ctx)
	m.AzureCloudConnection = vs
}

// GetDestinations returns the value of the Destinations field in PrivateNetworkGateway as
// a slice of networking_tf.PrivateNetworkGatewayDestination values.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGateway) GetDestinations(ctx context.Context) ([]networking_tf.PrivateNetworkGatewayDestination, bool) {
	if m.Destinations.IsNull() || m.Destinations.IsUnknown() {
		return nil, false
	}
	var v []networking_tf.PrivateNetworkGatewayDestination
	d := m.Destinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDestinations sets the value of the Destinations field in PrivateNetworkGateway.
func (m *PrivateNetworkGateway) SetDestinations(ctx context.Context, v []networking_tf.PrivateNetworkGatewayDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Destinations = types.ListValueMust(t, vs)
}

// GetPrivateDnsResolvers returns the value of the PrivateDnsResolvers field in PrivateNetworkGateway as
// a slice of networking_tf.PrivateNetworkGatewayPrivateDnsResolver values.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGateway) GetPrivateDnsResolvers(ctx context.Context) ([]networking_tf.PrivateNetworkGatewayPrivateDnsResolver, bool) {
	if m.PrivateDnsResolvers.IsNull() || m.PrivateDnsResolvers.IsUnknown() {
		return nil, false
	}
	var v []networking_tf.PrivateNetworkGatewayPrivateDnsResolver
	d := m.PrivateDnsResolvers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivateDnsResolvers sets the value of the PrivateDnsResolvers field in PrivateNetworkGateway.
func (m *PrivateNetworkGateway) SetPrivateDnsResolvers(ctx context.Context, v []networking_tf.PrivateNetworkGatewayPrivateDnsResolver) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["private_dns_resolvers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrivateDnsResolvers = types.ListValueMust(t, vs)
}

func (r *PrivateNetworkGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *PrivateNetworkGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, PrivateNetworkGateway{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks private_network_gateway",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *PrivateNetworkGatewayResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *PrivateNetworkGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan PrivateNetworkGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var private_network_gateway networking.PrivateNetworkGateway

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &private_network_gateway)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := networking.CreatePrivateNetworkGatewayRequest{
		PrivateNetworkGateway: private_network_gateway,
		Parent:                plan.Parent.ValueString(),
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.PrivateNetworkGateways.CreatePrivateNetworkGateway(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create private_network_gateway", err.Error())
		return
	}

	var newState PrivateNetworkGateway

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		resp.Diagnostics.AddError("error waiting for private_network_gateway to be ready", err.Error())
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *PrivateNetworkGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState PrivateNetworkGateway
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest networking.GetPrivateNetworkGatewayRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
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
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get private_network_gateway", err.Error())
		return
	}

	var newState PrivateNetworkGateway
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *PrivateNetworkGatewayResource) update(ctx context.Context, plan PrivateNetworkGateway, diags *diag.Diagnostics, state *tfsdk.State) {
	var private_network_gateway networking.PrivateNetworkGateway

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &private_network_gateway)...)
	if diags.HasError() {
		return
	}

	updateRequest := networking.UpdatePrivateNetworkGatewayRequest{
		PrivateNetworkGateway: private_network_gateway,
		Name:                  plan.Name.ValueString(),
		UpdateMask:            *fieldmask.New(strings.Split("bandwidth_tier_gigabits_per_second,destinations,display_name,private_dns_resolvers,traffic_mode", ",")),
	}

	client, clientDiags := r.Client.GetAccountClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.PrivateNetworkGateways.UpdatePrivateNetworkGateway(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update private_network_gateway", err.Error())
		return
	}

	var newState PrivateNetworkGateway

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *PrivateNetworkGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan PrivateNetworkGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *PrivateNetworkGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state PrivateNetworkGateway
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest networking.DeletePrivateNetworkGatewayRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.PrivateNetworkGateways.DeletePrivateNetworkGateway(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete private_network_gateway", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &PrivateNetworkGatewayResource{}

func (r *PrivateNetworkGatewayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: name. Got: %q",
				req.ID,
			),
		)
		return
	}

	name := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
}
