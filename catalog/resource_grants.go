package catalog

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/catalog/permissions"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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

// diffPermissions returns an array of catalog.PermissionsChange of this permissions list with `diff` privileges removed
func diffPermissions(pl catalog.PermissionsList, existing catalog.PermissionsList) (diff []catalog.PermissionsChange) {
	// diffs change sets
	configured := map[string]*schema.Set{}
	for _, v := range pl.PrivilegeAssignments {
		configured[v.Principal] = permissions.SliceToSet(v.Privileges)
	}
	// existing permissions that needs removal
	remote := map[string]*schema.Set{}
	for _, v := range existing.PrivilegeAssignments {
		remote[v.Principal] = permissions.SliceToSet(v.Privileges)
	}
	// STEP 1: detect overlaps
	for principal, confPrivs := range configured {
		remotePrivs, ok := remote[principal]
		if !ok {
			remotePrivs = permissions.SliceToSet([]catalog.Privilege{})
		}
		add := permissions.SetToSlice(confPrivs.Difference(remotePrivs))
		remove := permissions.SetToSlice(remotePrivs.Difference(confPrivs))
		if len(add) == 0 && len(remove) == 0 {
			continue
		}
		diff = append(diff, catalog.PermissionsChange{
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
		diff = append(diff, catalog.PermissionsChange{
			Principal: principal,
			Remove:    permissions.SetToSlice(remove),
		})
	}
	// so that we can deterministic tests
	sort.Slice(diff, func(i, j int) bool {
		return diff[i].Principal < diff[j].Principal
	})
	return diff
}

// replaceAllPermissions merges removal diff of existing permissions on the platform
func replaceAllPermissions(a permissions.UnityCatalogPermissionsAPI, securable string, name string, list catalog.PermissionsList) error {
	securableType := permissions.Mappings.GetSecurableType(securable)
	existing, err := a.GetPermissions(securableType, name)
	if err != nil {
		return err
	}
	err = a.UpdatePermissions(securableType, name, diffPermissions(list, *existing))
	if err != nil {
		return err
	}
	return a.WaitForUpdate(1*time.Minute, securableType, name, list, func(current *catalog.PermissionsList, desired catalog.PermissionsList) []catalog.PermissionsChange {
		return diffPermissions(desired, *current)
	})
}

type securableMapping map[string]map[string]bool

// reuse ResourceDiff and ResourceData
type attributeGetter interface {
	GetRawConfig() cty.Value
}

func (sm securableMapping) kv(d attributeGetter) (string, string) {
	rawConfig := d.GetRawConfig()
	if rawConfig.IsNull() {
		return "unknown", "unknown"
	}
	rawConfigValues := rawConfig.AsValueMap()
	for field := range sm {
		value := rawConfigValues[field]
		if value.IsNull() {
			continue
		}
		// We need this check because:
		// "AsString returns the native string from a non-null, non-unknown cty.String value, or panics if called on any other value."
		if !value.IsKnown() {
			// This is the case when the new value isn't known but the field is required for validation
			return field, "unknown"
		}
		return field, value.AsString()
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
				// check if user uses spaces instead of underscores
				if allowed[strings.ReplaceAll(priv, " ", "_")] {
					return fmt.Errorf(`%s is not allowed on %s. Did you mean %s?`, priv, securable, strings.ReplaceAll(priv, " ", "_"))
				}
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
		"APPLY_TAG":      true,
		"BROWSE":         true,
	},
	"catalog": {
		"CREATE": true,
		"USAGE":  true,

		// v1.0
		"ALL_PRIVILEGES":           true,
		"APPLY_TAG":                true,
		"USE_CATALOG":              true,
		"USE_SCHEMA":               true,
		"CREATE_SCHEMA":            true,
		"CREATE_TABLE":             true,
		"CREATE_FUNCTION":          true,
		"CREATE_MATERIALIZED_VIEW": true,
		"CREATE_MODEL":             true,
		"CREATE_VOLUME":            true,
		"READ_VOLUME":              true,
		"WRITE_VOLUME":             true,
		"EXECUTE":                  true,
		"MODIFY":                   true,
		"SELECT":                   true,
		"REFRESH":                  true,
		"BROWSE":                   true,
	},
	"schema": {
		"CREATE": true,
		"USAGE":  true,

		// v1.0
		"ALL_PRIVILEGES":           true,
		"APPLY_TAG":                true,
		"USE_SCHEMA":               true,
		"CREATE_TABLE":             true,
		"CREATE_FUNCTION":          true,
		"CREATE_MATERIALIZED_VIEW": true,
		"CREATE_MODEL":             true,
		"CREATE_VOLUME":            true,
		"READ_VOLUME":              true,
		"WRITE_VOLUME":             true,
		"EXECUTE":                  true,
		"MODIFY":                   true,
		"SELECT":                   true,
		"REFRESH":                  true,
		"BROWSE":                   true,
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
		"CREATE_MANAGED_STORAGE": true,
		"CREATE_EXTERNAL_VOLUME": true,
		"BROWSE":                 true,
	},
	"metastore": {
		// v1.0
		"CREATE_CATALOG":            true,
		"CREATE_CLEAN_ROOM":         true,
		"CREATE_CONNECTION":         true,
		"CREATE_EXTERNAL_LOCATION":  true,
		"CREATE_STORAGE_CREDENTIAL": true,
		"CREATE_SHARE":              true,
		"CREATE_RECIPIENT":          true,
		"CREATE_PROVIDER":           true,
		"MANAGE_ALLOWLIST":          true,
		"USE_CONNECTION":            true,
		"USE_PROVIDER":              true,
		"USE_SHARE":                 true,
		"USE_RECIPIENT":             true,
		"USE_MARKETPLACE_ASSETS":    true,
		"SET_SHARE_PERMISSION":      true,
	},
	"function": {
		"ALL_PRIVILEGES": true,
		"EXECUTE":        true,
	},
	"model": {
		"ALL_PRIVILEGES": true,
		"APPLY_TAG":      true,
		"EXECUTE":        true,
	},
	"share": {
		"SELECT": true,
	},
	"volume": {
		"ALL_PRIVILEGES": true,
		"READ_VOLUME":    true,
		"WRITE_VOLUME":   true,
	},
	// avoid reserved field
	"foreign_connection": {
		"ALL_PRIVILEGES":         true,
		"CREATE_FOREIGN_CATALOG": true,
		"CREATE_FOREIGN_SCHEMA":  true,
		"CREATE_FOREIGN_TABLE":   true,
		"USE_CONNECTION":         true,
	},
}

func (pl PermissionsList) toSdkPermissionsList() (out catalog.PermissionsList) {
	for _, v := range pl.Assignments {
		privileges := []catalog.Privilege{}
		for _, p := range v.Privileges {
			privileges = append(privileges, catalog.Privilege(p))
		}
		out.PrivilegeAssignments = append(out.PrivilegeAssignments, catalog.PrivilegeAssignment{
			Principal:  v.Principal,
			Privileges: privileges,
		})
	}
	return
}

func sdkPermissionsListToPermissionsList(sdkPermissionsList catalog.PermissionsList) (out PermissionsList) {
	for _, v := range sdkPermissionsList.PrivilegeAssignments {
		privileges := []string{}
		for _, p := range v.Privileges {
			privileges = append(privileges, p.String())
		}
		out.Assignments = append(out.Assignments, PrivilegeAssignment{
			Principal:  v.Principal,
			Privileges: privileges,
		})
	}
	return
}

func parseId(d *schema.ResourceData) (string, string, error) {
	split := strings.SplitN(d.Id(), "/", 2)
	if len(split) != 2 {
		return "", "", fmt.Errorf("ID must be two elements split by `/`: %s", d.Id())
	}
	return split[0], split[1], nil
}

func ResourceGrants() common.Resource {
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
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			if d.Id() == "" {
				// unfortunately we cannot do validation before dependent resources exist with tfsdkv2
				return nil
			}
			var grants PermissionsList
			common.DiffToStructPointer(d, s, &grants)
			return mapping.validate(d, grants)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore").(string))
			if err != nil {
				return err
			}
			var grants PermissionsList
			common.DataToStructPointer(d, s, &grants)
			securable, name := mapping.kv(d)
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			err = replaceAllPermissions(unityCatalogPermissionsAPI, securable, name, grants.toSdkPermissionsList())
			if err != nil {
				return err
			}
			d.SetId(mapping.id(d))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			securable, name, err := parseId(d)
			if err != nil {
				return err
			}
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			grants, err := unityCatalogPermissionsAPI.GetPermissions(permissions.Mappings.GetSecurableType(securable), name)
			if err != nil {
				return err
			}
			if len(grants.PrivilegeAssignments) == 0 {
				return apierr.NotFound("got empty permissions list")
			}

			err = common.StructToData(sdkPermissionsListToPermissionsList(*grants), s, d)
			if err != nil {
				return err
			}
			return d.Set(securable, name)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore").(string))
			if err != nil {
				return err
			}
			securable, name, err := parseId(d)
			if err != nil {
				return err
			}
			var grants PermissionsList
			common.DataToStructPointer(d, s, &grants)
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			return replaceAllPermissions(unityCatalogPermissionsAPI, securable, name, grants.toSdkPermissionsList())
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore").(string))
			if err != nil {
				return err
			}
			securable, name, err := parseId(d)
			if err != nil {
				return err
			}
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			return replaceAllPermissions(unityCatalogPermissionsAPI, securable, name, catalog.PermissionsList{})
		},
	}
}
