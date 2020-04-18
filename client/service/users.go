package service

import (
	"encoding/json"
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"net/http"
	"sort"
)

// UsersAPI exposes the scim user API
type UsersAPI struct {
	Client DBApiClient
}

func (a UsersAPI) Create(userName string, displayName string, entitlements []string, roles []string) (model.User, error) {
	var user model.User
	scimUserRequest := struct {
		Schemas      []model.URN                  `json:"schemas,omitempty"`
		UserName     string                       `json:"userName,omitempty"`
		Entitlements []model.EntitlementsListItem `json:"entitlements,omitempty"`
		DisplayName  string                       `json:"displayName,omitempty"`
		Roles        []model.RoleListItem         `json:"roles,omitempty"`
	}{}
	scimUserRequest.Schemas = []model.URN{model.UserSchema}
	scimUserRequest.UserName = userName
	scimUserRequest.DisplayName = displayName
	scimUserRequest.Entitlements = []model.EntitlementsListItem{}
	for _, entitlement := range entitlements {
		scimUserRequest.Entitlements = append(scimUserRequest.Entitlements, model.EntitlementsListItem{Value: model.Entitlement(entitlement)})
	}
	scimUserRequest.Roles = []model.RoleListItem{}
	for _, role := range roles {
		scimUserRequest.Roles = append(scimUserRequest.Roles, model.RoleListItem{Value: role})
	}

	resp, err := a.Client.performQuery(http.MethodPost, "/preview/scim/v2/Users", "2.0", scimHeaders, scimUserRequest)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)
	return user, err
}

func (a UsersAPI) Read(userId string) (model.User, error) {
	var user model.User
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userId)

	resp, err := a.Client.performQuery(http.MethodGet, userPath, "2.0", scimHeaders, nil)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	//get groups
	var groups []model.Group
	for _, group := range user.Groups {
		group, err := a.Client.Groups().Read(group.Value)
		if err != nil {
			return user, err
		}
		groups = append(groups, group)
	}
	inherited, unInherited, err := a.getInheritedAndNonInheritedRoles(user, groups)
	user.InheritedRoles = inherited
	user.UnInheritedRoles = unInherited
	return user, err
}

func (a UsersAPI) Update(userId string, userName string, displayName string, entitlements []string, roles []string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userId)
	scimUserUpdateRequest := struct {
		Schemas      []model.URN                  `json:"schemas,omitempty"`
		UserName     string                       `json:"userName,omitempty"`
		Entitlements []model.EntitlementsListItem `json:"entitlements,omitempty"`
		DisplayName  string                       `json:"displayName,omitempty"`
		Roles        []model.RoleListItem         `json:"roles,omitempty"`
		Groups       []model.GroupsListItem       `json:"groups,omitempty"`
	}{}
	scimUserUpdateRequest.Schemas = []model.URN{model.UserSchema}
	scimUserUpdateRequest.UserName = userName
	scimUserUpdateRequest.DisplayName = displayName
	scimUserUpdateRequest.Entitlements = []model.EntitlementsListItem{}
	for _, entitlement := range entitlements {
		scimUserUpdateRequest.Entitlements = append(scimUserUpdateRequest.Entitlements, model.EntitlementsListItem{Value: model.Entitlement(entitlement)})
	}
	scimUserUpdateRequest.Roles = []model.RoleListItem{}
	for _, role := range roles {
		scimUserUpdateRequest.Roles = append(scimUserUpdateRequest.Roles, model.RoleListItem{Value: role})
	}
	//Get any existing groups that the user is part of
	user, err := a.Read(userId)
	if err != nil {
		return err
	}
	scimUserUpdateRequest.Groups = user.Groups
	_, err = a.Client.performQuery(http.MethodPut, userPath, "2.0", scimHeaders, scimUserUpdateRequest)
	return err
}

func (a UsersAPI) Delete(userId string) error {

	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userId)

	_, err := a.Client.performQuery(http.MethodDelete, userPath, "2.0", scimHeaders, nil)
	return err
}

func (a UsersAPI) SetUserAsAdmin(userId string, adminGroupId string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userId)

	var addOperations model.UserPatchOperations

	userPatchRequest := model.UserPatchRequest{
		Schemas:    []model.URN{model.PatchOp},
		Operations: []model.UserPatchOperations{},
	}

	addOperations = model.UserPatchOperations{
		Op: "add",
		Value: &model.GroupsValue{
			Groups: []model.ValueListItem{model.ValueListItem{Value: adminGroupId}},
		},
	}
	userPatchRequest.Operations = append(userPatchRequest.Operations, addOperations)

	_, err := a.Client.performQuery(http.MethodPatch, userPath, "2.0", scimHeaders, userPatchRequest)

	return err
}

func (a UsersAPI) VerifyUserAsAdmin(userId string, adminGroupId string) (bool, error) {
	user, err := a.Read(userId)
	if err != nil {
		return false, err
	}
	for _, group := range user.Groups {
		if group.Value == adminGroupId {
			return true, nil
		}
	}

	return false, nil
}

func (a UsersAPI) RemoveUserAsAdmin(userId string, adminGroupId string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userId)

	var removeOperations model.UserPatchOperations

	userPatchRequest := model.UserPatchRequest{
		Schemas:    []model.URN{model.PatchOp},
		Operations: []model.UserPatchOperations{},
	}

	path := fmt.Sprintf("groups[value eq \"%s\"]", adminGroupId)
	removeOperations = model.UserPatchOperations{
		Op:   "remove",
		Path: path,
	}
	userPatchRequest.Operations = append(userPatchRequest.Operations, removeOperations)

	_, err := a.Client.performQuery(http.MethodPatch, userPath, "2.0", scimHeaders, userPatchRequest)

	return err
}

func (a UsersAPI) getInheritedAndNonInheritedRoles(user model.User, groups []model.Group) (inherited []model.RoleListItem, unInherited []model.RoleListItem, err error) {
	allRoles := user.Roles
	var inheritedRoles []model.RoleListItem
	inheritedRolesKeys := []string{}
	inheritedRolesMap := map[string]model.RoleListItem{}
	for _, group := range groups {
		for _, role := range group.Roles {
			inheritedRoles = append(inheritedRoles, role)
		}
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
	return inherited, unInherited, nil
}
