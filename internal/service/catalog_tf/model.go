// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package catalog_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AccountsCreateMetastore struct {
	MetastoreInfo types.Object `tfsdk:"metastore_info"`
}

func (newState *AccountsCreateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateMetastore) {
}

func (newState *AccountsCreateMetastore) SyncEffectiveFieldsDuringRead(existingState AccountsCreateMetastore) {
}

func (c AccountsCreateMetastore) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metastore_info"] = attrs["metastore_info"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsCreateMetastore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsCreateMetastore) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_info": reflect.TypeOf(CreateMetastore{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsCreateMetastore
// only implements ToObjectValue() and Type().
func (o AccountsCreateMetastore) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_info": o.MetastoreInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsCreateMetastore) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_info": CreateMetastore{}.Type(ctx),
		},
	}
}

// GetMetastoreInfo returns the value of the MetastoreInfo field in AccountsCreateMetastore as
// a CreateMetastore value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsCreateMetastore) GetMetastoreInfo(ctx context.Context) (CreateMetastore, bool) {
	var e CreateMetastore
	if o.MetastoreInfo.IsNull() || o.MetastoreInfo.IsUnknown() {
		return e, false
	}
	var v []CreateMetastore
	d := o.MetastoreInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetMetastoreInfo sets the value of the MetastoreInfo field in AccountsCreateMetastore.
func (o *AccountsCreateMetastore) SetMetastoreInfo(ctx context.Context, v CreateMetastore) {
	vs := v.ToObjectValue(ctx)
	o.MetastoreInfo = vs
}

type AccountsCreateMetastoreAssignment struct {
	MetastoreAssignment types.Object `tfsdk:"metastore_assignment"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *AccountsCreateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateMetastoreAssignment) {
}

func (newState *AccountsCreateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState AccountsCreateMetastoreAssignment) {
}

func (c AccountsCreateMetastoreAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metastore_assignment"] = attrs["metastore_assignment"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsCreateMetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsCreateMetastoreAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_assignment": reflect.TypeOf(CreateMetastoreAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsCreateMetastoreAssignment
// only implements ToObjectValue() and Type().
func (o AccountsCreateMetastoreAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_assignment": o.MetastoreAssignment,
			"metastore_id":         o.MetastoreId,
			"workspace_id":         o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsCreateMetastoreAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_assignment": CreateMetastoreAssignment{}.Type(ctx),
			"metastore_id":         types.StringType,
			"workspace_id":         types.Int64Type,
		},
	}
}

// GetMetastoreAssignment returns the value of the MetastoreAssignment field in AccountsCreateMetastoreAssignment as
// a CreateMetastoreAssignment value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsCreateMetastoreAssignment) GetMetastoreAssignment(ctx context.Context) (CreateMetastoreAssignment, bool) {
	var e CreateMetastoreAssignment
	if o.MetastoreAssignment.IsNull() || o.MetastoreAssignment.IsUnknown() {
		return e, false
	}
	var v []CreateMetastoreAssignment
	d := o.MetastoreAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetMetastoreAssignment sets the value of the MetastoreAssignment field in AccountsCreateMetastoreAssignment.
func (o *AccountsCreateMetastoreAssignment) SetMetastoreAssignment(ctx context.Context, v CreateMetastoreAssignment) {
	vs := v.ToObjectValue(ctx)
	o.MetastoreAssignment = vs
}

type AccountsCreateStorageCredential struct {
	CredentialInfo types.Object `tfsdk:"credential_info"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *AccountsCreateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateStorageCredential) {
}

func (newState *AccountsCreateStorageCredential) SyncEffectiveFieldsDuringRead(existingState AccountsCreateStorageCredential) {
}

func (c AccountsCreateStorageCredential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_info"] = attrs["credential_info"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsCreateStorageCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsCreateStorageCredential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credential_info": reflect.TypeOf(CreateStorageCredential{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsCreateStorageCredential
// only implements ToObjectValue() and Type().
func (o AccountsCreateStorageCredential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_info": o.CredentialInfo,
			"metastore_id":    o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsCreateStorageCredential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_info": CreateStorageCredential{}.Type(ctx),
			"metastore_id":    types.StringType,
		},
	}
}

// GetCredentialInfo returns the value of the CredentialInfo field in AccountsCreateStorageCredential as
// a CreateStorageCredential value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsCreateStorageCredential) GetCredentialInfo(ctx context.Context) (CreateStorageCredential, bool) {
	var e CreateStorageCredential
	if o.CredentialInfo.IsNull() || o.CredentialInfo.IsUnknown() {
		return e, false
	}
	var v []CreateStorageCredential
	d := o.CredentialInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetCredentialInfo sets the value of the CredentialInfo field in AccountsCreateStorageCredential.
func (o *AccountsCreateStorageCredential) SetCredentialInfo(ctx context.Context, v CreateStorageCredential) {
	vs := v.ToObjectValue(ctx)
	o.CredentialInfo = vs
}

type AccountsMetastoreAssignment struct {
	MetastoreAssignment types.Object `tfsdk:"metastore_assignment"`
}

func (newState *AccountsMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsMetastoreAssignment) {
}

func (newState *AccountsMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState AccountsMetastoreAssignment) {
}

func (c AccountsMetastoreAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metastore_assignment"] = attrs["metastore_assignment"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsMetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsMetastoreAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_assignment": reflect.TypeOf(MetastoreAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsMetastoreAssignment
// only implements ToObjectValue() and Type().
func (o AccountsMetastoreAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_assignment": o.MetastoreAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsMetastoreAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_assignment": MetastoreAssignment{}.Type(ctx),
		},
	}
}

// GetMetastoreAssignment returns the value of the MetastoreAssignment field in AccountsMetastoreAssignment as
// a MetastoreAssignment value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsMetastoreAssignment) GetMetastoreAssignment(ctx context.Context) (MetastoreAssignment, bool) {
	var e MetastoreAssignment
	if o.MetastoreAssignment.IsNull() || o.MetastoreAssignment.IsUnknown() {
		return e, false
	}
	var v []MetastoreAssignment
	d := o.MetastoreAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetMetastoreAssignment sets the value of the MetastoreAssignment field in AccountsMetastoreAssignment.
func (o *AccountsMetastoreAssignment) SetMetastoreAssignment(ctx context.Context, v MetastoreAssignment) {
	vs := v.ToObjectValue(ctx)
	o.MetastoreAssignment = vs
}

type AccountsMetastoreInfo struct {
	MetastoreInfo types.Object `tfsdk:"metastore_info"`
}

func (newState *AccountsMetastoreInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsMetastoreInfo) {
}

func (newState *AccountsMetastoreInfo) SyncEffectiveFieldsDuringRead(existingState AccountsMetastoreInfo) {
}

func (c AccountsMetastoreInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metastore_info"] = attrs["metastore_info"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsMetastoreInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsMetastoreInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_info": reflect.TypeOf(MetastoreInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsMetastoreInfo
// only implements ToObjectValue() and Type().
func (o AccountsMetastoreInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_info": o.MetastoreInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsMetastoreInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_info": MetastoreInfo{}.Type(ctx),
		},
	}
}

// GetMetastoreInfo returns the value of the MetastoreInfo field in AccountsMetastoreInfo as
// a MetastoreInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsMetastoreInfo) GetMetastoreInfo(ctx context.Context) (MetastoreInfo, bool) {
	var e MetastoreInfo
	if o.MetastoreInfo.IsNull() || o.MetastoreInfo.IsUnknown() {
		return e, false
	}
	var v []MetastoreInfo
	d := o.MetastoreInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetMetastoreInfo sets the value of the MetastoreInfo field in AccountsMetastoreInfo.
func (o *AccountsMetastoreInfo) SetMetastoreInfo(ctx context.Context, v MetastoreInfo) {
	vs := v.ToObjectValue(ctx)
	o.MetastoreInfo = vs
}

type AccountsStorageCredentialInfo struct {
	CredentialInfo types.Object `tfsdk:"credential_info"`
}

func (newState *AccountsStorageCredentialInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsStorageCredentialInfo) {
}

func (newState *AccountsStorageCredentialInfo) SyncEffectiveFieldsDuringRead(existingState AccountsStorageCredentialInfo) {
}

func (c AccountsStorageCredentialInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_info"] = attrs["credential_info"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsStorageCredentialInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsStorageCredentialInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credential_info": reflect.TypeOf(StorageCredentialInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsStorageCredentialInfo
// only implements ToObjectValue() and Type().
func (o AccountsStorageCredentialInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_info": o.CredentialInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsStorageCredentialInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_info": StorageCredentialInfo{}.Type(ctx),
		},
	}
}

// GetCredentialInfo returns the value of the CredentialInfo field in AccountsStorageCredentialInfo as
// a StorageCredentialInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsStorageCredentialInfo) GetCredentialInfo(ctx context.Context) (StorageCredentialInfo, bool) {
	var e StorageCredentialInfo
	if o.CredentialInfo.IsNull() || o.CredentialInfo.IsUnknown() {
		return e, false
	}
	var v []StorageCredentialInfo
	d := o.CredentialInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetCredentialInfo sets the value of the CredentialInfo field in AccountsStorageCredentialInfo.
func (o *AccountsStorageCredentialInfo) SetCredentialInfo(ctx context.Context, v StorageCredentialInfo) {
	vs := v.ToObjectValue(ctx)
	o.CredentialInfo = vs
}

type AccountsUpdateMetastore struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`

	MetastoreInfo types.Object `tfsdk:"metastore_info"`
}

func (newState *AccountsUpdateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateMetastore) {
}

func (newState *AccountsUpdateMetastore) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateMetastore) {
}

func (c AccountsUpdateMetastore) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metastore_id"] = attrs["metastore_id"].SetRequired()
	attrs["metastore_info"] = attrs["metastore_info"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsUpdateMetastore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsUpdateMetastore) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_info": reflect.TypeOf(UpdateMetastore{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsUpdateMetastore
// only implements ToObjectValue() and Type().
func (o AccountsUpdateMetastore) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id":   o.MetastoreId,
			"metastore_info": o.MetastoreInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsUpdateMetastore) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id":   types.StringType,
			"metastore_info": UpdateMetastore{}.Type(ctx),
		},
	}
}

// GetMetastoreInfo returns the value of the MetastoreInfo field in AccountsUpdateMetastore as
// a UpdateMetastore value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsUpdateMetastore) GetMetastoreInfo(ctx context.Context) (UpdateMetastore, bool) {
	var e UpdateMetastore
	if o.MetastoreInfo.IsNull() || o.MetastoreInfo.IsUnknown() {
		return e, false
	}
	var v []UpdateMetastore
	d := o.MetastoreInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetMetastoreInfo sets the value of the MetastoreInfo field in AccountsUpdateMetastore.
func (o *AccountsUpdateMetastore) SetMetastoreInfo(ctx context.Context, v UpdateMetastore) {
	vs := v.ToObjectValue(ctx)
	o.MetastoreInfo = vs
}

type AccountsUpdateMetastoreAssignment struct {
	MetastoreAssignment types.Object `tfsdk:"metastore_assignment"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *AccountsUpdateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateMetastoreAssignment) {
}

func (newState *AccountsUpdateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateMetastoreAssignment) {
}

func (c AccountsUpdateMetastoreAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metastore_assignment"] = attrs["metastore_assignment"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsUpdateMetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsUpdateMetastoreAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_assignment": reflect.TypeOf(UpdateMetastoreAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsUpdateMetastoreAssignment
// only implements ToObjectValue() and Type().
func (o AccountsUpdateMetastoreAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_assignment": o.MetastoreAssignment,
			"metastore_id":         o.MetastoreId,
			"workspace_id":         o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsUpdateMetastoreAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_assignment": UpdateMetastoreAssignment{}.Type(ctx),
			"metastore_id":         types.StringType,
			"workspace_id":         types.Int64Type,
		},
	}
}

// GetMetastoreAssignment returns the value of the MetastoreAssignment field in AccountsUpdateMetastoreAssignment as
// a UpdateMetastoreAssignment value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsUpdateMetastoreAssignment) GetMetastoreAssignment(ctx context.Context) (UpdateMetastoreAssignment, bool) {
	var e UpdateMetastoreAssignment
	if o.MetastoreAssignment.IsNull() || o.MetastoreAssignment.IsUnknown() {
		return e, false
	}
	var v []UpdateMetastoreAssignment
	d := o.MetastoreAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetMetastoreAssignment sets the value of the MetastoreAssignment field in AccountsUpdateMetastoreAssignment.
func (o *AccountsUpdateMetastoreAssignment) SetMetastoreAssignment(ctx context.Context, v UpdateMetastoreAssignment) {
	vs := v.ToObjectValue(ctx)
	o.MetastoreAssignment = vs
}

type AccountsUpdateStorageCredential struct {
	CredentialInfo types.Object `tfsdk:"credential_info"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Name of the storage credential.
	StorageCredentialName types.String `tfsdk:"-"`
}

func (newState *AccountsUpdateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateStorageCredential) {
}

func (newState *AccountsUpdateStorageCredential) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateStorageCredential) {
}

func (c AccountsUpdateStorageCredential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_info"] = attrs["credential_info"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetRequired()
	attrs["storage_credential_name"] = attrs["storage_credential_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsUpdateStorageCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsUpdateStorageCredential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credential_info": reflect.TypeOf(UpdateStorageCredential{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsUpdateStorageCredential
// only implements ToObjectValue() and Type().
func (o AccountsUpdateStorageCredential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_info":         o.CredentialInfo,
			"metastore_id":            o.MetastoreId,
			"storage_credential_name": o.StorageCredentialName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsUpdateStorageCredential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_info":         UpdateStorageCredential{}.Type(ctx),
			"metastore_id":            types.StringType,
			"storage_credential_name": types.StringType,
		},
	}
}

// GetCredentialInfo returns the value of the CredentialInfo field in AccountsUpdateStorageCredential as
// a UpdateStorageCredential value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsUpdateStorageCredential) GetCredentialInfo(ctx context.Context) (UpdateStorageCredential, bool) {
	var e UpdateStorageCredential
	if o.CredentialInfo.IsNull() || o.CredentialInfo.IsUnknown() {
		return e, false
	}
	var v []UpdateStorageCredential
	d := o.CredentialInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetCredentialInfo sets the value of the CredentialInfo field in AccountsUpdateStorageCredential.
func (o *AccountsUpdateStorageCredential) SetCredentialInfo(ctx context.Context, v UpdateStorageCredential) {
	vs := v.ToObjectValue(ctx)
	o.CredentialInfo = vs
}

type ArtifactAllowlistInfo struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers types.List `tfsdk:"artifact_matchers"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of the user who set the artifact allowlist.
	CreatedBy types.String `tfsdk:"created_by"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
}

func (newState *ArtifactAllowlistInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ArtifactAllowlistInfo) {
}

func (newState *ArtifactAllowlistInfo) SyncEffectiveFieldsDuringRead(existingState ArtifactAllowlistInfo) {
}

func (c ArtifactAllowlistInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["artifact_matchers"] = attrs["artifact_matchers"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ArtifactAllowlistInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ArtifactAllowlistInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"artifact_matchers": reflect.TypeOf(ArtifactMatcher{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ArtifactAllowlistInfo
// only implements ToObjectValue() and Type().
func (o ArtifactAllowlistInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_matchers": o.ArtifactMatchers,
			"created_at":        o.CreatedAt,
			"created_by":        o.CreatedBy,
			"metastore_id":      o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ArtifactAllowlistInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_matchers": basetypes.ListType{
				ElemType: ArtifactMatcher{}.Type(ctx),
			},
			"created_at":   types.Int64Type,
			"created_by":   types.StringType,
			"metastore_id": types.StringType,
		},
	}
}

// GetArtifactMatchers returns the value of the ArtifactMatchers field in ArtifactAllowlistInfo as
// a slice of ArtifactMatcher values.
// If the field is unknown or null, the boolean return value is false.
func (o *ArtifactAllowlistInfo) GetArtifactMatchers(ctx context.Context) ([]ArtifactMatcher, bool) {
	if o.ArtifactMatchers.IsNull() || o.ArtifactMatchers.IsUnknown() {
		return nil, false
	}
	var v []ArtifactMatcher
	d := o.ArtifactMatchers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetArtifactMatchers sets the value of the ArtifactMatchers field in ArtifactAllowlistInfo.
func (o *ArtifactAllowlistInfo) SetArtifactMatchers(ctx context.Context, v []ArtifactMatcher) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["artifact_matchers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ArtifactMatchers = types.ListValueMust(t, vs)
}

type ArtifactMatcher struct {
	// The artifact path or maven coordinate
	Artifact types.String `tfsdk:"artifact"`
	// The pattern matching type of the artifact
	MatchType types.String `tfsdk:"match_type"`
}

func (newState *ArtifactMatcher) SyncEffectiveFieldsDuringCreateOrUpdate(plan ArtifactMatcher) {
}

func (newState *ArtifactMatcher) SyncEffectiveFieldsDuringRead(existingState ArtifactMatcher) {
}

func (c ArtifactMatcher) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["artifact"] = attrs["artifact"].SetRequired()
	attrs["match_type"] = attrs["match_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ArtifactMatcher.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ArtifactMatcher) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ArtifactMatcher
// only implements ToObjectValue() and Type().
func (o ArtifactMatcher) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact":   o.Artifact,
			"match_type": o.MatchType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ArtifactMatcher) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact":   types.StringType,
			"match_type": types.StringType,
		},
	}
}

type AssignResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AssignResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AssignResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AssignResponse
// only implements ToObjectValue() and Type().
func (o AssignResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o AssignResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// AWS temporary credentials for API authentication. Read more at
// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
type AwsCredentials struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId types.String `tfsdk:"access_key_id"`
	// The Amazon Resource Name (ARN) of the S3 access point for temporary
	// credentials related the external location.
	AccessPoint types.String `tfsdk:"access_point"`
	// The secret access key that can be used to sign AWS API requests.
	SecretAccessKey types.String `tfsdk:"secret_access_key"`
	// The token that users must pass to AWS API to use the temporary
	// credentials.
	SessionToken types.String `tfsdk:"session_token"`
}

func (newState *AwsCredentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsCredentials) {
}

func (newState *AwsCredentials) SyncEffectiveFieldsDuringRead(existingState AwsCredentials) {
}

func (c AwsCredentials) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_key_id"] = attrs["access_key_id"].SetOptional()
	attrs["access_point"] = attrs["access_point"].SetOptional()
	attrs["secret_access_key"] = attrs["secret_access_key"].SetOptional()
	attrs["session_token"] = attrs["session_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsCredentials) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsCredentials
// only implements ToObjectValue() and Type().
func (o AwsCredentials) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_key_id":     o.AccessKeyId,
			"access_point":      o.AccessPoint,
			"secret_access_key": o.SecretAccessKey,
			"session_token":     o.SessionToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsCredentials) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_key_id":     types.StringType,
			"access_point":      types.StringType,
			"secret_access_key": types.StringType,
			"session_token":     types.StringType,
		},
	}
}

// The AWS IAM role configuration
type AwsIamRole struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	ExternalId types.String `tfsdk:"external_id"`
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	RoleArn types.String `tfsdk:"role_arn"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn types.String `tfsdk:"unity_catalog_iam_arn"`
}

func (newState *AwsIamRole) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRole) {
}

func (newState *AwsIamRole) SyncEffectiveFieldsDuringRead(existingState AwsIamRole) {
}

func (c AwsIamRole) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetComputed()
	attrs["role_arn"] = attrs["role_arn"].SetOptional()
	attrs["unity_catalog_iam_arn"] = attrs["unity_catalog_iam_arn"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsIamRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsIamRole) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsIamRole
// only implements ToObjectValue() and Type().
func (o AwsIamRole) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id":           o.ExternalId,
			"role_arn":              o.RoleArn,
			"unity_catalog_iam_arn": o.UnityCatalogIamArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsIamRole) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id":           types.StringType,
			"role_arn":              types.StringType,
			"unity_catalog_iam_arn": types.StringType,
		},
	}
}

type AwsIamRoleRequest struct {
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn types.String `tfsdk:"role_arn"`
}

func (newState *AwsIamRoleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRoleRequest) {
}

func (newState *AwsIamRoleRequest) SyncEffectiveFieldsDuringRead(existingState AwsIamRoleRequest) {
}

func (c AwsIamRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["role_arn"] = attrs["role_arn"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsIamRoleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsIamRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsIamRoleRequest
// only implements ToObjectValue() and Type().
func (o AwsIamRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"role_arn": o.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsIamRoleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"role_arn": types.StringType,
		},
	}
}

type AwsIamRoleResponse struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	ExternalId types.String `tfsdk:"external_id"`
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn types.String `tfsdk:"role_arn"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn types.String `tfsdk:"unity_catalog_iam_arn"`
}

func (newState *AwsIamRoleResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRoleResponse) {
}

func (newState *AwsIamRoleResponse) SyncEffectiveFieldsDuringRead(existingState AwsIamRoleResponse) {
}

func (c AwsIamRoleResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["role_arn"] = attrs["role_arn"].SetRequired()
	attrs["unity_catalog_iam_arn"] = attrs["unity_catalog_iam_arn"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsIamRoleResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsIamRoleResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsIamRoleResponse
// only implements ToObjectValue() and Type().
func (o AwsIamRoleResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id":           o.ExternalId,
			"role_arn":              o.RoleArn,
			"unity_catalog_iam_arn": o.UnityCatalogIamArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsIamRoleResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id":           types.StringType,
			"role_arn":              types.StringType,
			"unity_catalog_iam_arn": types.StringType,
		},
	}
}

type AwsSqsQueue struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	ManagedResourceId types.String `tfsdk:"managed_resource_id"`
	// The AQS queue url in the format
	// https://sqs.{region}.amazonaws.com/{account id}/{queue name} REQUIRED for
	// provided_sqs.
	QueueUrl types.String `tfsdk:"queue_url"`
}

func (newState *AwsSqsQueue) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsSqsQueue) {
}

func (newState *AwsSqsQueue) SyncEffectiveFieldsDuringRead(existingState AwsSqsQueue) {
}

func (c AwsSqsQueue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_resource_id"] = attrs["managed_resource_id"].SetOptional()
	attrs["queue_url"] = attrs["queue_url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsSqsQueue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsSqsQueue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsSqsQueue
// only implements ToObjectValue() and Type().
func (o AwsSqsQueue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_resource_id": o.ManagedResourceId,
			"queue_url":           o.QueueUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsSqsQueue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_resource_id": types.StringType,
			"queue_url":           types.StringType,
		},
	}
}

// Azure Active Directory token, essentially the Oauth token for Azure Service
// Principal or Managed Identity. Read more at
// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
type AzureActiveDirectoryToken struct {
	// Opaque token that contains claims that you can use in Azure Active
	// Directory to access cloud services.
	AadToken types.String `tfsdk:"aad_token"`
}

func (newState *AzureActiveDirectoryToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureActiveDirectoryToken) {
}

func (newState *AzureActiveDirectoryToken) SyncEffectiveFieldsDuringRead(existingState AzureActiveDirectoryToken) {
}

func (c AzureActiveDirectoryToken) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aad_token"] = attrs["aad_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureActiveDirectoryToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureActiveDirectoryToken) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureActiveDirectoryToken
// only implements ToObjectValue() and Type().
func (o AzureActiveDirectoryToken) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aad_token": o.AadToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureActiveDirectoryToken) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aad_token": types.StringType,
		},
	}
}

// The Azure managed identity configuration.
type AzureManagedIdentity struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	AccessConnectorId types.String `tfsdk:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database. .
	CredentialId types.String `tfsdk:"credential_id"`
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	ManagedIdentityId types.String `tfsdk:"managed_identity_id"`
}

func (newState *AzureManagedIdentity) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentity) {
}

func (newState *AzureManagedIdentity) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentity) {
}

func (c AzureManagedIdentity) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_connector_id"] = attrs["access_connector_id"].SetRequired()
	attrs["credential_id"] = attrs["credential_id"].SetComputed()
	attrs["managed_identity_id"] = attrs["managed_identity_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureManagedIdentity.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureManagedIdentity) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureManagedIdentity
// only implements ToObjectValue() and Type().
func (o AzureManagedIdentity) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_connector_id": o.AccessConnectorId,
			"credential_id":       o.CredentialId,
			"managed_identity_id": o.ManagedIdentityId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureManagedIdentity) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_connector_id": types.StringType,
			"credential_id":       types.StringType,
			"managed_identity_id": types.StringType,
		},
	}
}

type AzureManagedIdentityRequest struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId types.String `tfsdk:"access_connector_id"`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId types.String `tfsdk:"managed_identity_id"`
}

func (newState *AzureManagedIdentityRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentityRequest) {
}

func (newState *AzureManagedIdentityRequest) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentityRequest) {
}

func (c AzureManagedIdentityRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_connector_id"] = attrs["access_connector_id"].SetRequired()
	attrs["managed_identity_id"] = attrs["managed_identity_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureManagedIdentityRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureManagedIdentityRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureManagedIdentityRequest
// only implements ToObjectValue() and Type().
func (o AzureManagedIdentityRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_connector_id": o.AccessConnectorId,
			"managed_identity_id": o.ManagedIdentityId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureManagedIdentityRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_connector_id": types.StringType,
			"managed_identity_id": types.StringType,
		},
	}
}

type AzureManagedIdentityResponse struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId types.String `tfsdk:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity.
	CredentialId types.String `tfsdk:"credential_id"`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId types.String `tfsdk:"managed_identity_id"`
}

func (newState *AzureManagedIdentityResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentityResponse) {
}

func (newState *AzureManagedIdentityResponse) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentityResponse) {
}

func (c AzureManagedIdentityResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_connector_id"] = attrs["access_connector_id"].SetRequired()
	attrs["credential_id"] = attrs["credential_id"].SetOptional()
	attrs["managed_identity_id"] = attrs["managed_identity_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureManagedIdentityResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureManagedIdentityResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureManagedIdentityResponse
// only implements ToObjectValue() and Type().
func (o AzureManagedIdentityResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_connector_id": o.AccessConnectorId,
			"credential_id":       o.CredentialId,
			"managed_identity_id": o.ManagedIdentityId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureManagedIdentityResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_connector_id": types.StringType,
			"credential_id":       types.StringType,
			"managed_identity_id": types.StringType,
		},
	}
}

type AzureQueueStorage struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	ManagedResourceId types.String `tfsdk:"managed_resource_id"`
	// The AQS queue url in the format https://{storage
	// account}.queue.core.windows.net/{queue name} REQUIRED for provided_aqs.
	QueueUrl types.String `tfsdk:"queue_url"`
	// The resource group for the queue, event grid subscription, and external
	// location storage account. ONLY REQUIRED for locations with a service
	// principal storage credential
	ResourceGroup types.String `tfsdk:"resource_group"`
	// OPTIONAL: The subscription id for the queue, event grid subscription, and
	// external location storage account. REQUIRED for locations with a service
	// principal storage credential
	SubscriptionId types.String `tfsdk:"subscription_id"`
}

func (newState *AzureQueueStorage) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureQueueStorage) {
}

func (newState *AzureQueueStorage) SyncEffectiveFieldsDuringRead(existingState AzureQueueStorage) {
}

func (c AzureQueueStorage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_resource_id"] = attrs["managed_resource_id"].SetOptional()
	attrs["queue_url"] = attrs["queue_url"].SetOptional()
	attrs["resource_group"] = attrs["resource_group"].SetOptional()
	attrs["subscription_id"] = attrs["subscription_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureQueueStorage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureQueueStorage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureQueueStorage
// only implements ToObjectValue() and Type().
func (o AzureQueueStorage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_resource_id": o.ManagedResourceId,
			"queue_url":           o.QueueUrl,
			"resource_group":      o.ResourceGroup,
			"subscription_id":     o.SubscriptionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureQueueStorage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_resource_id": types.StringType,
			"queue_url":           types.StringType,
			"resource_group":      types.StringType,
			"subscription_id":     types.StringType,
		},
	}
}

// The Azure service principal configuration. Only applicable when purpose is
// **STORAGE**.
type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	ApplicationId types.String `tfsdk:"application_id"`
	// The client secret generated for the above app ID in AAD.
	ClientSecret types.String `tfsdk:"client_secret"`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	DirectoryId types.String `tfsdk:"directory_id"`
}

func (newState *AzureServicePrincipal) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureServicePrincipal) {
}

func (newState *AzureServicePrincipal) SyncEffectiveFieldsDuringRead(existingState AzureServicePrincipal) {
}

func (c AzureServicePrincipal) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["application_id"] = attrs["application_id"].SetRequired()
	attrs["client_secret"] = attrs["client_secret"].SetRequired()
	attrs["directory_id"] = attrs["directory_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureServicePrincipal.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureServicePrincipal) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureServicePrincipal
// only implements ToObjectValue() and Type().
func (o AzureServicePrincipal) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"application_id": o.ApplicationId,
			"client_secret":  o.ClientSecret,
			"directory_id":   o.DirectoryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureServicePrincipal) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"application_id": types.StringType,
			"client_secret":  types.StringType,
			"directory_id":   types.StringType,
		},
	}
}

// Azure temporary credentials for API authentication. Read more at
// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
type AzureUserDelegationSas struct {
	// The signed URI (SAS Token) used to access blob services for a given path
	SasToken types.String `tfsdk:"sas_token"`
}

func (newState *AzureUserDelegationSas) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureUserDelegationSas) {
}

func (newState *AzureUserDelegationSas) SyncEffectiveFieldsDuringRead(existingState AzureUserDelegationSas) {
}

func (c AzureUserDelegationSas) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sas_token"] = attrs["sas_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureUserDelegationSas.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureUserDelegationSas) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureUserDelegationSas
// only implements ToObjectValue() and Type().
func (o AzureUserDelegationSas) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sas_token": o.SasToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureUserDelegationSas) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sas_token": types.StringType,
		},
	}
}

// Cancel refresh
type CancelRefreshRequest struct {
	// ID of the refresh.
	RefreshId types.String `tfsdk:"-"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelRefreshRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelRefreshRequest
// only implements ToObjectValue() and Type().
func (o CancelRefreshRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"refresh_id": o.RefreshId,
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CancelRefreshRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"refresh_id": types.StringType,
			"table_name": types.StringType,
		},
	}
}

type CancelRefreshResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRefreshResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelRefreshResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelRefreshResponse
// only implements ToObjectValue() and Type().
func (o CancelRefreshResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CancelRefreshResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CatalogInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// The type of the catalog.
	CatalogType types.String `tfsdk:"catalog_type"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// The name of the connection to an external data source.
	ConnectionName types.String `tfsdk:"connection_name"`
	// Time at which this catalog was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of catalog creator.
	CreatedBy types.String `tfsdk:"created_by"`

	EffectivePredictiveOptimizationFlag types.Object `tfsdk:"effective_predictive_optimization_flag"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization"`
	// The full name of the catalog. Corresponds with the name field.
	FullName types.String `tfsdk:"full_name"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of catalog.
	Name types.String `tfsdk:"name"`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options"`
	// Username of current owner of catalog.
	Owner types.String `tfsdk:"owner"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName types.String `tfsdk:"provider_name"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo types.Object `tfsdk:"provisioning_info"`
	// The type of Unity Catalog securable.
	SecurableType types.String `tfsdk:"securable_type"`
	// The name of the share under the share provider.
	ShareName types.String `tfsdk:"share_name"`
	// Storage Location URL (full path) for managed tables within catalog.
	StorageLocation types.String `tfsdk:"storage_location"`
	// Storage root URL for managed tables within catalog.
	StorageRoot types.String `tfsdk:"storage_root"`
	// Time at which this catalog was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified catalog.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *CatalogInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CatalogInfo) {
}

func (newState *CatalogInfo) SyncEffectiveFieldsDuringRead(existingState CatalogInfo) {
}

func (c CatalogInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["browse_only"] = attrs["browse_only"].SetOptional()
	attrs["catalog_type"] = attrs["catalog_type"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["connection_name"] = attrs["connection_name"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["effective_predictive_optimization_flag"] = attrs["effective_predictive_optimization_flag"].SetOptional()
	attrs["enable_predictive_optimization"] = attrs["enable_predictive_optimization"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["isolation_mode"] = attrs["isolation_mode"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["provider_name"] = attrs["provider_name"].SetOptional()
	attrs["provisioning_info"] = attrs["provisioning_info"].SetOptional()
	attrs["securable_type"] = attrs["securable_type"].SetOptional()
	attrs["share_name"] = attrs["share_name"].SetOptional()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CatalogInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CatalogInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_predictive_optimization_flag": reflect.TypeOf(EffectivePredictiveOptimizationFlag{}),
		"options":                                reflect.TypeOf(types.String{}),
		"properties":                             reflect.TypeOf(types.String{}),
		"provisioning_info":                      reflect.TypeOf(ProvisioningInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CatalogInfo
// only implements ToObjectValue() and Type().
func (o CatalogInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"browse_only":                            o.BrowseOnly,
			"catalog_type":                           o.CatalogType,
			"comment":                                o.Comment,
			"connection_name":                        o.ConnectionName,
			"created_at":                             o.CreatedAt,
			"created_by":                             o.CreatedBy,
			"effective_predictive_optimization_flag": o.EffectivePredictiveOptimizationFlag,
			"enable_predictive_optimization":         o.EnablePredictiveOptimization,
			"full_name":                              o.FullName,
			"isolation_mode":                         o.IsolationMode,
			"metastore_id":                           o.MetastoreId,
			"name":                                   o.Name,
			"options":                                o.Options,
			"owner":                                  o.Owner,
			"properties":                             o.Properties,
			"provider_name":                          o.ProviderName,
			"provisioning_info":                      o.ProvisioningInfo,
			"securable_type":                         o.SecurableType,
			"share_name":                             o.ShareName,
			"storage_location":                       o.StorageLocation,
			"storage_root":                           o.StorageRoot,
			"updated_at":                             o.UpdatedAt,
			"updated_by":                             o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CatalogInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":                            types.BoolType,
			"catalog_type":                           types.StringType,
			"comment":                                types.StringType,
			"connection_name":                        types.StringType,
			"created_at":                             types.Int64Type,
			"created_by":                             types.StringType,
			"effective_predictive_optimization_flag": EffectivePredictiveOptimizationFlag{}.Type(ctx),
			"enable_predictive_optimization":         types.StringType,
			"full_name":                              types.StringType,
			"isolation_mode":                         types.StringType,
			"metastore_id":                           types.StringType,
			"name":                                   types.StringType,
			"options": basetypes.MapType{
				ElemType: types.StringType,
			},
			"owner": types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"provider_name":     types.StringType,
			"provisioning_info": ProvisioningInfo{}.Type(ctx),
			"securable_type":    types.StringType,
			"share_name":        types.StringType,
			"storage_location":  types.StringType,
			"storage_root":      types.StringType,
			"updated_at":        types.Int64Type,
			"updated_by":        types.StringType,
		},
	}
}

// GetEffectivePredictiveOptimizationFlag returns the value of the EffectivePredictiveOptimizationFlag field in CatalogInfo as
// a EffectivePredictiveOptimizationFlag value.
// If the field is unknown or null, the boolean return value is false.
func (o *CatalogInfo) GetEffectivePredictiveOptimizationFlag(ctx context.Context) (EffectivePredictiveOptimizationFlag, bool) {
	var e EffectivePredictiveOptimizationFlag
	if o.EffectivePredictiveOptimizationFlag.IsNull() || o.EffectivePredictiveOptimizationFlag.IsUnknown() {
		return e, false
	}
	var v []EffectivePredictiveOptimizationFlag
	d := o.EffectivePredictiveOptimizationFlag.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEffectivePredictiveOptimizationFlag sets the value of the EffectivePredictiveOptimizationFlag field in CatalogInfo.
func (o *CatalogInfo) SetEffectivePredictiveOptimizationFlag(ctx context.Context, v EffectivePredictiveOptimizationFlag) {
	vs := v.ToObjectValue(ctx)
	o.EffectivePredictiveOptimizationFlag = vs
}

// GetOptions returns the value of the Options field in CatalogInfo as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CatalogInfo) GetOptions(ctx context.Context) (map[string]types.String, bool) {
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in CatalogInfo.
func (o *CatalogInfo) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in CatalogInfo as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CatalogInfo) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in CatalogInfo.
func (o *CatalogInfo) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

// GetProvisioningInfo returns the value of the ProvisioningInfo field in CatalogInfo as
// a ProvisioningInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CatalogInfo) GetProvisioningInfo(ctx context.Context) (ProvisioningInfo, bool) {
	var e ProvisioningInfo
	if o.ProvisioningInfo.IsNull() || o.ProvisioningInfo.IsUnknown() {
		return e, false
	}
	var v []ProvisioningInfo
	d := o.ProvisioningInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetProvisioningInfo sets the value of the ProvisioningInfo field in CatalogInfo.
func (o *CatalogInfo) SetProvisioningInfo(ctx context.Context, v ProvisioningInfo) {
	vs := v.ToObjectValue(ctx)
	o.ProvisioningInfo = vs
}

type CloudflareApiToken struct {
	// The Cloudflare access key id of the token.
	AccessKeyId types.String `tfsdk:"access_key_id"`
	// The account id associated with the API token.
	AccountId types.String `tfsdk:"account_id"`
	// The secret access token generated for the access key id
	SecretAccessKey types.String `tfsdk:"secret_access_key"`
}

func (newState *CloudflareApiToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloudflareApiToken) {
}

func (newState *CloudflareApiToken) SyncEffectiveFieldsDuringRead(existingState CloudflareApiToken) {
}

func (c CloudflareApiToken) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_key_id"] = attrs["access_key_id"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["secret_access_key"] = attrs["secret_access_key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloudflareApiToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CloudflareApiToken) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudflareApiToken
// only implements ToObjectValue() and Type().
func (o CloudflareApiToken) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_key_id":     o.AccessKeyId,
			"account_id":        o.AccountId,
			"secret_access_key": o.SecretAccessKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CloudflareApiToken) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_key_id":     types.StringType,
			"account_id":        types.StringType,
			"secret_access_key": types.StringType,
		},
	}
}

type ColumnInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`

	Mask types.Object `tfsdk:"mask"`
	// Name of Column.
	Name types.String `tfsdk:"name"`
	// Whether field may be Null (default: true).
	Nullable types.Bool `tfsdk:"nullable"`
	// Partition index for column.
	PartitionIndex types.Int64 `tfsdk:"partition_index"`
	// Ordinal position of column (starting at position 0).
	Position types.Int64 `tfsdk:"position"`
	// Format of IntervalType.
	TypeIntervalType types.String `tfsdk:"type_interval_type"`
	// Full data type specification, JSON-serialized.
	TypeJson types.String `tfsdk:"type_json"`

	TypeName types.String `tfsdk:"type_name"`
	// Digits of precision; required for DecimalTypes.
	TypePrecision types.Int64 `tfsdk:"type_precision"`
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale types.Int64 `tfsdk:"type_scale"`
	// Full data type specification as SQL/catalogString text.
	TypeText types.String `tfsdk:"type_text"`
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnInfo) {
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringRead(existingState ColumnInfo) {
}

func (c ColumnInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["mask"] = attrs["mask"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["nullable"] = attrs["nullable"].SetOptional()
	attrs["partition_index"] = attrs["partition_index"].SetOptional()
	attrs["position"] = attrs["position"].SetOptional()
	attrs["type_interval_type"] = attrs["type_interval_type"].SetOptional()
	attrs["type_json"] = attrs["type_json"].SetOptional()
	attrs["type_name"] = attrs["type_name"].SetOptional()
	attrs["type_precision"] = attrs["type_precision"].SetOptional()
	attrs["type_scale"] = attrs["type_scale"].SetOptional()
	attrs["type_text"] = attrs["type_text"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ColumnInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ColumnInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"mask": reflect.TypeOf(ColumnMask{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo
// only implements ToObjectValue() and Type().
func (o ColumnInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":            o.Comment,
			"mask":               o.Mask,
			"name":               o.Name,
			"nullable":           o.Nullable,
			"partition_index":    o.PartitionIndex,
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
func (o ColumnInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":            types.StringType,
			"mask":               ColumnMask{}.Type(ctx),
			"name":               types.StringType,
			"nullable":           types.BoolType,
			"partition_index":    types.Int64Type,
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

// GetMask returns the value of the Mask field in ColumnInfo as
// a ColumnMask value.
// If the field is unknown or null, the boolean return value is false.
func (o *ColumnInfo) GetMask(ctx context.Context) (ColumnMask, bool) {
	var e ColumnMask
	if o.Mask.IsNull() || o.Mask.IsUnknown() {
		return e, false
	}
	var v []ColumnMask
	d := o.Mask.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetMask sets the value of the Mask field in ColumnInfo.
func (o *ColumnInfo) SetMask(ctx context.Context, v ColumnMask) {
	vs := v.ToObjectValue(ctx)
	o.Mask = vs
}

type ColumnMask struct {
	// The full name of the column mask SQL UDF.
	FunctionName types.String `tfsdk:"function_name"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames types.List `tfsdk:"using_column_names"`
}

func (newState *ColumnMask) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnMask) {
}

func (newState *ColumnMask) SyncEffectiveFieldsDuringRead(existingState ColumnMask) {
}

func (c ColumnMask) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["function_name"] = attrs["function_name"].SetOptional()
	attrs["using_column_names"] = attrs["using_column_names"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ColumnMask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ColumnMask) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"using_column_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnMask
// only implements ToObjectValue() and Type().
func (o ColumnMask) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_name":      o.FunctionName,
			"using_column_names": o.UsingColumnNames,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ColumnMask) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_name": types.StringType,
			"using_column_names": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetUsingColumnNames returns the value of the UsingColumnNames field in ColumnMask as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ColumnMask) GetUsingColumnNames(ctx context.Context) ([]types.String, bool) {
	if o.UsingColumnNames.IsNull() || o.UsingColumnNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.UsingColumnNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUsingColumnNames sets the value of the UsingColumnNames field in ColumnMask.
func (o *ColumnMask) SetUsingColumnNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["using_column_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UsingColumnNames = types.ListValueMust(t, vs)
}

type ConnectionInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Unique identifier of the Connection.
	ConnectionId types.String `tfsdk:"connection_id"`
	// The type of connection.
	ConnectionType types.String `tfsdk:"connection_type"`
	// Time at which this connection was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of connection creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// The type of credential.
	CredentialType types.String `tfsdk:"credential_type"`
	// Full name of connection.
	FullName types.String `tfsdk:"full_name"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of the connection.
	Name types.String `tfsdk:"name"`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options"`
	// Username of current owner of the connection.
	Owner types.String `tfsdk:"owner"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo types.Object `tfsdk:"provisioning_info"`
	// If the connection is read only.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// The type of Unity Catalog securable.
	SecurableType types.String `tfsdk:"securable_type"`
	// Time at which this connection was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified connection.
	UpdatedBy types.String `tfsdk:"updated_by"`
	// URL of the remote data source, extracted from options.
	Url types.String `tfsdk:"url"`
}

func (newState *ConnectionInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ConnectionInfo) {
}

func (newState *ConnectionInfo) SyncEffectiveFieldsDuringRead(existingState ConnectionInfo) {
}

func (c ConnectionInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["connection_id"] = attrs["connection_id"].SetOptional()
	attrs["connection_type"] = attrs["connection_type"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["credential_type"] = attrs["credential_type"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["provisioning_info"] = attrs["provisioning_info"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["securable_type"] = attrs["securable_type"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConnectionInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ConnectionInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options":           reflect.TypeOf(types.String{}),
		"properties":        reflect.TypeOf(types.String{}),
		"provisioning_info": reflect.TypeOf(ProvisioningInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConnectionInfo
// only implements ToObjectValue() and Type().
func (o ConnectionInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":           o.Comment,
			"connection_id":     o.ConnectionId,
			"connection_type":   o.ConnectionType,
			"created_at":        o.CreatedAt,
			"created_by":        o.CreatedBy,
			"credential_type":   o.CredentialType,
			"full_name":         o.FullName,
			"metastore_id":      o.MetastoreId,
			"name":              o.Name,
			"options":           o.Options,
			"owner":             o.Owner,
			"properties":        o.Properties,
			"provisioning_info": o.ProvisioningInfo,
			"read_only":         o.ReadOnly,
			"securable_type":    o.SecurableType,
			"updated_at":        o.UpdatedAt,
			"updated_by":        o.UpdatedBy,
			"url":               o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ConnectionInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":         types.StringType,
			"connection_id":   types.StringType,
			"connection_type": types.StringType,
			"created_at":      types.Int64Type,
			"created_by":      types.StringType,
			"credential_type": types.StringType,
			"full_name":       types.StringType,
			"metastore_id":    types.StringType,
			"name":            types.StringType,
			"options": basetypes.MapType{
				ElemType: types.StringType,
			},
			"owner": types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"provisioning_info": ProvisioningInfo{}.Type(ctx),
			"read_only":         types.BoolType,
			"securable_type":    types.StringType,
			"updated_at":        types.Int64Type,
			"updated_by":        types.StringType,
			"url":               types.StringType,
		},
	}
}

// GetOptions returns the value of the Options field in ConnectionInfo as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ConnectionInfo) GetOptions(ctx context.Context) (map[string]types.String, bool) {
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in ConnectionInfo.
func (o *ConnectionInfo) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in ConnectionInfo as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ConnectionInfo) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in ConnectionInfo.
func (o *ConnectionInfo) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

// GetProvisioningInfo returns the value of the ProvisioningInfo field in ConnectionInfo as
// a ProvisioningInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ConnectionInfo) GetProvisioningInfo(ctx context.Context) (ProvisioningInfo, bool) {
	var e ProvisioningInfo
	if o.ProvisioningInfo.IsNull() || o.ProvisioningInfo.IsUnknown() {
		return e, false
	}
	var v []ProvisioningInfo
	d := o.ProvisioningInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetProvisioningInfo sets the value of the ProvisioningInfo field in ConnectionInfo.
func (o *ConnectionInfo) SetProvisioningInfo(ctx context.Context, v ProvisioningInfo) {
	vs := v.ToObjectValue(ctx)
	o.ProvisioningInfo = vs
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
type ContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress types.Object `tfsdk:"initial_pipeline_sync_progress"`
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp types.String `tfsdk:"timestamp"`
}

func (newState *ContinuousUpdateStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ContinuousUpdateStatus) {
}

func (newState *ContinuousUpdateStatus) SyncEffectiveFieldsDuringRead(existingState ContinuousUpdateStatus) {
}

func (c ContinuousUpdateStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].SetOptional()
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ContinuousUpdateStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ContinuousUpdateStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(PipelineProgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContinuousUpdateStatus
// only implements ToObjectValue() and Type().
func (o ContinuousUpdateStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": o.InitialPipelineSyncProgress,
			"last_processed_commit_version":  o.LastProcessedCommitVersion,
			"timestamp":                      o.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ContinuousUpdateStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": PipelineProgress{}.Type(ctx),
			"last_processed_commit_version":  types.Int64Type,
			"timestamp":                      types.StringType,
		},
	}
}

// GetInitialPipelineSyncProgress returns the value of the InitialPipelineSyncProgress field in ContinuousUpdateStatus as
// a PipelineProgress value.
// If the field is unknown or null, the boolean return value is false.
func (o *ContinuousUpdateStatus) GetInitialPipelineSyncProgress(ctx context.Context) (PipelineProgress, bool) {
	var e PipelineProgress
	if o.InitialPipelineSyncProgress.IsNull() || o.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v []PipelineProgress
	d := o.InitialPipelineSyncProgress.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in ContinuousUpdateStatus.
func (o *ContinuousUpdateStatus) SetInitialPipelineSyncProgress(ctx context.Context, v PipelineProgress) {
	vs := v.ToObjectValue(ctx)
	o.InitialPipelineSyncProgress = vs
}

type CreateCatalog struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// The name of the connection to an external data source.
	ConnectionName types.String `tfsdk:"connection_name"`
	// Name of catalog.
	Name types.String `tfsdk:"name"`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName types.String `tfsdk:"provider_name"`
	// The name of the share under the share provider.
	ShareName types.String `tfsdk:"share_name"`
	// Storage root URL for managed tables within catalog.
	StorageRoot types.String `tfsdk:"storage_root"`
}

func (newState *CreateCatalog) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCatalog) {
}

func (newState *CreateCatalog) SyncEffectiveFieldsDuringRead(existingState CreateCatalog) {
}

func (c CreateCatalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["connection_name"] = attrs["connection_name"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["provider_name"] = attrs["provider_name"].SetOptional()
	attrs["share_name"] = attrs["share_name"].SetOptional()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCatalog.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCatalog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options":    reflect.TypeOf(types.String{}),
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCatalog
// only implements ToObjectValue() and Type().
func (o CreateCatalog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":         o.Comment,
			"connection_name": o.ConnectionName,
			"name":            o.Name,
			"options":         o.Options,
			"properties":      o.Properties,
			"provider_name":   o.ProviderName,
			"share_name":      o.ShareName,
			"storage_root":    o.StorageRoot,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCatalog) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":         types.StringType,
			"connection_name": types.StringType,
			"name":            types.StringType,
			"options": basetypes.MapType{
				ElemType: types.StringType,
			},
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"provider_name": types.StringType,
			"share_name":    types.StringType,
			"storage_root":  types.StringType,
		},
	}
}

// GetOptions returns the value of the Options field in CreateCatalog as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCatalog) GetOptions(ctx context.Context) (map[string]types.String, bool) {
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in CreateCatalog.
func (o *CreateCatalog) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in CreateCatalog as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCatalog) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in CreateCatalog.
func (o *CreateCatalog) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type CreateConnection struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// The type of connection.
	ConnectionType types.String `tfsdk:"connection_type"`
	// Name of the connection.
	Name types.String `tfsdk:"name"`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`
	// If the connection is read only.
	ReadOnly types.Bool `tfsdk:"read_only"`
}

func (newState *CreateConnection) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateConnection) {
}

func (newState *CreateConnection) SyncEffectiveFieldsDuringRead(existingState CreateConnection) {
}

func (c CreateConnection) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["connection_type"] = attrs["connection_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateConnection.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateConnection) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options":    reflect.TypeOf(types.String{}),
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateConnection
// only implements ToObjectValue() and Type().
func (o CreateConnection) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":         o.Comment,
			"connection_type": o.ConnectionType,
			"name":            o.Name,
			"options":         o.Options,
			"properties":      o.Properties,
			"read_only":       o.ReadOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateConnection) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":         types.StringType,
			"connection_type": types.StringType,
			"name":            types.StringType,
			"options": basetypes.MapType{
				ElemType: types.StringType,
			},
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"read_only": types.BoolType,
		},
	}
}

// GetOptions returns the value of the Options field in CreateConnection as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateConnection) GetOptions(ctx context.Context) (map[string]types.String, bool) {
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in CreateConnection.
func (o *CreateConnection) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in CreateConnection as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateConnection) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in CreateConnection.
func (o *CreateConnection) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type CreateCredentialRequest struct {
	// The AWS IAM role configuration
	AwsIamRole types.Object `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.Object `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal types.Object `tfsdk:"azure_service_principal"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount types.Object `tfsdk:"databricks_gcp_service_account"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name types.String `tfsdk:"name"`
	// Indicates the purpose of the credential.
	Purpose types.String `tfsdk:"purpose"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	SkipValidation types.Bool `tfsdk:"skip_validation"`
}

func (newState *CreateCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialRequest) {
}

func (newState *CreateCredentialRequest) SyncEffectiveFieldsDuringRead(existingState CreateCredentialRequest) {
}

func (c CreateCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_iam_role"] = attrs["aws_iam_role"].SetOptional()
	attrs["azure_managed_identity"] = attrs["azure_managed_identity"].SetOptional()
	attrs["azure_service_principal"] = attrs["azure_service_principal"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["databricks_gcp_service_account"] = attrs["databricks_gcp_service_account"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["purpose"] = attrs["purpose"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["skip_validation"] = attrs["skip_validation"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRole{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentity{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialRequest
// only implements ToObjectValue() and Type().
func (o CreateCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_iam_role":                   o.AwsIamRole,
			"azure_managed_identity":         o.AzureManagedIdentity,
			"azure_service_principal":        o.AzureServicePrincipal,
			"comment":                        o.Comment,
			"databricks_gcp_service_account": o.DatabricksGcpServiceAccount,
			"name":                           o.Name,
			"purpose":                        o.Purpose,
			"read_only":                      o.ReadOnly,
			"skip_validation":                o.SkipValidation,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role":                   AwsIamRole{}.Type(ctx),
			"azure_managed_identity":         AzureManagedIdentity{}.Type(ctx),
			"azure_service_principal":        AzureServicePrincipal{}.Type(ctx),
			"comment":                        types.StringType,
			"databricks_gcp_service_account": DatabricksGcpServiceAccount{}.Type(ctx),
			"name":                           types.StringType,
			"purpose":                        types.StringType,
			"read_only":                      types.BoolType,
			"skip_validation":                types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in CreateCredentialRequest as
// a AwsIamRole value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest) GetAwsIamRole(ctx context.Context) (AwsIamRole, bool) {
	var e AwsIamRole
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRole
	d := o.AwsIamRole.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsIamRole sets the value of the AwsIamRole field in CreateCredentialRequest.
func (o *CreateCredentialRequest) SetAwsIamRole(ctx context.Context, v AwsIamRole) {
	vs := v.ToObjectValue(ctx)
	o.AwsIamRole = vs
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in CreateCredentialRequest as
// a AzureManagedIdentity value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentity, bool) {
	var e AzureManagedIdentity
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentity
	d := o.AzureManagedIdentity.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in CreateCredentialRequest.
func (o *CreateCredentialRequest) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentity) {
	vs := v.ToObjectValue(ctx)
	o.AzureManagedIdentity = vs
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in CreateCredentialRequest as
// a AzureServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal, bool) {
	var e AzureServicePrincipal
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal
	d := o.AzureServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in CreateCredentialRequest.
func (o *CreateCredentialRequest) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	o.AzureServicePrincipal = vs
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in CreateCredentialRequest as
// a DatabricksGcpServiceAccount value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccount, bool) {
	var e DatabricksGcpServiceAccount
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccount
	d := o.DatabricksGcpServiceAccount.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in CreateCredentialRequest.
func (o *CreateCredentialRequest) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccount) {
	vs := v.ToObjectValue(ctx)
	o.DatabricksGcpServiceAccount = vs
}

// Create a Database Catalog
type CreateDatabaseCatalogRequest struct {
	Catalog types.Object `tfsdk:"catalog"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalog": reflect.TypeOf(DatabaseCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (o CreateDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": o.Catalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog": DatabaseCatalog{}.Type(ctx),
		},
	}
}

// GetCatalog returns the value of the Catalog field in CreateDatabaseCatalogRequest as
// a DatabaseCatalog value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateDatabaseCatalogRequest) GetCatalog(ctx context.Context) (DatabaseCatalog, bool) {
	var e DatabaseCatalog
	if o.Catalog.IsNull() || o.Catalog.IsUnknown() {
		return e, false
	}
	var v []DatabaseCatalog
	d := o.Catalog.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetCatalog sets the value of the Catalog field in CreateDatabaseCatalogRequest.
func (o *CreateDatabaseCatalogRequest) SetCatalog(ctx context.Context, v DatabaseCatalog) {
	vs := v.ToObjectValue(ctx)
	o.Catalog = vs
}

// Create a Database Instance
type CreateDatabaseInstanceRequest struct {
	// A DatabaseInstance represents a logical Postgres instance, comprised of
	// both compute and storage.
	DatabaseInstance types.Object `tfsdk:"database_instance"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance": reflect.TypeOf(DatabaseInstance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (o CreateDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance": o.DatabaseInstance,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance": DatabaseInstance{}.Type(ctx),
		},
	}
}

// GetDatabaseInstance returns the value of the DatabaseInstance field in CreateDatabaseInstanceRequest as
// a DatabaseInstance value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateDatabaseInstanceRequest) GetDatabaseInstance(ctx context.Context) (DatabaseInstance, bool) {
	var e DatabaseInstance
	if o.DatabaseInstance.IsNull() || o.DatabaseInstance.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstance
	d := o.DatabaseInstance.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDatabaseInstance sets the value of the DatabaseInstance field in CreateDatabaseInstanceRequest.
func (o *CreateDatabaseInstanceRequest) SetDatabaseInstance(ctx context.Context, v DatabaseInstance) {
	vs := v.ToObjectValue(ctx)
	o.DatabaseInstance = vs
}

type CreateExternalLocation struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Name of the storage credential used with this location.
	CredentialName types.String `tfsdk:"credential_name"`
	// [Create:OPT Update:OPT] Whether to enable file events on this external
	// location.
	EnableFileEvents types.Bool `tfsdk:"enable_file_events"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails types.Object `tfsdk:"encryption_details"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback types.Bool `tfsdk:"fallback"`
	// [Create:OPT Update:OPT] File event queue settings.
	FileEventQueue types.Object `tfsdk:"file_event_queue"`
	// Name of the external location.
	Name types.String `tfsdk:"name"`
	// Indicates whether the external location is read-only.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation types.Bool `tfsdk:"skip_validation"`
	// Path URL of the external location.
	Url types.String `tfsdk:"url"`
}

func (newState *CreateExternalLocation) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExternalLocation) {
}

func (newState *CreateExternalLocation) SyncEffectiveFieldsDuringRead(existingState CreateExternalLocation) {
}

func (c CreateExternalLocation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["credential_name"] = attrs["credential_name"].SetRequired()
	attrs["enable_file_events"] = attrs["enable_file_events"].SetOptional()
	attrs["encryption_details"] = attrs["encryption_details"].SetOptional()
	attrs["fallback"] = attrs["fallback"].SetOptional()
	attrs["file_event_queue"] = attrs["file_event_queue"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["skip_validation"] = attrs["skip_validation"].SetOptional()
	attrs["url"] = attrs["url"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExternalLocation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExternalLocation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"encryption_details": reflect.TypeOf(EncryptionDetails{}),
		"file_event_queue":   reflect.TypeOf(FileEventQueue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExternalLocation
// only implements ToObjectValue() and Type().
func (o CreateExternalLocation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":            o.Comment,
			"credential_name":    o.CredentialName,
			"enable_file_events": o.EnableFileEvents,
			"encryption_details": o.EncryptionDetails,
			"fallback":           o.Fallback,
			"file_event_queue":   o.FileEventQueue,
			"name":               o.Name,
			"read_only":          o.ReadOnly,
			"skip_validation":    o.SkipValidation,
			"url":                o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExternalLocation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":            types.StringType,
			"credential_name":    types.StringType,
			"enable_file_events": types.BoolType,
			"encryption_details": EncryptionDetails{}.Type(ctx),
			"fallback":           types.BoolType,
			"file_event_queue":   FileEventQueue{}.Type(ctx),
			"name":               types.StringType,
			"read_only":          types.BoolType,
			"skip_validation":    types.BoolType,
			"url":                types.StringType,
		},
	}
}

// GetEncryptionDetails returns the value of the EncryptionDetails field in CreateExternalLocation as
// a EncryptionDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateExternalLocation) GetEncryptionDetails(ctx context.Context) (EncryptionDetails, bool) {
	var e EncryptionDetails
	if o.EncryptionDetails.IsNull() || o.EncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []EncryptionDetails
	d := o.EncryptionDetails.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEncryptionDetails sets the value of the EncryptionDetails field in CreateExternalLocation.
func (o *CreateExternalLocation) SetEncryptionDetails(ctx context.Context, v EncryptionDetails) {
	vs := v.ToObjectValue(ctx)
	o.EncryptionDetails = vs
}

// GetFileEventQueue returns the value of the FileEventQueue field in CreateExternalLocation as
// a FileEventQueue value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateExternalLocation) GetFileEventQueue(ctx context.Context) (FileEventQueue, bool) {
	var e FileEventQueue
	if o.FileEventQueue.IsNull() || o.FileEventQueue.IsUnknown() {
		return e, false
	}
	var v []FileEventQueue
	d := o.FileEventQueue.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetFileEventQueue sets the value of the FileEventQueue field in CreateExternalLocation.
func (o *CreateExternalLocation) SetFileEventQueue(ctx context.Context, v FileEventQueue) {
	vs := v.ToObjectValue(ctx)
	o.FileEventQueue = vs
}

type CreateFunction struct {
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Scalar function return data type.
	DataType types.String `tfsdk:"data_type"`
	// External function language.
	ExternalLanguage types.String `tfsdk:"external_language"`
	// External function name.
	ExternalName types.String `tfsdk:"external_name"`
	// Pretty printed function data type.
	FullDataType types.String `tfsdk:"full_data_type"`

	InputParams types.Object `tfsdk:"input_params"`
	// Whether the function is deterministic.
	IsDeterministic types.Bool `tfsdk:"is_deterministic"`
	// Function null call.
	IsNullCall types.Bool `tfsdk:"is_null_call"`
	// Name of function, relative to parent schema.
	Name types.String `tfsdk:"name"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle types.String `tfsdk:"parameter_style"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties types.String `tfsdk:"properties"`
	// Table function return parameters.
	ReturnParams types.Object `tfsdk:"return_params"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody types.String `tfsdk:"routine_body"`
	// Function body.
	RoutineDefinition types.String `tfsdk:"routine_definition"`
	// Function dependencies.
	RoutineDependencies types.Object `tfsdk:"routine_dependencies"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name"`
	// Function security type.
	SecurityType types.String `tfsdk:"security_type"`
	// Specific name of the function; Reserved for future use.
	SpecificName types.String `tfsdk:"specific_name"`
	// Function SQL data access.
	SqlDataAccess types.String `tfsdk:"sql_data_access"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath types.String `tfsdk:"sql_path"`
}

func (newState *CreateFunction) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFunction) {
}

func (newState *CreateFunction) SyncEffectiveFieldsDuringRead(existingState CreateFunction) {
}

func (c CreateFunction) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["data_type"] = attrs["data_type"].SetRequired()
	attrs["external_language"] = attrs["external_language"].SetOptional()
	attrs["external_name"] = attrs["external_name"].SetOptional()
	attrs["full_data_type"] = attrs["full_data_type"].SetRequired()
	attrs["input_params"] = attrs["input_params"].SetRequired()
	attrs["is_deterministic"] = attrs["is_deterministic"].SetRequired()
	attrs["is_null_call"] = attrs["is_null_call"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["parameter_style"] = attrs["parameter_style"].SetRequired()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["return_params"] = attrs["return_params"].SetOptional()
	attrs["routine_body"] = attrs["routine_body"].SetRequired()
	attrs["routine_definition"] = attrs["routine_definition"].SetRequired()
	attrs["routine_dependencies"] = attrs["routine_dependencies"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetRequired()
	attrs["security_type"] = attrs["security_type"].SetRequired()
	attrs["specific_name"] = attrs["specific_name"].SetRequired()
	attrs["sql_data_access"] = attrs["sql_data_access"].SetRequired()
	attrs["sql_path"] = attrs["sql_path"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFunction.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateFunction) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input_params":         reflect.TypeOf(FunctionParameterInfos{}),
		"return_params":        reflect.TypeOf(FunctionParameterInfos{}),
		"routine_dependencies": reflect.TypeOf(DependencyList{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFunction
// only implements ToObjectValue() and Type().
func (o CreateFunction) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":         o.CatalogName,
			"comment":              o.Comment,
			"data_type":            o.DataType,
			"external_language":    o.ExternalLanguage,
			"external_name":        o.ExternalName,
			"full_data_type":       o.FullDataType,
			"input_params":         o.InputParams,
			"is_deterministic":     o.IsDeterministic,
			"is_null_call":         o.IsNullCall,
			"name":                 o.Name,
			"parameter_style":      o.ParameterStyle,
			"properties":           o.Properties,
			"return_params":        o.ReturnParams,
			"routine_body":         o.RoutineBody,
			"routine_definition":   o.RoutineDefinition,
			"routine_dependencies": o.RoutineDependencies,
			"schema_name":          o.SchemaName,
			"security_type":        o.SecurityType,
			"specific_name":        o.SpecificName,
			"sql_data_access":      o.SqlDataAccess,
			"sql_path":             o.SqlPath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateFunction) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":         types.StringType,
			"comment":              types.StringType,
			"data_type":            types.StringType,
			"external_language":    types.StringType,
			"external_name":        types.StringType,
			"full_data_type":       types.StringType,
			"input_params":         FunctionParameterInfos{}.Type(ctx),
			"is_deterministic":     types.BoolType,
			"is_null_call":         types.BoolType,
			"name":                 types.StringType,
			"parameter_style":      types.StringType,
			"properties":           types.StringType,
			"return_params":        FunctionParameterInfos{}.Type(ctx),
			"routine_body":         types.StringType,
			"routine_definition":   types.StringType,
			"routine_dependencies": DependencyList{}.Type(ctx),
			"schema_name":          types.StringType,
			"security_type":        types.StringType,
			"specific_name":        types.StringType,
			"sql_data_access":      types.StringType,
			"sql_path":             types.StringType,
		},
	}
}

// GetInputParams returns the value of the InputParams field in CreateFunction as
// a FunctionParameterInfos value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFunction) GetInputParams(ctx context.Context) (FunctionParameterInfos, bool) {
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

// SetInputParams sets the value of the InputParams field in CreateFunction.
func (o *CreateFunction) SetInputParams(ctx context.Context, v FunctionParameterInfos) {
	vs := v.ToObjectValue(ctx)
	o.InputParams = vs
}

// GetReturnParams returns the value of the ReturnParams field in CreateFunction as
// a FunctionParameterInfos value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFunction) GetReturnParams(ctx context.Context) (FunctionParameterInfos, bool) {
	var e FunctionParameterInfos
	if o.ReturnParams.IsNull() || o.ReturnParams.IsUnknown() {
		return e, false
	}
	var v []FunctionParameterInfos
	d := o.ReturnParams.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetReturnParams sets the value of the ReturnParams field in CreateFunction.
func (o *CreateFunction) SetReturnParams(ctx context.Context, v FunctionParameterInfos) {
	vs := v.ToObjectValue(ctx)
	o.ReturnParams = vs
}

// GetRoutineDependencies returns the value of the RoutineDependencies field in CreateFunction as
// a DependencyList value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFunction) GetRoutineDependencies(ctx context.Context) (DependencyList, bool) {
	var e DependencyList
	if o.RoutineDependencies.IsNull() || o.RoutineDependencies.IsUnknown() {
		return e, false
	}
	var v []DependencyList
	d := o.RoutineDependencies.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetRoutineDependencies sets the value of the RoutineDependencies field in CreateFunction.
func (o *CreateFunction) SetRoutineDependencies(ctx context.Context, v DependencyList) {
	vs := v.ToObjectValue(ctx)
	o.RoutineDependencies = vs
}

type CreateFunctionRequest struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	FunctionInfo types.Object `tfsdk:"function_info"`
}

func (newState *CreateFunctionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFunctionRequest) {
}

func (newState *CreateFunctionRequest) SyncEffectiveFieldsDuringRead(existingState CreateFunctionRequest) {
}

func (c CreateFunctionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["function_info"] = attrs["function_info"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFunctionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateFunctionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"function_info": reflect.TypeOf(CreateFunction{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFunctionRequest
// only implements ToObjectValue() and Type().
func (o CreateFunctionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_info": o.FunctionInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateFunctionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_info": CreateFunction{}.Type(ctx),
		},
	}
}

// GetFunctionInfo returns the value of the FunctionInfo field in CreateFunctionRequest as
// a CreateFunction value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFunctionRequest) GetFunctionInfo(ctx context.Context) (CreateFunction, bool) {
	var e CreateFunction
	if o.FunctionInfo.IsNull() || o.FunctionInfo.IsUnknown() {
		return e, false
	}
	var v []CreateFunction
	d := o.FunctionInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetFunctionInfo sets the value of the FunctionInfo field in CreateFunctionRequest.
func (o *CreateFunctionRequest) SetFunctionInfo(ctx context.Context, v CreateFunction) {
	vs := v.ToObjectValue(ctx)
	o.FunctionInfo = vs
}

type CreateMetastore struct {
	// The user-specified name of the metastore.
	Name types.String `tfsdk:"name"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// The field can be omitted in the __workspace-level__ __API__ but not in
	// the __account-level__ __API__. If this field is omitted, the region of
	// the workspace receiving the request will be used.
	Region types.String `tfsdk:"region"`
	// The storage root URL for metastore
	StorageRoot types.String `tfsdk:"storage_root"`
}

func (newState *CreateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMetastore) {
}

func (newState *CreateMetastore) SyncEffectiveFieldsDuringRead(existingState CreateMetastore) {
}

func (c CreateMetastore) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateMetastore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateMetastore) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateMetastore
// only implements ToObjectValue() and Type().
func (o CreateMetastore) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":         o.Name,
			"region":       o.Region,
			"storage_root": o.StorageRoot,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateMetastore) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":         types.StringType,
			"region":       types.StringType,
			"storage_root": types.StringType,
		},
	}
}

type CreateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name"`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *CreateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMetastoreAssignment) {
}

func (newState *CreateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState CreateMetastoreAssignment) {
}

func (c CreateMetastoreAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_catalog_name"] = attrs["default_catalog_name"].SetRequired()
	attrs["metastore_id"] = attrs["metastore_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateMetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateMetastoreAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateMetastoreAssignment
// only implements ToObjectValue() and Type().
func (o CreateMetastoreAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_catalog_name": o.DefaultCatalogName,
			"metastore_id":         o.MetastoreId,
			"workspace_id":         o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateMetastoreAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_catalog_name": types.StringType,
			"metastore_id":         types.StringType,
			"workspace_id":         types.Int64Type,
		},
	}
}

type CreateMonitor struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir types.String `tfsdk:"assets_dir"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName types.String `tfsdk:"baseline_table_name"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics types.List `tfsdk:"custom_metrics"`
	// The data classification config for the monitor.
	DataClassificationConfig types.Object `tfsdk:"data_classification_config"`
	// Configuration for monitoring inference logs.
	InferenceLog types.Object `tfsdk:"inference_log"`
	// The notification settings for the monitor.
	Notifications types.Object `tfsdk:"notifications"`
	// Schema where output metric tables are created.
	OutputSchemaName types.String `tfsdk:"output_schema_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule types.Object `tfsdk:"schedule"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard types.Bool `tfsdk:"skip_builtin_dashboard"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs types.List `tfsdk:"slicing_exprs"`
	// Configuration for monitoring snapshot tables.
	Snapshot types.Object `tfsdk:"snapshot"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries types.Object `tfsdk:"time_series"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *CreateMonitor) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMonitor) {
}

func (newState *CreateMonitor) SyncEffectiveFieldsDuringRead(existingState CreateMonitor) {
}

func (c CreateMonitor) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assets_dir"] = attrs["assets_dir"].SetRequired()
	attrs["baseline_table_name"] = attrs["baseline_table_name"].SetOptional()
	attrs["custom_metrics"] = attrs["custom_metrics"].SetOptional()
	attrs["data_classification_config"] = attrs["data_classification_config"].SetOptional()
	attrs["inference_log"] = attrs["inference_log"].SetOptional()
	attrs["notifications"] = attrs["notifications"].SetOptional()
	attrs["output_schema_name"] = attrs["output_schema_name"].SetRequired()
	attrs["schedule"] = attrs["schedule"].SetOptional()
	attrs["skip_builtin_dashboard"] = attrs["skip_builtin_dashboard"].SetOptional()
	attrs["slicing_exprs"] = attrs["slicing_exprs"].SetOptional()
	attrs["snapshot"] = attrs["snapshot"].SetOptional()
	attrs["table_name"] = attrs["table_name"].SetRequired()
	attrs["time_series"] = attrs["time_series"].SetOptional()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateMonitor.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateMonitor) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_metrics":             reflect.TypeOf(MonitorMetric{}),
		"data_classification_config": reflect.TypeOf(MonitorDataClassificationConfig{}),
		"inference_log":              reflect.TypeOf(MonitorInferenceLog{}),
		"notifications":              reflect.TypeOf(MonitorNotifications{}),
		"schedule":                   reflect.TypeOf(MonitorCronSchedule{}),
		"slicing_exprs":              reflect.TypeOf(types.String{}),
		"snapshot":                   reflect.TypeOf(MonitorSnapshot{}),
		"time_series":                reflect.TypeOf(MonitorTimeSeries{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateMonitor
// only implements ToObjectValue() and Type().
func (o CreateMonitor) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets_dir":                 o.AssetsDir,
			"baseline_table_name":        o.BaselineTableName,
			"custom_metrics":             o.CustomMetrics,
			"data_classification_config": o.DataClassificationConfig,
			"inference_log":              o.InferenceLog,
			"notifications":              o.Notifications,
			"output_schema_name":         o.OutputSchemaName,
			"schedule":                   o.Schedule,
			"skip_builtin_dashboard":     o.SkipBuiltinDashboard,
			"slicing_exprs":              o.SlicingExprs,
			"snapshot":                   o.Snapshot,
			"table_name":                 o.TableName,
			"time_series":                o.TimeSeries,
			"warehouse_id":               o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateMonitor) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets_dir":          types.StringType,
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: MonitorMetric{}.Type(ctx),
			},
			"data_classification_config": MonitorDataClassificationConfig{}.Type(ctx),
			"inference_log":              MonitorInferenceLog{}.Type(ctx),
			"notifications":              MonitorNotifications{}.Type(ctx),
			"output_schema_name":         types.StringType,
			"schedule":                   MonitorCronSchedule{}.Type(ctx),
			"skip_builtin_dashboard":     types.BoolType,
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot":     MonitorSnapshot{}.Type(ctx),
			"table_name":   types.StringType,
			"time_series":  MonitorTimeSeries{}.Type(ctx),
			"warehouse_id": types.StringType,
		},
	}
}

// GetCustomMetrics returns the value of the CustomMetrics field in CreateMonitor as
// a slice of MonitorMetric values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor) GetCustomMetrics(ctx context.Context) ([]MonitorMetric, bool) {
	if o.CustomMetrics.IsNull() || o.CustomMetrics.IsUnknown() {
		return nil, false
	}
	var v []MonitorMetric
	d := o.CustomMetrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomMetrics sets the value of the CustomMetrics field in CreateMonitor.
func (o *CreateMonitor) SetCustomMetrics(ctx context.Context, v []MonitorMetric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomMetrics = types.ListValueMust(t, vs)
}

// GetDataClassificationConfig returns the value of the DataClassificationConfig field in CreateMonitor as
// a MonitorDataClassificationConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor) GetDataClassificationConfig(ctx context.Context) (MonitorDataClassificationConfig, bool) {
	var e MonitorDataClassificationConfig
	if o.DataClassificationConfig.IsNull() || o.DataClassificationConfig.IsUnknown() {
		return e, false
	}
	var v []MonitorDataClassificationConfig
	d := o.DataClassificationConfig.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDataClassificationConfig sets the value of the DataClassificationConfig field in CreateMonitor.
func (o *CreateMonitor) SetDataClassificationConfig(ctx context.Context, v MonitorDataClassificationConfig) {
	vs := v.ToObjectValue(ctx)
	o.DataClassificationConfig = vs
}

// GetInferenceLog returns the value of the InferenceLog field in CreateMonitor as
// a MonitorInferenceLog value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor) GetInferenceLog(ctx context.Context) (MonitorInferenceLog, bool) {
	var e MonitorInferenceLog
	if o.InferenceLog.IsNull() || o.InferenceLog.IsUnknown() {
		return e, false
	}
	var v []MonitorInferenceLog
	d := o.InferenceLog.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetInferenceLog sets the value of the InferenceLog field in CreateMonitor.
func (o *CreateMonitor) SetInferenceLog(ctx context.Context, v MonitorInferenceLog) {
	vs := v.ToObjectValue(ctx)
	o.InferenceLog = vs
}

// GetNotifications returns the value of the Notifications field in CreateMonitor as
// a MonitorNotifications value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor) GetNotifications(ctx context.Context) (MonitorNotifications, bool) {
	var e MonitorNotifications
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return e, false
	}
	var v []MonitorNotifications
	d := o.Notifications.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNotifications sets the value of the Notifications field in CreateMonitor.
func (o *CreateMonitor) SetNotifications(ctx context.Context, v MonitorNotifications) {
	vs := v.ToObjectValue(ctx)
	o.Notifications = vs
}

// GetSchedule returns the value of the Schedule field in CreateMonitor as
// a MonitorCronSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor) GetSchedule(ctx context.Context) (MonitorCronSchedule, bool) {
	var e MonitorCronSchedule
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []MonitorCronSchedule
	d := o.Schedule.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSchedule sets the value of the Schedule field in CreateMonitor.
func (o *CreateMonitor) SetSchedule(ctx context.Context, v MonitorCronSchedule) {
	vs := v.ToObjectValue(ctx)
	o.Schedule = vs
}

// GetSlicingExprs returns the value of the SlicingExprs field in CreateMonitor as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor) GetSlicingExprs(ctx context.Context) ([]types.String, bool) {
	if o.SlicingExprs.IsNull() || o.SlicingExprs.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SlicingExprs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSlicingExprs sets the value of the SlicingExprs field in CreateMonitor.
func (o *CreateMonitor) SetSlicingExprs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["slicing_exprs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SlicingExprs = types.ListValueMust(t, vs)
}

// GetSnapshot returns the value of the Snapshot field in CreateMonitor as
// a MonitorSnapshot value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor) GetSnapshot(ctx context.Context) (MonitorSnapshot, bool) {
	var e MonitorSnapshot
	if o.Snapshot.IsNull() || o.Snapshot.IsUnknown() {
		return e, false
	}
	var v []MonitorSnapshot
	d := o.Snapshot.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSnapshot sets the value of the Snapshot field in CreateMonitor.
func (o *CreateMonitor) SetSnapshot(ctx context.Context, v MonitorSnapshot) {
	vs := v.ToObjectValue(ctx)
	o.Snapshot = vs
}

// GetTimeSeries returns the value of the TimeSeries field in CreateMonitor as
// a MonitorTimeSeries value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor) GetTimeSeries(ctx context.Context) (MonitorTimeSeries, bool) {
	var e MonitorTimeSeries
	if o.TimeSeries.IsNull() || o.TimeSeries.IsUnknown() {
		return e, false
	}
	var v []MonitorTimeSeries
	d := o.TimeSeries.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTimeSeries sets the value of the TimeSeries field in CreateMonitor.
func (o *CreateMonitor) SetTimeSeries(ctx context.Context, v MonitorTimeSeries) {
	vs := v.ToObjectValue(ctx)
	o.TimeSeries = vs
}

// Create an Online Table
type CreateOnlineTableRequest struct {
	// Online Table information.
	Table types.Object `tfsdk:"table"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateOnlineTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateOnlineTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table": reflect.TypeOf(OnlineTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOnlineTableRequest
// only implements ToObjectValue() and Type().
func (o CreateOnlineTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table": o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateOnlineTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table": OnlineTable{}.Type(ctx),
		},
	}
}

// GetTable returns the value of the Table field in CreateOnlineTableRequest as
// a OnlineTable value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateOnlineTableRequest) GetTable(ctx context.Context) (OnlineTable, bool) {
	var e OnlineTable
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []OnlineTable
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

// SetTable sets the value of the Table field in CreateOnlineTableRequest.
func (o *CreateOnlineTableRequest) SetTable(ctx context.Context, v OnlineTable) {
	vs := v.ToObjectValue(ctx)
	o.Table = vs
}

type CreateRegisteredModelRequest struct {
	// The name of the catalog where the schema and the registered model reside
	CatalogName types.String `tfsdk:"catalog_name"`
	// The comment attached to the registered model
	Comment types.String `tfsdk:"comment"`
	// The name of the registered model
	Name types.String `tfsdk:"name"`
	// The name of the schema where the registered model resides
	SchemaName types.String `tfsdk:"schema_name"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation types.String `tfsdk:"storage_location"`
}

func (newState *CreateRegisteredModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRegisteredModelRequest) {
}

func (newState *CreateRegisteredModelRequest) SyncEffectiveFieldsDuringRead(existingState CreateRegisteredModelRequest) {
}

func (c CreateRegisteredModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetRequired()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRegisteredModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateRegisteredModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRegisteredModelRequest
// only implements ToObjectValue() and Type().
func (o CreateRegisteredModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":     o.CatalogName,
			"comment":          o.Comment,
			"name":             o.Name,
			"schema_name":      o.SchemaName,
			"storage_location": o.StorageLocation,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateRegisteredModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":     types.StringType,
			"comment":          types.StringType,
			"name":             types.StringType,
			"schema_name":      types.StringType,
			"storage_location": types.StringType,
		},
	}
}

type CreateResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateResponse
// only implements ToObjectValue() and Type().
func (o CreateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CreateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CreateSchema struct {
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Name of schema, relative to parent catalog.
	Name types.String `tfsdk:"name"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`
	// Storage root URL for managed tables within schema.
	StorageRoot types.String `tfsdk:"storage_root"`
}

func (newState *CreateSchema) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateSchema) {
}

func (newState *CreateSchema) SyncEffectiveFieldsDuringRead(existingState CreateSchema) {
}

func (c CreateSchema) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSchema.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateSchema) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSchema
// only implements ToObjectValue() and Type().
func (o CreateSchema) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name": o.CatalogName,
			"comment":      o.Comment,
			"name":         o.Name,
			"properties":   o.Properties,
			"storage_root": o.StorageRoot,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateSchema) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name": types.StringType,
			"comment":      types.StringType,
			"name":         types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"storage_root": types.StringType,
		},
	}
}

// GetProperties returns the value of the Properties field in CreateSchema as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateSchema) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in CreateSchema.
func (o *CreateSchema) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole types.Object `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.Object `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.Object `tfsdk:"azure_service_principal"`
	// The Cloudflare API token configuration.
	CloudflareApiToken types.Object `tfsdk:"cloudflare_api_token"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount types.Object `tfsdk:"databricks_gcp_service_account"`
	// The credential name. The name must be unique within the metastore.
	Name types.String `tfsdk:"name"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// Supplying true to this argument skips validation of the created
	// credential.
	SkipValidation types.Bool `tfsdk:"skip_validation"`
}

func (newState *CreateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateStorageCredential) {
}

func (newState *CreateStorageCredential) SyncEffectiveFieldsDuringRead(existingState CreateStorageCredential) {
}

func (c CreateStorageCredential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_iam_role"] = attrs["aws_iam_role"].SetOptional()
	attrs["azure_managed_identity"] = attrs["azure_managed_identity"].SetOptional()
	attrs["azure_service_principal"] = attrs["azure_service_principal"].SetOptional()
	attrs["cloudflare_api_token"] = attrs["cloudflare_api_token"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["databricks_gcp_service_account"] = attrs["databricks_gcp_service_account"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["skip_validation"] = attrs["skip_validation"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateStorageCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateStorageCredential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRoleRequest{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentityRequest{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal{}),
		"cloudflare_api_token":           reflect.TypeOf(CloudflareApiToken{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccountRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateStorageCredential
// only implements ToObjectValue() and Type().
func (o CreateStorageCredential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_iam_role":                   o.AwsIamRole,
			"azure_managed_identity":         o.AzureManagedIdentity,
			"azure_service_principal":        o.AzureServicePrincipal,
			"cloudflare_api_token":           o.CloudflareApiToken,
			"comment":                        o.Comment,
			"databricks_gcp_service_account": o.DatabricksGcpServiceAccount,
			"name":                           o.Name,
			"read_only":                      o.ReadOnly,
			"skip_validation":                o.SkipValidation,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateStorageCredential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role":                   AwsIamRoleRequest{}.Type(ctx),
			"azure_managed_identity":         AzureManagedIdentityRequest{}.Type(ctx),
			"azure_service_principal":        AzureServicePrincipal{}.Type(ctx),
			"cloudflare_api_token":           CloudflareApiToken{}.Type(ctx),
			"comment":                        types.StringType,
			"databricks_gcp_service_account": DatabricksGcpServiceAccountRequest{}.Type(ctx),
			"name":                           types.StringType,
			"read_only":                      types.BoolType,
			"skip_validation":                types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in CreateStorageCredential as
// a AwsIamRoleRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageCredential) GetAwsIamRole(ctx context.Context) (AwsIamRoleRequest, bool) {
	var e AwsIamRoleRequest
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRoleRequest
	d := o.AwsIamRole.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsIamRole sets the value of the AwsIamRole field in CreateStorageCredential.
func (o *CreateStorageCredential) SetAwsIamRole(ctx context.Context, v AwsIamRoleRequest) {
	vs := v.ToObjectValue(ctx)
	o.AwsIamRole = vs
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in CreateStorageCredential as
// a AzureManagedIdentityRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageCredential) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentityRequest, bool) {
	var e AzureManagedIdentityRequest
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentityRequest
	d := o.AzureManagedIdentity.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in CreateStorageCredential.
func (o *CreateStorageCredential) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentityRequest) {
	vs := v.ToObjectValue(ctx)
	o.AzureManagedIdentity = vs
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in CreateStorageCredential as
// a AzureServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageCredential) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal, bool) {
	var e AzureServicePrincipal
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal
	d := o.AzureServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in CreateStorageCredential.
func (o *CreateStorageCredential) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	o.AzureServicePrincipal = vs
}

// GetCloudflareApiToken returns the value of the CloudflareApiToken field in CreateStorageCredential as
// a CloudflareApiToken value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageCredential) GetCloudflareApiToken(ctx context.Context) (CloudflareApiToken, bool) {
	var e CloudflareApiToken
	if o.CloudflareApiToken.IsNull() || o.CloudflareApiToken.IsUnknown() {
		return e, false
	}
	var v []CloudflareApiToken
	d := o.CloudflareApiToken.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetCloudflareApiToken sets the value of the CloudflareApiToken field in CreateStorageCredential.
func (o *CreateStorageCredential) SetCloudflareApiToken(ctx context.Context, v CloudflareApiToken) {
	vs := v.ToObjectValue(ctx)
	o.CloudflareApiToken = vs
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in CreateStorageCredential as
// a DatabricksGcpServiceAccountRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageCredential) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccountRequest, bool) {
	var e DatabricksGcpServiceAccountRequest
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccountRequest
	d := o.DatabricksGcpServiceAccount.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in CreateStorageCredential.
func (o *CreateStorageCredential) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccountRequest) {
	vs := v.ToObjectValue(ctx)
	o.DatabricksGcpServiceAccount = vs
}

// Create a Synced Database Table
type CreateSyncedDatabaseTableRequest struct {
	// Next field marker: 10
	SyncedTable types.Object `tfsdk:"synced_table"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_table": reflect.TypeOf(SyncedDatabaseTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (o CreateSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"synced_table": o.SyncedTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"synced_table": SyncedDatabaseTable{}.Type(ctx),
		},
	}
}

// GetSyncedTable returns the value of the SyncedTable field in CreateSyncedDatabaseTableRequest as
// a SyncedDatabaseTable value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateSyncedDatabaseTableRequest) GetSyncedTable(ctx context.Context) (SyncedDatabaseTable, bool) {
	var e SyncedDatabaseTable
	if o.SyncedTable.IsNull() || o.SyncedTable.IsUnknown() {
		return e, false
	}
	var v []SyncedDatabaseTable
	d := o.SyncedTable.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSyncedTable sets the value of the SyncedTable field in CreateSyncedDatabaseTableRequest.
func (o *CreateSyncedDatabaseTableRequest) SetSyncedTable(ctx context.Context, v SyncedDatabaseTable) {
	vs := v.ToObjectValue(ctx)
	o.SyncedTable = vs
}

type CreateTableConstraint struct {
	// A table constraint, as defined by *one* of the following fields being
	// set: __primary_key_constraint__, __foreign_key_constraint__,
	// __named_table_constraint__.
	Constraint types.Object `tfsdk:"constraint"`
	// The full name of the table referenced by the constraint.
	FullNameArg types.String `tfsdk:"full_name_arg"`
}

func (newState *CreateTableConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTableConstraint) {
}

func (newState *CreateTableConstraint) SyncEffectiveFieldsDuringRead(existingState CreateTableConstraint) {
}

func (c CreateTableConstraint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["constraint"] = attrs["constraint"].SetRequired()
	attrs["full_name_arg"] = attrs["full_name_arg"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTableConstraint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateTableConstraint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"constraint": reflect.TypeOf(TableConstraint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTableConstraint
// only implements ToObjectValue() and Type().
func (o CreateTableConstraint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"constraint":    o.Constraint,
			"full_name_arg": o.FullNameArg,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTableConstraint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"constraint":    TableConstraint{}.Type(ctx),
			"full_name_arg": types.StringType,
		},
	}
}

// GetConstraint returns the value of the Constraint field in CreateTableConstraint as
// a TableConstraint value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateTableConstraint) GetConstraint(ctx context.Context) (TableConstraint, bool) {
	var e TableConstraint
	if o.Constraint.IsNull() || o.Constraint.IsUnknown() {
		return e, false
	}
	var v []TableConstraint
	d := o.Constraint.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetConstraint sets the value of the Constraint field in CreateTableConstraint.
func (o *CreateTableConstraint) SetConstraint(ctx context.Context, v TableConstraint) {
	vs := v.ToObjectValue(ctx)
	o.Constraint = vs
}

type CreateVolumeRequestContent struct {
	// The name of the catalog where the schema and the volume are
	CatalogName types.String `tfsdk:"catalog_name"`
	// The comment attached to the volume
	Comment types.String `tfsdk:"comment"`
	// The name of the volume
	Name types.String `tfsdk:"name"`
	// The name of the schema where the volume is
	SchemaName types.String `tfsdk:"schema_name"`
	// The storage location on the cloud
	StorageLocation types.String `tfsdk:"storage_location"`
	// The type of the volume. An external volume is located in the specified
	// external location. A managed volume is located in the default location
	// which is specified by the parent schema, or the parent catalog, or the
	// Metastore. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
	VolumeType types.String `tfsdk:"volume_type"`
}

func (newState *CreateVolumeRequestContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVolumeRequestContent) {
}

func (newState *CreateVolumeRequestContent) SyncEffectiveFieldsDuringRead(existingState CreateVolumeRequestContent) {
}

func (c CreateVolumeRequestContent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetRequired()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
	attrs["volume_type"] = attrs["volume_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVolumeRequestContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVolumeRequestContent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVolumeRequestContent
// only implements ToObjectValue() and Type().
func (o CreateVolumeRequestContent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":     o.CatalogName,
			"comment":          o.Comment,
			"name":             o.Name,
			"schema_name":      o.SchemaName,
			"storage_location": o.StorageLocation,
			"volume_type":      o.VolumeType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateVolumeRequestContent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":     types.StringType,
			"comment":          types.StringType,
			"name":             types.StringType,
			"schema_name":      types.StringType,
			"storage_location": types.StringType,
			"volume_type":      types.StringType,
		},
	}
}

type CredentialInfo struct {
	// The AWS IAM role configuration
	AwsIamRole types.Object `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.Object `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal types.Object `tfsdk:"azure_service_principal"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment"`
	// Time at which this credential was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of credential creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount types.Object `tfsdk:"databricks_gcp_service_account"`
	// The full name of the credential.
	FullName types.String `tfsdk:"full_name"`
	// The unique identifier of the credential.
	Id types.String `tfsdk:"id"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode"`
	// Unique identifier of the parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name types.String `tfsdk:"name"`
	// Username of current owner of credential.
	Owner types.String `tfsdk:"owner"`
	// Indicates the purpose of the credential.
	Purpose types.String `tfsdk:"purpose"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified the credential.
	UpdatedBy types.String `tfsdk:"updated_by"`
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	UsedForManagedStorage types.Bool `tfsdk:"used_for_managed_storage"`
}

func (newState *CredentialInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CredentialInfo) {
}

func (newState *CredentialInfo) SyncEffectiveFieldsDuringRead(existingState CredentialInfo) {
}

func (c CredentialInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_iam_role"] = attrs["aws_iam_role"].SetOptional()
	attrs["azure_managed_identity"] = attrs["azure_managed_identity"].SetOptional()
	attrs["azure_service_principal"] = attrs["azure_service_principal"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["databricks_gcp_service_account"] = attrs["databricks_gcp_service_account"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["isolation_mode"] = attrs["isolation_mode"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["purpose"] = attrs["purpose"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()
	attrs["used_for_managed_storage"] = attrs["used_for_managed_storage"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CredentialInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CredentialInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRole{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentity{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CredentialInfo
// only implements ToObjectValue() and Type().
func (o CredentialInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_iam_role":                   o.AwsIamRole,
			"azure_managed_identity":         o.AzureManagedIdentity,
			"azure_service_principal":        o.AzureServicePrincipal,
			"comment":                        o.Comment,
			"created_at":                     o.CreatedAt,
			"created_by":                     o.CreatedBy,
			"databricks_gcp_service_account": o.DatabricksGcpServiceAccount,
			"full_name":                      o.FullName,
			"id":                             o.Id,
			"isolation_mode":                 o.IsolationMode,
			"metastore_id":                   o.MetastoreId,
			"name":                           o.Name,
			"owner":                          o.Owner,
			"purpose":                        o.Purpose,
			"read_only":                      o.ReadOnly,
			"updated_at":                     o.UpdatedAt,
			"updated_by":                     o.UpdatedBy,
			"used_for_managed_storage":       o.UsedForManagedStorage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CredentialInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role":                   AwsIamRole{}.Type(ctx),
			"azure_managed_identity":         AzureManagedIdentity{}.Type(ctx),
			"azure_service_principal":        AzureServicePrincipal{}.Type(ctx),
			"comment":                        types.StringType,
			"created_at":                     types.Int64Type,
			"created_by":                     types.StringType,
			"databricks_gcp_service_account": DatabricksGcpServiceAccount{}.Type(ctx),
			"full_name":                      types.StringType,
			"id":                             types.StringType,
			"isolation_mode":                 types.StringType,
			"metastore_id":                   types.StringType,
			"name":                           types.StringType,
			"owner":                          types.StringType,
			"purpose":                        types.StringType,
			"read_only":                      types.BoolType,
			"updated_at":                     types.Int64Type,
			"updated_by":                     types.StringType,
			"used_for_managed_storage":       types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in CredentialInfo as
// a AwsIamRole value.
// If the field is unknown or null, the boolean return value is false.
func (o *CredentialInfo) GetAwsIamRole(ctx context.Context) (AwsIamRole, bool) {
	var e AwsIamRole
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRole
	d := o.AwsIamRole.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsIamRole sets the value of the AwsIamRole field in CredentialInfo.
func (o *CredentialInfo) SetAwsIamRole(ctx context.Context, v AwsIamRole) {
	vs := v.ToObjectValue(ctx)
	o.AwsIamRole = vs
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in CredentialInfo as
// a AzureManagedIdentity value.
// If the field is unknown or null, the boolean return value is false.
func (o *CredentialInfo) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentity, bool) {
	var e AzureManagedIdentity
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentity
	d := o.AzureManagedIdentity.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in CredentialInfo.
func (o *CredentialInfo) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentity) {
	vs := v.ToObjectValue(ctx)
	o.AzureManagedIdentity = vs
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in CredentialInfo as
// a AzureServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (o *CredentialInfo) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal, bool) {
	var e AzureServicePrincipal
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal
	d := o.AzureServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in CredentialInfo.
func (o *CredentialInfo) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	o.AzureServicePrincipal = vs
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in CredentialInfo as
// a DatabricksGcpServiceAccount value.
// If the field is unknown or null, the boolean return value is false.
func (o *CredentialInfo) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccount, bool) {
	var e DatabricksGcpServiceAccount
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccount
	d := o.DatabricksGcpServiceAccount.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in CredentialInfo.
func (o *CredentialInfo) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccount) {
	vs := v.ToObjectValue(ctx)
	o.DatabricksGcpServiceAccount = vs
}

type CredentialValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message types.String `tfsdk:"message"`
	// The results of the tested operation.
	Result types.String `tfsdk:"result"`
}

func (newState *CredentialValidationResult) SyncEffectiveFieldsDuringCreateOrUpdate(plan CredentialValidationResult) {
}

func (newState *CredentialValidationResult) SyncEffectiveFieldsDuringRead(existingState CredentialValidationResult) {
}

func (c CredentialValidationResult) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetOptional()
	attrs["result"] = attrs["result"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CredentialValidationResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CredentialValidationResult) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CredentialValidationResult
// only implements ToObjectValue() and Type().
func (o CredentialValidationResult) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
			"result":  o.Result,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CredentialValidationResult) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"result":  types.StringType,
		},
	}
}

type DatabaseCatalog struct {
	CreateDatabaseIfNotExists types.Bool `tfsdk:"create_database_if_not_exists"`
	// The name of the DatabaseInstance housing the database.
	DatabaseInstanceName types.String `tfsdk:"database_instance_name"`
	// The name of the database (in a instance) associated with the catalog.
	DatabaseName types.String `tfsdk:"database_name"`
	// The name of the catalog in UC.
	Name types.String `tfsdk:"name"`

	Uid types.String `tfsdk:"uid"`
}

func (newState *DatabaseCatalog) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabaseCatalog) {
}

func (newState *DatabaseCatalog) SyncEffectiveFieldsDuringRead(existingState DatabaseCatalog) {
}

func (c DatabaseCatalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].SetOptional()
	attrs["database_instance_name"] = attrs["database_instance_name"].SetRequired()
	attrs["database_name"] = attrs["database_name"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["uid"] = attrs["uid"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseCatalog.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabaseCatalog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCatalog
// only implements ToObjectValue() and Type().
func (o DatabaseCatalog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_database_if_not_exists": o.CreateDatabaseIfNotExists,
			"database_instance_name":        o.DatabaseInstanceName,
			"database_name":                 o.DatabaseName,
			"name":                          o.Name,
			"uid":                           o.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseCatalog) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_database_if_not_exists": types.BoolType,
			"database_instance_name":        types.StringType,
			"database_name":                 types.StringType,
			"name":                          types.StringType,
			"uid":                           types.StringType,
		},
	}
}

// A DatabaseInstance represents a logical Postgres instance, comprised of both
// compute and storage.
type DatabaseInstance struct {
	// Password for admin user to create. If not provided, no user will be
	// created.
	AdminPassword types.String `tfsdk:"admin_password"`
	// Name of the admin role for the instance. If not provided, defaults to
	// 'databricks_admin'.
	AdminRolename types.String `tfsdk:"admin_rolename"`
	// The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4".
	Capacity types.String `tfsdk:"capacity"`
	// The timestamp when the instance was created.
	CreationTime types.String `tfsdk:"creation_time"`
	// The email of the creator of the instance.
	Creator types.String `tfsdk:"creator"`
	// The name of the instance. This is the unique identifier for the instance.
	Name types.String `tfsdk:"name"`
	// The version of Postgres running on the instance.
	PgVersion types.String `tfsdk:"pg_version"`
	// The DNS endpoint to connect to the instance for read+write access.
	ReadWriteDns types.String `tfsdk:"read_write_dns"`
	// The current state of the instance.
	State types.String `tfsdk:"state"`
	// Whether the instance is stopped.
	Stopped types.Bool `tfsdk:"stopped"`
	// An immutable UUID identifier for the instance.
	Uid types.String `tfsdk:"uid"`
}

func (newState *DatabaseInstance) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabaseInstance) {
}

func (newState *DatabaseInstance) SyncEffectiveFieldsDuringRead(existingState DatabaseInstance) {
}

func (c DatabaseInstance) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["admin_password"] = attrs["admin_password"].SetOptional()
	attrs["admin_rolename"] = attrs["admin_rolename"].SetOptional()
	attrs["capacity"] = attrs["capacity"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["pg_version"] = attrs["pg_version"].SetComputed()
	attrs["read_write_dns"] = attrs["read_write_dns"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["stopped"] = attrs["stopped"].SetOptional()
	attrs["uid"] = attrs["uid"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseInstance.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabaseInstance) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstance
// only implements ToObjectValue() and Type().
func (o DatabaseInstance) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"admin_password": o.AdminPassword,
			"admin_rolename": o.AdminRolename,
			"capacity":       o.Capacity,
			"creation_time":  o.CreationTime,
			"creator":        o.Creator,
			"name":           o.Name,
			"pg_version":     o.PgVersion,
			"read_write_dns": o.ReadWriteDns,
			"state":          o.State,
			"stopped":        o.Stopped,
			"uid":            o.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseInstance) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"admin_password": types.StringType,
			"admin_rolename": types.StringType,
			"capacity":       types.StringType,
			"creation_time":  types.StringType,
			"creator":        types.StringType,
			"name":           types.StringType,
			"pg_version":     types.StringType,
			"read_write_dns": types.StringType,
			"state":          types.StringType,
			"stopped":        types.BoolType,
			"uid":            types.StringType,
		},
	}
}

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccount struct {
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database
	CredentialId types.String `tfsdk:"credential_id"`
	// The email of the service account.
	Email types.String `tfsdk:"email"`
	// The ID that represents the private key for this Service Account
	PrivateKeyId types.String `tfsdk:"private_key_id"`
}

func (newState *DatabricksGcpServiceAccount) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksGcpServiceAccount) {
}

func (newState *DatabricksGcpServiceAccount) SyncEffectiveFieldsDuringRead(existingState DatabricksGcpServiceAccount) {
}

func (c DatabricksGcpServiceAccount) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetComputed()
	attrs["email"] = attrs["email"].SetComputed()
	attrs["private_key_id"] = attrs["private_key_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksGcpServiceAccount.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabricksGcpServiceAccount) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksGcpServiceAccount
// only implements ToObjectValue() and Type().
func (o DatabricksGcpServiceAccount) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":  o.CredentialId,
			"email":          o.Email,
			"private_key_id": o.PrivateKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabricksGcpServiceAccount) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id":  types.StringType,
			"email":          types.StringType,
			"private_key_id": types.StringType,
		},
	}
}

type DatabricksGcpServiceAccountRequest struct {
}

func (newState *DatabricksGcpServiceAccountRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksGcpServiceAccountRequest) {
}

func (newState *DatabricksGcpServiceAccountRequest) SyncEffectiveFieldsDuringRead(existingState DatabricksGcpServiceAccountRequest) {
}

func (c DatabricksGcpServiceAccountRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksGcpServiceAccountRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabricksGcpServiceAccountRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksGcpServiceAccountRequest
// only implements ToObjectValue() and Type().
func (o DatabricksGcpServiceAccountRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DatabricksGcpServiceAccountRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DatabricksGcpServiceAccountResponse struct {
	// The Databricks internal ID that represents this service account. This is
	// an output-only field.
	CredentialId types.String `tfsdk:"credential_id"`
	// The email of the service account. This is an output-only field.
	Email types.String `tfsdk:"email"`
}

func (newState *DatabricksGcpServiceAccountResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksGcpServiceAccountResponse) {
}

func (newState *DatabricksGcpServiceAccountResponse) SyncEffectiveFieldsDuringRead(existingState DatabricksGcpServiceAccountResponse) {
}

func (c DatabricksGcpServiceAccountResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetOptional()
	attrs["email"] = attrs["email"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksGcpServiceAccountResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabricksGcpServiceAccountResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksGcpServiceAccountResponse
// only implements ToObjectValue() and Type().
func (o DatabricksGcpServiceAccountResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
			"email":         o.Email,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabricksGcpServiceAccountResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.StringType,
			"email":         types.StringType,
		},
	}
}

// Delete a metastore assignment
type DeleteAccountMetastoreAssignmentRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountMetastoreAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountMetastoreAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountMetastoreAssignmentRequest
// only implements ToObjectValue() and Type().
func (o DeleteAccountMetastoreAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountMetastoreAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// Delete a metastore
type DeleteAccountMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force types.Bool `tfsdk:"-"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountMetastoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountMetastoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountMetastoreRequest
// only implements ToObjectValue() and Type().
func (o DeleteAccountMetastoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force":        o.Force,
			"metastore_id": o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountMetastoreRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force":        types.BoolType,
			"metastore_id": types.StringType,
		},
	}
}

// Delete a storage credential
type DeleteAccountStorageCredentialRequest struct {
	// Force deletion even if the Storage Credential is not empty. Default is
	// false.
	Force types.Bool `tfsdk:"-"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Name of the storage credential.
	StorageCredentialName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountStorageCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountStorageCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountStorageCredentialRequest
// only implements ToObjectValue() and Type().
func (o DeleteAccountStorageCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force":                   o.Force,
			"metastore_id":            o.MetastoreId,
			"storage_credential_name": o.StorageCredentialName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountStorageCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force":                   types.BoolType,
			"metastore_id":            types.StringType,
			"storage_credential_name": types.StringType,
		},
	}
}

// Delete a Registered Model Alias
type DeleteAliasRequest struct {
	// The name of the alias
	Alias types.String `tfsdk:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAliasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAliasRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAliasRequest
// only implements ToObjectValue() and Type().
func (o DeleteAliasRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":     o.Alias,
			"full_name": o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAliasRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":     types.StringType,
			"full_name": types.StringType,
		},
	}
}

type DeleteAliasResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAliasResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAliasResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAliasResponse
// only implements ToObjectValue() and Type().
func (o DeleteAliasResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAliasResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a catalog
type DeleteCatalogRequest struct {
	// Force deletion even if the catalog is not empty.
	Force types.Bool `tfsdk:"-"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCatalogRequest
// only implements ToObjectValue() and Type().
func (o DeleteCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
		},
	}
}

// Delete a connection
type DeleteConnectionRequest struct {
	// The name of the connection to be deleted.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteConnectionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteConnectionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteConnectionRequest
// only implements ToObjectValue() and Type().
func (o DeleteConnectionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteConnectionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Delete a credential
type DeleteCredentialRequest struct {
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force types.Bool `tfsdk:"-"`
	// Name of the credential.
	NameArg types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialRequest
// only implements ToObjectValue() and Type().
func (o DeleteCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force":    o.Force,
			"name_arg": o.NameArg,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force":    types.BoolType,
			"name_arg": types.StringType,
		},
	}
}

type DeleteCredentialResponse struct {
}

func (newState *DeleteCredentialResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCredentialResponse) {
}

func (newState *DeleteCredentialResponse) SyncEffectiveFieldsDuringRead(existingState DeleteCredentialResponse) {
}

func (c DeleteCredentialResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCredentialResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialResponse
// only implements ToObjectValue() and Type().
func (o DeleteCredentialResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a Database Catalog
type DeleteDatabaseCatalogRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteDatabaseCatalogResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseCatalogResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseCatalogResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseCatalogResponse
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseCatalogResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseCatalogResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a Database Instance
type DeleteDatabaseInstanceRequest struct {
	// By default, a instance cannot be deleted if it has descendant instances
	// created via PITR. If this flag is specified as true, all descendent
	// instances will be deleted as well.
	Force types.Bool `tfsdk:"-"`
	// Name of the instance to delete.
	Name types.String `tfsdk:"-"`
	// If false, the database instance is soft deleted. Soft deleted instances
	// behave as if they are deleted, and cannot be used for CRUD operations nor
	// connected to. However they can be undeleted by calling the undelete API
	// for a limited time. If true, the database instance is hard deleted and
	// cannot be undeleted.
	Purge types.Bool `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
			"purge": o.Purge,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
			"purge": types.BoolType,
		},
	}
}

type DeleteDatabaseInstanceResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseInstanceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseInstanceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceResponse
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseInstanceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseInstanceResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete an external location
type DeleteExternalLocationRequest struct {
	// Force deletion even if there are dependent external tables or mounts.
	Force types.Bool `tfsdk:"-"`
	// Name of the external location.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExternalLocationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExternalLocationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExternalLocationRequest
// only implements ToObjectValue() and Type().
func (o DeleteExternalLocationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExternalLocationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
		},
	}
}

// Delete a function
type DeleteFunctionRequest struct {
	// Force deletion even if the function is notempty.
	Force types.Bool `tfsdk:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFunctionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteFunctionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFunctionRequest
// only implements ToObjectValue() and Type().
func (o DeleteFunctionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteFunctionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
		},
	}
}

// Delete a metastore
type DeleteMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force types.Bool `tfsdk:"-"`
	// Unique ID of the metastore.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteMetastoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteMetastoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteMetastoreRequest
// only implements ToObjectValue() and Type().
func (o DeleteMetastoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"id":    o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteMetastoreRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"id":    types.StringType,
		},
	}
}

// Delete a Model Version
type DeleteModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	FullName types.String `tfsdk:"-"`
	// The integer version number of the model version
	Version types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionRequest
// only implements ToObjectValue() and Type().
func (o DeleteModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": o.FullName,
			"version":   o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
			"version":   types.Int64Type,
		},
	}
}

// Delete an Online Table
type DeleteOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteOnlineTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteOnlineTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteOnlineTableRequest
// only implements ToObjectValue() and Type().
func (o DeleteOnlineTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteOnlineTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Delete a table monitor
type DeleteQualityMonitorRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (o DeleteQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteQualityMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name": types.StringType,
		},
	}
}

// Delete a Registered Model
type DeleteRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRegisteredModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRegisteredModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRegisteredModelRequest
// only implements ToObjectValue() and Type().
func (o DeleteRegisteredModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRegisteredModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
		},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

func (c DeleteResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

// Delete a schema
type DeleteSchemaRequest struct {
	// Force deletion even if the schema is not empty.
	Force types.Bool `tfsdk:"-"`
	// Full name of the schema.
	FullName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSchemaRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSchemaRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSchemaRequest
// only implements ToObjectValue() and Type().
func (o DeleteSchemaRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force":     o.Force,
			"full_name": o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSchemaRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force":     types.BoolType,
			"full_name": types.StringType,
		},
	}
}

// Delete a credential
type DeleteStorageCredentialRequest struct {
	// Force deletion even if there are dependent external locations or external
	// tables.
	Force types.Bool `tfsdk:"-"`
	// Name of the storage credential.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteStorageCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteStorageCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteStorageCredentialRequest
// only implements ToObjectValue() and Type().
func (o DeleteStorageCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteStorageCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
		},
	}
}

// Delete a Synced Database Table
type DeleteSyncedDatabaseTableRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (o DeleteSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteSyncedDatabaseTableResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSyncedDatabaseTableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSyncedDatabaseTableResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSyncedDatabaseTableResponse
// only implements ToObjectValue() and Type().
func (o DeleteSyncedDatabaseTableResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSyncedDatabaseTableResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a table constraint
type DeleteTableConstraintRequest struct {
	// If true, try deleting all child constraints of the current constraint. If
	// false, reject this operation if the current constraint has any child
	// constraints.
	Cascade types.Bool `tfsdk:"-"`
	// The name of the constraint to delete.
	ConstraintName types.String `tfsdk:"-"`
	// Full name of the table referenced by the constraint.
	FullName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTableConstraintRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTableConstraintRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTableConstraintRequest
// only implements ToObjectValue() and Type().
func (o DeleteTableConstraintRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cascade":         o.Cascade,
			"constraint_name": o.ConstraintName,
			"full_name":       o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTableConstraintRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cascade":         types.BoolType,
			"constraint_name": types.StringType,
			"full_name":       types.StringType,
		},
	}
}

// Delete a table
type DeleteTableRequest struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTableRequest
// only implements ToObjectValue() and Type().
func (o DeleteTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
		},
	}
}

// Delete a Volume
type DeleteVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteVolumeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteVolumeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteVolumeRequest
// only implements ToObjectValue() and Type().
func (o DeleteVolumeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteVolumeRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Properties pertaining to the current state of the delta table as given by the
// commit server. This does not contain **delta.*** (input) properties in
// __TableInfo.properties__.
type DeltaRuntimePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	DeltaRuntimeProperties types.Map `tfsdk:"delta_runtime_properties"`
}

func (newState *DeltaRuntimePropertiesKvPairs) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaRuntimePropertiesKvPairs) {
}

func (newState *DeltaRuntimePropertiesKvPairs) SyncEffectiveFieldsDuringRead(existingState DeltaRuntimePropertiesKvPairs) {
}

func (c DeltaRuntimePropertiesKvPairs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["delta_runtime_properties"] = attrs["delta_runtime_properties"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaRuntimePropertiesKvPairs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeltaRuntimePropertiesKvPairs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_runtime_properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaRuntimePropertiesKvPairs
// only implements ToObjectValue() and Type().
func (o DeltaRuntimePropertiesKvPairs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_runtime_properties": o.DeltaRuntimeProperties,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaRuntimePropertiesKvPairs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_runtime_properties": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDeltaRuntimeProperties returns the value of the DeltaRuntimeProperties field in DeltaRuntimePropertiesKvPairs as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *DeltaRuntimePropertiesKvPairs) GetDeltaRuntimeProperties(ctx context.Context) (map[string]types.String, bool) {
	if o.DeltaRuntimeProperties.IsNull() || o.DeltaRuntimeProperties.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.DeltaRuntimeProperties.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeltaRuntimeProperties sets the value of the DeltaRuntimeProperties field in DeltaRuntimePropertiesKvPairs.
func (o *DeltaRuntimePropertiesKvPairs) SetDeltaRuntimeProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["delta_runtime_properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DeltaRuntimeProperties = types.MapValueMust(t, vs)
}

// A dependency of a SQL object. Either the __table__ field or the __function__
// field must be defined.
type Dependency struct {
	// A function that is dependent on a SQL object.
	Function types.Object `tfsdk:"function"`
	// A table that is dependent on a SQL object.
	Table types.Object `tfsdk:"table"`
}

func (newState *Dependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dependency) {
}

func (newState *Dependency) SyncEffectiveFieldsDuringRead(existingState Dependency) {
}

func (c Dependency) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["function"] = attrs["function"].SetOptional()
	attrs["table"] = attrs["table"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Dependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Dependency) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"function": reflect.TypeOf(FunctionDependency{}),
		"table":    reflect.TypeOf(TableDependency{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dependency
// only implements ToObjectValue() and Type().
func (o Dependency) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function": o.Function,
			"table":    o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Dependency) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function": FunctionDependency{}.Type(ctx),
			"table":    TableDependency{}.Type(ctx),
		},
	}
}

// GetFunction returns the value of the Function field in Dependency as
// a FunctionDependency value.
// If the field is unknown or null, the boolean return value is false.
func (o *Dependency) GetFunction(ctx context.Context) (FunctionDependency, bool) {
	var e FunctionDependency
	if o.Function.IsNull() || o.Function.IsUnknown() {
		return e, false
	}
	var v []FunctionDependency
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

// SetFunction sets the value of the Function field in Dependency.
func (o *Dependency) SetFunction(ctx context.Context, v FunctionDependency) {
	vs := v.ToObjectValue(ctx)
	o.Function = vs
}

// GetTable returns the value of the Table field in Dependency as
// a TableDependency value.
// If the field is unknown or null, the boolean return value is false.
func (o *Dependency) GetTable(ctx context.Context) (TableDependency, bool) {
	var e TableDependency
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []TableDependency
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

// SetTable sets the value of the Table field in Dependency.
func (o *Dependency) SetTable(ctx context.Context, v TableDependency) {
	vs := v.ToObjectValue(ctx)
	o.Table = vs
}

// A list of dependencies.
type DependencyList struct {
	// Array of dependencies.
	Dependencies types.List `tfsdk:"dependencies"`
}

func (newState *DependencyList) SyncEffectiveFieldsDuringCreateOrUpdate(plan DependencyList) {
}

func (newState *DependencyList) SyncEffectiveFieldsDuringRead(existingState DependencyList) {
}

func (c DependencyList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dependencies"] = attrs["dependencies"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DependencyList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DependencyList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies": reflect.TypeOf(Dependency{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DependencyList
// only implements ToObjectValue() and Type().
func (o DependencyList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dependencies": o.Dependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DependencyList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dependencies": basetypes.ListType{
				ElemType: Dependency{}.Type(ctx),
			},
		},
	}
}

// GetDependencies returns the value of the Dependencies field in DependencyList as
// a slice of Dependency values.
// If the field is unknown or null, the boolean return value is false.
func (o *DependencyList) GetDependencies(ctx context.Context) ([]Dependency, bool) {
	if o.Dependencies.IsNull() || o.Dependencies.IsUnknown() {
		return nil, false
	}
	var v []Dependency
	d := o.Dependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependencies sets the value of the Dependencies field in DependencyList.
func (o *DependencyList) SetDependencies(ctx context.Context, v []Dependency) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Dependencies = types.ListValueMust(t, vs)
}

// Disable a system schema
type DisableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId types.String `tfsdk:"-"`
	// Full name of the system schema.
	SchemaName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DisableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DisableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableRequest
// only implements ToObjectValue() and Type().
func (o DisableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
			"schema_name":  o.SchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DisableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
			"schema_name":  types.StringType,
		},
	}
}

type DisableResponse struct {
}

func (newState *DisableResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableResponse) {
}

func (newState *DisableResponse) SyncEffectiveFieldsDuringRead(existingState DisableResponse) {
}

func (c DisableResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DisableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DisableResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableResponse
// only implements ToObjectValue() and Type().
func (o DisableResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DisableResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EffectivePermissionsList struct {
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments"`
}

func (newState *EffectivePermissionsList) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePermissionsList) {
}

func (newState *EffectivePermissionsList) SyncEffectiveFieldsDuringRead(existingState EffectivePermissionsList) {
}

func (c EffectivePermissionsList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["privilege_assignments"] = attrs["privilege_assignments"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EffectivePermissionsList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EffectivePermissionsList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(EffectivePrivilegeAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EffectivePermissionsList
// only implements ToObjectValue() and Type().
func (o EffectivePermissionsList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"privilege_assignments": o.PrivilegeAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EffectivePermissionsList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: EffectivePrivilegeAssignment{}.Type(ctx),
			},
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in EffectivePermissionsList as
// a slice of EffectivePrivilegeAssignment values.
// If the field is unknown or null, the boolean return value is false.
func (o *EffectivePermissionsList) GetPrivilegeAssignments(ctx context.Context) ([]EffectivePrivilegeAssignment, bool) {
	if o.PrivilegeAssignments.IsNull() || o.PrivilegeAssignments.IsUnknown() {
		return nil, false
	}
	var v []EffectivePrivilegeAssignment
	d := o.PrivilegeAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in EffectivePermissionsList.
func (o *EffectivePermissionsList) SetPrivilegeAssignments(ctx context.Context, v []EffectivePrivilegeAssignment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrivilegeAssignments = types.ListValueMust(t, vs)
}

type EffectivePredictiveOptimizationFlag struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromName types.String `tfsdk:"inherited_from_name"`
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromType types.String `tfsdk:"inherited_from_type"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	Value types.String `tfsdk:"value"`
}

func (newState *EffectivePredictiveOptimizationFlag) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePredictiveOptimizationFlag) {
}

func (newState *EffectivePredictiveOptimizationFlag) SyncEffectiveFieldsDuringRead(existingState EffectivePredictiveOptimizationFlag) {
}

func (c EffectivePredictiveOptimizationFlag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited_from_name"] = attrs["inherited_from_name"].SetOptional()
	attrs["inherited_from_type"] = attrs["inherited_from_type"].SetOptional()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EffectivePredictiveOptimizationFlag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EffectivePredictiveOptimizationFlag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EffectivePredictiveOptimizationFlag
// only implements ToObjectValue() and Type().
func (o EffectivePredictiveOptimizationFlag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited_from_name": o.InheritedFromName,
			"inherited_from_type": o.InheritedFromType,
			"value":               o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EffectivePredictiveOptimizationFlag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited_from_name": types.StringType,
			"inherited_from_type": types.StringType,
			"value":               types.StringType,
		},
	}
}

type EffectivePrivilege struct {
	// The full name of the object that conveys this privilege via inheritance.
	// This field is omitted when privilege is not inherited (it's assigned to
	// the securable itself).
	InheritedFromName types.String `tfsdk:"inherited_from_name"`
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	InheritedFromType types.String `tfsdk:"inherited_from_type"`
	// The privilege assigned to the principal.
	Privilege types.String `tfsdk:"privilege"`
}

func (newState *EffectivePrivilege) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePrivilege) {
}

func (newState *EffectivePrivilege) SyncEffectiveFieldsDuringRead(existingState EffectivePrivilege) {
}

func (c EffectivePrivilege) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited_from_name"] = attrs["inherited_from_name"].SetOptional()
	attrs["inherited_from_type"] = attrs["inherited_from_type"].SetOptional()
	attrs["privilege"] = attrs["privilege"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EffectivePrivilege.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EffectivePrivilege) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EffectivePrivilege
// only implements ToObjectValue() and Type().
func (o EffectivePrivilege) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited_from_name": o.InheritedFromName,
			"inherited_from_type": o.InheritedFromType,
			"privilege":           o.Privilege,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EffectivePrivilege) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited_from_name": types.StringType,
			"inherited_from_type": types.StringType,
			"privilege":           types.StringType,
		},
	}
}

type EffectivePrivilegeAssignment struct {
	// The principal (user email address or group name).
	Principal types.String `tfsdk:"principal"`
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	Privileges types.List `tfsdk:"privileges"`
}

func (newState *EffectivePrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePrivilegeAssignment) {
}

func (newState *EffectivePrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState EffectivePrivilegeAssignment) {
}

func (c EffectivePrivilegeAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal"] = attrs["principal"].SetOptional()
	attrs["privileges"] = attrs["privileges"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EffectivePrivilegeAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EffectivePrivilegeAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privileges": reflect.TypeOf(EffectivePrivilege{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EffectivePrivilegeAssignment
// only implements ToObjectValue() and Type().
func (o EffectivePrivilegeAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal":  o.Principal,
			"privileges": o.Privileges,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EffectivePrivilegeAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"privileges": basetypes.ListType{
				ElemType: EffectivePrivilege{}.Type(ctx),
			},
		},
	}
}

// GetPrivileges returns the value of the Privileges field in EffectivePrivilegeAssignment as
// a slice of EffectivePrivilege values.
// If the field is unknown or null, the boolean return value is false.
func (o *EffectivePrivilegeAssignment) GetPrivileges(ctx context.Context) ([]EffectivePrivilege, bool) {
	if o.Privileges.IsNull() || o.Privileges.IsUnknown() {
		return nil, false
	}
	var v []EffectivePrivilege
	d := o.Privileges.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivileges sets the value of the Privileges field in EffectivePrivilegeAssignment.
func (o *EffectivePrivilegeAssignment) SetPrivileges(ctx context.Context, v []EffectivePrivilege) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privileges"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Privileges = types.ListValueMust(t, vs)
}

type EnableRequest struct {
	// the catalog for which the system schema is to enabled in
	CatalogName types.String `tfsdk:"catalog_name"`
	// The metastore ID under which the system schema lives.
	MetastoreId types.String `tfsdk:"-"`
	// Full name of the system schema.
	SchemaName types.String `tfsdk:"-"`
}

func (newState *EnableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnableRequest) {
}

func (newState *EnableRequest) SyncEffectiveFieldsDuringRead(existingState EnableRequest) {
}

func (c EnableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableRequest
// only implements ToObjectValue() and Type().
func (o EnableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name": o.CatalogName,
			"metastore_id": o.MetastoreId,
			"schema_name":  o.SchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name": types.StringType,
			"metastore_id": types.StringType,
			"schema_name":  types.StringType,
		},
	}
}

type EnableResponse struct {
}

func (newState *EnableResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnableResponse) {
}

func (newState *EnableResponse) SyncEffectiveFieldsDuringRead(existingState EnableResponse) {
}

func (c EnableResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnableResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableResponse
// only implements ToObjectValue() and Type().
func (o EnableResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EnableResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Encryption options that apply to clients connecting to cloud storage.
type EncryptionDetails struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	SseEncryptionDetails types.Object `tfsdk:"sse_encryption_details"`
}

func (newState *EncryptionDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan EncryptionDetails) {
}

func (newState *EncryptionDetails) SyncEffectiveFieldsDuringRead(existingState EncryptionDetails) {
}

func (c EncryptionDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sse_encryption_details"] = attrs["sse_encryption_details"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EncryptionDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EncryptionDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sse_encryption_details": reflect.TypeOf(SseEncryptionDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EncryptionDetails
// only implements ToObjectValue() and Type().
func (o EncryptionDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sse_encryption_details": o.SseEncryptionDetails,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EncryptionDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sse_encryption_details": SseEncryptionDetails{}.Type(ctx),
		},
	}
}

// GetSseEncryptionDetails returns the value of the SseEncryptionDetails field in EncryptionDetails as
// a SseEncryptionDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *EncryptionDetails) GetSseEncryptionDetails(ctx context.Context) (SseEncryptionDetails, bool) {
	var e SseEncryptionDetails
	if o.SseEncryptionDetails.IsNull() || o.SseEncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []SseEncryptionDetails
	d := o.SseEncryptionDetails.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSseEncryptionDetails sets the value of the SseEncryptionDetails field in EncryptionDetails.
func (o *EncryptionDetails) SetSseEncryptionDetails(ctx context.Context, v SseEncryptionDetails) {
	vs := v.ToObjectValue(ctx)
	o.SseEncryptionDetails = vs
}

// Get boolean reflecting if table exists
type ExistsRequest struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExistsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExistsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExistsRequest
// only implements ToObjectValue() and Type().
func (o ExistsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExistsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
		},
	}
}

type ExternalLocationInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Time at which this external location was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of external location creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Unique ID of the location's storage credential.
	CredentialId types.String `tfsdk:"credential_id"`
	// Name of the storage credential used with this location.
	CredentialName types.String `tfsdk:"credential_name"`
	// [Create:OPT Update:OPT] Whether to enable file events on this external
	// location.
	EnableFileEvents types.Bool `tfsdk:"enable_file_events"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails types.Object `tfsdk:"encryption_details"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback types.Bool `tfsdk:"fallback"`
	// [Create:OPT Update:OPT] File event queue settings.
	FileEventQueue types.Object `tfsdk:"file_event_queue"`

	IsolationMode types.String `tfsdk:"isolation_mode"`
	// Unique identifier of metastore hosting the external location.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of the external location.
	Name types.String `tfsdk:"name"`
	// The owner of the external location.
	Owner types.String `tfsdk:"owner"`
	// Indicates whether the external location is read-only.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified the external location.
	UpdatedBy types.String `tfsdk:"updated_by"`
	// Path URL of the external location.
	Url types.String `tfsdk:"url"`
}

func (newState *ExternalLocationInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalLocationInfo) {
}

func (newState *ExternalLocationInfo) SyncEffectiveFieldsDuringRead(existingState ExternalLocationInfo) {
}

func (c ExternalLocationInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["browse_only"] = attrs["browse_only"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["credential_id"] = attrs["credential_id"].SetOptional()
	attrs["credential_name"] = attrs["credential_name"].SetOptional()
	attrs["enable_file_events"] = attrs["enable_file_events"].SetOptional()
	attrs["encryption_details"] = attrs["encryption_details"].SetOptional()
	attrs["fallback"] = attrs["fallback"].SetOptional()
	attrs["file_event_queue"] = attrs["file_event_queue"].SetOptional()
	attrs["isolation_mode"] = attrs["isolation_mode"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalLocationInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalLocationInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"encryption_details": reflect.TypeOf(EncryptionDetails{}),
		"file_event_queue":   reflect.TypeOf(FileEventQueue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalLocationInfo
// only implements ToObjectValue() and Type().
func (o ExternalLocationInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"browse_only":        o.BrowseOnly,
			"comment":            o.Comment,
			"created_at":         o.CreatedAt,
			"created_by":         o.CreatedBy,
			"credential_id":      o.CredentialId,
			"credential_name":    o.CredentialName,
			"enable_file_events": o.EnableFileEvents,
			"encryption_details": o.EncryptionDetails,
			"fallback":           o.Fallback,
			"file_event_queue":   o.FileEventQueue,
			"isolation_mode":     o.IsolationMode,
			"metastore_id":       o.MetastoreId,
			"name":               o.Name,
			"owner":              o.Owner,
			"read_only":          o.ReadOnly,
			"updated_at":         o.UpdatedAt,
			"updated_by":         o.UpdatedBy,
			"url":                o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalLocationInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":        types.BoolType,
			"comment":            types.StringType,
			"created_at":         types.Int64Type,
			"created_by":         types.StringType,
			"credential_id":      types.StringType,
			"credential_name":    types.StringType,
			"enable_file_events": types.BoolType,
			"encryption_details": EncryptionDetails{}.Type(ctx),
			"fallback":           types.BoolType,
			"file_event_queue":   FileEventQueue{}.Type(ctx),
			"isolation_mode":     types.StringType,
			"metastore_id":       types.StringType,
			"name":               types.StringType,
			"owner":              types.StringType,
			"read_only":          types.BoolType,
			"updated_at":         types.Int64Type,
			"updated_by":         types.StringType,
			"url":                types.StringType,
		},
	}
}

// GetEncryptionDetails returns the value of the EncryptionDetails field in ExternalLocationInfo as
// a EncryptionDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalLocationInfo) GetEncryptionDetails(ctx context.Context) (EncryptionDetails, bool) {
	var e EncryptionDetails
	if o.EncryptionDetails.IsNull() || o.EncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []EncryptionDetails
	d := o.EncryptionDetails.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEncryptionDetails sets the value of the EncryptionDetails field in ExternalLocationInfo.
func (o *ExternalLocationInfo) SetEncryptionDetails(ctx context.Context, v EncryptionDetails) {
	vs := v.ToObjectValue(ctx)
	o.EncryptionDetails = vs
}

// GetFileEventQueue returns the value of the FileEventQueue field in ExternalLocationInfo as
// a FileEventQueue value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalLocationInfo) GetFileEventQueue(ctx context.Context) (FileEventQueue, bool) {
	var e FileEventQueue
	if o.FileEventQueue.IsNull() || o.FileEventQueue.IsUnknown() {
		return e, false
	}
	var v []FileEventQueue
	d := o.FileEventQueue.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetFileEventQueue sets the value of the FileEventQueue field in ExternalLocationInfo.
func (o *ExternalLocationInfo) SetFileEventQueue(ctx context.Context, v FileEventQueue) {
	vs := v.ToObjectValue(ctx)
	o.FileEventQueue = vs
}

// Detailed status of an online table. Shown if the online table is in the
// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
type FailedStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may only be partially synced to the online
	// table. Only populated if the table is still online and available for
	// serving.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	Timestamp types.String `tfsdk:"timestamp"`
}

func (newState *FailedStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan FailedStatus) {
}

func (newState *FailedStatus) SyncEffectiveFieldsDuringRead(existingState FailedStatus) {
}

func (c FailedStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FailedStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FailedStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FailedStatus
// only implements ToObjectValue() and Type().
func (o FailedStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": o.LastProcessedCommitVersion,
			"timestamp":                     o.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FailedStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
		},
	}
}

type FileEventQueue struct {
	ManagedAqs types.Object `tfsdk:"managed_aqs"`

	ManagedPubsub types.Object `tfsdk:"managed_pubsub"`

	ManagedSqs types.Object `tfsdk:"managed_sqs"`

	ProvidedAqs types.Object `tfsdk:"provided_aqs"`

	ProvidedPubsub types.Object `tfsdk:"provided_pubsub"`

	ProvidedSqs types.Object `tfsdk:"provided_sqs"`
}

func (newState *FileEventQueue) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileEventQueue) {
}

func (newState *FileEventQueue) SyncEffectiveFieldsDuringRead(existingState FileEventQueue) {
}

func (c FileEventQueue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_aqs"] = attrs["managed_aqs"].SetOptional()
	attrs["managed_pubsub"] = attrs["managed_pubsub"].SetOptional()
	attrs["managed_sqs"] = attrs["managed_sqs"].SetOptional()
	attrs["provided_aqs"] = attrs["provided_aqs"].SetOptional()
	attrs["provided_pubsub"] = attrs["provided_pubsub"].SetOptional()
	attrs["provided_sqs"] = attrs["provided_sqs"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileEventQueue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FileEventQueue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_aqs":     reflect.TypeOf(AzureQueueStorage{}),
		"managed_pubsub":  reflect.TypeOf(GcpPubsub{}),
		"managed_sqs":     reflect.TypeOf(AwsSqsQueue{}),
		"provided_aqs":    reflect.TypeOf(AzureQueueStorage{}),
		"provided_pubsub": reflect.TypeOf(GcpPubsub{}),
		"provided_sqs":    reflect.TypeOf(AwsSqsQueue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileEventQueue
// only implements ToObjectValue() and Type().
func (o FileEventQueue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_aqs":     o.ManagedAqs,
			"managed_pubsub":  o.ManagedPubsub,
			"managed_sqs":     o.ManagedSqs,
			"provided_aqs":    o.ProvidedAqs,
			"provided_pubsub": o.ProvidedPubsub,
			"provided_sqs":    o.ProvidedSqs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FileEventQueue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_aqs":     AzureQueueStorage{}.Type(ctx),
			"managed_pubsub":  GcpPubsub{}.Type(ctx),
			"managed_sqs":     AwsSqsQueue{}.Type(ctx),
			"provided_aqs":    AzureQueueStorage{}.Type(ctx),
			"provided_pubsub": GcpPubsub{}.Type(ctx),
			"provided_sqs":    AwsSqsQueue{}.Type(ctx),
		},
	}
}

// GetManagedAqs returns the value of the ManagedAqs field in FileEventQueue as
// a AzureQueueStorage value.
// If the field is unknown or null, the boolean return value is false.
func (o *FileEventQueue) GetManagedAqs(ctx context.Context) (AzureQueueStorage, bool) {
	var e AzureQueueStorage
	if o.ManagedAqs.IsNull() || o.ManagedAqs.IsUnknown() {
		return e, false
	}
	var v []AzureQueueStorage
	d := o.ManagedAqs.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetManagedAqs sets the value of the ManagedAqs field in FileEventQueue.
func (o *FileEventQueue) SetManagedAqs(ctx context.Context, v AzureQueueStorage) {
	vs := v.ToObjectValue(ctx)
	o.ManagedAqs = vs
}

// GetManagedPubsub returns the value of the ManagedPubsub field in FileEventQueue as
// a GcpPubsub value.
// If the field is unknown or null, the boolean return value is false.
func (o *FileEventQueue) GetManagedPubsub(ctx context.Context) (GcpPubsub, bool) {
	var e GcpPubsub
	if o.ManagedPubsub.IsNull() || o.ManagedPubsub.IsUnknown() {
		return e, false
	}
	var v []GcpPubsub
	d := o.ManagedPubsub.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetManagedPubsub sets the value of the ManagedPubsub field in FileEventQueue.
func (o *FileEventQueue) SetManagedPubsub(ctx context.Context, v GcpPubsub) {
	vs := v.ToObjectValue(ctx)
	o.ManagedPubsub = vs
}

// GetManagedSqs returns the value of the ManagedSqs field in FileEventQueue as
// a AwsSqsQueue value.
// If the field is unknown or null, the boolean return value is false.
func (o *FileEventQueue) GetManagedSqs(ctx context.Context) (AwsSqsQueue, bool) {
	var e AwsSqsQueue
	if o.ManagedSqs.IsNull() || o.ManagedSqs.IsUnknown() {
		return e, false
	}
	var v []AwsSqsQueue
	d := o.ManagedSqs.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetManagedSqs sets the value of the ManagedSqs field in FileEventQueue.
func (o *FileEventQueue) SetManagedSqs(ctx context.Context, v AwsSqsQueue) {
	vs := v.ToObjectValue(ctx)
	o.ManagedSqs = vs
}

// GetProvidedAqs returns the value of the ProvidedAqs field in FileEventQueue as
// a AzureQueueStorage value.
// If the field is unknown or null, the boolean return value is false.
func (o *FileEventQueue) GetProvidedAqs(ctx context.Context) (AzureQueueStorage, bool) {
	var e AzureQueueStorage
	if o.ProvidedAqs.IsNull() || o.ProvidedAqs.IsUnknown() {
		return e, false
	}
	var v []AzureQueueStorage
	d := o.ProvidedAqs.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetProvidedAqs sets the value of the ProvidedAqs field in FileEventQueue.
func (o *FileEventQueue) SetProvidedAqs(ctx context.Context, v AzureQueueStorage) {
	vs := v.ToObjectValue(ctx)
	o.ProvidedAqs = vs
}

// GetProvidedPubsub returns the value of the ProvidedPubsub field in FileEventQueue as
// a GcpPubsub value.
// If the field is unknown or null, the boolean return value is false.
func (o *FileEventQueue) GetProvidedPubsub(ctx context.Context) (GcpPubsub, bool) {
	var e GcpPubsub
	if o.ProvidedPubsub.IsNull() || o.ProvidedPubsub.IsUnknown() {
		return e, false
	}
	var v []GcpPubsub
	d := o.ProvidedPubsub.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetProvidedPubsub sets the value of the ProvidedPubsub field in FileEventQueue.
func (o *FileEventQueue) SetProvidedPubsub(ctx context.Context, v GcpPubsub) {
	vs := v.ToObjectValue(ctx)
	o.ProvidedPubsub = vs
}

// GetProvidedSqs returns the value of the ProvidedSqs field in FileEventQueue as
// a AwsSqsQueue value.
// If the field is unknown or null, the boolean return value is false.
func (o *FileEventQueue) GetProvidedSqs(ctx context.Context) (AwsSqsQueue, bool) {
	var e AwsSqsQueue
	if o.ProvidedSqs.IsNull() || o.ProvidedSqs.IsUnknown() {
		return e, false
	}
	var v []AwsSqsQueue
	d := o.ProvidedSqs.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetProvidedSqs sets the value of the ProvidedSqs field in FileEventQueue.
func (o *FileEventQueue) SetProvidedSqs(ctx context.Context, v AwsSqsQueue) {
	vs := v.ToObjectValue(ctx)
	o.ProvidedSqs = vs
}

// Find a Database Instance by uid
type FindDatabaseInstanceByUidRequest struct {
	// UID of the cluster to get.
	Uid types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FindDatabaseInstanceByUidRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FindDatabaseInstanceByUidRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FindDatabaseInstanceByUidRequest
// only implements ToObjectValue() and Type().
func (o FindDatabaseInstanceByUidRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"uid": o.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FindDatabaseInstanceByUidRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"uid": types.StringType,
		},
	}
}

type ForeignKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns types.List `tfsdk:"child_columns"`
	// The name of the constraint.
	Name types.String `tfsdk:"name"`
	// Column names for this constraint.
	ParentColumns types.List `tfsdk:"parent_columns"`
	// The full name of the parent constraint.
	ParentTable types.String `tfsdk:"parent_table"`
}

func (newState *ForeignKeyConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForeignKeyConstraint) {
}

func (newState *ForeignKeyConstraint) SyncEffectiveFieldsDuringRead(existingState ForeignKeyConstraint) {
}

func (c ForeignKeyConstraint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["child_columns"] = attrs["child_columns"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["parent_columns"] = attrs["parent_columns"].SetRequired()
	attrs["parent_table"] = attrs["parent_table"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForeignKeyConstraint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ForeignKeyConstraint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"child_columns":  reflect.TypeOf(types.String{}),
		"parent_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForeignKeyConstraint
// only implements ToObjectValue() and Type().
func (o ForeignKeyConstraint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"child_columns":  o.ChildColumns,
			"name":           o.Name,
			"parent_columns": o.ParentColumns,
			"parent_table":   o.ParentTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ForeignKeyConstraint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"child_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"name": types.StringType,
			"parent_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"parent_table": types.StringType,
		},
	}
}

// GetChildColumns returns the value of the ChildColumns field in ForeignKeyConstraint as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ForeignKeyConstraint) GetChildColumns(ctx context.Context) ([]types.String, bool) {
	if o.ChildColumns.IsNull() || o.ChildColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ChildColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChildColumns sets the value of the ChildColumns field in ForeignKeyConstraint.
func (o *ForeignKeyConstraint) SetChildColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["child_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ChildColumns = types.ListValueMust(t, vs)
}

// GetParentColumns returns the value of the ParentColumns field in ForeignKeyConstraint as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ForeignKeyConstraint) GetParentColumns(ctx context.Context) ([]types.String, bool) {
	if o.ParentColumns.IsNull() || o.ParentColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ParentColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParentColumns sets the value of the ParentColumns field in ForeignKeyConstraint.
func (o *ForeignKeyConstraint) SetParentColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parent_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ParentColumns = types.ListValueMust(t, vs)
}

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	FunctionFullName types.String `tfsdk:"function_full_name"`
}

func (newState *FunctionDependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionDependency) {
}

func (newState *FunctionDependency) SyncEffectiveFieldsDuringRead(existingState FunctionDependency) {
}

func (c FunctionDependency) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["function_full_name"] = attrs["function_full_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FunctionDependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FunctionDependency) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionDependency
// only implements ToObjectValue() and Type().
func (o FunctionDependency) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_full_name": o.FunctionFullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FunctionDependency) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_full_name": types.StringType,
		},
	}
}

type FunctionInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Time at which this function was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of function creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Scalar function return data type.
	DataType types.String `tfsdk:"data_type"`
	// External function language.
	ExternalLanguage types.String `tfsdk:"external_language"`
	// External function name.
	ExternalName types.String `tfsdk:"external_name"`
	// Pretty printed function data type.
	FullDataType types.String `tfsdk:"full_data_type"`
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	FullName types.String `tfsdk:"full_name"`
	// Id of Function, relative to parent schema.
	FunctionId types.String `tfsdk:"function_id"`

	InputParams types.Object `tfsdk:"input_params"`
	// Whether the function is deterministic.
	IsDeterministic types.Bool `tfsdk:"is_deterministic"`
	// Function null call.
	IsNullCall types.Bool `tfsdk:"is_null_call"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of function, relative to parent schema.
	Name types.String `tfsdk:"name"`
	// Username of current owner of function.
	Owner types.String `tfsdk:"owner"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle types.String `tfsdk:"parameter_style"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties types.String `tfsdk:"properties"`
	// Table function return parameters.
	ReturnParams types.Object `tfsdk:"return_params"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody types.String `tfsdk:"routine_body"`
	// Function body.
	RoutineDefinition types.String `tfsdk:"routine_definition"`
	// Function dependencies.
	RoutineDependencies types.Object `tfsdk:"routine_dependencies"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name"`
	// Function security type.
	SecurityType types.String `tfsdk:"security_type"`
	// Specific name of the function; Reserved for future use.
	SpecificName types.String `tfsdk:"specific_name"`
	// Function SQL data access.
	SqlDataAccess types.String `tfsdk:"sql_data_access"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath types.String `tfsdk:"sql_path"`
	// Time at which this function was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified function.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *FunctionInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionInfo) {
}

func (newState *FunctionInfo) SyncEffectiveFieldsDuringRead(existingState FunctionInfo) {
}

func (c FunctionInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["browse_only"] = attrs["browse_only"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["data_type"] = attrs["data_type"].SetOptional()
	attrs["external_language"] = attrs["external_language"].SetOptional()
	attrs["external_name"] = attrs["external_name"].SetOptional()
	attrs["full_data_type"] = attrs["full_data_type"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["function_id"] = attrs["function_id"].SetOptional()
	attrs["input_params"] = attrs["input_params"].SetOptional()
	attrs["is_deterministic"] = attrs["is_deterministic"].SetOptional()
	attrs["is_null_call"] = attrs["is_null_call"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["parameter_style"] = attrs["parameter_style"].SetOptional()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["return_params"] = attrs["return_params"].SetOptional()
	attrs["routine_body"] = attrs["routine_body"].SetOptional()
	attrs["routine_definition"] = attrs["routine_definition"].SetOptional()
	attrs["routine_dependencies"] = attrs["routine_dependencies"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["security_type"] = attrs["security_type"].SetOptional()
	attrs["specific_name"] = attrs["specific_name"].SetOptional()
	attrs["sql_data_access"] = attrs["sql_data_access"].SetOptional()
	attrs["sql_path"] = attrs["sql_path"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FunctionInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FunctionInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input_params":         reflect.TypeOf(FunctionParameterInfos{}),
		"return_params":        reflect.TypeOf(FunctionParameterInfos{}),
		"routine_dependencies": reflect.TypeOf(DependencyList{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionInfo
// only implements ToObjectValue() and Type().
func (o FunctionInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"browse_only":          o.BrowseOnly,
			"catalog_name":         o.CatalogName,
			"comment":              o.Comment,
			"created_at":           o.CreatedAt,
			"created_by":           o.CreatedBy,
			"data_type":            o.DataType,
			"external_language":    o.ExternalLanguage,
			"external_name":        o.ExternalName,
			"full_data_type":       o.FullDataType,
			"full_name":            o.FullName,
			"function_id":          o.FunctionId,
			"input_params":         o.InputParams,
			"is_deterministic":     o.IsDeterministic,
			"is_null_call":         o.IsNullCall,
			"metastore_id":         o.MetastoreId,
			"name":                 o.Name,
			"owner":                o.Owner,
			"parameter_style":      o.ParameterStyle,
			"properties":           o.Properties,
			"return_params":        o.ReturnParams,
			"routine_body":         o.RoutineBody,
			"routine_definition":   o.RoutineDefinition,
			"routine_dependencies": o.RoutineDependencies,
			"schema_name":          o.SchemaName,
			"security_type":        o.SecurityType,
			"specific_name":        o.SpecificName,
			"sql_data_access":      o.SqlDataAccess,
			"sql_path":             o.SqlPath,
			"updated_at":           o.UpdatedAt,
			"updated_by":           o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FunctionInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":          types.BoolType,
			"catalog_name":         types.StringType,
			"comment":              types.StringType,
			"created_at":           types.Int64Type,
			"created_by":           types.StringType,
			"data_type":            types.StringType,
			"external_language":    types.StringType,
			"external_name":        types.StringType,
			"full_data_type":       types.StringType,
			"full_name":            types.StringType,
			"function_id":          types.StringType,
			"input_params":         FunctionParameterInfos{}.Type(ctx),
			"is_deterministic":     types.BoolType,
			"is_null_call":         types.BoolType,
			"metastore_id":         types.StringType,
			"name":                 types.StringType,
			"owner":                types.StringType,
			"parameter_style":      types.StringType,
			"properties":           types.StringType,
			"return_params":        FunctionParameterInfos{}.Type(ctx),
			"routine_body":         types.StringType,
			"routine_definition":   types.StringType,
			"routine_dependencies": DependencyList{}.Type(ctx),
			"schema_name":          types.StringType,
			"security_type":        types.StringType,
			"specific_name":        types.StringType,
			"sql_data_access":      types.StringType,
			"sql_path":             types.StringType,
			"updated_at":           types.Int64Type,
			"updated_by":           types.StringType,
		},
	}
}

// GetInputParams returns the value of the InputParams field in FunctionInfo as
// a FunctionParameterInfos value.
// If the field is unknown or null, the boolean return value is false.
func (o *FunctionInfo) GetInputParams(ctx context.Context) (FunctionParameterInfos, bool) {
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

// SetInputParams sets the value of the InputParams field in FunctionInfo.
func (o *FunctionInfo) SetInputParams(ctx context.Context, v FunctionParameterInfos) {
	vs := v.ToObjectValue(ctx)
	o.InputParams = vs
}

// GetReturnParams returns the value of the ReturnParams field in FunctionInfo as
// a FunctionParameterInfos value.
// If the field is unknown or null, the boolean return value is false.
func (o *FunctionInfo) GetReturnParams(ctx context.Context) (FunctionParameterInfos, bool) {
	var e FunctionParameterInfos
	if o.ReturnParams.IsNull() || o.ReturnParams.IsUnknown() {
		return e, false
	}
	var v []FunctionParameterInfos
	d := o.ReturnParams.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetReturnParams sets the value of the ReturnParams field in FunctionInfo.
func (o *FunctionInfo) SetReturnParams(ctx context.Context, v FunctionParameterInfos) {
	vs := v.ToObjectValue(ctx)
	o.ReturnParams = vs
}

// GetRoutineDependencies returns the value of the RoutineDependencies field in FunctionInfo as
// a DependencyList value.
// If the field is unknown or null, the boolean return value is false.
func (o *FunctionInfo) GetRoutineDependencies(ctx context.Context) (DependencyList, bool) {
	var e DependencyList
	if o.RoutineDependencies.IsNull() || o.RoutineDependencies.IsUnknown() {
		return e, false
	}
	var v []DependencyList
	d := o.RoutineDependencies.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetRoutineDependencies sets the value of the RoutineDependencies field in FunctionInfo.
func (o *FunctionInfo) SetRoutineDependencies(ctx context.Context, v DependencyList) {
	vs := v.ToObjectValue(ctx)
	o.RoutineDependencies = vs
}

type FunctionParameterInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Name of parameter.
	Name types.String `tfsdk:"name"`
	// Default value of the parameter.
	ParameterDefault types.String `tfsdk:"parameter_default"`
	// The mode of the function parameter.
	ParameterMode types.String `tfsdk:"parameter_mode"`
	// The type of function parameter.
	ParameterType types.String `tfsdk:"parameter_type"`
	// Ordinal position of column (starting at position 0).
	Position types.Int64 `tfsdk:"position"`
	// Format of IntervalType.
	TypeIntervalType types.String `tfsdk:"type_interval_type"`
	// Full data type spec, JSON-serialized.
	TypeJson types.String `tfsdk:"type_json"`

	TypeName types.String `tfsdk:"type_name"`
	// Digits of precision; required on Create for DecimalTypes.
	TypePrecision types.Int64 `tfsdk:"type_precision"`
	// Digits to right of decimal; Required on Create for DecimalTypes.
	TypeScale types.Int64 `tfsdk:"type_scale"`
	// Full data type spec, SQL/catalogString text.
	TypeText types.String `tfsdk:"type_text"`
}

func (newState *FunctionParameterInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionParameterInfo) {
}

func (newState *FunctionParameterInfo) SyncEffectiveFieldsDuringRead(existingState FunctionParameterInfo) {
}

func (c FunctionParameterInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["parameter_default"] = attrs["parameter_default"].SetOptional()
	attrs["parameter_mode"] = attrs["parameter_mode"].SetOptional()
	attrs["parameter_type"] = attrs["parameter_type"].SetOptional()
	attrs["position"] = attrs["position"].SetRequired()
	attrs["type_interval_type"] = attrs["type_interval_type"].SetOptional()
	attrs["type_json"] = attrs["type_json"].SetOptional()
	attrs["type_name"] = attrs["type_name"].SetRequired()
	attrs["type_precision"] = attrs["type_precision"].SetOptional()
	attrs["type_scale"] = attrs["type_scale"].SetOptional()
	attrs["type_text"] = attrs["type_text"].SetRequired()

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
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
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

// GCP temporary credentials for API authentication. Read more at
// https://developers.google.com/identity/protocols/oauth2/service-account
type GcpOauthToken struct {
	OauthToken types.String `tfsdk:"oauth_token"`
}

func (newState *GcpOauthToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpOauthToken) {
}

func (newState *GcpOauthToken) SyncEffectiveFieldsDuringRead(existingState GcpOauthToken) {
}

func (c GcpOauthToken) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["oauth_token"] = attrs["oauth_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpOauthToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpOauthToken) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpOauthToken
// only implements ToObjectValue() and Type().
func (o GcpOauthToken) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"oauth_token": o.OauthToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpOauthToken) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"oauth_token": types.StringType,
		},
	}
}

type GcpPubsub struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	ManagedResourceId types.String `tfsdk:"managed_resource_id"`
	// The Pub/Sub subscription name in the format
	// projects/{project}/subscriptions/{subscription name} REQUIRED for
	// provided_pubsub.
	SubscriptionName types.String `tfsdk:"subscription_name"`
}

func (newState *GcpPubsub) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpPubsub) {
}

func (newState *GcpPubsub) SyncEffectiveFieldsDuringRead(existingState GcpPubsub) {
}

func (c GcpPubsub) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_resource_id"] = attrs["managed_resource_id"].SetOptional()
	attrs["subscription_name"] = attrs["subscription_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpPubsub.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpPubsub) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpPubsub
// only implements ToObjectValue() and Type().
func (o GcpPubsub) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_resource_id": o.ManagedResourceId,
			"subscription_name":   o.SubscriptionName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpPubsub) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_resource_id": types.StringType,
			"subscription_name":   types.StringType,
		},
	}
}

// The Azure cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialAzureOptions struct {
	// The resources to which the temporary Azure credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://learn.microsoft.com/python/api/azure-core/azure.core.credentials.tokencredential?view=azure-python)
	Resources types.List `tfsdk:"resources"`
}

func (newState *GenerateTemporaryServiceCredentialAzureOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryServiceCredentialAzureOptions) {
}

func (newState *GenerateTemporaryServiceCredentialAzureOptions) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryServiceCredentialAzureOptions) {
}

func (c GenerateTemporaryServiceCredentialAzureOptions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["resources"] = attrs["resources"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateTemporaryServiceCredentialAzureOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenerateTemporaryServiceCredentialAzureOptions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resources": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateTemporaryServiceCredentialAzureOptions
// only implements ToObjectValue() and Type().
func (o GenerateTemporaryServiceCredentialAzureOptions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resources": o.Resources,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateTemporaryServiceCredentialAzureOptions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resources": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetResources returns the value of the Resources field in GenerateTemporaryServiceCredentialAzureOptions as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryServiceCredentialAzureOptions) GetResources(ctx context.Context) ([]types.String, bool) {
	if o.Resources.IsNull() || o.Resources.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in GenerateTemporaryServiceCredentialAzureOptions.
func (o *GenerateTemporaryServiceCredentialAzureOptions) SetResources(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

// The GCP cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialGcpOptions struct {
	// The scopes to which the temporary GCP credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://google-auth.readthedocs.io/en/latest/reference/google.auth.html#google.auth.credentials.Credentials)
	Scopes types.List `tfsdk:"scopes"`
}

func (newState *GenerateTemporaryServiceCredentialGcpOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryServiceCredentialGcpOptions) {
}

func (newState *GenerateTemporaryServiceCredentialGcpOptions) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryServiceCredentialGcpOptions) {
}

func (c GenerateTemporaryServiceCredentialGcpOptions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["scopes"] = attrs["scopes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateTemporaryServiceCredentialGcpOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenerateTemporaryServiceCredentialGcpOptions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateTemporaryServiceCredentialGcpOptions
// only implements ToObjectValue() and Type().
func (o GenerateTemporaryServiceCredentialGcpOptions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scopes": o.Scopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateTemporaryServiceCredentialGcpOptions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetScopes returns the value of the Scopes field in GenerateTemporaryServiceCredentialGcpOptions as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryServiceCredentialGcpOptions) GetScopes(ctx context.Context) ([]types.String, bool) {
	if o.Scopes.IsNull() || o.Scopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in GenerateTemporaryServiceCredentialGcpOptions.
func (o *GenerateTemporaryServiceCredentialGcpOptions) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

type GenerateTemporaryServiceCredentialRequest struct {
	// The Azure cloud options to customize the requested temporary credential
	AzureOptions types.Object `tfsdk:"azure_options"`
	// The name of the service credential used to generate a temporary
	// credential
	CredentialName types.String `tfsdk:"credential_name"`
	// The GCP cloud options to customize the requested temporary credential
	GcpOptions types.Object `tfsdk:"gcp_options"`
}

func (newState *GenerateTemporaryServiceCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryServiceCredentialRequest) {
}

func (newState *GenerateTemporaryServiceCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryServiceCredentialRequest) {
}

func (c GenerateTemporaryServiceCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["azure_options"] = attrs["azure_options"].SetOptional()
	attrs["credential_name"] = attrs["credential_name"].SetRequired()
	attrs["gcp_options"] = attrs["gcp_options"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateTemporaryServiceCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenerateTemporaryServiceCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"azure_options": reflect.TypeOf(GenerateTemporaryServiceCredentialAzureOptions{}),
		"gcp_options":   reflect.TypeOf(GenerateTemporaryServiceCredentialGcpOptions{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateTemporaryServiceCredentialRequest
// only implements ToObjectValue() and Type().
func (o GenerateTemporaryServiceCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"azure_options":   o.AzureOptions,
			"credential_name": o.CredentialName,
			"gcp_options":     o.GcpOptions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateTemporaryServiceCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"azure_options":   GenerateTemporaryServiceCredentialAzureOptions{}.Type(ctx),
			"credential_name": types.StringType,
			"gcp_options":     GenerateTemporaryServiceCredentialGcpOptions{}.Type(ctx),
		},
	}
}

// GetAzureOptions returns the value of the AzureOptions field in GenerateTemporaryServiceCredentialRequest as
// a GenerateTemporaryServiceCredentialAzureOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryServiceCredentialRequest) GetAzureOptions(ctx context.Context) (GenerateTemporaryServiceCredentialAzureOptions, bool) {
	var e GenerateTemporaryServiceCredentialAzureOptions
	if o.AzureOptions.IsNull() || o.AzureOptions.IsUnknown() {
		return e, false
	}
	var v []GenerateTemporaryServiceCredentialAzureOptions
	d := o.AzureOptions.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureOptions sets the value of the AzureOptions field in GenerateTemporaryServiceCredentialRequest.
func (o *GenerateTemporaryServiceCredentialRequest) SetAzureOptions(ctx context.Context, v GenerateTemporaryServiceCredentialAzureOptions) {
	vs := v.ToObjectValue(ctx)
	o.AzureOptions = vs
}

// GetGcpOptions returns the value of the GcpOptions field in GenerateTemporaryServiceCredentialRequest as
// a GenerateTemporaryServiceCredentialGcpOptions value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryServiceCredentialRequest) GetGcpOptions(ctx context.Context) (GenerateTemporaryServiceCredentialGcpOptions, bool) {
	var e GenerateTemporaryServiceCredentialGcpOptions
	if o.GcpOptions.IsNull() || o.GcpOptions.IsUnknown() {
		return e, false
	}
	var v []GenerateTemporaryServiceCredentialGcpOptions
	d := o.GcpOptions.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetGcpOptions sets the value of the GcpOptions field in GenerateTemporaryServiceCredentialRequest.
func (o *GenerateTemporaryServiceCredentialRequest) SetGcpOptions(ctx context.Context, v GenerateTemporaryServiceCredentialGcpOptions) {
	vs := v.ToObjectValue(ctx)
	o.GcpOptions = vs
}

type GenerateTemporaryTableCredentialRequest struct {
	// The operation performed against the table data, either READ or
	// READ_WRITE. If READ_WRITE is specified, the credentials returned will
	// have write permissions, otherwise, it will be read only.
	Operation types.String `tfsdk:"operation"`
	// UUID of the table to read or write.
	TableId types.String `tfsdk:"table_id"`
}

func (newState *GenerateTemporaryTableCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryTableCredentialRequest) {
}

func (newState *GenerateTemporaryTableCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryTableCredentialRequest) {
}

func (c GenerateTemporaryTableCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["operation"] = attrs["operation"].SetOptional()
	attrs["table_id"] = attrs["table_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateTemporaryTableCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenerateTemporaryTableCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateTemporaryTableCredentialRequest
// only implements ToObjectValue() and Type().
func (o GenerateTemporaryTableCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"operation": o.Operation,
			"table_id":  o.TableId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateTemporaryTableCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"operation": types.StringType,
			"table_id":  types.StringType,
		},
	}
}

type GenerateTemporaryTableCredentialResponse struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	AwsTempCredentials types.Object `tfsdk:"aws_temp_credentials"`
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	AzureAad types.Object `tfsdk:"azure_aad"`
	// Azure temporary credentials for API authentication. Read more at
	// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
	AzureUserDelegationSas types.Object `tfsdk:"azure_user_delegation_sas"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	GcpOauthToken types.Object `tfsdk:"gcp_oauth_token"`
	// R2 temporary credentials for API authentication. Read more at
	// https://developers.cloudflare.com/r2/api/s3/tokens/.
	R2TempCredentials types.Object `tfsdk:"r2_temp_credentials"`
	// The URL of the storage path accessible by the temporary credential.
	Url types.String `tfsdk:"url"`
}

func (newState *GenerateTemporaryTableCredentialResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryTableCredentialResponse) {
}

func (newState *GenerateTemporaryTableCredentialResponse) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryTableCredentialResponse) {
}

func (c GenerateTemporaryTableCredentialResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_temp_credentials"] = attrs["aws_temp_credentials"].SetOptional()
	attrs["azure_aad"] = attrs["azure_aad"].SetOptional()
	attrs["azure_user_delegation_sas"] = attrs["azure_user_delegation_sas"].SetOptional()
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["gcp_oauth_token"] = attrs["gcp_oauth_token"].SetOptional()
	attrs["r2_temp_credentials"] = attrs["r2_temp_credentials"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateTemporaryTableCredentialResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenerateTemporaryTableCredentialResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_temp_credentials":      reflect.TypeOf(AwsCredentials{}),
		"azure_aad":                 reflect.TypeOf(AzureActiveDirectoryToken{}),
		"azure_user_delegation_sas": reflect.TypeOf(AzureUserDelegationSas{}),
		"gcp_oauth_token":           reflect.TypeOf(GcpOauthToken{}),
		"r2_temp_credentials":       reflect.TypeOf(R2Credentials{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateTemporaryTableCredentialResponse
// only implements ToObjectValue() and Type().
func (o GenerateTemporaryTableCredentialResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_temp_credentials":      o.AwsTempCredentials,
			"azure_aad":                 o.AzureAad,
			"azure_user_delegation_sas": o.AzureUserDelegationSas,
			"expiration_time":           o.ExpirationTime,
			"gcp_oauth_token":           o.GcpOauthToken,
			"r2_temp_credentials":       o.R2TempCredentials,
			"url":                       o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateTemporaryTableCredentialResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_temp_credentials":      AwsCredentials{}.Type(ctx),
			"azure_aad":                 AzureActiveDirectoryToken{}.Type(ctx),
			"azure_user_delegation_sas": AzureUserDelegationSas{}.Type(ctx),
			"expiration_time":           types.Int64Type,
			"gcp_oauth_token":           GcpOauthToken{}.Type(ctx),
			"r2_temp_credentials":       R2Credentials{}.Type(ctx),
			"url":                       types.StringType,
		},
	}
}

// GetAwsTempCredentials returns the value of the AwsTempCredentials field in GenerateTemporaryTableCredentialResponse as
// a AwsCredentials value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryTableCredentialResponse) GetAwsTempCredentials(ctx context.Context) (AwsCredentials, bool) {
	var e AwsCredentials
	if o.AwsTempCredentials.IsNull() || o.AwsTempCredentials.IsUnknown() {
		return e, false
	}
	var v []AwsCredentials
	d := o.AwsTempCredentials.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsTempCredentials sets the value of the AwsTempCredentials field in GenerateTemporaryTableCredentialResponse.
func (o *GenerateTemporaryTableCredentialResponse) SetAwsTempCredentials(ctx context.Context, v AwsCredentials) {
	vs := v.ToObjectValue(ctx)
	o.AwsTempCredentials = vs
}

// GetAzureAad returns the value of the AzureAad field in GenerateTemporaryTableCredentialResponse as
// a AzureActiveDirectoryToken value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryTableCredentialResponse) GetAzureAad(ctx context.Context) (AzureActiveDirectoryToken, bool) {
	var e AzureActiveDirectoryToken
	if o.AzureAad.IsNull() || o.AzureAad.IsUnknown() {
		return e, false
	}
	var v []AzureActiveDirectoryToken
	d := o.AzureAad.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureAad sets the value of the AzureAad field in GenerateTemporaryTableCredentialResponse.
func (o *GenerateTemporaryTableCredentialResponse) SetAzureAad(ctx context.Context, v AzureActiveDirectoryToken) {
	vs := v.ToObjectValue(ctx)
	o.AzureAad = vs
}

// GetAzureUserDelegationSas returns the value of the AzureUserDelegationSas field in GenerateTemporaryTableCredentialResponse as
// a AzureUserDelegationSas value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryTableCredentialResponse) GetAzureUserDelegationSas(ctx context.Context) (AzureUserDelegationSas, bool) {
	var e AzureUserDelegationSas
	if o.AzureUserDelegationSas.IsNull() || o.AzureUserDelegationSas.IsUnknown() {
		return e, false
	}
	var v []AzureUserDelegationSas
	d := o.AzureUserDelegationSas.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureUserDelegationSas sets the value of the AzureUserDelegationSas field in GenerateTemporaryTableCredentialResponse.
func (o *GenerateTemporaryTableCredentialResponse) SetAzureUserDelegationSas(ctx context.Context, v AzureUserDelegationSas) {
	vs := v.ToObjectValue(ctx)
	o.AzureUserDelegationSas = vs
}

// GetGcpOauthToken returns the value of the GcpOauthToken field in GenerateTemporaryTableCredentialResponse as
// a GcpOauthToken value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryTableCredentialResponse) GetGcpOauthToken(ctx context.Context) (GcpOauthToken, bool) {
	var e GcpOauthToken
	if o.GcpOauthToken.IsNull() || o.GcpOauthToken.IsUnknown() {
		return e, false
	}
	var v []GcpOauthToken
	d := o.GcpOauthToken.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetGcpOauthToken sets the value of the GcpOauthToken field in GenerateTemporaryTableCredentialResponse.
func (o *GenerateTemporaryTableCredentialResponse) SetGcpOauthToken(ctx context.Context, v GcpOauthToken) {
	vs := v.ToObjectValue(ctx)
	o.GcpOauthToken = vs
}

// GetR2TempCredentials returns the value of the R2TempCredentials field in GenerateTemporaryTableCredentialResponse as
// a R2Credentials value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryTableCredentialResponse) GetR2TempCredentials(ctx context.Context) (R2Credentials, bool) {
	var e R2Credentials
	if o.R2TempCredentials.IsNull() || o.R2TempCredentials.IsUnknown() {
		return e, false
	}
	var v []R2Credentials
	d := o.R2TempCredentials.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetR2TempCredentials sets the value of the R2TempCredentials field in GenerateTemporaryTableCredentialResponse.
func (o *GenerateTemporaryTableCredentialResponse) SetR2TempCredentials(ctx context.Context, v R2Credentials) {
	vs := v.ToObjectValue(ctx)
	o.R2TempCredentials = vs
}

// Gets the metastore assignment for a workspace
type GetAccountMetastoreAssignmentRequest struct {
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountMetastoreAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountMetastoreAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountMetastoreAssignmentRequest
// only implements ToObjectValue() and Type().
func (o GetAccountMetastoreAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountMetastoreAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

// Get a metastore
type GetAccountMetastoreRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountMetastoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountMetastoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountMetastoreRequest
// only implements ToObjectValue() and Type().
func (o GetAccountMetastoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountMetastoreRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
		},
	}
}

// Gets the named storage credential
type GetAccountStorageCredentialRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Name of the storage credential.
	StorageCredentialName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountStorageCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountStorageCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountStorageCredentialRequest
// only implements ToObjectValue() and Type().
func (o GetAccountStorageCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id":            o.MetastoreId,
			"storage_credential_name": o.StorageCredentialName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountStorageCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id":            types.StringType,
			"storage_credential_name": types.StringType,
		},
	}
}

// Get an artifact allowlist
type GetArtifactAllowlistRequest struct {
	// The artifact type of the allowlist.
	ArtifactType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetArtifactAllowlistRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetArtifactAllowlistRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetArtifactAllowlistRequest
// only implements ToObjectValue() and Type().
func (o GetArtifactAllowlistRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_type": o.ArtifactType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetArtifactAllowlistRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_type": types.StringType,
		},
	}
}

// Get securable workspace bindings
type GetBindingsRequest struct {
	// Maximum number of workspace bindings to return. - When set to 0, the page
	// length is set to a server configured value (recommended); - When set to a
	// value greater than 0, the page length is the minimum of this value and a
	// server configured value; - When set to a value less than 0, an invalid
	// parameter error is returned; - If not set, all the workspace bindings are
	// returned (not recommended).
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
	// The name of the securable.
	SecurableName types.String `tfsdk:"-"`
	// The type of the securable to bind to a workspace (catalog,
	// storage_credential, credential, or external_location).
	SecurableType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBindingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBindingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBindingsRequest
// only implements ToObjectValue() and Type().
func (o GetBindingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results":    o.MaxResults,
			"page_token":     o.PageToken,
			"securable_name": o.SecurableName,
			"securable_type": o.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetBindingsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
			"securable_name": types.StringType,
			"securable_type": types.StringType,
		},
	}
}

// Get Model Version By Alias
type GetByAliasRequest struct {
	// The name of the alias
	Alias types.String `tfsdk:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases types.Bool `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetByAliasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetByAliasRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetByAliasRequest
// only implements ToObjectValue() and Type().
func (o GetByAliasRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":           o.Alias,
			"full_name":       o.FullName,
			"include_aliases": o.IncludeAliases,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetByAliasRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":           types.StringType,
			"full_name":       types.StringType,
			"include_aliases": types.BoolType,
		},
	}
}

// Get a catalog
type GetCatalogRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCatalogRequest
// only implements ToObjectValue() and Type().
func (o GetCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"name":           o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"name":           types.StringType,
		},
	}
}

type GetCatalogWorkspaceBindingsResponse struct {
	// A list of workspace IDs
	Workspaces types.List `tfsdk:"workspaces"`
}

func (newState *GetCatalogWorkspaceBindingsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCatalogWorkspaceBindingsResponse) {
}

func (newState *GetCatalogWorkspaceBindingsResponse) SyncEffectiveFieldsDuringRead(existingState GetCatalogWorkspaceBindingsResponse) {
}

func (c GetCatalogWorkspaceBindingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspaces"] = attrs["workspaces"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCatalogWorkspaceBindingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCatalogWorkspaceBindingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspaces": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCatalogWorkspaceBindingsResponse
// only implements ToObjectValue() and Type().
func (o GetCatalogWorkspaceBindingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspaces": o.Workspaces,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCatalogWorkspaceBindingsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspaces": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetWorkspaces returns the value of the Workspaces field in GetCatalogWorkspaceBindingsResponse as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetCatalogWorkspaceBindingsResponse) GetWorkspaces(ctx context.Context) ([]types.Int64, bool) {
	if o.Workspaces.IsNull() || o.Workspaces.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.Workspaces.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaces sets the value of the Workspaces field in GetCatalogWorkspaceBindingsResponse.
func (o *GetCatalogWorkspaceBindingsResponse) SetWorkspaces(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspaces"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Workspaces = types.ListValueMust(t, vs)
}

// Get a connection
type GetConnectionRequest struct {
	// Name of the connection.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetConnectionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetConnectionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetConnectionRequest
// only implements ToObjectValue() and Type().
func (o GetConnectionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetConnectionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a credential
type GetCredentialRequest struct {
	// Name of the credential.
	NameArg types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialRequest
// only implements ToObjectValue() and Type().
func (o GetCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name_arg": o.NameArg,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name_arg": types.StringType,
		},
	}
}

// Get a Database Catalog
type GetDatabaseCatalogRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (o GetDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a Database Instance
type GetDatabaseInstanceRequest struct {
	// Name of the cluster to get.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (o GetDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get effective permissions
type GetEffectiveRequest struct {
	// Full name of securable.
	FullName types.String `tfsdk:"-"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal types.String `tfsdk:"-"`
	// Type of securable.
	SecurableType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEffectiveRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEffectiveRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEffectiveRequest
// only implements ToObjectValue() and Type().
func (o GetEffectiveRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":      o.FullName,
			"principal":      o.Principal,
			"securable_type": o.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEffectiveRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":      types.StringType,
			"principal":      types.StringType,
			"securable_type": types.StringType,
		},
	}
}

// Get an external location
type GetExternalLocationRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Name of the external location.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExternalLocationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetExternalLocationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalLocationRequest
// only implements ToObjectValue() and Type().
func (o GetExternalLocationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"name":           o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExternalLocationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"name":           types.StringType,
		},
	}
}

// Get a function
type GetFunctionRequest struct {
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFunctionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetFunctionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFunctionRequest
// only implements ToObjectValue() and Type().
func (o GetFunctionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"name":           o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetFunctionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"name":           types.StringType,
		},
	}
}

// Get permissions
type GetGrantRequest struct {
	// Full name of securable.
	FullName types.String `tfsdk:"-"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal types.String `tfsdk:"-"`
	// Type of securable.
	SecurableType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGrantRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetGrantRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGrantRequest
// only implements ToObjectValue() and Type().
func (o GetGrantRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":      o.FullName,
			"principal":      o.Principal,
			"securable_type": o.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetGrantRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":      types.StringType,
			"principal":      types.StringType,
			"securable_type": types.StringType,
		},
	}
}

// Get a metastore
type GetMetastoreRequest struct {
	// Unique ID of the metastore.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMetastoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetMetastoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetastoreRequest
// only implements ToObjectValue() and Type().
func (o GetMetastoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetMetastoreRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetMetastoreSummaryResponse struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud types.String `tfsdk:"cloud"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of metastore creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId types.String `tfsdk:"default_data_access_config_id"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName types.String `tfsdk:"delta_sharing_organization_name"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds types.Int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope types.String `tfsdk:"delta_sharing_scope"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled types.Bool `tfsdk:"external_access_enabled"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id"`
	// Unique identifier of metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The user-specified name of the metastore.
	Name types.String `tfsdk:"name"`
	// The owner of the metastore.
	Owner types.String `tfsdk:"owner"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion types.String `tfsdk:"privilege_model_version"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region types.String `tfsdk:"region"`
	// The storage root URL for metastore
	StorageRoot types.String `tfsdk:"storage_root"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId types.String `tfsdk:"storage_root_credential_id"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName types.String `tfsdk:"storage_root_credential_name"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified the metastore.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *GetMetastoreSummaryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetastoreSummaryResponse) {
}

func (newState *GetMetastoreSummaryResponse) SyncEffectiveFieldsDuringRead(existingState GetMetastoreSummaryResponse) {
}

func (c GetMetastoreSummaryResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["default_data_access_config_id"] = attrs["default_data_access_config_id"].SetOptional()
	attrs["delta_sharing_organization_name"] = attrs["delta_sharing_organization_name"].SetOptional()
	attrs["delta_sharing_recipient_token_lifetime_in_seconds"] = attrs["delta_sharing_recipient_token_lifetime_in_seconds"].SetOptional()
	attrs["delta_sharing_scope"] = attrs["delta_sharing_scope"].SetOptional()
	attrs["external_access_enabled"] = attrs["external_access_enabled"].SetOptional()
	attrs["global_metastore_id"] = attrs["global_metastore_id"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["privilege_model_version"] = attrs["privilege_model_version"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()
	attrs["storage_root_credential_id"] = attrs["storage_root_credential_id"].SetOptional()
	attrs["storage_root_credential_name"] = attrs["storage_root_credential_name"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMetastoreSummaryResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetMetastoreSummaryResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetastoreSummaryResponse
// only implements ToObjectValue() and Type().
func (o GetMetastoreSummaryResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cloud":                           o.Cloud,
			"created_at":                      o.CreatedAt,
			"created_by":                      o.CreatedBy,
			"default_data_access_config_id":   o.DefaultDataAccessConfigId,
			"delta_sharing_organization_name": o.DeltaSharingOrganizationName,
			"delta_sharing_recipient_token_lifetime_in_seconds": o.DeltaSharingRecipientTokenLifetimeInSeconds,
			"delta_sharing_scope":                               o.DeltaSharingScope,
			"external_access_enabled":                           o.ExternalAccessEnabled,
			"global_metastore_id":                               o.GlobalMetastoreId,
			"metastore_id":                                      o.MetastoreId,
			"name":                                              o.Name,
			"owner":                                             o.Owner,
			"privilege_model_version":                           o.PrivilegeModelVersion,
			"region":                                            o.Region,
			"storage_root":                                      o.StorageRoot,
			"storage_root_credential_id":                        o.StorageRootCredentialId,
			"storage_root_credential_name":                      o.StorageRootCredentialName,
			"updated_at":                                        o.UpdatedAt,
			"updated_by":                                        o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetMetastoreSummaryResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cloud":                           types.StringType,
			"created_at":                      types.Int64Type,
			"created_by":                      types.StringType,
			"default_data_access_config_id":   types.StringType,
			"delta_sharing_organization_name": types.StringType,
			"delta_sharing_recipient_token_lifetime_in_seconds": types.Int64Type,
			"delta_sharing_scope":                               types.StringType,
			"external_access_enabled":                           types.BoolType,
			"global_metastore_id":                               types.StringType,
			"metastore_id":                                      types.StringType,
			"name":                                              types.StringType,
			"owner":                                             types.StringType,
			"privilege_model_version":                           types.StringType,
			"region":                                            types.StringType,
			"storage_root":                                      types.StringType,
			"storage_root_credential_id":                        types.StringType,
			"storage_root_credential_name":                      types.StringType,
			"updated_at":                                        types.Int64Type,
			"updated_by":                                        types.StringType,
		},
	}
}

// Get a Model Version
type GetModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	FullName types.String `tfsdk:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases types.Bool `tfsdk:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The integer version number of the model version
	Version types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionRequest
// only implements ToObjectValue() and Type().
func (o GetModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":       o.FullName,
			"include_aliases": o.IncludeAliases,
			"include_browse":  o.IncludeBrowse,
			"version":         o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":       types.StringType,
			"include_aliases": types.BoolType,
			"include_browse":  types.BoolType,
			"version":         types.Int64Type,
		},
	}
}

// Get an Online Table
type GetOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOnlineTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetOnlineTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOnlineTableRequest
// only implements ToObjectValue() and Type().
func (o GetOnlineTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetOnlineTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a table monitor
type GetQualityMonitorRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQualityMonitorRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQualityMonitorRequest
// only implements ToObjectValue() and Type().
func (o GetQualityMonitorRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQualityMonitorRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name": types.StringType,
		},
	}
}

// Get information for a single resource quota.
type GetQuotaRequest struct {
	// Full name of the parent resource. Provide the metastore ID if the parent
	// is a metastore.
	ParentFullName types.String `tfsdk:"-"`
	// Securable type of the quota parent.
	ParentSecurableType types.String `tfsdk:"-"`
	// Name of the quota. Follows the pattern of the quota type, with "-quota"
	// added as a suffix.
	QuotaName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQuotaRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQuotaRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQuotaRequest
// only implements ToObjectValue() and Type().
func (o GetQuotaRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent_full_name":      o.ParentFullName,
			"parent_securable_type": o.ParentSecurableType,
			"quota_name":            o.QuotaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQuotaRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent_full_name":      types.StringType,
			"parent_securable_type": types.StringType,
			"quota_name":            types.StringType,
		},
	}
}

type GetQuotaResponse struct {
	// The returned QuotaInfo.
	QuotaInfo types.Object `tfsdk:"quota_info"`
}

func (newState *GetQuotaResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQuotaResponse) {
}

func (newState *GetQuotaResponse) SyncEffectiveFieldsDuringRead(existingState GetQuotaResponse) {
}

func (c GetQuotaResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["quota_info"] = attrs["quota_info"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQuotaResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQuotaResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quota_info": reflect.TypeOf(QuotaInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQuotaResponse
// only implements ToObjectValue() and Type().
func (o GetQuotaResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quota_info": o.QuotaInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQuotaResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quota_info": QuotaInfo{}.Type(ctx),
		},
	}
}

// GetQuotaInfo returns the value of the QuotaInfo field in GetQuotaResponse as
// a QuotaInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetQuotaResponse) GetQuotaInfo(ctx context.Context) (QuotaInfo, bool) {
	var e QuotaInfo
	if o.QuotaInfo.IsNull() || o.QuotaInfo.IsUnknown() {
		return e, false
	}
	var v []QuotaInfo
	d := o.QuotaInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetQuotaInfo sets the value of the QuotaInfo field in GetQuotaResponse.
func (o *GetQuotaResponse) SetQuotaInfo(ctx context.Context, v QuotaInfo) {
	vs := v.ToObjectValue(ctx)
	o.QuotaInfo = vs
}

// Get refresh
type GetRefreshRequest struct {
	// ID of the refresh.
	RefreshId types.String `tfsdk:"-"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRefreshRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRefreshRequest
// only implements ToObjectValue() and Type().
func (o GetRefreshRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"refresh_id": o.RefreshId,
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRefreshRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"refresh_id": types.StringType,
			"table_name": types.StringType,
		},
	}
}

// Get a Registered Model
type GetRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
	// Whether to include registered model aliases in the response
	IncludeAliases types.Bool `tfsdk:"-"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRegisteredModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRegisteredModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelRequest
// only implements ToObjectValue() and Type().
func (o GetRegisteredModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":       o.FullName,
			"include_aliases": o.IncludeAliases,
			"include_browse":  o.IncludeBrowse,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRegisteredModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":       types.StringType,
			"include_aliases": types.BoolType,
			"include_browse":  types.BoolType,
		},
	}
}

// Get a schema
type GetSchemaRequest struct {
	// Full name of the schema.
	FullName types.String `tfsdk:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSchemaRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSchemaRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSchemaRequest
// only implements ToObjectValue() and Type().
func (o GetSchemaRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":      o.FullName,
			"include_browse": o.IncludeBrowse,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSchemaRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":      types.StringType,
			"include_browse": types.BoolType,
		},
	}
}

// Get a credential
type GetStorageCredentialRequest struct {
	// Name of the storage credential.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStorageCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStorageCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStorageCredentialRequest
// only implements ToObjectValue() and Type().
func (o GetStorageCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStorageCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a Synced Database Table
type GetSyncedDatabaseTableRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (o GetSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a table
type GetTableRequest struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata types.Bool `tfsdk:"-"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities types.Bool `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTableRequest
// only implements ToObjectValue() and Type().
func (o GetTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":                     o.FullName,
			"include_browse":                o.IncludeBrowse,
			"include_delta_metadata":        o.IncludeDeltaMetadata,
			"include_manifest_capabilities": o.IncludeManifestCapabilities,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":                     types.StringType,
			"include_browse":                types.BoolType,
			"include_delta_metadata":        types.BoolType,
			"include_manifest_capabilities": types.BoolType,
		},
	}
}

// Get catalog workspace bindings
type GetWorkspaceBindingRequest struct {
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceBindingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceBindingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceBindingRequest
// only implements ToObjectValue() and Type().
func (o GetWorkspaceBindingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceBindingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetWorkspaceBindingsResponse struct {
	// List of workspace bindings
	Bindings types.List `tfsdk:"bindings"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *GetWorkspaceBindingsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceBindingsResponse) {
}

func (newState *GetWorkspaceBindingsResponse) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceBindingsResponse) {
}

func (c GetWorkspaceBindingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bindings"] = attrs["bindings"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceBindingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceBindingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"bindings": reflect.TypeOf(WorkspaceBinding{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceBindingsResponse
// only implements ToObjectValue() and Type().
func (o GetWorkspaceBindingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bindings":        o.Bindings,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceBindingsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bindings": basetypes.ListType{
				ElemType: WorkspaceBinding{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetBindings returns the value of the Bindings field in GetWorkspaceBindingsResponse as
// a slice of WorkspaceBinding values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceBindingsResponse) GetBindings(ctx context.Context) ([]WorkspaceBinding, bool) {
	if o.Bindings.IsNull() || o.Bindings.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceBinding
	d := o.Bindings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBindings sets the value of the Bindings field in GetWorkspaceBindingsResponse.
func (o *GetWorkspaceBindingsResponse) SetBindings(ctx context.Context, v []WorkspaceBinding) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["bindings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Bindings = types.ListValueMust(t, vs)
}

// Get all workspaces assigned to a metastore
type ListAccountMetastoreAssignmentsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountMetastoreAssignmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountMetastoreAssignmentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountMetastoreAssignmentsRequest
// only implements ToObjectValue() and Type().
func (o ListAccountMetastoreAssignmentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountMetastoreAssignmentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
		},
	}
}

// The list of workspaces to which the given metastore is assigned.
type ListAccountMetastoreAssignmentsResponse struct {
	WorkspaceIds types.List `tfsdk:"workspace_ids"`
}

func (newState *ListAccountMetastoreAssignmentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountMetastoreAssignmentsResponse) {
}

func (newState *ListAccountMetastoreAssignmentsResponse) SyncEffectiveFieldsDuringRead(existingState ListAccountMetastoreAssignmentsResponse) {
}

func (c ListAccountMetastoreAssignmentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_ids"] = attrs["workspace_ids"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountMetastoreAssignmentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountMetastoreAssignmentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_ids": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountMetastoreAssignmentsResponse
// only implements ToObjectValue() and Type().
func (o ListAccountMetastoreAssignmentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_ids": o.WorkspaceIds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountMetastoreAssignmentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetWorkspaceIds returns the value of the WorkspaceIds field in ListAccountMetastoreAssignmentsResponse as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAccountMetastoreAssignmentsResponse) GetWorkspaceIds(ctx context.Context) ([]types.Int64, bool) {
	if o.WorkspaceIds.IsNull() || o.WorkspaceIds.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.WorkspaceIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceIds sets the value of the WorkspaceIds field in ListAccountMetastoreAssignmentsResponse.
func (o *ListAccountMetastoreAssignmentsResponse) SetWorkspaceIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WorkspaceIds = types.ListValueMust(t, vs)
}

// Get all storage credentials assigned to a metastore
type ListAccountStorageCredentialsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountStorageCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountStorageCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountStorageCredentialsRequest
// only implements ToObjectValue() and Type().
func (o ListAccountStorageCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountStorageCredentialsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
		},
	}
}

type ListAccountStorageCredentialsResponse struct {
	// An array of metastore storage credentials.
	StorageCredentials types.List `tfsdk:"storage_credentials"`
}

func (newState *ListAccountStorageCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountStorageCredentialsResponse) {
}

func (newState *ListAccountStorageCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState ListAccountStorageCredentialsResponse) {
}

func (c ListAccountStorageCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["storage_credentials"] = attrs["storage_credentials"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountStorageCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountStorageCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"storage_credentials": reflect.TypeOf(StorageCredentialInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountStorageCredentialsResponse
// only implements ToObjectValue() and Type().
func (o ListAccountStorageCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_credentials": o.StorageCredentials,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountStorageCredentialsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_credentials": basetypes.ListType{
				ElemType: StorageCredentialInfo{}.Type(ctx),
			},
		},
	}
}

// GetStorageCredentials returns the value of the StorageCredentials field in ListAccountStorageCredentialsResponse as
// a slice of StorageCredentialInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAccountStorageCredentialsResponse) GetStorageCredentials(ctx context.Context) ([]StorageCredentialInfo, bool) {
	if o.StorageCredentials.IsNull() || o.StorageCredentials.IsUnknown() {
		return nil, false
	}
	var v []StorageCredentialInfo
	d := o.StorageCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStorageCredentials sets the value of the StorageCredentials field in ListAccountStorageCredentialsResponse.
func (o *ListAccountStorageCredentialsResponse) SetStorageCredentials(ctx context.Context, v []StorageCredentialInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["storage_credentials"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.StorageCredentials = types.ListValueMust(t, vs)
}

// List catalogs
type ListCatalogsRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of catalogs to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid catalogs are returned (not
	// recommended). - Note: The number of returned catalogs might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further catalogs can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCatalogsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCatalogsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCatalogsRequest
// only implements ToObjectValue() and Type().
func (o ListCatalogsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"max_results":    o.MaxResults,
			"page_token":     o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCatalogsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
		},
	}
}

type ListCatalogsResponse struct {
	// An array of catalog information objects.
	Catalogs types.List `tfsdk:"catalogs"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListCatalogsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCatalogsResponse) {
}

func (newState *ListCatalogsResponse) SyncEffectiveFieldsDuringRead(existingState ListCatalogsResponse) {
}

func (c ListCatalogsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalogs"] = attrs["catalogs"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCatalogsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCatalogsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalogs": reflect.TypeOf(CatalogInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCatalogsResponse
// only implements ToObjectValue() and Type().
func (o ListCatalogsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalogs":        o.Catalogs,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCatalogsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalogs": basetypes.ListType{
				ElemType: CatalogInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetCatalogs returns the value of the Catalogs field in ListCatalogsResponse as
// a slice of CatalogInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCatalogsResponse) GetCatalogs(ctx context.Context) ([]CatalogInfo, bool) {
	if o.Catalogs.IsNull() || o.Catalogs.IsUnknown() {
		return nil, false
	}
	var v []CatalogInfo
	d := o.Catalogs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCatalogs sets the value of the Catalogs field in ListCatalogsResponse.
func (o *ListCatalogsResponse) SetCatalogs(ctx context.Context, v []CatalogInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["catalogs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Catalogs = types.ListValueMust(t, vs)
}

// List connections
type ListConnectionsRequest struct {
	// Maximum number of connections to return. - If not set, all connections
	// are returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListConnectionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListConnectionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListConnectionsRequest
// only implements ToObjectValue() and Type().
func (o ListConnectionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListConnectionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListConnectionsResponse struct {
	// An array of connection information objects.
	Connections types.List `tfsdk:"connections"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListConnectionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListConnectionsResponse) {
}

func (newState *ListConnectionsResponse) SyncEffectiveFieldsDuringRead(existingState ListConnectionsResponse) {
}

func (c ListConnectionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connections"] = attrs["connections"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListConnectionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListConnectionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"connections": reflect.TypeOf(ConnectionInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListConnectionsResponse
// only implements ToObjectValue() and Type().
func (o ListConnectionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connections":     o.Connections,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListConnectionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connections": basetypes.ListType{
				ElemType: ConnectionInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetConnections returns the value of the Connections field in ListConnectionsResponse as
// a slice of ConnectionInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListConnectionsResponse) GetConnections(ctx context.Context) ([]ConnectionInfo, bool) {
	if o.Connections.IsNull() || o.Connections.IsUnknown() {
		return nil, false
	}
	var v []ConnectionInfo
	d := o.Connections.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConnections sets the value of the Connections field in ListConnectionsResponse.
func (o *ListConnectionsResponse) SetConnections(ctx context.Context, v []ConnectionInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["connections"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Connections = types.ListValueMust(t, vs)
}

// List credentials
type ListCredentialsRequest struct {
	// Maximum number of credentials to return. - If not set, the default max
	// page size is used. - When set to a value greater than 0, the page length
	// is the minimum of this value and a server-configured value. - When set to
	// 0, the page length is set to a server-configured value (recommended). -
	// When set to a value less than 0, an invalid parameter error is returned.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque token to retrieve the next page of results.
	PageToken types.String `tfsdk:"-"`
	// Return only credentials for the specified purpose.
	Purpose types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsRequest
// only implements ToObjectValue() and Type().
func (o ListCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
			"purpose":     o.Purpose,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCredentialsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
			"purpose":     types.StringType,
		},
	}
}

type ListCredentialsResponse struct {
	Credentials types.List `tfsdk:"credentials"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCredentialsResponse) {
}

func (newState *ListCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState ListCredentialsResponse) {
}

func (c ListCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credentials"] = attrs["credentials"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credentials": reflect.TypeOf(CredentialInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsResponse
// only implements ToObjectValue() and Type().
func (o ListCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials":     o.Credentials,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCredentialsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials": basetypes.ListType{
				ElemType: CredentialInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetCredentials returns the value of the Credentials field in ListCredentialsResponse as
// a slice of CredentialInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCredentialsResponse) GetCredentials(ctx context.Context) ([]CredentialInfo, bool) {
	if o.Credentials.IsNull() || o.Credentials.IsUnknown() {
		return nil, false
	}
	var v []CredentialInfo
	d := o.Credentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCredentials sets the value of the Credentials field in ListCredentialsResponse.
func (o *ListCredentialsResponse) SetCredentials(ctx context.Context, v []CredentialInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["credentials"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Credentials = types.ListValueMust(t, vs)
}

// List Database Instances
type ListDatabaseInstancesRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseInstancesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDatabaseInstancesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstancesRequest
// only implements ToObjectValue() and Type().
func (o ListDatabaseInstancesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDatabaseInstancesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListDatabaseInstancesResponse struct {
	// List of instances.
	DatabaseInstances types.List `tfsdk:"database_instances"`
	// Pagination token to request the next page of instances.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListDatabaseInstancesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDatabaseInstancesResponse) {
}

func (newState *ListDatabaseInstancesResponse) SyncEffectiveFieldsDuringRead(existingState ListDatabaseInstancesResponse) {
}

func (c ListDatabaseInstancesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instances"] = attrs["database_instances"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseInstancesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDatabaseInstancesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instances": reflect.TypeOf(DatabaseInstance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstancesResponse
// only implements ToObjectValue() and Type().
func (o ListDatabaseInstancesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instances": o.DatabaseInstances,
			"next_page_token":    o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDatabaseInstancesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instances": basetypes.ListType{
				ElemType: DatabaseInstance{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatabaseInstances returns the value of the DatabaseInstances field in ListDatabaseInstancesResponse as
// a slice of DatabaseInstance values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListDatabaseInstancesResponse) GetDatabaseInstances(ctx context.Context) ([]DatabaseInstance, bool) {
	if o.DatabaseInstances.IsNull() || o.DatabaseInstances.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstance
	d := o.DatabaseInstances.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstances sets the value of the DatabaseInstances field in ListDatabaseInstancesResponse.
func (o *ListDatabaseInstancesResponse) SetDatabaseInstances(ctx context.Context, v []DatabaseInstance) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instances"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DatabaseInstances = types.ListValueMust(t, vs)
}

// List external locations
type ListExternalLocationsRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of external locations to return. If not set, all the
	// external locations are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExternalLocationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExternalLocationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExternalLocationsRequest
// only implements ToObjectValue() and Type().
func (o ListExternalLocationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"max_results":    o.MaxResults,
			"page_token":     o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExternalLocationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
		},
	}
}

type ListExternalLocationsResponse struct {
	// An array of external locations.
	ExternalLocations types.List `tfsdk:"external_locations"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListExternalLocationsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExternalLocationsResponse) {
}

func (newState *ListExternalLocationsResponse) SyncEffectiveFieldsDuringRead(existingState ListExternalLocationsResponse) {
}

func (c ListExternalLocationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_locations"] = attrs["external_locations"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExternalLocationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExternalLocationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"external_locations": reflect.TypeOf(ExternalLocationInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExternalLocationsResponse
// only implements ToObjectValue() and Type().
func (o ListExternalLocationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_locations": o.ExternalLocations,
			"next_page_token":    o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExternalLocationsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_locations": basetypes.ListType{
				ElemType: ExternalLocationInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExternalLocations returns the value of the ExternalLocations field in ListExternalLocationsResponse as
// a slice of ExternalLocationInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListExternalLocationsResponse) GetExternalLocations(ctx context.Context) ([]ExternalLocationInfo, bool) {
	if o.ExternalLocations.IsNull() || o.ExternalLocations.IsUnknown() {
		return nil, false
	}
	var v []ExternalLocationInfo
	d := o.ExternalLocations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalLocations sets the value of the ExternalLocations field in ListExternalLocationsResponse.
func (o *ListExternalLocationsResponse) SetExternalLocations(ctx context.Context, v []ExternalLocationInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["external_locations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExternalLocations = types.ListValueMust(t, vs)
}

// List functions
type ListFunctionsRequest struct {
	// Name of parent catalog for functions of interest.
	CatalogName types.String `tfsdk:"-"`
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of functions to return. If not set, all the functions are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
	// Parent schema of functions.
	SchemaName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFunctionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFunctionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFunctionsRequest
// only implements ToObjectValue() and Type().
func (o ListFunctionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":   o.CatalogName,
			"include_browse": o.IncludeBrowse,
			"max_results":    o.MaxResults,
			"page_token":     o.PageToken,
			"schema_name":    o.SchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFunctionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":   types.StringType,
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
			"schema_name":    types.StringType,
		},
	}
}

type ListFunctionsResponse struct {
	// An array of function information objects.
	Functions types.List `tfsdk:"functions"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListFunctionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFunctionsResponse) {
}

func (newState *ListFunctionsResponse) SyncEffectiveFieldsDuringRead(existingState ListFunctionsResponse) {
}

func (c ListFunctionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["functions"] = attrs["functions"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFunctionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFunctionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"functions": reflect.TypeOf(FunctionInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFunctionsResponse
// only implements ToObjectValue() and Type().
func (o ListFunctionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"functions":       o.Functions,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFunctionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"functions": basetypes.ListType{
				ElemType: FunctionInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFunctions returns the value of the Functions field in ListFunctionsResponse as
// a slice of FunctionInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListFunctionsResponse) GetFunctions(ctx context.Context) ([]FunctionInfo, bool) {
	if o.Functions.IsNull() || o.Functions.IsUnknown() {
		return nil, false
	}
	var v []FunctionInfo
	d := o.Functions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFunctions sets the value of the Functions field in ListFunctionsResponse.
func (o *ListFunctionsResponse) SetFunctions(ctx context.Context, v []FunctionInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["functions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Functions = types.ListValueMust(t, vs)
}

type ListMetastoresResponse struct {
	// An array of metastore information objects.
	Metastores types.List `tfsdk:"metastores"`
}

func (newState *ListMetastoresResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListMetastoresResponse) {
}

func (newState *ListMetastoresResponse) SyncEffectiveFieldsDuringRead(existingState ListMetastoresResponse) {
}

func (c ListMetastoresResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metastores"] = attrs["metastores"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListMetastoresResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListMetastoresResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastores": reflect.TypeOf(MetastoreInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListMetastoresResponse
// only implements ToObjectValue() and Type().
func (o ListMetastoresResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastores": o.Metastores,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListMetastoresResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastores": basetypes.ListType{
				ElemType: MetastoreInfo{}.Type(ctx),
			},
		},
	}
}

// GetMetastores returns the value of the Metastores field in ListMetastoresResponse as
// a slice of MetastoreInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListMetastoresResponse) GetMetastores(ctx context.Context) ([]MetastoreInfo, bool) {
	if o.Metastores.IsNull() || o.Metastores.IsUnknown() {
		return nil, false
	}
	var v []MetastoreInfo
	d := o.Metastores.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetastores sets the value of the Metastores field in ListMetastoresResponse.
func (o *ListMetastoresResponse) SetMetastores(ctx context.Context, v []MetastoreInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metastores"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Metastores = types.ListValueMust(t, vs)
}

// List Model Versions
type ListModelVersionsRequest struct {
	// The full three-level name of the registered model under which to list
	// model versions
	FullName types.String `tfsdk:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of model versions to return. If not set, the page length
	// is set to a server configured value (100, as of 1/3/2024). - when set to
	// a value greater than 0, the page length is the minimum of this value and
	// a server configured value(1000, as of 1/3/2024); - when set to 0, the
	// page length is set to a server configured value (100, as of 1/3/2024)
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListModelVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListModelVersionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelVersionsRequest
// only implements ToObjectValue() and Type().
func (o ListModelVersionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":      o.FullName,
			"include_browse": o.IncludeBrowse,
			"max_results":    o.MaxResults,
			"page_token":     o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListModelVersionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":      types.StringType,
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
		},
	}
}

type ListModelVersionsResponse struct {
	ModelVersions types.List `tfsdk:"model_versions"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListModelVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelVersionsResponse) {
}

func (newState *ListModelVersionsResponse) SyncEffectiveFieldsDuringRead(existingState ListModelVersionsResponse) {
}

func (c ListModelVersionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_versions"] = attrs["model_versions"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListModelVersionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListModelVersionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(ModelVersionInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelVersionsResponse
// only implements ToObjectValue() and Type().
func (o ListModelVersionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_versions":  o.ModelVersions,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListModelVersionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_versions": basetypes.ListType{
				ElemType: ModelVersionInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetModelVersions returns the value of the ModelVersions field in ListModelVersionsResponse as
// a slice of ModelVersionInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListModelVersionsResponse) GetModelVersions(ctx context.Context) ([]ModelVersionInfo, bool) {
	if o.ModelVersions.IsNull() || o.ModelVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionInfo
	d := o.ModelVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersions sets the value of the ModelVersions field in ListModelVersionsResponse.
func (o *ListModelVersionsResponse) SetModelVersions(ctx context.Context, v []ModelVersionInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ModelVersions = types.ListValueMust(t, vs)
}

// List all resource quotas under a metastore.
type ListQuotasRequest struct {
	// The number of quotas to return.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque token for the next page of results.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQuotasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQuotasRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQuotasRequest
// only implements ToObjectValue() and Type().
func (o ListQuotasRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQuotasRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListQuotasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of returned QuotaInfos.
	Quotas types.List `tfsdk:"quotas"`
}

func (newState *ListQuotasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQuotasResponse) {
}

func (newState *ListQuotasResponse) SyncEffectiveFieldsDuringRead(existingState ListQuotasResponse) {
}

func (c ListQuotasResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["quotas"] = attrs["quotas"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQuotasResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQuotasResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quotas": reflect.TypeOf(QuotaInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQuotasResponse
// only implements ToObjectValue() and Type().
func (o ListQuotasResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"quotas":          o.Quotas,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQuotasResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"quotas": basetypes.ListType{
				ElemType: QuotaInfo{}.Type(ctx),
			},
		},
	}
}

// GetQuotas returns the value of the Quotas field in ListQuotasResponse as
// a slice of QuotaInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQuotasResponse) GetQuotas(ctx context.Context) ([]QuotaInfo, bool) {
	if o.Quotas.IsNull() || o.Quotas.IsUnknown() {
		return nil, false
	}
	var v []QuotaInfo
	d := o.Quotas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQuotas sets the value of the Quotas field in ListQuotasResponse.
func (o *ListQuotasResponse) SetQuotas(ctx context.Context, v []QuotaInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["quotas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Quotas = types.ListValueMust(t, vs)
}

// List refreshes
type ListRefreshesRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRefreshesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRefreshesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRefreshesRequest
// only implements ToObjectValue() and Type().
func (o ListRefreshesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRefreshesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name": types.StringType,
		},
	}
}

// List Registered Models
type ListRegisteredModelsRequest struct {
	// The identifier of the catalog under which to list registered models. If
	// specified, schema_name must be specified.
	CatalogName types.String `tfsdk:"-"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Max number of registered models to return.
	//
	// If both catalog and schema are specified: - when max_results is not
	// specified, the page length is set to a server configured value (10000, as
	// of 4/2/2024). - when set to a value greater than 0, the page length is
	// the minimum of this value and a server configured value (10000, as of
	// 4/2/2024); - when set to 0, the page length is set to a server configured
	// value (10000, as of 4/2/2024); - when set to a value less than 0, an
	// invalid parameter error is returned;
	//
	// If neither schema nor catalog is specified: - when max_results is not
	// specified, the page length is set to a server configured value (100, as
	// of 4/2/2024). - when set to a value greater than 0, the page length is
	// the minimum of this value and a server configured value (1000, as of
	// 4/2/2024); - when set to 0, the page length is set to a server configured
	// value (100, as of 4/2/2024); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque token to send for the next page of results (pagination).
	PageToken types.String `tfsdk:"-"`
	// The identifier of the schema under which to list registered models. If
	// specified, catalog_name must be specified.
	SchemaName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRegisteredModelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRegisteredModelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRegisteredModelsRequest
// only implements ToObjectValue() and Type().
func (o ListRegisteredModelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":   o.CatalogName,
			"include_browse": o.IncludeBrowse,
			"max_results":    o.MaxResults,
			"page_token":     o.PageToken,
			"schema_name":    o.SchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRegisteredModelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":   types.StringType,
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
			"schema_name":    types.StringType,
		},
	}
}

type ListRegisteredModelsResponse struct {
	// Opaque token for pagination. Omitted if there are no more results.
	// page_token should be set to this value for fetching the next page.
	NextPageToken types.String `tfsdk:"next_page_token"`

	RegisteredModels types.List `tfsdk:"registered_models"`
}

func (newState *ListRegisteredModelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRegisteredModelsResponse) {
}

func (newState *ListRegisteredModelsResponse) SyncEffectiveFieldsDuringRead(existingState ListRegisteredModelsResponse) {
}

func (c ListRegisteredModelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["registered_models"] = attrs["registered_models"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRegisteredModelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRegisteredModelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_models": reflect.TypeOf(RegisteredModelInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRegisteredModelsResponse
// only implements ToObjectValue() and Type().
func (o ListRegisteredModelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"registered_models": o.RegisteredModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRegisteredModelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"registered_models": basetypes.ListType{
				ElemType: RegisteredModelInfo{}.Type(ctx),
			},
		},
	}
}

// GetRegisteredModels returns the value of the RegisteredModels field in ListRegisteredModelsResponse as
// a slice of RegisteredModelInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListRegisteredModelsResponse) GetRegisteredModels(ctx context.Context) ([]RegisteredModelInfo, bool) {
	if o.RegisteredModels.IsNull() || o.RegisteredModels.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelInfo
	d := o.RegisteredModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModels sets the value of the RegisteredModels field in ListRegisteredModelsResponse.
func (o *ListRegisteredModelsResponse) SetRegisteredModels(ctx context.Context, v []RegisteredModelInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RegisteredModels = types.ListValueMust(t, vs)
}

// List schemas
type ListSchemasRequest struct {
	// Parent catalog for schemas of interest.
	CatalogName types.String `tfsdk:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of schemas to return. If not set, all the schemas are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSchemasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSchemasRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSchemasRequest
// only implements ToObjectValue() and Type().
func (o ListSchemasRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":   o.CatalogName,
			"include_browse": o.IncludeBrowse,
			"max_results":    o.MaxResults,
			"page_token":     o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSchemasRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":   types.StringType,
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
		},
	}
}

type ListSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of schema information objects.
	Schemas types.List `tfsdk:"schemas"`
}

func (newState *ListSchemasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchemasResponse) {
}

func (newState *ListSchemasResponse) SyncEffectiveFieldsDuringRead(existingState ListSchemasResponse) {
}

func (c ListSchemasResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["schemas"] = attrs["schemas"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSchemasResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSchemasResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schemas": reflect.TypeOf(SchemaInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSchemasResponse
// only implements ToObjectValue() and Type().
func (o ListSchemasResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"schemas":         o.Schemas,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSchemasResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"schemas": basetypes.ListType{
				ElemType: SchemaInfo{}.Type(ctx),
			},
		},
	}
}

// GetSchemas returns the value of the Schemas field in ListSchemasResponse as
// a slice of SchemaInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSchemasResponse) GetSchemas(ctx context.Context) ([]SchemaInfo, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []SchemaInfo
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in ListSchemasResponse.
func (o *ListSchemasResponse) SetSchemas(ctx context.Context, v []SchemaInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

// List credentials
type ListStorageCredentialsRequest struct {
	// Maximum number of storage credentials to return. If not set, all the
	// storage credentials are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStorageCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListStorageCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStorageCredentialsRequest
// only implements ToObjectValue() and Type().
func (o ListStorageCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListStorageCredentialsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListStorageCredentialsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`

	StorageCredentials types.List `tfsdk:"storage_credentials"`
}

func (newState *ListStorageCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListStorageCredentialsResponse) {
}

func (newState *ListStorageCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState ListStorageCredentialsResponse) {
}

func (c ListStorageCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["storage_credentials"] = attrs["storage_credentials"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStorageCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListStorageCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"storage_credentials": reflect.TypeOf(StorageCredentialInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStorageCredentialsResponse
// only implements ToObjectValue() and Type().
func (o ListStorageCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":     o.NextPageToken,
			"storage_credentials": o.StorageCredentials,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListStorageCredentialsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"storage_credentials": basetypes.ListType{
				ElemType: StorageCredentialInfo{}.Type(ctx),
			},
		},
	}
}

// GetStorageCredentials returns the value of the StorageCredentials field in ListStorageCredentialsResponse as
// a slice of StorageCredentialInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListStorageCredentialsResponse) GetStorageCredentials(ctx context.Context) ([]StorageCredentialInfo, bool) {
	if o.StorageCredentials.IsNull() || o.StorageCredentials.IsUnknown() {
		return nil, false
	}
	var v []StorageCredentialInfo
	d := o.StorageCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStorageCredentials sets the value of the StorageCredentials field in ListStorageCredentialsResponse.
func (o *ListStorageCredentialsResponse) SetStorageCredentials(ctx context.Context, v []StorageCredentialInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["storage_credentials"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.StorageCredentials = types.ListValueMust(t, vs)
}

// List table summaries
type ListSummariesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName types.String `tfsdk:"-"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities types.Bool `tfsdk:"-"`
	// Maximum number of summaries for tables to return. If not set, the page
	// length is set to a server configured value (10000, as of 1/5/2024). -
	// when set to a value greater than 0, the page length is the minimum of
	// this value and a server configured value (10000, as of 1/5/2024); - when
	// set to 0, the page length is set to a server configured value (10000, as
	// of 1/5/2024) (recommended); - when set to a value less than 0, an invalid
	// parameter error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
	// A sql LIKE pattern (% and _) for schema names. All schemas will be
	// returned if not set or empty.
	SchemaNamePattern types.String `tfsdk:"-"`
	// A sql LIKE pattern (% and _) for table names. All tables will be returned
	// if not set or empty.
	TableNamePattern types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSummariesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSummariesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSummariesRequest
// only implements ToObjectValue() and Type().
func (o ListSummariesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":                  o.CatalogName,
			"include_manifest_capabilities": o.IncludeManifestCapabilities,
			"max_results":                   o.MaxResults,
			"page_token":                    o.PageToken,
			"schema_name_pattern":           o.SchemaNamePattern,
			"table_name_pattern":            o.TableNamePattern,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSummariesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":                  types.StringType,
			"include_manifest_capabilities": types.BoolType,
			"max_results":                   types.Int64Type,
			"page_token":                    types.StringType,
			"schema_name_pattern":           types.StringType,
			"table_name_pattern":            types.StringType,
		},
	}
}

// List system schemas
type ListSystemSchemasRequest struct {
	// Maximum number of schemas to return. - When set to 0, the page length is
	// set to a server configured value (recommended); - When set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - When set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all the schemas are returned (not
	// recommended).
	MaxResults types.Int64 `tfsdk:"-"`
	// The ID for the metastore in which the system schema resides.
	MetastoreId types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSystemSchemasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSystemSchemasRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSystemSchemasRequest
// only implements ToObjectValue() and Type().
func (o ListSystemSchemasRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results":  o.MaxResults,
			"metastore_id": o.MetastoreId,
			"page_token":   o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSystemSchemasRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results":  types.Int64Type,
			"metastore_id": types.StringType,
			"page_token":   types.StringType,
		},
	}
}

type ListSystemSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of system schema information objects.
	Schemas types.List `tfsdk:"schemas"`
}

func (newState *ListSystemSchemasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSystemSchemasResponse) {
}

func (newState *ListSystemSchemasResponse) SyncEffectiveFieldsDuringRead(existingState ListSystemSchemasResponse) {
}

func (c ListSystemSchemasResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["schemas"] = attrs["schemas"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSystemSchemasResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSystemSchemasResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schemas": reflect.TypeOf(SystemSchemaInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSystemSchemasResponse
// only implements ToObjectValue() and Type().
func (o ListSystemSchemasResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"schemas":         o.Schemas,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSystemSchemasResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"schemas": basetypes.ListType{
				ElemType: SystemSchemaInfo{}.Type(ctx),
			},
		},
	}
}

// GetSchemas returns the value of the Schemas field in ListSystemSchemasResponse as
// a slice of SystemSchemaInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSystemSchemasResponse) GetSchemas(ctx context.Context) ([]SystemSchemaInfo, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []SystemSchemaInfo
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in ListSystemSchemasResponse.
func (o *ListSystemSchemasResponse) SetSchemas(ctx context.Context, v []SystemSchemaInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type ListTableSummariesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of table summaries.
	Tables types.List `tfsdk:"tables"`
}

func (newState *ListTableSummariesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTableSummariesResponse) {
}

func (newState *ListTableSummariesResponse) SyncEffectiveFieldsDuringRead(existingState ListTableSummariesResponse) {
}

func (c ListTableSummariesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["tables"] = attrs["tables"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTableSummariesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTableSummariesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tables": reflect.TypeOf(TableSummary{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTableSummariesResponse
// only implements ToObjectValue() and Type().
func (o ListTableSummariesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"tables":          o.Tables,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTableSummariesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tables": basetypes.ListType{
				ElemType: TableSummary{}.Type(ctx),
			},
		},
	}
}

// GetTables returns the value of the Tables field in ListTableSummariesResponse as
// a slice of TableSummary values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListTableSummariesResponse) GetTables(ctx context.Context) ([]TableSummary, bool) {
	if o.Tables.IsNull() || o.Tables.IsUnknown() {
		return nil, false
	}
	var v []TableSummary
	d := o.Tables.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTables sets the value of the Tables field in ListTableSummariesResponse.
func (o *ListTableSummariesResponse) SetTables(ctx context.Context, v []TableSummary) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tables"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tables = types.ListValueMust(t, vs)
}

// List tables
type ListTablesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName types.String `tfsdk:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata types.Bool `tfsdk:"-"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities types.Bool `tfsdk:"-"`
	// Maximum number of tables to return. If not set, all the tables are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults types.Int64 `tfsdk:"-"`
	// Whether to omit the columns of the table from the response or not.
	OmitColumns types.Bool `tfsdk:"-"`
	// Whether to omit the properties of the table from the response or not.
	OmitProperties types.Bool `tfsdk:"-"`
	// Whether to omit the username of the table (e.g. owner, updated_by,
	// created_by) from the response or not.
	OmitUsername types.Bool `tfsdk:"-"`
	// Opaque token to send for the next page of results (pagination).
	PageToken types.String `tfsdk:"-"`
	// Parent schema of tables.
	SchemaName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTablesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTablesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTablesRequest
// only implements ToObjectValue() and Type().
func (o ListTablesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":                  o.CatalogName,
			"include_browse":                o.IncludeBrowse,
			"include_delta_metadata":        o.IncludeDeltaMetadata,
			"include_manifest_capabilities": o.IncludeManifestCapabilities,
			"max_results":                   o.MaxResults,
			"omit_columns":                  o.OmitColumns,
			"omit_properties":               o.OmitProperties,
			"omit_username":                 o.OmitUsername,
			"page_token":                    o.PageToken,
			"schema_name":                   o.SchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTablesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":                  types.StringType,
			"include_browse":                types.BoolType,
			"include_delta_metadata":        types.BoolType,
			"include_manifest_capabilities": types.BoolType,
			"max_results":                   types.Int64Type,
			"omit_columns":                  types.BoolType,
			"omit_properties":               types.BoolType,
			"omit_username":                 types.BoolType,
			"page_token":                    types.StringType,
			"schema_name":                   types.StringType,
		},
	}
}

type ListTablesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// An array of table information objects.
	Tables types.List `tfsdk:"tables"`
}

func (newState *ListTablesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTablesResponse) {
}

func (newState *ListTablesResponse) SyncEffectiveFieldsDuringRead(existingState ListTablesResponse) {
}

func (c ListTablesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["tables"] = attrs["tables"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTablesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTablesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tables": reflect.TypeOf(TableInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTablesResponse
// only implements ToObjectValue() and Type().
func (o ListTablesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"tables":          o.Tables,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTablesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tables": basetypes.ListType{
				ElemType: TableInfo{}.Type(ctx),
			},
		},
	}
}

// GetTables returns the value of the Tables field in ListTablesResponse as
// a slice of TableInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListTablesResponse) GetTables(ctx context.Context) ([]TableInfo, bool) {
	if o.Tables.IsNull() || o.Tables.IsUnknown() {
		return nil, false
	}
	var v []TableInfo
	d := o.Tables.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTables sets the value of the Tables field in ListTablesResponse.
func (o *ListTablesResponse) SetTables(ctx context.Context, v []TableInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tables"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tables = types.ListValueMust(t, vs)
}

// List Volumes
type ListVolumesRequest struct {
	// The identifier of the catalog
	CatalogName types.String `tfsdk:"-"`
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Maximum number of volumes to return (page length).
	//
	// If not set, the page length is set to a server configured value (10000,
	// as of 1/29/2024). - when set to a value greater than 0, the page length
	// is the minimum of this value and a server configured value (10000, as of
	// 1/29/2024); - when set to 0, the page length is set to a server
	// configured value (10000, as of 1/29/2024) (recommended); - when set to a
	// value less than 0, an invalid parameter error is returned;
	//
	// Note: this parameter controls only the maximum number of volumes to
	// return. The actual number of volumes returned in a page may be smaller
	// than this value, including 0, even if there are more pages.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque token returned by a previous request. It must be included in the
	// request to retrieve the next page of results (pagination).
	PageToken types.String `tfsdk:"-"`
	// The identifier of the schema
	SchemaName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVolumesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVolumesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVolumesRequest
// only implements ToObjectValue() and Type().
func (o ListVolumesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":   o.CatalogName,
			"include_browse": o.IncludeBrowse,
			"max_results":    o.MaxResults,
			"page_token":     o.PageToken,
			"schema_name":    o.SchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListVolumesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":   types.StringType,
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
			"schema_name":    types.StringType,
		},
	}
}

type ListVolumesResponseContent struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request to retrieve the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Volumes types.List `tfsdk:"volumes"`
}

func (newState *ListVolumesResponseContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVolumesResponseContent) {
}

func (newState *ListVolumesResponseContent) SyncEffectiveFieldsDuringRead(existingState ListVolumesResponseContent) {
}

func (c ListVolumesResponseContent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["volumes"] = attrs["volumes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVolumesResponseContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVolumesResponseContent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"volumes": reflect.TypeOf(VolumeInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVolumesResponseContent
// only implements ToObjectValue() and Type().
func (o ListVolumesResponseContent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"volumes":         o.Volumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListVolumesResponseContent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"volumes": basetypes.ListType{
				ElemType: VolumeInfo{}.Type(ctx),
			},
		},
	}
}

// GetVolumes returns the value of the Volumes field in ListVolumesResponseContent as
// a slice of VolumeInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListVolumesResponseContent) GetVolumes(ctx context.Context) ([]VolumeInfo, bool) {
	if o.Volumes.IsNull() || o.Volumes.IsUnknown() {
		return nil, false
	}
	var v []VolumeInfo
	d := o.Volumes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolumes sets the value of the Volumes field in ListVolumesResponseContent.
func (o *ListVolumesResponseContent) SetVolumes(ctx context.Context, v []VolumeInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["volumes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Volumes = types.ListValueMust(t, vs)
}

type MetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name"`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The unique ID of the Databricks workspace.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (newState *MetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan MetastoreAssignment) {
}

func (newState *MetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState MetastoreAssignment) {
}

func (c MetastoreAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_catalog_name"] = attrs["default_catalog_name"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MetastoreAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MetastoreAssignment
// only implements ToObjectValue() and Type().
func (o MetastoreAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_catalog_name": o.DefaultCatalogName,
			"metastore_id":         o.MetastoreId,
			"workspace_id":         o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MetastoreAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_catalog_name": types.StringType,
			"metastore_id":         types.StringType,
			"workspace_id":         types.Int64Type,
		},
	}
}

type MetastoreInfo struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud types.String `tfsdk:"cloud"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of metastore creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId types.String `tfsdk:"default_data_access_config_id"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName types.String `tfsdk:"delta_sharing_organization_name"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds types.Int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope types.String `tfsdk:"delta_sharing_scope"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled types.Bool `tfsdk:"external_access_enabled"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id"`
	// Unique identifier of metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The user-specified name of the metastore.
	Name types.String `tfsdk:"name"`
	// The owner of the metastore.
	Owner types.String `tfsdk:"owner"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion types.String `tfsdk:"privilege_model_version"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region types.String `tfsdk:"region"`
	// The storage root URL for metastore
	StorageRoot types.String `tfsdk:"storage_root"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId types.String `tfsdk:"storage_root_credential_id"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName types.String `tfsdk:"storage_root_credential_name"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified the metastore.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *MetastoreInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan MetastoreInfo) {
}

func (newState *MetastoreInfo) SyncEffectiveFieldsDuringRead(existingState MetastoreInfo) {
}

func (c MetastoreInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["default_data_access_config_id"] = attrs["default_data_access_config_id"].SetOptional()
	attrs["delta_sharing_organization_name"] = attrs["delta_sharing_organization_name"].SetOptional()
	attrs["delta_sharing_recipient_token_lifetime_in_seconds"] = attrs["delta_sharing_recipient_token_lifetime_in_seconds"].SetOptional()
	attrs["delta_sharing_scope"] = attrs["delta_sharing_scope"].SetOptional()
	attrs["external_access_enabled"] = attrs["external_access_enabled"].SetOptional()
	attrs["global_metastore_id"] = attrs["global_metastore_id"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["privilege_model_version"] = attrs["privilege_model_version"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()
	attrs["storage_root_credential_id"] = attrs["storage_root_credential_id"].SetOptional()
	attrs["storage_root_credential_name"] = attrs["storage_root_credential_name"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MetastoreInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MetastoreInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MetastoreInfo
// only implements ToObjectValue() and Type().
func (o MetastoreInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cloud":                           o.Cloud,
			"created_at":                      o.CreatedAt,
			"created_by":                      o.CreatedBy,
			"default_data_access_config_id":   o.DefaultDataAccessConfigId,
			"delta_sharing_organization_name": o.DeltaSharingOrganizationName,
			"delta_sharing_recipient_token_lifetime_in_seconds": o.DeltaSharingRecipientTokenLifetimeInSeconds,
			"delta_sharing_scope":                               o.DeltaSharingScope,
			"external_access_enabled":                           o.ExternalAccessEnabled,
			"global_metastore_id":                               o.GlobalMetastoreId,
			"metastore_id":                                      o.MetastoreId,
			"name":                                              o.Name,
			"owner":                                             o.Owner,
			"privilege_model_version":                           o.PrivilegeModelVersion,
			"region":                                            o.Region,
			"storage_root":                                      o.StorageRoot,
			"storage_root_credential_id":                        o.StorageRootCredentialId,
			"storage_root_credential_name":                      o.StorageRootCredentialName,
			"updated_at":                                        o.UpdatedAt,
			"updated_by":                                        o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MetastoreInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cloud":                           types.StringType,
			"created_at":                      types.Int64Type,
			"created_by":                      types.StringType,
			"default_data_access_config_id":   types.StringType,
			"delta_sharing_organization_name": types.StringType,
			"delta_sharing_recipient_token_lifetime_in_seconds": types.Int64Type,
			"delta_sharing_scope":                               types.StringType,
			"external_access_enabled":                           types.BoolType,
			"global_metastore_id":                               types.StringType,
			"metastore_id":                                      types.StringType,
			"name":                                              types.StringType,
			"owner":                                             types.StringType,
			"privilege_model_version":                           types.StringType,
			"region":                                            types.StringType,
			"storage_root":                                      types.StringType,
			"storage_root_credential_id":                        types.StringType,
			"storage_root_credential_name":                      types.StringType,
			"updated_at":                                        types.Int64Type,
			"updated_by":                                        types.StringType,
		},
	}
}

type ModelVersionInfo struct {
	// List of aliases associated with the model version
	Aliases types.List `tfsdk:"aliases"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// The name of the catalog containing the model version
	CatalogName types.String `tfsdk:"catalog_name"`
	// The comment attached to the model version
	Comment types.String `tfsdk:"comment"`

	CreatedAt types.Int64 `tfsdk:"created_at"`
	// The identifier of the user who created the model version
	CreatedBy types.String `tfsdk:"created_by"`
	// The unique identifier of the model version
	Id types.String `tfsdk:"id"`
	// The unique identifier of the metastore containing the model version
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The name of the parent registered model of the model version, relative to
	// parent schema
	ModelName types.String `tfsdk:"model_name"`
	// Model version dependencies, for feature-store packaged models
	ModelVersionDependencies types.Object `tfsdk:"model_version_dependencies"`
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	RunId types.String `tfsdk:"run_id"`
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	RunWorkspaceId types.Int64 `tfsdk:"run_workspace_id"`
	// The name of the schema containing the model version, relative to parent
	// catalog
	SchemaName types.String `tfsdk:"schema_name"`
	// URI indicating the location of the source artifacts (files) for the model
	// version
	Source types.String `tfsdk:"source"`
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	Status types.String `tfsdk:"status"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation types.String `tfsdk:"storage_location"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// The identifier of the user who updated the model version last time
	UpdatedBy types.String `tfsdk:"updated_by"`
	// Integer model version number, used to reference the model version in API
	// requests.
	Version types.Int64 `tfsdk:"version"`
}

func (newState *ModelVersionInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersionInfo) {
}

func (newState *ModelVersionInfo) SyncEffectiveFieldsDuringRead(existingState ModelVersionInfo) {
}

func (c ModelVersionInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aliases"] = attrs["aliases"].SetOptional()
	attrs["browse_only"] = attrs["browse_only"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["model_name"] = attrs["model_name"].SetOptional()
	attrs["model_version_dependencies"] = attrs["model_version_dependencies"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_workspace_id"] = attrs["run_workspace_id"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["source"] = attrs["source"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelVersionInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ModelVersionInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aliases":                    reflect.TypeOf(RegisteredModelAlias{}),
		"model_version_dependencies": reflect.TypeOf(DependencyList{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersionInfo
// only implements ToObjectValue() and Type().
func (o ModelVersionInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aliases":                    o.Aliases,
			"browse_only":                o.BrowseOnly,
			"catalog_name":               o.CatalogName,
			"comment":                    o.Comment,
			"created_at":                 o.CreatedAt,
			"created_by":                 o.CreatedBy,
			"id":                         o.Id,
			"metastore_id":               o.MetastoreId,
			"model_name":                 o.ModelName,
			"model_version_dependencies": o.ModelVersionDependencies,
			"run_id":                     o.RunId,
			"run_workspace_id":           o.RunWorkspaceId,
			"schema_name":                o.SchemaName,
			"source":                     o.Source,
			"status":                     o.Status,
			"storage_location":           o.StorageLocation,
			"updated_at":                 o.UpdatedAt,
			"updated_by":                 o.UpdatedBy,
			"version":                    o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelVersionInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aliases": basetypes.ListType{
				ElemType: RegisteredModelAlias{}.Type(ctx),
			},
			"browse_only":                types.BoolType,
			"catalog_name":               types.StringType,
			"comment":                    types.StringType,
			"created_at":                 types.Int64Type,
			"created_by":                 types.StringType,
			"id":                         types.StringType,
			"metastore_id":               types.StringType,
			"model_name":                 types.StringType,
			"model_version_dependencies": DependencyList{}.Type(ctx),
			"run_id":                     types.StringType,
			"run_workspace_id":           types.Int64Type,
			"schema_name":                types.StringType,
			"source":                     types.StringType,
			"status":                     types.StringType,
			"storage_location":           types.StringType,
			"updated_at":                 types.Int64Type,
			"updated_by":                 types.StringType,
			"version":                    types.Int64Type,
		},
	}
}

// GetAliases returns the value of the Aliases field in ModelVersionInfo as
// a slice of RegisteredModelAlias values.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelVersionInfo) GetAliases(ctx context.Context) ([]RegisteredModelAlias, bool) {
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

// SetAliases sets the value of the Aliases field in ModelVersionInfo.
func (o *ModelVersionInfo) SetAliases(ctx context.Context, v []RegisteredModelAlias) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Aliases = types.ListValueMust(t, vs)
}

// GetModelVersionDependencies returns the value of the ModelVersionDependencies field in ModelVersionInfo as
// a DependencyList value.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelVersionInfo) GetModelVersionDependencies(ctx context.Context) (DependencyList, bool) {
	var e DependencyList
	if o.ModelVersionDependencies.IsNull() || o.ModelVersionDependencies.IsUnknown() {
		return e, false
	}
	var v []DependencyList
	d := o.ModelVersionDependencies.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetModelVersionDependencies sets the value of the ModelVersionDependencies field in ModelVersionInfo.
func (o *ModelVersionInfo) SetModelVersionDependencies(ctx context.Context, v DependencyList) {
	vs := v.ToObjectValue(ctx)
	o.ModelVersionDependencies = vs
}

type MonitorCronSchedule struct {
	// Read only field that indicates whether a schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression types.String `tfsdk:"quartz_cron_expression"`
	// The timezone id (e.g., ``"PST"``) in which to evaluate the quartz
	// expression.
	TimezoneId types.String `tfsdk:"timezone_id"`
}

func (newState *MonitorCronSchedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorCronSchedule) {
}

func (newState *MonitorCronSchedule) SyncEffectiveFieldsDuringRead(existingState MonitorCronSchedule) {
}

func (c MonitorCronSchedule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pause_status"] = attrs["pause_status"].SetOptional()
	attrs["quartz_cron_expression"] = attrs["quartz_cron_expression"].SetRequired()
	attrs["timezone_id"] = attrs["timezone_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorCronSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorCronSchedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorCronSchedule
// only implements ToObjectValue() and Type().
func (o MonitorCronSchedule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pause_status":           o.PauseStatus,
			"quartz_cron_expression": o.QuartzCronExpression,
			"timezone_id":            o.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorCronSchedule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pause_status":           types.StringType,
			"quartz_cron_expression": types.StringType,
			"timezone_id":            types.StringType,
		},
	}
}

type MonitorDataClassificationConfig struct {
	// Whether data classification is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
}

func (newState *MonitorDataClassificationConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorDataClassificationConfig) {
}

func (newState *MonitorDataClassificationConfig) SyncEffectiveFieldsDuringRead(existingState MonitorDataClassificationConfig) {
}

func (c MonitorDataClassificationConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorDataClassificationConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorDataClassificationConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorDataClassificationConfig
// only implements ToObjectValue() and Type().
func (o MonitorDataClassificationConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled": o.Enabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorDataClassificationConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

type MonitorDestination struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	EmailAddresses types.List `tfsdk:"email_addresses"`
}

func (newState *MonitorDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorDestination) {
}

func (newState *MonitorDestination) SyncEffectiveFieldsDuringRead(existingState MonitorDestination) {
}

func (c MonitorDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["email_addresses"] = attrs["email_addresses"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorDestination
// only implements ToObjectValue() and Type().
func (o MonitorDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email_addresses": o.EmailAddresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorDestination) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetEmailAddresses returns the value of the EmailAddresses field in MonitorDestination as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorDestination) GetEmailAddresses(ctx context.Context) ([]types.String, bool) {
	if o.EmailAddresses.IsNull() || o.EmailAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.EmailAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmailAddresses sets the value of the EmailAddresses field in MonitorDestination.
func (o *MonitorDestination) SetEmailAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmailAddresses = types.ListValueMust(t, vs)
}

type MonitorInferenceLog struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities types.List `tfsdk:"granularities"`
	// Optional column that contains the ground truth for the prediction.
	LabelCol types.String `tfsdk:"label_col"`
	// Column that contains the id of the model generating the predictions.
	// Metrics will be computed per model id by default, and also across all
	// model ids.
	ModelIdCol types.String `tfsdk:"model_id_col"`
	// Column that contains the output/prediction from the model.
	PredictionCol types.String `tfsdk:"prediction_col"`
	// Optional column that contains the prediction probabilities for each class
	// in a classification problem type. The values in this column should be a
	// map, mapping each class label to the prediction probability for a given
	// sample. The map should be of PySpark MapType().
	PredictionProbaCol types.String `tfsdk:"prediction_proba_col"`
	// Problem type the model aims to solve. Determines the type of
	// model-quality metrics that will be computed.
	ProblemType types.String `tfsdk:"problem_type"`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol types.String `tfsdk:"timestamp_col"`
}

func (newState *MonitorInferenceLog) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorInferenceLog) {
}

func (newState *MonitorInferenceLog) SyncEffectiveFieldsDuringRead(existingState MonitorInferenceLog) {
}

func (c MonitorInferenceLog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["granularities"] = attrs["granularities"].SetRequired()
	attrs["label_col"] = attrs["label_col"].SetOptional()
	attrs["model_id_col"] = attrs["model_id_col"].SetRequired()
	attrs["prediction_col"] = attrs["prediction_col"].SetRequired()
	attrs["prediction_proba_col"] = attrs["prediction_proba_col"].SetOptional()
	attrs["problem_type"] = attrs["problem_type"].SetRequired()
	attrs["timestamp_col"] = attrs["timestamp_col"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorInferenceLog.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorInferenceLog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"granularities": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorInferenceLog
// only implements ToObjectValue() and Type().
func (o MonitorInferenceLog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"granularities":        o.Granularities,
			"label_col":            o.LabelCol,
			"model_id_col":         o.ModelIdCol,
			"prediction_col":       o.PredictionCol,
			"prediction_proba_col": o.PredictionProbaCol,
			"problem_type":         o.ProblemType,
			"timestamp_col":        o.TimestampCol,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorInferenceLog) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"granularities": basetypes.ListType{
				ElemType: types.StringType,
			},
			"label_col":            types.StringType,
			"model_id_col":         types.StringType,
			"prediction_col":       types.StringType,
			"prediction_proba_col": types.StringType,
			"problem_type":         types.StringType,
			"timestamp_col":        types.StringType,
		},
	}
}

// GetGranularities returns the value of the Granularities field in MonitorInferenceLog as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInferenceLog) GetGranularities(ctx context.Context) ([]types.String, bool) {
	if o.Granularities.IsNull() || o.Granularities.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Granularities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGranularities sets the value of the Granularities field in MonitorInferenceLog.
func (o *MonitorInferenceLog) SetGranularities(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["granularities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Granularities = types.ListValueMust(t, vs)
}

type MonitorInfo struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir types.String `tfsdk:"assets_dir"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName types.String `tfsdk:"baseline_table_name"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics types.List `tfsdk:"custom_metrics"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The data classification config for the monitor.
	DataClassificationConfig types.Object `tfsdk:"data_classification_config"`
	// The full name of the drift metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	DriftMetricsTableName types.String `tfsdk:"drift_metrics_table_name"`
	// Configuration for monitoring inference logs.
	InferenceLog types.Object `tfsdk:"inference_log"`
	// The latest failure message of the monitor (if any).
	LatestMonitorFailureMsg types.String `tfsdk:"latest_monitor_failure_msg"`
	// The version of the monitor config (e.g. 1,2,3). If negative, the monitor
	// may be corrupted.
	MonitorVersion types.String `tfsdk:"monitor_version"`
	// The notification settings for the monitor.
	Notifications types.Object `tfsdk:"notifications"`
	// Schema where output metric tables are created.
	OutputSchemaName types.String `tfsdk:"output_schema_name"`
	// The full name of the profile metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	ProfileMetricsTableName types.String `tfsdk:"profile_metrics_table_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule types.Object `tfsdk:"schedule"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs types.List `tfsdk:"slicing_exprs"`
	// Configuration for monitoring snapshot tables.
	Snapshot types.Object `tfsdk:"snapshot"`
	// The status of the monitor.
	Status types.String `tfsdk:"status"`
	// The full name of the table to monitor. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	TableName types.String `tfsdk:"table_name"`
	// Configuration for monitoring time series tables.
	TimeSeries types.Object `tfsdk:"time_series"`
}

func (newState *MonitorInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorInfo) {
}

func (newState *MonitorInfo) SyncEffectiveFieldsDuringRead(existingState MonitorInfo) {
}

func (c MonitorInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assets_dir"] = attrs["assets_dir"].SetOptional()
	attrs["baseline_table_name"] = attrs["baseline_table_name"].SetOptional()
	attrs["custom_metrics"] = attrs["custom_metrics"].SetOptional()
	attrs["dashboard_id"] = attrs["dashboard_id"].SetOptional()
	attrs["data_classification_config"] = attrs["data_classification_config"].SetOptional()
	attrs["drift_metrics_table_name"] = attrs["drift_metrics_table_name"].SetRequired()
	attrs["inference_log"] = attrs["inference_log"].SetOptional()
	attrs["latest_monitor_failure_msg"] = attrs["latest_monitor_failure_msg"].SetOptional()
	attrs["monitor_version"] = attrs["monitor_version"].SetRequired()
	attrs["notifications"] = attrs["notifications"].SetOptional()
	attrs["output_schema_name"] = attrs["output_schema_name"].SetOptional()
	attrs["profile_metrics_table_name"] = attrs["profile_metrics_table_name"].SetRequired()
	attrs["schedule"] = attrs["schedule"].SetOptional()
	attrs["slicing_exprs"] = attrs["slicing_exprs"].SetOptional()
	attrs["snapshot"] = attrs["snapshot"].SetOptional()
	attrs["status"] = attrs["status"].SetRequired()
	attrs["table_name"] = attrs["table_name"].SetRequired()
	attrs["time_series"] = attrs["time_series"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_metrics":             reflect.TypeOf(MonitorMetric{}),
		"data_classification_config": reflect.TypeOf(MonitorDataClassificationConfig{}),
		"inference_log":              reflect.TypeOf(MonitorInferenceLog{}),
		"notifications":              reflect.TypeOf(MonitorNotifications{}),
		"schedule":                   reflect.TypeOf(MonitorCronSchedule{}),
		"slicing_exprs":              reflect.TypeOf(types.String{}),
		"snapshot":                   reflect.TypeOf(MonitorSnapshot{}),
		"time_series":                reflect.TypeOf(MonitorTimeSeries{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorInfo
// only implements ToObjectValue() and Type().
func (o MonitorInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets_dir":                 o.AssetsDir,
			"baseline_table_name":        o.BaselineTableName,
			"custom_metrics":             o.CustomMetrics,
			"dashboard_id":               o.DashboardId,
			"data_classification_config": o.DataClassificationConfig,
			"drift_metrics_table_name":   o.DriftMetricsTableName,
			"inference_log":              o.InferenceLog,
			"latest_monitor_failure_msg": o.LatestMonitorFailureMsg,
			"monitor_version":            o.MonitorVersion,
			"notifications":              o.Notifications,
			"output_schema_name":         o.OutputSchemaName,
			"profile_metrics_table_name": o.ProfileMetricsTableName,
			"schedule":                   o.Schedule,
			"slicing_exprs":              o.SlicingExprs,
			"snapshot":                   o.Snapshot,
			"status":                     o.Status,
			"table_name":                 o.TableName,
			"time_series":                o.TimeSeries,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets_dir":          types.StringType,
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: MonitorMetric{}.Type(ctx),
			},
			"dashboard_id":               types.StringType,
			"data_classification_config": MonitorDataClassificationConfig{}.Type(ctx),
			"drift_metrics_table_name":   types.StringType,
			"inference_log":              MonitorInferenceLog{}.Type(ctx),
			"latest_monitor_failure_msg": types.StringType,
			"monitor_version":            types.StringType,
			"notifications":              MonitorNotifications{}.Type(ctx),
			"output_schema_name":         types.StringType,
			"profile_metrics_table_name": types.StringType,
			"schedule":                   MonitorCronSchedule{}.Type(ctx),
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot":    MonitorSnapshot{}.Type(ctx),
			"status":      types.StringType,
			"table_name":  types.StringType,
			"time_series": MonitorTimeSeries{}.Type(ctx),
		},
	}
}

// GetCustomMetrics returns the value of the CustomMetrics field in MonitorInfo as
// a slice of MonitorMetric values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo) GetCustomMetrics(ctx context.Context) ([]MonitorMetric, bool) {
	if o.CustomMetrics.IsNull() || o.CustomMetrics.IsUnknown() {
		return nil, false
	}
	var v []MonitorMetric
	d := o.CustomMetrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomMetrics sets the value of the CustomMetrics field in MonitorInfo.
func (o *MonitorInfo) SetCustomMetrics(ctx context.Context, v []MonitorMetric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomMetrics = types.ListValueMust(t, vs)
}

// GetDataClassificationConfig returns the value of the DataClassificationConfig field in MonitorInfo as
// a MonitorDataClassificationConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo) GetDataClassificationConfig(ctx context.Context) (MonitorDataClassificationConfig, bool) {
	var e MonitorDataClassificationConfig
	if o.DataClassificationConfig.IsNull() || o.DataClassificationConfig.IsUnknown() {
		return e, false
	}
	var v []MonitorDataClassificationConfig
	d := o.DataClassificationConfig.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDataClassificationConfig sets the value of the DataClassificationConfig field in MonitorInfo.
func (o *MonitorInfo) SetDataClassificationConfig(ctx context.Context, v MonitorDataClassificationConfig) {
	vs := v.ToObjectValue(ctx)
	o.DataClassificationConfig = vs
}

// GetInferenceLog returns the value of the InferenceLog field in MonitorInfo as
// a MonitorInferenceLog value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo) GetInferenceLog(ctx context.Context) (MonitorInferenceLog, bool) {
	var e MonitorInferenceLog
	if o.InferenceLog.IsNull() || o.InferenceLog.IsUnknown() {
		return e, false
	}
	var v []MonitorInferenceLog
	d := o.InferenceLog.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetInferenceLog sets the value of the InferenceLog field in MonitorInfo.
func (o *MonitorInfo) SetInferenceLog(ctx context.Context, v MonitorInferenceLog) {
	vs := v.ToObjectValue(ctx)
	o.InferenceLog = vs
}

// GetNotifications returns the value of the Notifications field in MonitorInfo as
// a MonitorNotifications value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo) GetNotifications(ctx context.Context) (MonitorNotifications, bool) {
	var e MonitorNotifications
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return e, false
	}
	var v []MonitorNotifications
	d := o.Notifications.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNotifications sets the value of the Notifications field in MonitorInfo.
func (o *MonitorInfo) SetNotifications(ctx context.Context, v MonitorNotifications) {
	vs := v.ToObjectValue(ctx)
	o.Notifications = vs
}

// GetSchedule returns the value of the Schedule field in MonitorInfo as
// a MonitorCronSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo) GetSchedule(ctx context.Context) (MonitorCronSchedule, bool) {
	var e MonitorCronSchedule
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []MonitorCronSchedule
	d := o.Schedule.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSchedule sets the value of the Schedule field in MonitorInfo.
func (o *MonitorInfo) SetSchedule(ctx context.Context, v MonitorCronSchedule) {
	vs := v.ToObjectValue(ctx)
	o.Schedule = vs
}

// GetSlicingExprs returns the value of the SlicingExprs field in MonitorInfo as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo) GetSlicingExprs(ctx context.Context) ([]types.String, bool) {
	if o.SlicingExprs.IsNull() || o.SlicingExprs.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SlicingExprs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSlicingExprs sets the value of the SlicingExprs field in MonitorInfo.
func (o *MonitorInfo) SetSlicingExprs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["slicing_exprs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SlicingExprs = types.ListValueMust(t, vs)
}

// GetSnapshot returns the value of the Snapshot field in MonitorInfo as
// a MonitorSnapshot value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo) GetSnapshot(ctx context.Context) (MonitorSnapshot, bool) {
	var e MonitorSnapshot
	if o.Snapshot.IsNull() || o.Snapshot.IsUnknown() {
		return e, false
	}
	var v []MonitorSnapshot
	d := o.Snapshot.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSnapshot sets the value of the Snapshot field in MonitorInfo.
func (o *MonitorInfo) SetSnapshot(ctx context.Context, v MonitorSnapshot) {
	vs := v.ToObjectValue(ctx)
	o.Snapshot = vs
}

// GetTimeSeries returns the value of the TimeSeries field in MonitorInfo as
// a MonitorTimeSeries value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo) GetTimeSeries(ctx context.Context) (MonitorTimeSeries, bool) {
	var e MonitorTimeSeries
	if o.TimeSeries.IsNull() || o.TimeSeries.IsUnknown() {
		return e, false
	}
	var v []MonitorTimeSeries
	d := o.TimeSeries.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTimeSeries sets the value of the TimeSeries field in MonitorInfo.
func (o *MonitorInfo) SetTimeSeries(ctx context.Context, v MonitorTimeSeries) {
	vs := v.ToObjectValue(ctx)
	o.TimeSeries = vs
}

type MonitorMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	Definition types.String `tfsdk:"definition"`
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	InputColumns types.List `tfsdk:"input_columns"`
	// Name of the metric in the output tables.
	Name types.String `tfsdk:"name"`
	// The output type of the custom metric.
	OutputDataType types.String `tfsdk:"output_data_type"`
	// Can only be one of ``"CUSTOM_METRIC_TYPE_AGGREGATE"``,
	// ``"CUSTOM_METRIC_TYPE_DERIVED"``, or ``"CUSTOM_METRIC_TYPE_DRIFT"``. The
	// ``"CUSTOM_METRIC_TYPE_AGGREGATE"`` and ``"CUSTOM_METRIC_TYPE_DERIVED"``
	// metrics are computed on a single table, whereas the
	// ``"CUSTOM_METRIC_TYPE_DRIFT"`` compare metrics across baseline and input
	// table, or across the two consecutive time windows. -
	// CUSTOM_METRIC_TYPE_AGGREGATE: only depend on the existing columns in your
	// table - CUSTOM_METRIC_TYPE_DERIVED: depend on previously computed
	// aggregate metrics - CUSTOM_METRIC_TYPE_DRIFT: depend on previously
	// computed aggregate or derived metrics
	Type_ types.String `tfsdk:"type"`
}

func (newState *MonitorMetric) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorMetric) {
}

func (newState *MonitorMetric) SyncEffectiveFieldsDuringRead(existingState MonitorMetric) {
}

func (c MonitorMetric) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["definition"] = attrs["definition"].SetRequired()
	attrs["input_columns"] = attrs["input_columns"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["output_data_type"] = attrs["output_data_type"].SetRequired()
	attrs["type"] = attrs["type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorMetric.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorMetric) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorMetric
// only implements ToObjectValue() and Type().
func (o MonitorMetric) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"definition":       o.Definition,
			"input_columns":    o.InputColumns,
			"name":             o.Name,
			"output_data_type": o.OutputDataType,
			"type":             o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorMetric) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"definition": types.StringType,
			"input_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"name":             types.StringType,
			"output_data_type": types.StringType,
			"type":             types.StringType,
		},
	}
}

// GetInputColumns returns the value of the InputColumns field in MonitorMetric as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorMetric) GetInputColumns(ctx context.Context) ([]types.String, bool) {
	if o.InputColumns.IsNull() || o.InputColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InputColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInputColumns sets the value of the InputColumns field in MonitorMetric.
func (o *MonitorMetric) SetInputColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["input_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InputColumns = types.ListValueMust(t, vs)
}

type MonitorNotifications struct {
	// Who to send notifications to on monitor failure.
	OnFailure types.Object `tfsdk:"on_failure"`
	// Who to send notifications to when new data classification tags are
	// detected.
	OnNewClassificationTagDetected types.Object `tfsdk:"on_new_classification_tag_detected"`
}

func (newState *MonitorNotifications) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorNotifications) {
}

func (newState *MonitorNotifications) SyncEffectiveFieldsDuringRead(existingState MonitorNotifications) {
}

func (c MonitorNotifications) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["on_failure"] = attrs["on_failure"].SetOptional()
	attrs["on_new_classification_tag_detected"] = attrs["on_new_classification_tag_detected"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorNotifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorNotifications) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_failure":                         reflect.TypeOf(MonitorDestination{}),
		"on_new_classification_tag_detected": reflect.TypeOf(MonitorDestination{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorNotifications
// only implements ToObjectValue() and Type().
func (o MonitorNotifications) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"on_failure":                         o.OnFailure,
			"on_new_classification_tag_detected": o.OnNewClassificationTagDetected,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorNotifications) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"on_failure":                         MonitorDestination{}.Type(ctx),
			"on_new_classification_tag_detected": MonitorDestination{}.Type(ctx),
		},
	}
}

// GetOnFailure returns the value of the OnFailure field in MonitorNotifications as
// a MonitorDestination value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorNotifications) GetOnFailure(ctx context.Context) (MonitorDestination, bool) {
	var e MonitorDestination
	if o.OnFailure.IsNull() || o.OnFailure.IsUnknown() {
		return e, false
	}
	var v []MonitorDestination
	d := o.OnFailure.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetOnFailure sets the value of the OnFailure field in MonitorNotifications.
func (o *MonitorNotifications) SetOnFailure(ctx context.Context, v MonitorDestination) {
	vs := v.ToObjectValue(ctx)
	o.OnFailure = vs
}

// GetOnNewClassificationTagDetected returns the value of the OnNewClassificationTagDetected field in MonitorNotifications as
// a MonitorDestination value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorNotifications) GetOnNewClassificationTagDetected(ctx context.Context) (MonitorDestination, bool) {
	var e MonitorDestination
	if o.OnNewClassificationTagDetected.IsNull() || o.OnNewClassificationTagDetected.IsUnknown() {
		return e, false
	}
	var v []MonitorDestination
	d := o.OnNewClassificationTagDetected.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetOnNewClassificationTagDetected sets the value of the OnNewClassificationTagDetected field in MonitorNotifications.
func (o *MonitorNotifications) SetOnNewClassificationTagDetected(ctx context.Context, v MonitorDestination) {
	vs := v.ToObjectValue(ctx)
	o.OnNewClassificationTagDetected = vs
}

type MonitorRefreshInfo struct {
	// Time at which refresh operation completed (milliseconds since 1/1/1970
	// UTC).
	EndTimeMs types.Int64 `tfsdk:"end_time_ms"`
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	Message types.String `tfsdk:"message"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"refresh_id"`
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	StartTimeMs types.Int64 `tfsdk:"start_time_ms"`
	// The current state of the refresh.
	State types.String `tfsdk:"state"`
	// The method by which the refresh was triggered.
	Trigger types.String `tfsdk:"trigger"`
}

func (newState *MonitorRefreshInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorRefreshInfo) {
}

func (newState *MonitorRefreshInfo) SyncEffectiveFieldsDuringRead(existingState MonitorRefreshInfo) {
}

func (c MonitorRefreshInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["end_time_ms"] = attrs["end_time_ms"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["refresh_id"] = attrs["refresh_id"].SetRequired()
	attrs["start_time_ms"] = attrs["start_time_ms"].SetRequired()
	attrs["state"] = attrs["state"].SetRequired()
	attrs["trigger"] = attrs["trigger"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorRefreshInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorRefreshInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorRefreshInfo
// only implements ToObjectValue() and Type().
func (o MonitorRefreshInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time_ms":   o.EndTimeMs,
			"message":       o.Message,
			"refresh_id":    o.RefreshId,
			"start_time_ms": o.StartTimeMs,
			"state":         o.State,
			"trigger":       o.Trigger,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorRefreshInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time_ms":   types.Int64Type,
			"message":       types.StringType,
			"refresh_id":    types.Int64Type,
			"start_time_ms": types.Int64Type,
			"state":         types.StringType,
			"trigger":       types.StringType,
		},
	}
}

type MonitorRefreshListResponse struct {
	// List of refreshes.
	Refreshes types.List `tfsdk:"refreshes"`
}

func (newState *MonitorRefreshListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorRefreshListResponse) {
}

func (newState *MonitorRefreshListResponse) SyncEffectiveFieldsDuringRead(existingState MonitorRefreshListResponse) {
}

func (c MonitorRefreshListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["refreshes"] = attrs["refreshes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorRefreshListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorRefreshListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refreshes": reflect.TypeOf(MonitorRefreshInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorRefreshListResponse
// only implements ToObjectValue() and Type().
func (o MonitorRefreshListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"refreshes": o.Refreshes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorRefreshListResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"refreshes": basetypes.ListType{
				ElemType: MonitorRefreshInfo{}.Type(ctx),
			},
		},
	}
}

// GetRefreshes returns the value of the Refreshes field in MonitorRefreshListResponse as
// a slice of MonitorRefreshInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorRefreshListResponse) GetRefreshes(ctx context.Context) ([]MonitorRefreshInfo, bool) {
	if o.Refreshes.IsNull() || o.Refreshes.IsUnknown() {
		return nil, false
	}
	var v []MonitorRefreshInfo
	d := o.Refreshes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefreshes sets the value of the Refreshes field in MonitorRefreshListResponse.
func (o *MonitorRefreshListResponse) SetRefreshes(ctx context.Context, v []MonitorRefreshInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["refreshes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Refreshes = types.ListValueMust(t, vs)
}

type MonitorSnapshot struct {
}

func (newState *MonitorSnapshot) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorSnapshot) {
}

func (newState *MonitorSnapshot) SyncEffectiveFieldsDuringRead(existingState MonitorSnapshot) {
}

func (c MonitorSnapshot) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorSnapshot.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorSnapshot) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorSnapshot
// only implements ToObjectValue() and Type().
func (o MonitorSnapshot) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorSnapshot) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type MonitorTimeSeries struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities types.List `tfsdk:"granularities"`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol types.String `tfsdk:"timestamp_col"`
}

func (newState *MonitorTimeSeries) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorTimeSeries) {
}

func (newState *MonitorTimeSeries) SyncEffectiveFieldsDuringRead(existingState MonitorTimeSeries) {
}

func (c MonitorTimeSeries) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["granularities"] = attrs["granularities"].SetRequired()
	attrs["timestamp_col"] = attrs["timestamp_col"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorTimeSeries.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorTimeSeries) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"granularities": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorTimeSeries
// only implements ToObjectValue() and Type().
func (o MonitorTimeSeries) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"granularities": o.Granularities,
			"timestamp_col": o.TimestampCol,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorTimeSeries) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"granularities": basetypes.ListType{
				ElemType: types.StringType,
			},
			"timestamp_col": types.StringType,
		},
	}
}

// GetGranularities returns the value of the Granularities field in MonitorTimeSeries as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorTimeSeries) GetGranularities(ctx context.Context) ([]types.String, bool) {
	if o.Granularities.IsNull() || o.Granularities.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Granularities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGranularities sets the value of the Granularities field in MonitorTimeSeries.
func (o *MonitorTimeSeries) SetGranularities(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["granularities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Granularities = types.ListValueMust(t, vs)
}

type NamedTableConstraint struct {
	// The name of the constraint.
	Name types.String `tfsdk:"name"`
}

func (newState *NamedTableConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan NamedTableConstraint) {
}

func (newState *NamedTableConstraint) SyncEffectiveFieldsDuringRead(existingState NamedTableConstraint) {
}

func (c NamedTableConstraint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NamedTableConstraint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NamedTableConstraint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NamedTableConstraint
// only implements ToObjectValue() and Type().
func (o NamedTableConstraint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NamedTableConstraint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Custom fields that user can set for pipeline while creating
// SyncedDatabaseTable. Note that other fields of pipeline are still inferred by
// table def internally
type NewPipelineSpec struct {
	// UC catalog for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be a standard catalog where the user has
	// permissions to create Delta tables.
	StorageCatalog types.String `tfsdk:"storage_catalog"`
	// UC schema for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be in the standard catalog where the user
	// has permissions to create Delta tables.
	StorageSchema types.String `tfsdk:"storage_schema"`
}

func (newState *NewPipelineSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan NewPipelineSpec) {
}

func (newState *NewPipelineSpec) SyncEffectiveFieldsDuringRead(existingState NewPipelineSpec) {
}

func (c NewPipelineSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["storage_catalog"] = attrs["storage_catalog"].SetOptional()
	attrs["storage_schema"] = attrs["storage_schema"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NewPipelineSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NewPipelineSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NewPipelineSpec
// only implements ToObjectValue() and Type().
func (o NewPipelineSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_catalog": o.StorageCatalog,
			"storage_schema":  o.StorageSchema,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NewPipelineSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_catalog": types.StringType,
			"storage_schema":  types.StringType,
		},
	}
}

// Online Table information.
type OnlineTable struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"name"`
	// Specification of the online table.
	Spec types.Object `tfsdk:"spec"`
	// Online Table data synchronization status
	Status types.Object `tfsdk:"status"`
	// Data serving REST API URL for this table
	TableServingUrl types.String `tfsdk:"table_serving_url"`
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState types.String `tfsdk:"unity_catalog_provisioning_state"`
}

func (newState *OnlineTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTable) {
}

func (newState *OnlineTable) SyncEffectiveFieldsDuringRead(existingState OnlineTable) {
}

func (c OnlineTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["table_serving_url"] = attrs["table_serving_url"].SetComputed()
	attrs["unity_catalog_provisioning_state"] = attrs["unity_catalog_provisioning_state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OnlineTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(OnlineTableSpec{}),
		"status": reflect.TypeOf(OnlineTableStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineTable
// only implements ToObjectValue() and Type().
func (o OnlineTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":                             o.Name,
			"spec":                             o.Spec,
			"status":                           o.Status,
			"table_serving_url":                o.TableServingUrl,
			"unity_catalog_provisioning_state": o.UnityCatalogProvisioningState,
		})
}

// Type implements basetypes.ObjectValuable.
func (o OnlineTable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":                             types.StringType,
			"spec":                             OnlineTableSpec{}.Type(ctx),
			"status":                           OnlineTableStatus{}.Type(ctx),
			"table_serving_url":                types.StringType,
			"unity_catalog_provisioning_state": types.StringType,
		},
	}
}

// GetSpec returns the value of the Spec field in OnlineTable as
// a OnlineTableSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTable) GetSpec(ctx context.Context) (OnlineTableSpec, bool) {
	var e OnlineTableSpec
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v []OnlineTableSpec
	d := o.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSpec sets the value of the Spec field in OnlineTable.
func (o *OnlineTable) SetSpec(ctx context.Context, v OnlineTableSpec) {
	vs := v.ToObjectValue(ctx)
	o.Spec = vs
}

// GetStatus returns the value of the Status field in OnlineTable as
// a OnlineTableStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTable) GetStatus(ctx context.Context) (OnlineTableStatus, bool) {
	var e OnlineTableStatus
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []OnlineTableStatus
	d := o.Status.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetStatus sets the value of the Status field in OnlineTable.
func (o *OnlineTable) SetStatus(ctx context.Context, v OnlineTableStatus) {
	vs := v.ToObjectValue(ctx)
	o.Status = vs
}

// Specification of an online table.
type OnlineTableSpec struct {
	// Whether to create a full-copy pipeline -- a pipeline that stops after
	// creates a full copy of the source table upon initialization and does not
	// process any change data feeds (CDFs) afterwards. The pipeline can still
	// be manually triggered afterwards, but it always perform a full copy of
	// the source table and there are no incremental updates. This mode is
	// useful for syncing views or tables without CDFs to online tables. Note
	// that the full-copy pipeline only supports "triggered" scheduling policy.
	PerformFullCopy types.Bool `tfsdk:"perform_full_copy"`
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns types.List `tfsdk:"primary_key_columns"`
	// Pipeline runs continuously after generating the initial data.
	RunContinuously types.Object `tfsdk:"run_continuously"`
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	RunTriggered types.Object `tfsdk:"run_triggered"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName types.String `tfsdk:"source_table_full_name"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey types.String `tfsdk:"timeseries_key"`
}

func (newState *OnlineTableSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpec) {
}

func (newState *OnlineTableSpec) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpec) {
}

func (c OnlineTableSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["perform_full_copy"] = attrs["perform_full_copy"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetComputed()
	attrs["primary_key_columns"] = attrs["primary_key_columns"].SetOptional()
	attrs["run_continuously"] = attrs["run_continuously"].SetOptional()
	attrs["run_triggered"] = attrs["run_triggered"].SetOptional()
	attrs["source_table_full_name"] = attrs["source_table_full_name"].SetOptional()
	attrs["timeseries_key"] = attrs["timeseries_key"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineTableSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OnlineTableSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_key_columns": reflect.TypeOf(types.String{}),
		"run_continuously":    reflect.TypeOf(OnlineTableSpecContinuousSchedulingPolicy{}),
		"run_triggered":       reflect.TypeOf(OnlineTableSpecTriggeredSchedulingPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineTableSpec
// only implements ToObjectValue() and Type().
func (o OnlineTableSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"perform_full_copy":      o.PerformFullCopy,
			"pipeline_id":            o.PipelineId,
			"primary_key_columns":    o.PrimaryKeyColumns,
			"run_continuously":       o.RunContinuously,
			"run_triggered":          o.RunTriggered,
			"source_table_full_name": o.SourceTableFullName,
			"timeseries_key":         o.TimeseriesKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o OnlineTableSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"perform_full_copy": types.BoolType,
			"pipeline_id":       types.StringType,
			"primary_key_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"run_continuously":       OnlineTableSpecContinuousSchedulingPolicy{}.Type(ctx),
			"run_triggered":          OnlineTableSpecTriggeredSchedulingPolicy{}.Type(ctx),
			"source_table_full_name": types.StringType,
			"timeseries_key":         types.StringType,
		},
	}
}

// GetPrimaryKeyColumns returns the value of the PrimaryKeyColumns field in OnlineTableSpec as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableSpec) GetPrimaryKeyColumns(ctx context.Context) ([]types.String, bool) {
	if o.PrimaryKeyColumns.IsNull() || o.PrimaryKeyColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PrimaryKeyColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrimaryKeyColumns sets the value of the PrimaryKeyColumns field in OnlineTableSpec.
func (o *OnlineTableSpec) SetPrimaryKeyColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_key_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrimaryKeyColumns = types.ListValueMust(t, vs)
}

// GetRunContinuously returns the value of the RunContinuously field in OnlineTableSpec as
// a OnlineTableSpecContinuousSchedulingPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableSpec) GetRunContinuously(ctx context.Context) (OnlineTableSpecContinuousSchedulingPolicy, bool) {
	var e OnlineTableSpecContinuousSchedulingPolicy
	if o.RunContinuously.IsNull() || o.RunContinuously.IsUnknown() {
		return e, false
	}
	var v []OnlineTableSpecContinuousSchedulingPolicy
	d := o.RunContinuously.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetRunContinuously sets the value of the RunContinuously field in OnlineTableSpec.
func (o *OnlineTableSpec) SetRunContinuously(ctx context.Context, v OnlineTableSpecContinuousSchedulingPolicy) {
	vs := v.ToObjectValue(ctx)
	o.RunContinuously = vs
}

// GetRunTriggered returns the value of the RunTriggered field in OnlineTableSpec as
// a OnlineTableSpecTriggeredSchedulingPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableSpec) GetRunTriggered(ctx context.Context) (OnlineTableSpecTriggeredSchedulingPolicy, bool) {
	var e OnlineTableSpecTriggeredSchedulingPolicy
	if o.RunTriggered.IsNull() || o.RunTriggered.IsUnknown() {
		return e, false
	}
	var v []OnlineTableSpecTriggeredSchedulingPolicy
	d := o.RunTriggered.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetRunTriggered sets the value of the RunTriggered field in OnlineTableSpec.
func (o *OnlineTableSpec) SetRunTriggered(ctx context.Context, v OnlineTableSpecTriggeredSchedulingPolicy) {
	vs := v.ToObjectValue(ctx)
	o.RunTriggered = vs
}

type OnlineTableSpecContinuousSchedulingPolicy struct {
}

func (newState *OnlineTableSpecContinuousSchedulingPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpecContinuousSchedulingPolicy) {
}

func (newState *OnlineTableSpecContinuousSchedulingPolicy) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpecContinuousSchedulingPolicy) {
}

func (c OnlineTableSpecContinuousSchedulingPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineTableSpecContinuousSchedulingPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OnlineTableSpecContinuousSchedulingPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineTableSpecContinuousSchedulingPolicy
// only implements ToObjectValue() and Type().
func (o OnlineTableSpecContinuousSchedulingPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o OnlineTableSpecContinuousSchedulingPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type OnlineTableSpecTriggeredSchedulingPolicy struct {
}

func (newState *OnlineTableSpecTriggeredSchedulingPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpecTriggeredSchedulingPolicy) {
}

func (newState *OnlineTableSpecTriggeredSchedulingPolicy) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpecTriggeredSchedulingPolicy) {
}

func (c OnlineTableSpecTriggeredSchedulingPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineTableSpecTriggeredSchedulingPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OnlineTableSpecTriggeredSchedulingPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineTableSpecTriggeredSchedulingPolicy
// only implements ToObjectValue() and Type().
func (o OnlineTableSpecTriggeredSchedulingPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o OnlineTableSpecTriggeredSchedulingPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Status of an online table.
type OnlineTableStatus struct {
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
	ContinuousUpdateStatus types.Object `tfsdk:"continuous_update_status"`
	// The state of the online table.
	DetailedState types.String `tfsdk:"detailed_state"`
	// Detailed status of an online table. Shown if the online table is in the
	// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
	FailedStatus types.Object `tfsdk:"failed_status"`
	// A text description of the current state of the online table.
	Message types.String `tfsdk:"message"`
	// Detailed status of an online table. Shown if the online table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	ProvisioningStatus types.Object `tfsdk:"provisioning_status"`
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
	TriggeredUpdateStatus types.Object `tfsdk:"triggered_update_status"`
}

func (newState *OnlineTableStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableStatus) {
}

func (newState *OnlineTableStatus) SyncEffectiveFieldsDuringRead(existingState OnlineTableStatus) {
}

func (c OnlineTableStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["continuous_update_status"] = attrs["continuous_update_status"].SetOptional()
	attrs["detailed_state"] = attrs["detailed_state"].SetOptional()
	attrs["failed_status"] = attrs["failed_status"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["provisioning_status"] = attrs["provisioning_status"].SetOptional()
	attrs["triggered_update_status"] = attrs["triggered_update_status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineTableStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OnlineTableStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"continuous_update_status": reflect.TypeOf(ContinuousUpdateStatus{}),
		"failed_status":            reflect.TypeOf(FailedStatus{}),
		"provisioning_status":      reflect.TypeOf(ProvisioningStatus{}),
		"triggered_update_status":  reflect.TypeOf(TriggeredUpdateStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineTableStatus
// only implements ToObjectValue() and Type().
func (o OnlineTableStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"continuous_update_status": o.ContinuousUpdateStatus,
			"detailed_state":           o.DetailedState,
			"failed_status":            o.FailedStatus,
			"message":                  o.Message,
			"provisioning_status":      o.ProvisioningStatus,
			"triggered_update_status":  o.TriggeredUpdateStatus,
		})
}

// Type implements basetypes.ObjectValuable.
func (o OnlineTableStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"continuous_update_status": ContinuousUpdateStatus{}.Type(ctx),
			"detailed_state":           types.StringType,
			"failed_status":            FailedStatus{}.Type(ctx),
			"message":                  types.StringType,
			"provisioning_status":      ProvisioningStatus{}.Type(ctx),
			"triggered_update_status":  TriggeredUpdateStatus{}.Type(ctx),
		},
	}
}

// GetContinuousUpdateStatus returns the value of the ContinuousUpdateStatus field in OnlineTableStatus as
// a ContinuousUpdateStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableStatus) GetContinuousUpdateStatus(ctx context.Context) (ContinuousUpdateStatus, bool) {
	var e ContinuousUpdateStatus
	if o.ContinuousUpdateStatus.IsNull() || o.ContinuousUpdateStatus.IsUnknown() {
		return e, false
	}
	var v []ContinuousUpdateStatus
	d := o.ContinuousUpdateStatus.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetContinuousUpdateStatus sets the value of the ContinuousUpdateStatus field in OnlineTableStatus.
func (o *OnlineTableStatus) SetContinuousUpdateStatus(ctx context.Context, v ContinuousUpdateStatus) {
	vs := v.ToObjectValue(ctx)
	o.ContinuousUpdateStatus = vs
}

// GetFailedStatus returns the value of the FailedStatus field in OnlineTableStatus as
// a FailedStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableStatus) GetFailedStatus(ctx context.Context) (FailedStatus, bool) {
	var e FailedStatus
	if o.FailedStatus.IsNull() || o.FailedStatus.IsUnknown() {
		return e, false
	}
	var v []FailedStatus
	d := o.FailedStatus.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetFailedStatus sets the value of the FailedStatus field in OnlineTableStatus.
func (o *OnlineTableStatus) SetFailedStatus(ctx context.Context, v FailedStatus) {
	vs := v.ToObjectValue(ctx)
	o.FailedStatus = vs
}

// GetProvisioningStatus returns the value of the ProvisioningStatus field in OnlineTableStatus as
// a ProvisioningStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableStatus) GetProvisioningStatus(ctx context.Context) (ProvisioningStatus, bool) {
	var e ProvisioningStatus
	if o.ProvisioningStatus.IsNull() || o.ProvisioningStatus.IsUnknown() {
		return e, false
	}
	var v []ProvisioningStatus
	d := o.ProvisioningStatus.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetProvisioningStatus sets the value of the ProvisioningStatus field in OnlineTableStatus.
func (o *OnlineTableStatus) SetProvisioningStatus(ctx context.Context, v ProvisioningStatus) {
	vs := v.ToObjectValue(ctx)
	o.ProvisioningStatus = vs
}

// GetTriggeredUpdateStatus returns the value of the TriggeredUpdateStatus field in OnlineTableStatus as
// a TriggeredUpdateStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableStatus) GetTriggeredUpdateStatus(ctx context.Context) (TriggeredUpdateStatus, bool) {
	var e TriggeredUpdateStatus
	if o.TriggeredUpdateStatus.IsNull() || o.TriggeredUpdateStatus.IsUnknown() {
		return e, false
	}
	var v []TriggeredUpdateStatus
	d := o.TriggeredUpdateStatus.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTriggeredUpdateStatus sets the value of the TriggeredUpdateStatus field in OnlineTableStatus.
func (o *OnlineTableStatus) SetTriggeredUpdateStatus(ctx context.Context, v TriggeredUpdateStatus) {
	vs := v.ToObjectValue(ctx)
	o.TriggeredUpdateStatus = vs
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

type PermissionsList struct {
	// The privileges assigned to each principal
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments"`
}

func (newState *PermissionsList) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsList) {
}

func (newState *PermissionsList) SyncEffectiveFieldsDuringRead(existingState PermissionsList) {
}

func (c PermissionsList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["privilege_assignments"] = attrs["privilege_assignments"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionsList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionsList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(PrivilegeAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionsList
// only implements ToObjectValue() and Type().
func (o PermissionsList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"privilege_assignments": o.PrivilegeAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionsList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment{}.Type(ctx),
			},
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in PermissionsList as
// a slice of PrivilegeAssignment values.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionsList) GetPrivilegeAssignments(ctx context.Context) ([]PrivilegeAssignment, bool) {
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

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in PermissionsList.
func (o *PermissionsList) SetPrivilegeAssignments(ctx context.Context, v []PrivilegeAssignment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrivilegeAssignments = types.ListValueMust(t, vs)
}

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds types.Float64 `tfsdk:"estimated_completion_time_seconds"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing types.Int64 `tfsdk:"latest_version_currently_processing"`
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion types.Float64 `tfsdk:"sync_progress_completion"`
	// The number of rows that have been synced in this update.
	SyncedRowCount types.Int64 `tfsdk:"synced_row_count"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount types.Int64 `tfsdk:"total_row_count"`
}

func (newState *PipelineProgress) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineProgress) {
}

func (newState *PipelineProgress) SyncEffectiveFieldsDuringRead(existingState PipelineProgress) {
}

func (c PipelineProgress) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["estimated_completion_time_seconds"] = attrs["estimated_completion_time_seconds"].SetOptional()
	attrs["latest_version_currently_processing"] = attrs["latest_version_currently_processing"].SetOptional()
	attrs["sync_progress_completion"] = attrs["sync_progress_completion"].SetOptional()
	attrs["synced_row_count"] = attrs["synced_row_count"].SetOptional()
	attrs["total_row_count"] = attrs["total_row_count"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineProgress.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineProgress) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineProgress
// only implements ToObjectValue() and Type().
func (o PipelineProgress) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"estimated_completion_time_seconds":   o.EstimatedCompletionTimeSeconds,
			"latest_version_currently_processing": o.LatestVersionCurrentlyProcessing,
			"sync_progress_completion":            o.SyncProgressCompletion,
			"synced_row_count":                    o.SyncedRowCount,
			"total_row_count":                     o.TotalRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PipelineProgress) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"estimated_completion_time_seconds":   types.Float64Type,
			"latest_version_currently_processing": types.Int64Type,
			"sync_progress_completion":            types.Float64Type,
			"synced_row_count":                    types.Int64Type,
			"total_row_count":                     types.Int64Type,
		},
	}
}

type PrimaryKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns types.List `tfsdk:"child_columns"`
	// The name of the constraint.
	Name types.String `tfsdk:"name"`
	// Column names that represent a timeseries.
	TimeseriesColumns types.List `tfsdk:"timeseries_columns"`
}

func (newState *PrimaryKeyConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrimaryKeyConstraint) {
}

func (newState *PrimaryKeyConstraint) SyncEffectiveFieldsDuringRead(existingState PrimaryKeyConstraint) {
}

func (c PrimaryKeyConstraint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["child_columns"] = attrs["child_columns"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["timeseries_columns"] = attrs["timeseries_columns"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrimaryKeyConstraint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PrimaryKeyConstraint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"child_columns":      reflect.TypeOf(types.String{}),
		"timeseries_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrimaryKeyConstraint
// only implements ToObjectValue() and Type().
func (o PrimaryKeyConstraint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"child_columns":      o.ChildColumns,
			"name":               o.Name,
			"timeseries_columns": o.TimeseriesColumns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PrimaryKeyConstraint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"child_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"name": types.StringType,
			"timeseries_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetChildColumns returns the value of the ChildColumns field in PrimaryKeyConstraint as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PrimaryKeyConstraint) GetChildColumns(ctx context.Context) ([]types.String, bool) {
	if o.ChildColumns.IsNull() || o.ChildColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ChildColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChildColumns sets the value of the ChildColumns field in PrimaryKeyConstraint.
func (o *PrimaryKeyConstraint) SetChildColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["child_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ChildColumns = types.ListValueMust(t, vs)
}

// GetTimeseriesColumns returns the value of the TimeseriesColumns field in PrimaryKeyConstraint as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PrimaryKeyConstraint) GetTimeseriesColumns(ctx context.Context) ([]types.String, bool) {
	if o.TimeseriesColumns.IsNull() || o.TimeseriesColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.TimeseriesColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTimeseriesColumns sets the value of the TimeseriesColumns field in PrimaryKeyConstraint.
func (o *PrimaryKeyConstraint) SetTimeseriesColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["timeseries_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TimeseriesColumns = types.ListValueMust(t, vs)
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

// Status of an asynchronously provisioned resource.
type ProvisioningInfo struct {
	// The provisioning state of the resource.
	State types.String `tfsdk:"state"`
}

func (newState *ProvisioningInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProvisioningInfo) {
}

func (newState *ProvisioningInfo) SyncEffectiveFieldsDuringRead(existingState ProvisioningInfo) {
}

func (c ProvisioningInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["state"] = attrs["state"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProvisioningInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProvisioningInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProvisioningInfo
// only implements ToObjectValue() and Type().
func (o ProvisioningInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"state": o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ProvisioningInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"state": types.StringType,
		},
	}
}

// Detailed status of an online table. Shown if the online table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type ProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress types.Object `tfsdk:"initial_pipeline_sync_progress"`
}

func (newState *ProvisioningStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProvisioningStatus) {
}

func (newState *ProvisioningStatus) SyncEffectiveFieldsDuringRead(existingState ProvisioningStatus) {
}

func (c ProvisioningStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProvisioningStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProvisioningStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(PipelineProgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProvisioningStatus
// only implements ToObjectValue() and Type().
func (o ProvisioningStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": o.InitialPipelineSyncProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ProvisioningStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": PipelineProgress{}.Type(ctx),
		},
	}
}

// GetInitialPipelineSyncProgress returns the value of the InitialPipelineSyncProgress field in ProvisioningStatus as
// a PipelineProgress value.
// If the field is unknown or null, the boolean return value is false.
func (o *ProvisioningStatus) GetInitialPipelineSyncProgress(ctx context.Context) (PipelineProgress, bool) {
	var e PipelineProgress
	if o.InitialPipelineSyncProgress.IsNull() || o.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v []PipelineProgress
	d := o.InitialPipelineSyncProgress.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in ProvisioningStatus.
func (o *ProvisioningStatus) SetInitialPipelineSyncProgress(ctx context.Context, v PipelineProgress) {
	vs := v.ToObjectValue(ctx)
	o.InitialPipelineSyncProgress = vs
}

type QuotaInfo struct {
	// The timestamp that indicates when the quota count was last updated.
	LastRefreshedAt types.Int64 `tfsdk:"last_refreshed_at"`
	// Name of the parent resource. Returns metastore ID if the parent is a
	// metastore.
	ParentFullName types.String `tfsdk:"parent_full_name"`
	// The quota parent securable type.
	ParentSecurableType types.String `tfsdk:"parent_securable_type"`
	// The current usage of the resource quota.
	QuotaCount types.Int64 `tfsdk:"quota_count"`
	// The current limit of the resource quota.
	QuotaLimit types.Int64 `tfsdk:"quota_limit"`
	// The name of the quota.
	QuotaName types.String `tfsdk:"quota_name"`
}

func (newState *QuotaInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan QuotaInfo) {
}

func (newState *QuotaInfo) SyncEffectiveFieldsDuringRead(existingState QuotaInfo) {
}

func (c QuotaInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_refreshed_at"] = attrs["last_refreshed_at"].SetOptional()
	attrs["parent_full_name"] = attrs["parent_full_name"].SetOptional()
	attrs["parent_securable_type"] = attrs["parent_securable_type"].SetOptional()
	attrs["quota_count"] = attrs["quota_count"].SetOptional()
	attrs["quota_limit"] = attrs["quota_limit"].SetOptional()
	attrs["quota_name"] = attrs["quota_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QuotaInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QuotaInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QuotaInfo
// only implements ToObjectValue() and Type().
func (o QuotaInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_refreshed_at":     o.LastRefreshedAt,
			"parent_full_name":      o.ParentFullName,
			"parent_securable_type": o.ParentSecurableType,
			"quota_count":           o.QuotaCount,
			"quota_limit":           o.QuotaLimit,
			"quota_name":            o.QuotaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QuotaInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_refreshed_at":     types.Int64Type,
			"parent_full_name":      types.StringType,
			"parent_securable_type": types.StringType,
			"quota_count":           types.Int64Type,
			"quota_limit":           types.Int64Type,
			"quota_name":            types.StringType,
		},
	}
}

// R2 temporary credentials for API authentication. Read more at
// https://developers.cloudflare.com/r2/api/s3/tokens/.
type R2Credentials struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId types.String `tfsdk:"access_key_id"`
	// The secret access key associated with the access key.
	SecretAccessKey types.String `tfsdk:"secret_access_key"`
	// The generated JWT that users must pass to use the temporary credentials.
	SessionToken types.String `tfsdk:"session_token"`
}

func (newState *R2Credentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan R2Credentials) {
}

func (newState *R2Credentials) SyncEffectiveFieldsDuringRead(existingState R2Credentials) {
}

func (c R2Credentials) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_key_id"] = attrs["access_key_id"].SetOptional()
	attrs["secret_access_key"] = attrs["secret_access_key"].SetOptional()
	attrs["session_token"] = attrs["session_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in R2Credentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a R2Credentials) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, R2Credentials
// only implements ToObjectValue() and Type().
func (o R2Credentials) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_key_id":     o.AccessKeyId,
			"secret_access_key": o.SecretAccessKey,
			"session_token":     o.SessionToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o R2Credentials) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_key_id":     types.StringType,
			"secret_access_key": types.StringType,
			"session_token":     types.StringType,
		},
	}
}

// Get a Volume
type ReadVolumeRequest struct {
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The three-level (fully qualified) name of the volume
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReadVolumeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReadVolumeRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReadVolumeRequest
// only implements ToObjectValue() and Type().
func (o ReadVolumeRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"name":           o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ReadVolumeRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"name":           types.StringType,
		},
	}
}

type RegenerateDashboardRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Optional argument to specify the warehouse for dashboard regeneration. If
	// not specified, the first running warehouse will be used.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

func (newState *RegenerateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegenerateDashboardRequest) {
}

func (newState *RegenerateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState RegenerateDashboardRequest) {
}

func (c RegenerateDashboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table_name"] = attrs["table_name"].SetRequired()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegenerateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RegenerateDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegenerateDashboardRequest
// only implements ToObjectValue() and Type().
func (o RegenerateDashboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name":   o.TableName,
			"warehouse_id": o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegenerateDashboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name":   types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

type RegenerateDashboardResponse struct {
	// Id of the regenerated monitoring dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The directory where the regenerated dashboard is stored.
	ParentFolder types.String `tfsdk:"parent_folder"`
}

func (newState *RegenerateDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegenerateDashboardResponse) {
}

func (newState *RegenerateDashboardResponse) SyncEffectiveFieldsDuringRead(existingState RegenerateDashboardResponse) {
}

func (c RegenerateDashboardResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dashboard_id"] = attrs["dashboard_id"].SetOptional()
	attrs["parent_folder"] = attrs["parent_folder"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegenerateDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RegenerateDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegenerateDashboardResponse
// only implements ToObjectValue() and Type().
func (o RegenerateDashboardResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":  o.DashboardId,
			"parent_folder": o.ParentFolder,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegenerateDashboardResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":  types.StringType,
			"parent_folder": types.StringType,
		},
	}
}

// Registered model alias.
type RegisteredModelAlias struct {
	// Name of the alias, e.g. 'champion' or 'latest_stable'
	AliasName types.String `tfsdk:"alias_name"`
	// Integer version number of the model version to which this alias points.
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

type RegisteredModelInfo struct {
	// List of aliases associated with the registered model
	Aliases types.List `tfsdk:"aliases"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// The name of the catalog where the schema and the registered model reside
	CatalogName types.String `tfsdk:"catalog_name"`
	// The comment attached to the registered model
	Comment types.String `tfsdk:"comment"`
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// The identifier of the user who created the registered model
	CreatedBy types.String `tfsdk:"created_by"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"full_name"`
	// The unique identifier of the metastore
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The name of the registered model
	Name types.String `tfsdk:"name"`
	// The identifier of the user who owns the registered model
	Owner types.String `tfsdk:"owner"`
	// The name of the schema where the registered model resides
	SchemaName types.String `tfsdk:"schema_name"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation types.String `tfsdk:"storage_location"`
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// The identifier of the user who updated the registered model last time
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *RegisteredModelInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelInfo) {
}

func (newState *RegisteredModelInfo) SyncEffectiveFieldsDuringRead(existingState RegisteredModelInfo) {
}

func (c RegisteredModelInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aliases"] = attrs["aliases"].SetOptional()
	attrs["browse_only"] = attrs["browse_only"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RegisteredModelInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aliases": reflect.TypeOf(RegisteredModelAlias{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelInfo
// only implements ToObjectValue() and Type().
func (o RegisteredModelInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aliases":          o.Aliases,
			"browse_only":      o.BrowseOnly,
			"catalog_name":     o.CatalogName,
			"comment":          o.Comment,
			"created_at":       o.CreatedAt,
			"created_by":       o.CreatedBy,
			"full_name":        o.FullName,
			"metastore_id":     o.MetastoreId,
			"name":             o.Name,
			"owner":            o.Owner,
			"schema_name":      o.SchemaName,
			"storage_location": o.StorageLocation,
			"updated_at":       o.UpdatedAt,
			"updated_by":       o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aliases": basetypes.ListType{
				ElemType: RegisteredModelAlias{}.Type(ctx),
			},
			"browse_only":      types.BoolType,
			"catalog_name":     types.StringType,
			"comment":          types.StringType,
			"created_at":       types.Int64Type,
			"created_by":       types.StringType,
			"full_name":        types.StringType,
			"metastore_id":     types.StringType,
			"name":             types.StringType,
			"owner":            types.StringType,
			"schema_name":      types.StringType,
			"storage_location": types.StringType,
			"updated_at":       types.Int64Type,
			"updated_by":       types.StringType,
		},
	}
}

// GetAliases returns the value of the Aliases field in RegisteredModelInfo as
// a slice of RegisteredModelAlias values.
// If the field is unknown or null, the boolean return value is false.
func (o *RegisteredModelInfo) GetAliases(ctx context.Context) ([]RegisteredModelAlias, bool) {
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

// SetAliases sets the value of the Aliases field in RegisteredModelInfo.
func (o *RegisteredModelInfo) SetAliases(ctx context.Context, v []RegisteredModelAlias) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Aliases = types.ListValueMust(t, vs)
}

// Queue a metric refresh for a monitor
type RunRefreshRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunRefreshRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunRefreshRequest
// only implements ToObjectValue() and Type().
func (o RunRefreshRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunRefreshRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name": types.StringType,
		},
	}
}

// Next ID: 40
type SchemaInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name"`
	// The type of the parent catalog.
	CatalogType types.String `tfsdk:"catalog_type"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Time at which this schema was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of schema creator.
	CreatedBy types.String `tfsdk:"created_by"`

	EffectivePredictiveOptimizationFlag types.Object `tfsdk:"effective_predictive_optimization_flag"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization"`
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	FullName types.String `tfsdk:"full_name"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of schema, relative to parent catalog.
	Name types.String `tfsdk:"name"`
	// Username of current owner of schema.
	Owner types.String `tfsdk:"owner"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`
	// The unique identifier of the schema.
	SchemaId types.String `tfsdk:"schema_id"`
	// Storage location for managed tables within schema.
	StorageLocation types.String `tfsdk:"storage_location"`
	// Storage root URL for managed tables within schema.
	StorageRoot types.String `tfsdk:"storage_root"`
	// Time at which this schema was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified schema.
	UpdatedBy types.String `tfsdk:"updated_by"`
}

func (newState *SchemaInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan SchemaInfo) {
}

func (newState *SchemaInfo) SyncEffectiveFieldsDuringRead(existingState SchemaInfo) {
}

func (c SchemaInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["browse_only"] = attrs["browse_only"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["catalog_type"] = attrs["catalog_type"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["effective_predictive_optimization_flag"] = attrs["effective_predictive_optimization_flag"].SetOptional()
	attrs["enable_predictive_optimization"] = attrs["enable_predictive_optimization"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["schema_id"] = attrs["schema_id"].SetOptional()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
	attrs["storage_root"] = attrs["storage_root"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SchemaInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SchemaInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_predictive_optimization_flag": reflect.TypeOf(EffectivePredictiveOptimizationFlag{}),
		"properties":                             reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SchemaInfo
// only implements ToObjectValue() and Type().
func (o SchemaInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"browse_only":                            o.BrowseOnly,
			"catalog_name":                           o.CatalogName,
			"catalog_type":                           o.CatalogType,
			"comment":                                o.Comment,
			"created_at":                             o.CreatedAt,
			"created_by":                             o.CreatedBy,
			"effective_predictive_optimization_flag": o.EffectivePredictiveOptimizationFlag,
			"enable_predictive_optimization":         o.EnablePredictiveOptimization,
			"full_name":                              o.FullName,
			"metastore_id":                           o.MetastoreId,
			"name":                                   o.Name,
			"owner":                                  o.Owner,
			"properties":                             o.Properties,
			"schema_id":                              o.SchemaId,
			"storage_location":                       o.StorageLocation,
			"storage_root":                           o.StorageRoot,
			"updated_at":                             o.UpdatedAt,
			"updated_by":                             o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SchemaInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":                            types.BoolType,
			"catalog_name":                           types.StringType,
			"catalog_type":                           types.StringType,
			"comment":                                types.StringType,
			"created_at":                             types.Int64Type,
			"created_by":                             types.StringType,
			"effective_predictive_optimization_flag": EffectivePredictiveOptimizationFlag{}.Type(ctx),
			"enable_predictive_optimization":         types.StringType,
			"full_name":                              types.StringType,
			"metastore_id":                           types.StringType,
			"name":                                   types.StringType,
			"owner":                                  types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"schema_id":        types.StringType,
			"storage_location": types.StringType,
			"storage_root":     types.StringType,
			"updated_at":       types.Int64Type,
			"updated_by":       types.StringType,
		},
	}
}

// GetEffectivePredictiveOptimizationFlag returns the value of the EffectivePredictiveOptimizationFlag field in SchemaInfo as
// a EffectivePredictiveOptimizationFlag value.
// If the field is unknown or null, the boolean return value is false.
func (o *SchemaInfo) GetEffectivePredictiveOptimizationFlag(ctx context.Context) (EffectivePredictiveOptimizationFlag, bool) {
	var e EffectivePredictiveOptimizationFlag
	if o.EffectivePredictiveOptimizationFlag.IsNull() || o.EffectivePredictiveOptimizationFlag.IsUnknown() {
		return e, false
	}
	var v []EffectivePredictiveOptimizationFlag
	d := o.EffectivePredictiveOptimizationFlag.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEffectivePredictiveOptimizationFlag sets the value of the EffectivePredictiveOptimizationFlag field in SchemaInfo.
func (o *SchemaInfo) SetEffectivePredictiveOptimizationFlag(ctx context.Context, v EffectivePredictiveOptimizationFlag) {
	vs := v.ToObjectValue(ctx)
	o.EffectivePredictiveOptimizationFlag = vs
}

// GetProperties returns the value of the Properties field in SchemaInfo as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SchemaInfo) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in SchemaInfo.
func (o *SchemaInfo) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type SetArtifactAllowlist struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers types.List `tfsdk:"artifact_matchers"`
	// The artifact type of the allowlist.
	ArtifactType types.String `tfsdk:"-"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of the user who set the artifact allowlist.
	CreatedBy types.String `tfsdk:"created_by"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
}

func (newState *SetArtifactAllowlist) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetArtifactAllowlist) {
}

func (newState *SetArtifactAllowlist) SyncEffectiveFieldsDuringRead(existingState SetArtifactAllowlist) {
}

func (c SetArtifactAllowlist) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["artifact_matchers"] = attrs["artifact_matchers"].SetRequired()
	attrs["artifact_type"] = attrs["artifact_type"].SetRequired()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetArtifactAllowlist.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetArtifactAllowlist) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"artifact_matchers": reflect.TypeOf(ArtifactMatcher{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetArtifactAllowlist
// only implements ToObjectValue() and Type().
func (o SetArtifactAllowlist) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_matchers": o.ArtifactMatchers,
			"artifact_type":     o.ArtifactType,
			"created_at":        o.CreatedAt,
			"created_by":        o.CreatedBy,
			"metastore_id":      o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetArtifactAllowlist) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_matchers": basetypes.ListType{
				ElemType: ArtifactMatcher{}.Type(ctx),
			},
			"artifact_type": types.StringType,
			"created_at":    types.Int64Type,
			"created_by":    types.StringType,
			"metastore_id":  types.StringType,
		},
	}
}

// GetArtifactMatchers returns the value of the ArtifactMatchers field in SetArtifactAllowlist as
// a slice of ArtifactMatcher values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetArtifactAllowlist) GetArtifactMatchers(ctx context.Context) ([]ArtifactMatcher, bool) {
	if o.ArtifactMatchers.IsNull() || o.ArtifactMatchers.IsUnknown() {
		return nil, false
	}
	var v []ArtifactMatcher
	d := o.ArtifactMatchers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetArtifactMatchers sets the value of the ArtifactMatchers field in SetArtifactAllowlist.
func (o *SetArtifactAllowlist) SetArtifactMatchers(ctx context.Context, v []ArtifactMatcher) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["artifact_matchers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ArtifactMatchers = types.ListValueMust(t, vs)
}

type SetRegisteredModelAliasRequest struct {
	// The name of the alias
	Alias types.String `tfsdk:"alias"`
	// Full name of the registered model
	FullName types.String `tfsdk:"full_name"`
	// The version number of the model version to which the alias points
	VersionNum types.Int64 `tfsdk:"version_num"`
}

func (newState *SetRegisteredModelAliasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetRegisteredModelAliasRequest) {
}

func (newState *SetRegisteredModelAliasRequest) SyncEffectiveFieldsDuringRead(existingState SetRegisteredModelAliasRequest) {
}

func (c SetRegisteredModelAliasRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["alias"] = attrs["alias"].SetRequired()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["version_num"] = attrs["version_num"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetRegisteredModelAliasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetRegisteredModelAliasRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetRegisteredModelAliasRequest
// only implements ToObjectValue() and Type().
func (o SetRegisteredModelAliasRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":       o.Alias,
			"full_name":   o.FullName,
			"version_num": o.VersionNum,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetRegisteredModelAliasRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":       types.StringType,
			"full_name":   types.StringType,
			"version_num": types.Int64Type,
		},
	}
}

// Server-Side Encryption properties for clients communicating with AWS s3.
type SseEncryptionDetails struct {
	// Sets the value of the 'x-amz-server-side-encryption' header in S3
	// request.
	Algorithm types.String `tfsdk:"algorithm"`
	// Optional. The ARN of the SSE-KMS key used with the S3 location, when
	// algorithm = "SSE-KMS". Sets the value of the
	// 'x-amz-server-side-encryption-aws-kms-key-id' header.
	AwsKmsKeyArn types.String `tfsdk:"aws_kms_key_arn"`
}

func (newState *SseEncryptionDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan SseEncryptionDetails) {
}

func (newState *SseEncryptionDetails) SyncEffectiveFieldsDuringRead(existingState SseEncryptionDetails) {
}

func (c SseEncryptionDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["algorithm"] = attrs["algorithm"].SetOptional()
	attrs["aws_kms_key_arn"] = attrs["aws_kms_key_arn"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SseEncryptionDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SseEncryptionDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SseEncryptionDetails
// only implements ToObjectValue() and Type().
func (o SseEncryptionDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"algorithm":       o.Algorithm,
			"aws_kms_key_arn": o.AwsKmsKeyArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SseEncryptionDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"algorithm":       types.StringType,
			"aws_kms_key_arn": types.StringType,
		},
	}
}

type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	AwsIamRole types.Object `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.Object `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.Object `tfsdk:"azure_service_principal"`
	// The Cloudflare API token configuration.
	CloudflareApiToken types.Object `tfsdk:"cloudflare_api_token"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment"`
	// Time at which this Credential was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of credential creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount types.Object `tfsdk:"databricks_gcp_service_account"`
	// The full name of the credential.
	FullName types.String `tfsdk:"full_name"`
	// The unique identifier of the credential.
	Id types.String `tfsdk:"id"`

	IsolationMode types.String `tfsdk:"isolation_mode"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The credential name. The name must be unique within the metastore.
	Name types.String `tfsdk:"name"`
	// Username of current owner of credential.
	Owner types.String `tfsdk:"owner"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified the credential.
	UpdatedBy types.String `tfsdk:"updated_by"`
	// Whether this credential is the current metastore's root storage
	// credential.
	UsedForManagedStorage types.Bool `tfsdk:"used_for_managed_storage"`
}

func (newState *StorageCredentialInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan StorageCredentialInfo) {
}

func (newState *StorageCredentialInfo) SyncEffectiveFieldsDuringRead(existingState StorageCredentialInfo) {
}

func (c StorageCredentialInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_iam_role"] = attrs["aws_iam_role"].SetOptional()
	attrs["azure_managed_identity"] = attrs["azure_managed_identity"].SetOptional()
	attrs["azure_service_principal"] = attrs["azure_service_principal"].SetOptional()
	attrs["cloudflare_api_token"] = attrs["cloudflare_api_token"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["databricks_gcp_service_account"] = attrs["databricks_gcp_service_account"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["isolation_mode"] = attrs["isolation_mode"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()
	attrs["used_for_managed_storage"] = attrs["used_for_managed_storage"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StorageCredentialInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StorageCredentialInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRoleResponse{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentityResponse{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal{}),
		"cloudflare_api_token":           reflect.TypeOf(CloudflareApiToken{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccountResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StorageCredentialInfo
// only implements ToObjectValue() and Type().
func (o StorageCredentialInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_iam_role":                   o.AwsIamRole,
			"azure_managed_identity":         o.AzureManagedIdentity,
			"azure_service_principal":        o.AzureServicePrincipal,
			"cloudflare_api_token":           o.CloudflareApiToken,
			"comment":                        o.Comment,
			"created_at":                     o.CreatedAt,
			"created_by":                     o.CreatedBy,
			"databricks_gcp_service_account": o.DatabricksGcpServiceAccount,
			"full_name":                      o.FullName,
			"id":                             o.Id,
			"isolation_mode":                 o.IsolationMode,
			"metastore_id":                   o.MetastoreId,
			"name":                           o.Name,
			"owner":                          o.Owner,
			"read_only":                      o.ReadOnly,
			"updated_at":                     o.UpdatedAt,
			"updated_by":                     o.UpdatedBy,
			"used_for_managed_storage":       o.UsedForManagedStorage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StorageCredentialInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role":                   AwsIamRoleResponse{}.Type(ctx),
			"azure_managed_identity":         AzureManagedIdentityResponse{}.Type(ctx),
			"azure_service_principal":        AzureServicePrincipal{}.Type(ctx),
			"cloudflare_api_token":           CloudflareApiToken{}.Type(ctx),
			"comment":                        types.StringType,
			"created_at":                     types.Int64Type,
			"created_by":                     types.StringType,
			"databricks_gcp_service_account": DatabricksGcpServiceAccountResponse{}.Type(ctx),
			"full_name":                      types.StringType,
			"id":                             types.StringType,
			"isolation_mode":                 types.StringType,
			"metastore_id":                   types.StringType,
			"name":                           types.StringType,
			"owner":                          types.StringType,
			"read_only":                      types.BoolType,
			"updated_at":                     types.Int64Type,
			"updated_by":                     types.StringType,
			"used_for_managed_storage":       types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in StorageCredentialInfo as
// a AwsIamRoleResponse value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageCredentialInfo) GetAwsIamRole(ctx context.Context) (AwsIamRoleResponse, bool) {
	var e AwsIamRoleResponse
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRoleResponse
	d := o.AwsIamRole.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsIamRole sets the value of the AwsIamRole field in StorageCredentialInfo.
func (o *StorageCredentialInfo) SetAwsIamRole(ctx context.Context, v AwsIamRoleResponse) {
	vs := v.ToObjectValue(ctx)
	o.AwsIamRole = vs
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in StorageCredentialInfo as
// a AzureManagedIdentityResponse value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageCredentialInfo) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentityResponse, bool) {
	var e AzureManagedIdentityResponse
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentityResponse
	d := o.AzureManagedIdentity.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in StorageCredentialInfo.
func (o *StorageCredentialInfo) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentityResponse) {
	vs := v.ToObjectValue(ctx)
	o.AzureManagedIdentity = vs
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in StorageCredentialInfo as
// a AzureServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageCredentialInfo) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal, bool) {
	var e AzureServicePrincipal
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal
	d := o.AzureServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in StorageCredentialInfo.
func (o *StorageCredentialInfo) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	o.AzureServicePrincipal = vs
}

// GetCloudflareApiToken returns the value of the CloudflareApiToken field in StorageCredentialInfo as
// a CloudflareApiToken value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageCredentialInfo) GetCloudflareApiToken(ctx context.Context) (CloudflareApiToken, bool) {
	var e CloudflareApiToken
	if o.CloudflareApiToken.IsNull() || o.CloudflareApiToken.IsUnknown() {
		return e, false
	}
	var v []CloudflareApiToken
	d := o.CloudflareApiToken.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetCloudflareApiToken sets the value of the CloudflareApiToken field in StorageCredentialInfo.
func (o *StorageCredentialInfo) SetCloudflareApiToken(ctx context.Context, v CloudflareApiToken) {
	vs := v.ToObjectValue(ctx)
	o.CloudflareApiToken = vs
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in StorageCredentialInfo as
// a DatabricksGcpServiceAccountResponse value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageCredentialInfo) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccountResponse, bool) {
	var e DatabricksGcpServiceAccountResponse
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccountResponse
	d := o.DatabricksGcpServiceAccount.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in StorageCredentialInfo.
func (o *StorageCredentialInfo) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccountResponse) {
	vs := v.ToObjectValue(ctx)
	o.DatabricksGcpServiceAccount = vs
}

// Next field marker: 10
type SyncedDatabaseTable struct {
	// Synced Table data synchronization status
	DataSynchronizationStatus types.Object `tfsdk:"data_synchronization_status"`
	// Name of the target database instance. This is required when creating
	// synced database tables in standard catalogs. This is optional when
	// creating synced database tables in registered catalogs. If this field is
	// specified when creating synced database tables in registered catalogs,
	// the database instance name MUST match that of the registered catalog (or
	// the request will be rejected).
	DatabaseInstanceName types.String `tfsdk:"database_instance_name"`
	// Target Postgres database object (logical database) name for this table.
	// This field is optional in all scenarios.
	//
	// When creating a synced table in a registered Postgres catalog, the target
	// Postgres database name is inferred to be that of the registered catalog.
	// If this field is specified in this scenario, the Postgres database name
	// MUST match that of the registered catalog (or the request will be
	// rejected).
	//
	// When creating a synced table in a standard catalog, the target database
	// name is inferred to be that of the standard catalog. In this scenario,
	// specifying this field will allow targeting an arbitrary postgres
	// database.
	LogicalDatabaseName types.String `tfsdk:"logical_database_name"`
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"name"`
	// Specification of a synced database table.
	Spec types.Object `tfsdk:"spec"`
	// Data serving REST API URL for this table
	TableServingUrl types.String `tfsdk:"table_serving_url"`
	// The provisioning state of the synced table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState types.String `tfsdk:"unity_catalog_provisioning_state"`
}

func (newState *SyncedDatabaseTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedDatabaseTable) {
}

func (newState *SyncedDatabaseTable) SyncEffectiveFieldsDuringRead(existingState SyncedDatabaseTable) {
}

func (c SyncedDatabaseTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_synchronization_status"] = attrs["data_synchronization_status"].SetComputed()
	attrs["database_instance_name"] = attrs["database_instance_name"].SetOptional()
	attrs["logical_database_name"] = attrs["logical_database_name"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["table_serving_url"] = attrs["table_serving_url"].SetComputed()
	attrs["unity_catalog_provisioning_state"] = attrs["unity_catalog_provisioning_state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedDatabaseTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncedDatabaseTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_synchronization_status": reflect.TypeOf(OnlineTableStatus{}),
		"spec":                        reflect.TypeOf(SyncedTableSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedDatabaseTable
// only implements ToObjectValue() and Type().
func (o SyncedDatabaseTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_synchronization_status":      o.DataSynchronizationStatus,
			"database_instance_name":           o.DatabaseInstanceName,
			"logical_database_name":            o.LogicalDatabaseName,
			"name":                             o.Name,
			"spec":                             o.Spec,
			"table_serving_url":                o.TableServingUrl,
			"unity_catalog_provisioning_state": o.UnityCatalogProvisioningState,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedDatabaseTable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_synchronization_status":      OnlineTableStatus{}.Type(ctx),
			"database_instance_name":           types.StringType,
			"logical_database_name":            types.StringType,
			"name":                             types.StringType,
			"spec":                             SyncedTableSpec{}.Type(ctx),
			"table_serving_url":                types.StringType,
			"unity_catalog_provisioning_state": types.StringType,
		},
	}
}

// GetDataSynchronizationStatus returns the value of the DataSynchronizationStatus field in SyncedDatabaseTable as
// a OnlineTableStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedDatabaseTable) GetDataSynchronizationStatus(ctx context.Context) (OnlineTableStatus, bool) {
	var e OnlineTableStatus
	if o.DataSynchronizationStatus.IsNull() || o.DataSynchronizationStatus.IsUnknown() {
		return e, false
	}
	var v []OnlineTableStatus
	d := o.DataSynchronizationStatus.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDataSynchronizationStatus sets the value of the DataSynchronizationStatus field in SyncedDatabaseTable.
func (o *SyncedDatabaseTable) SetDataSynchronizationStatus(ctx context.Context, v OnlineTableStatus) {
	vs := v.ToObjectValue(ctx)
	o.DataSynchronizationStatus = vs
}

// GetSpec returns the value of the Spec field in SyncedDatabaseTable as
// a SyncedTableSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedDatabaseTable) GetSpec(ctx context.Context) (SyncedTableSpec, bool) {
	var e SyncedTableSpec
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v []SyncedTableSpec
	d := o.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSpec sets the value of the Spec field in SyncedDatabaseTable.
func (o *SyncedDatabaseTable) SetSpec(ctx context.Context, v SyncedTableSpec) {
	vs := v.ToObjectValue(ctx)
	o.Spec = vs
}

// Specification of a synced database table.
type SyncedTableSpec struct {
	// If true, the synced table's logical database and schema resources in PG
	// will be created if they do not already exist.
	CreateDatabaseObjectsIfMissing types.Bool `tfsdk:"create_database_objects_if_missing"`
	// Spec of new pipeline. Should be empty if pipeline_id is set
	NewPipelineSpec types.Object `tfsdk:"new_pipeline_spec"`
	// ID of the associated pipeline. Should be empty if new_pipeline_spec is
	// set
	PipelineId types.String `tfsdk:"pipeline_id"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns types.List `tfsdk:"primary_key_columns"`
	// Scheduling policy of the underlying pipeline.
	SchedulingPolicy types.String `tfsdk:"scheduling_policy"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName types.String `tfsdk:"source_table_full_name"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey types.String `tfsdk:"timeseries_key"`
}

func (newState *SyncedTableSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableSpec) {
}

func (newState *SyncedTableSpec) SyncEffectiveFieldsDuringRead(existingState SyncedTableSpec) {
}

func (c SyncedTableSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_database_objects_if_missing"] = attrs["create_database_objects_if_missing"].SetOptional()
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()
	attrs["primary_key_columns"] = attrs["primary_key_columns"].SetOptional()
	attrs["scheduling_policy"] = attrs["scheduling_policy"].SetOptional()
	attrs["source_table_full_name"] = attrs["source_table_full_name"].SetOptional()
	attrs["timeseries_key"] = attrs["timeseries_key"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncedTableSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"new_pipeline_spec":   reflect.TypeOf(NewPipelineSpec{}),
		"primary_key_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableSpec
// only implements ToObjectValue() and Type().
func (o SyncedTableSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_database_objects_if_missing": o.CreateDatabaseObjectsIfMissing,
			"new_pipeline_spec":                  o.NewPipelineSpec,
			"pipeline_id":                        o.PipelineId,
			"primary_key_columns":                o.PrimaryKeyColumns,
			"scheduling_policy":                  o.SchedulingPolicy,
			"source_table_full_name":             o.SourceTableFullName,
			"timeseries_key":                     o.TimeseriesKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_database_objects_if_missing": types.BoolType,
			"new_pipeline_spec":                  NewPipelineSpec{}.Type(ctx),
			"pipeline_id":                        types.StringType,
			"primary_key_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"scheduling_policy":      types.StringType,
			"source_table_full_name": types.StringType,
			"timeseries_key":         types.StringType,
		},
	}
}

// GetNewPipelineSpec returns the value of the NewPipelineSpec field in SyncedTableSpec as
// a NewPipelineSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableSpec) GetNewPipelineSpec(ctx context.Context) (NewPipelineSpec, bool) {
	var e NewPipelineSpec
	if o.NewPipelineSpec.IsNull() || o.NewPipelineSpec.IsUnknown() {
		return e, false
	}
	var v []NewPipelineSpec
	d := o.NewPipelineSpec.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNewPipelineSpec sets the value of the NewPipelineSpec field in SyncedTableSpec.
func (o *SyncedTableSpec) SetNewPipelineSpec(ctx context.Context, v NewPipelineSpec) {
	vs := v.ToObjectValue(ctx)
	o.NewPipelineSpec = vs
}

// GetPrimaryKeyColumns returns the value of the PrimaryKeyColumns field in SyncedTableSpec as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableSpec) GetPrimaryKeyColumns(ctx context.Context) ([]types.String, bool) {
	if o.PrimaryKeyColumns.IsNull() || o.PrimaryKeyColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PrimaryKeyColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrimaryKeyColumns sets the value of the PrimaryKeyColumns field in SyncedTableSpec.
func (o *SyncedTableSpec) SetPrimaryKeyColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_key_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrimaryKeyColumns = types.ListValueMust(t, vs)
}

type SystemSchemaInfo struct {
	// Name of the system schema.
	Schema types.String `tfsdk:"schema"`
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in. Possible
	// values: AVAILABLE | ENABLE_INITIALIZED | ENABLE_COMPLETED |
	// DISABLE_INITIALIZED | UNAVAILABLE
	State types.String `tfsdk:"state"`
}

func (newState *SystemSchemaInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan SystemSchemaInfo) {
}

func (newState *SystemSchemaInfo) SyncEffectiveFieldsDuringRead(existingState SystemSchemaInfo) {
}

func (c SystemSchemaInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["schema"] = attrs["schema"].SetRequired()
	attrs["state"] = attrs["state"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SystemSchemaInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SystemSchemaInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SystemSchemaInfo
// only implements ToObjectValue() and Type().
func (o SystemSchemaInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"schema": o.Schema,
			"state":  o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SystemSchemaInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"schema": types.StringType,
			"state":  types.StringType,
		},
	}
}

// A table constraint, as defined by *one* of the following fields being set:
// __primary_key_constraint__, __foreign_key_constraint__,
// __named_table_constraint__.
type TableConstraint struct {
	ForeignKeyConstraint types.Object `tfsdk:"foreign_key_constraint"`

	NamedTableConstraint types.Object `tfsdk:"named_table_constraint"`

	PrimaryKeyConstraint types.Object `tfsdk:"primary_key_constraint"`
}

func (newState *TableConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableConstraint) {
}

func (newState *TableConstraint) SyncEffectiveFieldsDuringRead(existingState TableConstraint) {
}

func (c TableConstraint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["foreign_key_constraint"] = attrs["foreign_key_constraint"].SetOptional()
	attrs["named_table_constraint"] = attrs["named_table_constraint"].SetOptional()
	attrs["primary_key_constraint"] = attrs["primary_key_constraint"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableConstraint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableConstraint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"foreign_key_constraint": reflect.TypeOf(ForeignKeyConstraint{}),
		"named_table_constraint": reflect.TypeOf(NamedTableConstraint{}),
		"primary_key_constraint": reflect.TypeOf(PrimaryKeyConstraint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableConstraint
// only implements ToObjectValue() and Type().
func (o TableConstraint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"foreign_key_constraint": o.ForeignKeyConstraint,
			"named_table_constraint": o.NamedTableConstraint,
			"primary_key_constraint": o.PrimaryKeyConstraint,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableConstraint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"foreign_key_constraint": ForeignKeyConstraint{}.Type(ctx),
			"named_table_constraint": NamedTableConstraint{}.Type(ctx),
			"primary_key_constraint": PrimaryKeyConstraint{}.Type(ctx),
		},
	}
}

// GetForeignKeyConstraint returns the value of the ForeignKeyConstraint field in TableConstraint as
// a ForeignKeyConstraint value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableConstraint) GetForeignKeyConstraint(ctx context.Context) (ForeignKeyConstraint, bool) {
	var e ForeignKeyConstraint
	if o.ForeignKeyConstraint.IsNull() || o.ForeignKeyConstraint.IsUnknown() {
		return e, false
	}
	var v []ForeignKeyConstraint
	d := o.ForeignKeyConstraint.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetForeignKeyConstraint sets the value of the ForeignKeyConstraint field in TableConstraint.
func (o *TableConstraint) SetForeignKeyConstraint(ctx context.Context, v ForeignKeyConstraint) {
	vs := v.ToObjectValue(ctx)
	o.ForeignKeyConstraint = vs
}

// GetNamedTableConstraint returns the value of the NamedTableConstraint field in TableConstraint as
// a NamedTableConstraint value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableConstraint) GetNamedTableConstraint(ctx context.Context) (NamedTableConstraint, bool) {
	var e NamedTableConstraint
	if o.NamedTableConstraint.IsNull() || o.NamedTableConstraint.IsUnknown() {
		return e, false
	}
	var v []NamedTableConstraint
	d := o.NamedTableConstraint.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNamedTableConstraint sets the value of the NamedTableConstraint field in TableConstraint.
func (o *TableConstraint) SetNamedTableConstraint(ctx context.Context, v NamedTableConstraint) {
	vs := v.ToObjectValue(ctx)
	o.NamedTableConstraint = vs
}

// GetPrimaryKeyConstraint returns the value of the PrimaryKeyConstraint field in TableConstraint as
// a PrimaryKeyConstraint value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableConstraint) GetPrimaryKeyConstraint(ctx context.Context) (PrimaryKeyConstraint, bool) {
	var e PrimaryKeyConstraint
	if o.PrimaryKeyConstraint.IsNull() || o.PrimaryKeyConstraint.IsUnknown() {
		return e, false
	}
	var v []PrimaryKeyConstraint
	d := o.PrimaryKeyConstraint.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPrimaryKeyConstraint sets the value of the PrimaryKeyConstraint field in TableConstraint.
func (o *TableConstraint) SetPrimaryKeyConstraint(ctx context.Context, v PrimaryKeyConstraint) {
	vs := v.ToObjectValue(ctx)
	o.PrimaryKeyConstraint = vs
}

// A table that is dependent on a SQL object.
type TableDependency struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	TableFullName types.String `tfsdk:"table_full_name"`
}

func (newState *TableDependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableDependency) {
}

func (newState *TableDependency) SyncEffectiveFieldsDuringRead(existingState TableDependency) {
}

func (c TableDependency) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table_full_name"] = attrs["table_full_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableDependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableDependency) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableDependency
// only implements ToObjectValue() and Type().
func (o TableDependency) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_full_name": o.TableFullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableDependency) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_full_name": types.StringType,
		},
	}
}

type TableExistsResponse struct {
	// Whether the table exists or not.
	TableExists types.Bool `tfsdk:"table_exists"`
}

func (newState *TableExistsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableExistsResponse) {
}

func (newState *TableExistsResponse) SyncEffectiveFieldsDuringRead(existingState TableExistsResponse) {
}

func (c TableExistsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table_exists"] = attrs["table_exists"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableExistsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableExistsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableExistsResponse
// only implements ToObjectValue() and Type().
func (o TableExistsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_exists": o.TableExists,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableExistsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_exists": types.BoolType,
		},
	}
}

type TableInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns types.List `tfsdk:"columns"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Time at which this table was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// Username of table creator.
	CreatedBy types.String `tfsdk:"created_by"`
	// Unique ID of the Data Access Configuration to use with the table data.
	DataAccessConfigurationId types.String `tfsdk:"data_access_configuration_id"`
	// Data source format
	DataSourceFormat types.String `tfsdk:"data_source_format"`
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	DeletedAt types.Int64 `tfsdk:"deleted_at"`
	// Information pertaining to current state of the delta table.
	DeltaRuntimePropertiesKvpairs types.Object `tfsdk:"delta_runtime_properties_kvpairs"`

	EffectivePredictiveOptimizationFlag types.Object `tfsdk:"effective_predictive_optimization_flag"`

	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails types.Object `tfsdk:"encryption_details"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName types.String `tfsdk:"full_name"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// Name of table, relative to parent schema.
	Name types.String `tfsdk:"name"`
	// Username of current owner of table.
	Owner types.String `tfsdk:"owner"`
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	PipelineId types.String `tfsdk:"pipeline_id"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`

	RowFilter types.Object `tfsdk:"row_filter"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath types.String `tfsdk:"sql_path"`
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	StorageCredentialName types.String `tfsdk:"storage_credential_name"`
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables)
	StorageLocation types.String `tfsdk:"storage_location"`
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	TableConstraints types.List `tfsdk:"table_constraints"`
	// The unique identifier of the table.
	TableId types.String `tfsdk:"table_id"`

	TableType types.String `tfsdk:"table_type"`
	// Time at which this table was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// Username of user who last modified the table.
	UpdatedBy types.String `tfsdk:"updated_by"`
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	ViewDefinition types.String `tfsdk:"view_definition"`
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	ViewDependencies types.Object `tfsdk:"view_dependencies"`
}

func (newState *TableInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableInfo) {
}

func (newState *TableInfo) SyncEffectiveFieldsDuringRead(existingState TableInfo) {
}

func (c TableInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_point"] = attrs["access_point"].SetOptional()
	attrs["browse_only"] = attrs["browse_only"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["columns"] = attrs["columns"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["data_access_configuration_id"] = attrs["data_access_configuration_id"].SetOptional()
	attrs["data_source_format"] = attrs["data_source_format"].SetOptional()
	attrs["deleted_at"] = attrs["deleted_at"].SetOptional()
	attrs["delta_runtime_properties_kvpairs"] = attrs["delta_runtime_properties_kvpairs"].SetOptional()
	attrs["effective_predictive_optimization_flag"] = attrs["effective_predictive_optimization_flag"].SetOptional()
	attrs["enable_predictive_optimization"] = attrs["enable_predictive_optimization"].SetOptional()
	attrs["encryption_details"] = attrs["encryption_details"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()
	attrs["properties"] = attrs["properties"].SetOptional()
	attrs["row_filter"] = attrs["row_filter"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["sql_path"] = attrs["sql_path"].SetOptional()
	attrs["storage_credential_name"] = attrs["storage_credential_name"].SetOptional()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
	attrs["table_constraints"] = attrs["table_constraints"].SetOptional()
	attrs["table_id"] = attrs["table_id"].SetOptional()
	attrs["table_type"] = attrs["table_type"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()
	attrs["view_definition"] = attrs["view_definition"].SetOptional()
	attrs["view_dependencies"] = attrs["view_dependencies"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns":                                reflect.TypeOf(ColumnInfo{}),
		"delta_runtime_properties_kvpairs":       reflect.TypeOf(DeltaRuntimePropertiesKvPairs{}),
		"effective_predictive_optimization_flag": reflect.TypeOf(EffectivePredictiveOptimizationFlag{}),
		"encryption_details":                     reflect.TypeOf(EncryptionDetails{}),
		"properties":                             reflect.TypeOf(types.String{}),
		"row_filter":                             reflect.TypeOf(TableRowFilter{}),
		"table_constraints":                      reflect.TypeOf(TableConstraint{}),
		"view_dependencies":                      reflect.TypeOf(DependencyList{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableInfo
// only implements ToObjectValue() and Type().
func (o TableInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_point":                           o.AccessPoint,
			"browse_only":                            o.BrowseOnly,
			"catalog_name":                           o.CatalogName,
			"columns":                                o.Columns,
			"comment":                                o.Comment,
			"created_at":                             o.CreatedAt,
			"created_by":                             o.CreatedBy,
			"data_access_configuration_id":           o.DataAccessConfigurationId,
			"data_source_format":                     o.DataSourceFormat,
			"deleted_at":                             o.DeletedAt,
			"delta_runtime_properties_kvpairs":       o.DeltaRuntimePropertiesKvpairs,
			"effective_predictive_optimization_flag": o.EffectivePredictiveOptimizationFlag,
			"enable_predictive_optimization":         o.EnablePredictiveOptimization,
			"encryption_details":                     o.EncryptionDetails,
			"full_name":                              o.FullName,
			"metastore_id":                           o.MetastoreId,
			"name":                                   o.Name,
			"owner":                                  o.Owner,
			"pipeline_id":                            o.PipelineId,
			"properties":                             o.Properties,
			"row_filter":                             o.RowFilter,
			"schema_name":                            o.SchemaName,
			"sql_path":                               o.SqlPath,
			"storage_credential_name":                o.StorageCredentialName,
			"storage_location":                       o.StorageLocation,
			"table_constraints":                      o.TableConstraints,
			"table_id":                               o.TableId,
			"table_type":                             o.TableType,
			"updated_at":                             o.UpdatedAt,
			"updated_by":                             o.UpdatedBy,
			"view_definition":                        o.ViewDefinition,
			"view_dependencies":                      o.ViewDependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point": types.StringType,
			"browse_only":  types.BoolType,
			"catalog_name": types.StringType,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo{}.Type(ctx),
			},
			"comment":                                types.StringType,
			"created_at":                             types.Int64Type,
			"created_by":                             types.StringType,
			"data_access_configuration_id":           types.StringType,
			"data_source_format":                     types.StringType,
			"deleted_at":                             types.Int64Type,
			"delta_runtime_properties_kvpairs":       DeltaRuntimePropertiesKvPairs{}.Type(ctx),
			"effective_predictive_optimization_flag": EffectivePredictiveOptimizationFlag{}.Type(ctx),
			"enable_predictive_optimization":         types.StringType,
			"encryption_details":                     EncryptionDetails{}.Type(ctx),
			"full_name":                              types.StringType,
			"metastore_id":                           types.StringType,
			"name":                                   types.StringType,
			"owner":                                  types.StringType,
			"pipeline_id":                            types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"row_filter":              TableRowFilter{}.Type(ctx),
			"schema_name":             types.StringType,
			"sql_path":                types.StringType,
			"storage_credential_name": types.StringType,
			"storage_location":        types.StringType,
			"table_constraints": basetypes.ListType{
				ElemType: TableConstraint{}.Type(ctx),
			},
			"table_id":          types.StringType,
			"table_type":        types.StringType,
			"updated_at":        types.Int64Type,
			"updated_by":        types.StringType,
			"view_definition":   types.StringType,
			"view_dependencies": DependencyList{}.Type(ctx),
		},
	}
}

// GetColumns returns the value of the Columns field in TableInfo as
// a slice of ColumnInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo) GetColumns(ctx context.Context) ([]ColumnInfo, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in TableInfo.
func (o *TableInfo) SetColumns(ctx context.Context, v []ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

// GetDeltaRuntimePropertiesKvpairs returns the value of the DeltaRuntimePropertiesKvpairs field in TableInfo as
// a DeltaRuntimePropertiesKvPairs value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo) GetDeltaRuntimePropertiesKvpairs(ctx context.Context) (DeltaRuntimePropertiesKvPairs, bool) {
	var e DeltaRuntimePropertiesKvPairs
	if o.DeltaRuntimePropertiesKvpairs.IsNull() || o.DeltaRuntimePropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []DeltaRuntimePropertiesKvPairs
	d := o.DeltaRuntimePropertiesKvpairs.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDeltaRuntimePropertiesKvpairs sets the value of the DeltaRuntimePropertiesKvpairs field in TableInfo.
func (o *TableInfo) SetDeltaRuntimePropertiesKvpairs(ctx context.Context, v DeltaRuntimePropertiesKvPairs) {
	vs := v.ToObjectValue(ctx)
	o.DeltaRuntimePropertiesKvpairs = vs
}

// GetEffectivePredictiveOptimizationFlag returns the value of the EffectivePredictiveOptimizationFlag field in TableInfo as
// a EffectivePredictiveOptimizationFlag value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo) GetEffectivePredictiveOptimizationFlag(ctx context.Context) (EffectivePredictiveOptimizationFlag, bool) {
	var e EffectivePredictiveOptimizationFlag
	if o.EffectivePredictiveOptimizationFlag.IsNull() || o.EffectivePredictiveOptimizationFlag.IsUnknown() {
		return e, false
	}
	var v []EffectivePredictiveOptimizationFlag
	d := o.EffectivePredictiveOptimizationFlag.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEffectivePredictiveOptimizationFlag sets the value of the EffectivePredictiveOptimizationFlag field in TableInfo.
func (o *TableInfo) SetEffectivePredictiveOptimizationFlag(ctx context.Context, v EffectivePredictiveOptimizationFlag) {
	vs := v.ToObjectValue(ctx)
	o.EffectivePredictiveOptimizationFlag = vs
}

// GetEncryptionDetails returns the value of the EncryptionDetails field in TableInfo as
// a EncryptionDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo) GetEncryptionDetails(ctx context.Context) (EncryptionDetails, bool) {
	var e EncryptionDetails
	if o.EncryptionDetails.IsNull() || o.EncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []EncryptionDetails
	d := o.EncryptionDetails.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEncryptionDetails sets the value of the EncryptionDetails field in TableInfo.
func (o *TableInfo) SetEncryptionDetails(ctx context.Context, v EncryptionDetails) {
	vs := v.ToObjectValue(ctx)
	o.EncryptionDetails = vs
}

// GetProperties returns the value of the Properties field in TableInfo as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in TableInfo.
func (o *TableInfo) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

// GetRowFilter returns the value of the RowFilter field in TableInfo as
// a TableRowFilter value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo) GetRowFilter(ctx context.Context) (TableRowFilter, bool) {
	var e TableRowFilter
	if o.RowFilter.IsNull() || o.RowFilter.IsUnknown() {
		return e, false
	}
	var v []TableRowFilter
	d := o.RowFilter.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetRowFilter sets the value of the RowFilter field in TableInfo.
func (o *TableInfo) SetRowFilter(ctx context.Context, v TableRowFilter) {
	vs := v.ToObjectValue(ctx)
	o.RowFilter = vs
}

// GetTableConstraints returns the value of the TableConstraints field in TableInfo as
// a slice of TableConstraint values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo) GetTableConstraints(ctx context.Context) ([]TableConstraint, bool) {
	if o.TableConstraints.IsNull() || o.TableConstraints.IsUnknown() {
		return nil, false
	}
	var v []TableConstraint
	d := o.TableConstraints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableConstraints sets the value of the TableConstraints field in TableInfo.
func (o *TableInfo) SetTableConstraints(ctx context.Context, v []TableConstraint) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table_constraints"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TableConstraints = types.ListValueMust(t, vs)
}

// GetViewDependencies returns the value of the ViewDependencies field in TableInfo as
// a DependencyList value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo) GetViewDependencies(ctx context.Context) (DependencyList, bool) {
	var e DependencyList
	if o.ViewDependencies.IsNull() || o.ViewDependencies.IsUnknown() {
		return e, false
	}
	var v []DependencyList
	d := o.ViewDependencies.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetViewDependencies sets the value of the ViewDependencies field in TableInfo.
func (o *TableInfo) SetViewDependencies(ctx context.Context, v DependencyList) {
	vs := v.ToObjectValue(ctx)
	o.ViewDependencies = vs
}

type TableRowFilter struct {
	// The full name of the row filter SQL UDF.
	FunctionName types.String `tfsdk:"function_name"`
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	InputColumnNames types.List `tfsdk:"input_column_names"`
}

func (newState *TableRowFilter) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableRowFilter) {
}

func (newState *TableRowFilter) SyncEffectiveFieldsDuringRead(existingState TableRowFilter) {
}

func (c TableRowFilter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["function_name"] = attrs["function_name"].SetRequired()
	attrs["input_column_names"] = attrs["input_column_names"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableRowFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableRowFilter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input_column_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableRowFilter
// only implements ToObjectValue() and Type().
func (o TableRowFilter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_name":      o.FunctionName,
			"input_column_names": o.InputColumnNames,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableRowFilter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_name": types.StringType,
			"input_column_names": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetInputColumnNames returns the value of the InputColumnNames field in TableRowFilter as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableRowFilter) GetInputColumnNames(ctx context.Context) ([]types.String, bool) {
	if o.InputColumnNames.IsNull() || o.InputColumnNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InputColumnNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInputColumnNames sets the value of the InputColumnNames field in TableRowFilter.
func (o *TableRowFilter) SetInputColumnNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["input_column_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InputColumnNames = types.ListValueMust(t, vs)
}

type TableSummary struct {
	// The full name of the table.
	FullName types.String `tfsdk:"full_name"`

	TableType types.String `tfsdk:"table_type"`
}

func (newState *TableSummary) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableSummary) {
}

func (newState *TableSummary) SyncEffectiveFieldsDuringRead(existingState TableSummary) {
}

func (c TableSummary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["table_type"] = attrs["table_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableSummary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableSummary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSummary
// only implements ToObjectValue() and Type().
func (o TableSummary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":  o.FullName,
			"table_type": o.TableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableSummary) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":  types.StringType,
			"table_type": types.StringType,
		},
	}
}

type TagKeyValue struct {
	// name of the tag
	Key types.String `tfsdk:"key"`
	// value of the tag associated with the key, could be optional
	Value types.String `tfsdk:"value"`
}

func (newState *TagKeyValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan TagKeyValue) {
}

func (newState *TagKeyValue) SyncEffectiveFieldsDuringRead(existingState TagKeyValue) {
}

func (c TagKeyValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TagKeyValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TagKeyValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagKeyValue
// only implements ToObjectValue() and Type().
func (o TagKeyValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TagKeyValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type TemporaryCredentials struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	AwsTempCredentials types.Object `tfsdk:"aws_temp_credentials"`
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	AzureAad types.Object `tfsdk:"azure_aad"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime types.Int64 `tfsdk:"expiration_time"`
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	GcpOauthToken types.Object `tfsdk:"gcp_oauth_token"`
}

func (newState *TemporaryCredentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan TemporaryCredentials) {
}

func (newState *TemporaryCredentials) SyncEffectiveFieldsDuringRead(existingState TemporaryCredentials) {
}

func (c TemporaryCredentials) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_temp_credentials"] = attrs["aws_temp_credentials"].SetOptional()
	attrs["azure_aad"] = attrs["azure_aad"].SetOptional()
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["gcp_oauth_token"] = attrs["gcp_oauth_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TemporaryCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TemporaryCredentials) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_temp_credentials": reflect.TypeOf(AwsCredentials{}),
		"azure_aad":            reflect.TypeOf(AzureActiveDirectoryToken{}),
		"gcp_oauth_token":      reflect.TypeOf(GcpOauthToken{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TemporaryCredentials
// only implements ToObjectValue() and Type().
func (o TemporaryCredentials) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_temp_credentials": o.AwsTempCredentials,
			"azure_aad":            o.AzureAad,
			"expiration_time":      o.ExpirationTime,
			"gcp_oauth_token":      o.GcpOauthToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TemporaryCredentials) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_temp_credentials": AwsCredentials{}.Type(ctx),
			"azure_aad":            AzureActiveDirectoryToken{}.Type(ctx),
			"expiration_time":      types.Int64Type,
			"gcp_oauth_token":      GcpOauthToken{}.Type(ctx),
		},
	}
}

// GetAwsTempCredentials returns the value of the AwsTempCredentials field in TemporaryCredentials as
// a AwsCredentials value.
// If the field is unknown or null, the boolean return value is false.
func (o *TemporaryCredentials) GetAwsTempCredentials(ctx context.Context) (AwsCredentials, bool) {
	var e AwsCredentials
	if o.AwsTempCredentials.IsNull() || o.AwsTempCredentials.IsUnknown() {
		return e, false
	}
	var v []AwsCredentials
	d := o.AwsTempCredentials.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsTempCredentials sets the value of the AwsTempCredentials field in TemporaryCredentials.
func (o *TemporaryCredentials) SetAwsTempCredentials(ctx context.Context, v AwsCredentials) {
	vs := v.ToObjectValue(ctx)
	o.AwsTempCredentials = vs
}

// GetAzureAad returns the value of the AzureAad field in TemporaryCredentials as
// a AzureActiveDirectoryToken value.
// If the field is unknown or null, the boolean return value is false.
func (o *TemporaryCredentials) GetAzureAad(ctx context.Context) (AzureActiveDirectoryToken, bool) {
	var e AzureActiveDirectoryToken
	if o.AzureAad.IsNull() || o.AzureAad.IsUnknown() {
		return e, false
	}
	var v []AzureActiveDirectoryToken
	d := o.AzureAad.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureAad sets the value of the AzureAad field in TemporaryCredentials.
func (o *TemporaryCredentials) SetAzureAad(ctx context.Context, v AzureActiveDirectoryToken) {
	vs := v.ToObjectValue(ctx)
	o.AzureAad = vs
}

// GetGcpOauthToken returns the value of the GcpOauthToken field in TemporaryCredentials as
// a GcpOauthToken value.
// If the field is unknown or null, the boolean return value is false.
func (o *TemporaryCredentials) GetGcpOauthToken(ctx context.Context) (GcpOauthToken, bool) {
	var e GcpOauthToken
	if o.GcpOauthToken.IsNull() || o.GcpOauthToken.IsUnknown() {
		return e, false
	}
	var v []GcpOauthToken
	d := o.GcpOauthToken.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetGcpOauthToken sets the value of the GcpOauthToken field in TemporaryCredentials.
func (o *TemporaryCredentials) SetGcpOauthToken(ctx context.Context, v GcpOauthToken) {
	vs := v.ToObjectValue(ctx)
	o.GcpOauthToken = vs
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
type TriggeredUpdateStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp types.String `tfsdk:"timestamp"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress types.Object `tfsdk:"triggered_update_progress"`
}

func (newState *TriggeredUpdateStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan TriggeredUpdateStatus) {
}

func (newState *TriggeredUpdateStatus) SyncEffectiveFieldsDuringRead(existingState TriggeredUpdateStatus) {
}

func (c TriggeredUpdateStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetOptional()
	attrs["triggered_update_progress"] = attrs["triggered_update_progress"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TriggeredUpdateStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TriggeredUpdateStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"triggered_update_progress": reflect.TypeOf(PipelineProgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TriggeredUpdateStatus
// only implements ToObjectValue() and Type().
func (o TriggeredUpdateStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": o.LastProcessedCommitVersion,
			"timestamp":                     o.Timestamp,
			"triggered_update_progress":     o.TriggeredUpdateProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TriggeredUpdateStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
			"triggered_update_progress":     PipelineProgress{}.Type(ctx),
		},
	}
}

// GetTriggeredUpdateProgress returns the value of the TriggeredUpdateProgress field in TriggeredUpdateStatus as
// a PipelineProgress value.
// If the field is unknown or null, the boolean return value is false.
func (o *TriggeredUpdateStatus) GetTriggeredUpdateProgress(ctx context.Context) (PipelineProgress, bool) {
	var e PipelineProgress
	if o.TriggeredUpdateProgress.IsNull() || o.TriggeredUpdateProgress.IsUnknown() {
		return e, false
	}
	var v []PipelineProgress
	d := o.TriggeredUpdateProgress.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTriggeredUpdateProgress sets the value of the TriggeredUpdateProgress field in TriggeredUpdateStatus.
func (o *TriggeredUpdateStatus) SetTriggeredUpdateProgress(ctx context.Context, v PipelineProgress) {
	vs := v.ToObjectValue(ctx)
	o.TriggeredUpdateProgress = vs
}

// Delete an assignment
type UnassignRequest struct {
	// Query for the ID of the metastore to delete.
	MetastoreId types.String `tfsdk:"-"`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnassignRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnassignRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnassignRequest
// only implements ToObjectValue() and Type().
func (o UnassignRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UnassignRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

type UnassignResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnassignResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnassignResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnassignResponse
// only implements ToObjectValue() and Type().
func (o UnassignResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UnassignResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateAssignmentResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAssignmentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAssignmentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAssignmentResponse
// only implements ToObjectValue() and Type().
func (o UpdateAssignmentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAssignmentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateCatalog struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
	// New name for the catalog.
	NewName types.String `tfsdk:"new_name"`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options"`
	// Username of current owner of catalog.
	Owner types.String `tfsdk:"owner"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`
}

func (newState *UpdateCatalog) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCatalog) {
}

func (newState *UpdateCatalog) SyncEffectiveFieldsDuringRead(existingState UpdateCatalog) {
}

func (c UpdateCatalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["enable_predictive_optimization"] = attrs["enable_predictive_optimization"].SetOptional()
	attrs["isolation_mode"] = attrs["isolation_mode"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["options"] = attrs["options"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties"] = attrs["properties"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCatalog.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCatalog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options":    reflect.TypeOf(types.String{}),
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCatalog
// only implements ToObjectValue() and Type().
func (o UpdateCatalog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":                        o.Comment,
			"enable_predictive_optimization": o.EnablePredictiveOptimization,
			"isolation_mode":                 o.IsolationMode,
			"name":                           o.Name,
			"new_name":                       o.NewName,
			"options":                        o.Options,
			"owner":                          o.Owner,
			"properties":                     o.Properties,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCatalog) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":                        types.StringType,
			"enable_predictive_optimization": types.StringType,
			"isolation_mode":                 types.StringType,
			"name":                           types.StringType,
			"new_name":                       types.StringType,
			"options": basetypes.MapType{
				ElemType: types.StringType,
			},
			"owner": types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetOptions returns the value of the Options field in UpdateCatalog as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCatalog) GetOptions(ctx context.Context) (map[string]types.String, bool) {
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in UpdateCatalog.
func (o *UpdateCatalog) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in UpdateCatalog as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCatalog) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in UpdateCatalog.
func (o *UpdateCatalog) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type UpdateCatalogWorkspaceBindingsResponse struct {
	// A list of workspace IDs
	Workspaces types.List `tfsdk:"workspaces"`
}

func (newState *UpdateCatalogWorkspaceBindingsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCatalogWorkspaceBindingsResponse) {
}

func (newState *UpdateCatalogWorkspaceBindingsResponse) SyncEffectiveFieldsDuringRead(existingState UpdateCatalogWorkspaceBindingsResponse) {
}

func (c UpdateCatalogWorkspaceBindingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspaces"] = attrs["workspaces"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCatalogWorkspaceBindingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCatalogWorkspaceBindingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspaces": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCatalogWorkspaceBindingsResponse
// only implements ToObjectValue() and Type().
func (o UpdateCatalogWorkspaceBindingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspaces": o.Workspaces,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCatalogWorkspaceBindingsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspaces": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetWorkspaces returns the value of the Workspaces field in UpdateCatalogWorkspaceBindingsResponse as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCatalogWorkspaceBindingsResponse) GetWorkspaces(ctx context.Context) ([]types.Int64, bool) {
	if o.Workspaces.IsNull() || o.Workspaces.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.Workspaces.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaces sets the value of the Workspaces field in UpdateCatalogWorkspaceBindingsResponse.
func (o *UpdateCatalogWorkspaceBindingsResponse) SetWorkspaces(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspaces"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Workspaces = types.ListValueMust(t, vs)
}

type UpdateConnection struct {
	// Name of the connection.
	Name types.String `tfsdk:"-"`
	// New name for the connection.
	NewName types.String `tfsdk:"new_name"`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options"`
	// Username of current owner of the connection.
	Owner types.String `tfsdk:"owner"`
}

func (newState *UpdateConnection) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateConnection) {
}

func (newState *UpdateConnection) SyncEffectiveFieldsDuringRead(existingState UpdateConnection) {
}

func (c UpdateConnection) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["options"] = attrs["options"].SetRequired()
	attrs["owner"] = attrs["owner"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateConnection.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateConnection) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateConnection
// only implements ToObjectValue() and Type().
func (o UpdateConnection) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     o.Name,
			"new_name": o.NewName,
			"options":  o.Options,
			"owner":    o.Owner,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateConnection) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":     types.StringType,
			"new_name": types.StringType,
			"options": basetypes.MapType{
				ElemType: types.StringType,
			},
			"owner": types.StringType,
		},
	}
}

// GetOptions returns the value of the Options field in UpdateConnection as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateConnection) GetOptions(ctx context.Context) (map[string]types.String, bool) {
	if o.Options.IsNull() || o.Options.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.Options.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOptions sets the value of the Options field in UpdateConnection.
func (o *UpdateConnection) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

type UpdateCredentialRequest struct {
	// The AWS IAM role configuration
	AwsIamRole types.Object `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.Object `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal types.Object `tfsdk:"azure_service_principal"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount types.Object `tfsdk:"databricks_gcp_service_account"`
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force types.Bool `tfsdk:"force"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode"`
	// Name of the credential.
	NameArg types.String `tfsdk:"-"`
	// New name of credential.
	NewName types.String `tfsdk:"new_name"`
	// Username of current owner of credential.
	Owner types.String `tfsdk:"owner"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// Supply true to this argument to skip validation of the updated
	// credential.
	SkipValidation types.Bool `tfsdk:"skip_validation"`
}

func (newState *UpdateCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCredentialRequest) {
}

func (newState *UpdateCredentialRequest) SyncEffectiveFieldsDuringRead(existingState UpdateCredentialRequest) {
}

func (c UpdateCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_iam_role"] = attrs["aws_iam_role"].SetOptional()
	attrs["azure_managed_identity"] = attrs["azure_managed_identity"].SetOptional()
	attrs["azure_service_principal"] = attrs["azure_service_principal"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["databricks_gcp_service_account"] = attrs["databricks_gcp_service_account"].SetOptional()
	attrs["force"] = attrs["force"].SetOptional()
	attrs["isolation_mode"] = attrs["isolation_mode"].SetOptional()
	attrs["name_arg"] = attrs["name_arg"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["skip_validation"] = attrs["skip_validation"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRole{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentity{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCredentialRequest
// only implements ToObjectValue() and Type().
func (o UpdateCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_iam_role":                   o.AwsIamRole,
			"azure_managed_identity":         o.AzureManagedIdentity,
			"azure_service_principal":        o.AzureServicePrincipal,
			"comment":                        o.Comment,
			"databricks_gcp_service_account": o.DatabricksGcpServiceAccount,
			"force":                          o.Force,
			"isolation_mode":                 o.IsolationMode,
			"name_arg":                       o.NameArg,
			"new_name":                       o.NewName,
			"owner":                          o.Owner,
			"read_only":                      o.ReadOnly,
			"skip_validation":                o.SkipValidation,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role":                   AwsIamRole{}.Type(ctx),
			"azure_managed_identity":         AzureManagedIdentity{}.Type(ctx),
			"azure_service_principal":        AzureServicePrincipal{}.Type(ctx),
			"comment":                        types.StringType,
			"databricks_gcp_service_account": DatabricksGcpServiceAccount{}.Type(ctx),
			"force":                          types.BoolType,
			"isolation_mode":                 types.StringType,
			"name_arg":                       types.StringType,
			"new_name":                       types.StringType,
			"owner":                          types.StringType,
			"read_only":                      types.BoolType,
			"skip_validation":                types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in UpdateCredentialRequest as
// a AwsIamRole value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCredentialRequest) GetAwsIamRole(ctx context.Context) (AwsIamRole, bool) {
	var e AwsIamRole
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRole
	d := o.AwsIamRole.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsIamRole sets the value of the AwsIamRole field in UpdateCredentialRequest.
func (o *UpdateCredentialRequest) SetAwsIamRole(ctx context.Context, v AwsIamRole) {
	vs := v.ToObjectValue(ctx)
	o.AwsIamRole = vs
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in UpdateCredentialRequest as
// a AzureManagedIdentity value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCredentialRequest) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentity, bool) {
	var e AzureManagedIdentity
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentity
	d := o.AzureManagedIdentity.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in UpdateCredentialRequest.
func (o *UpdateCredentialRequest) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentity) {
	vs := v.ToObjectValue(ctx)
	o.AzureManagedIdentity = vs
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in UpdateCredentialRequest as
// a AzureServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCredentialRequest) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal, bool) {
	var e AzureServicePrincipal
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal
	d := o.AzureServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in UpdateCredentialRequest.
func (o *UpdateCredentialRequest) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	o.AzureServicePrincipal = vs
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in UpdateCredentialRequest as
// a DatabricksGcpServiceAccount value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCredentialRequest) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccount, bool) {
	var e DatabricksGcpServiceAccount
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccount
	d := o.DatabricksGcpServiceAccount.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in UpdateCredentialRequest.
func (o *UpdateCredentialRequest) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccount) {
	vs := v.ToObjectValue(ctx)
	o.DatabricksGcpServiceAccount = vs
}

// Update a Database Instance
type UpdateDatabaseInstanceRequest struct {
	// A DatabaseInstance represents a logical Postgres instance, comprised of
	// both compute and storage.
	DatabaseInstance types.Object `tfsdk:"database_instance"`
	// The name of the instance. This is the unique identifier for the instance.
	Name types.String `tfsdk:"-"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance": reflect.TypeOf(DatabaseInstance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (o UpdateDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance": o.DatabaseInstance,
			"name":              o.Name,
			"update_mask":       o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance": DatabaseInstance{}.Type(ctx),
			"name":              types.StringType,
			"update_mask":       types.StringType,
		},
	}
}

// GetDatabaseInstance returns the value of the DatabaseInstance field in UpdateDatabaseInstanceRequest as
// a DatabaseInstance value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateDatabaseInstanceRequest) GetDatabaseInstance(ctx context.Context) (DatabaseInstance, bool) {
	var e DatabaseInstance
	if o.DatabaseInstance.IsNull() || o.DatabaseInstance.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstance
	d := o.DatabaseInstance.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDatabaseInstance sets the value of the DatabaseInstance field in UpdateDatabaseInstanceRequest.
func (o *UpdateDatabaseInstanceRequest) SetDatabaseInstance(ctx context.Context, v DatabaseInstance) {
	vs := v.ToObjectValue(ctx)
	o.DatabaseInstance = vs
}

type UpdateExternalLocation struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Name of the storage credential used with this location.
	CredentialName types.String `tfsdk:"credential_name"`
	// [Create:OPT Update:OPT] Whether to enable file events on this external
	// location.
	EnableFileEvents types.Bool `tfsdk:"enable_file_events"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails types.Object `tfsdk:"encryption_details"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback types.Bool `tfsdk:"fallback"`
	// [Create:OPT Update:OPT] File event queue settings.
	FileEventQueue types.Object `tfsdk:"file_event_queue"`
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	Force types.Bool `tfsdk:"force"`

	IsolationMode types.String `tfsdk:"isolation_mode"`
	// Name of the external location.
	Name types.String `tfsdk:"-"`
	// New name for the external location.
	NewName types.String `tfsdk:"new_name"`
	// The owner of the external location.
	Owner types.String `tfsdk:"owner"`
	// Indicates whether the external location is read-only.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation types.Bool `tfsdk:"skip_validation"`
	// Path URL of the external location.
	Url types.String `tfsdk:"url"`
}

func (newState *UpdateExternalLocation) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExternalLocation) {
}

func (newState *UpdateExternalLocation) SyncEffectiveFieldsDuringRead(existingState UpdateExternalLocation) {
}

func (c UpdateExternalLocation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["credential_name"] = attrs["credential_name"].SetOptional()
	attrs["enable_file_events"] = attrs["enable_file_events"].SetOptional()
	attrs["encryption_details"] = attrs["encryption_details"].SetOptional()
	attrs["fallback"] = attrs["fallback"].SetOptional()
	attrs["file_event_queue"] = attrs["file_event_queue"].SetOptional()
	attrs["force"] = attrs["force"].SetOptional()
	attrs["isolation_mode"] = attrs["isolation_mode"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["skip_validation"] = attrs["skip_validation"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExternalLocation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExternalLocation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"encryption_details": reflect.TypeOf(EncryptionDetails{}),
		"file_event_queue":   reflect.TypeOf(FileEventQueue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExternalLocation
// only implements ToObjectValue() and Type().
func (o UpdateExternalLocation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":            o.Comment,
			"credential_name":    o.CredentialName,
			"enable_file_events": o.EnableFileEvents,
			"encryption_details": o.EncryptionDetails,
			"fallback":           o.Fallback,
			"file_event_queue":   o.FileEventQueue,
			"force":              o.Force,
			"isolation_mode":     o.IsolationMode,
			"name":               o.Name,
			"new_name":           o.NewName,
			"owner":              o.Owner,
			"read_only":          o.ReadOnly,
			"skip_validation":    o.SkipValidation,
			"url":                o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExternalLocation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":            types.StringType,
			"credential_name":    types.StringType,
			"enable_file_events": types.BoolType,
			"encryption_details": EncryptionDetails{}.Type(ctx),
			"fallback":           types.BoolType,
			"file_event_queue":   FileEventQueue{}.Type(ctx),
			"force":              types.BoolType,
			"isolation_mode":     types.StringType,
			"name":               types.StringType,
			"new_name":           types.StringType,
			"owner":              types.StringType,
			"read_only":          types.BoolType,
			"skip_validation":    types.BoolType,
			"url":                types.StringType,
		},
	}
}

// GetEncryptionDetails returns the value of the EncryptionDetails field in UpdateExternalLocation as
// a EncryptionDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExternalLocation) GetEncryptionDetails(ctx context.Context) (EncryptionDetails, bool) {
	var e EncryptionDetails
	if o.EncryptionDetails.IsNull() || o.EncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []EncryptionDetails
	d := o.EncryptionDetails.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEncryptionDetails sets the value of the EncryptionDetails field in UpdateExternalLocation.
func (o *UpdateExternalLocation) SetEncryptionDetails(ctx context.Context, v EncryptionDetails) {
	vs := v.ToObjectValue(ctx)
	o.EncryptionDetails = vs
}

// GetFileEventQueue returns the value of the FileEventQueue field in UpdateExternalLocation as
// a FileEventQueue value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExternalLocation) GetFileEventQueue(ctx context.Context) (FileEventQueue, bool) {
	var e FileEventQueue
	if o.FileEventQueue.IsNull() || o.FileEventQueue.IsUnknown() {
		return e, false
	}
	var v []FileEventQueue
	d := o.FileEventQueue.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetFileEventQueue sets the value of the FileEventQueue field in UpdateExternalLocation.
func (o *UpdateExternalLocation) SetFileEventQueue(ctx context.Context, v FileEventQueue) {
	vs := v.ToObjectValue(ctx)
	o.FileEventQueue = vs
}

type UpdateFunction struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name types.String `tfsdk:"-"`
	// Username of current owner of function.
	Owner types.String `tfsdk:"owner"`
}

func (newState *UpdateFunction) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateFunction) {
}

func (newState *UpdateFunction) SyncEffectiveFieldsDuringRead(existingState UpdateFunction) {
}

func (c UpdateFunction) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["owner"] = attrs["owner"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateFunction.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateFunction) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateFunction
// only implements ToObjectValue() and Type().
func (o UpdateFunction) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":  o.Name,
			"owner": o.Owner,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateFunction) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":  types.StringType,
			"owner": types.StringType,
		},
	}
}

type UpdateMetastore struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName types.String `tfsdk:"delta_sharing_organization_name"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds types.Int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope types.String `tfsdk:"delta_sharing_scope"`
	// Unique ID of the metastore.
	Id types.String `tfsdk:"-"`
	// New name for the metastore.
	NewName types.String `tfsdk:"new_name"`
	// The owner of the metastore.
	Owner types.String `tfsdk:"owner"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion types.String `tfsdk:"privilege_model_version"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId types.String `tfsdk:"storage_root_credential_id"`
}

func (newState *UpdateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMetastore) {
}

func (newState *UpdateMetastore) SyncEffectiveFieldsDuringRead(existingState UpdateMetastore) {
}

func (c UpdateMetastore) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["delta_sharing_organization_name"] = attrs["delta_sharing_organization_name"].SetOptional()
	attrs["delta_sharing_recipient_token_lifetime_in_seconds"] = attrs["delta_sharing_recipient_token_lifetime_in_seconds"].SetOptional()
	attrs["delta_sharing_scope"] = attrs["delta_sharing_scope"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["privilege_model_version"] = attrs["privilege_model_version"].SetOptional()
	attrs["storage_root_credential_id"] = attrs["storage_root_credential_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateMetastore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateMetastore) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateMetastore
// only implements ToObjectValue() and Type().
func (o UpdateMetastore) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_sharing_organization_name":                   o.DeltaSharingOrganizationName,
			"delta_sharing_recipient_token_lifetime_in_seconds": o.DeltaSharingRecipientTokenLifetimeInSeconds,
			"delta_sharing_scope":                               o.DeltaSharingScope,
			"id":                                                o.Id,
			"new_name":                                          o.NewName,
			"owner":                                             o.Owner,
			"privilege_model_version":                           o.PrivilegeModelVersion,
			"storage_root_credential_id":                        o.StorageRootCredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateMetastore) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_sharing_organization_name":                   types.StringType,
			"delta_sharing_recipient_token_lifetime_in_seconds": types.Int64Type,
			"delta_sharing_scope":                               types.StringType,
			"id":                                                types.StringType,
			"new_name":                                          types.StringType,
			"owner":                                             types.StringType,
			"privilege_model_version":                           types.StringType,
			"storage_root_credential_id":                        types.StringType,
		},
	}
}

type UpdateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name"`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UpdateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMetastoreAssignment) {
}

func (newState *UpdateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState UpdateMetastoreAssignment) {
}

func (c UpdateMetastoreAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_catalog_name"] = attrs["default_catalog_name"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateMetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateMetastoreAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateMetastoreAssignment
// only implements ToObjectValue() and Type().
func (o UpdateMetastoreAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_catalog_name": o.DefaultCatalogName,
			"metastore_id":         o.MetastoreId,
			"workspace_id":         o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateMetastoreAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_catalog_name": types.StringType,
			"metastore_id":         types.StringType,
			"workspace_id":         types.Int64Type,
		},
	}
}

type UpdateModelVersionRequest struct {
	// The comment attached to the model version
	Comment types.String `tfsdk:"comment"`
	// The three-level (fully qualified) name of the model version
	FullName types.String `tfsdk:"-"`
	// The integer version number of the model version
	Version types.Int64 `tfsdk:"-"`
}

func (newState *UpdateModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelVersionRequest) {
}

func (newState *UpdateModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState UpdateModelVersionRequest) {
}

func (c UpdateModelVersionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelVersionRequest
// only implements ToObjectValue() and Type().
func (o UpdateModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":   o.Comment,
			"full_name": o.FullName,
			"version":   o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelVersionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":   types.StringType,
			"full_name": types.StringType,
			"version":   types.Int64Type,
		},
	}
}

type UpdateMonitor struct {
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName types.String `tfsdk:"baseline_table_name"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics types.List `tfsdk:"custom_metrics"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The data classification config for the monitor.
	DataClassificationConfig types.Object `tfsdk:"data_classification_config"`
	// Configuration for monitoring inference logs.
	InferenceLog types.Object `tfsdk:"inference_log"`
	// The notification settings for the monitor.
	Notifications types.Object `tfsdk:"notifications"`
	// Schema where output metric tables are created.
	OutputSchemaName types.String `tfsdk:"output_schema_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule types.Object `tfsdk:"schedule"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs types.List `tfsdk:"slicing_exprs"`
	// Configuration for monitoring snapshot tables.
	Snapshot types.Object `tfsdk:"snapshot"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries types.Object `tfsdk:"time_series"`
}

func (newState *UpdateMonitor) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMonitor) {
}

func (newState *UpdateMonitor) SyncEffectiveFieldsDuringRead(existingState UpdateMonitor) {
}

func (c UpdateMonitor) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["baseline_table_name"] = attrs["baseline_table_name"].SetOptional()
	attrs["custom_metrics"] = attrs["custom_metrics"].SetOptional()
	attrs["dashboard_id"] = attrs["dashboard_id"].SetOptional()
	attrs["data_classification_config"] = attrs["data_classification_config"].SetOptional()
	attrs["inference_log"] = attrs["inference_log"].SetOptional()
	attrs["notifications"] = attrs["notifications"].SetOptional()
	attrs["output_schema_name"] = attrs["output_schema_name"].SetRequired()
	attrs["schedule"] = attrs["schedule"].SetOptional()
	attrs["slicing_exprs"] = attrs["slicing_exprs"].SetOptional()
	attrs["snapshot"] = attrs["snapshot"].SetOptional()
	attrs["table_name"] = attrs["table_name"].SetRequired()
	attrs["time_series"] = attrs["time_series"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateMonitor.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateMonitor) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_metrics":             reflect.TypeOf(MonitorMetric{}),
		"data_classification_config": reflect.TypeOf(MonitorDataClassificationConfig{}),
		"inference_log":              reflect.TypeOf(MonitorInferenceLog{}),
		"notifications":              reflect.TypeOf(MonitorNotifications{}),
		"schedule":                   reflect.TypeOf(MonitorCronSchedule{}),
		"slicing_exprs":              reflect.TypeOf(types.String{}),
		"snapshot":                   reflect.TypeOf(MonitorSnapshot{}),
		"time_series":                reflect.TypeOf(MonitorTimeSeries{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateMonitor
// only implements ToObjectValue() and Type().
func (o UpdateMonitor) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"baseline_table_name":        o.BaselineTableName,
			"custom_metrics":             o.CustomMetrics,
			"dashboard_id":               o.DashboardId,
			"data_classification_config": o.DataClassificationConfig,
			"inference_log":              o.InferenceLog,
			"notifications":              o.Notifications,
			"output_schema_name":         o.OutputSchemaName,
			"schedule":                   o.Schedule,
			"slicing_exprs":              o.SlicingExprs,
			"snapshot":                   o.Snapshot,
			"table_name":                 o.TableName,
			"time_series":                o.TimeSeries,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateMonitor) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: MonitorMetric{}.Type(ctx),
			},
			"dashboard_id":               types.StringType,
			"data_classification_config": MonitorDataClassificationConfig{}.Type(ctx),
			"inference_log":              MonitorInferenceLog{}.Type(ctx),
			"notifications":              MonitorNotifications{}.Type(ctx),
			"output_schema_name":         types.StringType,
			"schedule":                   MonitorCronSchedule{}.Type(ctx),
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot":    MonitorSnapshot{}.Type(ctx),
			"table_name":  types.StringType,
			"time_series": MonitorTimeSeries{}.Type(ctx),
		},
	}
}

// GetCustomMetrics returns the value of the CustomMetrics field in UpdateMonitor as
// a slice of MonitorMetric values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor) GetCustomMetrics(ctx context.Context) ([]MonitorMetric, bool) {
	if o.CustomMetrics.IsNull() || o.CustomMetrics.IsUnknown() {
		return nil, false
	}
	var v []MonitorMetric
	d := o.CustomMetrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomMetrics sets the value of the CustomMetrics field in UpdateMonitor.
func (o *UpdateMonitor) SetCustomMetrics(ctx context.Context, v []MonitorMetric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomMetrics = types.ListValueMust(t, vs)
}

// GetDataClassificationConfig returns the value of the DataClassificationConfig field in UpdateMonitor as
// a MonitorDataClassificationConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor) GetDataClassificationConfig(ctx context.Context) (MonitorDataClassificationConfig, bool) {
	var e MonitorDataClassificationConfig
	if o.DataClassificationConfig.IsNull() || o.DataClassificationConfig.IsUnknown() {
		return e, false
	}
	var v []MonitorDataClassificationConfig
	d := o.DataClassificationConfig.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDataClassificationConfig sets the value of the DataClassificationConfig field in UpdateMonitor.
func (o *UpdateMonitor) SetDataClassificationConfig(ctx context.Context, v MonitorDataClassificationConfig) {
	vs := v.ToObjectValue(ctx)
	o.DataClassificationConfig = vs
}

// GetInferenceLog returns the value of the InferenceLog field in UpdateMonitor as
// a MonitorInferenceLog value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor) GetInferenceLog(ctx context.Context) (MonitorInferenceLog, bool) {
	var e MonitorInferenceLog
	if o.InferenceLog.IsNull() || o.InferenceLog.IsUnknown() {
		return e, false
	}
	var v []MonitorInferenceLog
	d := o.InferenceLog.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetInferenceLog sets the value of the InferenceLog field in UpdateMonitor.
func (o *UpdateMonitor) SetInferenceLog(ctx context.Context, v MonitorInferenceLog) {
	vs := v.ToObjectValue(ctx)
	o.InferenceLog = vs
}

// GetNotifications returns the value of the Notifications field in UpdateMonitor as
// a MonitorNotifications value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor) GetNotifications(ctx context.Context) (MonitorNotifications, bool) {
	var e MonitorNotifications
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return e, false
	}
	var v []MonitorNotifications
	d := o.Notifications.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNotifications sets the value of the Notifications field in UpdateMonitor.
func (o *UpdateMonitor) SetNotifications(ctx context.Context, v MonitorNotifications) {
	vs := v.ToObjectValue(ctx)
	o.Notifications = vs
}

// GetSchedule returns the value of the Schedule field in UpdateMonitor as
// a MonitorCronSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor) GetSchedule(ctx context.Context) (MonitorCronSchedule, bool) {
	var e MonitorCronSchedule
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []MonitorCronSchedule
	d := o.Schedule.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSchedule sets the value of the Schedule field in UpdateMonitor.
func (o *UpdateMonitor) SetSchedule(ctx context.Context, v MonitorCronSchedule) {
	vs := v.ToObjectValue(ctx)
	o.Schedule = vs
}

// GetSlicingExprs returns the value of the SlicingExprs field in UpdateMonitor as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor) GetSlicingExprs(ctx context.Context) ([]types.String, bool) {
	if o.SlicingExprs.IsNull() || o.SlicingExprs.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SlicingExprs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSlicingExprs sets the value of the SlicingExprs field in UpdateMonitor.
func (o *UpdateMonitor) SetSlicingExprs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["slicing_exprs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SlicingExprs = types.ListValueMust(t, vs)
}

// GetSnapshot returns the value of the Snapshot field in UpdateMonitor as
// a MonitorSnapshot value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor) GetSnapshot(ctx context.Context) (MonitorSnapshot, bool) {
	var e MonitorSnapshot
	if o.Snapshot.IsNull() || o.Snapshot.IsUnknown() {
		return e, false
	}
	var v []MonitorSnapshot
	d := o.Snapshot.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSnapshot sets the value of the Snapshot field in UpdateMonitor.
func (o *UpdateMonitor) SetSnapshot(ctx context.Context, v MonitorSnapshot) {
	vs := v.ToObjectValue(ctx)
	o.Snapshot = vs
}

// GetTimeSeries returns the value of the TimeSeries field in UpdateMonitor as
// a MonitorTimeSeries value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor) GetTimeSeries(ctx context.Context) (MonitorTimeSeries, bool) {
	var e MonitorTimeSeries
	if o.TimeSeries.IsNull() || o.TimeSeries.IsUnknown() {
		return e, false
	}
	var v []MonitorTimeSeries
	d := o.TimeSeries.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTimeSeries sets the value of the TimeSeries field in UpdateMonitor.
func (o *UpdateMonitor) SetTimeSeries(ctx context.Context, v MonitorTimeSeries) {
	vs := v.ToObjectValue(ctx)
	o.TimeSeries = vs
}

type UpdatePermissions struct {
	// Array of permissions change objects.
	Changes types.List `tfsdk:"changes"`
	// Full name of securable.
	FullName types.String `tfsdk:"-"`
	// Type of securable.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *UpdatePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePermissions) {
}

func (newState *UpdatePermissions) SyncEffectiveFieldsDuringRead(existingState UpdatePermissions) {
}

func (c UpdatePermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["changes"] = attrs["changes"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["securable_type"] = attrs["securable_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"changes": reflect.TypeOf(PermissionsChange{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePermissions
// only implements ToObjectValue() and Type().
func (o UpdatePermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"changes":        o.Changes,
			"full_name":      o.FullName,
			"securable_type": o.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"changes": basetypes.ListType{
				ElemType: PermissionsChange{}.Type(ctx),
			},
			"full_name":      types.StringType,
			"securable_type": types.StringType,
		},
	}
}

// GetChanges returns the value of the Changes field in UpdatePermissions as
// a slice of PermissionsChange values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdatePermissions) GetChanges(ctx context.Context) ([]PermissionsChange, bool) {
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

// SetChanges sets the value of the Changes field in UpdatePermissions.
func (o *UpdatePermissions) SetChanges(ctx context.Context, v []PermissionsChange) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["changes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Changes = types.ListValueMust(t, vs)
}

type UpdateRegisteredModelRequest struct {
	// The comment attached to the registered model
	Comment types.String `tfsdk:"comment"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
	// New name for the registered model.
	NewName types.String `tfsdk:"new_name"`
	// The identifier of the user who owns the registered model
	Owner types.String `tfsdk:"owner"`
}

func (newState *UpdateRegisteredModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRegisteredModelRequest) {
}

func (newState *UpdateRegisteredModelRequest) SyncEffectiveFieldsDuringRead(existingState UpdateRegisteredModelRequest) {
}

func (c UpdateRegisteredModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRegisteredModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRegisteredModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRegisteredModelRequest
// only implements ToObjectValue() and Type().
func (o UpdateRegisteredModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":   o.Comment,
			"full_name": o.FullName,
			"new_name":  o.NewName,
			"owner":     o.Owner,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRegisteredModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":   types.StringType,
			"full_name": types.StringType,
			"new_name":  types.StringType,
			"owner":     types.StringType,
		},
	}
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

type UpdateSchema struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization"`
	// Full name of the schema.
	FullName types.String `tfsdk:"-"`
	// New name for the schema.
	NewName types.String `tfsdk:"new_name"`
	// Username of current owner of schema.
	Owner types.String `tfsdk:"owner"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties"`
}

func (newState *UpdateSchema) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSchema) {
}

func (newState *UpdateSchema) SyncEffectiveFieldsDuringRead(existingState UpdateSchema) {
}

func (c UpdateSchema) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["enable_predictive_optimization"] = attrs["enable_predictive_optimization"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["properties"] = attrs["properties"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSchema.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateSchema) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSchema
// only implements ToObjectValue() and Type().
func (o UpdateSchema) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":                        o.Comment,
			"enable_predictive_optimization": o.EnablePredictiveOptimization,
			"full_name":                      o.FullName,
			"new_name":                       o.NewName,
			"owner":                          o.Owner,
			"properties":                     o.Properties,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateSchema) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":                        types.StringType,
			"enable_predictive_optimization": types.StringType,
			"full_name":                      types.StringType,
			"new_name":                       types.StringType,
			"owner":                          types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetProperties returns the value of the Properties field in UpdateSchema as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateSchema) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in UpdateSchema.
func (o *UpdateSchema) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type UpdateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole types.Object `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.Object `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.Object `tfsdk:"azure_service_principal"`
	// The Cloudflare API token configuration.
	CloudflareApiToken types.Object `tfsdk:"cloudflare_api_token"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount types.Object `tfsdk:"databricks_gcp_service_account"`
	// Force update even if there are dependent external locations or external
	// tables.
	Force types.Bool `tfsdk:"force"`

	IsolationMode types.String `tfsdk:"isolation_mode"`
	// Name of the storage credential.
	Name types.String `tfsdk:"-"`
	// New name for the storage credential.
	NewName types.String `tfsdk:"new_name"`
	// Username of current owner of credential.
	Owner types.String `tfsdk:"owner"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// Supplying true to this argument skips validation of the updated
	// credential.
	SkipValidation types.Bool `tfsdk:"skip_validation"`
}

func (newState *UpdateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateStorageCredential) {
}

func (newState *UpdateStorageCredential) SyncEffectiveFieldsDuringRead(existingState UpdateStorageCredential) {
}

func (c UpdateStorageCredential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_iam_role"] = attrs["aws_iam_role"].SetOptional()
	attrs["azure_managed_identity"] = attrs["azure_managed_identity"].SetOptional()
	attrs["azure_service_principal"] = attrs["azure_service_principal"].SetOptional()
	attrs["cloudflare_api_token"] = attrs["cloudflare_api_token"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["databricks_gcp_service_account"] = attrs["databricks_gcp_service_account"].SetOptional()
	attrs["force"] = attrs["force"].SetOptional()
	attrs["isolation_mode"] = attrs["isolation_mode"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["skip_validation"] = attrs["skip_validation"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateStorageCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateStorageCredential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRoleRequest{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentityResponse{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal{}),
		"cloudflare_api_token":           reflect.TypeOf(CloudflareApiToken{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccountRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateStorageCredential
// only implements ToObjectValue() and Type().
func (o UpdateStorageCredential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_iam_role":                   o.AwsIamRole,
			"azure_managed_identity":         o.AzureManagedIdentity,
			"azure_service_principal":        o.AzureServicePrincipal,
			"cloudflare_api_token":           o.CloudflareApiToken,
			"comment":                        o.Comment,
			"databricks_gcp_service_account": o.DatabricksGcpServiceAccount,
			"force":                          o.Force,
			"isolation_mode":                 o.IsolationMode,
			"name":                           o.Name,
			"new_name":                       o.NewName,
			"owner":                          o.Owner,
			"read_only":                      o.ReadOnly,
			"skip_validation":                o.SkipValidation,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateStorageCredential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role":                   AwsIamRoleRequest{}.Type(ctx),
			"azure_managed_identity":         AzureManagedIdentityResponse{}.Type(ctx),
			"azure_service_principal":        AzureServicePrincipal{}.Type(ctx),
			"cloudflare_api_token":           CloudflareApiToken{}.Type(ctx),
			"comment":                        types.StringType,
			"databricks_gcp_service_account": DatabricksGcpServiceAccountRequest{}.Type(ctx),
			"force":                          types.BoolType,
			"isolation_mode":                 types.StringType,
			"name":                           types.StringType,
			"new_name":                       types.StringType,
			"owner":                          types.StringType,
			"read_only":                      types.BoolType,
			"skip_validation":                types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in UpdateStorageCredential as
// a AwsIamRoleRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateStorageCredential) GetAwsIamRole(ctx context.Context) (AwsIamRoleRequest, bool) {
	var e AwsIamRoleRequest
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRoleRequest
	d := o.AwsIamRole.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsIamRole sets the value of the AwsIamRole field in UpdateStorageCredential.
func (o *UpdateStorageCredential) SetAwsIamRole(ctx context.Context, v AwsIamRoleRequest) {
	vs := v.ToObjectValue(ctx)
	o.AwsIamRole = vs
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in UpdateStorageCredential as
// a AzureManagedIdentityResponse value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateStorageCredential) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentityResponse, bool) {
	var e AzureManagedIdentityResponse
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentityResponse
	d := o.AzureManagedIdentity.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in UpdateStorageCredential.
func (o *UpdateStorageCredential) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentityResponse) {
	vs := v.ToObjectValue(ctx)
	o.AzureManagedIdentity = vs
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in UpdateStorageCredential as
// a AzureServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateStorageCredential) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal, bool) {
	var e AzureServicePrincipal
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal
	d := o.AzureServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in UpdateStorageCredential.
func (o *UpdateStorageCredential) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	o.AzureServicePrincipal = vs
}

// GetCloudflareApiToken returns the value of the CloudflareApiToken field in UpdateStorageCredential as
// a CloudflareApiToken value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateStorageCredential) GetCloudflareApiToken(ctx context.Context) (CloudflareApiToken, bool) {
	var e CloudflareApiToken
	if o.CloudflareApiToken.IsNull() || o.CloudflareApiToken.IsUnknown() {
		return e, false
	}
	var v []CloudflareApiToken
	d := o.CloudflareApiToken.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetCloudflareApiToken sets the value of the CloudflareApiToken field in UpdateStorageCredential.
func (o *UpdateStorageCredential) SetCloudflareApiToken(ctx context.Context, v CloudflareApiToken) {
	vs := v.ToObjectValue(ctx)
	o.CloudflareApiToken = vs
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in UpdateStorageCredential as
// a DatabricksGcpServiceAccountRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateStorageCredential) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccountRequest, bool) {
	var e DatabricksGcpServiceAccountRequest
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccountRequest
	d := o.DatabricksGcpServiceAccount.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in UpdateStorageCredential.
func (o *UpdateStorageCredential) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccountRequest) {
	vs := v.ToObjectValue(ctx)
	o.DatabricksGcpServiceAccount = vs
}

// Update a table owner.
type UpdateTableRequest struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`

	Owner types.String `tfsdk:"owner"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateTableRequest
// only implements ToObjectValue() and Type().
func (o UpdateTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": o.FullName,
			"owner":     o.Owner,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
			"owner":     types.StringType,
		},
	}
}

type UpdateVolumeRequestContent struct {
	// The comment attached to the volume
	Comment types.String `tfsdk:"comment"`
	// The three-level (fully qualified) name of the volume
	Name types.String `tfsdk:"-"`
	// New name for the volume.
	NewName types.String `tfsdk:"new_name"`
	// The identifier of the user who owns the volume
	Owner types.String `tfsdk:"owner"`
}

func (newState *UpdateVolumeRequestContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateVolumeRequestContent) {
}

func (newState *UpdateVolumeRequestContent) SyncEffectiveFieldsDuringRead(existingState UpdateVolumeRequestContent) {
}

func (c UpdateVolumeRequestContent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateVolumeRequestContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateVolumeRequestContent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateVolumeRequestContent
// only implements ToObjectValue() and Type().
func (o UpdateVolumeRequestContent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":  o.Comment,
			"name":     o.Name,
			"new_name": o.NewName,
			"owner":    o.Owner,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateVolumeRequestContent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":  types.StringType,
			"name":     types.StringType,
			"new_name": types.StringType,
			"owner":    types.StringType,
		},
	}
}

type UpdateWorkspaceBindings struct {
	// A list of workspace IDs.
	AssignWorkspaces types.List `tfsdk:"assign_workspaces"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
	// A list of workspace IDs.
	UnassignWorkspaces types.List `tfsdk:"unassign_workspaces"`
}

func (newState *UpdateWorkspaceBindings) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceBindings) {
}

func (newState *UpdateWorkspaceBindings) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceBindings) {
}

func (c UpdateWorkspaceBindings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assign_workspaces"] = attrs["assign_workspaces"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["unassign_workspaces"] = attrs["unassign_workspaces"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceBindings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceBindings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assign_workspaces":   reflect.TypeOf(types.Int64{}),
		"unassign_workspaces": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceBindings
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceBindings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assign_workspaces":   o.AssignWorkspaces,
			"name":                o.Name,
			"unassign_workspaces": o.UnassignWorkspaces,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceBindings) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assign_workspaces": basetypes.ListType{
				ElemType: types.Int64Type,
			},
			"name": types.StringType,
			"unassign_workspaces": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetAssignWorkspaces returns the value of the AssignWorkspaces field in UpdateWorkspaceBindings as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceBindings) GetAssignWorkspaces(ctx context.Context) ([]types.Int64, bool) {
	if o.AssignWorkspaces.IsNull() || o.AssignWorkspaces.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.AssignWorkspaces.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssignWorkspaces sets the value of the AssignWorkspaces field in UpdateWorkspaceBindings.
func (o *UpdateWorkspaceBindings) SetAssignWorkspaces(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["assign_workspaces"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AssignWorkspaces = types.ListValueMust(t, vs)
}

// GetUnassignWorkspaces returns the value of the UnassignWorkspaces field in UpdateWorkspaceBindings as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceBindings) GetUnassignWorkspaces(ctx context.Context) ([]types.Int64, bool) {
	if o.UnassignWorkspaces.IsNull() || o.UnassignWorkspaces.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.UnassignWorkspaces.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUnassignWorkspaces sets the value of the UnassignWorkspaces field in UpdateWorkspaceBindings.
func (o *UpdateWorkspaceBindings) SetUnassignWorkspaces(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["unassign_workspaces"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UnassignWorkspaces = types.ListValueMust(t, vs)
}

type UpdateWorkspaceBindingsParameters struct {
	// List of workspace bindings.
	Add types.List `tfsdk:"add"`
	// List of workspace bindings.
	Remove types.List `tfsdk:"remove"`
	// The name of the securable.
	SecurableName types.String `tfsdk:"-"`
	// The type of the securable to bind to a workspace (catalog,
	// storage_credential, credential, or external_location).
	SecurableType types.String `tfsdk:"-"`
}

func (newState *UpdateWorkspaceBindingsParameters) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceBindingsParameters) {
}

func (newState *UpdateWorkspaceBindingsParameters) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceBindingsParameters) {
}

func (c UpdateWorkspaceBindingsParameters) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["add"] = attrs["add"].SetOptional()
	attrs["remove"] = attrs["remove"].SetOptional()
	attrs["securable_name"] = attrs["securable_name"].SetRequired()
	attrs["securable_type"] = attrs["securable_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceBindingsParameters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceBindingsParameters) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"add":    reflect.TypeOf(WorkspaceBinding{}),
		"remove": reflect.TypeOf(WorkspaceBinding{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceBindingsParameters
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceBindingsParameters) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"add":            o.Add,
			"remove":         o.Remove,
			"securable_name": o.SecurableName,
			"securable_type": o.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceBindingsParameters) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"add": basetypes.ListType{
				ElemType: WorkspaceBinding{}.Type(ctx),
			},
			"remove": basetypes.ListType{
				ElemType: WorkspaceBinding{}.Type(ctx),
			},
			"securable_name": types.StringType,
			"securable_type": types.StringType,
		},
	}
}

// GetAdd returns the value of the Add field in UpdateWorkspaceBindingsParameters as
// a slice of WorkspaceBinding values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceBindingsParameters) GetAdd(ctx context.Context) ([]WorkspaceBinding, bool) {
	if o.Add.IsNull() || o.Add.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceBinding
	d := o.Add.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAdd sets the value of the Add field in UpdateWorkspaceBindingsParameters.
func (o *UpdateWorkspaceBindingsParameters) SetAdd(ctx context.Context, v []WorkspaceBinding) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["add"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Add = types.ListValueMust(t, vs)
}

// GetRemove returns the value of the Remove field in UpdateWorkspaceBindingsParameters as
// a slice of WorkspaceBinding values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceBindingsParameters) GetRemove(ctx context.Context) ([]WorkspaceBinding, bool) {
	if o.Remove.IsNull() || o.Remove.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceBinding
	d := o.Remove.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRemove sets the value of the Remove field in UpdateWorkspaceBindingsParameters.
func (o *UpdateWorkspaceBindingsParameters) SetRemove(ctx context.Context, v []WorkspaceBinding) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["remove"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Remove = types.ListValueMust(t, vs)
}

// A list of workspace IDs that are bound to the securable
type UpdateWorkspaceBindingsResponse struct {
	// List of workspace bindings.
	Bindings types.List `tfsdk:"bindings"`
}

func (newState *UpdateWorkspaceBindingsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceBindingsResponse) {
}

func (newState *UpdateWorkspaceBindingsResponse) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceBindingsResponse) {
}

func (c UpdateWorkspaceBindingsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bindings"] = attrs["bindings"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceBindingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceBindingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"bindings": reflect.TypeOf(WorkspaceBinding{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceBindingsResponse
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceBindingsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bindings": o.Bindings,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceBindingsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bindings": basetypes.ListType{
				ElemType: WorkspaceBinding{}.Type(ctx),
			},
		},
	}
}

// GetBindings returns the value of the Bindings field in UpdateWorkspaceBindingsResponse as
// a slice of WorkspaceBinding values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceBindingsResponse) GetBindings(ctx context.Context) ([]WorkspaceBinding, bool) {
	if o.Bindings.IsNull() || o.Bindings.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceBinding
	d := o.Bindings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBindings sets the value of the Bindings field in UpdateWorkspaceBindingsResponse.
func (o *UpdateWorkspaceBindingsResponse) SetBindings(ctx context.Context, v []WorkspaceBinding) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["bindings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Bindings = types.ListValueMust(t, vs)
}

// Next ID: 17
type ValidateCredentialRequest struct {
	// The AWS IAM role configuration
	AwsIamRole types.Object `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.Object `tfsdk:"azure_managed_identity"`
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	CredentialName types.String `tfsdk:"credential_name"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount types.Object `tfsdk:"databricks_gcp_service_account"`
	// The name of an existing external location to validate. Only applicable
	// for storage credentials (purpose is **STORAGE**.)
	ExternalLocationName types.String `tfsdk:"external_location_name"`
	// The purpose of the credential. This should only be used when the
	// credential is specified.
	Purpose types.String `tfsdk:"purpose"`
	// Whether the credential is only usable for read operations. Only
	// applicable for storage credentials (purpose is **STORAGE**.)
	ReadOnly types.Bool `tfsdk:"read_only"`
	// The external location url to validate. Only applicable when purpose is
	// **STORAGE**.
	Url types.String `tfsdk:"url"`
}

func (newState *ValidateCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateCredentialRequest) {
}

func (newState *ValidateCredentialRequest) SyncEffectiveFieldsDuringRead(existingState ValidateCredentialRequest) {
}

func (c ValidateCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_iam_role"] = attrs["aws_iam_role"].SetOptional()
	attrs["azure_managed_identity"] = attrs["azure_managed_identity"].SetOptional()
	attrs["credential_name"] = attrs["credential_name"].SetOptional()
	attrs["databricks_gcp_service_account"] = attrs["databricks_gcp_service_account"].SetOptional()
	attrs["external_location_name"] = attrs["external_location_name"].SetOptional()
	attrs["purpose"] = attrs["purpose"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidateCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ValidateCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRole{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentity{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidateCredentialRequest
// only implements ToObjectValue() and Type().
func (o ValidateCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_iam_role":                   o.AwsIamRole,
			"azure_managed_identity":         o.AzureManagedIdentity,
			"credential_name":                o.CredentialName,
			"databricks_gcp_service_account": o.DatabricksGcpServiceAccount,
			"external_location_name":         o.ExternalLocationName,
			"purpose":                        o.Purpose,
			"read_only":                      o.ReadOnly,
			"url":                            o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ValidateCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role":                   AwsIamRole{}.Type(ctx),
			"azure_managed_identity":         AzureManagedIdentity{}.Type(ctx),
			"credential_name":                types.StringType,
			"databricks_gcp_service_account": DatabricksGcpServiceAccount{}.Type(ctx),
			"external_location_name":         types.StringType,
			"purpose":                        types.StringType,
			"read_only":                      types.BoolType,
			"url":                            types.StringType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in ValidateCredentialRequest as
// a AwsIamRole value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateCredentialRequest) GetAwsIamRole(ctx context.Context) (AwsIamRole, bool) {
	var e AwsIamRole
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRole
	d := o.AwsIamRole.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsIamRole sets the value of the AwsIamRole field in ValidateCredentialRequest.
func (o *ValidateCredentialRequest) SetAwsIamRole(ctx context.Context, v AwsIamRole) {
	vs := v.ToObjectValue(ctx)
	o.AwsIamRole = vs
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in ValidateCredentialRequest as
// a AzureManagedIdentity value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateCredentialRequest) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentity, bool) {
	var e AzureManagedIdentity
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentity
	d := o.AzureManagedIdentity.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in ValidateCredentialRequest.
func (o *ValidateCredentialRequest) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentity) {
	vs := v.ToObjectValue(ctx)
	o.AzureManagedIdentity = vs
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in ValidateCredentialRequest as
// a DatabricksGcpServiceAccount value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateCredentialRequest) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccount, bool) {
	var e DatabricksGcpServiceAccount
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccount
	d := o.DatabricksGcpServiceAccount.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in ValidateCredentialRequest.
func (o *ValidateCredentialRequest) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccount) {
	vs := v.ToObjectValue(ctx)
	o.DatabricksGcpServiceAccount = vs
}

type ValidateCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage. Only
	// applicable for when purpose is **STORAGE**.
	IsDir types.Bool `tfsdk:"isDir"`
	// The results of the validation check.
	Results types.List `tfsdk:"results"`
}

func (newState *ValidateCredentialResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateCredentialResponse) {
}

func (newState *ValidateCredentialResponse) SyncEffectiveFieldsDuringRead(existingState ValidateCredentialResponse) {
}

func (c ValidateCredentialResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["isDir"] = attrs["isDir"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidateCredentialResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ValidateCredentialResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(CredentialValidationResult{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidateCredentialResponse
// only implements ToObjectValue() and Type().
func (o ValidateCredentialResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"isDir":   o.IsDir,
			"results": o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ValidateCredentialResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"isDir": types.BoolType,
			"results": basetypes.ListType{
				ElemType: CredentialValidationResult{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ValidateCredentialResponse as
// a slice of CredentialValidationResult values.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateCredentialResponse) GetResults(ctx context.Context) ([]CredentialValidationResult, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []CredentialValidationResult
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ValidateCredentialResponse.
func (o *ValidateCredentialResponse) SetResults(ctx context.Context, v []CredentialValidationResult) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ValidateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole types.Object `tfsdk:"aws_iam_role"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.Object `tfsdk:"azure_managed_identity"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.Object `tfsdk:"azure_service_principal"`
	// The Cloudflare API token configuration.
	CloudflareApiToken types.Object `tfsdk:"cloudflare_api_token"`
	// The Databricks created GCP service account configuration.
	DatabricksGcpServiceAccount types.Object `tfsdk:"databricks_gcp_service_account"`
	// The name of an existing external location to validate.
	ExternalLocationName types.String `tfsdk:"external_location_name"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only"`
	// The name of the storage credential to validate.
	StorageCredentialName types.String `tfsdk:"storage_credential_name"`
	// The external location url to validate.
	Url types.String `tfsdk:"url"`
}

func (newState *ValidateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateStorageCredential) {
}

func (newState *ValidateStorageCredential) SyncEffectiveFieldsDuringRead(existingState ValidateStorageCredential) {
}

func (c ValidateStorageCredential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_iam_role"] = attrs["aws_iam_role"].SetOptional()
	attrs["azure_managed_identity"] = attrs["azure_managed_identity"].SetOptional()
	attrs["azure_service_principal"] = attrs["azure_service_principal"].SetOptional()
	attrs["cloudflare_api_token"] = attrs["cloudflare_api_token"].SetOptional()
	attrs["databricks_gcp_service_account"] = attrs["databricks_gcp_service_account"].SetOptional()
	attrs["external_location_name"] = attrs["external_location_name"].SetOptional()
	attrs["read_only"] = attrs["read_only"].SetOptional()
	attrs["storage_credential_name"] = attrs["storage_credential_name"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidateStorageCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ValidateStorageCredential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRoleRequest{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentityRequest{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal{}),
		"cloudflare_api_token":           reflect.TypeOf(CloudflareApiToken{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccountRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidateStorageCredential
// only implements ToObjectValue() and Type().
func (o ValidateStorageCredential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_iam_role":                   o.AwsIamRole,
			"azure_managed_identity":         o.AzureManagedIdentity,
			"azure_service_principal":        o.AzureServicePrincipal,
			"cloudflare_api_token":           o.CloudflareApiToken,
			"databricks_gcp_service_account": o.DatabricksGcpServiceAccount,
			"external_location_name":         o.ExternalLocationName,
			"read_only":                      o.ReadOnly,
			"storage_credential_name":        o.StorageCredentialName,
			"url":                            o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ValidateStorageCredential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role":                   AwsIamRoleRequest{}.Type(ctx),
			"azure_managed_identity":         AzureManagedIdentityRequest{}.Type(ctx),
			"azure_service_principal":        AzureServicePrincipal{}.Type(ctx),
			"cloudflare_api_token":           CloudflareApiToken{}.Type(ctx),
			"databricks_gcp_service_account": DatabricksGcpServiceAccountRequest{}.Type(ctx),
			"external_location_name":         types.StringType,
			"read_only":                      types.BoolType,
			"storage_credential_name":        types.StringType,
			"url":                            types.StringType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in ValidateStorageCredential as
// a AwsIamRoleRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredential) GetAwsIamRole(ctx context.Context) (AwsIamRoleRequest, bool) {
	var e AwsIamRoleRequest
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRoleRequest
	d := o.AwsIamRole.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsIamRole sets the value of the AwsIamRole field in ValidateStorageCredential.
func (o *ValidateStorageCredential) SetAwsIamRole(ctx context.Context, v AwsIamRoleRequest) {
	vs := v.ToObjectValue(ctx)
	o.AwsIamRole = vs
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in ValidateStorageCredential as
// a AzureManagedIdentityRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredential) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentityRequest, bool) {
	var e AzureManagedIdentityRequest
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentityRequest
	d := o.AzureManagedIdentity.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in ValidateStorageCredential.
func (o *ValidateStorageCredential) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentityRequest) {
	vs := v.ToObjectValue(ctx)
	o.AzureManagedIdentity = vs
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in ValidateStorageCredential as
// a AzureServicePrincipal value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredential) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal, bool) {
	var e AzureServicePrincipal
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal
	d := o.AzureServicePrincipal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in ValidateStorageCredential.
func (o *ValidateStorageCredential) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal) {
	vs := v.ToObjectValue(ctx)
	o.AzureServicePrincipal = vs
}

// GetCloudflareApiToken returns the value of the CloudflareApiToken field in ValidateStorageCredential as
// a CloudflareApiToken value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredential) GetCloudflareApiToken(ctx context.Context) (CloudflareApiToken, bool) {
	var e CloudflareApiToken
	if o.CloudflareApiToken.IsNull() || o.CloudflareApiToken.IsUnknown() {
		return e, false
	}
	var v []CloudflareApiToken
	d := o.CloudflareApiToken.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetCloudflareApiToken sets the value of the CloudflareApiToken field in ValidateStorageCredential.
func (o *ValidateStorageCredential) SetCloudflareApiToken(ctx context.Context, v CloudflareApiToken) {
	vs := v.ToObjectValue(ctx)
	o.CloudflareApiToken = vs
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in ValidateStorageCredential as
// a DatabricksGcpServiceAccountRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredential) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccountRequest, bool) {
	var e DatabricksGcpServiceAccountRequest
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccountRequest
	d := o.DatabricksGcpServiceAccount.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in ValidateStorageCredential.
func (o *ValidateStorageCredential) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccountRequest) {
	vs := v.ToObjectValue(ctx)
	o.DatabricksGcpServiceAccount = vs
}

type ValidateStorageCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage.
	IsDir types.Bool `tfsdk:"isDir"`
	// The results of the validation check.
	Results types.List `tfsdk:"results"`
}

func (newState *ValidateStorageCredentialResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateStorageCredentialResponse) {
}

func (newState *ValidateStorageCredentialResponse) SyncEffectiveFieldsDuringRead(existingState ValidateStorageCredentialResponse) {
}

func (c ValidateStorageCredentialResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["isDir"] = attrs["isDir"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidateStorageCredentialResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ValidateStorageCredentialResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ValidationResult{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidateStorageCredentialResponse
// only implements ToObjectValue() and Type().
func (o ValidateStorageCredentialResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"isDir":   o.IsDir,
			"results": o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ValidateStorageCredentialResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"isDir": types.BoolType,
			"results": basetypes.ListType{
				ElemType: ValidationResult{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ValidateStorageCredentialResponse as
// a slice of ValidationResult values.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredentialResponse) GetResults(ctx context.Context) ([]ValidationResult, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []ValidationResult
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ValidateStorageCredentialResponse.
func (o *ValidateStorageCredentialResponse) SetResults(ctx context.Context, v []ValidationResult) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message types.String `tfsdk:"message"`
	// The operation tested.
	Operation types.String `tfsdk:"operation"`
	// The results of the tested operation.
	Result types.String `tfsdk:"result"`
}

func (newState *ValidationResult) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidationResult) {
}

func (newState *ValidationResult) SyncEffectiveFieldsDuringRead(existingState ValidationResult) {
}

func (c ValidationResult) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetOptional()
	attrs["operation"] = attrs["operation"].SetOptional()
	attrs["result"] = attrs["result"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidationResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ValidationResult) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidationResult
// only implements ToObjectValue() and Type().
func (o ValidationResult) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message":   o.Message,
			"operation": o.Operation,
			"result":    o.Result,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ValidationResult) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message":   types.StringType,
			"operation": types.StringType,
			"result":    types.StringType,
		},
	}
}

type VolumeInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// The name of the catalog where the schema and the volume are
	CatalogName types.String `tfsdk:"catalog_name"`
	// The comment attached to the volume
	Comment types.String `tfsdk:"comment"`

	CreatedAt types.Int64 `tfsdk:"created_at"`
	// The identifier of the user who created the volume
	CreatedBy types.String `tfsdk:"created_by"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails types.Object `tfsdk:"encryption_details"`
	// The three-level (fully qualified) name of the volume
	FullName types.String `tfsdk:"full_name"`
	// The unique identifier of the metastore
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The name of the volume
	Name types.String `tfsdk:"name"`
	// The identifier of the user who owns the volume
	Owner types.String `tfsdk:"owner"`
	// The name of the schema where the volume is
	SchemaName types.String `tfsdk:"schema_name"`
	// The storage location on the cloud
	StorageLocation types.String `tfsdk:"storage_location"`

	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// The identifier of the user who updated the volume last time
	UpdatedBy types.String `tfsdk:"updated_by"`
	// The unique identifier of the volume
	VolumeId types.String `tfsdk:"volume_id"`
	// The type of the volume. An external volume is located in the specified
	// external location. A managed volume is located in the default location
	// which is specified by the parent schema, or the parent catalog, or the
	// Metastore. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
	VolumeType types.String `tfsdk:"volume_type"`
}

func (newState *VolumeInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan VolumeInfo) {
}

func (newState *VolumeInfo) SyncEffectiveFieldsDuringRead(existingState VolumeInfo) {
}

func (c VolumeInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_point"] = attrs["access_point"].SetOptional()
	attrs["browse_only"] = attrs["browse_only"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["encryption_details"] = attrs["encryption_details"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetOptional()
	attrs["metastore_id"] = attrs["metastore_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["storage_location"] = attrs["storage_location"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()
	attrs["volume_id"] = attrs["volume_id"].SetOptional()
	attrs["volume_type"] = attrs["volume_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VolumeInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a VolumeInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"encryption_details": reflect.TypeOf(EncryptionDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VolumeInfo
// only implements ToObjectValue() and Type().
func (o VolumeInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_point":       o.AccessPoint,
			"browse_only":        o.BrowseOnly,
			"catalog_name":       o.CatalogName,
			"comment":            o.Comment,
			"created_at":         o.CreatedAt,
			"created_by":         o.CreatedBy,
			"encryption_details": o.EncryptionDetails,
			"full_name":          o.FullName,
			"metastore_id":       o.MetastoreId,
			"name":               o.Name,
			"owner":              o.Owner,
			"schema_name":        o.SchemaName,
			"storage_location":   o.StorageLocation,
			"updated_at":         o.UpdatedAt,
			"updated_by":         o.UpdatedBy,
			"volume_id":          o.VolumeId,
			"volume_type":        o.VolumeType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o VolumeInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point":       types.StringType,
			"browse_only":        types.BoolType,
			"catalog_name":       types.StringType,
			"comment":            types.StringType,
			"created_at":         types.Int64Type,
			"created_by":         types.StringType,
			"encryption_details": EncryptionDetails{}.Type(ctx),
			"full_name":          types.StringType,
			"metastore_id":       types.StringType,
			"name":               types.StringType,
			"owner":              types.StringType,
			"schema_name":        types.StringType,
			"storage_location":   types.StringType,
			"updated_at":         types.Int64Type,
			"updated_by":         types.StringType,
			"volume_id":          types.StringType,
			"volume_type":        types.StringType,
		},
	}
}

// GetEncryptionDetails returns the value of the EncryptionDetails field in VolumeInfo as
// a EncryptionDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *VolumeInfo) GetEncryptionDetails(ctx context.Context) (EncryptionDetails, bool) {
	var e EncryptionDetails
	if o.EncryptionDetails.IsNull() || o.EncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []EncryptionDetails
	d := o.EncryptionDetails.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEncryptionDetails sets the value of the EncryptionDetails field in VolumeInfo.
func (o *VolumeInfo) SetEncryptionDetails(ctx context.Context, v EncryptionDetails) {
	vs := v.ToObjectValue(ctx)
	o.EncryptionDetails = vs
}

type WorkspaceBinding struct {
	// One of READ_WRITE/READ_ONLY. Default is READ_WRITE.
	BindingType types.String `tfsdk:"binding_type"`
	// Required
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (newState *WorkspaceBinding) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceBinding) {
}

func (newState *WorkspaceBinding) SyncEffectiveFieldsDuringRead(existingState WorkspaceBinding) {
}

func (c WorkspaceBinding) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["binding_type"] = attrs["binding_type"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceBinding.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceBinding) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceBinding
// only implements ToObjectValue() and Type().
func (o WorkspaceBinding) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"binding_type": o.BindingType,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceBinding) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"binding_type": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}
