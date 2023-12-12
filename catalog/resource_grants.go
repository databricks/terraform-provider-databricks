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

// permissionsDiff is the inner structure of updatePermissions RPC
// diff returns permissionsDiff of this permissions list with `diff` privileges removed
func diffPermissions(pl permissions.UnityCatalogPermissionsList, existing permissions.UnityCatalogPermissionsList) (diff permissions.UnityCatalogPermissionsDiff) {
	// diffs change sets
	configured := map[string]*schema.Set{}
	for _, v := range pl.Assignments {
		configured[v.Principal] = util.SliceToSet(v.Privileges)
	}
	// existing permissions that needs removal
	remote := map[string]*schema.Set{}
	for _, v := range existing.Assignments {
		remote[v.Principal] = util.SliceToSet(v.Privileges)
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

// replaceAllPermissions merges removal diff of existing permissions on the platform
func replaceAllPermissions(a permissions.UnityCatalogPermissionsAPI, securable string, name string, list permissions.UnityCatalogPermissionsList) error {
	existing, err := a.GetPermissions(securable, name)
	if err != nil {
		return err
	}
	return a.UpdatePermissions(securable, name, diffPermissions(list, existing))
}

func ResourceGrants() *schema.Resource {
	s := common.StructToSchema(permissions.UnityCatalogPermissionsList{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
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
			return s
		})
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
			var grants permissions.UnityCatalogPermissionsList
			common.DataToStructPointer(d, s, &grants)
			securable, name := permissions.Mappings.KeyValue(d)
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			err := replaceAllPermissions(unityCatalogPermissionsAPI, securable, name, grants)
			if err != nil {
				return err
			}
			d.SetId(permissions.Mappings.Id(d))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			split := strings.SplitN(d.Id(), "/", 2)
			if len(split) != 2 {
				return fmt.Errorf("ID must be two elements split by `/`: %s", d.Id())
			}
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			grants, err := unityCatalogPermissionsAPI.GetPermissions(split[0], split[1])
			if err != nil {
				return err
			}
			if len(grants.Assignments) == 0 {
				return apierr.NotFound("got empty permissions list")
			}
			return common.StructToData(grants, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			securable, name := permissions.Mappings.KeyValue(d)
			var grants permissions.UnityCatalogPermissionsList
			common.DataToStructPointer(d, s, &grants)
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			return replaceAllPermissions(unityCatalogPermissionsAPI, securable, name, grants)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			split := strings.SplitN(d.Id(), "/", 2)
			if len(split) != 2 {
				return fmt.Errorf("ID must be two elements split by `/`: %s", d.Id())
			}
			unityCatalogPermissionsAPI := permissions.NewUnityCatalogPermissionsAPI(ctx, c)
			return replaceAllPermissions(unityCatalogPermissionsAPI, split[0], split[1], permissions.UnityCatalogPermissionsList{})
		},
	}.ToResource()
}
