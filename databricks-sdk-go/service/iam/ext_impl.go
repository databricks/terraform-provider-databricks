package iam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AccountGroups API methods
type accountGroupsImpl struct {
	client *client.DatabricksClient
}

func (a *accountGroupsImpl) Create(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *accountGroupsImpl) Delete(ctx context.Context, request DeleteAccountGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountGroupsImpl) Get(ctx context.Context, request GetAccountGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

// Gets all details of the groups associated with the Databricks account.
func (a *accountGroupsImpl) List(ctx context.Context, request ListAccountGroupsRequest) listing.Iterator[Group] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
	}
	getNextPage := func(ctx context.Context, req ListAccountGroupsRequest) (*ListGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListGroupsResponse) []Group {
		return resp.Resources
	}
	getNextReq := func(resp *ListGroupsResponse) *ListAccountGroupsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[Group, string](
		iterator,
		func(item Group) string {
			return item.Id
		})
	return dedupedIterator
}

// Gets all details of the groups associated with the Databricks account.
func (a *accountGroupsImpl) ListAll(ctx context.Context, request ListAccountGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

}

func (a *accountGroupsImpl) internalList(ctx context.Context, request ListAccountGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *accountGroupsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *accountGroupsImpl) Update(ctx context.Context, request Group) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just AccountServicePrincipals API methods
type accountServicePrincipalsImpl struct {
	client *client.DatabricksClient
}

func (a *accountServicePrincipalsImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountServicePrincipalsImpl) Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountServicePrincipalsImpl) Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

// Gets the set of service principals associated with a Databricks account.
func (a *accountServicePrincipalsImpl) List(ctx context.Context, request ListAccountServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
	}
	getNextPage := func(ctx context.Context, req ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalResponse) []ServicePrincipal {
		return resp.Resources
	}
	getNextReq := func(resp *ListServicePrincipalResponse) *ListAccountServicePrincipalsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[ServicePrincipal, string](
		iterator,
		func(item ServicePrincipal) string {
			return item.Id
		})
	return dedupedIterator
}

// Gets the set of service principals associated with a Databricks account.
func (a *accountServicePrincipalsImpl) ListAll(ctx context.Context, request ListAccountServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

}

func (a *accountServicePrincipalsImpl) internalList(ctx context.Context, request ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

func (a *accountServicePrincipalsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *accountServicePrincipalsImpl) Update(ctx context.Context, request ServicePrincipal) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just AccountUsers API methods
type accountUsersImpl struct {
	client *client.DatabricksClient
}

func (a *accountUsersImpl) Create(ctx context.Context, request User) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *accountUsersImpl) Delete(ctx context.Context, request DeleteAccountUserRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountUsersImpl) Get(ctx context.Context, request GetAccountUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
}

// Gets details for all the users associated with a Databricks account.
func (a *accountUsersImpl) List(ctx context.Context, request ListAccountUsersRequest) listing.Iterator[User] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
	}
	getNextPage := func(ctx context.Context, req ListAccountUsersRequest) (*ListUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListUsersResponse) []User {
		return resp.Resources
	}
	getNextReq := func(resp *ListUsersResponse) *ListAccountUsersRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[User, string](
		iterator,
		func(item User) string {
			return item.Id
		})
	return dedupedIterator
}

// Gets details for all the users associated with a Databricks account.
func (a *accountUsersImpl) ListAll(ctx context.Context, request ListAccountUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

}

func (a *accountUsersImpl) internalList(ctx context.Context, request ListAccountUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *accountUsersImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *accountUsersImpl) Update(ctx context.Context, request User) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just Groups API methods
type groupsImpl struct {
	client *client.DatabricksClient
}

func (a *groupsImpl) Create(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := "/api/2.0/preview/scim/v2/Groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *groupsImpl) Delete(ctx context.Context, request DeleteGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *groupsImpl) Get(ctx context.Context, request GetGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

// Gets all details of the groups associated with the Databricks workspace.
func (a *groupsImpl) List(ctx context.Context, request ListGroupsRequest) listing.Iterator[Group] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
	}
	getNextPage := func(ctx context.Context, req ListGroupsRequest) (*ListGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListGroupsResponse) []Group {
		return resp.Resources
	}
	getNextReq := func(resp *ListGroupsResponse) *ListGroupsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[Group, string](
		iterator,
		func(item Group) string {
			return item.Id
		})
	return dedupedIterator
}

// Gets all details of the groups associated with the Databricks workspace.
func (a *groupsImpl) ListAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

}

func (a *groupsImpl) internalList(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/preview/scim/v2/Groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *groupsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *groupsImpl) Update(ctx context.Context, request Group) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just ServicePrincipals API methods
type servicePrincipalsImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalsImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *servicePrincipalsImpl) Delete(ctx context.Context, request DeleteServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *servicePrincipalsImpl) Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

// Gets the set of service principals associated with a Databricks workspace.
func (a *servicePrincipalsImpl) List(ctx context.Context, request ListServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
	}
	getNextPage := func(ctx context.Context, req ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalResponse) []ServicePrincipal {
		return resp.Resources
	}
	getNextReq := func(resp *ListServicePrincipalResponse) *ListServicePrincipalsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[ServicePrincipal, string](
		iterator,
		func(item ServicePrincipal) string {
			return item.Id
		})
	return dedupedIterator
}

// Gets the set of service principals associated with a Databricks workspace.
func (a *servicePrincipalsImpl) ListAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

}

func (a *servicePrincipalsImpl) internalList(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

func (a *servicePrincipalsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *servicePrincipalsImpl) Update(ctx context.Context, request ServicePrincipal) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just Users API methods
type usersImpl struct {
	client *client.DatabricksClient
}

func (a *usersImpl) Create(ctx context.Context, request User) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *usersImpl) Delete(ctx context.Context, request DeleteUserRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *usersImpl) Get(ctx context.Context, request GetUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *usersImpl) GetPermissionLevels(ctx context.Context) (*GetPasswordPermissionLevelsResponse, error) {
	var getPasswordPermissionLevelsResponse GetPasswordPermissionLevelsResponse
	path := "/api/2.0/permissions/authorization/passwords/permissionLevels"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getPasswordPermissionLevelsResponse)
	return &getPasswordPermissionLevelsResponse, err
}

func (a *usersImpl) GetPermissions(ctx context.Context) (*PasswordPermissions, error) {
	var passwordPermissions PasswordPermissions
	path := "/api/2.0/permissions/authorization/passwords"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &passwordPermissions)
	return &passwordPermissions, err
}

// Gets details for all the users associated with a Databricks workspace.
func (a *usersImpl) List(ctx context.Context, request ListUsersRequest) listing.Iterator[User] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
	}
	getNextPage := func(ctx context.Context, req ListUsersRequest) (*ListUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListUsersResponse) []User {
		return resp.Resources
	}
	getNextReq := func(resp *ListUsersResponse) *ListUsersRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.StartIndex = resp.StartIndex + int64(len(resp.Resources))
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[User, string](
		iterator,
		func(item User) string {
			return item.Id
		})
	return dedupedIterator
}

// Gets details for all the users associated with a Databricks workspace.
func (a *usersImpl) ListAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

}

func (a *usersImpl) internalList(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0/preview/scim/v2/Users"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *usersImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *usersImpl) SetPermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error) {
	var passwordPermissions PasswordPermissions
	path := "/api/2.0/permissions/authorization/passwords"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &passwordPermissions)
	return &passwordPermissions, err
}

func (a *usersImpl) Update(ctx context.Context, request User) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

func (a *usersImpl) UpdatePermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error) {
	var passwordPermissions PasswordPermissions
	path := "/api/2.0/permissions/authorization/passwords"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &passwordPermissions)
	return &passwordPermissions, err
}
