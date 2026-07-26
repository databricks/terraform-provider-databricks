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
* `index_type` (string, required) - Type of index. Required on create and immutable thereafter. Possible values are: `DELTA_SYNC`, `DIRECT_ACCESS`
* `parent` (string, required) - The Endpoint where this Index will be created.
  Format: `workspaces/{workspace_id}/endpoints/{endpoint_id}`
* `primary_key` (string, required) - Primary key of the index. Set on create and immutable thereafter
* `delta_sync_index_spec` (DeltaSyncIndexSpec, optional) - Specification for a Delta Sync index. Set when `index_type` is `DELTA_SYNC`
* `direct_access_index_spec` (DirectAccessIndexSpec, optional) - Specification for a Direct Access index. Set when `index_type` is `DIRECT_ACCESS`
* `index_id` (string, optional) - The user-supplied Unity Catalog table name for the Index, per AIP-133. The server
  composes the full `Index.name` as `{parent}/indexes/{index_id}`. AIP-133 does not
  list `index_id` as a fields-may-be-required entry, so we annotate it OPTIONAL on the
  wire; the server still rejects empty values with INVALID_PARAMETER_VALUE
* `index_subtype` (string, optional) - The subtype of the index. Set on create and immutable thereafter. Possible values are: `FULL_TEXT`, `HYBRID`, `VECTOR`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### DeltaSyncIndexSpec
* `pipeline_type` (string, required) - Pipeline execution mode. Required on create — the backend rejects an unset value.
  Storage Optimized endpoints accept only `TRIGGERED`; Standard endpoints accept both.
  No explicit `stage` — a REQUIRED field staged below its service would be dropped from
  combined specs while remaining in `required`, tripping the OpenAPI required-vs-properties
  consistency check. The field inherits the service's launch stage. Possible values are: `CONTINUOUS`, `TRIGGERED`
* `columns_to_sync` (list of string, optional) - [Optional] Select the columns to sync with the index. If left blank, all columns
  from the source table are synced. The primary key column and embedding source or
  vector column are always synced
* `embedding_source_columns` (list of EmbeddingSourceColumn, optional) - The columns that contain the embedding source
* `embedding_vector_columns` (list of EmbeddingVectorColumn, optional) - The columns that contain the embedding vectors
* `embedding_writeback_table` (string, optional) - [Optional] Name of the Delta table to sync the index contents and computed embeddings to
* `source_table` (string, optional) - The full name of the source Delta table

### DirectAccessIndexSpec
* `embedding_source_columns` (list of EmbeddingSourceColumn, optional) - The columns that contain the embedding source
* `embedding_vector_columns` (list of EmbeddingVectorColumn, optional) - The columns that contain the embedding vectors
* `schema_json` (string, optional) - The schema of the index in JSON format. Supported types are `integer`, `long`,
  `float`, `double`, `boolean`, `string`, `date`, `timestamp`. Supported types for
  vector columns: `array<float>`, `array<double>`

### EmbeddingSourceColumn
* `embedding_model_endpoint` (string, optional) - Name of the embedding model endpoint, used by default for both ingestion and querying
* `model_endpoint_name_for_query` (string, optional) - Name of the embedding model endpoint which, if specified, is used for querying (not ingestion)
* `name` (string, optional) - Name of the source column

### EmbeddingVectorColumn
* `embedding_dimension` (integer, optional) - Dimension of the embedding vector
* `name` (string, optional) - Name of the column

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
* `ready` (boolean) - Whether the index is ready for search

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