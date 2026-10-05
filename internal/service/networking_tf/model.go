// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package networking_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AwsVpcEndpointInfo struct {
	// The AWS account ID in which this VPC endpoint lives.
	AwsAccountId types.String `tfsdk:"aws_account_id"`
	// The ID of the Databricks VPC endpoint service that this endpoint connects
	// to.
	AwsEndpointServiceId types.String `tfsdk:"aws_endpoint_service_id"`
	// The ID of the underlying VPC endpoint in AWS. Provided by the customer
	// when registering an existing AWS VPC endpoint.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id"`
}

func (to *AwsVpcEndpointInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsVpcEndpointInfo) {
}

func (to *AwsVpcEndpointInfo) SyncFieldsDuringRead(ctx context.Context, from AwsVpcEndpointInfo) {
}

func (m AwsVpcEndpointInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_account_id"] = attrs["aws_account_id"].SetComputed()
	attrs["aws_account_id"] = attrs["aws_account_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["aws_endpoint_service_id"] = attrs["aws_endpoint_service_id"].SetComputed()
	attrs["aws_endpoint_service_id"] = attrs["aws_endpoint_service_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["aws_vpc_endpoint_id"] = attrs["aws_vpc_endpoint_id"].SetRequired()
	attrs["aws_vpc_endpoint_id"] = attrs["aws_vpc_endpoint_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsVpcEndpointInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AwsVpcEndpointInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsVpcEndpointInfo
// only implements ToObjectValue() and Type().
func (m AwsVpcEndpointInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_account_id":          m.AwsAccountId,
			"aws_endpoint_service_id": m.AwsEndpointServiceId,
			"aws_vpc_endpoint_id":     m.AwsVpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AwsVpcEndpointInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_account_id":          types.StringType,
			"aws_endpoint_service_id": types.StringType,
			"aws_vpc_endpoint_id":     types.StringType,
		},
	}
}

type AzurePrivateEndpointInfo struct {
	// The name of the Private Endpoint in the Azure subscription.
	PrivateEndpointName types.String `tfsdk:"private_endpoint_name"`
	// The GUID of the Private Endpoint resource in the Azure subscription. This
	// is assigned by Azure when the user sets up the Private Endpoint.
	PrivateEndpointResourceGuid types.String `tfsdk:"private_endpoint_resource_guid"`
	// The full resource ID of the Private Endpoint.
	PrivateEndpointResourceId types.String `tfsdk:"private_endpoint_resource_id"`
	// The resource ID of the Databricks Private Link Service that this Private
	// Endpoint connects to.
	PrivateLinkServiceId types.String `tfsdk:"private_link_service_id"`
}

func (to *AzurePrivateEndpointInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AzurePrivateEndpointInfo) {
}

func (to *AzurePrivateEndpointInfo) SyncFieldsDuringRead(ctx context.Context, from AzurePrivateEndpointInfo) {
}

func (m AzurePrivateEndpointInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["private_endpoint_name"] = attrs["private_endpoint_name"].SetRequired()
	attrs["private_endpoint_resource_guid"] = attrs["private_endpoint_resource_guid"].SetRequired()
	attrs["private_endpoint_resource_id"] = attrs["private_endpoint_resource_id"].SetComputed()
	attrs["private_link_service_id"] = attrs["private_link_service_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzurePrivateEndpointInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AzurePrivateEndpointInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzurePrivateEndpointInfo
// only implements ToObjectValue() and Type().
func (m AzurePrivateEndpointInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_endpoint_name":          m.PrivateEndpointName,
			"private_endpoint_resource_guid": m.PrivateEndpointResourceGuid,
			"private_endpoint_resource_id":   m.PrivateEndpointResourceId,
			"private_link_service_id":        m.PrivateLinkServiceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AzurePrivateEndpointInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_endpoint_name":          types.StringType,
			"private_endpoint_resource_guid": types.StringType,
			"private_endpoint_resource_id":   types.StringType,
			"private_link_service_id":        types.StringType,
		},
	}
}

type CreateEndpointRequest struct {
	Endpoint types.Object `tfsdk:"endpoint"`
	// The parent resource name of the account under which the endpoint is
	// created. Format: `accounts/{account_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *CreateEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateEndpointRequest) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				// Recursively sync the fields of Endpoint
				toEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (to *CreateEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from CreateEndpointRequest) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m CreateEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateEndpointRequest
// only implements ToObjectValue() and Type().
func (m CreateEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint": m.Endpoint,
			"parent":   m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint": Endpoint{}.Type(ctx),
			"parent":   types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in CreateEndpointRequest as
// a Endpoint value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateEndpointRequest) GetEndpoint(ctx context.Context) (Endpoint, bool) {
	var e Endpoint
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v Endpoint
	d := m.Endpoint.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoint sets the value of the Endpoint field in CreateEndpointRequest.
func (m *CreateEndpointRequest) SetEndpoint(ctx context.Context, v Endpoint) {
	vs := v.ToObjectValue(ctx)
	m.Endpoint = vs
}

type CreatePrivateNetworkGatewayRequest struct {
	// The network connectivity configuration that will contain the gateway.
	Parent types.String `tfsdk:"-"`
	// The gateway to create.
	PrivateNetworkGateway types.Object `tfsdk:"private_network_gateway"`
	// A unique identifier for this request. The request is idempotent when this
	// is provided.
	RequestId types.String `tfsdk:"-"`
}

func (to *CreatePrivateNetworkGatewayRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePrivateNetworkGatewayRequest) {
	if !from.PrivateNetworkGateway.IsNull() && !from.PrivateNetworkGateway.IsUnknown() {
		if toPrivateNetworkGateway, ok := to.GetPrivateNetworkGateway(ctx); ok {
			if fromPrivateNetworkGateway, ok := from.GetPrivateNetworkGateway(ctx); ok {
				// Recursively sync the fields of PrivateNetworkGateway
				toPrivateNetworkGateway.SyncFieldsDuringCreateOrUpdate(ctx, fromPrivateNetworkGateway)
				to.SetPrivateNetworkGateway(ctx, toPrivateNetworkGateway)
			}
		}
	}
}

func (to *CreatePrivateNetworkGatewayRequest) SyncFieldsDuringRead(ctx context.Context, from CreatePrivateNetworkGatewayRequest) {
	if !from.PrivateNetworkGateway.IsNull() && !from.PrivateNetworkGateway.IsUnknown() {
		if toPrivateNetworkGateway, ok := to.GetPrivateNetworkGateway(ctx); ok {
			if fromPrivateNetworkGateway, ok := from.GetPrivateNetworkGateway(ctx); ok {
				toPrivateNetworkGateway.SyncFieldsDuringRead(ctx, fromPrivateNetworkGateway)
				to.SetPrivateNetworkGateway(ctx, toPrivateNetworkGateway)
			}
		}
	}
}

func (m CreatePrivateNetworkGatewayRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["private_network_gateway"] = attrs["private_network_gateway"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["request_id"] = attrs["request_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePrivateNetworkGatewayRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreatePrivateNetworkGatewayRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"private_network_gateway": reflect.TypeOf(PrivateNetworkGateway{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateNetworkGatewayRequest
// only implements ToObjectValue() and Type().
func (m CreatePrivateNetworkGatewayRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":                  m.Parent,
			"private_network_gateway": m.PrivateNetworkGateway,
			"request_id":              m.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePrivateNetworkGatewayRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent":                  types.StringType,
			"private_network_gateway": PrivateNetworkGateway{}.Type(ctx),
			"request_id":              types.StringType,
		},
	}
}

// GetPrivateNetworkGateway returns the value of the PrivateNetworkGateway field in CreatePrivateNetworkGatewayRequest as
// a PrivateNetworkGateway value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePrivateNetworkGatewayRequest) GetPrivateNetworkGateway(ctx context.Context) (PrivateNetworkGateway, bool) {
	var e PrivateNetworkGateway
	if m.PrivateNetworkGateway.IsNull() || m.PrivateNetworkGateway.IsUnknown() {
		return e, false
	}
	var v PrivateNetworkGateway
	d := m.PrivateNetworkGateway.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivateNetworkGateway sets the value of the PrivateNetworkGateway field in CreatePrivateNetworkGatewayRequest.
func (m *CreatePrivateNetworkGatewayRequest) SetPrivateNetworkGateway(ctx context.Context, v PrivateNetworkGateway) {
	vs := v.ToObjectValue(ctx)
	m.PrivateNetworkGateway = vs
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	Details types.List `tfsdk:"details"`

	ErrorCode types.String `tfsdk:"error_code"`

	Message types.String `tfsdk:"message"`

	StackTrace types.String `tfsdk:"stack_trace"`
}

func (to *DatabricksServiceExceptionWithDetailsProto) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (to *DatabricksServiceExceptionWithDetailsProto) SyncFieldsDuringRead(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (m DatabricksServiceExceptionWithDetailsProto) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["details"] = attrs["details"].SetOptional()
	attrs["error_code"] = attrs["error_code"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["stack_trace"] = attrs["stack_trace"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksServiceExceptionWithDetailsProto.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabricksServiceExceptionWithDetailsProto) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"details": reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksServiceExceptionWithDetailsProto
// only implements ToObjectValue() and Type().
func (m DatabricksServiceExceptionWithDetailsProto) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"details":     m.Details,
			"error_code":  m.ErrorCode,
			"message":     m.Message,
			"stack_trace": m.StackTrace,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabricksServiceExceptionWithDetailsProto) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"details": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"error_code":  types.StringType,
			"message":     types.StringType,
			"stack_trace": types.StringType,
		},
	}
}

// GetDetails returns the value of the Details field in DatabricksServiceExceptionWithDetailsProto as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabricksServiceExceptionWithDetailsProto) GetDetails(ctx context.Context) ([]types.Object, bool) {
	if m.Details.IsNull() || m.Details.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Details.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDetails sets the value of the Details field in DatabricksServiceExceptionWithDetailsProto.
func (m *DatabricksServiceExceptionWithDetailsProto) SetDetails(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["details"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Details = types.ListValueMust(t, vs)
}

type DeleteEndpointRequest struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEndpointRequest) {
}

func (to *DeleteEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteEndpointRequest) {
}

func (m DeleteEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointRequest
// only implements ToObjectValue() and Type().
func (m DeleteEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeletePrivateNetworkGatewayRequest struct {
	// The canonical resource name of the gateway.
	Name types.String `tfsdk:"-"`
}

func (to *DeletePrivateNetworkGatewayRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePrivateNetworkGatewayRequest) {
}

func (to *DeletePrivateNetworkGatewayRequest) SyncFieldsDuringRead(ctx context.Context, from DeletePrivateNetworkGatewayRequest) {
}

func (m DeletePrivateNetworkGatewayRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePrivateNetworkGatewayRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePrivateNetworkGatewayRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePrivateNetworkGatewayRequest
// only implements ToObjectValue() and Type().
func (m DeletePrivateNetworkGatewayRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePrivateNetworkGatewayRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Endpoint represents a cloud networking resource in a user's cloud account and
// binds it to the Databricks account.
type Endpoint struct {
	// The Databricks Account in which the endpoint object exists.
	AccountId types.String `tfsdk:"account_id"`
	// Info for an AWS VPC endpoint.
	AwsVpcEndpointInfo types.Object `tfsdk:"aws_vpc_endpoint_info"`
	// Info for an Azure private endpoint.
	AzurePrivateEndpointInfo types.Object `tfsdk:"azure_private_endpoint_info"`
	// The timestamp when the endpoint was created. The timestamp is in RFC 3339
	// format in UTC timezone.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The human-readable display name of this endpoint. The input should
	// conform to RFC-1034, which restricts to letters, numbers, and hyphens,
	// with the first character a letter, the last a letter or a number, and a
	// 63 character maximum.
	DisplayName types.String `tfsdk:"display_name"`
	// The unique identifier for this endpoint under the account. This field is
	// a UUID generated by Databricks.
	EndpointId types.String `tfsdk:"endpoint_id"`
	// Info for a GCP Private Service Connect endpoint.
	GcpPscEndpointInfo types.Object `tfsdk:"gcp_psc_endpoint_info"`
	// The resource name of the endpoint, which uniquely identifies the
	// endpoint.
	Name types.String `tfsdk:"name"`
	// The cloud provider region where this endpoint is located.
	Region types.String `tfsdk:"region"`
	// The state of the endpoint. The endpoint can only be used if the state is
	// `APPROVED`.
	State types.String `tfsdk:"state"`
	// The use case that determines the type of network connectivity this
	// endpoint provides. This field is automatically determined based on the
	// endpoint configuration and cloud-specific settings.
	UseCase types.String `tfsdk:"use_case"`
}

func (to *Endpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Endpoint) {
	if !from.AwsVpcEndpointInfo.IsNull() && !from.AwsVpcEndpointInfo.IsUnknown() {
		if toAwsVpcEndpointInfo, ok := to.GetAwsVpcEndpointInfo(ctx); ok {
			if fromAwsVpcEndpointInfo, ok := from.GetAwsVpcEndpointInfo(ctx); ok {
				// Recursively sync the fields of AwsVpcEndpointInfo
				toAwsVpcEndpointInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsVpcEndpointInfo)
				to.SetAwsVpcEndpointInfo(ctx, toAwsVpcEndpointInfo)
			}
		}
	}
	if !from.AzurePrivateEndpointInfo.IsNull() && !from.AzurePrivateEndpointInfo.IsUnknown() {
		if toAzurePrivateEndpointInfo, ok := to.GetAzurePrivateEndpointInfo(ctx); ok {
			if fromAzurePrivateEndpointInfo, ok := from.GetAzurePrivateEndpointInfo(ctx); ok {
				// Recursively sync the fields of AzurePrivateEndpointInfo
				toAzurePrivateEndpointInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAzurePrivateEndpointInfo)
				to.SetAzurePrivateEndpointInfo(ctx, toAzurePrivateEndpointInfo)
			}
		}
	}
	if !from.GcpPscEndpointInfo.IsNull() && !from.GcpPscEndpointInfo.IsUnknown() {
		if toGcpPscEndpointInfo, ok := to.GetGcpPscEndpointInfo(ctx); ok {
			if fromGcpPscEndpointInfo, ok := from.GetGcpPscEndpointInfo(ctx); ok {
				// Recursively sync the fields of GcpPscEndpointInfo
				toGcpPscEndpointInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpPscEndpointInfo)
				to.SetGcpPscEndpointInfo(ctx, toGcpPscEndpointInfo)
			}
		}
	}
}

func (to *Endpoint) SyncFieldsDuringRead(ctx context.Context, from Endpoint) {
	if !from.AwsVpcEndpointInfo.IsNull() && !from.AwsVpcEndpointInfo.IsUnknown() {
		if toAwsVpcEndpointInfo, ok := to.GetAwsVpcEndpointInfo(ctx); ok {
			if fromAwsVpcEndpointInfo, ok := from.GetAwsVpcEndpointInfo(ctx); ok {
				toAwsVpcEndpointInfo.SyncFieldsDuringRead(ctx, fromAwsVpcEndpointInfo)
				to.SetAwsVpcEndpointInfo(ctx, toAwsVpcEndpointInfo)
			}
		}
	}
	if !from.AzurePrivateEndpointInfo.IsNull() && !from.AzurePrivateEndpointInfo.IsUnknown() {
		if toAzurePrivateEndpointInfo, ok := to.GetAzurePrivateEndpointInfo(ctx); ok {
			if fromAzurePrivateEndpointInfo, ok := from.GetAzurePrivateEndpointInfo(ctx); ok {
				toAzurePrivateEndpointInfo.SyncFieldsDuringRead(ctx, fromAzurePrivateEndpointInfo)
				to.SetAzurePrivateEndpointInfo(ctx, toAzurePrivateEndpointInfo)
			}
		}
	}
	if !from.GcpPscEndpointInfo.IsNull() && !from.GcpPscEndpointInfo.IsUnknown() {
		if toGcpPscEndpointInfo, ok := to.GetGcpPscEndpointInfo(ctx); ok {
			if fromGcpPscEndpointInfo, ok := from.GetGcpPscEndpointInfo(ctx); ok {
				toGcpPscEndpointInfo.SyncFieldsDuringRead(ctx, fromGcpPscEndpointInfo)
				to.SetGcpPscEndpointInfo(ctx, toGcpPscEndpointInfo)
			}
		}
	}
}

func (m Endpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_id"] = attrs["account_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["aws_vpc_endpoint_info"] = attrs["aws_vpc_endpoint_info"].SetOptional()
	attrs["aws_vpc_endpoint_info"] = attrs["aws_vpc_endpoint_info"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["azure_private_endpoint_info"] = attrs["azure_private_endpoint_info"].SetOptional()
	attrs["azure_private_endpoint_info"] = attrs["azure_private_endpoint_info"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["display_name"] = attrs["display_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["endpoint_id"] = attrs["endpoint_id"].SetComputed()
	attrs["endpoint_id"] = attrs["endpoint_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["gcp_psc_endpoint_info"] = attrs["gcp_psc_endpoint_info"].SetOptional()
	attrs["gcp_psc_endpoint_info"] = attrs["gcp_psc_endpoint_info"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["region"] = attrs["region"].SetRequired()
	attrs["region"] = attrs["region"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["state"] = attrs["state"].SetComputed()
	attrs["use_case"] = attrs["use_case"].SetComputed()
	attrs["use_case"] = attrs["use_case"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Endpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Endpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_vpc_endpoint_info":       reflect.TypeOf(AwsVpcEndpointInfo{}),
		"azure_private_endpoint_info": reflect.TypeOf(AzurePrivateEndpointInfo{}),
		"gcp_psc_endpoint_info":       reflect.TypeOf(GcpPscEndpointInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Endpoint
// only implements ToObjectValue() and Type().
func (m Endpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                  m.AccountId,
			"aws_vpc_endpoint_info":       m.AwsVpcEndpointInfo,
			"azure_private_endpoint_info": m.AzurePrivateEndpointInfo,
			"create_time":                 m.CreateTime,
			"display_name":                m.DisplayName,
			"endpoint_id":                 m.EndpointId,
			"gcp_psc_endpoint_info":       m.GcpPscEndpointInfo,
			"name":                        m.Name,
			"region":                      m.Region,
			"state":                       m.State,
			"use_case":                    m.UseCase,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Endpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":                  types.StringType,
			"aws_vpc_endpoint_info":       AwsVpcEndpointInfo{}.Type(ctx),
			"azure_private_endpoint_info": AzurePrivateEndpointInfo{}.Type(ctx),
			"create_time":                 timetypes.RFC3339{}.Type(ctx),
			"display_name":                types.StringType,
			"endpoint_id":                 types.StringType,
			"gcp_psc_endpoint_info":       GcpPscEndpointInfo{}.Type(ctx),
			"name":                        types.StringType,
			"region":                      types.StringType,
			"state":                       types.StringType,
			"use_case":                    types.StringType,
		},
	}
}

// GetAwsVpcEndpointInfo returns the value of the AwsVpcEndpointInfo field in Endpoint as
// a AwsVpcEndpointInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetAwsVpcEndpointInfo(ctx context.Context) (AwsVpcEndpointInfo, bool) {
	var e AwsVpcEndpointInfo
	if m.AwsVpcEndpointInfo.IsNull() || m.AwsVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v AwsVpcEndpointInfo
	d := m.AwsVpcEndpointInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsVpcEndpointInfo sets the value of the AwsVpcEndpointInfo field in Endpoint.
func (m *Endpoint) SetAwsVpcEndpointInfo(ctx context.Context, v AwsVpcEndpointInfo) {
	vs := v.ToObjectValue(ctx)
	m.AwsVpcEndpointInfo = vs
}

// GetAzurePrivateEndpointInfo returns the value of the AzurePrivateEndpointInfo field in Endpoint as
// a AzurePrivateEndpointInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetAzurePrivateEndpointInfo(ctx context.Context) (AzurePrivateEndpointInfo, bool) {
	var e AzurePrivateEndpointInfo
	if m.AzurePrivateEndpointInfo.IsNull() || m.AzurePrivateEndpointInfo.IsUnknown() {
		return e, false
	}
	var v AzurePrivateEndpointInfo
	d := m.AzurePrivateEndpointInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzurePrivateEndpointInfo sets the value of the AzurePrivateEndpointInfo field in Endpoint.
func (m *Endpoint) SetAzurePrivateEndpointInfo(ctx context.Context, v AzurePrivateEndpointInfo) {
	vs := v.ToObjectValue(ctx)
	m.AzurePrivateEndpointInfo = vs
}

// GetGcpPscEndpointInfo returns the value of the GcpPscEndpointInfo field in Endpoint as
// a GcpPscEndpointInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetGcpPscEndpointInfo(ctx context.Context) (GcpPscEndpointInfo, bool) {
	var e GcpPscEndpointInfo
	if m.GcpPscEndpointInfo.IsNull() || m.GcpPscEndpointInfo.IsUnknown() {
		return e, false
	}
	var v GcpPscEndpointInfo
	d := m.GcpPscEndpointInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpPscEndpointInfo sets the value of the GcpPscEndpointInfo field in Endpoint.
func (m *Endpoint) SetGcpPscEndpointInfo(ctx context.Context, v GcpPscEndpointInfo) {
	vs := v.ToObjectValue(ctx)
	m.GcpPscEndpointInfo = vs
}

type GcpPscEndpointInfo struct {
	// The GCP region of the PSC connection endpoint. Provided by the customer
	// when registering an existing PSC endpoint. GCP supports only same-region
	// PSC, so this must match the workspace region.
	EndpointRegion types.String `tfsdk:"endpoint_region"`
	// The GCP consumer project ID in which this PSC endpoint is created.
	// Provided by the customer when registering an existing PSC endpoint.
	ProjectId types.String `tfsdk:"project_id"`
	// The ID of the underlying Private Service Connect connection in the GCP
	// consumer project, assigned by GCP when the PSC connection is created.
	PscConnectionId types.String `tfsdk:"psc_connection_id"`
	// The name of this PSC connection in the GCP consumer project. Provided by
	// the customer when registering an existing PSC endpoint.
	PscEndpoint types.String `tfsdk:"psc_endpoint"`
	// The ID of the Databricks service attachment this PSC endpoint connects
	// to.
	ServiceAttachmentId types.String `tfsdk:"service_attachment_id"`
}

func (to *GcpPscEndpointInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpPscEndpointInfo) {
}

func (to *GcpPscEndpointInfo) SyncFieldsDuringRead(ctx context.Context, from GcpPscEndpointInfo) {
}

func (m GcpPscEndpointInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint_region"] = attrs["endpoint_region"].SetRequired()
	attrs["endpoint_region"] = attrs["endpoint_region"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["psc_connection_id"] = attrs["psc_connection_id"].SetComputed()
	attrs["psc_connection_id"] = attrs["psc_connection_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["psc_endpoint"] = attrs["psc_endpoint"].SetRequired()
	attrs["psc_endpoint"] = attrs["psc_endpoint"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["service_attachment_id"] = attrs["service_attachment_id"].SetComputed()
	attrs["service_attachment_id"] = attrs["service_attachment_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpPscEndpointInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GcpPscEndpointInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpPscEndpointInfo
// only implements ToObjectValue() and Type().
func (m GcpPscEndpointInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_region":       m.EndpointRegion,
			"project_id":            m.ProjectId,
			"psc_connection_id":     m.PscConnectionId,
			"psc_endpoint":          m.PscEndpoint,
			"service_attachment_id": m.ServiceAttachmentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GcpPscEndpointInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_region":       types.StringType,
			"project_id":            types.StringType,
			"psc_connection_id":     types.StringType,
			"psc_endpoint":          types.StringType,
			"service_attachment_id": types.StringType,
		},
	}
}

type GetEndpointRequest struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEndpointRequest) {
}

func (to *GetEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from GetEndpointRequest) {
}

func (m GetEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEndpointRequest
// only implements ToObjectValue() and Type().
func (m GetEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name types.String `tfsdk:"-"`
}

func (to *GetOperationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOperationRequest) {
}

func (to *GetOperationRequest) SyncFieldsDuringRead(ctx context.Context, from GetOperationRequest) {
}

func (m GetOperationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOperationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetOperationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOperationRequest
// only implements ToObjectValue() and Type().
func (m GetOperationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOperationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetPrivateNetworkGatewayRequest struct {
	// The canonical resource name of the gateway.
	Name types.String `tfsdk:"-"`
}

func (to *GetPrivateNetworkGatewayRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPrivateNetworkGatewayRequest) {
}

func (to *GetPrivateNetworkGatewayRequest) SyncFieldsDuringRead(ctx context.Context, from GetPrivateNetworkGatewayRequest) {
}

func (m GetPrivateNetworkGatewayRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPrivateNetworkGatewayRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetPrivateNetworkGatewayRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPrivateNetworkGatewayRequest
// only implements ToObjectValue() and Type().
func (m GetPrivateNetworkGatewayRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPrivateNetworkGatewayRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListEndpointsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// The parent resource name of the account to list endpoints for. Format:
	// `accounts/{account_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListEndpointsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsRequest) {
}

func (to *ListEndpointsRequest) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsRequest) {
}

func (m ListEndpointsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEndpointsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsRequest
// only implements ToObjectValue() and Type().
func (m ListEndpointsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListEndpointsResponse struct {
	Items types.List `tfsdk:"items"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListEndpointsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
	if !from.Items.IsNull() && !from.Items.IsUnknown() {
		if toItems, ok := to.GetItems(ctx); ok {
			if fromItems, ok := from.GetItems(ctx); ok {
				// Recursively sync the fields of each Items element by position.
				for i := range toItems {
					if i < len(fromItems) {
						toItems[i].SyncFieldsDuringCreateOrUpdate(ctx, fromItems[i])
					}
				}
				to.SetItems(ctx, toItems)
			}
		}
	}
}

func (to *ListEndpointsResponse) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
	if !from.Items.IsNull() && !from.Items.IsUnknown() {
		if toItems, ok := to.GetItems(ctx); ok {
			if fromItems, ok := from.GetItems(ctx); ok {
				for i := range toItems {
					if i < len(fromItems) {
						toItems[i].SyncFieldsDuringRead(ctx, fromItems[i])
					}
				}
				to.SetItems(ctx, toItems)
			}
		}
	}
}

func (m ListEndpointsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["items"] = attrs["items"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEndpointsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(Endpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsResponse
// only implements ToObjectValue() and Type().
func (m ListEndpointsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           m.Items,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: Endpoint{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetItems returns the value of the Items field in ListEndpointsResponse as
// a slice of Endpoint values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListEndpointsResponse) GetItems(ctx context.Context) ([]Endpoint, bool) {
	if m.Items.IsNull() || m.Items.IsUnknown() {
		return nil, false
	}
	var v []Endpoint
	d := m.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListEndpointsResponse.
func (m *ListEndpointsResponse) SetItems(ctx context.Context, v []Endpoint) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Items = types.ListValueMust(t, vs)
}

type ListPrivateNetworkGatewaysRequest struct {
	// An opaque token returned by a previous list request.
	PageToken types.String `tfsdk:"-"`
	// The network connectivity configuration containing the gateways.
	Parent types.String `tfsdk:"-"`
}

func (to *ListPrivateNetworkGatewaysRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPrivateNetworkGatewaysRequest) {
}

func (to *ListPrivateNetworkGatewaysRequest) SyncFieldsDuringRead(ctx context.Context, from ListPrivateNetworkGatewaysRequest) {
}

func (m ListPrivateNetworkGatewaysRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPrivateNetworkGatewaysRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListPrivateNetworkGatewaysRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateNetworkGatewaysRequest
// only implements ToObjectValue() and Type().
func (m ListPrivateNetworkGatewaysRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPrivateNetworkGatewaysRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListPrivateNetworkGatewaysResponse struct {
	// An opaque token for the next page, or empty when there are no more
	// results.
	NextPageToken types.String `tfsdk:"next_page_token"`

	PrivateNetworkGateways types.List `tfsdk:"private_network_gateways"`
}

func (to *ListPrivateNetworkGatewaysResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPrivateNetworkGatewaysResponse) {
	if !from.PrivateNetworkGateways.IsNull() && !from.PrivateNetworkGateways.IsUnknown() && to.PrivateNetworkGateways.IsNull() && len(from.PrivateNetworkGateways.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrivateNetworkGateways, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrivateNetworkGateways = from.PrivateNetworkGateways
	}
	if !from.PrivateNetworkGateways.IsNull() && !from.PrivateNetworkGateways.IsUnknown() {
		if toPrivateNetworkGateways, ok := to.GetPrivateNetworkGateways(ctx); ok {
			if fromPrivateNetworkGateways, ok := from.GetPrivateNetworkGateways(ctx); ok {
				// Recursively sync the fields of each PrivateNetworkGateways element by position.
				for i := range toPrivateNetworkGateways {
					if i < len(fromPrivateNetworkGateways) {
						toPrivateNetworkGateways[i].SyncFieldsDuringCreateOrUpdate(ctx, fromPrivateNetworkGateways[i])
					}
				}
				to.SetPrivateNetworkGateways(ctx, toPrivateNetworkGateways)
			}
		}
	}
}

func (to *ListPrivateNetworkGatewaysResponse) SyncFieldsDuringRead(ctx context.Context, from ListPrivateNetworkGatewaysResponse) {
	if !from.PrivateNetworkGateways.IsNull() && !from.PrivateNetworkGateways.IsUnknown() && to.PrivateNetworkGateways.IsNull() && len(from.PrivateNetworkGateways.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrivateNetworkGateways, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrivateNetworkGateways = from.PrivateNetworkGateways
	}
	if !from.PrivateNetworkGateways.IsNull() && !from.PrivateNetworkGateways.IsUnknown() {
		if toPrivateNetworkGateways, ok := to.GetPrivateNetworkGateways(ctx); ok {
			if fromPrivateNetworkGateways, ok := from.GetPrivateNetworkGateways(ctx); ok {
				for i := range toPrivateNetworkGateways {
					if i < len(fromPrivateNetworkGateways) {
						toPrivateNetworkGateways[i].SyncFieldsDuringRead(ctx, fromPrivateNetworkGateways[i])
					}
				}
				to.SetPrivateNetworkGateways(ctx, toPrivateNetworkGateways)
			}
		}
	}
}

func (m ListPrivateNetworkGatewaysResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["private_network_gateways"] = attrs["private_network_gateways"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPrivateNetworkGatewaysResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListPrivateNetworkGatewaysResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"private_network_gateways": reflect.TypeOf(PrivateNetworkGateway{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateNetworkGatewaysResponse
// only implements ToObjectValue() and Type().
func (m ListPrivateNetworkGatewaysResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":          m.NextPageToken,
			"private_network_gateways": m.PrivateNetworkGateways,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPrivateNetworkGatewaysResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"private_network_gateways": basetypes.ListType{
				ElemType: PrivateNetworkGateway{}.Type(ctx),
			},
		},
	}
}

// GetPrivateNetworkGateways returns the value of the PrivateNetworkGateways field in ListPrivateNetworkGatewaysResponse as
// a slice of PrivateNetworkGateway values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListPrivateNetworkGatewaysResponse) GetPrivateNetworkGateways(ctx context.Context) ([]PrivateNetworkGateway, bool) {
	if m.PrivateNetworkGateways.IsNull() || m.PrivateNetworkGateways.IsUnknown() {
		return nil, false
	}
	var v []PrivateNetworkGateway
	d := m.PrivateNetworkGateways.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivateNetworkGateways sets the value of the PrivateNetworkGateways field in ListPrivateNetworkGatewaysResponse.
func (m *ListPrivateNetworkGatewaysResponse) SetPrivateNetworkGateways(ctx context.Context, v []PrivateNetworkGateway) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["private_network_gateways"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrivateNetworkGateways = types.ListValueMust(t, vs)
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done types.Bool `tfsdk:"done"`
	// The error result of the operation in case of failure or cancellation.
	Error types.Object `tfsdk:"error"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata types.Object `tfsdk:"metadata"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name types.String `tfsdk:"name"`
	// The normal, successful response of the operation.
	Response types.Object `tfsdk:"response"`
}

func (to *Operation) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Operation) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				// Recursively sync the fields of Error
				toError.SyncFieldsDuringCreateOrUpdate(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (to *Operation) SyncFieldsDuringRead(ctx context.Context, from Operation) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (m Operation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["done"] = attrs["done"].SetOptional()
	attrs["error"] = attrs["error"].SetOptional()
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["response"] = attrs["response"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Operation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Operation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(DatabricksServiceExceptionWithDetailsProto{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Operation
// only implements ToObjectValue() and Type().
func (m Operation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"done":     m.Done,
			"error":    m.Error,
			"metadata": m.Metadata,
			"name":     m.Name,
			"response": m.Response,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Operation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"done":     types.BoolType,
			"error":    DatabricksServiceExceptionWithDetailsProto{}.Type(ctx),
			"metadata": types.ObjectType{},
			"name":     types.StringType,
			"response": types.ObjectType{},
		},
	}
}

// GetError returns the value of the Error field in Operation as
// a DatabricksServiceExceptionWithDetailsProto value.
// If the field is unknown or null, the boolean return value is false.
func (m *Operation) GetError(ctx context.Context) (DatabricksServiceExceptionWithDetailsProto, bool) {
	var e DatabricksServiceExceptionWithDetailsProto
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v DatabricksServiceExceptionWithDetailsProto
	d := m.Error.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetError sets the value of the Error field in Operation.
func (m *Operation) SetError(ctx context.Context, v DatabricksServiceExceptionWithDetailsProto) {
	vs := v.ToObjectValue(ctx)
	m.Error = vs
}

// A private network gateway connects serverless compute to destinations in a
// customer-managed VPC or VNet.
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
	// The DNS resolvers used for private name resolution.
	PrivateDnsResolvers types.List `tfsdk:"private_dns_resolvers"`
	// The current lifecycle state of the gateway.
	State types.String `tfsdk:"state"`
	// The traffic routed through this gateway.
	TrafficMode types.String `tfsdk:"traffic_mode"`
	// The time when the gateway was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

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
	attrs["name"] = attrs["name"].SetOptional()
	attrs["private_dns_resolvers"] = attrs["private_dns_resolvers"].SetOptional()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["traffic_mode"] = attrs["traffic_mode"].SetRequired()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateNetworkGateway.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrivateNetworkGateway) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_cloud_connection":   reflect.TypeOf(PrivateNetworkGatewayAwsCloudConnection{}),
		"azure_cloud_connection": reflect.TypeOf(PrivateNetworkGatewayAzureCloudConnection{}),
		"destinations":           reflect.TypeOf(PrivateNetworkGatewayDestination{}),
		"private_dns_resolvers":  reflect.TypeOf(PrivateNetworkGatewayPrivateDnsResolver{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGateway
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGateway) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivateNetworkGateway) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_cloud_connection":               PrivateNetworkGatewayAwsCloudConnection{}.Type(ctx),
			"azure_cloud_connection":             PrivateNetworkGatewayAzureCloudConnection{}.Type(ctx),
			"bandwidth_tier_gigabits_per_second": types.Int64Type,
			"create_time":                        timetypes.RFC3339{}.Type(ctx),
			"destinations": basetypes.ListType{
				ElemType: PrivateNetworkGatewayDestination{}.Type(ctx),
			},
			"display_name":  types.StringType,
			"error_message": types.StringType,
			"name":          types.StringType,
			"private_dns_resolvers": basetypes.ListType{
				ElemType: PrivateNetworkGatewayPrivateDnsResolver{}.Type(ctx),
			},
			"state":        types.StringType,
			"traffic_mode": types.StringType,
			"update_time":  timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetAwsCloudConnection returns the value of the AwsCloudConnection field in PrivateNetworkGateway as
// a PrivateNetworkGatewayAwsCloudConnection value.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGateway) GetAwsCloudConnection(ctx context.Context) (PrivateNetworkGatewayAwsCloudConnection, bool) {
	var e PrivateNetworkGatewayAwsCloudConnection
	if m.AwsCloudConnection.IsNull() || m.AwsCloudConnection.IsUnknown() {
		return e, false
	}
	var v PrivateNetworkGatewayAwsCloudConnection
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
func (m *PrivateNetworkGateway) SetAwsCloudConnection(ctx context.Context, v PrivateNetworkGatewayAwsCloudConnection) {
	vs := v.ToObjectValue(ctx)
	m.AwsCloudConnection = vs
}

// GetAzureCloudConnection returns the value of the AzureCloudConnection field in PrivateNetworkGateway as
// a PrivateNetworkGatewayAzureCloudConnection value.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGateway) GetAzureCloudConnection(ctx context.Context) (PrivateNetworkGatewayAzureCloudConnection, bool) {
	var e PrivateNetworkGatewayAzureCloudConnection
	if m.AzureCloudConnection.IsNull() || m.AzureCloudConnection.IsUnknown() {
		return e, false
	}
	var v PrivateNetworkGatewayAzureCloudConnection
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
func (m *PrivateNetworkGateway) SetAzureCloudConnection(ctx context.Context, v PrivateNetworkGatewayAzureCloudConnection) {
	vs := v.ToObjectValue(ctx)
	m.AzureCloudConnection = vs
}

// GetDestinations returns the value of the Destinations field in PrivateNetworkGateway as
// a slice of PrivateNetworkGatewayDestination values.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGateway) GetDestinations(ctx context.Context) ([]PrivateNetworkGatewayDestination, bool) {
	if m.Destinations.IsNull() || m.Destinations.IsUnknown() {
		return nil, false
	}
	var v []PrivateNetworkGatewayDestination
	d := m.Destinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDestinations sets the value of the Destinations field in PrivateNetworkGateway.
func (m *PrivateNetworkGateway) SetDestinations(ctx context.Context, v []PrivateNetworkGatewayDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Destinations = types.ListValueMust(t, vs)
}

// GetPrivateDnsResolvers returns the value of the PrivateDnsResolvers field in PrivateNetworkGateway as
// a slice of PrivateNetworkGatewayPrivateDnsResolver values.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGateway) GetPrivateDnsResolvers(ctx context.Context) ([]PrivateNetworkGatewayPrivateDnsResolver, bool) {
	if m.PrivateDnsResolvers.IsNull() || m.PrivateDnsResolvers.IsUnknown() {
		return nil, false
	}
	var v []PrivateNetworkGatewayPrivateDnsResolver
	d := m.PrivateDnsResolvers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivateDnsResolvers sets the value of the PrivateDnsResolvers field in PrivateNetworkGateway.
func (m *PrivateNetworkGateway) SetPrivateDnsResolvers(ctx context.Context, v []PrivateNetworkGatewayPrivateDnsResolver) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["private_dns_resolvers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrivateDnsResolvers = types.ListValueMust(t, vs)
}

// AWS connection configuration.
type PrivateNetworkGatewayAwsCloudConnection struct {
	// The IAM role that Databricks assumes to manage gateway resources.
	CrossAccountRole types.Object `tfsdk:"cross_account_role"`
	// The subnets where the gateway establishes connectivity.
	GatewaySubnets types.List `tfsdk:"gateway_subnets"`
	// The security groups attached to the gateway network interface.
	SecurityGroupIds types.List `tfsdk:"security_group_ids"`
}

func (to *PrivateNetworkGatewayAwsCloudConnection) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateNetworkGatewayAwsCloudConnection) {
	if !from.CrossAccountRole.IsNull() && !from.CrossAccountRole.IsUnknown() {
		if toCrossAccountRole, ok := to.GetCrossAccountRole(ctx); ok {
			if fromCrossAccountRole, ok := from.GetCrossAccountRole(ctx); ok {
				// Recursively sync the fields of CrossAccountRole
				toCrossAccountRole.SyncFieldsDuringCreateOrUpdate(ctx, fromCrossAccountRole)
				to.SetCrossAccountRole(ctx, toCrossAccountRole)
			}
		}
	}
	if !from.GatewaySubnets.IsNull() && !from.GatewaySubnets.IsUnknown() {
		if toGatewaySubnets, ok := to.GetGatewaySubnets(ctx); ok {
			if fromGatewaySubnets, ok := from.GetGatewaySubnets(ctx); ok {
				// Recursively sync the fields of each GatewaySubnets element by position.
				for i := range toGatewaySubnets {
					if i < len(fromGatewaySubnets) {
						toGatewaySubnets[i].SyncFieldsDuringCreateOrUpdate(ctx, fromGatewaySubnets[i])
					}
				}
				to.SetGatewaySubnets(ctx, toGatewaySubnets)
			}
		}
	}
}

func (to *PrivateNetworkGatewayAwsCloudConnection) SyncFieldsDuringRead(ctx context.Context, from PrivateNetworkGatewayAwsCloudConnection) {
	if !from.CrossAccountRole.IsNull() && !from.CrossAccountRole.IsUnknown() {
		if toCrossAccountRole, ok := to.GetCrossAccountRole(ctx); ok {
			if fromCrossAccountRole, ok := from.GetCrossAccountRole(ctx); ok {
				toCrossAccountRole.SyncFieldsDuringRead(ctx, fromCrossAccountRole)
				to.SetCrossAccountRole(ctx, toCrossAccountRole)
			}
		}
	}
	if !from.GatewaySubnets.IsNull() && !from.GatewaySubnets.IsUnknown() {
		if toGatewaySubnets, ok := to.GetGatewaySubnets(ctx); ok {
			if fromGatewaySubnets, ok := from.GetGatewaySubnets(ctx); ok {
				for i := range toGatewaySubnets {
					if i < len(fromGatewaySubnets) {
						toGatewaySubnets[i].SyncFieldsDuringRead(ctx, fromGatewaySubnets[i])
					}
				}
				to.SetGatewaySubnets(ctx, toGatewaySubnets)
			}
		}
	}
}

func (m PrivateNetworkGatewayAwsCloudConnection) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cross_account_role"] = attrs["cross_account_role"].SetRequired()
	attrs["gateway_subnets"] = attrs["gateway_subnets"].SetRequired()
	attrs["security_group_ids"] = attrs["security_group_ids"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateNetworkGatewayAwsCloudConnection.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrivateNetworkGatewayAwsCloudConnection) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cross_account_role": reflect.TypeOf(PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole{}),
		"gateway_subnets":    reflect.TypeOf(PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet{}),
		"security_group_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGatewayAwsCloudConnection
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGatewayAwsCloudConnection) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cross_account_role": m.CrossAccountRole,
			"gateway_subnets":    m.GatewaySubnets,
			"security_group_ids": m.SecurityGroupIds,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivateNetworkGatewayAwsCloudConnection) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cross_account_role": PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole{}.Type(ctx),
			"gateway_subnets": basetypes.ListType{
				ElemType: PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet{}.Type(ctx),
			},
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetCrossAccountRole returns the value of the CrossAccountRole field in PrivateNetworkGatewayAwsCloudConnection as
// a PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole value.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGatewayAwsCloudConnection) GetCrossAccountRole(ctx context.Context) (PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole, bool) {
	var e PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole
	if m.CrossAccountRole.IsNull() || m.CrossAccountRole.IsUnknown() {
		return e, false
	}
	var v PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole
	d := m.CrossAccountRole.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCrossAccountRole sets the value of the CrossAccountRole field in PrivateNetworkGatewayAwsCloudConnection.
func (m *PrivateNetworkGatewayAwsCloudConnection) SetCrossAccountRole(ctx context.Context, v PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) {
	vs := v.ToObjectValue(ctx)
	m.CrossAccountRole = vs
}

// GetGatewaySubnets returns the value of the GatewaySubnets field in PrivateNetworkGatewayAwsCloudConnection as
// a slice of PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet values.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGatewayAwsCloudConnection) GetGatewaySubnets(ctx context.Context) ([]PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet, bool) {
	if m.GatewaySubnets.IsNull() || m.GatewaySubnets.IsUnknown() {
		return nil, false
	}
	var v []PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet
	d := m.GatewaySubnets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGatewaySubnets sets the value of the GatewaySubnets field in PrivateNetworkGatewayAwsCloudConnection.
func (m *PrivateNetworkGatewayAwsCloudConnection) SetGatewaySubnets(ctx context.Context, v []PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["gateway_subnets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.GatewaySubnets = types.ListValueMust(t, vs)
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in PrivateNetworkGatewayAwsCloudConnection as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGatewayAwsCloudConnection) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
	if m.SecurityGroupIds.IsNull() || m.SecurityGroupIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SecurityGroupIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in PrivateNetworkGatewayAwsCloudConnection.
func (m *PrivateNetworkGatewayAwsCloudConnection) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SecurityGroupIds = types.ListValueMust(t, vs)
}

// An AWS subnet used by the gateway.
type PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet struct {
	// The AWS subnet ID.
	SubnetId types.String `tfsdk:"subnet_id"`
}

func (to *PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) {
}

func (to *PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) SyncFieldsDuringRead(ctx context.Context, from PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) {
}

func (m PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["subnet_id"] = attrs["subnet_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"subnet_id": m.SubnetId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivateNetworkGatewayAwsCloudConnectionAwsGatewaySubnet) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"subnet_id": types.StringType,
		},
	}
}

// A cross-account IAM role used to manage gateway resources.
type PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole struct {
	// The ARN of the IAM role.
	RoleArn types.String `tfsdk:"role_arn"`
}

func (to *PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) {
}

func (to *PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) SyncFieldsDuringRead(ctx context.Context, from PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) {
}

func (m PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["role_arn"] = attrs["role_arn"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"role_arn": m.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"role_arn": types.StringType,
		},
	}
}

// Azure connection configuration.
type PrivateNetworkGatewayAzureCloudConnection struct {
	// The subnet where the gateway establishes connectivity.
	GatewaySubnet types.Object `tfsdk:"gateway_subnet"`
}

func (to *PrivateNetworkGatewayAzureCloudConnection) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateNetworkGatewayAzureCloudConnection) {
	if !from.GatewaySubnet.IsNull() && !from.GatewaySubnet.IsUnknown() {
		if toGatewaySubnet, ok := to.GetGatewaySubnet(ctx); ok {
			if fromGatewaySubnet, ok := from.GetGatewaySubnet(ctx); ok {
				// Recursively sync the fields of GatewaySubnet
				toGatewaySubnet.SyncFieldsDuringCreateOrUpdate(ctx, fromGatewaySubnet)
				to.SetGatewaySubnet(ctx, toGatewaySubnet)
			}
		}
	}
}

func (to *PrivateNetworkGatewayAzureCloudConnection) SyncFieldsDuringRead(ctx context.Context, from PrivateNetworkGatewayAzureCloudConnection) {
	if !from.GatewaySubnet.IsNull() && !from.GatewaySubnet.IsUnknown() {
		if toGatewaySubnet, ok := to.GetGatewaySubnet(ctx); ok {
			if fromGatewaySubnet, ok := from.GetGatewaySubnet(ctx); ok {
				toGatewaySubnet.SyncFieldsDuringRead(ctx, fromGatewaySubnet)
				to.SetGatewaySubnet(ctx, toGatewaySubnet)
			}
		}
	}
}

func (m PrivateNetworkGatewayAzureCloudConnection) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gateway_subnet"] = attrs["gateway_subnet"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateNetworkGatewayAzureCloudConnection.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrivateNetworkGatewayAzureCloudConnection) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gateway_subnet": reflect.TypeOf(PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGatewayAzureCloudConnection
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGatewayAzureCloudConnection) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gateway_subnet": m.GatewaySubnet,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivateNetworkGatewayAzureCloudConnection) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gateway_subnet": PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet{}.Type(ctx),
		},
	}
}

// GetGatewaySubnet returns the value of the GatewaySubnet field in PrivateNetworkGatewayAzureCloudConnection as
// a PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet value.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateNetworkGatewayAzureCloudConnection) GetGatewaySubnet(ctx context.Context) (PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet, bool) {
	var e PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet
	if m.GatewaySubnet.IsNull() || m.GatewaySubnet.IsUnknown() {
		return e, false
	}
	var v PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet
	d := m.GatewaySubnet.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGatewaySubnet sets the value of the GatewaySubnet field in PrivateNetworkGatewayAzureCloudConnection.
func (m *PrivateNetworkGatewayAzureCloudConnection) SetGatewaySubnet(ctx context.Context, v PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) {
	vs := v.ToObjectValue(ctx)
	m.GatewaySubnet = vs
}

// An Azure subnet used by the gateway.
type PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet struct {
	// The full Azure resource ID of the subnet.
	ResourceId types.String `tfsdk:"resource_id"`
}

func (to *PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) {
}

func (to *PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) SyncFieldsDuringRead(ctx context.Context, from PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) {
}

func (m PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["resource_id"] = attrs["resource_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resource_id": m.ResourceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivateNetworkGatewayAzureCloudConnectionAzureGatewaySubnet) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resource_id": types.StringType,
		},
	}
}

// A destination routed through the gateway.
type PrivateNetworkGatewayDestination struct {
	// The destination type.
	DestinationType types.String `tfsdk:"destination_type"`
	// The destination value.
	Value types.String `tfsdk:"value"`
}

func (to *PrivateNetworkGatewayDestination) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateNetworkGatewayDestination) {
}

func (to *PrivateNetworkGatewayDestination) SyncFieldsDuringRead(ctx context.Context, from PrivateNetworkGatewayDestination) {
}

func (m PrivateNetworkGatewayDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_type"] = attrs["destination_type"].SetRequired()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateNetworkGatewayDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrivateNetworkGatewayDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGatewayDestination
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGatewayDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_type": m.DestinationType,
			"value":            m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivateNetworkGatewayDestination) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_type": types.StringType,
			"value":            types.StringType,
		},
	}
}

type PrivateNetworkGatewayOperationMetadata struct {
	// The operation performed on the gateway.
	OperationType types.String `tfsdk:"operation_type"`
}

func (to *PrivateNetworkGatewayOperationMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateNetworkGatewayOperationMetadata) {
}

func (to *PrivateNetworkGatewayOperationMetadata) SyncFieldsDuringRead(ctx context.Context, from PrivateNetworkGatewayOperationMetadata) {
}

func (m PrivateNetworkGatewayOperationMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["operation_type"] = attrs["operation_type"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateNetworkGatewayOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrivateNetworkGatewayOperationMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGatewayOperationMetadata
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGatewayOperationMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"operation_type": m.OperationType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivateNetworkGatewayOperationMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"operation_type": types.StringType,
		},
	}
}

// A private DNS resolver used by the gateway.
type PrivateNetworkGatewayPrivateDnsResolver struct {
	// The resolver type.
	ResolverType types.String `tfsdk:"resolver_type"`
	// The resolver value.
	Value types.String `tfsdk:"value"`
}

func (to *PrivateNetworkGatewayPrivateDnsResolver) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateNetworkGatewayPrivateDnsResolver) {
}

func (to *PrivateNetworkGatewayPrivateDnsResolver) SyncFieldsDuringRead(ctx context.Context, from PrivateNetworkGatewayPrivateDnsResolver) {
}

func (m PrivateNetworkGatewayPrivateDnsResolver) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["resolver_type"] = attrs["resolver_type"].SetRequired()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateNetworkGatewayPrivateDnsResolver.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrivateNetworkGatewayPrivateDnsResolver) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateNetworkGatewayPrivateDnsResolver
// only implements ToObjectValue() and Type().
func (m PrivateNetworkGatewayPrivateDnsResolver) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resolver_type": m.ResolverType,
			"value":         m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivateNetworkGatewayPrivateDnsResolver) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resolver_type": types.StringType,
			"value":         types.StringType,
		},
	}
}

type UpdatePrivateNetworkGatewayRequest struct {
	// The canonical resource name of the gateway, in the form
	// `accounts/{account_id}/network-connectivity-configs/{ncc_id}/private-network-gateways/{gateway_id}`.
	Name types.String `tfsdk:"-"`
	// The gateway containing the desired mutable field values.
	PrivateNetworkGateway types.Object `tfsdk:"private_network_gateway"`
	// The fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdatePrivateNetworkGatewayRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePrivateNetworkGatewayRequest) {
	if !from.PrivateNetworkGateway.IsNull() && !from.PrivateNetworkGateway.IsUnknown() {
		if toPrivateNetworkGateway, ok := to.GetPrivateNetworkGateway(ctx); ok {
			if fromPrivateNetworkGateway, ok := from.GetPrivateNetworkGateway(ctx); ok {
				// Recursively sync the fields of PrivateNetworkGateway
				toPrivateNetworkGateway.SyncFieldsDuringCreateOrUpdate(ctx, fromPrivateNetworkGateway)
				to.SetPrivateNetworkGateway(ctx, toPrivateNetworkGateway)
			}
		}
	}
}

func (to *UpdatePrivateNetworkGatewayRequest) SyncFieldsDuringRead(ctx context.Context, from UpdatePrivateNetworkGatewayRequest) {
	if !from.PrivateNetworkGateway.IsNull() && !from.PrivateNetworkGateway.IsUnknown() {
		if toPrivateNetworkGateway, ok := to.GetPrivateNetworkGateway(ctx); ok {
			if fromPrivateNetworkGateway, ok := from.GetPrivateNetworkGateway(ctx); ok {
				toPrivateNetworkGateway.SyncFieldsDuringRead(ctx, fromPrivateNetworkGateway)
				to.SetPrivateNetworkGateway(ctx, toPrivateNetworkGateway)
			}
		}
	}
}

func (m UpdatePrivateNetworkGatewayRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["private_network_gateway"] = attrs["private_network_gateway"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePrivateNetworkGatewayRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdatePrivateNetworkGatewayRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"private_network_gateway": reflect.TypeOf(PrivateNetworkGateway{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePrivateNetworkGatewayRequest
// only implements ToObjectValue() and Type().
func (m UpdatePrivateNetworkGatewayRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":                    m.Name,
			"private_network_gateway": m.PrivateNetworkGateway,
			"update_mask":             m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePrivateNetworkGatewayRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":                    types.StringType,
			"private_network_gateway": PrivateNetworkGateway{}.Type(ctx),
			"update_mask":             types.StringType,
		},
	}
}

// GetPrivateNetworkGateway returns the value of the PrivateNetworkGateway field in UpdatePrivateNetworkGatewayRequest as
// a PrivateNetworkGateway value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdatePrivateNetworkGatewayRequest) GetPrivateNetworkGateway(ctx context.Context) (PrivateNetworkGateway, bool) {
	var e PrivateNetworkGateway
	if m.PrivateNetworkGateway.IsNull() || m.PrivateNetworkGateway.IsUnknown() {
		return e, false
	}
	var v PrivateNetworkGateway
	d := m.PrivateNetworkGateway.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivateNetworkGateway sets the value of the PrivateNetworkGateway field in UpdatePrivateNetworkGatewayRequest.
func (m *UpdatePrivateNetworkGatewayRequest) SetPrivateNetworkGateway(ctx context.Context, v PrivateNetworkGateway) {
	vs := v.ToObjectValue(ctx)
	m.PrivateNetworkGateway = vs
}
