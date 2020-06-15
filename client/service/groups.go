package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// GroupsAPI exposes the scim groups API
type GroupsAPI struct {
	Client *DBApiClient
}

// Create creates a scim group in the Databricks workspace
func (a GroupsAPI) Create(groupName string, members []string, roles []string, entitlements []string) (model.Group, error) {
	var group model.Group

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

	resp, err := a.Client.performQuery(http.MethodPost, "/preview/scim/v2/Groups", "2.0", scimHeaders, scimGroupRequest, nil)
	if err != nil {
		return group, err
	}

	err = json.Unmarshal(resp, &group)
	return group, err
}

// Read reads and returns a Group object via SCIM api
func (a GroupsAPI) Read(groupID string) (model.Group, error) {
	var group model.Group
	groupPath := fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID)

	resp, err := a.Client.performQuery(http.MethodGet, groupPath, "2.0", scimHeaders, nil, nil)
	if err != nil {
		return group, err
	}

	err = json.Unmarshal(resp, &group)
	if err != nil {
		return group, err
	}
	log.Println(group)
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

	return group, err
}

// GetAdminGroup returns the admin group in a given workspace by fetching with query "displayName+eq+admins"
func (a GroupsAPI) GetAdminGroup() (model.Group, error) {
	var group model.Group
	var groups model.GroupList

	adminsQuery := "/preview/scim/v2/Groups?filter=displayName+eq+admins"

	resp, err := a.Client.performQuery(http.MethodGet, adminsQuery, "2.0", scimHeaders, nil, nil)
	if err != nil {
		return group, err
	}
	err = json.Unmarshal(resp, &groups)

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

	_, err := a.Client.performQuery(http.MethodPatch, groupPath, "2.0", scimHeaders, groupPatchRequest, nil)

	return err
}

// Delete deletes a group given a group id
func (a GroupsAPI) Delete(groupID string) error {
	groupPath := fmt.Sprintf("/preview/scim/v2/Groups/%v", groupID)
	_, err := a.Client.performQuery(http.MethodDelete, groupPath, "2.0", scimHeaders, nil, nil)
	return err
}

func (a GroupsAPI) getInheritedAndNonInheritedRoles(group model.Group, groups []model.Group) (inherited []model.RoleListItem, unInherited []model.RoleListItem) {
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
