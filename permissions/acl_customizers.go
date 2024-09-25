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

func removeAdminPermissionsCustomizer(ctx aclReadCustomizerContext, acl iam.ObjectPermissions) iam.ObjectPermissions {
	// Remove all permissions for the 'admins' group.
	filteredAcl := make([]iam.AccessControlResponse, 0, len(acl.AccessControlList))
	for _, a := range acl.AccessControlList {
		if a.GroupName != "admins" {
			filteredAcl = append(filteredAcl, a)
		}
	}
	acl.AccessControlList = filteredAcl
	return acl
}

func rewritePermissionForUpdate(old, new iam.PermissionLevel) aclUpdateCustomizer {
	return func(ctx aclUpdateCustomizerContext, acl []iam.AccessControlRequest) ([]iam.AccessControlRequest, error) {
		for i := range acl {
			if acl[i].PermissionLevel == old {
				acl[i].PermissionLevel = new
			}
		}
		return acl, nil
	}
}

func rewritePermissionForRead(old, new iam.PermissionLevel) aclReadCustomizer {
	return func(ctx aclReadCustomizerContext, acl iam.ObjectPermissions) iam.ObjectPermissions {
		for i := range acl.AccessControlList {
			for j := range acl.AccessControlList[i].AllPermissions {
				if acl.AccessControlList[i].AllPermissions[j].PermissionLevel == old {
					acl.AccessControlList[i].AllPermissions[j].PermissionLevel = new
				}
			}
		}
		return acl
	}
}
