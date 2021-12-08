package scim

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databrickslabs/terraform-provider-databricks/common"
)

// NewUsersAPI creates UsersAPI instance from provider meta
func NewUsersAPI(ctx context.Context, m interface{}) UsersAPI {
	return UsersAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// UsersAPI exposes the scim user API
type UsersAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create user in the backend
func (a UsersAPI) Create(ru User) (user User, err error) {
	if ru.Schemas == nil {
		ru.Schemas = []URN{UserSchema}
	}
	err = a.client.Scim(a.context, http.MethodPost, "/preview/scim/v2/Users", ru, &user)
	return user, err
}

// Filter retrieves users by filter
func (a UsersAPI) Filter(filter string) (u []User, err error) {
	var users UserList
	req := map[string]string{}
	if filter != "" {
		req["filter"] = filter
	}
	err = a.client.Scim(a.context, http.MethodGet, "/preview/scim/v2/Users", req, &users)
	if err != nil {
		return
	}
	u = users.Resources
	return
}

func (a UsersAPI) Read(userID string) (User, error) {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	return a.readByPath(userPath)
}

// Me gets user information about caller
func (a UsersAPI) Me() (User, error) {
	return a.readByPath("/preview/scim/v2/Me")
}

func (a UsersAPI) readByPath(userPath string) (user User, err error) {
	err = a.client.Scim(a.context, http.MethodGet, userPath, nil, &user)
	return
}

// Update replaces user information for given ID
func (a UsersAPI) Update(userID string, updateRequest User) error {
	user, err := a.Read(userID)
	if err != nil {
		return err
	}
	updateRequest.Groups = user.Groups
	updateRequest.Roles = user.Roles
	if updateRequest.Schemas == nil {
		updateRequest.Schemas = []URN{UserSchema}
	}
	return a.client.Scim(a.context, http.MethodPut,
		fmt.Sprintf("/preview/scim/v2/Users/%v", userID),
		updateRequest, nil)
}

// Patch updates resource-friendly entity
func (a UsersAPI) Patch(userID string, r patchRequest) error {
	return a.client.Scim(a.context, http.MethodPatch, fmt.Sprintf("/preview/scim/v2/Users/%v", userID), r, nil)
}

// Delete will delete the user given the user id
func (a UsersAPI) Delete(userID string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	return a.client.Scim(a.context, http.MethodDelete, userPath, nil, nil)
}
