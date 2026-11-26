// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aisearch

import (
	"context"
)

// **AI Search Endpoint**: Represents the compute resources to host AI Search
// indexes. AIP-conformant replacement for the legacy VectorSearchEndpoints API;
// functionally equivalent.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AiSearchService interface {

	// Create a new AI Search endpoint.
	CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (*Endpoint, error)

	// Create a new AI Search index. The index is provisioned asynchronously;
	// `status.ready` on GetIndex reports when it can serve queries.
	CreateIndex(ctx context.Context, request CreateIndexRequest) (*Index, error)

	// Delete an AI Search endpoint.
	DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error

	// Delete an AI Search index.
	DeleteIndex(ctx context.Context, request DeleteIndexRequest) error

	// Get details for a single AI Search endpoint.
	GetEndpoint(ctx context.Context, request GetEndpointRequest) (*Endpoint, error)

	// Get details for a single AI Search index.
	GetIndex(ctx context.Context, request GetIndexRequest) (*Index, error)

	// List AI Search endpoints in a workspace.
	ListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointsResponse, error)

	// List AI Search indexes on an endpoint.
	ListIndexes(ctx context.Context, request ListIndexesRequest) (*ListIndexesResponse, error)

	// Query (search) an AI Search index. Read-only, so a read-scoped token may
	// invoke it.
	QueryIndex(ctx context.Context, request QueryIndexRequest) (*QueryIndexResponse, error)

	// Remove rows by primary key from a Direct Access AI Search index.
	RemoveData(ctx context.Context, request RemoveDataRequest) (*RemoveDataResponse, error)

	// Scan (paginate over) the rows of an AI Search index.
	ScanIndex(ctx context.Context, request ScanIndexRequest) (*ScanIndexResponse, error)

	// Start a sync of a Delta Sync AI Search index with its source Delta table.
	// Use it with `TRIGGERED` pipelines; `CONTINUOUS` pipelines sync
	// automatically. Direct Access indexes are written through UpsertData
	// instead.
	SyncIndex(ctx context.Context, request SyncIndexRequest) (*SyncIndexResponse, error)

	// Update an existing AI Search endpoint. Only the fields named in
	// `update_mask` change. The server applies a mask that spans several
	// settings, such as `budget_policy_id` and `custom_tags`, one setting at a
	// time, so if a later setting fails, the earlier ones may already have been
	// applied.
	UpdateEndpoint(ctx context.Context, request UpdateEndpointRequest) (*Endpoint, error)

	// Upsert rows into a Direct Access AI Search index.
	UpsertData(ctx context.Context, request UpsertDataRequest) (*UpsertDataResponse, error)
}
