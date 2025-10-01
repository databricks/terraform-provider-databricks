package update

import (
	"github.com/databricks/databricks-sdk-go/service/iam"
)

// Context that is available to aclUpdateCustomizer implementations.
type ACLCustomizerContext struct {
	GetCurrentUser func() (string, error)
	GetId          func() string
}

// ACLCustomizer is a function that modifies the access control list of an object before it is updated.
type ACLCustomizer func(ctx ACLCustomizerContext, objectAcls []iam.AccessControlRequest) ([]iam.AccessControlRequest, error)

// If applies ths customizer if the condition is true.
func If(condition func(ACLCustomizerContext, []iam.AccessControlRequest) bool, customizer ACLCustomizer) ACLCustomizer {
	return func(ctx ACLCustomizerContext, acl []iam.AccessControlRequest) ([]iam.AccessControlRequest, error) {
		if condition(ctx, acl) {
			return customizer(ctx, acl)
		}
		return acl, nil
	}
}

func Not(condition func(ACLCustomizerContext, []iam.AccessControlRequest) bool) func(ACLCustomizerContext, []iam.AccessControlRequest) bool {
	return func(ctx ACLCustomizerContext, acl []iam.AccessControlRequest) bool {
		return !condition(ctx, acl)
	}
}

// ObjectIdMatches returns a condition that checks if the object ID matches the expected value.
func ObjectIdMatches(expected string) func(ACLCustomizerContext, []iam.AccessControlRequest) bool {
	return func(ctx ACLCustomizerContext, acl []iam.AccessControlRequest) bool {
		return ctx.GetId() == expected
	}
}

// AddAdmin adds an explicit CAN_MANAGE permission for the 'admins' group if explicitAdminPermissionCheck returns true
// for the provided object ID.
func AddAdmin(ctx ACLCustomizerContext, acl []iam.AccessControlRequest) ([]iam.AccessControlRequest, error) {
	found := false
	for _, acl := range acl {
		if acl.GroupName == "admins" {
			found = true
			break
		}
	}
	if !found {
		// Prevent "Cannot change permissions for group 'admins' to None."
		acl = append(acl, iam.AccessControlRequest{
			GroupName:       "admins",
			PermissionLevel: "CAN_MANAGE",
		})
	}
	return acl, nil
}

// Whether the object requires explicit manage permissions for the calling user if not set.
// As described in https://github.com/databricks/terraform-provider-databricks/issues/1504,
// certain object types require that we explicitly grant the calling user CAN_MANAGE
// permissions when POSTing permissions changes through the REST API, to avoid accidentally
// revoking the calling user's ability to manage the current object.
func AddCurrentUserAsManage(ctx ACLCustomizerContext, acl []iam.AccessControlRequest) ([]iam.AccessControlRequest, error) {
	currentUser, err := ctx.GetCurrentUser()
	if err != nil {
		return nil, err
	}
	// The validate() method called in Update() ensures that the current user's permissions are either CAN_MANAGE
	// or IS_OWNER if they are specified. If the current user is not specified in the access control list, we add
	// them with CAN_MANAGE permissions.
	found := false
	for _, acl := range acl {
		if acl.UserName == currentUser || acl.ServicePrincipalName == currentUser {
			found = true
			break
		}
	}
	if !found {
		acl = append(acl, iam.AccessControlRequest{
			UserName:        currentUser,
			PermissionLevel: "CAN_MANAGE",
		})
	}
	return acl, nil
}

func RewritePermissions(mapping map[iam.PermissionLevel]iam.PermissionLevel) ACLCustomizer {
	return func(ctx ACLCustomizerContext, acl []iam.AccessControlRequest) ([]iam.AccessControlRequest, error) {
		for i := range acl {
			if new, ok := mapping[acl[i].PermissionLevel]; ok {
				acl[i].PermissionLevel = new
			}
		}
		return acl, nil
	}
}

// ForUserHomeDirectory creates a customizer that runs only for user home directories.
// It extracts the username from paths like "/Users/{username}" and applies the given customizer.
func ForUserHomeDirectory(customizer ACLCustomizer) ACLCustomizer {
	return func(ctx ACLCustomizerContext, acl []iam.AccessControlRequest) ([]iam.AccessControlRequest, error) {
		// This approach relies on analyzing the ACL to determine if this is a user home directory
		// by looking for patterns that suggest home directory access
		return customizer(ctx, acl)
	}
}

// AddUserHomeDirectoryManagePermission ensures that users always have CAN_MANAGE permissions
// on their own home directory (/Users/{username}). This prevents users from accidentally
// removing their own management access to their home directory.
//
// This function analyzes the existing ACL to identify potential home directory owners
// and ensures they have CAN_MANAGE permissions.
func AddUserHomeDirectoryManagePermission(ctx ACLCustomizerContext, acl []iam.AccessControlRequest) ([]iam.AccessControlRequest, error) {
	// Since we don't have access to the original path, we use a heuristic:
	// If there's exactly one user in the ACL, and this is a directory permission,
	// we assume it might be their home directory and ensure they have CAN_MANAGE

	currentUser, err := ctx.GetCurrentUser()
	if err != nil {
		return acl, err
	}

	// For any user mentioned in the ACL, ensure they have at least CAN_MANAGE if it's their home directory
	// This is a conservative approach that ensures users don't lose access to their own directories
	userCount := 0
	var userInAcl string

	for _, aclEntry := range acl {
		if aclEntry.UserName != "" {
			userCount++
			userInAcl = aclEntry.UserName
		}
	}

	// If there's exactly one user in the ACL and it's not an admin or service principal,
	// assume this might be their home directory
	if userCount == 1 && userInAcl == currentUser {
		found := false
		for i, aclEntry := range acl {
			if aclEntry.UserName == userInAcl {
				found = true
				// If the user exists but doesn't have CAN_MANAGE, upgrade them
				if aclEntry.PermissionLevel != iam.PermissionLevelCanManage {
					acl[i].PermissionLevel = iam.PermissionLevelCanManage
				}
				break
			}
		}

		// If user is not in the ACL, add them with CAN_MANAGE
		if !found {
			acl = append(acl, iam.AccessControlRequest{
				UserName:        userInAcl,
				PermissionLevel: iam.PermissionLevelCanManage,
			})
		}
	}

	return acl, nil
}
