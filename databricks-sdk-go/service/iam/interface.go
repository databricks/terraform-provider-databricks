// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"context"
)

// Rule based Access Control for Databricks Resources.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccessControlService interface {

	// Check access policy to a resource.
	CheckPolicy(ctx context.Context, request CheckPolicyRequest) (*CheckPolicyResponse, error)
}

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountAccessControlService interface {

	// Gets all the roles that can be granted on an account level resource. A
	// role is grantable if the rule set on the resource can contain an access
	// rule of the role.
	GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error)

	// Get a rule set by its name. A rule set is always attached to a resource
	// and contains a list of access rules on the said resource. Currently only
	// a default rule set for each resource is supported.
	GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error)

	// Replace the rules of a rule set. First, use get to read the current
	// version of the rule set before modifying it. This pattern helps prevent
	// conflicts between concurrent updates.
	UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error)
}

// These APIs manage access rules on resources in an account. Currently, only
// grant rules are supported. A grant rule specifies a role assigned to a set of
// principals. A list of rules attached to a resource is called a rule set. A
// workspace must belong to an account for these APIs to work
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountAccessControlProxyService interface {

	// Gets all the roles that can be granted on an account level resource. A
	// role is grantable if the rule set on the resource can contain an access
	// rule of the role.
	GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error)

	// Get a rule set by its name. A rule set is always attached to a resource
	// and contains a list of access rules on the said resource. Currently only
	// a default rule set for each resource is supported.
	GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error)

	// Replace the rules of a rule set. First, use get to read the current
	// version of the rule set before modifying it. This pattern helps prevent
	// conflicts between concurrent updates.
	UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error)
}

// Groups simplify identity management, making it easier to assign access to
// Databricks account, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks account identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountGroupsV2Service interface {

	// Creates a group in the Databricks account with a unique name, using the
	// supplied group details.
	Create(ctx context.Context, request CreateAccountGroupRequest) (*AccountGroup, error)

	// Deletes a group from the Databricks account.
	Delete(ctx context.Context, request DeleteAccountGroupRequest) error

	// Gets the information for a specific group in the Databricks account.
	Get(ctx context.Context, request GetAccountGroupRequest) (*AccountGroup, error)

	// Gets all details of the groups associated with the Databricks account. As
	// of 08/22/2025, this endpoint will no longer return members. Instead,
	// members should be retrieved by iterating through `Get group details`.
	// Existing accounts that rely on this attribute will not be impacted and
	// will continue receiving member data as before.
	List(ctx context.Context, request ListAccountGroupsRequest) (*ListAccountGroupsResponse, error)

	// Partially updates the details of a group.
	Patch(ctx context.Context, request PatchAccountGroupRequest) error

	// Updates the details of a group by replacing the entire group entity.
	Update(ctx context.Context, request UpdateAccountGroupRequest) error
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountServicePrincipalsV2Service interface {

	// Creates a new service principal in the Databricks account.
	Create(ctx context.Context, request CreateAccountServicePrincipalRequest) (*AccountServicePrincipal, error)

	// Deletes a single service principal from the Databricks account.
	Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error

	// Gets the details for a single service principal in the Databricks
	// account.
	Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*AccountServicePrincipal, error)

	// Gets the set of service principals associated with a Databricks account.
	List(ctx context.Context, request ListAccountServicePrincipalsRequest) (*ListAccountServicePrincipalsResponse, error)

	// Partially updates the details of a single service principal in the
	// Databricks account.
	Patch(ctx context.Context, request PatchAccountServicePrincipalRequest) error

	// Updates the details of a single service principal.
	//
	// This action replaces the existing service principal that has the given
	// ID.
	Update(ctx context.Context, request UpdateAccountServicePrincipalRequest) error
}

// User identities recognized by Databricks and represented by email addresses.
//
// Databricks recommends using SCIM provisioning to sync users and groups
// automatically from your identity provider to your Databricks account. SCIM
// streamlines onboarding a new employee or team by using your identity provider
// to create users and groups in Databricks account and give them the proper
// level of access. When a user leaves your organization or no longer needs
// access to Databricks account, admins can terminate the user in your identity
// provider and that user’s account will also be removed from Databricks
// account. This ensures a consistent offboarding process and prevents
// unauthorized users from accessing sensitive data.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountUsersV2Service interface {

	// Creates a new user in the Databricks account. This new user will also be
	// added to the Databricks account.
	Create(ctx context.Context, request CreateAccountUserRequest) (*AccountUser, error)

	// Deletes a user. Deleting a user from a Databricks account also removes
	// objects associated with the user.
	Delete(ctx context.Context, request DeleteAccountUserRequest) error

	// Gets information for a specific user in Databricks account.
	Get(ctx context.Context, request GetAccountUserRequest) (*AccountUser, error)

	// Gets details for all the users associated with a Databricks account.
	List(ctx context.Context, request ListAccountUsersRequest) (*ListAccountUsersResponse, error)

	// Partially updates a user resource by applying the supplied operations on
	// specific user attributes. The `userName` and `emails` attributes cannot
	// be updated through this API; any supplied changes to them are ignored
	// (no-op).
	Patch(ctx context.Context, request PatchAccountUserRequest) error

	// Replaces a user's information with the data supplied in request. The
	// `userName` and `emails` attributes cannot be updated through this API;
	// any supplied changes to them are ignored (no-op).
	Update(ctx context.Context, request UpdateAccountUserRequest) error
}

// Manage the workspace's restricted sharing groups.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AssumeGroupService interface {

	// Adds a group to the workspace's restricted list. The workspace allows at
	// most 100 restricted sharing groups. Requires workspace admin.
	CreateRestrictedSharingGroup(ctx context.Context, request CreateRestrictedSharingGroupRequest) (*RestrictedSharingGroup, error)

	// Removes a group from the workspace's restricted list. Requires workspace
	// admin.
	DeleteRestrictedSharingGroup(ctx context.Context, request DeleteRestrictedSharingGroupRequest) error

	// Gets a restricted sharing group in the workspace. Use it to check whether
	// a group is on the workspace's restricted list. Requires workspace admin.
	GetRestrictedSharingGroup(ctx context.Context, request GetRestrictedSharingGroupRequest) (*RestrictedSharingGroup, error)

	// Lists the restricted sharing groups in the workspace. Requires workspace
	// admin.
	ListRestrictedSharingGroups(ctx context.Context, request ListRestrictedSharingGroupsRequest) (*ListRestrictedSharingGroupsResponse, error)
}

// This API allows retrieving information about currently authenticated user or
// service principal.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CurrentUserService interface {

	// Get details about the current method caller's identity.
	Me(ctx context.Context, request MeRequest) (*User, error)
}

// Groups simplify identity management, making it easier to assign access to
// Databricks workspace, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks workspace identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type GroupsV2Service interface {

	// Creates a group in the Databricks workspace with a unique name, using the
	// supplied group details.
	Create(ctx context.Context, request CreateGroupRequest) (*Group, error)

	// Deletes a group from the Databricks workspace.
	Delete(ctx context.Context, request DeleteGroupRequest) error

	// Gets the information for a specific group in the Databricks workspace.
	Get(ctx context.Context, request GetGroupRequest) (*Group, error)

	// Gets all details of the groups associated with the Databricks workspace.
	List(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error)

	// Partially updates the details of a group.
	Patch(ctx context.Context, request PatchGroupRequest) error

	// Updates the details of a group by replacing the entire group entity.
	Update(ctx context.Context, request UpdateGroupRequest) error
}

// APIs for migrating acl permissions, used only by the ucx tool:
// https://github.com/databrickslabs/ucx
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PermissionMigrationService interface {

	// Migrate Permissions.
	MigratePermissions(ctx context.Context, request MigratePermissionsRequest) (*MigratePermissionsResponse, error)
}

// Permissions API are used to create read, write, edit, update and manage
// access for various users on different objects and endpoints. * **[Apps
// permissions](:service:apps)** — Manage which users can manage or use apps.
// * **[Cluster permissions](:service:clusters)** — Manage which users can
// manage, restart, or attach to clusters. * **[Cluster policy
// permissions](:service:clusterpolicies)** — Manage which users can use
// cluster policies. * **[Spark Declarative Pipelines
// permissions](:service:pipelines)** — Manage which users can view, manage,
// run, cancel, or own a Spark Declarative Pipeline. * **[Job
// permissions](:service:jobs)** — Manage which users can view, manage,
// trigger, cancel, or own a job. * **[MLflow experiment
// permissions](:service:experiments)** — Manage which users can read, edit,
// or manage MLflow experiments. * **[MLflow registered model
// permissions](:service:modelregistry)** — Manage which users can read, edit,
// or manage MLflow registered models. * **[Instance Pool
// permissions](:service:instancepools)** — Manage which users can manage or
// attach to pools. * **[Repo permissions](repos)** — Manage which users can
// read, run, edit, or manage a repo. * **[Serving endpoint
// permissions](:service:servingendpoints)** — Manage which users can view,
// query, or manage a serving endpoint. * **[SQL warehouse
// permissions](:service:warehouses)** — Manage which users can use or manage
// SQL warehouses. * **[Token permissions](:service:tokenmanagement)** —
// Manage which users can create or use tokens. * **[Workspace object
// permissions](:service:workspace)** — Manage which users can read, run,
// edit, or manage alerts, dbsql-dashboards, directories, files, notebooks and
// queries. For the mapping of the required permissions for specific actions or
// abilities and other important information, see [Access Control]. Note that to
// manage access control on service principals, use **[Account Access Control
// Proxy](:service:accountaccesscontrolproxy)**.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Access Control]: https://docs.databricks.com/security/auth-authz/access-control/index.html
type PermissionsService interface {

	// Gets the permissions of an object. Objects can inherit permissions from
	// their parent objects or root object.
	Get(ctx context.Context, request GetPermissionRequest) (*ObjectPermissions, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error)

	// Sets permissions on an object, replacing all direct permissions with the
	// ones in the request. Principals that are not in the request lose their
	// direct permissions on the object, and all direct permissions are deleted
	// if none are specified. Objects can inherit permissions from their parent
	// objects or root object.
	Set(ctx context.Context, request SetObjectPermissions) (*ObjectPermissions, error)

	// Updates the permissions on an object by adding or changing the
	// permissions of the principals in the request. Principals that are not in
	// the request keep their existing permissions. Objects can inherit
	// permissions from their parent objects or root object.
	Update(ctx context.Context, request UpdateObjectPermissions) (*ObjectPermissions, error)
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ServicePrincipalsV2Service interface {

	// Creates a new service principal in the Databricks workspace.
	Create(ctx context.Context, request CreateServicePrincipalRequest) (*ServicePrincipal, error)

	// Delete a single service principal in the Databricks workspace.
	Delete(ctx context.Context, request DeleteServicePrincipalRequest) error

	// Gets the details for a single service principal define in the Databricks
	// workspace.
	Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error)

	// Gets the set of service principals associated with a Databricks
	// workspace.
	List(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error)

	// Partially updates the details of a single service principal in the
	// Databricks workspace.
	Patch(ctx context.Context, request PatchServicePrincipalRequest) error

	// Updates the details of a single service principal.
	//
	// This action replaces the existing service principal with the same name.
	Update(ctx context.Context, request UpdateServicePrincipalRequest) error
}

// User identities recognized by Databricks and represented by email addresses.
//
// Databricks recommends using SCIM provisioning to sync users and groups
// automatically from your identity provider to your Databricks workspace. SCIM
// streamlines onboarding a new employee or team by using your identity provider
// to create users and groups in Databricks workspace and give them the proper
// level of access. When a user leaves your organization or no longer needs
// access to Databricks workspace, admins can terminate the user in your
// identity provider and that user’s account will also be removed from
// Databricks workspace. This ensures a consistent offboarding process and
// prevents unauthorized users from accessing sensitive data.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type UsersV2Service interface {

	// Creates a new user in the Databricks workspace. This new user will also
	// be added to the Databricks account.
	Create(ctx context.Context, request CreateUserRequest) (*User, error)

	// Deletes a user. Deleting a user from a Databricks workspace also removes
	// objects associated with the user.
	Delete(ctx context.Context, request DeleteUserRequest) error

	// Gets information for a specific user in Databricks workspace.
	Get(ctx context.Context, request GetUserRequest) (*User, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetPasswordPermissionLevelsRequest) (*GetPasswordPermissionLevelsResponse, error)

	// Gets the permissions of all passwords. Passwords can inherit permissions
	// from their root object.
	GetPermissions(ctx context.Context, request GetPasswordPermissionsRequest) (*PasswordPermissions, error)

	// Gets details for all the users associated with a Databricks workspace.
	List(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error)

	// Partially updates a user resource by applying the supplied operations on
	// specific user attributes. The `userName` and `emails` attributes cannot
	// be updated through this API; any supplied changes to them are ignored
	// (no-op).
	Patch(ctx context.Context, request PatchUserRequest) error

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error)

	// Replaces a user's information with the data supplied in request. The
	// `userName` and `emails` attributes cannot be updated through this API;
	// any supplied changes to them are ignored (no-op).
	Update(ctx context.Context, request UpdateUserRequest) error

	// Updates the permissions on all passwords. Passwords can inherit
	// permissions from their root object.
	UpdatePermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error)
}

// The Workspace Permission Assignment API allows you to manage workspace
// permissions for principals in your account.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceAssignmentService interface {

	// Deletes the workspace permissions assignment in a given account and
	// workspace for the specified principal.
	Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error

	// Get an array of workspace permissions for the specified account and
	// workspace.
	Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error)

	// Get the permission assignments for the specified Databricks account and
	// Databricks workspace.
	List(ctx context.Context, request ListWorkspaceAssignmentRequest) (*PermissionAssignments, error)

	// Creates or updates the workspace permissions assignment in a given
	// account and workspace for the specified principal.
	Update(ctx context.Context, request UpdateWorkspaceAssignments) (*PermissionAssignment, error)
}
