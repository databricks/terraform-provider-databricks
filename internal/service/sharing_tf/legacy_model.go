// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package sharing_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateFederationPolicyRequest_SdkV2 struct {
	// Name of the policy. This is the name of the policy to be created.
	Policy types.List `tfsdk:"policy"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being created.
	RecipientName types.String `tfsdk:"-"`
}

func (to *CreateFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateFederationPolicyRequest_SdkV2) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				// Recursively sync the fields of Policy
				toPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (to *CreateFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateFederationPolicyRequest_SdkV2) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m CreateFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["policy"] = attrs["policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["recipient_name"] = attrs["recipient_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":         m.Policy,
			"recipient_name": m.RecipientName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy": basetypes.ListType{
				ElemType: FederationPolicy_SdkV2{}.Type(ctx),
			},
			"recipient_name": types.StringType,
		},
	}
}

// GetPolicy returns the value of the Policy field in CreateFederationPolicyRequest_SdkV2 as
// a FederationPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateFederationPolicyRequest_SdkV2) GetPolicy(ctx context.Context) (FederationPolicy_SdkV2, bool) {
	var e FederationPolicy_SdkV2
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v []FederationPolicy_SdkV2
	d := m.Policy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPolicy sets the value of the Policy field in CreateFederationPolicyRequest_SdkV2.
func (m *CreateFederationPolicyRequest_SdkV2) SetPolicy(ctx context.Context, v FederationPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policy"]
	m.Policy = types.ListValueMust(t, vs)
}

type CreateProvider_SdkV2 struct {
	AuthenticationType types.String `tfsdk:"authentication_type"`
	// Description about the provider.
	Comment types.String `tfsdk:"comment"`
	// The name of the Provider.
	Name types.String `tfsdk:"name"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str"`
}

func (to *CreateProvider_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProvider_SdkV2) {
}

func (to *CreateProvider_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateProvider_SdkV2) {
}

func (m CreateProvider_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["authentication_type"] = attrs["authentication_type"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["recipient_profile_str"] = attrs["recipient_profile_str"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProvider.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateProvider_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProvider_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateProvider_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authentication_type":   m.AuthenticationType,
			"comment":               m.Comment,
			"name":                  m.Name,
			"recipient_profile_str": m.RecipientProfileStr,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateProvider_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authentication_type":   types.StringType,
			"comment":               types.StringType,
			"name":                  types.StringType,
			"recipient_profile_str": types.StringType,
		},
	}
}

type CreateRecipient_SdkV2 struct {
	AuthenticationType types.String `tfsdk:"authentication_type"`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// IP Access List
	IpAccessList types.List `tfsdk:"ip_access_list"`
	// Name of Recipient.
	Name types.String `tfsdk:"name"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs types.List `tfsdk:"properties_kvpairs"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code"`
}

func (to *CreateRecipient_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRecipient_SdkV2) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				// Recursively sync the fields of IpAccessList
				toIpAccessList.SyncFieldsDuringCreateOrUpdate(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
	if !from.PropertiesKvpairs.IsNull() && !from.PropertiesKvpairs.IsUnknown() {
		if toPropertiesKvpairs, ok := to.GetPropertiesKvpairs(ctx); ok {
			if fromPropertiesKvpairs, ok := from.GetPropertiesKvpairs(ctx); ok {
				// Recursively sync the fields of PropertiesKvpairs
				toPropertiesKvpairs.SyncFieldsDuringCreateOrUpdate(ctx, fromPropertiesKvpairs)
				to.SetPropertiesKvpairs(ctx, toPropertiesKvpairs)
			}
		}
	}
}

func (to *CreateRecipient_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateRecipient_SdkV2) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				toIpAccessList.SyncFieldsDuringRead(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
	if !from.PropertiesKvpairs.IsNull() && !from.PropertiesKvpairs.IsUnknown() {
		if toPropertiesKvpairs, ok := to.GetPropertiesKvpairs(ctx); ok {
			if fromPropertiesKvpairs, ok := from.GetPropertiesKvpairs(ctx); ok {
				toPropertiesKvpairs.SyncFieldsDuringRead(ctx, fromPropertiesKvpairs)
				to.SetPropertiesKvpairs(ctx, toPropertiesKvpairs)
			}
		}
	}
}

func (m CreateRecipient_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["authentication_type"] = attrs["authentication_type"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["data_recipient_global_metastore_id"] = attrs["data_recipient_global_metastore_id"].SetOptional()
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties_kvpairs"] = attrs["properties_kvpairs"].SetOptional()
	attrs["properties_kvpairs"] = attrs["properties_kvpairs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["sharing_code"] = attrs["sharing_code"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRecipient.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateRecipient_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList_SdkV2{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRecipient_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateRecipient_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authentication_type":                m.AuthenticationType,
			"comment":                            m.Comment,
			"data_recipient_global_metastore_id": m.DataRecipientGlobalMetastoreId,
			"expiration_time":                    m.ExpirationTime,
			"ip_access_list":                     m.IpAccessList,
			"name":                               m.Name,
			"owner":                              m.Owner,
			"properties_kvpairs":                 m.PropertiesKvpairs,
			"sharing_code":                       m.SharingCode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRecipient_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authentication_type":                types.StringType,
			"comment":                            types.StringType,
			"data_recipient_global_metastore_id": types.StringType,
			"expiration_time":                    types.Int64Type,
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessList_SdkV2{}.Type(ctx),
			},
			"name":  types.StringType,
			"owner": types.StringType,
			"properties_kvpairs": basetypes.ListType{
				ElemType: SecurablePropertiesKvPairs_SdkV2{}.Type(ctx),
			},
			"sharing_code": types.StringType,
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in CreateRecipient_SdkV2 as
// a IpAccessList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRecipient_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessList_SdkV2, bool) {
	var e IpAccessList_SdkV2
	if m.IpAccessList.IsNull() || m.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessList_SdkV2
	d := m.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in CreateRecipient_SdkV2.
func (m *CreateRecipient_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	m.IpAccessList = types.ListValueMust(t, vs)
}

// GetPropertiesKvpairs returns the value of the PropertiesKvpairs field in CreateRecipient_SdkV2 as
// a SecurablePropertiesKvPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRecipient_SdkV2) GetPropertiesKvpairs(ctx context.Context) (SecurablePropertiesKvPairs_SdkV2, bool) {
	var e SecurablePropertiesKvPairs_SdkV2
	if m.PropertiesKvpairs.IsNull() || m.PropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []SecurablePropertiesKvPairs_SdkV2
	d := m.PropertiesKvpairs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPropertiesKvpairs sets the value of the PropertiesKvpairs field in CreateRecipient_SdkV2.
func (m *CreateRecipient_SdkV2) SetPropertiesKvpairs(ctx context.Context, v SecurablePropertiesKvPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["properties_kvpairs"]
	m.PropertiesKvpairs = types.ListValueMust(t, vs)
}

type CreateShare_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Name of the share.
	Name types.String `tfsdk:"name"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root"`
}

func (to *CreateShare_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateShare_SdkV2) {
}

func (to *CreateShare_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateShare_SdkV2) {
}

func (m CreateShare_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateShare.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateShare_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateShare_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateShare_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":      m.Comment,
			"name":         m.Name,
			"storage_root": m.StorageRoot,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateShare_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":      types.StringType,
			"name":         types.StringType,
			"storage_root": types.StringType,
		},
	}
}

type DeleteFederationPolicyRequest_SdkV2 struct {
	// Name of the policy. This is the name of the policy to be deleted.
	Name types.String `tfsdk:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being deleted.
	RecipientName types.String `tfsdk:"-"`
}

func (to *DeleteFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFederationPolicyRequest_SdkV2) {
}

func (to *DeleteFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteFederationPolicyRequest_SdkV2) {
}

func (m DeleteFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["recipient_name"] = attrs["recipient_name"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":           m.Name,
			"recipient_name": m.RecipientName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":           types.StringType,
			"recipient_name": types.StringType,
		},
	}
}

type DeleteProviderRequest_SdkV2 struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteProviderRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteProviderRequest_SdkV2) {
}

func (to *DeleteProviderRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteProviderRequest_SdkV2) {
}

func (m DeleteProviderRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteProviderRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProviderRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteProviderRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteProviderRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteRecipientRequest_SdkV2 struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteRecipientRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRecipientRequest_SdkV2) {
}

func (to *DeleteRecipientRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRecipientRequest_SdkV2) {
}

func (m DeleteRecipientRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRecipientRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRecipientRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRecipientRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRecipientRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRecipientRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteShareRequest_SdkV2 struct {
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteShareRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteShareRequest_SdkV2) {
}

func (to *DeleteShareRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteShareRequest_SdkV2) {
}

func (m DeleteShareRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteShareRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteShareRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteShareRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteShareRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteShareRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Represents a UC dependency.
type DeltaSharingDependency_SdkV2 struct {
	Function types.List `tfsdk:"function"`

	Table types.List `tfsdk:"table"`
}

func (to *DeltaSharingDependency_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaSharingDependency_SdkV2) {
	if !from.Function.IsNull() && !from.Function.IsUnknown() {
		if toFunction, ok := to.GetFunction(ctx); ok {
			if fromFunction, ok := from.GetFunction(ctx); ok {
				// Recursively sync the fields of Function
				toFunction.SyncFieldsDuringCreateOrUpdate(ctx, fromFunction)
				to.SetFunction(ctx, toFunction)
			}
		}
	}
	if !from.Table.IsNull() && !from.Table.IsUnknown() {
		if toTable, ok := to.GetTable(ctx); ok {
			if fromTable, ok := from.GetTable(ctx); ok {
				// Recursively sync the fields of Table
				toTable.SyncFieldsDuringCreateOrUpdate(ctx, fromTable)
				to.SetTable(ctx, toTable)
			}
		}
	}
}

func (to *DeltaSharingDependency_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeltaSharingDependency_SdkV2) {
	if !from.Function.IsNull() && !from.Function.IsUnknown() {
		if toFunction, ok := to.GetFunction(ctx); ok {
			if fromFunction, ok := from.GetFunction(ctx); ok {
				toFunction.SyncFieldsDuringRead(ctx, fromFunction)
				to.SetFunction(ctx, toFunction)
			}
		}
	}
	if !from.Table.IsNull() && !from.Table.IsUnknown() {
		if toTable, ok := to.GetTable(ctx); ok {
			if fromTable, ok := from.GetTable(ctx); ok {
				toTable.SyncFieldsDuringRead(ctx, fromTable)
				to.SetTable(ctx, toTable)
			}
		}
	}
}

func (m DeltaSharingDependency_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["function"] = attrs["function"].SetOptional()
	attrs["function"] = attrs["function"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["table"] = attrs["table"].SetOptional()
	attrs["table"] = attrs["table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSharingDependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeltaSharingDependency_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"function": reflect.TypeOf(DeltaSharingFunctionDependency_SdkV2{}),
		"table":    reflect.TypeOf(DeltaSharingTableDependency_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSharingDependency_SdkV2
// only implements ToObjectValue() and Type().
func (m DeltaSharingDependency_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function": m.Function,
			"table":    m.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaSharingDependency_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function": basetypes.ListType{
				ElemType: DeltaSharingFunctionDependency_SdkV2{}.Type(ctx),
			},
			"table": basetypes.ListType{
				ElemType: DeltaSharingTableDependency_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFunction returns the value of the Function field in DeltaSharingDependency_SdkV2 as
// a DeltaSharingFunctionDependency_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSharingDependency_SdkV2) GetFunction(ctx context.Context) (DeltaSharingFunctionDependency_SdkV2, bool) {
	var e DeltaSharingFunctionDependency_SdkV2
	if m.Function.IsNull() || m.Function.IsUnknown() {
		return e, false
	}
	var v []DeltaSharingFunctionDependency_SdkV2
	d := m.Function.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFunction sets the value of the Function field in DeltaSharingDependency_SdkV2.
func (m *DeltaSharingDependency_SdkV2) SetFunction(ctx context.Context, v DeltaSharingFunctionDependency_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["function"]
	m.Function = types.ListValueMust(t, vs)
}

// GetTable returns the value of the Table field in DeltaSharingDependency_SdkV2 as
// a DeltaSharingTableDependency_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSharingDependency_SdkV2) GetTable(ctx context.Context) (DeltaSharingTableDependency_SdkV2, bool) {
	var e DeltaSharingTableDependency_SdkV2
	if m.Table.IsNull() || m.Table.IsUnknown() {
		return e, false
	}
	var v []DeltaSharingTableDependency_SdkV2
	d := m.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in DeltaSharingDependency_SdkV2.
func (m *DeltaSharingDependency_SdkV2) SetTable(ctx context.Context, v DeltaSharingTableDependency_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	m.Table = types.ListValueMust(t, vs)
}

// Represents a list of dependencies.
type DeltaSharingDependencyList_SdkV2 struct {
	// An array of Dependency.
	Dependencies types.List `tfsdk:"dependencies"`
}

func (to *DeltaSharingDependencyList_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaSharingDependencyList_SdkV2) {
	if !from.Dependencies.IsNull() && !from.Dependencies.IsUnknown() && to.Dependencies.IsNull() && len(from.Dependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Dependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Dependencies = from.Dependencies
	}
}

func (to *DeltaSharingDependencyList_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeltaSharingDependencyList_SdkV2) {
	if !from.Dependencies.IsNull() && !from.Dependencies.IsUnknown() && to.Dependencies.IsNull() && len(from.Dependencies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Dependencies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Dependencies = from.Dependencies
	}
}

func (m DeltaSharingDependencyList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dependencies"] = attrs["dependencies"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSharingDependencyList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeltaSharingDependencyList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies": reflect.TypeOf(DeltaSharingDependency_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSharingDependencyList_SdkV2
// only implements ToObjectValue() and Type().
func (m DeltaSharingDependencyList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dependencies": m.Dependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaSharingDependencyList_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dependencies": basetypes.ListType{
				ElemType: DeltaSharingDependency_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDependencies returns the value of the Dependencies field in DeltaSharingDependencyList_SdkV2 as
// a slice of DeltaSharingDependency_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSharingDependencyList_SdkV2) GetDependencies(ctx context.Context) ([]DeltaSharingDependency_SdkV2, bool) {
	if m.Dependencies.IsNull() || m.Dependencies.IsUnknown() {
		return nil, false
	}
	var v []DeltaSharingDependency_SdkV2
	d := m.Dependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependencies sets the value of the Dependencies field in DeltaSharingDependencyList_SdkV2.
func (m *DeltaSharingDependencyList_SdkV2) SetDependencies(ctx context.Context, v []DeltaSharingDependency_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Dependencies = types.ListValueMust(t, vs)
}

type DeltaSharingFunction_SdkV2 struct {
	// The aliass of registered model.
	Aliases types.List `tfsdk:"aliases"`
	// The comment of the function.
	Comment types.String `tfsdk:"comment"`
	// The data type of the function.
	DataType types.String `tfsdk:"data_type"`
	// The dependency list of the function.
	DependencyList types.List `tfsdk:"dependency_list"`
	// The full data type of the function.
	FullDataType types.String `tfsdk:"full_data_type"`
	// The id of the function.
	Id types.String `tfsdk:"id"`
	// The function parameter information.
	InputParams types.List `tfsdk:"input_params"`
	// The name of the function.
	Name types.String `tfsdk:"name"`
	// The properties of the function.
	Properties types.String `tfsdk:"properties"`
	// The routine definition of the function.
	RoutineDefinition types.String `tfsdk:"routine_definition"`
	// The name of the schema that the function belongs to.
	Schema types.String `tfsdk:"schema"`
	// The securable kind of the function.
	SecurableKind types.String `tfsdk:"securable_kind"`
	// The name of the share that the function belongs to.
	Share types.String `tfsdk:"share"`
	// The id of the share that the function belongs to.
	ShareId types.String `tfsdk:"share_id"`
	// The storage location of the function.
	StorageLocation types.String `tfsdk:"storage_location"`
	// The tags of the function.
	Tags types.List `tfsdk:"tags"`
}

func (to *DeltaSharingFunction_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaSharingFunction_SdkV2) {
	if !from.Aliases.IsNull() && !from.Aliases.IsUnknown() && to.Aliases.IsNull() && len(from.Aliases.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Aliases, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Aliases = from.Aliases
	}
	if !from.DependencyList.IsNull() && !from.DependencyList.IsUnknown() {
		if toDependencyList, ok := to.GetDependencyList(ctx); ok {
			if fromDependencyList, ok := from.GetDependencyList(ctx); ok {
				// Recursively sync the fields of DependencyList
				toDependencyList.SyncFieldsDuringCreateOrUpdate(ctx, fromDependencyList)
				to.SetDependencyList(ctx, toDependencyList)
			}
		}
	}
	if !from.InputParams.IsNull() && !from.InputParams.IsUnknown() {
		if toInputParams, ok := to.GetInputParams(ctx); ok {
			if fromInputParams, ok := from.GetInputParams(ctx); ok {
				// Recursively sync the fields of InputParams
				toInputParams.SyncFieldsDuringCreateOrUpdate(ctx, fromInputParams)
				to.SetInputParams(ctx, toInputParams)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *DeltaSharingFunction_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeltaSharingFunction_SdkV2) {
	if !from.Aliases.IsNull() && !from.Aliases.IsUnknown() && to.Aliases.IsNull() && len(from.Aliases.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Aliases, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Aliases = from.Aliases
	}
	if !from.DependencyList.IsNull() && !from.DependencyList.IsUnknown() {
		if toDependencyList, ok := to.GetDependencyList(ctx); ok {
			if fromDependencyList, ok := from.GetDependencyList(ctx); ok {
				toDependencyList.SyncFieldsDuringRead(ctx, fromDependencyList)
				to.SetDependencyList(ctx, toDependencyList)
			}
		}
	}
	if !from.InputParams.IsNull() && !from.InputParams.IsUnknown() {
		if toInputParams, ok := to.GetInputParams(ctx); ok {
			if fromInputParams, ok := from.GetInputParams(ctx); ok {
				toInputParams.SyncFieldsDuringRead(ctx, fromInputParams)
				to.SetInputParams(ctx, toInputParams)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m DeltaSharingFunction_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aliases"] = attrs["aliases"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["data_type"] = attrs["data_type"].SetOptional()
	attrs["dependency_list"] = attrs["dependency_list"].SetOptional()
	attrs["dependency_list"] = attrs["dependency_list"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["full_data_type"] = attrs["full_data_type"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["input_params"] = attrs["input_params"].SetOptional()
	attrs["input_params"] = attrs["input_params"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["routine_definition"] = attrs["routine_definition"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["securable_kind"] = attrs["securable_kind"].SetOptional()
	attrs["share"] = attrs["share"].SetOptional()
	attrs["share_id"] = attrs["share_id"].SetOptional()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSharingFunction.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeltaSharingFunction_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aliases":         reflect.TypeOf(RegisteredModelAlias_SdkV2{}),
		"dependency_list": reflect.TypeOf(DeltaSharingDependencyList_SdkV2{}),
		"input_params":    reflect.TypeOf(FunctionParameterInfos_SdkV2{}),
		"tags":            reflect.TypeOf(catalog_tf.TagKeyValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSharingFunction_SdkV2
// only implements ToObjectValue() and Type().
func (m DeltaSharingFunction_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aliases":            m.Aliases,
			"comment":            m.Comment,
			"data_type":          m.DataType,
			"dependency_list":    m.DependencyList,
			"full_data_type":     m.FullDataType,
			"id":                 m.Id,
			"input_params":       m.InputParams,
			"name":               m.Name,
			"properties":         m.Properties,
			"routine_definition": m.RoutineDefinition,
			"schema":             m.Schema,
			"securable_kind":     m.SecurableKind,
			"share":              m.Share,
			"share_id":           m.ShareId,
			"storage_location":   m.StorageLocation,
			"tags":               m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaSharingFunction_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aliases": basetypes.ListType{
				ElemType: RegisteredModelAlias_SdkV2{}.Type(ctx),
			},
			"comment":   types.StringType,
			"data_type": types.StringType,
			"dependency_list": basetypes.ListType{
				ElemType: DeltaSharingDependencyList_SdkV2{}.Type(ctx),
			},
			"full_data_type": types.StringType,
			"id":             types.StringType,
			"input_params": basetypes.ListType{
				ElemType: FunctionParameterInfos_SdkV2{}.Type(ctx),
			},
			"name":               types.StringType,
			"properties":         types.StringType,
			"routine_definition": types.StringType,
			"schema":             types.StringType,
			"securable_kind":     types.StringType,
			"share":              types.StringType,
			"share_id":           types.StringType,
			"storage_location":   types.StringType,
			"tags": basetypes.ListType{
				ElemType: catalog_tf.TagKeyValue_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAliases returns the value of the Aliases field in DeltaSharingFunction_SdkV2 as
// a slice of RegisteredModelAlias_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSharingFunction_SdkV2) GetAliases(ctx context.Context) ([]RegisteredModelAlias_SdkV2, bool) {
	if m.Aliases.IsNull() || m.Aliases.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAlias_SdkV2
	d := m.Aliases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAliases sets the value of the Aliases field in DeltaSharingFunction_SdkV2.
func (m *DeltaSharingFunction_SdkV2) SetAliases(ctx context.Context, v []RegisteredModelAlias_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Aliases = types.ListValueMust(t, vs)
}

// GetDependencyList returns the value of the DependencyList field in DeltaSharingFunction_SdkV2 as
// a DeltaSharingDependencyList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSharingFunction_SdkV2) GetDependencyList(ctx context.Context) (DeltaSharingDependencyList_SdkV2, bool) {
	var e DeltaSharingDependencyList_SdkV2
	if m.DependencyList.IsNull() || m.DependencyList.IsUnknown() {
		return e, false
	}
	var v []DeltaSharingDependencyList_SdkV2
	d := m.DependencyList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDependencyList sets the value of the DependencyList field in DeltaSharingFunction_SdkV2.
func (m *DeltaSharingFunction_SdkV2) SetDependencyList(ctx context.Context, v DeltaSharingDependencyList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dependency_list"]
	m.DependencyList = types.ListValueMust(t, vs)
}

// GetInputParams returns the value of the InputParams field in DeltaSharingFunction_SdkV2 as
// a FunctionParameterInfos_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSharingFunction_SdkV2) GetInputParams(ctx context.Context) (FunctionParameterInfos_SdkV2, bool) {
	var e FunctionParameterInfos_SdkV2
	if m.InputParams.IsNull() || m.InputParams.IsUnknown() {
		return e, false
	}
	var v []FunctionParameterInfos_SdkV2
	d := m.InputParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInputParams sets the value of the InputParams field in DeltaSharingFunction_SdkV2.
func (m *DeltaSharingFunction_SdkV2) SetInputParams(ctx context.Context, v FunctionParameterInfos_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["input_params"]
	m.InputParams = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in DeltaSharingFunction_SdkV2 as
// a slice of catalog_tf.TagKeyValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSharingFunction_SdkV2) GetTags(ctx context.Context) ([]catalog_tf.TagKeyValue_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.TagKeyValue_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in DeltaSharingFunction_SdkV2.
func (m *DeltaSharingFunction_SdkV2) SetTags(ctx context.Context, v []catalog_tf.TagKeyValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// A Function in UC as a dependency.
type DeltaSharingFunctionDependency_SdkV2 struct {
	FunctionName types.String `tfsdk:"function_name"`

	SchemaName types.String `tfsdk:"schema_name"`
}

func (to *DeltaSharingFunctionDependency_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaSharingFunctionDependency_SdkV2) {
}

func (to *DeltaSharingFunctionDependency_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeltaSharingFunctionDependency_SdkV2) {
}

func (m DeltaSharingFunctionDependency_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["function_name"] = attrs["function_name"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSharingFunctionDependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeltaSharingFunctionDependency_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSharingFunctionDependency_SdkV2
// only implements ToObjectValue() and Type().
func (m DeltaSharingFunctionDependency_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_name": m.FunctionName,
			"schema_name":   m.SchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaSharingFunctionDependency_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_name": types.StringType,
			"schema_name":   types.StringType,
		},
	}
}

// A Table in UC as a dependency.
type DeltaSharingTableDependency_SdkV2 struct {
	SchemaName types.String `tfsdk:"schema_name"`

	TableName types.String `tfsdk:"table_name"`
}

func (to *DeltaSharingTableDependency_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaSharingTableDependency_SdkV2) {
}

func (to *DeltaSharingTableDependency_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeltaSharingTableDependency_SdkV2) {
}

func (m DeltaSharingTableDependency_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["table_name"] = attrs["table_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSharingTableDependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeltaSharingTableDependency_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSharingTableDependency_SdkV2
// only implements ToObjectValue() and Type().
func (m DeltaSharingTableDependency_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"schema_name": m.SchemaName,
			"table_name":  m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaSharingTableDependency_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"schema_name": types.StringType,
			"table_name":  types.StringType,
		},
	}
}

type FederationPolicy_SdkV2 struct {
	// Description of the policy. This is a user-provided description.
	Comment types.String `tfsdk:"comment"`
	// System-generated timestamp indicating when the policy was created.
	CreateTime types.String `tfsdk:"create_time"`
	// Unique, immutable system-generated identifier for the federation policy.
	Id types.String `tfsdk:"id"`
	// Name of the federation policy. A recipient can have multiple policies
	// with different names. The name must contain only lowercase alphanumeric
	// characters, numbers, and hyphens.
	Name types.String `tfsdk:"name"`
	// Specifies the policy to use for validating OIDC claims in the federated
	// tokens.
	OidcPolicy types.List `tfsdk:"oidc_policy"`
	// System-generated timestamp indicating when the policy was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (to *FederationPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FederationPolicy_SdkV2) {
	if !from.OidcPolicy.IsNull() && !from.OidcPolicy.IsUnknown() {
		if toOidcPolicy, ok := to.GetOidcPolicy(ctx); ok {
			if fromOidcPolicy, ok := from.GetOidcPolicy(ctx); ok {
				// Recursively sync the fields of OidcPolicy
				toOidcPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromOidcPolicy)
				to.SetOidcPolicy(ctx, toOidcPolicy)
			}
		}
	}
}

func (to *FederationPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FederationPolicy_SdkV2) {
	if !from.OidcPolicy.IsNull() && !from.OidcPolicy.IsUnknown() {
		if toOidcPolicy, ok := to.GetOidcPolicy(ctx); ok {
			if fromOidcPolicy, ok := from.GetOidcPolicy(ctx); ok {
				toOidcPolicy.SyncFieldsDuringRead(ctx, fromOidcPolicy)
				to.SetOidcPolicy(ctx, toOidcPolicy)
			}
		}
	}
}

func (m FederationPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetComputed()
	attrs["id"] = attrs["id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["oidc_policy"] = attrs["oidc_policy"].SetOptional()
	attrs["oidc_policy"] = attrs["oidc_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FederationPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FederationPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"oidc_policy": reflect.TypeOf(OidcFederationPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FederationPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m FederationPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":     m.Comment,
			"create_time": m.CreateTime,
			"id":          m.Id,
			"name":        m.Name,
			"oidc_policy": m.OidcPolicy,
			"update_time": m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FederationPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":     types.StringType,
			"create_time": types.StringType,
			"id":          types.StringType,
			"name":        types.StringType,
			"oidc_policy": basetypes.ListType{
				ElemType: OidcFederationPolicy_SdkV2{}.Type(ctx),
			},
			"update_time": types.StringType,
		},
	}
}

// GetOidcPolicy returns the value of the OidcPolicy field in FederationPolicy_SdkV2 as
// a OidcFederationPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *FederationPolicy_SdkV2) GetOidcPolicy(ctx context.Context) (OidcFederationPolicy_SdkV2, bool) {
	var e OidcFederationPolicy_SdkV2
	if m.OidcPolicy.IsNull() || m.OidcPolicy.IsUnknown() {
		return e, false
	}
	var v []OidcFederationPolicy_SdkV2
	d := m.OidcPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOidcPolicy sets the value of the OidcPolicy field in FederationPolicy_SdkV2.
func (m *FederationPolicy_SdkV2) SetOidcPolicy(ctx context.Context, v OidcFederationPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["oidc_policy"]
	m.OidcPolicy = types.ListValueMust(t, vs)
}

// Represents a parameter of a function. The same message is used for both input
// and output columns.
type FunctionParameterInfo_SdkV2 struct {
	// The comment of the parameter.
	Comment types.String `tfsdk:"comment"`
	// The name of the parameter.
	Name types.String `tfsdk:"name"`
	// The default value of the parameter.
	ParameterDefault types.String `tfsdk:"parameter_default"`
	// The mode of the function parameter.
	ParameterMode types.String `tfsdk:"parameter_mode"`
	// The type of the function parameter.
	ParameterType types.String `tfsdk:"parameter_type"`
	// The position of the parameter.
	Position types.Int64 `tfsdk:"position"`
	// The interval type of the parameter type.
	TypeIntervalType types.String `tfsdk:"type_interval_type"`
	// The type of the parameter in JSON format.
	TypeJson types.String `tfsdk:"type_json"`
	// The type of the parameter in Enum format.
	TypeName types.String `tfsdk:"type_name"`
	// The precision of the parameter type.
	TypePrecision types.Int64 `tfsdk:"type_precision"`
	// The scale of the parameter type.
	TypeScale types.Int64 `tfsdk:"type_scale"`
	// The type of the parameter in text format.
	TypeText types.String `tfsdk:"type_text"`
}

func (to *FunctionParameterInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FunctionParameterInfo_SdkV2) {
}

func (to *FunctionParameterInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FunctionParameterInfo_SdkV2) {
}

func (m FunctionParameterInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parameter_default"] = attrs["parameter_default"].SetOptional()
	attrs["parameter_mode"] = attrs["parameter_mode"].SetOptional()
	attrs["parameter_type"] = attrs["parameter_type"].SetOptional()
	attrs["position"] = attrs["position"].SetOptional()
	attrs["type_interval_type"] = attrs["type_interval_type"].SetOptional()
	attrs["type_json"] = attrs["type_json"].SetOptional()
	attrs["type_name"] = attrs["type_name"].SetOptional()
	attrs["type_precision"] = attrs["type_precision"].SetOptional()
	attrs["type_scale"] = attrs["type_scale"].SetOptional()
	attrs["type_text"] = attrs["type_text"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FunctionParameterInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FunctionParameterInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionParameterInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m FunctionParameterInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":            m.Comment,
			"name":               m.Name,
			"parameter_default":  m.ParameterDefault,
			"parameter_mode":     m.ParameterMode,
			"parameter_type":     m.ParameterType,
			"position":           m.Position,
			"type_interval_type": m.TypeIntervalType,
			"type_json":          m.TypeJson,
			"type_name":          m.TypeName,
			"type_precision":     m.TypePrecision,
			"type_scale":         m.TypeScale,
			"type_text":          m.TypeText,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FunctionParameterInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":            types.StringType,
			"name":               types.StringType,
			"parameter_default":  types.StringType,
			"parameter_mode":     types.StringType,
			"parameter_type":     types.StringType,
			"position":           types.Int64Type,
			"type_interval_type": types.StringType,
			"type_json":          types.StringType,
			"type_name":          types.StringType,
			"type_precision":     types.Int64Type,
			"type_scale":         types.Int64Type,
			"type_text":          types.StringType,
		},
	}
}

type FunctionParameterInfos_SdkV2 struct {
	// The list of parameters of the function.
	Parameters types.List `tfsdk:"parameters"`
}

func (to *FunctionParameterInfos_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FunctionParameterInfos_SdkV2) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
}

func (to *FunctionParameterInfos_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FunctionParameterInfos_SdkV2) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() && to.Parameters.IsNull() && len(from.Parameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parameters = from.Parameters
	}
}

func (m FunctionParameterInfos_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parameters"] = attrs["parameters"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FunctionParameterInfos.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FunctionParameterInfos_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(FunctionParameterInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionParameterInfos_SdkV2
// only implements ToObjectValue() and Type().
func (m FunctionParameterInfos_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parameters": m.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FunctionParameterInfos_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parameters": basetypes.ListType{
				ElemType: FunctionParameterInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetParameters returns the value of the Parameters field in FunctionParameterInfos_SdkV2 as
// a slice of FunctionParameterInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *FunctionParameterInfos_SdkV2) GetParameters(ctx context.Context) ([]FunctionParameterInfo_SdkV2, bool) {
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return nil, false
	}
	var v []FunctionParameterInfo_SdkV2
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in FunctionParameterInfos_SdkV2.
func (m *FunctionParameterInfos_SdkV2) SetParameters(ctx context.Context, v []FunctionParameterInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parameters = types.ListValueMust(t, vs)
}

type GetActivationUrlInfoRequest_SdkV2 struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

func (to *GetActivationUrlInfoRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetActivationUrlInfoRequest_SdkV2) {
}

func (to *GetActivationUrlInfoRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetActivationUrlInfoRequest_SdkV2) {
}

func (m GetActivationUrlInfoRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activation_url"] = attrs["activation_url"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetActivationUrlInfoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetActivationUrlInfoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetActivationUrlInfoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetActivationUrlInfoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activation_url": m.ActivationUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetActivationUrlInfoRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activation_url": types.StringType,
		},
	}
}

type GetActivationUrlInfoResponse_SdkV2 struct {
}

func (to *GetActivationUrlInfoResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetActivationUrlInfoResponse_SdkV2) {
}

func (to *GetActivationUrlInfoResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetActivationUrlInfoResponse_SdkV2) {
}

func (m GetActivationUrlInfoResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetActivationUrlInfoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetActivationUrlInfoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetActivationUrlInfoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetActivationUrlInfoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetActivationUrlInfoResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetFederationPolicyRequest_SdkV2 struct {
	// Name of the policy. This is the name of the policy to be retrieved.
	Name types.String `tfsdk:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being retrieved.
	RecipientName types.String `tfsdk:"-"`
}

func (to *GetFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFederationPolicyRequest_SdkV2) {
}

func (to *GetFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetFederationPolicyRequest_SdkV2) {
}

func (m GetFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["recipient_name"] = attrs["recipient_name"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":           m.Name,
			"recipient_name": m.RecipientName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":           types.StringType,
			"recipient_name": types.StringType,
		},
	}
}

type GetProviderRequest_SdkV2 struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
}

func (to *GetProviderRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProviderRequest_SdkV2) {
}

func (to *GetProviderRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetProviderRequest_SdkV2) {
}

func (m GetProviderRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetProviderRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProviderRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetProviderRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetProviderRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetRecipientRequest_SdkV2 struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (to *GetRecipientRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRecipientRequest_SdkV2) {
}

func (to *GetRecipientRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRecipientRequest_SdkV2) {
}

func (m GetRecipientRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRecipientRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRecipientRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRecipientRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRecipientRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRecipientRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetRecipientSharePermissionsResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of data share permissions for a recipient.
	PermissionsOut types.List `tfsdk:"permissions_out"`
}

func (to *GetRecipientSharePermissionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRecipientSharePermissionsResponse_SdkV2) {
	if !from.PermissionsOut.IsNull() && !from.PermissionsOut.IsUnknown() && to.PermissionsOut.IsNull() && len(from.PermissionsOut.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionsOut, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionsOut = from.PermissionsOut
	}
}

func (to *GetRecipientSharePermissionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRecipientSharePermissionsResponse_SdkV2) {
	if !from.PermissionsOut.IsNull() && !from.PermissionsOut.IsUnknown() && to.PermissionsOut.IsNull() && len(from.PermissionsOut.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionsOut, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionsOut = from.PermissionsOut
	}
}

func (m GetRecipientSharePermissionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["permissions_out"] = attrs["permissions_out"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRecipientSharePermissionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRecipientSharePermissionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions_out": reflect.TypeOf(ShareToPrivilegeAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRecipientSharePermissionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRecipientSharePermissionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"permissions_out": m.PermissionsOut,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRecipientSharePermissionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"permissions_out": basetypes.ListType{
				ElemType: ShareToPrivilegeAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionsOut returns the value of the PermissionsOut field in GetRecipientSharePermissionsResponse_SdkV2 as
// a slice of ShareToPrivilegeAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetRecipientSharePermissionsResponse_SdkV2) GetPermissionsOut(ctx context.Context) ([]ShareToPrivilegeAssignment_SdkV2, bool) {
	if m.PermissionsOut.IsNull() || m.PermissionsOut.IsUnknown() {
		return nil, false
	}
	var v []ShareToPrivilegeAssignment_SdkV2
	d := m.PermissionsOut.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionsOut sets the value of the PermissionsOut field in GetRecipientSharePermissionsResponse_SdkV2.
func (m *GetRecipientSharePermissionsResponse_SdkV2) SetPermissionsOut(ctx context.Context, v []ShareToPrivilegeAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions_out"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionsOut = types.ListValueMust(t, vs)
}

type GetSharePermissionsResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The privileges assigned to each principal
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments"`
}

func (to *GetSharePermissionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSharePermissionsResponse_SdkV2) {
	if !from.PrivilegeAssignments.IsNull() && !from.PrivilegeAssignments.IsUnknown() && to.PrivilegeAssignments.IsNull() && len(from.PrivilegeAssignments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrivilegeAssignments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrivilegeAssignments = from.PrivilegeAssignments
	}
}

func (to *GetSharePermissionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSharePermissionsResponse_SdkV2) {
	if !from.PrivilegeAssignments.IsNull() && !from.PrivilegeAssignments.IsUnknown() && to.PrivilegeAssignments.IsNull() && len(from.PrivilegeAssignments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrivilegeAssignments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrivilegeAssignments = from.PrivilegeAssignments
	}
}

func (m GetSharePermissionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["privilege_assignments"] = attrs["privilege_assignments"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSharePermissionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSharePermissionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(PrivilegeAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSharePermissionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSharePermissionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":       m.NextPageToken,
			"privilege_assignments": m.PrivilegeAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSharePermissionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in GetSharePermissionsResponse_SdkV2 as
// a slice of PrivilegeAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetSharePermissionsResponse_SdkV2) GetPrivilegeAssignments(ctx context.Context) ([]PrivilegeAssignment_SdkV2, bool) {
	if m.PrivilegeAssignments.IsNull() || m.PrivilegeAssignments.IsUnknown() {
		return nil, false
	}
	var v []PrivilegeAssignment_SdkV2
	d := m.PrivilegeAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in GetSharePermissionsResponse_SdkV2.
func (m *GetSharePermissionsResponse_SdkV2) SetPrivilegeAssignments(ctx context.Context, v []PrivilegeAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrivilegeAssignments = types.ListValueMust(t, vs)
}

type GetShareRequest_SdkV2 struct {
	// Query for data to include in the share.
	IncludeSharedData types.Bool `tfsdk:"-"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (to *GetShareRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetShareRequest_SdkV2) {
}

func (to *GetShareRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetShareRequest_SdkV2) {
}

func (m GetShareRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["include_shared_data"] = attrs["include_shared_data"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetShareRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetShareRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetShareRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetShareRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_shared_data": m.IncludeSharedData,
			"name":                m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetShareRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_shared_data": types.BoolType,
			"name":                types.StringType,
		},
	}
}

type IpAccessList_SdkV2 struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses types.List `tfsdk:"allowed_ip_addresses"`
}

func (to *IpAccessList_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IpAccessList_SdkV2) {
	if !from.AllowedIpAddresses.IsNull() && !from.AllowedIpAddresses.IsUnknown() && to.AllowedIpAddresses.IsNull() && len(from.AllowedIpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedIpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedIpAddresses = from.AllowedIpAddresses
	}
}

func (to *IpAccessList_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IpAccessList_SdkV2) {
	if !from.AllowedIpAddresses.IsNull() && !from.AllowedIpAddresses.IsUnknown() && to.AllowedIpAddresses.IsNull() && len(from.AllowedIpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedIpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedIpAddresses = from.AllowedIpAddresses
	}
}

func (m IpAccessList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allowed_ip_addresses"] = attrs["allowed_ip_addresses"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m IpAccessList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IpAccessList_SdkV2
// only implements ToObjectValue() and Type().
func (m IpAccessList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_ip_addresses": m.AllowedIpAddresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IpAccessList_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAllowedIpAddresses returns the value of the AllowedIpAddresses field in IpAccessList_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *IpAccessList_SdkV2) GetAllowedIpAddresses(ctx context.Context) ([]types.String, bool) {
	if m.AllowedIpAddresses.IsNull() || m.AllowedIpAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.AllowedIpAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedIpAddresses sets the value of the AllowedIpAddresses field in IpAccessList_SdkV2.
func (m *IpAccessList_SdkV2) SetAllowedIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedIpAddresses = types.ListValueMust(t, vs)
}

type ListFederationPoliciesRequest_SdkV2 struct {
	MaxResults types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policies are being listed.
	RecipientName types.String `tfsdk:"-"`
}

func (to *ListFederationPoliciesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFederationPoliciesRequest_SdkV2) {
}

func (to *ListFederationPoliciesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFederationPoliciesRequest_SdkV2) {
}

func (m ListFederationPoliciesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["recipient_name"] = attrs["recipient_name"].SetRequired()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFederationPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListFederationPoliciesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFederationPoliciesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListFederationPoliciesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results":    m.MaxResults,
			"page_token":     m.PageToken,
			"recipient_name": m.RecipientName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFederationPoliciesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
			"recipient_name": types.StringType,
		},
	}
}

type ListFederationPoliciesResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Policies types.List `tfsdk:"policies"`
}

func (to *ListFederationPoliciesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFederationPoliciesResponse_SdkV2) {
	if !from.Policies.IsNull() && !from.Policies.IsUnknown() && to.Policies.IsNull() && len(from.Policies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Policies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Policies = from.Policies
	}
}

func (to *ListFederationPoliciesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFederationPoliciesResponse_SdkV2) {
	if !from.Policies.IsNull() && !from.Policies.IsUnknown() && to.Policies.IsNull() && len(from.Policies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Policies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Policies = from.Policies
	}
}

func (m ListFederationPoliciesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["policies"] = attrs["policies"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFederationPoliciesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListFederationPoliciesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(FederationPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFederationPoliciesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListFederationPoliciesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"policies":        m.Policies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFederationPoliciesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"policies": basetypes.ListType{
				ElemType: FederationPolicy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPolicies returns the value of the Policies field in ListFederationPoliciesResponse_SdkV2 as
// a slice of FederationPolicy_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListFederationPoliciesResponse_SdkV2) GetPolicies(ctx context.Context) ([]FederationPolicy_SdkV2, bool) {
	if m.Policies.IsNull() || m.Policies.IsUnknown() {
		return nil, false
	}
	var v []FederationPolicy_SdkV2
	d := m.Policies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicies sets the value of the Policies field in ListFederationPoliciesResponse_SdkV2.
func (m *ListFederationPoliciesResponse_SdkV2) SetPolicies(ctx context.Context, v []FederationPolicy_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Policies = types.ListValueMust(t, vs)
}

type ListProviderShareAssetsRequest_SdkV2 struct {
	// Maximum number of functions to return.
	FunctionMaxResults types.Int64 `tfsdk:"-"`
	// Maximum number of notebooks to return.
	NotebookMaxResults types.Int64 `tfsdk:"-"`
	// The name of the provider who owns the share.
	ProviderName types.String `tfsdk:"-"`
	// The name of the share.
	ShareName types.String `tfsdk:"-"`
	// Maximum number of tables to return.
	TableMaxResults types.Int64 `tfsdk:"-"`
	// Maximum number of volumes to return.
	VolumeMaxResults types.Int64 `tfsdk:"-"`
}

func (to *ListProviderShareAssetsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProviderShareAssetsRequest_SdkV2) {
}

func (to *ListProviderShareAssetsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProviderShareAssetsRequest_SdkV2) {
}

func (m ListProviderShareAssetsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["provider_name"] = attrs["provider_name"].SetRequired()
	attrs["share_name"] = attrs["share_name"].SetRequired()
	attrs["table_max_results"] = attrs["table_max_results"].SetOptional()
	attrs["function_max_results"] = attrs["function_max_results"].SetOptional()
	attrs["volume_max_results"] = attrs["volume_max_results"].SetOptional()
	attrs["notebook_max_results"] = attrs["notebook_max_results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProviderShareAssetsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProviderShareAssetsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderShareAssetsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListProviderShareAssetsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_max_results": m.FunctionMaxResults,
			"notebook_max_results": m.NotebookMaxResults,
			"provider_name":        m.ProviderName,
			"share_name":           m.ShareName,
			"table_max_results":    m.TableMaxResults,
			"volume_max_results":   m.VolumeMaxResults,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProviderShareAssetsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_max_results": types.Int64Type,
			"notebook_max_results": types.Int64Type,
			"provider_name":        types.StringType,
			"share_name":           types.StringType,
			"table_max_results":    types.Int64Type,
			"volume_max_results":   types.Int64Type,
		},
	}
}

// Response to ListProviderShareAssets, which contains the list of assets of a
// share.
type ListProviderShareAssetsResponse_SdkV2 struct {
	// The list of functions in the share.
	Functions types.List `tfsdk:"functions"`
	// The list of notebooks in the share.
	Notebooks types.List `tfsdk:"notebooks"`
	// The metadata of the share.
	Share types.List `tfsdk:"share"`
	// The list of tables in the share.
	Tables types.List `tfsdk:"tables"`
	// The list of volumes in the share.
	Volumes types.List `tfsdk:"volumes"`
}

func (to *ListProviderShareAssetsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProviderShareAssetsResponse_SdkV2) {
	if !from.Functions.IsNull() && !from.Functions.IsUnknown() && to.Functions.IsNull() && len(from.Functions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Functions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Functions = from.Functions
	}
	if !from.Notebooks.IsNull() && !from.Notebooks.IsUnknown() && to.Notebooks.IsNull() && len(from.Notebooks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Notebooks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Notebooks = from.Notebooks
	}
	if !from.Share.IsNull() && !from.Share.IsUnknown() {
		if toShare, ok := to.GetShare(ctx); ok {
			if fromShare, ok := from.GetShare(ctx); ok {
				// Recursively sync the fields of Share
				toShare.SyncFieldsDuringCreateOrUpdate(ctx, fromShare)
				to.SetShare(ctx, toShare)
			}
		}
	}
	if !from.Tables.IsNull() && !from.Tables.IsUnknown() && to.Tables.IsNull() && len(from.Tables.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tables, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tables = from.Tables
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() && to.Volumes.IsNull() && len(from.Volumes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Volumes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Volumes = from.Volumes
	}
}

func (to *ListProviderShareAssetsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProviderShareAssetsResponse_SdkV2) {
	if !from.Functions.IsNull() && !from.Functions.IsUnknown() && to.Functions.IsNull() && len(from.Functions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Functions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Functions = from.Functions
	}
	if !from.Notebooks.IsNull() && !from.Notebooks.IsUnknown() && to.Notebooks.IsNull() && len(from.Notebooks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Notebooks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Notebooks = from.Notebooks
	}
	if !from.Share.IsNull() && !from.Share.IsUnknown() {
		if toShare, ok := to.GetShare(ctx); ok {
			if fromShare, ok := from.GetShare(ctx); ok {
				toShare.SyncFieldsDuringRead(ctx, fromShare)
				to.SetShare(ctx, toShare)
			}
		}
	}
	if !from.Tables.IsNull() && !from.Tables.IsUnknown() && to.Tables.IsNull() && len(from.Tables.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tables, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tables = from.Tables
	}
	if !from.Volumes.IsNull() && !from.Volumes.IsUnknown() && to.Volumes.IsNull() && len(from.Volumes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Volumes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Volumes = from.Volumes
	}
}

func (m ListProviderShareAssetsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["functions"] = attrs["functions"].SetOptional()
	attrs["notebooks"] = attrs["notebooks"].SetOptional()
	attrs["share"] = attrs["share"].SetOptional()
	attrs["share"] = attrs["share"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["tables"] = attrs["tables"].SetOptional()
	attrs["volumes"] = attrs["volumes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProviderShareAssetsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProviderShareAssetsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"functions": reflect.TypeOf(DeltaSharingFunction_SdkV2{}),
		"notebooks": reflect.TypeOf(NotebookFile_SdkV2{}),
		"share":     reflect.TypeOf(Share_SdkV2{}),
		"tables":    reflect.TypeOf(Table_SdkV2{}),
		"volumes":   reflect.TypeOf(Volume_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderShareAssetsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListProviderShareAssetsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"functions": m.Functions,
			"notebooks": m.Notebooks,
			"share":     m.Share,
			"tables":    m.Tables,
			"volumes":   m.Volumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProviderShareAssetsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"functions": basetypes.ListType{
				ElemType: DeltaSharingFunction_SdkV2{}.Type(ctx),
			},
			"notebooks": basetypes.ListType{
				ElemType: NotebookFile_SdkV2{}.Type(ctx),
			},
			"share": basetypes.ListType{
				ElemType: Share_SdkV2{}.Type(ctx),
			},
			"tables": basetypes.ListType{
				ElemType: Table_SdkV2{}.Type(ctx),
			},
			"volumes": basetypes.ListType{
				ElemType: Volume_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFunctions returns the value of the Functions field in ListProviderShareAssetsResponse_SdkV2 as
// a slice of DeltaSharingFunction_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProviderShareAssetsResponse_SdkV2) GetFunctions(ctx context.Context) ([]DeltaSharingFunction_SdkV2, bool) {
	if m.Functions.IsNull() || m.Functions.IsUnknown() {
		return nil, false
	}
	var v []DeltaSharingFunction_SdkV2
	d := m.Functions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFunctions sets the value of the Functions field in ListProviderShareAssetsResponse_SdkV2.
func (m *ListProviderShareAssetsResponse_SdkV2) SetFunctions(ctx context.Context, v []DeltaSharingFunction_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["functions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Functions = types.ListValueMust(t, vs)
}

// GetNotebooks returns the value of the Notebooks field in ListProviderShareAssetsResponse_SdkV2 as
// a slice of NotebookFile_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProviderShareAssetsResponse_SdkV2) GetNotebooks(ctx context.Context) ([]NotebookFile_SdkV2, bool) {
	if m.Notebooks.IsNull() || m.Notebooks.IsUnknown() {
		return nil, false
	}
	var v []NotebookFile_SdkV2
	d := m.Notebooks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebooks sets the value of the Notebooks field in ListProviderShareAssetsResponse_SdkV2.
func (m *ListProviderShareAssetsResponse_SdkV2) SetNotebooks(ctx context.Context, v []NotebookFile_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notebooks"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Notebooks = types.ListValueMust(t, vs)
}

// GetShare returns the value of the Share field in ListProviderShareAssetsResponse_SdkV2 as
// a Share_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProviderShareAssetsResponse_SdkV2) GetShare(ctx context.Context) (Share_SdkV2, bool) {
	var e Share_SdkV2
	if m.Share.IsNull() || m.Share.IsUnknown() {
		return e, false
	}
	var v []Share_SdkV2
	d := m.Share.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetShare sets the value of the Share field in ListProviderShareAssetsResponse_SdkV2.
func (m *ListProviderShareAssetsResponse_SdkV2) SetShare(ctx context.Context, v Share_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["share"]
	m.Share = types.ListValueMust(t, vs)
}

// GetTables returns the value of the Tables field in ListProviderShareAssetsResponse_SdkV2 as
// a slice of Table_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProviderShareAssetsResponse_SdkV2) GetTables(ctx context.Context) ([]Table_SdkV2, bool) {
	if m.Tables.IsNull() || m.Tables.IsUnknown() {
		return nil, false
	}
	var v []Table_SdkV2
	d := m.Tables.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTables sets the value of the Tables field in ListProviderShareAssetsResponse_SdkV2.
func (m *ListProviderShareAssetsResponse_SdkV2) SetTables(ctx context.Context, v []Table_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tables"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tables = types.ListValueMust(t, vs)
}

// GetVolumes returns the value of the Volumes field in ListProviderShareAssetsResponse_SdkV2 as
// a slice of Volume_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProviderShareAssetsResponse_SdkV2) GetVolumes(ctx context.Context) ([]Volume_SdkV2, bool) {
	if m.Volumes.IsNull() || m.Volumes.IsUnknown() {
		return nil, false
	}
	var v []Volume_SdkV2
	d := m.Volumes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolumes sets the value of the Volumes field in ListProviderShareAssetsResponse_SdkV2.
func (m *ListProviderShareAssetsResponse_SdkV2) SetVolumes(ctx context.Context, v []Volume_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["volumes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Volumes = types.ListValueMust(t, vs)
}

type ListProviderSharesResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of provider shares.
	Shares types.List `tfsdk:"shares"`
}

func (to *ListProviderSharesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProviderSharesResponse_SdkV2) {
	if !from.Shares.IsNull() && !from.Shares.IsUnknown() && to.Shares.IsNull() && len(from.Shares.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Shares, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Shares = from.Shares
	}
}

func (to *ListProviderSharesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProviderSharesResponse_SdkV2) {
	if !from.Shares.IsNull() && !from.Shares.IsUnknown() && to.Shares.IsNull() && len(from.Shares.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Shares, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Shares = from.Shares
	}
}

func (m ListProviderSharesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["shares"] = attrs["shares"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProviderSharesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProviderSharesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"shares": reflect.TypeOf(ProviderShare_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderSharesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListProviderSharesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"shares":          m.Shares,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProviderSharesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shares": basetypes.ListType{
				ElemType: ProviderShare_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetShares returns the value of the Shares field in ListProviderSharesResponse_SdkV2 as
// a slice of ProviderShare_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProviderSharesResponse_SdkV2) GetShares(ctx context.Context) ([]ProviderShare_SdkV2, bool) {
	if m.Shares.IsNull() || m.Shares.IsUnknown() {
		return nil, false
	}
	var v []ProviderShare_SdkV2
	d := m.Shares.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetShares sets the value of the Shares field in ListProviderSharesResponse_SdkV2.
func (m *ListProviderSharesResponse_SdkV2) SetShares(ctx context.Context, v []ProviderShare_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["shares"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Shares = types.ListValueMust(t, vs)
}

type ListProvidersRequest_SdkV2 struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	DataProviderGlobalMetastoreId types.String `tfsdk:"-"`
	// Maximum number of providers to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid providers are returned (not
	// recommended). - Note: The number of returned providers might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further providers can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListProvidersRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProvidersRequest_SdkV2) {
}

func (to *ListProvidersRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProvidersRequest_SdkV2) {
}

func (m ListProvidersRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_provider_global_metastore_id"] = attrs["data_provider_global_metastore_id"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProvidersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProvidersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListProvidersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_provider_global_metastore_id": m.DataProviderGlobalMetastoreId,
			"max_results":                       m.MaxResults,
			"page_token":                        m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProvidersRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_provider_global_metastore_id": types.StringType,
			"max_results":                       types.Int64Type,
			"page_token":                        types.StringType,
		},
	}
}

type ListProvidersResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of provider information objects.
	Providers types.List `tfsdk:"providers"`
}

func (to *ListProvidersResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProvidersResponse_SdkV2) {
	if !from.Providers.IsNull() && !from.Providers.IsUnknown() && to.Providers.IsNull() && len(from.Providers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Providers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Providers = from.Providers
	}
}

func (to *ListProvidersResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProvidersResponse_SdkV2) {
	if !from.Providers.IsNull() && !from.Providers.IsUnknown() && to.Providers.IsNull() && len(from.Providers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Providers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Providers = from.Providers
	}
}

func (m ListProvidersResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["providers"] = attrs["providers"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProvidersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProvidersResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"providers": reflect.TypeOf(ProviderInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListProvidersResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"providers":       m.Providers,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProvidersResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"providers": basetypes.ListType{
				ElemType: ProviderInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetProviders returns the value of the Providers field in ListProvidersResponse_SdkV2 as
// a slice of ProviderInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProvidersResponse_SdkV2) GetProviders(ctx context.Context) ([]ProviderInfo_SdkV2, bool) {
	if m.Providers.IsNull() || m.Providers.IsUnknown() {
		return nil, false
	}
	var v []ProviderInfo_SdkV2
	d := m.Providers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviders sets the value of the Providers field in ListProvidersResponse_SdkV2.
func (m *ListProvidersResponse_SdkV2) SetProviders(ctx context.Context, v []ProviderInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["providers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Providers = types.ListValueMust(t, vs)
}

type ListRecipientsRequest_SdkV2 struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"-"`
	// Maximum number of recipients to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid recipients are returned (not
	// recommended). - Note: The number of returned recipients might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further recipients can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListRecipientsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRecipientsRequest_SdkV2) {
}

func (to *ListRecipientsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListRecipientsRequest_SdkV2) {
}

func (m ListRecipientsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_recipient_global_metastore_id"] = attrs["data_recipient_global_metastore_id"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRecipientsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListRecipientsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRecipientsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListRecipientsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_recipient_global_metastore_id": m.DataRecipientGlobalMetastoreId,
			"max_results":                        m.MaxResults,
			"page_token":                         m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRecipientsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_recipient_global_metastore_id": types.StringType,
			"max_results":                        types.Int64Type,
			"page_token":                         types.StringType,
		},
	}
}

type ListRecipientsResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of recipient information objects.
	Recipients types.List `tfsdk:"recipients"`
}

func (to *ListRecipientsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRecipientsResponse_SdkV2) {
	if !from.Recipients.IsNull() && !from.Recipients.IsUnknown() && to.Recipients.IsNull() && len(from.Recipients.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Recipients, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Recipients = from.Recipients
	}
}

func (to *ListRecipientsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListRecipientsResponse_SdkV2) {
	if !from.Recipients.IsNull() && !from.Recipients.IsUnknown() && to.Recipients.IsNull() && len(from.Recipients.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Recipients, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Recipients = from.Recipients
	}
}

func (m ListRecipientsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["recipients"] = attrs["recipients"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRecipientsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListRecipientsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"recipients": reflect.TypeOf(RecipientInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRecipientsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListRecipientsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"recipients":      m.Recipients,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRecipientsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"recipients": basetypes.ListType{
				ElemType: RecipientInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRecipients returns the value of the Recipients field in ListRecipientsResponse_SdkV2 as
// a slice of RecipientInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListRecipientsResponse_SdkV2) GetRecipients(ctx context.Context) ([]RecipientInfo_SdkV2, bool) {
	if m.Recipients.IsNull() || m.Recipients.IsUnknown() {
		return nil, false
	}
	var v []RecipientInfo_SdkV2
	d := m.Recipients.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRecipients sets the value of the Recipients field in ListRecipientsResponse_SdkV2.
func (m *ListRecipientsResponse_SdkV2) SetRecipients(ctx context.Context, v []RecipientInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["recipients"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Recipients = types.ListValueMust(t, vs)
}

type ListSharesRequest_SdkV2 struct {
	// Maximum number of shares to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid shares are returned (not
	// recommended). - Note: The number of returned shares might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further shares can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Name of the provider in which to list shares.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListSharesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSharesRequest_SdkV2) {
}

func (to *ListSharesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSharesRequest_SdkV2) {
}

func (m ListSharesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSharesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSharesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSharesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSharesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": m.MaxResults,
			"name":        m.Name,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSharesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"name":        types.StringType,
			"page_token":  types.StringType,
		},
	}
}

type ListSharesResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of data share information objects.
	Shares types.List `tfsdk:"shares"`
}

func (to *ListSharesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSharesResponse_SdkV2) {
	if !from.Shares.IsNull() && !from.Shares.IsUnknown() && to.Shares.IsNull() && len(from.Shares.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Shares, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Shares = from.Shares
	}
}

func (to *ListSharesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSharesResponse_SdkV2) {
	if !from.Shares.IsNull() && !from.Shares.IsUnknown() && to.Shares.IsNull() && len(from.Shares.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Shares, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Shares = from.Shares
	}
}

func (m ListSharesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["shares"] = attrs["shares"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSharesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSharesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"shares": reflect.TypeOf(ShareInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSharesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSharesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"shares":          m.Shares,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSharesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shares": basetypes.ListType{
				ElemType: ShareInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetShares returns the value of the Shares field in ListSharesResponse_SdkV2 as
// a slice of ShareInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSharesResponse_SdkV2) GetShares(ctx context.Context) ([]ShareInfo_SdkV2, bool) {
	if m.Shares.IsNull() || m.Shares.IsUnknown() {
		return nil, false
	}
	var v []ShareInfo_SdkV2
	d := m.Shares.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetShares sets the value of the Shares field in ListSharesResponse_SdkV2.
func (m *ListSharesResponse_SdkV2) SetShares(ctx context.Context, v []ShareInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["shares"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Shares = types.ListValueMust(t, vs)
}

type NotebookFile_SdkV2 struct {
	// The comment of the notebook file.
	Comment types.String `tfsdk:"comment"`
	// The id of the notebook file.
	Id types.String `tfsdk:"id"`
	// Name of the notebook file.
	Name types.String `tfsdk:"name"`
	// The name of the share that the notebook file belongs to.
	Share types.String `tfsdk:"share"`
	// The id of the share that the notebook file belongs to.
	ShareId types.String `tfsdk:"share_id"`
	// The tags of the notebook file.
	Tags types.List `tfsdk:"tags"`
}

func (to *NotebookFile_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotebookFile_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *NotebookFile_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NotebookFile_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m NotebookFile_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["share"] = attrs["share"].SetOptional()
	attrs["share_id"] = attrs["share_id"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotebookFile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NotebookFile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(catalog_tf.TagKeyValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookFile_SdkV2
// only implements ToObjectValue() and Type().
func (m NotebookFile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":  m.Comment,
			"id":       m.Id,
			"name":     m.Name,
			"share":    m.Share,
			"share_id": m.ShareId,
			"tags":     m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NotebookFile_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":  types.StringType,
			"id":       types.StringType,
			"name":     types.StringType,
			"share":    types.StringType,
			"share_id": types.StringType,
			"tags": basetypes.ListType{
				ElemType: catalog_tf.TagKeyValue_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in NotebookFile_SdkV2 as
// a slice of catalog_tf.TagKeyValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *NotebookFile_SdkV2) GetTags(ctx context.Context) ([]catalog_tf.TagKeyValue_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.TagKeyValue_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in NotebookFile_SdkV2.
func (m *NotebookFile_SdkV2) SetTags(ctx context.Context, v []catalog_tf.TagKeyValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// Specifies the policy to use for validating OIDC claims in your federated
// tokens from Delta Sharing Clients. Refer to
// https://docs.databricks.com/en/delta-sharing/create-recipient-oidc-fed for
// more details.
type OidcFederationPolicy_SdkV2 struct {
	// The allowed token audiences, as specified in the 'aud' claim of federated
	// tokens. The audience identifier is intended to represent the recipient of
	// the token. Can be any non-empty string value. As long as the audience in
	// the token matches at least one audience in the policy,
	Audiences types.List `tfsdk:"audiences"`
	// The required token issuer, as specified in the 'iss' claim of federated
	// tokens.
	Issuer types.String `tfsdk:"issuer"`
	// The required token subject, as specified in the subject claim of
	// federated tokens. The subject claim identifies the identity of the user
	// or machine accessing the resource. Examples for Entra ID (AAD): - U2M
	// flow (group access): If the subject claim is `groups`, this must be the
	// Object ID of the group in Entra ID. - U2M flow (user access): If the
	// subject claim is `oid`, this must be the Object ID of the user in Entra
	// ID. - M2M flow (OAuth App access): If the subject claim is `azp`, this
	// must be the client ID of the OAuth app registered in Entra ID.
	Subject types.String `tfsdk:"subject"`
	// The claim that contains the subject of the token. Depending on the
	// identity provider and the use case (U2M or M2M), this can vary: - For
	// Entra ID (AAD): * U2M flow (group access): Use `groups`. * U2M flow (user
	// access): Use `oid`. * M2M flow (OAuth App access): Use `azp`. - For other
	// IdPs, refer to the specific IdP documentation.
	//
	// Supported `subject_claim` values are: - `oid`: Object ID of the user. -
	// `azp`: Client ID of the OAuth app. - `groups`: Object ID of the group. -
	// `sub`: Subject identifier for other use cases.
	SubjectClaim types.String `tfsdk:"subject_claim"`
}

func (to *OidcFederationPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from OidcFederationPolicy_SdkV2) {
	if !from.Audiences.IsNull() && !from.Audiences.IsUnknown() && to.Audiences.IsNull() && len(from.Audiences.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Audiences, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Audiences = from.Audiences
	}
}

func (to *OidcFederationPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from OidcFederationPolicy_SdkV2) {
	if !from.Audiences.IsNull() && !from.Audiences.IsUnknown() && to.Audiences.IsNull() && len(from.Audiences.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Audiences, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Audiences = from.Audiences
	}
}

func (m OidcFederationPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["audiences"] = attrs["audiences"].SetOptional()
	attrs["issuer"] = attrs["issuer"].SetRequired()
	attrs["subject"] = attrs["subject"].SetRequired()
	attrs["subject_claim"] = attrs["subject_claim"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OidcFederationPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m OidcFederationPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"audiences": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OidcFederationPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m OidcFederationPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"audiences":     m.Audiences,
			"issuer":        m.Issuer,
			"subject":       m.Subject,
			"subject_claim": m.SubjectClaim,
		})
}

// Type implements basetypes.ObjectValuable.
func (m OidcFederationPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"audiences": basetypes.ListType{
				ElemType: types.StringType,
			},
			"issuer":        types.StringType,
			"subject":       types.StringType,
			"subject_claim": types.StringType,
		},
	}
}

// GetAudiences returns the value of the Audiences field in OidcFederationPolicy_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *OidcFederationPolicy_SdkV2) GetAudiences(ctx context.Context) ([]types.String, bool) {
	if m.Audiences.IsNull() || m.Audiences.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Audiences.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAudiences sets the value of the Audiences field in OidcFederationPolicy_SdkV2.
func (m *OidcFederationPolicy_SdkV2) SetAudiences(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["audiences"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Audiences = types.ListValueMust(t, vs)
}

type Partition_SdkV2 struct {
	// An array of partition values.
	Values types.List `tfsdk:"value"`
}

func (to *Partition_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Partition_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *Partition_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Partition_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (m Partition_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Partition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Partition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(PartitionValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Partition_SdkV2
// only implements ToObjectValue() and Type().
func (m Partition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Partition_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": basetypes.ListType{
				ElemType: PartitionValue_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in Partition_SdkV2 as
// a slice of PartitionValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Partition_SdkV2) GetValues(ctx context.Context) ([]PartitionValue_SdkV2, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []PartitionValue_SdkV2
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in Partition_SdkV2.
func (m *Partition_SdkV2) SetValues(ctx context.Context, v []PartitionValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

type PartitionValue_SdkV2 struct {
	// The name of the partition column.
	Name types.String `tfsdk:"name"`
	// The operator to apply for the value.
	Op types.String `tfsdk:"op"`
	// The key of a Delta Sharing recipient's property. For example
	// "databricks-account-id". When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey types.String `tfsdk:"recipient_property_key"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value types.String `tfsdk:"value"`
}

func (to *PartitionValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PartitionValue_SdkV2) {
}

func (to *PartitionValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PartitionValue_SdkV2) {
}

func (m PartitionValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["op"] = attrs["op"].SetOptional()
	attrs["recipient_property_key"] = attrs["recipient_property_key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartitionValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PartitionValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartitionValue_SdkV2
// only implements ToObjectValue() and Type().
func (m PartitionValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":                   m.Name,
			"op":                     m.Op,
			"recipient_property_key": m.RecipientPropertyKey,
			"value":                  m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PartitionValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":                   types.StringType,
			"op":                     types.StringType,
			"recipient_property_key": types.StringType,
			"value":                  types.StringType,
		},
	}
}

type PermissionsChange_SdkV2 struct {
	// The set of privileges to add.
	Add types.List `tfsdk:"add"`
	// The principal whose privileges we are changing. Only one of principal or
	// principal_id should be specified, never both at the same time.
	Principal types.String `tfsdk:"principal"`
	// The set of privileges to remove.
	Remove types.List `tfsdk:"remove"`
}

func (to *PermissionsChange_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PermissionsChange_SdkV2) {
	if !from.Add.IsNull() && !from.Add.IsUnknown() && to.Add.IsNull() && len(from.Add.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Add, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Add = from.Add
	}
	if !from.Remove.IsNull() && !from.Remove.IsUnknown() && to.Remove.IsNull() && len(from.Remove.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Remove, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Remove = from.Remove
	}
}

func (to *PermissionsChange_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PermissionsChange_SdkV2) {
	if !from.Add.IsNull() && !from.Add.IsUnknown() && to.Add.IsNull() && len(from.Add.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Add, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Add = from.Add
	}
	if !from.Remove.IsNull() && !from.Remove.IsUnknown() && to.Remove.IsNull() && len(from.Remove.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Remove, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Remove = from.Remove
	}
}

func (m PermissionsChange_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["add"] = attrs["add"].SetOptional()
	attrs["principal"] = attrs["principal"].SetOptional()
	attrs["remove"] = attrs["remove"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionsChange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PermissionsChange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"add":    reflect.TypeOf(types.String{}),
		"remove": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionsChange_SdkV2
// only implements ToObjectValue() and Type().
func (m PermissionsChange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"add":       m.Add,
			"principal": m.Principal,
			"remove":    m.Remove,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PermissionsChange_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"add": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal": types.StringType,
			"remove": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAdd returns the value of the Add field in PermissionsChange_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PermissionsChange_SdkV2) GetAdd(ctx context.Context) ([]types.String, bool) {
	if m.Add.IsNull() || m.Add.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Add.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAdd sets the value of the Add field in PermissionsChange_SdkV2.
func (m *PermissionsChange_SdkV2) SetAdd(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["add"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Add = types.ListValueMust(t, vs)
}

// GetRemove returns the value of the Remove field in PermissionsChange_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PermissionsChange_SdkV2) GetRemove(ctx context.Context) ([]types.String, bool) {
	if m.Remove.IsNull() || m.Remove.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Remove.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRemove sets the value of the Remove field in PermissionsChange_SdkV2.
func (m *PermissionsChange_SdkV2) SetRemove(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["remove"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Remove = types.ListValueMust(t, vs)
}

type PrivilegeAssignment_SdkV2 struct {
	// The principal (user email address or group name). For deleted principals,
	// `principal` is empty while `principal_id` is populated.
	Principal types.String `tfsdk:"principal"`
	// The privileges assigned to the principal.
	Privileges types.List `tfsdk:"privileges"`
}

func (to *PrivilegeAssignment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivilegeAssignment_SdkV2) {
	if !from.Privileges.IsNull() && !from.Privileges.IsUnknown() && to.Privileges.IsNull() && len(from.Privileges.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Privileges, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Privileges = from.Privileges
	}
}

func (to *PrivilegeAssignment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PrivilegeAssignment_SdkV2) {
	if !from.Privileges.IsNull() && !from.Privileges.IsUnknown() && to.Privileges.IsNull() && len(from.Privileges.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Privileges, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Privileges = from.Privileges
	}
}

func (m PrivilegeAssignment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal"] = attrs["principal"].SetOptional()
	attrs["privileges"] = attrs["privileges"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivilegeAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PrivilegeAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privileges": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivilegeAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (m PrivilegeAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal":  m.Principal,
			"privileges": m.Privileges,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivilegeAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"privileges": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPrivileges returns the value of the Privileges field in PrivilegeAssignment_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivilegeAssignment_SdkV2) GetPrivileges(ctx context.Context) ([]types.String, bool) {
	if m.Privileges.IsNull() || m.Privileges.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Privileges.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivileges sets the value of the Privileges field in PrivilegeAssignment_SdkV2.
func (m *PrivilegeAssignment_SdkV2) SetPrivileges(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["privileges"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Privileges = types.ListValueMust(t, vs)
}

type ProviderInfo_SdkV2 struct {
	AuthenticationType types.String `tfsdk:"authentication_type"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud types.String `tfsdk:"cloud"`
	// Description about the provider.
	Comment types.String `tfsdk:"comment"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of Provider creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format __cloud__:__region__:__metastore-uuid__.
	DataProviderGlobalMetastoreId types.String `tfsdk:"data_provider_global_metastore_id"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The name of the Provider.
	Name types.String `tfsdk:"name"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN` or `OAUTH_CLIENT_CREDENTIALS`.
	RecipientProfile types.List `tfsdk:"recipient_profile"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified Provider.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (to *ProviderInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProviderInfo_SdkV2) {
	if !from.RecipientProfile.IsNull() && !from.RecipientProfile.IsUnknown() {
		if toRecipientProfile, ok := to.GetRecipientProfile(ctx); ok {
			if fromRecipientProfile, ok := from.GetRecipientProfile(ctx); ok {
				// Recursively sync the fields of RecipientProfile
				toRecipientProfile.SyncFieldsDuringCreateOrUpdate(ctx, fromRecipientProfile)
				to.SetRecipientProfile(ctx, toRecipientProfile)
			}
		}
	}
}

func (to *ProviderInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProviderInfo_SdkV2) {
	if !from.RecipientProfile.IsNull() && !from.RecipientProfile.IsUnknown() {
		if toRecipientProfile, ok := to.GetRecipientProfile(ctx); ok {
			if fromRecipientProfile, ok := from.GetRecipientProfile(ctx); ok {
				toRecipientProfile.SyncFieldsDuringRead(ctx, fromRecipientProfile)
				to.SetRecipientProfile(ctx, toRecipientProfile)
			}
		}
	}
}

func (m ProviderInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["authentication_type"] = attrs["authentication_type"].SetOptional()
	attrs["cloud"] = attrs["cloud"].SetComputed()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["data_provider_global_metastore_id"] = attrs["data_provider_global_metastore_id"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["recipient_profile"] = attrs["recipient_profile"].SetOptional()
	attrs["recipient_profile"] = attrs["recipient_profile"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["recipient_profile_str"] = attrs["recipient_profile_str"].SetOptional()
	attrs["region"] = attrs["region"].SetComputed()
	attrs["updated_at"] = attrs["updated_at"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProviderInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"recipient_profile": reflect.TypeOf(RecipientProfile_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m ProviderInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authentication_type":               m.AuthenticationType,
			"cloud":                             m.Cloud,
			"comment":                           m.Comment,
			"created_at":                        m.CreatedAt,
			"created_by":                        m.CreatedBy,
			"data_provider_global_metastore_id": m.DataProviderGlobalMetastoreId,
			"metastore_id":                      m.MetastoreId,
			"name":                              m.Name,
			"owner":                             m.Owner,
			"recipient_profile":                 m.RecipientProfile,
			"recipient_profile_str":             m.RecipientProfileStr,
			"region":                            m.Region,
			"updated_at":                        m.UpdatedAt,
			"updated_by":                        m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProviderInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authentication_type":               types.StringType,
			"cloud":                             types.StringType,
			"comment":                           types.StringType,
			"created_at":                        types.Int64Type,
			"created_by":                        types.StringType,
			"data_provider_global_metastore_id": types.StringType,
			"metastore_id":                      types.StringType,
			"name":                              types.StringType,
			"owner":                             types.StringType,
			"recipient_profile": basetypes.ListType{
				ElemType: RecipientProfile_SdkV2{}.Type(ctx),
			},
			"recipient_profile_str": types.StringType,
			"region":                types.StringType,
			"updated_at":            types.Int64Type,
			"updated_by":            types.StringType,
		},
	}
}

// GetRecipientProfile returns the value of the RecipientProfile field in ProviderInfo_SdkV2 as
// a RecipientProfile_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProviderInfo_SdkV2) GetRecipientProfile(ctx context.Context) (RecipientProfile_SdkV2, bool) {
	var e RecipientProfile_SdkV2
	if m.RecipientProfile.IsNull() || m.RecipientProfile.IsUnknown() {
		return e, false
	}
	var v []RecipientProfile_SdkV2
	d := m.RecipientProfile.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRecipientProfile sets the value of the RecipientProfile field in ProviderInfo_SdkV2.
func (m *ProviderInfo_SdkV2) SetRecipientProfile(ctx context.Context, v RecipientProfile_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["recipient_profile"]
	m.RecipientProfile = types.ListValueMust(t, vs)
}

type ProviderShare_SdkV2 struct {
	// The name of the Provider Share.
	Name types.String `tfsdk:"name"`
}

func (to *ProviderShare_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProviderShare_SdkV2) {
}

func (to *ProviderShare_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProviderShare_SdkV2) {
}

func (m ProviderShare_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderShare.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProviderShare_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderShare_SdkV2
// only implements ToObjectValue() and Type().
func (m ProviderShare_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProviderShare_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type RecipientInfo_SdkV2 struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated types.Bool `tfsdk:"activated"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url"`

	AuthenticationType types.String `tfsdk:"authentication_type"`
	// Cloud vendor of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Cloud types.String `tfsdk:"cloud"`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of recipient creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// IP Access List
	IpAccessList types.List `tfsdk:"ip_access_list"`
	// Unique identifier of recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of Recipient.
	Name types.String `tfsdk:"name"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs types.List `tfsdk:"properties_kvpairs"`
	// Cloud region of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens types.List `tfsdk:"tokens"`
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of recipient updater.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (to *RecipientInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RecipientInfo_SdkV2) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				// Recursively sync the fields of IpAccessList
				toIpAccessList.SyncFieldsDuringCreateOrUpdate(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
	if !from.PropertiesKvpairs.IsNull() && !from.PropertiesKvpairs.IsUnknown() {
		if toPropertiesKvpairs, ok := to.GetPropertiesKvpairs(ctx); ok {
			if fromPropertiesKvpairs, ok := from.GetPropertiesKvpairs(ctx); ok {
				// Recursively sync the fields of PropertiesKvpairs
				toPropertiesKvpairs.SyncFieldsDuringCreateOrUpdate(ctx, fromPropertiesKvpairs)
				to.SetPropertiesKvpairs(ctx, toPropertiesKvpairs)
			}
		}
	}
	if !from.Tokens.IsNull() && !from.Tokens.IsUnknown() && to.Tokens.IsNull() && len(from.Tokens.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tokens, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tokens = from.Tokens
	}
}

func (to *RecipientInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RecipientInfo_SdkV2) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				toIpAccessList.SyncFieldsDuringRead(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
	if !from.PropertiesKvpairs.IsNull() && !from.PropertiesKvpairs.IsUnknown() {
		if toPropertiesKvpairs, ok := to.GetPropertiesKvpairs(ctx); ok {
			if fromPropertiesKvpairs, ok := from.GetPropertiesKvpairs(ctx); ok {
				toPropertiesKvpairs.SyncFieldsDuringRead(ctx, fromPropertiesKvpairs)
				to.SetPropertiesKvpairs(ctx, toPropertiesKvpairs)
			}
		}
	}
	if !from.Tokens.IsNull() && !from.Tokens.IsUnknown() && to.Tokens.IsNull() && len(from.Tokens.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tokens, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tokens = from.Tokens
	}
}

func (m RecipientInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activated"] = attrs["activated"].SetComputed()
	attrs["activation_url"] = attrs["activation_url"].SetComputed()
	attrs["authentication_type"] = attrs["authentication_type"].SetOptional()
	attrs["cloud"] = attrs["cloud"].SetComputed()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["data_recipient_global_metastore_id"] = attrs["data_recipient_global_metastore_id"].SetOptional()
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties_kvpairs"] = attrs["properties_kvpairs"].SetOptional()
	attrs["properties_kvpairs"] = attrs["properties_kvpairs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["region"] = attrs["region"].SetComputed()
	attrs["sharing_code"] = attrs["sharing_code"].SetOptional()
	attrs["tokens"] = attrs["tokens"].SetComputed()
	attrs["updated_at"] = attrs["updated_at"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RecipientInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RecipientInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList_SdkV2{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs_SdkV2{}),
		"tokens":             reflect.TypeOf(RecipientTokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RecipientInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m RecipientInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activated":                          m.Activated,
			"activation_url":                     m.ActivationUrl,
			"authentication_type":                m.AuthenticationType,
			"cloud":                              m.Cloud,
			"comment":                            m.Comment,
			"created_at":                         m.CreatedAt,
			"created_by":                         m.CreatedBy,
			"data_recipient_global_metastore_id": m.DataRecipientGlobalMetastoreId,
			"expiration_time":                    m.ExpirationTime,
			"ip_access_list":                     m.IpAccessList,
			"metastore_id":                       m.MetastoreId,
			"name":                               m.Name,
			"owner":                              m.Owner,
			"properties_kvpairs":                 m.PropertiesKvpairs,
			"region":                             m.Region,
			"sharing_code":                       m.SharingCode,
			"tokens":                             m.Tokens,
			"updated_at":                         m.UpdatedAt,
			"updated_by":                         m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RecipientInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activated":                          types.BoolType,
			"activation_url":                     types.StringType,
			"authentication_type":                types.StringType,
			"cloud":                              types.StringType,
			"comment":                            types.StringType,
			"created_at":                         types.Int64Type,
			"created_by":                         types.StringType,
			"data_recipient_global_metastore_id": types.StringType,
			"expiration_time":                    types.Int64Type,
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessList_SdkV2{}.Type(ctx),
			},
			"metastore_id": types.StringType,
			"name":         types.StringType,
			"owner":        types.StringType,
			"properties_kvpairs": basetypes.ListType{
				ElemType: SecurablePropertiesKvPairs_SdkV2{}.Type(ctx),
			},
			"region":       types.StringType,
			"sharing_code": types.StringType,
			"tokens": basetypes.ListType{
				ElemType: RecipientTokenInfo_SdkV2{}.Type(ctx),
			},
			"updated_at": types.Int64Type,
			"updated_by": types.StringType,
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in RecipientInfo_SdkV2 as
// a IpAccessList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RecipientInfo_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessList_SdkV2, bool) {
	var e IpAccessList_SdkV2
	if m.IpAccessList.IsNull() || m.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessList_SdkV2
	d := m.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in RecipientInfo_SdkV2.
func (m *RecipientInfo_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	m.IpAccessList = types.ListValueMust(t, vs)
}

// GetPropertiesKvpairs returns the value of the PropertiesKvpairs field in RecipientInfo_SdkV2 as
// a SecurablePropertiesKvPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RecipientInfo_SdkV2) GetPropertiesKvpairs(ctx context.Context) (SecurablePropertiesKvPairs_SdkV2, bool) {
	var e SecurablePropertiesKvPairs_SdkV2
	if m.PropertiesKvpairs.IsNull() || m.PropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []SecurablePropertiesKvPairs_SdkV2
	d := m.PropertiesKvpairs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPropertiesKvpairs sets the value of the PropertiesKvpairs field in RecipientInfo_SdkV2.
func (m *RecipientInfo_SdkV2) SetPropertiesKvpairs(ctx context.Context, v SecurablePropertiesKvPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["properties_kvpairs"]
	m.PropertiesKvpairs = types.ListValueMust(t, vs)
}

// GetTokens returns the value of the Tokens field in RecipientInfo_SdkV2 as
// a slice of RecipientTokenInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RecipientInfo_SdkV2) GetTokens(ctx context.Context) ([]RecipientTokenInfo_SdkV2, bool) {
	if m.Tokens.IsNull() || m.Tokens.IsUnknown() {
		return nil, false
	}
	var v []RecipientTokenInfo_SdkV2
	d := m.Tokens.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokens sets the value of the Tokens field in RecipientInfo_SdkV2.
func (m *RecipientInfo_SdkV2) SetTokens(ctx context.Context, v []RecipientTokenInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tokens"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tokens = types.ListValueMust(t, vs)
}

type RecipientProfile_SdkV2 struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearer_token"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version"`
}

func (to *RecipientProfile_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RecipientProfile_SdkV2) {
}

func (to *RecipientProfile_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RecipientProfile_SdkV2) {
}

func (m RecipientProfile_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bearer_token"] = attrs["bearer_token"].SetOptional()
	attrs["endpoint"] = attrs["endpoint"].SetOptional()
	attrs["share_credentials_version"] = attrs["share_credentials_version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RecipientProfile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RecipientProfile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RecipientProfile_SdkV2
// only implements ToObjectValue() and Type().
func (m RecipientProfile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bearer_token":              m.BearerToken,
			"endpoint":                  m.Endpoint,
			"share_credentials_version": m.ShareCredentialsVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RecipientProfile_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearer_token":              types.StringType,
			"endpoint":                  types.StringType,
			"share_credentials_version": types.Int64Type,
		},
	}
}

type RecipientTokenInfo_SdkV2 struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url"`
	// Time at which this recipient token was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of recipient token creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// Unique ID of the recipient token.
	Id types.String `tfsdk:"id"`
	// Time at which this recipient token was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of recipient token updater.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (to *RecipientTokenInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RecipientTokenInfo_SdkV2) {
}

func (to *RecipientTokenInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RecipientTokenInfo_SdkV2) {
}

func (m RecipientTokenInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activation_url"] = attrs["activation_url"].SetComputed()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["expiration_time"] = attrs["expiration_time"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["updated_at"] = attrs["updated_at"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RecipientTokenInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RecipientTokenInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RecipientTokenInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m RecipientTokenInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activation_url":  m.ActivationUrl,
			"created_at":      m.CreatedAt,
			"created_by":      m.CreatedBy,
			"expiration_time": m.ExpirationTime,
			"id":              m.Id,
			"updated_at":      m.UpdatedAt,
			"updated_by":      m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RecipientTokenInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activation_url":  types.StringType,
			"created_at":      types.Int64Type,
			"created_by":      types.StringType,
			"expiration_time": types.Int64Type,
			"id":              types.StringType,
			"updated_at":      types.Int64Type,
			"updated_by":      types.StringType,
		},
	}
}

type RegisteredModelAlias_SdkV2 struct {
	// Name of the alias.
	AliasName types.String `tfsdk:"alias_name"`
	// Numeric model version that alias will reference.
	VersionNum types.Int64 `tfsdk:"version_num"`
}

func (to *RegisteredModelAlias_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelAlias_SdkV2) {
}

func (to *RegisteredModelAlias_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelAlias_SdkV2) {
}

func (m RegisteredModelAlias_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alias_name"] = attrs["alias_name"].SetOptional()
	attrs["version_num"] = attrs["version_num"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelAlias.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RegisteredModelAlias_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAlias_SdkV2
// only implements ToObjectValue() and Type().
func (m RegisteredModelAlias_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias_name":  m.AliasName,
			"version_num": m.VersionNum,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegisteredModelAlias_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias_name":  types.StringType,
			"version_num": types.Int64Type,
		},
	}
}

type RetrieveTokenRequest_SdkV2 struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

func (to *RetrieveTokenRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RetrieveTokenRequest_SdkV2) {
}

func (to *RetrieveTokenRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RetrieveTokenRequest_SdkV2) {
}

func (m RetrieveTokenRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activation_url"] = attrs["activation_url"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RetrieveTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RetrieveTokenRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RetrieveTokenRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RetrieveTokenRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activation_url": m.ActivationUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RetrieveTokenRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activation_url": types.StringType,
		},
	}
}

type RetrieveTokenResponse_SdkV2 struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearer_token"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.String `tfsdk:"expiration_time"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version"`
}

func (to *RetrieveTokenResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RetrieveTokenResponse_SdkV2) {
}

func (to *RetrieveTokenResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RetrieveTokenResponse_SdkV2) {
}

func (m RetrieveTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bearer_token"] = attrs["bearer_token"].SetComputed()
	attrs["endpoint"] = attrs["endpoint"].SetComputed()
	attrs["expiration_time"] = attrs["expiration_time"].SetComputed()
	attrs["share_credentials_version"] = attrs["share_credentials_version"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RetrieveTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RetrieveTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RetrieveTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RetrieveTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bearer_token":              m.BearerToken,
			"endpoint":                  m.Endpoint,
			"expiration_time":           m.ExpirationTime,
			"share_credentials_version": m.ShareCredentialsVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RetrieveTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearer_token":              types.StringType,
			"endpoint":                  types.StringType,
			"expiration_time":           types.StringType,
			"share_credentials_version": types.Int64Type,
		},
	}
}

type RotateRecipientToken_SdkV2 struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds types.Int64 `tfsdk:"existing_token_expire_in_seconds"`
	// The name of the Recipient.
	Name types.String `tfsdk:"-"`
}

func (to *RotateRecipientToken_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RotateRecipientToken_SdkV2) {
}

func (to *RotateRecipientToken_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RotateRecipientToken_SdkV2) {
}

func (m RotateRecipientToken_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["existing_token_expire_in_seconds"] = attrs["existing_token_expire_in_seconds"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RotateRecipientToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RotateRecipientToken_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RotateRecipientToken_SdkV2
// only implements ToObjectValue() and Type().
func (m RotateRecipientToken_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"existing_token_expire_in_seconds": m.ExistingTokenExpireInSeconds,
			"name":                             m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RotateRecipientToken_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"existing_token_expire_in_seconds": types.Int64Type,
			"name":                             types.StringType,
		},
	}
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs_SdkV2 struct {
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`
}

func (to *SecurablePropertiesKvPairs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SecurablePropertiesKvPairs_SdkV2) {
}

func (to *SecurablePropertiesKvPairs_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SecurablePropertiesKvPairs_SdkV2) {
}

func (m SecurablePropertiesKvPairs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["properties"] = attrs["properties"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SecurablePropertiesKvPairs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SecurablePropertiesKvPairs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecurablePropertiesKvPairs_SdkV2
// only implements ToObjectValue() and Type().
func (m SecurablePropertiesKvPairs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"properties": m.Properties,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SecurablePropertiesKvPairs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetProperties returns the value of the Properties field in SecurablePropertiesKvPairs_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SecurablePropertiesKvPairs_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
	if m.Properties.IsNull() || m.Properties.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Properties.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProperties sets the value of the Properties field in SecurablePropertiesKvPairs_SdkV2.
func (m *SecurablePropertiesKvPairs_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Properties = types.MapValueMust(t, vs)
}

type Share_SdkV2 struct {
	Id types.String `tfsdk:"id"`

	Name types.String `tfsdk:"name"`
}

func (to *Share_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Share_SdkV2) {
}

func (to *Share_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Share_SdkV2) {
}

func (m Share_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Share.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Share_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Share_SdkV2
// only implements ToObjectValue() and Type().
func (m Share_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":   m.Id,
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Share_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":   types.StringType,
			"name": types.StringType,
		},
	}
}

type ShareInfo_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of share creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Name of the share.
	Name types.String `tfsdk:"name"`
	// A list of shared data objects within the share.
	Objects types.List `tfsdk:"object"`
	// Username of current owner of share.
	Owner          types.String `tfsdk:"owner"`
	EffectiveOwner types.String `tfsdk:"effective_owner"`
	// Storage Location URL (full path) for the share.
	StorageLocation types.String `tfsdk:"storage_location"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of share updater.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (to *ShareInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ShareInfo_SdkV2) {
	if !from.Objects.IsNull() && !from.Objects.IsUnknown() && to.Objects.IsNull() && len(from.Objects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Objects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Objects = from.Objects
	}
	to.EffectiveOwner = to.Owner
	to.Owner = from.Owner
}

func (to *ShareInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ShareInfo_SdkV2) {
	if !from.Objects.IsNull() && !from.Objects.IsUnknown() && to.Objects.IsNull() && len(from.Objects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Objects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Objects = from.Objects
	}
	to.EffectiveOwner = from.EffectiveOwner
	if from.EffectiveOwner.ValueString() == to.Owner.ValueString() {
		to.Owner = from.Owner
	}
}

func (m ShareInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["object"] = attrs["object"].SetOptional()
	attrs["effective_owner"] = attrs["effective_owner"].SetComputed()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["storage_location"] = attrs["storage_location"].SetComputed()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ShareInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ShareInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"object": reflect.TypeOf(SharedDataObject_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ShareInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m ShareInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":    m.Comment,
			"created_at": m.CreatedAt,
			"created_by": m.CreatedBy,
			"name":       m.Name,
			"object":     m.Objects,
			"owner":      m.Owner, "effective_owner": m.EffectiveOwner,
			"storage_location": m.StorageLocation,
			"storage_root":     m.StorageRoot,
			"updated_at":       m.UpdatedAt,
			"updated_by":       m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ShareInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":    types.StringType,
			"created_at": types.Int64Type,
			"created_by": types.StringType,
			"name":       types.StringType,
			"object": basetypes.ListType{
				ElemType: SharedDataObject_SdkV2{}.Type(ctx),
			},
			"owner":            types.StringType,
			"effective_owner":  types.StringType,
			"storage_location": types.StringType,
			"storage_root":     types.StringType,
			"updated_at":       types.Int64Type,
			"updated_by":       types.StringType,
		},
	}
}

// GetObjects returns the value of the Objects field in ShareInfo_SdkV2 as
// a slice of SharedDataObject_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ShareInfo_SdkV2) GetObjects(ctx context.Context) ([]SharedDataObject_SdkV2, bool) {
	if m.Objects.IsNull() || m.Objects.IsUnknown() {
		return nil, false
	}
	var v []SharedDataObject_SdkV2
	d := m.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in ShareInfo_SdkV2.
func (m *ShareInfo_SdkV2) SetObjects(ctx context.Context, v []SharedDataObject_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Objects = types.ListValueMust(t, vs)
}

type SharePermissionsRequest_SdkV2 struct {
	// Maximum number of permissions to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid permissions are returned (not
	// recommended). - Note: The number of returned permissions might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further permissions can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// The name of the Recipient.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *SharePermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SharePermissionsRequest_SdkV2) {
}

func (to *SharePermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SharePermissionsRequest_SdkV2) {
}

func (m SharePermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SharePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SharePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": m.MaxResults,
			"name":        m.Name,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SharePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"name":        types.StringType,
			"page_token":  types.StringType,
		},
	}
}

type ShareToPrivilegeAssignment_SdkV2 struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments"`
	// The share name.
	ShareName types.String `tfsdk:"share_name"`
}

func (to *ShareToPrivilegeAssignment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ShareToPrivilegeAssignment_SdkV2) {
	if !from.PrivilegeAssignments.IsNull() && !from.PrivilegeAssignments.IsUnknown() && to.PrivilegeAssignments.IsNull() && len(from.PrivilegeAssignments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrivilegeAssignments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrivilegeAssignments = from.PrivilegeAssignments
	}
}

func (to *ShareToPrivilegeAssignment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ShareToPrivilegeAssignment_SdkV2) {
	if !from.PrivilegeAssignments.IsNull() && !from.PrivilegeAssignments.IsUnknown() && to.PrivilegeAssignments.IsNull() && len(from.PrivilegeAssignments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrivilegeAssignments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrivilegeAssignments = from.PrivilegeAssignments
	}
}

func (m ShareToPrivilegeAssignment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["privilege_assignments"] = attrs["privilege_assignments"].SetOptional()
	attrs["share_name"] = attrs["share_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ShareToPrivilegeAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ShareToPrivilegeAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(PrivilegeAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ShareToPrivilegeAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (m ShareToPrivilegeAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"privilege_assignments": m.PrivilegeAssignments,
			"share_name":            m.ShareName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ShareToPrivilegeAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment_SdkV2{}.Type(ctx),
			},
			"share_name": types.StringType,
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in ShareToPrivilegeAssignment_SdkV2 as
// a slice of PrivilegeAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ShareToPrivilegeAssignment_SdkV2) GetPrivilegeAssignments(ctx context.Context) ([]PrivilegeAssignment_SdkV2, bool) {
	if m.PrivilegeAssignments.IsNull() || m.PrivilegeAssignments.IsUnknown() {
		return nil, false
	}
	var v []PrivilegeAssignment_SdkV2
	d := m.PrivilegeAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in ShareToPrivilegeAssignment_SdkV2.
func (m *ShareToPrivilegeAssignment_SdkV2) SetPrivilegeAssignments(ctx context.Context, v []PrivilegeAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrivilegeAssignments = types.ListValueMust(t, vs)
}

type SharedDataObject_SdkV2 struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at"`
	// Username of the sharer.
	AddedBy types.String `tfsdk:"added_by"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled          types.Bool `tfsdk:"cdf_enabled"`
	EffectiveCdfEnabled types.Bool `tfsdk:"effective_cdf_enabled"`
	// A user-provided comment when adding the data object to the share.
	Comment types.String `tfsdk:"comment"`
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	Content types.String `tfsdk:"content"`
	// The type of the data object.
	DataObjectType types.String `tfsdk:"data_object_type"`
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	HistoryDataSharingStatus          types.String `tfsdk:"history_data_sharing_status"`
	EffectiveHistoryDataSharingStatus types.String `tfsdk:"effective_history_data_sharing_status"`
	// A fully qualified name that uniquely identifies a data object. For
	// example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`,
	Name types.String `tfsdk:"name"`
	// Array of partitions for the shared data.
	Partitions types.List `tfsdk:"partition"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	SharedAs          types.String `tfsdk:"shared_as"`
	EffectiveSharedAs types.String `tfsdk:"effective_shared_as"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion          types.Int64 `tfsdk:"start_version"`
	EffectiveStartVersion types.Int64 `tfsdk:"effective_start_version"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status types.String `tfsdk:"status"`
	// A user-provided new name for the shared object within the share. If this
	// new name is not not provided, the object's original name will be used as
	// the `string_shared_as` name. The `string_shared_as` name must be unique
	// for objects of the same type within a Share. For notebooks, the new name
	// should be the new notebook file name.
	StringSharedAs          types.String `tfsdk:"string_shared_as"`
	EffectiveStringSharedAs types.String `tfsdk:"effective_string_shared_as"`
}

func (to *SharedDataObject_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SharedDataObject_SdkV2) {
	to.EffectiveCdfEnabled = to.CdfEnabled
	to.CdfEnabled = from.CdfEnabled
	to.EffectiveHistoryDataSharingStatus = to.HistoryDataSharingStatus
	to.HistoryDataSharingStatus = from.HistoryDataSharingStatus
	if !from.Partitions.IsNull() && !from.Partitions.IsUnknown() && to.Partitions.IsNull() && len(from.Partitions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Partitions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Partitions = from.Partitions
	}
	to.EffectiveSharedAs = to.SharedAs
	to.SharedAs = from.SharedAs
	to.EffectiveStartVersion = to.StartVersion
	to.StartVersion = from.StartVersion
	to.EffectiveStringSharedAs = to.StringSharedAs
	to.StringSharedAs = from.StringSharedAs
}

func (to *SharedDataObject_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SharedDataObject_SdkV2) {
	to.EffectiveCdfEnabled = from.EffectiveCdfEnabled
	if from.EffectiveCdfEnabled.ValueBool() == to.CdfEnabled.ValueBool() {
		to.CdfEnabled = from.CdfEnabled
	}
	to.EffectiveHistoryDataSharingStatus = from.EffectiveHistoryDataSharingStatus
	if from.EffectiveHistoryDataSharingStatus.ValueString() == to.HistoryDataSharingStatus.ValueString() {
		to.HistoryDataSharingStatus = from.HistoryDataSharingStatus
	}
	if !from.Partitions.IsNull() && !from.Partitions.IsUnknown() && to.Partitions.IsNull() && len(from.Partitions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Partitions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Partitions = from.Partitions
	}
	to.EffectiveSharedAs = from.EffectiveSharedAs
	if from.EffectiveSharedAs.ValueString() == to.SharedAs.ValueString() {
		to.SharedAs = from.SharedAs
	}
	to.EffectiveStartVersion = from.EffectiveStartVersion
	if from.EffectiveStartVersion.ValueInt64() == to.StartVersion.ValueInt64() {
		to.StartVersion = from.StartVersion
	}
	to.EffectiveStringSharedAs = from.EffectiveStringSharedAs
	if from.EffectiveStringSharedAs.ValueString() == to.StringSharedAs.ValueString() {
		to.StringSharedAs = from.StringSharedAs
	}
}

func (m SharedDataObject_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["added_at"] = attrs["added_at"].SetComputed()
	attrs["added_by"] = attrs["added_by"].SetComputed()
	attrs["effective_cdf_enabled"] = attrs["effective_cdf_enabled"].SetComputed()
	attrs["cdf_enabled"] = attrs["cdf_enabled"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["content"] = attrs["content"].SetOptional()
	attrs["data_object_type"] = attrs["data_object_type"].SetOptional()
	attrs["effective_history_data_sharing_status"] = attrs["effective_history_data_sharing_status"].SetComputed()
	attrs["history_data_sharing_status"] = attrs["history_data_sharing_status"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["partition"] = attrs["partition"].SetOptional()
	attrs["effective_shared_as"] = attrs["effective_shared_as"].SetComputed()
	attrs["shared_as"] = attrs["shared_as"].SetOptional()
	attrs["effective_start_version"] = attrs["effective_start_version"].SetComputed()
	attrs["start_version"] = attrs["start_version"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["effective_string_shared_as"] = attrs["effective_string_shared_as"].SetComputed()
	attrs["string_shared_as"] = attrs["string_shared_as"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharedDataObject.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SharedDataObject_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partition": reflect.TypeOf(Partition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharedDataObject_SdkV2
// only implements ToObjectValue() and Type().
func (m SharedDataObject_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"added_at":    m.AddedAt,
			"added_by":    m.AddedBy,
			"cdf_enabled": m.CdfEnabled, "effective_cdf_enabled": m.EffectiveCdfEnabled,
			"comment":                     m.Comment,
			"content":                     m.Content,
			"data_object_type":            m.DataObjectType,
			"history_data_sharing_status": m.HistoryDataSharingStatus, "effective_history_data_sharing_status": m.EffectiveHistoryDataSharingStatus,
			"name":      m.Name,
			"partition": m.Partitions,
			"shared_as": m.SharedAs, "effective_shared_as": m.EffectiveSharedAs,
			"start_version": m.StartVersion, "effective_start_version": m.EffectiveStartVersion,
			"status":           m.Status,
			"string_shared_as": m.StringSharedAs, "effective_string_shared_as": m.EffectiveStringSharedAs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SharedDataObject_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"added_at":                              types.Int64Type,
			"added_by":                              types.StringType,
			"cdf_enabled":                           types.BoolType,
			"effective_cdf_enabled":                 types.BoolType,
			"comment":                               types.StringType,
			"content":                               types.StringType,
			"data_object_type":                      types.StringType,
			"history_data_sharing_status":           types.StringType,
			"effective_history_data_sharing_status": types.StringType,
			"name":                                  types.StringType,
			"partition": basetypes.ListType{
				ElemType: Partition_SdkV2{}.Type(ctx),
			},
			"shared_as":                  types.StringType,
			"effective_shared_as":        types.StringType,
			"start_version":              types.Int64Type,
			"effective_start_version":    types.Int64Type,
			"status":                     types.StringType,
			"string_shared_as":           types.StringType,
			"effective_string_shared_as": types.StringType,
		},
	}
}

// GetPartitions returns the value of the Partitions field in SharedDataObject_SdkV2 as
// a slice of Partition_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SharedDataObject_SdkV2) GetPartitions(ctx context.Context) ([]Partition_SdkV2, bool) {
	if m.Partitions.IsNull() || m.Partitions.IsUnknown() {
		return nil, false
	}
	var v []Partition_SdkV2
	d := m.Partitions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPartitions sets the value of the Partitions field in SharedDataObject_SdkV2.
func (m *SharedDataObject_SdkV2) SetPartitions(ctx context.Context, v []Partition_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["partition"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Partitions = types.ListValueMust(t, vs)
}

type SharedDataObjectUpdate_SdkV2 struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	Action types.String `tfsdk:"action"`
	// The data object that is being added, removed, or updated. The maximum
	// number update data objects allowed is a 100.
	DataObject types.List `tfsdk:"data_object"`
}

func (to *SharedDataObjectUpdate_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SharedDataObjectUpdate_SdkV2) {
	if !from.DataObject.IsNull() && !from.DataObject.IsUnknown() {
		if toDataObject, ok := to.GetDataObject(ctx); ok {
			if fromDataObject, ok := from.GetDataObject(ctx); ok {
				// Recursively sync the fields of DataObject
				toDataObject.SyncFieldsDuringCreateOrUpdate(ctx, fromDataObject)
				to.SetDataObject(ctx, toDataObject)
			}
		}
	}
}

func (to *SharedDataObjectUpdate_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SharedDataObjectUpdate_SdkV2) {
	if !from.DataObject.IsNull() && !from.DataObject.IsUnknown() {
		if toDataObject, ok := to.GetDataObject(ctx); ok {
			if fromDataObject, ok := from.GetDataObject(ctx); ok {
				toDataObject.SyncFieldsDuringRead(ctx, fromDataObject)
				to.SetDataObject(ctx, toDataObject)
			}
		}
	}
}

func (m SharedDataObjectUpdate_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["action"] = attrs["action"].SetOptional()
	attrs["data_object"] = attrs["data_object"].SetOptional()
	attrs["data_object"] = attrs["data_object"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharedDataObjectUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SharedDataObjectUpdate_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_object": reflect.TypeOf(SharedDataObject_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharedDataObjectUpdate_SdkV2
// only implements ToObjectValue() and Type().
func (m SharedDataObjectUpdate_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action":      m.Action,
			"data_object": m.DataObject,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SharedDataObjectUpdate_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action": types.StringType,
			"data_object": basetypes.ListType{
				ElemType: SharedDataObject_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDataObject returns the value of the DataObject field in SharedDataObjectUpdate_SdkV2 as
// a SharedDataObject_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SharedDataObjectUpdate_SdkV2) GetDataObject(ctx context.Context) (SharedDataObject_SdkV2, bool) {
	var e SharedDataObject_SdkV2
	if m.DataObject.IsNull() || m.DataObject.IsUnknown() {
		return e, false
	}
	var v []SharedDataObject_SdkV2
	d := m.DataObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataObject sets the value of the DataObject field in SharedDataObjectUpdate_SdkV2.
func (m *SharedDataObjectUpdate_SdkV2) SetDataObject(ctx context.Context, v SharedDataObject_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_object"]
	m.DataObject = types.ListValueMust(t, vs)
}

type SharesListRequest_SdkV2 struct {
	// Maximum number of shares to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid shares are returned (not
	// recommended). - Note: The number of returned shares might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further shares can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *SharesListRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SharesListRequest_SdkV2) {
}

func (to *SharesListRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SharesListRequest_SdkV2) {
}

func (m SharesListRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharesListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SharesListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharesListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SharesListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": m.MaxResults,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SharesListRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type Table_SdkV2 struct {
	// The comment of the table.
	Comment types.String `tfsdk:"comment"`
	// The id of the table.
	Id types.String `tfsdk:"id"`
	// Internal information for D2D sharing that should not be disclosed to
	// external users.
	InternalAttributes types.List `tfsdk:"internal_attributes"`
	// The catalog and schema of the materialized table
	MaterializationNamespace types.String `tfsdk:"materialization_namespace"`
	// The name of a materialized table.
	MaterializedTableName types.String `tfsdk:"materialized_table_name"`
	// The name of the table.
	Name types.String `tfsdk:"name"`
	// The name of the schema that the table belongs to.
	Schema types.String `tfsdk:"schema"`
	// The name of the share that the table belongs to.
	Share types.String `tfsdk:"share"`
	// The id of the share that the table belongs to.
	ShareId types.String `tfsdk:"share_id"`
	// The Tags of the table.
	Tags types.List `tfsdk:"tags"`
}

func (to *Table_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Table_SdkV2) {
	if !from.InternalAttributes.IsNull() && !from.InternalAttributes.IsUnknown() {
		if toInternalAttributes, ok := to.GetInternalAttributes(ctx); ok {
			if fromInternalAttributes, ok := from.GetInternalAttributes(ctx); ok {
				// Recursively sync the fields of InternalAttributes
				toInternalAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromInternalAttributes)
				to.SetInternalAttributes(ctx, toInternalAttributes)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *Table_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Table_SdkV2) {
	if !from.InternalAttributes.IsNull() && !from.InternalAttributes.IsUnknown() {
		if toInternalAttributes, ok := to.GetInternalAttributes(ctx); ok {
			if fromInternalAttributes, ok := from.GetInternalAttributes(ctx); ok {
				toInternalAttributes.SyncFieldsDuringRead(ctx, fromInternalAttributes)
				to.SetInternalAttributes(ctx, toInternalAttributes)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m Table_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["internal_attributes"] = attrs["internal_attributes"].SetOptional()
	attrs["internal_attributes"] = attrs["internal_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["materialization_namespace"] = attrs["materialization_namespace"].SetOptional()
	attrs["materialized_table_name"] = attrs["materialized_table_name"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["share"] = attrs["share"].SetOptional()
	attrs["share_id"] = attrs["share_id"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Table.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Table_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"internal_attributes": reflect.TypeOf(TableInternalAttributes_SdkV2{}),
		"tags":                reflect.TypeOf(catalog_tf.TagKeyValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Table_SdkV2
// only implements ToObjectValue() and Type().
func (m Table_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":                   m.Comment,
			"id":                        m.Id,
			"internal_attributes":       m.InternalAttributes,
			"materialization_namespace": m.MaterializationNamespace,
			"materialized_table_name":   m.MaterializedTableName,
			"name":                      m.Name,
			"schema":                    m.Schema,
			"share":                     m.Share,
			"share_id":                  m.ShareId,
			"tags":                      m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Table_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"id":      types.StringType,
			"internal_attributes": basetypes.ListType{
				ElemType: TableInternalAttributes_SdkV2{}.Type(ctx),
			},
			"materialization_namespace": types.StringType,
			"materialized_table_name":   types.StringType,
			"name":                      types.StringType,
			"schema":                    types.StringType,
			"share":                     types.StringType,
			"share_id":                  types.StringType,
			"tags": basetypes.ListType{
				ElemType: catalog_tf.TagKeyValue_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetInternalAttributes returns the value of the InternalAttributes field in Table_SdkV2 as
// a TableInternalAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Table_SdkV2) GetInternalAttributes(ctx context.Context) (TableInternalAttributes_SdkV2, bool) {
	var e TableInternalAttributes_SdkV2
	if m.InternalAttributes.IsNull() || m.InternalAttributes.IsUnknown() {
		return e, false
	}
	var v []TableInternalAttributes_SdkV2
	d := m.InternalAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInternalAttributes sets the value of the InternalAttributes field in Table_SdkV2.
func (m *Table_SdkV2) SetInternalAttributes(ctx context.Context, v TableInternalAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["internal_attributes"]
	m.InternalAttributes = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Table_SdkV2 as
// a slice of catalog_tf.TagKeyValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Table_SdkV2) GetTags(ctx context.Context) ([]catalog_tf.TagKeyValue_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.TagKeyValue_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Table_SdkV2.
func (m *Table_SdkV2) SetTags(ctx context.Context, v []catalog_tf.TagKeyValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// Internal information for D2D sharing that should not be disclosed to external
// users.
type TableInternalAttributes_SdkV2 struct {
	// Managed Delta Metadata location for foreign iceberg tables.
	AuxiliaryManagedLocation types.String `tfsdk:"auxiliary_managed_location"`
	// Will be populated in the reconciliation response for VIEW and
	// FOREIGN_TABLE, with the value of the parent UC entity's storage_location,
	// following the same logic as getManagedEntityPath in
	// CreateStagingTableHandler, which is used to store the materialized table
	// for a shared VIEW/FOREIGN_TABLE for D2O queries. The value will be used
	// on the recipient side to be whitelisted when SEG is enabled on the
	// workspace of the recipient, to allow the recipient users to query this
	// shared VIEW/FOREIGN_TABLE.
	ParentStorageLocation types.String `tfsdk:"parent_storage_location"`
	// The cloud storage location of a shard table with DIRECTORY_BASED_TABLE
	// type.
	StorageLocation types.String `tfsdk:"storage_location"`
	// The type of the shared table.
	Type_ types.String `tfsdk:"type"`
	// The view definition of a shared view. DEPRECATED.
	ViewDefinition types.String `tfsdk:"view_definition"`
}

func (to *TableInternalAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TableInternalAttributes_SdkV2) {
}

func (to *TableInternalAttributes_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TableInternalAttributes_SdkV2) {
}

func (m TableInternalAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auxiliary_managed_location"] = attrs["auxiliary_managed_location"].SetOptional()
	attrs["parent_storage_location"] = attrs["parent_storage_location"].SetOptional()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()
	attrs["view_definition"] = attrs["view_definition"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableInternalAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TableInternalAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableInternalAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (m TableInternalAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auxiliary_managed_location": m.AuxiliaryManagedLocation,
			"parent_storage_location":    m.ParentStorageLocation,
			"storage_location":           m.StorageLocation,
			"type":                       m.Type_,
			"view_definition":            m.ViewDefinition,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TableInternalAttributes_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auxiliary_managed_location": types.StringType,
			"parent_storage_location":    types.StringType,
			"storage_location":           types.StringType,
			"type":                       types.StringType,
			"view_definition":            types.StringType,
		},
	}
}

type UpdateFederationPolicyRequest_SdkV2 struct {
	// Name of the policy. This is the name of the current name of the policy.
	Name types.String `tfsdk:"-"`

	Policy types.List `tfsdk:"policy"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being updated.
	RecipientName types.String `tfsdk:"-"`
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'comment,oidc_policy.audiences'.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateFederationPolicyRequest_SdkV2) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				// Recursively sync the fields of Policy
				toPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (to *UpdateFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateFederationPolicyRequest_SdkV2) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m UpdateFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["policy"] = attrs["policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["recipient_name"] = attrs["recipient_name"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":           m.Name,
			"policy":         m.Policy,
			"recipient_name": m.RecipientName,
			"update_mask":    m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"policy": basetypes.ListType{
				ElemType: FederationPolicy_SdkV2{}.Type(ctx),
			},
			"recipient_name": types.StringType,
			"update_mask":    types.StringType,
		},
	}
}

// GetPolicy returns the value of the Policy field in UpdateFederationPolicyRequest_SdkV2 as
// a FederationPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateFederationPolicyRequest_SdkV2) GetPolicy(ctx context.Context) (FederationPolicy_SdkV2, bool) {
	var e FederationPolicy_SdkV2
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v []FederationPolicy_SdkV2
	d := m.Policy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPolicy sets the value of the Policy field in UpdateFederationPolicyRequest_SdkV2.
func (m *UpdateFederationPolicyRequest_SdkV2) SetPolicy(ctx context.Context, v FederationPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policy"]
	m.Policy = types.ListValueMust(t, vs)
}

type UpdateProvider_SdkV2 struct {
	// Description about the provider.
	Comment types.String `tfsdk:"comment"`
	// Name of the provider.
	Name types.String `tfsdk:"-"`
	// New name for the provider.
	NewName types.String `tfsdk:"new_name"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str"`
}

func (to *UpdateProvider_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProvider_SdkV2) {
}

func (to *UpdateProvider_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateProvider_SdkV2) {
}

func (m UpdateProvider_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["recipient_profile_str"] = attrs["recipient_profile_str"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProvider.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateProvider_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProvider_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateProvider_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":               m.Comment,
			"name":                  m.Name,
			"new_name":              m.NewName,
			"owner":                 m.Owner,
			"recipient_profile_str": m.RecipientProfileStr,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProvider_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":               types.StringType,
			"name":                  types.StringType,
			"new_name":              types.StringType,
			"owner":                 types.StringType,
			"recipient_profile_str": types.StringType,
		},
	}
}

type UpdateRecipient_SdkV2 struct {
	// Description about the recipient.
	Comment types.String `tfsdk:"comment"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// IP Access List
	IpAccessList types.List `tfsdk:"ip_access_list"`
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
	// New name for the recipient. .
	NewName types.String `tfsdk:"new_name"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs types.List `tfsdk:"properties_kvpairs"`
}

func (to *UpdateRecipient_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRecipient_SdkV2) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				// Recursively sync the fields of IpAccessList
				toIpAccessList.SyncFieldsDuringCreateOrUpdate(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
	if !from.PropertiesKvpairs.IsNull() && !from.PropertiesKvpairs.IsUnknown() {
		if toPropertiesKvpairs, ok := to.GetPropertiesKvpairs(ctx); ok {
			if fromPropertiesKvpairs, ok := from.GetPropertiesKvpairs(ctx); ok {
				// Recursively sync the fields of PropertiesKvpairs
				toPropertiesKvpairs.SyncFieldsDuringCreateOrUpdate(ctx, fromPropertiesKvpairs)
				to.SetPropertiesKvpairs(ctx, toPropertiesKvpairs)
			}
		}
	}
}

func (to *UpdateRecipient_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRecipient_SdkV2) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				toIpAccessList.SyncFieldsDuringRead(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
	if !from.PropertiesKvpairs.IsNull() && !from.PropertiesKvpairs.IsUnknown() {
		if toPropertiesKvpairs, ok := to.GetPropertiesKvpairs(ctx); ok {
			if fromPropertiesKvpairs, ok := from.GetPropertiesKvpairs(ctx); ok {
				toPropertiesKvpairs.SyncFieldsDuringRead(ctx, fromPropertiesKvpairs)
				to.SetPropertiesKvpairs(ctx, toPropertiesKvpairs)
			}
		}
	}
}

func (m UpdateRecipient_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties_kvpairs"] = attrs["properties_kvpairs"].SetOptional()
	attrs["properties_kvpairs"] = attrs["properties_kvpairs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRecipient.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRecipient_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList_SdkV2{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRecipient_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateRecipient_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":            m.Comment,
			"expiration_time":    m.ExpirationTime,
			"ip_access_list":     m.IpAccessList,
			"name":               m.Name,
			"new_name":           m.NewName,
			"owner":              m.Owner,
			"properties_kvpairs": m.PropertiesKvpairs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRecipient_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":         types.StringType,
			"expiration_time": types.Int64Type,
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessList_SdkV2{}.Type(ctx),
			},
			"name":     types.StringType,
			"new_name": types.StringType,
			"owner":    types.StringType,
			"properties_kvpairs": basetypes.ListType{
				ElemType: SecurablePropertiesKvPairs_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in UpdateRecipient_SdkV2 as
// a IpAccessList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRecipient_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessList_SdkV2, bool) {
	var e IpAccessList_SdkV2
	if m.IpAccessList.IsNull() || m.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessList_SdkV2
	d := m.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in UpdateRecipient_SdkV2.
func (m *UpdateRecipient_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	m.IpAccessList = types.ListValueMust(t, vs)
}

// GetPropertiesKvpairs returns the value of the PropertiesKvpairs field in UpdateRecipient_SdkV2 as
// a SecurablePropertiesKvPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRecipient_SdkV2) GetPropertiesKvpairs(ctx context.Context) (SecurablePropertiesKvPairs_SdkV2, bool) {
	var e SecurablePropertiesKvPairs_SdkV2
	if m.PropertiesKvpairs.IsNull() || m.PropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []SecurablePropertiesKvPairs_SdkV2
	d := m.PropertiesKvpairs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPropertiesKvpairs sets the value of the PropertiesKvpairs field in UpdateRecipient_SdkV2.
func (m *UpdateRecipient_SdkV2) SetPropertiesKvpairs(ctx context.Context, v SecurablePropertiesKvPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["properties_kvpairs"]
	m.PropertiesKvpairs = types.ListValueMust(t, vs)
}

type UpdateShare_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// New name for the share.
	NewName types.String `tfsdk:"new_name"`
	// Username of current owner of share.
	Owner          types.String `tfsdk:"owner"`
	EffectiveOwner types.String `tfsdk:"effective_owner"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root"`
	// Array of shared data object updates.
	Updates types.List `tfsdk:"updates"`
}

func (to *UpdateShare_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateShare_SdkV2) {
	to.EffectiveOwner = to.Owner
	to.Owner = from.Owner
	if !from.Updates.IsNull() && !from.Updates.IsUnknown() && to.Updates.IsNull() && len(from.Updates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Updates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Updates = from.Updates
	}
}

func (to *UpdateShare_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateShare_SdkV2) {
	to.EffectiveOwner = from.EffectiveOwner
	if from.EffectiveOwner.ValueString() == to.Owner.ValueString() {
		to.Owner = from.Owner
	}
	if !from.Updates.IsNull() && !from.Updates.IsUnknown() && to.Updates.IsNull() && len(from.Updates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Updates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Updates = from.Updates
	}
}

func (m UpdateShare_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["effective_owner"] = attrs["effective_owner"].SetComputed()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()
	attrs["updates"] = attrs["updates"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateShare.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateShare_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"updates": reflect.TypeOf(SharedDataObjectUpdate_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateShare_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateShare_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":  m.Comment,
			"name":     m.Name,
			"new_name": m.NewName,
			"owner":    m.Owner, "effective_owner": m.EffectiveOwner,
			"storage_root": m.StorageRoot,
			"updates":      m.Updates,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateShare_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":         types.StringType,
			"name":            types.StringType,
			"new_name":        types.StringType,
			"owner":           types.StringType,
			"effective_owner": types.StringType,
			"storage_root":    types.StringType,
			"updates": basetypes.ListType{
				ElemType: SharedDataObjectUpdate_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetUpdates returns the value of the Updates field in UpdateShare_SdkV2 as
// a slice of SharedDataObjectUpdate_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateShare_SdkV2) GetUpdates(ctx context.Context) ([]SharedDataObjectUpdate_SdkV2, bool) {
	if m.Updates.IsNull() || m.Updates.IsUnknown() {
		return nil, false
	}
	var v []SharedDataObjectUpdate_SdkV2
	d := m.Updates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdates sets the value of the Updates field in UpdateShare_SdkV2.
func (m *UpdateShare_SdkV2) SetUpdates(ctx context.Context, v []SharedDataObjectUpdate_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Updates = types.ListValueMust(t, vs)
}

type UpdateSharePermissions_SdkV2 struct {
	// Array of permissions change objects.
	Changes types.List `tfsdk:"changes"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// Optional. Whether to return the latest permissions list of the share in
	// the response.
	OmitPermissionsList types.Bool `tfsdk:"omit_permissions_list"`
}

func (to *UpdateSharePermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSharePermissions_SdkV2) {
	if !from.Changes.IsNull() && !from.Changes.IsUnknown() && to.Changes.IsNull() && len(from.Changes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Changes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Changes = from.Changes
	}
}

func (to *UpdateSharePermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateSharePermissions_SdkV2) {
	if !from.Changes.IsNull() && !from.Changes.IsUnknown() && to.Changes.IsNull() && len(from.Changes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Changes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Changes = from.Changes
	}
}

func (m UpdateSharePermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["changes"] = attrs["changes"].SetOptional()
	attrs["omit_permissions_list"] = attrs["omit_permissions_list"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSharePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSharePermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"changes": reflect.TypeOf(PermissionsChange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSharePermissions_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateSharePermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"changes":               m.Changes,
			"name":                  m.Name,
			"omit_permissions_list": m.OmitPermissionsList,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSharePermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"changes": basetypes.ListType{
				ElemType: PermissionsChange_SdkV2{}.Type(ctx),
			},
			"name":                  types.StringType,
			"omit_permissions_list": types.BoolType,
		},
	}
}

// GetChanges returns the value of the Changes field in UpdateSharePermissions_SdkV2 as
// a slice of PermissionsChange_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSharePermissions_SdkV2) GetChanges(ctx context.Context) ([]PermissionsChange_SdkV2, bool) {
	if m.Changes.IsNull() || m.Changes.IsUnknown() {
		return nil, false
	}
	var v []PermissionsChange_SdkV2
	d := m.Changes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChanges sets the value of the Changes field in UpdateSharePermissions_SdkV2.
func (m *UpdateSharePermissions_SdkV2) SetChanges(ctx context.Context, v []PermissionsChange_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["changes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Changes = types.ListValueMust(t, vs)
}

type UpdateSharePermissionsResponse_SdkV2 struct {
	// The privileges assigned to each principal
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments"`
}

func (to *UpdateSharePermissionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSharePermissionsResponse_SdkV2) {
	if !from.PrivilegeAssignments.IsNull() && !from.PrivilegeAssignments.IsUnknown() && to.PrivilegeAssignments.IsNull() && len(from.PrivilegeAssignments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrivilegeAssignments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrivilegeAssignments = from.PrivilegeAssignments
	}
}

func (to *UpdateSharePermissionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateSharePermissionsResponse_SdkV2) {
	if !from.PrivilegeAssignments.IsNull() && !from.PrivilegeAssignments.IsUnknown() && to.PrivilegeAssignments.IsNull() && len(from.PrivilegeAssignments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrivilegeAssignments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrivilegeAssignments = from.PrivilegeAssignments
	}
}

func (m UpdateSharePermissionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["privilege_assignments"] = attrs["privilege_assignments"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSharePermissionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSharePermissionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(PrivilegeAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSharePermissionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateSharePermissionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"privilege_assignments": m.PrivilegeAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSharePermissionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in UpdateSharePermissionsResponse_SdkV2 as
// a slice of PrivilegeAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSharePermissionsResponse_SdkV2) GetPrivilegeAssignments(ctx context.Context) ([]PrivilegeAssignment_SdkV2, bool) {
	if m.PrivilegeAssignments.IsNull() || m.PrivilegeAssignments.IsUnknown() {
		return nil, false
	}
	var v []PrivilegeAssignment_SdkV2
	d := m.PrivilegeAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in UpdateSharePermissionsResponse_SdkV2.
func (m *UpdateSharePermissionsResponse_SdkV2) SetPrivilegeAssignments(ctx context.Context, v []PrivilegeAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrivilegeAssignments = types.ListValueMust(t, vs)
}

type Volume_SdkV2 struct {
	// The comment of the volume.
	Comment types.String `tfsdk:"comment"`
	// This id maps to the shared_volume_id in database Recipient needs
	// shared_volume_id for recon to check if this volume is already in
	// recipient's DB or not.
	Id types.String `tfsdk:"id"`
	// Internal attributes for D2D sharing that should not be disclosed to
	// external users.
	InternalAttributes types.List `tfsdk:"internal_attributes"`
	// The name of the volume.
	Name types.String `tfsdk:"name"`
	// The name of the schema that the volume belongs to.
	Schema types.String `tfsdk:"schema"`
	// The name of the share that the volume belongs to.
	Share types.String `tfsdk:"share"`
	// / The id of the share that the volume belongs to.
	ShareId types.String `tfsdk:"share_id"`
	// The tags of the volume.
	Tags types.List `tfsdk:"tags"`
}

func (to *Volume_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Volume_SdkV2) {
	if !from.InternalAttributes.IsNull() && !from.InternalAttributes.IsUnknown() {
		if toInternalAttributes, ok := to.GetInternalAttributes(ctx); ok {
			if fromInternalAttributes, ok := from.GetInternalAttributes(ctx); ok {
				// Recursively sync the fields of InternalAttributes
				toInternalAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromInternalAttributes)
				to.SetInternalAttributes(ctx, toInternalAttributes)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *Volume_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Volume_SdkV2) {
	if !from.InternalAttributes.IsNull() && !from.InternalAttributes.IsUnknown() {
		if toInternalAttributes, ok := to.GetInternalAttributes(ctx); ok {
			if fromInternalAttributes, ok := from.GetInternalAttributes(ctx); ok {
				toInternalAttributes.SyncFieldsDuringRead(ctx, fromInternalAttributes)
				to.SetInternalAttributes(ctx, toInternalAttributes)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m Volume_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["internal_attributes"] = attrs["internal_attributes"].SetOptional()
	attrs["internal_attributes"] = attrs["internal_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["share"] = attrs["share"].SetOptional()
	attrs["share_id"] = attrs["share_id"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Volume.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Volume_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"internal_attributes": reflect.TypeOf(VolumeInternalAttributes_SdkV2{}),
		"tags":                reflect.TypeOf(catalog_tf.TagKeyValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Volume_SdkV2
// only implements ToObjectValue() and Type().
func (m Volume_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":             m.Comment,
			"id":                  m.Id,
			"internal_attributes": m.InternalAttributes,
			"name":                m.Name,
			"schema":              m.Schema,
			"share":               m.Share,
			"share_id":            m.ShareId,
			"tags":                m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Volume_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"id":      types.StringType,
			"internal_attributes": basetypes.ListType{
				ElemType: VolumeInternalAttributes_SdkV2{}.Type(ctx),
			},
			"name":     types.StringType,
			"schema":   types.StringType,
			"share":    types.StringType,
			"share_id": types.StringType,
			"tags": basetypes.ListType{
				ElemType: catalog_tf.TagKeyValue_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetInternalAttributes returns the value of the InternalAttributes field in Volume_SdkV2 as
// a VolumeInternalAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Volume_SdkV2) GetInternalAttributes(ctx context.Context) (VolumeInternalAttributes_SdkV2, bool) {
	var e VolumeInternalAttributes_SdkV2
	if m.InternalAttributes.IsNull() || m.InternalAttributes.IsUnknown() {
		return e, false
	}
	var v []VolumeInternalAttributes_SdkV2
	d := m.InternalAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInternalAttributes sets the value of the InternalAttributes field in Volume_SdkV2.
func (m *Volume_SdkV2) SetInternalAttributes(ctx context.Context, v VolumeInternalAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["internal_attributes"]
	m.InternalAttributes = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Volume_SdkV2 as
// a slice of catalog_tf.TagKeyValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Volume_SdkV2) GetTags(ctx context.Context) ([]catalog_tf.TagKeyValue_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.TagKeyValue_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Volume_SdkV2.
func (m *Volume_SdkV2) SetTags(ctx context.Context, v []catalog_tf.TagKeyValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// Internal information for D2D sharing that should not be disclosed to external
// users.
type VolumeInternalAttributes_SdkV2 struct {
	// The cloud storage location of the volume
	StorageLocation types.String `tfsdk:"storage_location"`
	// The type of the shared volume.
	Type_ types.String `tfsdk:"type"`
}

func (to *VolumeInternalAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VolumeInternalAttributes_SdkV2) {
}

func (to *VolumeInternalAttributes_SdkV2) SyncFieldsDuringRead(ctx context.Context, from VolumeInternalAttributes_SdkV2) {
}

func (m VolumeInternalAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VolumeInternalAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VolumeInternalAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VolumeInternalAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (m VolumeInternalAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_location": m.StorageLocation,
			"type":             m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VolumeInternalAttributes_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_location": types.StringType,
			"type":             types.StringType,
		},
	}
}
