---
subcategory: "Serving"
---
# databricks_model_serving Resource

This resource allows you to manage [Model Serving](https://docs.databricks.com/machine-learning/model-serving/index.html) endpoints in Databricks, including custom models, external models, and foundation models. For newer foundation models, including Llama 4, please use the [databricks_model_serving_provisioned_throughput](model_serving_provisioned_throughput.md) resource.

-> This resource can only be used with a workspace-level provider!

-> If you replace `served_models` with `served_entities` in an existing serving endpoint, the serving endpoint will briefly go into an update state (~30 seconds) and increment the config version.

## Example Usage

Creating a CPU serving endpoint

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

Creating a Foundation Model endpoint

```hcl
resource "databricks_model_serving" "llama" {
  name = "llama_3_2_3b_instruct"
  ai_gateway {
    usage_tracking_config {
      enabled = true
    }
  }
  config {
    served_entities {
      name                       = "meta_llama_v3_2_3b_instruct-3"
      entity_name                = "system.ai.llama_v3_2_3b_instruct"
      entity_version             = "2"
      scale_to_zero_enabled      = true
      max_provisioned_throughput = 44000
    }
  }
}
```

Creating an External Model endpoint

```hcl
resource "databricks_model_serving" "gpt_4o" {
  name = "gpt-4o-mini"
  ai_gateway {
    usage_tracking_config {
      enabled = true
    }
    rate_limits {
      calls          = 10
      key            = "endpoint"
      renewal_period = "minute"
    }
    inference_table_config {
      enabled           = true
      table_name_prefix = "gpt-4o-mini"
      catalog_name      = "ml"
      schema_name       = "ai_gateway"
    }
    guardrails {
      input {
        invalid_keywords = ["SuperSecretProject"]
        pii {
          behavior = "BLOCK"
        }
      }
      output {
        pii {
          behavior = "BLOCK"
        }
      }
    }
  }
  config {
    served_entities {
      name = "gpt-4o-mini"
      external_model {
        name     = "gpt-4o-mini"
        provider = "openai"
        task     = "llm/v1/chat"
        openai_config {
          openai_api_key = "{{secrets/llm_scope/openai_api_key}}"
        }
      }
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the model serving endpoint. This field is required and must be unique across a workspace. An endpoint name can consist of alphanumeric characters, dashes, and underscores. NOTE: Changing this name will delete the existing endpoint and create a new endpoint with the updated name.
* `config` - The model serving endpoint configuration. This is optional and can be added and modified after creation. If `config` was provided in a previous apply but is not provided in the current apply, no change to the model serving endpoint will occur. To recreate the model serving endpoint without the `config` block, the model serving endpoint must be destroyed and recreated.
  * `served_entities` - A list of served entities for the endpoint to serve. A serving endpoint can have up to 10 served entities.
  * `served_models` - (Deprecated, use `served_entities` instead) Each block represents a served model for the endpoint to serve. A model serving endpoint can have up to 10 served models.
  * `traffic_config` - A single block represents the traffic split configuration amongst the served models.
  * `auto_capture_config` - Configuration for Inference Tables which automatically logs requests and responses to Unity Catalog.
* `tags` - Tags to be attached to the serving endpoint and automatically propagated to billing logs.
* `rate_limits` - (Deprecated, use `ai_gateway` to manage rate limits) A list of rate limit blocks to be applied to the serving endpoint. *Note: only external and foundation model endpoints are supported as of now.*
* `ai_gateway` - (Optional) A block with AI Gateway configuration for the serving endpoint. *Note: only external model endpoints are supported as of now.*
* `route_optimized` - (Optional) A boolean enabling route optimization for the endpoint. *Note: only available for custom models.*
* `budget_policy_id` - (Optiona) The Budget Policy ID set for this serving endpoint.
* `description` - (Optional) The description of the model serving endpoint.
* `email_notifications` - (Optional) A block with Email notification setting.

### served_entities Configuration Block

* `name` - The name of a served entity. It must be unique across an endpoint. A served entity name can consist of alphanumeric characters, dashes, and underscores. If not specified for an external model, this field defaults to `external_model.name`, with '.' and ':' replaced with '-', and if not specified for other entities, it defaults to -.
* `external_model` - The external model to be served. NOTE: Only one of `external_model` and (`entity_name`, `entity_version`, `workload_size`, `workload_type`, and `scale_to_zero_enabled`) can be specified with the latter set being used for custom model serving for a Databricks registered model. When an `external_model` is present, the served entities list can only have one `served_entity` object. An existing endpoint with `external_model` can not be updated to an endpoint without `external_model`. If the endpoint is created without `external_model`, users cannot update it to add `external_model` later.
  * `provider` - (Required) The name of the provider for the external model. Currently, the supported providers are `ai21labs`, `anthropic`, `amazon-bedrock`, `cohere`, `databricks-model-serving`, `google-cloud-vertex-ai`, `openai`, and `palm`.
  * `name` - The name of the external model.
  * `task` - The task type of the external model.
  * `config` - The config for the external model, which must match the provider. *Note that API keys could be provided either as a reference to the Databricks Secret (parameters without `_plaintext` suffix) or in plain text (parameters with `_plaintext` suffix)!*
    * `ai21labs_config` - AI21Labs Config
      * `ai21labs_api_key` - The Databricks secret key reference for an AI21Labs API key.
      * `ai21labs_api_key_plaintext` - An AI21 Labs API key provided as a plaintext string.
    * `anthropic_config` - Anthropic Config
      * `anthropic_api_key` - The Databricks secret key reference for an Anthropic API key.
      * `anthropic_api_key_plaintext` - The Anthropic API key provided as a plaintext string.
    * `amazon_bedrock_config` - Amazon Bedrock Config
      * `aws_region` - The AWS region to use. Bedrock has to be enabled there.
      * `aws_access_key_id` - The Databricks secret key reference for an AWS Access Key ID with permissions to interact with Bedrock services.
      * `aws_access_key_id_plaintext` - An AWS access key ID with permissions to interact with Bedrock services provided as a plaintext string.
      * `aws_secret_access_key` - The Databricks secret key reference for an AWS Secret Access Key paired with the access key ID, with permissions to interact with Bedrock services.
      * `aws_secret_access_key_plaintext` -  An AWS secret access key paired with the access key ID, with permissions to interact with Bedrock services provided as a plaintext string.
      * `bedrock_provider` - The underlying provider in Amazon Bedrock. Supported values (case insensitive) include: `Anthropic`, `Cohere`, `AI21Labs`, `Amazon`.
      * `instance_profile_arn` - Optional ARN of the instance profile that the external model will use to access AWS resources. You must authenticate using an instance profile or access keys.
    * `cohere_config` - Cohere Config
      * `cohere_api_key` - The Databricks secret key reference for a Cohere API key.
      * `cohere_api_key_plaintext` - The Cohere API key provided as a plaintext string.
    * `custom_provider_config` - Custom Provider Config. Only required if the provider is 'custom'.
      * `custom_provider_url` (Required) - URL of the custom provider API.
      * `api_key_auth` - (Optional) API key authentication for the custom provider API. Conflicts with `bearer_token_auth`.
        * `key` (Required) - The name of the API key parameter used for authentication.
        * `value` (Optional) - The Databricks secret key reference for an API Key.
        * `value_plaintext` (Optional) - The API Key provided as a plaintext string.
      * `bearer_token_auth` (Optional) - bearer token authentication for the custom provider API.  Conflicts with `api_key_auth`.
        * `token` (Optional) -  The Databricks secret key reference for a token.
        * `token_plaintext` (Optional) - The token provided as a plaintext string.
    * `databricks_model_serving_config` - Databricks Model Serving Config
      * `databricks_api_token` - The Databricks secret key reference for a Databricks API token that corresponds to a user or service principal with Can Query access to the model serving endpoint pointed to by this external model.
      * `databricks_api_token_plaintext` - The Databricks API token that corresponds to a user or service principal with Can Query access to the model serving endpoint pointed to by this external model provided as a plaintext string.
      * `databricks_workspace_url` - The URL of the Databricks workspace containing the model serving endpoint pointed to by this external model.
    * `google_cloud_vertex_ai_config` - Google Cloud Vertex AI Config.
      * `private_key` - The Databricks secret key reference for a private key for the service account that has access to the Google Cloud Vertex AI Service.
      * `private_key_plaintext` - The private key for the service account that has access to the Google Cloud Vertex AI Service is provided as a plaintext secret.
      * `project_id` - This is the Google Cloud project id that the service account is associated with.
      * `region` - This is the region for the Google Cloud Vertex AI Service.
    * `openai_config` - OpenAI Config
      * `openai_api_key` - The Databricks secret key reference for an OpenAI or Azure OpenAI API key.
      * `openai_api_key_plaintext` - The OpenAI API key using the OpenAI or Azure service provided as a plaintext string.
      * `openai_api_type` - This is an optional field to specify the type of OpenAI API to use. For Azure OpenAI, this field is required, and this parameter represents the preferred security access validation protocol. For access token validation, use `azure`. For authentication using Azure Active Directory (Azure AD) use, `azuread`.
      * `microsoft_entra_client_id` - This field is only required for Azure AD OpenAI and is the Microsoft Entra Client ID.
      * `microsoft_entra_client_secret` - The Databricks secret key reference for a client secret used for Microsoft Entra ID authentication.
      * `microsoft_entra_client_secret_plaintext` - The client secret used for Microsoft Entra ID authentication provided as a plaintext string.
      * `microsoft_entra_tenant_id` - This field is only required for Azure AD OpenAI and is the Microsoft Entra Tenant ID.
      * `openai_api_base` - This is the base URL for the OpenAI API (default: "<https://api.openai.com/v1>"). For Azure OpenAI, this field is required and is the base URL for the Azure OpenAI API service provided by Azure.
      * `openai_api_version` - This is an optional field to specify the OpenAI API version. For Azure OpenAI, this field is required and is the version of the Azure OpenAI service to utilize, specified by a date.
      * `openai_organization` - This is an optional field to specify the organization in OpenAI or Azure OpenAI.
      * `openai_deployment_name` - This field is only required for Azure OpenAI and is the name of the deployment resource for the Azure OpenAI service.
    * `palm_config` - PaLM Config
      * `palm_api_key` - The Databricks secret key reference for a PaLM API key.
      * `palm_api_key_plaintext` - The PaLM API key provided as a plaintext string.
* `entity_name` - The name of the entity to be served. The entity may be a model in the Databricks Model Registry, a model in the Unity Catalog (UC), or a function of type `FEATURE_SPEC` in the UC. If it is a UC object, the full name of the object should be given in the form of `catalog_name.schema_name.model_name`.
* `entity_version` - The version of the model in Databricks Model Registry to be served or empty if the entity is a `FEATURE_SPEC`.
* `workload_size` - The workload size of the served entity. The workload size corresponds to a range of provisioned concurrency that the compute autoscales between. A single unit of provisioned concurrency can process one request at a time. Valid workload sizes are `Small` (4 - 4 provisioned concurrency), `Medium` (8 - 16 provisioned concurrency), and `Large` (16 - 64 provisioned concurrency). If `scale-to-zero` is enabled, the lower bound of the provisioned concurrency for each workload size is 0. Conflicts with `min_provisioned_concurrency` and `max_provisioned_concurrency`.
* `min_provisioned_concurrency` - The minimum provisioned concurrency that the endpoint can scale down to. Conflicts with `workload_size`.
* `max_provisioned_concurrency` - The maximum provisioned concurrency that the endpoint can scale up to. Conflicts with `workload_size`.
* `min_provisioned_throughput` - The minimum tokens per second that the endpoint can scale down to.
* `max_provisioned_throughput` - The maximum tokens per second that the endpoint can scale up to.
* `workload_type` - The workload type of the served entity. The workload type selects which type of compute to use in the endpoint. The default value for this parameter is `CPU`. For deep learning workloads, GPU acceleration is available by selecting workload types like `GPU_SMALL` and others. See the available [GPU types](https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types).
* `scale_to_zero_enabled` - Whether the compute resources for the served entity should scale down to zero.
* `environment_vars` - An object containing a set of optional, user-specified environment variable key-value pairs used for serving this entity. Note: this is an experimental feature and is subject to change. Example entity environment variables that refer to Databricks secrets: ```{"OPENAI_API_KEY": "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN": "{{secrets/my_scope2/my_key2}}"}```
* `instance_profile_arn` - ARN of the instance profile that the served entity uses to access AWS resources.

### served_models Configuration Block (deprecated)

-> **Deprecated** Please use `served_entities` instead of `served_models`.

* `name` - The name of a served model. It must be unique across an endpoint. If not specified, this field will default to `modelname-modelversion`. A served model name can consist of alphanumeric characters, dashes, and underscores.
* `model_name` - (Required) The name of the model in Databricks Model Registry to be served.
* `model_version` - (Required) The version of the model in Databricks Model Registry to be served.
* `workload_size` - (Required) The workload size of the served model. The workload size corresponds to a range of provisioned concurrency that the compute will autoscale between. A single unit of provisioned concurrency can process one request at a time. Valid workload sizes are `Small` (4 - 4 provisioned concurrency), `Medium` (8 - 16 provisioned concurrency), and `Large` (16 - 64 provisioned concurrency).
* `scale_to_zero_enabled` - Whether the compute resources for the served model should scale down to zero. If `scale-to-zero` is enabled, the lower bound of the provisioned concurrency for each workload size will be 0. The default value is `true`.
* `workload_type` - The workload type of the served model. The workload type selects which type of compute to use in the endpoint. For deep learning workloads, GPU acceleration is available by selecting workload types like `GPU_SMALL` and others. See the documentation for all options. The default value is `CPU`.
* `environment_vars` - (Optional) a map of environment variable names/values that will be used for serving this model.  Environment variables may refer to Databricks secrets using the standard syntax: `{{secrets/secret_scope/secret_key}}`.
* `instance_profile_arn` - (Optional) ARN of the instance profile that the served model will use to access AWS resources.

### traffic_config Configuration Block

* `routes` - (Required) Each block represents a route that defines traffic to each served entity. Each `served_entity` block needs to have a corresponding `routes` block.
  * `served_entity_name` - (Required) The name of the served entity this route configures traffic for. This needs to match the name of a `served_entity` block.
  * `traffic_percentage` - (Required) The percentage of endpoint traffic to send to this route. It must be an integer between 0 and 100 inclusive.

### auto_capture_config Configuration Block

* `catalog_name` - The name of the catalog in Unity Catalog. NOTE: On update, you cannot change the catalog name if it was already set.
* `schema_name` - The name of the schema in Unity Catalog. NOTE: On update, you cannot change the schema name if it was already set.
* `table_name_prefix` - The prefix of the table in Unity Catalog. NOTE: On update, you cannot change the prefix name if it was already set.
* `enabled` - If inference tables are enabled or not. NOTE: If you have already disabled payload logging once, you cannot enable it again.

### tags Configuration Block

* `key` - The key field for a tag.
* `value` - The value field for a tag.

### rate_limits Configuration Block

* `calls` - (Required) Used to specify how many calls are allowed for a key within the renewal_period.
* `key` - (Optional) Key field for a serving endpoint rate limit. Currently, `user`, `user_group`, `service_principal`, and `endpoint` are supported, with `endpoint` being the default if not specified.
* `principal` - (Optional) Principal field for a user, user group, or service principal to apply rate limiting to. Accepts a user email, group name, or service principal application ID.
* `renewal_period` - (Required) Renewal period field for a serving endpoint rate limit. Currently, only `minute` is supported.
* `tokens` - (Optional, int) Specifies how many tokens are allowed for a key within the renewal_period.

### ai_gateway Configuration Block

* `fallback_config` - (Optional) block with configuration for traffic fallback which auto fallbacks to other served entities if the request to a served entity fails with certain error codes, to increase availability.
  * `enabled` -  Whether to enable traffic fallback. When a served entity in the serving endpoint returns specific error codes (e.g. 500), the request will automatically be round-robin attempted with other served entities in the same endpoint, following the order of served entity list, until a successful response is returned.
* `guardrails` - (Optional) Block with configuration for AI Guardrails to prevent unwanted data and unsafe data in requests and responses. Consists of the following attributes:
  * `input` - A block with configuration for input guardrail filters:
    * `invalid_keywords` - (Deprecated) List of invalid keywords. AI guardrail uses keyword or string matching to decide if the keyword exists in the request or response content.
    * `valid_topics` - (Deprecated) The list of allowed topics. Given a chat request, this guardrail flags the request if its topic is not in the allowed topics.
    * `safety` - the boolean flag that indicates whether the safety filter is enabled.
    * `pii` - Block with configuration for guardrail PII filter:
      * `behavior` - a string that describes the behavior for PII filter. Currently only `BLOCK` value is supported.
  * `output` - A block with configuration for output guardrail filters.  Has the same structure as `input` block.
* `rate_limits` - (Optional) Block describing rate limits for AI gateway. For details see the description of `rate_limits` block above.
* `usage_tracking_config` - (Optional) Block with configuration for payload logging using inference tables. For details see the description of `auto_capture_config` block above.
* `inference_table_config` - (Optional) Block describing the configuration of usage tracking. Consists of the following attributes:
  * `enabled` - boolean flag specifying if usage tracking is enabled.

### email_notifications Block

* `on_update_failure` - (Optional) a list of email addresses to be notified when an endpoint fails to update its configuration or state.
* `on_update_success` - (Optional) a list of email addresses to be notified when an endpoint successfully updates its configuration or state.

## Attribute Reference

In addition to all the arguments above, the following attributes are exported:

* `id` - Equal to the `name` argument and used to identify the serving endpoint.
* `serving_endpoint_id` - Unique identifier of the serving endpoint primarily used to set permissions and refer to this instance for other operations.
* `endpoint_url` - Invocation url of the endpoint.

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

```hcl
import {
  to = databricks_model_serving.this
  id = "<model-serving-endpoint-name>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_model_serving.this <model-serving-endpoint-name>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_model_serving_provisioned_throughput](model_serving.md) to create [Foundation Model provisioned throughput](https://docs.databricks.com/aws/en/machine-learning/foundation-model-apis/deploy-prov-throughput-foundation-model-apis) endpoints in Databricks.
* [databricks_registered_model](registered_model.md) to create [Models in Unity Catalog](https://docs.databricks.com/en/mlflow/models-in-uc.html) in Databricks.
* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_directory](directory.md) to manage directories in [Databricks Workspace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_mlflow_model](mlflow_model.md) to create models in the [workspace model registry](https://docs.databricks.com/en/mlflow/model-registry.html) in Databricks.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_notebook](../data-sources/notebook.md) data to export a notebook from Databricks Workspace.
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
