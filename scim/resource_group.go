package scim

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"golang.org/x/sync/singleflight"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// groupListAttrs are the SCIM attributes fetched when bulk-listing groups for the read cache.
const groupListAttrs = "id,displayName,externalId,entitlements"

type groupsEntryItem struct {
	mu          sync.RWMutex
	initialized bool
	byID        map[string]Group
	sg          singleflight.Group
}

// groupsListCache caches a ListAll result per (host, apiLevel) so that
// N concurrent reads for N distinct databricks_group resources only issue
// one SCIM list call instead of N individual GET-by-ID calls.
type groupsListCache struct {
	mu    sync.Mutex
	cache map[string]*groupsEntryItem
}

func newGroupsListCache() *groupsListCache {
	return &groupsListCache{cache: make(map[string]*groupsEntryItem)}
}

func (c *groupsListCache) entry(key string) *groupsEntryItem {
	c.mu.Lock()
	defer c.mu.Unlock()
	if e, ok := c.cache[key]; ok {
		return e
	}
	e := &groupsEntryItem{byID: make(map[string]Group)}
	c.cache[key] = e
	return e
}

func (c *groupsListCache) lookup(api GroupsAPI, groupID string) (Group, error) {
	key := api.client.Config.Host + "|" + api.ApiLevel
	e := c.entry(key)

	// Fast path: warm cache, concurrent readers proceed simultaneously.
	e.mu.RLock()
	if e.initialized {
		if g, ok := e.byID[groupID]; ok {
			e.mu.RUnlock()
			return g, nil
		}
		e.mu.RUnlock()
		// Cache populated but group absent (created externally); fall through to direct read.
		return api.Read(groupID, "displayName,externalId,entitlements")
	}
	e.mu.RUnlock()

	// Slow path: populate cache via singleflight — at most one ListAll in-flight
	// per (host, apiLevel); all other goroutines join and wake simultaneously.
	_, err, _ := e.sg.Do("list", func() (interface{}, error) {
		// Double-check after acquiring the singleflight slot.
		e.mu.RLock()
		if e.initialized {
			e.mu.RUnlock()
			return nil, nil
		}
		e.mu.RUnlock()

		groups, err := api.ListAll(groupListAttrs)
		if err != nil {
			return nil, err
		}
		m := make(map[string]Group, len(groups))
		for _, g := range groups {
			m[g.ID] = g
		}
		e.mu.Lock()
		e.byID = m
		e.initialized = true
		e.mu.Unlock()
		return nil, nil
	})
	if err != nil {
		return Group{}, err
	}

	e.mu.RLock()
	g, ok := e.byID[groupID]
	e.mu.RUnlock()
	if !ok {
		return api.Read(groupID, "displayName,externalId,entitlements")
	}
	return g, nil
}

func (c *groupsListCache) invalidate(key string) {
	c.mu.Lock()
	e, ok := c.cache[key]
	c.mu.Unlock()
	if ok {
		e.mu.Lock()
		e.initialized = false
		e.byID = make(map[string]Group)
		e.mu.Unlock()
	}
}

var globalGroupsListCache = newGroupsListCache()

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
			defer globalGroupsListCache.invalidate(c.Config.Host + "|" + common.GetApiLevel(d))
			g := Group{
				DisplayName:  d.Get("display_name").(string),
				Entitlements: readEntitlementsFromData(d),
				ExternalID:   d.Get("external_id").(string),
			}
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
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
			group, err := globalGroupsListCache.lookup(groupsAPI, d.Id())
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
			defer globalGroupsListCache.invalidate(c.Config.Host + "|" + common.GetApiLevel(d))
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
			groupName := d.Get("display_name").(string)
			return groupsAPI.UpdateNameAndEntitlements(d.Id(), groupName,
				d.Get("external_id").(string), readEntitlementsFromData(d))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			defer globalGroupsListCache.invalidate(c.Config.Host + "|" + common.GetApiLevel(d))
			groupsAPI := NewGroupsAPI(ctx, c, common.GetApiLevel(d))
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
