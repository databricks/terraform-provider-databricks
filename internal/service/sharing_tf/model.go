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
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type CreateProvider struct {
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

func (newState *CreateProvider) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateProvider) {
}

func (newState *CreateProvider) SyncEffectiveFieldsDuringRead(existingState CreateProvider) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateProvider{}

// Equal implements basetypes.ObjectValuable.
func (o CreateProvider) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateProvider) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateProvider) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateProvider) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateProvider) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateProvider) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *CreateRecipient) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRecipient) {
}

func (newState *CreateRecipient) SyncEffectiveFieldsDuringRead(existingState CreateRecipient) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateRecipient{}

// Equal implements basetypes.ObjectValuable.
func (o CreateRecipient) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateRecipient) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateRecipient) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateRecipient) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateRecipient) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateRecipient) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateRecipient) Type(ctx context.Context) attr.Type {
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

type CreateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the share.
	Name types.String `tfsdk:"name" tf:""`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
}

func (newState *CreateShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateShare) {
}

func (newState *CreateShare) SyncEffectiveFieldsDuringRead(existingState CreateShare) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateShare{}

// Equal implements basetypes.ObjectValuable.
func (o CreateShare) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateShare) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateShare) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateShare) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateShare) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateShare) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *DeleteProviderRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteProviderRequest) {
}

func (newState *DeleteProviderRequest) SyncEffectiveFieldsDuringRead(existingState DeleteProviderRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteProviderRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteProviderRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteProviderRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteProviderRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteProviderRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteProviderRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteProviderRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *DeleteRecipientRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRecipientRequest) {
}

func (newState *DeleteRecipientRequest) SyncEffectiveFieldsDuringRead(existingState DeleteRecipientRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteRecipientRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteRecipientRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteRecipientRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteRecipientRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteRecipientRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteRecipientRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteRecipientRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Delete a share
type DeleteShareRequest struct {
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteShareRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteShareRequest) {
}

func (newState *DeleteShareRequest) SyncEffectiveFieldsDuringRead(existingState DeleteShareRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteShareRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteShareRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteShareRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteShareRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteShareRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteShareRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteShareRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *GetActivationUrlInfoRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetActivationUrlInfoRequest) {
}

func (newState *GetActivationUrlInfoRequest) SyncEffectiveFieldsDuringRead(existingState GetActivationUrlInfoRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetActivationUrlInfoRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetActivationUrlInfoResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetActivationUrlInfoResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *GetProviderRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetProviderRequest) {
}

func (newState *GetProviderRequest) SyncEffectiveFieldsDuringRead(existingState GetProviderRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetProviderRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetProviderRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetProviderRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetProviderRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetProviderRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetProviderRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetProviderRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *GetRecipientRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRecipientRequest) {
}

func (newState *GetRecipientRequest) SyncEffectiveFieldsDuringRead(existingState GetRecipientRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetRecipientRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetRecipientRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetRecipientRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetRecipientRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetRecipientRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetRecipientRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetRecipientRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of data share permissions for a recipient.
	PermissionsOut types.List `tfsdk:"permissions_out" tf:"optional"`
}

func (newState *GetRecipientSharePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRecipientSharePermissionsResponse) {
}

func (newState *GetRecipientSharePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState GetRecipientSharePermissionsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetRecipientSharePermissionsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetRecipientSharePermissionsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetRecipientSharePermissionsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetRecipientSharePermissionsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetRecipientSharePermissionsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetRecipientSharePermissionsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetRecipientSharePermissionsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get a share
type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData types.Bool `tfsdk:"-"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
}

func (newState *GetShareRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetShareRequest) {
}

func (newState *GetShareRequest) SyncEffectiveFieldsDuringRead(existingState GetShareRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetShareRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetShareRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetShareRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetShareRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetShareRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetShareRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetShareRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AllowedIpAddresses types.List `tfsdk:"allowed_ip_addresses" tf:"optional"`
}

func (newState *IpAccessList) SyncEffectiveFieldsDuringCreateOrUpdate(plan IpAccessList) {
}

func (newState *IpAccessList) SyncEffectiveFieldsDuringRead(existingState IpAccessList) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = IpAccessList{}

// Equal implements basetypes.ObjectValuable.
func (o IpAccessList) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o IpAccessList) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o IpAccessList) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o IpAccessList) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o IpAccessList) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o IpAccessList) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ListProviderSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of provider shares.
	Shares types.List `tfsdk:"shares" tf:"optional"`
}

func (newState *ListProviderSharesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProviderSharesResponse) {
}

func (newState *ListProviderSharesResponse) SyncEffectiveFieldsDuringRead(existingState ListProviderSharesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListProviderSharesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListProviderSharesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListProviderSharesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListProviderSharesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListProviderSharesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListProviderSharesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListProviderSharesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *ListProvidersRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersRequest) {
}

func (newState *ListProvidersRequest) SyncEffectiveFieldsDuringRead(existingState ListProvidersRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListProvidersRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListProvidersRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListProvidersRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListProvidersRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListProvidersRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListProvidersRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListProvidersRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of provider information objects.
	Providers types.List `tfsdk:"providers" tf:"optional"`
}

func (newState *ListProvidersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListProvidersResponse) {
}

func (newState *ListProvidersResponse) SyncEffectiveFieldsDuringRead(existingState ListProvidersResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListProvidersResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListProvidersResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListProvidersResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListProvidersResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListProvidersResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListProvidersResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListProvidersResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *ListRecipientsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRecipientsRequest) {
}

func (newState *ListRecipientsRequest) SyncEffectiveFieldsDuringRead(existingState ListRecipientsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListRecipientsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListRecipientsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListRecipientsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListRecipientsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListRecipientsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListRecipientsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListRecipientsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of recipient information objects.
	Recipients types.List `tfsdk:"recipients" tf:"optional"`
}

func (newState *ListRecipientsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRecipientsResponse) {
}

func (newState *ListRecipientsResponse) SyncEffectiveFieldsDuringRead(existingState ListRecipientsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListRecipientsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListRecipientsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListRecipientsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListRecipientsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListRecipientsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListRecipientsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListRecipientsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *ListSharesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSharesRequest) {
}

func (newState *ListSharesRequest) SyncEffectiveFieldsDuringRead(existingState ListSharesRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListSharesRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListSharesRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListSharesRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListSharesRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListSharesRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListSharesRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListSharesRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of data share information objects.
	Shares types.List `tfsdk:"shares" tf:"optional"`
}

func (newState *ListSharesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSharesResponse) {
}

func (newState *ListSharesResponse) SyncEffectiveFieldsDuringRead(existingState ListSharesResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListSharesResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListSharesResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListSharesResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListSharesResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListSharesResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListSharesResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListSharesResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type Partition struct {
	// An array of partition values.
	Values types.List `tfsdk:"value" tf:"optional"`
}

func (newState *Partition) SyncEffectiveFieldsDuringCreateOrUpdate(plan Partition) {
}

func (newState *Partition) SyncEffectiveFieldsDuringRead(existingState Partition) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Partition{}

// Equal implements basetypes.ObjectValuable.
func (o Partition) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Partition) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Partition) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Partition) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Partition) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Partition) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type PartitionValue struct {
	// The name of the partition column.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The operator to apply for the value.
	Op types.String `tfsdk:"op" tf:"optional"`
	// The key of a Delta Sharing recipient's property. For example
	// `databricks-account-id`. When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey types.String `tfsdk:"recipient_property_key" tf:"optional"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *PartitionValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartitionValue) {
}

func (newState *PartitionValue) SyncEffectiveFieldsDuringRead(existingState PartitionValue) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PartitionValue{}

// Equal implements basetypes.ObjectValuable.
func (o PartitionValue) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PartitionValue) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PartitionValue) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PartitionValue) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PartitionValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PartitionValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The privileges assigned to the principal.
	Privileges types.List `tfsdk:"privileges" tf:"optional"`
}

func (newState *PrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrivilegeAssignment) {
}

func (newState *PrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState PrivilegeAssignment) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = PrivilegeAssignment{}

// Equal implements basetypes.ObjectValuable.
func (o PrivilegeAssignment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o PrivilegeAssignment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o PrivilegeAssignment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o PrivilegeAssignment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o PrivilegeAssignment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o PrivilegeAssignment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ProviderInfo struct {
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

func (newState *ProviderInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderInfo) {
}

func (newState *ProviderInfo) SyncEffectiveFieldsDuringRead(existingState ProviderInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ProviderInfo{}

// Equal implements basetypes.ObjectValuable.
func (o ProviderInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ProviderInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ProviderInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ProviderInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ProviderInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ProviderInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ProviderShare struct {
	// The name of the Provider Share.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ProviderShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProviderShare) {
}

func (newState *ProviderShare) SyncEffectiveFieldsDuringRead(existingState ProviderShare) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ProviderShare{}

// Equal implements basetypes.ObjectValuable.
func (o ProviderShare) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ProviderShare) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ProviderShare) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ProviderShare) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ProviderShare) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ProviderShare) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *RecipientInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientInfo) {
}

func (newState *RecipientInfo) SyncEffectiveFieldsDuringRead(existingState RecipientInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RecipientInfo{}

// Equal implements basetypes.ObjectValuable.
func (o RecipientInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RecipientInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RecipientInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RecipientInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RecipientInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RecipientInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type RecipientProfile struct {
	// The token used to authorize the recipient.
	BearerToken types.String `tfsdk:"bearer_token" tf:"optional"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion types.Int64 `tfsdk:"share_credentials_version" tf:"optional"`
}

func (newState *RecipientProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientProfile) {
}

func (newState *RecipientProfile) SyncEffectiveFieldsDuringRead(existingState RecipientProfile) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RecipientProfile{}

// Equal implements basetypes.ObjectValuable.
func (o RecipientProfile) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RecipientProfile) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RecipientProfile) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RecipientProfile) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RecipientProfile) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RecipientProfile) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *RecipientTokenInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RecipientTokenInfo) {
}

func (newState *RecipientTokenInfo) SyncEffectiveFieldsDuringRead(existingState RecipientTokenInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RecipientTokenInfo{}

// Equal implements basetypes.ObjectValuable.
func (o RecipientTokenInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RecipientTokenInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RecipientTokenInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RecipientTokenInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RecipientTokenInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RecipientTokenInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *RetrieveTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RetrieveTokenRequest) {
}

func (newState *RetrieveTokenRequest) SyncEffectiveFieldsDuringRead(existingState RetrieveTokenRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RetrieveTokenRequest{}

// Equal implements basetypes.ObjectValuable.
func (o RetrieveTokenRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RetrieveTokenRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RetrieveTokenRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RetrieveTokenRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RetrieveTokenRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RetrieveTokenRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	BearerToken types.String `tfsdk:"bearerToken" tf:"optional"`
	// The endpoint for the share to be used by the recipient.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime types.String `tfsdk:"expirationTime" tf:"optional"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion types.Int64 `tfsdk:"shareCredentialsVersion" tf:"optional"`
}

func (newState *RetrieveTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RetrieveTokenResponse) {
}

func (newState *RetrieveTokenResponse) SyncEffectiveFieldsDuringRead(existingState RetrieveTokenResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RetrieveTokenResponse{}

// Equal implements basetypes.ObjectValuable.
func (o RetrieveTokenResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RetrieveTokenResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RetrieveTokenResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RetrieveTokenResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RetrieveTokenResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RetrieveTokenResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ExistingTokenExpireInSeconds types.Int64 `tfsdk:"existing_token_expire_in_seconds" tf:""`
	// The name of the recipient.
	Name types.String `tfsdk:"-"`
}

func (newState *RotateRecipientToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan RotateRecipientToken) {
}

func (newState *RotateRecipientToken) SyncEffectiveFieldsDuringRead(existingState RotateRecipientToken) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RotateRecipientToken{}

// Equal implements basetypes.ObjectValuable.
func (o RotateRecipientToken) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RotateRecipientToken) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RotateRecipientToken) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RotateRecipientToken) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RotateRecipientToken) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RotateRecipientToken) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Properties types.Map `tfsdk:"properties" tf:""`
}

func (newState *SecurablePropertiesKvPairs) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecurablePropertiesKvPairs) {
}

func (newState *SecurablePropertiesKvPairs) SyncEffectiveFieldsDuringRead(existingState SecurablePropertiesKvPairs) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SecurablePropertiesKvPairs{}

// Equal implements basetypes.ObjectValuable.
func (o SecurablePropertiesKvPairs) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SecurablePropertiesKvPairs) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SecurablePropertiesKvPairs) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SecurablePropertiesKvPairs) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SecurablePropertiesKvPairs) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SecurablePropertiesKvPairs) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ShareInfo struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"computed,optional"`
	// Username of share creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"computed,optional"`
	// Name of the share.
	Name types.String `tfsdk:"name" tf:"optional"`
	// A list of shared data objects within the share.
	Objects types.List `tfsdk:"object" tf:"optional"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner" tf:"computed,optional"`
	// Storage Location URL (full path) for the share.
	StorageLocation types.String `tfsdk:"storage_location" tf:"optional"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"computed,optional"`
	// Username of share updater.
	UpdatedBy types.String `tfsdk:"updated_by" tf:"computed,optional"`
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareInfo) {
}

func (newState *ShareInfo) SyncEffectiveFieldsDuringRead(existingState ShareInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ShareInfo{}

// Equal implements basetypes.ObjectValuable.
func (o ShareInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ShareInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ShareInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ShareInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ShareInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ShareInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *SharePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharePermissionsRequest) {
}

func (newState *SharePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState SharePermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SharePermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o SharePermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SharePermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SharePermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SharePermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SharePermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SharePermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments" tf:"optional"`
	// The share name.
	ShareName types.String `tfsdk:"share_name" tf:"optional"`
}

func (newState *ShareToPrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan ShareToPrivilegeAssignment) {
}

func (newState *ShareToPrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState ShareToPrivilegeAssignment) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ShareToPrivilegeAssignment{}

// Equal implements basetypes.ObjectValuable.
func (o ShareToPrivilegeAssignment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ShareToPrivilegeAssignment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ShareToPrivilegeAssignment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ShareToPrivilegeAssignment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ShareToPrivilegeAssignment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ShareToPrivilegeAssignment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at" tf:"computed,optional"`
	// Username of the sharer.
	AddedBy types.String `tfsdk:"added_by" tf:"computed,optional"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled types.Bool `tfsdk:"cdf_enabled" tf:"computed,optional"`
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
	HistoryDataSharingStatus types.String `tfsdk:"history_data_sharing_status" tf:"computed,optional"`
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
	SharedAs types.String `tfsdk:"shared_as" tf:"computed,optional"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion types.Int64 `tfsdk:"start_version" tf:"computed,optional"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status types.String `tfsdk:"status" tf:"computed,optional"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `string_shared_as` name. The `string_shared_as` name must be unique
	// within a share. For notebooks, the new name should be the new notebook
	// file name.
	StringSharedAs types.String `tfsdk:"string_shared_as" tf:"optional"`
}

func (newState *SharedDataObject) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObject) {
}

func (newState *SharedDataObject) SyncEffectiveFieldsDuringRead(existingState SharedDataObject) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SharedDataObject{}

// Equal implements basetypes.ObjectValuable.
func (o SharedDataObject) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SharedDataObject) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SharedDataObject) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SharedDataObject) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SharedDataObject) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SharedDataObject) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o SharedDataObject) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"added_at":                    types.Int64Type,
			"added_by":                    types.StringType,
			"cdf_enabled":                 types.BoolType,
			"comment":                     types.StringType,
			"content":                     types.StringType,
			"data_object_type":            types.StringType,
			"history_data_sharing_status": types.StringType,
			"name":                        types.StringType,
			"partition": basetypes.ListType{
				ElemType: Partition{}.Type(ctx),
			},
			"shared_as":        types.StringType,
			"start_version":    types.Int64Type,
			"status":           types.StringType,
			"string_shared_as": types.StringType,
		},
	}
}

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	Action types.String `tfsdk:"action" tf:"optional"`
	// The data object that is being added, removed, or updated.
	DataObject types.List `tfsdk:"data_object" tf:"optional,object"`
}

func (newState *SharedDataObjectUpdate) SyncEffectiveFieldsDuringCreateOrUpdate(plan SharedDataObjectUpdate) {
}

func (newState *SharedDataObjectUpdate) SyncEffectiveFieldsDuringRead(existingState SharedDataObjectUpdate) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SharedDataObjectUpdate{}

// Equal implements basetypes.ObjectValuable.
func (o SharedDataObjectUpdate) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SharedDataObjectUpdate) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SharedDataObjectUpdate) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SharedDataObjectUpdate) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SharedDataObjectUpdate) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SharedDataObjectUpdate) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o SharedDataObjectUpdate) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"action": types.StringType,
			"data_object": basetypes.ListType{
				ElemType: SharedDataObject{}.Type(ctx),
			},
		},
	}
}

type UpdatePermissionsResponse struct {
}

func (newState *UpdatePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePermissionsResponse) {
}

func (newState *UpdatePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState UpdatePermissionsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdatePermissionsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdatePermissionsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdatePermissionsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdatePermissionsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdatePermissionsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdatePermissionsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdatePermissionsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePermissionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateProvider struct {
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

func (newState *UpdateProvider) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateProvider) {
}

func (newState *UpdateProvider) SyncEffectiveFieldsDuringRead(existingState UpdateProvider) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateProvider{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateProvider) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateProvider) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateProvider) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateProvider) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateProvider) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateProvider) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

func (newState *UpdateRecipient) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRecipient) {
}

func (newState *UpdateRecipient) SyncEffectiveFieldsDuringRead(existingState UpdateRecipient) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateRecipient{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateRecipient) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateRecipient) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateRecipient) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateRecipient) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateRecipient) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateRecipient) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRecipient) Type(ctx context.Context) attr.Type {
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

type UpdateShare struct {
	// User-provided free-form text description.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The name of the share.
	Name types.String `tfsdk:"-"`
	// New name for the share.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// Username of current owner of share.
	Owner types.String `tfsdk:"owner" tf:"computed,optional"`
	// Storage root URL for the share.
	StorageRoot types.String `tfsdk:"storage_root" tf:"optional"`
	// Array of shared data object updates.
	Updates types.List `tfsdk:"updates" tf:"optional"`
}

func (newState *UpdateShare) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateShare) {
}

func (newState *UpdateShare) SyncEffectiveFieldsDuringRead(existingState UpdateShare) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateShare{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateShare) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateShare) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateShare) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateShare) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateShare) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateShare) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type UpdateSharePermissions struct {
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

func (newState *UpdateSharePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSharePermissions) {
}

func (newState *UpdateSharePermissions) SyncEffectiveFieldsDuringRead(existingState UpdateSharePermissions) {
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
		"changes": reflect.TypeOf(catalog.PermissionsChange{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateSharePermissions{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateSharePermissions) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateSharePermissions) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateSharePermissions) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateSharePermissions) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateSharePermissions) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateSharePermissions) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// The delta sharing authentication type.

// The operator to apply for the value.

// The type of the data object.

// Whether to enable or disable sharing of data history. If not specified, the
// default is **DISABLED**.

// One of: **ACTIVE**, **PERMISSION_DENIED**.

// One of: **ADD**, **REMOVE**, **UPDATE**.
