// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iamv2

import (
	"context"
)

// These APIs are used to manage identities and the workspace access of these
// identities in Databricks.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountIamV2Service interface {

	// Adds a user, service principal, or group as a direct member of a group.
	//
	// If automatic identity management is on and the group has an
	// `external_id`, its members come from your identity provider and can't be
	// changed in Databricks.
	CreateDirectGroupMember(ctx context.Context, request CreateDirectGroupMemberRequest) (*DirectGroupMember, error)

	// Creates a local group, which is a group that isn't synced from your
	// identity provider (IdP). You become a manager of the new group.
	//
	// If automatic identity management is on, you can't set `external_id`. To
	// add a group from your IdP, get it as an external group instead.
	CreateGroup(ctx context.Context, request CreateGroupRequest) (*Group, error)

	// Creates a group visibility filter. It applies to every workspace in the
	// account.
	//
	// Returns `ALREADY_EXISTS` if a filter with the same ID, or the same type
	// and match value, already exists. Returns `RESOURCE_EXHAUSTED` if the
	// account already has the maximum number of filters of that type. You can't
	// change a filter: to change a rule, delete it and create a new one.
	CreateIdentityVisibilityFilter(ctx context.Context, request CreateIdentityVisibilityFilterRequest) (*IdentityVisibilityFilter, error)

	// Creates a local service principal, which is a service principal that
	// isn't synced from your identity provider (IdP).
	//
	// If automatic identity management is on, you can't set `external_id`. To
	// add a service principal from your IdP, get it as an external service
	// principal instead.
	CreateServicePrincipal(ctx context.Context, request CreateServicePrincipalRequest) (*ServicePrincipal, error)

	// Creates a local user, which is a user that isn't synced from your
	// identity provider (IdP).
	//
	// If automatic identity management is on, you can't set `external_id`. To
	// add a user from your IdP, get them as an external user instead.
	CreateUser(ctx context.Context, request CreateUserRequest) (*User, error)

	// Assigns a user, service principal, or group directly to a workspace with
	// the given entitlements. Returns an error if the principal is already
	// assigned directly.
	//
	// Entitlements are granted one at a time. If the request fails partway
	// through, the principal stays assigned with only some of the entitlements.
	// Get the assignment afterwards to check which were granted.
	CreateWorkspaceAssignment(ctx context.Context, request CreateWorkspaceAssignmentRequest) (*WorkspaceAssignment, error)

	// Creates a workspace assignment detail for a principal.
	CreateWorkspaceAssignmentDetail(ctx context.Context, request CreateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)

	// Removes a direct member from a group.
	//
	// If automatic identity management is on and the group has an
	// `external_id`, its members come from your identity provider and can't be
	// changed in Databricks.
	DeleteDirectGroupMember(ctx context.Context, request DeleteDirectGroupMemberRequest) error

	// Deletes a group from your account and removes it from the workspaces it's
	// assigned to.
	DeleteGroup(ctx context.Context, request DeleteGroupRequest) error

	// Deletes a group visibility filter. Deleting a filter doesn't delete
	// groups or remove access. Groups that only this filter matched stop being
	// visible, unless they or a parent group are already in your account. If no
	// group filters remain, all groups become visible again. Deleting a filter
	// that doesn't exist succeeds.
	DeleteIdentityVisibilityFilter(ctx context.Context, request DeleteIdentityVisibilityFilterRequest) error

	// Deletes a service principal from your account and removes it from the
	// workspaces it's assigned to. You can't delete service principals that
	// Databricks manages.
	DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error

	// Deletes a user from your account and removes them from the workspaces
	// they're assigned to.
	DeleteUser(ctx context.Context, request DeleteUserRequest) error

	// Removes a principal's direct assignment to a workspace and revokes the
	// entitlements it granted. The principal keeps any access it gets through
	// groups.
	//
	// Entitlements are revoked one at a time. If the request fails partway
	// through, the principal stays assigned with some of its entitlements. You
	// can retry the request.
	DeleteWorkspaceAssignment(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error

	// Deletes a workspace assignment detail for a principal, revoking all of
	// its entitlements. Entitlements are revoked one at a time rather than
	// atomically. If the request fails partway through, the principal stays
	// assigned with some of its original entitlements. Retrying is safe.
	DeleteWorkspaceAssignmentDetail(ctx context.Context, request DeleteWorkspaceAssignmentDetailRequest) error

	// Gets a direct member of a group, including where the membership comes
	// from. Returns an error if the principal isn't a direct member.
	GetDirectGroupMember(ctx context.Context, request GetDirectGroupMemberRequest) (*DirectGroupMember, error)

	// Gets a group by its ID in your identity provider (IdP) and syncs its name
	// and parent groups from your IdP. If the group isn't in your account yet,
	// it's added. Requires automatic identity management. If your account has
	// group visibility filters, the group must match them.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	GetExternalGroup(ctx context.Context, request GetExternalGroupRequest) (*ExternalGroup, error)

	// Gets a service principal by its ID in your identity provider (IdP) and
	// syncs its details and group memberships from your IdP. If the service
	// principal isn't in your account yet, it's added. Requires automatic
	// identity management with Microsoft Entra ID as your IdP.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	GetExternalServicePrincipal(ctx context.Context, request GetExternalServicePrincipalRequest) (*ExternalServicePrincipal, error)

	// Gets a user by their ID in your identity provider (IdP) and syncs their
	// details and group memberships from your IdP. If the user isn't in your
	// account yet, they're added. Requires automatic identity management.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	GetExternalUser(ctx context.Context, request GetExternalUserRequest) (*ExternalUser, error)

	// Gets a group by its Databricks ID.
	GetGroup(ctx context.Context, request GetGroupRequest) (*Group, error)

	// Gets a group visibility filter. It returns the rule, not the groups it
	// matches.
	GetIdentityVisibilityFilter(ctx context.Context, request GetIdentityVisibilityFilterRequest) (*IdentityVisibilityFilter, error)

	// Gets a service principal by its Databricks ID.
	GetServicePrincipal(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error)

	// Gets a user by their Databricks ID.
	GetUser(ctx context.Context, request GetUserRequest) (*User, error)

	// Returns the access details for a principal in a workspace. Allows for
	// checking access details for any provisioned principal (user, service
	// principal, or group) in a workspace. * Provisioned principal here refers
	// to one that has been synced into Databricks from the customer's IdP or
	// added explicitly to Databricks via SCIM/UI. Allows for passing in a
	// "view" parameter to control what fields are returned (BASIC by default or
	// FULL).
	GetWorkspaceAccessDetail(ctx context.Context, request GetWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error)

	// Gets a principal's direct assignment to a workspace, including its
	// effective entitlements. Returns an error if the principal isn't assigned
	// directly.
	GetWorkspaceAssignment(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspaceAssignment, error)

	// Returns the assignment details for a principal in a workspace.
	GetWorkspaceAssignmentDetail(ctx context.Context, request GetWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)

	// Lists the direct members of a group, sorted by display name, with where
	// each membership comes from. Results combine memberships added in
	// Databricks (in the UI, through the API, or through SCIM provisioning)
	// with, if automatic identity management is on and the group has an
	// `external_id`, memberships from your identity provider (IdP). Only
	// members in your Databricks account are returned: an IdP member of the
	// group appears only if that user, service principal, or group has been
	// added to Databricks. Nested groups aren't expanded. To get the full
	// membership, list the direct members of each member group.
	ListDirectGroupMembers(ctx context.Context, request ListDirectGroupMembersRequest) (*ListDirectGroupMembersResponse, error)

	// Lists the groups in your account, sorted by Databricks ID. Groups in your
	// identity provider that haven't been added to Databricks aren't included.
	ListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error)

	// Lists the group visibility filters in an account. It returns the rules,
	// not the groups they match.
	ListIdentityVisibilityFilters(ctx context.Context, request ListIdentityVisibilityFiltersRequest) (*ListIdentityVisibilityFiltersResponse, error)

	// Lists the service principals in your account, sorted by Databricks ID.
	// Service principals in your identity provider that haven't been added to
	// Databricks aren't included.
	ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalsResponse, error)

	// Lists every group that a user, service principal, or group belongs to,
	// directly or through other groups, sorted by Databricks ID. Results
	// combine memberships from every source: added in the UI, through the API,
	// through SCIM provisioning, or synced from your identity provider (IdP) by
	// automatic identity management. Only groups that exist in your Databricks
	// account are returned, not every group that the principal belongs to in
	// your IdP. If automatic identity management is on, only the principal's
	// own group memberships are refreshed from your IdP before the list is
	// read.
	ListTransitiveParentGroups(ctx context.Context, request ListTransitiveParentGroupsRequest) (*ListTransitiveParentGroupsResponse, error)

	// Lists the users in your account, sorted by Databricks ID. Users in your
	// identity provider who haven't been added to Databricks aren't included.
	ListUsers(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error)

	// Lists workspace assignment details for a workspace. The response omits
	// the per-principal entitlement fields (`entitlements` and
	// `effective_entitlements`). To read the entitlements for a single
	// principal, get that principal's assignment detail.
	ListWorkspaceAssignmentDetails(ctx context.Context, request ListWorkspaceAssignmentDetailsRequest) (*ListWorkspaceAssignmentDetailsResponse, error)

	// Lists the principals assigned directly to a workspace. The response
	// leaves `entitlements` and `effective_entitlements` empty. To read a
	// principal's entitlements, get its assignment.
	ListWorkspaceAssignments(ctx context.Context, request ListWorkspaceAssignmentsRequest) (*ListWorkspaceAssignmentsResponse, error)

	// Gets a group by its ID in your identity provider (IdP) and syncs its name
	// and parent groups from your IdP. If the group isn't in your account yet,
	// it's added. Requires automatic identity management. If your account has
	// group visibility filters, the group must match them.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	ResolveGroup(ctx context.Context, request ResolveGroupRequest) (*ResolveGroupResponse, error)

	// Gets a service principal by its ID in your identity provider (IdP) and
	// syncs its details and group memberships from your IdP. If the service
	// principal isn't in your account yet, it's added. Requires automatic
	// identity management with Microsoft Entra ID as your IdP.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	ResolveServicePrincipal(ctx context.Context, request ResolveServicePrincipalRequest) (*ResolveServicePrincipalResponse, error)

	// Gets a user by their ID in your identity provider (IdP) and syncs their
	// details and group memberships from your IdP. If the user isn't in your
	// account yet, they're added. Requires automatic identity management.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	ResolveUser(ctx context.Context, request ResolveUserRequest) (*ResolveUserResponse, error)

	// Updates a group. Only the fields in `update_mask` change.
	//
	// If automatic identity management is on and the group has an
	// `external_id`, you can only update `external_id`. The group's other
	// fields come from your identity provider. If your account also has group
	// visibility filters, a new `external_id` must match them. You can always
	// clear `external_id`.
	UpdateGroup(ctx context.Context, request UpdateGroupRequest) (*Group, error)

	// Updates a service principal. Only the fields in `update_mask` change.
	//
	// If automatic identity management is on and the service principal has an
	// `external_id`, you can only update `external_id`. Its other fields come
	// from your identity provider. A service principal can't deactivate itself,
	// and you can't update service principals that Databricks manages.
	UpdateServicePrincipal(ctx context.Context, request UpdateServicePrincipalRequest) (*ServicePrincipal, error)

	// Updates a user. Only the fields in `update_mask` change.
	//
	// If automatic identity management is on and the user has an `external_id`,
	// you can only update `external_id`. The user's other fields come from your
	// identity provider. You can't deactivate yourself or change the status of
	// the account owner.
	UpdateUser(ctx context.Context, request UpdateUserRequest) (*User, error)

	// Replaces the entitlements of a principal's direct assignment to a
	// workspace. The new list can't be empty.
	//
	// Entitlements are changed one at a time. If the request fails partway
	// through, only some of the changes take effect. Get the assignment
	// afterwards to check the result.
	UpdateWorkspaceAssignment(ctx context.Context, request UpdateWorkspaceAssignmentRequest) (*WorkspaceAssignment, error)

	// Updates the entitlements of a directly assigned principal in a workspace.
	// Changes are applied one at a time rather than atomically. If the request
	// fails partway through, only some of the requested changes take effect.
	// Get the assignment detail afterwards to confirm the final state.
	UpdateWorkspaceAssignmentDetail(ctx context.Context, request UpdateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error)
}

// These APIs are used to manage identities and the workspace access of these
// identities in Databricks.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceIamV2Service interface {

	// Adds a user, service principal, or group as a direct member of a group in
	// the account that contains this workspace.
	//
	// If automatic identity management is on and the group has an
	// `external_id`, its members come from your identity provider and can't be
	// changed in Databricks.
	CreateDirectGroupMemberProxy(ctx context.Context, request CreateDirectGroupMemberProxyRequest) (*DirectGroupMember, error)

	// Creates a local group in the account that contains this workspace. A
	// local group is a group that isn't synced from your identity provider
	// (IdP). You become a manager of the new group.
	//
	// If automatic identity management is on, you can't set `external_id`. To
	// add a group from your IdP, get it as an external group instead.
	CreateGroupProxy(ctx context.Context, request CreateGroupProxyRequest) (*Group, error)

	// Creates a local service principal in the account that contains this
	// workspace. A local service principal is one that isn't synced from your
	// identity provider (IdP).
	//
	// If automatic identity management is on, you can't set `external_id`. To
	// add a service principal from your IdP, get it as an external service
	// principal instead.
	CreateServicePrincipalProxy(ctx context.Context, request CreateServicePrincipalProxyRequest) (*ServicePrincipal, error)

	// Creates a local user in the account that contains this workspace. A local
	// user is a user that isn't synced from your identity provider (IdP).
	//
	// If automatic identity management is on, you can't set `external_id`. To
	// add a user from your IdP, get them as an external user instead.
	CreateUserProxy(ctx context.Context, request CreateUserProxyRequest) (*User, error)

	// Creates a workspace assignment detail for a principal in the calling
	// workspace.
	CreateWorkspaceAssignmentDetailProxy(ctx context.Context, request CreateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)

	// Assigns a user, service principal, or group directly to this workspace
	// with the given entitlements. Returns an error if the principal is already
	// assigned directly.
	//
	// Entitlements are granted one at a time. If the request fails partway
	// through, the principal stays assigned with only some of the entitlements.
	// Get the assignment afterwards to check which were granted.
	CreateWorkspaceAssignmentProxy(ctx context.Context, request CreateWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error)

	// Removes a direct member from a group in the account that contains this
	// workspace.
	//
	// If automatic identity management is on and the group has an
	// `external_id`, its members come from your identity provider and can't be
	// changed in Databricks.
	DeleteDirectGroupMemberProxy(ctx context.Context, request DeleteDirectGroupMemberProxyRequest) error

	// Deletes a group from the account that contains this workspace and removes
	// it from the workspaces it's assigned to.
	DeleteGroupProxy(ctx context.Context, request DeleteGroupProxyRequest) error

	// Deletes a service principal from the account that contains this workspace
	// and removes it from the workspaces it's assigned to. You can't delete
	// service principals that Databricks manages.
	DeleteServicePrincipalProxy(ctx context.Context, request DeleteServicePrincipalProxyRequest) error

	// Deletes a user from the account that contains this workspace and removes
	// them from the workspaces they're assigned to.
	DeleteUserProxy(ctx context.Context, request DeleteUserProxyRequest) error

	// Deletes a workspace assignment detail for a principal in the calling
	// workspace, revoking all of its entitlements. Entitlements are revoked one
	// at a time rather than atomically. If the request fails partway through,
	// the principal stays assigned with some of its original entitlements.
	// Retrying is safe.
	DeleteWorkspaceAssignmentDetailProxy(ctx context.Context, request DeleteWorkspaceAssignmentDetailProxyRequest) error

	// Removes a principal's direct assignment to this workspace and revokes the
	// entitlements it granted. The principal keeps any access it gets through
	// groups.
	//
	// Entitlements are revoked one at a time. If the request fails partway
	// through, the principal stays assigned with some of its entitlements. You
	// can retry the request.
	DeleteWorkspaceAssignmentProxy(ctx context.Context, request DeleteWorkspaceAssignmentProxyRequest) error

	// Gets a direct member of a group in the account that contains this
	// workspace, including where the membership comes from. Returns an error if
	// the principal isn't a direct member.
	GetDirectGroupMemberProxy(ctx context.Context, request GetDirectGroupMemberProxyRequest) (*DirectGroupMember, error)

	// Gets a group by its ID in your identity provider (IdP) and syncs its name
	// and parent groups from your IdP. If the group isn't in the account that
	// contains this workspace yet, it's added. Requires automatic identity
	// management. If your account has group visibility filters, the group must
	// match them.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	GetExternalGroupProxy(ctx context.Context, request GetExternalGroupProxyRequest) (*ExternalGroup, error)

	// Gets a service principal by its ID in your identity provider (IdP) and
	// syncs its details and group memberships from your IdP. If the service
	// principal isn't in the account that contains this workspace yet, it's
	// added. Requires automatic identity management with Microsoft Entra ID as
	// your IdP.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	GetExternalServicePrincipalProxy(ctx context.Context, request GetExternalServicePrincipalProxyRequest) (*ExternalServicePrincipal, error)

	// Gets a user by their ID in your identity provider (IdP) and syncs their
	// details and group memberships from your IdP. If the user isn't in the
	// account that contains this workspace yet, they're added. Requires
	// automatic identity management.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	GetExternalUserProxy(ctx context.Context, request GetExternalUserProxyRequest) (*ExternalUser, error)

	// Gets a group in the account that contains this workspace by its
	// Databricks ID.
	GetGroupProxy(ctx context.Context, request GetGroupProxyRequest) (*Group, error)

	// Gets a service principal in the account that contains this workspace by
	// its Databricks ID.
	GetServicePrincipalProxy(ctx context.Context, request GetServicePrincipalProxyRequest) (*ServicePrincipal, error)

	// Gets a user in the account that contains this workspace by their
	// Databricks ID.
	GetUserProxy(ctx context.Context, request GetUserProxyRequest) (*User, error)

	// Returns the access details for a principal in the current workspace.
	// Allows for checking access details for any provisioned principal (user,
	// service principal, or group) in the current workspace. * Provisioned
	// principal here refers to one that has been synced into Databricks from
	// the customer's IdP or added explicitly to Databricks via SCIM/UI. Allows
	// for passing in a "view" parameter to control what fields are returned
	// (BASIC by default or FULL).
	GetWorkspaceAccessDetailLocal(ctx context.Context, request GetWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error)

	// Returns the assignment details for a principal in the calling workspace.
	GetWorkspaceAssignmentDetailProxy(ctx context.Context, request GetWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)

	// Gets a principal's direct assignment to this workspace, including its
	// effective entitlements. Returns an error if the principal isn't assigned
	// directly.
	GetWorkspaceAssignmentProxy(ctx context.Context, request GetWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error)

	// Gets a principal's status in this workspace and whether it has access
	// directly or through a group. Returns an error if the principal doesn't
	// have access to the workspace.
	GetWorkspaceIdentityDetail(ctx context.Context, request GetWorkspaceIdentityDetailRequest) (*WorkspaceIdentityDetail, error)

	// Lists the direct members of a group in the account that contains this
	// workspace, sorted by display name, with where each membership comes from.
	// Results combine memberships added in Databricks (in the UI, through the
	// API, or through SCIM provisioning) with, if automatic identity management
	// is on and the group has an `external_id`, memberships from your identity
	// provider (IdP). Only members in your Databricks account are returned: an
	// IdP member of the group appears only if that user, service principal, or
	// group has been added to Databricks. Nested groups aren't expanded. To get
	// the full membership, list the direct members of each member group.
	ListDirectGroupMembersProxy(ctx context.Context, request ListDirectGroupMembersProxyRequest) (*ListDirectGroupMembersResponse, error)

	// Lists the groups in the account that contains this workspace, sorted by
	// Databricks ID. Groups in your identity provider that haven't been added
	// to Databricks aren't included.
	ListGroupsProxy(ctx context.Context, request ListGroupsProxyRequest) (*ListGroupsResponse, error)

	// Lists the service principals in the account that contains this workspace,
	// sorted by Databricks ID. Service principals in your identity provider
	// that haven't been added to Databricks aren't included.
	ListServicePrincipalsProxy(ctx context.Context, request ListServicePrincipalsProxyRequest) (*ListServicePrincipalsResponse, error)

	// Lists every group that a user, service principal, or group belongs to,
	// directly or through other groups, in the account that contains this
	// workspace. Results are sorted by Databricks ID and combine memberships
	// from every source: added in the UI, through the API, through SCIM
	// provisioning, or synced from your identity provider (IdP) by automatic
	// identity management. Only groups that exist in your Databricks account
	// are returned, not every group that the principal belongs to in your IdP.
	// If automatic identity management is on, only the principal's own group
	// memberships are refreshed from your IdP before the list is read.
	ListTransitiveParentGroupsProxy(ctx context.Context, request ListTransitiveParentGroupsProxyRequest) (*ListTransitiveParentGroupsResponse, error)

	// Lists the users in the account that contains this workspace, sorted by
	// Databricks ID. Users in your identity provider who haven't been added to
	// Databricks aren't included.
	ListUsersProxy(ctx context.Context, request ListUsersProxyRequest) (*ListUsersResponse, error)

	// Lists workspace assignment details for the calling workspace. The
	// response omits the per-principal entitlement fields (`entitlements` and
	// `effective_entitlements`). To read the entitlements for a single
	// principal, get that principal's assignment detail.
	ListWorkspaceAssignmentDetailsProxy(ctx context.Context, request ListWorkspaceAssignmentDetailsProxyRequest) (*ListWorkspaceAssignmentDetailsResponse, error)

	// Lists the principals assigned directly to this workspace. The response
	// leaves `entitlements` and `effective_entitlements` empty. To read a
	// principal's entitlements, get its assignment.
	ListWorkspaceAssignmentsProxy(ctx context.Context, request ListWorkspaceAssignmentsProxyRequest) (*ListWorkspaceAssignmentsResponse, error)

	// Gets a group by its ID in your identity provider (IdP) and syncs its name
	// and parent groups from your IdP. If the group isn't in the account that
	// contains this workspace yet, it's added. Requires automatic identity
	// management. If your account has group visibility filters, the group must
	// match them.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	ResolveGroupProxy(ctx context.Context, request ResolveGroupProxyRequest) (*ResolveGroupResponse, error)

	// Gets a service principal by its ID in your identity provider (IdP) and
	// syncs its details and group memberships from your IdP. If the service
	// principal isn't in the account that contains this workspace yet, it's
	// added. Requires automatic identity management with Microsoft Entra ID as
	// your IdP.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	ResolveServicePrincipalProxy(ctx context.Context, request ResolveServicePrincipalProxyRequest) (*ResolveServicePrincipalResponse, error)

	// Gets a user by their ID in your identity provider (IdP) and syncs their
	// details and group memberships from your IdP. If the user isn't in the
	// account that contains this workspace yet, they're added. Requires
	// automatic identity management.
	//
	// To control how fresh the synced data is, send a `Cache-Control` request
	// header:
	//
	// - `no-cache` syncs from your IdP before the call returns. - `max-age=<n>,
	// max-stale=0` syncs from your IdP before the call returns if the last sync
	// is more than `<n>` seconds old.
	//
	// Without the header, the call skips the sync if the last sync is within
	// the default freshness window. Databricks might sync more often than you
	// request. If you send `max-age` without `max-stale=0`, the call can return
	// the data from the last sync and sync in the background. Databricks
	// ignores other `Cache-Control` directives. A forced sync makes the call
	// wait for your IdP. A forced sync on every call adds load on your IdP and
	// can hit rate limits.
	ResolveUserProxy(ctx context.Context, request ResolveUserProxyRequest) (*ResolveUserResponse, error)

	// Updates a group in the account that contains this workspace. Only the
	// fields in `update_mask` change.
	//
	// If automatic identity management is on and the group has an
	// `external_id`, you can only update `external_id`. The group's other
	// fields come from your identity provider. If your account also has group
	// visibility filters, a new `external_id` must match them. You can always
	// clear `external_id`.
	UpdateGroupProxy(ctx context.Context, request UpdateGroupProxyRequest) (*Group, error)

	// Updates a service principal in the account that contains this workspace.
	// Only the fields in `update_mask` change.
	//
	// If automatic identity management is on and the service principal has an
	// `external_id`, you can only update `external_id`. Its other fields come
	// from your identity provider. A service principal can't deactivate itself,
	// and you can't update service principals that Databricks manages.
	UpdateServicePrincipalProxy(ctx context.Context, request UpdateServicePrincipalProxyRequest) (*ServicePrincipal, error)

	// Updates a user in the account that contains this workspace. Only the
	// fields in `update_mask` change.
	//
	// If automatic identity management is on and the user has an `external_id`,
	// you can only update `external_id`. The user's other fields come from your
	// identity provider. You can't deactivate yourself or change the status of
	// the account owner.
	UpdateUserProxy(ctx context.Context, request UpdateUserProxyRequest) (*User, error)

	// Updates the entitlements of a directly assigned principal in the calling
	// workspace. Changes are applied one at a time rather than atomically. If
	// the request fails partway through, only some of the requested changes
	// take effect. Get the assignment detail afterwards to confirm the final
	// state.
	UpdateWorkspaceAssignmentDetailProxy(ctx context.Context, request UpdateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error)

	// Replaces the entitlements of a principal's direct assignment to this
	// workspace. The new list can't be empty.
	//
	// Entitlements are changed one at a time. If the request fails partway
	// through, only some of the changes take effect. Get the assignment
	// afterwards to check the result.
	UpdateWorkspaceAssignmentProxy(ctx context.Context, request UpdateWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error)

	// Activates or deactivates a user or service principal in this workspace.
	// You can only update `workspace_identity_status`. You can't change your
	// own status or the status of a group.
	UpdateWorkspaceIdentityDetail(ctx context.Context, request UpdateWorkspaceIdentityDetailRequest) (*WorkspaceIdentityDetail, error)
}
