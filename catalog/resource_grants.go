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

// permissionsDiff is the inner structure of updatePermissions RPC
type permissionsDiff struct {
	Changes []permissionsChange `json:"changes"`
}

type permissionsChange struct {
	Principal string   `json:"principal"`
	Add       []string `json:"add,omitempty"`
	Remove    []string `json:"remove,omitempty"`
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

// diff returns permissionsDiff of this permissions list with `diff` privileges removed
func (pl PermissionsList) diff(existing PermissionsList) (diff permissionsDiff) {
	// diffs change sets
	new := map[string]*schema.Set{}
	for _, v := range pl.Assignments {
		new[v.Principal] = newStringSet(v.Privileges)
	}
	// existing permissions that needs removal
	old := map[string]*schema.Set{}
	for _, v := range existing.Assignments {
		old[v.Principal] = newStringSet(v.Privileges)
	}
	// STEP 1: detect overlaps
	for principal, add := range new {
		remove, ok := old[principal]
		if ok {
			add = add.Difference(remove)
		} else {
			remove = newStringSet([]string{})
		}
		diff.Changes = append(diff.Changes, permissionsChange{
			Principal: principal,
			Add:       setToStrings(add),
			Remove:    setToStrings(remove),
		})
	}
	// STEP 2: non overlap - simply remove
	for principal, remove := range old {
		_, ok := new[principal]
		if ok { // already handled in STEP 1
			continue
		}
		diff.Changes = append(diff.Changes, permissionsChange{
			Principal: principal,
			Remove:    setToStrings(remove),
		})
	}
	return diff
}

func newStringSet(in []string) *schema.Set {
	var out []interface{}
	for _, v := range in {
		out = append(out, v)
	}
	return schema.NewSet(schema.HashString, out)
}

func NewPermissionsAPI(ctx context.Context, m interface{}) PermissionsAPI {
	return PermissionsAPI{m.(*common.DatabricksClient), ctx}
}

func (a PermissionsAPI) getPermissions(securable, name string) (list PermissionsList, err error) {
	err = a.client.Get(a.context, fmt.Sprintf("/unity-catalog/permissions/%s/%s", securable, name), nil, &list)
	return
}

func (a PermissionsAPI) updatePermissions(securable, name string, diff permissionsDiff) error {
	return a.client.Patch(a.context, fmt.Sprintf("/unity-catalog/permissions/%s/%s", securable, name), diff)
}

// replacePermissions merges removal diff of existing permissions on the platform
func (a PermissionsAPI) replacePermissions(securable, name string, list PermissionsList) error {
	existing, err := a.getPermissions(securable, name)
	if err != nil {
		return err
	}
	return a.updatePermissions(securable, name, list.diff(existing))
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
	"storage_credential": {
		"CREATE TABLE":   true,
		"READ FILES":     true,
		"WRITE FILES":    true,
		"ALL PRIVILEGES": true,
	},
	"external_location": {
		"CREATE TABLE":   true,
		"READ FILES":     true,
		"WRITE FILES":    true,
		"ALL PRIVILEGES": true,
	},
}

func setToStrings(set *schema.Set) (ss []string) {
	for _, v := range set.List() {
		ss = append(ss, v.(string))
	}
	return
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
			if d.Id() == "" {
				// unfortunately we cannot do validation before dependent resources exist with tfsdkv2
				return nil
			}
			var grants PermissionsList
			common.DiffToStructPointer(d, s, &grants)
			return mapping.validate(d, grants)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var grants PermissionsList
			common.DataToStructPointer(d, s, &grants)
			securable, name := mapping.kv(d)
			err := NewPermissionsAPI(ctx, c).replacePermissions(securable, name, grants)
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
			var grants PermissionsList
			common.DataToStructPointer(d, s, &grants)
			return NewPermissionsAPI(ctx, c).replacePermissions(securable, name, grants)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			split := strings.SplitN(d.Id(), "/", 2)
			if len(split) != 2 {
				return fmt.Errorf("ID must be two elements split by `/`: %s", d.Id())
			}
			return NewPermissionsAPI(ctx, c).replacePermissions(split[0], split[1], PermissionsList{})
		},
	}.ToResource()
}
