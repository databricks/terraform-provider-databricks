package model

// ObjectACL ...
type ObjectACL struct {
	ObjectID          string           `json:"object_id,omitempty"`
	ObjectType        string           `json:"object_type,omitempty"`
	AccessControlList []*AccessControl `json:"access_control_list"`
}

// AccessControlChangeList is wrapper around ACL changes
type AccessControlChangeList struct {
	AccessControlList []*AccessControlChange `json:"access_control_list"`
}

// AccessControlChange is API wrapper for changing permissions
type AccessControlChange struct {
	UserName             *string `json:"user_name,omitempty"`
	GroupName            *string `json:"group_name,omitempty"`
	ServicePrincipalName *string `json:"service_principal_name,omitempty"`
	PermissionLevel      string  `json:"permission_level"`
}

// AccessControl ...
type AccessControl struct {
	UserName       *string       `json:"user_name,omitempty"`
	GroupName      *string       `json:"group_name,omitempty"`
	AllPermissions []*Permission `json:"all_permissions,omitempty"`
}

// Permission ...
type Permission struct {
	PermissionLevel     string   `json:"permission_level"`
	Inherited           bool     `json:"inherited,omitempty"`
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
}

// ToAccessControlChangeList converts data formats
func (oa *ObjectACL) ToAccessControlChangeList() *AccessControlChangeList {
	acl := new(AccessControlChangeList)
	for _, accessControl := range oa.AccessControlList {
		for _, permission := range accessControl.AllPermissions {
			if permission.Inherited {
				continue
			}
			item := new(AccessControlChange)
			acl.AccessControlList = append(acl.AccessControlList, item)
			item.PermissionLevel = permission.PermissionLevel
			if accessControl.UserName != nil {
				item.UserName = accessControl.UserName
			} else if accessControl.GroupName != nil {
				item.GroupName = accessControl.GroupName
			}
		}
	}
	return acl
}

// AccessControl exports data for TF
func (acl *AccessControlChangeList) AccessControl(me string) []map[string]string {
	result := []map[string]string{}
	for _, control := range acl.AccessControlList {
		item := map[string]string{}
		if control.UserName != nil && *control.UserName != "" {
			if me == *control.UserName {
				continue
			}
			item["user_name"] = *control.UserName
		} else if control.GroupName != nil && *control.GroupName != "" {
			item["group_name"] = *control.GroupName
		}
		item["permission_level"] = control.PermissionLevel
		result = append(result, item)
	}
	return result
}
