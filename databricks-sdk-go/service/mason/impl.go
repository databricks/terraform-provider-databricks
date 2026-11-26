// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mason

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

// unexported type that holds implementations of just Mason API methods
type masonImpl struct {
	client *client.DatabricksClient
}

func (a *masonImpl) AppendSessionItems(ctx context.Context, request AppendSessionItemsRequest) (*AppendSessionItemsResponse, error) {
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

func (a *masonImpl) ClearSessionItems(ctx context.Context, request ClearSessionItemsRequest) (*ClearSessionItemsResponse, error) {
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

func (a *masonImpl) CreateMemory(ctx context.Context, request CreateManagedMemoryEntryRequest) (*ManagedMemoryEntry, error) {
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

func (a *masonImpl) CreateMemoryStore(ctx context.Context, request CreateManagedMemoryStoreRequest) (*ManagedMemoryStore, error) {
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

func (a *masonImpl) CreateSession(ctx context.Context, request CreateSessionRequest) (*Session, error) {
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

func (a *masonImpl) CreateSessionStore(ctx context.Context, request CreateSessionStoreRequest) (*SessionStore, error) {
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

func (a *masonImpl) DeleteMemory(ctx context.Context, request DeleteManagedMemoryEntryRequest) error {
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

func (a *masonImpl) DeleteMemoryStore(ctx context.Context, request DeleteManagedMemoryStoreRequest) error {
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

func (a *masonImpl) DeleteSession(ctx context.Context, request DeleteSessionRequest) error {
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

func (a *masonImpl) DeleteSessionStore(ctx context.Context, request DeleteSessionStoreRequest) error {
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

func (a *masonImpl) ExtractMemories(ctx context.Context, request ExtractMemoriesRequest) (*ExtractMemoriesResponse, error) {
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

func (a *masonImpl) ForkSession(ctx context.Context, request ForkSessionRequest) (*ForkSessionResponse, error) {
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

func (a *masonImpl) GetMemory(ctx context.Context, request GetManagedMemoryEntryRequest) (*ManagedMemoryEntry, error) {
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

func (a *masonImpl) GetMemoryStore(ctx context.Context, request GetManagedMemoryStoreRequest) (*ManagedMemoryStore, error) {
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

func (a *masonImpl) GetSession(ctx context.Context, request GetSessionRequest) (*Session, error) {
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

func (a *masonImpl) GetSessionStore(ctx context.Context, request GetSessionStoreRequest) (*SessionStore, error) {
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

// Lists managed memory entries for one actor. An exact `path` filters entries
// across sessions, ignoring session metadata. Otherwise, `session_id` and
// `path_prefix` restrict the actor partition. `read_mask` selects fields in
// each returned entry.
func (a *masonImpl) ListMemories(ctx context.Context, request ListManagedMemoryEntriesRequest) listing.Iterator[ManagedMemoryEntry] {

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

// Lists managed memory entries for one actor. An exact `path` filters entries
// across sessions, ignoring session metadata. Otherwise, `session_id` and
// `path_prefix` restrict the actor partition. `read_mask` selects fields in
// each returned entry.
func (a *masonImpl) ListMemoriesAll(ctx context.Context, request ListManagedMemoryEntriesRequest) ([]ManagedMemoryEntry, error) {
	iterator := a.ListMemories(ctx, request)
	return listing.ToSlice[ManagedMemoryEntry](ctx, iterator)
}

func (a *masonImpl) internalListMemories(ctx context.Context, request ListManagedMemoryEntriesRequest) (*ListManagedMemoryEntriesResponse, error) {
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

// Lists managed memory stores in the caller's workspace.
func (a *masonImpl) ListMemoryStores(ctx context.Context, request ListManagedMemoryStoresRequest) listing.Iterator[ManagedMemoryStore] {

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

// Lists managed memory stores in the caller's workspace.
func (a *masonImpl) ListMemoryStoresAll(ctx context.Context, request ListManagedMemoryStoresRequest) ([]ManagedMemoryStore, error) {
	iterator := a.ListMemoryStores(ctx, request)
	return listing.ToSlice[ManagedMemoryStore](ctx, iterator)
}

func (a *masonImpl) internalListMemoryStores(ctx context.Context, request ListManagedMemoryStoresRequest) (*ListManagedMemoryStoresResponse, error) {
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
func (a *masonImpl) ListSessionItems(ctx context.Context, request ListSessionItemsRequest) listing.Iterator[SessionItem] {

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
func (a *masonImpl) ListSessionItemsAll(ctx context.Context, request ListSessionItemsRequest) ([]SessionItem, error) {
	iterator := a.ListSessionItems(ctx, request)
	return listing.ToSlice[SessionItem](ctx, iterator)
}

func (a *masonImpl) internalListSessionItems(ctx context.Context, request ListSessionItemsRequest) (*ListSessionItemsResponse, error) {
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

// Lists session stores.
func (a *masonImpl) ListSessionStores(ctx context.Context, request ListSessionStoresRequest) listing.Iterator[SessionStore] {

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
func (a *masonImpl) ListSessionStoresAll(ctx context.Context, request ListSessionStoresRequest) ([]SessionStore, error) {
	iterator := a.ListSessionStores(ctx, request)
	return listing.ToSlice[SessionStore](ctx, iterator)
}

func (a *masonImpl) internalListSessionStores(ctx context.Context, request ListSessionStoresRequest) (*ListSessionStoresResponse, error) {
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
func (a *masonImpl) ListSessions(ctx context.Context, request ListSessionsRequest) listing.Iterator[Session] {

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
func (a *masonImpl) ListSessionsAll(ctx context.Context, request ListSessionsRequest) ([]Session, error) {
	iterator := a.ListSessions(ctx, request)
	return listing.ToSlice[Session](ctx, iterator)
}

func (a *masonImpl) internalListSessions(ctx context.Context, request ListSessionsRequest) (*ListSessionsResponse, error) {
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

func (a *masonImpl) PopSessionItem(ctx context.Context, request PopSessionItemRequest) (*PopSessionItemResponse, error) {
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

// Searches managed memory entries by text query for one actor. Returns matching
// entries and scores ranked by relevance; `read_mask` selects fields in each
// returned entry.
func (a *masonImpl) SearchMemories(ctx context.Context, request SearchManagedMemoryEntriesRequest) listing.Iterator[ManagedMemoryEntrySearchResult] {

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

// Searches managed memory entries by text query for one actor. Returns matching
// entries and scores ranked by relevance; `read_mask` selects fields in each
// returned entry.
func (a *masonImpl) SearchMemoriesAll(ctx context.Context, request SearchManagedMemoryEntriesRequest) ([]ManagedMemoryEntrySearchResult, error) {
	iterator := a.SearchMemories(ctx, request)
	return listing.ToSlice[ManagedMemoryEntrySearchResult](ctx, iterator)
}

func (a *masonImpl) internalSearchMemories(ctx context.Context, request SearchManagedMemoryEntriesRequest) (*SearchManagedMemoryEntriesResponse, error) {
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

func (a *masonImpl) UpdateMemory(ctx context.Context, request UpdateManagedMemoryEntryRequest) (*ManagedMemoryEntry, error) {
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

func (a *masonImpl) UpdateMemoryStore(ctx context.Context, request UpdateManagedMemoryStoreRequest) (*ManagedMemoryStore, error) {
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

func (a *masonImpl) UpdateSession(ctx context.Context, request UpdateSessionRequest) (*Session, error) {
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

func (a *masonImpl) UpdateSessionStore(ctx context.Context, request UpdateSessionStoreRequest) (*SessionStore, error) {
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
