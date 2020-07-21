package service

import (
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// GroupsAPI exposes the scim groups API
type GroupsAPI struct {
	client *DatabricksClient
}

// Create creates a scim group in the Databricks workspace
func (a GroupsAPI) Create(groupName string, members []string, roles []string, entitlements []string) (group model.Group, err error) {
	scimGroupRequest := struct {
		Schemas      []model.URN           `json:"schemas,omitempty"`
		DisplayName  string                `json:"displayName,omitempty"`
		Members      []model.ValueListItem `json:"members,omitempty"`
		Entitlements []model.ValueListItem `json:"entitlements,omitempty"`
		Roles        []model.ValueListItem `json:"roles,omitempty"`
	}{}
	scimGroupRequest.Schemas = []model.URN{model.GroupSchema}
	scimGroupRequest.DisplayName = groupName

	scimGroupRequest.Members = []model.ValueListItem{}
	for _, member := range members {
		scimGroupRequest.Members = append(scimGroupRequest.Members, model.ValueListItem{Value: member})
	}

	scimGroupRequest.Roles = []model.ValueListItem{}
	for _, role := range roles {
		scimGroupRequest.Roles = append(scimGroupRequest.Roles, model.ValueListItem{Value: role})
	}

	scimGroupRequest.Entitlements = []model.ValueListItem{}
	for _, entitlement := range entitlements {
		scimGroupRequest.Entitlements = append(scimGroupRequest.Entitlements, model.ValueListItem{Value: entitlement})
	}
	err = a.client.performScim(http.MethodPost, "/preview/scim/v2/Groups", scimGroupRequest, &group)
	return
}

// Read reads and returns a Group object via SCIM api
func (a GroupsAPI) Read(groupID string) (group model.Group, err error) {
	err = a.scimClient.client.get(fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID), nil, &group)
	if err != nil {
		return
	}
	//get inherited groups
	var groups []model.Group
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

// GetAdminGroup returns the admin group in a given workspace by fetching with query "displayName+eq+admins"
func (a GroupsAPI) GetAdminGroup() (model.Group, error) {
	var group model.Group
	var groups model.GroupList
	err := a.client.performScim(http.MethodGet, "/preview/scim/v2/Groups", map[string]string{
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

// Patch applys a patch request for a group given a path attribute
func (a GroupsAPI) Patch(groupID string, addList []string, removeList []string, path model.GroupPathType) error {
	groupPath := fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID)

	var addOperations model.GroupPatchOperations
	var removeOperations model.GroupPatchOperations

	groupPatchRequest := model.GroupPatchRequest{
		Schemas:    []model.URN{model.PatchOp},
		Operations: []model.GroupPatchOperations{},
	}

	if addList == nil && removeList == nil {
		return errors.New("empty members list to add or to remove")
	}

	if len(addList) > 0 {
		addOperations = model.GroupPatchOperations{
			Op:    "add",
			Path:  path,
			Value: []model.ValueListItem{},
		}
		for _, addItem := range addList {
			addOperations.Value = append(addOperations.Value, model.ValueListItem{Value: addItem})
		}
		groupPatchRequest.Operations = append(groupPatchRequest.Operations, addOperations)
	}

	for _, removeItem := range removeList {
		path := fmt.Sprintf("%s[value eq \"%s\"]", string(path), removeItem)
		removeOperations = model.GroupPatchOperations{
			Op:   "remove",
			Path: model.GroupPathType(path),
		}
		groupPatchRequest.Operations = append(groupPatchRequest.Operations, removeOperations)
	}
	return a.client.performScim(http.MethodPatch, groupPath, groupPatchRequest)
}

// Delete deletes a group given a group id
func (a GroupsAPI) Delete(groupID string) error {
	return a.client.performScim(http.MethodDelete,
		fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID), nil)
}

func (a GroupsAPI) getInheritedAndNonInheritedRoles(
	group model.Group, groups []model.Group) (inherited []model.RoleListItem,
	unInherited []model.RoleListItem) {
	allRoles := group.Roles
	var inheritedRoles []model.RoleListItem
	inheritedRolesKeys := []string{}
	inheritedRolesMap := map[string]model.RoleListItem{}
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
