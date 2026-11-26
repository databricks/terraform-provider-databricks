// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iamv2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just AccountIamV2 API methods
type accountIamV2Impl struct {
	client *client.DatabricksClient
}

func (a *accountIamV2Impl) CreateDirectGroupMember(ctx context.Context, request CreateDirectGroupMemberRequest) (*DirectGroupMember, error) {
	var directGroupMember DirectGroupMember
	path := fmt.Sprintf("/api/2.0/identity/%v/direct-members", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DirectGroupMember, &directGroupMember)
	return &directGroupMember, err
}

func (a *accountIamV2Impl) CreateGroup(ctx context.Context, request CreateGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/%v/groups", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Group, &group)
	return &group, err
}

func (a *accountIamV2Impl) CreateIdentityVisibilityFilter(ctx context.Context, request CreateIdentityVisibilityFilterRequest) (*IdentityVisibilityFilter, error) {
	var identityVisibilityFilter IdentityVisibilityFilter
	path := fmt.Sprintf("/api/2.0/identity/%v/identity-visibility-filters", request.Parent)
	queryParams := make(map[string]any)

	if request.IdentityVisibilityFilterId != "" || slices.Contains(request.ForceSendFields, "IdentityVisibilityFilterId") {
		queryParams["identity_visibility_filter_id"] = request.IdentityVisibilityFilterId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.IdentityVisibilityFilter, &identityVisibilityFilter)
	return &identityVisibilityFilter, err
}

func (a *accountIamV2Impl) CreateServicePrincipal(ctx context.Context, request CreateServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/%v/service-principals", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ServicePrincipal, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountIamV2Impl) CreateUser(ctx context.Context, request CreateUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/%v/users", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.User, &user)
	return &user, err
}

func (a *accountIamV2Impl) CreateWorkspaceAssignment(ctx context.Context, request CreateWorkspaceAssignmentRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := fmt.Sprintf("/api/2.0/identity/%v/workspace-assignments", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAssignment, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *accountIamV2Impl) CreateWorkspaceAssignmentDetail(ctx context.Context, request CreateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignment-details", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *accountIamV2Impl) DeleteDirectGroupMember(ctx context.Context, request DeleteDirectGroupMemberRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteGroup(ctx context.Context, request DeleteGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteIdentityVisibilityFilter(ctx context.Context, request DeleteIdentityVisibilityFilterRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteUser(ctx context.Context, request DeleteUserRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteWorkspaceAssignment(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) DeleteWorkspaceAssignmentDetail(ctx context.Context, request DeleteWorkspaceAssignmentDetailRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignment-details/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *accountIamV2Impl) GetDirectGroupMember(ctx context.Context, request GetDirectGroupMemberRequest) (*DirectGroupMember, error) {
	var directGroupMember DirectGroupMember
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &directGroupMember)
	return &directGroupMember, err
}

func (a *accountIamV2Impl) GetExternalGroup(ctx context.Context, request GetExternalGroupRequest) (*ExternalGroup, error) {
	var externalGroup ExternalGroup
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &externalGroup)
	return &externalGroup, err
}

func (a *accountIamV2Impl) GetExternalServicePrincipal(ctx context.Context, request GetExternalServicePrincipalRequest) (*ExternalServicePrincipal, error) {
	var externalServicePrincipal ExternalServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &externalServicePrincipal)
	return &externalServicePrincipal, err
}

func (a *accountIamV2Impl) GetExternalUser(ctx context.Context, request GetExternalUserRequest) (*ExternalUser, error) {
	var externalUser ExternalUser
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &externalUser)
	return &externalUser, err
}

func (a *accountIamV2Impl) GetGroup(ctx context.Context, request GetGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &group)
	return &group, err
}

func (a *accountIamV2Impl) GetIdentityVisibilityFilter(ctx context.Context, request GetIdentityVisibilityFilterRequest) (*IdentityVisibilityFilter, error) {
	var identityVisibilityFilter IdentityVisibilityFilter
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &identityVisibilityFilter)
	return &identityVisibilityFilter, err
}

func (a *accountIamV2Impl) GetServicePrincipal(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountIamV2Impl) GetUser(ctx context.Context, request GetUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &user)
	return &user, err
}

func (a *accountIamV2Impl) GetWorkspaceAccessDetail(ctx context.Context, request GetWorkspaceAccessDetailRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}

func (a *accountIamV2Impl) GetWorkspaceAssignment(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *accountIamV2Impl) GetWorkspaceAssignmentDetail(ctx context.Context, request GetWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignment-details/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

// Lists the direct members of a group, sorted by display name, with where each
// membership comes from. Results combine memberships added in Databricks (in
// the UI, through the API, or through SCIM provisioning) with, if automatic
// identity management is on and the group has an `external_id`, memberships
// from your identity provider (IdP). Only members in your Databricks account
// are returned: an IdP member of the group appears only if that user, service
// principal, or group has been added to Databricks. Nested groups aren't
// expanded. To get the full membership, list the direct members of each member
// group.
func (a *accountIamV2Impl) ListDirectGroupMembers(ctx context.Context, request ListDirectGroupMembersRequest) listing.Iterator[DirectGroupMember] {

	getNextPage := func(ctx context.Context, req ListDirectGroupMembersRequest) (*ListDirectGroupMembersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDirectGroupMembers(ctx, req)
	}
	getItems := func(resp *ListDirectGroupMembersResponse) []DirectGroupMember {
		return resp.DirectGroupMembers
	}
	getNextReq := func(resp *ListDirectGroupMembersResponse) *ListDirectGroupMembersRequest {
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

// Lists the direct members of a group, sorted by display name, with where each
// membership comes from. Results combine memberships added in Databricks (in
// the UI, through the API, or through SCIM provisioning) with, if automatic
// identity management is on and the group has an `external_id`, memberships
// from your identity provider (IdP). Only members in your Databricks account
// are returned: an IdP member of the group appears only if that user, service
// principal, or group has been added to Databricks. Nested groups aren't
// expanded. To get the full membership, list the direct members of each member
// group.
func (a *accountIamV2Impl) ListDirectGroupMembersAll(ctx context.Context, request ListDirectGroupMembersRequest) ([]DirectGroupMember, error) {
	iterator := a.ListDirectGroupMembers(ctx, request)
	return listing.ToSlice[DirectGroupMember](ctx, iterator)
}

func (a *accountIamV2Impl) internalListDirectGroupMembers(ctx context.Context, request ListDirectGroupMembersRequest) (*ListDirectGroupMembersResponse, error) {
	var listDirectGroupMembersResponse ListDirectGroupMembersResponse
	path := fmt.Sprintf("/api/2.0/identity/%v/direct-members", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDirectGroupMembersResponse)
	return &listDirectGroupMembersResponse, err
}

// Lists the groups in your account, sorted by Databricks ID. Groups in your
// identity provider that haven't been added to Databricks aren't included.
func (a *accountIamV2Impl) ListGroups(ctx context.Context, request ListGroupsRequest) listing.Iterator[Group] {

	getNextPage := func(ctx context.Context, req ListGroupsRequest) (*ListGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListGroups(ctx, req)
	}
	getItems := func(resp *ListGroupsResponse) []Group {
		return resp.Groups
	}
	getNextReq := func(resp *ListGroupsResponse) *ListGroupsRequest {
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

// Lists the groups in your account, sorted by Databricks ID. Groups in your
// identity provider that haven't been added to Databricks aren't included.
func (a *accountIamV2Impl) ListGroupsAll(ctx context.Context, request ListGroupsRequest) ([]Group, error) {
	iterator := a.ListGroups(ctx, request)
	return listing.ToSlice[Group](ctx, iterator)
}

func (a *accountIamV2Impl) internalListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := fmt.Sprintf("/api/2.0/identity/%v/groups", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

// Lists the group visibility filters in an account. It returns the rules, not
// the groups they match.
func (a *accountIamV2Impl) ListIdentityVisibilityFilters(ctx context.Context, request ListIdentityVisibilityFiltersRequest) listing.Iterator[IdentityVisibilityFilter] {

	getNextPage := func(ctx context.Context, req ListIdentityVisibilityFiltersRequest) (*ListIdentityVisibilityFiltersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListIdentityVisibilityFilters(ctx, req)
	}
	getItems := func(resp *ListIdentityVisibilityFiltersResponse) []IdentityVisibilityFilter {
		return resp.IdentityVisibilityFilters
	}
	getNextReq := func(resp *ListIdentityVisibilityFiltersResponse) *ListIdentityVisibilityFiltersRequest {
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

// Lists the group visibility filters in an account. It returns the rules, not
// the groups they match.
func (a *accountIamV2Impl) ListIdentityVisibilityFiltersAll(ctx context.Context, request ListIdentityVisibilityFiltersRequest) ([]IdentityVisibilityFilter, error) {
	iterator := a.ListIdentityVisibilityFilters(ctx, request)
	return listing.ToSlice[IdentityVisibilityFilter](ctx, iterator)
}

func (a *accountIamV2Impl) internalListIdentityVisibilityFilters(ctx context.Context, request ListIdentityVisibilityFiltersRequest) (*ListIdentityVisibilityFiltersResponse, error) {
	var listIdentityVisibilityFiltersResponse ListIdentityVisibilityFiltersResponse
	path := fmt.Sprintf("/api/2.0/identity/%v/identity-visibility-filters", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listIdentityVisibilityFiltersResponse)
	return &listIdentityVisibilityFiltersResponse, err
}

// Lists the service principals in your account, sorted by Databricks ID.
// Service principals in your identity provider that haven't been added to
// Databricks aren't included.
func (a *accountIamV2Impl) ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) listing.Iterator[ServicePrincipal] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalsRequest) (*ListServicePrincipalsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListServicePrincipals(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalsResponse) []ServicePrincipal {
		return resp.ServicePrincipals
	}
	getNextReq := func(resp *ListServicePrincipalsResponse) *ListServicePrincipalsRequest {
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

// Lists the service principals in your account, sorted by Databricks ID.
// Service principals in your identity provider that haven't been added to
// Databricks aren't included.
func (a *accountIamV2Impl) ListServicePrincipalsAll(ctx context.Context, request ListServicePrincipalsRequest) ([]ServicePrincipal, error) {
	iterator := a.ListServicePrincipals(ctx, request)
	return listing.ToSlice[ServicePrincipal](ctx, iterator)
}

func (a *accountIamV2Impl) internalListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalsResponse, error) {
	var listServicePrincipalsResponse ListServicePrincipalsResponse
	path := fmt.Sprintf("/api/2.0/identity/%v/service-principals", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalsResponse)
	return &listServicePrincipalsResponse, err
}

func (a *accountIamV2Impl) ListTransitiveParentGroups(ctx context.Context, request ListTransitiveParentGroupsRequest) (*ListTransitiveParentGroupsResponse, error) {
	var listTransitiveParentGroupsResponse ListTransitiveParentGroupsResponse
	path := fmt.Sprintf("/api/2.0/identity/%v/transitive-parent-groups", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTransitiveParentGroupsResponse)
	return &listTransitiveParentGroupsResponse, err
}

// Lists the users in your account, sorted by Databricks ID. Users in your
// identity provider who haven't been added to Databricks aren't included.
func (a *accountIamV2Impl) ListUsers(ctx context.Context, request ListUsersRequest) listing.Iterator[User] {

	getNextPage := func(ctx context.Context, req ListUsersRequest) (*ListUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListUsers(ctx, req)
	}
	getItems := func(resp *ListUsersResponse) []User {
		return resp.Users
	}
	getNextReq := func(resp *ListUsersResponse) *ListUsersRequest {
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

// Lists the users in your account, sorted by Databricks ID. Users in your
// identity provider who haven't been added to Databricks aren't included.
func (a *accountIamV2Impl) ListUsersAll(ctx context.Context, request ListUsersRequest) ([]User, error) {
	iterator := a.ListUsers(ctx, request)
	return listing.ToSlice[User](ctx, iterator)
}

func (a *accountIamV2Impl) internalListUsers(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := fmt.Sprintf("/api/2.0/identity/%v/users", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listUsersResponse)
	return &listUsersResponse, err
}

// Lists workspace assignment details for a workspace. The response omits the
// per-principal entitlement fields (`entitlements` and
// `effective_entitlements`). To read the entitlements for a single principal,
// get that principal's assignment detail.
func (a *accountIamV2Impl) ListWorkspaceAssignmentDetails(ctx context.Context, request ListWorkspaceAssignmentDetailsRequest) listing.Iterator[WorkspaceAssignmentDetail] {

	getNextPage := func(ctx context.Context, req ListWorkspaceAssignmentDetailsRequest) (*ListWorkspaceAssignmentDetailsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListWorkspaceAssignmentDetails(ctx, req)
	}
	getItems := func(resp *ListWorkspaceAssignmentDetailsResponse) []WorkspaceAssignmentDetail {
		return resp.WorkspaceAssignmentDetails
	}
	getNextReq := func(resp *ListWorkspaceAssignmentDetailsResponse) *ListWorkspaceAssignmentDetailsRequest {
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

// Lists workspace assignment details for a workspace. The response omits the
// per-principal entitlement fields (`entitlements` and
// `effective_entitlements`). To read the entitlements for a single principal,
// get that principal's assignment detail.
func (a *accountIamV2Impl) ListWorkspaceAssignmentDetailsAll(ctx context.Context, request ListWorkspaceAssignmentDetailsRequest) ([]WorkspaceAssignmentDetail, error) {
	iterator := a.ListWorkspaceAssignmentDetails(ctx, request)
	return listing.ToSlice[WorkspaceAssignmentDetail](ctx, iterator)
}

func (a *accountIamV2Impl) internalListWorkspaceAssignmentDetails(ctx context.Context, request ListWorkspaceAssignmentDetailsRequest) (*ListWorkspaceAssignmentDetailsResponse, error) {
	var listWorkspaceAssignmentDetailsResponse ListWorkspaceAssignmentDetailsResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignment-details", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAssignmentDetailsResponse)
	return &listWorkspaceAssignmentDetailsResponse, err
}

// Lists the principals assigned directly to a workspace. The response leaves
// `entitlements` and `effective_entitlements` empty. To read a principal's
// entitlements, get its assignment.
func (a *accountIamV2Impl) ListWorkspaceAssignments(ctx context.Context, request ListWorkspaceAssignmentsRequest) listing.Iterator[WorkspaceAssignment] {

	getNextPage := func(ctx context.Context, req ListWorkspaceAssignmentsRequest) (*ListWorkspaceAssignmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListWorkspaceAssignments(ctx, req)
	}
	getItems := func(resp *ListWorkspaceAssignmentsResponse) []WorkspaceAssignment {
		return resp.WorkspaceAssignments
	}
	getNextReq := func(resp *ListWorkspaceAssignmentsResponse) *ListWorkspaceAssignmentsRequest {
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

// Lists the principals assigned directly to a workspace. The response leaves
// `entitlements` and `effective_entitlements` empty. To read a principal's
// entitlements, get its assignment.
func (a *accountIamV2Impl) ListWorkspaceAssignmentsAll(ctx context.Context, request ListWorkspaceAssignmentsRequest) ([]WorkspaceAssignment, error) {
	iterator := a.ListWorkspaceAssignments(ctx, request)
	return listing.ToSlice[WorkspaceAssignment](ctx, iterator)
}

func (a *accountIamV2Impl) internalListWorkspaceAssignments(ctx context.Context, request ListWorkspaceAssignmentsRequest) (*ListWorkspaceAssignmentsResponse, error) {
	var listWorkspaceAssignmentsResponse ListWorkspaceAssignmentsResponse
	path := fmt.Sprintf("/api/2.0/identity/%v/workspace-assignments", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAssignmentsResponse)
	return &listWorkspaceAssignmentsResponse, err
}

func (a *accountIamV2Impl) ResolveGroup(ctx context.Context, request ResolveGroupRequest) (*ResolveGroupResponse, error) {
	var resolveGroupResponse ResolveGroupResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/groups/resolve-by-external-id", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveGroupResponse)
	return &resolveGroupResponse, err
}

func (a *accountIamV2Impl) ResolveServicePrincipal(ctx context.Context, request ResolveServicePrincipalRequest) (*ResolveServicePrincipalResponse, error) {
	var resolveServicePrincipalResponse ResolveServicePrincipalResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/service-principals/resolve-by-external-id", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveServicePrincipalResponse)
	return &resolveServicePrincipalResponse, err
}

func (a *accountIamV2Impl) ResolveUser(ctx context.Context, request ResolveUserRequest) (*ResolveUserResponse, error) {
	var resolveUserResponse ResolveUserResponse
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/users/resolve-by-external-id", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveUserResponse)
	return &resolveUserResponse, err
}

func (a *accountIamV2Impl) UpdateGroup(ctx context.Context, request UpdateGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Group, &group)
	return &group, err
}

func (a *accountIamV2Impl) UpdateServicePrincipal(ctx context.Context, request UpdateServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ServicePrincipal, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountIamV2Impl) UpdateUser(ctx context.Context, request UpdateUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.User, &user)
	return &user, err
}

func (a *accountIamV2Impl) UpdateWorkspaceAssignment(ctx context.Context, request UpdateWorkspaceAssignmentRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAssignment, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *accountIamV2Impl) UpdateWorkspaceAssignmentDetail(ctx context.Context, request UpdateWorkspaceAssignmentDetailRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/accounts/%v/workspaces/%v/workspace-assignment-details/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

// unexported type that holds implementations of just WorkspaceIamV2 API methods
type workspaceIamV2Impl struct {
	client *client.DatabricksClient
}

func (a *workspaceIamV2Impl) CreateDirectGroupMemberProxy(ctx context.Context, request CreateDirectGroupMemberProxyRequest) (*DirectGroupMember, error) {
	var directGroupMember DirectGroupMember
	path := fmt.Sprintf("/api/2.0/identity/%v/direct-members", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DirectGroupMember, &directGroupMember)
	return &directGroupMember, err
}

func (a *workspaceIamV2Impl) CreateGroupProxy(ctx context.Context, request CreateGroupProxyRequest) (*Group, error) {
	var group Group
	path := "/api/2.0/identity/groups"
	queryParams := make(map[string]any)

	if request.Parent != "" || slices.Contains(request.ForceSendFields, "Parent") {
		queryParams["parent"] = request.Parent
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Group, &group)
	return &group, err
}

func (a *workspaceIamV2Impl) CreateServicePrincipalProxy(ctx context.Context, request CreateServicePrincipalProxyRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/identity/service-principals"
	queryParams := make(map[string]any)

	if request.Parent != "" || slices.Contains(request.ForceSendFields, "Parent") {
		queryParams["parent"] = request.Parent
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ServicePrincipal, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *workspaceIamV2Impl) CreateUserProxy(ctx context.Context, request CreateUserProxyRequest) (*User, error) {
	var user User
	path := "/api/2.0/identity/users"
	queryParams := make(map[string]any)

	if request.Parent != "" || slices.Contains(request.ForceSendFields, "Parent") {
		queryParams["parent"] = request.Parent
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.User, &user)
	return &user, err
}

func (a *workspaceIamV2Impl) CreateWorkspaceAssignmentDetailProxy(ctx context.Context, request CreateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := "/api/2.0/identity/workspace-assignment-details"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *workspaceIamV2Impl) CreateWorkspaceAssignmentProxy(ctx context.Context, request CreateWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := "/api/2.0/identity/workspace-assignments"
	queryParams := make(map[string]any)

	if request.Parent != "" || slices.Contains(request.ForceSendFields, "Parent") {
		queryParams["parent"] = request.Parent
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.WorkspaceAssignment, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *workspaceIamV2Impl) DeleteDirectGroupMemberProxy(ctx context.Context, request DeleteDirectGroupMemberProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
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

func (a *workspaceIamV2Impl) DeleteGroupProxy(ctx context.Context, request DeleteGroupProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
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

func (a *workspaceIamV2Impl) DeleteServicePrincipalProxy(ctx context.Context, request DeleteServicePrincipalProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
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

func (a *workspaceIamV2Impl) DeleteUserProxy(ctx context.Context, request DeleteUserProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
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

func (a *workspaceIamV2Impl) DeleteWorkspaceAssignmentDetailProxy(ctx context.Context, request DeleteWorkspaceAssignmentDetailProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/workspace-assignment-details/%v", request.PrincipalId)
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

func (a *workspaceIamV2Impl) DeleteWorkspaceAssignmentProxy(ctx context.Context, request DeleteWorkspaceAssignmentProxyRequest) error {
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
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

func (a *workspaceIamV2Impl) GetDirectGroupMemberProxy(ctx context.Context, request GetDirectGroupMemberProxyRequest) (*DirectGroupMember, error) {
	var directGroupMember DirectGroupMember
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &directGroupMember)
	return &directGroupMember, err
}

func (a *workspaceIamV2Impl) GetExternalGroupProxy(ctx context.Context, request GetExternalGroupProxyRequest) (*ExternalGroup, error) {
	var externalGroup ExternalGroup
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &externalGroup)
	return &externalGroup, err
}

func (a *workspaceIamV2Impl) GetExternalServicePrincipalProxy(ctx context.Context, request GetExternalServicePrincipalProxyRequest) (*ExternalServicePrincipal, error) {
	var externalServicePrincipal ExternalServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &externalServicePrincipal)
	return &externalServicePrincipal, err
}

func (a *workspaceIamV2Impl) GetExternalUserProxy(ctx context.Context, request GetExternalUserProxyRequest) (*ExternalUser, error) {
	var externalUser ExternalUser
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &externalUser)
	return &externalUser, err
}

func (a *workspaceIamV2Impl) GetGroupProxy(ctx context.Context, request GetGroupProxyRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
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

func (a *workspaceIamV2Impl) GetServicePrincipalProxy(ctx context.Context, request GetServicePrincipalProxyRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
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

func (a *workspaceIamV2Impl) GetUserProxy(ctx context.Context, request GetUserProxyRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
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

func (a *workspaceIamV2Impl) GetWorkspaceAccessDetailLocal(ctx context.Context, request GetWorkspaceAccessDetailLocalRequest) (*WorkspaceAccessDetail, error) {
	var workspaceAccessDetail WorkspaceAccessDetail
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAccessDetail)
	return &workspaceAccessDetail, err
}

func (a *workspaceIamV2Impl) GetWorkspaceAssignmentDetailProxy(ctx context.Context, request GetWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/workspace-assignment-details/%v", request.PrincipalId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *workspaceIamV2Impl) GetWorkspaceAssignmentProxy(ctx context.Context, request GetWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *workspaceIamV2Impl) GetWorkspaceIdentityDetail(ctx context.Context, request GetWorkspaceIdentityDetailRequest) (*WorkspaceIdentityDetail, error) {
	var workspaceIdentityDetail WorkspaceIdentityDetail
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceIdentityDetail)
	return &workspaceIdentityDetail, err
}

// Lists the direct members of a group in the account that contains this
// workspace, sorted by display name, with where each membership comes from.
// Results combine memberships added in Databricks (in the UI, through the API,
// or through SCIM provisioning) with, if automatic identity management is on
// and the group has an `external_id`, memberships from your identity provider
// (IdP). Only members in your Databricks account are returned: an IdP member of
// the group appears only if that user, service principal, or group has been
// added to Databricks. Nested groups aren't expanded. To get the full
// membership, list the direct members of each member group.
func (a *workspaceIamV2Impl) ListDirectGroupMembersProxy(ctx context.Context, request ListDirectGroupMembersProxyRequest) listing.Iterator[DirectGroupMember] {

	getNextPage := func(ctx context.Context, req ListDirectGroupMembersProxyRequest) (*ListDirectGroupMembersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDirectGroupMembersProxy(ctx, req)
	}
	getItems := func(resp *ListDirectGroupMembersResponse) []DirectGroupMember {
		return resp.DirectGroupMembers
	}
	getNextReq := func(resp *ListDirectGroupMembersResponse) *ListDirectGroupMembersProxyRequest {
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

// Lists the direct members of a group in the account that contains this
// workspace, sorted by display name, with where each membership comes from.
// Results combine memberships added in Databricks (in the UI, through the API,
// or through SCIM provisioning) with, if automatic identity management is on
// and the group has an `external_id`, memberships from your identity provider
// (IdP). Only members in your Databricks account are returned: an IdP member of
// the group appears only if that user, service principal, or group has been
// added to Databricks. Nested groups aren't expanded. To get the full
// membership, list the direct members of each member group.
func (a *workspaceIamV2Impl) ListDirectGroupMembersProxyAll(ctx context.Context, request ListDirectGroupMembersProxyRequest) ([]DirectGroupMember, error) {
	iterator := a.ListDirectGroupMembersProxy(ctx, request)
	return listing.ToSlice[DirectGroupMember](ctx, iterator)
}

func (a *workspaceIamV2Impl) internalListDirectGroupMembersProxy(ctx context.Context, request ListDirectGroupMembersProxyRequest) (*ListDirectGroupMembersResponse, error) {
	var listDirectGroupMembersResponse ListDirectGroupMembersResponse
	path := fmt.Sprintf("/api/2.0/identity/%v/direct-members", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDirectGroupMembersResponse)
	return &listDirectGroupMembersResponse, err
}

// Lists the groups in the account that contains this workspace, sorted by
// Databricks ID. Groups in your identity provider that haven't been added to
// Databricks aren't included.
func (a *workspaceIamV2Impl) ListGroupsProxy(ctx context.Context, request ListGroupsProxyRequest) listing.Iterator[Group] {

	getNextPage := func(ctx context.Context, req ListGroupsProxyRequest) (*ListGroupsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListGroupsProxy(ctx, req)
	}
	getItems := func(resp *ListGroupsResponse) []Group {
		return resp.Groups
	}
	getNextReq := func(resp *ListGroupsResponse) *ListGroupsProxyRequest {
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

// Lists the groups in the account that contains this workspace, sorted by
// Databricks ID. Groups in your identity provider that haven't been added to
// Databricks aren't included.
func (a *workspaceIamV2Impl) ListGroupsProxyAll(ctx context.Context, request ListGroupsProxyRequest) ([]Group, error) {
	iterator := a.ListGroupsProxy(ctx, request)
	return listing.ToSlice[Group](ctx, iterator)
}

func (a *workspaceIamV2Impl) internalListGroupsProxy(ctx context.Context, request ListGroupsProxyRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/identity/groups"
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

// Lists the service principals in the account that contains this workspace,
// sorted by Databricks ID. Service principals in your identity provider that
// haven't been added to Databricks aren't included.
func (a *workspaceIamV2Impl) ListServicePrincipalsProxy(ctx context.Context, request ListServicePrincipalsProxyRequest) listing.Iterator[ServicePrincipal] {

	getNextPage := func(ctx context.Context, req ListServicePrincipalsProxyRequest) (*ListServicePrincipalsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListServicePrincipalsProxy(ctx, req)
	}
	getItems := func(resp *ListServicePrincipalsResponse) []ServicePrincipal {
		return resp.ServicePrincipals
	}
	getNextReq := func(resp *ListServicePrincipalsResponse) *ListServicePrincipalsProxyRequest {
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

// Lists the service principals in the account that contains this workspace,
// sorted by Databricks ID. Service principals in your identity provider that
// haven't been added to Databricks aren't included.
func (a *workspaceIamV2Impl) ListServicePrincipalsProxyAll(ctx context.Context, request ListServicePrincipalsProxyRequest) ([]ServicePrincipal, error) {
	iterator := a.ListServicePrincipalsProxy(ctx, request)
	return listing.ToSlice[ServicePrincipal](ctx, iterator)
}

func (a *workspaceIamV2Impl) internalListServicePrincipalsProxy(ctx context.Context, request ListServicePrincipalsProxyRequest) (*ListServicePrincipalsResponse, error) {
	var listServicePrincipalsResponse ListServicePrincipalsResponse
	path := "/api/2.0/identity/service-principals"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listServicePrincipalsResponse)
	return &listServicePrincipalsResponse, err
}

func (a *workspaceIamV2Impl) ListTransitiveParentGroupsProxy(ctx context.Context, request ListTransitiveParentGroupsProxyRequest) (*ListTransitiveParentGroupsResponse, error) {
	var listTransitiveParentGroupsResponse ListTransitiveParentGroupsResponse
	path := fmt.Sprintf("/api/2.0/identity/%v/transitive-parent-groups", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTransitiveParentGroupsResponse)
	return &listTransitiveParentGroupsResponse, err
}

// Lists the users in the account that contains this workspace, sorted by
// Databricks ID. Users in your identity provider who haven't been added to
// Databricks aren't included.
func (a *workspaceIamV2Impl) ListUsersProxy(ctx context.Context, request ListUsersProxyRequest) listing.Iterator[User] {

	getNextPage := func(ctx context.Context, req ListUsersProxyRequest) (*ListUsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListUsersProxy(ctx, req)
	}
	getItems := func(resp *ListUsersResponse) []User {
		return resp.Users
	}
	getNextReq := func(resp *ListUsersResponse) *ListUsersProxyRequest {
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

// Lists the users in the account that contains this workspace, sorted by
// Databricks ID. Users in your identity provider who haven't been added to
// Databricks aren't included.
func (a *workspaceIamV2Impl) ListUsersProxyAll(ctx context.Context, request ListUsersProxyRequest) ([]User, error) {
	iterator := a.ListUsersProxy(ctx, request)
	return listing.ToSlice[User](ctx, iterator)
}

func (a *workspaceIamV2Impl) internalListUsersProxy(ctx context.Context, request ListUsersProxyRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0/identity/users"
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

// Lists workspace assignment details for the calling workspace. The response
// omits the per-principal entitlement fields (`entitlements` and
// `effective_entitlements`). To read the entitlements for a single principal,
// get that principal's assignment detail.
func (a *workspaceIamV2Impl) ListWorkspaceAssignmentDetailsProxy(ctx context.Context, request ListWorkspaceAssignmentDetailsProxyRequest) listing.Iterator[WorkspaceAssignmentDetail] {

	getNextPage := func(ctx context.Context, req ListWorkspaceAssignmentDetailsProxyRequest) (*ListWorkspaceAssignmentDetailsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListWorkspaceAssignmentDetailsProxy(ctx, req)
	}
	getItems := func(resp *ListWorkspaceAssignmentDetailsResponse) []WorkspaceAssignmentDetail {
		return resp.WorkspaceAssignmentDetails
	}
	getNextReq := func(resp *ListWorkspaceAssignmentDetailsResponse) *ListWorkspaceAssignmentDetailsProxyRequest {
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

// Lists workspace assignment details for the calling workspace. The response
// omits the per-principal entitlement fields (`entitlements` and
// `effective_entitlements`). To read the entitlements for a single principal,
// get that principal's assignment detail.
func (a *workspaceIamV2Impl) ListWorkspaceAssignmentDetailsProxyAll(ctx context.Context, request ListWorkspaceAssignmentDetailsProxyRequest) ([]WorkspaceAssignmentDetail, error) {
	iterator := a.ListWorkspaceAssignmentDetailsProxy(ctx, request)
	return listing.ToSlice[WorkspaceAssignmentDetail](ctx, iterator)
}

func (a *workspaceIamV2Impl) internalListWorkspaceAssignmentDetailsProxy(ctx context.Context, request ListWorkspaceAssignmentDetailsProxyRequest) (*ListWorkspaceAssignmentDetailsResponse, error) {
	var listWorkspaceAssignmentDetailsResponse ListWorkspaceAssignmentDetailsResponse
	path := "/api/2.0/identity/workspace-assignment-details"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAssignmentDetailsResponse)
	return &listWorkspaceAssignmentDetailsResponse, err
}

// Lists the principals assigned directly to this workspace. The response leaves
// `entitlements` and `effective_entitlements` empty. To read a principal's
// entitlements, get its assignment.
func (a *workspaceIamV2Impl) ListWorkspaceAssignmentsProxy(ctx context.Context, request ListWorkspaceAssignmentsProxyRequest) listing.Iterator[WorkspaceAssignment] {

	getNextPage := func(ctx context.Context, req ListWorkspaceAssignmentsProxyRequest) (*ListWorkspaceAssignmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListWorkspaceAssignmentsProxy(ctx, req)
	}
	getItems := func(resp *ListWorkspaceAssignmentsResponse) []WorkspaceAssignment {
		return resp.WorkspaceAssignments
	}
	getNextReq := func(resp *ListWorkspaceAssignmentsResponse) *ListWorkspaceAssignmentsProxyRequest {
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

// Lists the principals assigned directly to this workspace. The response leaves
// `entitlements` and `effective_entitlements` empty. To read a principal's
// entitlements, get its assignment.
func (a *workspaceIamV2Impl) ListWorkspaceAssignmentsProxyAll(ctx context.Context, request ListWorkspaceAssignmentsProxyRequest) ([]WorkspaceAssignment, error) {
	iterator := a.ListWorkspaceAssignmentsProxy(ctx, request)
	return listing.ToSlice[WorkspaceAssignment](ctx, iterator)
}

func (a *workspaceIamV2Impl) internalListWorkspaceAssignmentsProxy(ctx context.Context, request ListWorkspaceAssignmentsProxyRequest) (*ListWorkspaceAssignmentsResponse, error) {
	var listWorkspaceAssignmentsResponse ListWorkspaceAssignmentsResponse
	path := "/api/2.0/identity/workspace-assignments"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWorkspaceAssignmentsResponse)
	return &listWorkspaceAssignmentsResponse, err
}

func (a *workspaceIamV2Impl) ResolveGroupProxy(ctx context.Context, request ResolveGroupProxyRequest) (*ResolveGroupResponse, error) {
	var resolveGroupResponse ResolveGroupResponse
	path := "/api/2.0/identity/groups/resolve-by-external-id"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveGroupResponse)
	return &resolveGroupResponse, err
}

func (a *workspaceIamV2Impl) ResolveServicePrincipalProxy(ctx context.Context, request ResolveServicePrincipalProxyRequest) (*ResolveServicePrincipalResponse, error) {
	var resolveServicePrincipalResponse ResolveServicePrincipalResponse
	path := "/api/2.0/identity/service-principals/resolve-by-external-id"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveServicePrincipalResponse)
	return &resolveServicePrincipalResponse, err
}

func (a *workspaceIamV2Impl) ResolveUserProxy(ctx context.Context, request ResolveUserProxyRequest) (*ResolveUserResponse, error) {
	var resolveUserResponse ResolveUserResponse
	path := "/api/2.0/identity/users/resolve-by-external-id"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resolveUserResponse)
	return &resolveUserResponse, err
}

func (a *workspaceIamV2Impl) UpdateGroupProxy(ctx context.Context, request UpdateGroupProxyRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Group, &group)
	return &group, err
}

func (a *workspaceIamV2Impl) UpdateServicePrincipalProxy(ctx context.Context, request UpdateServicePrincipalProxyRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ServicePrincipal, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *workspaceIamV2Impl) UpdateUserProxy(ctx context.Context, request UpdateUserProxyRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.User, &user)
	return &user, err
}

func (a *workspaceIamV2Impl) UpdateWorkspaceAssignmentDetailProxy(ctx context.Context, request UpdateWorkspaceAssignmentDetailProxyRequest) (*WorkspaceAssignmentDetail, error) {
	var workspaceAssignmentDetail WorkspaceAssignmentDetail
	path := fmt.Sprintf("/api/2.0/identity/workspace-assignment-details/%v", request.PrincipalId)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAssignmentDetail, &workspaceAssignmentDetail)
	return &workspaceAssignmentDetail, err
}

func (a *workspaceIamV2Impl) UpdateWorkspaceAssignmentProxy(ctx context.Context, request UpdateWorkspaceAssignmentProxyRequest) (*WorkspaceAssignment, error) {
	var workspaceAssignment WorkspaceAssignment
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceAssignment, &workspaceAssignment)
	return &workspaceAssignment, err
}

func (a *workspaceIamV2Impl) UpdateWorkspaceIdentityDetail(ctx context.Context, request UpdateWorkspaceIdentityDetailRequest) (*WorkspaceIdentityDetail, error) {
	var workspaceIdentityDetail WorkspaceIdentityDetail
	path := fmt.Sprintf("/api/2.0/identity/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.WorkspaceIdentityDetail, &workspaceIdentityDetail)
	return &workspaceIdentityDetail, err
}
