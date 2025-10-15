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

// An item representing an ACL rule applied to the given principal (user or
// group) on the associated scope point.
type AclItem struct {
	// The permission level applied to the principal.
	Permission types.String `tfsdk:"permission"`
	// The principal in which the permission is applied.
	Principal types.String `tfsdk:"principal"`
}

func (to *AclItem) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AclItem) {
}

func (to *AclItem) SyncFieldsDuringRead(ctx context.Context, from AclItem) {
}

func (m AclItem) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AclItem) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AclItem
// only implements ToObjectValue() and Type().
func (m AclItem) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": m.Permission,
			"principal":  m.Principal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AclItem) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
			"principal":  types.StringType,
		},
	}
}

// The metadata of the Azure KeyVault for a secret scope of type
// `AZURE_KEYVAULT`
type AzureKeyVaultSecretScopeMetadata struct {
	// The DNS of the KeyVault
	DnsName types.String `tfsdk:"dns_name"`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	ResourceId types.String `tfsdk:"resource_id"`
}

func (to *AzureKeyVaultSecretScopeMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AzureKeyVaultSecretScopeMetadata) {
}

func (to *AzureKeyVaultSecretScopeMetadata) SyncFieldsDuringRead(ctx context.Context, from AzureKeyVaultSecretScopeMetadata) {
}

func (m AzureKeyVaultSecretScopeMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AzureKeyVaultSecretScopeMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureKeyVaultSecretScopeMetadata
// only implements ToObjectValue() and Type().
func (m AzureKeyVaultSecretScopeMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dns_name":    m.DnsName,
			"resource_id": m.ResourceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AzureKeyVaultSecretScopeMetadata) Type(ctx context.Context) attr.Type {
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
	// if the credential is the default for the given provider
	IsDefaultForProvider types.Bool `tfsdk:"is_default_for_provider"`
	// the name of the git credential, used for identification and ease of
	// lookup
	Name types.String `tfsdk:"name"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken types.String `tfsdk:"personal_access_token"`
}

func (to *CreateCredentialsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialsRequest) {
}

func (to *CreateCredentialsRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialsRequest) {
}

func (m CreateCredentialsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["git_provider"] = attrs["git_provider"].SetRequired()
	attrs["git_username"] = attrs["git_username"].SetOptional()
	attrs["is_default_for_provider"] = attrs["is_default_for_provider"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["personal_access_token"] = attrs["personal_access_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialsRequest
// only implements ToObjectValue() and Type().
func (m CreateCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"git_provider":            m.GitProvider,
			"git_username":            m.GitUsername,
			"is_default_for_provider": m.IsDefaultForProvider,
			"name":                    m.Name,
			"personal_access_token":   m.PersonalAccessToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCredentialsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"git_provider":            types.StringType,
			"git_username":            types.StringType,
			"is_default_for_provider": types.BoolType,
			"name":                    types.StringType,
			"personal_access_token":   types.StringType,
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
	// if the credential is the default for the given provider
	IsDefaultForProvider types.Bool `tfsdk:"is_default_for_provider"`
	// the name of the git credential, used for identification and ease of
	// lookup
	Name types.String `tfsdk:"name"`
}

func (to *CreateCredentialsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialsResponse) {
}

func (to *CreateCredentialsResponse) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialsResponse) {
}

func (m CreateCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetRequired()
	attrs["git_provider"] = attrs["git_provider"].SetRequired()
	attrs["git_username"] = attrs["git_username"].SetOptional()
	attrs["is_default_for_provider"] = attrs["is_default_for_provider"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialsResponse
// only implements ToObjectValue() and Type().
func (m CreateCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":           m.CredentialId,
			"git_provider":            m.GitProvider,
			"git_username":            m.GitUsername,
			"is_default_for_provider": m.IsDefaultForProvider,
			"name":                    m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCredentialsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id":           types.Int64Type,
			"git_provider":            types.StringType,
			"git_username":            types.StringType,
			"is_default_for_provider": types.BoolType,
			"name":                    types.StringType,
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

func (to *CreateRepoRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRepoRequest) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				// Recursively sync the fields of SparseCheckout
				toSparseCheckout.SyncFieldsDuringCreateOrUpdate(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (to *CreateRepoRequest) SyncFieldsDuringRead(ctx context.Context, from CreateRepoRequest) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				toSparseCheckout.SyncFieldsDuringRead(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (m CreateRepoRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetOptional()
	attrs["provider"] = attrs["provider"].SetRequired()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].SetOptional()
	attrs["url"] = attrs["url"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRepoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRepoRequest
// only implements ToObjectValue() and Type().
func (m CreateRepoRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path":            m.Path,
			"provider":        m.Provider,
			"sparse_checkout": m.SparseCheckout,
			"url":             m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRepoRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateRepoRequest) GetSparseCheckout(ctx context.Context) (SparseCheckout, bool) {
	var e SparseCheckout
	if m.SparseCheckout.IsNull() || m.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v SparseCheckout
	d := m.SparseCheckout.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparseCheckout sets the value of the SparseCheckout field in CreateRepoRequest.
func (m *CreateRepoRequest) SetSparseCheckout(ctx context.Context, v SparseCheckout) {
	vs := v.ToObjectValue(ctx)
	m.SparseCheckout = vs
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

func (to *CreateRepoResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRepoResponse) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				// Recursively sync the fields of SparseCheckout
				toSparseCheckout.SyncFieldsDuringCreateOrUpdate(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (to *CreateRepoResponse) SyncFieldsDuringRead(ctx context.Context, from CreateRepoResponse) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				toSparseCheckout.SyncFieldsDuringRead(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (m CreateRepoResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRepoResponse
// only implements ToObjectValue() and Type().
func (m CreateRepoResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":          m.Branch,
			"head_commit_id":  m.HeadCommitId,
			"id":              m.Id,
			"path":            m.Path,
			"provider":        m.Provider,
			"sparse_checkout": m.SparseCheckout,
			"url":             m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRepoResponse) Type(ctx context.Context) attr.Type {
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
func (m *CreateRepoResponse) GetSparseCheckout(ctx context.Context) (SparseCheckout, bool) {
	var e SparseCheckout
	if m.SparseCheckout.IsNull() || m.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v SparseCheckout
	d := m.SparseCheckout.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparseCheckout sets the value of the SparseCheckout field in CreateRepoResponse.
func (m *CreateRepoResponse) SetSparseCheckout(ctx context.Context, v SparseCheckout) {
	vs := v.ToObjectValue(ctx)
	m.SparseCheckout = vs
}

type CreateScope struct {
	// The metadata for the secret scope if the type is ``AZURE_KEYVAULT``
	BackendAzureKeyvault types.Object `tfsdk:"backend_azure_keyvault"`
	// The principal that is initially granted ``MANAGE`` permission to the
	// created scope.
	InitialManagePrincipal types.String `tfsdk:"initial_manage_principal"`
	// Scope name requested by the user. Scope names are unique.
	Scope types.String `tfsdk:"scope"`
	// The backend type the scope will be created with. If not specified, will
	// default to ``DATABRICKS``
	ScopeBackendType types.String `tfsdk:"scope_backend_type"`
}

func (to *CreateScope) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateScope) {
	if !from.BackendAzureKeyvault.IsNull() && !from.BackendAzureKeyvault.IsUnknown() {
		if toBackendAzureKeyvault, ok := to.GetBackendAzureKeyvault(ctx); ok {
			if fromBackendAzureKeyvault, ok := from.GetBackendAzureKeyvault(ctx); ok {
				// Recursively sync the fields of BackendAzureKeyvault
				toBackendAzureKeyvault.SyncFieldsDuringCreateOrUpdate(ctx, fromBackendAzureKeyvault)
				to.SetBackendAzureKeyvault(ctx, toBackendAzureKeyvault)
			}
		}
	}
}

func (to *CreateScope) SyncFieldsDuringRead(ctx context.Context, from CreateScope) {
	if !from.BackendAzureKeyvault.IsNull() && !from.BackendAzureKeyvault.IsUnknown() {
		if toBackendAzureKeyvault, ok := to.GetBackendAzureKeyvault(ctx); ok {
			if fromBackendAzureKeyvault, ok := from.GetBackendAzureKeyvault(ctx); ok {
				toBackendAzureKeyvault.SyncFieldsDuringRead(ctx, fromBackendAzureKeyvault)
				to.SetBackendAzureKeyvault(ctx, toBackendAzureKeyvault)
			}
		}
	}
}

func (m CreateScope) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["backend_azure_keyvault"] = attrs["backend_azure_keyvault"].SetOptional()
	attrs["initial_manage_principal"] = attrs["initial_manage_principal"].SetOptional()
	attrs["scope"] = attrs["scope"].SetRequired()
	attrs["scope_backend_type"] = attrs["scope_backend_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateScope.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateScope) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"backend_azure_keyvault": reflect.TypeOf(AzureKeyVaultSecretScopeMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateScope
// only implements ToObjectValue() and Type().
func (m CreateScope) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"backend_azure_keyvault":   m.BackendAzureKeyvault,
			"initial_manage_principal": m.InitialManagePrincipal,
			"scope":                    m.Scope,
			"scope_backend_type":       m.ScopeBackendType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateScope) Type(ctx context.Context) attr.Type {
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
func (m *CreateScope) GetBackendAzureKeyvault(ctx context.Context) (AzureKeyVaultSecretScopeMetadata, bool) {
	var e AzureKeyVaultSecretScopeMetadata
	if m.BackendAzureKeyvault.IsNull() || m.BackendAzureKeyvault.IsUnknown() {
		return e, false
	}
	var v AzureKeyVaultSecretScopeMetadata
	d := m.BackendAzureKeyvault.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBackendAzureKeyvault sets the value of the BackendAzureKeyvault field in CreateScope.
func (m *CreateScope) SetBackendAzureKeyvault(ctx context.Context, v AzureKeyVaultSecretScopeMetadata) {
	vs := v.ToObjectValue(ctx)
	m.BackendAzureKeyvault = vs
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	CredentialId types.Int64 `tfsdk:"credential_id"`
	// The Git provider associated with the credential.
	GitProvider types.String `tfsdk:"git_provider"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername types.String `tfsdk:"git_username"`
	// if the credential is the default for the given provider
	IsDefaultForProvider types.Bool `tfsdk:"is_default_for_provider"`
	// the name of the git credential, used for identification and ease of
	// lookup
	Name types.String `tfsdk:"name"`
}

func (to *CredentialInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CredentialInfo) {
}

func (to *CredentialInfo) SyncFieldsDuringRead(ctx context.Context, from CredentialInfo) {
}

func (m CredentialInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetRequired()
	attrs["git_provider"] = attrs["git_provider"].SetOptional()
	attrs["git_username"] = attrs["git_username"].SetOptional()
	attrs["is_default_for_provider"] = attrs["is_default_for_provider"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CredentialInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CredentialInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CredentialInfo
// only implements ToObjectValue() and Type().
func (m CredentialInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":           m.CredentialId,
			"git_provider":            m.GitProvider,
			"git_username":            m.GitUsername,
			"is_default_for_provider": m.IsDefaultForProvider,
			"name":                    m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CredentialInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id":           types.Int64Type,
			"git_provider":            types.StringType,
			"git_username":            types.StringType,
			"is_default_for_provider": types.BoolType,
			"name":                    types.StringType,
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

func (to *Delete) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Delete) {
}

func (to *Delete) SyncFieldsDuringRead(ctx context.Context, from Delete) {
}

func (m Delete) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()
	attrs["recursive"] = attrs["recursive"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Delete.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Delete) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Delete
// only implements ToObjectValue() and Type().
func (m Delete) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path":      m.Path,
			"recursive": m.Recursive,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Delete) Type(ctx context.Context) attr.Type {
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

func (to *DeleteAcl) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAcl) {
}

func (to *DeleteAcl) SyncFieldsDuringRead(ctx context.Context, from DeleteAcl) {
}

func (m DeleteAcl) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["principal"] = attrs["principal"].SetRequired()
	attrs["scope"] = attrs["scope"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAcl.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteAcl) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAcl
// only implements ToObjectValue() and Type().
func (m DeleteAcl) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal": m.Principal,
			"scope":     m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAcl) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"scope":     types.StringType,
		},
	}
}

type DeleteCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-"`
}

func (to *DeleteCredentialsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCredentialsRequest) {
}

func (to *DeleteCredentialsRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteCredentialsRequest) {
}

func (m DeleteCredentialsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialsRequest
// only implements ToObjectValue() and Type().
func (m DeleteCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": m.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCredentialsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.Int64Type,
		},
	}
}

type DeleteCredentialsResponse struct {
}

func (to *DeleteCredentialsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCredentialsResponse) {
}

func (to *DeleteCredentialsResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteCredentialsResponse) {
}

func (m DeleteCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialsResponse
// only implements ToObjectValue() and Type().
func (m DeleteCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCredentialsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRepoRequest struct {
	// The ID for the corresponding repo to delete.
	RepoId types.Int64 `tfsdk:"-"`
}

func (to *DeleteRepoRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRepoRequest) {
}

func (to *DeleteRepoRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteRepoRequest) {
}

func (m DeleteRepoRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["repo_id"] = attrs["repo_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRepoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRepoRequest
// only implements ToObjectValue() and Type().
func (m DeleteRepoRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": m.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRepoRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.Int64Type,
		},
	}
}

type DeleteRepoResponse struct {
}

func (to *DeleteRepoResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRepoResponse) {
}

func (to *DeleteRepoResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteRepoResponse) {
}

func (m DeleteRepoResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRepoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRepoResponse
// only implements ToObjectValue() and Type().
func (m DeleteRepoResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRepoResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteResponse struct {
}

func (to *DeleteResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteResponse) {
}

func (to *DeleteResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteResponse) {
}

func (m DeleteResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse
// only implements ToObjectValue() and Type().
func (m DeleteResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteScope struct {
	// Name of the scope to delete.
	Scope types.String `tfsdk:"scope"`
}

func (to *DeleteScope) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteScope) {
}

func (to *DeleteScope) SyncFieldsDuringRead(ctx context.Context, from DeleteScope) {
}

func (m DeleteScope) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["scope"] = attrs["scope"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteScope.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteScope) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteScope
// only implements ToObjectValue() and Type().
func (m DeleteScope) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteScope) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope": types.StringType,
		},
	}
}

type DeleteSecret struct {
	// Name of the secret to delete.
	Key types.String `tfsdk:"key"`
	// The name of the scope that contains the secret to delete.
	Scope types.String `tfsdk:"scope"`
}

func (to *DeleteSecret) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSecret) {
}

func (to *DeleteSecret) SyncFieldsDuringRead(ctx context.Context, from DeleteSecret) {
}

func (m DeleteSecret) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["scope"] = attrs["scope"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSecret.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSecret
// only implements ToObjectValue() and Type().
func (m DeleteSecret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"scope": m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSecret) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"scope": types.StringType,
		},
	}
}

type DeleteSecretResponse struct {
}

func (to *DeleteSecretResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSecretResponse) {
}

func (to *DeleteSecretResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteSecretResponse) {
}

func (m DeleteSecretResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSecretResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSecretResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSecretResponse
// only implements ToObjectValue() and Type().
func (m DeleteSecretResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSecretResponse) Type(ctx context.Context) attr.Type {
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
	// This specifies which cell outputs should be included in the export (if
	// the export format allows it). If not specified, the behavior is
	// determined by the format. For JUPYTER format, the default is to include
	// all outputs. This is a public endpoint, but only ALL or NONE is
	// documented publically, DATABRICKS is internal only
	Outputs types.String `tfsdk:"-"`
	// The absolute path of the object or directory. Exporting a directory is
	// only supported for the `DBC`, `SOURCE`, and `AUTO` format.
	Path types.String `tfsdk:"-"`
}

func (to *ExportRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExportRequest) {
}

func (to *ExportRequest) SyncFieldsDuringRead(ctx context.Context, from ExportRequest) {
}

func (m ExportRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()
	attrs["format"] = attrs["format"].SetOptional()
	attrs["outputs"] = attrs["outputs"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExportRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportRequest
// only implements ToObjectValue() and Type().
func (m ExportRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"format":  m.Format,
			"outputs": m.Outputs,
			"path":    m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExportRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"format":  types.StringType,
			"outputs": types.StringType,
			"path":    types.StringType,
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

func (to *ExportResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExportResponse) {
}

func (to *ExportResponse) SyncFieldsDuringRead(ctx context.Context, from ExportResponse) {
}

func (m ExportResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExportResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportResponse
// only implements ToObjectValue() and Type().
func (m ExportResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content":   m.Content,
			"file_type": m.FileType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExportResponse) Type(ctx context.Context) attr.Type {
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

func (to *GetAclRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAclRequest) {
}

func (to *GetAclRequest) SyncFieldsDuringRead(ctx context.Context, from GetAclRequest) {
}

func (m GetAclRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["scope"] = attrs["scope"].SetRequired()
	attrs["principal"] = attrs["principal"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAclRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAclRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAclRequest
// only implements ToObjectValue() and Type().
func (m GetAclRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal": m.Principal,
			"scope":     m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAclRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetCredentialsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCredentialsRequest) {
}

func (to *GetCredentialsRequest) SyncFieldsDuringRead(ctx context.Context, from GetCredentialsRequest) {
}

func (m GetCredentialsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialsRequest
// only implements ToObjectValue() and Type().
func (m GetCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": m.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCredentialsRequest) Type(ctx context.Context) attr.Type {
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
	// if the credential is the default for the given provider
	IsDefaultForProvider types.Bool `tfsdk:"is_default_for_provider"`
	// the name of the git credential, used for identification and ease of
	// lookup
	Name types.String `tfsdk:"name"`
}

func (to *GetCredentialsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCredentialsResponse) {
}

func (to *GetCredentialsResponse) SyncFieldsDuringRead(ctx context.Context, from GetCredentialsResponse) {
}

func (m GetCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential_id"] = attrs["credential_id"].SetRequired()
	attrs["git_provider"] = attrs["git_provider"].SetOptional()
	attrs["git_username"] = attrs["git_username"].SetOptional()
	attrs["is_default_for_provider"] = attrs["is_default_for_provider"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialsResponse
// only implements ToObjectValue() and Type().
func (m GetCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":           m.CredentialId,
			"git_provider":            m.GitProvider,
			"git_username":            m.GitUsername,
			"is_default_for_provider": m.IsDefaultForProvider,
			"name":                    m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCredentialsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id":           types.Int64Type,
			"git_provider":            types.StringType,
			"git_username":            types.StringType,
			"is_default_for_provider": types.BoolType,
			"name":                    types.StringType,
		},
	}
}

type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

func (to *GetRepoPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoPermissionLevelsRequest) {
}

func (to *GetRepoPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetRepoPermissionLevelsRequest) {
}

func (m GetRepoPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["repo_id"] = attrs["repo_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRepoPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRepoPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetRepoPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": m.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRepoPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetRepoPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetRepoPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetRepoPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetRepoPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRepoPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(RepoPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetRepoPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRepoPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
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
func (m *GetRepoPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]RepoPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []RepoPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetRepoPermissionLevelsResponse.
func (m *GetRepoPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []RepoPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetRepoPermissionsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

func (to *GetRepoPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoPermissionsRequest) {
}

func (to *GetRepoPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetRepoPermissionsRequest) {
}

func (m GetRepoPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["repo_id"] = attrs["repo_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRepoPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRepoPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetRepoPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": m.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRepoPermissionsRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetRepoRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoRequest) {
}

func (to *GetRepoRequest) SyncFieldsDuringRead(ctx context.Context, from GetRepoRequest) {
}

func (m GetRepoRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["repo_id"] = attrs["repo_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRepoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoRequest
// only implements ToObjectValue() and Type().
func (m GetRepoRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": m.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRepoRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetRepoResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoResponse) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				// Recursively sync the fields of SparseCheckout
				toSparseCheckout.SyncFieldsDuringCreateOrUpdate(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (to *GetRepoResponse) SyncFieldsDuringRead(ctx context.Context, from GetRepoResponse) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				toSparseCheckout.SyncFieldsDuringRead(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (m GetRepoResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoResponse
// only implements ToObjectValue() and Type().
func (m GetRepoResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":          m.Branch,
			"head_commit_id":  m.HeadCommitId,
			"id":              m.Id,
			"path":            m.Path,
			"provider":        m.Provider,
			"sparse_checkout": m.SparseCheckout,
			"url":             m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRepoResponse) Type(ctx context.Context) attr.Type {
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
func (m *GetRepoResponse) GetSparseCheckout(ctx context.Context) (SparseCheckout, bool) {
	var e SparseCheckout
	if m.SparseCheckout.IsNull() || m.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v SparseCheckout
	d := m.SparseCheckout.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparseCheckout sets the value of the SparseCheckout field in GetRepoResponse.
func (m *GetRepoResponse) SetSparseCheckout(ctx context.Context, v SparseCheckout) {
	vs := v.ToObjectValue(ctx)
	m.SparseCheckout = vs
}

type GetSecretRequest struct {
	// Name of the secret to fetch value information.
	Key types.String `tfsdk:"-"`
	// The name of the scope that contains the secret.
	Scope types.String `tfsdk:"-"`
}

func (to *GetSecretRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSecretRequest) {
}

func (to *GetSecretRequest) SyncFieldsDuringRead(ctx context.Context, from GetSecretRequest) {
}

func (m GetSecretRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["scope"] = attrs["scope"].SetRequired()
	attrs["key"] = attrs["key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSecretRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSecretRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSecretRequest
// only implements ToObjectValue() and Type().
func (m GetSecretRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"scope": m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSecretRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetSecretResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSecretResponse) {
}

func (to *GetSecretResponse) SyncFieldsDuringRead(ctx context.Context, from GetSecretResponse) {
}

func (m GetSecretResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetSecretResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSecretResponse
// only implements ToObjectValue() and Type().
func (m GetSecretResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSecretResponse) Type(ctx context.Context) attr.Type {
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

func (to *GetStatusRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatusRequest) {
}

func (to *GetStatusRequest) SyncFieldsDuringRead(ctx context.Context, from GetStatusRequest) {
}

func (m GetStatusRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatusRequest
// only implements ToObjectValue() and Type().
func (m GetStatusRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStatusRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetWorkspaceObjectPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceObjectPermissionLevelsRequest) {
}

func (to *GetWorkspaceObjectPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceObjectPermissionLevelsRequest) {
}

func (m GetWorkspaceObjectPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_object_type"] = attrs["workspace_object_type"].SetRequired()
	attrs["workspace_object_id"] = attrs["workspace_object_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceObjectPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceObjectPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceObjectPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_object_id":   m.WorkspaceObjectId,
			"workspace_object_type": m.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceObjectPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetWorkspaceObjectPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceObjectPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetWorkspaceObjectPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceObjectPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetWorkspaceObjectPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWorkspaceObjectPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(WorkspaceObjectPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetWorkspaceObjectPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceObjectPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
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
func (m *GetWorkspaceObjectPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]WorkspaceObjectPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetWorkspaceObjectPermissionLevelsResponse.
func (m *GetWorkspaceObjectPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []WorkspaceObjectPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetWorkspaceObjectPermissionsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-"`
}

func (to *GetWorkspaceObjectPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceObjectPermissionsRequest) {
}

func (to *GetWorkspaceObjectPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceObjectPermissionsRequest) {
}

func (m GetWorkspaceObjectPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_object_type"] = attrs["workspace_object_type"].SetRequired()
	attrs["workspace_object_id"] = attrs["workspace_object_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceObjectPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceObjectPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceObjectPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_object_id":   m.WorkspaceObjectId,
			"workspace_object_type": m.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceObjectPermissionsRequest) Type(ctx context.Context) attr.Type {
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

func (to *Import) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Import) {
}

func (to *Import) SyncFieldsDuringRead(ctx context.Context, from Import) {
}

func (m Import) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetOptional()
	attrs["format"] = attrs["format"].SetOptional()
	attrs["language"] = attrs["language"].SetOptional()
	attrs["overwrite"] = attrs["overwrite"].SetOptional()
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Import.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Import) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Import
// only implements ToObjectValue() and Type().
func (m Import) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content":   m.Content,
			"format":    m.Format,
			"language":  m.Language,
			"overwrite": m.Overwrite,
			"path":      m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Import) Type(ctx context.Context) attr.Type {
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

func (to *ImportResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ImportResponse) {
}

func (to *ImportResponse) SyncFieldsDuringRead(ctx context.Context, from ImportResponse) {
}

func (m ImportResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ImportResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ImportResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ImportResponse
// only implements ToObjectValue() and Type().
func (m ImportResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ImportResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListAclsRequest struct {
	// The name of the scope to fetch ACL information from.
	Scope types.String `tfsdk:"-"`
}

func (to *ListAclsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAclsRequest) {
}

func (to *ListAclsRequest) SyncFieldsDuringRead(ctx context.Context, from ListAclsRequest) {
}

func (m ListAclsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["scope"] = attrs["scope"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAclsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListAclsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAclsRequest
// only implements ToObjectValue() and Type().
func (m ListAclsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAclsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListAclsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAclsResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (to *ListAclsResponse) SyncFieldsDuringRead(ctx context.Context, from ListAclsResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (m ListAclsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAclsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(AclItem{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAclsResponse
// only implements ToObjectValue() and Type().
func (m ListAclsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items": m.Items,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAclsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListAclsResponse) GetItems(ctx context.Context) ([]AclItem, bool) {
	if m.Items.IsNull() || m.Items.IsUnknown() {
		return nil, false
	}
	var v []AclItem
	d := m.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListAclsResponse.
func (m *ListAclsResponse) SetItems(ctx context.Context, v []AclItem) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Items = types.ListValueMust(t, vs)
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

type ListCredentialsResponse struct {
	// List of credentials.
	Credentials types.List `tfsdk:"credentials"`
}

func (to *ListCredentialsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCredentialsResponse) {
	if !from.Credentials.IsNull() && !from.Credentials.IsUnknown() && to.Credentials.IsNull() && len(from.Credentials.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Credentials, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Credentials = from.Credentials
	}
}

func (to *ListCredentialsResponse) SyncFieldsDuringRead(ctx context.Context, from ListCredentialsResponse) {
	if !from.Credentials.IsNull() && !from.Credentials.IsUnknown() && to.Credentials.IsNull() && len(from.Credentials.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Credentials, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Credentials = from.Credentials
	}
}

func (m ListCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credentials": reflect.TypeOf(CredentialInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsResponse
// only implements ToObjectValue() and Type().
func (m ListCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials": m.Credentials,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCredentialsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListCredentialsResponse) GetCredentials(ctx context.Context) ([]CredentialInfo, bool) {
	if m.Credentials.IsNull() || m.Credentials.IsUnknown() {
		return nil, false
	}
	var v []CredentialInfo
	d := m.Credentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCredentials sets the value of the Credentials field in ListCredentialsResponse.
func (m *ListCredentialsResponse) SetCredentials(ctx context.Context, v []CredentialInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["credentials"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Credentials = types.ListValueMust(t, vs)
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

func (to *ListReposRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListReposRequest) {
}

func (to *ListReposRequest) SyncFieldsDuringRead(ctx context.Context, from ListReposRequest) {
}

func (m ListReposRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path_prefix"] = attrs["path_prefix"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListReposRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListReposRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListReposRequest
// only implements ToObjectValue() and Type().
func (m ListReposRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"path_prefix":     m.PathPrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListReposRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListReposResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListReposResponse) {
	if !from.Repos.IsNull() && !from.Repos.IsUnknown() && to.Repos.IsNull() && len(from.Repos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Repos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Repos = from.Repos
	}
}

func (to *ListReposResponse) SyncFieldsDuringRead(ctx context.Context, from ListReposResponse) {
	if !from.Repos.IsNull() && !from.Repos.IsUnknown() && to.Repos.IsNull() && len(from.Repos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Repos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Repos = from.Repos
	}
}

func (m ListReposResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListReposResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"repos": reflect.TypeOf(RepoInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListReposResponse
// only implements ToObjectValue() and Type().
func (m ListReposResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"repos":           m.Repos,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListReposResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListReposResponse) GetRepos(ctx context.Context) ([]RepoInfo, bool) {
	if m.Repos.IsNull() || m.Repos.IsUnknown() {
		return nil, false
	}
	var v []RepoInfo
	d := m.Repos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRepos sets the value of the Repos field in ListReposResponse.
func (m *ListReposResponse) SetRepos(ctx context.Context, v []RepoInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["repos"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Repos = types.ListValueMust(t, vs)
}

type ListResponse struct {
	// List of objects.
	Objects types.List `tfsdk:"objects"`
}

func (to *ListResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListResponse) {
	if !from.Objects.IsNull() && !from.Objects.IsUnknown() && to.Objects.IsNull() && len(from.Objects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Objects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Objects = from.Objects
	}
}

func (to *ListResponse) SyncFieldsDuringRead(ctx context.Context, from ListResponse) {
	if !from.Objects.IsNull() && !from.Objects.IsUnknown() && to.Objects.IsNull() && len(from.Objects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Objects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Objects = from.Objects
	}
}

func (m ListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objects": reflect.TypeOf(ObjectInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListResponse
// only implements ToObjectValue() and Type().
func (m ListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"objects": m.Objects,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListResponse) GetObjects(ctx context.Context) ([]ObjectInfo, bool) {
	if m.Objects.IsNull() || m.Objects.IsUnknown() {
		return nil, false
	}
	var v []ObjectInfo
	d := m.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in ListResponse.
func (m *ListResponse) SetObjects(ctx context.Context, v []ObjectInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Objects = types.ListValueMust(t, vs)
}

type ListScopesRequest struct {
}

func (to *ListScopesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListScopesRequest) {
}

func (to *ListScopesRequest) SyncFieldsDuringRead(ctx context.Context, from ListScopesRequest) {
}

func (m ListScopesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListScopesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListScopesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListScopesRequest
// only implements ToObjectValue() and Type().
func (m ListScopesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListScopesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListScopesResponse struct {
	// The available secret scopes.
	Scopes types.List `tfsdk:"scopes"`
}

func (to *ListScopesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListScopesResponse) {
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
}

func (to *ListScopesResponse) SyncFieldsDuringRead(ctx context.Context, from ListScopesResponse) {
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
}

func (m ListScopesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListScopesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(SecretScope{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListScopesResponse
// only implements ToObjectValue() and Type().
func (m ListScopesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scopes": m.Scopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListScopesResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListScopesResponse) GetScopes(ctx context.Context) ([]SecretScope, bool) {
	if m.Scopes.IsNull() || m.Scopes.IsUnknown() {
		return nil, false
	}
	var v []SecretScope
	d := m.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in ListScopesResponse.
func (m *ListScopesResponse) SetScopes(ctx context.Context, v []SecretScope) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	Scope types.String `tfsdk:"-"`
}

func (to *ListSecretsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSecretsRequest) {
}

func (to *ListSecretsRequest) SyncFieldsDuringRead(ctx context.Context, from ListSecretsRequest) {
}

func (m ListSecretsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["scope"] = attrs["scope"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSecretsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSecretsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSecretsRequest
// only implements ToObjectValue() and Type().
func (m ListSecretsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSecretsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListSecretsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSecretsResponse) {
	if !from.Secrets.IsNull() && !from.Secrets.IsUnknown() && to.Secrets.IsNull() && len(from.Secrets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Secrets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Secrets = from.Secrets
	}
}

func (to *ListSecretsResponse) SyncFieldsDuringRead(ctx context.Context, from ListSecretsResponse) {
	if !from.Secrets.IsNull() && !from.Secrets.IsUnknown() && to.Secrets.IsNull() && len(from.Secrets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Secrets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Secrets = from.Secrets
	}
}

func (m ListSecretsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSecretsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"secrets": reflect.TypeOf(SecretMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSecretsResponse
// only implements ToObjectValue() and Type().
func (m ListSecretsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"secrets": m.Secrets,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSecretsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListSecretsResponse) GetSecrets(ctx context.Context) ([]SecretMetadata, bool) {
	if m.Secrets.IsNull() || m.Secrets.IsUnknown() {
		return nil, false
	}
	var v []SecretMetadata
	d := m.Secrets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecrets sets the value of the Secrets field in ListSecretsResponse.
func (m *ListSecretsResponse) SetSecrets(ctx context.Context, v []SecretMetadata) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["secrets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Secrets = types.ListValueMust(t, vs)
}

type ListWorkspaceRequest struct {
	// UTC timestamp in milliseconds
	NotebooksModifiedAfter types.Int64 `tfsdk:"-"`
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"-"`
}

func (to *ListWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceRequest) {
}

func (to *ListWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceRequest) {
}

func (m ListWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()
	attrs["notebooks_modified_after"] = attrs["notebooks_modified_after"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceRequest
// only implements ToObjectValue() and Type().
func (m ListWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notebooks_modified_after": m.NotebooksModifiedAfter,
			"path":                     m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceRequest) Type(ctx context.Context) attr.Type {
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

func (to *Mkdirs) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Mkdirs) {
}

func (to *Mkdirs) SyncFieldsDuringRead(ctx context.Context, from Mkdirs) {
}

func (m Mkdirs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Mkdirs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Mkdirs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Mkdirs
// only implements ToObjectValue() and Type().
func (m Mkdirs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Mkdirs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type MkdirsResponse struct {
}

func (to *MkdirsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MkdirsResponse) {
}

func (to *MkdirsResponse) SyncFieldsDuringRead(ctx context.Context, from MkdirsResponse) {
}

func (m MkdirsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MkdirsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MkdirsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MkdirsResponse
// only implements ToObjectValue() and Type().
func (m MkdirsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m MkdirsResponse) Type(ctx context.Context) attr.Type {
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

func (to *ObjectInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ObjectInfo) {
}

func (to *ObjectInfo) SyncFieldsDuringRead(ctx context.Context, from ObjectInfo) {
}

func (m ObjectInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ObjectInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ObjectInfo
// only implements ToObjectValue() and Type().
func (m ObjectInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_at":  m.CreatedAt,
			"language":    m.Language,
			"modified_at": m.ModifiedAt,
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,
			"path":        m.Path,
			"resource_id": m.ResourceId,
			"size":        m.Size,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ObjectInfo) Type(ctx context.Context) attr.Type {
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

func (to *PutAcl) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutAcl) {
}

func (to *PutAcl) SyncFieldsDuringRead(ctx context.Context, from PutAcl) {
}

func (m PutAcl) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()
	attrs["principal"] = attrs["principal"].SetRequired()
	attrs["scope"] = attrs["scope"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutAcl.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PutAcl) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAcl
// only implements ToObjectValue() and Type().
func (m PutAcl) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": m.Permission,
			"principal":  m.Principal,
			"scope":      m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PutAcl) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
			"principal":  types.StringType,
			"scope":      types.StringType,
		},
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

func (to *PutSecret) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutSecret) {
}

func (to *PutSecret) SyncFieldsDuringRead(ctx context.Context, from PutSecret) {
}

func (m PutSecret) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bytes_value"] = attrs["bytes_value"].SetOptional()
	attrs["key"] = attrs["key"].SetRequired()
	attrs["scope"] = attrs["scope"].SetRequired()
	attrs["string_value"] = attrs["string_value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutSecret.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PutSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutSecret
// only implements ToObjectValue() and Type().
func (m PutSecret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bytes_value":  m.BytesValue,
			"key":          m.Key,
			"scope":        m.Scope,
			"string_value": m.StringValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PutSecret) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bytes_value":  types.StringType,
			"key":          types.StringType,
			"scope":        types.StringType,
			"string_value": types.StringType,
		},
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

func (to *RepoAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoAccessControlRequest) {
}

func (to *RepoAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from RepoAccessControlRequest) {
}

func (m RepoAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoAccessControlRequest
// only implements ToObjectValue() and Type().
func (m RepoAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoAccessControlRequest) Type(ctx context.Context) attr.Type {
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

func (to *RepoAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *RepoAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from RepoAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m RepoAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(RepoPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoAccessControlResponse
// only implements ToObjectValue() and Type().
func (m RepoAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoAccessControlResponse) Type(ctx context.Context) attr.Type {
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
func (m *RepoAccessControlResponse) GetAllPermissions(ctx context.Context) ([]RepoPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []RepoPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in RepoAccessControlResponse.
func (m *RepoAccessControlResponse) SetAllPermissions(ctx context.Context, v []RepoPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
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

func (to *RepoInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoInfo) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				// Recursively sync the fields of SparseCheckout
				toSparseCheckout.SyncFieldsDuringCreateOrUpdate(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (to *RepoInfo) SyncFieldsDuringRead(ctx context.Context, from RepoInfo) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				toSparseCheckout.SyncFieldsDuringRead(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (m RepoInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoInfo
// only implements ToObjectValue() and Type().
func (m RepoInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":          m.Branch,
			"head_commit_id":  m.HeadCommitId,
			"id":              m.Id,
			"path":            m.Path,
			"provider":        m.Provider,
			"sparse_checkout": m.SparseCheckout,
			"url":             m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoInfo) Type(ctx context.Context) attr.Type {
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
func (m *RepoInfo) GetSparseCheckout(ctx context.Context) (SparseCheckout, bool) {
	var e SparseCheckout
	if m.SparseCheckout.IsNull() || m.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v SparseCheckout
	d := m.SparseCheckout.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparseCheckout sets the value of the SparseCheckout field in RepoInfo.
func (m *RepoInfo) SetSparseCheckout(ctx context.Context, v SparseCheckout) {
	vs := v.ToObjectValue(ctx)
	m.SparseCheckout = vs
}

type RepoPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *RepoPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *RepoPermission) SyncFieldsDuringRead(ctx context.Context, from RepoPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m RepoPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermission
// only implements ToObjectValue() and Type().
func (m RepoPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoPermission) Type(ctx context.Context) attr.Type {
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
func (m *RepoPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in RepoPermission.
func (m *RepoPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type RepoPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *RepoPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *RepoPermissions) SyncFieldsDuringRead(ctx context.Context, from RepoPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m RepoPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RepoAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissions
// only implements ToObjectValue() and Type().
func (m RepoPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoPermissions) Type(ctx context.Context) attr.Type {
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
func (m *RepoPermissions) GetAccessControlList(ctx context.Context) ([]RepoAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RepoAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RepoPermissions.
func (m *RepoPermissions) SetAccessControlList(ctx context.Context, v []RepoAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type RepoPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *RepoPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoPermissionsDescription) {
}

func (to *RepoPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from RepoPermissionsDescription) {
}

func (m RepoPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissionsDescription
// only implements ToObjectValue() and Type().
func (m RepoPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoPermissionsDescription) Type(ctx context.Context) attr.Type {
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

func (to *RepoPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *RepoPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from RepoPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m RepoPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["repo_id"] = attrs["repo_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RepoPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RepoPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RepoAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissionsRequest
// only implements ToObjectValue() and Type().
func (m RepoPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"repo_id":             m.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoPermissionsRequest) Type(ctx context.Context) attr.Type {
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
func (m *RepoPermissionsRequest) GetAccessControlList(ctx context.Context) ([]RepoAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RepoAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RepoPermissionsRequest.
func (m *RepoPermissionsRequest) SetAccessControlList(ctx context.Context, v []RepoAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

// The metadata about a secret. Returned when listing secrets. Does not contain
// the actual secret value.
type SecretMetadata struct {
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key"`
	// The last updated timestamp (in milliseconds) for the secret.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
}

func (to *SecretMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SecretMetadata) {
}

func (to *SecretMetadata) SyncFieldsDuringRead(ctx context.Context, from SecretMetadata) {
}

func (m SecretMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SecretMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretMetadata
// only implements ToObjectValue() and Type().
func (m SecretMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":                    m.Key,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SecretMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":                    types.StringType,
			"last_updated_timestamp": types.Int64Type,
		},
	}
}

// An organizational resource for storing secrets. Secret scopes can be
// different types (Databricks-managed, Azure KeyVault backed, etc), and ACLs
// can be applied to control permissions for all secrets within a scope.
type SecretScope struct {
	// The type of secret scope backend.
	BackendType types.String `tfsdk:"backend_type"`
	// The metadata for the secret scope if the type is ``AZURE_KEYVAULT``
	KeyvaultMetadata types.Object `tfsdk:"keyvault_metadata"`
	// A unique name to identify the secret scope.
	Name types.String `tfsdk:"name"`
}

func (to *SecretScope) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SecretScope) {
	if !from.KeyvaultMetadata.IsNull() && !from.KeyvaultMetadata.IsUnknown() {
		if toKeyvaultMetadata, ok := to.GetKeyvaultMetadata(ctx); ok {
			if fromKeyvaultMetadata, ok := from.GetKeyvaultMetadata(ctx); ok {
				// Recursively sync the fields of KeyvaultMetadata
				toKeyvaultMetadata.SyncFieldsDuringCreateOrUpdate(ctx, fromKeyvaultMetadata)
				to.SetKeyvaultMetadata(ctx, toKeyvaultMetadata)
			}
		}
	}
}

func (to *SecretScope) SyncFieldsDuringRead(ctx context.Context, from SecretScope) {
	if !from.KeyvaultMetadata.IsNull() && !from.KeyvaultMetadata.IsUnknown() {
		if toKeyvaultMetadata, ok := to.GetKeyvaultMetadata(ctx); ok {
			if fromKeyvaultMetadata, ok := from.GetKeyvaultMetadata(ctx); ok {
				toKeyvaultMetadata.SyncFieldsDuringRead(ctx, fromKeyvaultMetadata)
				to.SetKeyvaultMetadata(ctx, toKeyvaultMetadata)
			}
		}
	}
}

func (m SecretScope) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SecretScope) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"keyvault_metadata": reflect.TypeOf(AzureKeyVaultSecretScopeMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretScope
// only implements ToObjectValue() and Type().
func (m SecretScope) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"backend_type":      m.BackendType,
			"keyvault_metadata": m.KeyvaultMetadata,
			"name":              m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SecretScope) Type(ctx context.Context) attr.Type {
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
func (m *SecretScope) GetKeyvaultMetadata(ctx context.Context) (AzureKeyVaultSecretScopeMetadata, bool) {
	var e AzureKeyVaultSecretScopeMetadata
	if m.KeyvaultMetadata.IsNull() || m.KeyvaultMetadata.IsUnknown() {
		return e, false
	}
	var v AzureKeyVaultSecretScopeMetadata
	d := m.KeyvaultMetadata.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKeyvaultMetadata sets the value of the KeyvaultMetadata field in SecretScope.
func (m *SecretScope) SetKeyvaultMetadata(ctx context.Context, v AzureKeyVaultSecretScopeMetadata) {
	vs := v.ToObjectValue(ctx)
	m.KeyvaultMetadata = vs
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckout struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns types.List `tfsdk:"patterns"`
}

func (to *SparseCheckout) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SparseCheckout) {
	if !from.Patterns.IsNull() && !from.Patterns.IsUnknown() && to.Patterns.IsNull() && len(from.Patterns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Patterns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Patterns = from.Patterns
	}
}

func (to *SparseCheckout) SyncFieldsDuringRead(ctx context.Context, from SparseCheckout) {
	if !from.Patterns.IsNull() && !from.Patterns.IsUnknown() && to.Patterns.IsNull() && len(from.Patterns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Patterns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Patterns = from.Patterns
	}
}

func (m SparseCheckout) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SparseCheckout) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"patterns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparseCheckout
// only implements ToObjectValue() and Type().
func (m SparseCheckout) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"patterns": m.Patterns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SparseCheckout) Type(ctx context.Context) attr.Type {
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
func (m *SparseCheckout) GetPatterns(ctx context.Context) ([]types.String, bool) {
	if m.Patterns.IsNull() || m.Patterns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Patterns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPatterns sets the value of the Patterns field in SparseCheckout.
func (m *SparseCheckout) SetPatterns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["patterns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Patterns = types.ListValueMust(t, vs)
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckoutUpdate struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns types.List `tfsdk:"patterns"`
}

func (to *SparseCheckoutUpdate) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SparseCheckoutUpdate) {
	if !from.Patterns.IsNull() && !from.Patterns.IsUnknown() && to.Patterns.IsNull() && len(from.Patterns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Patterns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Patterns = from.Patterns
	}
}

func (to *SparseCheckoutUpdate) SyncFieldsDuringRead(ctx context.Context, from SparseCheckoutUpdate) {
	if !from.Patterns.IsNull() && !from.Patterns.IsUnknown() && to.Patterns.IsNull() && len(from.Patterns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Patterns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Patterns = from.Patterns
	}
}

func (m SparseCheckoutUpdate) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SparseCheckoutUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"patterns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparseCheckoutUpdate
// only implements ToObjectValue() and Type().
func (m SparseCheckoutUpdate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"patterns": m.Patterns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SparseCheckoutUpdate) Type(ctx context.Context) attr.Type {
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
func (m *SparseCheckoutUpdate) GetPatterns(ctx context.Context) ([]types.String, bool) {
	if m.Patterns.IsNull() || m.Patterns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Patterns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPatterns sets the value of the Patterns field in SparseCheckoutUpdate.
func (m *SparseCheckoutUpdate) SetPatterns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["patterns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Patterns = types.ListValueMust(t, vs)
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
	// if the credential is the default for the given provider
	IsDefaultForProvider types.Bool `tfsdk:"is_default_for_provider"`
	// the name of the git credential, used for identification and ease of
	// lookup
	Name types.String `tfsdk:"name"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken types.String `tfsdk:"personal_access_token"`
}

func (to *UpdateCredentialsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCredentialsRequest) {
}

func (to *UpdateCredentialsRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateCredentialsRequest) {
}

func (m UpdateCredentialsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["git_provider"] = attrs["git_provider"].SetRequired()
	attrs["git_username"] = attrs["git_username"].SetOptional()
	attrs["is_default_for_provider"] = attrs["is_default_for_provider"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["personal_access_token"] = attrs["personal_access_token"].SetOptional()
	attrs["credential_id"] = attrs["credential_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCredentialsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCredentialsRequest
// only implements ToObjectValue() and Type().
func (m UpdateCredentialsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":           m.CredentialId,
			"git_provider":            m.GitProvider,
			"git_username":            m.GitUsername,
			"is_default_for_provider": m.IsDefaultForProvider,
			"name":                    m.Name,
			"personal_access_token":   m.PersonalAccessToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCredentialsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id":           types.Int64Type,
			"git_provider":            types.StringType,
			"git_username":            types.StringType,
			"is_default_for_provider": types.BoolType,
			"name":                    types.StringType,
			"personal_access_token":   types.StringType,
		},
	}
}

type UpdateCredentialsResponse struct {
}

func (to *UpdateCredentialsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCredentialsResponse) {
}

func (to *UpdateCredentialsResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateCredentialsResponse) {
}

func (m UpdateCredentialsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCredentialsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCredentialsResponse
// only implements ToObjectValue() and Type().
func (m UpdateCredentialsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCredentialsResponse) Type(ctx context.Context) attr.Type {
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

func (to *UpdateRepoRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRepoRequest) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				// Recursively sync the fields of SparseCheckout
				toSparseCheckout.SyncFieldsDuringCreateOrUpdate(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (to *UpdateRepoRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateRepoRequest) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				toSparseCheckout.SyncFieldsDuringRead(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (m UpdateRepoRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].SetOptional()
	attrs["tag"] = attrs["tag"].SetOptional()
	attrs["repo_id"] = attrs["repo_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRepoRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRepoRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckoutUpdate{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRepoRequest
// only implements ToObjectValue() and Type().
func (m UpdateRepoRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":          m.Branch,
			"repo_id":         m.RepoId,
			"sparse_checkout": m.SparseCheckout,
			"tag":             m.Tag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRepoRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateRepoRequest) GetSparseCheckout(ctx context.Context) (SparseCheckoutUpdate, bool) {
	var e SparseCheckoutUpdate
	if m.SparseCheckout.IsNull() || m.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v SparseCheckoutUpdate
	d := m.SparseCheckout.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSparseCheckout sets the value of the SparseCheckout field in UpdateRepoRequest.
func (m *UpdateRepoRequest) SetSparseCheckout(ctx context.Context, v SparseCheckoutUpdate) {
	vs := v.ToObjectValue(ctx)
	m.SparseCheckout = vs
}

type UpdateRepoResponse struct {
}

func (to *UpdateRepoResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRepoResponse) {
}

func (to *UpdateRepoResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateRepoResponse) {
}

func (m UpdateRepoResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRepoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRepoResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRepoResponse
// only implements ToObjectValue() and Type().
func (m UpdateRepoResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRepoResponse) Type(ctx context.Context) attr.Type {
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

func (to *WorkspaceObjectAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectAccessControlRequest) {
}

func (to *WorkspaceObjectAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectAccessControlRequest) {
}

func (m WorkspaceObjectAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectAccessControlRequest
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceObjectAccessControlRequest) Type(ctx context.Context) attr.Type {
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

func (to *WorkspaceObjectAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *WorkspaceObjectAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m WorkspaceObjectAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(WorkspaceObjectPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectAccessControlResponse
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceObjectAccessControlResponse) Type(ctx context.Context) attr.Type {
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
func (m *WorkspaceObjectAccessControlResponse) GetAllPermissions(ctx context.Context) ([]WorkspaceObjectPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in WorkspaceObjectAccessControlResponse.
func (m *WorkspaceObjectAccessControlResponse) SetAllPermissions(ctx context.Context, v []WorkspaceObjectPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type WorkspaceObjectPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *WorkspaceObjectPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *WorkspaceObjectPermission) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m WorkspaceObjectPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermission
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceObjectPermission) Type(ctx context.Context) attr.Type {
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
func (m *WorkspaceObjectPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in WorkspaceObjectPermission.
func (m *WorkspaceObjectPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type WorkspaceObjectPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *WorkspaceObjectPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *WorkspaceObjectPermissions) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m WorkspaceObjectPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WorkspaceObjectAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissions
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceObjectPermissions) Type(ctx context.Context) attr.Type {
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
func (m *WorkspaceObjectPermissions) GetAccessControlList(ctx context.Context) ([]WorkspaceObjectAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WorkspaceObjectPermissions.
func (m *WorkspaceObjectPermissions) SetAccessControlList(ctx context.Context, v []WorkspaceObjectAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type WorkspaceObjectPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *WorkspaceObjectPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectPermissionsDescription) {
}

func (to *WorkspaceObjectPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectPermissionsDescription) {
}

func (m WorkspaceObjectPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissionsDescription
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceObjectPermissionsDescription) Type(ctx context.Context) attr.Type {
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

func (to *WorkspaceObjectPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *WorkspaceObjectPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m WorkspaceObjectPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["workspace_object_type"] = attrs["workspace_object_type"].SetRequired()
	attrs["workspace_object_id"] = attrs["workspace_object_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceObjectPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceObjectPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WorkspaceObjectAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissionsRequest
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list":   m.AccessControlList,
			"workspace_object_id":   m.WorkspaceObjectId,
			"workspace_object_type": m.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceObjectPermissionsRequest) Type(ctx context.Context) attr.Type {
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
func (m *WorkspaceObjectPermissionsRequest) GetAccessControlList(ctx context.Context) ([]WorkspaceObjectAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WorkspaceObjectPermissionsRequest.
func (m *WorkspaceObjectPermissionsRequest) SetAccessControlList(ctx context.Context, v []WorkspaceObjectAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}
