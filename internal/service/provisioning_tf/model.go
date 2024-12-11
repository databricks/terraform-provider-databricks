// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package provisioning_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AwsCredentials struct {
	StsRole types.List `tfsdk:"sts_role" tf:"optional,object"`
}

func (newState *AwsCredentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsCredentials) {
}

func (newState *AwsCredentials) SyncEffectiveFieldsDuringRead(existingState AwsCredentials) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsCredentials) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sts_role": reflect.TypeOf(StsRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsCredentials
// only implements ToObjectValue() and Type().
func (o AwsCredentials) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sts_role": o.StsRole,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsCredentials) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sts_role": basetypes.ListType{ElemType: StsRole{}.Type(ctx)},
		},
	}
}

// GetStsRole returns the value of the StsRole field in AwsCredentials as
// a StsRole value.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsCredentials) GetStsRole(ctx context.Context) (StsRole, bool) {
	var e StsRole
	if o.StsRole.IsNull() || o.StsRole.IsUnknown() {
		return e, false
	}
	var v []StsRole
	d := o.StsRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStsRole sets the value of the StsRole field in AwsCredentials.
func (o *AwsCredentials) SetStsRole(ctx context.Context, v StsRole) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sts_role"]
	o.StsRole = types.ListValueMust(t, vs)
}

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias types.String `tfsdk:"key_alias" tf:"optional"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn types.String `tfsdk:"key_arn" tf:""`
	// The AWS KMS key region.
	KeyRegion types.String `tfsdk:"key_region" tf:""`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	ReuseKeyForClusterVolumes types.Bool `tfsdk:"reuse_key_for_cluster_volumes" tf:"optional"`
}

func (newState *AwsKeyInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsKeyInfo) {
}

func (newState *AwsKeyInfo) SyncEffectiveFieldsDuringRead(existingState AwsKeyInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsKeyInfo
// only implements ToObjectValue() and Type().
func (o AwsKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key_alias":                     o.KeyAlias,
			"key_arn":                       o.KeyArn,
			"key_region":                    o.KeyRegion,
			"reuse_key_for_cluster_volumes": o.ReuseKeyForClusterVolumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key_alias":                     types.StringType,
			"key_arn":                       types.StringType,
			"key_region":                    types.StringType,
			"reuse_key_for_cluster_volumes": types.BoolType,
		},
	}
}

type AzureWorkspaceInfo struct {
	// Azure Resource Group name
	ResourceGroup types.String `tfsdk:"resource_group" tf:"optional"`
	// Azure Subscription ID
	SubscriptionId types.String `tfsdk:"subscription_id" tf:"optional"`
}

func (newState *AzureWorkspaceInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureWorkspaceInfo) {
}

func (newState *AzureWorkspaceInfo) SyncEffectiveFieldsDuringRead(existingState AzureWorkspaceInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureWorkspaceInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureWorkspaceInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureWorkspaceInfo
// only implements ToObjectValue() and Type().
func (o AzureWorkspaceInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resource_group":  o.ResourceGroup,
			"subscription_id": o.SubscriptionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureWorkspaceInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resource_group":  types.StringType,
			"subscription_id": types.StringType,
		},
	}
}

// The general workspace configurations that are specific to cloud providers.
type CloudResourceContainer struct {
	// The general workspace configurations that are specific to Google Cloud.
	Gcp types.List `tfsdk:"gcp" tf:"optional,object"`
}

func (newState *CloudResourceContainer) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloudResourceContainer) {
}

func (newState *CloudResourceContainer) SyncEffectiveFieldsDuringRead(existingState CloudResourceContainer) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloudResourceContainer.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CloudResourceContainer) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp": reflect.TypeOf(CustomerFacingGcpCloudResourceContainer{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudResourceContainer
// only implements ToObjectValue() and Type().
func (o CloudResourceContainer) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp": o.Gcp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CloudResourceContainer) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp": basetypes.ListType{ElemType: CustomerFacingGcpCloudResourceContainer{}.Type(ctx)},
		},
	}
}

// GetGcp returns the value of the Gcp field in CloudResourceContainer as
// a CustomerFacingGcpCloudResourceContainer value.
// If the field is unknown or null, the boolean return value is false.
func (o *CloudResourceContainer) GetGcp(ctx context.Context) (CustomerFacingGcpCloudResourceContainer, bool) {
	var e CustomerFacingGcpCloudResourceContainer
	if o.Gcp.IsNull() || o.Gcp.IsUnknown() {
		return e, false
	}
	var v []CustomerFacingGcpCloudResourceContainer
	d := o.Gcp.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcp sets the value of the Gcp field in CloudResourceContainer.
func (o *CloudResourceContainer) SetGcp(ctx context.Context, v CustomerFacingGcpCloudResourceContainer) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp"]
	o.Gcp = types.ListValueMust(t, vs)
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias types.String `tfsdk:"key_alias" tf:"optional"`
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	KeyArn types.String `tfsdk:"key_arn" tf:""`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
	ReuseKeyForClusterVolumes types.Bool `tfsdk:"reuse_key_for_cluster_volumes" tf:"optional"`
}

func (newState *CreateAwsKeyInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAwsKeyInfo) {
}

func (newState *CreateAwsKeyInfo) SyncEffectiveFieldsDuringRead(existingState CreateAwsKeyInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAwsKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAwsKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAwsKeyInfo
// only implements ToObjectValue() and Type().
func (o CreateAwsKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key_alias":                     o.KeyAlias,
			"key_arn":                       o.KeyArn,
			"reuse_key_for_cluster_volumes": o.ReuseKeyForClusterVolumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAwsKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key_alias":                     types.StringType,
			"key_arn":                       types.StringType,
			"reuse_key_for_cluster_volumes": types.BoolType,
		},
	}
}

type CreateCredentialAwsCredentials struct {
	StsRole types.List `tfsdk:"sts_role" tf:"optional,object"`
}

func (newState *CreateCredentialAwsCredentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialAwsCredentials) {
}

func (newState *CreateCredentialAwsCredentials) SyncEffectiveFieldsDuringRead(existingState CreateCredentialAwsCredentials) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialAwsCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialAwsCredentials) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sts_role": reflect.TypeOf(CreateCredentialStsRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialAwsCredentials
// only implements ToObjectValue() and Type().
func (o CreateCredentialAwsCredentials) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sts_role": o.StsRole,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialAwsCredentials) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sts_role": basetypes.ListType{ElemType: CreateCredentialStsRole{}.Type(ctx)},
		},
	}
}

// GetStsRole returns the value of the StsRole field in CreateCredentialAwsCredentials as
// a CreateCredentialStsRole value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialAwsCredentials) GetStsRole(ctx context.Context) (CreateCredentialStsRole, bool) {
	var e CreateCredentialStsRole
	if o.StsRole.IsNull() || o.StsRole.IsUnknown() {
		return e, false
	}
	var v []CreateCredentialStsRole
	d := o.StsRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStsRole sets the value of the StsRole field in CreateCredentialAwsCredentials.
func (o *CreateCredentialAwsCredentials) SetStsRole(ctx context.Context, v CreateCredentialStsRole) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sts_role"]
	o.StsRole = types.ListValueMust(t, vs)
}

type CreateCredentialRequest struct {
	AwsCredentials types.List `tfsdk:"aws_credentials" tf:"object"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name" tf:""`
}

func (newState *CreateCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialRequest) {
}

func (newState *CreateCredentialRequest) SyncEffectiveFieldsDuringRead(existingState CreateCredentialRequest) {
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
		"aws_credentials": reflect.TypeOf(CreateCredentialAwsCredentials{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialRequest
// only implements ToObjectValue() and Type().
func (o CreateCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_credentials":  o.AwsCredentials,
			"credentials_name": o.CredentialsName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_credentials":  basetypes.ListType{ElemType: CreateCredentialAwsCredentials{}.Type(ctx)},
			"credentials_name": types.StringType,
		},
	}
}

// GetAwsCredentials returns the value of the AwsCredentials field in CreateCredentialRequest as
// a CreateCredentialAwsCredentials value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest) GetAwsCredentials(ctx context.Context) (CreateCredentialAwsCredentials, bool) {
	var e CreateCredentialAwsCredentials
	if o.AwsCredentials.IsNull() || o.AwsCredentials.IsUnknown() {
		return e, false
	}
	var v []CreateCredentialAwsCredentials
	d := o.AwsCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsCredentials sets the value of the AwsCredentials field in CreateCredentialRequest.
func (o *CreateCredentialRequest) SetAwsCredentials(ctx context.Context, v CreateCredentialAwsCredentials) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_credentials"]
	o.AwsCredentials = types.ListValueMust(t, vs)
}

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn types.String `tfsdk:"role_arn" tf:"optional"`
}

func (newState *CreateCredentialStsRole) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialStsRole) {
}

func (newState *CreateCredentialStsRole) SyncEffectiveFieldsDuringRead(existingState CreateCredentialStsRole) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialStsRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialStsRole) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialStsRole
// only implements ToObjectValue() and Type().
func (o CreateCredentialStsRole) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"role_arn": o.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialStsRole) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"role_arn": types.StringType,
		},
	}
}

type CreateCustomerManagedKeyRequest struct {
	AwsKeyInfo types.List `tfsdk:"aws_key_info" tf:"optional,object"`

	GcpKeyInfo types.List `tfsdk:"gcp_key_info" tf:"optional,object"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases" tf:""`
}

func (newState *CreateCustomerManagedKeyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCustomerManagedKeyRequest) {
}

func (newState *CreateCustomerManagedKeyRequest) SyncEffectiveFieldsDuringRead(existingState CreateCustomerManagedKeyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomerManagedKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCustomerManagedKeyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_key_info": reflect.TypeOf(CreateAwsKeyInfo{}),
		"gcp_key_info": reflect.TypeOf(CreateGcpKeyInfo{}),
		"use_cases":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomerManagedKeyRequest
// only implements ToObjectValue() and Type().
func (o CreateCustomerManagedKeyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_key_info": o.AwsKeyInfo,
			"gcp_key_info": o.GcpKeyInfo,
			"use_cases":    o.UseCases,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCustomerManagedKeyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_key_info": basetypes.ListType{ElemType: CreateAwsKeyInfo{}.Type(ctx)},
			"gcp_key_info": basetypes.ListType{ElemType: CreateGcpKeyInfo{}.Type(ctx)},
			"use_cases":    basetypes.ListType{ElemType: types.StringType},
		},
	}
}

// GetAwsKeyInfo returns the value of the AwsKeyInfo field in CreateCustomerManagedKeyRequest as
// a CreateAwsKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest) GetAwsKeyInfo(ctx context.Context) (CreateAwsKeyInfo, bool) {
	var e CreateAwsKeyInfo
	if o.AwsKeyInfo.IsNull() || o.AwsKeyInfo.IsUnknown() {
		return e, false
	}
	var v []CreateAwsKeyInfo
	d := o.AwsKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CreateCustomerManagedKeyRequest.
func (o *CreateCustomerManagedKeyRequest) SetAwsKeyInfo(ctx context.Context, v CreateAwsKeyInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_key_info"]
	o.AwsKeyInfo = types.ListValueMust(t, vs)
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest as
// a CreateGcpKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest) GetGcpKeyInfo(ctx context.Context) (CreateGcpKeyInfo, bool) {
	var e CreateGcpKeyInfo
	if o.GcpKeyInfo.IsNull() || o.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v []CreateGcpKeyInfo
	d := o.GcpKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest.
func (o *CreateCustomerManagedKeyRequest) SetGcpKeyInfo(ctx context.Context, v CreateGcpKeyInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_key_info"]
	o.GcpKeyInfo = types.ListValueMust(t, vs)
}

// GetUseCases returns the value of the UseCases field in CreateCustomerManagedKeyRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest) GetUseCases(ctx context.Context) ([]types.String, bool) {
	if o.UseCases.IsNull() || o.UseCases.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.UseCases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUseCases sets the value of the UseCases field in CreateCustomerManagedKeyRequest.
func (o *CreateCustomerManagedKeyRequest) SetUseCases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["use_cases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UseCases = types.ListValueMust(t, vs)
}

type CreateGcpKeyInfo struct {
	// The GCP KMS key's resource name
	KmsKeyId types.String `tfsdk:"kms_key_id" tf:""`
}

func (newState *CreateGcpKeyInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateGcpKeyInfo) {
}

func (newState *CreateGcpKeyInfo) SyncEffectiveFieldsDuringRead(existingState CreateGcpKeyInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateGcpKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateGcpKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateGcpKeyInfo
// only implements ToObjectValue() and Type().
func (o CreateGcpKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kms_key_id": o.KmsKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateGcpKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"kms_key_id": types.StringType,
		},
	}
}

type CreateNetworkRequest struct {
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	GcpNetworkInfo types.List `tfsdk:"gcp_network_info" tf:"optional,object"`
	// The human-readable name of the network configuration.
	NetworkName types.String `tfsdk:"network_name" tf:""`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds types.List `tfsdk:"security_group_ids" tf:"optional"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds types.List `tfsdk:"subnet_ids" tf:"optional"`
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	VpcEndpoints types.List `tfsdk:"vpc_endpoints" tf:"optional,object"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId types.String `tfsdk:"vpc_id" tf:"optional"`
}

func (newState *CreateNetworkRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateNetworkRequest) {
}

func (newState *CreateNetworkRequest) SyncEffectiveFieldsDuringRead(existingState CreateNetworkRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateNetworkRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_network_info":   reflect.TypeOf(GcpNetworkInfo{}),
		"security_group_ids": reflect.TypeOf(types.String{}),
		"subnet_ids":         reflect.TypeOf(types.String{}),
		"vpc_endpoints":      reflect.TypeOf(NetworkVpcEndpoints{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkRequest
// only implements ToObjectValue() and Type().
func (o CreateNetworkRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_network_info":   o.GcpNetworkInfo,
			"network_name":       o.NetworkName,
			"security_group_ids": o.SecurityGroupIds,
			"subnet_ids":         o.SubnetIds,
			"vpc_endpoints":      o.VpcEndpoints,
			"vpc_id":             o.VpcId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateNetworkRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_network_info":   basetypes.ListType{ElemType: GcpNetworkInfo{}.Type(ctx)},
			"network_name":       types.StringType,
			"security_group_ids": basetypes.ListType{ElemType: types.StringType},
			"subnet_ids":         basetypes.ListType{ElemType: types.StringType},
			"vpc_endpoints":      basetypes.ListType{ElemType: NetworkVpcEndpoints{}.Type(ctx)},
			"vpc_id":             types.StringType,
		},
	}
}

// GetGcpNetworkInfo returns the value of the GcpNetworkInfo field in CreateNetworkRequest as
// a GcpNetworkInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest) GetGcpNetworkInfo(ctx context.Context) (GcpNetworkInfo, bool) {
	var e GcpNetworkInfo
	if o.GcpNetworkInfo.IsNull() || o.GcpNetworkInfo.IsUnknown() {
		return e, false
	}
	var v []GcpNetworkInfo
	d := o.GcpNetworkInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in CreateNetworkRequest.
func (o *CreateNetworkRequest) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_network_info"]
	o.GcpNetworkInfo = types.ListValueMust(t, vs)
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in CreateNetworkRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
	if o.SecurityGroupIds.IsNull() || o.SecurityGroupIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SecurityGroupIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in CreateNetworkRequest.
func (o *CreateNetworkRequest) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in CreateNetworkRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
	if o.SubnetIds.IsNull() || o.SubnetIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SubnetIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnetIds sets the value of the SubnetIds field in CreateNetworkRequest.
func (o *CreateNetworkRequest) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SubnetIds = types.ListValueMust(t, vs)
}

// GetVpcEndpoints returns the value of the VpcEndpoints field in CreateNetworkRequest as
// a NetworkVpcEndpoints value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest) GetVpcEndpoints(ctx context.Context) (NetworkVpcEndpoints, bool) {
	var e NetworkVpcEndpoints
	if o.VpcEndpoints.IsNull() || o.VpcEndpoints.IsUnknown() {
		return e, false
	}
	var v []NetworkVpcEndpoints
	d := o.VpcEndpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in CreateNetworkRequest.
func (o *CreateNetworkRequest) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["vpc_endpoints"]
	o.VpcEndpoints = types.ListValueMust(t, vs)
}

type CreateStorageConfigurationRequest struct {
	// Root S3 bucket information.
	RootBucketInfo types.List `tfsdk:"root_bucket_info" tf:"object"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName types.String `tfsdk:"storage_configuration_name" tf:""`
}

func (newState *CreateStorageConfigurationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateStorageConfigurationRequest) {
}

func (newState *CreateStorageConfigurationRequest) SyncEffectiveFieldsDuringRead(existingState CreateStorageConfigurationRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateStorageConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateStorageConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"root_bucket_info": reflect.TypeOf(RootBucketInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateStorageConfigurationRequest
// only implements ToObjectValue() and Type().
func (o CreateStorageConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"root_bucket_info":           o.RootBucketInfo,
			"storage_configuration_name": o.StorageConfigurationName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateStorageConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"root_bucket_info":           basetypes.ListType{ElemType: RootBucketInfo{}.Type(ctx)},
			"storage_configuration_name": types.StringType,
		},
	}
}

// GetRootBucketInfo returns the value of the RootBucketInfo field in CreateStorageConfigurationRequest as
// a RootBucketInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageConfigurationRequest) GetRootBucketInfo(ctx context.Context) (RootBucketInfo, bool) {
	var e RootBucketInfo
	if o.RootBucketInfo.IsNull() || o.RootBucketInfo.IsUnknown() {
		return e, false
	}
	var v []RootBucketInfo
	d := o.RootBucketInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in CreateStorageConfigurationRequest.
func (o *CreateStorageConfigurationRequest) SetRootBucketInfo(ctx context.Context, v RootBucketInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["root_bucket_info"]
	o.RootBucketInfo = types.ListValueMust(t, vs)
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id" tf:"optional"`
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	GcpVpcEndpointInfo types.List `tfsdk:"gcp_vpc_endpoint_info" tf:"optional,object"`
	// The AWS region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name" tf:""`
}

func (newState *CreateVpcEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVpcEndpointRequest) {
}

func (newState *CreateVpcEndpointRequest) SyncEffectiveFieldsDuringRead(existingState CreateVpcEndpointRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVpcEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_vpc_endpoint_info": reflect.TypeOf(GcpVpcEndpointInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVpcEndpointRequest
// only implements ToObjectValue() and Type().
func (o CreateVpcEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_vpc_endpoint_id":   o.AwsVpcEndpointId,
			"gcp_vpc_endpoint_info": o.GcpVpcEndpointInfo,
			"region":                o.Region,
			"vpc_endpoint_name":     o.VpcEndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateVpcEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_vpc_endpoint_id":   types.StringType,
			"gcp_vpc_endpoint_info": basetypes.ListType{ElemType: GcpVpcEndpointInfo{}.Type(ctx)},
			"region":                types.StringType,
			"vpc_endpoint_name":     types.StringType,
		},
	}
}

// GetGcpVpcEndpointInfo returns the value of the GcpVpcEndpointInfo field in CreateVpcEndpointRequest as
// a GcpVpcEndpointInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateVpcEndpointRequest) GetGcpVpcEndpointInfo(ctx context.Context) (GcpVpcEndpointInfo, bool) {
	var e GcpVpcEndpointInfo
	if o.GcpVpcEndpointInfo.IsNull() || o.GcpVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v []GcpVpcEndpointInfo
	d := o.GcpVpcEndpointInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in CreateVpcEndpointRequest.
func (o *CreateVpcEndpointRequest) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_vpc_endpoint_info"]
	o.GcpVpcEndpointInfo = types.ListValueMust(t, vs)
}

type CreateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane.
	AwsRegion types.String `tfsdk:"aws_region" tf:"optional"`
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// The general workspace configurations that are specific to cloud
	// providers.
	CloudResourceContainer types.List `tfsdk:"cloud_resource_container" tf:"optional,object"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id" tf:"optional"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for the web application and REST APIs is
	// `<workspace-deployment-name>.cloud.databricks.com`. For example, if the
	// deployment name is `abcsales`, your workspace URL will be
	// `https://abcsales.cloud.databricks.com`. Hyphens are allowed. This
	// property supports only the set of characters that are allowed in a
	// subdomain.
	//
	// To set this value, you must have a deployment name prefix. Contact your
	// Databricks account team to add an account deployment name prefix to your
	// account.
	//
	// Workspace deployment names follow the account prefix and a hyphen. For
	// example, if your account's deployment prefix is `acme` and the workspace
	// deployment name is `workspace-1`, the JSON response for the
	// `deployment_name` field becomes `acme-workspace-1`. The workspace URL
	// would be `acme-workspace-1.cloud.databricks.com`.
	//
	// You can also set the `deployment_name` to the reserved keyword `EMPTY` if
	// you want the deployment name to only include the deployment prefix. For
	// example, if your account's deployment prefix is `acme` and the workspace
	// deployment name is `EMPTY`, the `deployment_name` becomes `acme` only and
	// the workspace URL is `acme.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	//
	// If a new workspace omits this property, the server generates a unique
	// deployment name for you with the pattern `dbc-xxxxxxxx-xxxx`.
	DeploymentName types.String `tfsdk:"deployment_name" tf:"optional"`
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.", All the IP range configurations must be
	// mutually exclusive. An attempt to create a workspace fails if Databricks
	// detects an IP range overlap.
	//
	// Specify custom IP ranges in CIDR format. The IP ranges for these fields
	// must not overlap, and all IP addresses must be entirely within the
	// following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
	// `192.168.0.0/16`, and `240.0.0.0/4`.
	//
	// The sizes of these IP ranges affect the maximum number of nodes for the
	// workspace.
	//
	// **Important**: Confirm the IP ranges used by your Databricks workspace
	// before creating the workspace. You cannot change them after your
	// workspace is deployed. If the IP address ranges for your Databricks are
	// too small, IP exhaustion can occur, causing your Databricks jobs to fail.
	// To determine the address range sizes that you need, Databricks provides a
	// calculator as a Microsoft Excel spreadsheet. See [calculate subnet sizes
	// for a new workspace].
	//
	// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
	GcpManagedNetworkConfig types.List `tfsdk:"gcp_managed_network_config" tf:"optional,object"`
	// The configurations for the GKE cluster of a Databricks workspace.
	GkeConfig types.List `tfsdk:"gke_config" tf:"optional,object"`
	// Whether no public IP is enabled for the workspace.
	IsNoPublicIpEnabled types.Bool `tfsdk:"is_no_public_ip_enabled" tf:"optional"`
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	Location types.String `tfsdk:"location" tf:"optional"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id" tf:"optional"`

	NetworkId types.String `tfsdk:"network_id" tf:"optional"`
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
	PricingTier types.String `tfsdk:"pricing_tier" tf:"optional"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. This ID must be specified for customers using [AWS
	// PrivateLink] for either front-end (user-to-workspace connection),
	// back-end (data plane to control plane connection), or both connection
	// types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id" tf:"optional"`
	// The ID of the workspace's storage configuration object.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:"optional"`
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id" tf:"optional"`
	// The workspace's human-readable name.
	WorkspaceName types.String `tfsdk:"workspace_name" tf:""`
}

func (newState *CreateWorkspaceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWorkspaceRequest) {
}

func (newState *CreateWorkspaceRequest) SyncEffectiveFieldsDuringRead(existingState CreateWorkspaceRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cloud_resource_container":   reflect.TypeOf(CloudResourceContainer{}),
		"custom_tags":                reflect.TypeOf(types.String{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig{}),
		"gke_config":                 reflect.TypeOf(GkeConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o CreateWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_region":                 o.AwsRegion,
			"cloud":                      o.Cloud,
			"cloud_resource_container":   o.CloudResourceContainer,
			"credentials_id":             o.CredentialsId,
			"custom_tags":                o.CustomTags,
			"deployment_name":            o.DeploymentName,
			"gcp_managed_network_config": o.GcpManagedNetworkConfig,
			"gke_config":                 o.GkeConfig,
			"is_no_public_ip_enabled":    o.IsNoPublicIpEnabled,
			"location":                   o.Location,
			"managed_services_customer_managed_key_id": o.ManagedServicesCustomerManagedKeyId,
			"network_id":                      o.NetworkId,
			"pricing_tier":                    o.PricingTier,
			"private_access_settings_id":      o.PrivateAccessSettingsId,
			"storage_configuration_id":        o.StorageConfigurationId,
			"storage_customer_managed_key_id": o.StorageCustomerManagedKeyId,
			"workspace_name":                  o.WorkspaceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_region":                 types.StringType,
			"cloud":                      types.StringType,
			"cloud_resource_container":   basetypes.ListType{ElemType: CloudResourceContainer{}.Type(ctx)},
			"credentials_id":             types.StringType,
			"custom_tags":                basetypes.MapType{ElemType: types.StringType},
			"deployment_name":            types.StringType,
			"gcp_managed_network_config": basetypes.ListType{ElemType: GcpManagedNetworkConfig{}.Type(ctx)},
			"gke_config":                 basetypes.ListType{ElemType: GkeConfig{}.Type(ctx)},
			"is_no_public_ip_enabled":    types.BoolType,
			"location":                   types.StringType,
			"managed_services_customer_managed_key_id": types.StringType,
			"network_id":                      types.StringType,
			"pricing_tier":                    types.StringType,
			"private_access_settings_id":      types.StringType,
			"storage_configuration_id":        types.StringType,
			"storage_customer_managed_key_id": types.StringType,
			"workspace_name":                  types.StringType,
		},
	}
}

// GetCloudResourceContainer returns the value of the CloudResourceContainer field in CreateWorkspaceRequest as
// a CloudResourceContainer value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer, bool) {
	var e CloudResourceContainer
	if o.CloudResourceContainer.IsNull() || o.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v []CloudResourceContainer
	d := o.CloudResourceContainer.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cloud_resource_container"]
	o.CloudResourceContainer = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in CreateWorkspaceRequest as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest as
// a GcpManagedNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig, bool) {
	var e GcpManagedNetworkConfig
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v []GcpManagedNetworkConfig
	d := o.GcpManagedNetworkConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_managed_network_config"]
	o.GcpManagedNetworkConfig = types.ListValueMust(t, vs)
}

// GetGkeConfig returns the value of the GkeConfig field in CreateWorkspaceRequest as
// a GkeConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest) GetGkeConfig(ctx context.Context) (GkeConfig, bool) {
	var e GkeConfig
	if o.GkeConfig.IsNull() || o.GkeConfig.IsUnknown() {
		return e, false
	}
	var v []GkeConfig
	d := o.GkeConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGkeConfig sets the value of the GkeConfig field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetGkeConfig(ctx context.Context, v GkeConfig) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gke_config"]
	o.GkeConfig = types.ListValueMust(t, vs)
}

type Credential struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`

	AwsCredentials types.List `tfsdk:"aws_credentials" tf:"optional,object"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed,optional"`
	// Databricks credential configuration ID.
	CredentialsId types.String `tfsdk:"credentials_id" tf:"optional"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name" tf:"optional"`
}

func (newState *Credential) SyncEffectiveFieldsDuringCreateOrUpdate(plan Credential) {
}

func (newState *Credential) SyncEffectiveFieldsDuringRead(existingState Credential) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Credential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Credential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_credentials": reflect.TypeOf(AwsCredentials{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Credential
// only implements ToObjectValue() and Type().
func (o Credential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":       o.AccountId,
			"aws_credentials":  o.AwsCredentials,
			"creation_time":    o.CreationTime,
			"credentials_id":   o.CredentialsId,
			"credentials_name": o.CredentialsName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Credential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":       types.StringType,
			"aws_credentials":  basetypes.ListType{ElemType: AwsCredentials{}.Type(ctx)},
			"creation_time":    types.Int64Type,
			"credentials_id":   types.StringType,
			"credentials_name": types.StringType,
		},
	}
}

// GetAwsCredentials returns the value of the AwsCredentials field in Credential as
// a AwsCredentials value.
// If the field is unknown or null, the boolean return value is false.
func (o *Credential) GetAwsCredentials(ctx context.Context) (AwsCredentials, bool) {
	var e AwsCredentials
	if o.AwsCredentials.IsNull() || o.AwsCredentials.IsUnknown() {
		return e, false
	}
	var v []AwsCredentials
	d := o.AwsCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsCredentials sets the value of the AwsCredentials field in Credential.
func (o *Credential) SetAwsCredentials(ctx context.Context, v AwsCredentials) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_credentials"]
	o.AwsCredentials = types.ListValueMust(t, vs)
}

// The general workspace configurations that are specific to Google Cloud.
type CustomerFacingGcpCloudResourceContainer struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	ProjectId types.String `tfsdk:"project_id" tf:"optional"`
}

func (newState *CustomerFacingGcpCloudResourceContainer) SyncEffectiveFieldsDuringCreateOrUpdate(plan CustomerFacingGcpCloudResourceContainer) {
}

func (newState *CustomerFacingGcpCloudResourceContainer) SyncEffectiveFieldsDuringRead(existingState CustomerFacingGcpCloudResourceContainer) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomerFacingGcpCloudResourceContainer.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CustomerFacingGcpCloudResourceContainer) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerFacingGcpCloudResourceContainer
// only implements ToObjectValue() and Type().
func (o CustomerFacingGcpCloudResourceContainer) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project_id": o.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomerFacingGcpCloudResourceContainer) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project_id": types.StringType,
		},
	}
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`

	AwsKeyInfo types.List `tfsdk:"aws_key_info" tf:"optional,object"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed,optional"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId types.String `tfsdk:"customer_managed_key_id" tf:"optional"`

	GcpKeyInfo types.List `tfsdk:"gcp_key_info" tf:"optional,object"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases" tf:"optional"`
}

func (newState *CustomerManagedKey) SyncEffectiveFieldsDuringCreateOrUpdate(plan CustomerManagedKey) {
}

func (newState *CustomerManagedKey) SyncEffectiveFieldsDuringRead(existingState CustomerManagedKey) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomerManagedKey.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CustomerManagedKey) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_key_info": reflect.TypeOf(AwsKeyInfo{}),
		"gcp_key_info": reflect.TypeOf(GcpKeyInfo{}),
		"use_cases":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerManagedKey
// only implements ToObjectValue() and Type().
func (o CustomerManagedKey) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":              o.AccountId,
			"aws_key_info":            o.AwsKeyInfo,
			"creation_time":           o.CreationTime,
			"customer_managed_key_id": o.CustomerManagedKeyId,
			"gcp_key_info":            o.GcpKeyInfo,
			"use_cases":               o.UseCases,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomerManagedKey) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":              types.StringType,
			"aws_key_info":            basetypes.ListType{ElemType: AwsKeyInfo{}.Type(ctx)},
			"creation_time":           types.Int64Type,
			"customer_managed_key_id": types.StringType,
			"gcp_key_info":            basetypes.ListType{ElemType: GcpKeyInfo{}.Type(ctx)},
			"use_cases":               basetypes.ListType{ElemType: types.StringType},
		},
	}
}

// GetAwsKeyInfo returns the value of the AwsKeyInfo field in CustomerManagedKey as
// a AwsKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey) GetAwsKeyInfo(ctx context.Context) (AwsKeyInfo, bool) {
	var e AwsKeyInfo
	if o.AwsKeyInfo.IsNull() || o.AwsKeyInfo.IsUnknown() {
		return e, false
	}
	var v []AwsKeyInfo
	d := o.AwsKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CustomerManagedKey.
func (o *CustomerManagedKey) SetAwsKeyInfo(ctx context.Context, v AwsKeyInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_key_info"]
	o.AwsKeyInfo = types.ListValueMust(t, vs)
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CustomerManagedKey as
// a GcpKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey) GetGcpKeyInfo(ctx context.Context) (GcpKeyInfo, bool) {
	var e GcpKeyInfo
	if o.GcpKeyInfo.IsNull() || o.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v []GcpKeyInfo
	d := o.GcpKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CustomerManagedKey.
func (o *CustomerManagedKey) SetGcpKeyInfo(ctx context.Context, v GcpKeyInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_key_info"]
	o.GcpKeyInfo = types.ListValueMust(t, vs)
}

// GetUseCases returns the value of the UseCases field in CustomerManagedKey as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey) GetUseCases(ctx context.Context) ([]types.String, bool) {
	if o.UseCases.IsNull() || o.UseCases.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.UseCases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUseCases sets the value of the UseCases field in CustomerManagedKey.
func (o *CustomerManagedKey) SetUseCases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["use_cases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UseCases = types.ListValueMust(t, vs)
}

// Delete credential configuration
type DeleteCredentialRequest struct {
	// Databricks Account API credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
}

func (newState *DeleteCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCredentialRequest) {
}

func (newState *DeleteCredentialRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCredentialRequest) {
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
			"credentials_id": o.CredentialsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id": types.StringType,
		},
	}
}

// Delete encryption key configuration
type DeleteEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId types.String `tfsdk:"-"`
}

func (newState *DeleteEncryptionKeyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteEncryptionKeyRequest) {
}

func (newState *DeleteEncryptionKeyRequest) SyncEffectiveFieldsDuringRead(existingState DeleteEncryptionKeyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEncryptionKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteEncryptionKeyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEncryptionKeyRequest
// only implements ToObjectValue() and Type().
func (o DeleteEncryptionKeyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_managed_key_id": o.CustomerManagedKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteEncryptionKeyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_managed_key_id": types.StringType,
		},
	}
}

// Delete a network configuration
type DeleteNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId types.String `tfsdk:"-"`
}

func (newState *DeleteNetworkRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteNetworkRequest) {
}

func (newState *DeleteNetworkRequest) SyncEffectiveFieldsDuringRead(existingState DeleteNetworkRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkRequest
// only implements ToObjectValue() and Type().
func (o DeleteNetworkRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_id": o.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNetworkRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_id": types.StringType,
		},
	}
}

// Delete a private access settings object
type DeletePrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (newState *DeletePrivateAccesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePrivateAccesRequest) {
}

func (newState *DeletePrivateAccesRequest) SyncEffectiveFieldsDuringRead(existingState DeletePrivateAccesRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePrivateAccesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePrivateAccesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePrivateAccesRequest
// only implements ToObjectValue() and Type().
func (o DeletePrivateAccesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_access_settings_id": o.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePrivateAccesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_access_settings_id": types.StringType,
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

// Delete storage configuration
type DeleteStorageRequest struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (newState *DeleteStorageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteStorageRequest) {
}

func (newState *DeleteStorageRequest) SyncEffectiveFieldsDuringRead(existingState DeleteStorageRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteStorageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteStorageRequest
// only implements ToObjectValue() and Type().
func (o DeleteStorageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteStorageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_configuration_id": types.StringType,
		},
	}
}

// Delete VPC endpoint configuration
type DeleteVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId types.String `tfsdk:"-"`
}

func (newState *DeleteVpcEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteVpcEndpointRequest) {
}

func (newState *DeleteVpcEndpointRequest) SyncEffectiveFieldsDuringRead(existingState DeleteVpcEndpointRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteVpcEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteVpcEndpointRequest
// only implements ToObjectValue() and Type().
func (o DeleteVpcEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vpc_endpoint_id": o.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteVpcEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vpc_endpoint_id": types.StringType,
		},
	}
}

// Delete a workspace
type DeleteWorkspaceRequest struct {
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *DeleteWorkspaceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWorkspaceRequest) {
}

func (newState *DeleteWorkspaceRequest) SyncEffectiveFieldsDuringRead(existingState DeleteWorkspaceRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o DeleteWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type ExternalCustomerInfo struct {
	// Email of the authoritative user.
	AuthoritativeUserEmail types.String `tfsdk:"authoritative_user_email" tf:"optional"`
	// The authoritative user full name.
	AuthoritativeUserFullName types.String `tfsdk:"authoritative_user_full_name" tf:"optional"`
	// The legal entity name for the external workspace
	CustomerName types.String `tfsdk:"customer_name" tf:"optional"`
}

func (newState *ExternalCustomerInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalCustomerInfo) {
}

func (newState *ExternalCustomerInfo) SyncEffectiveFieldsDuringRead(existingState ExternalCustomerInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalCustomerInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalCustomerInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalCustomerInfo
// only implements ToObjectValue() and Type().
func (o ExternalCustomerInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authoritative_user_email":     o.AuthoritativeUserEmail,
			"authoritative_user_full_name": o.AuthoritativeUserFullName,
			"customer_name":                o.CustomerName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalCustomerInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authoritative_user_email":     types.StringType,
			"authoritative_user_full_name": types.StringType,
			"customer_name":                types.StringType,
		},
	}
}

type GcpKeyInfo struct {
	// The GCP KMS key's resource name
	KmsKeyId types.String `tfsdk:"kms_key_id" tf:""`
}

func (newState *GcpKeyInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpKeyInfo) {
}

func (newState *GcpKeyInfo) SyncEffectiveFieldsDuringRead(existingState GcpKeyInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpKeyInfo
// only implements ToObjectValue() and Type().
func (o GcpKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kms_key_id": o.KmsKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"kms_key_id": types.StringType,
		},
	}
}

// The network settings for the workspace. The configurations are only for
// Databricks-managed VPCs. It is ignored if you specify a customer-managed VPC
// in the `network_id` field.", All the IP range configurations must be mutually
// exclusive. An attempt to create a workspace fails if Databricks detects an IP
// range overlap.
//
// Specify custom IP ranges in CIDR format. The IP ranges for these fields must
// not overlap, and all IP addresses must be entirely within the following
// ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`, `192.168.0.0/16`, and
// `240.0.0.0/4`.
//
// The sizes of these IP ranges affect the maximum number of nodes for the
// workspace.
//
// **Important**: Confirm the IP ranges used by your Databricks workspace before
// creating the workspace. You cannot change them after your workspace is
// deployed. If the IP address ranges for your Databricks are too small, IP
// exhaustion can occur, causing your Databricks jobs to fail. To determine the
// address range sizes that you need, Databricks provides a calculator as a
// Microsoft Excel spreadsheet. See [calculate subnet sizes for a new
// workspace].
//
// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
type GcpManagedNetworkConfig struct {
	// The IP range from which to allocate GKE cluster pods. No bigger than `/9`
	// and no smaller than `/21`.
	GkeClusterPodIpRange types.String `tfsdk:"gke_cluster_pod_ip_range" tf:"optional"`
	// The IP range from which to allocate GKE cluster services. No bigger than
	// `/16` and no smaller than `/27`.
	GkeClusterServiceIpRange types.String `tfsdk:"gke_cluster_service_ip_range" tf:"optional"`
	// The IP range from which to allocate GKE cluster nodes. No bigger than
	// `/9` and no smaller than `/29`.
	SubnetCidr types.String `tfsdk:"subnet_cidr" tf:"optional"`
}

func (newState *GcpManagedNetworkConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpManagedNetworkConfig) {
}

func (newState *GcpManagedNetworkConfig) SyncEffectiveFieldsDuringRead(existingState GcpManagedNetworkConfig) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpManagedNetworkConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpManagedNetworkConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpManagedNetworkConfig
// only implements ToObjectValue() and Type().
func (o GcpManagedNetworkConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gke_cluster_pod_ip_range":     o.GkeClusterPodIpRange,
			"gke_cluster_service_ip_range": o.GkeClusterServiceIpRange,
			"subnet_cidr":                  o.SubnetCidr,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpManagedNetworkConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gke_cluster_pod_ip_range":     types.StringType,
			"gke_cluster_service_ip_range": types.StringType,
			"subnet_cidr":                  types.StringType,
		},
	}
}

// The Google Cloud specific information for this network (for example, the VPC
// ID, subnet ID, and secondary IP ranges).
type GcpNetworkInfo struct {
	// The Google Cloud project ID of the VPC network.
	NetworkProjectId types.String `tfsdk:"network_project_id" tf:""`
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can be
	// used by only one workspace.
	PodIpRangeName types.String `tfsdk:"pod_ip_range_name" tf:""`
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// be used by only one workspace.
	ServiceIpRangeName types.String `tfsdk:"service_ip_range_name" tf:""`
	// The ID of the subnet associated with this network.
	SubnetId types.String `tfsdk:"subnet_id" tf:""`
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	SubnetRegion types.String `tfsdk:"subnet_region" tf:""`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId types.String `tfsdk:"vpc_id" tf:""`
}

func (newState *GcpNetworkInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpNetworkInfo) {
}

func (newState *GcpNetworkInfo) SyncEffectiveFieldsDuringRead(existingState GcpNetworkInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpNetworkInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpNetworkInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpNetworkInfo
// only implements ToObjectValue() and Type().
func (o GcpNetworkInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_project_id":    o.NetworkProjectId,
			"pod_ip_range_name":     o.PodIpRangeName,
			"service_ip_range_name": o.ServiceIpRangeName,
			"subnet_id":             o.SubnetId,
			"subnet_region":         o.SubnetRegion,
			"vpc_id":                o.VpcId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpNetworkInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_project_id":    types.StringType,
			"pod_ip_range_name":     types.StringType,
			"service_ip_range_name": types.StringType,
			"subnet_id":             types.StringType,
			"subnet_region":         types.StringType,
			"vpc_id":                types.StringType,
		},
	}
}

// The Google Cloud specific information for this Private Service Connect
// endpoint.
type GcpVpcEndpointInfo struct {
	// Region of the PSC endpoint.
	EndpointRegion types.String `tfsdk:"endpoint_region" tf:""`
	// The Google Cloud project ID of the VPC network where the PSC connection
	// resides.
	ProjectId types.String `tfsdk:"project_id" tf:""`
	// The unique ID of this PSC connection.
	PscConnectionId types.String `tfsdk:"psc_connection_id" tf:"optional"`
	// The name of the PSC endpoint in the Google Cloud project.
	PscEndpointName types.String `tfsdk:"psc_endpoint_name" tf:""`
	// The service attachment this PSC connection connects to.
	ServiceAttachmentId types.String `tfsdk:"service_attachment_id" tf:"optional"`
}

func (newState *GcpVpcEndpointInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpVpcEndpointInfo) {
}

func (newState *GcpVpcEndpointInfo) SyncEffectiveFieldsDuringRead(existingState GcpVpcEndpointInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpVpcEndpointInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpVpcEndpointInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpVpcEndpointInfo
// only implements ToObjectValue() and Type().
func (o GcpVpcEndpointInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_region":       o.EndpointRegion,
			"project_id":            o.ProjectId,
			"psc_connection_id":     o.PscConnectionId,
			"psc_endpoint_name":     o.PscEndpointName,
			"service_attachment_id": o.ServiceAttachmentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpVpcEndpointInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_region":       types.StringType,
			"project_id":            types.StringType,
			"psc_connection_id":     types.StringType,
			"psc_endpoint_name":     types.StringType,
			"service_attachment_id": types.StringType,
		},
	}
}

// Get credential configuration
type GetCredentialRequest struct {
	// Databricks Account API credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
}

func (newState *GetCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCredentialRequest) {
}

func (newState *GetCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GetCredentialRequest) {
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
			"credentials_id": o.CredentialsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id": types.StringType,
		},
	}
}

// Get encryption key configuration
type GetEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId types.String `tfsdk:"-"`
}

func (newState *GetEncryptionKeyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEncryptionKeyRequest) {
}

func (newState *GetEncryptionKeyRequest) SyncEffectiveFieldsDuringRead(existingState GetEncryptionKeyRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEncryptionKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEncryptionKeyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEncryptionKeyRequest
// only implements ToObjectValue() and Type().
func (o GetEncryptionKeyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_managed_key_id": o.CustomerManagedKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEncryptionKeyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_managed_key_id": types.StringType,
		},
	}
}

// Get a network configuration
type GetNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId types.String `tfsdk:"-"`
}

func (newState *GetNetworkRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetNetworkRequest) {
}

func (newState *GetNetworkRequest) SyncEffectiveFieldsDuringRead(existingState GetNetworkRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetNetworkRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkRequest
// only implements ToObjectValue() and Type().
func (o GetNetworkRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_id": o.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetNetworkRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_id": types.StringType,
		},
	}
}

// Get a private access settings object
type GetPrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (newState *GetPrivateAccesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPrivateAccesRequest) {
}

func (newState *GetPrivateAccesRequest) SyncEffectiveFieldsDuringRead(existingState GetPrivateAccesRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPrivateAccesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPrivateAccesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPrivateAccesRequest
// only implements ToObjectValue() and Type().
func (o GetPrivateAccesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_access_settings_id": o.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPrivateAccesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_access_settings_id": types.StringType,
		},
	}
}

// Get storage configuration
type GetStorageRequest struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (newState *GetStorageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStorageRequest) {
}

func (newState *GetStorageRequest) SyncEffectiveFieldsDuringRead(existingState GetStorageRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStorageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStorageRequest
// only implements ToObjectValue() and Type().
func (o GetStorageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStorageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_configuration_id": types.StringType,
		},
	}
}

// Get a VPC endpoint configuration
type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId types.String `tfsdk:"-"`
}

func (newState *GetVpcEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetVpcEndpointRequest) {
}

func (newState *GetVpcEndpointRequest) SyncEffectiveFieldsDuringRead(existingState GetVpcEndpointRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetVpcEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetVpcEndpointRequest
// only implements ToObjectValue() and Type().
func (o GetVpcEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vpc_endpoint_id": o.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetVpcEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vpc_endpoint_id": types.StringType,
		},
	}
}

// Get a workspace
type GetWorkspaceRequest struct {
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *GetWorkspaceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceRequest) {
}

func (newState *GetWorkspaceRequest) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o GetWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

// The configurations for the GKE cluster of a Databricks workspace.
type GkeConfig struct {
	// Specifies the network connectivity types for the GKE nodes and the GKE
	// master network.
	//
	// Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE cluster for the
	// workspace. The GKE nodes will not have public IPs.
	//
	// Set to `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of
	// a public GKE cluster have public IP addresses.
	ConnectivityType types.String `tfsdk:"connectivity_type" tf:"optional"`
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
	MasterIpRange types.String `tfsdk:"master_ip_range" tf:"optional"`
}

func (newState *GkeConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan GkeConfig) {
}

func (newState *GkeConfig) SyncEffectiveFieldsDuringRead(existingState GkeConfig) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GkeConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GkeConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GkeConfig
// only implements ToObjectValue() and Type().
func (o GkeConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connectivity_type": o.ConnectivityType,
			"master_ip_range":   o.MasterIpRange,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GkeConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connectivity_type": types.StringType,
			"master_ip_range":   types.StringType,
		},
	}
}

type Network struct {
	// The Databricks account ID associated with this network configuration.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// Time in epoch milliseconds when the network was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed,optional"`
	// Array of error messages about the network configuration.
	ErrorMessages types.List `tfsdk:"error_messages" tf:"computed,optional"`
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	GcpNetworkInfo types.List `tfsdk:"gcp_network_info" tf:"optional,object"`
	// The Databricks network configuration ID.
	NetworkId types.String `tfsdk:"network_id" tf:"optional"`
	// The human-readable name of the network configuration.
	NetworkName types.String `tfsdk:"network_name" tf:"optional"`

	SecurityGroupIds types.List `tfsdk:"security_group_ids" tf:"optional"`

	SubnetIds types.List `tfsdk:"subnet_ids" tf:"optional"`
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	VpcEndpoints types.List `tfsdk:"vpc_endpoints" tf:"optional,object"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId types.String `tfsdk:"vpc_id" tf:"optional"`
	// The status of this network configuration object in terms of its use in a
	// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`:
	// Broken. * `WARNED`: Warned.
	VpcStatus types.String `tfsdk:"vpc_status" tf:"computed,optional"`
	// Array of warning messages about the network configuration.
	WarningMessages types.List `tfsdk:"warning_messages" tf:"computed,optional"`
	// Workspace ID associated with this network configuration.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
}

func (newState *Network) SyncEffectiveFieldsDuringCreateOrUpdate(plan Network) {
}

func (newState *Network) SyncEffectiveFieldsDuringRead(existingState Network) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Network.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Network) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error_messages":     reflect.TypeOf(NetworkHealth{}),
		"gcp_network_info":   reflect.TypeOf(GcpNetworkInfo{}),
		"security_group_ids": reflect.TypeOf(types.String{}),
		"subnet_ids":         reflect.TypeOf(types.String{}),
		"vpc_endpoints":      reflect.TypeOf(NetworkVpcEndpoints{}),
		"warning_messages":   reflect.TypeOf(NetworkWarning{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Network
// only implements ToObjectValue() and Type().
func (o Network) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":         o.AccountId,
			"creation_time":      o.CreationTime,
			"error_messages":     o.ErrorMessages,
			"gcp_network_info":   o.GcpNetworkInfo,
			"network_id":         o.NetworkId,
			"network_name":       o.NetworkName,
			"security_group_ids": o.SecurityGroupIds,
			"subnet_ids":         o.SubnetIds,
			"vpc_endpoints":      o.VpcEndpoints,
			"vpc_id":             o.VpcId,
			"vpc_status":         o.VpcStatus,
			"warning_messages":   o.WarningMessages,
			"workspace_id":       o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Network) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":         types.StringType,
			"creation_time":      types.Int64Type,
			"error_messages":     basetypes.ListType{ElemType: basetypes.ListType{ElemType: NetworkHealth{}.Type(ctx)}},
			"gcp_network_info":   basetypes.ListType{ElemType: GcpNetworkInfo{}.Type(ctx)},
			"network_id":         types.StringType,
			"network_name":       types.StringType,
			"security_group_ids": basetypes.ListType{ElemType: types.StringType},
			"subnet_ids":         basetypes.ListType{ElemType: types.StringType},
			"vpc_endpoints":      basetypes.ListType{ElemType: NetworkVpcEndpoints{}.Type(ctx)},
			"vpc_id":             types.StringType,
			"vpc_status":         types.StringType,
			"warning_messages":   basetypes.ListType{ElemType: basetypes.ListType{ElemType: NetworkWarning{}.Type(ctx)}},
			"workspace_id":       types.Int64Type,
		},
	}
}

// GetErrorMessages returns the value of the ErrorMessages field in Network as
// a slice of NetworkHealth values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetErrorMessages(ctx context.Context) ([]NetworkHealth, bool) {
	if o.ErrorMessages.IsNull() || o.ErrorMessages.IsUnknown() {
		return nil, false
	}
	var v []NetworkHealth
	d := o.ErrorMessages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetErrorMessages sets the value of the ErrorMessages field in Network.
func (o *Network) SetErrorMessages(ctx context.Context, v []NetworkHealth) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error_messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ErrorMessages = types.ListValueMust(t, vs)
}

// GetGcpNetworkInfo returns the value of the GcpNetworkInfo field in Network as
// a GcpNetworkInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetGcpNetworkInfo(ctx context.Context) (GcpNetworkInfo, bool) {
	var e GcpNetworkInfo
	if o.GcpNetworkInfo.IsNull() || o.GcpNetworkInfo.IsUnknown() {
		return e, false
	}
	var v []GcpNetworkInfo
	d := o.GcpNetworkInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in Network.
func (o *Network) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_network_info"]
	o.GcpNetworkInfo = types.ListValueMust(t, vs)
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in Network as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
	if o.SecurityGroupIds.IsNull() || o.SecurityGroupIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SecurityGroupIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in Network.
func (o *Network) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in Network as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
	if o.SubnetIds.IsNull() || o.SubnetIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.SubnetIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnetIds sets the value of the SubnetIds field in Network.
func (o *Network) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SubnetIds = types.ListValueMust(t, vs)
}

// GetVpcEndpoints returns the value of the VpcEndpoints field in Network as
// a NetworkVpcEndpoints value.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetVpcEndpoints(ctx context.Context) (NetworkVpcEndpoints, bool) {
	var e NetworkVpcEndpoints
	if o.VpcEndpoints.IsNull() || o.VpcEndpoints.IsUnknown() {
		return e, false
	}
	var v []NetworkVpcEndpoints
	d := o.VpcEndpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in Network.
func (o *Network) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["vpc_endpoints"]
	o.VpcEndpoints = types.ListValueMust(t, vs)
}

// GetWarningMessages returns the value of the WarningMessages field in Network as
// a slice of NetworkWarning values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network) GetWarningMessages(ctx context.Context) ([]NetworkWarning, bool) {
	if o.WarningMessages.IsNull() || o.WarningMessages.IsUnknown() {
		return nil, false
	}
	var v []NetworkWarning
	d := o.WarningMessages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWarningMessages sets the value of the WarningMessages field in Network.
func (o *Network) SetWarningMessages(ctx context.Context, v []NetworkWarning) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["warning_messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WarningMessages = types.ListValueMust(t, vs)
}

type NetworkHealth struct {
	// Details of the error.
	ErrorMessage types.String `tfsdk:"error_message" tf:"optional"`
	// The AWS resource associated with this error: credentials, VPC, subnet,
	// security group, or network ACL.
	ErrorType types.String `tfsdk:"error_type" tf:"optional"`
}

func (newState *NetworkHealth) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkHealth) {
}

func (newState *NetworkHealth) SyncEffectiveFieldsDuringRead(existingState NetworkHealth) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkHealth.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkHealth) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkHealth
// only implements ToObjectValue() and Type().
func (o NetworkHealth) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_message": o.ErrorMessage,
			"error_type":    o.ErrorType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkHealth) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_message": types.StringType,
			"error_type":    types.StringType,
		},
	}
}

// If specified, contains the VPC endpoints used to allow cluster communication
// from this VPC over [AWS PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	DataplaneRelay types.List `tfsdk:"dataplane_relay" tf:""`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi types.List `tfsdk:"rest_api" tf:""`
}

func (newState *NetworkVpcEndpoints) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkVpcEndpoints) {
}

func (newState *NetworkVpcEndpoints) SyncEffectiveFieldsDuringRead(existingState NetworkVpcEndpoints) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkVpcEndpoints.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkVpcEndpoints) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataplane_relay": reflect.TypeOf(types.String{}),
		"rest_api":        reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkVpcEndpoints
// only implements ToObjectValue() and Type().
func (o NetworkVpcEndpoints) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataplane_relay": o.DataplaneRelay,
			"rest_api":        o.RestApi,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkVpcEndpoints) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataplane_relay": basetypes.ListType{ElemType: types.StringType},
			"rest_api":        basetypes.ListType{ElemType: types.StringType},
		},
	}
}

// GetDataplaneRelay returns the value of the DataplaneRelay field in NetworkVpcEndpoints as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NetworkVpcEndpoints) GetDataplaneRelay(ctx context.Context) ([]types.String, bool) {
	if o.DataplaneRelay.IsNull() || o.DataplaneRelay.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DataplaneRelay.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataplaneRelay sets the value of the DataplaneRelay field in NetworkVpcEndpoints.
func (o *NetworkVpcEndpoints) SetDataplaneRelay(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dataplane_relay"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DataplaneRelay = types.ListValueMust(t, vs)
}

// GetRestApi returns the value of the RestApi field in NetworkVpcEndpoints as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NetworkVpcEndpoints) GetRestApi(ctx context.Context) ([]types.String, bool) {
	if o.RestApi.IsNull() || o.RestApi.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.RestApi.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestApi sets the value of the RestApi field in NetworkVpcEndpoints.
func (o *NetworkVpcEndpoints) SetRestApi(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rest_api"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RestApi = types.ListValueMust(t, vs)
}

type NetworkWarning struct {
	// Details of the warning.
	WarningMessage types.String `tfsdk:"warning_message" tf:"optional"`
	// The AWS resource associated with this warning: a subnet or a security
	// group.
	WarningType types.String `tfsdk:"warning_type" tf:"optional"`
}

func (newState *NetworkWarning) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkWarning) {
}

func (newState *NetworkWarning) SyncEffectiveFieldsDuringRead(existingState NetworkWarning) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkWarning.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkWarning) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkWarning
// only implements ToObjectValue() and Type().
func (o NetworkWarning) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warning_message": o.WarningMessage,
			"warning_type":    o.WarningType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkWarning) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warning_message": types.StringType,
			"warning_type":    types.StringType,
		},
	}
}

type PrivateAccessSettings struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// An array of Databricks VPC endpoint IDs.
	AllowedVpcEndpointIds types.List `tfsdk:"allowed_vpc_endpoint_ids" tf:"optional"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel types.String `tfsdk:"private_access_level" tf:"optional"`
	// Databricks private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id" tf:"optional"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name" tf:"optional"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled" tf:"optional"`
	// The cloud region for workspaces attached to this private access settings
	// object.
	Region types.String `tfsdk:"region" tf:"optional"`
}

func (newState *PrivateAccessSettings) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrivateAccessSettings) {
}

func (newState *PrivateAccessSettings) SyncEffectiveFieldsDuringRead(existingState PrivateAccessSettings) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateAccessSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PrivateAccessSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateAccessSettings
// only implements ToObjectValue() and Type().
func (o PrivateAccessSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                   o.AccountId,
			"allowed_vpc_endpoint_ids":     o.AllowedVpcEndpointIds,
			"private_access_level":         o.PrivateAccessLevel,
			"private_access_settings_id":   o.PrivateAccessSettingsId,
			"private_access_settings_name": o.PrivateAccessSettingsName,
			"public_access_enabled":        o.PublicAccessEnabled,
			"region":                       o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PrivateAccessSettings) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":                   types.StringType,
			"allowed_vpc_endpoint_ids":     basetypes.ListType{ElemType: types.StringType},
			"private_access_level":         types.StringType,
			"private_access_settings_id":   types.StringType,
			"private_access_settings_name": types.StringType,
			"public_access_enabled":        types.BoolType,
			"region":                       types.StringType,
		},
	}
}

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in PrivateAccessSettings as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PrivateAccessSettings) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
	if o.AllowedVpcEndpointIds.IsNull() || o.AllowedVpcEndpointIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.AllowedVpcEndpointIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in PrivateAccessSettings.
func (o *PrivateAccessSettings) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type ReplaceResponse struct {
}

func (newState *ReplaceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReplaceResponse) {
}

func (newState *ReplaceResponse) SyncEffectiveFieldsDuringRead(existingState ReplaceResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplaceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReplaceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplaceResponse
// only implements ToObjectValue() and Type().
func (o ReplaceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ReplaceResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Root S3 bucket information.
type RootBucketInfo struct {
	// The name of the S3 bucket.
	BucketName types.String `tfsdk:"bucket_name" tf:"optional"`
}

func (newState *RootBucketInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RootBucketInfo) {
}

func (newState *RootBucketInfo) SyncEffectiveFieldsDuringRead(existingState RootBucketInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RootBucketInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RootBucketInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RootBucketInfo
// only implements ToObjectValue() and Type().
func (o RootBucketInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bucket_name": o.BucketName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RootBucketInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bucket_name": types.StringType,
		},
	}
}

type StorageConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id" tf:"computed,optional"`
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed,optional"`
	// Root S3 bucket information.
	RootBucketInfo types.List `tfsdk:"root_bucket_info" tf:"optional,object"`
	// Databricks storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:"optional"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName types.String `tfsdk:"storage_configuration_name" tf:"optional"`
}

func (newState *StorageConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan StorageConfiguration) {
}

func (newState *StorageConfiguration) SyncEffectiveFieldsDuringRead(existingState StorageConfiguration) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StorageConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StorageConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"root_bucket_info": reflect.TypeOf(RootBucketInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StorageConfiguration
// only implements ToObjectValue() and Type().
func (o StorageConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                 o.AccountId,
			"creation_time":              o.CreationTime,
			"root_bucket_info":           o.RootBucketInfo,
			"storage_configuration_id":   o.StorageConfigurationId,
			"storage_configuration_name": o.StorageConfigurationName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StorageConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":                 types.StringType,
			"creation_time":              types.Int64Type,
			"root_bucket_info":           basetypes.ListType{ElemType: RootBucketInfo{}.Type(ctx)},
			"storage_configuration_id":   types.StringType,
			"storage_configuration_name": types.StringType,
		},
	}
}

// GetRootBucketInfo returns the value of the RootBucketInfo field in StorageConfiguration as
// a RootBucketInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageConfiguration) GetRootBucketInfo(ctx context.Context) (RootBucketInfo, bool) {
	var e RootBucketInfo
	if o.RootBucketInfo.IsNull() || o.RootBucketInfo.IsUnknown() {
		return e, false
	}
	var v []RootBucketInfo
	d := o.RootBucketInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in StorageConfiguration.
func (o *StorageConfiguration) SetRootBucketInfo(ctx context.Context, v RootBucketInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["root_bucket_info"]
	o.RootBucketInfo = types.ListValueMust(t, vs)
}

type StsRole struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	ExternalId types.String `tfsdk:"external_id" tf:"optional"`
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn types.String `tfsdk:"role_arn" tf:"optional"`
}

func (newState *StsRole) SyncEffectiveFieldsDuringCreateOrUpdate(plan StsRole) {
}

func (newState *StsRole) SyncEffectiveFieldsDuringRead(existingState StsRole) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StsRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StsRole) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StsRole
// only implements ToObjectValue() and Type().
func (o StsRole) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
			"role_arn":    o.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StsRole) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
			"role_arn":    types.StringType,
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

type UpdateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane (for example, `us-west-2`).
	// This parameter is available only for updating failed workspaces.
	AwsRegion types.String `tfsdk:"aws_region" tf:"optional"`
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	CredentialsId types.String `tfsdk:"credentials_id" tf:"optional"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id" tf:"optional"`

	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id" tf:"optional"`
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. For failed workspaces only, you can
	// switch from a Databricks-managed VPC to a customer-managed VPC by
	// updating the workspace to add a network configuration ID.
	NetworkId types.String `tfsdk:"network_id" tf:"optional"`
	// The ID of the workspace's private access settings configuration object.
	// This parameter is available only for updating failed workspaces.
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id" tf:"optional"`
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:"optional"`
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id" tf:"optional"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UpdateWorkspaceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceRequest) {
}

func (newState *UpdateWorkspaceRequest) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_region":     o.AwsRegion,
			"credentials_id": o.CredentialsId,
			"custom_tags":    o.CustomTags,
			"managed_services_customer_managed_key_id": o.ManagedServicesCustomerManagedKeyId,
			"network_connectivity_config_id":           o.NetworkConnectivityConfigId,
			"network_id":                               o.NetworkId,
			"private_access_settings_id":               o.PrivateAccessSettingsId,
			"storage_configuration_id":                 o.StorageConfigurationId,
			"storage_customer_managed_key_id":          o.StorageCustomerManagedKeyId,
			"workspace_id":                             o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_region":     types.StringType,
			"credentials_id": types.StringType,
			"custom_tags":    basetypes.MapType{ElemType: types.StringType},
			"managed_services_customer_managed_key_id": types.StringType,
			"network_connectivity_config_id":           types.StringType,
			"network_id":                               types.StringType,
			"private_access_settings_id":               types.StringType,
			"storage_configuration_id":                 types.StringType,
			"storage_customer_managed_key_id":          types.StringType,
			"workspace_id":                             types.Int64Type,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in UpdateWorkspaceRequest as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceRequest) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in UpdateWorkspaceRequest.
func (o *UpdateWorkspaceRequest) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

type UpsertPrivateAccessSettingsRequest struct {
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID that
	// is returned when registering the VPC endpoint configuration in your
	// Databricks account. This is not the ID of the VPC endpoint in AWS.
	//
	// Only used when `private_access_level` is set to `ENDPOINT`. This is an
	// allow list of VPC endpoints that in your account that can connect to your
	// workspace over AWS PrivateLink.
	//
	// If hybrid access to your workspace is enabled by setting
	// `public_access_enabled` to `true`, this control only works for
	// PrivateLink connections. To control how your workspace is accessed via
	// public internet, see [IP access lists].
	//
	// [IP access lists]: https://docs.databricks.com/security/network/ip-access-list.html
	AllowedVpcEndpointIds types.List `tfsdk:"allowed_vpc_endpoint_ids" tf:"optional"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel types.String `tfsdk:"private_access_level" tf:"optional"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name" tf:""`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled" tf:"optional"`
	// The cloud region for workspaces associated with this private access
	// settings object.
	Region types.String `tfsdk:"region" tf:""`
}

func (newState *UpsertPrivateAccessSettingsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpsertPrivateAccessSettingsRequest) {
}

func (newState *UpsertPrivateAccessSettingsRequest) SyncEffectiveFieldsDuringRead(existingState UpsertPrivateAccessSettingsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertPrivateAccessSettingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpsertPrivateAccessSettingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertPrivateAccessSettingsRequest
// only implements ToObjectValue() and Type().
func (o UpsertPrivateAccessSettingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_vpc_endpoint_ids":     o.AllowedVpcEndpointIds,
			"private_access_level":         o.PrivateAccessLevel,
			"private_access_settings_id":   o.PrivateAccessSettingsId,
			"private_access_settings_name": o.PrivateAccessSettingsName,
			"public_access_enabled":        o.PublicAccessEnabled,
			"region":                       o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpsertPrivateAccessSettingsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_vpc_endpoint_ids":     basetypes.ListType{ElemType: types.StringType},
			"private_access_level":         types.StringType,
			"private_access_settings_id":   types.StringType,
			"private_access_settings_name": types.StringType,
			"public_access_enabled":        types.BoolType,
			"region":                       types.StringType,
		},
	}
}

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in UpsertPrivateAccessSettingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpsertPrivateAccessSettingsRequest) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
	if o.AllowedVpcEndpointIds.IsNull() || o.AllowedVpcEndpointIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.AllowedVpcEndpointIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in UpsertPrivateAccessSettingsRequest.
func (o *UpsertPrivateAccessSettingsRequest) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// The AWS Account in which the VPC endpoint object exists.
	AwsAccountId types.String `tfsdk:"aws_account_id" tf:"optional"`
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	AwsEndpointServiceId types.String `tfsdk:"aws_endpoint_service_id" tf:"optional"`
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id" tf:"optional"`
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	GcpVpcEndpointInfo types.List `tfsdk:"gcp_vpc_endpoint_info" tf:"optional,object"`
	// The AWS region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region" tf:"optional"`
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	State types.String `tfsdk:"state" tf:"optional"`
	// This enumeration represents the type of Databricks VPC [endpoint service]
	// that was used when creating this VPC endpoint.
	//
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	UseCase types.String `tfsdk:"use_case" tf:"optional"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId types.String `tfsdk:"vpc_endpoint_id" tf:"optional"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name" tf:"optional"`
}

func (newState *VpcEndpoint) SyncEffectiveFieldsDuringCreateOrUpdate(plan VpcEndpoint) {
}

func (newState *VpcEndpoint) SyncEffectiveFieldsDuringRead(existingState VpcEndpoint) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VpcEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a VpcEndpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_vpc_endpoint_info": reflect.TypeOf(GcpVpcEndpointInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VpcEndpoint
// only implements ToObjectValue() and Type().
func (o VpcEndpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":              o.AccountId,
			"aws_account_id":          o.AwsAccountId,
			"aws_endpoint_service_id": o.AwsEndpointServiceId,
			"aws_vpc_endpoint_id":     o.AwsVpcEndpointId,
			"gcp_vpc_endpoint_info":   o.GcpVpcEndpointInfo,
			"region":                  o.Region,
			"state":                   o.State,
			"use_case":                o.UseCase,
			"vpc_endpoint_id":         o.VpcEndpointId,
			"vpc_endpoint_name":       o.VpcEndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o VpcEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":              types.StringType,
			"aws_account_id":          types.StringType,
			"aws_endpoint_service_id": types.StringType,
			"aws_vpc_endpoint_id":     types.StringType,
			"gcp_vpc_endpoint_info":   basetypes.ListType{ElemType: GcpVpcEndpointInfo{}.Type(ctx)},
			"region":                  types.StringType,
			"state":                   types.StringType,
			"use_case":                types.StringType,
			"vpc_endpoint_id":         types.StringType,
			"vpc_endpoint_name":       types.StringType,
		},
	}
}

// GetGcpVpcEndpointInfo returns the value of the GcpVpcEndpointInfo field in VpcEndpoint as
// a GcpVpcEndpointInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *VpcEndpoint) GetGcpVpcEndpointInfo(ctx context.Context) (GcpVpcEndpointInfo, bool) {
	var e GcpVpcEndpointInfo
	if o.GcpVpcEndpointInfo.IsNull() || o.GcpVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v []GcpVpcEndpointInfo
	d := o.GcpVpcEndpointInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in VpcEndpoint.
func (o *VpcEndpoint) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_vpc_endpoint_info"]
	o.GcpVpcEndpointInfo = types.ListValueMust(t, vs)
}

type Workspace struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	AwsRegion types.String `tfsdk:"aws_region" tf:"optional"`

	AzureWorkspaceInfo types.List `tfsdk:"azure_workspace_info" tf:"optional,object"`
	// The cloud name. This field always has the value `gcp`.
	Cloud types.String `tfsdk:"cloud" tf:"optional"`
	// The general workspace configurations that are specific to cloud
	// providers.
	CloudResourceContainer types.List `tfsdk:"cloud_resource_container" tf:"optional,object"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed,optional"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id" tf:"optional"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags types.Map `tfsdk:"custom_tags" tf:"optional"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	DeploymentName types.String `tfsdk:"deployment_name" tf:"optional"`
	// If this workspace is for a external customer, then external_customer_info
	// is populated. If this workspace is not for a external customer, then
	// external_customer_info is empty.
	ExternalCustomerInfo types.List `tfsdk:"external_customer_info" tf:"optional,object"`
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.", All the IP range configurations must be
	// mutually exclusive. An attempt to create a workspace fails if Databricks
	// detects an IP range overlap.
	//
	// Specify custom IP ranges in CIDR format. The IP ranges for these fields
	// must not overlap, and all IP addresses must be entirely within the
	// following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
	// `192.168.0.0/16`, and `240.0.0.0/4`.
	//
	// The sizes of these IP ranges affect the maximum number of nodes for the
	// workspace.
	//
	// **Important**: Confirm the IP ranges used by your Databricks workspace
	// before creating the workspace. You cannot change them after your
	// workspace is deployed. If the IP address ranges for your Databricks are
	// too small, IP exhaustion can occur, causing your Databricks jobs to fail.
	// To determine the address range sizes that you need, Databricks provides a
	// calculator as a Microsoft Excel spreadsheet. See [calculate subnet sizes
	// for a new workspace].
	//
	// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
	GcpManagedNetworkConfig types.List `tfsdk:"gcp_managed_network_config" tf:"optional,object"`
	// The configurations for the GKE cluster of a Databricks workspace.
	GkeConfig types.List `tfsdk:"gke_config" tf:"optional,object"`
	// Whether no public IP is enabled for the workspace.
	IsNoPublicIpEnabled types.Bool `tfsdk:"is_no_public_ip_enabled" tf:"optional"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location types.String `tfsdk:"location" tf:"optional"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id" tf:"optional"`
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	NetworkId types.String `tfsdk:"network_id" tf:"optional"`
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
	PricingTier types.String `tfsdk:"pricing_tier" tf:"optional"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. You must specify this ID if you are using [AWS PrivateLink]
	// for either front-end (user-to-workspace connection), back-end (data plane
	// to control plane connection), or both connection types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id" tf:"optional"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:"optional"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id" tf:"optional"`
	// A unique integer ID for the workspace
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
	// The human-readable name of the workspace.
	WorkspaceName types.String `tfsdk:"workspace_name" tf:"optional"`
	// The status of the workspace. For workspace creation, usually it is set to
	// `PROVISIONING` initially. Continue to check the status until the status
	// is `RUNNING`.
	WorkspaceStatus types.String `tfsdk:"workspace_status" tf:"computed,optional"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage types.String `tfsdk:"workspace_status_message" tf:"computed,optional"`
}

func (newState *Workspace) SyncEffectiveFieldsDuringCreateOrUpdate(plan Workspace) {
}

func (newState *Workspace) SyncEffectiveFieldsDuringRead(existingState Workspace) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Workspace.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Workspace) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"azure_workspace_info":       reflect.TypeOf(AzureWorkspaceInfo{}),
		"cloud_resource_container":   reflect.TypeOf(CloudResourceContainer{}),
		"custom_tags":                reflect.TypeOf(types.String{}),
		"external_customer_info":     reflect.TypeOf(ExternalCustomerInfo{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig{}),
		"gke_config":                 reflect.TypeOf(GkeConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Workspace
// only implements ToObjectValue() and Type().
func (o Workspace) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                 o.AccountId,
			"aws_region":                 o.AwsRegion,
			"azure_workspace_info":       o.AzureWorkspaceInfo,
			"cloud":                      o.Cloud,
			"cloud_resource_container":   o.CloudResourceContainer,
			"creation_time":              o.CreationTime,
			"credentials_id":             o.CredentialsId,
			"custom_tags":                o.CustomTags,
			"deployment_name":            o.DeploymentName,
			"external_customer_info":     o.ExternalCustomerInfo,
			"gcp_managed_network_config": o.GcpManagedNetworkConfig,
			"gke_config":                 o.GkeConfig,
			"is_no_public_ip_enabled":    o.IsNoPublicIpEnabled,
			"location":                   o.Location,
			"managed_services_customer_managed_key_id": o.ManagedServicesCustomerManagedKeyId,
			"network_id":                      o.NetworkId,
			"pricing_tier":                    o.PricingTier,
			"private_access_settings_id":      o.PrivateAccessSettingsId,
			"storage_configuration_id":        o.StorageConfigurationId,
			"storage_customer_managed_key_id": o.StorageCustomerManagedKeyId,
			"workspace_id":                    o.WorkspaceId,
			"workspace_name":                  o.WorkspaceName,
			"workspace_status":                o.WorkspaceStatus,
			"workspace_status_message":        o.WorkspaceStatusMessage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Workspace) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":                 types.StringType,
			"aws_region":                 types.StringType,
			"azure_workspace_info":       basetypes.ListType{ElemType: AzureWorkspaceInfo{}.Type(ctx)},
			"cloud":                      types.StringType,
			"cloud_resource_container":   basetypes.ListType{ElemType: CloudResourceContainer{}.Type(ctx)},
			"creation_time":              types.Int64Type,
			"credentials_id":             types.StringType,
			"custom_tags":                basetypes.MapType{ElemType: types.StringType},
			"deployment_name":            types.StringType,
			"external_customer_info":     basetypes.ListType{ElemType: ExternalCustomerInfo{}.Type(ctx)},
			"gcp_managed_network_config": basetypes.ListType{ElemType: GcpManagedNetworkConfig{}.Type(ctx)},
			"gke_config":                 basetypes.ListType{ElemType: GkeConfig{}.Type(ctx)},
			"is_no_public_ip_enabled":    types.BoolType,
			"location":                   types.StringType,
			"managed_services_customer_managed_key_id": types.StringType,
			"network_id":                      types.StringType,
			"pricing_tier":                    types.StringType,
			"private_access_settings_id":      types.StringType,
			"storage_configuration_id":        types.StringType,
			"storage_customer_managed_key_id": types.StringType,
			"workspace_id":                    types.Int64Type,
			"workspace_name":                  types.StringType,
			"workspace_status":                types.StringType,
			"workspace_status_message":        types.StringType,
		},
	}
}

// GetAzureWorkspaceInfo returns the value of the AzureWorkspaceInfo field in Workspace as
// a AzureWorkspaceInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetAzureWorkspaceInfo(ctx context.Context) (AzureWorkspaceInfo, bool) {
	var e AzureWorkspaceInfo
	if o.AzureWorkspaceInfo.IsNull() || o.AzureWorkspaceInfo.IsUnknown() {
		return e, false
	}
	var v []AzureWorkspaceInfo
	d := o.AzureWorkspaceInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureWorkspaceInfo sets the value of the AzureWorkspaceInfo field in Workspace.
func (o *Workspace) SetAzureWorkspaceInfo(ctx context.Context, v AzureWorkspaceInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_workspace_info"]
	o.AzureWorkspaceInfo = types.ListValueMust(t, vs)
}

// GetCloudResourceContainer returns the value of the CloudResourceContainer field in Workspace as
// a CloudResourceContainer value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer, bool) {
	var e CloudResourceContainer
	if o.CloudResourceContainer.IsNull() || o.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v []CloudResourceContainer
	d := o.CloudResourceContainer.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in Workspace.
func (o *Workspace) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cloud_resource_container"]
	o.CloudResourceContainer = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in Workspace as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if o.CustomTags.IsNull() || o.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in Workspace.
func (o *Workspace) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetExternalCustomerInfo returns the value of the ExternalCustomerInfo field in Workspace as
// a ExternalCustomerInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetExternalCustomerInfo(ctx context.Context) (ExternalCustomerInfo, bool) {
	var e ExternalCustomerInfo
	if o.ExternalCustomerInfo.IsNull() || o.ExternalCustomerInfo.IsUnknown() {
		return e, false
	}
	var v []ExternalCustomerInfo
	d := o.ExternalCustomerInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExternalCustomerInfo sets the value of the ExternalCustomerInfo field in Workspace.
func (o *Workspace) SetExternalCustomerInfo(ctx context.Context, v ExternalCustomerInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["external_customer_info"]
	o.ExternalCustomerInfo = types.ListValueMust(t, vs)
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in Workspace as
// a GcpManagedNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig, bool) {
	var e GcpManagedNetworkConfig
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v []GcpManagedNetworkConfig
	d := o.GcpManagedNetworkConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in Workspace.
func (o *Workspace) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_managed_network_config"]
	o.GcpManagedNetworkConfig = types.ListValueMust(t, vs)
}

// GetGkeConfig returns the value of the GkeConfig field in Workspace as
// a GkeConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetGkeConfig(ctx context.Context) (GkeConfig, bool) {
	var e GkeConfig
	if o.GkeConfig.IsNull() || o.GkeConfig.IsUnknown() {
		return e, false
	}
	var v []GkeConfig
	d := o.GkeConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGkeConfig sets the value of the GkeConfig field in Workspace.
func (o *Workspace) SetGkeConfig(ctx context.Context, v GkeConfig) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gke_config"]
	o.GkeConfig = types.ListValueMust(t, vs)
}
