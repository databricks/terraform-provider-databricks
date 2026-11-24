// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mason

import (
	"context"
)

// APIs for managing agent memory and durable session state. This interface is
// under active development and may change.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type MasonService interface {

	// Appends items to a session.
	AppendSessionItems(ctx context.Context, request AppendSessionItemsRequest) (*AppendSessionItemsResponse, error)

	// Clears all items from a session.
	ClearSessionItems(ctx context.Context, request ClearSessionItemsRequest) (*ClearSessionItemsResponse, error)

	// Creates a managed memory entry using exclusive-create semantics. Callers
	// may choose the entry ID; the service generates one when it is omitted.
	// Returns `ALREADY_EXISTS` when an entry with the same actor, session, and
	// path already exists. Omitted `session_id` is its own uniqueness key: two
	// omitted-session entries with the same actor and path conflict, but an
	// omitted-session entry does not conflict with a session-scoped entry at
	// the same actor and path.
	CreateMemory(ctx context.Context, request CreateManagedMemoryEntryRequest) (*ManagedMemoryEntry, error)

	// Creates a managed memory store in the caller's workspace.
	CreateMemoryStore(ctx context.Context, request CreateManagedMemoryStoreRequest) (*ManagedMemoryStore, error)

	// Creates a session within a session store.
	CreateSession(ctx context.Context, request CreateSessionRequest) (*Session, error)

	// Creates a session store.
	CreateSessionStore(ctx context.Context, request CreateSessionStoreRequest) (*SessionStore, error)

	// Deletes a managed memory entry by resource name. Returns `NOT_FOUND` when
	// the entry does not exist in the caller's workspace.
	DeleteMemory(ctx context.Context, request DeleteManagedMemoryEntryRequest) error

	// Deletes a managed memory store by resource name. Returns `NOT_FOUND` when
	// the store does not exist in the caller's workspace.
	DeleteMemoryStore(ctx context.Context, request DeleteManagedMemoryStoreRequest) error

	// Deletes a session, its items, and any descendant sessions recursively.
	// Independently retained memory is not deleted.
	DeleteSession(ctx context.Context, request DeleteSessionRequest) error

	// Deletes a session store, its sessions and items, and its service-managed
	// storage. Memory entries retained by a separate Memory Store are not
	// deleted.
	DeleteSessionStore(ctx context.Context, request DeleteSessionStoreRequest) error

	// Synchronously extracts memories from a single session into the given
	// memory store, returning the entries that were written.
	ExtractMemories(ctx context.Context, request ExtractMemoriesRequest) (*ExtractMemoriesResponse, error)

	// Forks a session into an independent top-level copy.
	ForkSession(ctx context.Context, request ForkSessionRequest) (*ForkSessionResponse, error)

	// Retrieves a managed memory entry, including its content, by resource
	// name. Returns `NOT_FOUND` when the entry does not exist in the caller's
	// workspace.
	GetMemory(ctx context.Context, request GetManagedMemoryEntryRequest) (*ManagedMemoryEntry, error)

	// Retrieves a managed memory store by resource name. Returns `NOT_FOUND`
	// when the store does not exist in the caller's workspace.
	GetMemoryStore(ctx context.Context, request GetManagedMemoryStoreRequest) (*ManagedMemoryStore, error)

	// Gets a session by resource name.
	GetSession(ctx context.Context, request GetSessionRequest) (*Session, error)

	// Gets a session store by resource name.
	GetSessionStore(ctx context.Context, request GetSessionStoreRequest) (*SessionStore, error)

	// Lists managed memory entries for one actor. An exact `path` filters
	// entries across sessions, ignoring session metadata. Otherwise,
	// `session_id` and `path_prefix` restrict the actor partition. `read_mask`
	// selects fields in each returned entry.
	ListMemories(ctx context.Context, request ListManagedMemoryEntriesRequest) (*ListManagedMemoryEntriesResponse, error)

	// Lists managed memory stores in the caller's workspace.
	ListMemoryStores(ctx context.Context, request ListManagedMemoryStoresRequest) (*ListManagedMemoryStoresResponse, error)

	// Lists items in a session.
	ListSessionItems(ctx context.Context, request ListSessionItemsRequest) (*ListSessionItemsResponse, error)

	// Lists session stores.
	ListSessionStores(ctx context.Context, request ListSessionStoresRequest) (*ListSessionStoresResponse, error)

	// Lists sessions within a session store.
	ListSessions(ctx context.Context, request ListSessionsRequest) (*ListSessionsResponse, error)

	// Pops the newest item from a session.
	PopSessionItem(ctx context.Context, request PopSessionItemRequest) (*PopSessionItemResponse, error)

	// Searches managed memory entries by text query for one actor. Returns
	// matching entries and scores ranked by relevance; `read_mask` selects
	// fields in each returned entry.
	SearchMemories(ctx context.Context, request SearchManagedMemoryEntriesRequest) (*SearchManagedMemoryEntriesResponse, error)

	// Updates selected mutable fields on a managed memory entry. Identity
	// fields are immutable.
	UpdateMemory(ctx context.Context, request UpdateManagedMemoryEntryRequest) (*ManagedMemoryEntry, error)

	// Updates a managed memory store's description.
	UpdateMemoryStore(ctx context.Context, request UpdateManagedMemoryStoreRequest) (*ManagedMemoryStore, error)

	// Updates a session's mutable fields.
	UpdateSession(ctx context.Context, request UpdateSessionRequest) (*Session, error)

	// Updates a session store's description and metadata.
	UpdateSessionStore(ctx context.Context, request UpdateSessionStoreRequest) (*SessionStore, error)
}
