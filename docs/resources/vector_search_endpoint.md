---
subcategory: "Mosaic AI Vector Search"
---
# databricks_vector_search_endpoint Resource

This resource allows you to create [Mosaic AI Vector Search Endpoint](https://docs.databricks.com/en/generative-ai/vector-search.html) in Databricks.  Mosaic AI Vector Search is a serverless similarity search engine that allows you to store a vector representation of your data, including metadata, in a vector database.  The Mosaic AI Vector Search Endpoint is used to create and access vector search indexes.

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
resource "databricks_vector_search_endpoint" "this" {
  name          = "vector-search-test"
  endpoint_type = "STANDARD"
}
```

## Argument Reference

The following arguments are supported (change of any parameter leads to recreation of the resource):

* `name` - (Required) Name of the Mosaic AI Vector Search Endpoint to create.
* `endpoint_type` (Required) Type of Mosaic AI Vector Search Endpoint.  Currently only accepting single value: `STANDARD` (See [documentation](https://docs.databricks.com/api/workspace/vectorsearchendpoints/createendpoint) for the list of currently supported values).

## Attribute Reference

In addition to all the arguments above, the following attributes are exported:

* `id` - The same as the name of the endpoint.
* `creator` - Creator of the endpoint.
* `creation_timestamp` - Timestamp of endpoint creation (milliseconds).
* `last_updated_user` - User who last updated the endpoint.
* `last_updated_timestamp` - Timestamp of the last update to the endpoint (milliseconds).
* `endpoint_id` - Unique internal identifier of the endpoint (UUID).
* `num_indexes` - Number of indexes on the endpoint.
* `endpoint_status` - Object describing the current status of the endpoint consisting of the following fields:
  * `state` - Current state of the endpoint. Currently following values are supported: `PROVISIONING`, `ONLINE`, and `OFFLINE`.
  * `message` - Additional status message.

## Import

The resource can be imported using the name of the Mosaic AI Vector Search Endpoint

```hcl
import {
  to = databricks_vector_search_endpoint.this
  id = "<endpoint-name>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_vector_search_endpoint.this <endpoint-name>
```
