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

// Endpoint represents a cloud networking resource in a user's cloud account and
// binds it to the Databricks account.
type Endpoint struct {
	// The Databricks Account in which the endpoint object exists.
	AccountId types.String `tfsdk:"account_id"`
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
	if !from.AzurePrivateEndpointInfo.IsNull() && !from.AzurePrivateEndpointInfo.IsUnknown() {
		if toAzurePrivateEndpointInfo, ok := to.GetAzurePrivateEndpointInfo(ctx); ok {
			if fromAzurePrivateEndpointInfo, ok := from.GetAzurePrivateEndpointInfo(ctx); ok {
				// Recursively sync the fields of AzurePrivateEndpointInfo
				toAzurePrivateEndpointInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAzurePrivateEndpointInfo)
				to.SetAzurePrivateEndpointInfo(ctx, toAzurePrivateEndpointInfo)
			}
		}
	}
}

func (to *Endpoint) SyncFieldsDuringRead(ctx context.Context, from Endpoint) {
	if !from.AzurePrivateEndpointInfo.IsNull() && !from.AzurePrivateEndpointInfo.IsUnknown() {
		if toAzurePrivateEndpointInfo, ok := to.GetAzurePrivateEndpointInfo(ctx); ok {
			if fromAzurePrivateEndpointInfo, ok := from.GetAzurePrivateEndpointInfo(ctx); ok {
				toAzurePrivateEndpointInfo.SyncFieldsDuringRead(ctx, fromAzurePrivateEndpointInfo)
				to.SetAzurePrivateEndpointInfo(ctx, toAzurePrivateEndpointInfo)
			}
		}
	}
}

func (m Endpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_id"] = attrs["account_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["azure_private_endpoint_info"] = attrs["azure_private_endpoint_info"].SetOptional()
	attrs["azure_private_endpoint_info"] = attrs["azure_private_endpoint_info"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["display_name"] = attrs["display_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["endpoint_id"] = attrs["endpoint_id"].SetComputed()
	attrs["endpoint_id"] = attrs["endpoint_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
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
		"azure_private_endpoint_info": reflect.TypeOf(AzurePrivateEndpointInfo{}),
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
			"azure_private_endpoint_info": m.AzurePrivateEndpointInfo,
			"create_time":                 m.CreateTime,
			"display_name":                m.DisplayName,
			"endpoint_id":                 m.EndpointId,
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
			"azure_private_endpoint_info": AzurePrivateEndpointInfo{}.Type(ctx),
			"create_time":                 timetypes.RFC3339{}.Type(ctx),
			"display_name":                types.StringType,
			"endpoint_id":                 types.StringType,
			"name":                        types.StringType,
			"region":                      types.StringType,
			"state":                       types.StringType,
			"use_case":                    types.StringType,
		},
	}
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

type ListEndpointsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`

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
}

func (to *ListEndpointsResponse) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
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
