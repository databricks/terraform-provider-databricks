package iam

import "context"

// Groups simplify identity management, making it easier to assign access to
// Databricks account, data, and other securable objects.
//
// It is best practice to assign access to workspaces and access-control
// policies in Unity Catalog to groups, instead of to users individually. All
// Databricks account identities can be assigned as members of groups, and
// members inherit permissions that are assigned to their group.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountGroupsService interface {

	// Creates a group in the Databricks account with a unique name, using the
	// supplied group details.
	Create(ctx context.Context, request Group) (*Group, error)

	// Deletes a group from the Databricks account.
	Delete(ctx context.Context, request DeleteAccountGroupRequest) error

	// Gets the information for a specific group in the Databricks account.
	Get(ctx context.Context, request GetAccountGroupRequest) (*Group, error)

	// Gets all details of the groups associated with the Databricks account.
	List(ctx context.Context, request ListAccountGroupsRequest) (*ListGroupsResponse, error)

	// Partially updates the details of a group.
	Patch(ctx context.Context, request PartialUpdate) error

	// Updates the details of a group by replacing the entire group entity.
	Update(ctx context.Context, request Group) error
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountServicePrincipalsService interface {

	// Creates a new service principal in the Databricks account.
	Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error)

	// Delete a single service principal in the Databricks account.
	Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error

	// Gets the details for a single service principal define in the Databricks
	// account.
	Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*ServicePrincipal, error)

	// Gets the set of service principals associated with a Databricks account.
	List(ctx context.Context, request ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error)

	// Partially updates the details of a single service principal in the
	// Databricks account.
	Patch(ctx context.Context, request PartialUpdate) error

	// Updates the details of a single service principal.
	//
	// This action replaces the existing service principal with the same name.
	Update(ctx context.Context, request ServicePrincipal) error
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
type AccountUsersService interface {

	// Creates a new user in the Databricks account. This new user will also be
	// added to the Databricks account.
	Create(ctx context.Context, request User) (*User, error)

	// Deletes a user. Deleting a user from a Databricks account also removes
	// objects associated with the user.
	Delete(ctx context.Context, request DeleteAccountUserRequest) error

	// Gets information for a specific user in Databricks account.
	Get(ctx context.Context, request GetAccountUserRequest) (*User, error)

	// Gets details for all the users associated with a Databricks account.
	List(ctx context.Context, request ListAccountUsersRequest) (*ListUsersResponse, error)

	// Partially updates a user resource by applying the supplied operations on
	// specific user attributes.
	Patch(ctx context.Context, request PartialUpdate) error

	// Replaces a user's information with the data supplied in request.
	Update(ctx context.Context, request User) error
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
type GroupsService interface {

	// Creates a group in the Databricks workspace with a unique name, using the
	// supplied group details.
	Create(ctx context.Context, request Group) (*Group, error)

	// Deletes a group from the Databricks workspace.
	Delete(ctx context.Context, request DeleteGroupRequest) error

	// Gets the information for a specific group in the Databricks workspace.
	Get(ctx context.Context, request GetGroupRequest) (*Group, error)

	// Gets all details of the groups associated with the Databricks workspace.
	List(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error)

	// Partially updates the details of a group.
	Patch(ctx context.Context, request PartialUpdate) error

	// Updates the details of a group by replacing the entire group entity.
	Update(ctx context.Context, request Group) error
}

// Identities for use with jobs, automated tools, and systems such as scripts,
// apps, and CI/CD platforms. Databricks recommends creating service principals
// to run production jobs or modify production data. If all processes that act
// on production data run with service principals, interactive users do not need
// any write, delete, or modify privileges in production. This eliminates the
// risk of a user overwriting production data by accident.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ServicePrincipalsService interface {

	// Creates a new service principal in the Databricks workspace.
	Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error)

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
	Patch(ctx context.Context, request PartialUpdate) error

	// Updates the details of a single service principal.
	//
	// This action replaces the existing service principal with the same name.
	Update(ctx context.Context, request ServicePrincipal) error
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
type UsersService interface {

	// Creates a new user in the Databricks workspace. This new user will also
	// be added to the Databricks account.
	Create(ctx context.Context, request User) (*User, error)

	// Deletes a user. Deleting a user from a Databricks workspace also removes
	// objects associated with the user.
	Delete(ctx context.Context, request DeleteUserRequest) error

	// Gets information for a specific user in Databricks workspace.
	Get(ctx context.Context, request GetUserRequest) (*User, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context) (*GetPasswordPermissionLevelsResponse, error)

	// Gets the permissions of all passwords. Passwords can inherit permissions
	// from their root object.
	GetPermissions(ctx context.Context) (*PasswordPermissions, error)

	// Gets details for all the users associated with a Databricks workspace.
	List(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error)

	// Partially updates a user resource by applying the supplied operations on
	// specific user attributes.
	Patch(ctx context.Context, request PartialUpdate) error

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error)

	// Replaces a user's information with the data supplied in request.
	Update(ctx context.Context, request User) error

	// Updates the permissions on all passwords. Passwords can inherit
	// permissions from their root object.
	UpdatePermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error)
}
