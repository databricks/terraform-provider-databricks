package permissions

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func toPermissionsEntity(oa *iam.ObjectPermissions, d *schema.ResourceData, existing PermissionsEntity, me string, mapping resourcePermissions) (PermissionsEntity, error) {
	entity := PermissionsEntity{}
	for _, accessControl := range oa.AccessControlList {
		if me == accessControl.UserName || me == accessControl.ServicePrincipalName {
			// If the user doesn't include an access_control block for themselves, do not include it in the state.
			// On create/update, the provider will automatically include the current user in the access_control block
			// for appropriate resources. Otherwise, it must be included in state to prevent configuration drift.
			if !existing.containsUserOrServicePrincipal(me) {
				continue
			}
		}
		entity.AccessControlList = append(entity.AccessControlList, iam.AccessControlRequest{
			GroupName:            accessControl.GroupName,
			UserName:             accessControl.UserName,
			ServicePrincipalName: accessControl.ServicePrincipalName,
			PermissionLevel:      accessControl.AllPermissions[0].PermissionLevel,
		})
	}
	entity.ObjectType = mapping.objectType
	pathVariant := d.Get(mapping.getPathVariant())
	if pathVariant != nil && pathVariant.(string) != "" {
		// we're not importing and it's a path... it's set, so let's not re-set it
		return entity, nil
	}
	identifier := path.Base(oa.ObjectId)
	return entity, d.Set(mapping.field, identifier)
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

// safePutWithOwner is a workaround for the limitation where warehouse without owners cannot have IS_OWNER set
func (a PermissionsAPI) safePutWithOwner(objectID string, objectACL []iam.AccessControlRequest, mapping resourcePermissions, ownerOpt string) error {
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return err
	}
	idParts := strings.Split(objectID, "/")
	id := idParts[len(idParts)-1]
	withOwner := mapping.addOwnerPermissionIfNeeded(objectACL, ownerOpt)
	_, err = w.Permissions.Set(a.context, iam.PermissionsRequest{
		RequestObjectId:   id,
		RequestObjectType: mapping.resourceType,
		AccessControlList: withOwner,
	})
	if err != nil {
		if strings.Contains(err.Error(), "with no existing owner must provide a new owner") {
			_, err = w.Permissions.Set(a.context, iam.PermissionsRequest{
				RequestObjectId:   objectID,
				RequestObjectType: mapping.resourceType,
				AccessControlList: objectACL,
			})
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
func (a PermissionsAPI) Update(objectID string, objectACL []iam.AccessControlRequest, mapping resourcePermissions) error {
	currentUser, err := a.getCurrentUser()
	if err != nil {
		return err
	}
	// this logic was moved from CustomizeDiff because of undeterministic auth behavior
	// in the corner-case scenarios.
	// see https://github.com/databricks/terraform-provider-databricks/issues/2052
	err = mapping.validate(objectACL, currentUser)
	if err != nil {
		return err
	}
	accl, err := mapping.prepareForUpdate(objectID, objectACL, currentUser)
	if err != nil {
		return err
	}
	return a.safePutWithOwner(objectID, accl, mapping, currentUser)
}

// Delete gracefully removes permissions of non-admin users. After this operation, the object is managed
// by the current user and admin group. If the resource has IS_OWNER permissions, they are reset to the
// object creator, if it can be determined.
func (a PermissionsAPI) Delete(objectID string, mapping resourcePermissions) error {
	objectACL, err := a.Read(objectID, mapping)
	if err != nil {
		return err
	}
	accl, err := mapping.prepareForDelete(objectACL, a.getCurrentUser)
	if err != nil {
		return err
	}
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return err
	}
	resourceStatus, err := mapping.getObjectStatus(a.context, w, objectID)
	if err != nil {
		return err
	}
	// Do not bother resetting permissions for deleted resources
	if !resourceStatus.exists {
		return nil
	}
	return a.safePutWithOwner(objectID, accl, mapping, resourceStatus.creator)
}

// Read gets all relevant permissions for the object, including inherited ones
func (a PermissionsAPI) Read(objectID string, mapping resourcePermissions) (objectACL *iam.ObjectPermissions, err error) {
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return objectACL, err
	}
	idParts := strings.Split(objectID, "/")
	id := idParts[len(idParts)-1]
	permissions, err := w.Permissions.Get(a.context, iam.GetPermissionRequest{
		RequestObjectId:   id,
		RequestObjectType: mapping.resourceType,
	})
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
	if err != nil {
		return objectACL, err
	}
	return mapping.prepareResponse(objectID, permissions)
}

// PermissionsEntity is the one used for resource metadata
type PermissionsEntity struct {
	ObjectType        string                     `json:"object_type,omitempty" tf:"computed"`
	AccessControlList []iam.AccessControlRequest `json:"access_control" tf:"slice_set"`
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
		for _, mapping := range allResourcePermissions() {
			s[mapping.field] = &schema.Schema{
				ForceNew: true,
				Type:     schema.TypeString,
				Optional: true,
			}
			for _, m := range allResourcePermissions() {
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
				// No diff in permission level, so don't need to check.
				if permissionLevel == "" {
					continue
				}
				if _, ok := mapping.allowedPermissionLevels[string(permissionLevel)]; !ok {
					return fmt.Errorf(`permission_level %s is not supported with %s objects; allowed levels: %s`,
						permissionLevel, mapping.field, strings.Join(mapping.getAllowedPermissionLevels(true), ", "))
				}
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			id := d.Id()
			a := NewPermissionsAPI(ctx, c)
			mapping, err := getResourcePermissions(id)
			if err != nil {
				return err
			}
			objectACL, err := a.Read(id, mapping)
			if err != nil {
				return err
			}
			me, err := a.getCurrentUser()
			if err != nil {
				return err
			}
			var existing PermissionsEntity
			common.DataToStructPointer(d, s, &existing)
			entity, err := toPermissionsEntity(objectACL, d, existing, me, mapping)
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
			mapping, id, err := getResourcePermissionsFromState(d)
			if err != nil {
				return err
			}
			objectID, err := mapping.getID(ctx, w, id)
			if err != nil {
				return err
			}
			err = NewPermissionsAPI(ctx, c).Update(objectID, entity.AccessControlList, mapping)
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
			return NewPermissionsAPI(ctx, c).Update(d.Id(), entity.AccessControlList, mapping)
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
