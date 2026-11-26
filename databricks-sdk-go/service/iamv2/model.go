// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iamv2

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateDirectGroupMemberProxyRequest struct {
	// The member to add. Set its `principal_id`.
	DirectGroupMember DirectGroupMember `json:"direct_group_member"`
	// The resource name of the group to add the member to, in the format
	// `groups/{group_id}`.
	Parent string `json:"-" url:"-"`
}

func (s *CreateDirectGroupMemberProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateDirectGroupMemberRequest struct {
	// The member to add. Set its `principal_id`.
	DirectGroupMember DirectGroupMember `json:"direct_group_member"`
	// The resource name of the group to add the member to, in the format
	// `accounts/{account_id}/groups/{group_id}`.
	Parent string `json:"-" url:"-"`
}

func (s *CreateDirectGroupMemberRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateGroupProxyRequest struct {
	// The group to create.
	Group Group `json:"group"`
	// The resource name of the account to create the group in, in the format
	// `accounts/{account_id}`. Defaults to the account that contains this
	// workspace. Any other account is rejected.
	Parent string `json:"-" url:"parent,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateGroupProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateGroupRequest struct {
	// The group to create.
	Group Group `json:"group"`
	// The resource name of the account to create the group in, in the format
	// `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`
}

func (s *CreateGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateIdentityVisibilityFilterRequest struct {
	// The filter to create. Set `principal_type` to `GROUP` and exactly one of
	// `exact`, `prefix`, or `regex`. Databricks sets `name`, so any value you
	// set is ignored.
	IdentityVisibilityFilter IdentityVisibilityFilter `json:"identity_visibility_filter"`
	// The ID to give the filter, which becomes the last part of its resource
	// name. If you leave it empty, Databricks generates one. It can't contain
	// `/`.
	IdentityVisibilityFilterId string `json:"-" url:"identity_visibility_filter_id,omitempty"`
	// The resource name of the account, in the format `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateIdentityVisibilityFilterRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateIdentityVisibilityFilterRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateServicePrincipalProxyRequest struct {
	// The resource name of the account to create the service principal in, in
	// the format `accounts/{account_id}`. Defaults to the account that contains
	// this workspace. Any other account is rejected.
	Parent string `json:"-" url:"parent,omitempty"`
	// The service principal to create.
	ServicePrincipal ServicePrincipal `json:"service_principal"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateServicePrincipalProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateServicePrincipalRequest struct {
	// The resource name of the account to create the service principal in, in
	// the format `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`
	// The service principal to create.
	ServicePrincipal ServicePrincipal `json:"service_principal"`
}

func (s *CreateServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateUserProxyRequest struct {
	// The resource name of the account to create the user in, in the format
	// `accounts/{account_id}`. Defaults to the account that contains this
	// workspace. Any other account is rejected.
	Parent string `json:"-" url:"parent,omitempty"`
	// The user to create.
	User User `json:"user"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateUserProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateUserRequest struct {
	// The resource name of the account to create the user in, in the format
	// `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`
	// The user to create.
	User User `json:"user"`
}

func (s *CreateUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateWorkspaceAssignmentDetailProxyRequest struct {
	// Required. Workspace assignment detail to be created in Databricks.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
}

func (s *CreateWorkspaceAssignmentDetailProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateWorkspaceAssignmentDetailRequest struct {
	// Required. Workspace assignment detail to be created in Databricks.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
	// Required. The workspace ID for which the workspace assignment detail is
	// being created.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *CreateWorkspaceAssignmentDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateWorkspaceAssignmentProxyRequest struct {
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`. Defaults to this
	// workspace. Any other workspace is rejected.
	Parent string `json:"-" url:"parent,omitempty"`
	// The assignment to create. Set its `principal_id` and `entitlements`.
	WorkspaceAssignment WorkspaceAssignment `json:"workspace_assignment"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateWorkspaceAssignmentProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWorkspaceAssignmentProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateWorkspaceAssignmentRequest struct {
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`.
	Parent string `json:"-" url:"-"`
	// The assignment to create. Set its `principal_id` and `entitlements`.
	WorkspaceAssignment WorkspaceAssignment `json:"workspace_assignment"`
}

func (s *CreateWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDirectGroupMemberProxyRequest struct {
	// The resource name of the membership, in the format
	// `groups/{group_id}/direct-members/{principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteDirectGroupMemberProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDirectGroupMemberRequest struct {
	// The resource name of the membership, in the format
	// `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteDirectGroupMemberRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteGroupProxyRequest struct {
	// The resource name of the group, in the format `groups/{group_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteGroupRequest struct {
	// The resource name of the group, in the format
	// `accounts/{account_id}/groups/{group_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteIdentityVisibilityFilterRequest struct {
	// The resource name of the filter, in the format
	// `accounts/{account_id}/identity-visibility-filters/{identity_visibility_filter_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteIdentityVisibilityFilterRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteServicePrincipalProxyRequest struct {
	// The resource name of the service principal, in the format
	// `service-principals/{service_principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteServicePrincipalRequest struct {
	// The resource name of the service principal, in the format
	// `accounts/{account_id}/service-principals/{service_principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteUserProxyRequest struct {
	// The resource name of the user, in the format `users/{user_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteUserRequest struct {
	// The resource name of the user, in the format
	// `accounts/{account_id}/users/{user_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWorkspaceAssignmentDetailProxyRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *DeleteWorkspaceAssignmentDetailProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWorkspaceAssignmentDetailRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID where the principal has access.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *DeleteWorkspaceAssignmentDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWorkspaceAssignmentProxyRequest struct {
	// The resource name of the assignment, in the format
	// `workspace-assignments/{principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteWorkspaceAssignmentProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWorkspaceAssignmentRequest struct {
	// The resource name of the assignment, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A user, service principal, or group in your Databricks account that's a
// direct member of a group. To address a membership, put its resource name in
// the URL: `GET
// /api/2.0/identity/accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}`
// in account APIs, or `GET
// /api/2.0/identity/groups/{group_id}/direct-members/{principal_id}` in
// workspace APIs.
type DirectGroupMember struct {
	// The member's display name.
	DisplayName string `json:"display_name,omitempty"`
	// The ID of the member in your identity provider (IdP).
	ExternalId string `json:"external_id,omitempty"`
	// The Databricks ID of the group.
	GroupId int64 `json:"group_id,omitempty"`
	// Where the membership comes from: added in Databricks (in the UI, through
	// the API, or through SCIM provisioning), or synced from your IdP by
	// automatic identity management.
	MembershipSource GroupMembershipSource `json:"membership_source,omitempty"`
	// The resource name of the membership, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}` -
	// Workspace APIs: `groups/{group_id}/direct-members/{principal_id}`
	Name string `json:"name,omitempty"`
	// The resource name of the group, in one of these formats: - Account APIs:
	// `accounts/{account_id}/groups/{group_id}` - Workspace APIs:
	// `groups/{group_id}`
	Parent string `json:"parent,omitempty"`
	// The Databricks ID of the member.
	PrincipalId int64 `json:"principal_id"`
	// Whether the member is a user, service principal, or group.
	PrincipalType PrincipalType `json:"principal_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DirectGroupMember) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DirectGroupMember) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A permission to use part of a Databricks workspace.
type Entitlement string

const EntitlementAllowClusterCreate Entitlement = `ALLOW_CLUSTER_CREATE`

const EntitlementAllowInstancePoolCreate Entitlement = `ALLOW_INSTANCE_POOL_CREATE`

const EntitlementDatabricksSqlAccess Entitlement = `DATABRICKS_SQL_ACCESS`

const EntitlementWorkspaceAccess Entitlement = `WORKSPACE_ACCESS`

const EntitlementWorkspaceAdmin Entitlement = `WORKSPACE_ADMIN`

const EntitlementWorkspaceConsume Entitlement = `WORKSPACE_CONSUME`

// String representation for [fmt.Print]
func (f *Entitlement) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Entitlement) Set(v string) error {
	switch v {
	case `ALLOW_CLUSTER_CREATE`, `ALLOW_INSTANCE_POOL_CREATE`, `DATABRICKS_SQL_ACCESS`, `WORKSPACE_ACCESS`, `WORKSPACE_ADMIN`, `WORKSPACE_CONSUME`:
		*f = Entitlement(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_CLUSTER_CREATE", "ALLOW_INSTANCE_POOL_CREATE", "DATABRICKS_SQL_ACCESS", "WORKSPACE_ACCESS", "WORKSPACE_ADMIN", "WORKSPACE_CONSUME"`, v)
	}
}

// Values returns all possible values for Entitlement.
//
// There is no guarantee on the order of the values in the slice.
func (f *Entitlement) Values() []Entitlement {
	return []Entitlement{
		EntitlementAllowClusterCreate,
		EntitlementAllowInstancePoolCreate,
		EntitlementDatabricksSqlAccess,
		EntitlementWorkspaceAccess,
		EntitlementWorkspaceAdmin,
		EntitlementWorkspaceConsume,
	}
}

// Type always returns Entitlement to satisfy [pflag.Value] interface
func (f *Entitlement) Type() string {
	return "Entitlement"
}

// A rule that matches a group by its ID in your identity provider (IdP).
type ExactMatchRule struct {
	// The ID of the group in your IdP, not its Databricks ID. It's matched
	// exactly, including case.
	ExternalId string `json:"external_id"`
}

func (s *ExactMatchRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A group from your identity provider (IdP), identified by its ID in your IdP.
// This resource is read-only. Getting an external group syncs its name and
// parent groups from your IdP, and adds it to your account if it isn't already
// there. If your account has group visibility filters, the group must match
// them. To address an external group, put its resource name in the URL: `GET
// /api/2.0/identity/accounts/{account_id}/external-groups/{external_group_id}`
// in account APIs, or `GET
// /api/2.0/identity/external-groups/{external_group_id}` in workspace APIs.
type ExternalGroup struct {
	// The ID of the account that contains the group.
	AccountId string `json:"account_id,omitempty"`
	// The name of the group, from your IdP. It's the same as the group's
	// `group_name`.
	DisplayName string `json:"display_name,omitempty"`
	// The ID of the group in your IdP.
	ExternalGroupId string `json:"external_group_id,omitempty"`
	// The Databricks ID of the group. Use it as the `group_id` in the group
	// APIs.
	InternalId string `json:"internal_id,omitempty"`
	// The resource name of the external group, in one of these formats: -
	// Account APIs: `accounts/{account_id}/external-groups/{external_group_id}`
	// - Workspace APIs: `external-groups/{external_group_id}`
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalGroup) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalGroup) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A service principal from your identity provider (IdP), identified by its ID
// in your IdP. This resource is read-only and requires Microsoft Entra ID as
// your IdP. Getting an external service principal syncs its details and group
// memberships from your IdP, and adds it to your account if it isn't already
// there. To address an external service principal, put its resource name in the
// URL: `GET
// /api/2.0/identity/accounts/{account_id}/external-service-principals/{external_service_principal_id}`
// in account APIs, or `GET
// /api/2.0/identity/external-service-principals/{external_service_principal_id}`
// in workspace APIs.
type ExternalServicePrincipal struct {
	// The ID of the account that contains the service principal.
	AccountId string `json:"account_id,omitempty"`
	// Whether the service principal is active in the account.
	AccountSpStatus State `json:"account_sp_status,omitempty"`
	// The service principal's application ID, from your IdP.
	ApplicationId string `json:"application_id,omitempty"`
	// The service principal's display name, from your IdP.
	DisplayName string `json:"display_name,omitempty"`
	// The ID of the service principal in your IdP.
	ExternalServicePrincipalId string `json:"external_service_principal_id,omitempty"`
	// The Databricks ID of the service principal. Use it as the
	// `service_principal_id` in the service principal APIs.
	InternalId string `json:"internal_id,omitempty"`
	// The resource name of the external service principal, in one of these
	// formats: - Account APIs:
	// `accounts/{account_id}/external-service-principals/{external_service_principal_id}`
	// - Workspace APIs:
	// `external-service-principals/{external_service_principal_id}`
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalServicePrincipal) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalServicePrincipal) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A user from your identity provider (IdP), identified by their ID in your IdP.
// This resource is read-only. Getting an external user syncs the user's details
// and group memberships from your IdP, and adds the user to your account if
// they aren't already there. To address an external user, put its resource name
// in the URL: `GET
// /api/2.0/identity/accounts/{account_id}/external-users/{external_user_id}` in
// account APIs, or `GET /api/2.0/identity/external-users/{external_user_id}` in
// workspace APIs.
type ExternalUser struct {
	// The ID of the account that contains the user.
	AccountId string `json:"account_id,omitempty"`
	// Whether the user is active in the account.
	AccountUserStatus State `json:"account_user_status,omitempty"`
	// The user's given and family names from your IdP, separated by a space.
	DisplayName string `json:"display_name,omitempty"`
	// The ID of the user in your IdP.
	ExternalUserId string `json:"external_user_id,omitempty"`
	// The user's name, from your IdP.
	FullName *FullName `json:"full_name,omitempty"`
	// The Databricks ID of the user. Use it as the `user_id` in the user APIs.
	InternalId string `json:"internal_id,omitempty"`
	// The resource name of the external user, in one of these formats: -
	// Account APIs: `accounts/{account_id}/external-users/{external_user_id}` -
	// Workspace APIs: `external-users/{external_user_id}`
	Name string `json:"name,omitempty"`
	// The user's username, which is their email address.
	Username string `json:"username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalUser) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalUser) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A user's name, from your identity provider (IdP).
type FullName struct {
	// The user's family (last) name.
	FamilyName string `json:"family_name,omitempty"`
	// The user's given (first) name.
	GivenName string `json:"given_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FullName) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FullName) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetDirectGroupMemberProxyRequest struct {
	// The resource name of the membership, in the format
	// `groups/{group_id}/direct-members/{principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetDirectGroupMemberProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDirectGroupMemberRequest struct {
	// The resource name of the membership, in the format
	// `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetDirectGroupMemberRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExternalGroupProxyRequest struct {
	// The resource name of the external group, in the format
	// `external-groups/{external_group_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetExternalGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExternalGroupRequest struct {
	// The resource name of the external group, in the format
	// `accounts/{account_id}/external-groups/{external_group_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetExternalGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExternalServicePrincipalProxyRequest struct {
	// The resource name of the external service principal, in the format
	// `external-service-principals/{external_service_principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetExternalServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExternalServicePrincipalRequest struct {
	// The resource name of the external service principal, in the format
	// `accounts/{account_id}/external-service-principals/{external_service_principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetExternalServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExternalUserProxyRequest struct {
	// The resource name of the external user, in the format
	// `external-users/{external_user_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetExternalUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExternalUserRequest struct {
	// The resource name of the external user, in the format
	// `accounts/{account_id}/external-users/{external_user_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetExternalUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetGroupProxyRequest struct {
	// The resource name of the group, in the format `groups/{group_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetGroupRequest struct {
	// The resource name of the group, in the format
	// `accounts/{account_id}/groups/{group_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetIdentityVisibilityFilterRequest struct {
	// The resource name of the filter, in the format
	// `accounts/{account_id}/identity-visibility-filters/{identity_visibility_filter_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetIdentityVisibilityFilterRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetServicePrincipalProxyRequest struct {
	// The resource name of the service principal, in the format
	// `service-principals/{service_principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetServicePrincipalRequest struct {
	// The resource name of the service principal, in the format
	// `accounts/{account_id}/service-principals/{service_principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetUserProxyRequest struct {
	// The resource name of the user, in the format `users/{user_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetUserRequest struct {
	// The resource name of the user, in the format
	// `accounts/{account_id}/users/{user_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAccessDetailLocalRequest struct {
	// The resource name of the workspace access detail. Format:
	// workspace-access-details/{principal_id}
	Name string `json:"-" url:"-"`
	// Controls what fields are returned.
	View WorkspaceAccessDetailView `json:"-" url:"view,omitempty"`
}

func (s *GetWorkspaceAccessDetailLocalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAccessDetailRequest struct {
	// The resource name of the workspace access detail. Format:
	// accounts/{account_id}/workspaces/{workspace_id}/workspace-access-details/{principal_id}
	Name string `json:"-" url:"-"`
	// Controls what fields are returned.
	View WorkspaceAccessDetailView `json:"-" url:"view,omitempty"`
}

func (s *GetWorkspaceAccessDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAssignmentDetailProxyRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment details are being requested.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *GetWorkspaceAssignmentDetailProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAssignmentDetailRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment details are being requested.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The workspace ID for which the assignment details are being
	// requested.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *GetWorkspaceAssignmentDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAssignmentProxyRequest struct {
	// The resource name of the assignment, in the format
	// `workspace-assignments/{principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetWorkspaceAssignmentProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAssignmentRequest struct {
	// The resource name of the assignment, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceIdentityDetailRequest struct {
	// The resource name of the workspace identity detail, in the format
	// `principals/{principal_id}/workspace-identity-detail`.
	Name string `json:"-" url:"-"`
}

func (s *GetWorkspaceIdentityDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A group in your Databricks account. To address a group, put its resource name
// in the URL: `GET /api/2.0/identity/accounts/{account_id}/groups/{group_id}`
// in account APIs, or `GET /api/2.0/identity/groups/{group_id}` in workspace
// APIs.
type Group struct {
	// The ID of the account that contains the group.
	AccountId string `json:"account_id,omitempty"`
	// The ID of the group in your identity provider (IdP).
	ExternalId string `json:"external_id,omitempty"`
	// The Databricks ID of the group. It is a numeric string that Databricks
	// assigns when the group is created, and it is the `{group_id}` segment of
	// `name`.
	GroupId string `json:"group_id,omitempty"`
	// The display name of the group. It must be unique in the account, must not
	// be blank, and must be shorter than 255 characters. Databricks trims
	// surrounding whitespace. You can't use the name of a group that Databricks
	// manages, such as `admins` or `users`.
	GroupName string `json:"group_name,omitempty"`
	// The resource name of the group, in one of these formats: - Account APIs:
	// `accounts/{account_id}/groups/{group_id}` - Workspace APIs:
	// `groups/{group_id}`
	Name string `json:"name,omitempty"`
	// The resource name of the account that contains the group, in the format
	// `accounts/{account_id}`. Workspace APIs return the same format.
	Parent string `json:"parent,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Group) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Group) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Where a group membership comes from.
type GroupMembershipSource string

const GroupMembershipSourceIdentityProvider GroupMembershipSource = `IDENTITY_PROVIDER`

const GroupMembershipSourceInternal GroupMembershipSource = `INTERNAL`

// String representation for [fmt.Print]
func (f *GroupMembershipSource) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GroupMembershipSource) Set(v string) error {
	switch v {
	case `IDENTITY_PROVIDER`, `INTERNAL`:
		*f = GroupMembershipSource(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IDENTITY_PROVIDER", "INTERNAL"`, v)
	}
}

// Values returns all possible values for GroupMembershipSource.
//
// There is no guarantee on the order of the values in the slice.
func (f *GroupMembershipSource) Values() []GroupMembershipSource {
	return []GroupMembershipSource{
		GroupMembershipSourceIdentityProvider,
		GroupMembershipSourceInternal,
	}
}

// Type always returns GroupMembershipSource to satisfy [pflag.Value] interface
func (f *GroupMembershipSource) Type() string {
	return "GroupMembershipSource"
}

// An account-level rule that controls which groups from your identity provider
// (IdP) can be found in directory search and group pickers when automatic
// identity management is on. Filters apply to every workspace in the account.
//
// If you have no group filters, all groups are visible. Otherwise, a group is
// visible if it or any of its parent groups matches at least one filter. A
// group is also visible if it or any of its parent groups is already in your
// account. Local Databricks groups are always visible.
//
// You can only create group filters. Filters don't affect provisioning or
// sign-in for users and service principals. You can't change a filter: to
// change a rule, delete it and create a new one.
//
// To address a filter, put its resource name in the URL: `GET
// /api/2.0/identity/accounts/{account_id}/identity-visibility-filters/{identity_visibility_filter_id}`.
type IdentityVisibilityFilter struct {
	// Matches a group by its ID in your IdP.
	Exact *ExactMatchRule `json:"exact,omitempty"`
	// The resource name of the filter, in the format
	// `accounts/{account_id}/identity-visibility-filters/{identity_visibility_filter_id}`.
	Name string `json:"name,omitempty"`
	// Matches groups whose display name starts with a prefix.
	Prefix *PrefixMatchRule `json:"prefix,omitempty"`
	// The type of principal the filter applies to. Set it to `GROUP`. Other
	// types aren't supported.
	PrincipalType PrincipalType `json:"principal_type"`
	// Matches groups whose display name matches a regular expression.
	Regex *RegexMatchRule `json:"regex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IdentityVisibilityFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IdentityVisibilityFilter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDirectGroupMembersProxyRequest struct {
	// The maximum number of members to return, up to 1000. The default is 1000.
	// A page can contain fewer members.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the group, in the format `groups/{group_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDirectGroupMembersProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDirectGroupMembersProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDirectGroupMembersRequest struct {
	// The maximum number of members to return, up to 1000. The default is 1000.
	// A page can contain fewer members.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the group, in the format
	// `accounts/{account_id}/groups/{group_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDirectGroupMembersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDirectGroupMembersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of direct members of a group.
type ListDirectGroupMembersResponse struct {
	// The members on this page.
	DirectGroupMembers []DirectGroupMember `json:"direct_group_members,omitempty"`
	// The token for the next page. Not set on the last page.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDirectGroupMembersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDirectGroupMembersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListGroupsProxyRequest struct {
	// Filters groups by `group_name` or `external_id`, for example `group_name
	// eq "Data engineers"`. Put the value in quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of groups to return, up to 1000. The default is 1000.
	// A page can contain fewer groups.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the account to list groups in, in the format
	// `accounts/{account_id}`. Defaults to the account that contains this
	// workspace. Any other account is rejected.
	Parent string `json:"-" url:"parent,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListGroupsProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGroupsProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListGroupsRequest struct {
	// Filters groups by `group_name` or `external_id`, for example `group_name
	// eq "Data engineers"`. Put the value in quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of groups to return, up to 1000. The default is 1000.
	// A page can contain fewer groups.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the account to list groups in, in the format
	// `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListGroupsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGroupsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of groups.
type ListGroupsResponse struct {
	// The groups on this page.
	Groups []Group `json:"groups,omitempty"`
	// The token for the next page. Not set on the last page.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListGroupsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGroupsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListIdentityVisibilityFiltersRequest struct {
	// The maximum number of filters to return, up to 1000. The default is 100.
	// A page can contain fewer filters.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the account, in the format `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListIdentityVisibilityFiltersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListIdentityVisibilityFiltersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of group visibility filters.
type ListIdentityVisibilityFiltersResponse struct {
	// The filters on this page. These are the rules, not the groups they match.
	IdentityVisibilityFilters []IdentityVisibilityFilter `json:"identity_visibility_filters,omitempty"`
	// The token for the next page. Not set on the last page.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListIdentityVisibilityFiltersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListIdentityVisibilityFiltersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListServicePrincipalsProxyRequest struct {
	// Filters service principals by `application_id` or `external_id`, for
	// example `application_id eq "<application-id>"`. Put the value in quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of service principals to return, up to 1000. The
	// default is 1000. A page can contain fewer service principals.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the account to list service principals in, in the
	// format `accounts/{account_id}`. Defaults to the account that contains
	// this workspace. Any other account is rejected.
	Parent string `json:"-" url:"parent,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalsProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalsProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListServicePrincipalsRequest struct {
	// Filters service principals by `application_id` or `external_id`, for
	// example `application_id eq "<application-id>"`. Put the value in quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of service principals to return, up to 1000. The
	// default is 1000. A page can contain fewer service principals.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the account to list service principals in, in the
	// format `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of service principals.
type ListServicePrincipalsResponse struct {
	// The token for the next page. Not set on the last page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The service principals on this page.
	ServicePrincipals []ServicePrincipal `json:"service_principals,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTransitiveParentGroupsProxyRequest struct {
	// The maximum number of groups to return, up to 1000. The default is 1000.
	// A page can contain fewer groups.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The user, service principal, or group whose parent groups to list, in the
	// format `principals/{principal_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTransitiveParentGroupsProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTransitiveParentGroupsProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTransitiveParentGroupsRequest struct {
	// The maximum number of groups to return, up to 1000. The default is 1000.
	// A page can contain fewer groups.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The user, service principal, or group whose parent groups to list, in the
	// format `accounts/{account_id}/principals/{principal_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTransitiveParentGroupsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTransitiveParentGroupsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of the groups a principal belongs to.
type ListTransitiveParentGroupsResponse struct {
	// The token for the next page. Not set on the last page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The groups on this page.
	TransitiveParentGroups []TransitiveParentGroup `json:"transitive_parent_groups,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTransitiveParentGroupsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTransitiveParentGroupsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListUsersProxyRequest struct {
	// Filters users by `username` or `external_id`, for example `username eq
	// "someone@example.com"`. Put the value in quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of users to return, up to 1000. The default is 1000. A
	// page can contain fewer users.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the account to list users in, in the format
	// `accounts/{account_id}`. Defaults to the account that contains this
	// workspace. Any other account is rejected.
	Parent string `json:"-" url:"parent,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUsersProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUsersProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListUsersRequest struct {
	// Filters users by `username` or `external_id`, for example `username eq
	// "someone@example.com"`. Put the value in quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of users to return, up to 1000. The default is 1000. A
	// page can contain fewer users.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the account to list users in, in the format
	// `accounts/{account_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUsersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUsersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of users.
type ListUsersResponse struct {
	// The token for the next page. Not set on the last page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The users on this page.
	Users []User `json:"users,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUsersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUsersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceAssignmentDetailsProxyRequest struct {
	// The maximum number of workspace assignment details to return. The service
	// may return fewer than this value. If not provided, defaults to 1000,
	// which is also the maximum allowed. Requests for more than the maximum are
	// clamped to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentDetailsProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentDetailsProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceAssignmentDetailsRequest struct {
	// The maximum number of workspace assignment details to return. The service
	// may return fewer than this value. If not provided, defaults to 1000,
	// which is also the maximum allowed. Requests for more than the maximum are
	// clamped to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListWorkspaceAssignmentDetails
	// call. Provide this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Required. The workspace ID for which the workspace assignment details are
	// being fetched.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentDetailsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentDetailsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message for listing workspace assignment details.
type ListWorkspaceAssignmentDetailsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	WorkspaceAssignmentDetails []WorkspaceAssignmentDetail `json:"workspace_assignment_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentDetailsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentDetailsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceAssignmentsProxyRequest struct {
	// The maximum number of assignments to return, up to 1000. The default is
	// 1000. A page can contain fewer assignments.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`. Defaults to this
	// workspace. Any other workspace is rejected.
	Parent string `json:"-" url:"parent,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentsProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentsProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceAssignmentsRequest struct {
	// The maximum number of assignments to return, up to 1000. The default is
	// 1000. A page can contain fewer assignments.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The `next_page_token` from a previous response, to get the next page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of workspace assignments.
type ListWorkspaceAssignmentsResponse struct {
	// The token for the next page. Not set on the last page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The assignments on this page.
	WorkspaceAssignments []WorkspaceAssignment `json:"workspace_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A rule that matches groups whose display name starts with a prefix. Matching
// is case-sensitive. For example, the prefix `data-` matches `data-platform`.
type PrefixMatchRule struct {
	// The prefix to match. It must contain at least one non-whitespace
	// character. Leading and trailing spaces are kept.
	DisplayNamePrefix string `json:"display_name_prefix"`
}

func (s *PrefixMatchRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The type of a principal.
type PrincipalType string

const PrincipalTypeGroup PrincipalType = `GROUP`

const PrincipalTypeServicePrincipal PrincipalType = `SERVICE_PRINCIPAL`

const PrincipalTypeUser PrincipalType = `USER`

// String representation for [fmt.Print]
func (f *PrincipalType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrincipalType) Set(v string) error {
	switch v {
	case `GROUP`, `SERVICE_PRINCIPAL`, `USER`:
		*f = PrincipalType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GROUP", "SERVICE_PRINCIPAL", "USER"`, v)
	}
}

// Values returns all possible values for PrincipalType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrincipalType) Values() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeGroup,
		PrincipalTypeServicePrincipal,
		PrincipalTypeUser,
	}
}

// Type always returns PrincipalType to satisfy [pflag.Value] interface
func (f *PrincipalType) Type() string {
	return "PrincipalType"
}

// A rule that matches groups whose whole display name matches an RE2 regular
// expression. Matching is case-sensitive. To match part of a name, add `.*`,
// for example `.*-admins`.
type RegexMatchRule struct {
	// The RE2 regular expression to match against the whole display name. At
	// most 255 characters.
	DisplayNameRegex string `json:"display_name_regex"`
}

func (s *RegexMatchRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to resolve a group by its ID in your identity provider (IdP).
type ResolveGroupProxyRequest struct {
	// The ID of the group in your IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to resolve a group by its ID in your identity provider (IdP).
type ResolveGroupRequest struct {
	// The ID of the group in your IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The response to resolving a group.
type ResolveGroupResponse struct {
	// The resolved group.
	Group *Group `json:"group,omitempty"`
}

func (s *ResolveGroupResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to resolve a service principal by its ID in your identity provider
// (IdP).
type ResolveServicePrincipalProxyRequest struct {
	// The ID of the service principal in your IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to resolve a service principal by its ID in your identity provider
// (IdP).
type ResolveServicePrincipalRequest struct {
	// The ID of the service principal in your IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The response to resolving a service principal.
type ResolveServicePrincipalResponse struct {
	// The resolved service principal.
	ServicePrincipal *ServicePrincipal `json:"service_principal,omitempty"`
}

func (s *ResolveServicePrincipalResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to resolve a user by their ID in your identity provider (IdP).
type ResolveUserProxyRequest struct {
	// The ID of the user in your IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to resolve a user by their ID in your identity provider (IdP).
type ResolveUserRequest struct {
	// The ID of the user in your IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The response to resolving a user.
type ResolveUserResponse struct {
	// The resolved user.
	User *User `json:"user,omitempty"`
}

func (s *ResolveUserResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A service principal in your Databricks account. To address a service
// principal, put its resource name in the URL: `GET
// /api/2.0/identity/accounts/{account_id}/service-principals/{service_principal_id}`
// in account APIs, or `GET
// /api/2.0/identity/service-principals/{service_principal_id}` in workspace
// APIs.
type ServicePrincipal struct {
	// The ID of the account that contains the service principal.
	AccountId string `json:"account_id,omitempty"`
	// Whether the service principal is active in the account.
	AccountSpStatus State `json:"account_sp_status"`
	// The service principal's application ID, a UUID. If you don't set it when
	// you create the service principal, Databricks generates one. It must be
	// unique in the account and can't be changed.
	ApplicationId string `json:"application_id,omitempty"`
	// The service principal's display name.
	DisplayName string `json:"display_name"`
	// The ID of the service principal in your identity provider (IdP).
	ExternalId string `json:"external_id,omitempty"`
	// The resource name of the service principal, in one of these formats: -
	// Account APIs:
	// `accounts/{account_id}/service-principals/{service_principal_id}` -
	// Workspace APIs: `service-principals/{service_principal_id}`
	Name string `json:"name,omitempty"`
	// The resource name of the account that contains the service principal, in
	// the format `accounts/{account_id}`. Workspace APIs return the same
	// format.
	Parent string `json:"parent,omitempty"`
	// The Databricks ID of the service principal.
	ServicePrincipalId string `json:"service_principal_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServicePrincipal) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServicePrincipal) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Whether a user or service principal is active in an account or workspace.
type State string

const StateActive State = `ACTIVE`

const StateInactive State = `INACTIVE`

// String representation for [fmt.Print]
func (f *State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *State) Set(v string) error {
	switch v {
	case `ACTIVE`, `INACTIVE`:
		*f = State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "INACTIVE"`, v)
	}
}

// Values returns all possible values for State.
//
// There is no guarantee on the order of the values in the slice.
func (f *State) Values() []State {
	return []State{
		StateActive,
		StateInactive,
	}
}

// Type always returns State to satisfy [pflag.Value] interface
func (f *State) Type() string {
	return "State"
}

// A group in your Databricks account that a principal belongs to, either
// directly or through another group. To list a principal's parent groups, put
// the principal's resource name in the URL: `GET
// /api/2.0/identity/accounts/{account_id}/principals/{principal_id}/transitive-parent-groups`
// in account APIs, or `GET
// /api/2.0/identity/principals/{principal_id}/transitive-parent-groups` in
// workspace APIs.
type TransitiveParentGroup struct {
	// The ID of the account that contains the group.
	AccountId string `json:"account_id,omitempty"`
	// The ID of the group in your identity provider (IdP).
	ExternalId string `json:"external_id,omitempty"`
	// The Databricks ID of the group.
	GroupId string `json:"group_id,omitempty"`
	// The resource name of the parent group, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/principals/{principal_id}/transitive-parent-groups/{group_id}`
	// - Workspace APIs:
	// `principals/{principal_id}/transitive-parent-groups/{group_id}`
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TransitiveParentGroup) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TransitiveParentGroup) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateGroupProxyRequest struct {
	// The group to update. Its `name` identifies the group.
	Group Group `json:"group"`
	// The resource name of the group, in one of these formats: - Account APIs:
	// `accounts/{account_id}/groups/{group_id}` - Workspace APIs:
	// `groups/{group_id}`
	Name string `json:"-" url:"-"`
	// The fields to update, separated by commas. You can update `group_name`
	// and `external_id`.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateGroupRequest struct {
	// The group to update. Its `name` identifies the group.
	Group Group `json:"group"`
	// The resource name of the group, in one of these formats: - Account APIs:
	// `accounts/{account_id}/groups/{group_id}` - Workspace APIs:
	// `groups/{group_id}`
	Name string `json:"-" url:"-"`
	// The fields to update, separated by commas. You can update `group_name`
	// and `external_id`.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateServicePrincipalProxyRequest struct {
	// The resource name of the service principal, in one of these formats: -
	// Account APIs:
	// `accounts/{account_id}/service-principals/{service_principal_id}` -
	// Workspace APIs: `service-principals/{service_principal_id}`
	Name string `json:"-" url:"-"`
	// The service principal to update. Its `name` identifies the service
	// principal.
	ServicePrincipal ServicePrincipal `json:"service_principal"`
	// The fields to update, separated by commas. You can update `display_name`,
	// `account_sp_status`, and `external_id`.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateServicePrincipalRequest struct {
	// The resource name of the service principal, in one of these formats: -
	// Account APIs:
	// `accounts/{account_id}/service-principals/{service_principal_id}` -
	// Workspace APIs: `service-principals/{service_principal_id}`
	Name string `json:"-" url:"-"`
	// The service principal to update. Its `name` identifies the service
	// principal.
	ServicePrincipal ServicePrincipal `json:"service_principal"`
	// The fields to update, separated by commas. You can update `display_name`,
	// `account_sp_status`, and `external_id`.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateUserProxyRequest struct {
	// The resource name of the user, in one of these formats: - Account APIs:
	// `accounts/{account_id}/users/{user_id}` - Workspace APIs:
	// `users/{user_id}`
	Name string `json:"-" url:"-"`
	// The fields to update, separated by commas. You can update
	// `full_name.given_name`, `full_name.family_name`, `account_user_status`,
	// and `external_id`.
	UpdateMask string `json:"-" url:"update_mask"`
	// The user to update. Its `name` identifies the user.
	User User `json:"user"`
}

func (s *UpdateUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateUserRequest struct {
	// The resource name of the user, in one of these formats: - Account APIs:
	// `accounts/{account_id}/users/{user_id}` - Workspace APIs:
	// `users/{user_id}`
	Name string `json:"-" url:"-"`
	// The fields to update, separated by commas. You can update
	// `full_name.given_name`, `full_name.family_name`, `account_user_status`,
	// and `external_id`.
	UpdateMask string `json:"-" url:"update_mask"`
	// The user to update. Its `name` identifies the user.
	User User `json:"user"`
}

func (s *UpdateUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceAssignmentDetailProxyRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// Required. Workspace assignment detail to be updated in Databricks.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
}

func (s *UpdateWorkspaceAssignmentDetailProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceAssignmentDetailRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// Required. Workspace assignment detail to be updated in Databricks.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
	// Required. The workspace ID for which the workspace assignment detail is
	// being updated.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *UpdateWorkspaceAssignmentDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceAssignmentProxyRequest struct {
	// The resource name of the assignment, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
	// - Workspace APIs: `workspace-assignments/{principal_id}`
	Name string `json:"-" url:"-"`
	// The fields to update. The only supported field is `entitlements`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// The assignment to update. Its `name` identifies the assignment.
	WorkspaceAssignment WorkspaceAssignment `json:"workspace_assignment"`
}

func (s *UpdateWorkspaceAssignmentProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceAssignmentRequest struct {
	// The resource name of the assignment, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
	// - Workspace APIs: `workspace-assignments/{principal_id}`
	Name string `json:"-" url:"-"`
	// The fields to update. The only supported field is `entitlements`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// The assignment to update. Its `name` identifies the assignment.
	WorkspaceAssignment WorkspaceAssignment `json:"workspace_assignment"`
}

func (s *UpdateWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceIdentityDetailRequest struct {
	// The resource name of the workspace identity detail, in the format
	// `principals/{principal_id}/workspace-identity-detail`.
	Name string `json:"-" url:"-"`
	// The fields to update. The only supported field is
	// `workspace_identity_status`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// The workspace identity detail to update. Its `name` identifies the
	// principal.
	WorkspaceIdentityDetail WorkspaceIdentityDetail `json:"workspace_identity_detail"`
}

func (s *UpdateWorkspaceIdentityDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A user in your Databricks account. To address a user, put its resource name
// in the URL: `GET /api/2.0/identity/accounts/{account_id}/users/{user_id}` in
// account APIs, or `GET /api/2.0/identity/users/{user_id}` in workspace APIs.
type User struct {
	// The ID of the account that contains the user.
	AccountId string `json:"account_id,omitempty"`
	// Whether the user is active in the account. New users are active unless
	// you set this to `INACTIVE`.
	AccountUserStatus State `json:"account_user_status"`
	// The ID of the user in your identity provider (IdP).
	ExternalId string `json:"external_id,omitempty"`
	// The user's name.
	FullName UserFullName `json:"full_name"`
	// The resource name of the user, in one of these formats: - Account APIs:
	// `accounts/{account_id}/users/{user_id}` - Workspace APIs:
	// `users/{user_id}`
	Name string `json:"name,omitempty"`
	// The resource name of the account that contains the user, in the format
	// `accounts/{account_id}`. Workspace APIs return the same format.
	Parent string `json:"parent,omitempty"`
	// The Databricks ID of the user.
	UserId string `json:"user_id,omitempty"`
	// The user's username, which is their email address. It can't be changed
	// after the user is created.
	Username string `json:"username"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *User) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s User) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A user's name.
type UserFullName struct {
	// The user's family (last) name.
	FamilyName string `json:"family_name,omitempty"`
	// The user's given (first) name.
	GivenName string `json:"given_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UserFullName) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UserFullName) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The details of a principal's access to a workspace.
type WorkspaceAccessDetail struct {
	AccessType WorkspaceAccessDetailAccessType `json:"access_type,omitempty"`
	// The account ID parent of the workspace where the principal has access.
	AccountId string `json:"account_id,omitempty"`
	// The resource name of the workspace access detail. The format depends on
	// the API that returned it: - Account-scoped:
	// accounts/{account_id}/workspaces/{workspace_id}/workspace-access-details/{principal_id}
	// - Workspace-scoped: workspace-access-details/{principal_id}
	Name string `json:"name,omitempty"`
	// The permissions granted to the principal in the workspace.
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId int64 `json:"principal_id,omitempty"`

	PrincipalType PrincipalType `json:"principal_type,omitempty"`
	// The activity status of the principal in the workspace. Not applicable for
	// groups at the moment.
	Status State `json:"status,omitempty"`
	// The workspace ID where the principal has access.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceAccessDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceAccessDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of access the principal has to the workspace.
type WorkspaceAccessDetailAccessType string

const WorkspaceAccessDetailAccessTypeDirect WorkspaceAccessDetailAccessType = `DIRECT`

const WorkspaceAccessDetailAccessTypeIndirect WorkspaceAccessDetailAccessType = `INDIRECT`

// String representation for [fmt.Print]
func (f *WorkspaceAccessDetailAccessType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceAccessDetailAccessType) Set(v string) error {
	switch v {
	case `DIRECT`, `INDIRECT`:
		*f = WorkspaceAccessDetailAccessType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECT", "INDIRECT"`, v)
	}
}

// Values returns all possible values for WorkspaceAccessDetailAccessType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceAccessDetailAccessType) Values() []WorkspaceAccessDetailAccessType {
	return []WorkspaceAccessDetailAccessType{
		WorkspaceAccessDetailAccessTypeDirect,
		WorkspaceAccessDetailAccessTypeIndirect,
	}
}

// Type always returns WorkspaceAccessDetailAccessType to satisfy [pflag.Value] interface
func (f *WorkspaceAccessDetailAccessType) Type() string {
	return "WorkspaceAccessDetailAccessType"
}

// Controls what fields are returned in the GetWorkspaceAccessDetail response.
type WorkspaceAccessDetailView string

const WorkspaceAccessDetailViewBasic WorkspaceAccessDetailView = `BASIC`

const WorkspaceAccessDetailViewFull WorkspaceAccessDetailView = `FULL`

// String representation for [fmt.Print]
func (f *WorkspaceAccessDetailView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceAccessDetailView) Set(v string) error {
	switch v {
	case `BASIC`, `FULL`:
		*f = WorkspaceAccessDetailView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BASIC", "FULL"`, v)
	}
}

// Values returns all possible values for WorkspaceAccessDetailView.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceAccessDetailView) Values() []WorkspaceAccessDetailView {
	return []WorkspaceAccessDetailView{
		WorkspaceAccessDetailViewBasic,
		WorkspaceAccessDetailViewFull,
	}
}

// Type always returns WorkspaceAccessDetailView to satisfy [pflag.Value] interface
func (f *WorkspaceAccessDetailView) Type() string {
	return "WorkspaceAccessDetailView"
}

// A principal's direct assignment to a workspace, and the entitlements it
// grants there.
//
// Only principals assigned directly to the workspace have an assignment. A
// principal that gets access only through a group doesn't. Creating an
// assignment gives the principal access to the workspace, and deleting it
// removes that direct access. These APIs replace managing workspace assignments
// through SCIM and the permission assignment APIs. They require a workspace
// with identity federation enabled. You can't assign, change, or unassign
// principals that Databricks manages.
//
// `entitlements` is the only field you set. `effective_entitlements` also
// includes the entitlements the principal gets through its groups. An
// assignment must grant at least one entitlement. To remove a principal's
// direct access, delete the assignment.
//
// To address an assignment, put its resource name in the URL: `GET
// /api/2.0/identity/accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
// in account APIs, or `GET
// /api/2.0/identity/workspace-assignments/{principal_id}` in workspace APIs.
type WorkspaceAssignment struct {
	// The ID of the account that contains the workspace.
	AccountId string `json:"account_id,omitempty"`
	// Every entitlement the principal has in the workspace, granted directly or
	// through a group. List responses leave it empty.
	EffectiveEntitlements []Entitlement `json:"effective_entitlements,omitempty"`
	// The entitlements granted directly by this assignment. Create and update
	// set exactly this list, even if the principal also gets some of these
	// entitlements through a group. List responses leave it empty.
	Entitlements []Entitlement `json:"entitlements,omitempty"`
	// The resource name of the assignment, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
	// - Workspace APIs: `workspace-assignments/{principal_id}`
	Name string `json:"name,omitempty"`
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`. Workspace APIs return
	// the same format.
	Parent string `json:"parent,omitempty"`
	// The Databricks ID of the assigned user, service principal, or group. It
	// can't be changed after the assignment is created.
	PrincipalId int64 `json:"principal_id"`
	// Whether the assigned principal is a user, service principal, or group.
	PrincipalType PrincipalType `json:"principal_type,omitempty"`
	// The ID of the workspace.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The direct assignment of a provisioned account-level principal (user, service
// principal, or group) to a workspace, together with the entitlements that
// assignment grants in the workspace.
//
// This resource covers only principals assigned directly to the workspace.
// Principals that inherit workspace access through a group are not represented
// here. See WorkspaceAccessDetail and WorkspaceIdentityDetail for the
// effective, direct-or-indirect view. Creating the resource assigns the
// principal to the workspace, and deleting it removes the assignment.
//
// `entitlements` is the only client-settable field. It holds the entitlements
// granted directly on this assignment, including any the principal also holds
// through a group. `effective_entitlements` is the read-only union of those and
// any granted through group membership.
//
// A direct assignment always carries at least one directly-assigned
// entitlement, because the assignment is what grants it. Create and update both
// reject an empty `entitlements` set. To remove a principal's assignment
// entirely, delete the resource.
//
// This resource replaces workspace assignment previously managed through the
// workspace SCIM and permission-assignment APIs, and is intended for account
// and workspace admins.
type WorkspaceAssignmentDetail struct {
	// The account ID parent of the workspace where the principal is assigned
	AccountId string `json:"account_id,omitempty"`
	// Every entitlement the principal holds in this workspace, whether granted
	// directly or through group membership. Get responses populate this field.
	// List responses leave it empty.
	EffectiveEntitlements []Entitlement `json:"effective_entitlements,omitempty"`
	// Entitlements granted directly to the principal on this workspace. This is
	// the only client-settable field. Create and update manage exactly this
	// set, including entitlements the principal also holds through a group.
	// List responses leave this field empty. Get a single principal to read its
	// entitlements.
	Entitlements []Entitlement `json:"entitlements,omitempty"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId int64 `json:"principal_id"`

	PrincipalType PrincipalType `json:"principal_type,omitempty"`
	// The workspace ID where the principal is assigned.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceAssignmentDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceAssignmentDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A user, service principal, or group's status in a workspace, and whether it
// has access to the workspace directly or through a group. Update it to
// activate or deactivate a user or service principal in the workspace. To
// address it, put its resource name in the URL: `GET
// /api/2.0/identity/principals/{principal_id}/workspace-identity-detail`.
type WorkspaceIdentityDetail struct {
	// Whether the principal is assigned to the workspace directly or has access
	// through a group.
	AssignmentType WorkspaceIdentityDetailAssignmentType `json:"assignment_type,omitempty"`
	// The resource name of the workspace identity detail, in the format
	// `principals/{principal_id}/workspace-identity-detail`.
	Name string `json:"name,omitempty"`
	// The Databricks ID of the user, service principal, or group.
	PrincipalId int64 `json:"principal_id,omitempty"`
	// Whether the principal is a user, service principal, or group.
	PrincipalType PrincipalType `json:"principal_type,omitempty"`
	// Whether the principal is active in the workspace. Groups are always
	// `ACTIVE`.
	WorkspaceIdentityStatus State `json:"workspace_identity_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceIdentityDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceIdentityDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// How a principal has access to a workspace.
type WorkspaceIdentityDetailAssignmentType string

const WorkspaceIdentityDetailAssignmentTypeDirect WorkspaceIdentityDetailAssignmentType = `DIRECT`

const WorkspaceIdentityDetailAssignmentTypeIndirect WorkspaceIdentityDetailAssignmentType = `INDIRECT`

// String representation for [fmt.Print]
func (f *WorkspaceIdentityDetailAssignmentType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceIdentityDetailAssignmentType) Set(v string) error {
	switch v {
	case `DIRECT`, `INDIRECT`:
		*f = WorkspaceIdentityDetailAssignmentType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECT", "INDIRECT"`, v)
	}
}

// Values returns all possible values for WorkspaceIdentityDetailAssignmentType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceIdentityDetailAssignmentType) Values() []WorkspaceIdentityDetailAssignmentType {
	return []WorkspaceIdentityDetailAssignmentType{
		WorkspaceIdentityDetailAssignmentTypeDirect,
		WorkspaceIdentityDetailAssignmentTypeIndirect,
	}
}

// Type always returns WorkspaceIdentityDetailAssignmentType to satisfy [pflag.Value] interface
func (f *WorkspaceIdentityDetailAssignmentType) Type() string {
	return "WorkspaceIdentityDetailAssignmentType"
}

// The type of permission a principal has to a workspace (admin/user).
type WorkspacePermission string

const WorkspacePermissionAdminPermission WorkspacePermission = `ADMIN_PERMISSION`

const WorkspacePermissionUserPermission WorkspacePermission = `USER_PERMISSION`

// String representation for [fmt.Print]
func (f *WorkspacePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspacePermission) Set(v string) error {
	switch v {
	case `ADMIN_PERMISSION`, `USER_PERMISSION`:
		*f = WorkspacePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADMIN_PERMISSION", "USER_PERMISSION"`, v)
	}
}

// Values returns all possible values for WorkspacePermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspacePermission) Values() []WorkspacePermission {
	return []WorkspacePermission{
		WorkspacePermissionAdminPermission,
		WorkspacePermissionUserPermission,
	}
}

// Type always returns WorkspacePermission to satisfy [pflag.Value] interface
func (f *WorkspacePermission) Type() string {
	return "WorkspacePermission"
}
