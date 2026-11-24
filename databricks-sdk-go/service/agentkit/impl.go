// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package agentkit

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

// unexported type that holds implementations of just AgentKit API methods
type agentKitImpl struct {
	client *client.DatabricksClient
}

func (a *agentKitImpl) AppendSessionItems(ctx context.Context, request AppendSessionItemsRequest) (*AppendSessionItemsResponse, error) {
	var appendSessionItemsResponse AppendSessionItemsResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/items:append", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &appendSessionItemsResponse)
	return &appendSessionItemsResponse, err
}

func (a *agentKitImpl) ClearSessionItems(ctx context.Context, request ClearSessionItemsRequest) (*ClearSessionItemsResponse, error) {
	var clearSessionItemsResponse ClearSessionItemsResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/items:clear", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &clearSessionItemsResponse)
	return &clearSessionItemsResponse, err
}

func (a *agentKitImpl) CreateMemory(ctx context.Context, request CreateManagedMemoryEntryRequest) (*ManagedMemoryEntry, error) {
	var managedMemoryEntry ManagedMemoryEntry
	path := fmt.Sprintf("/api/2.0/agents/%v/entries", request.Parent)
	queryParams := make(map[string]any)

	if request.ManagedMemoryEntryId != "" || slices.Contains(request.ForceSendFields, "ManagedMemoryEntryId") {
		queryParams["managed_memory_entry_id"] = request.ManagedMemoryEntryId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ManagedMemoryEntry, &managedMemoryEntry)
	return &managedMemoryEntry, err
}

func (a *agentKitImpl) CreateMemoryStore(ctx context.Context, request CreateManagedMemoryStoreRequest) (*ManagedMemoryStore, error) {
	var managedMemoryStore ManagedMemoryStore
	path := "/api/2.0/agents/memory-stores"
	queryParams := make(map[string]any)

	if request.ManagedMemoryStoreId != "" {
		queryParams["managed_memory_store_id"] = request.ManagedMemoryStoreId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ManagedMemoryStore, &managedMemoryStore)
	return &managedMemoryStore, err
}

func (a *agentKitImpl) CreateSession(ctx context.Context, request CreateSessionRequest) (*Session, error) {
	var session Session
	path := fmt.Sprintf("/api/2.0/agents/%v/sessions", request.Parent)
	queryParams := make(map[string]any)

	if request.SessionId != "" || slices.Contains(request.ForceSendFields, "SessionId") {
		queryParams["session_id"] = request.SessionId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Session, &session)
	return &session, err
}

func (a *agentKitImpl) CreateSessionStore(ctx context.Context, request CreateSessionStoreRequest) (*SessionStore, error) {
	var sessionStore SessionStore
	path := "/api/2.0/agents/session-stores"
	queryParams := make(map[string]any)

	if request.SessionStoreId != "" {
		queryParams["session_store_id"] = request.SessionStoreId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.SessionStore, &sessionStore)
	return &sessionStore, err
}

func (a *agentKitImpl) DeleteMemory(ctx context.Context, request DeleteManagedMemoryEntryRequest) error {
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
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

func (a *agentKitImpl) DeleteMemoryStore(ctx context.Context, request DeleteManagedMemoryStoreRequest) error {
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
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

func (a *agentKitImpl) DeleteSession(ctx context.Context, request DeleteSessionRequest) error {
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
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

func (a *agentKitImpl) DeleteSessionStore(ctx context.Context, request DeleteSessionStoreRequest) error {
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
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

func (a *agentKitImpl) ExtractMemories(ctx context.Context, request ExtractMemoriesRequest) (*ExtractMemoriesResponse, error) {
	var extractMemoriesResponse ExtractMemoriesResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/sessions/%v/extractions", request.SessionStore, request.SessionId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &extractMemoriesResponse)
	return &extractMemoriesResponse, err
}

func (a *agentKitImpl) ForkSession(ctx context.Context, request ForkSessionRequest) (*ForkSessionResponse, error) {
	var forkSessionResponse ForkSessionResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/sessions:fork", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &forkSessionResponse)
	return &forkSessionResponse, err
}

func (a *agentKitImpl) GetMemory(ctx context.Context, request GetManagedMemoryEntryRequest) (*ManagedMemoryEntry, error) {
	var managedMemoryEntry ManagedMemoryEntry
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &managedMemoryEntry)
	return &managedMemoryEntry, err
}

func (a *agentKitImpl) GetMemoryStore(ctx context.Context, request GetManagedMemoryStoreRequest) (*ManagedMemoryStore, error) {
	var managedMemoryStore ManagedMemoryStore
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &managedMemoryStore)
	return &managedMemoryStore, err
}

func (a *agentKitImpl) GetMemoryStorePermission(ctx context.Context, request GetMemoryStorePermissionRequest) (*MemoryStorePermission, error) {
	var memoryStorePermission MemoryStorePermission
	path := fmt.Sprintf("/api/2.0/agents/%v/permissions/self", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &memoryStorePermission)
	return &memoryStorePermission, err
}

func (a *agentKitImpl) GetSession(ctx context.Context, request GetSessionRequest) (*Session, error) {
	var session Session
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &session)
	return &session, err
}

func (a *agentKitImpl) GetSessionStore(ctx context.Context, request GetSessionStoreRequest) (*SessionStore, error) {
	var sessionStore SessionStore
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &sessionStore)
	return &sessionStore, err
}

func (a *agentKitImpl) GetSessionStorePermission(ctx context.Context, request GetSessionStorePermissionRequest) (*SessionStorePermission, error) {
	var sessionStorePermission SessionStorePermission
	path := fmt.Sprintf("/api/2.0/agents/%v/permissions/self", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &sessionStorePermission)
	return &sessionStorePermission, err
}

func (a *agentKitImpl) GrantMemoryStorePermission(ctx context.Context, request GrantMemoryStorePermissionRequest) (*GrantMemoryStorePermissionResponse, error) {
	var grantMemoryStorePermissionResponse GrantMemoryStorePermissionResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/permissions:grant", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &grantMemoryStorePermissionResponse)
	return &grantMemoryStorePermissionResponse, err
}

func (a *agentKitImpl) GrantSessionStorePermission(ctx context.Context, request GrantSessionStorePermissionRequest) (*GrantSessionStorePermissionResponse, error) {
	var grantSessionStorePermissionResponse GrantSessionStorePermissionResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/permissions:grant", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &grantSessionStorePermissionResponse)
	return &grantSessionStorePermissionResponse, err
}

// Lists memory entries for one actor. An exact `path` filters entries across
// sessions, ignoring session metadata. Otherwise, `path_prefix` and the
// deprecated `session_id` restrict the actor partition. `view` controls whether
// each returned entry includes memory content.
func (a *agentKitImpl) ListMemories(ctx context.Context, request ListManagedMemoryEntriesRequest) listing.Iterator[ManagedMemoryEntry] {

	getNextPage := func(ctx context.Context, req ListManagedMemoryEntriesRequest) (*ListManagedMemoryEntriesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListMemories(ctx, req)
	}
	getItems := func(resp *ListManagedMemoryEntriesResponse) []ManagedMemoryEntry {
		return resp.ManagedMemoryEntries
	}
	getNextReq := func(resp *ListManagedMemoryEntriesResponse) *ListManagedMemoryEntriesRequest {
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

// Lists memory entries for one actor. An exact `path` filters entries across
// sessions, ignoring session metadata. Otherwise, `path_prefix` and the
// deprecated `session_id` restrict the actor partition. `view` controls whether
// each returned entry includes memory content.
func (a *agentKitImpl) ListMemoriesAll(ctx context.Context, request ListManagedMemoryEntriesRequest) ([]ManagedMemoryEntry, error) {
	iterator := a.ListMemories(ctx, request)
	return listing.ToSlice[ManagedMemoryEntry](ctx, iterator)
}

func (a *agentKitImpl) internalListMemories(ctx context.Context, request ListManagedMemoryEntriesRequest) (*ListManagedMemoryEntriesResponse, error) {
	var listManagedMemoryEntriesResponse ListManagedMemoryEntriesResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/entries", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listManagedMemoryEntriesResponse)
	return &listManagedMemoryEntriesResponse, err
}

// Lists explicit memory store grants. Only the owner or a workspace admin may
// list permissions.
func (a *agentKitImpl) ListMemoryStorePermissions(ctx context.Context, request ListMemoryStorePermissionsRequest) listing.Iterator[MemoryStorePermission] {

	getNextPage := func(ctx context.Context, req ListMemoryStorePermissionsRequest) (*ListMemoryStorePermissionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListMemoryStorePermissions(ctx, req)
	}
	getItems := func(resp *ListMemoryStorePermissionsResponse) []MemoryStorePermission {
		return resp.MemoryStorePermissions
	}
	getNextReq := func(resp *ListMemoryStorePermissionsResponse) *ListMemoryStorePermissionsRequest {
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

// Lists explicit memory store grants. Only the owner or a workspace admin may
// list permissions.
func (a *agentKitImpl) ListMemoryStorePermissionsAll(ctx context.Context, request ListMemoryStorePermissionsRequest) ([]MemoryStorePermission, error) {
	iterator := a.ListMemoryStorePermissions(ctx, request)
	return listing.ToSlice[MemoryStorePermission](ctx, iterator)
}

func (a *agentKitImpl) internalListMemoryStorePermissions(ctx context.Context, request ListMemoryStorePermissionsRequest) (*ListMemoryStorePermissionsResponse, error) {
	var listMemoryStorePermissionsResponse ListMemoryStorePermissionsResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/permissions", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listMemoryStorePermissionsResponse)
	return &listMemoryStorePermissionsResponse, err
}

// Lists memory stores in the caller's workspace.
func (a *agentKitImpl) ListMemoryStores(ctx context.Context, request ListManagedMemoryStoresRequest) listing.Iterator[ManagedMemoryStore] {

	getNextPage := func(ctx context.Context, req ListManagedMemoryStoresRequest) (*ListManagedMemoryStoresResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListMemoryStores(ctx, req)
	}
	getItems := func(resp *ListManagedMemoryStoresResponse) []ManagedMemoryStore {
		return resp.ManagedMemoryStores
	}
	getNextReq := func(resp *ListManagedMemoryStoresResponse) *ListManagedMemoryStoresRequest {
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

// Lists memory stores in the caller's workspace.
func (a *agentKitImpl) ListMemoryStoresAll(ctx context.Context, request ListManagedMemoryStoresRequest) ([]ManagedMemoryStore, error) {
	iterator := a.ListMemoryStores(ctx, request)
	return listing.ToSlice[ManagedMemoryStore](ctx, iterator)
}

func (a *agentKitImpl) internalListMemoryStores(ctx context.Context, request ListManagedMemoryStoresRequest) (*ListManagedMemoryStoresResponse, error) {
	var listManagedMemoryStoresResponse ListManagedMemoryStoresResponse
	path := "/api/2.0/agents/memory-stores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listManagedMemoryStoresResponse)
	return &listManagedMemoryStoresResponse, err
}

// Lists items in a session.
func (a *agentKitImpl) ListSessionItems(ctx context.Context, request ListSessionItemsRequest) listing.Iterator[SessionItem] {

	getNextPage := func(ctx context.Context, req ListSessionItemsRequest) (*ListSessionItemsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSessionItems(ctx, req)
	}
	getItems := func(resp *ListSessionItemsResponse) []SessionItem {
		return resp.SessionItems
	}
	getNextReq := func(resp *ListSessionItemsResponse) *ListSessionItemsRequest {
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

// Lists items in a session.
func (a *agentKitImpl) ListSessionItemsAll(ctx context.Context, request ListSessionItemsRequest) ([]SessionItem, error) {
	iterator := a.ListSessionItems(ctx, request)
	return listing.ToSlice[SessionItem](ctx, iterator)
}

func (a *agentKitImpl) internalListSessionItems(ctx context.Context, request ListSessionItemsRequest) (*ListSessionItemsResponse, error) {
	var listSessionItemsResponse ListSessionItemsResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/items", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSessionItemsResponse)
	return &listSessionItemsResponse, err
}

// Lists explicit session store grants. Only the owner or a workspace admin may
// list permissions.
func (a *agentKitImpl) ListSessionStorePermissions(ctx context.Context, request ListSessionStorePermissionsRequest) listing.Iterator[SessionStorePermission] {

	getNextPage := func(ctx context.Context, req ListSessionStorePermissionsRequest) (*ListSessionStorePermissionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSessionStorePermissions(ctx, req)
	}
	getItems := func(resp *ListSessionStorePermissionsResponse) []SessionStorePermission {
		return resp.SessionStorePermissions
	}
	getNextReq := func(resp *ListSessionStorePermissionsResponse) *ListSessionStorePermissionsRequest {
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

// Lists explicit session store grants. Only the owner or a workspace admin may
// list permissions.
func (a *agentKitImpl) ListSessionStorePermissionsAll(ctx context.Context, request ListSessionStorePermissionsRequest) ([]SessionStorePermission, error) {
	iterator := a.ListSessionStorePermissions(ctx, request)
	return listing.ToSlice[SessionStorePermission](ctx, iterator)
}

func (a *agentKitImpl) internalListSessionStorePermissions(ctx context.Context, request ListSessionStorePermissionsRequest) (*ListSessionStorePermissionsResponse, error) {
	var listSessionStorePermissionsResponse ListSessionStorePermissionsResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/permissions", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSessionStorePermissionsResponse)
	return &listSessionStorePermissionsResponse, err
}

// Lists session stores.
func (a *agentKitImpl) ListSessionStores(ctx context.Context, request ListSessionStoresRequest) listing.Iterator[SessionStore] {

	getNextPage := func(ctx context.Context, req ListSessionStoresRequest) (*ListSessionStoresResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSessionStores(ctx, req)
	}
	getItems := func(resp *ListSessionStoresResponse) []SessionStore {
		return resp.SessionStores
	}
	getNextReq := func(resp *ListSessionStoresResponse) *ListSessionStoresRequest {
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

// Lists session stores.
func (a *agentKitImpl) ListSessionStoresAll(ctx context.Context, request ListSessionStoresRequest) ([]SessionStore, error) {
	iterator := a.ListSessionStores(ctx, request)
	return listing.ToSlice[SessionStore](ctx, iterator)
}

func (a *agentKitImpl) internalListSessionStores(ctx context.Context, request ListSessionStoresRequest) (*ListSessionStoresResponse, error) {
	var listSessionStoresResponse ListSessionStoresResponse
	path := "/api/2.0/agents/session-stores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSessionStoresResponse)
	return &listSessionStoresResponse, err
}

// Lists sessions within a session store.
func (a *agentKitImpl) ListSessions(ctx context.Context, request ListSessionsRequest) listing.Iterator[Session] {

	getNextPage := func(ctx context.Context, req ListSessionsRequest) (*ListSessionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSessions(ctx, req)
	}
	getItems := func(resp *ListSessionsResponse) []Session {
		return resp.Sessions
	}
	getNextReq := func(resp *ListSessionsResponse) *ListSessionsRequest {
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

// Lists sessions within a session store.
func (a *agentKitImpl) ListSessionsAll(ctx context.Context, request ListSessionsRequest) ([]Session, error) {
	iterator := a.ListSessions(ctx, request)
	return listing.ToSlice[Session](ctx, iterator)
}

func (a *agentKitImpl) internalListSessions(ctx context.Context, request ListSessionsRequest) (*ListSessionsResponse, error) {
	var listSessionsResponse ListSessionsResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/sessions", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSessionsResponse)
	return &listSessionsResponse, err
}

func (a *agentKitImpl) PopSessionItem(ctx context.Context, request PopSessionItemRequest) (*PopSessionItemResponse, error) {
	var popSessionItemResponse PopSessionItemResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/items:pop", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &popSessionItemResponse)
	return &popSessionItemResponse, err
}

func (a *agentKitImpl) RevokeMemoryStorePermission(ctx context.Context, request RevokeMemoryStorePermissionRequest) (*RevokeMemoryStorePermissionResponse, error) {
	var revokeMemoryStorePermissionResponse RevokeMemoryStorePermissionResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/permissions:revoke", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &revokeMemoryStorePermissionResponse)
	return &revokeMemoryStorePermissionResponse, err
}

func (a *agentKitImpl) RevokeSessionStorePermission(ctx context.Context, request RevokeSessionStorePermissionRequest) (*RevokeSessionStorePermissionResponse, error) {
	var revokeSessionStorePermissionResponse RevokeSessionStorePermissionResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/permissions:revoke", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &revokeSessionStorePermissionResponse)
	return &revokeSessionStorePermissionResponse, err
}

// Searches memory entries by text query for one actor. Returns matching entries
// and scores ranked by relevance; `view` controls whether each returned entry
// includes content.
func (a *agentKitImpl) SearchMemories(ctx context.Context, request SearchManagedMemoryEntriesRequest) listing.Iterator[ManagedMemoryEntrySearchResult] {

	getNextPage := func(ctx context.Context, req SearchManagedMemoryEntriesRequest) (*SearchManagedMemoryEntriesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalSearchMemories(ctx, req)
	}
	getItems := func(resp *SearchManagedMemoryEntriesResponse) []ManagedMemoryEntrySearchResult {
		return resp.Results
	}
	getNextReq := func(resp *SearchManagedMemoryEntriesResponse) *SearchManagedMemoryEntriesRequest {
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

// Searches memory entries by text query for one actor. Returns matching entries
// and scores ranked by relevance; `view` controls whether each returned entry
// includes content.
func (a *agentKitImpl) SearchMemoriesAll(ctx context.Context, request SearchManagedMemoryEntriesRequest) ([]ManagedMemoryEntrySearchResult, error) {
	iterator := a.SearchMemories(ctx, request)
	return listing.ToSlice[ManagedMemoryEntrySearchResult](ctx, iterator)
}

func (a *agentKitImpl) internalSearchMemories(ctx context.Context, request SearchManagedMemoryEntriesRequest) (*SearchManagedMemoryEntriesResponse, error) {
	var searchManagedMemoryEntriesResponse SearchManagedMemoryEntriesResponse
	path := fmt.Sprintf("/api/2.0/agents/%v/entries:search", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &searchManagedMemoryEntriesResponse)
	return &searchManagedMemoryEntriesResponse, err
}

func (a *agentKitImpl) UpdateMemory(ctx context.Context, request UpdateManagedMemoryEntryRequest) (*ManagedMemoryEntry, error) {
	var managedMemoryEntry ManagedMemoryEntry
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ManagedMemoryEntry, &managedMemoryEntry)
	return &managedMemoryEntry, err
}

func (a *agentKitImpl) UpdateMemoryStore(ctx context.Context, request UpdateManagedMemoryStoreRequest) (*ManagedMemoryStore, error) {
	var managedMemoryStore ManagedMemoryStore
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.ManagedMemoryStore, &managedMemoryStore)
	return &managedMemoryStore, err
}

func (a *agentKitImpl) UpdateSession(ctx context.Context, request UpdateSessionRequest) (*Session, error) {
	var session Session
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Session, &session)
	return &session, err
}

func (a *agentKitImpl) UpdateSessionStore(ctx context.Context, request UpdateSessionStoreRequest) (*SessionStore, error) {
	var sessionStore SessionStore
	path := fmt.Sprintf("/api/2.0/agents/%v", request.Name)
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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.SessionStore, &sessionStore)
	return &sessionStore, err
}
