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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateProvider struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type"`
	// Description about the provider.
	Comment types.String `tfsdk:"comment"`
	// The name of the Provider.
	Name types.String `tfsdk:"name"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str"`
}

func (newState *CreateProvider) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateProvider) {
}

func (newState *CreateProvider) SyncEffectiveFieldsDuringRead(existingState CreateProvider) {
}

func (c CreateProvider) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateProvider) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProvider
// only implements ToObjectValue() and Type().
func (o CreateProvider) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateProvider) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authentication_type":   types.StringType,
			"comment":               types.StringType,
			"name":                  types.StringType,
			"recipient_profile_str": types.StringType,
		},
	}
}

type CreateRecipient struct {
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type"`
	// Description about the recipient.
	Comment types.String `tfsdk:"comment"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is required when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId types.String `tfsdk:"data_recipient_global_metastore_id"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// IP Access List
	IpAccessList types.Object `tfsdk:"ip_access_list"`
	// Name of Recipient.
	Name types.String `tfsdk:"name"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs types.Object `tfsdk:"properties_kvpairs"`
	// The one-time sharing code provided by the data recipient. This field is
	// required when the __authentication_type__ is **DATABRICKS**.
	SharingCode types.String `tfsdk:"sharing_code"`
}

func (newState *CreateRecipient) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRecipient) {
}

func (newState *CreateRecipient) SyncEffectiveFieldsDuringRead(existingState CreateRecipient) {
}

func (c CreateRecipient) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["authentication_type"] = attrs["authentication_type"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["data_recipient_global_metastore_id"] = attrs["data_recipient_global_metastore_id"].SetOptional()
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties_kvpairs"] = attrs["properties_kvpairs"].SetOptional()
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
func (a CreateRecipient) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRecipient
// only implements ToObjectValue() and Type().
func (o CreateRecipient) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateRecipient) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authentication_type":                types.StringType,
			"comment":                            types.StringType,
			"data_recipient_global_metastore_id": types.StringType,
			"expiration_time":                    types.Int64Type,
			"ip_access_list":                     IpAccessList{}.Type(ctx),
			"name":                               types.StringType,
			"owner":                              types.StringType,
			"properties_kvpairs":                 SecurablePropertiesKvPairs{}.Type(ctx),
			"sharing_code":                       types.StringType,
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in CreateRecipient as
// a IpAccessList value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRecipient) GetIpAccessList(ctx context.Context) (IpAccessList, bool) {
	var e IpAccessList
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessList
	d := o.IpAccessList.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in CreateRecipient.
func (o *CreateRecipient) SetIpAccessList(ctx context.Context, v IpAccessList) {
	vs := v.ToObjectValue(ctx)
	o.IpAccessList = vs
}

// GetPropertiesKvpairs returns the value of the PropertiesKvpairs field in CreateRecipient as
// a SecurablePropertiesKvPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRecipient) GetPropertiesKvpairs(ctx context.Context) (SecurablePropertiesKvPairs, bool) {
	var e SecurablePropertiesKvPairs
	if o.PropertiesKvpairs.IsNull() || o.PropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []SecurablePropertiesKvPairs
	d := o.PropertiesKvpairs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPropertiesKvpairs sets the value of the PropertiesKvpairs field in CreateRecipient.
func (o *CreateRecipient) SetPropertiesKvpairs(ctx context.Context, v SecurablePropertiesKvPairs) {
	vs := v.ToObjectValue(ctx)
	o.PropertiesKvpairs = vs
}

type CreateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Name of the share.
	Name types.String `tfsdk:"name"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root"`
}

func (newState *CreateShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateShare) {
}

func (newState *CreateShare) SyncEffectiveFieldsDuringRead(existingState CreateShare) {
}

func (c CreateShare) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateShare) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateShare
// only implements ToObjectValue() and Type().
func (o CreateShare) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":      o.Comment,
			"name":         o.Name,
			"storage_root": o.StorageRoot,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateShare) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":      types.StringType,
			"name":         types.StringType,
			"storage_root": types.StringType,
		},
	}
}

// Delete a provider
type DeleteProviderRequest struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteProviderRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProviderRequest
// only implements ToObjectValue() and Type().
func (o DeleteProviderRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteProviderRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Delete a share recipient
type DeleteRecipientRequest struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRecipientRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRecipientRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRecipientRequest
// only implements ToObjectValue() and Type().
func (o DeleteRecipientRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRecipientRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse
// only implements ToObjectValue() and Type().
func (o DeleteResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a share
type DeleteShareRequest struct {
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteShareRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteShareRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteShareRequest
// only implements ToObjectValue() and Type().
func (o DeleteShareRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteShareRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a share activation URL
type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetActivationUrlInfoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetActivationUrlInfoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetActivationUrlInfoRequest
// only implements ToObjectValue() and Type().
func (o GetActivationUrlInfoRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activation_url": o.ActivationUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activation_url": types.StringType,
		},
	}
}

type GetActivationUrlInfoResponse struct {
}

func (newState *GetActivationUrlInfoResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetActivationUrlInfoResponse) {
}

func (newState *GetActivationUrlInfoResponse) SyncEffectiveFieldsDuringRead(existingState GetActivationUrlInfoResponse) {
}

func (c GetActivationUrlInfoResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetActivationUrlInfoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetActivationUrlInfoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetActivationUrlInfoResponse
// only implements ToObjectValue() and Type().
func (o GetActivationUrlInfoResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Get a provider
type GetProviderRequest struct {
	// Name of the provider.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProviderRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetProviderRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProviderRequest
// only implements ToObjectValue() and Type().
func (o GetProviderRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetProviderRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a share recipient
type GetRecipientRequest struct {
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRecipientRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRecipientRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRecipientRequest
// only implements ToObjectValue() and Type().
func (o GetRecipientRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRecipientRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetRecipientSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of data share permissions for a recipient.
	PermissionsOut types.List `tfsdk:"permissions_out"`
}

func (newState *GetRecipientSharePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRecipientSharePermissionsResponse) {
}

func (newState *GetRecipientSharePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState GetRecipientSharePermissionsResponse) {
}

func (c GetRecipientSharePermissionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRecipientSharePermissionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions_out": reflect.TypeOf(ShareToPrivilegeAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRecipientSharePermissionsResponse
// only implements ToObjectValue() and Type().
func (o GetRecipientSharePermissionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"permissions_out": o.PermissionsOut,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRecipientSharePermissionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"permissions_out": basetypes.ListType{
				ElemType: ShareToPrivilegeAssignment{}.Type(ctx),
			},
		},
	}
}

// GetPermissionsOut returns the value of the PermissionsOut field in GetRecipientSharePermissionsResponse as
// a slice of ShareToPrivilegeAssignment values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetRecipientSharePermissionsResponse) GetPermissionsOut(ctx context.Context) ([]ShareToPrivilegeAssignment, bool) {
	if o.PermissionsOut.IsNull() || o.PermissionsOut.IsUnknown() {
		return nil, false
	}
	var v []ShareToPrivilegeAssignment
	d := o.PermissionsOut.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionsOut sets the value of the PermissionsOut field in GetRecipientSharePermissionsResponse.
func (o *GetRecipientSharePermissionsResponse) SetPermissionsOut(ctx context.Context, v []ShareToPrivilegeAssignment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permissions_out"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionsOut = types.ListValueMust(t, vs)
}

// Get a share
type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData types.Bool `tfsdk:"-"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetShareRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetShareRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetShareRequest
// only implements ToObjectValue() and Type().
func (o GetShareRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_shared_data": o.IncludeSharedData,
			"name":                o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetShareRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_shared_data": types.BoolType,
			"name":                types.StringType,
		},
	}
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	AllowedIpAddresses types.List `tfsdk:"allowed_ip_addresses"`
}

func (newState *IpAccessList) SyncEffectiveFieldsDuringCreateOrUpdate(plan IpAccessList) {
}

func (newState *IpAccessList) SyncEffectiveFieldsDuringRead(existingState IpAccessList) {
}

func (c IpAccessList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a IpAccessList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IpAccessList
// only implements ToObjectValue() and Type().
func (o IpAccessList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_ip_addresses": o.AllowedIpAddresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IpAccessList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAllowedIpAddresses returns the value of the AllowedIpAddresses field in IpAccessList as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *IpAccessList) GetAllowedIpAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetAllowedIpAddresses sets the value of the AllowedIpAddresses field in IpAccessList.
func (o *IpAccessList) SetAllowedIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedIpAddresses = types.ListValueMust(t, vs)
}

type ListProviderSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of provider shares.
	Shares types.List `tfsdk:"shares"`
}

func (newState *ListProviderSharesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProviderSharesResponse) {
}

func (newState *ListProviderSharesResponse) SyncEffectiveFieldsDuringRead(existingState ListProviderSharesResponse) {
}

func (c ListProviderSharesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListProviderSharesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"shares": reflect.TypeOf(ProviderShare{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderSharesResponse
// only implements ToObjectValue() and Type().
func (o ListProviderSharesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"shares":          o.Shares,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProviderSharesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shares": basetypes.ListType{
				ElemType: ProviderShare{}.Type(ctx),
			},
		},
	}
}

// GetShares returns the value of the Shares field in ListProviderSharesResponse as
// a slice of ProviderShare values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListProviderSharesResponse) GetShares(ctx context.Context) ([]ProviderShare, bool) {
	if o.Shares.IsNull() || o.Shares.IsUnknown() {
		return nil, false
	}
	var v []ProviderShare
	d := o.Shares.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetShares sets the value of the Shares field in ListProviderSharesResponse.
func (o *ListProviderSharesResponse) SetShares(ctx context.Context, v []ProviderShare) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["shares"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Shares = types.ListValueMust(t, vs)
}

// List providers
type ListProvidersRequest struct {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProvidersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProvidersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersRequest
// only implements ToObjectValue() and Type().
func (o ListProvidersRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_provider_global_metastore_id": o.DataProviderGlobalMetastoreId,
			"max_results":                       o.MaxResults,
			"page_token":                        o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProvidersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_provider_global_metastore_id": types.StringType,
			"max_results":                       types.Int64Type,
			"page_token":                        types.StringType,
		},
	}
}

type ListProvidersResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of provider information objects.
	Providers types.List `tfsdk:"providers"`
}

func (newState *ListProvidersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersResponse) {
}

func (newState *ListProvidersResponse) SyncEffectiveFieldsDuringRead(existingState ListProvidersResponse) {
}

func (c ListProvidersResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListProvidersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"providers": reflect.TypeOf(ProviderInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProvidersResponse
// only implements ToObjectValue() and Type().
func (o ListProvidersResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"providers":       o.Providers,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProvidersResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"providers": basetypes.ListType{
				ElemType: ProviderInfo{}.Type(ctx),
			},
		},
	}
}

// GetProviders returns the value of the Providers field in ListProvidersResponse as
// a slice of ProviderInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListProvidersResponse) GetProviders(ctx context.Context) ([]ProviderInfo, bool) {
	if o.Providers.IsNull() || o.Providers.IsUnknown() {
		return nil, false
	}
	var v []ProviderInfo
	d := o.Providers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProviders sets the value of the Providers field in ListProvidersResponse.
func (o *ListProvidersResponse) SetProviders(ctx context.Context, v []ProviderInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["providers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Providers = types.ListValueMust(t, vs)
}

// List share recipients
type ListRecipientsRequest struct {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRecipientsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRecipientsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRecipientsRequest
// only implements ToObjectValue() and Type().
func (o ListRecipientsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_recipient_global_metastore_id": o.DataRecipientGlobalMetastoreId,
			"max_results":                        o.MaxResults,
			"page_token":                         o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRecipientsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_recipient_global_metastore_id": types.StringType,
			"max_results":                        types.Int64Type,
			"page_token":                         types.StringType,
		},
	}
}

type ListRecipientsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of recipient information objects.
	Recipients types.List `tfsdk:"recipients"`
}

func (newState *ListRecipientsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRecipientsResponse) {
}

func (newState *ListRecipientsResponse) SyncEffectiveFieldsDuringRead(existingState ListRecipientsResponse) {
}

func (c ListRecipientsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListRecipientsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"recipients": reflect.TypeOf(RecipientInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRecipientsResponse
// only implements ToObjectValue() and Type().
func (o ListRecipientsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"recipients":      o.Recipients,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRecipientsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"recipients": basetypes.ListType{
				ElemType: RecipientInfo{}.Type(ctx),
			},
		},
	}
}

// GetRecipients returns the value of the Recipients field in ListRecipientsResponse as
// a slice of RecipientInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListRecipientsResponse) GetRecipients(ctx context.Context) ([]RecipientInfo, bool) {
	if o.Recipients.IsNull() || o.Recipients.IsUnknown() {
		return nil, false
	}
	var v []RecipientInfo
	d := o.Recipients.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRecipients sets the value of the Recipients field in ListRecipientsResponse.
func (o *ListRecipientsResponse) SetRecipients(ctx context.Context, v []RecipientInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["recipients"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Recipients = types.ListValueMust(t, vs)
}

// List shares by Provider
type ListSharesRequest struct {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSharesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSharesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSharesRequest
// only implements ToObjectValue() and Type().
func (o ListSharesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"name":        o.Name,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSharesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"name":        types.StringType,
			"page_token":  types.StringType,
		},
	}
}

type ListSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of data share information objects.
	Shares types.List `tfsdk:"shares"`
}

func (newState *ListSharesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSharesResponse) {
}

func (newState *ListSharesResponse) SyncEffectiveFieldsDuringRead(existingState ListSharesResponse) {
}

func (c ListSharesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListSharesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"shares": reflect.TypeOf(ShareInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSharesResponse
// only implements ToObjectValue() and Type().
func (o ListSharesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"shares":          o.Shares,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSharesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"shares": basetypes.ListType{
				ElemType: ShareInfo{}.Type(ctx),
			},
		},
	}
}

// GetShares returns the value of the Shares field in ListSharesResponse as
// a slice of ShareInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSharesResponse) GetShares(ctx context.Context) ([]ShareInfo, bool) {
	if o.Shares.IsNull() || o.Shares.IsUnknown() {
		return nil, false
	}
	var v []ShareInfo
	d := o.Shares.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetShares sets the value of the Shares field in ListSharesResponse.
func (o *ListSharesResponse) SetShares(ctx context.Context, v []ShareInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["shares"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Shares = types.ListValueMust(t, vs)
}

type Partition struct {
	// An array of partition values.
	Values types.List `tfsdk:"value"`
}

func (newState *Partition) SyncEffectiveFieldsDuringCreateOrUpdate(plan Partition) {
}

func (newState *Partition) SyncEffectiveFieldsDuringRead(existingState Partition) {
}

func (c Partition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Partition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(PartitionValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Partition
// only implements ToObjectValue() and Type().
func (o Partition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Partition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": basetypes.ListType{
				ElemType: PartitionValue{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in Partition as
// a slice of PartitionValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *Partition) GetValues(ctx context.Context) ([]PartitionValue, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []PartitionValue
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in Partition.
func (o *Partition) SetValues(ctx context.Context, v []PartitionValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type PartitionSpecificationPartition struct {
	// An array of partition values.
	Values types.List `tfsdk:"value"`
}

func (newState *PartitionSpecificationPartition) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartitionSpecificationPartition) {
}

func (newState *PartitionSpecificationPartition) SyncEffectiveFieldsDuringRead(existingState PartitionSpecificationPartition) {
}

func (c PartitionSpecificationPartition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartitionSpecificationPartition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PartitionSpecificationPartition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(PartitionValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartitionSpecificationPartition
// only implements ToObjectValue() and Type().
func (o PartitionSpecificationPartition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PartitionSpecificationPartition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": basetypes.ListType{
				ElemType: PartitionValue{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in PartitionSpecificationPartition as
// a slice of PartitionValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *PartitionSpecificationPartition) GetValues(ctx context.Context) ([]PartitionValue, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []PartitionValue
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in PartitionSpecificationPartition.
func (o *PartitionSpecificationPartition) SetValues(ctx context.Context, v []PartitionValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type PartitionValue struct {
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

func (newState *PartitionValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartitionValue) {
}

func (newState *PartitionValue) SyncEffectiveFieldsDuringRead(existingState PartitionValue) {
}

func (c PartitionValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PartitionValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartitionValue
// only implements ToObjectValue() and Type().
func (o PartitionValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PartitionValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":                   types.StringType,
			"op":                     types.StringType,
			"recipient_property_key": types.StringType,
			"value":                  types.StringType,
		},
	}
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal types.String `tfsdk:"principal"`
	// The privileges assigned to the principal.
	Privileges types.List `tfsdk:"privileges"`
}

func (newState *PrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrivilegeAssignment) {
}

func (newState *PrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState PrivilegeAssignment) {
}

func (c PrivilegeAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PrivilegeAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privileges": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivilegeAssignment
// only implements ToObjectValue() and Type().
func (o PrivilegeAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal":  o.Principal,
			"privileges": o.Privileges,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PrivilegeAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"privileges": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPrivileges returns the value of the Privileges field in PrivilegeAssignment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PrivilegeAssignment) GetPrivileges(ctx context.Context) ([]types.String, bool) {
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

// SetPrivileges sets the value of the Privileges field in PrivilegeAssignment.
func (o *PrivilegeAssignment) SetPrivileges(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privileges"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Privileges = types.ListValueMust(t, vs)
}

type ProviderInfo struct {
	// The delta sharing authentication type.
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
	// identifier is of format <cloud>:<region>:<metastore-uuid>.
	DataProviderGlobalMetastoreId types.String `tfsdk:"data_provider_global_metastore_id"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The name of the Provider.
	Name types.String `tfsdk:"name"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN`.
	RecipientProfile types.Object `tfsdk:"recipient_profile"`
	// This field is only present when the authentication_type is `TOKEN` or not
	// provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region types.String `tfsdk:"region"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified Share.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *ProviderInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderInfo) {
}

func (newState *ProviderInfo) SyncEffectiveFieldsDuringRead(existingState ProviderInfo) {
}

func (c ProviderInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["authentication_type"] = attrs["authentication_type"].SetOptional()
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["data_provider_global_metastore_id"] = attrs["data_provider_global_metastore_id"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["recipient_profile"] = attrs["recipient_profile"].SetOptional()
	attrs["recipient_profile_str"] = attrs["recipient_profile_str"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProviderInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"recipient_profile": reflect.TypeOf(RecipientProfile{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderInfo
// only implements ToObjectValue() and Type().
func (o ProviderInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ProviderInfo) Type(ctx context.Context) attr.Type {
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
			"recipient_profile":                 RecipientProfile{}.Type(ctx),
			"recipient_profile_str":             types.StringType,
			"region":                            types.StringType,
			"updated_at":                        types.Int64Type,
			"updated_by":                        types.StringType,
		},
	}
}

// GetRecipientProfile returns the value of the RecipientProfile field in ProviderInfo as
// a RecipientProfile value.
// If the field is unknown or null, the boolean return value is false.
func (o *ProviderInfo) GetRecipientProfile(ctx context.Context) (RecipientProfile, bool) {
	var e RecipientProfile
	if o.RecipientProfile.IsNull() || o.RecipientProfile.IsUnknown() {
		return e, false
	}
	var v []RecipientProfile
	d := o.RecipientProfile.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRecipientProfile sets the value of the RecipientProfile field in ProviderInfo.
func (o *ProviderInfo) SetRecipientProfile(ctx context.Context, v RecipientProfile) {
	vs := v.ToObjectValue(ctx)
	o.RecipientProfile = vs
}

type ProviderShare struct {
	// The name of the Provider Share.
	Name types.String `tfsdk:"name"`
}

func (newState *ProviderShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderShare) {
}

func (newState *ProviderShare) SyncEffectiveFieldsDuringRead(existingState ProviderShare) {
}

func (c ProviderShare) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ProviderShare) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderShare
// only implements ToObjectValue() and Type().
func (o ProviderShare) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ProviderShare) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated types.Bool `tfsdk:"activated"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url"`
	// The delta sharing authentication type.
	AuthenticationType types.String `tfsdk:"authentication_type"`
	// Cloud vendor of the recipient's Unity Catalog Metstore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**`.
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
	// IP Access List
	IpAccessList types.Object `tfsdk:"ip_access_list"`
	// Unique identifier of recipient's Unity Catalog metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of Recipient.
	Name types.String `tfsdk:"name"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs.
	PropertiesKvpairs types.Object `tfsdk:"properties_kvpairs"`
	// Cloud region of the recipient's Unity Catalog Metstore. This field is
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

func (newState *RecipientInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientInfo) {
}

func (newState *RecipientInfo) SyncEffectiveFieldsDuringRead(existingState RecipientInfo) {
}

func (c RecipientInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activated"] = attrs["activated"].SetOptional()
	attrs["activation_url"] = attrs["activation_url"].SetOptional()
	attrs["authentication_type"] = attrs["authentication_type"].SetOptional()
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["data_recipient_global_metastore_id"] = attrs["data_recipient_global_metastore_id"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties_kvpairs"] = attrs["properties_kvpairs"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["sharing_code"] = attrs["sharing_code"].SetOptional()
	attrs["tokens"] = attrs["tokens"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RecipientInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RecipientInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs{}),
		"tokens":             reflect.TypeOf(RecipientTokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RecipientInfo
// only implements ToObjectValue() and Type().
func (o RecipientInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RecipientInfo) Type(ctx context.Context) attr.Type {
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
			"ip_access_list":                     IpAccessList{}.Type(ctx),
			"metastore_id":                       types.StringType,
			"name":                               types.StringType,
			"owner":                              types.StringType,
			"properties_kvpairs":                 SecurablePropertiesKvPairs{}.Type(ctx),
			"region":                             types.StringType,
			"sharing_code":                       types.StringType,
			"tokens": basetypes.ListType{
				ElemType: RecipientTokenInfo{}.Type(ctx),
			},
			"updated_at": types.Int64Type,
			"updated_by": types.StringType,
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in RecipientInfo as
// a IpAccessList value.
// If the field is unknown or null, the boolean return value is false.
func (o *RecipientInfo) GetIpAccessList(ctx context.Context) (IpAccessList, bool) {
	var e IpAccessList
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessList
	d := o.IpAccessList.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in RecipientInfo.
func (o *RecipientInfo) SetIpAccessList(ctx context.Context, v IpAccessList) {
	vs := v.ToObjectValue(ctx)
	o.IpAccessList = vs
}

// GetPropertiesKvpairs returns the value of the PropertiesKvpairs field in RecipientInfo as
// a SecurablePropertiesKvPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *RecipientInfo) GetPropertiesKvpairs(ctx context.Context) (SecurablePropertiesKvPairs, bool) {
	var e SecurablePropertiesKvPairs
	if o.PropertiesKvpairs.IsNull() || o.PropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []SecurablePropertiesKvPairs
	d := o.PropertiesKvpairs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPropertiesKvpairs sets the value of the PropertiesKvpairs field in RecipientInfo.
func (o *RecipientInfo) SetPropertiesKvpairs(ctx context.Context, v SecurablePropertiesKvPairs) {
	vs := v.ToObjectValue(ctx)
	o.PropertiesKvpairs = vs
}

// GetTokens returns the value of the Tokens field in RecipientInfo as
// a slice of RecipientTokenInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *RecipientInfo) GetTokens(ctx context.Context) ([]RecipientTokenInfo, bool) {
	if o.Tokens.IsNull() || o.Tokens.IsUnknown() {
		return nil, false
	}
	var v []RecipientTokenInfo
	d := o.Tokens.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokens sets the value of the Tokens field in RecipientInfo.
func (o *RecipientInfo) SetTokens(ctx context.Context, v []RecipientTokenInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tokens"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tokens = types.ListValueMust(t, vs)
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearer_token"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version"`
}

func (newState *RecipientProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientProfile) {
}

func (newState *RecipientProfile) SyncEffectiveFieldsDuringRead(existingState RecipientProfile) {
}

func (c RecipientProfile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RecipientProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RecipientProfile
// only implements ToObjectValue() and Type().
func (o RecipientProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bearer_token":              o.BearerToken,
			"endpoint":                  o.Endpoint,
			"share_credentials_version": o.ShareCredentialsVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RecipientProfile) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearer_token":              types.StringType,
			"endpoint":                  types.StringType,
			"share_credentials_version": types.Int64Type,
		},
	}
}

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl types.String `tfsdk:"activation_url"`
	// Time at which this recipient Token was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of recipient token creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// Unique ID of the recipient token.
	Id types.String `tfsdk:"id"`
	// Time at which this recipient Token was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of recipient Token updater.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *RecipientTokenInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientTokenInfo) {
}

func (newState *RecipientTokenInfo) SyncEffectiveFieldsDuringRead(existingState RecipientTokenInfo) {
}

func (c RecipientTokenInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activation_url"] = attrs["activation_url"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RecipientTokenInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RecipientTokenInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RecipientTokenInfo
// only implements ToObjectValue() and Type().
func (o RecipientTokenInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RecipientTokenInfo) Type(ctx context.Context) attr.Type {
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
type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RetrieveTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RetrieveTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RetrieveTokenRequest
// only implements ToObjectValue() and Type().
func (o RetrieveTokenRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activation_url": o.ActivationUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RetrieveTokenRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activation_url": types.StringType,
		},
	}
}

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearerToken"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.String `tfsdk:"expirationTime"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion types.Int64 `tfsdk:"shareCredentialsVersion"`
}

func (newState *RetrieveTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RetrieveTokenResponse) {
}

func (newState *RetrieveTokenResponse) SyncEffectiveFieldsDuringRead(existingState RetrieveTokenResponse) {
}

func (c RetrieveTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bearerToken"] = attrs["bearerToken"].SetOptional()
	attrs["endpoint"] = attrs["endpoint"].SetOptional()
	attrs["expirationTime"] = attrs["expirationTime"].SetOptional()
	attrs["shareCredentialsVersion"] = attrs["shareCredentialsVersion"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RetrieveTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RetrieveTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RetrieveTokenResponse
// only implements ToObjectValue() and Type().
func (o RetrieveTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RetrieveTokenResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bearerToken":             types.StringType,
			"endpoint":                types.StringType,
			"expirationTime":          types.StringType,
			"shareCredentialsVersion": types.Int64Type,
		},
	}
}

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds types.Int64 `tfsdk:"existing_token_expire_in_seconds"`
	// The name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *RotateRecipientToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan RotateRecipientToken) {
}

func (newState *RotateRecipientToken) SyncEffectiveFieldsDuringRead(existingState RotateRecipientToken) {
}

func (c RotateRecipientToken) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RotateRecipientToken) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RotateRecipientToken
// only implements ToObjectValue() and Type().
func (o RotateRecipientToken) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"existing_token_expire_in_seconds": o.ExistingTokenExpireInSeconds,
			"name":                             o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RotateRecipientToken) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"existing_token_expire_in_seconds": types.Int64Type,
			"name":                             types.StringType,
		},
	}
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`
}

func (newState *SecurablePropertiesKvPairs) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecurablePropertiesKvPairs) {
}

func (newState *SecurablePropertiesKvPairs) SyncEffectiveFieldsDuringRead(existingState SecurablePropertiesKvPairs) {
}

func (c SecurablePropertiesKvPairs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SecurablePropertiesKvPairs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecurablePropertiesKvPairs
// only implements ToObjectValue() and Type().
func (o SecurablePropertiesKvPairs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"properties": o.Properties,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SecurablePropertiesKvPairs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetProperties returns the value of the Properties field in SecurablePropertiesKvPairs as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SecurablePropertiesKvPairs) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in SecurablePropertiesKvPairs.
func (o *SecurablePropertiesKvPairs) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type ShareInfo struct {
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
	Owner types.String `tfsdk:"owner"`
	// Storage Location URL (full path) for the share.
	StorageLocation types.String `tfsdk:"storage_location"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of share updater.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareInfo) {
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringRead(existingState ShareInfo) {
}

func (c ShareInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["object"] = attrs["object"].SetOptional()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
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
func (a ShareInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"object": reflect.TypeOf(SharedDataObject{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ShareInfo
// only implements ToObjectValue() and Type().
func (o ShareInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ShareInfo) Type(ctx context.Context) attr.Type {
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

// GetObjects returns the value of the Objects field in ShareInfo as
// a slice of SharedDataObject values.
// If the field is unknown or null, the boolean return value is false.
func (o *ShareInfo) GetObjects(ctx context.Context) ([]SharedDataObject, bool) {
	if o.Objects.IsNull() || o.Objects.IsUnknown() {
		return nil, false
	}
	var v []SharedDataObject
	d := o.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in ShareInfo.
func (o *ShareInfo) SetObjects(ctx context.Context, v []SharedDataObject) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Objects = types.ListValueMust(t, vs)
}

// Get recipient share permissions
type SharePermissionsRequest struct {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SharePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharePermissionsRequest
// only implements ToObjectValue() and Type().
func (o SharePermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"name":        o.Name,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SharePermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"name":        types.StringType,
			"page_token":  types.StringType,
		},
	}
}

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments"`
	// The share name.
	ShareName types.String `tfsdk:"share_name"`
}

func (newState *ShareToPrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareToPrivilegeAssignment) {
}

func (newState *ShareToPrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState ShareToPrivilegeAssignment) {
}

func (c ShareToPrivilegeAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ShareToPrivilegeAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(PrivilegeAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ShareToPrivilegeAssignment
// only implements ToObjectValue() and Type().
func (o ShareToPrivilegeAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"privilege_assignments": o.PrivilegeAssignments,
			"share_name":            o.ShareName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ShareToPrivilegeAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment{}.Type(ctx),
			},
			"share_name": types.StringType,
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in ShareToPrivilegeAssignment as
// a slice of PrivilegeAssignment values.
// If the field is unknown or null, the boolean return value is false.
func (o *ShareToPrivilegeAssignment) GetPrivilegeAssignments(ctx context.Context) ([]PrivilegeAssignment, bool) {
	if o.PrivilegeAssignments.IsNull() || o.PrivilegeAssignments.IsUnknown() {
		return nil, false
	}
	var v []PrivilegeAssignment
	d := o.PrivilegeAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in ShareToPrivilegeAssignment.
func (o *ShareToPrivilegeAssignment) SetPrivilegeAssignments(ctx context.Context, v []PrivilegeAssignment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrivilegeAssignments = types.ListValueMust(t, vs)
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at"`
	// Username of the sharer.
	AddedBy types.String `tfsdk:"added_by"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled          types.Bool `tfsdk:"cdf_enabled"`
	EffectiveCdfEnabled types.Bool `tfsdk:"effective_cdf_enabled"`
	// A user-provided comment when adding the data object to the share.
	// [Update:OPT]
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
	// A fully qualified name that uniquely identifies a data object.
	//
	// For example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`.
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
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `string_shared_as` name. The `string_shared_as` name must be unique
	// within a share. For notebooks, the new name should be the new notebook
	// file name.
	StringSharedAs types.String `tfsdk:"string_shared_as"`
}

func (newState *SharedDataObject) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObject) {
	newState.EffectiveCdfEnabled = newState.CdfEnabled
	newState.CdfEnabled = plan.CdfEnabled
	newState.EffectiveHistoryDataSharingStatus = newState.HistoryDataSharingStatus
	newState.HistoryDataSharingStatus = plan.HistoryDataSharingStatus
	newState.EffectiveSharedAs = newState.SharedAs
	newState.SharedAs = plan.SharedAs
	newState.EffectiveStartVersion = newState.StartVersion
	newState.StartVersion = plan.StartVersion
}

func (newState *SharedDataObject) SyncEffectiveFieldsDuringRead(existingState SharedDataObject) {
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

func (c SharedDataObject) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SharedDataObject) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partition": reflect.TypeOf(Partition{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharedDataObject
// only implements ToObjectValue() and Type().
func (o SharedDataObject) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SharedDataObject) Type(ctx context.Context) attr.Type {
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

// GetPartitions returns the value of the Partitions field in SharedDataObject as
// a slice of Partition values.
// If the field is unknown or null, the boolean return value is false.
func (o *SharedDataObject) GetPartitions(ctx context.Context) ([]Partition, bool) {
	if o.Partitions.IsNull() || o.Partitions.IsUnknown() {
		return nil, false
	}
	var v []Partition
	d := o.Partitions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPartitions sets the value of the Partitions field in SharedDataObject.
func (o *SharedDataObject) SetPartitions(ctx context.Context, v []Partition) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["partition"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Partitions = types.ListValueMust(t, vs)
}

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	Action types.String `tfsdk:"action"`
	// The data object that is being added, removed, or updated.
	DataObject types.Object `tfsdk:"data_object"`
}

func (newState *SharedDataObjectUpdate) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObjectUpdate) {
}

func (newState *SharedDataObjectUpdate) SyncEffectiveFieldsDuringRead(existingState SharedDataObjectUpdate) {
}

func (c SharedDataObjectUpdate) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["action"] = attrs["action"].SetOptional()
	attrs["data_object"] = attrs["data_object"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SharedDataObjectUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SharedDataObjectUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_object": reflect.TypeOf(SharedDataObject{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SharedDataObjectUpdate
// only implements ToObjectValue() and Type().
func (o SharedDataObjectUpdate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"action":      o.Action,
			"data_object": o.DataObject,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SharedDataObjectUpdate) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action":      types.StringType,
			"data_object": SharedDataObject{}.Type(ctx),
		},
	}
}

// GetDataObject returns the value of the DataObject field in SharedDataObjectUpdate as
// a SharedDataObject value.
// If the field is unknown or null, the boolean return value is false.
func (o *SharedDataObjectUpdate) GetDataObject(ctx context.Context) (SharedDataObject, bool) {
	var e SharedDataObject
	if o.DataObject.IsNull() || o.DataObject.IsUnknown() {
		return e, false
	}
	var v []SharedDataObject
	d := o.DataObject.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataObject sets the value of the DataObject field in SharedDataObjectUpdate.
func (o *SharedDataObjectUpdate) SetDataObject(ctx context.Context, v SharedDataObject) {
	vs := v.ToObjectValue(ctx)
	o.DataObject = vs
}

type UpdatePermissionsResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePermissionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePermissionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePermissionsResponse
// only implements ToObjectValue() and Type().
func (o UpdatePermissionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePermissionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateProvider struct {
	// Description about the provider.
	Comment types.String `tfsdk:"comment"`
	// Name of the provider.
	Name types.String `tfsdk:"-"`
	// New name for the provider.
	NewName types.String `tfsdk:"new_name"`
	// Username of Provider owner.
	Owner types.String `tfsdk:"owner"`
	// This field is required when the __authentication_type__ is **TOKEN** or
	// not provided.
	RecipientProfileStr types.String `tfsdk:"recipient_profile_str"`
}

func (newState *UpdateProvider) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProvider) {
}

func (newState *UpdateProvider) SyncEffectiveFieldsDuringRead(existingState UpdateProvider) {
}

func (c UpdateProvider) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["recipient_profile_str"] = attrs["recipient_profile_str"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProvider.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateProvider) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProvider
// only implements ToObjectValue() and Type().
func (o UpdateProvider) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateProvider) Type(ctx context.Context) attr.Type {
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

type UpdateRecipient struct {
	// Description about the recipient.
	Comment types.String `tfsdk:"comment"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// IP Access List
	IpAccessList types.Object `tfsdk:"ip_access_list"`
	// Name of the recipient.
	Name types.String `tfsdk:"-"`
	// New name for the recipient.
	NewName types.String `tfsdk:"new_name"`
	// Username of the recipient owner.
	Owner types.String `tfsdk:"owner"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs types.Object `tfsdk:"properties_kvpairs"`
}

func (newState *UpdateRecipient) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRecipient) {
}

func (newState *UpdateRecipient) SyncEffectiveFieldsDuringRead(existingState UpdateRecipient) {
}

func (c UpdateRecipient) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties_kvpairs"] = attrs["properties_kvpairs"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRecipient.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRecipient) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list":     reflect.TypeOf(IpAccessList{}),
		"properties_kvpairs": reflect.TypeOf(SecurablePropertiesKvPairs{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRecipient
// only implements ToObjectValue() and Type().
func (o UpdateRecipient) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateRecipient) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":            types.StringType,
			"expiration_time":    types.Int64Type,
			"ip_access_list":     IpAccessList{}.Type(ctx),
			"name":               types.StringType,
			"new_name":           types.StringType,
			"owner":              types.StringType,
			"properties_kvpairs": SecurablePropertiesKvPairs{}.Type(ctx),
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in UpdateRecipient as
// a IpAccessList value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRecipient) GetIpAccessList(ctx context.Context) (IpAccessList, bool) {
	var e IpAccessList
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessList
	d := o.IpAccessList.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in UpdateRecipient.
func (o *UpdateRecipient) SetIpAccessList(ctx context.Context, v IpAccessList) {
	vs := v.ToObjectValue(ctx)
	o.IpAccessList = vs
}

// GetPropertiesKvpairs returns the value of the PropertiesKvpairs field in UpdateRecipient as
// a SecurablePropertiesKvPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRecipient) GetPropertiesKvpairs(ctx context.Context) (SecurablePropertiesKvPairs, bool) {
	var e SecurablePropertiesKvPairs
	if o.PropertiesKvpairs.IsNull() || o.PropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []SecurablePropertiesKvPairs
	d := o.PropertiesKvpairs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPropertiesKvpairs sets the value of the PropertiesKvpairs field in UpdateRecipient.
func (o *UpdateRecipient) SetPropertiesKvpairs(ctx context.Context, v SecurablePropertiesKvPairs) {
	vs := v.ToObjectValue(ctx)
	o.PropertiesKvpairs = vs
}

type UpdateResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse
// only implements ToObjectValue() and Type().
func (o UpdateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// New name for the share.
	NewName types.String `tfsdk:"new_name"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root"`
	// Array of shared data object updates.
	Updates types.List `tfsdk:"updates"`
}

func (newState *UpdateShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateShare) {
}

func (newState *UpdateShare) SyncEffectiveFieldsDuringRead(existingState UpdateShare) {
}

func (c UpdateShare) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()
	attrs["updates"] = attrs["updates"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateShare.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateShare) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"updates": reflect.TypeOf(SharedDataObjectUpdate{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateShare
// only implements ToObjectValue() and Type().
func (o UpdateShare) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateShare) Type(ctx context.Context) attr.Type {
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

// GetUpdates returns the value of the Updates field in UpdateShare as
// a slice of SharedDataObjectUpdate values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateShare) GetUpdates(ctx context.Context) ([]SharedDataObjectUpdate, bool) {
	if o.Updates.IsNull() || o.Updates.IsUnknown() {
		return nil, false
	}
	var v []SharedDataObjectUpdate
	d := o.Updates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdates sets the value of the Updates field in UpdateShare.
func (o *UpdateShare) SetUpdates(ctx context.Context, v []SharedDataObjectUpdate) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["updates"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Updates = types.ListValueMust(t, vs)
}

type UpdateSharePermissions struct {
	// Array of permission changes.
	Changes types.List `tfsdk:"changes"`
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

func (newState *UpdateSharePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSharePermissions) {
}

func (newState *UpdateSharePermissions) SyncEffectiveFieldsDuringRead(existingState UpdateSharePermissions) {
}

func (c UpdateSharePermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["changes"] = attrs["changes"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSharePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateSharePermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"changes": reflect.TypeOf(catalog_tf.PermissionsChange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSharePermissions
// only implements ToObjectValue() and Type().
func (o UpdateSharePermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateSharePermissions) Type(ctx context.Context) attr.Type {
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

// GetChanges returns the value of the Changes field in UpdateSharePermissions as
// a slice of catalog_tf.PermissionsChange values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateSharePermissions) GetChanges(ctx context.Context) ([]catalog_tf.PermissionsChange, bool) {
	if o.Changes.IsNull() || o.Changes.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.PermissionsChange
	d := o.Changes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChanges sets the value of the Changes field in UpdateSharePermissions.
func (o *UpdateSharePermissions) SetChanges(ctx context.Context, v []catalog_tf.PermissionsChange) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["changes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Changes = types.ListValueMust(t, vs)
}
