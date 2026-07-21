package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// groupListAttrs are the SCIM attributes fetched when bulk-listing groups for the read cache.
const groupListAttrs = "id,displayName,externalId,entitlements"

// globalGroupsListCache caches a ListAll result per (host, apiLevel, accountID) so
// that N concurrent reads for N distinct databricks_group resources only issue one
// SCIM list call instead of N individual GET-by-ID calls.
//
// The cache key includes accountID so that two provider instances targeting the
// same API host but different Databricks accounts never share cached groups.
var globalGroupsListCache = newGroupsListCache()

func newGroupsListCache() *common.KeyedCache[string, map[string]Group] {
	return common.NewKeyedCache[string, map[string]Group]()
}

func groupsCacheKey(api GroupsAPI) string {
	return api.client.Config.Host + "|" + api.ApiLevel + "|" + api.client.Config.AccountID
}

func groupsListCacheLookup(api GroupsAPI, groupID string) (Group, error) {
	byID, err := globalGroupsListCache.Get(groupsCacheKey(api), func() (map[string]Group, error) {
		groups, err := api.ListAll(groupListAttrs)
		if err != nil {
			return nil, err
		}
		m := make(map[string]Group, len(groups))
		for _, g := range groups {
			m[g.ID] = g
		}
		return m, nil
	})
	if err != nil {
		return Group{}, err
	}
	if g, ok := byID[groupID]; ok {
		return g, nil
	}
	// Cache populated but group absent (e.g. created concurrently); fall back to a direct read.
	return api.Read(groupID, "displayName,externalId,entitlements")
}

// ResourceGroup manages user groups
func ResourceGroup() common.Resource {
	type entity struct {
		DisplayName string `json:"display_name"`
		ExternalID  string `json:"external_id,omitempty" tf:"force_new,suppress_diff"`
		URL         string `json:"url,omitempty" tf:"computed"`
	}
	groupSchema := common.StructToSchema(entity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			addEntitlementsToSchema(m)
			common.AddApiField(m)
			// https://github.com/databricks/terraform-provider-databricks/issues/1089
			m["display_name"].ValidateDiagFunc = validation.ToDiagFunc(
				validation.StringNotInSlice([]string{"users", "admins"}, false))
			m["force"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["acl_principal_id"] = &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}
			return m
		})
	addEntitlementsToSchema(groupSchema)
	common.AddNamespaceInSchema(groupSchema)
	common.NamespaceCustomizeSchemaMap(groupSchema)
	return common.Resource{
		IsDual: true,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.CustomizeDiffDualResources(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
			defer globalGroupsListCache.Invalidate(groupsCacheKey(groupsAPI))
			g := Group{
				DisplayName:  d.Get("display_name").(string),
				Entitlements: readEntitlementsFromData(d),
				ExternalID:   d.Get("external_id").(string),
			}
			group, err := groupsAPI.Create(g)
			if err != nil {
				return createForceOverridesManuallyAddedGroup(err, d, groupsAPI, g)
			}
			d.SetId(group.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
			group, err := groupsListCacheLookup(groupsAPI, d.Id())
			if err != nil {
				return err
			}
			d.Set("display_name", group.DisplayName)
			d.Set("external_id", group.ExternalID)
			d.Set("acl_principal_id", fmt.Sprintf("groups/%s", group.DisplayName))
			if common.IsAccountLevel(d, c) {
				d.Set("url", c.FormatURL("users/groups/", d.Id(), "/information"))
			} else {
				d.Set("url", c.FormatURL("#setting/accounts/groups/", d.Id()))
			}
			return group.Entitlements.readIntoData(d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
			defer globalGroupsListCache.Invalidate(groupsCacheKey(groupsAPI))
			groupName := d.Get("display_name").(string)
			return groupsAPI.UpdateNameAndEntitlements(d.Id(), groupName,
				d.Get("external_id").(string), readEntitlementsFromData(d))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
			defer globalGroupsListCache.Invalidate(groupsCacheKey(groupsAPI))
			return groupsAPI.Delete(d.Id())
		},
		Schema: groupSchema,
	}
}

func createForceOverridesManuallyAddedGroup(err error, d *schema.ResourceData, groupsAPI GroupsAPI, g Group) error {
	forceCreate := d.Get("force").(bool)
	if !forceCreate {
		return err
	}
	// corner-case for overriding manually provisioned groups
	groupName := strings.ReplaceAll(g.DisplayName, "'", "")
	force := fmt.Sprintf("Group with name %s already exists.", groupName)
	if err.Error() != force {
		return err
	}
	group, err := groupsAPI.ReadByDisplayName(groupName, "")
	if err != nil {
		return err
	}
	d.SetId(group.ID)
	return groupsAPI.UpdateNameAndEntitlements(d.Id(), g.DisplayName, g.ExternalID, g.Entitlements)
}
