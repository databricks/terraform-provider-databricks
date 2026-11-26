// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package agentkit

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// Request to append items to a session.
type AppendSessionItemsRequest struct {
	// Items to append atomically in request order. Concurrent append requests
	// are serialized into one committed order without exposing a numeric
	// sequence in the public contract.
	Items []SessionItem `json:"items"`
	// Resource name of the containing session, in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Parent string `json:"-" url:"-"`
}

func (s *AppendSessionItemsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response containing appended items.
type AppendSessionItemsResponse struct {
	// Persisted session items with service-assigned fields.
	SessionItems []SessionItem `json:"session_items,omitempty"`
}

func (s *AppendSessionItemsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to clear all items from a session.
type ClearSessionItemsRequest struct {
	// Resource name of the containing session, in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Parent string `json:"-" url:"-"`
}

func (s *ClearSessionItemsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response from clearing items from a session.
type ClearSessionItemsResponse struct {
}

func (s *ClearSessionItemsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateManagedMemoryEntryRequest struct {
	// The memory entry to create.
	ManagedMemoryEntry ManagedMemoryEntry `json:"managed_memory_entry"`
	// Caller-selected memory entry ID. It must contain 1-255 letters, digits,
	// hyphens, or underscores, and it becomes the last segment of the entry's
	// `name`. The service generates a UUID when it is omitted.
	ManagedMemoryEntryId string `json:"-" url:"managed_memory_entry_id,omitempty"`
	// Memory store that will contain the entry, in the form
	// `memory-stores/{managed_memory_store_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateManagedMemoryEntryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateManagedMemoryEntryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateManagedMemoryStoreRequest struct {
	// The memory store to create.
	ManagedMemoryStore ManagedMemoryStore `json:"managed_memory_store"`
	// Caller-provided, workspace-unique memory store ID. It must be 3-56
	// characters, begin with a lowercase letter, contain only lowercase
	// letters, digits, and hyphens, and end with a letter or digit. It becomes
	// the last segment of the store's `name` and cannot be changed after
	// creation.
	ManagedMemoryStoreId string `json:"-" url:"managed_memory_store_id"`
}

func (s *CreateManagedMemoryStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateSessionRequest struct {
	// Resource name of the containing session store, in the form
	// `session-stores/{session_store_id}`.
	Parent string `json:"-" url:"-"`
	// The session to create. `actor_id` is required. A session with
	// `parent_session_id` is a child and must use its parent's `actor_id`.
	// Independent forks are created only through `ForkSession`.
	Session Session `json:"session"`
	// Optional caller-selected session ID. The service generates a UUID when
	// this field is omitted. The ID must be unique; a collision returns
	// `ALREADY_EXISTS`.
	SessionId string `json:"-" url:"session_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateSessionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateSessionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateSessionStoreRequest struct {
	// The session store to create.
	SessionStore SessionStore `json:"session_store"`
	// Caller-provided, workspace-unique session store ID. It must be 3-55
	// characters, begin with a lowercase letter, and contain only lowercase
	// letters, digits, and hyphens.
	SessionStoreId string `json:"-" url:"session_store_id"`
}

func (s *CreateSessionStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteManagedMemoryEntryRequest struct {
	// Resource name in the form
	// `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteManagedMemoryEntryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteManagedMemoryStoreRequest struct {
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteManagedMemoryStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteSessionRequest struct {
	// Resource name in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteSessionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteSessionStoreRequest struct {
	// Resource name in the form `session-stores/{session_store_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteSessionStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to synchronously extract memories from a single session.
type ExtractMemoriesRequest struct {
	// When true, extract and return the entries without writing them to the
	// memory store. Defaults to false, which persists the extracted entries and
	// returns them.
	DryRun bool `json:"dry_run,omitempty"`
	// Instructions steering what is extracted from the session.
	Instructions string `json:"instructions,omitempty"`
	// Memory store the extracted entries are written to, in the form
	// `memory-stores/{managed_memory_store_id}`.
	MemoryStore string `json:"memory_store"`
	// Identifier of the session whose transcript is distilled into memories.
	SessionId string `json:"-" url:"-"`
	// Session store containing the session, in the form
	// `session-stores/{session_store_id}`.
	SessionStore string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExtractMemoriesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExtractMemoriesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Result of a single-session memory extraction.
type ExtractMemoriesResponse struct {
	// The memory entries written by this extraction.
	Entries []ManagedMemoryEntry `json:"entries,omitempty"`
	// Correlation identifier for this extraction, for logging and tracing. Not
	// a fetchable resource.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExtractMemoriesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExtractMemoriesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to fork a session.
type ForkSessionRequest struct {
	// Opaque caller-provided identifier for the application actor associated
	// with the forked session.
	ActorId string `json:"actor_id"`
	// Optional metadata for the fork.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Resource name of the containing session store, in the form
	// `session-stores/{session_store_id}`.
	Parent string `json:"-" url:"-"`
	// Optional unique ID for the forked session. A collision returns
	// `ALREADY_EXISTS`.
	SessionId string `json:"session_id,omitempty"`
	// ID of the session to copy.
	SourceSessionId string `json:"source_session_id"`
	// Optional last item ID to copy through, inclusively. When omitted, the
	// fork atomically copies all items committed before the fork operation
	// begins.
	UpToItemId string `json:"up_to_item_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ForkSessionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForkSessionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response from forking a session.
type ForkSessionResponse struct {
	// The newly-created independent top-level session.
	Session *Session `json:"session,omitempty"`
}

func (s *ForkSessionResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetManagedMemoryEntryRequest struct {
	// Resource name in the form
	// `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`.
	Name string `json:"-" url:"-"`
	// Controls whether the response includes memory content. Defaults to
	// `FULL`.
	View ManagedMemoryEntryView `json:"-" url:"view,omitempty"`
}

func (s *GetManagedMemoryEntryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetManagedMemoryStoreRequest struct {
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetManagedMemoryStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetMemoryStorePermissionRequest struct {
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetMemoryStorePermissionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSessionRequest struct {
	// Resource name in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetSessionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSessionStorePermissionRequest struct {
	// Resource name in the form `session-stores/{session_store_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetSessionStorePermissionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSessionStoreRequest struct {
	// Resource name in the form `session-stores/{session_store_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetSessionStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to grant a READ or WRITE permission on a memory store to a principal.
// Only the store owner or a workspace admin may grant permissions.
type GrantMemoryStorePermissionRequest struct {
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name string `json:"-" url:"-"`
	// Permission level to grant. `WRITE` includes `READ`. Grants are additive:
	// granting `READ` to a principal that already holds `WRITE` does not lower
	// its access. To downgrade a principal, revoke its grants and grant again.
	Permission PermissionLevel `json:"permission"`
	// Principal to grant the permission to.
	Principal StorePrincipal `json:"principal"`
}

func (s *GrantMemoryStorePermissionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response from granting a permission on a memory store.
type GrantMemoryStorePermissionResponse struct {
}

func (s *GrantMemoryStorePermissionResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to grant a READ or WRITE permission on a session store to a
// principal. Only the store owner or a workspace admin may grant permissions.
type GrantSessionStorePermissionRequest struct {
	// Resource name in the form `session-stores/{session_store_id}`.
	Name string `json:"-" url:"-"`
	// Permission level to grant.
	Permission PermissionLevel `json:"permission"`
	// Principal to grant the permission to.
	Principal StorePrincipal `json:"principal"`
}

func (s *GrantSessionStorePermissionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response from granting a permission on a session store.
type GrantSessionStorePermissionResponse struct {
}

func (s *GrantSessionStorePermissionResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListManagedMemoryEntriesRequest struct {
	// Customer-provided identifier for the actor whose entries are listed.
	ActorId string `json:"-" url:"actor_id"`
	// Maximum number of entries to return. The service may return fewer entries
	// than requested. Defaults to 10; must be between 1 and 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token, taken from the `next_page_token` of a previous
	// response. The other request parameters must match the request that
	// returned the token.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Memory store whose entries are listed, in the form
	// `memory-stores/{managed_memory_store_id}`.
	Parent string `json:"-" url:"-"`
	// Optional path prefix used to restrict entries within the actor partition.
	PathPrefix string `json:"-" url:"path_prefix,omitempty"`
	// Deprecated: the service will ignore this filter. Until then, when set,
	// only entries with this exact `session_id` are returned, and
	// omitted-session (cross-session) entries are not included. Ignored when
	// path is set.
	SessionId string `json:"-" url:"session_id,omitempty"`
	// Controls whether each returned entry includes memory content. Defaults to
	// `FULL`.
	View ManagedMemoryEntryView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListManagedMemoryEntriesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListManagedMemoryEntriesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing memory entries.
type ListManagedMemoryEntriesResponse struct {
	// Memory entries matching the request and selected response view, ordered
	// by creation time, oldest first.
	ManagedMemoryEntries []ManagedMemoryEntry `json:"managed_memory_entries,omitempty"`
	// Opaque pagination token. This field is omitted when there are no more
	// results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListManagedMemoryEntriesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListManagedMemoryEntriesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListManagedMemoryStoresRequest struct {
	// Maximum number of stores to return. The service may return fewer stores
	// than requested. Defaults to 10; must be between 1 and 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token, taken from the `next_page_token` of a previous
	// response.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListManagedMemoryStoresRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListManagedMemoryStoresRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing memory stores in the caller's workspace.
type ListManagedMemoryStoresResponse struct {
	// Memory stores in the caller's workspace.
	ManagedMemoryStores []ManagedMemoryStore `json:"managed_memory_stores,omitempty"`
	// Opaque pagination token. This field is omitted when there are no more
	// results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListManagedMemoryStoresResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListManagedMemoryStoresResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListMemoryStorePermissionsRequest struct {
	// Maximum number of permissions to return. Defaults to 100; maximum is
	// 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Continuation token from a previous response.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Resource name of the store, in the form
	// `memory-stores/{managed_memory_store_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMemoryStorePermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMemoryStorePermissionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of explicit memory store permissions.
type ListMemoryStorePermissionsResponse struct {
	// Explicit READ or WRITE grants. Implicit owner/admin access and inherited
	// grants are excluded.
	MemoryStorePermissions []MemoryStorePermission `json:"memory_store_permissions,omitempty"`
	// Token for the next page, omitted when all permissions have been returned.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMemoryStorePermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMemoryStorePermissionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSessionItemsRequest struct {
	// Sort order. Supported values are `create_time asc` and `create_time
	// desc`. The default is `create_time desc`, which returns the most recently
	// appended items first. Equal timestamps are resolved by committed append
	// order in the requested direction.
	OrderBy string `json:"-" url:"order_by,omitempty"`
	// Maximum number of items to return. Defaults to 10; must be between 1 and
	// 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token returned by a previous list request.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Resource name of the containing session, in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSessionItemsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSessionItemsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing a page of session items.
type ListSessionItemsResponse struct {
	// Token to retrieve the next page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Session items in the requested page.
	SessionItems []SessionItem `json:"session_items,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSessionItemsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSessionItemsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSessionStorePermissionsRequest struct {
	// Maximum number of permissions to return. Defaults to 100; maximum is
	// 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Continuation token from a previous response.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Resource name of the store, in the form
	// `session-stores/{session_store_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSessionStorePermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSessionStorePermissionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A page of explicit session store permissions.
type ListSessionStorePermissionsResponse struct {
	// Token for the next page, omitted when all permissions have been returned.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Explicit READ or WRITE grants. Implicit owner/admin access and inherited
	// grants are excluded.
	SessionStorePermissions []SessionStorePermission `json:"session_store_permissions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSessionStorePermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSessionStorePermissionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSessionStoresRequest struct {
	// Maximum number of session stores to return. Defaults to 10; must be
	// between 1 and 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token returned by a previous list request.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSessionStoresRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSessionStoresRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing a page of session stores.
type ListSessionStoresResponse struct {
	// Token to retrieve the next page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Session stores in the requested page.
	SessionStores []SessionStore `json:"session_stores,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSessionStoresResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSessionStoresResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSessionsRequest struct {
	// Filter expression. Supported fields include `actor_id` and `metadata`;
	// for example, `actor_id = "support-customer-123"`.
	Filter string `json:"-" url:"filter,omitempty"`
	// Sort order. Defaults to `last_activity_time desc`. Page-token
	// continuation is exactly-once when ordering by `create_time` (immutable);
	// ordering by `last_activity_time` is best-effort, because that value
	// changes as a session gains activity, so a session updated between page
	// requests may be repeated or skipped. To enumerate every session exactly
	// once, order by `create_time`.
	OrderBy string `json:"-" url:"order_by,omitempty"`
	// Maximum number of sessions to return. Defaults to 10; must be between 1
	// and 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token returned by a previous list request.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Resource name of the containing session store, in the form
	// `session-stores/{session_store_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSessionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSessionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing a page of sessions.
type ListSessionsResponse struct {
	// Token to retrieve the next page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Sessions in the requested page.
	Sessions []Session `json:"sessions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSessionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSessionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A workspace-scoped entry in a memory store. Each entry belongs to one actor
// and, optionally, one session, and its `path` is unique within that actor and
// session.
type ManagedMemoryEntry struct {
	// Customer-provided identifier for the actor whose memory this entry
	// represents.
	ActorId string `json:"actor_id"`
	// Optional free-form memory content.
	Content string `json:"content,omitempty"`
	// Time when the entry was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Human-readable description of the memory entry.
	Description string `json:"description,omitempty"`
	// Resource name in the form
	// `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`.
	// Required for updates; assigned by the service on creation.
	Name string `json:"name,omitempty"`
	// Absolute, case-sensitive path identifying the entry within its actor and,
	// until the deprecated `session_id` is ignored, its optional session. Paths
	// must begin with `/` and must not contain empty, `.` or `..` segments.
	Path string `json:"path"`
	// Deprecated: memory entries will be identified by `(actor_id, path)`, and
	// the service will ignore this field on writes and omit it from responses.
	// Until then, an omitted value means the entry applies across the actor's
	// sessions.
	SessionId string `json:"session_id,omitempty"`
	// Which writer created this entry. Caller sets this on Create; immutable
	// after creation. The server may store an omitted or UNSPECIFIED value as
	// AGENT.
	SourceType ManagedMemoryEntrySourceType `json:"source_type,omitempty"`
	// Time when the entry was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ManagedMemoryEntry) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ManagedMemoryEntry) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// One relevance-ranked memory search result.
type ManagedMemoryEntrySearchResult struct {
	// Memory entry matching the query.
	ManagedMemoryEntry *ManagedMemoryEntry `json:"managed_memory_entry,omitempty"`
	// Relevance score for the result. Higher scores are more relevant.
	Score float64 `json:"score,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ManagedMemoryEntrySearchResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ManagedMemoryEntrySearchResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Identifies the source that created a memory entry.
type ManagedMemoryEntrySourceType string

const ManagedMemoryEntrySourceTypeManagedMemoryEntrySourceTypeAgent ManagedMemoryEntrySourceType = `MANAGED_MEMORY_ENTRY_SOURCE_TYPE_AGENT`

const ManagedMemoryEntrySourceTypeManagedMemoryEntrySourceTypeDreamer ManagedMemoryEntrySourceType = `MANAGED_MEMORY_ENTRY_SOURCE_TYPE_DREAMER`

// String representation for [fmt.Print]
func (f *ManagedMemoryEntrySourceType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ManagedMemoryEntrySourceType) Set(v string) error {
	switch v {
	case `MANAGED_MEMORY_ENTRY_SOURCE_TYPE_AGENT`, `MANAGED_MEMORY_ENTRY_SOURCE_TYPE_DREAMER`:
		*f = ManagedMemoryEntrySourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANAGED_MEMORY_ENTRY_SOURCE_TYPE_AGENT", "MANAGED_MEMORY_ENTRY_SOURCE_TYPE_DREAMER"`, v)
	}
}

// Values returns all possible values for ManagedMemoryEntrySourceType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ManagedMemoryEntrySourceType) Values() []ManagedMemoryEntrySourceType {
	return []ManagedMemoryEntrySourceType{
		ManagedMemoryEntrySourceTypeManagedMemoryEntrySourceTypeAgent,
		ManagedMemoryEntrySourceTypeManagedMemoryEntrySourceTypeDreamer,
	}
}

// Type always returns ManagedMemoryEntrySourceType to satisfy [pflag.Value] interface
func (f *ManagedMemoryEntrySourceType) Type() string {
	return "ManagedMemoryEntrySourceType"
}

// Controls how much memory entry data is returned.
type ManagedMemoryEntryView string

const ManagedMemoryEntryViewManagedMemoryEntryViewBasic ManagedMemoryEntryView = `MANAGED_MEMORY_ENTRY_VIEW_BASIC`

const ManagedMemoryEntryViewManagedMemoryEntryViewFull ManagedMemoryEntryView = `MANAGED_MEMORY_ENTRY_VIEW_FULL`

// String representation for [fmt.Print]
func (f *ManagedMemoryEntryView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ManagedMemoryEntryView) Set(v string) error {
	switch v {
	case `MANAGED_MEMORY_ENTRY_VIEW_BASIC`, `MANAGED_MEMORY_ENTRY_VIEW_FULL`:
		*f = ManagedMemoryEntryView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANAGED_MEMORY_ENTRY_VIEW_BASIC", "MANAGED_MEMORY_ENTRY_VIEW_FULL"`, v)
	}
}

// Values returns all possible values for ManagedMemoryEntryView.
//
// There is no guarantee on the order of the values in the slice.
func (f *ManagedMemoryEntryView) Values() []ManagedMemoryEntryView {
	return []ManagedMemoryEntryView{
		ManagedMemoryEntryViewManagedMemoryEntryViewBasic,
		ManagedMemoryEntryViewManagedMemoryEntryViewFull,
	}
}

// Type always returns ManagedMemoryEntryView to satisfy [pflag.Value] interface
func (f *ManagedMemoryEntryView) Type() string {
	return "ManagedMemoryEntryView"
}

// A workspace-scoped memory store backed by service-managed storage. A store is
// the container for an agent's memory entries and the boundary at which access
// is granted; creating one provisions its Lakebase storage automatically. See
// https://docs.databricks.com/aws/en/agents/agent-memory/managed-memory for an
// overview.
type ManagedMemoryStore struct {
	// Time when the store was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Workspace-local user ID of the authenticated principal that created the
	// store. This is immutable server-set attribution and does not grant
	// access; authorization is evaluated from the authenticated request
	// context.
	CreatorUserId string `json:"creator_user_id,omitempty"`
	// Human-readable description of the memory store.
	Description string `json:"description,omitempty"`
	// Deprecated compatibility alias for the caller-provided memory store ID.
	// Canonical clients provide the ID through
	// `CreateMemoryStoreRequest.managed_memory_store_id` and use `name` as the
	// resource identifier.
	DisplayName string `json:"display_name,omitempty"`
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name string `json:"name,omitempty"`
	// Deprecated alias for `creator_user_id`. This identifies the original
	// creator, not a transferable owner. Use `creator_user_id` instead.
	OwnerUserId string `json:"owner_user_id,omitempty"`
	// Service-managed storage backing this memory store.
	StorageBackend *StorageBackend `json:"storage_backend,omitempty"`
	// Time when the store was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Workspace that owns the memory store.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ManagedMemoryStore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ManagedMemoryStore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A principal's permission on a memory store.
type MemoryStorePermission struct {
	// The permission level. WRITE includes READ; PERMISSION_LEVEL_UNSPECIFIED
	// means neither is available.
	Permission PermissionLevel `json:"permission,omitempty"`
	// The principal whose permission is reported.
	Principal *StorePrincipal `json:"principal,omitempty"`
}

func (s *MemoryStorePermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Permission level granted on a store. WRITE is a superset of READ.
type PermissionLevel string

const PermissionLevelRead PermissionLevel = `READ`

const PermissionLevelWrite PermissionLevel = `WRITE`

// String representation for [fmt.Print]
func (f *PermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PermissionLevel) Set(v string) error {
	switch v {
	case `READ`, `WRITE`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "READ", "WRITE"`, v)
	}
}

// Values returns all possible values for PermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PermissionLevel) Values() []PermissionLevel {
	return []PermissionLevel{
		PermissionLevelRead,
		PermissionLevelWrite,
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

// Request to pop an item from a session.
type PopSessionItemRequest struct {
	// Resource name of the containing session, in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Parent string `json:"-" url:"-"`
}

func (s *PopSessionItemRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response containing the popped item.
type PopSessionItemResponse struct {
	// Removed item, if any.
	Item *SessionItem `json:"item,omitempty"`
}

func (s *PopSessionItemResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The kind of Databricks principal a store permission is granted to.
type PrincipalType string

const PrincipalTypeServicePrincipal PrincipalType = `SERVICE_PRINCIPAL`

const PrincipalTypeUser PrincipalType = `USER`

// String representation for [fmt.Print]
func (f *PrincipalType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrincipalType) Set(v string) error {
	switch v {
	case `SERVICE_PRINCIPAL`, `USER`:
		*f = PrincipalType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SERVICE_PRINCIPAL", "USER"`, v)
	}
}

// Values returns all possible values for PrincipalType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrincipalType) Values() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeServicePrincipal,
		PrincipalTypeUser,
	}
}

// Type always returns PrincipalType to satisfy [pflag.Value] interface
func (f *PrincipalType) Type() string {
	return "PrincipalType"
}

// Request to revoke all explicit access to a memory store. Only the owner or a
// workspace admin may revoke permissions.
type RevokeMemoryStorePermissionRequest struct {
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name string `json:"-" url:"-"`
	// Principal whose READ and WRITE grants will be removed.
	Principal StorePrincipal `json:"principal"`
}

func (s *RevokeMemoryStorePermissionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response from revoking a memory store permission.
type RevokeMemoryStorePermissionResponse struct {
}

func (s *RevokeMemoryStorePermissionResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to revoke all explicit access to a session store. Only the owner or a
// workspace admin may revoke permissions.
type RevokeSessionStorePermissionRequest struct {
	// Resource name in the form `session-stores/{session_store_id}`.
	Name string `json:"-" url:"-"`
	// Principal whose READ and WRITE grants will be removed.
	Principal StorePrincipal `json:"principal"`
}

func (s *RevokeSessionStorePermissionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response from revoking a session store permission.
type RevokeSessionStorePermissionResponse struct {
}

func (s *RevokeSessionStorePermissionResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to search memory entries by text query for one actor. Search returns
// a relevance-ranked top-N result set and does not currently paginate.
type SearchManagedMemoryEntriesRequest struct {
	// Customer-provided identifier for the actor whose entries are searched.
	ActorId string `json:"actor_id"`
	// Deprecated alias for `page_size`. When both fields are set, their values
	// must match.
	Limit int `json:"limit,omitempty"`
	// Maximum number of relevance-ranked entries to return. Defaults to 10 and
	// must be between 1 and 100.
	PageSize int `json:"page_size,omitempty"`
	// Reserved for pagination compatibility. The server currently ignores this
	// field because Search returns a ranked top-N result set.
	PageToken string `json:"page_token,omitempty"`
	// Memory store whose entries are searched, in the form
	// `memory-stores/{managed_memory_store_id}`.
	Parent string `json:"-" url:"-"`
	// Optional absolute, case-sensitive path prefix used to restrict searched
	// entries within the actor partition. The prefix must begin with `/` and
	// must not contain empty, `.` or `..` segments.
	PathPrefix string `json:"path_prefix,omitempty"`
	// Free-form search query. Entries are ranked by how well their `content`
	// and `description` match it.
	Query string `json:"query"`
	// Deprecated: the service will ignore this filter. Until then, when set,
	// only entries with this exact `session_id` are searched, and
	// omitted-session (cross-session) entries are not included.
	SessionId string `json:"session_id,omitempty"`
	// Controls whether each returned entry includes memory content. Search
	// scores are always returned. Defaults to `FULL`.
	View ManagedMemoryEntryView `json:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchManagedMemoryEntriesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchManagedMemoryEntriesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing memory entries ranked by relevance.
type SearchManagedMemoryEntriesResponse struct {
	// Deprecated compatibility alias for clients migrating to `results`. This
	// contains the same entries in the same order, but omits their relevance
	// scores.
	ManagedMemoryEntries []ManagedMemoryEntry `json:"managed_memory_entries,omitempty"`
	// Opaque pagination token. Search currently returns an unpaginated ranked
	// top-N result set, so the server does not populate this field.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Canonical matching entries and relevance scores, ordered most relevant
	// first.
	Results []ManagedMemoryEntrySearchResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchManagedMemoryEntriesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchManagedMemoryEntriesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A durable logical interaction stored within a Session Store.
type Session struct {
	// Opaque caller-provided identifier for the application actor associated
	// with the session.
	//
	// This is application data and has no Databricks authentication or
	// authorization semantics. Use the same value as the Managed Memory Entry
	// `actor_id` when storing memories associated with this actor. Every
	// session must set it. A child session must use the same value as its
	// parent.
	ActorId string `json:"actor_id"`
	// Time when the session was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Time when the session's item history was last mutated.
	LastActivityTime *time.Time `json:"last_activity_time,omitempty"`
	// Mutable caller-defined string labels.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Resource name in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Name string `json:"name,omitempty"`
	// Immediate parent session ID. Set only at creation for child sessions,
	// immutable thereafter, and restricted to the same store.
	ParentSessionId string `json:"parent_session_id,omitempty"`
	// Top-level session ID in the spawn tree. This equals `session_id` for a
	// root or fork and is inherited transitively by child sessions.
	RootSessionId string `json:"root_session_id,omitempty"`
	// Unique session ID. The service generates a UUID unless the caller
	// supplies `CreateSessionRequest.session_id`.
	SessionId string `json:"session_id,omitempty"`
	// Time when session resource fields last changed.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Session) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Session) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A transcript entry in a session's history.
type SessionItem struct {
	// Server-assigned time when the append commits. Values are nondecreasing
	// within a session. Item listing orders by this timestamp; equal timestamps
	// are resolved by committed append order.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Complete SDK-native, JSON-compatible item. The service stores and returns
	// this value without interpreting provider-specific fields such as `type`,
	// `role`, or `content`.
	Data json.RawMessage `json:"data"`
	// Stable service-generated item ID.
	ItemId string `json:"item_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SessionItem) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SessionItem) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A workspace-scoped session store.
type SessionStore struct {
	// Time when the store was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Workspace-local user ID of the authenticated principal that created the
	// store. This is immutable server-set attribution and does not grant
	// access; authorization is evaluated from the authenticated request
	// context.
	CreatorUserId string `json:"creator_user_id,omitempty"`
	// Human-readable description of the session store.
	Description string `json:"description,omitempty"`
	// Mutable caller-defined string labels.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Resource name in the form `session-stores/{session_store_id}`.
	Name string `json:"name,omitempty"`
	// Time when the store was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SessionStore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SessionStore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A principal's permission on a session store.
type SessionStorePermission struct {
	// The permission level. WRITE includes READ; PERMISSION_LEVEL_UNSPECIFIED
	// means neither is available.
	Permission PermissionLevel `json:"permission,omitempty"`
	// The principal whose permission is reported.
	Principal *StorePrincipal `json:"principal,omitempty"`
}

func (s *SessionStorePermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Service-managed storage backing a memory store.
type StorageBackend struct {
	// Backend-specific identifier. For Lakebase, this is the resource name of
	// the database that holds the store's entries, in the form
	// `projects/{project_id}/branches/{branch_id}/databases/{database_id}`.
	BackendId string `json:"backend_id,omitempty"`
	// Type of the storage backend.
	BackendType StorageBackendType `json:"backend_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StorageBackend) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StorageBackend) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Type of service-managed storage backing a memory store.
type StorageBackendType string

const StorageBackendTypeStorageBackendTypeLakebase StorageBackendType = `STORAGE_BACKEND_TYPE_LAKEBASE`

// String representation for [fmt.Print]
func (f *StorageBackendType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *StorageBackendType) Set(v string) error {
	switch v {
	case `STORAGE_BACKEND_TYPE_LAKEBASE`:
		*f = StorageBackendType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "STORAGE_BACKEND_TYPE_LAKEBASE"`, v)
	}
}

// Values returns all possible values for StorageBackendType.
//
// There is no guarantee on the order of the values in the slice.
func (f *StorageBackendType) Values() []StorageBackendType {
	return []StorageBackendType{
		StorageBackendTypeStorageBackendTypeLakebase,
	}
}

// Type always returns StorageBackendType to satisfy [pflag.Value] interface
func (f *StorageBackendType) Type() string {
	return "StorageBackendType"
}

// A principal a store permission is granted to. `name` is the principal's
// Lakebase identifier: a user's email or a service principal's application
// (client) id.
type StorePrincipal struct {
	// Identifier of the principal: the email address of a user, or the
	// application (client) ID of a service principal. Groups are not supported.
	Name string `json:"name"`
	// The kind of principal, which determines how `name` is interpreted. A
	// principal without a type is rejected.
	PrincipalType PrincipalType `json:"principal_type,omitempty"`
}

func (s *StorePrincipal) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateManagedMemoryEntryRequest struct {
	// The memory entry to update.
	ManagedMemoryEntry ManagedMemoryEntry `json:"managed_memory_entry"`
	// Resource name in the form
	// `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`.
	// Required for updates; assigned by the service on creation.
	Name string `json:"-" url:"-"`
	// Fields to update. Only `content` and `description` may be updated.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateManagedMemoryEntryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateManagedMemoryStoreRequest struct {
	// The memory store to update. `name` is taken from the URL.
	ManagedMemoryStore ManagedMemoryStore `json:"managed_memory_store"`
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name string `json:"-" url:"-"`
	// Only `description` and `metadata` may be updated. Selecting `description`
	// while leaving it unset clears it. Selecting `metadata` replaces the
	// entire map.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateManagedMemoryStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateSessionRequest struct {
	// Resource name in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Name string `json:"-" url:"-"`
	// Session to update.
	Session Session `json:"session"`
	// Fields to update. Only `metadata` is mutable; any other path returns
	// `INVALID_PARAMETER_VALUE`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateSessionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateSessionStoreRequest struct {
	// Resource name in the form `session-stores/{session_store_id}`.
	Name string `json:"-" url:"-"`
	// Session store to update.
	SessionStore SessionStore `json:"session_store"`
	// Fields to update. Only `description` and `metadata` are mutable; any
	// other path returns `INVALID_PARAMETER_VALUE`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateSessionStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
