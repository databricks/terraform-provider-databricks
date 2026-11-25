// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

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

// unexported type that holds implementations of just AccountMetastoreAssignments API methods
type accountMetastoreAssignmentsImpl struct {
	client *client.DatabricksClient
}

func (a *accountMetastoreAssignmentsImpl) Create(ctx context.Context, request AccountsCreateMetastoreAssignment) (*AccountsCreateMetastoreAssignmentResponse, error) {
	var accountsCreateMetastoreAssignmentResponse AccountsCreateMetastoreAssignmentResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &accountsCreateMetastoreAssignmentResponse)
	return &accountsCreateMetastoreAssignmentResponse, err
}

func (a *accountMetastoreAssignmentsImpl) Delete(ctx context.Context, request DeleteAccountMetastoreAssignmentRequest) (*AccountsDeleteMetastoreAssignmentResponse, error) {
	var accountsDeleteMetastoreAssignmentResponse AccountsDeleteMetastoreAssignmentResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &accountsDeleteMetastoreAssignmentResponse)
	return &accountsDeleteMetastoreAssignmentResponse, err
}

func (a *accountMetastoreAssignmentsImpl) Get(ctx context.Context, request GetAccountMetastoreAssignmentRequest) (*AccountsMetastoreAssignment, error) {
	var accountsMetastoreAssignment AccountsMetastoreAssignment
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastore", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &accountsMetastoreAssignment)
	return &accountsMetastoreAssignment, err
}

// Gets a list of all Databricks workspace IDs that have been assigned to given
// metastore.
func (a *accountMetastoreAssignmentsImpl) List(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) listing.Iterator[int64] {

	getNextPage := func(ctx context.Context, req ListAccountMetastoreAssignmentsRequest) (*ListAccountMetastoreAssignmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAccountMetastoreAssignmentsResponse) []int64 {
		return resp.WorkspaceIds
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Gets a list of all Databricks workspace IDs that have been assigned to given
// metastore.
func (a *accountMetastoreAssignmentsImpl) ListAll(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) ([]int64, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[int64](ctx, iterator)
}

func (a *accountMetastoreAssignmentsImpl) internalList(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) (*ListAccountMetastoreAssignmentsResponse, error) {
	var listAccountMetastoreAssignmentsResponse ListAccountMetastoreAssignmentsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/workspaces", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAccountMetastoreAssignmentsResponse)
	return &listAccountMetastoreAssignmentsResponse, err
}

func (a *accountMetastoreAssignmentsImpl) Update(ctx context.Context, request AccountsUpdateMetastoreAssignment) (*AccountsUpdateMetastoreAssignmentResponse, error) {
	var accountsUpdateMetastoreAssignmentResponse AccountsUpdateMetastoreAssignmentResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &accountsUpdateMetastoreAssignmentResponse)
	return &accountsUpdateMetastoreAssignmentResponse, err
}

// unexported type that holds implementations of just AccountMetastores API methods
type accountMetastoresImpl struct {
	client *client.DatabricksClient
}

func (a *accountMetastoresImpl) Create(ctx context.Context, request AccountsCreateMetastore) (*AccountsCreateMetastoreResponse, error) {
	var accountsCreateMetastoreResponse AccountsCreateMetastoreResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &accountsCreateMetastoreResponse)
	return &accountsCreateMetastoreResponse, err
}

func (a *accountMetastoresImpl) Delete(ctx context.Context, request DeleteAccountMetastoreRequest) (*AccountsDeleteMetastoreResponse, error) {
	var accountsDeleteMetastoreResponse AccountsDeleteMetastoreResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &accountsDeleteMetastoreResponse)
	return &accountsDeleteMetastoreResponse, err
}

func (a *accountMetastoresImpl) Get(ctx context.Context, request GetAccountMetastoreRequest) (*AccountsGetMetastoreResponse, error) {
	var accountsGetMetastoreResponse AccountsGetMetastoreResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &accountsGetMetastoreResponse)
	return &accountsGetMetastoreResponse, err
}

// Gets all Unity Catalog metastores associated with an account specified by ID.
func (a *accountMetastoresImpl) List(ctx context.Context) listing.Iterator[MetastoreInfo] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*AccountsListMetastoresResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *AccountsListMetastoresResponse) []MetastoreInfo {
		return resp.Metastores
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Gets all Unity Catalog metastores associated with an account specified by ID.
func (a *accountMetastoresImpl) ListAll(ctx context.Context) ([]MetastoreInfo, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[MetastoreInfo](ctx, iterator)
}

func (a *accountMetastoresImpl) internalList(ctx context.Context) (*AccountsListMetastoresResponse, error) {
	var accountsListMetastoresResponse AccountsListMetastoresResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &accountsListMetastoresResponse)
	return &accountsListMetastoresResponse, err
}

func (a *accountMetastoresImpl) Update(ctx context.Context, request AccountsUpdateMetastore) (*AccountsUpdateMetastoreResponse, error) {
	var accountsUpdateMetastoreResponse AccountsUpdateMetastoreResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &accountsUpdateMetastoreResponse)
	return &accountsUpdateMetastoreResponse, err
}

// unexported type that holds implementations of just AccountStorageCredentials API methods
type accountStorageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *accountStorageCredentialsImpl) Create(ctx context.Context, request AccountsCreateStorageCredential) (*AccountsCreateStorageCredentialInfo, error) {
	var accountsCreateStorageCredentialInfo AccountsCreateStorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &accountsCreateStorageCredentialInfo)
	return &accountsCreateStorageCredentialInfo, err
}

func (a *accountStorageCredentialsImpl) Delete(ctx context.Context, request DeleteAccountStorageCredentialRequest) (*AccountsDeleteStorageCredentialResponse, error) {
	var accountsDeleteStorageCredentialResponse AccountsDeleteStorageCredentialResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &accountsDeleteStorageCredentialResponse)
	return &accountsDeleteStorageCredentialResponse, err
}

func (a *accountStorageCredentialsImpl) Get(ctx context.Context, request GetAccountStorageCredentialRequest) (*AccountsStorageCredentialInfo, error) {
	var accountsStorageCredentialInfo AccountsStorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &accountsStorageCredentialInfo)
	return &accountsStorageCredentialInfo, err
}

// Gets a list of all storage credentials that have been assigned to given
// metastore.
func (a *accountStorageCredentialsImpl) List(ctx context.Context, request ListAccountStorageCredentialsRequest) listing.Iterator[StorageCredentialInfo] {

	getNextPage := func(ctx context.Context, req ListAccountStorageCredentialsRequest) (*ListAccountStorageCredentialsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAccountStorageCredentialsResponse) []StorageCredentialInfo {
		return resp.StorageCredentials
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Gets a list of all storage credentials that have been assigned to given
// metastore.
func (a *accountStorageCredentialsImpl) ListAll(ctx context.Context, request ListAccountStorageCredentialsRequest) ([]StorageCredentialInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[StorageCredentialInfo](ctx, iterator)
}

func (a *accountStorageCredentialsImpl) internalList(ctx context.Context, request ListAccountStorageCredentialsRequest) (*ListAccountStorageCredentialsResponse, error) {
	var listAccountStorageCredentialsResponse ListAccountStorageCredentialsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAccountStorageCredentialsResponse)
	return &listAccountStorageCredentialsResponse, err
}

func (a *accountStorageCredentialsImpl) Update(ctx context.Context, request AccountsUpdateStorageCredential) (*AccountsUpdateStorageCredentialResponse, error) {
	var accountsUpdateStorageCredentialResponse AccountsUpdateStorageCredentialResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &accountsUpdateStorageCredentialResponse)
	return &accountsUpdateStorageCredentialResponse, err
}

// unexported type that holds implementations of just AiGateway API methods
type aiGatewayImpl struct {
	client *client.DatabricksClient
}

func (a *aiGatewayImpl) CreateMcpService(ctx context.Context, request CreateMcpServiceRequest) (*McpService, error) {
	var mcpService McpService
	path := "/api/2.1/unity-catalog/mcp-services"
	queryParams := make(map[string]any)

	if request.McpServiceId != "" {
		queryParams["mcp_service_id"] = request.McpServiceId
	}

	if request.Parent != "" {
		queryParams["parent"] = request.Parent
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.McpService, &mcpService)
	return &mcpService, err
}

func (a *aiGatewayImpl) CreateMcpServiceUserMappedCredential(ctx context.Context, request CreateMcpServiceUserMappedCredentialRequest) (*McpServiceUserMappedCredential, error) {
	var mcpServiceUserMappedCredential McpServiceUserMappedCredential
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v/user-credentials", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Login, &mcpServiceUserMappedCredential)
	return &mcpServiceUserMappedCredential, err
}

func (a *aiGatewayImpl) CreateModelProviderService(ctx context.Context, request CreateModelProviderServiceRequest) (*ModelProviderService, error) {
	var modelProviderService ModelProviderService
	path := "/api/2.1/unity-catalog/model-provider-services"
	queryParams := make(map[string]any)

	if request.ModelProviderServiceId != "" {
		queryParams["model_provider_service_id"] = request.ModelProviderServiceId
	}

	if request.Parent != "" {
		queryParams["parent"] = request.Parent
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ModelProviderService, &modelProviderService)
	return &modelProviderService, err
}

func (a *aiGatewayImpl) CreateModelService(ctx context.Context, request CreateModelServiceRequest) (*ModelService, error) {
	var modelService ModelService
	path := "/api/2.1/unity-catalog/model-services"
	queryParams := make(map[string]any)

	if request.ModelServiceId != "" {
		queryParams["model_service_id"] = request.ModelServiceId
	}

	if request.Parent != "" {
		queryParams["parent"] = request.Parent
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ModelService, &modelService)
	return &modelService, err
}

func (a *aiGatewayImpl) CreateSkill(ctx context.Context, request CreateSkillRequest) (*Skill, error) {
	var skill Skill
	path := "/api/2.1/unity-catalog/skills"
	queryParams := make(map[string]any)

	if request.Parent != "" {
		queryParams["parent"] = request.Parent
	}

	if request.SkillId != "" {
		queryParams["skill_id"] = request.SkillId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Skill, &skill)
	return &skill, err
}

func (a *aiGatewayImpl) DeleteMcpService(ctx context.Context, request DeleteMcpServiceRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
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

func (a *aiGatewayImpl) DeleteMcpServiceUserMappedCredential(ctx context.Context, request DeleteMcpServiceUserMappedCredentialRequest) (*DeleteMcpServiceUserMappedCredentialResponse, error) {
	var deleteMcpServiceUserMappedCredentialResponse DeleteMcpServiceUserMappedCredentialResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v/user-credentials", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteMcpServiceUserMappedCredentialResponse)
	return &deleteMcpServiceUserMappedCredentialResponse, err
}

func (a *aiGatewayImpl) DeleteModelProviderService(ctx context.Context, request DeleteModelProviderServiceRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
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

func (a *aiGatewayImpl) DeleteModelService(ctx context.Context, request DeleteModelServiceRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
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

func (a *aiGatewayImpl) DeleteSkill(ctx context.Context, request DeleteSkillRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
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

func (a *aiGatewayImpl) FinalizeSkill(ctx context.Context, request FinalizeSkillRequest) (*Skill, error) {
	var skill Skill
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v/finalize", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &skill)
	return &skill, err
}

func (a *aiGatewayImpl) GetMcpService(ctx context.Context, request GetMcpServiceRequest) (*McpService, error) {
	var mcpService McpService
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &mcpService)
	return &mcpService, err
}

func (a *aiGatewayImpl) GetMcpServiceUserMappedCredential(ctx context.Context, request GetMcpServiceUserMappedCredentialRequest) (*McpServiceUserMappedCredential, error) {
	var mcpServiceUserMappedCredential McpServiceUserMappedCredential
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v/user-credentials", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &mcpServiceUserMappedCredential)
	return &mcpServiceUserMappedCredential, err
}

func (a *aiGatewayImpl) GetModelProviderService(ctx context.Context, request GetModelProviderServiceRequest) (*ModelProviderService, error) {
	var modelProviderService ModelProviderService
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &modelProviderService)
	return &modelProviderService, err
}

func (a *aiGatewayImpl) GetModelService(ctx context.Context, request GetModelServiceRequest) (*ModelService, error) {
	var modelService ModelService
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &modelService)
	return &modelService, err
}

func (a *aiGatewayImpl) GetSkill(ctx context.Context, request GetSkillRequest) (*Skill, error) {
	var skill Skill
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &skill)
	return &skill, err
}

// Lists the MCP services in a Unity Catalog schema. Provide `parent` as
// `schemas/{catalog}.{schema}`. Results are paginated; pass the returned
// `next_page_token` to fetch subsequent pages.
//
// Requires `USE_CATALOG` on the parent catalog and `USE_SCHEMA` on the parent
// schema. Only MCP services the caller can access (as owner or through
// `EXECUTE`, `READ_METADATA`, or `MANAGE`) are returned.
func (a *aiGatewayImpl) ListMcpServices(ctx context.Context, request ListMcpServicesRequest) listing.Iterator[McpService] {

	getNextPage := func(ctx context.Context, req ListMcpServicesRequest) (*ListMcpServicesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListMcpServices(ctx, req)
	}
	getItems := func(resp *ListMcpServicesResponse) []McpService {
		return resp.McpServices
	}
	getNextReq := func(resp *ListMcpServicesResponse) *ListMcpServicesRequest {
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

// Lists the MCP services in a Unity Catalog schema. Provide `parent` as
// `schemas/{catalog}.{schema}`. Results are paginated; pass the returned
// `next_page_token` to fetch subsequent pages.
//
// Requires `USE_CATALOG` on the parent catalog and `USE_SCHEMA` on the parent
// schema. Only MCP services the caller can access (as owner or through
// `EXECUTE`, `READ_METADATA`, or `MANAGE`) are returned.
func (a *aiGatewayImpl) ListMcpServicesAll(ctx context.Context, request ListMcpServicesRequest) ([]McpService, error) {
	iterator := a.ListMcpServices(ctx, request)
	return listing.ToSlice[McpService](ctx, iterator)
}

func (a *aiGatewayImpl) internalListMcpServices(ctx context.Context, request ListMcpServicesRequest) (*ListMcpServicesResponse, error) {
	var listMcpServicesResponse ListMcpServicesResponse
	path := "/api/2.1/unity-catalog/mcp-services"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listMcpServicesResponse)
	return &listMcpServicesResponse, err
}

// Lists the model provider services in a Unity Catalog schema. Provide `parent`
// as `schemas/{catalog}.{schema}`. Results are paginated; pass the returned
// `next_page_token` to fetch subsequent pages.
//
// Requires `USE_CATALOG` on the parent catalog and `USE_SCHEMA` on the parent
// schema. Only model provider services the caller can access (as owner or
// through `EXECUTE`, `READ_METADATA`, or `MANAGE`) are returned.
func (a *aiGatewayImpl) ListModelProviderServices(ctx context.Context, request ListModelProviderServicesRequest) listing.Iterator[ModelProviderService] {

	getNextPage := func(ctx context.Context, req ListModelProviderServicesRequest) (*ListModelProviderServicesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListModelProviderServices(ctx, req)
	}
	getItems := func(resp *ListModelProviderServicesResponse) []ModelProviderService {
		return resp.ModelProviderServices
	}
	getNextReq := func(resp *ListModelProviderServicesResponse) *ListModelProviderServicesRequest {
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

// Lists the model provider services in a Unity Catalog schema. Provide `parent`
// as `schemas/{catalog}.{schema}`. Results are paginated; pass the returned
// `next_page_token` to fetch subsequent pages.
//
// Requires `USE_CATALOG` on the parent catalog and `USE_SCHEMA` on the parent
// schema. Only model provider services the caller can access (as owner or
// through `EXECUTE`, `READ_METADATA`, or `MANAGE`) are returned.
func (a *aiGatewayImpl) ListModelProviderServicesAll(ctx context.Context, request ListModelProviderServicesRequest) ([]ModelProviderService, error) {
	iterator := a.ListModelProviderServices(ctx, request)
	return listing.ToSlice[ModelProviderService](ctx, iterator)
}

func (a *aiGatewayImpl) internalListModelProviderServices(ctx context.Context, request ListModelProviderServicesRequest) (*ListModelProviderServicesResponse, error) {
	var listModelProviderServicesResponse ListModelProviderServicesResponse
	path := "/api/2.1/unity-catalog/model-provider-services"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listModelProviderServicesResponse)
	return &listModelProviderServicesResponse, err
}

// Lists the model services in a Unity Catalog schema. Provide `parent` as
// `schemas/{catalog}.{schema}`. Results are paginated; pass the returned
// `next_page_token` to fetch subsequent pages.
//
// Requires `USE_CATALOG` on the parent catalog and `USE_SCHEMA` on the parent
// schema. Only model services the caller can access (as owner or through
// `EXECUTE`, `READ_METADATA`, or `MANAGE`) are returned by default.
func (a *aiGatewayImpl) ListModelServices(ctx context.Context, request ListModelServicesRequest) listing.Iterator[ModelService] {

	getNextPage := func(ctx context.Context, req ListModelServicesRequest) (*ListModelServicesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListModelServices(ctx, req)
	}
	getItems := func(resp *ListModelServicesResponse) []ModelService {
		return resp.ModelServices
	}
	getNextReq := func(resp *ListModelServicesResponse) *ListModelServicesRequest {
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

// Lists the model services in a Unity Catalog schema. Provide `parent` as
// `schemas/{catalog}.{schema}`. Results are paginated; pass the returned
// `next_page_token` to fetch subsequent pages.
//
// Requires `USE_CATALOG` on the parent catalog and `USE_SCHEMA` on the parent
// schema. Only model services the caller can access (as owner or through
// `EXECUTE`, `READ_METADATA`, or `MANAGE`) are returned by default.
func (a *aiGatewayImpl) ListModelServicesAll(ctx context.Context, request ListModelServicesRequest) ([]ModelService, error) {
	iterator := a.ListModelServices(ctx, request)
	return listing.ToSlice[ModelService](ctx, iterator)
}

func (a *aiGatewayImpl) internalListModelServices(ctx context.Context, request ListModelServicesRequest) (*ListModelServicesResponse, error) {
	var listModelServicesResponse ListModelServicesResponse
	path := "/api/2.1/unity-catalog/model-services"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listModelServicesResponse)
	return &listModelServicesResponse, err
}

// Lists skills in a Unity Catalog schema. Provide `parent` as
// `schemas/{catalog}.{schema}`. Results are paginated; pass the returned
// `next_page_token` to fetch subsequent pages.
//
// Requires `USE_CATALOG` on the parent catalog and `USE_SCHEMA` on the parent
// schema. Only skills the caller can access as owner or through `READ_VOLUME`,
// `READ_METADATA`, or `MANAGE` are returned.
func (a *aiGatewayImpl) ListSkills(ctx context.Context, request ListSkillsRequest) listing.Iterator[Skill] {

	getNextPage := func(ctx context.Context, req ListSkillsRequest) (*ListSkillsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSkills(ctx, req)
	}
	getItems := func(resp *ListSkillsResponse) []Skill {
		return resp.Skills
	}
	getNextReq := func(resp *ListSkillsResponse) *ListSkillsRequest {
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

// Lists skills in a Unity Catalog schema. Provide `parent` as
// `schemas/{catalog}.{schema}`. Results are paginated; pass the returned
// `next_page_token` to fetch subsequent pages.
//
// Requires `USE_CATALOG` on the parent catalog and `USE_SCHEMA` on the parent
// schema. Only skills the caller can access as owner or through `READ_VOLUME`,
// `READ_METADATA`, or `MANAGE` are returned.
func (a *aiGatewayImpl) ListSkillsAll(ctx context.Context, request ListSkillsRequest) ([]Skill, error) {
	iterator := a.ListSkills(ctx, request)
	return listing.ToSlice[Skill](ctx, iterator)
}

func (a *aiGatewayImpl) internalListSkills(ctx context.Context, request ListSkillsRequest) (*ListSkillsResponse, error) {
	var listSkillsResponse ListSkillsResponse
	path := "/api/2.1/unity-catalog/skills"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSkillsResponse)
	return &listSkillsResponse, err
}

func (a *aiGatewayImpl) UpdateMcpService(ctx context.Context, request UpdateMcpServiceRequest) (*McpService, error) {
	var mcpService McpService
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.McpService, &mcpService)
	return &mcpService, err
}

func (a *aiGatewayImpl) UpdateModelProviderService(ctx context.Context, request UpdateModelProviderServiceRequest) (*ModelProviderService, error) {
	var modelProviderService ModelProviderService
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ModelProviderService, &modelProviderService)
	return &modelProviderService, err
}

func (a *aiGatewayImpl) UpdateModelService(ctx context.Context, request UpdateModelServiceRequest) (*ModelService, error) {
	var modelService ModelService
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ModelService, &modelService)
	return &modelService, err
}

func (a *aiGatewayImpl) UpdateSkill(ctx context.Context, request UpdateSkillRequest) (*Skill, error) {
	var skill Skill
	path := fmt.Sprintf("/api/2.1/unity-catalog/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Skill, &skill)
	return &skill, err
}

// unexported type that holds implementations of just ArtifactAllowlists API methods
type artifactAllowlistsImpl struct {
	client *client.DatabricksClient
}

func (a *artifactAllowlistsImpl) Get(ctx context.Context, request GetArtifactAllowlistRequest) (*ArtifactAllowlistInfo, error) {
	var artifactAllowlistInfo ArtifactAllowlistInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/artifact-allowlists/%v", request.ArtifactType)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &artifactAllowlistInfo)
	return &artifactAllowlistInfo, err
}

func (a *artifactAllowlistsImpl) Update(ctx context.Context, request SetArtifactAllowlist) (*ArtifactAllowlistInfo, error) {
	var artifactAllowlistInfo ArtifactAllowlistInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/artifact-allowlists/%v", request.ArtifactType)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &artifactAllowlistInfo)
	return &artifactAllowlistInfo, err
}

// unexported type that holds implementations of just Catalogs API methods
type catalogsImpl struct {
	client *client.DatabricksClient
}

func (a *catalogsImpl) Create(ctx context.Context, request CreateCatalog) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := "/api/2.1/unity-catalog/catalogs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &catalogInfo)
	return &catalogInfo, err
}

func (a *catalogsImpl) Delete(ctx context.Context, request DeleteCatalogRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
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

func (a *catalogsImpl) Get(ctx context.Context, request GetCatalogRequest) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &catalogInfo)
	return &catalogInfo, err
}

// Gets an array of catalogs in the metastore. If the caller is the metastore
// admin, all catalogs will be retrieved. Otherwise, only catalogs owned by the
// caller (or for which the caller has the **USE_CATALOG** privilege) will be
// retrieved. There is no guarantee of a specific ordering of the elements in
// the array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *catalogsImpl) List(ctx context.Context, request ListCatalogsRequest) listing.Iterator[CatalogInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListCatalogsRequest) (*ListCatalogsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCatalogsResponse) []CatalogInfo {
		return resp.Catalogs
	}
	getNextReq := func(resp *ListCatalogsResponse) *ListCatalogsRequest {
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

// Gets an array of catalogs in the metastore. If the caller is the metastore
// admin, all catalogs will be retrieved. Otherwise, only catalogs owned by the
// caller (or for which the caller has the **USE_CATALOG** privilege) will be
// retrieved. There is no guarantee of a specific ordering of the elements in
// the array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *catalogsImpl) ListAll(ctx context.Context, request ListCatalogsRequest) ([]CatalogInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CatalogInfo](ctx, iterator)
}

func (a *catalogsImpl) internalList(ctx context.Context, request ListCatalogsRequest) (*ListCatalogsResponse, error) {
	var listCatalogsResponse ListCatalogsResponse
	path := "/api/2.1/unity-catalog/catalogs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCatalogsResponse)
	return &listCatalogsResponse, err
}

func (a *catalogsImpl) Update(ctx context.Context, request UpdateCatalog) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &catalogInfo)
	return &catalogInfo, err
}

// unexported type that holds implementations of just Connections API methods
type connectionsImpl struct {
	client *client.DatabricksClient
}

func (a *connectionsImpl) Create(ctx context.Context, request CreateConnection) (*ConnectionInfo, error) {
	var connectionInfo ConnectionInfo
	path := "/api/2.1/unity-catalog/connections"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &connectionInfo)
	return &connectionInfo, err
}

func (a *connectionsImpl) Delete(ctx context.Context, request DeleteConnectionRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
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

func (a *connectionsImpl) Get(ctx context.Context, request GetConnectionRequest) (*ConnectionInfo, error) {
	var connectionInfo ConnectionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &connectionInfo)
	return &connectionInfo, err
}

// List all connections.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *connectionsImpl) List(ctx context.Context, request ListConnectionsRequest) listing.Iterator[ConnectionInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListConnectionsRequest) (*ListConnectionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListConnectionsResponse) []ConnectionInfo {
		return resp.Connections
	}
	getNextReq := func(resp *ListConnectionsResponse) *ListConnectionsRequest {
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

// List all connections.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *connectionsImpl) ListAll(ctx context.Context, request ListConnectionsRequest) ([]ConnectionInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ConnectionInfo](ctx, iterator)
}

func (a *connectionsImpl) internalList(ctx context.Context, request ListConnectionsRequest) (*ListConnectionsResponse, error) {
	var listConnectionsResponse ListConnectionsResponse
	path := "/api/2.1/unity-catalog/connections"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listConnectionsResponse)
	return &listConnectionsResponse, err
}

func (a *connectionsImpl) Update(ctx context.Context, request UpdateConnection) (*ConnectionInfo, error) {
	var connectionInfo ConnectionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &connectionInfo)
	return &connectionInfo, err
}

// unexported type that holds implementations of just Credentials API methods
type credentialsImpl struct {
	client *client.DatabricksClient
}

func (a *credentialsImpl) CreateCredential(ctx context.Context, request CreateCredentialRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := "/api/2.1/unity-catalog/credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &credentialInfo)
	return &credentialInfo, err
}

func (a *credentialsImpl) DeleteCredential(ctx context.Context, request DeleteCredentialRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
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

func (a *credentialsImpl) GenerateTemporaryServiceCredential(ctx context.Context, request GenerateTemporaryServiceCredentialRequest) (*TemporaryCredentials, error) {
	var temporaryCredentials TemporaryCredentials
	path := "/api/2.1/unity-catalog/temporary-service-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &temporaryCredentials)
	return &temporaryCredentials, err
}

func (a *credentialsImpl) GetCredential(ctx context.Context, request GetCredentialRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &credentialInfo)
	return &credentialInfo, err
}

// Gets an array of credentials (as __CredentialInfo__ objects).
//
// The array is limited to only the credentials that the caller has permission
// to access. If the caller is a metastore admin, retrieval of credentials is
// unrestricted. There is no guarantee of a specific ordering of the elements in
// the array.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *credentialsImpl) ListCredentials(ctx context.Context, request ListCredentialsRequest) listing.Iterator[CredentialInfo] {

	getNextPage := func(ctx context.Context, req ListCredentialsRequest) (*ListCredentialsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListCredentials(ctx, req)
	}
	getItems := func(resp *ListCredentialsResponse) []CredentialInfo {
		return resp.Credentials
	}
	getNextReq := func(resp *ListCredentialsResponse) *ListCredentialsRequest {
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

// Gets an array of credentials (as __CredentialInfo__ objects).
//
// The array is limited to only the credentials that the caller has permission
// to access. If the caller is a metastore admin, retrieval of credentials is
// unrestricted. There is no guarantee of a specific ordering of the elements in
// the array.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *credentialsImpl) ListCredentialsAll(ctx context.Context, request ListCredentialsRequest) ([]CredentialInfo, error) {
	iterator := a.ListCredentials(ctx, request)
	return listing.ToSlice[CredentialInfo](ctx, iterator)
}

func (a *credentialsImpl) internalListCredentials(ctx context.Context, request ListCredentialsRequest) (*ListCredentialsResponse, error) {
	var listCredentialsResponse ListCredentialsResponse
	path := "/api/2.1/unity-catalog/credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCredentialsResponse)
	return &listCredentialsResponse, err
}

func (a *credentialsImpl) UpdateCredential(ctx context.Context, request UpdateCredentialRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &credentialInfo)
	return &credentialInfo, err
}

func (a *credentialsImpl) ValidateCredential(ctx context.Context, request ValidateCredentialRequest) (*ValidateCredentialResponse, error) {
	var validateCredentialResponse ValidateCredentialResponse
	path := "/api/2.1/unity-catalog/validate-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &validateCredentialResponse)
	return &validateCredentialResponse, err
}

// unexported type that holds implementations of just EntityTagAssignments API methods
type entityTagAssignmentsImpl struct {
	client *client.DatabricksClient
}

func (a *entityTagAssignmentsImpl) Create(ctx context.Context, request CreateEntityTagAssignmentRequest) (*EntityTagAssignment, error) {
	var entityTagAssignment EntityTagAssignment
	path := "/api/2.1/unity-catalog/entity-tag-assignments"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.TagAssignment, &entityTagAssignment)
	return &entityTagAssignment, err
}

func (a *entityTagAssignmentsImpl) Delete(ctx context.Context, request DeleteEntityTagAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/entity-tag-assignments/%v/%v/tags/%v", request.EntityType, request.EntityName, request.TagKey)
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

func (a *entityTagAssignmentsImpl) Get(ctx context.Context, request GetEntityTagAssignmentRequest) (*EntityTagAssignment, error) {
	var entityTagAssignment EntityTagAssignment
	path := fmt.Sprintf("/api/2.1/unity-catalog/entity-tag-assignments/%v/%v/tags/%v", request.EntityType, request.EntityName, request.TagKey)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &entityTagAssignment)
	return &entityTagAssignment, err
}

// List tag assignments for an Unity Catalog entity
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *entityTagAssignmentsImpl) List(ctx context.Context, request ListEntityTagAssignmentsRequest) listing.Iterator[EntityTagAssignment] {

	getNextPage := func(ctx context.Context, req ListEntityTagAssignmentsRequest) (*ListEntityTagAssignmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListEntityTagAssignmentsResponse) []EntityTagAssignment {
		return resp.TagAssignments
	}
	getNextReq := func(resp *ListEntityTagAssignmentsResponse) *ListEntityTagAssignmentsRequest {
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

// List tag assignments for an Unity Catalog entity
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *entityTagAssignmentsImpl) ListAll(ctx context.Context, request ListEntityTagAssignmentsRequest) ([]EntityTagAssignment, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[EntityTagAssignment](ctx, iterator)
}

func (a *entityTagAssignmentsImpl) internalList(ctx context.Context, request ListEntityTagAssignmentsRequest) (*ListEntityTagAssignmentsResponse, error) {
	var listEntityTagAssignmentsResponse ListEntityTagAssignmentsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/entity-tag-assignments/%v/%v/tags", request.EntityType, request.EntityName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listEntityTagAssignmentsResponse)
	return &listEntityTagAssignmentsResponse, err
}

func (a *entityTagAssignmentsImpl) Update(ctx context.Context, request UpdateEntityTagAssignmentRequest) (*EntityTagAssignment, error) {
	var entityTagAssignment EntityTagAssignment
	path := fmt.Sprintf("/api/2.1/unity-catalog/entity-tag-assignments/%v/%v/tags/%v", request.EntityType, request.EntityName, request.TagKey)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.TagAssignment, &entityTagAssignment)
	return &entityTagAssignment, err
}

// unexported type that holds implementations of just ExternalLineage API methods
type externalLineageImpl struct {
	client *client.DatabricksClient
}

func (a *externalLineageImpl) CreateExternalLineageRelationship(ctx context.Context, request CreateExternalLineageRelationshipRequest) (*ExternalLineageRelationship, error) {
	var externalLineageRelationship ExternalLineageRelationship
	path := "/api/2.0/lineage-tracking/external-lineage"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ExternalLineageRelationship, &externalLineageRelationship)
	return &externalLineageRelationship, err
}

func (a *externalLineageImpl) DeleteExternalLineageRelationship(ctx context.Context, request DeleteExternalLineageRelationshipRequest) error {
	path := "/api/2.0/lineage-tracking/external-lineage"
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

// Lists external lineage relationships of a Databricks object or external
// metadata given a supplied direction.
func (a *externalLineageImpl) ListExternalLineageRelationships(ctx context.Context, request ListExternalLineageRelationshipsRequest) listing.Iterator[ExternalLineageInfo] {

	getNextPage := func(ctx context.Context, req ListExternalLineageRelationshipsRequest) (*ListExternalLineageRelationshipsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListExternalLineageRelationships(ctx, req)
	}
	getItems := func(resp *ListExternalLineageRelationshipsResponse) []ExternalLineageInfo {
		return resp.ExternalLineageRelationships
	}
	getNextReq := func(resp *ListExternalLineageRelationshipsResponse) *ListExternalLineageRelationshipsRequest {
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

// Lists external lineage relationships of a Databricks object or external
// metadata given a supplied direction.
func (a *externalLineageImpl) ListExternalLineageRelationshipsAll(ctx context.Context, request ListExternalLineageRelationshipsRequest) ([]ExternalLineageInfo, error) {
	iterator := a.ListExternalLineageRelationships(ctx, request)
	return listing.ToSlice[ExternalLineageInfo](ctx, iterator)
}

func (a *externalLineageImpl) internalListExternalLineageRelationships(ctx context.Context, request ListExternalLineageRelationshipsRequest) (*ListExternalLineageRelationshipsResponse, error) {
	var listExternalLineageRelationshipsResponse ListExternalLineageRelationshipsResponse
	path := "/api/2.0/lineage-tracking/external-lineage"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExternalLineageRelationshipsResponse)
	return &listExternalLineageRelationshipsResponse, err
}

func (a *externalLineageImpl) UpdateExternalLineageRelationship(ctx context.Context, request UpdateExternalLineageRelationshipRequest) (*ExternalLineageRelationship, error) {
	var externalLineageRelationship ExternalLineageRelationship
	path := "/api/2.0/lineage-tracking/external-lineage"
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ExternalLineageRelationship, &externalLineageRelationship)
	return &externalLineageRelationship, err
}

// unexported type that holds implementations of just ExternalLocations API methods
type externalLocationsImpl struct {
	client *client.DatabricksClient
}

func (a *externalLocationsImpl) Create(ctx context.Context, request CreateExternalLocation) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := "/api/2.1/unity-catalog/external-locations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

func (a *externalLocationsImpl) Delete(ctx context.Context, request DeleteExternalLocationRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
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

func (a *externalLocationsImpl) Get(ctx context.Context, request GetExternalLocationRequest) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

// Gets an array of external locations (__ExternalLocationInfo__ objects) from
// the metastore. The caller must be a metastore admin, the owner of the
// external location, or a user that has some privilege on the external
// location. There is no guarantee of a specific ordering of the elements in the
// array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *externalLocationsImpl) List(ctx context.Context, request ListExternalLocationsRequest) listing.Iterator[ExternalLocationInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListExternalLocationsRequest) (*ListExternalLocationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListExternalLocationsResponse) []ExternalLocationInfo {
		return resp.ExternalLocations
	}
	getNextReq := func(resp *ListExternalLocationsResponse) *ListExternalLocationsRequest {
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

// Gets an array of external locations (__ExternalLocationInfo__ objects) from
// the metastore. The caller must be a metastore admin, the owner of the
// external location, or a user that has some privilege on the external
// location. There is no guarantee of a specific ordering of the elements in the
// array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *externalLocationsImpl) ListAll(ctx context.Context, request ListExternalLocationsRequest) ([]ExternalLocationInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ExternalLocationInfo](ctx, iterator)
}

func (a *externalLocationsImpl) internalList(ctx context.Context, request ListExternalLocationsRequest) (*ListExternalLocationsResponse, error) {
	var listExternalLocationsResponse ListExternalLocationsResponse
	path := "/api/2.1/unity-catalog/external-locations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExternalLocationsResponse)
	return &listExternalLocationsResponse, err
}

func (a *externalLocationsImpl) Update(ctx context.Context, request UpdateExternalLocation) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

// unexported type that holds implementations of just ExternalMetadata API methods
type externalMetadataImpl struct {
	client *client.DatabricksClient
}

func (a *externalMetadataImpl) CreateExternalMetadata(ctx context.Context, request CreateExternalMetadataRequest) (*ExternalMetadata, error) {
	var externalMetadata ExternalMetadata
	path := "/api/2.0/lineage-tracking/external-metadata"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ExternalMetadata, &externalMetadata)
	return &externalMetadata, err
}

func (a *externalMetadataImpl) DeleteExternalMetadata(ctx context.Context, request DeleteExternalMetadataRequest) error {
	path := fmt.Sprintf("/api/2.0/lineage-tracking/external-metadata/%v", request.Name)
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

func (a *externalMetadataImpl) GetExternalMetadata(ctx context.Context, request GetExternalMetadataRequest) (*ExternalMetadata, error) {
	var externalMetadata ExternalMetadata
	path := fmt.Sprintf("/api/2.0/lineage-tracking/external-metadata/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &externalMetadata)
	return &externalMetadata, err
}

// Gets an array of external metadata objects in the metastore. If the caller is
// the metastore admin, all external metadata objects will be retrieved.
// Otherwise, only external metadata objects that the caller has **BROWSE** on
// will be retrieved. There is no guarantee of a specific ordering of the
// elements in the array.
func (a *externalMetadataImpl) ListExternalMetadata(ctx context.Context, request ListExternalMetadataRequest) listing.Iterator[ExternalMetadata] {

	getNextPage := func(ctx context.Context, req ListExternalMetadataRequest) (*ListExternalMetadataResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListExternalMetadata(ctx, req)
	}
	getItems := func(resp *ListExternalMetadataResponse) []ExternalMetadata {
		return resp.ExternalMetadata
	}
	getNextReq := func(resp *ListExternalMetadataResponse) *ListExternalMetadataRequest {
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

// Gets an array of external metadata objects in the metastore. If the caller is
// the metastore admin, all external metadata objects will be retrieved.
// Otherwise, only external metadata objects that the caller has **BROWSE** on
// will be retrieved. There is no guarantee of a specific ordering of the
// elements in the array.
func (a *externalMetadataImpl) ListExternalMetadataAll(ctx context.Context, request ListExternalMetadataRequest) ([]ExternalMetadata, error) {
	iterator := a.ListExternalMetadata(ctx, request)
	return listing.ToSlice[ExternalMetadata](ctx, iterator)
}

func (a *externalMetadataImpl) internalListExternalMetadata(ctx context.Context, request ListExternalMetadataRequest) (*ListExternalMetadataResponse, error) {
	var listExternalMetadataResponse ListExternalMetadataResponse
	path := "/api/2.0/lineage-tracking/external-metadata"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExternalMetadataResponse)
	return &listExternalMetadataResponse, err
}

func (a *externalMetadataImpl) UpdateExternalMetadata(ctx context.Context, request UpdateExternalMetadataRequest) (*ExternalMetadata, error) {
	var externalMetadata ExternalMetadata
	path := fmt.Sprintf("/api/2.0/lineage-tracking/external-metadata/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ExternalMetadata, &externalMetadata)
	return &externalMetadata, err
}

// unexported type that holds implementations of just Functions API methods
type functionsImpl struct {
	client *client.DatabricksClient
}

func (a *functionsImpl) Create(ctx context.Context, request CreateFunctionRequest) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := "/api/2.1/unity-catalog/functions"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &functionInfo)
	return &functionInfo, err
}

func (a *functionsImpl) Delete(ctx context.Context, request DeleteFunctionRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *functionsImpl) Get(ctx context.Context, request GetFunctionRequest) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &functionInfo)
	return &functionInfo, err
}

// List functions within the specified parent catalog and schema. If the user is
// a metastore admin, all functions are returned in the output list. Otherwise,
// the user must have the **USE_CATALOG** privilege on the catalog and the
// **USE_SCHEMA** privilege on the schema, and the output list contains only
// functions for which either the user has the **EXECUTE** privilege or the user
// is the owner. There is no guarantee of a specific ordering of the elements in
// the array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *functionsImpl) List(ctx context.Context, request ListFunctionsRequest) listing.Iterator[FunctionInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListFunctionsRequest) (*ListFunctionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFunctionsResponse) []FunctionInfo {
		return resp.Functions
	}
	getNextReq := func(resp *ListFunctionsResponse) *ListFunctionsRequest {
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

// List functions within the specified parent catalog and schema. If the user is
// a metastore admin, all functions are returned in the output list. Otherwise,
// the user must have the **USE_CATALOG** privilege on the catalog and the
// **USE_SCHEMA** privilege on the schema, and the output list contains only
// functions for which either the user has the **EXECUTE** privilege or the user
// is the owner. There is no guarantee of a specific ordering of the elements in
// the array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *functionsImpl) ListAll(ctx context.Context, request ListFunctionsRequest) ([]FunctionInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FunctionInfo](ctx, iterator)
}

func (a *functionsImpl) internalList(ctx context.Context, request ListFunctionsRequest) (*ListFunctionsResponse, error) {
	var listFunctionsResponse ListFunctionsResponse
	path := "/api/2.1/unity-catalog/functions"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFunctionsResponse)
	return &listFunctionsResponse, err
}

func (a *functionsImpl) Update(ctx context.Context, request UpdateFunction) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &functionInfo)
	return &functionInfo, err
}

// unexported type that holds implementations of just Grants API methods
type grantsImpl struct {
	client *client.DatabricksClient
}

func (a *grantsImpl) Get(ctx context.Context, request GetGrantRequest) (*GetPermissionsResponse, error) {
	var getPermissionsResponse GetPermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPermissionsResponse)
	return &getPermissionsResponse, err
}

func (a *grantsImpl) GetEffective(ctx context.Context, request GetEffectiveRequest) (*EffectivePermissionsList, error) {
	var effectivePermissionsList EffectivePermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/effective-permissions/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &effectivePermissionsList)
	return &effectivePermissionsList, err
}

// Lists the privilege assignments for a securable. Does not include inherited
// privileges. Paginated version of Get Permissions API.
func (a *grantsImpl) List(ctx context.Context, request ListPrivilegeAssignmentsRequest) listing.Iterator[PrivilegeAssignment] {

	getNextPage := func(ctx context.Context, req ListPrivilegeAssignmentsRequest) (*ListPrivilegeAssignmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListPrivilegeAssignmentsResponse) []PrivilegeAssignment {
		return resp.PrivilegeAssignments
	}
	getNextReq := func(resp *ListPrivilegeAssignmentsResponse) *ListPrivilegeAssignmentsRequest {
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

// Lists the privilege assignments for a securable. Does not include inherited
// privileges. Paginated version of Get Permissions API.
func (a *grantsImpl) ListAll(ctx context.Context, request ListPrivilegeAssignmentsRequest) ([]PrivilegeAssignment, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PrivilegeAssignment](ctx, iterator)
}

func (a *grantsImpl) internalList(ctx context.Context, request ListPrivilegeAssignmentsRequest) (*ListPrivilegeAssignmentsResponse, error) {
	var listPrivilegeAssignmentsResponse ListPrivilegeAssignmentsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/privilege-assignments/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listPrivilegeAssignmentsResponse)
	return &listPrivilegeAssignmentsResponse, err
}

// Lists the effective privilege assignments for a securable. Includes inherited
// privileges. Paginated version of Get Effective Permissions API.
func (a *grantsImpl) ListEffective(ctx context.Context, request ListEffectivePrivilegeAssignmentsRequest) listing.Iterator[EffectivePrivilegeAssignment] {

	getNextPage := func(ctx context.Context, req ListEffectivePrivilegeAssignmentsRequest) (*ListEffectivePrivilegeAssignmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListEffective(ctx, req)
	}
	getItems := func(resp *ListEffectivePrivilegeAssignmentsResponse) []EffectivePrivilegeAssignment {
		return resp.EffectivePrivilegeAssignments
	}
	getNextReq := func(resp *ListEffectivePrivilegeAssignmentsResponse) *ListEffectivePrivilegeAssignmentsRequest {
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

// Lists the effective privilege assignments for a securable. Includes inherited
// privileges. Paginated version of Get Effective Permissions API.
func (a *grantsImpl) ListEffectiveAll(ctx context.Context, request ListEffectivePrivilegeAssignmentsRequest) ([]EffectivePrivilegeAssignment, error) {
	iterator := a.ListEffective(ctx, request)
	return listing.ToSlice[EffectivePrivilegeAssignment](ctx, iterator)
}

func (a *grantsImpl) internalListEffective(ctx context.Context, request ListEffectivePrivilegeAssignmentsRequest) (*ListEffectivePrivilegeAssignmentsResponse, error) {
	var listEffectivePrivilegeAssignmentsResponse ListEffectivePrivilegeAssignmentsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/effective-privilege-assignments/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listEffectivePrivilegeAssignmentsResponse)
	return &listEffectivePrivilegeAssignmentsResponse, err
}

func (a *grantsImpl) Update(ctx context.Context, request UpdatePermissions) (*UpdatePermissionsResponse, error) {
	var updatePermissionsResponse UpdatePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updatePermissionsResponse)
	return &updatePermissionsResponse, err
}

// unexported type that holds implementations of just Metastores API methods
type metastoresImpl struct {
	client *client.DatabricksClient
}

func (a *metastoresImpl) Assign(ctx context.Context, request CreateMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
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

func (a *metastoresImpl) Create(ctx context.Context, request CreateMetastore) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := "/api/2.1/unity-catalog/metastores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) Current(ctx context.Context) (*MetastoreAssignment, error) {
	var metastoreAssignment MetastoreAssignment
	path := "/api/2.1/unity-catalog/current-metastore-assignment"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &metastoreAssignment)
	return &metastoreAssignment, err
}

func (a *metastoresImpl) Delete(ctx context.Context, request DeleteMetastoreRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
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

func (a *metastoresImpl) Get(ctx context.Context, request GetMetastoreRequest) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &metastoreInfo)
	return &metastoreInfo, err
}

// Gets an array of the available metastores (as __MetastoreInfo__ objects). The
// caller must be an admin to retrieve this info. There is no guarantee of a
// specific ordering of the elements in the array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *metastoresImpl) List(ctx context.Context, request ListMetastoresRequest) listing.Iterator[MetastoreInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListMetastoresRequest) (*ListMetastoresResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListMetastoresResponse) []MetastoreInfo {
		return resp.Metastores
	}
	getNextReq := func(resp *ListMetastoresResponse) *ListMetastoresRequest {
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

// Gets an array of the available metastores (as __MetastoreInfo__ objects). The
// caller must be an admin to retrieve this info. There is no guarantee of a
// specific ordering of the elements in the array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *metastoresImpl) ListAll(ctx context.Context, request ListMetastoresRequest) ([]MetastoreInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[MetastoreInfo](ctx, iterator)
}

func (a *metastoresImpl) internalList(ctx context.Context, request ListMetastoresRequest) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := "/api/2.1/unity-catalog/metastores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *metastoresImpl) Summary(ctx context.Context) (*GetMetastoreSummaryResponse, error) {
	var getMetastoreSummaryResponse GetMetastoreSummaryResponse
	path := "/api/2.1/unity-catalog/metastore_summary"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getMetastoreSummaryResponse)
	return &getMetastoreSummaryResponse, err
}

func (a *metastoresImpl) Unassign(ctx context.Context, request UnassignRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
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

func (a *metastoresImpl) UndropMetastoreById(ctx context.Context, request UndropMetastoreByIdRequest) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := "/api/2.1/unity-catalog/metastores:undropById"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) Update(ctx context.Context, request UpdateMetastore) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) UpdateAssignment(ctx context.Context, request UpdateMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
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

// unexported type that holds implementations of just ModelVersions API methods
type modelVersionsImpl struct {
	client *client.DatabricksClient
}

func (a *modelVersionsImpl) Delete(ctx context.Context, request DeleteModelVersionRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *modelVersionsImpl) Get(ctx context.Context, request GetModelVersionRequest) (*ModelVersionInfo, error) {
	var modelVersionInfo ModelVersionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &modelVersionInfo)
	return &modelVersionInfo, err
}

func (a *modelVersionsImpl) GetByAlias(ctx context.Context, request GetByAliasRequest) (*ModelVersionInfo, error) {
	var modelVersionInfo ModelVersionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &modelVersionInfo)
	return &modelVersionInfo, err
}

// List the model versions of the specified registered model, identified by its
// full three-level name (catalog.schema.model).
//
// The returned models are filtered based on the privileges of the calling user.
// For example, the metastore admin is able to list all the model versions. A
// regular user needs to be the owner or have the **EXECUTE** privilege on the
// parent registered model to receive the model versions in the response. For
// the latter case, the caller must also be the owner or have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema.
//
// There is no guarantee of a specific ordering of the elements in the response.
// The elements in the response will not contain any aliases or tags.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *modelVersionsImpl) List(ctx context.Context, request ListModelVersionsRequest) listing.Iterator[ModelVersionInfo] {

	getNextPage := func(ctx context.Context, req ListModelVersionsRequest) (*ListModelVersionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListModelVersionsResponse) []ModelVersionInfo {
		return resp.ModelVersions
	}
	getNextReq := func(resp *ListModelVersionsResponse) *ListModelVersionsRequest {
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

// List the model versions of the specified registered model, identified by its
// full three-level name (catalog.schema.model).
//
// The returned models are filtered based on the privileges of the calling user.
// For example, the metastore admin is able to list all the model versions. A
// regular user needs to be the owner or have the **EXECUTE** privilege on the
// parent registered model to receive the model versions in the response. For
// the latter case, the caller must also be the owner or have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema.
//
// There is no guarantee of a specific ordering of the elements in the response.
// The elements in the response will not contain any aliases or tags.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *modelVersionsImpl) ListAll(ctx context.Context, request ListModelVersionsRequest) ([]ModelVersionInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ModelVersionInfo](ctx, iterator)
}

func (a *modelVersionsImpl) internalList(ctx context.Context, request ListModelVersionsRequest) (*ListModelVersionsResponse, error) {
	var listModelVersionsResponse ListModelVersionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listModelVersionsResponse)
	return &listModelVersionsResponse, err
}

func (a *modelVersionsImpl) Update(ctx context.Context, request UpdateModelVersionRequest) (*ModelVersionInfo, error) {
	var modelVersionInfo ModelVersionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &modelVersionInfo)
	return &modelVersionInfo, err
}

// unexported type that holds implementations of just OnlineTables API methods
type onlineTablesImpl struct {
	client *client.DatabricksClient
}

func (a *onlineTablesImpl) Create(ctx context.Context, request CreateOnlineTableRequest) (*OnlineTable, error) {
	var onlineTable OnlineTable
	path := "/api/2.0/online-tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Table, &onlineTable)
	return &onlineTable, err
}

func (a *onlineTablesImpl) Delete(ctx context.Context, request DeleteOnlineTableRequest) error {
	path := fmt.Sprintf("/api/2.0/online-tables/%v", request.Name)
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

func (a *onlineTablesImpl) Get(ctx context.Context, request GetOnlineTableRequest) (*OnlineTable, error) {
	var onlineTable OnlineTable
	path := fmt.Sprintf("/api/2.0/online-tables/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &onlineTable)
	return &onlineTable, err
}

// unexported type that holds implementations of just Policies API methods
type policiesImpl struct {
	client *client.DatabricksClient
}

func (a *policiesImpl) CreatePolicy(ctx context.Context, request CreatePolicyRequest) (*PolicyInfo, error) {
	var policyInfo PolicyInfo
	path := "/api/2.1/unity-catalog/policies"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.PolicyInfo, &policyInfo)
	return &policyInfo, err
}

func (a *policiesImpl) DeletePolicy(ctx context.Context, request DeletePolicyRequest) (*DeletePolicyResponse, error) {
	var deletePolicyResponse DeletePolicyResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/policies/%v/%v/%v", request.OnSecurableType, request.OnSecurableFullname, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deletePolicyResponse)
	return &deletePolicyResponse, err
}

func (a *policiesImpl) GetPolicy(ctx context.Context, request GetPolicyRequest) (*PolicyInfo, error) {
	var policyInfo PolicyInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/policies/%v/%v/%v", request.OnSecurableType, request.OnSecurableFullname, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &policyInfo)
	return &policyInfo, err
}

// List all policies defined on a securable. Optionally, the list can include
// inherited policies defined on the securable's parent schema or catalog.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *policiesImpl) ListPolicies(ctx context.Context, request ListPoliciesRequest) listing.Iterator[PolicyInfo] {

	getNextPage := func(ctx context.Context, req ListPoliciesRequest) (*ListPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListPolicies(ctx, req)
	}
	getItems := func(resp *ListPoliciesResponse) []PolicyInfo {
		return resp.Policies
	}
	getNextReq := func(resp *ListPoliciesResponse) *ListPoliciesRequest {
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

// List all policies defined on a securable. Optionally, the list can include
// inherited policies defined on the securable's parent schema or catalog.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *policiesImpl) ListPoliciesAll(ctx context.Context, request ListPoliciesRequest) ([]PolicyInfo, error) {
	iterator := a.ListPolicies(ctx, request)
	return listing.ToSlice[PolicyInfo](ctx, iterator)
}

func (a *policiesImpl) internalListPolicies(ctx context.Context, request ListPoliciesRequest) (*ListPoliciesResponse, error) {
	var listPoliciesResponse ListPoliciesResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/policies/%v/%v", request.OnSecurableType, request.OnSecurableFullname)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listPoliciesResponse)
	return &listPoliciesResponse, err
}

func (a *policiesImpl) UpdatePolicy(ctx context.Context, request UpdatePolicyRequest) (*PolicyInfo, error) {
	var policyInfo PolicyInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/policies/%v/%v/%v", request.OnSecurableType, request.OnSecurableFullname, request.Name)
	queryParams := make(map[string]any)

	if request.UpdateMask != "" || slices.Contains(request.ForceSendFields, "UpdateMask") {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.PolicyInfo, &policyInfo)
	return &policyInfo, err
}

// unexported type that holds implementations of just QualityMonitors API methods
type qualityMonitorsImpl struct {
	client *client.DatabricksClient
}

func (a *qualityMonitorsImpl) CancelRefresh(ctx context.Context, request CancelRefreshRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes/%v/cancel", request.TableName, request.RefreshId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, nil)
	return err
}

func (a *qualityMonitorsImpl) Create(ctx context.Context, request CreateMonitor) (*MonitorInfo, error) {
	var monitorInfo MonitorInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &monitorInfo)
	return &monitorInfo, err
}

func (a *qualityMonitorsImpl) Delete(ctx context.Context, request DeleteQualityMonitorRequest) (*DeleteMonitorResponse, error) {
	var deleteMonitorResponse DeleteMonitorResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteMonitorResponse)
	return &deleteMonitorResponse, err
}

func (a *qualityMonitorsImpl) Get(ctx context.Context, request GetQualityMonitorRequest) (*MonitorInfo, error) {
	var monitorInfo MonitorInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &monitorInfo)
	return &monitorInfo, err
}

func (a *qualityMonitorsImpl) GetRefresh(ctx context.Context, request GetRefreshRequest) (*MonitorRefreshInfo, error) {
	var monitorRefreshInfo MonitorRefreshInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes/%v", request.TableName, request.RefreshId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &monitorRefreshInfo)
	return &monitorRefreshInfo, err
}

func (a *qualityMonitorsImpl) ListRefreshes(ctx context.Context, request ListRefreshesRequest) (*MonitorRefreshListResponse, error) {
	var monitorRefreshListResponse MonitorRefreshListResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &monitorRefreshListResponse)
	return &monitorRefreshListResponse, err
}

func (a *qualityMonitorsImpl) RegenerateDashboard(ctx context.Context, request RegenerateDashboardRequest) (*RegenerateDashboardResponse, error) {
	var regenerateDashboardResponse RegenerateDashboardResponse
	path := fmt.Sprintf("/api/2.1/quality-monitoring/tables/%v/monitor/dashboard", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &regenerateDashboardResponse)
	return &regenerateDashboardResponse, err
}

func (a *qualityMonitorsImpl) RunRefresh(ctx context.Context, request RunRefreshRequest) (*MonitorRefreshInfo, error) {
	var monitorRefreshInfo MonitorRefreshInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &monitorRefreshInfo)
	return &monitorRefreshInfo, err
}

func (a *qualityMonitorsImpl) Update(ctx context.Context, request UpdateMonitor) (*MonitorInfo, error) {
	var monitorInfo MonitorInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &monitorInfo)
	return &monitorInfo, err
}

// unexported type that holds implementations of just RegisteredModels API methods
type registeredModelsImpl struct {
	client *client.DatabricksClient
}

func (a *registeredModelsImpl) Create(ctx context.Context, request CreateRegisteredModelRequest) (*RegisteredModelInfo, error) {
	var registeredModelInfo RegisteredModelInfo
	path := "/api/2.1/unity-catalog/models"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &registeredModelInfo)
	return &registeredModelInfo, err
}

func (a *registeredModelsImpl) Delete(ctx context.Context, request DeleteRegisteredModelRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *registeredModelsImpl) DeleteAlias(ctx context.Context, request DeleteAliasRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *registeredModelsImpl) Get(ctx context.Context, request GetRegisteredModelRequest) (*RegisteredModelInfo, error) {
	var registeredModelInfo RegisteredModelInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &registeredModelInfo)
	return &registeredModelInfo, err
}

// List registered models. You can list registered models under a particular
// schema, or list all registered models in the current metastore.
//
// The returned models are filtered based on the privileges of the calling user.
// For example, the metastore admin is able to list all the registered models. A
// regular user needs to be the owner or have the **EXECUTE** privilege on the
// registered model to receive the registered models in the response. For the
// latter case, the caller must also be the owner or have the **USE_CATALOG**
// privilege on the parent catalog and the **USE_SCHEMA** privilege on the
// parent schema.
//
// There is no guarantee of a specific ordering of the elements in the response.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *registeredModelsImpl) List(ctx context.Context, request ListRegisteredModelsRequest) listing.Iterator[RegisteredModelInfo] {

	getNextPage := func(ctx context.Context, req ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListRegisteredModelsResponse) []RegisteredModelInfo {
		return resp.RegisteredModels
	}
	getNextReq := func(resp *ListRegisteredModelsResponse) *ListRegisteredModelsRequest {
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

// List registered models. You can list registered models under a particular
// schema, or list all registered models in the current metastore.
//
// The returned models are filtered based on the privileges of the calling user.
// For example, the metastore admin is able to list all the registered models. A
// regular user needs to be the owner or have the **EXECUTE** privilege on the
// registered model to receive the registered models in the response. For the
// latter case, the caller must also be the owner or have the **USE_CATALOG**
// privilege on the parent catalog and the **USE_SCHEMA** privilege on the
// parent schema.
//
// There is no guarantee of a specific ordering of the elements in the response.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *registeredModelsImpl) ListAll(ctx context.Context, request ListRegisteredModelsRequest) ([]RegisteredModelInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[RegisteredModelInfo](ctx, iterator)
}

func (a *registeredModelsImpl) internalList(ctx context.Context, request ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error) {
	var listRegisteredModelsResponse ListRegisteredModelsResponse
	path := "/api/2.1/unity-catalog/models"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRegisteredModelsResponse)
	return &listRegisteredModelsResponse, err
}

func (a *registeredModelsImpl) SetAlias(ctx context.Context, request SetRegisteredModelAliasRequest) (*RegisteredModelAlias, error) {
	var registeredModelAlias RegisteredModelAlias
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &registeredModelAlias)
	return &registeredModelAlias, err
}

func (a *registeredModelsImpl) Update(ctx context.Context, request UpdateRegisteredModelRequest) (*RegisteredModelInfo, error) {
	var registeredModelInfo RegisteredModelInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &registeredModelInfo)
	return &registeredModelInfo, err
}

// unexported type that holds implementations of just ResourceQuotas API methods
type resourceQuotasImpl struct {
	client *client.DatabricksClient
}

func (a *resourceQuotasImpl) GetQuota(ctx context.Context, request GetQuotaRequest) (*GetQuotaResponse, error) {
	var getQuotaResponse GetQuotaResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/resource-quotas/%v/%v/%v", request.ParentSecurableType, request.ParentFullName, request.QuotaName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getQuotaResponse)
	return &getQuotaResponse, err
}

// ListQuotas returns all quota values under the metastore. There are no SLAs on
// the freshness of the counts returned. This API does not trigger a refresh of
// quota counts.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *resourceQuotasImpl) ListQuotas(ctx context.Context, request ListQuotasRequest) listing.Iterator[QuotaInfo] {

	getNextPage := func(ctx context.Context, req ListQuotasRequest) (*ListQuotasResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListQuotas(ctx, req)
	}
	getItems := func(resp *ListQuotasResponse) []QuotaInfo {
		return resp.Quotas
	}
	getNextReq := func(resp *ListQuotasResponse) *ListQuotasRequest {
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

// ListQuotas returns all quota values under the metastore. There are no SLAs on
// the freshness of the counts returned. This API does not trigger a refresh of
// quota counts.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *resourceQuotasImpl) ListQuotasAll(ctx context.Context, request ListQuotasRequest) ([]QuotaInfo, error) {
	iterator := a.ListQuotas(ctx, request)
	return listing.ToSlice[QuotaInfo](ctx, iterator)
}

func (a *resourceQuotasImpl) internalListQuotas(ctx context.Context, request ListQuotasRequest) (*ListQuotasResponse, error) {
	var listQuotasResponse ListQuotasResponse
	path := "/api/2.1/unity-catalog/resource-quotas/all-resource-quotas"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listQuotasResponse)
	return &listQuotasResponse, err
}

// unexported type that holds implementations of just Rfa API methods
type rfaImpl struct {
	client *client.DatabricksClient
}

func (a *rfaImpl) BatchCreateAccessRequests(ctx context.Context, request BatchCreateAccessRequestsRequest) (*BatchCreateAccessRequestsResponse, error) {
	var batchCreateAccessRequestsResponse BatchCreateAccessRequestsResponse
	path := "/api/3.0/rfa/requests"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &batchCreateAccessRequestsResponse)
	return &batchCreateAccessRequestsResponse, err
}

func (a *rfaImpl) GetAccessRequestDestinations(ctx context.Context, request GetAccessRequestDestinationsRequest) (*AccessRequestDestinations, error) {
	var accessRequestDestinations AccessRequestDestinations
	path := fmt.Sprintf("/api/3.0/rfa/destinations/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &accessRequestDestinations)
	return &accessRequestDestinations, err
}

func (a *rfaImpl) UpdateAccessRequestDestinations(ctx context.Context, request UpdateAccessRequestDestinationsRequest) (*AccessRequestDestinations, error) {
	var accessRequestDestinations AccessRequestDestinations
	path := "/api/3.0/rfa/destinations"
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.AccessRequestDestinations, &accessRequestDestinations)
	return &accessRequestDestinations, err
}

// unexported type that holds implementations of just Schemas API methods
type schemasImpl struct {
	client *client.DatabricksClient
}

func (a *schemasImpl) Create(ctx context.Context, request CreateSchema) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := "/api/2.1/unity-catalog/schemas"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &schemaInfo)
	return &schemaInfo, err
}

func (a *schemasImpl) Delete(ctx context.Context, request DeleteSchemaRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
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

func (a *schemasImpl) Get(ctx context.Context, request GetSchemaRequest) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &schemaInfo)
	return &schemaInfo, err
}

// Gets an array of schemas for a catalog in the metastore. If the caller is the
// metastore admin or the owner of the parent catalog, all schemas for the
// catalog will be retrieved. Otherwise, only schemas owned by the caller (or
// for which the caller has the **USE_SCHEMA** privilege) will be retrieved.
// There is no guarantee of a specific ordering of the elements in the array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *schemasImpl) List(ctx context.Context, request ListSchemasRequest) listing.Iterator[SchemaInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListSchemasRequest) (*ListSchemasResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListSchemasResponse) []SchemaInfo {
		return resp.Schemas
	}
	getNextReq := func(resp *ListSchemasResponse) *ListSchemasRequest {
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

// Gets an array of schemas for a catalog in the metastore. If the caller is the
// metastore admin or the owner of the parent catalog, all schemas for the
// catalog will be retrieved. Otherwise, only schemas owned by the caller (or
// for which the caller has the **USE_SCHEMA** privilege) will be retrieved.
// There is no guarantee of a specific ordering of the elements in the array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *schemasImpl) ListAll(ctx context.Context, request ListSchemasRequest) ([]SchemaInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[SchemaInfo](ctx, iterator)
}

func (a *schemasImpl) internalList(ctx context.Context, request ListSchemasRequest) (*ListSchemasResponse, error) {
	var listSchemasResponse ListSchemasResponse
	path := "/api/2.1/unity-catalog/schemas"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSchemasResponse)
	return &listSchemasResponse, err
}

func (a *schemasImpl) Update(ctx context.Context, request UpdateSchema) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &schemaInfo)
	return &schemaInfo, err
}

// unexported type that holds implementations of just SecretsUc API methods
type secretsUcImpl struct {
	client *client.DatabricksClient
}

func (a *secretsUcImpl) CreateSecret(ctx context.Context, request CreateSecretRequest) (*Secret, error) {
	var secret Secret
	path := "/api/2.1/unity-catalog/secrets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Secret, &secret)
	return &secret, err
}

func (a *secretsUcImpl) DeleteSecret(ctx context.Context, request DeleteSecretRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/secrets/%v", request.FullName)
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

func (a *secretsUcImpl) GetSecret(ctx context.Context, request GetSecretRequest) (*Secret, error) {
	var secret Secret
	path := fmt.Sprintf("/api/2.1/unity-catalog/secrets/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &secret)
	return &secret, err
}

// Lists the secrets in a schema.
//
// With **USE CATALOG** on the catalog and **USE SCHEMA** on the schema, the
// response contains the secrets that you own or hold a privilege on. Metastore
// admins, principals with the **MANAGE** privilege on the catalog, and
// principals with the **MANAGE** privilege on the schema together with **USE
// CATALOG** on the catalog see every secret in the schema.
//
// Secrets are listed within a single schema, so both **catalog_name** and
// **schema_name** must be specified. Results are paginated; pass the
// **next_page_token** value from the response as **page_token** to retrieve the
// next page.
func (a *secretsUcImpl) ListSecrets(ctx context.Context, request ListSecretsRequest) listing.Iterator[Secret] {

	getNextPage := func(ctx context.Context, req ListSecretsRequest) (*ListSecretsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSecrets(ctx, req)
	}
	getItems := func(resp *ListSecretsResponse) []Secret {
		return resp.Secrets
	}
	getNextReq := func(resp *ListSecretsResponse) *ListSecretsRequest {
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

// Lists the secrets in a schema.
//
// With **USE CATALOG** on the catalog and **USE SCHEMA** on the schema, the
// response contains the secrets that you own or hold a privilege on. Metastore
// admins, principals with the **MANAGE** privilege on the catalog, and
// principals with the **MANAGE** privilege on the schema together with **USE
// CATALOG** on the catalog see every secret in the schema.
//
// Secrets are listed within a single schema, so both **catalog_name** and
// **schema_name** must be specified. Results are paginated; pass the
// **next_page_token** value from the response as **page_token** to retrieve the
// next page.
func (a *secretsUcImpl) ListSecretsAll(ctx context.Context, request ListSecretsRequest) ([]Secret, error) {
	iterator := a.ListSecrets(ctx, request)
	return listing.ToSlice[Secret](ctx, iterator)
}

func (a *secretsUcImpl) internalListSecrets(ctx context.Context, request ListSecretsRequest) (*ListSecretsResponse, error) {
	var listSecretsResponse ListSecretsResponse
	path := "/api/2.1/unity-catalog/secrets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSecretsResponse)
	return &listSecretsResponse, err
}

func (a *secretsUcImpl) UpdateSecret(ctx context.Context, request UpdateSecretRequest) (*Secret, error) {
	var secret Secret
	path := fmt.Sprintf("/api/2.1/unity-catalog/secrets/%v", request.FullName)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Secret, &secret)
	return &secret, err
}

// unexported type that holds implementations of just StorageCredentials API methods
type storageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *storageCredentialsImpl) Create(ctx context.Context, request CreateStorageCredential) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := "/api/2.1/unity-catalog/storage-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *storageCredentialsImpl) Delete(ctx context.Context, request DeleteStorageCredentialRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
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

func (a *storageCredentialsImpl) Get(ctx context.Context, request GetStorageCredentialRequest) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

// Gets an array of storage credentials (as __StorageCredentialInfo__ objects).
// The array is limited to only those storage credentials the caller has
// permission to access. If the caller is a metastore admin, retrieval of
// credentials is unrestricted. There is no guarantee of a specific ordering of
// the elements in the array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *storageCredentialsImpl) List(ctx context.Context, request ListStorageCredentialsRequest) listing.Iterator[StorageCredentialInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListStorageCredentialsRequest) (*ListStorageCredentialsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListStorageCredentialsResponse) []StorageCredentialInfo {
		return resp.StorageCredentials
	}
	getNextReq := func(resp *ListStorageCredentialsResponse) *ListStorageCredentialsRequest {
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

// Gets an array of storage credentials (as __StorageCredentialInfo__ objects).
// The array is limited to only those storage credentials the caller has
// permission to access. If the caller is a metastore admin, retrieval of
// credentials is unrestricted. There is no guarantee of a specific ordering of
// the elements in the array.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *storageCredentialsImpl) ListAll(ctx context.Context, request ListStorageCredentialsRequest) ([]StorageCredentialInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[StorageCredentialInfo](ctx, iterator)
}

func (a *storageCredentialsImpl) internalList(ctx context.Context, request ListStorageCredentialsRequest) (*ListStorageCredentialsResponse, error) {
	var listStorageCredentialsResponse ListStorageCredentialsResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listStorageCredentialsResponse)
	return &listStorageCredentialsResponse, err
}

func (a *storageCredentialsImpl) Update(ctx context.Context, request UpdateStorageCredential) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *storageCredentialsImpl) Validate(ctx context.Context, request ValidateStorageCredential) (*ValidateStorageCredentialResponse, error) {
	var validateStorageCredentialResponse ValidateStorageCredentialResponse
	path := "/api/2.1/unity-catalog/validate-storage-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &validateStorageCredentialResponse)
	return &validateStorageCredentialResponse, err
}

// unexported type that holds implementations of just SystemSchemas API methods
type systemSchemasImpl struct {
	client *client.DatabricksClient
}

func (a *systemSchemasImpl) Disable(ctx context.Context, request DisableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas/%v", request.MetastoreId, request.SchemaName)
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

func (a *systemSchemasImpl) Enable(ctx context.Context, request EnableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas/%v", request.MetastoreId, request.SchemaName)
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

// Gets an array of system schemas for a metastore. The caller must be an
// account admin or a metastore admin.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *systemSchemasImpl) List(ctx context.Context, request ListSystemSchemasRequest) listing.Iterator[SystemSchemaInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListSystemSchemasRequest) (*ListSystemSchemasResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListSystemSchemasResponse) []SystemSchemaInfo {
		return resp.Schemas
	}
	getNextReq := func(resp *ListSystemSchemasResponse) *ListSystemSchemasRequest {
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

// Gets an array of system schemas for a metastore. The caller must be an
// account admin or a metastore admin.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *systemSchemasImpl) ListAll(ctx context.Context, request ListSystemSchemasRequest) ([]SystemSchemaInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[SystemSchemaInfo](ctx, iterator)
}

func (a *systemSchemasImpl) internalList(ctx context.Context, request ListSystemSchemasRequest) (*ListSystemSchemasResponse, error) {
	var listSystemSchemasResponse ListSystemSchemasResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas", request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSystemSchemasResponse)
	return &listSystemSchemasResponse, err
}

// unexported type that holds implementations of just TableConstraints API methods
type tableConstraintsImpl struct {
	client *client.DatabricksClient
}

func (a *tableConstraintsImpl) Create(ctx context.Context, request CreateTableConstraint) (*TableConstraint, error) {
	var tableConstraint TableConstraint
	path := "/api/2.1/unity-catalog/constraints"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &tableConstraint)
	return &tableConstraint, err
}

func (a *tableConstraintsImpl) Delete(ctx context.Context, request DeleteTableConstraintRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/constraints/%v", request.FullName)
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

// unexported type that holds implementations of just Tables API methods
type tablesImpl struct {
	client *client.DatabricksClient
}

func (a *tablesImpl) Create(ctx context.Context, request CreateTableRequest) (*TableInfo, error) {
	var tableInfo TableInfo
	path := "/api/2.1/unity-catalog/tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &tableInfo)
	return &tableInfo, err
}

func (a *tablesImpl) Delete(ctx context.Context, request DeleteTableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
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

func (a *tablesImpl) Exists(ctx context.Context, request ExistsRequest) (*TableExistsResponse, error) {
	var tableExistsResponse TableExistsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/exists", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &tableExistsResponse)
	return &tableExistsResponse, err
}

func (a *tablesImpl) Get(ctx context.Context, request GetTableRequest) (*TableInfo, error) {
	var tableInfo TableInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &tableInfo)
	return &tableInfo, err
}

// Gets an array of all tables for the current metastore under the parent
// catalog and schema. The caller must be a metastore admin or an owner of (or
// have the **SELECT** privilege on) the table. For the latter case, the caller
// must also be the owner or have the **USE_CATALOG** privilege on the parent
// catalog and the **USE_SCHEMA** privilege on the parent schema. There is no
// guarantee of a specific ordering of the elements in the array.
//
// NOTE: **view_dependencies** and **table_constraints** are not returned by
// ListTables queries.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *tablesImpl) List(ctx context.Context, request ListTablesRequest) listing.Iterator[TableInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListTablesRequest) (*ListTablesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListTablesResponse) []TableInfo {
		return resp.Tables
	}
	getNextReq := func(resp *ListTablesResponse) *ListTablesRequest {
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

// Gets an array of all tables for the current metastore under the parent
// catalog and schema. The caller must be a metastore admin or an owner of (or
// have the **SELECT** privilege on) the table. For the latter case, the caller
// must also be the owner or have the **USE_CATALOG** privilege on the parent
// catalog and the **USE_SCHEMA** privilege on the parent schema. There is no
// guarantee of a specific ordering of the elements in the array.
//
// NOTE: **view_dependencies** and **table_constraints** are not returned by
// ListTables queries.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *tablesImpl) ListAll(ctx context.Context, request ListTablesRequest) ([]TableInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[TableInfo](ctx, iterator)
}

func (a *tablesImpl) internalList(ctx context.Context, request ListTablesRequest) (*ListTablesResponse, error) {
	var listTablesResponse ListTablesResponse
	path := "/api/2.1/unity-catalog/tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTablesResponse)
	return &listTablesResponse, err
}

// Gets an array of summaries for tables for a schema and catalog within the
// metastore. The table summaries returned are either:
//
// * summaries for tables (within the current metastore and parent catalog and
// schema), when the user is a metastore admin, or: * summaries for tables and
// schemas (within the current metastore and parent catalog) for which the user
// has ownership or the **SELECT** privilege on the table and ownership or
// **USE_SCHEMA** privilege on the schema, provided that the user also has
// ownership or the **USE_CATALOG** privilege on the parent catalog.
//
// There is no guarantee of a specific ordering of the elements in the array.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *tablesImpl) ListSummaries(ctx context.Context, request ListSummariesRequest) listing.Iterator[TableSummary] {

	getNextPage := func(ctx context.Context, req ListSummariesRequest) (*ListTableSummariesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSummaries(ctx, req)
	}
	getItems := func(resp *ListTableSummariesResponse) []TableSummary {
		return resp.Tables
	}
	getNextReq := func(resp *ListTableSummariesResponse) *ListSummariesRequest {
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

// Gets an array of summaries for tables for a schema and catalog within the
// metastore. The table summaries returned are either:
//
// * summaries for tables (within the current metastore and parent catalog and
// schema), when the user is a metastore admin, or: * summaries for tables and
// schemas (within the current metastore and parent catalog) for which the user
// has ownership or the **SELECT** privilege on the table and ownership or
// **USE_SCHEMA** privilege on the schema, provided that the user also has
// ownership or the **USE_CATALOG** privilege on the parent catalog.
//
// There is no guarantee of a specific ordering of the elements in the array.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *tablesImpl) ListSummariesAll(ctx context.Context, request ListSummariesRequest) ([]TableSummary, error) {
	iterator := a.ListSummaries(ctx, request)
	return listing.ToSlice[TableSummary](ctx, iterator)
}

func (a *tablesImpl) internalListSummaries(ctx context.Context, request ListSummariesRequest) (*ListTableSummariesResponse, error) {
	var listTableSummariesResponse ListTableSummariesResponse
	path := "/api/2.1/unity-catalog/table-summaries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTableSummariesResponse)
	return &listTableSummariesResponse, err
}

func (a *tablesImpl) Update(ctx context.Context, request UpdateTableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
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

// unexported type that holds implementations of just TemporaryPathCredentials API methods
type temporaryPathCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *temporaryPathCredentialsImpl) GenerateTemporaryPathCredentials(ctx context.Context, request GenerateTemporaryPathCredentialRequest) (*GenerateTemporaryPathCredentialResponse, error) {
	var generateTemporaryPathCredentialResponse GenerateTemporaryPathCredentialResponse
	path := "/api/2.0/unity-catalog/temporary-path-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &generateTemporaryPathCredentialResponse)
	return &generateTemporaryPathCredentialResponse, err
}

// unexported type that holds implementations of just TemporaryTableCredentials API methods
type temporaryTableCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *temporaryTableCredentialsImpl) GenerateTemporaryTableCredentials(ctx context.Context, request GenerateTemporaryTableCredentialRequest) (*GenerateTemporaryTableCredentialResponse, error) {
	var generateTemporaryTableCredentialResponse GenerateTemporaryTableCredentialResponse
	path := "/api/2.0/unity-catalog/temporary-table-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &generateTemporaryTableCredentialResponse)
	return &generateTemporaryTableCredentialResponse, err
}

// unexported type that holds implementations of just TemporaryVolumeCredentials API methods
type temporaryVolumeCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *temporaryVolumeCredentialsImpl) GenerateTemporaryVolumeCredentials(ctx context.Context, request GenerateTemporaryVolumeCredentialRequest) (*GenerateTemporaryVolumeCredentialResponse, error) {
	var generateTemporaryVolumeCredentialResponse GenerateTemporaryVolumeCredentialResponse
	path := "/api/2.0/unity-catalog/temporary-volume-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &generateTemporaryVolumeCredentialResponse)
	return &generateTemporaryVolumeCredentialResponse, err
}

// unexported type that holds implementations of just Volumes API methods
type volumesImpl struct {
	client *client.DatabricksClient
}

func (a *volumesImpl) Create(ctx context.Context, request CreateVolumeRequestContent) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := "/api/2.1/unity-catalog/volumes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &volumeInfo)
	return &volumeInfo, err
}

func (a *volumesImpl) Delete(ctx context.Context, request DeleteVolumeRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

// Gets an array of volumes for the current metastore under the parent catalog
// and schema.
//
// The returned volumes are filtered based on the privileges of the calling
// user. For example, the metastore admin is able to list all the volumes. A
// regular user needs to be the owner or have the **READ VOLUME** privilege on
// the volume to receive the volumes in the response. For the latter case, the
// caller must also be the owner or have the **USE_CATALOG** privilege on the
// parent catalog and the **USE_SCHEMA** privilege on the parent schema.
//
// There is no guarantee of a specific ordering of the elements in the array.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *volumesImpl) List(ctx context.Context, request ListVolumesRequest) listing.Iterator[VolumeInfo] {

	getNextPage := func(ctx context.Context, req ListVolumesRequest) (*ListVolumesResponseContent, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListVolumesResponseContent) []VolumeInfo {
		return resp.Volumes
	}
	getNextReq := func(resp *ListVolumesResponseContent) *ListVolumesRequest {
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

// Gets an array of volumes for the current metastore under the parent catalog
// and schema.
//
// The returned volumes are filtered based on the privileges of the calling
// user. For example, the metastore admin is able to list all the volumes. A
// regular user needs to be the owner or have the **READ VOLUME** privilege on
// the volume to receive the volumes in the response. For the latter case, the
// caller must also be the owner or have the **USE_CATALOG** privilege on the
// parent catalog and the **USE_SCHEMA** privilege on the parent schema.
//
// There is no guarantee of a specific ordering of the elements in the array.
//
// PAGINATION BEHAVIOR: The API is by default paginated, a page may contain zero
// results while still providing a next_page_token. Clients must continue
// reading pages until next_page_token is absent, which is the only indication
// that the end of results has been reached.
func (a *volumesImpl) ListAll(ctx context.Context, request ListVolumesRequest) ([]VolumeInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[VolumeInfo](ctx, iterator)
}

func (a *volumesImpl) internalList(ctx context.Context, request ListVolumesRequest) (*ListVolumesResponseContent, error) {
	var listVolumesResponseContent ListVolumesResponseContent
	path := "/api/2.1/unity-catalog/volumes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listVolumesResponseContent)
	return &listVolumesResponseContent, err
}

func (a *volumesImpl) Read(ctx context.Context, request ReadVolumeRequest) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &volumeInfo)
	return &volumeInfo, err
}

func (a *volumesImpl) Update(ctx context.Context, request UpdateVolumeRequestContent) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &volumeInfo)
	return &volumeInfo, err
}

// unexported type that holds implementations of just WorkspaceBindings API methods
type workspaceBindingsImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceBindingsImpl) Get(ctx context.Context, request GetWorkspaceBindingRequest) (*GetCatalogWorkspaceBindingsResponse, error) {
	var getCatalogWorkspaceBindingsResponse GetCatalogWorkspaceBindingsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspace-bindings/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getCatalogWorkspaceBindingsResponse)
	return &getCatalogWorkspaceBindingsResponse, err
}

// Gets workspace bindings of the securable. The caller must be a metastore
// admin or an owner of the securable.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *workspaceBindingsImpl) GetBindings(ctx context.Context, request GetBindingsRequest) listing.Iterator[WorkspaceBinding] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req GetBindingsRequest) (*GetWorkspaceBindingsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalGetBindings(ctx, req)
	}
	getItems := func(resp *GetWorkspaceBindingsResponse) []WorkspaceBinding {
		return resp.Bindings
	}
	getNextReq := func(resp *GetWorkspaceBindingsResponse) *GetBindingsRequest {
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

// Gets workspace bindings of the securable. The caller must be a metastore
// admin or an owner of the securable.
//
// NOTE: we recommend using max_results=0 to use the paginated version of this
// API. Unpaginated calls will be deprecated soon.
//
// PAGINATION BEHAVIOR: When using pagination (max_results >= 0), a page may
// contain zero results while still providing a next_page_token. Clients must
// continue reading pages until next_page_token is absent, which is the only
// indication that the end of results has been reached.
func (a *workspaceBindingsImpl) GetBindingsAll(ctx context.Context, request GetBindingsRequest) ([]WorkspaceBinding, error) {
	iterator := a.GetBindings(ctx, request)
	return listing.ToSlice[WorkspaceBinding](ctx, iterator)
}

func (a *workspaceBindingsImpl) internalGetBindings(ctx context.Context, request GetBindingsRequest) (*GetWorkspaceBindingsResponse, error) {
	var getWorkspaceBindingsResponse GetWorkspaceBindingsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/bindings/%v/%v", request.SecurableType, request.SecurableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getWorkspaceBindingsResponse)
	return &getWorkspaceBindingsResponse, err
}

func (a *workspaceBindingsImpl) Update(ctx context.Context, request UpdateWorkspaceBindings) (*UpdateCatalogWorkspaceBindingsResponse, error) {
	var updateCatalogWorkspaceBindingsResponse UpdateCatalogWorkspaceBindingsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspace-bindings/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateCatalogWorkspaceBindingsResponse)
	return &updateCatalogWorkspaceBindingsResponse, err
}

func (a *workspaceBindingsImpl) UpdateBindings(ctx context.Context, request UpdateWorkspaceBindingsParameters) (*UpdateWorkspaceBindingsResponse, error) {
	var updateWorkspaceBindingsResponse UpdateWorkspaceBindingsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/bindings/%v/%v", request.SecurableType, request.SecurableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateWorkspaceBindingsResponse)
	return &updateWorkspaceBindingsResponse, err
}
