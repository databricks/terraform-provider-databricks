---
subcategory: "AI Search"
---
# databricks_ai_search_index Resource
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/aisearch)

# AI Search Index Resource

An AI Search Index is a searchable collection of records hosted on an AI Search endpoint. An index has a primary key and an index type (`DELTA_SYNC`, which keeps the index in sync with a source Delta table, or `DIRECT_ACCESS`, which is written to directly through the API). Indexes are immutable once created. This resource is the AIP-conformant replacement for the legacy `databricks_vector_search_index` resource and is functionally equivalent.


## Example Usage
# Example: AI Search Index Resource

An index is nested under an endpoint: the parent endpoint must be supplied, and
`index_id` is the index's Unity Catalog table name.

```hcl
resource "databricks_ai_search_index" "this" {
  endpoint_name = "example-ai-search-endpoint"
  index_id      = "main.default.example_index"
  primary_key   = "id"
  index_type    = "DELTA_SYNC"
}
```


## Arguments
The following arguments are supported:
* `index_type` (string, required) - Type of index. A `DELTA_SYNC` index is kept in sync with a source Delta table; a
  `DIRECT_ACCESS` index holds rows that you write through UpsertData and RemoveData. Set on
  create and immutable thereafter. Possible values are: `DELTA_SYNC`, `DIRECT_ACCESS`
* `parent` (string, required) - The Endpoint where this Index will be created.
  Format: `workspaces/{workspace_id}/endpoints/{endpoint_id}`
* `primary_key` (string, required) - Name of the column that uniquely identifies each row. It must exist in the source table
  (Delta Sync) or in `schema_json` (Direct Access) and cannot be the embedding source or vector
  column. Set on create and immutable thereafter
* `delta_sync_index_spec` (DeltaSyncIndexSpec, optional) - Specification for a Delta Sync index. Set when `index_type` is `DELTA_SYNC`
* `direct_access_index_spec` (DirectAccessIndexSpec, optional) - Specification for a Direct Access index. Set when `index_type` is `DIRECT_ACCESS`
* `index_id` (string, optional) - Three-level Unity Catalog table name for the new index, in the form `catalog.schema.name`.
  The server composes the full `Index.name` as `{parent}/indexes/{index_id}` and rejects an
  empty value
* `index_subtype` (string, optional) - Subtype of the index. Use `HYBRID` for vector and hybrid search, or `FULL_TEXT` for
  keyword-only search without embeddings; `VECTOR` is not a supported create value. Set on
  create and immutable thereafter. Possible values are: `FULL_TEXT`, `HYBRID`, `VECTOR`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### DeltaSyncIndexSpec
* `pipeline_type` (string, required) - Sync mode of the pipeline that keeps the index up to date. `TRIGGERED` syncs at creation and
  then each time SyncIndex is called; `CONTINUOUS` processes source changes as they arrive.
  Storage Optimized endpoints accept only `TRIGGERED`, and an unset value is rejected. Possible values are: `CONTINUOUS`, `TRIGGERED`
* `columns_to_sync` (list of string, optional) - Columns of the source table to include in the index. When empty, all columns are synced. The
  primary key and the embedding source or vector column are always synced. Only synced columns
  can be returned in query results or used in filters
* `embedding_source_columns` (list of EmbeddingSourceColumn, optional) - Columns whose text Databricks embeds during sync, using the model endpoint named on each
  column. Set either this field or `embedding_vector_columns`, not both
* `embedding_vector_columns` (list of EmbeddingVectorColumn, optional) - Columns of the source table that already hold precomputed embedding vectors. Set either this
  field or `embedding_source_columns`, not both
* `embedding_writeback_table` (string, optional) - Three-level name of a Delta table to write the synced rows and computed embeddings to. It
  applies only when embeddings are computed from `embedding_source_columns`, and the table must
  not already exist
* `source_table` (string, optional) - Three-level Unity Catalog name of the source Delta table, in the form `catalog.schema.table`

### DirectAccessIndexSpec
* `embedding_source_columns` (list of EmbeddingSourceColumn, optional) - Columns that contain the text to embed, each with the model endpoint that computes its
  embeddings
* `embedding_vector_columns` (list of EmbeddingVectorColumn, optional) - Columns that hold precomputed embedding vectors. A single vector column is supported, and it
  cannot be the primary key
* `schema_json` (string, optional) - The schema of the index in JSON format, mapping each column name to its type. Supported
  types are `integer`, `long`, `float`, `double`, `boolean`, `string`, `date`, `timestamp`.
  Supported types for vector columns: `array<float>`, `array<double>`. The schema must include
  the primary key column and the embedding vector column

### EmbeddingSourceColumn
* `embedding_model_endpoint` (string, optional) - Name of the embedding model endpoint, used by default for both ingestion and querying
* `model_endpoint_name_for_query` (string, optional) - Name of the embedding model endpoint which, if specified, is used for querying (not ingestion)
* `name` (string, optional) - Name of the column whose text is embedded

### EmbeddingVectorColumn
* `embedding_dimension` (integer, optional) - Number of elements in each embedding vector. It must be positive
* `name` (string, optional) - Name of the column that holds the precomputed embedding vectors

## Attributes
In addition to the above arguments, the following attributes are exported:
* `creator` (string) - Creator of the index
* `endpoint` (string) - Name of the endpoint associated with the index. Ignored on create — the endpoint is
  taken from `CreateIndexRequest.parent`; populated only on output
* `name` (string) - Name of the AI Search index. Server-assigned full resource path
  (`workspaces/{workspace}/endpoints/{endpoint}/indexes/{index}`) on output, where
  `{index}` is the index's Unity Catalog table name. On create, the user-supplied UC
  table name is conveyed via `CreateIndexRequest.index_id`; the server composes the
  full `name` and returns it on the response
* `status` (IndexStatus) - Current status of the index

### DeltaSyncIndexSpec
* `pipeline_id` (string) - The ID of the pipeline that is used to sync the index

### IndexStatus
* `index_url` (string) - Index API URL used to perform operations on the index
* `indexed_row_count` (integer) - Number of rows indexed
* `message` (string) - Human-readable detail about the index's current state
* `ready` (boolean) - Whether the index is ready to serve queries. After CreateIndex, poll GetIndex until this is
  true; `message` describes the current state in the meantime

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_ai_search_index.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_ai_search_index.this "name"
```