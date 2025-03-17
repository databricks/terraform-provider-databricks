package scim

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bluele/gcache"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type GroupMembersInfo struct {
	initialized bool
	members     map[string]struct{}
	lock        *sync.Mutex
}

type GroupCache struct {
	cache gcache.Cache
}

func NewGroupsCache(expiration time.Duration) *GroupCache {
	return &GroupCache{
		cache: gcache.New(1000).LRU().LoaderFunc(func(key interface{}) (interface{}, error) {
			return &GroupMembersInfo{
				initialized: false,
				members:     make(map[string]struct{}),
				lock:        &sync.Mutex{},
			}, nil
		}).Expiration(expiration).Build(),
	}
}

func (c *GroupCache) GetMembers(api GroupsAPI, groupID string) (map[string]struct{}, error) {
	entry, err := c.cache.Get(groupID)
	if err != nil {
		return nil, err
	}
	groupInfo := entry.(*GroupMembersInfo)
	groupInfo.lock.Lock()
	defer groupInfo.lock.Unlock()

	if !groupInfo.initialized {
		tflog.Debug(api.context, fmt.Sprintf("Getting members for group %s", groupID))
		group, err := api.Read(groupID, "members")
		if err != nil {
			return nil, err
		}
		tflog.Debug(api.context, fmt.Sprintf("Group %s has %d members", groupID, len(group.Members)))
		for _, member := range group.Members {
			memberKey := strings.ToLower(member.Value)
			groupInfo.members[memberKey] = struct{}{}
		}
		groupInfo.initialized = true
		return groupInfo.members, nil
	}
	tflog.Debug(api.context, fmt.Sprintf("Group %s has %d members (cached)", groupID, len(groupInfo.members)))
	return groupInfo.members, nil
}

func (c *GroupCache) removeMember(api GroupsAPI, groupID string, memberID string) error {
	entry, err := c.cache.Get(groupID)
	if err != nil {
		return err
	}
	groupInfo := entry.(*GroupMembersInfo)
	groupInfo.lock.Lock()
	defer groupInfo.lock.Unlock()

	err = api.Patch(groupID, PatchRequest(
		"remove", fmt.Sprintf(`members[value eq "%s"]`, memberID)))
	if err != nil {
		return err
	}

	if groupInfo.initialized {
		memberKey := strings.ToLower(memberID)
		delete(groupInfo.members, memberKey)
	}
	return err

}

func (c *GroupCache) addMember(api GroupsAPI, groupID string, memberID string) error {
	entry, err := c.cache.Get(groupID)
	if err != nil {
		return err
	}
	groupInfo := entry.(*GroupMembersInfo)
	groupInfo.lock.Lock()
	defer groupInfo.lock.Unlock()

	err = api.Patch(groupID, PatchRequestWithValue("add", "members", memberID))
	if err != nil {
		return err
	}
	if groupInfo.initialized {
		memberKey := strings.ToLower(memberID)
		groupInfo.members[memberKey] = struct{}{}
	}
	return err
}

func hasMember(members map[string]struct{}, memberID string) bool {
	_, ok := members[memberID]
	return ok
}

var globalGroupsCache = NewGroupsCache(time.Second * 300)

// ResourceGroupMember bind group with member
func ResourceGroupMember() common.Resource {
	return common.NewPairID("group_id", "member_id").BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			return globalGroupsCache.addMember(NewGroupsAPI(ctx, c), groupID, memberID)
		},
		ReadContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			members, err := globalGroupsCache.GetMembers(NewGroupsAPI(ctx, c), groupID)
			if err == nil && !hasMember(members, memberID) {
				return apierr.NotFound(fmt.Sprintf("Group has no member %s %d", memberID, len(members)))
			}
			return err
		},
		DeleteContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			return globalGroupsCache.removeMember(NewGroupsAPI(ctx, c), groupID, memberID)
		},
	})
}
