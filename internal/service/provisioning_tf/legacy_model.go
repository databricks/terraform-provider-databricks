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
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AwsCredentials_SdkV2 struct {
	StsRole types.List `tfsdk:"sts_role" tf:"optional,object"`
}

func (newState *AwsCredentials_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsCredentials_SdkV2) {
}

func (newState *AwsCredentials_SdkV2) SyncEffectiveFieldsDuringRead(existingState AwsCredentials_SdkV2) {
}

func (c AwsCredentials_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	StsRole_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "sts_role")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsCredentials_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sts_role": reflect.TypeOf(StsRole_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsCredentials_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsCredentials_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sts_role": o.StsRole,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AwsCredentials_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sts_role": basetypes.ListType{
				ElemType: StsRole_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetStsRole returns the value of the StsRole field in AwsCredentials_SdkV2 as
// a StsRole_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AwsCredentials_SdkV2) GetStsRole(ctx context.Context) (StsRole_SdkV2, bool) {
	var e StsRole_SdkV2
	if o.StsRole.IsNull() || o.StsRole.IsUnknown() {
		return e, false
	}
	var v []StsRole_SdkV2
	d := o.StsRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStsRole sets the value of the StsRole field in AwsCredentials_SdkV2.
func (o *AwsCredentials_SdkV2) SetStsRole(ctx context.Context, v StsRole_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sts_role"]
	o.StsRole = types.ListValueMust(t, vs)
}

type AwsKeyInfo_SdkV2 struct {
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

func (newState *AwsKeyInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsKeyInfo_SdkV2) {
}

func (newState *AwsKeyInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState AwsKeyInfo_SdkV2) {
}

func (c AwsKeyInfo_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "key_arn")...)
	cs.SetRequired(append(path, "key_region")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AwsKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AwsKeyInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsKeyInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o AwsKeyInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AwsKeyInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key_alias":                     types.StringType,
			"key_arn":                       types.StringType,
			"key_region":                    types.StringType,
			"reuse_key_for_cluster_volumes": types.BoolType,
		},
	}
}

type AzureWorkspaceInfo_SdkV2 struct {
	// Azure Resource Group name
	ResourceGroup types.String `tfsdk:"resource_group" tf:"optional"`
	// Azure Subscription ID
	SubscriptionId types.String `tfsdk:"subscription_id" tf:"optional"`
}

func (newState *AzureWorkspaceInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureWorkspaceInfo_SdkV2) {
}

func (newState *AzureWorkspaceInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState AzureWorkspaceInfo_SdkV2) {
}

func (c AzureWorkspaceInfo_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureWorkspaceInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureWorkspaceInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureWorkspaceInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureWorkspaceInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resource_group":  o.ResourceGroup,
			"subscription_id": o.SubscriptionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureWorkspaceInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resource_group":  types.StringType,
			"subscription_id": types.StringType,
		},
	}
}

// The general workspace configurations that are specific to cloud providers.
type CloudResourceContainer_SdkV2 struct {
	// The general workspace configurations that are specific to Google Cloud.
	Gcp types.List `tfsdk:"gcp" tf:"optional,object"`
}

func (newState *CloudResourceContainer_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloudResourceContainer_SdkV2) {
}

func (newState *CloudResourceContainer_SdkV2) SyncEffectiveFieldsDuringRead(existingState CloudResourceContainer_SdkV2) {
}

func (c CloudResourceContainer_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	CustomerFacingGcpCloudResourceContainer_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gcp")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CloudResourceContainer.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CloudResourceContainer_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp": reflect.TypeOf(CustomerFacingGcpCloudResourceContainer_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudResourceContainer_SdkV2
// only implements ToObjectValue() and Type().
func (o CloudResourceContainer_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp": o.Gcp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CloudResourceContainer_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp": basetypes.ListType{
				ElemType: CustomerFacingGcpCloudResourceContainer_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetGcp returns the value of the Gcp field in CloudResourceContainer_SdkV2 as
// a CustomerFacingGcpCloudResourceContainer_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CloudResourceContainer_SdkV2) GetGcp(ctx context.Context) (CustomerFacingGcpCloudResourceContainer_SdkV2, bool) {
	var e CustomerFacingGcpCloudResourceContainer_SdkV2
	if o.Gcp.IsNull() || o.Gcp.IsUnknown() {
		return e, false
	}
	var v []CustomerFacingGcpCloudResourceContainer_SdkV2
	d := o.Gcp.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcp sets the value of the Gcp field in CloudResourceContainer_SdkV2.
func (o *CloudResourceContainer_SdkV2) SetGcp(ctx context.Context, v CustomerFacingGcpCloudResourceContainer_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp"]
	o.Gcp = types.ListValueMust(t, vs)
}

type CreateAwsKeyInfo_SdkV2 struct {
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

func (newState *CreateAwsKeyInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAwsKeyInfo_SdkV2) {
}

func (newState *CreateAwsKeyInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateAwsKeyInfo_SdkV2) {
}

func (c CreateAwsKeyInfo_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "key_arn")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAwsKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAwsKeyInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAwsKeyInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateAwsKeyInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key_alias":                     o.KeyAlias,
			"key_arn":                       o.KeyArn,
			"reuse_key_for_cluster_volumes": o.ReuseKeyForClusterVolumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAwsKeyInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key_alias":                     types.StringType,
			"key_arn":                       types.StringType,
			"reuse_key_for_cluster_volumes": types.BoolType,
		},
	}
}

type CreateCredentialAwsCredentials_SdkV2 struct {
	StsRole types.List `tfsdk:"sts_role" tf:"optional,object"`
}

func (newState *CreateCredentialAwsCredentials_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialAwsCredentials_SdkV2) {
}

func (newState *CreateCredentialAwsCredentials_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateCredentialAwsCredentials_SdkV2) {
}

func (c CreateCredentialAwsCredentials_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	CreateCredentialStsRole_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "sts_role")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialAwsCredentials.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialAwsCredentials_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sts_role": reflect.TypeOf(CreateCredentialStsRole_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialAwsCredentials_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCredentialAwsCredentials_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sts_role": o.StsRole,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialAwsCredentials_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sts_role": basetypes.ListType{
				ElemType: CreateCredentialStsRole_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetStsRole returns the value of the StsRole field in CreateCredentialAwsCredentials_SdkV2 as
// a CreateCredentialStsRole_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialAwsCredentials_SdkV2) GetStsRole(ctx context.Context) (CreateCredentialStsRole_SdkV2, bool) {
	var e CreateCredentialStsRole_SdkV2
	if o.StsRole.IsNull() || o.StsRole.IsUnknown() {
		return e, false
	}
	var v []CreateCredentialStsRole_SdkV2
	d := o.StsRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStsRole sets the value of the StsRole field in CreateCredentialAwsCredentials_SdkV2.
func (o *CreateCredentialAwsCredentials_SdkV2) SetStsRole(ctx context.Context, v CreateCredentialStsRole_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sts_role"]
	o.StsRole = types.ListValueMust(t, vs)
}

type CreateCredentialRequest_SdkV2 struct {
	AwsCredentials types.List `tfsdk:"aws_credentials" tf:"object"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name" tf:""`
}

func (newState *CreateCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialRequest_SdkV2) {
}

func (newState *CreateCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateCredentialRequest_SdkV2) {
}

func (c CreateCredentialRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "aws_credentials")...)
	CreateCredentialAwsCredentials_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "aws_credentials")...)
	cs.SetRequired(append(path, "credentials_name")...)

	return cs
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
		"aws_credentials": reflect.TypeOf(CreateCredentialAwsCredentials_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_credentials":  o.AwsCredentials,
			"credentials_name": o.CredentialsName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_credentials": basetypes.ListType{
				ElemType: CreateCredentialAwsCredentials_SdkV2{}.Type(ctx),
			},
			"credentials_name": types.StringType,
		},
	}
}

// GetAwsCredentials returns the value of the AwsCredentials field in CreateCredentialRequest_SdkV2 as
// a CreateCredentialAwsCredentials_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCredentialRequest_SdkV2) GetAwsCredentials(ctx context.Context) (CreateCredentialAwsCredentials_SdkV2, bool) {
	var e CreateCredentialAwsCredentials_SdkV2
	if o.AwsCredentials.IsNull() || o.AwsCredentials.IsUnknown() {
		return e, false
	}
	var v []CreateCredentialAwsCredentials_SdkV2
	d := o.AwsCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsCredentials sets the value of the AwsCredentials field in CreateCredentialRequest_SdkV2.
func (o *CreateCredentialRequest_SdkV2) SetAwsCredentials(ctx context.Context, v CreateCredentialAwsCredentials_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_credentials"]
	o.AwsCredentials = types.ListValueMust(t, vs)
}

type CreateCredentialStsRole_SdkV2 struct {
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn types.String `tfsdk:"role_arn" tf:"optional"`
}

func (newState *CreateCredentialStsRole_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialStsRole_SdkV2) {
}

func (newState *CreateCredentialStsRole_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateCredentialStsRole_SdkV2) {
}

func (c CreateCredentialStsRole_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialStsRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialStsRole_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialStsRole_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCredentialStsRole_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"role_arn": o.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialStsRole_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"role_arn": types.StringType,
		},
	}
}

type CreateCustomerManagedKeyRequest_SdkV2 struct {
	AwsKeyInfo types.List `tfsdk:"aws_key_info" tf:"optional,object"`

	GcpKeyInfo types.List `tfsdk:"gcp_key_info" tf:"optional,object"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases" tf:""`
}

func (newState *CreateCustomerManagedKeyRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCustomerManagedKeyRequest_SdkV2) {
}

func (newState *CreateCustomerManagedKeyRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateCustomerManagedKeyRequest_SdkV2) {
}

func (c CreateCustomerManagedKeyRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	CreateAwsKeyInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "aws_key_info")...)
	CreateGcpKeyInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gcp_key_info")...)
	cs.SetRequired(append(path, "use_cases")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomerManagedKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCustomerManagedKeyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_key_info": reflect.TypeOf(CreateAwsKeyInfo_SdkV2{}),
		"gcp_key_info": reflect.TypeOf(CreateGcpKeyInfo_SdkV2{}),
		"use_cases":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomerManagedKeyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCustomerManagedKeyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_key_info": o.AwsKeyInfo,
			"gcp_key_info": o.GcpKeyInfo,
			"use_cases":    o.UseCases,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCustomerManagedKeyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_key_info": basetypes.ListType{
				ElemType: CreateAwsKeyInfo_SdkV2{}.Type(ctx),
			},
			"gcp_key_info": basetypes.ListType{
				ElemType: CreateGcpKeyInfo_SdkV2{}.Type(ctx),
			},
			"use_cases": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAwsKeyInfo returns the value of the AwsKeyInfo field in CreateCustomerManagedKeyRequest_SdkV2 as
// a CreateAwsKeyInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest_SdkV2) GetAwsKeyInfo(ctx context.Context) (CreateAwsKeyInfo_SdkV2, bool) {
	var e CreateAwsKeyInfo_SdkV2
	if o.AwsKeyInfo.IsNull() || o.AwsKeyInfo.IsUnknown() {
		return e, false
	}
	var v []CreateAwsKeyInfo_SdkV2
	d := o.AwsKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CreateCustomerManagedKeyRequest_SdkV2.
func (o *CreateCustomerManagedKeyRequest_SdkV2) SetAwsKeyInfo(ctx context.Context, v CreateAwsKeyInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_key_info"]
	o.AwsKeyInfo = types.ListValueMust(t, vs)
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest_SdkV2 as
// a CreateGcpKeyInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest_SdkV2) GetGcpKeyInfo(ctx context.Context) (CreateGcpKeyInfo_SdkV2, bool) {
	var e CreateGcpKeyInfo_SdkV2
	if o.GcpKeyInfo.IsNull() || o.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v []CreateGcpKeyInfo_SdkV2
	d := o.GcpKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest_SdkV2.
func (o *CreateCustomerManagedKeyRequest_SdkV2) SetGcpKeyInfo(ctx context.Context, v CreateGcpKeyInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_key_info"]
	o.GcpKeyInfo = types.ListValueMust(t, vs)
}

// GetUseCases returns the value of the UseCases field in CreateCustomerManagedKeyRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest_SdkV2) GetUseCases(ctx context.Context) ([]types.String, bool) {
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

// SetUseCases sets the value of the UseCases field in CreateCustomerManagedKeyRequest_SdkV2.
func (o *CreateCustomerManagedKeyRequest_SdkV2) SetUseCases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["use_cases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UseCases = types.ListValueMust(t, vs)
}

type CreateGcpKeyInfo_SdkV2 struct {
	// The GCP KMS key's resource name
	KmsKeyId types.String `tfsdk:"kms_key_id" tf:""`
}

func (newState *CreateGcpKeyInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateGcpKeyInfo_SdkV2) {
}

func (newState *CreateGcpKeyInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateGcpKeyInfo_SdkV2) {
}

func (c CreateGcpKeyInfo_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "kms_key_id")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateGcpKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateGcpKeyInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateGcpKeyInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateGcpKeyInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kms_key_id": o.KmsKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateGcpKeyInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"kms_key_id": types.StringType,
		},
	}
}

type CreateNetworkRequest_SdkV2 struct {
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

func (newState *CreateNetworkRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateNetworkRequest_SdkV2) {
}

func (newState *CreateNetworkRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateNetworkRequest_SdkV2) {
}

func (c CreateNetworkRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	GcpNetworkInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gcp_network_info")...)
	cs.SetRequired(append(path, "network_name")...)
	NetworkVpcEndpoints_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "vpc_endpoints")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateNetworkRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_network_info":   reflect.TypeOf(GcpNetworkInfo_SdkV2{}),
		"security_group_ids": reflect.TypeOf(types.String{}),
		"subnet_ids":         reflect.TypeOf(types.String{}),
		"vpc_endpoints":      reflect.TypeOf(NetworkVpcEndpoints_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateNetworkRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateNetworkRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_network_info": basetypes.ListType{
				ElemType: GcpNetworkInfo_SdkV2{}.Type(ctx),
			},
			"network_name": types.StringType,
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnet_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"vpc_endpoints": basetypes.ListType{
				ElemType: NetworkVpcEndpoints_SdkV2{}.Type(ctx),
			},
			"vpc_id": types.StringType,
		},
	}
}

// GetGcpNetworkInfo returns the value of the GcpNetworkInfo field in CreateNetworkRequest_SdkV2 as
// a GcpNetworkInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest_SdkV2) GetGcpNetworkInfo(ctx context.Context) (GcpNetworkInfo_SdkV2, bool) {
	var e GcpNetworkInfo_SdkV2
	if o.GcpNetworkInfo.IsNull() || o.GcpNetworkInfo.IsUnknown() {
		return e, false
	}
	var v []GcpNetworkInfo_SdkV2
	d := o.GcpNetworkInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in CreateNetworkRequest_SdkV2.
func (o *CreateNetworkRequest_SdkV2) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_network_info"]
	o.GcpNetworkInfo = types.ListValueMust(t, vs)
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in CreateNetworkRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest_SdkV2) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
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

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in CreateNetworkRequest_SdkV2.
func (o *CreateNetworkRequest_SdkV2) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in CreateNetworkRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest_SdkV2) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
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

// SetSubnetIds sets the value of the SubnetIds field in CreateNetworkRequest_SdkV2.
func (o *CreateNetworkRequest_SdkV2) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SubnetIds = types.ListValueMust(t, vs)
}

// GetVpcEndpoints returns the value of the VpcEndpoints field in CreateNetworkRequest_SdkV2 as
// a NetworkVpcEndpoints_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkRequest_SdkV2) GetVpcEndpoints(ctx context.Context) (NetworkVpcEndpoints_SdkV2, bool) {
	var e NetworkVpcEndpoints_SdkV2
	if o.VpcEndpoints.IsNull() || o.VpcEndpoints.IsUnknown() {
		return e, false
	}
	var v []NetworkVpcEndpoints_SdkV2
	d := o.VpcEndpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in CreateNetworkRequest_SdkV2.
func (o *CreateNetworkRequest_SdkV2) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["vpc_endpoints"]
	o.VpcEndpoints = types.ListValueMust(t, vs)
}

type CreateStorageConfigurationRequest_SdkV2 struct {
	// Root S3 bucket information.
	RootBucketInfo types.List `tfsdk:"root_bucket_info" tf:"object"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName types.String `tfsdk:"storage_configuration_name" tf:""`
}

func (newState *CreateStorageConfigurationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateStorageConfigurationRequest_SdkV2) {
}

func (newState *CreateStorageConfigurationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateStorageConfigurationRequest_SdkV2) {
}

func (c CreateStorageConfigurationRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "root_bucket_info")...)
	RootBucketInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "root_bucket_info")...)
	cs.SetRequired(append(path, "storage_configuration_name")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateStorageConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateStorageConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"root_bucket_info": reflect.TypeOf(RootBucketInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateStorageConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateStorageConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"root_bucket_info":           o.RootBucketInfo,
			"storage_configuration_name": o.StorageConfigurationName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateStorageConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"root_bucket_info": basetypes.ListType{
				ElemType: RootBucketInfo_SdkV2{}.Type(ctx),
			},
			"storage_configuration_name": types.StringType,
		},
	}
}

// GetRootBucketInfo returns the value of the RootBucketInfo field in CreateStorageConfigurationRequest_SdkV2 as
// a RootBucketInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateStorageConfigurationRequest_SdkV2) GetRootBucketInfo(ctx context.Context) (RootBucketInfo_SdkV2, bool) {
	var e RootBucketInfo_SdkV2
	if o.RootBucketInfo.IsNull() || o.RootBucketInfo.IsUnknown() {
		return e, false
	}
	var v []RootBucketInfo_SdkV2
	d := o.RootBucketInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in CreateStorageConfigurationRequest_SdkV2.
func (o *CreateStorageConfigurationRequest_SdkV2) SetRootBucketInfo(ctx context.Context, v RootBucketInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["root_bucket_info"]
	o.RootBucketInfo = types.ListValueMust(t, vs)
}

type CreateVpcEndpointRequest_SdkV2 struct {
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

func (newState *CreateVpcEndpointRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVpcEndpointRequest_SdkV2) {
}

func (newState *CreateVpcEndpointRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateVpcEndpointRequest_SdkV2) {
}

func (c CreateVpcEndpointRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	GcpVpcEndpointInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gcp_vpc_endpoint_info")...)
	cs.SetRequired(append(path, "vpc_endpoint_name")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVpcEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_vpc_endpoint_info": reflect.TypeOf(GcpVpcEndpointInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVpcEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateVpcEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateVpcEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_vpc_endpoint_id": types.StringType,
			"gcp_vpc_endpoint_info": basetypes.ListType{
				ElemType: GcpVpcEndpointInfo_SdkV2{}.Type(ctx),
			},
			"region":            types.StringType,
			"vpc_endpoint_name": types.StringType,
		},
	}
}

// GetGcpVpcEndpointInfo returns the value of the GcpVpcEndpointInfo field in CreateVpcEndpointRequest_SdkV2 as
// a GcpVpcEndpointInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateVpcEndpointRequest_SdkV2) GetGcpVpcEndpointInfo(ctx context.Context) (GcpVpcEndpointInfo_SdkV2, bool) {
	var e GcpVpcEndpointInfo_SdkV2
	if o.GcpVpcEndpointInfo.IsNull() || o.GcpVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v []GcpVpcEndpointInfo_SdkV2
	d := o.GcpVpcEndpointInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in CreateVpcEndpointRequest_SdkV2.
func (o *CreateVpcEndpointRequest_SdkV2) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_vpc_endpoint_info"]
	o.GcpVpcEndpointInfo = types.ListValueMust(t, vs)
}

type CreateWorkspaceRequest_SdkV2 struct {
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

func (newState *CreateWorkspaceRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWorkspaceRequest_SdkV2) {
}

func (newState *CreateWorkspaceRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateWorkspaceRequest_SdkV2) {
}

func (c CreateWorkspaceRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	CloudResourceContainer_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "cloud_resource_container")...)
	GcpManagedNetworkConfig_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gcp_managed_network_config")...)
	GkeConfig_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gke_config")...)
	cs.SetRequired(append(path, "workspace_name")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cloud_resource_container":   reflect.TypeOf(CloudResourceContainer_SdkV2{}),
		"custom_tags":                reflect.TypeOf(types.String{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig_SdkV2{}),
		"gke_config":                 reflect.TypeOf(GkeConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_region": types.StringType,
			"cloud":      types.StringType,
			"cloud_resource_container": basetypes.ListType{
				ElemType: CloudResourceContainer_SdkV2{}.Type(ctx),
			},
			"credentials_id": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"deployment_name": types.StringType,
			"gcp_managed_network_config": basetypes.ListType{
				ElemType: GcpManagedNetworkConfig_SdkV2{}.Type(ctx),
			},
			"gke_config": basetypes.ListType{
				ElemType: GkeConfig_SdkV2{}.Type(ctx),
			},
			"is_no_public_ip_enabled": types.BoolType,
			"location":                types.StringType,
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

// GetCloudResourceContainer returns the value of the CloudResourceContainer field in CreateWorkspaceRequest_SdkV2 as
// a CloudResourceContainer_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest_SdkV2) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer_SdkV2, bool) {
	var e CloudResourceContainer_SdkV2
	if o.CloudResourceContainer.IsNull() || o.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v []CloudResourceContainer_SdkV2
	d := o.CloudResourceContainer.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in CreateWorkspaceRequest_SdkV2.
func (o *CreateWorkspaceRequest_SdkV2) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cloud_resource_container"]
	o.CloudResourceContainer = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in CreateWorkspaceRequest_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in CreateWorkspaceRequest_SdkV2.
func (o *CreateWorkspaceRequest_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest_SdkV2 as
// a GcpManagedNetworkConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest_SdkV2) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig_SdkV2, bool) {
	var e GcpManagedNetworkConfig_SdkV2
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v []GcpManagedNetworkConfig_SdkV2
	d := o.GcpManagedNetworkConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest_SdkV2.
func (o *CreateWorkspaceRequest_SdkV2) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_managed_network_config"]
	o.GcpManagedNetworkConfig = types.ListValueMust(t, vs)
}

// GetGkeConfig returns the value of the GkeConfig field in CreateWorkspaceRequest_SdkV2 as
// a GkeConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest_SdkV2) GetGkeConfig(ctx context.Context) (GkeConfig_SdkV2, bool) {
	var e GkeConfig_SdkV2
	if o.GkeConfig.IsNull() || o.GkeConfig.IsUnknown() {
		return e, false
	}
	var v []GkeConfig_SdkV2
	d := o.GkeConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGkeConfig sets the value of the GkeConfig field in CreateWorkspaceRequest_SdkV2.
func (o *CreateWorkspaceRequest_SdkV2) SetGkeConfig(ctx context.Context, v GkeConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gke_config"]
	o.GkeConfig = types.ListValueMust(t, vs)
}

type Credential_SdkV2 struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`

	AwsCredentials types.List `tfsdk:"aws_credentials" tf:"optional,object"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed"`
	// Databricks credential configuration ID.
	CredentialsId types.String `tfsdk:"credentials_id" tf:"optional"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name" tf:"optional"`
}

func (newState *Credential_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Credential_SdkV2) {
}

func (newState *Credential_SdkV2) SyncEffectiveFieldsDuringRead(existingState Credential_SdkV2) {
}

func (c Credential_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AwsCredentials_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "aws_credentials")...)
	cs.SetComputed(append(path, "creation_time")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Credential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Credential_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_credentials": reflect.TypeOf(AwsCredentials_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Credential_SdkV2
// only implements ToObjectValue() and Type().
func (o Credential_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Credential_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"aws_credentials": basetypes.ListType{
				ElemType: AwsCredentials_SdkV2{}.Type(ctx),
			},
			"creation_time":    types.Int64Type,
			"credentials_id":   types.StringType,
			"credentials_name": types.StringType,
		},
	}
}

// GetAwsCredentials returns the value of the AwsCredentials field in Credential_SdkV2 as
// a AwsCredentials_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Credential_SdkV2) GetAwsCredentials(ctx context.Context) (AwsCredentials_SdkV2, bool) {
	var e AwsCredentials_SdkV2
	if o.AwsCredentials.IsNull() || o.AwsCredentials.IsUnknown() {
		return e, false
	}
	var v []AwsCredentials_SdkV2
	d := o.AwsCredentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsCredentials sets the value of the AwsCredentials field in Credential_SdkV2.
func (o *Credential_SdkV2) SetAwsCredentials(ctx context.Context, v AwsCredentials_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_credentials"]
	o.AwsCredentials = types.ListValueMust(t, vs)
}

// The general workspace configurations that are specific to Google Cloud.
type CustomerFacingGcpCloudResourceContainer_SdkV2 struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	ProjectId types.String `tfsdk:"project_id" tf:"optional"`
}

func (newState *CustomerFacingGcpCloudResourceContainer_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CustomerFacingGcpCloudResourceContainer_SdkV2) {
}

func (newState *CustomerFacingGcpCloudResourceContainer_SdkV2) SyncEffectiveFieldsDuringRead(existingState CustomerFacingGcpCloudResourceContainer_SdkV2) {
}

func (c CustomerFacingGcpCloudResourceContainer_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomerFacingGcpCloudResourceContainer.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CustomerFacingGcpCloudResourceContainer_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerFacingGcpCloudResourceContainer_SdkV2
// only implements ToObjectValue() and Type().
func (o CustomerFacingGcpCloudResourceContainer_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project_id": o.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomerFacingGcpCloudResourceContainer_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project_id": types.StringType,
		},
	}
}

type CustomerManagedKey_SdkV2 struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`

	AwsKeyInfo types.List `tfsdk:"aws_key_info" tf:"optional,object"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId types.String `tfsdk:"customer_managed_key_id" tf:"optional"`

	GcpKeyInfo types.List `tfsdk:"gcp_key_info" tf:"optional,object"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases" tf:"optional"`
}

func (newState *CustomerManagedKey_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CustomerManagedKey_SdkV2) {
}

func (newState *CustomerManagedKey_SdkV2) SyncEffectiveFieldsDuringRead(existingState CustomerManagedKey_SdkV2) {
}

func (c CustomerManagedKey_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AwsKeyInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "aws_key_info")...)
	cs.SetComputed(append(path, "creation_time")...)
	GcpKeyInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gcp_key_info")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomerManagedKey.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CustomerManagedKey_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_key_info": reflect.TypeOf(AwsKeyInfo_SdkV2{}),
		"gcp_key_info": reflect.TypeOf(GcpKeyInfo_SdkV2{}),
		"use_cases":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerManagedKey_SdkV2
// only implements ToObjectValue() and Type().
func (o CustomerManagedKey_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CustomerManagedKey_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"aws_key_info": basetypes.ListType{
				ElemType: AwsKeyInfo_SdkV2{}.Type(ctx),
			},
			"creation_time":           types.Int64Type,
			"customer_managed_key_id": types.StringType,
			"gcp_key_info": basetypes.ListType{
				ElemType: GcpKeyInfo_SdkV2{}.Type(ctx),
			},
			"use_cases": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAwsKeyInfo returns the value of the AwsKeyInfo field in CustomerManagedKey_SdkV2 as
// a AwsKeyInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey_SdkV2) GetAwsKeyInfo(ctx context.Context) (AwsKeyInfo_SdkV2, bool) {
	var e AwsKeyInfo_SdkV2
	if o.AwsKeyInfo.IsNull() || o.AwsKeyInfo.IsUnknown() {
		return e, false
	}
	var v []AwsKeyInfo_SdkV2
	d := o.AwsKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CustomerManagedKey_SdkV2.
func (o *CustomerManagedKey_SdkV2) SetAwsKeyInfo(ctx context.Context, v AwsKeyInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_key_info"]
	o.AwsKeyInfo = types.ListValueMust(t, vs)
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CustomerManagedKey_SdkV2 as
// a GcpKeyInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey_SdkV2) GetGcpKeyInfo(ctx context.Context) (GcpKeyInfo_SdkV2, bool) {
	var e GcpKeyInfo_SdkV2
	if o.GcpKeyInfo.IsNull() || o.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v []GcpKeyInfo_SdkV2
	d := o.GcpKeyInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CustomerManagedKey_SdkV2.
func (o *CustomerManagedKey_SdkV2) SetGcpKeyInfo(ctx context.Context, v GcpKeyInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_key_info"]
	o.GcpKeyInfo = types.ListValueMust(t, vs)
}

// GetUseCases returns the value of the UseCases field in CustomerManagedKey_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey_SdkV2) GetUseCases(ctx context.Context) ([]types.String, bool) {
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

// SetUseCases sets the value of the UseCases field in CustomerManagedKey_SdkV2.
func (o *CustomerManagedKey_SdkV2) SetUseCases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["use_cases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UseCases = types.ListValueMust(t, vs)
}

// Delete credential configuration
type DeleteCredentialRequest_SdkV2 struct {
	// Databricks Account API credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
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
			"credentials_id": o.CredentialsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id": types.StringType,
		},
	}
}

// Delete encryption key configuration
type DeleteEncryptionKeyRequest_SdkV2 struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEncryptionKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteEncryptionKeyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEncryptionKeyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteEncryptionKeyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_managed_key_id": o.CustomerManagedKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteEncryptionKeyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_managed_key_id": types.StringType,
		},
	}
}

// Delete a network configuration
type DeleteNetworkRequest_SdkV2 struct {
	// Databricks Account API network configuration ID.
	NetworkId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteNetworkRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_id": o.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNetworkRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_id": types.StringType,
		},
	}
}

// Delete a private access settings object
type DeletePrivateAccesRequest_SdkV2 struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePrivateAccesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePrivateAccesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePrivateAccesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeletePrivateAccesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_access_settings_id": o.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePrivateAccesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_access_settings_id": types.StringType,
		},
	}
}

type DeleteResponse_SdkV2 struct {
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

// Delete storage configuration
type DeleteStorageRequest_SdkV2 struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteStorageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteStorageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteStorageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteStorageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_configuration_id": types.StringType,
		},
	}
}

// Delete VPC endpoint configuration
type DeleteVpcEndpointRequest_SdkV2 struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteVpcEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteVpcEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteVpcEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vpc_endpoint_id": o.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteVpcEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vpc_endpoint_id": types.StringType,
		},
	}
}

// Delete a workspace
type DeleteWorkspaceRequest_SdkV2 struct {
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type ExternalCustomerInfo_SdkV2 struct {
	// Email of the authoritative user.
	AuthoritativeUserEmail types.String `tfsdk:"authoritative_user_email" tf:"optional"`
	// The authoritative user full name.
	AuthoritativeUserFullName types.String `tfsdk:"authoritative_user_full_name" tf:"optional"`
	// The legal entity name for the external workspace
	CustomerName types.String `tfsdk:"customer_name" tf:"optional"`
}

func (newState *ExternalCustomerInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalCustomerInfo_SdkV2) {
}

func (newState *ExternalCustomerInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExternalCustomerInfo_SdkV2) {
}

func (c ExternalCustomerInfo_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalCustomerInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalCustomerInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalCustomerInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ExternalCustomerInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authoritative_user_email":     o.AuthoritativeUserEmail,
			"authoritative_user_full_name": o.AuthoritativeUserFullName,
			"customer_name":                o.CustomerName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalCustomerInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authoritative_user_email":     types.StringType,
			"authoritative_user_full_name": types.StringType,
			"customer_name":                types.StringType,
		},
	}
}

type GcpKeyInfo_SdkV2 struct {
	// The GCP KMS key's resource name
	KmsKeyId types.String `tfsdk:"kms_key_id" tf:""`
}

func (newState *GcpKeyInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpKeyInfo_SdkV2) {
}

func (newState *GcpKeyInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState GcpKeyInfo_SdkV2) {
}

func (c GcpKeyInfo_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "kms_key_id")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpKeyInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpKeyInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpKeyInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kms_key_id": o.KmsKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpKeyInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
type GcpManagedNetworkConfig_SdkV2 struct {
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

func (newState *GcpManagedNetworkConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpManagedNetworkConfig_SdkV2) {
}

func (newState *GcpManagedNetworkConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState GcpManagedNetworkConfig_SdkV2) {
}

func (c GcpManagedNetworkConfig_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpManagedNetworkConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpManagedNetworkConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpManagedNetworkConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpManagedNetworkConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gke_cluster_pod_ip_range":     o.GkeClusterPodIpRange,
			"gke_cluster_service_ip_range": o.GkeClusterServiceIpRange,
			"subnet_cidr":                  o.SubnetCidr,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GcpManagedNetworkConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
type GcpNetworkInfo_SdkV2 struct {
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

func (newState *GcpNetworkInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpNetworkInfo_SdkV2) {
}

func (newState *GcpNetworkInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState GcpNetworkInfo_SdkV2) {
}

func (c GcpNetworkInfo_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "network_project_id")...)
	cs.SetRequired(append(path, "pod_ip_range_name")...)
	cs.SetRequired(append(path, "service_ip_range_name")...)
	cs.SetRequired(append(path, "subnet_id")...)
	cs.SetRequired(append(path, "subnet_region")...)
	cs.SetRequired(append(path, "vpc_id")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpNetworkInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpNetworkInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpNetworkInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpNetworkInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GcpNetworkInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
type GcpVpcEndpointInfo_SdkV2 struct {
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

func (newState *GcpVpcEndpointInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpVpcEndpointInfo_SdkV2) {
}

func (newState *GcpVpcEndpointInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState GcpVpcEndpointInfo_SdkV2) {
}

func (c GcpVpcEndpointInfo_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "endpoint_region")...)
	cs.SetRequired(append(path, "project_id")...)
	cs.SetRequired(append(path, "psc_endpoint_name")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpVpcEndpointInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpVpcEndpointInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpVpcEndpointInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o GcpVpcEndpointInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GcpVpcEndpointInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
type GetCredentialRequest_SdkV2 struct {
	// Databricks Account API credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
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
			"credentials_id": o.CredentialsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials_id": types.StringType,
		},
	}
}

// Get encryption key configuration
type GetEncryptionKeyRequest_SdkV2 struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEncryptionKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEncryptionKeyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEncryptionKeyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetEncryptionKeyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_managed_key_id": o.CustomerManagedKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEncryptionKeyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_managed_key_id": types.StringType,
		},
	}
}

// Get a network configuration
type GetNetworkRequest_SdkV2 struct {
	// Databricks Account API network configuration ID.
	NetworkId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetNetworkRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetNetworkRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_id": o.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetNetworkRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_id": types.StringType,
		},
	}
}

// Get a private access settings object
type GetPrivateAccesRequest_SdkV2 struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPrivateAccesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPrivateAccesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPrivateAccesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPrivateAccesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_access_settings_id": o.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPrivateAccesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_access_settings_id": types.StringType,
		},
	}
}

// Get storage configuration
type GetStorageRequest_SdkV2 struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStorageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStorageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetStorageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_configuration_id": o.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStorageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_configuration_id": types.StringType,
		},
	}
}

// Get a VPC endpoint configuration
type GetVpcEndpointRequest_SdkV2 struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetVpcEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetVpcEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetVpcEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vpc_endpoint_id": o.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetVpcEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vpc_endpoint_id": types.StringType,
		},
	}
}

// Get a workspace
type GetWorkspaceRequest_SdkV2 struct {
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

// The configurations for the GKE cluster of a Databricks workspace.
type GkeConfig_SdkV2 struct {
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

func (newState *GkeConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GkeConfig_SdkV2) {
}

func (newState *GkeConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState GkeConfig_SdkV2) {
}

func (c GkeConfig_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GkeConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GkeConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GkeConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o GkeConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connectivity_type": o.ConnectivityType,
			"master_ip_range":   o.MasterIpRange,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GkeConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connectivity_type": types.StringType,
			"master_ip_range":   types.StringType,
		},
	}
}

type Network_SdkV2 struct {
	// The Databricks account ID associated with this network configuration.
	AccountId types.String `tfsdk:"account_id" tf:"optional"`
	// Time in epoch milliseconds when the network was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed"`
	// Array of error messages about the network configuration.
	ErrorMessages types.List `tfsdk:"error_messages" tf:"computed"`
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
	VpcStatus types.String `tfsdk:"vpc_status" tf:"computed"`
	// Array of warning messages about the network configuration.
	WarningMessages types.List `tfsdk:"warning_messages" tf:"computed"`
	// Workspace ID associated with this network configuration.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
}

func (newState *Network_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Network_SdkV2) {
}

func (newState *Network_SdkV2) SyncEffectiveFieldsDuringRead(existingState Network_SdkV2) {
}

func (c Network_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetComputed(append(path, "creation_time")...)
	cs.SetComputed(append(path, "error_messages")...)
	NetworkHealth_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "error_messages")...)
	GcpNetworkInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gcp_network_info")...)
	NetworkVpcEndpoints_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "vpc_endpoints")...)
	cs.SetComputed(append(path, "vpc_status")...)
	cs.SetComputed(append(path, "warning_messages")...)
	NetworkWarning_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "warning_messages")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Network.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Network_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error_messages":     reflect.TypeOf(NetworkHealth_SdkV2{}),
		"gcp_network_info":   reflect.TypeOf(GcpNetworkInfo_SdkV2{}),
		"security_group_ids": reflect.TypeOf(types.String{}),
		"subnet_ids":         reflect.TypeOf(types.String{}),
		"vpc_endpoints":      reflect.TypeOf(NetworkVpcEndpoints_SdkV2{}),
		"warning_messages":   reflect.TypeOf(NetworkWarning_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Network_SdkV2
// only implements ToObjectValue() and Type().
func (o Network_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Network_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":    types.StringType,
			"creation_time": types.Int64Type,
			"error_messages": basetypes.ListType{
				ElemType: NetworkHealth_SdkV2{}.Type(ctx),
			},
			"gcp_network_info": basetypes.ListType{
				ElemType: GcpNetworkInfo_SdkV2{}.Type(ctx),
			},
			"network_id":   types.StringType,
			"network_name": types.StringType,
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnet_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"vpc_endpoints": basetypes.ListType{
				ElemType: NetworkVpcEndpoints_SdkV2{}.Type(ctx),
			},
			"vpc_id":     types.StringType,
			"vpc_status": types.StringType,
			"warning_messages": basetypes.ListType{
				ElemType: NetworkWarning_SdkV2{}.Type(ctx),
			},
			"workspace_id": types.Int64Type,
		},
	}
}

// GetErrorMessages returns the value of the ErrorMessages field in Network_SdkV2 as
// a slice of NetworkHealth_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetErrorMessages(ctx context.Context) ([]NetworkHealth_SdkV2, bool) {
	if o.ErrorMessages.IsNull() || o.ErrorMessages.IsUnknown() {
		return nil, false
	}
	var v []NetworkHealth_SdkV2
	d := o.ErrorMessages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetErrorMessages sets the value of the ErrorMessages field in Network_SdkV2.
func (o *Network_SdkV2) SetErrorMessages(ctx context.Context, v []NetworkHealth_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["error_messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ErrorMessages = types.ListValueMust(t, vs)
}

// GetGcpNetworkInfo returns the value of the GcpNetworkInfo field in Network_SdkV2 as
// a GcpNetworkInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetGcpNetworkInfo(ctx context.Context) (GcpNetworkInfo_SdkV2, bool) {
	var e GcpNetworkInfo_SdkV2
	if o.GcpNetworkInfo.IsNull() || o.GcpNetworkInfo.IsUnknown() {
		return e, false
	}
	var v []GcpNetworkInfo_SdkV2
	d := o.GcpNetworkInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in Network_SdkV2.
func (o *Network_SdkV2) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_network_info"]
	o.GcpNetworkInfo = types.ListValueMust(t, vs)
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in Network_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
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

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in Network_SdkV2.
func (o *Network_SdkV2) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in Network_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
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

// SetSubnetIds sets the value of the SubnetIds field in Network_SdkV2.
func (o *Network_SdkV2) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SubnetIds = types.ListValueMust(t, vs)
}

// GetVpcEndpoints returns the value of the VpcEndpoints field in Network_SdkV2 as
// a NetworkVpcEndpoints_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetVpcEndpoints(ctx context.Context) (NetworkVpcEndpoints_SdkV2, bool) {
	var e NetworkVpcEndpoints_SdkV2
	if o.VpcEndpoints.IsNull() || o.VpcEndpoints.IsUnknown() {
		return e, false
	}
	var v []NetworkVpcEndpoints_SdkV2
	d := o.VpcEndpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in Network_SdkV2.
func (o *Network_SdkV2) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["vpc_endpoints"]
	o.VpcEndpoints = types.ListValueMust(t, vs)
}

// GetWarningMessages returns the value of the WarningMessages field in Network_SdkV2 as
// a slice of NetworkWarning_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Network_SdkV2) GetWarningMessages(ctx context.Context) ([]NetworkWarning_SdkV2, bool) {
	if o.WarningMessages.IsNull() || o.WarningMessages.IsUnknown() {
		return nil, false
	}
	var v []NetworkWarning_SdkV2
	d := o.WarningMessages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWarningMessages sets the value of the WarningMessages field in Network_SdkV2.
func (o *Network_SdkV2) SetWarningMessages(ctx context.Context, v []NetworkWarning_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["warning_messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.WarningMessages = types.ListValueMust(t, vs)
}

type NetworkHealth_SdkV2 struct {
	// Details of the error.
	ErrorMessage types.String `tfsdk:"error_message" tf:"optional"`
	// The AWS resource associated with this error: credentials, VPC, subnet,
	// security group, or network ACL.
	ErrorType types.String `tfsdk:"error_type" tf:"optional"`
}

func (newState *NetworkHealth_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkHealth_SdkV2) {
}

func (newState *NetworkHealth_SdkV2) SyncEffectiveFieldsDuringRead(existingState NetworkHealth_SdkV2) {
}

func (c NetworkHealth_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkHealth.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkHealth_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkHealth_SdkV2
// only implements ToObjectValue() and Type().
func (o NetworkHealth_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_message": o.ErrorMessage,
			"error_type":    o.ErrorType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkHealth_SdkV2) Type(ctx context.Context) attr.Type {
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
type NetworkVpcEndpoints_SdkV2 struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	DataplaneRelay types.List `tfsdk:"dataplane_relay" tf:""`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi types.List `tfsdk:"rest_api" tf:""`
}

func (newState *NetworkVpcEndpoints_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkVpcEndpoints_SdkV2) {
}

func (newState *NetworkVpcEndpoints_SdkV2) SyncEffectiveFieldsDuringRead(existingState NetworkVpcEndpoints_SdkV2) {
}

func (c NetworkVpcEndpoints_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "dataplane_relay")...)
	cs.SetRequired(append(path, "rest_api")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkVpcEndpoints.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkVpcEndpoints_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataplane_relay": reflect.TypeOf(types.String{}),
		"rest_api":        reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkVpcEndpoints_SdkV2
// only implements ToObjectValue() and Type().
func (o NetworkVpcEndpoints_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataplane_relay": o.DataplaneRelay,
			"rest_api":        o.RestApi,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkVpcEndpoints_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataplane_relay": basetypes.ListType{
				ElemType: types.StringType,
			},
			"rest_api": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDataplaneRelay returns the value of the DataplaneRelay field in NetworkVpcEndpoints_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NetworkVpcEndpoints_SdkV2) GetDataplaneRelay(ctx context.Context) ([]types.String, bool) {
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

// SetDataplaneRelay sets the value of the DataplaneRelay field in NetworkVpcEndpoints_SdkV2.
func (o *NetworkVpcEndpoints_SdkV2) SetDataplaneRelay(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dataplane_relay"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DataplaneRelay = types.ListValueMust(t, vs)
}

// GetRestApi returns the value of the RestApi field in NetworkVpcEndpoints_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NetworkVpcEndpoints_SdkV2) GetRestApi(ctx context.Context) ([]types.String, bool) {
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

// SetRestApi sets the value of the RestApi field in NetworkVpcEndpoints_SdkV2.
func (o *NetworkVpcEndpoints_SdkV2) SetRestApi(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rest_api"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RestApi = types.ListValueMust(t, vs)
}

type NetworkWarning_SdkV2 struct {
	// Details of the warning.
	WarningMessage types.String `tfsdk:"warning_message" tf:"optional"`
	// The AWS resource associated with this warning: a subnet or a security
	// group.
	WarningType types.String `tfsdk:"warning_type" tf:"optional"`
}

func (newState *NetworkWarning_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkWarning_SdkV2) {
}

func (newState *NetworkWarning_SdkV2) SyncEffectiveFieldsDuringRead(existingState NetworkWarning_SdkV2) {
}

func (c NetworkWarning_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkWarning.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkWarning_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkWarning_SdkV2
// only implements ToObjectValue() and Type().
func (o NetworkWarning_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warning_message": o.WarningMessage,
			"warning_type":    o.WarningType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkWarning_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warning_message": types.StringType,
			"warning_type":    types.StringType,
		},
	}
}

type PrivateAccessSettings_SdkV2 struct {
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

func (newState *PrivateAccessSettings_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrivateAccessSettings_SdkV2) {
}

func (newState *PrivateAccessSettings_SdkV2) SyncEffectiveFieldsDuringRead(existingState PrivateAccessSettings_SdkV2) {
}

func (c PrivateAccessSettings_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrivateAccessSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PrivateAccessSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateAccessSettings_SdkV2
// only implements ToObjectValue() and Type().
func (o PrivateAccessSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PrivateAccessSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"allowed_vpc_endpoint_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"private_access_level":         types.StringType,
			"private_access_settings_id":   types.StringType,
			"private_access_settings_name": types.StringType,
			"public_access_enabled":        types.BoolType,
			"region":                       types.StringType,
		},
	}
}

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in PrivateAccessSettings_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PrivateAccessSettings_SdkV2) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
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

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in PrivateAccessSettings_SdkV2.
func (o *PrivateAccessSettings_SdkV2) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type ReplaceResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplaceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReplaceResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplaceResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ReplaceResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ReplaceResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Root S3 bucket information.
type RootBucketInfo_SdkV2 struct {
	// The name of the S3 bucket.
	BucketName types.String `tfsdk:"bucket_name" tf:"optional"`
}

func (newState *RootBucketInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RootBucketInfo_SdkV2) {
}

func (newState *RootBucketInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState RootBucketInfo_SdkV2) {
}

func (c RootBucketInfo_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RootBucketInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RootBucketInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RootBucketInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o RootBucketInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bucket_name": o.BucketName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RootBucketInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bucket_name": types.StringType,
		},
	}
}

type StorageConfiguration_SdkV2 struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id" tf:"computed"`
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed"`
	// Root S3 bucket information.
	RootBucketInfo types.List `tfsdk:"root_bucket_info" tf:"optional,object"`
	// Databricks storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id" tf:"optional"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName types.String `tfsdk:"storage_configuration_name" tf:"optional"`
}

func (newState *StorageConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StorageConfiguration_SdkV2) {
}

func (newState *StorageConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState StorageConfiguration_SdkV2) {
}

func (c StorageConfiguration_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetComputed(append(path, "account_id")...)
	cs.SetComputed(append(path, "creation_time")...)
	RootBucketInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "root_bucket_info")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StorageConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StorageConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"root_bucket_info": reflect.TypeOf(RootBucketInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StorageConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o StorageConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o StorageConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":    types.StringType,
			"creation_time": types.Int64Type,
			"root_bucket_info": basetypes.ListType{
				ElemType: RootBucketInfo_SdkV2{}.Type(ctx),
			},
			"storage_configuration_id":   types.StringType,
			"storage_configuration_name": types.StringType,
		},
	}
}

// GetRootBucketInfo returns the value of the RootBucketInfo field in StorageConfiguration_SdkV2 as
// a RootBucketInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *StorageConfiguration_SdkV2) GetRootBucketInfo(ctx context.Context) (RootBucketInfo_SdkV2, bool) {
	var e RootBucketInfo_SdkV2
	if o.RootBucketInfo.IsNull() || o.RootBucketInfo.IsUnknown() {
		return e, false
	}
	var v []RootBucketInfo_SdkV2
	d := o.RootBucketInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in StorageConfiguration_SdkV2.
func (o *StorageConfiguration_SdkV2) SetRootBucketInfo(ctx context.Context, v RootBucketInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["root_bucket_info"]
	o.RootBucketInfo = types.ListValueMust(t, vs)
}

type StsRole_SdkV2 struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	ExternalId types.String `tfsdk:"external_id" tf:"optional"`
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn types.String `tfsdk:"role_arn" tf:"optional"`
}

func (newState *StsRole_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StsRole_SdkV2) {
}

func (newState *StsRole_SdkV2) SyncEffectiveFieldsDuringRead(existingState StsRole_SdkV2) {
}

func (c StsRole_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StsRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StsRole_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StsRole_SdkV2
// only implements ToObjectValue() and Type().
func (o StsRole_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"external_id": o.ExternalId,
			"role_arn":    o.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StsRole_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_id": types.StringType,
			"role_arn":    types.StringType,
		},
	}
}

type UpdateResponse_SdkV2 struct {
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

type UpdateWorkspaceRequest_SdkV2 struct {
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

func (newState *UpdateWorkspaceRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceRequest_SdkV2) {
}

func (newState *UpdateWorkspaceRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceRequest_SdkV2) {
}

func (c UpdateWorkspaceRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "workspace_id")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_region":     types.StringType,
			"credentials_id": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
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

// GetCustomTags returns the value of the CustomTags field in UpdateWorkspaceRequest_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWorkspaceRequest_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in UpdateWorkspaceRequest_SdkV2.
func (o *UpdateWorkspaceRequest_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

type UpsertPrivateAccessSettingsRequest_SdkV2 struct {
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

func (newState *UpsertPrivateAccessSettingsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpsertPrivateAccessSettingsRequest_SdkV2) {
}

func (newState *UpsertPrivateAccessSettingsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpsertPrivateAccessSettingsRequest_SdkV2) {
}

func (c UpsertPrivateAccessSettingsRequest_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	cs.SetRequired(append(path, "private_access_settings_id")...)
	cs.SetRequired(append(path, "private_access_settings_name")...)
	cs.SetRequired(append(path, "region")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertPrivateAccessSettingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpsertPrivateAccessSettingsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertPrivateAccessSettingsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpsertPrivateAccessSettingsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpsertPrivateAccessSettingsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_vpc_endpoint_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"private_access_level":         types.StringType,
			"private_access_settings_id":   types.StringType,
			"private_access_settings_name": types.StringType,
			"public_access_enabled":        types.BoolType,
			"region":                       types.StringType,
		},
	}
}

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in UpsertPrivateAccessSettingsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpsertPrivateAccessSettingsRequest_SdkV2) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
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

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in UpsertPrivateAccessSettingsRequest_SdkV2.
func (o *UpsertPrivateAccessSettingsRequest_SdkV2) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type VpcEndpoint_SdkV2 struct {
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

func (newState *VpcEndpoint_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan VpcEndpoint_SdkV2) {
}

func (newState *VpcEndpoint_SdkV2) SyncEffectiveFieldsDuringRead(existingState VpcEndpoint_SdkV2) {
}

func (c VpcEndpoint_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	GcpVpcEndpointInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gcp_vpc_endpoint_info")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VpcEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a VpcEndpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_vpc_endpoint_info": reflect.TypeOf(GcpVpcEndpointInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VpcEndpoint_SdkV2
// only implements ToObjectValue() and Type().
func (o VpcEndpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o VpcEndpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":              types.StringType,
			"aws_account_id":          types.StringType,
			"aws_endpoint_service_id": types.StringType,
			"aws_vpc_endpoint_id":     types.StringType,
			"gcp_vpc_endpoint_info": basetypes.ListType{
				ElemType: GcpVpcEndpointInfo_SdkV2{}.Type(ctx),
			},
			"region":            types.StringType,
			"state":             types.StringType,
			"use_case":          types.StringType,
			"vpc_endpoint_id":   types.StringType,
			"vpc_endpoint_name": types.StringType,
		},
	}
}

// GetGcpVpcEndpointInfo returns the value of the GcpVpcEndpointInfo field in VpcEndpoint_SdkV2 as
// a GcpVpcEndpointInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *VpcEndpoint_SdkV2) GetGcpVpcEndpointInfo(ctx context.Context) (GcpVpcEndpointInfo_SdkV2, bool) {
	var e GcpVpcEndpointInfo_SdkV2
	if o.GcpVpcEndpointInfo.IsNull() || o.GcpVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v []GcpVpcEndpointInfo_SdkV2
	d := o.GcpVpcEndpointInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in VpcEndpoint_SdkV2.
func (o *VpcEndpoint_SdkV2) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_vpc_endpoint_info"]
	o.GcpVpcEndpointInfo = types.ListValueMust(t, vs)
}

type Workspace_SdkV2 struct {
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
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"computed"`
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
	WorkspaceStatus types.String `tfsdk:"workspace_status" tf:"computed"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage types.String `tfsdk:"workspace_status_message" tf:"computed"`
}

func (newState *Workspace_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Workspace_SdkV2) {
}

func (newState *Workspace_SdkV2) SyncEffectiveFieldsDuringRead(existingState Workspace_SdkV2) {
}

func (c Workspace_SdkV2) ApplySchemaCustomizations(cs tfschema.CustomizableSchema, path ...string) tfschema.CustomizableSchema {
	AzureWorkspaceInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "azure_workspace_info")...)
	CloudResourceContainer_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "cloud_resource_container")...)
	cs.SetComputed(append(path, "creation_time")...)
	ExternalCustomerInfo_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "external_customer_info")...)
	GcpManagedNetworkConfig_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gcp_managed_network_config")...)
	GkeConfig_SdkV2{}.ApplySchemaCustomizations(cs, append(path, "gke_config")...)
	cs.SetComputed(append(path, "workspace_status")...)
	cs.SetComputed(append(path, "workspace_status_message")...)

	return cs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Workspace.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Workspace_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"azure_workspace_info":       reflect.TypeOf(AzureWorkspaceInfo_SdkV2{}),
		"cloud_resource_container":   reflect.TypeOf(CloudResourceContainer_SdkV2{}),
		"custom_tags":                reflect.TypeOf(types.String{}),
		"external_customer_info":     reflect.TypeOf(ExternalCustomerInfo_SdkV2{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig_SdkV2{}),
		"gke_config":                 reflect.TypeOf(GkeConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Workspace_SdkV2
// only implements ToObjectValue() and Type().
func (o Workspace_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Workspace_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"aws_region": types.StringType,
			"azure_workspace_info": basetypes.ListType{
				ElemType: AzureWorkspaceInfo_SdkV2{}.Type(ctx),
			},
			"cloud": types.StringType,
			"cloud_resource_container": basetypes.ListType{
				ElemType: CloudResourceContainer_SdkV2{}.Type(ctx),
			},
			"creation_time":  types.Int64Type,
			"credentials_id": types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"deployment_name": types.StringType,
			"external_customer_info": basetypes.ListType{
				ElemType: ExternalCustomerInfo_SdkV2{}.Type(ctx),
			},
			"gcp_managed_network_config": basetypes.ListType{
				ElemType: GcpManagedNetworkConfig_SdkV2{}.Type(ctx),
			},
			"gke_config": basetypes.ListType{
				ElemType: GkeConfig_SdkV2{}.Type(ctx),
			},
			"is_no_public_ip_enabled": types.BoolType,
			"location":                types.StringType,
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

// GetAzureWorkspaceInfo returns the value of the AzureWorkspaceInfo field in Workspace_SdkV2 as
// a AzureWorkspaceInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetAzureWorkspaceInfo(ctx context.Context) (AzureWorkspaceInfo_SdkV2, bool) {
	var e AzureWorkspaceInfo_SdkV2
	if o.AzureWorkspaceInfo.IsNull() || o.AzureWorkspaceInfo.IsUnknown() {
		return e, false
	}
	var v []AzureWorkspaceInfo_SdkV2
	d := o.AzureWorkspaceInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureWorkspaceInfo sets the value of the AzureWorkspaceInfo field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetAzureWorkspaceInfo(ctx context.Context, v AzureWorkspaceInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_workspace_info"]
	o.AzureWorkspaceInfo = types.ListValueMust(t, vs)
}

// GetCloudResourceContainer returns the value of the CloudResourceContainer field in Workspace_SdkV2 as
// a CloudResourceContainer_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer_SdkV2, bool) {
	var e CloudResourceContainer_SdkV2
	if o.CloudResourceContainer.IsNull() || o.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v []CloudResourceContainer_SdkV2
	d := o.CloudResourceContainer.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cloud_resource_container"]
	o.CloudResourceContainer = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in Workspace_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
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

// SetCustomTags sets the value of the CustomTags field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CustomTags = types.MapValueMust(t, vs)
}

// GetExternalCustomerInfo returns the value of the ExternalCustomerInfo field in Workspace_SdkV2 as
// a ExternalCustomerInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetExternalCustomerInfo(ctx context.Context) (ExternalCustomerInfo_SdkV2, bool) {
	var e ExternalCustomerInfo_SdkV2
	if o.ExternalCustomerInfo.IsNull() || o.ExternalCustomerInfo.IsUnknown() {
		return e, false
	}
	var v []ExternalCustomerInfo_SdkV2
	d := o.ExternalCustomerInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExternalCustomerInfo sets the value of the ExternalCustomerInfo field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetExternalCustomerInfo(ctx context.Context, v ExternalCustomerInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["external_customer_info"]
	o.ExternalCustomerInfo = types.ListValueMust(t, vs)
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in Workspace_SdkV2 as
// a GcpManagedNetworkConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig_SdkV2, bool) {
	var e GcpManagedNetworkConfig_SdkV2
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v []GcpManagedNetworkConfig_SdkV2
	d := o.GcpManagedNetworkConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gcp_managed_network_config"]
	o.GcpManagedNetworkConfig = types.ListValueMust(t, vs)
}

// GetGkeConfig returns the value of the GkeConfig field in Workspace_SdkV2 as
// a GkeConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace_SdkV2) GetGkeConfig(ctx context.Context) (GkeConfig_SdkV2, bool) {
	var e GkeConfig_SdkV2
	if o.GkeConfig.IsNull() || o.GkeConfig.IsUnknown() {
		return e, false
	}
	var v []GkeConfig_SdkV2
	d := o.GkeConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGkeConfig sets the value of the GkeConfig field in Workspace_SdkV2.
func (o *Workspace_SdkV2) SetGkeConfig(ctx context.Context, v GkeConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["gke_config"]
	o.GkeConfig = types.ListValueMust(t, vs)
}
