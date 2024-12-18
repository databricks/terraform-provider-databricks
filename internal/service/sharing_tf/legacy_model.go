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

	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateProvider_SdkV2 struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:""`
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the Provider.
	Name types.String `tfsdk:"name" tf:""`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
}

func (newState *CreateProvider_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateProvider_SdkV2) {
}

func (newState *CreateProvider_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateProvider_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProvider.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateProvider_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProvider_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateProvider_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authentication_type":   o.AuthenticationType,
			"comment":               o.Comment,
			"name":                  o.Name,
			"recipient_profile_str": o.RecipientProfileStr,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateProvider_SdkV2) Type(ctx context.Context) attr.Type {
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
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:""`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is required when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id" tf:"optional"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// IP Access List
	IpAccessList types.List `tfsdk:"ip_access_list" tf:"optional,object"`
	// Name of Recipient.
	Name types.String `tfsdk:"name" tf:""`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs types.List `tfsdk:"properties_kvpairs" tf:"optional,object"`
	// The one-time sharing code provided by the data recipient. This field is
	// required when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code" tf:"optional"`
}

func (newState *CreateRecipient_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRecipient_SdkV2) {
}

func (newState *CreateRecipient_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateRecipient_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRecipient.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateRecipient_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList_SdkV2{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRecipient_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateRecipient_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authentication_type":                o.AuthenticationType,
			"comment":                            o.Comment,
			"data_recipient_global_metastore_id": o.DataRecipientGlobalMetastoreId,
			"expiration_time":                    o.ExpirationTime,
			"ip_access_list":                     o.IpAccessList,
			"name":                               o.Name,
			"owner":                              o.Owner,
			"properties_kvpairs":                 o.PropertiesKvpairs,
			"sharing_code":                       o.SharingCode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateRecipient_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authentication_type":                types.StringType,
			"comment":                            types.StringType,
			"data_recipient_global_metastore_id": types.StringType,
			"expiration_time":                    types.Int64Type,
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessList{}.Type(ctx),
			},
			"name":  types.StringType,
			"owner": types.StringType,
			"properties_kvpairs": basetypes.ListType{
				ElemType: SecurablePropertiesKvPairs{}.Type(ctx),
			},
			"sharing_code": types.StringType,
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in CreateRecipient_SdkV2 as
// a IpAccessList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRecipient_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessList_SdkV2, bool) {
	var e IpAccessList_SdkV2
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessList_SdkV2
	d := o.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in CreateRecipient_SdkV2.
func (o *CreateRecipient_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	o.IpAccessList = types.ListValueMust(t, vs)
}

// GetPropertiesKvpairs returns the value of the PropertiesKvpairs field in CreateRecipient_SdkV2 as
// a SecurablePropertiesKvPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRecipient_SdkV2) GetPropertiesKvpairs(ctx context.Context) (SecurablePropertiesKvPairs_SdkV2, bool) {
	var e SecurablePropertiesKvPairs_SdkV2
	if o.PropertiesKvpairs.IsNull() || o.PropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []SecurablePropertiesKvPairs_SdkV2
	d := o.PropertiesKvpairs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPropertiesKvpairs sets the value of the PropertiesKvpairs field in CreateRecipient_SdkV2.
func (o *CreateRecipient_SdkV2) SetPropertiesKvpairs(ctx context.Context, v SecurablePropertiesKvPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties_kvpairs"]
	o.PropertiesKvpairs = types.ListValueMust(t, vs)
}

type CreateShare_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the share.
	Name types.String `tfsdk:"name" tf:""`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

func (newState *CreateShare_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateShare_SdkV2) {
}

func (newState *CreateShare_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateShare_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateShare.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateShare_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateShare_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateShare_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":      o.Comment,
			"name":         o.Name,
			"storage_root": o.StorageRoot,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateShare_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":      types.StringType,
			"name":         types.StringType,
			"storage_root": types.StringType,
		},
	}
}

// Delete a provider
type DeleteProviderRequest_SdkV2 struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteProviderRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteProviderRequest_SdkV2) {
}

func (newState *DeleteProviderRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteProviderRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteProviderRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProviderRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteProviderRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteProviderRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Delete a share recipient
type DeleteRecipientRequest_SdkV2 struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteRecipientRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRecipientRequest_SdkV2) {
}

func (newState *DeleteRecipientRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteRecipientRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRecipientRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRecipientRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRecipientRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRecipientRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRecipientRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteResponse_SdkV2 struct {
}

func (newState *DeleteResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse_SdkV2) {
}

func (newState *DeleteResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a share
type DeleteShareRequest_SdkV2 struct {
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteShareRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteShareRequest_SdkV2) {
}

func (newState *DeleteShareRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteShareRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteShareRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteShareRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteShareRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteShareRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteShareRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a share activation URL
type GetActivationUrlInfoRequest_SdkV2 struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

func (newState *GetActivationUrlInfoRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetActivationUrlInfoRequest_SdkV2) {
}

func (newState *GetActivationUrlInfoRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetActivationUrlInfoRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetActivationUrlInfoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetActivationUrlInfoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetActivationUrlInfoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetActivationUrlInfoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activation_url": o.ActivationUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activation_url": types.StringType,
		},
	}
}

type GetActivationUrlInfoResponse_SdkV2 struct {
}

func (newState *GetActivationUrlInfoResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetActivationUrlInfoResponse_SdkV2) {
}

func (newState *GetActivationUrlInfoResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetActivationUrlInfoResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetActivationUrlInfoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetActivationUrlInfoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetActivationUrlInfoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetActivationUrlInfoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Get a provider
type GetProviderRequest_SdkV2 struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
}

func (newState *GetProviderRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetProviderRequest_SdkV2) {
}

func (newState *GetProviderRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetProviderRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetProviderRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProviderRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetProviderRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetProviderRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a share recipient
type GetRecipientRequest_SdkV2 struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *GetRecipientRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRecipientRequest_SdkV2) {
}

func (newState *GetRecipientRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetRecipientRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRecipientRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRecipientRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRecipientRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRecipientRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRecipientRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of data share permissions for a recipient.
	PermissionsOut types.List `tfsdk:"permissions_out" tf:"optional"`
}

func (newState *GetRecipientSharePermissionsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRecipientSharePermissionsResponse_SdkV2) {
}

func (newState *GetRecipientSharePermissionsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetRecipientSharePermissionsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRecipientSharePermissionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRecipientSharePermissionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions_out": reflect.TypeOf(ShareToPrivilegeAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRecipientSharePermissionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRecipientSharePermissionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"permissions_out": o.PermissionsOut,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRecipientSharePermissionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"permissions_out": basetypes.ListType{
				ElemType: ShareToPrivilegeAssignment{}.Type(ctx),
			},
		},
	}
}

// GetPermissionsOut returns the value of the PermissionsOut field in GetRecipientSharePermissionsResponse_SdkV2 as
// a slice of ShareToPrivilegeAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetRecipientSharePermissionsResponse_SdkV2) GetPermissionsOut(ctx context.Context) ([]ShareToPrivilegeAssignment_SdkV2, bool) {
	if o.PermissionsOut.IsNull() || o.PermissionsOut.IsUnknown() {
		return nil, false
	}
	var v []ShareToPrivilegeAssignment_SdkV2
	d := o.PermissionsOut.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionsOut sets the value of the PermissionsOut field in GetRecipientSharePermissionsResponse_SdkV2.
func (o *GetRecipientSharePermissionsResponse_SdkV2) SetPermissionsOut(ctx context.Context, v []ShareToPrivilegeAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions_out"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionsOut = types.ListValueMust(t, vs)
}

// Get a share
type GetShareRequest_SdkV2 struct {
	// Query for data to include in the share.
	IncludeSharedData types.Bool `tfsdk:"-"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (newState *GetShareRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetShareRequest_SdkV2) {
}

func (newState *GetShareRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetShareRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetShareRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetShareRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetShareRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetShareRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_shared_data": o.IncludeSharedData,
			"name":                o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetShareRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_shared_data": types.BoolType,
			"name":                types.StringType,
		},
	}
}

type IpAccessList_SdkV2 struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses types.List `tfsdk:"allowed_ip_addresses" tf:"optional"`
}

func (newState *IpAccessList_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan IpAccessList_SdkV2) {
}

func (newState *IpAccessList_SdkV2) SyncEffectiveFieldsDuringRead(existingState IpAccessList_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IpAccessList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IpAccessList_SdkV2
// only implements ToObjectValue() and Type().
func (o IpAccessList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_ip_addresses": o.AllowedIpAddresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IpAccessList_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *IpAccessList_SdkV2) GetAllowedIpAddresses(ctx context.Context) ([]types.String, bool) {
	if o.AllowedIpAddresses.IsNull() || o.AllowedIpAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.AllowedIpAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedIpAddresses sets the value of the AllowedIpAddresses field in IpAccessList_SdkV2.
func (o *IpAccessList_SdkV2) SetAllowedIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedIpAddresses = types.ListValueMust(t, vs)
}

type ListProviderSharesResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of provider shares.
	Shares types.List `tfsdk:"shares" tf:"optional"`
}

func (newState *ListProviderSharesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProviderSharesResponse_SdkV2) {
}

func (newState *ListProviderSharesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListProviderSharesResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProviderSharesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProviderSharesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"shares": reflect.TypeOf(ProviderShare_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderSharesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListProviderSharesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"shares":          o.Shares,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProviderSharesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shares": basetypes.ListType{
				ElemType: ProviderShare{}.Type(ctx),
			},
		},
	}
}

// GetShares returns the value of the Shares field in ListProviderSharesResponse_SdkV2 as
// a slice of ProviderShare_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListProviderSharesResponse_SdkV2) GetShares(ctx context.Context) ([]ProviderShare_SdkV2, bool) {
	if o.Shares.IsNull() || o.Shares.IsUnknown() {
		return nil, false
	}
	var v []ProviderShare_SdkV2
	d := o.Shares.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetShares sets the value of the Shares field in ListProviderSharesResponse_SdkV2.
func (o *ListProviderSharesResponse_SdkV2) SetShares(ctx context.Context, v []ProviderShare_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["shares"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Shares = types.ListValueMust(t, vs)
}

// List providers
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

func (newState *ListProvidersRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersRequest_SdkV2) {
}

func (newState *ListProvidersRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListProvidersRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProvidersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProvidersRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListProvidersRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_provider_global_metastore_id": o.DataProviderGlobalMetastoreId,
			"max_results":                       o.MaxResults,
			"page_token":                        o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProvidersRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of provider information objects.
	Providers types.List `tfsdk:"providers" tf:"optional"`
}

func (newState *ListProvidersResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersResponse_SdkV2) {
}

func (newState *ListProvidersResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListProvidersResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProvidersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProvidersResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"providers": reflect.TypeOf(ProviderInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListProvidersResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"providers":       o.Providers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProvidersResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"providers": basetypes.ListType{
				ElemType: ProviderInfo{}.Type(ctx),
			},
		},
	}
}

// GetProviders returns the value of the Providers field in ListProvidersResponse_SdkV2 as
// a slice of ProviderInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListProvidersResponse_SdkV2) GetProviders(ctx context.Context) ([]ProviderInfo_SdkV2, bool) {
	if o.Providers.IsNull() || o.Providers.IsUnknown() {
		return nil, false
	}
	var v []ProviderInfo_SdkV2
	d := o.Providers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviders sets the value of the Providers field in ListProvidersResponse_SdkV2.
func (o *ListProvidersResponse_SdkV2) SetProviders(ctx context.Context, v []ProviderInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["providers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Providers = types.ListValueMust(t, vs)
}

// List share recipients
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

func (newState *ListRecipientsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRecipientsRequest_SdkV2) {
}

func (newState *ListRecipientsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListRecipientsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRecipientsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRecipientsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRecipientsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListRecipientsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_recipient_global_metastore_id": o.DataRecipientGlobalMetastoreId,
			"max_results":                        o.MaxResults,
			"page_token":                         o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRecipientsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of recipient information objects.
	Recipients types.List `tfsdk:"recipients" tf:"optional"`
}

func (newState *ListRecipientsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRecipientsResponse_SdkV2) {
}

func (newState *ListRecipientsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListRecipientsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRecipientsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRecipientsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"recipients": reflect.TypeOf(RecipientInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRecipientsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListRecipientsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"recipients":      o.Recipients,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRecipientsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"recipients": basetypes.ListType{
				ElemType: RecipientInfo{}.Type(ctx),
			},
		},
	}
}

// GetRecipients returns the value of the Recipients field in ListRecipientsResponse_SdkV2 as
// a slice of RecipientInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListRecipientsResponse_SdkV2) GetRecipients(ctx context.Context) ([]RecipientInfo_SdkV2, bool) {
	if o.Recipients.IsNull() || o.Recipients.IsUnknown() {
		return nil, false
	}
	var v []RecipientInfo_SdkV2
	d := o.Recipients.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRecipients sets the value of the Recipients field in ListRecipientsResponse_SdkV2.
func (o *ListRecipientsResponse_SdkV2) SetRecipients(ctx context.Context, v []RecipientInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["recipients"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Recipients = types.ListValueMust(t, vs)
}

// List shares by Provider
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

func (newState *ListSharesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSharesRequest_SdkV2) {
}

func (newState *ListSharesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSharesRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSharesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSharesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSharesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSharesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"name":        o.Name,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSharesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of data share information objects.
	Shares types.List `tfsdk:"shares" tf:"optional"`
}

func (newState *ListSharesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSharesResponse_SdkV2) {
}

func (newState *ListSharesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSharesResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSharesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSharesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"shares": reflect.TypeOf(ShareInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSharesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSharesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"shares":          o.Shares,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSharesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shares": basetypes.ListType{
				ElemType: ShareInfo{}.Type(ctx),
			},
		},
	}
}

// GetShares returns the value of the Shares field in ListSharesResponse_SdkV2 as
// a slice of ShareInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSharesResponse_SdkV2) GetShares(ctx context.Context) ([]ShareInfo_SdkV2, bool) {
	if o.Shares.IsNull() || o.Shares.IsUnknown() {
		return nil, false
	}
	var v []ShareInfo_SdkV2
	d := o.Shares.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetShares sets the value of the Shares field in ListSharesResponse_SdkV2.
func (o *ListSharesResponse_SdkV2) SetShares(ctx context.Context, v []ShareInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["shares"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Shares = types.ListValueMust(t, vs)
}

type Partition_SdkV2 struct {
	// An array of partition values.
	Values types.List `tfsdk:"value" tf:"optional"`
}

func (newState *Partition_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Partition_SdkV2) {
}

func (newState *Partition_SdkV2) SyncEffectiveFieldsDuringRead(existingState Partition_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Partition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Partition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(PartitionValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Partition_SdkV2
// only implements ToObjectValue() and Type().
func (o Partition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Partition_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": basetypes.ListType{
				ElemType: PartitionValue{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in Partition_SdkV2 as
// a slice of PartitionValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Partition_SdkV2) GetValues(ctx context.Context) ([]PartitionValue_SdkV2, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []PartitionValue_SdkV2
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in Partition_SdkV2.
func (o *Partition_SdkV2) SetValues(ctx context.Context, v []PartitionValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type PartitionSpecificationPartition_SdkV2 struct {
	// An array of partition values.
	Values types.List `tfsdk:"value" tf:"optional"`
}

func (newState *PartitionSpecificationPartition_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartitionSpecificationPartition_SdkV2) {
}

func (newState *PartitionSpecificationPartition_SdkV2) SyncEffectiveFieldsDuringRead(existingState PartitionSpecificationPartition_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartitionSpecificationPartition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PartitionSpecificationPartition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(PartitionValue_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartitionSpecificationPartition_SdkV2
// only implements ToObjectValue() and Type().
func (o PartitionSpecificationPartition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PartitionSpecificationPartition_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": basetypes.ListType{
				ElemType: PartitionValue{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in PartitionSpecificationPartition_SdkV2 as
// a slice of PartitionValue_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PartitionSpecificationPartition_SdkV2) GetValues(ctx context.Context) ([]PartitionValue_SdkV2, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []PartitionValue_SdkV2
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in PartitionSpecificationPartition_SdkV2.
func (o *PartitionSpecificationPartition_SdkV2) SetValues(ctx context.Context, v []PartitionValue_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type PartitionValue_SdkV2 struct {
	// The name of the partition column.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The operator to apply for the value.
	Op types.String `tfsdk:"op" tf:"optional"`
	// The key of a Delta Sharing recipient's property. For example
	// "databricks-account-id". When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey types.String `tfsdk:"recipient_property_key" tf:"optional"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *PartitionValue_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartitionValue_SdkV2) {
}

func (newState *PartitionValue_SdkV2) SyncEffectiveFieldsDuringRead(existingState PartitionValue_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartitionValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PartitionValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartitionValue_SdkV2
// only implements ToObjectValue() and Type().
func (o PartitionValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":                   o.Name,
			"op":                     o.Op,
			"recipient_property_key": o.RecipientPropertyKey,
			"value":                  o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PartitionValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":                   types.StringType,
			"op":                     types.StringType,
			"recipient_property_key": types.StringType,
			"value":                  types.StringType,
		},
	}
}

type PrivilegeAssignment_SdkV2 struct {
	// The principal (user email address or group name).
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The privileges assigned to the principal.
	Privileges types.List `tfsdk:"privileges" tf:"optional"`
}

func (newState *PrivilegeAssignment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrivilegeAssignment_SdkV2) {
}

func (newState *PrivilegeAssignment_SdkV2) SyncEffectiveFieldsDuringRead(existingState PrivilegeAssignment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivilegeAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PrivilegeAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privileges": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivilegeAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o PrivilegeAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal":  o.Principal,
			"privileges": o.Privileges,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PrivilegeAssignment_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *PrivilegeAssignment_SdkV2) GetPrivileges(ctx context.Context) ([]types.String, bool) {
	if o.Privileges.IsNull() || o.Privileges.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Privileges.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivileges sets the value of the Privileges field in PrivilegeAssignment_SdkV2.
func (o *PrivilegeAssignment_SdkV2) SetPrivileges(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privileges"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Privileges = types.ListValueMust(t, vs)
}

type ProviderInfo_SdkV2 struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:"optional"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of Provider creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format <cloud>:<region>:<metastore-uuid>.
	DataProviderGlobalMetastoreId types.String `tfsdk:"data_provider_global_metastore_id" tf:"optional"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The name of the Provider.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN`.
	RecipientProfile types.List `tfsdk:"recipient_profile" tf:"optional,object"`
	// This field is only present when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region" tf:"optional"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified Share.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *ProviderInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderInfo_SdkV2) {
}

func (newState *ProviderInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ProviderInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProviderInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"recipient_profile": reflect.TypeOf(RecipientProfile_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ProviderInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authentication_type":               o.AuthenticationType,
			"cloud":                             o.Cloud,
			"comment":                           o.Comment,
			"created_at":                        o.CreatedAt,
			"created_by":                        o.CreatedBy,
			"data_provider_global_metastore_id": o.DataProviderGlobalMetastoreId,
			"metastore_id":                      o.MetastoreId,
			"name":                              o.Name,
			"owner":                             o.Owner,
			"recipient_profile":                 o.RecipientProfile,
			"recipient_profile_str":             o.RecipientProfileStr,
			"region":                            o.Region,
			"updated_at":                        o.UpdatedAt,
			"updated_by":                        o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ProviderInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: RecipientProfile{}.Type(ctx),
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
func (o *ProviderInfo_SdkV2) GetRecipientProfile(ctx context.Context) (RecipientProfile_SdkV2, bool) {
	var e RecipientProfile_SdkV2
	if o.RecipientProfile.IsNull() || o.RecipientProfile.IsUnknown() {
		return e, false
	}
	var v []RecipientProfile_SdkV2
	d := o.RecipientProfile.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRecipientProfile sets the value of the RecipientProfile field in ProviderInfo_SdkV2.
func (o *ProviderInfo_SdkV2) SetRecipientProfile(ctx context.Context, v RecipientProfile_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["recipient_profile"]
	o.RecipientProfile = types.ListValueMust(t, vs)
}

type ProviderShare_SdkV2 struct {
	// The name of the Provider Share.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ProviderShare_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderShare_SdkV2) {
}

func (newState *ProviderShare_SdkV2) SyncEffectiveFieldsDuringRead(existingState ProviderShare_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderShare.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProviderShare_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderShare_SdkV2
// only implements ToObjectValue() and Type().
func (o ProviderShare_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ProviderShare_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type RecipientInfo_SdkV2 struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated types.Bool `tfsdk:"activated" tf:"optional"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url" tf:"optional"`
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type" tf:"optional"`
	// Cloud vendor of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**`.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of recipient creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id" tf:"optional"`
	// IP Access List
	IpAccessList types.List `tfsdk:"ip_access_list" tf:"optional,object"`
	// Unique identifier of recipient's Unity Catalog metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of Recipient.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs types.List `tfsdk:"properties_kvpairs" tf:"optional,object"`
	// Cloud region of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code" tf:"optional"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens types.List `tfsdk:"tokens" tf:"optional"`
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of recipient updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *RecipientInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientInfo_SdkV2) {
}

func (newState *RecipientInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState RecipientInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RecipientInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RecipientInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList_SdkV2{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs_SdkV2{}),
		"tokens":             reflect.TypeOf(RecipientTokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RecipientInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o RecipientInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activated":                          o.Activated,
			"activation_url":                     o.ActivationUrl,
			"authentication_type":                o.AuthenticationType,
			"cloud":                              o.Cloud,
			"comment":                            o.Comment,
			"created_at":                         o.CreatedAt,
			"created_by":                         o.CreatedBy,
			"data_recipient_global_metastore_id": o.DataRecipientGlobalMetastoreId,
			"ip_access_list":                     o.IpAccessList,
			"metastore_id":                       o.MetastoreId,
			"name":                               o.Name,
			"owner":                              o.Owner,
			"properties_kvpairs":                 o.PropertiesKvpairs,
			"region":                             o.Region,
			"sharing_code":                       o.SharingCode,
			"tokens":                             o.Tokens,
			"updated_at":                         o.UpdatedAt,
			"updated_by":                         o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RecipientInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessList{}.Type(ctx),
			},
			"metastore_id": types.StringType,
			"name":         types.StringType,
			"owner":        types.StringType,
			"properties_kvpairs": basetypes.ListType{
				ElemType: SecurablePropertiesKvPairs{}.Type(ctx),
			},
			"region":       types.StringType,
			"sharing_code": types.StringType,
			"tokens": basetypes.ListType{
				ElemType: RecipientTokenInfo{}.Type(ctx),
			},
			"updated_at": types.Int64Type,
			"updated_by": types.StringType,
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in RecipientInfo_SdkV2 as
// a IpAccessList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RecipientInfo_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessList_SdkV2, bool) {
	var e IpAccessList_SdkV2
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessList_SdkV2
	d := o.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in RecipientInfo_SdkV2.
func (o *RecipientInfo_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	o.IpAccessList = types.ListValueMust(t, vs)
}

// GetPropertiesKvpairs returns the value of the PropertiesKvpairs field in RecipientInfo_SdkV2 as
// a SecurablePropertiesKvPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RecipientInfo_SdkV2) GetPropertiesKvpairs(ctx context.Context) (SecurablePropertiesKvPairs_SdkV2, bool) {
	var e SecurablePropertiesKvPairs_SdkV2
	if o.PropertiesKvpairs.IsNull() || o.PropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []SecurablePropertiesKvPairs_SdkV2
	d := o.PropertiesKvpairs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPropertiesKvpairs sets the value of the PropertiesKvpairs field in RecipientInfo_SdkV2.
func (o *RecipientInfo_SdkV2) SetPropertiesKvpairs(ctx context.Context, v SecurablePropertiesKvPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties_kvpairs"]
	o.PropertiesKvpairs = types.ListValueMust(t, vs)
}

// GetTokens returns the value of the Tokens field in RecipientInfo_SdkV2 as
// a slice of RecipientTokenInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RecipientInfo_SdkV2) GetTokens(ctx context.Context) ([]RecipientTokenInfo_SdkV2, bool) {
	if o.Tokens.IsNull() || o.Tokens.IsUnknown() {
		return nil, false
	}
	var v []RecipientTokenInfo_SdkV2
	d := o.Tokens.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokens sets the value of the Tokens field in RecipientInfo_SdkV2.
func (o *RecipientInfo_SdkV2) SetTokens(ctx context.Context, v []RecipientTokenInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tokens"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tokens = types.ListValueMust(t, vs)
}

type RecipientProfile_SdkV2 struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearer_token" tf:"optional"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version" tf:"optional"`
}

func (newState *RecipientProfile_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientProfile_SdkV2) {
}

func (newState *RecipientProfile_SdkV2) SyncEffectiveFieldsDuringRead(existingState RecipientProfile_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RecipientProfile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RecipientProfile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RecipientProfile_SdkV2
// only implements ToObjectValue() and Type().
func (o RecipientProfile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bearer_token":              o.BearerToken,
			"endpoint":                  o.Endpoint,
			"share_credentials_version": o.ShareCredentialsVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RecipientProfile_SdkV2) Type(ctx context.Context) attr.Type {
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
	ActivationUrl types.String `tfsdk:"activation_url" tf:"optional"`
	// Time at which this recipient Token was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of recipient token creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// Unique ID of the recipient token.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time at which this recipient Token was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of recipient Token updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *RecipientTokenInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientTokenInfo_SdkV2) {
}

func (newState *RecipientTokenInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState RecipientTokenInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RecipientTokenInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RecipientTokenInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RecipientTokenInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o RecipientTokenInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activation_url":  o.ActivationUrl,
			"created_at":      o.CreatedAt,
			"created_by":      o.CreatedBy,
			"expiration_time": o.ExpirationTime,
			"id":              o.Id,
			"updated_at":      o.UpdatedAt,
			"updated_by":      o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RecipientTokenInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

// Get an access token
type RetrieveTokenRequest_SdkV2 struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

func (newState *RetrieveTokenRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RetrieveTokenRequest_SdkV2) {
}

func (newState *RetrieveTokenRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState RetrieveTokenRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RetrieveTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RetrieveTokenRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RetrieveTokenRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RetrieveTokenRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activation_url": o.ActivationUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RetrieveTokenRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activation_url": types.StringType,
		},
	}
}

type RetrieveTokenResponse_SdkV2 struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearerToken" tf:"optional"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.String `tfsdk:"expirationTime" tf:"optional"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion types.Int64 `tfsdk:"shareCredentialsVersion" tf:"optional"`
}

func (newState *RetrieveTokenResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RetrieveTokenResponse_SdkV2) {
}

func (newState *RetrieveTokenResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState RetrieveTokenResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RetrieveTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RetrieveTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RetrieveTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RetrieveTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bearerToken":             o.BearerToken,
			"endpoint":                o.Endpoint,
			"expirationTime":          o.ExpirationTime,
			"shareCredentialsVersion": o.ShareCredentialsVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RetrieveTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearerToken":             types.StringType,
			"endpoint":                types.StringType,
			"expirationTime":          types.StringType,
			"shareCredentialsVersion": types.Int64Type,
		},
	}
}

type RotateRecipientToken_SdkV2 struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds types.Int64 `tfsdk:"existing_token_expire_in_seconds" tf:""`
	// The name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *RotateRecipientToken_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RotateRecipientToken_SdkV2) {
}

func (newState *RotateRecipientToken_SdkV2) SyncEffectiveFieldsDuringRead(existingState RotateRecipientToken_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RotateRecipientToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RotateRecipientToken_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RotateRecipientToken_SdkV2
// only implements ToObjectValue() and Type().
func (o RotateRecipientToken_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"existing_token_expire_in_seconds": o.ExistingTokenExpireInSeconds,
			"name":                             o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RotateRecipientToken_SdkV2) Type(ctx context.Context) attr.Type {
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
	Properties types.Map `tfsdk:"properties" tf:""`
}

func (newState *SecurablePropertiesKvPairs_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecurablePropertiesKvPairs_SdkV2) {
}

func (newState *SecurablePropertiesKvPairs_SdkV2) SyncEffectiveFieldsDuringRead(existingState SecurablePropertiesKvPairs_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SecurablePropertiesKvPairs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SecurablePropertiesKvPairs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecurablePropertiesKvPairs_SdkV2
// only implements ToObjectValue() and Type().
func (o SecurablePropertiesKvPairs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"properties": o.Properties,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SecurablePropertiesKvPairs_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SecurablePropertiesKvPairs_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
	if o.Properties.IsNull() || o.Properties.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Properties.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProperties sets the value of the Properties field in SecurablePropertiesKvPairs_SdkV2.
func (o *SecurablePropertiesKvPairs_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type ShareInfo_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"computed"`
	// Username of share creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"computed"`
	// Name of the share.
	Name types.String `tfsdk:"name" tf:"optional"`
	// A list of shared data objects within the share.
	Objects types.List `tfsdk:"object" tf:"optional"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner" tf:"computed"`
	// Storage Location URL (full path) for the share.
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"computed"`
	// Username of share updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"computed"`
}

func (newState *ShareInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareInfo_SdkV2) {
}

func (newState *ShareInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ShareInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ShareInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ShareInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"object": reflect.TypeOf(SharedDataObject_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ShareInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ShareInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":          o.Comment,
			"created_at":       o.CreatedAt,
			"created_by":       o.CreatedBy,
			"name":             o.Name,
			"object":           o.Objects,
			"owner":            o.Owner,
			"storage_location": o.StorageLocation,
			"storage_root":     o.StorageRoot,
			"updated_at":       o.UpdatedAt,
			"updated_by":       o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ShareInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":    types.StringType,
			"created_at": types.Int64Type,
			"created_by": types.StringType,
			"name":       types.StringType,
			"object": basetypes.ListType{
				ElemType: SharedDataObject{}.Type(ctx),
			},
			"owner":            types.StringType,
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
func (o *ShareInfo_SdkV2) GetObjects(ctx context.Context) ([]SharedDataObject_SdkV2, bool) {
	if o.Objects.IsNull() || o.Objects.IsUnknown() {
		return nil, false
	}
	var v []SharedDataObject_SdkV2
	d := o.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in ShareInfo_SdkV2.
func (o *ShareInfo_SdkV2) SetObjects(ctx context.Context, v []SharedDataObject_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Objects = types.ListValueMust(t, vs)
}

// Get recipient share permissions
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

func (newState *SharePermissionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharePermissionsRequest_SdkV2) {
}

func (newState *SharePermissionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState SharePermissionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SharePermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharePermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SharePermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"name":        o.Name,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SharePermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments" tf:"optional"`
	// The share name.
	ShareName types.String `tfsdk:"share_name" tf:"optional"`
}

func (newState *ShareToPrivilegeAssignment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareToPrivilegeAssignment_SdkV2) {
}

func (newState *ShareToPrivilegeAssignment_SdkV2) SyncEffectiveFieldsDuringRead(existingState ShareToPrivilegeAssignment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ShareToPrivilegeAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ShareToPrivilegeAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(PrivilegeAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ShareToPrivilegeAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o ShareToPrivilegeAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"privilege_assignments": o.PrivilegeAssignments,
			"share_name":            o.ShareName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ShareToPrivilegeAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment{}.Type(ctx),
			},
			"share_name": types.StringType,
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in ShareToPrivilegeAssignment_SdkV2 as
// a slice of PrivilegeAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ShareToPrivilegeAssignment_SdkV2) GetPrivilegeAssignments(ctx context.Context) ([]PrivilegeAssignment_SdkV2, bool) {
	if o.PrivilegeAssignments.IsNull() || o.PrivilegeAssignments.IsUnknown() {
		return nil, false
	}
	var v []PrivilegeAssignment_SdkV2
	d := o.PrivilegeAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in ShareToPrivilegeAssignment_SdkV2.
func (o *ShareToPrivilegeAssignment_SdkV2) SetPrivilegeAssignments(ctx context.Context, v []PrivilegeAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrivilegeAssignments = types.ListValueMust(t, vs)
}

type SharedDataObject_SdkV2 struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at" tf:"computed"`
	// Username of the sharer.
	AddedBy types.String `tfsdk:"added_by" tf:"computed"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled          types.Bool `tfsdk:"cdf_enabled" tf:"optional"`
	EffectiveCdfEnabled types.Bool `tfsdk:"effective_cdf_enabled" tf:"computed"`
	// A user-provided comment when adding the data object to the share.
	// [Update:OPT]
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	Content types.String `tfsdk:"content" tf:"optional"`
	// The type of the data object.
	DataObjectType types.String `tfsdk:"data_object_type" tf:"optional"`
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	HistoryDataSharingStatus          types.String `tfsdk:"history_data_sharing_status" tf:"optional"`
	EffectiveHistoryDataSharingStatus types.String `tfsdk:"effective_history_data_sharing_status" tf:"computed"`
	// A fully qualified name that uniquely identifies a data object.
	//
	// For example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`.
	Name types.String `tfsdk:"name" tf:""`
	// Array of partitions for the shared data.
	Partitions types.List `tfsdk:"partition" tf:"optional"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	SharedAs          types.String `tfsdk:"shared_as" tf:"optional"`
	EffectiveSharedAs types.String `tfsdk:"effective_shared_as" tf:"computed"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion          types.Int64 `tfsdk:"start_version" tf:"optional"`
	EffectiveStartVersion types.Int64 `tfsdk:"effective_start_version" tf:"computed"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status types.String `tfsdk:"status" tf:"computed"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `string_shared_as` name. The `string_shared_as` name must be unique
	// within a share. For notebooks, the new name should be the new notebook
	// file name.
	StringSharedAs types.String `tfsdk:"string_shared_as" tf:"optional"`
}

func (newState *SharedDataObject_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObject_SdkV2) {
	newState.EffectiveCdfEnabled = newState.CdfEnabled
	newState.CdfEnabled = plan.CdfEnabled
	newState.EffectiveHistoryDataSharingStatus = newState.HistoryDataSharingStatus
	newState.HistoryDataSharingStatus = plan.HistoryDataSharingStatus
	newState.EffectiveSharedAs = newState.SharedAs
	newState.SharedAs = plan.SharedAs
	newState.EffectiveStartVersion = newState.StartVersion
	newState.StartVersion = plan.StartVersion
}

func (newState *SharedDataObject_SdkV2) SyncEffectiveFieldsDuringRead(existingState SharedDataObject_SdkV2) {
	newState.EffectiveCdfEnabled = existingState.EffectiveCdfEnabled
	if existingState.EffectiveCdfEnabled.ValueBool() == newState.CdfEnabled.ValueBool() {
		newState.CdfEnabled = existingState.CdfEnabled
	}
	newState.EffectiveHistoryDataSharingStatus = existingState.EffectiveHistoryDataSharingStatus
	if existingState.EffectiveHistoryDataSharingStatus.ValueString() == newState.HistoryDataSharingStatus.ValueString() {
		newState.HistoryDataSharingStatus = existingState.HistoryDataSharingStatus
	}
	newState.EffectiveSharedAs = existingState.EffectiveSharedAs
	if existingState.EffectiveSharedAs.ValueString() == newState.SharedAs.ValueString() {
		newState.SharedAs = existingState.SharedAs
	}
	newState.EffectiveStartVersion = existingState.EffectiveStartVersion
	if existingState.EffectiveStartVersion.ValueInt64() == newState.StartVersion.ValueInt64() {
		newState.StartVersion = existingState.StartVersion
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharedDataObject.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SharedDataObject_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partition": reflect.TypeOf(Partition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharedDataObject_SdkV2
// only implements ToObjectValue() and Type().
func (o SharedDataObject_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"added_at":                              o.AddedAt,
			"added_by":                              o.AddedBy,
			"cdf_enabled":                           o.CdfEnabled,
			"effective_cdf_enabled":                 o.EffectiveCdfEnabled,
			"comment":                               o.Comment,
			"content":                               o.Content,
			"data_object_type":                      o.DataObjectType,
			"history_data_sharing_status":           o.HistoryDataSharingStatus,
			"effective_history_data_sharing_status": o.EffectiveHistoryDataSharingStatus,
			"name":                                  o.Name,
			"partition":                             o.Partitions,
			"shared_as":                             o.SharedAs,
			"effective_shared_as":                   o.EffectiveSharedAs,
			"start_version":                         o.StartVersion,
			"effective_start_version":               o.EffectiveStartVersion,
			"status":                                o.Status,
			"string_shared_as":                      o.StringSharedAs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SharedDataObject_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: Partition{}.Type(ctx),
			},
			"shared_as":               types.StringType,
			"effective_shared_as":     types.StringType,
			"start_version":           types.Int64Type,
			"effective_start_version": types.Int64Type,
			"status":                  types.StringType,
			"string_shared_as":        types.StringType,
		},
	}
}

// GetPartitions returns the value of the Partitions field in SharedDataObject_SdkV2 as
// a slice of Partition_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SharedDataObject_SdkV2) GetPartitions(ctx context.Context) ([]Partition_SdkV2, bool) {
	if o.Partitions.IsNull() || o.Partitions.IsUnknown() {
		return nil, false
	}
	var v []Partition_SdkV2
	d := o.Partitions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPartitions sets the value of the Partitions field in SharedDataObject_SdkV2.
func (o *SharedDataObject_SdkV2) SetPartitions(ctx context.Context, v []Partition_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["partition"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Partitions = types.ListValueMust(t, vs)
}

type SharedDataObjectUpdate_SdkV2 struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	Action types.String `tfsdk:"action" tf:"optional"`
	// The data object that is being added, removed, or updated.
	DataObject types.List `tfsdk:"data_object" tf:"optional,object"`
}

func (newState *SharedDataObjectUpdate_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObjectUpdate_SdkV2) {
}

func (newState *SharedDataObjectUpdate_SdkV2) SyncEffectiveFieldsDuringRead(existingState SharedDataObjectUpdate_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharedDataObjectUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SharedDataObjectUpdate_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_object": reflect.TypeOf(SharedDataObject_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharedDataObjectUpdate_SdkV2
// only implements ToObjectValue() and Type().
func (o SharedDataObjectUpdate_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action":      o.Action,
			"data_object": o.DataObject,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SharedDataObjectUpdate_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action": types.StringType,
			"data_object": basetypes.ListType{
				ElemType: SharedDataObject{}.Type(ctx),
			},
		},
	}
}

// GetDataObject returns the value of the DataObject field in SharedDataObjectUpdate_SdkV2 as
// a SharedDataObject_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SharedDataObjectUpdate_SdkV2) GetDataObject(ctx context.Context) (SharedDataObject_SdkV2, bool) {
	var e SharedDataObject_SdkV2
	if o.DataObject.IsNull() || o.DataObject.IsUnknown() {
		return e, false
	}
	var v []SharedDataObject_SdkV2
	d := o.DataObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataObject sets the value of the DataObject field in SharedDataObjectUpdate_SdkV2.
func (o *SharedDataObjectUpdate_SdkV2) SetDataObject(ctx context.Context, v SharedDataObject_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_object"]
	o.DataObject = types.ListValueMust(t, vs)
}

type UpdatePermissionsResponse_SdkV2 struct {
}

func (newState *UpdatePermissionsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePermissionsResponse_SdkV2) {
}

func (newState *UpdatePermissionsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdatePermissionsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePermissionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePermissionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePermissionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdatePermissionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePermissionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateProvider_SdkV2 struct {
	// Description about the provider.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the provider.
	Name types.String `tfsdk:"-"`
	// New name for the provider.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str" tf:"optional"`
}

func (newState *UpdateProvider_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProvider_SdkV2) {
}

func (newState *UpdateProvider_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateProvider_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProvider.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateProvider_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProvider_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateProvider_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":               o.Comment,
			"name":                  o.Name,
			"new_name":              o.NewName,
			"owner":                 o.Owner,
			"recipient_profile_str": o.RecipientProfileStr,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateProvider_SdkV2) Type(ctx context.Context) attr.Type {
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
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// IP Access List
	IpAccessList types.List `tfsdk:"ip_access_list" tf:"optional,object"`
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
	// New name for the recipient.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs types.List `tfsdk:"properties_kvpairs" tf:"optional,object"`
}

func (newState *UpdateRecipient_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRecipient_SdkV2) {
}

func (newState *UpdateRecipient_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateRecipient_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRecipient.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRecipient_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList_SdkV2{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRecipient_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateRecipient_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":            o.Comment,
			"expiration_time":    o.ExpirationTime,
			"ip_access_list":     o.IpAccessList,
			"name":               o.Name,
			"new_name":           o.NewName,
			"owner":              o.Owner,
			"properties_kvpairs": o.PropertiesKvpairs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRecipient_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":         types.StringType,
			"expiration_time": types.Int64Type,
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessList{}.Type(ctx),
			},
			"name":     types.StringType,
			"new_name": types.StringType,
			"owner":    types.StringType,
			"properties_kvpairs": basetypes.ListType{
				ElemType: SecurablePropertiesKvPairs{}.Type(ctx),
			},
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in UpdateRecipient_SdkV2 as
// a IpAccessList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRecipient_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessList_SdkV2, bool) {
	var e IpAccessList_SdkV2
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessList_SdkV2
	d := o.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in UpdateRecipient_SdkV2.
func (o *UpdateRecipient_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	o.IpAccessList = types.ListValueMust(t, vs)
}

// GetPropertiesKvpairs returns the value of the PropertiesKvpairs field in UpdateRecipient_SdkV2 as
// a SecurablePropertiesKvPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRecipient_SdkV2) GetPropertiesKvpairs(ctx context.Context) (SecurablePropertiesKvPairs_SdkV2, bool) {
	var e SecurablePropertiesKvPairs_SdkV2
	if o.PropertiesKvpairs.IsNull() || o.PropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []SecurablePropertiesKvPairs_SdkV2
	d := o.PropertiesKvpairs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPropertiesKvpairs sets the value of the PropertiesKvpairs field in UpdateRecipient_SdkV2.
func (o *UpdateRecipient_SdkV2) SetPropertiesKvpairs(ctx context.Context, v SecurablePropertiesKvPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties_kvpairs"]
	o.PropertiesKvpairs = types.ListValueMust(t, vs)
}

type UpdateResponse_SdkV2 struct {
}

func (newState *UpdateResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse_SdkV2) {
}

func (newState *UpdateResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateShare_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// New name for the share.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner" tf:"computed"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Array of shared data object updates.
	Updates types.List `tfsdk:"updates" tf:"optional"`
}

func (newState *UpdateShare_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateShare_SdkV2) {
}

func (newState *UpdateShare_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateShare_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateShare.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateShare_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"updates": reflect.TypeOf(SharedDataObjectUpdate_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateShare_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateShare_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":      o.Comment,
			"name":         o.Name,
			"new_name":     o.NewName,
			"owner":        o.Owner,
			"storage_root": o.StorageRoot,
			"updates":      o.Updates,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateShare_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":      types.StringType,
			"name":         types.StringType,
			"new_name":     types.StringType,
			"owner":        types.StringType,
			"storage_root": types.StringType,
			"updates": basetypes.ListType{
				ElemType: SharedDataObjectUpdate{}.Type(ctx),
			},
		},
	}
}

// GetUpdates returns the value of the Updates field in UpdateShare_SdkV2 as
// a slice of SharedDataObjectUpdate_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateShare_SdkV2) GetUpdates(ctx context.Context) ([]SharedDataObjectUpdate_SdkV2, bool) {
	if o.Updates.IsNull() || o.Updates.IsUnknown() {
		return nil, false
	}
	var v []SharedDataObjectUpdate_SdkV2
	d := o.Updates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdates sets the value of the Updates field in UpdateShare_SdkV2.
func (o *UpdateShare_SdkV2) SetUpdates(ctx context.Context, v []SharedDataObjectUpdate_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Updates = types.ListValueMust(t, vs)
}

type UpdateSharePermissions_SdkV2 struct {
	// Array of permission changes.
	Changes types.List `tfsdk:"changes" tf:"optional"`
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
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *UpdateSharePermissions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSharePermissions_SdkV2) {
}

func (newState *UpdateSharePermissions_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateSharePermissions_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSharePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateSharePermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"changes": reflect.TypeOf(catalog_tf.PermissionsChange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSharePermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateSharePermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"changes":     o.Changes,
			"max_results": o.MaxResults,
			"name":        o.Name,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateSharePermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"changes": basetypes.ListType{
				ElemType: catalog_tf.PermissionsChange{}.Type(ctx),
			},
			"max_results": types.Int64Type,
			"name":        types.StringType,
			"page_token":  types.StringType,
		},
	}
}

// GetChanges returns the value of the Changes field in UpdateSharePermissions_SdkV2 as
// a slice of catalog_tf.PermissionsChange_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateSharePermissions_SdkV2) GetChanges(ctx context.Context) ([]catalog_tf.PermissionsChange_SdkV2, bool) {
	if o.Changes.IsNull() || o.Changes.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.PermissionsChange_SdkV2
	d := o.Changes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChanges sets the value of the Changes field in UpdateSharePermissions_SdkV2.
func (o *UpdateSharePermissions_SdkV2) SetChanges(ctx context.Context, v []catalog_tf.PermissionsChange_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["changes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Changes = types.ListValueMust(t, vs)
}
