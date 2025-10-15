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
	common.Namespace
	Assignments []PrivilegeAssignment `json:"privilege_assignments" tf:"slice_set,alias:grant"`
}

// diffPermissions returns an array of catalog.PermissionsChange of this permissions list with `diff` privileges removed
func diffPermissions(pl []catalog.PrivilegeAssignment, existing []catalog.PrivilegeAssignment) (diff []catalog.PermissionsChange) {
	// diffs change sets
	configured := map[string]*schema.Set{}
	for _, v := range pl {
		configured[v.Principal] = permissions.SliceToSet(v.Privileges)
	}
	// existing permissions that needs removal
	remote := map[string]*schema.Set{}
	for _, v := range existing {
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
func replaceAllPermissions(a permissions.UnityCatalogPermissionsAPI, securable string, name string, list catalog.GetPermissionsResponse) error {
	securableType := permissions.Mappings.GetSecurableType(securable)
	existing, err := a.GetPermissions(securableType, name)
	if err != nil {
		return err
	}
	err = a.UpdatePermissions(securableType, name, diffPermissions(list.PrivilegeAssignments, existing.PrivilegeAssignments))
	if err != nil {
		return err
	}
	return a.WaitForUpdate(1*time.Minute, securableType, name, list.PrivilegeAssignments, func(current []catalog.PrivilegeAssignment, desired []catalog.PrivilegeAssignment) []catalog.PermissionsChange {
		return diffPermissions(desired, current)
	})
}

func (pl PermissionsList) toSdkPermissionsList() (out catalog.GetPermissionsResponse) {
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

func sdkPermissionsListToPermissionsList(sdkPermissionsList []catalog.PrivilegeAssignment) (out PermissionsList) {
	for _, v := range sdkPermissionsList {
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
			// set custom hash function for principal and privileges
			common.MustSchemaPath(s, "grant", "privileges").Set = func(i any) int {
				privilege := i.(string)
				return schema.HashString(permissions.NormalizePrivilege(privilege))
			}
			common.MustSchemaPath(s, "grant").Set = func(i any) int {
				objectStruct := i.(map[string]any)
				principal := objectStruct["principal"].(string)
				privileges := objectStruct["privileges"].(*schema.Set)
				hashString := strings.ToLower(principal)
				for _, privilege := range privileges.List() {
					hashString += "|" + permissions.NormalizePrivilege(privilege.(string))
				}
				return schema.HashString(hashString)
			}
			alof := []string{}
			for field := range permissions.Mappings {
				s[field] = &schema.Schema{
					Type:     schema.TypeString,
					ForceNew: true,
					Optional: true,
				}
				alof = append(alof, field)
			}
			for field := range permissions.Mappings {
				s[field].AtLeastOneOf = alof
			}
			common.NamespaceCustomizeSchemaMap(s)
			return s
		})
	return common.Resource{
		Schema: s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore").(string))
			if err != nil {
				return err
			}
			var grants PermissionsList
			common.DataToStructPointer(d, s, &grants)
			securable, name := permissions.Mappings.KeyValue(d)
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			err = replaceAllPermissions(unityCatalogPermissionsAPI, securable, name, grants.toSdkPermissionsList())
			if err != nil {
				return err
			}
			d.SetId(permissions.Mappings.Id(d))
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
				return &apierr.APIError{
					ErrorCode:  "NOT_FOUND",
					StatusCode: 404,
					Message:    "got empty permissions list",
				}
			}

			err = common.StructToData(sdkPermissionsListToPermissionsList(grants.PrivilegeAssignments), s, d)
			if err != nil {
				return err
			}
			return d.Set(securable, name)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
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
			return replaceAllPermissions(unityCatalogPermissionsAPI, securable, name, catalog.GetPermissionsResponse{})
		},
	}
}
