---
subcategory: "Serving"
---
# databricks_model_serving_provisioned_throughput Resource

This resource allows you to manage [Foundation Model provisioned throughput](https://docs.databricks.com/aws/en/machine-learning/foundation-model-apis/deploy-prov-throughput-foundation-model-apis) endpoints in Databricks.

~> This resource is currently in private preview, and only available for enrolled customers.

-> This resource can only be used with a workspace-level provider!

## Example Usage

Creating a Foundation Model provisioned throughput endpoint

```hcl
resource "databricks_model_serving_provisioned_throughput" "llama" {
  ai_gateway {
    usage_tracking_config {
      enabled = true
    }
  }
  config {
    served_entities {
      entity_name                = "system.ai.llama-4-maverick"
      entity_version             = "1"
      provisioned_model_units    = 100
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the model serving endpoint. This field is required and must be unique across a workspace. An endpoint name can consist of alphanumeric characters, dashes, and underscores. NOTE: Changing this name will delete the existing endpoint and create a new endpoint with the updated name.
* `config` - The model serving endpoint configuration.
  * `served_entities` - A list of served entities for the endpoint to serve.
  * `traffic_config` - A single block represents the traffic split configuration amongst the served models.
* `tags` - Tags to be attached to the serving endpoint and automatically propagated to billing logs.
* `ai_gateway` - (Optional) A block with AI Gateway configuration for the serving endpoint. *Note: only external model endpoints are supported as of now.*
* `budget_policy_id` - (Optiona) The Budget Policy ID set for this serving endpoint.

### served_entities Configuration Block

* `name` - The name of a served entity. It must be unique across an endpoint. A served entity name can consist of alphanumeric characters, dashes, and underscores. If not specified for an external model, this field will be created from the `entity_name` and `entity_version`
* `entity_name` - The full path of the UC model to be served, given in the form of `catalog_name.schema_name.model_name`.
* `entity_version` - The version of the model in UC to be served.
* `provisioned_model_units` - The number of model units to be provisioned.

### traffic_config Configuration Block

* `routes` - (Required) Each block represents a route that defines traffic to each served entity. Each `served_entity` block needs to have a corresponding `routes` block.
  * `served_entity_name` - (Required) The name of the served entity this route configures traffic for. This needs to match the name of a `served_entity` block.
  * `traffic_percentage` - (Required) The percentage of endpoint traffic to send to this route. It must be an integer between 0 and 100 inclusive.


### tags Configuration Block

* `key` - The key field for a tag.
* `value` - The value field for a tag.

### ai_gateway Configuration Block

* `guardrails` - (Optional) Block with configuration for AI Guardrails to prevent unwanted data and unsafe data in requests and responses. Consists of the following attributes:
  * `input` - A block with configuration for input guardrail filters:
    * `invalid_keywords` - List of invalid keywords. AI guardrail uses keyword or string matching to decide if the keyword exists in the request or response content.
    * `valid_topics` - The list of allowed topics. Given a chat request, this guardrail flags the request if its topic is not in the allowed topics.
    * `safety` - the boolean flag that indicates whether the safety filter is enabled.
    * `pii` - Block with configuration for guardrail PII filter:
      * `behavior` - a string that describes the behavior for PII filter. Currently only `BLOCK` value is supported.
  * `output` - A block with configuration for output guardrail filters.  Has the same structure as `input` block.
* `rate_limits` - (Optional) Block describing rate limits for AI gateway. For details see the description of `rate_limits` block above.
* `usage_tracking_config` - (Optional) Block with configuration for payload logging using inference tables. For details see the description of `auto_capture_config` block above.
* `inference_table_config` - (Optional) Block describing the configuration of usage tracking. Consists of the following attributes:
  * `enabled` - boolean flag specifying if usage tracking is enabled.

## Attribute Reference

In addition to all the arguments above, the following attributes are exported:

* `id` - Equal to the `name` argument and used to identify the serving endpoint.
* `serving_endpoint_id` - Unique identifier of the serving endpoint primarily used to set permissions and refer to this instance for other operations.

## Access Control

* [databricks_permissions](permissions.md#model-serving-usage) can control which groups or individual users can *Manage*, *Query* or *View* individual serving endpoints.

## Timeouts

The `timeouts` block allows you to specify `create` and `update` timeouts. The default right now is 10 minutes for both operations.

```hcl
timeouts {
  create = "30m"
}
```

## Import

The model serving provisioned throughput resource can be imported using the name of the endpoint.

```bash
terraform import databricks_model_serving_provisioned_throughput.this <model-serving-endpoint-name>
```


## Related Resources

The following resources are often used in the same context:

* [databricks_registered_model](registered_model.md) to create [Models in Unity Catalog](https://docs.databricks.com/en/mlflow/models-in-uc.html) in Databricks.
* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_directory](directory.md) to manage directories in [Databricks Workspace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_mlflow_model](mlflow_model.md) to create models in the [workspace model registry](https://docs.databricks.com/en/mlflow/model-registry.html) in Databricks.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_notebook](../data-sources/notebook.md) data to export a notebook from Databricks Workspace.
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
