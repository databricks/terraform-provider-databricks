package permissions

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// API
func NewUnityCatalogPermissionsAPI(ctx context.Context, m any) UnityCatalogPermissionsAPI {
	return UnityCatalogPermissionsAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

func (a UnityCatalogPermissionsAPI) GetPermissions(securable, name string) (list UnityCatalogPermissionsList, err error) {
	err = a.client.Get(a.context, getPermissionEndpoint(securable, name), nil, &list)
	return
}

func (a UnityCatalogPermissionsAPI) UpdatePermissions(securable, name string, diff UnityCatalogPermissionsDiff) error {
	return a.client.Patch(a.context, getPermissionEndpoint(securable, name), diff)
}

func getPermissionEndpoint(securable, name string) string {
	if securable == "share" {
		return fmt.Sprintf("/unity-catalog/shares/%s/permissions", name)
	}
	if securable == "foreign_connection" {
		return fmt.Sprintf("/unity-catalog/permissions/connection/%s", name)
	}
	if securable == "model" {
		return fmt.Sprintf("/unity-catalog/permissions/function/%s", name)
	}
	return fmt.Sprintf("/unity-catalog/permissions/%s/%s", securable, name)
}

//Responses

// UnityCatalogPermissionsDiff is the inner structure of updatePermissions RPC
type UnityCatalogPermissionsDiff struct {
	Changes []UnityCatalogPermissionsChange `json:"changes"`
}

type UnityCatalogPermissionsChange struct {
	Principal string   `json:"principal"`
	Add       []string `json:"add,omitempty"`
	Remove    []string `json:"remove,omitempty"`
}

// PrivilegeAssignment reflects on `grant` block
type UnityCatalogPrivilegeAssignment struct {
	Principal  string   `json:"principal"`
	Privileges []string `json:"privileges" tf:"slice_set"`
}

// PermissionsList reflects it's shape on terraform resource with
// privilege_assignments column renamed to `grant` block for simplicity
type UnityCatalogPermissionsList struct {
	Assignments []UnityCatalogPrivilegeAssignment `json:"privilege_assignments" tf:"slice_set,alias:grant"`
}

type UnityCatalogPermissionsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Permission Mappings

type SecurableMapping map[string]map[string]bool

// reuse ResourceDiff and ResourceData
type attributeGetter interface {
	Get(key string) any
}

func (sm SecurableMapping) KeyValue(d attributeGetter) (string, string) {
	for field := range sm {
		v := d.Get(field).(string)
		if v == "" {
			continue
		}
		return field, v
	}
	return "unknown", "unknown"
}
func (sm SecurableMapping) Id(d *schema.ResourceData) string {
	securable, name := sm.KeyValue(d)
	return fmt.Sprintf("%s/%s", securable, name)
}

func (sm SecurableMapping) Validate(d attributeGetter, pl UnityCatalogPermissionsList) error {
	securable, _ := sm.KeyValue(d)
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

var Mappings = SecurableMapping{
	// add other securable mappings once needed
	"table": {
		"MODIFY": true,
		"SELECT": true,

		// v1.0
		"ALL_PRIVILEGES": true,
		"APPLY_TAG":      true,
		"BROWSE":         true,
	},
	"view": {
		"SELECT":    true,
		"APPLY_TAG": true,
		"BROWSE":    true,
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
	"materialized_view": {
		"ALL_PRIVILEGES": true,
		"SELECT":         true,
		"REFRESH":        true,
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
