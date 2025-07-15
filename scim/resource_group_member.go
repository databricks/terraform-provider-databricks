package scim

import (
	"context"
	"fmt"
	"strings"
	"sync"

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
	cache map[string]*GroupMembersInfo
	lock  sync.Mutex
}

func NewGroupsCache() *GroupCache {
	return &GroupCache{
		cache: make(map[string]*GroupMembersInfo),
		lock:  sync.Mutex{},
	}
}

func (c *GroupCache) getOrCreateGroupInfo(groupID string) *GroupMembersInfo {
	c.lock.Lock()
	defer c.lock.Unlock()

	groupInfo, exists := c.cache[groupID]
	if !exists {
		groupInfo = &GroupMembersInfo{
			initialized: false,
			members:     make(map[string]struct{}),
			lock:        &sync.Mutex{},
		}
		c.cache[groupID] = groupInfo
	}
	return groupInfo
}

func (c *GroupCache) GetMembers(api GroupsAPI, groupID string) (map[string]struct{}, error) {
	groupInfo := c.getOrCreateGroupInfo(groupID)
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
	groupInfo := c.getOrCreateGroupInfo(groupID)
	groupInfo.lock.Lock()
	defer groupInfo.lock.Unlock()

	err := api.Patch(groupID, PatchRequest(
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
	groupInfo := c.getOrCreateGroupInfo(groupID)
	groupInfo.lock.Lock()
	defer groupInfo.lock.Unlock()

	err := api.Patch(groupID, PatchRequestWithValue("add", "members", memberID))
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
	memberKey := strings.ToLower(memberID)
	_, ok := members[memberKey]
	return ok
}

var globalGroupsCache = NewGroupsCache()

// ResourceGroupMember bind group with member
func ResourceGroupMember() common.Resource {
	return common.NewPairID("group_id", "member_id").BindResource(common.BindResource{
		CreateContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			return globalGroupsCache.addMember(NewGroupsAPI(ctx, c), groupID, memberID)
		},
		ReadContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			members, err := globalGroupsCache.GetMembers(NewGroupsAPI(ctx, c), groupID)
			if err == nil && !hasMember(members, memberID) {
				return &apierr.APIError{
					ErrorCode:  "NOT_FOUND",
					StatusCode: 404,
					Message:    "Group has no member",
				}
			}
			return err
		},
		DeleteContext: func(ctx context.Context, groupID, memberID string, c *common.DatabricksClient) error {
			return globalGroupsCache.removeMember(NewGroupsAPI(ctx, c), groupID, memberID)
		},
	})
}
