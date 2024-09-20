package permissions

import "github.com/databricks/terraform-provider-databricks/common"

// Context that is available to aclUpdateCustomizer implementations.
type aclUpdateCustomizerContext struct {
	getCurrentUser func() (string, error)
	getId          func() string
}

// aclUpdateCustomizer is a function that modifies the access control list of an object before it is updated.
type aclUpdateCustomizer func(ctx aclUpdateCustomizerContext, objectAcls []AccessControlChangeApiRequest) ([]AccessControlChangeApiRequest, error)

// Context that is available to aclReadCustomizer implementations.
type aclReadCustomizerContext struct {
	getId func() string
}

// aclReadCustomizer is a function that modifies the access control list of an object after it is read.
type aclReadCustomizer func(ctx aclReadCustomizerContext, objectAcls ObjectAclApiResponse) (ObjectAclApiResponse, error)

// addAdminAclCustomizer adds an explicit CAN_MANAGE permission for the 'admins' group if explicitAdminPermissionCheck returns true
// for the provided object ID.
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

// Copies the username to service principal name if the username is a UUID.
// The SQL permissions API only accepts usernames and determines on the backend if the provided username
// is actually a service principal ID. The API puts service principal IDs in the username field, so we need
// to copy the ID to the service_principal_name field when handling the response.
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

// Copies the username to service principal name if the username is a UUID.
// The SQL permissions API only accepts usernames and determines on the backend if the provided username
// is actually a service principal ID. Users may still specify service_principal_id directly, so we need
// to copy the ID to the user_name field before making a request.
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
