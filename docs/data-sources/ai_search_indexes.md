---
subcategory: "AI Search"
---
# databricks_ai_search_indexes Data Source
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/aisearch)

# Datasource (Plural) Artifact

This data source can be used to fetch the list of AI Search indexes on an endpoint.


## Example Usage
# Example: AI Search Index Datasource (Plural)

```hcl
data "databricks_ai_search_indexes" "all" {
  endpoint_name = "example-ai-search-endpoint"
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The Endpoint that owns this collection of indexes.
  Format: `workspaces/{workspace_id}/endpoints/{endpoint_id}`
* `page_size` (integer, optional) - Maximum number of indexes to return. When unset, every index on the endpoint is returned
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `indexes`. It is a list of resources, each with the following attributes:
* `creator` (string) - Creator of the index
* `delta_sync_index_spec` (DeltaSyncIndexSpec) - Specification for a Delta Sync index. Set when `index_type` is `DELTA_SYNC`
* `direct_access_index_spec` (DirectAccessIndexSpec) - Specification for a Direct Access index. Set when `index_type` is `DIRECT_ACCESS`
* `endpoint` (string) - Name of the endpoint associated with the index. Ignored on create — the endpoint is
  taken from `CreateIndexRequest.parent`; populated only on output
* `index_subtype` (string) - Subtype of the index. Use `HYBRID` for vector and hybrid search, or `FULL_TEXT` for
  keyword-only search without embeddings; `VECTOR` is not a supported create value. Set on
  create and immutable thereafter. Possible values are: `FULL_TEXT`, `HYBRID`, `VECTOR`
* `index_type` (string) - Type of index. A `DELTA_SYNC` index is kept in sync with a source Delta table; a
  `DIRECT_ACCESS` index holds rows that you write through UpsertData and RemoveData. Set on
  create and immutable thereafter. Possible values are: `DELTA_SYNC`, `DIRECT_ACCESS`
* `name` (string) - Name of the AI Search index. Server-assigned full resource path
  (`workspaces/{workspace}/endpoints/{endpoint}/indexes/{index}`) on output, where
  `{index}` is the index's Unity Catalog table name. On create, the user-supplied UC
  table name is conveyed via `CreateIndexRequest.index_id`; the server composes the
  full `name` and returns it on the response
* `primary_key` (string) - Name of the column that uniquely identifies each row. It must exist in the source table
  (Delta Sync) or in `schema_json` (Direct Access) and cannot be the embedding source or vector
  column. Set on create and immutable thereafter
* `status` (IndexStatus) - Current status of the index

### DeltaSyncIndexSpec
* `columns_to_sync` (list of string) - Columns of the source table to include in the index. When empty, all columns are synced. The
  primary key and the embedding source or vector column are always synced. Only synced columns
  can be returned in query results or used in filters
* `embedding_source_columns` (list of EmbeddingSourceColumn) - Columns whose text Databricks embeds during sync, using the model endpoint named on each
  column. Set either this field or `embedding_vector_columns`, not both
* `embedding_vector_columns` (list of EmbeddingVectorColumn) - Columns of the source table that already hold precomputed embedding vectors. Set either this
  field or `embedding_source_columns`, not both
* `embedding_writeback_table` (string) - Three-level name of a Delta table to write the synced rows and computed embeddings to. It
  applies only when embeddings are computed from `embedding_source_columns`, and the table must
  not already exist
* `pipeline_id` (string) - The ID of the pipeline that is used to sync the index
* `pipeline_type` (string) - Sync mode of the pipeline that keeps the index up to date. `TRIGGERED` syncs at creation and
  then each time SyncIndex is called; `CONTINUOUS` processes source changes as they arrive.
  Storage Optimized endpoints accept only `TRIGGERED`, and an unset value is rejected. Possible values are: `CONTINUOUS`, `TRIGGERED`
* `source_table` (string) - Three-level Unity Catalog name of the source Delta table, in the form `catalog.schema.table`

### DirectAccessIndexSpec
* `embedding_source_columns` (list of EmbeddingSourceColumn) - Columns that contain the text to embed, each with the model endpoint that computes its
  embeddings
* `embedding_vector_columns` (list of EmbeddingVectorColumn) - Columns that hold precomputed embedding vectors. A single vector column is supported, and it
  cannot be the primary key
* `schema_json` (string) - The schema of the index in JSON format, mapping each column name to its type. Supported
  types are `integer`, `long`, `float`, `double`, `boolean`, `string`, `date`, `timestamp`.
  Supported types for vector columns: `array<float>`, `array<double>`. The schema must include
  the primary key column and the embedding vector column

### EmbeddingSourceColumn
* `embedding_model_endpoint` (string) - Name of the embedding model endpoint, used by default for both ingestion and querying
* `model_endpoint_name_for_query` (string) - Name of the embedding model endpoint which, if specified, is used for querying (not ingestion)
* `name` (string) - Name of the column whose text is embedded

### EmbeddingVectorColumn
* `embedding_dimension` (integer) - Number of elements in each embedding vector. It must be positive
* `name` (string) - Name of the column that holds the precomputed embedding vectors

### IndexStatus
* `index_url` (string) - Index API URL used to perform operations on the index
* `indexed_row_count` (integer) - Number of rows indexed
* `message` (string) - Human-readable detail about the index's current state
* `ready` (boolean) - Whether the index is ready to serve queries. After CreateIndex, poll GetIndex until this is
  true; `message` describes the current state in the meantime