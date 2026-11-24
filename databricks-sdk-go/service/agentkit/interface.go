// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package agentkit

import (
	"context"
)

// APIs for managing agent memory and durable session state. This interface is
// under active development and may change.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AgentKitService interface {

	// Appends items to a session.
	AppendSessionItems(ctx context.Context, request AppendSessionItemsRequest) (*AppendSessionItemsResponse, error)

	// Clears all items from a session.
	ClearSessionItems(ctx context.Context, request ClearSessionItemsRequest) (*ClearSessionItemsResponse, error)

	// Creates a memory entry using exclusive-create semantics. Callers may
	// choose the entry ID; the service generates one when it is omitted.
	// Returns `ALREADY_EXISTS` when an entry with the same actor, session, and
	// path already exists. Omitted `session_id` is its own uniqueness key: two
	// omitted-session entries with the same actor and path conflict, but an
	// omitted-session entry does not conflict with a session-scoped entry at
	// the same actor and path. `session_id` is deprecated: once it is ignored,
	// the uniqueness key becomes actor and path.
	CreateMemory(ctx context.Context, request CreateManagedMemoryEntryRequest) (*ManagedMemoryEntry, error)

	// Creates a memory store in the caller's workspace. Returns
	// `ALREADY_EXISTS` when a store with the requested ID already exists.
	CreateMemoryStore(ctx context.Context, request CreateManagedMemoryStoreRequest) (*ManagedMemoryStore, error)

	// Creates a session within a session store.
	CreateSession(ctx context.Context, request CreateSessionRequest) (*Session, error)

	// Creates a session store.
	CreateSessionStore(ctx context.Context, request CreateSessionStoreRequest) (*SessionStore, error)

	// Deletes a memory entry by resource name. Returns `NOT_FOUND` when the
	// entry does not exist in the caller's workspace. The entry stops appearing
	// in Get, List and Search. In a store that records memory history, the
	// versions written before the delete, including their content, stay
	// readable through `ListMemoryChanges` and `GetMemoryChange` to anyone who
	// can read the store, until the store itself is deleted.
	DeleteMemory(ctx context.Context, request DeleteManagedMemoryEntryRequest) error

	// Deletes a memory store by resource name. Deleting a store also deletes
	// all of its memory entries, and its ID becomes available for reuse.
	// Returns `NOT_FOUND` when the store does not exist in the caller's
	// workspace.
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

	// Retrieves a memory entry by resource name. `view` controls whether the
	// response includes memory content. Returns `NOT_FOUND` when the entry does
	// not exist in the caller's workspace.
	GetMemory(ctx context.Context, request GetManagedMemoryEntryRequest) (*ManagedMemoryEntry, error)

	// Retrieves a memory store by resource name. Returns `NOT_FOUND` when the
	// store does not exist in the caller's workspace.
	GetMemoryStore(ctx context.Context, request GetManagedMemoryStoreRequest) (*ManagedMemoryStore, error)

	// Retrieves the current caller's effective memory store permission. Returns
	// PERMISSION_LEVEL_UNSPECIFIED when the caller has no access.
	GetMemoryStorePermission(ctx context.Context, request GetMemoryStorePermissionRequest) (*MemoryStorePermission, error)

	// Gets a session by resource name.
	GetSession(ctx context.Context, request GetSessionRequest) (*Session, error)

	// Gets a session store by resource name.
	GetSessionStore(ctx context.Context, request GetSessionStoreRequest) (*SessionStore, error)

	// Retrieves the current caller's effective session store permission.
	// Returns PERMISSION_LEVEL_UNSPECIFIED when the caller has no access.
	GetSessionStorePermission(ctx context.Context, request GetSessionStorePermissionRequest) (*SessionStorePermission, error)

	// Grants a READ or WRITE permission on a memory store to a principal. Only
	// the owner or a workspace admin may grant permissions.
	GrantMemoryStorePermission(ctx context.Context, request GrantMemoryStorePermissionRequest) (*GrantMemoryStorePermissionResponse, error)

	// Grants a READ or WRITE permission on a session store to a principal. Only
	// the owner or a workspace admin may grant permissions.
	GrantSessionStorePermission(ctx context.Context, request GrantSessionStorePermissionRequest) (*GrantSessionStorePermissionResponse, error)

	// Lists memory entries for one actor. An exact `path` filters entries
	// across sessions, ignoring session metadata. Otherwise, `path_prefix` and
	// the deprecated `session_id` restrict the actor partition. `view` controls
	// whether each returned entry includes memory content.
	ListMemories(ctx context.Context, request ListManagedMemoryEntriesRequest) (*ListManagedMemoryEntriesResponse, error)

	// Lists explicit memory store grants. Only the owner or a workspace admin
	// may list permissions.
	ListMemoryStorePermissions(ctx context.Context, request ListMemoryStorePermissionsRequest) (*ListMemoryStorePermissionsResponse, error)

	// Lists memory stores in the caller's workspace.
	ListMemoryStores(ctx context.Context, request ListManagedMemoryStoresRequest) (*ListManagedMemoryStoresResponse, error)

	// Lists items in a session.
	ListSessionItems(ctx context.Context, request ListSessionItemsRequest) (*ListSessionItemsResponse, error)

	// Lists explicit session store grants. Only the owner or a workspace admin
	// may list permissions.
	ListSessionStorePermissions(ctx context.Context, request ListSessionStorePermissionsRequest) (*ListSessionStorePermissionsResponse, error)

	// Lists session stores.
	ListSessionStores(ctx context.Context, request ListSessionStoresRequest) (*ListSessionStoresResponse, error)

	// Lists sessions within a session store.
	ListSessions(ctx context.Context, request ListSessionsRequest) (*ListSessionsResponse, error)

	// Pops the newest item from a session.
	PopSessionItem(ctx context.Context, request PopSessionItemRequest) (*PopSessionItemResponse, error)

	// Removes a principal's explicit READ and WRITE grants on a memory store.
	// Only the owner or a workspace admin may revoke permissions. Ownership and
	// inherited access are unaffected.
	RevokeMemoryStorePermission(ctx context.Context, request RevokeMemoryStorePermissionRequest) (*RevokeMemoryStorePermissionResponse, error)

	// Removes a principal's explicit READ and WRITE grants on a session store.
	// Only the owner or a workspace admin may revoke permissions. Ownership and
	// inherited access are unaffected.
	RevokeSessionStorePermission(ctx context.Context, request RevokeSessionStorePermissionRequest) (*RevokeSessionStorePermissionResponse, error)

	// Searches memory entries by text query for one actor. Returns matching
	// entries and scores ranked by relevance; `view` controls whether each
	// returned entry includes content.
	SearchMemories(ctx context.Context, request SearchManagedMemoryEntriesRequest) (*SearchManagedMemoryEntriesResponse, error)

	// Updates selected mutable fields on a memory entry. Identity fields are
	// immutable.
	UpdateMemory(ctx context.Context, request UpdateManagedMemoryEntryRequest) (*ManagedMemoryEntry, error)

	// Updates a memory store's description and metadata.
	UpdateMemoryStore(ctx context.Context, request UpdateManagedMemoryStoreRequest) (*ManagedMemoryStore, error)

	// Updates a session's mutable fields.
	UpdateSession(ctx context.Context, request UpdateSessionRequest) (*Session, error)

	// Updates a session store's description and metadata.
	UpdateSessionStore(ctx context.Context, request UpdateSessionStoreRequest) (*SessionStore, error)
}
