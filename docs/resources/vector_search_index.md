---
subcategory: "Mosaic AI Vector Search"
---
# databricks_vector_search_index Resource

-> **Note** This resource could be only used on Unity Catalog-enabled workspace!

This resource allows you to create [Mosaic AI Vector Search Index](https://docs.databricks.com/en/generative-ai/create-query-vector-search.html) in Databricks.  Mosaic AI Vector Search is a serverless similarity search engine that allows you to store a vector representation of your data, including metadata, in a vector database.  The Mosaic AI Vector Search Index provides the ability to search data in the linked Delta Table.

## Example Usage

```hcl
resource "databricks_vector_search_index" "sync" {
  name          = "main.default.vector_search_index"
  endpoint_name = databricks_vector_search_endpoint.this.name
  primary_key   = "id"
  index_type    = "DELTA_SYNC"
  delta_sync_index_spec {
    source_table  = "main.default.source_table"
    pipeline_type = "TRIGGERED"
    embedding_source_columns {
      name                          = "text"
      embedding_model_endpoint_name = databricks_model_serving.this.name
    }
  }
}
```

## Argument Reference

The following arguments are supported (change of any parameter leads to recreation of the resource):

* `name` - (required) Three-level name of the Mosaic AI Vector Search Index to create (`catalog.schema.index_name`).
* `endpoint_name` - (required) The name of the Mosaic AI Vector Search Endpoint that will be used for indexing the data.
* `primary_key` - (required) The column name that will be used as a primary key.
* `index_type` - (required) Mosaic AI Vector Search index type. Currently supported values are:
  * `DELTA_SYNC`: An index that automatically syncs with a source Delta Table, automatically and incrementally updating the index as the underlying data in the Delta Table changes.
  * `DIRECT_ACCESS`: An index that supports the direct read and write of vectors and metadata through our REST and SDK APIs. With this model, the user manages index updates.
* `delta_sync_index_spec` - (object) Specification for Delta Sync Index. Required if `index_type` is `DELTA_SYNC`.
  * `source_table` (required) The name of the source table.
  * `embedding_source_columns` - (required if `embedding_vector_columns` isn't provided) array of objects representing columns that contain the embedding source.  Each entry consists of:
	* `name` - The name of the column
	* `embedding_model_endpoint_name` - The name of the embedding model endpoint
  * `embedding_vector_columns`  - (required if `embedding_source_columns` isn't provided)  array of objects representing columns that contain the embedding vectors. Each entry consists of:
	* `name` - The name of the column.
	* `embedding_dimension` - Dimension of the embedding vector.
  * `pipeline_type` - Pipeline execution mode. Possible values are:
	* `TRIGGERED`: If the pipeline uses the triggered execution mode, the system stops processing after successfully refreshing the source table in the pipeline once, ensuring the table is updated based on the data available when the update started.
	* `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline processes new data as it arrives in the source table to keep the vector index fresh.
* `direct_access_index_spec` - (object) Specification for Direct Vector Access Index. Required if `index_type` is `DIRECT_ACCESS`.
  * `schema_json` - The schema of the index in JSON format.  Check the [API documentation](https://docs.databricks.com/api/workspace/vectorsearchindexes/createindex#direct_access_index_spec-schema_json) for a list of supported data types.
  * `embedding_source_columns` - (required if `embedding_vector_columns` isn't provided) array of objects representing columns that contain the embedding source.  Each entry consists of:
	* `name` - The name of the column
	* `embedding_model_endpoint_name` - The name of the embedding model endpoint
  * `embedding_vector_columns`  - (required if `embedding_source_columns` isn't provided)  array of objects representing columns that contain the embedding vectors. Each entry consists of:
	* `name` - The name of the column.
	* `embedding_dimension` - Dimension of the embedding vector.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The same as the name of the index.
* `creator` - Creator of the endpoint.
* `delta_sync_index_spec`:
  * `pipeline_id` - ID of the associated Delta Live Table pipeline.
* `status` - Object describing the current status of the index consisting of the following fields:
  * `message` - Message associated with the index status
  * `indexed_row_count` - Number of rows indexed
  * `ready` - Whether the index is ready for search
  * `index_url` - Index API Url to be used to perform operations on the index

## Import

The resource can be imported using the name of the Mosaic AI Vector Search Index

```bash
terraform import databricks_vector_search_index.this <index-name>
```
