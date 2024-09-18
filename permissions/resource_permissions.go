package permissions

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ObjectAclApiResponse is a structure to generically describe access control.
// It represents the responses from the permissions APIs.
type ObjectAclApiResponse struct {
	ObjectID          string                     `json:"object_id,omitempty"`
	ObjectType        string                     `json:"object_type,omitempty"`
	AccessControlList []AccessControlApiResponse `json:"access_control_list"`
}

func (oa ObjectAclApiResponse) ToPermissionsEntity(d *schema.ResourceData, existing PermissionsEntity, me string) (PermissionsEntity, error) {
	entity := PermissionsEntity{}
	mapping, _, err := getResourcePermissionsForObjectAcl(oa)
	if err != nil {
		return entity, err
	}
	for _, accessControl := range oa.AccessControlList {
		if accessControl.GroupName == "admins" && !mapping.allowAdminGroup {
			// admin permission is always returned but can only be explicitly set for certain resources
			// For other resources, admin permissions are not included in the resource state.
			continue
		}
		if me == accessControl.UserName || me == accessControl.ServicePrincipalName {
			// If the user doesn't include an access_control block for themselves, do not include it in the state.
			// On create/update, the provider will automatically include the current user in the access_control block
			// for appropriate resources. Otherwise, it must be included in state to prevent configuration drift.
			if !existing.containsUserOrServicePrincipal(me) {
				continue
			}
		}
		if change, direct := accessControl.toAccessControlChange(); direct {
			entity.AccessControlList = append(entity.AccessControlList, change)
		}
	}
	entity.ObjectType = mapping.objectType
	pathVariant := d.Get(mapping.getPathVariant())
	if pathVariant != nil && pathVariant.(string) != "" {
		// we're not importing and it's a path... it's set, so let's not re-set it
		return entity, nil
	}
	identifier := path.Base(oa.ObjectID)
	return entity, d.Set(mapping.field, identifier)
}

// AccessControlApiResponse is a structure to describe user/group permissions.
type AccessControlApiResponse struct {
	UserName             string                  `json:"user_name,omitempty"`
	GroupName            string                  `json:"group_name,omitempty"`
	ServicePrincipalName string                  `json:"service_principal_name,omitempty"`
	AllPermissions       []PermissionApiResponse `json:"all_permissions,omitempty"`

	// SQLA entities don't use the `all_permissions` nesting, but rather a simple
	// top level string with the permission level when retrieving permissions.
	PermissionLevel string `json:"permission_level,omitempty"`
}

func (ac AccessControlApiResponse) toAccessControlChange() (AccessControlChangeApiRequest, bool) {
	for _, permission := range ac.AllPermissions {
		if permission.Inherited {
			continue
		}
		return AccessControlChangeApiRequest{
			PermissionLevel:      permission.PermissionLevel,
			UserName:             ac.UserName,
			GroupName:            ac.GroupName,
			ServicePrincipalName: ac.ServicePrincipalName,
		}, true
	}
	if ac.PermissionLevel != "" {
		return AccessControlChangeApiRequest{
			PermissionLevel:      ac.PermissionLevel,
			UserName:             ac.UserName,
			GroupName:            ac.GroupName,
			ServicePrincipalName: ac.ServicePrincipalName,
		}, true
	}
	return AccessControlChangeApiRequest{}, false
}

func (ac AccessControlApiResponse) String() string {
	return fmt.Sprintf("%s%s%s%v", ac.GroupName, ac.UserName, ac.ServicePrincipalName, ac.AllPermissions)
}

// PermissionApiResponse is a structure to describe permission level
type PermissionApiResponse struct {
	PermissionLevel     string   `json:"permission_level"`
	Inherited           bool     `json:"inherited,omitempty"`
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
}

func (p PermissionApiResponse) String() string {
	if len(p.InheritedFromObject) > 0 {
		return fmt.Sprintf("%s (from %s)", p.PermissionLevel, p.InheritedFromObject)
	}
	return p.PermissionLevel
}

// AccessControlChangeListApiRequest is wrapper around ACL changes for REST API
// This is the structure expected by the REST API when changing permissions.
type AccessControlChangeListApiRequest struct {
	AccessControlList []AccessControlChangeApiRequest `json:"access_control_list"`
}

// AccessControlChangeApiRequest is API wrapper for changing permissions
type AccessControlChangeApiRequest struct {
	UserName             string `json:"user_name,omitempty"`
	GroupName            string `json:"group_name,omitempty"`
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	PermissionLevel      string `json:"permission_level"`
}

func (acc AccessControlChangeApiRequest) String() string {
	return fmt.Sprintf("%v%v%v %s", acc.UserName, acc.GroupName, acc.ServicePrincipalName,
		acc.PermissionLevel)
}

// NewPermissionsAPI creates PermissionsAPI instance from provider meta
func NewPermissionsAPI(ctx context.Context, m any) PermissionsAPI {
	return PermissionsAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// PermissionsAPI exposes general permission related methods
type PermissionsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Helper function for applying permissions changes. Ensures that
// we select the correct HTTP method based on the object type and preserve the calling
// user's ability to manage the specified object when applying permissions changes.
func (a PermissionsAPI) put(mapping resourcePermissions, objectID string, objectACL AccessControlChangeListApiRequest) error {
	urlPath := mapping.getRequestPath(objectID)
	if mapping.usePost {
		// SQLA entities use POST for permission updates.
		return a.client.Post(a.context, urlPath, objectACL, nil)
	}
	log.Printf("[DEBUG] PUT %s %v", objectID, objectACL)
	return a.client.Put(a.context, urlPath, objectACL)
}

// safePutWithOwner is a workaround for the limitation where warehouse without owners cannot have IS_OWNER set
func (a PermissionsAPI) safePutWithOwner(objectID string, objectACL AccessControlChangeListApiRequest, getCurrentUser, getOwner func() (string, error)) error {
	mapping, err := getResourcePermissions(objectID)
	if mapping.shouldExplicitlyGrantCallingUserManagePermissions {
		currentUser, err := getCurrentUser()
		if err != nil {
			return err
		}
		objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChangeApiRequest{
			UserName:        currentUser,
			PermissionLevel: "CAN_MANAGE",
		})
	}
	originalAcl := make([]AccessControlChangeApiRequest, len(objectACL.AccessControlList))
	copy(originalAcl, objectACL.AccessControlList)
	if err != nil {
		return err
	}
	if mapping.hasOwnerPermissionLevel() {
		owners := 0
		for _, acl := range objectACL.AccessControlList {
			if acl.PermissionLevel == "IS_OWNER" {
				owners++
			}
		}
		if owners == 0 {
			// add owner if it's missing, otherwise automated planning might be difficult
			owner, err := getOwner()
			if err != nil {
				return err
			}
			if owner == "" {
				return nil
			}
			objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChangeApiRequest{
				UserName:        owner,
				PermissionLevel: "IS_OWNER",
			})
		}
	}
	err = a.put(mapping, objectID, objectACL)
	if err != nil {
		if strings.Contains(err.Error(), "with no existing owner must provide a new owner") {
			objectACL.AccessControlList = originalAcl
			return a.put(mapping, objectID, objectACL)
		}
		return err
	}
	return nil
}

func (a PermissionsAPI) getCurrentUser() (string, error) {
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return "", err
	}
	me, err := w.CurrentUser.Me(a.context)
	if err != nil {
		return "", err
	}
	return me.UserName, nil
}

// Update updates object permissions. Technically, it's using method named SetOrDelete, but here we do more
func (a PermissionsAPI) Update(objectID string, objectACL AccessControlChangeListApiRequest, mapping resourcePermissions) error {
	currentUser, err := a.getCurrentUser()
	if err != nil {
		return err
	}
	// this logic was moved from CustomizeDiff because of undeterministic auth behavior
	// in the corner-case scenarios.
	// see https://github.com/databricks/terraform-provider-databricks/issues/2052
	err = mapping.validate(objectACL.AccessControlList, currentUser)
	if err != nil {
		return err
	}
	if mapping.requiresExplicitAdminPermissions(objectID) {
		// Prevent "Cannot change permissions for group 'admins' to None."
		objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChangeApiRequest{
			GroupName:       "admins",
			PermissionLevel: "CAN_MANAGE",
		})
	}
	getCurrentUser := func() (string, error) { return currentUser, nil }
	return a.safePutWithOwner(objectID, objectACL, getCurrentUser, getCurrentUser)
}

// Delete gracefully removes permissions of non-admin users. After this operation, the object is managed
// by the current user and admin group.
func (a PermissionsAPI) Delete(objectID string, mapping resourcePermissions) error {
	objectACL, err := a.Read(objectID)
	if err != nil {
		return err
	}
	accl := AccessControlChangeListApiRequest{}
	for _, acl := range objectACL.AccessControlList {
		// When deleting permissions for a resource with explicit admin permissions, delete should remove
		// admin permissions as well. Otherwise, admin permissions should be left as-is.
		if acl.GroupName == "admins" && !mapping.allowAdminGroup {
			if change, direct := acl.toAccessControlChange(); direct {
				// keep everything direct for admin group
				accl.AccessControlList = append(accl.AccessControlList, change)
			}
		}
	}
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return err
	}
	return a.safePutWithOwner(objectID, accl, a.getCurrentUser, func() (string, error) { return mapping.getObjectCreator(a.context, w, objectID) })
}

// Read gets all relevant permissions for the object, including inherited ones
func (a PermissionsAPI) Read(objectID string) (objectACL ObjectAclApiResponse, err error) {
	mapping, err := getResourcePermissions(objectID)
	if err != nil {
		return objectACL, err
	}
	err = a.client.Get(a.context, mapping.getRequestPath(objectID), nil, &objectACL)
	var apiErr *apierr.APIError
	// https://github.com/databricks/terraform-provider-databricks/issues/1227
	// platform propagates INVALID_STATE error for auto-purged clusters in
	// the permissions api. this adds "a logical fix" also here, not to introduce
	// cross-package dependency on "clusters".
	if errors.As(err, &apiErr) && strings.Contains(apiErr.Message, "Cannot access cluster") && apiErr.StatusCode == 400 {
		apiErr.StatusCode = 404
		err = apiErr
		return
	}
	if strings.HasPrefix(objectID, "/dashboards/") {
		// workaround for inconsistent API response returning object ID of file in the workspace
		objectACL.ObjectID = objectID
	}
	return
}

// PermissionsEntity is the one used for resource metadata
type PermissionsEntity struct {
	ObjectType        string                          `json:"object_type,omitempty" tf:"computed"`
	AccessControlList []AccessControlChangeApiRequest `json:"access_control" tf:"slice_set"`
}

func (p PermissionsEntity) containsUserOrServicePrincipal(name string) bool {
	for _, ac := range p.AccessControlList {
		if ac.UserName == name || ac.ServicePrincipalName == name {
			return true
		}
	}
	return false
}

// ResourcePermissions definition
func ResourcePermissions() common.Resource {
	s := common.StructToSchema(PermissionsEntity{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		for _, mapping := range permissionsResourceIDFields() {
			s[mapping.field] = &schema.Schema{
				ForceNew: true,
				Type:     schema.TypeString,
				Optional: true,
			}
			for _, m := range permissionsResourceIDFields() {
				if m.field == mapping.field {
					continue
				}
				s[mapping.field].ConflictsWith = append(s[mapping.field].ConflictsWith, m.field)
			}
		}
		s["access_control"].MinItems = 1
		return s
	})
	return common.Resource{
		Schema: s,
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff) error {
			mapping, _, err := getResourcePermissionsFromState(diff)
			if err != nil {
				// This preserves current behavior but is likely only exercised in tests where
				// the original config is not specified.
				return nil
			}
			planned := PermissionsEntity{}
			common.DiffToStructPointer(diff, s, &planned)
			// Plan time validation for object permission levels
			for _, accessControl := range planned.AccessControlList {
				permissionLevel := accessControl.PermissionLevel
				if _, ok := mapping.allowedPermissionLevels[permissionLevel]; !ok {
					return fmt.Errorf(`permission_level %s is not supported with %s objects; allowed levels: %s`,
						permissionLevel, mapping.field, strings.Join(mapping.getAllowedPermissionLevels(true), ", "))
				}
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			id := d.Id()
			a := NewPermissionsAPI(ctx, c)
			objectACL, err := a.Read(id)
			if err != nil {
				return err
			}
			me, err := a.getCurrentUser()
			if err != nil {
				return err
			}
			var existing PermissionsEntity
			common.DataToStructPointer(d, s, &existing)
			entity, err := objectACL.ToPermissionsEntity(d, existing, me)
			if err != nil {
				return err
			}
			if len(entity.AccessControlList) == 0 {
				// empty "modifiable" access control list is the same as resource absence
				d.SetId("")
				return nil
			}
			return common.StructToData(entity, s, d)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var entity PermissionsEntity
			common.DataToStructPointer(d, s, &entity)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			mapping, v, err := getResourcePermissionsFromState(d)
			if err != nil {
				return err
			}
			objectID, err := mapping.getID(ctx, w, v)
			if err != nil {
				return err
			}
			err = NewPermissionsAPI(ctx, c).Update(objectID, AccessControlChangeListApiRequest{
				AccessControlList: entity.AccessControlList,
			}, mapping)
			if err != nil {
				return err
			}
			d.SetId(objectID)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var entity PermissionsEntity
			common.DataToStructPointer(d, s, &entity)
			mapping, err := getResourcePermissions(d.Id())
			if err != nil {
				return err
			}
			return NewPermissionsAPI(ctx, c).Update(d.Id(), AccessControlChangeListApiRequest{
				AccessControlList: entity.AccessControlList,
			}, mapping)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			mapping, err := getResourcePermissions(d.Id())
			if err != nil {
				return err
			}
			return NewPermissionsAPI(ctx, c).Delete(d.Id(), mapping)
		},
	}
}
