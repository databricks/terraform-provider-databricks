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

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AclItem struct {
	// The permission level applied to the principal.
	Permission types.String `tfsdk:"permission"`
	// The principal in which the permission is applied.
	Principal types.String `tfsdk:"principal"`
}

func (newState *AclItem) SyncEffectiveFieldsDuringCreateOrUpdate(plan AclItem) {
}

func (newState *AclItem) SyncEffectiveFieldsDuringRead(existingState AclItem) {
}

func (c AclItem) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()
	attrs["principal"] = attrs["principal"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AclItem.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AclItem) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AclItem
// only implements ToObjectValue() and Type().
func (o AclItem) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": o.Permission,
			"principal":  o.Principal,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AclItem) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
			"principal":  types.StringType,
		},
	}
}

type AzureKeyVaultSecretScopeMetadata struct {
	// The DNS of the KeyVault
	DnsName types.String `tfsdk:"dns_name"`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	ResourceId types.String `tfsdk:"resource_id"`
}

func (newState *AzureKeyVaultSecretScopeMetadata) SyncEffectiveFieldsDuringCreateOrUpdate(plan AzureKeyVaultSecretScopeMetadata) {
}

func (newState *AzureKeyVaultSecretScopeMetadata) SyncEffectiveFieldsDuringRead(existingState AzureKeyVaultSecretScopeMetadata) {
}

func (c AzureKeyVaultSecretScopeMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dns_name"] = attrs["dns_name"].SetRequired()
	attrs["resource_id"] = attrs["resource_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AzureKeyVaultSecretScopeMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AzureKeyVaultSecretScopeMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureKeyVaultSecretScopeMetadata
// only implements ToObjectValue() and Type().
func (o AzureKeyVaultSecretScopeMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dns_name":    o.DnsName,
			"resource_id": o.ResourceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureKeyVaultSecretScopeMetadata) Type(ctx context.Context) attr.Type {
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
	GitProvider types.String `tfsdk:"git_provider"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername types.String `tfsdk:"git_username"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken types.String `tfsdk:"personal_access_token"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialsRequest
// only implements ToObjectValue() and Type().
func (o CreateCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"git_provider":          o.GitProvider,
			"git_username":          o.GitUsername,
			"personal_access_token": o.PersonalAccessToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialsRequest) Type(ctx context.Context) attr.Type {
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
	CredentialId types.Int64 `tfsdk:"credential_id"`
	// The Git provider associated with the credential.
	GitProvider types.String `tfsdk:"git_provider"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername types.String `tfsdk:"git_username"`
}

func (newState *CreateCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCredentialsResponse) {
}

func (newState *CreateCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState CreateCredentialsResponse) {
}

func (c CreateCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetRequired()
	attrs["git_provider"] = attrs["git_provider"].SetRequired()
	attrs["git_username"] = attrs["git_username"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialsResponse
// only implements ToObjectValue() and Type().
func (o CreateCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
			"git_provider":  o.GitProvider,
			"git_username":  o.GitUsername,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialsResponse) Type(ctx context.Context) attr.Type {
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
	Path types.String `tfsdk:"path"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	Provider types.String `tfsdk:"provider"`
	// If specified, the repo will be created with sparse checkout enabled. You
	// cannot enable/disable sparse checkout after the repo is created.
	SparseCheckout types.Object `tfsdk:"sparse_checkout"`
	// URL of the Git repository to be linked.
	Url types.String `tfsdk:"url"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRepoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRepoRequest
// only implements ToObjectValue() and Type().
func (o CreateRepoRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path":            o.Path,
			"provider":        o.Provider,
			"sparse_checkout": o.SparseCheckout,
			"url":             o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateRepoRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":            types.StringType,
			"provider":        types.StringType,
			"sparse_checkout": SparseCheckout{}.Type(ctx),
			"url":             types.StringType,
		},
	}
}

// GetSparseCheckout returns the value of the SparseCheckout field in CreateRepoRequest as
// a SparseCheckout value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRepoRequest) GetSparseCheckout(ctx context.Context) (SparseCheckout, bool) {
	var e SparseCheckout
	if o.SparseCheckout.IsNull() || o.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout
	d := o.SparseCheckout.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSparseCheckout sets the value of the SparseCheckout field in CreateRepoRequest.
func (o *CreateRepoRequest) SetSparseCheckout(ctx context.Context, v SparseCheckout) {
	vs := v.ToObjectValue(ctx)
	o.SparseCheckout = vs
}

type CreateRepoResponse struct {
	// Branch that the Git folder (repo) is checked out to.
	Branch types.String `tfsdk:"branch"`
	// SHA-1 hash representing the commit ID of the current HEAD of the Git
	// folder (repo).
	HeadCommitId types.String `tfsdk:"head_commit_id"`
	// ID of the Git folder (repo) object in the workspace.
	Id types.Int64 `tfsdk:"id"`
	// Path of the Git folder (repo) in the workspace.
	Path types.String `tfsdk:"path"`
	// Git provider of the linked Git repository.
	Provider types.String `tfsdk:"provider"`
	// Sparse checkout settings for the Git folder (repo).
	SparseCheckout types.Object `tfsdk:"sparse_checkout"`
	// URL of the linked Git repository.
	Url types.String `tfsdk:"url"`
}

func (newState *CreateRepoResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRepoResponse) {
}

func (newState *CreateRepoResponse) SyncEffectiveFieldsDuringRead(existingState CreateRepoResponse) {
}

func (c CreateRepoResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetOptional()
	attrs["head_commit_id"] = attrs["head_commit_id"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()
	attrs["provider"] = attrs["provider"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRepoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRepoResponse
// only implements ToObjectValue() and Type().
func (o CreateRepoResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":          o.Branch,
			"head_commit_id":  o.HeadCommitId,
			"id":              o.Id,
			"path":            o.Path,
			"provider":        o.Provider,
			"sparse_checkout": o.SparseCheckout,
			"url":             o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateRepoResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":          types.StringType,
			"head_commit_id":  types.StringType,
			"id":              types.Int64Type,
			"path":            types.StringType,
			"provider":        types.StringType,
			"sparse_checkout": SparseCheckout{}.Type(ctx),
			"url":             types.StringType,
		},
	}
}

// GetSparseCheckout returns the value of the SparseCheckout field in CreateRepoResponse as
// a SparseCheckout value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRepoResponse) GetSparseCheckout(ctx context.Context) (SparseCheckout, bool) {
	var e SparseCheckout
	if o.SparseCheckout.IsNull() || o.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout
	d := o.SparseCheckout.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSparseCheckout sets the value of the SparseCheckout field in CreateRepoResponse.
func (o *CreateRepoResponse) SetSparseCheckout(ctx context.Context, v SparseCheckout) {
	vs := v.ToObjectValue(ctx)
	o.SparseCheckout = vs
}

type CreateScope struct {
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	BackendAzureKeyvault types.Object `tfsdk:"backend_azure_keyvault"`
	// The principal that is initially granted `MANAGE` permission to the
	// created scope.
	InitialManagePrincipal types.String `tfsdk:"initial_manage_principal"`
	// Scope name requested by the user. Scope names are unique.
	Scope types.String `tfsdk:"scope"`
	// The backend type the scope will be created with. If not specified, will
	// default to `DATABRICKS`
	ScopeBackendType types.String `tfsdk:"scope_backend_type"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateScope.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateScope) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"backend_azure_keyvault": reflect.TypeOf(AzureKeyVaultSecretScopeMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateScope
// only implements ToObjectValue() and Type().
func (o CreateScope) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"backend_azure_keyvault":   o.BackendAzureKeyvault,
			"initial_manage_principal": o.InitialManagePrincipal,
			"scope":                    o.Scope,
			"scope_backend_type":       o.ScopeBackendType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateScope) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"backend_azure_keyvault":   AzureKeyVaultSecretScopeMetadata{}.Type(ctx),
			"initial_manage_principal": types.StringType,
			"scope":                    types.StringType,
			"scope_backend_type":       types.StringType,
		},
	}
}

// GetBackendAzureKeyvault returns the value of the BackendAzureKeyvault field in CreateScope as
// a AzureKeyVaultSecretScopeMetadata value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateScope) GetBackendAzureKeyvault(ctx context.Context) (AzureKeyVaultSecretScopeMetadata, bool) {
	var e AzureKeyVaultSecretScopeMetadata
	if o.BackendAzureKeyvault.IsNull() || o.BackendAzureKeyvault.IsUnknown() {
		return e, false
	}
	var v []AzureKeyVaultSecretScopeMetadata
	d := o.BackendAzureKeyvault.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetBackendAzureKeyvault sets the value of the BackendAzureKeyvault field in CreateScope.
func (o *CreateScope) SetBackendAzureKeyvault(ctx context.Context, v AzureKeyVaultSecretScopeMetadata) {
	vs := v.ToObjectValue(ctx)
	o.BackendAzureKeyvault = vs
}

type CreateScopeResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateScopeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateScopeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateScopeResponse
// only implements ToObjectValue() and Type().
func (o CreateScopeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o CreateScopeResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	CredentialId types.Int64 `tfsdk:"credential_id"`
	// The Git provider associated with the credential.
	GitProvider types.String `tfsdk:"git_provider"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername types.String `tfsdk:"git_username"`
}

func (newState *CredentialInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CredentialInfo) {
}

func (newState *CredentialInfo) SyncEffectiveFieldsDuringRead(existingState CredentialInfo) {
}

func (c CredentialInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetRequired()
	attrs["git_provider"] = attrs["git_provider"].SetOptional()
	attrs["git_username"] = attrs["git_username"].SetOptional()

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
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CredentialInfo
// only implements ToObjectValue() and Type().
func (o CredentialInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
			"git_provider":  o.GitProvider,
			"git_username":  o.GitUsername,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CredentialInfo) Type(ctx context.Context) attr.Type {
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
	Path types.String `tfsdk:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive types.Bool `tfsdk:"recursive"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Delete.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Delete) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Delete
// only implements ToObjectValue() and Type().
func (o Delete) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path":      o.Path,
			"recursive": o.Recursive,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Delete) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":      types.StringType,
			"recursive": types.BoolType,
		},
	}
}

type DeleteAcl struct {
	// The principal to remove an existing ACL from.
	Principal types.String `tfsdk:"principal"`
	// The name of the scope to remove permissions from.
	Scope types.String `tfsdk:"scope"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAcl.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAcl) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAcl
// only implements ToObjectValue() and Type().
func (o DeleteAcl) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal": o.Principal,
			"scope":     o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAcl) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"scope":     types.StringType,
		},
	}
}

type DeleteAclResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAclResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAclResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAclResponse
// only implements ToObjectValue() and Type().
func (o DeleteAclResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAclResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialsRequest
// only implements ToObjectValue() and Type().
func (o DeleteCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialsRequest) Type(ctx context.Context) attr.Type {
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

func (c DeleteCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialsResponse
// only implements ToObjectValue() and Type().
func (o DeleteCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRepoRequest struct {
	// The ID for the corresponding repo to delete.
	RepoId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRepoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRepoRequest
// only implements ToObjectValue() and Type().
func (o DeleteRepoRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": o.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRepoRequest) Type(ctx context.Context) attr.Type {
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

func (c DeleteRepoResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRepoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRepoResponse
// only implements ToObjectValue() and Type().
func (o DeleteRepoResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRepoResponse) Type(ctx context.Context) attr.Type {
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

type DeleteScope struct {
	// Name of the scope to delete.
	Scope types.String `tfsdk:"scope"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteScope.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteScope) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteScope
// only implements ToObjectValue() and Type().
func (o DeleteScope) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteScope) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope": types.StringType,
		},
	}
}

type DeleteScopeResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteScopeResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteScopeResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteScopeResponse
// only implements ToObjectValue() and Type().
func (o DeleteScopeResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteScopeResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteSecret struct {
	// Name of the secret to delete.
	Key types.String `tfsdk:"key"`
	// The name of the scope that contains the secret to delete.
	Scope types.String `tfsdk:"scope"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSecret.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSecret
// only implements ToObjectValue() and Type().
func (o DeleteSecret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"scope": o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSecret) Type(ctx context.Context) attr.Type {
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

func (c DeleteSecretResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSecretResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSecretResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSecretResponse
// only implements ToObjectValue() and Type().
func (o DeleteSecretResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSecretResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExportRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportRequest
// only implements ToObjectValue() and Type().
func (o ExportRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"format": o.Format,
			"path":   o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExportRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"format": types.StringType,
			"path":   types.StringType,
		},
	}
}

// The request field `direct_download` determines whether a JSON response or
// binary contents are returned by this endpoint.
type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	Content types.String `tfsdk:"content"`
	// The file type of the exported file.
	FileType types.String `tfsdk:"file_type"`
}

func (newState *ExportResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportResponse) {
}

func (newState *ExportResponse) SyncEffectiveFieldsDuringRead(existingState ExportResponse) {
}

func (c ExportResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetOptional()
	attrs["file_type"] = attrs["file_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExportResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportResponse
// only implements ToObjectValue() and Type().
func (o ExportResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content":   o.Content,
			"file_type": o.FileType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExportResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content":   types.StringType,
			"file_type": types.StringType,
		},
	}
}

type GetAclRequest struct {
	// The principal to fetch ACL information for.
	Principal types.String `tfsdk:"-"`
	// The name of the scope to fetch ACL information from.
	Scope types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAclRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAclRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAclRequest
// only implements ToObjectValue() and Type().
func (o GetAclRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal": o.Principal,
			"scope":     o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAclRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"scope":     types.StringType,
		},
	}
}

type GetCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialsRequest
// only implements ToObjectValue() and Type().
func (o GetCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCredentialsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.Int64Type,
		},
	}
}

type GetCredentialsResponse struct {
	// ID of the credential object in the workspace.
	CredentialId types.Int64 `tfsdk:"credential_id"`
	// The Git provider associated with the credential.
	GitProvider types.String `tfsdk:"git_provider"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername types.String `tfsdk:"git_username"`
}

func (newState *GetCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCredentialsResponse) {
}

func (newState *GetCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState GetCredentialsResponse) {
}

func (c GetCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetRequired()
	attrs["git_provider"] = attrs["git_provider"].SetOptional()
	attrs["git_username"] = attrs["git_username"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialsResponse
// only implements ToObjectValue() and Type().
func (o GetCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
			"git_provider":  o.GitProvider,
			"git_username":  o.GitUsername,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCredentialsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.Int64Type,
			"git_provider":  types.StringType,
			"git_username":  types.StringType,
		},
	}
}

type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRepoPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRepoPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetRepoPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": o.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRepoPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.StringType,
		},
	}
}

type GetRepoPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetRepoPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRepoPermissionLevelsResponse) {
}

func (newState *GetRepoPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetRepoPermissionLevelsResponse) {
}

func (c GetRepoPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRepoPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRepoPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(RepoPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetRepoPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRepoPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: RepoPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetRepoPermissionLevelsResponse as
// a slice of RepoPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetRepoPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]RepoPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []RepoPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetRepoPermissionLevelsResponse.
func (o *GetRepoPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []RepoPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetRepoPermissionsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRepoPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRepoPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetRepoPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": o.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRepoPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.StringType,
		},
	}
}

type GetRepoRequest struct {
	// ID of the Git folder (repo) object in the workspace.
	RepoId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRepoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoRequest
// only implements ToObjectValue() and Type().
func (o GetRepoRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": o.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRepoRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.Int64Type,
		},
	}
}

type GetRepoResponse struct {
	// Branch that the local version of the repo is checked out to.
	Branch types.String `tfsdk:"branch"`
	// SHA-1 hash representing the commit ID of the current HEAD of the repo.
	HeadCommitId types.String `tfsdk:"head_commit_id"`
	// ID of the Git folder (repo) object in the workspace.
	Id types.Int64 `tfsdk:"id"`
	// Path of the Git folder (repo) in the workspace.
	Path types.String `tfsdk:"path"`
	// Git provider of the linked Git repository.
	Provider types.String `tfsdk:"provider"`
	// Sparse checkout settings for the Git folder (repo).
	SparseCheckout types.Object `tfsdk:"sparse_checkout"`
	// URL of the linked Git repository.
	Url types.String `tfsdk:"url"`
}

func (newState *GetRepoResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRepoResponse) {
}

func (newState *GetRepoResponse) SyncEffectiveFieldsDuringRead(existingState GetRepoResponse) {
}

func (c GetRepoResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetOptional()
	attrs["head_commit_id"] = attrs["head_commit_id"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()
	attrs["provider"] = attrs["provider"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRepoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoResponse
// only implements ToObjectValue() and Type().
func (o GetRepoResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":          o.Branch,
			"head_commit_id":  o.HeadCommitId,
			"id":              o.Id,
			"path":            o.Path,
			"provider":        o.Provider,
			"sparse_checkout": o.SparseCheckout,
			"url":             o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRepoResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":          types.StringType,
			"head_commit_id":  types.StringType,
			"id":              types.Int64Type,
			"path":            types.StringType,
			"provider":        types.StringType,
			"sparse_checkout": SparseCheckout{}.Type(ctx),
			"url":             types.StringType,
		},
	}
}

// GetSparseCheckout returns the value of the SparseCheckout field in GetRepoResponse as
// a SparseCheckout value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetRepoResponse) GetSparseCheckout(ctx context.Context) (SparseCheckout, bool) {
	var e SparseCheckout
	if o.SparseCheckout.IsNull() || o.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout
	d := o.SparseCheckout.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSparseCheckout sets the value of the SparseCheckout field in GetRepoResponse.
func (o *GetRepoResponse) SetSparseCheckout(ctx context.Context, v SparseCheckout) {
	vs := v.ToObjectValue(ctx)
	o.SparseCheckout = vs
}

type GetSecretRequest struct {
	// The key to fetch secret for.
	Key types.String `tfsdk:"-"`
	// The name of the scope to fetch secret information from.
	Scope types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSecretRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSecretRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSecretRequest
// only implements ToObjectValue() and Type().
func (o GetSecretRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"scope": o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSecretRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"scope": types.StringType,
		},
	}
}

type GetSecretResponse struct {
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key"`
	// The value of the secret in its byte representation.
	Value types.String `tfsdk:"value"`
}

func (newState *GetSecretResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSecretResponse) {
}

func (newState *GetSecretResponse) SyncEffectiveFieldsDuringRead(existingState GetSecretResponse) {
}

func (c GetSecretResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSecretResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSecretResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSecretResponse
// only implements ToObjectValue() and Type().
func (o GetSecretResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSecretResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatusRequest
// only implements ToObjectValue() and Type().
func (o GetStatusRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStatusRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type GetWorkspaceObjectPermissionLevelsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceObjectPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceObjectPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetWorkspaceObjectPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_object_id":   o.WorkspaceObjectId,
			"workspace_object_type": o.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceObjectPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_object_id":   types.StringType,
			"workspace_object_type": types.StringType,
		},
	}
}

type GetWorkspaceObjectPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetWorkspaceObjectPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceObjectPermissionLevelsResponse) {
}

func (newState *GetWorkspaceObjectPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceObjectPermissionLevelsResponse) {
}

func (c GetWorkspaceObjectPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceObjectPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceObjectPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(WorkspaceObjectPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetWorkspaceObjectPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceObjectPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: WorkspaceObjectPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetWorkspaceObjectPermissionLevelsResponse as
// a slice of WorkspaceObjectPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetWorkspaceObjectPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]WorkspaceObjectPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetWorkspaceObjectPermissionLevelsResponse.
func (o *GetWorkspaceObjectPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []WorkspaceObjectPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetWorkspaceObjectPermissionsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceObjectPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceObjectPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetWorkspaceObjectPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_object_id":   o.WorkspaceObjectId,
			"workspace_object_type": o.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceObjectPermissionsRequest) Type(ctx context.Context) attr.Type {
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
	Content types.String `tfsdk:"content"`
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
	Format types.String `tfsdk:"format"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language types.String `tfsdk:"language"`
	// The flag that specifies whether to overwrite existing object. It is
	// `false` by default. For `DBC` format, `overwrite` is not supported since
	// it may contain a directory.
	Overwrite types.Bool `tfsdk:"overwrite"`
	// The absolute path of the object or directory. Importing a directory is
	// only supported for the `DBC` and `SOURCE` formats.
	Path types.String `tfsdk:"path"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Import.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Import) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Import
// only implements ToObjectValue() and Type().
func (o Import) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content":   o.Content,
			"format":    o.Format,
			"language":  o.Language,
			"overwrite": o.Overwrite,
			"path":      o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Import) Type(ctx context.Context) attr.Type {
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

func (c ImportResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ImportResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ImportResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ImportResponse
// only implements ToObjectValue() and Type().
func (o ImportResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ImportResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListAclsRequest struct {
	// The name of the scope to fetch ACL information from.
	Scope types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAclsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAclsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAclsRequest
// only implements ToObjectValue() and Type().
func (o ListAclsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAclsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope": types.StringType,
		},
	}
}

type ListAclsResponse struct {
	// The associated ACLs rule applied to principals in the given scope.
	Items types.List `tfsdk:"items"`
}

func (newState *ListAclsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAclsResponse) {
}

func (newState *ListAclsResponse) SyncEffectiveFieldsDuringRead(existingState ListAclsResponse) {
}

func (c ListAclsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["items"] = attrs["items"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAclsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAclsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(AclItem{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAclsResponse
// only implements ToObjectValue() and Type().
func (o ListAclsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items": o.Items,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAclsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: AclItem{}.Type(ctx),
			},
		},
	}
}

// GetItems returns the value of the Items field in ListAclsResponse as
// a slice of AclItem values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAclsResponse) GetItems(ctx context.Context) ([]AclItem, bool) {
	if o.Items.IsNull() || o.Items.IsUnknown() {
		return nil, false
	}
	var v []AclItem
	d := o.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListAclsResponse.
func (o *ListAclsResponse) SetItems(ctx context.Context, v []AclItem) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Items = types.ListValueMust(t, vs)
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

type ListCredentialsResponse struct {
	// List of credentials.
	Credentials types.List `tfsdk:"credentials"`
}

func (newState *ListCredentialsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCredentialsResponse) {
}

func (newState *ListCredentialsResponse) SyncEffectiveFieldsDuringRead(existingState ListCredentialsResponse) {
}

func (c ListCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credentials"] = attrs["credentials"].SetOptional()

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
			"credentials": o.Credentials,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCredentialsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials": basetypes.ListType{
				ElemType: CredentialInfo{}.Type(ctx),
			},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListReposRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListReposRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListReposRequest
// only implements ToObjectValue() and Type().
func (o ListReposRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"path_prefix":     o.PathPrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListReposRequest) Type(ctx context.Context) attr.Type {
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
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of Git folders (repos).
	Repos types.List `tfsdk:"repos"`
}

func (newState *ListReposResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListReposResponse) {
}

func (newState *ListReposResponse) SyncEffectiveFieldsDuringRead(existingState ListReposResponse) {
}

func (c ListReposResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["repos"] = attrs["repos"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListReposResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListReposResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"repos": reflect.TypeOf(RepoInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListReposResponse
// only implements ToObjectValue() and Type().
func (o ListReposResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"repos":           o.Repos,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListReposResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"repos": basetypes.ListType{
				ElemType: RepoInfo{}.Type(ctx),
			},
		},
	}
}

// GetRepos returns the value of the Repos field in ListReposResponse as
// a slice of RepoInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListReposResponse) GetRepos(ctx context.Context) ([]RepoInfo, bool) {
	if o.Repos.IsNull() || o.Repos.IsUnknown() {
		return nil, false
	}
	var v []RepoInfo
	d := o.Repos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRepos sets the value of the Repos field in ListReposResponse.
func (o *ListReposResponse) SetRepos(ctx context.Context, v []RepoInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["repos"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Repos = types.ListValueMust(t, vs)
}

type ListResponse struct {
	// List of objects.
	Objects types.List `tfsdk:"objects"`
}

func (newState *ListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListResponse) {
}

func (newState *ListResponse) SyncEffectiveFieldsDuringRead(existingState ListResponse) {
}

func (c ListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["objects"] = attrs["objects"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objects": reflect.TypeOf(ObjectInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListResponse
// only implements ToObjectValue() and Type().
func (o ListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"objects": o.Objects,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"objects": basetypes.ListType{
				ElemType: ObjectInfo{}.Type(ctx),
			},
		},
	}
}

// GetObjects returns the value of the Objects field in ListResponse as
// a slice of ObjectInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListResponse) GetObjects(ctx context.Context) ([]ObjectInfo, bool) {
	if o.Objects.IsNull() || o.Objects.IsUnknown() {
		return nil, false
	}
	var v []ObjectInfo
	d := o.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in ListResponse.
func (o *ListResponse) SetObjects(ctx context.Context, v []ObjectInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Objects = types.ListValueMust(t, vs)
}

type ListScopesRequest struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListScopesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListScopesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListScopesRequest
// only implements ToObjectValue() and Type().
func (o ListScopesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListScopesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListScopesResponse struct {
	// The available secret scopes.
	Scopes types.List `tfsdk:"scopes"`
}

func (newState *ListScopesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListScopesResponse) {
}

func (newState *ListScopesResponse) SyncEffectiveFieldsDuringRead(existingState ListScopesResponse) {
}

func (c ListScopesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["scopes"] = attrs["scopes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListScopesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListScopesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(SecretScope{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListScopesResponse
// only implements ToObjectValue() and Type().
func (o ListScopesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scopes": o.Scopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListScopesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scopes": basetypes.ListType{
				ElemType: SecretScope{}.Type(ctx),
			},
		},
	}
}

// GetScopes returns the value of the Scopes field in ListScopesResponse as
// a slice of SecretScope values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListScopesResponse) GetScopes(ctx context.Context) ([]SecretScope, bool) {
	if o.Scopes.IsNull() || o.Scopes.IsUnknown() {
		return nil, false
	}
	var v []SecretScope
	d := o.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in ListScopesResponse.
func (o *ListScopesResponse) SetScopes(ctx context.Context, v []SecretScope) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	Scope types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSecretsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSecretsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSecretsRequest
// only implements ToObjectValue() and Type().
func (o ListSecretsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSecretsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope": types.StringType,
		},
	}
}

type ListSecretsResponse struct {
	// Metadata information of all secrets contained within the given scope.
	Secrets types.List `tfsdk:"secrets"`
}

func (newState *ListSecretsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSecretsResponse) {
}

func (newState *ListSecretsResponse) SyncEffectiveFieldsDuringRead(existingState ListSecretsResponse) {
}

func (c ListSecretsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["secrets"] = attrs["secrets"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSecretsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSecretsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"secrets": reflect.TypeOf(SecretMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSecretsResponse
// only implements ToObjectValue() and Type().
func (o ListSecretsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"secrets": o.Secrets,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSecretsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"secrets": basetypes.ListType{
				ElemType: SecretMetadata{}.Type(ctx),
			},
		},
	}
}

// GetSecrets returns the value of the Secrets field in ListSecretsResponse as
// a slice of SecretMetadata values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSecretsResponse) GetSecrets(ctx context.Context) ([]SecretMetadata, bool) {
	if o.Secrets.IsNull() || o.Secrets.IsUnknown() {
		return nil, false
	}
	var v []SecretMetadata
	d := o.Secrets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecrets sets the value of the Secrets field in ListSecretsResponse.
func (o *ListSecretsResponse) SetSecrets(ctx context.Context, v []SecretMetadata) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["secrets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Secrets = types.ListValueMust(t, vs)
}

type ListWorkspaceRequest struct {
	// UTC timestamp in milliseconds
	NotebooksModifiedAfter types.Int64 `tfsdk:"-"`
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o ListWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notebooks_modified_after": o.NotebooksModifiedAfter,
			"path":                     o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceRequest) Type(ctx context.Context) attr.Type {
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
	Path types.String `tfsdk:"path"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Mkdirs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Mkdirs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Mkdirs
// only implements ToObjectValue() and Type().
func (o Mkdirs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Mkdirs) Type(ctx context.Context) attr.Type {
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

func (c MkdirsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MkdirsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MkdirsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MkdirsResponse
// only implements ToObjectValue() and Type().
func (o MkdirsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o MkdirsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// The information of the object in workspace. It will be returned by “list“
// and “get-status“.
type ObjectInfo struct {
	// Only applicable to files. The creation UTC timestamp.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// The language of the object. This value is set only if the object type is
	// ``NOTEBOOK``.
	Language types.String `tfsdk:"language"`
	// Only applicable to files, the last modified UTC timestamp.
	ModifiedAt types.Int64 `tfsdk:"modified_at"`
	// Unique identifier for the object.
	ObjectId types.Int64 `tfsdk:"object_id"`
	// The type of the object in workspace.
	//
	// - `NOTEBOOK`: document that contains runnable code, visualizations, and
	// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
	// file - `REPO`: repository - `DASHBOARD`: Lakeview dashboard
	ObjectType types.String `tfsdk:"object_type"`
	// The absolute path of the object.
	Path types.String `tfsdk:"path"`
	// A unique identifier for the object that is consistent across all
	// Databricks APIs.
	ResourceId types.String `tfsdk:"resource_id"`
	// Only applicable to files. The file size in bytes can be returned.
	Size types.Int64 `tfsdk:"size"`
}

func (newState *ObjectInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ObjectInfo) {
}

func (newState *ObjectInfo) SyncEffectiveFieldsDuringRead(existingState ObjectInfo) {
}

func (c ObjectInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["language"] = attrs["language"].SetOptional()
	attrs["modified_at"] = attrs["modified_at"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()
	attrs["resource_id"] = attrs["resource_id"].SetOptional()
	attrs["size"] = attrs["size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ObjectInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ObjectInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ObjectInfo
// only implements ToObjectValue() and Type().
func (o ObjectInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":  o.CreatedAt,
			"language":    o.Language,
			"modified_at": o.ModifiedAt,
			"object_id":   o.ObjectId,
			"object_type": o.ObjectType,
			"path":        o.Path,
			"resource_id": o.ResourceId,
			"size":        o.Size,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ObjectInfo) Type(ctx context.Context) attr.Type {
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
	Permission types.String `tfsdk:"permission"`
	// The principal in which the permission is applied.
	Principal types.String `tfsdk:"principal"`
	// The name of the scope to apply permissions to.
	Scope types.String `tfsdk:"scope"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutAcl.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutAcl) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAcl
// only implements ToObjectValue() and Type().
func (o PutAcl) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": o.Permission,
			"principal":  o.Principal,
			"scope":      o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutAcl) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutAclResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutAclResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAclResponse
// only implements ToObjectValue() and Type().
func (o PutAclResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PutAclResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type PutSecret struct {
	// If specified, value will be stored as bytes.
	BytesValue types.String `tfsdk:"bytes_value"`
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key"`
	// The name of the scope to which the secret will be associated with.
	Scope types.String `tfsdk:"scope"`
	// If specified, note that the value will be stored in UTF-8 (MB4) form.
	StringValue types.String `tfsdk:"string_value"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutSecret.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutSecret
// only implements ToObjectValue() and Type().
func (o PutSecret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bytes_value":  o.BytesValue,
			"key":          o.Key,
			"scope":        o.Scope,
			"string_value": o.StringValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutSecret) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutSecretResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutSecretResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutSecretResponse
// only implements ToObjectValue() and Type().
func (o PutSecretResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o PutSecretResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RepoAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *RepoAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoAccessControlRequest) {
}

func (newState *RepoAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState RepoAccessControlRequest) {
}

func (c RepoAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoAccessControlRequest
// only implements ToObjectValue() and Type().
func (o RepoAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoAccessControlRequest) Type(ctx context.Context) attr.Type {
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
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *RepoAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoAccessControlResponse) {
}

func (newState *RepoAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState RepoAccessControlResponse) {
}

func (c RepoAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(RepoPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoAccessControlResponse
// only implements ToObjectValue() and Type().
func (o RepoAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: RepoPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in RepoAccessControlResponse as
// a slice of RepoPermission values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepoAccessControlResponse) GetAllPermissions(ctx context.Context) ([]RepoPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []RepoPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in RepoAccessControlResponse.
func (o *RepoAccessControlResponse) SetAllPermissions(ctx context.Context, v []RepoPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

// Git folder (repo) information.
type RepoInfo struct {
	// Name of the current git branch of the git folder (repo).
	Branch types.String `tfsdk:"branch"`
	// Current git commit id of the git folder (repo).
	HeadCommitId types.String `tfsdk:"head_commit_id"`
	// Id of the git folder (repo) in the Workspace.
	Id types.Int64 `tfsdk:"id"`
	// Root path of the git folder (repo) in the Workspace.
	Path types.String `tfsdk:"path"`
	// Git provider of the remote git repository, e.g. `gitHub`.
	Provider types.String `tfsdk:"provider"`
	// Sparse checkout config for the git folder (repo).
	SparseCheckout types.Object `tfsdk:"sparse_checkout"`
	// URL of the remote git repository.
	Url types.String `tfsdk:"url"`
}

func (newState *RepoInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoInfo) {
}

func (newState *RepoInfo) SyncEffectiveFieldsDuringRead(existingState RepoInfo) {
}

func (c RepoInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetOptional()
	attrs["head_commit_id"] = attrs["head_commit_id"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()
	attrs["provider"] = attrs["provider"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoInfo
// only implements ToObjectValue() and Type().
func (o RepoInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":          o.Branch,
			"head_commit_id":  o.HeadCommitId,
			"id":              o.Id,
			"path":            o.Path,
			"provider":        o.Provider,
			"sparse_checkout": o.SparseCheckout,
			"url":             o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":          types.StringType,
			"head_commit_id":  types.StringType,
			"id":              types.Int64Type,
			"path":            types.StringType,
			"provider":        types.StringType,
			"sparse_checkout": SparseCheckout{}.Type(ctx),
			"url":             types.StringType,
		},
	}
}

// GetSparseCheckout returns the value of the SparseCheckout field in RepoInfo as
// a SparseCheckout value.
// If the field is unknown or null, the boolean return value is false.
func (o *RepoInfo) GetSparseCheckout(ctx context.Context) (SparseCheckout, bool) {
	var e SparseCheckout
	if o.SparseCheckout.IsNull() || o.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout
	d := o.SparseCheckout.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSparseCheckout sets the value of the SparseCheckout field in RepoInfo.
func (o *RepoInfo) SetSparseCheckout(ctx context.Context, v SparseCheckout) {
	vs := v.ToObjectValue(ctx)
	o.SparseCheckout = vs
}

type RepoPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *RepoPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoPermission) {
}

func (newState *RepoPermission) SyncEffectiveFieldsDuringRead(existingState RepoPermission) {
}

func (c RepoPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermission
// only implements ToObjectValue() and Type().
func (o RepoPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoPermission) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in RepoPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepoPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in RepoPermission.
func (o *RepoPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type RepoPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *RepoPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoPermissions) {
}

func (newState *RepoPermissions) SyncEffectiveFieldsDuringRead(existingState RepoPermissions) {
}

func (c RepoPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RepoAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissions
// only implements ToObjectValue() and Type().
func (o RepoPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: RepoAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in RepoPermissions as
// a slice of RepoAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepoPermissions) GetAccessControlList(ctx context.Context) ([]RepoAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RepoAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RepoPermissions.
func (o *RepoPermissions) SetAccessControlList(ctx context.Context, v []RepoAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type RepoPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *RepoPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepoPermissionsDescription) {
}

func (newState *RepoPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState RepoPermissionsDescription) {
}

func (c RepoPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissionsDescription
// only implements ToObjectValue() and Type().
func (o RepoPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type RepoPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RepoPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RepoAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissionsRequest
// only implements ToObjectValue() and Type().
func (o RepoPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"repo_id":             o.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: RepoAccessControlRequest{}.Type(ctx),
			},
			"repo_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in RepoPermissionsRequest as
// a slice of RepoAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *RepoPermissionsRequest) GetAccessControlList(ctx context.Context) ([]RepoAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RepoAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RepoPermissionsRequest.
func (o *RepoPermissionsRequest) SetAccessControlList(ctx context.Context, v []RepoAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type SecretMetadata struct {
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key"`
	// The last updated timestamp (in milliseconds) for the secret.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
}

func (newState *SecretMetadata) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecretMetadata) {
}

func (newState *SecretMetadata) SyncEffectiveFieldsDuringRead(existingState SecretMetadata) {
}

func (c SecretMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SecretMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SecretMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretMetadata
// only implements ToObjectValue() and Type().
func (o SecretMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":                    o.Key,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SecretMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":                    types.StringType,
			"last_updated_timestamp": types.Int64Type,
		},
	}
}

type SecretScope struct {
	// The type of secret scope backend.
	BackendType types.String `tfsdk:"backend_type"`
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	KeyvaultMetadata types.Object `tfsdk:"keyvault_metadata"`
	// A unique name to identify the secret scope.
	Name types.String `tfsdk:"name"`
}

func (newState *SecretScope) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecretScope) {
}

func (newState *SecretScope) SyncEffectiveFieldsDuringRead(existingState SecretScope) {
}

func (c SecretScope) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["backend_type"] = attrs["backend_type"].SetOptional()
	attrs["keyvault_metadata"] = attrs["keyvault_metadata"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SecretScope.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SecretScope) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"keyvault_metadata": reflect.TypeOf(AzureKeyVaultSecretScopeMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretScope
// only implements ToObjectValue() and Type().
func (o SecretScope) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"backend_type":      o.BackendType,
			"keyvault_metadata": o.KeyvaultMetadata,
			"name":              o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SecretScope) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"backend_type":      types.StringType,
			"keyvault_metadata": AzureKeyVaultSecretScopeMetadata{}.Type(ctx),
			"name":              types.StringType,
		},
	}
}

// GetKeyvaultMetadata returns the value of the KeyvaultMetadata field in SecretScope as
// a AzureKeyVaultSecretScopeMetadata value.
// If the field is unknown or null, the boolean return value is false.
func (o *SecretScope) GetKeyvaultMetadata(ctx context.Context) (AzureKeyVaultSecretScopeMetadata, bool) {
	var e AzureKeyVaultSecretScopeMetadata
	if o.KeyvaultMetadata.IsNull() || o.KeyvaultMetadata.IsUnknown() {
		return e, false
	}
	var v []AzureKeyVaultSecretScopeMetadata
	d := o.KeyvaultMetadata.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetKeyvaultMetadata sets the value of the KeyvaultMetadata field in SecretScope.
func (o *SecretScope) SetKeyvaultMetadata(ctx context.Context, v AzureKeyVaultSecretScopeMetadata) {
	vs := v.ToObjectValue(ctx)
	o.KeyvaultMetadata = vs
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckout struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns types.List `tfsdk:"patterns"`
}

func (newState *SparseCheckout) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparseCheckout) {
}

func (newState *SparseCheckout) SyncEffectiveFieldsDuringRead(existingState SparseCheckout) {
}

func (c SparseCheckout) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["patterns"] = attrs["patterns"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparseCheckout.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SparseCheckout) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"patterns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparseCheckout
// only implements ToObjectValue() and Type().
func (o SparseCheckout) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"patterns": o.Patterns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparseCheckout) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"patterns": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPatterns returns the value of the Patterns field in SparseCheckout as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SparseCheckout) GetPatterns(ctx context.Context) ([]types.String, bool) {
	if o.Patterns.IsNull() || o.Patterns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Patterns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPatterns sets the value of the Patterns field in SparseCheckout.
func (o *SparseCheckout) SetPatterns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["patterns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Patterns = types.ListValueMust(t, vs)
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckoutUpdate struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns types.List `tfsdk:"patterns"`
}

func (newState *SparseCheckoutUpdate) SyncEffectiveFieldsDuringCreateOrUpdate(plan SparseCheckoutUpdate) {
}

func (newState *SparseCheckoutUpdate) SyncEffectiveFieldsDuringRead(existingState SparseCheckoutUpdate) {
}

func (c SparseCheckoutUpdate) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["patterns"] = attrs["patterns"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SparseCheckoutUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SparseCheckoutUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"patterns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparseCheckoutUpdate
// only implements ToObjectValue() and Type().
func (o SparseCheckoutUpdate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"patterns": o.Patterns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparseCheckoutUpdate) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"patterns": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPatterns returns the value of the Patterns field in SparseCheckoutUpdate as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SparseCheckoutUpdate) GetPatterns(ctx context.Context) ([]types.String, bool) {
	if o.Patterns.IsNull() || o.Patterns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Patterns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPatterns sets the value of the Patterns field in SparseCheckoutUpdate.
func (o *SparseCheckoutUpdate) SetPatterns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["patterns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Patterns = types.ListValueMust(t, vs)
}

type UpdateCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	GitProvider types.String `tfsdk:"git_provider"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername types.String `tfsdk:"git_username"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken types.String `tfsdk:"personal_access_token"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCredentialsRequest
// only implements ToObjectValue() and Type().
func (o UpdateCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":         o.CredentialId,
			"git_provider":          o.GitProvider,
			"git_username":          o.GitUsername,
			"personal_access_token": o.PersonalAccessToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCredentialsRequest) Type(ctx context.Context) attr.Type {
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

func (c UpdateCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCredentialsResponse
// only implements ToObjectValue() and Type().
func (o UpdateCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCredentialsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateRepoRequest struct {
	// Branch that the local version of the repo is checked out to.
	Branch types.String `tfsdk:"branch"`
	// ID of the Git folder (repo) object in the workspace.
	RepoId types.Int64 `tfsdk:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	SparseCheckout types.Object `tfsdk:"sparse_checkout"`
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	Tag types.String `tfsdk:"tag"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRepoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckoutUpdate{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRepoRequest
// only implements ToObjectValue() and Type().
func (o UpdateRepoRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":          o.Branch,
			"repo_id":         o.RepoId,
			"sparse_checkout": o.SparseCheckout,
			"tag":             o.Tag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRepoRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":          types.StringType,
			"repo_id":         types.Int64Type,
			"sparse_checkout": SparseCheckoutUpdate{}.Type(ctx),
			"tag":             types.StringType,
		},
	}
}

// GetSparseCheckout returns the value of the SparseCheckout field in UpdateRepoRequest as
// a SparseCheckoutUpdate value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRepoRequest) GetSparseCheckout(ctx context.Context) (SparseCheckoutUpdate, bool) {
	var e SparseCheckoutUpdate
	if o.SparseCheckout.IsNull() || o.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckoutUpdate
	d := o.SparseCheckout.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSparseCheckout sets the value of the SparseCheckout field in UpdateRepoRequest.
func (o *UpdateRepoRequest) SetSparseCheckout(ctx context.Context, v SparseCheckoutUpdate) {
	vs := v.ToObjectValue(ctx)
	o.SparseCheckout = vs
}

type UpdateRepoResponse struct {
}

func (newState *UpdateRepoResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRepoResponse) {
}

func (newState *UpdateRepoResponse) SyncEffectiveFieldsDuringRead(existingState UpdateRepoResponse) {
}

func (c UpdateRepoResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRepoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRepoResponse
// only implements ToObjectValue() and Type().
func (o UpdateRepoResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRepoResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type WorkspaceObjectAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *WorkspaceObjectAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectAccessControlRequest) {
}

func (newState *WorkspaceObjectAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectAccessControlRequest) {
}

func (c WorkspaceObjectAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceObjectAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceObjectAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectAccessControlRequest
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceObjectAccessControlRequest) Type(ctx context.Context) attr.Type {
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
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *WorkspaceObjectAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectAccessControlResponse) {
}

func (newState *WorkspaceObjectAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectAccessControlResponse) {
}

func (c WorkspaceObjectAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceObjectAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceObjectAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(WorkspaceObjectPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectAccessControlResponse
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceObjectAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: WorkspaceObjectPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in WorkspaceObjectAccessControlResponse as
// a slice of WorkspaceObjectPermission values.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceObjectAccessControlResponse) GetAllPermissions(ctx context.Context) ([]WorkspaceObjectPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in WorkspaceObjectAccessControlResponse.
func (o *WorkspaceObjectAccessControlResponse) SetAllPermissions(ctx context.Context, v []WorkspaceObjectPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type WorkspaceObjectPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *WorkspaceObjectPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectPermission) {
}

func (newState *WorkspaceObjectPermission) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectPermission) {
}

func (c WorkspaceObjectPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceObjectPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceObjectPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermission
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceObjectPermission) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in WorkspaceObjectPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceObjectPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in WorkspaceObjectPermission.
func (o *WorkspaceObjectPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type WorkspaceObjectPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *WorkspaceObjectPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectPermissions) {
}

func (newState *WorkspaceObjectPermissions) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectPermissions) {
}

func (c WorkspaceObjectPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceObjectPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceObjectPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WorkspaceObjectAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissions
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceObjectPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: WorkspaceObjectAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in WorkspaceObjectPermissions as
// a slice of WorkspaceObjectAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceObjectPermissions) GetAccessControlList(ctx context.Context) ([]WorkspaceObjectAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WorkspaceObjectPermissions.
func (o *WorkspaceObjectPermissions) SetAccessControlList(ctx context.Context, v []WorkspaceObjectAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type WorkspaceObjectPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *WorkspaceObjectPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceObjectPermissionsDescription) {
}

func (newState *WorkspaceObjectPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState WorkspaceObjectPermissionsDescription) {
}

func (c WorkspaceObjectPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceObjectPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceObjectPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissionsDescription
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceObjectPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type WorkspaceObjectPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceObjectPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspaceObjectPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WorkspaceObjectAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissionsRequest
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list":   o.AccessControlList,
			"workspace_object_id":   o.WorkspaceObjectId,
			"workspace_object_type": o.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceObjectPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: WorkspaceObjectAccessControlRequest{}.Type(ctx),
			},
			"workspace_object_id":   types.StringType,
			"workspace_object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in WorkspaceObjectPermissionsRequest as
// a slice of WorkspaceObjectAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *WorkspaceObjectPermissionsRequest) GetAccessControlList(ctx context.Context) ([]WorkspaceObjectAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WorkspaceObjectPermissionsRequest.
func (o *WorkspaceObjectPermissionsRequest) SetAccessControlList(ctx context.Context, v []WorkspaceObjectAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}
