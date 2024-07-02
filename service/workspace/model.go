// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AclItem struct {
	// The permission level applied to the principal.
	Permission AclPermission `tfsdk:"permission"`
	// The principal in which the permission is applied.
	Principal string `tfsdk:"principal"`
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
	DnsName string `tfsdk:"dns_name"`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	ResourceId string `tfsdk:"resource_id"`
}

type CreateCredentials struct {
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider string `tfsdk:"git_provider"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername string `tfsdk:"git_username"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more]. The personal access token used to
	// authenticate to the corresponding Git
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken string `tfsdk:"personal_access_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateCredentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCredentialsResponse struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `tfsdk:"credential_id"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider string `tfsdk:"git_provider"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername string `tfsdk:"git_username"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRepo struct {
	// Desired path for the repo in the workspace. Almost any path in the
	// workspace can be chosen. If repo is created in /Repos, path must be in
	// the format /Repos/{folder}/{repo-name}.
	Path string `tfsdk:"path"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	Provider string `tfsdk:"provider"`
	// If specified, the repo will be created with sparse checkout enabled. You
	// cannot enable/disable sparse checkout after the repo is created.
	SparseCheckout *SparseCheckout `tfsdk:"sparse_checkout"`
	// URL of the Git repository to be linked.
	Url string `tfsdk:"url"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateRepo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRepo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateScope struct {
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	BackendAzureKeyvault *AzureKeyVaultSecretScopeMetadata `tfsdk:"backend_azure_keyvault"`
	// The principal that is initially granted `MANAGE` permission to the
	// created scope.
	InitialManagePrincipal string `tfsdk:"initial_manage_principal"`
	// Scope name requested by the user. Scope names are unique.
	Scope string `tfsdk:"scope"`
	// The backend type the scope will be created with. If not specified, will
	// default to `DATABRICKS`
	ScopeBackendType ScopeBackendType `tfsdk:"scope_backend_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateScope) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateScope) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateScopeResponse struct {
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `tfsdk:"credential_id"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, gitHubOAuth, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider string `tfsdk:"git_provider"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername string `tfsdk:"git_username"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Delete struct {
	// The absolute path of the notebook or directory.
	Path string `tfsdk:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive bool `tfsdk:"recursive"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Delete) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Delete) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteAcl struct {
	// The principal to remove an existing ACL from.
	Principal string `tfsdk:"principal"`
	// The name of the scope to remove permissions from.
	Scope string `tfsdk:"scope"`
}

type DeleteAclResponse struct {
}

// Delete a credential
type DeleteGitCredentialRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `tfsdk:"-" url:"-"`
}

// Delete a repo
type DeleteRepoRequest struct {
	// The ID for the corresponding repo to access.
	RepoId int64 `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

type DeleteScope struct {
	// Name of the scope to delete.
	Scope string `tfsdk:"scope"`
}

type DeleteScopeResponse struct {
}

type DeleteSecret struct {
	// Name of the secret to delete.
	Key string `tfsdk:"key"`
	// The name of the scope that contains the secret to delete.
	Scope string `tfsdk:"scope"`
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
	Path string `tfsdk:"-" url:"path"`
}

type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	Content string `tfsdk:"content"`
	// The file type of the exported file.
	FileType string `tfsdk:"file_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ExportResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExportResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get secret ACL details
type GetAclRequest struct {
	// The principal to fetch ACL information for.
	Principal string `tfsdk:"-" url:"principal"`
	// The name of the scope to fetch ACL information from.
	Scope string `tfsdk:"-" url:"scope"`
}

type GetCredentialsResponse struct {
	Credentials []CredentialInfo `tfsdk:"credentials"`
}

// Get a credential entry
type GetGitCredentialRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `tfsdk:"-" url:"-"`
}

// Get repo permission levels
type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId string `tfsdk:"-" url:"-"`
}

type GetRepoPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []RepoPermissionsDescription `tfsdk:"permission_levels"`
}

// Get repo permissions
type GetRepoPermissionsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId string `tfsdk:"-" url:"-"`
}

// Get a repo
type GetRepoRequest struct {
	// The ID for the corresponding repo to access.
	RepoId int64 `tfsdk:"-" url:"-"`
}

// Get a secret
type GetSecretRequest struct {
	// The key to fetch secret for.
	Key string `tfsdk:"-" url:"key"`
	// The name of the scope to fetch secret information from.
	Scope string `tfsdk:"-" url:"scope"`
}

type GetSecretResponse struct {
	// A unique name to identify the secret.
	Key string `tfsdk:"key"`
	// The value of the secret in its byte representation.
	Value string `tfsdk:"value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetSecretResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSecretResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get status
type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	Path string `tfsdk:"-" url:"path"`
}

// Get workspace object permission levels
type GetWorkspaceObjectPermissionLevelsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `tfsdk:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `tfsdk:"-" url:"-"`
}

type GetWorkspaceObjectPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []WorkspaceObjectPermissionsDescription `tfsdk:"permission_levels"`
}

// Get workspace object permissions
type GetWorkspaceObjectPermissionsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `tfsdk:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `tfsdk:"-" url:"-"`
}

type Import struct {
	// The base64-encoded content. This has a limit of 10 MB.
	//
	// If the limit (10MB) is exceeded, exception with error code
	// **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown. This parameter might be absent,
	// and instead a posted file is used.
	Content string `tfsdk:"content"`
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
	Format ImportFormat `tfsdk:"format"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language Language `tfsdk:"language"`
	// The flag that specifies whether to overwrite existing object. It is
	// `false` by default. For `DBC` format, `overwrite` is not supported since
	// it may contain a directory.
	Overwrite bool `tfsdk:"overwrite"`
	// The absolute path of the object or directory. Importing a directory is
	// only supported for the `DBC` and `SOURCE` formats.
	Path string `tfsdk:"path"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Import) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Import) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Scope string `tfsdk:"-" url:"scope"`
}

type ListAclsResponse struct {
	// The associated ACLs rule applied to principals in the given scope.
	Items []AclItem `tfsdk:"items"`
}

// Get repos
type ListReposRequest struct {
	// Token used to get the next page of results. If not specified, returns the
	// first page of results as well as a next page token if there are more
	// results.
	NextPageToken string `tfsdk:"-" url:"next_page_token,omitempty"`
	// Filters repos that have paths starting with the given path prefix. If not
	// provided repos from /Repos will be served.
	PathPrefix string `tfsdk:"-" url:"path_prefix,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListReposRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListReposRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListReposResponse struct {
	// Token that can be specified as a query parameter to the GET /repos
	// endpoint to retrieve the next page of results.
	NextPageToken string `tfsdk:"next_page_token"`

	Repos []RepoInfo `tfsdk:"repos"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListReposResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListReposResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListResponse struct {
	// List of objects.
	Objects []ObjectInfo `tfsdk:"objects"`
}

type ListScopesResponse struct {
	// The available secret scopes.
	Scopes []SecretScope `tfsdk:"scopes"`
}

// List secret keys
type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	Scope string `tfsdk:"-" url:"scope"`
}

type ListSecretsResponse struct {
	// Metadata information of all secrets contained within the given scope.
	Secrets []SecretMetadata `tfsdk:"secrets"`
}

// List contents
type ListWorkspaceRequest struct {
	// UTC timestamp in milliseconds
	NotebooksModifiedAfter int `tfsdk:"-" url:"notebooks_modified_after,omitempty"`
	// The absolute path of the notebook or directory.
	Path string `tfsdk:"-" url:"path"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path string `tfsdk:"path"`
}

type MkdirsResponse struct {
}

type ObjectInfo struct {
	// Only applicable to files. The creation UTC timestamp.
	CreatedAt int64 `tfsdk:"created_at"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language Language `tfsdk:"language"`
	// Only applicable to files, the last modified UTC timestamp.
	ModifiedAt int64 `tfsdk:"modified_at"`
	// Unique identifier for the object.
	ObjectId int64 `tfsdk:"object_id"`
	// The type of the object in workspace.
	//
	// - `NOTEBOOK`: document that contains runnable code, visualizations, and
	// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
	// file - `REPO`: repository - `DASHBOARD`: Lakeview dashboard
	ObjectType ObjectType `tfsdk:"object_type"`
	// The absolute path of the object.
	Path string `tfsdk:"path"`
	// A unique identifier for the object that is consistent across all
	// Databricks APIs.
	ResourceId string `tfsdk:"resource_id"`
	// Only applicable to files. The file size in bytes can be returned.
	Size int64 `tfsdk:"size"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ObjectInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ObjectInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Permission AclPermission `tfsdk:"permission"`
	// The principal in which the permission is applied.
	Principal string `tfsdk:"principal"`
	// The name of the scope to apply permissions to.
	Scope string `tfsdk:"scope"`
}

type PutAclResponse struct {
}

type PutSecret struct {
	// If specified, value will be stored as bytes.
	BytesValue string `tfsdk:"bytes_value"`
	// A unique name to identify the secret.
	Key string `tfsdk:"key"`
	// The name of the scope to which the secret will be associated with.
	Scope string `tfsdk:"scope"`
	// If specified, note that the value will be stored in UTF-8 (MB4) form.
	StringValue string `tfsdk:"string_value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PutSecret) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PutSecret) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PutSecretResponse struct {
}

type RepoAccessControlRequest struct {
	// name of the group
	GroupName string `tfsdk:"group_name"`
	// Permission level
	PermissionLevel RepoPermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// name of the user
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RepoAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepoAccessControlResponse struct {
	// All permissions.
	AllPermissions []RepoPermission `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName string `tfsdk:"display_name"`
	// name of the group
	GroupName string `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// name of the user
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RepoAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepoInfo struct {
	// Branch that the local version of the repo is checked out to.
	Branch string `tfsdk:"branch"`
	// SHA-1 hash representing the commit ID of the current HEAD of the repo.
	HeadCommitId string `tfsdk:"head_commit_id"`
	// ID of the repo object in the workspace.
	Id int64 `tfsdk:"id"`
	// Desired path for the repo in the workspace. Almost any path in the
	// workspace can be chosen. If repo is created in /Repos, path must be in
	// the format /Repos/{folder}/{repo-name}.
	Path string `tfsdk:"path"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	Provider string `tfsdk:"provider"`

	SparseCheckout *SparseCheckout `tfsdk:"sparse_checkout"`
	// URL of the Git repository to be linked.
	Url string `tfsdk:"url"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RepoInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepoPermission struct {
	Inherited bool `tfsdk:"inherited"`

	InheritedFromObject []string `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel RepoPermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RepoPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	AccessControlList []RepoAccessControlResponse `tfsdk:"access_control_list"`

	ObjectId string `tfsdk:"object_id"`

	ObjectType string `tfsdk:"object_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RepoPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepoPermissionsDescription struct {
	Description string `tfsdk:"description"`
	// Permission level
	PermissionLevel RepoPermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RepoPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepoPermissionsRequest struct {
	AccessControlList []RepoAccessControlRequest `tfsdk:"access_control_list"`
	// The repo for which to get or manage permissions.
	RepoId string `tfsdk:"-" url:"-"`
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
	Key string `tfsdk:"key"`
	// The last updated timestamp (in milliseconds) for the secret.
	LastUpdatedTimestamp int64 `tfsdk:"last_updated_timestamp"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *SecretMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SecretMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SecretScope struct {
	// The type of secret scope backend.
	BackendType ScopeBackendType `tfsdk:"backend_type"`
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata `tfsdk:"keyvault_metadata"`
	// A unique name to identify the secret scope.
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *SecretScope) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SecretScope) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SparseCheckout struct {
	// List of patterns to include for sparse checkout.
	Patterns []string `tfsdk:"patterns"`
}

type SparseCheckoutUpdate struct {
	// List of patterns to include for sparse checkout.
	Patterns []string `tfsdk:"patterns"`
}

type UpdateCredentials struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `tfsdk:"-" url:"-"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider string `tfsdk:"git_provider"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername string `tfsdk:"git_username"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more]. The personal access token used to
	// authenticate to the corresponding Git
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken string `tfsdk:"personal_access_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateCredentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCredentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRepo struct {
	// Branch that the local version of the repo is checked out to.
	Branch string `tfsdk:"branch"`
	// The ID for the corresponding repo to access.
	RepoId int64 `tfsdk:"-" url:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	SparseCheckout *SparseCheckoutUpdate `tfsdk:"sparse_checkout"`
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	Tag string `tfsdk:"tag"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateRepo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRepo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateResponse struct {
}

type WorkspaceObjectAccessControlRequest struct {
	// name of the group
	GroupName string `tfsdk:"group_name"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// name of the user
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *WorkspaceObjectAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceObjectAccessControlResponse struct {
	// All permissions.
	AllPermissions []WorkspaceObjectPermission `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName string `tfsdk:"display_name"`
	// name of the group
	GroupName string `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// name of the user
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *WorkspaceObjectAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceObjectPermission struct {
	Inherited bool `tfsdk:"inherited"`

	InheritedFromObject []string `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *WorkspaceObjectPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	AccessControlList []WorkspaceObjectAccessControlResponse `tfsdk:"access_control_list"`

	ObjectId string `tfsdk:"object_id"`

	ObjectType string `tfsdk:"object_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *WorkspaceObjectPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceObjectPermissionsDescription struct {
	Description string `tfsdk:"description"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *WorkspaceObjectPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceObjectPermissionsRequest struct {
	AccessControlList []WorkspaceObjectAccessControlRequest `tfsdk:"access_control_list"`
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `tfsdk:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `tfsdk:"-" url:"-"`
}
