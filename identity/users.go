package identity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/common"
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

// UserEntity entity from which resource schema is made
type UserEntity struct {
	UserName                string `json:"user_name"`
	DisplayName             string `json:"display_name,omitempty" tf:"computed"`
	Active                  bool   `json:"active,omitempty"`
	AllowClusterCreate      bool   `json:"allow_cluster_create,omitempty"`
	AllowInstancePoolCreate bool   `json:"allow_instance_pool_create,omitempty"`
}

func (u UserEntity) toRequest() ScimUser {
	entitlements := []entitlementsListItem{}
	if u.AllowClusterCreate {
		entitlements = append(entitlements, entitlementsListItem{
			Value: Entitlement("allow-cluster-create"),
		})
	}
	if u.AllowInstancePoolCreate {
		entitlements = append(entitlements, entitlementsListItem{
			Value: Entitlement("allow-instance-pool-create"),
		})
	}
	return ScimUser{
		Schemas:      []URN{UserSchema},
		UserName:     u.UserName,
		Active:       u.Active,
		DisplayName:  u.DisplayName,
		Entitlements: entitlements,
	}
}

// Create ..
func (a UsersAPI) Create(ru UserEntity) (user ScimUser, err error) {
	err = a.client.Scim(a.context, http.MethodPost, "/preview/scim/v2/Users", ru.toRequest(), &user)
	return user, err
}

// Read reads resource-friendly entity
func (a UsersAPI) Read(userID string) (ru UserEntity, err error) {
	user, err := a.read(userID)
	if err != nil {
		return
	}
	ru.UserName = user.UserName
	ru.DisplayName = user.DisplayName
	ru.Active = user.Active
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

func (a UsersAPI) read(userID string) (ScimUser, error) {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	return a.readByPath(userPath)
}

// Me gets user information about caller
func (a UsersAPI) Me() (ScimUser, error) {
	return a.readByPath("/preview/scim/v2/Me")
}

func (a UsersAPI) readByPath(userPath string) (user ScimUser, err error) {
	err = a.client.Scim(a.context, http.MethodGet, userPath, nil, &user)
	return
}

// Update replaces resource-friendly-entity
func (a UsersAPI) Update(userID string, ru UserEntity) error {
	user, err := a.read(userID)
	if err != nil {
		return err
	}
	updateRequest := ru.toRequest()
	updateRequest.Groups = user.Groups
	updateRequest.Roles = user.Roles
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
