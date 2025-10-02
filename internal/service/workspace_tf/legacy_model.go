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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// An item representing an ACL rule applied to the given principal (user or
// group) on the associated scope point.
type AclItem_SdkV2 struct {
	// The permission level applied to the principal.
	Permission types.String `tfsdk:"permission"`
	// The principal in which the permission is applied.
	Principal types.String `tfsdk:"principal"`
}

func (to *AclItem_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AclItem_SdkV2) {
}

func (to *AclItem_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AclItem_SdkV2) {
}

func (m AclItem_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AclItem_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AclItem_SdkV2
// only implements ToObjectValue() and Type().
func (m AclItem_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": m.Permission,
			"principal":  m.Principal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AclItem_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
			"principal":  types.StringType,
		},
	}
}

// The metadata of the Azure KeyVault for a secret scope of type
// `AZURE_KEYVAULT`
type AzureKeyVaultSecretScopeMetadata_SdkV2 struct {
	// The DNS of the KeyVault
	DnsName types.String `tfsdk:"dns_name"`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	ResourceId types.String `tfsdk:"resource_id"`
}

func (to *AzureKeyVaultSecretScopeMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AzureKeyVaultSecretScopeMetadata_SdkV2) {
}

func (to *AzureKeyVaultSecretScopeMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AzureKeyVaultSecretScopeMetadata_SdkV2) {
}

func (m AzureKeyVaultSecretScopeMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AzureKeyVaultSecretScopeMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureKeyVaultSecretScopeMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m AzureKeyVaultSecretScopeMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dns_name":    m.DnsName,
			"resource_id": m.ResourceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AzureKeyVaultSecretScopeMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dns_name":    types.StringType,
			"resource_id": types.StringType,
		},
	}
}

type CreateCredentialsRequest_SdkV2 struct {
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

func (to *CreateCredentialsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialsRequest_SdkV2) {
}

func (to *CreateCredentialsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialsRequest_SdkV2) {
}

func (m CreateCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type CreateCredentialsResponse_SdkV2 struct {
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

func (to *CreateCredentialsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCredentialsResponse_SdkV2) {
}

func (to *CreateCredentialsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCredentialsResponse_SdkV2) {
}

func (m CreateCredentialsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

type CreateRepoRequest_SdkV2 struct {
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
	SparseCheckout types.List `tfsdk:"sparse_checkout"`
	// URL of the Git repository to be linked.
	Url types.String `tfsdk:"url"`
}

func (to *CreateRepoRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRepoRequest_SdkV2) {
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

func (to *CreateRepoRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateRepoRequest_SdkV2) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				toSparseCheckout.SyncFieldsDuringRead(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (m CreateRepoRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetOptional()
	attrs["provider"] = attrs["provider"].SetRequired()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateRepoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRepoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateRepoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateRepoRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":     types.StringType,
			"provider": types.StringType,
			"sparse_checkout": basetypes.ListType{
				ElemType: SparseCheckout_SdkV2{}.Type(ctx),
			},
			"url": types.StringType,
		},
	}
}

// GetSparseCheckout returns the value of the SparseCheckout field in CreateRepoRequest_SdkV2 as
// a SparseCheckout_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRepoRequest_SdkV2) GetSparseCheckout(ctx context.Context) (SparseCheckout_SdkV2, bool) {
	var e SparseCheckout_SdkV2
	if m.SparseCheckout.IsNull() || m.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout_SdkV2
	d := m.SparseCheckout.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparseCheckout sets the value of the SparseCheckout field in CreateRepoRequest_SdkV2.
func (m *CreateRepoRequest_SdkV2) SetSparseCheckout(ctx context.Context, v SparseCheckout_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sparse_checkout"]
	m.SparseCheckout = types.ListValueMust(t, vs)
}

type CreateRepoResponse_SdkV2 struct {
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
	SparseCheckout types.List `tfsdk:"sparse_checkout"`
	// URL of the linked Git repository.
	Url types.String `tfsdk:"url"`
}

func (to *CreateRepoResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRepoResponse_SdkV2) {
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

func (to *CreateRepoResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateRepoResponse_SdkV2) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				toSparseCheckout.SyncFieldsDuringRead(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (m CreateRepoResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetOptional()
	attrs["head_commit_id"] = attrs["head_commit_id"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()
	attrs["provider"] = attrs["provider"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateRepoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRepoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateRepoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateRepoResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":         types.StringType,
			"head_commit_id": types.StringType,
			"id":             types.Int64Type,
			"path":           types.StringType,
			"provider":       types.StringType,
			"sparse_checkout": basetypes.ListType{
				ElemType: SparseCheckout_SdkV2{}.Type(ctx),
			},
			"url": types.StringType,
		},
	}
}

// GetSparseCheckout returns the value of the SparseCheckout field in CreateRepoResponse_SdkV2 as
// a SparseCheckout_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRepoResponse_SdkV2) GetSparseCheckout(ctx context.Context) (SparseCheckout_SdkV2, bool) {
	var e SparseCheckout_SdkV2
	if m.SparseCheckout.IsNull() || m.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout_SdkV2
	d := m.SparseCheckout.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparseCheckout sets the value of the SparseCheckout field in CreateRepoResponse_SdkV2.
func (m *CreateRepoResponse_SdkV2) SetSparseCheckout(ctx context.Context, v SparseCheckout_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sparse_checkout"]
	m.SparseCheckout = types.ListValueMust(t, vs)
}

type CreateScope_SdkV2 struct {
	// The metadata for the secret scope if the type is ``AZURE_KEYVAULT``
	BackendAzureKeyvault types.List `tfsdk:"backend_azure_keyvault"`
	// The principal that is initially granted ``MANAGE`` permission to the
	// created scope.
	InitialManagePrincipal types.String `tfsdk:"initial_manage_principal"`
	// Scope name requested by the user. Scope names are unique.
	Scope types.String `tfsdk:"scope"`
	// The backend type the scope will be created with. If not specified, will
	// default to ``DATABRICKS``
	ScopeBackendType types.String `tfsdk:"scope_backend_type"`
}

func (to *CreateScope_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateScope_SdkV2) {
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

func (to *CreateScope_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateScope_SdkV2) {
	if !from.BackendAzureKeyvault.IsNull() && !from.BackendAzureKeyvault.IsUnknown() {
		if toBackendAzureKeyvault, ok := to.GetBackendAzureKeyvault(ctx); ok {
			if fromBackendAzureKeyvault, ok := from.GetBackendAzureKeyvault(ctx); ok {
				toBackendAzureKeyvault.SyncFieldsDuringRead(ctx, fromBackendAzureKeyvault)
				to.SetBackendAzureKeyvault(ctx, toBackendAzureKeyvault)
			}
		}
	}
}

func (m CreateScope_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["backend_azure_keyvault"] = attrs["backend_azure_keyvault"].SetOptional()
	attrs["backend_azure_keyvault"] = attrs["backend_azure_keyvault"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateScope_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"backend_azure_keyvault": reflect.TypeOf(AzureKeyVaultSecretScopeMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateScope_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateScope_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateScope_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"backend_azure_keyvault": basetypes.ListType{
				ElemType: AzureKeyVaultSecretScopeMetadata_SdkV2{}.Type(ctx),
			},
			"initial_manage_principal": types.StringType,
			"scope":                    types.StringType,
			"scope_backend_type":       types.StringType,
		},
	}
}

// GetBackendAzureKeyvault returns the value of the BackendAzureKeyvault field in CreateScope_SdkV2 as
// a AzureKeyVaultSecretScopeMetadata_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateScope_SdkV2) GetBackendAzureKeyvault(ctx context.Context) (AzureKeyVaultSecretScopeMetadata_SdkV2, bool) {
	var e AzureKeyVaultSecretScopeMetadata_SdkV2
	if m.BackendAzureKeyvault.IsNull() || m.BackendAzureKeyvault.IsUnknown() {
		return e, false
	}
	var v []AzureKeyVaultSecretScopeMetadata_SdkV2
	d := m.BackendAzureKeyvault.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBackendAzureKeyvault sets the value of the BackendAzureKeyvault field in CreateScope_SdkV2.
func (m *CreateScope_SdkV2) SetBackendAzureKeyvault(ctx context.Context, v AzureKeyVaultSecretScopeMetadata_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["backend_azure_keyvault"]
	m.BackendAzureKeyvault = types.ListValueMust(t, vs)
}

type CredentialInfo_SdkV2 struct {
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

func (to *CredentialInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CredentialInfo_SdkV2) {
}

func (to *CredentialInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CredentialInfo_SdkV2) {
}

func (m CredentialInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CredentialInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CredentialInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m CredentialInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CredentialInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

type Delete_SdkV2 struct {
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive types.Bool `tfsdk:"recursive"`
}

func (to *Delete_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Delete_SdkV2) {
}

func (to *Delete_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Delete_SdkV2) {
}

func (m Delete_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Delete_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Delete_SdkV2
// only implements ToObjectValue() and Type().
func (m Delete_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path":      m.Path,
			"recursive": m.Recursive,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Delete_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":      types.StringType,
			"recursive": types.BoolType,
		},
	}
}

type DeleteAcl_SdkV2 struct {
	// The principal to remove an existing ACL from.
	Principal types.String `tfsdk:"principal"`
	// The name of the scope to remove permissions from.
	Scope types.String `tfsdk:"scope"`
}

func (to *DeleteAcl_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAcl_SdkV2) {
}

func (to *DeleteAcl_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteAcl_SdkV2) {
}

func (m DeleteAcl_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteAcl_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAcl_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteAcl_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal": m.Principal,
			"scope":     m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAcl_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"scope":     types.StringType,
		},
	}
}

type DeleteCredentialsRequest_SdkV2 struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-"`
}

func (to *DeleteCredentialsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCredentialsRequest_SdkV2) {
}

func (to *DeleteCredentialsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCredentialsRequest_SdkV2) {
}

func (m DeleteCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": m.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.Int64Type,
		},
	}
}

type DeleteCredentialsResponse_SdkV2 struct {
}

func (to *DeleteCredentialsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCredentialsResponse_SdkV2) {
}

func (to *DeleteCredentialsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCredentialsResponse_SdkV2) {
}

func (m DeleteCredentialsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRepoRequest_SdkV2 struct {
	// The ID for the corresponding repo to delete.
	RepoId types.Int64 `tfsdk:"-"`
}

func (to *DeleteRepoRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRepoRequest_SdkV2) {
}

func (to *DeleteRepoRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRepoRequest_SdkV2) {
}

func (m DeleteRepoRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteRepoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRepoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRepoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": m.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRepoRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.Int64Type,
		},
	}
}

type DeleteRepoResponse_SdkV2 struct {
}

func (to *DeleteRepoResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRepoResponse_SdkV2) {
}

func (to *DeleteRepoResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRepoResponse_SdkV2) {
}

func (m DeleteRepoResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRepoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRepoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRepoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRepoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRepoResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteResponse_SdkV2 struct {
}

func (to *DeleteResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteResponse_SdkV2) {
}

func (to *DeleteResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteResponse_SdkV2) {
}

func (m DeleteResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteScope_SdkV2 struct {
	// Name of the scope to delete.
	Scope types.String `tfsdk:"scope"`
}

func (to *DeleteScope_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteScope_SdkV2) {
}

func (to *DeleteScope_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteScope_SdkV2) {
}

func (m DeleteScope_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteScope_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteScope_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteScope_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteScope_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope": types.StringType,
		},
	}
}

type DeleteSecret_SdkV2 struct {
	// Name of the secret to delete.
	Key types.String `tfsdk:"key"`
	// The name of the scope that contains the secret to delete.
	Scope types.String `tfsdk:"scope"`
}

func (to *DeleteSecret_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSecret_SdkV2) {
}

func (to *DeleteSecret_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteSecret_SdkV2) {
}

func (m DeleteSecret_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteSecret_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSecret_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteSecret_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"scope": m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSecret_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"scope": types.StringType,
		},
	}
}

type DeleteSecretResponse_SdkV2 struct {
}

func (to *DeleteSecretResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSecretResponse_SdkV2) {
}

func (to *DeleteSecretResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteSecretResponse_SdkV2) {
}

func (m DeleteSecretResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSecretResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSecretResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSecretResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteSecretResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSecretResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ExportRequest_SdkV2 struct {
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

func (to *ExportRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExportRequest_SdkV2) {
}

func (to *ExportRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExportRequest_SdkV2) {
}

func (m ExportRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()
	attrs["format"] = attrs["format"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExportRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ExportRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"format": m.Format,
			"path":   m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExportRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"format": types.StringType,
			"path":   types.StringType,
		},
	}
}

// The request field `direct_download` determines whether a JSON response or
// binary contents are returned by this endpoint.
type ExportResponse_SdkV2 struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	Content types.String `tfsdk:"content"`
	// The file type of the exported file.
	FileType types.String `tfsdk:"file_type"`
}

func (to *ExportResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExportResponse_SdkV2) {
}

func (to *ExportResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExportResponse_SdkV2) {
}

func (m ExportResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExportResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ExportResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content":   m.Content,
			"file_type": m.FileType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExportResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content":   types.StringType,
			"file_type": types.StringType,
		},
	}
}

type GetAclRequest_SdkV2 struct {
	// The principal to fetch ACL information for.
	Principal types.String `tfsdk:"-"`
	// The name of the scope to fetch ACL information from.
	Scope types.String `tfsdk:"-"`
}

func (to *GetAclRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAclRequest_SdkV2) {
}

func (to *GetAclRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAclRequest_SdkV2) {
}

func (m GetAclRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetAclRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAclRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAclRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal": m.Principal,
			"scope":     m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAclRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal": types.StringType,
			"scope":     types.StringType,
		},
	}
}

type GetCredentialsRequest_SdkV2 struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-"`
}

func (to *GetCredentialsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCredentialsRequest_SdkV2) {
}

func (to *GetCredentialsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCredentialsRequest_SdkV2) {
}

func (m GetCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": m.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential_id": types.Int64Type,
		},
	}
}

type GetCredentialsResponse_SdkV2 struct {
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

func (to *GetCredentialsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCredentialsResponse_SdkV2) {
}

func (to *GetCredentialsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCredentialsResponse_SdkV2) {
}

func (m GetCredentialsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m GetCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

type GetRepoPermissionLevelsRequest_SdkV2 struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

func (to *GetRepoPermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoPermissionLevelsRequest_SdkV2) {
}

func (to *GetRepoPermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRepoPermissionLevelsRequest_SdkV2) {
}

func (m GetRepoPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRepoPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRepoPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": m.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRepoPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.StringType,
		},
	}
}

type GetRepoPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetRepoPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetRepoPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRepoPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetRepoPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRepoPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(RepoPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRepoPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRepoPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: RepoPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetRepoPermissionLevelsResponse_SdkV2 as
// a slice of RepoPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetRepoPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]RepoPermissionsDescription_SdkV2, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []RepoPermissionsDescription_SdkV2
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetRepoPermissionLevelsResponse_SdkV2.
func (m *GetRepoPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []RepoPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetRepoPermissionsRequest_SdkV2 struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

func (to *GetRepoPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoPermissionsRequest_SdkV2) {
}

func (to *GetRepoPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRepoPermissionsRequest_SdkV2) {
}

func (m GetRepoPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRepoPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRepoPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": m.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRepoPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.StringType,
		},
	}
}

type GetRepoRequest_SdkV2 struct {
	// ID of the Git folder (repo) object in the workspace.
	RepoId types.Int64 `tfsdk:"-"`
}

func (to *GetRepoRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoRequest_SdkV2) {
}

func (to *GetRepoRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRepoRequest_SdkV2) {
}

func (m GetRepoRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRepoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRepoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": m.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRepoRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"repo_id": types.Int64Type,
		},
	}
}

type GetRepoResponse_SdkV2 struct {
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
	SparseCheckout types.List `tfsdk:"sparse_checkout"`
	// URL of the linked Git repository.
	Url types.String `tfsdk:"url"`
}

func (to *GetRepoResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoResponse_SdkV2) {
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

func (to *GetRepoResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRepoResponse_SdkV2) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				toSparseCheckout.SyncFieldsDuringRead(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (m GetRepoResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetOptional()
	attrs["head_commit_id"] = attrs["head_commit_id"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()
	attrs["provider"] = attrs["provider"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m GetRepoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRepoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m GetRepoResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":         types.StringType,
			"head_commit_id": types.StringType,
			"id":             types.Int64Type,
			"path":           types.StringType,
			"provider":       types.StringType,
			"sparse_checkout": basetypes.ListType{
				ElemType: SparseCheckout_SdkV2{}.Type(ctx),
			},
			"url": types.StringType,
		},
	}
}

// GetSparseCheckout returns the value of the SparseCheckout field in GetRepoResponse_SdkV2 as
// a SparseCheckout_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetRepoResponse_SdkV2) GetSparseCheckout(ctx context.Context) (SparseCheckout_SdkV2, bool) {
	var e SparseCheckout_SdkV2
	if m.SparseCheckout.IsNull() || m.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout_SdkV2
	d := m.SparseCheckout.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparseCheckout sets the value of the SparseCheckout field in GetRepoResponse_SdkV2.
func (m *GetRepoResponse_SdkV2) SetSparseCheckout(ctx context.Context, v SparseCheckout_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sparse_checkout"]
	m.SparseCheckout = types.ListValueMust(t, vs)
}

type GetSecretRequest_SdkV2 struct {
	// Name of the secret to fetch value information.
	Key types.String `tfsdk:"-"`
	// The name of the scope that contains the secret.
	Scope types.String `tfsdk:"-"`
}

func (to *GetSecretRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSecretRequest_SdkV2) {
}

func (to *GetSecretRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSecretRequest_SdkV2) {
}

func (m GetSecretRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetSecretRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSecretRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSecretRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"scope": m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSecretRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"scope": types.StringType,
		},
	}
}

type GetSecretResponse_SdkV2 struct {
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key"`
	// The value of the secret in its byte representation.
	Value types.String `tfsdk:"value"`
}

func (to *GetSecretResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSecretResponse_SdkV2) {
}

func (to *GetSecretResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSecretResponse_SdkV2) {
}

func (m GetSecretResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetSecretResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSecretResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSecretResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSecretResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type GetStatusRequest_SdkV2 struct {
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"-"`
}

func (to *GetStatusRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatusRequest_SdkV2) {
}

func (to *GetStatusRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetStatusRequest_SdkV2) {
}

func (m GetStatusRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetStatusRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatusRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetStatusRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStatusRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type GetWorkspaceObjectPermissionLevelsRequest_SdkV2 struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-"`
}

func (to *GetWorkspaceObjectPermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceObjectPermissionLevelsRequest_SdkV2) {
}

func (to *GetWorkspaceObjectPermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceObjectPermissionLevelsRequest_SdkV2) {
}

func (m GetWorkspaceObjectPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWorkspaceObjectPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceObjectPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_object_id":   m.WorkspaceObjectId,
			"workspace_object_type": m.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceObjectPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_object_id":   types.StringType,
			"workspace_object_type": types.StringType,
		},
	}
}

type GetWorkspaceObjectPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetWorkspaceObjectPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceObjectPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetWorkspaceObjectPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceObjectPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetWorkspaceObjectPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWorkspaceObjectPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(WorkspaceObjectPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceObjectPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceObjectPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: WorkspaceObjectPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetWorkspaceObjectPermissionLevelsResponse_SdkV2 as
// a slice of WorkspaceObjectPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetWorkspaceObjectPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]WorkspaceObjectPermissionsDescription_SdkV2, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectPermissionsDescription_SdkV2
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetWorkspaceObjectPermissionLevelsResponse_SdkV2.
func (m *GetWorkspaceObjectPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []WorkspaceObjectPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetWorkspaceObjectPermissionsRequest_SdkV2 struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-"`
}

func (to *GetWorkspaceObjectPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceObjectPermissionsRequest_SdkV2) {
}

func (to *GetWorkspaceObjectPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceObjectPermissionsRequest_SdkV2) {
}

func (m GetWorkspaceObjectPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWorkspaceObjectPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceObjectPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_object_id":   m.WorkspaceObjectId,
			"workspace_object_type": m.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceObjectPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_object_id":   types.StringType,
			"workspace_object_type": types.StringType,
		},
	}
}

type Import_SdkV2 struct {
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

func (to *Import_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Import_SdkV2) {
}

func (to *Import_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Import_SdkV2) {
}

func (m Import_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Import_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Import_SdkV2
// only implements ToObjectValue() and Type().
func (m Import_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Import_SdkV2) Type(ctx context.Context) attr.Type {
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

type ImportResponse_SdkV2 struct {
}

func (to *ImportResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ImportResponse_SdkV2) {
}

func (to *ImportResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ImportResponse_SdkV2) {
}

func (m ImportResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ImportResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ImportResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ImportResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ImportResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ImportResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListAclsRequest_SdkV2 struct {
	// The name of the scope to fetch ACL information from.
	Scope types.String `tfsdk:"-"`
}

func (to *ListAclsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAclsRequest_SdkV2) {
}

func (to *ListAclsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAclsRequest_SdkV2) {
}

func (m ListAclsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAclsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAclsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListAclsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAclsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope": types.StringType,
		},
	}
}

type ListAclsResponse_SdkV2 struct {
	// The associated ACLs rule applied to principals in the given scope.
	Items types.List `tfsdk:"items"`
}

func (to *ListAclsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAclsResponse_SdkV2) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (to *ListAclsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAclsResponse_SdkV2) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (m ListAclsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAclsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(AclItem_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAclsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListAclsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items": m.Items,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAclsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: AclItem_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetItems returns the value of the Items field in ListAclsResponse_SdkV2 as
// a slice of AclItem_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListAclsResponse_SdkV2) GetItems(ctx context.Context) ([]AclItem_SdkV2, bool) {
	if m.Items.IsNull() || m.Items.IsUnknown() {
		return nil, false
	}
	var v []AclItem_SdkV2
	d := m.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListAclsResponse_SdkV2.
func (m *ListAclsResponse_SdkV2) SetItems(ctx context.Context, v []AclItem_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Items = types.ListValueMust(t, vs)
}

type ListCredentialsRequest_SdkV2 struct {
}

func (to *ListCredentialsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCredentialsRequest_SdkV2) {
}

func (to *ListCredentialsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCredentialsRequest_SdkV2) {
}

func (m ListCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListCredentialsResponse_SdkV2 struct {
	// List of credentials.
	Credentials types.List `tfsdk:"credentials"`
}

func (to *ListCredentialsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCredentialsResponse_SdkV2) {
	if !from.Credentials.IsNull() && !from.Credentials.IsUnknown() && to.Credentials.IsNull() && len(from.Credentials.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Credentials, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Credentials = from.Credentials
	}
}

func (to *ListCredentialsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCredentialsResponse_SdkV2) {
	if !from.Credentials.IsNull() && !from.Credentials.IsUnknown() && to.Credentials.IsNull() && len(from.Credentials.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Credentials, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Credentials = from.Credentials
	}
}

func (m ListCredentialsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credentials": reflect.TypeOf(CredentialInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials": m.Credentials,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credentials": basetypes.ListType{
				ElemType: CredentialInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetCredentials returns the value of the Credentials field in ListCredentialsResponse_SdkV2 as
// a slice of CredentialInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListCredentialsResponse_SdkV2) GetCredentials(ctx context.Context) ([]CredentialInfo_SdkV2, bool) {
	if m.Credentials.IsNull() || m.Credentials.IsUnknown() {
		return nil, false
	}
	var v []CredentialInfo_SdkV2
	d := m.Credentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCredentials sets the value of the Credentials field in ListCredentialsResponse_SdkV2.
func (m *ListCredentialsResponse_SdkV2) SetCredentials(ctx context.Context, v []CredentialInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["credentials"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Credentials = types.ListValueMust(t, vs)
}

type ListReposRequest_SdkV2 struct {
	// Token used to get the next page of results. If not specified, returns the
	// first page of results as well as a next page token if there are more
	// results.
	NextPageToken types.String `tfsdk:"-"`
	// Filters repos that have paths starting with the given path prefix. If not
	// provided or when provided an effectively empty prefix (`/` or
	// `/Workspace`) Git folders (repos) from `/Workspace/Repos` will be served.
	PathPrefix types.String `tfsdk:"-"`
}

func (to *ListReposRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListReposRequest_SdkV2) {
}

func (to *ListReposRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListReposRequest_SdkV2) {
}

func (m ListReposRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListReposRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListReposRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListReposRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"path_prefix":     m.PathPrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListReposRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"path_prefix":     types.StringType,
		},
	}
}

type ListReposResponse_SdkV2 struct {
	// Token that can be specified as a query parameter to the `GET /repos`
	// endpoint to retrieve the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of Git folders (repos).
	Repos types.List `tfsdk:"repos"`
}

func (to *ListReposResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListReposResponse_SdkV2) {
	if !from.Repos.IsNull() && !from.Repos.IsUnknown() && to.Repos.IsNull() && len(from.Repos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Repos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Repos = from.Repos
	}
}

func (to *ListReposResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListReposResponse_SdkV2) {
	if !from.Repos.IsNull() && !from.Repos.IsUnknown() && to.Repos.IsNull() && len(from.Repos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Repos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Repos = from.Repos
	}
}

func (m ListReposResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListReposResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"repos": reflect.TypeOf(RepoInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListReposResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListReposResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"repos":           m.Repos,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListReposResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"repos": basetypes.ListType{
				ElemType: RepoInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRepos returns the value of the Repos field in ListReposResponse_SdkV2 as
// a slice of RepoInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListReposResponse_SdkV2) GetRepos(ctx context.Context) ([]RepoInfo_SdkV2, bool) {
	if m.Repos.IsNull() || m.Repos.IsUnknown() {
		return nil, false
	}
	var v []RepoInfo_SdkV2
	d := m.Repos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRepos sets the value of the Repos field in ListReposResponse_SdkV2.
func (m *ListReposResponse_SdkV2) SetRepos(ctx context.Context, v []RepoInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["repos"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Repos = types.ListValueMust(t, vs)
}

type ListResponse_SdkV2 struct {
	// List of objects.
	Objects types.List `tfsdk:"objects"`
}

func (to *ListResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListResponse_SdkV2) {
	if !from.Objects.IsNull() && !from.Objects.IsUnknown() && to.Objects.IsNull() && len(from.Objects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Objects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Objects = from.Objects
	}
}

func (to *ListResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListResponse_SdkV2) {
	if !from.Objects.IsNull() && !from.Objects.IsUnknown() && to.Objects.IsNull() && len(from.Objects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Objects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Objects = from.Objects
	}
}

func (m ListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objects": reflect.TypeOf(ObjectInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"objects": m.Objects,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"objects": basetypes.ListType{
				ElemType: ObjectInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetObjects returns the value of the Objects field in ListResponse_SdkV2 as
// a slice of ObjectInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListResponse_SdkV2) GetObjects(ctx context.Context) ([]ObjectInfo_SdkV2, bool) {
	if m.Objects.IsNull() || m.Objects.IsUnknown() {
		return nil, false
	}
	var v []ObjectInfo_SdkV2
	d := m.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in ListResponse_SdkV2.
func (m *ListResponse_SdkV2) SetObjects(ctx context.Context, v []ObjectInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Objects = types.ListValueMust(t, vs)
}

type ListScopesRequest_SdkV2 struct {
}

func (to *ListScopesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListScopesRequest_SdkV2) {
}

func (to *ListScopesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListScopesRequest_SdkV2) {
}

func (m ListScopesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListScopesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListScopesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListScopesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListScopesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListScopesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListScopesResponse_SdkV2 struct {
	// The available secret scopes.
	Scopes types.List `tfsdk:"scopes"`
}

func (to *ListScopesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListScopesResponse_SdkV2) {
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
}

func (to *ListScopesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListScopesResponse_SdkV2) {
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
}

func (m ListScopesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListScopesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(SecretScope_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListScopesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListScopesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scopes": m.Scopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListScopesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scopes": basetypes.ListType{
				ElemType: SecretScope_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetScopes returns the value of the Scopes field in ListScopesResponse_SdkV2 as
// a slice of SecretScope_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListScopesResponse_SdkV2) GetScopes(ctx context.Context) ([]SecretScope_SdkV2, bool) {
	if m.Scopes.IsNull() || m.Scopes.IsUnknown() {
		return nil, false
	}
	var v []SecretScope_SdkV2
	d := m.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in ListScopesResponse_SdkV2.
func (m *ListScopesResponse_SdkV2) SetScopes(ctx context.Context, v []SecretScope_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

type ListSecretsRequest_SdkV2 struct {
	// The name of the scope to list secrets within.
	Scope types.String `tfsdk:"-"`
}

func (to *ListSecretsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSecretsRequest_SdkV2) {
}

func (to *ListSecretsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSecretsRequest_SdkV2) {
}

func (m ListSecretsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSecretsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSecretsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSecretsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSecretsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"scope": types.StringType,
		},
	}
}

type ListSecretsResponse_SdkV2 struct {
	// Metadata information of all secrets contained within the given scope.
	Secrets types.List `tfsdk:"secrets"`
}

func (to *ListSecretsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSecretsResponse_SdkV2) {
	if !from.Secrets.IsNull() && !from.Secrets.IsUnknown() && to.Secrets.IsNull() && len(from.Secrets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Secrets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Secrets = from.Secrets
	}
}

func (to *ListSecretsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSecretsResponse_SdkV2) {
	if !from.Secrets.IsNull() && !from.Secrets.IsUnknown() && to.Secrets.IsNull() && len(from.Secrets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Secrets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Secrets = from.Secrets
	}
}

func (m ListSecretsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSecretsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"secrets": reflect.TypeOf(SecretMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSecretsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSecretsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"secrets": m.Secrets,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSecretsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"secrets": basetypes.ListType{
				ElemType: SecretMetadata_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSecrets returns the value of the Secrets field in ListSecretsResponse_SdkV2 as
// a slice of SecretMetadata_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSecretsResponse_SdkV2) GetSecrets(ctx context.Context) ([]SecretMetadata_SdkV2, bool) {
	if m.Secrets.IsNull() || m.Secrets.IsUnknown() {
		return nil, false
	}
	var v []SecretMetadata_SdkV2
	d := m.Secrets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecrets sets the value of the Secrets field in ListSecretsResponse_SdkV2.
func (m *ListSecretsResponse_SdkV2) SetSecrets(ctx context.Context, v []SecretMetadata_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["secrets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Secrets = types.ListValueMust(t, vs)
}

type ListWorkspaceRequest_SdkV2 struct {
	// UTC timestamp in milliseconds
	NotebooksModifiedAfter types.Int64 `tfsdk:"-"`
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"-"`
}

func (to *ListWorkspaceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWorkspaceRequest_SdkV2) {
}

func (to *ListWorkspaceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWorkspaceRequest_SdkV2) {
}

func (m ListWorkspaceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notebooks_modified_after": m.NotebooksModifiedAfter,
			"path":                     m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"notebooks_modified_after": types.Int64Type,
			"path":                     types.StringType,
		},
	}
}

type Mkdirs_SdkV2 struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path types.String `tfsdk:"path"`
}

func (to *Mkdirs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Mkdirs_SdkV2) {
}

func (to *Mkdirs_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Mkdirs_SdkV2) {
}

func (m Mkdirs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Mkdirs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Mkdirs_SdkV2
// only implements ToObjectValue() and Type().
func (m Mkdirs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Mkdirs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type MkdirsResponse_SdkV2 struct {
}

func (to *MkdirsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MkdirsResponse_SdkV2) {
}

func (to *MkdirsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from MkdirsResponse_SdkV2) {
}

func (m MkdirsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MkdirsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MkdirsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MkdirsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m MkdirsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m MkdirsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// The information of the object in workspace. It will be returned by “list“
// and “get-status“.
type ObjectInfo_SdkV2 struct {
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

func (to *ObjectInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ObjectInfo_SdkV2) {
}

func (to *ObjectInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ObjectInfo_SdkV2) {
}

func (m ObjectInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ObjectInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ObjectInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m ObjectInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ObjectInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

type PutAcl_SdkV2 struct {
	// The permission level applied to the principal.
	Permission types.String `tfsdk:"permission"`
	// The principal in which the permission is applied.
	Principal types.String `tfsdk:"principal"`
	// The name of the scope to apply permissions to.
	Scope types.String `tfsdk:"scope"`
}

func (to *PutAcl_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutAcl_SdkV2) {
}

func (to *PutAcl_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PutAcl_SdkV2) {
}

func (m PutAcl_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PutAcl_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAcl_SdkV2
// only implements ToObjectValue() and Type().
func (m PutAcl_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": m.Permission,
			"principal":  m.Principal,
			"scope":      m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PutAcl_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
			"principal":  types.StringType,
			"scope":      types.StringType,
		},
	}
}

type PutSecret_SdkV2 struct {
	// If specified, value will be stored as bytes.
	BytesValue types.String `tfsdk:"bytes_value"`
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key"`
	// The name of the scope to which the secret will be associated with.
	Scope types.String `tfsdk:"scope"`
	// If specified, note that the value will be stored in UTF-8 (MB4) form.
	StringValue types.String `tfsdk:"string_value"`
}

func (to *PutSecret_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutSecret_SdkV2) {
}

func (to *PutSecret_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PutSecret_SdkV2) {
}

func (m PutSecret_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PutSecret_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutSecret_SdkV2
// only implements ToObjectValue() and Type().
func (m PutSecret_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m PutSecret_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bytes_value":  types.StringType,
			"key":          types.StringType,
			"scope":        types.StringType,
			"string_value": types.StringType,
		},
	}
}

type RepoAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *RepoAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoAccessControlRequest_SdkV2) {
}

func (to *RepoAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepoAccessControlRequest_SdkV2) {
}

func (m RepoAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RepoAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m RepoAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type RepoAccessControlResponse_SdkV2 struct {
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

func (to *RepoAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *RepoAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepoAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m RepoAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(RepoPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RepoAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m RepoAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: RepoPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in RepoAccessControlResponse_SdkV2 as
// a slice of RepoPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RepoAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]RepoPermission_SdkV2, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []RepoPermission_SdkV2
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in RepoAccessControlResponse_SdkV2.
func (m *RepoAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []RepoPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

// Git folder (repo) information.
type RepoInfo_SdkV2 struct {
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
	SparseCheckout types.List `tfsdk:"sparse_checkout"`
	// URL of the remote git repository.
	Url types.String `tfsdk:"url"`
}

func (to *RepoInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoInfo_SdkV2) {
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

func (to *RepoInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepoInfo_SdkV2) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				toSparseCheckout.SyncFieldsDuringRead(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (m RepoInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetOptional()
	attrs["head_commit_id"] = attrs["head_commit_id"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()
	attrs["provider"] = attrs["provider"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m RepoInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m RepoInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m RepoInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":         types.StringType,
			"head_commit_id": types.StringType,
			"id":             types.Int64Type,
			"path":           types.StringType,
			"provider":       types.StringType,
			"sparse_checkout": basetypes.ListType{
				ElemType: SparseCheckout_SdkV2{}.Type(ctx),
			},
			"url": types.StringType,
		},
	}
}

// GetSparseCheckout returns the value of the SparseCheckout field in RepoInfo_SdkV2 as
// a SparseCheckout_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RepoInfo_SdkV2) GetSparseCheckout(ctx context.Context) (SparseCheckout_SdkV2, bool) {
	var e SparseCheckout_SdkV2
	if m.SparseCheckout.IsNull() || m.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout_SdkV2
	d := m.SparseCheckout.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparseCheckout sets the value of the SparseCheckout field in RepoInfo_SdkV2.
func (m *RepoInfo_SdkV2) SetSparseCheckout(ctx context.Context, v SparseCheckout_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sparse_checkout"]
	m.SparseCheckout = types.ListValueMust(t, vs)
}

type RepoPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *RepoPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *RepoPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepoPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m RepoPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermission_SdkV2
// only implements ToObjectValue() and Type().
func (m RepoPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in RepoPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RepoPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in RepoPermission_SdkV2.
func (m *RepoPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type RepoPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *RepoPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *RepoPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepoPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m RepoPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RepoAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (m RepoPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: RepoAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in RepoPermissions_SdkV2 as
// a slice of RepoAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RepoPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]RepoAccessControlResponse_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RepoAccessControlResponse_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RepoPermissions_SdkV2.
func (m *RepoPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []RepoAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type RepoPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *RepoPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoPermissionsDescription_SdkV2) {
}

func (to *RepoPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepoPermissionsDescription_SdkV2) {
}

func (m RepoPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (m RepoPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type RepoPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

func (to *RepoPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *RepoPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepoPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m RepoPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RepoPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RepoAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RepoPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"repo_id":             m.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RepoPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: RepoAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"repo_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in RepoPermissionsRequest_SdkV2 as
// a slice of RepoAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RepoPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]RepoAccessControlRequest_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RepoAccessControlRequest_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RepoPermissionsRequest_SdkV2.
func (m *RepoPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []RepoAccessControlRequest_SdkV2) {
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
type SecretMetadata_SdkV2 struct {
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key"`
	// The last updated timestamp (in milliseconds) for the secret.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
}

func (to *SecretMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SecretMetadata_SdkV2) {
}

func (to *SecretMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SecretMetadata_SdkV2) {
}

func (m SecretMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SecretMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m SecretMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":                    m.Key,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SecretMetadata_SdkV2) Type(ctx context.Context) attr.Type {
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
type SecretScope_SdkV2 struct {
	// The type of secret scope backend.
	BackendType types.String `tfsdk:"backend_type"`
	// The metadata for the secret scope if the type is ``AZURE_KEYVAULT``
	KeyvaultMetadata types.List `tfsdk:"keyvault_metadata"`
	// A unique name to identify the secret scope.
	Name types.String `tfsdk:"name"`
}

func (to *SecretScope_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SecretScope_SdkV2) {
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

func (to *SecretScope_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SecretScope_SdkV2) {
	if !from.KeyvaultMetadata.IsNull() && !from.KeyvaultMetadata.IsUnknown() {
		if toKeyvaultMetadata, ok := to.GetKeyvaultMetadata(ctx); ok {
			if fromKeyvaultMetadata, ok := from.GetKeyvaultMetadata(ctx); ok {
				toKeyvaultMetadata.SyncFieldsDuringRead(ctx, fromKeyvaultMetadata)
				to.SetKeyvaultMetadata(ctx, toKeyvaultMetadata)
			}
		}
	}
}

func (m SecretScope_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["backend_type"] = attrs["backend_type"].SetOptional()
	attrs["keyvault_metadata"] = attrs["keyvault_metadata"].SetOptional()
	attrs["keyvault_metadata"] = attrs["keyvault_metadata"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m SecretScope_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"keyvault_metadata": reflect.TypeOf(AzureKeyVaultSecretScopeMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretScope_SdkV2
// only implements ToObjectValue() and Type().
func (m SecretScope_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"backend_type":      m.BackendType,
			"keyvault_metadata": m.KeyvaultMetadata,
			"name":              m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SecretScope_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"backend_type": types.StringType,
			"keyvault_metadata": basetypes.ListType{
				ElemType: AzureKeyVaultSecretScopeMetadata_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetKeyvaultMetadata returns the value of the KeyvaultMetadata field in SecretScope_SdkV2 as
// a AzureKeyVaultSecretScopeMetadata_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SecretScope_SdkV2) GetKeyvaultMetadata(ctx context.Context) (AzureKeyVaultSecretScopeMetadata_SdkV2, bool) {
	var e AzureKeyVaultSecretScopeMetadata_SdkV2
	if m.KeyvaultMetadata.IsNull() || m.KeyvaultMetadata.IsUnknown() {
		return e, false
	}
	var v []AzureKeyVaultSecretScopeMetadata_SdkV2
	d := m.KeyvaultMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetKeyvaultMetadata sets the value of the KeyvaultMetadata field in SecretScope_SdkV2.
func (m *SecretScope_SdkV2) SetKeyvaultMetadata(ctx context.Context, v AzureKeyVaultSecretScopeMetadata_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["keyvault_metadata"]
	m.KeyvaultMetadata = types.ListValueMust(t, vs)
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckout_SdkV2 struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns types.List `tfsdk:"patterns"`
}

func (to *SparseCheckout_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SparseCheckout_SdkV2) {
	if !from.Patterns.IsNull() && !from.Patterns.IsUnknown() && to.Patterns.IsNull() && len(from.Patterns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Patterns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Patterns = from.Patterns
	}
}

func (to *SparseCheckout_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SparseCheckout_SdkV2) {
	if !from.Patterns.IsNull() && !from.Patterns.IsUnknown() && to.Patterns.IsNull() && len(from.Patterns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Patterns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Patterns = from.Patterns
	}
}

func (m SparseCheckout_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SparseCheckout_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"patterns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparseCheckout_SdkV2
// only implements ToObjectValue() and Type().
func (m SparseCheckout_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"patterns": m.Patterns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SparseCheckout_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"patterns": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPatterns returns the value of the Patterns field in SparseCheckout_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SparseCheckout_SdkV2) GetPatterns(ctx context.Context) ([]types.String, bool) {
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

// SetPatterns sets the value of the Patterns field in SparseCheckout_SdkV2.
func (m *SparseCheckout_SdkV2) SetPatterns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["patterns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Patterns = types.ListValueMust(t, vs)
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckoutUpdate_SdkV2 struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns types.List `tfsdk:"patterns"`
}

func (to *SparseCheckoutUpdate_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SparseCheckoutUpdate_SdkV2) {
	if !from.Patterns.IsNull() && !from.Patterns.IsUnknown() && to.Patterns.IsNull() && len(from.Patterns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Patterns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Patterns = from.Patterns
	}
}

func (to *SparseCheckoutUpdate_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SparseCheckoutUpdate_SdkV2) {
	if !from.Patterns.IsNull() && !from.Patterns.IsUnknown() && to.Patterns.IsNull() && len(from.Patterns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Patterns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Patterns = from.Patterns
	}
}

func (m SparseCheckoutUpdate_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SparseCheckoutUpdate_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"patterns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparseCheckoutUpdate_SdkV2
// only implements ToObjectValue() and Type().
func (m SparseCheckoutUpdate_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"patterns": m.Patterns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SparseCheckoutUpdate_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"patterns": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPatterns returns the value of the Patterns field in SparseCheckoutUpdate_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SparseCheckoutUpdate_SdkV2) GetPatterns(ctx context.Context) ([]types.String, bool) {
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

// SetPatterns sets the value of the Patterns field in SparseCheckoutUpdate_SdkV2.
func (m *SparseCheckoutUpdate_SdkV2) SetPatterns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["patterns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Patterns = types.ListValueMust(t, vs)
}

type UpdateCredentialsRequest_SdkV2 struct {
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

func (to *UpdateCredentialsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCredentialsRequest_SdkV2) {
}

func (to *UpdateCredentialsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateCredentialsRequest_SdkV2) {
}

func (m UpdateCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type UpdateCredentialsResponse_SdkV2 struct {
}

func (to *UpdateCredentialsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCredentialsResponse_SdkV2) {
}

func (to *UpdateCredentialsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateCredentialsResponse_SdkV2) {
}

func (m UpdateCredentialsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateRepoRequest_SdkV2 struct {
	// Branch that the local version of the repo is checked out to.
	Branch types.String `tfsdk:"branch"`
	// ID of the Git folder (repo) object in the workspace.
	RepoId types.Int64 `tfsdk:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	SparseCheckout types.List `tfsdk:"sparse_checkout"`
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	Tag types.String `tfsdk:"tag"`
}

func (to *UpdateRepoRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRepoRequest_SdkV2) {
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

func (to *UpdateRepoRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRepoRequest_SdkV2) {
	if !from.SparseCheckout.IsNull() && !from.SparseCheckout.IsUnknown() {
		if toSparseCheckout, ok := to.GetSparseCheckout(ctx); ok {
			if fromSparseCheckout, ok := from.GetSparseCheckout(ctx); ok {
				toSparseCheckout.SyncFieldsDuringRead(ctx, fromSparseCheckout)
				to.SetSparseCheckout(ctx, toSparseCheckout)
			}
		}
	}
}

func (m UpdateRepoRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].SetOptional()
	attrs["sparse_checkout"] = attrs["sparse_checkout"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateRepoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckoutUpdate_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRepoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateRepoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateRepoRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":  types.StringType,
			"repo_id": types.Int64Type,
			"sparse_checkout": basetypes.ListType{
				ElemType: SparseCheckoutUpdate_SdkV2{}.Type(ctx),
			},
			"tag": types.StringType,
		},
	}
}

// GetSparseCheckout returns the value of the SparseCheckout field in UpdateRepoRequest_SdkV2 as
// a SparseCheckoutUpdate_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRepoRequest_SdkV2) GetSparseCheckout(ctx context.Context) (SparseCheckoutUpdate_SdkV2, bool) {
	var e SparseCheckoutUpdate_SdkV2
	if m.SparseCheckout.IsNull() || m.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckoutUpdate_SdkV2
	d := m.SparseCheckout.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparseCheckout sets the value of the SparseCheckout field in UpdateRepoRequest_SdkV2.
func (m *UpdateRepoRequest_SdkV2) SetSparseCheckout(ctx context.Context, v SparseCheckoutUpdate_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sparse_checkout"]
	m.SparseCheckout = types.ListValueMust(t, vs)
}

type UpdateRepoResponse_SdkV2 struct {
}

func (to *UpdateRepoResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRepoResponse_SdkV2) {
}

func (to *UpdateRepoResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRepoResponse_SdkV2) {
}

func (m UpdateRepoResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRepoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRepoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRepoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateRepoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRepoResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type WorkspaceObjectAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *WorkspaceObjectAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectAccessControlRequest_SdkV2) {
}

func (to *WorkspaceObjectAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectAccessControlRequest_SdkV2) {
}

func (m WorkspaceObjectAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m WorkspaceObjectAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type WorkspaceObjectAccessControlResponse_SdkV2 struct {
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

func (to *WorkspaceObjectAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *WorkspaceObjectAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m WorkspaceObjectAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(WorkspaceObjectPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m WorkspaceObjectAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: WorkspaceObjectPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in WorkspaceObjectAccessControlResponse_SdkV2 as
// a slice of WorkspaceObjectPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceObjectAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]WorkspaceObjectPermission_SdkV2, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectPermission_SdkV2
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in WorkspaceObjectAccessControlResponse_SdkV2.
func (m *WorkspaceObjectAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []WorkspaceObjectPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type WorkspaceObjectPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *WorkspaceObjectPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *WorkspaceObjectPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m WorkspaceObjectPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermission_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceObjectPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in WorkspaceObjectPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceObjectPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in WorkspaceObjectPermission_SdkV2.
func (m *WorkspaceObjectPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type WorkspaceObjectPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *WorkspaceObjectPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *WorkspaceObjectPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m WorkspaceObjectPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WorkspaceObjectAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceObjectPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: WorkspaceObjectAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in WorkspaceObjectPermissions_SdkV2 as
// a slice of WorkspaceObjectAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceObjectPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]WorkspaceObjectAccessControlResponse_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectAccessControlResponse_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WorkspaceObjectPermissions_SdkV2.
func (m *WorkspaceObjectPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []WorkspaceObjectAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type WorkspaceObjectPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *WorkspaceObjectPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectPermissionsDescription_SdkV2) {
}

func (to *WorkspaceObjectPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectPermissionsDescription_SdkV2) {
}

func (m WorkspaceObjectPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceObjectPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type WorkspaceObjectPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-"`
}

func (to *WorkspaceObjectPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *WorkspaceObjectPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m WorkspaceObjectPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceObjectPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WorkspaceObjectAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceObjectPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list":   m.AccessControlList,
			"workspace_object_id":   m.WorkspaceObjectId,
			"workspace_object_type": m.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceObjectPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: WorkspaceObjectAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"workspace_object_id":   types.StringType,
			"workspace_object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in WorkspaceObjectPermissionsRequest_SdkV2 as
// a slice of WorkspaceObjectAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceObjectPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]WorkspaceObjectAccessControlRequest_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectAccessControlRequest_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WorkspaceObjectPermissionsRequest_SdkV2.
func (m *WorkspaceObjectPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []WorkspaceObjectAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}
