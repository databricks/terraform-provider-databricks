package permissions

import (
	"github.com/databricks/databricks-sdk-go/service/iam"
)

// Context that is available to aclUpdateCustomizer implementations.
type aclUpdateCustomizerContext struct {
	getCurrentUser func() (string, error)
	getId          func() string
}

// aclUpdateCustomizer is a function that modifies the access control list of an object before it is updated.
type aclUpdateCustomizer func(ctx aclUpdateCustomizerContext, objectAcls []iam.AccessControlRequest) ([]iam.AccessControlRequest, error)

// Context that is available to aclReadCustomizer implementations.
type aclReadCustomizerContext struct {
	getId                        func() string
	getExistingPermissionsEntity func() PermissionsEntity
}

// aclReadCustomizer is a function that modifies the access control list of an object after it is read.
type aclReadCustomizer func(ctx aclReadCustomizerContext, objectAcls iam.ObjectPermissions) iam.ObjectPermissions

// addAdminAclCustomizer adds an explicit CAN_MANAGE permission for the 'admins' group if explicitAdminPermissionCheck returns true
// for the provided object ID.
func addAdminAclCustomizer(explicitAdminPermissionCheck func(string) bool) aclUpdateCustomizer {
	return func(ctx aclUpdateCustomizerContext, acl []iam.AccessControlRequest) ([]iam.AccessControlRequest, error) {
		if explicitAdminPermissionCheck(ctx.getId()) {
			// Prevent "Cannot change permissions for group 'admins' to None."
			acl = append(acl, iam.AccessControlRequest{
				GroupName:       "admins",
				PermissionLevel: "CAN_MANAGE",
			})
		}
		return acl, nil
	}
}

// Whether the object requires explicit manage permissions for the calling user if not set.
// As described in https://github.com/databricks/terraform-provider-databricks/issues/1504,
// certain object types require that we explicitly grant the calling user CAN_MANAGE
// permissions when POSTing permissions changes through the REST API, to avoid accidentally
// revoking the calling user's ability to manage the current object.
func addCurrentUserAsManageCustomizer(ctx aclUpdateCustomizerContext, acl []iam.AccessControlRequest) ([]iam.AccessControlRequest, error) {
	currentUser, err := ctx.getCurrentUser()
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

func rewritePermissionForUpdate(mapping map[iam.PermissionLevel]iam.PermissionLevel) aclUpdateCustomizer {
	return func(ctx aclUpdateCustomizerContext, acl []iam.AccessControlRequest) ([]iam.AccessControlRequest, error) {
		for i := range acl {
			if new, ok := mapping[acl[i].PermissionLevel]; ok {
				acl[i].PermissionLevel = new
			}
		}
		return acl, nil
	}
}

// Rewrites the permission level of the access control list of an object after it is read.
// This is done only for resources in state where the permission level is equal to the replacement value
// in the mapping. For example, the permissons endpoint used to use the "CAN_VIEW" permission level for
// read-only access, but this was changed to "CAN_READ". Users who previously used "CAN_VIEW" should not
// be forced to change to "CAN_READ". This customizer will rewrite "CAN_READ" to "CAN_VIEW" when the
// user-specified value is CAN_VIEW and the API response is CAN_READ.
func rewritePermissionForRead(mapping map[iam.PermissionLevel]iam.PermissionLevel) aclReadCustomizer {
	findOriginalAcl := func(new iam.AccessControlResponse, original PermissionsEntity) (iam.AccessControlRequest, bool) {
		for _, old := range original.AccessControlList {
			if new.GroupName != "" && old.GroupName == new.GroupName {
				return old, true
			}
			if new.UserName != "" && old.UserName == new.UserName {
				return old, true
			}
			if new.ServicePrincipalName != "" && old.ServicePrincipalName == new.ServicePrincipalName {
				return old, true
			}
		}
		return iam.AccessControlRequest{}, false
	}
	return func(ctx aclReadCustomizerContext, acl iam.ObjectPermissions) iam.ObjectPermissions {
		original := ctx.getExistingPermissionsEntity()
		for i := range acl.AccessControlList {
			inState, found := findOriginalAcl(acl.AccessControlList[i], original)
			for j := range acl.AccessControlList[i].AllPermissions {
				// If the original permission level is remapped to a replacement level, and the permission level
				// in state is equal to the replacement level, we rewrite it to the replacement level.
				original := acl.AccessControlList[i].AllPermissions[j].PermissionLevel
				replacement, ok := mapping[original]
				if ok && found && inState.PermissionLevel == replacement {
					acl.AccessControlList[i].AllPermissions[j].PermissionLevel = replacement
				}
			}
		}
		return acl
	}
}
