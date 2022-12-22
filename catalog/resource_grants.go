package catalog

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
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
	configured := map[string]*schema.Set{}
	for _, v := range pl.Assignments {
		configured[v.Principal] = newStringSet(v.Privileges)
	}
	// existing permissions that needs removal
	remote := map[string]*schema.Set{}
	for _, v := range existing.Assignments {
		remote[v.Principal] = newStringSet(v.Privileges)
	}
	// STEP 1: detect overlaps
	for principal, confPrivs := range configured {
		remotePrivs, ok := remote[principal]
		if !ok {
			remotePrivs = newStringSet([]string{})
		}
		add := setToStrings(confPrivs.Difference(remotePrivs))
		remove := setToStrings(remotePrivs.Difference(confPrivs))
		if len(add) == 0 && len(remove) == 0 {
			continue
		}
		diff.Changes = append(diff.Changes, permissionsChange{
			Principal: principal,
			Add:       add,
			Remove:    remove,
		})
	}
	// STEP 2: non overlap - simply remove
	for principal, remove := range remote {
		_, ok := configured[principal]
		if ok { // already handled in STEP 1
			continue
		}
		diff.Changes = append(diff.Changes, permissionsChange{
			Principal: principal,
			Remove:    setToStrings(remove),
		})
	}
	// so that we can deterministic tests
	sort.Slice(diff.Changes, func(i, j int) bool {
		return diff.Changes[i].Principal < diff.Changes[j].Principal
	})
	return diff
}

func newStringSet(in []string) *schema.Set {
	var out []any
	for _, v := range in {
		out = append(out, v)
	}
	return schema.NewSet(schema.HashString, out)
}

func NewPermissionsAPI(ctx context.Context, m any) PermissionsAPI {
	return PermissionsAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

func getPermissionEndpoint(securable, name string) string {
	if securable == "share" {
		return fmt.Sprintf("/unity-catalog/shares/%s/permissions", name)
	}
	return fmt.Sprintf("/unity-catalog/permissions/%s/%s", securable, name)
}

func (a PermissionsAPI) getPermissions(securable, name string) (list PermissionsList, err error) {
	err = a.client.Get(a.context, getPermissionEndpoint(securable, name), nil, &list)
	return
}

func (a PermissionsAPI) updatePermissions(securable, name string, diff permissionsDiff) error {
	return a.client.Patch(a.context, getPermissionEndpoint(securable, name), diff)
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
	Get(key string) any
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
		"MODIFY": true,
		"SELECT": true,

		// v1.0
		"ALL_PRIVILEGES": true,
	},
	"view": {
		"SELECT": true,
	},
	"catalog": {
		"CREATE": true,
		"USAGE":  true,

		// v1.0
		"ALL_PRIVILEGES":           true,
		"USE_CATALOG":              true,
		"USE_SCHEMA":               true,
		"CREATE_SCHEMA":            true,
		"CREATE_TABLE":             true,
		"CREATE_VIEW":              true,
		"CREATE_FUNCTION":          true,
		"CREATE_MATERIALIZED_VIEW": true,
		"EXECUTE":                  true,
		"MODIFY":                   true,
		"SELECT":                   true,
		"REFRESH":                  true,
	},
	"schema": {
		"CREATE": true,
		"USAGE":  true,

		// v1.0
		"ALL_PRIVILEGES":           true,
		"USE_SCHEMA":               true,
		"CREATE_TABLE":             true,
		"CREATE_VIEW":              true,
		"CREATE_FUNCTION":          true,
		"CREATE_MATERIALIZED_VIEW": true,
		"EXECUTE":                  true,
		"MODIFY":                   true,
		"SELECT":                   true,
		"REFRESH":                  true,
	},
	"storage_credential": {
		"CREATE_TABLE":             true,
		"READ_FILES":               true,
		"WRITE_FILES":              true,
		"CREATE_EXTERNAL_LOCATION": true,

		// v1.0
		"ALL_PRIVILEGES":        true,
		"CREATE_EXTERNAL_TABLE": true,
	},
	"external_location": {
		"CREATE_TABLE": true,
		"READ_FILES":   true,
		"WRITE_FILES":  true,

		// v1.0
		"ALL_PRIVILEGES":         true,
		"CREATE_EXTERNAL_TABLE":  true,
		"CREATE MANAGED STORAGE": true,
	},
	"metastore": {
		// v1.0
		"CREATE_CATALOG":            true,
		"CREATE_EXTERNAL_LOCATION":  true,
		"CREATE_STORAGE_CREDENTIAL": true,
		"CREATE_SHARE":              true,
		"CREATE_RECIPIENT":          true,
		"CREATE_PROVIDER":           true,
	},
	"function": {
		"ALL_PRIVILEGES": true,
		"EXECUTE":        true,
	},
	"materialized_view": {
		"ALL_PRIVILEGES": true,
		"SELECT":         true,
		"REFRESH":        true,
	},
	"share": {
		"SELECT": true,
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
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c any) error {
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
			if len(grants.Assignments) == 0 {
				return common.NotFound("got empty permissions list")
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
