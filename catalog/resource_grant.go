package catalog

import (
	"context"
	"errors"
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

// diffPermissionsForPrincipal returns an array of catalog.PermissionsChange of this permissions list with `diff` privileges removed
func diffPermissionsForPrincipal(principal string, desired catalog.PermissionsList, existing catalog.PermissionsList) (diff []catalog.PermissionsChange) {
	// diffs change sets for principal
	configured := map[string]*schema.Set{}
	for _, v := range desired.PrivilegeAssignments {
		if v.Principal == principal {
			normalizedPrivileges := []catalog.Privilege{}
			for _, p := range v.Privileges {
				normalizedPriv := strings.ReplaceAll(p.String(), " ", "_")
				normalizedPrivileges = append(normalizedPrivileges, catalog.Privilege(normalizedPriv))
			}
			configured[v.Principal] = permissions.SliceToSet(normalizedPrivileges)
		}
	}
	// existing permissions that needs removal for principal
	remote := map[string]*schema.Set{}
	for _, v := range existing.PrivilegeAssignments {
		if v.Principal == principal {
			remote[v.Principal] = permissions.SliceToSet(v.Privileges)
		}
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

// replacePermissionsForPrincipal merges removal diff of existing permissions on the platform
func replacePermissionsForPrincipal(a permissions.UnityCatalogPermissionsAPI, securable string, name string, principal string, list catalog.PermissionsList) error {
	securableType := permissions.Mappings.GetSecurableType(securable)
	existing, err := a.GetPermissions(securableType, name)
	if err != nil {
		return err
	}
	err = a.UpdatePermissions(securableType, name, diffPermissionsForPrincipal(principal, list, *existing))
	if err != nil {
		return err
	}
	return a.WaitForUpdate(1*time.Minute, securableType, name, list, func(current *catalog.PermissionsList, desired catalog.PermissionsList) []catalog.PermissionsChange {
		return diffPermissionsForPrincipal(principal, desired, *current)
	})
}

// filterPermissionsForPrincipal extracts permissions for the given principal and transforms to permissions.UnityCatalogPrivilegeAssignment to match Schema
func filterPermissionsForPrincipal(in catalog.PermissionsList, principal string) (*permissions.UnityCatalogPrivilegeAssignment, error) {
	grantsForPrincipal := []permissions.UnityCatalogPrivilegeAssignment{}
	for _, v := range in.PrivilegeAssignments {
		privileges := []string{}
		if v.Principal == principal {
			for _, p := range v.Privileges {
				privileges = append(privileges, p.String())
			}
			grantsForPrincipal = append(grantsForPrincipal, permissions.UnityCatalogPrivilegeAssignment{
				Principal:  v.Principal,
				Privileges: privileges,
			})
		}
	}
	if len(grantsForPrincipal) == 0 {
		return nil, apierr.NotFound("got empty permissions list")
	}
	if len(grantsForPrincipal) > 1 {
		return nil, errors.New("got more than one principal in permissions list")
	}
	return &grantsForPrincipal[0], nil
}

func toSecurableId(d *schema.ResourceData) string {
	principal := d.Get("principal").(string)
	return fmt.Sprintf("%s/%s", permissions.Mappings.Id(d), principal)
}

func parseSecurableId(d *schema.ResourceData) (string, string, string, error) {
	split := strings.SplitN(d.Id(), "/", 3)
	if len(split) != 3 {
		return "", "", "", fmt.Errorf("ID must be three elements split by `/`: %s", d.Id())
	}
	return split[0], split[1], split[2], nil
}

func ResourceGrant() common.Resource {
	s := common.StructToSchema(permissions.UnityCatalogPrivilegeAssignment{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {

			m["principal"].ForceNew = true

			allFields := []string{}
			for field := range permissions.Mappings {
				allFields = append(allFields, field)
			}
			for field := range permissions.Mappings {
				m[field] = &schema.Schema{
					Type:          schema.TypeString,
					Optional:      true,
					ForceNew:      true,
					AtLeastOneOf:  allFields,
					ConflictsWith: permissions.SliceWithoutString(allFields, field),
				}
			}
			return m
		})

	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore").(string))
			if err != nil {
				return err
			}
			principal := d.Get("principal").(string)
			privileges := permissions.SetToSlice(d.Get("privileges").(*schema.Set))
			var grants = catalog.PermissionsList{
				PrivilegeAssignments: []catalog.PrivilegeAssignment{
					{
						Principal:  principal,
						Privileges: privileges,
					},
				},
			}
			securable, name := permissions.Mappings.KeyValue(d)
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			err = replacePermissionsForPrincipal(unityCatalogPermissionsAPI, securable, name, principal, grants)
			if err != nil {
				return err
			}
			d.SetId(toSecurableId(d))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			securable, name, principal, err := parseSecurableId(d)
			if err != nil {
				return err
			}
			grants, err := permissions.NewUnityCatalogPermissionsAPI(ctx, c).GetPermissions(permissions.Mappings.GetSecurableType(securable), name)
			if err != nil {
				return err
			}
			grantsForPrincipal, err := filterPermissionsForPrincipal(*grants, principal)
			if err != nil {
				return err
			}
			err = common.StructToData(*grantsForPrincipal, s, d)
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
			securable, name, principal, err := parseSecurableId(d)
			if err != nil {
				return err
			}
			privileges := permissions.SetToSlice(d.Get("privileges").(*schema.Set))
			var grants = catalog.PermissionsList{
				PrivilegeAssignments: []catalog.PrivilegeAssignment{
					{
						Principal:  principal,
						Privileges: privileges,
					},
				},
			}
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			return replacePermissionsForPrincipal(unityCatalogPermissionsAPI, securable, name, principal, grants)
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
			securable, name, principal, err := parseSecurableId(d)
			if err != nil {
				return err
			}
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			return replacePermissionsForPrincipal(unityCatalogPermissionsAPI, securable, name, principal, catalog.PermissionsList{})
		},
	}
}
