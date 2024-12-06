// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package iam_tf

import (
	"context"
	"fmt"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type AccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *AccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccessControlRequest) {
}

func (newState *AccessControlRequest) SyncEffectiveFieldsDuringRead(existingState AccessControlRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = AccessControlRequest{}

// Equal implements basetypes.ObjectValuable.
func (o AccessControlRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o AccessControlRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o AccessControlRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o AccessControlRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o AccessControlRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o AccessControlRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o AccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type AccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *AccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccessControlResponse) {
}

func (newState *AccessControlResponse) SyncEffectiveFieldsDuringRead(existingState AccessControlResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(Permission{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = AccessControlResponse{}

// Equal implements basetypes.ObjectValuable.
func (o AccessControlResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o AccessControlResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o AccessControlResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o AccessControlResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o AccessControlResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o AccessControlResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o AccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: Permission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ComplexValue struct {
	Display types.String `tfsdk:"display" tf:"optional"`

	Primary types.Bool `tfsdk:"primary" tf:"optional"`

	Ref types.String `tfsdk:"$ref" tf:"optional"`

	Type_ types.String `tfsdk:"type" tf:"optional"`

	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *ComplexValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplexValue) {
}

func (newState *ComplexValue) SyncEffectiveFieldsDuringRead(existingState ComplexValue) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComplexValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ComplexValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ComplexValue{}

// Equal implements basetypes.ObjectValuable.
func (o ComplexValue) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ComplexValue) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ComplexValue) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ComplexValue) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ComplexValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ComplexValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ComplexValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display": types.StringType,
			"primary": types.BoolType,
			"$ref":    types.StringType,
			"type":    types.StringType,
			"value":   types.StringType,
		},
	}
}

// Delete a group
type DeleteAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteAccountGroupRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountGroupRequest) {
}

func (newState *DeleteAccountGroupRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountGroupRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteAccountGroupRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteAccountGroupRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteAccountGroupRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteAccountGroupRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteAccountGroupRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteAccountGroupRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteAccountGroupRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a service principal
type DeleteAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteAccountServicePrincipalRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountServicePrincipalRequest) {
}

func (newState *DeleteAccountServicePrincipalRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountServicePrincipalRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteAccountServicePrincipalRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteAccountServicePrincipalRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteAccountServicePrincipalRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteAccountServicePrincipalRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteAccountServicePrincipalRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteAccountServicePrincipalRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteAccountServicePrincipalRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a user
type DeleteAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteAccountUserRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountUserRequest) {
}

func (newState *DeleteAccountUserRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountUserRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteAccountUserRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteAccountUserRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteAccountUserRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteAccountUserRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteAccountUserRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteAccountUserRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteAccountUserRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a group
type DeleteGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteGroupRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteGroupRequest) {
}

func (newState *DeleteGroupRequest) SyncEffectiveFieldsDuringRead(existingState DeleteGroupRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteGroupRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteGroupRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteGroupRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteGroupRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteGroupRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteGroupRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteGroupRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a service principal
type DeleteServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteServicePrincipalRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteServicePrincipalRequest) {
}

func (newState *DeleteServicePrincipalRequest) SyncEffectiveFieldsDuringRead(existingState DeleteServicePrincipalRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteServicePrincipalRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a user
type DeleteUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteUserRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteUserRequest) {
}

func (newState *DeleteUserRequest) SyncEffectiveFieldsDuringRead(existingState DeleteUserRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteUserRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteUserRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteUserRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteUserRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteUserRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteUserRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteUserRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete permissions assignment
type DeleteWorkspaceAssignmentRequest struct {
	// The ID of the user, service principal, or group.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID for the account.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *DeleteWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWorkspaceAssignmentRequest) {
}

func (newState *DeleteWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringRead(existingState DeleteWorkspaceAssignmentRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteWorkspaceAssignmentRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteWorkspaceAssignmentRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteWorkspaceAssignmentRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteWorkspaceAssignmentRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteWorkspaceAssignmentRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteWorkspaceAssignmentRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteWorkspaceAssignmentRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWorkspaceAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type DeleteWorkspacePermissionAssignmentResponse struct {
}

func (newState *DeleteWorkspacePermissionAssignmentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWorkspacePermissionAssignmentResponse) {
}

func (newState *DeleteWorkspacePermissionAssignmentResponse) SyncEffectiveFieldsDuringRead(existingState DeleteWorkspacePermissionAssignmentResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspacePermissionAssignmentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWorkspacePermissionAssignmentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteWorkspacePermissionAssignmentResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteWorkspacePermissionAssignmentResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteWorkspacePermissionAssignmentResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteWorkspacePermissionAssignmentResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteWorkspacePermissionAssignmentResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteWorkspacePermissionAssignmentResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteWorkspacePermissionAssignmentResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWorkspacePermissionAssignmentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Get group details
type GetAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id types.String `tfsdk:"-"`
}

func (newState *GetAccountGroupRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountGroupRequest) {
}

func (newState *GetAccountGroupRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountGroupRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetAccountGroupRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetAccountGroupRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetAccountGroupRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetAccountGroupRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetAccountGroupRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetAccountGroupRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetAccountGroupRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get service principal details
type GetAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id types.String `tfsdk:"-"`
}

func (newState *GetAccountServicePrincipalRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountServicePrincipalRequest) {
}

func (newState *GetAccountServicePrincipalRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountServicePrincipalRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetAccountServicePrincipalRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetAccountServicePrincipalRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetAccountServicePrincipalRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetAccountServicePrincipalRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetAccountServicePrincipalRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetAccountServicePrincipalRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetAccountServicePrincipalRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get user details
type GetAccountUserRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Unique ID for a user in the Databricks account.
	Id types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *GetAccountUserRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountUserRequest) {
}

func (newState *GetAccountUserRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountUserRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetAccountUserRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetAccountUserRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetAccountUserRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetAccountUserRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetAccountUserRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetAccountUserRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetAccountUserRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"id":                 types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// Get assignable roles for a resource
type GetAssignableRolesForResourceRequest struct {
	// The resource name for which assignable roles will be listed.
	Resource types.String `tfsdk:"-"`
}

func (newState *GetAssignableRolesForResourceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAssignableRolesForResourceRequest) {
}

func (newState *GetAssignableRolesForResourceRequest) SyncEffectiveFieldsDuringRead(existingState GetAssignableRolesForResourceRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAssignableRolesForResourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAssignableRolesForResourceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetAssignableRolesForResourceRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resource": types.StringType,
		},
	}
}

type GetAssignableRolesForResourceResponse struct {
	Roles types.List `tfsdk:"roles" tf:"optional"`
}

func (newState *GetAssignableRolesForResourceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAssignableRolesForResourceResponse) {
}

func (newState *GetAssignableRolesForResourceResponse) SyncEffectiveFieldsDuringRead(existingState GetAssignableRolesForResourceResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAssignableRolesForResourceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAssignableRolesForResourceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"roles": reflect.TypeOf(Role{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetAssignableRolesForResourceResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetAssignableRolesForResourceResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"roles": basetypes.ListType{
				ElemType: Role{}.Type(ctx),
			},
		},
	}
}

// Get group details
type GetGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

func (newState *GetGroupRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetGroupRequest) {
}

func (newState *GetGroupRequest) SyncEffectiveFieldsDuringRead(existingState GetGroupRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetGroupRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetGroupRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetGroupRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetGroupRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetGroupRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetGroupRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetGroupRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetPasswordPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetPasswordPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPasswordPermissionLevelsResponse) {
}

func (newState *GetPasswordPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetPasswordPermissionLevelsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPasswordPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPasswordPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PasswordPermissionsDescription{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetPasswordPermissionLevelsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetPasswordPermissionLevelsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetPasswordPermissionLevelsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetPasswordPermissionLevelsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetPasswordPermissionLevelsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetPasswordPermissionLevelsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetPasswordPermissionLevelsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetPasswordPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PasswordPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// Get object permission levels
type GetPermissionLevelsRequest struct {
	// <needs content>
	RequestObjectId types.String `tfsdk:"-"`
	// <needs content>
	RequestObjectType types.String `tfsdk:"-"`
}

func (newState *GetPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPermissionLevelsRequest) {
}

func (newState *GetPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetPermissionLevelsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetPermissionLevelsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetPermissionLevelsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetPermissionLevelsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetPermissionLevelsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetPermissionLevelsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetPermissionLevelsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetPermissionLevelsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

type GetPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPermissionLevelsResponse) {
}

func (newState *GetPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetPermissionLevelsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PermissionsDescription{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetPermissionLevelsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetPermissionLevelsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetPermissionLevelsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetPermissionLevelsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetPermissionLevelsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetPermissionLevelsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetPermissionLevelsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// Get object permissions
type GetPermissionRequest struct {
	// The id of the request object.
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

func (newState *GetPermissionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPermissionRequest) {
}

func (newState *GetPermissionRequest) SyncEffectiveFieldsDuringRead(existingState GetPermissionRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetPermissionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetPermissionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetPermissionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetPermissionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetPermissionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetPermissionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetPermissionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetPermissionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

// Get a rule set
type GetRuleSetRequest struct {
	// Etag used for versioning. The response is at least as fresh as the eTag
	// provided. Etag is used for optimistic concurrency control as a way to
	// help prevent simultaneous updates of a rule set from overwriting each
	// other. It is strongly suggested that systems make use of the etag in the
	// read -> modify -> write pattern to perform rule set updates in order to
	// avoid race conditions that is get an etag from a GET rule set request,
	// and pass it with the PUT update request to identify the rule set version
	// you are updating.
	Etag types.String `tfsdk:"-"`
	// The ruleset name associated with the request.
	Name types.String `tfsdk:"-"`
}

func (newState *GetRuleSetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRuleSetRequest) {
}

func (newState *GetRuleSetRequest) SyncEffectiveFieldsDuringRead(existingState GetRuleSetRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRuleSetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRuleSetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetRuleSetRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetRuleSetRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetRuleSetRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetRuleSetRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetRuleSetRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetRuleSetRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetRuleSetRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetRuleSetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"name": types.StringType,
		},
	}
}

// Get service principal details
type GetServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

func (newState *GetServicePrincipalRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServicePrincipalRequest) {
}

func (newState *GetServicePrincipalRequest) SyncEffectiveFieldsDuringRead(existingState GetServicePrincipalRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetServicePrincipalRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetServicePrincipalRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetServicePrincipalRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetServicePrincipalRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetServicePrincipalRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetServicePrincipalRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetServicePrincipalRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetServicePrincipalRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get user details
type GetUserRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Unique ID for a user in the Databricks workspace.
	Id types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *GetUserRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetUserRequest) {
}

func (newState *GetUserRequest) SyncEffectiveFieldsDuringRead(existingState GetUserRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetUserRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetUserRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetUserRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetUserRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetUserRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetUserRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetUserRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetUserRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"id":                 types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// List workspace permissions
type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *GetWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceAssignmentRequest) {
}

func (newState *GetWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceAssignmentRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetWorkspaceAssignmentRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetWorkspaceAssignmentRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetWorkspaceAssignmentRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetWorkspaceAssignmentRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetWorkspaceAssignmentRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetWorkspaceAssignmentRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetWorkspaceAssignmentRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type GrantRule struct {
	// Principals this grant rule applies to.
	Principals types.List `tfsdk:"principals" tf:"optional"`
	// Role that is assigned to the list of principals.
	Role types.String `tfsdk:"role" tf:""`
}

func (newState *GrantRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan GrantRule) {
}

func (newState *GrantRule) SyncEffectiveFieldsDuringRead(existingState GrantRule) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GrantRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GrantRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"principals": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GrantRule{}

// Equal implements basetypes.ObjectValuable.
func (o GrantRule) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GrantRule) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GrantRule) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GrantRule) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GrantRule) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GrantRule) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GrantRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principals": basetypes.ListType{
				ElemType: types.StringType,
			},
			"role": types.StringType,
		},
	}
}

type Group struct {
	// String that represents a human-readable group name
	DisplayName types.String `tfsdk:"displayName" tf:"optional"`
	// Entitlements assigned to the group. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements types.List `tfsdk:"entitlements" tf:"optional"`

	ExternalId types.String `tfsdk:"externalId" tf:"optional"`

	Groups types.List `tfsdk:"groups" tf:"optional"`
	// Databricks group ID
	Id types.String `tfsdk:"id" tf:"optional"`

	Members types.List `tfsdk:"members" tf:"optional"`
	// Container for the group identifier. Workspace local versus account.
	Meta types.List `tfsdk:"meta" tf:"optional,object"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles" tf:"optional"`
	// The schema of the group.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
}

func (newState *Group) SyncEffectiveFieldsDuringCreateOrUpdate(plan Group) {
}

func (newState *Group) SyncEffectiveFieldsDuringRead(existingState Group) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Group.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Group) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"members":      reflect.TypeOf(ComplexValue{}),
		"meta":         reflect.TypeOf(ResourceMeta{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Group{}

// Equal implements basetypes.ObjectValuable.
func (o Group) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Group) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Group) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Group) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Group) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Group) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Group) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"displayName": types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"id": types.StringType,
			"members": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"meta": basetypes.ListType{
				ElemType: ResourceMeta{}.Type(ctx),
			},
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// List group details
type ListAccountGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListAccountGroupsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountGroupsRequest) {
}

func (newState *ListAccountGroupsRequest) SyncEffectiveFieldsDuringRead(existingState ListAccountGroupsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountGroupsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListAccountGroupsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListAccountGroupsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListAccountGroupsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListAccountGroupsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListAccountGroupsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListAccountGroupsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListAccountGroupsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountGroupsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// List service principals
type ListAccountServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListAccountServicePrincipalsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountServicePrincipalsRequest) {
}

func (newState *ListAccountServicePrincipalsRequest) SyncEffectiveFieldsDuringRead(existingState ListAccountServicePrincipalsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountServicePrincipalsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountServicePrincipalsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListAccountServicePrincipalsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListAccountServicePrincipalsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListAccountServicePrincipalsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListAccountServicePrincipalsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListAccountServicePrincipalsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListAccountServicePrincipalsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListAccountServicePrincipalsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountServicePrincipalsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// List users
type ListAccountUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListAccountUsersRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountUsersRequest) {
}

func (newState *ListAccountUsersRequest) SyncEffectiveFieldsDuringRead(existingState ListAccountUsersRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountUsersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountUsersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListAccountUsersRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListAccountUsersRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListAccountUsersRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListAccountUsersRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListAccountUsersRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListAccountUsersRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListAccountUsersRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountUsersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// List group details
type ListGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListGroupsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListGroupsRequest) {
}

func (newState *ListGroupsRequest) SyncEffectiveFieldsDuringRead(existingState ListGroupsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListGroupsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListGroupsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListGroupsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListGroupsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListGroupsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListGroupsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListGroupsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListGroupsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type ListGroupsResponse struct {
	// Total results returned in the response.
	ItemsPerPage types.Int64 `tfsdk:"itemsPerPage" tf:"optional"`
	// User objects returned in the response.
	Resources types.List `tfsdk:"Resources" tf:"optional"`
	// The schema of the service principal.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex types.Int64 `tfsdk:"startIndex" tf:"optional"`
	// Total results that match the request filters.
	TotalResults types.Int64 `tfsdk:"totalResults" tf:"optional"`
}

func (newState *ListGroupsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListGroupsResponse) {
}

func (newState *ListGroupsResponse) SyncEffectiveFieldsDuringRead(existingState ListGroupsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListGroupsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(Group{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListGroupsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListGroupsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListGroupsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListGroupsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListGroupsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListGroupsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListGroupsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListGroupsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: Group{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

type ListServicePrincipalResponse struct {
	// Total results returned in the response.
	ItemsPerPage types.Int64 `tfsdk:"itemsPerPage" tf:"optional"`
	// User objects returned in the response.
	Resources types.List `tfsdk:"Resources" tf:"optional"`
	// The schema of the List response.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex types.Int64 `tfsdk:"startIndex" tf:"optional"`
	// Total results that match the request filters.
	TotalResults types.Int64 `tfsdk:"totalResults" tf:"optional"`
}

func (newState *ListServicePrincipalResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListServicePrincipalResponse) {
}

func (newState *ListServicePrincipalResponse) SyncEffectiveFieldsDuringRead(existingState ListServicePrincipalResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListServicePrincipalResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(ServicePrincipal{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListServicePrincipalResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListServicePrincipalResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListServicePrincipalResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListServicePrincipalResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListServicePrincipalResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListServicePrincipalResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListServicePrincipalResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListServicePrincipalResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: ServicePrincipal{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

// List service principals
type ListServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListServicePrincipalsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListServicePrincipalsRequest) {
}

func (newState *ListServicePrincipalsRequest) SyncEffectiveFieldsDuringRead(existingState ListServicePrincipalsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListServicePrincipalsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListServicePrincipalsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListServicePrincipalsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListServicePrincipalsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListServicePrincipalsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListServicePrincipalsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListServicePrincipalsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListServicePrincipalsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListServicePrincipalsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// List users
type ListUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListUsersRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListUsersRequest) {
}

func (newState *ListUsersRequest) SyncEffectiveFieldsDuringRead(existingState ListUsersRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListUsersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListUsersRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListUsersRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListUsersRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListUsersRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListUsersRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListUsersRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListUsersRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListUsersRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type ListUsersResponse struct {
	// Total results returned in the response.
	ItemsPerPage types.Int64 `tfsdk:"itemsPerPage" tf:"optional"`
	// User objects returned in the response.
	Resources types.List `tfsdk:"Resources" tf:"optional"`
	// The schema of the List response.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex types.Int64 `tfsdk:"startIndex" tf:"optional"`
	// Total results that match the request filters.
	TotalResults types.Int64 `tfsdk:"totalResults" tf:"optional"`
}

func (newState *ListUsersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListUsersResponse) {
}

func (newState *ListUsersResponse) SyncEffectiveFieldsDuringRead(existingState ListUsersResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListUsersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(User{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListUsersResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListUsersResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListUsersResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListUsersResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListUsersResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListUsersResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListUsersResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListUsersResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: User{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

// Get permission assignments
type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *ListWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListWorkspaceAssignmentRequest) {
}

func (newState *ListWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringRead(existingState ListWorkspaceAssignmentRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListWorkspaceAssignmentRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListWorkspaceAssignmentRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListWorkspaceAssignmentRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListWorkspaceAssignmentRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListWorkspaceAssignmentRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListWorkspaceAssignmentRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListWorkspaceAssignmentRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type MigratePermissionsRequest struct {
	// The name of the workspace group that permissions will be migrated from.
	FromWorkspaceGroupName types.String `tfsdk:"from_workspace_group_name" tf:""`
	// The maximum number of permissions that will be migrated.
	Size types.Int64 `tfsdk:"size" tf:"optional"`
	// The name of the account group that permissions will be migrated to.
	ToAccountGroupName types.String `tfsdk:"to_account_group_name" tf:""`
	// WorkspaceId of the associated workspace where the permission migration
	// will occur.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:""`
}

func (newState *MigratePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan MigratePermissionsRequest) {
}

func (newState *MigratePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState MigratePermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigratePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MigratePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = MigratePermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o MigratePermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o MigratePermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o MigratePermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o MigratePermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o MigratePermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o MigratePermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o MigratePermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"from_workspace_group_name": types.StringType,
			"size":                      types.Int64Type,
			"to_account_group_name":     types.StringType,
			"workspace_id":              types.Int64Type,
		},
	}
}

type MigratePermissionsResponse struct {
	// Number of permissions migrated.
	PermissionsMigrated types.Int64 `tfsdk:"permissions_migrated" tf:"optional"`
}

func (newState *MigratePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MigratePermissionsResponse) {
}

func (newState *MigratePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState MigratePermissionsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigratePermissionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MigratePermissionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = MigratePermissionsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o MigratePermissionsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o MigratePermissionsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o MigratePermissionsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o MigratePermissionsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o MigratePermissionsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o MigratePermissionsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o MigratePermissionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions_migrated": types.Int64Type,
		},
	}
}

type Name struct {
	// Family name of the Databricks user.
	FamilyName types.String `tfsdk:"familyName" tf:"optional"`
	// Given name of the Databricks user.
	GivenName types.String `tfsdk:"givenName" tf:"optional"`
}

func (newState *Name) SyncEffectiveFieldsDuringCreateOrUpdate(plan Name) {
}

func (newState *Name) SyncEffectiveFieldsDuringRead(existingState Name) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Name.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Name) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Name{}

// Equal implements basetypes.ObjectValuable.
func (o Name) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Name) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Name) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Name) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Name) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Name) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Name) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"familyName": types.StringType,
			"givenName":  types.StringType,
		},
	}
}

type ObjectPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *ObjectPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ObjectPermissions) {
}

func (newState *ObjectPermissions) SyncEffectiveFieldsDuringRead(existingState ObjectPermissions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ObjectPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ObjectPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlResponse{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ObjectPermissions{}

// Equal implements basetypes.ObjectValuable.
func (o ObjectPermissions) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ObjectPermissions) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ObjectPermissions) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ObjectPermissions) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ObjectPermissions) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ObjectPermissions) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ObjectPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type PartialUpdate struct {
	// Unique ID for a user in the Databricks workspace.
	Id types.String `tfsdk:"-"`

	Operations types.List `tfsdk:"Operations" tf:"optional"`
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
}

func (newState *PartialUpdate) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartialUpdate) {
}

func (newState *PartialUpdate) SyncEffectiveFieldsDuringRead(existingState PartialUpdate) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartialUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PartialUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Operations": reflect.TypeOf(Patch{}),
		"schemas":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PartialUpdate{}

// Equal implements basetypes.ObjectValuable.
func (o PartialUpdate) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PartialUpdate) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PartialUpdate) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PartialUpdate) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PartialUpdate) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PartialUpdate) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PartialUpdate) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"Operations": basetypes.ListType{
				ElemType: Patch{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type PasswordAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *PasswordAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordAccessControlRequest) {
}

func (newState *PasswordAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState PasswordAccessControlRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PasswordAccessControlRequest{}

// Equal implements basetypes.ObjectValuable.
func (o PasswordAccessControlRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PasswordAccessControlRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PasswordAccessControlRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PasswordAccessControlRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PasswordAccessControlRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PasswordAccessControlRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PasswordAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type PasswordAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *PasswordAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordAccessControlResponse) {
}

func (newState *PasswordAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState PasswordAccessControlResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(PasswordPermission{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PasswordAccessControlResponse{}

// Equal implements basetypes.ObjectValuable.
func (o PasswordAccessControlResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PasswordAccessControlResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PasswordAccessControlResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PasswordAccessControlResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PasswordAccessControlResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PasswordAccessControlResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PasswordAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: PasswordPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type PasswordPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PasswordPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermission) {
}

func (newState *PasswordPermission) SyncEffectiveFieldsDuringRead(existingState PasswordPermission) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PasswordPermission{}

// Equal implements basetypes.ObjectValuable.
func (o PasswordPermission) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PasswordPermission) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PasswordPermission) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PasswordPermission) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PasswordPermission) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PasswordPermission) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermission) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

type PasswordPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *PasswordPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermissions) {
}

func (newState *PasswordPermissions) SyncEffectiveFieldsDuringRead(existingState PasswordPermissions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PasswordAccessControlResponse{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PasswordPermissions{}

// Equal implements basetypes.ObjectValuable.
func (o PasswordPermissions) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PasswordPermissions) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PasswordPermissions) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PasswordPermissions) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PasswordPermissions) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PasswordPermissions) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PasswordAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type PasswordPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PasswordPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermissionsDescription) {
}

func (newState *PasswordPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState PasswordPermissionsDescription) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PasswordPermissionsDescription{}

// Equal implements basetypes.ObjectValuable.
func (o PasswordPermissionsDescription) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PasswordPermissionsDescription) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PasswordPermissionsDescription) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PasswordPermissionsDescription) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PasswordPermissionsDescription) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PasswordPermissionsDescription) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PasswordPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
}

func (newState *PasswordPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermissionsRequest) {
}

func (newState *PasswordPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState PasswordPermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PasswordAccessControlRequest{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PasswordPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o PasswordPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PasswordPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PasswordPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PasswordPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PasswordPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PasswordPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PasswordPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PasswordAccessControlRequest{}.Type(ctx),
			},
		},
	}
}

type Patch struct {
	// Type of patch operation.
	Op types.String `tfsdk:"op" tf:"optional"`
	// Selection of patch operation
	Path types.String `tfsdk:"path" tf:"optional"`
	// Value to modify
	Value any `tfsdk:"value" tf:"optional"`
}

func (newState *Patch) SyncEffectiveFieldsDuringCreateOrUpdate(plan Patch) {
}

func (newState *Patch) SyncEffectiveFieldsDuringRead(existingState Patch) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Patch.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Patch) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Patch{}

// Equal implements basetypes.ObjectValuable.
func (o Patch) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Patch) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Patch) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Patch) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Patch) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Patch) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Patch) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"op":    types.StringType,
			"path":  types.StringType,
			"value": types.ObjectType{},
		},
	}
}

type PatchResponse struct {
}

func (newState *PatchResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PatchResponse) {
}

func (newState *PatchResponse) SyncEffectiveFieldsDuringRead(existingState PatchResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PatchResponse{}

// Equal implements basetypes.ObjectValuable.
func (o PatchResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PatchResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PatchResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PatchResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PatchResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PatchResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PatchResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Permission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *Permission) SyncEffectiveFieldsDuringCreateOrUpdate(plan Permission) {
}

func (newState *Permission) SyncEffectiveFieldsDuringRead(existingState Permission) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Permission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Permission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Permission{}

// Equal implements basetypes.ObjectValuable.
func (o Permission) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Permission) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Permission) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Permission) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Permission) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Permission) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Permission) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

// The output format for existing workspace PermissionAssignment records, which
// contains some info for user consumption.
type PermissionAssignment struct {
	// Error response associated with a workspace permission assignment, if any.
	Error types.String `tfsdk:"error" tf:"optional"`
	// The permissions level of the principal.
	Permissions types.List `tfsdk:"permissions" tf:"optional"`
	// Information about the principal assigned to the workspace.
	Principal types.List `tfsdk:"principal" tf:"optional,object"`
}

func (newState *PermissionAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionAssignment) {
}

func (newState *PermissionAssignment) SyncEffectiveFieldsDuringRead(existingState PermissionAssignment) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
		"principal":   reflect.TypeOf(PrincipalOutput{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PermissionAssignment{}

// Equal implements basetypes.ObjectValuable.
func (o PermissionAssignment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PermissionAssignment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PermissionAssignment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PermissionAssignment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PermissionAssignment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PermissionAssignment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PermissionAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": types.StringType,
			"permissions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal": basetypes.ListType{
				ElemType: PrincipalOutput{}.Type(ctx),
			},
		},
	}
}

type PermissionAssignments struct {
	// Array of permissions assignments defined for a workspace.
	PermissionAssignments types.List `tfsdk:"permission_assignments" tf:"optional"`
}

func (newState *PermissionAssignments) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionAssignments) {
}

func (newState *PermissionAssignments) SyncEffectiveFieldsDuringRead(existingState PermissionAssignments) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionAssignments.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionAssignments) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_assignments": reflect.TypeOf(PermissionAssignment{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PermissionAssignments{}

// Equal implements basetypes.ObjectValuable.
func (o PermissionAssignments) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PermissionAssignments) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PermissionAssignments) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PermissionAssignments) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PermissionAssignments) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PermissionAssignments) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PermissionAssignments) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_assignments": basetypes.ListType{
				ElemType: PermissionAssignment{}.Type(ctx),
			},
		},
	}
}

type PermissionOutput struct {
	// The results of a permissions query.
	Description types.String `tfsdk:"description" tf:"optional"`

	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PermissionOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionOutput) {
}

func (newState *PermissionOutput) SyncEffectiveFieldsDuringRead(existingState PermissionOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PermissionOutput{}

// Equal implements basetypes.ObjectValuable.
func (o PermissionOutput) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PermissionOutput) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PermissionOutput) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PermissionOutput) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PermissionOutput) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PermissionOutput) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PermissionOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsDescription) {
}

func (newState *PermissionsDescription) SyncEffectiveFieldsDuringRead(existingState PermissionsDescription) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PermissionsDescription{}

// Equal implements basetypes.ObjectValuable.
func (o PermissionsDescription) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PermissionsDescription) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PermissionsDescription) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PermissionsDescription) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PermissionsDescription) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PermissionsDescription) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The id of the request object.
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

func (newState *PermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsRequest) {
}

func (newState *PermissionsRequest) SyncEffectiveFieldsDuringRead(existingState PermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlRequest{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o PermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlRequest{}.Type(ctx),
			},
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

// Information about the principal assigned to the workspace.
type PrincipalOutput struct {
	// The display name of the principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The group name of the group. Present only if the principal is a group.
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// The unique, opaque id of the principal.
	PrincipalId types.Int64 `tfsdk:"principal_id" tf:"optional"`
	// The name of the service principal. Present only if the principal is a
	// service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// The username of the user. Present only if the principal is a user.
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *PrincipalOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrincipalOutput) {
}

func (newState *PrincipalOutput) SyncEffectiveFieldsDuringRead(existingState PrincipalOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrincipalOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PrincipalOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PrincipalOutput{}

// Equal implements basetypes.ObjectValuable.
func (o PrincipalOutput) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PrincipalOutput) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PrincipalOutput) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PrincipalOutput) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PrincipalOutput) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PrincipalOutput) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o PrincipalOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"principal_id":           types.Int64Type,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ResourceMeta struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	ResourceType types.String `tfsdk:"resourceType" tf:"optional"`
}

func (newState *ResourceMeta) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResourceMeta) {
}

func (newState *ResourceMeta) SyncEffectiveFieldsDuringRead(existingState ResourceMeta) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResourceMeta.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResourceMeta) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ResourceMeta{}

// Equal implements basetypes.ObjectValuable.
func (o ResourceMeta) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ResourceMeta) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ResourceMeta) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ResourceMeta) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ResourceMeta) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ResourceMeta) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ResourceMeta) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resourceType": types.StringType,
		},
	}
}

type Role struct {
	// Role to assign to a principal or a list of principals on a resource.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *Role) SyncEffectiveFieldsDuringCreateOrUpdate(plan Role) {
}

func (newState *Role) SyncEffectiveFieldsDuringRead(existingState Role) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Role.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Role) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Role{}

// Equal implements basetypes.ObjectValuable.
func (o Role) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Role) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Role) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Role) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Role) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Role) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Role) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type RuleSetResponse struct {
	// Identifies the version of the rule set returned.
	Etag types.String `tfsdk:"etag" tf:"optional"`

	GrantRules types.List `tfsdk:"grant_rules" tf:"optional"`
	// Name of the rule set.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *RuleSetResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RuleSetResponse) {
}

func (newState *RuleSetResponse) SyncEffectiveFieldsDuringRead(existingState RuleSetResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleSetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RuleSetResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(GrantRule{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RuleSetResponse{}

// Equal implements basetypes.ObjectValuable.
func (o RuleSetResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RuleSetResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RuleSetResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RuleSetResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RuleSetResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RuleSetResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o RuleSetResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"grant_rules": basetypes.ListType{
				ElemType: GrantRule{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

type RuleSetUpdateRequest struct {
	// The expected etag of the rule set to update. The update will fail if the
	// value does not match the value that is stored in account access control
	// service.
	Etag types.String `tfsdk:"etag" tf:""`

	GrantRules types.List `tfsdk:"grant_rules" tf:"optional"`
	// Name of the rule set.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *RuleSetUpdateRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RuleSetUpdateRequest) {
}

func (newState *RuleSetUpdateRequest) SyncEffectiveFieldsDuringRead(existingState RuleSetUpdateRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleSetUpdateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RuleSetUpdateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(GrantRule{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RuleSetUpdateRequest{}

// Equal implements basetypes.ObjectValuable.
func (o RuleSetUpdateRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RuleSetUpdateRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RuleSetUpdateRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RuleSetUpdateRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RuleSetUpdateRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RuleSetUpdateRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o RuleSetUpdateRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"grant_rules": basetypes.ListType{
				ElemType: GrantRule{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

type ServicePrincipal struct {
	// If this user is active
	Active types.Bool `tfsdk:"active" tf:"optional"`
	// UUID relating to the service principal
	ApplicationId types.String `tfsdk:"applicationId" tf:"optional"`
	// String that represents a concatenation of given and family names.
	DisplayName types.String `tfsdk:"displayName" tf:"optional"`
	// Entitlements assigned to the service principal. See [assigning
	// entitlements] for a full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements types.List `tfsdk:"entitlements" tf:"optional"`

	ExternalId types.String `tfsdk:"externalId" tf:"optional"`

	Groups types.List `tfsdk:"groups" tf:"optional"`
	// Databricks service principal ID.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles" tf:"optional"`
	// The schema of the List response.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
}

func (newState *ServicePrincipal) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServicePrincipal) {
}

func (newState *ServicePrincipal) SyncEffectiveFieldsDuringRead(existingState ServicePrincipal) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServicePrincipal.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServicePrincipal) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ServicePrincipal{}

// Equal implements basetypes.ObjectValuable.
func (o ServicePrincipal) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ServicePrincipal) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ServicePrincipal) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ServicePrincipal) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ServicePrincipal) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ServicePrincipal) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ServicePrincipal) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active":        types.BoolType,
			"applicationId": types.StringType,
			"displayName":   types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"id": types.StringType,
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateRuleSetRequest struct {
	// Name of the rule set.
	Name types.String `tfsdk:"name" tf:""`

	RuleSet types.List `tfsdk:"rule_set" tf:"object"`
}

func (newState *UpdateRuleSetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRuleSetRequest) {
}

func (newState *UpdateRuleSetRequest) SyncEffectiveFieldsDuringRead(existingState UpdateRuleSetRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRuleSetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRuleSetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rule_set": reflect.TypeOf(RuleSetUpdateRequest{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateRuleSetRequest{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateRuleSetRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateRuleSetRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateRuleSetRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateRuleSetRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateRuleSetRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateRuleSetRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRuleSetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"rule_set": basetypes.ListType{
				ElemType: RuleSetUpdateRequest{}.Type(ctx),
			},
		},
	}
}

type UpdateWorkspaceAssignments struct {
	// Array of permissions assignments to update on the workspace. Valid values
	// are "USER" and "ADMIN" (case-sensitive). If both "USER" and "ADMIN" are
	// provided, "ADMIN" takes precedence. Other values will be ignored. Note
	// that excluding this field, or providing unsupported values, will have the
	// same effect as providing an empty list, which will result in the deletion
	// of all permissions for the principal.
	Permissions types.List `tfsdk:"permissions" tf:"optional"`
	// The ID of the user, service principal, or group.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UpdateWorkspaceAssignments) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceAssignments) {
}

func (newState *UpdateWorkspaceAssignments) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceAssignments) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceAssignments.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceAssignments) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateWorkspaceAssignments{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateWorkspaceAssignments) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateWorkspaceAssignments) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateWorkspaceAssignments) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateWorkspaceAssignments) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateWorkspaceAssignments) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateWorkspaceAssignments) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceAssignments) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type User struct {
	// If this user is active
	Active types.Bool `tfsdk:"active" tf:"optional"`
	// String that represents a concatenation of given and family names. For
	// example `John Smith`. This field cannot be updated through the Workspace
	// SCIM APIs when [identity federation is enabled]. Use Account SCIM APIs to
	// update `displayName`.
	//
	// [identity federation is enabled]: https://docs.databricks.com/administration-guide/users-groups/best-practices.html#enable-identity-federation
	DisplayName types.String `tfsdk:"displayName" tf:"optional"`
	// All the emails associated with the Databricks user.
	Emails types.List `tfsdk:"emails" tf:"optional"`
	// Entitlements assigned to the user. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements types.List `tfsdk:"entitlements" tf:"optional"`
	// External ID is not currently supported. It is reserved for future use.
	ExternalId types.String `tfsdk:"externalId" tf:"optional"`

	Groups types.List `tfsdk:"groups" tf:"optional"`
	// Databricks user ID. This is automatically set by Databricks. Any value
	// provided by the client will be ignored.
	Id types.String `tfsdk:"id" tf:"optional"`

	Name types.List `tfsdk:"name" tf:"optional,object"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles" tf:"optional"`
	// The schema of the user.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
	// Email address of the Databricks user.
	UserName types.String `tfsdk:"userName" tf:"optional"`
}

func (newState *User) SyncEffectiveFieldsDuringCreateOrUpdate(plan User) {
}

func (newState *User) SyncEffectiveFieldsDuringRead(existingState User) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in User.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a User) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"emails":       reflect.TypeOf(ComplexValue{}),
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"name":         reflect.TypeOf(Name{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = User{}

// Equal implements basetypes.ObjectValuable.
func (o User) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o User) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o User) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o User) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o User) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o User) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o User) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active":      types.BoolType,
			"displayName": types.StringType,
			"emails": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"id": types.StringType,
			"name": basetypes.ListType{
				ElemType: Name{}.Type(ctx),
			},
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.Type(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"userName": types.StringType,
		},
	}
}

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	Permissions types.List `tfsdk:"permissions" tf:"optional"`
}

func (newState *WorkspacePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspacePermissions) {
}

func (newState *WorkspacePermissions) SyncEffectiveFieldsDuringRead(existingState WorkspacePermissions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspacePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspacePermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(PermissionOutput{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = WorkspacePermissions{}

// Equal implements basetypes.ObjectValuable.
func (o WorkspacePermissions) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o WorkspacePermissions) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o WorkspacePermissions) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o WorkspacePermissions) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o WorkspacePermissions) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o WorkspacePermissions) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o WorkspacePermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions": basetypes.ListType{
				ElemType: PermissionOutput{}.Type(ctx),
			},
		},
	}
}

// Permission level

// Type of patch operation.

// Permission level
