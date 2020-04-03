package service

import (
	"encoding/json"
	"fmt"
	"github.com/stikkireddy/databricks-tf-provider/client/model"
	"net/http"
)

// TokensAPI exposes the Secrets API
type UsersAPI struct {
	Client DBApiClient
}

func (a UsersAPI) init(client DBApiClient) UsersAPI {
	a.Client = client
	return a
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
	_, err := a.Client.performQuery(http.MethodPut, userPath, "2.0", scimHeaders, scimUserUpdateRequest)
	return err
}

func (a UsersAPI) Delete(userId string) error {

	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userId)

	_, err := a.Client.performQuery(http.MethodDelete, userPath, "2.0", scimHeaders, nil)
	return err
}
