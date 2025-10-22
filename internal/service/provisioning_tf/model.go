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

func (to *AwsCredentials) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsCredentials) {
	if !from.StsRole.IsNull() && !from.StsRole.IsUnknown() {
		if toStsRole, ok := to.GetStsRole(ctx); ok {
			if fromStsRole, ok := from.GetStsRole(ctx); ok {
				// Recursively sync the fields of StsRole
				toStsRole.SyncFieldsDuringCreateOrUpdate(ctx, fromStsRole)
				to.SetStsRole(ctx, toStsRole)
			}
		}
	}
}

func (to *AwsCredentials) SyncFieldsDuringRead(ctx context.Context, from AwsCredentials) {
	if !from.StsRole.IsNull() && !from.StsRole.IsUnknown() {
		if toStsRole, ok := to.GetStsRole(ctx); ok {
			if fromStsRole, ok := from.GetStsRole(ctx); ok {
				toStsRole.SyncFieldsDuringRead(ctx, fromStsRole)
				to.SetStsRole(ctx, toStsRole)
			}
		}
	}
}

func (m AwsCredentials) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AwsCredentials) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sts_role": reflect.TypeOf(StsRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsCredentials
// only implements ToObjectValue() and Type().
func (m AwsCredentials) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sts_role": m.StsRole,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AwsCredentials) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sts_role": StsRole{}.Type(ctx),
		},
	}
}

// GetStsRole returns the value of the StsRole field in AwsCredentials as
// a StsRole value.
// If the field is unknown or null, the boolean return value is false.
func (m *AwsCredentials) GetStsRole(ctx context.Context) (StsRole, bool) {
	var e StsRole
	if m.StsRole.IsNull() || m.StsRole.IsUnknown() {
		return e, false
	}
	var v StsRole
	d := m.StsRole.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStsRole sets the value of the StsRole field in AwsCredentials.
func (m *AwsCredentials) SetStsRole(ctx context.Context, v StsRole) {
	vs := v.ToObjectValue(ctx)
	m.StsRole = vs
}

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias types.String `tfsdk:"key_alias"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn types.String `tfsdk:"key_arn"`
	// The AWS KMS key region.
	KeyRegion types.String `tfsdk:"key_region"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to true or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to false.
	ReuseKeyForClusterVolumes types.Bool `tfsdk:"reuse_key_for_cluster_volumes"`
}

func (to *AwsKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AwsKeyInfo) {
}

func (to *AwsKeyInfo) SyncFieldsDuringRead(ctx context.Context, from AwsKeyInfo) {
}

func (m AwsKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AwsKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AwsKeyInfo
// only implements ToObjectValue() and Type().
func (m AwsKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key_alias":                     m.KeyAlias,
			"key_arn":                       m.KeyArn,
			"key_region":                    m.KeyRegion,
			"reuse_key_for_cluster_volumes": m.ReuseKeyForClusterVolumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AwsKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key_alias":                     types.StringType,
			"key_arn":                       types.StringType,
			"key_region":                    types.StringType,
			"reuse_key_for_cluster_volumes": types.BoolType,
		},
	}
}

type AzureKeyInfo struct {
	// The Disk Encryption Set id that is used to represent the key info used
	// for Managed Disk BYOK use case
	DiskEncryptionSetId types.String `tfsdk:"disk_encryption_set_id"`
	// The structure to store key access credential This is set if the Managed
	// Identity is being used to access the Azure Key Vault key.
	KeyAccessConfiguration types.Object `tfsdk:"key_access_configuration"`
	// The name of the key in KeyVault.
	KeyName types.String `tfsdk:"key_name"`
	// The base URI of the KeyVault.
	KeyVaultUri types.String `tfsdk:"key_vault_uri"`
	// The tenant id where the KeyVault lives.
	TenantId types.String `tfsdk:"tenant_id"`
	// The current key version.
	Version types.String `tfsdk:"version"`
}

func (to *AzureKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AzureKeyInfo) {
	if !from.KeyAccessConfiguration.IsNull() && !from.KeyAccessConfiguration.IsUnknown() {
		if toKeyAccessConfiguration, ok := to.GetKeyAccessConfiguration(ctx); ok {
			if fromKeyAccessConfiguration, ok := from.GetKeyAccessConfiguration(ctx); ok {
				// Recursively sync the fields of KeyAccessConfiguration
				toKeyAccessConfiguration.SyncFieldsDuringCreateOrUpdate(ctx, fromKeyAccessConfiguration)
				to.SetKeyAccessConfiguration(ctx, toKeyAccessConfiguration)
			}
		}
	}
}

func (to *AzureKeyInfo) SyncFieldsDuringRead(ctx context.Context, from AzureKeyInfo) {
	if !from.KeyAccessConfiguration.IsNull() && !from.KeyAccessConfiguration.IsUnknown() {
		if toKeyAccessConfiguration, ok := to.GetKeyAccessConfiguration(ctx); ok {
			if fromKeyAccessConfiguration, ok := from.GetKeyAccessConfiguration(ctx); ok {
				toKeyAccessConfiguration.SyncFieldsDuringRead(ctx, fromKeyAccessConfiguration)
				to.SetKeyAccessConfiguration(ctx, toKeyAccessConfiguration)
			}
		}
	}
}

func (m AzureKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disk_encryption_set_id"] = attrs["disk_encryption_set_id"].SetOptional()
	attrs["key_access_configuration"] = attrs["key_access_configuration"].SetOptional()
	attrs["key_name"] = attrs["key_name"].SetOptional()
	attrs["key_vault_uri"] = attrs["key_vault_uri"].SetOptional()
	attrs["tenant_id"] = attrs["tenant_id"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureKeyInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AzureKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"key_access_configuration": reflect.TypeOf(KeyAccessConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureKeyInfo
// only implements ToObjectValue() and Type().
func (m AzureKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disk_encryption_set_id":   m.DiskEncryptionSetId,
			"key_access_configuration": m.KeyAccessConfiguration,
			"key_name":                 m.KeyName,
			"key_vault_uri":            m.KeyVaultUri,
			"tenant_id":                m.TenantId,
			"version":                  m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AzureKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disk_encryption_set_id":   types.StringType,
			"key_access_configuration": KeyAccessConfiguration{}.Type(ctx),
			"key_name":                 types.StringType,
			"key_vault_uri":            types.StringType,
			"tenant_id":                types.StringType,
			"version":                  types.StringType,
		},
	}
}

// GetKeyAccessConfiguration returns the value of the KeyAccessConfiguration field in AzureKeyInfo as
// a KeyAccessConfiguration value.
// If the field is unknown or null, the boolean return value is false.
func (m *AzureKeyInfo) GetKeyAccessConfiguration(ctx context.Context) (KeyAccessConfiguration, bool) {
	var e KeyAccessConfiguration
	if m.KeyAccessConfiguration.IsNull() || m.KeyAccessConfiguration.IsUnknown() {
		return e, false
	}
	var v KeyAccessConfiguration
	d := m.KeyAccessConfiguration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKeyAccessConfiguration sets the value of the KeyAccessConfiguration field in AzureKeyInfo.
func (m *AzureKeyInfo) SetKeyAccessConfiguration(ctx context.Context, v KeyAccessConfiguration) {
	vs := v.ToObjectValue(ctx)
	m.KeyAccessConfiguration = vs
}

type AzureWorkspaceInfo struct {
	// Azure Resource Group name
	ResourceGroup types.String `tfsdk:"resource_group"`
	// Azure Subscription ID
	SubscriptionId types.String `tfsdk:"subscription_id"`
}

func (to *AzureWorkspaceInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AzureWorkspaceInfo) {
}

func (to *AzureWorkspaceInfo) SyncFieldsDuringRead(ctx context.Context, from AzureWorkspaceInfo) {
}

func (m AzureWorkspaceInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AzureWorkspaceInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureWorkspaceInfo
// only implements ToObjectValue() and Type().
func (m AzureWorkspaceInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"resource_group":  m.ResourceGroup,
			"subscription_id": m.SubscriptionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AzureWorkspaceInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resource_group":  types.StringType,
			"subscription_id": types.StringType,
		},
	}
}

type CloudResourceContainer struct {
	Gcp types.Object `tfsdk:"gcp"`
}

func (to *CloudResourceContainer) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CloudResourceContainer) {
	if !from.Gcp.IsNull() && !from.Gcp.IsUnknown() {
		if toGcp, ok := to.GetGcp(ctx); ok {
			if fromGcp, ok := from.GetGcp(ctx); ok {
				// Recursively sync the fields of Gcp
				toGcp.SyncFieldsDuringCreateOrUpdate(ctx, fromGcp)
				to.SetGcp(ctx, toGcp)
			}
		}
	}
}

func (to *CloudResourceContainer) SyncFieldsDuringRead(ctx context.Context, from CloudResourceContainer) {
	if !from.Gcp.IsNull() && !from.Gcp.IsUnknown() {
		if toGcp, ok := to.GetGcp(ctx); ok {
			if fromGcp, ok := from.GetGcp(ctx); ok {
				toGcp.SyncFieldsDuringRead(ctx, fromGcp)
				to.SetGcp(ctx, toGcp)
			}
		}
	}
}

func (m CloudResourceContainer) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CloudResourceContainer) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp": reflect.TypeOf(CustomerFacingGcpCloudResourceContainer{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CloudResourceContainer
// only implements ToObjectValue() and Type().
func (m CloudResourceContainer) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp": m.Gcp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CloudResourceContainer) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp": CustomerFacingGcpCloudResourceContainer{}.Type(ctx),
		},
	}
}

// GetGcp returns the value of the Gcp field in CloudResourceContainer as
// a CustomerFacingGcpCloudResourceContainer value.
// If the field is unknown or null, the boolean return value is false.
func (m *CloudResourceContainer) GetGcp(ctx context.Context) (CustomerFacingGcpCloudResourceContainer, bool) {
	var e CustomerFacingGcpCloudResourceContainer
	if m.Gcp.IsNull() || m.Gcp.IsUnknown() {
		return e, false
	}
	var v CustomerFacingGcpCloudResourceContainer
	d := m.Gcp.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcp sets the value of the Gcp field in CloudResourceContainer.
func (m *CloudResourceContainer) SetGcp(ctx context.Context, v CustomerFacingGcpCloudResourceContainer) {
	vs := v.ToObjectValue(ctx)
	m.Gcp = vs
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias types.String `tfsdk:"key_alias"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn types.String `tfsdk:"key_arn"`
	// The AWS KMS key region.
	KeyRegion types.String `tfsdk:"key_region"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to true or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to false.
	ReuseKeyForClusterVolumes types.Bool `tfsdk:"reuse_key_for_cluster_volumes"`
}

func (to *CreateAwsKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAwsKeyInfo) {
}

func (to *CreateAwsKeyInfo) SyncFieldsDuringRead(ctx context.Context, from CreateAwsKeyInfo) {
}

func (m CreateAwsKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key_alias"] = attrs["key_alias"].SetOptional()
	attrs["key_arn"] = attrs["key_arn"].SetRequired()
	attrs["key_region"] = attrs["key_region"].SetOptional()
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
func (m CreateAwsKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAwsKeyInfo
// only implements ToObjectValue() and Type().
func (m CreateAwsKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key_alias":                     m.KeyAlias,
			"key_arn":                       m.KeyArn,
			"key_region":                    m.KeyRegion,
			"reuse_key_for_cluster_volumes": m.ReuseKeyForClusterVolumes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateAwsKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key_alias":                     types.StringType,
			"key_arn":                       types.StringType,
			"key_region":                    types.StringType,
			"reuse_key_for_cluster_volumes": types.BoolType,
		},
	}
}

type CreateCredentialAwsCredentials struct {
	StsRole types.Object `tfsdk:"sts_role"`
}

func (to *CreateCredentialAwsCredentials) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialAwsCredentials) {
	if !from.StsRole.IsNull() && !from.StsRole.IsUnknown() {
		if toStsRole, ok := to.GetStsRole(ctx); ok {
			if fromStsRole, ok := from.GetStsRole(ctx); ok {
				// Recursively sync the fields of StsRole
				toStsRole.SyncFieldsDuringCreateOrUpdate(ctx, fromStsRole)
				to.SetStsRole(ctx, toStsRole)
			}
		}
	}
}

func (to *CreateCredentialAwsCredentials) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialAwsCredentials) {
	if !from.StsRole.IsNull() && !from.StsRole.IsUnknown() {
		if toStsRole, ok := to.GetStsRole(ctx); ok {
			if fromStsRole, ok := from.GetStsRole(ctx); ok {
				toStsRole.SyncFieldsDuringRead(ctx, fromStsRole)
				to.SetStsRole(ctx, toStsRole)
			}
		}
	}
}

func (m CreateCredentialAwsCredentials) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCredentialAwsCredentials) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sts_role": reflect.TypeOf(CreateCredentialStsRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialAwsCredentials
// only implements ToObjectValue() and Type().
func (m CreateCredentialAwsCredentials) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"sts_role": m.StsRole,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCredentialAwsCredentials) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sts_role": CreateCredentialStsRole{}.Type(ctx),
		},
	}
}

// GetStsRole returns the value of the StsRole field in CreateCredentialAwsCredentials as
// a CreateCredentialStsRole value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCredentialAwsCredentials) GetStsRole(ctx context.Context) (CreateCredentialStsRole, bool) {
	var e CreateCredentialStsRole
	if m.StsRole.IsNull() || m.StsRole.IsUnknown() {
		return e, false
	}
	var v CreateCredentialStsRole
	d := m.StsRole.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStsRole sets the value of the StsRole field in CreateCredentialAwsCredentials.
func (m *CreateCredentialAwsCredentials) SetStsRole(ctx context.Context, v CreateCredentialStsRole) {
	vs := v.ToObjectValue(ctx)
	m.StsRole = vs
}

type CreateCredentialRequest struct {
	AwsCredentials types.Object `tfsdk:"aws_credentials"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name"`
}

func (to *CreateCredentialRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialRequest) {
	if !from.AwsCredentials.IsNull() && !from.AwsCredentials.IsUnknown() {
		if toAwsCredentials, ok := to.GetAwsCredentials(ctx); ok {
			if fromAwsCredentials, ok := from.GetAwsCredentials(ctx); ok {
				// Recursively sync the fields of AwsCredentials
				toAwsCredentials.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsCredentials)
				to.SetAwsCredentials(ctx, toAwsCredentials)
			}
		}
	}
}

func (to *CreateCredentialRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialRequest) {
	if !from.AwsCredentials.IsNull() && !from.AwsCredentials.IsUnknown() {
		if toAwsCredentials, ok := to.GetAwsCredentials(ctx); ok {
			if fromAwsCredentials, ok := from.GetAwsCredentials(ctx); ok {
				toAwsCredentials.SyncFieldsDuringRead(ctx, fromAwsCredentials)
				to.SetAwsCredentials(ctx, toAwsCredentials)
			}
		}
	}
}

func (m CreateCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_credentials"] = attrs["aws_credentials"].SetRequired()
	attrs["credentials_name"] = attrs["credentials_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_credentials": reflect.TypeOf(CreateCredentialAwsCredentials{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialRequest
// only implements ToObjectValue() and Type().
func (m CreateCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_credentials":  m.AwsCredentials,
			"credentials_name": m.CredentialsName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCredentialRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateCredentialRequest) GetAwsCredentials(ctx context.Context) (CreateCredentialAwsCredentials, bool) {
	var e CreateCredentialAwsCredentials
	if m.AwsCredentials.IsNull() || m.AwsCredentials.IsUnknown() {
		return e, false
	}
	var v CreateCredentialAwsCredentials
	d := m.AwsCredentials.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsCredentials sets the value of the AwsCredentials field in CreateCredentialRequest.
func (m *CreateCredentialRequest) SetAwsCredentials(ctx context.Context, v CreateCredentialAwsCredentials) {
	vs := v.ToObjectValue(ctx)
	m.AwsCredentials = vs
}

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account IAM role.
	RoleArn types.String `tfsdk:"role_arn"`
}

func (to *CreateCredentialStsRole) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialStsRole) {
}

func (to *CreateCredentialStsRole) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialStsRole) {
}

func (m CreateCredentialStsRole) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCredentialStsRole) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialStsRole
// only implements ToObjectValue() and Type().
func (m CreateCredentialStsRole) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"role_arn": m.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCredentialStsRole) Type(ctx context.Context) attr.Type {
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

func (to *CreateCustomerManagedKeyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCustomerManagedKeyRequest) {
	if !from.AwsKeyInfo.IsNull() && !from.AwsKeyInfo.IsUnknown() {
		if toAwsKeyInfo, ok := to.GetAwsKeyInfo(ctx); ok {
			if fromAwsKeyInfo, ok := from.GetAwsKeyInfo(ctx); ok {
				// Recursively sync the fields of AwsKeyInfo
				toAwsKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsKeyInfo)
				to.SetAwsKeyInfo(ctx, toAwsKeyInfo)
			}
		}
	}
	if !from.GcpKeyInfo.IsNull() && !from.GcpKeyInfo.IsUnknown() {
		if toGcpKeyInfo, ok := to.GetGcpKeyInfo(ctx); ok {
			if fromGcpKeyInfo, ok := from.GetGcpKeyInfo(ctx); ok {
				// Recursively sync the fields of GcpKeyInfo
				toGcpKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpKeyInfo)
				to.SetGcpKeyInfo(ctx, toGcpKeyInfo)
			}
		}
	}
}

func (to *CreateCustomerManagedKeyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCustomerManagedKeyRequest) {
	if !from.AwsKeyInfo.IsNull() && !from.AwsKeyInfo.IsUnknown() {
		if toAwsKeyInfo, ok := to.GetAwsKeyInfo(ctx); ok {
			if fromAwsKeyInfo, ok := from.GetAwsKeyInfo(ctx); ok {
				toAwsKeyInfo.SyncFieldsDuringRead(ctx, fromAwsKeyInfo)
				to.SetAwsKeyInfo(ctx, toAwsKeyInfo)
			}
		}
	}
	if !from.GcpKeyInfo.IsNull() && !from.GcpKeyInfo.IsUnknown() {
		if toGcpKeyInfo, ok := to.GetGcpKeyInfo(ctx); ok {
			if fromGcpKeyInfo, ok := from.GetGcpKeyInfo(ctx); ok {
				toGcpKeyInfo.SyncFieldsDuringRead(ctx, fromGcpKeyInfo)
				to.SetGcpKeyInfo(ctx, toGcpKeyInfo)
			}
		}
	}
}

func (m CreateCustomerManagedKeyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_key_info"] = attrs["aws_key_info"].SetOptional()
	attrs["gcp_key_info"] = attrs["gcp_key_info"].SetOptional()
	attrs["use_cases"] = attrs["use_cases"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomerManagedKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCustomerManagedKeyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_key_info": reflect.TypeOf(CreateAwsKeyInfo{}),
		"gcp_key_info": reflect.TypeOf(CreateGcpKeyInfo{}),
		"use_cases":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomerManagedKeyRequest
// only implements ToObjectValue() and Type().
func (m CreateCustomerManagedKeyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_key_info": m.AwsKeyInfo,
			"gcp_key_info": m.GcpKeyInfo,
			"use_cases":    m.UseCases,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCustomerManagedKeyRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateCustomerManagedKeyRequest) GetAwsKeyInfo(ctx context.Context) (CreateAwsKeyInfo, bool) {
	var e CreateAwsKeyInfo
	if m.AwsKeyInfo.IsNull() || m.AwsKeyInfo.IsUnknown() {
		return e, false
	}
	var v CreateAwsKeyInfo
	d := m.AwsKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CreateCustomerManagedKeyRequest.
func (m *CreateCustomerManagedKeyRequest) SetAwsKeyInfo(ctx context.Context, v CreateAwsKeyInfo) {
	vs := v.ToObjectValue(ctx)
	m.AwsKeyInfo = vs
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest as
// a CreateGcpKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomerManagedKeyRequest) GetGcpKeyInfo(ctx context.Context) (CreateGcpKeyInfo, bool) {
	var e CreateGcpKeyInfo
	if m.GcpKeyInfo.IsNull() || m.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v CreateGcpKeyInfo
	d := m.GcpKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CreateCustomerManagedKeyRequest.
func (m *CreateCustomerManagedKeyRequest) SetGcpKeyInfo(ctx context.Context, v CreateGcpKeyInfo) {
	vs := v.ToObjectValue(ctx)
	m.GcpKeyInfo = vs
}

// GetUseCases returns the value of the UseCases field in CreateCustomerManagedKeyRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomerManagedKeyRequest) GetUseCases(ctx context.Context) ([]types.String, bool) {
	if m.UseCases.IsNull() || m.UseCases.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.UseCases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUseCases sets the value of the UseCases field in CreateCustomerManagedKeyRequest.
func (m *CreateCustomerManagedKeyRequest) SetUseCases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["use_cases"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UseCases = types.ListValueMust(t, vs)
}

type CreateGcpKeyInfo struct {
	// Globally unique kms key resource id of the form
	// projects/testProjectId/locations/us-east4/keyRings/gcpCmkKeyRing/cryptoKeys/cmk-eastus4
	KmsKeyId types.String `tfsdk:"kms_key_id"`
}

func (to *CreateGcpKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateGcpKeyInfo) {
}

func (to *CreateGcpKeyInfo) SyncFieldsDuringRead(ctx context.Context, from CreateGcpKeyInfo) {
}

func (m CreateGcpKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateGcpKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateGcpKeyInfo
// only implements ToObjectValue() and Type().
func (m CreateGcpKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kms_key_id": m.KmsKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateGcpKeyInfo) Type(ctx context.Context) attr.Type {
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
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId types.String `tfsdk:"vpc_id"`
}

func (to *CreateNetworkRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNetworkRequest) {
	if !from.GcpNetworkInfo.IsNull() && !from.GcpNetworkInfo.IsUnknown() {
		if toGcpNetworkInfo, ok := to.GetGcpNetworkInfo(ctx); ok {
			if fromGcpNetworkInfo, ok := from.GetGcpNetworkInfo(ctx); ok {
				// Recursively sync the fields of GcpNetworkInfo
				toGcpNetworkInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpNetworkInfo)
				to.SetGcpNetworkInfo(ctx, toGcpNetworkInfo)
			}
		}
	}
	if !from.SecurityGroupIds.IsNull() && !from.SecurityGroupIds.IsUnknown() && to.SecurityGroupIds.IsNull() && len(from.SecurityGroupIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecurityGroupIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecurityGroupIds = from.SecurityGroupIds
	}
	if !from.SubnetIds.IsNull() && !from.SubnetIds.IsUnknown() && to.SubnetIds.IsNull() && len(from.SubnetIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SubnetIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SubnetIds = from.SubnetIds
	}
	if !from.VpcEndpoints.IsNull() && !from.VpcEndpoints.IsUnknown() {
		if toVpcEndpoints, ok := to.GetVpcEndpoints(ctx); ok {
			if fromVpcEndpoints, ok := from.GetVpcEndpoints(ctx); ok {
				// Recursively sync the fields of VpcEndpoints
				toVpcEndpoints.SyncFieldsDuringCreateOrUpdate(ctx, fromVpcEndpoints)
				to.SetVpcEndpoints(ctx, toVpcEndpoints)
			}
		}
	}
}

func (to *CreateNetworkRequest) SyncFieldsDuringRead(ctx context.Context, from CreateNetworkRequest) {
	if !from.GcpNetworkInfo.IsNull() && !from.GcpNetworkInfo.IsUnknown() {
		if toGcpNetworkInfo, ok := to.GetGcpNetworkInfo(ctx); ok {
			if fromGcpNetworkInfo, ok := from.GetGcpNetworkInfo(ctx); ok {
				toGcpNetworkInfo.SyncFieldsDuringRead(ctx, fromGcpNetworkInfo)
				to.SetGcpNetworkInfo(ctx, toGcpNetworkInfo)
			}
		}
	}
	if !from.SecurityGroupIds.IsNull() && !from.SecurityGroupIds.IsUnknown() && to.SecurityGroupIds.IsNull() && len(from.SecurityGroupIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecurityGroupIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecurityGroupIds = from.SecurityGroupIds
	}
	if !from.SubnetIds.IsNull() && !from.SubnetIds.IsUnknown() && to.SubnetIds.IsNull() && len(from.SubnetIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SubnetIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SubnetIds = from.SubnetIds
	}
	if !from.VpcEndpoints.IsNull() && !from.VpcEndpoints.IsUnknown() {
		if toVpcEndpoints, ok := to.GetVpcEndpoints(ctx); ok {
			if fromVpcEndpoints, ok := from.GetVpcEndpoints(ctx); ok {
				toVpcEndpoints.SyncFieldsDuringRead(ctx, fromVpcEndpoints)
				to.SetVpcEndpoints(ctx, toVpcEndpoints)
			}
		}
	}
}

func (m CreateNetworkRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_network_info"] = attrs["gcp_network_info"].SetOptional()
	attrs["network_name"] = attrs["network_name"].SetOptional()
	attrs["security_group_ids"] = attrs["security_group_ids"].SetOptional()
	attrs["subnet_ids"] = attrs["subnet_ids"].SetOptional()
	attrs["vpc_endpoints"] = attrs["vpc_endpoints"].SetOptional()
	attrs["vpc_id"] = attrs["vpc_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateNetworkRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m CreateNetworkRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_network_info":   m.GcpNetworkInfo,
			"network_name":       m.NetworkName,
			"security_group_ids": m.SecurityGroupIds,
			"subnet_ids":         m.SubnetIds,
			"vpc_endpoints":      m.VpcEndpoints,
			"vpc_id":             m.VpcId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateNetworkRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateNetworkRequest) GetGcpNetworkInfo(ctx context.Context) (GcpNetworkInfo, bool) {
	var e GcpNetworkInfo
	if m.GcpNetworkInfo.IsNull() || m.GcpNetworkInfo.IsUnknown() {
		return e, false
	}
	var v GcpNetworkInfo
	d := m.GcpNetworkInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in CreateNetworkRequest.
func (m *CreateNetworkRequest) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo) {
	vs := v.ToObjectValue(ctx)
	m.GcpNetworkInfo = vs
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in CreateNetworkRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateNetworkRequest) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
	if m.SecurityGroupIds.IsNull() || m.SecurityGroupIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SecurityGroupIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in CreateNetworkRequest.
func (m *CreateNetworkRequest) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in CreateNetworkRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateNetworkRequest) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
	if m.SubnetIds.IsNull() || m.SubnetIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SubnetIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnetIds sets the value of the SubnetIds field in CreateNetworkRequest.
func (m *CreateNetworkRequest) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SubnetIds = types.ListValueMust(t, vs)
}

// GetVpcEndpoints returns the value of the VpcEndpoints field in CreateNetworkRequest as
// a NetworkVpcEndpoints value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateNetworkRequest) GetVpcEndpoints(ctx context.Context) (NetworkVpcEndpoints, bool) {
	var e NetworkVpcEndpoints
	if m.VpcEndpoints.IsNull() || m.VpcEndpoints.IsUnknown() {
		return e, false
	}
	var v NetworkVpcEndpoints
	d := m.VpcEndpoints.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in CreateNetworkRequest.
func (m *CreateNetworkRequest) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints) {
	vs := v.ToObjectValue(ctx)
	m.VpcEndpoints = vs
}

type CreatePrivateAccessSettingsRequest struct {
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID
	// returned when registering the VPC endpoint configuration in your
	// Databricks account. This is not the ID of the VPC endpoint in AWS. Only
	// used when private_access_level is set to ENDPOINT. This is an allow list
	// of VPC endpoints registered in your Databricks account that can connect
	// to your workspace over AWS PrivateLink. Note: If hybrid access to your
	// workspace is enabled by setting public_access_enabled to true, this
	// control only works for PrivateLink connections. To control how your
	// workspace is accessed via public internet, see IP access lists.
	AllowedVpcEndpointIds types.List `tfsdk:"allowed_vpc_endpoint_ids"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see allowed_vpc_endpoint_ids.
	PrivateAccessLevel types.String `tfsdk:"private_access_level"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify false, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify true, which means that public access is
	// enabled.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled"`
	// The AWS region for workspaces attached to this private access settings
	// object.
	Region types.String `tfsdk:"region"`
}

func (to *CreatePrivateAccessSettingsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePrivateAccessSettingsRequest) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (to *CreatePrivateAccessSettingsRequest) SyncFieldsDuringRead(ctx context.Context, from CreatePrivateAccessSettingsRequest) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (m CreatePrivateAccessSettingsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allowed_vpc_endpoint_ids"] = attrs["allowed_vpc_endpoint_ids"].SetOptional()
	attrs["private_access_level"] = attrs["private_access_level"].SetOptional()
	attrs["private_access_settings_name"] = attrs["private_access_settings_name"].SetOptional()
	attrs["public_access_enabled"] = attrs["public_access_enabled"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePrivateAccessSettingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreatePrivateAccessSettingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateAccessSettingsRequest
// only implements ToObjectValue() and Type().
func (m CreatePrivateAccessSettingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_vpc_endpoint_ids":     m.AllowedVpcEndpointIds,
			"private_access_level":         m.PrivateAccessLevel,
			"private_access_settings_name": m.PrivateAccessSettingsName,
			"public_access_enabled":        m.PublicAccessEnabled,
			"region":                       m.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePrivateAccessSettingsRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreatePrivateAccessSettingsRequest) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
	if m.AllowedVpcEndpointIds.IsNull() || m.AllowedVpcEndpointIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.AllowedVpcEndpointIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in CreatePrivateAccessSettingsRequest.
func (m *CreatePrivateAccessSettingsRequest) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type CreateStorageConfigurationRequest struct {
	// Optional IAM role that is used to access the workspace catalog which is
	// created during workspace creation for UC by Default. If a storage
	// configuration with this field populated is used to create a workspace,
	// then a workspace catalog is created together with the workspace. The
	// workspace catalog shares the root bucket with internal workspace storage
	// (including DBFS root) but uses a dedicated bucket path prefix.
	RoleArn types.String `tfsdk:"role_arn"`
	// Root S3 bucket information.
	RootBucketInfo types.Object `tfsdk:"root_bucket_info"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName types.String `tfsdk:"storage_configuration_name"`
}

func (to *CreateStorageConfigurationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateStorageConfigurationRequest) {
	if !from.RootBucketInfo.IsNull() && !from.RootBucketInfo.IsUnknown() {
		if toRootBucketInfo, ok := to.GetRootBucketInfo(ctx); ok {
			if fromRootBucketInfo, ok := from.GetRootBucketInfo(ctx); ok {
				// Recursively sync the fields of RootBucketInfo
				toRootBucketInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromRootBucketInfo)
				to.SetRootBucketInfo(ctx, toRootBucketInfo)
			}
		}
	}
}

func (to *CreateStorageConfigurationRequest) SyncFieldsDuringRead(ctx context.Context, from CreateStorageConfigurationRequest) {
	if !from.RootBucketInfo.IsNull() && !from.RootBucketInfo.IsUnknown() {
		if toRootBucketInfo, ok := to.GetRootBucketInfo(ctx); ok {
			if fromRootBucketInfo, ok := from.GetRootBucketInfo(ctx); ok {
				toRootBucketInfo.SyncFieldsDuringRead(ctx, fromRootBucketInfo)
				to.SetRootBucketInfo(ctx, toRootBucketInfo)
			}
		}
	}
}

func (m CreateStorageConfigurationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["role_arn"] = attrs["role_arn"].SetOptional()
	attrs["root_bucket_info"] = attrs["root_bucket_info"].SetRequired()
	attrs["storage_configuration_name"] = attrs["storage_configuration_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateStorageConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateStorageConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"root_bucket_info": reflect.TypeOf(RootBucketInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateStorageConfigurationRequest
// only implements ToObjectValue() and Type().
func (m CreateStorageConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"role_arn":                   m.RoleArn,
			"root_bucket_info":           m.RootBucketInfo,
			"storage_configuration_name": m.StorageConfigurationName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateStorageConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"role_arn":                   types.StringType,
			"root_bucket_info":           RootBucketInfo{}.Type(ctx),
			"storage_configuration_name": types.StringType,
		},
	}
}

// GetRootBucketInfo returns the value of the RootBucketInfo field in CreateStorageConfigurationRequest as
// a RootBucketInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateStorageConfigurationRequest) GetRootBucketInfo(ctx context.Context) (RootBucketInfo, bool) {
	var e RootBucketInfo
	if m.RootBucketInfo.IsNull() || m.RootBucketInfo.IsUnknown() {
		return e, false
	}
	var v RootBucketInfo
	d := m.RootBucketInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in CreateStorageConfigurationRequest.
func (m *CreateStorageConfigurationRequest) SetRootBucketInfo(ctx context.Context, v RootBucketInfo) {
	vs := v.ToObjectValue(ctx)
	m.RootBucketInfo = vs
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId types.String `tfsdk:"aws_vpc_endpoint_id"`
	// The cloud info of this vpc endpoint.
	GcpVpcEndpointInfo types.Object `tfsdk:"gcp_vpc_endpoint_info"`
	// The region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name"`
}

func (to *CreateVpcEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateVpcEndpointRequest) {
	if !from.GcpVpcEndpointInfo.IsNull() && !from.GcpVpcEndpointInfo.IsUnknown() {
		if toGcpVpcEndpointInfo, ok := to.GetGcpVpcEndpointInfo(ctx); ok {
			if fromGcpVpcEndpointInfo, ok := from.GetGcpVpcEndpointInfo(ctx); ok {
				// Recursively sync the fields of GcpVpcEndpointInfo
				toGcpVpcEndpointInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpVpcEndpointInfo)
				to.SetGcpVpcEndpointInfo(ctx, toGcpVpcEndpointInfo)
			}
		}
	}
}

func (to *CreateVpcEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from CreateVpcEndpointRequest) {
	if !from.GcpVpcEndpointInfo.IsNull() && !from.GcpVpcEndpointInfo.IsUnknown() {
		if toGcpVpcEndpointInfo, ok := to.GetGcpVpcEndpointInfo(ctx); ok {
			if fromGcpVpcEndpointInfo, ok := from.GetGcpVpcEndpointInfo(ctx); ok {
				toGcpVpcEndpointInfo.SyncFieldsDuringRead(ctx, fromGcpVpcEndpointInfo)
				to.SetGcpVpcEndpointInfo(ctx, toGcpVpcEndpointInfo)
			}
		}
	}
}

func (m CreateVpcEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_vpc_endpoint_id"] = attrs["aws_vpc_endpoint_id"].SetOptional()
	attrs["gcp_vpc_endpoint_info"] = attrs["gcp_vpc_endpoint_info"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["vpc_endpoint_name"] = attrs["vpc_endpoint_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateVpcEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_vpc_endpoint_info": reflect.TypeOf(GcpVpcEndpointInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVpcEndpointRequest
// only implements ToObjectValue() and Type().
func (m CreateVpcEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_vpc_endpoint_id":   m.AwsVpcEndpointId,
			"gcp_vpc_endpoint_info": m.GcpVpcEndpointInfo,
			"region":                m.Region,
			"vpc_endpoint_name":     m.VpcEndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateVpcEndpointRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateVpcEndpointRequest) GetGcpVpcEndpointInfo(ctx context.Context) (GcpVpcEndpointInfo, bool) {
	var e GcpVpcEndpointInfo
	if m.GcpVpcEndpointInfo.IsNull() || m.GcpVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v GcpVpcEndpointInfo
	d := m.GcpVpcEndpointInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in CreateVpcEndpointRequest.
func (m *CreateVpcEndpointRequest) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo) {
	vs := v.ToObjectValue(ctx)
	m.GcpVpcEndpointInfo = vs
}

type CreateWorkspaceRequest struct {
	AwsRegion types.String `tfsdk:"aws_region"`
	// The cloud name. This field always has the value `gcp`.
	Cloud types.String `tfsdk:"cloud"`

	CloudResourceContainer types.Object `tfsdk:"cloud_resource_container"`
	// If the compute mode is `SERVERLESS`, a serverless workspace is created
	// that comes pre-configured with serverless compute and default storage,
	// providing a fully-managed, enterprise-ready SaaS experience. This means
	// you don't need to provide any resources managed by you, such as
	// credentials, storage, or network. If the compute mode is `HYBRID` (which
	// is the default option), a classic workspace is created that uses
	// customer-managed resources.
	ComputeMode types.String `tfsdk:"compute_mode"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags types.Map `tfsdk:"custom_tags"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for the web application and REST APIs is
	// <workspace-deployment-name>.cloud.databricks.com. For example, if the
	// deployment name is abcsales, your workspace URL will be
	// https://abcsales.cloud.databricks.com. Hyphens are allowed. This property
	// supports only the set of characters that are allowed in a subdomain. To
	// set this value, you must have a deployment name prefix. Contact your
	// Databricks account team to add an account deployment name prefix to your
	// account. Workspace deployment names follow the account prefix and a
	// hyphen. For example, if your account's deployment prefix is acme and the
	// workspace deployment name is workspace-1, the JSON response for the
	// deployment_name field becomes acme-workspace-1. The workspace URL would
	// be acme-workspace-1.cloud.databricks.com. You can also set the
	// deployment_name to the reserved keyword EMPTY if you want the deployment
	// name to only include the deployment prefix. For example, if your
	// account's deployment prefix is acme and the workspace deployment name is
	// EMPTY, the deployment_name becomes acme only and the workspace URL is
	// acme.cloud.databricks.com. This value must be unique across all
	// non-deleted deployments across all AWS regions. If a new workspace omits
	// this property, the server generates a unique deployment name for you with
	// the pattern dbc-xxxxxxxx-xxxx.
	DeploymentName types.String `tfsdk:"deployment_name"`

	GcpManagedNetworkConfig types.Object `tfsdk:"gcp_managed_network_config"`

	GkeConfig types.Object `tfsdk:"gke_config"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location types.String `tfsdk:"location"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property use_cases must
	// contain MANAGED_SERVICES.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id"`
	// The object ID of network connectivity config. Once assigned, the
	// workspace serverless compute resources use the same set of stable IP CIDR
	// blocks and optional private link to access your resources.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id"`
	// The ID of the workspace's network configuration object. To use AWS
	// PrivateLink, this field is required.
	NetworkId types.String `tfsdk:"network_id"`

	PricingTier types.String `tfsdk:"pricing_tier"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. You must specify this ID if you are using [AWS PrivateLink]
	// for either front-end (user-to-workspace connection), back-end (data plane
	// to control plane connection), or both connection types. Before
	// configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property use_cases must contain STORAGE.
	StorageCustomerManagedKeyId types.String `tfsdk:"storage_customer_managed_key_id"`
	// The human-readable name of the workspace.
	WorkspaceName types.String `tfsdk:"workspace_name"`
}

func (to *CreateWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWorkspaceRequest) {
	if !from.CloudResourceContainer.IsNull() && !from.CloudResourceContainer.IsUnknown() {
		if toCloudResourceContainer, ok := to.GetCloudResourceContainer(ctx); ok {
			if fromCloudResourceContainer, ok := from.GetCloudResourceContainer(ctx); ok {
				// Recursively sync the fields of CloudResourceContainer
				toCloudResourceContainer.SyncFieldsDuringCreateOrUpdate(ctx, fromCloudResourceContainer)
				to.SetCloudResourceContainer(ctx, toCloudResourceContainer)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				// Recursively sync the fields of GcpManagedNetworkConfig
				toGcpManagedNetworkConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
	if !from.GkeConfig.IsNull() && !from.GkeConfig.IsUnknown() {
		if toGkeConfig, ok := to.GetGkeConfig(ctx); ok {
			if fromGkeConfig, ok := from.GetGkeConfig(ctx); ok {
				// Recursively sync the fields of GkeConfig
				toGkeConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGkeConfig)
				to.SetGkeConfig(ctx, toGkeConfig)
			}
		}
	}
}

func (to *CreateWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from CreateWorkspaceRequest) {
	if !from.CloudResourceContainer.IsNull() && !from.CloudResourceContainer.IsUnknown() {
		if toCloudResourceContainer, ok := to.GetCloudResourceContainer(ctx); ok {
			if fromCloudResourceContainer, ok := from.GetCloudResourceContainer(ctx); ok {
				toCloudResourceContainer.SyncFieldsDuringRead(ctx, fromCloudResourceContainer)
				to.SetCloudResourceContainer(ctx, toCloudResourceContainer)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				toGcpManagedNetworkConfig.SyncFieldsDuringRead(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
	if !from.GkeConfig.IsNull() && !from.GkeConfig.IsUnknown() {
		if toGkeConfig, ok := to.GetGkeConfig(ctx); ok {
			if fromGkeConfig, ok := from.GetGkeConfig(ctx); ok {
				toGkeConfig.SyncFieldsDuringRead(ctx, fromGkeConfig)
				to.SetGkeConfig(ctx, toGkeConfig)
			}
		}
	}
}

func (m CreateWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_region"] = attrs["aws_region"].SetOptional()
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["cloud_resource_container"] = attrs["cloud_resource_container"].SetOptional()
	attrs["compute_mode"] = attrs["compute_mode"].SetOptional()
	attrs["credentials_id"] = attrs["credentials_id"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["deployment_name"] = attrs["deployment_name"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].SetOptional()
	attrs["gke_config"] = attrs["gke_config"].SetOptional()
	attrs["location"] = attrs["location"].SetOptional()
	attrs["managed_services_customer_managed_key_id"] = attrs["managed_services_customer_managed_key_id"].SetOptional()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetOptional()
	attrs["network_id"] = attrs["network_id"].SetOptional()
	attrs["pricing_tier"] = attrs["pricing_tier"].SetOptional()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetOptional()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetOptional()
	attrs["storage_customer_managed_key_id"] = attrs["storage_customer_managed_key_id"].SetOptional()
	attrs["workspace_name"] = attrs["workspace_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m CreateWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_region":                 m.AwsRegion,
			"cloud":                      m.Cloud,
			"cloud_resource_container":   m.CloudResourceContainer,
			"compute_mode":               m.ComputeMode,
			"credentials_id":             m.CredentialsId,
			"custom_tags":                m.CustomTags,
			"deployment_name":            m.DeploymentName,
			"gcp_managed_network_config": m.GcpManagedNetworkConfig,
			"gke_config":                 m.GkeConfig,
			"location":                   m.Location,
			"managed_services_customer_managed_key_id": m.ManagedServicesCustomerManagedKeyId,
			"network_connectivity_config_id":           m.NetworkConnectivityConfigId,
			"network_id":                               m.NetworkId,
			"pricing_tier":                             m.PricingTier,
			"private_access_settings_id":               m.PrivateAccessSettingsId,
			"storage_configuration_id":                 m.StorageConfigurationId,
			"storage_customer_managed_key_id":          m.StorageCustomerManagedKeyId,
			"workspace_name":                           m.WorkspaceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_region":               types.StringType,
			"cloud":                    types.StringType,
			"cloud_resource_container": CloudResourceContainer{}.Type(ctx),
			"compute_mode":             types.StringType,
			"credentials_id":           types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"deployment_name":            types.StringType,
			"gcp_managed_network_config": GcpManagedNetworkConfig{}.Type(ctx),
			"gke_config":                 GkeConfig{}.Type(ctx),
			"location":                   types.StringType,
			"managed_services_customer_managed_key_id": types.StringType,
			"network_connectivity_config_id":           types.StringType,
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
func (m *CreateWorkspaceRequest) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer, bool) {
	var e CloudResourceContainer
	if m.CloudResourceContainer.IsNull() || m.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v CloudResourceContainer
	d := m.CloudResourceContainer.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in CreateWorkspaceRequest.
func (m *CreateWorkspaceRequest) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer) {
	vs := v.ToObjectValue(ctx)
	m.CloudResourceContainer = vs
}

// GetCustomTags returns the value of the CustomTags field in CreateWorkspaceRequest as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceRequest) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in CreateWorkspaceRequest.
func (m *CreateWorkspaceRequest) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest as
// a GcpManagedNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceRequest) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig, bool) {
	var e GcpManagedNetworkConfig
	if m.GcpManagedNetworkConfig.IsNull() || m.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v GcpManagedNetworkConfig
	d := m.GcpManagedNetworkConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in CreateWorkspaceRequest.
func (m *CreateWorkspaceRequest) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig) {
	vs := v.ToObjectValue(ctx)
	m.GcpManagedNetworkConfig = vs
}

// GetGkeConfig returns the value of the GkeConfig field in CreateWorkspaceRequest as
// a GkeConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWorkspaceRequest) GetGkeConfig(ctx context.Context) (GkeConfig, bool) {
	var e GkeConfig
	if m.GkeConfig.IsNull() || m.GkeConfig.IsUnknown() {
		return e, false
	}
	var v GkeConfig
	d := m.GkeConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGkeConfig sets the value of the GkeConfig field in CreateWorkspaceRequest.
func (m *CreateWorkspaceRequest) SetGkeConfig(ctx context.Context, v GkeConfig) {
	vs := v.ToObjectValue(ctx)
	m.GkeConfig = vs
}

type Credential struct {
	AwsCredentials types.Object `tfsdk:"aws_credentials"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Databricks credential configuration ID.
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The human-readable name of the credential configuration object.
	CredentialsName types.String `tfsdk:"credentials_name"`
}

func (to *Credential) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Credential) {
	if !from.AwsCredentials.IsNull() && !from.AwsCredentials.IsUnknown() {
		if toAwsCredentials, ok := to.GetAwsCredentials(ctx); ok {
			if fromAwsCredentials, ok := from.GetAwsCredentials(ctx); ok {
				// Recursively sync the fields of AwsCredentials
				toAwsCredentials.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsCredentials)
				to.SetAwsCredentials(ctx, toAwsCredentials)
			}
		}
	}
}

func (to *Credential) SyncFieldsDuringRead(ctx context.Context, from Credential) {
	if !from.AwsCredentials.IsNull() && !from.AwsCredentials.IsUnknown() {
		if toAwsCredentials, ok := to.GetAwsCredentials(ctx); ok {
			if fromAwsCredentials, ok := from.GetAwsCredentials(ctx); ok {
				toAwsCredentials.SyncFieldsDuringRead(ctx, fromAwsCredentials)
				to.SetAwsCredentials(ctx, toAwsCredentials)
			}
		}
	}
}

func (m Credential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Credential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_credentials": reflect.TypeOf(AwsCredentials{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Credential
// only implements ToObjectValue() and Type().
func (m Credential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_credentials":  m.AwsCredentials,
			"creation_time":    m.CreationTime,
			"credentials_id":   m.CredentialsId,
			"credentials_name": m.CredentialsName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Credential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
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
func (m *Credential) GetAwsCredentials(ctx context.Context) (AwsCredentials, bool) {
	var e AwsCredentials
	if m.AwsCredentials.IsNull() || m.AwsCredentials.IsUnknown() {
		return e, false
	}
	var v AwsCredentials
	d := m.AwsCredentials.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsCredentials sets the value of the AwsCredentials field in Credential.
func (m *Credential) SetAwsCredentials(ctx context.Context, v AwsCredentials) {
	vs := v.ToObjectValue(ctx)
	m.AwsCredentials = vs
}

type CustomerFacingGcpCloudResourceContainer struct {
	ProjectId types.String `tfsdk:"project_id"`
}

func (to *CustomerFacingGcpCloudResourceContainer) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomerFacingGcpCloudResourceContainer) {
}

func (to *CustomerFacingGcpCloudResourceContainer) SyncFieldsDuringRead(ctx context.Context, from CustomerFacingGcpCloudResourceContainer) {
}

func (m CustomerFacingGcpCloudResourceContainer) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CustomerFacingGcpCloudResourceContainer) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerFacingGcpCloudResourceContainer
// only implements ToObjectValue() and Type().
func (m CustomerFacingGcpCloudResourceContainer) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomerFacingGcpCloudResourceContainer) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project_id": types.StringType,
		},
	}
}

type CustomerManagedKey struct {
	AwsKeyInfo types.Object `tfsdk:"aws_key_info"`

	AzureKeyInfo types.Object `tfsdk:"azure_key_info"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId types.String `tfsdk:"customer_managed_key_id"`

	GcpKeyInfo types.Object `tfsdk:"gcp_key_info"`
	// The cases that the key can be used for.
	UseCases types.List `tfsdk:"use_cases"`
}

func (to *CustomerManagedKey) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomerManagedKey) {
	if !from.AwsKeyInfo.IsNull() && !from.AwsKeyInfo.IsUnknown() {
		if toAwsKeyInfo, ok := to.GetAwsKeyInfo(ctx); ok {
			if fromAwsKeyInfo, ok := from.GetAwsKeyInfo(ctx); ok {
				// Recursively sync the fields of AwsKeyInfo
				toAwsKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsKeyInfo)
				to.SetAwsKeyInfo(ctx, toAwsKeyInfo)
			}
		}
	}
	if !from.AzureKeyInfo.IsNull() && !from.AzureKeyInfo.IsUnknown() {
		if toAzureKeyInfo, ok := to.GetAzureKeyInfo(ctx); ok {
			if fromAzureKeyInfo, ok := from.GetAzureKeyInfo(ctx); ok {
				// Recursively sync the fields of AzureKeyInfo
				toAzureKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureKeyInfo)
				to.SetAzureKeyInfo(ctx, toAzureKeyInfo)
			}
		}
	}
	if !from.GcpKeyInfo.IsNull() && !from.GcpKeyInfo.IsUnknown() {
		if toGcpKeyInfo, ok := to.GetGcpKeyInfo(ctx); ok {
			if fromGcpKeyInfo, ok := from.GetGcpKeyInfo(ctx); ok {
				// Recursively sync the fields of GcpKeyInfo
				toGcpKeyInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpKeyInfo)
				to.SetGcpKeyInfo(ctx, toGcpKeyInfo)
			}
		}
	}
	if !from.UseCases.IsNull() && !from.UseCases.IsUnknown() && to.UseCases.IsNull() && len(from.UseCases.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UseCases, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UseCases = from.UseCases
	}
}

func (to *CustomerManagedKey) SyncFieldsDuringRead(ctx context.Context, from CustomerManagedKey) {
	if !from.AwsKeyInfo.IsNull() && !from.AwsKeyInfo.IsUnknown() {
		if toAwsKeyInfo, ok := to.GetAwsKeyInfo(ctx); ok {
			if fromAwsKeyInfo, ok := from.GetAwsKeyInfo(ctx); ok {
				toAwsKeyInfo.SyncFieldsDuringRead(ctx, fromAwsKeyInfo)
				to.SetAwsKeyInfo(ctx, toAwsKeyInfo)
			}
		}
	}
	if !from.AzureKeyInfo.IsNull() && !from.AzureKeyInfo.IsUnknown() {
		if toAzureKeyInfo, ok := to.GetAzureKeyInfo(ctx); ok {
			if fromAzureKeyInfo, ok := from.GetAzureKeyInfo(ctx); ok {
				toAzureKeyInfo.SyncFieldsDuringRead(ctx, fromAzureKeyInfo)
				to.SetAzureKeyInfo(ctx, toAzureKeyInfo)
			}
		}
	}
	if !from.GcpKeyInfo.IsNull() && !from.GcpKeyInfo.IsUnknown() {
		if toGcpKeyInfo, ok := to.GetGcpKeyInfo(ctx); ok {
			if fromGcpKeyInfo, ok := from.GetGcpKeyInfo(ctx); ok {
				toGcpKeyInfo.SyncFieldsDuringRead(ctx, fromGcpKeyInfo)
				to.SetGcpKeyInfo(ctx, toGcpKeyInfo)
			}
		}
	}
	if !from.UseCases.IsNull() && !from.UseCases.IsUnknown() && to.UseCases.IsNull() && len(from.UseCases.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UseCases, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UseCases = from.UseCases
	}
}

func (m CustomerManagedKey) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_key_info"] = attrs["aws_key_info"].SetOptional()
	attrs["azure_key_info"] = attrs["azure_key_info"].SetOptional()
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
func (m CustomerManagedKey) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_key_info":   reflect.TypeOf(AwsKeyInfo{}),
		"azure_key_info": reflect.TypeOf(AzureKeyInfo{}),
		"gcp_key_info":   reflect.TypeOf(GcpKeyInfo{}),
		"use_cases":      reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerManagedKey
// only implements ToObjectValue() and Type().
func (m CustomerManagedKey) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_key_info":            m.AwsKeyInfo,
			"azure_key_info":          m.AzureKeyInfo,
			"creation_time":           m.CreationTime,
			"customer_managed_key_id": m.CustomerManagedKeyId,
			"gcp_key_info":            m.GcpKeyInfo,
			"use_cases":               m.UseCases,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomerManagedKey) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_key_info":            AwsKeyInfo{}.Type(ctx),
			"azure_key_info":          AzureKeyInfo{}.Type(ctx),
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
func (m *CustomerManagedKey) GetAwsKeyInfo(ctx context.Context) (AwsKeyInfo, bool) {
	var e AwsKeyInfo
	if m.AwsKeyInfo.IsNull() || m.AwsKeyInfo.IsUnknown() {
		return e, false
	}
	var v AwsKeyInfo
	d := m.AwsKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsKeyInfo sets the value of the AwsKeyInfo field in CustomerManagedKey.
func (m *CustomerManagedKey) SetAwsKeyInfo(ctx context.Context, v AwsKeyInfo) {
	vs := v.ToObjectValue(ctx)
	m.AwsKeyInfo = vs
}

// GetAzureKeyInfo returns the value of the AzureKeyInfo field in CustomerManagedKey as
// a AzureKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomerManagedKey) GetAzureKeyInfo(ctx context.Context) (AzureKeyInfo, bool) {
	var e AzureKeyInfo
	if m.AzureKeyInfo.IsNull() || m.AzureKeyInfo.IsUnknown() {
		return e, false
	}
	var v AzureKeyInfo
	d := m.AzureKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureKeyInfo sets the value of the AzureKeyInfo field in CustomerManagedKey.
func (m *CustomerManagedKey) SetAzureKeyInfo(ctx context.Context, v AzureKeyInfo) {
	vs := v.ToObjectValue(ctx)
	m.AzureKeyInfo = vs
}

// GetGcpKeyInfo returns the value of the GcpKeyInfo field in CustomerManagedKey as
// a GcpKeyInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomerManagedKey) GetGcpKeyInfo(ctx context.Context) (GcpKeyInfo, bool) {
	var e GcpKeyInfo
	if m.GcpKeyInfo.IsNull() || m.GcpKeyInfo.IsUnknown() {
		return e, false
	}
	var v GcpKeyInfo
	d := m.GcpKeyInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpKeyInfo sets the value of the GcpKeyInfo field in CustomerManagedKey.
func (m *CustomerManagedKey) SetGcpKeyInfo(ctx context.Context, v GcpKeyInfo) {
	vs := v.ToObjectValue(ctx)
	m.GcpKeyInfo = vs
}

// GetUseCases returns the value of the UseCases field in CustomerManagedKey as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomerManagedKey) GetUseCases(ctx context.Context) ([]types.String, bool) {
	if m.UseCases.IsNull() || m.UseCases.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.UseCases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUseCases sets the value of the UseCases field in CustomerManagedKey.
func (m *CustomerManagedKey) SetUseCases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["use_cases"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UseCases = types.ListValueMust(t, vs)
}

type DeleteCredentialRequest struct {
	// Databricks Account API credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
}

func (to *DeleteCredentialRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCredentialRequest) {
}

func (to *DeleteCredentialRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteCredentialRequest) {
}

func (m DeleteCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credentials_id"] = attrs["credentials_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialRequest
// only implements ToObjectValue() and Type().
func (m DeleteCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials_id": m.CredentialsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCredentialRequest) Type(ctx context.Context) attr.Type {
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

func (to *DeleteEncryptionKeyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEncryptionKeyRequest) {
}

func (to *DeleteEncryptionKeyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteEncryptionKeyRequest) {
}

func (m DeleteEncryptionKeyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["customer_managed_key_id"] = attrs["customer_managed_key_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEncryptionKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteEncryptionKeyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEncryptionKeyRequest
// only implements ToObjectValue() and Type().
func (m DeleteEncryptionKeyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_managed_key_id": m.CustomerManagedKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEncryptionKeyRequest) Type(ctx context.Context) attr.Type {
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

func (to *DeleteNetworkRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteNetworkRequest) {
}

func (to *DeleteNetworkRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteNetworkRequest) {
}

func (m DeleteNetworkRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_id"] = attrs["network_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteNetworkRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkRequest
// only implements ToObjectValue() and Type().
func (m DeleteNetworkRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_id": m.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteNetworkRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_id": types.StringType,
		},
	}
}

type DeletePrivateAccesRequest struct {
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (to *DeletePrivateAccesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePrivateAccesRequest) {
}

func (to *DeletePrivateAccesRequest) SyncFieldsDuringRead(ctx context.Context, from DeletePrivateAccesRequest) {
}

func (m DeletePrivateAccesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePrivateAccesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePrivateAccesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePrivateAccesRequest
// only implements ToObjectValue() and Type().
func (m DeletePrivateAccesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_access_settings_id": m.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePrivateAccesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_access_settings_id": types.StringType,
		},
	}
}

type DeleteStorageRequest struct {
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (to *DeleteStorageRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteStorageRequest) {
}

func (to *DeleteStorageRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteStorageRequest) {
}

func (m DeleteStorageRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteStorageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteStorageRequest
// only implements ToObjectValue() and Type().
func (m DeleteStorageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_configuration_id": m.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteStorageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_configuration_id": types.StringType,
		},
	}
}

type DeleteVpcEndpointRequest struct {
	VpcEndpointId types.String `tfsdk:"-"`
}

func (to *DeleteVpcEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteVpcEndpointRequest) {
}

func (to *DeleteVpcEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteVpcEndpointRequest) {
}

func (m DeleteVpcEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["vpc_endpoint_id"] = attrs["vpc_endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteVpcEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteVpcEndpointRequest
// only implements ToObjectValue() and Type().
func (m DeleteVpcEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vpc_endpoint_id": m.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteVpcEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vpc_endpoint_id": types.StringType,
		},
	}
}

type DeleteWorkspaceRequest struct {
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *DeleteWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWorkspaceRequest) {
}

func (to *DeleteWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteWorkspaceRequest) {
}

func (m DeleteWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWorkspaceRequest
// only implements ToObjectValue() and Type().
func (m DeleteWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

// The shared network config for GCP workspace. This object has common network
// configurations that are network attributions of a workspace. DEPRECATED. Use
// GkeConfig instead.
type GcpCommonNetworkConfig struct {
	// The IP range that will be used to allocate GKE cluster master resources
	// from. This field must not be set if
	// gke_cluster_type=PUBLIC_NODE_PUBLIC_MASTER.
	GkeClusterMasterIpRange types.String `tfsdk:"gke_cluster_master_ip_range"`
	// The type of network connectivity of the GKE cluster.
	GkeConnectivityType types.String `tfsdk:"gke_connectivity_type"`
}

func (to *GcpCommonNetworkConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpCommonNetworkConfig) {
}

func (to *GcpCommonNetworkConfig) SyncFieldsDuringRead(ctx context.Context, from GcpCommonNetworkConfig) {
}

func (m GcpCommonNetworkConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gke_cluster_master_ip_range"] = attrs["gke_cluster_master_ip_range"].SetOptional()
	attrs["gke_connectivity_type"] = attrs["gke_connectivity_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpCommonNetworkConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GcpCommonNetworkConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpCommonNetworkConfig
// only implements ToObjectValue() and Type().
func (m GcpCommonNetworkConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gke_cluster_master_ip_range": m.GkeClusterMasterIpRange,
			"gke_connectivity_type":       m.GkeConnectivityType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GcpCommonNetworkConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gke_cluster_master_ip_range": types.StringType,
			"gke_connectivity_type":       types.StringType,
		},
	}
}

type GcpKeyInfo struct {
	// Globally unique kms key resource id of the form
	// projects/testProjectId/locations/us-east4/keyRings/gcpCmkKeyRing/cryptoKeys/cmk-eastus4
	KmsKeyId types.String `tfsdk:"kms_key_id"`
}

func (to *GcpKeyInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpKeyInfo) {
}

func (to *GcpKeyInfo) SyncFieldsDuringRead(ctx context.Context, from GcpKeyInfo) {
}

func (m GcpKeyInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GcpKeyInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpKeyInfo
// only implements ToObjectValue() and Type().
func (m GcpKeyInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"kms_key_id": m.KmsKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GcpKeyInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"kms_key_id": types.StringType,
		},
	}
}

// The network configuration for the workspace.
type GcpManagedNetworkConfig struct {
	// The IP range that will be used to allocate GKE cluster Pods from.
	GkeClusterPodIpRange types.String `tfsdk:"gke_cluster_pod_ip_range"`
	// The IP range that will be used to allocate GKE cluster Services from.
	GkeClusterServiceIpRange types.String `tfsdk:"gke_cluster_service_ip_range"`
	// The IP range which will be used to allocate GKE cluster nodes from. Note:
	// Pods, services and master IP range must be mutually exclusive.
	SubnetCidr types.String `tfsdk:"subnet_cidr"`
}

func (to *GcpManagedNetworkConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpManagedNetworkConfig) {
}

func (to *GcpManagedNetworkConfig) SyncFieldsDuringRead(ctx context.Context, from GcpManagedNetworkConfig) {
}

func (m GcpManagedNetworkConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GcpManagedNetworkConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpManagedNetworkConfig
// only implements ToObjectValue() and Type().
func (m GcpManagedNetworkConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gke_cluster_pod_ip_range":     m.GkeClusterPodIpRange,
			"gke_cluster_service_ip_range": m.GkeClusterServiceIpRange,
			"subnet_cidr":                  m.SubnetCidr,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GcpManagedNetworkConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gke_cluster_pod_ip_range":     types.StringType,
			"gke_cluster_service_ip_range": types.StringType,
			"subnet_cidr":                  types.StringType,
		},
	}
}

type GcpNetworkInfo struct {
	// The GCP project ID for network resources. This project is where the VPC
	// and subnet resides.
	NetworkProjectId types.String `tfsdk:"network_project_id"`
	// Name of the secondary range within the subnet that will be used by GKE as
	// Pod IP range. This is BYO VPC specific. DB VPC uses
	// network.getGcpManagedNetworkConfig.getGkeClusterPodIpRange
	PodIpRangeName types.String `tfsdk:"pod_ip_range_name"`
	// Name of the secondary range within the subnet that will be used by GKE as
	// Service IP range.
	ServiceIpRangeName types.String `tfsdk:"service_ip_range_name"`
	// The customer-provided Subnet ID that will be available to Clusters in
	// Workspaces using this Network.
	SubnetId types.String `tfsdk:"subnet_id"`

	SubnetRegion types.String `tfsdk:"subnet_region"`
	// The customer-provided VPC ID.
	VpcId types.String `tfsdk:"vpc_id"`
}

func (to *GcpNetworkInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpNetworkInfo) {
}

func (to *GcpNetworkInfo) SyncFieldsDuringRead(ctx context.Context, from GcpNetworkInfo) {
}

func (m GcpNetworkInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GcpNetworkInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpNetworkInfo
// only implements ToObjectValue() and Type().
func (m GcpNetworkInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_project_id":    m.NetworkProjectId,
			"pod_ip_range_name":     m.PodIpRangeName,
			"service_ip_range_name": m.ServiceIpRangeName,
			"subnet_id":             m.SubnetId,
			"subnet_region":         m.SubnetRegion,
			"vpc_id":                m.VpcId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GcpNetworkInfo) Type(ctx context.Context) attr.Type {
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

type GcpVpcEndpointInfo struct {
	EndpointRegion types.String `tfsdk:"endpoint_region"`

	ProjectId types.String `tfsdk:"project_id"`

	PscConnectionId types.String `tfsdk:"psc_connection_id"`

	PscEndpointName types.String `tfsdk:"psc_endpoint_name"`

	ServiceAttachmentId types.String `tfsdk:"service_attachment_id"`
}

func (to *GcpVpcEndpointInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GcpVpcEndpointInfo) {
}

func (to *GcpVpcEndpointInfo) SyncFieldsDuringRead(ctx context.Context, from GcpVpcEndpointInfo) {
}

func (m GcpVpcEndpointInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GcpVpcEndpointInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GcpVpcEndpointInfo
// only implements ToObjectValue() and Type().
func (m GcpVpcEndpointInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_region":       m.EndpointRegion,
			"project_id":            m.ProjectId,
			"psc_connection_id":     m.PscConnectionId,
			"psc_endpoint_name":     m.PscEndpointName,
			"service_attachment_id": m.ServiceAttachmentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GcpVpcEndpointInfo) Type(ctx context.Context) attr.Type {
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
	// Credential configuration ID
	CredentialsId types.String `tfsdk:"-"`
}

func (to *GetCredentialRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCredentialRequest) {
}

func (to *GetCredentialRequest) SyncFieldsDuringRead(ctx context.Context, from GetCredentialRequest) {
}

func (m GetCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credentials_id"] = attrs["credentials_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialRequest
// only implements ToObjectValue() and Type().
func (m GetCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials_id": m.CredentialsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCredentialRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetEncryptionKeyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEncryptionKeyRequest) {
}

func (to *GetEncryptionKeyRequest) SyncFieldsDuringRead(ctx context.Context, from GetEncryptionKeyRequest) {
}

func (m GetEncryptionKeyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["customer_managed_key_id"] = attrs["customer_managed_key_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEncryptionKeyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEncryptionKeyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEncryptionKeyRequest
// only implements ToObjectValue() and Type().
func (m GetEncryptionKeyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_managed_key_id": m.CustomerManagedKeyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEncryptionKeyRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetNetworkRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetNetworkRequest) {
}

func (to *GetNetworkRequest) SyncFieldsDuringRead(ctx context.Context, from GetNetworkRequest) {
}

func (m GetNetworkRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_id"] = attrs["network_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNetworkRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetNetworkRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkRequest
// only implements ToObjectValue() and Type().
func (m GetNetworkRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_id": m.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetNetworkRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_id": types.StringType,
		},
	}
}

type GetPrivateAccesRequest struct {
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (to *GetPrivateAccesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPrivateAccesRequest) {
}

func (to *GetPrivateAccesRequest) SyncFieldsDuringRead(ctx context.Context, from GetPrivateAccesRequest) {
}

func (m GetPrivateAccesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPrivateAccesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetPrivateAccesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPrivateAccesRequest
// only implements ToObjectValue() and Type().
func (m GetPrivateAccesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_access_settings_id": m.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPrivateAccesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_access_settings_id": types.StringType,
		},
	}
}

type GetStorageRequest struct {
	StorageConfigurationId types.String `tfsdk:"-"`
}

func (to *GetStorageRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStorageRequest) {
}

func (to *GetStorageRequest) SyncFieldsDuringRead(ctx context.Context, from GetStorageRequest) {
}

func (m GetStorageRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetStorageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStorageRequest
// only implements ToObjectValue() and Type().
func (m GetStorageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_configuration_id": m.StorageConfigurationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStorageRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetVpcEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetVpcEndpointRequest) {
}

func (to *GetVpcEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from GetVpcEndpointRequest) {
}

func (m GetVpcEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["vpc_endpoint_id"] = attrs["vpc_endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetVpcEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetVpcEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetVpcEndpointRequest
// only implements ToObjectValue() and Type().
func (m GetVpcEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vpc_endpoint_id": m.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetVpcEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vpc_endpoint_id": types.StringType,
		},
	}
}

type GetWorkspaceRequest struct {
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceRequest) {
}

func (to *GetWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceRequest) {
}

func (m GetWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

// The configurations of the GKE cluster used by the GCP workspace.
type GkeConfig struct {
	// The type of network connectivity of the GKE cluster.
	ConnectivityType types.String `tfsdk:"connectivity_type"`
	// The IP range that will be used to allocate GKE cluster master resources
	// from. This field must not be set if
	// gke_cluster_type=PUBLIC_NODE_PUBLIC_MASTER.
	MasterIpRange types.String `tfsdk:"master_ip_range"`
}

func (to *GkeConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GkeConfig) {
}

func (to *GkeConfig) SyncFieldsDuringRead(ctx context.Context, from GkeConfig) {
}

func (m GkeConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GkeConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GkeConfig
// only implements ToObjectValue() and Type().
func (m GkeConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connectivity_type": m.ConnectivityType,
			"master_ip_range":   m.MasterIpRange,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GkeConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connectivity_type": types.StringType,
			"master_ip_range":   types.StringType,
		},
	}
}

// The credential ID that is used to access the key vault.
type KeyAccessConfiguration struct {
	CredentialId types.String `tfsdk:"credential_id"`
}

func (to *KeyAccessConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KeyAccessConfiguration) {
}

func (to *KeyAccessConfiguration) SyncFieldsDuringRead(ctx context.Context, from KeyAccessConfiguration) {
}

func (m KeyAccessConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KeyAccessConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KeyAccessConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KeyAccessConfiguration
// only implements ToObjectValue() and Type().
func (m KeyAccessConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": m.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KeyAccessConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.StringType,
		},
	}
}

type ListCredentialsRequest struct {
}

func (to *ListCredentialsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCredentialsRequest) {
}

func (to *ListCredentialsRequest) SyncFieldsDuringRead(ctx context.Context, from ListCredentialsRequest) {
}

func (m ListCredentialsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsRequest
// only implements ToObjectValue() and Type().
func (m ListCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListCredentialsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListEncryptionKeysRequest struct {
}

func (to *ListEncryptionKeysRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEncryptionKeysRequest) {
}

func (to *ListEncryptionKeysRequest) SyncFieldsDuringRead(ctx context.Context, from ListEncryptionKeysRequest) {
}

func (m ListEncryptionKeysRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEncryptionKeysRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEncryptionKeysRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEncryptionKeysRequest
// only implements ToObjectValue() and Type().
func (m ListEncryptionKeysRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListEncryptionKeysRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListNetworksRequest struct {
}

func (to *ListNetworksRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworksRequest) {
}

func (to *ListNetworksRequest) SyncFieldsDuringRead(ctx context.Context, from ListNetworksRequest) {
}

func (m ListNetworksRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworksRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListNetworksRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworksRequest
// only implements ToObjectValue() and Type().
func (m ListNetworksRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListNetworksRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListPrivateAccessRequest struct {
}

func (to *ListPrivateAccessRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPrivateAccessRequest) {
}

func (to *ListPrivateAccessRequest) SyncFieldsDuringRead(ctx context.Context, from ListPrivateAccessRequest) {
}

func (m ListPrivateAccessRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPrivateAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListPrivateAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateAccessRequest
// only implements ToObjectValue() and Type().
func (m ListPrivateAccessRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListPrivateAccessRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListStorageRequest struct {
}

func (to *ListStorageRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListStorageRequest) {
}

func (to *ListStorageRequest) SyncFieldsDuringRead(ctx context.Context, from ListStorageRequest) {
}

func (m ListStorageRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStorageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListStorageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStorageRequest
// only implements ToObjectValue() and Type().
func (m ListStorageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListStorageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListVpcEndpointsRequest struct {
}

func (to *ListVpcEndpointsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVpcEndpointsRequest) {
}

func (to *ListVpcEndpointsRequest) SyncFieldsDuringRead(ctx context.Context, from ListVpcEndpointsRequest) {
}

func (m ListVpcEndpointsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVpcEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListVpcEndpointsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVpcEndpointsRequest
// only implements ToObjectValue() and Type().
func (m ListVpcEndpointsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListVpcEndpointsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListWorkspacesRequest struct {
}

func (to *ListWorkspacesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspacesRequest) {
}

func (to *ListWorkspacesRequest) SyncFieldsDuringRead(ctx context.Context, from ListWorkspacesRequest) {
}

func (m ListWorkspacesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspacesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspacesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspacesRequest
// only implements ToObjectValue() and Type().
func (m ListWorkspacesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspacesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Network struct {
	// Time in epoch milliseconds when the network was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Array of error messages about the network configuration.
	ErrorMessages types.List `tfsdk:"error_messages"`

	GcpNetworkInfo types.Object `tfsdk:"gcp_network_info"`
	// The Databricks network configuration ID.
	NetworkId types.String `tfsdk:"network_id"`
	// The human-readable name of the network configuration.
	NetworkName types.String `tfsdk:"network_name"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds types.List `tfsdk:"security_group_ids"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
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

func (to *Network) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Network) {
	if !from.ErrorMessages.IsNull() && !from.ErrorMessages.IsUnknown() && to.ErrorMessages.IsNull() && len(from.ErrorMessages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ErrorMessages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ErrorMessages = from.ErrorMessages
	}
	if !from.GcpNetworkInfo.IsNull() && !from.GcpNetworkInfo.IsUnknown() {
		if toGcpNetworkInfo, ok := to.GetGcpNetworkInfo(ctx); ok {
			if fromGcpNetworkInfo, ok := from.GetGcpNetworkInfo(ctx); ok {
				// Recursively sync the fields of GcpNetworkInfo
				toGcpNetworkInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpNetworkInfo)
				to.SetGcpNetworkInfo(ctx, toGcpNetworkInfo)
			}
		}
	}
	if !from.SecurityGroupIds.IsNull() && !from.SecurityGroupIds.IsUnknown() && to.SecurityGroupIds.IsNull() && len(from.SecurityGroupIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecurityGroupIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecurityGroupIds = from.SecurityGroupIds
	}
	if !from.SubnetIds.IsNull() && !from.SubnetIds.IsUnknown() && to.SubnetIds.IsNull() && len(from.SubnetIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SubnetIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SubnetIds = from.SubnetIds
	}
	if !from.VpcEndpoints.IsNull() && !from.VpcEndpoints.IsUnknown() {
		if toVpcEndpoints, ok := to.GetVpcEndpoints(ctx); ok {
			if fromVpcEndpoints, ok := from.GetVpcEndpoints(ctx); ok {
				// Recursively sync the fields of VpcEndpoints
				toVpcEndpoints.SyncFieldsDuringCreateOrUpdate(ctx, fromVpcEndpoints)
				to.SetVpcEndpoints(ctx, toVpcEndpoints)
			}
		}
	}
	if !from.WarningMessages.IsNull() && !from.WarningMessages.IsUnknown() && to.WarningMessages.IsNull() && len(from.WarningMessages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WarningMessages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WarningMessages = from.WarningMessages
	}
}

func (to *Network) SyncFieldsDuringRead(ctx context.Context, from Network) {
	if !from.ErrorMessages.IsNull() && !from.ErrorMessages.IsUnknown() && to.ErrorMessages.IsNull() && len(from.ErrorMessages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ErrorMessages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ErrorMessages = from.ErrorMessages
	}
	if !from.GcpNetworkInfo.IsNull() && !from.GcpNetworkInfo.IsUnknown() {
		if toGcpNetworkInfo, ok := to.GetGcpNetworkInfo(ctx); ok {
			if fromGcpNetworkInfo, ok := from.GetGcpNetworkInfo(ctx); ok {
				toGcpNetworkInfo.SyncFieldsDuringRead(ctx, fromGcpNetworkInfo)
				to.SetGcpNetworkInfo(ctx, toGcpNetworkInfo)
			}
		}
	}
	if !from.SecurityGroupIds.IsNull() && !from.SecurityGroupIds.IsUnknown() && to.SecurityGroupIds.IsNull() && len(from.SecurityGroupIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecurityGroupIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecurityGroupIds = from.SecurityGroupIds
	}
	if !from.SubnetIds.IsNull() && !from.SubnetIds.IsUnknown() && to.SubnetIds.IsNull() && len(from.SubnetIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SubnetIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SubnetIds = from.SubnetIds
	}
	if !from.VpcEndpoints.IsNull() && !from.VpcEndpoints.IsUnknown() {
		if toVpcEndpoints, ok := to.GetVpcEndpoints(ctx); ok {
			if fromVpcEndpoints, ok := from.GetVpcEndpoints(ctx); ok {
				toVpcEndpoints.SyncFieldsDuringRead(ctx, fromVpcEndpoints)
				to.SetVpcEndpoints(ctx, toVpcEndpoints)
			}
		}
	}
	if !from.WarningMessages.IsNull() && !from.WarningMessages.IsUnknown() && to.WarningMessages.IsNull() && len(from.WarningMessages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for WarningMessages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.WarningMessages = from.WarningMessages
	}
}

func (m Network) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Network) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m Network) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_time":      m.CreationTime,
			"error_messages":     m.ErrorMessages,
			"gcp_network_info":   m.GcpNetworkInfo,
			"network_id":         m.NetworkId,
			"network_name":       m.NetworkName,
			"security_group_ids": m.SecurityGroupIds,
			"subnet_ids":         m.SubnetIds,
			"vpc_endpoints":      m.VpcEndpoints,
			"vpc_id":             m.VpcId,
			"vpc_status":         m.VpcStatus,
			"warning_messages":   m.WarningMessages,
			"workspace_id":       m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Network) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
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
func (m *Network) GetErrorMessages(ctx context.Context) ([]NetworkHealth, bool) {
	if m.ErrorMessages.IsNull() || m.ErrorMessages.IsUnknown() {
		return nil, false
	}
	var v []NetworkHealth
	d := m.ErrorMessages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetErrorMessages sets the value of the ErrorMessages field in Network.
func (m *Network) SetErrorMessages(ctx context.Context, v []NetworkHealth) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["error_messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ErrorMessages = types.ListValueMust(t, vs)
}

// GetGcpNetworkInfo returns the value of the GcpNetworkInfo field in Network as
// a GcpNetworkInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *Network) GetGcpNetworkInfo(ctx context.Context) (GcpNetworkInfo, bool) {
	var e GcpNetworkInfo
	if m.GcpNetworkInfo.IsNull() || m.GcpNetworkInfo.IsUnknown() {
		return e, false
	}
	var v GcpNetworkInfo
	d := m.GcpNetworkInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpNetworkInfo sets the value of the GcpNetworkInfo field in Network.
func (m *Network) SetGcpNetworkInfo(ctx context.Context, v GcpNetworkInfo) {
	vs := v.ToObjectValue(ctx)
	m.GcpNetworkInfo = vs
}

// GetSecurityGroupIds returns the value of the SecurityGroupIds field in Network as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Network) GetSecurityGroupIds(ctx context.Context) ([]types.String, bool) {
	if m.SecurityGroupIds.IsNull() || m.SecurityGroupIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SecurityGroupIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecurityGroupIds sets the value of the SecurityGroupIds field in Network.
func (m *Network) SetSecurityGroupIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["security_group_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SecurityGroupIds = types.ListValueMust(t, vs)
}

// GetSubnetIds returns the value of the SubnetIds field in Network as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Network) GetSubnetIds(ctx context.Context) ([]types.String, bool) {
	if m.SubnetIds.IsNull() || m.SubnetIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SubnetIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnetIds sets the value of the SubnetIds field in Network.
func (m *Network) SetSubnetIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["subnet_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SubnetIds = types.ListValueMust(t, vs)
}

// GetVpcEndpoints returns the value of the VpcEndpoints field in Network as
// a NetworkVpcEndpoints value.
// If the field is unknown or null, the boolean return value is false.
func (m *Network) GetVpcEndpoints(ctx context.Context) (NetworkVpcEndpoints, bool) {
	var e NetworkVpcEndpoints
	if m.VpcEndpoints.IsNull() || m.VpcEndpoints.IsUnknown() {
		return e, false
	}
	var v NetworkVpcEndpoints
	d := m.VpcEndpoints.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVpcEndpoints sets the value of the VpcEndpoints field in Network.
func (m *Network) SetVpcEndpoints(ctx context.Context, v NetworkVpcEndpoints) {
	vs := v.ToObjectValue(ctx)
	m.VpcEndpoints = vs
}

// GetWarningMessages returns the value of the WarningMessages field in Network as
// a slice of NetworkWarning values.
// If the field is unknown or null, the boolean return value is false.
func (m *Network) GetWarningMessages(ctx context.Context) ([]NetworkWarning, bool) {
	if m.WarningMessages.IsNull() || m.WarningMessages.IsUnknown() {
		return nil, false
	}
	var v []NetworkWarning
	d := m.WarningMessages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWarningMessages sets the value of the WarningMessages field in Network.
func (m *Network) SetWarningMessages(ctx context.Context, v []NetworkWarning) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["warning_messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WarningMessages = types.ListValueMust(t, vs)
}

type NetworkHealth struct {
	// Details of the error.
	ErrorMessage types.String `tfsdk:"error_message"`

	ErrorType types.String `tfsdk:"error_type"`
}

func (to *NetworkHealth) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkHealth) {
}

func (to *NetworkHealth) SyncFieldsDuringRead(ctx context.Context, from NetworkHealth) {
}

func (m NetworkHealth) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NetworkHealth) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkHealth
// only implements ToObjectValue() and Type().
func (m NetworkHealth) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"error_message": m.ErrorMessage,
			"error_type":    m.ErrorType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NetworkHealth) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error_message": types.StringType,
			"error_type":    types.StringType,
		},
	}
}

type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	DataplaneRelay types.List `tfsdk:"dataplane_relay"`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi types.List `tfsdk:"rest_api"`
}

func (to *NetworkVpcEndpoints) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkVpcEndpoints) {
	if !from.DataplaneRelay.IsNull() && !from.DataplaneRelay.IsUnknown() && to.DataplaneRelay.IsNull() && len(from.DataplaneRelay.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataplaneRelay, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataplaneRelay = from.DataplaneRelay
	}
	if !from.RestApi.IsNull() && !from.RestApi.IsUnknown() && to.RestApi.IsNull() && len(from.RestApi.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RestApi, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RestApi = from.RestApi
	}
}

func (to *NetworkVpcEndpoints) SyncFieldsDuringRead(ctx context.Context, from NetworkVpcEndpoints) {
	if !from.DataplaneRelay.IsNull() && !from.DataplaneRelay.IsUnknown() && to.DataplaneRelay.IsNull() && len(from.DataplaneRelay.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataplaneRelay, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataplaneRelay = from.DataplaneRelay
	}
	if !from.RestApi.IsNull() && !from.RestApi.IsUnknown() && to.RestApi.IsNull() && len(from.RestApi.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RestApi, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RestApi = from.RestApi
	}
}

func (m NetworkVpcEndpoints) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dataplane_relay"] = attrs["dataplane_relay"].SetOptional()
	attrs["rest_api"] = attrs["rest_api"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkVpcEndpoints.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NetworkVpcEndpoints) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataplane_relay": reflect.TypeOf(types.String{}),
		"rest_api":        reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkVpcEndpoints
// only implements ToObjectValue() and Type().
func (m NetworkVpcEndpoints) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataplane_relay": m.DataplaneRelay,
			"rest_api":        m.RestApi,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NetworkVpcEndpoints) Type(ctx context.Context) attr.Type {
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
func (m *NetworkVpcEndpoints) GetDataplaneRelay(ctx context.Context) ([]types.String, bool) {
	if m.DataplaneRelay.IsNull() || m.DataplaneRelay.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DataplaneRelay.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataplaneRelay sets the value of the DataplaneRelay field in NetworkVpcEndpoints.
func (m *NetworkVpcEndpoints) SetDataplaneRelay(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dataplane_relay"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DataplaneRelay = types.ListValueMust(t, vs)
}

// GetRestApi returns the value of the RestApi field in NetworkVpcEndpoints as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NetworkVpcEndpoints) GetRestApi(ctx context.Context) ([]types.String, bool) {
	if m.RestApi.IsNull() || m.RestApi.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.RestApi.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestApi sets the value of the RestApi field in NetworkVpcEndpoints.
func (m *NetworkVpcEndpoints) SetRestApi(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rest_api"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RestApi = types.ListValueMust(t, vs)
}

type NetworkWarning struct {
	// Details of the warning.
	WarningMessage types.String `tfsdk:"warning_message"`

	WarningType types.String `tfsdk:"warning_type"`
}

func (to *NetworkWarning) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkWarning) {
}

func (to *NetworkWarning) SyncFieldsDuringRead(ctx context.Context, from NetworkWarning) {
}

func (m NetworkWarning) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NetworkWarning) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkWarning
// only implements ToObjectValue() and Type().
func (m NetworkWarning) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"warning_message": m.WarningMessage,
			"warning_type":    m.WarningType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NetworkWarning) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"warning_message": types.StringType,
			"warning_type":    types.StringType,
		},
	}
}

// *
type PrivateAccessSettings struct {
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID that
	// is returned when registering the VPC endpoint configuration in your
	// Databricks account. This is not the ID of the VPC endpoint in AWS. Only
	// used when private_access_level is set to ENDPOINT. This is an allow list
	// of VPC endpoints that in your account that can connect to your workspace
	// over AWS PrivateLink. If hybrid access to your workspace is enabled by
	// setting public_access_enabled to true, this control only works for
	// PrivateLink connections. To control how your workspace is accessed via
	// public internet, see IP access lists.
	AllowedVpcEndpointIds types.List `tfsdk:"allowed_vpc_endpoint_ids"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see allowed_vpc_endpoint_ids.
	PrivateAccessLevel types.String `tfsdk:"private_access_level"`
	// Databricks private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"private_access_settings_id"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName types.String `tfsdk:"private_access_settings_name"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify false, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify true, which means that public access is
	// enabled.
	PublicAccessEnabled types.Bool `tfsdk:"public_access_enabled"`
	// The AWS region for workspaces attached to this private access settings
	// object.
	Region types.String `tfsdk:"region"`
}

func (to *PrivateAccessSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PrivateAccessSettings) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (to *PrivateAccessSettings) SyncFieldsDuringRead(ctx context.Context, from PrivateAccessSettings) {
	if !from.AllowedVpcEndpointIds.IsNull() && !from.AllowedVpcEndpointIds.IsUnknown() && to.AllowedVpcEndpointIds.IsNull() && len(from.AllowedVpcEndpointIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedVpcEndpointIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedVpcEndpointIds = from.AllowedVpcEndpointIds
	}
}

func (m PrivateAccessSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allowed_vpc_endpoint_ids"] = attrs["allowed_vpc_endpoint_ids"].SetOptional()
	attrs["private_access_level"] = attrs["private_access_level"].SetOptional()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetComputed()
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
func (m PrivateAccessSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_vpc_endpoint_ids": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PrivateAccessSettings
// only implements ToObjectValue() and Type().
func (m PrivateAccessSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_vpc_endpoint_ids":     m.AllowedVpcEndpointIds,
			"private_access_level":         m.PrivateAccessLevel,
			"private_access_settings_id":   m.PrivateAccessSettingsId,
			"private_access_settings_name": m.PrivateAccessSettingsName,
			"public_access_enabled":        m.PublicAccessEnabled,
			"region":                       m.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PrivateAccessSettings) Type(ctx context.Context) attr.Type {
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

// GetAllowedVpcEndpointIds returns the value of the AllowedVpcEndpointIds field in PrivateAccessSettings as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PrivateAccessSettings) GetAllowedVpcEndpointIds(ctx context.Context) ([]types.String, bool) {
	if m.AllowedVpcEndpointIds.IsNull() || m.AllowedVpcEndpointIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.AllowedVpcEndpointIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedVpcEndpointIds sets the value of the AllowedVpcEndpointIds field in PrivateAccessSettings.
func (m *PrivateAccessSettings) SetAllowedVpcEndpointIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_vpc_endpoint_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedVpcEndpointIds = types.ListValueMust(t, vs)
}

type ReplacePrivateAccessSettingsRequest struct {
	// Properties of the new private access settings object.
	CustomerFacingPrivateAccessSettings types.Object `tfsdk:"customer_facing_private_access_settings"`
	// Databricks private access settings ID.
	PrivateAccessSettingsId types.String `tfsdk:"-"`
}

func (to *ReplacePrivateAccessSettingsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReplacePrivateAccessSettingsRequest) {
	if !from.CustomerFacingPrivateAccessSettings.IsNull() && !from.CustomerFacingPrivateAccessSettings.IsUnknown() {
		if toCustomerFacingPrivateAccessSettings, ok := to.GetCustomerFacingPrivateAccessSettings(ctx); ok {
			if fromCustomerFacingPrivateAccessSettings, ok := from.GetCustomerFacingPrivateAccessSettings(ctx); ok {
				// Recursively sync the fields of CustomerFacingPrivateAccessSettings
				toCustomerFacingPrivateAccessSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromCustomerFacingPrivateAccessSettings)
				to.SetCustomerFacingPrivateAccessSettings(ctx, toCustomerFacingPrivateAccessSettings)
			}
		}
	}
}

func (to *ReplacePrivateAccessSettingsRequest) SyncFieldsDuringRead(ctx context.Context, from ReplacePrivateAccessSettingsRequest) {
	if !from.CustomerFacingPrivateAccessSettings.IsNull() && !from.CustomerFacingPrivateAccessSettings.IsUnknown() {
		if toCustomerFacingPrivateAccessSettings, ok := to.GetCustomerFacingPrivateAccessSettings(ctx); ok {
			if fromCustomerFacingPrivateAccessSettings, ok := from.GetCustomerFacingPrivateAccessSettings(ctx); ok {
				toCustomerFacingPrivateAccessSettings.SyncFieldsDuringRead(ctx, fromCustomerFacingPrivateAccessSettings)
				to.SetCustomerFacingPrivateAccessSettings(ctx, toCustomerFacingPrivateAccessSettings)
			}
		}
	}
}

func (m ReplacePrivateAccessSettingsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["customer_facing_private_access_settings"] = attrs["customer_facing_private_access_settings"].SetRequired()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplacePrivateAccessSettingsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ReplacePrivateAccessSettingsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"customer_facing_private_access_settings": reflect.TypeOf(PrivateAccessSettings{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplacePrivateAccessSettingsRequest
// only implements ToObjectValue() and Type().
func (m ReplacePrivateAccessSettingsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_facing_private_access_settings": m.CustomerFacingPrivateAccessSettings,
			"private_access_settings_id":              m.PrivateAccessSettingsId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ReplacePrivateAccessSettingsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_facing_private_access_settings": PrivateAccessSettings{}.Type(ctx),
			"private_access_settings_id":              types.StringType,
		},
	}
}

// GetCustomerFacingPrivateAccessSettings returns the value of the CustomerFacingPrivateAccessSettings field in ReplacePrivateAccessSettingsRequest as
// a PrivateAccessSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *ReplacePrivateAccessSettingsRequest) GetCustomerFacingPrivateAccessSettings(ctx context.Context) (PrivateAccessSettings, bool) {
	var e PrivateAccessSettings
	if m.CustomerFacingPrivateAccessSettings.IsNull() || m.CustomerFacingPrivateAccessSettings.IsUnknown() {
		return e, false
	}
	var v PrivateAccessSettings
	d := m.CustomerFacingPrivateAccessSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomerFacingPrivateAccessSettings sets the value of the CustomerFacingPrivateAccessSettings field in ReplacePrivateAccessSettingsRequest.
func (m *ReplacePrivateAccessSettingsRequest) SetCustomerFacingPrivateAccessSettings(ctx context.Context, v PrivateAccessSettings) {
	vs := v.ToObjectValue(ctx)
	m.CustomerFacingPrivateAccessSettings = vs
}

type RootBucketInfo struct {
	// Name of the S3 bucket
	BucketName types.String `tfsdk:"bucket_name"`
}

func (to *RootBucketInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RootBucketInfo) {
}

func (to *RootBucketInfo) SyncFieldsDuringRead(ctx context.Context, from RootBucketInfo) {
}

func (m RootBucketInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RootBucketInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RootBucketInfo
// only implements ToObjectValue() and Type().
func (m RootBucketInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bucket_name": m.BucketName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RootBucketInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bucket_name": types.StringType,
		},
	}
}

type StorageConfiguration struct {
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Optional IAM role that is used to access the workspace catalog which is
	// created during workspace creation for UC by Default. If a storage
	// configuration with this field populated is used to create a workspace,
	// then a workspace catalog is created together with the workspace. The
	// workspace catalog shares the root bucket with internal workspace storage
	// (including DBFS root) but uses a dedicated bucket path prefix.
	RoleArn types.String `tfsdk:"role_arn"`
	// The root bucket information for the storage configuration.
	RootBucketInfo types.Object `tfsdk:"root_bucket_info"`
	// Databricks storage configuration ID.
	StorageConfigurationId types.String `tfsdk:"storage_configuration_id"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName types.String `tfsdk:"storage_configuration_name"`
}

func (to *StorageConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StorageConfiguration) {
	if !from.RootBucketInfo.IsNull() && !from.RootBucketInfo.IsUnknown() {
		if toRootBucketInfo, ok := to.GetRootBucketInfo(ctx); ok {
			if fromRootBucketInfo, ok := from.GetRootBucketInfo(ctx); ok {
				// Recursively sync the fields of RootBucketInfo
				toRootBucketInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromRootBucketInfo)
				to.SetRootBucketInfo(ctx, toRootBucketInfo)
			}
		}
	}
}

func (to *StorageConfiguration) SyncFieldsDuringRead(ctx context.Context, from StorageConfiguration) {
	if !from.RootBucketInfo.IsNull() && !from.RootBucketInfo.IsUnknown() {
		if toRootBucketInfo, ok := to.GetRootBucketInfo(ctx); ok {
			if fromRootBucketInfo, ok := from.GetRootBucketInfo(ctx); ok {
				toRootBucketInfo.SyncFieldsDuringRead(ctx, fromRootBucketInfo)
				to.SetRootBucketInfo(ctx, toRootBucketInfo)
			}
		}
	}
}

func (m StorageConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["role_arn"] = attrs["role_arn"].SetOptional()
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
func (m StorageConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"root_bucket_info": reflect.TypeOf(RootBucketInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StorageConfiguration
// only implements ToObjectValue() and Type().
func (m StorageConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_time":              m.CreationTime,
			"role_arn":                   m.RoleArn,
			"root_bucket_info":           m.RootBucketInfo,
			"storage_configuration_id":   m.StorageConfigurationId,
			"storage_configuration_name": m.StorageConfigurationName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StorageConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_time":              types.Int64Type,
			"role_arn":                   types.StringType,
			"root_bucket_info":           RootBucketInfo{}.Type(ctx),
			"storage_configuration_id":   types.StringType,
			"storage_configuration_name": types.StringType,
		},
	}
}

// GetRootBucketInfo returns the value of the RootBucketInfo field in StorageConfiguration as
// a RootBucketInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *StorageConfiguration) GetRootBucketInfo(ctx context.Context) (RootBucketInfo, bool) {
	var e RootBucketInfo
	if m.RootBucketInfo.IsNull() || m.RootBucketInfo.IsUnknown() {
		return e, false
	}
	var v RootBucketInfo
	d := m.RootBucketInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRootBucketInfo sets the value of the RootBucketInfo field in StorageConfiguration.
func (m *StorageConfiguration) SetRootBucketInfo(ctx context.Context, v RootBucketInfo) {
	vs := v.ToObjectValue(ctx)
	m.RootBucketInfo = vs
}

type StsRole struct {
	// The Amazon Resource Name (ARN) of the cross account IAM role.
	RoleArn types.String `tfsdk:"role_arn"`
}

func (to *StsRole) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StsRole) {
}

func (to *StsRole) SyncFieldsDuringRead(ctx context.Context, from StsRole) {
}

func (m StsRole) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StsRole) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StsRole
// only implements ToObjectValue() and Type().
func (m StsRole) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"role_arn": m.RoleArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StsRole) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"role_arn": types.StringType,
		},
	}
}

type UpdateWorkspaceRequest struct {
	CustomerFacingWorkspace types.Object `tfsdk:"customer_facing_workspace"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"-"`
	// A unique integer ID for the workspace
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *UpdateWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceRequest) {
	if !from.CustomerFacingWorkspace.IsNull() && !from.CustomerFacingWorkspace.IsUnknown() {
		if toCustomerFacingWorkspace, ok := to.GetCustomerFacingWorkspace(ctx); ok {
			if fromCustomerFacingWorkspace, ok := from.GetCustomerFacingWorkspace(ctx); ok {
				// Recursively sync the fields of CustomerFacingWorkspace
				toCustomerFacingWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromCustomerFacingWorkspace)
				to.SetCustomerFacingWorkspace(ctx, toCustomerFacingWorkspace)
			}
		}
	}
}

func (to *UpdateWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceRequest) {
	if !from.CustomerFacingWorkspace.IsNull() && !from.CustomerFacingWorkspace.IsUnknown() {
		if toCustomerFacingWorkspace, ok := to.GetCustomerFacingWorkspace(ctx); ok {
			if fromCustomerFacingWorkspace, ok := from.GetCustomerFacingWorkspace(ctx); ok {
				toCustomerFacingWorkspace.SyncFieldsDuringRead(ctx, fromCustomerFacingWorkspace)
				to.SetCustomerFacingWorkspace(ctx, toCustomerFacingWorkspace)
			}
		}
	}
}

func (m UpdateWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["customer_facing_workspace"] = attrs["customer_facing_workspace"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
	attrs["update_mask"] = attrs["update_mask"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"customer_facing_workspace": reflect.TypeOf(Workspace{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceRequest
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"customer_facing_workspace": m.CustomerFacingWorkspace,
			"update_mask":               m.UpdateMask,
			"workspace_id":              m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"customer_facing_workspace": Workspace{}.Type(ctx),
			"update_mask":               types.StringType,
			"workspace_id":              types.Int64Type,
		},
	}
}

// GetCustomerFacingWorkspace returns the value of the CustomerFacingWorkspace field in UpdateWorkspaceRequest as
// a Workspace value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceRequest) GetCustomerFacingWorkspace(ctx context.Context) (Workspace, bool) {
	var e Workspace
	if m.CustomerFacingWorkspace.IsNull() || m.CustomerFacingWorkspace.IsUnknown() {
		return e, false
	}
	var v Workspace
	d := m.CustomerFacingWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomerFacingWorkspace sets the value of the CustomerFacingWorkspace field in UpdateWorkspaceRequest.
func (m *UpdateWorkspaceRequest) SetCustomerFacingWorkspace(ctx context.Context, v Workspace) {
	vs := v.ToObjectValue(ctx)
	m.CustomerFacingWorkspace = vs
}

// *
type VpcEndpoint struct {
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
	// The cloud info of this vpc endpoint. Info for a GCP vpc endpoint.
	GcpVpcEndpointInfo types.Object `tfsdk:"gcp_vpc_endpoint_info"`
	// The AWS region in which this VPC endpoint object exists.
	Region types.String `tfsdk:"region"`
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	State types.String `tfsdk:"state"`
	// This enumeration represents the type of Databricks VPC endpoint service
	// that was used when creating this VPC endpoint. If the VPC endpoint
	// connects to the Databricks control plane for either the front-end
	// connection or the back-end REST API connection, the value is
	// WORKSPACE_ACCESS. If the VPC endpoint connects to the Databricks
	// workspace for the back-end secure cluster connectivity relay, the value
	// is DATAPLANE_RELAY_ACCESS.
	UseCase types.String `tfsdk:"use_case"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId types.String `tfsdk:"vpc_endpoint_id"`
	// The human-readable name of the storage configuration.
	VpcEndpointName types.String `tfsdk:"vpc_endpoint_name"`
}

func (to *VpcEndpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VpcEndpoint) {
	if !from.GcpVpcEndpointInfo.IsNull() && !from.GcpVpcEndpointInfo.IsUnknown() {
		if toGcpVpcEndpointInfo, ok := to.GetGcpVpcEndpointInfo(ctx); ok {
			if fromGcpVpcEndpointInfo, ok := from.GetGcpVpcEndpointInfo(ctx); ok {
				// Recursively sync the fields of GcpVpcEndpointInfo
				toGcpVpcEndpointInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpVpcEndpointInfo)
				to.SetGcpVpcEndpointInfo(ctx, toGcpVpcEndpointInfo)
			}
		}
	}
}

func (to *VpcEndpoint) SyncFieldsDuringRead(ctx context.Context, from VpcEndpoint) {
	if !from.GcpVpcEndpointInfo.IsNull() && !from.GcpVpcEndpointInfo.IsUnknown() {
		if toGcpVpcEndpointInfo, ok := to.GetGcpVpcEndpointInfo(ctx); ok {
			if fromGcpVpcEndpointInfo, ok := from.GetGcpVpcEndpointInfo(ctx); ok {
				toGcpVpcEndpointInfo.SyncFieldsDuringRead(ctx, fromGcpVpcEndpointInfo)
				to.SetGcpVpcEndpointInfo(ctx, toGcpVpcEndpointInfo)
			}
		}
	}
}

func (m VpcEndpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m VpcEndpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_vpc_endpoint_info": reflect.TypeOf(GcpVpcEndpointInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VpcEndpoint
// only implements ToObjectValue() and Type().
func (m VpcEndpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_account_id":          m.AwsAccountId,
			"aws_endpoint_service_id": m.AwsEndpointServiceId,
			"aws_vpc_endpoint_id":     m.AwsVpcEndpointId,
			"gcp_vpc_endpoint_info":   m.GcpVpcEndpointInfo,
			"region":                  m.Region,
			"state":                   m.State,
			"use_case":                m.UseCase,
			"vpc_endpoint_id":         m.VpcEndpointId,
			"vpc_endpoint_name":       m.VpcEndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VpcEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
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
func (m *VpcEndpoint) GetGcpVpcEndpointInfo(ctx context.Context) (GcpVpcEndpointInfo, bool) {
	var e GcpVpcEndpointInfo
	if m.GcpVpcEndpointInfo.IsNull() || m.GcpVpcEndpointInfo.IsUnknown() {
		return e, false
	}
	var v GcpVpcEndpointInfo
	d := m.GcpVpcEndpointInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpVpcEndpointInfo sets the value of the GcpVpcEndpointInfo field in VpcEndpoint.
func (m *VpcEndpoint) SetGcpVpcEndpointInfo(ctx context.Context, v GcpVpcEndpointInfo) {
	vs := v.ToObjectValue(ctx)
	m.GcpVpcEndpointInfo = vs
}

type Workspace struct {
	AwsRegion types.String `tfsdk:"aws_region"`

	AzureWorkspaceInfo types.Object `tfsdk:"azure_workspace_info"`
	// The cloud name. This field can have values like `azure`, `gcp`.
	Cloud types.String `tfsdk:"cloud"`

	CloudResourceContainer types.Object `tfsdk:"cloud_resource_container"`
	// The compute mode of the workspace.
	ComputeMode types.String `tfsdk:"compute_mode"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// ID of the workspace's credential configuration object.
	CredentialsId types.String `tfsdk:"credentials_id"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags types.Map `tfsdk:"custom_tags"`

	DeploymentName types.String `tfsdk:"deployment_name"`
	// A client owned field used to indicate the workspace status that the
	// client expects to be in. For now this is only used to unblock Temporal
	// workflow for GCP least privileged workspace.
	ExpectedWorkspaceStatus types.String `tfsdk:"expected_workspace_status"`

	GcpManagedNetworkConfig types.Object `tfsdk:"gcp_managed_network_config"`

	GkeConfig types.Object `tfsdk:"gke_config"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location types.String `tfsdk:"location"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId types.String `tfsdk:"managed_services_customer_managed_key_id"`
	// The network configuration for the workspace.
	//
	// DEPRECATED. Use `network_id` instead.
	Network types.Object `tfsdk:"network"`
	// The object ID of network connectivity config.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id"`
	// If this workspace is BYO VPC, then the network_id will be populated. If
	// this workspace is not BYO VPC, then the network_id will be empty.
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
	// The storage mode of the workspace.
	StorageMode types.String `tfsdk:"storage_mode"`
	// A unique integer ID for the workspace
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
	// The human-readable name of the workspace.
	WorkspaceName types.String `tfsdk:"workspace_name"`
	// The status of a workspace
	WorkspaceStatus types.String `tfsdk:"workspace_status"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage types.String `tfsdk:"workspace_status_message"`
}

func (to *Workspace) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Workspace) {
	if !from.AzureWorkspaceInfo.IsNull() && !from.AzureWorkspaceInfo.IsUnknown() {
		if toAzureWorkspaceInfo, ok := to.GetAzureWorkspaceInfo(ctx); ok {
			if fromAzureWorkspaceInfo, ok := from.GetAzureWorkspaceInfo(ctx); ok {
				// Recursively sync the fields of AzureWorkspaceInfo
				toAzureWorkspaceInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureWorkspaceInfo)
				to.SetAzureWorkspaceInfo(ctx, toAzureWorkspaceInfo)
			}
		}
	}
	if !from.CloudResourceContainer.IsNull() && !from.CloudResourceContainer.IsUnknown() {
		if toCloudResourceContainer, ok := to.GetCloudResourceContainer(ctx); ok {
			if fromCloudResourceContainer, ok := from.GetCloudResourceContainer(ctx); ok {
				// Recursively sync the fields of CloudResourceContainer
				toCloudResourceContainer.SyncFieldsDuringCreateOrUpdate(ctx, fromCloudResourceContainer)
				to.SetCloudResourceContainer(ctx, toCloudResourceContainer)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				// Recursively sync the fields of GcpManagedNetworkConfig
				toGcpManagedNetworkConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
	if !from.GkeConfig.IsNull() && !from.GkeConfig.IsUnknown() {
		if toGkeConfig, ok := to.GetGkeConfig(ctx); ok {
			if fromGkeConfig, ok := from.GetGkeConfig(ctx); ok {
				// Recursively sync the fields of GkeConfig
				toGkeConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGkeConfig)
				to.SetGkeConfig(ctx, toGkeConfig)
			}
		}
	}
	if !from.Network.IsNull() && !from.Network.IsUnknown() {
		if toNetwork, ok := to.GetNetwork(ctx); ok {
			if fromNetwork, ok := from.GetNetwork(ctx); ok {
				// Recursively sync the fields of Network
				toNetwork.SyncFieldsDuringCreateOrUpdate(ctx, fromNetwork)
				to.SetNetwork(ctx, toNetwork)
			}
		}
	}
}

func (to *Workspace) SyncFieldsDuringRead(ctx context.Context, from Workspace) {
	if !from.AzureWorkspaceInfo.IsNull() && !from.AzureWorkspaceInfo.IsUnknown() {
		if toAzureWorkspaceInfo, ok := to.GetAzureWorkspaceInfo(ctx); ok {
			if fromAzureWorkspaceInfo, ok := from.GetAzureWorkspaceInfo(ctx); ok {
				toAzureWorkspaceInfo.SyncFieldsDuringRead(ctx, fromAzureWorkspaceInfo)
				to.SetAzureWorkspaceInfo(ctx, toAzureWorkspaceInfo)
			}
		}
	}
	if !from.CloudResourceContainer.IsNull() && !from.CloudResourceContainer.IsUnknown() {
		if toCloudResourceContainer, ok := to.GetCloudResourceContainer(ctx); ok {
			if fromCloudResourceContainer, ok := from.GetCloudResourceContainer(ctx); ok {
				toCloudResourceContainer.SyncFieldsDuringRead(ctx, fromCloudResourceContainer)
				to.SetCloudResourceContainer(ctx, toCloudResourceContainer)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				toGcpManagedNetworkConfig.SyncFieldsDuringRead(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
	if !from.GkeConfig.IsNull() && !from.GkeConfig.IsUnknown() {
		if toGkeConfig, ok := to.GetGkeConfig(ctx); ok {
			if fromGkeConfig, ok := from.GetGkeConfig(ctx); ok {
				toGkeConfig.SyncFieldsDuringRead(ctx, fromGkeConfig)
				to.SetGkeConfig(ctx, toGkeConfig)
			}
		}
	}
	if !from.Network.IsNull() && !from.Network.IsUnknown() {
		if toNetwork, ok := to.GetNetwork(ctx); ok {
			if fromNetwork, ok := from.GetNetwork(ctx); ok {
				toNetwork.SyncFieldsDuringRead(ctx, fromNetwork)
				to.SetNetwork(ctx, toNetwork)
			}
		}
	}
}

func (m Workspace) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_region"] = attrs["aws_region"].SetOptional()
	attrs["azure_workspace_info"] = attrs["azure_workspace_info"].SetComputed()
	attrs["cloud"] = attrs["cloud"].SetOptional()
	attrs["cloud_resource_container"] = attrs["cloud_resource_container"].SetOptional()
	attrs["compute_mode"] = attrs["compute_mode"].SetComputed()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["credentials_id"] = attrs["credentials_id"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["deployment_name"] = attrs["deployment_name"].SetOptional()
	attrs["expected_workspace_status"] = attrs["expected_workspace_status"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].SetOptional()
	attrs["gke_config"] = attrs["gke_config"].SetOptional()
	attrs["location"] = attrs["location"].SetOptional()
	attrs["managed_services_customer_managed_key_id"] = attrs["managed_services_customer_managed_key_id"].SetOptional()
	attrs["network"] = attrs["network"].SetOptional()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetOptional()
	attrs["network_id"] = attrs["network_id"].SetOptional()
	attrs["pricing_tier"] = attrs["pricing_tier"].SetComputed()
	attrs["private_access_settings_id"] = attrs["private_access_settings_id"].SetOptional()
	attrs["storage_configuration_id"] = attrs["storage_configuration_id"].SetOptional()
	attrs["storage_customer_managed_key_id"] = attrs["storage_customer_managed_key_id"].SetOptional()
	attrs["storage_mode"] = attrs["storage_mode"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
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
func (m Workspace) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"azure_workspace_info":       reflect.TypeOf(AzureWorkspaceInfo{}),
		"cloud_resource_container":   reflect.TypeOf(CloudResourceContainer{}),
		"custom_tags":                reflect.TypeOf(types.String{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig{}),
		"gke_config":                 reflect.TypeOf(GkeConfig{}),
		"network":                    reflect.TypeOf(WorkspaceNetwork{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Workspace
// only implements ToObjectValue() and Type().
func (m Workspace) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_region":                 m.AwsRegion,
			"azure_workspace_info":       m.AzureWorkspaceInfo,
			"cloud":                      m.Cloud,
			"cloud_resource_container":   m.CloudResourceContainer,
			"compute_mode":               m.ComputeMode,
			"creation_time":              m.CreationTime,
			"credentials_id":             m.CredentialsId,
			"custom_tags":                m.CustomTags,
			"deployment_name":            m.DeploymentName,
			"expected_workspace_status":  m.ExpectedWorkspaceStatus,
			"gcp_managed_network_config": m.GcpManagedNetworkConfig,
			"gke_config":                 m.GkeConfig,
			"location":                   m.Location,
			"managed_services_customer_managed_key_id": m.ManagedServicesCustomerManagedKeyId,
			"network":                         m.Network,
			"network_connectivity_config_id":  m.NetworkConnectivityConfigId,
			"network_id":                      m.NetworkId,
			"pricing_tier":                    m.PricingTier,
			"private_access_settings_id":      m.PrivateAccessSettingsId,
			"storage_configuration_id":        m.StorageConfigurationId,
			"storage_customer_managed_key_id": m.StorageCustomerManagedKeyId,
			"storage_mode":                    m.StorageMode,
			"workspace_id":                    m.WorkspaceId,
			"workspace_name":                  m.WorkspaceName,
			"workspace_status":                m.WorkspaceStatus,
			"workspace_status_message":        m.WorkspaceStatusMessage,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Workspace) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_region":               types.StringType,
			"azure_workspace_info":     AzureWorkspaceInfo{}.Type(ctx),
			"cloud":                    types.StringType,
			"cloud_resource_container": CloudResourceContainer{}.Type(ctx),
			"compute_mode":             types.StringType,
			"creation_time":            types.Int64Type,
			"credentials_id":           types.StringType,
			"custom_tags": basetypes.MapType{
				ElemType: types.StringType,
			},
			"deployment_name":                          types.StringType,
			"expected_workspace_status":                types.StringType,
			"gcp_managed_network_config":               GcpManagedNetworkConfig{}.Type(ctx),
			"gke_config":                               GkeConfig{}.Type(ctx),
			"location":                                 types.StringType,
			"managed_services_customer_managed_key_id": types.StringType,
			"network":                                  WorkspaceNetwork{}.Type(ctx),
			"network_connectivity_config_id":           types.StringType,
			"network_id":                               types.StringType,
			"pricing_tier":                             types.StringType,
			"private_access_settings_id":               types.StringType,
			"storage_configuration_id":                 types.StringType,
			"storage_customer_managed_key_id":          types.StringType,
			"storage_mode":                             types.StringType,
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
func (m *Workspace) GetAzureWorkspaceInfo(ctx context.Context) (AzureWorkspaceInfo, bool) {
	var e AzureWorkspaceInfo
	if m.AzureWorkspaceInfo.IsNull() || m.AzureWorkspaceInfo.IsUnknown() {
		return e, false
	}
	var v AzureWorkspaceInfo
	d := m.AzureWorkspaceInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureWorkspaceInfo sets the value of the AzureWorkspaceInfo field in Workspace.
func (m *Workspace) SetAzureWorkspaceInfo(ctx context.Context, v AzureWorkspaceInfo) {
	vs := v.ToObjectValue(ctx)
	m.AzureWorkspaceInfo = vs
}

// GetCloudResourceContainer returns the value of the CloudResourceContainer field in Workspace as
// a CloudResourceContainer value.
// If the field is unknown or null, the boolean return value is false.
func (m *Workspace) GetCloudResourceContainer(ctx context.Context) (CloudResourceContainer, bool) {
	var e CloudResourceContainer
	if m.CloudResourceContainer.IsNull() || m.CloudResourceContainer.IsUnknown() {
		return e, false
	}
	var v CloudResourceContainer
	d := m.CloudResourceContainer.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCloudResourceContainer sets the value of the CloudResourceContainer field in Workspace.
func (m *Workspace) SetCloudResourceContainer(ctx context.Context, v CloudResourceContainer) {
	vs := v.ToObjectValue(ctx)
	m.CloudResourceContainer = vs
}

// GetCustomTags returns the value of the CustomTags field in Workspace as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Workspace) GetCustomTags(ctx context.Context) (map[string]types.String, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in Workspace.
func (m *Workspace) SetCustomTags(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.MapValueMust(t, vs)
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in Workspace as
// a GcpManagedNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Workspace) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig, bool) {
	var e GcpManagedNetworkConfig
	if m.GcpManagedNetworkConfig.IsNull() || m.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v GcpManagedNetworkConfig
	d := m.GcpManagedNetworkConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in Workspace.
func (m *Workspace) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig) {
	vs := v.ToObjectValue(ctx)
	m.GcpManagedNetworkConfig = vs
}

// GetGkeConfig returns the value of the GkeConfig field in Workspace as
// a GkeConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Workspace) GetGkeConfig(ctx context.Context) (GkeConfig, bool) {
	var e GkeConfig
	if m.GkeConfig.IsNull() || m.GkeConfig.IsUnknown() {
		return e, false
	}
	var v GkeConfig
	d := m.GkeConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGkeConfig sets the value of the GkeConfig field in Workspace.
func (m *Workspace) SetGkeConfig(ctx context.Context, v GkeConfig) {
	vs := v.ToObjectValue(ctx)
	m.GkeConfig = vs
}

// GetNetwork returns the value of the Network field in Workspace as
// a WorkspaceNetwork value.
// If the field is unknown or null, the boolean return value is false.
func (m *Workspace) GetNetwork(ctx context.Context) (WorkspaceNetwork, bool) {
	var e WorkspaceNetwork
	if m.Network.IsNull() || m.Network.IsUnknown() {
		return e, false
	}
	var v WorkspaceNetwork
	d := m.Network.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNetwork sets the value of the Network field in Workspace.
func (m *Workspace) SetNetwork(ctx context.Context, v WorkspaceNetwork) {
	vs := v.ToObjectValue(ctx)
	m.Network = vs
}

// The network configuration for workspaces.
type WorkspaceNetwork struct {
	// The shared network config for GCP workspace. This object has common
	// network configurations that are network attributions of a workspace. This
	// object is input-only.
	GcpCommonNetworkConfig types.Object `tfsdk:"gcp_common_network_config"`
	// The mutually exclusive network deployment modes. The option decides which
	// network mode the workspace will use. The network config for GCP workspace
	// with Databricks managed network. This object is input-only and will not
	// be provided when listing workspaces. See go/gcp-byovpc-alpha-design for
	// interface decisions.
	GcpManagedNetworkConfig types.Object `tfsdk:"gcp_managed_network_config"`
	// The ID of the network object, if the workspace is a BYOVPC workspace.
	// This should apply to workspaces on all clouds in internal services. In
	// accounts-rest-api, user will use workspace.network_id for input and
	// output instead. Currently (2021-06-19) the network ID is only used by
	// GCP.
	NetworkId types.String `tfsdk:"network_id"`
}

func (to *WorkspaceNetwork) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceNetwork) {
	if !from.GcpCommonNetworkConfig.IsNull() && !from.GcpCommonNetworkConfig.IsUnknown() {
		if toGcpCommonNetworkConfig, ok := to.GetGcpCommonNetworkConfig(ctx); ok {
			if fromGcpCommonNetworkConfig, ok := from.GetGcpCommonNetworkConfig(ctx); ok {
				// Recursively sync the fields of GcpCommonNetworkConfig
				toGcpCommonNetworkConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpCommonNetworkConfig)
				to.SetGcpCommonNetworkConfig(ctx, toGcpCommonNetworkConfig)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				// Recursively sync the fields of GcpManagedNetworkConfig
				toGcpManagedNetworkConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
}

func (to *WorkspaceNetwork) SyncFieldsDuringRead(ctx context.Context, from WorkspaceNetwork) {
	if !from.GcpCommonNetworkConfig.IsNull() && !from.GcpCommonNetworkConfig.IsUnknown() {
		if toGcpCommonNetworkConfig, ok := to.GetGcpCommonNetworkConfig(ctx); ok {
			if fromGcpCommonNetworkConfig, ok := from.GetGcpCommonNetworkConfig(ctx); ok {
				toGcpCommonNetworkConfig.SyncFieldsDuringRead(ctx, fromGcpCommonNetworkConfig)
				to.SetGcpCommonNetworkConfig(ctx, toGcpCommonNetworkConfig)
			}
		}
	}
	if !from.GcpManagedNetworkConfig.IsNull() && !from.GcpManagedNetworkConfig.IsUnknown() {
		if toGcpManagedNetworkConfig, ok := to.GetGcpManagedNetworkConfig(ctx); ok {
			if fromGcpManagedNetworkConfig, ok := from.GetGcpManagedNetworkConfig(ctx); ok {
				toGcpManagedNetworkConfig.SyncFieldsDuringRead(ctx, fromGcpManagedNetworkConfig)
				to.SetGcpManagedNetworkConfig(ctx, toGcpManagedNetworkConfig)
			}
		}
	}
}

func (m WorkspaceNetwork) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["gcp_common_network_config"] = attrs["gcp_common_network_config"].SetOptional()
	attrs["gcp_managed_network_config"] = attrs["gcp_managed_network_config"].SetOptional()
	attrs["network_id"] = attrs["network_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceNetwork.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceNetwork) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"gcp_common_network_config":  reflect.TypeOf(GcpCommonNetworkConfig{}),
		"gcp_managed_network_config": reflect.TypeOf(GcpManagedNetworkConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceNetwork
// only implements ToObjectValue() and Type().
func (m WorkspaceNetwork) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"gcp_common_network_config":  m.GcpCommonNetworkConfig,
			"gcp_managed_network_config": m.GcpManagedNetworkConfig,
			"network_id":                 m.NetworkId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceNetwork) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"gcp_common_network_config":  GcpCommonNetworkConfig{}.Type(ctx),
			"gcp_managed_network_config": GcpManagedNetworkConfig{}.Type(ctx),
			"network_id":                 types.StringType,
		},
	}
}

// GetGcpCommonNetworkConfig returns the value of the GcpCommonNetworkConfig field in WorkspaceNetwork as
// a GcpCommonNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceNetwork) GetGcpCommonNetworkConfig(ctx context.Context) (GcpCommonNetworkConfig, bool) {
	var e GcpCommonNetworkConfig
	if m.GcpCommonNetworkConfig.IsNull() || m.GcpCommonNetworkConfig.IsUnknown() {
		return e, false
	}
	var v GcpCommonNetworkConfig
	d := m.GcpCommonNetworkConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpCommonNetworkConfig sets the value of the GcpCommonNetworkConfig field in WorkspaceNetwork.
func (m *WorkspaceNetwork) SetGcpCommonNetworkConfig(ctx context.Context, v GcpCommonNetworkConfig) {
	vs := v.ToObjectValue(ctx)
	m.GcpCommonNetworkConfig = vs
}

// GetGcpManagedNetworkConfig returns the value of the GcpManagedNetworkConfig field in WorkspaceNetwork as
// a GcpManagedNetworkConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceNetwork) GetGcpManagedNetworkConfig(ctx context.Context) (GcpManagedNetworkConfig, bool) {
	var e GcpManagedNetworkConfig
	if m.GcpManagedNetworkConfig.IsNull() || m.GcpManagedNetworkConfig.IsUnknown() {
		return e, false
	}
	var v GcpManagedNetworkConfig
	d := m.GcpManagedNetworkConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGcpManagedNetworkConfig sets the value of the GcpManagedNetworkConfig field in WorkspaceNetwork.
func (m *WorkspaceNetwork) SetGcpManagedNetworkConfig(ctx context.Context, v GcpManagedNetworkConfig) {
	vs := v.ToObjectValue(ctx)
	m.GcpManagedNetworkConfig = vs
}
