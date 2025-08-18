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

type AwsCredentials struct {
	StsRole types.Object `tfsdk:"sts_role"`
}

func (toState *AwsCredentials) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AwsCredentials) {
	if !fromPlan.StsRole.IsNull() && !fromPlan.StsRole.IsUnknown() {
		if toStateStsRole, ok := toState.GetStsRole(ctx); ok {
			if fromPlanStsRole, ok := fromPlan.GetStsRole(ctx); ok {
				toStateStsRole.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanStsRole)
				toState.SetStsRole(ctx, toStateStsRole)
			}
		}
	}
}

func (toState *AwsCredentials) SyncFieldsDuringRead(ctx context.Context, fromState AwsCredentials) {
	if !fromState.StsRole.IsNull() && !fromState.StsRole.IsUnknown() {
		if toStateStsRole, ok := toState.GetStsRole(ctx); ok {
			if fromStateStsRole, ok := fromState.GetStsRole(ctx); ok {
				toStateStsRole.SyncFieldsDuringRead(ctx, fromStateStsRole)
				toState.SetStsRole(ctx, toStateStsRole)
			}
		}
	}
}

func (c AwsCredentials) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sts_role"] = attrs["sts_role"].SetOptional()

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
			"sts_role": StsRole{}.Type(ctx),
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
	var v StsRole
	d := o.StsRole.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStsRole sets the value of the StsRole field in AwsCredentials.
func (o *AwsCredentials) SetStsRole(ctx context.Context, v StsRole) {
	vs := v.ToObjectValue(ctx)
	o.StsRole = vs
}

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias types.String `tfsdk:"key_alias"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn types.String `tfsdk:"key_arn"`
	// The AWS KMS key region.
	KeyRegion types.String `tfsdk:"key_region"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	ReuseKeyForClusterVolumes types.Bool `tfsdk:"reuse_key_for_cluster_volumes"`
}

func (toState *AwsKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AwsKeyInfo) {
}

func (toState *AwsKeyInfo) SyncFieldsDuringRead(ctx context.Context, fromState AwsKeyInfo) {
}

func (c AwsKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key_alias"] = attrs["key_alias"].SetOptional()
	attrs["key_arn"] = attrs["key_arn"].SetRequired()
	attrs["key_region"] = attrs["key_region"].SetRequired()
	attrs["reuse_key_for_cluster_volumes"] = attrs["reuse_key_for_cluster_volumes"].SetOptional()

	return attrs
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
	ResourceGroup types.String `tfsdk:"resource_group"`
	// Azure Subscription ID
	SubscriptionId types.String `tfsdk:"subscription_id"`
}

func (toState *AzureWorkspaceInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan AzureWorkspaceInfo) {
}

func (toState *AzureWorkspaceInfo) SyncFieldsDuringRead(ctx context.Context, fromState AzureWorkspaceInfo) {
}

func (c AzureWorkspaceInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["resource_group"] = attrs["resource_group"].SetOptional()
	attrs["subscription_id"] = attrs["subscription_id"].SetOptional()

	return attrs
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
	Gcp types.Object `tfsdk:"gcp"`
}

func (toState *CloudResourceContainer) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CloudResourceContainer) {
	if !fromPlan.Gcp.IsNull() && !fromPlan.Gcp.IsUnknown() {
		if toStateGcp, ok := toState.GetGcp(ctx); ok {
			if fromPlanGcp, ok := fromPlan.GetGcp(ctx); ok {
				toStateGcp.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGcp)
				toState.SetGcp(ctx, toStateGcp)
			}
		}
	}
}

func (toState *CloudResourceContainer) SyncFieldsDuringRead(ctx context.Context, fromState CloudResourceContainer) {
	if !fromState.Gcp.IsNull() && !fromState.Gcp.IsUnknown() {
		if toStateGcp, ok := toState.GetGcp(ctx); ok {
			if fromStateGcp, ok := fromState.GetGcp(ctx); ok {
				toStateGcp.SyncFieldsDuringRead(ctx, fromStateGcp)
				toState.SetGcp(ctx, toStateGcp)
			}
		}
	}
}

func (c CloudResourceContainer) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp"] = attrs["gcp"].SetOptional()

	return attrs
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
			"gcp": CustomerFacingGcpCloudResourceContainer{}.Type(ctx),
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
	var v CustomerFacingGcpCloudResourceContainer
	d := o.Gcp.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcp sets the value of the Gcp field in CloudResourceContainer.
func (o *CloudResourceContainer) SetGcp(ctx context.Context, v CustomerFacingGcpCloudResourceContainer) {
	vs := v.ToObjectValue(ctx)
	o.Gcp = vs
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias types.String `tfsdk:"key_alias"`
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	KeyArn types.String `tfsdk:"key_arn"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
	ReuseKeyForClusterVolumes types.Bool `tfsdk:"reuse_key_for_cluster_volumes"`
}

func (toState *CreateAwsKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateAwsKeyInfo) {
}

func (toState *CreateAwsKeyInfo) SyncFieldsDuringRead(ctx context.Context, fromState CreateAwsKeyInfo) {
}

func (c CreateAwsKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key_alias"] = attrs["key_alias"].SetOptional()
	attrs["key_arn"] = attrs["key_arn"].SetRequired()
	attrs["reuse_key_for_cluster_volumes"] = attrs["reuse_key_for_cluster_volumes"].SetOptional()

	return attrs
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
	StsRole types.Object `tfsdk:"sts_role"`
}

func (toState *CreateCredentialAwsCredentials) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCredentialAwsCredentials) {
	if !fromPlan.StsRole.IsNull() && !fromPlan.StsRole.IsUnknown() {
		if toStateStsRole, ok := toState.GetStsRole(ctx); ok {
			if fromPlanStsRole, ok := fromPlan.GetStsRole(ctx); ok {
				toStateStsRole.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanStsRole)
				toState.SetStsRole(ctx, toStateStsRole)
			}
		}
	}
}

func (toState *CreateCredentialAwsCredentials) SyncFieldsDuringRead(ctx context.Context, fromState CreateCredentialAwsCredentials) {
	if !fromState.StsRole.IsNull() && !fromState.StsRole.IsUnknown() {
		if toStateStsRole, ok := toState.GetStsRole(ctx); ok {
			if fromStateStsRole, ok := fromState.GetStsRole(ctx); ok {
				toStateStsRole.SyncFieldsDuringRead(ctx, fromStateStsRole)
				toState.SetStsRole(ctx, toStateStsRole)
			}
		}
	}
}

func (c CreateCredentialAwsCredentials) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["sts_role"] = attrs["sts_role"].SetOptional()

	return attrs
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
			"sts_role": CreateCredentialStsRole{}.Type(ctx),
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
	var v CreateCredentialStsRole
	d := o.StsRole.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStsRole sets the value of the StsRole field in CreateCredentialAwsCredentials.
func (o *CreateCredentialAwsCredentials) SetStsRole(ctx context.Context, v CreateCredentialStsRole) {
	vs := v.ToObjectValue(ctx)
	o.StsRole = vs
}

type CreateCredentialRequest struct {
	AwsCredentials types.Object `tfsdk:"aws_credentials"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name"`
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
			"aws_credentials":  CreateCredentialAwsCredentials{}.Type(ctx),
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
	var v CreateCredentialAwsCredentials
	d := o.AwsCredentials.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsCredentials sets the value of the AwsCredentials field in CreateCredentialRequest.
func (o *CreateCredentialRequest) SetAwsCredentials(ctx context.Context, v CreateCredentialAwsCredentials) {
	vs := v.ToObjectValue(ctx)
	o.AwsCredentials = vs
}

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn types.String `tfsdk:"role_arn"`
}

func (toState *CreateCredentialStsRole) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCredentialStsRole) {
}

func (toState *CreateCredentialStsRole) SyncFieldsDuringRead(ctx context.Context, fromState CreateCredentialStsRole) {
}

func (c CreateCredentialStsRole) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["role_arn"] = attrs["role_arn"].SetOptional()

	return attrs
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
	AwsKeyInfo types.Object `tfsdk:"aws_key_info"`

	GcpKeyInfo types.Object `tfsdk:"gcp_key_info"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases"`
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
			"aws_key_info": CreateAwsKeyInfo{}.Type(ctx),
			"gcp_key_info": CreateGcpKeyInfo{}.Type(ctx),
			"use_cases": basetypes.ListType{
				ElemType: types.StringType,
			},
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
	var v CreateAwsKeyInfo
	d := o.AwsKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CreateCustomerManagedKeyRequest.
func (o *CreateCustomerManagedKeyRequest) SetAwsKeyInfo(ctx context.Context, v CreateAwsKeyInfo) {
	vs := v.ToObjectValue(ctx)
	o.AwsKeyInfo = vs
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest as
// a CreateGcpKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomerManagedKeyRequest) GetGcpKeyInfo(ctx context.Context) (CreateGcpKeyInfo, bool) {
	var e CreateGcpKeyInfo
	if o.GcpKeyInfo.IsNull() || o.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v CreateGcpKeyInfo
	d := o.GcpKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest.
func (o *CreateCustomerManagedKeyRequest) SetGcpKeyInfo(ctx context.Context, v CreateGcpKeyInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpKeyInfo = vs
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
	KmsKeyId types.String `tfsdk:"kms_key_id"`
}

func (toState *CreateGcpKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateGcpKeyInfo) {
}

func (toState *CreateGcpKeyInfo) SyncFieldsDuringRead(ctx context.Context, fromState CreateGcpKeyInfo) {
}

func (c CreateGcpKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["kms_key_id"] = attrs["kms_key_id"].SetRequired()

	return attrs
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
	GcpNetworkInfo types.Object `tfsdk:"gcp_network_info"`
	// The human-readable name of the network configuration.
	NetworkName types.String `tfsdk:"network_name"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds types.List `tfsdk:"security_group_ids"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds types.List `tfsdk:"subnet_ids"`

	VpcEndpoints types.Object `tfsdk:"vpc_endpoints"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId types.String `tfsdk:"vpc_id"`
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
			"gcp_network_info": GcpNetworkInfo{}.Type(ctx),
			"network_name":     types.StringType,
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnet_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"vpc_endpoints": NetworkVpcEndpoints{}.Type(ctx),
			"vpc_id":        types.StringType,
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
	var v GcpNetworkInfo
	d := o.GcpNetworkInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in CreateNetworkRequest.
func (o *CreateNetworkRequest) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpNetworkInfo = vs
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
	var v NetworkVpcEndpoints
	d := o.VpcEndpoints.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in CreateNetworkRequest.
func (o *CreateNetworkRequest) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints) {
	vs := v.ToObjectValue(ctx)
	o.VpcEndpoints = vs
}

type CreatePrivateAccessSettingsRequest struct {
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
	AllowedVpcEndpointIds types.List `tfsdk:"allowed_vpc_endpoint_ids"`

	PrivateAccessLevel types.String `tfsdk:"private_access_level"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled"`
	// The cloud region for workspaces associated with this private access
	// settings object.
	Region types.String `tfsdk:"region"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePrivateAccessSettingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePrivateAccessSettingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateAccessSettingsRequest
// only implements ToObjectValue() and Type().
func (o CreatePrivateAccessSettingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_vpc_endpoint_ids":     o.AllowedVpcEndpointIds,
			"private_access_level":         o.PrivateAccessLevel,
			"private_access_settings_name": o.PrivateAccessSettingsName,
			"public_access_enabled":        o.PublicAccessEnabled,
			"region":                       o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePrivateAccessSettingsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_vpc_endpoint_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"private_access_level":         types.StringType,
			"private_access_settings_name": types.StringType,
			"public_access_enabled":        types.BoolType,
			"region":                       types.StringType,
		},
	}
}

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in CreatePrivateAccessSettingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePrivateAccessSettingsRequest) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
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

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in CreatePrivateAccessSettingsRequest.
func (o *CreatePrivateAccessSettingsRequest) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type CreateStorageConfigurationRequest struct {
	RootBucketInfo types.Object `tfsdk:"root_bucket_info"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName types.String `tfsdk:"storage_configuration_name"`
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
			"root_bucket_info":           RootBucketInfo{}.Type(ctx),
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
	var v RootBucketInfo
	d := o.RootBucketInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in CreateStorageConfigurationRequest.
func (o *CreateStorageConfigurationRequest) SetRootBucketInfo(ctx context.Context, v RootBucketInfo) {
	vs := v.ToObjectValue(ctx)
	o.RootBucketInfo = vs
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id"`

	GcpVpcEndpointInfo types.Object `tfsdk:"gcp_vpc_endpoint_info"`
	// The AWS region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name"`
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
			"gcp_vpc_endpoint_info": GcpVpcEndpointInfo{}.Type(ctx),
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
	var v GcpVpcEndpointInfo
	d := o.GcpVpcEndpointInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in CreateVpcEndpointRequest.
func (o *CreateVpcEndpointRequest) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpVpcEndpointInfo = vs
}

type CreateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane.
	AwsRegion types.String `tfsdk:"aws_region"`
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	Cloud types.String `tfsdk:"cloud"`

	CloudResourceContainer types.Object `tfsdk:"cloud_resource_container"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags types.Map `tfsdk:"custom_tags"`
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
	DeploymentName types.String `tfsdk:"deployment_name"`

	GcpManagedNetworkConfig types.Object `tfsdk:"gcp_managed_network_config"`

	GkeConfig types.Object `tfsdk:"gke_config"`
	// Whether no public IP is enabled for the workspace.
	IsNoPublicIpEnabled types.Bool `tfsdk:"is_no_public_ip_enabled"`
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	Location types.String `tfsdk:"location"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id"`

	NetworkId types.String `tfsdk:"network_id"`

	PricingTier types.String `tfsdk:"pricing_tier"`
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
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id"`
	// The ID of the workspace's storage configuration object.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id"`
	// The workspace's human-readable name.
	WorkspaceName types.String `tfsdk:"workspace_name"`
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
			"aws_region":               types.StringType,
			"cloud":                    types.StringType,
			"cloud_resource_container": CloudResourceContainer{}.Type(ctx),
			"credentials_id":           types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"deployment_name":                          types.StringType,
			"gcp_managed_network_config":               GcpManagedNetworkConfig{}.Type(ctx),
			"gke_config":                               GkeConfig{}.Type(ctx),
			"is_no_public_ip_enabled":                  types.BoolType,
			"location":                                 types.StringType,
			"managed_services_customer_managed_key_id": types.StringType,
			"network_id":                               types.StringType,
			"pricing_tier":                             types.StringType,
			"private_access_settings_id":               types.StringType,
			"storage_configuration_id":                 types.StringType,
			"storage_customer_managed_key_id":          types.StringType,
			"workspace_name":                           types.StringType,
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
	var v CloudResourceContainer
	d := o.CloudResourceContainer.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer) {
	vs := v.ToObjectValue(ctx)
	o.CloudResourceContainer = vs
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
	var v GcpManagedNetworkConfig
	d := o.GcpManagedNetworkConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig) {
	vs := v.ToObjectValue(ctx)
	o.GcpManagedNetworkConfig = vs
}

// GetGkeConfig returns the value of the GkeConfig field in CreateWorkspaceRequest as
// a GkeConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWorkspaceRequest) GetGkeConfig(ctx context.Context) (GkeConfig, bool) {
	var e GkeConfig
	if o.GkeConfig.IsNull() || o.GkeConfig.IsUnknown() {
		return e, false
	}
	var v GkeConfig
	d := o.GkeConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGkeConfig sets the value of the GkeConfig field in CreateWorkspaceRequest.
func (o *CreateWorkspaceRequest) SetGkeConfig(ctx context.Context, v GkeConfig) {
	vs := v.ToObjectValue(ctx)
	o.GkeConfig = vs
}

type Credential struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id"`

	AwsCredentials types.Object `tfsdk:"aws_credentials"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Databricks credential configuration ID.
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name"`
}

func (toState *Credential) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Credential) {
	if !fromPlan.AwsCredentials.IsNull() && !fromPlan.AwsCredentials.IsUnknown() {
		if toStateAwsCredentials, ok := toState.GetAwsCredentials(ctx); ok {
			if fromPlanAwsCredentials, ok := fromPlan.GetAwsCredentials(ctx); ok {
				toStateAwsCredentials.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAwsCredentials)
				toState.SetAwsCredentials(ctx, toStateAwsCredentials)
			}
		}
	}
}

func (toState *Credential) SyncFieldsDuringRead(ctx context.Context, fromState Credential) {
	if !fromState.AwsCredentials.IsNull() && !fromState.AwsCredentials.IsUnknown() {
		if toStateAwsCredentials, ok := toState.GetAwsCredentials(ctx); ok {
			if fromStateAwsCredentials, ok := fromState.GetAwsCredentials(ctx); ok {
				toStateAwsCredentials.SyncFieldsDuringRead(ctx, fromStateAwsCredentials)
				toState.SetAwsCredentials(ctx, toStateAwsCredentials)
			}
		}
	}
}

func (c Credential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_credentials"] = attrs["aws_credentials"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["credentials_id"] = attrs["credentials_id"].SetOptional()
	attrs["credentials_name"] = attrs["credentials_name"].SetOptional()

	return attrs
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
			"aws_credentials":  AwsCredentials{}.Type(ctx),
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
	var v AwsCredentials
	d := o.AwsCredentials.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsCredentials sets the value of the AwsCredentials field in Credential.
func (o *Credential) SetAwsCredentials(ctx context.Context, v AwsCredentials) {
	vs := v.ToObjectValue(ctx)
	o.AwsCredentials = vs
}

// The general workspace configurations that are specific to Google Cloud.
type CustomerFacingGcpCloudResourceContainer struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	ProjectId types.String `tfsdk:"project_id"`
}

func (toState *CustomerFacingGcpCloudResourceContainer) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CustomerFacingGcpCloudResourceContainer) {
}

func (toState *CustomerFacingGcpCloudResourceContainer) SyncFieldsDuringRead(ctx context.Context, fromState CustomerFacingGcpCloudResourceContainer) {
}

func (c CustomerFacingGcpCloudResourceContainer) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetOptional()

	return attrs
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
	AccountId types.String `tfsdk:"account_id"`

	AwsKeyInfo types.Object `tfsdk:"aws_key_info"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId types.String `tfsdk:"customer_managed_key_id"`

	GcpKeyInfo types.Object `tfsdk:"gcp_key_info"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases"`
}

func (toState *CustomerManagedKey) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CustomerManagedKey) {
	if !fromPlan.AwsKeyInfo.IsNull() && !fromPlan.AwsKeyInfo.IsUnknown() {
		if toStateAwsKeyInfo, ok := toState.GetAwsKeyInfo(ctx); ok {
			if fromPlanAwsKeyInfo, ok := fromPlan.GetAwsKeyInfo(ctx); ok {
				toStateAwsKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAwsKeyInfo)
				toState.SetAwsKeyInfo(ctx, toStateAwsKeyInfo)
			}
		}
	}
	if !fromPlan.GcpKeyInfo.IsNull() && !fromPlan.GcpKeyInfo.IsUnknown() {
		if toStateGcpKeyInfo, ok := toState.GetGcpKeyInfo(ctx); ok {
			if fromPlanGcpKeyInfo, ok := fromPlan.GetGcpKeyInfo(ctx); ok {
				toStateGcpKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGcpKeyInfo)
				toState.SetGcpKeyInfo(ctx, toStateGcpKeyInfo)
			}
		}
	}
}

func (toState *CustomerManagedKey) SyncFieldsDuringRead(ctx context.Context, fromState CustomerManagedKey) {
	if !fromState.AwsKeyInfo.IsNull() && !fromState.AwsKeyInfo.IsUnknown() {
		if toStateAwsKeyInfo, ok := toState.GetAwsKeyInfo(ctx); ok {
			if fromStateAwsKeyInfo, ok := fromState.GetAwsKeyInfo(ctx); ok {
				toStateAwsKeyInfo.SyncFieldsDuringRead(ctx, fromStateAwsKeyInfo)
				toState.SetAwsKeyInfo(ctx, toStateAwsKeyInfo)
			}
		}
	}
	if !fromState.GcpKeyInfo.IsNull() && !fromState.GcpKeyInfo.IsUnknown() {
		if toStateGcpKeyInfo, ok := toState.GetGcpKeyInfo(ctx); ok {
			if fromStateGcpKeyInfo, ok := fromState.GetGcpKeyInfo(ctx); ok {
				toStateGcpKeyInfo.SyncFieldsDuringRead(ctx, fromStateGcpKeyInfo)
				toState.SetGcpKeyInfo(ctx, toStateGcpKeyInfo)
			}
		}
	}
}

func (c CustomerManagedKey) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_key_info"] = attrs["aws_key_info"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["customer_managed_key_id"] = attrs["customer_managed_key_id"].SetOptional()
	attrs["gcp_key_info"] = attrs["gcp_key_info"].SetOptional()
	attrs["use_cases"] = attrs["use_cases"].SetOptional()

	return attrs
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
			"aws_key_info":            AwsKeyInfo{}.Type(ctx),
			"creation_time":           types.Int64Type,
			"customer_managed_key_id": types.StringType,
			"gcp_key_info":            GcpKeyInfo{}.Type(ctx),
			"use_cases": basetypes.ListType{
				ElemType: types.StringType,
			},
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
	var v AwsKeyInfo
	d := o.AwsKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CustomerManagedKey.
func (o *CustomerManagedKey) SetAwsKeyInfo(ctx context.Context, v AwsKeyInfo) {
	vs := v.ToObjectValue(ctx)
	o.AwsKeyInfo = vs
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CustomerManagedKey as
// a GcpKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerManagedKey) GetGcpKeyInfo(ctx context.Context) (GcpKeyInfo, bool) {
	var e GcpKeyInfo
	if o.GcpKeyInfo.IsNull() || o.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v GcpKeyInfo
	d := o.GcpKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CustomerManagedKey.
func (o *CustomerManagedKey) SetGcpKeyInfo(ctx context.Context, v GcpKeyInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpKeyInfo = vs
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

type DeleteCredentialRequest struct {
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

type DeleteEncryptionKeyRequest struct {
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

type DeleteNetworkRequest struct {
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

type DeletePrivateAccesRequest struct {
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

type DeleteStorageRequest struct {
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

type DeleteVpcEndpointRequest struct {
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

type DeleteWorkspaceRequest struct {
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
	AuthoritativeUserEmail types.String `tfsdk:"authoritative_user_email"`
	// The authoritative user full name.
	AuthoritativeUserFullName types.String `tfsdk:"authoritative_user_full_name"`
	// The legal entity name for the external workspace
	CustomerName types.String `tfsdk:"customer_name"`
}

func (toState *ExternalCustomerInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ExternalCustomerInfo) {
}

func (toState *ExternalCustomerInfo) SyncFieldsDuringRead(ctx context.Context, fromState ExternalCustomerInfo) {
}

func (c ExternalCustomerInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["authoritative_user_email"] = attrs["authoritative_user_email"].SetOptional()
	attrs["authoritative_user_full_name"] = attrs["authoritative_user_full_name"].SetOptional()
	attrs["customer_name"] = attrs["customer_name"].SetOptional()

	return attrs
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
	KmsKeyId types.String `tfsdk:"kms_key_id"`
}

func (toState *GcpKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GcpKeyInfo) {
}

func (toState *GcpKeyInfo) SyncFieldsDuringRead(ctx context.Context, fromState GcpKeyInfo) {
}

func (c GcpKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["kms_key_id"] = attrs["kms_key_id"].SetRequired()

	return attrs
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
	GkeClusterPodIpRange types.String `tfsdk:"gke_cluster_pod_ip_range"`
	// The IP range from which to allocate GKE cluster services. No bigger than
	// `/16` and no smaller than `/27`.
	GkeClusterServiceIpRange types.String `tfsdk:"gke_cluster_service_ip_range"`
	// The IP range from which to allocate GKE cluster nodes. No bigger than
	// `/9` and no smaller than `/29`.
	SubnetCidr types.String `tfsdk:"subnet_cidr"`
}

func (toState *GcpManagedNetworkConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GcpManagedNetworkConfig) {
}

func (toState *GcpManagedNetworkConfig) SyncFieldsDuringRead(ctx context.Context, fromState GcpManagedNetworkConfig) {
}

func (c GcpManagedNetworkConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gke_cluster_pod_ip_range"] = attrs["gke_cluster_pod_ip_range"].SetOptional()
	attrs["gke_cluster_service_ip_range"] = attrs["gke_cluster_service_ip_range"].SetOptional()
	attrs["subnet_cidr"] = attrs["subnet_cidr"].SetOptional()

	return attrs
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
	NetworkProjectId types.String `tfsdk:"network_project_id"`
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can be
	// used by only one workspace.
	PodIpRangeName types.String `tfsdk:"pod_ip_range_name"`
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// be used by only one workspace.
	ServiceIpRangeName types.String `tfsdk:"service_ip_range_name"`
	// The ID of the subnet associated with this network.
	SubnetId types.String `tfsdk:"subnet_id"`
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	SubnetRegion types.String `tfsdk:"subnet_region"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId types.String `tfsdk:"vpc_id"`
}

func (toState *GcpNetworkInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GcpNetworkInfo) {
}

func (toState *GcpNetworkInfo) SyncFieldsDuringRead(ctx context.Context, fromState GcpNetworkInfo) {
}

func (c GcpNetworkInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_project_id"] = attrs["network_project_id"].SetRequired()
	attrs["pod_ip_range_name"] = attrs["pod_ip_range_name"].SetRequired()
	attrs["service_ip_range_name"] = attrs["service_ip_range_name"].SetRequired()
	attrs["subnet_id"] = attrs["subnet_id"].SetRequired()
	attrs["subnet_region"] = attrs["subnet_region"].SetRequired()
	attrs["vpc_id"] = attrs["vpc_id"].SetRequired()

	return attrs
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
	EndpointRegion types.String `tfsdk:"endpoint_region"`
	// The Google Cloud project ID of the VPC network where the PSC connection
	// resides.
	ProjectId types.String `tfsdk:"project_id"`
	// The unique ID of this PSC connection.
	PscConnectionId types.String `tfsdk:"psc_connection_id"`
	// The name of the PSC endpoint in the Google Cloud project.
	PscEndpointName types.String `tfsdk:"psc_endpoint_name"`
	// The service attachment this PSC connection connects to.
	ServiceAttachmentId types.String `tfsdk:"service_attachment_id"`
}

func (toState *GcpVpcEndpointInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GcpVpcEndpointInfo) {
}

func (toState *GcpVpcEndpointInfo) SyncFieldsDuringRead(ctx context.Context, fromState GcpVpcEndpointInfo) {
}

func (c GcpVpcEndpointInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint_region"] = attrs["endpoint_region"].SetRequired()
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["psc_connection_id"] = attrs["psc_connection_id"].SetOptional()
	attrs["psc_endpoint_name"] = attrs["psc_endpoint_name"].SetRequired()
	attrs["service_attachment_id"] = attrs["service_attachment_id"].SetOptional()

	return attrs
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

type GetCredentialRequest struct {
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

type GetEncryptionKeyRequest struct {
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

type GetNetworkRequest struct {
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

type GetPrivateAccesRequest struct {
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

type GetStorageRequest struct {
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

type GetVpcEndpointRequest struct {
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

type GetWorkspaceRequest struct {
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
	ConnectivityType types.String `tfsdk:"connectivity_type"`
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
	MasterIpRange types.String `tfsdk:"master_ip_range"`
}

func (toState *GkeConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GkeConfig) {
}

func (toState *GkeConfig) SyncFieldsDuringRead(ctx context.Context, fromState GkeConfig) {
}

func (c GkeConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connectivity_type"] = attrs["connectivity_type"].SetOptional()
	attrs["master_ip_range"] = attrs["master_ip_range"].SetOptional()

	return attrs
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

type ListCredentialsRequest struct {
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
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListCredentialsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListEncryptionKeysRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEncryptionKeysRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListEncryptionKeysRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEncryptionKeysRequest
// only implements ToObjectValue() and Type().
func (o ListEncryptionKeysRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListEncryptionKeysRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListNetworksRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworksRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNetworksRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworksRequest
// only implements ToObjectValue() and Type().
func (o ListNetworksRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListNetworksRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListPrivateAccessRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPrivateAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPrivateAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateAccessRequest
// only implements ToObjectValue() and Type().
func (o ListPrivateAccessRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListPrivateAccessRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListStorageRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListStorageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStorageRequest
// only implements ToObjectValue() and Type().
func (o ListStorageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListStorageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListVpcEndpointsRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVpcEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListVpcEndpointsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVpcEndpointsRequest
// only implements ToObjectValue() and Type().
func (o ListVpcEndpointsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListVpcEndpointsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListWorkspacesRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspacesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspacesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspacesRequest
// only implements ToObjectValue() and Type().
func (o ListWorkspacesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspacesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Network struct {
	// The Databricks account ID associated with this network configuration.
	AccountId types.String `tfsdk:"account_id"`
	// Time in epoch milliseconds when the network was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Array of error messages about the network configuration.
	ErrorMessages types.List `tfsdk:"error_messages"`

	GcpNetworkInfo types.Object `tfsdk:"gcp_network_info"`
	// The Databricks network configuration ID.
	NetworkId types.String `tfsdk:"network_id"`
	// The human-readable name of the network configuration.
	NetworkName types.String `tfsdk:"network_name"`

	SecurityGroupIds types.List `tfsdk:"security_group_ids"`

	SubnetIds types.List `tfsdk:"subnet_ids"`

	VpcEndpoints types.Object `tfsdk:"vpc_endpoints"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId types.String `tfsdk:"vpc_id"`

	VpcStatus types.String `tfsdk:"vpc_status"`
	// Array of warning messages about the network configuration.
	WarningMessages types.List `tfsdk:"warning_messages"`
	// Workspace ID associated with this network configuration.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (toState *Network) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Network) {
	if !fromPlan.GcpNetworkInfo.IsNull() && !fromPlan.GcpNetworkInfo.IsUnknown() {
		if toStateGcpNetworkInfo, ok := toState.GetGcpNetworkInfo(ctx); ok {
			if fromPlanGcpNetworkInfo, ok := fromPlan.GetGcpNetworkInfo(ctx); ok {
				toStateGcpNetworkInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGcpNetworkInfo)
				toState.SetGcpNetworkInfo(ctx, toStateGcpNetworkInfo)
			}
		}
	}
	if !fromPlan.VpcEndpoints.IsNull() && !fromPlan.VpcEndpoints.IsUnknown() {
		if toStateVpcEndpoints, ok := toState.GetVpcEndpoints(ctx); ok {
			if fromPlanVpcEndpoints, ok := fromPlan.GetVpcEndpoints(ctx); ok {
				toStateVpcEndpoints.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanVpcEndpoints)
				toState.SetVpcEndpoints(ctx, toStateVpcEndpoints)
			}
		}
	}
}

func (toState *Network) SyncFieldsDuringRead(ctx context.Context, fromState Network) {
	if !fromState.GcpNetworkInfo.IsNull() && !fromState.GcpNetworkInfo.IsUnknown() {
		if toStateGcpNetworkInfo, ok := toState.GetGcpNetworkInfo(ctx); ok {
			if fromStateGcpNetworkInfo, ok := fromState.GetGcpNetworkInfo(ctx); ok {
				toStateGcpNetworkInfo.SyncFieldsDuringRead(ctx, fromStateGcpNetworkInfo)
				toState.SetGcpNetworkInfo(ctx, toStateGcpNetworkInfo)
			}
		}
	}
	if !fromState.VpcEndpoints.IsNull() && !fromState.VpcEndpoints.IsUnknown() {
		if toStateVpcEndpoints, ok := toState.GetVpcEndpoints(ctx); ok {
			if fromStateVpcEndpoints, ok := fromState.GetVpcEndpoints(ctx); ok {
				toStateVpcEndpoints.SyncFieldsDuringRead(ctx, fromStateVpcEndpoints)
				toState.SetVpcEndpoints(ctx, toStateVpcEndpoints)
			}
		}
	}
}

func (c Network) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["error_messages"] = attrs["error_messages"].SetComputed()
	attrs["gcp_network_info"] = attrs["gcp_network_info"].SetOptional()
	attrs["network_id"] = attrs["network_id"].SetOptional()
	attrs["network_name"] = attrs["network_name"].SetOptional()
	attrs["security_group_ids"] = attrs["security_group_ids"].SetOptional()
	attrs["subnet_ids"] = attrs["subnet_ids"].SetOptional()
	attrs["vpc_endpoints"] = attrs["vpc_endpoints"].SetComputed()
	attrs["vpc_id"] = attrs["vpc_id"].SetOptional()
	attrs["vpc_status"] = attrs["vpc_status"].SetComputed()
	attrs["warning_messages"] = attrs["warning_messages"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()

	return attrs
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
			"account_id":    types.StringType,
			"creation_time": types.Int64Type,
			"error_messages": basetypes.ListType{
				ElemType: NetworkHealth{}.Type(ctx),
			},
			"gcp_network_info": GcpNetworkInfo{}.Type(ctx),
			"network_id":       types.StringType,
			"network_name":     types.StringType,
			"security_group_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"subnet_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"vpc_endpoints": NetworkVpcEndpoints{}.Type(ctx),
			"vpc_id":        types.StringType,
			"vpc_status":    types.StringType,
			"warning_messages": basetypes.ListType{
				ElemType: NetworkWarning{}.Type(ctx),
			},
			"workspace_id": types.Int64Type,
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
	var v GcpNetworkInfo
	d := o.GcpNetworkInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in Network.
func (o *Network) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpNetworkInfo = vs
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
	var v NetworkVpcEndpoints
	d := o.VpcEndpoints.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in Network.
func (o *Network) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints) {
	vs := v.ToObjectValue(ctx)
	o.VpcEndpoints = vs
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
	ErrorMessage types.String `tfsdk:"error_message"`

	ErrorType types.String `tfsdk:"error_type"`
}

func (toState *NetworkHealth) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan NetworkHealth) {
}

func (toState *NetworkHealth) SyncFieldsDuringRead(ctx context.Context, fromState NetworkHealth) {
}

func (c NetworkHealth) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["error_message"] = attrs["error_message"].SetOptional()
	attrs["error_type"] = attrs["error_type"].SetOptional()

	return attrs
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
	DataplaneRelay types.List `tfsdk:"dataplane_relay"`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi types.List `tfsdk:"rest_api"`
}

func (toState *NetworkVpcEndpoints) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan NetworkVpcEndpoints) {
}

func (toState *NetworkVpcEndpoints) SyncFieldsDuringRead(ctx context.Context, fromState NetworkVpcEndpoints) {
}

func (c NetworkVpcEndpoints) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dataplane_relay"] = attrs["dataplane_relay"].SetRequired()
	attrs["rest_api"] = attrs["rest_api"].SetRequired()

	return attrs
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
			"dataplane_relay": basetypes.ListType{
				ElemType: types.StringType,
			},
			"rest_api": basetypes.ListType{
				ElemType: types.StringType,
			},
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
	WarningMessage types.String `tfsdk:"warning_message"`

	WarningType types.String `tfsdk:"warning_type"`
}

func (toState *NetworkWarning) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan NetworkWarning) {
}

func (toState *NetworkWarning) SyncFieldsDuringRead(ctx context.Context, fromState NetworkWarning) {
}

func (c NetworkWarning) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["warning_message"] = attrs["warning_message"].SetOptional()
	attrs["warning_type"] = attrs["warning_type"].SetOptional()

	return attrs
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
	AccountId types.String `tfsdk:"account_id"`
	// An array of Databricks VPC endpoint IDs.
	AllowedVpcEndpointIds types.List `tfsdk:"allowed_vpc_endpoint_ids"`

	PrivateAccessLevel types.String `tfsdk:"private_access_level"`
	// Databricks private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled"`
	// The cloud region for workspaces attached to this private access settings
	// object.
	Region types.String `tfsdk:"region"`
}

func (toState *PrivateAccessSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PrivateAccessSettings) {
}

func (toState *PrivateAccessSettings) SyncFieldsDuringRead(ctx context.Context, fromState PrivateAccessSettings) {
}

func (c PrivateAccessSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["allowed_vpc_endpoint_ids"] = attrs["allowed_vpc_endpoint_ids"].SetOptional()
	attrs["private_access_level"] = attrs["private_access_level"].SetOptional()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetOptional()
	attrs["private_access_settings_name"] = attrs["private_access_settings_name"].SetOptional()
	attrs["public_access_enabled"] = attrs["public_access_enabled"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()

	return attrs
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

type ReplacePrivateAccessSettingsRequest struct {
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
	AllowedVpcEndpointIds types.List `tfsdk:"allowed_vpc_endpoint_ids"`

	PrivateAccessLevel types.String `tfsdk:"private_access_level"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled"`
	// The cloud region for workspaces associated with this private access
	// settings object.
	Region types.String `tfsdk:"region"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplacePrivateAccessSettingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReplacePrivateAccessSettingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplacePrivateAccessSettingsRequest
// only implements ToObjectValue() and Type().
func (o ReplacePrivateAccessSettingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ReplacePrivateAccessSettingsRequest) Type(ctx context.Context) attr.Type {
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

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in ReplacePrivateAccessSettingsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ReplacePrivateAccessSettingsRequest) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
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

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in ReplacePrivateAccessSettingsRequest.
func (o *ReplacePrivateAccessSettingsRequest) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
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
	BucketName types.String `tfsdk:"bucket_name"`
}

func (toState *RootBucketInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RootBucketInfo) {
}

func (toState *RootBucketInfo) SyncFieldsDuringRead(ctx context.Context, fromState RootBucketInfo) {
}

func (c RootBucketInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bucket_name"] = attrs["bucket_name"].SetOptional()

	return attrs
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
	AccountId types.String `tfsdk:"account_id"`
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`

	RootBucketInfo types.Object `tfsdk:"root_bucket_info"`
	// Databricks storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName types.String `tfsdk:"storage_configuration_name"`
}

func (toState *StorageConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan StorageConfiguration) {
	if !fromPlan.RootBucketInfo.IsNull() && !fromPlan.RootBucketInfo.IsUnknown() {
		if toStateRootBucketInfo, ok := toState.GetRootBucketInfo(ctx); ok {
			if fromPlanRootBucketInfo, ok := fromPlan.GetRootBucketInfo(ctx); ok {
				toStateRootBucketInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRootBucketInfo)
				toState.SetRootBucketInfo(ctx, toStateRootBucketInfo)
			}
		}
	}
}

func (toState *StorageConfiguration) SyncFieldsDuringRead(ctx context.Context, fromState StorageConfiguration) {
	if !fromState.RootBucketInfo.IsNull() && !fromState.RootBucketInfo.IsUnknown() {
		if toStateRootBucketInfo, ok := toState.GetRootBucketInfo(ctx); ok {
			if fromStateRootBucketInfo, ok := fromState.GetRootBucketInfo(ctx); ok {
				toStateRootBucketInfo.SyncFieldsDuringRead(ctx, fromStateRootBucketInfo)
				toState.SetRootBucketInfo(ctx, toStateRootBucketInfo)
			}
		}
	}
}

func (c StorageConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["root_bucket_info"] = attrs["root_bucket_info"].SetOptional()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetOptional()
	attrs["storage_configuration_name"] = attrs["storage_configuration_name"].SetOptional()

	return attrs
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
			"root_bucket_info":           RootBucketInfo{}.Type(ctx),
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
	var v RootBucketInfo
	d := o.RootBucketInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in StorageConfiguration.
func (o *StorageConfiguration) SetRootBucketInfo(ctx context.Context, v RootBucketInfo) {
	vs := v.ToObjectValue(ctx)
	o.RootBucketInfo = vs
}

type StsRole struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	ExternalId types.String `tfsdk:"external_id"`
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn types.String `tfsdk:"role_arn"`
}

func (toState *StsRole) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan StsRole) {
}

func (toState *StsRole) SyncFieldsDuringRead(ctx context.Context, fromState StsRole) {
}

func (c StsRole) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetOptional()
	attrs["role_arn"] = attrs["role_arn"].SetOptional()

	return attrs
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
	AwsRegion types.String `tfsdk:"aws_region"`
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags types.Map `tfsdk:"custom_tags"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id"`

	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id"`
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. For failed workspaces only, you can
	// switch from a Databricks-managed VPC to a customer-managed VPC by
	// updating the workspace to add a network configuration ID.
	NetworkId types.String `tfsdk:"network_id"`
	// The ID of the workspace's private access settings configuration object.
	// This parameter is available only for updating failed workspaces.
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id"`
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
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

type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
	AccountId types.String `tfsdk:"account_id"`
	// The AWS Account in which the VPC endpoint object exists.
	AwsAccountId types.String `tfsdk:"aws_account_id"`
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	AwsEndpointServiceId types.String `tfsdk:"aws_endpoint_service_id"`
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id"`

	GcpVpcEndpointInfo types.Object `tfsdk:"gcp_vpc_endpoint_info"`
	// The AWS region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region"`
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	State types.String `tfsdk:"state"`

	UseCase types.String `tfsdk:"use_case"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId types.String `tfsdk:"vpc_endpoint_id"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name"`
}

func (toState *VpcEndpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan VpcEndpoint) {
	if !fromPlan.GcpVpcEndpointInfo.IsNull() && !fromPlan.GcpVpcEndpointInfo.IsUnknown() {
		if toStateGcpVpcEndpointInfo, ok := toState.GetGcpVpcEndpointInfo(ctx); ok {
			if fromPlanGcpVpcEndpointInfo, ok := fromPlan.GetGcpVpcEndpointInfo(ctx); ok {
				toStateGcpVpcEndpointInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGcpVpcEndpointInfo)
				toState.SetGcpVpcEndpointInfo(ctx, toStateGcpVpcEndpointInfo)
			}
		}
	}
}

func (toState *VpcEndpoint) SyncFieldsDuringRead(ctx context.Context, fromState VpcEndpoint) {
	if !fromState.GcpVpcEndpointInfo.IsNull() && !fromState.GcpVpcEndpointInfo.IsUnknown() {
		if toStateGcpVpcEndpointInfo, ok := toState.GetGcpVpcEndpointInfo(ctx); ok {
			if fromStateGcpVpcEndpointInfo, ok := fromState.GetGcpVpcEndpointInfo(ctx); ok {
				toStateGcpVpcEndpointInfo.SyncFieldsDuringRead(ctx, fromStateGcpVpcEndpointInfo)
				toState.SetGcpVpcEndpointInfo(ctx, toStateGcpVpcEndpointInfo)
			}
		}
	}
}

func (c VpcEndpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_account_id"] = attrs["aws_account_id"].SetOptional()
	attrs["aws_endpoint_service_id"] = attrs["aws_endpoint_service_id"].SetOptional()
	attrs["aws_vpc_endpoint_id"] = attrs["aws_vpc_endpoint_id"].SetOptional()
	attrs["gcp_vpc_endpoint_info"] = attrs["gcp_vpc_endpoint_info"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["use_case"] = attrs["use_case"].SetOptional()
	attrs["vpc_endpoint_id"] = attrs["vpc_endpoint_id"].SetOptional()
	attrs["vpc_endpoint_name"] = attrs["vpc_endpoint_name"].SetOptional()

	return attrs
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
			"gcp_vpc_endpoint_info":   GcpVpcEndpointInfo{}.Type(ctx),
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
	var v GcpVpcEndpointInfo
	d := o.GcpVpcEndpointInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in VpcEndpoint.
func (o *VpcEndpoint) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo) {
	vs := v.ToObjectValue(ctx)
	o.GcpVpcEndpointInfo = vs
}

type Workspace struct {
	// Databricks account ID.
	AccountId types.String `tfsdk:"account_id"`
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	AwsRegion types.String `tfsdk:"aws_region"`

	AzureWorkspaceInfo types.Object `tfsdk:"azure_workspace_info"`
	// The cloud name. This field always has the value `gcp`.
	Cloud types.String `tfsdk:"cloud"`

	CloudResourceContainer types.Object `tfsdk:"cloud_resource_container"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags types.Map `tfsdk:"custom_tags"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	DeploymentName types.String `tfsdk:"deployment_name"`
	// If this workspace is for a external customer, then external_customer_info
	// is populated. If this workspace is not for a external customer, then
	// external_customer_info is empty.
	ExternalCustomerInfo types.Object `tfsdk:"external_customer_info"`

	GcpManagedNetworkConfig types.Object `tfsdk:"gcp_managed_network_config"`

	GkeConfig types.Object `tfsdk:"gke_config"`
	// Whether no public IP is enabled for the workspace.
	IsNoPublicIpEnabled types.Bool `tfsdk:"is_no_public_ip_enabled"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location types.String `tfsdk:"location"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id"`
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	NetworkId types.String `tfsdk:"network_id"`

	PricingTier types.String `tfsdk:"pricing_tier"`
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
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id"`
	// A unique integer ID for the workspace
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
	// The human-readable name of the workspace.
	WorkspaceName types.String `tfsdk:"workspace_name"`

	WorkspaceStatus types.String `tfsdk:"workspace_status"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage types.String `tfsdk:"workspace_status_message"`
}

func (toState *Workspace) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Workspace) {
	if !fromPlan.AzureWorkspaceInfo.IsNull() && !fromPlan.AzureWorkspaceInfo.IsUnknown() {
		if toStateAzureWorkspaceInfo, ok := toState.GetAzureWorkspaceInfo(ctx); ok {
			if fromPlanAzureWorkspaceInfo, ok := fromPlan.GetAzureWorkspaceInfo(ctx); ok {
				toStateAzureWorkspaceInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAzureWorkspaceInfo)
				toState.SetAzureWorkspaceInfo(ctx, toStateAzureWorkspaceInfo)
			}
		}
	}
	if !fromPlan.CloudResourceContainer.IsNull() && !fromPlan.CloudResourceContainer.IsUnknown() {
		if toStateCloudResourceContainer, ok := toState.GetCloudResourceContainer(ctx); ok {
			if fromPlanCloudResourceContainer, ok := fromPlan.GetCloudResourceContainer(ctx); ok {
				toStateCloudResourceContainer.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanCloudResourceContainer)
				toState.SetCloudResourceContainer(ctx, toStateCloudResourceContainer)
			}
		}
	}
	if !fromPlan.ExternalCustomerInfo.IsNull() && !fromPlan.ExternalCustomerInfo.IsUnknown() {
		if toStateExternalCustomerInfo, ok := toState.GetExternalCustomerInfo(ctx); ok {
			if fromPlanExternalCustomerInfo, ok := fromPlan.GetExternalCustomerInfo(ctx); ok {
				toStateExternalCustomerInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanExternalCustomerInfo)
				toState.SetExternalCustomerInfo(ctx, toStateExternalCustomerInfo)
			}
		}
	}
	if !fromPlan.GcpManagedNetworkConfig.IsNull() && !fromPlan.GcpManagedNetworkConfig.IsUnknown() {
		if toStateGcpManagedNetworkConfig, ok := toState.GetGcpManagedNetworkConfig(ctx); ok {
			if fromPlanGcpManagedNetworkConfig, ok := fromPlan.GetGcpManagedNetworkConfig(ctx); ok {
				toStateGcpManagedNetworkConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGcpManagedNetworkConfig)
				toState.SetGcpManagedNetworkConfig(ctx, toStateGcpManagedNetworkConfig)
			}
		}
	}
	if !fromPlan.GkeConfig.IsNull() && !fromPlan.GkeConfig.IsUnknown() {
		if toStateGkeConfig, ok := toState.GetGkeConfig(ctx); ok {
			if fromPlanGkeConfig, ok := fromPlan.GetGkeConfig(ctx); ok {
				toStateGkeConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanGkeConfig)
				toState.SetGkeConfig(ctx, toStateGkeConfig)
			}
		}
	}
}

func (toState *Workspace) SyncFieldsDuringRead(ctx context.Context, fromState Workspace) {
	if !fromState.AzureWorkspaceInfo.IsNull() && !fromState.AzureWorkspaceInfo.IsUnknown() {
		if toStateAzureWorkspaceInfo, ok := toState.GetAzureWorkspaceInfo(ctx); ok {
			if fromStateAzureWorkspaceInfo, ok := fromState.GetAzureWorkspaceInfo(ctx); ok {
				toStateAzureWorkspaceInfo.SyncFieldsDuringRead(ctx, fromStateAzureWorkspaceInfo)
				toState.SetAzureWorkspaceInfo(ctx, toStateAzureWorkspaceInfo)
			}
		}
	}
	if !fromState.CloudResourceContainer.IsNull() && !fromState.CloudResourceContainer.IsUnknown() {
		if toStateCloudResourceContainer, ok := toState.GetCloudResourceContainer(ctx); ok {
			if fromStateCloudResourceContainer, ok := fromState.GetCloudResourceContainer(ctx); ok {
				toStateCloudResourceContainer.SyncFieldsDuringRead(ctx, fromStateCloudResourceContainer)
				toState.SetCloudResourceContainer(ctx, toStateCloudResourceContainer)
			}
		}
	}
	if !fromState.ExternalCustomerInfo.IsNull() && !fromState.ExternalCustomerInfo.IsUnknown() {
		if toStateExternalCustomerInfo, ok := toState.GetExternalCustomerInfo(ctx); ok {
			if fromStateExternalCustomerInfo, ok := fromState.GetExternalCustomerInfo(ctx); ok {
				toStateExternalCustomerInfo.SyncFieldsDuringRead(ctx, fromStateExternalCustomerInfo)
				toState.SetExternalCustomerInfo(ctx, toStateExternalCustomerInfo)
			}
		}
	}
	if !fromState.GcpManagedNetworkConfig.IsNull() && !fromState.GcpManagedNetworkConfig.IsUnknown() {
		if toStateGcpManagedNetworkConfig, ok := toState.GetGcpManagedNetworkConfig(ctx); ok {
			if fromStateGcpManagedNetworkConfig, ok := fromState.GetGcpManagedNetworkConfig(ctx); ok {
				toStateGcpManagedNetworkConfig.SyncFieldsDuringRead(ctx, fromStateGcpManagedNetworkConfig)
				toState.SetGcpManagedNetworkConfig(ctx, toStateGcpManagedNetworkConfig)
			}
		}
	}
	if !fromState.GkeConfig.IsNull() && !fromState.GkeConfig.IsUnknown() {
		if toStateGkeConfig, ok := toState.GetGkeConfig(ctx); ok {
			if fromStateGkeConfig, ok := fromState.GetGkeConfig(ctx); ok {
				toStateGkeConfig.SyncFieldsDuringRead(ctx, fromStateGkeConfig)
				toState.SetGkeConfig(ctx, toStateGkeConfig)
			}
		}
	}
}

func (c Workspace) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["aws_region"] = attrs["aws_region"].SetOptional()
	attrs["azure_workspace_info"] = attrs["azure_workspace_info"].SetComputed()
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["cloud_resource_container"] = attrs["cloud_resource_container"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["credentials_id"] = attrs["credentials_id"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["deployment_name"] = attrs["deployment_name"].SetOptional()
	attrs["external_customer_info"] = attrs["external_customer_info"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].SetOptional()
	attrs["gke_config"] = attrs["gke_config"].SetOptional()
	attrs["is_no_public_ip_enabled"] = attrs["is_no_public_ip_enabled"].SetOptional()
	attrs["location"] = attrs["location"].SetOptional()
	attrs["managed_services_customer_managed_key_id"] = attrs["managed_services_customer_managed_key_id"].SetOptional()
	attrs["network_id"] = attrs["network_id"].SetOptional()
	attrs["pricing_tier"] = attrs["pricing_tier"].SetOptional()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetOptional()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetOptional()
	attrs["storage_customer_managed_key_id"] = attrs["storage_customer_managed_key_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_name"] = attrs["workspace_name"].SetOptional()
	attrs["workspace_status"] = attrs["workspace_status"].SetComputed()
	attrs["workspace_status_message"] = attrs["workspace_status_message"].SetComputed()

	return attrs
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
			"account_id":               types.StringType,
			"aws_region":               types.StringType,
			"azure_workspace_info":     AzureWorkspaceInfo{}.Type(ctx),
			"cloud":                    types.StringType,
			"cloud_resource_container": CloudResourceContainer{}.Type(ctx),
			"creation_time":            types.Int64Type,
			"credentials_id":           types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"deployment_name":                          types.StringType,
			"external_customer_info":                   ExternalCustomerInfo{}.Type(ctx),
			"gcp_managed_network_config":               GcpManagedNetworkConfig{}.Type(ctx),
			"gke_config":                               GkeConfig{}.Type(ctx),
			"is_no_public_ip_enabled":                  types.BoolType,
			"location":                                 types.StringType,
			"managed_services_customer_managed_key_id": types.StringType,
			"network_id":                               types.StringType,
			"pricing_tier":                             types.StringType,
			"private_access_settings_id":               types.StringType,
			"storage_configuration_id":                 types.StringType,
			"storage_customer_managed_key_id":          types.StringType,
			"workspace_id":                             types.Int64Type,
			"workspace_name":                           types.StringType,
			"workspace_status":                         types.StringType,
			"workspace_status_message":                 types.StringType,
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
	var v AzureWorkspaceInfo
	d := o.AzureWorkspaceInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureWorkspaceInfo sets the value of the AzureWorkspaceInfo field in Workspace.
func (o *Workspace) SetAzureWorkspaceInfo(ctx context.Context, v AzureWorkspaceInfo) {
	vs := v.ToObjectValue(ctx)
	o.AzureWorkspaceInfo = vs
}

// GetCloudResourceContainer returns the value of the CloudResourceContainer field in Workspace as
// a CloudResourceContainer value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer, bool) {
	var e CloudResourceContainer
	if o.CloudResourceContainer.IsNull() || o.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v CloudResourceContainer
	d := o.CloudResourceContainer.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in Workspace.
func (o *Workspace) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer) {
	vs := v.ToObjectValue(ctx)
	o.CloudResourceContainer = vs
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
	var v ExternalCustomerInfo
	d := o.ExternalCustomerInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalCustomerInfo sets the value of the ExternalCustomerInfo field in Workspace.
func (o *Workspace) SetExternalCustomerInfo(ctx context.Context, v ExternalCustomerInfo) {
	vs := v.ToObjectValue(ctx)
	o.ExternalCustomerInfo = vs
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in Workspace as
// a GcpManagedNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig, bool) {
	var e GcpManagedNetworkConfig
	if o.GcpManagedNetworkConfig.IsNull() || o.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v GcpManagedNetworkConfig
	d := o.GcpManagedNetworkConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in Workspace.
func (o *Workspace) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig) {
	vs := v.ToObjectValue(ctx)
	o.GcpManagedNetworkConfig = vs
}

// GetGkeConfig returns the value of the GkeConfig field in Workspace as
// a GkeConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *Workspace) GetGkeConfig(ctx context.Context) (GkeConfig, bool) {
	var e GkeConfig
	if o.GkeConfig.IsNull() || o.GkeConfig.IsUnknown() {
		return e, false
	}
	var v GkeConfig
	d := o.GkeConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGkeConfig sets the value of the GkeConfig field in Workspace.
func (o *Workspace) SetGkeConfig(ctx context.Context, v GkeConfig) {
	vs := v.ToObjectValue(ctx)
	o.GkeConfig = vs
}
