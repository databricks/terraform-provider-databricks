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
	// mutatedInBulkGen records the bulk-fetch generation (groupEndpointCache.bulkGen)
	// that was active when this group was last mutated by addMember/removeMember.
	// The bulk commit uses this to skip groups that were mutated while the specific
	// ListAll request was in-flight: if mutatedInBulkGen equals the current bulk
	// generation the snapshot was captured before the mutation, so committing it
	// would reintroduce stale data.  Groups flagged this way fall back to the
	// individual-read path in getMembers to obtain fresh data.
	mutatedInBulkGen uint64
	members          map[string]struct{}
	lock             sync.RWMutex
	// sg deduplicates concurrent in-flight API calls when the cache is cold.
	// All goroutines waiting on a cold cache join the single fetch and wake
	// simultaneously, rather than serialising through a write lock.
	sg singleflight.Group
}

// groupEndpointCache is the per-endpoint (host|apiLevel|accountID) member cache.
// Scoping by endpoint prevents two provider instances targeting different accounts
// (or different API levels) from sharing cached group-member data.
type groupEndpointCache struct {
	lock sync.Mutex
	// cache maps group ID to its member info for this endpoint.
	cache map[string]*groupMembersInfo
	// bulkMu / bulkGen / bulkInitialized / bulkSg guard the one-time ListAll("id,members") bulk fetch.
	// bulkGen is monotonically incremented each time a bulk fetch starts; addMember/removeMember
	// record the current bulkGen so the bulk commit can detect mid-flight mutations.
	// After the bulk is done, per-group fast-path reads skip any API calls entirely.
	bulkMu          sync.RWMutex
	bulkGen         uint64
	bulkInitialized bool
	bulkSg          singleflight.Group
}

func newGroupEndpointCache() *groupEndpointCache {
	return &groupEndpointCache{cache: make(map[string]*groupMembersInfo)}
}

func (ep *groupEndpointCache) getOrCreateGroupInfo(groupID string) *groupMembersInfo {
	ep.lock.Lock()
	defer ep.lock.Unlock()
	if info, ok := ep.cache[groupID]; ok {
		return info
	}
	info := &groupMembersInfo{members: make(map[string]struct{})}
	ep.cache[groupID] = info
	return info
}

func (ep *groupEndpointCache) getMembers(api GroupsAPI, groupID string) (map[string]struct{}, error) {
	groupInfo := ep.getOrCreateGroupInfo(groupID)

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
	ep.bulkMu.RLock()
	bulkDone := ep.bulkInitialized
	ep.bulkMu.RUnlock()
	if !bulkDone {
		_, err, _ := ep.bulkSg.Do("bulk", func() (interface{}, error) {
			// Double-check after acquiring the singleflight slot.
			ep.bulkMu.RLock()
			if ep.bulkInitialized {
				ep.bulkMu.RUnlock()
				return nil, nil
			}
			ep.bulkMu.RUnlock()

			// Assign a new bulk generation before starting the API call.
			// addMember/removeMember that run concurrently will record this
			// generation so the commit loop below can skip those groups.
			ep.bulkMu.Lock()
			ep.bulkGen++
			myBulkGen := ep.bulkGen
			ep.bulkMu.Unlock()

			tflog.Debug(api.context, "Fetching all group memberships in bulk")
			groups, err := api.ListAll("id,members")
			if err != nil {
				return nil, err
			}
			for _, g := range groups {
				info := ep.getOrCreateGroupInfo(g.ID)
				info.lock.Lock()
				// Skip groups whose mutatedInBulkGen matches the current bulk
				// generation: that means addMember/removeMember ran concurrently
				// with this ListAll, so our snapshot predates the mutation and
				// committing it would cache stale membership data.  Those groups
				// will be served by the individual-read fallback instead.
				if !info.initialized && info.mutatedInBulkGen != myBulkGen {
					info.members = make(map[string]struct{}, len(g.Members))
					for _, m := range g.Members {
						info.members[m.Value] = struct{}{}
					}
					info.initialized = true
				}
				info.lock.Unlock()
			}
			ep.bulkMu.Lock()
			ep.bulkInitialized = true
			ep.bulkMu.Unlock()
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

func (ep *groupEndpointCache) removeMember(api GroupsAPI, groupID string, memberID string) error {
	groupInfo := ep.getOrCreateGroupInfo(groupID)
	groupInfo.lock.Lock()
	defer groupInfo.lock.Unlock()

	err := api.Patch(groupID, PatchRequest(
		"remove", fmt.Sprintf(`members[value eq "%s"]`, memberID)))
	if err != nil {
		return err
	}
	// Record the current bulk generation before updating the map. The bulk
	// commit skips any group whose mutatedInBulkGen equals the in-flight
	// generation, preventing stale snapshot data from overwriting this mutation.
	ep.bulkMu.RLock()
	groupInfo.mutatedInBulkGen = ep.bulkGen
	ep.bulkMu.RUnlock()
	if groupInfo.initialized {
		delete(groupInfo.members, memberID)
	}
	return nil
}

func (ep *groupEndpointCache) addMember(api GroupsAPI, groupID string, memberID string) error {
	groupInfo := ep.getOrCreateGroupInfo(groupID)
	groupInfo.lock.Lock()
	defer groupInfo.lock.Unlock()

	err := api.Patch(groupID, PatchRequestWithValue("add", "members", memberID))
	if err != nil {
		return err
	}
	// Record the current bulk generation before updating the map. The bulk
	// commit skips any group whose mutatedInBulkGen equals the in-flight
	// generation, preventing stale snapshot data from overwriting this mutation.
	ep.bulkMu.RLock()
	groupInfo.mutatedInBulkGen = ep.bulkGen
	ep.bulkMu.RUnlock()
	if groupInfo.initialized {
		groupInfo.members[memberID] = struct{}{}
	}
	return nil
}

// groupCache maps endpoint keys to per-endpoint caches.
// Scoping by (host, apiLevel, accountID) prevents two providers targeting
// different accounts or API levels from sharing cached group-member data.
type groupCache struct {
	mu        sync.Mutex
	endpoints map[string]*groupEndpointCache
}

func newGroupCache() *groupCache {
	return &groupCache{endpoints: make(map[string]*groupEndpointCache)}
}

func groupMembersCacheKey(api GroupsAPI) string {
	return api.client.Config.Host + "|" + api.ApiLevel + "|" + api.client.Config.AccountID
}

func (gc *groupCache) endpointFor(api GroupsAPI) *groupEndpointCache {
	key := groupMembersCacheKey(api)
	gc.mu.Lock()
	defer gc.mu.Unlock()
	if ep, ok := gc.endpoints[key]; ok {
		return ep
	}
	ep := newGroupEndpointCache()
	gc.endpoints[key] = ep
	return ep
}

func (gc *groupCache) getMembers(api GroupsAPI, groupID string) (map[string]struct{}, error) {
	return gc.endpointFor(api).getMembers(api, groupID)
}

func (gc *groupCache) addMember(api GroupsAPI, groupID string, memberID string) error {
	return gc.endpointFor(api).addMember(api, groupID, memberID)
}

func (gc *groupCache) removeMember(api GroupsAPI, groupID string, memberID string) error {
	return gc.endpointFor(api).removeMember(api, groupID, memberID)
}

func copyGroupMembers(m map[string]struct{}) map[string]struct{} {
	cp := make(map[string]struct{}, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return cp
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
