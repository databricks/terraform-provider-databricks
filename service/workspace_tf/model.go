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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AclItem struct {
	// The permission level applied to the principal.
	Permission AclPermission `tfsdk:"permission" tf:""`
	// The principal in which the permission is applied.
	Principal types.String `tfsdk:"principal" tf:""`
}

type AclPermission string

const AclPermissionManage AclPermission = `MANAGE`

const AclPermissionRead AclPermission = `READ`

const AclPermissionWrite AclPermission = `WRITE`

// String representation for [fmt.Print]
func (f *AclPermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AclPermission) Set(v string) error {
	switch v {
	case `MANAGE`, `READ`, `WRITE`:
		*f = AclPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANAGE", "READ", "WRITE"`, v)
	}
}

// Type always returns AclPermission to satisfy [pflag.Value] interface
func (f *AclPermission) Type() string {
	return "AclPermission"
}

type AzureKeyVaultSecretScopeMetadata struct {
	// The DNS of the KeyVault
	DnsName types.String `tfsdk:"dns_name" tf:""`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	ResourceId types.String `tfsdk:"resource_id" tf:""`
}

type CreateCredentials struct {
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
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
	// scoped access tokens. [Learn more]. The personal access token used to
	// authenticate to the corresponding Git
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken types.String `tfsdk:"personal_access_token" tf:"optional"`
}

type CreateCredentialsResponse struct {
	// ID of the credential object in the workspace.
	CredentialId types.Int64 `tfsdk:"credential_id" tf:"optional"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider types.String `tfsdk:"git_provider" tf:"optional"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername types.String `tfsdk:"git_username" tf:"optional"`
}

type CreateRepo struct {
	// Desired path for the repo in the workspace. Almost any path in the
	// workspace can be chosen. If repo is created in /Repos, path must be in
	// the format /Repos/{folder}/{repo-name}.
	Path types.String `tfsdk:"path" tf:"optional"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	Provider types.String `tfsdk:"provider" tf:""`
	// If specified, the repo will be created with sparse checkout enabled. You
	// cannot enable/disable sparse checkout after the repo is created.
	SparseCheckout *SparseCheckout `tfsdk:"sparse_checkout" tf:"optional"`
	// URL of the Git repository to be linked.
	Url types.String `tfsdk:"url" tf:""`
}

type CreateScope struct {
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	BackendAzureKeyvault *AzureKeyVaultSecretScopeMetadata `tfsdk:"backend_azure_keyvault" tf:"optional"`
	// The principal that is initially granted `MANAGE` permission to the
	// created scope.
	InitialManagePrincipal types.String `tfsdk:"initial_manage_principal" tf:"optional"`
	// Scope name requested by the user. Scope names are unique.
	Scope types.String `tfsdk:"scope" tf:""`
	// The backend type the scope will be created with. If not specified, will
	// default to `DATABRICKS`
	ScopeBackendType ScopeBackendType `tfsdk:"scope_backend_type" tf:"optional"`
}

type CreateScopeResponse struct {
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	CredentialId types.Int64 `tfsdk:"credential_id" tf:"optional"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, gitHubOAuth, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider types.String `tfsdk:"git_provider" tf:"optional"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername types.String `tfsdk:"git_username" tf:"optional"`
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

type DeleteAcl struct {
	// The principal to remove an existing ACL from.
	Principal types.String `tfsdk:"principal" tf:""`
	// The name of the scope to remove permissions from.
	Scope types.String `tfsdk:"scope" tf:""`
}

type DeleteAclResponse struct {
}

// Delete a credential
type DeleteGitCredentialRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-" url:"-"`
}

// Delete a repo
type DeleteRepoRequest struct {
	// The ID for the corresponding repo to access.
	RepoId types.Int64 `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

type DeleteScope struct {
	// Name of the scope to delete.
	Scope types.String `tfsdk:"scope" tf:""`
}

type DeleteScopeResponse struct {
}

type DeleteSecret struct {
	// Name of the secret to delete.
	Key types.String `tfsdk:"key" tf:""`
	// The name of the scope that contains the secret to delete.
	Scope types.String `tfsdk:"scope" tf:""`
}

type DeleteSecretResponse struct {
}

type ExportFormat string

const ExportFormatAuto ExportFormat = `AUTO`

const ExportFormatDbc ExportFormat = `DBC`

const ExportFormatHtml ExportFormat = `HTML`

const ExportFormatJupyter ExportFormat = `JUPYTER`

const ExportFormatRMarkdown ExportFormat = `R_MARKDOWN`

const ExportFormatSource ExportFormat = `SOURCE`

// String representation for [fmt.Print]
func (f *ExportFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExportFormat) Set(v string) error {
	switch v {
	case `AUTO`, `DBC`, `HTML`, `JUPYTER`, `R_MARKDOWN`, `SOURCE`:
		*f = ExportFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO", "DBC", "HTML", "JUPYTER", "R_MARKDOWN", "SOURCE"`, v)
	}
}

// Type always returns ExportFormat to satisfy [pflag.Value] interface
func (f *ExportFormat) Type() string {
	return "ExportFormat"
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
	Format ExportFormat `tfsdk:"-" url:"format,omitempty"`
	// The absolute path of the object or directory. Exporting a directory is
	// only supported for the `DBC`, `SOURCE`, and `AUTO` format.
	Path types.String `tfsdk:"-" url:"path"`
}

type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	Content types.String `tfsdk:"content" tf:"optional"`
	// The file type of the exported file.
	FileType types.String `tfsdk:"file_type" tf:"optional"`
}

// Get secret ACL details
type GetAclRequest struct {
	// The principal to fetch ACL information for.
	Principal types.String `tfsdk:"-" url:"principal"`
	// The name of the scope to fetch ACL information from.
	Scope types.String `tfsdk:"-" url:"scope"`
}

type GetCredentialsResponse struct {
	Credentials []CredentialInfo `tfsdk:"credentials" tf:"optional"`
}

// Get a credential entry
type GetGitCredentialRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-" url:"-"`
}

// Get repo permission levels
type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-" url:"-"`
}

type GetRepoPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []RepoPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Get repo permissions
type GetRepoPermissionsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-" url:"-"`
}

// Get a repo
type GetRepoRequest struct {
	// The ID for the corresponding repo to access.
	RepoId types.Int64 `tfsdk:"-" url:"-"`
}

// Get a secret
type GetSecretRequest struct {
	// The key to fetch secret for.
	Key types.String `tfsdk:"-" url:"key"`
	// The name of the scope to fetch secret information from.
	Scope types.String `tfsdk:"-" url:"scope"`
}

type GetSecretResponse struct {
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The value of the secret in its byte representation.
	Value types.String `tfsdk:"value" tf:"optional"`
}

// Get status
type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"-" url:"path"`
}

// Get workspace object permission levels
type GetWorkspaceObjectPermissionLevelsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-" url:"-"`
}

type GetWorkspaceObjectPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []WorkspaceObjectPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Get workspace object permissions
type GetWorkspaceObjectPermissionsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-" url:"-"`
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
	Format ImportFormat `tfsdk:"format" tf:"optional"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language Language `tfsdk:"language" tf:"optional"`
	// The flag that specifies whether to overwrite existing object. It is
	// `false` by default. For `DBC` format, `overwrite` is not supported since
	// it may contain a directory.
	Overwrite types.Bool `tfsdk:"overwrite" tf:"optional"`
	// The absolute path of the object or directory. Importing a directory is
	// only supported for the `DBC` and `SOURCE` formats.
	Path types.String `tfsdk:"path" tf:""`
}

// This specifies the format of the file to be imported.
//
// The value is case sensitive.
//
// - `AUTO`: The item is imported depending on an analysis of the item's
// extension and the header content provided in the request. If the item is
// imported as a notebook, then the item's extension is automatically removed. -
// `SOURCE`: The notebook or directory is imported as source code. - `HTML`: The
// notebook is imported as an HTML file. - `JUPYTER`: The notebook is imported
// as a Jupyter/IPython Notebook file. - `DBC`: The notebook is imported in
// Databricks archive format. Required for directories. - `R_MARKDOWN`: The
// notebook is imported from R Markdown format.
type ImportFormat string

// The item is imported depending on an analysis of the item's extension and
const ImportFormatAuto ImportFormat = `AUTO`

// The notebook is imported in <Databricks> archive format. Required for
// directories.
const ImportFormatDbc ImportFormat = `DBC`

// The notebook is imported as an HTML file.
const ImportFormatHtml ImportFormat = `HTML`

// The notebook is imported as a Jupyter/IPython Notebook file.
const ImportFormatJupyter ImportFormat = `JUPYTER`

// The notebook is imported from R Markdown format.
const ImportFormatRMarkdown ImportFormat = `R_MARKDOWN`

// The notebook or directory is imported as source code.
const ImportFormatSource ImportFormat = `SOURCE`

// String representation for [fmt.Print]
func (f *ImportFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ImportFormat) Set(v string) error {
	switch v {
	case `AUTO`, `DBC`, `HTML`, `JUPYTER`, `R_MARKDOWN`, `SOURCE`:
		*f = ImportFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO", "DBC", "HTML", "JUPYTER", "R_MARKDOWN", "SOURCE"`, v)
	}
}

// Type always returns ImportFormat to satisfy [pflag.Value] interface
func (f *ImportFormat) Type() string {
	return "ImportFormat"
}

type ImportResponse struct {
}

// The language of the object. This value is set only if the object type is
// `NOTEBOOK`.
type Language string

const LanguagePython Language = `PYTHON`

const LanguageR Language = `R`

const LanguageScala Language = `SCALA`

const LanguageSql Language = `SQL`

// String representation for [fmt.Print]
func (f *Language) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Language) Set(v string) error {
	switch v {
	case `PYTHON`, `R`, `SCALA`, `SQL`:
		*f = Language(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PYTHON", "R", "SCALA", "SQL"`, v)
	}
}

// Type always returns Language to satisfy [pflag.Value] interface
func (f *Language) Type() string {
	return "Language"
}

// Lists ACLs
type ListAclsRequest struct {
	// The name of the scope to fetch ACL information from.
	Scope types.String `tfsdk:"-" url:"scope"`
}

type ListAclsResponse struct {
	// The associated ACLs rule applied to principals in the given scope.
	Items []AclItem `tfsdk:"items" tf:"optional"`
}

// Get repos
type ListReposRequest struct {
	// Token used to get the next page of results. If not specified, returns the
	// first page of results as well as a next page token if there are more
	// results.
	NextPageToken types.String `tfsdk:"-" url:"next_page_token,omitempty"`
	// Filters repos that have paths starting with the given path prefix. If not
	// provided repos from /Repos will be served.
	PathPrefix types.String `tfsdk:"-" url:"path_prefix,omitempty"`
}

type ListReposResponse struct {
	// Token that can be specified as a query parameter to the GET /repos
	// endpoint to retrieve the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Repos []RepoInfo `tfsdk:"repos" tf:"optional"`
}

type ListResponse struct {
	// List of objects.
	Objects []ObjectInfo `tfsdk:"objects" tf:"optional"`
}

type ListScopesResponse struct {
	// The available secret scopes.
	Scopes []SecretScope `tfsdk:"scopes" tf:"optional"`
}

// List secret keys
type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	Scope types.String `tfsdk:"-" url:"scope"`
}

type ListSecretsResponse struct {
	// Metadata information of all secrets contained within the given scope.
	Secrets []SecretMetadata `tfsdk:"secrets" tf:"optional"`
}

// List contents
type ListWorkspaceRequest struct {
	// UTC timestamp in milliseconds
	NotebooksModifiedAfter types.Int64 `tfsdk:"-" url:"notebooks_modified_after,omitempty"`
	// The absolute path of the notebook or directory.
	Path types.String `tfsdk:"-" url:"path"`
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path types.String `tfsdk:"path" tf:""`
}

type MkdirsResponse struct {
}

type ObjectInfo struct {
	// Only applicable to files. The creation UTC timestamp.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language Language `tfsdk:"language" tf:"optional"`
	// Only applicable to files, the last modified UTC timestamp.
	ModifiedAt types.Int64 `tfsdk:"modified_at" tf:"optional"`
	// Unique identifier for the object.
	ObjectId types.Int64 `tfsdk:"object_id" tf:"optional"`
	// The type of the object in workspace.
	//
	// - `NOTEBOOK`: document that contains runnable code, visualizations, and
	// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
	// file - `REPO`: repository - `DASHBOARD`: Lakeview dashboard
	ObjectType ObjectType `tfsdk:"object_type" tf:"optional"`
	// The absolute path of the object.
	Path types.String `tfsdk:"path" tf:"optional"`
	// A unique identifier for the object that is consistent across all
	// Databricks APIs.
	ResourceId types.String `tfsdk:"resource_id" tf:"optional"`
	// Only applicable to files. The file size in bytes can be returned.
	Size types.Int64 `tfsdk:"size" tf:"optional"`
}

// The type of the object in workspace.
//
// - `NOTEBOOK`: document that contains runnable code, visualizations, and
// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
// file - `REPO`: repository - `DASHBOARD`: Lakeview dashboard
type ObjectType string

// Lakeview dashboard
const ObjectTypeDashboard ObjectType = `DASHBOARD`

// directory
const ObjectTypeDirectory ObjectType = `DIRECTORY`

// file
const ObjectTypeFile ObjectType = `FILE`

// library
const ObjectTypeLibrary ObjectType = `LIBRARY`

// document that contains runnable code, visualizations, and explanatory text.
const ObjectTypeNotebook ObjectType = `NOTEBOOK`

// repository
const ObjectTypeRepo ObjectType = `REPO`

// String representation for [fmt.Print]
func (f *ObjectType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ObjectType) Set(v string) error {
	switch v {
	case `DASHBOARD`, `DIRECTORY`, `FILE`, `LIBRARY`, `NOTEBOOK`, `REPO`:
		*f = ObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DASHBOARD", "DIRECTORY", "FILE", "LIBRARY", "NOTEBOOK", "REPO"`, v)
	}
}

// Type always returns ObjectType to satisfy [pflag.Value] interface
func (f *ObjectType) Type() string {
	return "ObjectType"
}

type PutAcl struct {
	// The permission level applied to the principal.
	Permission AclPermission `tfsdk:"permission" tf:""`
	// The principal in which the permission is applied.
	Principal types.String `tfsdk:"principal" tf:""`
	// The name of the scope to apply permissions to.
	Scope types.String `tfsdk:"scope" tf:""`
}

type PutAclResponse struct {
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

type PutSecretResponse struct {
}

type RepoAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel RepoPermissionLevel `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type RepoAccessControlResponse struct {
	// All permissions.
	AllPermissions []RepoPermission `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type RepoInfo struct {
	// Branch that the local version of the repo is checked out to.
	Branch types.String `tfsdk:"branch" tf:"optional"`
	// SHA-1 hash representing the commit ID of the current HEAD of the repo.
	HeadCommitId types.String `tfsdk:"head_commit_id" tf:"optional"`
	// ID of the repo object in the workspace.
	Id types.Int64 `tfsdk:"id" tf:"optional"`
	// Desired path for the repo in the workspace. Almost any path in the
	// workspace can be chosen. If repo is created in /Repos, path must be in
	// the format /Repos/{folder}/{repo-name}.
	Path types.String `tfsdk:"path" tf:"optional"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	Provider types.String `tfsdk:"provider" tf:"optional"`

	SparseCheckout *SparseCheckout `tfsdk:"sparse_checkout" tf:"optional"`
	// URL of the Git repository to be linked.
	Url types.String `tfsdk:"url" tf:"optional"`
}

type RepoPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel RepoPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

// Permission level
type RepoPermissionLevel string

const RepoPermissionLevelCanEdit RepoPermissionLevel = `CAN_EDIT`

const RepoPermissionLevelCanManage RepoPermissionLevel = `CAN_MANAGE`

const RepoPermissionLevelCanRead RepoPermissionLevel = `CAN_READ`

const RepoPermissionLevelCanRun RepoPermissionLevel = `CAN_RUN`

// String representation for [fmt.Print]
func (f *RepoPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RepoPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`, `CAN_RUN`:
		*f = RepoPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ", "CAN_RUN"`, v)
	}
}

// Type always returns RepoPermissionLevel to satisfy [pflag.Value] interface
func (f *RepoPermissionLevel) Type() string {
	return "RepoPermissionLevel"
}

type RepoPermissions struct {
	AccessControlList []RepoAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type RepoPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel RepoPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

type RepoPermissionsRequest struct {
	AccessControlList []RepoAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The repo for which to get or manage permissions.
	RepoId types.String `tfsdk:"-" url:"-"`
}

type ScopeBackendType string

const ScopeBackendTypeAzureKeyvault ScopeBackendType = `AZURE_KEYVAULT`

const ScopeBackendTypeDatabricks ScopeBackendType = `DATABRICKS`

// String representation for [fmt.Print]
func (f *ScopeBackendType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ScopeBackendType) Set(v string) error {
	switch v {
	case `AZURE_KEYVAULT`, `DATABRICKS`:
		*f = ScopeBackendType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AZURE_KEYVAULT", "DATABRICKS"`, v)
	}
}

// Type always returns ScopeBackendType to satisfy [pflag.Value] interface
func (f *ScopeBackendType) Type() string {
	return "ScopeBackendType"
}

type SecretMetadata struct {
	// A unique name to identify the secret.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The last updated timestamp (in milliseconds) for the secret.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
}

type SecretScope struct {
	// The type of secret scope backend.
	BackendType ScopeBackendType `tfsdk:"backend_type" tf:"optional"`
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata `tfsdk:"keyvault_metadata" tf:"optional"`
	// A unique name to identify the secret scope.
	Name types.String `tfsdk:"name" tf:"optional"`
}

type SparseCheckout struct {
	// List of patterns to include for sparse checkout.
	Patterns []types.String `tfsdk:"patterns" tf:"optional"`
}

type SparseCheckoutUpdate struct {
	// List of patterns to include for sparse checkout.
	Patterns []types.String `tfsdk:"patterns" tf:"optional"`
}

type UpdateCredentials struct {
	// The ID for the corresponding credential to access.
	CredentialId types.Int64 `tfsdk:"-" url:"-"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider types.String `tfsdk:"git_provider" tf:"optional"`
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
	// scoped access tokens. [Learn more]. The personal access token used to
	// authenticate to the corresponding Git
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken types.String `tfsdk:"personal_access_token" tf:"optional"`
}

type UpdateRepo struct {
	// Branch that the local version of the repo is checked out to.
	Branch types.String `tfsdk:"branch" tf:"optional"`
	// The ID for the corresponding repo to access.
	RepoId types.Int64 `tfsdk:"-" url:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	SparseCheckout *SparseCheckoutUpdate `tfsdk:"sparse_checkout" tf:"optional"`
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	Tag types.String `tfsdk:"tag" tf:"optional"`
}

type UpdateResponse struct {
}

type WorkspaceObjectAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type WorkspaceObjectAccessControlResponse struct {
	// All permissions.
	AllPermissions []WorkspaceObjectPermission `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type WorkspaceObjectPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

// Permission level
type WorkspaceObjectPermissionLevel string

const WorkspaceObjectPermissionLevelCanEdit WorkspaceObjectPermissionLevel = `CAN_EDIT`

const WorkspaceObjectPermissionLevelCanManage WorkspaceObjectPermissionLevel = `CAN_MANAGE`

const WorkspaceObjectPermissionLevelCanRead WorkspaceObjectPermissionLevel = `CAN_READ`

const WorkspaceObjectPermissionLevelCanRun WorkspaceObjectPermissionLevel = `CAN_RUN`

// String representation for [fmt.Print]
func (f *WorkspaceObjectPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceObjectPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`, `CAN_RUN`:
		*f = WorkspaceObjectPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ", "CAN_RUN"`, v)
	}
}

// Type always returns WorkspaceObjectPermissionLevel to satisfy [pflag.Value] interface
func (f *WorkspaceObjectPermissionLevel) Type() string {
	return "WorkspaceObjectPermissionLevel"
}

type WorkspaceObjectPermissions struct {
	AccessControlList []WorkspaceObjectAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type WorkspaceObjectPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

type WorkspaceObjectPermissionsRequest struct {
	AccessControlList []WorkspaceObjectAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId types.String `tfsdk:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType types.String `tfsdk:"-" url:"-"`
}
