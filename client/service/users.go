package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// UsersAPI exposes the scim user API
type UsersAPI struct {
	client *DatabricksClient
}

// Create given a username, displayname, entitlements, and roles will create a scim user via SCIM api
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
	err := a.client.performScim(http.MethodPost, "/preview/scim/v2/Users", scimUserRequest, &user)
	return user, err
}

// Read returns the user object and all the attributes of a scim user
func (a UsersAPI) Read(userID string) (model.User, error) {
	user, err := a.read(userID)
	if err != nil {
		return user, err
	}

	//get groups
	var groups []model.Group
	for _, group := range user.Groups {
		group, err := a.client.Groups().Read(group.Value)
		if err != nil {
			return user, err
		}
		groups = append(groups, group)
	}
	inherited, unInherited := a.getInheritedAndNonInheritedRoles(user, groups)
	user.InheritedRoles = inherited
	user.UnInheritedRoles = unInherited
	return user, err
}

func (a UsersAPI) read(userID string) (model.User, error) {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	return a.readByPath(userPath)
}

// Me gets user information about caller
func (a UsersAPI) Me() (model.User, error) {
	return a.readByPath("/preview/scim/v2/Me")
}

func (a UsersAPI) readByPath(userPath string) (model.User, error) {
	var user model.User
	err := a.client.performScim(http.MethodGet, userPath, nil, &user)
	return user, err
}

// Update will update the user given the user id, username, display name, entitlements and roles
func (a UsersAPI) Update(userID string, userName string, displayName string, entitlements []string, roles []string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
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
	user, err := a.read(userID)
	if err != nil {
		return err
	}
	scimUserUpdateRequest.Groups = user.Groups
	return a.client.performScim(http.MethodPut, userPath, scimUserUpdateRequest, nil)
}

// Delete will delete the user given the user id
func (a UsersAPI) Delete(userID string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	return a.client.performScim(http.MethodDelete, userPath, nil, nil)
}

// SetUserAsAdmin will add the user to a admin group given the admin group id and user id
func (a UsersAPI) SetUserAsAdmin(userID string, adminGroupID string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	var addOperations model.UserPatchOperations
	userPatchRequest := model.UserPatchRequest{
		Schemas:    []model.URN{model.PatchOp},
		Operations: []model.UserPatchOperations{},
	}
	addOperations = model.UserPatchOperations{
		Op: "add",
		Value: &model.GroupsValue{
			Groups: []model.ValueListItem{{Value: adminGroupID}},
		},
	}
	userPatchRequest.Operations = append(userPatchRequest.Operations, addOperations)
	return a.client.performScim(http.MethodPatch, userPath, userPatchRequest, nil)
}

// VerifyUserAsAdmin will verify the user belongs to the admin group given the admin group id and user id
func (a UsersAPI) VerifyUserAsAdmin(userID string, adminGroupID string) (bool, error) {
	user, err := a.read(userID)
	if err != nil {
		return false, err
	}
	for _, group := range user.Groups {
		if group.Value == adminGroupID {
			return true, nil
		}
	}
	return false, nil
}

// RemoveUserAsAdmin will remove the user from the admin group given the admin group id and user id
func (a UsersAPI) RemoveUserAsAdmin(userID string, adminGroupID string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	var removeOperations model.UserPatchOperations
	userPatchRequest := model.UserPatchRequest{
		Schemas:    []model.URN{model.PatchOp},
		Operations: []model.UserPatchOperations{},
	}
	path := fmt.Sprintf("groups[value eq \"%s\"]", adminGroupID)
	removeOperations = model.UserPatchOperations{
		Op:   "remove",
		Path: path,
	}
	userPatchRequest.Operations = append(userPatchRequest.Operations, removeOperations)
	return a.client.performScim(http.MethodPatch, userPath, userPatchRequest, nil)
}

// GetOrCreateDefaultMetaUser ...
func (a UsersAPI) GetOrCreateDefaultMetaUser(metaUserDisplayName string, metaUserName string, deleteAfterCreate bool) (user model.User, err error) {
	var users model.UserList
	err = a.client.performScim(http.MethodGet, "/preview/scim/v2/Users", map[string]string{
		"filter": "displayName+eq+" + metaUserDisplayName,
	}, &users)
	if err != nil {
		return user, err
	}
	resources := users.Resources
	if len(resources) == 1 {
		return resources[0], err
	} else if len(resources) > 1 {
		return model.User{}, errors.New("more than one meta user")
	}

	log.Printf("Meta User not found will create a new meta user with name: %s\n", metaUserDisplayName)

	newCreatedUser, err := a.Create(metaUserName, metaUserDisplayName, nil, nil)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			time.Sleep(time.Second * 1)
			return a.GetOrCreateDefaultMetaUser(metaUserDisplayName, metaUserName, deleteAfterCreate)
		}
		return user, err
	}
	if deleteAfterCreate {
		defer func() {
			deferErr := a.Delete(newCreatedUser.ID)
			err = deferErr
		}()
	}
	return newCreatedUser, err
}

func (a UsersAPI) getInheritedAndNonInheritedRoles(user model.User, groups []model.Group) (inherited []model.RoleListItem, unInherited []model.RoleListItem) {
	allRoles := user.Roles
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
