package catalog

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/catalog/permissions"
	"github.com/databricks/terraform-provider-databricks/catalog/util"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// diffPermissionsForPrincipal returns UnityCatalogPermissionsDiff of this permissions list with `diff` privileges removed
func diffPermissionsForPrincipal(principal string, pl permissions.UnityCatalogPermissionsList, existing permissions.UnityCatalogPermissionsList) (diff permissions.UnityCatalogPermissionsDiff) {
	// diffs change sets for principal
	configured := map[string]*schema.Set{}
	for _, v := range pl.Assignments {
		if v.Principal == principal {
			configured[v.Principal] = util.SliceToSet(v.Privileges)
		}
	}
	// existing permissions that needs removal for principal
	remote := map[string]*schema.Set{}
	for _, v := range existing.Assignments {
		if v.Principal == principal {
			remote[v.Principal] = util.SliceToSet(v.Privileges)
		}
	}
	// STEP 1: detect overlaps
	for principal, confPrivs := range configured {
		remotePrivs, ok := remote[principal]
		if !ok {
			remotePrivs = util.SliceToSet([]string{})
		}
		add := util.SetToSlice(confPrivs.Difference(remotePrivs))
		remove := util.SetToSlice(remotePrivs.Difference(confPrivs))
		if len(add) == 0 && len(remove) == 0 {
			continue
		}
		diff.Changes = append(diff.Changes, permissions.UnityCatalogPermissionsChange{
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
		diff.Changes = append(diff.Changes, permissions.UnityCatalogPermissionsChange{
			Principal: principal,
			Remove:    util.SetToSlice(remove),
		})
	}
	// so that we can deterministic tests
	sort.Slice(diff.Changes, func(i, j int) bool {
		return diff.Changes[i].Principal < diff.Changes[j].Principal
	})
	return diff
}

// replacePermissionsForPrincipal merges removal diff of existing permissions on the platform
func replacePermissionsForPrincipal(a permissions.UnityCatalogPermissionsAPI, securable string, name string, principal string, list permissions.UnityCatalogPermissionsList) error {
	existing, err := a.GetPermissions(securable, name)
	if err != nil {
		return err
	}
	return a.UpdatePermissions(securable, name, diffPermissionsForPrincipal(principal, list, existing))
}

func filterPermissionsForPrincipal(in permissions.UnityCatalogPermissionsList, principal string) (out permissions.UnityCatalogPermissionsList) {

	for _, v := range in.Assignments {
		if v.Principal == principal {
			out.Assignments = append(out.Assignments, v)
		}
	}
	return
}

func ResourceGrant() *schema.Resource {
	s := map[string]*schema.Schema{
		"principal": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"privileges": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
		},
	}
	// Handle all of the securable resource types
	allFields := []string{}
	for field := range permissions.Mappings {
		allFields = append(allFields, field)
	}

	for _, field := range allFields {
		s[field] = &schema.Schema{
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: util.SliceWithoutString(allFields, field),
		}
	}

	return common.Resource{
		Schema: s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			if d.Id() == "" {
				// unfortunately we cannot do validation before dependent resources exist with tfsdkv2
				return nil
			}
			var grants permissions.UnityCatalogPermissionsList
			common.DiffToStructPointer(d, s, &grants)
			return permissions.Mappings.Validate(d, grants)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			principal := d.Get("principal").(string)
			privileges := util.ToStringSlice(d.Get("privileges").(*schema.Set).List())
			var grants = permissions.UnityCatalogPermissionsList{
				Assignments: []permissions.UnityCatalogPrivilegeAssignment{
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
			grants, err := permissions.NewUnityCatalogPermissionsAPI(ctx, c).GetPermissions(split[0], split[1])
			grantsForPrincipal := filterPermissionsForPrincipal(grants, principal)
			if err != nil {
				return err
			}
			if len(grantsForPrincipal.Assignments) == 0 {
				return apierr.NotFound("got empty permissions list")
			}
			return common.StructToData(grantsForPrincipal, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			principal := d.Get("principal").(string)
			privileges := util.ToStringSlice(d.Get("privileges").(*schema.Set).List())
			var grants = permissions.UnityCatalogPermissionsList{
				Assignments: []permissions.UnityCatalogPrivilegeAssignment{
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
			return replacePermissionsForPrincipal(unityCatalogPermissionsAPI, split[0], split[1], principal, permissions.UnityCatalogPermissionsList{})
		},
	}.ToResource()
}
