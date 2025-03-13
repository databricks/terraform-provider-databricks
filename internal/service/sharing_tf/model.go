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
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
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
	// This field is only present when the __authentication_type__ is
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
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs types.Object `tfsdk:"properties_kvpairs"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
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

// Represents a UC dependency.
type DeltaSharingDependency struct {
	// A Function in UC as a dependency.
	Function types.Object `tfsdk:"function"`
	// A Table in UC as a dependency.
	Table types.Object `tfsdk:"table"`
}

func (newState *DeltaSharingDependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaSharingDependency) {
}

func (newState *DeltaSharingDependency) SyncEffectiveFieldsDuringRead(existingState DeltaSharingDependency) {
}

func (c DeltaSharingDependency) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["function"] = attrs["function"].SetOptional()
	attrs["table"] = attrs["table"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSharingDependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeltaSharingDependency) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"function": reflect.TypeOf(DeltaSharingFunctionDependency{}),
		"table":    reflect.TypeOf(DeltaSharingTableDependency{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSharingDependency
// only implements ToObjectValue() and Type().
func (o DeltaSharingDependency) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function": o.Function,
			"table":    o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaSharingDependency) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function": DeltaSharingFunctionDependency{}.Type(ctx),
			"table":    DeltaSharingTableDependency{}.Type(ctx),
		},
	}
}

// GetFunction returns the value of the Function field in DeltaSharingDependency as
// a DeltaSharingFunctionDependency value.
// If the field is unknown or null, the boolean return value is false.
func (o *DeltaSharingDependency) GetFunction(ctx context.Context) (DeltaSharingFunctionDependency, bool) {
	var e DeltaSharingFunctionDependency
	if o.Function.IsNull() || o.Function.IsUnknown() {
		return e, false
	}
	var v []DeltaSharingFunctionDependency
	d := o.Function.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetFunction sets the value of the Function field in DeltaSharingDependency.
func (o *DeltaSharingDependency) SetFunction(ctx context.Context, v DeltaSharingFunctionDependency) {
	vs := v.ToObjectValue(ctx)
	o.Function = vs
}

// GetTable returns the value of the Table field in DeltaSharingDependency as
// a DeltaSharingTableDependency value.
// If the field is unknown or null, the boolean return value is false.
func (o *DeltaSharingDependency) GetTable(ctx context.Context) (DeltaSharingTableDependency, bool) {
	var e DeltaSharingTableDependency
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []DeltaSharingTableDependency
	d := o.Table.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTable sets the value of the Table field in DeltaSharingDependency.
func (o *DeltaSharingDependency) SetTable(ctx context.Context, v DeltaSharingTableDependency) {
	vs := v.ToObjectValue(ctx)
	o.Table = vs
}

// Represents a list of dependencies.
type DeltaSharingDependencyList struct {
	// An array of Dependency.
	Dependencies types.List `tfsdk:"dependencies"`
}

func (newState *DeltaSharingDependencyList) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaSharingDependencyList) {
}

func (newState *DeltaSharingDependencyList) SyncEffectiveFieldsDuringRead(existingState DeltaSharingDependencyList) {
}

func (c DeltaSharingDependencyList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeltaSharingDependencyList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies": reflect.TypeOf(DeltaSharingDependency{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSharingDependencyList
// only implements ToObjectValue() and Type().
func (o DeltaSharingDependencyList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dependencies": o.Dependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaSharingDependencyList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dependencies": basetypes.ListType{
				ElemType: DeltaSharingDependency{}.Type(ctx),
			},
		},
	}
}

// GetDependencies returns the value of the Dependencies field in DeltaSharingDependencyList as
// a slice of DeltaSharingDependency values.
// If the field is unknown or null, the boolean return value is false.
func (o *DeltaSharingDependencyList) GetDependencies(ctx context.Context) ([]DeltaSharingDependency, bool) {
	if o.Dependencies.IsNull() || o.Dependencies.IsUnknown() {
		return nil, false
	}
	var v []DeltaSharingDependency
	d := o.Dependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependencies sets the value of the Dependencies field in DeltaSharingDependencyList.
func (o *DeltaSharingDependencyList) SetDependencies(ctx context.Context, v []DeltaSharingDependency) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Dependencies = types.ListValueMust(t, vs)
}

// A Function in UC as a dependency.
type DeltaSharingFunctionDependency struct {
	FunctionName types.String `tfsdk:"function_name"`

	SchemaName types.String `tfsdk:"schema_name"`
}

func (newState *DeltaSharingFunctionDependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaSharingFunctionDependency) {
}

func (newState *DeltaSharingFunctionDependency) SyncEffectiveFieldsDuringRead(existingState DeltaSharingFunctionDependency) {
}

func (c DeltaSharingFunctionDependency) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeltaSharingFunctionDependency) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSharingFunctionDependency
// only implements ToObjectValue() and Type().
func (o DeltaSharingFunctionDependency) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_name": o.FunctionName,
			"schema_name":   o.SchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaSharingFunctionDependency) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_name": types.StringType,
			"schema_name":   types.StringType,
		},
	}
}

// A Table in UC as a dependency.
type DeltaSharingTableDependency struct {
	SchemaName types.String `tfsdk:"schema_name"`

	TableName types.String `tfsdk:"table_name"`
}

func (newState *DeltaSharingTableDependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaSharingTableDependency) {
}

func (newState *DeltaSharingTableDependency) SyncEffectiveFieldsDuringRead(existingState DeltaSharingTableDependency) {
}

func (c DeltaSharingTableDependency) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeltaSharingTableDependency) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSharingTableDependency
// only implements ToObjectValue() and Type().
func (o DeltaSharingTableDependency) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"schema_name": o.SchemaName,
			"table_name":  o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaSharingTableDependency) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"schema_name": types.StringType,
			"table_name":  types.StringType,
		},
	}
}

type Function struct {
	// The aliass of registered model.
	Aliases types.List `tfsdk:"aliases"`
	// The comment of the function.
	Comment types.String `tfsdk:"comment"`
	// The data type of the function.
	DataType types.String `tfsdk:"data_type"`
	// The dependency list of the function.
	DependencyList types.Object `tfsdk:"dependency_list"`
	// The full data type of the function.
	FullDataType types.String `tfsdk:"full_data_type"`
	// The id of the function.
	Id types.String `tfsdk:"id"`
	// The function parameter information.
	InputParams types.Object `tfsdk:"input_params"`
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

func (newState *Function) SyncEffectiveFieldsDuringCreateOrUpdate(plan Function) {
}

func (newState *Function) SyncEffectiveFieldsDuringRead(existingState Function) {
}

func (c Function) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aliases"] = attrs["aliases"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["data_type"] = attrs["data_type"].SetOptional()
	attrs["dependency_list"] = attrs["dependency_list"].SetOptional()
	attrs["full_data_type"] = attrs["full_data_type"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["input_params"] = attrs["input_params"].SetOptional()
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Function.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Function) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aliases":         reflect.TypeOf(RegisteredModelAlias{}),
		"dependency_list": reflect.TypeOf(DeltaSharingDependencyList{}),
		"input_params":    reflect.TypeOf(FunctionParameterInfos{}),
		"tags":            reflect.TypeOf(catalog_tf.TagKeyValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Function
// only implements ToObjectValue() and Type().
func (o Function) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aliases":            o.Aliases,
			"comment":            o.Comment,
			"data_type":          o.DataType,
			"dependency_list":    o.DependencyList,
			"full_data_type":     o.FullDataType,
			"id":                 o.Id,
			"input_params":       o.InputParams,
			"name":               o.Name,
			"properties":         o.Properties,
			"routine_definition": o.RoutineDefinition,
			"schema":             o.Schema,
			"securable_kind":     o.SecurableKind,
			"share":              o.Share,
			"share_id":           o.ShareId,
			"storage_location":   o.StorageLocation,
			"tags":               o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Function) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aliases": basetypes.ListType{
				ElemType: RegisteredModelAlias{}.Type(ctx),
			},
			"comment":            types.StringType,
			"data_type":          types.StringType,
			"dependency_list":    DeltaSharingDependencyList{}.Type(ctx),
			"full_data_type":     types.StringType,
			"id":                 types.StringType,
			"input_params":       FunctionParameterInfos{}.Type(ctx),
			"name":               types.StringType,
			"properties":         types.StringType,
			"routine_definition": types.StringType,
			"schema":             types.StringType,
			"securable_kind":     types.StringType,
			"share":              types.StringType,
			"share_id":           types.StringType,
			"storage_location":   types.StringType,
			"tags": basetypes.ListType{
				ElemType: catalog_tf.TagKeyValue{}.Type(ctx),
			},
		},
	}
}

// GetAliases returns the value of the Aliases field in Function as
// a slice of RegisteredModelAlias values.
// If the field is unknown or null, the boolean return value is false.
func (o *Function) GetAliases(ctx context.Context) ([]RegisteredModelAlias, bool) {
	if o.Aliases.IsNull() || o.Aliases.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAlias
	d := o.Aliases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAliases sets the value of the Aliases field in Function.
func (o *Function) SetAliases(ctx context.Context, v []RegisteredModelAlias) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Aliases = types.ListValueMust(t, vs)
}

// GetDependencyList returns the value of the DependencyList field in Function as
// a DeltaSharingDependencyList value.
// If the field is unknown or null, the boolean return value is false.
func (o *Function) GetDependencyList(ctx context.Context) (DeltaSharingDependencyList, bool) {
	var e DeltaSharingDependencyList
	if o.DependencyList.IsNull() || o.DependencyList.IsUnknown() {
		return e, false
	}
	var v []DeltaSharingDependencyList
	d := o.DependencyList.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDependencyList sets the value of the DependencyList field in Function.
func (o *Function) SetDependencyList(ctx context.Context, v DeltaSharingDependencyList) {
	vs := v.ToObjectValue(ctx)
	o.DependencyList = vs
}

// GetInputParams returns the value of the InputParams field in Function as
// a FunctionParameterInfos value.
// If the field is unknown or null, the boolean return value is false.
func (o *Function) GetInputParams(ctx context.Context) (FunctionParameterInfos, bool) {
	var e FunctionParameterInfos
	if o.InputParams.IsNull() || o.InputParams.IsUnknown() {
		return e, false
	}
	var v []FunctionParameterInfos
	d := o.InputParams.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetInputParams sets the value of the InputParams field in Function.
func (o *Function) SetInputParams(ctx context.Context, v FunctionParameterInfos) {
	vs := v.ToObjectValue(ctx)
	o.InputParams = vs
}

// GetTags returns the value of the Tags field in Function as
// a slice of catalog_tf.TagKeyValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *Function) GetTags(ctx context.Context) ([]catalog_tf.TagKeyValue, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.TagKeyValue
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Function.
func (o *Function) SetTags(ctx context.Context, v []catalog_tf.TagKeyValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// Represents a parameter of a function. The same message is used for both input
// and output columns.
type FunctionParameterInfo struct {
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

func (newState *FunctionParameterInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionParameterInfo) {
}

func (newState *FunctionParameterInfo) SyncEffectiveFieldsDuringRead(existingState FunctionParameterInfo) {
}

func (c FunctionParameterInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FunctionParameterInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionParameterInfo
// only implements ToObjectValue() and Type().
func (o FunctionParameterInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":            o.Comment,
			"name":               o.Name,
			"parameter_default":  o.ParameterDefault,
			"parameter_mode":     o.ParameterMode,
			"parameter_type":     o.ParameterType,
			"position":           o.Position,
			"type_interval_type": o.TypeIntervalType,
			"type_json":          o.TypeJson,
			"type_name":          o.TypeName,
			"type_precision":     o.TypePrecision,
			"type_scale":         o.TypeScale,
			"type_text":          o.TypeText,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FunctionParameterInfo) Type(ctx context.Context) attr.Type {
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

type FunctionParameterInfos struct {
	// The list of parameters of the function.
	Parameters types.List `tfsdk:"parameters"`
}

func (newState *FunctionParameterInfos) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionParameterInfos) {
}

func (newState *FunctionParameterInfos) SyncEffectiveFieldsDuringRead(existingState FunctionParameterInfos) {
}

func (c FunctionParameterInfos) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FunctionParameterInfos) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(FunctionParameterInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionParameterInfos
// only implements ToObjectValue() and Type().
func (o FunctionParameterInfos) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parameters": o.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FunctionParameterInfos) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parameters": basetypes.ListType{
				ElemType: FunctionParameterInfo{}.Type(ctx),
			},
		},
	}
}

// GetParameters returns the value of the Parameters field in FunctionParameterInfos as
// a slice of FunctionParameterInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *FunctionParameterInfos) GetParameters(ctx context.Context) ([]FunctionParameterInfo, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []FunctionParameterInfo
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in FunctionParameterInfos.
func (o *FunctionParameterInfos) SetParameters(ctx context.Context, v []FunctionParameterInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Parameters = types.ListValueMust(t, vs)
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

type GetSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The privileges assigned to each principal
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments"`
}

func (newState *GetSharePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSharePermissionsResponse) {
}

func (newState *GetSharePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState GetSharePermissionsResponse) {
}

func (c GetSharePermissionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetSharePermissionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(PrivilegeAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSharePermissionsResponse
// only implements ToObjectValue() and Type().
func (o GetSharePermissionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":       o.NextPageToken,
			"privilege_assignments": o.PrivilegeAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSharePermissionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment{}.Type(ctx),
			},
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in GetSharePermissionsResponse as
// a slice of PrivilegeAssignment values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetSharePermissionsResponse) GetPrivilegeAssignments(ctx context.Context) ([]PrivilegeAssignment, bool) {
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

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in GetSharePermissionsResponse.
func (o *GetSharePermissionsResponse) SetPrivilegeAssignments(ctx context.Context, v []PrivilegeAssignment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrivilegeAssignments = types.ListValueMust(t, vs)
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

// List assets by provider share
type ListProviderShareAssetsRequest struct {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProviderShareAssetsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListProviderShareAssetsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderShareAssetsRequest
// only implements ToObjectValue() and Type().
func (o ListProviderShareAssetsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_max_results": o.FunctionMaxResults,
			"notebook_max_results": o.NotebookMaxResults,
			"provider_name":        o.ProviderName,
			"share_name":           o.ShareName,
			"table_max_results":    o.TableMaxResults,
			"volume_max_results":   o.VolumeMaxResults,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProviderShareAssetsRequest) Type(ctx context.Context) attr.Type {
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
type ListProviderShareAssetsResponse struct {
	// The list of functions in the share.
	Functions types.List `tfsdk:"functions"`
	// The list of notebooks in the share.
	Notebooks types.List `tfsdk:"notebooks"`
	// The list of tables in the share.
	Tables types.List `tfsdk:"tables"`
	// The list of volumes in the share.
	Volumes types.List `tfsdk:"volumes"`
}

func (newState *ListProviderShareAssetsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProviderShareAssetsResponse) {
}

func (newState *ListProviderShareAssetsResponse) SyncEffectiveFieldsDuringRead(existingState ListProviderShareAssetsResponse) {
}

func (c ListProviderShareAssetsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["functions"] = attrs["functions"].SetOptional()
	attrs["notebooks"] = attrs["notebooks"].SetOptional()
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
func (a ListProviderShareAssetsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"functions": reflect.TypeOf(Function{}),
		"notebooks": reflect.TypeOf(NotebookFile{}),
		"tables":    reflect.TypeOf(Table{}),
		"volumes":   reflect.TypeOf(Volume{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProviderShareAssetsResponse
// only implements ToObjectValue() and Type().
func (o ListProviderShareAssetsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"functions": o.Functions,
			"notebooks": o.Notebooks,
			"tables":    o.Tables,
			"volumes":   o.Volumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListProviderShareAssetsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"functions": basetypes.ListType{
				ElemType: Function{}.Type(ctx),
			},
			"notebooks": basetypes.ListType{
				ElemType: NotebookFile{}.Type(ctx),
			},
			"tables": basetypes.ListType{
				ElemType: Table{}.Type(ctx),
			},
			"volumes": basetypes.ListType{
				ElemType: Volume{}.Type(ctx),
			},
		},
	}
}

// GetFunctions returns the value of the Functions field in ListProviderShareAssetsResponse as
// a slice of Function values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListProviderShareAssetsResponse) GetFunctions(ctx context.Context) ([]Function, bool) {
	if o.Functions.IsNull() || o.Functions.IsUnknown() {
		return nil, false
	}
	var v []Function
	d := o.Functions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFunctions sets the value of the Functions field in ListProviderShareAssetsResponse.
func (o *ListProviderShareAssetsResponse) SetFunctions(ctx context.Context, v []Function) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["functions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Functions = types.ListValueMust(t, vs)
}

// GetNotebooks returns the value of the Notebooks field in ListProviderShareAssetsResponse as
// a slice of NotebookFile values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListProviderShareAssetsResponse) GetNotebooks(ctx context.Context) ([]NotebookFile, bool) {
	if o.Notebooks.IsNull() || o.Notebooks.IsUnknown() {
		return nil, false
	}
	var v []NotebookFile
	d := o.Notebooks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebooks sets the value of the Notebooks field in ListProviderShareAssetsResponse.
func (o *ListProviderShareAssetsResponse) SetNotebooks(ctx context.Context, v []NotebookFile) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebooks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Notebooks = types.ListValueMust(t, vs)
}

// GetTables returns the value of the Tables field in ListProviderShareAssetsResponse as
// a slice of Table values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListProviderShareAssetsResponse) GetTables(ctx context.Context) ([]Table, bool) {
	if o.Tables.IsNull() || o.Tables.IsUnknown() {
		return nil, false
	}
	var v []Table
	d := o.Tables.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTables sets the value of the Tables field in ListProviderShareAssetsResponse.
func (o *ListProviderShareAssetsResponse) SetTables(ctx context.Context, v []Table) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tables"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tables = types.ListValueMust(t, vs)
}

// GetVolumes returns the value of the Volumes field in ListProviderShareAssetsResponse as
// a slice of Volume values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListProviderShareAssetsResponse) GetVolumes(ctx context.Context) ([]Volume, bool) {
	if o.Volumes.IsNull() || o.Volumes.IsUnknown() {
		return nil, false
	}
	var v []Volume
	d := o.Volumes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolumes sets the value of the Volumes field in ListProviderShareAssetsResponse.
func (o *ListProviderShareAssetsResponse) SetVolumes(ctx context.Context, v []Volume) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["volumes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Volumes = types.ListValueMust(t, vs)
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

type NotebookFile struct {
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

func (newState *NotebookFile) SyncEffectiveFieldsDuringCreateOrUpdate(plan NotebookFile) {
}

func (newState *NotebookFile) SyncEffectiveFieldsDuringRead(existingState NotebookFile) {
}

func (c NotebookFile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NotebookFile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(catalog_tf.TagKeyValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookFile
// only implements ToObjectValue() and Type().
func (o NotebookFile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":  o.Comment,
			"id":       o.Id,
			"name":     o.Name,
			"share":    o.Share,
			"share_id": o.ShareId,
			"tags":     o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NotebookFile) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":  types.StringType,
			"id":       types.StringType,
			"name":     types.StringType,
			"share":    types.StringType,
			"share_id": types.StringType,
			"tags": basetypes.ListType{
				ElemType: catalog_tf.TagKeyValue{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in NotebookFile as
// a slice of catalog_tf.TagKeyValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *NotebookFile) GetTags(ctx context.Context) ([]catalog_tf.TagKeyValue, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.TagKeyValue
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in NotebookFile.
func (o *NotebookFile) SetTags(ctx context.Context, v []catalog_tf.TagKeyValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
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

type PermissionsChange struct {
	// The set of privileges to add.
	Add types.List `tfsdk:"add"`
	// The principal whose privileges we are changing.
	Principal types.String `tfsdk:"principal"`
	// The set of privileges to remove.
	Remove types.List `tfsdk:"remove"`
}

func (newState *PermissionsChange) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsChange) {
}

func (newState *PermissionsChange) SyncEffectiveFieldsDuringRead(existingState PermissionsChange) {
}

func (c PermissionsChange) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PermissionsChange) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"add":    reflect.TypeOf(types.String{}),
		"remove": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionsChange
// only implements ToObjectValue() and Type().
func (o PermissionsChange) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"add":       o.Add,
			"principal": o.Principal,
			"remove":    o.Remove,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionsChange) Type(ctx context.Context) attr.Type {
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

// GetAdd returns the value of the Add field in PermissionsChange as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionsChange) GetAdd(ctx context.Context) ([]types.String, bool) {
	if o.Add.IsNull() || o.Add.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Add.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAdd sets the value of the Add field in PermissionsChange.
func (o *PermissionsChange) SetAdd(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["add"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Add = types.ListValueMust(t, vs)
}

// GetRemove returns the value of the Remove field in PermissionsChange as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionsChange) GetRemove(ctx context.Context) ([]types.String, bool) {
	if o.Remove.IsNull() || o.Remove.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Remove.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRemove sets the value of the Remove field in PermissionsChange.
func (o *PermissionsChange) SetRemove(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["remove"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Remove = types.ListValueMust(t, vs)
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
	RecipientProfile types.Object `tfsdk:"recipient_profile"`
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

func (newState *ProviderInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderInfo) {
}

func (newState *ProviderInfo) SyncEffectiveFieldsDuringRead(existingState ProviderInfo) {
}

func (c ProviderInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	IpAccessList types.Object `tfsdk:"ip_access_list"`
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
	PropertiesKvpairs types.Object `tfsdk:"properties_kvpairs"`
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

func (newState *RecipientInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientInfo) {
}

func (newState *RecipientInfo) SyncEffectiveFieldsDuringRead(existingState RecipientInfo) {
}

func (c RecipientInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties_kvpairs"] = attrs["properties_kvpairs"].SetOptional()
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
			"expiration_time":                    o.ExpirationTime,
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
			"expiration_time":                    types.Int64Type,
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

func (newState *RecipientTokenInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientTokenInfo) {
}

func (newState *RecipientTokenInfo) SyncEffectiveFieldsDuringRead(existingState RecipientTokenInfo) {
}

func (c RecipientTokenInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type RegisteredModelAlias struct {
	// Name of the alias.
	AliasName types.String `tfsdk:"alias_name"`
	// Numeric model version that alias will reference.
	VersionNum types.Int64 `tfsdk:"version_num"`
}

func (newState *RegisteredModelAlias) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelAlias) {
}

func (newState *RegisteredModelAlias) SyncEffectiveFieldsDuringRead(existingState RegisteredModelAlias) {
}

func (c RegisteredModelAlias) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelAlias) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAlias
// only implements ToObjectValue() and Type().
func (o RegisteredModelAlias) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias_name":  o.AliasName,
			"version_num": o.VersionNum,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelAlias) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias_name":  types.StringType,
			"version_num": types.Int64Type,
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
	attrs["bearerToken"] = attrs["bearerToken"].SetComputed()
	attrs["endpoint"] = attrs["endpoint"].SetComputed()
	attrs["expirationTime"] = attrs["expirationTime"].SetComputed()
	attrs["shareCredentialsVersion"] = attrs["shareCredentialsVersion"].SetComputed()

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
	// The name of the Recipient.
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

func (newState *ShareInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareInfo) {
	newState.EffectiveOwner = newState.Owner
	newState.Owner = plan.Owner
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringRead(existingState ShareInfo) {
	newState.EffectiveOwner = existingState.EffectiveOwner
	if existingState.EffectiveOwner.ValueString() == newState.Owner.ValueString() {
		newState.Owner = existingState.Owner
	}
}

func (c ShareInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"effective_owner":  o.EffectiveOwner,
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
			"effective_owner":  types.StringType,
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
	// A user-provided new name for the shared object within the share. If this
	// new name is not not provided, the object's original name will be used as
	// the `string_shared_as` name. The `string_shared_as` name must be unique
	// for objects of the same type within a Share. For notebooks, the new name
	// should be the new notebook file name.
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

type Table struct {
	// The comment of the table.
	Comment types.String `tfsdk:"comment"`
	// The id of the table.
	Id types.String `tfsdk:"id"`
	// Internal information for D2D sharing that should not be disclosed to
	// external users.
	InternalAttributes types.Object `tfsdk:"internal_attributes"`
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

func (newState *Table) SyncEffectiveFieldsDuringCreateOrUpdate(plan Table) {
}

func (newState *Table) SyncEffectiveFieldsDuringRead(existingState Table) {
}

func (c Table) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["internal_attributes"] = attrs["internal_attributes"].SetOptional()
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
func (a Table) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"internal_attributes": reflect.TypeOf(TableInternalAttributes{}),
		"tags":                reflect.TypeOf(catalog_tf.TagKeyValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Table
// only implements ToObjectValue() and Type().
func (o Table) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":                 o.Comment,
			"id":                      o.Id,
			"internal_attributes":     o.InternalAttributes,
			"materialized_table_name": o.MaterializedTableName,
			"name":                    o.Name,
			"schema":                  o.Schema,
			"share":                   o.Share,
			"share_id":                o.ShareId,
			"tags":                    o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Table) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":                 types.StringType,
			"id":                      types.StringType,
			"internal_attributes":     TableInternalAttributes{}.Type(ctx),
			"materialized_table_name": types.StringType,
			"name":                    types.StringType,
			"schema":                  types.StringType,
			"share":                   types.StringType,
			"share_id":                types.StringType,
			"tags": basetypes.ListType{
				ElemType: catalog_tf.TagKeyValue{}.Type(ctx),
			},
		},
	}
}

// GetInternalAttributes returns the value of the InternalAttributes field in Table as
// a TableInternalAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *Table) GetInternalAttributes(ctx context.Context) (TableInternalAttributes, bool) {
	var e TableInternalAttributes
	if o.InternalAttributes.IsNull() || o.InternalAttributes.IsUnknown() {
		return e, false
	}
	var v []TableInternalAttributes
	d := o.InternalAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetInternalAttributes sets the value of the InternalAttributes field in Table.
func (o *Table) SetInternalAttributes(ctx context.Context, v TableInternalAttributes) {
	vs := v.ToObjectValue(ctx)
	o.InternalAttributes = vs
}

// GetTags returns the value of the Tags field in Table as
// a slice of catalog_tf.TagKeyValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *Table) GetTags(ctx context.Context) ([]catalog_tf.TagKeyValue, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.TagKeyValue
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Table.
func (o *Table) SetTags(ctx context.Context, v []catalog_tf.TagKeyValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// Internal information for D2D sharing that should not be disclosed to external
// users.
type TableInternalAttributes struct {
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

func (newState *TableInternalAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableInternalAttributes) {
}

func (newState *TableInternalAttributes) SyncEffectiveFieldsDuringRead(existingState TableInternalAttributes) {
}

func (c TableInternalAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TableInternalAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableInternalAttributes
// only implements ToObjectValue() and Type().
func (o TableInternalAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent_storage_location": o.ParentStorageLocation,
			"storage_location":        o.StorageLocation,
			"type":                    o.Type_,
			"view_definition":         o.ViewDefinition,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableInternalAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent_storage_location": types.StringType,
			"storage_location":        types.StringType,
			"type":                    types.StringType,
			"view_definition":         types.StringType,
		},
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
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
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
	// New name for the recipient. .
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

type UpdateShare struct {
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

func (newState *UpdateShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateShare) {
	newState.EffectiveOwner = newState.Owner
	newState.Owner = plan.Owner
}

func (newState *UpdateShare) SyncEffectiveFieldsDuringRead(existingState UpdateShare) {
	newState.EffectiveOwner = existingState.EffectiveOwner
	if existingState.EffectiveOwner.ValueString() == newState.Owner.ValueString() {
		newState.Owner = existingState.Owner
	}
}

func (c UpdateShare) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["effective_owner"] = attrs["effective_owner"].SetComputed()
	attrs["owner"] = attrs["owner"].SetOptional()
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
			"comment":         o.Comment,
			"name":            o.Name,
			"new_name":        o.NewName,
			"owner":           o.Owner,
			"effective_owner": o.EffectiveOwner,
			"storage_root":    o.StorageRoot,
			"updates":         o.Updates,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateShare) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":         types.StringType,
			"name":            types.StringType,
			"new_name":        types.StringType,
			"owner":           types.StringType,
			"effective_owner": types.StringType,
			"storage_root":    types.StringType,
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
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (newState *UpdateSharePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSharePermissions) {
}

func (newState *UpdateSharePermissions) SyncEffectiveFieldsDuringRead(existingState UpdateSharePermissions) {
}

func (c UpdateSharePermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["changes"] = attrs["changes"].SetOptional()
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
func (a UpdateSharePermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"changes": reflect.TypeOf(PermissionsChange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSharePermissions
// only implements ToObjectValue() and Type().
func (o UpdateSharePermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"changes": o.Changes,
			"name":    o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateSharePermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"changes": basetypes.ListType{
				ElemType: PermissionsChange{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetChanges returns the value of the Changes field in UpdateSharePermissions as
// a slice of PermissionsChange values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateSharePermissions) GetChanges(ctx context.Context) ([]PermissionsChange, bool) {
	if o.Changes.IsNull() || o.Changes.IsUnknown() {
		return nil, false
	}
	var v []PermissionsChange
	d := o.Changes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChanges sets the value of the Changes field in UpdateSharePermissions.
func (o *UpdateSharePermissions) SetChanges(ctx context.Context, v []PermissionsChange) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["changes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Changes = types.ListValueMust(t, vs)
}

type UpdateSharePermissionsResponse struct {
	// The privileges assigned to each principal
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments"`
}

func (newState *UpdateSharePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSharePermissionsResponse) {
}

func (newState *UpdateSharePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState UpdateSharePermissionsResponse) {
}

func (c UpdateSharePermissionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateSharePermissionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(PrivilegeAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSharePermissionsResponse
// only implements ToObjectValue() and Type().
func (o UpdateSharePermissionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"privilege_assignments": o.PrivilegeAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateSharePermissionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment{}.Type(ctx),
			},
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in UpdateSharePermissionsResponse as
// a slice of PrivilegeAssignment values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateSharePermissionsResponse) GetPrivilegeAssignments(ctx context.Context) ([]PrivilegeAssignment, bool) {
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

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in UpdateSharePermissionsResponse.
func (o *UpdateSharePermissionsResponse) SetPrivilegeAssignments(ctx context.Context, v []PrivilegeAssignment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrivilegeAssignments = types.ListValueMust(t, vs)
}

type Volume struct {
	// The comment of the volume.
	Comment types.String `tfsdk:"comment"`
	// This id maps to the shared_volume_id in database Recipient needs
	// shared_volume_id for recon to check if this volume is already in
	// recipient's DB or not.
	Id types.String `tfsdk:"id"`
	// Internal attributes for D2D sharing that should not be disclosed to
	// external users.
	InternalAttributes types.Object `tfsdk:"internal_attributes"`
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

func (newState *Volume) SyncEffectiveFieldsDuringCreateOrUpdate(plan Volume) {
}

func (newState *Volume) SyncEffectiveFieldsDuringRead(existingState Volume) {
}

func (c Volume) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["internal_attributes"] = attrs["internal_attributes"].SetOptional()
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
func (a Volume) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"internal_attributes": reflect.TypeOf(VolumeInternalAttributes{}),
		"tags":                reflect.TypeOf(catalog_tf.TagKeyValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Volume
// only implements ToObjectValue() and Type().
func (o Volume) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":             o.Comment,
			"id":                  o.Id,
			"internal_attributes": o.InternalAttributes,
			"name":                o.Name,
			"schema":              o.Schema,
			"share":               o.Share,
			"share_id":            o.ShareId,
			"tags":                o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Volume) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":             types.StringType,
			"id":                  types.StringType,
			"internal_attributes": VolumeInternalAttributes{}.Type(ctx),
			"name":                types.StringType,
			"schema":              types.StringType,
			"share":               types.StringType,
			"share_id":            types.StringType,
			"tags": basetypes.ListType{
				ElemType: catalog_tf.TagKeyValue{}.Type(ctx),
			},
		},
	}
}

// GetInternalAttributes returns the value of the InternalAttributes field in Volume as
// a VolumeInternalAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *Volume) GetInternalAttributes(ctx context.Context) (VolumeInternalAttributes, bool) {
	var e VolumeInternalAttributes
	if o.InternalAttributes.IsNull() || o.InternalAttributes.IsUnknown() {
		return e, false
	}
	var v []VolumeInternalAttributes
	d := o.InternalAttributes.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetInternalAttributes sets the value of the InternalAttributes field in Volume.
func (o *Volume) SetInternalAttributes(ctx context.Context, v VolumeInternalAttributes) {
	vs := v.ToObjectValue(ctx)
	o.InternalAttributes = vs
}

// GetTags returns the value of the Tags field in Volume as
// a slice of catalog_tf.TagKeyValue values.
// If the field is unknown or null, the boolean return value is false.
func (o *Volume) GetTags(ctx context.Context) ([]catalog_tf.TagKeyValue, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.TagKeyValue
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Volume.
func (o *Volume) SetTags(ctx context.Context, v []catalog_tf.TagKeyValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// Internal information for D2D sharing that should not be disclosed to external
// users.
type VolumeInternalAttributes struct {
	// The cloud storage location of the volume
	StorageLocation types.String `tfsdk:"storage_location"`
	// The type of the shared volume.
	Type_ types.String `tfsdk:"type"`
}

func (newState *VolumeInternalAttributes) SyncEffectiveFieldsDuringCreateOrUpdate(plan VolumeInternalAttributes) {
}

func (newState *VolumeInternalAttributes) SyncEffectiveFieldsDuringRead(existingState VolumeInternalAttributes) {
}

func (c VolumeInternalAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a VolumeInternalAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VolumeInternalAttributes
// only implements ToObjectValue() and Type().
func (o VolumeInternalAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_location": o.StorageLocation,
			"type":             o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o VolumeInternalAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_location": types.StringType,
			"type":             types.StringType,
		},
	}
}
