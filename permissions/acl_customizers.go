package permissions

import "github.com/databricks/terraform-provider-databricks/common"

func addAdminAclCustomizer(explicitAdminPermissionCheck func(string) bool) aclUpdateCustomizer {
	return func(ctx aclUpdateCustomizerContext, acl []AccessControlChangeApiRequest) ([]AccessControlChangeApiRequest, error) {
		if explicitAdminPermissionCheck(ctx.getId()) {
			// Prevent "Cannot change permissions for group 'admins' to None."
			acl = append(acl, AccessControlChangeApiRequest{
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
func addCurrentUserAsManageCustomizer(ctx aclUpdateCustomizerContext, acl []AccessControlChangeApiRequest) ([]AccessControlChangeApiRequest, error) {
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
		acl = append(acl, AccessControlChangeApiRequest{
			UserName:        currentUser,
			PermissionLevel: "CAN_MANAGE",
		})
	}
	return acl, nil
}

func copyUserToServicePrincipalCustomizer(ctx aclReadCustomizerContext, objectAcls ObjectAclApiResponse) (ObjectAclApiResponse, error) {
	for i, acl := range objectAcls.AccessControlList {
		// If the username is a UUID, it's probably a service principal.
		if common.StringIsUUID(acl.UserName) {
			objectAcls.AccessControlList[i].ServicePrincipalName = acl.UserName
			objectAcls.AccessControlList[i].UserName = ""
		}
	}
	return objectAcls, nil
}

func copyServicePrincipalToUserCustomizer(ctx aclUpdateCustomizerContext, objectAcls []AccessControlChangeApiRequest) ([]AccessControlChangeApiRequest, error) {
	acl := make([]AccessControlChangeApiRequest, 0, len(objectAcls))
	for _, change := range objectAcls {
		if change.ServicePrincipalName != "" {
			acl = append(acl, AccessControlChangeApiRequest{
				UserName:        change.ServicePrincipalName,
				PermissionLevel: change.PermissionLevel,
			})
		} else {
			acl = append(acl, change)
		}
	}
	return acl, nil
}
