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

func (c AclItem_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AclItem_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AclItem_SdkV2
// only implements ToObjectValue() and Type().
func (o AclItem_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": o.Permission,
			"principal":  o.Principal,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AclItem_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c AzureKeyVaultSecretScopeMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AzureKeyVaultSecretScopeMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AzureKeyVaultSecretScopeMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (o AzureKeyVaultSecretScopeMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dns_name":    o.DnsName,
			"resource_id": o.ResourceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AzureKeyVaultSecretScopeMetadata_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c CreateCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"git_provider":            o.GitProvider,
			"git_username":            o.GitUsername,
			"is_default_for_provider": o.IsDefaultForProvider,
			"name":                    o.Name,
			"personal_access_token":   o.PersonalAccessToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c CreateCredentialsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":           o.CredentialId,
			"git_provider":            o.GitProvider,
			"git_username":            o.GitUsername,
			"is_default_for_provider": o.IsDefaultForProvider,
			"name":                    o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c CreateRepoRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateRepoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRepoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateRepoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateRepoRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateRepoRequest_SdkV2) GetSparseCheckout(ctx context.Context) (SparseCheckout_SdkV2, bool) {
	var e SparseCheckout_SdkV2
	if o.SparseCheckout.IsNull() || o.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout_SdkV2
	d := o.SparseCheckout.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparseCheckout sets the value of the SparseCheckout field in CreateRepoRequest_SdkV2.
func (o *CreateRepoRequest_SdkV2) SetSparseCheckout(ctx context.Context, v SparseCheckout_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sparse_checkout"]
	o.SparseCheckout = types.ListValueMust(t, vs)
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

func (c CreateRepoResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateRepoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRepoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateRepoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateRepoResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateRepoResponse_SdkV2) GetSparseCheckout(ctx context.Context) (SparseCheckout_SdkV2, bool) {
	var e SparseCheckout_SdkV2
	if o.SparseCheckout.IsNull() || o.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout_SdkV2
	d := o.SparseCheckout.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparseCheckout sets the value of the SparseCheckout field in CreateRepoResponse_SdkV2.
func (o *CreateRepoResponse_SdkV2) SetSparseCheckout(ctx context.Context, v SparseCheckout_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sparse_checkout"]
	o.SparseCheckout = types.ListValueMust(t, vs)
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

func (c CreateScope_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateScope_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"backend_azure_keyvault": reflect.TypeOf(AzureKeyVaultSecretScopeMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateScope_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateScope_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateScope_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateScope_SdkV2) GetBackendAzureKeyvault(ctx context.Context) (AzureKeyVaultSecretScopeMetadata_SdkV2, bool) {
	var e AzureKeyVaultSecretScopeMetadata_SdkV2
	if o.BackendAzureKeyvault.IsNull() || o.BackendAzureKeyvault.IsUnknown() {
		return e, false
	}
	var v []AzureKeyVaultSecretScopeMetadata_SdkV2
	d := o.BackendAzureKeyvault.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBackendAzureKeyvault sets the value of the BackendAzureKeyvault field in CreateScope_SdkV2.
func (o *CreateScope_SdkV2) SetBackendAzureKeyvault(ctx context.Context, v AzureKeyVaultSecretScopeMetadata_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["backend_azure_keyvault"]
	o.BackendAzureKeyvault = types.ListValueMust(t, vs)
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

func (c CredentialInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CredentialInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CredentialInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o CredentialInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":           o.CredentialId,
			"git_provider":            o.GitProvider,
			"git_username":            o.GitUsername,
			"is_default_for_provider": o.IsDefaultForProvider,
			"name":                    o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CredentialInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c Delete_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Delete_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Delete_SdkV2
// only implements ToObjectValue() and Type().
func (o Delete_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path":      o.Path,
			"recursive": o.Recursive,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Delete_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c DeleteAcl_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteAcl_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAcl_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAcl_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal": o.Principal,
			"scope":     o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAcl_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c DeleteCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c DeleteCredentialsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c DeleteRepoRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteRepoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRepoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRepoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": o.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRepoRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c DeleteRepoResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRepoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRepoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRepoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRepoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRepoResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c DeleteResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type DeleteScope_SdkV2 struct {
	// Name of the scope to delete.
	Scope types.String `tfsdk:"scope"`
}

func (to *DeleteScope_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteScope_SdkV2) {
}

func (to *DeleteScope_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteScope_SdkV2) {
}

func (c DeleteScope_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteScope_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteScope_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteScope_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteScope_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c DeleteSecret_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteSecret_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSecret_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteSecret_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"scope": o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSecret_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c DeleteSecretResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSecretResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSecretResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSecretResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteSecretResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSecretResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c ExportRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExportRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ExportRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"format": o.Format,
			"path":   o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExportRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c ExportResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExportResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ExportResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content":   o.Content,
			"file_type": o.FileType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExportResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetAclRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetAclRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAclRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAclRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"principal": o.Principal,
			"scope":     o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAclRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id": o.CredentialId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetCredentialsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":           o.CredentialId,
			"git_provider":            o.GitProvider,
			"git_username":            o.GitUsername,
			"is_default_for_provider": o.IsDefaultForProvider,
			"name":                    o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetRepoPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRepoPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRepoPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": o.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRepoPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetRepoPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRepoPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(RepoPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRepoPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRepoPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetRepoPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]RepoPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []RepoPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetRepoPermissionLevelsResponse_SdkV2.
func (o *GetRepoPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []RepoPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetRepoPermissionsRequest_SdkV2 struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-"`
}

func (to *GetRepoPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRepoPermissionsRequest_SdkV2) {
}

func (to *GetRepoPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRepoPermissionsRequest_SdkV2) {
}

func (c GetRepoPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRepoPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRepoPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": o.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRepoPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetRepoRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRepoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRepoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"repo_id": o.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRepoRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetRepoResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRepoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRepoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRepoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetRepoResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetRepoResponse_SdkV2) GetSparseCheckout(ctx context.Context) (SparseCheckout_SdkV2, bool) {
	var e SparseCheckout_SdkV2
	if o.SparseCheckout.IsNull() || o.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout_SdkV2
	d := o.SparseCheckout.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparseCheckout sets the value of the SparseCheckout field in GetRepoResponse_SdkV2.
func (o *GetRepoResponse_SdkV2) SetSparseCheckout(ctx context.Context, v SparseCheckout_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sparse_checkout"]
	o.SparseCheckout = types.ListValueMust(t, vs)
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

func (c GetSecretRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetSecretRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSecretRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetSecretRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"scope": o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSecretRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetSecretResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetSecretResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSecretResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetSecretResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSecretResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetStatusRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetStatusRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatusRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetStatusRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStatusRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetWorkspaceObjectPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetWorkspaceObjectPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceObjectPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_object_id":   o.WorkspaceObjectId,
			"workspace_object_type": o.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceObjectPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c GetWorkspaceObjectPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetWorkspaceObjectPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(WorkspaceObjectPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceObjectPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceObjectPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetWorkspaceObjectPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]WorkspaceObjectPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetWorkspaceObjectPermissionLevelsResponse_SdkV2.
func (o *GetWorkspaceObjectPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []WorkspaceObjectPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
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

func (c GetWorkspaceObjectPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetWorkspaceObjectPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceObjectPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetWorkspaceObjectPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_object_id":   o.WorkspaceObjectId,
			"workspace_object_type": o.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceObjectPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c Import_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Import_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Import_SdkV2
// only implements ToObjectValue() and Type().
func (o Import_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Import_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c ImportResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ImportResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ImportResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ImportResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ImportResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ImportResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c ListAclsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListAclsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAclsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAclsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAclsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c ListAclsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListAclsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(AclItem_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAclsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAclsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items": o.Items,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAclsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListAclsResponse_SdkV2) GetItems(ctx context.Context) ([]AclItem_SdkV2, bool) {
	if o.Items.IsNull() || o.Items.IsUnknown() {
		return nil, false
	}
	var v []AclItem_SdkV2
	d := o.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListAclsResponse_SdkV2.
func (o *ListAclsResponse_SdkV2) SetItems(ctx context.Context, v []AclItem_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Items = types.ListValueMust(t, vs)
}

type ListCredentialsRequest_SdkV2 struct {
}

func (to *ListCredentialsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCredentialsRequest_SdkV2) {
}

func (to *ListCredentialsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCredentialsRequest_SdkV2) {
}

func (c ListCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCredentialsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c ListCredentialsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"credentials": reflect.TypeOf(CredentialInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credentials": o.Credentials,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListCredentialsResponse_SdkV2) GetCredentials(ctx context.Context) ([]CredentialInfo_SdkV2, bool) {
	if o.Credentials.IsNull() || o.Credentials.IsUnknown() {
		return nil, false
	}
	var v []CredentialInfo_SdkV2
	d := o.Credentials.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCredentials sets the value of the Credentials field in ListCredentialsResponse_SdkV2.
func (o *ListCredentialsResponse_SdkV2) SetCredentials(ctx context.Context, v []CredentialInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["credentials"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Credentials = types.ListValueMust(t, vs)
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

func (c ListReposRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListReposRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListReposRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListReposRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"path_prefix":     o.PathPrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListReposRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c ListReposResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListReposResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"repos": reflect.TypeOf(RepoInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListReposResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListReposResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"repos":           o.Repos,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListReposResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListReposResponse_SdkV2) GetRepos(ctx context.Context) ([]RepoInfo_SdkV2, bool) {
	if o.Repos.IsNull() || o.Repos.IsUnknown() {
		return nil, false
	}
	var v []RepoInfo_SdkV2
	d := o.Repos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRepos sets the value of the Repos field in ListReposResponse_SdkV2.
func (o *ListReposResponse_SdkV2) SetRepos(ctx context.Context, v []RepoInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["repos"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Repos = types.ListValueMust(t, vs)
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

func (c ListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"objects": reflect.TypeOf(ObjectInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"objects": o.Objects,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListResponse_SdkV2) GetObjects(ctx context.Context) ([]ObjectInfo_SdkV2, bool) {
	if o.Objects.IsNull() || o.Objects.IsUnknown() {
		return nil, false
	}
	var v []ObjectInfo_SdkV2
	d := o.Objects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetObjects sets the value of the Objects field in ListResponse_SdkV2.
func (o *ListResponse_SdkV2) SetObjects(ctx context.Context, v []ObjectInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["objects"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Objects = types.ListValueMust(t, vs)
}

type ListScopesRequest_SdkV2 struct {
}

func (to *ListScopesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListScopesRequest_SdkV2) {
}

func (to *ListScopesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListScopesRequest_SdkV2) {
}

func (c ListScopesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListScopesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListScopesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListScopesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListScopesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListScopesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c ListScopesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListScopesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(SecretScope_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListScopesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListScopesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scopes": o.Scopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListScopesResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListScopesResponse_SdkV2) GetScopes(ctx context.Context) ([]SecretScope_SdkV2, bool) {
	if o.Scopes.IsNull() || o.Scopes.IsUnknown() {
		return nil, false
	}
	var v []SecretScope_SdkV2
	d := o.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in ListScopesResponse_SdkV2.
func (o *ListScopesResponse_SdkV2) SetScopes(ctx context.Context, v []SecretScope_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

type ListSecretsRequest_SdkV2 struct {
	// The name of the scope to list secrets within.
	Scope types.String `tfsdk:"-"`
}

func (to *ListSecretsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSecretsRequest_SdkV2) {
}

func (to *ListSecretsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSecretsRequest_SdkV2) {
}

func (c ListSecretsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListSecretsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSecretsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSecretsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"scope": o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSecretsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c ListSecretsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListSecretsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"secrets": reflect.TypeOf(SecretMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSecretsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListSecretsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"secrets": o.Secrets,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSecretsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListSecretsResponse_SdkV2) GetSecrets(ctx context.Context) ([]SecretMetadata_SdkV2, bool) {
	if o.Secrets.IsNull() || o.Secrets.IsUnknown() {
		return nil, false
	}
	var v []SecretMetadata_SdkV2
	d := o.Secrets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecrets sets the value of the Secrets field in ListSecretsResponse_SdkV2.
func (o *ListSecretsResponse_SdkV2) SetSecrets(ctx context.Context, v []SecretMetadata_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["secrets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Secrets = types.ListValueMust(t, vs)
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

func (c ListWorkspaceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notebooks_modified_after": o.NotebooksModifiedAfter,
			"path":                     o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c Mkdirs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Mkdirs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Mkdirs_SdkV2
// only implements ToObjectValue() and Type().
func (o Mkdirs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Mkdirs_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c MkdirsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MkdirsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MkdirsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MkdirsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o MkdirsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o MkdirsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c ObjectInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ObjectInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ObjectInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ObjectInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ObjectInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c PutAcl_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PutAcl_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAcl_SdkV2
// only implements ToObjectValue() and Type().
func (o PutAcl_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": o.Permission,
			"principal":  o.Principal,
			"scope":      o.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutAcl_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c PutSecret_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PutSecret_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutSecret_SdkV2
// only implements ToObjectValue() and Type().
func (o PutSecret_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PutSecret_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c RepoAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RepoAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RepoAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RepoAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c RepoAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RepoAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(RepoPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RepoAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RepoAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RepoAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]RepoPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []RepoPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in RepoAccessControlResponse_SdkV2.
func (o *RepoAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []RepoPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
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

func (c RepoInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RepoInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckout_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o RepoInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RepoInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RepoInfo_SdkV2) GetSparseCheckout(ctx context.Context) (SparseCheckout_SdkV2, bool) {
	var e SparseCheckout_SdkV2
	if o.SparseCheckout.IsNull() || o.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckout_SdkV2
	d := o.SparseCheckout.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparseCheckout sets the value of the SparseCheckout field in RepoInfo_SdkV2.
func (o *RepoInfo_SdkV2) SetSparseCheckout(ctx context.Context, v SparseCheckout_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sparse_checkout"]
	o.SparseCheckout = types.ListValueMust(t, vs)
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

func (c RepoPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RepoPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o RepoPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoPermission_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RepoPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in RepoPermission_SdkV2.
func (o *RepoPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
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

func (c RepoPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RepoPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RepoAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o RepoPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoPermissions_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RepoPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]RepoAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RepoAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RepoPermissions_SdkV2.
func (o *RepoPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []RepoAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type RepoPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *RepoPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RepoPermissionsDescription_SdkV2) {
}

func (to *RepoPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RepoPermissionsDescription_SdkV2) {
}

func (c RepoPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RepoPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o RepoPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c RepoPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RepoPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RepoAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RepoPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RepoPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"repo_id":             o.RepoId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RepoPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RepoPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]RepoAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RepoAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RepoPermissionsRequest_SdkV2.
func (o *RepoPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []RepoAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
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

func (c SecretMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SecretMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (o SecretMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":                    o.Key,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SecretMetadata_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c SecretScope_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SecretScope_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"keyvault_metadata": reflect.TypeOf(AzureKeyVaultSecretScopeMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretScope_SdkV2
// only implements ToObjectValue() and Type().
func (o SecretScope_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"backend_type":      o.BackendType,
			"keyvault_metadata": o.KeyvaultMetadata,
			"name":              o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SecretScope_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SecretScope_SdkV2) GetKeyvaultMetadata(ctx context.Context) (AzureKeyVaultSecretScopeMetadata_SdkV2, bool) {
	var e AzureKeyVaultSecretScopeMetadata_SdkV2
	if o.KeyvaultMetadata.IsNull() || o.KeyvaultMetadata.IsUnknown() {
		return e, false
	}
	var v []AzureKeyVaultSecretScopeMetadata_SdkV2
	d := o.KeyvaultMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetKeyvaultMetadata sets the value of the KeyvaultMetadata field in SecretScope_SdkV2.
func (o *SecretScope_SdkV2) SetKeyvaultMetadata(ctx context.Context, v AzureKeyVaultSecretScopeMetadata_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["keyvault_metadata"]
	o.KeyvaultMetadata = types.ListValueMust(t, vs)
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

func (c SparseCheckout_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SparseCheckout_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"patterns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparseCheckout_SdkV2
// only implements ToObjectValue() and Type().
func (o SparseCheckout_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"patterns": o.Patterns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparseCheckout_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SparseCheckout_SdkV2) GetPatterns(ctx context.Context) ([]types.String, bool) {
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

// SetPatterns sets the value of the Patterns field in SparseCheckout_SdkV2.
func (o *SparseCheckout_SdkV2) SetPatterns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["patterns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Patterns = types.ListValueMust(t, vs)
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

func (c SparseCheckoutUpdate_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SparseCheckoutUpdate_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"patterns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SparseCheckoutUpdate_SdkV2
// only implements ToObjectValue() and Type().
func (o SparseCheckoutUpdate_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"patterns": o.Patterns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SparseCheckoutUpdate_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SparseCheckoutUpdate_SdkV2) GetPatterns(ctx context.Context) ([]types.String, bool) {
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

// SetPatterns sets the value of the Patterns field in SparseCheckoutUpdate_SdkV2.
func (o *SparseCheckoutUpdate_SdkV2) SetPatterns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["patterns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Patterns = types.ListValueMust(t, vs)
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

func (c UpdateCredentialsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateCredentialsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCredentialsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCredentialsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential_id":           o.CredentialId,
			"git_provider":            o.GitProvider,
			"git_username":            o.GitUsername,
			"is_default_for_provider": o.IsDefaultForProvider,
			"name":                    o.Name,
			"personal_access_token":   o.PersonalAccessToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCredentialsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c UpdateCredentialsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCredentialsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCredentialsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCredentialsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCredentialsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCredentialsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c UpdateRepoRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateRepoRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sparse_checkout": reflect.TypeOf(SparseCheckoutUpdate_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRepoRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateRepoRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateRepoRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateRepoRequest_SdkV2) GetSparseCheckout(ctx context.Context) (SparseCheckoutUpdate_SdkV2, bool) {
	var e SparseCheckoutUpdate_SdkV2
	if o.SparseCheckout.IsNull() || o.SparseCheckout.IsUnknown() {
		return e, false
	}
	var v []SparseCheckoutUpdate_SdkV2
	d := o.SparseCheckout.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSparseCheckout sets the value of the SparseCheckout field in UpdateRepoRequest_SdkV2.
func (o *UpdateRepoRequest_SdkV2) SetSparseCheckout(ctx context.Context, v SparseCheckoutUpdate_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["sparse_checkout"]
	o.SparseCheckout = types.ListValueMust(t, vs)
}

type UpdateRepoResponse_SdkV2 struct {
}

func (to *UpdateRepoResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRepoResponse_SdkV2) {
}

func (to *UpdateRepoResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRepoResponse_SdkV2) {
}

func (c UpdateRepoResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRepoResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRepoResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRepoResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateRepoResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRepoResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c WorkspaceObjectAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspaceObjectAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o WorkspaceObjectAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c WorkspaceObjectAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspaceObjectAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(WorkspaceObjectPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o WorkspaceObjectAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *WorkspaceObjectAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]WorkspaceObjectPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in WorkspaceObjectAccessControlResponse_SdkV2.
func (o *WorkspaceObjectAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []WorkspaceObjectPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
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

func (c WorkspaceObjectPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspaceObjectPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceObjectPermission_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *WorkspaceObjectPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in WorkspaceObjectPermission_SdkV2.
func (o *WorkspaceObjectPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
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

func (c WorkspaceObjectPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspaceObjectPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WorkspaceObjectAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceObjectPermissions_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *WorkspaceObjectPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]WorkspaceObjectAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WorkspaceObjectPermissions_SdkV2.
func (o *WorkspaceObjectPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []WorkspaceObjectAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type WorkspaceObjectPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *WorkspaceObjectPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceObjectPermissionsDescription_SdkV2) {
}

func (to *WorkspaceObjectPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceObjectPermissionsDescription_SdkV2) {
}

func (c WorkspaceObjectPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspaceObjectPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceObjectPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
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

func (c WorkspaceObjectPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspaceObjectPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(WorkspaceObjectAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceObjectPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o WorkspaceObjectPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list":   o.AccessControlList,
			"workspace_object_id":   o.WorkspaceObjectId,
			"workspace_object_type": o.WorkspaceObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceObjectPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *WorkspaceObjectPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]WorkspaceObjectAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceObjectAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in WorkspaceObjectPermissionsRequest_SdkV2.
func (o *WorkspaceObjectPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []WorkspaceObjectAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}
