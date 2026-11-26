// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AccessControl API methods
type accessControlImpl struct {
	client *client.DatabricksClient
}

func (a *accessControlImpl) CheckPolicy(ctx context.Context, request CheckPolicyRequest) (*CheckPolicyResponse, error) {
	var checkPolicyResponse CheckPolicyResponse
	path := "/api/2.0/access-control/check-policy-v2"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &checkPolicyResponse)
	return &checkPolicyResponse, err
}

// unexported type that holds implementations of just AccountAccessControl API methods
type accountAccessControlImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	var getAssignableRolesForResourceResponse GetAssignableRolesForResourceResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/assignable-roles", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getAssignableRolesForResourceResponse)
	return &getAssignableRolesForResourceResponse, err
}

func (a *accountAccessControlImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

func (a *accountAccessControlImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

// unexported type that holds implementations of just AccountAccessControlProxy API methods
type accountAccessControlProxyImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlProxyImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	var getAssignableRolesForResourceResponse GetAssignableRolesForResourceResponse
	path := "/api/2.0/preview/accounts/access-control/assignable-roles"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getAssignableRolesForResourceResponse)
	return &getAssignableRolesForResourceResponse, err
}

func (a *accountAccessControlProxyImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := "/api/2.0/preview/accounts/access-control/rule-sets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

func (a *accountAccessControlProxyImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := "/api/2.0/preview/accounts/access-control/rule-sets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

// unexported type that holds implementations of just AccountGroupsV2 API methods
type accountGroupsV2Impl struct {
	client *client.DatabricksClient
}

func (a *accountGroupsV2Impl) Create(ctx context.Context, request CreateAccountGroupRequest) (*AccountGroup, error) {
	var accountGroup AccountGroup
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &accountGroup)
	return &accountGroup, err
}

func (a *accountGroupsV2Impl) Delete(ctx context.Context, request DeleteAccountGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountGroupsV2Impl) Get(ctx context.Context, request GetAccountGroupRequest) (*AccountGroup, error) {
	var accountGroup AccountGroup
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &accountGroup)
	return &accountGroup, err
}

// Gets all details of the groups associated with the Databricks account. As of
// 08/22/2025, this endpoint will no longer return members. Instead, members
// should be retrieved by iterating through `Get group details`. Existing
// accounts that rely on this attribute will not be impacted and will continue
// receiving member data as before.
func (a *accountGroupsV2Impl) List(ctx context.Context, request ListAccountGroupsRequest) listing.Iterator[AccountGroup] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
	}
	getNextPage := func(ctx context.Context, req ListAccountGroupsRequest) (*ListAccountGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAccountGroupsResponse) []AccountGroup {
		return resp.Resources
	}
	getNextReq := func(resp *ListAccountGroupsResponse) *ListAccountGroupsRequest {
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
	return iterator
}

// Gets all details of the groups associated with the Databricks account. As of
// 08/22/2025, this endpoint will no longer return members. Instead, members
// should be retrieved by iterating through `Get group details`. Existing
// accounts that rely on this attribute will not be impacted and will continue
// receiving member data as before.
func (a *accountGroupsV2Impl) ListAll(ctx context.Context, request ListAccountGroupsRequest) ([]AccountGroup, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[AccountGroup, int64](ctx, iterator, request.Count)

}

func (a *accountGroupsV2Impl) internalList(ctx context.Context, request ListAccountGroupsRequest) (*ListAccountGroupsResponse, error) {
	var listAccountGroupsResponse ListAccountGroupsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAccountGroupsResponse)
	return &listAccountGroupsResponse, err
}

func (a *accountGroupsV2Impl) Patch(ctx context.Context, request PatchAccountGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *accountGroupsV2Impl) Update(ctx context.Context, request UpdateAccountGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just AccountServicePrincipalsV2 API methods
type accountServicePrincipalsV2Impl struct {
	client *client.DatabricksClient
}

func (a *accountServicePrincipalsV2Impl) Create(ctx context.Context, request CreateAccountServicePrincipalRequest) (*AccountServicePrincipal, error) {
	var accountServicePrincipal AccountServicePrincipal
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &accountServicePrincipal)
	return &accountServicePrincipal, err
}

func (a *accountServicePrincipalsV2Impl) Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountServicePrincipalsV2Impl) Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*AccountServicePrincipal, error) {
	var accountServicePrincipal AccountServicePrincipal
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &accountServicePrincipal)
	return &accountServicePrincipal, err
}

// Gets the set of service principals associated with a Databricks account.
func (a *accountServicePrincipalsV2Impl) List(ctx context.Context, request ListAccountServicePrincipalsRequest) listing.Iterator[AccountServicePrincipal] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
	}
	getNextPage := func(ctx context.Context, req ListAccountServicePrincipalsRequest) (*ListAccountServicePrincipalsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAccountServicePrincipalsResponse) []AccountServicePrincipal {
		return resp.Resources
	}
	getNextReq := func(resp *ListAccountServicePrincipalsResponse) *ListAccountServicePrincipalsRequest {
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
	return iterator
}

// Gets the set of service principals associated with a Databricks account.
func (a *accountServicePrincipalsV2Impl) ListAll(ctx context.Context, request ListAccountServicePrincipalsRequest) ([]AccountServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[AccountServicePrincipal, int64](ctx, iterator, request.Count)

}

func (a *accountServicePrincipalsV2Impl) internalList(ctx context.Context, request ListAccountServicePrincipalsRequest) (*ListAccountServicePrincipalsResponse, error) {
	var listAccountServicePrincipalsResponse ListAccountServicePrincipalsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAccountServicePrincipalsResponse)
	return &listAccountServicePrincipalsResponse, err
}

func (a *accountServicePrincipalsV2Impl) Patch(ctx context.Context, request PatchAccountServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *accountServicePrincipalsV2Impl) Update(ctx context.Context, request UpdateAccountServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just AccountUsersV2 API methods
type accountUsersV2Impl struct {
	client *client.DatabricksClient
}

func (a *accountUsersV2Impl) Create(ctx context.Context, request CreateAccountUserRequest) (*AccountUser, error) {
	var accountUser AccountUser
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &accountUser)
	return &accountUser, err
}

func (a *accountUsersV2Impl) Delete(ctx context.Context, request DeleteAccountUserRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountUsersV2Impl) Get(ctx context.Context, request GetAccountUserRequest) (*AccountUser, error) {
	var accountUser AccountUser
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &accountUser)
	return &accountUser, err
}

// Gets details for all the users associated with a Databricks account.
func (a *accountUsersV2Impl) List(ctx context.Context, request ListAccountUsersRequest) listing.Iterator[AccountUser] {

	request.StartIndex = 1 // SCIM offset starts from 1
	if request.Count == 0 {
		request.Count = 10000
	}
	getNextPage := func(ctx context.Context, req ListAccountUsersRequest) (*ListAccountUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAccountUsersResponse) []AccountUser {
		return resp.Resources
	}
	getNextReq := func(resp *ListAccountUsersResponse) *ListAccountUsersRequest {
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
	return iterator
}

// Gets details for all the users associated with a Databricks account.
func (a *accountUsersV2Impl) ListAll(ctx context.Context, request ListAccountUsersRequest) ([]AccountUser, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[AccountUser, int64](ctx, iterator, request.Count)

}

func (a *accountUsersV2Impl) internalList(ctx context.Context, request ListAccountUsersRequest) (*ListAccountUsersResponse, error) {
	var listAccountUsersResponse ListAccountUsersResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAccountUsersResponse)
	return &listAccountUsersResponse, err
}

func (a *accountUsersV2Impl) Patch(ctx context.Context, request PatchAccountUserRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *accountUsersV2Impl) Update(ctx context.Context, request UpdateAccountUserRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just AssumeGroup API methods
type assumeGroupImpl struct {
	client *client.DatabricksClient
}

func (a *assumeGroupImpl) CreateRestrictedSharingGroup(ctx context.Context, request CreateRestrictedSharingGroupRequest) (*RestrictedSharingGroup, error) {
	var restrictedSharingGroup RestrictedSharingGroup
	path := "/api/2.0/access-control/restricted-sharing-groups"
	queryParams := make(map[string]any)

	if request.RestrictedSharingGroupId != "" {
		queryParams["restricted_sharing_group_id"] = request.RestrictedSharingGroupId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.RestrictedSharingGroup, &restrictedSharingGroup)
	return &restrictedSharingGroup, err
}

func (a *assumeGroupImpl) DeleteRestrictedSharingGroup(ctx context.Context, request DeleteRestrictedSharingGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/access-control/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *assumeGroupImpl) GetRestrictedSharingGroup(ctx context.Context, request GetRestrictedSharingGroupRequest) (*RestrictedSharingGroup, error) {
	var restrictedSharingGroup RestrictedSharingGroup
	path := fmt.Sprintf("/api/2.0/access-control/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &restrictedSharingGroup)
	return &restrictedSharingGroup, err
}

// Lists the restricted sharing groups in the workspace. Requires workspace
// admin.
func (a *assumeGroupImpl) ListRestrictedSharingGroups(ctx context.Context, request ListRestrictedSharingGroupsRequest) listing.Iterator[RestrictedSharingGroup] {

	getNextPage := func(ctx context.Context, req ListRestrictedSharingGroupsRequest) (*ListRestrictedSharingGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListRestrictedSharingGroups(ctx, req)
	}
	getItems := func(resp *ListRestrictedSharingGroupsResponse) []RestrictedSharingGroup {
		return resp.RestrictedSharingGroups
	}
	getNextReq := func(resp *ListRestrictedSharingGroupsResponse) *ListRestrictedSharingGroupsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the restricted sharing groups in the workspace. Requires workspace
// admin.
func (a *assumeGroupImpl) ListRestrictedSharingGroupsAll(ctx context.Context, request ListRestrictedSharingGroupsRequest) ([]RestrictedSharingGroup, error) {
	iterator := a.ListRestrictedSharingGroups(ctx, request)
	return listing.ToSlice[RestrictedSharingGroup](ctx, iterator)
}

func (a *assumeGroupImpl) internalListRestrictedSharingGroups(ctx context.Context, request ListRestrictedSharingGroupsRequest) (*ListRestrictedSharingGroupsResponse, error) {
	var listRestrictedSharingGroupsResponse ListRestrictedSharingGroupsResponse
	path := "/api/2.0/access-control/restricted-sharing-groups"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRestrictedSharingGroupsResponse)
	return &listRestrictedSharingGroupsResponse, err
}

// unexported type that holds implementations of just CurrentUser API methods
type currentUserImpl struct {
	client *client.DatabricksClient
}

func (a *currentUserImpl) Me(ctx context.Context, request MeRequest) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Me"
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

// unexported type that holds implementations of just GroupsV2 API methods
type groupsV2Impl struct {
	client *client.DatabricksClient
}

func (a *groupsV2Impl) Create(ctx context.Context, request CreateGroupRequest) (*Group, error) {
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

func (a *groupsV2Impl) Delete(ctx context.Context, request DeleteGroupRequest) error {
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

func (a *groupsV2Impl) Get(ctx context.Context, request GetGroupRequest) (*Group, error) {
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
func (a *groupsV2Impl) List(ctx context.Context, request ListGroupsRequest) listing.Iterator[Group] {

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
	return iterator
}

// Gets all details of the groups associated with the Databricks workspace.
func (a *groupsV2Impl) ListAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Group, int64](ctx, iterator, request.Count)

}

func (a *groupsV2Impl) internalList(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
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

func (a *groupsV2Impl) Patch(ctx context.Context, request PatchGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *groupsV2Impl) Update(ctx context.Context, request UpdateGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just PermissionMigration API methods
type permissionMigrationImpl struct {
	client *client.DatabricksClient
}

func (a *permissionMigrationImpl) MigratePermissions(ctx context.Context, request MigratePermissionsRequest) (*MigratePermissionsResponse, error) {
	var migratePermissionsResponse MigratePermissionsResponse
	path := "/api/2.0/permissionmigration"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &migratePermissionsResponse)
	return &migratePermissionsResponse, err
}

// unexported type that holds implementations of just Permissions API methods
type permissionsImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsImpl) Get(ctx context.Context, request GetPermissionRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &objectPermissions)
	return &objectPermissions, err
}

func (a *permissionsImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	var getPermissionLevelsResponse GetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPermissionLevelsResponse)
	return &getPermissionLevelsResponse, err
}

func (a *permissionsImpl) Set(ctx context.Context, request SetObjectPermissions) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &objectPermissions)
	return &objectPermissions, err
}

func (a *permissionsImpl) Update(ctx context.Context, request UpdateObjectPermissions) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.RequestObjectType, request.RequestObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &objectPermissions)
	return &objectPermissions, err
}

// unexported type that holds implementations of just ServicePrincipalsV2 API methods
type servicePrincipalsV2Impl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalsV2Impl) Create(ctx context.Context, request CreateServicePrincipalRequest) (*ServicePrincipal, error) {
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

func (a *servicePrincipalsV2Impl) Delete(ctx context.Context, request DeleteServicePrincipalRequest) error {
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

func (a *servicePrincipalsV2Impl) Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
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
func (a *servicePrincipalsV2Impl) List(ctx context.Context, request ListServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

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
	return iterator
}

// Gets the set of service principals associated with a Databricks workspace.
func (a *servicePrincipalsV2Impl) ListAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ServicePrincipal, int64](ctx, iterator, request.Count)

}

func (a *servicePrincipalsV2Impl) internalList(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
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

func (a *servicePrincipalsV2Impl) Patch(ctx context.Context, request PatchServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *servicePrincipalsV2Impl) Update(ctx context.Context, request UpdateServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

// unexported type that holds implementations of just UsersV2 API methods
type usersV2Impl struct {
	client *client.DatabricksClient
}

func (a *usersV2Impl) Create(ctx context.Context, request CreateUserRequest) (*User, error) {
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

func (a *usersV2Impl) Delete(ctx context.Context, request DeleteUserRequest) error {
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

func (a *usersV2Impl) Get(ctx context.Context, request GetUserRequest) (*User, error) {
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

func (a *usersV2Impl) GetPermissionLevels(ctx context.Context, request GetPasswordPermissionLevelsRequest) (*GetPasswordPermissionLevelsResponse, error) {
	var getPasswordPermissionLevelsResponse GetPasswordPermissionLevelsResponse
	path := "/api/2.0/permissions/authorization/passwords/permissionLevels"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPasswordPermissionLevelsResponse)
	return &getPasswordPermissionLevelsResponse, err
}

func (a *usersV2Impl) GetPermissions(ctx context.Context, request GetPasswordPermissionsRequest) (*PasswordPermissions, error) {
	var passwordPermissions PasswordPermissions
	path := "/api/2.0/permissions/authorization/passwords"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &passwordPermissions)
	return &passwordPermissions, err
}

// Gets details for all the users associated with a Databricks workspace.
func (a *usersV2Impl) List(ctx context.Context, request ListUsersRequest) listing.Iterator[User] {

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
	return iterator
}

// Gets details for all the users associated with a Databricks workspace.
func (a *usersV2Impl) ListAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[User, int64](ctx, iterator, request.Count)

}

func (a *usersV2Impl) internalList(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
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

func (a *usersV2Impl) Patch(ctx context.Context, request PatchUserRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, nil)
	return err
}

func (a *usersV2Impl) SetPermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error) {
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

func (a *usersV2Impl) Update(ctx context.Context, request UpdateUserRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, nil)
	return err
}

func (a *usersV2Impl) UpdatePermissions(ctx context.Context, request PasswordPermissionsRequest) (*PasswordPermissions, error) {
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

// unexported type that holds implementations of just WorkspaceAssignment API methods
type workspaceAssignmentImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceAssignmentImpl) Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *workspaceAssignmentImpl) Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error) {
	var workspacePermissions WorkspacePermissions
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/permissions", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspacePermissions)
	return &workspacePermissions, err
}

// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
func (a *workspaceAssignmentImpl) List(ctx context.Context, request ListWorkspaceAssignmentRequest) listing.Iterator[PermissionAssignment] {

	getNextPage := func(ctx context.Context, req ListWorkspaceAssignmentRequest) (*PermissionAssignments, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *PermissionAssignments) []PermissionAssignment {
		return resp.PermissionAssignments
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get the permission assignments for the specified Databricks account and
// Databricks workspace.
func (a *workspaceAssignmentImpl) ListAll(ctx context.Context, request ListWorkspaceAssignmentRequest) ([]PermissionAssignment, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PermissionAssignment](ctx, iterator)
}

func (a *workspaceAssignmentImpl) internalList(ctx context.Context, request ListWorkspaceAssignmentRequest) (*PermissionAssignments, error) {
	var permissionAssignments PermissionAssignments
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &permissionAssignments)
	return &permissionAssignments, err
}

func (a *workspaceAssignmentImpl) Update(ctx context.Context, request UpdateWorkspaceAssignments) (*PermissionAssignment, error) {
	var permissionAssignment PermissionAssignment
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &permissionAssignment)
	return &permissionAssignment, err
}
