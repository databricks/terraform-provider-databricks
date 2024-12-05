// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package workspace_tf

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AclItem struct {
	// The permission level applied to the principal.
	Permission types.String `tfsdk:"permission" tf:""`
	// The principal in which the permission is applied.
	Principal types.String `tfsdk:"principal" tf:""`
}

func (newState *AclItem) SyncEffectiveFieldsDuringCreateOrUpdate(plan AclItem) {
}

func (newState *AclItem) SyncEffectiveFieldsDuringRead(existingState AclItem) {
}

func (a AclItem) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a AclItem) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
			"principal":  types.StringType,
		},
	}
}

type AzureKeyVaultSecretScopeMetadata struct {
	// The DNS of the KeyVault
	DnsName types.String `tfsdk:"dns_name" tf:""`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	ResourceId types.String `tfsdk:"resource_id" tf:""`
}

func (newState *AzureKeyVaultSecretScopeMetadata) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureKeyVaultSecretScopeMetadata) {
}

func (newState *AzureKeyVaultSecretScopeMetadata) SyncEffectiveFieldsDuringRead(existingState AzureKeyVaultSecretScopeMetadata) {
}

func (a AzureKeyVaultSecretScopeMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a AzureKeyVaultSecretScopeMetadata) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dns_name":    types.StringType,
			"resource_id": types.StringType,
		},
	}
}

type CreateCredentialsRequest struct {
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	GitProvider types.String `tfsdk:"git_provider" tf:""`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername types.String `tfsdk:"git_username" tf:"optional"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken types.String `tfsdk:"personal_access_token" tf:"optional"`
}

func (newState *CreateCredentialsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialsRequest) {
}

func (newState *CreateCredentialsRequest) SyncEffectiveFieldsDuringRead(existingState CreateCredentialsRequest) {
}

func (a CreateCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateCredentialsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"git_provider":          types.StringType,
			"git_username":          types.StringType,
			"personal_access_token": types.StringType,
		},
	}
}

type CreateCredentialsResponse struct {
	// ID of the credential object in the workspace.
	CredentialId types.Int64 `tfsdk:"credential_id" tf:""`
	// The Git provider associated with the credential.
	GitProvider types.String `tfsdk:"git_provider" tf:""`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername types.String `tfsdk:"git_username" tf:"optional"`
}

func (newState *CreateCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialsResponse) {
}

func (newState *CreateCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState CreateCredentialsResponse) {
}

func (a CreateCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateCredentialsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.Int64Type,
			"git_provider":  types.StringType,
			"git_username":  types.StringType,
		},
	}
}

type CreateRepoRequest struct {
	// Desired path for the repo in the workspace. Almost any path in the
	// workspace can be chosen. If repo is created in `/Repos`, path must be in
	// the format `/Repos/{folder}/{repo-name}`.
	Path types.String `tfsdk:"path" tf:"optional"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	Provider types.String `tfsdk:"provider" tf:""`
	// If specified, the repo will be created with sparse checkout enabled. You
	// cannot enable/disable sparse checkout after the repo is created.
	SparseCheckout types.List `tfsdk:"sparse_checkout" tf:"optional,object"`
	// URL of the Git repository to be linked.
	Url types.String `tfsdk:"url" tf:""`
}

func (newState *CreateRepoRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRepoRequest) {
}

func (newState *CreateRepoRequest) SyncEffectiveFieldsDuringRead(existingState CreateRepoRequest) {
}

func (a CreateRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

func (a CreateRepoRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":     types.StringType,
			"provider": types.StringType,
			"sparse_checkout": basetypes.ListType{
				ElemType: SparseCheckout{}.ToAttrType(ctx),
			},
			"url": types.StringType,
		},
	}
}

type CreateRepoResponse struct {
	// Branch that the Git folder (repo) is checked out to.
	Branch types.String `tfsdk:"branch" tf:"optional"`
	// SHA-1 hash representing the commit ID of the current HEAD of the Git
	// folder (repo).
	HeadCommitId types.String `tfsdk:"head_commit_id" tf:"optional"`
	// ID of the Git folder (repo) object in the workspace.
	Id types.Int64 `tfsdk:"id" tf:"optional"`
	// Path of the Git folder (repo) in the workspace.
	Path types.String `tfsdk:"path" tf:"optional"`
	// Git provider of the linked Git repository.
	Provider types.String `tfsdk:"provider" tf:"optional"`
	// Sparse checkout settings for the Git folder (repo).
	SparseCheckout types.List `tfsdk:"sparse_checkout" tf:"optional,object"`
	// URL of the linked Git repository.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *CreateRepoResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRepoResponse) {
}

func (newState *CreateRepoResponse) SyncEffectiveFieldsDuringRead(existingState CreateRepoResponse) {
}

func (a CreateRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

func (a CreateRepoResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":         types.StringType,
			"head_commit_id": types.StringType,
			"id":             types.Int64Type,
			"path":           types.StringType,
			"provider":       types.StringType,
			"sparse_checkout": basetypes.ListType{
				ElemType: SparseCheckout{}.ToAttrType(ctx),
			},
			"url": types.StringType,
		},
	}
}

type CreateScope struct {
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	BackendAzureKeyvault types.List `tfsdk:"backend_azure_keyvault" tf:"optional,object"`
	// The principal that is initially granted `MANAGE` permission to the
	// created scope.
	InitialManagePrincipal types.String `tfsdk:"initial_manage_principal" tf:"optional"`
	// Scope name requested by the user. Scope names are unique.
	Scope types.String `tfsdk:"scope" tf:""`
	// The backend type the scope will be created with. If not specified, will
	// default to `DATABRICKS`
	ScopeBackendType types.String `tfsdk:"scope_backend_type" tf:"optional"`
}

func (newState *CreateScope) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateScope) {
}

func (newState *CreateScope) SyncEffectiveFieldsDuringRead(existingState CreateScope) {
}

func (a CreateScope) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"backend_azure_keyvault": reflect.TypeOf(AzureKeyVaultSecretScopeMetadata{}),
	}
}

func (a CreateScope) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"backend_azure_keyvault": basetypes.ListType{
				ElemType: AzureKeyVaultSecretScopeMetadata{}.ToAttrType(ctx),
			},
			"initial_manage_principal": types.StringType,
			"scope":                    types.StringType,
			"scope_backend_type":       types.StringType,
		},
	}
}

type CreateScopeResponse struct {
}

func (newState *CreateScopeResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateScopeResponse) {
}

func (newState *CreateScopeResponse) SyncEffectiveFieldsDuringRead(existingState CreateScopeResponse) {
}

func (a CreateScopeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CreateScopeResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	CredentialId types.Int64 `tfsdk:"credential_id" tf:""`
	// The Git provider associated with the credential.
	GitProvider types.String `tfsdk:"git_provider" tf:"optional"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername types.String `tfsdk:"git_username" tf:"optional"`
}

func (newState *CredentialInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CredentialInfo) {
}

func (newState *CredentialInfo) SyncEffectiveFieldsDuringRead(existingState CredentialInfo) {
}

func (a CredentialInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a CredentialInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.Int64Type,
			"git_provider":  types.StringType,
			"git_username":  types.StringType,
		},
	}
}

type Delete struct {
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"path" tf:""`
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive types.Bool `tfsdk:"recursive" tf:"optional"`
}

func (newState *Delete) SyncEffectiveFieldsDuringCreateOrUpdate(plan Delete) {
}

func (newState *Delete) SyncEffectiveFieldsDuringRead(existingState Delete) {
}

func (a Delete) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Delete) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":      types.StringType,
			"recursive": types.BoolType,
		},
	}
}

type DeleteAcl struct {
	// The principal to remove an existing ACL from.
	Principal types.String `tfsdk:"principal" tf:""`
	// The name of the scope to remove permissions from.
	Scope types.String `tfsdk:"scope" tf:""`
}

func (newState *DeleteAcl) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAcl) {
}

func (newState *DeleteAcl) SyncEffectiveFieldsDuringRead(existingState DeleteAcl) {
}

func (a DeleteAcl) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteAcl) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"scope":     types.StringType,
		},
	}
}

type DeleteAclResponse struct {
}

func (newState *DeleteAclResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAclResponse) {
}

func (newState *DeleteAclResponse) SyncEffectiveFieldsDuringRead(existingState DeleteAclResponse) {
}

func (a DeleteAclResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteAclResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a credential
type DeleteCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-"`
}

func (newState *DeleteCredentialsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCredentialsRequest) {
}

func (newState *DeleteCredentialsRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCredentialsRequest) {
}

func (a DeleteCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteCredentialsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.Int64Type,
		},
	}
}

type DeleteCredentialsResponse struct {
}

func (newState *DeleteCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCredentialsResponse) {
}

func (newState *DeleteCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState DeleteCredentialsResponse) {
}

func (a DeleteCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteCredentialsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a repo
type DeleteRepoRequest struct {
	// The ID for the corresponding repo to delete.
	RepoId types.Int64 `tfsdk:"-"`
}

func (newState *DeleteRepoRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRepoRequest) {
}

func (newState *DeleteRepoRequest) SyncEffectiveFieldsDuringRead(existingState DeleteRepoRequest) {
}

func (a DeleteRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteRepoRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.Int64Type,
		},
	}
}

type DeleteRepoResponse struct {
}

func (newState *DeleteRepoResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRepoResponse) {
}

func (newState *DeleteRepoResponse) SyncEffectiveFieldsDuringRead(existingState DeleteRepoResponse) {
}

func (a DeleteRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteRepoResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

func (a DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteScope struct {
	// Name of the scope to delete.
	Scope types.String `tfsdk:"scope" tf:""`
}

func (newState *DeleteScope) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteScope) {
}

func (newState *DeleteScope) SyncEffectiveFieldsDuringRead(existingState DeleteScope) {
}

func (a DeleteScope) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteScope) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope": types.StringType,
		},
	}
}

type DeleteScopeResponse struct {
}

func (newState *DeleteScopeResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteScopeResponse) {
}

func (newState *DeleteScopeResponse) SyncEffectiveFieldsDuringRead(existingState DeleteScopeResponse) {
}

func (a DeleteScopeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteScopeResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteSecret struct {
	// Name of the secret to delete.
	Key types.String `tfsdk:"key" tf:""`
	// The name of the scope that contains the secret to delete.
	Scope types.String `tfsdk:"scope" tf:""`
}

func (newState *DeleteSecret) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSecret) {
}

func (newState *DeleteSecret) SyncEffectiveFieldsDuringRead(existingState DeleteSecret) {
}

func (a DeleteSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteSecret) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"scope": types.StringType,
		},
	}
}

type DeleteSecretResponse struct {
}

func (newState *DeleteSecretResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSecretResponse) {
}

func (newState *DeleteSecretResponse) SyncEffectiveFieldsDuringRead(existingState DeleteSecretResponse) {
}

func (a DeleteSecretResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteSecretResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Export a workspace object
type ExportRequest struct {
	// This specifies the format of the exported file. By default, this is
	// `SOURCE`.
	//
	// The value is case sensitive.
	//
	// - `SOURCE`: The notebook is exported as source code. Directory exports
	// will not include non-notebook entries. - `HTML`: The notebook is exported
	// as an HTML file. - `JUPYTER`: The notebook is exported as a
	// Jupyter/IPython Notebook file. - `DBC`: The notebook is exported in
	// Databricks archive format. Directory exports will not include
	// non-notebook entries. - `R_MARKDOWN`: The notebook is exported to R
	// Markdown format. - `AUTO`: The object or directory is exported depending
	// on the objects type. Directory exports will include notebooks and
	// workspace files.
	Format types.String `tfsdk:"-"`
	// The absolute path of the object or directory. Exporting a directory is
	// only supported for the `DBC`, `SOURCE`, and `AUTO` format.
	Path types.String `tfsdk:"-"`
}

func (newState *ExportRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportRequest) {
}

func (newState *ExportRequest) SyncEffectiveFieldsDuringRead(existingState ExportRequest) {
}

func (a ExportRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ExportRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"format": types.StringType,
			"path":   types.StringType,
		},
	}
}

type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	Content types.String `tfsdk:"content" tf:"optional"`
	// The file type of the exported file.
	FileType types.String `tfsdk:"file_type" tf:"optional"`
}

func (newState *ExportResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportResponse) {
}

func (newState *ExportResponse) SyncEffectiveFieldsDuringRead(existingState ExportResponse) {
}

func (a ExportResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ExportResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content":   types.StringType,
			"file_type": types.StringType,
		},
	}
}

// Get secret ACL details
type GetAclRequest struct {
	// The principal to fetch ACL information for.
	Principal types.String `tfsdk:"-"`
	// The name of the scope to fetch ACL information from.
	Scope types.String `tfsdk:"-"`
}

func (newState *GetAclRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAclRequest) {
}

func (newState *GetAclRequest) SyncEffectiveFieldsDuringRead(existingState GetAclRequest) {
}

func (a GetAclRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetAclRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"scope":     types.StringType,
		},
	}
}

// Get a credential entry
type GetCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-"`
}

func (newState *GetCredentialsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCredentialsRequest) {
}

func (newState *GetCredentialsRequest) SyncEffectiveFieldsDuringRead(existingState GetCredentialsRequest) {
}

func (a GetCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetCredentialsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.Int64Type,
		},
	}
}

type GetCredentialsResponse struct {
	// ID of the credential object in the workspace.
	CredentialId types.Int64 `tfsdk:"credential_id" tf:""`
	// The Git provider associated with the credential.
	GitProvider types.String `tfsdk:"git_provider" tf:"optional"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername types.String `tfsdk:"git_username" tf:"optional"`
}

func (newState *GetCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCredentialsResponse) {
}

func (newState *GetCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState GetCredentialsResponse) {
}

func (a GetCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetCredentialsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.Int64Type,
			"git_provider":  types.StringType,
			"git_username":  types.StringType,
		},
	}
}

// Get repo permission levels
type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

func (newState *GetRepoPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRepoPermissionLevelsRequest) {
}

func (newState *GetRepoPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetRepoPermissionLevelsRequest) {
}

func (a GetRepoPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetRepoPermissionLevelsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.StringType,
		},
	}
}

type GetRepoPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetRepoPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRepoPermissionLevelsResponse) {
}

func (newState *GetRepoPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetRepoPermissionLevelsResponse) {
}

func (a GetRepoPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(RepoPermissionsDescription{}),
	}
}

func (a GetRepoPermissionLevelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: RepoPermissionsDescription{}.ToAttrType(ctx),
			},
		},
	}
}

// Get repo permissions
type GetRepoPermissionsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

func (newState *GetRepoPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRepoPermissionsRequest) {
}

func (newState *GetRepoPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetRepoPermissionsRequest) {
}

func (a GetRepoPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetRepoPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.StringType,
		},
	}
}

// Get a repo
type GetRepoRequest struct {
	// ID of the Git folder (repo) object in the workspace.
	RepoId types.Int64 `tfsdk:"-"`
}

func (newState *GetRepoRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRepoRequest) {
}

func (newState *GetRepoRequest) SyncEffectiveFieldsDuringRead(existingState GetRepoRequest) {
}

func (a GetRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetRepoRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.Int64Type,
		},
	}
}

type GetRepoResponse struct {
	// Branch that the local version of the repo is checked out to.
	Branch types.String `tfsdk:"branch" tf:"optional"`
	// SHA-1 hash representing the commit ID of the current HEAD of the repo.
	HeadCommitId types.String `tfsdk:"head_commit_id" tf:"optional"`
	// ID of the Git folder (repo) object in the workspace.
	Id types.Int64 `tfsdk:"id" tf:"optional"`
	// Path of the Git folder (repo) in the workspace.
	Path types.String `tfsdk:"path" tf:"optional"`
	// Git provider of the linked Git repository.
	Provider types.String `tfsdk:"provider" tf:"optional"`
	// Sparse checkout settings for the Git folder (repo).
	SparseCheckout types.List `tfsdk:"sparse_checkout" tf:"optional,object"`
	// URL of the linked Git repository.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *GetRepoResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRepoResponse) {
}

func (newState *GetRepoResponse) SyncEffectiveFieldsDuringRead(existingState GetRepoResponse) {
}

func (a GetRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

func (a GetRepoResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":         types.StringType,
			"head_commit_id": types.StringType,
			"id":             types.Int64Type,
			"path":           types.StringType,
			"provider":       types.StringType,
			"sparse_checkout": basetypes.ListType{
				ElemType: SparseCheckout{}.ToAttrType(ctx),
			},
			"url": types.StringType,
		},
	}
}

// Get a secret
type GetSecretRequest struct {
	// The key to fetch secret for.
	Key types.String `tfsdk:"-"`
	// The name of the scope to fetch secret information from.
	Scope types.String `tfsdk:"-"`
}

func (newState *GetSecretRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSecretRequest) {
}

func (newState *GetSecretRequest) SyncEffectiveFieldsDuringRead(existingState GetSecretRequest) {
}

func (a GetSecretRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetSecretRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"scope": types.StringType,
		},
	}
}

type GetSecretResponse struct {
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The value of the secret in its byte representation.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *GetSecretResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSecretResponse) {
}

func (newState *GetSecretResponse) SyncEffectiveFieldsDuringRead(existingState GetSecretResponse) {
}

func (a GetSecretResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetSecretResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Get status
type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"-"`
}

func (newState *GetStatusRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStatusRequest) {
}

func (newState *GetStatusRequest) SyncEffectiveFieldsDuringRead(existingState GetStatusRequest) {
}

func (a GetStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetStatusRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

// Get workspace object permission levels
type GetWorkspaceObjectPermissionLevelsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-"`
}

func (newState *GetWorkspaceObjectPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceObjectPermissionLevelsRequest) {
}

func (newState *GetWorkspaceObjectPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceObjectPermissionLevelsRequest) {
}

func (a GetWorkspaceObjectPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetWorkspaceObjectPermissionLevelsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_object_id":   types.StringType,
			"workspace_object_type": types.StringType,
		},
	}
}

type GetWorkspaceObjectPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetWorkspaceObjectPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceObjectPermissionLevelsResponse) {
}

func (newState *GetWorkspaceObjectPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceObjectPermissionLevelsResponse) {
}

func (a GetWorkspaceObjectPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(WorkspaceObjectPermissionsDescription{}),
	}
}

func (a GetWorkspaceObjectPermissionLevelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: WorkspaceObjectPermissionsDescription{}.ToAttrType(ctx),
			},
		},
	}
}

// Get workspace object permissions
type GetWorkspaceObjectPermissionsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-"`
}

func (newState *GetWorkspaceObjectPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceObjectPermissionsRequest) {
}

func (newState *GetWorkspaceObjectPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceObjectPermissionsRequest) {
}

func (a GetWorkspaceObjectPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetWorkspaceObjectPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_object_id":   types.StringType,
			"workspace_object_type": types.StringType,
		},
	}
}

type Import struct {
	// The base64-encoded content. This has a limit of 10 MB.
	//
	// If the limit (10MB) is exceeded, exception with error code
	// **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown. This parameter might be absent,
	// and instead a posted file is used.
	Content types.String `tfsdk:"content" tf:"optional"`
	// This specifies the format of the file to be imported.
	//
	// The value is case sensitive.
	//
	// - `AUTO`: The item is imported depending on an analysis of the item's
	// extension and the header content provided in the request. If the item is
	// imported as a notebook, then the item's extension is automatically
	// removed. - `SOURCE`: The notebook or directory is imported as source
	// code. - `HTML`: The notebook is imported as an HTML file. - `JUPYTER`:
	// The notebook is imported as a Jupyter/IPython Notebook file. - `DBC`: The
	// notebook is imported in Databricks archive format. Required for
	// directories. - `R_MARKDOWN`: The notebook is imported from R Markdown
	// format.
	Format types.String `tfsdk:"format" tf:"optional"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language types.String `tfsdk:"language" tf:"optional"`
	// The flag that specifies whether to overwrite existing object. It is
	// `false` by default. For `DBC` format, `overwrite` is not supported since
	// it may contain a directory.
	Overwrite types.Bool `tfsdk:"overwrite" tf:"optional"`
	// The absolute path of the object or directory. Importing a directory is
	// only supported for the `DBC` and `SOURCE` formats.
	Path types.String `tfsdk:"path" tf:""`
}

func (newState *Import) SyncEffectiveFieldsDuringCreateOrUpdate(plan Import) {
}

func (newState *Import) SyncEffectiveFieldsDuringRead(existingState Import) {
}

func (a Import) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Import) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content":   types.StringType,
			"format":    types.StringType,
			"language":  types.StringType,
			"overwrite": types.BoolType,
			"path":      types.StringType,
		},
	}
}

type ImportResponse struct {
}

func (newState *ImportResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ImportResponse) {
}

func (newState *ImportResponse) SyncEffectiveFieldsDuringRead(existingState ImportResponse) {
}

func (a ImportResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ImportResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Lists ACLs
type ListAclsRequest struct {
	// The name of the scope to fetch ACL information from.
	Scope types.String `tfsdk:"-"`
}

func (newState *ListAclsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAclsRequest) {
}

func (newState *ListAclsRequest) SyncEffectiveFieldsDuringRead(existingState ListAclsRequest) {
}

func (a ListAclsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListAclsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope": types.StringType,
		},
	}
}

type ListAclsResponse struct {
	// The associated ACLs rule applied to principals in the given scope.
	Items types.List `tfsdk:"items" tf:"optional"`
}

func (newState *ListAclsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAclsResponse) {
}

func (newState *ListAclsResponse) SyncEffectiveFieldsDuringRead(existingState ListAclsResponse) {
}

func (a ListAclsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(AclItem{}),
	}
}

func (a ListAclsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: AclItem{}.ToAttrType(ctx),
			},
		},
	}
}

type ListCredentialsResponse struct {
	// List of credentials.
	Credentials types.List `tfsdk:"credentials" tf:"optional"`
}

func (newState *ListCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCredentialsResponse) {
}

func (newState *ListCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState ListCredentialsResponse) {
}

func (a ListCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credentials": reflect.TypeOf(CredentialInfo{}),
	}
}

func (a ListCredentialsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials": basetypes.ListType{
				ElemType: CredentialInfo{}.ToAttrType(ctx),
			},
		},
	}
}

// Get repos
type ListReposRequest struct {
	// Token used to get the next page of results. If not specified, returns the
	// first page of results as well as a next page token if there are more
	// results.
	NextPageToken types.String `tfsdk:"-"`
	// Filters repos that have paths starting with the given path prefix. If not
	// provided or when provided an effectively empty prefix (`/` or
	// `/Workspace`) Git folders (repos) from `/Workspace/Repos` will be served.
	PathPrefix types.String `tfsdk:"-"`
}

func (newState *ListReposRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListReposRequest) {
}

func (newState *ListReposRequest) SyncEffectiveFieldsDuringRead(existingState ListReposRequest) {
}

func (a ListReposRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListReposRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"path_prefix":     types.StringType,
		},
	}
}

type ListReposResponse struct {
	// Token that can be specified as a query parameter to the `GET /repos`
	// endpoint to retrieve the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// List of Git folders (repos).
	Repos types.List `tfsdk:"repos" tf:"optional"`
}

func (newState *ListReposResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListReposResponse) {
}

func (newState *ListReposResponse) SyncEffectiveFieldsDuringRead(existingState ListReposResponse) {
}

func (a ListReposResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"repos": reflect.TypeOf(RepoInfo{}),
	}
}

func (a ListReposResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"repos": basetypes.ListType{
				ElemType: RepoInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type ListResponse struct {
	// List of objects.
	Objects types.List `tfsdk:"objects" tf:"optional"`
}

func (newState *ListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListResponse) {
}

func (newState *ListResponse) SyncEffectiveFieldsDuringRead(existingState ListResponse) {
}

func (a ListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objects": reflect.TypeOf(ObjectInfo{}),
	}
}

func (a ListResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"objects": basetypes.ListType{
				ElemType: ObjectInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type ListScopesResponse struct {
	// The available secret scopes.
	Scopes types.List `tfsdk:"scopes" tf:"optional"`
}

func (newState *ListScopesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListScopesResponse) {
}

func (newState *ListScopesResponse) SyncEffectiveFieldsDuringRead(existingState ListScopesResponse) {
}

func (a ListScopesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(SecretScope{}),
	}
}

func (a ListScopesResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scopes": basetypes.ListType{
				ElemType: SecretScope{}.ToAttrType(ctx),
			},
		},
	}
}

// List secret keys
type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	Scope types.String `tfsdk:"-"`
}

func (newState *ListSecretsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSecretsRequest) {
}

func (newState *ListSecretsRequest) SyncEffectiveFieldsDuringRead(existingState ListSecretsRequest) {
}

func (a ListSecretsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListSecretsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope": types.StringType,
		},
	}
}

type ListSecretsResponse struct {
	// Metadata information of all secrets contained within the given scope.
	Secrets types.List `tfsdk:"secrets" tf:"optional"`
}

func (newState *ListSecretsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSecretsResponse) {
}

func (newState *ListSecretsResponse) SyncEffectiveFieldsDuringRead(existingState ListSecretsResponse) {
}

func (a ListSecretsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"secrets": reflect.TypeOf(SecretMetadata{}),
	}
}

func (a ListSecretsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"secrets": basetypes.ListType{
				ElemType: SecretMetadata{}.ToAttrType(ctx),
			},
		},
	}
}

// List contents
type ListWorkspaceRequest struct {
	// UTC timestamp in milliseconds
	NotebooksModifiedAfter types.Int64 `tfsdk:"-"`
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"-"`
}

func (newState *ListWorkspaceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListWorkspaceRequest) {
}

func (newState *ListWorkspaceRequest) SyncEffectiveFieldsDuringRead(existingState ListWorkspaceRequest) {
}

func (a ListWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListWorkspaceRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"notebooks_modified_after": types.Int64Type,
			"path":                     types.StringType,
		},
	}
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path types.String `tfsdk:"path" tf:""`
}

func (newState *Mkdirs) SyncEffectiveFieldsDuringCreateOrUpdate(plan Mkdirs) {
}

func (newState *Mkdirs) SyncEffectiveFieldsDuringRead(existingState Mkdirs) {
}

func (a Mkdirs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Mkdirs) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type MkdirsResponse struct {
}

func (newState *MkdirsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MkdirsResponse) {
}

func (newState *MkdirsResponse) SyncEffectiveFieldsDuringRead(existingState MkdirsResponse) {
}

func (a MkdirsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a MkdirsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ObjectInfo struct {
	// Only applicable to files. The creation UTC timestamp.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language types.String `tfsdk:"language" tf:"optional"`
	// Only applicable to files, the last modified UTC timestamp.
	ModifiedAt types.Int64 `tfsdk:"modified_at" tf:"optional"`
	// Unique identifier for the object.
	ObjectId types.Int64 `tfsdk:"object_id" tf:"optional"`
	// The type of the object in workspace.
	//
	// - `NOTEBOOK`: document that contains runnable code, visualizations, and
	// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
	// file - `REPO`: repository - `DASHBOARD`: Lakeview dashboard
	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
	// The absolute path of the object.
	Path types.String `tfsdk:"path" tf:"optional"`
	// A unique identifier for the object that is consistent across all
	// Databricks APIs.
	ResourceId types.String `tfsdk:"resource_id" tf:"optional"`
	// Only applicable to files. The file size in bytes can be returned.
	Size types.Int64 `tfsdk:"size" tf:"optional"`
}

func (newState *ObjectInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ObjectInfo) {
}

func (newState *ObjectInfo) SyncEffectiveFieldsDuringRead(existingState ObjectInfo) {
}

func (a ObjectInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ObjectInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_at":  types.Int64Type,
			"language":    types.StringType,
			"modified_at": types.Int64Type,
			"object_id":   types.Int64Type,
			"object_type": types.StringType,
			"path":        types.StringType,
			"resource_id": types.StringType,
			"size":        types.Int64Type,
		},
	}
}

type PutAcl struct {
	// The permission level applied to the principal.
	Permission types.String `tfsdk:"permission" tf:""`
	// The principal in which the permission is applied.
	Principal types.String `tfsdk:"principal" tf:""`
	// The name of the scope to apply permissions to.
	Scope types.String `tfsdk:"scope" tf:""`
}

func (newState *PutAcl) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutAcl) {
}

func (newState *PutAcl) SyncEffectiveFieldsDuringRead(existingState PutAcl) {
}

func (a PutAcl) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PutAcl) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
			"principal":  types.StringType,
			"scope":      types.StringType,
		},
	}
}

type PutAclResponse struct {
}

func (newState *PutAclResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutAclResponse) {
}

func (newState *PutAclResponse) SyncEffectiveFieldsDuringRead(existingState PutAclResponse) {
}

func (a PutAclResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PutAclResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type PutSecret struct {
	// If specified, value will be stored as bytes.
	BytesValue types.String `tfsdk:"bytes_value" tf:"optional"`
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key" tf:""`
	// The name of the scope to which the secret will be associated with.
	Scope types.String `tfsdk:"scope" tf:""`
	// If specified, note that the value will be stored in UTF-8 (MB4) form.
	StringValue types.String `tfsdk:"string_value" tf:"optional"`
}

func (newState *PutSecret) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutSecret) {
}

func (newState *PutSecret) SyncEffectiveFieldsDuringRead(existingState PutSecret) {
}

func (a PutSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PutSecret) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bytes_value":  types.StringType,
			"key":          types.StringType,
			"scope":        types.StringType,
			"string_value": types.StringType,
		},
	}
}

type PutSecretResponse struct {
}

func (newState *PutSecretResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutSecretResponse) {
}

func (newState *PutSecretResponse) SyncEffectiveFieldsDuringRead(existingState PutSecretResponse) {
}

func (a PutSecretResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PutSecretResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RepoAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *RepoAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoAccessControlRequest) {
}

func (newState *RepoAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState RepoAccessControlRequest) {
}

func (a RepoAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RepoAccessControlRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type RepoAccessControlResponse struct {
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

func (newState *RepoAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoAccessControlResponse) {
}

func (newState *RepoAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState RepoAccessControlResponse) {
}

func (a RepoAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(RepoPermission{}),
	}
}

func (a RepoAccessControlResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: RepoPermission{}.ToAttrType(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// Git folder (repo) information.
type RepoInfo struct {
	// Name of the current git branch of the git folder (repo).
	Branch types.String `tfsdk:"branch" tf:"optional"`
	// Current git commit id of the git folder (repo).
	HeadCommitId types.String `tfsdk:"head_commit_id" tf:"optional"`
	// Id of the git folder (repo) in the Workspace.
	Id types.Int64 `tfsdk:"id" tf:"optional"`
	// Root path of the git folder (repo) in the Workspace.
	Path types.String `tfsdk:"path" tf:"optional"`
	// Git provider of the remote git repository, e.g. `gitHub`.
	Provider types.String `tfsdk:"provider" tf:"optional"`
	// Sparse checkout config for the git folder (repo).
	SparseCheckout types.List `tfsdk:"sparse_checkout" tf:"optional,object"`
	// URL of the remote git repository.
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *RepoInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoInfo) {
}

func (newState *RepoInfo) SyncEffectiveFieldsDuringRead(existingState RepoInfo) {
}

func (a RepoInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

func (a RepoInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":         types.StringType,
			"head_commit_id": types.StringType,
			"id":             types.Int64Type,
			"path":           types.StringType,
			"provider":       types.StringType,
			"sparse_checkout": basetypes.ListType{
				ElemType: SparseCheckout{}.ToAttrType(ctx),
			},
			"url": types.StringType,
		},
	}
}

type RepoPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *RepoPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoPermission) {
}

func (newState *RepoPermission) SyncEffectiveFieldsDuringRead(existingState RepoPermission) {
}

func (a RepoPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

func (a RepoPermission) ToAttrType(ctx context.Context) types.ObjectType {
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

type RepoPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *RepoPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoPermissions) {
}

func (newState *RepoPermissions) SyncEffectiveFieldsDuringRead(existingState RepoPermissions) {
}

func (a RepoPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RepoAccessControlResponse{}),
	}
}

func (a RepoPermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: RepoAccessControlResponse{}.ToAttrType(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type RepoPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *RepoPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoPermissionsDescription) {
}

func (newState *RepoPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState RepoPermissionsDescription) {
}

func (a RepoPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a RepoPermissionsDescription) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type RepoPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

func (newState *RepoPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoPermissionsRequest) {
}

func (newState *RepoPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState RepoPermissionsRequest) {
}

func (a RepoPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RepoAccessControlRequest{}),
	}
}

func (a RepoPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: RepoAccessControlRequest{}.ToAttrType(ctx),
			},
			"repo_id": types.StringType,
		},
	}
}

type SecretMetadata struct {
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The last updated timestamp (in milliseconds) for the secret.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
}

func (newState *SecretMetadata) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecretMetadata) {
}

func (newState *SecretMetadata) SyncEffectiveFieldsDuringRead(existingState SecretMetadata) {
}

func (a SecretMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a SecretMetadata) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":                    types.StringType,
			"last_updated_timestamp": types.Int64Type,
		},
	}
}

type SecretScope struct {
	// The type of secret scope backend.
	BackendType types.String `tfsdk:"backend_type" tf:"optional"`
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	KeyvaultMetadata types.List `tfsdk:"keyvault_metadata" tf:"optional,object"`
	// A unique name to identify the secret scope.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *SecretScope) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecretScope) {
}

func (newState *SecretScope) SyncEffectiveFieldsDuringRead(existingState SecretScope) {
}

func (a SecretScope) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"keyvault_metadata": reflect.TypeOf(AzureKeyVaultSecretScopeMetadata{}),
	}
}

func (a SecretScope) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"backend_type": types.StringType,
			"keyvault_metadata": basetypes.ListType{
				ElemType: AzureKeyVaultSecretScopeMetadata{}.ToAttrType(ctx),
			},
			"name": types.StringType,
		},
	}
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckout struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns types.List `tfsdk:"patterns" tf:"optional"`
}

func (newState *SparseCheckout) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparseCheckout) {
}

func (newState *SparseCheckout) SyncEffectiveFieldsDuringRead(existingState SparseCheckout) {
}

func (a SparseCheckout) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"patterns": reflect.TypeOf(types.String{}),
	}
}

func (a SparseCheckout) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"patterns": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckoutUpdate struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns types.List `tfsdk:"patterns" tf:"optional"`
}

func (newState *SparseCheckoutUpdate) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparseCheckoutUpdate) {
}

func (newState *SparseCheckoutUpdate) SyncEffectiveFieldsDuringRead(existingState SparseCheckoutUpdate) {
}

func (a SparseCheckoutUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"patterns": reflect.TypeOf(types.String{}),
	}
}

func (a SparseCheckoutUpdate) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"patterns": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type UpdateCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	GitProvider types.String `tfsdk:"git_provider" tf:""`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername types.String `tfsdk:"git_username" tf:"optional"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken types.String `tfsdk:"personal_access_token" tf:"optional"`
}

func (newState *UpdateCredentialsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCredentialsRequest) {
}

func (newState *UpdateCredentialsRequest) SyncEffectiveFieldsDuringRead(existingState UpdateCredentialsRequest) {
}

func (a UpdateCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateCredentialsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id":         types.Int64Type,
			"git_provider":          types.StringType,
			"git_username":          types.StringType,
			"personal_access_token": types.StringType,
		},
	}
}

type UpdateCredentialsResponse struct {
}

func (newState *UpdateCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCredentialsResponse) {
}

func (newState *UpdateCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState UpdateCredentialsResponse) {
}

func (a UpdateCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateCredentialsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateRepoRequest struct {
	// Branch that the local version of the repo is checked out to.
	Branch types.String `tfsdk:"branch" tf:"optional"`
	// ID of the Git folder (repo) object in the workspace.
	RepoId types.Int64 `tfsdk:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	SparseCheckout types.List `tfsdk:"sparse_checkout" tf:"optional,object"`
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	Tag types.String `tfsdk:"tag" tf:"optional"`
}

func (newState *UpdateRepoRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRepoRequest) {
}

func (newState *UpdateRepoRequest) SyncEffectiveFieldsDuringRead(existingState UpdateRepoRequest) {
}

func (a UpdateRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckoutUpdate{}),
	}
}

func (a UpdateRepoRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":  types.StringType,
			"repo_id": types.Int64Type,
			"sparse_checkout": basetypes.ListType{
				ElemType: SparseCheckoutUpdate{}.ToAttrType(ctx),
			},
			"tag": types.StringType,
		},
	}
}

type UpdateRepoResponse struct {
}

func (newState *UpdateRepoResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRepoResponse) {
}

func (newState *UpdateRepoResponse) SyncEffectiveFieldsDuringRead(existingState UpdateRepoResponse) {
}

func (a UpdateRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateRepoResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type WorkspaceObjectAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *WorkspaceObjectAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectAccessControlRequest) {
}

func (newState *WorkspaceObjectAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectAccessControlRequest) {
}

func (a WorkspaceObjectAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a WorkspaceObjectAccessControlRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type WorkspaceObjectAccessControlResponse struct {
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

func (newState *WorkspaceObjectAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectAccessControlResponse) {
}

func (newState *WorkspaceObjectAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectAccessControlResponse) {
}

func (a WorkspaceObjectAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(WorkspaceObjectPermission{}),
	}
}

func (a WorkspaceObjectAccessControlResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: WorkspaceObjectPermission{}.ToAttrType(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type WorkspaceObjectPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *WorkspaceObjectPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectPermission) {
}

func (newState *WorkspaceObjectPermission) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectPermission) {
}

func (a WorkspaceObjectPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

func (a WorkspaceObjectPermission) ToAttrType(ctx context.Context) types.ObjectType {
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

type WorkspaceObjectPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *WorkspaceObjectPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectPermissions) {
}

func (newState *WorkspaceObjectPermissions) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectPermissions) {
}

func (a WorkspaceObjectPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WorkspaceObjectAccessControlResponse{}),
	}
}

func (a WorkspaceObjectPermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: WorkspaceObjectAccessControlResponse{}.ToAttrType(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type WorkspaceObjectPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *WorkspaceObjectPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectPermissionsDescription) {
}

func (newState *WorkspaceObjectPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectPermissionsDescription) {
}

func (a WorkspaceObjectPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a WorkspaceObjectPermissionsDescription) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type WorkspaceObjectPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-"`
}

func (newState *WorkspaceObjectPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectPermissionsRequest) {
}

func (newState *WorkspaceObjectPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectPermissionsRequest) {
}

func (a WorkspaceObjectPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WorkspaceObjectAccessControlRequest{}),
	}
}

func (a WorkspaceObjectPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: WorkspaceObjectAccessControlRequest{}.ToAttrType(ctx),
			},
			"workspace_object_id":   types.StringType,
			"workspace_object_type": types.StringType,
		},
	}
}
