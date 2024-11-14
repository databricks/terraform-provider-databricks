package read

import (
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/permissions/entity"
)

// Context that is available to aclReadCustomizer implementations.
type ACLCustomizerContext struct {
	GetId                        func() string
	GetExistingPermissionsEntity func() entity.PermissionsEntity
}

// ACLCustomizer is a function that modifies the access control list of an object after it is read.
type ACLCustomizer func(ctx ACLCustomizerContext, objectAcls iam.ObjectPermissions) iam.ObjectPermissions

// Rewrites the permission level of the access control list of an object after it is read.
// This is done only for resources in state where the permission level is equal to the replacement value
// in the mapping. For example, the permissons endpoint used to use the "CAN_VIEW" permission level for
// read-only access, but this was changed to "CAN_READ". Users who previously used "CAN_VIEW" should not
// be forced to change to "CAN_READ". This customizer will rewrite "CAN_READ" to "CAN_VIEW" when the
// user-specified value is CAN_VIEW and the API response is CAN_READ.
func RewritePermissions(mapping map[iam.PermissionLevel]iam.PermissionLevel) ACLCustomizer {
	findOriginalAcl := func(new iam.AccessControlResponse, original entity.PermissionsEntity) (iam.AccessControlRequest, bool) {
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
	return func(ctx ACLCustomizerContext, acl iam.ObjectPermissions) iam.ObjectPermissions {
		original := ctx.GetExistingPermissionsEntity()
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
