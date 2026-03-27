package scim

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/singleflight"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// groupMembersInfo is the set of members that belong to a group. Read and write access
// to the underlying members map must be protected by the lock.
type groupMembersInfo struct {
	initialized bool
	members     map[string]struct{}
	lock        sync.RWMutex
	// sg deduplicates concurrent in-flight API calls when the cache is cold.
	// All goroutines waiting on a cold cache join the single fetch and wake
	// simultaneously, rather than serialising through a write lock.
	sg singleflight.Group
}

// groupCache is a cache of groups to their members by their ID. Access to the cache
// must be protected by the lock.
type groupCache struct {
	// The mapping of cached members for this group. The cache key is the group ID.
	// TODO: add workspace ID to the cache key when account-level and workspace-level providers are unified.
	cache map[string]*groupMembersInfo
	lock  sync.Mutex
	// bulkInitialized is true once a ListAll("id,members") has populated all entries.
	// After that, per-group fast-path reads skip any API calls entirely.
	bulkMu          sync.RWMutex
	bulkInitialized bool
	bulkSg          singleflight.Group
}

func newGroupCache() *groupCache {
	return &groupCache{
		cache: make(map[string]*groupMembersInfo),
	}
}

func (gc *groupCache) getOrCreateGroupInfo(groupID string) *groupMembersInfo {
	gc.lock.Lock()
	defer gc.lock.Unlock()

	groupInfo, exists := gc.cache[groupID]
	if !exists {
		groupInfo = &groupMembersInfo{
			members: make(map[string]struct{}),
		}
		gc.cache[groupID] = groupInfo
	}
	return groupInfo
}

func copyGroupMembers(m map[string]struct{}) map[string]struct{} {
	cp := make(map[string]struct{}, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func (gc *groupCache) getMembers(api GroupsAPI, groupID string) (map[string]struct{}, error) {
	groupInfo := gc.getOrCreateGroupInfo(groupID)

	// Fast path: per-group cache is warm — concurrent readers hold RLock simultaneously.
	groupInfo.lock.RLock()
	if groupInfo.initialized {
		tflog.Debug(api.context, fmt.Sprintf("Group %s has %d members (cached)", groupID, len(groupInfo.members)))
		membersCopy := copyGroupMembers(groupInfo.members)
		groupInfo.lock.RUnlock()
		return membersCopy, nil
	}
	groupInfo.lock.RUnlock()

	// Bulk path: fetch all groups+members in one ListAll call;
	// all goroutines share a single in-flight request via singleflight.
	gc.bulkMu.RLock()
	bulkDone := gc.bulkInitialized
	gc.bulkMu.RUnlock()
	if !bulkDone {
		_, err, _ := gc.bulkSg.Do("bulk", func() (interface{}, error) {
			// Double-check after acquiring the singleflight slot.
			gc.bulkMu.RLock()
			if gc.bulkInitialized {
				gc.bulkMu.RUnlock()
				return nil, nil
			}
			gc.bulkMu.RUnlock()

			tflog.Debug(api.context, "Fetching all group memberships in bulk")
			groups, err := api.ListAll("id,members")
			if err != nil {
				return nil, err
			}
			for _, g := range groups {
				info := gc.getOrCreateGroupInfo(g.ID)
				info.lock.Lock()
				if !info.initialized {
					info.members = make(map[string]struct{}, len(g.Members))
					for _, m := range g.Members {
						info.members[m.Value] = struct{}{}
					}
					info.initialized = true
				}
				info.lock.Unlock()
			}
			gc.bulkMu.Lock()
			gc.bulkInitialized = true
			gc.bulkMu.Unlock()
			return nil, nil
		})
		if err != nil {
			return nil, err
		}
	}

	// Re-check per-group cache after bulk.
	groupInfo.lock.RLock()
	if groupInfo.initialized {
		tflog.Debug(api.context, fmt.Sprintf("Group %s has %d members (initialized)", groupID, len(groupInfo.members)))
		membersCopy := copyGroupMembers(groupInfo.members)
		groupInfo.lock.RUnlock()
		return membersCopy, nil
	}
	groupInfo.lock.RUnlock()

	// Group was not returned by the bulk fetch (e.g., created externally after the
	// bulk ran). Fall back to an individual read, deduped via singleflight.
	_, err, _ := groupInfo.sg.Do("fetch", func() (interface{}, error) {
		// Double-check after acquiring the singleflight slot.
		groupInfo.lock.RLock()
		if groupInfo.initialized {
			groupInfo.lock.RUnlock()
			return nil, nil
		}
		groupInfo.lock.RUnlock()

		tflog.Debug(api.context, fmt.Sprintf("Getting members for group %s (not in bulk fetch)", groupID))
		group, err := api.Read(groupID, "members")
		if err != nil {
			return nil, err
		}
		tflog.Debug(api.context, fmt.Sprintf("Group %s has %d members", groupID, len(group.Members)))
		groupInfo.lock.Lock()
		groupInfo.members = make(map[string]struct{}, len(group.Members))
		for _, member := range group.Members {
			groupInfo.members[member.Value] = struct{}{}
		}
		groupInfo.initialized = true
		groupInfo.lock.Unlock()
		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	groupInfo.lock.RLock()
	tflog.Debug(api.context, fmt.Sprintf("Group %s has %d members (initialized)", groupID, len(groupInfo.members)))
	membersCopy := copyGroupMembers(groupInfo.members)
	groupInfo.lock.RUnlock()
	return membersCopy, nil
}

func (gc *groupCache) removeMember(api GroupsAPI, groupID string, memberID string) error {
	groupInfo := gc.getOrCreateGroupInfo(groupID)
	groupInfo.lock.Lock()
	defer groupInfo.lock.Unlock()

	err := api.Patch(groupID, PatchRequest(
		"remove", fmt.Sprintf(`members[value eq "%s"]`, memberID)))
	if err != nil {
		return err
	}

	if groupInfo.initialized {
		delete(groupInfo.members, memberID)
	}
	return nil
}

func (gc *groupCache) addMember(api GroupsAPI, groupID string, memberID string) error {
	groupInfo := gc.getOrCreateGroupInfo(groupID)
	groupInfo.lock.Lock()
	defer groupInfo.lock.Unlock()

	err := api.Patch(groupID, PatchRequestWithValue("add", "members", memberID))
	if err != nil {
		return err
	}
	if groupInfo.initialized {
		groupInfo.members[memberID] = struct{}{}
	}
	return nil
}

func hasMember(members map[string]struct{}, memberID string) bool {
	_, ok := members[memberID]
	return ok
}

var globalGroupsCache = newGroupCache()

// ResourceGroupMember bind group with member
func ResourceGroupMember() common.Resource {
	r := common.NewPairID("group_id", "member_id").Schema(func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.AddApiField(m)
		common.AddNamespaceInSchema(m)
		common.NamespaceCustomizeSchemaMap(m)
		return m
	}).BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			return globalGroupsCache.addMember(NewGroupsAPI(ctx, c, common.GetApiLevel(d)), groupID, memberID)
		},
		ReadContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			members, err := globalGroupsCache.getMembers(NewGroupsAPI(ctx, c, common.GetApiLevel(d)), groupID)
			if err == nil && !hasMember(members, memberID) {
				return &apierr.APIError{
					ErrorCode:  "NOT_FOUND",
					StatusCode: 404,
					Message:    "Group has no member",
				}
			}
			return err
		},
		DeleteContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient, d *schema.ResourceData) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			return globalGroupsCache.removeMember(NewGroupsAPI(ctx, c, common.GetApiLevel(d)), groupID, memberID)
		},
	})
	r.CustomizeDiff = func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
		return common.CustomizeDiffDualResources(ctx, d, c)
	}
	r.IsDual = true
	return r
}
