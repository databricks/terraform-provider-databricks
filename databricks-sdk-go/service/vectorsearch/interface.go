// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"context"
)

// **Endpoint**: Represents the compute resources to host AI Search indexes.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type VectorSearchEndpointsService interface {

	// Create a new endpoint.
	CreateEndpoint(ctx context.Context, request CreateEndpoint) (*EndpointInfo, error)

	// Delete an AI Search endpoint.
	DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error

	// Get details for a single AI Search endpoint.
	GetEndpoint(ctx context.Context, request GetEndpointRequest) (*EndpointInfo, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetVectorSearchEndpointPermissionLevelsRequest) (*GetVectorSearchEndpointPermissionLevelsResponse, error)

	// Gets the permissions of a vector search endpoint. Vector search endpoints
	// can inherit permissions from their root object.
	GetPermissions(ctx context.Context, request GetVectorSearchEndpointPermissionsRequest) (*VectorSearchEndpointPermissions, error)

	// List all AI Search endpoints in the workspace.
	ListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointResponse, error)

	// Update the Target QPS of a STANDARD endpoint. The response is the updated
	// endpoint; its scaling_info shows whether the resulting scaling change is
	// applied or still in progress.
	PatchEndpoint(ctx context.Context, request PatchEndpointRequest) (*EndpointInfo, error)

	// Retrieve time series of usage and health metrics for an endpoint and its
	// indexes, such as request counts, request latency percentiles and endpoint
	// utilization.
	RetrieveUserVisibleMetrics(ctx context.Context, request RetrieveUserVisibleMetricsRequest) (*RetrieveUserVisibleMetricsResponse, error)

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request VectorSearchEndpointPermissionsRequest) (*VectorSearchEndpointPermissions, error)

	// Start a stopped endpoint. The endpoint serves the same indexes it had
	// when it stopped.
	StartEndpoint(ctx context.Context, request StartEndpointRequest) (*EndpointInfo, error)

	// Stop an endpoint. Stopping removes the endpoint's serving compute and
	// retains its indexes. A stopped endpoint can't serve queries or sync
	// indexes until it is started.
	StopEndpoint(ctx context.Context, request StopEndpointRequest) (*EndpointInfo, error)

	// Update the budget policy that the endpoint's usage is attributed to.
	UpdateEndpointBudgetPolicy(ctx context.Context, request PatchEndpointBudgetPolicyRequest) (*PatchEndpointBudgetPolicyResponse, error)

	// Update the custom tags of an endpoint.
	UpdateEndpointCustomTags(ctx context.Context, request UpdateEndpointCustomTagsRequest) (*UpdateEndpointCustomTagsResponse, error)

	// Updates the permissions on a vector search endpoint. Vector search
	// endpoints can inherit permissions from their root object.
	UpdatePermissions(ctx context.Context, request VectorSearchEndpointPermissionsRequest) (*VectorSearchEndpointPermissions, error)
}

// **Index**: An efficient representation of your embedding vectors that
// supports real-time and efficient approximate nearest neighbor (ANN) search
// queries.
//
// There are 2 types of AI Search indexes: - **Delta Sync Index**: An index that
// automatically syncs with a source Delta Table, automatically and
// incrementally updating the index as the underlying data in the Delta Table
// changes. - **Direct Vector Access Index**: An index that supports direct read
// and write of vectors and metadata through our REST and SDK APIs. With this
// model, the user manages index updates.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type VectorSearchIndexesService interface {

	// Create a new index on an endpoint. Creating a Delta Sync index also
	// creates the pipeline that populates it from the source table; a Direct
	// Vector Access index starts empty and is populated through Upsert data
	// into an index.
	CreateIndex(ctx context.Context, request CreateVectorIndexRequest) (*VectorIndex, error)

	// Delete rows from a Direct Vector Access index by primary key. The
	// response reports how many rows were deleted and the primary keys of the
	// rows that failed.
	DeleteDataVectorIndex(ctx context.Context, request DeleteDataVectorIndexRequest) (*DeleteDataVectorIndexResponse, error)

	// Delete an index and the data it holds. The source table of a Delta Sync
	// index is not affected.
	DeleteIndex(ctx context.Context, request DeleteIndexRequest) error

	// Get an index.
	GetIndex(ctx context.Context, request GetIndexRequest) (*VectorIndex, error)

	// List all indexes in the given endpoint.
	ListIndexes(ctx context.Context, request ListIndexesRequest) (*ListVectorIndexesResponse, error)

	// Query an index with a vector, text, or both, and return the requested
	// columns of the best-matching rows.
	QueryIndex(ctx context.Context, request QueryVectorIndexRequest) (*QueryVectorIndexResponse, error)

	// Use `next_page_token` returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` request to fetch next page of results.
	QueryNextPage(ctx context.Context, request QueryVectorIndexNextPageRequest) (*QueryVectorIndexResponse, error)

	// Scan an index and return the first `num_results` rows after the exclusive
	// `last_primary_key`. Repeat the call with the returned `last_primary_key`
	// to read the whole index.
	ScanIndex(ctx context.Context, request ScanVectorIndexRequest) (*ScanVectorIndexResponse, error)

	// Start a sync of a Delta Sync index with a `TRIGGERED` pipeline from its
	// source table. An index with a `CONTINUOUS` pipeline syncs automatically
	// and rejects manual syncs.
	SyncIndex(ctx context.Context, request SyncIndexRequest) error

	// Insert rows into a Direct Vector Access index, replacing existing rows
	// that have the same primary key. The response reports how many rows
	// succeeded and the primary keys of the rows that failed.
	UpsertDataVectorIndex(ctx context.Context, request UpsertDataVectorIndexRequest) (*UpsertDataVectorIndexResponse, error)
}
