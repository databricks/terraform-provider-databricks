package scim

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/databricks/terraform-provider-databricks/common"
)

// NewUsersAPI creates UsersAPI instance from provider meta.
// apiLevel controls whether account-level or workspace-level SCIM endpoints are used.
// Pass "" to infer from the provider host.
func NewUsersAPI(ctx context.Context, m any, apiLevel string) UsersAPI {
	return UsersAPI{
		client:   m.(*common.DatabricksClient),
		context:  ctx,
		ApiLevel: apiLevel,
	}
}

// UsersAPI exposes the scim user API
type UsersAPI struct {
	client   *common.DatabricksClient
	context  context.Context
	ApiLevel string
}

// Create user in the backend
func (a UsersAPI) Create(ru User) (user User, err error) {
	if ru.Schemas == nil {
		ru.Schemas = []URN{UserSchema}
	}
	err = a.client.Scim(a.context, http.MethodPost, "/preview/scim/v2/Users", ru, &user, a.ApiLevel)
	return user, err
}

// Filter retrieves users by filter
func (a UsersAPI) Filter(filter string, excludeRoles bool) (u []User, err error) {
	var users UserList
	req := map[string]string{}
	if filter != "" {
		req["filter"] = filter
	}
	// We exclude roles to reduce load on the scim service
	if excludeRoles {
		req["excludedAttributes"] = "roles"
	}
	err = a.client.Scim(a.context, http.MethodGet, "/preview/scim/v2/Users", req, &users, a.ApiLevel)
	if err != nil {
		return
	}
	u = users.Resources
	return
}

// ListAll retrieves all users with the given attributes, handling SCIM pagination.
func (a UsersAPI) ListAll(attributes string) ([]User, error) {
	startIndex := 1
	var result []User
	for {
		req := map[string]string{
			"count":      "10000",
			"startIndex": strconv.Itoa(startIndex),
		}
		if attributes != "" {
			req["attributes"] = attributes
		}
		var page UserList
		err := a.client.Scim(a.context, http.MethodGet, "/preview/scim/v2/Users", req, &page, a.ApiLevel)
		if err != nil {
			return nil, err
		}
		result = append(result, page.Resources...)
		if len(page.Resources) == 0 || len(result) >= int(page.TotalResults) {
			break
		}
		startIndex += len(page.Resources)
	}
	return result, nil
}

func (a UsersAPI) Read(userID, attributes string) (User, error) {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v?attributes=%s", userID, attributes)
	return a.readByPath(userPath)
}

// Me gets user information about caller
func (a UsersAPI) Me() (User, error) {
	return a.readByPath("/preview/scim/v2/Me")
}

func (a UsersAPI) readByPath(userPath string) (user User, err error) {
	err = a.client.Scim(a.context, http.MethodGet, userPath, nil, &user, a.ApiLevel)
	return
}

// Update replaces user information for given ID
func (a UsersAPI) Update(userID string, updateRequest User) error {
	user, err := a.Read(userID, "groups,roles")
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
		updateRequest, nil, a.ApiLevel)
}

// Patch updates resource-friendly entity
func (a UsersAPI) Patch(userID string, r patchRequest) error {
	return a.client.Scim(a.context, http.MethodPatch, fmt.Sprintf("/preview/scim/v2/Users/%v", userID), r, nil, a.ApiLevel)
}

// Delete will delete the user given the user id
func (a UsersAPI) Delete(userID string) error {
	userPath := fmt.Sprintf("/preview/scim/v2/Users/%v", userID)
	return a.client.Scim(a.context, http.MethodDelete, userPath, nil, nil, a.ApiLevel)
}

func (a UsersAPI) UpdateEntitlements(userID string, entitlements patchRequest) error {
	return a.client.Scim(a.context, http.MethodPatch,
		fmt.Sprintf("/preview/scim/v2/Users/%v", userID), entitlements, nil, a.ApiLevel)
}
