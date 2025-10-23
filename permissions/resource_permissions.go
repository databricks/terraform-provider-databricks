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
	"github.com/databricks/terraform-provider-databricks/permissions/entity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
	_, err = w.Permissions.Set(a.context, iam.SetObjectPermissions{
		RequestObjectId:   id,
		RequestObjectType: mapping.requestObjectType,
		AccessControlList: withOwner,
	})
	if err != nil {
		if strings.Contains(err.Error(), "with no existing owner must provide a new owner") {
			_, err = w.Permissions.Set(a.context, iam.SetObjectPermissions{
				RequestObjectId:   id,
				RequestObjectType: mapping.requestObjectType,
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
func (a PermissionsAPI) Update(objectID string, entity entity.PermissionsEntity, mapping resourcePermissions) error {
	currentUser, err := a.getCurrentUser()
	if err != nil {
		return err
	}
	// this logic was moved from CustomizeDiff because of undeterministic auth behavior
	// in the corner-case scenarios.
	// see https://github.com/databricks/terraform-provider-databricks/issues/2052
	err = mapping.validate(a.context, entity, currentUser)
	if err != nil {
		return err
	}
	prepared, err := mapping.prepareForUpdate(objectID, entity, currentUser)
	if err != nil {
		return err
	}
	return a.safePutWithOwner(objectID, prepared.AccessControlList, mapping, currentUser)
}

// Delete gracefully removes permissions of non-admin users. After this operation, the object is managed
// by the current user and admin group. If the resource has IS_OWNER permissions, they are reset to the
// object creator, if it can be determined.
func (a PermissionsAPI) Delete(objectID string, mapping resourcePermissions) error {
	if mapping.objectType == "pipelines" {
		// There is a bug which causes the code below send IS_OWNER with run_as identity
		// Which is of course wrong thing to do.
		// For non-admin users this results in the error: https://community.databricks.com/t5/data-engineering/dab-dlt-destroy-fails-due-to-ownership-permissions-mismatch/td-p/132101
		// For admin users situation is worse but there is no error, it silently changes owner to wrong identity.
		return nil
	}
	objectACL, err := a.readRaw(objectID, mapping)
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

func (a PermissionsAPI) readRaw(objectID string, mapping resourcePermissions) (*iam.ObjectPermissions, error) {
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return nil, err
	}
	idParts := strings.Split(objectID, "/")
	id := idParts[len(idParts)-1]

	// TODO: This a temporary measure to implement retry on 504 until this is
	// supported natively in the Go SDK.
	permissions, err := common.RetryOn504(a.context, func(ctx context.Context) (*iam.ObjectPermissions, error) {
		return w.Permissions.Get(a.context, iam.GetPermissionRequest{
			RequestObjectId:   id,
			RequestObjectType: mapping.requestObjectType,
		})
	})

	var apiErr *apierr.APIError
	// https://github.com/databricks/terraform-provider-databricks/issues/1227
	// platform propagates INVALID_STATE error for auto-purged clusters in
	// the permissions api. this adds "a logical fix" also here, not to introduce
	// cross-package dependency on "clusters".
	if errors.As(err, &apiErr) && strings.Contains(apiErr.Message, "Cannot access cluster") && apiErr.StatusCode == 400 {
		apiErr.StatusCode = 404
		apiErr.ErrorCode = "RESOURCE_DOES_NOT_EXIST"
		err = apiErr
	}
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// Read gets all relevant permissions for the object, including inherited ones
func (a PermissionsAPI) Read(objectID string, mapping resourcePermissions, existing entity.PermissionsEntity, me string) (entity.PermissionsEntity, error) {
	permissions, err := a.readRaw(objectID, mapping)
	if err != nil {
		return entity.PermissionsEntity{}, err
	}
	return mapping.prepareResponse(objectID, permissions, existing, me)
}

// ResourcePermissions definition
func ResourcePermissions() common.Resource {
	s := common.StructToSchema(entity.PermissionsEntity{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
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

		// Use a custom hash function that only considers non-empty fields.
		// This prevents spurious diffs when comparing {group_name: "X", permission_level: "Y"}
		// with {group_name: "X", permission_level: "Y", service_principal_name: "", user_name: ""}
		acSchema := s["access_control"].Elem.(*schema.Resource).Schema
		s["access_control"].Set = func(v interface{}) int {
			m, ok := v.(map[string]interface{})
			if !ok {
				return 0
			}
			// Build a normalized map with only non-empty string fields for hashing
			normalized := make(map[string]interface{})
			for key, val := range m {
				if _, exists := acSchema[key]; !exists {
					continue // Skip fields not in schema
				}
				if strVal, ok := val.(string); ok {
					if strVal != "" {
						normalized[key] = strVal
					}
				}
			}
			// Use HashResource with a schema that only includes the fields we care about
			hashSchema := &schema.Resource{Schema: acSchema}
			return schema.HashResource(hashSchema)(normalized)
		}
		return s
	})
	return common.Resource{
		Schema: s,
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, c *common.DatabricksClient) error {
			mapping, _, err := getResourcePermissionsFromState(diff)
			if err != nil {
				// This preserves current behavior but is likely only exercised in tests where
				// the original config is not specified.
				return nil
			}
			planned := entity.PermissionsEntity{}
			common.DiffToStructPointer(diff, s, &planned)
			// Plan time validation for object permission levels
			for _, accessControl := range planned.AccessControlList {
				permissionLevel := accessControl.PermissionLevel
				// No diff in permission level, so don't need to check.
				if permissionLevel == "" {
					continue
				}
				// TODO: only warn on unknown permission levels, as new levels may be released that the TF provider
				// is not aware of.
				if _, ok := mapping.allowedPermissionLevels[string(permissionLevel)]; !ok {
					return fmt.Errorf(`permission_level %s is not supported with %s objects; allowed levels: %s`,
						permissionLevel, mapping.field, strings.Join(mapping.getAllowedPermissionLevels(true), ", "))
				}
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			a := NewPermissionsAPI(ctx, c)
			mapping, err := getResourcePermissionsFromId(d.Id())
			if err != nil {
				return err
			}
			var existing entity.PermissionsEntity
			common.DataToStructPointer(d, s, &existing)
			me, err := a.getCurrentUser()
			if err != nil {
				return err
			}
			id := d.Id()
			entity, err := a.Read(id, mapping, existing, me)
			if err != nil {
				return err
			}
			if len(entity.AccessControlList) == 0 {
				// empty "modifiable" access control list is the same as resource absence
				d.SetId("")
				return nil
			}
			entity.ObjectType = mapping.objectType
			pathVariant := d.Get(mapping.getPathVariant())
			if pathVariant == nil || pathVariant.(string) == "" {
				identifier := path.Base(id)
				if err = d.Set(mapping.field, identifier); err != nil {
					return err
				}
			}
			return common.StructToData(entity, s, d)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var entity entity.PermissionsEntity
			common.DataToStructPointer(d, s, &entity)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			mapping, configuredValue, err := getResourcePermissionsFromState(d)
			if err != nil {
				return err
			}
			objectID, err := mapping.getID(ctx, w, configuredValue)
			if err != nil {
				return err
			}
			err = NewPermissionsAPI(ctx, c).Update(objectID, entity, mapping)
			if err != nil {
				return err
			}
			d.SetId(objectID)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var entity entity.PermissionsEntity
			common.DataToStructPointer(d, s, &entity)
			mapping, err := getResourcePermissionsFromId(d.Id())
			if err != nil {
				return err
			}
			return NewPermissionsAPI(ctx, c).Update(d.Id(), entity, mapping)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			mapping, err := getResourcePermissionsFromId(d.Id())
			if err != nil {
				return err
			}
			return NewPermissionsAPI(ctx, c).Delete(d.Id(), mapping)
		},
	}
}
