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

// diffPermissionsForPrincipal returns an array of catalog.PermissionsChange of this permissions list with `diff` privileges removed
func diffPermissionsForPrincipal(principal string, desired catalog.PermissionsList, existing catalog.PermissionsList) (diff []catalog.PermissionsChange) {
	// diffs change sets for principal
	configured := map[string]*schema.Set{}
	for _, v := range desired.PrivilegeAssignments {
		if v.Principal == principal {
			configured[v.Principal] = permissions.SliceToSet(v.Privileges)
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

func filterPermissionsForPrincipal(in catalog.PermissionsList, principal string) (out catalog.PermissionsList) {

	for _, v := range in.PrivilegeAssignments {
		if v.Principal == principal {
			out.PrivilegeAssignments = append(out.PrivilegeAssignments, v)
		}
	}
	return
}

func ResourceGrant() *schema.Resource {
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
			principal := d.Get("principal").(string)
			privileges := permissions.ToPrivilegeSlice(d.Get("privileges").(*schema.Set).List())
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
			err := replacePermissionsForPrincipal(unityCatalogPermissionsAPI, securable, name, principal, grants)
			if err != nil {
				return err
			}
			d.SetId(fmt.Sprintf("%s/%s", permissions.Mappings.Id(d), principal))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			principal := d.Get("principal").(string)
			split := strings.SplitN(d.Id(), "/", 3)
			if len(split) != 3 {
				return fmt.Errorf("ID must be three elements split by `/`: %s", d.Id())
			}
			grants, err := permissions.NewUnityCatalogPermissionsAPI(ctx, c).GetPermissions(permissions.Mappings.GetSecurableType(split[0]), split[1])
			grantsForPrincipal := filterPermissionsForPrincipal(*grants, principal)
			if err != nil {
				return err
			}
			if len(grantsForPrincipal.PrivilegeAssignments) == 0 {
				return apierr.NotFound("got empty permissions list")
			}
			return common.StructToData(grantsForPrincipal, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			principal := d.Get("principal").(string)
			privileges := permissions.ToPrivilegeSlice(d.Get("privileges").(*schema.Set).List())
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
			return replacePermissionsForPrincipal(unityCatalogPermissionsAPI, securable, name, principal, grants)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			principal := d.Get("principal").(string)
			split := strings.SplitN(d.Id(), "/", 3)
			if len(split) != 3 {
				return fmt.Errorf("ID must be three elements split by `/`: %s", d.Id())
			}
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			return replacePermissionsForPrincipal(unityCatalogPermissionsAPI, split[0], split[1], principal, catalog.PermissionsList{})
		},
	}.ToResource()
}
