// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccessControlRequest struct {
	// name of the group
	GroupName string `tfsdk:"group_name"`
	// Permission level
	PermissionLevel PermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// name of the user
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AccessControlResponse struct {
	// All permissions.
	AllPermissions []Permission `tfsdk:"all_permissions"`
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

func (s *AccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ComplexValue struct {
	Display string `tfsdk:"display"`

	Primary bool `tfsdk:"primary"`

	Ref string `tfsdk:"$ref"`

	Type string `tfsdk:"type"`

	Value string `tfsdk:"value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ComplexValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplexValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a group
type DeleteAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id string `tfsdk:"-" url:"-"`
}

// Delete a service principal
type DeleteAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `tfsdk:"-" url:"-"`
}

// Delete a user
type DeleteAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	Id string `tfsdk:"-" url:"-"`
}

// Delete a group
type DeleteGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete a service principal
type DeleteServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `tfsdk:"-" url:"-"`
}

// Delete a user
type DeleteUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `tfsdk:"-" url:"-"`
}

// Delete permissions assignment
type DeleteWorkspaceAssignmentRequest struct {
	// The ID of the user, service principal, or group.
	PrincipalId int64 `tfsdk:"-" url:"-"`
	// The workspace ID.
	WorkspaceId int64 `tfsdk:"-" url:"-"`
}

type DeleteWorkspaceAssignments struct {
}

// Get group details
type GetAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id string `tfsdk:"-" url:"-"`
}

// Get service principal details
type GetAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `tfsdk:"-" url:"-"`
}

// Get user details
type GetAccountUserRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `tfsdk:"-" url:"attributes,omitempty"`
	// Desired number of results per page. Default is 10000.
	Count int `tfsdk:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `tfsdk:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `tfsdk:"-" url:"filter,omitempty"`
	// Unique ID for a user in the Databricks account.
	Id string `tfsdk:"-" url:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `tfsdk:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder GetSortOrder `tfsdk:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `tfsdk:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetAccountUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAccountUserRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get assignable roles for a resource
type GetAssignableRolesForResourceRequest struct {
	// The resource name for which assignable roles will be listed.
	Resource string `tfsdk:"-" url:"resource"`
}

type GetAssignableRolesForResourceResponse struct {
	Roles []Role `tfsdk:"roles"`
}

// Get group details
type GetGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `tfsdk:"-" url:"-"`
}

type GetPasswordPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PasswordPermissionsDescription `tfsdk:"permission_levels"`
}

// Get object permission levels
type GetPermissionLevelsRequest struct {
	// <needs content>
	RequestObjectId string `tfsdk:"-" url:"-"`
	// <needs content>
	RequestObjectType string `tfsdk:"-" url:"-"`
}

type GetPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PermissionsDescription `tfsdk:"permission_levels"`
}

// Get object permissions
type GetPermissionRequest struct {
	// The id of the request object.
	RequestObjectId string `tfsdk:"-" url:"-"`
	// The type of the request object. Can be one of the following:
	// authorization, clusters, cluster-policies, directories, experiments,
	// files, instance-pools, jobs, notebooks, pipelines, registered-models,
	// repos, serving-endpoints, or warehouses.
	RequestObjectType string `tfsdk:"-" url:"-"`
}

// Get a rule set
type GetRuleSetRequest struct {
	// Etag used for versioning. The response is at least as fresh as the eTag
	// provided. Etag is used for optimistic concurrency control as a way to
	// help prevent simultaneous updates of a rule set from overwriting each
	// other. It is strongly suggested that systems make use of the etag in the
	// read -> modify -> write pattern to perform rule set updates in order to
	// avoid race conditions that is get an etag from a GET rule set request,
	// and pass it with the PUT update request to identify the rule set version
	// you are updating.
	Etag string `tfsdk:"-" url:"etag"`
	// The ruleset name associated with the request.
	Name string `tfsdk:"-" url:"name"`
}

// Get service principal details
type GetServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `tfsdk:"-" url:"-"`
}

type GetSortOrder string

const GetSortOrderAscending GetSortOrder = `ascending`

const GetSortOrderDescending GetSortOrder = `descending`

// String representation for [fmt.Print]
func (f *GetSortOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetSortOrder) Set(v string) error {
	switch v {
	case `ascending`, `descending`:
		*f = GetSortOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ascending", "descending"`, v)
	}
}

// Type always returns GetSortOrder to satisfy [pflag.Value] interface
func (f *GetSortOrder) Type() string {
	return "GetSortOrder"
}

// Get user details
type GetUserRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `tfsdk:"-" url:"attributes,omitempty"`
	// Desired number of results per page.
	Count int `tfsdk:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `tfsdk:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `tfsdk:"-" url:"filter,omitempty"`
	// Unique ID for a user in the Databricks workspace.
	Id string `tfsdk:"-" url:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `tfsdk:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder GetSortOrder `tfsdk:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `tfsdk:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *GetUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetUserRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List workspace permissions
type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	WorkspaceId int64 `tfsdk:"-" url:"-"`
}

type GrantRule struct {
	// Principals this grant rule applies to.
	Principals []string `tfsdk:"principals"`
	// Role that is assigned to the list of principals.
	Role string `tfsdk:"role"`
}

type Group struct {
	// String that represents a human-readable group name
	DisplayName string `tfsdk:"displayName"`
	// Entitlements assigned to the group. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements []ComplexValue `tfsdk:"entitlements"`

	ExternalId string `tfsdk:"externalId"`

	Groups []ComplexValue `tfsdk:"groups"`
	// Databricks group ID
	Id string `tfsdk:"id" url:"-"`

	Members []ComplexValue `tfsdk:"members"`
	// Container for the group identifier. Workspace local versus account.
	Meta *ResourceMeta `tfsdk:"meta"`
	// Corresponds to AWS instance profile/arn role.
	Roles []ComplexValue `tfsdk:"roles"`
	// The schema of the group.
	Schemas []GroupSchema `tfsdk:"schemas"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Group) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Group) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GroupSchema string

const GroupSchemaUrnIetfParamsScimSchemasCore20Group GroupSchema = `urn:ietf:params:scim:schemas:core:2.0:Group`

// String representation for [fmt.Print]
func (f *GroupSchema) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GroupSchema) Set(v string) error {
	switch v {
	case `urn:ietf:params:scim:schemas:core:2.0:Group`:
		*f = GroupSchema(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "urn:ietf:params:scim:schemas:core:2.0:Group"`, v)
	}
}

// Type always returns GroupSchema to satisfy [pflag.Value] interface
func (f *GroupSchema) Type() string {
	return "GroupSchema"
}

// List group details
type ListAccountGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `tfsdk:"-" url:"attributes,omitempty"`
	// Desired number of results per page. Default is 10000.
	Count int64 `tfsdk:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `tfsdk:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `tfsdk:"-" url:"filter,omitempty"`
	// Attribute to sort the results.
	SortBy string `tfsdk:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `tfsdk:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int64 `tfsdk:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAccountGroupsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountGroupsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List service principals
type ListAccountServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `tfsdk:"-" url:"attributes,omitempty"`
	// Desired number of results per page. Default is 10000.
	Count int64 `tfsdk:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `tfsdk:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `tfsdk:"-" url:"filter,omitempty"`
	// Attribute to sort the results.
	SortBy string `tfsdk:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `tfsdk:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int64 `tfsdk:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAccountServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List users
type ListAccountUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `tfsdk:"-" url:"attributes,omitempty"`
	// Desired number of results per page. Default is 10000.
	Count int64 `tfsdk:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `tfsdk:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `tfsdk:"-" url:"filter,omitempty"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `tfsdk:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `tfsdk:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int64 `tfsdk:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAccountUsersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountUsersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List group details
type ListGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `tfsdk:"-" url:"attributes,omitempty"`
	// Desired number of results per page.
	Count int64 `tfsdk:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `tfsdk:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `tfsdk:"-" url:"filter,omitempty"`
	// Attribute to sort the results.
	SortBy string `tfsdk:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `tfsdk:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int64 `tfsdk:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListGroupsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGroupsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListGroupsResponse struct {
	// Total results returned in the response.
	ItemsPerPage int64 `tfsdk:"itemsPerPage"`
	// User objects returned in the response.
	Resources []Group `tfsdk:"Resources"`
	// The schema of the service principal.
	Schemas []ListResponseSchema `tfsdk:"schemas"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `tfsdk:"startIndex"`
	// Total results that match the request filters.
	TotalResults int64 `tfsdk:"totalResults"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListGroupsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGroupsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListResponseSchema string

const ListResponseSchemaUrnIetfParamsScimApiMessages20ListResponse ListResponseSchema = `urn:ietf:params:scim:api:messages:2.0:ListResponse`

// String representation for [fmt.Print]
func (f *ListResponseSchema) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListResponseSchema) Set(v string) error {
	switch v {
	case `urn:ietf:params:scim:api:messages:2.0:ListResponse`:
		*f = ListResponseSchema(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "urn:ietf:params:scim:api:messages:2.0:ListResponse"`, v)
	}
}

// Type always returns ListResponseSchema to satisfy [pflag.Value] interface
func (f *ListResponseSchema) Type() string {
	return "ListResponseSchema"
}

type ListServicePrincipalResponse struct {
	// Total results returned in the response.
	ItemsPerPage int64 `tfsdk:"itemsPerPage"`
	// User objects returned in the response.
	Resources []ServicePrincipal `tfsdk:"Resources"`
	// The schema of the List response.
	Schemas []ListResponseSchema `tfsdk:"schemas"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `tfsdk:"startIndex"`
	// Total results that match the request filters.
	TotalResults int64 `tfsdk:"totalResults"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListServicePrincipalResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List service principals
type ListServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `tfsdk:"-" url:"attributes,omitempty"`
	// Desired number of results per page.
	Count int64 `tfsdk:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `tfsdk:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `tfsdk:"-" url:"filter,omitempty"`
	// Attribute to sort the results.
	SortBy string `tfsdk:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `tfsdk:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int64 `tfsdk:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSortOrder string

const ListSortOrderAscending ListSortOrder = `ascending`

const ListSortOrderDescending ListSortOrder = `descending`

// String representation for [fmt.Print]
func (f *ListSortOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListSortOrder) Set(v string) error {
	switch v {
	case `ascending`, `descending`:
		*f = ListSortOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ascending", "descending"`, v)
	}
}

// Type always returns ListSortOrder to satisfy [pflag.Value] interface
func (f *ListSortOrder) Type() string {
	return "ListSortOrder"
}

// List users
type ListUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `tfsdk:"-" url:"attributes,omitempty"`
	// Desired number of results per page.
	Count int64 `tfsdk:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `tfsdk:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `tfsdk:"-" url:"filter,omitempty"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `tfsdk:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `tfsdk:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int64 `tfsdk:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListUsersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUsersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListUsersResponse struct {
	// Total results returned in the response.
	ItemsPerPage int64 `tfsdk:"itemsPerPage"`
	// User objects returned in the response.
	Resources []User `tfsdk:"Resources"`
	// The schema of the List response.
	Schemas []ListResponseSchema `tfsdk:"schemas"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `tfsdk:"startIndex"`
	// Total results that match the request filters.
	TotalResults int64 `tfsdk:"totalResults"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListUsersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUsersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get permission assignments
type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	WorkspaceId int64 `tfsdk:"-" url:"-"`
}

type Name struct {
	// Family name of the Databricks user.
	FamilyName string `tfsdk:"familyName"`
	// Given name of the Databricks user.
	GivenName string `tfsdk:"givenName"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Name) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Name) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ObjectPermissions struct {
	AccessControlList []AccessControlResponse `tfsdk:"access_control_list"`

	ObjectId string `tfsdk:"object_id"`

	ObjectType string `tfsdk:"object_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ObjectPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ObjectPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PartialUpdate struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `tfsdk:"-" url:"-"`

	Operations []Patch `tfsdk:"Operations"`
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	Schemas []PatchSchema `tfsdk:"schemas"`
}

type PasswordAccessControlRequest struct {
	// name of the group
	GroupName string `tfsdk:"group_name"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// name of the user
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PasswordAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PasswordAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PasswordAccessControlResponse struct {
	// All permissions.
	AllPermissions []PasswordPermission `tfsdk:"all_permissions"`
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

func (s *PasswordAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PasswordAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PasswordPermission struct {
	Inherited bool `tfsdk:"inherited"`

	InheritedFromObject []string `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PasswordPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PasswordPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type PasswordPermissionLevel string

const PasswordPermissionLevelCanUse PasswordPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (f *PasswordPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PasswordPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_USE`:
		*f = PasswordPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_USE"`, v)
	}
}

// Type always returns PasswordPermissionLevel to satisfy [pflag.Value] interface
func (f *PasswordPermissionLevel) Type() string {
	return "PasswordPermissionLevel"
}

type PasswordPermissions struct {
	AccessControlList []PasswordAccessControlResponse `tfsdk:"access_control_list"`

	ObjectId string `tfsdk:"object_id"`

	ObjectType string `tfsdk:"object_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PasswordPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PasswordPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PasswordPermissionsDescription struct {
	Description string `tfsdk:"description"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PasswordPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PasswordPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PasswordPermissionsRequest struct {
	AccessControlList []PasswordAccessControlRequest `tfsdk:"access_control_list"`
}

type Patch struct {
	// Type of patch operation.
	Op PatchOp `tfsdk:"op"`
	// Selection of patch operation
	Path string `tfsdk:"path"`
	// Value to modify
	Value any `tfsdk:"value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Patch) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Patch) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Type of patch operation.
type PatchOp string

const PatchOpAdd PatchOp = `add`

const PatchOpRemove PatchOp = `remove`

const PatchOpReplace PatchOp = `replace`

// String representation for [fmt.Print]
func (f *PatchOp) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PatchOp) Set(v string) error {
	switch v {
	case `add`, `remove`, `replace`:
		*f = PatchOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "add", "remove", "replace"`, v)
	}
}

// Type always returns PatchOp to satisfy [pflag.Value] interface
func (f *PatchOp) Type() string {
	return "PatchOp"
}

type PatchResponse struct {
}

type PatchSchema string

const PatchSchemaUrnIetfParamsScimApiMessages20PatchOp PatchSchema = `urn:ietf:params:scim:api:messages:2.0:PatchOp`

// String representation for [fmt.Print]
func (f *PatchSchema) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PatchSchema) Set(v string) error {
	switch v {
	case `urn:ietf:params:scim:api:messages:2.0:PatchOp`:
		*f = PatchSchema(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "urn:ietf:params:scim:api:messages:2.0:PatchOp"`, v)
	}
}

// Type always returns PatchSchema to satisfy [pflag.Value] interface
func (f *PatchSchema) Type() string {
	return "PatchSchema"
}

type Permission struct {
	Inherited bool `tfsdk:"inherited"`

	InheritedFromObject []string `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel PermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *Permission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Permission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionAssignment struct {
	// Error response associated with a workspace permission assignment, if any.
	Error string `tfsdk:"error"`
	// The permissions level of the principal.
	Permissions []WorkspacePermission `tfsdk:"permissions"`
	// Information about the principal assigned to the workspace.
	Principal *PrincipalOutput `tfsdk:"principal"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PermissionAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionAssignments struct {
	// Array of permissions assignments defined for a workspace.
	PermissionAssignments []PermissionAssignment `tfsdk:"permission_assignments"`
}

// Permission level
type PermissionLevel string

const PermissionLevelCanAttachTo PermissionLevel = `CAN_ATTACH_TO`

const PermissionLevelCanBind PermissionLevel = `CAN_BIND`

const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

const PermissionLevelCanEditMetadata PermissionLevel = `CAN_EDIT_METADATA`

const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageRun PermissionLevel = `CAN_MANAGE_RUN`

const PermissionLevelCanManageStagingVersions PermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanQuery PermissionLevel = `CAN_QUERY`

const PermissionLevelCanRead PermissionLevel = `CAN_READ`

const PermissionLevelCanRestart PermissionLevel = `CAN_RESTART`

const PermissionLevelCanRun PermissionLevel = `CAN_RUN`

const PermissionLevelCanUse PermissionLevel = `CAN_USE`

const PermissionLevelCanView PermissionLevel = `CAN_VIEW`

const PermissionLevelCanViewMetadata PermissionLevel = `CAN_VIEW_METADATA`

const PermissionLevelIsOwner PermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (f *PermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PermissionLevel) Set(v string) error {
	switch v {
	case `CAN_ATTACH_TO`, `CAN_BIND`, `CAN_EDIT`, `CAN_EDIT_METADATA`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_RUN`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_QUERY`, `CAN_READ`, `CAN_RESTART`, `CAN_RUN`, `CAN_USE`, `CAN_VIEW`, `CAN_VIEW_METADATA`, `IS_OWNER`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_BIND", "CAN_EDIT", "CAN_EDIT_METADATA", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_RUN", "CAN_MANAGE_STAGING_VERSIONS", "CAN_QUERY", "CAN_READ", "CAN_RESTART", "CAN_RUN", "CAN_USE", "CAN_VIEW", "CAN_VIEW_METADATA", "IS_OWNER"`, v)
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

type PermissionMigrationRequest struct {
	// The name of the workspace group that permissions will be migrated from.
	FromWorkspaceGroupName string `tfsdk:"from_workspace_group_name"`
	// The maximum number of permissions that will be migrated.
	Size int `tfsdk:"size"`
	// The name of the account group that permissions will be migrated to.
	ToAccountGroupName string `tfsdk:"to_account_group_name"`
	// WorkspaceId of the associated workspace where the permission migration
	// will occur. Both workspace group and account group must be in this
	// workspace.
	WorkspaceId int64 `tfsdk:"workspace_id"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PermissionMigrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionMigrationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionMigrationResponse struct {
	// Number of permissions migrated.
	PermissionsMigrated int `tfsdk:"permissions_migrated"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PermissionMigrationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionMigrationResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionOutput struct {
	// The results of a permissions query.
	Description string `tfsdk:"description"`

	PermissionLevel WorkspacePermission `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PermissionOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionsDescription struct {
	Description string `tfsdk:"description"`
	// Permission level
	PermissionLevel PermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionsRequest struct {
	AccessControlList []AccessControlRequest `tfsdk:"access_control_list"`
	// The id of the request object.
	RequestObjectId string `tfsdk:"-" url:"-"`
	// The type of the request object. Can be one of the following:
	// authorization, clusters, cluster-policies, directories, experiments,
	// files, instance-pools, jobs, notebooks, pipelines, registered-models,
	// repos, serving-endpoints, or warehouses.
	RequestObjectType string `tfsdk:"-" url:"-"`
}

type PrincipalOutput struct {
	// The display name of the principal.
	DisplayName string `tfsdk:"display_name"`
	// The group name of the group. Present only if the principal is a group.
	GroupName string `tfsdk:"group_name"`
	// The unique, opaque id of the principal.
	PrincipalId int64 `tfsdk:"principal_id"`
	// The name of the service principal. Present only if the principal is a
	// service principal.
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// The username of the user. Present only if the principal is a user.
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PrincipalOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrincipalOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ResourceMeta struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	ResourceType string `tfsdk:"resourceType"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ResourceMeta) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResourceMeta) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Role struct {
	// Role to assign to a principal or a list of principals on a resource.
	Name string `tfsdk:"name"`
}

type RuleSetResponse struct {
	// Identifies the version of the rule set returned.
	Etag string `tfsdk:"etag"`

	GrantRules []GrantRule `tfsdk:"grant_rules"`
	// Name of the rule set.
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *RuleSetResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RuleSetResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RuleSetUpdateRequest struct {
	// The expected etag of the rule set to update. The update will fail if the
	// value does not match the value that is stored in account access control
	// service.
	Etag string `tfsdk:"etag"`

	GrantRules []GrantRule `tfsdk:"grant_rules"`
	// Name of the rule set.
	Name string `tfsdk:"name"`
}

type ServicePrincipal struct {
	// If this user is active
	Active bool `tfsdk:"active"`
	// UUID relating to the service principal
	ApplicationId string `tfsdk:"applicationId"`
	// String that represents a concatenation of given and family names.
	DisplayName string `tfsdk:"displayName"`
	// Entitlements assigned to the service principal. See [assigning
	// entitlements] for a full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements []ComplexValue `tfsdk:"entitlements"`

	ExternalId string `tfsdk:"externalId"`

	Groups []ComplexValue `tfsdk:"groups"`
	// Databricks service principal ID.
	Id string `tfsdk:"id" url:"-"`
	// Corresponds to AWS instance profile/arn role.
	Roles []ComplexValue `tfsdk:"roles"`
	// The schema of the List response.
	Schemas []ServicePrincipalSchema `tfsdk:"schemas"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServicePrincipal) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServicePrincipal) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServicePrincipalSchema string

const ServicePrincipalSchemaUrnIetfParamsScimSchemasCore20ServicePrincipal ServicePrincipalSchema = `urn:ietf:params:scim:schemas:core:2.0:ServicePrincipal`

// String representation for [fmt.Print]
func (f *ServicePrincipalSchema) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServicePrincipalSchema) Set(v string) error {
	switch v {
	case `urn:ietf:params:scim:schemas:core:2.0:ServicePrincipal`:
		*f = ServicePrincipalSchema(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "urn:ietf:params:scim:schemas:core:2.0:ServicePrincipal"`, v)
	}
}

// Type always returns ServicePrincipalSchema to satisfy [pflag.Value] interface
func (f *ServicePrincipalSchema) Type() string {
	return "ServicePrincipalSchema"
}

type UpdateResponse struct {
}

type UpdateRuleSetRequest struct {
	// Name of the rule set.
	Name string `tfsdk:"name"`

	RuleSet RuleSetUpdateRequest `tfsdk:"rule_set"`
}

type UpdateWorkspaceAssignments struct {
	// Array of permissions assignments to update on the workspace. Note that
	// excluding this field will have the same effect as providing an empty list
	// which will result in the deletion of all permissions for the principal.
	Permissions []WorkspacePermission `tfsdk:"permissions"`
	// The ID of the user, service principal, or group.
	PrincipalId int64 `tfsdk:"-" url:"-"`
	// The workspace ID.
	WorkspaceId int64 `tfsdk:"-" url:"-"`
}

type User struct {
	// If this user is active
	Active bool `tfsdk:"active"`
	// String that represents a concatenation of given and family names. For
	// example `John Smith`. This field cannot be updated through the Workspace
	// SCIM APIs when [identity federation is enabled]. Use Account SCIM APIs to
	// update `displayName`.
	//
	// [identity federation is enabled]: https://docs.databricks.com/administration-guide/users-groups/best-practices.html#enable-identity-federation
	DisplayName string `tfsdk:"displayName"`
	// All the emails associated with the Databricks user.
	Emails []ComplexValue `tfsdk:"emails"`
	// Entitlements assigned to the user. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements []ComplexValue `tfsdk:"entitlements"`
	// External ID is not currently supported. It is reserved for future use.
	ExternalId string `tfsdk:"externalId"`

	Groups []ComplexValue `tfsdk:"groups"`
	// Databricks user ID. This is automatically set by Databricks. Any value
	// provided by the client will be ignored.
	Id string `tfsdk:"id" url:"-"`

	Name *Name `tfsdk:"name"`
	// Corresponds to AWS instance profile/arn role.
	Roles []ComplexValue `tfsdk:"roles"`
	// The schema of the user.
	Schemas []UserSchema `tfsdk:"schemas"`
	// Email address of the Databricks user.
	UserName string `tfsdk:"userName"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *User) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s User) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UserSchema string

const UserSchemaUrnIetfParamsScimSchemasCore20User UserSchema = `urn:ietf:params:scim:schemas:core:2.0:User`

const UserSchemaUrnIetfParamsScimSchemasExtensionWorkspace20User UserSchema = `urn:ietf:params:scim:schemas:extension:workspace:2.0:User`

// String representation for [fmt.Print]
func (f *UserSchema) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UserSchema) Set(v string) error {
	switch v {
	case `urn:ietf:params:scim:schemas:core:2.0:User`, `urn:ietf:params:scim:schemas:extension:workspace:2.0:User`:
		*f = UserSchema(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:workspace:2.0:User"`, v)
	}
}

// Type always returns UserSchema to satisfy [pflag.Value] interface
func (f *UserSchema) Type() string {
	return "UserSchema"
}

type WorkspacePermission string

const WorkspacePermissionAdmin WorkspacePermission = `ADMIN`

const WorkspacePermissionUnknown WorkspacePermission = `UNKNOWN`

const WorkspacePermissionUser WorkspacePermission = `USER`

// String representation for [fmt.Print]
func (f *WorkspacePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspacePermission) Set(v string) error {
	switch v {
	case `ADMIN`, `UNKNOWN`, `USER`:
		*f = WorkspacePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADMIN", "UNKNOWN", "USER"`, v)
	}
}

// Type always returns WorkspacePermission to satisfy [pflag.Value] interface
func (f *WorkspacePermission) Type() string {
	return "WorkspacePermission"
}

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	Permissions []PermissionOutput `tfsdk:"permissions"`
}
