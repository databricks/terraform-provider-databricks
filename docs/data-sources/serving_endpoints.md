---
subcategory: "Serving"
---
# databricks_serving_endpoints Data Source

-> This resource can only be used with a workspace-level provider!

This resource allows you to get information about [Model Serving](https://docs.databricks.com/machine-learning/model-serving/index.html) endpoints in Databricks.

## Example Usage

```hcl
data "databricks_serving_endpoints" "all" {
}

resource "databricks_permissions" "ml_serving_usage" {
  for_each            = databricks_serving_endpoints.all.endpoints
  serving_endpoint_id = each.value.id

  access_control {
    group_name       = "users"
    permission_level = "CAN_VIEW"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_MANAGE"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_QUERY"
  }
}
```

## Attribute Reference

The following attributes are exported:

* `endpoints` - List of objects describing the serving endpoints. Each object consists of following attributes:
  * `name` - The name of the model serving endpoint.
  * `config` - The model serving endpoint configuration.
  * `tags` - Tags to be attached to the serving endpoint and automatically propagated to billing logs.
  * `rate_limits` - A list of rate limit blocks to be applied to the serving endpoint.
  * `ai_gateway` - A block with AI Gateway configuration for the serving endpoint.
  * `route_optimized` - A boolean enabling route optimization for the endpoint.

See [`databricks_model_serving` resource](../resources/model_serving.md) for the full list of attributes for each block

## Related Resources

The following resources are often used in the same context:

* [databricks_permissions](../resources/permissions.md#model-serving-usage) can control which groups or individual users can *Manage*, *Query* or *View* individual serving endpoints.
