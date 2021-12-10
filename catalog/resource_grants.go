package catalog

import (
	"context"
	"fmt"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type PermissionsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// PrivilegeAssignment reflects on `grant` block
type PrivilegeAssignment struct {
	Principal  string   `json:"principal"`
	Privileges []string `json:"privileges" tf:"slice_set"`
}

// PermissionsList reflects it's shape on terraform resource with
// privilege_assignments column renamed to `grant` block for simplicity
type PermissionsList struct {
	Assignments []PrivilegeAssignment `json:"privilege_assignments" tf:"slice_set,alias:grant"`
}

func (pl PermissionsList) toCreate() (diff PermissionsDiff) {
	for _, v := range pl.Assignments {
		diff.Changes = append(diff.Changes, PermissionsChange{
			Principal: v.Principal,
			Add:       v.Privileges,
		})
	}
	return
}

func (pl PermissionsList) toDelete() (diff PermissionsDiff) {
	for _, v := range pl.Assignments {
		diff.Changes = append(diff.Changes, PermissionsChange{
			Principal: v.Principal,
			Remove:    v.Privileges,
		})
	}
	return
}

type PermissionsDiff struct {
	Changes []PermissionsChange `json:"changes"`
}

type PermissionsChange struct {
	Principal string   `json:"principal"`
	Add       []string `json:"add,omitempty"`
	Remove    []string `json:"remove,omitempty"`
}

func NewPermissionsAPI(ctx context.Context, m interface{}) PermissionsAPI {
	return PermissionsAPI{m.(*common.DatabricksClient), ctx}
}

func (a PermissionsAPI) getPermissions(securable, name string) (list PermissionsList, err error) {
	err = a.client.Get(a.context, fmt.Sprintf("/unity-catalog/permissions/%s/%s", securable, name), nil, &list)
	return
}

func (a PermissionsAPI) updatePermissions(securable, name string, diff PermissionsDiff) error {
	return a.client.Patch(a.context, fmt.Sprintf("/unity-catalog/permissions/%s/%s", securable, name), diff)
}

type securableMapping map[string]map[string]bool

// reuse ResourceDiff and ResourceData
type attributeGetter interface {
	Get(key string) interface{}
}

func (sm securableMapping) kv(d attributeGetter) (string, string) {
	for field := range sm {
		v := d.Get(field).(string)
		if v == "" {
			continue
		}
		return field, v
	}
	return "unknown", "unknown"
}

func (sm securableMapping) id(d *schema.ResourceData) string {
	securable, name := sm.kv(d)
	return fmt.Sprintf("%s/%s", securable, name)
}

func (sm securableMapping) validate(d attributeGetter, pl PermissionsList) error {
	securable, _ := sm.kv(d)
	allowed, ok := sm[securable]
	if !ok {
		return fmt.Errorf(`%s is not fully supported yet`, securable)
	}
	for _, v := range pl.Assignments {
		for _, priv := range v.Privileges {
			if !allowed[strings.ToUpper(priv)] {
				return fmt.Errorf(`%s is not allowed on %s`, priv, securable)
			}
		}
	}
	return nil
}

var mapping = securableMapping{
	// add other securable mappings once needed
	"table": {
		"MODIFY":         true,
		"SELECT":         true,
		"ALL PRIVILEGES": true,
	},
	"view": {
		"SELECT":         true,
		"ALL PRIVILEGES": true,
	},
	"catalog": {
		"CREATE":         true,
		"USAGE":          true,
		"ALL PRIVILEGES": true,
	},
	"schema": {
		"CREATE":         true,
		"USAGE":          true,
		"ALL PRIVILEGES": true,
	},
}

func setToStrings(set *schema.Set) (ss []string) {
	for _, v := range set.List() {
		ss = append(ss, v.(string))
	}
	return
}

func principalAndPrivsFromRaw(v interface{}) (string, *schema.Set) {
	item := v.(map[string]interface{})
	principal := item["principal"].(string)
	privileges := item["privileges"].(*schema.Set)
	return principal, privileges
}

func permissionDiffFromRaw(old, new interface{}) PermissionsDiff {
	o := old.(*schema.Set)
	n := new.(*schema.Set)
	prev := map[string]*schema.Set{}
	for _, v := range o.List() {
		principal, privileges := principalAndPrivsFromRaw(v)
		prev[principal] = privileges
	}
	diff := PermissionsDiff{Changes: []PermissionsChange{}}
	mix := map[string]bool{}
	for _, v := range n.List() {
		principal, newPrivs := principalAndPrivsFromRaw(v)
		oldPrivs, exist := prev[principal]
		if !exist {
			// add new principal privileges
			diff.Changes = append(diff.Changes, PermissionsChange{
				Principal: principal,
				Add:       setToStrings(newPrivs),
			})
			continue
		}
		// add or remove principal privileges
		change := PermissionsChange{
			Principal: principal,
			Remove:    setToStrings(oldPrivs.Difference(newPrivs)),
			Add:       setToStrings(newPrivs.Difference(oldPrivs)),
		}
		if len(change.Add) == 0 && len(change.Remove) == 0 {
			continue
		}
		diff.Changes = append(diff.Changes, change)
		mix[principal] = true
	}
	for _, v := range o.Difference(n).List() {
		// remove old principal privs
		principal, oldPrivs := principalAndPrivsFromRaw(v)
		if mix[principal] {
			// skip if mixed removes/adds were detected
			continue
		}
		diff.Changes = append(diff.Changes, PermissionsChange{
			Principal: principal,
			Remove:    setToStrings(oldPrivs),
		})
	}
	return diff
}

func ResourceGrants() *schema.Resource {
	s := common.StructToSchema(PermissionsList{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			alof := []string{}
			for field := range mapping {
				s[field] = &schema.Schema{
					Type:     schema.TypeString,
					ForceNew: true,
					Optional: true,
				}
				alof = append(alof, field)
			}
			for field := range mapping {
				s[field].AtLeastOneOf = alof
			}
			return s
		})
	return common.Resource{
		Schema: s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c interface{}) error {
			var grants PermissionsList
			if err := common.DiffToStructPointer(d, s, &grants); err != nil {
				return err
			}
			return mapping.validate(d, grants)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var grants PermissionsList
			if err := common.DataToStructPointer(d, s, &grants); err != nil {
				return err
			}
			securable, name := mapping.kv(d)
			err := NewPermissionsAPI(ctx, c).updatePermissions(securable, name, grants.toCreate())
			if err != nil {
				return err
			}
			d.SetId(mapping.id(d))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			split := strings.SplitN(d.Id(), "/", 2)
			if len(split) != 2 {
				return fmt.Errorf("ID must be two elements split by `/`: %s", d.Id())
			}
			grants, err := NewPermissionsAPI(ctx, c).getPermissions(split[0], split[1])
			if err != nil {
				return err
			}
			return common.StructToData(grants, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			securable, name := mapping.kv(d)
			diff := permissionDiffFromRaw(d.GetChange("grant"))
			return NewPermissionsAPI(ctx, c).updatePermissions(securable, name, diff)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var grants PermissionsList
			if err := common.DataToStructPointer(d, s, &grants); err != nil {
				return err
			}
			securable, name := mapping.kv(d)
			return NewPermissionsAPI(ctx, c).updatePermissions(securable, name, grants.toDelete())
		},
	}.ToResource()
}
