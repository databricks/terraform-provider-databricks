// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mason

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
	// The managed memory entry to create.
	ManagedMemoryEntry ManagedMemoryEntry `json:"managed_memory_entry"`
	// Optional caller-selected managed memory entry ID. The service generates
	// an ID when omitted.
	ManagedMemoryEntryId string `json:"-" url:"managed_memory_entry_id,omitempty"`
	// Managed memory store that will contain the entry, in the form
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
	// The managed memory store to create.
	ManagedMemoryStore ManagedMemoryStore `json:"managed_memory_store"`
	// Caller-provided, workspace-unique managed memory store ID. It must be
	// 3-56 characters, begin with a lowercase letter, contain only lowercase
	// letters, digits, and hyphens, and end with a letter or digit.
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
	// Managed memory store the extracted entries are written to, in the form
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
	// Fields to return, using proto field names such as `content` (not
	// `contents`). An omitted or empty mask returns the full entry, including
	// `content`; a non-empty mask returns only the requested fields.
	ReadMask *fieldmask.FieldMask `json:"-" url:"read_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetManagedMemoryEntryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetManagedMemoryEntryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetManagedMemoryStoreRequest struct {
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetManagedMemoryStoreRequest) UnmarshalJSON(b []byte) error {
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

type GetSessionStoreRequest struct {
	// Resource name in the form `session-stores/{session_store_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetSessionStoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListManagedMemoryEntriesRequest struct {
	// Customer-provided identifier for the actor whose entries are listed.
	ActorId string `json:"-" url:"actor_id"`
	// Maximum number of entries to return. The service may return fewer entries
	// than requested. Defaults to 10; must be between 1 and 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque pagination token from a previous ListManagedMemoryEntries
	// response.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Managed memory store whose entries are listed, in the form
	// `memory-stores/{managed_memory_store_id}`.
	Parent string `json:"-" url:"-"`
	// Optional path prefix used to restrict entries within the actor partition.
	PathPrefix string `json:"-" url:"path_prefix,omitempty"`
	// Fields to return in each entry, using proto field names such as `content`
	// (not `contents`). An omitted or empty mask returns each full entry,
	// including `content`; a non-empty mask returns only the requested fields.
	ReadMask *fieldmask.FieldMask `json:"-" url:"read_mask,omitempty"`
	// Optional session identifier. When set, only entries with this exact
	// `session_id` are returned. Omitted-session (cross-session) entries are
	// not included. Ignored when path is set.
	SessionId string `json:"-" url:"session_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListManagedMemoryEntriesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListManagedMemoryEntriesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing managed memory entries.
type ListManagedMemoryEntriesResponse struct {
	// Managed memory entries matching the request and its read mask.
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
	// Opaque pagination token from a previous ListManagedMemoryStores response.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListManagedMemoryStoresRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListManagedMemoryStoresRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing managed memory stores in the caller's workspace.
type ListManagedMemoryStoresResponse struct {
	// Managed memory stores in the caller's workspace.
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

// A workspace-scoped entry in a managed memory store.
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
	Name string `json:"name,omitempty"`
	// Absolute, case-sensitive path identifying the entry within its actor and
	// optional session. Paths must begin with `/` and must not contain empty,
	// `.` or `..` segments.
	Path string `json:"path"`
	// Optional identifier for the session associated with this memory entry.
	// When omitted, the entry applies across the actor's sessions.
	SessionId string `json:"session_id,omitempty"`
	// Which writer created this entry. Caller sets this on Create; immutable
	// after creation.
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

// One relevance-ranked managed memory search result.
type ManagedMemoryEntrySearchResult struct {
	// Managed memory entry matching the query.
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

// Identifies the source that created a managed memory entry.
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

// A workspace-scoped managed memory store backed by service-managed storage.
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
	// Deprecated compatibility alias for the caller-provided managed memory
	// store ID. Canonical clients provide the ID through
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

// Request to search managed memory entries by text query for one actor. Search
// returns a relevance-ranked top-N result set and does not currently paginate.
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
	// Managed memory store whose entries are searched, in the form
	// `memory-stores/{managed_memory_store_id}`.
	Parent string `json:"-" url:"-"`
	// Optional absolute, case-sensitive path prefix used to restrict searched
	// entries within the actor partition. The prefix must begin with `/` and
	// must not contain empty, `.` or `..` segments.
	PathPrefix string `json:"path_prefix,omitempty"`
	// Free-form search query.
	Query string `json:"query"`
	// Fields to return in each matching entry, using proto field names such as
	// `content` (not `contents`). An omitted or empty mask returns each full
	// entry, including `content`; a non-empty mask returns only the requested
	// fields. Search scores are always returned.
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	ReadMask *fieldmask.FieldMask `json:"read_mask,omitempty"`
	// Optional session identifier. When set, only entries with this exact
	// `session_id` are searched. Omitted-session (cross-session) entries are
	// not included.
	SessionId string `json:"session_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SearchManagedMemoryEntriesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SearchManagedMemoryEntriesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing managed memory entries ranked by relevance.
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

// Service-managed storage backing a managed memory store.
type StorageBackend struct {
	// Backend-specific identifier. For Lakebase, this is the project ID.
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

// Type of service-managed storage backing a managed memory store.
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

type UpdateManagedMemoryEntryRequest struct {
	// The managed memory entry to update.
	ManagedMemoryEntry ManagedMemoryEntry `json:"managed_memory_entry"`
	// Resource name in the form
	// `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`.
	Name string `json:"-" url:"-"`
	// Fields to update. Only `content` and `description` may be updated.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateManagedMemoryEntryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateManagedMemoryStoreRequest struct {
	// The managed memory store to update. `name` is taken from the URL.
	ManagedMemoryStore ManagedMemoryStore `json:"managed_memory_store"`
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name string `json:"-" url:"-"`
	// Only `description` may be updated.
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
