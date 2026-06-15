---
subcategory: "AI Search"
---
# databricks_ai_search_indexes Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `parent` (string, required) - The Endpoint that owns this collection of indexes.
  Format: `workspaces/{workspace_id}/endpoints/{endpoint_id}`
* `page_size` (integer, optional) - Best-effort upper bound on the number of results to return. Honored as an upper
  bound by the shim: `page_size` only narrows the legacy backend's response, never
  widens it, so the practical cap is `min(page_size, legacy_fixed_page_size)`
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
* `index_subtype` (string) - The subtype of the index. Set on create and immutable thereafter. Possible values are: `FULL_TEXT`, `HYBRID`, `VECTOR`
* `index_type` (string) - Type of index. Required on create and immutable thereafter. Possible values are: `DELTA_SYNC`, `DIRECT_ACCESS`
* `name` (string) - Name of the AI Search index. Server-assigned full resource path
  (`workspaces/{workspace}/endpoints/{endpoint}/indexes/{index}`) on output, where
  `{index}` is the index's Unity Catalog table name. On create, the user-supplied UC
  table name is conveyed via `CreateIndexRequest.index_id`; the server composes the
  full `name` and returns it on the response
* `primary_key` (string) - Primary key of the index. Set on create and immutable thereafter
* `status` (IndexStatus) - Current status of the index

### DeltaSyncIndexSpec
* `columns_to_sync` (list of string) - [Optional] Select the columns to sync with the index. If left blank, all columns
  from the source table are synced. The primary key column and embedding source or
  vector column are always synced
* `embedding_source_columns` (list of EmbeddingSourceColumn) - The columns that contain the embedding source
* `embedding_vector_columns` (list of EmbeddingVectorColumn) - The columns that contain the embedding vectors
* `embedding_writeback_table` (string) - [Optional] Name of the Delta table to sync the index contents and computed embeddings to
* `pipeline_id` (string) - The ID of the pipeline that is used to sync the index
* `pipeline_type` (string) - Pipeline execution mode. Required on create — the backend rejects an unset value.
  Storage Optimized endpoints accept only `TRIGGERED`; Standard endpoints accept both.
  No explicit `stage` — a REQUIRED field staged below its service would be dropped from
  combined specs while remaining in `required`, tripping the OpenAPI required-vs-properties
  consistency check. The field inherits the service's launch stage. Possible values are: `CONTINUOUS`, `TRIGGERED`
* `source_table` (string) - The full name of the source Delta table

### DirectAccessIndexSpec
* `embedding_source_columns` (list of EmbeddingSourceColumn) - The columns that contain the embedding source
* `embedding_vector_columns` (list of EmbeddingVectorColumn) - The columns that contain the embedding vectors
* `schema_json` (string) - The schema of the index in JSON format. Supported types are `integer`, `long`,
  `float`, `double`, `boolean`, `string`, `date`, `timestamp`. Supported types for
  vector columns: `array<float>`, `array<double>`

### EmbeddingSourceColumn
* `embedding_model_endpoint` (string) - Name of the embedding model endpoint, used by default for both ingestion and querying
* `model_endpoint_name_for_query` (string) - Name of the embedding model endpoint which, if specified, is used for querying (not ingestion)
* `name` (string) - Name of the source column

### EmbeddingVectorColumn
* `embedding_dimension` (integer) - Dimension of the embedding vector
* `name` (string) - Name of the column

### IndexStatus
* `index_url` (string) - Index API URL used to perform operations on the index
* `indexed_row_count` (integer) - Number of rows indexed
* `message` (string) - Human-readable detail about the index's current state
* `ready` (boolean) - Whether the index is ready for search