package identity

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
)

// NewUsersAPI creates UsersAPI instance from provider meta
func NewUsersAPI(m interface{}) UsersAPI {
	return UsersAPI{C: m.(*common.DatabricksClient)}
}

// UsersAPI exposes the scim user API
type UsersAPI struct {
	C *common.DatabricksClient
}

// UserEntity entity from which resource schema is made
type UserEntity struct {
	UserName                string `json:"user_name"`
	DisplayName             string `json:"display_name,omitempty"`
	AllowClusterCreate      bool   `json:"allow_cluster_create,omitempty"`
	AllowInstancePoolCreate bool   `json:"allow_instance_pool_create,omitempty"`
}

func (u UserEntity) toRequest() ScimUser {
	entitlements := []EntitlementsListItem{}
	if u.AllowClusterCreate {
		entitlements = append(entitlements, EntitlementsListItem{
			Value: Entitlement("allow-cluster-create"),
		})
	}
	if u.AllowInstancePoolCreate {
		entitlements = append(entitlements, EntitlementsListItem{
			Value: Entitlement("allow-instance-pool-create"),
		})
	}
	return ScimUser{
		Schemas:      []URN{UserSchema},
		UserName:     u.UserName,
		DisplayName:  u.DisplayName,
		Entitlements: entitlements,
	}
}

// CreateR ..
func (a UsersAPI) CreateR(ru UserEntity) (user ScimUser, err error) {
	err = a.C.Scim(http.MethodPost, "/preview/scim/v2/Users", ru.toRequest(), &user)
	return user, err
}

// Create given a username, displayname, entitlements, and roles will create a scim user via SCIM api
func (a UsersAPI) Create(userName string, displayName string, entitlements []string, roles []string) (ScimUser, error) {
	var user ScimUser
	createRequest := ScimUser{
		Schemas:      []URN{UserSchema},
		UserName:     userName,
		DisplayName:  displayName,
		Entitlements: []EntitlementsListItem{},
		Roles:        []RoleListItem{},
	}
	for _, entitlement := range entitlements {
		createRequest.Entitlements = append(createRequest.Entitlements, EntitlementsListItem{Value: Entitlement(entitlement)})
	}
	for _, role := range roles {
		createRequest.Roles = append(createRequest.Roles, RoleListItem{Value: role})
	}
	err := a.C.Scim(http.MethodPost, "/preview/scim/v2/Users", createRequest, &user)
	return user, err
}

func (a UsersAPI) ReadR(userID string) (ru UserEntity, err error) {
	user, err := a.read(userID)
	if err != nil {
		return
	}
	ru.UserName = user.UserName
	ru.DisplayName = user.DisplayName
	for _, ent := range user.Entitlements {
		switch ent.Value {
		case AllowClusterCreateEntitlement:
			ru.AllowClusterCreate = true
		case AllowInstancePoolCreateEntitlement:
			ru.AllowInstancePoolCreate = true
		}
	}
	return
}

// Read returns the user object and all the attributes of a scim user
func (a UsersAPI) Read(userID string) (ScimUser, error) {
	user, err := a.read(userID)
	if err != nil {
		return user, err
	}

	//get groups
	var groups []ScimGroup
	for _, group := range user.Groups {
		group, err := GroupsAPI{a.C}.Read(group.Value)
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

func (a UsersAPI) read(userID string) (ScimUser, error) {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	return a.readByPath(userPath)
}

// Me gets user information about caller
func (a UsersAPI) Me() (ScimUser, error) {
	return a.readByPath("/preview/scim/v2/Me")
}

func (a UsersAPI) readByPath(userPath string) (user ScimUser, err error) {
	err = a.C.Scim(http.MethodGet, userPath, nil, &user)
	return
}

func (a UsersAPI) UpdateR(userID string, ru UserEntity) error {
	user, err := a.read(userID)
	if err != nil {
		return err
	}
	updateRequest := ru.toRequest()
	updateRequest.Groups = user.Groups
	updateRequest.Roles = user.Roles
	return a.C.Scim(http.MethodPut,
		fmt.Sprintf("/preview/scim/v2/Users/%v", userID),
		updateRequest, nil)
}

// Update will update the user given the user id, username, display name, entitlements and roles
func (a UsersAPI) Update(userID string, userName string, displayName string, entitlements []string, roles []string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	scimUserUpdateRequest := struct {
		Schemas      []URN                  `json:"schemas,omitempty"`
		UserName     string                 `json:"userName,omitempty"`
		Entitlements []EntitlementsListItem `json:"entitlements,omitempty"`
		DisplayName  string                 `json:"displayName,omitempty"`
		Roles        []RoleListItem         `json:"roles,omitempty"`
		Groups       []GroupsListItem       `json:"groups,omitempty"`
	}{}
	scimUserUpdateRequest.Schemas = []URN{UserSchema}
	scimUserUpdateRequest.UserName = userName
	scimUserUpdateRequest.DisplayName = displayName
	scimUserUpdateRequest.Entitlements = []EntitlementsListItem{}
	for _, entitlement := range entitlements {
		scimUserUpdateRequest.Entitlements = append(scimUserUpdateRequest.Entitlements, EntitlementsListItem{Value: Entitlement(entitlement)})
	}
	scimUserUpdateRequest.Roles = []RoleListItem{}
	for _, role := range roles {
		scimUserUpdateRequest.Roles = append(scimUserUpdateRequest.Roles, RoleListItem{Value: role})
	}
	//Get any existing groups that the user is part of
	user, err := a.read(userID)
	if err != nil {
		return err
	}
	scimUserUpdateRequest.Groups = user.Groups
	return a.C.Scim(http.MethodPut, userPath, scimUserUpdateRequest, nil)
}

// Delete will delete the user given the user id
func (a UsersAPI) Delete(userID string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	return a.C.Scim(http.MethodDelete, userPath, nil, nil)
}

// Deprecated SetUserAsAdmin will add the user to a admin group given the admin group id and user id
func (a UsersAPI) SetUserAsAdmin(userID string, adminGroupID string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	var addOperations UserPatchOperations
	userPatchRequest := UserPatchRequest{
		Schemas:    []URN{PatchOp},
		Operations: []UserPatchOperations{},
	}
	addOperations = UserPatchOperations{
		Op: "add",
		Value: &GroupsValue{
			Groups: []ValueListItem{{Value: adminGroupID}},
		},
	}
	userPatchRequest.Operations = append(userPatchRequest.Operations, addOperations)
	return a.C.Scim(http.MethodPatch, userPath, userPatchRequest, nil)
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
	var removeOperations UserPatchOperations
	userPatchRequest := UserPatchRequest{
		Schemas:    []URN{PatchOp},
		Operations: []UserPatchOperations{},
	}
	path := fmt.Sprintf("groups[value eq \"%s\"]", adminGroupID)
	removeOperations = UserPatchOperations{
		Op:   "remove",
		Path: path,
	}
	userPatchRequest.Operations = append(userPatchRequest.Operations, removeOperations)
	return a.C.Scim(http.MethodPatch, userPath, userPatchRequest, nil)
}

// GetOrCreateDefaultMetaUser ...
func (a UsersAPI) GetOrCreateDefaultMetaUser(metaUserDisplayName string, metaUserName string, deleteAfterCreate bool) (user ScimUser, err error) {
	var users UserList
	err = a.C.Scim(http.MethodGet, "/preview/scim/v2/Users", map[string]string{
		"filter": "displayName+eq+" + metaUserDisplayName,
	}, &users)
	if err != nil {
		return user, err
	}
	resources := users.Resources
	if len(resources) == 1 {
		return resources[0], err
	} else if len(resources) > 1 {
		return ScimUser{}, errors.New("more than one meta user")
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

func (a UsersAPI) getInheritedAndNonInheritedRoles(user ScimUser, groups []ScimGroup) (inherited []RoleListItem, unInherited []RoleListItem) {
	allRoles := user.Roles
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
