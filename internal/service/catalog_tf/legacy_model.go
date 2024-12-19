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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AccountsCreateMetastore_SdkV2 struct {
	MetastoreInfo types.List `tfsdk:"metastore_info" tf:"optional,object"`
}

func (newState *AccountsCreateMetastore_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateMetastore_SdkV2) {
}

func (newState *AccountsCreateMetastore_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccountsCreateMetastore_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsCreateMetastore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsCreateMetastore_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_info": reflect.TypeOf(CreateMetastore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsCreateMetastore_SdkV2
// only implements ToObjectValue() and Type().
func (o AccountsCreateMetastore_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_info": o.MetastoreInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsCreateMetastore_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_info": basetypes.ListType{
				ElemType: CreateMetastore_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetMetastoreInfo returns the value of the MetastoreInfo field in AccountsCreateMetastore_SdkV2 as
// a CreateMetastore_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsCreateMetastore_SdkV2) GetMetastoreInfo(ctx context.Context) (CreateMetastore_SdkV2, bool) {
	var e CreateMetastore_SdkV2
	if o.MetastoreInfo.IsNull() || o.MetastoreInfo.IsUnknown() {
		return e, false
	}
	var v []CreateMetastore_SdkV2
	d := o.MetastoreInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetastoreInfo sets the value of the MetastoreInfo field in AccountsCreateMetastore_SdkV2.
func (o *AccountsCreateMetastore_SdkV2) SetMetastoreInfo(ctx context.Context, v CreateMetastore_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metastore_info"]
	o.MetastoreInfo = types.ListValueMust(t, vs)
}

type AccountsCreateMetastoreAssignment_SdkV2 struct {
	MetastoreAssignment types.List `tfsdk:"metastore_assignment" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *AccountsCreateMetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateMetastoreAssignment_SdkV2) {
}

func (newState *AccountsCreateMetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccountsCreateMetastoreAssignment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsCreateMetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsCreateMetastoreAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_assignment": reflect.TypeOf(CreateMetastoreAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsCreateMetastoreAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o AccountsCreateMetastoreAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_assignment": o.MetastoreAssignment,
			"metastore_id":         o.MetastoreId,
			"workspace_id":         o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsCreateMetastoreAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_assignment": basetypes.ListType{
				ElemType: CreateMetastoreAssignment_SdkV2{}.Type(ctx),
			},
			"metastore_id": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// GetMetastoreAssignment returns the value of the MetastoreAssignment field in AccountsCreateMetastoreAssignment_SdkV2 as
// a CreateMetastoreAssignment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsCreateMetastoreAssignment_SdkV2) GetMetastoreAssignment(ctx context.Context) (CreateMetastoreAssignment_SdkV2, bool) {
	var e CreateMetastoreAssignment_SdkV2
	if o.MetastoreAssignment.IsNull() || o.MetastoreAssignment.IsUnknown() {
		return e, false
	}
	var v []CreateMetastoreAssignment_SdkV2
	d := o.MetastoreAssignment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetastoreAssignment sets the value of the MetastoreAssignment field in AccountsCreateMetastoreAssignment_SdkV2.
func (o *AccountsCreateMetastoreAssignment_SdkV2) SetMetastoreAssignment(ctx context.Context, v CreateMetastoreAssignment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metastore_assignment"]
	o.MetastoreAssignment = types.ListValueMust(t, vs)
}

type AccountsCreateStorageCredential_SdkV2 struct {
	CredentialInfo types.List `tfsdk:"credential_info" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *AccountsCreateStorageCredential_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateStorageCredential_SdkV2) {
}

func (newState *AccountsCreateStorageCredential_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccountsCreateStorageCredential_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsCreateStorageCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsCreateStorageCredential_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credential_info": reflect.TypeOf(CreateStorageCredential_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsCreateStorageCredential_SdkV2
// only implements ToObjectValue() and Type().
func (o AccountsCreateStorageCredential_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_info": o.CredentialInfo,
			"metastore_id":    o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsCreateStorageCredential_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_info": basetypes.ListType{
				ElemType: CreateStorageCredential_SdkV2{}.Type(ctx),
			},
			"metastore_id": types.StringType,
		},
	}
}

// GetCredentialInfo returns the value of the CredentialInfo field in AccountsCreateStorageCredential_SdkV2 as
// a CreateStorageCredential_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsCreateStorageCredential_SdkV2) GetCredentialInfo(ctx context.Context) (CreateStorageCredential_SdkV2, bool) {
	var e CreateStorageCredential_SdkV2
	if o.CredentialInfo.IsNull() || o.CredentialInfo.IsUnknown() {
		return e, false
	}
	var v []CreateStorageCredential_SdkV2
	d := o.CredentialInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCredentialInfo sets the value of the CredentialInfo field in AccountsCreateStorageCredential_SdkV2.
func (o *AccountsCreateStorageCredential_SdkV2) SetCredentialInfo(ctx context.Context, v CreateStorageCredential_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["credential_info"]
	o.CredentialInfo = types.ListValueMust(t, vs)
}

type AccountsMetastoreAssignment_SdkV2 struct {
	MetastoreAssignment types.List `tfsdk:"metastore_assignment" tf:"optional,object"`
}

func (newState *AccountsMetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsMetastoreAssignment_SdkV2) {
}

func (newState *AccountsMetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccountsMetastoreAssignment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsMetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsMetastoreAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_assignment": reflect.TypeOf(MetastoreAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsMetastoreAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o AccountsMetastoreAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_assignment": o.MetastoreAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsMetastoreAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_assignment": basetypes.ListType{
				ElemType: MetastoreAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetMetastoreAssignment returns the value of the MetastoreAssignment field in AccountsMetastoreAssignment_SdkV2 as
// a MetastoreAssignment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsMetastoreAssignment_SdkV2) GetMetastoreAssignment(ctx context.Context) (MetastoreAssignment_SdkV2, bool) {
	var e MetastoreAssignment_SdkV2
	if o.MetastoreAssignment.IsNull() || o.MetastoreAssignment.IsUnknown() {
		return e, false
	}
	var v []MetastoreAssignment_SdkV2
	d := o.MetastoreAssignment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetastoreAssignment sets the value of the MetastoreAssignment field in AccountsMetastoreAssignment_SdkV2.
func (o *AccountsMetastoreAssignment_SdkV2) SetMetastoreAssignment(ctx context.Context, v MetastoreAssignment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metastore_assignment"]
	o.MetastoreAssignment = types.ListValueMust(t, vs)
}

type AccountsMetastoreInfo_SdkV2 struct {
	MetastoreInfo types.List `tfsdk:"metastore_info" tf:"optional,object"`
}

func (newState *AccountsMetastoreInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsMetastoreInfo_SdkV2) {
}

func (newState *AccountsMetastoreInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccountsMetastoreInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsMetastoreInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsMetastoreInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_info": reflect.TypeOf(MetastoreInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsMetastoreInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o AccountsMetastoreInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_info": o.MetastoreInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsMetastoreInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_info": basetypes.ListType{
				ElemType: MetastoreInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetMetastoreInfo returns the value of the MetastoreInfo field in AccountsMetastoreInfo_SdkV2 as
// a MetastoreInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsMetastoreInfo_SdkV2) GetMetastoreInfo(ctx context.Context) (MetastoreInfo_SdkV2, bool) {
	var e MetastoreInfo_SdkV2
	if o.MetastoreInfo.IsNull() || o.MetastoreInfo.IsUnknown() {
		return e, false
	}
	var v []MetastoreInfo_SdkV2
	d := o.MetastoreInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetastoreInfo sets the value of the MetastoreInfo field in AccountsMetastoreInfo_SdkV2.
func (o *AccountsMetastoreInfo_SdkV2) SetMetastoreInfo(ctx context.Context, v MetastoreInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metastore_info"]
	o.MetastoreInfo = types.ListValueMust(t, vs)
}

type AccountsStorageCredentialInfo_SdkV2 struct {
	CredentialInfo types.List `tfsdk:"credential_info" tf:"optional,object"`
}

func (newState *AccountsStorageCredentialInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsStorageCredentialInfo_SdkV2) {
}

func (newState *AccountsStorageCredentialInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccountsStorageCredentialInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsStorageCredentialInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsStorageCredentialInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credential_info": reflect.TypeOf(StorageCredentialInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsStorageCredentialInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o AccountsStorageCredentialInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_info": o.CredentialInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsStorageCredentialInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_info": basetypes.ListType{
				ElemType: StorageCredentialInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetCredentialInfo returns the value of the CredentialInfo field in AccountsStorageCredentialInfo_SdkV2 as
// a StorageCredentialInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsStorageCredentialInfo_SdkV2) GetCredentialInfo(ctx context.Context) (StorageCredentialInfo_SdkV2, bool) {
	var e StorageCredentialInfo_SdkV2
	if o.CredentialInfo.IsNull() || o.CredentialInfo.IsUnknown() {
		return e, false
	}
	var v []StorageCredentialInfo_SdkV2
	d := o.CredentialInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCredentialInfo sets the value of the CredentialInfo field in AccountsStorageCredentialInfo_SdkV2.
func (o *AccountsStorageCredentialInfo_SdkV2) SetCredentialInfo(ctx context.Context, v StorageCredentialInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["credential_info"]
	o.CredentialInfo = types.ListValueMust(t, vs)
}

type AccountsUpdateMetastore_SdkV2 struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`

	MetastoreInfo types.List `tfsdk:"metastore_info" tf:"optional,object"`
}

func (newState *AccountsUpdateMetastore_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateMetastore_SdkV2) {
}

func (newState *AccountsUpdateMetastore_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateMetastore_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsUpdateMetastore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsUpdateMetastore_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_info": reflect.TypeOf(UpdateMetastore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsUpdateMetastore_SdkV2
// only implements ToObjectValue() and Type().
func (o AccountsUpdateMetastore_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id":   o.MetastoreId,
			"metastore_info": o.MetastoreInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsUpdateMetastore_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
			"metastore_info": basetypes.ListType{
				ElemType: UpdateMetastore_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetMetastoreInfo returns the value of the MetastoreInfo field in AccountsUpdateMetastore_SdkV2 as
// a UpdateMetastore_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsUpdateMetastore_SdkV2) GetMetastoreInfo(ctx context.Context) (UpdateMetastore_SdkV2, bool) {
	var e UpdateMetastore_SdkV2
	if o.MetastoreInfo.IsNull() || o.MetastoreInfo.IsUnknown() {
		return e, false
	}
	var v []UpdateMetastore_SdkV2
	d := o.MetastoreInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetastoreInfo sets the value of the MetastoreInfo field in AccountsUpdateMetastore_SdkV2.
func (o *AccountsUpdateMetastore_SdkV2) SetMetastoreInfo(ctx context.Context, v UpdateMetastore_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metastore_info"]
	o.MetastoreInfo = types.ListValueMust(t, vs)
}

type AccountsUpdateMetastoreAssignment_SdkV2 struct {
	MetastoreAssignment types.List `tfsdk:"metastore_assignment" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *AccountsUpdateMetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateMetastoreAssignment_SdkV2) {
}

func (newState *AccountsUpdateMetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateMetastoreAssignment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsUpdateMetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsUpdateMetastoreAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastore_assignment": reflect.TypeOf(UpdateMetastoreAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsUpdateMetastoreAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o AccountsUpdateMetastoreAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_assignment": o.MetastoreAssignment,
			"metastore_id":         o.MetastoreId,
			"workspace_id":         o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsUpdateMetastoreAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_assignment": basetypes.ListType{
				ElemType: UpdateMetastoreAssignment_SdkV2{}.Type(ctx),
			},
			"metastore_id": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// GetMetastoreAssignment returns the value of the MetastoreAssignment field in AccountsUpdateMetastoreAssignment_SdkV2 as
// a UpdateMetastoreAssignment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsUpdateMetastoreAssignment_SdkV2) GetMetastoreAssignment(ctx context.Context) (UpdateMetastoreAssignment_SdkV2, bool) {
	var e UpdateMetastoreAssignment_SdkV2
	if o.MetastoreAssignment.IsNull() || o.MetastoreAssignment.IsUnknown() {
		return e, false
	}
	var v []UpdateMetastoreAssignment_SdkV2
	d := o.MetastoreAssignment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMetastoreAssignment sets the value of the MetastoreAssignment field in AccountsUpdateMetastoreAssignment_SdkV2.
func (o *AccountsUpdateMetastoreAssignment_SdkV2) SetMetastoreAssignment(ctx context.Context, v UpdateMetastoreAssignment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metastore_assignment"]
	o.MetastoreAssignment = types.ListValueMust(t, vs)
}

type AccountsUpdateStorageCredential_SdkV2 struct {
	CredentialInfo types.List `tfsdk:"credential_info" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Name of the storage credential.
	StorageCredentialName types.String `tfsdk:"-"`
}

func (newState *AccountsUpdateStorageCredential_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateStorageCredential_SdkV2) {
}

func (newState *AccountsUpdateStorageCredential_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateStorageCredential_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountsUpdateStorageCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccountsUpdateStorageCredential_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credential_info": reflect.TypeOf(UpdateStorageCredential_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountsUpdateStorageCredential_SdkV2
// only implements ToObjectValue() and Type().
func (o AccountsUpdateStorageCredential_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_info":         o.CredentialInfo,
			"metastore_id":            o.MetastoreId,
			"storage_credential_name": o.StorageCredentialName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountsUpdateStorageCredential_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_info": basetypes.ListType{
				ElemType: UpdateStorageCredential_SdkV2{}.Type(ctx),
			},
			"metastore_id":            types.StringType,
			"storage_credential_name": types.StringType,
		},
	}
}

// GetCredentialInfo returns the value of the CredentialInfo field in AccountsUpdateStorageCredential_SdkV2 as
// a UpdateStorageCredential_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AccountsUpdateStorageCredential_SdkV2) GetCredentialInfo(ctx context.Context) (UpdateStorageCredential_SdkV2, bool) {
	var e UpdateStorageCredential_SdkV2
	if o.CredentialInfo.IsNull() || o.CredentialInfo.IsUnknown() {
		return e, false
	}
	var v []UpdateStorageCredential_SdkV2
	d := o.CredentialInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCredentialInfo sets the value of the CredentialInfo field in AccountsUpdateStorageCredential_SdkV2.
func (o *AccountsUpdateStorageCredential_SdkV2) SetCredentialInfo(ctx context.Context, v UpdateStorageCredential_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["credential_info"]
	o.CredentialInfo = types.ListValueMust(t, vs)
}

type ArtifactAllowlistInfo_SdkV2 struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers types.List `tfsdk:"artifact_matchers" tf:"optional"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of the user who set the artifact allowlist.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
}

func (newState *ArtifactAllowlistInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ArtifactAllowlistInfo_SdkV2) {
}

func (newState *ArtifactAllowlistInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ArtifactAllowlistInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ArtifactAllowlistInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ArtifactAllowlistInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"artifact_matchers": reflect.TypeOf(ArtifactMatcher_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ArtifactAllowlistInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ArtifactAllowlistInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ArtifactAllowlistInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_matchers": basetypes.ListType{
				ElemType: ArtifactMatcher_SdkV2{}.Type(ctx),
			},
			"created_at":   types.Int64Type,
			"created_by":   types.StringType,
			"metastore_id": types.StringType,
		},
	}
}

// GetArtifactMatchers returns the value of the ArtifactMatchers field in ArtifactAllowlistInfo_SdkV2 as
// a slice of ArtifactMatcher_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ArtifactAllowlistInfo_SdkV2) GetArtifactMatchers(ctx context.Context) ([]ArtifactMatcher_SdkV2, bool) {
	if o.ArtifactMatchers.IsNull() || o.ArtifactMatchers.IsUnknown() {
		return nil, false
	}
	var v []ArtifactMatcher_SdkV2
	d := o.ArtifactMatchers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetArtifactMatchers sets the value of the ArtifactMatchers field in ArtifactAllowlistInfo_SdkV2.
func (o *ArtifactAllowlistInfo_SdkV2) SetArtifactMatchers(ctx context.Context, v []ArtifactMatcher_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["artifact_matchers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ArtifactMatchers = types.ListValueMust(t, vs)
}

type ArtifactMatcher_SdkV2 struct {
	// The artifact path or maven coordinate
	Artifact types.String `tfsdk:"artifact" tf:""`
	// The pattern matching type of the artifact
	MatchType types.String `tfsdk:"match_type" tf:""`
}

func (newState *ArtifactMatcher_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ArtifactMatcher_SdkV2) {
}

func (newState *ArtifactMatcher_SdkV2) SyncEffectiveFieldsDuringRead(existingState ArtifactMatcher_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ArtifactMatcher.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ArtifactMatcher_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ArtifactMatcher_SdkV2
// only implements ToObjectValue() and Type().
func (o ArtifactMatcher_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact":   o.Artifact,
			"match_type": o.MatchType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ArtifactMatcher_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact":   types.StringType,
			"match_type": types.StringType,
		},
	}
}

type AssignResponse_SdkV2 struct {
}

func (newState *AssignResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AssignResponse_SdkV2) {
}

func (newState *AssignResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState AssignResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AssignResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AssignResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AssignResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o AssignResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o AssignResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// AWS temporary credentials for API authentication. Read more at
// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
type AwsCredentials_SdkV2 struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId types.String `tfsdk:"access_key_id" tf:"optional"`
	// The Amazon Resource Name (ARN) of the S3 access point for temporary
	// credentials related the external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// The secret access key that can be used to sign AWS API requests.
	SecretAccessKey types.String `tfsdk:"secret_access_key" tf:"optional"`
	// The token that users must pass to AWS API to use the temporary
	// credentials.
	SessionToken types.String `tfsdk:"session_token" tf:"optional"`
}

func (newState *AwsCredentials_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsCredentials_SdkV2) {
}

func (newState *AwsCredentials_SdkV2) SyncEffectiveFieldsDuringRead(existingState AwsCredentials_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsCredentials_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsCredentials_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsCredentials_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AwsCredentials_SdkV2) Type(ctx context.Context) attr.Type {
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
type AwsIamRole_SdkV2 struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	ExternalId types.String `tfsdk:"external_id" tf:"computed"`
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	RoleArn types.String `tfsdk:"role_arn" tf:"optional"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn types.String `tfsdk:"unity_catalog_iam_arn" tf:"computed"`
}

func (newState *AwsIamRole_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRole_SdkV2) {
}

func (newState *AwsIamRole_SdkV2) SyncEffectiveFieldsDuringRead(existingState AwsIamRole_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsIamRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsIamRole_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsIamRole_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsIamRole_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id":           o.ExternalId,
			"role_arn":              o.RoleArn,
			"unity_catalog_iam_arn": o.UnityCatalogIamArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsIamRole_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id":           types.StringType,
			"role_arn":              types.StringType,
			"unity_catalog_iam_arn": types.StringType,
		},
	}
}

type AwsIamRoleRequest_SdkV2 struct {
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn types.String `tfsdk:"role_arn" tf:""`
}

func (newState *AwsIamRoleRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRoleRequest_SdkV2) {
}

func (newState *AwsIamRoleRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState AwsIamRoleRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsIamRoleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsIamRoleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsIamRoleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsIamRoleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"role_arn": o.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsIamRoleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"role_arn": types.StringType,
		},
	}
}

type AwsIamRoleResponse_SdkV2 struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	ExternalId types.String `tfsdk:"external_id" tf:"optional"`
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn types.String `tfsdk:"role_arn" tf:""`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn types.String `tfsdk:"unity_catalog_iam_arn" tf:"optional"`
}

func (newState *AwsIamRoleResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRoleResponse_SdkV2) {
}

func (newState *AwsIamRoleResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState AwsIamRoleResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsIamRoleResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsIamRoleResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsIamRoleResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsIamRoleResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id":           o.ExternalId,
			"role_arn":              o.RoleArn,
			"unity_catalog_iam_arn": o.UnityCatalogIamArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsIamRoleResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id":           types.StringType,
			"role_arn":              types.StringType,
			"unity_catalog_iam_arn": types.StringType,
		},
	}
}

// Azure Active Directory token, essentially the Oauth token for Azure Service
// Principal or Managed Identity. Read more at
// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
type AzureActiveDirectoryToken_SdkV2 struct {
	// Opaque token that contains claims that you can use in Azure Active
	// Directory to access cloud services.
	AadToken types.String `tfsdk:"aad_token" tf:"optional"`
}

func (newState *AzureActiveDirectoryToken_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureActiveDirectoryToken_SdkV2) {
}

func (newState *AzureActiveDirectoryToken_SdkV2) SyncEffectiveFieldsDuringRead(existingState AzureActiveDirectoryToken_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureActiveDirectoryToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureActiveDirectoryToken_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureActiveDirectoryToken_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureActiveDirectoryToken_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aad_token": o.AadToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureActiveDirectoryToken_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aad_token": types.StringType,
		},
	}
}

// The Azure managed identity configuration.
type AzureManagedIdentity_SdkV2 struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	AccessConnectorId types.String `tfsdk:"access_connector_id" tf:""`
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database. .
	CredentialId types.String `tfsdk:"credential_id" tf:"optional"`
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	ManagedIdentityId types.String `tfsdk:"managed_identity_id" tf:"optional"`
}

func (newState *AzureManagedIdentity_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentity_SdkV2) {
}

func (newState *AzureManagedIdentity_SdkV2) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentity_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureManagedIdentity.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureManagedIdentity_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureManagedIdentity_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureManagedIdentity_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_connector_id": o.AccessConnectorId,
			"credential_id":       o.CredentialId,
			"managed_identity_id": o.ManagedIdentityId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureManagedIdentity_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_connector_id": types.StringType,
			"credential_id":       types.StringType,
			"managed_identity_id": types.StringType,
		},
	}
}

type AzureManagedIdentityRequest_SdkV2 struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId types.String `tfsdk:"access_connector_id" tf:""`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId types.String `tfsdk:"managed_identity_id" tf:"optional"`
}

func (newState *AzureManagedIdentityRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentityRequest_SdkV2) {
}

func (newState *AzureManagedIdentityRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentityRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureManagedIdentityRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureManagedIdentityRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureManagedIdentityRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureManagedIdentityRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_connector_id": o.AccessConnectorId,
			"managed_identity_id": o.ManagedIdentityId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureManagedIdentityRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_connector_id": types.StringType,
			"managed_identity_id": types.StringType,
		},
	}
}

type AzureManagedIdentityResponse_SdkV2 struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId types.String `tfsdk:"access_connector_id" tf:""`
	// The Databricks internal ID that represents this managed identity.
	CredentialId types.String `tfsdk:"credential_id" tf:"optional"`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId types.String `tfsdk:"managed_identity_id" tf:"optional"`
}

func (newState *AzureManagedIdentityResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentityResponse_SdkV2) {
}

func (newState *AzureManagedIdentityResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentityResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureManagedIdentityResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureManagedIdentityResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureManagedIdentityResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureManagedIdentityResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_connector_id": o.AccessConnectorId,
			"credential_id":       o.CredentialId,
			"managed_identity_id": o.ManagedIdentityId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureManagedIdentityResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_connector_id": types.StringType,
			"credential_id":       types.StringType,
			"managed_identity_id": types.StringType,
		},
	}
}

// The Azure service principal configuration. Only applicable when purpose is
// **STORAGE**.
type AzureServicePrincipal_SdkV2 struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	ApplicationId types.String `tfsdk:"application_id" tf:""`
	// The client secret generated for the above app ID in AAD.
	ClientSecret types.String `tfsdk:"client_secret" tf:""`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	DirectoryId types.String `tfsdk:"directory_id" tf:""`
}

func (newState *AzureServicePrincipal_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureServicePrincipal_SdkV2) {
}

func (newState *AzureServicePrincipal_SdkV2) SyncEffectiveFieldsDuringRead(existingState AzureServicePrincipal_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureServicePrincipal.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureServicePrincipal_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureServicePrincipal_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureServicePrincipal_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"application_id": o.ApplicationId,
			"client_secret":  o.ClientSecret,
			"directory_id":   o.DirectoryId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureServicePrincipal_SdkV2) Type(ctx context.Context) attr.Type {
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
type AzureUserDelegationSas_SdkV2 struct {
	// The signed URI (SAS Token) used to access blob services for a given path
	SasToken types.String `tfsdk:"sas_token" tf:"optional"`
}

func (newState *AzureUserDelegationSas_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureUserDelegationSas_SdkV2) {
}

func (newState *AzureUserDelegationSas_SdkV2) SyncEffectiveFieldsDuringRead(existingState AzureUserDelegationSas_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureUserDelegationSas.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureUserDelegationSas_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureUserDelegationSas_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureUserDelegationSas_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sas_token": o.SasToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureUserDelegationSas_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sas_token": types.StringType,
		},
	}
}

// Cancel refresh
type CancelRefreshRequest_SdkV2 struct {
	// ID of the refresh.
	RefreshId types.String `tfsdk:"-"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *CancelRefreshRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelRefreshRequest_SdkV2) {
}

func (newState *CancelRefreshRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CancelRefreshRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelRefreshRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelRefreshRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelRefreshRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"refresh_id": o.RefreshId,
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CancelRefreshRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"refresh_id": types.StringType,
			"table_name": types.StringType,
		},
	}
}

type CancelRefreshResponse_SdkV2 struct {
}

func (newState *CancelRefreshResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelRefreshResponse_SdkV2) {
}

func (newState *CancelRefreshResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CancelRefreshResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelRefreshResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelRefreshResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelRefreshResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelRefreshResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CancelRefreshResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CatalogInfo_SdkV2 struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// The type of the catalog.
	CatalogType types.String `tfsdk:"catalog_type" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the connection to an external data source.
	ConnectionName types.String `tfsdk:"connection_name" tf:"optional"`
	// Time at which this catalog was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of catalog creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	EffectivePredictiveOptimizationFlag types.List `tfsdk:"effective_predictive_optimization_flag" tf:"optional,object"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization" tf:"optional"`
	// The full name of the catalog. Corresponds with the name field.
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of catalog.
	Name types.String `tfsdk:"name" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options" tf:"optional"`
	// Username of current owner of catalog.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties" tf:"optional"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName types.String `tfsdk:"provider_name" tf:"optional"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo types.List `tfsdk:"provisioning_info" tf:"optional,object"`
	// Kind of catalog securable.
	SecurableKind types.String `tfsdk:"securable_kind" tf:"optional"`

	SecurableType types.String `tfsdk:"securable_type" tf:"optional"`
	// The name of the share under the share provider.
	ShareName types.String `tfsdk:"share_name" tf:"optional"`
	// Storage Location URL (full path) for managed tables within catalog.
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Storage root URL for managed tables within catalog.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Time at which this catalog was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified catalog.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *CatalogInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CatalogInfo_SdkV2) {
}

func (newState *CatalogInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState CatalogInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CatalogInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CatalogInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_predictive_optimization_flag": reflect.TypeOf(EffectivePredictiveOptimizationFlag_SdkV2{}),
		"options":                                reflect.TypeOf(types.String{}),
		"properties":                             reflect.TypeOf(types.String{}),
		"provisioning_info":                      reflect.TypeOf(ProvisioningInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CatalogInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o CatalogInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
			"securable_kind":                         o.SecurableKind,
			"securable_type":                         o.SecurableType,
			"share_name":                             o.ShareName,
			"storage_location":                       o.StorageLocation,
			"storage_root":                           o.StorageRoot,
			"updated_at":                             o.UpdatedAt,
			"updated_by":                             o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CatalogInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":     types.BoolType,
			"catalog_type":    types.StringType,
			"comment":         types.StringType,
			"connection_name": types.StringType,
			"created_at":      types.Int64Type,
			"created_by":      types.StringType,
			"effective_predictive_optimization_flag": basetypes.ListType{
				ElemType: EffectivePredictiveOptimizationFlag_SdkV2{}.Type(ctx),
			},
			"enable_predictive_optimization": types.StringType,
			"full_name":                      types.StringType,
			"isolation_mode":                 types.StringType,
			"metastore_id":                   types.StringType,
			"name":                           types.StringType,
			"options": basetypes.MapType{
				ElemType: types.StringType,
			},
			"owner": types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"provider_name": types.StringType,
			"provisioning_info": basetypes.ListType{
				ElemType: ProvisioningInfo_SdkV2{}.Type(ctx),
			},
			"securable_kind":   types.StringType,
			"securable_type":   types.StringType,
			"share_name":       types.StringType,
			"storage_location": types.StringType,
			"storage_root":     types.StringType,
			"updated_at":       types.Int64Type,
			"updated_by":       types.StringType,
		},
	}
}

// GetEffectivePredictiveOptimizationFlag returns the value of the EffectivePredictiveOptimizationFlag field in CatalogInfo_SdkV2 as
// a EffectivePredictiveOptimizationFlag_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CatalogInfo_SdkV2) GetEffectivePredictiveOptimizationFlag(ctx context.Context) (EffectivePredictiveOptimizationFlag_SdkV2, bool) {
	var e EffectivePredictiveOptimizationFlag_SdkV2
	if o.EffectivePredictiveOptimizationFlag.IsNull() || o.EffectivePredictiveOptimizationFlag.IsUnknown() {
		return e, false
	}
	var v []EffectivePredictiveOptimizationFlag_SdkV2
	d := o.EffectivePredictiveOptimizationFlag.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectivePredictiveOptimizationFlag sets the value of the EffectivePredictiveOptimizationFlag field in CatalogInfo_SdkV2.
func (o *CatalogInfo_SdkV2) SetEffectivePredictiveOptimizationFlag(ctx context.Context, v EffectivePredictiveOptimizationFlag_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_predictive_optimization_flag"]
	o.EffectivePredictiveOptimizationFlag = types.ListValueMust(t, vs)
}

// GetOptions returns the value of the Options field in CatalogInfo_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CatalogInfo_SdkV2) GetOptions(ctx context.Context) (map[string]types.String, bool) {
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

// SetOptions sets the value of the Options field in CatalogInfo_SdkV2.
func (o *CatalogInfo_SdkV2) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in CatalogInfo_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CatalogInfo_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in CatalogInfo_SdkV2.
func (o *CatalogInfo_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

// GetProvisioningInfo returns the value of the ProvisioningInfo field in CatalogInfo_SdkV2 as
// a ProvisioningInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CatalogInfo_SdkV2) GetProvisioningInfo(ctx context.Context) (ProvisioningInfo_SdkV2, bool) {
	var e ProvisioningInfo_SdkV2
	if o.ProvisioningInfo.IsNull() || o.ProvisioningInfo.IsUnknown() {
		return e, false
	}
	var v []ProvisioningInfo_SdkV2
	d := o.ProvisioningInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvisioningInfo sets the value of the ProvisioningInfo field in CatalogInfo_SdkV2.
func (o *CatalogInfo_SdkV2) SetProvisioningInfo(ctx context.Context, v ProvisioningInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provisioning_info"]
	o.ProvisioningInfo = types.ListValueMust(t, vs)
}

type CloudflareApiToken_SdkV2 struct {
	// The Cloudflare access key id of the token.
	AccessKeyId types.String `tfsdk:"access_key_id" tf:""`
	// The account id associated with the API token.
	AccountId types.String `tfsdk:"account_id" tf:""`
	// The secret access token generated for the access key id
	SecretAccessKey types.String `tfsdk:"secret_access_key" tf:""`
}

func (newState *CloudflareApiToken_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloudflareApiToken_SdkV2) {
}

func (newState *CloudflareApiToken_SdkV2) SyncEffectiveFieldsDuringRead(existingState CloudflareApiToken_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloudflareApiToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CloudflareApiToken_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudflareApiToken_SdkV2
// only implements ToObjectValue() and Type().
func (o CloudflareApiToken_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_key_id":     o.AccessKeyId,
			"account_id":        o.AccountId,
			"secret_access_key": o.SecretAccessKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CloudflareApiToken_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_key_id":     types.StringType,
			"account_id":        types.StringType,
			"secret_access_key": types.StringType,
		},
	}
}

type ColumnInfo_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`

	Mask types.List `tfsdk:"mask" tf:"optional,object"`
	// Name of Column.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Whether field may be Null (default: true).
	Nullable types.Bool `tfsdk:"nullable" tf:"optional"`
	// Partition index for column.
	PartitionIndex types.Int64 `tfsdk:"partition_index" tf:"optional"`
	// Ordinal position of column (starting at position 0).
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// Format of IntervalType.
	TypeIntervalType types.String `tfsdk:"type_interval_type" tf:"optional"`
	// Full data type specification, JSON-serialized.
	TypeJson types.String `tfsdk:"type_json" tf:"optional"`

	TypeName types.String `tfsdk:"type_name" tf:"optional"`
	// Digits of precision; required for DecimalTypes.
	TypePrecision types.Int64 `tfsdk:"type_precision" tf:"optional"`
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale types.Int64 `tfsdk:"type_scale" tf:"optional"`
	// Full data type specification as SQL/catalogString text.
	TypeText types.String `tfsdk:"type_text" tf:"optional"`
}

func (newState *ColumnInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnInfo_SdkV2) {
}

func (newState *ColumnInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ColumnInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ColumnInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ColumnInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"mask": reflect.TypeOf(ColumnMask_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ColumnInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ColumnInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"mask": basetypes.ListType{
				ElemType: ColumnMask_SdkV2{}.Type(ctx),
			},
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

// GetMask returns the value of the Mask field in ColumnInfo_SdkV2 as
// a ColumnMask_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ColumnInfo_SdkV2) GetMask(ctx context.Context) (ColumnMask_SdkV2, bool) {
	var e ColumnMask_SdkV2
	if o.Mask.IsNull() || o.Mask.IsUnknown() {
		return e, false
	}
	var v []ColumnMask_SdkV2
	d := o.Mask.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMask sets the value of the Mask field in ColumnInfo_SdkV2.
func (o *ColumnInfo_SdkV2) SetMask(ctx context.Context, v ColumnMask_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["mask"]
	o.Mask = types.ListValueMust(t, vs)
}

type ColumnMask_SdkV2 struct {
	// The full name of the column mask SQL UDF.
	FunctionName types.String `tfsdk:"function_name" tf:"optional"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames types.List `tfsdk:"using_column_names" tf:"optional"`
}

func (newState *ColumnMask_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnMask_SdkV2) {
}

func (newState *ColumnMask_SdkV2) SyncEffectiveFieldsDuringRead(existingState ColumnMask_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ColumnMask.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ColumnMask_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"using_column_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnMask_SdkV2
// only implements ToObjectValue() and Type().
func (o ColumnMask_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_name":      o.FunctionName,
			"using_column_names": o.UsingColumnNames,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ColumnMask_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_name": types.StringType,
			"using_column_names": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetUsingColumnNames returns the value of the UsingColumnNames field in ColumnMask_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ColumnMask_SdkV2) GetUsingColumnNames(ctx context.Context) ([]types.String, bool) {
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

// SetUsingColumnNames sets the value of the UsingColumnNames field in ColumnMask_SdkV2.
func (o *ColumnMask_SdkV2) SetUsingColumnNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["using_column_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UsingColumnNames = types.ListValueMust(t, vs)
}

type ConnectionInfo_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Unique identifier of the Connection.
	ConnectionId types.String `tfsdk:"connection_id" tf:"optional"`
	// The type of connection.
	ConnectionType types.String `tfsdk:"connection_type" tf:"optional"`
	// Time at which this connection was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of connection creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The type of credential.
	CredentialType types.String `tfsdk:"credential_type" tf:"optional"`
	// Full name of connection.
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of the connection.
	Name types.String `tfsdk:"name" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options" tf:"optional"`
	// Username of current owner of the connection.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties types.Map `tfsdk:"properties" tf:"optional"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo types.List `tfsdk:"provisioning_info" tf:"optional,object"`
	// If the connection is read only.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Kind of connection securable.
	SecurableKind types.String `tfsdk:"securable_kind" tf:"optional"`

	SecurableType types.String `tfsdk:"securable_type" tf:"optional"`
	// Time at which this connection was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified connection.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// URL of the remote data source, extracted from options.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *ConnectionInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ConnectionInfo_SdkV2) {
}

func (newState *ConnectionInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ConnectionInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ConnectionInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ConnectionInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options":           reflect.TypeOf(types.String{}),
		"properties":        reflect.TypeOf(types.String{}),
		"provisioning_info": reflect.TypeOf(ProvisioningInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ConnectionInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ConnectionInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
			"securable_kind":    o.SecurableKind,
			"securable_type":    o.SecurableType,
			"updated_at":        o.UpdatedAt,
			"updated_by":        o.UpdatedBy,
			"url":               o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ConnectionInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
			"provisioning_info": basetypes.ListType{
				ElemType: ProvisioningInfo_SdkV2{}.Type(ctx),
			},
			"read_only":      types.BoolType,
			"securable_kind": types.StringType,
			"securable_type": types.StringType,
			"updated_at":     types.Int64Type,
			"updated_by":     types.StringType,
			"url":            types.StringType,
		},
	}
}

// GetOptions returns the value of the Options field in ConnectionInfo_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ConnectionInfo_SdkV2) GetOptions(ctx context.Context) (map[string]types.String, bool) {
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

// SetOptions sets the value of the Options field in ConnectionInfo_SdkV2.
func (o *ConnectionInfo_SdkV2) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in ConnectionInfo_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ConnectionInfo_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in ConnectionInfo_SdkV2.
func (o *ConnectionInfo_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

// GetProvisioningInfo returns the value of the ProvisioningInfo field in ConnectionInfo_SdkV2 as
// a ProvisioningInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ConnectionInfo_SdkV2) GetProvisioningInfo(ctx context.Context) (ProvisioningInfo_SdkV2, bool) {
	var e ProvisioningInfo_SdkV2
	if o.ProvisioningInfo.IsNull() || o.ProvisioningInfo.IsUnknown() {
		return e, false
	}
	var v []ProvisioningInfo_SdkV2
	d := o.ProvisioningInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvisioningInfo sets the value of the ProvisioningInfo field in ConnectionInfo_SdkV2.
func (o *ConnectionInfo_SdkV2) SetProvisioningInfo(ctx context.Context, v ProvisioningInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provisioning_info"]
	o.ProvisioningInfo = types.ListValueMust(t, vs)
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
type ContinuousUpdateStatus_SdkV2 struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress types.List `tfsdk:"initial_pipeline_sync_progress" tf:"optional,object"`
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version" tf:"optional"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp types.String `tfsdk:"timestamp" tf:"optional"`
}

func (newState *ContinuousUpdateStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ContinuousUpdateStatus_SdkV2) {
}

func (newState *ContinuousUpdateStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState ContinuousUpdateStatus_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ContinuousUpdateStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ContinuousUpdateStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(PipelineProgress_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ContinuousUpdateStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o ContinuousUpdateStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": o.InitialPipelineSyncProgress,
			"last_processed_commit_version":  o.LastProcessedCommitVersion,
			"timestamp":                      o.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ContinuousUpdateStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": basetypes.ListType{
				ElemType: PipelineProgress_SdkV2{}.Type(ctx),
			},
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
		},
	}
}

// GetInitialPipelineSyncProgress returns the value of the InitialPipelineSyncProgress field in ContinuousUpdateStatus_SdkV2 as
// a PipelineProgress_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ContinuousUpdateStatus_SdkV2) GetInitialPipelineSyncProgress(ctx context.Context) (PipelineProgress_SdkV2, bool) {
	var e PipelineProgress_SdkV2
	if o.InitialPipelineSyncProgress.IsNull() || o.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v []PipelineProgress_SdkV2
	d := o.InitialPipelineSyncProgress.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in ContinuousUpdateStatus_SdkV2.
func (o *ContinuousUpdateStatus_SdkV2) SetInitialPipelineSyncProgress(ctx context.Context, v PipelineProgress_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["initial_pipeline_sync_progress"]
	o.InitialPipelineSyncProgress = types.ListValueMust(t, vs)
}

type CreateCatalog_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the connection to an external data source.
	ConnectionName types.String `tfsdk:"connection_name" tf:"optional"`
	// Name of catalog.
	Name types.String `tfsdk:"name" tf:""`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties" tf:"optional"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName types.String `tfsdk:"provider_name" tf:"optional"`
	// The name of the share under the share provider.
	ShareName types.String `tfsdk:"share_name" tf:"optional"`
	// Storage root URL for managed tables within catalog.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

func (newState *CreateCatalog_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCatalog_SdkV2) {
}

func (newState *CreateCatalog_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateCatalog_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCatalog.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCatalog_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options":    reflect.TypeOf(types.String{}),
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCatalog_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCatalog_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateCatalog_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetOptions returns the value of the Options field in CreateCatalog_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCatalog_SdkV2) GetOptions(ctx context.Context) (map[string]types.String, bool) {
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

// SetOptions sets the value of the Options field in CreateCatalog_SdkV2.
func (o *CreateCatalog_SdkV2) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in CreateCatalog_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCatalog_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in CreateCatalog_SdkV2.
func (o *CreateCatalog_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type CreateConnection_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The type of connection.
	ConnectionType types.String `tfsdk:"connection_type" tf:""`
	// Name of the connection.
	Name types.String `tfsdk:"name" tf:""`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options" tf:""`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties types.Map `tfsdk:"properties" tf:"optional"`
	// If the connection is read only.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
}

func (newState *CreateConnection_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateConnection_SdkV2) {
}

func (newState *CreateConnection_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateConnection_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateConnection.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateConnection_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options":    reflect.TypeOf(types.String{}),
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateConnection_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateConnection_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateConnection_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetOptions returns the value of the Options field in CreateConnection_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateConnection_SdkV2) GetOptions(ctx context.Context) (map[string]types.String, bool) {
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

// SetOptions sets the value of the Options field in CreateConnection_SdkV2.
func (o *CreateConnection_SdkV2) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in CreateConnection_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateConnection_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in CreateConnection_SdkV2.
func (o *CreateConnection_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type CreateCredentialRequest_SdkV2 struct {
	// The AWS IAM role configuration
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount types.List `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name types.String `tfsdk:"name" tf:""`
	// Indicates the purpose of the credential.
	Purpose types.String `tfsdk:"purpose" tf:"optional"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
}

func (newState *CreateCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialRequest_SdkV2) {
}

func (newState *CreateCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRole_SdkV2{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentity_SdkV2{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal_SdkV2{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccount_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRole_SdkV2{}.Type(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentity_SdkV2{}.Type(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal_SdkV2{}.Type(ctx),
			},
			"comment": types.StringType,
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccount_SdkV2{}.Type(ctx),
			},
			"name":            types.StringType,
			"purpose":         types.StringType,
			"read_only":       types.BoolType,
			"skip_validation": types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in CreateCredentialRequest_SdkV2 as
// a AwsIamRole_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest_SdkV2) GetAwsIamRole(ctx context.Context) (AwsIamRole_SdkV2, bool) {
	var e AwsIamRole_SdkV2
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRole_SdkV2
	d := o.AwsIamRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsIamRole sets the value of the AwsIamRole field in CreateCredentialRequest_SdkV2.
func (o *CreateCredentialRequest_SdkV2) SetAwsIamRole(ctx context.Context, v AwsIamRole_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_iam_role"]
	o.AwsIamRole = types.ListValueMust(t, vs)
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in CreateCredentialRequest_SdkV2 as
// a AzureManagedIdentity_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest_SdkV2) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentity_SdkV2, bool) {
	var e AzureManagedIdentity_SdkV2
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentity_SdkV2
	d := o.AzureManagedIdentity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in CreateCredentialRequest_SdkV2.
func (o *CreateCredentialRequest_SdkV2) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentity_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_managed_identity"]
	o.AzureManagedIdentity = types.ListValueMust(t, vs)
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in CreateCredentialRequest_SdkV2 as
// a AzureServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest_SdkV2) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal_SdkV2, bool) {
	var e AzureServicePrincipal_SdkV2
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal_SdkV2
	d := o.AzureServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in CreateCredentialRequest_SdkV2.
func (o *CreateCredentialRequest_SdkV2) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_service_principal"]
	o.AzureServicePrincipal = types.ListValueMust(t, vs)
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in CreateCredentialRequest_SdkV2 as
// a DatabricksGcpServiceAccount_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest_SdkV2) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccount_SdkV2, bool) {
	var e DatabricksGcpServiceAccount_SdkV2
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccount_SdkV2
	d := o.DatabricksGcpServiceAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in CreateCredentialRequest_SdkV2.
func (o *CreateCredentialRequest_SdkV2) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccount_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["databricks_gcp_service_account"]
	o.DatabricksGcpServiceAccount = types.ListValueMust(t, vs)
}

type CreateExternalLocation_SdkV2 struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the storage credential used with this location.
	CredentialName types.String `tfsdk:"credential_name" tf:""`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails types.List `tfsdk:"encryption_details" tf:"optional,object"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback types.Bool `tfsdk:"fallback" tf:"optional"`
	// Name of the external location.
	Name types.String `tfsdk:"name" tf:""`
	// Indicates whether the external location is read-only.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
	// Path URL of the external location.
	Url types.String `tfsdk:"url" tf:""`
}

func (newState *CreateExternalLocation_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExternalLocation_SdkV2) {
}

func (newState *CreateExternalLocation_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateExternalLocation_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExternalLocation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExternalLocation_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"encryption_details": reflect.TypeOf(EncryptionDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExternalLocation_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateExternalLocation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_point":       o.AccessPoint,
			"comment":            o.Comment,
			"credential_name":    o.CredentialName,
			"encryption_details": o.EncryptionDetails,
			"fallback":           o.Fallback,
			"name":               o.Name,
			"read_only":          o.ReadOnly,
			"skip_validation":    o.SkipValidation,
			"url":                o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExternalLocation_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point":    types.StringType,
			"comment":         types.StringType,
			"credential_name": types.StringType,
			"encryption_details": basetypes.ListType{
				ElemType: EncryptionDetails_SdkV2{}.Type(ctx),
			},
			"fallback":        types.BoolType,
			"name":            types.StringType,
			"read_only":       types.BoolType,
			"skip_validation": types.BoolType,
			"url":             types.StringType,
		},
	}
}

// GetEncryptionDetails returns the value of the EncryptionDetails field in CreateExternalLocation_SdkV2 as
// a EncryptionDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateExternalLocation_SdkV2) GetEncryptionDetails(ctx context.Context) (EncryptionDetails_SdkV2, bool) {
	var e EncryptionDetails_SdkV2
	if o.EncryptionDetails.IsNull() || o.EncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []EncryptionDetails_SdkV2
	d := o.EncryptionDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEncryptionDetails sets the value of the EncryptionDetails field in CreateExternalLocation_SdkV2.
func (o *CreateExternalLocation_SdkV2) SetEncryptionDetails(ctx context.Context, v EncryptionDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["encryption_details"]
	o.EncryptionDetails = types.ListValueMust(t, vs)
}

type CreateFunction_SdkV2 struct {
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:""`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Scalar function return data type.
	DataType types.String `tfsdk:"data_type" tf:""`
	// External function language.
	ExternalLanguage types.String `tfsdk:"external_language" tf:"optional"`
	// External function name.
	ExternalName types.String `tfsdk:"external_name" tf:"optional"`
	// Pretty printed function data type.
	FullDataType types.String `tfsdk:"full_data_type" tf:""`

	InputParams types.List `tfsdk:"input_params" tf:"object"`
	// Whether the function is deterministic.
	IsDeterministic types.Bool `tfsdk:"is_deterministic" tf:""`
	// Function null call.
	IsNullCall types.Bool `tfsdk:"is_null_call" tf:""`
	// Name of function, relative to parent schema.
	Name types.String `tfsdk:"name" tf:""`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle types.String `tfsdk:"parameter_style" tf:""`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties types.String `tfsdk:"properties" tf:"optional"`
	// Table function return parameters.
	ReturnParams types.List `tfsdk:"return_params" tf:"optional,object"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody types.String `tfsdk:"routine_body" tf:""`
	// Function body.
	RoutineDefinition types.String `tfsdk:"routine_definition" tf:""`
	// Function dependencies.
	RoutineDependencies types.List `tfsdk:"routine_dependencies" tf:"optional,object"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name" tf:""`
	// Function security type.
	SecurityType types.String `tfsdk:"security_type" tf:""`
	// Specific name of the function; Reserved for future use.
	SpecificName types.String `tfsdk:"specific_name" tf:""`
	// Function SQL data access.
	SqlDataAccess types.String `tfsdk:"sql_data_access" tf:""`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath types.String `tfsdk:"sql_path" tf:"optional"`
}

func (newState *CreateFunction_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFunction_SdkV2) {
}

func (newState *CreateFunction_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateFunction_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFunction.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateFunction_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input_params":         reflect.TypeOf(FunctionParameterInfos_SdkV2{}),
		"return_params":        reflect.TypeOf(FunctionParameterInfos_SdkV2{}),
		"routine_dependencies": reflect.TypeOf(DependencyList_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFunction_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateFunction_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateFunction_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":      types.StringType,
			"comment":           types.StringType,
			"data_type":         types.StringType,
			"external_language": types.StringType,
			"external_name":     types.StringType,
			"full_data_type":    types.StringType,
			"input_params": basetypes.ListType{
				ElemType: FunctionParameterInfos_SdkV2{}.Type(ctx),
			},
			"is_deterministic": types.BoolType,
			"is_null_call":     types.BoolType,
			"name":             types.StringType,
			"parameter_style":  types.StringType,
			"properties":       types.StringType,
			"return_params": basetypes.ListType{
				ElemType: FunctionParameterInfos_SdkV2{}.Type(ctx),
			},
			"routine_body":       types.StringType,
			"routine_definition": types.StringType,
			"routine_dependencies": basetypes.ListType{
				ElemType: DependencyList_SdkV2{}.Type(ctx),
			},
			"schema_name":     types.StringType,
			"security_type":   types.StringType,
			"specific_name":   types.StringType,
			"sql_data_access": types.StringType,
			"sql_path":        types.StringType,
		},
	}
}

// GetInputParams returns the value of the InputParams field in CreateFunction_SdkV2 as
// a FunctionParameterInfos_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFunction_SdkV2) GetInputParams(ctx context.Context) (FunctionParameterInfos_SdkV2, bool) {
	var e FunctionParameterInfos_SdkV2
	if o.InputParams.IsNull() || o.InputParams.IsUnknown() {
		return e, false
	}
	var v []FunctionParameterInfos_SdkV2
	d := o.InputParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInputParams sets the value of the InputParams field in CreateFunction_SdkV2.
func (o *CreateFunction_SdkV2) SetInputParams(ctx context.Context, v FunctionParameterInfos_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["input_params"]
	o.InputParams = types.ListValueMust(t, vs)
}

// GetReturnParams returns the value of the ReturnParams field in CreateFunction_SdkV2 as
// a FunctionParameterInfos_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFunction_SdkV2) GetReturnParams(ctx context.Context) (FunctionParameterInfos_SdkV2, bool) {
	var e FunctionParameterInfos_SdkV2
	if o.ReturnParams.IsNull() || o.ReturnParams.IsUnknown() {
		return e, false
	}
	var v []FunctionParameterInfos_SdkV2
	d := o.ReturnParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetReturnParams sets the value of the ReturnParams field in CreateFunction_SdkV2.
func (o *CreateFunction_SdkV2) SetReturnParams(ctx context.Context, v FunctionParameterInfos_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["return_params"]
	o.ReturnParams = types.ListValueMust(t, vs)
}

// GetRoutineDependencies returns the value of the RoutineDependencies field in CreateFunction_SdkV2 as
// a DependencyList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFunction_SdkV2) GetRoutineDependencies(ctx context.Context) (DependencyList_SdkV2, bool) {
	var e DependencyList_SdkV2
	if o.RoutineDependencies.IsNull() || o.RoutineDependencies.IsUnknown() {
		return e, false
	}
	var v []DependencyList_SdkV2
	d := o.RoutineDependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRoutineDependencies sets the value of the RoutineDependencies field in CreateFunction_SdkV2.
func (o *CreateFunction_SdkV2) SetRoutineDependencies(ctx context.Context, v DependencyList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["routine_dependencies"]
	o.RoutineDependencies = types.ListValueMust(t, vs)
}

type CreateFunctionRequest_SdkV2 struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	FunctionInfo types.List `tfsdk:"function_info" tf:"object"`
}

func (newState *CreateFunctionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFunctionRequest_SdkV2) {
}

func (newState *CreateFunctionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateFunctionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFunctionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateFunctionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"function_info": reflect.TypeOf(CreateFunction_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFunctionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateFunctionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_info": o.FunctionInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateFunctionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_info": basetypes.ListType{
				ElemType: CreateFunction_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFunctionInfo returns the value of the FunctionInfo field in CreateFunctionRequest_SdkV2 as
// a CreateFunction_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateFunctionRequest_SdkV2) GetFunctionInfo(ctx context.Context) (CreateFunction_SdkV2, bool) {
	var e CreateFunction_SdkV2
	if o.FunctionInfo.IsNull() || o.FunctionInfo.IsUnknown() {
		return e, false
	}
	var v []CreateFunction_SdkV2
	d := o.FunctionInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFunctionInfo sets the value of the FunctionInfo field in CreateFunctionRequest_SdkV2.
func (o *CreateFunctionRequest_SdkV2) SetFunctionInfo(ctx context.Context, v CreateFunction_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["function_info"]
	o.FunctionInfo = types.ListValueMust(t, vs)
}

type CreateMetastore_SdkV2 struct {
	// The user-specified name of the metastore.
	Name types.String `tfsdk:"name" tf:""`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// The field can be omitted in the __workspace-level__ __API__ but not in
	// the __account-level__ __API__. If this field is omitted, the region of
	// the workspace receiving the request will be used.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The storage root URL for metastore
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

func (newState *CreateMetastore_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMetastore_SdkV2) {
}

func (newState *CreateMetastore_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateMetastore_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateMetastore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateMetastore_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateMetastore_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateMetastore_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":         o.Name,
			"region":       o.Region,
			"storage_root": o.StorageRoot,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateMetastore_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":         types.StringType,
			"region":       types.StringType,
			"storage_root": types.StringType,
		},
	}
}

type CreateMetastoreAssignment_SdkV2 struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name" tf:""`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:""`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *CreateMetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMetastoreAssignment_SdkV2) {
}

func (newState *CreateMetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateMetastoreAssignment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateMetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateMetastoreAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateMetastoreAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateMetastoreAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_catalog_name": o.DefaultCatalogName,
			"metastore_id":         o.MetastoreId,
			"workspace_id":         o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateMetastoreAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_catalog_name": types.StringType,
			"metastore_id":         types.StringType,
			"workspace_id":         types.Int64Type,
		},
	}
}

type CreateMonitor_SdkV2 struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir types.String `tfsdk:"assets_dir" tf:""`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName types.String `tfsdk:"baseline_table_name" tf:"optional"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics types.List `tfsdk:"custom_metrics" tf:"optional"`
	// The data classification config for the monitor.
	DataClassificationConfig types.List `tfsdk:"data_classification_config" tf:"optional,object"`
	// Configuration for monitoring inference logs.
	InferenceLog types.List `tfsdk:"inference_log" tf:"optional,object"`
	// The notification settings for the monitor.
	Notifications types.List `tfsdk:"notifications" tf:"optional,object"`
	// Schema where output metric tables are created.
	OutputSchemaName types.String `tfsdk:"output_schema_name" tf:""`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard types.Bool `tfsdk:"skip_builtin_dashboard" tf:"optional"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs types.List `tfsdk:"slicing_exprs" tf:"optional"`
	// Configuration for monitoring snapshot tables.
	Snapshot types.List `tfsdk:"snapshot" tf:"optional,object"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries types.List `tfsdk:"time_series" tf:"optional,object"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *CreateMonitor_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMonitor_SdkV2) {
}

func (newState *CreateMonitor_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateMonitor_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateMonitor.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateMonitor_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_metrics":             reflect.TypeOf(MonitorMetric_SdkV2{}),
		"data_classification_config": reflect.TypeOf(MonitorDataClassificationConfig_SdkV2{}),
		"inference_log":              reflect.TypeOf(MonitorInferenceLog_SdkV2{}),
		"notifications":              reflect.TypeOf(MonitorNotifications_SdkV2{}),
		"schedule":                   reflect.TypeOf(MonitorCronSchedule_SdkV2{}),
		"slicing_exprs":              reflect.TypeOf(types.String{}),
		"snapshot":                   reflect.TypeOf(MonitorSnapshot_SdkV2{}),
		"time_series":                reflect.TypeOf(MonitorTimeSeries_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateMonitor_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateMonitor_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateMonitor_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets_dir":          types.StringType,
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: MonitorMetric_SdkV2{}.Type(ctx),
			},
			"data_classification_config": basetypes.ListType{
				ElemType: MonitorDataClassificationConfig_SdkV2{}.Type(ctx),
			},
			"inference_log": basetypes.ListType{
				ElemType: MonitorInferenceLog_SdkV2{}.Type(ctx),
			},
			"notifications": basetypes.ListType{
				ElemType: MonitorNotifications_SdkV2{}.Type(ctx),
			},
			"output_schema_name": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: MonitorCronSchedule_SdkV2{}.Type(ctx),
			},
			"skip_builtin_dashboard": types.BoolType,
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot": basetypes.ListType{
				ElemType: MonitorSnapshot_SdkV2{}.Type(ctx),
			},
			"table_name": types.StringType,
			"time_series": basetypes.ListType{
				ElemType: MonitorTimeSeries_SdkV2{}.Type(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

// GetCustomMetrics returns the value of the CustomMetrics field in CreateMonitor_SdkV2 as
// a slice of MonitorMetric_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor_SdkV2) GetCustomMetrics(ctx context.Context) ([]MonitorMetric_SdkV2, bool) {
	if o.CustomMetrics.IsNull() || o.CustomMetrics.IsUnknown() {
		return nil, false
	}
	var v []MonitorMetric_SdkV2
	d := o.CustomMetrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomMetrics sets the value of the CustomMetrics field in CreateMonitor_SdkV2.
func (o *CreateMonitor_SdkV2) SetCustomMetrics(ctx context.Context, v []MonitorMetric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomMetrics = types.ListValueMust(t, vs)
}

// GetDataClassificationConfig returns the value of the DataClassificationConfig field in CreateMonitor_SdkV2 as
// a MonitorDataClassificationConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor_SdkV2) GetDataClassificationConfig(ctx context.Context) (MonitorDataClassificationConfig_SdkV2, bool) {
	var e MonitorDataClassificationConfig_SdkV2
	if o.DataClassificationConfig.IsNull() || o.DataClassificationConfig.IsUnknown() {
		return e, false
	}
	var v []MonitorDataClassificationConfig_SdkV2
	d := o.DataClassificationConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataClassificationConfig sets the value of the DataClassificationConfig field in CreateMonitor_SdkV2.
func (o *CreateMonitor_SdkV2) SetDataClassificationConfig(ctx context.Context, v MonitorDataClassificationConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_classification_config"]
	o.DataClassificationConfig = types.ListValueMust(t, vs)
}

// GetInferenceLog returns the value of the InferenceLog field in CreateMonitor_SdkV2 as
// a MonitorInferenceLog_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor_SdkV2) GetInferenceLog(ctx context.Context) (MonitorInferenceLog_SdkV2, bool) {
	var e MonitorInferenceLog_SdkV2
	if o.InferenceLog.IsNull() || o.InferenceLog.IsUnknown() {
		return e, false
	}
	var v []MonitorInferenceLog_SdkV2
	d := o.InferenceLog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInferenceLog sets the value of the InferenceLog field in CreateMonitor_SdkV2.
func (o *CreateMonitor_SdkV2) SetInferenceLog(ctx context.Context, v MonitorInferenceLog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inference_log"]
	o.InferenceLog = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in CreateMonitor_SdkV2 as
// a MonitorNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor_SdkV2) GetNotifications(ctx context.Context) (MonitorNotifications_SdkV2, bool) {
	var e MonitorNotifications_SdkV2
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return e, false
	}
	var v []MonitorNotifications_SdkV2
	d := o.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotifications sets the value of the Notifications field in CreateMonitor_SdkV2.
func (o *CreateMonitor_SdkV2) SetNotifications(ctx context.Context, v MonitorNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	o.Notifications = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in CreateMonitor_SdkV2 as
// a MonitorCronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor_SdkV2) GetSchedule(ctx context.Context) (MonitorCronSchedule_SdkV2, bool) {
	var e MonitorCronSchedule_SdkV2
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []MonitorCronSchedule_SdkV2
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in CreateMonitor_SdkV2.
func (o *CreateMonitor_SdkV2) SetSchedule(ctx context.Context, v MonitorCronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}

// GetSlicingExprs returns the value of the SlicingExprs field in CreateMonitor_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor_SdkV2) GetSlicingExprs(ctx context.Context) ([]types.String, bool) {
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

// SetSlicingExprs sets the value of the SlicingExprs field in CreateMonitor_SdkV2.
func (o *CreateMonitor_SdkV2) SetSlicingExprs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["slicing_exprs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SlicingExprs = types.ListValueMust(t, vs)
}

// GetSnapshot returns the value of the Snapshot field in CreateMonitor_SdkV2 as
// a MonitorSnapshot_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor_SdkV2) GetSnapshot(ctx context.Context) (MonitorSnapshot_SdkV2, bool) {
	var e MonitorSnapshot_SdkV2
	if o.Snapshot.IsNull() || o.Snapshot.IsUnknown() {
		return e, false
	}
	var v []MonitorSnapshot_SdkV2
	d := o.Snapshot.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSnapshot sets the value of the Snapshot field in CreateMonitor_SdkV2.
func (o *CreateMonitor_SdkV2) SetSnapshot(ctx context.Context, v MonitorSnapshot_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["snapshot"]
	o.Snapshot = types.ListValueMust(t, vs)
}

// GetTimeSeries returns the value of the TimeSeries field in CreateMonitor_SdkV2 as
// a MonitorTimeSeries_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateMonitor_SdkV2) GetTimeSeries(ctx context.Context) (MonitorTimeSeries_SdkV2, bool) {
	var e MonitorTimeSeries_SdkV2
	if o.TimeSeries.IsNull() || o.TimeSeries.IsUnknown() {
		return e, false
	}
	var v []MonitorTimeSeries_SdkV2
	d := o.TimeSeries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTimeSeries sets the value of the TimeSeries field in CreateMonitor_SdkV2.
func (o *CreateMonitor_SdkV2) SetTimeSeries(ctx context.Context, v MonitorTimeSeries_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["time_series"]
	o.TimeSeries = types.ListValueMust(t, vs)
}

// Create an Online Table
type CreateOnlineTableRequest_SdkV2 struct {
	// Online Table information.
	Table types.List `tfsdk:"table" tf:"optional,object"`
}

func (newState *CreateOnlineTableRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateOnlineTableRequest_SdkV2) {
}

func (newState *CreateOnlineTableRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateOnlineTableRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateOnlineTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateOnlineTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table": reflect.TypeOf(OnlineTable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOnlineTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateOnlineTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table": o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateOnlineTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table": basetypes.ListType{
				ElemType: OnlineTable_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTable returns the value of the Table field in CreateOnlineTableRequest_SdkV2 as
// a OnlineTable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateOnlineTableRequest_SdkV2) GetTable(ctx context.Context) (OnlineTable_SdkV2, bool) {
	var e OnlineTable_SdkV2
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []OnlineTable_SdkV2
	d := o.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in CreateOnlineTableRequest_SdkV2.
func (o *CreateOnlineTableRequest_SdkV2) SetTable(ctx context.Context, v OnlineTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	o.Table = types.ListValueMust(t, vs)
}

type CreateRegisteredModelRequest_SdkV2 struct {
	// The name of the catalog where the schema and the registered model reside
	CatalogName types.String `tfsdk:"catalog_name" tf:""`
	// The comment attached to the registered model
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the registered model
	Name types.String `tfsdk:"name" tf:""`
	// The name of the schema where the registered model resides
	SchemaName types.String `tfsdk:"schema_name" tf:""`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
}

func (newState *CreateRegisteredModelRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRegisteredModelRequest_SdkV2) {
}

func (newState *CreateRegisteredModelRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateRegisteredModelRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRegisteredModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateRegisteredModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRegisteredModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateRegisteredModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateRegisteredModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type CreateResponse_SdkV2 struct {
}

func (newState *CreateResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateResponse_SdkV2) {
}

func (newState *CreateResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CreateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CreateSchema_SdkV2 struct {
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:""`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of schema, relative to parent catalog.
	Name types.String `tfsdk:"name" tf:""`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties" tf:"optional"`
	// Storage root URL for managed tables within schema.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

func (newState *CreateSchema_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateSchema_SdkV2) {
}

func (newState *CreateSchema_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateSchema_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSchema.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateSchema_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSchema_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateSchema_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateSchema_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetProperties returns the value of the Properties field in CreateSchema_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateSchema_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in CreateSchema_SdkV2.
func (o *CreateSchema_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type CreateStorageCredential_SdkV2 struct {
	// The AWS IAM role configuration.
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// The Cloudflare API token configuration.
	CloudflareApiToken types.List `tfsdk:"cloudflare_api_token" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount types.List `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// The credential name. The name must be unique within the metastore.
	Name types.String `tfsdk:"name" tf:""`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Supplying true to this argument skips validation of the created
	// credential.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
}

func (newState *CreateStorageCredential_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateStorageCredential_SdkV2) {
}

func (newState *CreateStorageCredential_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateStorageCredential_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateStorageCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateStorageCredential_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRoleRequest_SdkV2{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentityRequest_SdkV2{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal_SdkV2{}),
		"cloudflare_api_token":           reflect.TypeOf(CloudflareApiToken_SdkV2{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccountRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateStorageCredential_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateStorageCredential_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateStorageCredential_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRoleRequest_SdkV2{}.Type(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentityRequest_SdkV2{}.Type(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal_SdkV2{}.Type(ctx),
			},
			"cloudflare_api_token": basetypes.ListType{
				ElemType: CloudflareApiToken_SdkV2{}.Type(ctx),
			},
			"comment": types.StringType,
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccountRequest_SdkV2{}.Type(ctx),
			},
			"name":            types.StringType,
			"read_only":       types.BoolType,
			"skip_validation": types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in CreateStorageCredential_SdkV2 as
// a AwsIamRoleRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageCredential_SdkV2) GetAwsIamRole(ctx context.Context) (AwsIamRoleRequest_SdkV2, bool) {
	var e AwsIamRoleRequest_SdkV2
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRoleRequest_SdkV2
	d := o.AwsIamRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsIamRole sets the value of the AwsIamRole field in CreateStorageCredential_SdkV2.
func (o *CreateStorageCredential_SdkV2) SetAwsIamRole(ctx context.Context, v AwsIamRoleRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_iam_role"]
	o.AwsIamRole = types.ListValueMust(t, vs)
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in CreateStorageCredential_SdkV2 as
// a AzureManagedIdentityRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageCredential_SdkV2) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentityRequest_SdkV2, bool) {
	var e AzureManagedIdentityRequest_SdkV2
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentityRequest_SdkV2
	d := o.AzureManagedIdentity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in CreateStorageCredential_SdkV2.
func (o *CreateStorageCredential_SdkV2) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentityRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_managed_identity"]
	o.AzureManagedIdentity = types.ListValueMust(t, vs)
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in CreateStorageCredential_SdkV2 as
// a AzureServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageCredential_SdkV2) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal_SdkV2, bool) {
	var e AzureServicePrincipal_SdkV2
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal_SdkV2
	d := o.AzureServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in CreateStorageCredential_SdkV2.
func (o *CreateStorageCredential_SdkV2) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_service_principal"]
	o.AzureServicePrincipal = types.ListValueMust(t, vs)
}

// GetCloudflareApiToken returns the value of the CloudflareApiToken field in CreateStorageCredential_SdkV2 as
// a CloudflareApiToken_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageCredential_SdkV2) GetCloudflareApiToken(ctx context.Context) (CloudflareApiToken_SdkV2, bool) {
	var e CloudflareApiToken_SdkV2
	if o.CloudflareApiToken.IsNull() || o.CloudflareApiToken.IsUnknown() {
		return e, false
	}
	var v []CloudflareApiToken_SdkV2
	d := o.CloudflareApiToken.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudflareApiToken sets the value of the CloudflareApiToken field in CreateStorageCredential_SdkV2.
func (o *CreateStorageCredential_SdkV2) SetCloudflareApiToken(ctx context.Context, v CloudflareApiToken_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cloudflare_api_token"]
	o.CloudflareApiToken = types.ListValueMust(t, vs)
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in CreateStorageCredential_SdkV2 as
// a DatabricksGcpServiceAccountRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageCredential_SdkV2) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccountRequest_SdkV2, bool) {
	var e DatabricksGcpServiceAccountRequest_SdkV2
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccountRequest_SdkV2
	d := o.DatabricksGcpServiceAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in CreateStorageCredential_SdkV2.
func (o *CreateStorageCredential_SdkV2) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccountRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["databricks_gcp_service_account"]
	o.DatabricksGcpServiceAccount = types.ListValueMust(t, vs)
}

type CreateTableConstraint_SdkV2 struct {
	// A table constraint, as defined by *one* of the following fields being
	// set: __primary_key_constraint__, __foreign_key_constraint__,
	// __named_table_constraint__.
	Constraint types.List `tfsdk:"constraint" tf:"object"`
	// The full name of the table referenced by the constraint.
	FullNameArg types.String `tfsdk:"full_name_arg" tf:""`
}

func (newState *CreateTableConstraint_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTableConstraint_SdkV2) {
}

func (newState *CreateTableConstraint_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateTableConstraint_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTableConstraint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateTableConstraint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"constraint": reflect.TypeOf(TableConstraint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTableConstraint_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateTableConstraint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"constraint":    o.Constraint,
			"full_name_arg": o.FullNameArg,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTableConstraint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"constraint": basetypes.ListType{
				ElemType: TableConstraint_SdkV2{}.Type(ctx),
			},
			"full_name_arg": types.StringType,
		},
	}
}

// GetConstraint returns the value of the Constraint field in CreateTableConstraint_SdkV2 as
// a TableConstraint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateTableConstraint_SdkV2) GetConstraint(ctx context.Context) (TableConstraint_SdkV2, bool) {
	var e TableConstraint_SdkV2
	if o.Constraint.IsNull() || o.Constraint.IsUnknown() {
		return e, false
	}
	var v []TableConstraint_SdkV2
	d := o.Constraint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConstraint sets the value of the Constraint field in CreateTableConstraint_SdkV2.
func (o *CreateTableConstraint_SdkV2) SetConstraint(ctx context.Context, v TableConstraint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["constraint"]
	o.Constraint = types.ListValueMust(t, vs)
}

type CreateVolumeRequestContent_SdkV2 struct {
	// The name of the catalog where the schema and the volume are
	CatalogName types.String `tfsdk:"catalog_name" tf:""`
	// The comment attached to the volume
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the volume
	Name types.String `tfsdk:"name" tf:""`
	// The name of the schema where the volume is
	SchemaName types.String `tfsdk:"schema_name" tf:""`
	// The storage location on the cloud
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`

	VolumeType types.String `tfsdk:"volume_type" tf:""`
}

func (newState *CreateVolumeRequestContent_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVolumeRequestContent_SdkV2) {
}

func (newState *CreateVolumeRequestContent_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateVolumeRequestContent_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVolumeRequestContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVolumeRequestContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVolumeRequestContent_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateVolumeRequestContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateVolumeRequestContent_SdkV2) Type(ctx context.Context) attr.Type {
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

type CredentialInfo_SdkV2 struct {
	// The AWS IAM role configuration
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this credential was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of credential creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount types.List `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// The full name of the credential.
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// The unique identifier of the credential.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Unique identifier of the parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of current owner of credential.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Indicates the purpose of the credential.
	Purpose types.String `tfsdk:"purpose" tf:"optional"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the credential.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	UsedForManagedStorage types.Bool `tfsdk:"used_for_managed_storage" tf:"optional"`
}

func (newState *CredentialInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CredentialInfo_SdkV2) {
}

func (newState *CredentialInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState CredentialInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CredentialInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CredentialInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRole_SdkV2{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentity_SdkV2{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal_SdkV2{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccount_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CredentialInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o CredentialInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CredentialInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRole_SdkV2{}.Type(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentity_SdkV2{}.Type(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal_SdkV2{}.Type(ctx),
			},
			"comment":    types.StringType,
			"created_at": types.Int64Type,
			"created_by": types.StringType,
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccount_SdkV2{}.Type(ctx),
			},
			"full_name":                types.StringType,
			"id":                       types.StringType,
			"isolation_mode":           types.StringType,
			"metastore_id":             types.StringType,
			"name":                     types.StringType,
			"owner":                    types.StringType,
			"purpose":                  types.StringType,
			"read_only":                types.BoolType,
			"updated_at":               types.Int64Type,
			"updated_by":               types.StringType,
			"used_for_managed_storage": types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in CredentialInfo_SdkV2 as
// a AwsIamRole_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CredentialInfo_SdkV2) GetAwsIamRole(ctx context.Context) (AwsIamRole_SdkV2, bool) {
	var e AwsIamRole_SdkV2
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRole_SdkV2
	d := o.AwsIamRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsIamRole sets the value of the AwsIamRole field in CredentialInfo_SdkV2.
func (o *CredentialInfo_SdkV2) SetAwsIamRole(ctx context.Context, v AwsIamRole_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_iam_role"]
	o.AwsIamRole = types.ListValueMust(t, vs)
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in CredentialInfo_SdkV2 as
// a AzureManagedIdentity_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CredentialInfo_SdkV2) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentity_SdkV2, bool) {
	var e AzureManagedIdentity_SdkV2
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentity_SdkV2
	d := o.AzureManagedIdentity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in CredentialInfo_SdkV2.
func (o *CredentialInfo_SdkV2) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentity_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_managed_identity"]
	o.AzureManagedIdentity = types.ListValueMust(t, vs)
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in CredentialInfo_SdkV2 as
// a AzureServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CredentialInfo_SdkV2) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal_SdkV2, bool) {
	var e AzureServicePrincipal_SdkV2
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal_SdkV2
	d := o.AzureServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in CredentialInfo_SdkV2.
func (o *CredentialInfo_SdkV2) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_service_principal"]
	o.AzureServicePrincipal = types.ListValueMust(t, vs)
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in CredentialInfo_SdkV2 as
// a DatabricksGcpServiceAccount_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CredentialInfo_SdkV2) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccount_SdkV2, bool) {
	var e DatabricksGcpServiceAccount_SdkV2
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccount_SdkV2
	d := o.DatabricksGcpServiceAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in CredentialInfo_SdkV2.
func (o *CredentialInfo_SdkV2) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccount_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["databricks_gcp_service_account"]
	o.DatabricksGcpServiceAccount = types.ListValueMust(t, vs)
}

type CredentialValidationResult_SdkV2 struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message types.String `tfsdk:"message" tf:"optional"`
	// The results of the tested operation.
	Result types.String `tfsdk:"result" tf:"optional"`
}

func (newState *CredentialValidationResult_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CredentialValidationResult_SdkV2) {
}

func (newState *CredentialValidationResult_SdkV2) SyncEffectiveFieldsDuringRead(existingState CredentialValidationResult_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CredentialValidationResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CredentialValidationResult_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CredentialValidationResult_SdkV2
// only implements ToObjectValue() and Type().
func (o CredentialValidationResult_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
			"result":  o.Result,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CredentialValidationResult_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"result":  types.StringType,
		},
	}
}

// Currently assigned workspaces
type CurrentWorkspaceBindings_SdkV2 struct {
	// A list of workspace IDs.
	Workspaces types.List `tfsdk:"workspaces" tf:"optional"`
}

func (newState *CurrentWorkspaceBindings_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CurrentWorkspaceBindings_SdkV2) {
}

func (newState *CurrentWorkspaceBindings_SdkV2) SyncEffectiveFieldsDuringRead(existingState CurrentWorkspaceBindings_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CurrentWorkspaceBindings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CurrentWorkspaceBindings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspaces": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CurrentWorkspaceBindings_SdkV2
// only implements ToObjectValue() and Type().
func (o CurrentWorkspaceBindings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspaces": o.Workspaces,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CurrentWorkspaceBindings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspaces": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetWorkspaces returns the value of the Workspaces field in CurrentWorkspaceBindings_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CurrentWorkspaceBindings_SdkV2) GetWorkspaces(ctx context.Context) ([]types.Int64, bool) {
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

// SetWorkspaces sets the value of the Workspaces field in CurrentWorkspaceBindings_SdkV2.
func (o *CurrentWorkspaceBindings_SdkV2) SetWorkspaces(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspaces"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Workspaces = types.ListValueMust(t, vs)
}

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccount_SdkV2 struct {
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database
	CredentialId types.String `tfsdk:"credential_id" tf:"optional"`
	// The email of the service account.
	Email types.String `tfsdk:"email" tf:"optional"`
	// The ID that represents the private key for this Service Account
	PrivateKeyId types.String `tfsdk:"private_key_id" tf:"optional"`
}

func (newState *DatabricksGcpServiceAccount_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksGcpServiceAccount_SdkV2) {
}

func (newState *DatabricksGcpServiceAccount_SdkV2) SyncEffectiveFieldsDuringRead(existingState DatabricksGcpServiceAccount_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksGcpServiceAccount.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabricksGcpServiceAccount_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksGcpServiceAccount_SdkV2
// only implements ToObjectValue() and Type().
func (o DatabricksGcpServiceAccount_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":  o.CredentialId,
			"email":          o.Email,
			"private_key_id": o.PrivateKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabricksGcpServiceAccount_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id":  types.StringType,
			"email":          types.StringType,
			"private_key_id": types.StringType,
		},
	}
}

type DatabricksGcpServiceAccountRequest_SdkV2 struct {
}

func (newState *DatabricksGcpServiceAccountRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksGcpServiceAccountRequest_SdkV2) {
}

func (newState *DatabricksGcpServiceAccountRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DatabricksGcpServiceAccountRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksGcpServiceAccountRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabricksGcpServiceAccountRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksGcpServiceAccountRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DatabricksGcpServiceAccountRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DatabricksGcpServiceAccountRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DatabricksGcpServiceAccountResponse_SdkV2 struct {
	// The Databricks internal ID that represents this service account. This is
	// an output-only field.
	CredentialId types.String `tfsdk:"credential_id" tf:"optional"`
	// The email of the service account. This is an output-only field.
	Email types.String `tfsdk:"email" tf:"optional"`
}

func (newState *DatabricksGcpServiceAccountResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksGcpServiceAccountResponse_SdkV2) {
}

func (newState *DatabricksGcpServiceAccountResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DatabricksGcpServiceAccountResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksGcpServiceAccountResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabricksGcpServiceAccountResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksGcpServiceAccountResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DatabricksGcpServiceAccountResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
			"email":         o.Email,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabricksGcpServiceAccountResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.StringType,
			"email":         types.StringType,
		},
	}
}

// Delete a metastore assignment
type DeleteAccountMetastoreAssignmentRequest_SdkV2 struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *DeleteAccountMetastoreAssignmentRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountMetastoreAssignmentRequest_SdkV2) {
}

func (newState *DeleteAccountMetastoreAssignmentRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteAccountMetastoreAssignmentRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountMetastoreAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountMetastoreAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountMetastoreAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAccountMetastoreAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountMetastoreAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// Delete a metastore
type DeleteAccountMetastoreRequest_SdkV2 struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force types.Bool `tfsdk:"-"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *DeleteAccountMetastoreRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountMetastoreRequest_SdkV2) {
}

func (newState *DeleteAccountMetastoreRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteAccountMetastoreRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountMetastoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountMetastoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountMetastoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAccountMetastoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force":        o.Force,
			"metastore_id": o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountMetastoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force":        types.BoolType,
			"metastore_id": types.StringType,
		},
	}
}

// Delete a storage credential
type DeleteAccountStorageCredentialRequest_SdkV2 struct {
	// Force deletion even if the Storage Credential is not empty. Default is
	// false.
	Force types.Bool `tfsdk:"-"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Name of the storage credential.
	StorageCredentialName types.String `tfsdk:"-"`
}

func (newState *DeleteAccountStorageCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountStorageCredentialRequest_SdkV2) {
}

func (newState *DeleteAccountStorageCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteAccountStorageCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountStorageCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountStorageCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountStorageCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAccountStorageCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force":                   o.Force,
			"metastore_id":            o.MetastoreId,
			"storage_credential_name": o.StorageCredentialName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountStorageCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force":                   types.BoolType,
			"metastore_id":            types.StringType,
			"storage_credential_name": types.StringType,
		},
	}
}

// Delete a Registered Model Alias
type DeleteAliasRequest_SdkV2 struct {
	// The name of the alias
	Alias types.String `tfsdk:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
}

func (newState *DeleteAliasRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAliasRequest_SdkV2) {
}

func (newState *DeleteAliasRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteAliasRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAliasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAliasRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAliasRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAliasRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":     o.Alias,
			"full_name": o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAliasRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":     types.StringType,
			"full_name": types.StringType,
		},
	}
}

type DeleteAliasResponse_SdkV2 struct {
}

func (newState *DeleteAliasResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAliasResponse_SdkV2) {
}

func (newState *DeleteAliasResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteAliasResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAliasResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAliasResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAliasResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAliasResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAliasResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a catalog
type DeleteCatalogRequest_SdkV2 struct {
	// Force deletion even if the catalog is not empty.
	Force types.Bool `tfsdk:"-"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteCatalogRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCatalogRequest_SdkV2) {
}

func (newState *DeleteCatalogRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteCatalogRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
		},
	}
}

// Delete a connection
type DeleteConnectionRequest_SdkV2 struct {
	// The name of the connection to be deleted.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteConnectionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteConnectionRequest_SdkV2) {
}

func (newState *DeleteConnectionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteConnectionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteConnectionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteConnectionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteConnectionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteConnectionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteConnectionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Delete a credential
type DeleteCredentialRequest_SdkV2 struct {
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force types.Bool `tfsdk:"-"`
	// Name of the credential.
	NameArg types.String `tfsdk:"-"`
}

func (newState *DeleteCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCredentialRequest_SdkV2) {
}

func (newState *DeleteCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force":    o.Force,
			"name_arg": o.NameArg,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force":    types.BoolType,
			"name_arg": types.StringType,
		},
	}
}

type DeleteCredentialResponse_SdkV2 struct {
}

func (newState *DeleteCredentialResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCredentialResponse_SdkV2) {
}

func (newState *DeleteCredentialResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteCredentialResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCredentialResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCredentialResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete an external location
type DeleteExternalLocationRequest_SdkV2 struct {
	// Force deletion even if there are dependent external tables or mounts.
	Force types.Bool `tfsdk:"-"`
	// Name of the external location.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteExternalLocationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExternalLocationRequest_SdkV2) {
}

func (newState *DeleteExternalLocationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteExternalLocationRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExternalLocationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExternalLocationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExternalLocationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteExternalLocationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExternalLocationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
		},
	}
}

// Delete a function
type DeleteFunctionRequest_SdkV2 struct {
	// Force deletion even if the function is notempty.
	Force types.Bool `tfsdk:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteFunctionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteFunctionRequest_SdkV2) {
}

func (newState *DeleteFunctionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteFunctionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFunctionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteFunctionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFunctionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteFunctionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteFunctionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
		},
	}
}

// Delete a metastore
type DeleteMetastoreRequest_SdkV2 struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force types.Bool `tfsdk:"-"`
	// Unique ID of the metastore.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteMetastoreRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteMetastoreRequest_SdkV2) {
}

func (newState *DeleteMetastoreRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteMetastoreRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteMetastoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteMetastoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteMetastoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteMetastoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"id":    o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteMetastoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"id":    types.StringType,
		},
	}
}

// Delete a Model Version
type DeleteModelVersionRequest_SdkV2 struct {
	// The three-level (fully qualified) name of the model version
	FullName types.String `tfsdk:"-"`
	// The integer version number of the model version
	Version types.Int64 `tfsdk:"-"`
}

func (newState *DeleteModelVersionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionRequest_SdkV2) {
}

func (newState *DeleteModelVersionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": o.FullName,
			"version":   o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
			"version":   types.Int64Type,
		},
	}
}

// Delete an Online Table
type DeleteOnlineTableRequest_SdkV2 struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteOnlineTableRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteOnlineTableRequest_SdkV2) {
}

func (newState *DeleteOnlineTableRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteOnlineTableRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteOnlineTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteOnlineTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteOnlineTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteOnlineTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteOnlineTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Delete a table monitor
type DeleteQualityMonitorRequest_SdkV2 struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *DeleteQualityMonitorRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteQualityMonitorRequest_SdkV2) {
}

func (newState *DeleteQualityMonitorRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteQualityMonitorRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteQualityMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteQualityMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteQualityMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteQualityMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name": types.StringType,
		},
	}
}

// Delete a Registered Model
type DeleteRegisteredModelRequest_SdkV2 struct {
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
}

func (newState *DeleteRegisteredModelRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRegisteredModelRequest_SdkV2) {
}

func (newState *DeleteRegisteredModelRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteRegisteredModelRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRegisteredModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRegisteredModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRegisteredModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRegisteredModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRegisteredModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
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

// Delete a schema
type DeleteSchemaRequest_SdkV2 struct {
	// Force deletion even if the schema is not empty.
	Force types.Bool `tfsdk:"-"`
	// Full name of the schema.
	FullName types.String `tfsdk:"-"`
}

func (newState *DeleteSchemaRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSchemaRequest_SdkV2) {
}

func (newState *DeleteSchemaRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteSchemaRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSchemaRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSchemaRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSchemaRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteSchemaRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force":     o.Force,
			"full_name": o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSchemaRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force":     types.BoolType,
			"full_name": types.StringType,
		},
	}
}

// Delete a credential
type DeleteStorageCredentialRequest_SdkV2 struct {
	// Force deletion even if there are dependent external locations or external
	// tables.
	Force types.Bool `tfsdk:"-"`
	// Name of the storage credential.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteStorageCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteStorageCredentialRequest_SdkV2) {
}

func (newState *DeleteStorageCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteStorageCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteStorageCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteStorageCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteStorageCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteStorageCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteStorageCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
		},
	}
}

// Delete a table constraint
type DeleteTableConstraintRequest_SdkV2 struct {
	// If true, try deleting all child constraints of the current constraint. If
	// false, reject this operation if the current constraint has any child
	// constraints.
	Cascade types.Bool `tfsdk:"-"`
	// The name of the constraint to delete.
	ConstraintName types.String `tfsdk:"-"`
	// Full name of the table referenced by the constraint.
	FullName types.String `tfsdk:"-"`
}

func (newState *DeleteTableConstraintRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTableConstraintRequest_SdkV2) {
}

func (newState *DeleteTableConstraintRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteTableConstraintRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTableConstraintRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTableConstraintRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTableConstraintRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteTableConstraintRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cascade":         o.Cascade,
			"constraint_name": o.ConstraintName,
			"full_name":       o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTableConstraintRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cascade":         types.BoolType,
			"constraint_name": types.StringType,
			"full_name":       types.StringType,
		},
	}
}

// Delete a table
type DeleteTableRequest_SdkV2 struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`
}

func (newState *DeleteTableRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTableRequest_SdkV2) {
}

func (newState *DeleteTableRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteTableRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
		},
	}
}

// Delete a Volume
type DeleteVolumeRequest_SdkV2 struct {
	// The three-level (fully qualified) name of the volume
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteVolumeRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteVolumeRequest_SdkV2) {
}

func (newState *DeleteVolumeRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteVolumeRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteVolumeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteVolumeRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteVolumeRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteVolumeRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteVolumeRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Properties pertaining to the current state of the delta table as given by the
// commit server. This does not contain **delta.*** (input) properties in
// __TableInfo.properties__.
type DeltaRuntimePropertiesKvPairs_SdkV2 struct {
	// A map of key-value properties attached to the securable.
	DeltaRuntimeProperties types.Map `tfsdk:"delta_runtime_properties" tf:""`
}

func (newState *DeltaRuntimePropertiesKvPairs_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaRuntimePropertiesKvPairs_SdkV2) {
}

func (newState *DeltaRuntimePropertiesKvPairs_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeltaRuntimePropertiesKvPairs_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaRuntimePropertiesKvPairs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeltaRuntimePropertiesKvPairs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_runtime_properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaRuntimePropertiesKvPairs_SdkV2
// only implements ToObjectValue() and Type().
func (o DeltaRuntimePropertiesKvPairs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_runtime_properties": o.DeltaRuntimeProperties,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaRuntimePropertiesKvPairs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_runtime_properties": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDeltaRuntimeProperties returns the value of the DeltaRuntimeProperties field in DeltaRuntimePropertiesKvPairs_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *DeltaRuntimePropertiesKvPairs_SdkV2) GetDeltaRuntimeProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetDeltaRuntimeProperties sets the value of the DeltaRuntimeProperties field in DeltaRuntimePropertiesKvPairs_SdkV2.
func (o *DeltaRuntimePropertiesKvPairs_SdkV2) SetDeltaRuntimeProperties(ctx context.Context, v map[string]types.String) {
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
type Dependency_SdkV2 struct {
	// A function that is dependent on a SQL object.
	Function types.List `tfsdk:"function" tf:"optional,object"`
	// A table that is dependent on a SQL object.
	Table types.List `tfsdk:"table" tf:"optional,object"`
}

func (newState *Dependency_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dependency_SdkV2) {
}

func (newState *Dependency_SdkV2) SyncEffectiveFieldsDuringRead(existingState Dependency_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Dependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Dependency_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"function": reflect.TypeOf(FunctionDependency_SdkV2{}),
		"table":    reflect.TypeOf(TableDependency_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dependency_SdkV2
// only implements ToObjectValue() and Type().
func (o Dependency_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function": o.Function,
			"table":    o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Dependency_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function": basetypes.ListType{
				ElemType: FunctionDependency_SdkV2{}.Type(ctx),
			},
			"table": basetypes.ListType{
				ElemType: TableDependency_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFunction returns the value of the Function field in Dependency_SdkV2 as
// a FunctionDependency_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Dependency_SdkV2) GetFunction(ctx context.Context) (FunctionDependency_SdkV2, bool) {
	var e FunctionDependency_SdkV2
	if o.Function.IsNull() || o.Function.IsUnknown() {
		return e, false
	}
	var v []FunctionDependency_SdkV2
	d := o.Function.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFunction sets the value of the Function field in Dependency_SdkV2.
func (o *Dependency_SdkV2) SetFunction(ctx context.Context, v FunctionDependency_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["function"]
	o.Function = types.ListValueMust(t, vs)
}

// GetTable returns the value of the Table field in Dependency_SdkV2 as
// a TableDependency_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Dependency_SdkV2) GetTable(ctx context.Context) (TableDependency_SdkV2, bool) {
	var e TableDependency_SdkV2
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []TableDependency_SdkV2
	d := o.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in Dependency_SdkV2.
func (o *Dependency_SdkV2) SetTable(ctx context.Context, v TableDependency_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	o.Table = types.ListValueMust(t, vs)
}

// A list of dependencies.
type DependencyList_SdkV2 struct {
	// Array of dependencies.
	Dependencies types.List `tfsdk:"dependencies" tf:"optional"`
}

func (newState *DependencyList_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DependencyList_SdkV2) {
}

func (newState *DependencyList_SdkV2) SyncEffectiveFieldsDuringRead(existingState DependencyList_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DependencyList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DependencyList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dependencies": reflect.TypeOf(Dependency_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DependencyList_SdkV2
// only implements ToObjectValue() and Type().
func (o DependencyList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dependencies": o.Dependencies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DependencyList_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dependencies": basetypes.ListType{
				ElemType: Dependency_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDependencies returns the value of the Dependencies field in DependencyList_SdkV2 as
// a slice of Dependency_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *DependencyList_SdkV2) GetDependencies(ctx context.Context) ([]Dependency_SdkV2, bool) {
	if o.Dependencies.IsNull() || o.Dependencies.IsUnknown() {
		return nil, false
	}
	var v []Dependency_SdkV2
	d := o.Dependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDependencies sets the value of the Dependencies field in DependencyList_SdkV2.
func (o *DependencyList_SdkV2) SetDependencies(ctx context.Context, v []Dependency_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dependencies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Dependencies = types.ListValueMust(t, vs)
}

// Disable a system schema
type DisableRequest_SdkV2 struct {
	// The metastore ID under which the system schema lives.
	MetastoreId types.String `tfsdk:"-"`
	// Full name of the system schema.
	SchemaName types.String `tfsdk:"-"`
}

func (newState *DisableRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableRequest_SdkV2) {
}

func (newState *DisableRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DisableRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DisableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DisableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DisableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
			"schema_name":  o.SchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DisableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
			"schema_name":  types.StringType,
		},
	}
}

type DisableResponse_SdkV2 struct {
}

func (newState *DisableResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableResponse_SdkV2) {
}

func (newState *DisableResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DisableResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DisableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DisableResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DisableResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DisableResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EffectivePermissionsList_SdkV2 struct {
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments" tf:"optional"`
}

func (newState *EffectivePermissionsList_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePermissionsList_SdkV2) {
}

func (newState *EffectivePermissionsList_SdkV2) SyncEffectiveFieldsDuringRead(existingState EffectivePermissionsList_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EffectivePermissionsList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EffectivePermissionsList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(EffectivePrivilegeAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EffectivePermissionsList_SdkV2
// only implements ToObjectValue() and Type().
func (o EffectivePermissionsList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"privilege_assignments": o.PrivilegeAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EffectivePermissionsList_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: EffectivePrivilegeAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in EffectivePermissionsList_SdkV2 as
// a slice of EffectivePrivilegeAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EffectivePermissionsList_SdkV2) GetPrivilegeAssignments(ctx context.Context) ([]EffectivePrivilegeAssignment_SdkV2, bool) {
	if o.PrivilegeAssignments.IsNull() || o.PrivilegeAssignments.IsUnknown() {
		return nil, false
	}
	var v []EffectivePrivilegeAssignment_SdkV2
	d := o.PrivilegeAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in EffectivePermissionsList_SdkV2.
func (o *EffectivePermissionsList_SdkV2) SetPrivilegeAssignments(ctx context.Context, v []EffectivePrivilegeAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrivilegeAssignments = types.ListValueMust(t, vs)
}

type EffectivePredictiveOptimizationFlag_SdkV2 struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromName types.String `tfsdk:"inherited_from_name" tf:"optional"`
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromType types.String `tfsdk:"inherited_from_type" tf:"optional"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	Value types.String `tfsdk:"value" tf:""`
}

func (newState *EffectivePredictiveOptimizationFlag_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePredictiveOptimizationFlag_SdkV2) {
}

func (newState *EffectivePredictiveOptimizationFlag_SdkV2) SyncEffectiveFieldsDuringRead(existingState EffectivePredictiveOptimizationFlag_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EffectivePredictiveOptimizationFlag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EffectivePredictiveOptimizationFlag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EffectivePredictiveOptimizationFlag_SdkV2
// only implements ToObjectValue() and Type().
func (o EffectivePredictiveOptimizationFlag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited_from_name": o.InheritedFromName,
			"inherited_from_type": o.InheritedFromType,
			"value":               o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EffectivePredictiveOptimizationFlag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited_from_name": types.StringType,
			"inherited_from_type": types.StringType,
			"value":               types.StringType,
		},
	}
}

type EffectivePrivilege_SdkV2 struct {
	// The full name of the object that conveys this privilege via inheritance.
	// This field is omitted when privilege is not inherited (it's assigned to
	// the securable itself).
	InheritedFromName types.String `tfsdk:"inherited_from_name" tf:"optional"`
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	InheritedFromType types.String `tfsdk:"inherited_from_type" tf:"optional"`
	// The privilege assigned to the principal.
	Privilege types.String `tfsdk:"privilege" tf:"optional"`
}

func (newState *EffectivePrivilege_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePrivilege_SdkV2) {
}

func (newState *EffectivePrivilege_SdkV2) SyncEffectiveFieldsDuringRead(existingState EffectivePrivilege_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EffectivePrivilege.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EffectivePrivilege_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EffectivePrivilege_SdkV2
// only implements ToObjectValue() and Type().
func (o EffectivePrivilege_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited_from_name": o.InheritedFromName,
			"inherited_from_type": o.InheritedFromType,
			"privilege":           o.Privilege,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EffectivePrivilege_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited_from_name": types.StringType,
			"inherited_from_type": types.StringType,
			"privilege":           types.StringType,
		},
	}
}

type EffectivePrivilegeAssignment_SdkV2 struct {
	// The principal (user email address or group name).
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	Privileges types.List `tfsdk:"privileges" tf:"optional"`
}

func (newState *EffectivePrivilegeAssignment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePrivilegeAssignment_SdkV2) {
}

func (newState *EffectivePrivilegeAssignment_SdkV2) SyncEffectiveFieldsDuringRead(existingState EffectivePrivilegeAssignment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EffectivePrivilegeAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EffectivePrivilegeAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privileges": reflect.TypeOf(EffectivePrivilege_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EffectivePrivilegeAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o EffectivePrivilegeAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal":  o.Principal,
			"privileges": o.Privileges,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EffectivePrivilegeAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"privileges": basetypes.ListType{
				ElemType: EffectivePrivilege_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPrivileges returns the value of the Privileges field in EffectivePrivilegeAssignment_SdkV2 as
// a slice of EffectivePrivilege_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EffectivePrivilegeAssignment_SdkV2) GetPrivileges(ctx context.Context) ([]EffectivePrivilege_SdkV2, bool) {
	if o.Privileges.IsNull() || o.Privileges.IsUnknown() {
		return nil, false
	}
	var v []EffectivePrivilege_SdkV2
	d := o.Privileges.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivileges sets the value of the Privileges field in EffectivePrivilegeAssignment_SdkV2.
func (o *EffectivePrivilegeAssignment_SdkV2) SetPrivileges(ctx context.Context, v []EffectivePrivilege_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privileges"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Privileges = types.ListValueMust(t, vs)
}

// Enable a system schema
type EnableRequest_SdkV2 struct {
	// The metastore ID under which the system schema lives.
	MetastoreId types.String `tfsdk:"-"`
	// Full name of the system schema.
	SchemaName types.String `tfsdk:"-"`
}

func (newState *EnableRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnableRequest_SdkV2) {
}

func (newState *EnableRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState EnableRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o EnableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
			"schema_name":  o.SchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
			"schema_name":  types.StringType,
		},
	}
}

type EnableResponse_SdkV2 struct {
}

func (newState *EnableResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnableResponse_SdkV2) {
}

func (newState *EnableResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState EnableResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EnableResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o EnableResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o EnableResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Encryption options that apply to clients connecting to cloud storage.
type EncryptionDetails_SdkV2 struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	SseEncryptionDetails types.List `tfsdk:"sse_encryption_details" tf:"optional,object"`
}

func (newState *EncryptionDetails_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EncryptionDetails_SdkV2) {
}

func (newState *EncryptionDetails_SdkV2) SyncEffectiveFieldsDuringRead(existingState EncryptionDetails_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EncryptionDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EncryptionDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sse_encryption_details": reflect.TypeOf(SseEncryptionDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EncryptionDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o EncryptionDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sse_encryption_details": o.SseEncryptionDetails,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EncryptionDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sse_encryption_details": basetypes.ListType{
				ElemType: SseEncryptionDetails_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSseEncryptionDetails returns the value of the SseEncryptionDetails field in EncryptionDetails_SdkV2 as
// a SseEncryptionDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EncryptionDetails_SdkV2) GetSseEncryptionDetails(ctx context.Context) (SseEncryptionDetails_SdkV2, bool) {
	var e SseEncryptionDetails_SdkV2
	if o.SseEncryptionDetails.IsNull() || o.SseEncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []SseEncryptionDetails_SdkV2
	d := o.SseEncryptionDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSseEncryptionDetails sets the value of the SseEncryptionDetails field in EncryptionDetails_SdkV2.
func (o *EncryptionDetails_SdkV2) SetSseEncryptionDetails(ctx context.Context, v SseEncryptionDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sse_encryption_details"]
	o.SseEncryptionDetails = types.ListValueMust(t, vs)
}

// Get boolean reflecting if table exists
type ExistsRequest_SdkV2 struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`
}

func (newState *ExistsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExistsRequest_SdkV2) {
}

func (newState *ExistsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExistsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExistsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExistsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExistsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ExistsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": o.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExistsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
		},
	}
}

type ExternalLocationInfo_SdkV2 struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this external location was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of external location creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique ID of the location's storage credential.
	CredentialId types.String `tfsdk:"credential_id" tf:"optional"`
	// Name of the storage credential used with this location.
	CredentialName types.String `tfsdk:"credential_name" tf:"optional"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails types.List `tfsdk:"encryption_details" tf:"optional,object"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback types.Bool `tfsdk:"fallback" tf:"optional"`

	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Unique identifier of metastore hosting the external location.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of the external location.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The owner of the external location.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Indicates whether the external location is read-only.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the external location.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// Path URL of the external location.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *ExternalLocationInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalLocationInfo_SdkV2) {
}

func (newState *ExternalLocationInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExternalLocationInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalLocationInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalLocationInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"encryption_details": reflect.TypeOf(EncryptionDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalLocationInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ExternalLocationInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_point":       o.AccessPoint,
			"browse_only":        o.BrowseOnly,
			"comment":            o.Comment,
			"created_at":         o.CreatedAt,
			"created_by":         o.CreatedBy,
			"credential_id":      o.CredentialId,
			"credential_name":    o.CredentialName,
			"encryption_details": o.EncryptionDetails,
			"fallback":           o.Fallback,
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
func (o ExternalLocationInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point":    types.StringType,
			"browse_only":     types.BoolType,
			"comment":         types.StringType,
			"created_at":      types.Int64Type,
			"created_by":      types.StringType,
			"credential_id":   types.StringType,
			"credential_name": types.StringType,
			"encryption_details": basetypes.ListType{
				ElemType: EncryptionDetails_SdkV2{}.Type(ctx),
			},
			"fallback":       types.BoolType,
			"isolation_mode": types.StringType,
			"metastore_id":   types.StringType,
			"name":           types.StringType,
			"owner":          types.StringType,
			"read_only":      types.BoolType,
			"updated_at":     types.Int64Type,
			"updated_by":     types.StringType,
			"url":            types.StringType,
		},
	}
}

// GetEncryptionDetails returns the value of the EncryptionDetails field in ExternalLocationInfo_SdkV2 as
// a EncryptionDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalLocationInfo_SdkV2) GetEncryptionDetails(ctx context.Context) (EncryptionDetails_SdkV2, bool) {
	var e EncryptionDetails_SdkV2
	if o.EncryptionDetails.IsNull() || o.EncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []EncryptionDetails_SdkV2
	d := o.EncryptionDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEncryptionDetails sets the value of the EncryptionDetails field in ExternalLocationInfo_SdkV2.
func (o *ExternalLocationInfo_SdkV2) SetEncryptionDetails(ctx context.Context, v EncryptionDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["encryption_details"]
	o.EncryptionDetails = types.ListValueMust(t, vs)
}

// Detailed status of an online table. Shown if the online table is in the
// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
type FailedStatus_SdkV2 struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may only be partially synced to the online
	// table. Only populated if the table is still online and available for
	// serving.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version" tf:"optional"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	Timestamp types.String `tfsdk:"timestamp" tf:"optional"`
}

func (newState *FailedStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FailedStatus_SdkV2) {
}

func (newState *FailedStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState FailedStatus_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FailedStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FailedStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FailedStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o FailedStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": o.LastProcessedCommitVersion,
			"timestamp":                     o.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FailedStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
		},
	}
}

type ForeignKeyConstraint_SdkV2 struct {
	// Column names for this constraint.
	ChildColumns types.List `tfsdk:"child_columns" tf:""`
	// The name of the constraint.
	Name types.String `tfsdk:"name" tf:""`
	// Column names for this constraint.
	ParentColumns types.List `tfsdk:"parent_columns" tf:""`
	// The full name of the parent constraint.
	ParentTable types.String `tfsdk:"parent_table" tf:""`
}

func (newState *ForeignKeyConstraint_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForeignKeyConstraint_SdkV2) {
}

func (newState *ForeignKeyConstraint_SdkV2) SyncEffectiveFieldsDuringRead(existingState ForeignKeyConstraint_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForeignKeyConstraint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ForeignKeyConstraint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"child_columns":  reflect.TypeOf(types.String{}),
		"parent_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForeignKeyConstraint_SdkV2
// only implements ToObjectValue() and Type().
func (o ForeignKeyConstraint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ForeignKeyConstraint_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetChildColumns returns the value of the ChildColumns field in ForeignKeyConstraint_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ForeignKeyConstraint_SdkV2) GetChildColumns(ctx context.Context) ([]types.String, bool) {
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

// SetChildColumns sets the value of the ChildColumns field in ForeignKeyConstraint_SdkV2.
func (o *ForeignKeyConstraint_SdkV2) SetChildColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["child_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ChildColumns = types.ListValueMust(t, vs)
}

// GetParentColumns returns the value of the ParentColumns field in ForeignKeyConstraint_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ForeignKeyConstraint_SdkV2) GetParentColumns(ctx context.Context) ([]types.String, bool) {
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

// SetParentColumns sets the value of the ParentColumns field in ForeignKeyConstraint_SdkV2.
func (o *ForeignKeyConstraint_SdkV2) SetParentColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["parent_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ParentColumns = types.ListValueMust(t, vs)
}

// A function that is dependent on a SQL object.
type FunctionDependency_SdkV2 struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	FunctionFullName types.String `tfsdk:"function_full_name" tf:""`
}

func (newState *FunctionDependency_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionDependency_SdkV2) {
}

func (newState *FunctionDependency_SdkV2) SyncEffectiveFieldsDuringRead(existingState FunctionDependency_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FunctionDependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FunctionDependency_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionDependency_SdkV2
// only implements ToObjectValue() and Type().
func (o FunctionDependency_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_full_name": o.FunctionFullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FunctionDependency_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_full_name": types.StringType,
		},
	}
}

type FunctionInfo_SdkV2 struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this function was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of function creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Scalar function return data type.
	DataType types.String `tfsdk:"data_type" tf:"optional"`
	// External function language.
	ExternalLanguage types.String `tfsdk:"external_language" tf:"optional"`
	// External function name.
	ExternalName types.String `tfsdk:"external_name" tf:"optional"`
	// Pretty printed function data type.
	FullDataType types.String `tfsdk:"full_data_type" tf:"optional"`
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Id of Function, relative to parent schema.
	FunctionId types.String `tfsdk:"function_id" tf:"optional"`

	InputParams types.List `tfsdk:"input_params" tf:"optional,object"`
	// Whether the function is deterministic.
	IsDeterministic types.Bool `tfsdk:"is_deterministic" tf:"optional"`
	// Function null call.
	IsNullCall types.Bool `tfsdk:"is_null_call" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of function, relative to parent schema.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of current owner of function.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle types.String `tfsdk:"parameter_style" tf:"optional"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties types.String `tfsdk:"properties" tf:"optional"`
	// Table function return parameters.
	ReturnParams types.List `tfsdk:"return_params" tf:"optional,object"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody types.String `tfsdk:"routine_body" tf:"optional"`
	// Function body.
	RoutineDefinition types.String `tfsdk:"routine_definition" tf:"optional"`
	// Function dependencies.
	RoutineDependencies types.List `tfsdk:"routine_dependencies" tf:"optional,object"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// Function security type.
	SecurityType types.String `tfsdk:"security_type" tf:"optional"`
	// Specific name of the function; Reserved for future use.
	SpecificName types.String `tfsdk:"specific_name" tf:"optional"`
	// Function SQL data access.
	SqlDataAccess types.String `tfsdk:"sql_data_access" tf:"optional"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath types.String `tfsdk:"sql_path" tf:"optional"`
	// Time at which this function was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified function.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *FunctionInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionInfo_SdkV2) {
}

func (newState *FunctionInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState FunctionInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FunctionInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FunctionInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input_params":         reflect.TypeOf(FunctionParameterInfos_SdkV2{}),
		"return_params":        reflect.TypeOf(FunctionParameterInfos_SdkV2{}),
		"routine_dependencies": reflect.TypeOf(DependencyList_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o FunctionInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o FunctionInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":       types.BoolType,
			"catalog_name":      types.StringType,
			"comment":           types.StringType,
			"created_at":        types.Int64Type,
			"created_by":        types.StringType,
			"data_type":         types.StringType,
			"external_language": types.StringType,
			"external_name":     types.StringType,
			"full_data_type":    types.StringType,
			"full_name":         types.StringType,
			"function_id":       types.StringType,
			"input_params": basetypes.ListType{
				ElemType: FunctionParameterInfos_SdkV2{}.Type(ctx),
			},
			"is_deterministic": types.BoolType,
			"is_null_call":     types.BoolType,
			"metastore_id":     types.StringType,
			"name":             types.StringType,
			"owner":            types.StringType,
			"parameter_style":  types.StringType,
			"properties":       types.StringType,
			"return_params": basetypes.ListType{
				ElemType: FunctionParameterInfos_SdkV2{}.Type(ctx),
			},
			"routine_body":       types.StringType,
			"routine_definition": types.StringType,
			"routine_dependencies": basetypes.ListType{
				ElemType: DependencyList_SdkV2{}.Type(ctx),
			},
			"schema_name":     types.StringType,
			"security_type":   types.StringType,
			"specific_name":   types.StringType,
			"sql_data_access": types.StringType,
			"sql_path":        types.StringType,
			"updated_at":      types.Int64Type,
			"updated_by":      types.StringType,
		},
	}
}

// GetInputParams returns the value of the InputParams field in FunctionInfo_SdkV2 as
// a FunctionParameterInfos_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *FunctionInfo_SdkV2) GetInputParams(ctx context.Context) (FunctionParameterInfos_SdkV2, bool) {
	var e FunctionParameterInfos_SdkV2
	if o.InputParams.IsNull() || o.InputParams.IsUnknown() {
		return e, false
	}
	var v []FunctionParameterInfos_SdkV2
	d := o.InputParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInputParams sets the value of the InputParams field in FunctionInfo_SdkV2.
func (o *FunctionInfo_SdkV2) SetInputParams(ctx context.Context, v FunctionParameterInfos_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["input_params"]
	o.InputParams = types.ListValueMust(t, vs)
}

// GetReturnParams returns the value of the ReturnParams field in FunctionInfo_SdkV2 as
// a FunctionParameterInfos_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *FunctionInfo_SdkV2) GetReturnParams(ctx context.Context) (FunctionParameterInfos_SdkV2, bool) {
	var e FunctionParameterInfos_SdkV2
	if o.ReturnParams.IsNull() || o.ReturnParams.IsUnknown() {
		return e, false
	}
	var v []FunctionParameterInfos_SdkV2
	d := o.ReturnParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetReturnParams sets the value of the ReturnParams field in FunctionInfo_SdkV2.
func (o *FunctionInfo_SdkV2) SetReturnParams(ctx context.Context, v FunctionParameterInfos_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["return_params"]
	o.ReturnParams = types.ListValueMust(t, vs)
}

// GetRoutineDependencies returns the value of the RoutineDependencies field in FunctionInfo_SdkV2 as
// a DependencyList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *FunctionInfo_SdkV2) GetRoutineDependencies(ctx context.Context) (DependencyList_SdkV2, bool) {
	var e DependencyList_SdkV2
	if o.RoutineDependencies.IsNull() || o.RoutineDependencies.IsUnknown() {
		return e, false
	}
	var v []DependencyList_SdkV2
	d := o.RoutineDependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRoutineDependencies sets the value of the RoutineDependencies field in FunctionInfo_SdkV2.
func (o *FunctionInfo_SdkV2) SetRoutineDependencies(ctx context.Context, v DependencyList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["routine_dependencies"]
	o.RoutineDependencies = types.ListValueMust(t, vs)
}

type FunctionParameterInfo_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of parameter.
	Name types.String `tfsdk:"name" tf:""`
	// Default value of the parameter.
	ParameterDefault types.String `tfsdk:"parameter_default" tf:"optional"`
	// The mode of the function parameter.
	ParameterMode types.String `tfsdk:"parameter_mode" tf:"optional"`
	// The type of function parameter.
	ParameterType types.String `tfsdk:"parameter_type" tf:"optional"`
	// Ordinal position of column (starting at position 0).
	Position types.Int64 `tfsdk:"position" tf:""`
	// Format of IntervalType.
	TypeIntervalType types.String `tfsdk:"type_interval_type" tf:"optional"`
	// Full data type spec, JSON-serialized.
	TypeJson types.String `tfsdk:"type_json" tf:"optional"`

	TypeName types.String `tfsdk:"type_name" tf:""`
	// Digits of precision; required on Create for DecimalTypes.
	TypePrecision types.Int64 `tfsdk:"type_precision" tf:"optional"`
	// Digits to right of decimal; Required on Create for DecimalTypes.
	TypeScale types.Int64 `tfsdk:"type_scale" tf:"optional"`
	// Full data type spec, SQL/catalogString text.
	TypeText types.String `tfsdk:"type_text" tf:""`
}

func (newState *FunctionParameterInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionParameterInfo_SdkV2) {
}

func (newState *FunctionParameterInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState FunctionParameterInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FunctionParameterInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FunctionParameterInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionParameterInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o FunctionParameterInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o FunctionParameterInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
}

func (newState *FunctionParameterInfos_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionParameterInfos_SdkV2) {
}

func (newState *FunctionParameterInfos_SdkV2) SyncEffectiveFieldsDuringRead(existingState FunctionParameterInfos_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FunctionParameterInfos.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FunctionParameterInfos_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(FunctionParameterInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionParameterInfos_SdkV2
// only implements ToObjectValue() and Type().
func (o FunctionParameterInfos_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parameters": o.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FunctionParameterInfos_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *FunctionParameterInfos_SdkV2) GetParameters(ctx context.Context) ([]FunctionParameterInfo_SdkV2, bool) {
	if o.Parameters.IsNull() || o.Parameters.IsUnknown() {
		return nil, false
	}
	var v []FunctionParameterInfo_SdkV2
	d := o.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in FunctionParameterInfos_SdkV2.
func (o *FunctionParameterInfos_SdkV2) SetParameters(ctx context.Context, v []FunctionParameterInfo_SdkV2) {
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
type GcpOauthToken_SdkV2 struct {
	OauthToken types.String `tfsdk:"oauth_token" tf:"optional"`
}

func (newState *GcpOauthToken_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpOauthToken_SdkV2) {
}

func (newState *GcpOauthToken_SdkV2) SyncEffectiveFieldsDuringRead(existingState GcpOauthToken_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpOauthToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpOauthToken_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpOauthToken_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpOauthToken_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"oauth_token": o.OauthToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpOauthToken_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"oauth_token": types.StringType,
		},
	}
}

// The Azure cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialAzureOptions_SdkV2 struct {
	// The resources to which the temporary Azure credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://learn.microsoft.com/python/api/azure-core/azure.core.credentials.tokencredential?view=azure-python)
	Resources types.List `tfsdk:"resources" tf:"optional"`
}

func (newState *GenerateTemporaryServiceCredentialAzureOptions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryServiceCredentialAzureOptions_SdkV2) {
}

func (newState *GenerateTemporaryServiceCredentialAzureOptions_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryServiceCredentialAzureOptions_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateTemporaryServiceCredentialAzureOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenerateTemporaryServiceCredentialAzureOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resources": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateTemporaryServiceCredentialAzureOptions_SdkV2
// only implements ToObjectValue() and Type().
func (o GenerateTemporaryServiceCredentialAzureOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resources": o.Resources,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateTemporaryServiceCredentialAzureOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resources": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetResources returns the value of the Resources field in GenerateTemporaryServiceCredentialAzureOptions_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryServiceCredentialAzureOptions_SdkV2) GetResources(ctx context.Context) ([]types.String, bool) {
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

// SetResources sets the value of the Resources field in GenerateTemporaryServiceCredentialAzureOptions_SdkV2.
func (o *GenerateTemporaryServiceCredentialAzureOptions_SdkV2) SetResources(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

// The GCP cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialGcpOptions_SdkV2 struct {
	// The scopes to which the temporary GCP credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://google-auth.readthedocs.io/en/latest/reference/google.auth.html#google.auth.credentials.Credentials)
	Scopes types.List `tfsdk:"scopes" tf:"optional"`
}

func (newState *GenerateTemporaryServiceCredentialGcpOptions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryServiceCredentialGcpOptions_SdkV2) {
}

func (newState *GenerateTemporaryServiceCredentialGcpOptions_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryServiceCredentialGcpOptions_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateTemporaryServiceCredentialGcpOptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenerateTemporaryServiceCredentialGcpOptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateTemporaryServiceCredentialGcpOptions_SdkV2
// only implements ToObjectValue() and Type().
func (o GenerateTemporaryServiceCredentialGcpOptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scopes": o.Scopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateTemporaryServiceCredentialGcpOptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetScopes returns the value of the Scopes field in GenerateTemporaryServiceCredentialGcpOptions_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryServiceCredentialGcpOptions_SdkV2) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in GenerateTemporaryServiceCredentialGcpOptions_SdkV2.
func (o *GenerateTemporaryServiceCredentialGcpOptions_SdkV2) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

type GenerateTemporaryServiceCredentialRequest_SdkV2 struct {
	// The Azure cloud options to customize the requested temporary credential
	AzureOptions types.List `tfsdk:"azure_options" tf:"optional,object"`
	// The name of the service credential used to generate a temporary
	// credential
	CredentialName types.String `tfsdk:"credential_name" tf:""`
	// The GCP cloud options to customize the requested temporary credential
	GcpOptions types.List `tfsdk:"gcp_options" tf:"optional,object"`
}

func (newState *GenerateTemporaryServiceCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryServiceCredentialRequest_SdkV2) {
}

func (newState *GenerateTemporaryServiceCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryServiceCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateTemporaryServiceCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenerateTemporaryServiceCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"azure_options": reflect.TypeOf(GenerateTemporaryServiceCredentialAzureOptions_SdkV2{}),
		"gcp_options":   reflect.TypeOf(GenerateTemporaryServiceCredentialGcpOptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateTemporaryServiceCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenerateTemporaryServiceCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"azure_options":   o.AzureOptions,
			"credential_name": o.CredentialName,
			"gcp_options":     o.GcpOptions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateTemporaryServiceCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"azure_options": basetypes.ListType{
				ElemType: GenerateTemporaryServiceCredentialAzureOptions_SdkV2{}.Type(ctx),
			},
			"credential_name": types.StringType,
			"gcp_options": basetypes.ListType{
				ElemType: GenerateTemporaryServiceCredentialGcpOptions_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAzureOptions returns the value of the AzureOptions field in GenerateTemporaryServiceCredentialRequest_SdkV2 as
// a GenerateTemporaryServiceCredentialAzureOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryServiceCredentialRequest_SdkV2) GetAzureOptions(ctx context.Context) (GenerateTemporaryServiceCredentialAzureOptions_SdkV2, bool) {
	var e GenerateTemporaryServiceCredentialAzureOptions_SdkV2
	if o.AzureOptions.IsNull() || o.AzureOptions.IsUnknown() {
		return e, false
	}
	var v []GenerateTemporaryServiceCredentialAzureOptions_SdkV2
	d := o.AzureOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureOptions sets the value of the AzureOptions field in GenerateTemporaryServiceCredentialRequest_SdkV2.
func (o *GenerateTemporaryServiceCredentialRequest_SdkV2) SetAzureOptions(ctx context.Context, v GenerateTemporaryServiceCredentialAzureOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_options"]
	o.AzureOptions = types.ListValueMust(t, vs)
}

// GetGcpOptions returns the value of the GcpOptions field in GenerateTemporaryServiceCredentialRequest_SdkV2 as
// a GenerateTemporaryServiceCredentialGcpOptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryServiceCredentialRequest_SdkV2) GetGcpOptions(ctx context.Context) (GenerateTemporaryServiceCredentialGcpOptions_SdkV2, bool) {
	var e GenerateTemporaryServiceCredentialGcpOptions_SdkV2
	if o.GcpOptions.IsNull() || o.GcpOptions.IsUnknown() {
		return e, false
	}
	var v []GenerateTemporaryServiceCredentialGcpOptions_SdkV2
	d := o.GcpOptions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpOptions sets the value of the GcpOptions field in GenerateTemporaryServiceCredentialRequest_SdkV2.
func (o *GenerateTemporaryServiceCredentialRequest_SdkV2) SetGcpOptions(ctx context.Context, v GenerateTemporaryServiceCredentialGcpOptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_options"]
	o.GcpOptions = types.ListValueMust(t, vs)
}

type GenerateTemporaryTableCredentialRequest_SdkV2 struct {
	// The operation performed against the table data, either READ or
	// READ_WRITE. If READ_WRITE is specified, the credentials returned will
	// have write permissions, otherwise, it will be read only.
	Operation types.String `tfsdk:"operation" tf:"optional"`
	// UUID of the table to read or write.
	TableId types.String `tfsdk:"table_id" tf:"optional"`
}

func (newState *GenerateTemporaryTableCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryTableCredentialRequest_SdkV2) {
}

func (newState *GenerateTemporaryTableCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryTableCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateTemporaryTableCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenerateTemporaryTableCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateTemporaryTableCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenerateTemporaryTableCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"operation": o.Operation,
			"table_id":  o.TableId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateTemporaryTableCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"operation": types.StringType,
			"table_id":  types.StringType,
		},
	}
}

type GenerateTemporaryTableCredentialResponse_SdkV2 struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	AwsTempCredentials types.List `tfsdk:"aws_temp_credentials" tf:"optional,object"`
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	AzureAad types.List `tfsdk:"azure_aad" tf:"optional,object"`
	// Azure temporary credentials for API authentication. Read more at
	// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
	AzureUserDelegationSas types.List `tfsdk:"azure_user_delegation_sas" tf:"optional,object"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	GcpOauthToken types.List `tfsdk:"gcp_oauth_token" tf:"optional,object"`
	// R2 temporary credentials for API authentication. Read more at
	// https://developers.cloudflare.com/r2/api/s3/tokens/.
	R2TempCredentials types.List `tfsdk:"r2_temp_credentials" tf:"optional,object"`
	// The URL of the storage path accessible by the temporary credential.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *GenerateTemporaryTableCredentialResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryTableCredentialResponse_SdkV2) {
}

func (newState *GenerateTemporaryTableCredentialResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryTableCredentialResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateTemporaryTableCredentialResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenerateTemporaryTableCredentialResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_temp_credentials":      reflect.TypeOf(AwsCredentials_SdkV2{}),
		"azure_aad":                 reflect.TypeOf(AzureActiveDirectoryToken_SdkV2{}),
		"azure_user_delegation_sas": reflect.TypeOf(AzureUserDelegationSas_SdkV2{}),
		"gcp_oauth_token":           reflect.TypeOf(GcpOauthToken_SdkV2{}),
		"r2_temp_credentials":       reflect.TypeOf(R2Credentials_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateTemporaryTableCredentialResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GenerateTemporaryTableCredentialResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GenerateTemporaryTableCredentialResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_temp_credentials": basetypes.ListType{
				ElemType: AwsCredentials_SdkV2{}.Type(ctx),
			},
			"azure_aad": basetypes.ListType{
				ElemType: AzureActiveDirectoryToken_SdkV2{}.Type(ctx),
			},
			"azure_user_delegation_sas": basetypes.ListType{
				ElemType: AzureUserDelegationSas_SdkV2{}.Type(ctx),
			},
			"expiration_time": types.Int64Type,
			"gcp_oauth_token": basetypes.ListType{
				ElemType: GcpOauthToken_SdkV2{}.Type(ctx),
			},
			"r2_temp_credentials": basetypes.ListType{
				ElemType: R2Credentials_SdkV2{}.Type(ctx),
			},
			"url": types.StringType,
		},
	}
}

// GetAwsTempCredentials returns the value of the AwsTempCredentials field in GenerateTemporaryTableCredentialResponse_SdkV2 as
// a AwsCredentials_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryTableCredentialResponse_SdkV2) GetAwsTempCredentials(ctx context.Context) (AwsCredentials_SdkV2, bool) {
	var e AwsCredentials_SdkV2
	if o.AwsTempCredentials.IsNull() || o.AwsTempCredentials.IsUnknown() {
		return e, false
	}
	var v []AwsCredentials_SdkV2
	d := o.AwsTempCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsTempCredentials sets the value of the AwsTempCredentials field in GenerateTemporaryTableCredentialResponse_SdkV2.
func (o *GenerateTemporaryTableCredentialResponse_SdkV2) SetAwsTempCredentials(ctx context.Context, v AwsCredentials_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_temp_credentials"]
	o.AwsTempCredentials = types.ListValueMust(t, vs)
}

// GetAzureAad returns the value of the AzureAad field in GenerateTemporaryTableCredentialResponse_SdkV2 as
// a AzureActiveDirectoryToken_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryTableCredentialResponse_SdkV2) GetAzureAad(ctx context.Context) (AzureActiveDirectoryToken_SdkV2, bool) {
	var e AzureActiveDirectoryToken_SdkV2
	if o.AzureAad.IsNull() || o.AzureAad.IsUnknown() {
		return e, false
	}
	var v []AzureActiveDirectoryToken_SdkV2
	d := o.AzureAad.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAad sets the value of the AzureAad field in GenerateTemporaryTableCredentialResponse_SdkV2.
func (o *GenerateTemporaryTableCredentialResponse_SdkV2) SetAzureAad(ctx context.Context, v AzureActiveDirectoryToken_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_aad"]
	o.AzureAad = types.ListValueMust(t, vs)
}

// GetAzureUserDelegationSas returns the value of the AzureUserDelegationSas field in GenerateTemporaryTableCredentialResponse_SdkV2 as
// a AzureUserDelegationSas_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryTableCredentialResponse_SdkV2) GetAzureUserDelegationSas(ctx context.Context) (AzureUserDelegationSas_SdkV2, bool) {
	var e AzureUserDelegationSas_SdkV2
	if o.AzureUserDelegationSas.IsNull() || o.AzureUserDelegationSas.IsUnknown() {
		return e, false
	}
	var v []AzureUserDelegationSas_SdkV2
	d := o.AzureUserDelegationSas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureUserDelegationSas sets the value of the AzureUserDelegationSas field in GenerateTemporaryTableCredentialResponse_SdkV2.
func (o *GenerateTemporaryTableCredentialResponse_SdkV2) SetAzureUserDelegationSas(ctx context.Context, v AzureUserDelegationSas_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_user_delegation_sas"]
	o.AzureUserDelegationSas = types.ListValueMust(t, vs)
}

// GetGcpOauthToken returns the value of the GcpOauthToken field in GenerateTemporaryTableCredentialResponse_SdkV2 as
// a GcpOauthToken_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryTableCredentialResponse_SdkV2) GetGcpOauthToken(ctx context.Context) (GcpOauthToken_SdkV2, bool) {
	var e GcpOauthToken_SdkV2
	if o.GcpOauthToken.IsNull() || o.GcpOauthToken.IsUnknown() {
		return e, false
	}
	var v []GcpOauthToken_SdkV2
	d := o.GcpOauthToken.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpOauthToken sets the value of the GcpOauthToken field in GenerateTemporaryTableCredentialResponse_SdkV2.
func (o *GenerateTemporaryTableCredentialResponse_SdkV2) SetGcpOauthToken(ctx context.Context, v GcpOauthToken_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_oauth_token"]
	o.GcpOauthToken = types.ListValueMust(t, vs)
}

// GetR2TempCredentials returns the value of the R2TempCredentials field in GenerateTemporaryTableCredentialResponse_SdkV2 as
// a R2Credentials_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateTemporaryTableCredentialResponse_SdkV2) GetR2TempCredentials(ctx context.Context) (R2Credentials_SdkV2, bool) {
	var e R2Credentials_SdkV2
	if o.R2TempCredentials.IsNull() || o.R2TempCredentials.IsUnknown() {
		return e, false
	}
	var v []R2Credentials_SdkV2
	d := o.R2TempCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetR2TempCredentials sets the value of the R2TempCredentials field in GenerateTemporaryTableCredentialResponse_SdkV2.
func (o *GenerateTemporaryTableCredentialResponse_SdkV2) SetR2TempCredentials(ctx context.Context, v R2Credentials_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["r2_temp_credentials"]
	o.R2TempCredentials = types.ListValueMust(t, vs)
}

// Gets the metastore assignment for a workspace
type GetAccountMetastoreAssignmentRequest_SdkV2 struct {
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *GetAccountMetastoreAssignmentRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountMetastoreAssignmentRequest_SdkV2) {
}

func (newState *GetAccountMetastoreAssignmentRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetAccountMetastoreAssignmentRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountMetastoreAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountMetastoreAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountMetastoreAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAccountMetastoreAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountMetastoreAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

// Get a metastore
type GetAccountMetastoreRequest_SdkV2 struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *GetAccountMetastoreRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountMetastoreRequest_SdkV2) {
}

func (newState *GetAccountMetastoreRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetAccountMetastoreRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountMetastoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountMetastoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountMetastoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAccountMetastoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountMetastoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
		},
	}
}

// Gets the named storage credential
type GetAccountStorageCredentialRequest_SdkV2 struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Name of the storage credential.
	StorageCredentialName types.String `tfsdk:"-"`
}

func (newState *GetAccountStorageCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountStorageCredentialRequest_SdkV2) {
}

func (newState *GetAccountStorageCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetAccountStorageCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountStorageCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountStorageCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountStorageCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAccountStorageCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id":            o.MetastoreId,
			"storage_credential_name": o.StorageCredentialName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountStorageCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id":            types.StringType,
			"storage_credential_name": types.StringType,
		},
	}
}

// Get an artifact allowlist
type GetArtifactAllowlistRequest_SdkV2 struct {
	// The artifact type of the allowlist.
	ArtifactType types.String `tfsdk:"-"`
}

func (newState *GetArtifactAllowlistRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetArtifactAllowlistRequest_SdkV2) {
}

func (newState *GetArtifactAllowlistRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetArtifactAllowlistRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetArtifactAllowlistRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetArtifactAllowlistRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetArtifactAllowlistRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetArtifactAllowlistRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_type": o.ArtifactType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetArtifactAllowlistRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_type": types.StringType,
		},
	}
}

// Get securable workspace bindings
type GetBindingsRequest_SdkV2 struct {
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
	// The type of the securable to bind to a workspace.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *GetBindingsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetBindingsRequest_SdkV2) {
}

func (newState *GetBindingsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetBindingsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBindingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetBindingsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBindingsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetBindingsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetBindingsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
type GetByAliasRequest_SdkV2 struct {
	// The name of the alias
	Alias types.String `tfsdk:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases types.Bool `tfsdk:"-"`
}

func (newState *GetByAliasRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetByAliasRequest_SdkV2) {
}

func (newState *GetByAliasRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetByAliasRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetByAliasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetByAliasRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetByAliasRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetByAliasRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":           o.Alias,
			"full_name":       o.FullName,
			"include_aliases": o.IncludeAliases,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetByAliasRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":           types.StringType,
			"full_name":       types.StringType,
			"include_aliases": types.BoolType,
		},
	}
}

// Get a catalog
type GetCatalogRequest_SdkV2 struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
}

func (newState *GetCatalogRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCatalogRequest_SdkV2) {
}

func (newState *GetCatalogRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetCatalogRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"name":           o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"name":           types.StringType,
		},
	}
}

// Get a connection
type GetConnectionRequest_SdkV2 struct {
	// Name of the connection.
	Name types.String `tfsdk:"-"`
}

func (newState *GetConnectionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetConnectionRequest_SdkV2) {
}

func (newState *GetConnectionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetConnectionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetConnectionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetConnectionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetConnectionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetConnectionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetConnectionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a credential
type GetCredentialRequest_SdkV2 struct {
	// Name of the credential.
	NameArg types.String `tfsdk:"-"`
}

func (newState *GetCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCredentialRequest_SdkV2) {
}

func (newState *GetCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name_arg": o.NameArg,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name_arg": types.StringType,
		},
	}
}

// Get effective permissions
type GetEffectiveRequest_SdkV2 struct {
	// Full name of securable.
	FullName types.String `tfsdk:"-"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal types.String `tfsdk:"-"`
	// Type of securable.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *GetEffectiveRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEffectiveRequest_SdkV2) {
}

func (newState *GetEffectiveRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetEffectiveRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEffectiveRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEffectiveRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEffectiveRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetEffectiveRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":      o.FullName,
			"principal":      o.Principal,
			"securable_type": o.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEffectiveRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":      types.StringType,
			"principal":      types.StringType,
			"securable_type": types.StringType,
		},
	}
}

// Get an external location
type GetExternalLocationRequest_SdkV2 struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// Name of the external location.
	Name types.String `tfsdk:"-"`
}

func (newState *GetExternalLocationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExternalLocationRequest_SdkV2) {
}

func (newState *GetExternalLocationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetExternalLocationRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExternalLocationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetExternalLocationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExternalLocationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetExternalLocationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"name":           o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExternalLocationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"name":           types.StringType,
		},
	}
}

// Get a function
type GetFunctionRequest_SdkV2 struct {
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name types.String `tfsdk:"-"`
}

func (newState *GetFunctionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetFunctionRequest_SdkV2) {
}

func (newState *GetFunctionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetFunctionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFunctionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetFunctionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFunctionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetFunctionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"name":           o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetFunctionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"name":           types.StringType,
		},
	}
}

// Get permissions
type GetGrantRequest_SdkV2 struct {
	// Full name of securable.
	FullName types.String `tfsdk:"-"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal types.String `tfsdk:"-"`
	// Type of securable.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *GetGrantRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetGrantRequest_SdkV2) {
}

func (newState *GetGrantRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetGrantRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGrantRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetGrantRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetGrantRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetGrantRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":      o.FullName,
			"principal":      o.Principal,
			"securable_type": o.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetGrantRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":      types.StringType,
			"principal":      types.StringType,
			"securable_type": types.StringType,
		},
	}
}

// Get a metastore
type GetMetastoreRequest_SdkV2 struct {
	// Unique ID of the metastore.
	Id types.String `tfsdk:"-"`
}

func (newState *GetMetastoreRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetastoreRequest_SdkV2) {
}

func (newState *GetMetastoreRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetMetastoreRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMetastoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetMetastoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetastoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetMetastoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetMetastoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetMetastoreSummaryResponse_SdkV2 struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of metastore creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId types.String `tfsdk:"default_data_access_config_id" tf:"optional"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName types.String `tfsdk:"delta_sharing_organization_name" tf:"optional"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds types.Int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds" tf:"optional"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope types.String `tfsdk:"delta_sharing_scope" tf:"optional"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled types.Bool `tfsdk:"external_access_enabled" tf:"optional"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id" tf:"optional"`
	// Unique identifier of metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The user-specified name of the metastore.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The owner of the metastore.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion types.String `tfsdk:"privilege_model_version" tf:"optional"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region types.String `tfsdk:"region" tf:"optional"`
	// The storage root URL for metastore
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId types.String `tfsdk:"storage_root_credential_id" tf:"optional"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName types.String `tfsdk:"storage_root_credential_name" tf:"optional"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the metastore.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *GetMetastoreSummaryResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetastoreSummaryResponse_SdkV2) {
}

func (newState *GetMetastoreSummaryResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetMetastoreSummaryResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMetastoreSummaryResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetMetastoreSummaryResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetastoreSummaryResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetMetastoreSummaryResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetMetastoreSummaryResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
type GetModelVersionRequest_SdkV2 struct {
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

func (newState *GetModelVersionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionRequest_SdkV2) {
}

func (newState *GetModelVersionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetModelVersionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
type GetOnlineTableRequest_SdkV2 struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"-"`
}

func (newState *GetOnlineTableRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetOnlineTableRequest_SdkV2) {
}

func (newState *GetOnlineTableRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetOnlineTableRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOnlineTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetOnlineTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOnlineTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetOnlineTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetOnlineTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a table monitor
type GetQualityMonitorRequest_SdkV2 struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *GetQualityMonitorRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQualityMonitorRequest_SdkV2) {
}

func (newState *GetQualityMonitorRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetQualityMonitorRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQualityMonitorRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQualityMonitorRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQualityMonitorRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetQualityMonitorRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQualityMonitorRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name": types.StringType,
		},
	}
}

// Get information for a single resource quota.
type GetQuotaRequest_SdkV2 struct {
	// Full name of the parent resource. Provide the metastore ID if the parent
	// is a metastore.
	ParentFullName types.String `tfsdk:"-"`
	// Securable type of the quota parent.
	ParentSecurableType types.String `tfsdk:"-"`
	// Name of the quota. Follows the pattern of the quota type, with "-quota"
	// added as a suffix.
	QuotaName types.String `tfsdk:"-"`
}

func (newState *GetQuotaRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQuotaRequest_SdkV2) {
}

func (newState *GetQuotaRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetQuotaRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQuotaRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQuotaRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQuotaRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetQuotaRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent_full_name":      o.ParentFullName,
			"parent_securable_type": o.ParentSecurableType,
			"quota_name":            o.QuotaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQuotaRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent_full_name":      types.StringType,
			"parent_securable_type": types.StringType,
			"quota_name":            types.StringType,
		},
	}
}

type GetQuotaResponse_SdkV2 struct {
	// The returned QuotaInfo.
	QuotaInfo types.List `tfsdk:"quota_info" tf:"optional,object"`
}

func (newState *GetQuotaResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQuotaResponse_SdkV2) {
}

func (newState *GetQuotaResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetQuotaResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetQuotaResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetQuotaResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quota_info": reflect.TypeOf(QuotaInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetQuotaResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetQuotaResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"quota_info": o.QuotaInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetQuotaResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quota_info": basetypes.ListType{
				ElemType: QuotaInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetQuotaInfo returns the value of the QuotaInfo field in GetQuotaResponse_SdkV2 as
// a QuotaInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetQuotaResponse_SdkV2) GetQuotaInfo(ctx context.Context) (QuotaInfo_SdkV2, bool) {
	var e QuotaInfo_SdkV2
	if o.QuotaInfo.IsNull() || o.QuotaInfo.IsUnknown() {
		return e, false
	}
	var v []QuotaInfo_SdkV2
	d := o.QuotaInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQuotaInfo sets the value of the QuotaInfo field in GetQuotaResponse_SdkV2.
func (o *GetQuotaResponse_SdkV2) SetQuotaInfo(ctx context.Context, v QuotaInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["quota_info"]
	o.QuotaInfo = types.ListValueMust(t, vs)
}

// Get refresh
type GetRefreshRequest_SdkV2 struct {
	// ID of the refresh.
	RefreshId types.String `tfsdk:"-"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *GetRefreshRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRefreshRequest_SdkV2) {
}

func (newState *GetRefreshRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetRefreshRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRefreshRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRefreshRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRefreshRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"refresh_id": o.RefreshId,
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRefreshRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"refresh_id": types.StringType,
			"table_name": types.StringType,
		},
	}
}

// Get a Registered Model
type GetRegisteredModelRequest_SdkV2 struct {
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
	// Whether to include registered model aliases in the response
	IncludeAliases types.Bool `tfsdk:"-"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
}

func (newState *GetRegisteredModelRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRegisteredModelRequest_SdkV2) {
}

func (newState *GetRegisteredModelRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetRegisteredModelRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRegisteredModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRegisteredModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRegisteredModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":       o.FullName,
			"include_aliases": o.IncludeAliases,
			"include_browse":  o.IncludeBrowse,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRegisteredModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":       types.StringType,
			"include_aliases": types.BoolType,
			"include_browse":  types.BoolType,
		},
	}
}

// Get a schema
type GetSchemaRequest_SdkV2 struct {
	// Full name of the schema.
	FullName types.String `tfsdk:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
}

func (newState *GetSchemaRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSchemaRequest_SdkV2) {
}

func (newState *GetSchemaRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetSchemaRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSchemaRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSchemaRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSchemaRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetSchemaRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":      o.FullName,
			"include_browse": o.IncludeBrowse,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSchemaRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":      types.StringType,
			"include_browse": types.BoolType,
		},
	}
}

// Get a credential
type GetStorageCredentialRequest_SdkV2 struct {
	// Name of the storage credential.
	Name types.String `tfsdk:"-"`
}

func (newState *GetStorageCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStorageCredentialRequest_SdkV2) {
}

func (newState *GetStorageCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetStorageCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStorageCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStorageCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStorageCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetStorageCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStorageCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a table
type GetTableRequest_SdkV2 struct {
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

func (newState *GetTableRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetTableRequest_SdkV2) {
}

func (newState *GetTableRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetTableRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
type GetWorkspaceBindingRequest_SdkV2 struct {
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
}

func (newState *GetWorkspaceBindingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceBindingRequest_SdkV2) {
}

func (newState *GetWorkspaceBindingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceBindingRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceBindingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceBindingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceBindingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceBindingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceBindingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get all workspaces assigned to a metastore
type ListAccountMetastoreAssignmentsRequest_SdkV2 struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *ListAccountMetastoreAssignmentsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountMetastoreAssignmentsRequest_SdkV2) {
}

func (newState *ListAccountMetastoreAssignmentsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAccountMetastoreAssignmentsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountMetastoreAssignmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountMetastoreAssignmentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountMetastoreAssignmentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountMetastoreAssignmentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountMetastoreAssignmentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
		},
	}
}

// The list of workspaces to which the given metastore is assigned.
type ListAccountMetastoreAssignmentsResponse_SdkV2 struct {
	WorkspaceIds types.List `tfsdk:"workspace_ids" tf:"optional"`
}

func (newState *ListAccountMetastoreAssignmentsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountMetastoreAssignmentsResponse_SdkV2) {
}

func (newState *ListAccountMetastoreAssignmentsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAccountMetastoreAssignmentsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountMetastoreAssignmentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountMetastoreAssignmentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_ids": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountMetastoreAssignmentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountMetastoreAssignmentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_ids": o.WorkspaceIds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountMetastoreAssignmentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetWorkspaceIds returns the value of the WorkspaceIds field in ListAccountMetastoreAssignmentsResponse_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAccountMetastoreAssignmentsResponse_SdkV2) GetWorkspaceIds(ctx context.Context) ([]types.Int64, bool) {
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

// SetWorkspaceIds sets the value of the WorkspaceIds field in ListAccountMetastoreAssignmentsResponse_SdkV2.
func (o *ListAccountMetastoreAssignmentsResponse_SdkV2) SetWorkspaceIds(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WorkspaceIds = types.ListValueMust(t, vs)
}

// Get all storage credentials assigned to a metastore
type ListAccountStorageCredentialsRequest_SdkV2 struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *ListAccountStorageCredentialsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountStorageCredentialsRequest_SdkV2) {
}

func (newState *ListAccountStorageCredentialsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAccountStorageCredentialsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountStorageCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountStorageCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountStorageCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountStorageCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountStorageCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
		},
	}
}

type ListAccountStorageCredentialsResponse_SdkV2 struct {
	// An array of metastore storage credentials.
	StorageCredentials types.List `tfsdk:"storage_credentials" tf:"optional"`
}

func (newState *ListAccountStorageCredentialsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountStorageCredentialsResponse_SdkV2) {
}

func (newState *ListAccountStorageCredentialsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListAccountStorageCredentialsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountStorageCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountStorageCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"storage_credentials": reflect.TypeOf(StorageCredentialInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountStorageCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountStorageCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_credentials": o.StorageCredentials,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountStorageCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_credentials": basetypes.ListType{
				ElemType: StorageCredentialInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetStorageCredentials returns the value of the StorageCredentials field in ListAccountStorageCredentialsResponse_SdkV2 as
// a slice of StorageCredentialInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAccountStorageCredentialsResponse_SdkV2) GetStorageCredentials(ctx context.Context) ([]StorageCredentialInfo_SdkV2, bool) {
	if o.StorageCredentials.IsNull() || o.StorageCredentials.IsUnknown() {
		return nil, false
	}
	var v []StorageCredentialInfo_SdkV2
	d := o.StorageCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStorageCredentials sets the value of the StorageCredentials field in ListAccountStorageCredentialsResponse_SdkV2.
func (o *ListAccountStorageCredentialsResponse_SdkV2) SetStorageCredentials(ctx context.Context, v []StorageCredentialInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["storage_credentials"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.StorageCredentials = types.ListValueMust(t, vs)
}

// List catalogs
type ListCatalogsRequest_SdkV2 struct {
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

func (newState *ListCatalogsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCatalogsRequest_SdkV2) {
}

func (newState *ListCatalogsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListCatalogsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCatalogsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCatalogsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCatalogsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCatalogsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"max_results":    o.MaxResults,
			"page_token":     o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCatalogsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
		},
	}
}

type ListCatalogsResponse_SdkV2 struct {
	// An array of catalog information objects.
	Catalogs types.List `tfsdk:"catalogs" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListCatalogsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCatalogsResponse_SdkV2) {
}

func (newState *ListCatalogsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListCatalogsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCatalogsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCatalogsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalogs": reflect.TypeOf(CatalogInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCatalogsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCatalogsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalogs":        o.Catalogs,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCatalogsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalogs": basetypes.ListType{
				ElemType: CatalogInfo_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetCatalogs returns the value of the Catalogs field in ListCatalogsResponse_SdkV2 as
// a slice of CatalogInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCatalogsResponse_SdkV2) GetCatalogs(ctx context.Context) ([]CatalogInfo_SdkV2, bool) {
	if o.Catalogs.IsNull() || o.Catalogs.IsUnknown() {
		return nil, false
	}
	var v []CatalogInfo_SdkV2
	d := o.Catalogs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCatalogs sets the value of the Catalogs field in ListCatalogsResponse_SdkV2.
func (o *ListCatalogsResponse_SdkV2) SetCatalogs(ctx context.Context, v []CatalogInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["catalogs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Catalogs = types.ListValueMust(t, vs)
}

// List connections
type ListConnectionsRequest_SdkV2 struct {
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

func (newState *ListConnectionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListConnectionsRequest_SdkV2) {
}

func (newState *ListConnectionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListConnectionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListConnectionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListConnectionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListConnectionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListConnectionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListConnectionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListConnectionsResponse_SdkV2 struct {
	// An array of connection information objects.
	Connections types.List `tfsdk:"connections" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListConnectionsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListConnectionsResponse_SdkV2) {
}

func (newState *ListConnectionsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListConnectionsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListConnectionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListConnectionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"connections": reflect.TypeOf(ConnectionInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListConnectionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListConnectionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connections":     o.Connections,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListConnectionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connections": basetypes.ListType{
				ElemType: ConnectionInfo_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetConnections returns the value of the Connections field in ListConnectionsResponse_SdkV2 as
// a slice of ConnectionInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListConnectionsResponse_SdkV2) GetConnections(ctx context.Context) ([]ConnectionInfo_SdkV2, bool) {
	if o.Connections.IsNull() || o.Connections.IsUnknown() {
		return nil, false
	}
	var v []ConnectionInfo_SdkV2
	d := o.Connections.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConnections sets the value of the Connections field in ListConnectionsResponse_SdkV2.
func (o *ListConnectionsResponse_SdkV2) SetConnections(ctx context.Context, v []ConnectionInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["connections"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Connections = types.ListValueMust(t, vs)
}

// List credentials
type ListCredentialsRequest_SdkV2 struct {
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

func (newState *ListCredentialsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCredentialsRequest_SdkV2) {
}

func (newState *ListCredentialsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListCredentialsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
			"purpose":     o.Purpose,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
			"purpose":     types.StringType,
		},
	}
}

type ListCredentialsResponse_SdkV2 struct {
	Credentials types.List `tfsdk:"credentials" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListCredentialsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCredentialsResponse_SdkV2) {
}

func (newState *ListCredentialsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListCredentialsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credentials": reflect.TypeOf(CredentialInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials":     o.Credentials,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials": basetypes.ListType{
				ElemType: CredentialInfo_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetCredentials returns the value of the Credentials field in ListCredentialsResponse_SdkV2 as
// a slice of CredentialInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCredentialsResponse_SdkV2) GetCredentials(ctx context.Context) ([]CredentialInfo_SdkV2, bool) {
	if o.Credentials.IsNull() || o.Credentials.IsUnknown() {
		return nil, false
	}
	var v []CredentialInfo_SdkV2
	d := o.Credentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCredentials sets the value of the Credentials field in ListCredentialsResponse_SdkV2.
func (o *ListCredentialsResponse_SdkV2) SetCredentials(ctx context.Context, v []CredentialInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["credentials"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Credentials = types.ListValueMust(t, vs)
}

// List external locations
type ListExternalLocationsRequest_SdkV2 struct {
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

func (newState *ListExternalLocationsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExternalLocationsRequest_SdkV2) {
}

func (newState *ListExternalLocationsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListExternalLocationsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExternalLocationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExternalLocationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExternalLocationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListExternalLocationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"max_results":    o.MaxResults,
			"page_token":     o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExternalLocationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
		},
	}
}

type ListExternalLocationsResponse_SdkV2 struct {
	// An array of external locations.
	ExternalLocations types.List `tfsdk:"external_locations" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExternalLocationsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExternalLocationsResponse_SdkV2) {
}

func (newState *ListExternalLocationsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListExternalLocationsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExternalLocationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExternalLocationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"external_locations": reflect.TypeOf(ExternalLocationInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExternalLocationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListExternalLocationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_locations": o.ExternalLocations,
			"next_page_token":    o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExternalLocationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_locations": basetypes.ListType{
				ElemType: ExternalLocationInfo_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExternalLocations returns the value of the ExternalLocations field in ListExternalLocationsResponse_SdkV2 as
// a slice of ExternalLocationInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListExternalLocationsResponse_SdkV2) GetExternalLocations(ctx context.Context) ([]ExternalLocationInfo_SdkV2, bool) {
	if o.ExternalLocations.IsNull() || o.ExternalLocations.IsUnknown() {
		return nil, false
	}
	var v []ExternalLocationInfo_SdkV2
	d := o.ExternalLocations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalLocations sets the value of the ExternalLocations field in ListExternalLocationsResponse_SdkV2.
func (o *ListExternalLocationsResponse_SdkV2) SetExternalLocations(ctx context.Context, v []ExternalLocationInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["external_locations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExternalLocations = types.ListValueMust(t, vs)
}

// List functions
type ListFunctionsRequest_SdkV2 struct {
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

func (newState *ListFunctionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFunctionsRequest_SdkV2) {
}

func (newState *ListFunctionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListFunctionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFunctionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFunctionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFunctionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListFunctionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListFunctionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type ListFunctionsResponse_SdkV2 struct {
	// An array of function information objects.
	Functions types.List `tfsdk:"functions" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListFunctionsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFunctionsResponse_SdkV2) {
}

func (newState *ListFunctionsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListFunctionsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFunctionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFunctionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"functions": reflect.TypeOf(FunctionInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFunctionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListFunctionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"functions":       o.Functions,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFunctionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"functions": basetypes.ListType{
				ElemType: FunctionInfo_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFunctions returns the value of the Functions field in ListFunctionsResponse_SdkV2 as
// a slice of FunctionInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListFunctionsResponse_SdkV2) GetFunctions(ctx context.Context) ([]FunctionInfo_SdkV2, bool) {
	if o.Functions.IsNull() || o.Functions.IsUnknown() {
		return nil, false
	}
	var v []FunctionInfo_SdkV2
	d := o.Functions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFunctions sets the value of the Functions field in ListFunctionsResponse_SdkV2.
func (o *ListFunctionsResponse_SdkV2) SetFunctions(ctx context.Context, v []FunctionInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["functions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Functions = types.ListValueMust(t, vs)
}

type ListMetastoresResponse_SdkV2 struct {
	// An array of metastore information objects.
	Metastores types.List `tfsdk:"metastores" tf:"optional"`
}

func (newState *ListMetastoresResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListMetastoresResponse_SdkV2) {
}

func (newState *ListMetastoresResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListMetastoresResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListMetastoresResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListMetastoresResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metastores": reflect.TypeOf(MetastoreInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListMetastoresResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListMetastoresResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastores": o.Metastores,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListMetastoresResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastores": basetypes.ListType{
				ElemType: MetastoreInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetMetastores returns the value of the Metastores field in ListMetastoresResponse_SdkV2 as
// a slice of MetastoreInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListMetastoresResponse_SdkV2) GetMetastores(ctx context.Context) ([]MetastoreInfo_SdkV2, bool) {
	if o.Metastores.IsNull() || o.Metastores.IsUnknown() {
		return nil, false
	}
	var v []MetastoreInfo_SdkV2
	d := o.Metastores.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetastores sets the value of the Metastores field in ListMetastoresResponse_SdkV2.
func (o *ListMetastoresResponse_SdkV2) SetMetastores(ctx context.Context, v []MetastoreInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metastores"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Metastores = types.ListValueMust(t, vs)
}

// List Model Versions
type ListModelVersionsRequest_SdkV2 struct {
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

func (newState *ListModelVersionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelVersionsRequest_SdkV2) {
}

func (newState *ListModelVersionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListModelVersionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListModelVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListModelVersionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelVersionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListModelVersionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListModelVersionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":      types.StringType,
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
		},
	}
}

type ListModelVersionsResponse_SdkV2 struct {
	ModelVersions types.List `tfsdk:"model_versions" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListModelVersionsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelVersionsResponse_SdkV2) {
}

func (newState *ListModelVersionsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListModelVersionsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListModelVersionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListModelVersionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(ModelVersionInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelVersionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListModelVersionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_versions":  o.ModelVersions,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListModelVersionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_versions": basetypes.ListType{
				ElemType: ModelVersionInfo_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetModelVersions returns the value of the ModelVersions field in ListModelVersionsResponse_SdkV2 as
// a slice of ModelVersionInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListModelVersionsResponse_SdkV2) GetModelVersions(ctx context.Context) ([]ModelVersionInfo_SdkV2, bool) {
	if o.ModelVersions.IsNull() || o.ModelVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionInfo_SdkV2
	d := o.ModelVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersions sets the value of the ModelVersions field in ListModelVersionsResponse_SdkV2.
func (o *ListModelVersionsResponse_SdkV2) SetModelVersions(ctx context.Context, v []ModelVersionInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ModelVersions = types.ListValueMust(t, vs)
}

// List all resource quotas under a metastore.
type ListQuotasRequest_SdkV2 struct {
	// The number of quotas to return.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque token for the next page of results.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListQuotasRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQuotasRequest_SdkV2) {
}

func (newState *ListQuotasRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListQuotasRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQuotasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQuotasRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQuotasRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListQuotasRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQuotasRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListQuotasResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of returned QuotaInfos.
	Quotas types.List `tfsdk:"quotas" tf:"optional"`
}

func (newState *ListQuotasResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQuotasResponse_SdkV2) {
}

func (newState *ListQuotasResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListQuotasResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListQuotasResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListQuotasResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"quotas": reflect.TypeOf(QuotaInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListQuotasResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListQuotasResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"quotas":          o.Quotas,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListQuotasResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"quotas": basetypes.ListType{
				ElemType: QuotaInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetQuotas returns the value of the Quotas field in ListQuotasResponse_SdkV2 as
// a slice of QuotaInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListQuotasResponse_SdkV2) GetQuotas(ctx context.Context) ([]QuotaInfo_SdkV2, bool) {
	if o.Quotas.IsNull() || o.Quotas.IsUnknown() {
		return nil, false
	}
	var v []QuotaInfo_SdkV2
	d := o.Quotas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQuotas sets the value of the Quotas field in ListQuotasResponse_SdkV2.
func (o *ListQuotasResponse_SdkV2) SetQuotas(ctx context.Context, v []QuotaInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["quotas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Quotas = types.ListValueMust(t, vs)
}

// List refreshes
type ListRefreshesRequest_SdkV2 struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *ListRefreshesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRefreshesRequest_SdkV2) {
}

func (newState *ListRefreshesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListRefreshesRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRefreshesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRefreshesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRefreshesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListRefreshesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRefreshesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name": types.StringType,
		},
	}
}

// List Registered Models
type ListRegisteredModelsRequest_SdkV2 struct {
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

func (newState *ListRegisteredModelsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRegisteredModelsRequest_SdkV2) {
}

func (newState *ListRegisteredModelsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListRegisteredModelsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRegisteredModelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRegisteredModelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRegisteredModelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListRegisteredModelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListRegisteredModelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type ListRegisteredModelsResponse_SdkV2 struct {
	// Opaque token for pagination. Omitted if there are no more results.
	// page_token should be set to this value for fetching the next page.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	RegisteredModels types.List `tfsdk:"registered_models" tf:"optional"`
}

func (newState *ListRegisteredModelsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRegisteredModelsResponse_SdkV2) {
}

func (newState *ListRegisteredModelsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListRegisteredModelsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRegisteredModelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListRegisteredModelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_models": reflect.TypeOf(RegisteredModelInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRegisteredModelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListRegisteredModelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"registered_models": o.RegisteredModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRegisteredModelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"registered_models": basetypes.ListType{
				ElemType: RegisteredModelInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRegisteredModels returns the value of the RegisteredModels field in ListRegisteredModelsResponse_SdkV2 as
// a slice of RegisteredModelInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListRegisteredModelsResponse_SdkV2) GetRegisteredModels(ctx context.Context) ([]RegisteredModelInfo_SdkV2, bool) {
	if o.RegisteredModels.IsNull() || o.RegisteredModels.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelInfo_SdkV2
	d := o.RegisteredModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModels sets the value of the RegisteredModels field in ListRegisteredModelsResponse_SdkV2.
func (o *ListRegisteredModelsResponse_SdkV2) SetRegisteredModels(ctx context.Context, v []RegisteredModelInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RegisteredModels = types.ListValueMust(t, vs)
}

// List schemas
type ListSchemasRequest_SdkV2 struct {
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

func (newState *ListSchemasRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchemasRequest_SdkV2) {
}

func (newState *ListSchemasRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSchemasRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSchemasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSchemasRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSchemasRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSchemasRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListSchemasRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":   types.StringType,
			"include_browse": types.BoolType,
			"max_results":    types.Int64Type,
			"page_token":     types.StringType,
		},
	}
}

type ListSchemasResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of schema information objects.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
}

func (newState *ListSchemasResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchemasResponse_SdkV2) {
}

func (newState *ListSchemasResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSchemasResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSchemasResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSchemasResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schemas": reflect.TypeOf(SchemaInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSchemasResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSchemasResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"schemas":         o.Schemas,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSchemasResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"schemas": basetypes.ListType{
				ElemType: SchemaInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSchemas returns the value of the Schemas field in ListSchemasResponse_SdkV2 as
// a slice of SchemaInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSchemasResponse_SdkV2) GetSchemas(ctx context.Context) ([]SchemaInfo_SdkV2, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []SchemaInfo_SdkV2
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in ListSchemasResponse_SdkV2.
func (o *ListSchemasResponse_SdkV2) SetSchemas(ctx context.Context, v []SchemaInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

// List credentials
type ListStorageCredentialsRequest_SdkV2 struct {
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

func (newState *ListStorageCredentialsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListStorageCredentialsRequest_SdkV2) {
}

func (newState *ListStorageCredentialsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListStorageCredentialsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStorageCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListStorageCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStorageCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListStorageCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListStorageCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListStorageCredentialsResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	StorageCredentials types.List `tfsdk:"storage_credentials" tf:"optional"`
}

func (newState *ListStorageCredentialsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListStorageCredentialsResponse_SdkV2) {
}

func (newState *ListStorageCredentialsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListStorageCredentialsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStorageCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListStorageCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"storage_credentials": reflect.TypeOf(StorageCredentialInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStorageCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListStorageCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":     o.NextPageToken,
			"storage_credentials": o.StorageCredentials,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListStorageCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"storage_credentials": basetypes.ListType{
				ElemType: StorageCredentialInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetStorageCredentials returns the value of the StorageCredentials field in ListStorageCredentialsResponse_SdkV2 as
// a slice of StorageCredentialInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListStorageCredentialsResponse_SdkV2) GetStorageCredentials(ctx context.Context) ([]StorageCredentialInfo_SdkV2, bool) {
	if o.StorageCredentials.IsNull() || o.StorageCredentials.IsUnknown() {
		return nil, false
	}
	var v []StorageCredentialInfo_SdkV2
	d := o.StorageCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStorageCredentials sets the value of the StorageCredentials field in ListStorageCredentialsResponse_SdkV2.
func (o *ListStorageCredentialsResponse_SdkV2) SetStorageCredentials(ctx context.Context, v []StorageCredentialInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["storage_credentials"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.StorageCredentials = types.ListValueMust(t, vs)
}

// List table summaries
type ListSummariesRequest_SdkV2 struct {
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

func (newState *ListSummariesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSummariesRequest_SdkV2) {
}

func (newState *ListSummariesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSummariesRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSummariesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSummariesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSummariesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSummariesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListSummariesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
type ListSystemSchemasRequest_SdkV2 struct {
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

func (newState *ListSystemSchemasRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSystemSchemasRequest_SdkV2) {
}

func (newState *ListSystemSchemasRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSystemSchemasRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSystemSchemasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSystemSchemasRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSystemSchemasRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSystemSchemasRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results":  o.MaxResults,
			"metastore_id": o.MetastoreId,
			"page_token":   o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSystemSchemasRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results":  types.Int64Type,
			"metastore_id": types.StringType,
			"page_token":   types.StringType,
		},
	}
}

type ListSystemSchemasResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of system schema information objects.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
}

func (newState *ListSystemSchemasResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSystemSchemasResponse_SdkV2) {
}

func (newState *ListSystemSchemasResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListSystemSchemasResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSystemSchemasResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSystemSchemasResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schemas": reflect.TypeOf(SystemSchemaInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSystemSchemasResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSystemSchemasResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"schemas":         o.Schemas,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSystemSchemasResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"schemas": basetypes.ListType{
				ElemType: SystemSchemaInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSchemas returns the value of the Schemas field in ListSystemSchemasResponse_SdkV2 as
// a slice of SystemSchemaInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSystemSchemasResponse_SdkV2) GetSchemas(ctx context.Context) ([]SystemSchemaInfo_SdkV2, bool) {
	if o.Schemas.IsNull() || o.Schemas.IsUnknown() {
		return nil, false
	}
	var v []SystemSchemaInfo_SdkV2
	d := o.Schemas.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchemas sets the value of the Schemas field in ListSystemSchemasResponse_SdkV2.
func (o *ListSystemSchemasResponse_SdkV2) SetSchemas(ctx context.Context, v []SystemSchemaInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schemas"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Schemas = types.ListValueMust(t, vs)
}

type ListTableSummariesResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// List of table summaries.
	Tables types.List `tfsdk:"tables" tf:"optional"`
}

func (newState *ListTableSummariesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTableSummariesResponse_SdkV2) {
}

func (newState *ListTableSummariesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListTableSummariesResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTableSummariesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTableSummariesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tables": reflect.TypeOf(TableSummary_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTableSummariesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTableSummariesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"tables":          o.Tables,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTableSummariesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tables": basetypes.ListType{
				ElemType: TableSummary_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTables returns the value of the Tables field in ListTableSummariesResponse_SdkV2 as
// a slice of TableSummary_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListTableSummariesResponse_SdkV2) GetTables(ctx context.Context) ([]TableSummary_SdkV2, bool) {
	if o.Tables.IsNull() || o.Tables.IsUnknown() {
		return nil, false
	}
	var v []TableSummary_SdkV2
	d := o.Tables.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTables sets the value of the Tables field in ListTableSummariesResponse_SdkV2.
func (o *ListTableSummariesResponse_SdkV2) SetTables(ctx context.Context, v []TableSummary_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tables"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tables = types.ListValueMust(t, vs)
}

// List tables
type ListTablesRequest_SdkV2 struct {
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

func (newState *ListTablesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTablesRequest_SdkV2) {
}

func (newState *ListTablesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListTablesRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTablesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTablesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTablesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTablesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListTablesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type ListTablesResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of table information objects.
	Tables types.List `tfsdk:"tables" tf:"optional"`
}

func (newState *ListTablesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTablesResponse_SdkV2) {
}

func (newState *ListTablesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListTablesResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTablesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTablesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tables": reflect.TypeOf(TableInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTablesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTablesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"tables":          o.Tables,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTablesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tables": basetypes.ListType{
				ElemType: TableInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTables returns the value of the Tables field in ListTablesResponse_SdkV2 as
// a slice of TableInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListTablesResponse_SdkV2) GetTables(ctx context.Context) ([]TableInfo_SdkV2, bool) {
	if o.Tables.IsNull() || o.Tables.IsUnknown() {
		return nil, false
	}
	var v []TableInfo_SdkV2
	d := o.Tables.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTables sets the value of the Tables field in ListTablesResponse_SdkV2.
func (o *ListTablesResponse_SdkV2) SetTables(ctx context.Context, v []TableInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tables"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tables = types.ListValueMust(t, vs)
}

// List Volumes
type ListVolumesRequest_SdkV2 struct {
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

func (newState *ListVolumesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVolumesRequest_SdkV2) {
}

func (newState *ListVolumesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListVolumesRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVolumesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVolumesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVolumesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListVolumesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListVolumesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type ListVolumesResponseContent_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request to retrieve the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Volumes types.List `tfsdk:"volumes" tf:"optional"`
}

func (newState *ListVolumesResponseContent_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVolumesResponseContent_SdkV2) {
}

func (newState *ListVolumesResponseContent_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListVolumesResponseContent_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVolumesResponseContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVolumesResponseContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"volumes": reflect.TypeOf(VolumeInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVolumesResponseContent_SdkV2
// only implements ToObjectValue() and Type().
func (o ListVolumesResponseContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"volumes":         o.Volumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListVolumesResponseContent_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"volumes": basetypes.ListType{
				ElemType: VolumeInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetVolumes returns the value of the Volumes field in ListVolumesResponseContent_SdkV2 as
// a slice of VolumeInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListVolumesResponseContent_SdkV2) GetVolumes(ctx context.Context) ([]VolumeInfo_SdkV2, bool) {
	if o.Volumes.IsNull() || o.Volumes.IsUnknown() {
		return nil, false
	}
	var v []VolumeInfo_SdkV2
	d := o.Volumes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolumes sets the value of the Volumes field in ListVolumesResponseContent_SdkV2.
func (o *ListVolumesResponseContent_SdkV2) SetVolumes(ctx context.Context, v []VolumeInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["volumes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Volumes = types.ListValueMust(t, vs)
}

type MetastoreAssignment_SdkV2 struct {
	// The name of the default catalog in the metastore.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name" tf:"optional"`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:""`
	// The unique ID of the Databricks workspace.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:""`
}

func (newState *MetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MetastoreAssignment_SdkV2) {
}

func (newState *MetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringRead(existingState MetastoreAssignment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MetastoreAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MetastoreAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o MetastoreAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_catalog_name": o.DefaultCatalogName,
			"metastore_id":         o.MetastoreId,
			"workspace_id":         o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MetastoreAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_catalog_name": types.StringType,
			"metastore_id":         types.StringType,
			"workspace_id":         types.Int64Type,
		},
	}
}

type MetastoreInfo_SdkV2 struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of metastore creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId types.String `tfsdk:"default_data_access_config_id" tf:"optional"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName types.String `tfsdk:"delta_sharing_organization_name" tf:"optional"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds types.Int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds" tf:"optional"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope types.String `tfsdk:"delta_sharing_scope" tf:"optional"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled types.Bool `tfsdk:"external_access_enabled" tf:"optional"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id" tf:"optional"`
	// Unique identifier of metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The user-specified name of the metastore.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The owner of the metastore.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion types.String `tfsdk:"privilege_model_version" tf:"optional"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region types.String `tfsdk:"region" tf:"optional"`
	// The storage root URL for metastore
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId types.String `tfsdk:"storage_root_credential_id" tf:"optional"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName types.String `tfsdk:"storage_root_credential_name" tf:"optional"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the metastore.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *MetastoreInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MetastoreInfo_SdkV2) {
}

func (newState *MetastoreInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState MetastoreInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MetastoreInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MetastoreInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MetastoreInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o MetastoreInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o MetastoreInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

type ModelVersionInfo_SdkV2 struct {
	// List of aliases associated with the model version
	Aliases types.List `tfsdk:"aliases" tf:"optional"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// The name of the catalog containing the model version
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The comment attached to the model version
	Comment types.String `tfsdk:"comment" tf:"optional"`

	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The identifier of the user who created the model version
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The unique identifier of the model version
	Id types.String `tfsdk:"id" tf:"optional"`
	// The unique identifier of the metastore containing the model version
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The name of the parent registered model of the model version, relative to
	// parent schema
	ModelName types.String `tfsdk:"model_name" tf:"optional"`
	// Model version dependencies, for feature-store packaged models
	ModelVersionDependencies types.List `tfsdk:"model_version_dependencies" tf:"optional,object"`
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	RunWorkspaceId types.Int64 `tfsdk:"run_workspace_id" tf:"optional"`
	// The name of the schema containing the model version, relative to parent
	// catalog
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// URI indicating the location of the source artifacts (files) for the model
	// version
	Source types.String `tfsdk:"source" tf:"optional"`
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	Status types.String `tfsdk:"status" tf:"optional"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// The identifier of the user who updated the model version last time
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// Integer model version number, used to reference the model version in API
	// requests.
	Version types.Int64 `tfsdk:"version" tf:"optional"`
}

func (newState *ModelVersionInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersionInfo_SdkV2) {
}

func (newState *ModelVersionInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ModelVersionInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelVersionInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ModelVersionInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aliases":                    reflect.TypeOf(RegisteredModelAlias_SdkV2{}),
		"model_version_dependencies": reflect.TypeOf(DependencyList_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersionInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ModelVersionInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ModelVersionInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aliases": basetypes.ListType{
				ElemType: RegisteredModelAlias_SdkV2{}.Type(ctx),
			},
			"browse_only":  types.BoolType,
			"catalog_name": types.StringType,
			"comment":      types.StringType,
			"created_at":   types.Int64Type,
			"created_by":   types.StringType,
			"id":           types.StringType,
			"metastore_id": types.StringType,
			"model_name":   types.StringType,
			"model_version_dependencies": basetypes.ListType{
				ElemType: DependencyList_SdkV2{}.Type(ctx),
			},
			"run_id":           types.StringType,
			"run_workspace_id": types.Int64Type,
			"schema_name":      types.StringType,
			"source":           types.StringType,
			"status":           types.StringType,
			"storage_location": types.StringType,
			"updated_at":       types.Int64Type,
			"updated_by":       types.StringType,
			"version":          types.Int64Type,
		},
	}
}

// GetAliases returns the value of the Aliases field in ModelVersionInfo_SdkV2 as
// a slice of RegisteredModelAlias_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelVersionInfo_SdkV2) GetAliases(ctx context.Context) ([]RegisteredModelAlias_SdkV2, bool) {
	if o.Aliases.IsNull() || o.Aliases.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAlias_SdkV2
	d := o.Aliases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAliases sets the value of the Aliases field in ModelVersionInfo_SdkV2.
func (o *ModelVersionInfo_SdkV2) SetAliases(ctx context.Context, v []RegisteredModelAlias_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Aliases = types.ListValueMust(t, vs)
}

// GetModelVersionDependencies returns the value of the ModelVersionDependencies field in ModelVersionInfo_SdkV2 as
// a DependencyList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelVersionInfo_SdkV2) GetModelVersionDependencies(ctx context.Context) (DependencyList_SdkV2, bool) {
	var e DependencyList_SdkV2
	if o.ModelVersionDependencies.IsNull() || o.ModelVersionDependencies.IsUnknown() {
		return e, false
	}
	var v []DependencyList_SdkV2
	d := o.ModelVersionDependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersionDependencies sets the value of the ModelVersionDependencies field in ModelVersionInfo_SdkV2.
func (o *ModelVersionInfo_SdkV2) SetModelVersionDependencies(ctx context.Context, v DependencyList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_version_dependencies"]
	o.ModelVersionDependencies = types.ListValueMust(t, vs)
}

type MonitorCronSchedule_SdkV2 struct {
	// Read only field that indicates whether a schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression types.String `tfsdk:"quartz_cron_expression" tf:""`
	// The timezone id (e.g., ``"PST"``) in which to evaluate the quartz
	// expression.
	TimezoneId types.String `tfsdk:"timezone_id" tf:""`
}

func (newState *MonitorCronSchedule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorCronSchedule_SdkV2) {
}

func (newState *MonitorCronSchedule_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorCronSchedule_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorCronSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorCronSchedule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorCronSchedule_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorCronSchedule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pause_status":           o.PauseStatus,
			"quartz_cron_expression": o.QuartzCronExpression,
			"timezone_id":            o.TimezoneId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorCronSchedule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pause_status":           types.StringType,
			"quartz_cron_expression": types.StringType,
			"timezone_id":            types.StringType,
		},
	}
}

type MonitorDataClassificationConfig_SdkV2 struct {
	// Whether data classification is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
}

func (newState *MonitorDataClassificationConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorDataClassificationConfig_SdkV2) {
}

func (newState *MonitorDataClassificationConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorDataClassificationConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorDataClassificationConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorDataClassificationConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorDataClassificationConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorDataClassificationConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled": o.Enabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorDataClassificationConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

type MonitorDestination_SdkV2 struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	EmailAddresses types.List `tfsdk:"email_addresses" tf:"optional"`
}

func (newState *MonitorDestination_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorDestination_SdkV2) {
}

func (newState *MonitorDestination_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorDestination_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorDestination_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email_addresses": o.EmailAddresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorDestination_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetEmailAddresses returns the value of the EmailAddresses field in MonitorDestination_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorDestination_SdkV2) GetEmailAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetEmailAddresses sets the value of the EmailAddresses field in MonitorDestination_SdkV2.
func (o *MonitorDestination_SdkV2) SetEmailAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmailAddresses = types.ListValueMust(t, vs)
}

type MonitorInferenceLog_SdkV2 struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities types.List `tfsdk:"granularities" tf:""`
	// Optional column that contains the ground truth for the prediction.
	LabelCol types.String `tfsdk:"label_col" tf:"optional"`
	// Column that contains the id of the model generating the predictions.
	// Metrics will be computed per model id by default, and also across all
	// model ids.
	ModelIdCol types.String `tfsdk:"model_id_col" tf:""`
	// Column that contains the output/prediction from the model.
	PredictionCol types.String `tfsdk:"prediction_col" tf:""`
	// Optional column that contains the prediction probabilities for each class
	// in a classification problem type. The values in this column should be a
	// map, mapping each class label to the prediction probability for a given
	// sample. The map should be of PySpark MapType().
	PredictionProbaCol types.String `tfsdk:"prediction_proba_col" tf:"optional"`
	// Problem type the model aims to solve. Determines the type of
	// model-quality metrics that will be computed.
	ProblemType types.String `tfsdk:"problem_type" tf:""`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol types.String `tfsdk:"timestamp_col" tf:""`
}

func (newState *MonitorInferenceLog_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorInferenceLog_SdkV2) {
}

func (newState *MonitorInferenceLog_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorInferenceLog_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorInferenceLog.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorInferenceLog_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"granularities": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorInferenceLog_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorInferenceLog_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o MonitorInferenceLog_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetGranularities returns the value of the Granularities field in MonitorInferenceLog_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInferenceLog_SdkV2) GetGranularities(ctx context.Context) ([]types.String, bool) {
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

// SetGranularities sets the value of the Granularities field in MonitorInferenceLog_SdkV2.
func (o *MonitorInferenceLog_SdkV2) SetGranularities(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["granularities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Granularities = types.ListValueMust(t, vs)
}

type MonitorInfo_SdkV2 struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir types.String `tfsdk:"assets_dir" tf:"optional"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName types.String `tfsdk:"baseline_table_name" tf:"optional"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics types.List `tfsdk:"custom_metrics" tf:"optional"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The data classification config for the monitor.
	DataClassificationConfig types.List `tfsdk:"data_classification_config" tf:"optional,object"`
	// The full name of the drift metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	DriftMetricsTableName types.String `tfsdk:"drift_metrics_table_name" tf:""`
	// Configuration for monitoring inference logs.
	InferenceLog types.List `tfsdk:"inference_log" tf:"optional,object"`
	// The latest failure message of the monitor (if any).
	LatestMonitorFailureMsg types.String `tfsdk:"latest_monitor_failure_msg" tf:"optional"`
	// The version of the monitor config (e.g. 1,2,3). If negative, the monitor
	// may be corrupted.
	MonitorVersion types.String `tfsdk:"monitor_version" tf:""`
	// The notification settings for the monitor.
	Notifications types.List `tfsdk:"notifications" tf:"optional,object"`
	// Schema where output metric tables are created.
	OutputSchemaName types.String `tfsdk:"output_schema_name" tf:"optional"`
	// The full name of the profile metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	ProfileMetricsTableName types.String `tfsdk:"profile_metrics_table_name" tf:""`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs types.List `tfsdk:"slicing_exprs" tf:"optional"`
	// Configuration for monitoring snapshot tables.
	Snapshot types.List `tfsdk:"snapshot" tf:"optional,object"`
	// The status of the monitor.
	Status types.String `tfsdk:"status" tf:""`
	// The full name of the table to monitor. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	TableName types.String `tfsdk:"table_name" tf:""`
	// Configuration for monitoring time series tables.
	TimeSeries types.List `tfsdk:"time_series" tf:"optional,object"`
}

func (newState *MonitorInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorInfo_SdkV2) {
}

func (newState *MonitorInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_metrics":             reflect.TypeOf(MonitorMetric_SdkV2{}),
		"data_classification_config": reflect.TypeOf(MonitorDataClassificationConfig_SdkV2{}),
		"inference_log":              reflect.TypeOf(MonitorInferenceLog_SdkV2{}),
		"notifications":              reflect.TypeOf(MonitorNotifications_SdkV2{}),
		"schedule":                   reflect.TypeOf(MonitorCronSchedule_SdkV2{}),
		"slicing_exprs":              reflect.TypeOf(types.String{}),
		"snapshot":                   reflect.TypeOf(MonitorSnapshot_SdkV2{}),
		"time_series":                reflect.TypeOf(MonitorTimeSeries_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o MonitorInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets_dir":          types.StringType,
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: MonitorMetric_SdkV2{}.Type(ctx),
			},
			"dashboard_id": types.StringType,
			"data_classification_config": basetypes.ListType{
				ElemType: MonitorDataClassificationConfig_SdkV2{}.Type(ctx),
			},
			"drift_metrics_table_name": types.StringType,
			"inference_log": basetypes.ListType{
				ElemType: MonitorInferenceLog_SdkV2{}.Type(ctx),
			},
			"latest_monitor_failure_msg": types.StringType,
			"monitor_version":            types.StringType,
			"notifications": basetypes.ListType{
				ElemType: MonitorNotifications_SdkV2{}.Type(ctx),
			},
			"output_schema_name":         types.StringType,
			"profile_metrics_table_name": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: MonitorCronSchedule_SdkV2{}.Type(ctx),
			},
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot": basetypes.ListType{
				ElemType: MonitorSnapshot_SdkV2{}.Type(ctx),
			},
			"status":     types.StringType,
			"table_name": types.StringType,
			"time_series": basetypes.ListType{
				ElemType: MonitorTimeSeries_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetCustomMetrics returns the value of the CustomMetrics field in MonitorInfo_SdkV2 as
// a slice of MonitorMetric_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo_SdkV2) GetCustomMetrics(ctx context.Context) ([]MonitorMetric_SdkV2, bool) {
	if o.CustomMetrics.IsNull() || o.CustomMetrics.IsUnknown() {
		return nil, false
	}
	var v []MonitorMetric_SdkV2
	d := o.CustomMetrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomMetrics sets the value of the CustomMetrics field in MonitorInfo_SdkV2.
func (o *MonitorInfo_SdkV2) SetCustomMetrics(ctx context.Context, v []MonitorMetric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomMetrics = types.ListValueMust(t, vs)
}

// GetDataClassificationConfig returns the value of the DataClassificationConfig field in MonitorInfo_SdkV2 as
// a MonitorDataClassificationConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo_SdkV2) GetDataClassificationConfig(ctx context.Context) (MonitorDataClassificationConfig_SdkV2, bool) {
	var e MonitorDataClassificationConfig_SdkV2
	if o.DataClassificationConfig.IsNull() || o.DataClassificationConfig.IsUnknown() {
		return e, false
	}
	var v []MonitorDataClassificationConfig_SdkV2
	d := o.DataClassificationConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataClassificationConfig sets the value of the DataClassificationConfig field in MonitorInfo_SdkV2.
func (o *MonitorInfo_SdkV2) SetDataClassificationConfig(ctx context.Context, v MonitorDataClassificationConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_classification_config"]
	o.DataClassificationConfig = types.ListValueMust(t, vs)
}

// GetInferenceLog returns the value of the InferenceLog field in MonitorInfo_SdkV2 as
// a MonitorInferenceLog_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo_SdkV2) GetInferenceLog(ctx context.Context) (MonitorInferenceLog_SdkV2, bool) {
	var e MonitorInferenceLog_SdkV2
	if o.InferenceLog.IsNull() || o.InferenceLog.IsUnknown() {
		return e, false
	}
	var v []MonitorInferenceLog_SdkV2
	d := o.InferenceLog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInferenceLog sets the value of the InferenceLog field in MonitorInfo_SdkV2.
func (o *MonitorInfo_SdkV2) SetInferenceLog(ctx context.Context, v MonitorInferenceLog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inference_log"]
	o.InferenceLog = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in MonitorInfo_SdkV2 as
// a MonitorNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo_SdkV2) GetNotifications(ctx context.Context) (MonitorNotifications_SdkV2, bool) {
	var e MonitorNotifications_SdkV2
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return e, false
	}
	var v []MonitorNotifications_SdkV2
	d := o.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotifications sets the value of the Notifications field in MonitorInfo_SdkV2.
func (o *MonitorInfo_SdkV2) SetNotifications(ctx context.Context, v MonitorNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	o.Notifications = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in MonitorInfo_SdkV2 as
// a MonitorCronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo_SdkV2) GetSchedule(ctx context.Context) (MonitorCronSchedule_SdkV2, bool) {
	var e MonitorCronSchedule_SdkV2
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []MonitorCronSchedule_SdkV2
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in MonitorInfo_SdkV2.
func (o *MonitorInfo_SdkV2) SetSchedule(ctx context.Context, v MonitorCronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}

// GetSlicingExprs returns the value of the SlicingExprs field in MonitorInfo_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo_SdkV2) GetSlicingExprs(ctx context.Context) ([]types.String, bool) {
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

// SetSlicingExprs sets the value of the SlicingExprs field in MonitorInfo_SdkV2.
func (o *MonitorInfo_SdkV2) SetSlicingExprs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["slicing_exprs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SlicingExprs = types.ListValueMust(t, vs)
}

// GetSnapshot returns the value of the Snapshot field in MonitorInfo_SdkV2 as
// a MonitorSnapshot_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo_SdkV2) GetSnapshot(ctx context.Context) (MonitorSnapshot_SdkV2, bool) {
	var e MonitorSnapshot_SdkV2
	if o.Snapshot.IsNull() || o.Snapshot.IsUnknown() {
		return e, false
	}
	var v []MonitorSnapshot_SdkV2
	d := o.Snapshot.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSnapshot sets the value of the Snapshot field in MonitorInfo_SdkV2.
func (o *MonitorInfo_SdkV2) SetSnapshot(ctx context.Context, v MonitorSnapshot_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["snapshot"]
	o.Snapshot = types.ListValueMust(t, vs)
}

// GetTimeSeries returns the value of the TimeSeries field in MonitorInfo_SdkV2 as
// a MonitorTimeSeries_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorInfo_SdkV2) GetTimeSeries(ctx context.Context) (MonitorTimeSeries_SdkV2, bool) {
	var e MonitorTimeSeries_SdkV2
	if o.TimeSeries.IsNull() || o.TimeSeries.IsUnknown() {
		return e, false
	}
	var v []MonitorTimeSeries_SdkV2
	d := o.TimeSeries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTimeSeries sets the value of the TimeSeries field in MonitorInfo_SdkV2.
func (o *MonitorInfo_SdkV2) SetTimeSeries(ctx context.Context, v MonitorTimeSeries_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["time_series"]
	o.TimeSeries = types.ListValueMust(t, vs)
}

type MonitorMetric_SdkV2 struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	Definition types.String `tfsdk:"definition" tf:""`
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	InputColumns types.List `tfsdk:"input_columns" tf:""`
	// Name of the metric in the output tables.
	Name types.String `tfsdk:"name" tf:""`
	// The output type of the custom metric.
	OutputDataType types.String `tfsdk:"output_data_type" tf:""`
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
	Type_ types.String `tfsdk:"type" tf:""`
}

func (newState *MonitorMetric_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorMetric_SdkV2) {
}

func (newState *MonitorMetric_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorMetric_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorMetric.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorMetric_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorMetric_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorMetric_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o MonitorMetric_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInputColumns returns the value of the InputColumns field in MonitorMetric_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorMetric_SdkV2) GetInputColumns(ctx context.Context) ([]types.String, bool) {
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

// SetInputColumns sets the value of the InputColumns field in MonitorMetric_SdkV2.
func (o *MonitorMetric_SdkV2) SetInputColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["input_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InputColumns = types.ListValueMust(t, vs)
}

type MonitorNotifications_SdkV2 struct {
	// Who to send notifications to on monitor failure.
	OnFailure types.List `tfsdk:"on_failure" tf:"optional,object"`
	// Who to send notifications to when new data classification tags are
	// detected.
	OnNewClassificationTagDetected types.List `tfsdk:"on_new_classification_tag_detected" tf:"optional,object"`
}

func (newState *MonitorNotifications_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorNotifications_SdkV2) {
}

func (newState *MonitorNotifications_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorNotifications_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorNotifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorNotifications_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_failure":                         reflect.TypeOf(MonitorDestination_SdkV2{}),
		"on_new_classification_tag_detected": reflect.TypeOf(MonitorDestination_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorNotifications_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorNotifications_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"on_failure":                         o.OnFailure,
			"on_new_classification_tag_detected": o.OnNewClassificationTagDetected,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorNotifications_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"on_failure": basetypes.ListType{
				ElemType: MonitorDestination_SdkV2{}.Type(ctx),
			},
			"on_new_classification_tag_detected": basetypes.ListType{
				ElemType: MonitorDestination_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOnFailure returns the value of the OnFailure field in MonitorNotifications_SdkV2 as
// a MonitorDestination_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorNotifications_SdkV2) GetOnFailure(ctx context.Context) (MonitorDestination_SdkV2, bool) {
	var e MonitorDestination_SdkV2
	if o.OnFailure.IsNull() || o.OnFailure.IsUnknown() {
		return e, false
	}
	var v []MonitorDestination_SdkV2
	d := o.OnFailure.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOnFailure sets the value of the OnFailure field in MonitorNotifications_SdkV2.
func (o *MonitorNotifications_SdkV2) SetOnFailure(ctx context.Context, v MonitorDestination_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_failure"]
	o.OnFailure = types.ListValueMust(t, vs)
}

// GetOnNewClassificationTagDetected returns the value of the OnNewClassificationTagDetected field in MonitorNotifications_SdkV2 as
// a MonitorDestination_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorNotifications_SdkV2) GetOnNewClassificationTagDetected(ctx context.Context) (MonitorDestination_SdkV2, bool) {
	var e MonitorDestination_SdkV2
	if o.OnNewClassificationTagDetected.IsNull() || o.OnNewClassificationTagDetected.IsUnknown() {
		return e, false
	}
	var v []MonitorDestination_SdkV2
	d := o.OnNewClassificationTagDetected.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOnNewClassificationTagDetected sets the value of the OnNewClassificationTagDetected field in MonitorNotifications_SdkV2.
func (o *MonitorNotifications_SdkV2) SetOnNewClassificationTagDetected(ctx context.Context, v MonitorDestination_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_new_classification_tag_detected"]
	o.OnNewClassificationTagDetected = types.ListValueMust(t, vs)
}

type MonitorRefreshInfo_SdkV2 struct {
	// Time at which refresh operation completed (milliseconds since 1/1/1970
	// UTC).
	EndTimeMs types.Int64 `tfsdk:"end_time_ms" tf:"optional"`
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	Message types.String `tfsdk:"message" tf:"optional"`
	// Unique id of the refresh operation.
	RefreshId types.Int64 `tfsdk:"refresh_id" tf:""`
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	StartTimeMs types.Int64 `tfsdk:"start_time_ms" tf:""`
	// The current state of the refresh.
	State types.String `tfsdk:"state" tf:""`
	// The method by which the refresh was triggered.
	Trigger types.String `tfsdk:"trigger" tf:"optional"`
}

func (newState *MonitorRefreshInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorRefreshInfo_SdkV2) {
}

func (newState *MonitorRefreshInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorRefreshInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorRefreshInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorRefreshInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorRefreshInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorRefreshInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o MonitorRefreshInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

type MonitorRefreshListResponse_SdkV2 struct {
	// List of refreshes.
	Refreshes types.List `tfsdk:"refreshes" tf:"optional"`
}

func (newState *MonitorRefreshListResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorRefreshListResponse_SdkV2) {
}

func (newState *MonitorRefreshListResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorRefreshListResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorRefreshListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorRefreshListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refreshes": reflect.TypeOf(MonitorRefreshInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorRefreshListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorRefreshListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"refreshes": o.Refreshes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorRefreshListResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"refreshes": basetypes.ListType{
				ElemType: MonitorRefreshInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRefreshes returns the value of the Refreshes field in MonitorRefreshListResponse_SdkV2 as
// a slice of MonitorRefreshInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorRefreshListResponse_SdkV2) GetRefreshes(ctx context.Context) ([]MonitorRefreshInfo_SdkV2, bool) {
	if o.Refreshes.IsNull() || o.Refreshes.IsUnknown() {
		return nil, false
	}
	var v []MonitorRefreshInfo_SdkV2
	d := o.Refreshes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRefreshes sets the value of the Refreshes field in MonitorRefreshListResponse_SdkV2.
func (o *MonitorRefreshListResponse_SdkV2) SetRefreshes(ctx context.Context, v []MonitorRefreshInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["refreshes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Refreshes = types.ListValueMust(t, vs)
}

type MonitorSnapshot_SdkV2 struct {
}

func (newState *MonitorSnapshot_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorSnapshot_SdkV2) {
}

func (newState *MonitorSnapshot_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorSnapshot_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorSnapshot.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorSnapshot_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorSnapshot_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorSnapshot_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorSnapshot_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type MonitorTimeSeries_SdkV2 struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities types.List `tfsdk:"granularities" tf:""`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol types.String `tfsdk:"timestamp_col" tf:""`
}

func (newState *MonitorTimeSeries_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorTimeSeries_SdkV2) {
}

func (newState *MonitorTimeSeries_SdkV2) SyncEffectiveFieldsDuringRead(existingState MonitorTimeSeries_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MonitorTimeSeries.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MonitorTimeSeries_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"granularities": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MonitorTimeSeries_SdkV2
// only implements ToObjectValue() and Type().
func (o MonitorTimeSeries_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"granularities": o.Granularities,
			"timestamp_col": o.TimestampCol,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MonitorTimeSeries_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"granularities": basetypes.ListType{
				ElemType: types.StringType,
			},
			"timestamp_col": types.StringType,
		},
	}
}

// GetGranularities returns the value of the Granularities field in MonitorTimeSeries_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *MonitorTimeSeries_SdkV2) GetGranularities(ctx context.Context) ([]types.String, bool) {
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

// SetGranularities sets the value of the Granularities field in MonitorTimeSeries_SdkV2.
func (o *MonitorTimeSeries_SdkV2) SetGranularities(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["granularities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Granularities = types.ListValueMust(t, vs)
}

type NamedTableConstraint_SdkV2 struct {
	// The name of the constraint.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *NamedTableConstraint_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NamedTableConstraint_SdkV2) {
}

func (newState *NamedTableConstraint_SdkV2) SyncEffectiveFieldsDuringRead(existingState NamedTableConstraint_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NamedTableConstraint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NamedTableConstraint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NamedTableConstraint_SdkV2
// only implements ToObjectValue() and Type().
func (o NamedTableConstraint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NamedTableConstraint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Online Table information.
type OnlineTable_SdkV2 struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Specification of the online table.
	Spec types.List `tfsdk:"spec" tf:"optional,object"`
	// Online Table data synchronization status
	Status types.List `tfsdk:"status" tf:"computed,object"`
	// Data serving REST API URL for this table
	TableServingUrl types.String `tfsdk:"table_serving_url" tf:"computed"`
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState types.String `tfsdk:"unity_catalog_provisioning_state" tf:"computed"`
}

func (newState *OnlineTable_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTable_SdkV2) {
}

func (newState *OnlineTable_SdkV2) SyncEffectiveFieldsDuringRead(existingState OnlineTable_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OnlineTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(OnlineTableSpec_SdkV2{}),
		"status": reflect.TypeOf(OnlineTableStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineTable_SdkV2
// only implements ToObjectValue() and Type().
func (o OnlineTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o OnlineTable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"spec": basetypes.ListType{
				ElemType: OnlineTableSpec_SdkV2{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: OnlineTableStatus_SdkV2{}.Type(ctx),
			},
			"table_serving_url":                types.StringType,
			"unity_catalog_provisioning_state": types.StringType,
		},
	}
}

// GetSpec returns the value of the Spec field in OnlineTable_SdkV2 as
// a OnlineTableSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTable_SdkV2) GetSpec(ctx context.Context) (OnlineTableSpec_SdkV2, bool) {
	var e OnlineTableSpec_SdkV2
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v []OnlineTableSpec_SdkV2
	d := o.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in OnlineTable_SdkV2.
func (o *OnlineTable_SdkV2) SetSpec(ctx context.Context, v OnlineTableSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	o.Spec = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in OnlineTable_SdkV2 as
// a OnlineTableStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTable_SdkV2) GetStatus(ctx context.Context) (OnlineTableStatus_SdkV2, bool) {
	var e OnlineTableStatus_SdkV2
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []OnlineTableStatus_SdkV2
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in OnlineTable_SdkV2.
func (o *OnlineTable_SdkV2) SetStatus(ctx context.Context, v OnlineTableStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
}

// Specification of an online table.
type OnlineTableSpec_SdkV2 struct {
	// Whether to create a full-copy pipeline -- a pipeline that stops after
	// creates a full copy of the source table upon initialization and does not
	// process any change data feeds (CDFs) afterwards. The pipeline can still
	// be manually triggered afterwards, but it always perform a full copy of
	// the source table and there are no incremental updates. This mode is
	// useful for syncing views or tables without CDFs to online tables. Note
	// that the full-copy pipeline only supports "triggered" scheduling policy.
	PerformFullCopy types.Bool `tfsdk:"perform_full_copy" tf:"optional"`
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"computed"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns types.List `tfsdk:"primary_key_columns" tf:"optional"`
	// Pipeline runs continuously after generating the initial data.
	RunContinuously types.List `tfsdk:"run_continuously" tf:"optional,object"`
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	RunTriggered types.List `tfsdk:"run_triggered" tf:"optional,object"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName types.String `tfsdk:"source_table_full_name" tf:"optional"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey types.String `tfsdk:"timeseries_key" tf:"optional"`
}

func (newState *OnlineTableSpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpec_SdkV2) {
}

func (newState *OnlineTableSpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpec_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineTableSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OnlineTableSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_key_columns": reflect.TypeOf(types.String{}),
		"run_continuously":    reflect.TypeOf(OnlineTableSpecContinuousSchedulingPolicy_SdkV2{}),
		"run_triggered":       reflect.TypeOf(OnlineTableSpecTriggeredSchedulingPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineTableSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o OnlineTableSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o OnlineTableSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"perform_full_copy": types.BoolType,
			"pipeline_id":       types.StringType,
			"primary_key_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"run_continuously": basetypes.ListType{
				ElemType: OnlineTableSpecContinuousSchedulingPolicy_SdkV2{}.Type(ctx),
			},
			"run_triggered": basetypes.ListType{
				ElemType: OnlineTableSpecTriggeredSchedulingPolicy_SdkV2{}.Type(ctx),
			},
			"source_table_full_name": types.StringType,
			"timeseries_key":         types.StringType,
		},
	}
}

// GetPrimaryKeyColumns returns the value of the PrimaryKeyColumns field in OnlineTableSpec_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableSpec_SdkV2) GetPrimaryKeyColumns(ctx context.Context) ([]types.String, bool) {
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

// SetPrimaryKeyColumns sets the value of the PrimaryKeyColumns field in OnlineTableSpec_SdkV2.
func (o *OnlineTableSpec_SdkV2) SetPrimaryKeyColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_key_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrimaryKeyColumns = types.ListValueMust(t, vs)
}

// GetRunContinuously returns the value of the RunContinuously field in OnlineTableSpec_SdkV2 as
// a OnlineTableSpecContinuousSchedulingPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableSpec_SdkV2) GetRunContinuously(ctx context.Context) (OnlineTableSpecContinuousSchedulingPolicy_SdkV2, bool) {
	var e OnlineTableSpecContinuousSchedulingPolicy_SdkV2
	if o.RunContinuously.IsNull() || o.RunContinuously.IsUnknown() {
		return e, false
	}
	var v []OnlineTableSpecContinuousSchedulingPolicy_SdkV2
	d := o.RunContinuously.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunContinuously sets the value of the RunContinuously field in OnlineTableSpec_SdkV2.
func (o *OnlineTableSpec_SdkV2) SetRunContinuously(ctx context.Context, v OnlineTableSpecContinuousSchedulingPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_continuously"]
	o.RunContinuously = types.ListValueMust(t, vs)
}

// GetRunTriggered returns the value of the RunTriggered field in OnlineTableSpec_SdkV2 as
// a OnlineTableSpecTriggeredSchedulingPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableSpec_SdkV2) GetRunTriggered(ctx context.Context) (OnlineTableSpecTriggeredSchedulingPolicy_SdkV2, bool) {
	var e OnlineTableSpecTriggeredSchedulingPolicy_SdkV2
	if o.RunTriggered.IsNull() || o.RunTriggered.IsUnknown() {
		return e, false
	}
	var v []OnlineTableSpecTriggeredSchedulingPolicy_SdkV2
	d := o.RunTriggered.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunTriggered sets the value of the RunTriggered field in OnlineTableSpec_SdkV2.
func (o *OnlineTableSpec_SdkV2) SetRunTriggered(ctx context.Context, v OnlineTableSpecTriggeredSchedulingPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_triggered"]
	o.RunTriggered = types.ListValueMust(t, vs)
}

type OnlineTableSpecContinuousSchedulingPolicy_SdkV2 struct {
}

func (newState *OnlineTableSpecContinuousSchedulingPolicy_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpecContinuousSchedulingPolicy_SdkV2) {
}

func (newState *OnlineTableSpecContinuousSchedulingPolicy_SdkV2) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpecContinuousSchedulingPolicy_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineTableSpecContinuousSchedulingPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OnlineTableSpecContinuousSchedulingPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineTableSpecContinuousSchedulingPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (o OnlineTableSpecContinuousSchedulingPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o OnlineTableSpecContinuousSchedulingPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type OnlineTableSpecTriggeredSchedulingPolicy_SdkV2 struct {
}

func (newState *OnlineTableSpecTriggeredSchedulingPolicy_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpecTriggeredSchedulingPolicy_SdkV2) {
}

func (newState *OnlineTableSpecTriggeredSchedulingPolicy_SdkV2) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpecTriggeredSchedulingPolicy_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineTableSpecTriggeredSchedulingPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OnlineTableSpecTriggeredSchedulingPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineTableSpecTriggeredSchedulingPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (o OnlineTableSpecTriggeredSchedulingPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o OnlineTableSpecTriggeredSchedulingPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Status of an online table.
type OnlineTableStatus_SdkV2 struct {
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
	ContinuousUpdateStatus types.List `tfsdk:"continuous_update_status" tf:"optional,object"`
	// The state of the online table.
	DetailedState types.String `tfsdk:"detailed_state" tf:"optional"`
	// Detailed status of an online table. Shown if the online table is in the
	// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
	FailedStatus types.List `tfsdk:"failed_status" tf:"optional,object"`
	// A text description of the current state of the online table.
	Message types.String `tfsdk:"message" tf:"optional"`
	// Detailed status of an online table. Shown if the online table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	ProvisioningStatus types.List `tfsdk:"provisioning_status" tf:"optional,object"`
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
	TriggeredUpdateStatus types.List `tfsdk:"triggered_update_status" tf:"optional,object"`
}

func (newState *OnlineTableStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableStatus_SdkV2) {
}

func (newState *OnlineTableStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState OnlineTableStatus_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineTableStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OnlineTableStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"continuous_update_status": reflect.TypeOf(ContinuousUpdateStatus_SdkV2{}),
		"failed_status":            reflect.TypeOf(FailedStatus_SdkV2{}),
		"provisioning_status":      reflect.TypeOf(ProvisioningStatus_SdkV2{}),
		"triggered_update_status":  reflect.TypeOf(TriggeredUpdateStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineTableStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o OnlineTableStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o OnlineTableStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"continuous_update_status": basetypes.ListType{
				ElemType: ContinuousUpdateStatus_SdkV2{}.Type(ctx),
			},
			"detailed_state": types.StringType,
			"failed_status": basetypes.ListType{
				ElemType: FailedStatus_SdkV2{}.Type(ctx),
			},
			"message": types.StringType,
			"provisioning_status": basetypes.ListType{
				ElemType: ProvisioningStatus_SdkV2{}.Type(ctx),
			},
			"triggered_update_status": basetypes.ListType{
				ElemType: TriggeredUpdateStatus_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetContinuousUpdateStatus returns the value of the ContinuousUpdateStatus field in OnlineTableStatus_SdkV2 as
// a ContinuousUpdateStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableStatus_SdkV2) GetContinuousUpdateStatus(ctx context.Context) (ContinuousUpdateStatus_SdkV2, bool) {
	var e ContinuousUpdateStatus_SdkV2
	if o.ContinuousUpdateStatus.IsNull() || o.ContinuousUpdateStatus.IsUnknown() {
		return e, false
	}
	var v []ContinuousUpdateStatus_SdkV2
	d := o.ContinuousUpdateStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetContinuousUpdateStatus sets the value of the ContinuousUpdateStatus field in OnlineTableStatus_SdkV2.
func (o *OnlineTableStatus_SdkV2) SetContinuousUpdateStatus(ctx context.Context, v ContinuousUpdateStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["continuous_update_status"]
	o.ContinuousUpdateStatus = types.ListValueMust(t, vs)
}

// GetFailedStatus returns the value of the FailedStatus field in OnlineTableStatus_SdkV2 as
// a FailedStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableStatus_SdkV2) GetFailedStatus(ctx context.Context) (FailedStatus_SdkV2, bool) {
	var e FailedStatus_SdkV2
	if o.FailedStatus.IsNull() || o.FailedStatus.IsUnknown() {
		return e, false
	}
	var v []FailedStatus_SdkV2
	d := o.FailedStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFailedStatus sets the value of the FailedStatus field in OnlineTableStatus_SdkV2.
func (o *OnlineTableStatus_SdkV2) SetFailedStatus(ctx context.Context, v FailedStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["failed_status"]
	o.FailedStatus = types.ListValueMust(t, vs)
}

// GetProvisioningStatus returns the value of the ProvisioningStatus field in OnlineTableStatus_SdkV2 as
// a ProvisioningStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableStatus_SdkV2) GetProvisioningStatus(ctx context.Context) (ProvisioningStatus_SdkV2, bool) {
	var e ProvisioningStatus_SdkV2
	if o.ProvisioningStatus.IsNull() || o.ProvisioningStatus.IsUnknown() {
		return e, false
	}
	var v []ProvisioningStatus_SdkV2
	d := o.ProvisioningStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvisioningStatus sets the value of the ProvisioningStatus field in OnlineTableStatus_SdkV2.
func (o *OnlineTableStatus_SdkV2) SetProvisioningStatus(ctx context.Context, v ProvisioningStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provisioning_status"]
	o.ProvisioningStatus = types.ListValueMust(t, vs)
}

// GetTriggeredUpdateStatus returns the value of the TriggeredUpdateStatus field in OnlineTableStatus_SdkV2 as
// a TriggeredUpdateStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *OnlineTableStatus_SdkV2) GetTriggeredUpdateStatus(ctx context.Context) (TriggeredUpdateStatus_SdkV2, bool) {
	var e TriggeredUpdateStatus_SdkV2
	if o.TriggeredUpdateStatus.IsNull() || o.TriggeredUpdateStatus.IsUnknown() {
		return e, false
	}
	var v []TriggeredUpdateStatus_SdkV2
	d := o.TriggeredUpdateStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTriggeredUpdateStatus sets the value of the TriggeredUpdateStatus field in OnlineTableStatus_SdkV2.
func (o *OnlineTableStatus_SdkV2) SetTriggeredUpdateStatus(ctx context.Context, v TriggeredUpdateStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["triggered_update_status"]
	o.TriggeredUpdateStatus = types.ListValueMust(t, vs)
}

type PermissionsChange_SdkV2 struct {
	// The set of privileges to add.
	Add types.List `tfsdk:"add" tf:"optional"`
	// The principal whose privileges we are changing.
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The set of privileges to remove.
	Remove types.List `tfsdk:"remove" tf:"optional"`
}

func (newState *PermissionsChange_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsChange_SdkV2) {
}

func (newState *PermissionsChange_SdkV2) SyncEffectiveFieldsDuringRead(existingState PermissionsChange_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionsChange.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionsChange_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"add":    reflect.TypeOf(types.String{}),
		"remove": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionsChange_SdkV2
// only implements ToObjectValue() and Type().
func (o PermissionsChange_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"add":       o.Add,
			"principal": o.Principal,
			"remove":    o.Remove,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionsChange_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *PermissionsChange_SdkV2) GetAdd(ctx context.Context) ([]types.String, bool) {
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

// SetAdd sets the value of the Add field in PermissionsChange_SdkV2.
func (o *PermissionsChange_SdkV2) SetAdd(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["add"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Add = types.ListValueMust(t, vs)
}

// GetRemove returns the value of the Remove field in PermissionsChange_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionsChange_SdkV2) GetRemove(ctx context.Context) ([]types.String, bool) {
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

// SetRemove sets the value of the Remove field in PermissionsChange_SdkV2.
func (o *PermissionsChange_SdkV2) SetRemove(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["remove"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Remove = types.ListValueMust(t, vs)
}

type PermissionsList_SdkV2 struct {
	// The privileges assigned to each principal
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments" tf:"optional"`
}

func (newState *PermissionsList_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsList_SdkV2) {
}

func (newState *PermissionsList_SdkV2) SyncEffectiveFieldsDuringRead(existingState PermissionsList_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionsList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionsList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"privilege_assignments": reflect.TypeOf(PrivilegeAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PermissionsList_SdkV2
// only implements ToObjectValue() and Type().
func (o PermissionsList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"privilege_assignments": o.PrivilegeAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PermissionsList_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPrivilegeAssignments returns the value of the PrivilegeAssignments field in PermissionsList_SdkV2 as
// a slice of PrivilegeAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PermissionsList_SdkV2) GetPrivilegeAssignments(ctx context.Context) ([]PrivilegeAssignment_SdkV2, bool) {
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

// SetPrivilegeAssignments sets the value of the PrivilegeAssignments field in PermissionsList_SdkV2.
func (o *PermissionsList_SdkV2) SetPrivilegeAssignments(ctx context.Context, v []PrivilegeAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["privilege_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrivilegeAssignments = types.ListValueMust(t, vs)
}

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress_SdkV2 struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds types.Float64 `tfsdk:"estimated_completion_time_seconds" tf:"optional"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing types.Int64 `tfsdk:"latest_version_currently_processing" tf:"optional"`
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion types.Float64 `tfsdk:"sync_progress_completion" tf:"optional"`
	// The number of rows that have been synced in this update.
	SyncedRowCount types.Int64 `tfsdk:"synced_row_count" tf:"optional"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount types.Int64 `tfsdk:"total_row_count" tf:"optional"`
}

func (newState *PipelineProgress_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineProgress_SdkV2) {
}

func (newState *PipelineProgress_SdkV2) SyncEffectiveFieldsDuringRead(existingState PipelineProgress_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PipelineProgress.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PipelineProgress_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PipelineProgress_SdkV2
// only implements ToObjectValue() and Type().
func (o PipelineProgress_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PipelineProgress_SdkV2) Type(ctx context.Context) attr.Type {
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

type PrimaryKeyConstraint_SdkV2 struct {
	// Column names for this constraint.
	ChildColumns types.List `tfsdk:"child_columns" tf:""`
	// The name of the constraint.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *PrimaryKeyConstraint_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrimaryKeyConstraint_SdkV2) {
}

func (newState *PrimaryKeyConstraint_SdkV2) SyncEffectiveFieldsDuringRead(existingState PrimaryKeyConstraint_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrimaryKeyConstraint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PrimaryKeyConstraint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"child_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrimaryKeyConstraint_SdkV2
// only implements ToObjectValue() and Type().
func (o PrimaryKeyConstraint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"child_columns": o.ChildColumns,
			"name":          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PrimaryKeyConstraint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"child_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"name": types.StringType,
		},
	}
}

// GetChildColumns returns the value of the ChildColumns field in PrimaryKeyConstraint_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PrimaryKeyConstraint_SdkV2) GetChildColumns(ctx context.Context) ([]types.String, bool) {
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

// SetChildColumns sets the value of the ChildColumns field in PrimaryKeyConstraint_SdkV2.
func (o *PrimaryKeyConstraint_SdkV2) SetChildColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["child_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ChildColumns = types.ListValueMust(t, vs)
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

// Status of an asynchronously provisioned resource.
type ProvisioningInfo_SdkV2 struct {
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *ProvisioningInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProvisioningInfo_SdkV2) {
}

func (newState *ProvisioningInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ProvisioningInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProvisioningInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProvisioningInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProvisioningInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ProvisioningInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"state": o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ProvisioningInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"state": types.StringType,
		},
	}
}

// Detailed status of an online table. Shown if the online table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type ProvisioningStatus_SdkV2 struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress types.List `tfsdk:"initial_pipeline_sync_progress" tf:"optional,object"`
}

func (newState *ProvisioningStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProvisioningStatus_SdkV2) {
}

func (newState *ProvisioningStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState ProvisioningStatus_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProvisioningStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ProvisioningStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(PipelineProgress_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProvisioningStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o ProvisioningStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": o.InitialPipelineSyncProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ProvisioningStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": basetypes.ListType{
				ElemType: PipelineProgress_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetInitialPipelineSyncProgress returns the value of the InitialPipelineSyncProgress field in ProvisioningStatus_SdkV2 as
// a PipelineProgress_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ProvisioningStatus_SdkV2) GetInitialPipelineSyncProgress(ctx context.Context) (PipelineProgress_SdkV2, bool) {
	var e PipelineProgress_SdkV2
	if o.InitialPipelineSyncProgress.IsNull() || o.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v []PipelineProgress_SdkV2
	d := o.InitialPipelineSyncProgress.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in ProvisioningStatus_SdkV2.
func (o *ProvisioningStatus_SdkV2) SetInitialPipelineSyncProgress(ctx context.Context, v PipelineProgress_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["initial_pipeline_sync_progress"]
	o.InitialPipelineSyncProgress = types.ListValueMust(t, vs)
}

type QuotaInfo_SdkV2 struct {
	// The timestamp that indicates when the quota count was last updated.
	LastRefreshedAt types.Int64 `tfsdk:"last_refreshed_at" tf:"optional"`
	// Name of the parent resource. Returns metastore ID if the parent is a
	// metastore.
	ParentFullName types.String `tfsdk:"parent_full_name" tf:"optional"`
	// The quota parent securable type.
	ParentSecurableType types.String `tfsdk:"parent_securable_type" tf:"optional"`
	// The current usage of the resource quota.
	QuotaCount types.Int64 `tfsdk:"quota_count" tf:"optional"`
	// The current limit of the resource quota.
	QuotaLimit types.Int64 `tfsdk:"quota_limit" tf:"optional"`
	// The name of the quota.
	QuotaName types.String `tfsdk:"quota_name" tf:"optional"`
}

func (newState *QuotaInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QuotaInfo_SdkV2) {
}

func (newState *QuotaInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState QuotaInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QuotaInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QuotaInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QuotaInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o QuotaInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o QuotaInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
type R2Credentials_SdkV2 struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId types.String `tfsdk:"access_key_id" tf:"optional"`
	// The secret access key associated with the access key.
	SecretAccessKey types.String `tfsdk:"secret_access_key" tf:"optional"`
	// The generated JWT that users must pass to use the temporary credentials.
	SessionToken types.String `tfsdk:"session_token" tf:"optional"`
}

func (newState *R2Credentials_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan R2Credentials_SdkV2) {
}

func (newState *R2Credentials_SdkV2) SyncEffectiveFieldsDuringRead(existingState R2Credentials_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in R2Credentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a R2Credentials_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, R2Credentials_SdkV2
// only implements ToObjectValue() and Type().
func (o R2Credentials_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_key_id":     o.AccessKeyId,
			"secret_access_key": o.SecretAccessKey,
			"session_token":     o.SessionToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o R2Credentials_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_key_id":     types.StringType,
			"secret_access_key": types.StringType,
			"session_token":     types.StringType,
		},
	}
}

// Get a Volume
type ReadVolumeRequest_SdkV2 struct {
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse types.Bool `tfsdk:"-"`
	// The three-level (fully qualified) name of the volume
	Name types.String `tfsdk:"-"`
}

func (newState *ReadVolumeRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReadVolumeRequest_SdkV2) {
}

func (newState *ReadVolumeRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ReadVolumeRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReadVolumeRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReadVolumeRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReadVolumeRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ReadVolumeRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_browse": o.IncludeBrowse,
			"name":           o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ReadVolumeRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"name":           types.StringType,
		},
	}
}

type RegenerateDashboardRequest_SdkV2 struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Optional argument to specify the warehouse for dashboard regeneration. If
	// not specified, the first running warehouse will be used.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *RegenerateDashboardRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegenerateDashboardRequest_SdkV2) {
}

func (newState *RegenerateDashboardRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState RegenerateDashboardRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegenerateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RegenerateDashboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegenerateDashboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RegenerateDashboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name":   o.TableName,
			"warehouse_id": o.WarehouseId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegenerateDashboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name":   types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

type RegenerateDashboardResponse_SdkV2 struct {
	// Id of the regenerated monitoring dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The directory where the regenerated dashboard is stored.
	ParentFolder types.String `tfsdk:"parent_folder" tf:"optional"`
}

func (newState *RegenerateDashboardResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegenerateDashboardResponse_SdkV2) {
}

func (newState *RegenerateDashboardResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState RegenerateDashboardResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegenerateDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RegenerateDashboardResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegenerateDashboardResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RegenerateDashboardResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dashboard_id":  o.DashboardId,
			"parent_folder": o.ParentFolder,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegenerateDashboardResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":  types.StringType,
			"parent_folder": types.StringType,
		},
	}
}

// Registered model alias.
type RegisteredModelAlias_SdkV2 struct {
	// Name of the alias, e.g. 'champion' or 'latest_stable'
	AliasName types.String `tfsdk:"alias_name" tf:"optional"`
	// Integer version number of the model version to which this alias points.
	VersionNum types.Int64 `tfsdk:"version_num" tf:"optional"`
}

func (newState *RegisteredModelAlias_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelAlias_SdkV2) {
}

func (newState *RegisteredModelAlias_SdkV2) SyncEffectiveFieldsDuringRead(existingState RegisteredModelAlias_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelAlias.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RegisteredModelAlias_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAlias_SdkV2
// only implements ToObjectValue() and Type().
func (o RegisteredModelAlias_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias_name":  o.AliasName,
			"version_num": o.VersionNum,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelAlias_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias_name":  types.StringType,
			"version_num": types.Int64Type,
		},
	}
}

type RegisteredModelInfo_SdkV2 struct {
	// List of aliases associated with the registered model
	Aliases types.List `tfsdk:"aliases" tf:"optional"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// The name of the catalog where the schema and the registered model reside
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The comment attached to the registered model
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The identifier of the user who created the registered model
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// The unique identifier of the metastore
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The name of the registered model
	Name types.String `tfsdk:"name" tf:"optional"`
	// The identifier of the user who owns the registered model
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// The name of the schema where the registered model resides
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// The identifier of the user who updated the registered model last time
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *RegisteredModelInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelInfo_SdkV2) {
}

func (newState *RegisteredModelInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState RegisteredModelInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RegisteredModelInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aliases": reflect.TypeOf(RegisteredModelAlias_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o RegisteredModelInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RegisteredModelInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aliases": basetypes.ListType{
				ElemType: RegisteredModelAlias_SdkV2{}.Type(ctx),
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

// GetAliases returns the value of the Aliases field in RegisteredModelInfo_SdkV2 as
// a slice of RegisteredModelAlias_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RegisteredModelInfo_SdkV2) GetAliases(ctx context.Context) ([]RegisteredModelAlias_SdkV2, bool) {
	if o.Aliases.IsNull() || o.Aliases.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAlias_SdkV2
	d := o.Aliases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAliases sets the value of the Aliases field in RegisteredModelInfo_SdkV2.
func (o *RegisteredModelInfo_SdkV2) SetAliases(ctx context.Context, v []RegisteredModelAlias_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Aliases = types.ListValueMust(t, vs)
}

// Queue a metric refresh for a monitor
type RunRefreshRequest_SdkV2 struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *RunRefreshRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunRefreshRequest_SdkV2) {
}

func (newState *RunRefreshRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState RunRefreshRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunRefreshRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RunRefreshRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunRefreshRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RunRefreshRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name": o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunRefreshRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name": types.StringType,
		},
	}
}

type SchemaInfo_SdkV2 struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The type of the parent catalog.
	CatalogType types.String `tfsdk:"catalog_type" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this schema was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of schema creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`

	EffectivePredictiveOptimizationFlag types.List `tfsdk:"effective_predictive_optimization_flag" tf:"optional,object"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization" tf:"optional"`
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of schema, relative to parent catalog.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of current owner of schema.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties" tf:"optional"`
	// The unique identifier of the schema.
	SchemaId types.String `tfsdk:"schema_id" tf:"optional"`
	// Storage location for managed tables within schema.
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Storage root URL for managed tables within schema.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Time at which this schema was created, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified schema.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

func (newState *SchemaInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SchemaInfo_SdkV2) {
}

func (newState *SchemaInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState SchemaInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SchemaInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SchemaInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_predictive_optimization_flag": reflect.TypeOf(EffectivePredictiveOptimizationFlag_SdkV2{}),
		"properties":                             reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SchemaInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o SchemaInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SchemaInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":  types.BoolType,
			"catalog_name": types.StringType,
			"catalog_type": types.StringType,
			"comment":      types.StringType,
			"created_at":   types.Int64Type,
			"created_by":   types.StringType,
			"effective_predictive_optimization_flag": basetypes.ListType{
				ElemType: EffectivePredictiveOptimizationFlag_SdkV2{}.Type(ctx),
			},
			"enable_predictive_optimization": types.StringType,
			"full_name":                      types.StringType,
			"metastore_id":                   types.StringType,
			"name":                           types.StringType,
			"owner":                          types.StringType,
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

// GetEffectivePredictiveOptimizationFlag returns the value of the EffectivePredictiveOptimizationFlag field in SchemaInfo_SdkV2 as
// a EffectivePredictiveOptimizationFlag_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SchemaInfo_SdkV2) GetEffectivePredictiveOptimizationFlag(ctx context.Context) (EffectivePredictiveOptimizationFlag_SdkV2, bool) {
	var e EffectivePredictiveOptimizationFlag_SdkV2
	if o.EffectivePredictiveOptimizationFlag.IsNull() || o.EffectivePredictiveOptimizationFlag.IsUnknown() {
		return e, false
	}
	var v []EffectivePredictiveOptimizationFlag_SdkV2
	d := o.EffectivePredictiveOptimizationFlag.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectivePredictiveOptimizationFlag sets the value of the EffectivePredictiveOptimizationFlag field in SchemaInfo_SdkV2.
func (o *SchemaInfo_SdkV2) SetEffectivePredictiveOptimizationFlag(ctx context.Context, v EffectivePredictiveOptimizationFlag_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_predictive_optimization_flag"]
	o.EffectivePredictiveOptimizationFlag = types.ListValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in SchemaInfo_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SchemaInfo_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in SchemaInfo_SdkV2.
func (o *SchemaInfo_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type SetArtifactAllowlist_SdkV2 struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers types.List `tfsdk:"artifact_matchers" tf:""`
	// The artifact type of the allowlist.
	ArtifactType types.String `tfsdk:"-"`
}

func (newState *SetArtifactAllowlist_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetArtifactAllowlist_SdkV2) {
}

func (newState *SetArtifactAllowlist_SdkV2) SyncEffectiveFieldsDuringRead(existingState SetArtifactAllowlist_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetArtifactAllowlist.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetArtifactAllowlist_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"artifact_matchers": reflect.TypeOf(ArtifactMatcher_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetArtifactAllowlist_SdkV2
// only implements ToObjectValue() and Type().
func (o SetArtifactAllowlist_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_matchers": o.ArtifactMatchers,
			"artifact_type":     o.ArtifactType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetArtifactAllowlist_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_matchers": basetypes.ListType{
				ElemType: ArtifactMatcher_SdkV2{}.Type(ctx),
			},
			"artifact_type": types.StringType,
		},
	}
}

// GetArtifactMatchers returns the value of the ArtifactMatchers field in SetArtifactAllowlist_SdkV2 as
// a slice of ArtifactMatcher_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SetArtifactAllowlist_SdkV2) GetArtifactMatchers(ctx context.Context) ([]ArtifactMatcher_SdkV2, bool) {
	if o.ArtifactMatchers.IsNull() || o.ArtifactMatchers.IsUnknown() {
		return nil, false
	}
	var v []ArtifactMatcher_SdkV2
	d := o.ArtifactMatchers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetArtifactMatchers sets the value of the ArtifactMatchers field in SetArtifactAllowlist_SdkV2.
func (o *SetArtifactAllowlist_SdkV2) SetArtifactMatchers(ctx context.Context, v []ArtifactMatcher_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["artifact_matchers"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ArtifactMatchers = types.ListValueMust(t, vs)
}

type SetRegisteredModelAliasRequest_SdkV2 struct {
	// The name of the alias
	Alias types.String `tfsdk:"alias" tf:""`
	// Full name of the registered model
	FullName types.String `tfsdk:"full_name" tf:""`
	// The version number of the model version to which the alias points
	VersionNum types.Int64 `tfsdk:"version_num" tf:""`
}

func (newState *SetRegisteredModelAliasRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetRegisteredModelAliasRequest_SdkV2) {
}

func (newState *SetRegisteredModelAliasRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState SetRegisteredModelAliasRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetRegisteredModelAliasRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetRegisteredModelAliasRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetRegisteredModelAliasRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SetRegisteredModelAliasRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"alias":       o.Alias,
			"full_name":   o.FullName,
			"version_num": o.VersionNum,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetRegisteredModelAliasRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":       types.StringType,
			"full_name":   types.StringType,
			"version_num": types.Int64Type,
		},
	}
}

// Server-Side Encryption properties for clients communicating with AWS s3.
type SseEncryptionDetails_SdkV2 struct {
	// The type of key encryption to use (affects headers from s3 client).
	Algorithm types.String `tfsdk:"algorithm" tf:"optional"`
	// When algorithm is **AWS_SSE_KMS** this field specifies the ARN of the SSE
	// key to use.
	AwsKmsKeyArn types.String `tfsdk:"aws_kms_key_arn" tf:"optional"`
}

func (newState *SseEncryptionDetails_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SseEncryptionDetails_SdkV2) {
}

func (newState *SseEncryptionDetails_SdkV2) SyncEffectiveFieldsDuringRead(existingState SseEncryptionDetails_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SseEncryptionDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SseEncryptionDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SseEncryptionDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o SseEncryptionDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"algorithm":       o.Algorithm,
			"aws_kms_key_arn": o.AwsKmsKeyArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SseEncryptionDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"algorithm":       types.StringType,
			"aws_kms_key_arn": types.StringType,
		},
	}
}

type StorageCredentialInfo_SdkV2 struct {
	// The AWS IAM role configuration.
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// The Cloudflare API token configuration.
	CloudflareApiToken types.List `tfsdk:"cloudflare_api_token" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this Credential was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of credential creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount types.List `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// The full name of the credential.
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// The unique identifier of the credential.
	Id types.String `tfsdk:"id" tf:"optional"`

	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The credential name. The name must be unique within the metastore.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of current owner of credential.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the credential.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// Whether this credential is the current metastore's root storage
	// credential.
	UsedForManagedStorage types.Bool `tfsdk:"used_for_managed_storage" tf:"optional"`
}

func (newState *StorageCredentialInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StorageCredentialInfo_SdkV2) {
}

func (newState *StorageCredentialInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState StorageCredentialInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StorageCredentialInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StorageCredentialInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRoleResponse_SdkV2{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentityResponse_SdkV2{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal_SdkV2{}),
		"cloudflare_api_token":           reflect.TypeOf(CloudflareApiToken_SdkV2{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccountResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StorageCredentialInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o StorageCredentialInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o StorageCredentialInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRoleResponse_SdkV2{}.Type(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentityResponse_SdkV2{}.Type(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal_SdkV2{}.Type(ctx),
			},
			"cloudflare_api_token": basetypes.ListType{
				ElemType: CloudflareApiToken_SdkV2{}.Type(ctx),
			},
			"comment":    types.StringType,
			"created_at": types.Int64Type,
			"created_by": types.StringType,
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccountResponse_SdkV2{}.Type(ctx),
			},
			"full_name":                types.StringType,
			"id":                       types.StringType,
			"isolation_mode":           types.StringType,
			"metastore_id":             types.StringType,
			"name":                     types.StringType,
			"owner":                    types.StringType,
			"read_only":                types.BoolType,
			"updated_at":               types.Int64Type,
			"updated_by":               types.StringType,
			"used_for_managed_storage": types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in StorageCredentialInfo_SdkV2 as
// a AwsIamRoleResponse_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageCredentialInfo_SdkV2) GetAwsIamRole(ctx context.Context) (AwsIamRoleResponse_SdkV2, bool) {
	var e AwsIamRoleResponse_SdkV2
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRoleResponse_SdkV2
	d := o.AwsIamRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsIamRole sets the value of the AwsIamRole field in StorageCredentialInfo_SdkV2.
func (o *StorageCredentialInfo_SdkV2) SetAwsIamRole(ctx context.Context, v AwsIamRoleResponse_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_iam_role"]
	o.AwsIamRole = types.ListValueMust(t, vs)
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in StorageCredentialInfo_SdkV2 as
// a AzureManagedIdentityResponse_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageCredentialInfo_SdkV2) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentityResponse_SdkV2, bool) {
	var e AzureManagedIdentityResponse_SdkV2
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentityResponse_SdkV2
	d := o.AzureManagedIdentity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in StorageCredentialInfo_SdkV2.
func (o *StorageCredentialInfo_SdkV2) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentityResponse_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_managed_identity"]
	o.AzureManagedIdentity = types.ListValueMust(t, vs)
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in StorageCredentialInfo_SdkV2 as
// a AzureServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageCredentialInfo_SdkV2) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal_SdkV2, bool) {
	var e AzureServicePrincipal_SdkV2
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal_SdkV2
	d := o.AzureServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in StorageCredentialInfo_SdkV2.
func (o *StorageCredentialInfo_SdkV2) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_service_principal"]
	o.AzureServicePrincipal = types.ListValueMust(t, vs)
}

// GetCloudflareApiToken returns the value of the CloudflareApiToken field in StorageCredentialInfo_SdkV2 as
// a CloudflareApiToken_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageCredentialInfo_SdkV2) GetCloudflareApiToken(ctx context.Context) (CloudflareApiToken_SdkV2, bool) {
	var e CloudflareApiToken_SdkV2
	if o.CloudflareApiToken.IsNull() || o.CloudflareApiToken.IsUnknown() {
		return e, false
	}
	var v []CloudflareApiToken_SdkV2
	d := o.CloudflareApiToken.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudflareApiToken sets the value of the CloudflareApiToken field in StorageCredentialInfo_SdkV2.
func (o *StorageCredentialInfo_SdkV2) SetCloudflareApiToken(ctx context.Context, v CloudflareApiToken_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cloudflare_api_token"]
	o.CloudflareApiToken = types.ListValueMust(t, vs)
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in StorageCredentialInfo_SdkV2 as
// a DatabricksGcpServiceAccountResponse_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageCredentialInfo_SdkV2) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccountResponse_SdkV2, bool) {
	var e DatabricksGcpServiceAccountResponse_SdkV2
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccountResponse_SdkV2
	d := o.DatabricksGcpServiceAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in StorageCredentialInfo_SdkV2.
func (o *StorageCredentialInfo_SdkV2) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccountResponse_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["databricks_gcp_service_account"]
	o.DatabricksGcpServiceAccount = types.ListValueMust(t, vs)
}

type SystemSchemaInfo_SdkV2 struct {
	// Name of the system schema.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *SystemSchemaInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SystemSchemaInfo_SdkV2) {
}

func (newState *SystemSchemaInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState SystemSchemaInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SystemSchemaInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SystemSchemaInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SystemSchemaInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o SystemSchemaInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"schema": o.Schema,
			"state":  o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SystemSchemaInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
type TableConstraint_SdkV2 struct {
	ForeignKeyConstraint types.List `tfsdk:"foreign_key_constraint" tf:"optional,object"`

	NamedTableConstraint types.List `tfsdk:"named_table_constraint" tf:"optional,object"`

	PrimaryKeyConstraint types.List `tfsdk:"primary_key_constraint" tf:"optional,object"`
}

func (newState *TableConstraint_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableConstraint_SdkV2) {
}

func (newState *TableConstraint_SdkV2) SyncEffectiveFieldsDuringRead(existingState TableConstraint_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableConstraint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableConstraint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"foreign_key_constraint": reflect.TypeOf(ForeignKeyConstraint_SdkV2{}),
		"named_table_constraint": reflect.TypeOf(NamedTableConstraint_SdkV2{}),
		"primary_key_constraint": reflect.TypeOf(PrimaryKeyConstraint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableConstraint_SdkV2
// only implements ToObjectValue() and Type().
func (o TableConstraint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"foreign_key_constraint": o.ForeignKeyConstraint,
			"named_table_constraint": o.NamedTableConstraint,
			"primary_key_constraint": o.PrimaryKeyConstraint,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableConstraint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"foreign_key_constraint": basetypes.ListType{
				ElemType: ForeignKeyConstraint_SdkV2{}.Type(ctx),
			},
			"named_table_constraint": basetypes.ListType{
				ElemType: NamedTableConstraint_SdkV2{}.Type(ctx),
			},
			"primary_key_constraint": basetypes.ListType{
				ElemType: PrimaryKeyConstraint_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetForeignKeyConstraint returns the value of the ForeignKeyConstraint field in TableConstraint_SdkV2 as
// a ForeignKeyConstraint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableConstraint_SdkV2) GetForeignKeyConstraint(ctx context.Context) (ForeignKeyConstraint_SdkV2, bool) {
	var e ForeignKeyConstraint_SdkV2
	if o.ForeignKeyConstraint.IsNull() || o.ForeignKeyConstraint.IsUnknown() {
		return e, false
	}
	var v []ForeignKeyConstraint_SdkV2
	d := o.ForeignKeyConstraint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetForeignKeyConstraint sets the value of the ForeignKeyConstraint field in TableConstraint_SdkV2.
func (o *TableConstraint_SdkV2) SetForeignKeyConstraint(ctx context.Context, v ForeignKeyConstraint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["foreign_key_constraint"]
	o.ForeignKeyConstraint = types.ListValueMust(t, vs)
}

// GetNamedTableConstraint returns the value of the NamedTableConstraint field in TableConstraint_SdkV2 as
// a NamedTableConstraint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableConstraint_SdkV2) GetNamedTableConstraint(ctx context.Context) (NamedTableConstraint_SdkV2, bool) {
	var e NamedTableConstraint_SdkV2
	if o.NamedTableConstraint.IsNull() || o.NamedTableConstraint.IsUnknown() {
		return e, false
	}
	var v []NamedTableConstraint_SdkV2
	d := o.NamedTableConstraint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNamedTableConstraint sets the value of the NamedTableConstraint field in TableConstraint_SdkV2.
func (o *TableConstraint_SdkV2) SetNamedTableConstraint(ctx context.Context, v NamedTableConstraint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["named_table_constraint"]
	o.NamedTableConstraint = types.ListValueMust(t, vs)
}

// GetPrimaryKeyConstraint returns the value of the PrimaryKeyConstraint field in TableConstraint_SdkV2 as
// a PrimaryKeyConstraint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableConstraint_SdkV2) GetPrimaryKeyConstraint(ctx context.Context) (PrimaryKeyConstraint_SdkV2, bool) {
	var e PrimaryKeyConstraint_SdkV2
	if o.PrimaryKeyConstraint.IsNull() || o.PrimaryKeyConstraint.IsUnknown() {
		return e, false
	}
	var v []PrimaryKeyConstraint_SdkV2
	d := o.PrimaryKeyConstraint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPrimaryKeyConstraint sets the value of the PrimaryKeyConstraint field in TableConstraint_SdkV2.
func (o *TableConstraint_SdkV2) SetPrimaryKeyConstraint(ctx context.Context, v PrimaryKeyConstraint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_key_constraint"]
	o.PrimaryKeyConstraint = types.ListValueMust(t, vs)
}

// A table that is dependent on a SQL object.
type TableDependency_SdkV2 struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	TableFullName types.String `tfsdk:"table_full_name" tf:""`
}

func (newState *TableDependency_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableDependency_SdkV2) {
}

func (newState *TableDependency_SdkV2) SyncEffectiveFieldsDuringRead(existingState TableDependency_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableDependency.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableDependency_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableDependency_SdkV2
// only implements ToObjectValue() and Type().
func (o TableDependency_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_full_name": o.TableFullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableDependency_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_full_name": types.StringType,
		},
	}
}

type TableExistsResponse_SdkV2 struct {
	// Whether the table exists or not.
	TableExists types.Bool `tfsdk:"table_exists" tf:"optional"`
}

func (newState *TableExistsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableExistsResponse_SdkV2) {
}

func (newState *TableExistsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState TableExistsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableExistsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableExistsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableExistsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o TableExistsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_exists": o.TableExists,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableExistsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_exists": types.BoolType,
		},
	}
}

type TableInfo_SdkV2 struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// Name of parent catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns types.List `tfsdk:"columns" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this table was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of table creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique ID of the Data Access Configuration to use with the table data.
	DataAccessConfigurationId types.String `tfsdk:"data_access_configuration_id" tf:"optional"`
	// Data source format
	DataSourceFormat types.String `tfsdk:"data_source_format" tf:"optional"`
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	DeletedAt types.Int64 `tfsdk:"deleted_at" tf:"optional"`
	// Information pertaining to current state of the delta table.
	DeltaRuntimePropertiesKvpairs types.List `tfsdk:"delta_runtime_properties_kvpairs" tf:"optional,object"`

	EffectivePredictiveOptimizationFlag types.List `tfsdk:"effective_predictive_optimization_flag" tf:"optional,object"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization" tf:"optional"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails types.List `tfsdk:"encryption_details" tf:"optional,object"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// Name of table, relative to parent schema.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Username of current owner of table.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	PipelineId types.String `tfsdk:"pipeline_id" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties" tf:"optional"`

	RowFilter types.List `tfsdk:"row_filter" tf:"optional,object"`
	// Name of parent schema relative to its parent catalog.
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath types.String `tfsdk:"sql_path" tf:"optional"`
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	StorageCredentialName types.String `tfsdk:"storage_credential_name" tf:"optional"`
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables)
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	TableConstraints types.List `tfsdk:"table_constraints" tf:"optional"`
	// The unique identifier of the table.
	TableId types.String `tfsdk:"table_id" tf:"optional"`

	TableType types.String `tfsdk:"table_type" tf:"optional"`
	// Time at which this table was last modified, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// Username of user who last modified the table.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	ViewDefinition types.String `tfsdk:"view_definition" tf:"optional"`
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	ViewDependencies types.List `tfsdk:"view_dependencies" tf:"optional,object"`
}

func (newState *TableInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableInfo_SdkV2) {
}

func (newState *TableInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState TableInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns":                                reflect.TypeOf(ColumnInfo_SdkV2{}),
		"delta_runtime_properties_kvpairs":       reflect.TypeOf(DeltaRuntimePropertiesKvPairs_SdkV2{}),
		"effective_predictive_optimization_flag": reflect.TypeOf(EffectivePredictiveOptimizationFlag_SdkV2{}),
		"encryption_details":                     reflect.TypeOf(EncryptionDetails_SdkV2{}),
		"properties":                             reflect.TypeOf(types.String{}),
		"row_filter":                             reflect.TypeOf(TableRowFilter_SdkV2{}),
		"table_constraints":                      reflect.TypeOf(TableConstraint_SdkV2{}),
		"view_dependencies":                      reflect.TypeOf(DependencyList_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o TableInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TableInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point": types.StringType,
			"browse_only":  types.BoolType,
			"catalog_name": types.StringType,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo_SdkV2{}.Type(ctx),
			},
			"comment":                      types.StringType,
			"created_at":                   types.Int64Type,
			"created_by":                   types.StringType,
			"data_access_configuration_id": types.StringType,
			"data_source_format":           types.StringType,
			"deleted_at":                   types.Int64Type,
			"delta_runtime_properties_kvpairs": basetypes.ListType{
				ElemType: DeltaRuntimePropertiesKvPairs_SdkV2{}.Type(ctx),
			},
			"effective_predictive_optimization_flag": basetypes.ListType{
				ElemType: EffectivePredictiveOptimizationFlag_SdkV2{}.Type(ctx),
			},
			"enable_predictive_optimization": types.StringType,
			"encryption_details": basetypes.ListType{
				ElemType: EncryptionDetails_SdkV2{}.Type(ctx),
			},
			"full_name":    types.StringType,
			"metastore_id": types.StringType,
			"name":         types.StringType,
			"owner":        types.StringType,
			"pipeline_id":  types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
			"row_filter": basetypes.ListType{
				ElemType: TableRowFilter_SdkV2{}.Type(ctx),
			},
			"schema_name":             types.StringType,
			"sql_path":                types.StringType,
			"storage_credential_name": types.StringType,
			"storage_location":        types.StringType,
			"table_constraints": basetypes.ListType{
				ElemType: TableConstraint_SdkV2{}.Type(ctx),
			},
			"table_id":        types.StringType,
			"table_type":      types.StringType,
			"updated_at":      types.Int64Type,
			"updated_by":      types.StringType,
			"view_definition": types.StringType,
			"view_dependencies": basetypes.ListType{
				ElemType: DependencyList_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in TableInfo_SdkV2 as
// a slice of ColumnInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo_SdkV2) GetColumns(ctx context.Context) ([]ColumnInfo_SdkV2, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo_SdkV2
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in TableInfo_SdkV2.
func (o *TableInfo_SdkV2) SetColumns(ctx context.Context, v []ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

// GetDeltaRuntimePropertiesKvpairs returns the value of the DeltaRuntimePropertiesKvpairs field in TableInfo_SdkV2 as
// a DeltaRuntimePropertiesKvPairs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo_SdkV2) GetDeltaRuntimePropertiesKvpairs(ctx context.Context) (DeltaRuntimePropertiesKvPairs_SdkV2, bool) {
	var e DeltaRuntimePropertiesKvPairs_SdkV2
	if o.DeltaRuntimePropertiesKvpairs.IsNull() || o.DeltaRuntimePropertiesKvpairs.IsUnknown() {
		return e, false
	}
	var v []DeltaRuntimePropertiesKvPairs_SdkV2
	d := o.DeltaRuntimePropertiesKvpairs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeltaRuntimePropertiesKvpairs sets the value of the DeltaRuntimePropertiesKvpairs field in TableInfo_SdkV2.
func (o *TableInfo_SdkV2) SetDeltaRuntimePropertiesKvpairs(ctx context.Context, v DeltaRuntimePropertiesKvPairs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["delta_runtime_properties_kvpairs"]
	o.DeltaRuntimePropertiesKvpairs = types.ListValueMust(t, vs)
}

// GetEffectivePredictiveOptimizationFlag returns the value of the EffectivePredictiveOptimizationFlag field in TableInfo_SdkV2 as
// a EffectivePredictiveOptimizationFlag_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo_SdkV2) GetEffectivePredictiveOptimizationFlag(ctx context.Context) (EffectivePredictiveOptimizationFlag_SdkV2, bool) {
	var e EffectivePredictiveOptimizationFlag_SdkV2
	if o.EffectivePredictiveOptimizationFlag.IsNull() || o.EffectivePredictiveOptimizationFlag.IsUnknown() {
		return e, false
	}
	var v []EffectivePredictiveOptimizationFlag_SdkV2
	d := o.EffectivePredictiveOptimizationFlag.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectivePredictiveOptimizationFlag sets the value of the EffectivePredictiveOptimizationFlag field in TableInfo_SdkV2.
func (o *TableInfo_SdkV2) SetEffectivePredictiveOptimizationFlag(ctx context.Context, v EffectivePredictiveOptimizationFlag_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_predictive_optimization_flag"]
	o.EffectivePredictiveOptimizationFlag = types.ListValueMust(t, vs)
}

// GetEncryptionDetails returns the value of the EncryptionDetails field in TableInfo_SdkV2 as
// a EncryptionDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo_SdkV2) GetEncryptionDetails(ctx context.Context) (EncryptionDetails_SdkV2, bool) {
	var e EncryptionDetails_SdkV2
	if o.EncryptionDetails.IsNull() || o.EncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []EncryptionDetails_SdkV2
	d := o.EncryptionDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEncryptionDetails sets the value of the EncryptionDetails field in TableInfo_SdkV2.
func (o *TableInfo_SdkV2) SetEncryptionDetails(ctx context.Context, v EncryptionDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["encryption_details"]
	o.EncryptionDetails = types.ListValueMust(t, vs)
}

// GetProperties returns the value of the Properties field in TableInfo_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in TableInfo_SdkV2.
func (o *TableInfo_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

// GetRowFilter returns the value of the RowFilter field in TableInfo_SdkV2 as
// a TableRowFilter_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo_SdkV2) GetRowFilter(ctx context.Context) (TableRowFilter_SdkV2, bool) {
	var e TableRowFilter_SdkV2
	if o.RowFilter.IsNull() || o.RowFilter.IsUnknown() {
		return e, false
	}
	var v []TableRowFilter_SdkV2
	d := o.RowFilter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRowFilter sets the value of the RowFilter field in TableInfo_SdkV2.
func (o *TableInfo_SdkV2) SetRowFilter(ctx context.Context, v TableRowFilter_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["row_filter"]
	o.RowFilter = types.ListValueMust(t, vs)
}

// GetTableConstraints returns the value of the TableConstraints field in TableInfo_SdkV2 as
// a slice of TableConstraint_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo_SdkV2) GetTableConstraints(ctx context.Context) ([]TableConstraint_SdkV2, bool) {
	if o.TableConstraints.IsNull() || o.TableConstraints.IsUnknown() {
		return nil, false
	}
	var v []TableConstraint_SdkV2
	d := o.TableConstraints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableConstraints sets the value of the TableConstraints field in TableInfo_SdkV2.
func (o *TableInfo_SdkV2) SetTableConstraints(ctx context.Context, v []TableConstraint_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table_constraints"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TableConstraints = types.ListValueMust(t, vs)
}

// GetViewDependencies returns the value of the ViewDependencies field in TableInfo_SdkV2 as
// a DependencyList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TableInfo_SdkV2) GetViewDependencies(ctx context.Context) (DependencyList_SdkV2, bool) {
	var e DependencyList_SdkV2
	if o.ViewDependencies.IsNull() || o.ViewDependencies.IsUnknown() {
		return e, false
	}
	var v []DependencyList_SdkV2
	d := o.ViewDependencies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetViewDependencies sets the value of the ViewDependencies field in TableInfo_SdkV2.
func (o *TableInfo_SdkV2) SetViewDependencies(ctx context.Context, v DependencyList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["view_dependencies"]
	o.ViewDependencies = types.ListValueMust(t, vs)
}

type TableRowFilter_SdkV2 struct {
	// The full name of the row filter SQL UDF.
	FunctionName types.String `tfsdk:"function_name" tf:""`
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	InputColumnNames types.List `tfsdk:"input_column_names" tf:""`
}

func (newState *TableRowFilter_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableRowFilter_SdkV2) {
}

func (newState *TableRowFilter_SdkV2) SyncEffectiveFieldsDuringRead(existingState TableRowFilter_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableRowFilter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableRowFilter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input_column_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableRowFilter_SdkV2
// only implements ToObjectValue() and Type().
func (o TableRowFilter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"function_name":      o.FunctionName,
			"input_column_names": o.InputColumnNames,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableRowFilter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_name": types.StringType,
			"input_column_names": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetInputColumnNames returns the value of the InputColumnNames field in TableRowFilter_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *TableRowFilter_SdkV2) GetInputColumnNames(ctx context.Context) ([]types.String, bool) {
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

// SetInputColumnNames sets the value of the InputColumnNames field in TableRowFilter_SdkV2.
func (o *TableRowFilter_SdkV2) SetInputColumnNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["input_column_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InputColumnNames = types.ListValueMust(t, vs)
}

type TableSummary_SdkV2 struct {
	// The full name of the table.
	FullName types.String `tfsdk:"full_name" tf:"optional"`

	TableType types.String `tfsdk:"table_type" tf:"optional"`
}

func (newState *TableSummary_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableSummary_SdkV2) {
}

func (newState *TableSummary_SdkV2) SyncEffectiveFieldsDuringRead(existingState TableSummary_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TableSummary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TableSummary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TableSummary_SdkV2
// only implements ToObjectValue() and Type().
func (o TableSummary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name":  o.FullName,
			"table_type": o.TableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TableSummary_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":  types.StringType,
			"table_type": types.StringType,
		},
	}
}

type TemporaryCredentials_SdkV2 struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	AwsTempCredentials types.List `tfsdk:"aws_temp_credentials" tf:"optional,object"`
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	AzureAad types.List `tfsdk:"azure_aad" tf:"optional,object"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime types.Int64 `tfsdk:"expiration_time" tf:"optional"`
}

func (newState *TemporaryCredentials_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TemporaryCredentials_SdkV2) {
}

func (newState *TemporaryCredentials_SdkV2) SyncEffectiveFieldsDuringRead(existingState TemporaryCredentials_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TemporaryCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TemporaryCredentials_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_temp_credentials": reflect.TypeOf(AwsCredentials_SdkV2{}),
		"azure_aad":            reflect.TypeOf(AzureActiveDirectoryToken_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TemporaryCredentials_SdkV2
// only implements ToObjectValue() and Type().
func (o TemporaryCredentials_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_temp_credentials": o.AwsTempCredentials,
			"azure_aad":            o.AzureAad,
			"expiration_time":      o.ExpirationTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TemporaryCredentials_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_temp_credentials": basetypes.ListType{
				ElemType: AwsCredentials_SdkV2{}.Type(ctx),
			},
			"azure_aad": basetypes.ListType{
				ElemType: AzureActiveDirectoryToken_SdkV2{}.Type(ctx),
			},
			"expiration_time": types.Int64Type,
		},
	}
}

// GetAwsTempCredentials returns the value of the AwsTempCredentials field in TemporaryCredentials_SdkV2 as
// a AwsCredentials_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TemporaryCredentials_SdkV2) GetAwsTempCredentials(ctx context.Context) (AwsCredentials_SdkV2, bool) {
	var e AwsCredentials_SdkV2
	if o.AwsTempCredentials.IsNull() || o.AwsTempCredentials.IsUnknown() {
		return e, false
	}
	var v []AwsCredentials_SdkV2
	d := o.AwsTempCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsTempCredentials sets the value of the AwsTempCredentials field in TemporaryCredentials_SdkV2.
func (o *TemporaryCredentials_SdkV2) SetAwsTempCredentials(ctx context.Context, v AwsCredentials_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_temp_credentials"]
	o.AwsTempCredentials = types.ListValueMust(t, vs)
}

// GetAzureAad returns the value of the AzureAad field in TemporaryCredentials_SdkV2 as
// a AzureActiveDirectoryToken_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TemporaryCredentials_SdkV2) GetAzureAad(ctx context.Context) (AzureActiveDirectoryToken_SdkV2, bool) {
	var e AzureActiveDirectoryToken_SdkV2
	if o.AzureAad.IsNull() || o.AzureAad.IsUnknown() {
		return e, false
	}
	var v []AzureActiveDirectoryToken_SdkV2
	d := o.AzureAad.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureAad sets the value of the AzureAad field in TemporaryCredentials_SdkV2.
func (o *TemporaryCredentials_SdkV2) SetAzureAad(ctx context.Context, v AzureActiveDirectoryToken_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_aad"]
	o.AzureAad = types.ListValueMust(t, vs)
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
type TriggeredUpdateStatus_SdkV2 struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version" tf:"optional"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp types.String `tfsdk:"timestamp" tf:"optional"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress types.List `tfsdk:"triggered_update_progress" tf:"optional,object"`
}

func (newState *TriggeredUpdateStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TriggeredUpdateStatus_SdkV2) {
}

func (newState *TriggeredUpdateStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState TriggeredUpdateStatus_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TriggeredUpdateStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TriggeredUpdateStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"triggered_update_progress": reflect.TypeOf(PipelineProgress_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TriggeredUpdateStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o TriggeredUpdateStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": o.LastProcessedCommitVersion,
			"timestamp":                     o.Timestamp,
			"triggered_update_progress":     o.TriggeredUpdateProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TriggeredUpdateStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
			"triggered_update_progress": basetypes.ListType{
				ElemType: PipelineProgress_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTriggeredUpdateProgress returns the value of the TriggeredUpdateProgress field in TriggeredUpdateStatus_SdkV2 as
// a PipelineProgress_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *TriggeredUpdateStatus_SdkV2) GetTriggeredUpdateProgress(ctx context.Context) (PipelineProgress_SdkV2, bool) {
	var e PipelineProgress_SdkV2
	if o.TriggeredUpdateProgress.IsNull() || o.TriggeredUpdateProgress.IsUnknown() {
		return e, false
	}
	var v []PipelineProgress_SdkV2
	d := o.TriggeredUpdateProgress.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTriggeredUpdateProgress sets the value of the TriggeredUpdateProgress field in TriggeredUpdateStatus_SdkV2.
func (o *TriggeredUpdateStatus_SdkV2) SetTriggeredUpdateProgress(ctx context.Context, v PipelineProgress_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["triggered_update_progress"]
	o.TriggeredUpdateProgress = types.ListValueMust(t, vs)
}

// Delete an assignment
type UnassignRequest_SdkV2 struct {
	// Query for the ID of the metastore to delete.
	MetastoreId types.String `tfsdk:"-"`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UnassignRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnassignRequest_SdkV2) {
}

func (newState *UnassignRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UnassignRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnassignRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnassignRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnassignRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UnassignRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metastore_id": o.MetastoreId,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UnassignRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

type UnassignResponse_SdkV2 struct {
}

func (newState *UnassignResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnassignResponse_SdkV2) {
}

func (newState *UnassignResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UnassignResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnassignResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnassignResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UnassignResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UnassignResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UnassignResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateAssignmentResponse_SdkV2 struct {
}

func (newState *UpdateAssignmentResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAssignmentResponse_SdkV2) {
}

func (newState *UpdateAssignmentResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateAssignmentResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAssignmentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAssignmentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAssignmentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateAssignmentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAssignmentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateCatalog_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization" tf:"optional"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
	// New name for the catalog.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of catalog.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties" tf:"optional"`
}

func (newState *UpdateCatalog_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCatalog_SdkV2) {
}

func (newState *UpdateCatalog_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateCatalog_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCatalog.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCatalog_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCatalog_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCatalog_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":                        o.Comment,
			"enable_predictive_optimization": o.EnablePredictiveOptimization,
			"isolation_mode":                 o.IsolationMode,
			"name":                           o.Name,
			"new_name":                       o.NewName,
			"owner":                          o.Owner,
			"properties":                     o.Properties,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCatalog_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":                        types.StringType,
			"enable_predictive_optimization": types.StringType,
			"isolation_mode":                 types.StringType,
			"name":                           types.StringType,
			"new_name":                       types.StringType,
			"owner":                          types.StringType,
			"properties": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetProperties returns the value of the Properties field in UpdateCatalog_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCatalog_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in UpdateCatalog_SdkV2.
func (o *UpdateCatalog_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type UpdateConnection_SdkV2 struct {
	// Name of the connection.
	Name types.String `tfsdk:"-"`
	// New name for the connection.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options" tf:""`
	// Username of current owner of the connection.
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateConnection_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateConnection_SdkV2) {
}

func (newState *UpdateConnection_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateConnection_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateConnection.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateConnection_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"options": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateConnection_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateConnection_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateConnection_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetOptions returns the value of the Options field in UpdateConnection_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateConnection_SdkV2) GetOptions(ctx context.Context) (map[string]types.String, bool) {
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

// SetOptions sets the value of the Options field in UpdateConnection_SdkV2.
func (o *UpdateConnection_SdkV2) SetOptions(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["options"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Options = types.MapValueMust(t, vs)
}

type UpdateCredentialRequest_SdkV2 struct {
	// The AWS IAM role configuration
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount types.List `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force types.Bool `tfsdk:"force" tf:"optional"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Name of the credential.
	NameArg types.String `tfsdk:"-"`
	// New name of credential.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of credential.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Supply true to this argument to skip validation of the updated
	// credential.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
}

func (newState *UpdateCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCredentialRequest_SdkV2) {
}

func (newState *UpdateCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRole_SdkV2{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentity_SdkV2{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal_SdkV2{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccount_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRole_SdkV2{}.Type(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentity_SdkV2{}.Type(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal_SdkV2{}.Type(ctx),
			},
			"comment": types.StringType,
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccount_SdkV2{}.Type(ctx),
			},
			"force":           types.BoolType,
			"isolation_mode":  types.StringType,
			"name_arg":        types.StringType,
			"new_name":        types.StringType,
			"owner":           types.StringType,
			"read_only":       types.BoolType,
			"skip_validation": types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in UpdateCredentialRequest_SdkV2 as
// a AwsIamRole_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCredentialRequest_SdkV2) GetAwsIamRole(ctx context.Context) (AwsIamRole_SdkV2, bool) {
	var e AwsIamRole_SdkV2
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRole_SdkV2
	d := o.AwsIamRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsIamRole sets the value of the AwsIamRole field in UpdateCredentialRequest_SdkV2.
func (o *UpdateCredentialRequest_SdkV2) SetAwsIamRole(ctx context.Context, v AwsIamRole_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_iam_role"]
	o.AwsIamRole = types.ListValueMust(t, vs)
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in UpdateCredentialRequest_SdkV2 as
// a AzureManagedIdentity_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCredentialRequest_SdkV2) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentity_SdkV2, bool) {
	var e AzureManagedIdentity_SdkV2
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentity_SdkV2
	d := o.AzureManagedIdentity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in UpdateCredentialRequest_SdkV2.
func (o *UpdateCredentialRequest_SdkV2) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentity_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_managed_identity"]
	o.AzureManagedIdentity = types.ListValueMust(t, vs)
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in UpdateCredentialRequest_SdkV2 as
// a AzureServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCredentialRequest_SdkV2) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal_SdkV2, bool) {
	var e AzureServicePrincipal_SdkV2
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal_SdkV2
	d := o.AzureServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in UpdateCredentialRequest_SdkV2.
func (o *UpdateCredentialRequest_SdkV2) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_service_principal"]
	o.AzureServicePrincipal = types.ListValueMust(t, vs)
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in UpdateCredentialRequest_SdkV2 as
// a DatabricksGcpServiceAccount_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCredentialRequest_SdkV2) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccount_SdkV2, bool) {
	var e DatabricksGcpServiceAccount_SdkV2
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccount_SdkV2
	d := o.DatabricksGcpServiceAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in UpdateCredentialRequest_SdkV2.
func (o *UpdateCredentialRequest_SdkV2) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccount_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["databricks_gcp_service_account"]
	o.DatabricksGcpServiceAccount = types.ListValueMust(t, vs)
}

type UpdateExternalLocation_SdkV2 struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the storage credential used with this location.
	CredentialName types.String `tfsdk:"credential_name" tf:"optional"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails types.List `tfsdk:"encryption_details" tf:"optional,object"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback types.Bool `tfsdk:"fallback" tf:"optional"`
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	Force types.Bool `tfsdk:"force" tf:"optional"`

	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Name of the external location.
	Name types.String `tfsdk:"-"`
	// New name for the external location.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// The owner of the external location.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Indicates whether the external location is read-only.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
	// Path URL of the external location.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *UpdateExternalLocation_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExternalLocation_SdkV2) {
}

func (newState *UpdateExternalLocation_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateExternalLocation_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExternalLocation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExternalLocation_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"encryption_details": reflect.TypeOf(EncryptionDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExternalLocation_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateExternalLocation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_point":       o.AccessPoint,
			"comment":            o.Comment,
			"credential_name":    o.CredentialName,
			"encryption_details": o.EncryptionDetails,
			"fallback":           o.Fallback,
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
func (o UpdateExternalLocation_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point":    types.StringType,
			"comment":         types.StringType,
			"credential_name": types.StringType,
			"encryption_details": basetypes.ListType{
				ElemType: EncryptionDetails_SdkV2{}.Type(ctx),
			},
			"fallback":        types.BoolType,
			"force":           types.BoolType,
			"isolation_mode":  types.StringType,
			"name":            types.StringType,
			"new_name":        types.StringType,
			"owner":           types.StringType,
			"read_only":       types.BoolType,
			"skip_validation": types.BoolType,
			"url":             types.StringType,
		},
	}
}

// GetEncryptionDetails returns the value of the EncryptionDetails field in UpdateExternalLocation_SdkV2 as
// a EncryptionDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateExternalLocation_SdkV2) GetEncryptionDetails(ctx context.Context) (EncryptionDetails_SdkV2, bool) {
	var e EncryptionDetails_SdkV2
	if o.EncryptionDetails.IsNull() || o.EncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []EncryptionDetails_SdkV2
	d := o.EncryptionDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEncryptionDetails sets the value of the EncryptionDetails field in UpdateExternalLocation_SdkV2.
func (o *UpdateExternalLocation_SdkV2) SetEncryptionDetails(ctx context.Context, v EncryptionDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["encryption_details"]
	o.EncryptionDetails = types.ListValueMust(t, vs)
}

type UpdateFunction_SdkV2 struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name types.String `tfsdk:"-"`
	// Username of current owner of function.
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateFunction_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateFunction_SdkV2) {
}

func (newState *UpdateFunction_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateFunction_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateFunction.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateFunction_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateFunction_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateFunction_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":  o.Name,
			"owner": o.Owner,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateFunction_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":  types.StringType,
			"owner": types.StringType,
		},
	}
}

type UpdateMetastore_SdkV2 struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName types.String `tfsdk:"delta_sharing_organization_name" tf:"optional"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds types.Int64 `tfsdk:"delta_sharing_recipient_token_lifetime_in_seconds" tf:"optional"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope types.String `tfsdk:"delta_sharing_scope" tf:"optional"`
	// Unique ID of the metastore.
	Id types.String `tfsdk:"-"`
	// New name for the metastore.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// The owner of the metastore.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion types.String `tfsdk:"privilege_model_version" tf:"optional"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId types.String `tfsdk:"storage_root_credential_id" tf:"optional"`
}

func (newState *UpdateMetastore_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMetastore_SdkV2) {
}

func (newState *UpdateMetastore_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateMetastore_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateMetastore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateMetastore_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateMetastore_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateMetastore_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateMetastore_SdkV2) Type(ctx context.Context) attr.Type {
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

type UpdateMetastoreAssignment_SdkV2 struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name" tf:"optional"`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UpdateMetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMetastoreAssignment_SdkV2) {
}

func (newState *UpdateMetastoreAssignment_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateMetastoreAssignment_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateMetastoreAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateMetastoreAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateMetastoreAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateMetastoreAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_catalog_name": o.DefaultCatalogName,
			"metastore_id":         o.MetastoreId,
			"workspace_id":         o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateMetastoreAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_catalog_name": types.StringType,
			"metastore_id":         types.StringType,
			"workspace_id":         types.Int64Type,
		},
	}
}

type UpdateModelVersionRequest_SdkV2 struct {
	// The comment attached to the model version
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The three-level (fully qualified) name of the model version
	FullName types.String `tfsdk:"-"`
	// The integer version number of the model version
	Version types.Int64 `tfsdk:"-"`
}

func (newState *UpdateModelVersionRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelVersionRequest_SdkV2) {
}

func (newState *UpdateModelVersionRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateModelVersionRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":   o.Comment,
			"full_name": o.FullName,
			"version":   o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":   types.StringType,
			"full_name": types.StringType,
			"version":   types.Int64Type,
		},
	}
}

type UpdateMonitor_SdkV2 struct {
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName types.String `tfsdk:"baseline_table_name" tf:"optional"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics types.List `tfsdk:"custom_metrics" tf:"optional"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The data classification config for the monitor.
	DataClassificationConfig types.List `tfsdk:"data_classification_config" tf:"optional,object"`
	// Configuration for monitoring inference logs.
	InferenceLog types.List `tfsdk:"inference_log" tf:"optional,object"`
	// The notification settings for the monitor.
	Notifications types.List `tfsdk:"notifications" tf:"optional,object"`
	// Schema where output metric tables are created.
	OutputSchemaName types.String `tfsdk:"output_schema_name" tf:""`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs types.List `tfsdk:"slicing_exprs" tf:"optional"`
	// Configuration for monitoring snapshot tables.
	Snapshot types.List `tfsdk:"snapshot" tf:"optional,object"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries types.List `tfsdk:"time_series" tf:"optional,object"`
}

func (newState *UpdateMonitor_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMonitor_SdkV2) {
}

func (newState *UpdateMonitor_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateMonitor_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateMonitor.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateMonitor_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_metrics":             reflect.TypeOf(MonitorMetric_SdkV2{}),
		"data_classification_config": reflect.TypeOf(MonitorDataClassificationConfig_SdkV2{}),
		"inference_log":              reflect.TypeOf(MonitorInferenceLog_SdkV2{}),
		"notifications":              reflect.TypeOf(MonitorNotifications_SdkV2{}),
		"schedule":                   reflect.TypeOf(MonitorCronSchedule_SdkV2{}),
		"slicing_exprs":              reflect.TypeOf(types.String{}),
		"snapshot":                   reflect.TypeOf(MonitorSnapshot_SdkV2{}),
		"time_series":                reflect.TypeOf(MonitorTimeSeries_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateMonitor_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateMonitor_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateMonitor_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: MonitorMetric_SdkV2{}.Type(ctx),
			},
			"dashboard_id": types.StringType,
			"data_classification_config": basetypes.ListType{
				ElemType: MonitorDataClassificationConfig_SdkV2{}.Type(ctx),
			},
			"inference_log": basetypes.ListType{
				ElemType: MonitorInferenceLog_SdkV2{}.Type(ctx),
			},
			"notifications": basetypes.ListType{
				ElemType: MonitorNotifications_SdkV2{}.Type(ctx),
			},
			"output_schema_name": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: MonitorCronSchedule_SdkV2{}.Type(ctx),
			},
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot": basetypes.ListType{
				ElemType: MonitorSnapshot_SdkV2{}.Type(ctx),
			},
			"table_name": types.StringType,
			"time_series": basetypes.ListType{
				ElemType: MonitorTimeSeries_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetCustomMetrics returns the value of the CustomMetrics field in UpdateMonitor_SdkV2 as
// a slice of MonitorMetric_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor_SdkV2) GetCustomMetrics(ctx context.Context) ([]MonitorMetric_SdkV2, bool) {
	if o.CustomMetrics.IsNull() || o.CustomMetrics.IsUnknown() {
		return nil, false
	}
	var v []MonitorMetric_SdkV2
	d := o.CustomMetrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomMetrics sets the value of the CustomMetrics field in UpdateMonitor_SdkV2.
func (o *UpdateMonitor_SdkV2) SetCustomMetrics(ctx context.Context, v []MonitorMetric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomMetrics = types.ListValueMust(t, vs)
}

// GetDataClassificationConfig returns the value of the DataClassificationConfig field in UpdateMonitor_SdkV2 as
// a MonitorDataClassificationConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor_SdkV2) GetDataClassificationConfig(ctx context.Context) (MonitorDataClassificationConfig_SdkV2, bool) {
	var e MonitorDataClassificationConfig_SdkV2
	if o.DataClassificationConfig.IsNull() || o.DataClassificationConfig.IsUnknown() {
		return e, false
	}
	var v []MonitorDataClassificationConfig_SdkV2
	d := o.DataClassificationConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataClassificationConfig sets the value of the DataClassificationConfig field in UpdateMonitor_SdkV2.
func (o *UpdateMonitor_SdkV2) SetDataClassificationConfig(ctx context.Context, v MonitorDataClassificationConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_classification_config"]
	o.DataClassificationConfig = types.ListValueMust(t, vs)
}

// GetInferenceLog returns the value of the InferenceLog field in UpdateMonitor_SdkV2 as
// a MonitorInferenceLog_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor_SdkV2) GetInferenceLog(ctx context.Context) (MonitorInferenceLog_SdkV2, bool) {
	var e MonitorInferenceLog_SdkV2
	if o.InferenceLog.IsNull() || o.InferenceLog.IsUnknown() {
		return e, false
	}
	var v []MonitorInferenceLog_SdkV2
	d := o.InferenceLog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInferenceLog sets the value of the InferenceLog field in UpdateMonitor_SdkV2.
func (o *UpdateMonitor_SdkV2) SetInferenceLog(ctx context.Context, v MonitorInferenceLog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inference_log"]
	o.InferenceLog = types.ListValueMust(t, vs)
}

// GetNotifications returns the value of the Notifications field in UpdateMonitor_SdkV2 as
// a MonitorNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor_SdkV2) GetNotifications(ctx context.Context) (MonitorNotifications_SdkV2, bool) {
	var e MonitorNotifications_SdkV2
	if o.Notifications.IsNull() || o.Notifications.IsUnknown() {
		return e, false
	}
	var v []MonitorNotifications_SdkV2
	d := o.Notifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotifications sets the value of the Notifications field in UpdateMonitor_SdkV2.
func (o *UpdateMonitor_SdkV2) SetNotifications(ctx context.Context, v MonitorNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notifications"]
	o.Notifications = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in UpdateMonitor_SdkV2 as
// a MonitorCronSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor_SdkV2) GetSchedule(ctx context.Context) (MonitorCronSchedule_SdkV2, bool) {
	var e MonitorCronSchedule_SdkV2
	if o.Schedule.IsNull() || o.Schedule.IsUnknown() {
		return e, false
	}
	var v []MonitorCronSchedule_SdkV2
	d := o.Schedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSchedule sets the value of the Schedule field in UpdateMonitor_SdkV2.
func (o *UpdateMonitor_SdkV2) SetSchedule(ctx context.Context, v MonitorCronSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["schedule"]
	o.Schedule = types.ListValueMust(t, vs)
}

// GetSlicingExprs returns the value of the SlicingExprs field in UpdateMonitor_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor_SdkV2) GetSlicingExprs(ctx context.Context) ([]types.String, bool) {
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

// SetSlicingExprs sets the value of the SlicingExprs field in UpdateMonitor_SdkV2.
func (o *UpdateMonitor_SdkV2) SetSlicingExprs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["slicing_exprs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SlicingExprs = types.ListValueMust(t, vs)
}

// GetSnapshot returns the value of the Snapshot field in UpdateMonitor_SdkV2 as
// a MonitorSnapshot_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor_SdkV2) GetSnapshot(ctx context.Context) (MonitorSnapshot_SdkV2, bool) {
	var e MonitorSnapshot_SdkV2
	if o.Snapshot.IsNull() || o.Snapshot.IsUnknown() {
		return e, false
	}
	var v []MonitorSnapshot_SdkV2
	d := o.Snapshot.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSnapshot sets the value of the Snapshot field in UpdateMonitor_SdkV2.
func (o *UpdateMonitor_SdkV2) SetSnapshot(ctx context.Context, v MonitorSnapshot_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["snapshot"]
	o.Snapshot = types.ListValueMust(t, vs)
}

// GetTimeSeries returns the value of the TimeSeries field in UpdateMonitor_SdkV2 as
// a MonitorTimeSeries_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateMonitor_SdkV2) GetTimeSeries(ctx context.Context) (MonitorTimeSeries_SdkV2, bool) {
	var e MonitorTimeSeries_SdkV2
	if o.TimeSeries.IsNull() || o.TimeSeries.IsUnknown() {
		return e, false
	}
	var v []MonitorTimeSeries_SdkV2
	d := o.TimeSeries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTimeSeries sets the value of the TimeSeries field in UpdateMonitor_SdkV2.
func (o *UpdateMonitor_SdkV2) SetTimeSeries(ctx context.Context, v MonitorTimeSeries_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["time_series"]
	o.TimeSeries = types.ListValueMust(t, vs)
}

type UpdatePermissions_SdkV2 struct {
	// Array of permissions change objects.
	Changes types.List `tfsdk:"changes" tf:"optional"`
	// Full name of securable.
	FullName types.String `tfsdk:"-"`
	// Type of securable.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *UpdatePermissions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePermissions_SdkV2) {
}

func (newState *UpdatePermissions_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdatePermissions_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"changes": reflect.TypeOf(PermissionsChange_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdatePermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"changes":        o.Changes,
			"full_name":      o.FullName,
			"securable_type": o.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"changes": basetypes.ListType{
				ElemType: PermissionsChange_SdkV2{}.Type(ctx),
			},
			"full_name":      types.StringType,
			"securable_type": types.StringType,
		},
	}
}

// GetChanges returns the value of the Changes field in UpdatePermissions_SdkV2 as
// a slice of PermissionsChange_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdatePermissions_SdkV2) GetChanges(ctx context.Context) ([]PermissionsChange_SdkV2, bool) {
	if o.Changes.IsNull() || o.Changes.IsUnknown() {
		return nil, false
	}
	var v []PermissionsChange_SdkV2
	d := o.Changes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChanges sets the value of the Changes field in UpdatePermissions_SdkV2.
func (o *UpdatePermissions_SdkV2) SetChanges(ctx context.Context, v []PermissionsChange_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["changes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Changes = types.ListValueMust(t, vs)
}

type UpdateRegisteredModelRequest_SdkV2 struct {
	// The comment attached to the registered model
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
	// New name for the registered model.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// The identifier of the user who owns the registered model
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateRegisteredModelRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRegisteredModelRequest_SdkV2) {
}

func (newState *UpdateRegisteredModelRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateRegisteredModelRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRegisteredModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRegisteredModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRegisteredModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateRegisteredModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateRegisteredModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":   types.StringType,
			"full_name": types.StringType,
			"new_name":  types.StringType,
			"owner":     types.StringType,
		},
	}
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

type UpdateSchema_SdkV2 struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization types.String `tfsdk:"enable_predictive_optimization" tf:"optional"`
	// Full name of the schema.
	FullName types.String `tfsdk:"-"`
	// New name for the schema.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of schema.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Properties types.Map `tfsdk:"properties" tf:"optional"`
}

func (newState *UpdateSchema_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSchema_SdkV2) {
}

func (newState *UpdateSchema_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateSchema_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSchema.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateSchema_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"properties": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSchema_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateSchema_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateSchema_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetProperties returns the value of the Properties field in UpdateSchema_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateSchema_SdkV2) GetProperties(ctx context.Context) (map[string]types.String, bool) {
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

// SetProperties sets the value of the Properties field in UpdateSchema_SdkV2.
func (o *UpdateSchema_SdkV2) SetProperties(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["properties"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Properties = types.MapValueMust(t, vs)
}

type UpdateStorageCredential_SdkV2 struct {
	// The AWS IAM role configuration.
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// The Cloudflare API token configuration.
	CloudflareApiToken types.List `tfsdk:"cloudflare_api_token" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount types.List `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// Force update even if there are dependent external locations or external
	// tables.
	Force types.Bool `tfsdk:"force" tf:"optional"`

	IsolationMode types.String `tfsdk:"isolation_mode" tf:"optional"`
	// Name of the storage credential.
	Name types.String `tfsdk:"-"`
	// New name for the storage credential.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of credential.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Supplying true to this argument skips validation of the updated
	// credential.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
}

func (newState *UpdateStorageCredential_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateStorageCredential_SdkV2) {
}

func (newState *UpdateStorageCredential_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateStorageCredential_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateStorageCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateStorageCredential_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRoleRequest_SdkV2{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentityResponse_SdkV2{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal_SdkV2{}),
		"cloudflare_api_token":           reflect.TypeOf(CloudflareApiToken_SdkV2{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccountRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateStorageCredential_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateStorageCredential_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateStorageCredential_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRoleRequest_SdkV2{}.Type(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentityResponse_SdkV2{}.Type(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal_SdkV2{}.Type(ctx),
			},
			"cloudflare_api_token": basetypes.ListType{
				ElemType: CloudflareApiToken_SdkV2{}.Type(ctx),
			},
			"comment": types.StringType,
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccountRequest_SdkV2{}.Type(ctx),
			},
			"force":           types.BoolType,
			"isolation_mode":  types.StringType,
			"name":            types.StringType,
			"new_name":        types.StringType,
			"owner":           types.StringType,
			"read_only":       types.BoolType,
			"skip_validation": types.BoolType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in UpdateStorageCredential_SdkV2 as
// a AwsIamRoleRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateStorageCredential_SdkV2) GetAwsIamRole(ctx context.Context) (AwsIamRoleRequest_SdkV2, bool) {
	var e AwsIamRoleRequest_SdkV2
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRoleRequest_SdkV2
	d := o.AwsIamRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsIamRole sets the value of the AwsIamRole field in UpdateStorageCredential_SdkV2.
func (o *UpdateStorageCredential_SdkV2) SetAwsIamRole(ctx context.Context, v AwsIamRoleRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_iam_role"]
	o.AwsIamRole = types.ListValueMust(t, vs)
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in UpdateStorageCredential_SdkV2 as
// a AzureManagedIdentityResponse_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateStorageCredential_SdkV2) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentityResponse_SdkV2, bool) {
	var e AzureManagedIdentityResponse_SdkV2
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentityResponse_SdkV2
	d := o.AzureManagedIdentity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in UpdateStorageCredential_SdkV2.
func (o *UpdateStorageCredential_SdkV2) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentityResponse_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_managed_identity"]
	o.AzureManagedIdentity = types.ListValueMust(t, vs)
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in UpdateStorageCredential_SdkV2 as
// a AzureServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateStorageCredential_SdkV2) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal_SdkV2, bool) {
	var e AzureServicePrincipal_SdkV2
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal_SdkV2
	d := o.AzureServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in UpdateStorageCredential_SdkV2.
func (o *UpdateStorageCredential_SdkV2) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_service_principal"]
	o.AzureServicePrincipal = types.ListValueMust(t, vs)
}

// GetCloudflareApiToken returns the value of the CloudflareApiToken field in UpdateStorageCredential_SdkV2 as
// a CloudflareApiToken_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateStorageCredential_SdkV2) GetCloudflareApiToken(ctx context.Context) (CloudflareApiToken_SdkV2, bool) {
	var e CloudflareApiToken_SdkV2
	if o.CloudflareApiToken.IsNull() || o.CloudflareApiToken.IsUnknown() {
		return e, false
	}
	var v []CloudflareApiToken_SdkV2
	d := o.CloudflareApiToken.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudflareApiToken sets the value of the CloudflareApiToken field in UpdateStorageCredential_SdkV2.
func (o *UpdateStorageCredential_SdkV2) SetCloudflareApiToken(ctx context.Context, v CloudflareApiToken_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cloudflare_api_token"]
	o.CloudflareApiToken = types.ListValueMust(t, vs)
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in UpdateStorageCredential_SdkV2 as
// a DatabricksGcpServiceAccountRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateStorageCredential_SdkV2) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccountRequest_SdkV2, bool) {
	var e DatabricksGcpServiceAccountRequest_SdkV2
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccountRequest_SdkV2
	d := o.DatabricksGcpServiceAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in UpdateStorageCredential_SdkV2.
func (o *UpdateStorageCredential_SdkV2) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccountRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["databricks_gcp_service_account"]
	o.DatabricksGcpServiceAccount = types.ListValueMust(t, vs)
}

// Update a table owner.
type UpdateTableRequest_SdkV2 struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`

	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateTableRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateTableRequest_SdkV2) {
}

func (newState *UpdateTableRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateTableRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": o.FullName,
			"owner":     o.Owner,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
			"owner":     types.StringType,
		},
	}
}

type UpdateVolumeRequestContent_SdkV2 struct {
	// The comment attached to the volume
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The three-level (fully qualified) name of the volume
	Name types.String `tfsdk:"-"`
	// New name for the volume.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// The identifier of the user who owns the volume
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateVolumeRequestContent_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateVolumeRequestContent_SdkV2) {
}

func (newState *UpdateVolumeRequestContent_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateVolumeRequestContent_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateVolumeRequestContent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateVolumeRequestContent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateVolumeRequestContent_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateVolumeRequestContent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateVolumeRequestContent_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":  types.StringType,
			"name":     types.StringType,
			"new_name": types.StringType,
			"owner":    types.StringType,
		},
	}
}

type UpdateWorkspaceBindings_SdkV2 struct {
	// A list of workspace IDs.
	AssignWorkspaces types.List `tfsdk:"assign_workspaces" tf:"optional"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
	// A list of workspace IDs.
	UnassignWorkspaces types.List `tfsdk:"unassign_workspaces" tf:"optional"`
}

func (newState *UpdateWorkspaceBindings_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceBindings_SdkV2) {
}

func (newState *UpdateWorkspaceBindings_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceBindings_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceBindings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceBindings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assign_workspaces":   reflect.TypeOf(types.Int64{}),
		"unassign_workspaces": reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceBindings_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceBindings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assign_workspaces":   o.AssignWorkspaces,
			"name":                o.Name,
			"unassign_workspaces": o.UnassignWorkspaces,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceBindings_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetAssignWorkspaces returns the value of the AssignWorkspaces field in UpdateWorkspaceBindings_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceBindings_SdkV2) GetAssignWorkspaces(ctx context.Context) ([]types.Int64, bool) {
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

// SetAssignWorkspaces sets the value of the AssignWorkspaces field in UpdateWorkspaceBindings_SdkV2.
func (o *UpdateWorkspaceBindings_SdkV2) SetAssignWorkspaces(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["assign_workspaces"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AssignWorkspaces = types.ListValueMust(t, vs)
}

// GetUnassignWorkspaces returns the value of the UnassignWorkspaces field in UpdateWorkspaceBindings_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceBindings_SdkV2) GetUnassignWorkspaces(ctx context.Context) ([]types.Int64, bool) {
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

// SetUnassignWorkspaces sets the value of the UnassignWorkspaces field in UpdateWorkspaceBindings_SdkV2.
func (o *UpdateWorkspaceBindings_SdkV2) SetUnassignWorkspaces(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["unassign_workspaces"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UnassignWorkspaces = types.ListValueMust(t, vs)
}

type UpdateWorkspaceBindingsParameters_SdkV2 struct {
	// List of workspace bindings
	Add types.List `tfsdk:"add" tf:"optional"`
	// List of workspace bindings
	Remove types.List `tfsdk:"remove" tf:"optional"`
	// The name of the securable.
	SecurableName types.String `tfsdk:"-"`
	// The type of the securable to bind to a workspace.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *UpdateWorkspaceBindingsParameters_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceBindingsParameters_SdkV2) {
}

func (newState *UpdateWorkspaceBindingsParameters_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceBindingsParameters_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceBindingsParameters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceBindingsParameters_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"add":    reflect.TypeOf(WorkspaceBinding_SdkV2{}),
		"remove": reflect.TypeOf(WorkspaceBinding_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceBindingsParameters_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceBindingsParameters_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateWorkspaceBindingsParameters_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"add": basetypes.ListType{
				ElemType: WorkspaceBinding_SdkV2{}.Type(ctx),
			},
			"remove": basetypes.ListType{
				ElemType: WorkspaceBinding_SdkV2{}.Type(ctx),
			},
			"securable_name": types.StringType,
			"securable_type": types.StringType,
		},
	}
}

// GetAdd returns the value of the Add field in UpdateWorkspaceBindingsParameters_SdkV2 as
// a slice of WorkspaceBinding_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceBindingsParameters_SdkV2) GetAdd(ctx context.Context) ([]WorkspaceBinding_SdkV2, bool) {
	if o.Add.IsNull() || o.Add.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceBinding_SdkV2
	d := o.Add.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAdd sets the value of the Add field in UpdateWorkspaceBindingsParameters_SdkV2.
func (o *UpdateWorkspaceBindingsParameters_SdkV2) SetAdd(ctx context.Context, v []WorkspaceBinding_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["add"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Add = types.ListValueMust(t, vs)
}

// GetRemove returns the value of the Remove field in UpdateWorkspaceBindingsParameters_SdkV2 as
// a slice of WorkspaceBinding_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceBindingsParameters_SdkV2) GetRemove(ctx context.Context) ([]WorkspaceBinding_SdkV2, bool) {
	if o.Remove.IsNull() || o.Remove.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceBinding_SdkV2
	d := o.Remove.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRemove sets the value of the Remove field in UpdateWorkspaceBindingsParameters_SdkV2.
func (o *UpdateWorkspaceBindingsParameters_SdkV2) SetRemove(ctx context.Context, v []WorkspaceBinding_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["remove"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Remove = types.ListValueMust(t, vs)
}

type ValidateCredentialRequest_SdkV2 struct {
	// The AWS IAM role configuration
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	CredentialName types.String `tfsdk:"credential_name" tf:"optional"`
	// The name of an existing external location to validate. Only applicable
	// for storage credentials (purpose is **STORAGE**.)
	ExternalLocationName types.String `tfsdk:"external_location_name" tf:"optional"`
	// The purpose of the credential. This should only be used when the
	// credential is specified.
	Purpose types.String `tfsdk:"purpose" tf:"optional"`
	// Whether the credential is only usable for read operations. Only
	// applicable for storage credentials (purpose is **STORAGE**.)
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// The external location url to validate. Only applicable when purpose is
	// **STORAGE**.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *ValidateCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateCredentialRequest_SdkV2) {
}

func (newState *ValidateCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ValidateCredentialRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidateCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ValidateCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":           reflect.TypeOf(AwsIamRole_SdkV2{}),
		"azure_managed_identity": reflect.TypeOf(AzureManagedIdentity_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidateCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ValidateCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_iam_role":           o.AwsIamRole,
			"azure_managed_identity": o.AzureManagedIdentity,
			"credential_name":        o.CredentialName,
			"external_location_name": o.ExternalLocationName,
			"purpose":                o.Purpose,
			"read_only":              o.ReadOnly,
			"url":                    o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ValidateCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRole_SdkV2{}.Type(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentity_SdkV2{}.Type(ctx),
			},
			"credential_name":        types.StringType,
			"external_location_name": types.StringType,
			"purpose":                types.StringType,
			"read_only":              types.BoolType,
			"url":                    types.StringType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in ValidateCredentialRequest_SdkV2 as
// a AwsIamRole_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateCredentialRequest_SdkV2) GetAwsIamRole(ctx context.Context) (AwsIamRole_SdkV2, bool) {
	var e AwsIamRole_SdkV2
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRole_SdkV2
	d := o.AwsIamRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsIamRole sets the value of the AwsIamRole field in ValidateCredentialRequest_SdkV2.
func (o *ValidateCredentialRequest_SdkV2) SetAwsIamRole(ctx context.Context, v AwsIamRole_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_iam_role"]
	o.AwsIamRole = types.ListValueMust(t, vs)
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in ValidateCredentialRequest_SdkV2 as
// a AzureManagedIdentity_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateCredentialRequest_SdkV2) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentity_SdkV2, bool) {
	var e AzureManagedIdentity_SdkV2
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentity_SdkV2
	d := o.AzureManagedIdentity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in ValidateCredentialRequest_SdkV2.
func (o *ValidateCredentialRequest_SdkV2) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentity_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_managed_identity"]
	o.AzureManagedIdentity = types.ListValueMust(t, vs)
}

type ValidateCredentialResponse_SdkV2 struct {
	// Whether the tested location is a directory in cloud storage. Only
	// applicable for when purpose is **STORAGE**.
	IsDir types.Bool `tfsdk:"isDir" tf:"optional"`
	// The results of the validation check.
	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ValidateCredentialResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateCredentialResponse_SdkV2) {
}

func (newState *ValidateCredentialResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ValidateCredentialResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidateCredentialResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ValidateCredentialResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(CredentialValidationResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidateCredentialResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ValidateCredentialResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"isDir":   o.IsDir,
			"results": o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ValidateCredentialResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"isDir": types.BoolType,
			"results": basetypes.ListType{
				ElemType: CredentialValidationResult_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ValidateCredentialResponse_SdkV2 as
// a slice of CredentialValidationResult_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateCredentialResponse_SdkV2) GetResults(ctx context.Context) ([]CredentialValidationResult_SdkV2, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []CredentialValidationResult_SdkV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ValidateCredentialResponse_SdkV2.
func (o *ValidateCredentialResponse_SdkV2) SetResults(ctx context.Context, v []CredentialValidationResult_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ValidateStorageCredential_SdkV2 struct {
	// The AWS IAM role configuration.
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// The Cloudflare API token configuration.
	CloudflareApiToken types.List `tfsdk:"cloudflare_api_token" tf:"optional,object"`
	// The Databricks created GCP service account configuration.
	DatabricksGcpServiceAccount types.List `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// The name of an existing external location to validate.
	ExternalLocationName types.String `tfsdk:"external_location_name" tf:"optional"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// The name of the storage credential to validate.
	StorageCredentialName types.String `tfsdk:"storage_credential_name" tf:"optional"`
	// The external location url to validate.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *ValidateStorageCredential_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateStorageCredential_SdkV2) {
}

func (newState *ValidateStorageCredential_SdkV2) SyncEffectiveFieldsDuringRead(existingState ValidateStorageCredential_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidateStorageCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ValidateStorageCredential_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_iam_role":                   reflect.TypeOf(AwsIamRoleRequest_SdkV2{}),
		"azure_managed_identity":         reflect.TypeOf(AzureManagedIdentityRequest_SdkV2{}),
		"azure_service_principal":        reflect.TypeOf(AzureServicePrincipal_SdkV2{}),
		"cloudflare_api_token":           reflect.TypeOf(CloudflareApiToken_SdkV2{}),
		"databricks_gcp_service_account": reflect.TypeOf(DatabricksGcpServiceAccountRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidateStorageCredential_SdkV2
// only implements ToObjectValue() and Type().
func (o ValidateStorageCredential_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ValidateStorageCredential_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRoleRequest_SdkV2{}.Type(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentityRequest_SdkV2{}.Type(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal_SdkV2{}.Type(ctx),
			},
			"cloudflare_api_token": basetypes.ListType{
				ElemType: CloudflareApiToken_SdkV2{}.Type(ctx),
			},
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccountRequest_SdkV2{}.Type(ctx),
			},
			"external_location_name":  types.StringType,
			"read_only":               types.BoolType,
			"storage_credential_name": types.StringType,
			"url":                     types.StringType,
		},
	}
}

// GetAwsIamRole returns the value of the AwsIamRole field in ValidateStorageCredential_SdkV2 as
// a AwsIamRoleRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredential_SdkV2) GetAwsIamRole(ctx context.Context) (AwsIamRoleRequest_SdkV2, bool) {
	var e AwsIamRoleRequest_SdkV2
	if o.AwsIamRole.IsNull() || o.AwsIamRole.IsUnknown() {
		return e, false
	}
	var v []AwsIamRoleRequest_SdkV2
	d := o.AwsIamRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsIamRole sets the value of the AwsIamRole field in ValidateStorageCredential_SdkV2.
func (o *ValidateStorageCredential_SdkV2) SetAwsIamRole(ctx context.Context, v AwsIamRoleRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_iam_role"]
	o.AwsIamRole = types.ListValueMust(t, vs)
}

// GetAzureManagedIdentity returns the value of the AzureManagedIdentity field in ValidateStorageCredential_SdkV2 as
// a AzureManagedIdentityRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredential_SdkV2) GetAzureManagedIdentity(ctx context.Context) (AzureManagedIdentityRequest_SdkV2, bool) {
	var e AzureManagedIdentityRequest_SdkV2
	if o.AzureManagedIdentity.IsNull() || o.AzureManagedIdentity.IsUnknown() {
		return e, false
	}
	var v []AzureManagedIdentityRequest_SdkV2
	d := o.AzureManagedIdentity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureManagedIdentity sets the value of the AzureManagedIdentity field in ValidateStorageCredential_SdkV2.
func (o *ValidateStorageCredential_SdkV2) SetAzureManagedIdentity(ctx context.Context, v AzureManagedIdentityRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_managed_identity"]
	o.AzureManagedIdentity = types.ListValueMust(t, vs)
}

// GetAzureServicePrincipal returns the value of the AzureServicePrincipal field in ValidateStorageCredential_SdkV2 as
// a AzureServicePrincipal_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredential_SdkV2) GetAzureServicePrincipal(ctx context.Context) (AzureServicePrincipal_SdkV2, bool) {
	var e AzureServicePrincipal_SdkV2
	if o.AzureServicePrincipal.IsNull() || o.AzureServicePrincipal.IsUnknown() {
		return e, false
	}
	var v []AzureServicePrincipal_SdkV2
	d := o.AzureServicePrincipal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureServicePrincipal sets the value of the AzureServicePrincipal field in ValidateStorageCredential_SdkV2.
func (o *ValidateStorageCredential_SdkV2) SetAzureServicePrincipal(ctx context.Context, v AzureServicePrincipal_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_service_principal"]
	o.AzureServicePrincipal = types.ListValueMust(t, vs)
}

// GetCloudflareApiToken returns the value of the CloudflareApiToken field in ValidateStorageCredential_SdkV2 as
// a CloudflareApiToken_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredential_SdkV2) GetCloudflareApiToken(ctx context.Context) (CloudflareApiToken_SdkV2, bool) {
	var e CloudflareApiToken_SdkV2
	if o.CloudflareApiToken.IsNull() || o.CloudflareApiToken.IsUnknown() {
		return e, false
	}
	var v []CloudflareApiToken_SdkV2
	d := o.CloudflareApiToken.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudflareApiToken sets the value of the CloudflareApiToken field in ValidateStorageCredential_SdkV2.
func (o *ValidateStorageCredential_SdkV2) SetCloudflareApiToken(ctx context.Context, v CloudflareApiToken_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cloudflare_api_token"]
	o.CloudflareApiToken = types.ListValueMust(t, vs)
}

// GetDatabricksGcpServiceAccount returns the value of the DatabricksGcpServiceAccount field in ValidateStorageCredential_SdkV2 as
// a DatabricksGcpServiceAccountRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredential_SdkV2) GetDatabricksGcpServiceAccount(ctx context.Context) (DatabricksGcpServiceAccountRequest_SdkV2, bool) {
	var e DatabricksGcpServiceAccountRequest_SdkV2
	if o.DatabricksGcpServiceAccount.IsNull() || o.DatabricksGcpServiceAccount.IsUnknown() {
		return e, false
	}
	var v []DatabricksGcpServiceAccountRequest_SdkV2
	d := o.DatabricksGcpServiceAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabricksGcpServiceAccount sets the value of the DatabricksGcpServiceAccount field in ValidateStorageCredential_SdkV2.
func (o *ValidateStorageCredential_SdkV2) SetDatabricksGcpServiceAccount(ctx context.Context, v DatabricksGcpServiceAccountRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["databricks_gcp_service_account"]
	o.DatabricksGcpServiceAccount = types.ListValueMust(t, vs)
}

type ValidateStorageCredentialResponse_SdkV2 struct {
	// Whether the tested location is a directory in cloud storage.
	IsDir types.Bool `tfsdk:"isDir" tf:"optional"`
	// The results of the validation check.
	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ValidateStorageCredentialResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateStorageCredentialResponse_SdkV2) {
}

func (newState *ValidateStorageCredentialResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ValidateStorageCredentialResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidateStorageCredentialResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ValidateStorageCredentialResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ValidationResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidateStorageCredentialResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ValidateStorageCredentialResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"isDir":   o.IsDir,
			"results": o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ValidateStorageCredentialResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"isDir": types.BoolType,
			"results": basetypes.ListType{
				ElemType: ValidationResult_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ValidateStorageCredentialResponse_SdkV2 as
// a slice of ValidationResult_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ValidateStorageCredentialResponse_SdkV2) GetResults(ctx context.Context) ([]ValidationResult_SdkV2, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []ValidationResult_SdkV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ValidateStorageCredentialResponse_SdkV2.
func (o *ValidateStorageCredentialResponse_SdkV2) SetResults(ctx context.Context, v []ValidationResult_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
}

type ValidationResult_SdkV2 struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message types.String `tfsdk:"message" tf:"optional"`
	// The operation tested.
	Operation types.String `tfsdk:"operation" tf:"optional"`
	// The results of the tested operation.
	Result types.String `tfsdk:"result" tf:"optional"`
}

func (newState *ValidationResult_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidationResult_SdkV2) {
}

func (newState *ValidationResult_SdkV2) SyncEffectiveFieldsDuringRead(existingState ValidationResult_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidationResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ValidationResult_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidationResult_SdkV2
// only implements ToObjectValue() and Type().
func (o ValidationResult_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message":   o.Message,
			"operation": o.Operation,
			"result":    o.Result,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ValidationResult_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message":   types.StringType,
			"operation": types.StringType,
			"result":    types.StringType,
		},
	}
}

type VolumeInfo_SdkV2 struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint types.String `tfsdk:"access_point" tf:"optional"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only" tf:"optional"`
	// The name of the catalog where the schema and the volume are
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// The comment attached to the volume
	Comment types.String `tfsdk:"comment" tf:"optional"`

	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The identifier of the user who created the volume
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails types.List `tfsdk:"encryption_details" tf:"optional,object"`
	// The three-level (fully qualified) name of the volume
	FullName types.String `tfsdk:"full_name" tf:"optional"`
	// The unique identifier of the metastore
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// The name of the volume
	Name types.String `tfsdk:"name" tf:"optional"`
	// The identifier of the user who owns the volume
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// The name of the schema where the volume is
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// The storage location on the cloud
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`

	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// The identifier of the user who updated the volume last time
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
	// The unique identifier of the volume
	VolumeId types.String `tfsdk:"volume_id" tf:"optional"`

	VolumeType types.String `tfsdk:"volume_type" tf:"optional"`
}

func (newState *VolumeInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan VolumeInfo_SdkV2) {
}

func (newState *VolumeInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState VolumeInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VolumeInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a VolumeInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"encryption_details": reflect.TypeOf(EncryptionDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VolumeInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o VolumeInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o VolumeInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point": types.StringType,
			"browse_only":  types.BoolType,
			"catalog_name": types.StringType,
			"comment":      types.StringType,
			"created_at":   types.Int64Type,
			"created_by":   types.StringType,
			"encryption_details": basetypes.ListType{
				ElemType: EncryptionDetails_SdkV2{}.Type(ctx),
			},
			"full_name":        types.StringType,
			"metastore_id":     types.StringType,
			"name":             types.StringType,
			"owner":            types.StringType,
			"schema_name":      types.StringType,
			"storage_location": types.StringType,
			"updated_at":       types.Int64Type,
			"updated_by":       types.StringType,
			"volume_id":        types.StringType,
			"volume_type":      types.StringType,
		},
	}
}

// GetEncryptionDetails returns the value of the EncryptionDetails field in VolumeInfo_SdkV2 as
// a EncryptionDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *VolumeInfo_SdkV2) GetEncryptionDetails(ctx context.Context) (EncryptionDetails_SdkV2, bool) {
	var e EncryptionDetails_SdkV2
	if o.EncryptionDetails.IsNull() || o.EncryptionDetails.IsUnknown() {
		return e, false
	}
	var v []EncryptionDetails_SdkV2
	d := o.EncryptionDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEncryptionDetails sets the value of the EncryptionDetails field in VolumeInfo_SdkV2.
func (o *VolumeInfo_SdkV2) SetEncryptionDetails(ctx context.Context, v EncryptionDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["encryption_details"]
	o.EncryptionDetails = types.ListValueMust(t, vs)
}

type WorkspaceBinding_SdkV2 struct {
	BindingType types.String `tfsdk:"binding_type" tf:"optional"`

	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
}

func (newState *WorkspaceBinding_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceBinding_SdkV2) {
}

func (newState *WorkspaceBinding_SdkV2) SyncEffectiveFieldsDuringRead(existingState WorkspaceBinding_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceBinding.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceBinding_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceBinding_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceBinding_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"binding_type": o.BindingType,
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceBinding_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"binding_type": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// Currently assigned workspace bindings
type WorkspaceBindingsResponse_SdkV2 struct {
	// List of workspace bindings
	Bindings types.List `tfsdk:"bindings" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *WorkspaceBindingsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceBindingsResponse_SdkV2) {
}

func (newState *WorkspaceBindingsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState WorkspaceBindingsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceBindingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceBindingsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"bindings": reflect.TypeOf(WorkspaceBinding_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceBindingsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceBindingsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bindings":        o.Bindings,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceBindingsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bindings": basetypes.ListType{
				ElemType: WorkspaceBinding_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetBindings returns the value of the Bindings field in WorkspaceBindingsResponse_SdkV2 as
// a slice of WorkspaceBinding_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceBindingsResponse_SdkV2) GetBindings(ctx context.Context) ([]WorkspaceBinding_SdkV2, bool) {
	if o.Bindings.IsNull() || o.Bindings.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceBinding_SdkV2
	d := o.Bindings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBindings sets the value of the Bindings field in WorkspaceBindingsResponse_SdkV2.
func (o *WorkspaceBindingsResponse_SdkV2) SetBindings(ctx context.Context, v []WorkspaceBinding_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["bindings"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Bindings = types.ListValueMust(t, vs)
}
