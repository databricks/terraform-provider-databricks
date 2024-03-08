---
subcategory: "Vector Search"
---
# databricks_vector_search_endpoint Resource

-> **Note** This resource could be only used on Unity Catalog-enabled workspace!

This resource allows you to create [Vector Search Endpoint](https://docs.databricks.com/en/generative-ai/vector-search.html) in Databricks.  Vector Search is a serverless similarity search engine that allows you to store a vector representation of your data, including metadata, in a vector database.  The Vector Search Endpoint is used to create and access vector search indexes.

## Example Usage

```hcl
resource "databricks_vector_search_endpoint" "this" {
  name          = "vector-search-test"
  endpoint_type = "STANDARD"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the Vector Search Endpoint to create.  If name is changed, Vector Search Endpoint is recreated.
* `endpoint_type` (Required) type of Vector Search Endpoint.  Currently only accepting single value: `STANDARD` (See [documentation](https://docs.databricks.com/api/workspace/vectorsearchendpoints/createendpoint) for the list of currently supported values).  If it's changed, Vector Search Endpoint is recreated.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The same as the name of the endpoint.
* `creator` - Creator of the endpoint.
* `creation_timestamp` - Timestamp of endpoint creation (milliseconds).
* `last_updated_user` - User who last updated the endpoint.
* `last_updated_timestamp` - Timestamp of last update to the endpoint (milliseconds).
* `endpoint_id` - Unique internal identifier of the endpoint (UUID).
* `num_indexes` - Number of indexes on the endpoint.
* `endpoint_status` - Object describing the current status of the endpoint consisting of following fields:
  * `state` - Current state of the endpoint. Currently following values are supported: `PROVISIONING`, `ONLINE`, `OFFLINE`.
  * `message` - Additional status message.

## Import

The resource can be imported using the name of the Vector Search Endpoint

```bash
terraform import databricks_vector_search_endpoint.this <endpoint-name>
```
