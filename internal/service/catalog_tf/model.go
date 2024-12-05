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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AccountsCreateMetastore struct {
	MetastoreInfo types.List `tfsdk:"metastore_info" tf:"optional,object"`
}

func (newState *AccountsCreateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateMetastore) {
}

func (newState *AccountsCreateMetastore) SyncEffectiveFieldsDuringRead(existingState AccountsCreateMetastore) {
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

// ToAttrType returns the representation of AccountsCreateMetastore in the Terraform plugin framework type
// system.
func (a AccountsCreateMetastore) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_info": basetypes.ListType{
				ElemType: CreateMetastore{}.ToAttrType(ctx),
			},
		},
	}
}

type AccountsCreateMetastoreAssignment struct {
	MetastoreAssignment types.List `tfsdk:"metastore_assignment" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *AccountsCreateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateMetastoreAssignment) {
}

func (newState *AccountsCreateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState AccountsCreateMetastoreAssignment) {
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

// ToAttrType returns the representation of AccountsCreateMetastoreAssignment in the Terraform plugin framework type
// system.
func (a AccountsCreateMetastoreAssignment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_assignment": basetypes.ListType{
				ElemType: CreateMetastoreAssignment{}.ToAttrType(ctx),
			},
			"metastore_id": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

type AccountsCreateStorageCredential struct {
	CredentialInfo types.List `tfsdk:"credential_info" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *AccountsCreateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsCreateStorageCredential) {
}

func (newState *AccountsCreateStorageCredential) SyncEffectiveFieldsDuringRead(existingState AccountsCreateStorageCredential) {
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

// ToAttrType returns the representation of AccountsCreateStorageCredential in the Terraform plugin framework type
// system.
func (a AccountsCreateStorageCredential) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_info": basetypes.ListType{
				ElemType: CreateStorageCredential{}.ToAttrType(ctx),
			},
			"metastore_id": types.StringType,
		},
	}
}

type AccountsMetastoreAssignment struct {
	MetastoreAssignment types.List `tfsdk:"metastore_assignment" tf:"optional,object"`
}

func (newState *AccountsMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsMetastoreAssignment) {
}

func (newState *AccountsMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState AccountsMetastoreAssignment) {
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

// ToAttrType returns the representation of AccountsMetastoreAssignment in the Terraform plugin framework type
// system.
func (a AccountsMetastoreAssignment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_assignment": basetypes.ListType{
				ElemType: MetastoreAssignment{}.ToAttrType(ctx),
			},
		},
	}
}

type AccountsMetastoreInfo struct {
	MetastoreInfo types.List `tfsdk:"metastore_info" tf:"optional,object"`
}

func (newState *AccountsMetastoreInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsMetastoreInfo) {
}

func (newState *AccountsMetastoreInfo) SyncEffectiveFieldsDuringRead(existingState AccountsMetastoreInfo) {
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

// ToAttrType returns the representation of AccountsMetastoreInfo in the Terraform plugin framework type
// system.
func (a AccountsMetastoreInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_info": basetypes.ListType{
				ElemType: MetastoreInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type AccountsStorageCredentialInfo struct {
	CredentialInfo types.List `tfsdk:"credential_info" tf:"optional,object"`
}

func (newState *AccountsStorageCredentialInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsStorageCredentialInfo) {
}

func (newState *AccountsStorageCredentialInfo) SyncEffectiveFieldsDuringRead(existingState AccountsStorageCredentialInfo) {
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

// ToAttrType returns the representation of AccountsStorageCredentialInfo in the Terraform plugin framework type
// system.
func (a AccountsStorageCredentialInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_info": basetypes.ListType{
				ElemType: StorageCredentialInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type AccountsUpdateMetastore struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`

	MetastoreInfo types.List `tfsdk:"metastore_info" tf:"optional,object"`
}

func (newState *AccountsUpdateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateMetastore) {
}

func (newState *AccountsUpdateMetastore) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateMetastore) {
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

// ToAttrType returns the representation of AccountsUpdateMetastore in the Terraform plugin framework type
// system.
func (a AccountsUpdateMetastore) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
			"metastore_info": basetypes.ListType{
				ElemType: UpdateMetastore{}.ToAttrType(ctx),
			},
		},
	}
}

type AccountsUpdateMetastoreAssignment struct {
	MetastoreAssignment types.List `tfsdk:"metastore_assignment" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *AccountsUpdateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateMetastoreAssignment) {
}

func (newState *AccountsUpdateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateMetastoreAssignment) {
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

// ToAttrType returns the representation of AccountsUpdateMetastoreAssignment in the Terraform plugin framework type
// system.
func (a AccountsUpdateMetastoreAssignment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_assignment": basetypes.ListType{
				ElemType: UpdateMetastoreAssignment{}.ToAttrType(ctx),
			},
			"metastore_id": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

type AccountsUpdateStorageCredential struct {
	CredentialInfo types.List `tfsdk:"credential_info" tf:"optional,object"`
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
	// Name of the storage credential.
	StorageCredentialName types.String `tfsdk:"-"`
}

func (newState *AccountsUpdateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountsUpdateStorageCredential) {
}

func (newState *AccountsUpdateStorageCredential) SyncEffectiveFieldsDuringRead(existingState AccountsUpdateStorageCredential) {
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

// ToAttrType returns the representation of AccountsUpdateStorageCredential in the Terraform plugin framework type
// system.
func (a AccountsUpdateStorageCredential) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_info": basetypes.ListType{
				ElemType: UpdateStorageCredential{}.ToAttrType(ctx),
			},
			"metastore_id":            types.StringType,
			"storage_credential_name": types.StringType,
		},
	}
}

type ArtifactAllowlistInfo struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers types.List `tfsdk:"artifact_matchers" tf:"optional"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of the user who set the artifact allowlist.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Unique identifier of parent metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
}

func (newState *ArtifactAllowlistInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ArtifactAllowlistInfo) {
}

func (newState *ArtifactAllowlistInfo) SyncEffectiveFieldsDuringRead(existingState ArtifactAllowlistInfo) {
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

// ToAttrType returns the representation of ArtifactAllowlistInfo in the Terraform plugin framework type
// system.
func (a ArtifactAllowlistInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_matchers": basetypes.ListType{
				ElemType: ArtifactMatcher{}.ToAttrType(ctx),
			},
			"created_at":   types.Int64Type,
			"created_by":   types.StringType,
			"metastore_id": types.StringType,
		},
	}
}

type ArtifactMatcher struct {
	// The artifact path or maven coordinate
	Artifact types.String `tfsdk:"artifact" tf:""`
	// The pattern matching type of the artifact
	MatchType types.String `tfsdk:"match_type" tf:""`
}

func (newState *ArtifactMatcher) SyncEffectiveFieldsDuringCreateOrUpdate(plan ArtifactMatcher) {
}

func (newState *ArtifactMatcher) SyncEffectiveFieldsDuringRead(existingState ArtifactMatcher) {
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

// ToAttrType returns the representation of ArtifactMatcher in the Terraform plugin framework type
// system.
func (a ArtifactMatcher) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact":   types.StringType,
			"match_type": types.StringType,
		},
	}
}

type AssignResponse struct {
}

func (newState *AssignResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AssignResponse) {
}

func (newState *AssignResponse) SyncEffectiveFieldsDuringRead(existingState AssignResponse) {
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

// ToAttrType returns the representation of AssignResponse in the Terraform plugin framework type
// system.
func (a AssignResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// AWS temporary credentials for API authentication. Read more at
// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
type AwsCredentials struct {
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
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of AwsCredentials in the Terraform plugin framework type
// system.
func (a AwsCredentials) ToAttrType(ctx context.Context) types.ObjectType {
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
	ExternalId types.String `tfsdk:"external_id" tf:"optional"`
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	RoleArn types.String `tfsdk:"role_arn" tf:"optional"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn types.String `tfsdk:"unity_catalog_iam_arn" tf:"optional"`
}

func (newState *AwsIamRole) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRole) {
}

func (newState *AwsIamRole) SyncEffectiveFieldsDuringRead(existingState AwsIamRole) {
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

// ToAttrType returns the representation of AwsIamRole in the Terraform plugin framework type
// system.
func (a AwsIamRole) ToAttrType(ctx context.Context) types.ObjectType {
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
	RoleArn types.String `tfsdk:"role_arn" tf:""`
}

func (newState *AwsIamRoleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRoleRequest) {
}

func (newState *AwsIamRoleRequest) SyncEffectiveFieldsDuringRead(existingState AwsIamRoleRequest) {
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

// ToAttrType returns the representation of AwsIamRoleRequest in the Terraform plugin framework type
// system.
func (a AwsIamRoleRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"role_arn": types.StringType,
		},
	}
}

type AwsIamRoleResponse struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	ExternalId types.String `tfsdk:"external_id" tf:"optional"`
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn types.String `tfsdk:"role_arn" tf:""`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn types.String `tfsdk:"unity_catalog_iam_arn" tf:"optional"`
}

func (newState *AwsIamRoleResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AwsIamRoleResponse) {
}

func (newState *AwsIamRoleResponse) SyncEffectiveFieldsDuringRead(existingState AwsIamRoleResponse) {
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

// ToAttrType returns the representation of AwsIamRoleResponse in the Terraform plugin framework type
// system.
func (a AwsIamRoleResponse) ToAttrType(ctx context.Context) types.ObjectType {
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
type AzureActiveDirectoryToken struct {
	// Opaque token that contains claims that you can use in Azure Active
	// Directory to access cloud services.
	AadToken types.String `tfsdk:"aad_token" tf:"optional"`
}

func (newState *AzureActiveDirectoryToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureActiveDirectoryToken) {
}

func (newState *AzureActiveDirectoryToken) SyncEffectiveFieldsDuringRead(existingState AzureActiveDirectoryToken) {
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

// ToAttrType returns the representation of AzureActiveDirectoryToken in the Terraform plugin framework type
// system.
func (a AzureActiveDirectoryToken) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *AzureManagedIdentity) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentity) {
}

func (newState *AzureManagedIdentity) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentity) {
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

// ToAttrType returns the representation of AzureManagedIdentity in the Terraform plugin framework type
// system.
func (a AzureManagedIdentity) ToAttrType(ctx context.Context) types.ObjectType {
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
	AccessConnectorId types.String `tfsdk:"access_connector_id" tf:""`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId types.String `tfsdk:"managed_identity_id" tf:"optional"`
}

func (newState *AzureManagedIdentityRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentityRequest) {
}

func (newState *AzureManagedIdentityRequest) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentityRequest) {
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

// ToAttrType returns the representation of AzureManagedIdentityRequest in the Terraform plugin framework type
// system.
func (a AzureManagedIdentityRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *AzureManagedIdentityResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureManagedIdentityResponse) {
}

func (newState *AzureManagedIdentityResponse) SyncEffectiveFieldsDuringRead(existingState AzureManagedIdentityResponse) {
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

// ToAttrType returns the representation of AzureManagedIdentityResponse in the Terraform plugin framework type
// system.
func (a AzureManagedIdentityResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_connector_id": types.StringType,
			"credential_id":       types.StringType,
			"managed_identity_id": types.StringType,
		},
	}
}

// The Azure service principal configuration.
type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	ApplicationId types.String `tfsdk:"application_id" tf:""`
	// The client secret generated for the above app ID in AAD.
	ClientSecret types.String `tfsdk:"client_secret" tf:""`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	DirectoryId types.String `tfsdk:"directory_id" tf:""`
}

func (newState *AzureServicePrincipal) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureServicePrincipal) {
}

func (newState *AzureServicePrincipal) SyncEffectiveFieldsDuringRead(existingState AzureServicePrincipal) {
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

// ToAttrType returns the representation of AzureServicePrincipal in the Terraform plugin framework type
// system.
func (a AzureServicePrincipal) ToAttrType(ctx context.Context) types.ObjectType {
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
	SasToken types.String `tfsdk:"sas_token" tf:"optional"`
}

func (newState *AzureUserDelegationSas) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureUserDelegationSas) {
}

func (newState *AzureUserDelegationSas) SyncEffectiveFieldsDuringRead(existingState AzureUserDelegationSas) {
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

// ToAttrType returns the representation of AzureUserDelegationSas in the Terraform plugin framework type
// system.
func (a AzureUserDelegationSas) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *CancelRefreshRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelRefreshRequest) {
}

func (newState *CancelRefreshRequest) SyncEffectiveFieldsDuringRead(existingState CancelRefreshRequest) {
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

// ToAttrType returns the representation of CancelRefreshRequest in the Terraform plugin framework type
// system.
func (a CancelRefreshRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"refresh_id": types.StringType,
			"table_name": types.StringType,
		},
	}
}

type CancelRefreshResponse struct {
}

func (newState *CancelRefreshResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelRefreshResponse) {
}

func (newState *CancelRefreshResponse) SyncEffectiveFieldsDuringRead(existingState CancelRefreshResponse) {
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

// ToAttrType returns the representation of CancelRefreshResponse in the Terraform plugin framework type
// system.
func (a CancelRefreshResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CatalogInfo struct {
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

func (newState *CatalogInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CatalogInfo) {
}

func (newState *CatalogInfo) SyncEffectiveFieldsDuringRead(existingState CatalogInfo) {
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

// ToAttrType returns the representation of CatalogInfo in the Terraform plugin framework type
// system.
func (a CatalogInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":     types.BoolType,
			"catalog_type":    types.StringType,
			"comment":         types.StringType,
			"connection_name": types.StringType,
			"created_at":      types.Int64Type,
			"created_by":      types.StringType,
			"effective_predictive_optimization_flag": basetypes.ListType{
				ElemType: EffectivePredictiveOptimizationFlag{}.ToAttrType(ctx),
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
				ElemType: ProvisioningInfo{}.ToAttrType(ctx),
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

type CloudflareApiToken struct {
	// The Cloudflare access key id of the token.
	AccessKeyId types.String `tfsdk:"access_key_id" tf:""`
	// The account id associated with the API token.
	AccountId types.String `tfsdk:"account_id" tf:""`
	// The secret access token generated for the access key id
	SecretAccessKey types.String `tfsdk:"secret_access_key" tf:""`
}

func (newState *CloudflareApiToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan CloudflareApiToken) {
}

func (newState *CloudflareApiToken) SyncEffectiveFieldsDuringRead(existingState CloudflareApiToken) {
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

// ToAttrType returns the representation of CloudflareApiToken in the Terraform plugin framework type
// system.
func (a CloudflareApiToken) ToAttrType(ctx context.Context) types.ObjectType {
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
	// Name of type (INT, STRUCT, MAP, etc.).
	TypeName types.String `tfsdk:"type_name" tf:"optional"`
	// Digits of precision; required for DecimalTypes.
	TypePrecision types.Int64 `tfsdk:"type_precision" tf:"optional"`
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale types.Int64 `tfsdk:"type_scale" tf:"optional"`
	// Full data type specification as SQL/catalogString text.
	TypeText types.String `tfsdk:"type_text" tf:"optional"`
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnInfo) {
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringRead(existingState ColumnInfo) {
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

// ToAttrType returns the representation of ColumnInfo in the Terraform plugin framework type
// system.
func (a ColumnInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"mask": basetypes.ListType{
				ElemType: ColumnMask{}.ToAttrType(ctx),
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

type ColumnMask struct {
	// The full name of the column mask SQL UDF.
	FunctionName types.String `tfsdk:"function_name" tf:"optional"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames types.List `tfsdk:"using_column_names" tf:"optional"`
}

func (newState *ColumnMask) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnMask) {
}

func (newState *ColumnMask) SyncEffectiveFieldsDuringRead(existingState ColumnMask) {
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

// ToAttrType returns the representation of ColumnMask in the Terraform plugin framework type
// system.
func (a ColumnMask) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_name": types.StringType,
			"using_column_names": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type ConnectionInfo struct {
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

func (newState *ConnectionInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ConnectionInfo) {
}

func (newState *ConnectionInfo) SyncEffectiveFieldsDuringRead(existingState ConnectionInfo) {
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

// ToAttrType returns the representation of ConnectionInfo in the Terraform plugin framework type
// system.
func (a ConnectionInfo) ToAttrType(ctx context.Context) types.ObjectType {
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
				ElemType: ProvisioningInfo{}.ToAttrType(ctx),
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

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
type ContinuousUpdateStatus struct {
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

func (newState *ContinuousUpdateStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ContinuousUpdateStatus) {
}

func (newState *ContinuousUpdateStatus) SyncEffectiveFieldsDuringRead(existingState ContinuousUpdateStatus) {
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

// ToAttrType returns the representation of ContinuousUpdateStatus in the Terraform plugin framework type
// system.
func (a ContinuousUpdateStatus) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": basetypes.ListType{
				ElemType: PipelineProgress{}.ToAttrType(ctx),
			},
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
		},
	}
}

type CreateCatalog struct {
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

func (newState *CreateCatalog) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCatalog) {
}

func (newState *CreateCatalog) SyncEffectiveFieldsDuringRead(existingState CreateCatalog) {
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

// ToAttrType returns the representation of CreateCatalog in the Terraform plugin framework type
// system.
func (a CreateCatalog) ToAttrType(ctx context.Context) types.ObjectType {
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

type CreateConnection struct {
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

func (newState *CreateConnection) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateConnection) {
}

func (newState *CreateConnection) SyncEffectiveFieldsDuringRead(existingState CreateConnection) {
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

// ToAttrType returns the representation of CreateConnection in the Terraform plugin framework type
// system.
func (a CreateConnection) ToAttrType(ctx context.Context) types.ObjectType {
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

type CreateCredentialRequest struct {
	// The AWS IAM role configuration
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// TODO(UC-978): Document GCP service account key usage for service
	// credentials.
	GcpServiceAccountKey types.List `tfsdk:"gcp_service_account_key" tf:"optional,object"`
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
		"aws_iam_role":            reflect.TypeOf(AwsIamRole{}),
		"azure_managed_identity":  reflect.TypeOf(AzureManagedIdentity{}),
		"azure_service_principal": reflect.TypeOf(AzureServicePrincipal{}),
		"gcp_service_account_key": reflect.TypeOf(GcpServiceAccountKey{}),
	}
}

// ToAttrType returns the representation of CreateCredentialRequest in the Terraform plugin framework type
// system.
func (a CreateCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRole{}.ToAttrType(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentity{}.ToAttrType(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal{}.ToAttrType(ctx),
			},
			"comment": types.StringType,
			"gcp_service_account_key": basetypes.ListType{
				ElemType: GcpServiceAccountKey{}.ToAttrType(ctx),
			},
			"name":            types.StringType,
			"purpose":         types.StringType,
			"read_only":       types.BoolType,
			"skip_validation": types.BoolType,
		},
	}
}

type CreateExternalLocation struct {
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

func (newState *CreateExternalLocation) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExternalLocation) {
}

func (newState *CreateExternalLocation) SyncEffectiveFieldsDuringRead(existingState CreateExternalLocation) {
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
	}
}

// ToAttrType returns the representation of CreateExternalLocation in the Terraform plugin framework type
// system.
func (a CreateExternalLocation) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point":    types.StringType,
			"comment":         types.StringType,
			"credential_name": types.StringType,
			"encryption_details": basetypes.ListType{
				ElemType: EncryptionDetails{}.ToAttrType(ctx),
			},
			"fallback":        types.BoolType,
			"name":            types.StringType,
			"read_only":       types.BoolType,
			"skip_validation": types.BoolType,
			"url":             types.StringType,
		},
	}
}

type CreateFunction struct {
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

func (newState *CreateFunction) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFunction) {
}

func (newState *CreateFunction) SyncEffectiveFieldsDuringRead(existingState CreateFunction) {
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

// ToAttrType returns the representation of CreateFunction in the Terraform plugin framework type
// system.
func (a CreateFunction) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":      types.StringType,
			"comment":           types.StringType,
			"data_type":         types.StringType,
			"external_language": types.StringType,
			"external_name":     types.StringType,
			"full_data_type":    types.StringType,
			"input_params": basetypes.ListType{
				ElemType: FunctionParameterInfos{}.ToAttrType(ctx),
			},
			"is_deterministic": types.BoolType,
			"is_null_call":     types.BoolType,
			"name":             types.StringType,
			"parameter_style":  types.StringType,
			"properties":       types.StringType,
			"return_params": basetypes.ListType{
				ElemType: FunctionParameterInfos{}.ToAttrType(ctx),
			},
			"routine_body":       types.StringType,
			"routine_definition": types.StringType,
			"routine_dependencies": basetypes.ListType{
				ElemType: DependencyList{}.ToAttrType(ctx),
			},
			"schema_name":     types.StringType,
			"security_type":   types.StringType,
			"specific_name":   types.StringType,
			"sql_data_access": types.StringType,
			"sql_path":        types.StringType,
		},
	}
}

type CreateFunctionRequest struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	FunctionInfo types.List `tfsdk:"function_info" tf:"object"`
}

func (newState *CreateFunctionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateFunctionRequest) {
}

func (newState *CreateFunctionRequest) SyncEffectiveFieldsDuringRead(existingState CreateFunctionRequest) {
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

// ToAttrType returns the representation of CreateFunctionRequest in the Terraform plugin framework type
// system.
func (a CreateFunctionRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_info": basetypes.ListType{
				ElemType: CreateFunction{}.ToAttrType(ctx),
			},
		},
	}
}

type CreateMetastore struct {
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

func (newState *CreateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMetastore) {
}

func (newState *CreateMetastore) SyncEffectiveFieldsDuringRead(existingState CreateMetastore) {
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

// ToAttrType returns the representation of CreateMetastore in the Terraform plugin framework type
// system.
func (a CreateMetastore) ToAttrType(ctx context.Context) types.ObjectType {
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
	DefaultCatalogName types.String `tfsdk:"default_catalog_name" tf:""`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:""`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *CreateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMetastoreAssignment) {
}

func (newState *CreateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState CreateMetastoreAssignment) {
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

// ToAttrType returns the representation of CreateMetastoreAssignment in the Terraform plugin framework type
// system.
func (a CreateMetastoreAssignment) ToAttrType(ctx context.Context) types.ObjectType {
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
	Snapshot []MonitorSnapshot `tfsdk:"snapshot" tf:"optional,object"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries types.List `tfsdk:"time_series" tf:"optional,object"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *CreateMonitor) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateMonitor) {
}

func (newState *CreateMonitor) SyncEffectiveFieldsDuringRead(existingState CreateMonitor) {
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

// ToAttrType returns the representation of CreateMonitor in the Terraform plugin framework type
// system.
func (a CreateMonitor) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets_dir":          types.StringType,
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: MonitorMetric{}.ToAttrType(ctx),
			},
			"data_classification_config": basetypes.ListType{
				ElemType: MonitorDataClassificationConfig{}.ToAttrType(ctx),
			},
			"inference_log": basetypes.ListType{
				ElemType: MonitorInferenceLog{}.ToAttrType(ctx),
			},
			"notifications": basetypes.ListType{
				ElemType: MonitorNotifications{}.ToAttrType(ctx),
			},
			"output_schema_name": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: MonitorCronSchedule{}.ToAttrType(ctx),
			},
			"skip_builtin_dashboard": types.BoolType,
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot": basetypes.ListType{
				ElemType: MonitorSnapshot{}.ToAttrType(ctx),
			},
			"table_name": types.StringType,
			"time_series": basetypes.ListType{
				ElemType: MonitorTimeSeries{}.ToAttrType(ctx),
			},
			"warehouse_id": types.StringType,
		},
	}
}

// Create an Online Table
type CreateOnlineTableRequest struct {
	// Online Table information.
	Table types.List `tfsdk:"table" tf:"optional,object"`
}

func (newState *CreateOnlineTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateOnlineTableRequest) {
}

func (newState *CreateOnlineTableRequest) SyncEffectiveFieldsDuringRead(existingState CreateOnlineTableRequest) {
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

// ToAttrType returns the representation of CreateOnlineTableRequest in the Terraform plugin framework type
// system.
func (a CreateOnlineTableRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table": basetypes.ListType{
				ElemType: OnlineTable{}.ToAttrType(ctx),
			},
		},
	}
}

type CreateRegisteredModelRequest struct {
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

func (newState *CreateRegisteredModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRegisteredModelRequest) {
}

func (newState *CreateRegisteredModelRequest) SyncEffectiveFieldsDuringRead(existingState CreateRegisteredModelRequest) {
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

// ToAttrType returns the representation of CreateRegisteredModelRequest in the Terraform plugin framework type
// system.
func (a CreateRegisteredModelRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *CreateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateResponse) {
}

func (newState *CreateResponse) SyncEffectiveFieldsDuringRead(existingState CreateResponse) {
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

// ToAttrType returns the representation of CreateResponse in the Terraform plugin framework type
// system.
func (a CreateResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CreateSchema struct {
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

func (newState *CreateSchema) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateSchema) {
}

func (newState *CreateSchema) SyncEffectiveFieldsDuringRead(existingState CreateSchema) {
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

// ToAttrType returns the representation of CreateSchema in the Terraform plugin framework type
// system.
func (a CreateSchema) ToAttrType(ctx context.Context) types.ObjectType {
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

type CreateStorageCredential struct {
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
	DatabricksGcpServiceAccount []DatabricksGcpServiceAccountRequest `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// The credential name. The name must be unique within the metastore.
	Name types.String `tfsdk:"name" tf:""`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// Supplying true to this argument skips validation of the created
	// credential.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
}

func (newState *CreateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateStorageCredential) {
}

func (newState *CreateStorageCredential) SyncEffectiveFieldsDuringRead(existingState CreateStorageCredential) {
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

// ToAttrType returns the representation of CreateStorageCredential in the Terraform plugin framework type
// system.
func (a CreateStorageCredential) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRoleRequest{}.ToAttrType(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentityRequest{}.ToAttrType(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal{}.ToAttrType(ctx),
			},
			"cloudflare_api_token": basetypes.ListType{
				ElemType: CloudflareApiToken{}.ToAttrType(ctx),
			},
			"comment": types.StringType,
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccountRequest{}.ToAttrType(ctx),
			},
			"name":            types.StringType,
			"read_only":       types.BoolType,
			"skip_validation": types.BoolType,
		},
	}
}

type CreateTableConstraint struct {
	// A table constraint, as defined by *one* of the following fields being
	// set: __primary_key_constraint__, __foreign_key_constraint__,
	// __named_table_constraint__.
	Constraint types.List `tfsdk:"constraint" tf:"object"`
	// The full name of the table referenced by the constraint.
	FullNameArg types.String `tfsdk:"full_name_arg" tf:""`
}

func (newState *CreateTableConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTableConstraint) {
}

func (newState *CreateTableConstraint) SyncEffectiveFieldsDuringRead(existingState CreateTableConstraint) {
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

// ToAttrType returns the representation of CreateTableConstraint in the Terraform plugin framework type
// system.
func (a CreateTableConstraint) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"constraint": basetypes.ListType{
				ElemType: TableConstraint{}.ToAttrType(ctx),
			},
			"full_name_arg": types.StringType,
		},
	}
}

type CreateVolumeRequestContent struct {
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

func (newState *CreateVolumeRequestContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVolumeRequestContent) {
}

func (newState *CreateVolumeRequestContent) SyncEffectiveFieldsDuringRead(existingState CreateVolumeRequestContent) {
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

// ToAttrType returns the representation of CreateVolumeRequestContent in the Terraform plugin framework type
// system.
func (a CreateVolumeRequestContent) ToAttrType(ctx context.Context) types.ObjectType {
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
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Time at which this credential was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// Username of credential creator.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
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

func (newState *CredentialInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CredentialInfo) {
}

func (newState *CredentialInfo) SyncEffectiveFieldsDuringRead(existingState CredentialInfo) {
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
		"aws_iam_role":            reflect.TypeOf(AwsIamRole{}),
		"azure_managed_identity":  reflect.TypeOf(AzureManagedIdentity{}),
		"azure_service_principal": reflect.TypeOf(AzureServicePrincipal{}),
	}
}

// ToAttrType returns the representation of CredentialInfo in the Terraform plugin framework type
// system.
func (a CredentialInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRole{}.ToAttrType(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentity{}.ToAttrType(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal{}.ToAttrType(ctx),
			},
			"comment":                  types.StringType,
			"created_at":               types.Int64Type,
			"created_by":               types.StringType,
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

type CredentialValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message types.String `tfsdk:"message" tf:"optional"`
	// The results of the tested operation.
	Result types.String `tfsdk:"result" tf:"optional"`
}

func (newState *CredentialValidationResult) SyncEffectiveFieldsDuringCreateOrUpdate(plan CredentialValidationResult) {
}

func (newState *CredentialValidationResult) SyncEffectiveFieldsDuringRead(existingState CredentialValidationResult) {
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

// ToAttrType returns the representation of CredentialValidationResult in the Terraform plugin framework type
// system.
func (a CredentialValidationResult) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"result":  types.StringType,
		},
	}
}

// Currently assigned workspaces
type CurrentWorkspaceBindings struct {
	// A list of workspace IDs.
	Workspaces types.List `tfsdk:"workspaces" tf:"optional"`
}

func (newState *CurrentWorkspaceBindings) SyncEffectiveFieldsDuringCreateOrUpdate(plan CurrentWorkspaceBindings) {
}

func (newState *CurrentWorkspaceBindings) SyncEffectiveFieldsDuringRead(existingState CurrentWorkspaceBindings) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CurrentWorkspaceBindings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CurrentWorkspaceBindings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspaces": reflect.TypeOf(types.Int64{}),
	}
}

// ToAttrType returns the representation of CurrentWorkspaceBindings in the Terraform plugin framework type
// system.
func (a CurrentWorkspaceBindings) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspaces": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

type DatabricksGcpServiceAccountRequest struct {
}

func (newState *DatabricksGcpServiceAccountRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksGcpServiceAccountRequest) {
}

func (newState *DatabricksGcpServiceAccountRequest) SyncEffectiveFieldsDuringRead(existingState DatabricksGcpServiceAccountRequest) {
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

// ToAttrType returns the representation of DatabricksGcpServiceAccountRequest in the Terraform plugin framework type
// system.
func (a DatabricksGcpServiceAccountRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DatabricksGcpServiceAccountResponse struct {
	// The Databricks internal ID that represents this service account. This is
	// an output-only field.
	CredentialId types.String `tfsdk:"credential_id" tf:"optional"`
	// The email of the service account. This is an output-only field.
	Email types.String `tfsdk:"email" tf:"optional"`
}

func (newState *DatabricksGcpServiceAccountResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksGcpServiceAccountResponse) {
}

func (newState *DatabricksGcpServiceAccountResponse) SyncEffectiveFieldsDuringRead(existingState DatabricksGcpServiceAccountResponse) {
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

// ToAttrType returns the representation of DatabricksGcpServiceAccountResponse in the Terraform plugin framework type
// system.
func (a DatabricksGcpServiceAccountResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteAccountMetastoreAssignmentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountMetastoreAssignmentRequest) {
}

func (newState *DeleteAccountMetastoreAssignmentRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountMetastoreAssignmentRequest) {
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

// ToAttrType returns the representation of DeleteAccountMetastoreAssignmentRequest in the Terraform plugin framework type
// system.
func (a DeleteAccountMetastoreAssignmentRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteAccountMetastoreRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountMetastoreRequest) {
}

func (newState *DeleteAccountMetastoreRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountMetastoreRequest) {
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

// ToAttrType returns the representation of DeleteAccountMetastoreRequest in the Terraform plugin framework type
// system.
func (a DeleteAccountMetastoreRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteAccountStorageCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountStorageCredentialRequest) {
}

func (newState *DeleteAccountStorageCredentialRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountStorageCredentialRequest) {
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

// ToAttrType returns the representation of DeleteAccountStorageCredentialRequest in the Terraform plugin framework type
// system.
func (a DeleteAccountStorageCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteAliasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAliasRequest) {
}

func (newState *DeleteAliasRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAliasRequest) {
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

// ToAttrType returns the representation of DeleteAliasRequest in the Terraform plugin framework type
// system.
func (a DeleteAliasRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias":     types.StringType,
			"full_name": types.StringType,
		},
	}
}

type DeleteAliasResponse struct {
}

func (newState *DeleteAliasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAliasResponse) {
}

func (newState *DeleteAliasResponse) SyncEffectiveFieldsDuringRead(existingState DeleteAliasResponse) {
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

// ToAttrType returns the representation of DeleteAliasResponse in the Terraform plugin framework type
// system.
func (a DeleteAliasResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteCatalogRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCatalogRequest) {
}

func (newState *DeleteCatalogRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCatalogRequest) {
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

// ToAttrType returns the representation of DeleteCatalogRequest in the Terraform plugin framework type
// system.
func (a DeleteCatalogRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteConnectionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteConnectionRequest) {
}

func (newState *DeleteConnectionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteConnectionRequest) {
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

// ToAttrType returns the representation of DeleteConnectionRequest in the Terraform plugin framework type
// system.
func (a DeleteConnectionRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

// ToAttrType returns the representation of DeleteCredentialRequest in the Terraform plugin framework type
// system.
func (a DeleteCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

// ToAttrType returns the representation of DeleteCredentialResponse in the Terraform plugin framework type
// system.
func (a DeleteCredentialResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteExternalLocationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExternalLocationRequest) {
}

func (newState *DeleteExternalLocationRequest) SyncEffectiveFieldsDuringRead(existingState DeleteExternalLocationRequest) {
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

// ToAttrType returns the representation of DeleteExternalLocationRequest in the Terraform plugin framework type
// system.
func (a DeleteExternalLocationRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteFunctionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteFunctionRequest) {
}

func (newState *DeleteFunctionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteFunctionRequest) {
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

// ToAttrType returns the representation of DeleteFunctionRequest in the Terraform plugin framework type
// system.
func (a DeleteFunctionRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteMetastoreRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteMetastoreRequest) {
}

func (newState *DeleteMetastoreRequest) SyncEffectiveFieldsDuringRead(existingState DeleteMetastoreRequest) {
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

// ToAttrType returns the representation of DeleteMetastoreRequest in the Terraform plugin framework type
// system.
func (a DeleteMetastoreRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionRequest) {
}

func (newState *DeleteModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionRequest) {
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

// ToAttrType returns the representation of DeleteModelVersionRequest in the Terraform plugin framework type
// system.
func (a DeleteModelVersionRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteOnlineTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteOnlineTableRequest) {
}

func (newState *DeleteOnlineTableRequest) SyncEffectiveFieldsDuringRead(existingState DeleteOnlineTableRequest) {
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

// ToAttrType returns the representation of DeleteOnlineTableRequest in the Terraform plugin framework type
// system.
func (a DeleteOnlineTableRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteQualityMonitorRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteQualityMonitorRequest) {
}

func (newState *DeleteQualityMonitorRequest) SyncEffectiveFieldsDuringRead(existingState DeleteQualityMonitorRequest) {
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

// ToAttrType returns the representation of DeleteQualityMonitorRequest in the Terraform plugin framework type
// system.
func (a DeleteQualityMonitorRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteRegisteredModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRegisteredModelRequest) {
}

func (newState *DeleteRegisteredModelRequest) SyncEffectiveFieldsDuringRead(existingState DeleteRegisteredModelRequest) {
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

// ToAttrType returns the representation of DeleteRegisteredModelRequest in the Terraform plugin framework type
// system.
func (a DeleteRegisteredModelRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

// ToAttrType returns the representation of DeleteResponse in the Terraform plugin framework type
// system.
func (a DeleteResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteSchemaRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSchemaRequest) {
}

func (newState *DeleteSchemaRequest) SyncEffectiveFieldsDuringRead(existingState DeleteSchemaRequest) {
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

// ToAttrType returns the representation of DeleteSchemaRequest in the Terraform plugin framework type
// system.
func (a DeleteSchemaRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteStorageCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteStorageCredentialRequest) {
}

func (newState *DeleteStorageCredentialRequest) SyncEffectiveFieldsDuringRead(existingState DeleteStorageCredentialRequest) {
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

// ToAttrType returns the representation of DeleteStorageCredentialRequest in the Terraform plugin framework type
// system.
func (a DeleteStorageCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
		},
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

func (newState *DeleteTableConstraintRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTableConstraintRequest) {
}

func (newState *DeleteTableConstraintRequest) SyncEffectiveFieldsDuringRead(existingState DeleteTableConstraintRequest) {
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

// ToAttrType returns the representation of DeleteTableConstraintRequest in the Terraform plugin framework type
// system.
func (a DeleteTableConstraintRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTableRequest) {
}

func (newState *DeleteTableRequest) SyncEffectiveFieldsDuringRead(existingState DeleteTableRequest) {
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

// ToAttrType returns the representation of DeleteTableRequest in the Terraform plugin framework type
// system.
func (a DeleteTableRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *DeleteVolumeRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteVolumeRequest) {
}

func (newState *DeleteVolumeRequest) SyncEffectiveFieldsDuringRead(existingState DeleteVolumeRequest) {
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

// ToAttrType returns the representation of DeleteVolumeRequest in the Terraform plugin framework type
// system.
func (a DeleteVolumeRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	DeltaRuntimeProperties types.Map `tfsdk:"delta_runtime_properties" tf:""`
}

func (newState *DeltaRuntimePropertiesKvPairs) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaRuntimePropertiesKvPairs) {
}

func (newState *DeltaRuntimePropertiesKvPairs) SyncEffectiveFieldsDuringRead(existingState DeltaRuntimePropertiesKvPairs) {
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

// ToAttrType returns the representation of DeltaRuntimePropertiesKvPairs in the Terraform plugin framework type
// system.
func (a DeltaRuntimePropertiesKvPairs) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_runtime_properties": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// A dependency of a SQL object. Either the __table__ field or the __function__
// field must be defined.
type Dependency struct {
	// A function that is dependent on a SQL object.
	Function types.List `tfsdk:"function" tf:"optional,object"`
	// A table that is dependent on a SQL object.
	Table types.List `tfsdk:"table" tf:"optional,object"`
}

func (newState *Dependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dependency) {
}

func (newState *Dependency) SyncEffectiveFieldsDuringRead(existingState Dependency) {
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

// ToAttrType returns the representation of Dependency in the Terraform plugin framework type
// system.
func (a Dependency) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function": basetypes.ListType{
				ElemType: FunctionDependency{}.ToAttrType(ctx),
			},
			"table": basetypes.ListType{
				ElemType: TableDependency{}.ToAttrType(ctx),
			},
		},
	}
}

// A list of dependencies.
type DependencyList struct {
	// Array of dependencies.
	Dependencies types.List `tfsdk:"dependencies" tf:"optional"`
}

func (newState *DependencyList) SyncEffectiveFieldsDuringCreateOrUpdate(plan DependencyList) {
}

func (newState *DependencyList) SyncEffectiveFieldsDuringRead(existingState DependencyList) {
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

// ToAttrType returns the representation of DependencyList in the Terraform plugin framework type
// system.
func (a DependencyList) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dependencies": basetypes.ListType{
				ElemType: Dependency{}.ToAttrType(ctx),
			},
		},
	}
}

// Disable a system schema
type DisableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId types.String `tfsdk:"-"`
	// Full name of the system schema.
	SchemaName types.String `tfsdk:"-"`
}

func (newState *DisableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableRequest) {
}

func (newState *DisableRequest) SyncEffectiveFieldsDuringRead(existingState DisableRequest) {
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

// ToAttrType returns the representation of DisableRequest in the Terraform plugin framework type
// system.
func (a DisableRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

// ToAttrType returns the representation of DisableResponse in the Terraform plugin framework type
// system.
func (a DisableResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EffectivePermissionsList struct {
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments" tf:"optional"`
}

func (newState *EffectivePermissionsList) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePermissionsList) {
}

func (newState *EffectivePermissionsList) SyncEffectiveFieldsDuringRead(existingState EffectivePermissionsList) {
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

// ToAttrType returns the representation of EffectivePermissionsList in the Terraform plugin framework type
// system.
func (a EffectivePermissionsList) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: EffectivePrivilegeAssignment{}.ToAttrType(ctx),
			},
		},
	}
}

type EffectivePredictiveOptimizationFlag struct {
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

func (newState *EffectivePredictiveOptimizationFlag) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePredictiveOptimizationFlag) {
}

func (newState *EffectivePredictiveOptimizationFlag) SyncEffectiveFieldsDuringRead(existingState EffectivePredictiveOptimizationFlag) {
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

// ToAttrType returns the representation of EffectivePredictiveOptimizationFlag in the Terraform plugin framework type
// system.
func (a EffectivePredictiveOptimizationFlag) ToAttrType(ctx context.Context) types.ObjectType {
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
	InheritedFromName types.String `tfsdk:"inherited_from_name" tf:"optional"`
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	InheritedFromType types.String `tfsdk:"inherited_from_type" tf:"optional"`
	// The privilege assigned to the principal.
	Privilege types.String `tfsdk:"privilege" tf:"optional"`
}

func (newState *EffectivePrivilege) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePrivilege) {
}

func (newState *EffectivePrivilege) SyncEffectiveFieldsDuringRead(existingState EffectivePrivilege) {
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

// ToAttrType returns the representation of EffectivePrivilege in the Terraform plugin framework type
// system.
func (a EffectivePrivilege) ToAttrType(ctx context.Context) types.ObjectType {
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
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	Privileges types.List `tfsdk:"privileges" tf:"optional"`
}

func (newState *EffectivePrivilegeAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan EffectivePrivilegeAssignment) {
}

func (newState *EffectivePrivilegeAssignment) SyncEffectiveFieldsDuringRead(existingState EffectivePrivilegeAssignment) {
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

// ToAttrType returns the representation of EffectivePrivilegeAssignment in the Terraform plugin framework type
// system.
func (a EffectivePrivilegeAssignment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"privileges": basetypes.ListType{
				ElemType: EffectivePrivilege{}.ToAttrType(ctx),
			},
		},
	}
}

// Enable a system schema
type EnableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId types.String `tfsdk:"-"`
	// Full name of the system schema.
	SchemaName types.String `tfsdk:"-"`
}

func (newState *EnableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnableRequest) {
}

func (newState *EnableRequest) SyncEffectiveFieldsDuringRead(existingState EnableRequest) {
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

// ToAttrType returns the representation of EnableRequest in the Terraform plugin framework type
// system.
func (a EnableRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
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

// ToAttrType returns the representation of EnableResponse in the Terraform plugin framework type
// system.
func (a EnableResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Encryption options that apply to clients connecting to cloud storage.
type EncryptionDetails struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	SseEncryptionDetails types.List `tfsdk:"sse_encryption_details" tf:"optional,object"`
}

func (newState *EncryptionDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan EncryptionDetails) {
}

func (newState *EncryptionDetails) SyncEffectiveFieldsDuringRead(existingState EncryptionDetails) {
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

// ToAttrType returns the representation of EncryptionDetails in the Terraform plugin framework type
// system.
func (a EncryptionDetails) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"sse_encryption_details": basetypes.ListType{
				ElemType: SseEncryptionDetails{}.ToAttrType(ctx),
			},
		},
	}
}

// Get boolean reflecting if table exists
type ExistsRequest struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`
}

func (newState *ExistsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExistsRequest) {
}

func (newState *ExistsRequest) SyncEffectiveFieldsDuringRead(existingState ExistsRequest) {
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

// ToAttrType returns the representation of ExistsRequest in the Terraform plugin framework type
// system.
func (a ExistsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
		},
	}
}

type ExternalLocationInfo struct {
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

func (newState *ExternalLocationInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalLocationInfo) {
}

func (newState *ExternalLocationInfo) SyncEffectiveFieldsDuringRead(existingState ExternalLocationInfo) {
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
	}
}

// ToAttrType returns the representation of ExternalLocationInfo in the Terraform plugin framework type
// system.
func (a ExternalLocationInfo) ToAttrType(ctx context.Context) types.ObjectType {
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
				ElemType: EncryptionDetails{}.ToAttrType(ctx),
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

// Detailed status of an online table. Shown if the online table is in the
// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
type FailedStatus struct {
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

func (newState *FailedStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan FailedStatus) {
}

func (newState *FailedStatus) SyncEffectiveFieldsDuringRead(existingState FailedStatus) {
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

// ToAttrType returns the representation of FailedStatus in the Terraform plugin framework type
// system.
func (a FailedStatus) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
		},
	}
}

type ForeignKeyConstraint struct {
	// Column names for this constraint.
	ChildColumns types.List `tfsdk:"child_columns" tf:""`
	// The name of the constraint.
	Name types.String `tfsdk:"name" tf:""`
	// Column names for this constraint.
	ParentColumns types.List `tfsdk:"parent_columns" tf:""`
	// The full name of the parent constraint.
	ParentTable types.String `tfsdk:"parent_table" tf:""`
}

func (newState *ForeignKeyConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForeignKeyConstraint) {
}

func (newState *ForeignKeyConstraint) SyncEffectiveFieldsDuringRead(existingState ForeignKeyConstraint) {
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

// ToAttrType returns the representation of ForeignKeyConstraint in the Terraform plugin framework type
// system.
func (a ForeignKeyConstraint) ToAttrType(ctx context.Context) types.ObjectType {
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

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	FunctionFullName types.String `tfsdk:"function_full_name" tf:""`
}

func (newState *FunctionDependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionDependency) {
}

func (newState *FunctionDependency) SyncEffectiveFieldsDuringRead(existingState FunctionDependency) {
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

// ToAttrType returns the representation of FunctionDependency in the Terraform plugin framework type
// system.
func (a FunctionDependency) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *FunctionInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionInfo) {
}

func (newState *FunctionInfo) SyncEffectiveFieldsDuringRead(existingState FunctionInfo) {
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

// ToAttrType returns the representation of FunctionInfo in the Terraform plugin framework type
// system.
func (a FunctionInfo) ToAttrType(ctx context.Context) types.ObjectType {
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
				ElemType: FunctionParameterInfos{}.ToAttrType(ctx),
			},
			"is_deterministic": types.BoolType,
			"is_null_call":     types.BoolType,
			"metastore_id":     types.StringType,
			"name":             types.StringType,
			"owner":            types.StringType,
			"parameter_style":  types.StringType,
			"properties":       types.StringType,
			"return_params": basetypes.ListType{
				ElemType: FunctionParameterInfos{}.ToAttrType(ctx),
			},
			"routine_body":       types.StringType,
			"routine_definition": types.StringType,
			"routine_dependencies": basetypes.ListType{
				ElemType: DependencyList{}.ToAttrType(ctx),
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

type FunctionParameterInfo struct {
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
	// Name of type (INT, STRUCT, MAP, etc.).
	TypeName types.String `tfsdk:"type_name" tf:""`
	// Digits of precision; required on Create for DecimalTypes.
	TypePrecision types.Int64 `tfsdk:"type_precision" tf:"optional"`
	// Digits to right of decimal; Required on Create for DecimalTypes.
	TypeScale types.Int64 `tfsdk:"type_scale" tf:"optional"`
	// Full data type spec, SQL/catalogString text.
	TypeText types.String `tfsdk:"type_text" tf:""`
}

func (newState *FunctionParameterInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionParameterInfo) {
}

func (newState *FunctionParameterInfo) SyncEffectiveFieldsDuringRead(existingState FunctionParameterInfo) {
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

// ToAttrType returns the representation of FunctionParameterInfo in the Terraform plugin framework type
// system.
func (a FunctionParameterInfo) ToAttrType(ctx context.Context) types.ObjectType {
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
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
}

func (newState *FunctionParameterInfos) SyncEffectiveFieldsDuringCreateOrUpdate(plan FunctionParameterInfos) {
}

func (newState *FunctionParameterInfos) SyncEffectiveFieldsDuringRead(existingState FunctionParameterInfos) {
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

// ToAttrType returns the representation of FunctionParameterInfos in the Terraform plugin framework type
// system.
func (a FunctionParameterInfos) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parameters": basetypes.ListType{
				ElemType: FunctionParameterInfo{}.ToAttrType(ctx),
			},
		},
	}
}

// GCP temporary credentials for API authentication. Read more at
// https://developers.google.com/identity/protocols/oauth2/service-account
type GcpOauthToken struct {
	OauthToken types.String `tfsdk:"oauth_token" tf:"optional"`
}

func (newState *GcpOauthToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpOauthToken) {
}

func (newState *GcpOauthToken) SyncEffectiveFieldsDuringRead(existingState GcpOauthToken) {
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

// ToAttrType returns the representation of GcpOauthToken in the Terraform plugin framework type
// system.
func (a GcpOauthToken) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"oauth_token": types.StringType,
		},
	}
}

// GCP long-lived credential. GCP Service Account.
type GcpServiceAccountKey struct {
	// The email of the service account. [Create:REQ Update:OPT].
	Email types.String `tfsdk:"email" tf:"optional"`
	// The service account's RSA private key. [Create:REQ Update:OPT]
	PrivateKey types.String `tfsdk:"private_key" tf:"optional"`
	// The ID of the service account's private key. [Create:REQ Update:OPT]
	PrivateKeyId types.String `tfsdk:"private_key_id" tf:"optional"`
}

func (newState *GcpServiceAccountKey) SyncEffectiveFieldsDuringCreateOrUpdate(plan GcpServiceAccountKey) {
}

func (newState *GcpServiceAccountKey) SyncEffectiveFieldsDuringRead(existingState GcpServiceAccountKey) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GcpServiceAccountKey.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GcpServiceAccountKey) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToAttrType returns the representation of GcpServiceAccountKey in the Terraform plugin framework type
// system.
func (a GcpServiceAccountKey) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email":          types.StringType,
			"private_key":    types.StringType,
			"private_key_id": types.StringType,
		},
	}
}

// Options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialAzureOptions struct {
	// The resources to which the temporary Azure credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://learn.microsoft.com/python/api/azure-core/azure.core.credentials.tokencredential?view=azure-python)
	Resources types.List `tfsdk:"resources" tf:"optional"`
}

func (newState *GenerateTemporaryServiceCredentialAzureOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryServiceCredentialAzureOptions) {
}

func (newState *GenerateTemporaryServiceCredentialAzureOptions) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryServiceCredentialAzureOptions) {
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

// ToAttrType returns the representation of GenerateTemporaryServiceCredentialAzureOptions in the Terraform plugin framework type
// system.
func (a GenerateTemporaryServiceCredentialAzureOptions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resources": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type GenerateTemporaryServiceCredentialRequest struct {
	// Options to customize the requested temporary credential
	AzureOptions types.List `tfsdk:"azure_options" tf:"optional,object"`
	// The name of the service credential used to generate a temporary
	// credential
	CredentialName types.String `tfsdk:"credential_name" tf:""`
}

func (newState *GenerateTemporaryServiceCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryServiceCredentialRequest) {
}

func (newState *GenerateTemporaryServiceCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryServiceCredentialRequest) {
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
	}
}

// ToAttrType returns the representation of GenerateTemporaryServiceCredentialRequest in the Terraform plugin framework type
// system.
func (a GenerateTemporaryServiceCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"azure_options": basetypes.ListType{
				ElemType: GenerateTemporaryServiceCredentialAzureOptions{}.ToAttrType(ctx),
			},
			"credential_name": types.StringType,
		},
	}
}

type GenerateTemporaryTableCredentialRequest struct {
	// The operation performed against the table data, either READ or
	// READ_WRITE. If READ_WRITE is specified, the credentials returned will
	// have write permissions, otherwise, it will be read only.
	Operation types.String `tfsdk:"operation" tf:"optional"`
	// UUID of the table to read or write.
	TableId types.String `tfsdk:"table_id" tf:"optional"`
}

func (newState *GenerateTemporaryTableCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryTableCredentialRequest) {
}

func (newState *GenerateTemporaryTableCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryTableCredentialRequest) {
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

// ToAttrType returns the representation of GenerateTemporaryTableCredentialRequest in the Terraform plugin framework type
// system.
func (a GenerateTemporaryTableCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GenerateTemporaryTableCredentialResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateTemporaryTableCredentialResponse) {
}

func (newState *GenerateTemporaryTableCredentialResponse) SyncEffectiveFieldsDuringRead(existingState GenerateTemporaryTableCredentialResponse) {
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

// ToAttrType returns the representation of GenerateTemporaryTableCredentialResponse in the Terraform plugin framework type
// system.
func (a GenerateTemporaryTableCredentialResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_temp_credentials": basetypes.ListType{
				ElemType: AwsCredentials{}.ToAttrType(ctx),
			},
			"azure_aad": basetypes.ListType{
				ElemType: AzureActiveDirectoryToken{}.ToAttrType(ctx),
			},
			"azure_user_delegation_sas": basetypes.ListType{
				ElemType: AzureUserDelegationSas{}.ToAttrType(ctx),
			},
			"expiration_time": types.Int64Type,
			"gcp_oauth_token": basetypes.ListType{
				ElemType: GcpOauthToken{}.ToAttrType(ctx),
			},
			"r2_temp_credentials": basetypes.ListType{
				ElemType: R2Credentials{}.ToAttrType(ctx),
			},
			"url": types.StringType,
		},
	}
}

// Gets the metastore assignment for a workspace
type GetAccountMetastoreAssignmentRequest struct {
	// Workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *GetAccountMetastoreAssignmentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountMetastoreAssignmentRequest) {
}

func (newState *GetAccountMetastoreAssignmentRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountMetastoreAssignmentRequest) {
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

// ToAttrType returns the representation of GetAccountMetastoreAssignmentRequest in the Terraform plugin framework type
// system.
func (a GetAccountMetastoreAssignmentRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetAccountMetastoreRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountMetastoreRequest) {
}

func (newState *GetAccountMetastoreRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountMetastoreRequest) {
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

// ToAttrType returns the representation of GetAccountMetastoreRequest in the Terraform plugin framework type
// system.
func (a GetAccountMetastoreRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetAccountStorageCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountStorageCredentialRequest) {
}

func (newState *GetAccountStorageCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountStorageCredentialRequest) {
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

// ToAttrType returns the representation of GetAccountStorageCredentialRequest in the Terraform plugin framework type
// system.
func (a GetAccountStorageCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetArtifactAllowlistRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetArtifactAllowlistRequest) {
}

func (newState *GetArtifactAllowlistRequest) SyncEffectiveFieldsDuringRead(existingState GetArtifactAllowlistRequest) {
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

// ToAttrType returns the representation of GetArtifactAllowlistRequest in the Terraform plugin framework type
// system.
func (a GetArtifactAllowlistRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	// The type of the securable to bind to a workspace.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *GetBindingsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetBindingsRequest) {
}

func (newState *GetBindingsRequest) SyncEffectiveFieldsDuringRead(existingState GetBindingsRequest) {
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

// ToAttrType returns the representation of GetBindingsRequest in the Terraform plugin framework type
// system.
func (a GetBindingsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetByAliasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetByAliasRequest) {
}

func (newState *GetByAliasRequest) SyncEffectiveFieldsDuringRead(existingState GetByAliasRequest) {
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

// ToAttrType returns the representation of GetByAliasRequest in the Terraform plugin framework type
// system.
func (a GetByAliasRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetCatalogRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCatalogRequest) {
}

func (newState *GetCatalogRequest) SyncEffectiveFieldsDuringRead(existingState GetCatalogRequest) {
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

// ToAttrType returns the representation of GetCatalogRequest in the Terraform plugin framework type
// system.
func (a GetCatalogRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_browse": types.BoolType,
			"name":           types.StringType,
		},
	}
}

// Get a connection
type GetConnectionRequest struct {
	// Name of the connection.
	Name types.String `tfsdk:"-"`
}

func (newState *GetConnectionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetConnectionRequest) {
}

func (newState *GetConnectionRequest) SyncEffectiveFieldsDuringRead(existingState GetConnectionRequest) {
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

// ToAttrType returns the representation of GetConnectionRequest in the Terraform plugin framework type
// system.
func (a GetConnectionRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

// ToAttrType returns the representation of GetCredentialRequest in the Terraform plugin framework type
// system.
func (a GetCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name_arg": types.StringType,
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

func (newState *GetEffectiveRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetEffectiveRequest) {
}

func (newState *GetEffectiveRequest) SyncEffectiveFieldsDuringRead(existingState GetEffectiveRequest) {
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

// ToAttrType returns the representation of GetEffectiveRequest in the Terraform plugin framework type
// system.
func (a GetEffectiveRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetExternalLocationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExternalLocationRequest) {
}

func (newState *GetExternalLocationRequest) SyncEffectiveFieldsDuringRead(existingState GetExternalLocationRequest) {
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

// ToAttrType returns the representation of GetExternalLocationRequest in the Terraform plugin framework type
// system.
func (a GetExternalLocationRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetFunctionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetFunctionRequest) {
}

func (newState *GetFunctionRequest) SyncEffectiveFieldsDuringRead(existingState GetFunctionRequest) {
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

// ToAttrType returns the representation of GetFunctionRequest in the Terraform plugin framework type
// system.
func (a GetFunctionRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetGrantRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetGrantRequest) {
}

func (newState *GetGrantRequest) SyncEffectiveFieldsDuringRead(existingState GetGrantRequest) {
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

// ToAttrType returns the representation of GetGrantRequest in the Terraform plugin framework type
// system.
func (a GetGrantRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetMetastoreRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetastoreRequest) {
}

func (newState *GetMetastoreRequest) SyncEffectiveFieldsDuringRead(existingState GetMetastoreRequest) {
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

// ToAttrType returns the representation of GetMetastoreRequest in the Terraform plugin framework type
// system.
func (a GetMetastoreRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetMetastoreSummaryResponse struct {
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

func (newState *GetMetastoreSummaryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetastoreSummaryResponse) {
}

func (newState *GetMetastoreSummaryResponse) SyncEffectiveFieldsDuringRead(existingState GetMetastoreSummaryResponse) {
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

// ToAttrType returns the representation of GetMetastoreSummaryResponse in the Terraform plugin framework type
// system.
func (a GetMetastoreSummaryResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionRequest) {
}

func (newState *GetModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState GetModelVersionRequest) {
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

// ToAttrType returns the representation of GetModelVersionRequest in the Terraform plugin framework type
// system.
func (a GetModelVersionRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetOnlineTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetOnlineTableRequest) {
}

func (newState *GetOnlineTableRequest) SyncEffectiveFieldsDuringRead(existingState GetOnlineTableRequest) {
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

// ToAttrType returns the representation of GetOnlineTableRequest in the Terraform plugin framework type
// system.
func (a GetOnlineTableRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetQualityMonitorRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQualityMonitorRequest) {
}

func (newState *GetQualityMonitorRequest) SyncEffectiveFieldsDuringRead(existingState GetQualityMonitorRequest) {
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

// ToAttrType returns the representation of GetQualityMonitorRequest in the Terraform plugin framework type
// system.
func (a GetQualityMonitorRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetQuotaRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQuotaRequest) {
}

func (newState *GetQuotaRequest) SyncEffectiveFieldsDuringRead(existingState GetQuotaRequest) {
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

// ToAttrType returns the representation of GetQuotaRequest in the Terraform plugin framework type
// system.
func (a GetQuotaRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	QuotaInfo types.List `tfsdk:"quota_info" tf:"optional,object"`
}

func (newState *GetQuotaResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQuotaResponse) {
}

func (newState *GetQuotaResponse) SyncEffectiveFieldsDuringRead(existingState GetQuotaResponse) {
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

// ToAttrType returns the representation of GetQuotaResponse in the Terraform plugin framework type
// system.
func (a GetQuotaResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quota_info": basetypes.ListType{
				ElemType: QuotaInfo{}.ToAttrType(ctx),
			},
		},
	}
}

// Get refresh
type GetRefreshRequest struct {
	// ID of the refresh.
	RefreshId types.String `tfsdk:"-"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *GetRefreshRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRefreshRequest) {
}

func (newState *GetRefreshRequest) SyncEffectiveFieldsDuringRead(existingState GetRefreshRequest) {
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

// ToAttrType returns the representation of GetRefreshRequest in the Terraform plugin framework type
// system.
func (a GetRefreshRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetRegisteredModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRegisteredModelRequest) {
}

func (newState *GetRegisteredModelRequest) SyncEffectiveFieldsDuringRead(existingState GetRegisteredModelRequest) {
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

// ToAttrType returns the representation of GetRegisteredModelRequest in the Terraform plugin framework type
// system.
func (a GetRegisteredModelRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetSchemaRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSchemaRequest) {
}

func (newState *GetSchemaRequest) SyncEffectiveFieldsDuringRead(existingState GetSchemaRequest) {
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

// ToAttrType returns the representation of GetSchemaRequest in the Terraform plugin framework type
// system.
func (a GetSchemaRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetStorageCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStorageCredentialRequest) {
}

func (newState *GetStorageCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GetStorageCredentialRequest) {
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

// ToAttrType returns the representation of GetStorageCredentialRequest in the Terraform plugin framework type
// system.
func (a GetStorageCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetTableRequest) {
}

func (newState *GetTableRequest) SyncEffectiveFieldsDuringRead(existingState GetTableRequest) {
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

// ToAttrType returns the representation of GetTableRequest in the Terraform plugin framework type
// system.
func (a GetTableRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GetWorkspaceBindingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceBindingRequest) {
}

func (newState *GetWorkspaceBindingRequest) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceBindingRequest) {
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

// ToAttrType returns the representation of GetWorkspaceBindingRequest in the Terraform plugin framework type
// system.
func (a GetWorkspaceBindingRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get all workspaces assigned to a metastore
type ListAccountMetastoreAssignmentsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *ListAccountMetastoreAssignmentsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountMetastoreAssignmentsRequest) {
}

func (newState *ListAccountMetastoreAssignmentsRequest) SyncEffectiveFieldsDuringRead(existingState ListAccountMetastoreAssignmentsRequest) {
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

// ToAttrType returns the representation of ListAccountMetastoreAssignmentsRequest in the Terraform plugin framework type
// system.
func (a ListAccountMetastoreAssignmentsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
		},
	}
}

// The list of workspaces to which the given metastore is assigned.
type ListAccountMetastoreAssignmentsResponse struct {
	WorkspaceIds types.List `tfsdk:"workspace_ids" tf:"optional"`
}

func (newState *ListAccountMetastoreAssignmentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountMetastoreAssignmentsResponse) {
}

func (newState *ListAccountMetastoreAssignmentsResponse) SyncEffectiveFieldsDuringRead(existingState ListAccountMetastoreAssignmentsResponse) {
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

// ToAttrType returns the representation of ListAccountMetastoreAssignmentsResponse in the Terraform plugin framework type
// system.
func (a ListAccountMetastoreAssignmentsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_ids": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// Get all storage credentials assigned to a metastore
type ListAccountStorageCredentialsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId types.String `tfsdk:"-"`
}

func (newState *ListAccountStorageCredentialsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountStorageCredentialsRequest) {
}

func (newState *ListAccountStorageCredentialsRequest) SyncEffectiveFieldsDuringRead(existingState ListAccountStorageCredentialsRequest) {
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

// ToAttrType returns the representation of ListAccountStorageCredentialsRequest in the Terraform plugin framework type
// system.
func (a ListAccountStorageCredentialsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
		},
	}
}

type ListAccountStorageCredentialsResponse struct {
	// An array of metastore storage credentials.
	StorageCredentials types.List `tfsdk:"storage_credentials" tf:"optional"`
}

func (newState *ListAccountStorageCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountStorageCredentialsResponse) {
}

func (newState *ListAccountStorageCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState ListAccountStorageCredentialsResponse) {
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

// ToAttrType returns the representation of ListAccountStorageCredentialsResponse in the Terraform plugin framework type
// system.
func (a ListAccountStorageCredentialsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_credentials": basetypes.ListType{
				ElemType: StorageCredentialInfo{}.ToAttrType(ctx),
			},
		},
	}
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

func (newState *ListCatalogsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCatalogsRequest) {
}

func (newState *ListCatalogsRequest) SyncEffectiveFieldsDuringRead(existingState ListCatalogsRequest) {
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

// ToAttrType returns the representation of ListCatalogsRequest in the Terraform plugin framework type
// system.
func (a ListCatalogsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	Catalogs types.List `tfsdk:"catalogs" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListCatalogsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCatalogsResponse) {
}

func (newState *ListCatalogsResponse) SyncEffectiveFieldsDuringRead(existingState ListCatalogsResponse) {
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

// ToAttrType returns the representation of ListCatalogsResponse in the Terraform plugin framework type
// system.
func (a ListCatalogsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalogs": basetypes.ListType{
				ElemType: CatalogInfo{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
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

func (newState *ListConnectionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListConnectionsRequest) {
}

func (newState *ListConnectionsRequest) SyncEffectiveFieldsDuringRead(existingState ListConnectionsRequest) {
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

// ToAttrType returns the representation of ListConnectionsRequest in the Terraform plugin framework type
// system.
func (a ListConnectionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListConnectionsResponse struct {
	// An array of connection information objects.
	Connections types.List `tfsdk:"connections" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListConnectionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListConnectionsResponse) {
}

func (newState *ListConnectionsResponse) SyncEffectiveFieldsDuringRead(existingState ListConnectionsResponse) {
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

// ToAttrType returns the representation of ListConnectionsResponse in the Terraform plugin framework type
// system.
func (a ListConnectionsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connections": basetypes.ListType{
				ElemType: ConnectionInfo{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
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

func (newState *ListCredentialsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCredentialsRequest) {
}

func (newState *ListCredentialsRequest) SyncEffectiveFieldsDuringRead(existingState ListCredentialsRequest) {
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

// ToAttrType returns the representation of ListCredentialsRequest in the Terraform plugin framework type
// system.
func (a ListCredentialsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
			"purpose":     types.StringType,
		},
	}
}

type ListCredentialsResponse struct {
	Credentials types.List `tfsdk:"credentials" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCredentialsResponse) {
}

func (newState *ListCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState ListCredentialsResponse) {
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

// ToAttrType returns the representation of ListCredentialsResponse in the Terraform plugin framework type
// system.
func (a ListCredentialsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials": basetypes.ListType{
				ElemType: CredentialInfo{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
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

func (newState *ListExternalLocationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExternalLocationsRequest) {
}

func (newState *ListExternalLocationsRequest) SyncEffectiveFieldsDuringRead(existingState ListExternalLocationsRequest) {
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

// ToAttrType returns the representation of ListExternalLocationsRequest in the Terraform plugin framework type
// system.
func (a ListExternalLocationsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	ExternalLocations types.List `tfsdk:"external_locations" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExternalLocationsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExternalLocationsResponse) {
}

func (newState *ListExternalLocationsResponse) SyncEffectiveFieldsDuringRead(existingState ListExternalLocationsResponse) {
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

// ToAttrType returns the representation of ListExternalLocationsResponse in the Terraform plugin framework type
// system.
func (a ListExternalLocationsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"external_locations": basetypes.ListType{
				ElemType: ExternalLocationInfo{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
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

func (newState *ListFunctionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFunctionsRequest) {
}

func (newState *ListFunctionsRequest) SyncEffectiveFieldsDuringRead(existingState ListFunctionsRequest) {
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

// ToAttrType returns the representation of ListFunctionsRequest in the Terraform plugin framework type
// system.
func (a ListFunctionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	Functions types.List `tfsdk:"functions" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListFunctionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFunctionsResponse) {
}

func (newState *ListFunctionsResponse) SyncEffectiveFieldsDuringRead(existingState ListFunctionsResponse) {
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

// ToAttrType returns the representation of ListFunctionsResponse in the Terraform plugin framework type
// system.
func (a ListFunctionsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"functions": basetypes.ListType{
				ElemType: FunctionInfo{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

type ListMetastoresResponse struct {
	// An array of metastore information objects.
	Metastores types.List `tfsdk:"metastores" tf:"optional"`
}

func (newState *ListMetastoresResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListMetastoresResponse) {
}

func (newState *ListMetastoresResponse) SyncEffectiveFieldsDuringRead(existingState ListMetastoresResponse) {
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

// ToAttrType returns the representation of ListMetastoresResponse in the Terraform plugin framework type
// system.
func (a ListMetastoresResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastores": basetypes.ListType{
				ElemType: MetastoreInfo{}.ToAttrType(ctx),
			},
		},
	}
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

func (newState *ListModelVersionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelVersionsRequest) {
}

func (newState *ListModelVersionsRequest) SyncEffectiveFieldsDuringRead(existingState ListModelVersionsRequest) {
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

// ToAttrType returns the representation of ListModelVersionsRequest in the Terraform plugin framework type
// system.
func (a ListModelVersionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	ModelVersions types.List `tfsdk:"model_versions" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListModelVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelVersionsResponse) {
}

func (newState *ListModelVersionsResponse) SyncEffectiveFieldsDuringRead(existingState ListModelVersionsResponse) {
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

// ToAttrType returns the representation of ListModelVersionsResponse in the Terraform plugin framework type
// system.
func (a ListModelVersionsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_versions": basetypes.ListType{
				ElemType: ModelVersionInfo{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// List all resource quotas under a metastore.
type ListQuotasRequest struct {
	// The number of quotas to return.
	MaxResults types.Int64 `tfsdk:"-"`
	// Opaque token for the next page of results.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListQuotasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQuotasRequest) {
}

func (newState *ListQuotasRequest) SyncEffectiveFieldsDuringRead(existingState ListQuotasRequest) {
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

// ToAttrType returns the representation of ListQuotasRequest in the Terraform plugin framework type
// system.
func (a ListQuotasRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of returned QuotaInfos.
	Quotas types.List `tfsdk:"quotas" tf:"optional"`
}

func (newState *ListQuotasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQuotasResponse) {
}

func (newState *ListQuotasResponse) SyncEffectiveFieldsDuringRead(existingState ListQuotasResponse) {
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

// ToAttrType returns the representation of ListQuotasResponse in the Terraform plugin framework type
// system.
func (a ListQuotasResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"quotas": basetypes.ListType{
				ElemType: QuotaInfo{}.ToAttrType(ctx),
			},
		},
	}
}

// List refreshes
type ListRefreshesRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *ListRefreshesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRefreshesRequest) {
}

func (newState *ListRefreshesRequest) SyncEffectiveFieldsDuringRead(existingState ListRefreshesRequest) {
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

// ToAttrType returns the representation of ListRefreshesRequest in the Terraform plugin framework type
// system.
func (a ListRefreshesRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *ListRegisteredModelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRegisteredModelsRequest) {
}

func (newState *ListRegisteredModelsRequest) SyncEffectiveFieldsDuringRead(existingState ListRegisteredModelsRequest) {
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

// ToAttrType returns the representation of ListRegisteredModelsRequest in the Terraform plugin framework type
// system.
func (a ListRegisteredModelsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	RegisteredModels types.List `tfsdk:"registered_models" tf:"optional"`
}

func (newState *ListRegisteredModelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRegisteredModelsResponse) {
}

func (newState *ListRegisteredModelsResponse) SyncEffectiveFieldsDuringRead(existingState ListRegisteredModelsResponse) {
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

// ToAttrType returns the representation of ListRegisteredModelsResponse in the Terraform plugin framework type
// system.
func (a ListRegisteredModelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"registered_models": basetypes.ListType{
				ElemType: RegisteredModelInfo{}.ToAttrType(ctx),
			},
		},
	}
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

func (newState *ListSchemasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchemasRequest) {
}

func (newState *ListSchemasRequest) SyncEffectiveFieldsDuringRead(existingState ListSchemasRequest) {
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

// ToAttrType returns the representation of ListSchemasRequest in the Terraform plugin framework type
// system.
func (a ListSchemasRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of schema information objects.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
}

func (newState *ListSchemasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchemasResponse) {
}

func (newState *ListSchemasResponse) SyncEffectiveFieldsDuringRead(existingState ListSchemasResponse) {
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

// ToAttrType returns the representation of ListSchemasResponse in the Terraform plugin framework type
// system.
func (a ListSchemasResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"schemas": basetypes.ListType{
				ElemType: SchemaInfo{}.ToAttrType(ctx),
			},
		},
	}
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

func (newState *ListStorageCredentialsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListStorageCredentialsRequest) {
}

func (newState *ListStorageCredentialsRequest) SyncEffectiveFieldsDuringRead(existingState ListStorageCredentialsRequest) {
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

// ToAttrType returns the representation of ListStorageCredentialsRequest in the Terraform plugin framework type
// system.
func (a ListStorageCredentialsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	StorageCredentials types.List `tfsdk:"storage_credentials" tf:"optional"`
}

func (newState *ListStorageCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListStorageCredentialsResponse) {
}

func (newState *ListStorageCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState ListStorageCredentialsResponse) {
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

// ToAttrType returns the representation of ListStorageCredentialsResponse in the Terraform plugin framework type
// system.
func (a ListStorageCredentialsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"storage_credentials": basetypes.ListType{
				ElemType: StorageCredentialInfo{}.ToAttrType(ctx),
			},
		},
	}
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

func (newState *ListSummariesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSummariesRequest) {
}

func (newState *ListSummariesRequest) SyncEffectiveFieldsDuringRead(existingState ListSummariesRequest) {
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

// ToAttrType returns the representation of ListSummariesRequest in the Terraform plugin framework type
// system.
func (a ListSummariesRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *ListSystemSchemasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSystemSchemasRequest) {
}

func (newState *ListSystemSchemasRequest) SyncEffectiveFieldsDuringRead(existingState ListSystemSchemasRequest) {
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

// ToAttrType returns the representation of ListSystemSchemasRequest in the Terraform plugin framework type
// system.
func (a ListSystemSchemasRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of system schema information objects.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
}

func (newState *ListSystemSchemasResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSystemSchemasResponse) {
}

func (newState *ListSystemSchemasResponse) SyncEffectiveFieldsDuringRead(existingState ListSystemSchemasResponse) {
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

// ToAttrType returns the representation of ListSystemSchemasResponse in the Terraform plugin framework type
// system.
func (a ListSystemSchemasResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"schemas": basetypes.ListType{
				ElemType: SystemSchemaInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type ListTableSummariesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// List of table summaries.
	Tables types.List `tfsdk:"tables" tf:"optional"`
}

func (newState *ListTableSummariesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTableSummariesResponse) {
}

func (newState *ListTableSummariesResponse) SyncEffectiveFieldsDuringRead(existingState ListTableSummariesResponse) {
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

// ToAttrType returns the representation of ListTableSummariesResponse in the Terraform plugin framework type
// system.
func (a ListTableSummariesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tables": basetypes.ListType{
				ElemType: TableSummary{}.ToAttrType(ctx),
			},
		},
	}
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

func (newState *ListTablesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTablesRequest) {
}

func (newState *ListTablesRequest) SyncEffectiveFieldsDuringRead(existingState ListTablesRequest) {
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

// ToAttrType returns the representation of ListTablesRequest in the Terraform plugin framework type
// system.
func (a ListTablesRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// An array of table information objects.
	Tables types.List `tfsdk:"tables" tf:"optional"`
}

func (newState *ListTablesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTablesResponse) {
}

func (newState *ListTablesResponse) SyncEffectiveFieldsDuringRead(existingState ListTablesResponse) {
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

// ToAttrType returns the representation of ListTablesResponse in the Terraform plugin framework type
// system.
func (a ListTablesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tables": basetypes.ListType{
				ElemType: TableInfo{}.ToAttrType(ctx),
			},
		},
	}
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

func (newState *ListVolumesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVolumesRequest) {
}

func (newState *ListVolumesRequest) SyncEffectiveFieldsDuringRead(existingState ListVolumesRequest) {
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

// ToAttrType returns the representation of ListVolumesRequest in the Terraform plugin framework type
// system.
func (a ListVolumesRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Volumes types.List `tfsdk:"volumes" tf:"optional"`
}

func (newState *ListVolumesResponseContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVolumesResponseContent) {
}

func (newState *ListVolumesResponseContent) SyncEffectiveFieldsDuringRead(existingState ListVolumesResponseContent) {
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

// ToAttrType returns the representation of ListVolumesResponseContent in the Terraform plugin framework type
// system.
func (a ListVolumesResponseContent) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"volumes": basetypes.ListType{
				ElemType: VolumeInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type MetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	DefaultCatalogName types.String `tfsdk:"default_catalog_name" tf:"optional"`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:""`
	// The unique ID of the Databricks workspace.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:""`
}

func (newState *MetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan MetastoreAssignment) {
}

func (newState *MetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState MetastoreAssignment) {
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

// ToAttrType returns the representation of MetastoreAssignment in the Terraform plugin framework type
// system.
func (a MetastoreAssignment) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *MetastoreInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan MetastoreInfo) {
}

func (newState *MetastoreInfo) SyncEffectiveFieldsDuringRead(existingState MetastoreInfo) {
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

// ToAttrType returns the representation of MetastoreInfo in the Terraform plugin framework type
// system.
func (a MetastoreInfo) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *ModelVersionInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersionInfo) {
}

func (newState *ModelVersionInfo) SyncEffectiveFieldsDuringRead(existingState ModelVersionInfo) {
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

// ToAttrType returns the representation of ModelVersionInfo in the Terraform plugin framework type
// system.
func (a ModelVersionInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aliases": basetypes.ListType{
				ElemType: RegisteredModelAlias{}.ToAttrType(ctx),
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
				ElemType: DependencyList{}.ToAttrType(ctx),
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

type MonitorCronSchedule struct {
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

func (newState *MonitorCronSchedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorCronSchedule) {
}

func (newState *MonitorCronSchedule) SyncEffectiveFieldsDuringRead(existingState MonitorCronSchedule) {
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

// ToAttrType returns the representation of MonitorCronSchedule in the Terraform plugin framework type
// system.
func (a MonitorCronSchedule) ToAttrType(ctx context.Context) types.ObjectType {
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
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
}

func (newState *MonitorDataClassificationConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorDataClassificationConfig) {
}

func (newState *MonitorDataClassificationConfig) SyncEffectiveFieldsDuringRead(existingState MonitorDataClassificationConfig) {
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

// ToAttrType returns the representation of MonitorDataClassificationConfig in the Terraform plugin framework type
// system.
func (a MonitorDataClassificationConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

type MonitorDestination struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	EmailAddresses types.List `tfsdk:"email_addresses" tf:"optional"`
}

func (newState *MonitorDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorDestination) {
}

func (newState *MonitorDestination) SyncEffectiveFieldsDuringRead(existingState MonitorDestination) {
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

// ToAttrType returns the representation of MonitorDestination in the Terraform plugin framework type
// system.
func (a MonitorDestination) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type MonitorInferenceLog struct {
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

func (newState *MonitorInferenceLog) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorInferenceLog) {
}

func (newState *MonitorInferenceLog) SyncEffectiveFieldsDuringRead(existingState MonitorInferenceLog) {
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

// ToAttrType returns the representation of MonitorInferenceLog in the Terraform plugin framework type
// system.
func (a MonitorInferenceLog) ToAttrType(ctx context.Context) types.ObjectType {
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

type MonitorInfo struct {
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
	Snapshot []MonitorSnapshot `tfsdk:"snapshot" tf:"optional,object"`
	// The status of the monitor.
	Status types.String `tfsdk:"status" tf:""`
	// The full name of the table to monitor. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	TableName types.String `tfsdk:"table_name" tf:""`
	// Configuration for monitoring time series tables.
	TimeSeries types.List `tfsdk:"time_series" tf:"optional,object"`
}

func (newState *MonitorInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorInfo) {
}

func (newState *MonitorInfo) SyncEffectiveFieldsDuringRead(existingState MonitorInfo) {
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

// ToAttrType returns the representation of MonitorInfo in the Terraform plugin framework type
// system.
func (a MonitorInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets_dir":          types.StringType,
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: MonitorMetric{}.ToAttrType(ctx),
			},
			"dashboard_id": types.StringType,
			"data_classification_config": basetypes.ListType{
				ElemType: MonitorDataClassificationConfig{}.ToAttrType(ctx),
			},
			"drift_metrics_table_name": types.StringType,
			"inference_log": basetypes.ListType{
				ElemType: MonitorInferenceLog{}.ToAttrType(ctx),
			},
			"latest_monitor_failure_msg": types.StringType,
			"monitor_version":            types.StringType,
			"notifications": basetypes.ListType{
				ElemType: MonitorNotifications{}.ToAttrType(ctx),
			},
			"output_schema_name":         types.StringType,
			"profile_metrics_table_name": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: MonitorCronSchedule{}.ToAttrType(ctx),
			},
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot": basetypes.ListType{
				ElemType: MonitorSnapshot{}.ToAttrType(ctx),
			},
			"status":     types.StringType,
			"table_name": types.StringType,
			"time_series": basetypes.ListType{
				ElemType: MonitorTimeSeries{}.ToAttrType(ctx),
			},
		},
	}
}

type MonitorMetric struct {
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
	Type types.String `tfsdk:"type" tf:""`
}

func (newState *MonitorMetric) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorMetric) {
}

func (newState *MonitorMetric) SyncEffectiveFieldsDuringRead(existingState MonitorMetric) {
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

// ToAttrType returns the representation of MonitorMetric in the Terraform plugin framework type
// system.
func (a MonitorMetric) ToAttrType(ctx context.Context) types.ObjectType {
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

type MonitorNotifications struct {
	// Who to send notifications to on monitor failure.
	OnFailure types.List `tfsdk:"on_failure" tf:"optional,object"`
	// Who to send notifications to when new data classification tags are
	// detected.
	OnNewClassificationTagDetected types.List `tfsdk:"on_new_classification_tag_detected" tf:"optional,object"`
}

func (newState *MonitorNotifications) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorNotifications) {
}

func (newState *MonitorNotifications) SyncEffectiveFieldsDuringRead(existingState MonitorNotifications) {
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

// ToAttrType returns the representation of MonitorNotifications in the Terraform plugin framework type
// system.
func (a MonitorNotifications) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"on_failure": basetypes.ListType{
				ElemType: MonitorDestination{}.ToAttrType(ctx),
			},
			"on_new_classification_tag_detected": basetypes.ListType{
				ElemType: MonitorDestination{}.ToAttrType(ctx),
			},
		},
	}
}

type MonitorRefreshInfo struct {
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

func (newState *MonitorRefreshInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorRefreshInfo) {
}

func (newState *MonitorRefreshInfo) SyncEffectiveFieldsDuringRead(existingState MonitorRefreshInfo) {
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

// ToAttrType returns the representation of MonitorRefreshInfo in the Terraform plugin framework type
// system.
func (a MonitorRefreshInfo) ToAttrType(ctx context.Context) types.ObjectType {
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
	Refreshes types.List `tfsdk:"refreshes" tf:"optional"`
}

func (newState *MonitorRefreshListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorRefreshListResponse) {
}

func (newState *MonitorRefreshListResponse) SyncEffectiveFieldsDuringRead(existingState MonitorRefreshListResponse) {
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

// ToAttrType returns the representation of MonitorRefreshListResponse in the Terraform plugin framework type
// system.
func (a MonitorRefreshListResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"refreshes": basetypes.ListType{
				ElemType: MonitorRefreshInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type MonitorSnapshot struct {
}

func (newState *MonitorSnapshot) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorSnapshot) {
}

func (newState *MonitorSnapshot) SyncEffectiveFieldsDuringRead(existingState MonitorSnapshot) {
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

// ToAttrType returns the representation of MonitorSnapshot in the Terraform plugin framework type
// system.
func (a MonitorSnapshot) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type MonitorTimeSeries struct {
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

func (newState *MonitorTimeSeries) SyncEffectiveFieldsDuringCreateOrUpdate(plan MonitorTimeSeries) {
}

func (newState *MonitorTimeSeries) SyncEffectiveFieldsDuringRead(existingState MonitorTimeSeries) {
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

// ToAttrType returns the representation of MonitorTimeSeries in the Terraform plugin framework type
// system.
func (a MonitorTimeSeries) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"granularities": basetypes.ListType{
				ElemType: types.StringType,
			},
			"timestamp_col": types.StringType,
		},
	}
}

type NamedTableConstraint struct {
	// The name of the constraint.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *NamedTableConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan NamedTableConstraint) {
}

func (newState *NamedTableConstraint) SyncEffectiveFieldsDuringRead(existingState NamedTableConstraint) {
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

// ToAttrType returns the representation of NamedTableConstraint in the Terraform plugin framework type
// system.
func (a NamedTableConstraint) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Online Table information.
type OnlineTable struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Specification of the online table.
	Spec types.List `tfsdk:"spec" tf:"optional,object"`
	// Online Table data synchronization status
	Status types.List `tfsdk:"status" tf:"optional,object"`
	// Data serving REST API URL for this table
	TableServingUrl types.String `tfsdk:"table_serving_url" tf:"computed,optional"`
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState types.String `tfsdk:"unity_catalog_provisioning_state" tf:"optional"`
}

func (newState *OnlineTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTable) {
}

func (newState *OnlineTable) SyncEffectiveFieldsDuringRead(existingState OnlineTable) {
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

// ToAttrType returns the representation of OnlineTable in the Terraform plugin framework type
// system.
func (a OnlineTable) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"spec": basetypes.ListType{
				ElemType: OnlineTableSpec{}.ToAttrType(ctx),
			},
			"status": basetypes.ListType{
				ElemType: OnlineTableStatus{}.ToAttrType(ctx),
			},
			"table_serving_url":                types.StringType,
			"unity_catalog_provisioning_state": types.StringType,
		},
	}
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
	PerformFullCopy types.Bool `tfsdk:"perform_full_copy" tf:"optional"`
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	PipelineId types.String `tfsdk:"pipeline_id" tf:"computed,optional"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns types.List `tfsdk:"primary_key_columns" tf:"optional"`
	// Pipeline runs continuously after generating the initial data.
	RunContinuously []OnlineTableSpecContinuousSchedulingPolicy `tfsdk:"run_continuously" tf:"optional,object"`
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	RunTriggered []OnlineTableSpecTriggeredSchedulingPolicy `tfsdk:"run_triggered" tf:"optional,object"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName types.String `tfsdk:"source_table_full_name" tf:"optional"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey types.String `tfsdk:"timeseries_key" tf:"optional"`
}

func (newState *OnlineTableSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpec) {
}

func (newState *OnlineTableSpec) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpec) {
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

// ToAttrType returns the representation of OnlineTableSpec in the Terraform plugin framework type
// system.
func (a OnlineTableSpec) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"perform_full_copy": types.BoolType,
			"pipeline_id":       types.StringType,
			"primary_key_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"run_continuously": basetypes.ListType{
				ElemType: OnlineTableSpecContinuousSchedulingPolicy{}.ToAttrType(ctx),
			},
			"run_triggered": basetypes.ListType{
				ElemType: OnlineTableSpecTriggeredSchedulingPolicy{}.ToAttrType(ctx),
			},
			"source_table_full_name": types.StringType,
			"timeseries_key":         types.StringType,
		},
	}
}

type OnlineTableSpecContinuousSchedulingPolicy struct {
}

func (newState *OnlineTableSpecContinuousSchedulingPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableSpecContinuousSchedulingPolicy) {
}

func (newState *OnlineTableSpecContinuousSchedulingPolicy) SyncEffectiveFieldsDuringRead(existingState OnlineTableSpecContinuousSchedulingPolicy) {
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

// ToAttrType returns the representation of OnlineTableSpecContinuousSchedulingPolicy in the Terraform plugin framework type
// system.
func (a OnlineTableSpecContinuousSchedulingPolicy) ToAttrType(ctx context.Context) types.ObjectType {
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

// ToAttrType returns the representation of OnlineTableSpecTriggeredSchedulingPolicy in the Terraform plugin framework type
// system.
func (a OnlineTableSpecTriggeredSchedulingPolicy) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Status of an online table.
type OnlineTableStatus struct {
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

func (newState *OnlineTableStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineTableStatus) {
}

func (newState *OnlineTableStatus) SyncEffectiveFieldsDuringRead(existingState OnlineTableStatus) {
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

// ToAttrType returns the representation of OnlineTableStatus in the Terraform plugin framework type
// system.
func (a OnlineTableStatus) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"continuous_update_status": basetypes.ListType{
				ElemType: ContinuousUpdateStatus{}.ToAttrType(ctx),
			},
			"detailed_state": types.StringType,
			"failed_status": basetypes.ListType{
				ElemType: FailedStatus{}.ToAttrType(ctx),
			},
			"message": types.StringType,
			"provisioning_status": basetypes.ListType{
				ElemType: ProvisioningStatus{}.ToAttrType(ctx),
			},
			"triggered_update_status": basetypes.ListType{
				ElemType: TriggeredUpdateStatus{}.ToAttrType(ctx),
			},
		},
	}
}

type PermissionsChange struct {
	// The set of privileges to add.
	Add types.List `tfsdk:"add" tf:"optional"`
	// The principal whose privileges we are changing.
	Principal types.String `tfsdk:"principal" tf:"optional"`
	// The set of privileges to remove.
	Remove types.List `tfsdk:"remove" tf:"optional"`
}

func (newState *PermissionsChange) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsChange) {
}

func (newState *PermissionsChange) SyncEffectiveFieldsDuringRead(existingState PermissionsChange) {
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

// ToAttrType returns the representation of PermissionsChange in the Terraform plugin framework type
// system.
func (a PermissionsChange) ToAttrType(ctx context.Context) types.ObjectType {
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

type PermissionsList struct {
	// The privileges assigned to each principal
	PrivilegeAssignments types.List `tfsdk:"privilege_assignments" tf:"optional"`
}

func (newState *PermissionsList) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsList) {
}

func (newState *PermissionsList) SyncEffectiveFieldsDuringRead(existingState PermissionsList) {
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

// ToAttrType returns the representation of PermissionsList in the Terraform plugin framework type
// system.
func (a PermissionsList) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"privilege_assignments": basetypes.ListType{
				ElemType: PrivilegeAssignment{}.ToAttrType(ctx),
			},
		},
	}
}

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress struct {
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

func (newState *PipelineProgress) SyncEffectiveFieldsDuringCreateOrUpdate(plan PipelineProgress) {
}

func (newState *PipelineProgress) SyncEffectiveFieldsDuringRead(existingState PipelineProgress) {
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

// ToAttrType returns the representation of PipelineProgress in the Terraform plugin framework type
// system.
func (a PipelineProgress) ToAttrType(ctx context.Context) types.ObjectType {
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
	ChildColumns types.List `tfsdk:"child_columns" tf:""`
	// The name of the constraint.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *PrimaryKeyConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrimaryKeyConstraint) {
}

func (newState *PrimaryKeyConstraint) SyncEffectiveFieldsDuringRead(existingState PrimaryKeyConstraint) {
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
		"child_columns": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of PrimaryKeyConstraint in the Terraform plugin framework type
// system.
func (a PrimaryKeyConstraint) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"child_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"name": types.StringType,
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

// ToAttrType returns the representation of PrivilegeAssignment in the Terraform plugin framework type
// system.
func (a PrivilegeAssignment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"privileges": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// Status of an asynchronously provisioned resource.
type ProvisioningInfo struct {
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *ProvisioningInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProvisioningInfo) {
}

func (newState *ProvisioningInfo) SyncEffectiveFieldsDuringRead(existingState ProvisioningInfo) {
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

// ToAttrType returns the representation of ProvisioningInfo in the Terraform plugin framework type
// system.
func (a ProvisioningInfo) ToAttrType(ctx context.Context) types.ObjectType {
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
	InitialPipelineSyncProgress types.List `tfsdk:"initial_pipeline_sync_progress" tf:"optional,object"`
}

func (newState *ProvisioningStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan ProvisioningStatus) {
}

func (newState *ProvisioningStatus) SyncEffectiveFieldsDuringRead(existingState ProvisioningStatus) {
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

// ToAttrType returns the representation of ProvisioningStatus in the Terraform plugin framework type
// system.
func (a ProvisioningStatus) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": basetypes.ListType{
				ElemType: PipelineProgress{}.ToAttrType(ctx),
			},
		},
	}
}

type QuotaInfo struct {
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

func (newState *QuotaInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan QuotaInfo) {
}

func (newState *QuotaInfo) SyncEffectiveFieldsDuringRead(existingState QuotaInfo) {
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

// ToAttrType returns the representation of QuotaInfo in the Terraform plugin framework type
// system.
func (a QuotaInfo) ToAttrType(ctx context.Context) types.ObjectType {
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
	AccessKeyId types.String `tfsdk:"access_key_id" tf:"optional"`
	// The secret access key associated with the access key.
	SecretAccessKey types.String `tfsdk:"secret_access_key" tf:"optional"`
	// The generated JWT that users must pass to use the temporary credentials.
	SessionToken types.String `tfsdk:"session_token" tf:"optional"`
}

func (newState *R2Credentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan R2Credentials) {
}

func (newState *R2Credentials) SyncEffectiveFieldsDuringRead(existingState R2Credentials) {
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

// ToAttrType returns the representation of R2Credentials in the Terraform plugin framework type
// system.
func (a R2Credentials) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *ReadVolumeRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReadVolumeRequest) {
}

func (newState *ReadVolumeRequest) SyncEffectiveFieldsDuringRead(existingState ReadVolumeRequest) {
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

// ToAttrType returns the representation of ReadVolumeRequest in the Terraform plugin framework type
// system.
func (a ReadVolumeRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *RegenerateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegenerateDashboardRequest) {
}

func (newState *RegenerateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState RegenerateDashboardRequest) {
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

// ToAttrType returns the representation of RegenerateDashboardRequest in the Terraform plugin framework type
// system.
func (a RegenerateDashboardRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name":   types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

type RegenerateDashboardResponse struct {
	// Id of the regenerated monitoring dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The directory where the regenerated dashboard is stored.
	ParentFolder types.String `tfsdk:"parent_folder" tf:"optional"`
}

func (newState *RegenerateDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegenerateDashboardResponse) {
}

func (newState *RegenerateDashboardResponse) SyncEffectiveFieldsDuringRead(existingState RegenerateDashboardResponse) {
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

// ToAttrType returns the representation of RegenerateDashboardResponse in the Terraform plugin framework type
// system.
func (a RegenerateDashboardResponse) ToAttrType(ctx context.Context) types.ObjectType {
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
	AliasName types.String `tfsdk:"alias_name" tf:"optional"`
	// Integer version number of the model version to which this alias points.
	VersionNum types.Int64 `tfsdk:"version_num" tf:"optional"`
}

func (newState *RegisteredModelAlias) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelAlias) {
}

func (newState *RegisteredModelAlias) SyncEffectiveFieldsDuringRead(existingState RegisteredModelAlias) {
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

// ToAttrType returns the representation of RegisteredModelAlias in the Terraform plugin framework type
// system.
func (a RegisteredModelAlias) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"alias_name":  types.StringType,
			"version_num": types.Int64Type,
		},
	}
}

type RegisteredModelInfo struct {
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

func (newState *RegisteredModelInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelInfo) {
}

func (newState *RegisteredModelInfo) SyncEffectiveFieldsDuringRead(existingState RegisteredModelInfo) {
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

// ToAttrType returns the representation of RegisteredModelInfo in the Terraform plugin framework type
// system.
func (a RegisteredModelInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aliases": basetypes.ListType{
				ElemType: RegisteredModelAlias{}.ToAttrType(ctx),
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

// Queue a metric refresh for a monitor
type RunRefreshRequest struct {
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
}

func (newState *RunRefreshRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunRefreshRequest) {
}

func (newState *RunRefreshRequest) SyncEffectiveFieldsDuringRead(existingState RunRefreshRequest) {
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

// ToAttrType returns the representation of RunRefreshRequest in the Terraform plugin framework type
// system.
func (a RunRefreshRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name": types.StringType,
		},
	}
}

type SchemaInfo struct {
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

func (newState *SchemaInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan SchemaInfo) {
}

func (newState *SchemaInfo) SyncEffectiveFieldsDuringRead(existingState SchemaInfo) {
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

// ToAttrType returns the representation of SchemaInfo in the Terraform plugin framework type
// system.
func (a SchemaInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"browse_only":  types.BoolType,
			"catalog_name": types.StringType,
			"catalog_type": types.StringType,
			"comment":      types.StringType,
			"created_at":   types.Int64Type,
			"created_by":   types.StringType,
			"effective_predictive_optimization_flag": basetypes.ListType{
				ElemType: EffectivePredictiveOptimizationFlag{}.ToAttrType(ctx),
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

type SetArtifactAllowlist struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers types.List `tfsdk:"artifact_matchers" tf:""`
	// The artifact type of the allowlist.
	ArtifactType types.String `tfsdk:"-"`
}

func (newState *SetArtifactAllowlist) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetArtifactAllowlist) {
}

func (newState *SetArtifactAllowlist) SyncEffectiveFieldsDuringRead(existingState SetArtifactAllowlist) {
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

// ToAttrType returns the representation of SetArtifactAllowlist in the Terraform plugin framework type
// system.
func (a SetArtifactAllowlist) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_matchers": basetypes.ListType{
				ElemType: ArtifactMatcher{}.ToAttrType(ctx),
			},
			"artifact_type": types.StringType,
		},
	}
}

type SetRegisteredModelAliasRequest struct {
	// The name of the alias
	Alias types.String `tfsdk:"alias" tf:""`
	// Full name of the registered model
	FullName types.String `tfsdk:"full_name" tf:""`
	// The version number of the model version to which the alias points
	VersionNum types.Int64 `tfsdk:"version_num" tf:""`
}

func (newState *SetRegisteredModelAliasRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetRegisteredModelAliasRequest) {
}

func (newState *SetRegisteredModelAliasRequest) SyncEffectiveFieldsDuringRead(existingState SetRegisteredModelAliasRequest) {
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

// ToAttrType returns the representation of SetRegisteredModelAliasRequest in the Terraform plugin framework type
// system.
func (a SetRegisteredModelAliasRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	// The type of key encryption to use (affects headers from s3 client).
	Algorithm types.String `tfsdk:"algorithm" tf:"optional"`
	// When algorithm is **AWS_SSE_KMS** this field specifies the ARN of the SSE
	// key to use.
	AwsKmsKeyArn types.String `tfsdk:"aws_kms_key_arn" tf:"optional"`
}

func (newState *SseEncryptionDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan SseEncryptionDetails) {
}

func (newState *SseEncryptionDetails) SyncEffectiveFieldsDuringRead(existingState SseEncryptionDetails) {
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

// ToAttrType returns the representation of SseEncryptionDetails in the Terraform plugin framework type
// system.
func (a SseEncryptionDetails) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"algorithm":       types.StringType,
			"aws_kms_key_arn": types.StringType,
		},
	}
}

type StorageCredentialInfo struct {
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

func (newState *StorageCredentialInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan StorageCredentialInfo) {
}

func (newState *StorageCredentialInfo) SyncEffectiveFieldsDuringRead(existingState StorageCredentialInfo) {
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

// ToAttrType returns the representation of StorageCredentialInfo in the Terraform plugin framework type
// system.
func (a StorageCredentialInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRoleResponse{}.ToAttrType(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentityResponse{}.ToAttrType(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal{}.ToAttrType(ctx),
			},
			"cloudflare_api_token": basetypes.ListType{
				ElemType: CloudflareApiToken{}.ToAttrType(ctx),
			},
			"comment":    types.StringType,
			"created_at": types.Int64Type,
			"created_by": types.StringType,
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccountResponse{}.ToAttrType(ctx),
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

type SystemSchemaInfo struct {
	// Name of the system schema.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in.
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *SystemSchemaInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan SystemSchemaInfo) {
}

func (newState *SystemSchemaInfo) SyncEffectiveFieldsDuringRead(existingState SystemSchemaInfo) {
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

// ToAttrType returns the representation of SystemSchemaInfo in the Terraform plugin framework type
// system.
func (a SystemSchemaInfo) ToAttrType(ctx context.Context) types.ObjectType {
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
	ForeignKeyConstraint types.List `tfsdk:"foreign_key_constraint" tf:"optional,object"`

	NamedTableConstraint types.List `tfsdk:"named_table_constraint" tf:"optional,object"`

	PrimaryKeyConstraint types.List `tfsdk:"primary_key_constraint" tf:"optional,object"`
}

func (newState *TableConstraint) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableConstraint) {
}

func (newState *TableConstraint) SyncEffectiveFieldsDuringRead(existingState TableConstraint) {
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

// ToAttrType returns the representation of TableConstraint in the Terraform plugin framework type
// system.
func (a TableConstraint) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"foreign_key_constraint": basetypes.ListType{
				ElemType: ForeignKeyConstraint{}.ToAttrType(ctx),
			},
			"named_table_constraint": basetypes.ListType{
				ElemType: NamedTableConstraint{}.ToAttrType(ctx),
			},
			"primary_key_constraint": basetypes.ListType{
				ElemType: PrimaryKeyConstraint{}.ToAttrType(ctx),
			},
		},
	}
}

// A table that is dependent on a SQL object.
type TableDependency struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	TableFullName types.String `tfsdk:"table_full_name" tf:""`
}

func (newState *TableDependency) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableDependency) {
}

func (newState *TableDependency) SyncEffectiveFieldsDuringRead(existingState TableDependency) {
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

// ToAttrType returns the representation of TableDependency in the Terraform plugin framework type
// system.
func (a TableDependency) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_full_name": types.StringType,
		},
	}
}

type TableExistsResponse struct {
	// Whether the table exists or not.
	TableExists types.Bool `tfsdk:"table_exists" tf:"optional"`
}

func (newState *TableExistsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableExistsResponse) {
}

func (newState *TableExistsResponse) SyncEffectiveFieldsDuringRead(existingState TableExistsResponse) {
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

// ToAttrType returns the representation of TableExistsResponse in the Terraform plugin framework type
// system.
func (a TableExistsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_exists": types.BoolType,
		},
	}
}

type TableInfo struct {
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

func (newState *TableInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableInfo) {
}

func (newState *TableInfo) SyncEffectiveFieldsDuringRead(existingState TableInfo) {
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

// ToAttrType returns the representation of TableInfo in the Terraform plugin framework type
// system.
func (a TableInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point": types.StringType,
			"browse_only":  types.BoolType,
			"catalog_name": types.StringType,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo{}.ToAttrType(ctx),
			},
			"comment":                      types.StringType,
			"created_at":                   types.Int64Type,
			"created_by":                   types.StringType,
			"data_access_configuration_id": types.StringType,
			"data_source_format":           types.StringType,
			"deleted_at":                   types.Int64Type,
			"delta_runtime_properties_kvpairs": basetypes.ListType{
				ElemType: DeltaRuntimePropertiesKvPairs{}.ToAttrType(ctx),
			},
			"effective_predictive_optimization_flag": basetypes.ListType{
				ElemType: EffectivePredictiveOptimizationFlag{}.ToAttrType(ctx),
			},
			"enable_predictive_optimization": types.StringType,
			"encryption_details": basetypes.ListType{
				ElemType: EncryptionDetails{}.ToAttrType(ctx),
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
				ElemType: TableRowFilter{}.ToAttrType(ctx),
			},
			"schema_name":             types.StringType,
			"sql_path":                types.StringType,
			"storage_credential_name": types.StringType,
			"storage_location":        types.StringType,
			"table_constraints": basetypes.ListType{
				ElemType: TableConstraint{}.ToAttrType(ctx),
			},
			"table_id":        types.StringType,
			"table_type":      types.StringType,
			"updated_at":      types.Int64Type,
			"updated_by":      types.StringType,
			"view_definition": types.StringType,
			"view_dependencies": basetypes.ListType{
				ElemType: DependencyList{}.ToAttrType(ctx),
			},
		},
	}
}

type TableRowFilter struct {
	// The full name of the row filter SQL UDF.
	FunctionName types.String `tfsdk:"function_name" tf:""`
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	InputColumnNames types.List `tfsdk:"input_column_names" tf:""`
}

func (newState *TableRowFilter) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableRowFilter) {
}

func (newState *TableRowFilter) SyncEffectiveFieldsDuringRead(existingState TableRowFilter) {
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

// ToAttrType returns the representation of TableRowFilter in the Terraform plugin framework type
// system.
func (a TableRowFilter) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"function_name": types.StringType,
			"input_column_names": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type TableSummary struct {
	// The full name of the table.
	FullName types.String `tfsdk:"full_name" tf:"optional"`

	TableType types.String `tfsdk:"table_type" tf:"optional"`
}

func (newState *TableSummary) SyncEffectiveFieldsDuringCreateOrUpdate(plan TableSummary) {
}

func (newState *TableSummary) SyncEffectiveFieldsDuringRead(existingState TableSummary) {
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

// ToAttrType returns the representation of TableSummary in the Terraform plugin framework type
// system.
func (a TableSummary) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name":  types.StringType,
			"table_type": types.StringType,
		},
	}
}

type TemporaryCredentials struct {
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

func (newState *TemporaryCredentials) SyncEffectiveFieldsDuringCreateOrUpdate(plan TemporaryCredentials) {
}

func (newState *TemporaryCredentials) SyncEffectiveFieldsDuringRead(existingState TemporaryCredentials) {
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
	}
}

// ToAttrType returns the representation of TemporaryCredentials in the Terraform plugin framework type
// system.
func (a TemporaryCredentials) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_temp_credentials": basetypes.ListType{
				ElemType: AwsCredentials{}.ToAttrType(ctx),
			},
			"azure_aad": basetypes.ListType{
				ElemType: AzureActiveDirectoryToken{}.ToAttrType(ctx),
			},
			"expiration_time": types.Int64Type,
		},
	}
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
type TriggeredUpdateStatus struct {
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

func (newState *TriggeredUpdateStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan TriggeredUpdateStatus) {
}

func (newState *TriggeredUpdateStatus) SyncEffectiveFieldsDuringRead(existingState TriggeredUpdateStatus) {
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

// ToAttrType returns the representation of TriggeredUpdateStatus in the Terraform plugin framework type
// system.
func (a TriggeredUpdateStatus) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
			"triggered_update_progress": basetypes.ListType{
				ElemType: PipelineProgress{}.ToAttrType(ctx),
			},
		},
	}
}

// Delete an assignment
type UnassignRequest struct {
	// Query for the ID of the metastore to delete.
	MetastoreId types.String `tfsdk:"-"`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UnassignRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnassignRequest) {
}

func (newState *UnassignRequest) SyncEffectiveFieldsDuringRead(existingState UnassignRequest) {
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

// ToAttrType returns the representation of UnassignRequest in the Terraform plugin framework type
// system.
func (a UnassignRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metastore_id": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

type UnassignResponse struct {
}

func (newState *UnassignResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnassignResponse) {
}

func (newState *UnassignResponse) SyncEffectiveFieldsDuringRead(existingState UnassignResponse) {
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

// ToAttrType returns the representation of UnassignResponse in the Terraform plugin framework type
// system.
func (a UnassignResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateAssignmentResponse struct {
}

func (newState *UpdateAssignmentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAssignmentResponse) {
}

func (newState *UpdateAssignmentResponse) SyncEffectiveFieldsDuringRead(existingState UpdateAssignmentResponse) {
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

// ToAttrType returns the representation of UpdateAssignmentResponse in the Terraform plugin framework type
// system.
func (a UpdateAssignmentResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateCatalog struct {
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

func (newState *UpdateCatalog) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCatalog) {
}

func (newState *UpdateCatalog) SyncEffectiveFieldsDuringRead(existingState UpdateCatalog) {
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
		"properties": reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of UpdateCatalog in the Terraform plugin framework type
// system.
func (a UpdateCatalog) ToAttrType(ctx context.Context) types.ObjectType {
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

type UpdateConnection struct {
	// Name of the connection.
	Name types.String `tfsdk:"-"`
	// New name for the connection.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// A map of key-value properties attached to the securable.
	Options types.Map `tfsdk:"options" tf:""`
	// Username of current owner of the connection.
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateConnection) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateConnection) {
}

func (newState *UpdateConnection) SyncEffectiveFieldsDuringRead(existingState UpdateConnection) {
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

// ToAttrType returns the representation of UpdateConnection in the Terraform plugin framework type
// system.
func (a UpdateConnection) ToAttrType(ctx context.Context) types.ObjectType {
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

type UpdateCredentialRequest struct {
	// The AWS IAM role configuration
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// Comment associated with the credential.
	Comment types.String `tfsdk:"comment" tf:"optional"`
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

func (newState *UpdateCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCredentialRequest) {
}

func (newState *UpdateCredentialRequest) SyncEffectiveFieldsDuringRead(existingState UpdateCredentialRequest) {
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
		"aws_iam_role":            reflect.TypeOf(AwsIamRole{}),
		"azure_managed_identity":  reflect.TypeOf(AzureManagedIdentity{}),
		"azure_service_principal": reflect.TypeOf(AzureServicePrincipal{}),
	}
}

// ToAttrType returns the representation of UpdateCredentialRequest in the Terraform plugin framework type
// system.
func (a UpdateCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRole{}.ToAttrType(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentity{}.ToAttrType(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal{}.ToAttrType(ctx),
			},
			"comment":         types.StringType,
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

type UpdateExternalLocation struct {
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

func (newState *UpdateExternalLocation) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExternalLocation) {
}

func (newState *UpdateExternalLocation) SyncEffectiveFieldsDuringRead(existingState UpdateExternalLocation) {
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
	}
}

// ToAttrType returns the representation of UpdateExternalLocation in the Terraform plugin framework type
// system.
func (a UpdateExternalLocation) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point":    types.StringType,
			"comment":         types.StringType,
			"credential_name": types.StringType,
			"encryption_details": basetypes.ListType{
				ElemType: EncryptionDetails{}.ToAttrType(ctx),
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

type UpdateFunction struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name types.String `tfsdk:"-"`
	// Username of current owner of function.
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateFunction) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateFunction) {
}

func (newState *UpdateFunction) SyncEffectiveFieldsDuringRead(existingState UpdateFunction) {
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

// ToAttrType returns the representation of UpdateFunction in the Terraform plugin framework type
// system.
func (a UpdateFunction) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *UpdateMetastore) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMetastore) {
}

func (newState *UpdateMetastore) SyncEffectiveFieldsDuringRead(existingState UpdateMetastore) {
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

// ToAttrType returns the representation of UpdateMetastore in the Terraform plugin framework type
// system.
func (a UpdateMetastore) ToAttrType(ctx context.Context) types.ObjectType {
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
	DefaultCatalogName types.String `tfsdk:"default_catalog_name" tf:"optional"`
	// The unique ID of the metastore.
	MetastoreId types.String `tfsdk:"metastore_id" tf:"optional"`
	// A workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UpdateMetastoreAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMetastoreAssignment) {
}

func (newState *UpdateMetastoreAssignment) SyncEffectiveFieldsDuringRead(existingState UpdateMetastoreAssignment) {
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

// ToAttrType returns the representation of UpdateMetastoreAssignment in the Terraform plugin framework type
// system.
func (a UpdateMetastoreAssignment) ToAttrType(ctx context.Context) types.ObjectType {
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
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The three-level (fully qualified) name of the model version
	FullName types.String `tfsdk:"-"`
	// The integer version number of the model version
	Version types.Int64 `tfsdk:"-"`
}

func (newState *UpdateModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelVersionRequest) {
}

func (newState *UpdateModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState UpdateModelVersionRequest) {
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

// ToAttrType returns the representation of UpdateModelVersionRequest in the Terraform plugin framework type
// system.
func (a UpdateModelVersionRequest) ToAttrType(ctx context.Context) types.ObjectType {
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
	Snapshot []MonitorSnapshot `tfsdk:"snapshot" tf:"optional,object"`
	// Full name of the table.
	TableName types.String `tfsdk:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries types.List `tfsdk:"time_series" tf:"optional,object"`
}

func (newState *UpdateMonitor) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateMonitor) {
}

func (newState *UpdateMonitor) SyncEffectiveFieldsDuringRead(existingState UpdateMonitor) {
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

// ToAttrType returns the representation of UpdateMonitor in the Terraform plugin framework type
// system.
func (a UpdateMonitor) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"baseline_table_name": types.StringType,
			"custom_metrics": basetypes.ListType{
				ElemType: MonitorMetric{}.ToAttrType(ctx),
			},
			"dashboard_id": types.StringType,
			"data_classification_config": basetypes.ListType{
				ElemType: MonitorDataClassificationConfig{}.ToAttrType(ctx),
			},
			"inference_log": basetypes.ListType{
				ElemType: MonitorInferenceLog{}.ToAttrType(ctx),
			},
			"notifications": basetypes.ListType{
				ElemType: MonitorNotifications{}.ToAttrType(ctx),
			},
			"output_schema_name": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: MonitorCronSchedule{}.ToAttrType(ctx),
			},
			"slicing_exprs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"snapshot": basetypes.ListType{
				ElemType: MonitorSnapshot{}.ToAttrType(ctx),
			},
			"table_name": types.StringType,
			"time_series": basetypes.ListType{
				ElemType: MonitorTimeSeries{}.ToAttrType(ctx),
			},
		},
	}
}

type UpdatePermissions struct {
	// Array of permissions change objects.
	Changes types.List `tfsdk:"changes" tf:"optional"`
	// Full name of securable.
	FullName types.String `tfsdk:"-"`
	// Type of securable.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *UpdatePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePermissions) {
}

func (newState *UpdatePermissions) SyncEffectiveFieldsDuringRead(existingState UpdatePermissions) {
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

// ToAttrType returns the representation of UpdatePermissions in the Terraform plugin framework type
// system.
func (a UpdatePermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"changes": basetypes.ListType{
				ElemType: PermissionsChange{}.ToAttrType(ctx),
			},
			"full_name":      types.StringType,
			"securable_type": types.StringType,
		},
	}
}

type UpdateRegisteredModelRequest struct {
	// The comment attached to the registered model
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The three-level (fully qualified) name of the registered model
	FullName types.String `tfsdk:"-"`
	// New name for the registered model.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// The identifier of the user who owns the registered model
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateRegisteredModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRegisteredModelRequest) {
}

func (newState *UpdateRegisteredModelRequest) SyncEffectiveFieldsDuringRead(existingState UpdateRegisteredModelRequest) {
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

// ToAttrType returns the representation of UpdateRegisteredModelRequest in the Terraform plugin framework type
// system.
func (a UpdateRegisteredModelRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

// ToAttrType returns the representation of UpdateResponse in the Terraform plugin framework type
// system.
func (a UpdateResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateSchema struct {
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

func (newState *UpdateSchema) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSchema) {
}

func (newState *UpdateSchema) SyncEffectiveFieldsDuringRead(existingState UpdateSchema) {
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

// ToAttrType returns the representation of UpdateSchema in the Terraform plugin framework type
// system.
func (a UpdateSchema) ToAttrType(ctx context.Context) types.ObjectType {
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

type UpdateStorageCredential struct {
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
	DatabricksGcpServiceAccount []DatabricksGcpServiceAccountRequest `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
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

func (newState *UpdateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateStorageCredential) {
}

func (newState *UpdateStorageCredential) SyncEffectiveFieldsDuringRead(existingState UpdateStorageCredential) {
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

// ToAttrType returns the representation of UpdateStorageCredential in the Terraform plugin framework type
// system.
func (a UpdateStorageCredential) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRoleRequest{}.ToAttrType(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentityResponse{}.ToAttrType(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal{}.ToAttrType(ctx),
			},
			"cloudflare_api_token": basetypes.ListType{
				ElemType: CloudflareApiToken{}.ToAttrType(ctx),
			},
			"comment": types.StringType,
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccountRequest{}.ToAttrType(ctx),
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

// Update a table owner.
type UpdateTableRequest struct {
	// Full name of the table.
	FullName types.String `tfsdk:"-"`

	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateTableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateTableRequest) {
}

func (newState *UpdateTableRequest) SyncEffectiveFieldsDuringRead(existingState UpdateTableRequest) {
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

// ToAttrType returns the representation of UpdateTableRequest in the Terraform plugin framework type
// system.
func (a UpdateTableRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
			"owner":     types.StringType,
		},
	}
}

type UpdateVolumeRequestContent struct {
	// The comment attached to the volume
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// The three-level (fully qualified) name of the volume
	Name types.String `tfsdk:"-"`
	// New name for the volume.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
	// The identifier of the user who owns the volume
	Owner types.String `tfsdk:"owner" tf:"optional"`
}

func (newState *UpdateVolumeRequestContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateVolumeRequestContent) {
}

func (newState *UpdateVolumeRequestContent) SyncEffectiveFieldsDuringRead(existingState UpdateVolumeRequestContent) {
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

// ToAttrType returns the representation of UpdateVolumeRequestContent in the Terraform plugin framework type
// system.
func (a UpdateVolumeRequestContent) ToAttrType(ctx context.Context) types.ObjectType {
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
	AssignWorkspaces types.List `tfsdk:"assign_workspaces" tf:"optional"`
	// The name of the catalog.
	Name types.String `tfsdk:"-"`
	// A list of workspace IDs.
	UnassignWorkspaces types.List `tfsdk:"unassign_workspaces" tf:"optional"`
}

func (newState *UpdateWorkspaceBindings) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceBindings) {
}

func (newState *UpdateWorkspaceBindings) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceBindings) {
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

// ToAttrType returns the representation of UpdateWorkspaceBindings in the Terraform plugin framework type
// system.
func (a UpdateWorkspaceBindings) ToAttrType(ctx context.Context) types.ObjectType {
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

type UpdateWorkspaceBindingsParameters struct {
	// List of workspace bindings
	Add types.List `tfsdk:"add" tf:"optional"`
	// List of workspace bindings
	Remove types.List `tfsdk:"remove" tf:"optional"`
	// The name of the securable.
	SecurableName types.String `tfsdk:"-"`
	// The type of the securable to bind to a workspace.
	SecurableType types.String `tfsdk:"-"`
}

func (newState *UpdateWorkspaceBindingsParameters) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceBindingsParameters) {
}

func (newState *UpdateWorkspaceBindingsParameters) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceBindingsParameters) {
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

// ToAttrType returns the representation of UpdateWorkspaceBindingsParameters in the Terraform plugin framework type
// system.
func (a UpdateWorkspaceBindingsParameters) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"add": basetypes.ListType{
				ElemType: WorkspaceBinding{}.ToAttrType(ctx),
			},
			"remove": basetypes.ListType{
				ElemType: WorkspaceBinding{}.ToAttrType(ctx),
			},
			"securable_name": types.StringType,
			"securable_type": types.StringType,
		},
	}
}

type ValidateCredentialRequest struct {
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

func (newState *ValidateCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateCredentialRequest) {
}

func (newState *ValidateCredentialRequest) SyncEffectiveFieldsDuringRead(existingState ValidateCredentialRequest) {
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
		"aws_iam_role":           reflect.TypeOf(AwsIamRole{}),
		"azure_managed_identity": reflect.TypeOf(AzureManagedIdentity{}),
	}
}

// ToAttrType returns the representation of ValidateCredentialRequest in the Terraform plugin framework type
// system.
func (a ValidateCredentialRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRole{}.ToAttrType(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentity{}.ToAttrType(ctx),
			},
			"credential_name":        types.StringType,
			"external_location_name": types.StringType,
			"purpose":                types.StringType,
			"read_only":              types.BoolType,
			"url":                    types.StringType,
		},
	}
}

type ValidateCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage. Only
	// applicable for when purpose is **STORAGE**.
	IsDir types.Bool `tfsdk:"isDir" tf:"optional"`
	// The results of the validation check.
	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ValidateCredentialResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateCredentialResponse) {
}

func (newState *ValidateCredentialResponse) SyncEffectiveFieldsDuringRead(existingState ValidateCredentialResponse) {
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

// ToAttrType returns the representation of ValidateCredentialResponse in the Terraform plugin framework type
// system.
func (a ValidateCredentialResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"isDir": types.BoolType,
			"results": basetypes.ListType{
				ElemType: CredentialValidationResult{}.ToAttrType(ctx),
			},
		},
	}
}

type ValidateStorageCredential struct {
	// The AWS IAM role configuration.
	AwsIamRole types.List `tfsdk:"aws_iam_role" tf:"optional,object"`
	// The Azure managed identity configuration.
	AzureManagedIdentity types.List `tfsdk:"azure_managed_identity" tf:"optional,object"`
	// The Azure service principal configuration.
	AzureServicePrincipal types.List `tfsdk:"azure_service_principal" tf:"optional,object"`
	// The Cloudflare API token configuration.
	CloudflareApiToken types.List `tfsdk:"cloudflare_api_token" tf:"optional,object"`
	// The Databricks created GCP service account configuration.
	DatabricksGcpServiceAccount []DatabricksGcpServiceAccountRequest `tfsdk:"databricks_gcp_service_account" tf:"optional,object"`
	// The name of an existing external location to validate.
	ExternalLocationName types.String `tfsdk:"external_location_name" tf:"optional"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly types.Bool `tfsdk:"read_only" tf:"optional"`
	// The name of the storage credential to validate.
	StorageCredentialName types.String `tfsdk:"storage_credential_name" tf:"optional"`
	// The external location url to validate.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *ValidateStorageCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateStorageCredential) {
}

func (newState *ValidateStorageCredential) SyncEffectiveFieldsDuringRead(existingState ValidateStorageCredential) {
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

// ToAttrType returns the representation of ValidateStorageCredential in the Terraform plugin framework type
// system.
func (a ValidateStorageCredential) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_iam_role": basetypes.ListType{
				ElemType: AwsIamRoleRequest{}.ToAttrType(ctx),
			},
			"azure_managed_identity": basetypes.ListType{
				ElemType: AzureManagedIdentityRequest{}.ToAttrType(ctx),
			},
			"azure_service_principal": basetypes.ListType{
				ElemType: AzureServicePrincipal{}.ToAttrType(ctx),
			},
			"cloudflare_api_token": basetypes.ListType{
				ElemType: CloudflareApiToken{}.ToAttrType(ctx),
			},
			"databricks_gcp_service_account": basetypes.ListType{
				ElemType: DatabricksGcpServiceAccountRequest{}.ToAttrType(ctx),
			},
			"external_location_name":  types.StringType,
			"read_only":               types.BoolType,
			"storage_credential_name": types.StringType,
			"url":                     types.StringType,
		},
	}
}

type ValidateStorageCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage.
	IsDir types.Bool `tfsdk:"isDir" tf:"optional"`
	// The results of the validation check.
	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ValidateStorageCredentialResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidateStorageCredentialResponse) {
}

func (newState *ValidateStorageCredentialResponse) SyncEffectiveFieldsDuringRead(existingState ValidateStorageCredentialResponse) {
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

// ToAttrType returns the representation of ValidateStorageCredentialResponse in the Terraform plugin framework type
// system.
func (a ValidateStorageCredentialResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"isDir": types.BoolType,
			"results": basetypes.ListType{
				ElemType: ValidationResult{}.ToAttrType(ctx),
			},
		},
	}
}

type ValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message types.String `tfsdk:"message" tf:"optional"`
	// The operation tested.
	Operation types.String `tfsdk:"operation" tf:"optional"`
	// The results of the tested operation.
	Result types.String `tfsdk:"result" tf:"optional"`
}

func (newState *ValidationResult) SyncEffectiveFieldsDuringCreateOrUpdate(plan ValidationResult) {
}

func (newState *ValidationResult) SyncEffectiveFieldsDuringRead(existingState ValidationResult) {
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

// ToAttrType returns the representation of ValidationResult in the Terraform plugin framework type
// system.
func (a ValidationResult) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *VolumeInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan VolumeInfo) {
}

func (newState *VolumeInfo) SyncEffectiveFieldsDuringRead(existingState VolumeInfo) {
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

// ToAttrType returns the representation of VolumeInfo in the Terraform plugin framework type
// system.
func (a VolumeInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_point": types.StringType,
			"browse_only":  types.BoolType,
			"catalog_name": types.StringType,
			"comment":      types.StringType,
			"created_at":   types.Int64Type,
			"created_by":   types.StringType,
			"encryption_details": basetypes.ListType{
				ElemType: EncryptionDetails{}.ToAttrType(ctx),
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

type WorkspaceBinding struct {
	BindingType types.String `tfsdk:"binding_type" tf:"optional"`

	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:"optional"`
}

func (newState *WorkspaceBinding) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceBinding) {
}

func (newState *WorkspaceBinding) SyncEffectiveFieldsDuringRead(existingState WorkspaceBinding) {
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

// ToAttrType returns the representation of WorkspaceBinding in the Terraform plugin framework type
// system.
func (a WorkspaceBinding) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"binding_type": types.StringType,
			"workspace_id": types.Int64Type,
		},
	}
}

// Currently assigned workspace bindings
type WorkspaceBindingsResponse struct {
	// List of workspace bindings
	Bindings types.List `tfsdk:"bindings" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *WorkspaceBindingsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceBindingsResponse) {
}

func (newState *WorkspaceBindingsResponse) SyncEffectiveFieldsDuringRead(existingState WorkspaceBindingsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceBindingsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceBindingsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"bindings": reflect.TypeOf(WorkspaceBinding{}),
	}
}

// ToAttrType returns the representation of WorkspaceBindingsResponse in the Terraform plugin framework type
// system.
func (a WorkspaceBindingsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bindings": basetypes.ListType{
				ElemType: WorkspaceBinding{}.ToAttrType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}
