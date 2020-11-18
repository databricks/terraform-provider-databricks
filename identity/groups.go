package identity

import (
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/databrickslabs/databricks-terraform/common"
)

// NewGroupsAPI creates GroupsAPI instance from provider meta
func NewGroupsAPI(m interface{}) GroupsAPI {
	return GroupsAPI{client: m.(*common.DatabricksClient)}
}

// GroupsAPI exposes the scim groups API
type GroupsAPI struct {
	client *common.DatabricksClient
}

// Create creates a scim group in the Databricks workspace
func (a GroupsAPI) Create(groupName string, members []string, roles []string, entitlements []string) (group ScimGroup, err error) {
	scimGroupRequest := struct {
		Schemas      []URN           `json:"schemas,omitempty"`
		DisplayName  string          `json:"displayName,omitempty"`
		Members      []ValueListItem `json:"members,omitempty"`
		Entitlements []ValueListItem `json:"entitlements,omitempty"`
		Roles        []ValueListItem `json:"roles,omitempty"`
	}{}
	scimGroupRequest.Schemas = []URN{GroupSchema}
	scimGroupRequest.DisplayName = groupName

	scimGroupRequest.Members = []ValueListItem{}
	for _, member := range members {
		scimGroupRequest.Members = append(scimGroupRequest.Members, ValueListItem{Value: member})
	}

	scimGroupRequest.Roles = []ValueListItem{}
	for _, role := range roles {
		scimGroupRequest.Roles = append(scimGroupRequest.Roles, ValueListItem{Value: role})
	}

	scimGroupRequest.Entitlements = []ValueListItem{}
	for _, entitlement := range entitlements {
		scimGroupRequest.Entitlements = append(scimGroupRequest.Entitlements, ValueListItem{Value: entitlement})
	}
	err = a.client.Scim(http.MethodPost, "/preview/scim/v2/Groups", scimGroupRequest, &group)
	return
}

// Read reads and returns a Group object via SCIM api
func (a GroupsAPI) Read(groupID string) (group ScimGroup, err error) {
	err = a.client.Scim(http.MethodGet, fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID), nil, &group)
	if err != nil {
		return
	}
	//get inherited groups
	var groups []ScimGroup
	for _, inheritedGroup := range group.Groups {
		inheritedGroupFull, err := a.Read(inheritedGroup.Value)
		if err != nil {
			return group, err
		}
		groups = append(groups, inheritedGroupFull)
	}
	inherited, unInherited := a.getInheritedAndNonInheritedRoles(group, groups)
	group.InheritedRoles = inherited
	group.UnInheritedRoles = unInherited
	return
}

// Filter returns groups matching the filter
func (a GroupsAPI) Filter(filter string) (GroupList, error) {
	var groups GroupList
	req := map[string]string{}
	if filter != "" {
		req["filter"] = filter
	}
	err := a.client.Scim(http.MethodGet, "/preview/scim/v2/Groups", req, &groups)
	return groups, err
}

// GetAdminGroup returns the admin group in a given workspace by fetching with query "displayName+eq+admins"
func (a GroupsAPI) GetAdminGroup() (ScimGroup, error) {
	var group ScimGroup
	var groups GroupList
	err := a.client.Scim(http.MethodGet, "/preview/scim/v2/Groups", map[string]string{
		"filter": "displayName+eq+admins",
	}, &groups)
	if err != nil {
		return group, err
	}
	resources := groups.Resources
	if len(resources) == 1 {
		return resources[0], err
	}
	return group, errors.New("Unable to identify the admin group! ")
}

func (a GroupsAPI) PatchR(groupID string, r patchRequest) error {
	return a.client.Scim(http.MethodPatch, fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID), r, nil)
}

// Patch applys a patch request for a group given a path attribute
func (a GroupsAPI) Patch(groupID string, addList []string, removeList []string, path GroupPathType) error {
	groupPath := fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID)

	var addOperations GroupPatchOperations
	var removeOperations GroupPatchOperations

	groupPatchRequest := GroupPatchRequest{
		Schemas:    []URN{PatchOp},
		Operations: []GroupPatchOperations{},
	}

	if addList == nil && removeList == nil {
		return errors.New("empty members list to add or to remove")
	}

	if len(addList) > 0 {
		addOperations = GroupPatchOperations{
			Op:    "add",
			Path:  path,
			Value: []ValueListItem{},
		}
		for _, addItem := range addList {
			addOperations.Value = append(addOperations.Value, ValueListItem{Value: addItem})
		}
		groupPatchRequest.Operations = append(groupPatchRequest.Operations, addOperations)
	}

	for _, removeItem := range removeList {
		path := fmt.Sprintf("%s[value eq \"%s\"]", string(path), removeItem)
		removeOperations = GroupPatchOperations{
			Op:   "remove",
			Path: GroupPathType(path),
		}
		groupPatchRequest.Operations = append(groupPatchRequest.Operations, removeOperations)
	}
	return a.client.Scim(http.MethodPatch, groupPath, groupPatchRequest, nil)
}

// Delete deletes a group given a group id
func (a GroupsAPI) Delete(groupID string) error {
	return a.client.Scim(http.MethodDelete,
		fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID),
		nil, nil)
}

func (a GroupsAPI) getInheritedAndNonInheritedRoles(
	group ScimGroup, groups []ScimGroup) (inherited []RoleListItem,
	unInherited []RoleListItem) {
	allRoles := group.Roles
	var inheritedRoles []RoleListItem
	inheritedRolesKeys := []string{}
	inheritedRolesMap := map[string]RoleListItem{}
	for _, group := range groups {
		inheritedRoles = append(inheritedRoles, group.Roles...)
	}
	for _, role := range inheritedRoles {
		inheritedRolesKeys = append(inheritedRolesKeys, role.Value)
		inheritedRolesMap[role.Value] = role
	}
	sort.Strings(inheritedRolesKeys)
	for _, roleKey := range inheritedRolesKeys {
		inherited = append(inherited, inheritedRolesMap[roleKey])
	}
	for _, role := range allRoles {
		if _, ok := inheritedRolesMap[role.Value]; !ok {
			unInherited = append(unInherited, role)
		}
	}
	return inherited, unInherited
}
