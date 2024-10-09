---
subcategory: "Serving"
---
# databricks_model_serving Resource

This resource allows you to manage [Model Serving](https://docs.databricks.com/machine-learning/model-serving/index.html) endpoints in Databricks.

-> If you replace `served_models` with `served_entities` in an existing serving endpoint, the serving endpoint will briefly go into an update state (~30 seconds) and increment the config version.

## Example Usage

```hcl
resource "databricks_model_serving" "this" {
  name = "ads-serving-endpoint"
  config {
    served_entities {
      name                  = "prod_model"
      entity_name           = "ads-model"
      entity_version        = "2"
      workload_size         = "Small"
      scale_to_zero_enabled = true
    }
    served_entities {
      name                  = "candidate_model"
      entity_name           = "ads-model"
      entity_version        = "4"
      workload_size         = "Small"
      scale_to_zero_enabled = false
    }
    traffic_config {
      routes {
        served_model_name  = "prod_model"
        traffic_percentage = 90
      }
      routes {
        served_model_name  = "candidate_model"
        traffic_percentage = 10
      }
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the model serving endpoint. This field is required and must be unique across a workspace. An endpoint name can consist of alphanumeric characters, dashes, and underscores. NOTE: Changing this name will delete the existing endpoint and create a new endpoint with the update name.
* `config` - (Required) The model serving endpoint configuration.
  * `served_entities` - A list of served entities for the endpoint to serve. A serving endpoint can have up to 10 served entities.
  * `served_models` - (Deprecated, use `served_entities` instead) Each block represents a served model for the endpoint to serve. A model serving endpoint can have up to 10 served models.
  * `traffic_config` - A single block represents the traffic split configuration amongst the served models.
  * `auto_capture_config` - Configuration for Inference Tables which automatically logs requests and responses to Unity Catalog.
* `tags` - Tags to be attached to the serving endpoint and automatically propagated to billing logs.
* `rate_limits` - A list of rate limits to be applied to the serving endpoint. NOTE: only external and foundation model endpoints are supported as of now.
* `route_optimized` - (Optional) A boolean enabling route optimization for the endpoint. NOTE: only available for custom models.

### served_entities Configuration Block

* `name` - The name of a served entity. It must be unique across an endpoint. A served entity name can consist of alphanumeric characters, dashes, and underscores. If not specified for an external model, this field defaults to `external_model.name`, with '.' and ':' replaced with '-', and if not specified for other entities, it defaults to -.
* `external_model` - The external model to be served. NOTE: Only one of `external_model` and (`entity_name`, `entity_version`, `workload_size`, `workload_type`, and `scale_to_zero_enabled`) can be specified with the latter set being used for custom model serving for a Databricks registered model. When an `external_model` is present, the served entities list can only have one `served_entity` object. For an existing endpoint with `external_model`, it can not be updated to an endpoint without `external_model`. If the endpoint is created without `external_model`, users cannot update it to add `external_model` later.
  * `provider` - (Required) The name of the provider for the external model. Currently, the supported providers are `ai21labs`, `anthropic`, `amazon-bedrock`, `cohere`, `databricks-model-serving`, `openai`, and `palm`.
  * `name` - The name of the external model.
  * `task` - The task type of the external model.
  * `config` - The config for the external model, which must match the provider.
    * `ai21labs_config` - AI21Labs Config
      * `ai21labs_api_key` - The Databricks secret key reference for an AI21Labs API key.
    * `anthropic_config` - Anthropic Config
      * `anthropic_api_key` - The Databricks secret key reference for an Anthropic API key.
        The Databricks secret key reference for an Anthropic API key.
    * `amazon_bedrock_config` - Amazon Bedrock Config
      * `aws_region` - The AWS region to use. Bedrock has to be enabled there.
      * `aws_access_key_id` - The Databricks secret key reference for an AWS Access Key ID with permissions to interact with Bedrock services.
      * `aws_secret_access_key` - The Databricks secret key reference for an AWS Secret Access Key paired with the access key ID, with permissions to interact with Bedrock services.
      * `bedrock_provider` - The underlying provider in Amazon Bedrock. Supported values (case insensitive) include: `Anthropic`, `Cohere`, `AI21Labs`, `Amazon`.
    * `cohere_config` - Cohere Config
      * `cohere_api_key` - The Databricks secret key reference for a Cohere API key.
    * `databricks_model_serving_config` - Databricks Model Serving Config
      * `databricks_api_token` - The Databricks secret key reference for a Databricks API token that corresponds to a user or service principal with Can Query access to the model serving endpoint pointed to by this external model.
      * `databricks_workspace_url` - The URL of the Databricks workspace containing the model serving endpoint pointed to by this external model.
    * `openai_config` - OpenAI Config
      * `openai_api_key` - The Databricks secret key reference for an OpenAI or Azure OpenAI API key.
      * `openai_api_type` - This is an optional field to specify the type of OpenAI API to use. For Azure OpenAI, this field is required, and adjust this parameter to represent the preferred security access validation protocol. For access token validation, use azure. For authentication using Azure Active Directory (Azure AD) use, azuread.
      * `openai_api_base` - This is the base URL for the OpenAI API (default: "https://api.openai.com/v1"). For Azure OpenAI, this field is required, and is the base URL for the Azure OpenAI API service provided by Azure.
      * `openai_api_version` - This is an optional field to specify the OpenAI API version. For Azure OpenAI, this field is required, and is the version of the Azure OpenAI service to utilize, specified by a date.
      * `openai_organization` - This is an optional field to specify the organization in OpenAI or Azure OpenAI.
      * `openai_deployment_name` - This field is only required for Azure OpenAI and is the name of the deployment resource for the Azure OpenAI service.
    * `palm_config` - PaLM Config
      * `palm_api_key` - The Databricks secret key reference for a PaLM API key.
* `entity_name` - The name of the entity to be served. The entity may be a model in the Databricks Model Registry, a model in the Unity Catalog (UC), or a function of type `FEATURE_SPEC` in the UC. If it is a UC object, the full name of the object should be given in the form of `catalog_name.schema_name.model_name`.
* `entity_version` - The version of the model in Databricks Model Registry to be served or empty if the entity is a `FEATURE_SPEC`.
* `min_provisioned_throughput`- The minimum tokens per second that the endpoint can scale down to.
* `max_provisioned_throughput` -  The maximum tokens per second that the endpoint can scale up to.
* `workload_size` - The workload size of the served entity. The workload size corresponds to a range of provisioned concurrency that the compute autoscales between. A single unit of provisioned concurrency can process one request at a time. Valid workload sizes are `Small` (4 - 4 provisioned concurrency), `Medium` (8 - 16 provisioned concurrency), and `Large` (16 - 64 provisioned concurrency). If `scale-to-zero` is enabled, the lower bound of the provisioned concurrency for each workload size is 0.
* `workload_type` - The workload type of the served entity. The workload type selects which type of compute to use in the endpoint. The default value for this parameter is `CPU`. For deep learning workloads, GPU acceleration is available by selecting workload types like `GPU_SMALL` and others. See the available [GPU types](https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types).
* `scale_to_zero_enabled` - Whether the compute resources for the served entity should scale down to zero.
* `environment_vars` - An object containing a set of optional, user-specified environment variable key-value pairs used for serving this entity. Note: this is an experimental feature and subject to change. Example entity environment variables that refer to Databricks secrets: ```{"OPENAI_API_KEY": "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN": "{{secrets/my_scope2/my_key2}}"}```
* `instance_profile_arn` - ARN of the instance profile that the served entity uses to access AWS resources.

### served_models Configuration Block (deprecated)

-> **Deprecated** Please use `served_entities` instead of `served_models`.

* `name` - The name of a served model. It must be unique across an endpoint. If not specified, this field will default to `modelname-modelversion`. A served model name can consist of alphanumeric characters, dashes, and underscores.
* `model_name` - (Required) The name of the model in Databricks Model Registry to be served.
* `model_version` - (Required) The version of the model in Databricks Model Registry to be served.
* `workload_size` - (Required) The workload size of the served model. The workload size corresponds to a range of provisioned concurrency that the compute will autoscale between. A single unit of provisioned concurrency can process one request at a time. Valid workload sizes are `Small` (4 - 4 provisioned concurrency), `Medium` (8 - 16 provisioned concurrency), and `Large` (16 - 64 provisioned concurrency).
* `scale_to_zero_enabled` - Whether the compute resources for the served model should scale down to zero. If `scale-to-zero` is enabled, the lower bound of the provisioned concurrency for each workload size will be 0. The default value is `true`.
* `workload_type` - The workload type of the served model. The workload type selects which type of compute to use in the endpoint. For deep learning workloads, GPU acceleration is available by selecting workload types like `GPU_SMALL` and others. See documentation for all options. The default value is `CPU`.
* `environment_vars` - (Optional) a map of environment variable name/values that will be used for serving this model.  Environment variables may refer to Databricks secrets using the standard syntax: `{{secrets/secret_scope/secret_key}}`.
* `instance_profile_arn` - (Optional) ARN of the instance profile that the served model will use to access AWS resources.

### traffic_config Configuration Block

* `routes` - (Required) Each block represents a route that defines traffic to each served entity. Each `served_entity` block needs to have a corresponding `routes` block.
  * `served_entity_name` - (Required) The name of the served entity this route configures traffic for. This needs to match the name of a `served_entity` block.
  * `traffic_percentage` - (Required) The percentage of endpoint traffic to send to this route. It must be an integer between 0 and 100 inclusive.

### auto_capture_config Configuration Block

* `catalog_name` - The name of the catalog in Unity Catalog. NOTE: On update, you cannot change the catalog name if it was already set.
* `schema_name` - The name of the schema in Unity Catalog. NOTE: On update, you cannot change the schema name if it was already set.
* `table_name_prefix` - The prefix of the table in Unity Catalog. NOTE: On update, you cannot change the prefix name if it was already set.
* `enabled` - If inference tables are enabled or not. NOTE: If you have already disabled payload logging once, you cannot enable again.

### tags Configuration Block

* `key` - The key field for a tag.
* `value` - The value field for a tag.

### rate_limits Configuration Block

* `calls` - (Required) Used to specify how many calls are allowed for a key within the renewal_period.
* `key` - Key field for a serving endpoint rate limit. Currently, only `user` and `endpoint` are supported, with `endpoint` being the default if not specified.
* `renewal_period` - (Required) Renewal period field for a serving endpoint rate limit. Currently, only `minute` is supported.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Equal to the `name` argument and used to identify the serving endpoint.
* `serving_endpoint_id` - Unique identifier of the serving endpoint primarily used to set permissions and refer to this instance for other operations.

## Access Control

* [databricks_permissions](permissions.md#model-serving-usage) can control which groups or individual users can *Manage*, *Query* or *View* individual serving endpoints.

## Timeouts

The `timeouts` block allows you to specify `create` and `update` timeouts. The default right now is 45 minutes for both operations.

```hcl
timeouts {
  create = "30m"
}
```

## Import

The model serving resource can be imported using the name of the endpoint.

```bash
terraform import databricks_model_serving.this <model-serving-endpoint-name>
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
