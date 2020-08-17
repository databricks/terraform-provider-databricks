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

// UserManagement abstracts user management
type UserManagement interface {
	Create(userName string, displayName string, entitlements []string, roles []string) (User, error)
	Read(userID string) (User, error)
	Update(userID string, userName string, displayName string, entitlements []string, roles []string) error
	Delete(userID string) error
	Me() (User, error)
}

// CommandMock mocks the execution of command
type CommandMock func(commandStr string) (string, error)

// CommandExecutorMock simplifies command testing
type commandExecutorMock struct {
	mock CommandMock
}

// Execute mock command with given mock function
func (c commandExecutorMock) Execute(clusterID, language, commandStr string) (string, error) {
	return c.mock(commandStr)
}

// UsersAPI exposes the scim user API
type UsersAPI struct {
	C *common.DatabricksClient
}

// ResourceUser entity from which resource schema is made
type ResourceUser struct {
	UserName       string   `json:"user_name"`
	DisplayName    string   `json:"display_name,omitempty"`
	Roles          []string `json:"roles,omitempty" tf:"slice_set"`
	Entitlements   []string `json:"entitlements,omitempty" tf:"slice_set"`
	DefaultRoles   []string `json:"default_roles" tf:"slice_set"`
	InheritedRoles []string `json:"inherited_roles,omitempty" tf:"slice_set,computed"`
	SetAdmin       bool     `json:"set_admin,omitempty"`
}

type scimUserRequest struct {
	Schemas      []URN                  `json:"schemas,omitempty"`
	UserName     string                 `json:"userName,omitempty"`
	Entitlements []EntitlementsListItem `json:"entitlements,omitempty"`
	DisplayName  string                 `json:"displayName,omitempty"`
	Roles        []RoleListItem         `json:"roles,omitempty"`
	Groups       []GroupsListItem       `json:"groups,omitempty"`
}

func (u ResourceUser) toRequest() scimUserRequest {
	entitlements := []EntitlementsListItem{}
	for _, entitlement := range u.Entitlements {
		entitlements = append(entitlements, EntitlementsListItem{Value: Entitlement(entitlement)})
	}
	roles := []RoleListItem{}
	for _, role := range u.Roles {
		roles = append(roles, RoleListItem{Value: role})
	}
	return scimUserRequest{
		Schemas:      []URN{UserSchema},
		UserName:     u.UserName,
		DisplayName:  u.DisplayName,
		Entitlements: entitlements,
		Roles:        roles,
	}
}

// CreateR ..
func (a UsersAPI) CreateR(ru ResourceUser) (user User, err error) {
	err = a.C.Scim(http.MethodPost, "/preview/scim/v2/Users", ru.toRequest(), &user)
	return user, err
}

// Create given a username, displayname, entitlements, and roles will create a scim user via SCIM api
func (a UsersAPI) Create(userName string, displayName string, entitlements []string, roles []string) (User, error) {
	var user User
	scimUserRequest := struct {
		Schemas      []URN                  `json:"schemas,omitempty"`
		UserName     string                 `json:"userName,omitempty"`
		Entitlements []EntitlementsListItem `json:"entitlements,omitempty"`
		DisplayName  string                 `json:"displayName,omitempty"`
		Roles        []RoleListItem         `json:"roles,omitempty"`
	}{}
	scimUserRequest.Schemas = []URN{UserSchema}
	scimUserRequest.UserName = userName
	scimUserRequest.DisplayName = displayName
	scimUserRequest.Entitlements = []EntitlementsListItem{}
	for _, entitlement := range entitlements {
		scimUserRequest.Entitlements = append(scimUserRequest.Entitlements, EntitlementsListItem{Value: Entitlement(entitlement)})
	}
	scimUserRequest.Roles = []RoleListItem{}
	for _, role := range roles {
		scimUserRequest.Roles = append(scimUserRequest.Roles, RoleListItem{Value: role})
	}
	err := a.C.Scim(http.MethodPost, "/preview/scim/v2/Users", scimUserRequest, &user)
	return user, err
}

// Read returns the user object and all the attributes of a scim user
func (a UsersAPI) Read(userID string) (User, error) {
	user, err := a.read(userID)
	if err != nil {
		return user, err
	}

	//get groups
	var groups []Group
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

func (a UsersAPI) read(userID string) (User, error) {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	return a.readByPath(userPath)
}

// Me gets user information about caller
func (a UsersAPI) Me() (User, error) {
	return a.readByPath("/preview/scim/v2/Me")
}

func (a UsersAPI) readByPath(userPath string) (User, error) {
	var user User
	err := a.C.Scim(http.MethodGet, userPath, nil, &user)
	// TODO: add fetching of nested groups
	return user, err
}

// UpdateR ...
func (a UsersAPI) UpdateR(userID string, ru ResourceUser) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	scimUserUpdateRequest := ru.toRequest()
	//Get any existing groups that the user is part of
	user, err := a.read(userID)
	if err != nil {
		return err
	}
	scimUserUpdateRequest.Groups = user.Groups
	return a.C.Scim(http.MethodPut, userPath, scimUserUpdateRequest, nil)
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

// SetUserAsAdmin will add the user to a admin group given the admin group id and user id
func (a UsersAPI) SetUserAsAdmin(userID string, adminGroupID string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	userPatchRequest := UserPatchRequest{}
	userPatchRequest.GroupOperation("add", adminGroupID)
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
	userPatchRequest := UserPatchRequest{}
	userPatchRequest.GroupOperation("remove", adminGroupID)
	return a.C.Scim(http.MethodPatch, userPath, userPatchRequest, nil)
}

// GetOrCreateDefaultMetaUser ...
func (a UsersAPI) GetOrCreateDefaultMetaUser(metaUserDisplayName string, metaUserName string, deleteAfterCreate bool) (user User, err error) {
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
		return User{}, errors.New("more than one meta user")
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

func (a UsersAPI) getInheritedAndNonInheritedRoles(user User, groups []Group) (inherited []RoleListItem, unInherited []RoleListItem) {
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
